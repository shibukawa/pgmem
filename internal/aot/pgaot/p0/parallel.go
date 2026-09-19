package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
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
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
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
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v99 != 0 {
		goto L23
	} else {
		goto L24
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
	v87 = v21
	goto L8
L8:
	;
	v91 = v20 + int32(1)
	if v91 < v87 {
		v20 = v91
		v21 = v87
		goto L4
	} else {
		goto L22
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
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v77+v26)+4))
	F_shm_mq_detach(m, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L9
	} else {
		goto L21
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
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_DestroyParallelContext[1]))
	F_LWLockRelease(m, v72+int32(_a_F_DestroyParallelContext_0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L9
	} else {
		goto L20
	}
L15:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_DestroyParallelContext[2])))
	if v57 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L11
L17:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_DestroyParallelContext[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v61+int32(24)))) = int32(1)
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_DestroyParallelContext[4]))
	v70 = F_pgmem_kill(m, v68, int32(10))
	mBase = m.M
	goto L19
L18:
	;
	goto L19
L19:
	;
	goto L16
L20:
	;
	goto L11
L21:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v82+v26)+4)) = int32(0)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v87 = v86
	goto L8
L22:
	;
	goto L5
L23:
	;
	F_dsm_detach(m, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L9
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v104 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(0)
	goto L25
L27:
	;
	F_pfree(m, v104)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L9
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v109 = int32(_a_F_DestroyParallelContext_1)
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_DestroyParallelContext[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_DestroyParallelContext[5])) = v111 + int32(1)
	F_WaitForParallelWorkersToExit(m, l0)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L9
	} else {
		goto L31
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
	goto L29
L31:
	;
	v117 = int32(_a_F_DestroyParallelContext_1)
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_DestroyParallelContext[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_DestroyParallelContext[5])) = v119 - int32(1)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v123 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	F_pfree(m, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L9
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_pfree(m, v128)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L9
	} else {
		goto L36
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = int32(0)
	goto L34
L36:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	F_pfree(m, v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L9
	} else {
		goto L37
	}
L37:
	;
	F_pfree(m, l0)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L9
	} else {
		goto L38
	}
L38:
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
	var v134 int32
	_ = v134
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
	var v193 int32
	_ = v193
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
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
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
	var v399 int32
	_ = v399
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
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v500 int32
	_ = v500
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v598 int32
	_ = v598
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v622 int32
	_ = v622
	var v628 int32
	_ = v628
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v647 int32
	_ = v647
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v688 int32
	_ = v688
	var v693 int32
	_ = v693
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v826 int32
	_ = v826
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v840 int32
	_ = v840
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v860 int32
	_ = v860
	var v869 int32
	_ = v869
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v922 int32
	_ = v922
	var v932 int32
	_ = v932
	var v935 int32
	_ = v935
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
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
	v973 = F_BarrierArriveAndWait(m, v23, int32(134217751))
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L7
	} else {
		goto L172
	}
L3:
	;
	v803 = F_BarrierArriveAndWait(m, v23, int32(134217749))
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L7
	} else {
		goto L139
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
	v134 = int32(0)
	goto L45
L43:
	;
	v174 = v84
	goto L44
L44:
	;
	v191 = v174
	v193 = int32(0)
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
	v169 = v134 + v166
	if v169 != v108&int32(-8) {
		v132 = v167
		v134 = v169
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
	v214 = v193 + v211
	if v214 != v126 {
		v191 = v191 + v211
		v193 = v214
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
	v373 = v363
	v377 = v367
	goto L72
L70:
	;
	v567 = v363
	goto L71
L71:
	;
	F_LWLockRelease(m, v567)
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L7
	} else {
		goto L104
	}
L72:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v387 = F_dsa_get_address(m, v386, v377)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L7
	} else {
		goto L74
	}
L73:
	;
	v567 = v557
	goto L71
L74:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v387)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v371))) = v389
	F_LWLockRelease(m, v373)
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
	v399 = int32(0)
	goto L79
L77:
	;
	goto L78
L78:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	F_dsa_free(m, v548, v377)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L7
	} else {
		goto L97
	}
L79:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v414 = v399 + (v387 + int32(16))
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
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v511)+20))
	v513 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v511)+20)) = v512 + v513
	v516 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v519 = v516 + v500*int32(36)
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v519)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v519)+8)) = v520 + v513
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v419)))
	v529 = (v524+int32(15))&int32(-8) + v399
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v387)+8))
	if base.Ui32(v529) < base.Ui32(v530) {
		v399 = v529
		goto L79
	} else {
		goto L96
	}
L82:
	;
	v479 = v427 * int32(36)
	v480 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v481 = v479 + v480
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v481)+16))
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v419)))
	*(*int32)(unsafe.Add(mBase, uint32(v481)+16)) = v482 + (v483+int32(15))&int32(-8)
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v490+v479)+28))
	F_sts_puttuple(m, v492, v417, v419)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L7
	} else {
		goto L95
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
	v452 = int32(0)
	v454 = base.AtomicRmwCmpxchg32(m, v449, v452, v450, v445)
	if v450 == v454 {
		v500 = v452
		goto L81
	} else {
		goto L91
	}
L91:
	;
	v459 = v454
	goto L92
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v437))) = v459
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v449)))
	*(*int32)(unsafe.Add(mBase, uint32(v437))) = v473
	v476 = base.AtomicRmwCmpxchg32(m, v449, int32(0), v473, v445)
	if v473 != v476 {
		v459 = v476
		goto L92
	} else {
		goto L94
	}
L93:
	;
	v500 = v452
	goto L81
L94:
	;
	goto L93
L95:
	;
	v500 = v427
	goto L81
L96:
	;
	goto L80
L97:
	;
	v552 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashIncreaseNumBatches[2]))
	if v552 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L7
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v557 = v555 + int32(40)
	v559 = F_LWLockAcquire(m, v557, int32(0))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L7
	} else {
		goto L102
	}
L101:
	;
	goto L100
L102:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v555)+24))
	if v563 != 0 {
		v371 = v555 + int32(24)
		v373 = v557
		v377 = v563
		goto L72
	} else {
		goto L103
	}
L103:
	;
	goto L73
L104:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v582)+12))
	v584 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v582)+4))
	v586 = F_dsa_get_address(m, v584, v585)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L7
	} else {
		goto L105
	}
L105:
	;
	v590 = F_palloc0(m, v583<<(uint(int32(2))%32))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L7
	} else {
		goto L106
	}
L106:
	;
	if int32(2) <= v583 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v598 = int32(1)
	goto L110
L108:
	;
	goto L109
L109:
	;
	F_pfree(m, v590)
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L7
	} else {
		goto L136
	}
L110:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v616)+28))
	goto L112
L111:
	;
	v647 = int32(1)
	goto L115
L112:
	;
	v622 = int32(1)
	v628 = int32(-64)
	v635 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashIncreaseNumBatches[3]))
	v638 = F_sts_attach(m, v586+(((v617*int32(28)+int32(76))<<(uint(v622)%32)+int32(14))&int32(-16)-v628)*v598-v628, v635+v622, v582+int32(168))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L7
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v590+v598<<(uint(int32(2))%32)))) = v638
	v642 = v598 + int32(1)
	if v642 != v583 {
		v598 = v642
		goto L110
	} else {
		goto L114
	}
L114:
	;
	goto L111
L115:
	;
	v663 = v590 + v647<<(uint(int32(2))%32)
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v663)))
	F_sts_begin_parallel_scan(m, v664)
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L7
	} else {
		goto L117
	}
L116:
	;
	goto L109
L117:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v663)))
	v670 = F_sts_parallel_scan_next(m, v667, v19+int32(12))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L7
	} else {
		goto L118
	}
L118:
	;
	if v670 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v673 = v670
	goto L122
L120:
	;
	goto L121
L121:
	;
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v663)))
	F_sts_end_parallel_scan(m, v757)
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L7
	} else {
		goto L134
	}
L122:
	;
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v673)))
	v693 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(int32(2)) <= base.Ui32(v693) {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	goto L121
L124:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v699 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v703 = (v693 - int32(1)) & base.I32_rotr(v698, v699)
	goto L126
L125:
	;
	v703 = int32(0)
	goto L126
L126:
	;
	v704 = int32(36)
	v705 = v703 * v704
	v706 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v707 = v705 + v706
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v707)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v707)+16)) = v708 + (v688+int32(15))&int32(-8)
	v711 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v712 = v711 + v705
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v712)+8))
	v714 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v712)+8)) = v713 + v714
	v717 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v720 = v717 + v647*v704
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v720)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v720)+20)) = v721 + v714
	v725 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v725+v705)+28))
	F_sts_puttuple(m, v727, v19+int32(12), v673)
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L7
	} else {
		goto L127
	}
L127:
	;
	v733 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashIncreaseNumBatches[2]))
	if v733 != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L7
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v663)))
	v739 = F_sts_parallel_scan_next(m, v736, v19+int32(12))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L7
	} else {
		goto L132
	}
L131:
	;
	goto L130
L132:
	;
	if v739 != 0 {
		v673 = v739
		goto L122
	} else {
		goto L133
	}
L133:
	;
	goto L123
L134:
	;
	v761 = v647 + int32(1)
	if v761 != v583 {
		v647 = v761
		goto L115
	} else {
		goto L135
	}
L135:
	;
	goto L116
L136:
	;
	F_ExecParallelHashMergeCounters(m, l0)
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L7
	} else {
		goto L137
	}
L137:
	;
	v784 = F_BarrierArriveAndWait(m, v23, int32(134217753))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L7
	} else {
		goto L138
	}
L138:
	;
	goto L3
L139:
	;
	if v803 == int32(0) {
		goto L2
	} else {
		goto L140
	}
L140:
	;
	F_ExecParallelHashEnsureBatchAccessors(m, l0)
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L7
	} else {
		goto L141
	}
