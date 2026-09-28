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
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
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
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v11 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v50 != 0 {
		goto L13
	} else {
		goto L14
	}
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v14 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = int32(0)
	v20 = v14
	goto L4
L4:
	;
	v24 = v19 << (uint(int32(3)) % 32)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v26 = v24 + v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v27 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L1
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	F_TerminateBackgroundWorker(m, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v41 = v20
	goto L8
L8:
	;
	v43 = v19 + int32(1)
	if v43 < v41 {
		v19 = v43
		v20 = v41
		goto L4
	} else {
		goto L12
	}
L9:
	;
	return
L10:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v31+v24)+4))
	F_shm_mq_detach(m, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v36+v24)+4)) = int32(0)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v41 = v40
	goto L8
L12:
	;
	goto L5
L13:
	;
	F_dsm_detach(m, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L9
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v55 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(0)
	goto L15
L17:
	;
	F_pfree(m, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L9
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v60 = int32(_a_F_DestroyParallelContext_0)
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_DestroyParallelContext[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_DestroyParallelContext[0])) = v62 + int32(1)
	F_WaitForParallelWorkersToExit(m, l0)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L9
	} else {
		goto L21
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
	goto L19
L21:
	;
	v68 = int32(_a_F_DestroyParallelContext_0)
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_DestroyParallelContext[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_DestroyParallelContext[0])) = v70 - int32(1)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v74 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	F_pfree(m, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L9
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_pfree(m, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L9
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = int32(0)
	goto L24
L26:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	F_pfree(m, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L9
	} else {
		goto L27
	}
L27:
	;
	F_pfree(m, l0)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
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
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
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
	v24 = F_palloc0_mul(m, int32(36), v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v24
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v29 = F_dsa_get_address(m, v27, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if int32(0) < v31 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v35 = v9 + int32(168)
	v40 = int32(0)
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
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+28))
	goto L15
L14:
	;
	goto L12
L15:
	;
	v53 = v44 + v40*int32(36)
	v54 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = v54
	*(*uint16)(unsafe.Add(mBase, uint32(v53)+25)) = uint16(v54)
	v58 = int32(1)
	v64 = int32(-64)
	v67 = v29 + (((v46*int32(28)+int32(76))<<(uint(v58)%32)+int32(14))&int32(-16)-v64)*v40
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = v67
	v70 = v67 - v64
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashEnsureBatchAccessors[1]))
	v75 = F_sts_attach(m, v70, v72+v58, v35)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+28)) = v75
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	goto L17
L17:
	;
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashEnsureBatchAccessors[1]))
	v92 = F_sts_attach(m, v70+(v78*int32(28)+int32(83))&int32(-8), v89+int32(1), v35)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+32)) = v92
	v96 = v40 + int32(1)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v96 < v97 {
		v40 = v96
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
	var v339 int32
	_ = v339
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v395 int32
	_ = v395
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v496 int32
	_ = v496
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v593 int32
	_ = v593
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v617 int32
	_ = v617
	var v623 int32
	_ = v623
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v642 int32
	_ = v642
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v683 int32
	_ = v683
	var v688 int32
	_ = v688
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v841 int32
	_ = v841
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v851 int32
	_ = v851
	var v860 int32
	_ = v860
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v944 int32
	_ = v944
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
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
	v964 = F_BarrierArriveAndWait(m, v23, int32(134217751))
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L7
	} else {
		goto L164
	}
L3:
	;
	v798 = F_BarrierArriveAndWait(m, v23, int32(134217749))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L7
	} else {
		goto L135
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
	v339 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v339
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v339
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v338
	if base.Ui32(int32(2)) <= base.Ui32(v338) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v352 = int32(32) - base.I32_clz(v338-int32(1))
	goto L63
L62:
	;
	v352 = v339
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v352
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v355 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v354)+24)) = uint8(v355)
	v357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v359 = v357 + int32(40)
	v361 = F_LWLockAcquire(m, v359, v355)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L7
	} else {
		goto L64
	}
L64:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v357)+24))
	if v363 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v367 = v357 + int32(24)
	v369 = v359
	v373 = v363
	goto L68
L66:
	;
	v563 = v359
	goto L67
L67:
	;
	F_LWLockRelease(m, v563)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L7
	} else {
		goto L100
	}
L68:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v383 = F_dsa_get_address(m, v382, v373)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L7
	} else {
		goto L70
	}
L69:
	;
	v563 = v553
	goto L67
L70:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v383)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v367))) = v385
	F_LWLockRelease(m, v369)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L7
	} else {
		goto L71
	}
L71:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v383)+8))
	if v389 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v395 = int32(0)
	goto L75
L73:
	;
	goto L74
L74:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	F_dsa_free(m, v544, v373)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L7
	} else {
		goto L93
	}
L75:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v410 = v395 + (v383 + int32(16))
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v410)+4))
	v413 = v410 + int32(4)
	v415 = v410 + int32(8)
	v416 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(int32(2)) <= base.Ui32(v416) {
		goto L79
	} else {
		goto L80
	}
L76:
	;
	goto L74
L77:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v507)+20))
	v509 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v507)+20)) = v508 + v509
	v512 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v515 = v512 + v496*int32(36)
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v515)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v515)+8)) = v516 + v509
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v415)))
	v525 = (v520+int32(15))&int32(-8) + v395
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v383)+8))
	if base.Ui32(v525) < base.Ui32(v526) {
		v395 = v525
		goto L75
	} else {
		goto L92
	}
L78:
	;
	v475 = v423 * int32(36)
	v476 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v477 = v475 + v476
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v477)+16))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v415)))
	*(*int32)(unsafe.Add(mBase, uint32(v477)+16)) = v478 + (v479+int32(15))&int32(-8)
	v486 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v486+v475)+28))
	F_sts_puttuple(m, v488, v413, v415)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L7
	} else {
		goto L91
	}
L79:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v423 = (v416 - int32(1)) & base.I32_rotr(v411, v421)
	if v423 != 0 {
		goto L78
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v415)))
	v429 = int32(8)
	v433 = F_ExecParallelHashTupleAlloc(m, l0, v428+v429, v19+v429)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L7
	} else {
		goto L83
	}
L82:
	;
	goto L81
L83:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v413)))
	*(*int32)(unsafe.Add(mBase, uint32(v433)+4)) = v435
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v415)))
	if v437 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	base.MemoryCopy(m, v433+int32(8), v415, v437)
	goto L86
L85:
	;
	goto L86
L86:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v445 = v442 + (v409-int32(1))&v411<<(uint(int32(2))%32)
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v445)))
	*(*int32)(unsafe.Add(mBase, uint32(v433))) = v446
	v448 = int32(0)
	v450 = base.AtomicRmwCmpxchg32(m, v445, v448, v446, v441)
	if v446 == v450 {
		v496 = v448
		goto L77
	} else {
		goto L87
	}
L87:
	;
	v455 = v450
	goto L88
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v433))) = v455
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v445)))
	*(*int32)(unsafe.Add(mBase, uint32(v433))) = v469
	v472 = base.AtomicRmwCmpxchg32(m, v445, int32(0), v469, v441)
	if v469 != v472 {
		v455 = v472
		goto L88
	} else {
		goto L90
	}
L89:
	;
	v496 = v448
	goto L77
L90:
	;
	goto L89
L91:
	;
	v496 = v423
	goto L77
L92:
	;
	goto L76
L93:
	;
	v548 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashIncreaseNumBatches[2]))
	if v548 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L7
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v553 = v551 + int32(40)
	v555 = F_LWLockAcquire(m, v553, int32(0))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L7
	} else {
		goto L98
	}
L97:
	;
	goto L96
L98:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v551)+24))
	if v559 != 0 {
		v367 = v551 + int32(24)
		v369 = v553
		v373 = v559
		goto L68
	} else {
		goto L99
	}
L99:
	;
	goto L69
L100:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v578)+12))
	v580 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v578)+4))
	v582 = F_dsa_get_address(m, v580, v581)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L7
	} else {
		goto L101
	}
L101:
	;
	v585 = F_palloc0_mul(m, int32(4), v579)
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L7
	} else {
		goto L102
	}
L102:
	;
	if int32(2) <= v579 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v593 = int32(1)
	goto L106
L104:
	;
	goto L105
L105:
	;
	F_pfree(m, v585)
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L7
	} else {
		goto L132
	}
L106:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v611)+28))
	goto L108
L107:
	;
	v642 = int32(1)
	goto L111
L108:
	;
	v617 = int32(1)
	v623 = int32(-64)
	v630 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashIncreaseNumBatches[3]))
	v633 = F_sts_attach(m, v582+(((v612*int32(28)+int32(76))<<(uint(v617)%32)+int32(14))&int32(-16)-v623)*v593-v623, v630+v617, v578+int32(168))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L7
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v585+v593<<(uint(int32(2))%32)))) = v633
	v637 = v593 + int32(1)
	if v637 != v579 {
		v593 = v637
		goto L106
	} else {
		goto L110
	}
L110:
	;
	goto L107
L111:
	;
	v658 = v585 + v642<<(uint(int32(2))%32)
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v658)))
	F_sts_begin_parallel_scan(m, v659)
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L7
	} else {
		goto L113
	}
L112:
	;
	goto L105
L113:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v658)))
	v665 = F_sts_parallel_scan_next(m, v662, v19+int32(12))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L7
	} else {
		goto L114
	}
L114:
	;
	if v665 != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v668 = v665
	goto L118
L116:
	;
	goto L117
L117:
	;
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v658)))
	F_sts_end_parallel_scan(m, v752)
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L7
	} else {
		goto L130
	}
L118:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v668)))
	v688 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(int32(2)) <= base.Ui32(v688) {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	goto L117
L120:
	;
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v694 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v698 = (v688 - int32(1)) & base.I32_rotr(v693, v694)
	goto L122
L121:
	;
	v698 = int32(0)
	goto L122
L122:
	;
	v699 = int32(36)
	v700 = v698 * v699
	v701 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v702 = v700 + v701
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v702)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v702)+16)) = v703 + (v683+int32(15))&int32(-8)
	v706 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v707 = v706 + v700
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v707)+8))
	v709 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v707)+8)) = v708 + v709
	v712 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v715 = v712 + v642*v699
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v715)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v715)+20)) = v716 + v709
	v720 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v720+v700)+28))
	F_sts_puttuple(m, v722, v19+int32(12), v668)
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L7
	} else {
		goto L123
	}
L123:
	;
	v728 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashIncreaseNumBatches[2]))
	if v728 != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L7
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v658)))
	v734 = F_sts_parallel_scan_next(m, v731, v19+int32(12))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L7
	} else {
		goto L128
	}
L127:
	;
	goto L126
L128:
	;
	if v734 != 0 {
		v668 = v734
		goto L118
	} else {
		goto L129
	}
L129:
	;
	goto L119
L130:
	;
	v756 = v642 + int32(1)
	if v756 != v579 {
		v642 = v756
		goto L111
	} else {
		goto L131
	}
L131:
	;
	goto L112
L132:
	;
	F_ExecParallelHashMergeCounters(m, l0)
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L7
	} else {
		goto L133
	}
L133:
	;
	v779 = F_BarrierArriveAndWait(m, v23, int32(134217753))
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L7
	} else {
		goto L134
	}
L134:
	;
	goto L3
L135:
	;
	if v798 == int32(0) {
		goto L2
	} else {
		goto L136
	}
L136:
	;
	F_ExecParallelHashEnsureBatchAccessors(m, l0)
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L7
	} else {
		goto L137
	}
L137:
	;
	v804 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v804
	v807 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v808 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v808)))
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v809)))
	v811 = F_dsa_get_address(m, v807, v810)
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L7
	} else {
		goto L138
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v811
	v814 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v814)+16))
	v816 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v816
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v816
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v815
	if base.Ui32(int32(2)) <= base.Ui32(v815) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v829 = int32(32) - base.I32_clz(v815-int32(1))
	goto L141
L140:
	;
	v829 = v816
	goto L141
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v829
	v831 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v832 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v831)+24)) = uint8(v832)
	v834 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v836 = F_dsa_get_address(m, v834, v835)
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L7
	} else {
		goto L142
	}
L142:
	;
	v838 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v838 <= int32(0) {
		v926 = v804
		goto L143
	} else {
		goto L144
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v926
	v941 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	F_dsa_free(m, v941, v942)
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		goto L7
	} else {
		goto L163
	}
L144:
	;
	v841 = int32(0)
	v845 = v841
	v847 = v841
	v851 = v841
	goto L145
L145:
	;
	v860 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v860+v845*int32(36))))
	v865 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v864)+60)))
	if v865 == int32(0) {
		goto L148
	} else {
		goto L149
	}
L146:
	;
	v913 = v906 | base.B2i32(int32(1073741822) < v909)
	if (v913|v872)&int32(1) == int32(0) {
		v926 = v804
		goto L143
	} else {
		goto L159
	}
L147:
	;
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v874 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v874)+28))
	goto L153
L148:
	;
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v864)+48))
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v21)+32))
	if base.Ui32(v868) <= base.Ui32(v869) {
		v872 = v851
		goto L147
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	v872 = int32(1)
	goto L147
L151:
	;
	goto L150
L152:
	;
	v908 = v845 + int32(1)
	v909 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v908 < v909 {
		v845 = v908
		v847 = v906
		v851 = v872
		goto L145
	} else {
		goto L158
	}
L153:
	;
	v888 = base.I32_rem_s(v845, v873)
	v891 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v836+(((v875*int32(28)+int32(76))<<(uint(int32(1))%32)+int32(14))&int32(-16)-int32(-64))*v888)+60)))
	if v891 == int32(0) {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v864)+48))
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v21)+32))
	if base.Ui32(v894) <= base.Ui32(v895) {
		v906 = v847
		goto L152
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v864)+52))
	v898 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v898+v888*int32(36))))
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v902)+56))
	v906 = base.B2i32(v897 == v903) | v847
	goto L152
L157:
	;
	goto L156
L158:
	;
	goto L146
L159:
	;
	if v913&int32(1) != 0 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v923 = int32(3)
	goto L162
L161:
	;
	v923 = int32(2)
	goto L162
L162:
	;
	v926 = v923
	goto L143
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = int32(0)
	goto L2
