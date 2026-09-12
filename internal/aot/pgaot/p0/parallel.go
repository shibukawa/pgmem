package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CreateParallelContext(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	v6 = int32(4470752)
	v7 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v10 = *(*int32)(unsafe.Add(mBase, _consts[76]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v10
	v13 = F_palloc0(m, int32(68))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, _consts[37]))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v19
		v24 = F_pstrdup(m, int32(159252))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v24
			v27 = F_pstrdup(m, l0)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v27
				v31 = *(*int32)(unsafe.Add(mBase, _consts[77]))
				*(*int64)(unsafe.Add(mBase, uint32(v13)+36)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v31
				v36 = *(*int32)(unsafe.Add(mBase, _consts[78]))
				if v36 == int32(0) {
					v39 = int32(4080156)
					*(*int32)(unsafe.Add(mBase, _consts[79])) = v39
					v43 = v39
				} else {
					v43 = v36
				}
				*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(4080156)
				*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v43
				*(*int32)(unsafe.Add(mBase, uint32(v43))) = v13
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v7
				*(*int32)(unsafe.Add(mBase, _consts[78])) = v13
				return v13
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
	v31 = *(*int32)(unsafe.Add(mBase, _consts[82]))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v35 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v39 = F_LWLockAcquire(m, v35+int32(4224), int32(0))
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
	v50 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	F_LWLockRelease(m, v50+int32(4224))
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
	v59 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	F_LWLockRelease(m, v59+int32(4224))
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
	v96 = int32(4465404)
	v98 = *(*int32)(unsafe.Add(mBase, _consts[83]))
	*(*int32)(unsafe.Add(mBase, _consts[83])) = v98 + int32(1)
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
	v104 = int32(4465404)
	v106 = *(*int32)(unsafe.Add(mBase, _consts[83]))
	*(*int32)(unsafe.Add(mBase, _consts[83])) = v106 - int32(1)
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v64 int64
	_ = v64
	var v65 int64
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
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
	var v115 int32
	_ = v115
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
	if v12 != 0 {
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
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+184))
	if v15 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v98 != 0 {
		goto L26
	} else {
		goto L27
	}
L9:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+100))
	v20 = F_MemoryContextAllocZero(m, v18, int32(48))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L12
	}
L10:
	;
	v27 = v15
	goto L11
L11:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if int32(0) < v28 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+184)) = v20
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+184))
	v27 = v25
	goto L11
L13:
	;
	v35 = int32(0)
	goto L16
L14:
	;
	v76 = v28
	goto L15
L15:
	;
	v80 = F_mul_size(m, v76, int32(48))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L4
	} else {
		goto L20
	}
L16:
	;
	v43 = v12 + int32(8) + v35*int32(48)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v44 + v45
	v48 = *(*int64)(unsafe.Add(mBase, uint32(v27)+8))
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v43)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+8)) = v48 + v49
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v27)+16))
	v53 = *(*int64)(unsafe.Add(mBase, uint32(v43)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+16)) = v52 + v53
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v27)+24))
	v57 = *(*int64)(unsafe.Add(mBase, uint32(v43)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+24)) = v56 + v57
	v60 = *(*int64)(unsafe.Add(mBase, uint32(v27)+32))
	v61 = *(*int64)(unsafe.Add(mBase, uint32(v43)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+32)) = v60 + v61
	v64 = *(*int64)(unsafe.Add(mBase, uint32(v27)+40))
	v65 = *(*int64)(unsafe.Add(mBase, uint32(v43)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+40)) = v64 + v65
	goto L18
L17:
	;
	v76 = v70
	goto L15
L18:
	;
	v69 = v35 + int32(1)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v69 < v70 {
		v35 = v69
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+100))
	v85 = v80 + int32(8)
	v86 = F_MemoryContextAlloc(m, v83, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v86
	if v85 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L8
L23:
	;
	v89 = F__emscripten_memcpy_bulkmem(m, v86, v12, v85)
	mBase = m.M
	goto L25
L24:
	;
	goto L25
L25:
	;
	goto L22
L26:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_dsa_free(m, v99, v98)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v104 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
	goto L28
L30:
	;
	F_dsa_detach(m, v104)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v109 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(0)
	goto L32
L34:
	;
	F_DestroyParallelContext(m, v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	F_pfree(m, l0)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L4
	} else {
		goto L38
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(0)
	goto L36
L38:
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
	var v8 int32
	_ = v8
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
	v8 = v2
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
	v11 = v8 * int32(36)
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
	v33 = v8 + int32(1)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v33 < v34 {
		v8 = v33
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
	v16 = int32(4470752)
	v17 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v19
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
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v17
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
	v73 = *(*int32)(unsafe.Add(mBase, _consts[55]))
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
	v90 = *(*int32)(unsafe.Add(mBase, _consts[55]))
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 float64
	_ = v45
	var v47 int32
	_ = v47
	var v51 float64
	_ = v51
	var v52 float64
	_ = v52
	var v55 float64
	_ = v55
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 float64
	_ = v92
	var v96 float64
	_ = v96
	var v97 float64
	_ = v97
	var v100 float64
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
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
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v425 int32
	_ = v425
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v534 int32
	_ = v534
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v636 int32
	_ = v636
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v661 int32
	_ = v661
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v691 int32
	_ = v691
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v727 int32
	_ = v727
	var v732 int32
	_ = v732
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v760 int32
	_ = v760
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v827 int32
	_ = v827
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
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
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v874 int32
	_ = v874
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v900 int32
	_ = v900
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v920 int32
	_ = v920
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v973 int32
	_ = v973
	var v983 int32
	_ = v983
	var v990 int32
	_ = v990
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v24 = v22 + int32(92)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v27 = base.I32_rem_s(v25, int32(5))
	switch v27 {
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
	m.G0 = v20 + int32(16)
	return
L2:
	;
	v1026 = F_BarrierArriveAndWait(m, v24, int32(134217751))
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L7
	} else {
		goto L187
	}
L3:
	;
	v851 = F_BarrierArriveAndWait(m, v24, int32(134217749))
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L7
	} else {
		goto L154
	}
L4:
	;
	F_ExecParallelHashEnsureBatchAccessors(m, l0)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L7
	} else {
		goto L67
	}
L5:
	;
	v328 = F_BarrierArriveAndWait(m, v24, int32(134217752))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L7
	} else {
		goto L66
	}
L6:
	;
	v29 = F_BarrierArriveAndWait(m, v24, int32(134217750))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	if v29 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v33
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = int32(0)
	F_ExecParallelHashCloseBatchAccessors(m, l0)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
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
	v309 = m.ExcPending
	if v309 != 0 {
		goto L7
	} else {
		goto L65
	}
L12:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v41 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	F_ExecParallelHashJoinSetUpBatches(m, l0, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L7
	} else {
		goto L27
	}
L14:
	;
	v45 = *(*float64)(unsafe.Add(mBase, _consts[345]))
	v47 = *(*int32)(unsafe.Add(mBase, _consts[326]))
	v51 = base.F64_mul(base.F64_mul(v45, base.F64_convert_i32_s(v47)), float64(1024))
	v52 = float64(4.294967295e+09)
	if base.F64_lt(v51, v52) != 0 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	goto L16
L16:
	;
	v81 = v41 << (uint(int32(1)) % 32)
	goto L13
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v63
	v65 = int32(1)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v69 = v67 << (uint(v65) % 32)
	if v69&(v69-v65) != 0 {
		goto L24
	} else {
		goto L25
	}
L18:
	;
	v55 = v51
	goto L20
L19:
	;
	v55 = v52
	goto L20
L20:
	;
	if base.F64_lt(v55, float64(4.294967296e+09))&base.F64_ge(v55, float64(0)) != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v61 = base.I32_trunc_f64_u(v55)
	v63 = v61
	goto L17
L22:
	;
	goto L23
L23:
	;
	v63 = int32(0)
	goto L17
L24:
	;
	v76 = v65 << (uint(int32(32)-base.I32_clz(v69)) % 32)
	goto L26
L25:
	;
	v76 = v69
	goto L26
L26:
	;
	v81 = v76
	goto L13
L27:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if v84 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v32)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v304
	goto L5
L29:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v32)+52))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	F_dsa_free(m, v88, v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L7
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	*(*int32)(unsafe.Add(mBase, uint32(v251))) = v252
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v255 = F_dsa_get_address(m, v254, v252)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L7
	} else {
		goto L60
	}
L32:
	;
	v92 = base.F64_convert_i32_u(v87)
	v96 = base.F64_ceil(base.F64_div(base.F64_add(v92, v92), base.F64_convert_i32_s(v81)))
	v97 = float64(1.34217728e+08)
	if base.F64_lt(v96, v97) != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v107 = int32(0)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v111 = int32(1024)
	if v106 <= v111 {
		goto L40
	} else {
		goto L41
	}
L34:
	;
	v100 = v96
	goto L36
L35:
	;
	v100 = v97
	goto L36
L36:
	;
	if base.F64_lt(base.F64_abs(v100), float64(2.147483648e+09)) != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v104 = base.I32_trunc_f64_s(v100)
	v106 = v104
	goto L33
L38:
	;
	goto L39
L39:
	;
	v106 = int32(-2147483648)
	goto L33
L40:
	;
	v114 = v111
	goto L42
L41:
	;
	v114 = v106
	goto L42
L42:
	;
	if v114&(v114-int32(1)) != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v121 = int32(1) << (uint(int32(32)-base.I32_clz(v114)) % 32)
	goto L45
L44:
	;
	v121 = v114
	goto L45
L45:
	;
	v125 = F_dsa_allocate_extended(m, v108, v121<<(uint(int32(2))%32), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L7
	} else {
		goto L46
	}
L46:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	*(*int32)(unsafe.Add(mBase, uint32(v128))) = v125
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v134 = F_dsa_get_address(m, v130, v133)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L7
	} else {
		goto L47
	}
L47:
	;
	if v121 <= int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v121
	goto L28
L49:
	;
	v139 = v121 & int32(7)
	if base.Ui32(int32(8)) <= base.Ui32(v121) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v146 = v107
	v148 = int32(0)
	goto L53
L51:
	;
	v187 = v107
	goto L52
L52:
	;
	if v139 == int32(0) {
		goto L48
	} else {
		goto L56
	}
L53:
	;
	v163 = v134 + v146<<(uint(int32(2))%32)
	v164 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = v164
	*(*int32)(unsafe.Add(mBase, uint32(v163)+4)) = v164
	*(*int32)(unsafe.Add(mBase, uint32(v163)+8)) = v164
	*(*int32)(unsafe.Add(mBase, uint32(v163)+12)) = v164
	*(*int32)(unsafe.Add(mBase, uint32(v163)+16)) = v164
	*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = v164
	*(*int32)(unsafe.Add(mBase, uint32(v163)+24)) = v164
	*(*int32)(unsafe.Add(mBase, uint32(v163)+28)) = v164
	v180 = int32(8)
	v181 = v146 + v180
	v183 = v148 + v180
	if v183 != v121&int32(-8) {
		v146 = v181
		v148 = v183
		goto L53
	} else {
		goto L55
	}
L54:
	;
	v187 = v181
	goto L52
L55:
	;
	goto L54
L56:
	;
	v206 = int32(0)
	v207 = v187
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v134+v207<<(uint(int32(2))%32)))) = int32(0)
	v227 = int32(1)
	v230 = v206 + v227
	if v230 != v139 {
		v206 = v230
		v207 = v207 + v227
		goto L57
	} else {
		goto L59
	}
L58:
	;
	goto L48
L59:
	;
	goto L58
L60:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v257 <= int32(0) {
		goto L28
	} else {
		goto L61
	}
L61:
	;
	v262 = int32(0)
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v255+v262<<(uint(int32(2))%32)))) = int32(0)
	v284 = v262 + int32(1)
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v284 < v285 {
		v262 = v284
		goto L62
	} else {
		goto L64
	}
L63:
	;
	goto L28
L64:
	;
	goto L63
L65:
	;
	goto L5
L66:
	;
	goto L4
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v352)))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v353)))
	v355 = F_dsa_get_address(m, v351, v354)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L7
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v355
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v358)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v359
	v362 = int32(1073741823)
	if v362 <= v359 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v374 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v374
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v374
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v373
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	*(*uint8)(unsafe.Add(mBase, uint32(v379)+24)) = uint8(v374)
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v384 = v382 + int32(40)
	v386 = F_LWLockAcquire(m, v384, v374)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L7
	} else {
		goto L76
	}
L70:
	;
	v365 = v362
	goto L72
L71:
	;
	v365 = v359
	goto L72
L72:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v365) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v373 = int32(32) - base.I32_clz(v365-int32(1))
	goto L75
L74:
	;
	v373 = int32(0)
	goto L75
L75:
	;
	goto L69
L76:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v382)+24))
	if v388 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v393 = v384
	v394 = v382 + int32(24)
	v403 = v388
	goto L80
L78:
	;
	v601 = v384
	goto L79
L79:
	;
	F_LWLockRelease(m, v601)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L7
	} else {
		goto L119
	}
L80:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v409 = F_dsa_get_address(m, v408, v403)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L7
	} else {
		goto L82
	}
L81:
	;
	v601 = v592
	goto L79
L82:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v409)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v394))) = v411
	F_LWLockRelease(m, v393)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L7
	} else {
		goto L83
	}
L83:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v409)+8))
	if v415 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v425 = int32(0)
	goto L87
L85:
	;
	goto L86