L141:
	;
	v809 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v809
	v812 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v813 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v813)))
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v814)))
	v816 = F_dsa_get_address(m, v812, v815)
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L7
	} else {
		goto L142
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v816
	v819 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v819)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v820
	v823 = int32(1073741823)
	if v823 <= v820 {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	v835 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v835
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v835
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v834
	v840 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	*(*uint8)(unsafe.Add(mBase, uint32(v840)+24)) = uint8(v835)
	v843 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v845 = F_dsa_get_address(m, v843, v844)
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L7
	} else {
		goto L150
	}
L144:
	;
	v826 = v823
	goto L146
L145:
	;
	v826 = v820
	goto L146
L146:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v826) {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v834 = int32(32) - base.I32_clz(v826-int32(1))
	goto L149
L148:
	;
	v834 = int32(0)
	goto L149
L149:
	;
	goto L143
L150:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v847 <= int32(0) {
		v935 = v809
		goto L151
	} else {
		goto L152
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v935
	v950 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	F_dsa_free(m, v950, v951)
	mBase = m.M
	v953 = m.ExcPending
	if v953 != 0 {
		goto L7
	} else {
		goto L171
	}
L152:
	;
	v850 = int32(0)
	v854 = v850
	v856 = v850
	v860 = v850
	goto L153
L153:
	;
	v869 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v869+v854*int32(36))))
	v874 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v873)+60)))
	if v874 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L154:
	;
	v922 = v915 | base.B2i32(int32(1073741822) < v918)
	if (v922|v881)&int32(1) == int32(0) {
		v935 = v809
		goto L151
	} else {
		goto L167
	}
L155:
	;
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v883 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v883)+28))
	goto L161
L156:
	;
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v873)+48))
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v21)+32))
	if base.Ui32(v877) <= base.Ui32(v878) {
		v881 = v860
		goto L155
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	v881 = int32(1)
	goto L155
L159:
	;
	goto L158
L160:
	;
	v917 = v854 + int32(1)
	v918 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v917 < v918 {
		v854 = v917
		v856 = v915
		v860 = v881
		goto L153
	} else {
		goto L166
	}
L161:
	;
	v897 = base.I32_rem_s(v854, v882)
	v900 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v845+(((v884*int32(28)+int32(76))<<(uint(int32(1))%32)+int32(14))&int32(-16)-int32(-64))*v897)+60)))
	if v900 == int32(0) {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v873)+48))
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v21)+32))
	if base.Ui32(v903) <= base.Ui32(v904) {
		v915 = v856
		goto L160
	} else {
		goto L165
	}
L163:
	;
	goto L164
L164:
	;
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v873)+52))
	v907 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v907+v897*int32(36))))
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v911)+56))
	v915 = base.B2i32(v906 == v912) | v856
	goto L160
L165:
	;
	goto L164
L166:
	;
	goto L154