L164:
	;
	goto L1
}
func F_ExecParallelInitializeDSM(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
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
	var v93 int32
	_ = v93
	var v94 int64
	_ = v94
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
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
	var v119 int64
	_ = v119
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
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
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
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
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
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int64
	_ = v172
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
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
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int64
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
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
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int64
	_ = v255
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int64
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int64
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
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
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int64
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int64
	_ = v378
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int64
	_ = v403
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int64
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v449 int32
	_ = v449
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int64
	_ = v457
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int64
	_ = v488
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int64
	_ = v503
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v523 int64
	_ = v523
	var v527 int32
	_ = v527
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
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
	var v617 int32
	_ = v617
	var v618 int64
	_ = v618
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
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
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int64
	_ = v646
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int64
	_ = v674
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int64
	_ = v702
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int64
	_ = v730
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	if l0 == int32(0) {
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
		switch v25 - int32(403) {
		case 0:
			v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v386)+36)))
			if v387 != int32(1) {
				v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
				mBase = m.M
				v744 = m.ExcPending
				if v744 != 0 {
					return int32(0)
				} else {
					return v743
				}
			} else {
				v390 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v391 = *(*int32)(unsafe.Add(mBase, uint32(v390)+52))
				v392 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
				v393 = F_shm_toc_allocate(m, v391, v392)
				mBase = m.M
				v394 = m.ExcPending
				if v394 != 0 {
					return int32(0)
				} else {
					v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
					if v395 != 0 {
						base.MemoryFill(m, v393, int32(0), v395)
					} else {
					}
					F_LWLockInitialize(m, v393, int32(81))
					mBase = m.M
					v400 = m.ExcPending
					if v400 != 0 {
						return int32(0)
					} else {
						v401 = *(*int32)(unsafe.Add(mBase, uint32(v390)+52))
						v402 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v403 = int64(*(*int32)(unsafe.Add(mBase, uint32(v402)+40)))
						F_shm_toc_insert(m, v401, v403, v393)
						mBase = m.M
						v405 = m.ExcPending
						if v405 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = int32(742)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v393
							v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
							mBase = m.M
							v744 = m.ExcPending
							if v744 != 0 {
								return int32(0)
							} else {
								return v743
							}
						}
					}
				}
			}
		default:
			v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
			mBase = m.M
			v744 = m.ExcPending
			if v744 != 0 {
				return int32(0)
			} else {
				return v743
			}
		case 6:
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+36)))
			if v29 == int32(1) {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v34 = F_ScanRelIsReadOnly(m, l0)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v33)+132))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v32)+52))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
					v41 = F_shm_toc_allocate(m, v39, v40)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
						F_table_parallelscan_initialize(m, v43, v41, v44)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v32)+52))
							v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v49 = int64(*(*int32)(unsafe.Add(mBase, uint32(v48)+40)))
							F_shm_toc_insert(m, v47, v49, v41)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
								if v34 != 0 {
									v59 = int32(1024)
								} else {
									v59 = int32(0)
								}
								v61 = F_table_beginscan_parallel(m, v52, v41, v38<<(uint(int32(7))%32)&int32(2048)|v59)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v61
									v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+132)))
									if v71&int32(16) == int32(0) {
										v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
										mBase = m.M
										v744 = m.ExcPending
										if v744 != 0 {
											return int32(0)
										} else {
											return v743
										}
									} else {
										v76 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
										if v76 == int32(0) {
											v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
											mBase = m.M
											v744 = m.ExcPending
											if v744 != 0 {
												return int32(0)
											} else {
												return v743
											}
										} else {
											v81 = F_mul_size(m, v76, int32(56))
											mBase = m.M
											v82 = m.ExcPending
											if v82 != 0 {
												return int32(0)
											} else {
												v83 = F_add_size(m, int32(8), v81)
												mBase = m.M
												v84 = m.ExcPending
												if v84 != 0 {
													return int32(0)
												} else {
													v85 = *(*int32)(unsafe.Add(mBase, uint32(v69)+52))
													v86 = F_shm_toc_allocate(m, v85, v83)
													mBase = m.M
													v87 = m.ExcPending
													if v87 != 0 {
														return int32(0)
													} else {
														if v83 != 0 {
															base.MemoryFill(m, v86, int32(0), v83)
														} else {
														}
														v90 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
														*(*int32)(unsafe.Add(mBase, uint32(v86))) = v90
														v92 = *(*int32)(unsafe.Add(mBase, uint32(v69)+52))
														v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
														v94 = int64(*(*int32)(unsafe.Add(mBase, uint32(v93)+40)))
														F_shm_toc_insert(m, v92, v94-int64(3458764513820540928), v86)
														mBase = m.M
														v98 = m.ExcPending
														if v98 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v86
															v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
															mBase = m.M
															v744 = m.ExcPending
															if v744 != 0 {
																return int32(0)
															} else {
																return v743
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
				v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+132)))
				if v71&int32(16) == int32(0) {
					v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
					mBase = m.M
					v744 = m.ExcPending
					if v744 != 0 {
						return int32(0)
					} else {
						return v743
					}
				} else {
					v76 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
					if v76 == int32(0) {
						v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
						mBase = m.M
						v744 = m.ExcPending
						if v744 != 0 {
							return int32(0)
						} else {
							return v743
						}
					} else {
						v81 = F_mul_size(m, v76, int32(56))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return int32(0)
						} else {
							v83 = F_add_size(m, int32(8), v81)
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int32(0)
							} else {
								v85 = *(*int32)(unsafe.Add(mBase, uint32(v69)+52))
								v86 = F_shm_toc_allocate(m, v85, v83)
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return int32(0)
								} else {
									if v83 != 0 {
										base.MemoryFill(m, v86, int32(0), v83)
									} else {
									}
									v90 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
									*(*int32)(unsafe.Add(mBase, uint32(v86))) = v90
									v92 = *(*int32)(unsafe.Add(mBase, uint32(v69)+52))
									v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v94 = int64(*(*int32)(unsafe.Add(mBase, uint32(v93)+40)))
									F_shm_toc_insert(m, v92, v94-int64(3458764513820540928), v86)
									mBase = m.M
									v98 = m.ExcPending
									if v98 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v86
										v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
										mBase = m.M
										v744 = m.ExcPending
										if v744 != 0 {
											return int32(0)
										} else {
											return v743
										}
									}
								}
							}
						}
					}
				}
			}
		case 8:
			v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+36)))
			if v103 == int32(1) {
				v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v107 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+52))
				v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
				v110 = F_shm_toc_allocate(m, v108, v109)
				mBase = m.M
				v111 = m.ExcPending
				if v111 != 0 {
					return int32(0)
				} else {
					v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
					v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
					v114 = *(*int32)(unsafe.Add(mBase, uint32(v106)+8))
					F_index_parallelscan_initialize(m, v112, v113, v114, v110)
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return int32(0)
					} else {
						v117 = *(*int32)(unsafe.Add(mBase, uint32(v107)+52))
						v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v119 = int64(*(*int32)(unsafe.Add(mBase, uint32(v118)+40)))
						F_shm_toc_insert(m, v117, v119, v110)
						mBase = m.M
						v121 = m.ExcPending
						if v121 != 0 {
							return int32(0)
						} else {
							v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
							v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
							v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
							v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
							v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
							v129 = F_ScanRelIsReadOnly(m, l0)
							mBase = m.M
							v130 = m.ExcPending
							if v130 != 0 {
								return int32(0)
							} else {
								if v129 != 0 {
									v131 = int32(1024)
								} else {
									v131 = int32(0)
								}
								v132 = F_index_beginscan_parallel(m, v122, v123, v124, v125, v126, v110, v131)
								mBase = m.M
								v133 = m.ExcPending
								if v133 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v132
									v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
									if v135 != 0 {
										v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+148)))
										if v136 != int32(1) {
											v148 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
											v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											if v149 == int32(0) {
												v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
												mBase = m.M
												v744 = m.ExcPending
												if v744 != 0 {
													return int32(0)
												} else {
													return v743
												}
											} else {
												v152 = *(*int32)(unsafe.Add(mBase, uint32(v148)+12))
												if v152 == int32(0) {
													v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
													mBase = m.M
													v744 = m.ExcPending
													if v744 != 0 {
														return int32(0)
													} else {
														return v743
													}
												} else {
													v155 = int32(8)
													v157 = F_mul_size(m, v152, v155)
													mBase = m.M
													v158 = m.ExcPending
													if v158 != 0 {
														return int32(0)
													} else {
														v159 = F_add_size(m, v155, v157)
														mBase = m.M
														v160 = m.ExcPending
														if v160 != 0 {
															return int32(0)
														} else {
															v161 = *(*int32)(unsafe.Add(mBase, uint32(v148)+52))
															v162 = F_shm_toc_allocate(m, v161, v159)
															mBase = m.M
															v163 = m.ExcPending
															if v163 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v162
																if v159 != 0 {
																	base.MemoryFill(m, v162, int32(0), v159)
																} else {
																}
																v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
																v168 = *(*int32)(unsafe.Add(mBase, uint32(v148)+12))
																*(*int32)(unsafe.Add(mBase, uint32(v167))) = v168
																v170 = *(*int32)(unsafe.Add(mBase, uint32(v148)+52))
																v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
																v172 = int64(*(*int32)(unsafe.Add(mBase, uint32(v171)+40)))
																v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
																F_shm_toc_insert(m, v170, v172-int64(3458764513820540928), v175)
																mBase = m.M
																v177 = m.ExcPending
																if v177 != 0 {
																	return int32(0)
																} else {
																	v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
																	mBase = m.M
																	v744 = m.ExcPending
																	if v744 != 0 {
																		return int32(0)
																	} else {
																		return v743
																	}
																}
															}
														}
													}
												}
											}
										} else {
											v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
											v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
											v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
											v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
											F_index_rescan(m, v132, v139, v140, v141, v142)
											mBase = m.M
											v144 = m.ExcPending
											if v144 != 0 {
												return int32(0)
											} else {
												v148 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
												v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
												if v149 == int32(0) {
													v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
													mBase = m.M
													v744 = m.ExcPending
													if v744 != 0 {
														return int32(0)
													} else {
														return v743
													}
												} else {
													v152 = *(*int32)(unsafe.Add(mBase, uint32(v148)+12))
													if v152 == int32(0) {
														v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
														mBase = m.M
														v744 = m.ExcPending
														if v744 != 0 {
															return int32(0)
														} else {
															return v743
														}
													} else {
														v155 = int32(8)
														v157 = F_mul_size(m, v152, v155)
														mBase = m.M
														v158 = m.ExcPending
														if v158 != 0 {
															return int32(0)
														} else {
															v159 = F_add_size(m, v155, v157)
															mBase = m.M
															v160 = m.ExcPending
															if v160 != 0 {
																return int32(0)
															} else {
																v161 = *(*int32)(unsafe.Add(mBase, uint32(v148)+52))
																v162 = F_shm_toc_allocate(m, v161, v159)
																mBase = m.M
																v163 = m.ExcPending
																if v163 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v162
																	if v159 != 0 {
																		base.MemoryFill(m, v162, int32(0), v159)
																	} else {
																	}
																	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
																	v168 = *(*int32)(unsafe.Add(mBase, uint32(v148)+12))
																	*(*int32)(unsafe.Add(mBase, uint32(v167))) = v168
																	v170 = *(*int32)(unsafe.Add(mBase, uint32(v148)+52))
																	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
																	v172 = int64(*(*int32)(unsafe.Add(mBase, uint32(v171)+40)))
																	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
																	F_shm_toc_insert(m, v170, v172-int64(3458764513820540928), v175)
																	mBase = m.M
																	v177 = m.ExcPending
																	if v177 != 0 {
																		return int32(0)
																	} else {
																		v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
																		mBase = m.M
																		v744 = m.ExcPending
																		if v744 != 0 {
																			return int32(0)
																		} else {
																			return v743
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
										v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
										v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
										v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
										v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
										F_index_rescan(m, v132, v139, v140, v141, v142)
										mBase = m.M
										v144 = m.ExcPending
										if v144 != 0 {
											return int32(0)
										} else {
											v148 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
											v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											if v149 == int32(0) {
												v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
												mBase = m.M
												v744 = m.ExcPending
												if v744 != 0 {
													return int32(0)
												} else {
													return v743
												}
											} else {
												v152 = *(*int32)(unsafe.Add(mBase, uint32(v148)+12))
												if v152 == int32(0) {
													v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
													mBase = m.M
													v744 = m.ExcPending
													if v744 != 0 {
														return int32(0)
													} else {
														return v743
													}
												} else {
													v155 = int32(8)
													v157 = F_mul_size(m, v152, v155)
													mBase = m.M
													v158 = m.ExcPending
													if v158 != 0 {
														return int32(0)
													} else {
														v159 = F_add_size(m, v155, v157)
														mBase = m.M
														v160 = m.ExcPending
														if v160 != 0 {
															return int32(0)
														} else {
															v161 = *(*int32)(unsafe.Add(mBase, uint32(v148)+52))
															v162 = F_shm_toc_allocate(m, v161, v159)
															mBase = m.M
															v163 = m.ExcPending
															if v163 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v162
																if v159 != 0 {
																	base.MemoryFill(m, v162, int32(0), v159)
																} else {
																}
																v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
																v168 = *(*int32)(unsafe.Add(mBase, uint32(v148)+12))
																*(*int32)(unsafe.Add(mBase, uint32(v167))) = v168
																v170 = *(*int32)(unsafe.Add(mBase, uint32(v148)+52))
																v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
																v172 = int64(*(*int32)(unsafe.Add(mBase, uint32(v171)+40)))
																v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
																F_shm_toc_insert(m, v170, v172-int64(3458764513820540928), v175)
																mBase = m.M
																v177 = m.ExcPending
																if v177 != 0 {
																	return int32(0)
																} else {
																	v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
																	mBase = m.M
																	v744 = m.ExcPending
																	if v744 != 0 {
																		return int32(0)
																	} else {
																		return v743
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
				v148 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v149 == int32(0) {
					v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
					mBase = m.M
					v744 = m.ExcPending
					if v744 != 0 {
						return int32(0)
					} else {
						return v743
					}
				} else {
					v152 = *(*int32)(unsafe.Add(mBase, uint32(v148)+12))
					if v152 == int32(0) {
						v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
						mBase = m.M
						v744 = m.ExcPending
						if v744 != 0 {
							return int32(0)
						} else {
							return v743
						}
					} else {
						v155 = int32(8)
						v157 = F_mul_size(m, v152, v155)
						mBase = m.M
						v158 = m.ExcPending
						if v158 != 0 {
							return int32(0)
						} else {
							v159 = F_add_size(m, v155, v157)
							mBase = m.M
							v160 = m.ExcPending
							if v160 != 0 {
								return int32(0)
							} else {
								v161 = *(*int32)(unsafe.Add(mBase, uint32(v148)+52))
								v162 = F_shm_toc_allocate(m, v161, v159)
								mBase = m.M
								v163 = m.ExcPending
								if v163 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v162
									if v159 != 0 {
										base.MemoryFill(m, v162, int32(0), v159)
									} else {
									}
									v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
									v168 = *(*int32)(unsafe.Add(mBase, uint32(v148)+12))
									*(*int32)(unsafe.Add(mBase, uint32(v167))) = v168
									v170 = *(*int32)(unsafe.Add(mBase, uint32(v148)+52))
									v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v172 = int64(*(*int32)(unsafe.Add(mBase, uint32(v171)+40)))
									v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
									F_shm_toc_insert(m, v170, v172-int64(3458764513820540928), v175)
									mBase = m.M
									v177 = m.ExcPending
									if v177 != 0 {
										return int32(0)
									} else {
										v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
										mBase = m.M
										v744 = m.ExcPending
										if v744 != 0 {
											return int32(0)
										} else {
											return v743
										}
									}
								}
							}
						}
					}
				}
			}
		case 9:
			v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+36)))
			if v181 == int32(1) {
				v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v185 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v186 = *(*int32)(unsafe.Add(mBase, uint32(v185)+52))
				v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
				v188 = F_shm_toc_allocate(m, v186, v187)
				mBase = m.M
				v189 = m.ExcPending
				if v189 != 0 {
					return int32(0)
				} else {
					v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
					v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
					v192 = *(*int32)(unsafe.Add(mBase, uint32(v184)+8))
					F_index_parallelscan_initialize(m, v190, v191, v192, v188)
					mBase = m.M
					v194 = m.ExcPending
					if v194 != 0 {
						return int32(0)
					} else {
						v195 = *(*int32)(unsafe.Add(mBase, uint32(v185)+52))
						v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v197 = int64(*(*int32)(unsafe.Add(mBase, uint32(v196)+40)))
						F_shm_toc_insert(m, v195, v197, v188)
						mBase = m.M
						v199 = m.ExcPending
						if v199 != 0 {
							return int32(0)
						} else {
							v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
							v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
							v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
							v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
							v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
							v207 = F_ScanRelIsReadOnly(m, l0)
							mBase = m.M
							v208 = m.ExcPending
							if v208 != 0 {
								return int32(0)
							} else {
								if v207 != 0 {
									v209 = int32(1024)
								} else {
									v209 = int32(0)
								}
								v210 = F_index_beginscan_parallel(m, v200, v201, v202, v203, v204, v188, v209)
								mBase = m.M
								v211 = m.ExcPending
								if v211 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v210
									v213 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v210)+28)) = uint8(v213)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = int32(0)
									v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
									if v217 != 0 {
										v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+144)))
										if v218 != int32(1) {
											v231 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
											v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											if v232 == int32(0) {
												v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
												mBase = m.M
												v744 = m.ExcPending
												if v744 != 0 {
													return int32(0)
												} else {
													return v743
												}
											} else {
												v235 = *(*int32)(unsafe.Add(mBase, uint32(v231)+12))
												if v235 == int32(0) {
													v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
													mBase = m.M
													v744 = m.ExcPending
													if v744 != 0 {
														return int32(0)
													} else {
														return v743
													}
												} else {
													v238 = int32(8)
													v240 = F_mul_size(m, v235, v238)
													mBase = m.M
													v241 = m.ExcPending
													if v241 != 0 {
														return int32(0)
													} else {
														v242 = F_add_size(m, v238, v240)
														mBase = m.M
														v243 = m.ExcPending
														if v243 != 0 {
															return int32(0)
														} else {
															v244 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
															v245 = F_shm_toc_allocate(m, v244, v242)
															mBase = m.M
															v246 = m.ExcPending
															if v246 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v245
																if v242 != 0 {
																	base.MemoryFill(m, v245, int32(0), v242)
																} else {
																}
																v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
																v251 = *(*int32)(unsafe.Add(mBase, uint32(v231)+12))
																*(*int32)(unsafe.Add(mBase, uint32(v250))) = v251
																v253 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
																v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
																v255 = int64(*(*int32)(unsafe.Add(mBase, uint32(v254)+40)))
																v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
																F_shm_toc_insert(m, v253, v255-int64(3458764513820540928), v258)
																mBase = m.M
																v260 = m.ExcPending
																if v260 != 0 {
																	return int32(0)
																} else {
																	v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
																	mBase = m.M
																	v744 = m.ExcPending
																	if v744 != 0 {
																		return int32(0)
																	} else {
																		return v743
																	}
																}
															}
														}
													}
												}
											}
										} else {
											v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
											v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
											v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
											v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
											v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
											F_index_rescan(m, v221, v222, v223, v224, v225)
											mBase = m.M
											v227 = m.ExcPending
											if v227 != 0 {
												return int32(0)
											} else {
												v231 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
												v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
												if v232 == int32(0) {
													v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
													mBase = m.M
													v744 = m.ExcPending
													if v744 != 0 {
														return int32(0)
													} else {
														return v743
													}
												} else {
													v235 = *(*int32)(unsafe.Add(mBase, uint32(v231)+12))
													if v235 == int32(0) {
														v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
														mBase = m.M
														v744 = m.ExcPending
														if v744 != 0 {
															return int32(0)
														} else {
															return v743
														}
													} else {
														v238 = int32(8)
														v240 = F_mul_size(m, v235, v238)
														mBase = m.M
														v241 = m.ExcPending
														if v241 != 0 {
															return int32(0)
														} else {
															v242 = F_add_size(m, v238, v240)
															mBase = m.M
															v243 = m.ExcPending
															if v243 != 0 {
																return int32(0)
															} else {
																v244 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
																v245 = F_shm_toc_allocate(m, v244, v242)
																mBase = m.M
																v246 = m.ExcPending
																if v246 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v245
																	if v242 != 0 {
																		base.MemoryFill(m, v245, int32(0), v242)
																	} else {
																	}
																	v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
																	v251 = *(*int32)(unsafe.Add(mBase, uint32(v231)+12))
																	*(*int32)(unsafe.Add(mBase, uint32(v250))) = v251
																	v253 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
																	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
																	v255 = int64(*(*int32)(unsafe.Add(mBase, uint32(v254)+40)))
																	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
																	F_shm_toc_insert(m, v253, v255-int64(3458764513820540928), v258)
																	mBase = m.M
																	v260 = m.ExcPending
																	if v260 != 0 {
																		return int32(0)
																	} else {
																		v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
																		mBase = m.M
																		v744 = m.ExcPending
																		if v744 != 0 {
																			return int32(0)
																		} else {
																			return v743
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
										v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
										v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
										v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
										v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
										v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
										F_index_rescan(m, v221, v222, v223, v224, v225)
										mBase = m.M
										v227 = m.ExcPending
										if v227 != 0 {
											return int32(0)
										} else {
											v231 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
											v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											if v232 == int32(0) {
												v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
												mBase = m.M
												v744 = m.ExcPending
												if v744 != 0 {
													return int32(0)
												} else {
													return v743
												}
											} else {
												v235 = *(*int32)(unsafe.Add(mBase, uint32(v231)+12))
												if v235 == int32(0) {
													v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
													mBase = m.M
													v744 = m.ExcPending
													if v744 != 0 {
														return int32(0)
													} else {
														return v743
													}
												} else {
													v238 = int32(8)
													v240 = F_mul_size(m, v235, v238)
													mBase = m.M
													v241 = m.ExcPending
													if v241 != 0 {
														return int32(0)
													} else {
														v242 = F_add_size(m, v238, v240)
														mBase = m.M
														v243 = m.ExcPending
														if v243 != 0 {
															return int32(0)
														} else {
															v244 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
															v245 = F_shm_toc_allocate(m, v244, v242)
															mBase = m.M
															v246 = m.ExcPending
															if v246 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v245
																if v242 != 0 {
																	base.MemoryFill(m, v245, int32(0), v242)
																} else {
																}
																v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
																v251 = *(*int32)(unsafe.Add(mBase, uint32(v231)+12))
																*(*int32)(unsafe.Add(mBase, uint32(v250))) = v251
																v253 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
																v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
																v255 = int64(*(*int32)(unsafe.Add(mBase, uint32(v254)+40)))
																v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
																F_shm_toc_insert(m, v253, v255-int64(3458764513820540928), v258)
																mBase = m.M
																v260 = m.ExcPending
																if v260 != 0 {
																	return int32(0)
																} else {
																	v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
																	mBase = m.M
																	v744 = m.ExcPending
																	if v744 != 0 {
																		return int32(0)
																	} else {
																		return v743
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
				v231 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v232 == int32(0) {
					v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
					mBase = m.M
					v744 = m.ExcPending
					if v744 != 0 {
						return int32(0)
					} else {
						return v743
					}
				} else {
					v235 = *(*int32)(unsafe.Add(mBase, uint32(v231)+12))
					if v235 == int32(0) {
						v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
						mBase = m.M
						v744 = m.ExcPending
						if v744 != 0 {
							return int32(0)
						} else {
							return v743
						}
					} else {
						v238 = int32(8)
						v240 = F_mul_size(m, v235, v238)
						mBase = m.M
						v241 = m.ExcPending
						if v241 != 0 {
							return int32(0)
						} else {
							v242 = F_add_size(m, v238, v240)
							mBase = m.M
							v243 = m.ExcPending
							if v243 != 0 {
								return int32(0)
							} else {
								v244 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
								v245 = F_shm_toc_allocate(m, v244, v242)
								mBase = m.M
								v246 = m.ExcPending
								if v246 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v245
									if v242 != 0 {
										base.MemoryFill(m, v245, int32(0), v242)
									} else {
									}
									v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
									v251 = *(*int32)(unsafe.Add(mBase, uint32(v231)+12))
									*(*int32)(unsafe.Add(mBase, uint32(v250))) = v251
									v253 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
									v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v255 = int64(*(*int32)(unsafe.Add(mBase, uint32(v254)+40)))
									v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
									F_shm_toc_insert(m, v253, v255-int64(3458764513820540928), v258)
									mBase = m.M
									v260 = m.ExcPending
									if v260 != 0 {
										return int32(0)
									} else {
										v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
										mBase = m.M
										v744 = m.ExcPending
										if v744 != 0 {
											return int32(0)
										} else {
											return v743
										}
									}
								}
							}
						}
					}
				}
			}
		case 10:
			v263 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v264 == int32(0) {
				v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
				mBase = m.M
				v744 = m.ExcPending
				if v744 != 0 {
					return int32(0)
				} else {
					return v743
				}
			} else {
				v267 = *(*int32)(unsafe.Add(mBase, uint32(v263)+12))
				if v267 == int32(0) {
					v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
					mBase = m.M
					v744 = m.ExcPending
					if v744 != 0 {
						return int32(0)
					} else {
						return v743
					}
				} else {
					v270 = *(*int32)(unsafe.Add(mBase, uint32(v263)+52))
					v274 = v267<<(uint(int32(3))%32) + int32(8)
					v275 = F_shm_toc_allocate(m, v270, v274)
					mBase = m.M
					v276 = m.ExcPending
					if v276 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v275
						v278 = *(*int32)(unsafe.Add(mBase, uint32(v263)+52))
						v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v280 = int64(*(*int32)(unsafe.Add(mBase, uint32(v279)+40)))
						F_shm_toc_insert(m, v278, v280-int64(3458764513820540928), v275)
						mBase = m.M
						v284 = m.ExcPending
						if v284 != 0 {
							return int32(0)
						} else {
							if v274 != 0 {
								v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
								base.MemoryFill(m, v285, int32(0), v274)
							} else {
							}
							v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
							v289 = *(*int32)(unsafe.Add(mBase, uint32(v263)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v288))) = v289
							v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
							mBase = m.M
							v744 = m.ExcPending
							if v744 != 0 {
								return int32(0)
							} else {
								return v743
							}
						}
					}
				}
			}
		case 11:
			v430 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v430)+36)))
			if v431 == int32(1) {
				v434 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v435 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v436 = *(*int32)(unsafe.Add(mBase, uint32(v435)+172))
				if v436 != 0 {
					v437 = *(*int32)(unsafe.Add(mBase, uint32(v434)+52))
					v439 = F_shm_toc_allocate(m, v437, int32(24))
					mBase = m.M
					v440 = m.ExcPending
					if v440 != 0 {
						return int32(0)
					} else {
						v441 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v439))) = v441
						atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v439)+4)), uint32(v441))
						*(*int32)(unsafe.Add(mBase, uint32(v439)+8)) = v441
						v449 = v439 + int32(12)
						atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v449))), uint32(v441))
						*(*int64)(unsafe.Add(mBase, uint32(v449)+4)) = int64(-1)
						v455 = *(*int32)(unsafe.Add(mBase, uint32(v434)+52))
						v456 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v457 = int64(*(*int32)(unsafe.Add(mBase, uint32(v456)+40)))
						F_shm_toc_insert(m, v455, v457, v439)
						mBase = m.M
						v459 = m.ExcPending
						if v459 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v439
							v464 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							v465 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							if v465 == int32(0) {
								v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
								mBase = m.M
								v744 = m.ExcPending
								if v744 != 0 {
									return int32(0)
								} else {
									return v743
								}
							} else {
								v468 = *(*int32)(unsafe.Add(mBase, uint32(v464)+12))
								if v468 == int32(0) {
									v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
									mBase = m.M
									v744 = m.ExcPending
									if v744 != 0 {
										return int32(0)
									} else {
										return v743
									}
								} else {
									v473 = F_mul_size(m, v468, int32(72))
									mBase = m.M
									v474 = m.ExcPending
									if v474 != 0 {
										return int32(0)
									} else {
										v475 = F_add_size(m, int32(8), v473)
										mBase = m.M
										v476 = m.ExcPending
										if v476 != 0 {
											return int32(0)
										} else {
											v477 = *(*int32)(unsafe.Add(mBase, uint32(v464)+52))
											v478 = F_shm_toc_allocate(m, v477, v475)
											mBase = m.M
											v479 = m.ExcPending
											if v479 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = v478
												if v475 != 0 {
													base.MemoryFill(m, v478, int32(0), v475)
												} else {
												}
												v483 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
												v484 = *(*int32)(unsafe.Add(mBase, uint32(v464)+12))
												*(*int32)(unsafe.Add(mBase, uint32(v483))) = v484
												v486 = *(*int32)(unsafe.Add(mBase, uint32(v464)+52))
												v487 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												v488 = int64(*(*int32)(unsafe.Add(mBase, uint32(v487)+40)))
												v491 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
												F_shm_toc_insert(m, v486, v488-int64(3458764513820540928), v491)
												mBase = m.M
												v493 = m.ExcPending
												if v493 != 0 {
													return int32(0)
												} else {
													v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
													mBase = m.M
													v744 = m.ExcPending
													if v744 != 0 {
														return int32(0)
													} else {
														return v743
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
					v464 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v465 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v465 == int32(0) {
						v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
						mBase = m.M
						v744 = m.ExcPending
						if v744 != 0 {
							return int32(0)
						} else {
							return v743
						}
					} else {
						v468 = *(*int32)(unsafe.Add(mBase, uint32(v464)+12))
						if v468 == int32(0) {
							v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
							mBase = m.M
							v744 = m.ExcPending
							if v744 != 0 {
								return int32(0)
							} else {
								return v743
							}
						} else {
							v473 = F_mul_size(m, v468, int32(72))
							mBase = m.M
							v474 = m.ExcPending
							if v474 != 0 {
								return int32(0)
							} else {
								v475 = F_add_size(m, int32(8), v473)
								mBase = m.M
								v476 = m.ExcPending
								if v476 != 0 {
									return int32(0)
								} else {
									v477 = *(*int32)(unsafe.Add(mBase, uint32(v464)+52))
									v478 = F_shm_toc_allocate(m, v477, v475)
									mBase = m.M
									v479 = m.ExcPending
									if v479 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = v478
										if v475 != 0 {
											base.MemoryFill(m, v478, int32(0), v475)
										} else {
										}
										v483 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
										v484 = *(*int32)(unsafe.Add(mBase, uint32(v464)+12))
										*(*int32)(unsafe.Add(mBase, uint32(v483))) = v484
										v486 = *(*int32)(unsafe.Add(mBase, uint32(v464)+52))
										v487 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										v488 = int64(*(*int32)(unsafe.Add(mBase, uint32(v487)+40)))
										v491 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
										F_shm_toc_insert(m, v486, v488-int64(3458764513820540928), v491)
										mBase = m.M
										v493 = m.ExcPending
										if v493 != 0 {
											return int32(0)
										} else {
											v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
											mBase = m.M
											v744 = m.ExcPending
											if v744 != 0 {
												return int32(0)
											} else {
												return v743
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v464 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v465 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v465 == int32(0) {
					v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
					mBase = m.M
					v744 = m.ExcPending
					if v744 != 0 {
						return int32(0)
					} else {
						return v743
					}
				} else {
					v468 = *(*int32)(unsafe.Add(mBase, uint32(v464)+12))
					if v468 == int32(0) {
						v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
						mBase = m.M
						v744 = m.ExcPending
						if v744 != 0 {
							return int32(0)
						} else {
							return v743
						}
					} else {
						v473 = F_mul_size(m, v468, int32(72))
						mBase = m.M
						v474 = m.ExcPending
						if v474 != 0 {
							return int32(0)
						} else {
							v475 = F_add_size(m, int32(8), v473)
							mBase = m.M
							v476 = m.ExcPending
							if v476 != 0 {
								return int32(0)
							} else {
								v477 = *(*int32)(unsafe.Add(mBase, uint32(v464)+52))
								v478 = F_shm_toc_allocate(m, v477, v475)
								mBase = m.M
								v479 = m.ExcPending
								if v479 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = v478
									if v475 != 0 {
										base.MemoryFill(m, v478, int32(0), v475)
									} else {
									}
									v483 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
									v484 = *(*int32)(unsafe.Add(mBase, uint32(v464)+12))
									*(*int32)(unsafe.Add(mBase, uint32(v483))) = v484
									v486 = *(*int32)(unsafe.Add(mBase, uint32(v464)+52))
									v487 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v488 = int64(*(*int32)(unsafe.Add(mBase, uint32(v487)+40)))
									v491 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
									F_shm_toc_insert(m, v486, v488-int64(3458764513820540928), v491)
									mBase = m.M
									v493 = m.ExcPending
									if v493 != 0 {
										return int32(0)
									} else {
										v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
										mBase = m.M
										v744 = m.ExcPending
										if v744 != 0 {
											return int32(0)
										} else {
											return v743
										}
									}
								}
							}
						}
					}
				}
			}
		case 13:
			v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314)+36)))
			if v315 == int32(1) {
				v318 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v319 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v320 = F_ScanRelIsReadOnly(m, l0)
				mBase = m.M
				v321 = m.ExcPending
				if v321 != 0 {
					return int32(0)
				} else {
					v322 = *(*int32)(unsafe.Add(mBase, uint32(v319)+132))
					v323 = *(*int32)(unsafe.Add(mBase, uint32(v318)+52))
					v324 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
					v325 = F_shm_toc_allocate(m, v323, v324)
					mBase = m.M
					v326 = m.ExcPending
					if v326 != 0 {
						return int32(0)
					} else {
						v327 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
						v328 = *(*int32)(unsafe.Add(mBase, uint32(v319)+8))
						F_table_parallelscan_initialize(m, v327, v325, v328)
						mBase = m.M
						v330 = m.ExcPending
						if v330 != 0 {
							return int32(0)
						} else {
							v331 = *(*int32)(unsafe.Add(mBase, uint32(v318)+52))
							v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v333 = int64(*(*int32)(unsafe.Add(mBase, uint32(v332)+40)))
							F_shm_toc_insert(m, v331, v333, v325)
							mBase = m.M
							v335 = m.ExcPending
							if v335 != 0 {
								return int32(0)
							} else {
								v336 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
								if v320 != 0 {
									v343 = int32(1024)
								} else {
									v343 = int32(0)
								}
								v345 = F_table_beginscan_parallel_tidrange(m, v336, v325, v322<<(uint(int32(7))%32)&int32(2048)|v343)
								mBase = m.M
								v346 = m.ExcPending
								if v346 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v345
									v353 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354)+132)))
									if v355&int32(16) == int32(0) {
										v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
										mBase = m.M
										v744 = m.ExcPending
										if v744 != 0 {
											return int32(0)
										} else {
											return v743
										}
									} else {
										v360 = *(*int32)(unsafe.Add(mBase, uint32(v353)+12))
										if v360 == int32(0) {
											v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
											mBase = m.M
											v744 = m.ExcPending
											if v744 != 0 {
												return int32(0)
											} else {
												return v743
											}
										} else {
											v365 = F_mul_size(m, v360, int32(56))
											mBase = m.M
											v366 = m.ExcPending
											if v366 != 0 {
												return int32(0)
											} else {
												v367 = F_add_size(m, int32(8), v365)
												mBase = m.M
												v368 = m.ExcPending
												if v368 != 0 {
													return int32(0)
												} else {
													v369 = *(*int32)(unsafe.Add(mBase, uint32(v353)+52))
													v370 = F_shm_toc_allocate(m, v369, v367)
													mBase = m.M
													v371 = m.ExcPending
													if v371 != 0 {
														return int32(0)
													} else {
														if v367 != 0 {
															base.MemoryFill(m, v370, int32(0), v367)
														} else {
														}
														v374 = *(*int32)(unsafe.Add(mBase, uint32(v353)+12))
														*(*int32)(unsafe.Add(mBase, uint32(v370))) = v374
														v376 = *(*int32)(unsafe.Add(mBase, uint32(v353)+52))
														v377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
														v378 = int64(*(*int32)(unsafe.Add(mBase, uint32(v377)+40)))
														F_shm_toc_insert(m, v376, v378-int64(3458764513820540928), v370)
														mBase = m.M
														v382 = m.ExcPending
														if v382 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v370
															v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
															mBase = m.M
															v744 = m.ExcPending
															if v744 != 0 {
																return int32(0)
															} else {
																return v743
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
				v353 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354)+132)))
				if v355&int32(16) == int32(0) {
					v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
					mBase = m.M
					v744 = m.ExcPending
					if v744 != 0 {
						return int32(0)
					} else {
						return v743
					}
				} else {
					v360 = *(*int32)(unsafe.Add(mBase, uint32(v353)+12))
					if v360 == int32(0) {
						v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
						mBase = m.M
						v744 = m.ExcPending
						if v744 != 0 {
							return int32(0)
						} else {
							return v743
						}
					} else {
						v365 = F_mul_size(m, v360, int32(56))
						mBase = m.M
						v366 = m.ExcPending
						if v366 != 0 {
							return int32(0)
						} else {
							v367 = F_add_size(m, int32(8), v365)
							mBase = m.M
							v368 = m.ExcPending
							if v368 != 0 {
								return int32(0)
							} else {
								v369 = *(*int32)(unsafe.Add(mBase, uint32(v353)+52))
								v370 = F_shm_toc_allocate(m, v369, v367)
								mBase = m.M
								v371 = m.ExcPending
								if v371 != 0 {
									return int32(0)
								} else {
									if v367 != 0 {
										base.MemoryFill(m, v370, int32(0), v367)
									} else {
									}
									v374 = *(*int32)(unsafe.Add(mBase, uint32(v353)+12))
									*(*int32)(unsafe.Add(mBase, uint32(v370))) = v374
									v376 = *(*int32)(unsafe.Add(mBase, uint32(v353)+52))
									v377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v378 = int64(*(*int32)(unsafe.Add(mBase, uint32(v377)+40)))
									F_shm_toc_insert(m, v376, v378-int64(3458764513820540928), v370)
									mBase = m.M
									v382 = m.ExcPending
									if v382 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v370
										v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
										mBase = m.M
										v744 = m.ExcPending
										if v744 != 0 {
											return int32(0)
										} else {
											return v743
										}
									}
								}
							}
						}
					}
				}
			}
		case 21:
			v293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293)+36)))
			if v294 != int32(1) {
				v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
				mBase = m.M
				v744 = m.ExcPending
				if v744 != 0 {
					return int32(0)
				} else {
					return v743
				}
			} else {
				v297 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v299 = *(*int32)(unsafe.Add(mBase, uint32(v298)+152))
				if v299 != 0 {
					v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v301 = int64(*(*int32)(unsafe.Add(mBase, uint32(v300)+40)))
					v302 = *(*int32)(unsafe.Add(mBase, uint32(v297)+52))
					v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
					v304 = F_shm_toc_allocate(m, v302, v303)
					mBase = m.M
					v305 = m.ExcPending
					if v305 != 0 {
						return int32(0)
					} else {
						v306 = *(*int32)(unsafe.Add(mBase, uint32(v298)+152))
						m.T0[v306].(func(*base.Module, int32, int32, int32))(m, l0, v297, v304)
						mBase = m.M
						v308 = m.ExcPending
						if v308 != 0 {
							return int32(0)
						} else {
							v309 = *(*int32)(unsafe.Add(mBase, uint32(v297)+52))
							F_shm_toc_insert(m, v309, v301, v304)
							mBase = m.M
							v311 = m.ExcPending
							if v311 != 0 {
								return int32(0)
							} else {
								v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
								mBase = m.M
								v744 = m.ExcPending
								if v744 != 0 {
									return int32(0)
								} else {
									return v743
								}
							}
						}
					}
				} else {
					v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
					mBase = m.M
					v744 = m.ExcPending
					if v744 != 0 {
						return int32(0)
					} else {
						return v743
					}
				}
			}
		case 22:
			v409 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409)+36)))
			if v410 != int32(1) {
				v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
				mBase = m.M
				v744 = m.ExcPending
				if v744 != 0 {
					return int32(0)
				} else {
					return v743
				}
			} else {
				v413 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v414 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v415 = *(*int32)(unsafe.Add(mBase, uint32(v414)+32))
				if v415 != 0 {
					v416 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v417 = int64(*(*int32)(unsafe.Add(mBase, uint32(v416)+40)))
					v418 = *(*int32)(unsafe.Add(mBase, uint32(v413)+52))
					v419 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
					v420 = F_shm_toc_allocate(m, v418, v419)
					mBase = m.M
					v421 = m.ExcPending
					if v421 != 0 {
						return int32(0)
					} else {
						v422 = *(*int32)(unsafe.Add(mBase, uint32(v414)+32))
						m.T0[v422].(func(*base.Module, int32, int32, int32))(m, l0, v413, v420)
						mBase = m.M
						v424 = m.ExcPending
						if v424 != 0 {
							return int32(0)
						} else {
							v425 = *(*int32)(unsafe.Add(mBase, uint32(v413)+52))
							F_shm_toc_insert(m, v425, v417, v420)
							mBase = m.M
							v427 = m.ExcPending
							if v427 != 0 {
								return int32(0)
							} else {
								v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
								mBase = m.M
								v744 = m.ExcPending
								if v744 != 0 {
									return int32(0)
								} else {
									return v743
								}
							}
						}
					}
				} else {
					v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
					mBase = m.M
					v744 = m.ExcPending
					if v744 != 0 {
						return int32(0)
					} else {
						return v743
					}
				}
			}
		case 26:
			v496 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+36)))
			if v497 != int32(1) {
				v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
				mBase = m.M
				v744 = m.ExcPending
				if v744 != 0 {
					return int32(0)
				} else {
					return v743
				}
			} else {
				v500 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v501 = *(*int32)(unsafe.Add(mBase, uint32(v500)+44))
				if v501 != 0 {
					v502 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v503 = int64(*(*int32)(unsafe.Add(mBase, uint32(v502)+40)))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(678)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(766)
					v508 = *(*int32)(unsafe.Add(mBase, uint32(v500)+52))
					v510 = F_shm_toc_allocate(m, v508, int32(220))
					mBase = m.M
					v511 = m.ExcPending
					if v511 != 0 {
						return int32(0)
					} else {
						v512 = *(*int32)(unsafe.Add(mBase, uint32(v500)+52))
						F_shm_toc_insert(m, v512, v503, v510)
						mBase = m.M
						v514 = m.ExcPending
						if v514 != 0 {
							return int32(0)
						} else {
							v515 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v510)+32)) = v515
							*(*int32)(unsafe.Add(mBase, uint32(v510)+8)) = v515
							*(*int32)(unsafe.Add(mBase, uint32(v510)+164)) = v515
							*(*int32)(unsafe.Add(mBase, uint32(v510)+24)) = v515
							v523 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v510)+16)) = v523
							*(*int64)(unsafe.Add(mBase, uint32(v510))) = v523
							v527 = *(*int32)(unsafe.Add(mBase, uint32(v500)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v510)+36)) = v515
							*(*int32)(unsafe.Add(mBase, uint32(v510)+28)) = v527 + int32(1)
							F_LWLockInitialize(m, v510+int32(40), int32(73))
							mBase = m.M
							v537 = m.ExcPending
							if v537 != 0 {
								return int32(0)
							} else {
								v539 = v510 + int32(56)
								v540 = int32(0)
								atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v539))), uint32(v540))
								*(*int32)(unsafe.Add(mBase, uint32(v539)+8)) = v540
								*(*uint8)(unsafe.Add(mBase, uint32(v539)+20)) = uint8(v540)
								*(*int64)(unsafe.Add(mBase, uint32(v539)+12)) = int64(0)
								*(*int32)(unsafe.Add(mBase, uint32(v539)+4)) = v540
								F_ConditionVariableInit(m, v510+int32(80))
								mBase = m.M
								v555 = v510 + int32(92)
								v556 = int32(0)
								atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v555))), uint32(v556))
								*(*int32)(unsafe.Add(mBase, uint32(v555)+8)) = v556
								*(*uint8)(unsafe.Add(mBase, uint32(v555)+20)) = uint8(v556)
								*(*int64)(unsafe.Add(mBase, uint32(v555)+12)) = int64(0)
								*(*int32)(unsafe.Add(mBase, uint32(v555)+4)) = v556
								F_ConditionVariableInit(m, v510+int32(116))
								mBase = m.M
								v571 = v510 + int32(128)
								v572 = int32(0)
								atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v571))), uint32(v572))
								*(*int32)(unsafe.Add(mBase, uint32(v571)+8)) = v572
								*(*uint8)(unsafe.Add(mBase, uint32(v571)+20)) = uint8(v572)
								*(*int64)(unsafe.Add(mBase, uint32(v571)+12)) = int64(0)
								*(*int32)(unsafe.Add(mBase, uint32(v571)+4)) = v572
								F_ConditionVariableInit(m, v510+int32(152))
								mBase = m.M
								v588 = *(*int32)(unsafe.Add(mBase, uint32(v500)+44))
								F_SharedFileSetInit(m, v510+int32(168), v588)
								mBase = m.M
								v590 = m.ExcPending
								if v590 != 0 {
									return int32(0)
								} else {
									v591 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									*(*int32)(unsafe.Add(mBase, uint32(v591)+136)) = v510
									v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
									mBase = m.M
									v744 = m.ExcPending
									if v744 != 0 {
										return int32(0)
									} else {
										return v743
									}
								}
							}
						}
					}
				} else {
					v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
					mBase = m.M
					v744 = m.ExcPending
					if v744 != 0 {
						return int32(0)
					} else {
						return v743
					}
				}
			}
		case 28:
			v708 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v709 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v709 == int32(0) {
				v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
				mBase = m.M
				v744 = m.ExcPending
				if v744 != 0 {
					return int32(0)
				} else {
					return v743
				}
			} else {
				v712 = *(*int32)(unsafe.Add(mBase, uint32(v708)+12))
				if v712 == int32(0) {
					v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
					mBase = m.M
					v744 = m.ExcPending
					if v744 != 0 {
						return int32(0)
					} else {
						return v743
					}
				} else {
					v715 = *(*int32)(unsafe.Add(mBase, uint32(v708)+52))
					v719 = v712*int32(40) + int32(8)
					v720 = F_shm_toc_allocate(m, v715, v719)
					mBase = m.M
					v721 = m.ExcPending
					if v721 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = v720
						if v719 != 0 {
							base.MemoryFill(m, v720, int32(0), v719)
						} else {
						}
						v725 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
						v726 = *(*int32)(unsafe.Add(mBase, uint32(v708)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v725))) = v726
						v728 = *(*int32)(unsafe.Add(mBase, uint32(v708)+52))
						v729 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v730 = int64(*(*int32)(unsafe.Add(mBase, uint32(v729)+40)))
						v731 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
						F_shm_toc_insert(m, v728, v730, v731)
						mBase = m.M
						v733 = m.ExcPending
						if v733 != 0 {
							return int32(0)
						} else {
							v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
							mBase = m.M
							v744 = m.ExcPending
							if v744 != 0 {
								return int32(0)
							} else {
								return v743
							}
						}
					}
				}
			}
		case 29:
			v624 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v625 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v625 == int32(0) {
				v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
				mBase = m.M
				v744 = m.ExcPending
				if v744 != 0 {
					return int32(0)
				} else {
					return v743
				}
			} else {
				v628 = *(*int32)(unsafe.Add(mBase, uint32(v624)+12))
				if v628 == int32(0) {
					v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
					mBase = m.M
					v744 = m.ExcPending
					if v744 != 0 {
						return int32(0)
					} else {
						return v743
					}
				} else {
					v631 = *(*int32)(unsafe.Add(mBase, uint32(v624)+52))
					v635 = v628<<(uint(int32(4))%32) | int32(8)
					v636 = F_shm_toc_allocate(m, v631, v635)
					mBase = m.M
					v637 = m.ExcPending
					if v637 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v636
						if v635 != 0 {
							base.MemoryFill(m, v636, int32(0), v635)
						} else {
						}
						v641 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
						v642 = *(*int32)(unsafe.Add(mBase, uint32(v624)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v641))) = v642
						v644 = *(*int32)(unsafe.Add(mBase, uint32(v624)+52))
						v645 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v646 = int64(*(*int32)(unsafe.Add(mBase, uint32(v645)+40)))
						v647 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
						F_shm_toc_insert(m, v644, v646, v647)
						mBase = m.M
						v649 = m.ExcPending
						if v649 != 0 {
							return int32(0)
						} else {
							v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
							mBase = m.M
							v744 = m.ExcPending
							if v744 != 0 {
								return int32(0)
							} else {
								return v743
							}
						}
					}
				}
			}
		case 30:
			v652 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v653 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v653 == int32(0) {
				v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
				mBase = m.M
				v744 = m.ExcPending
				if v744 != 0 {
					return int32(0)
				} else {
					return v743
				}
			} else {
				v656 = *(*int32)(unsafe.Add(mBase, uint32(v652)+12))
				if v656 == int32(0) {
					v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
					mBase = m.M
					v744 = m.ExcPending
					if v744 != 0 {
						return int32(0)
					} else {
						return v743
					}
				} else {
					v659 = *(*int32)(unsafe.Add(mBase, uint32(v652)+52))
					v663 = v656*int32(96) | int32(8)
					v664 = F_shm_toc_allocate(m, v659, v663)
					mBase = m.M
					v665 = m.ExcPending
					if v665 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+284)) = v664
						if v663 != 0 {
							base.MemoryFill(m, v664, int32(0), v663)
						} else {
						}
						v669 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
						v670 = *(*int32)(unsafe.Add(mBase, uint32(v652)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v669))) = v670
						v672 = *(*int32)(unsafe.Add(mBase, uint32(v652)+52))
						v673 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v674 = int64(*(*int32)(unsafe.Add(mBase, uint32(v673)+40)))
						v675 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
						F_shm_toc_insert(m, v672, v674, v675)
						mBase = m.M
						v677 = m.ExcPending
						if v677 != 0 {
							return int32(0)
						} else {
							v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
							mBase = m.M
							v744 = m.ExcPending
							if v744 != 0 {
								return int32(0)
							} else {
								return v743
							}
						}
					}
				}
			}
		case 32:
			v680 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v681 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v681 == int32(0) {
				v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
				mBase = m.M
				v744 = m.ExcPending
				if v744 != 0 {
					return int32(0)
				} else {
					return v743
				}
			} else {
				v684 = *(*int32)(unsafe.Add(mBase, uint32(v680)+12))
				if v684 == int32(0) {
					v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
					mBase = m.M
					v744 = m.ExcPending
					if v744 != 0 {
						return int32(0)
					} else {
						return v743
					}
				} else {
					v687 = *(*int32)(unsafe.Add(mBase, uint32(v680)+52))
					v691 = v684*int32(24) + int32(8)
					v692 = F_shm_toc_allocate(m, v687, v691)
					mBase = m.M
					v693 = m.ExcPending
					if v693 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+352)) = v692
						if v691 != 0 {
							base.MemoryFill(m, v692, int32(0), v691)
						} else {
						}
						v697 = *(*int32)(unsafe.Add(mBase, uint32(l0)+352))
						v698 = *(*int32)(unsafe.Add(mBase, uint32(v680)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v697))) = v698
						v700 = *(*int32)(unsafe.Add(mBase, uint32(v680)+52))
						v701 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v702 = int64(*(*int32)(unsafe.Add(mBase, uint32(v701)+40)))
						v703 = *(*int32)(unsafe.Add(mBase, uint32(l0)+352))
						F_shm_toc_insert(m, v700, v702, v703)
						mBase = m.M
						v705 = m.ExcPending
						if v705 != 0 {
							return int32(0)
						} else {
							v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
							mBase = m.M
							v744 = m.ExcPending
							if v744 != 0 {
								return int32(0)
							} else {
								return v743
							}
						}
					}
				}
			}
		case 37:
			v596 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v597 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v597 == int32(0) {
				v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
				mBase = m.M
				v744 = m.ExcPending
				if v744 != 0 {
					return int32(0)
				} else {
					return v743
				}
			} else {
				v600 = *(*int32)(unsafe.Add(mBase, uint32(v596)+12))
				if v600 == int32(0) {
					v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
					mBase = m.M
					v744 = m.ExcPending
					if v744 != 0 {
						return int32(0)
					} else {
						return v743
					}
				} else {
					v603 = *(*int32)(unsafe.Add(mBase, uint32(v596)+52))
					v607 = v600*int32(20) + int32(4)
					v608 = F_shm_toc_allocate(m, v603, v607)
					mBase = m.M
					v609 = m.ExcPending
					if v609 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v608
						if v607 != 0 {
							base.MemoryFill(m, v608, int32(0), v607)
						} else {
						}
						v613 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
						v614 = *(*int32)(unsafe.Add(mBase, uint32(v596)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v613))) = v614
						v616 = *(*int32)(unsafe.Add(mBase, uint32(v596)+52))
						v617 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v618 = int64(*(*int32)(unsafe.Add(mBase, uint32(v617)+40)))
						v619 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
						F_shm_toc_insert(m, v616, v618, v619)
						mBase = m.M
						v621 = m.ExcPending
						if v621 != 0 {
							return int32(0)
						} else {
							v743 = F_planstate_tree_walker_impl(m, l0, int32(673), l1)
							mBase = m.M
							v744 = m.ExcPending
							if v744 != 0 {
								return int32(0)
							} else {
								return v743
							}
						}
					}
				}
			}
		}
	}
}
func F_ParallelWorkerMain(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int64
	_ = v47
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
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
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int64
	_ = v95
	var v97 int64
	_ = v97
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
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
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v383 int64
	_ = v383
	var v386 int32
	_ = v386
	var v387 int64
	_ = v387
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v427 int32
	_ = v427
	var v433 int32
	_ = v433
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v592 int32
	_ = v592
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v620 int32
	_ = v620
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v677 int64
	_ = v677
	var v679 int32
	_ = v679
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
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v775 int32
	_ = v775
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v828 int32
	_ = v828
	var v832 int32
	_ = v832
	var v835 int32
	_ = v835
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v843 int32
	_ = v843
	var v848 int32
	_ = v848
	var v853 int32
	_ = v853
	var v856 int32
	_ = v856
	var v860 int32
	_ = v860
	var v864 int32
	_ = v864
	var v871 int32
	_ = v871
	var v874 int32
	_ = v874
	var v878 int32
	_ = v878
	var v883 int32
	_ = v883
	var v905 int32
	_ = v905
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v918 int32
	_ = v918
	var v921 int32
	_ = v921
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v952 int32
	_ = v952
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v961 int32
	_ = v961
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v984 int32
	_ = v984
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1002 int32
	_ = v1002
	var v1011 int32
	_ = v1011
	var v1028 int32
	_ = v1028
	var v1036 int32
	_ = v1036
	var v1039 int32
	_ = v1039
	var v1043 int32
	_ = v1043
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
	var v1062 int32
	_ = v1062
	var v1065 int32
	_ = v1065
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1073 int32
	_ = v1073
	var v1075 int32
	_ = v1075
	var v1079 int32
	_ = v1079
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1089 int32
	_ = v1089
	var v1105 int32
	_ = v1105
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1114 int32
	_ = v1114
	var v1129 int32
	_ = v1129
	var v1135 int32
	_ = v1135
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1146 int32
	_ = v1146
	var v1161 int32
	_ = v1161
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1194 int32
	_ = v1194
	var v1197 int32
	_ = v1197
	var v1203 int32
	_ = v1203
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1210 int32
	_ = v1210
	var v1212 int32
	_ = v1212
	var v1215 int32
	_ = v1215
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1221 int32
	_ = v1221
	var v1230 int64
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1239 int32
	_ = v1239
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1249 int32
	_ = v1249
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1258 int32
	_ = v1258
	var v1260 int32
	_ = v1260
	var v1262 int32
	_ = v1262
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1269 int32
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1278 int32
	_ = v1278
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1313 int32
	_ = v1313
	var v1316 int32
	_ = v1316
	var v1320 int32
	_ = v1320
	var v1325 int32
	_ = v1325
	var v1329 int32
	_ = v1329
	var v1332 int32
	_ = v1332
	var v1336 int32
	_ = v1336
	var v1341 int32
	_ = v1341
	var v1345 int32
	_ = v1345
	var v1351 int32
	_ = v1351
	var v1356 int32
	_ = v1356
	var v1360 int32
	_ = v1360
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1367 int32
	_ = v1367
	var v1372 int32
	_ = v1372
	var v1381 int32
	_ = v1381
	var v1385 int32
	_ = v1385
	var v1390 int32
	_ = v1390
	v17 = m.G0
	v19 = v17 - int32(32)
	m.G0 = v19
	v22 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[0])) = uint8(v22)
	F_BackgroundWorkerUnblockSignals(m)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[1]))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+1336))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[2])) = v29
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[3]))
	v38 = F_AllocSetContextCreateInternal(m, v33, int32(_a_F_ParallelWorkerMain_0), int32(0), int32(_a_F_ParallelWorkerMain_1), int32(_a_F_ParallelWorkerMain_2))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[4])) = v38
	v42 = F_dsm_attach(m, base.I32_wrap_i64(l0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L8
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1381 = m.ExcPending
	if v1381 != 0 {
		goto L1
	} else {
		goto L330
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1360 = m.ExcPending
	if v1360 != 0 {
		goto L1
	} else {
		goto L326
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1345 = m.ExcPending
	if v1345 != 0 {
		goto L1
	} else {
		goto L323
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1329 = m.ExcPending
	if v1329 != 0 {
		goto L1
	} else {
		goto L319
	}
L8:
	;
	if v42 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+24))
	v47 = *(*int64)(unsafe.Add(mBase, uint32(v45)))
	if v47 == int64(1346862204) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	goto L11
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1313 = m.ExcPending
	if v1313 != 0 {
		goto L1
	} else {
		goto L315
	}
L12:
	;
	if v49 == int32(0) {
		goto L7
	} else {
		goto L16
	}
L13:
	;
	v49 = v45
	goto L15
L14:
	;
	v49 = int32(0)
	goto L15
L15:
	;
	goto L12
L16:
	;
	v55 = F_shm_toc_lookup(m, v49, int64(-65535), int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[5])) = v55
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)+40))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[6])) = v59
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v55)+44))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[7])) = v62
	F_before_shmem_exit(m, int32(313), base.I64_extend_i32_u(v42))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v70 = F_shm_toc_lookup(m, v49, int64(-65534), int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[2]))
	v76 = v70 + v73<<(uint(int32(14))%32)
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[8]))
	F_shm_mq_set_sender(m, v76, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v81 = F_shm_mq_attach(m, v76, v42)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_pq_redirect_to_shm_mq(m, v42, v81)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v55)+40))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v55)+44))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[9])) = v86
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[10])) = v85
	goto L23