L86:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	F_dsa_free(m, v583, v403)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L7
	} else {
		goto L112
	}
L87:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v437 = v425 + (v409 + int32(16))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v437)+4))
	v440 = v437 + int32(4)
	v442 = v437 + int32(8)
	v443 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(int32(2)) <= base.Ui32(v443) {
		goto L91
	} else {
		goto L92
	}
L88:
	;
	goto L86
L89:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v543)+20))
	v545 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v543)+20)) = v544 + v545
	v548 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v553 = v548 + v534*int32(36) + int32(8)
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v553)))
	*(*int32)(unsafe.Add(mBase, uint32(v553))) = v554 + v545
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v442)))
	v563 = (v558+int32(15))&int32(-8) + v425
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v409)+8))
	if base.Ui32(v563) < base.Ui32(v564) {
		v425 = v563
		goto L87
	} else {
		goto L111
	}
L90:
	;
	v508 = v450 * int32(36)
	v509 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v512 = v508 + v509 + int32(16)
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v512)))
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v442)))
	*(*int32)(unsafe.Add(mBase, uint32(v512))) = v513 + (v514+int32(15))&int32(-8)
	v521 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v521+v508)+28))
	F_sts_puttuple(m, v523, v440, v442)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L7
	} else {
		goto L110
	}
L91:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v450 = (v443 - int32(1)) & base.I32_rotr(v438, v448)
	if v450 != 0 {
		goto L90
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v442)))
	v453 = int32(8)
	v457 = F_ExecParallelHashTupleAlloc(m, l0, v452+v453, v20+v453)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L7
	} else {
		goto L95
	}
L94:
	;
	goto L93
L95:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v440)))
	*(*int32)(unsafe.Add(mBase, uint32(v457)+4)) = v459
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v442)))
	if v463 != 0 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v473 = v467 + (v436-int32(1))&v438<<(uint(int32(2))%32)
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v473)))
	*(*int32)(unsafe.Add(mBase, uint32(v457))) = v474
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v473)))
	v477 = base.B2i32(v476 == v474)
	if v476 == v474 {
		goto L100
	} else {
		goto L101
	}
L97:
	;
	v464 = F__emscripten_memcpy_bulkmem(m, v457+int32(8), v442, v463)
	mBase = m.M
	goto L99
L98:
	;
	goto L99
L99:
	;
	goto L96
L100:
	;
	v478 = v466
	goto L102
L101:
	;
	v478 = v476
	goto L102
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v473))) = v478
	v480 = int32(0)
	if v476 == v474 {
		v534 = v480
		goto L89
	} else {
		goto L103
	}
L103:
	;
	v483 = v476
	goto L104
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v457))) = v483
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v473)))
	*(*int32)(unsafe.Add(mBase, uint32(v457))) = v499
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v473)))
	v502 = base.B2i32(v501 == v499)
	if v501 == v499 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v534 = v480
	goto L89
L106:
	;
	v503 = v466
	goto L108
L107:
	;
	v503 = v501
	goto L108
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v473))) = v503
	if v502 == int32(0) {
		v483 = v501
		goto L104
	} else {
		goto L109
	}
L109:
	;
	goto L105
L110:
	;
	v534 = v450
	goto L89
L111:
	;
	goto L88
L112:
	;
	v587 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v587 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L7
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v592 = v590 + int32(40)
	v594 = F_LWLockAcquire(m, v592, int32(0))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L7
	} else {
		goto L117
	}
L116:
	;
	goto L115
L117:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v590)+24))
	if v598 != 0 {
		v393 = v592
		v394 = v590 + int32(24)
		v403 = v598
		goto L80
	} else {
		goto L118
	}
L118:
	;
	goto L81
L119:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v618)+12))
	v620 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v618)+4))
	v622 = F_dsa_get_address(m, v620, v621)
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L7
	} else {
		goto L120
	}
L120:
	;
	v626 = F_palloc0(m, v619<<(uint(int32(2))%32))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L7
	} else {
		goto L121
	}
L121:
	;
	if int32(2) <= v619 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v636 = int32(1)
	goto L125
L123:
	;
	goto L124
L124:
	;
	F_pfree(m, v626)
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L7
	} else {
		goto L151
	}
L125:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v655)+28))
	goto L127
L126:
	;
	v691 = int32(1)
	goto L130
L127:
	;
	v661 = int32(1)
	v672 = *(*int32)(unsafe.Add(mBase, _consts[55]))
	v675 = F_sts_attach(m, v622-int32(-64)+(((v656*int32(28)+int32(76))<<(uint(v661)%32)+int32(14))&int32(-16)-int32(-64))*v636, v672+v661, v618+int32(168))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L7
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v626+v636<<(uint(int32(2))%32)))) = v675
	v679 = v636 + int32(1)
	if v679 != v619 {
		v636 = v679
		goto L125
	} else {
		goto L129
	}
L129:
	;
	goto L126
L130:
	;
	v701 = v626 + v691<<(uint(int32(2))%32)
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v701)))
	F_sts_begin_parallel_scan(m, v702)
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L7
	} else {
		goto L132
	}
L131:
	;
	goto L124
L132:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v701)))
	v708 = F_sts_parallel_scan_next(m, v705, v20+int32(12))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L7
	} else {
		goto L133
	}
L133:
	;
	if v708 != 0 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v711 = v708
	goto L137
L135:
	;
	goto L136
L136:
	;
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v701)))
	F_sts_end_parallel_scan(m, v803)
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L7
	} else {
		goto L149
	}
L137:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v711)))
	v732 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(int32(2)) <= base.Ui32(v732) {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	goto L136
L139:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v738 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v742 = (v732 - int32(1)) & base.I32_rotr(v737, v738)
	goto L141
L140:
	;
	v742 = int32(0)
	goto L141
L141:
	;
	v743 = int32(36)
	v744 = v742 * v743
	v745 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v748 = v744 + v745 + int32(16)
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v748)))
	*(*int32)(unsafe.Add(mBase, uint32(v748))) = v749 + (v727+int32(15))&int32(-8)
	v752 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v755 = v752 + v744 + int32(8)
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v755)))
	v757 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v755))) = v756 + v757
	v760 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v765 = v760 + v691*v743 + int32(20)
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v765)))
	*(*int32)(unsafe.Add(mBase, uint32(v765))) = v766 + v757
	v770 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v770+v744)+28))
	F_sts_puttuple(m, v772, v20+int32(12), v711)
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L7
	} else {
		goto L142
	}
L142:
	;
	v778 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v778 != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L7
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v701)))
	v784 = F_sts_parallel_scan_next(m, v781, v20+int32(12))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L7
	} else {
		goto L147
	}
L146:
	;
	goto L145
L147:
	;
	if v784 != 0 {
		v711 = v784
		goto L137
	} else {
		goto L148
	}
L148:
	;
	goto L138
L149:
	;
	v807 = v691 + int32(1)
	if v807 != v619 {
		v691 = v807
		goto L130
	} else {
		goto L150
	}
L150:
	;
	goto L131
L151:
	;
	F_ExecParallelHashMergeCounters(m, l0)
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L7
	} else {
		goto L152
	}
L152:
	;
	v831 = F_BarrierArriveAndWait(m, v24, int32(134217753))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L7
	} else {
		goto L153
	}
L153:
	;
	goto L3
L154:
	;
	if v851 == int32(0) {
		goto L2
	} else {
		goto L155
	}
L155:
	;
	F_ExecParallelHashEnsureBatchAccessors(m, l0)
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L7
	} else {
		goto L156
	}
L156:
	;
	v857 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v857
	v860 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v861 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v861)))
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v862)))
	v864 = F_dsa_get_address(m, v860, v863)
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L7
	} else {
		goto L157
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v864
	v867 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v867)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v868
	v871 = int32(1073741823)
	if v871 <= v868 {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	v883 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v883
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v883
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v882
	v888 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	*(*uint8)(unsafe.Add(mBase, uint32(v888)+24)) = uint8(v883)
	v891 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v893 = F_dsa_get_address(m, v891, v892)
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L7
	} else {
		goto L165
	}
L159:
	;
	v874 = v871
	goto L161
L160:
	;
	v874 = v868
	goto L161
L161:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v874) {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v882 = int32(32) - base.I32_clz(v874-int32(1))
	goto L164
L163:
	;
	v882 = int32(0)
	goto L164
L164:
	;
	goto L158
L165:
	;
	v895 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v895 <= int32(0) {
		v990 = v857
		goto L166
	} else {
		goto L167
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v990
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	F_dsa_free(m, v1002, v1003)
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L7
	} else {
		goto L186
	}
L167:
	;
	v900 = int32(0)
	v904 = v900
	v907 = v900
	v910 = v900
	goto L168
L168:
	;
	v920 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v920+v904*int32(36))))
	v925 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v924)+60)))
	if v925 == int32(0) {
		goto L171
	} else {
		goto L172
	}
L169:
	;
	v973 = v966 | base.B2i32(int32(1073741822) < v969)
	if (v973|v932)&int32(1) == int32(0) {
		v990 = v857
		goto L166
	} else {
		goto L182
	}
L170:
	;
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v934 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v935 = *(*int32)(unsafe.Add(mBase, uint32(v934)+28))
	goto L176
L171:
	;
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v924)+48))
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
	if base.Ui32(v928) <= base.Ui32(v929) {
		v932 = v910
		goto L170
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v932 = int32(1)
	goto L170
L174:
	;
	goto L173
L175:
	;
	v968 = v904 + int32(1)
	v969 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v968 < v969 {
		v904 = v968
		v907 = v966
		v910 = v932
		goto L168
	} else {
		goto L181
	}
L176:
	;
	v948 = base.I32_rem_s(v904, v933)
	v951 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v893+int32(60)+(((v935*int32(28)+int32(76))<<(uint(int32(1))%32)+int32(14))&int32(-16)-int32(-64))*v948))))
	if v951 == int32(0) {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v924)+48))
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
	if base.Ui32(v954) <= base.Ui32(v955) {
		v966 = v907
		goto L175
	} else {
		goto L180
	}
L178:
	;
	goto L179
L179:
	;
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v924)+52))
	v958 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v958+v948*int32(36))))
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v962)+56))
	v966 = base.B2i32(v957 == v963) | v907
	goto L175
L180:
	;
	goto L179
L181:
	;
	goto L169
L182:
	;
	if v973&int32(1) != 0 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v983 = int32(3)
	goto L185
L184:
	;
	v983 = int32(2)
	goto L185
L185:
	;
	v990 = v983
	goto L166
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = int32(0)
	goto L2