L167:
	;
	if v922&int32(1) != 0 {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v932 = int32(3)
	goto L170
L169:
	;
	v932 = int32(2)
	goto L170
L170:
	;
	v935 = v932
	goto L151
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = int32(0)
	goto L2
L172:
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
	var v303 int32
	_ = v303
	var v311 int32
	_ = v311
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int64
	_ = v332
	var v334 int32
	_ = v334
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int64
	_ = v347
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v367 int64
	_ = v367
	var v371 int32
	_ = v371
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int64
	_ = v465
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int64
	_ = v493
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int64
	_ = v521
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int64
	_ = v549
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int64
	_ = v577
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
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
				v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
				mBase = m.M
				v591 = m.ExcPending
				if v591 != 0 {
					return int32(0)
				} else {
					return v590
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
						v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
						mBase = m.M
						v591 = m.ExcPending
						if v591 != 0 {
							return int32(0)
						} else {
							return v590
						}
					}
				}
			}
		default:
			v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
			mBase = m.M
			v591 = m.ExcPending
			if v591 != 0 {
				return int32(0)
			} else {
				return v590
			}
		case 6:
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+36)))
			if v29 != int32(1) {
				v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
				mBase = m.M
				v591 = m.ExcPending
				if v591 != 0 {
					return int32(0)
				} else {
					return v590
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
								v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
								mBase = m.M
								v591 = m.ExcPending
								if v591 != 0 {
									return int32(0)
								} else {
									return v590
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
				v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
				mBase = m.M
				v591 = m.ExcPending
				if v591 != 0 {
					return int32(0)
				} else {
					return v590
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
								v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
								mBase = m.M
								v591 = m.ExcPending
								if v591 != 0 {
									return int32(0)
								} else {
									return v590
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
											v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
											mBase = m.M
											v591 = m.ExcPending
											if v591 != 0 {
												return int32(0)
											} else {
												return v590
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
												v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
												mBase = m.M
												v591 = m.ExcPending
												if v591 != 0 {
													return int32(0)
												} else {
													return v590
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
											v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
											mBase = m.M
											v591 = m.ExcPending
											if v591 != 0 {
												return int32(0)
											} else {
												return v590
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
				v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
				mBase = m.M
				v591 = m.ExcPending
				if v591 != 0 {
					return int32(0)
				} else {
					return v590
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
								v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
								mBase = m.M
								v591 = m.ExcPending
								if v591 != 0 {
									return int32(0)
								} else {
									return v590
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
											v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
											mBase = m.M
											v591 = m.ExcPending
											if v591 != 0 {
												return int32(0)
											} else {
												return v590
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
												v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
												mBase = m.M
												v591 = m.ExcPending
												if v591 != 0 {
													return int32(0)
												} else {
													return v590
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
											v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
											mBase = m.M
											v591 = m.ExcPending
											if v591 != 0 {
												return int32(0)
											} else {
												return v590
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
				v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
				mBase = m.M
				v591 = m.ExcPending
				if v591 != 0 {
					return int32(0)
				} else {
					return v590
				}
			} else {
				v174 = *(*int32)(unsafe.Add(mBase, uint32(v170)+12))
				if v174 == int32(0) {
					v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
					mBase = m.M
					v591 = m.ExcPending
					if v591 != 0 {
						return int32(0)
					} else {
						return v590
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
							v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
							mBase = m.M
							v591 = m.ExcPending
							if v591 != 0 {
								return int32(0)
							} else {
								return v590
							}
						}
					}
				}
			}
		case 11:
			v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266)+36)))
			if v267 != int32(1) {
				v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
				mBase = m.M
				v591 = m.ExcPending
				if v591 != 0 {
					return int32(0)
				} else {
					return v590
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
							v303 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v292))) = v303
							atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v292)+4)), uint32(v303))
							*(*int32)(unsafe.Add(mBase, uint32(v292)+8)) = v303
							v311 = v292 + int32(12)
							atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v311))), uint32(v303))
							*(*int64)(unsafe.Add(mBase, uint32(v311)+4)) = int64(-1)
							if v302 == int32(0) {
							} else {
								v319 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v302))) = v319
								v322 = v319 << (uint(int32(4)) % 32)
								if v322 == int32(0) {
								} else {
									base.MemoryFill(m, v302+int32(8), int32(0), v322)
								}
							}
							v330 = *(*int32)(unsafe.Add(mBase, uint32(v270)+52))
							v331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v332 = int64(*(*int32)(unsafe.Add(mBase, uint32(v331)+40)))
							F_shm_toc_insert(m, v330, v332, v292)
							mBase = m.M
							v334 = m.ExcPending
							if v334 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v302
								*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v292
								v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
								mBase = m.M
								v591 = m.ExcPending
								if v591 != 0 {
									return int32(0)
								} else {
									return v590
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
								v303 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v292))) = v303
								atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v292)+4)), uint32(v303))
								*(*int32)(unsafe.Add(mBase, uint32(v292)+8)) = v303
								v311 = v292 + int32(12)
								atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v311))), uint32(v303))
								*(*int64)(unsafe.Add(mBase, uint32(v311)+4)) = int64(-1)
								if v302 == int32(0) {
								} else {
									v319 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
									*(*int32)(unsafe.Add(mBase, uint32(v302))) = v319
									v322 = v319 << (uint(int32(4)) % 32)
									if v322 == int32(0) {
									} else {
										base.MemoryFill(m, v302+int32(8), int32(0), v322)
									}
								}
								v330 = *(*int32)(unsafe.Add(mBase, uint32(v270)+52))
								v331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v332 = int64(*(*int32)(unsafe.Add(mBase, uint32(v331)+40)))
								F_shm_toc_insert(m, v330, v332, v292)
								mBase = m.M
								v334 = m.ExcPending
								if v334 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v302
									*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v292
									v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
									mBase = m.M
									v591 = m.ExcPending
									if v591 != 0 {
										return int32(0)
									} else {
										return v590
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
											v303 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v292))) = v303
											atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v292)+4)), uint32(v303))
											*(*int32)(unsafe.Add(mBase, uint32(v292)+8)) = v303
											v311 = v292 + int32(12)
											atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v311))), uint32(v303))
											*(*int64)(unsafe.Add(mBase, uint32(v311)+4)) = int64(-1)
											if v302 == int32(0) {
											} else {
												v319 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
												*(*int32)(unsafe.Add(mBase, uint32(v302))) = v319
												v322 = v319 << (uint(int32(4)) % 32)
												if v322 == int32(0) {
												} else {
													base.MemoryFill(m, v302+int32(8), int32(0), v322)
												}
											}
											v330 = *(*int32)(unsafe.Add(mBase, uint32(v270)+52))
											v331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											v332 = int64(*(*int32)(unsafe.Add(mBase, uint32(v331)+40)))
											F_shm_toc_insert(m, v330, v332, v292)
											mBase = m.M
											v334 = m.ExcPending
											if v334 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v302
												*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v292
												v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
												mBase = m.M
												v591 = m.ExcPending
												if v591 != 0 {
													return int32(0)
												} else {
													return v590
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
					mBase = m.M
					v591 = m.ExcPending
					if v591 != 0 {
						return int32(0)
					} else {
						return v590
					}
				}
			}
		case 21:
			v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+36)))
			if v199 != int32(1) {
				v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
				mBase = m.M
				v591 = m.ExcPending
				if v591 != 0 {
					return int32(0)
				} else {
					return v590
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
								v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
								mBase = m.M
								v591 = m.ExcPending
								if v591 != 0 {
									return int32(0)
								} else {
									return v590
								}
							}
						}
					}
				} else {
					v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
					mBase = m.M
					v591 = m.ExcPending
					if v591 != 0 {
						return int32(0)
					} else {
						return v590
					}
				}
			}
		case 22:
			v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245)+36)))
			if v246 != int32(1) {
				v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
				mBase = m.M
				v591 = m.ExcPending
				if v591 != 0 {
					return int32(0)
				} else {
					return v590
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
								v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
								mBase = m.M
								v591 = m.ExcPending
								if v591 != 0 {
									return int32(0)
								} else {
									return v590
								}
							}
						}
					}
				} else {
					v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
					mBase = m.M
					v591 = m.ExcPending
					if v591 != 0 {
						return int32(0)
					} else {
						return v590
					}
				}
			}
		case 26:
			v340 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340)+36)))
			if v341 != int32(1) {
				v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
				mBase = m.M
				v591 = m.ExcPending
				if v591 != 0 {
					return int32(0)
				} else {
					return v590
				}
			} else {
				v344 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v345 = *(*int32)(unsafe.Add(mBase, uint32(v344)+44))
				if v345 != 0 {
					v346 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v347 = int64(*(*int32)(unsafe.Add(mBase, uint32(v346)+40)))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(632)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(720)
					v352 = *(*int32)(unsafe.Add(mBase, uint32(v344)+52))
					v354 = F_shm_toc_allocate(m, v352, int32(220))
					mBase = m.M
					v355 = m.ExcPending
					if v355 != 0 {
						return int32(0)
					} else {
						v356 = *(*int32)(unsafe.Add(mBase, uint32(v344)+52))
						F_shm_toc_insert(m, v356, v347, v354)
						mBase = m.M
						v358 = m.ExcPending
						if v358 != 0 {
							return int32(0)
						} else {
							v359 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v354)+32)) = v359
							*(*int32)(unsafe.Add(mBase, uint32(v354)+8)) = v359
							*(*int32)(unsafe.Add(mBase, uint32(v354)+164)) = v359
							*(*int32)(unsafe.Add(mBase, uint32(v354)+24)) = v359
							v367 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v354)+16)) = v367
							*(*int64)(unsafe.Add(mBase, uint32(v354))) = v367
							v371 = *(*int32)(unsafe.Add(mBase, uint32(v344)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v354)+36)) = v359
							*(*int32)(unsafe.Add(mBase, uint32(v354)+28)) = v371 + int32(1)
							v378 = v354 + int32(40)
							v379 = int32(69)
							*(*uint16)(unsafe.Add(mBase, uint32(v378))) = uint16(v379)
							*(*int32)(unsafe.Add(mBase, uint32(v378)+4)) = int32(1073741824)
							*(*int64)(unsafe.Add(mBase, uint32(v378)+8)) = int64(-1)
							v386 = v354 + int32(56)
							v387 = int32(0)
							atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v386))), uint32(v387))
							*(*int32)(unsafe.Add(mBase, uint32(v386)+8)) = v387
							*(*uint8)(unsafe.Add(mBase, uint32(v386)+20)) = uint8(v387)
							*(*int64)(unsafe.Add(mBase, uint32(v386)+12)) = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(v386)+4)) = v387
							F_ConditionVariableInit(m, v354+int32(80))
							mBase = m.M
							v402 = v354 + int32(92)
							v403 = int32(0)
							atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v402))), uint32(v403))
							*(*int32)(unsafe.Add(mBase, uint32(v402)+8)) = v403
							*(*uint8)(unsafe.Add(mBase, uint32(v402)+20)) = uint8(v403)
							*(*int64)(unsafe.Add(mBase, uint32(v402)+12)) = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(v402)+4)) = v403
							F_ConditionVariableInit(m, v354+int32(116))
							mBase = m.M
							v418 = v354 + int32(128)
							v419 = int32(0)
							atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v418))), uint32(v419))
							*(*int32)(unsafe.Add(mBase, uint32(v418)+8)) = v419
							*(*uint8)(unsafe.Add(mBase, uint32(v418)+20)) = uint8(v419)
							*(*int64)(unsafe.Add(mBase, uint32(v418)+12)) = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(v418)+4)) = v419
							F_ConditionVariableInit(m, v354+int32(152))
							mBase = m.M
							v435 = *(*int32)(unsafe.Add(mBase, uint32(v344)+44))
							F_SharedFileSetInit(m, v354+int32(168), v435)
							mBase = m.M
							v437 = m.ExcPending
							if v437 != 0 {
								return int32(0)
							} else {
								v438 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								*(*int32)(unsafe.Add(mBase, uint32(v438)+128)) = v354
								v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
								mBase = m.M
								v591 = m.ExcPending
								if v591 != 0 {
									return int32(0)
								} else {
									return v590
								}
							}
						}
					}
				} else {
					v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
					mBase = m.M
					v591 = m.ExcPending
					if v591 != 0 {
						return int32(0)
					} else {
						return v590
					}
				}
			}
		case 28:
			v555 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v556 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v556 == int32(0) {
				v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
				mBase = m.M
				v591 = m.ExcPending
				if v591 != 0 {
					return int32(0)
				} else {
					return v590
				}
			} else {
				v559 = *(*int32)(unsafe.Add(mBase, uint32(v555)+12))
				if v559 == int32(0) {
					v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
					mBase = m.M
					v591 = m.ExcPending
					if v591 != 0 {
						return int32(0)
					} else {
						return v590
					}
				} else {
					v562 = *(*int32)(unsafe.Add(mBase, uint32(v555)+52))
					v566 = v559*int32(40) + int32(8)
					v567 = F_shm_toc_allocate(m, v562, v566)
					mBase = m.M
					v568 = m.ExcPending
					if v568 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = v567
						if v566 != 0 {
							base.MemoryFill(m, v567, int32(0), v566)
						} else {
						}
						v572 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
						v573 = *(*int32)(unsafe.Add(mBase, uint32(v555)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v572))) = v573
						v575 = *(*int32)(unsafe.Add(mBase, uint32(v555)+52))
						v576 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v577 = int64(*(*int32)(unsafe.Add(mBase, uint32(v576)+40)))
						v578 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
						F_shm_toc_insert(m, v575, v577, v578)
						mBase = m.M
						v580 = m.ExcPending
						if v580 != 0 {
							return int32(0)
						} else {
							v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
							mBase = m.M
							v591 = m.ExcPending
							if v591 != 0 {
								return int32(0)
							} else {
								return v590
							}
						}
					}
				}
			}
		case 29:
			v471 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v472 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v472 == int32(0) {
				v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
				mBase = m.M
				v591 = m.ExcPending
				if v591 != 0 {
					return int32(0)
				} else {
					return v590
				}
			} else {
				v475 = *(*int32)(unsafe.Add(mBase, uint32(v471)+12))
				if v475 == int32(0) {
					v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
					mBase = m.M
					v591 = m.ExcPending
					if v591 != 0 {
						return int32(0)
					} else {
						return v590
					}
				} else {
					v478 = *(*int32)(unsafe.Add(mBase, uint32(v471)+52))
					v482 = v475<<(uint(int32(4))%32) | int32(8)
					v483 = F_shm_toc_allocate(m, v478, v482)
					mBase = m.M
					v484 = m.ExcPending
					if v484 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v483
						if v482 != 0 {
							base.MemoryFill(m, v483, int32(0), v482)
						} else {
						}
						v488 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
						v489 = *(*int32)(unsafe.Add(mBase, uint32(v471)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v488))) = v489
						v491 = *(*int32)(unsafe.Add(mBase, uint32(v471)+52))
						v492 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v493 = int64(*(*int32)(unsafe.Add(mBase, uint32(v492)+40)))
						v494 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
						F_shm_toc_insert(m, v491, v493, v494)
						mBase = m.M
						v496 = m.ExcPending
						if v496 != 0 {
							return int32(0)
						} else {
							v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
							mBase = m.M
							v591 = m.ExcPending
							if v591 != 0 {
								return int32(0)
							} else {
								return v590
							}
						}
					}
				}
			}
		case 30:
			v499 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v500 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v500 == int32(0) {
				v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
				mBase = m.M
				v591 = m.ExcPending
				if v591 != 0 {
					return int32(0)
				} else {
					return v590
				}
			} else {
				v503 = *(*int32)(unsafe.Add(mBase, uint32(v499)+12))
				if v503 == int32(0) {
					v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
					mBase = m.M
					v591 = m.ExcPending
					if v591 != 0 {
						return int32(0)
					} else {
						return v590
					}
				} else {
					v506 = *(*int32)(unsafe.Add(mBase, uint32(v499)+52))
					v510 = v503*int32(96) | int32(8)
					v511 = F_shm_toc_allocate(m, v506, v510)
					mBase = m.M
					v512 = m.ExcPending
					if v512 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+284)) = v511
						if v510 != 0 {
							base.MemoryFill(m, v511, int32(0), v510)
						} else {
						}
						v516 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
						v517 = *(*int32)(unsafe.Add(mBase, uint32(v499)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v516))) = v517
						v519 = *(*int32)(unsafe.Add(mBase, uint32(v499)+52))
						v520 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v521 = int64(*(*int32)(unsafe.Add(mBase, uint32(v520)+40)))
						v522 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
						F_shm_toc_insert(m, v519, v521, v522)
						mBase = m.M
						v524 = m.ExcPending
						if v524 != 0 {
							return int32(0)
						} else {
							v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
							mBase = m.M
							v591 = m.ExcPending
							if v591 != 0 {
								return int32(0)
							} else {
								return v590
							}
						}
					}
				}
			}
		case 32:
			v527 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v528 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v528 == int32(0) {
				v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
				mBase = m.M
				v591 = m.ExcPending
				if v591 != 0 {
					return int32(0)
				} else {
					return v590
				}
			} else {
				v531 = *(*int32)(unsafe.Add(mBase, uint32(v527)+12))
				if v531 == int32(0) {
					v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
					mBase = m.M
					v591 = m.ExcPending
					if v591 != 0 {
						return int32(0)
					} else {
						return v590
					}
				} else {
					v534 = *(*int32)(unsafe.Add(mBase, uint32(v527)+52))
					v538 = v531*int32(24) + int32(8)
					v539 = F_shm_toc_allocate(m, v534, v538)
					mBase = m.M
					v540 = m.ExcPending
					if v540 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+352)) = v539
						if v538 != 0 {
							base.MemoryFill(m, v539, int32(0), v538)
						} else {
						}
						v544 = *(*int32)(unsafe.Add(mBase, uint32(l0)+352))
						v545 = *(*int32)(unsafe.Add(mBase, uint32(v527)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v544))) = v545
						v547 = *(*int32)(unsafe.Add(mBase, uint32(v527)+52))
						v548 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v549 = int64(*(*int32)(unsafe.Add(mBase, uint32(v548)+40)))
						v550 = *(*int32)(unsafe.Add(mBase, uint32(l0)+352))
						F_shm_toc_insert(m, v547, v549, v550)
						mBase = m.M
						v552 = m.ExcPending
						if v552 != 0 {
							return int32(0)
						} else {
							v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
							mBase = m.M
							v591 = m.ExcPending
							if v591 != 0 {
								return int32(0)
							} else {
								return v590
							}
						}
					}
				}
			}
		case 37:
			v443 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v444 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v444 == int32(0) {
				v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
				mBase = m.M
				v591 = m.ExcPending
				if v591 != 0 {
					return int32(0)
				} else {
					return v590
				}
			} else {
				v447 = *(*int32)(unsafe.Add(mBase, uint32(v443)+12))
				if v447 == int32(0) {
					v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
					mBase = m.M
					v591 = m.ExcPending
					if v591 != 0 {
						return int32(0)
					} else {
						return v590
					}
				} else {
					v450 = *(*int32)(unsafe.Add(mBase, uint32(v443)+52))
					v454 = v447*int32(20) + int32(4)
					v455 = F_shm_toc_allocate(m, v450, v454)
					mBase = m.M
					v456 = m.ExcPending
					if v456 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v455
						if v454 != 0 {
							base.MemoryFill(m, v455, int32(0), v454)
						} else {
						}
						v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
						v461 = *(*int32)(unsafe.Add(mBase, uint32(v443)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v460))) = v461
						v463 = *(*int32)(unsafe.Add(mBase, uint32(v443)+52))
						v464 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v465 = int64(*(*int32)(unsafe.Add(mBase, uint32(v464)+40)))
						v466 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
						F_shm_toc_insert(m, v463, v465, v466)
						mBase = m.M
						v468 = m.ExcPending
						if v468 != 0 {
							return int32(0)
						} else {
							v590 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
							mBase = m.M
							v591 = m.ExcPending
							if v591 != 0 {
								return int32(0)
							} else {
								return v590
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
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int64
	_ = v79
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
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
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int64
	_ = v169
	var v171 int64
	_ = v171
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v457 int64
	_ = v457
	var v460 int32
	_ = v460
	var v461 int64
	_ = v461
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v501 int32
	_ = v501
	var v507 int32
	_ = v507
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v665 int32
	_ = v665
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v703 int32
	_ = v703
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v738 int32
	_ = v738
	var v742 int32
	_ = v742
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v751 int64
	_ = v751
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
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
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v847 int32
	_ = v847
	var v863 int32
	_ = v863
	var v866 int32
	_ = v866
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v920 int32
	_ = v920
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v927 int32
	_ = v927
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v938 int32
	_ = v938
	var v943 int32
	_ = v943
	var v946 int32
	_ = v946
	var v950 int32
	_ = v950
	var v954 int32
	_ = v954
	var v961 int32
	_ = v961
	var v964 int32
	_ = v964
	var v968 int32
	_ = v968
	var v973 int32
	_ = v973
	var v995 int32
	_ = v995
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1008 int32
	_ = v1008
	var v1011 int32
	_ = v1011
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1042 int32
	_ = v1042
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1048 int32
	_ = v1048
	var v1051 int32
	_ = v1051
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1065 int32
	_ = v1065
	var v1067 int32
	_ = v1067
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1074 int32
	_ = v1074
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1092 int32
	_ = v1092
	var v1101 int32
	_ = v1101
	var v1118 int32
	_ = v1118
	var v1126 int32
	_ = v1126
	var v1129 int32
	_ = v1129
	var v1133 int32
	_ = v1133
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1145 int32
	_ = v1145
	var v1147 int32
	_ = v1147
	var v1152 int32
	_ = v1152
	var v1155 int32
	_ = v1155
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1163 int32
	_ = v1163
	var v1165 int32
	_ = v1165
	var v1169 int32
	_ = v1169
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1178 int32
	_ = v1178
	var v1195 int32
	_ = v1195
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1219 int32
	_ = v1219
	var v1225 int32
	_ = v1225
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1235 int32
	_ = v1235
	var v1251 int32
	_ = v1251
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1284 int32
	_ = v1284
	var v1287 int32
	_ = v1287
	var v1293 int32
	_ = v1293
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1300 int32
	_ = v1300
	var v1302 int32
	_ = v1302
	var v1305 int32
	_ = v1305
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1311 int32
	_ = v1311
	var v1320 int32
	_ = v1320
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1329 int32
	_ = v1329
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1339 int32
	_ = v1339
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1348 int32
	_ = v1348
	var v1350 int32
	_ = v1350
	var v1352 int32
	_ = v1352
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1359 int32
	_ = v1359
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1368 int32
	_ = v1368
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1403 int32
	_ = v1403
	var v1406 int32
	_ = v1406
	var v1410 int32
	_ = v1410
	var v1415 int32
	_ = v1415
	var v1419 int32
	_ = v1419
	var v1422 int32
	_ = v1422
	var v1426 int32
	_ = v1426
	var v1431 int32
	_ = v1431
	var v1435 int32
	_ = v1435
	var v1441 int32
	_ = v1441
	var v1446 int32
	_ = v1446
	var v1450 int32
	_ = v1450
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1457 int32
	_ = v1457
	var v1462 int32
	_ = v1462
	var v1471 int32
	_ = v1471
	var v1475 int32
	_ = v1475
	var v1480 int32
	_ = v1480
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
	v58 = m.ExcPending
	if v58 != 0 {
		goto L7
	} else {
		goto L8
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
	v53 = F___sigaction(m, int32(15), v29+int32(12), int32(0))
	mBase = m.M
	m.G0 = v29 + int32(32)
	goto L1
L7:
	;
	return
L8:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[2]))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+1328))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[3])) = v62
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[4]))
	v71 = F_AllocSetContextCreateInternal(m, v66, int32(_a_F_ParallelWorkerMain_1), int32(0), int32(_a_F_ParallelWorkerMain_2), int32(_a_F_ParallelWorkerMain_3))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[5])) = v71
	v74 = F_dsm_attach(m, l0)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L7
	} else {
		goto L14
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1471 = m.ExcPending
	if v1471 != 0 {
		goto L7
	} else {
		goto L356
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1450 = m.ExcPending
	if v1450 != 0 {
		goto L7
	} else {
		goto L352
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1435 = m.ExcPending
	if v1435 != 0 {
		goto L7
	} else {
		goto L349
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1419 = m.ExcPending
	if v1419 != 0 {
		goto L7
	} else {
		goto L345
	}
L14:
	;
	if v74 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v74)+24))
	v79 = *(*int64)(unsafe.Add(mBase, uint32(v77)))
	if v79 == int64(1346862204) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1403 = m.ExcPending
	if v1403 != 0 {
		goto L7
	} else {
		goto L341
	}
L18:
	;
	if v81 == int32(0) {
		goto L13
	} else {
		goto L22
	}
L19:
	;
	v81 = v77
	goto L21
L20:
	;
	v81 = int32(0)
	goto L21
L21:
	;
	goto L18
L22:
	;
	v87 = F_shm_toc_lookup(m, v81, int64(-65535), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[6])) = v87
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v87)+40))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[7])) = v91
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v87)+44))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[8])) = v94
	F_before_shmem_exit(m, int32(296), v74)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L7
	} else {
		goto L24
	}