L23:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v55)+36))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v55)+40))
	v93 = F_BecomeLockGroupMember(m, v91, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if v93 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v95 = *(*int64)(unsafe.Add(mBase, uint32(v55)+48))
	v97 = *(*int64)(unsafe.Add(mBase, uint32(v55)+56))
	*(*int64)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[11])) = v97
	*(*int64)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[12])) = v95
	v103 = F_shm_toc_lookup(m, v49, int64(-65527), int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	m.G0 = v19 + int32(32)
	return
L28:
	;
	v105 = F_strlen(m, v103)
	mBase = m.M
	v108 = v105 + v103 + int32(1)
	v109 = int32(_a_F_ParallelWorkerMain_3)
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	v115 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[13])))
	if base.B2i32(v112 == int32(0))|base.B2i32(v112 != v115) != 0 {
		v133 = v112
		v134 = v115
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[14])) = v297
	v300 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v300)+24)) = v297
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+32)))
	F_SetSessionAuthorization(m, v302, v303)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L89
	}
L30:
	;
	if v133-v134 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L31:
	;
	goto L30
L32:
	;
	v118 = v103
	v119 = v109
	goto L33
L33:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+1)))
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+1)))
	if v123 == int32(0) {
		v133 = v123
		v134 = v122
		goto L31
	} else {
		goto L35
	}