L187:
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
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int64
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
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
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int64
	_ = v241
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int64
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
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
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int64
	_ = v327
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int64
	_ = v342
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v362 int64
	_ = v362
	var v366 int32
	_ = v366
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int64
	_ = v384
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int64
	_ = v397
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int64
	_ = v410
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
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
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int64
	_ = v481
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int64
	_ = v510
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int64
	_ = v539
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int64
	_ = v568
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
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
			v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220)+36)))
			if v221 != int32(1) {
				v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
				mBase = m.M
				v582 = m.ExcPending
				if v582 != 0 {
					return int32(0)
				} else {
					return v581
				}
			} else {
				v224 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+52))
				v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
				v227 = F_shm_toc_allocate(m, v225, v226)
				mBase = m.M
				v228 = m.ExcPending
				if v228 != 0 {
					return int32(0)
				} else {
					v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
					v232 = F__emscripten_memset_bulkmem(m, v227, base.I32_extend8_s(int32(0)), v230)
					mBase = m.M
					v233 = int32(77)
					*(*uint16)(unsafe.Add(mBase, uint32(v232))) = uint16(v233)
					*(*int32)(unsafe.Add(mBase, uint32(v232)+4)) = int32(1073741824)
					*(*int64)(unsafe.Add(mBase, uint32(v232)+8)) = int64(-1)
					v239 = *(*int32)(unsafe.Add(mBase, uint32(v224)+52))
					v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v241 = int64(*(*int32)(unsafe.Add(mBase, uint32(v240)+40)))
					F_shm_toc_insert(m, v239, v241, v232)
					mBase = m.M
					v243 = m.ExcPending
					if v243 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = int32(695)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v232
						v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
						mBase = m.M
						v582 = m.ExcPending
						if v582 != 0 {
							return int32(0)
						} else {
							return v581
						}
					}
				}
			}
		default:
			v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
			mBase = m.M
			v582 = m.ExcPending
			if v582 != 0 {
				return int32(0)
			} else {
				return v581
			}
		case 6:
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+36)))
			if v29 != int32(1) {
				v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
				mBase = m.M
				v582 = m.ExcPending
				if v582 != 0 {
					return int32(0)
				} else {
					return v581
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
								v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
								mBase = m.M
								v582 = m.ExcPending
								if v582 != 0 {
									return int32(0)
								} else {
									return v581
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
				v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
				mBase = m.M
				v582 = m.ExcPending
				if v582 != 0 {
					return int32(0)
				} else {
					return v581
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
								v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
								mBase = m.M
								v582 = m.ExcPending
								if v582 != 0 {
									return int32(0)
								} else {
									return v581
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
											v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
											mBase = m.M
											v582 = m.ExcPending
											if v582 != 0 {
												return int32(0)
											} else {
												return v581
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
												v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
												mBase = m.M
												v582 = m.ExcPending
												if v582 != 0 {
													return int32(0)
												} else {
													return v581
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
											v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
											mBase = m.M
											v582 = m.ExcPending
											if v582 != 0 {
												return int32(0)
											} else {
												return v581
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
				v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
				mBase = m.M
				v582 = m.ExcPending
				if v582 != 0 {
					return int32(0)
				} else {
					return v581
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
								v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
								mBase = m.M
								v582 = m.ExcPending
								if v582 != 0 {
									return int32(0)
								} else {
									return v581
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
											v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
											mBase = m.M
											v582 = m.ExcPending
											if v582 != 0 {
												return int32(0)
											} else {
												return v581
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
												v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
												mBase = m.M
												v582 = m.ExcPending
												if v582 != 0 {
													return int32(0)
												} else {
													return v581
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
											v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
											mBase = m.M
											v582 = m.ExcPending
											if v582 != 0 {
												return int32(0)
											} else {
												return v581
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
				v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
				mBase = m.M
				v582 = m.ExcPending
				if v582 != 0 {
					return int32(0)
				} else {
					return v581
				}
			} else {
				v174 = *(*int32)(unsafe.Add(mBase, uint32(v170)+12))
				if v174 == int32(0) {
					v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
					mBase = m.M
					v582 = m.ExcPending
					if v582 != 0 {
						return int32(0)
					} else {
						return v581
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
							v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
							v193 = F__emscripten_memset_bulkmem(m, v190, base.I32_extend8_s(int32(0)), v181)
							mBase = m.M
							v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
							v195 = *(*int32)(unsafe.Add(mBase, uint32(v170)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v194))) = v195
							v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
							mBase = m.M
							v582 = m.ExcPending
							if v582 != 0 {
								return int32(0)
							} else {
								return v581
							}
						}
					}
				}
			}
		case 11:
			v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268)+36)))
			if v269 != int32(1) {
				v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
				mBase = m.M
				v582 = m.ExcPending
				if v582 != 0 {
					return int32(0)
				} else {
					return v581
				}
			} else {
				v272 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v274 = *(*int32)(unsafe.Add(mBase, uint32(v273)+172))
				if v274 != 0 {
					v275 = int32(24)
					v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v276 == int32(0) {
						v292 = v275
						v293 = *(*int32)(unsafe.Add(mBase, uint32(v272)+52))
						v294 = F_shm_toc_allocate(m, v293, v292)
						mBase = m.M
						v295 = m.ExcPending
						if v295 != 0 {
							return int32(0)
						} else {
							v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							if v296 != 0 {
								v299 = int32(0)
								v300 = *(*int32)(unsafe.Add(mBase, uint32(v272)+12))
								if v299 < v300 {
									v303 = v294 + int32(24)
								} else {
									v303 = v299
								}
								v304 = v303
							} else {
								v304 = v3
							}
							*(*int32)(unsafe.Add(mBase, uint32(v294)+8)) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v294))) = int64(0)
							v310 = v294 + int32(12)
							*(*int32)(unsafe.Add(mBase, uint32(v310)+8)) = int32(-1)
							*(*int64)(unsafe.Add(mBase, uint32(v310))) = int64(-4294967296)
							if v304 != 0 {
								v315 = *(*int32)(unsafe.Add(mBase, uint32(v272)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v304))) = v315
								v323 = F__emscripten_memset_bulkmem(m, v304+int32(8), base.I32_extend8_s(int32(0)), v315<<(uint(int32(4))%32))
								mBase = m.M
							} else {
							}
							v325 = *(*int32)(unsafe.Add(mBase, uint32(v272)+52))
							v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v327 = int64(*(*int32)(unsafe.Add(mBase, uint32(v326)+40)))
							F_shm_toc_insert(m, v325, v327, v294)
							mBase = m.M
							v329 = m.ExcPending
							if v329 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v304
								*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v294
								v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
								mBase = m.M
								v582 = m.ExcPending
								if v582 != 0 {
									return int32(0)
								} else {
									return v581
								}
							}
						}
					} else {
						v279 = *(*int32)(unsafe.Add(mBase, uint32(v272)+12))
						if v279 <= int32(0) {
							v292 = v275
							v293 = *(*int32)(unsafe.Add(mBase, uint32(v272)+52))
							v294 = F_shm_toc_allocate(m, v293, v292)
							mBase = m.M
							v295 = m.ExcPending
							if v295 != 0 {
								return int32(0)
							} else {
								v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								if v296 != 0 {
									v299 = int32(0)
									v300 = *(*int32)(unsafe.Add(mBase, uint32(v272)+12))
									if v299 < v300 {
										v303 = v294 + int32(24)
									} else {
										v303 = v299
									}
									v304 = v303
								} else {
									v304 = v3
								}
								*(*int32)(unsafe.Add(mBase, uint32(v294)+8)) = int32(0)
								*(*int64)(unsafe.Add(mBase, uint32(v294))) = int64(0)
								v310 = v294 + int32(12)
								*(*int32)(unsafe.Add(mBase, uint32(v310)+8)) = int32(-1)
								*(*int64)(unsafe.Add(mBase, uint32(v310))) = int64(-4294967296)
								if v304 != 0 {
									v315 = *(*int32)(unsafe.Add(mBase, uint32(v272)+12))
									*(*int32)(unsafe.Add(mBase, uint32(v304))) = v315
									v323 = F__emscripten_memset_bulkmem(m, v304+int32(8), base.I32_extend8_s(int32(0)), v315<<(uint(int32(4))%32))
									mBase = m.M
								} else {
								}
								v325 = *(*int32)(unsafe.Add(mBase, uint32(v272)+52))
								v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v327 = int64(*(*int32)(unsafe.Add(mBase, uint32(v326)+40)))
								F_shm_toc_insert(m, v325, v327, v294)
								mBase = m.M
								v329 = m.ExcPending
								if v329 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v304
									*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v294
									v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
									mBase = m.M
									v582 = m.ExcPending
									if v582 != 0 {
										return int32(0)
									} else {
										return v581
									}
								}
							}
						} else {
							v284 = F_add_size(m, int32(24), int32(8))
							mBase = m.M
							v285 = m.ExcPending
							if v285 != 0 {
								return int32(0)
							} else {
								v286 = *(*int32)(unsafe.Add(mBase, uint32(v272)+12))
								v288 = F_mul_size(m, v286, int32(16))
								mBase = m.M
								v289 = m.ExcPending
								if v289 != 0 {
									return int32(0)
								} else {
									v290 = F_add_size(m, v284, v288)
									mBase = m.M
									v291 = m.ExcPending
									if v291 != 0 {
										return int32(0)
									} else {
										v292 = v290
										v293 = *(*int32)(unsafe.Add(mBase, uint32(v272)+52))
										v294 = F_shm_toc_allocate(m, v293, v292)
										mBase = m.M
										v295 = m.ExcPending
										if v295 != 0 {
											return int32(0)
										} else {
											v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											if v296 != 0 {
												v299 = int32(0)
												v300 = *(*int32)(unsafe.Add(mBase, uint32(v272)+12))
												if v299 < v300 {
													v303 = v294 + int32(24)
												} else {
													v303 = v299
												}
												v304 = v303
											} else {
												v304 = v3
											}
											*(*int32)(unsafe.Add(mBase, uint32(v294)+8)) = int32(0)
											*(*int64)(unsafe.Add(mBase, uint32(v294))) = int64(0)
											v310 = v294 + int32(12)
											*(*int32)(unsafe.Add(mBase, uint32(v310)+8)) = int32(-1)
											*(*int64)(unsafe.Add(mBase, uint32(v310))) = int64(-4294967296)
											if v304 != 0 {
												v315 = *(*int32)(unsafe.Add(mBase, uint32(v272)+12))
												*(*int32)(unsafe.Add(mBase, uint32(v304))) = v315
												v323 = F__emscripten_memset_bulkmem(m, v304+int32(8), base.I32_extend8_s(int32(0)), v315<<(uint(int32(4))%32))
												mBase = m.M
											} else {
											}
											v325 = *(*int32)(unsafe.Add(mBase, uint32(v272)+52))
											v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											v327 = int64(*(*int32)(unsafe.Add(mBase, uint32(v326)+40)))
											F_shm_toc_insert(m, v325, v327, v294)
											mBase = m.M
											v329 = m.ExcPending
											if v329 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v304
												*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v294
												v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
												mBase = m.M
												v582 = m.ExcPending
												if v582 != 0 {
													return int32(0)
												} else {
													return v581
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
					mBase = m.M
					v582 = m.ExcPending
					if v582 != 0 {
						return int32(0)
					} else {
						return v581
					}
				}
			}
		case 21:
			v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199)+36)))
			if v200 != int32(1) {
				v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
				mBase = m.M
				v582 = m.ExcPending
				if v582 != 0 {
					return int32(0)
				} else {
					return v581
				}
			} else {
				v203 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)+148))
				if v205 != 0 {
					v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v207 = int64(*(*int32)(unsafe.Add(mBase, uint32(v206)+40)))
					v208 = *(*int32)(unsafe.Add(mBase, uint32(v203)+52))
					v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
					v210 = F_shm_toc_allocate(m, v208, v209)
					mBase = m.M
					v211 = m.ExcPending
					if v211 != 0 {
						return int32(0)
					} else {
						v212 = *(*int32)(unsafe.Add(mBase, uint32(v204)+148))
						m.T0[v212].(func(*base.Module, int32, int32, int32))(m, l0, v203, v210)
						mBase = m.M
						v214 = m.ExcPending
						if v214 != 0 {
							return int32(0)
						} else {
							v215 = *(*int32)(unsafe.Add(mBase, uint32(v203)+52))
							F_shm_toc_insert(m, v215, v207, v210)
							mBase = m.M
							v217 = m.ExcPending
							if v217 != 0 {
								return int32(0)
							} else {
								v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
								mBase = m.M
								v582 = m.ExcPending
								if v582 != 0 {
									return int32(0)
								} else {
									return v581
								}
							}
						}
					}
				} else {
					v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
					mBase = m.M
					v582 = m.ExcPending
					if v582 != 0 {
						return int32(0)
					} else {
						return v581
					}
				}
			}
		case 22:
			v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247)+36)))
			if v248 != int32(1) {
				v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
				mBase = m.M
				v582 = m.ExcPending
				if v582 != 0 {
					return int32(0)
				} else {
					return v581
				}
			} else {
				v251 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v253 = *(*int32)(unsafe.Add(mBase, uint32(v252)+32))
				if v253 != 0 {
					v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v255 = int64(*(*int32)(unsafe.Add(mBase, uint32(v254)+40)))
					v256 = *(*int32)(unsafe.Add(mBase, uint32(v251)+52))
					v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
					v258 = F_shm_toc_allocate(m, v256, v257)
					mBase = m.M
					v259 = m.ExcPending
					if v259 != 0 {
						return int32(0)
					} else {
						v260 = *(*int32)(unsafe.Add(mBase, uint32(v252)+32))
						m.T0[v260].(func(*base.Module, int32, int32, int32))(m, l0, v251, v258)
						mBase = m.M
						v262 = m.ExcPending
						if v262 != 0 {
							return int32(0)
						} else {
							v263 = *(*int32)(unsafe.Add(mBase, uint32(v251)+52))
							F_shm_toc_insert(m, v263, v255, v258)
							mBase = m.M
							v265 = m.ExcPending
							if v265 != 0 {
								return int32(0)
							} else {
								v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
								mBase = m.M
								v582 = m.ExcPending
								if v582 != 0 {
									return int32(0)
								} else {
									return v581
								}
							}
						}
					}
				} else {
					v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
					mBase = m.M
					v582 = m.ExcPending
					if v582 != 0 {
						return int32(0)
					} else {
						return v581
					}
				}
			}
		case 26:
			v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+36)))
			if v336 != int32(1) {
				v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
				mBase = m.M
				v582 = m.ExcPending
				if v582 != 0 {
					return int32(0)
				} else {
					return v581
				}
			} else {
				v339 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v340 = *(*int32)(unsafe.Add(mBase, uint32(v339)+44))
				if v340 != 0 {
					v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v342 = int64(*(*int32)(unsafe.Add(mBase, uint32(v341)+40)))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(631)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(719)
					v347 = *(*int32)(unsafe.Add(mBase, uint32(v339)+52))
					v349 = F_shm_toc_allocate(m, v347, int32(220))
					mBase = m.M
					v350 = m.ExcPending
					if v350 != 0 {
						return int32(0)
					} else {
						v351 = *(*int32)(unsafe.Add(mBase, uint32(v339)+52))
						F_shm_toc_insert(m, v351, v342, v349)
						mBase = m.M
						v353 = m.ExcPending
						if v353 != 0 {
							return int32(0)
						} else {
							v354 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v349)+32)) = v354
							*(*int32)(unsafe.Add(mBase, uint32(v349)+8)) = v354
							*(*int32)(unsafe.Add(mBase, uint32(v349)+164)) = v354
							*(*int32)(unsafe.Add(mBase, uint32(v349)+24)) = v354
							v362 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v349)+16)) = v362
							*(*int64)(unsafe.Add(mBase, uint32(v349))) = v362
							v366 = *(*int32)(unsafe.Add(mBase, uint32(v339)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v349)+36)) = v354
							*(*int32)(unsafe.Add(mBase, uint32(v349)+28)) = v366 + int32(1)
							v373 = v349 + int32(40)
							v374 = int32(69)
							*(*uint16)(unsafe.Add(mBase, uint32(v373))) = uint16(v374)
							*(*int32)(unsafe.Add(mBase, uint32(v373)+4)) = int32(1073741824)
							*(*int64)(unsafe.Add(mBase, uint32(v373)+8)) = int64(-1)
							v381 = v349 + int32(56)
							v382 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v381)+8)) = v382
							v384 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v381))) = v384
							*(*int64)(unsafe.Add(mBase, uint32(v381)+12)) = v384
							*(*uint8)(unsafe.Add(mBase, uint32(v381)+20)) = uint8(v382)
							F_ConditionVariableInit(m, v349+int32(80))
							mBase = m.M
							v394 = v349 + int32(92)
							v395 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v394)+8)) = v395
							v397 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v394))) = v397
							*(*int64)(unsafe.Add(mBase, uint32(v394)+12)) = v397
							*(*uint8)(unsafe.Add(mBase, uint32(v394)+20)) = uint8(v395)
							F_ConditionVariableInit(m, v349+int32(116))
							mBase = m.M
							v407 = v349 + int32(128)
							v408 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v407)+8)) = v408
							v410 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v407))) = v410
							*(*int64)(unsafe.Add(mBase, uint32(v407)+12)) = v410
							*(*uint8)(unsafe.Add(mBase, uint32(v407)+20)) = uint8(v408)
							F_ConditionVariableInit(m, v349+int32(152))
							mBase = m.M
							v421 = *(*int32)(unsafe.Add(mBase, uint32(v339)+44))
							F_SharedFileSetInit(m, v349+int32(168), v421)
							mBase = m.M
							v423 = m.ExcPending
							if v423 != 0 {
								return int32(0)
							} else {
								v424 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								*(*int32)(unsafe.Add(mBase, uint32(v424)+128)) = v349
								v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
								mBase = m.M
								v582 = m.ExcPending
								if v582 != 0 {
									return int32(0)
								} else {
									return v581
								}
							}
						}
					}
				} else {
					v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
					mBase = m.M
					v582 = m.ExcPending
					if v582 != 0 {
						return int32(0)
					} else {
						return v581
					}
				}
			}
		case 28:
			v545 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v546 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v546 == int32(0) {
				v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
				mBase = m.M
				v582 = m.ExcPending
				if v582 != 0 {
					return int32(0)
				} else {
					return v581
				}
			} else {
				v549 = *(*int32)(unsafe.Add(mBase, uint32(v545)+12))
				if v549 == int32(0) {
					v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
					mBase = m.M
					v582 = m.ExcPending
					if v582 != 0 {
						return int32(0)
					} else {
						return v581
					}
				} else {
					v552 = *(*int32)(unsafe.Add(mBase, uint32(v545)+52))
					v556 = v549*int32(40) + int32(8)
					v557 = F_shm_toc_allocate(m, v552, v556)
					mBase = m.M
					v558 = m.ExcPending
					if v558 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = v557
						v562 = F__emscripten_memset_bulkmem(m, v557, base.I32_extend8_s(int32(0)), v556)
						mBase = m.M
						v563 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
						v564 = *(*int32)(unsafe.Add(mBase, uint32(v545)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v563))) = v564
						v566 = *(*int32)(unsafe.Add(mBase, uint32(v545)+52))
						v567 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v568 = int64(*(*int32)(unsafe.Add(mBase, uint32(v567)+40)))
						v569 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
						F_shm_toc_insert(m, v566, v568, v569)
						mBase = m.M
						v571 = m.ExcPending
						if v571 != 0 {
							return int32(0)
						} else {
							v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
							mBase = m.M
							v582 = m.ExcPending
							if v582 != 0 {
								return int32(0)
							} else {
								return v581
							}
						}
					}
				}
			}
		case 29:
			v458 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v459 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v459 == int32(0) {
				v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
				mBase = m.M
				v582 = m.ExcPending
				if v582 != 0 {
					return int32(0)
				} else {
					return v581
				}
			} else {
				v462 = *(*int32)(unsafe.Add(mBase, uint32(v458)+12))
				if v462 == int32(0) {
					v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
					mBase = m.M
					v582 = m.ExcPending
					if v582 != 0 {
						return int32(0)
					} else {
						return v581
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
						v475 = F__emscripten_memset_bulkmem(m, v470, base.I32_extend8_s(int32(0)), v469)
						mBase = m.M
						v476 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
						v477 = *(*int32)(unsafe.Add(mBase, uint32(v458)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v476))) = v477
						v479 = *(*int32)(unsafe.Add(mBase, uint32(v458)+52))
						v480 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v481 = int64(*(*int32)(unsafe.Add(mBase, uint32(v480)+40)))
						v482 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
						F_shm_toc_insert(m, v479, v481, v482)
						mBase = m.M
						v484 = m.ExcPending
						if v484 != 0 {
							return int32(0)
						} else {
							v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
							mBase = m.M
							v582 = m.ExcPending
							if v582 != 0 {
								return int32(0)
							} else {
								return v581
							}
						}
					}
				}
			}
		case 30:
			v487 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v488 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v488 == int32(0) {
				v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
				mBase = m.M
				v582 = m.ExcPending
				if v582 != 0 {
					return int32(0)
				} else {
					return v581
				}
			} else {
				v491 = *(*int32)(unsafe.Add(mBase, uint32(v487)+12))
				if v491 == int32(0) {
					v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
					mBase = m.M
					v582 = m.ExcPending
					if v582 != 0 {
						return int32(0)
					} else {
						return v581
					}
				} else {
					v494 = *(*int32)(unsafe.Add(mBase, uint32(v487)+52))
					v498 = v491*int32(96) | int32(8)
					v499 = F_shm_toc_allocate(m, v494, v498)
					mBase = m.M
					v500 = m.ExcPending
					if v500 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+284)) = v499
						v504 = F__emscripten_memset_bulkmem(m, v499, base.I32_extend8_s(int32(0)), v498)
						mBase = m.M
						v505 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
						v506 = *(*int32)(unsafe.Add(mBase, uint32(v487)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v505))) = v506
						v508 = *(*int32)(unsafe.Add(mBase, uint32(v487)+52))
						v509 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v510 = int64(*(*int32)(unsafe.Add(mBase, uint32(v509)+40)))
						v511 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
						F_shm_toc_insert(m, v508, v510, v511)
						mBase = m.M
						v513 = m.ExcPending
						if v513 != 0 {
							return int32(0)
						} else {
							v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
							mBase = m.M
							v582 = m.ExcPending
							if v582 != 0 {
								return int32(0)
							} else {
								return v581
							}
						}
					}
				}
			}
		case 32:
			v516 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v517 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v517 == int32(0) {
				v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
				mBase = m.M
				v582 = m.ExcPending
				if v582 != 0 {
					return int32(0)
				} else {
					return v581
				}
			} else {
				v520 = *(*int32)(unsafe.Add(mBase, uint32(v516)+12))
				if v520 == int32(0) {
					v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
					mBase = m.M
					v582 = m.ExcPending
					if v582 != 0 {
						return int32(0)
					} else {
						return v581
					}
				} else {
					v523 = *(*int32)(unsafe.Add(mBase, uint32(v516)+52))
					v527 = v520*int32(24) + int32(8)
					v528 = F_shm_toc_allocate(m, v523, v527)
					mBase = m.M
					v529 = m.ExcPending
					if v529 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+352)) = v528
						v533 = F__emscripten_memset_bulkmem(m, v528, base.I32_extend8_s(int32(0)), v527)
						mBase = m.M
						v534 = *(*int32)(unsafe.Add(mBase, uint32(l0)+352))
						v535 = *(*int32)(unsafe.Add(mBase, uint32(v516)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v534))) = v535
						v537 = *(*int32)(unsafe.Add(mBase, uint32(v516)+52))
						v538 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v539 = int64(*(*int32)(unsafe.Add(mBase, uint32(v538)+40)))
						v540 = *(*int32)(unsafe.Add(mBase, uint32(l0)+352))
						F_shm_toc_insert(m, v537, v539, v540)
						mBase = m.M
						v542 = m.ExcPending
						if v542 != 0 {
							return int32(0)
						} else {
							v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
							mBase = m.M
							v582 = m.ExcPending
							if v582 != 0 {
								return int32(0)
							} else {
								return v581
							}
						}
					}
				}
			}
		case 37:
			v429 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v430 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v430 == int32(0) {
				v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
				mBase = m.M
				v582 = m.ExcPending
				if v582 != 0 {
					return int32(0)
				} else {
					return v581
				}
			} else {
				v433 = *(*int32)(unsafe.Add(mBase, uint32(v429)+12))
				if v433 == int32(0) {
					v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
					mBase = m.M
					v582 = m.ExcPending
					if v582 != 0 {
						return int32(0)
					} else {
						return v581
					}
				} else {
					v436 = *(*int32)(unsafe.Add(mBase, uint32(v429)+52))
					v440 = v433*int32(20) + int32(4)
					v441 = F_shm_toc_allocate(m, v436, v440)
					mBase = m.M
					v442 = m.ExcPending
					if v442 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v441
						v446 = F__emscripten_memset_bulkmem(m, v441, base.I32_extend8_s(int32(0)), v440)
						mBase = m.M
						v447 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
						v448 = *(*int32)(unsafe.Add(mBase, uint32(v429)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v447))) = v448
						v450 = *(*int32)(unsafe.Add(mBase, uint32(v429)+52))
						v451 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v452 = int64(*(*int32)(unsafe.Add(mBase, uint32(v451)+40)))
						v453 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
						F_shm_toc_insert(m, v450, v452, v453)
						mBase = m.M
						v455 = m.ExcPending
						if v455 != 0 {
							return int32(0)
						} else {
							v581 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
							mBase = m.M
							v582 = m.ExcPending
							if v582 != 0 {
								return int32(0)
							} else {
								return v581
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
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int64
	_ = v88
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int64
	_ = v178
	var v180 int64
	_ = v180
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
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
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
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
	var v478 int32
	_ = v478
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
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v572 int64
	_ = v572
	var v575 int32
	_ = v575
	var v576 int64
	_ = v576
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v616 int32
	_ = v616
	var v622 int32
	_ = v622
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v641 int32
	_ = v641
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v785 int32
	_ = v785
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v808 int32
	_ = v808
	var v813 int32
	_ = v813
	var v817 int32
	_ = v817
	var v822 int32
	_ = v822
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v857 int32
	_ = v857
	var v861 int32
	_ = v861
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v870 int64
	_ = v870
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v958 int32
	_ = v958
	var v964 int32
	_ = v964
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v984 int32
	_ = v984
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1002 int32
	_ = v1002
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1009 int32
	_ = v1009
	var v1011 int32
	_ = v1011
	var v1014 int32
	_ = v1014
	var v1016 int32
	_ = v1016
	var v1019 int32
	_ = v1019
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1043 int32
	_ = v1043
	var v1045 int32
	_ = v1045
	var v1048 int32
	_ = v1048
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1055 int32
	_ = v1055
	var v1060 int32
	_ = v1060
	var v1065 int32
	_ = v1065
	var v1068 int32
	_ = v1068
	var v1072 int32
	_ = v1072
	var v1076 int32
	_ = v1076
	var v1083 int32
	_ = v1083
	var v1086 int32
	_ = v1086
	var v1090 int32
	_ = v1090
	var v1095 int32
	_ = v1095
	var v1119 int32
	_ = v1119
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1132 int32
	_ = v1132
	var v1135 int32
	_ = v1135
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1171 int32
	_ = v1171
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1177 int32
	_ = v1177
	var v1180 int32
	_ = v1180
	var v1183 int32
	_ = v1183
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1194 int32
	_ = v1194
	var v1196 int32
	_ = v1196
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1203 int32
	_ = v1203
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1221 int32
	_ = v1221
	var v1231 int32
	_ = v1231
	var v1248 int32
	_ = v1248
	var v1256 int32
	_ = v1256
	var v1260 int32
	_ = v1260
	var v1265 int32
	_ = v1265
	var v1269 int32
	_ = v1269
	var v1273 int32
	_ = v1273
	var v1278 int32
	_ = v1278
	var v1282 int32
	_ = v1282
	var v1286 int32
	_ = v1286
	var v1291 int32
	_ = v1291
	var v1295 int32
	_ = v1295
	var v1299 int32
	_ = v1299
	var v1304 int32
	_ = v1304
	var v1308 int32
	_ = v1308
	var v1311 int32
	_ = v1311
	var v1315 int32
	_ = v1315
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1327 int32
	_ = v1327
	var v1329 int32
	_ = v1329
	var v1334 int32
	_ = v1334
	var v1337 int32
	_ = v1337
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1345 int32
	_ = v1345
	var v1347 int32
	_ = v1347
	var v1351 int32
	_ = v1351
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1361 int32
	_ = v1361
	var v1377 int32
	_ = v1377
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1386 int32
	_ = v1386
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1407 int32
	_ = v1407
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1417 int32
	_ = v1417
	var v1433 int32
	_ = v1433
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1466 int32
	_ = v1466
	var v1469 int32
	_ = v1469
	var v1475 int32
	_ = v1475
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1482 int32
	_ = v1482
	var v1484 int32
	_ = v1484
	var v1489 int32
	_ = v1489
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1495 int32
	_ = v1495
	var v1504 int32
	_ = v1504
	var v1506 int32
	_ = v1506
	var v1507 int32
	_ = v1507
	var v1513 int32
	_ = v1513
	var v1517 int32
	_ = v1517
	var v1518 int32
	_ = v1518
	var v1523 int32
	_ = v1523
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1532 int32
	_ = v1532
	var v1534 int32
	_ = v1534
	var v1536 int32
	_ = v1536
	var v1540 int32
	_ = v1540
	var v1541 int32
	_ = v1541
	var v1543 int32
	_ = v1543
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1549 int32
	_ = v1549
	var v1550 int32
	_ = v1550
	var v1552 int32
	_ = v1552
	var v1554 int32
	_ = v1554
	var v1555 int32
	_ = v1555
	var v1561 int32
	_ = v1561
	var v1562 int32
	_ = v1562
	var v1563 int32
	_ = v1563
	var v1564 int32
	_ = v1564
	var v1587 int32
	_ = v1587
	var v1590 int32
	_ = v1590
	var v1594 int32
	_ = v1594
	var v1599 int32
	_ = v1599
	var v1603 int32
	_ = v1603
	var v1606 int32
	_ = v1606
	var v1610 int32
	_ = v1610
	var v1615 int32
	_ = v1615
	var v1619 int32
	_ = v1619
	var v1625 int32
	_ = v1625
	var v1630 int32
	_ = v1630
	var v1634 int32
	_ = v1634
	var v1636 int32
	_ = v1636
	var v1637 int32
	_ = v1637
	var v1641 int32
	_ = v1641
	var v1646 int32
	_ = v1646
	v2 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(32)
	m.G0 = v19
	v22 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[84])) = uint8(v22)
	v25 = int32(295)
	v27 = m.G0
	v29 = v27 - int32(144)
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
	v67 = m.ExcPending
	if v67 != 0 {
		goto L14
	} else {
		goto L15
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v39
	F_sigemptyset(m, v29+int32(8))
	mBase = m.M
	goto L5
L3:
	;
	*(*int32)(unsafe.Add(mBase, _consts[85])) = v25
	v39 = int32(4729)
	goto L2
L5:
	;
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+136)) = int32(268435456)
	v51 = v29 + int32(4)
	goto L9