L24:
	;
	v101 = F_shm_toc_lookup(m, v81, int64(-65534), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[3]))
	v107 = v101 + v104<<(uint(int32(14))%32)
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[9]))
	F_shm_mq_set_sender(m, v107, v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L7
	} else {
		goto L26
	}
L26:
	;
	v112 = F_shm_mq_attach(m, v107, v74)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L7
	} else {
		goto L27
	}
L27:
	;
	F_pq_redirect_to_shm_mq(m, v74, v112)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L7
	} else {
		goto L28
	}
L28:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v87)+40))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v87)+44))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[10])) = v117
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[11])) = v116
	goto L29
L29:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v87)+40))
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[12]))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v87)+36))
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[13]))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	v131 = base.I32_div_s(v125-v128, int32(640))
	v133 = base.I32_rem_s(v131, int32(16))
	v138 = v124 + v133<<(uint(int32(7))%32) + int32(_a_F_ParallelWorkerMain_4)
	v140 = F_LWLockAcquire(m, v138, int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L7
	} else {
		goto L30
	}
L30:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v125)+44))
	if v142 != v122 {
		v165 = v2
		goto L31
	} else {
		goto L32
	}
L31:
	;
	F_LWLockRelease(m, v138)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L7
	} else {
		goto L37
	}
L32:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v125)+616))
	if v144 != v125 {
		v165 = v2
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v147)+616)) = v125
	v150 = v125 + int32(620)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v125)+624))
	if v151 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v125)+624)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v125)+620)) = v150
	goto L36
L35:
	;
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147)+632)) = v150
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	*(*int32)(unsafe.Add(mBase, uint32(v147)+628)) = v157
	v160 = v147 + int32(628)
	*(*int32)(unsafe.Add(mBase, uint32(v157)+4)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v150))) = v160
	v165 = int32(1)
	goto L31
