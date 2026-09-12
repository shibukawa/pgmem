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
	v6 = int32(4480304)
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
		v24 = F_pstrdup(m, int32(160375))
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
					v39 = int32(4086764)
					*(*int32)(unsafe.Add(mBase, _consts[79])) = v39
					v43 = v39
				} else {
					v43 = v36
				}
				*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(4086764)
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
	v96 = int32(4474956)
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
	v104 = int32(4474956)
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
	v16 = int32(4480304)
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
				v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
						*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = int32(696)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v232
						v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
			v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
				v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
								v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
				v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
								v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
											v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
												v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
											v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
				v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
								v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
											v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
												v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
											v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
				v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
					v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
							v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
				v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
								v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
									v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
												v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
					v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
				v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
								v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
					v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
				v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
								v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
					v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
				v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(632)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(720)
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
								v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
					v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
				v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
					v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
							v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
				v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
					v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
							v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
				v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
					v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
							v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
				v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
					v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
							v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
				v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
					v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
							v581 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
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
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
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
	var v457 int32
	_ = v457
	var v460 int64
	_ = v460
	var v463 int32
	_ = v463
	var v464 int64
	_ = v464
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v504 int32
	_ = v504
	var v510 int32
	_ = v510
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v529 int32
	_ = v529
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v673 int32
	_ = v673
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v710 int32
	_ = v710
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v745 int32
	_ = v745
	var v749 int32
	_ = v749
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v758 int64
	_ = v758
	var v760 int32
	_ = v760
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
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
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v852 int32
	_ = v852
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v872 int32
	_ = v872
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v899 int32
	_ = v899
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v936 int32
	_ = v936
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v948 int32
	_ = v948
	var v953 int32
	_ = v953
	var v956 int32
	_ = v956
	var v960 int32
	_ = v960
	var v964 int32
	_ = v964
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v978 int32
	_ = v978
	var v983 int32
	_ = v983
	var v1007 int32
	_ = v1007
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1020 int32
	_ = v1020
	var v1023 int32
	_ = v1023
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1059 int32
	_ = v1059
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1065 int32
	_ = v1065
	var v1068 int32
	_ = v1068
	var v1071 int32
	_ = v1071
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1082 int32
	_ = v1082
	var v1084 int32
	_ = v1084
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1091 int32
	_ = v1091
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1109 int32
	_ = v1109
	var v1119 int32
	_ = v1119
	var v1136 int32
	_ = v1136
	var v1144 int32
	_ = v1144
	var v1148 int32
	_ = v1148
	var v1153 int32
	_ = v1153
	var v1157 int32
	_ = v1157
	var v1161 int32
	_ = v1161
	var v1166 int32
	_ = v1166
	var v1170 int32
	_ = v1170
	var v1174 int32
	_ = v1174
	var v1179 int32
	_ = v1179
	var v1183 int32
	_ = v1183
	var v1187 int32
	_ = v1187
	var v1192 int32
	_ = v1192
	var v1196 int32
	_ = v1196
	var v1199 int32
	_ = v1199
	var v1203 int32
	_ = v1203
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1215 int32
	_ = v1215
	var v1217 int32
	_ = v1217
	var v1222 int32
	_ = v1222
	var v1225 int32
	_ = v1225
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1233 int32
	_ = v1233
	var v1235 int32
	_ = v1235
	var v1239 int32
	_ = v1239
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1249 int32
	_ = v1249
	var v1265 int32
	_ = v1265
	var v1268 int32
	_ = v1268
	var v1269 int32
	_ = v1269
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1274 int32
	_ = v1274
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1295 int32
	_ = v1295
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1305 int32
	_ = v1305
	var v1321 int32
	_ = v1321
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1354 int32
	_ = v1354
	var v1357 int32
	_ = v1357
	var v1363 int32
	_ = v1363
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1370 int32
	_ = v1370
	var v1372 int32
	_ = v1372
	var v1377 int32
	_ = v1377
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1383 int32
	_ = v1383
	var v1392 int32
	_ = v1392
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1401 int32
	_ = v1401
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1411 int32
	_ = v1411
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1420 int32
	_ = v1420
	var v1422 int32
	_ = v1422
	var v1424 int32
	_ = v1424
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1440 int32
	_ = v1440
	var v1442 int32
	_ = v1442
	var v1443 int32
	_ = v1443
	var v1449 int32
	_ = v1449
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1452 int32
	_ = v1452
	var v1475 int32
	_ = v1475
	var v1478 int32
	_ = v1478
	var v1482 int32
	_ = v1482
	var v1487 int32
	_ = v1487
	var v1491 int32
	_ = v1491
	var v1494 int32
	_ = v1494
	var v1498 int32
	_ = v1498
	var v1503 int32
	_ = v1503
	var v1507 int32
	_ = v1507
	var v1513 int32
	_ = v1513
	var v1518 int32
	_ = v1518
	var v1522 int32
	_ = v1522
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1529 int32
	_ = v1529
	var v1534 int32
	_ = v1534
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
	v39 = int32(4730)
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
	v62 = F___memcpy(m, int32(4647316), v51, int32(140))
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
	v80 = F_AllocSetContextCreateInternal(m, v75, int32(219264), int32(0), int32(8192), int32(8388608))
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
	v1522 = m.ExcPending
	if v1522 != 0 {
		goto L14
	} else {
		goto L388
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1507 = m.ExcPending
	if v1507 != 0 {
		goto L14
	} else {
		goto L385
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1491 = m.ExcPending
	if v1491 != 0 {
		goto L14
	} else {
		goto L381
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
	v1475 = m.ExcPending
	if v1475 != 0 {
		goto L14
	} else {
		goto L377
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
	v188 = F_strlen(m, v186)
	mBase = m.M
	v191 = v188 + v186 + int32(1)
	v192 = int32(160375)
	v195 = int32(*(*uint8)(unsafe.Add(mBase, _consts[97])))
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
	if v196 == int32(0) {
		v215 = v195
		v216 = v196
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	*(*int32)(unsafe.Add(mBase, _consts[98])) = v374
	v377 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	*(*int32)(unsafe.Add(mBase, uint32(v377)+64)) = v374
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+32)))
	F_SetSessionAuthorization(m, v379, v380)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L14
	} else {
		goto L114
	}
L49:
	;
	if v216-v215 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L50:
	;
	goto L49
L51:
	;
	if v195 != v196 {
		v215 = v195
		v216 = v196
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v200 = v186
	v201 = v192
	goto L53
L53:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+1)))
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200)+1)))
	if v205 == int32(0) {
		v215 = v204
		v216 = v205
		goto L50
	} else {
		goto L55
	}