L34:
	;
	v133 = v123
	v134 = v122
	goto L31
L35:
	;
	v126 = int32(1)
	if v123 == v122 {
		v118 = v118 + v126
		v119 = v119 + v126
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v138 = int32(_a_F_ParallelWorkerMain_4)
	v141 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[15])))
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	if base.B2i32(v141 == int32(0))|base.B2i32(v141 != v144) != 0 {
		v162 = v141
		v163 = v144
		goto L41
	} else {
		goto L42
	}
L38:
	;
	goto L39
L39:
	;
	v293 = F_load_external_function(m, v103, v108, int32(1), int32(0))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L88
	}
L40:
	;
	if v162-v163 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L41:
	;
	goto L40
L42:
	;
	v147 = v138
	v148 = v108
	goto L43
L43:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+1)))
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+1)))
	if v152 == int32(0) {
		v162 = v152
		v163 = v151
		goto L41
	} else {
		goto L45
	}
L44:
	;
	v162 = v152
	v163 = v151
	goto L41
L45:
	;
	v155 = int32(1)
	if v152 == v151 {
		v147 = v147 + v155
		v148 = v148 + v155
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[16]))
	v295 = v168
	goto L29
L48:
	;
	goto L49
L49:
	;
	v169 = int32(_a_F_ParallelWorkerMain_5)
	v172 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[17])))
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	if base.B2i32(v172 == int32(0))|base.B2i32(v172 != v175) != 0 {
		v193 = v172
		v194 = v175
		goto L51
	} else {
		goto L52
	}