L37:
	;
	if v165 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v169 = *(*int64)(unsafe.Add(mBase, uint32(v87)+48))
	v171 = *(*int64)(unsafe.Add(mBase, uint32(v87)+56))
	*(*int64)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[14])) = v171
	*(*int64)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[15])) = v169
	v177 = F_shm_toc_lookup(m, v81, int64(-65527), int32(0))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L7
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	m.G0 = v19 + int32(32)
	return
L41:
	;
	v179 = F_strlen(m, v177)
	mBase = m.M
	v182 = v179 + v177 + int32(1)
	v183 = int32(_a_F_ParallelWorkerMain_5)
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
	v189 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[16])))
	if base.B2i32(v186 == int32(0))|base.B2i32(v186 != v189) != 0 {
		v207 = v186
		v208 = v189
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[17])) = v371
	v374 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v374)+64)) = v371
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+32)))
	F_SetSessionAuthorization(m, v376, v377)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L7
	} else {
		goto L102
	}
L43:
	;
	if v207-v208 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L44:
	;
	goto L43
L45:
	;
	v192 = v177
	v193 = v183
	goto L46
L46:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+1)))
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+1)))
	if v197 == int32(0) {
		v207 = v197
		v208 = v196
		goto L44
	} else {
		goto L48
	}
L47:
	;
	v207 = v197
	v208 = v196
	goto L44
L48:
	;
	v200 = int32(1)
	if v197 == v196 {
		v192 = v192 + v200
		v193 = v193 + v200
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v212 = int32(_a_F_ParallelWorkerMain_6)
	v215 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[18])))
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	if base.B2i32(v215 == int32(0))|base.B2i32(v215 != v218) != 0 {
		v236 = v215
		v237 = v218
		goto L54
	} else {
		goto L55
	}
L51:
	;
	goto L52
L52:
	;
	v367 = F_load_external_function(m, v177, v182, int32(1), int32(0))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L7
	} else {
		goto L101
	}
L53:
	;
	if v236-v237 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L54:
	;
	goto L53
L55:
	;
	v221 = v212
	v222 = v182
	goto L56
L56:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+1)))
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+1)))
	if v226 == int32(0) {
		v236 = v226
		v237 = v225
		goto L54
	} else {
		goto L58
	}
L57:
	;
	v236 = v226
	v237 = v225
	goto L54
L58:
	;
	v229 = int32(1)
	if v226 == v225 {
		v221 = v221 + v229
		v222 = v222 + v229
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	v242 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[19]))
	v369 = v242
	goto L42
L61:
	;
	goto L62
L62:
	;
	v243 = int32(_a_F_ParallelWorkerMain_7)
	v246 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[20])))
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	if base.B2i32(v246 == int32(0))|base.B2i32(v246 != v249) != 0 {
		v267 = v246
		v268 = v249
		goto L64
	} else {
		goto L65
	}
L63:
	;
	if v267-v268 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L64:
	;
	goto L63
L65:
	;
	v252 = v243
	v253 = v182
	goto L66
L66:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253)+1)))
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252)+1)))
	if v257 == int32(0) {
		v267 = v257
		v268 = v256
		goto L64
	} else {
		goto L68
	}
L67:
	;
	v267 = v257
	v268 = v256
	goto L64
L68:
	;
	v260 = int32(1)
	if v257 == v256 {
		v252 = v252 + v260
		v253 = v253 + v260
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	v273 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[21]))
	v369 = v273
	goto L42
L71:
	;
	goto L72
L72:
	;
	v274 = int32(_a_F_ParallelWorkerMain_8)
	v277 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[22])))
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	if base.B2i32(v277 == int32(0))|base.B2i32(v277 != v280) != 0 {
		v298 = v277
		v299 = v280
		goto L74
	} else {
		goto L75
	}
L73:
	;
	if v298-v299 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L74:
	;
	goto L73
L75:
	;
	v283 = v274
	v284 = v182
	goto L76
L76:
	;
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284)+1)))
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+1)))
	if v288 == int32(0) {
		v298 = v288
		v299 = v287
		goto L74
	} else {
		goto L78
	}
L77:
	;
	v298 = v288
	v299 = v287
	goto L74
L78:
	;
	v291 = int32(1)
	if v288 == v287 {
		v283 = v283 + v291
		v284 = v284 + v291
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	v304 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[23]))
	v369 = v304
	goto L42
L81:
	;
	goto L82
L82:
	;
	v305 = int32(_a_F_ParallelWorkerMain_9)
	v308 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[24])))
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	if base.B2i32(v308 == int32(0))|base.B2i32(v308 != v311) != 0 {
		v329 = v308
		v330 = v311
		goto L84
	} else {
		goto L85
	}
L83:
	;
	if v329-v330 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L84:
	;
	goto L83
L85:
	;
	v314 = v305
	v315 = v182
	goto L86
L86:
	;
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+1)))
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314)+1)))
	if v319 == int32(0) {
		v329 = v319
		v330 = v318
		goto L84
	} else {
		goto L88
	}
L87:
	;
	v329 = v319
	v330 = v318
	goto L84
L88:
	;
	v322 = int32(1)
	if v319 == v318 {
		v314 = v314 + v322
		v315 = v315 + v322
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	v335 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[25]))
	v369 = v335
	goto L42
L91:
	;
	goto L92
L92:
	;
	v336 = int32(_a_F_ParallelWorkerMain_10)
	v339 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[26])))
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	if base.B2i32(v339 == int32(0))|base.B2i32(v339 != v342) != 0 {
		v360 = v339
		v361 = v342
		goto L94
	} else {
		goto L95
	}
L93:
	;
	if v360-v361 != 0 {
		goto L12
	} else {
		goto L100
	}
L94:
	;
	goto L93
L95:
	;
	v345 = v336
	v346 = v182
	goto L96
L96:
	;
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+1)))
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345)+1)))
	if v350 == int32(0) {
		v360 = v350
		v361 = v349
		goto L94
	} else {
		goto L98
	}
L97:
	;
	v360 = v350
	v361 = v349
	goto L94
L98:
	;
	v353 = int32(1)
	if v350 == v349 {
		v345 = v345 + v353
		v346 = v346 + v353
		goto L96
	} else {
		goto L99
	}
L99:
	;
	goto L97
L100:
	;
	v364 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[27]))
	v369 = v364
	goto L42
L101:
	;
	v369 = v367
	goto L42
L102:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+33)))
	F_SetCurrentRoleId(m, v380, v381)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L7
	} else {
		goto L103
	}
L103:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	F_BackgroundWorkerInitializeConnectionByOid(m, v384, v385, int32(3))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L7
	} else {
		goto L104
	}
L104:
	;
	v390 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[28]))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v390)+4))
	goto L105
L105:
	;
	v392 = F_SetClientEncoding(m, v391)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L7
	} else {
		goto L106
	}
L106:
	;
	if v392 < int32(0) {
		goto L11
	} else {
		goto L107
	}
L107:
	;
	v398 = F_shm_toc_lookup(m, v81, int64(-65533), int32(0))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L7
	} else {
		goto L108
	}
L108:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L7
	} else {
		goto L109
	}
L109:
	;
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v398))))
	if v402 != 0 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v403 = v398
	goto L113
L111:
	;
	goto L112
L112:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L7
	} else {
		goto L117
	}
L113:
	;
	v419 = F_internal_load_library(m, v403)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L7
	} else {
		goto L115
	}
L114:
	;
	goto L112
L115:
	;
	v421 = F_strlen(m, v403)
	mBase = m.M
	v424 = v421 + v403 + int32(1)
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v424))))
	if v425 != 0 {
		v403 = v424
		goto L113
	} else {
		goto L116
	}
L116:
	;
	goto L114
L117:
	;
	v446 = F_shm_toc_lookup(m, v81, int64(-65528), int32(0))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L7
	} else {
		goto L118
	}
L118:
	;
	F_StartTransaction(m)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L7
	} else {
		goto L119
	}
L119:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v446)))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[29])) = v451
	v454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v446)+4)))
	*(*uint8)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[30])) = uint8(v454)
	v457 = *(*int64)(unsafe.Add(mBase, uint32(v446)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[31])) = v457
	v460 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[32]))
	v461 = *(*int64)(unsafe.Add(mBase, uint32(v446)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v460))) = v461
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v446)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[33])) = v464
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v446)+28))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[34])) = v446 + int32(32)
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[35])) = v466
	*(*int32)(unsafe.Add(mBase, uint32(v460)+24)) = int32(5)
	v477 = F_shm_toc_lookup(m, v81, int64(-65525), int32(0))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L7
	} else {
		goto L120
	}
L120:
	;
	v479 = m.G0
	v481 = v479 - int32(48)
	m.G0 = v481
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v477)+8))
	if v483 != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v484 = v477
	goto L124
L122:
	;
	goto L123
L123:
	;
	m.G0 = v481 + int32(48)
	v546 = F_shm_toc_lookup(m, v81, int64(-65523), int32(0))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L7
	} else {
		goto L132
	}
L124:
	;
	v501 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[36]))
	if v501 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	goto L123
L126:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v481)+16)) = int64(68719476748)
	v507 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[37]))
	*(*int32)(unsafe.Add(mBase, uint32(v481)+40)) = v507
	v513 = F_hash_create(m, int32(_a_F_ParallelWorkerMain_11), int32(16), v481, int32(1064))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L7
	} else {
		goto L129
	}
L127:
	;
	v516 = v501
	goto L128
L128:
	;
	v518 = F_hash_search(m, v516, v484, int32(1), v481)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L7
	} else {
		goto L130
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[36])) = v513
	v516 = v513
	goto L128
L130:
	;
	v520 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v518)+12)) = uint8(v520)
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v484)+20))
	if v522 != 0 {
		v484 = v484 + int32(12)
		goto L124
	} else {
		goto L131
	}
L131:
	;
	goto L125