L7:
	;
	m.G0 = v29 + int32(144)
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
	v62 = F___memcpy(m, int32(4637764), v51, int32(140))
	mBase = m.M
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
	v70 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+1328))
	*(*int32)(unsafe.Add(mBase, _consts[55])) = v71
	v75 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	v80 = F_AllocSetContextCreateInternal(m, v75, int32(218095), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v80
	v83 = F_dsm_attach(m, l0)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L14
	} else {
		goto L20
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1634 = m.ExcPending
	if v1634 != 0 {
		goto L14
	} else {
		goto L422
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1619 = m.ExcPending
	if v1619 != 0 {
		goto L14
	} else {
		goto L419
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1603 = m.ExcPending
	if v1603 != 0 {
		goto L14
	} else {
		goto L415
	}
L20:
	;
	if v83 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v83)+24))
	v88 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
	if v88 == int64(1346862204) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	goto L23
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1587 = m.ExcPending
	if v1587 != 0 {
		goto L14
	} else {
		goto L411
	}
L24:
	;
	if v90 == int32(0) {
		goto L19
	} else {
		goto L28
	}
L25:
	;
	v90 = v86
	goto L27
L26:
	;
	v90 = int32(0)
	goto L27
L27:
	;
	goto L24
L28:
	;
	v96 = F_shm_toc_lookup(m, v90, int64(-65535), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L14
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, _consts[88])) = v96
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v96)+40))
	*(*int32)(unsafe.Add(mBase, _consts[89])) = v100
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v96)+44))
	*(*int32)(unsafe.Add(mBase, _consts[90])) = v103
	F_before_shmem_exit(m, int32(296), v83)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L14
	} else {
		goto L30
	}