L50:
	;
	if v193-v194 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L51:
	;
	goto L50
L52:
	;
	v178 = v169
	v179 = v108
	goto L53
L53:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+1)))
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
	if v183 == int32(0) {
		v193 = v183
		v194 = v182
		goto L51
	} else {
		goto L55
	}
L54:
	;
	v193 = v183
	v194 = v182
	goto L51
L55:
	;
	v186 = int32(1)
	if v183 == v182 {
		v178 = v178 + v186
		v179 = v179 + v186
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v199 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[18]))
	v295 = v199
	goto L29
L58:
	;
	goto L59
L59:
	;
	v200 = int32(_a_F_ParallelWorkerMain_6)
	v203 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[19])))
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	if base.B2i32(v203 == int32(0))|base.B2i32(v203 != v206) != 0 {
		v224 = v203
		v225 = v206
		goto L61
	} else {
		goto L62
	}
L60:
	;
	if v224-v225 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L61:
	;
	goto L60
L62:
	;
	v209 = v200
	v210 = v108
	goto L63
L63:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+1)))
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+1)))
	if v214 == int32(0) {
		v224 = v214
		v225 = v213
		goto L61
	} else {
		goto L65
	}
L64:
	;
	v224 = v214
	v225 = v213
	goto L61
L65:
	;
	v217 = int32(1)
	if v214 == v213 {
		v209 = v209 + v217
		v210 = v210 + v217
		goto L63
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	v230 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[20]))
	v295 = v230
	goto L29