L132:
	;
	v549 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[38]))
	if v549 != 0 {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	v572 = int32(524)
	base.MemoryCopy(m, int32(_a_F_ParallelWorkerMain_12), v546, v572)
	base.MemoryCopy(m, int32(_a_F_ParallelWorkerMain_13), v546+v572, v572)
	v581 = F_shm_toc_lookup(m, v81, int64(-65524), int32(0))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L7
	} else {
		goto L142
	}
L134:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L7
	} else {
		goto L139
	}
L135:
	;
	v551 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[39]))
	if v551 != 0 {
		goto L134
	} else {
		goto L136
	}
L136:
	;
	v553 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[40]))
	if v553 != 0 {
		goto L134
	} else {
		goto L137
	}
L137:
	;
	v555 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[41]))
	if v555 == int32(0) {
		goto L133
	} else {
		goto L138
	}
L138:
	;
	goto L134
L139:
	;
	F_errmsg_internal(m, int32(_a_F_ParallelWorkerMain_14), int32(0))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L7
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_ParallelWorkerMain_15), int32(749), int32(_a_F_ParallelWorkerMain_16))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L7
	} else {
		goto L141
	}
L141:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L142:
	;
	v583 = int32(0)
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v581)))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[42])) = v585
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v581)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[43])) = v588
	v590 = int32(_a_F_ParallelWorkerMain_17)
	v591 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[5]))
	v594 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[5])) = v594
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v581)+8))
	if v583 < v596 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v602 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[44]))
	v604 = v583
	v605 = v602
	goto L146
L144:
	;
	goto L145
L145:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[5])) = v591
	v651 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[32]))
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v651)+28))
	goto L150
L146:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v581+int32(12)+v604<<(uint(int32(2))%32))))
	v624 = F_lappend_oid(m, v605, v623)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L7
	} else {
		goto L148
	}
L147:
	;
	goto L145
L148:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[44])) = v624
	v628 = v604 + int32(1)
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v581)+8))
	if v628 < v629 {
		v604 = v628
		v605 = v624
		goto L146
	} else {
		goto L149
	}
L149:
	;
	goto L147
L150:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[45])) = v652
	v656 = F_shm_toc_lookup(m, v81, int64(-65531), int32(0))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L7
	} else {
		goto L151
	}
L151:
	;
	v658 = int32(0)
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v656)))
	if v659 <= v658 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v722 = F_shm_toc_lookup(m, v81, int64(-65526), int32(0))
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L7
	} else {
		goto L164
	}
L153:
	;
	v665 = v658
	goto L154
L154:
	;
	v682 = v656 + int32(4) + v665<<(uint(int32(3))%32)
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v682)))
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v682)+4))
	v685 = F_GetComboCommandId(m, v683, v684)
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L7
	} else {
		goto L156
	}
L155:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L7
	} else {
		goto L161
	}
L156:
	;
	if v685 == v665 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v689 = v665 + int32(1)
	if v659 != v689 {
		v665 = v689
		goto L154
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	goto L155
L160:
	;
	goto L152
L161:
	;
	F_errmsg_internal(m, int32(_a_F_ParallelWorkerMain_18), int32(0))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L7
	} else {
		goto L162
	}
L162:
	;
	F_errfinish(m, int32(_a_F_ParallelWorkerMain_19), int32(362), int32(_a_F_ParallelWorkerMain_20))
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L7
	} else {
		goto L163
	}
L163:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L164:
	;
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v722)))
	v725 = int32(_a_F_ParallelWorkerMain_17)
	v726 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[5]))
	v729 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[5])) = v729
	v731 = F_dsm_attach(m, v724)
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L7
	} else {
		goto L165
	}
L165:
	;
	if v731 == int32(0) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L7
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v731)+24))
	v751 = *(*int64)(unsafe.Add(mBase, uint32(v749)))
	if v751 == int64(2880502729) {
		goto L173
	} else {
		goto L174
	}
L169:
	;
	F_errmsg_internal(m, int32(_a_F_ParallelWorkerMain_21), int32(0))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L7
	} else {
		goto L170
	}
L170:
	;
	F_errfinish(m, int32(_a_F_ParallelWorkerMain_22), int32(169), int32(_a_F_ParallelWorkerMain_23))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L7
	} else {
		goto L171
	}
L171:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L172:
	;
	v756 = F_shm_toc_lookup(m, v753, int64(-65535), int32(0))
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L7
	} else {
		goto L176
	}
L173:
	;
	v753 = v749
	goto L175
L174:
	;
	v753 = int32(0)
	goto L175
L175:
	;
	goto L172
L176:
	;
	v758 = F_dsa_attach_in_place(m, v756, v731)
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L7
	} else {
		goto L177
	}
L177:
	;
	v760 = int32(_a_F_ParallelWorkerMain_24)
	v761 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[46]))
	*(*int32)(unsafe.Add(mBase, uint32(v761))) = v731
	v764 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[46]))
	*(*int32)(unsafe.Add(mBase, uint32(v764)+4)) = v758
	v768 = F_shm_toc_lookup(m, v753, int64(-65534), int32(0))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L7
	} else {
		goto L178
	}
L178:
	;
	v770 = int32(_a_F_ParallelWorkerMain_17)
	v771 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[5]))
	v774 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[5])) = v774
	v777 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[46]))
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v777)+4))
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v768)))
	v781 = F_dshash_attach(m, v778, int32(_a_F_ParallelWorkerMain_25), v780, v778)
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L7
	} else {
		goto L179
	}
L179:
	;
	v784 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[46]))
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v784)+4))
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v768)+4))
	v789 = F_dshash_attach(m, v785, int32(_a_F_ParallelWorkerMain_26), v787, int32(0))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L7
	} else {
		goto L180
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[5])) = v771
	v794 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[46]))
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v794)))
	F_on_dsm_detach(m, v795, int32(1607), v768)
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L7
	} else {
		goto L181
	}
L181:
	;
	v800 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[46]))
	*(*int32)(unsafe.Add(mBase, uint32(v800)+16)) = v789
	*(*int32)(unsafe.Add(mBase, uint32(v800)+12)) = v781
	*(*int32)(unsafe.Add(mBase, uint32(v800)+8)) = v768
	F_dsm_pin_mapping(m, v731)
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L7
	} else {
		goto L182
	}
L182:
	;
	F_dsa_pin_mapping(m, v758)
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L7
	} else {
		goto L183
	}
L183:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[5])) = v726
	v812 = F_shm_toc_lookup(m, v81, int64(-65529), int32(0))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L7
	} else {
		goto L184
	}
L184:
	;
	v816 = F_shm_toc_lookup(m, v81, int64(-65530), int32(1))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L7
	} else {
		goto L185
	}
L185:
	;
	v818 = F_RestoreSnapshot(m, v812)
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L7
	} else {
		goto L186
	}
L186:
	;
	if v816 != 0 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v820 = F_RestoreSnapshot(m, v816)
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L7
	} else {
		goto L190
	}
L188:
	;
	v822 = v818
	goto L189
L189:
	;
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v87)+36))
	F_RestoreTransactionSnapshot(m, v822, v823)
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L7
	} else {
		goto L191
	}
L190:
	;
	v822 = v820
	goto L189
L191:
	;
	F_PushActiveSnapshot(m, v818)
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L7
	} else {
		goto L192
	}
L192:
	;
	F_InvalidateSystemCachesExtended(m)
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L7
	} else {
		goto L193
	}
L193:
	;
	v832 = F_shm_toc_lookup(m, v81, int64(-65532), int32(0))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L7
	} else {
		goto L194
	}
L194:
	;
	v834 = m.G0
	v836 = v834 - int32(32)
	m.G0 = v836
	v839 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[47]))
	v840 = int32(0)
	if base.B2i32(v839 == v840)|base.B2i32(v839 == int32(_a_F_ParallelWorkerMain_27)) == v840 {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v847 = v839
	goto L198
L196:
	;
	goto L197
L197:
	;
	v995 = *(*int32)(unsafe.Add(mBase, uint32(v832)))
	*(*int32)(unsafe.Add(mBase, uint32(v836)+20)) = int32(1643)
	v998 = int32(_a_F_ParallelWorkerMain_28)
	v999 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[48]))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[48])) = v836 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v836)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v836)+16)) = v999
	v1008 = v832 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v836)+28)) = v1008
	if v995 != 0 {
		goto L266
	} else {
		goto L267
	}
L198:
	;
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v847)+4))
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v847-int32(60))))
	if base.Ui32(v866) < base.Ui32(int32(2)) {
		goto L200
	} else {
		goto L201
	}
L199:
	;
	goto L197
L200:
	;
	if v863 != int32(_a_F_ParallelWorkerMain_27) {
		v847 = v863
		goto L198
	} else {
		goto L263
	}
L201:
	;
	v870 = v847 - int32(32)
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v870)))
	if v871 == int32(0) {
		goto L200
	} else {
		goto L202
	}
L202:
	;
	v875 = v847 - int32(4)
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v875)))
	if v876 != 0 {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	F_pfree(m, v876)
	mBase = m.M
	v878 = m.ExcPending
	if v878 != 0 {
		goto L7
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v847)+16))
	if v879 != 0 {
		goto L207
	} else {
		goto L208
	}
L206:
	;
	goto L205
L207:
	;
	F_pfree(m, v879)
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L7
	} else {
		goto L210
	}
L208:
	;
	goto L209
L209:
	;
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v847)+20))
	if v882 != 0 {
		goto L211
	} else {
		goto L212
	}
L210:
	;
	goto L209
L211:
	;
	F_pfree(m, v882)
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L7
	} else {
		goto L214
	}
L212:
	;
	goto L213
L213:
	;
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v847-int32(40))))
	switch v887 {
	case 0:
		goto L221
	case 1:
		goto L220
	case 2:
		goto L219
	case 3:
		goto L218
	case 4:
		goto L217
	default:
		goto L215
	}
L214:
	;
	goto L213
L215:
	;
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v870)))
	if v929 != 0 {
		goto L241
	} else {
		goto L242
	}