L30:
	;
	v110 = F_shm_toc_lookup(m, v90, int64(-65534), int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L14
	} else {
		goto L31
	}
L31:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _consts[55]))
	v116 = v110 + v113<<(uint(int32(14))%32)
	v118 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	F_shm_mq_set_sender(m, v116, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L14
	} else {
		goto L32
	}
L32:
	;
	v121 = F_shm_mq_attach(m, v116, v83)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L14
	} else {
		goto L33
	}
L33:
	;
	F_pq_redirect_to_shm_mq(m, v83, v121)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L14
	} else {
		goto L34
	}
L34:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v96)+40))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v96)+44))
	*(*int32)(unsafe.Add(mBase, _consts[92])) = v126
	*(*int32)(unsafe.Add(mBase, _consts[93])) = v125
	goto L35
L35:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v96)+40))
	v133 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v96)+36))
	v136 = *(*int32)(unsafe.Add(mBase, _consts[94]))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	v140 = base.I32_div_s(v134-v137, int32(640))
	v142 = base.I32_rem_s(v140, int32(16))
	v147 = v133 + v142<<(uint(int32(7))%32) + int32(23296)
	v149 = F_LWLockAcquire(m, v147, int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L14
	} else {
		goto L36
	}
L36:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v134)+44))
	if v151 != v131 {
		v174 = v2
		goto L37
	} else {
		goto L38
	}
L37:
	;
	F_LWLockRelease(m, v147)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L14
	} else {
		goto L43
	}
L38:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v134)+616))
	if v153 != v134 {
		v174 = v2
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v156 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	*(*int32)(unsafe.Add(mBase, uint32(v156)+616)) = v134
	v159 = v134 + int32(620)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v134)+624))
	if v160 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v134)+624)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v134)+620)) = v159
	goto L42
L41:
	;
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v156)+632)) = v159
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	*(*int32)(unsafe.Add(mBase, uint32(v156)+628)) = v166
	v169 = v156 + int32(628)
	*(*int32)(unsafe.Add(mBase, uint32(v166)+4)) = v169
	*(*int32)(unsafe.Add(mBase, uint32(v159))) = v169
	v174 = int32(1)
	goto L37
L43:
	;
	if v174 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v178 = *(*int64)(unsafe.Add(mBase, uint32(v96)+48))
	v180 = *(*int64)(unsafe.Add(mBase, uint32(v96)+56))
	*(*int64)(unsafe.Add(mBase, _consts[95])) = v180
	*(*int64)(unsafe.Add(mBase, _consts[96])) = v178
	v186 = F_shm_toc_lookup(m, v90, int64(-65527), int32(0))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L14
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	m.G0 = v19 + int32(32)
	return
L47:
	;
	if v186&int32(3) == int32(0) {
		v211 = v186
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v247 = v244 + v186 + int32(1)
	v248 = int32(159252)
	v251 = int32(*(*uint8)(unsafe.Add(mBase, _consts[97])))
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
	if v252 == int32(0) {
		v271 = v251
		v272 = v252
		goto L67
	} else {
		goto L68
	}
L49:
	;
	v244 = v236 - v186
	goto L48
L50:
	;
	v215 = v211
	goto L59
L51:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
	if v195 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v244 = int32(0)
	goto L48
L53:
	;
	goto L54
L54:
	;
	v200 = v186
	goto L55
L55:
	;
	v204 = v200 + int32(1)
	if v204&int32(3) == int32(0) {
		v211 = v204
		goto L50
	} else {
		goto L57
	}
L56:
	;
	v236 = v204
	goto L49
L57:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	if v209 != 0 {
		v200 = v204
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
	v224 = int32(-2139062144)
	if (int32(16843008)-v221|v221)&v224 == v224 {
		v215 = v215 + int32(4)
		goto L59
	} else {
		goto L61
	}
L60:
	;
	v230 = v215
	goto L62
L61:
	;
	goto L60
L62:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230))))
	if v234 != 0 {
		v230 = v230 + int32(1)
		goto L62
	} else {
		goto L64
	}
L63:
	;
	v236 = v230
	goto L49
L64:
	;
	goto L63
L65:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	*(*int32)(unsafe.Add(mBase, _consts[98])) = v430
	v433 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	*(*int32)(unsafe.Add(mBase, uint32(v433)+64)) = v430
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+32)))
	F_SetSessionAuthorization(m, v435, v436)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L14
	} else {
		goto L131
	}
L66:
	;
	if v272-v271 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L67:
	;
	goto L66
L68:
	;
	if v251 != v252 {
		v271 = v251
		v272 = v252
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v256 = v186
	v257 = v248
	goto L70
L70:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257)+1)))
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256)+1)))
	if v261 == int32(0) {
		v271 = v260
		v272 = v261
		goto L67
	} else {
		goto L72
	}
L71:
	;
	v271 = v260
	v272 = v261
	goto L67
L72:
	;
	v264 = int32(1)
	if v260 == v261 {
		v256 = v256 + v264
		v257 = v257 + v264
		goto L70
	} else {
		goto L73
	}
L73:
	;
	goto L71
L74:
	;
	v276 = int32(274390)
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247))))
	v280 = int32(*(*uint8)(unsafe.Add(mBase, _consts[99])))
	if v280 == int32(0) {
		v299 = v279
		v300 = v280
		goto L78
	} else {
		goto L79
	}
L75:
	;
	goto L76
L76:
	;
	v426 = F_load_external_function(m, v186, v247, int32(1), int32(0))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L14
	} else {
		goto L130
	}
L77:
	;
	if v300-v299 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L78:
	;
	goto L77
L79:
	;
	if v279 != v280 {
		v299 = v279
		v300 = v280
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v284 = v276
	v285 = v247
	goto L81
L81:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285)+1)))
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284)+1)))
	if v289 == int32(0) {
		v299 = v288
		v300 = v289
		goto L78
	} else {
		goto L83
	}
L82:
	;
	v299 = v288
	v300 = v289
	goto L78
L83:
	;
	v292 = int32(1)
	if v288 == v289 {
		v284 = v284 + v292
		v285 = v285 + v292
		goto L81
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	v305 = *(*int32)(unsafe.Add(mBase, _consts[100]))
	v428 = v305
	goto L65
L86:
	;
	goto L87
L87:
	;
	v306 = int32(274197)
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247))))
	v310 = int32(*(*uint8)(unsafe.Add(mBase, _consts[101])))
	if v310 == int32(0) {
		v329 = v309
		v330 = v310
		goto L89
	} else {
		goto L90
	}
L88:
	;
	if v330-v329 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L89:
	;
	goto L88
L90:
	;
	if v309 != v310 {
		v329 = v309
		v330 = v310
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v314 = v306
	v315 = v247
	goto L92
L92:
	;
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+1)))
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314)+1)))
	if v319 == int32(0) {
		v329 = v318
		v330 = v319
		goto L89
	} else {
		goto L94
	}
L93:
	;
	v329 = v318
	v330 = v319
	goto L89
L94:
	;
	v322 = int32(1)
	if v318 == v319 {
		v314 = v314 + v322
		v315 = v315 + v322
		goto L92
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	v335 = *(*int32)(unsafe.Add(mBase, _consts[102]))
	v428 = v335
	goto L65
L97:
	;
	goto L98
L98:
	;
	v336 = int32(274221)
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247))))
	v340 = int32(*(*uint8)(unsafe.Add(mBase, _consts[103])))
	if v340 == int32(0) {
		v359 = v339
		v360 = v340
		goto L100
	} else {
		goto L101
	}
L99:
	;
	if v360-v359 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L100:
	;
	goto L99
L101:
	;
	if v339 != v340 {
		v359 = v339
		v360 = v340
		goto L100
	} else {
		goto L102
	}
L102:
	;
	v344 = v336
	v345 = v247
	goto L103
L103:
	;
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345)+1)))
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344)+1)))
	if v349 == int32(0) {
		v359 = v348
		v360 = v349
		goto L100
	} else {
		goto L105
	}
L104:
	;
	v359 = v348
	v360 = v349
	goto L100
L105:
	;
	v352 = int32(1)
	if v348 == v349 {
		v344 = v344 + v352
		v345 = v345 + v352
		goto L103
	} else {
		goto L106
	}
L106:
	;
	goto L104
L107:
	;
	v365 = *(*int32)(unsafe.Add(mBase, _consts[104]))
	v428 = v365
	goto L65
L108:
	;
	goto L109
L109:
	;
	v366 = int32(274247)
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247))))
	v370 = int32(*(*uint8)(unsafe.Add(mBase, _consts[105])))
	if v370 == int32(0) {
		v389 = v369
		v390 = v370
		goto L111
	} else {
		goto L112
	}
L110:
	;
	if v390-v389 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L111:
	;
	goto L110
L112:
	;
	if v369 != v370 {
		v389 = v369
		v390 = v370
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v374 = v366
	v375 = v247
	goto L114
L114:
	;
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v375)+1)))
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374)+1)))
	if v379 == int32(0) {
		v389 = v378
		v390 = v379
		goto L111
	} else {
		goto L116
	}
L115:
	;
	v389 = v378
	v390 = v379
	goto L111
L116:
	;
	v382 = int32(1)
	if v378 == v379 {
		v374 = v374 + v382
		v375 = v375 + v382
		goto L114
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	v395 = *(*int32)(unsafe.Add(mBase, _consts[106]))
	v428 = v395
	goto L65
L119:
	;
	goto L120
L120:
	;
	v396 = int32(274176)
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247))))
	v400 = int32(*(*uint8)(unsafe.Add(mBase, _consts[107])))
	if v400 == int32(0) {
		v419 = v399
		v420 = v400
		goto L122
	} else {
		goto L123
	}
L121:
	;
	if v420-v419 != 0 {
		goto L18
	} else {
		goto L129
	}
L122:
	;
	goto L121
L123:
	;
	if v399 != v400 {
		v419 = v399
		v420 = v400
		goto L122
	} else {
		goto L124
	}
L124:
	;
	v404 = v396
	v405 = v247
	goto L125