L68:
	;
	goto L69
L69:
	;
	v231 = int32(_a_F_ParallelWorkerMain_7)
	v234 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[21])))
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	if base.B2i32(v234 == int32(0))|base.B2i32(v234 != v237) != 0 {
		v255 = v234
		v256 = v237
		goto L71
	} else {
		goto L72
	}
L70:
	;
	if v255-v256 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L71:
	;
	goto L70
L72:
	;
	v240 = v231
	v241 = v108
	goto L73
L73:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241)+1)))
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240)+1)))
	if v245 == int32(0) {
		v255 = v245
		v256 = v244
		goto L71
	} else {
		goto L75
	}
L74:
	;
	v255 = v245
	v256 = v244
	goto L71
L75:
	;
	v248 = int32(1)
	if v245 == v244 {
		v240 = v240 + v248
		v241 = v241 + v248
		goto L73
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	v261 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[22]))
	v295 = v261
	goto L29
L78:
	;
	goto L79
L79:
	;
	v262 = int32(_a_F_ParallelWorkerMain_8)
	v265 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[23])))
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	if base.B2i32(v265 == int32(0))|base.B2i32(v265 != v268) != 0 {
		v286 = v265
		v287 = v268
		goto L81
	} else {
		goto L82
	}
L80:
	;
	if v286-v287 != 0 {
		goto L6
	} else {
		goto L87
	}
L81:
	;
	goto L80
L82:
	;
	v271 = v262
	v272 = v108
	goto L83
L83:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+1)))
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271)+1)))
	if v276 == int32(0) {
		v286 = v276
		v287 = v275
		goto L81
	} else {
		goto L85
	}
L84:
	;
	v286 = v276
	v287 = v275
	goto L81
L85:
	;
	v279 = int32(1)
	if v276 == v275 {
		v271 = v271 + v279
		v272 = v272 + v279
		goto L83
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	v290 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[24]))
	v295 = v290
	goto L29
L88:
	;
	v295 = v293
	goto L29
L89:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v55)+12))
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+33)))
	F_SetCurrentRoleId(m, v306, v307)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	F_BackgroundWorkerInitializeConnectionByOid(m, v310, v311, int32(3))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	v316 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[25]))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v316)+4))
	goto L92
L92:
	;
	v318 = F_SetClientEncoding(m, v317)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	if v318 < int32(0) {
		goto L5
	} else {
		goto L94
	}
L94:
	;
	v324 = F_shm_toc_lookup(m, v49, int64(-65533), int32(0))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v324))))
	if v328 != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v330 = v324
	goto L100
L98:
	;
	goto L99
L99:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L104
	}
L100:
	;
	v345 = F_internal_load_library(m, v330)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L1
	} else {
		goto L102
	}
L101:
	;
	goto L99
L102:
	;
	v347 = F_strlen(m, v330)
	mBase = m.M
	v350 = v347 + v330 + int32(1)
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350))))
	if v351 != 0 {
		v330 = v350
		goto L100
	} else {
		goto L103
	}
L103:
	;
	goto L101
L104:
	;
	v372 = F_shm_toc_lookup(m, v49, int64(-65528), int32(0))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	F_StartTransaction(m)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[26])) = v377
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372)+4)))
	*(*uint8)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[27])) = uint8(v380)
	v383 = *(*int64)(unsafe.Add(mBase, uint32(v372)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[28])) = v383
	v386 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[29]))
	v387 = *(*int64)(unsafe.Add(mBase, uint32(v372)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v386))) = v387
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v372)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[30])) = v390
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v372)+28))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[31])) = v372 + int32(32)
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[32])) = v392
	*(*int32)(unsafe.Add(mBase, uint32(v386)+24)) = int32(5)
	v403 = F_shm_toc_lookup(m, v49, int64(-65525), int32(0))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	v405 = m.G0
	v407 = v405 - int32(48)
	m.G0 = v407
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v403)+8))
	if v409 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v411 = v403
	goto L111
L109:
	;
	goto L110
L110:
	;
	m.G0 = v407 + int32(48)
	v472 = F_shm_toc_lookup(m, v49, int64(-65523), int32(0))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L1
	} else {
		goto L119
	}
L111:
	;
	v427 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[33]))
	if v427 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	goto L110
L113:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v407)+8)) = int64(68719476748)
	v433 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[34]))
	*(*int32)(unsafe.Add(mBase, uint32(v407)+36)) = v433
	v439 = F_hash_create(m, int32(_a_F_ParallelWorkerMain_9), int64(16), v407, int32(1064))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L1
	} else {
		goto L116
	}
L114:
	;
	v442 = v427
	goto L115
L115:
	;
	v444 = F_hash_search(m, v442, v411, int32(1), v407)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L117
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[33])) = v439
	v442 = v439
	goto L115
L117:
	;
	v446 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v444)+12)) = uint8(v446)
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v411)+20))
	if v448 != 0 {
		v411 = v411 + int32(12)
		goto L111
	} else {
		goto L118
	}
L118:
	;
	goto L112
L119:
	;
	v475 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[35]))
	if v475 != 0 {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v498 = int32(524)
	base.MemoryCopy(m, int32(_a_F_ParallelWorkerMain_10), v472, v498)
	base.MemoryCopy(m, int32(_a_F_ParallelWorkerMain_11), v472+v498, v498)
	v507 = F_shm_toc_lookup(m, v49, int64(-65524), int32(0))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L1
	} else {
		goto L129
	}
L121:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L1
	} else {
		goto L126
	}
L122:
	;
	v477 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[36]))
	if v477 != 0 {
		goto L121
	} else {
		goto L123
	}
L123:
	;
	v479 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[37]))
	if v479 != 0 {
		goto L121
	} else {
		goto L124
	}
L124:
	;
	v481 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[38]))
	if v481 == int32(0) {
		goto L120
	} else {
		goto L125
	}
L125:
	;
	goto L121
L126:
	;
	F_errmsg_internal(m, int32(_a_F_ParallelWorkerMain_12), int32(0))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(_a_F_ParallelWorkerMain_13), int32(750), int32(_a_F_ParallelWorkerMain_14))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L129:
	;
	v509 = int32(0)
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v507)))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[39])) = v511
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v507)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[40])) = v514
	v516 = int32(_a_F_ParallelWorkerMain_15)
	v517 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[4]))
	v520 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[4])) = v520
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v507)+8))
	if v509 < v522 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v528 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[41]))
	v531 = v509
	v532 = v528
	goto L133
L131:
	;
	goto L132
L132:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[4])) = v517
	v577 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[29]))
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v577)+28))
	goto L137
L133:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v507+int32(12)+v531<<(uint(int32(2))%32))))
	v550 = F_lappend_oid(m, v532, v549)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L1
	} else {
		goto L135
	}
L134:
	;
	goto L132
L135:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[41])) = v550
	v554 = v531 + int32(1)
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v507)+8))
	if v554 < v555 {
		v531 = v554
		v532 = v550
		goto L133
	} else {
		goto L136
	}
L136:
	;
	goto L134
L137:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[42])) = v578
	v582 = F_shm_toc_lookup(m, v49, int64(-65531), int32(0))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	v584 = int32(0)
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v582)))
	if v585 <= v584 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v648 = F_shm_toc_lookup(m, v49, int64(-65526), int32(0))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L1
	} else {
		goto L151
	}
L140:
	;
	v592 = v584
	goto L141
L141:
	;
	v608 = v582 + int32(4) + v592<<(uint(int32(3))%32)
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v608)))
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v608)+4))
	v611 = F_GetComboCommandId(m, v609, v610)
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L1
	} else {
		goto L143
	}
L142:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L1
	} else {
		goto L148
	}
L143:
	;
	if v611 == v592 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v615 = v592 + int32(1)
	if v585 != v615 {
		v592 = v615
		goto L141
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	goto L142
L147:
	;
	goto L139
L148:
	;
	F_errmsg_internal(m, int32(_a_F_ParallelWorkerMain_16), int32(0))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	F_errfinish(m, int32(_a_F_ParallelWorkerMain_17), int32(362), int32(_a_F_ParallelWorkerMain_18))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L151:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v648)))
	v651 = int32(_a_F_ParallelWorkerMain_15)
	v652 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[4]))
	v655 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[4])) = v655
	v657 = F_dsm_attach(m, v650)
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	if v657 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L1
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v657)+24))
	v677 = *(*int64)(unsafe.Add(mBase, uint32(v675)))
	if v677 == int64(2880502729) {
		goto L160
	} else {
		goto L161
	}
L156:
	;
	F_errmsg_internal(m, int32(_a_F_ParallelWorkerMain_19), int32(0))
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	F_errfinish(m, int32(_a_F_ParallelWorkerMain_20), int32(169), int32(_a_F_ParallelWorkerMain_21))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L159:
	;
	v682 = F_shm_toc_lookup(m, v679, int64(-65535), int32(0))
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L1
	} else {
		goto L163
	}
L160:
	;
	v679 = v675
	goto L162
L161:
	;
	v679 = int32(0)
	goto L162
L162:
	;
	goto L159