L216:
	;
	F_pfree(m, v925)
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L7
	} else {
		goto L240
	}
L217:
	;
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v847)+56))
	if v920 == int32(0) {
		goto L215
	} else {
		goto L238
	}
L218:
	;
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v847)+28))
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v903)))
	if v904 != 0 {
		goto L228
	} else {
		goto L229
	}
L219:
	;
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v847)+80))
	if v898 == int32(0) {
		goto L215
	} else {
		goto L226
	}
L220:
	;
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v847)+60))
	if v893 == int32(0) {
		goto L215
	} else {
		goto L224
	}
L221:
	;
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v847)+52))
	if v888 == int32(0) {
		goto L215
	} else {
		goto L222
	}
L222:
	;
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v875)))
	if v888 != v891 {
		v925 = v888
		goto L216
	} else {
		goto L223
	}
L223:
	;
	goto L215
L224:
	;
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v875)))
	if v893 != v896 {
		v925 = v893
		goto L216
	} else {
		goto L225
	}
L225:
	;
	goto L215
L226:
	;
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v875)))
	if v898 != v901 {
		v925 = v898
		goto L216
	} else {
		goto L227
	}
L227:
	;
	goto L215
L228:
	;
	F_pfree(m, v904)
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L7
	} else {
		goto L231
	}
L229:
	;
	goto L230
L230:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v847)+48))
	if v907 == int32(0) {
		goto L232
	} else {
		goto L233
	}
L231:
	;
	goto L230
L232:
	;
	v915 = *(*int32)(unsafe.Add(mBase, uint32(v847)+52))
	if v915 == int32(0) {
		goto L215
	} else {
		goto L236
	}
L233:
	;
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v847)+28))
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v910)))
	if v907 == v911 {
		goto L232
	} else {
		goto L234
	}
L234:
	;
	F_pfree(m, v907)
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L7
	} else {
		goto L235
	}
L235:
	;
	goto L232
L236:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v875)))
	if v915 != v918 {
		v925 = v915
		goto L216
	} else {
		goto L237
	}
L237:
	;
	goto L215
L238:
	;
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v875)))
	if v920 == v923 {
		goto L215
	} else {
		goto L239
	}
L239:
	;
	v925 = v920
	goto L216
L240:
	;
	goto L215
L241:
	;
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v847)))
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v847)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v930)+4)) = v931
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v847)))
	*(*int32)(unsafe.Add(mBase, uint32(v931))) = v933
	goto L243
L242:
	;
	goto L243
L243:
	;
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v847-int32(8))))
	if v938 != 0 {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v943 = int32(_a_F_ParallelWorkerMain_29)
	goto L249
L245:
	;
	goto L246
L246:
	;
	v954 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v847-int32(36)))))
	if v954&int32(4) != 0 {
		goto L253
	} else {
		goto L254
	}
L247:
	;
	goto L246
L248:
	;
	goto L247
L249:
	;
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v943)))
	if v946 == int32(0) {
		goto L248
	} else {
		goto L251
	}
L250:
	;
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v946)))
	*(*int32)(unsafe.Add(mBase, uint32(v943))) = v950
	goto L248
L251:
	;
	if v946 != v847+int32(8) {
		v943 = v946
		goto L249
	} else {
		goto L252
	}
L252:
	;
	goto L250
L253:
	;
	v961 = int32(_a_F_ParallelWorkerMain_30)
	goto L258
L254:
	;
	goto L255
L255:
	;
	F_InitializeOneGUCOption(m, v847+int32(-64))
	mBase = m.M
	v973 = m.ExcPending
	if v973 != 0 {
		goto L7
	} else {
		goto L262
	}
L256:
	;
	goto L255
L257:
	;
	goto L256
L258:
	;
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v961)))
	if v964 == int32(0) {
		goto L257
	} else {
		goto L260
	}
L259:
	;
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v964)))
	*(*int32)(unsafe.Add(mBase, uint32(v961))) = v968
	goto L257
L260:
	;
	if v964 != v847+int32(12) {
		v961 = v964
		goto L258
	} else {
		goto L261
	}
L261:
	;
	goto L259
L262:
	;
	goto L200
L263:
	;
	goto L199
L264:
	;
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(v87)+16))
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(v87)+28))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[49])) = v1140
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[50])) = v1139
	goto L301
L265:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1126 = m.ExcPending
	if v1126 != 0 {
		goto L7
	} else {
		goto L297
	}
L266:
	;
	v1011 = v995 + v1008
	goto L269
L267:
	;
	v1118 = v999
	goto L268
L268:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[48])) = v1118
	m.G0 = v836 + int32(32)
	goto L264
L269:
	;
	v1029 = v836 + int32(28)
	v1030 = F_read_gucstate(m, v1029, v1011)
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L7
	} else {
		goto L271
	}
L270:
	;
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(v836)+16))
	v1118 = v1101
	goto L268
L271:
	;
	v1032 = F_read_gucstate(m, v1029, v1011)
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L7
	} else {
		goto L272
	}
L272:
	;
	v1034 = F_read_gucstate(m, v1029, v1011)
	mBase = m.M
	v1035 = m.ExcPending
	if v1035 != 0 {
		goto L7
	} else {
		goto L273
	}
L273:
	;
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v836)+28))
	v1037 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1034))))
	if v1037 == int32(0) {
		goto L275
	} else {
		goto L276
	}
L274:
	;
	v1048 = v1045 + int32(4)
	if base.Ui32(v1011) < base.Ui32(v1048) {
		goto L10
	} else {
		goto L279
	}
L275:
	;
	v1045 = v1036
	v1046 = int32(0)
	goto L274
L276:
	;
	goto L277
L277:
	;
	v1042 = v1036 + int32(4)
	if base.Ui32(v1011) < base.Ui32(v1042) {
		goto L10
	} else {
		goto L278
	}
L278:
	;
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(v1036)))
	v1045 = v1042
	v1046 = v1044
	goto L274
L279:
	;
	v1051 = v1045 + int32(8)
	if base.Ui32(v1011) < base.Ui32(v1051) {
		goto L10
	} else {
		goto L280
	}
L280:
	;
	v1054 = v1045 + int32(12)
	if base.Ui32(v1011) < base.Ui32(v1054) {
		goto L10
	} else {
		goto L281
	}
L281:
	;
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v1045)))
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(v1051)))
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v1048)))
	*(*int32)(unsafe.Add(mBase, uint32(v836)+12)) = v1032
	*(*int32)(unsafe.Add(mBase, uint32(v836)+8)) = v1030
	*(*int32)(unsafe.Add(mBase, uint32(v836)+28)) = v1054
	*(*int32)(unsafe.Add(mBase, uint32(v836)+24)) = v836 + int32(8)
	v1065 = int32(0)
	v1067 = int32(1)
	v1070 = F_set_config_with_handle(m, v1030, v1065, v1032, v1058, v1056, v1057, v1065, v1067, int32(21), v1067)
	mBase = m.M
	v1071 = m.ExcPending
	if v1071 != 0 {
		goto L7
	} else {
		goto L282
	}
L282:
	;
	if v1070 <= int32(0) {
		goto L265
	} else {
		goto L283
	}
L283:
	;
	v1074 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1034))))
	if v1074 == int32(0) {
		goto L284
	} else {
		goto L285
	}
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v836)+24)) = int32(0)
	if base.Ui32(v1054) < base.Ui32(v1011) {
		goto L269
	} else {
		goto L296
	}
L285:
	;
	v1082 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[51])))
	if v1082 != 0 {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	v1083 = int32(12)
	goto L288
L287:
	;
	v1083 = int32(15)
	goto L288
L288:
	;
	v1084 = F_find_option(m, v1030, int32(1), int32(0), v1083)
	mBase = m.M
	v1085 = m.ExcPending
	if v1085 != 0 {
		goto L7
	} else {
		goto L289
	}
L289:
	;
	if v1084 == int32(0) {
		goto L284
	} else {
		goto L290
	}
L290:
	;
	v1088 = F_guc_strdup(m, v1083, v1034)
	mBase = m.M
	v1089 = m.ExcPending
	if v1089 != 0 {
		goto L7
	} else {
		goto L291
	}
L291:
	;
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(v1084)+84))
	if v1090 != 0 {
		goto L292
	} else {
		goto L293
	}
L292:
	;
	F_pfree(m, v1090)
	mBase = m.M
	v1092 = m.ExcPending
	if v1092 != 0 {
		goto L7
	} else {
		goto L295
	}
L293:
	;
	goto L294
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1084)+88)) = v1046
	*(*int32)(unsafe.Add(mBase, uint32(v1084)+84)) = v1088
	goto L284
L295:
	;
	goto L294
L296:
	;
	goto L270
L297:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v1129 = m.ExcPending
	if v1129 != 0 {
		goto L7
	} else {
		goto L298
	}
L298:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v836))) = v1030
	F_errmsg(m, int32(_a_F_ParallelWorkerMain_31), v836)
	mBase = m.M
	v1133 = m.ExcPending
	if v1133 != 0 {
		goto L7
	} else {
		goto L299
	}
L299:
	;
	F_errfinish(m, int32(_a_F_ParallelWorkerMain_32), int32(_a_F_ParallelWorkerMain_33), int32(_a_F_ParallelWorkerMain_34))
	mBase = m.M
	v1138 = m.ExcPending
	if v1138 != 0 {
		goto L7
	} else {
		goto L300
	}
L300:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L301:
	;
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v87)+20))
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(v87)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[52])) = v1147
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[53])) = v1145
	v1152 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[54])) = uint8(v1152)
	v1155 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[55])) = uint8(v1155)
	v1159 = F_shm_toc_lookup(m, v81, int64(-65522), v1155)
	mBase = m.M
	v1160 = m.ExcPending
	if v1160 != 0 {
		goto L7
	} else {
		goto L302
	}
L302:
	;
	v1161 = m.G0
	v1163 = v1161 - int32(48)
	m.G0 = v1163
	v1165 = *(*int32)(unsafe.Add(mBase, uint32(v1159)))
	if v1165 != 0 {
		goto L303
	} else {
		goto L304
	}