L125:
	;
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v405)+1)))
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404)+1)))
	if v409 == int32(0) {
		v419 = v408
		v420 = v409
		goto L122
	} else {
		goto L127
	}
L126:
	;
	v419 = v408
	v420 = v409
	goto L122
L127:
	;
	v412 = int32(1)
	if v408 == v409 {
		v404 = v404 + v412
		v405 = v405 + v412
		goto L125
	} else {
		goto L128
	}
L128:
	;
	goto L126
L129:
	;
	v423 = *(*int32)(unsafe.Add(mBase, _consts[108]))
	v428 = v423
	goto L65
L130:
	;
	v428 = v426
	goto L65
L131:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v96)+12))
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+33)))
	F_SetCurrentRoleId(m, v439, v440)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L14
	} else {
		goto L132
	}
L132:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	F_BackgroundWorkerInitializeConnectionByOid(m, v443, v444, int32(3))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L14
	} else {
		goto L133
	}
L133:
	;
	v449 = *(*int32)(unsafe.Add(mBase, _consts[109]))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v449)+4))
	goto L134
L134:
	;
	v451 = F_SetClientEncoding(m, v450)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L14
	} else {
		goto L135
	}
L135:
	;
	if v451 < int32(0) {
		goto L17
	} else {
		goto L136
	}
L136:
	;
	v457 = F_shm_toc_lookup(m, v90, int64(-65533), int32(0))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L14
	} else {
		goto L137
	}
L137:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L14
	} else {
		goto L138
	}
L138:
	;
	v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v457))))
	if v461 != 0 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v463 = v457
	goto L142
L140:
	;
	goto L141
L141:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L14
	} else {
		goto L163
	}
L142:
	;
	v478 = F_internal_load_library(m, v463)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L14
	} else {
		goto L144
	}
L143:
	;
	goto L141
L144:
	;
	if v463&int32(3) == int32(0) {
		v503 = v463
		goto L147
	} else {
		goto L148
	}
L145:
	;
	v539 = v536 + v463 + int32(1)
	v540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v539))))
	if v540 != 0 {
		v463 = v539
		goto L142
	} else {
		goto L162
	}
L146:
	;
	v536 = v528 - v463
	goto L145
L147:
	;
	v507 = v503
	goto L156
L148:
	;
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v463))))
	if v487 == int32(0) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v536 = int32(0)
	goto L145
L150:
	;
	goto L151
L151:
	;
	v492 = v463
	goto L152
L152:
	;
	v496 = v492 + int32(1)
	if v496&int32(3) == int32(0) {
		v503 = v496
		goto L147
	} else {
		goto L154
	}
L153:
	;
	v528 = v496
	goto L146
L154:
	;
	v501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496))))
	if v501 != 0 {
		v492 = v496
		goto L152
	} else {
		goto L155
	}
L155:
	;
	goto L153
L156:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v507)))
	v516 = int32(-2139062144)
	if (int32(16843008)-v513|v513)&v516 == v516 {
		v507 = v507 + int32(4)
		goto L156
	} else {
		goto L158
	}
L157:
	;
	v522 = v507
	goto L159
L158:
	;
	goto L157
L159:
	;
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v522))))
	if v526 != 0 {
		v522 = v522 + int32(1)
		goto L159
	} else {
		goto L161
	}
L160:
	;
	v528 = v522
	goto L146
L161:
	;
	goto L160
L162:
	;
	goto L143
L163:
	;
	v561 = F_shm_toc_lookup(m, v90, int64(-65528), int32(0))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L14
	} else {
		goto L164
	}
L164:
	;
	F_StartTransaction(m)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L14
	} else {
		goto L165
	}
L165:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v561)))
	*(*int32)(unsafe.Add(mBase, _consts[110])) = v566
	v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v561)+4)))
	*(*uint8)(unsafe.Add(mBase, _consts[111])) = uint8(v569)
	v572 = *(*int64)(unsafe.Add(mBase, uint32(v561)+8))
	*(*int64)(unsafe.Add(mBase, _consts[35])) = v572
	v575 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	v576 = *(*int64)(unsafe.Add(mBase, uint32(v561)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v575))) = v576
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v561)+24))
	*(*int32)(unsafe.Add(mBase, _consts[112])) = v579
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v561)+28))
	*(*int32)(unsafe.Add(mBase, _consts[38])) = v561 + int32(32)
	*(*int32)(unsafe.Add(mBase, _consts[36])) = v581
	*(*int32)(unsafe.Add(mBase, uint32(v575)+24)) = int32(5)
	v592 = F_shm_toc_lookup(m, v90, int64(-65525), int32(0))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L14
	} else {
		goto L166
	}
L166:
	;
	v594 = m.G0
	v596 = v594 - int32(48)
	m.G0 = v596
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v592)+8))
	if v598 != 0 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v600 = v592
	goto L170
L168:
	;
	goto L169
L169:
	;
	m.G0 = v596 + int32(48)
	v663 = F_shm_toc_lookup(m, v90, int64(-65523), int32(0))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L14
	} else {
		goto L178
	}
L170:
	;
	v616 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	if v616 == int32(0) {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	goto L169
L172:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v596)+16)) = int64(68719476748)
	v622 = *(*int32)(unsafe.Add(mBase, _consts[76]))
	*(*int32)(unsafe.Add(mBase, uint32(v596)+40)) = v622
	v628 = F_hash_create(m, int32(318475), int32(16), v596, int32(1064))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L14
	} else {
		goto L175
	}
L173:
	;
	v631 = v616
	goto L174
L174:
	;
	v633 = F_hash_search(m, v631, v600, int32(1), v596)
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L14
	} else {
		goto L176
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, _consts[113])) = v628
	v631 = v628
	goto L174
L176:
	;
	v635 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v633)+12)) = uint8(v635)
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v600+int32(20))))
	if v641 != 0 {
		v600 = v600 + int32(12)
		goto L170
	} else {
		goto L177
	}
L177:
	;
	goto L171
L178:
	;
	v666 = *(*int32)(unsafe.Add(mBase, _consts[114]))
	if v666 != 0 {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	goto L189
L180:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L14
	} else {
		goto L185
	}
L181:
	;
	v668 = *(*int32)(unsafe.Add(mBase, _consts[115]))
	if v668 != 0 {
		goto L180
	} else {
		goto L182
	}
L182:
	;
	v670 = *(*int32)(unsafe.Add(mBase, _consts[116]))
	if v670 != 0 {
		goto L180
	} else {
		goto L183
	}
L183:
	;
	v672 = *(*int32)(unsafe.Add(mBase, _consts[117]))
	if v672 == int32(0) {
		goto L179
	} else {
		goto L184
	}
L184:
	;
	goto L180
L185:
	;
	F_errmsg_internal(m, int32(153917), int32(0))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L14
	} else {
		goto L186
	}
L186:
	;
	F_errfinish(m, int32(487689), int32(749), int32(235508))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L14
	} else {
		goto L187
	}
L187:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L188:
	;
	v693 = int32(524)
	goto L193
L189:
	;
	v690 = F__emscripten_memcpy_bulkmem(m, int32(4459024), v663, int32(524))
	mBase = m.M
	goto L191
L191:
	;
	goto L188
L192:
	;
	v700 = F_shm_toc_lookup(m, v90, int64(-65524), int32(0))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L14
	} else {
		goto L196
	}
L193:
	;
	v696 = F__emscripten_memcpy_bulkmem(m, int32(4460072), v663+v693, v693)
	mBase = m.M
	goto L195
L195:
	;
	goto L192
L196:
	;
	v702 = int32(0)
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v700)))
	*(*int32)(unsafe.Add(mBase, _consts[118])) = v704
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v700)+4))
	*(*int32)(unsafe.Add(mBase, _consts[43])) = v707
	v709 = int32(4470752)
	v710 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v713 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v713
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v700)+8))
	if v702 < v715 {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v721 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v724 = v702
	v725 = v721
	goto L200
L198:
	;
	goto L199
L199:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v710
	v770 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v770)+28))
	goto L204
L200:
	;
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v700+int32(12)+v724<<(uint(int32(2))%32))))
	v743 = F_lappend_oid(m, v725, v742)
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L14
	} else {
		goto L202
	}
L201:
	;
	goto L199
L202:
	;
	*(*int32)(unsafe.Add(mBase, _consts[44])) = v743
	v747 = v724 + int32(1)
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v700)+8))
	if v747 < v748 {
		v724 = v747
		v725 = v743
		goto L200
	} else {
		goto L203
	}
L203:
	;
	goto L201
L204:
	;
	*(*int32)(unsafe.Add(mBase, _consts[119])) = v771
	v775 = F_shm_toc_lookup(m, v90, int64(-65531), int32(0))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L14
	} else {
		goto L205
	}
L205:
	;
	v777 = int32(0)
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v775)))
	if v778 <= v777 {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v841 = F_shm_toc_lookup(m, v90, int64(-65526), int32(0))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L14
	} else {
		goto L218
	}
L207:
	;
	v785 = v777
	goto L208
L208:
	;
	v801 = v775 + int32(4) + v785<<(uint(int32(3))%32)
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v801)))
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v801)+4))
	v804 = F_GetComboCommandId(m, v802, v803)
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L14
	} else {
		goto L210
	}
L209:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L14
	} else {
		goto L215
	}
L210:
	;
	if v804 == v785 {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v808 = v785 + int32(1)
	if v778 != v808 {
		v785 = v808
		goto L208
	} else {
		goto L214
	}
L212:
	;
	goto L213
L213:
	;
	goto L209
L214:
	;
	goto L206
L215:
	;
	F_errmsg_internal(m, int32(172472), int32(0))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L14
	} else {
		goto L216
	}
L216:
	;
	F_errfinish(m, int32(492021), int32(362), int32(348771))
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L14
	} else {
		goto L217
	}
L217:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L218:
	;
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v841)))
	v844 = int32(4470752)
	v845 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v848 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v848
	v850 = F_dsm_attach(m, v843)
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L14
	} else {
		goto L219
	}
L219:
	;
	if v850 == int32(0) {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L14
	} else {
		goto L223
	}
L221:
	;
	goto L222
L222:
	;
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v850)+24))
	v870 = *(*int64)(unsafe.Add(mBase, uint32(v868)))
	if v870 == int64(2880502729) {
		goto L227
	} else {
		goto L228
	}
L223:
	;
	F_errmsg_internal(m, int32(94019), int32(0))
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L14
	} else {
		goto L224
	}
L224:
	;
	F_errfinish(m, int32(488783), int32(169), int32(267111))
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L14
	} else {
		goto L225
	}
L225:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L226:
	;
	v875 = F_shm_toc_lookup(m, v872, int64(-65535), int32(0))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L14
	} else {
		goto L230
	}
L227:
	;
	v872 = v868
	goto L229
L228:
	;
	v872 = int32(0)
	goto L229
L229:
	;
	goto L226
L230:
	;
	v877 = F_dsa_attach_in_place(m, v875, v850)
	mBase = m.M
	v878 = m.ExcPending
	if v878 != 0 {
		goto L14
	} else {
		goto L231
	}
L231:
	;
	v879 = int32(4365212)
	v880 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	*(*int32)(unsafe.Add(mBase, uint32(v880))) = v850
	v883 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	*(*int32)(unsafe.Add(mBase, uint32(v883)+4)) = v877
	v887 = F_shm_toc_lookup(m, v872, int64(-65534), int32(0))
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L14
	} else {
		goto L232
	}
L232:
	;
	v889 = int32(4470752)
	v890 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v893 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v893
	v896 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v896)+4))
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v887)))
	v900 = F_dshash_attach(m, v897, int32(1737548), v899, v897)
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L14
	} else {
		goto L233
	}
L233:
	;
	v903 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v903)+4))
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v887)+4))
	v908 = F_dshash_attach(m, v904, int32(1737572), v906, int32(0))
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L14
	} else {
		goto L234
	}
L234:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v890
	v913 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v913)))
	F_on_dsm_detach(m, v914, int32(1622), v887)
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L14
	} else {
		goto L235
	}
L235:
	;
	v919 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	*(*int32)(unsafe.Add(mBase, uint32(v919)+16)) = v908
	*(*int32)(unsafe.Add(mBase, uint32(v919)+12)) = v900
	*(*int32)(unsafe.Add(mBase, uint32(v919)+8)) = v887
	F_dsm_pin_mapping(m, v850)
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L14
	} else {
		goto L236
	}
L236:
	;
	F_dsa_pin_mapping(m, v877)
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L14
	} else {
		goto L237
	}
L237:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v845
	v931 = F_shm_toc_lookup(m, v90, int64(-65529), int32(0))
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L14
	} else {
		goto L238
	}
L238:
	;
	v935 = F_shm_toc_lookup(m, v90, int64(-65530), int32(1))
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L14
	} else {
		goto L239
	}
L239:
	;
	v937 = F_RestoreSnapshot(m, v931)
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		goto L14
	} else {
		goto L240
	}
L240:
	;
	if v935 != 0 {
		goto L241
	} else {
		goto L242
	}
L241:
	;
	v939 = F_RestoreSnapshot(m, v935)
	mBase = m.M
	v940 = m.ExcPending
	if v940 != 0 {
		goto L14
	} else {
		goto L244
	}
L242:
	;
	v941 = v937
	goto L243
L243:
	;
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v96)+36))
	F_RestoreTransactionSnapshot(m, v941, v942)
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		goto L14
	} else {
		goto L245
	}
L244:
	;
	v941 = v939
	goto L243