L54:
	;
	v215 = v204
	v216 = v205
	goto L50
L55:
	;
	v208 = int32(1)
	if v204 == v205 {
		v200 = v200 + v208
		v201 = v201 + v208
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v220 = int32(275816)
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191))))
	v224 = int32(*(*uint8)(unsafe.Add(mBase, _consts[99])))
	if v224 == int32(0) {
		v243 = v223
		v244 = v224
		goto L61
	} else {
		goto L62
	}
L58:
	;
	goto L59
L59:
	;
	v370 = F_load_external_function(m, v186, v191, int32(1), int32(0))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L14
	} else {
		goto L113
	}
L60:
	;
	if v244-v243 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L61:
	;
	goto L60
L62:
	;
	if v223 != v224 {
		v243 = v223
		v244 = v224
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v228 = v220
	v229 = v191
	goto L64
L64:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229)+1)))
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+1)))
	if v233 == int32(0) {
		v243 = v232
		v244 = v233
		goto L61
	} else {
		goto L66
	}
L65:
	;
	v243 = v232
	v244 = v233
	goto L61
L66:
	;
	v236 = int32(1)
	if v232 == v233 {
		v228 = v228 + v236
		v229 = v229 + v236
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	v249 = *(*int32)(unsafe.Add(mBase, _consts[100]))
	v372 = v249
	goto L48
L69:
	;
	goto L70
L70:
	;
	v250 = int32(275623)
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191))))
	v254 = int32(*(*uint8)(unsafe.Add(mBase, _consts[101])))
	if v254 == int32(0) {
		v273 = v253
		v274 = v254
		goto L72
	} else {
		goto L73
	}
L71:
	;
	if v274-v273 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L72:
	;
	goto L71
L73:
	;
	if v253 != v254 {
		v273 = v253
		v274 = v254
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v258 = v250
	v259 = v191
	goto L75
L75:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+1)))
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+1)))
	if v263 == int32(0) {
		v273 = v262
		v274 = v263
		goto L72
	} else {
		goto L77
	}
L76:
	;
	v273 = v262
	v274 = v263
	goto L72
L77:
	;
	v266 = int32(1)
	if v262 == v263 {
		v258 = v258 + v266
		v259 = v259 + v266
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v279 = *(*int32)(unsafe.Add(mBase, _consts[102]))
	v372 = v279
	goto L48
L80:
	;
	goto L81
L81:
	;
	v280 = int32(275647)
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191))))
	v284 = int32(*(*uint8)(unsafe.Add(mBase, _consts[103])))
	if v284 == int32(0) {
		v303 = v283
		v304 = v284
		goto L83
	} else {
		goto L84
	}
L82:
	;
	if v304-v303 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L83:
	;
	goto L82
L84:
	;
	if v283 != v284 {
		v303 = v283
		v304 = v284
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v288 = v280
	v289 = v191
	goto L86
L86:
	;
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289)+1)))
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v288)+1)))
	if v293 == int32(0) {
		v303 = v292
		v304 = v293
		goto L83
	} else {
		goto L88
	}
L87:
	;
	v303 = v292
	v304 = v293
	goto L83
L88:
	;
	v296 = int32(1)
	if v292 == v293 {
		v288 = v288 + v296
		v289 = v289 + v296
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	v309 = *(*int32)(unsafe.Add(mBase, _consts[104]))
	v372 = v309
	goto L48
L91:
	;
	goto L92
L92:
	;
	v310 = int32(275673)
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191))))
	v314 = int32(*(*uint8)(unsafe.Add(mBase, _consts[105])))
	if v314 == int32(0) {
		v333 = v313
		v334 = v314
		goto L94
	} else {
		goto L95
	}
L93:
	;
	if v334-v333 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L94:
	;
	goto L93
L95:
	;
	if v313 != v314 {
		v333 = v313
		v334 = v314
		goto L94
	} else {
		goto L96
	}
L96:
	;
	v318 = v310
	v319 = v191
	goto L97
L97:
	;
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319)+1)))
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318)+1)))
	if v323 == int32(0) {
		v333 = v322
		v334 = v323
		goto L94
	} else {
		goto L99
	}
L98:
	;
	v333 = v322
	v334 = v323
	goto L94
L99:
	;
	v326 = int32(1)
	if v322 == v323 {
		v318 = v318 + v326
		v319 = v319 + v326
		goto L97
	} else {
		goto L100
	}
L100:
	;
	goto L98
L101:
	;
	v339 = *(*int32)(unsafe.Add(mBase, _consts[106]))
	v372 = v339
	goto L48
L102:
	;
	goto L103
L103:
	;
	v340 = int32(275602)
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191))))
	v344 = int32(*(*uint8)(unsafe.Add(mBase, _consts[107])))
	if v344 == int32(0) {
		v363 = v343
		v364 = v344
		goto L105
	} else {
		goto L106
	}
L104:
	;
	if v364-v363 != 0 {
		goto L18
	} else {
		goto L112
	}
L105:
	;
	goto L104
L106:
	;
	if v343 != v344 {
		v363 = v343
		v364 = v344
		goto L105
	} else {
		goto L107
	}
L107:
	;
	v348 = v340
	v349 = v191
	goto L108
L108:
	;
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349)+1)))
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v348)+1)))
	if v353 == int32(0) {
		v363 = v352
		v364 = v353
		goto L105
	} else {
		goto L110
	}
L109:
	;
	v363 = v352
	v364 = v353
	goto L105
L110:
	;
	v356 = int32(1)
	if v352 == v353 {
		v348 = v348 + v356
		v349 = v349 + v356
		goto L108
	} else {
		goto L111
	}
L111:
	;
	goto L109
L112:
	;
	v367 = *(*int32)(unsafe.Add(mBase, _consts[108]))
	v372 = v367
	goto L48