L303:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1163)+16)) = int64(17179869188)
	v1169 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[37]))
	*(*int32)(unsafe.Add(mBase, uint32(v1163)+40)) = v1169
	v1175 = F_hash_create(m, int32(_a_F_ParallelWorkerMain_35), int32(32), v1163, int32(1064))
	mBase = m.M
	v1176 = m.ExcPending
	if v1176 != 0 {
		goto L7
	} else {
		goto L306
	}
L304:
	;
	v1203 = v1159
	goto L305
L305:
	;
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(v1203)+4))
	if v1219 != 0 {
		goto L311
	} else {
		goto L312
	}
L306:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[56])) = v1175
	v1178 = v1159
	goto L307
L307:
	;
	v1195 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[56]))
	v1198 = F_hash_search(m, v1195, v1178, int32(1), int32(0))
	mBase = m.M
	v1199 = m.ExcPending
	if v1199 != 0 {
		goto L7
	} else {
		goto L309
	}
L308:
	;
	v1203 = v1201
	goto L305
L309:
	;
	v1201 = v1178 + int32(4)
	v1202 = *(*int32)(unsafe.Add(mBase, uint32(v1178)+4))
	if v1202 != 0 {
		v1178 = v1201
		goto L307
	} else {
		goto L310
	}
L310:
	;
	goto L308
L311:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1163)+16)) = int64(17179869188)
	v1225 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[37]))
	*(*int32)(unsafe.Add(mBase, uint32(v1163)+40)) = v1225
	v1231 = F_hash_create(m, int32(_a_F_ParallelWorkerMain_36), int32(32), v1163, int32(1064))
	mBase = m.M
	v1232 = m.ExcPending
	if v1232 != 0 {
		goto L7
	} else {
		goto L314
	}
L312:
	;
	goto L313
L313:
	;
	m.G0 = v1163 + int32(48)
	v1280 = F_shm_toc_lookup(m, v81, int64(-65521), int32(0))
	mBase = m.M
	v1281 = m.ExcPending
	if v1281 != 0 {
		goto L7
	} else {
		goto L319
	}
L314:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[57])) = v1231
	v1235 = v1203 + int32(4)
	goto L315
L315:
	;
	v1251 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[57]))
	v1254 = F_hash_search(m, v1251, v1235, int32(1), int32(0))
	mBase = m.M
	v1255 = m.ExcPending
	if v1255 != 0 {
		goto L7
	} else {
		goto L317
	}
L316:
	;
	goto L313
L317:
	;
	v1256 = *(*int32)(unsafe.Add(mBase, uint32(v1235)+4))
	if v1256 != 0 {
		v1235 = v1235 + int32(4)
		goto L315
	} else {
		goto L318
	}
L318:
	;
	goto L316
L319:
	;
	v1282 = *(*int32)(unsafe.Add(mBase, uint32(v1280)))
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(v1280)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[58])) = v1284
	v1287 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[59])) = v1287
	if v1287 <= v1282 {
		goto L320
	} else {
		goto L321
	}
L320:
	;
	v1293 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[4]))
	v1296 = F_MemoryContextStrdup(m, v1293, v1280+int32(8))
	mBase = m.M
	v1297 = m.ExcPending
	if v1297 != 0 {
		goto L7
	} else {
		goto L323
	}
L321:
	;
	goto L322
L322:
	;
	v1300 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[59]))
	if v1300 != 0 {
		goto L324
	} else {
		goto L325
	}
L323:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[59])) = v1296
	goto L322
L324:
	;
	v1302 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[58]))
	v1305 = *(*int32)(unsafe.Add(mBase, uint32(v1302<<(uint(int32(2))%32))+uint32(_c_F_ParallelWorkerMain[60])))
	goto L327
L325:
	;
	goto L326
L326:
	;
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(v87)+64))
	v1309 = m.G0
	v1311 = v1309 - int32(48)
	m.G0 = v1311
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[61])) = v1308
	if v1308 != 0 {
		goto L329
	} else {
		goto L330
	}
L327:
	;
	F_InitializeSystemUser(m, v1300, v1305)
	mBase = m.M
	v1307 = m.ExcPending
	if v1307 != 0 {
		goto L7
	} else {
		goto L328
	}
L328:
	;
	goto L326
L329:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1311)+16)) = int64(103079215120)
	v1320 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[62]))
	v1322 = F_hash_create(m, int32(_a_F_ParallelWorkerMain_37), v1320, v1311, int32(40))
	mBase = m.M
	v1323 = m.ExcPending
	if v1323 != 0 {
		goto L7
	} else {
		goto L332
	}
L330:
	;
	goto L331
L331:
	;
	m.G0 = v1311 + int32(48)
	v1329 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[0])) = uint8(v1329)
	v1333 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[32]))
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(v1333)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v1333)+72)) = v1334 + int32(1)
	goto L333
L332:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[63])) = v1322
	goto L331
L333:
	;
	m.T0[v369].(func(*base.Module, int32, int32))(m, v74, v81)
	mBase = m.M
	v1339 = m.ExcPending
	if v1339 != 0 {
		goto L7
	} else {
		goto L334
	}
L334:
	;
	v1342 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[32]))
	v1343 = *(*int32)(unsafe.Add(mBase, uint32(v1342)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v1342)+72)) = v1343 - int32(1)
	goto L335
L335:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v1348 = m.ExcPending
	if v1348 != 0 {
		goto L7
	} else {
		goto L336
	}
L336:
	;
	F_CommitTransaction(m)
	mBase = m.M
	v1350 = m.ExcPending
	if v1350 != 0 {
		goto L7
	} else {
		goto L337
	}
L337:
	;
	v1352 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[32]))
	*(*int32)(unsafe.Add(mBase, uint32(v1352)+24)) = int32(0)
	v1356 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[46]))
	v1357 = *(*int32)(unsafe.Add(mBase, uint32(v1356)))
	F_dsm_detach(m, v1357)
	mBase = m.M
	v1359 = m.ExcPending
	if v1359 != 0 {
		goto L7
	} else {
		goto L338
	}
L338:
	;
	v1360 = int32(_a_F_ParallelWorkerMain_24)
	v1361 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[46]))
	*(*int32)(unsafe.Add(mBase, uint32(v1361))) = int32(0)
	v1365 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[46]))
	v1366 = *(*int32)(unsafe.Add(mBase, uint32(v1365)+4))
	F_dsa_detach(m, v1366)
	mBase = m.M
	v1368 = m.ExcPending
	if v1368 != 0 {
		goto L7
	} else {
		goto L339
	}
L339:
	;
	v1370 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[46]))
	v1371 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1370)+4)) = v1371
	v1377 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[64]))
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+16))
	v1379 = m.T0[v1378].(func(*base.Module, int32, int32, int32) int32)(m, int32(88), v1371, v1371)
	mBase = m.M
	v1380 = m.ExcPending
	if v1380 != 0 {
		goto L7
	} else {
		goto L340
	}
L340:
	;
	goto L40
L341:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1406 = m.ExcPending
	if v1406 != 0 {
		goto L7
	} else {
		goto L342
	}
L342:
	;
	F_errmsg(m, int32(_a_F_ParallelWorkerMain_38), int32(0))
	mBase = m.M
	v1410 = m.ExcPending
	if v1410 != 0 {
		goto L7
	} else {
		goto L343
	}
L343:
	;
	F_errfinish(m, int32(_a_F_ParallelWorkerMain_39), int32(1355), int32(_a_F_ParallelWorkerMain_40))
	mBase = m.M
	v1415 = m.ExcPending
	if v1415 != 0 {
		goto L7
	} else {
		goto L344
	}
L344:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L345:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1422 = m.ExcPending
	if v1422 != 0 {
		goto L7
	} else {
		goto L346
	}
L346:
	;
	F_errmsg(m, int32(_a_F_ParallelWorkerMain_41), int32(0))
	mBase = m.M
	v1426 = m.ExcPending
	if v1426 != 0 {
		goto L7
	} else {
		goto L347
	}
L347:
	;
	F_errfinish(m, int32(_a_F_ParallelWorkerMain_39), int32(1360), int32(_a_F_ParallelWorkerMain_40))
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L7
	} else {
		goto L348
	}
L348:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L349:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v182
	F_errmsg_internal(m, int32(_a_F_ParallelWorkerMain_42), v19+int32(16))
	mBase = m.M
	v1441 = m.ExcPending
	if v1441 != 0 {
		goto L7
	} else {
		goto L350
	}
L350:
	;
	F_errfinish(m, int32(_a_F_ParallelWorkerMain_39), int32(1666), int32(_a_F_ParallelWorkerMain_43))
	mBase = m.M
	v1446 = m.ExcPending
	if v1446 != 0 {
		goto L7
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
	v1452 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[28]))
	v1453 = *(*int32)(unsafe.Add(mBase, uint32(v1452)+4))
	goto L353
L353:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v1453
	F_errmsg_internal(m, int32(_a_F_ParallelWorkerMain_44), v19)
	mBase = m.M
	v1457 = m.ExcPending
	if v1457 != 0 {
		goto L7
	} else {
		goto L354
	}
L354:
	;
	F_errfinish(m, int32(_a_F_ParallelWorkerMain_39), int32(1452), int32(_a_F_ParallelWorkerMain_40))
	mBase = m.M
	v1462 = m.ExcPending
	if v1462 != 0 {
		goto L7
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
	F_errmsg_internal(m, int32(_a_F_ParallelWorkerMain_45), int32(0))
	mBase = m.M
	v1475 = m.ExcPending
	if v1475 != 0 {
		goto L7
	} else {
		goto L357
	}
L357:
	;
	F_errfinish(m, int32(_a_F_ParallelWorkerMain_32), int32(_a_F_ParallelWorkerMain_46), int32(_a_F_ParallelWorkerMain_47))
	mBase = m.M
	v1480 = m.ExcPending
	if v1480 != 0 {
		goto L7
	} else {
		goto L358
	}
L358:
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