L163:
	;
	v684 = F_dsa_attach_in_place(m, v682, v657)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	v686 = int32(_a_F_ParallelWorkerMain_22)
	v687 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[43]))
	*(*int32)(unsafe.Add(mBase, uint32(v687))) = v657
	v690 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[43]))
	*(*int32)(unsafe.Add(mBase, uint32(v690)+4)) = v684
	v694 = F_shm_toc_lookup(m, v679, int64(-65534), int32(0))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	v696 = int32(_a_F_ParallelWorkerMain_15)
	v697 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[4]))
	v700 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[4])) = v700
	v703 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[43]))
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v703)+4))
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v694)))
	v707 = F_dshash_attach(m, v704, int32(_a_F_ParallelWorkerMain_23), v706, v704)
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	v710 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[43]))
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v710)+4))
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v694)+4))
	v715 = F_dshash_attach(m, v711, int32(_a_F_ParallelWorkerMain_24), v713, int32(0))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[4])) = v697
	v720 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[43]))
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v720)))
	F_on_dsm_detach(m, v721, int32(1821), base.I64_extend_i32_u(v694))
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	v727 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[43]))
	*(*int32)(unsafe.Add(mBase, uint32(v727)+16)) = v715
	*(*int32)(unsafe.Add(mBase, uint32(v727)+12)) = v707
	*(*int32)(unsafe.Add(mBase, uint32(v727)+8)) = v694
	F_dsm_pin_mapping(m, v657)
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	F_dsa_pin_mapping(m, v684)
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[4])) = v652
	v739 = F_shm_toc_lookup(m, v49, int64(-65529), int32(0))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	v743 = F_shm_toc_lookup(m, v49, int64(-65530), int32(1))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	v745 = F_RestoreSnapshot(m, v739)
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	if v743 != 0 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v747 = F_RestoreSnapshot(m, v743)
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L1
	} else {
		goto L177
	}
L175:
	;
	v749 = v745
	goto L176
L176:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v55)+36))
	F_RestoreTransactionSnapshot(m, v749, v750)
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L1
	} else {
		goto L178
	}
L177:
	;
	v749 = v747
	goto L176
L178:
	;
	F_PushActiveSnapshot(m, v745)
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L1
	} else {
		goto L179
	}
L179:
	;
	F_InvalidateSystemCachesExtended(m)
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	v759 = F_shm_toc_lookup(m, v49, int64(-65532), int32(0))
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L1
	} else {
		goto L181
	}
L181:
	;
	v761 = m.G0
	v763 = v761 - int32(32)
	m.G0 = v763
	v766 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[44]))
	v767 = int32(0)
	if base.B2i32(v766 == v767)|base.B2i32(v766 == int32(_a_F_ParallelWorkerMain_25)) == v767 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v775 = v766
	goto L185
L183:
	;
	goto L184
L184:
	;
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v759)))
	*(*int32)(unsafe.Add(mBase, uint32(v763)+20)) = int32(1858)
	v908 = int32(_a_F_ParallelWorkerMain_26)
	v909 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[45]))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[45])) = v763 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v763)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v763)+16)) = v909
	v918 = v759 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v763)+28)) = v918
	if v905 != 0 {
		goto L240
	} else {
		goto L241
	}
L185:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v775)+4))
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v775+int32(-64))))
	if base.Ui32(v793) < base.Ui32(int32(2)) {
		goto L187
	} else {
		goto L188
	}
L186:
	;
	goto L184
L187:
	;
	if v790 != int32(_a_F_ParallelWorkerMain_25) {
		v775 = v790
		goto L185
	} else {
		goto L237
	}
L188:
	;
	v797 = v775 - int32(36)
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v797)))
	if v798 == int32(0) {
		goto L187
	} else {
		goto L189
	}
L189:
	;
	v802 = v775 - int32(8)
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v802)))
	if v803 != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	F_pfree(m, v803)
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L1
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v775)+16))
	if v806 != 0 {
		goto L194
	} else {
		goto L195
	}
L193:
	;
	goto L192
L194:
	;
	F_pfree(m, v806)
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L1
	} else {
		goto L197
	}
L195:
	;
	goto L196
L196:
	;
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v775)+20))
	if v809 != 0 {
		goto L198
	} else {
		goto L199
	}
L197:
	;
	goto L196
L198:
	;
	F_pfree(m, v809)
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L1
	} else {
		goto L201
	}
L199:
	;
	goto L200
L200:
	;
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v775-int32(44))))
	if v814 != int32(3) {
		goto L202
	} else {
		goto L203
	}
L201:
	;
	goto L200
L202:
	;
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v775-int32(4))))
	if v832 == int32(0) {
		goto L211
	} else {
		goto L212
	}
L203:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v775)+28))
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v817)))
	if v818 != 0 {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	F_pfree(m, v818)
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L1
	} else {
		goto L207
	}
L205:
	;
	goto L206
L206:
	;
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v775)+48))
	if v821 == int32(0) {
		goto L202
	} else {
		goto L208
	}
L207:
	;
	goto L206
L208:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v775)+28))
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v824)))
	if v821 == v825 {
		goto L202
	} else {
		goto L209
	}
L209:
	;
	F_pfree(m, v821)
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L1
	} else {
		goto L210
	}
L210:
	;
	goto L202
L211:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v797)))
	if v839 != 0 {
		goto L215
	} else {
		goto L216
	}
L212:
	;
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v802)))
	if v832 == v835 {
		goto L211
	} else {
		goto L213
	}
L213:
	;
	F_pfree(m, v832)
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L1
	} else {
		goto L214
	}
L214:
	;
	goto L211
L215:
	;
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v775)))
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v775)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v840)+4)) = v841
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v775)))
	*(*int32)(unsafe.Add(mBase, uint32(v841))) = v843
	goto L217
L216:
	;
	goto L217
L217:
	;
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v775-int32(12))))
	if v848 != 0 {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v853 = int32(_a_F_ParallelWorkerMain_27)
	goto L223
L219:
	;
	goto L220
L220:
	;
	v864 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v775-int32(40)))))
	if v864&int32(4) != 0 {
		goto L227
	} else {
		goto L228
	}
L221:
	;
	goto L220
L222:
	;
	goto L221
L223:
	;
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v853)))
	if v856 == int32(0) {
		goto L222
	} else {
		goto L225
	}
L224:
	;
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v856)))
	*(*int32)(unsafe.Add(mBase, uint32(v853))) = v860
	goto L222
L225:
	;
	if v856 != v775+int32(8) {
		v853 = v856
		goto L223
	} else {
		goto L226
	}
L226:
	;
	goto L224
L227:
	;
	v871 = int32(_a_F_ParallelWorkerMain_28)
	goto L232
L228:
	;
	goto L229
L229:
	;
	F_InitializeOneGUCOption(m, v775-int32(68))
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L1
	} else {
		goto L236
	}
L230:
	;
	goto L229
L231:
	;
	goto L230
L232:
	;
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v871)))
	if v874 == int32(0) {
		goto L231
	} else {
		goto L234
	}
L233:
	;
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v874)))
	*(*int32)(unsafe.Add(mBase, uint32(v871))) = v878
	goto L231
L234:
	;
	if v874 != v775+int32(12) {
		v871 = v874
		goto L232
	} else {
		goto L235
	}
L235:
	;
	goto L233
L236:
	;
	goto L187
L237:
	;
	goto L186
L238:
	;
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v55)+28))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[46])) = v1050
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[47])) = v1049
	goto L275
L239:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1036 = m.ExcPending
	if v1036 != 0 {
		goto L1
	} else {
		goto L271
	}
L240:
	;
	v921 = v905 + v918
	goto L243
L241:
	;
	v1028 = v909
	goto L242
L242:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[45])) = v1028
	m.G0 = v763 + int32(32)
	goto L238
L243:
	;
	v939 = v763 + int32(28)
	v940 = F_read_gucstate(m, v939, v921)
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L1
	} else {
		goto L245
	}
L244:
	;
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(v763)+16))
	v1028 = v1011
	goto L242
L245:
	;
	v942 = F_read_gucstate(m, v939, v921)
	mBase = m.M
	v943 = m.ExcPending
	if v943 != 0 {
		goto L1
	} else {
		goto L246
	}
L246:
	;
	v944 = F_read_gucstate(m, v939, v921)
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L1
	} else {
		goto L247
	}
L247:
	;
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v763)+28))
	v947 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v944))))
	if v947 == int32(0) {
		goto L249
	} else {
		goto L250
	}
L248:
	;
	v958 = v955 + int32(4)
	if base.Ui32(v921) < base.Ui32(v958) {
		goto L4
	} else {
		goto L253
	}
L249:
	;
	v955 = v946
	v956 = int32(0)
	goto L248
L250:
	;
	goto L251
L251:
	;
	v952 = v946 + int32(4)
	if base.Ui32(v921) < base.Ui32(v952) {
		goto L4
	} else {
		goto L252
	}
L252:
	;
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v946)))
	v955 = v952
	v956 = v954
	goto L248
L253:
	;
	v961 = v955 + int32(8)
	if base.Ui32(v921) < base.Ui32(v961) {
		goto L4
	} else {
		goto L254
	}
L254:
	;
	v964 = v955 + int32(12)
	if base.Ui32(v921) < base.Ui32(v964) {
		goto L4
	} else {
		goto L255
	}
L255:
	;
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v955)))
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v961)))
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v958)))
	*(*int32)(unsafe.Add(mBase, uint32(v763)+12)) = v942
	*(*int32)(unsafe.Add(mBase, uint32(v763)+8)) = v940
	*(*int32)(unsafe.Add(mBase, uint32(v763)+28)) = v964
	*(*int32)(unsafe.Add(mBase, uint32(v763)+24)) = v763 + int32(8)
	v975 = int32(0)
	v977 = int32(1)
	v980 = F_set_config_with_handle(m, v940, v975, v942, v968, v966, v967, v975, v977, int32(21), v977)
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L1
	} else {
		goto L256
	}
L256:
	;
	if v980 <= int32(0) {
		goto L239
	} else {
		goto L257
	}
L257:
	;
	v984 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v944))))
	if v984 == int32(0) {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v763)+24)) = int32(0)
	if base.Ui32(v964) < base.Ui32(v921) {
		goto L243
	} else {
		goto L270
	}
L259:
	;
	v992 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[48])))
	if v992 != 0 {
		goto L260
	} else {
		goto L261
	}
L260:
	;
	v993 = int32(12)
	goto L262
L261:
	;
	v993 = int32(15)
	goto L262
L262:
	;
	v994 = F_find_option(m, v940, int32(1), int32(0), v993)
	mBase = m.M
	v995 = m.ExcPending
	if v995 != 0 {
		goto L1
	} else {
		goto L263
	}
L263:
	;
	if v994 == int32(0) {
		goto L258
	} else {
		goto L264
	}
L264:
	;
	v998 = F_guc_strdup(m, v993, v944)
	mBase = m.M
	v999 = m.ExcPending
	if v999 != 0 {
		goto L1
	} else {
		goto L265
	}
L265:
	;
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(v994)+88))
	if v1000 != 0 {
		goto L266
	} else {
		goto L267
	}
L266:
	;
	F_pfree(m, v1000)
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L1
	} else {
		goto L269
	}
L267:
	;
	goto L268
L268:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v994)+92)) = v956
	*(*int32)(unsafe.Add(mBase, uint32(v994)+88)) = v998
	goto L258
L269:
	;
	goto L268
L270:
	;
	goto L244
L271:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v1039 = m.ExcPending
	if v1039 != 0 {
		goto L1
	} else {
		goto L272
	}
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v763))) = v940
	F_errmsg(m, int32(_a_F_ParallelWorkerMain_29), v763)
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L1
	} else {
		goto L273
	}
L273:
	;
	F_errfinish(m, int32(_a_F_ParallelWorkerMain_30), int32(_a_F_ParallelWorkerMain_31), int32(_a_F_ParallelWorkerMain_32))
	mBase = m.M
	v1048 = m.ExcPending
	if v1048 != 0 {
		goto L1
	} else {
		goto L274
	}
L274:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L275:
	;
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v55)+20))
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(v55)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[49])) = v1057
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[50])) = v1055
	v1062 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[51])) = uint8(v1062)
	v1065 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[52])) = uint8(v1065)
	v1069 = F_shm_toc_lookup(m, v49, int64(-65522), v1065)
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L1
	} else {
		goto L276
	}
L276:
	;
	v1071 = m.G0
	v1073 = v1071 - int32(48)
	m.G0 = v1073
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(v1069)))
	if v1075 != 0 {
		goto L277
	} else {
		goto L278
	}
L277:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1073)+8)) = int64(17179869188)
	v1079 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[34]))
	*(*int32)(unsafe.Add(mBase, uint32(v1073)+36)) = v1079
	v1085 = F_hash_create(m, int32(_a_F_ParallelWorkerMain_33), int64(32), v1073, int32(1064))
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		goto L1
	} else {
		goto L280
	}
L278:
	;
	v1114 = v1069
	goto L279
L279:
	;
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v1114)+4))
	if v1129 != 0 {
		goto L285
	} else {
		goto L286
	}
L280:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[53])) = v1085
	v1089 = v1069
	goto L281
L281:
	;
	v1105 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[53]))
	v1108 = F_hash_search(m, v1105, v1089, int32(1), int32(0))
	mBase = m.M
	v1109 = m.ExcPending
	if v1109 != 0 {
		goto L1
	} else {
		goto L283
	}
L282:
	;
	v1114 = v1111
	goto L279
L283:
	;
	v1111 = v1089 + int32(4)
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v1089)+4))
	if v1112 != 0 {
		v1089 = v1111
		goto L281
	} else {
		goto L284
	}
L284:
	;
	goto L282
L285:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1073)+8)) = int64(17179869188)
	v1135 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[34]))
	*(*int32)(unsafe.Add(mBase, uint32(v1073)+36)) = v1135
	v1141 = F_hash_create(m, int32(_a_F_ParallelWorkerMain_34), int64(32), v1073, int32(1064))
	mBase = m.M
	v1142 = m.ExcPending
	if v1142 != 0 {
		goto L1
	} else {
		goto L288
	}
L286:
	;
	goto L287
L287:
	;
	m.G0 = v1073 + int32(48)
	v1190 = F_shm_toc_lookup(m, v49, int64(-65521), int32(0))
	mBase = m.M
	v1191 = m.ExcPending
	if v1191 != 0 {
		goto L1
	} else {
		goto L293
	}
L288:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[54])) = v1141
	v1146 = v1114 + int32(4)
	goto L289
L289:
	;
	v1161 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[54]))
	v1164 = F_hash_search(m, v1161, v1146, int32(1), int32(0))
	mBase = m.M
	v1165 = m.ExcPending
	if v1165 != 0 {
		goto L1
	} else {
		goto L291
	}
L290:
	;
	goto L287
L291:
	;
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(v1146)+4))
	if v1166 != 0 {
		v1146 = v1146 + int32(4)
		goto L289
	} else {
		goto L292
	}
L292:
	;
	goto L290
L293:
	;
	v1192 = *(*int32)(unsafe.Add(mBase, uint32(v1190)))
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(v1190)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[55])) = v1194
	v1197 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[56])) = v1197
	if v1197 <= v1192 {
		goto L294
	} else {
		goto L295
	}