L113:
	;
	v372 = v370
	goto L48
L114:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v96)+12))
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+33)))
	F_SetCurrentRoleId(m, v383, v384)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L14
	} else {
		goto L115
	}
L115:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	F_BackgroundWorkerInitializeConnectionByOid(m, v387, v388, int32(3))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L14
	} else {
		goto L116
	}
L116:
	;
	v393 = *(*int32)(unsafe.Add(mBase, _consts[109]))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v393)+4))
	goto L117
L117:
	;
	v395 = F_SetClientEncoding(m, v394)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L14
	} else {
		goto L118
	}
L118:
	;
	if v395 < int32(0) {
		goto L17
	} else {
		goto L119
	}
L119:
	;
	v401 = F_shm_toc_lookup(m, v90, int64(-65533), int32(0))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L14
	} else {
		goto L120
	}
L120:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L14
	} else {
		goto L121
	}
L121:
	;
	v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v401))))
	if v405 != 0 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v407 = v401
	goto L125
L123:
	;
	goto L124
L124:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L14
	} else {
		goto L129
	}
L125:
	;
	v422 = F_internal_load_library(m, v407)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L14
	} else {
		goto L127
	}
L126:
	;
	goto L124
L127:
	;
	v424 = F_strlen(m, v407)
	mBase = m.M
	v427 = v424 + v407 + int32(1)
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v427))))
	if v428 != 0 {
		v407 = v427
		goto L125
	} else {
		goto L128
	}
L128:
	;
	goto L126
L129:
	;
	v449 = F_shm_toc_lookup(m, v90, int64(-65528), int32(0))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L14
	} else {
		goto L130
	}
L130:
	;
	F_StartTransaction(m)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L14
	} else {
		goto L131
	}
L131:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v449)))
	*(*int32)(unsafe.Add(mBase, _consts[110])) = v454
	v457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v449)+4)))
	*(*uint8)(unsafe.Add(mBase, _consts[111])) = uint8(v457)
	v460 = *(*int64)(unsafe.Add(mBase, uint32(v449)+8))
	*(*int64)(unsafe.Add(mBase, _consts[35])) = v460
	v463 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	v464 = *(*int64)(unsafe.Add(mBase, uint32(v449)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v463))) = v464
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v449)+24))
	*(*int32)(unsafe.Add(mBase, _consts[112])) = v467
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v449)+28))
	*(*int32)(unsafe.Add(mBase, _consts[38])) = v449 + int32(32)
	*(*int32)(unsafe.Add(mBase, _consts[36])) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v463)+24)) = int32(5)
	v480 = F_shm_toc_lookup(m, v90, int64(-65525), int32(0))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L14
	} else {
		goto L132
	}
L132:
	;
	v482 = m.G0
	v484 = v482 - int32(48)
	m.G0 = v484
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v480)+8))
	if v486 != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v488 = v480
	goto L136
L134:
	;
	goto L135
L135:
	;
	m.G0 = v484 + int32(48)
	v551 = F_shm_toc_lookup(m, v90, int64(-65523), int32(0))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L14
	} else {
		goto L144
	}
L136:
	;
	v504 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	if v504 == int32(0) {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	goto L135
L138:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v484)+16)) = int64(68719476748)
	v510 = *(*int32)(unsafe.Add(mBase, _consts[76]))
	*(*int32)(unsafe.Add(mBase, uint32(v484)+40)) = v510
	v516 = F_hash_create(m, int32(320276), int32(16), v484, int32(1064))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L14
	} else {
		goto L141
	}
L139:
	;
	v519 = v504
	goto L140
L140:
	;
	v521 = F_hash_search(m, v519, v488, int32(1), v484)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L14
	} else {
		goto L142
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, _consts[113])) = v516
	v519 = v516
	goto L140
L142:
	;
	v523 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v521)+12)) = uint8(v523)
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v488+int32(20))))
	if v529 != 0 {
		v488 = v488 + int32(12)
		goto L136
	} else {
		goto L143
	}
L143:
	;
	goto L137
L144:
	;
	v554 = *(*int32)(unsafe.Add(mBase, _consts[114]))
	if v554 != 0 {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	goto L155
L146:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L14
	} else {
		goto L151
	}
L147:
	;
	v556 = *(*int32)(unsafe.Add(mBase, _consts[115]))
	if v556 != 0 {
		goto L146
	} else {
		goto L148
	}
L148:
	;
	v558 = *(*int32)(unsafe.Add(mBase, _consts[116]))
	if v558 != 0 {
		goto L146
	} else {
		goto L149
	}
L149:
	;
	v560 = *(*int32)(unsafe.Add(mBase, _consts[117]))
	if v560 == int32(0) {
		goto L145
	} else {
		goto L150
	}
L150:
	;
	goto L146
L151:
	;
	F_errmsg_internal(m, int32(154920), int32(0))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L14
	} else {
		goto L152
	}
L152:
	;
	F_errfinish(m, int32(490311), int32(749), int32(236780))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L14
	} else {
		goto L153
	}
L153:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L154:
	;
	v581 = int32(524)
	goto L159
L155:
	;
	v578 = F__emscripten_memcpy_bulkmem(m, int32(4468576), v551, int32(524))
	mBase = m.M
	goto L157
L157:
	;
	goto L154
L158:
	;
	v588 = F_shm_toc_lookup(m, v90, int64(-65524), int32(0))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L14
	} else {
		goto L162
	}
L159:
	;
	v584 = F__emscripten_memcpy_bulkmem(m, int32(4469624), v551+v581, v581)
	mBase = m.M
	goto L161
L161:
	;
	goto L158
L162:
	;
	v590 = int32(0)
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v588)))
	*(*int32)(unsafe.Add(mBase, _consts[118])) = v592
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v588)+4))
	*(*int32)(unsafe.Add(mBase, _consts[43])) = v595
	v597 = int32(4480304)
	v598 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v601 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v601
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v588)+8))
	if v590 < v603 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v609 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v612 = v590
	v613 = v609
	goto L166
L164:
	;
	goto L165
L165:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v598
	v658 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v658)+28))
	goto L170