L245:
	;
	F_PushActiveSnapshot(m, v937)
	mBase = m.M
	v946 = m.ExcPending
	if v946 != 0 {
		goto L14
	} else {
		goto L246
	}
L246:
	;
	F_InvalidateSystemCachesExtended(m)
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L14
	} else {
		goto L247
	}
L247:
	;
	v951 = F_shm_toc_lookup(m, v90, int64(-65532), int32(0))
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L14
	} else {
		goto L248
	}
L248:
	;
	v953 = m.G0
	v955 = v953 - int32(32)
	m.G0 = v955
	v958 = *(*int32)(unsafe.Add(mBase, _consts[121]))
	if v958 == int32(0) {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(v951)))
	*(*int32)(unsafe.Add(mBase, uint32(v955)+20)) = int32(1658)
	v1122 = int32(4463656)
	v1123 = *(*int32)(unsafe.Add(mBase, _consts[77]))
	*(*int32)(unsafe.Add(mBase, _consts[77])) = v955 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v955)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v955)+16)) = v1123
	v1132 = v951 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v955)+28)) = v1132
	v1135 = v1119 + v1132
	if base.Ui32(v1132) < base.Ui32(v1135) {
		goto L324
	} else {
		goto L325
	}
L250:
	;
	if v958 == int32(4468780) {
		goto L249
	} else {
		goto L251
	}
L251:
	;
	v964 = v958
	goto L252
L252:
	;
	v980 = v964 + int32(4)
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v980)))
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v964-int32(60))))
	if base.Ui32(v984) < base.Ui32(int32(2)) {
		goto L254
	} else {
		goto L255
	}
L253:
	;
	goto L249
L254:
	;
	if v981 != int32(4468780) {
		v964 = v981
		goto L252
	} else {
		goto L317
	}
L255:
	;
	v988 = v964 - int32(32)
	v989 = *(*int32)(unsafe.Add(mBase, uint32(v988)))
	if v989 == int32(0) {
		goto L254
	} else {
		goto L256
	}
L256:
	;
	v993 = v964 - int32(4)
	v994 = *(*int32)(unsafe.Add(mBase, uint32(v993)))
	if v994 != 0 {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	F_pfree(m, v994)
	mBase = m.M
	v996 = m.ExcPending
	if v996 != 0 {
		goto L14
	} else {
		goto L260
	}
L258:
	;
	goto L259
L259:
	;
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v964)+16))
	if v997 != 0 {
		goto L261
	} else {
		goto L262
	}
L260:
	;
	goto L259
L261:
	;
	F_pfree(m, v997)
	mBase = m.M
	v999 = m.ExcPending
	if v999 != 0 {
		goto L14
	} else {
		goto L264
	}
L262:
	;
	goto L263
L263:
	;
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(v964)+20))
	if v1000 != 0 {
		goto L265
	} else {
		goto L266
	}
L264:
	;
	goto L263
L265:
	;
	F_pfree(m, v1000)
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L14
	} else {
		goto L268
	}
L266:
	;
	goto L267
L267:
	;
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(v964-int32(40))))
	switch v1005 {
	case 0:
		goto L275
	case 1:
		goto L274
	case 2:
		goto L273
	case 3:
		goto L272
	case 4:
		goto L271
	default:
		goto L269
	}
L268:
	;
	goto L267
L269:
	;
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(v988)))
	if v1051 != 0 {
		goto L295
	} else {
		goto L296
	}
L270:
	;
	F_pfree(m, v1045)
	mBase = m.M
	v1048 = m.ExcPending
	if v1048 != 0 {
		goto L14
	} else {
		goto L294
	}
L271:
	;
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v964)+56))
	if v1040 == int32(0) {
		goto L269
	} else {
		goto L292
	}
L272:
	;
	v1022 = v964 + int32(28)
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(v1022)))
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(v1023)))
	if v1024 != 0 {
		goto L282
	} else {
		goto L283
	}
L273:
	;
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(v964)+80))
	if v1016 == int32(0) {
		goto L269
	} else {
		goto L280
	}
L274:
	;
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(v964)+60))
	if v1011 == int32(0) {
		goto L269
	} else {
		goto L278
	}
L275:
	;
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v964)+52))
	if v1006 == int32(0) {
		goto L269
	} else {
		goto L276
	}
L276:
	;
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v993)))
	if v1006 != v1009 {
		v1045 = v1006
		goto L270
	} else {
		goto L277
	}
L277:
	;
	goto L269
L278:
	;
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(v993)))
	if v1011 != v1014 {
		v1045 = v1011
		goto L270
	} else {
		goto L279
	}
L279:
	;
	goto L269
L280:
	;
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(v993)))
	if v1016 != v1019 {
		v1045 = v1016
		goto L270
	} else {
		goto L281
	}
L281:
	;
	goto L269
L282:
	;
	F_pfree(m, v1024)
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L14
	} else {
		goto L285
	}
L283:
	;
	goto L284
L284:
	;
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(v964)+48))
	if v1027 == int32(0) {
		goto L286
	} else {
		goto L287
	}
L285:
	;
	goto L284
L286:
	;
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(v964)+52))
	if v1035 == int32(0) {
		goto L269
	} else {
		goto L290
	}
L287:
	;
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v1022)))
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(v1030)))
	if v1027 == v1031 {
		goto L286
	} else {
		goto L288
	}
L288:
	;
	F_pfree(m, v1027)
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L14
	} else {
		goto L289
	}
L289:
	;
	goto L286
L290:
	;
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v993)))
	if v1035 != v1038 {
		v1045 = v1035
		goto L270
	} else {
		goto L291
	}
L291:
	;
	goto L269
L292:
	;
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v993)))
	if v1040 == v1043 {
		goto L269
	} else {
		goto L293
	}
L293:
	;
	v1045 = v1040
	goto L270
L294:
	;
	goto L269
L295:
	;
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(v964)))
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v980)))
	*(*int32)(unsafe.Add(mBase, uint32(v1052)+4)) = v1053
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v964)))
	*(*int32)(unsafe.Add(mBase, uint32(v1053))) = v1055
	goto L297
L296:
	;
	goto L297
L297:
	;
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(v964-int32(8))))
	if v1060 != 0 {
		goto L298
	} else {
		goto L299
	}
L298:
	;
	v1065 = int32(4468796)
	goto L303
L299:
	;
	goto L300
L300:
	;
	v1076 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v964-int32(36)))))
	if v1076&int32(4) != 0 {
		goto L307
	} else {
		goto L308
	}
L301:
	;
	goto L300
L302:
	;
	goto L301
L303:
	;
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(v1065)))
	if v1068 == int32(0) {
		goto L302
	} else {
		goto L305
	}
L304:
	;
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v1068)))
	*(*int32)(unsafe.Add(mBase, uint32(v1065))) = v1072
	goto L302
L305:
	;
	if v1068 != v964+int32(8) {
		v1065 = v1068
		goto L303
	} else {
		goto L306
	}
L306:
	;
	goto L304
L307:
	;
	v1083 = int32(4468788)
	goto L312
L308:
	;
	goto L309
L309:
	;
	F_InitializeOneGUCOption(m, v964+int32(-64))
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L14
	} else {
		goto L316
	}
L310:
	;
	goto L309
L311:
	;
	goto L310
L312:
	;
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v1083)))
	if v1086 == int32(0) {
		goto L311
	} else {
		goto L314
	}
L313:
	;
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(v1086)))
	*(*int32)(unsafe.Add(mBase, uint32(v1083))) = v1090
	goto L311
L314:
	;
	if v1086 != v964+int32(12) {
		v1083 = v1086
		goto L312
	} else {
		goto L315
	}
L315:
	;
	goto L313
L316:
	;
	goto L254
L317:
	;
	goto L253
L318:
	;
	v1321 = *(*int32)(unsafe.Add(mBase, uint32(v96)+16))
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(v96)+28))
	*(*int32)(unsafe.Add(mBase, _consts[122])) = v1322
	*(*int32)(unsafe.Add(mBase, _consts[4])) = v1321
	goto L371
L319:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1308 = m.ExcPending
	if v1308 != 0 {
		goto L14
	} else {
		goto L367
	}
L320:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1295 = m.ExcPending
	if v1295 != 0 {
		goto L14
	} else {
		goto L364
	}
L321:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1282 = m.ExcPending
	if v1282 != 0 {
		goto L14
	} else {
		goto L361
	}
L322:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1269 = m.ExcPending
	if v1269 != 0 {
		goto L14
	} else {
		goto L358
	}
L323:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1256 = m.ExcPending
	if v1256 != 0 {
		goto L14
	} else {
		goto L355
	}
L324:
	;
	goto L327
L325:
	;
	v1248 = v1123
	goto L326
L326:
	;
	*(*int32)(unsafe.Add(mBase, _consts[77])) = v1248
	m.G0 = v955 + int32(32)
	goto L318
L327:
	;
	v1155 = F_read_gucstate(m, v955+int32(28), v1135)
	mBase = m.M
	v1156 = m.ExcPending
	if v1156 != 0 {
		goto L14
	} else {
		goto L329
	}
L328:
	;
	v1231 = *(*int32)(unsafe.Add(mBase, uint32(v955)+16))
	v1248 = v1231
	goto L326
L329:
	;
	v1159 = F_read_gucstate(m, v955+int32(28), v1135)
	mBase = m.M
	v1160 = m.ExcPending
	if v1160 != 0 {
		goto L14
	} else {
		goto L330
	}
L330:
	;
	v1163 = F_read_gucstate(m, v955+int32(28), v1135)
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L14
	} else {
		goto L331
	}
L331:
	;
	v1165 = *(*int32)(unsafe.Add(mBase, uint32(v955)+28))
	v1166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1163))))
	if v1166 == int32(0) {
		goto L333
	} else {
		goto L334
	}
L332:
	;
	v1177 = v1174 + int32(4)
	if base.Ui32(v1135) < base.Ui32(v1177) {
		goto L322
	} else {
		goto L337
	}
L333:
	;
	v1174 = v1165
	v1175 = int32(0)
	goto L332
L334:
	;
	goto L335
L335:
	;
	v1171 = v1165 + int32(4)
	if base.Ui32(v1135) < base.Ui32(v1171) {
		goto L323
	} else {
		goto L336
	}
L336:
	;
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(v1165)))
	v1174 = v1171
	v1175 = v1173
	goto L332
L337:
	;
	v1180 = v1174 + int32(8)
	if base.Ui32(v1135) < base.Ui32(v1180) {
		goto L321
	} else {
		goto L338
	}
L338:
	;
	v1183 = v1174 + int32(12)
	if base.Ui32(v1135) < base.Ui32(v1183) {
		goto L320
	} else {
		goto L339
	}
L339:
	;
	v1185 = *(*int32)(unsafe.Add(mBase, uint32(v1174)))
	v1186 = *(*int32)(unsafe.Add(mBase, uint32(v1180)))
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(v1177)))
	*(*int32)(unsafe.Add(mBase, uint32(v955)+12)) = v1159
	*(*int32)(unsafe.Add(mBase, uint32(v955)+8)) = v1155
	*(*int32)(unsafe.Add(mBase, uint32(v955)+28)) = v1183
	*(*int32)(unsafe.Add(mBase, uint32(v955)+24)) = v955 + int32(8)
	v1194 = int32(0)
	v1196 = int32(1)
	v1199 = F_set_config_with_handle(m, v1155, v1194, v1159, v1187, v1185, v1186, v1194, v1196, int32(21), v1196)
	mBase = m.M
	v1200 = m.ExcPending
	if v1200 != 0 {
		goto L14
	} else {
		goto L340
	}
L340:
	;
	if v1199 <= int32(0) {
		goto L319
	} else {
		goto L341
	}
L341:
	;
	v1203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1163))))
	if v1203 == int32(0) {
		goto L342
	} else {
		goto L343
	}
L342:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v955)+24)) = int32(0)
	if base.Ui32(v1183) < base.Ui32(v1135) {
		goto L327
	} else {
		goto L354
	}
L343:
	;
	v1211 = int32(*(*uint8)(unsafe.Add(mBase, _consts[123])))
	if v1211 != 0 {
		goto L344
	} else {
		goto L345
	}
L344:
	;
	v1212 = int32(12)
	goto L346
L345:
	;
	v1212 = int32(15)
	goto L346
L346:
	;
	v1213 = F_find_option(m, v1155, int32(1), int32(0), v1212)
	mBase = m.M
	v1214 = m.ExcPending
	if v1214 != 0 {
		goto L14
	} else {
		goto L347
	}
L347:
	;
	if v1213 == int32(0) {
		goto L342
	} else {
		goto L348
	}
L348:
	;
	v1217 = F_guc_strdup(m, v1212, v1163)
	mBase = m.M
	v1218 = m.ExcPending
	if v1218 != 0 {
		goto L14
	} else {
		goto L349
	}
L349:
	;
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(v1213)+84))
	if v1219 != 0 {
		goto L350
	} else {
		goto L351
	}
L350:
	;
	F_pfree(m, v1219)
	mBase = m.M
	v1221 = m.ExcPending
	if v1221 != 0 {
		goto L14
	} else {
		goto L353
	}
L351:
	;
	goto L352
L352:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1213)+88)) = v1175
	*(*int32)(unsafe.Add(mBase, uint32(v1213)+84)) = v1217
	goto L342
L353:
	;
	goto L352
L354:
	;
	goto L328