L294:
	;
	v1203 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[3]))
	v1206 = F_MemoryContextStrdup(m, v1203, v1190+int32(8))
	mBase = m.M
	v1207 = m.ExcPending
	if v1207 != 0 {
		goto L1
	} else {
		goto L297
	}
L295:
	;
	goto L296
L296:
	;
	v1210 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[56]))
	if v1210 != 0 {
		goto L298
	} else {
		goto L299
	}
L297:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[56])) = v1206
	goto L296
L298:
	;
	v1212 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[55]))
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(v1212<<(uint(int32(2))%32))+uint32(_c_F_ParallelWorkerMain[57])))
	goto L301
L299:
	;
	goto L300
L300:
	;
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(v55)+64))
	v1219 = m.G0
	v1221 = v1219 - int32(48)
	m.G0 = v1221
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[58])) = v1218
	if v1218 != 0 {
		goto L303
	} else {
		goto L304
	}
L301:
	;
	F_InitializeSystemUser(m, v1210, v1215)
	mBase = m.M
	v1217 = m.ExcPending
	if v1217 != 0 {
		goto L1
	} else {
		goto L302
	}
L302:
	;
	goto L300
L303:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1221)+8)) = int64(103079215120)
	v1230 = int64(*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[59])))
	v1232 = F_hash_create(m, int32(_a_F_ParallelWorkerMain_35), v1230, v1221, int32(40))
	mBase = m.M
	v1233 = m.ExcPending
	if v1233 != 0 {
		goto L1
	} else {
		goto L306
	}
L304:
	;
	goto L305
L305:
	;
	m.G0 = v1221 + int32(48)
	v1239 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[0])) = uint8(v1239)
	v1243 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[29]))
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(v1243)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v1243)+72)) = v1244 + int32(1)
	goto L307
L306:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[60])) = v1232
	goto L305
L307:
	;
	m.T0[v295].(func(*base.Module, int32, int32))(m, v42, v49)
	mBase = m.M
	v1249 = m.ExcPending
	if v1249 != 0 {
		goto L1
	} else {
		goto L308
	}
L308:
	;
	v1252 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[29]))
	v1253 = *(*int32)(unsafe.Add(mBase, uint32(v1252)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v1252)+72)) = v1253 - int32(1)
	goto L309
L309:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v1258 = m.ExcPending
	if v1258 != 0 {
		goto L1
	} else {
		goto L310
	}
L310:
	;
	F_CommitTransaction(m)
	mBase = m.M
	v1260 = m.ExcPending
	if v1260 != 0 {
		goto L1
	} else {
		goto L311
	}
L311:
	;
	v1262 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[29]))
	*(*int32)(unsafe.Add(mBase, uint32(v1262)+24)) = int32(0)
	v1266 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[43]))
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(v1266)))
	F_dsm_detach(m, v1267)
	mBase = m.M
	v1269 = m.ExcPending
	if v1269 != 0 {
		goto L1
	} else {
		goto L312
	}
L312:
	;
	v1270 = int32(_a_F_ParallelWorkerMain_22)
	v1271 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[43]))
	*(*int32)(unsafe.Add(mBase, uint32(v1271))) = int32(0)
	v1275 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[43]))
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(v1275)+4))
	F_dsa_detach(m, v1276)
	mBase = m.M
	v1278 = m.ExcPending
	if v1278 != 0 {
		goto L1
	} else {
		goto L313
	}
L313:
	;
	v1280 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[43]))
	v1281 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1280)+4)) = v1281
	v1287 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[61]))
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(v1287)+16))
	v1289 = m.T0[v1288].(func(*base.Module, int32, int32, int32) int32)(m, int32(88), v1281, v1281)
	mBase = m.M
	v1290 = m.ExcPending
	if v1290 != 0 {
		goto L1
	} else {
		goto L314
	}
L314:
	;
	goto L27
L315:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1316 = m.ExcPending
	if v1316 != 0 {
		goto L1
	} else {
		goto L316
	}
L316:
	;
	F_errmsg(m, int32(_a_F_ParallelWorkerMain_36), int32(0))
	mBase = m.M
	v1320 = m.ExcPending
	if v1320 != 0 {
		goto L1
	} else {
		goto L317
	}
L317:
	;
	F_errfinish(m, int32(_a_F_ParallelWorkerMain_37), int32(1356), int32(_a_F_ParallelWorkerMain_38))
	mBase = m.M
	v1325 = m.ExcPending
	if v1325 != 0 {
		goto L1
	} else {
		goto L318
	}
L318:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L319:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1332 = m.ExcPending
	if v1332 != 0 {
		goto L1
	} else {
		goto L320
	}
L320:
	;
	F_errmsg(m, int32(_a_F_ParallelWorkerMain_39), int32(0))
	mBase = m.M
	v1336 = m.ExcPending
	if v1336 != 0 {
		goto L1
	} else {
		goto L321
	}
L321:
	;
	F_errfinish(m, int32(_a_F_ParallelWorkerMain_37), int32(1361), int32(_a_F_ParallelWorkerMain_38))
	mBase = m.M
	v1341 = m.ExcPending
	if v1341 != 0 {
		goto L1
	} else {
		goto L322
	}
L322:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L323:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v108
	F_errmsg_internal(m, int32(_a_F_ParallelWorkerMain_40), v19+int32(16))
	mBase = m.M
	v1351 = m.ExcPending
	if v1351 != 0 {
		goto L1
	} else {
		goto L324
	}
L324:
	;
	F_errfinish(m, int32(_a_F_ParallelWorkerMain_37), int32(1667), int32(_a_F_ParallelWorkerMain_41))
	mBase = m.M
	v1356 = m.ExcPending
	if v1356 != 0 {
		goto L1
	} else {
		goto L325
	}
L325:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L326:
	;
	v1362 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[25]))
	v1363 = *(*int32)(unsafe.Add(mBase, uint32(v1362)+4))
	goto L327
L327:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v1363
	F_errmsg_internal(m, int32(_a_F_ParallelWorkerMain_42), v19)
	mBase = m.M
	v1367 = m.ExcPending
	if v1367 != 0 {
		goto L1
	} else {
		goto L328
	}
L328:
	;
	F_errfinish(m, int32(_a_F_ParallelWorkerMain_37), int32(1453), int32(_a_F_ParallelWorkerMain_38))
	mBase = m.M
	v1372 = m.ExcPending
	if v1372 != 0 {
		goto L1
	} else {
		goto L329
	}
L329:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L330:
	;
	F_errmsg_internal(m, int32(_a_F_ParallelWorkerMain_43), int32(0))
	mBase = m.M
	v1385 = m.ExcPending
	if v1385 != 0 {
		goto L1
	} else {
		goto L331
	}
L331:
	;
	F_errfinish(m, int32(_a_F_ParallelWorkerMain_30), int32(_a_F_ParallelWorkerMain_44), int32(_a_F_ParallelWorkerMain_45))
	mBase = m.M
	v1390 = m.ExcPending
	if v1390 != 0 {
		goto L1
	} else {
		goto L332
	}
L332:
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
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
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
func F_parallel_vacuum_dsm_detach(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_dsm_detach[0])) = int32(0)
	return
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
func F_parallel_vacuum_propagate_shared_delay_params(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 float64
	_ = v7
	var v8 float64
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 float64
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[0]))
	if v3 == int32(0) {
		return
	} else {
		v7 = *(*float64)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[1]))
		v8 = *(*float64)(unsafe.Add(mBase, uint32(v3)+8))
		if base.F64_ne(v7, v8) != 0 {
			v28 = base.AtomicRmwXchg32(m, v3, int32(4), int32(1))
			if v28 != 0 {
				F_s_lock(m, v3+int32(4), int32(_a_F_parallel_vacuum_propagate_shared_delay_params_0))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[0]))
					v37 = *(*float64)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[1]))
					*(*float64)(unsafe.Add(mBase, uint32(v35)+8)) = v37
					v40 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[2]))
					*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v40
					v43 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[3]))
					*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = v43
					v46 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[4]))
					*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = v46
					v49 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[5]))
					*(*int32)(unsafe.Add(mBase, uint32(v35)+28)) = v49
					v51 = int32(0)
					atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v35)+4)), uint32(v51))
					v56 = base.AtomicRmwAdd32(m, v35, v51, int32(1))
					return
				}
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[0]))
				v37 = *(*float64)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[1]))
				*(*float64)(unsafe.Add(mBase, uint32(v35)+8)) = v37
				v40 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[2]))
				*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v40
				v43 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[3]))
				*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = v43
				v46 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[4]))
				*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = v46
				v49 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[5]))
				*(*int32)(unsafe.Add(mBase, uint32(v35)+28)) = v49
				v51 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v35)+4)), uint32(v51))
				v56 = base.AtomicRmwAdd32(m, v35, v51, int32(1))
				return
			}
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[2]))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v3)+16))
			if v11 != v12 {
				v28 = base.AtomicRmwXchg32(m, v3, int32(4), int32(1))
				if v28 != 0 {
					F_s_lock(m, v3+int32(4), int32(_a_F_parallel_vacuum_propagate_shared_delay_params_0))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[0]))
						v37 = *(*float64)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[1]))
						*(*float64)(unsafe.Add(mBase, uint32(v35)+8)) = v37
						v40 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[2]))
						*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v40
						v43 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[3]))
						*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = v43
						v46 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[4]))
						*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = v46
						v49 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[5]))
						*(*int32)(unsafe.Add(mBase, uint32(v35)+28)) = v49
						v51 = int32(0)
						atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v35)+4)), uint32(v51))
						v56 = base.AtomicRmwAdd32(m, v35, v51, int32(1))
						return
					}
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[0]))
					v37 = *(*float64)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[1]))
					*(*float64)(unsafe.Add(mBase, uint32(v35)+8)) = v37
					v40 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[2]))
					*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v40
					v43 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[3]))
					*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = v43
					v46 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[4]))
					*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = v46
					v49 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[5]))
					*(*int32)(unsafe.Add(mBase, uint32(v35)+28)) = v49
					v51 = int32(0)
					atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v35)+4)), uint32(v51))
					v56 = base.AtomicRmwAdd32(m, v35, v51, int32(1))
					return
				}
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[3]))
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v3)+20))
				if v15 != v16 {
					v28 = base.AtomicRmwXchg32(m, v3, int32(4), int32(1))
					if v28 != 0 {
						F_s_lock(m, v3+int32(4), int32(_a_F_parallel_vacuum_propagate_shared_delay_params_0))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[0]))
							v37 = *(*float64)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[1]))
							*(*float64)(unsafe.Add(mBase, uint32(v35)+8)) = v37
							v40 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[2]))
							*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v40
							v43 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[3]))
							*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = v43
							v46 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[4]))
							*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = v46
							v49 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[5]))
							*(*int32)(unsafe.Add(mBase, uint32(v35)+28)) = v49
							v51 = int32(0)
							atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v35)+4)), uint32(v51))
							v56 = base.AtomicRmwAdd32(m, v35, v51, int32(1))
							return
						}
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[0]))
						v37 = *(*float64)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[1]))
						*(*float64)(unsafe.Add(mBase, uint32(v35)+8)) = v37
						v40 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[2]))
						*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v40
						v43 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[3]))
						*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = v43
						v46 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[4]))
						*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = v46
						v49 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[5]))
						*(*int32)(unsafe.Add(mBase, uint32(v35)+28)) = v49
						v51 = int32(0)
						atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v35)+4)), uint32(v51))
						v56 = base.AtomicRmwAdd32(m, v35, v51, int32(1))
						return
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[4]))
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v3)+24))
					if v19 != v20 {
						v28 = base.AtomicRmwXchg32(m, v3, int32(4), int32(1))
						if v28 != 0 {
							F_s_lock(m, v3+int32(4), int32(_a_F_parallel_vacuum_propagate_shared_delay_params_0))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return
							} else {
								v35 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[0]))
								v37 = *(*float64)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[1]))
								*(*float64)(unsafe.Add(mBase, uint32(v35)+8)) = v37
								v40 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[2]))
								*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v40
								v43 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[3]))
								*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = v43
								v46 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[4]))
								*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = v46
								v49 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[5]))
								*(*int32)(unsafe.Add(mBase, uint32(v35)+28)) = v49
								v51 = int32(0)
								atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v35)+4)), uint32(v51))
								v56 = base.AtomicRmwAdd32(m, v35, v51, int32(1))
								return
							}
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[0]))
							v37 = *(*float64)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[1]))
							*(*float64)(unsafe.Add(mBase, uint32(v35)+8)) = v37
							v40 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[2]))
							*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v40
							v43 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[3]))
							*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = v43
							v46 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[4]))
							*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = v46
							v49 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[5]))
							*(*int32)(unsafe.Add(mBase, uint32(v35)+28)) = v49
							v51 = int32(0)
							atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v35)+4)), uint32(v51))
							v56 = base.AtomicRmwAdd32(m, v35, v51, int32(1))
							return
						}
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[5]))
						v24 = *(*int32)(unsafe.Add(mBase, uint32(v3)+28))
						if v23 == v24 {
							return
						} else {
							v28 = base.AtomicRmwXchg32(m, v3, int32(4), int32(1))
							if v28 != 0 {
								F_s_lock(m, v3+int32(4), int32(_a_F_parallel_vacuum_propagate_shared_delay_params_0))
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return
								} else {
									v35 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[0]))
									v37 = *(*float64)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[1]))
									*(*float64)(unsafe.Add(mBase, uint32(v35)+8)) = v37
									v40 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[2]))
									*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v40
									v43 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[3]))
									*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = v43
									v46 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[4]))
									*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = v46
									v49 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[5]))
									*(*int32)(unsafe.Add(mBase, uint32(v35)+28)) = v49
									v51 = int32(0)
									atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v35)+4)), uint32(v51))
									v56 = base.AtomicRmwAdd32(m, v35, v51, int32(1))
									return
								}
							} else {
								v35 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[0]))
								v37 = *(*float64)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[1]))
								*(*float64)(unsafe.Add(mBase, uint32(v35)+8)) = v37
								v40 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[2]))
								*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v40
								v43 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[3]))
								*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = v43
								v46 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[4]))
								*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = v46
								v49 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_propagate_shared_delay_params[5]))
								*(*int32)(unsafe.Add(mBase, uint32(v35)+28)) = v49
								v51 = int32(0)
								atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v35)+4)), uint32(v51))
								v56 = base.AtomicRmwAdd32(m, v35, v51, int32(1))
								return
							}
						}
					}
				}
			}
		}
	}
}