L166:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v588+int32(12)+v612<<(uint(int32(2))%32))))
	v631 = F_lappend_oid(m, v613, v630)
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L14
	} else {
		goto L168
	}
L167:
	;
	goto L165
L168:
	;
	*(*int32)(unsafe.Add(mBase, _consts[44])) = v631
	v635 = v612 + int32(1)
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v588)+8))
	if v635 < v636 {
		v612 = v635
		v613 = v631
		goto L166
	} else {
		goto L169
	}
L169:
	;
	goto L167
L170:
	;
	*(*int32)(unsafe.Add(mBase, _consts[119])) = v659
	v663 = F_shm_toc_lookup(m, v90, int64(-65531), int32(0))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L14
	} else {
		goto L171
	}
L171:
	;
	v665 = int32(0)
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v663)))
	if v666 <= v665 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v729 = F_shm_toc_lookup(m, v90, int64(-65526), int32(0))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L14
	} else {
		goto L184
	}
L173:
	;
	v673 = v665
	goto L174
L174:
	;
	v689 = v663 + int32(4) + v673<<(uint(int32(3))%32)
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v689)))
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v689)+4))
	v692 = F_GetComboCommandId(m, v690, v691)
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L14
	} else {
		goto L176
	}
L175:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L14
	} else {
		goto L181
	}
L176:
	;
	if v692 == v673 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v696 = v673 + int32(1)
	if v666 != v696 {
		v673 = v696
		goto L174
	} else {
		goto L180
	}
L178:
	;
	goto L179
L179:
	;
	goto L175
L180:
	;
	goto L172
L181:
	;
	F_errmsg_internal(m, int32(173605), int32(0))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L14
	} else {
		goto L182
	}
L182:
	;
	F_errfinish(m, int32(494697), int32(362), int32(350694))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L14
	} else {
		goto L183
	}
L183:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L184:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v729)))
	v732 = int32(4480304)
	v733 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v736 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v736
	v738 = F_dsm_attach(m, v731)
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L14
	} else {
		goto L185
	}
L185:
	;
	if v738 == int32(0) {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L14
	} else {
		goto L189
	}
L187:
	;
	goto L188
L188:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v738)+24))
	v758 = *(*int64)(unsafe.Add(mBase, uint32(v756)))
	if v758 == int64(2880502729) {
		goto L193
	} else {
		goto L194
	}
L189:
	;
	F_errmsg_internal(m, int32(94554), int32(0))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L14
	} else {
		goto L190
	}
L190:
	;
	F_errfinish(m, int32(491415), int32(169), int32(268416))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L14
	} else {
		goto L191
	}
L191:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L192:
	;
	v763 = F_shm_toc_lookup(m, v760, int64(-65535), int32(0))
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L14
	} else {
		goto L196
	}
L193:
	;
	v760 = v756
	goto L195
L194:
	;
	v760 = int32(0)
	goto L195
L195:
	;
	goto L192
L196:
	;
	v765 = F_dsa_attach_in_place(m, v763, v738)
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L14
	} else {
		goto L197
	}
L197:
	;
	v767 = int32(4374764)
	v768 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	*(*int32)(unsafe.Add(mBase, uint32(v768))) = v738
	v771 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	*(*int32)(unsafe.Add(mBase, uint32(v771)+4)) = v765
	v775 = F_shm_toc_lookup(m, v760, int64(-65534), int32(0))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L14
	} else {
		goto L198
	}
L198:
	;
	v777 = int32(4480304)
	v778 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v781 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v781
	v784 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v784)+4))
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v775)))
	v788 = F_dshash_attach(m, v785, int32(1741644), v787, v785)
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L14
	} else {
		goto L199
	}
L199:
	;
	v791 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v791)+4))
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v775)+4))
	v796 = F_dshash_attach(m, v792, int32(1741668), v794, int32(0))
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L14
	} else {
		goto L200
	}
L200:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v778
	v801 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v801)))
	F_on_dsm_detach(m, v802, int32(1623), v775)
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L14
	} else {
		goto L201
	}
L201:
	;
	v807 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	*(*int32)(unsafe.Add(mBase, uint32(v807)+16)) = v796
	*(*int32)(unsafe.Add(mBase, uint32(v807)+12)) = v788
	*(*int32)(unsafe.Add(mBase, uint32(v807)+8)) = v775
	F_dsm_pin_mapping(m, v738)
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L14
	} else {
		goto L202
	}
L202:
	;
	F_dsa_pin_mapping(m, v765)
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L14
	} else {
		goto L203
	}
L203:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v733
	v819 = F_shm_toc_lookup(m, v90, int64(-65529), int32(0))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L14
	} else {
		goto L204
	}
L204:
	;
	v823 = F_shm_toc_lookup(m, v90, int64(-65530), int32(1))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L14
	} else {
		goto L205
	}
L205:
	;
	v825 = F_RestoreSnapshot(m, v819)
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L14
	} else {
		goto L206
	}
L206:
	;
	if v823 != 0 {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v827 = F_RestoreSnapshot(m, v823)
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L14
	} else {
		goto L210
	}
L208:
	;
	v829 = v825
	goto L209
L209:
	;
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v96)+36))
	F_RestoreTransactionSnapshot(m, v829, v830)
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L14
	} else {
		goto L211
	}
L210:
	;
	v829 = v827
	goto L209
L211:
	;
	F_PushActiveSnapshot(m, v825)
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L14
	} else {
		goto L212
	}
L212:
	;
	F_InvalidateSystemCachesExtended(m)
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L14
	} else {
		goto L213
	}
L213:
	;
	v839 = F_shm_toc_lookup(m, v90, int64(-65532), int32(0))
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L14
	} else {
		goto L214
	}
L214:
	;
	v841 = m.G0
	v843 = v841 - int32(32)
	m.G0 = v843
	v846 = *(*int32)(unsafe.Add(mBase, _consts[121]))
	if v846 == int32(0) {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(v839)))
	*(*int32)(unsafe.Add(mBase, uint32(v843)+20)) = int32(1659)
	v1010 = int32(4473208)
	v1011 = *(*int32)(unsafe.Add(mBase, _consts[77]))
	*(*int32)(unsafe.Add(mBase, _consts[77])) = v843 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v843)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v843)+16)) = v1011
	v1020 = v839 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v843)+28)) = v1020
	v1023 = v1007 + v1020
	if base.Ui32(v1020) < base.Ui32(v1023) {
		goto L290
	} else {
		goto L291
	}