L355:
	;
	F_errmsg_internal(m, int32(348126), int32(0))
	mBase = m.M
	v1260 = m.ExcPending
	if v1260 != 0 {
		goto L14
	} else {
		goto L356
	}
L356:
	;
	F_errfinish(m, int32(492061), int32(6168), int32(17667))
	mBase = m.M
	v1265 = m.ExcPending
	if v1265 != 0 {
		goto L14
	} else {
		goto L357
	}
L357:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L358:
	;
	F_errmsg_internal(m, int32(348126), int32(0))
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
		goto L14
	} else {
		goto L359
	}
L359:
	;
	F_errfinish(m, int32(492061), int32(6168), int32(17667))
	mBase = m.M
	v1278 = m.ExcPending
	if v1278 != 0 {
		goto L14
	} else {
		goto L360
	}
L360:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L361:
	;
	F_errmsg_internal(m, int32(348126), int32(0))
	mBase = m.M
	v1286 = m.ExcPending
	if v1286 != 0 {
		goto L14
	} else {
		goto L362
	}
L362:
	;
	F_errfinish(m, int32(492061), int32(6168), int32(17667))
	mBase = m.M
	v1291 = m.ExcPending
	if v1291 != 0 {
		goto L14
	} else {
		goto L363
	}
L363:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L364:
	;
	F_errmsg_internal(m, int32(348126), int32(0))
	mBase = m.M
	v1299 = m.ExcPending
	if v1299 != 0 {
		goto L14
	} else {
		goto L365
	}
L365:
	;
	F_errfinish(m, int32(492061), int32(6168), int32(17667))
	mBase = m.M
	v1304 = m.ExcPending
	if v1304 != 0 {
		goto L14
	} else {
		goto L366
	}
L366:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L367:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v1311 = m.ExcPending
	if v1311 != 0 {
		goto L14
	} else {
		goto L368
	}
L368:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v955))) = v1155
	F_errmsg(m, int32(104913), v955)
	mBase = m.M
	v1315 = m.ExcPending
	if v1315 != 0 {
		goto L14
	} else {
		goto L369
	}
L369:
	;
	F_errfinish(m, int32(492061), int32(6353), int32(348792))
	mBase = m.M
	v1320 = m.ExcPending
	if v1320 != 0 {
		goto L14
	} else {
		goto L370
	}
L370:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L371:
	;
	v1327 = *(*int32)(unsafe.Add(mBase, uint32(v96)+20))
	v1329 = *(*int32)(unsafe.Add(mBase, uint32(v96)+24))
	*(*int32)(unsafe.Add(mBase, _consts[124])) = v1329
	*(*int32)(unsafe.Add(mBase, _consts[125])) = v1327
	v1334 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[126])) = uint8(v1334)
	v1337 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[127])) = uint8(v1337)
	v1341 = F_shm_toc_lookup(m, v90, int64(-65522), v1337)
	mBase = m.M
	v1342 = m.ExcPending
	if v1342 != 0 {
		goto L14
	} else {
		goto L372
	}
L372:
	;
	v1343 = m.G0
	v1345 = v1343 - int32(48)
	m.G0 = v1345
	v1347 = *(*int32)(unsafe.Add(mBase, uint32(v1341)))
	if v1347 != 0 {
		goto L373
	} else {
		goto L374
	}
L373:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1345)+16)) = int64(17179869188)
	v1351 = *(*int32)(unsafe.Add(mBase, _consts[76]))
	*(*int32)(unsafe.Add(mBase, uint32(v1345)+40)) = v1351
	v1357 = F_hash_create(m, int32(159917), int32(32), v1345, int32(1064))
	mBase = m.M
	v1358 = m.ExcPending
	if v1358 != 0 {
		goto L14
	} else {
		goto L376
	}
L374:
	;
	v1386 = v1341
	goto L375
L375:
	;
	v1402 = v1386 + int32(4)
	v1403 = *(*int32)(unsafe.Add(mBase, uint32(v1402)))
	if v1403 != 0 {
		goto L381
	} else {
		goto L382
	}
L376:
	;
	*(*int32)(unsafe.Add(mBase, _consts[128])) = v1357
	v1361 = v1341
	goto L377
L377:
	;
	v1377 = *(*int32)(unsafe.Add(mBase, _consts[128]))
	v1380 = F_hash_search(m, v1377, v1361, int32(1), int32(0))
	mBase = m.M
	v1381 = m.ExcPending
	if v1381 != 0 {
		goto L14
	} else {
		goto L379
	}
L378:
	;
	v1386 = v1383
	goto L375
L379:
	;
	v1383 = v1361 + int32(4)
	v1384 = *(*int32)(unsafe.Add(mBase, uint32(v1383)))
	if v1384 != 0 {
		v1361 = v1383
		goto L377
	} else {
		goto L380
	}
L380:
	;
	goto L378
L381:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1345)+16)) = int64(17179869188)
	v1407 = *(*int32)(unsafe.Add(mBase, _consts[76]))
	*(*int32)(unsafe.Add(mBase, uint32(v1345)+40)) = v1407
	v1413 = F_hash_create(m, int32(155767), int32(32), v1345, int32(1064))
	mBase = m.M
	v1414 = m.ExcPending
	if v1414 != 0 {
		goto L14
	} else {
		goto L384
	}
L382:
	;
	goto L383
L383:
	;
	m.G0 = v1345 + int32(48)
	v1462 = F_shm_toc_lookup(m, v90, int64(-65521), int32(0))
	mBase = m.M
	v1463 = m.ExcPending
	if v1463 != 0 {
		goto L14
	} else {
		goto L389
	}
L384:
	;
	*(*int32)(unsafe.Add(mBase, _consts[129])) = v1413
	v1417 = v1402
	goto L385
L385:
	;
	v1433 = *(*int32)(unsafe.Add(mBase, _consts[129]))
	v1436 = F_hash_search(m, v1433, v1417, int32(1), int32(0))
	mBase = m.M
	v1437 = m.ExcPending
	if v1437 != 0 {
		goto L14
	} else {
		goto L387
	}
L386:
	;
	goto L383
L387:
	;
	v1439 = v1417 + int32(4)
	v1440 = *(*int32)(unsafe.Add(mBase, uint32(v1439)))
	if v1440 != 0 {
		v1417 = v1439
		goto L385
	} else {
		goto L388
	}
L388:
	;
	goto L386
L389:
	;
	v1464 = *(*int32)(unsafe.Add(mBase, uint32(v1462)))
	v1466 = *(*int32)(unsafe.Add(mBase, uint32(v1462)+4))
	*(*int32)(unsafe.Add(mBase, _consts[130])) = v1466
	v1469 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[131])) = v1469
	if v1469 <= v1464 {
		goto L390
	} else {
		goto L391
	}
L390:
	;
	v1475 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	v1478 = F_MemoryContextStrdup(m, v1475, v1462+int32(8))
	mBase = m.M
	v1479 = m.ExcPending
	if v1479 != 0 {
		goto L14
	} else {
		goto L393
	}
L391:
	;
	goto L392
L392:
	;
	v1482 = *(*int32)(unsafe.Add(mBase, _consts[131]))
	if v1482 != 0 {
		goto L394
	} else {
		goto L395
	}
L393:
	;
	*(*int32)(unsafe.Add(mBase, _consts[131])) = v1478
	goto L392
L394:
	;
	v1484 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	v1489 = *(*int32)(unsafe.Add(mBase, uint32(v1484<<(uint(int32(2))%32))+uint32(_consts[132])))
	goto L397
L395:
	;
	goto L396
L396:
	;
	v1492 = *(*int32)(unsafe.Add(mBase, uint32(v96)+64))
	v1493 = m.G0
	v1495 = v1493 - int32(48)
	m.G0 = v1495
	*(*int32)(unsafe.Add(mBase, _consts[133])) = v1492
	if v1492 != 0 {
		goto L399
	} else {
		goto L400
	}
L397:
	;
	F_InitializeSystemUser(m, v1482, v1489)
	mBase = m.M
	v1491 = m.ExcPending
	if v1491 != 0 {
		goto L14
	} else {
		goto L398
	}
L398:
	;
	goto L396
L399:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1495)+16)) = int64(103079215120)
	v1504 = *(*int32)(unsafe.Add(mBase, _consts[134]))
	v1506 = F_hash_create(m, int32(312653), v1504, v1495, int32(40))
	mBase = m.M
	v1507 = m.ExcPending
	if v1507 != 0 {
		goto L14
	} else {
		goto L402
	}
L400:
	;
	goto L401
L401:
	;
	m.G0 = v1495 + int32(48)
	v1513 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[84])) = uint8(v1513)
	v1517 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	v1518 = *(*int32)(unsafe.Add(mBase, uint32(v1517)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v1517)+72)) = v1518 + int32(1)
	goto L403
L402:
	;
	*(*int32)(unsafe.Add(mBase, _consts[135])) = v1506
	goto L401
L403:
	;
	m.T0[v428].(func(*base.Module, int32, int32))(m, v83, v90)
	mBase = m.M
	v1523 = m.ExcPending
	if v1523 != 0 {
		goto L14
	} else {
		goto L404
	}
L404:
	;
	v1526 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	v1527 = *(*int32)(unsafe.Add(mBase, uint32(v1526)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v1526)+72)) = v1527 - int32(1)
	goto L405
L405:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v1532 = m.ExcPending
	if v1532 != 0 {
		goto L14
	} else {
		goto L406
	}
L406:
	;
	F_CommitTransaction(m)
	mBase = m.M
	v1534 = m.ExcPending
	if v1534 != 0 {
		goto L14
	} else {
		goto L407
	}
L407:
	;
	v1536 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	*(*int32)(unsafe.Add(mBase, uint32(v1536)+24)) = int32(0)
	v1540 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	v1541 = *(*int32)(unsafe.Add(mBase, uint32(v1540)))
	F_dsm_detach(m, v1541)
	mBase = m.M
	v1543 = m.ExcPending
	if v1543 != 0 {
		goto L14
	} else {
		goto L408
	}
L408:
	;
	v1544 = int32(4365212)
	v1545 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	*(*int32)(unsafe.Add(mBase, uint32(v1545))) = int32(0)
	v1549 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	v1550 = *(*int32)(unsafe.Add(mBase, uint32(v1549)+4))
	F_dsa_detach(m, v1550)
	mBase = m.M
	v1552 = m.ExcPending
	if v1552 != 0 {
		goto L14
	} else {
		goto L409
	}
L409:
	;
	v1554 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	v1555 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1554)+4)) = v1555
	v1561 = *(*int32)(unsafe.Add(mBase, _consts[136]))
	v1562 = *(*int32)(unsafe.Add(mBase, uint32(v1561)+16))
	v1563 = m.T0[v1562].(func(*base.Module, int32, int32, int32) int32)(m, int32(88), v1555, v1555)
	mBase = m.M
	v1564 = m.ExcPending
	if v1564 != 0 {
		goto L14
	} else {
		goto L410
	}
L410:
	;
	goto L46
L411:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1590 = m.ExcPending
	if v1590 != 0 {
		goto L14
	} else {
		goto L412
	}
L412:
	;
	F_errmsg(m, int32(93884), int32(0))
	mBase = m.M
	v1594 = m.ExcPending
	if v1594 != 0 {
		goto L14
	} else {
		goto L413
	}
L413:
	;
	F_errfinish(m, int32(489815), int32(1355), int32(274629))
	mBase = m.M
	v1599 = m.ExcPending
	if v1599 != 0 {
		goto L14
	} else {
		goto L414
	}
L414:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L415:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1606 = m.ExcPending
	if v1606 != 0 {
		goto L14
	} else {
		goto L416
	}
L416:
	;
	F_errmsg(m, int32(93928), int32(0))
	mBase = m.M
	v1610 = m.ExcPending
	if v1610 != 0 {
		goto L14
	} else {
		goto L417
	}
L417:
	;
	F_errfinish(m, int32(489815), int32(1360), int32(274629))
	mBase = m.M
	v1615 = m.ExcPending
	if v1615 != 0 {
		goto L14
	} else {
		goto L418
	}
L418:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L419:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v247
	F_errmsg_internal(m, int32(417619), v19+int32(16))
	mBase = m.M
	v1625 = m.ExcPending
	if v1625 != 0 {
		goto L14
	} else {
		goto L420
	}
L420:
	;
	F_errfinish(m, int32(489815), int32(1666), int32(250932))
	mBase = m.M
	v1630 = m.ExcPending
	if v1630 != 0 {
		goto L14
	} else {
		goto L421
	}
L421:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L422:
	;
	v1636 = *(*int32)(unsafe.Add(mBase, _consts[109]))
	v1637 = *(*int32)(unsafe.Add(mBase, uint32(v1636)+4))
	goto L423
L423:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v1637
	F_errmsg_internal(m, int32(447822), v19)
	mBase = m.M
	v1641 = m.ExcPending
	if v1641 != 0 {
		goto L14
	} else {
		goto L424
	}
L424:
	;
	F_errfinish(m, int32(489815), int32(1452), int32(274629))
	mBase = m.M
	v1646 = m.ExcPending
	if v1646 != 0 {
		goto L14
	} else {
		goto L425
	}
L425:
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
	v15 = *(*int32)(unsafe.Add(mBase, _consts[378]))
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
	v24 = *(*int32)(unsafe.Add(mBase, _consts[379]))
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
	v34 = *(*int32)(unsafe.Add(mBase, _consts[378]))
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
	v67 = *(*int32)(unsafe.Add(mBase, _consts[379]))
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