L216:
	;
	if v846 == int32(4478332) {
		goto L215
	} else {
		goto L217
	}
L217:
	;
	v852 = v846
	goto L218
L218:
	;
	v868 = v852 + int32(4)
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v868)))
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v852-int32(60))))
	if base.Ui32(v872) < base.Ui32(int32(2)) {
		goto L220
	} else {
		goto L221
	}
L219:
	;
	goto L215
L220:
	;
	if v869 != int32(4478332) {
		v852 = v869
		goto L218
	} else {
		goto L283
	}
L221:
	;
	v876 = v852 - int32(32)
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v876)))
	if v877 == int32(0) {
		goto L220
	} else {
		goto L222
	}
L222:
	;
	v881 = v852 - int32(4)
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v881)))
	if v882 != 0 {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	F_pfree(m, v882)
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L14
	} else {
		goto L226
	}
L224:
	;
	goto L225
L225:
	;
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v852)+16))
	if v885 != 0 {
		goto L227
	} else {
		goto L228
	}
L226:
	;
	goto L225
L227:
	;
	F_pfree(m, v885)
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L14
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v852)+20))
	if v888 != 0 {
		goto L231
	} else {
		goto L232
	}
L230:
	;
	goto L229
L231:
	;
	F_pfree(m, v888)
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L14
	} else {
		goto L234
	}
L232:
	;
	goto L233
L233:
	;
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v852-int32(40))))
	switch v893 {
	case 0:
		goto L241
	case 1:
		goto L240
	case 2:
		goto L239
	case 3:
		goto L238
	case 4:
		goto L237
	default:
		goto L235
	}
L234:
	;
	goto L233
L235:
	;
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v876)))
	if v939 != 0 {
		goto L261
	} else {
		goto L262
	}
L236:
	;
	F_pfree(m, v933)
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L14
	} else {
		goto L260
	}
L237:
	;
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v852)+56))
	if v928 == int32(0) {
		goto L235
	} else {
		goto L258
	}
L238:
	;
	v910 = v852 + int32(28)
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v910)))
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v911)))
	if v912 != 0 {
		goto L248
	} else {
		goto L249
	}
L239:
	;
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v852)+80))
	if v904 == int32(0) {
		goto L235
	} else {
		goto L246
	}
L240:
	;
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v852)+60))
	if v899 == int32(0) {
		goto L235
	} else {
		goto L244
	}
L241:
	;
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v852)+52))
	if v894 == int32(0) {
		goto L235
	} else {
		goto L242
	}
L242:
	;
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v881)))
	if v894 != v897 {
		v933 = v894
		goto L236
	} else {
		goto L243
	}
L243:
	;
	goto L235
L244:
	;
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v881)))
	if v899 != v902 {
		v933 = v899
		goto L236
	} else {
		goto L245
	}
L245:
	;
	goto L235
L246:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v881)))
	if v904 != v907 {
		v933 = v904
		goto L236
	} else {
		goto L247
	}
L247:
	;
	goto L235
L248:
	;
	F_pfree(m, v912)
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L14
	} else {
		goto L251
	}
L249:
	;
	goto L250
L250:
	;
	v915 = *(*int32)(unsafe.Add(mBase, uint32(v852)+48))
	if v915 == int32(0) {
		goto L252
	} else {
		goto L253
	}
L251:
	;
	goto L250
L252:
	;
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v852)+52))
	if v923 == int32(0) {
		goto L235
	} else {
		goto L256
	}
L253:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v910)))
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v918)))
	if v915 == v919 {
		goto L252
	} else {
		goto L254
	}
L254:
	;
	F_pfree(m, v915)
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L14
	} else {
		goto L255
	}
L255:
	;
	goto L252
L256:
	;
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v881)))
	if v923 != v926 {
		v933 = v923
		goto L236
	} else {
		goto L257
	}
L257:
	;
	goto L235
L258:
	;
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v881)))
	if v928 == v931 {
		goto L235
	} else {
		goto L259
	}
L259:
	;
	v933 = v928
	goto L236
L260:
	;
	goto L235
L261:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v852)))
	v941 = *(*int32)(unsafe.Add(mBase, uint32(v868)))
	*(*int32)(unsafe.Add(mBase, uint32(v940)+4)) = v941
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v852)))
	*(*int32)(unsafe.Add(mBase, uint32(v941))) = v943
	goto L263
L262:
	;
	goto L263
L263:
	;
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v852-int32(8))))
	if v948 != 0 {
		goto L264
	} else {
		goto L265
	}
L264:
	;
	v953 = int32(4478348)
	goto L269
L265:
	;
	goto L266
L266:
	;
	v964 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852-int32(36)))))
	if v964&int32(4) != 0 {
		goto L273
	} else {
		goto L274
	}
L267:
	;
	goto L266
L268:
	;
	goto L267
L269:
	;
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v953)))
	if v956 == int32(0) {
		goto L268
	} else {
		goto L271
	}
L270:
	;
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v956)))
	*(*int32)(unsafe.Add(mBase, uint32(v953))) = v960
	goto L268
L271:
	;
	if v956 != v852+int32(8) {
		v953 = v956
		goto L269
	} else {
		goto L272
	}
L272:
	;
	goto L270
L273:
	;
	v971 = int32(4478340)
	goto L278
L274:
	;
	goto L275
L275:
	;
	F_InitializeOneGUCOption(m, v852+int32(-64))
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L14
	} else {
		goto L282
	}
L276:
	;
	goto L275
L277:
	;
	goto L276
L278:
	;
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v971)))
	if v974 == int32(0) {
		goto L277
	} else {
		goto L280
	}
L279:
	;
	v978 = *(*int32)(unsafe.Add(mBase, uint32(v974)))
	*(*int32)(unsafe.Add(mBase, uint32(v971))) = v978
	goto L277
L280:
	;
	if v974 != v852+int32(12) {
		v971 = v974
		goto L278
	} else {
		goto L281
	}
L281:
	;
	goto L279
L282:
	;
	goto L220
L283:
	;
	goto L219
L284:
	;
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(v96)+16))
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(v96)+28))
	*(*int32)(unsafe.Add(mBase, _consts[122])) = v1210
	*(*int32)(unsafe.Add(mBase, _consts[4])) = v1209
	goto L337
L285:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1196 = m.ExcPending
	if v1196 != 0 {
		goto L14
	} else {
		goto L333
	}
L286:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1183 = m.ExcPending
	if v1183 != 0 {
		goto L14
	} else {
		goto L330
	}
L287:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1170 = m.ExcPending
	if v1170 != 0 {
		goto L14
	} else {
		goto L327
	}
L288:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1157 = m.ExcPending
	if v1157 != 0 {
		goto L14
	} else {
		goto L324
	}
L289:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1144 = m.ExcPending
	if v1144 != 0 {
		goto L14
	} else {
		goto L321
	}
L290:
	;
	goto L293
L291:
	;
	v1136 = v1011
	goto L292
L292:
	;
	*(*int32)(unsafe.Add(mBase, _consts[77])) = v1136
	m.G0 = v843 + int32(32)
	goto L284
L293:
	;
	v1043 = F_read_gucstate(m, v843+int32(28), v1023)
	mBase = m.M
	v1044 = m.ExcPending
	if v1044 != 0 {
		goto L14
	} else {
		goto L295
	}
L294:
	;
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(v843)+16))
	v1136 = v1119
	goto L292
L295:
	;
	v1047 = F_read_gucstate(m, v843+int32(28), v1023)
	mBase = m.M
	v1048 = m.ExcPending
	if v1048 != 0 {
		goto L14
	} else {
		goto L296
	}
L296:
	;
	v1051 = F_read_gucstate(m, v843+int32(28), v1023)
	mBase = m.M
	v1052 = m.ExcPending
	if v1052 != 0 {
		goto L14
	} else {
		goto L297
	}
L297:
	;
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v843)+28))
	v1054 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1051))))
	if v1054 == int32(0) {
		goto L299
	} else {
		goto L300
	}
L298:
	;
	v1065 = v1062 + int32(4)
	if base.Ui32(v1023) < base.Ui32(v1065) {
		goto L288
	} else {
		goto L303
	}
L299:
	;
	v1062 = v1053
	v1063 = int32(0)
	goto L298
L300:
	;
	goto L301
L301:
	;
	v1059 = v1053 + int32(4)
	if base.Ui32(v1023) < base.Ui32(v1059) {
		goto L289
	} else {
		goto L302
	}
L302:
	;
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(v1053)))
	v1062 = v1059
	v1063 = v1061
	goto L298
L303:
	;
	v1068 = v1062 + int32(8)
	if base.Ui32(v1023) < base.Ui32(v1068) {
		goto L287
	} else {
		goto L304
	}
L304:
	;
	v1071 = v1062 + int32(12)
	if base.Ui32(v1023) < base.Ui32(v1071) {
		goto L286
	} else {
		goto L305
	}
L305:
	;
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(v1062)))
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(v1068)))
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(v1065)))
	*(*int32)(unsafe.Add(mBase, uint32(v843)+12)) = v1047
	*(*int32)(unsafe.Add(mBase, uint32(v843)+8)) = v1043
	*(*int32)(unsafe.Add(mBase, uint32(v843)+28)) = v1071
	*(*int32)(unsafe.Add(mBase, uint32(v843)+24)) = v843 + int32(8)
	v1082 = int32(0)
	v1084 = int32(1)
	v1087 = F_set_config_with_handle(m, v1043, v1082, v1047, v1075, v1073, v1074, v1082, v1084, int32(21), v1084)
	mBase = m.M
	v1088 = m.ExcPending
	if v1088 != 0 {
		goto L14
	} else {
		goto L306
	}
L306:
	;
	if v1087 <= int32(0) {
		goto L285
	} else {
		goto L307
	}
L307:
	;
	v1091 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1051))))
	if v1091 == int32(0) {
		goto L308
	} else {
		goto L309
	}
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v843)+24)) = int32(0)
	if base.Ui32(v1071) < base.Ui32(v1023) {
		goto L293
	} else {
		goto L320
	}
L309:
	;
	v1099 = int32(*(*uint8)(unsafe.Add(mBase, _consts[123])))
	if v1099 != 0 {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	v1100 = int32(12)
	goto L312
L311:
	;
	v1100 = int32(15)
	goto L312
L312:
	;
	v1101 = F_find_option(m, v1043, int32(1), int32(0), v1100)
	mBase = m.M
	v1102 = m.ExcPending
	if v1102 != 0 {
		goto L14
	} else {
		goto L313
	}
L313:
	;
	if v1101 == int32(0) {
		goto L308
	} else {
		goto L314
	}
L314:
	;
	v1105 = F_guc_strdup(m, v1100, v1051)
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L14
	} else {
		goto L315
	}
L315:
	;
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+84))
	if v1107 != 0 {
		goto L316
	} else {
		goto L317
	}
L316:
	;
	F_pfree(m, v1107)
	mBase = m.M
	v1109 = m.ExcPending
	if v1109 != 0 {
		goto L14
	} else {
		goto L319
	}
L317:
	;
	goto L318
L318:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1101)+88)) = v1063
	*(*int32)(unsafe.Add(mBase, uint32(v1101)+84)) = v1105
	goto L308
L319:
	;
	goto L318
L320:
	;
	goto L294
L321:
	;
	F_errmsg_internal(m, int32(350049), int32(0))
	mBase = m.M
	v1148 = m.ExcPending
	if v1148 != 0 {
		goto L14
	} else {
		goto L322
	}
L322:
	;
	F_errfinish(m, int32(494737), int32(6168), int32(17728))
	mBase = m.M
	v1153 = m.ExcPending
	if v1153 != 0 {
		goto L14
	} else {
		goto L323
	}
L323:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L324:
	;
	F_errmsg_internal(m, int32(350049), int32(0))
	mBase = m.M
	v1161 = m.ExcPending
	if v1161 != 0 {
		goto L14
	} else {
		goto L325
	}
L325:
	;
	F_errfinish(m, int32(494737), int32(6168), int32(17728))
	mBase = m.M
	v1166 = m.ExcPending
	if v1166 != 0 {
		goto L14
	} else {
		goto L326
	}
L326:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L327:
	;
	F_errmsg_internal(m, int32(350049), int32(0))
	mBase = m.M
	v1174 = m.ExcPending
	if v1174 != 0 {
		goto L14
	} else {
		goto L328
	}
L328:
	;
	F_errfinish(m, int32(494737), int32(6168), int32(17728))
	mBase = m.M
	v1179 = m.ExcPending
	if v1179 != 0 {
		goto L14
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
	F_errmsg_internal(m, int32(350049), int32(0))
	mBase = m.M
	v1187 = m.ExcPending
	if v1187 != 0 {
		goto L14
	} else {
		goto L331
	}
L331:
	;
	F_errfinish(m, int32(494737), int32(6168), int32(17728))
	mBase = m.M
	v1192 = m.ExcPending
	if v1192 != 0 {
		goto L14
	} else {
		goto L332
	}
L332:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L333:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v1199 = m.ExcPending
	if v1199 != 0 {
		goto L14
	} else {
		goto L334
	}
L334:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v843))) = v1043
	F_errmsg(m, int32(105555), v843)
	mBase = m.M
	v1203 = m.ExcPending
	if v1203 != 0 {
		goto L14
	} else {
		goto L335
	}
L335:
	;
	F_errfinish(m, int32(494737), int32(6353), int32(350715))
	mBase = m.M
	v1208 = m.ExcPending
	if v1208 != 0 {
		goto L14
	} else {
		goto L336
	}
L336:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L337:
	;
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(v96)+20))
	v1217 = *(*int32)(unsafe.Add(mBase, uint32(v96)+24))
	*(*int32)(unsafe.Add(mBase, _consts[124])) = v1217
	*(*int32)(unsafe.Add(mBase, _consts[125])) = v1215
	v1222 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[126])) = uint8(v1222)
	v1225 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[127])) = uint8(v1225)
	v1229 = F_shm_toc_lookup(m, v90, int64(-65522), v1225)
	mBase = m.M
	v1230 = m.ExcPending
	if v1230 != 0 {
		goto L14
	} else {
		goto L338
	}
L338:
	;
	v1231 = m.G0
	v1233 = v1231 - int32(48)
	m.G0 = v1233
	v1235 = *(*int32)(unsafe.Add(mBase, uint32(v1229)))
	if v1235 != 0 {
		goto L339
	} else {
		goto L340
	}
L339:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1233)+16)) = int64(17179869188)
	v1239 = *(*int32)(unsafe.Add(mBase, _consts[76]))
	*(*int32)(unsafe.Add(mBase, uint32(v1233)+40)) = v1239
	v1245 = F_hash_create(m, int32(161040), int32(32), v1233, int32(1064))
	mBase = m.M
	v1246 = m.ExcPending
	if v1246 != 0 {
		goto L14
	} else {
		goto L342
	}
L340:
	;
	v1274 = v1229
	goto L341
L341:
	;
	v1290 = v1274 + int32(4)
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(v1290)))
	if v1291 != 0 {
		goto L347
	} else {
		goto L348
	}
L342:
	;
	*(*int32)(unsafe.Add(mBase, _consts[128])) = v1245
	v1249 = v1229
	goto L343
L343:
	;
	v1265 = *(*int32)(unsafe.Add(mBase, _consts[128]))
	v1268 = F_hash_search(m, v1265, v1249, int32(1), int32(0))
	mBase = m.M
	v1269 = m.ExcPending
	if v1269 != 0 {
		goto L14
	} else {
		goto L345
	}
L344:
	;
	v1274 = v1271
	goto L341
L345:
	;
	v1271 = v1249 + int32(4)
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(v1271)))
	if v1272 != 0 {
		v1249 = v1271
		goto L343
	} else {
		goto L346
	}
L346:
	;
	goto L344
L347:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1233)+16)) = int64(17179869188)
	v1295 = *(*int32)(unsafe.Add(mBase, _consts[76]))
	*(*int32)(unsafe.Add(mBase, uint32(v1233)+40)) = v1295
	v1301 = F_hash_create(m, int32(156796), int32(32), v1233, int32(1064))
	mBase = m.M
	v1302 = m.ExcPending
	if v1302 != 0 {
		goto L14
	} else {
		goto L350
	}
L348:
	;
	goto L349
L349:
	;
	m.G0 = v1233 + int32(48)
	v1350 = F_shm_toc_lookup(m, v90, int64(-65521), int32(0))
	mBase = m.M
	v1351 = m.ExcPending
	if v1351 != 0 {
		goto L14
	} else {
		goto L355
	}
L350:
	;
	*(*int32)(unsafe.Add(mBase, _consts[129])) = v1301
	v1305 = v1290
	goto L351
L351:
	;
	v1321 = *(*int32)(unsafe.Add(mBase, _consts[129]))
	v1324 = F_hash_search(m, v1321, v1305, int32(1), int32(0))
	mBase = m.M
	v1325 = m.ExcPending
	if v1325 != 0 {
		goto L14
	} else {
		goto L353
	}
L352:
	;
	goto L349
L353:
	;
	v1327 = v1305 + int32(4)
	v1328 = *(*int32)(unsafe.Add(mBase, uint32(v1327)))
	if v1328 != 0 {
		v1305 = v1327
		goto L351
	} else {
		goto L354
	}
L354:
	;
	goto L352
L355:
	;
	v1352 = *(*int32)(unsafe.Add(mBase, uint32(v1350)))
	v1354 = *(*int32)(unsafe.Add(mBase, uint32(v1350)+4))
	*(*int32)(unsafe.Add(mBase, _consts[130])) = v1354
	v1357 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[131])) = v1357
	if v1357 <= v1352 {
		goto L356
	} else {
		goto L357
	}
L356:
	;
	v1363 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	v1366 = F_MemoryContextStrdup(m, v1363, v1350+int32(8))
	mBase = m.M
	v1367 = m.ExcPending
	if v1367 != 0 {
		goto L14
	} else {
		goto L359
	}
L357:
	;
	goto L358
L358:
	;
	v1370 = *(*int32)(unsafe.Add(mBase, _consts[131]))
	if v1370 != 0 {
		goto L360
	} else {
		goto L361
	}
L359:
	;
	*(*int32)(unsafe.Add(mBase, _consts[131])) = v1366
	goto L358
L360:
	;
	v1372 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	v1377 = *(*int32)(unsafe.Add(mBase, uint32(v1372<<(uint(int32(2))%32))+uint32(_consts[132])))
	goto L363
L361:
	;
	goto L362
L362:
	;
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(v96)+64))
	v1381 = m.G0
	v1383 = v1381 - int32(48)
	m.G0 = v1383
	*(*int32)(unsafe.Add(mBase, _consts[133])) = v1380
	if v1380 != 0 {
		goto L365
	} else {
		goto L366
	}
L363:
	;
	F_InitializeSystemUser(m, v1370, v1377)
	mBase = m.M
	v1379 = m.ExcPending
	if v1379 != 0 {
		goto L14
	} else {
		goto L364
	}
L364:
	;
	goto L362
L365:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1383)+16)) = int64(103079215120)
	v1392 = *(*int32)(unsafe.Add(mBase, _consts[134]))
	v1394 = F_hash_create(m, int32(314454), v1392, v1383, int32(40))
	mBase = m.M
	v1395 = m.ExcPending
	if v1395 != 0 {
		goto L14
	} else {
		goto L368
	}
L366:
	;
	goto L367
L367:
	;
	m.G0 = v1383 + int32(48)
	v1401 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[84])) = uint8(v1401)
	v1405 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	v1406 = *(*int32)(unsafe.Add(mBase, uint32(v1405)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v1405)+72)) = v1406 + int32(1)
	goto L369
L368:
	;
	*(*int32)(unsafe.Add(mBase, _consts[135])) = v1394
	goto L367
L369:
	;
	m.T0[v372].(func(*base.Module, int32, int32))(m, v83, v90)
	mBase = m.M
	v1411 = m.ExcPending
	if v1411 != 0 {
		goto L14
	} else {
		goto L370
	}
L370:
	;
	v1414 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	v1415 = *(*int32)(unsafe.Add(mBase, uint32(v1414)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v1414)+72)) = v1415 - int32(1)
	goto L371
L371:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v1420 = m.ExcPending
	if v1420 != 0 {
		goto L14
	} else {
		goto L372
	}
L372:
	;
	F_CommitTransaction(m)
	mBase = m.M
	v1422 = m.ExcPending
	if v1422 != 0 {
		goto L14
	} else {
		goto L373
	}
L373:
	;
	v1424 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	*(*int32)(unsafe.Add(mBase, uint32(v1424)+24)) = int32(0)
	v1428 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	v1429 = *(*int32)(unsafe.Add(mBase, uint32(v1428)))
	F_dsm_detach(m, v1429)
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L14
	} else {
		goto L374
	}
L374:
	;
	v1432 = int32(4374764)
	v1433 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	*(*int32)(unsafe.Add(mBase, uint32(v1433))) = int32(0)
	v1437 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	v1438 = *(*int32)(unsafe.Add(mBase, uint32(v1437)+4))
	F_dsa_detach(m, v1438)
	mBase = m.M
	v1440 = m.ExcPending
	if v1440 != 0 {
		goto L14
	} else {
		goto L375
	}
L375:
	;
	v1442 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	v1443 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1442)+4)) = v1443
	v1449 = *(*int32)(unsafe.Add(mBase, _consts[136]))
	v1450 = *(*int32)(unsafe.Add(mBase, uint32(v1449)+16))
	v1451 = m.T0[v1450].(func(*base.Module, int32, int32, int32) int32)(m, int32(88), v1443, v1443)
	mBase = m.M
	v1452 = m.ExcPending
	if v1452 != 0 {
		goto L14
	} else {
		goto L376
	}
L376:
	;
	goto L46
L377:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1478 = m.ExcPending
	if v1478 != 0 {
		goto L14
	} else {
		goto L378
	}
L378:
	;
	F_errmsg(m, int32(94419), int32(0))
	mBase = m.M
	v1482 = m.ExcPending
	if v1482 != 0 {
		goto L14
	} else {
		goto L379
	}
L379:
	;
	F_errfinish(m, int32(492475), int32(1355), int32(276055))
	mBase = m.M
	v1487 = m.ExcPending
	if v1487 != 0 {
		goto L14
	} else {
		goto L380
	}
L380:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L381:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1494 = m.ExcPending
	if v1494 != 0 {
		goto L14
	} else {
		goto L382
	}
L382:
	;
	F_errmsg(m, int32(94463), int32(0))
	mBase = m.M
	v1498 = m.ExcPending
	if v1498 != 0 {
		goto L14
	} else {
		goto L383
	}
L383:
	;
	F_errfinish(m, int32(492475), int32(1360), int32(276055))
	mBase = m.M
	v1503 = m.ExcPending
	if v1503 != 0 {
		goto L14
	} else {
		goto L384
	}
L384:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L385:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v191
	F_errmsg_internal(m, int32(419737), v19+int32(16))
	mBase = m.M
	v1513 = m.ExcPending
	if v1513 != 0 {
		goto L14
	} else {
		goto L386
	}
L386:
	;
	F_errfinish(m, int32(492475), int32(1666), int32(252204))
	mBase = m.M
	v1518 = m.ExcPending
	if v1518 != 0 {
		goto L14
	} else {
		goto L387
	}
L387:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L388:
	;
	v1524 = *(*int32)(unsafe.Add(mBase, _consts[109]))
	v1525 = *(*int32)(unsafe.Add(mBase, uint32(v1524)+4))
	goto L389
L389:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v1525
	F_errmsg_internal(m, int32(449999), v19)
	mBase = m.M
	v1529 = m.ExcPending
	if v1529 != 0 {
		goto L14
	} else {
		goto L390
	}
L390:
	;
	F_errfinish(m, int32(492475), int32(1452), int32(276055))
	mBase = m.M
	v1534 = m.ExcPending
	if v1534 != 0 {
		goto L14
	} else {
		goto L391
	}
L391:
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
