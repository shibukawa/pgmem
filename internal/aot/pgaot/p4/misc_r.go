package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_RecordKnownAssignedTransactionIds(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v41 int32
	_ = v41
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v11 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v11 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_RecordKnownAssignedTransactionIds[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v15
	F_errmsg_internal(m, int32(_a_F_RecordKnownAssignedTransactionIds_0), v7)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v25 = int32(3)
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_RecordKnownAssignedTransactionIds[0]))
	if base.B2i32(base.Ui32(l0) < base.Ui32(v25))|base.B2i32(base.Ui32(v28) < base.Ui32(v25)) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	F_errfinish(m, int32(_a_F_RecordKnownAssignedTransactionIds_1), int32(_a_F_RecordKnownAssignedTransactionIds_2), int32(_a_F_RecordKnownAssignedTransactionIds_3))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	m.G0 = v7 + int32(16)
	return
L9:
	;
	v41 = v28
	goto L15
L10:
	;
	if int32(0) < l0-v28 {
		goto L9
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	if base.Ui32(l0) <= base.Ui32(v28) {
		goto L8
	} else {
		goto L14
	}
L13:
	;
	goto L8
L14:
	;
	goto L9
L15:
	;
	if base.B2i32(base.Ui32(l0) < base.Ui32(int32(3)))|base.B2i32(base.Ui32(v41) < base.Ui32(int32(3))) == int32(0) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_RecordKnownAssignedTransactionIds[1]))
	if base.Ui32(v62) <= base.Ui32(int32(1)) {
		goto L28
	} else {
		goto L29
	}
L17:
	;
	goto L16
L18:
	;
	v53 = int32(3)
	v55 = v41 + int32(1)
	if base.Ui32(v55) <= base.Ui32(v53) {
		goto L24
	} else {
		goto L25
	}
L19:
	;
	if v41-l0 < int32(0) {
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if base.Ui32(l0) <= base.Ui32(v41) {
		goto L17
	} else {
		goto L23
	}
L22:
	;
	goto L17
L23:
	;
	goto L18
L24:
	;
	v58 = v53
	goto L26
L25:
	;
	v58 = v55
	goto L26
L26:
	;
	F_ExtendSUBTRANS(m, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v41 = v58
	goto L15
L28:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RecordKnownAssignedTransactionIds[0])) = l0
	goto L8
L29:
	;
	goto L30
L30:
	;
	v67 = int32(3)
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_RecordKnownAssignedTransactionIds[0]))
	v71 = v69 + int32(1)
	if base.Ui32(v71) <= base.Ui32(v67) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v74 = v67
	goto L33
L32:
	;
	v74 = v71
	goto L33
L33:
	;
	F_KnownAssignedXidsAdd(m, v74, l0, int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RecordKnownAssignedTransactionIds[0])) = l0
	F_AdvanceNextFullTransactionIdPastXid(m, l0)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	goto L8
}
func F_RegisterDynamicBackgroundWorker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v76 int64
	_ = v76
	var v78 int64
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	v3 = int32(0)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RegisterDynamicBackgroundWorker[0])))
	if v10 != int32(1) {
		v149 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v149
L2:
	;
	v14 = F_SanityCheckBackgroundWorker(m, l0, int32(21))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	if v14 == int32(0) {
		v149 = v3
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterDynamicBackgroundWorker[1]))
	v26 = F_LWLockAcquire(m, v22+int32(_a_F_RegisterDynamicBackgroundWorker_0), int32(0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterDynamicBackgroundWorker[2]))
	v31 = v20 & int32(16)
	if v31 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if int32(0) < v48 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterDynamicBackgroundWorker[3]))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	if base.Ui32(v36-v37) < base.Ui32(v35) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterDynamicBackgroundWorker[1]))
	F_LWLockRelease(m, v41+int32(_a_F_RegisterDynamicBackgroundWorker_0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	return int32(0)
L11:
	;
	v57 = int32(0)
	goto L14
L12:
	;
	goto L13
L13:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterDynamicBackgroundWorker[1]))
	F_LWLockRelease(m, v140+int32(_a_F_RegisterDynamicBackgroundWorker_0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L3
	} else {
		goto L30
	}
L14:
	;
	v64 = v29 + int32(16) + v57*int32(1488)
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	if v65 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L13
L16:
	;
	base.MemoryCopy(m, v64+int32(16), l0, int32(1472))
	*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = int32(-1)
	v74 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v64)+1)) = uint8(v74)
	v76 = *(*int64)(unsafe.Add(mBase, uint32(v64)+8))
	v78 = v76 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v64)+8)) = v78
	if v31 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v128 = v57 + int32(1)
	if v128 != v48 {
		v57 = v128
		goto L14
	} else {
		goto L29
	}
L19:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterDynamicBackgroundWorker[2]))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v81)+4)) = v82 + int32(1)
	goto L21
L20:
	;
	goto L21
L21:
	;
	v87 = int32(0)
	v90 = base.AtomicRmwOr32(m, v87, int32(_a_F_RegisterDynamicBackgroundWorker_1), v87)
	v91 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v64))) = uint8(v91)
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterDynamicBackgroundWorker[1]))
	F_LWLockRelease(m, v95+int32(_a_F_RegisterDynamicBackgroundWorker_0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RegisterDynamicBackgroundWorker[0])))
	if v102 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if l1 == int32(0) {
		v149 = v91
		goto L1
	} else {
		goto L27
	}
L24:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterDynamicBackgroundWorker[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v106+int32(28)))) = int32(1)
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterDynamicBackgroundWorker[5]))
	v115 = F_pgmem_kill(m, v113, int32(10))
	mBase = m.M
	goto L26
L25:
	;
	goto L26
L26:
	;
	goto L23
L27:
	;
	v119 = F_palloc(m, int32(16))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v119))) = v57
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v123)+8)) = v78
	return int32(1)
L29:
	;
	goto L15
L30:
	;
	v149 = int32(0)
	goto L1
}
func F_ReleaseExternalFD(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v3 int32
	_ = v3
	v1 = int32(_a_F_ReleaseExternalFD_0)
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseExternalFD[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReleaseExternalFD[0])) = v3 - int32(1)
	return
}
func F_RelfilenumberMapInvalidateCallback(m *base.Module, l0 int64, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = v8 + int32(12)
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_RelfilenumberMapInvalidateCallback[0]))
	F_hash_seq_init(m, v11, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v16 = F_hash_seq_search(m, v11)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L19
	}
L4:
	;
	if v16 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v20 = v16
	goto L8
L6:
	;
	goto L7
L7:
	;
	m.G0 = v8 + int32(32)
	return
L8:
	;
	if l1 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L7
L10:
	;
	v41 = F_hash_seq_search(m, v8+int32(12))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L17
	}
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_RelfilenumberMapInvalidateCallback[0]))
	v34 = F_hash_search(m, v31, v20, int32(2), int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L15
	}
L12:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	if v25 == int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	if l1 != v25 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	if v34 == int32(0) {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	goto L10
L17:
	;
	if v41 != 0 {
		v20 = v41
		goto L8
	} else {
		goto L18
	}
L18:
	;
	goto L9
L19:
	;
	F_errmsg_internal(m, int32(_a_F_RelfilenumberMapInvalidateCallback_0), int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(_a_F_RelfilenumberMapInvalidateCallback_1), int32(76), int32(_a_F_RelfilenumberMapInvalidateCallback_2))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RememberConstraintForRebuilding(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
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
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
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
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
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
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	if v11 == v3 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L18
	} else {
		goto L44
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L18
	} else {
		goto L41
	}
L3:
	;
	m.G0 = v9 + int32(32)
	return
L4:
	;
	if v50 != 0 {
		goto L3
	} else {
		goto L17
	}
L5:
	;
	v50 = int32(0)
	goto L4
L6:
	;
	goto L7
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v18 <= int32(0) {
		v44 = v3
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v50 = v44
	goto L4
L9:
	;
	v21 = int32(0)
	if v21 < v18 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v24 = v18
	goto L12
L11:
	;
	v24 = v21
	goto L12
L12:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v27 = int32(0)
	goto L13
L13:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v25+v27<<(uint(int32(2))%32))))
	v36 = base.B2i32(v35 == l0)
	if v35 == l0 {
		v44 = v36
		goto L8
	} else {
		goto L15
	}
L14:
	;
	v44 = v36
	goto L8
L15:
	;
	v38 = v27 + int32(1)
	if v38 != v24 {
		v27 = v38
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v52 = int32(0)
	v54 = F_pg_get_constraintdef_worker(m, l0, int32(1), v52, v52)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return
L19:
	;
	v56 = F_get_constraint_type(m, l0)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	if v56 == int32(110) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+112)) = v73
	v75 = F_get_constraint_index(m, l0)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L18
	} else {
		goto L29
	}
L22:
	;
	v61 = F_lcons_oid(m, l0, v58)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L18
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v67 = F_lappend_oid(m, v58, l0)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L18
	} else {
		goto L27
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+108)) = v61
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
	v65 = F_lcons(m, v54, v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L18
	} else {
		goto L26
	}
L26:
	;
	v73 = v65
	goto L21
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+108)) = v67
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
	v71 = F_lappend(m, v70, v54)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L18
	} else {
		goto L28
	}
L28:
	;
	v73 = v71
	goto L21
L29:
	;
	if v75 == int32(0) {
		goto L3
	} else {
		goto L30
	}
L30:
	;
	v79 = F_get_index_isreplident(m, v75)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L18
	} else {
		goto L31
	}
L31:
	;
	if v79 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+124))
	if v81 != 0 {
		goto L2
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v85 = F_get_index_isclustered(m, v75)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L18
	} else {
		goto L37
	}
L35:
	;
	v82 = F_get_rel_name(m, v75)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L18
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+124)) = v82
	goto L34
L37:
	;
	if v85 == int32(0) {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)+128))
	if v89 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v90 = F_get_rel_name(m, v75)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L18
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+128)) = v90
	goto L3
L41:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v104
	F_errmsg_internal(m, int32(_a_F_RememberConstraintForRebuilding_0), v9+int32(16))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L18
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_RememberConstraintForRebuilding_1), int32(_a_F_RememberConstraintForRebuilding_2), int32(_a_F_RememberConstraintForRebuilding_3))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L18
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L44:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v120
	F_errmsg_internal(m, int32(_a_F_RememberConstraintForRebuilding_4), v9)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L18
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_RememberConstraintForRebuilding_1), int32(_a_F_RememberConstraintForRebuilding_5), int32(_a_F_RememberConstraintForRebuilding_6))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L18
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RemoveInheritance(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v14 int32
	_ = v14
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int64
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
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
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v143 int32
	_ = v143
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
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v161 int64
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
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
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v242 int64
	_ = v242
	var v244 int32
	_ = v244
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
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
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
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	v4 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(240)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+119)))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v25 = F_DeleteInheritsTuple(m, v20, v21, l2, v22+int32(4))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v513 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v514 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v513 + v514
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v512 + v514
	F_errmsg(m, int32(_a_F_RemoveInheritance_0), v16+int32(32))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L2
	} else {
		goto L142
	}
L2:
	;
	return
L3:
	;
	if v25 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v58 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L2
	} else {
		goto L12
	}
L7:
	;
	F_errcode(m, int32(16908420))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	if v19 == int32(112) {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v40 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v39 + v40
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v38 + v40
	F_errmsg(m, int32(_a_F_RemoveInheritance_1), v16+int32(48))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	F_errfinish(m, int32(_a_F_RemoveInheritance_2), int32(_a_F_RemoveInheritance_3), int32(_a_F_RemoveInheritance_4))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L12:
	;
	v61 = v16 - int32(-64)
	v65 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
	F_ScanKeyInit(m, v61, int32(1), int32(3), int32(184), v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v69 = int32(1)
	v72 = F_systable_beginscan(m, v58, int32(2659), v69, int32(0), v69, v61)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	v74 = F_systable_getnext(m, v72)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	if v74 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v78 = v74
	goto L19
L17:
	;
	goto L18
L18:
	;
	F_systable_endscan(m, v72)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L2
	} else {
		goto L34
	}
L19:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+22)))
	v91 = v89 + v90
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+91)))
	if v92 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L18
L21:
	;
	v127 = F_systable_getnext(m, v72)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L2
	} else {
		goto L32
	}
L22:
	;
	v93 = int32(*(*int16)(unsafe.Add(mBase, uint32(v91)+94)))
	if v93 <= int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v99 = F_SearchSysCacheExistsAttName(m, v96, v91+int32(4))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	if v99 == int32(0) {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	v103 = F_heap_copytuple(m, v78)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v103)+16))
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+22)))
	v107 = v105 + v106
	v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v107)+94)))
	v110 = v108 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v107)+94)) = uint16(v110)
	if v110&int32(_a_F_RemoveInheritance_5) == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v116 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v107)+92)) = uint8(v116)
	goto L29
L28:
	;
	goto L29
L29:
	;
	F_CatalogTupleUpdate(m, v58, v103+int32(4), v103)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L2
	} else {
		goto L30
	}
L30:
	;
	F_pfree(m, v103)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	goto L21
L32:
	;
	if v127 != 0 {
		v78 = v127
		goto L19
	} else {
		goto L33
	}
L33:
	;
	goto L20
L34:
	;
	F_relation_close(m, v58, int32(3))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v150 = F_build_attrmap_by_name(m, v147, v148, int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	v154 = F_table_open(m, int32(2606), int32(3))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v157 = v16 - int32(-64)
	v161 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1)+56)))
	F_ScanKeyInit(m, v157, int32(9), int32(3), int32(184), v161)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L2
	} else {
		goto L38
	}
L38:
	;
	v165 = int32(1)
	v168 = F_systable_beginscan(m, v154, int32(2665), v165, int32(0), v165, v157)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	v170 = F_systable_getnext(m, v168)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L2
	} else {
		goto L40
	}
L40:
	;
	if v170 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v174 = v170
	v179 = v4
	v180 = v4
	goto L44
L42:
	;
	v229 = v4
	v230 = v4
	goto L43
L43:
	;
	F_systable_endscan(m, v168)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L2
	} else {
		goto L58
	}
L44:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v174)+16))
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185)+22)))
	v187 = v185 + v186
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+106)))
	if v188 != 0 {
		v217 = v179
		v218 = v180
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v229 = v217
	v230 = v218
	goto L43
L46:
	;
	v220 = F_systable_getnext(m, v168)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L2
	} else {
		goto L56
	}
L47:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+72)))
	if v189 == int32(99) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v194 = F_pstrdup(m, v187+int32(4))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L2
	} else {
		goto L51
	}
L49:
	;
	v199 = v179
	v200 = v189
	goto L50
L50:
	;
	if v200&int32(255) != int32(110) {
		v217 = v199
		v218 = v180
		goto L46
	} else {
		goto L53
	}
L51:
	;
	v196 = F_lappend(m, v179, v194)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L2
	} else {
		goto L52
	}
L52:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+72)))
	v199 = v196
	v200 = v198
	goto L50
L53:
	;
	v205 = F_extractNotNullColumn(m, v174)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L2
	} else {
		goto L54
	}
L54:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	v213 = int32(*(*int16)(unsafe.Add(mBase, uint32(v207+v205<<(uint(int32(1))%32)-int32(2)))))
	v214 = F_lappend_int(m, v180, v213)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L2
	} else {
		goto L55
	}
L55:
	;
	v217 = v199
	v218 = v214
	goto L46
L56:
	;
	if v220 != 0 {
		v174 = v220
		v179 = v217
		v180 = v218
		goto L44
	} else {
		goto L57
	}
L57:
	;
	goto L45
L58:
	;
	v238 = v16 - int32(-64)
	v242 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
	F_ScanKeyInit(m, v238, int32(9), int32(3), int32(184), v242)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L2
	} else {
		goto L59
	}
L59:
	;
	v246 = int32(1)
	v249 = F_systable_beginscan(m, v154, int32(2665), v246, int32(0), v246, v238)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L2
	} else {
		goto L61
	}
L60:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L2
	} else {
		goto L139
	}
L61:
	;
	v251 = F_systable_getnext(m, v249)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L2
	} else {
		goto L62
	}
L62:
	;
	if v251 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v256 = v251
	v260 = v229
	v261 = v230
	goto L66
L64:
	;
	v431 = v229
	v432 = v230
	goto L65
L65:
	;
	if v431|v432 != 0 {
		goto L117
	} else {
		goto L118
	}
L66:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v256)+16))
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266)+22)))
	v268 = v266 + v267
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268)+72)))
	switch v269 - int32(99) {
	case 0:
		goto L71
	default:
		v416 = v260
		v417 = v261
		goto L68
	case 11:
		goto L70
	}
L67:
	;
	v431 = v416
	v432 = v417
	goto L65
L68:
	;
	v422 = F_systable_getnext(m, v249)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L2
	} else {
		goto L115
	}
L69:
	;
	v386 = F_heap_copytuple(m, v256)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L2
	} else {
		goto L108
	}
L70:
	;
	v336 = F_extractNotNullColumn(m, v256)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L2
	} else {
		goto L93
	}
L71:
	;
	if v260 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v416 = int32(0)
	v417 = v261
	goto L68
L73:
	;
	goto L74
L74:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v260)+4))
	if v275 <= int32(0) {
		v416 = v260
		v417 = v261
		goto L68
	} else {
		goto L75
	}
L75:
	;
	v279 = v268 + int32(4)
	v280 = int32(0)
	if v280 < v275 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v283 = v275
	goto L78
L77:
	;
	v283 = v280
	goto L78
L78:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
	v288 = int32(0)
	goto L79
L79:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v284+v288<<(uint(int32(2))%32))))
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279))))
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302))))
	if base.B2i32(v305 == int32(0))|base.B2i32(v305 != v308) != 0 {
		v326 = v305
		v327 = v308
		goto L82
	} else {
		goto L83
	}
L80:
	;
	v416 = v260
	v417 = v261
	goto L68
L81:
	;
	if v326-v327 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L82:
	;
	goto L81
L83:
	;
	v311 = v279
	v312 = v302
	goto L84
L84:
	;
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v312)+1)))
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311)+1)))
	if v316 == int32(0) {
		v326 = v316
		v327 = v315
		goto L82
	} else {
		goto L86
	}
L85:
	;
	v326 = v316
	v327 = v315
	goto L82
L86:
	;
	v319 = int32(1)
	if v316 == v315 {
		v311 = v311 + v319
		v312 = v312 + v319
		goto L84
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	v331 = F_list_delete_nth_cell(m, v260, v288)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L2
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v334 = v288 + int32(1)
	if v334 != v283 {
		v288 = v334
		goto L79
	} else {
		goto L92
	}
L91:
	;
	v380 = v331
	v381 = v261
	goto L69
L92:
	;
	goto L80
L93:
	;
	if v261 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v416 = v260
	v417 = int32(0)
	goto L68
L95:
	;
	goto L96
L96:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v261)+4))
	if v341 <= int32(0) {
		v416 = v260
		v417 = v261
		goto L68
	} else {
		goto L97
	}
L97:
	;
	v344 = int32(0)
	if v344 < v341 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v347 = v341
	goto L100
L99:
	;
	v347 = v344
	goto L100
L100:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v261)+12))
	v352 = int32(0)
	goto L101
L101:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v348+v352<<(uint(int32(2))%32))))
	if v336 == v366 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v416 = v260
	v417 = v261
	goto L68
L103:
	;
	v368 = F_list_delete_nth_cell(m, v261, v352)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L2
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v371 = v352 + int32(1)
	if v371 != v347 {
		v352 = v371
		goto L101
	} else {
		goto L107
	}
L106:
	;
	v380 = v260
	v381 = v368
	goto L69
L107:
	;
	goto L102
L108:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v386)+16))
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388)+22)))
	v390 = v388 + v389
	v391 = int32(*(*int16)(unsafe.Add(mBase, uint32(v390)+104)))
	if v391 <= int32(0) {
		goto L60
	} else {
		goto L109
	}
L109:
	;
	v395 = v391 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v390)+104)) = uint16(v395)
	if v395&int32(_a_F_RemoveInheritance_5) == int32(0) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v401 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v390)+103)) = uint8(v401)
	goto L112
L111:
	;
	goto L112
L112:
	;
	F_CatalogTupleUpdate(m, v154, v386+int32(4), v386)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L2
	} else {
		goto L113
	}
L113:
	;
	F_pfree(m, v386)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L2
	} else {
		goto L114
	}
L114:
	;
	v416 = v380
	v417 = v381
	goto L68
L115:
	;
	if v422 != 0 {
		v256 = v422
		v260 = v416
		v261 = v417
		goto L66
	} else {
		goto L116
	}
L116:
	;
	goto L67
L117:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L2
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	F_systable_endscan(m, v249)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L2
	} else {
		goto L129
	}
L120:
	;
	if v431 != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v431)+4))
	v444 = v443
	goto L123
L122:
	;
	v444 = int32(0)
	goto L123
L123:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	if v432 != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v432)+4))
	v449 = v447
	goto L126
L125:
	;
	v449 = int32(0)
	goto L126
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v449 + v444
	v452 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v446 + v452
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v445 + v452
	F_errmsg_internal(m, int32(_a_F_RemoveInheritance_6), v16)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L2
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(_a_F_RemoveInheritance_2), int32(_a_F_RemoveInheritance_7), int32(_a_F_RemoveInheritance_4))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L2
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
	F_relation_close(m, v154, int32(3))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L2
	} else {
		goto L130
	}
L130:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	if v19 == int32(112) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v478 = int32(97)
	goto L133
L132:
	;
	v478 = int32(110)
	goto L133
L133:
	;
	F_drop_parent_dependency(m, v471, int32(1259), v473, v478)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L2
	} else {
		goto L134
	}
L134:
	;
	v482 = *(*int32)(unsafe.Add(mBase, _c_F_RemoveInheritance[0]))
	if v482 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v485 = int32(0)
	v486 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	F_RunObjectPostAlterHook(m, int32(2611), v484, v485, v486, v485)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L2
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	m.G0 = v16 + int32(240)
	return
L138:
	;
	goto L137
L139:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v390 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v497
	F_errmsg_internal(m, int32(_a_F_RemoveInheritance_8), v16+int32(16))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L2
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_RemoveInheritance_2), int32(_a_F_RemoveInheritance_9), int32(_a_F_RemoveInheritance_4))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L2
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
	F_errfinish(m, int32(_a_F_RemoveInheritance_2), int32(_a_F_RemoveInheritance_10), int32(_a_F_RemoveInheritance_4))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L2
	} else {
		goto L143
	}
L143:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RepackCommandAsString(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v3 = l0 - int32(1)
	if base.Ui32(v3) <= base.Ui32(int32(2)) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v3<<(uint(int32(2))%32))+uint32(_c_F_RepackCommandAsString[0])))
		v10 = v8
	} else {
		v10 = int32(_a_F_RepackCommandAsString_0)
	}
	return v10
}
func F_RepackWorkerShutdown(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
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
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v5 = base.I32_wrap_i64(l1)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+24))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+116))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v6)+112))
	F_dsm_detach(m, v5)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		v12 = F_SendProcSignal(m, v8, int32(8), v7)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			return
		}
	}
}
func F_ReportApplyConflict(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int64
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
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
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v426 int32
	_ = v426
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v459 int32
	_ = v459
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v484 int32
	_ = v484
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v508 int32
	_ = v508
	var v517 int32
	_ = v517
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v627 int32
	_ = v627
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v648 int32
	_ = v648
	var v656 int32
	_ = v656
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v679 int32
	_ = v679
	var v688 int32
	_ = v688
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v730 int32
	_ = v730
	var v734 int32
	_ = v734
	var v743 int32
	_ = v743
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v788 int32
	_ = v788
	var v794 int32
	_ = v794
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
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
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v837 int32
	_ = v837
	var v842 int32
	_ = v842
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v884 int32
	_ = v884
	var v888 int32
	_ = v888
	var v897 int32
	_ = v897
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v942 int32
	_ = v942
	var v948 int32
	_ = v948
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v994 int32
	_ = v994
	var v998 int32
	_ = v998
	var v1007 int32
	_ = v1007
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1050 int32
	_ = v1050
	var v1058 int32
	_ = v1058
	var v1067 int32
	_ = v1067
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1079 int32
	_ = v1079
	var v1087 int32
	_ = v1087
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1102 int32
	_ = v1102
	var v1110 int32
	_ = v1110
	var v1119 int32
	_ = v1119
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1145 int32
	_ = v1145
	var v1158 int32
	_ = v1158
	var v1162 int32
	_ = v1162
	var v1171 int32
	_ = v1171
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1216 int32
	_ = v1216
	var v1222 int32
	_ = v1222
	var v1247 int32
	_ = v1247
	var v1254 int32
	_ = v1254
	var v1257 int32
	_ = v1257
	var v1259 int32
	_ = v1259
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1289 int32
	_ = v1289
	var v1291 int32
	_ = v1291
	var v1292 int64
	_ = v1292
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1299 int32
	_ = v1299
	var v1300 int64
	_ = v1300
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1311 int32
	_ = v1311
	var v1313 int32
	_ = v1313
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
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1336 int32
	_ = v1336
	var v1341 int32
	_ = v1341
	v25 = m.G0
	v27 = v25 - int32(672)
	m.G0 = v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_initStringInfo(m, v27+int32(604))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if l6 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v1289 = int32(0)
	v1291 = *(*int32)(unsafe.Add(mBase, _c_F_ReportApplyConflict[0]))
	v1292 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1291)+4)))
	v1294 = F_pgstat_prep_pending_entry(m, int32(5), v1289, v1292, v1289)
	mBase = m.M
	v1295 = m.ExcPending
	if v1295 != 0 {
		goto L1
	} else {
		goto L287
	}
L4:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v36 <= int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v62 = int32(0)
	goto L6
L6:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v67+v62<<(uint(int32(2))%32))))
	v72 = *(*int64)(unsafe.Add(mBase, uint32(v71)+16))
	v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+12)))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+52))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v77)+56))
	v80 = int32(0)
	if base.B2i32(base.Ui32(int32(7)) < base.Ui32(l3))|base.B2i32(int32(1)<<(uint(l3)%32)&int32(133) == v80) != 0 {
		v99 = v80
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L3
L8:
	;
	v100 = int32(0)
	if v76 == v100 {
		v116 = v100
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v88 = F_build_index_value_desc(m, l0, v77, v76, v75)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v88 == int32(0) {
		v99 = v80
		goto L8
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+592)) = v88
	v96 = F_psprintf(m, int32(_a_F_ReportApplyConflict_0), v27+int32(592))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v99 = v96
	goto L8
L13:
	;
	if l5 == int32(0) {
		v135 = v100
		goto L18
	} else {
		goto L19
	}
L14:
	;
	v105 = F_ExecBuildSlotValueDescription(m, v79, v76, v78, int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if v105 == int32(0) {
		v116 = v100
		goto L13
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+576)) = v105
	v113 = F_psprintf(m, int32(_a_F_ReportApplyConflict_1), v27+int32(576))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v116 = v113
	goto L13
L18:
	;
	v137 = int32(0)
	if l4 == v137 {
		v165 = v137
		goto L26
	} else {
		goto L27
	}
L19:
	;
	v119 = F_ExecGetInsertedCols(m, l1, l0)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v121 = F_ExecGetUpdatedCols(m, l1, l0)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v123 = F_bms_union(m, v119, v121)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v125 = F_ExecBuildSlotValueDescription(m, v79, l5, v78, v123)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	if v125 == int32(0) {
		v135 = v100
		goto L18
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+560)) = v125
	v133 = F_psprintf(m, int32(_a_F_ReportApplyConflict_2), v27+int32(560))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v135 = v133
	goto L18
L26:
	;
	F_initStringInfo(m, v27+int32(656))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L38
	}
L27:
	;
	v140 = F_GetRelationIdentityOrPK(m, v77)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	if v140 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v145 = F_ExecBuildSlotValueDescription(m, v79, l4, v78, int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v155 = F_build_index_value_desc(m, l0, v77, l4, v140)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L35
	}
L32:
	;
	if v145 == int32(0) {
		v165 = v137
		goto L26
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+528)) = v145
	v153 = F_psprintf(m, int32(_a_F_ReportApplyConflict_3), v27+int32(528))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v165 = v153
	goto L26
L35:
	;
	if v155 == int32(0) {
		v165 = v137
		goto L26
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+544)) = v155
	v163 = F_psprintf(m, int32(_a_F_ReportApplyConflict_4), v27+int32(544))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v165 = v163
	goto L26
L38:
	;
	F_initStringInfo(m, v27+int32(640))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	switch l3 {
	case 0, 2, 7:
		goto L46
	case 1:
		goto L45
	case 3:
		goto L44
	case 4:
		goto L43
	case 5:
		goto L42
	case 6:
		goto L41
	default:
		goto L40
	}
L40:
	;
	v1247 = *(*int32)(unsafe.Add(mBase, uint32(v27)+608))
	if int32(0) < v1247 {
		goto L281
	} else {
		goto L282
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+524)) = v165
	*(*int32)(unsafe.Add(mBase, uint32(v27)+636)) = v165
	v1125 = F_list_make1_impl(m, int32(1), v27+int32(524))
	mBase = m.M
	v1126 = m.ExcPending
	if v1126 != 0 {
		goto L1
	} else {
		goto L262
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+628)) = v135
	*(*int32)(unsafe.Add(mBase, uint32(v27)+632)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v27)+624)) = v165
	*(*int32)(unsafe.Add(mBase, uint32(v27)+508)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v27)+504)) = v135
	*(*int32)(unsafe.Add(mBase, uint32(v27)+500)) = v165
	v961 = F_list_make3_impl(m, v27+int32(508), v27+int32(504), v27+int32(500))
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L1
	} else {
		goto L222
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+632)) = v165
	*(*int32)(unsafe.Add(mBase, uint32(v27)+636)) = v135
	*(*int32)(unsafe.Add(mBase, uint32(v27)+396)) = v135
	*(*int32)(unsafe.Add(mBase, uint32(v27)+392)) = v165
	v851 = F_list_make2_impl(m, v27+int32(396), v27+int32(392))
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L1
	} else {
		goto L202
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+628)) = v165
	*(*int32)(unsafe.Add(mBase, uint32(v27)+632)) = v135
	*(*int32)(unsafe.Add(mBase, uint32(v27)+380)) = v135
	*(*int32)(unsafe.Add(mBase, uint32(v27)+376)) = v165
	v697 = F_list_make2_impl(m, v27+int32(380), v27+int32(376))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L1
	} else {
		goto L164
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+628)) = v135
	*(*int32)(unsafe.Add(mBase, uint32(v27)+632)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v27)+624)) = v165
	*(*int32)(unsafe.Add(mBase, uint32(v27)+316)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v27)+312)) = v135
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v165
	v530 = F_list_make3_impl(m, v27+int32(316), v27+int32(312), v27+int32(308))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L1
	} else {
		goto L124
	}
L46:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v27)+608))
	if v176 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+628)) = v165
	*(*int32)(unsafe.Add(mBase, uint32(v27)+632)) = v135
	*(*int32)(unsafe.Add(mBase, uint32(v27)+204)) = v135
	*(*int32)(unsafe.Add(mBase, uint32(v27)+200)) = v165
	v187 = F_list_make2_impl(m, v27+int32(204), v27+int32(200))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+620)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v27)+624)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v27)+188)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v27)+184)) = v116
	v324 = F_list_make2_impl(m, v27+int32(188), v27+int32(184))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L73
	}
L50:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v27)+644))
	if v267 != 0 {
		goto L66
	} else {
		goto L67
	}
L51:
	;
	if v187 == int32(0) {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v192 = int32(0)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
	if v193 <= v192 {
		goto L50
	} else {
		goto L53
	}
L53:
	;
	v204 = v193
	v207 = v192
	v211 = int32(1)
	goto L54
L54:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v187)+12))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v220+v207<<(uint(int32(2))%32))))
	if v224 != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L50
L56:
	;
	if v211 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	v238 = v204
	v239 = v211
	goto L58
L58:
	;
	v241 = v207 + int32(1)
	if v241 < v238 {
		v204 = v238
		v207 = v241
		v211 = v239
		goto L54
	} else {
		goto L64
	}
L59:
	;
	F_appendStringInfoString(m, v27+int32(640), int32(_a_F_ReportApplyConflict_5))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	F_appendStringInfoString(m, v27+int32(640), v224)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L63
	}
L62:
	;
	goto L61
L63:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
	v238 = v237
	v239 = int32(0)
	goto L58
L64:
	;
	goto L55
L65:
	;
	v284 = v27 + int32(640)
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v284)))
	v286 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v285))) = uint8(v286)
	*(*int32)(unsafe.Add(mBase, uint32(v284)+12)) = v286
	*(*int32)(unsafe.Add(mBase, uint32(v284)+4)) = v286
	goto L71
L66:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v27)+640))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+192)) = v268
	F_appendStringInfo(m, v27+int32(656), int32(_a_F_ReportApplyConflict_6), v27+int32(192))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	F_appendStringInfo(m, v27+int32(656), int32(_a_F_ReportApplyConflict_7), int32(0))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L70
	}
L69:
	;
	goto L65
L70:
	;
	goto L65
L71:
	;
	goto L49
L72:
	;
	if v72 != int64(0) {
		goto L87
	} else {
		goto L88
	}
L73:
	;
	if v324 == int32(0) {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v329 = int32(0)
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v324)+4))
	if v330 <= v329 {
		goto L72
	} else {
		goto L75
	}
L75:
	;
	v341 = int32(1)
	v342 = v330
	v344 = v329
	goto L76
L76:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v324)+12))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v357+v344<<(uint(int32(2))%32))))
	if v361 != 0 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	goto L72
L78:
	;
	if v341&int32(1) == int32(0) {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	v377 = v341
	v378 = v342
	goto L80
L80:
	;
	v380 = v344 + int32(1)
	if v380 < v378 {
		v341 = v377
		v342 = v378
		v344 = v380
		goto L76
	} else {
		goto L86
	}
L81:
	;
	F_appendStringInfoString(m, v27+int32(640), int32(_a_F_ReportApplyConflict_5))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	F_appendStringInfoString(m, v27+int32(640), v361)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L85
	}
L84:
	;
	goto L83
L85:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v324)+4))
	v377 = int32(0)
	v378 = v375
	goto L80
L86:
	;
	goto L77
L87:
	;
	if v73 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	goto L89
L89:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v27)+644))
	v496 = F_get_rel_name(m, v75)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L1
	} else {
		goto L117
	}
L90:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v27)+644))
	v411 = F_get_rel_name(m, v75)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v439 = F_replorigin_by_oid(m, v73, v27+int32(636))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L1
	} else {
		goto L100
	}
L93:
	;
	v413 = F_timestamptz_to_str(m, v72)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	if v410 != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+80)) = v411
	*(*int32)(unsafe.Add(mBase, uint32(v27)+84)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v27)+88)) = v413
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v27)+640))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+92)) = v418
	F_appendStringInfo(m, v27+int32(656), int32(_a_F_ReportApplyConflict_8), v27+int32(80))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+72)) = v413
	*(*int32)(unsafe.Add(mBase, uint32(v27)+68)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v27)+64)) = v411
	F_appendStringInfo(m, v27+int32(656), int32(_a_F_ReportApplyConflict_9), v27-int32(-64))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L1
	} else {
		goto L99
	}
L98:
	;
	goto L40
L99:
	;
	goto L40
L100:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v27)+644))
	v442 = F_get_rel_name(m, v75)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	if v439 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v27)+636))
	v445 = F_timestamptz_to_str(m, v72)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v471 = F_timestamptz_to_str(m, v72)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L1
	} else {
		goto L111
	}
L105:
	;
	if v441 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v27)+640))
	*(*int32)(unsafe.Add(mBase, uint32(v27+int32(128)))) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v27)+112)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v27)+116)) = v444
	*(*int32)(unsafe.Add(mBase, uint32(v27)+120)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v27)+124)) = v445
	F_appendStringInfo(m, v27+int32(656), int32(_a_F_ReportApplyConflict_10), v27+int32(112))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+108)) = v445
	*(*int32)(unsafe.Add(mBase, uint32(v27)+104)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v27)+100)) = v444
	*(*int32)(unsafe.Add(mBase, uint32(v27)+96)) = v442
	F_appendStringInfo(m, v27+int32(656), int32(_a_F_ReportApplyConflict_11), v27+int32(96))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L1
	} else {
		goto L110
	}
L109:
	;
	goto L40
L110:
	;
	goto L40
L111:
	;
	if v441 != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+160)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v27)+164)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v27)+168)) = v471
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v27)+640))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+172)) = v476
	F_appendStringInfo(m, v27+int32(656), int32(_a_F_ReportApplyConflict_12), v27+int32(160))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L1
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+152)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v27)+148)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v27)+144)) = v442
	F_appendStringInfo(m, v27+int32(656), int32(_a_F_ReportApplyConflict_13), v27+int32(144))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L116
	}
L115:
	;
	goto L40
L116:
	;
	goto L40
L117:
	;
	if v495 != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+48)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v27)+52)) = v74
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v27)+640))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+56)) = v500
	F_appendStringInfo(m, v27+int32(656), int32(_a_F_ReportApplyConflict_14), v27+int32(48))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L1
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+36)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v496
	F_appendStringInfo(m, v27+int32(656), int32(_a_F_ReportApplyConflict_15), v27+int32(32))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L1
	} else {
		goto L122
	}
L121:
	;
	goto L40
L122:
	;
	goto L40
L123:
	;
	if v73 == int32(0) {
		goto L138
	} else {
		goto L139
	}
L124:
	;
	if v530 == int32(0) {
		goto L123
	} else {
		goto L125
	}
L125:
	;
	v535 = int32(0)
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v530)+4))
	if v536 <= v535 {
		goto L123
	} else {
		goto L126
	}
L126:
	;
	v547 = int32(1)
	v548 = v536
	v550 = v535
	goto L127
L127:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v530)+12))
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v563+v550<<(uint(int32(2))%32))))
	if v567 != 0 {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	goto L123
L129:
	;
	if v547&int32(1) == int32(0) {
		goto L132
	} else {
		goto L133
	}
L130:
	;
	v583 = v547
	v584 = v548
	goto L131
L131:
	;
	v586 = v550 + int32(1)
	if v586 < v584 {
		v547 = v583
		v548 = v584
		v550 = v586
		goto L127
	} else {
		goto L137
	}
L132:
	;
	F_appendStringInfoString(m, v27+int32(640), int32(_a_F_ReportApplyConflict_5))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L1
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	F_appendStringInfoString(m, v27+int32(640), v567)
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L1
	} else {
		goto L136
	}
L135:
	;
	goto L134
L136:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v530)+4))
	v583 = int32(0)
	v584 = v581
	goto L131
L137:
	;
	goto L128
L138:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v27)+644))
	v615 = F_timestamptz_to_str(m, v72)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L1
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	v639 = F_replorigin_by_oid(m, v73, v27+int32(636))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L1
	} else {
		goto L147
	}
L141:
	;
	if v614 != 0 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+224)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v27)+228)) = v615
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v27)+640))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+232)) = v619
	F_appendStringInfo(m, v27+int32(656), int32(_a_F_ReportApplyConflict_16), v27+int32(224))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L1
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+212)) = v615
	*(*int32)(unsafe.Add(mBase, uint32(v27)+208)) = v74
	F_appendStringInfo(m, v27+int32(656), int32(_a_F_ReportApplyConflict_17), v27+int32(208))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L1
	} else {
		goto L146
	}
L145:
	;
	goto L40
L146:
	;
	goto L40
L147:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v27)+644))
	if v639 != 0 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v27)+636))
	v643 = F_timestamptz_to_str(m, v72)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L1
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	v667 = F_timestamptz_to_str(m, v72)
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L1
	} else {
		goto L157
	}
L151:
	;
	if v641 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v642
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v643
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v27)+640))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v648
	F_appendStringInfo(m, v27+int32(656), int32(_a_F_ReportApplyConflict_18), v27+int32(256))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L1
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+248)) = v643
	*(*int32)(unsafe.Add(mBase, uint32(v27)+244)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v27)+240)) = v642
	F_appendStringInfo(m, v27+int32(656), int32(_a_F_ReportApplyConflict_19), v27+int32(240))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L1
	} else {
		goto L156
	}
L155:
	;
	goto L40
L156:
	;
	goto L40
L157:
	;
	if v641 != 0 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+288)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v27)+292)) = v667
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v27)+640))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+296)) = v671
	F_appendStringInfo(m, v27+int32(656), int32(_a_F_ReportApplyConflict_20), v27+int32(288))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L1
	} else {
		goto L161
	}
L159:
	;
	goto L160
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+276)) = v667
	*(*int32)(unsafe.Add(mBase, uint32(v27)+272)) = v74
	F_appendStringInfo(m, v27+int32(656), int32(_a_F_ReportApplyConflict_21), v27+int32(272))
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L1
	} else {
		goto L162
	}
L161:
	;
	goto L40
L162:
	;
	goto L40
L163:
	;
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v27)+644))
	if v779 != 0 {
		goto L179
	} else {
		goto L180
	}
L164:
	;
	if v697 == int32(0) {
		goto L163
	} else {
		goto L165
	}
L165:
	;
	v702 = int32(0)
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v697)+4))
	if v703 <= v702 {
		goto L163
	} else {
		goto L166
	}
L166:
	;
	v714 = int32(1)
	v715 = v703
	v717 = v702
	goto L167
L167:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v697)+12))
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v730+v717<<(uint(int32(2))%32))))
	if v734 != 0 {
		goto L169
	} else {
		goto L170
	}
L168:
	;
	goto L163
L169:
	;
	if v714&int32(1) == int32(0) {
		goto L172
	} else {
		goto L173
	}
L170:
	;
	v750 = v714
	v751 = v715
	goto L171
L171:
	;
	v753 = v717 + int32(1)
	if v753 < v751 {
		v714 = v750
		v715 = v751
		v717 = v753
		goto L167
	} else {
		goto L177
	}
L172:
	;
	F_appendStringInfoString(m, v27+int32(640), int32(_a_F_ReportApplyConflict_5))
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L1
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	F_appendStringInfoString(m, v27+int32(640), v734)
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L1
	} else {
		goto L176
	}
L175:
	;
	goto L174
L176:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v697)+4))
	v750 = int32(0)
	v751 = v748
	goto L171
L177:
	;
	goto L168
L178:
	;
	if v72 != int64(0) {
		goto L184
	} else {
		goto L185
	}
L179:
	;
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v27)+640))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+368)) = v780
	F_appendStringInfo(m, v27+int32(656), int32(_a_F_ReportApplyConflict_22), v27+int32(368))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L1
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	F_appendStringInfo(m, v27+int32(656), int32(_a_F_ReportApplyConflict_23), int32(0))
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L1
	} else {
		goto L183
	}
L182:
	;
	goto L178
L183:
	;
	goto L178
L184:
	;
	if v73 == int32(0) {
		goto L187
	} else {
		goto L188
	}
L185:
	;
	goto L186
L186:
	;
	F_appendStringInfoString(m, v27+int32(656), int32(_a_F_ReportApplyConflict_24))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L1
	} else {
		goto L200
	}
L187:
	;
	v799 = F_timestamptz_to_str(m, v72)
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L1
	} else {
		goto L190
	}
L188:
	;
	goto L189
L189:
	;
	v812 = F_replorigin_by_oid(m, v73, v27+int32(636))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L1
	} else {
		goto L192
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+324)) = v799
	*(*int32)(unsafe.Add(mBase, uint32(v27)+320)) = v74
	F_appendStringInfo(m, v27+int32(656), int32(_a_F_ReportApplyConflict_25), v27+int32(320))
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	goto L40
L192:
	;
	if v812 != 0 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v27)+636))
	v815 = F_timestamptz_to_str(m, v72)
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L1
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	v827 = F_timestamptz_to_str(m, v72)
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L1
	} else {
		goto L198
	}
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+344)) = v815
	*(*int32)(unsafe.Add(mBase, uint32(v27)+340)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v27)+336)) = v814
	F_appendStringInfo(m, v27+int32(656), int32(_a_F_ReportApplyConflict_26), v27+int32(336))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	goto L40
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+356)) = v827
	*(*int32)(unsafe.Add(mBase, uint32(v27)+352)) = v74
	F_appendStringInfo(m, v27+int32(656), int32(_a_F_ReportApplyConflict_27), v27+int32(352))
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L1
	} else {
		goto L199
	}
L199:
	;
	goto L40
L200:
	;
	goto L40
L201:
	;
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v27)+644))
	if v933 != 0 {
		goto L216
	} else {
		goto L217
	}
L202:
	;
	if v851 == int32(0) {
		goto L201
	} else {
		goto L203
	}
L203:
	;
	v856 = int32(0)
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v851)+4))
	if v857 <= v856 {
		goto L201
	} else {
		goto L204
	}
L204:
	;
	v868 = int32(1)
	v869 = v857
	v871 = v856
	goto L205
L205:
	;
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v851)+12))
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v884+v871<<(uint(int32(2))%32))))
	if v888 != 0 {
		goto L207
	} else {
		goto L208
	}
L206:
	;
	goto L201
L207:
	;
	if v868&int32(1) == int32(0) {
		goto L210
	} else {
		goto L211
	}
L208:
	;
	v904 = v868
	v905 = v869
	goto L209
L209:
	;
	v907 = v871 + int32(1)
	if v907 < v905 {
		v868 = v904
		v869 = v905
		v871 = v907
		goto L205
	} else {
		goto L215
	}
L210:
	;
	F_appendStringInfoString(m, v27+int32(640), int32(_a_F_ReportApplyConflict_5))
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L1
	} else {
		goto L213
	}
L211:
	;
	goto L212
L212:
	;
	F_appendStringInfoString(m, v27+int32(640), v888)
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L1
	} else {
		goto L214
	}
L213:
	;
	goto L212
L214:
	;
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v851)+4))
	v904 = int32(0)
	v905 = v902
	goto L209
L215:
	;
	goto L206
L216:
	;
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v27)+640))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+384)) = v934
	F_appendStringInfo(m, v27+int32(656), int32(_a_F_ReportApplyConflict_28), v27+int32(384))
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L1
	} else {
		goto L219
	}
L217:
	;
	goto L218
L218:
	;
	F_appendStringInfo(m, v27+int32(656), int32(_a_F_ReportApplyConflict_29), int32(0))
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L1
	} else {
		goto L220
	}
L219:
	;
	goto L40
L220:
	;
	goto L40
L221:
	;
	if v73 == int32(0) {
		goto L236
	} else {
		goto L237
	}
L222:
	;
	if v961 == int32(0) {
		goto L221
	} else {
		goto L223
	}
L223:
	;
	v966 = int32(0)
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v961)+4))
	if v967 <= v966 {
		goto L221
	} else {
		goto L224
	}
L224:
	;
	v978 = int32(1)
	v979 = v967
	v981 = v966
	goto L225
L225:
	;
	v994 = *(*int32)(unsafe.Add(mBase, uint32(v961)+12))
	v998 = *(*int32)(unsafe.Add(mBase, uint32(v994+v981<<(uint(int32(2))%32))))
	if v998 != 0 {
		goto L227
	} else {
		goto L228
	}
L226:
	;
	goto L221
L227:
	;
	if v978&int32(1) == int32(0) {
		goto L230
	} else {
		goto L231
	}
L228:
	;
	v1014 = v978
	v1015 = v979
	goto L229
L229:
	;
	v1017 = v981 + int32(1)
	if v1017 < v1015 {
		v978 = v1014
		v979 = v1015
		v981 = v1017
		goto L225
	} else {
		goto L235
	}
L230:
	;
	F_appendStringInfoString(m, v27+int32(640), int32(_a_F_ReportApplyConflict_5))
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L1
	} else {
		goto L233
	}
L231:
	;
	goto L232
L232:
	;
	F_appendStringInfoString(m, v27+int32(640), v998)
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L1
	} else {
		goto L234
	}
L233:
	;
	goto L232
L234:
	;
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(v961)+4))
	v1014 = int32(0)
	v1015 = v1012
	goto L229
L235:
	;
	goto L226
L236:
	;
	v1045 = *(*int32)(unsafe.Add(mBase, uint32(v27)+644))
	v1046 = F_timestamptz_to_str(m, v72)
	mBase = m.M
	v1047 = m.ExcPending
	if v1047 != 0 {
		goto L1
	} else {
		goto L239
	}
L237:
	;
	goto L238
L238:
	;
	v1070 = F_replorigin_by_oid(m, v73, v27+int32(636))
	mBase = m.M
	v1071 = m.ExcPending
	if v1071 != 0 {
		goto L1
	} else {
		goto L245
	}
L239:
	;
	if v1045 != 0 {
		goto L240
	} else {
		goto L241
	}
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+416)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v27)+420)) = v1046
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v27)+640))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+424)) = v1050
	F_appendStringInfo(m, v27+int32(656), int32(_a_F_ReportApplyConflict_30), v27+int32(416))
	mBase = m.M
	v1058 = m.ExcPending
	if v1058 != 0 {
		goto L1
	} else {
		goto L243
	}
L241:
	;
	goto L242
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+404)) = v1046
	*(*int32)(unsafe.Add(mBase, uint32(v27)+400)) = v74
	F_appendStringInfo(m, v27+int32(656), int32(_a_F_ReportApplyConflict_31), v27+int32(400))
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L1
	} else {
		goto L244
	}
L243:
	;
	goto L40
L244:
	;
	goto L40
L245:
	;
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v27)+644))
	if v1070 != 0 {
		goto L246
	} else {
		goto L247
	}
L246:
	;
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(v27)+636))
	v1074 = F_timestamptz_to_str(m, v72)
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L1
	} else {
		goto L249
	}
L247:
	;
	goto L248
L248:
	;
	v1098 = F_timestamptz_to_str(m, v72)
	mBase = m.M
	v1099 = m.ExcPending
	if v1099 != 0 {
		goto L1
	} else {
		goto L255
	}
L249:
	;
	if v1072 != 0 {
		goto L250
	} else {
		goto L251
	}
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+448)) = v1073
	*(*int32)(unsafe.Add(mBase, uint32(v27)+452)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v27)+456)) = v1074
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(v27)+640))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+460)) = v1079
	F_appendStringInfo(m, v27+int32(656), int32(_a_F_ReportApplyConflict_32), v27+int32(448))
	mBase = m.M
	v1087 = m.ExcPending
	if v1087 != 0 {
		goto L1
	} else {
		goto L253
	}
L251:
	;
	goto L252
L252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+440)) = v1074
	*(*int32)(unsafe.Add(mBase, uint32(v27)+436)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v27)+432)) = v1073
	F_appendStringInfo(m, v27+int32(656), int32(_a_F_ReportApplyConflict_33), v27+int32(432))
	mBase = m.M
	v1097 = m.ExcPending
	if v1097 != 0 {
		goto L1
	} else {
		goto L254
	}
L253:
	;
	goto L40
L254:
	;
	goto L40
L255:
	;
	if v1072 != 0 {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+480)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v27)+484)) = v1098
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(v27)+640))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+488)) = v1102
	F_appendStringInfo(m, v27+int32(656), int32(_a_F_ReportApplyConflict_34), v27+int32(480))
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L1
	} else {
		goto L259
	}
L257:
	;
	goto L258
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+468)) = v1098
	*(*int32)(unsafe.Add(mBase, uint32(v27)+464)) = v74
	F_appendStringInfo(m, v27+int32(656), int32(_a_F_ReportApplyConflict_35), v27+int32(464))
	mBase = m.M
	v1119 = m.ExcPending
	if v1119 != 0 {
		goto L1
	} else {
		goto L260
	}
L259:
	;
	goto L40
L260:
	;
	goto L40
L261:
	;
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(v27)+644))
	if v1207 != 0 {
		goto L276
	} else {
		goto L277
	}
L262:
	;
	if v1125 == int32(0) {
		goto L261
	} else {
		goto L263
	}
L263:
	;
	v1130 = int32(0)
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(v1125)+4))
	if v1131 <= v1130 {
		goto L261
	} else {
		goto L264
	}
L264:
	;
	v1142 = int32(1)
	v1143 = v1131
	v1145 = v1130
	goto L265
L265:
	;
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(v1125)+12))
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(v1158+v1145<<(uint(int32(2))%32))))
	if v1162 != 0 {
		goto L267
	} else {
		goto L268
	}
L266:
	;
	goto L261
L267:
	;
	if v1142&int32(1) == int32(0) {
		goto L270
	} else {
		goto L271
	}
L268:
	;
	v1178 = v1142
	v1179 = v1143
	goto L269
L269:
	;
	v1181 = v1145 + int32(1)
	if v1181 < v1179 {
		v1142 = v1178
		v1143 = v1179
		v1145 = v1181
		goto L265
	} else {
		goto L275
	}
L270:
	;
	F_appendStringInfoString(m, v27+int32(640), int32(_a_F_ReportApplyConflict_5))
	mBase = m.M
	v1171 = m.ExcPending
	if v1171 != 0 {
		goto L1
	} else {
		goto L273
	}
L271:
	;
	goto L272
L272:
	;
	F_appendStringInfoString(m, v27+int32(640), v1162)
	mBase = m.M
	v1175 = m.ExcPending
	if v1175 != 0 {
		goto L1
	} else {
		goto L274
	}
L273:
	;
	goto L272
L274:
	;
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(v1125)+4))
	v1178 = int32(0)
	v1179 = v1176
	goto L269
L275:
	;
	goto L266
L276:
	;
	v1208 = *(*int32)(unsafe.Add(mBase, uint32(v27)+640))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+512)) = v1208
	F_appendStringInfo(m, v27+int32(656), int32(_a_F_ReportApplyConflict_36), v27+int32(512))
	mBase = m.M
	v1216 = m.ExcPending
	if v1216 != 0 {
		goto L1
	} else {
		goto L279
	}
L277:
	;
	goto L278
L278:
	;
	F_appendStringInfo(m, v27+int32(656), int32(_a_F_ReportApplyConflict_37), int32(0))
	mBase = m.M
	v1222 = m.ExcPending
	if v1222 != 0 {
		goto L1
	} else {
		goto L280
	}
L279:
	;
	goto L40
L280:
	;
	goto L40
L281:
	;
	F_appendStringInfoChar(m, v27+int32(604), int32(10))
	mBase = m.M
	v1254 = m.ExcPending
	if v1254 != 0 {
		goto L1
	} else {
		goto L284
	}
L282:
	;
	goto L283
L283:
	;
	v1257 = *(*int32)(unsafe.Add(mBase, uint32(v27)+656))
	F_appendStringInfoString(m, v27+int32(604), v1257)
	mBase = m.M
	v1259 = m.ExcPending
	if v1259 != 0 {
		goto L1
	} else {
		goto L285
	}
L284:
	;
	goto L283
L285:
	;
	v1261 = v62 + int32(1)
	v1262 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v1261 < v1262 {
		v62 = v1261
		goto L6
	} else {
		goto L286
	}
L286:
	;
	goto L7
L287:
	;
	v1296 = *(*int32)(unsafe.Add(mBase, uint32(v1294)+12))
	v1299 = v1296 + l3<<(uint(int32(3))%32)
	v1300 = *(*int64)(unsafe.Add(mBase, uint32(v1299)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1299)+24)) = v1300 + int64(1)
	v1305 = F_errstart(m, l2, int32(0))
	mBase = m.M
	v1306 = m.ExcPending
	if v1306 != 0 {
		goto L1
	} else {
		goto L288
	}
L288:
	;
	if v1305 != 0 {
		goto L289
	} else {
		goto L290
	}
L289:
	;
	if base.Ui32(l3) <= base.Ui32(int32(7)) {
		goto L292
	} else {
		goto L293
	}
L290:
	;
	goto L291
L291:
	;
	m.G0 = v27 + int32(672)
	return
L292:
	;
	v1311 = *(*int32)(unsafe.Add(mBase, uint32(l3<<(uint(int32(2))%32))+uint32(_c_F_ReportApplyConflict[1])))
	F_errcode(m, v1311)
	mBase = m.M
	v1313 = m.ExcPending
	if v1313 != 0 {
		goto L1
	} else {
		goto L295
	}
L293:
	;
	goto L294
L294:
	;
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(l3<<(uint(int32(2))%32))+uint32(_c_F_ReportApplyConflict[2])))
	v1317 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	v1318 = *(*int32)(unsafe.Add(mBase, uint32(v1317)+68))
	v1319 = F_get_namespace_name(m, v1318)
	mBase = m.M
	v1320 = m.ExcPending
	if v1320 != 0 {
		goto L1
	} else {
		goto L296
	}
L295:
	;
	goto L294
L296:
	;
	v1321 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+24)) = v1316
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v1319
	*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v1321 + int32(4)
	F_errmsg(m, int32(_a_F_ReportApplyConflict_38), v27+int32(16))
	mBase = m.M
	v1331 = m.ExcPending
	if v1331 != 0 {
		goto L1
	} else {
		goto L297
	}
L297:
	;
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(v27)+604))
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v1332
	F_errdetail_internal(m, int32(_a_F_ReportApplyConflict_39), v27)
	mBase = m.M
	v1336 = m.ExcPending
	if v1336 != 0 {
		goto L1
	} else {
		goto L298
	}
L298:
	;
	F_errfinish(m, int32(_a_F_ReportApplyConflict_40), int32(132), int32(_a_F_ReportApplyConflict_41))
	mBase = m.M
	v1341 = m.ExcPending
	if v1341 != 0 {
		goto L1
	} else {
		goto L299
	}
L299:
	;
	goto L291
}
func F_ReservedPLKeywords_hash_func(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
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
	var v94 int32
	_ = v94
	v3 = int32(0)
	if l1 != 0 {
		if l1 == int32(1) {
			v50 = l0
			v51 = int32(0)
			v52 = v3
			v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
			v60 = v58 | int32(32)
			v68 = v51*int32(257) + v60
			v69 = v60 + v52*int32(31)
		} else {
			v17 = l0
			v18 = int32(0)
			v19 = v3
			v24 = v3
			for {
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
				v26 = int32(32)
				v27 = v25 | v26
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
				v30 = v28 | v26
				v31 = int32(31)
				v36 = v27 + (v30+v19*v31)*v31
				v37 = int32(257)
				v42 = (v18*v37+v30)*v37 + v27
				v43 = int32(2)
				v44 = v17 + v43
				v46 = v24 + v43
				if v46 != l1&int32(-2) {
					v17 = v44
					v18 = v42
					v19 = v36
					v24 = v46
					continue
				} else {
					break
				}
				break
			}
			if l1&int32(1) == int32(0) {
				v68 = v42
				v69 = v36
			} else {
				v50 = v44
				v51 = v42
				v52 = v36
				v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
				v60 = v58 | int32(32)
				v68 = v51*int32(257) + v60
				v69 = v60 + v52*int32(31)
			}
		}
		v75 = int32(45)
		v76 = base.I32_rem_u_s(v69, v75)
		v78 = base.I32_rem_u_s(v68, v75)
		v82 = v76
		v88 = v78
	} else {
		v82 = v3
		v88 = int32(0)
	}
	v91 = int32(*(*int8)(unsafe.Add(mBase, uint32(v82)+uint32(_c_F_ReservedPLKeywords_hash_func[0]))))
	v94 = int32(*(*int8)(unsafe.Add(mBase, uint32(v88)+uint32(_c_F_ReservedPLKeywords_hash_func[0]))))
	return v91 + v94
}
func F__readBitmapset(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v46 int64
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v10 = v7 + int32(44)
	v11 = F_pg_strtok(m, v10)
	mBase = m.M
	if v11 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L23
	} else {
		goto L41
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L23
	} else {
		goto L38
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L23
	} else {
		goto L35
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L23
	} else {
		goto L32
	}
L5:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
	if v12 != int32(1) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L23
	} else {
		goto L29
	}
L8:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v15 != int32(40) {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v18 = F_pg_strtok(m, v10)
	mBase = m.M
	if v18 == int32(0) {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
	if v21 != int32(1) {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v24 != int32(98) {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v27 = F_pg_strtok(m, v10)
	mBase = m.M
	if v27 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v29 = v27
	v31 = int32(0)
	goto L16
L14:
	;
	goto L15
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L23
	} else {
		goto L26
	}
L16:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
	if v32 != int32(1) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L15
L18:
	;
	v46 = F_strtox_2(m, v29, v7+int32(40), int32(10), int64(2147483648))
	mBase = m.M
	goto L21
L19:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v35 != int32(41) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	m.G0 = v7 + int32(48)
	return v31
L21:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v7)+40))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
	if v48 != v29+v49 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v52 = F_bms_add_member(m, v31, base.I32_wrap_i64(v46))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	return int32(0)
L24:
	;
	v58 = F_pg_strtok(m, v7+int32(44))
	mBase = m.M
	if v58 != 0 {
		v29 = v58
		v31 = v52
		goto L16
	} else {
		goto L25
	}
L25:
	;
	goto L17
L26:
	;
	F_errmsg_internal(m, int32(_a_F__readBitmapset_0), int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F__readBitmapset_1), int32(232), int32(_a_F__readBitmapset_2))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L23
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L29:
	;
	F_errmsg_internal(m, int32(_a_F__readBitmapset_3), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L23
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F__readBitmapset_1), int32(215), int32(_a_F__readBitmapset_2))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L23
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v11
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v94
	F_errmsg_internal(m, int32(_a_F__readBitmapset_4), v7+int32(32))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L23
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F__readBitmapset_1), int32(217), int32(_a_F__readBitmapset_2))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L23
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	F_errmsg_internal(m, int32(_a_F__readBitmapset_3), int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L23
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F__readBitmapset_1), int32(221), int32(_a_F__readBitmapset_2))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L23
	} else {
		goto L37
	}
L37:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v29
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v124
	F_errmsg_internal(m, int32(_a_F__readBitmapset_5), v7)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L23
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F__readBitmapset_1), int32(237), int32(_a_F__readBitmapset_2))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L23
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v18
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v139
	F_errmsg_internal(m, int32(_a_F__readBitmapset_4), v7+int32(16))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L23
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F__readBitmapset_1), int32(223), int32(_a_F__readBitmapset_2))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L23
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_r_et_condition_1(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v7 <= v16 {
		v56 = int32(-1)
	} else {
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+v7-int32(1)))))
		if int32(246) < v31 {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7 - int32(1)
			v53 = int32(0)
		} else {
			v33 = v31 - int32(97)
			if v33 < int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7 - int32(1)
				v53 = int32(0)
			} else {
				v36 = int32(1)
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v33)>>(uint(int32(3))%32)))+uint32(_c_F_r_et_condition_1[0]))))
				if int32(base.Ui32(v40)>>(uint(v33&int32(7))%32))&v36 != 0 {
					v53 = v36
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7 - int32(1)
					v53 = int32(0)
				}
			}
		}
		v56 = v53
	}
	if v56 != 0 {
		v148 = v2
		return v148
	} else {
		v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v65 <= v66 {
			v109 = int32(-1)
		} else {
			v78 = int32(1)
			v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79+v65-v78))))
			if int32(246) < v83 {
				v105 = v78
			} else {
				v85 = v83 - int32(97)
				if v85 < int32(0) {
					v105 = v78
				} else {
					v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v85)>>(uint(int32(3))%32)))+uint32(_c_F_r_et_condition_1[0]))))
					if int32(base.Ui32(v91)>>(uint(v85&int32(7))%32))&int32(1) == int32(0) {
						v105 = v78
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v65 - int32(1)
						v105 = int32(0)
					}
				}
			}
			v109 = v105
		}
		if v109 != 0 {
			v148 = v2
			return v148
		} else {
			v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v111 <= v110 {
				v148 = v2
				return v148
			} else {
				v113 = v7 - v6
				v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v115 = v113 + v114
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v115
				if v115 <= v110 {
					v144 = v115
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v144
					v148 = int32(1)
					return v148
				} else {
					v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v120 = int32(1)
					v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118+v115-v120))))
					if base.B2i32(v122&int32(224) != int32(96))|base.B2i32(v120<<(uint(v122)%32)&int32(_a_F_r_et_condition_1_0) == int32(0)) != 0 {
						v144 = v115
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v144
						v148 = int32(1)
						return v148
					} else {
						v134 = int32(0)
						v138 = F_find_among_b(m, l0, int32(_a_F_r_et_condition_1_1), int32(21), v134)
						mBase = m.M
						v141 = m.ExcPending
						if v141 != 0 {
							return int32(0)
						} else {
							if v138 != 0 {
								v148 = v134
							} else {
								v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v144 = v142 + v113
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v144
								v148 = int32(1)
							}
							return v148
						}
					}
				}
			}
		}
	}
}
func F_r_fix_chdz(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v4
	v7 = v4 - int32(1)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v7 <= v8 {
		v46 = v2
		return v46
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v7))))
		if base.B2i32(v12 != int32(190))&base.B2i32(v12 != int32(141)) != 0 {
			v46 = v2
			return v46
		} else {
			v21 = F_find_among_b(m, l0, int32(_a_F_r_fix_chdz_0), int32(2), int32(0))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				if v21 == int32(0) {
					v46 = v2
					return v46
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v27
					switch v21 - int32(1) {
					case 0:
						v33 = F_slice_from_s(m, l0, int32(1), int32(_a_F_r_fix_chdz_1))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							if int32(0) <= v33 {
								v46 = int32(1)
							} else {
								v46 = v33
							}
							return v46
						}
					case 1:
						v39 = F_slice_from_s(m, l0, int32(1), int32(_a_F_r_fix_chdz_2))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							if v39 < int32(0) {
								v46 = v39
							} else {
								v46 = int32(1)
							}
							return v46
						}
					default:
						v46 = int32(1)
						return v46
					}
				}
			}
		}
	}
}
func F_r_fix_ending(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
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
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
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
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v6-int32(4))))
	if v14 == v2 {
		v88 = int32(0)
	} else {
		v19 = v14 & int32(3)
		if base.Ui32(v14) < base.Ui32(int32(4)) {
			v55 = v6
			v56 = int32(0)
			v61 = v55
			v62 = v56
			v66 = v2
			for {
				v67 = int32(*(*int8)(unsafe.Add(mBase, uint32(v61))))
				v70 = v62 + base.B2i32(int32(-65) < v67)
				v71 = int32(1)
				v74 = v66 + v71
				if v74 != v19 {
					v61 = v61 + v71
					v62 = v70
					v66 = v74
					continue
				} else {
					break
				}
				break
			}
			v77 = v70
		} else {
			v26 = v6
			v27 = int32(0)
			v30 = v2
			for {
				v32 = int32(*(*int8)(unsafe.Add(mBase, uint32(v26))))
				v33 = int32(-65)
				v36 = int32(*(*int8)(unsafe.Add(mBase, uint32(v26)+1)))
				v40 = int32(*(*int8)(unsafe.Add(mBase, uint32(v26)+2)))
				v44 = int32(*(*int8)(unsafe.Add(mBase, uint32(v26)+3)))
				v47 = v27 + base.B2i32(v33 < v32) + base.B2i32(v33 < v36) + base.B2i32(v33 < v40) + base.B2i32(v33 < v44)
				v48 = int32(4)
				v49 = v26 + v48
				v51 = v30 + v48
				if v51 != v14&int32(-4) {
					v26 = v49
					v27 = v47
					v30 = v51
					continue
				} else {
					break
				}
				break
			}
			if v19 == int32(0) {
				v77 = v47
			} else {
				v55 = v49
				v56 = v47
				v61 = v55
				v62 = v56
				v66 = v2
				for {
					v67 = int32(*(*int8)(unsafe.Add(mBase, uint32(v61))))
					v70 = v62 + base.B2i32(int32(-65) < v67)
					v71 = int32(1)
					v74 = v66 + v71
					if v74 != v19 {
						v61 = v61 + v71
						v62 = v70
						v66 = v74
						continue
					} else {
						break
					}
					break
				}
				v77 = v70
			}
		}
		v88 = v77
	}
	if v88 < int32(4) {
		v347 = v2
		return v347
	} else {
		v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v91
		v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v93
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v93
		v96 = int32(0)
		v100 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_0), int32(17), v96)
		mBase = m.M
		v103 = m.ExcPending
		if v103 != 0 {
			return int32(0)
		} else {
			if v100 == int32(0) {
				v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v224
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v224
				v227 = int32(3)
				v229 = int32(0)
				v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v224-v232 < v227 {
					v242 = v229
				} else {
					v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v238 = F_memcmp(m, v235+v224-v227, int32(_a_F_r_fix_ending_1), v227)
					mBase = m.M
					if v238 != 0 {
						v242 = v229
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v224 - v227
						v242 = int32(1)
					}
				}
				if v242 == int32(0) {
					v347 = v96
					return v347
				} else {
					v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v250 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_2), int32(6), int32(0))
					mBase = m.M
					v251 = m.ExcPending
					if v251 != 0 {
						return int32(0)
					} else {
						v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						if v250 != 0 {
							v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v254 = v252 - v253
							v255 = int32(3)
							v257 = int32(0)
							v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v253-v260 < v255 {
								v270 = v257
							} else {
								v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v266 = F_memcmp(m, v263+v253-v255, int32(_a_F_r_fix_ending_3), v255)
								mBase = m.M
								if v266 != 0 {
									v270 = v257
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v253 - v255
									v270 = int32(1)
								}
							}
							if v270 == int32(0) {
								v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v283 = v273 - v254
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v283
								v285 = v283
								*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v285
								v287 = F_slice_del(m, l0)
								mBase = m.M
								if int32(0) <= v287 {
									v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
									v347 = int32(1)
								} else {
									v347 = v287
								}
								return v347
							} else {
								v278 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_4), int32(6), int32(0))
								mBase = m.M
								v279 = m.ExcPending
								if v279 != 0 {
									return int32(0)
								} else {
									if v278 != 0 {
										v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										v285 = v280
									} else {
										v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v283 = v281 - v254
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v283
										v285 = v283
									}
									*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v285
									v287 = F_slice_del(m, l0)
									mBase = m.M
									if int32(0) <= v287 {
										v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
										v347 = int32(1)
									} else {
										v347 = v287
									}
									return v347
								}
							}
						} else {
							v290 = v246 - v245
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v252 - v290
							v296 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_5), int32(11), int32(0))
							mBase = m.M
							v297 = m.ExcPending
							if v297 != 0 {
								return int32(0)
							} else {
								if v296 == int32(0) {
									v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v323 - v290
									v326 = int32(0)
									v330 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_6), int32(9), v326)
									mBase = m.M
									v331 = m.ExcPending
									if v331 != 0 {
										return int32(0)
									} else {
										if v330 == int32(0) {
											v347 = v326
										} else {
											v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v335 = v334 - v290
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v335
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v335
											v338 = F_slice_del(m, l0)
											mBase = m.M
											if v338 < int32(0) {
												v347 = v338
											} else {
												v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
												v347 = int32(1)
											}
										}
										return v347
									}
								} else {
									v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v300
									v302 = int32(3)
									v304 = int32(0)
									v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v300-v307 < v302 {
										v317 = v304
									} else {
										v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v313 = F_memcmp(m, v310+v300-v302, int32(_a_F_r_fix_ending_7), v302)
										mBase = m.M
										if v313 != 0 {
											v317 = v304
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v300 - v302
											v317 = int32(1)
										}
									}
									if v317 == int32(0) {
										v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v323 - v290
										v326 = int32(0)
										v330 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_6), int32(9), v326)
										mBase = m.M
										v331 = m.ExcPending
										if v331 != 0 {
											return int32(0)
										} else {
											if v330 == int32(0) {
												v347 = v326
											} else {
												v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												v335 = v334 - v290
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v335
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v335
												v338 = F_slice_del(m, l0)
												mBase = m.M
												if v338 < int32(0) {
													v347 = v338
												} else {
													v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
													v347 = int32(1)
												}
											}
											return v347
										}
									} else {
										v320 = F_slice_del(m, l0)
										mBase = m.M
										if int32(0) <= v320 {
											v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
											v347 = int32(1)
										} else {
											v347 = v320
										}
										return v347
									}
								}
							}
						}
					}
				}
			} else {
				v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v106
				switch v100 - int32(1) {
				case 0:
					v110 = F_slice_del(m, l0)
					mBase = m.M
					if int32(0) <= v110 {
						v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
						v347 = int32(1)
					} else {
						v347 = v110
					}
					return v347
				case 1:
					v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v117 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_8), int32(3), int32(0))
					mBase = m.M
					v118 = m.ExcPending
					if v118 != 0 {
						return int32(0)
					} else {
						if v117 == int32(0) {
							v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v224
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v224
							v227 = int32(3)
							v229 = int32(0)
							v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v224-v232 < v227 {
								v242 = v229
							} else {
								v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v238 = F_memcmp(m, v235+v224-v227, int32(_a_F_r_fix_ending_1), v227)
								mBase = m.M
								if v238 != 0 {
									v242 = v229
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v224 - v227
									v242 = int32(1)
								}
							}
							if v242 == int32(0) {
								v347 = v96
								return v347
							} else {
								v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v250 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_2), int32(6), int32(0))
								mBase = m.M
								v251 = m.ExcPending
								if v251 != 0 {
									return int32(0)
								} else {
									v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									if v250 != 0 {
										v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										v254 = v252 - v253
										v255 = int32(3)
										v257 = int32(0)
										v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v253-v260 < v255 {
											v270 = v257
										} else {
											v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v266 = F_memcmp(m, v263+v253-v255, int32(_a_F_r_fix_ending_3), v255)
											mBase = m.M
											if v266 != 0 {
												v270 = v257
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v253 - v255
												v270 = int32(1)
											}
										}
										if v270 == int32(0) {
											v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v283 = v273 - v254
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v283
											v285 = v283
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v285
											v287 = F_slice_del(m, l0)
											mBase = m.M
											if int32(0) <= v287 {
												v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
												v347 = int32(1)
											} else {
												v347 = v287
											}
											return v347
										} else {
											v278 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_4), int32(6), int32(0))
											mBase = m.M
											v279 = m.ExcPending
											if v279 != 0 {
												return int32(0)
											} else {
												if v278 != 0 {
													v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
													v285 = v280
												} else {
													v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													v283 = v281 - v254
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v283
													v285 = v283
												}
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v285
												v287 = F_slice_del(m, l0)
												mBase = m.M
												if int32(0) <= v287 {
													v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
													v347 = int32(1)
												} else {
													v347 = v287
												}
												return v347
											}
										}
									} else {
										v290 = v246 - v245
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v252 - v290
										v296 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_5), int32(11), int32(0))
										mBase = m.M
										v297 = m.ExcPending
										if v297 != 0 {
											return int32(0)
										} else {
											if v296 == int32(0) {
												v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v323 - v290
												v326 = int32(0)
												v330 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_6), int32(9), v326)
												mBase = m.M
												v331 = m.ExcPending
												if v331 != 0 {
													return int32(0)
												} else {
													if v330 == int32(0) {
														v347 = v326
													} else {
														v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														v335 = v334 - v290
														*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v335
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v335
														v338 = F_slice_del(m, l0)
														mBase = m.M
														if v338 < int32(0) {
															v347 = v338
														} else {
															v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
															*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
															v347 = int32(1)
														}
													}
													return v347
												}
											} else {
												v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v300
												v302 = int32(3)
												v304 = int32(0)
												v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												if v300-v307 < v302 {
													v317 = v304
												} else {
													v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													v313 = F_memcmp(m, v310+v300-v302, int32(_a_F_r_fix_ending_7), v302)
													mBase = m.M
													if v313 != 0 {
														v317 = v304
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v300 - v302
														v317 = int32(1)
													}
												}
												if v317 == int32(0) {
													v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v323 - v290
													v326 = int32(0)
													v330 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_6), int32(9), v326)
													mBase = m.M
													v331 = m.ExcPending
													if v331 != 0 {
														return int32(0)
													} else {
														if v330 == int32(0) {
															v347 = v326
														} else {
															v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															v335 = v334 - v290
															*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v335
															*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v335
															v338 = F_slice_del(m, l0)
															mBase = m.M
															if v338 < int32(0) {
																v347 = v338
															} else {
																v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
																v347 = int32(1)
															}
														}
														return v347
													}
												} else {
													v320 = F_slice_del(m, l0)
													mBase = m.M
													if int32(0) <= v320 {
														v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
														v347 = int32(1)
													} else {
														v347 = v320
													}
													return v347
												}
											}
										}
									}
								}
							}
						} else {
							v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v121 + (v106 - v113)
							v125 = F_slice_del(m, l0)
							mBase = m.M
							if int32(0) <= v125 {
								v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
								v347 = int32(1)
							} else {
								v347 = v125
							}
							return v347
						}
					}
				case 2:
					v130 = F_slice_from_s(m, l0, int32(6), int32(_a_F_r_fix_ending_9))
					mBase = m.M
					v131 = m.ExcPending
					if v131 != 0 {
						return int32(0)
					} else {
						if int32(0) <= v130 {
							v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
							v347 = int32(1)
						} else {
							v347 = v130
						}
						return v347
					}
				case 3:
					v136 = F_slice_from_s(m, l0, int32(6), int32(_a_F_r_fix_ending_10))
					mBase = m.M
					v137 = m.ExcPending
					if v137 != 0 {
						return int32(0)
					} else {
						if int32(0) <= v136 {
							v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
							v347 = int32(1)
						} else {
							v347 = v136
						}
						return v347
					}
				case 4:
					v142 = F_slice_from_s(m, l0, int32(6), int32(_a_F_r_fix_ending_11))
					mBase = m.M
					v143 = m.ExcPending
					if v143 != 0 {
						return int32(0)
					} else {
						if int32(0) <= v142 {
							v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
							v347 = int32(1)
						} else {
							v347 = v142
						}
						return v347
					}
				case 5:
					v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
					if v146 == int32(0) {
						v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v224
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v224
						v227 = int32(3)
						v229 = int32(0)
						v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v224-v232 < v227 {
							v242 = v229
						} else {
							v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v238 = F_memcmp(m, v235+v224-v227, int32(_a_F_r_fix_ending_1), v227)
							mBase = m.M
							if v238 != 0 {
								v242 = v229
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v224 - v227
								v242 = int32(1)
							}
						}
						if v242 == int32(0) {
							v347 = v96
							return v347
						} else {
							v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v250 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_2), int32(6), int32(0))
							mBase = m.M
							v251 = m.ExcPending
							if v251 != 0 {
								return int32(0)
							} else {
								v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								if v250 != 0 {
									v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v254 = v252 - v253
									v255 = int32(3)
									v257 = int32(0)
									v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v253-v260 < v255 {
										v270 = v257
									} else {
										v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v266 = F_memcmp(m, v263+v253-v255, int32(_a_F_r_fix_ending_3), v255)
										mBase = m.M
										if v266 != 0 {
											v270 = v257
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v253 - v255
											v270 = int32(1)
										}
									}
									if v270 == int32(0) {
										v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v283 = v273 - v254
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v283
										v285 = v283
										*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v285
										v287 = F_slice_del(m, l0)
										mBase = m.M
										if int32(0) <= v287 {
											v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
											v347 = int32(1)
										} else {
											v347 = v287
										}
										return v347
									} else {
										v278 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_4), int32(6), int32(0))
										mBase = m.M
										v279 = m.ExcPending
										if v279 != 0 {
											return int32(0)
										} else {
											if v278 != 0 {
												v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												v285 = v280
											} else {
												v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												v283 = v281 - v254
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v283
												v285 = v283
											}
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v285
											v287 = F_slice_del(m, l0)
											mBase = m.M
											if int32(0) <= v287 {
												v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
												v347 = int32(1)
											} else {
												v347 = v287
											}
											return v347
										}
									}
								} else {
									v290 = v246 - v245
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v252 - v290
									v296 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_5), int32(11), int32(0))
									mBase = m.M
									v297 = m.ExcPending
									if v297 != 0 {
										return int32(0)
									} else {
										if v296 == int32(0) {
											v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v323 - v290
											v326 = int32(0)
											v330 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_6), int32(9), v326)
											mBase = m.M
											v331 = m.ExcPending
											if v331 != 0 {
												return int32(0)
											} else {
												if v330 == int32(0) {
													v347 = v326
												} else {
													v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													v335 = v334 - v290
													*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v335
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v335
													v338 = F_slice_del(m, l0)
													mBase = m.M
													if v338 < int32(0) {
														v347 = v338
													} else {
														v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
														v347 = int32(1)
													}
												}
												return v347
											}
										} else {
											v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v300
											v302 = int32(3)
											v304 = int32(0)
											v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v300-v307 < v302 {
												v317 = v304
											} else {
												v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v313 = F_memcmp(m, v310+v300-v302, int32(_a_F_r_fix_ending_7), v302)
												mBase = m.M
												if v313 != 0 {
													v317 = v304
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v300 - v302
													v317 = int32(1)
												}
											}
											if v317 == int32(0) {
												v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v323 - v290
												v326 = int32(0)
												v330 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_6), int32(9), v326)
												mBase = m.M
												v331 = m.ExcPending
												if v331 != 0 {
													return int32(0)
												} else {
													if v330 == int32(0) {
														v347 = v326
													} else {
														v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														v335 = v334 - v290
														*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v335
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v335
														v338 = F_slice_del(m, l0)
														mBase = m.M
														if v338 < int32(0) {
															v347 = v338
														} else {
															v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
															*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
															v347 = int32(1)
														}
													}
													return v347
												}
											} else {
												v320 = F_slice_del(m, l0)
												mBase = m.M
												if int32(0) <= v320 {
													v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
													v347 = int32(1)
												} else {
													v347 = v320
												}
												return v347
											}
										}
									}
								}
							}
						}
					} else {
						v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v150 = int32(3)
						v152 = int32(0)
						v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v154-v155 < v150 {
							v165 = v152
						} else {
							v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v161 = F_memcmp(m, v158+v154-v150, int32(_a_F_r_fix_ending_12), v150)
							mBase = m.M
							if v161 != 0 {
								v165 = v152
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v154 - v150
								v165 = int32(1)
							}
						}
						if v165 != 0 {
							v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v224
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v224
							v227 = int32(3)
							v229 = int32(0)
							v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v224-v232 < v227 {
								v242 = v229
							} else {
								v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v238 = F_memcmp(m, v235+v224-v227, int32(_a_F_r_fix_ending_1), v227)
								mBase = m.M
								if v238 != 0 {
									v242 = v229
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v224 - v227
									v242 = int32(1)
								}
							}
							if v242 == int32(0) {
								v347 = v96
								return v347
							} else {
								v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v250 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_2), int32(6), int32(0))
								mBase = m.M
								v251 = m.ExcPending
								if v251 != 0 {
									return int32(0)
								} else {
									v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									if v250 != 0 {
										v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										v254 = v252 - v253
										v255 = int32(3)
										v257 = int32(0)
										v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v253-v260 < v255 {
											v270 = v257
										} else {
											v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v266 = F_memcmp(m, v263+v253-v255, int32(_a_F_r_fix_ending_3), v255)
											mBase = m.M
											if v266 != 0 {
												v270 = v257
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v253 - v255
												v270 = int32(1)
											}
										}
										if v270 == int32(0) {
											v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v283 = v273 - v254
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v283
											v285 = v283
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v285
											v287 = F_slice_del(m, l0)
											mBase = m.M
											if int32(0) <= v287 {
												v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
												v347 = int32(1)
											} else {
												v347 = v287
											}
											return v347
										} else {
											v278 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_4), int32(6), int32(0))
											mBase = m.M
											v279 = m.ExcPending
											if v279 != 0 {
												return int32(0)
											} else {
												if v278 != 0 {
													v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
													v285 = v280
												} else {
													v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													v283 = v281 - v254
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v283
													v285 = v283
												}
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v285
												v287 = F_slice_del(m, l0)
												mBase = m.M
												if int32(0) <= v287 {
													v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
													v347 = int32(1)
												} else {
													v347 = v287
												}
												return v347
											}
										}
									} else {
										v290 = v246 - v245
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v252 - v290
										v296 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_5), int32(11), int32(0))
										mBase = m.M
										v297 = m.ExcPending
										if v297 != 0 {
											return int32(0)
										} else {
											if v296 == int32(0) {
												v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v323 - v290
												v326 = int32(0)
												v330 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_6), int32(9), v326)
												mBase = m.M
												v331 = m.ExcPending
												if v331 != 0 {
													return int32(0)
												} else {
													if v330 == int32(0) {
														v347 = v326
													} else {
														v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														v335 = v334 - v290
														*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v335
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v335
														v338 = F_slice_del(m, l0)
														mBase = m.M
														if v338 < int32(0) {
															v347 = v338
														} else {
															v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
															*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
															v347 = int32(1)
														}
													}
													return v347
												}
											} else {
												v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v300
												v302 = int32(3)
												v304 = int32(0)
												v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												if v300-v307 < v302 {
													v317 = v304
												} else {
													v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													v313 = F_memcmp(m, v310+v300-v302, int32(_a_F_r_fix_ending_7), v302)
													mBase = m.M
													if v313 != 0 {
														v317 = v304
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v300 - v302
														v317 = int32(1)
													}
												}
												if v317 == int32(0) {
													v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v323 - v290
													v326 = int32(0)
													v330 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_6), int32(9), v326)
													mBase = m.M
													v331 = m.ExcPending
													if v331 != 0 {
														return int32(0)
													} else {
														if v330 == int32(0) {
															v347 = v326
														} else {
															v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															v335 = v334 - v290
															*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v335
															*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v335
															v338 = F_slice_del(m, l0)
															mBase = m.M
															if v338 < int32(0) {
																v347 = v338
															} else {
																v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
																v347 = int32(1)
															}
														}
														return v347
													}
												} else {
													v320 = F_slice_del(m, l0)
													mBase = m.M
													if int32(0) <= v320 {
														v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
														v347 = int32(1)
													} else {
														v347 = v320
													}
													return v347
												}
											}
										}
									}
								}
							}
						} else {
							v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v166 + (v106 - v149)
							v172 = F_slice_from_s(m, l0, int32(6), int32(_a_F_r_fix_ending_13))
							mBase = m.M
							v173 = m.ExcPending
							if v173 != 0 {
								return int32(0)
							} else {
								if int32(0) <= v172 {
									v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
									v347 = int32(1)
								} else {
									v347 = v172
								}
								return v347
							}
						}
					}
				case 6:
					v178 = F_slice_from_s(m, l0, int32(3), int32(_a_F_r_fix_ending_14))
					mBase = m.M
					v179 = m.ExcPending
					if v179 != 0 {
						return int32(0)
					} else {
						if int32(0) <= v178 {
							v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
							v347 = int32(1)
						} else {
							v347 = v178
						}
						return v347
					}
				case 7:
					v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v186 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_15), int32(8), int32(0))
					mBase = m.M
					v187 = m.ExcPending
					if v187 != 0 {
						return int32(0)
					} else {
						if v186 != 0 {
							v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v224
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v224
							v227 = int32(3)
							v229 = int32(0)
							v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v224-v232 < v227 {
								v242 = v229
							} else {
								v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v238 = F_memcmp(m, v235+v224-v227, int32(_a_F_r_fix_ending_1), v227)
								mBase = m.M
								if v238 != 0 {
									v242 = v229
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v224 - v227
									v242 = int32(1)
								}
							}
							if v242 == int32(0) {
								v347 = v96
								return v347
							} else {
								v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v250 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_2), int32(6), int32(0))
								mBase = m.M
								v251 = m.ExcPending
								if v251 != 0 {
									return int32(0)
								} else {
									v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									if v250 != 0 {
										v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										v254 = v252 - v253
										v255 = int32(3)
										v257 = int32(0)
										v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v253-v260 < v255 {
											v270 = v257
										} else {
											v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v266 = F_memcmp(m, v263+v253-v255, int32(_a_F_r_fix_ending_3), v255)
											mBase = m.M
											if v266 != 0 {
												v270 = v257
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v253 - v255
												v270 = int32(1)
											}
										}
										if v270 == int32(0) {
											v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v283 = v273 - v254
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v283
											v285 = v283
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v285
											v287 = F_slice_del(m, l0)
											mBase = m.M
											if int32(0) <= v287 {
												v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
												v347 = int32(1)
											} else {
												v347 = v287
											}
											return v347
										} else {
											v278 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_4), int32(6), int32(0))
											mBase = m.M
											v279 = m.ExcPending
											if v279 != 0 {
												return int32(0)
											} else {
												if v278 != 0 {
													v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
													v285 = v280
												} else {
													v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													v283 = v281 - v254
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v283
													v285 = v283
												}
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v285
												v287 = F_slice_del(m, l0)
												mBase = m.M
												if int32(0) <= v287 {
													v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
													v347 = int32(1)
												} else {
													v347 = v287
												}
												return v347
											}
										}
									} else {
										v290 = v246 - v245
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v252 - v290
										v296 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_5), int32(11), int32(0))
										mBase = m.M
										v297 = m.ExcPending
										if v297 != 0 {
											return int32(0)
										} else {
											if v296 == int32(0) {
												v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v323 - v290
												v326 = int32(0)
												v330 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_6), int32(9), v326)
												mBase = m.M
												v331 = m.ExcPending
												if v331 != 0 {
													return int32(0)
												} else {
													if v330 == int32(0) {
														v347 = v326
													} else {
														v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														v335 = v334 - v290
														*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v335
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v335
														v338 = F_slice_del(m, l0)
														mBase = m.M
														if v338 < int32(0) {
															v347 = v338
														} else {
															v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
															*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
															v347 = int32(1)
														}
													}
													return v347
												}
											} else {
												v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v300
												v302 = int32(3)
												v304 = int32(0)
												v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												if v300-v307 < v302 {
													v317 = v304
												} else {
													v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													v313 = F_memcmp(m, v310+v300-v302, int32(_a_F_r_fix_ending_7), v302)
													mBase = m.M
													if v313 != 0 {
														v317 = v304
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v300 - v302
														v317 = int32(1)
													}
												}
												if v317 == int32(0) {
													v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v323 - v290
													v326 = int32(0)
													v330 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_6), int32(9), v326)
													mBase = m.M
													v331 = m.ExcPending
													if v331 != 0 {
														return int32(0)
													} else {
														if v330 == int32(0) {
															v347 = v326
														} else {
															v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															v335 = v334 - v290
															*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v335
															*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v335
															v338 = F_slice_del(m, l0)
															mBase = m.M
															if v338 < int32(0) {
																v347 = v338
															} else {
																v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
																v347 = int32(1)
															}
														}
														return v347
													}
												} else {
													v320 = F_slice_del(m, l0)
													mBase = m.M
													if int32(0) <= v320 {
														v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
														v347 = int32(1)
													} else {
														v347 = v320
													}
													return v347
												}
											}
										}
									}
								}
							}
						} else {
							v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v188 + (v106 - v182)
							v192 = F_slice_del(m, l0)
							mBase = m.M
							if int32(0) <= v192 {
								v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
								v347 = int32(1)
							} else {
								v347 = v192
							}
							return v347
						}
					}
				case 8:
					v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v106-int32(2) <= v195 {
						v218 = F_slice_from_s(m, l0, int32(6), int32(_a_F_r_fix_ending_16))
						mBase = m.M
						v219 = m.ExcPending
						if v219 != 0 {
							return int32(0)
						} else {
							if int32(0) <= v218 {
								v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
								v347 = int32(1)
							} else {
								v347 = v218
							}
							return v347
						}
					} else {
						v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199+v106-int32(1)))))
						switch v203 - int32(136) {
						case 0, 5:
							v209 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_17), int32(3), int32(0))
							mBase = m.M
							v210 = m.ExcPending
							if v210 != 0 {
								return int32(0)
							} else {
								switch v209 - int32(1) {
								case 0:
									v213 = F_slice_del(m, l0)
									mBase = m.M
									if int32(0) <= v213 {
										v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
										v347 = int32(1)
									} else {
										v347 = v213
									}
									return v347
								case 1:
									v218 = F_slice_from_s(m, l0, int32(6), int32(_a_F_r_fix_ending_16))
									mBase = m.M
									v219 = m.ExcPending
									if v219 != 0 {
										return int32(0)
									} else {
										if int32(0) <= v218 {
											v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
											v347 = int32(1)
										} else {
											v347 = v218
										}
										return v347
									}
								default:
									v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
									v347 = int32(1)
									return v347
								}
							}
						default:
							v218 = F_slice_from_s(m, l0, int32(6), int32(_a_F_r_fix_ending_16))
							mBase = m.M
							v219 = m.ExcPending
							if v219 != 0 {
								return int32(0)
							} else {
								if int32(0) <= v218 {
									v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
									v347 = int32(1)
								} else {
									v347 = v218
								}
								return v347
							}
						}
					}
				default:
					v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
					v347 = int32(1)
					return v347
				}
			}
		}
	}
}
func F_r_lengthen_V_2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v277 int32
	_ = v277
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
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
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v411 int32
	_ = v411
	var v428 int32
	_ = v428
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v550 int32
	_ = v550
	var v566 int32
	_ = v566
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v673 int32
	_ = v673
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v685 int32
	_ = v685
	var v701 int32
	_ = v701
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v739 int32
	_ = v739
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v751 int32
	_ = v751
	var v755 int32
	_ = v755
	var v769 int32
	_ = v769
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v866 int32
	_ = v866
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v878 int32
	_ = v878
	var v894 int32
	_ = v894
	var v901 int32
	_ = v901
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v937 int32
	_ = v937
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v963 int32
	_ = v963
	var v965 int32
	_ = v965
	var v967 int32
	_ = v967
	var v985 int32
	_ = v985
	var v987 int32
	_ = v987
	var v995 int32
	_ = v995
	var v999 int32
	_ = v999
	var v1001 int32
	_ = v1001
	var v1007 int32
	_ = v1007
	var v1024 int32
	_ = v1024
	var v1031 int32
	_ = v1031
	var v1034 int32
	_ = v1034
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1065 int32
	_ = v1065
	var v1070 int32
	_ = v1070
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L5
L1:
	;
	return v1070
L2:
	;
	v1065 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1065 + (v8 - v7)
	v1070 = int32(1)
	goto L1
L3:
	;
	if v138 != 0 {
		goto L2
	} else {
		goto L21
	}
L4:
	;
	v138 = v131
	goto L3
L5:
	;
	if v8 <= v22 {
		v131 = int32(-1)
		goto L4
	} else {
		goto L7
	}
L6:
	;
	v131 = int32(0)
	goto L4
L7:
	;
	v39 = int32(1)
	v40 = v8 - v39
	v42 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23+v40))))
	v44 = v42 & int32(255)
	if base.B2i32(v40 == v22)|base.B2i32(int32(0) <= v42) != 0 {
		v102 = v44
		v106 = v39
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if int32(252) < v102 {
		goto L16
	} else {
		goto L17
	}
L9:
	;
	v51 = v44 & int32(63)
	v53 = v8 - int32(2)
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+v53))))
	v57 = v55 << (uint(int32(6)) % 32)
	if base.B2i32(v53 != v22)&base.B2i32(base.Ui32(v55) < base.Ui32(int32(192))) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v102 = v57&int32(1984) | v51
	v106 = int32(2)
	goto L8
L11:
	;
	goto L12
L12:
	;
	v70 = v57&int32(4032) | v51
	v72 = v8 - int32(3)
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+v72))))
	if base.B2i32(v72 != v22)&base.B2i32(base.Ui32(v74) < base.Ui32(int32(224))) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v102 = v74<<(uint(int32(12))%32)&int32(_a_F_r_lengthen_V_2_0) | v70
	v106 = int32(3)
	goto L8
L14:
	;
	goto L15
L15:
	;
	v92 = int32(4)
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+v23-v92))))
	v102 = v74<<(uint(int32(12))%32)&int32(_a_F_r_lengthen_V_2_1) | v94&int32(7)<<(uint(int32(18))%32) | v70
	v106 = v92
	goto L8
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v8 - v106
	goto L20
L17:
	;
	v108 = v102 - int32(97)
	if v108 < int32(0) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v108)>>(uint(int32(3))%32)))+uint32(_c_F_r_lengthen_V_2[0]))))
	if int32(base.Ui32(v114)>>(uint(v108&int32(7))%32))&int32(1) == int32(0) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v138 = v106
	goto L3
L20:
	;
	goto L6
L21:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v139
	v144 = F_find_among_b(m, l0, int32(_a_F_r_lengthen_V_2_2), int32(21), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	return int32(0)
L23:
	;
	if v144 == int32(0) {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v150
	switch v144 - int32(1) {
	case 0:
		goto L28
	case 1:
		goto L27
	case 2:
		goto L26
	case 3:
		goto L25
	default:
		goto L2
	}
L25:
	;
	v1058 = F_slice_from_s(m, l0, int32(3), int32(_a_F_r_lengthen_V_2_3))
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L22
	} else {
		goto L200
	}
L26:
	;
	v1052 = F_slice_from_s(m, l0, int32(4), int32(_a_F_r_lengthen_V_2_4))
	mBase = m.M
	v1053 = m.ExcPending
	if v1053 != 0 {
		goto L22
	} else {
		goto L198
	}
L27:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v305 = v304 - v150
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L57
L28:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L31
L29:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v287 = v285 + (v150 - v154)
	if v284 != 0 {
		goto L47
	} else {
		goto L48
	}
L30:
	;
	v284 = v277
	goto L29
L31:
	;
	if v167 <= v168 {
		v277 = int32(-1)
		goto L30
	} else {
		goto L33
	}
L32:
	;
	v277 = int32(0)
	goto L30
L33:
	;
	v185 = int32(1)
	v186 = v167 - v185
	v188 = int32(*(*int8)(unsafe.Add(mBase, uint32(v169+v186))))
	v190 = v188 & int32(255)
	if base.B2i32(v186 == v168)|base.B2i32(int32(0) <= v188) != 0 {
		v248 = v190
		v252 = v185
		goto L34
	} else {
		goto L35
	}
L34:
	;
	if int32(252) < v248 {
		goto L42
	} else {
		goto L43
	}
L35:
	;
	v197 = v190 & int32(63)
	v199 = v167 - int32(2)
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169+v199))))
	v203 = v201 << (uint(int32(6)) % 32)
	if base.B2i32(v199 != v168)&base.B2i32(base.Ui32(v201) < base.Ui32(int32(192))) == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v248 = v203&int32(1984) | v197
	v252 = int32(2)
	goto L34
L37:
	;
	goto L38
L38:
	;
	v216 = v203&int32(4032) | v197
	v218 = v167 - int32(3)
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169+v218))))
	if base.B2i32(v218 != v168)&base.B2i32(base.Ui32(v220) < base.Ui32(int32(224))) == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v248 = v220<<(uint(int32(12))%32)&int32(_a_F_r_lengthen_V_2_0) | v216
	v252 = int32(3)
	goto L34
L40:
	;
	goto L41
L41:
	;
	v238 = int32(4)
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167+v169-v238))))
	v248 = v220<<(uint(int32(12))%32)&int32(_a_F_r_lengthen_V_2_1) | v240&int32(7)<<(uint(int32(18))%32) | v216
	v252 = v238
	goto L34
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v167 - v252
	goto L46
L43:
	;
	v254 = v248 - int32(97)
	if v254 < int32(0) {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v254)>>(uint(int32(3))%32)))+uint32(_c_F_r_lengthen_V_2[1]))))
	if int32(base.Ui32(v260)>>(uint(v254&int32(7))%32))&int32(1) == int32(0) {
		goto L42
	} else {
		goto L45
	}
L45:
	;
	v284 = v252
	goto L29
L46:
	;
	goto L32
L47:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v288 < v287 {
		goto L2
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v287
	v293 = F_slice_to(m, l0, l0+int32(40))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L22
	} else {
		goto L51
	}
L50:
	;
	goto L49
L51:
	;
	if v293 < int32(0) {
		v1070 = v293
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v299 = F_insert_v(m, l0, v297, v297, v298)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L22
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v297
	if int32(0) <= v299 {
		goto L2
	} else {
		goto L54
	}
L54:
	;
	v1070 = v299
	goto L1
L55:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v435 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L56:
	;
	v435 = v428
	goto L55
L57:
	;
	if v318 <= v319 {
		v428 = int32(-1)
		goto L56
	} else {
		goto L59
	}
L58:
	;
	v428 = int32(0)
	goto L56
L59:
	;
	v336 = int32(1)
	v337 = v318 - v336
	v339 = int32(*(*int8)(unsafe.Add(mBase, uint32(v320+v337))))
	v341 = v339 & int32(255)
	if base.B2i32(v337 == v319)|base.B2i32(int32(0) <= v339) != 0 {
		v399 = v341
		v403 = v336
		goto L60
	} else {
		goto L61
	}
L60:
	;
	if int32(252) < v399 {
		goto L68
	} else {
		goto L69
	}
L61:
	;
	v348 = v341 & int32(63)
	v350 = v318 - int32(2)
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320+v350))))
	v354 = v352 << (uint(int32(6)) % 32)
	if base.B2i32(v350 != v319)&base.B2i32(base.Ui32(v352) < base.Ui32(int32(192))) == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v399 = v354&int32(1984) | v348
	v403 = int32(2)
	goto L60
L63:
	;
	goto L64
L64:
	;
	v367 = v354&int32(4032) | v348
	v369 = v318 - int32(3)
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320+v369))))
	if base.B2i32(v369 != v319)&base.B2i32(base.Ui32(v371) < base.Ui32(int32(224))) == int32(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v399 = v371<<(uint(int32(12))%32)&int32(_a_F_r_lengthen_V_2_0) | v367
	v403 = int32(3)
	goto L60
L66:
	;
	goto L67
L67:
	;
	v389 = int32(4)
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318+v320-v389))))
	v399 = v371<<(uint(int32(12))%32)&int32(_a_F_r_lengthen_V_2_1) | v391&int32(7)<<(uint(int32(18))%32) | v367
	v403 = v389
	goto L60
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v318 - v403
	goto L72
L69:
	;
	v405 = v399 - int32(97)
	if v405 < int32(0) {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v405)>>(uint(int32(3))%32)))+uint32(_c_F_r_lengthen_V_2[1]))))
	if int32(base.Ui32(v411)>>(uint(v405&int32(7))%32))&int32(1) == int32(0) {
		goto L68
	} else {
		goto L71
	}
L71:
	;
	v435 = v403
	goto L55
L72:
	;
	goto L58
L73:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v459 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L80
L74:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v444 = v439
	goto L73
L75:
	;
	goto L76
L76:
	;
	v440 = v436 - v305
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v440
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v442 < v440 {
		goto L2
	} else {
		goto L77
	}
L77:
	;
	v444 = v440
	goto L73
L78:
	;
	if v573 == int32(0) {
		goto L2
	} else {
		goto L101
	}
L79:
	;
	v573 = v566
	goto L78
L80:
	;
	if v457 <= v458 {
		v566 = int32(-1)
		goto L79
	} else {
		goto L82
	}
L81:
	;
	v566 = int32(0)
	goto L79
L82:
	;
	v475 = int32(1)
	v476 = v457 - v475
	v478 = int32(*(*int8)(unsafe.Add(mBase, uint32(v459+v476))))
	v480 = v478 & int32(255)
	if base.B2i32(v476 == v458)|base.B2i32(int32(0) <= v478) != 0 {
		v538 = v480
		v542 = v475
		goto L83
	} else {
		goto L84
	}
L83:
	;
	if int32(252) < v538 {
		goto L91
	} else {
		goto L92
	}
L84:
	;
	v487 = v480 & int32(63)
	v489 = v457 - int32(2)
	v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v459+v489))))
	v493 = v491 << (uint(int32(6)) % 32)
	if base.B2i32(v489 != v458)&base.B2i32(base.Ui32(v491) < base.Ui32(int32(192))) == int32(0) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v538 = v493&int32(1984) | v487
	v542 = int32(2)
	goto L83
L86:
	;
	goto L87
L87:
	;
	v506 = v493&int32(4032) | v487
	v508 = v457 - int32(3)
	v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v459+v508))))
	if base.B2i32(v508 != v458)&base.B2i32(base.Ui32(v510) < base.Ui32(int32(224))) == int32(0) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v538 = v510<<(uint(int32(12))%32)&int32(_a_F_r_lengthen_V_2_0) | v506
	v542 = int32(3)
	goto L83
L89:
	;
	goto L90
L90:
	;
	v528 = int32(4)
	v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v457+v459-v528))))
	v538 = v510<<(uint(int32(12))%32)&int32(_a_F_r_lengthen_V_2_1) | v530&int32(7)<<(uint(int32(18))%32) | v506
	v542 = v528
	goto L83
L91:
	;
	v573 = v542
	goto L78
L92:
	;
	goto L93
L93:
	;
	v544 = v538 - int32(97)
	if v544 < int32(0) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v573 = v542
	goto L78
L95:
	;
	goto L96
L96:
	;
	v550 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v544)>>(uint(int32(3))%32)))+uint32(_c_F_r_lengthen_V_2[2]))))
	if int32(base.Ui32(v550)>>(uint(v544&int32(7))%32))&int32(1) == int32(0) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v573 = v542
	goto L78
L98:
	;
	goto L99
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v457 - v542
	goto L100
L100:
	;
	goto L81
L101:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v577 = v436 - v444
	v578 = v576 - v577
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v578
	v593 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v594 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L105
L102:
	;
	v714 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v715 = v714 - v577
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v715
	v717 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L133
L103:
	;
	if v708 != 0 {
		goto L126
	} else {
		goto L127
	}
L104:
	;
	v708 = v701
	goto L103
L105:
	;
	if v578 <= v593 {
		v701 = int32(-1)
		goto L104
	} else {
		goto L107
	}
L106:
	;
	v701 = int32(0)
	goto L104
L107:
	;
	v610 = int32(1)
	v611 = v578 - v610
	v613 = int32(*(*int8)(unsafe.Add(mBase, uint32(v594+v611))))
	v615 = v613 & int32(255)
	if base.B2i32(v611 == v593)|base.B2i32(int32(0) <= v613) != 0 {
		v673 = v615
		v677 = v610
		goto L108
	} else {
		goto L109
	}
L108:
	;
	if int32(235) < v673 {
		goto L116
	} else {
		goto L117
	}
L109:
	;
	v622 = v615 & int32(63)
	v624 = v578 - int32(2)
	v626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v594+v624))))
	v628 = v626 << (uint(int32(6)) % 32)
	if base.B2i32(v624 != v593)&base.B2i32(base.Ui32(v626) < base.Ui32(int32(192))) == int32(0) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v673 = v628&int32(1984) | v622
	v677 = int32(2)
	goto L108
L111:
	;
	goto L112
L112:
	;
	v641 = v628&int32(4032) | v622
	v643 = v578 - int32(3)
	v645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v594+v643))))
	if base.B2i32(v643 != v593)&base.B2i32(base.Ui32(v645) < base.Ui32(int32(224))) == int32(0) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v673 = v645<<(uint(int32(12))%32)&int32(_a_F_r_lengthen_V_2_0) | v641
	v677 = int32(3)
	goto L108
L114:
	;
	goto L115
L115:
	;
	v663 = int32(4)
	v665 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v578+v594-v663))))
	v673 = v645<<(uint(int32(12))%32)&int32(_a_F_r_lengthen_V_2_1) | v665&int32(7)<<(uint(int32(18))%32) | v641
	v677 = v663
	goto L108
L116:
	;
	v708 = v677
	goto L103
L117:
	;
	goto L118
L118:
	;
	v679 = v673 - int32(101)
	if v679 < int32(0) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v708 = v677
	goto L103
L120:
	;
	goto L121
L121:
	;
	v685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v679)>>(uint(int32(3))%32)))+uint32(_c_F_r_lengthen_V_2[3]))))
	if int32(base.Ui32(v685)>>(uint(v679&int32(7))%32))&int32(1) == int32(0) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v708 = v677
	goto L103
L123:
	;
	goto L124
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v578 - v677
	goto L125
L125:
	;
	goto L106
L126:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v713 = v709
	goto L102
L127:
	;
	goto L128
L128:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v711 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v711 <= v710 {
		goto L2
	} else {
		goto L129
	}
L129:
	;
	v713 = v710
	goto L102
L130:
	;
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1034 - v305
	v1039 = F_slice_to(m, l0, l0+int32(40))
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L22
	} else {
		goto L194
	}
L131:
	;
	if v769 < int32(0) {
		goto L130
	} else {
		goto L150
	}
L133:
	;
	goto L134
L134:
	;
	goto L135
L135:
	;
	v724 = v715
	v726 = int32(1)
	goto L138
L137:
	;
	v769 = v751
	goto L131
L138:
	;
	if v724 <= v713 {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	goto L137
L140:
	;
	v769 = int32(-1)
	goto L131
L141:
	;
	goto L142
L142:
	;
	v731 = v724 - int32(1)
	v733 = int32(*(*int8)(unsafe.Add(mBase, uint32(v717+v731))))
	if base.B2i32(int32(0) <= v733)|base.B2i32(v731 <= v713) != 0 {
		v751 = v731
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v755 = int32(1)
	if v755 < v726 {
		v724 = v751
		v726 = v726 - v755
		goto L138
	} else {
		goto L149
	}
L144:
	;
	v739 = v731
	goto L145
L145:
	;
	v744 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v717+v739))))
	if base.Ui32(int32(191)) < base.Ui32(v744) {
		v751 = v739
		goto L143
	} else {
		goto L147
	}
L146:
	;
	v751 = v713
	goto L143
L147:
	;
	v748 = v739 - int32(1)
	if v713 < v748 {
		v739 = v748
		goto L145
	} else {
		goto L148
	}
L148:
	;
	goto L146
L149:
	;
	goto L139
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v769
	v786 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v787 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L153
L151:
	;
	if v901 != 0 {
		goto L130
	} else {
		goto L174
	}
L152:
	;
	v901 = v894
	goto L151
L153:
	;
	if v769 <= v786 {
		v894 = int32(-1)
		goto L152
	} else {
		goto L155
	}
L154:
	;
	v894 = int32(0)
	goto L152
L155:
	;
	v803 = int32(1)
	v804 = v769 - v803
	v806 = int32(*(*int8)(unsafe.Add(mBase, uint32(v787+v804))))
	v808 = v806 & int32(255)
	if base.B2i32(v804 == v786)|base.B2i32(int32(0) <= v806) != 0 {
		v866 = v808
		v870 = v803
		goto L156
	} else {
		goto L157
	}
L156:
	;
	if int32(252) < v866 {
		goto L164
	} else {
		goto L165
	}
L157:
	;
	v815 = v808 & int32(63)
	v817 = v769 - int32(2)
	v819 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v787+v817))))
	v821 = v819 << (uint(int32(6)) % 32)
	if base.B2i32(v817 != v786)&base.B2i32(base.Ui32(v819) < base.Ui32(int32(192))) == int32(0) {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v866 = v821&int32(1984) | v815
	v870 = int32(2)
	goto L156
L159:
	;
	goto L160
L160:
	;
	v834 = v821&int32(4032) | v815
	v836 = v769 - int32(3)
	v838 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v787+v836))))
	if base.B2i32(v836 != v786)&base.B2i32(base.Ui32(v838) < base.Ui32(int32(224))) == int32(0) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v866 = v838<<(uint(int32(12))%32)&int32(_a_F_r_lengthen_V_2_0) | v834
	v870 = int32(3)
	goto L156
L162:
	;
	goto L163
L163:
	;
	v856 = int32(4)
	v858 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v769+v787-v856))))
	v866 = v838<<(uint(int32(12))%32)&int32(_a_F_r_lengthen_V_2_1) | v858&int32(7)<<(uint(int32(18))%32) | v834
	v870 = v856
	goto L156
L164:
	;
	v901 = v870
	goto L151
L165:
	;
	goto L166
L166:
	;
	v872 = v866 - int32(97)
	if v872 < int32(0) {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v901 = v870
	goto L151
L168:
	;
	goto L169
L169:
	;
	v878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v872)>>(uint(int32(3))%32)))+uint32(_c_F_r_lengthen_V_2[2]))))
	if int32(base.Ui32(v878)>>(uint(v872&int32(7))%32))&int32(1) == int32(0) {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v901 = v870
	goto L151
L171:
	;
	goto L172
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v769 - v870
	goto L173
L173:
	;
	goto L154
L174:
	;
	v914 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v915 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v916 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L177
L175:
	;
	if v1031 == int32(0) {
		goto L2
	} else {
		goto L193
	}
L176:
	;
	v1031 = v1024
	goto L175
L177:
	;
	if v914 <= v915 {
		v1024 = int32(-1)
		goto L176
	} else {
		goto L179
	}
L178:
	;
	v1024 = int32(0)
	goto L176
L179:
	;
	v932 = int32(1)
	v933 = v914 - v932
	v935 = int32(*(*int8)(unsafe.Add(mBase, uint32(v916+v933))))
	v937 = v935 & int32(255)
	if base.B2i32(v933 == v915)|base.B2i32(int32(0) <= v935) != 0 {
		v995 = v937
		v999 = v932
		goto L180
	} else {
		goto L181
	}
L180:
	;
	if int32(252) < v995 {
		goto L188
	} else {
		goto L189
	}
L181:
	;
	v944 = v937 & int32(63)
	v946 = v914 - int32(2)
	v948 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v916+v946))))
	v950 = v948 << (uint(int32(6)) % 32)
	if base.B2i32(v946 != v915)&base.B2i32(base.Ui32(v948) < base.Ui32(int32(192))) == int32(0) {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v995 = v950&int32(1984) | v944
	v999 = int32(2)
	goto L180
L183:
	;
	goto L184
L184:
	;
	v963 = v950&int32(4032) | v944
	v965 = v914 - int32(3)
	v967 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v916+v965))))
	if base.B2i32(v965 != v915)&base.B2i32(base.Ui32(v967) < base.Ui32(int32(224))) == int32(0) {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v995 = v967<<(uint(int32(12))%32)&int32(_a_F_r_lengthen_V_2_0) | v963
	v999 = int32(3)
	goto L180
L186:
	;
	goto L187
L187:
	;
	v985 = int32(4)
	v987 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v914+v916-v985))))
	v995 = v967<<(uint(int32(12))%32)&int32(_a_F_r_lengthen_V_2_1) | v987&int32(7)<<(uint(int32(18))%32) | v963
	v999 = v985
	goto L180
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v914 - v999
	goto L192
L189:
	;
	v1001 = v995 - int32(97)
	if v1001 < int32(0) {
		goto L188
	} else {
		goto L190
	}
L190:
	;
	v1007 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1001)>>(uint(int32(3))%32)))+uint32(_c_F_r_lengthen_V_2[1]))))
	if int32(base.Ui32(v1007)>>(uint(v1001&int32(7))%32))&int32(1) == int32(0) {
		goto L188
	} else {
		goto L191
	}
L191:
	;
	v1031 = v999
	goto L175
L192:
	;
	goto L178
L193:
	;
	goto L130
L194:
	;
	if v1039 < int32(0) {
		v1070 = v1039
		goto L1
	} else {
		goto L195
	}
L195:
	;
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v1045 = F_insert_v(m, l0, v1043, v1043, v1044)
	mBase = m.M
	v1046 = m.ExcPending
	if v1046 != 0 {
		goto L22
	} else {
		goto L196
	}
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1043
	if int32(0) <= v1045 {
		goto L2
	} else {
		goto L197
	}
L197:
	;
	v1070 = v1045
	goto L1
L198:
	;
	if int32(0) <= v1052 {
		goto L2
	} else {
		goto L199
	}
L199:
	;
	v1070 = v1052
	goto L1
L200:
	;
	if v1058 < int32(0) {
		v1070 = v1058
		goto L1
	} else {
		goto L201
	}
L201:
	;
	goto L2
}
func F_r_measure_2(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v85 int32
	_ = v85
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v228 int32
	_ = v228
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v351 int32
	_ = v351
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v388 int32
	_ = v388
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v474 int32
	_ = v474
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v500 int32
	_ = v500
	var v511 int32
	_ = v511
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v617 int32
	_ = v617
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v643 int32
	_ = v643
	var v655 int32
	_ = v655
	var v662 int32
	_ = v662
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v727 int32
	_ = v727
	var v740 int32
	_ = v740
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v766 int32
	_ = v766
	var v777 int32
	_ = v777
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v5
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v5
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	for {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v25 <= v24 {
			v129 = int32(-1)
		} else {
			v41 = int32(1)
			v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v26))))
			if base.Ui32(v43) < base.Ui32(int32(192)) {
				v100 = v43
				v101 = v41
			} else {
				v47 = v24 + int32(1)
				if v47 == v25 {
					v100 = v43
					v101 = v41
				} else {
					v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47+v26))))
					v52 = v50 & int32(63)
					if base.Ui32(int32(224)) <= base.Ui32(v43) {
						v56 = v24 + int32(2)
						if v56 != v25 {
							v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+v26))))
							v68 = v66 & int32(63)
							if base.Ui32(int32(240)) <= base.Ui32(v43) {
								v72 = v24 + int32(3)
								if v72 != v25 {
									v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+v72))))
									v100 = v85&int32(63) | (v43<<(uint(int32(18))%32)&int32(_a_F_r_measure_2_0) | v52<<(uint(int32(12))%32) | v68<<(uint(int32(6))%32))
									v101 = int32(4)
								} else {
									v100 = v43<<(uint(int32(12))%32)&int32(_a_F_r_measure_2_1) | v52<<(uint(int32(6))%32) | v68
									v101 = int32(3)
								}
							} else {
								v100 = v43<<(uint(int32(12))%32)&int32(_a_F_r_measure_2_1) | v52<<(uint(int32(6))%32) | v68
								v101 = int32(3)
							}
						} else {
							v100 = v43<<(uint(int32(6))%32)&int32(1984) | v52
							v101 = int32(2)
						}
					} else {
						v100 = v43<<(uint(int32(6))%32)&int32(1984) | v52
						v101 = int32(2)
					}
				}
			}
			if int32(252) < v100 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v101 + v24
				v122 = int32(0)
			} else {
				v105 = v100 - int32(97)
				if v105 < int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v101 + v24
					v122 = int32(0)
				} else {
					v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v105)>>(uint(int32(3))%32)))+uint32(_c_F_r_measure_2[0]))))
					if int32(base.Ui32(v111)>>(uint(v105&int32(7))%32))&int32(1) != 0 {
						v122 = v101
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v101 + v24
						v122 = int32(0)
					}
				}
			}
			v129 = v122
		}
		if v129 == int32(0) {
			continue
		} else {
			break
		}
		break
	}
	v134 = int32(1)
	for {
		v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v138 = int32(2)
		v140 = int32(0)
		v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v142-v137 < v138 {
			v152 = v140
		} else {
			v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v148 = F_memcmp(m, v146+v137, int32(_a_F_r_measure_2_2), v138)
			mBase = m.M
			if v148 != 0 {
				v152 = v140
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v138 + v137
				v152 = int32(1)
			}
		}
		if v152 == int32(0) {
		} else {
			v134 = v134 - int32(1)
			continue
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v137
		v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v168 <= v137 {
			v273 = int32(-1)
		} else {
			v184 = int32(1)
			v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137+v169))))
			if base.Ui32(v186) < base.Ui32(int32(192)) {
				v243 = v186
				v244 = v184
			} else {
				v190 = v137 + int32(1)
				if v190 == v168 {
					v243 = v186
					v244 = v184
				} else {
					v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190+v169))))
					v195 = v193 & int32(63)
					if base.Ui32(int32(224)) <= base.Ui32(v186) {
						v199 = v137 + int32(2)
						if v199 != v168 {
							v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199+v169))))
							v211 = v209 & int32(63)
							if base.Ui32(int32(240)) <= base.Ui32(v186) {
								v215 = v137 + int32(3)
								if v215 != v168 {
									v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169+v215))))
									v243 = v228&int32(63) | (v186<<(uint(int32(18))%32)&int32(_a_F_r_measure_2_0) | v195<<(uint(int32(12))%32) | v211<<(uint(int32(6))%32))
									v244 = int32(4)
								} else {
									v243 = v186<<(uint(int32(12))%32)&int32(_a_F_r_measure_2_1) | v195<<(uint(int32(6))%32) | v211
									v244 = int32(3)
								}
							} else {
								v243 = v186<<(uint(int32(12))%32)&int32(_a_F_r_measure_2_1) | v195<<(uint(int32(6))%32) | v211
								v244 = int32(3)
							}
						} else {
							v243 = v186<<(uint(int32(6))%32)&int32(1984) | v195
							v244 = int32(2)
						}
					} else {
						v243 = v186<<(uint(int32(6))%32)&int32(1984) | v195
						v244 = int32(2)
					}
				}
			}
			if int32(252) < v243 {
				v266 = v244
			} else {
				v248 = v243 - int32(97)
				if v248 < int32(0) {
					v266 = v244
				} else {
					v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v248)>>(uint(int32(3))%32)))+uint32(_c_F_r_measure_2[0]))))
					if int32(base.Ui32(v254)>>(uint(v248&int32(7))%32))&int32(1) == int32(0) {
						v266 = v244
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v244 + v137
						v266 = int32(0)
					}
				}
			}
			v273 = v266
		}
		if v273 != 0 {
			break
		} else {
			v134 = v134 - int32(1)
			continue
		}
		break
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v137
	if int32(0) < v134 {
	} else {
		v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v292 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v291 <= v290 {
			v395 = int32(-1)
		} else {
			v307 = int32(1)
			v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290+v292))))
			if base.Ui32(v309) < base.Ui32(int32(192)) {
				v366 = v309
				v367 = v307
			} else {
				v313 = v290 + int32(1)
				if v313 == v291 {
					v366 = v309
					v367 = v307
				} else {
					v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v313+v292))))
					v318 = v316 & int32(63)
					if base.Ui32(int32(224)) <= base.Ui32(v309) {
						v322 = v290 + int32(2)
						if v322 != v291 {
							v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v322+v292))))
							v334 = v332 & int32(63)
							if base.Ui32(int32(240)) <= base.Ui32(v309) {
								v338 = v290 + int32(3)
								if v338 != v291 {
									v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292+v338))))
									v366 = v351&int32(63) | (v309<<(uint(int32(18))%32)&int32(_a_F_r_measure_2_0) | v318<<(uint(int32(12))%32) | v334<<(uint(int32(6))%32))
									v367 = int32(4)
								} else {
									v366 = v309<<(uint(int32(12))%32)&int32(_a_F_r_measure_2_1) | v318<<(uint(int32(6))%32) | v334
									v367 = int32(3)
								}
							} else {
								v366 = v309<<(uint(int32(12))%32)&int32(_a_F_r_measure_2_1) | v318<<(uint(int32(6))%32) | v334
								v367 = int32(3)
							}
						} else {
							v366 = v309<<(uint(int32(6))%32)&int32(1984) | v318
							v367 = int32(2)
						}
					} else {
						v366 = v309<<(uint(int32(6))%32)&int32(1984) | v318
						v367 = int32(2)
					}
				}
			}
			if int32(252) < v366 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v367 + v290
				v388 = int32(0)
			} else {
				v371 = v366 - int32(97)
				if v371 < int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v367 + v290
					v388 = int32(0)
				} else {
					v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v371)>>(uint(int32(3))%32)))+uint32(_c_F_r_measure_2[0]))))
					if int32(base.Ui32(v377)>>(uint(v371&int32(7))%32))&int32(1) != 0 {
						v388 = v367
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v367 + v290
						v388 = int32(0)
					}
				}
			}
			v395 = v388
		}
		if v395 != 0 {
		} else {
			v396 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v396
			for {
				v413 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v414 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v415 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v414 <= v413 {
					v518 = int32(-1)
				} else {
					v430 = int32(1)
					v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v413+v415))))
					if base.Ui32(v432) < base.Ui32(int32(192)) {
						v489 = v432
						v490 = v430
					} else {
						v436 = v413 + int32(1)
						if v436 == v414 {
							v489 = v432
							v490 = v430
						} else {
							v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v436+v415))))
							v441 = v439 & int32(63)
							if base.Ui32(int32(224)) <= base.Ui32(v432) {
								v445 = v413 + int32(2)
								if v445 != v414 {
									v455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v445+v415))))
									v457 = v455 & int32(63)
									if base.Ui32(int32(240)) <= base.Ui32(v432) {
										v461 = v413 + int32(3)
										if v461 != v414 {
											v474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415+v461))))
											v489 = v474&int32(63) | (v432<<(uint(int32(18))%32)&int32(_a_F_r_measure_2_0) | v441<<(uint(int32(12))%32) | v457<<(uint(int32(6))%32))
											v490 = int32(4)
										} else {
											v489 = v432<<(uint(int32(12))%32)&int32(_a_F_r_measure_2_1) | v441<<(uint(int32(6))%32) | v457
											v490 = int32(3)
										}
									} else {
										v489 = v432<<(uint(int32(12))%32)&int32(_a_F_r_measure_2_1) | v441<<(uint(int32(6))%32) | v457
										v490 = int32(3)
									}
								} else {
									v489 = v432<<(uint(int32(6))%32)&int32(1984) | v441
									v490 = int32(2)
								}
							} else {
								v489 = v432<<(uint(int32(6))%32)&int32(1984) | v441
								v490 = int32(2)
							}
						}
					}
					if int32(252) < v489 {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v490 + v413
						v511 = int32(0)
					} else {
						v494 = v489 - int32(97)
						if v494 < int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v490 + v413
							v511 = int32(0)
						} else {
							v500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v494)>>(uint(int32(3))%32)))+uint32(_c_F_r_measure_2[0]))))
							if int32(base.Ui32(v500)>>(uint(v494&int32(7))%32))&int32(1) != 0 {
								v511 = v490
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v490 + v413
								v511 = int32(0)
							}
						}
					}
					v518 = v511
				}
				if v518 == int32(0) {
					continue
				} else {
					break
				}
				break
			}
			v523 = int32(1)
			for {
				v526 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v527 = int32(2)
				v529 = int32(0)
				v531 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v531-v526 < v527 {
					v541 = v529
				} else {
					v535 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v537 = F_memcmp(m, v535+v526, int32(_a_F_r_measure_2_3), v527)
					mBase = m.M
					if v537 != 0 {
						v541 = v529
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v527 + v526
						v541 = int32(1)
					}
				}
				if v541 == int32(0) {
				} else {
					v523 = v523 - int32(1)
					continue
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v526
				v557 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v558 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v557 <= v526 {
					v662 = int32(-1)
				} else {
					v573 = int32(1)
					v575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v526+v558))))
					if base.Ui32(v575) < base.Ui32(int32(192)) {
						v632 = v575
						v633 = v573
					} else {
						v579 = v526 + int32(1)
						if v579 == v557 {
							v632 = v575
							v633 = v573
						} else {
							v582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v579+v558))))
							v584 = v582 & int32(63)
							if base.Ui32(int32(224)) <= base.Ui32(v575) {
								v588 = v526 + int32(2)
								if v588 != v557 {
									v598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v588+v558))))
									v600 = v598 & int32(63)
									if base.Ui32(int32(240)) <= base.Ui32(v575) {
										v604 = v526 + int32(3)
										if v604 != v557 {
											v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v558+v604))))
											v632 = v617&int32(63) | (v575<<(uint(int32(18))%32)&int32(_a_F_r_measure_2_0) | v584<<(uint(int32(12))%32) | v600<<(uint(int32(6))%32))
											v633 = int32(4)
										} else {
											v632 = v575<<(uint(int32(12))%32)&int32(_a_F_r_measure_2_1) | v584<<(uint(int32(6))%32) | v600
											v633 = int32(3)
										}
									} else {
										v632 = v575<<(uint(int32(12))%32)&int32(_a_F_r_measure_2_1) | v584<<(uint(int32(6))%32) | v600
										v633 = int32(3)
									}
								} else {
									v632 = v575<<(uint(int32(6))%32)&int32(1984) | v584
									v633 = int32(2)
								}
							} else {
								v632 = v575<<(uint(int32(6))%32)&int32(1984) | v584
								v633 = int32(2)
							}
						}
					}
					if int32(252) < v632 {
						v655 = v633
					} else {
						v637 = v632 - int32(97)
						if v637 < int32(0) {
							v655 = v633
						} else {
							v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v637)>>(uint(int32(3))%32)))+uint32(_c_F_r_measure_2[0]))))
							if int32(base.Ui32(v643)>>(uint(v637&int32(7))%32))&int32(1) == int32(0) {
								v655 = v633
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v633 + v526
								v655 = int32(0)
							}
						}
					}
					v662 = v655
				}
				if v662 != 0 {
					break
				} else {
					v523 = v523 - int32(1)
					continue
				}
				break
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v526
			if int32(0) < v523 {
			} else {
				v679 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v680 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v681 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v680 <= v679 {
					v784 = int32(-1)
				} else {
					v696 = int32(1)
					v698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v679+v681))))
					if base.Ui32(v698) < base.Ui32(int32(192)) {
						v755 = v698
						v756 = v696
					} else {
						v702 = v679 + int32(1)
						if v702 == v680 {
							v755 = v698
							v756 = v696
						} else {
							v705 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v702+v681))))
							v707 = v705 & int32(63)
							if base.Ui32(int32(224)) <= base.Ui32(v698) {
								v711 = v679 + int32(2)
								if v711 != v680 {
									v721 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v711+v681))))
									v723 = v721 & int32(63)
									if base.Ui32(int32(240)) <= base.Ui32(v698) {
										v727 = v679 + int32(3)
										if v727 != v680 {
											v740 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v681+v727))))
											v755 = v740&int32(63) | (v698<<(uint(int32(18))%32)&int32(_a_F_r_measure_2_0) | v707<<(uint(int32(12))%32) | v723<<(uint(int32(6))%32))
											v756 = int32(4)
										} else {
											v755 = v698<<(uint(int32(12))%32)&int32(_a_F_r_measure_2_1) | v707<<(uint(int32(6))%32) | v723
											v756 = int32(3)
										}
									} else {
										v755 = v698<<(uint(int32(12))%32)&int32(_a_F_r_measure_2_1) | v707<<(uint(int32(6))%32) | v723
										v756 = int32(3)
									}
								} else {
									v755 = v698<<(uint(int32(6))%32)&int32(1984) | v707
									v756 = int32(2)
								}
							} else {
								v755 = v698<<(uint(int32(6))%32)&int32(1984) | v707
								v756 = int32(2)
							}
						}
					}
					if int32(252) < v755 {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v756 + v679
						v777 = int32(0)
					} else {
						v760 = v755 - int32(97)
						if v760 < int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v756 + v679
							v777 = int32(0)
						} else {
							v766 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v760)>>(uint(int32(3))%32)))+uint32(_c_F_r_measure_2[0]))))
							if int32(base.Ui32(v766)>>(uint(v760&int32(7))%32))&int32(1) != 0 {
								v777 = v756
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v756 + v679
								v777 = int32(0)
							}
						}
					}
					v784 = v777
				}
				if v784 != 0 {
				} else {
					v785 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v785
				}
			}
		}
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v8
	return
}
func F_radians(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 float64
	_ = v4
	var v6 float64
	_ = v6
	var v8 float64
	_ = v8
	var v17 float64
	_ = v17
	var v20 int32
	_ = v20
	var v21 float64
	_ = v21
	var v27 float64
	_ = v27
	var v28 int32
	_ = v28
	var v30 float64
	_ = v30
	v4 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = base.F64_mul(v4, float64(0.017453292519943295))
	v8 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v6), v8)|base.F64_eq(base.F64_abs(v4), v8) == int32(0) {
		v17 = F_float_overflow_error_ext(m, int32(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int64(0)
		} else {
			v30 = float64(0)
			return base.I64_reinterpret_f64(v30)
		}
	} else {
		v21 = float64(0)
		if base.F64_eq(v4, v21)|base.F64_ne(v6, v21) != 0 {
			v30 = v6
			return base.I64_reinterpret_f64(v30)
		} else {
			v27 = F_float_underflow_error_ext(m, int32(0))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int64(0)
			} else {
				v30 = float64(0)
				return base.I64_reinterpret_f64(v30)
			}
		}
	}
}
func F_raise(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	v4 = m.G0
	v6 = v4 - int32(128)
	m.G0 = v6
	v10 = l0 - int32(1)
	if base.Ui32(v10) <= base.Ui32(int32(63)) {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(v10)>>(uint(int32(3))%32))&int32(536870908))+uint32(_c_F_raise[0])))
		v23 = int32(base.Ui32(v18)>>(uint(v10)%32)) & int32(1)
	} else {
		v23 = int32(0)
	}
	if v23 != 0 {
		v27 = l0 - int32(1)
		if base.B2i32(base.Ui32(v27) <= base.Ui32(int32(63)))&base.B2i32(base.Ui32(int32(2)) < base.Ui32(l0-int32(32))) == int32(0) {
			*(*int32)(unsafe.Add(mBase, _c_F_raise[1])) = int32(28)
		} else {
			v43 = int32(base.Ui32(v27)>>(uint(int32(3))%32)) & int32(536870908)
			v45 = *(*int32)(unsafe.Add(mBase, uint32(v43)+uint32(_c_F_raise[2])))
			*(*int32)(unsafe.Add(mBase, uint32(v43)+uint32(_c_F_raise[2]))) = v45 | int32(1)<<(uint(v27)%32)
		}
		m.G0 = v6 + int32(128)
		return
	} else {
		v52 = l0 * int32(20)
		v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+uint32(_c_F_raise[3]))))
		if v55&int32(4) != 0 {
			v58 = int32(0)
			base.MemoryFill(m, v6, v58, int32(128))
			v62 = *(*int32)(unsafe.Add(mBase, uint32(v52)+uint32(_c_F_raise[4])))
			m.T0[v62].(func(*base.Module, int32, int32, int32))(m, l0, v6, v58)
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return
			} else {
				m.G0 = v6 + int32(128)
				return
			}
		} else {
			v65 = *(*int32)(unsafe.Add(mBase, uint32(v52)+uint32(_c_F_raise[4])))
			switch v65 + int32(2) {
			case 0:
				m.G0 = v6 + int32(128)
				return
			default:
				m.Env.X__call_sighandler(m, v65, l0)
				mBase = m.M
				m.G0 = v6 + int32(128)
				return
			case 2:
				v72 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_raise[5])))
				if v72 == int32(0) {
					m.G0 = v6 + int32(128)
					return
				} else {
					m.T0[v72].(func(*base.Module, int32))(m, l0)
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return
					} else {
						m.G0 = v6 + int32(128)
						return
					}
				}
			}
		}
	}
}
func F_rangeTableEntry_used_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	v3 = int32(0)
	if l0 == v3 {
		v63 = v3
		return v63
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		switch v7 - int32(58) {
		case 0:
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v28 != 0 {
				v63 = v3
				return v63
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				if v29 != v30 {
					v63 = v3
					return v63
				} else {
					return int32(1)
				}
			}
		case 1, 2, 3, 4, 7, 8:
			v60 = F_expression_tree_walker_impl(m, l0, int32(1129), l1)
			mBase = m.M
			v61 = m.ExcPending
			if v61 != 0 {
				return int32(0)
			} else {
				v63 = v60
				return v63
			}
		case 5:
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			if v34 != v35 {
				v63 = v3
				return v63
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if v37 != 0 {
					v63 = v3
					return v63
				} else {
					return int32(1)
				}
			}
		case 6:
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			if v40 != v41 {
				v60 = F_expression_tree_walker_impl(m, l0, int32(1129), l1)
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int32(0)
				} else {
					v63 = v60
					return v63
				}
			} else {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if v43 != 0 {
					v60 = F_expression_tree_walker_impl(m, l0, int32(1129), l1)
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						v63 = v60
						return v63
					}
				} else {
					return int32(1)
				}
			}
		case 9:
			v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v46 + int32(1)
			v52 = F_query_tree_walker_impl(m, l0, int32(1129), l1, int32(0))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return int32(0)
			} else {
				v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v54 - int32(1)
				return v52
			}
		default:
			if v7 != int32(6) {
				v60 = F_expression_tree_walker_impl(m, l0, int32(1129), l1)
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int32(0)
				} else {
					v63 = v60
					return v63
				}
			} else {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if v12 == v13 {
					v15 = int32(1)
					v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					if v16 == v17 {
						v63 = v15
						return v63
					} else {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						v20 = F_bms_is_member(m, v16, v19)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int32(0)
						} else {
							if v20 != 0 {
								v63 = v15
								return v63
							} else {
								return int32(0)
							}
						}
					}
				} else {
					return int32(0)
				}
			}
		}
	}
}
func F_rbt_left_right_iterator(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v5 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v44
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v12 = v9
	goto L5
L3:
	;
	goto L4
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
	if v18 != int32(_a_F_rbt_left_right_iterator_0) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v12
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v15 != int32(_a_F_rbt_left_right_iterator_0) {
		v12 = v15
		goto L5
	} else {
		goto L7
	}
L6:
	;
	v44 = v12
	goto L1
L7:
	;
	goto L6
L8:
	;
	v23 = v18
	goto L11
L9:
	;
	goto L10
L10:
	;
	v32 = v5
	goto L14
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v23
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v26 != int32(_a_F_rbt_left_right_iterator_0) {
		v23 = v26
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v44 = v23
	goto L1
L14:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v33
	if v33 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v44 = v33
	goto L1
L16:
	;
	v37 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)) = uint8(v37)
	return int32(0)
L17:
	;
	goto L18
L18:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v41 != v32 {
		v32 = v33
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L15
}
func F_rbt_right_left_iterator(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v5 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v44
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v12 = v9
	goto L5
L3:
	;
	goto L4
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	if v18 != int32(_a_F_rbt_right_left_iterator_0) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v12
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if v15 != int32(_a_F_rbt_right_left_iterator_0) {
		v12 = v15
		goto L5
	} else {
		goto L7
	}
L6:
	;
	v44 = v12
	goto L1
L7:
	;
	goto L6
L8:
	;
	v23 = v18
	goto L11
L9:
	;
	goto L10
L10:
	;
	v32 = v5
	goto L14
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v23
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	if v26 != int32(_a_F_rbt_right_left_iterator_0) {
		v23 = v26
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v44 = v23
	goto L1
L14:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v33
	if v33 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v44 = v33
	goto L1
L16:
	;
	v37 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)) = uint8(v37)
	return int32(0)
L17:
	;
	goto L18
L18:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	if v41 != v32 {
		v32 = v33
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L15
}
func F_read(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l1
	v16 = m.Wasi_snapshot_preview1.Fd_read(m, l0, v7+int32(8), int32(1), v7+int32(4))
	mBase = m.M
	if v16 == int32(0) {
		v23 = int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_read[0])) = v16
		v23 = int32(-1)
	}
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	m.G0 = v7 + int32(16)
	if v23 != 0 {
		v29 = int32(-1)
	} else {
		v29 = v24
	}
	return v29
}
func F_readlink(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = v8 + int32(15)
	if l2 != 0 {
		v13 = l1
	} else {
		v13 = v12
	}
	v14 = int32(1)
	if base.Ui32(l2) <= base.Ui32(v14) {
		v17 = v14
	} else {
		v17 = l2
	}
	v18 = m.Env.X__syscall_readlinkat(m, int32(-100), l0, v13, v17)
	mBase = m.M
	if v13 == v12 {
		v23 = v18 >> (uint(int32(31)) % 32) & v18
	} else {
		v23 = v18
	}
	if base.Ui32(int32(-4095)) <= base.Ui32(v23) {
		*(*int32)(unsafe.Add(mBase, _c_F_readlink[0])) = int32(0) - v23
		v31 = int32(-1)
	} else {
		v31 = v23
	}
	m.G0 = v8 + int32(16)
	return v31
}
func F_readtup_cluster(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int64
	_ = v58
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l3
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v16 = F_tuplesort_readtup_alloc(m, l0, l3+int32(14))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v16))) = l3 - int32(10)
		*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v16 + int32(24)
		v27 = F_LogicalTapeRead(m, l2, v16+int32(4), int32(6))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return
		} else {
			if v27 == int32(6) {
				*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(0)
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
				v35 = F_LogicalTapeRead(m, l2, v33, v34)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
					if v35 != v37 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return
						} else {
							F_errmsg_internal(m, int32(_a_F_readtup_cluster_0), int32(0))
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_readtup_cluster_1), int32(1522), int32(_a_F_readtup_cluster_2))
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
						if v39&int32(1) != 0 {
							v45 = F_LogicalTapeRead(m, l2, v10+int32(12), int32(4))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return
							} else {
								if v45 != int32(4) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return
									} else {
										F_errmsg_internal(m, int32(_a_F_readtup_cluster_0), int32(0))
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_readtup_cluster_1), int32(1524), int32(_a_F_readtup_cluster_2))
											mBase = m.M
											v102 = m.ExcPending
											if v102 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l1))) = v16
									v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
									if v50 == int32(1) {
										v53 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
										v54 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53)+12)))
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
										v58 = F_heap_getattr_1(m, v16, v54, v55, l1+int32(16))
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v58
											m.G0 = v10 + int32(16)
											return
										}
									} else {
										m.G0 = v10 + int32(16)
										return
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v16
							v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
							if v50 == int32(1) {
								v53 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
								v54 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53)+12)))
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
								v58 = F_heap_getattr_1(m, v16, v54, v55, l1+int32(16))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v58
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
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(_a_F_readtup_cluster_0), int32(0))
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_readtup_cluster_1), int32(1518), int32(_a_F_readtup_cluster_2))
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
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
func F_readtup_heap_1(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int64
	_ = v50
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v13 = l3 + int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v13
	v15 = F_tuplesort_readtup_alloc(m, l0, v13)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v15))) = v13
		v21 = l3 - int32(4)
		v22 = F_LogicalTapeRead(m, l2, v15+int32(10), v21)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			if v22 == v21 {
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
				if v25&int32(1) != 0 {
					v31 = F_LogicalTapeRead(m, l2, v10+int32(28), int32(4))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						if v31 != int32(4) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return
							} else {
								F_errmsg_internal(m, int32(_a_F_readtup_heap_1_0), int32(0))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_readtup_heap_1_1), int32(1327), int32(_a_F_readtup_heap_1_2))
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v15
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
							v37 = int32(8)
							*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v15 - v37
							*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v36 + v37
							v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
							v46 = int32(*(*int16)(unsafe.Add(mBase, uint32(v45)+10)))
							v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
							v50 = F_heap_getattr_1(m, v10+v37, v46, v47, l1+int32(16))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v50
								m.G0 = v10 + int32(32)
								return
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v15
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
					v37 = int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v15 - v37
					*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v36 + v37
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					v46 = int32(*(*int16)(unsafe.Add(mBase, uint32(v45)+10)))
					v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
					v50 = F_heap_getattr_1(m, v10+v37, v46, v47, l1+int32(16))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v50
						m.G0 = v10 + int32(32)
						return
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(_a_F_readtup_heap_1_0), int32(0))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_readtup_heap_1_1), int32(1325), int32(_a_F_readtup_heap_1_2))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
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
func F_reapply_stacked_values(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
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
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if l2 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	F_reapply_stacked_values(m, l0, l1, v12, v13, v14, v15, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
	if l3 != v47 {
		goto L15
	} else {
		goto L16
	}
L5:
	;
	return
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	switch v20 {
	case 0:
		v35 = int32(2)
		goto L8
	case 1:
		goto L11
	case 2:
		goto L10
	case 3:
		goto L9
	default:
		goto L7
	}
L7:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v43 == v11 {
		goto L1
	} else {
		goto L14
	}
L8:
	;
	v36 = int32(0)
	v40 = F_set_config_with_handle(m, v10, v36, l3, l4, l5, l6, v35, int32(1), int32(19), v36)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L5
	} else {
		goto L13
	}
L9:
	;
	v23 = int32(1)
	v24 = int32(0)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v33 = F_set_config_with_handle(m, v10, v24, v25, v26, int32(13), v28, v24, v23, int32(19), v24)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L5
	} else {
		goto L12
	}
L10:
	;
	v35 = int32(1)
	goto L8
L11:
	;
	v35 = int32(0)
	goto L8
L12:
	;
	v35 = v23
	goto L8
L13:
	;
	goto L7
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v45
	return
L15:
	;
	v55 = int32(0)
	v60 = F_set_config_with_handle(m, v10, v55, l3, l4, l5, l6, v55, int32(1), int32(19), v55)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L5
	} else {
		goto L20
	}
L16:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	if l4 != v49 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if l5 != v51 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if l6 == v53 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	goto L15
L20:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v62 == int32(0) {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v69 = int32(_a_F_reapply_stacked_values_0)
	goto L24
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = int32(0)
	goto L1
L23:
	;
	goto L22
L24:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	if v72 == int32(0) {
		goto L23
	} else {
		goto L26
	}
L25:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	*(*int32)(unsafe.Add(mBase, uint32(v69))) = v76
	goto L23
L26:
	;
	if v72 != l0+int32(76) {
		v69 = v72
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
}
func F_record_cmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
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
	var v41 int32
	_ = v41
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
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
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
	var v89 int64
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
	var v299 int64
	_ = v299
	var v300 int32
	_ = v300
	var v306 int64
	_ = v306
	var v313 int32
	_ = v313
	var v314 int64
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v373 int32
	_ = v373
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v483 int32
	_ = v483
	v23 = m.G0
	v25 = v23 - int32(128)
	m.G0 = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v28 = F_pg_detoast_datum(m, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v33 = F_pg_detoast_datum(m, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v39 = F_lookup_rowtype_tupdesc(m, v37, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v44 = F_lookup_rowtype_tupdesc(m, v42, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+124)) = v28
	v49 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+120)) = v49
	*(*uint16)(unsafe.Add(mBase, uint32(v25)+116)) = uint16(v49)
	v53 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+112)) = v53
	v55 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+108)) = int32(base.Ui32(v47) >> (uint(v55) % 32))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+104)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v25)+100)) = v49
	*(*uint16)(unsafe.Add(mBase, uint32(v25)+96)) = uint16(v49)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+92)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v25)+88)) = int32(base.Ui32(v58) >> (uint(v55) % 32))
	if v46 < v41 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v70 = v41
	goto L9
L8:
	;
	v70 = v46
	goto L9
L9:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+16))
	if v72 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v37 != v95 {
		goto L16
	} else {
		goto L17
	}
L11:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v71)+20))
	v83 = F_MemoryContextAlloc(m, v78, v70<<(uint(int32(2))%32)+int32(20))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	if v75 < v70 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	v94 = v72
	v95 = v77
	goto L10
L14:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+16)) = v83
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+16))
	v89 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+4)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v88))) = v70
	*(*int64)(unsafe.Add(mBase, uint32(v88)+12)) = v89
	v94 = v88
	v95 = int32(0)
	goto L10
L15:
	;
	v151 = F_palloc(m, v41<<(uint(int32(3))%32))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L31
	}
L16:
	;
	v104 = v94 + int32(20)
	v108 = v70 << (uint(int32(2)) % 32)
	if v104&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v108)) == int32(0) {
		goto L22
	} else {
		goto L23
	}
L17:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
	if v97 != v38 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v94)+12))
	if v99 != v42 {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v94)+16))
	if v101 == v43 {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	goto L16
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94)+16)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v94)+12)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v94)+8)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = v37
	goto L15
L22:
	;
	if v108 == int32(0) {
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if v108 == int32(0) {
		goto L21
	} else {
		goto L30
	}
L25:
	;
	v118 = v94 + v108 + int32(20)
	v120 = v94 + int32(24)
	if base.Ui32(v120) < base.Ui32(v118) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v122 = v118
	goto L28
L27:
	;
	v122 = v120
	goto L28
L28:
	;
	v129 = (v122-v94-int32(21))&int32(-4) + int32(4)
	if v129 == int32(0) {
		goto L21
	} else {
		goto L29
	}
L29:
	;
	base.MemoryFill(m, v104, int32(0), v129)
	goto L21
L30:
	;
	base.MemoryFill(m, v104, int32(0), v108)
	goto L21
L31:
	;
	v153 = F_palloc(m, v41)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_heap_deform_tuple(m, v25+int32(108), v39, v151, v153)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v161 = F_palloc(m, v46<<(uint(int32(3))%32))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v163 = F_palloc(m, v46)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_heap_deform_tuple(m, v25+int32(88), v44, v161, v163)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v167 = int32(0)
	v173 = base.B2i32(v167 < v41)
	if v167 < v41 {
		goto L43
	} else {
		goto L44
	}
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L1
	} else {
		goto L113
	}
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L1
	} else {
		goto L108
	}
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L102
	}
L40:
	;
	F_pfree(m, v151)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L82
	}
L41:
	;
	if base.B2i32(v346 != v41)|base.B2i32(v344 != v46) != 0 {
		goto L37
	} else {
		goto L81
	}
L42:
	;
	v183 = v180
	v185 = v167
	v186 = base.B2i32(v167 < v46)
	v187 = v173
	v188 = v181
	goto L47
L43:
	;
	v174 = int32(0)
	v180 = v174
	v181 = v174
	goto L42
L44:
	;
	goto L45
L45:
	;
	v176 = int32(0)
	if v46 <= v176 {
		v344 = v176
		v346 = v167
		goto L41
	} else {
		goto L46
	}
L46:
	;
	v180 = v176
	v181 = v176
	goto L42
L47:
	;
	if v187&int32(1) == int32(0) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v344 = v331
	v346 = v332
	goto L41
L49:
	;
	v340 = base.B2i32(v331 < v46)
	v341 = base.B2i32(v332 < v41)
	if v340|v341 != 0 {
		v183 = v331
		v185 = v332
		v186 = v340
		v187 = v341
		v188 = v335
		goto L47
	} else {
		goto L80
	}
L50:
	;
	if v186&int32(1) == int32(0) {
		v344 = v183
		v346 = v185
		goto L41
	} else {
		goto L53
	}
L51:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+v208<<(uint(int32(3))%32)+v185*int32(100))+119)))
	if v215 != int32(1) {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v331 = v183
	v332 = v185 + int32(1)
	v335 = v188
	goto L49
L53:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v227 = v44 + v224<<(uint(int32(3))%32)
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227+v183*int32(100))+119)))
	if v231 == int32(1) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v331 = v183 + int32(1)
	v332 = v185
	v335 = v188
	goto L49
L55:
	;
	goto L56
L56:
	;
	if v187&int32(1) == int32(0) {
		v344 = v183
		v346 = v185
		goto L41
	} else {
		goto L57
	}
L57:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v244 = int32(100)
	v246 = v39 + v240<<(uint(int32(3))%32) + v185*v244
	v247 = int32(28)
	v248 = v246 + v247
	v251 = v227 + v183*v244
	v253 = v251 + v247
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v246)+96))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v251)+96))
	if v254 != v255 {
		goto L39
	} else {
		goto L58
	}
L58:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v253)+96))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v248)+96))
	v261 = v94 + int32(20) + v188<<(uint(int32(2))%32)
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v261)))
	if v262 != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185+v153))))
	if v274 == int32(1) {
		goto L67
	} else {
		goto L68
	}
L60:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v262)))
	if v263 == v254 {
		v272 = v262
		goto L59
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v266 = F_lookup_type_cache(m, v254, int32(64))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L64
	}
L63:
	;
	goto L62
L64:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v266)+108))
	if v268 == int32(0) {
		goto L38
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v261))) = v266
	v272 = v266
	goto L59
L66:
	;
	v325 = int32(1)
	v331 = v183 + v325
	v332 = v185 + v325
	v335 = v188 + v325
	goto L49
L67:
	;
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183+v163))))
	if v278 != 0 {
		goto L66
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183+v163))))
	if v281 != 0 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v373 = int32(1)
	goto L40
L71:
	;
	v373 = int32(-1)
	goto L40
L72:
	;
	goto L73
L73:
	;
	v283 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v25)+50)) = uint16(v283)
	v285 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+48)) = uint8(v285)
	if v258 == v257 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v289 = v258
	goto L76
L75:
	;
	v289 = v285
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+44)) = v289
	*(*int64)(unsafe.Add(mBase, uint32(v25)+36)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = v272 + int32(104)
	v296 = int32(3)
	v299 = *(*int64)(unsafe.Add(mBase, uint32(v151+v185<<(uint(v296)%32))))
	v300 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+64)) = uint8(v300)
	*(*int64)(unsafe.Add(mBase, uint32(v25)+56)) = v299
	v306 = *(*int64)(unsafe.Add(mBase, uint32(v161+v183<<(uint(v296)%32))))
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+80)) = uint8(v300)
	*(*int64)(unsafe.Add(mBase, uint32(v25)+72)) = v306
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v272)+104))
	v314 = m.T0[v313].(func(*base.Module, int32) int64)(m, v25+int32(32))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v316 = base.I32_wrap_i64(v314)
	if v316 < int32(0) {
		v373 = int32(-1)
		goto L40
	} else {
		goto L78
	}
L78:
	;
	if v316 == int32(0) {
		goto L66
	} else {
		goto L79
	}
L79:
	;
	v373 = int32(1)
	goto L40
L80:
	;
	goto L48
L81:
	;
	v373 = int32(0)
	goto L40
L82:
	;
	F_pfree(m, v153)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_pfree(m, v161)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	F_pfree(m, v163)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	if int32(0) <= v399 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	F_DecrTupleDescRefCount(m, v39)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	if int32(0) <= v404 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	goto L88
L90:
	;
	F_DecrTupleDescRefCount(m, v44)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L1
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v409 != v28 {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	goto L92
L94:
	;
	F_pfree(m, v28)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v413 != v33 {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	goto L96
L98:
	;
	F_pfree(m, v33)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	m.G0 = v25 + int32(128)
	return v373
L101:
	;
	goto L100
L102:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v248)+68))
	v429 = F_format_type_be(m, v428)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v253)+68))
	v432 = F_format_type_be(m, v431)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v188 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v429
	F_errmsg(m, int32(_a_F_record_cmp_0), v25+int32(16))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(_a_F_record_cmp_1), int32(952), int32(_a_F_record_cmp_2))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L108:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v266)))
	v457 = F_format_type_be(m, v456)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v457
	F_errmsg(m, int32(_a_F_record_cmp_3), v25)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_record_cmp_1), int32(975), int32(_a_F_record_cmp_2))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L113:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	F_errmsg(m, int32(_a_F_record_cmp_4), int32(0))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_record_cmp_1), int32(1040), int32(_a_F_record_cmp_2))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_record_image_eq(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
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
	var v85 int64
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
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
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v256 int64
	_ = v256
	var v260 int64
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v326 int64
	_ = v326
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
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
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
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
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	v20 = int64(0)
	v21 = m.G0
	v23 = v21 + int32(-64)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v26 = F_pg_detoast_datum(m, v25)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v31 = F_pg_detoast_datum(m, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v35 = F_lookup_rowtype_tupdesc(m, v33, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v40 = F_lookup_rowtype_tupdesc(m, v38, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+60)) = v26
	v45 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+56)) = v45
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+52)) = uint16(v45)
	v49 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v49
	v51 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = int32(base.Ui32(v43) >> (uint(v51) % 32))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v45
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+32)) = uint16(v45)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = int32(base.Ui32(v54) >> (uint(v51) % 32))
	if v42 < v37 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v66 = v37
	goto L8
L7:
	;
	v66 = v42
	goto L8
L8:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
	if v68 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v33 != v91 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	v79 = F_MemoryContextAlloc(m, v74, v66<<(uint(int32(2))%32)+int32(20))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	if v71 < v66 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v90 = v68
	v91 = v73
	goto L9
L13:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v81)+16)) = v79
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+16))
	v85 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v84)+4)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v84))) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v84)+12)) = v85
	v90 = v84
	v91 = int32(0)
	goto L9
L14:
	;
	v147 = F_palloc(m, v37<<(uint(int32(3))%32))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L30
	}
L15:
	;
	v100 = v90 + int32(20)
	v104 = v66 << (uint(int32(2)) % 32)
	if v100&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v104)) == int32(0) {
		goto L21
	} else {
		goto L22
	}
L16:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v90)+8))
	if v93 != v34 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
	if v95 != v38 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
	if v97 == v39 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L15
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v90)+16)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v90)+12)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v90)+8)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v90)+4)) = v33
	goto L14
L21:
	;
	if v104 == int32(0) {
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	if v104 == int32(0) {
		goto L20
	} else {
		goto L29
	}
L24:
	;
	v114 = v90 + v104 + int32(20)
	v116 = v90 + int32(24)
	if base.Ui32(v116) < base.Ui32(v114) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v118 = v114
	goto L27
L26:
	;
	v118 = v116
	goto L27
L27:
	;
	v125 = (v118-v90-int32(21))&int32(-4) + int32(4)
	if v125 == int32(0) {
		goto L20
	} else {
		goto L28
	}
L28:
	;
	base.MemoryFill(m, v100, int32(0), v125)
	goto L20
L29:
	;
	base.MemoryFill(m, v100, int32(0), v104)
	goto L20
L30:
	;
	v149 = F_palloc(m, v37)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_heap_deform_tuple(m, v21+int32(-20), v35, v147, v149)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v157 = F_palloc(m, v42<<(uint(int32(3))%32))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v159 = F_palloc(m, v42)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_heap_deform_tuple(m, v21+int32(-40), v40, v157, v159)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v163 = int32(0)
	v167 = base.B2i32(v163 < v37)
	if v163 < v37 {
		goto L41
	} else {
		goto L42
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L93
	}
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L87
	}
L38:
	;
	F_pfree(m, v147)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L67
	}
L39:
	;
	if base.B2i32(v285 != v37)|base.B2i32(v284 != v42) != 0 {
		goto L36
	} else {
		goto L66
	}
L40:
	;
	v177 = v174
	v178 = v163
	v180 = v167
	v181 = v175
	v183 = base.B2i32(v163 < v42)
	goto L45
L41:
	;
	v168 = int32(0)
	v174 = v168
	v175 = v168
	goto L40
L42:
	;
	goto L43
L43:
	;
	v170 = int32(0)
	if v42 <= v170 {
		v284 = v170
		v285 = v163
		goto L39
	} else {
		goto L44
	}
L44:
	;
	v174 = v170
	v175 = v170
	goto L40
L45:
	;
	if v180 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	v284 = v273
	v285 = v274
	goto L39
L47:
	;
	v280 = base.B2i32(v273 < v42)
	v281 = base.B2i32(v274 < v37)
	if v280|v281 != 0 {
		v177 = v273
		v178 = v274
		v180 = v281
		v181 = v276
		v183 = v280
		goto L45
	} else {
		goto L65
	}
L48:
	;
	if v183 == int32(0) {
		v284 = v177
		v285 = v178
		goto L39
	} else {
		goto L51
	}
L49:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35+v198<<(uint(int32(3))%32)+v178*int32(100))+119)))
	if v205 != int32(1) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v273 = v177
	v274 = v178 + int32(1)
	v276 = v181
	goto L47
L51:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v215 = v40 + v212<<(uint(int32(3))%32)
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215+v177*int32(100))+119)))
	if v219 == int32(1) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v273 = v177 + int32(1)
	v274 = v178
	v276 = v181
	goto L47
L53:
	;
	goto L54
L54:
	;
	if v180 == int32(0) {
		v284 = v177
		v285 = v178
		goto L39
	} else {
		goto L55
	}
L55:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v230 = int32(100)
	v232 = v35 + v226<<(uint(int32(3))%32) + v178*v230
	v233 = int32(28)
	v234 = v232 + v233
	v237 = v215 + v177*v230
	v239 = v237 + v233
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v232)+96))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v237)+96))
	if v240 != v241 {
		goto L37
	} else {
		goto L56
	}
L56:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177+v159))))
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178+v149))))
	if v246 == int32(1) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v267 = int32(1)
	v273 = v177 + v267
	v274 = v178 + v267
	v276 = v181 + v267
	goto L47
L58:
	;
	if v244&int32(1) != 0 {
		goto L57
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	if v244&int32(1) != 0 {
		v326 = v20
		goto L38
	} else {
		goto L62
	}
L61:
	;
	v326 = v20
	goto L38
L62:
	;
	v253 = int32(3)
	v256 = *(*int64)(unsafe.Add(mBase, uint32(v147+v178<<(uint(v253)%32))))
	v260 = *(*int64)(unsafe.Add(mBase, uint32(v157+v177<<(uint(v253)%32))))
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234)+82)))
	v262 = int32(*(*int16)(unsafe.Add(mBase, uint32(v239)+72)))
	v263 = F_datum_image_eq(m, v256, v260, v261, v262)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	if v263 == int32(0) {
		v326 = v20
		goto L38
	} else {
		goto L64
	}
L64:
	;
	goto L57
L65:
	;
	goto L46
L66:
	;
	v326 = int64(1)
	goto L38
L67:
	;
	F_pfree(m, v149)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_pfree(m, v157)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	F_pfree(m, v159)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	if int32(0) <= v335 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	F_DecrTupleDescRefCount(m, v35)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	if int32(0) <= v340 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	goto L73
L75:
	;
	F_DecrTupleDescRefCount(m, v40)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v345 != v26 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	goto L77
L79:
	;
	F_pfree(m, v26)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v349 != v31 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	goto L81
L83:
	;
	F_pfree(m, v31)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	m.G0 = v23 - int32(-64)
	return v326
L86:
	;
	goto L85
L87:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v234)+68))
	v365 = F_format_type_be(m, v364)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v239)+68))
	v368 = F_format_type_be(m, v367)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v181 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v368
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v365
	F_errmsg(m, int32(_a_F_record_image_eq_0), v23)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(_a_F_record_image_eq_1), int32(1720), int32(_a_F_record_image_eq_2))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L93:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	F_errmsg(m, int32(_a_F_record_image_eq_3), int32(0))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_record_image_eq_1), int32(1753), int32(_a_F_record_image_eq_2))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_record_image_ge(m *base.Module, l0 int32) int64 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_record_image_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(base.B2i32(int32(0) <= v2))
	}
}
func F_record_send(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
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
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v263 int32
	_ = v263
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v345 int64
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v376 int32
	_ = v376
	var v382 int32
	_ = v382
	var v388 int32
	_ = v388
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	v2 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(48)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v21 = F_pg_detoast_datum(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v29 = F_lookup_rowtype_tupdesc(m, v27, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v21
	v34 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v34
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+36)) = uint16(v34)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = int32(base.Ui32(v32) >> (uint(int32(2)) % 32))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
	if v45 == v34 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v65 == v27 {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v44)+20))
	v56 = F_MemoryContextAlloc(m, v51, v31*int32(44)+int32(12))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	if v48 != v31 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	v64 = v45
	v65 = v50
	goto L5
L9:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+16)) = v56
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v61))) = int64(0)
	v64 = v61
	v65 = v34
	goto L5
L10:
	;
	v114 = F_palloc_mul(m, int32(8), v31)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L25
	}
L11:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if v67 == v28 {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v72 = v31 * int32(44)
	v74 = v72 + int32(12)
	if v64&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v74)) == int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L13
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+8)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v27
	goto L10
L16:
	;
	if v74 == int32(0) {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	if v74 == int32(0) {
		goto L15
	} else {
		goto L24
	}
L19:
	;
	v86 = v64 + v72 + int32(12)
	v88 = v64 + int32(4)
	if base.Ui32(v88) < base.Ui32(v86) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v90 = v86
	goto L22
L21:
	;
	v90 = v88
	goto L22
L22:
	;
	v95 = (v64^int32(-1)+v90)&int32(-4) + int32(4)
	if v95 == int32(0) {
		goto L15
	} else {
		goto L23
	}
L23:
	;
	base.MemoryFill(m, v64, int32(0), v95)
	goto L15
L24:
	;
	base.MemoryFill(m, v64, int32(0), v74)
	goto L15
L25:
	;
	v117 = F_palloc_mul(m, int32(1), v31)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_heap_deform_tuple(m, v18+int32(28), v29, v114, v117)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_pq_begintypsend(m, v18+int32(12))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	if int32(0) < v31 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	F_pfree(m, v114)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L1
	} else {
		goto L64
	}
L30:
	;
	v127 = int32(3)
	v128 = v31 & v127
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v132 = v29 + v129<<(uint(v127)%32)
	v133 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v31) {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	goto L32
L32:
	;
	F_enlargeStringInfo(m, v18+int32(12), int32(4))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L63
	}
L33:
	;
	F_enlargeStringInfo(m, v18+int32(12), int32(4))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L44
	}
L34:
	;
	v141 = v133
	v143 = v133
	v153 = v2
	goto L37
L35:
	;
	v182 = v133
	v184 = v133
	goto L36
L36:
	;
	v197 = v182
	v199 = v184
	v207 = v2
	goto L41
L37:
	;
	v156 = v132 + v141*int32(100)
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+119)))
	v158 = int32(1)
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+219)))
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+319)))
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+419)))
	v172 = v143 + (v157 ^ v158) + (v161 ^ v158) + (v165 ^ v158) + (v169 ^ v158)
	v173 = int32(4)
	v174 = v141 + v173
	v176 = v153 + v173
	if v176 != v31&int32(2147483644) {
		v141 = v174
		v143 = v172
		v153 = v176
		goto L37
	} else {
		goto L39
	}
L38:
	;
	if v128 == int32(0) {
		v226 = v172
		goto L33
	} else {
		goto L40
	}
L39:
	;
	goto L38
L40:
	;
	v182 = v174
	v184 = v172
	goto L36
L41:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132+v197*int32(100))+119)))
	v214 = int32(1)
	v216 = v199 + (v213 ^ v214)
	v220 = v207 + v214
	if v220 != v128 {
		v197 = v197 + v214
		v199 = v216
		v207 = v220
		goto L41
	} else {
		goto L43
	}
L42:
	;
	v226 = v216
	goto L33
L43:
	;
	goto L42
L44:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v247 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v242+v243))) = base.I32_rotr(v226, int32(24))&v247 | base.I32_rotr(v226&v247, int32(8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v242 + int32(4)
	v263 = int32(0)
	goto L45
L45:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v282 = v29 + v276<<(uint(int32(3))%32) + v263*int32(100)
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282)+119)))
	if v283 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L29
L47:
	;
	v388 = v263 + int32(1)
	if v388 != v31 {
		v263 = v388
		goto L45
	} else {
		goto L62
	}
L48:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v282)+96))
	v286 = v18 + int32(12)
	F_enlargeStringInfo(m, v286, int32(4))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v295 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v290+v291))) = base.I32_rotr(v284, int32(24))&v295 | base.I32_rotr(v284&v295, int32(8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v290 + int32(4)
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263+v117))))
	if v307 == int32(1) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	F_enlargeStringInfo(m, v286, int32(4))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v323 = v64 + int32(12) + v263*int32(44)
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v323)))
	if v284 != v324 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v313+v314))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v313 + int32(4)
	goto L47
L54:
	;
	F_getTypeBinaryOutputInfo(m, v284, v323+int32(4), v323+int32(12))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v345 = *(*int64)(unsafe.Add(mBase, uint32(v114+v263<<(uint(int32(3))%32))))
	v346 = F_SendFunctionCall(m, v323+int32(16), v345)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L59
	}
L57:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v323)+4))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v335)+20))
	F_fmgr_info_cxt(m, v332, v323+int32(16), v336)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323))) = v284
	goto L56
L59:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v346)))
	v350 = v18 + int32(12)
	F_enlargeStringInfo(m, v350, int32(4))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v357 = int32(2)
	v359 = int32(4)
	v360 = int32(base.Ui32(v348)>>(uint(v357)%32)) - v359
	v361 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v354+v355))) = base.I32_rotr(v360&v361, int32(8)) | base.I32_rotr(v360, int32(24))&v361
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v354 + v359
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v346)))
	F_appendBinaryStringInfo(m, v350, v346+v359, int32(base.Ui32(v376)>>(uint(v357)%32))-v359)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	goto L47
L62:
	;
	goto L46
L63:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v395+v396))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v395 + int32(4)
	goto L29
L64:
	;
	F_pfree(m, v117)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	if int32(0) <= v422 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	F_DecrTupleDescRefCount(m, v29)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v428 = v18 + int32(12)
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v428)))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v428)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v430))) = v431 << (uint(int32(2)) % 32)
	goto L70
L69:
	;
	goto L68
L70:
	;
	m.G0 = v18 + int32(48)
	return base.I64_extend_i32_u(v430)
}
func F_record_smaller(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	v4 = F_record_cmp(m, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		if v4 < int32(0) {
			v10 = int32(24)
		} else {
			v10 = int32(40)
		}
		v12 = *(*int64)(unsafe.Add(mBase, uint32(l0+v10)))
		return v12
	}
}
func F_recurse_push_qual(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = l0
	goto L1
L1:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v19 != int32(142) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L8
	} else {
		goto L11
	}
L3:
	;
	goto L2
L4:
	;
	if v19 != int32(63) {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	F_recurse_push_qual(m, v39, l1, l2, l3, l4)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L8
	} else {
		goto L10
	}
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v25+v26<<(uint(int32(2))%32)-int32(4))))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+36))
	F_subquery_push_qual(m, v33, l2, l3, l4)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	m.G0 = v10 + int32(16)
	return
L10:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v12 = v42
	goto L1
L11:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v47
	F_errmsg_internal(m, int32(_a_F_recurse_push_qual_0), v10)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	F_errfinish(m, int32(_a_F_recurse_push_qual_1), int32(_a_F_recurse_push_qual_2), int32(_a_F_recurse_push_qual_3))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_reduce_outer_joins_pass2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
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
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
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
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v339 int64
	_ = v339
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v377 int32
	_ = v377
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v525 int32
	_ = v525
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v621 int32
	_ = v621
	var v628 int32
	_ = v628
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v693 int32
	_ = v693
	var v706 int32
	_ = v706
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v752 int32
	_ = v752
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v781 int32
	_ = v781
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v889 int32
	_ = v889
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v918 int32
	_ = v918
	var v928 int32
	_ = v928
	var v952 int32
	_ = v952
	var v956 int32
	_ = v956
	var v961 int32
	_ = v961
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v970 int32
	_ = v970
	var v975 int32
	_ = v975
	v7 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(32)
	m.G0 = v20
	if l0 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L10
	} else {
		goto L255
	}
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v22 - int32(63) {
	case 0:
		goto L9
	case 1:
		goto L7
	case 2:
		goto L8
	default:
		goto L1
	}
L3:
	;
	goto L4
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L10
	} else {
		goto L252
	}
L5:
	;
	m.G0 = v20 + int32(32)
	return
L6:
	;
	F_bms_free(m, v918)
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
		goto L10
	} else {
		goto L251
	}
L7:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+12))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v106 {
	case 0, 4, 5:
		v354 = v106
		v356 = v104
		v357 = v105
		goto L31
	case 1:
		goto L38
	case 2:
		goto L36
	case 3:
		goto L37
	default:
		goto L35
	}
L8:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v39 = F_find_nonnullable_rels(m, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L10
	} else {
		goto L14
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return
L11:
	;
	F_errmsg_internal(m, int32(_a_F_reduce_outer_joins_pass2_0), int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	F_errfinish(m, int32(_a_F_reduce_outer_joins_pass2_1), int32(3522), int32(_a_F_reduce_outer_joins_pass2_2))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L14:
	;
	v41 = F_bms_add_members(m, v39, l4)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v44 = F_find_forced_null_vars(m, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	v46 = F_mbms_add_members(m, v44, l5)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v51 = int32(0)
	goto L18
L18:
	;
	v68 = int32(0)
	if v49 == v68 {
		v78 = v68
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if v48 == int32(0) {
		v918 = v41
		goto L6
	} else {
		goto L23
	}
L21:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v72 <= v51 {
		v78 = int32(0)
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	v78 = v74 + v51<<(uint(int32(2))%32)
	goto L20
L23:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	if base.B2i32(v78 == int32(0))|base.B2i32(v83 <= v51) != 0 {
		v918 = v41
		goto L6
	} else {
		goto L24
	}
L24:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	if v86 == int32(0) {
		v918 = v41
		goto L6
	} else {
		goto L25
	}
L25:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v86+v51<<(uint(int32(2))%32))))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+4)))
	if v93 == int32(1) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	F_reduce_outer_joins_pass2(m, v96, v92, l2, l3, v41, v46)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L10
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v51 = v51 + int32(1)
	goto L18
L29:
	;
	goto L28
L30:
	;
	if v101 == int32(0) {
		goto L215
	} else {
		goto L216
	}
L31:
	;
	if base.B2i32(l5 == int32(0))|base.B2i32(v354 != int32(1)) != 0 {
		v827 = v354
		v830 = v356
		v832 = v357
		goto L30
	} else {
		goto L106
	}
L32:
	;
	v354 = int32(1)
	v356 = v350
	v357 = v351
	goto L31
L33:
	;
	v339 = *(*int64)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = base.I64_rotl(v339, int64(32))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)+12))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v344)+4))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v344)))
	v350 = v346
	v351 = v345
	goto L32
L34:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v329 = F_palloc(m, int32(8))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L10
	} else {
		goto L104
	}
L35:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L10
	} else {
		goto L101
	}
L36:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v207 = int32(0)
	if base.B2i32(l4 == v207)|base.B2i32(v206 == v207) != 0 {
		v252 = v207
		goto L67
	} else {
		goto L68
	}
L37:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v157 = int32(0)
	if base.B2i32(l4 == v157)|base.B2i32(v156 == v157) != 0 {
		v202 = v157
		goto L53
	} else {
		goto L54
	}
L38:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	v108 = int32(0)
	if base.B2i32(l4 == v108)|base.B2i32(v107 == v108) != 0 {
		v153 = v108
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v354 = v153 ^ int32(1)
	v356 = v104
	v357 = v105
	goto L31
L40:
	;
	goto L39
L41:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	if v118 < v119 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v121 = v118
	goto L44
L43:
	;
	v121 = v119
	goto L44
L44:
	;
	if v121 <= int32(1) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v124 = int32(1)
	goto L47
L46:
	;
	v124 = v121
	goto L47
L47:
	;
	v125 = int32(8)
	v130 = int32(0)
	goto L48
L48:
	;
	v137 = v130 << (uint(int32(2)) % 32)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v107+v125+v137)))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l4+v125+v137)))
	v142 = v139 & v141
	v144 = base.B2i32(v142 != int32(0))
	if v142 != 0 {
		v153 = v144
		goto L40
	} else {
		goto L50
	}
L49:
	;
	v153 = v144
	goto L40
L50:
	;
	v146 = v130 + int32(1)
	if v146 != v124 {
		v130 = v146
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	if v202 == int32(0) {
		goto L33
	} else {
		goto L65
	}
L53:
	;
	goto L52
L54:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	if v167 < v168 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v170 = v167
	goto L57
L56:
	;
	v170 = v168
	goto L57
L57:
	;
	if v170 <= int32(1) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v173 = int32(1)
	goto L60
L59:
	;
	v173 = v170
	goto L60
L60:
	;
	v174 = int32(8)
	v179 = int32(0)
	goto L61
L61:
	;
	v186 = v179 << (uint(int32(2)) % 32)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v156+v174+v186)))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l4+v174+v186)))
	v191 = v188 & v190
	v193 = base.B2i32(v191 != int32(0))
	if v191 != 0 {
		v202 = v193
		goto L53
	} else {
		goto L63
	}
L62:
	;
	v202 = v193
	goto L53
L63:
	;
	v195 = v179 + int32(1)
	if v195 != v173 {
		v179 = v195
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	v827 = int32(0)
	v830 = v104
	v832 = v105
	goto L30
L66:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	v254 = int32(0)
	if base.B2i32(l4 == v254)|base.B2i32(v253 == v254) != 0 {
		v299 = v254
		goto L80
	} else {
		goto L81
	}
L67:
	;
	goto L66
L68:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v206)+4))
	if v217 < v218 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v220 = v217
	goto L71
L70:
	;
	v220 = v218
	goto L71
L71:
	;
	if v220 <= int32(1) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v223 = int32(1)
	goto L74
L73:
	;
	v223 = v220
	goto L74
L74:
	;
	v224 = int32(8)
	v229 = int32(0)
	goto L75
L75:
	;
	v236 = v229 << (uint(int32(2)) % 32)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v206+v224+v236)))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l4+v224+v236)))
	v241 = v238 & v240
	v243 = base.B2i32(v241 != int32(0))
	if v241 != 0 {
		v252 = v243
		goto L67
	} else {
		goto L77
	}
L76:
	;
	v252 = v243
	goto L67
L77:
	;
	v245 = v229 + int32(1)
	if v245 != v223 {
		v229 = v245
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	if v252 != 0 {
		goto L92
	} else {
		goto L93
	}
L80:
	;
	goto L79
L81:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v253)+4))
	if v264 < v265 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v267 = v264
	goto L84
L83:
	;
	v267 = v265
	goto L84
L84:
	;
	if v267 <= int32(1) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v270 = int32(1)
	goto L87
L86:
	;
	v270 = v267
	goto L87
L87:
	;
	v271 = int32(8)
	v276 = int32(0)
	goto L88
L88:
	;
	v283 = v276 << (uint(int32(2)) % 32)
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v253+v271+v283)))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l4+v271+v283)))
	v288 = v285 & v287
	v290 = base.B2i32(v288 != int32(0))
	if v288 != 0 {
		v299 = v290
		goto L80
	} else {
		goto L90
	}
L89:
	;
	v299 = v290
	goto L80
L90:
	;
	v292 = v276 + int32(1)
	if v292 != v270 {
		v276 = v292
		goto L88
	} else {
		goto L91
	}
L91:
	;
	goto L89
L92:
	;
	if v299 != 0 {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	goto L94
L94:
	;
	if v299 != 0 {
		goto L34
	} else {
		goto L100
	}
L95:
	;
	v827 = int32(0)
	v830 = v104
	v832 = v105
	goto L30
L96:
	;
	goto L97
L97:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	v303 = F_palloc(m, int32(8))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L10
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v303)+4)) = v301
	*(*int32)(unsafe.Add(mBase, uint32(v303))) = v101
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v308 = F_lappend(m, v307, v303)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L10
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v308
	v350 = v104
	v351 = v105
	goto L32
L100:
	;
	v827 = int32(2)
	v830 = v104
	v832 = v105
	goto L30
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v106
	F_errmsg_internal(m, int32(_a_F_reduce_outer_joins_pass2_3), v20+int32(16))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L10
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(_a_F_reduce_outer_joins_pass2_1), int32(3610), int32(_a_F_reduce_outer_joins_pass2_2))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L10
	} else {
		goto L103
	}
L103:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v329)+4)) = v327
	*(*int32)(unsafe.Add(mBase, uint32(v329))) = v101
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v334 = F_lappend(m, v333, v329)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L10
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v334
	goto L33
L106:
	;
	v363 = int32(5)
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v366 = F_find_nonnullable_vars_walker(m, v364, int32(1))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L10
	} else {
		goto L107
	}
L107:
	;
	v370 = int32(0)
	v377 = v7
	goto L108
L108:
	;
	v386 = int32(0)
	if v366 == v386 {
		v396 = v386
		goto L110
	} else {
		goto L111
	}
L109:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v356)))
	v467 = int32(0)
	if base.B2i32(v464 == v467)|base.B2i32(v466 == v467) != 0 {
		v512 = v467
		goto L137
	} else {
		goto L138
	}
L110:
	;
	if l5 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L111:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v366)+4))
	if v390 <= v370 {
		v396 = int32(0)
		goto L110
	} else {
		goto L112
	}
L112:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v366)+12))
	v396 = v392 + v370<<(uint(int32(2))%32)
	goto L110
L113:
	;
	goto L109
L114:
	;
	v464 = int32(0)
	goto L113
L115:
	;
	goto L116
L116:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if base.B2i32(v396 == int32(0))|base.B2i32(v402 <= v370) != 0 {
		v464 = v377
		goto L113
	} else {
		goto L117
	}
L117:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	if v405 == int32(0) {
		v464 = v377
		goto L113
	} else {
		goto L118
	}
L118:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v396)))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v405+v370<<(uint(int32(2))%32))))
	v413 = int32(0)
	if base.B2i32(v408 == v413)|base.B2i32(v412 == v413) != 0 {
		v458 = v413
		goto L120
	} else {
		goto L121
	}
L119:
	;
	if v458 != 0 {
		goto L132
	} else {
		goto L133
	}
L120:
	;
	goto L119
L121:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v408)+4))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v412)+4))
	if v423 < v424 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v426 = v423
	goto L124
L123:
	;
	v426 = v424
	goto L124
L124:
	;
	if v426 <= int32(1) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v429 = int32(1)
	goto L127
L126:
	;
	v429 = v426
	goto L127
L127:
	;
	v430 = int32(8)
	v435 = int32(0)
	goto L128
L128:
	;
	v442 = v435 << (uint(int32(2)) % 32)
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v412+v430+v442)))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v408+v430+v442)))
	v447 = v444 & v446
	v449 = base.B2i32(v447 != int32(0))
	if v447 != 0 {
		v458 = v449
		goto L120
	} else {
		goto L130
	}
L129:
	;
	v458 = v449
	goto L120
L130:
	;
	v451 = v435 + int32(1)
	if v451 != v429 {
		v435 = v451
		goto L128
	} else {
		goto L131
	}
L131:
	;
	goto L129
L132:
	;
	v459 = F_bms_add_member(m, v377, v370)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L10
	} else {
		goto L135
	}
L133:
	;
	v461 = v377
	goto L134
L134:
	;
	v370 = v370 + int32(1)
	v377 = v461
	goto L108
L135:
	;
	v461 = v459
	goto L134
L136:
	;
	if v512 != 0 {
		v827 = v363
		v830 = v356
		v832 = v357
		goto L30
	} else {
		goto L149
	}
L137:
	;
	goto L136
L138:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v464)+4))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v466)+4))
	if v477 < v478 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v480 = v477
	goto L141
L140:
	;
	v480 = v478
	goto L141
L141:
	;
	if v480 <= int32(1) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v483 = int32(1)
	goto L144
L143:
	;
	v483 = v480
	goto L144
L144:
	;
	v484 = int32(8)
	v489 = int32(0)
	goto L145
L145:
	;
	v496 = v489 << (uint(int32(2)) % 32)
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v466+v484+v496)))
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v464+v484+v496)))
	v501 = v498 & v500
	v503 = base.B2i32(v501 != int32(0))
	if v501 != 0 {
		v512 = v503
		goto L137
	} else {
		goto L147
	}
L146:
	;
	v512 = v503
	goto L137
L147:
	;
	v505 = v489 + int32(1)
	if v505 != v483 {
		v489 = v505
		goto L145
	} else {
		goto L148
	}
L148:
	;
	goto L146
L149:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if int32(0) < v513 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v525 = int32(-1)
	v532 = v7
	goto L153
L151:
	;
	goto L152
L152:
	;
	v827 = int32(1)
	v830 = v356
	v832 = v357
	goto L30
L153:
	;
	v535 = v525 + int32(1)
	v536 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v536+v532<<(uint(int32(2))%32))))
	if v540 == int32(0) {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	goto L152
L155:
	;
	v800 = v532 + int32(1)
	v801 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v800 < v801 {
		v525 = v535
		v532 = v800
		goto L153
	} else {
		goto L214
	}
L156:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v356)))
	v544 = F_bms_is_member(m, v535, v543)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L10
	} else {
		goto L157
	}
L157:
	;
	if v544 == int32(0) {
		goto L155
	} else {
		goto L158
	}
L158:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v356)+8))
	v549 = F_bms_is_member(m, v535, v548)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L10
	} else {
		goto L159
	}
L159:
	;
	if v549 != 0 {
		goto L155
	} else {
		goto L160
	}
L160:
	;
	v551 = int32(0)
	if v540 == v551 {
		goto L163
	} else {
		goto L164
	}
L161:
	;
	if int32(0) <= v608 {
		goto L172
	} else {
		goto L173
	}
L162:
	;
	v608 = base.I32_ctz(v594) | v595<<(uint(int32(5))%32)
	goto L161
L163:
	;
	v608 = int32(-2)
	goto L161
L164:
	;
	v559 = int32(0)
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v540)+4))
	if v562 <= v559 {
		goto L163
	} else {
		goto L165
	}
L165:
	;
	v565 = v540 + int32(8)
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v565)))
	v572 = v569 & int32(-1)
	if v572 != 0 {
		v594 = v572
		v595 = v559
		goto L162
	} else {
		goto L166
	}
L166:
	;
	v573 = int32(1)
	if v573 == v562 {
		goto L163
	} else {
		goto L167
	}
L167:
	;
	v577 = v573
	goto L168
L168:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v565+v577<<(uint(int32(2))%32))))
	if v584 != 0 {
		v594 = v584
		v595 = v577
		goto L162
	} else {
		goto L170
	}
L169:
	;
	goto L163
L170:
	;
	v586 = v577 + int32(1)
	if v586 != v562 {
		v577 = v586
		goto L168
	} else {
		goto L171
	}
L171:
	;
	goto L169
L172:
	;
	v612 = v608
	v621 = v551
	goto L175
L173:
	;
	v706 = v551
	goto L174
L174:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v713)+52))
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v714)+12))
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v715+v525<<(uint(int32(2))%32))))
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v719)+12))
	if v720 != 0 {
		goto L191
	} else {
		goto L192
	}
L175:
	;
	v628 = int32(16)
	v633 = (v612<<(uint(v628)%32) - int32(_a_F_reduce_outer_joins_pass2_4)) >> (uint(v628) % 32)
	if v633 < int32(0) {
		v827 = v363
		v830 = v356
		v832 = v357
		goto L30
	} else {
		goto L177
	}
L176:
	;
	v706 = v636
	goto L174
L177:
	;
	v636 = F_bms_add_member(m, v621, v633)
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L10
	} else {
		goto L178
	}
L178:
	;
	if v540 == int32(0) {
		goto L181
	} else {
		goto L182
	}
L179:
	;
	if int32(0) <= v693 {
		v612 = v693
		v621 = v636
		goto L175
	} else {
		goto L190
	}
L180:
	;
	v693 = base.I32_ctz(v679) | v680<<(uint(int32(5))%32)
	goto L179
L181:
	;
	v693 = int32(-2)
	goto L179
L182:
	;
	v644 = v612 + int32(1)
	v646 = int32(base.Ui32(v644) >> (uint(int32(5)) % 32))
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v540)+4))
	if v647 <= v646 {
		goto L181
	} else {
		goto L183
	}
L183:
	;
	v650 = v540 + int32(8)
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v650+v646<<(uint(int32(2))%32))))
	v657 = v654 & (int32(-1) << (uint(v644) % 32))
	if v657 != 0 {
		v679 = v657
		v680 = v646
		goto L180
	} else {
		goto L184
	}
L184:
	;
	v659 = v646 + int32(1)
	if v659 == v647 {
		goto L181
	} else {
		goto L185
	}
L185:
	;
	v662 = v659
	goto L186
L186:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v650+v662<<(uint(int32(2))%32))))
	if v669 != 0 {
		v679 = v669
		v680 = v662
		goto L180
	} else {
		goto L188
	}
L187:
	;
	goto L181
L188:
	;
	v671 = v662 + int32(1)
	if v671 != v647 {
		v662 = v671
		goto L186
	} else {
		goto L189
	}
L189:
	;
	goto L187
L190:
	;
	goto L176
L191:
	;
	F_bms_free(m, v706)
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L10
	} else {
		goto L213
	}
L192:
	;
	v721 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v719)+20)))
	if v721 == int32(1) {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v724 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v719)+21)))
	if v724 != int32(112) {
		goto L191
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v719)+16))
	v728 = F_find_relation_notnullatts(m, l3, v727)
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L10
	} else {
		goto L197
	}
L196:
	;
	goto L195
L197:
	;
	v730 = int32(0)
	if base.B2i32(v728 == v730)|base.B2i32(v706 == v730) != 0 {
		v775 = v730
		goto L199
	} else {
		goto L200
	}
L198:
	;
	F_bms_free(m, v706)
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L10
	} else {
		goto L211
	}
L199:
	;
	goto L198
L200:
	;
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v728)+4))
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v706)+4))
	if v740 < v741 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v743 = v740
	goto L203
L202:
	;
	v743 = v741
	goto L203
L203:
	;
	if v743 <= int32(1) {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v746 = int32(1)
	goto L206
L205:
	;
	v746 = v743
	goto L206
L206:
	;
	v747 = int32(8)
	v752 = int32(0)
	goto L207
L207:
	;
	v759 = v752 << (uint(int32(2)) % 32)
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v706+v747+v759)))
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v728+v747+v759)))
	v764 = v761 & v763
	v766 = base.B2i32(v764 != int32(0))
	if v764 != 0 {
		v775 = v766
		goto L199
	} else {
		goto L209
	}
L208:
	;
	v775 = v766
	goto L199
L209:
	;
	v768 = v752 + int32(1)
	if v768 != v746 {
		v752 = v768
		goto L207
	} else {
		goto L210
	}
L210:
	;
	goto L208
L211:
	;
	if v775 == int32(0) {
		goto L155
	} else {
		goto L212
	}
L212:
	;
	v827 = v363
	v830 = v356
	v832 = v357
	goto L30
L213:
	;
	goto L155
L214:
	;
	goto L154
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v827
	v861 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v832)+4)))
	if v861 == int32(0) {
		goto L222
	} else {
		goto L223
	}
L216:
	;
	v840 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v827 == v840 {
		goto L215
	} else {
		goto L217
	}
L217:
	;
	v842 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v842)+52))
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v843)+12))
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v844+v101<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v850)+44)) = v827
	switch v827 {
	case 0:
		goto L219
	default:
		goto L215
	case 5:
		goto L218
	}
L218:
	;
	v856 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v857 = F_bms_add_member(m, v856, v101)
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L10
	} else {
		goto L221
	}
L219:
	;
	v852 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v853 = F_bms_add_member(m, v852, v101)
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L10
	} else {
		goto L220
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v853
	goto L215
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v857
	goto L215
L222:
	;
	v864 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v830)+4)))
	if v864 != int32(1) {
		goto L5
	} else {
		goto L225
	}
L223:
	;
	goto L224
L224:
	;
	v867 = int32(0)
	if v827 == int32(2) {
		v884 = v867
		v885 = v867
		goto L226
	} else {
		goto L227
	}
L225:
	;
	goto L224
L226:
	;
	v886 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v832)+4)))
	if v886 == int32(1) {
		goto L233
	} else {
		goto L234
	}
L227:
	;
	v871 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v872 = F_find_nonnullable_rels(m, v871)
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L10
	} else {
		goto L228
	}
L228:
	;
	v874 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v875 = F_find_forced_null_vars(m, v874)
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L10
	} else {
		goto L229
	}
L229:
	;
	if v827&int32(-5) != 0 {
		v884 = v872
		v885 = v875
		goto L226
	} else {
		goto L230
	}
L230:
	;
	v879 = F_bms_add_members(m, v872, l4)
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L10
	} else {
		goto L231
	}
L231:
	;
	v881 = F_mbms_add_members(m, v875, l5)
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L10
	} else {
		goto L232
	}
L232:
	;
	v884 = v879
	v885 = v881
	goto L226
L233:
	;
	v889 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v892 = base.B2i32(v827 == int32(2))
	if v827 == int32(2) {
		goto L236
	} else {
		goto L237
	}
L234:
	;
	goto L235
L235:
	;
	v904 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v830)+4)))
	if v904 != int32(1) {
		v918 = v884
		goto L6
	} else {
		goto L249
	}
L236:
	;
	v893 = int32(0)
	goto L238
L237:
	;
	v893 = l4
	goto L238
L238:
	;
	v895 = v827 & int32(-5)
	if v895 != 0 {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v896 = v893
	goto L241
L240:
	;
	v896 = v884
	goto L241
L241:
	;
	if v827 == int32(2) {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	v898 = int32(0)
	goto L244
L243:
	;
	v898 = l5
	goto L244
L244:
	;
	if v895 != 0 {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v899 = v898
	goto L247
L246:
	;
	v899 = v885
	goto L247
L247:
	;
	F_reduce_outer_joins_pass2(m, v889, v832, l2, l3, v896, v899)
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L10
	} else {
		goto L248
	}
L248:
	;
	goto L235
L249:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	F_reduce_outer_joins_pass2(m, v907, v830, l2, l3, v884, v885)
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L10
	} else {
		goto L250
	}
L250:
	;
	v918 = v884
	goto L6
L251:
	;
	goto L5
L252:
	;
	F_errmsg_internal(m, int32(_a_F_reduce_outer_joins_pass2_5), int32(0))
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L10
	} else {
		goto L253
	}
L253:
	;
	F_errfinish(m, int32(_a_F_reduce_outer_joins_pass2_1), int32(3520), int32(_a_F_reduce_outer_joins_pass2_2))
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L10
	} else {
		goto L254
	}
L254:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L255:
	;
	v966 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v966
	F_errmsg_internal(m, int32(_a_F_reduce_outer_joins_pass2_6), v20)
	mBase = m.M
	v970 = m.ExcPending
	if v970 != 0 {
		goto L10
	} else {
		goto L256
	}
L256:
	;
	F_errfinish(m, int32(_a_F_reduce_outer_joins_pass2_1), int32(3786), int32(_a_F_reduce_outer_joins_pass2_2))
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L10
	} else {
		goto L257
	}
L257:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_regcollationin(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
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
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int64
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v158 int64
	_ = v158
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	v6 = int64(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v13 == int32(45) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L29
	} else {
		goto L47
	}
L2:
	;
	m.G0 = v9 + int32(16)
	return v158
L3:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_regcollationin[0]))
	if v119 == int32(0) {
		goto L1
	} else {
		goto L31
	}
L4:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
	if v16 != 0 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if base.Ui32(int32(9)) < base.Ui32((v13-int32(48))&int32(255)) {
		goto L3
	} else {
		goto L8
	}
L7:
	;
	v158 = v6
	goto L2
L8:
	;
	v23 = int32(_a_F_regcollationin_0)
	v27 = m.G0
	v29 = v27 - int32(32)
	v30 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+24)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29)+16)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29))) = v30
	v38 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_regcollationin[1])))
	if v38 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v107 = F_strlen(m, v12)
	mBase = m.M
	if v106 != v107 {
		goto L3
	} else {
		goto L28
	}
L10:
	;
	v106 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_regcollationin[2])))
	if v42 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v46 = v12
	goto L16
L14:
	;
	goto L15
L15:
	;
	v56 = v23
	v57 = v38
	goto L19
L16:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v52 == v38 {
		v46 = v46 + int32(1)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v106 = v46 - v12
	goto L9
L18:
	;
	goto L17
L19:
	;
	v64 = v29 + int32(base.Ui32(v57)>>(uint(int32(3))%32))&int32(28)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v66 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v65 | v66<<(uint(v57)%32)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	if v70 != 0 {
		v56 = v56 + v66
		v57 = v70
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v73 == int32(0) {
		v96 = v12
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v106 = v96 - v12
	goto L9
L23:
	;
	v77 = v12
	v78 = v73
	goto L24
L24:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(base.Ui32(v78)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v86)>>(uint(v78)%32))&int32(1) == int32(0) {
		v96 = v77
		goto L22
	} else {
		goto L26
	}
L25:
	;
	v96 = v94
	goto L22
L26:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+1)))
	v94 = v77 + int32(1)
	if v92 != 0 {
		v77 = v94
		v78 = v92
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v113 = F_DirectInputFunctionCallSafe(m, int32(588), v12, int32(-1), v11, v9+int32(8))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	return int64(0)
L30:
	;
	v117 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v9)+8)))
	v158 = v117
	goto L2
L31:
	;
	v122 = F_stringToQualifiedNameList(m, v12, v11)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	if v122 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v126 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v126)
	v158 = v6
	goto L2
L34:
	;
	goto L35
L35:
	;
	v129 = F_get_collation_oid(m, v122, int32(1))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L29
	} else {
		goto L36
	}
L36:
	;
	if v129 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v133 = F_errsave_start(m, v11)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L29
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v158 = base.I64_extend_i32_u(v129)
	goto L2
L40:
	;
	if v133 == int32(0) {
		v158 = v6
		goto L2
	} else {
		goto L41
	}
L41:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L29
	} else {
		goto L42
	}
L42:
	;
	v140 = F_NameListToString(m, v122)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L29
	} else {
		goto L43
	}
L43:
	;
	v143 = *(*int32)(unsafe.Add(mBase, _c_F_regcollationin[3]))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v140
	F_errmsg(m, int32(_a_F_regcollationin_1), v9)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L29
	} else {
		goto L45
	}
L45:
	;
	F_errsave_finish(m, v11, int32(_a_F_regcollationin_2), int32(1065), int32(_a_F_regcollationin_3))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L29
	} else {
		goto L46
	}
L46:
	;
	v158 = v6
	goto L2
L47:
	;
	F_errmsg_internal(m, int32(_a_F_regcollationin_4), int32(0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L29
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_regcollationin_2), int32(1049), int32(_a_F_regcollationin_3))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L29
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_regconfigout(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
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
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = base.I32_wrap_i64(v11)
	if v12 == int32(0) {
		v16 = F_pstrdup(m, int32(_a_F_regconfigout_0))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			v49 = v16
			m.G0 = v9 + int32(16)
			return base.I64_extend_i32_u(v49)
		}
	} else {
		v23 = F_SearchSysCache1(m, int32(74), v11&int64(4294967295))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int64(0)
		} else {
			if v23 != 0 {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+22)))
				v27 = v25 + v26
				v30 = F_TSConfigIsVisible(m, v12)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int64(0)
				} else {
					if v30 != 0 {
						v36 = int32(0)
						v37 = F_quote_qualified_identifier(m, v36, v27+int32(4))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int64(0)
						} else {
							F_ReleaseCatCache(m, v23)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int64(0)
							} else {
								v49 = v37
								m.G0 = v9 + int32(16)
								return base.I64_extend_i32_u(v49)
							}
						}
					} else {
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v27)+68))
						v34 = F_get_namespace_name(m, v33)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int64(0)
						} else {
							v36 = v34
							v37 = F_quote_qualified_identifier(m, v36, v27+int32(4))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int64(0)
							} else {
								F_ReleaseCatCache(m, v23)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int64(0)
								} else {
									v49 = v37
									m.G0 = v9 + int32(16)
									return base.I64_extend_i32_u(v49)
								}
							}
						}
					}
				}
			} else {
				v42 = F_palloc(m, int32(64))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v12
					v47 = F_pg_snprintf(m, v42, int32(64), int32(_a_F_regconfigout_1), v9)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int64(0)
					} else {
						v49 = v42
						m.G0 = v9 + int32(16)
						return base.I64_extend_i32_u(v49)
					}
				}
			}
		}
	}
}
func F_regnamespacein(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
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
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int64
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
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
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v177 int64
	_ = v177
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	v6 = int64(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v13 == int32(45) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L29
	} else {
		goto L53
	}
L2:
	;
	m.G0 = v9 + int32(16)
	return v177
L3:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_regnamespacein[0]))
	if v119 == int32(0) {
		goto L1
	} else {
		goto L31
	}
L4:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
	if v16 != 0 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if base.Ui32(int32(9)) < base.Ui32((v13-int32(48))&int32(255)) {
		goto L3
	} else {
		goto L8
	}
L7:
	;
	v177 = v6
	goto L2
L8:
	;
	v23 = int32(_a_F_regnamespacein_0)
	v27 = m.G0
	v29 = v27 - int32(32)
	v30 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+24)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29)+16)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29))) = v30
	v38 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_regnamespacein[1])))
	if v38 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v107 = F_strlen(m, v12)
	mBase = m.M
	if v106 != v107 {
		goto L3
	} else {
		goto L28
	}
L10:
	;
	v106 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_regnamespacein[2])))
	if v42 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v46 = v12
	goto L16
L14:
	;
	goto L15
L15:
	;
	v56 = v23
	v57 = v38
	goto L19
L16:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v52 == v38 {
		v46 = v46 + int32(1)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v106 = v46 - v12
	goto L9
L18:
	;
	goto L17
L19:
	;
	v64 = v29 + int32(base.Ui32(v57)>>(uint(int32(3))%32))&int32(28)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v66 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v65 | v66<<(uint(v57)%32)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	if v70 != 0 {
		v56 = v56 + v66
		v57 = v70
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v73 == int32(0) {
		v96 = v12
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v106 = v96 - v12
	goto L9
L23:
	;
	v77 = v12
	v78 = v73
	goto L24
L24:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(base.Ui32(v78)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v86)>>(uint(v78)%32))&int32(1) == int32(0) {
		v96 = v77
		goto L22
	} else {
		goto L26
	}
L25:
	;
	v96 = v94
	goto L22
L26:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+1)))
	v94 = v77 + int32(1)
	if v92 != 0 {
		v77 = v94
		v78 = v92
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v113 = F_DirectInputFunctionCallSafe(m, int32(588), v12, int32(-1), v11, v9+int32(8))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	return int64(0)
L30:
	;
	v117 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v9)+8)))
	v177 = v117
	goto L2
L31:
	;
	v122 = F_stringToQualifiedNameList(m, v12, v11)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	if v122 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v126 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v126)
	v177 = v6
	goto L2
L34:
	;
	goto L35
L35:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	if v128 != int32(1) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v131 = F_errsave_start(m, v11)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L29
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	v151 = F_get_namespace_oid(m, v149, int32(1))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L29
	} else {
		goto L44
	}
L39:
	;
	if v131 == int32(0) {
		v177 = v6
		goto L2
	} else {
		goto L40
	}
L40:
	;
	F_errcode(m, int32(33579140))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L29
	} else {
		goto L41
	}
L41:
	;
	F_errmsg(m, int32(_a_F_regnamespacein_1), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L29
	} else {
		goto L42
	}
L42:
	;
	F_errsave_finish(m, v11, int32(_a_F_regnamespacein_2), int32(1689), int32(_a_F_regnamespacein_3))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L29
	} else {
		goto L43
	}
L43:
	;
	v177 = v6
	goto L2
L44:
	;
	if v151 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v155 = F_errsave_start(m, v11)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L29
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v177 = base.I64_extend_i32_u(v151)
	goto L2
L48:
	;
	if v155 == int32(0) {
		v177 = v6
		goto L2
	} else {
		goto L49
	}
L49:
	;
	F_errcode(m, int32(1411))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L29
	} else {
		goto L50
	}
L50:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v164
	F_errmsg(m, int32(_a_F_regnamespacein_4), v9)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L29
	} else {
		goto L51
	}
L51:
	;
	F_errsave_finish(m, v11, int32(_a_F_regnamespacein_2), int32(1697), int32(_a_F_regnamespacein_3))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L29
	} else {
		goto L52
	}
L52:
	;
	v177 = v6
	goto L2
L53:
	;
	F_errmsg_internal(m, int32(_a_F_regnamespacein_5), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L29
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_regnamespacein_2), int32(1679), int32(_a_F_regnamespacein_3))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L29
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_regprocedurein(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
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
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int64
	_ = v117
	var v119 int32
	_ = v119
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
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
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v245 int64
	_ = v245
	var v251 int64
	_ = v251
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	v6 = int64(0)
	v7 = m.G0
	v9 = v7 - int32(432)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v13 == int32(45) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L29
	} else {
		goto L68
	}
L2:
	;
	m.G0 = v9 + int32(432)
	return v251
L3:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_regprocedurein[0]))
	if v119 == int32(0) {
		goto L1
	} else {
		goto L31
	}
L4:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
	if v16 != 0 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if base.Ui32(int32(9)) < base.Ui32((v13-int32(48))&int32(255)) {
		goto L3
	} else {
		goto L8
	}
L7:
	;
	v251 = v6
	goto L2
L8:
	;
	v23 = int32(_a_F_regprocedurein_0)
	v27 = m.G0
	v29 = v27 - int32(32)
	v30 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+24)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29)+16)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29))) = v30
	v38 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_regprocedurein[1])))
	if v38 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v107 = F_strlen(m, v12)
	mBase = m.M
	if v106 != v107 {
		goto L3
	} else {
		goto L28
	}
L10:
	;
	v106 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_regprocedurein[2])))
	if v42 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v46 = v12
	goto L16
L14:
	;
	goto L15
L15:
	;
	v56 = v23
	v57 = v38
	goto L19
L16:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v52 == v38 {
		v46 = v46 + int32(1)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v106 = v46 - v12
	goto L9
L18:
	;
	goto L17
L19:
	;
	v64 = v29 + int32(base.Ui32(v57)>>(uint(int32(3))%32))&int32(28)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v66 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v65 | v66<<(uint(v57)%32)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	if v70 != 0 {
		v56 = v56 + v66
		v57 = v70
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v73 == int32(0) {
		v96 = v12
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v106 = v96 - v12
	goto L9
L23:
	;
	v77 = v12
	v78 = v73
	goto L24
L24:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(base.Ui32(v78)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v86)>>(uint(v78)%32))&int32(1) == int32(0) {
		v96 = v77
		goto L22
	} else {
		goto L26
	}
L25:
	;
	v96 = v94
	goto L22
L26:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+1)))
	v94 = v77 + int32(1)
	if v92 != 0 {
		v77 = v94
		v78 = v92
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v113 = F_DirectInputFunctionCallSafe(m, int32(588), v12, int32(-1), v11, v9+int32(16))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	return int64(0)
L30:
	;
	v117 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v9)+16)))
	v251 = v117
	goto L2
L31:
	;
	v129 = F_parseNameAndArgTypes(m, v12, int32(0), v9+int32(428), v9+int32(424), v9+int32(16), v11)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	if v129 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v133 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v133)
	v251 = v6
	goto L2
L34:
	;
	goto L35
L35:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v9)+428))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v9)+424))
	v137 = int32(0)
	v144 = F_FuncnameGetCandidates(m, v135, v136, v137, v137, v137, v137, int32(1), v9+int32(12))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L29
	} else {
		goto L37
	}
L36:
	;
	v245 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v148)+8)))
	v251 = v245
	goto L2
L37:
	;
	if v144 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v147 = v136 << (uint(int32(2)) % 32)
	v148 = v144
	goto L41
L39:
	;
	goto L40
L40:
	;
	v229 = F_errsave_start(m, v11)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L29
	} else {
		goto L63
	}
L41:
	;
	v155 = v148 + int32(32)
	v157 = v9 + int32(16)
	if base.Ui32(int32(4)) <= base.Ui32(v147) {
		goto L46
	} else {
		goto L47
	}
L42:
	;
	goto L40
L43:
	;
	if v219 == int32(0) {
		goto L36
	} else {
		goto L61
	}
L44:
	;
	v219 = int32(0)
	goto L43
L45:
	;
	v193 = v188
	v194 = v189
	v195 = v190
	goto L55
L46:
	;
	if (v155|v157)&int32(3) != 0 {
		v188 = v155
		v189 = v157
		v190 = v147
		goto L45
	} else {
		goto L49
	}
L47:
	;
	v181 = v155
	v182 = v157
	v183 = v147
	goto L48
L48:
	;
	if v183 == int32(0) {
		goto L44
	} else {
		goto L54
	}
L49:
	;
	v165 = v155
	v166 = v157
	v167 = v147
	goto L50
L50:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
	if v170 != v171 {
		v188 = v165
		v189 = v166
		v190 = v167
		goto L45
	} else {
		goto L52
	}
L51:
	;
	v181 = v176
	v182 = v174
	v183 = v178
	goto L48
L52:
	;
	v173 = int32(4)
	v174 = v166 + v173
	v176 = v165 + v173
	v178 = v167 - v173
	if base.Ui32(int32(3)) < base.Ui32(v178) {
		v165 = v176
		v166 = v174
		v167 = v178
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v188 = v181
	v189 = v182
	v190 = v183
	goto L45
L55:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193))))
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194))))
	if v198 == v199 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v219 = v198 - v199
	goto L43
L57:
	;
	v201 = int32(1)
	v206 = v195 - v201
	if v206 != 0 {
		v193 = v193 + v201
		v194 = v194 + v201
		v195 = v206
		goto L55
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	goto L56
L60:
	;
	goto L44
L61:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	if v222 != 0 {
		v148 = v222
		goto L41
	} else {
		goto L62
	}
L62:
	;
	goto L42
L63:
	;
	if v229 == int32(0) {
		v251 = v6
		goto L2
	} else {
		goto L64
	}
L64:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L29
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v12
	F_errmsg(m, int32(_a_F_regprocedurein_1), v9)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L29
	} else {
		goto L66
	}
L66:
	;
	F_errsave_finish(m, v11, int32(_a_F_regprocedurein_2), int32(271), int32(_a_F_regprocedurein_3))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L29
	} else {
		goto L67
	}
L67:
	;
	v251 = v6
	goto L2
L68:
	;
	F_errmsg_internal(m, int32(_a_F_regprocedurein_4), int32(0))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L29
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_regprocedurein_2), int32(246), int32(_a_F_regprocedurein_3))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L29
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_removeabbrev_datum(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v29 int64
	_ = v29
	var v31 int64
	_ = v31
	var v33 int64
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	v4 = int32(0)
	if l2 <= v4 {
	} else {
		v11 = l2 & int32(3)
		v12 = int32(0)
		if base.Ui32(int32(4)) <= base.Ui32(l2) {
			v17 = v12
			v22 = v4
			for {
				v26 = l1 + v17*int32(24)
				v27 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v26))))
				*(*int64)(unsafe.Add(mBase, uint32(v26)+8)) = v27
				v29 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v26)+24)))
				*(*int64)(unsafe.Add(mBase, uint32(v26)+32)) = v29
				v31 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v26)+48)))
				*(*int64)(unsafe.Add(mBase, uint32(v26)+56)) = v31
				v33 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v26)+72)))
				*(*int64)(unsafe.Add(mBase, uint32(v26)+80)) = v33
				v35 = int32(4)
				v36 = v17 + v35
				v38 = v22 + v35
				if v38 != l2&int32(2147483644) {
					v17 = v36
					v22 = v38
					continue
				} else {
					break
				}
				break
			}
			if v11 == int32(0) {
			} else {
				v42 = v36
				v49 = v42
				v55 = v4
				for {
					v58 = l1 + v49*int32(24)
					v59 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v58))))
					*(*int64)(unsafe.Add(mBase, uint32(v58)+8)) = v59
					v61 = int32(1)
					v64 = v55 + v61
					if v64 != v11 {
						v49 = v49 + v61
						v55 = v64
						continue
					} else {
						break
					}
					break
				}
			}
		} else {
			v42 = v12
			v49 = v42
			v55 = v4
			for {
				v58 = l1 + v49*int32(24)
				v59 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v58))))
				*(*int64)(unsafe.Add(mBase, uint32(v58)+8)) = v59
				v61 = int32(1)
				v64 = v55 + v61
				if v64 != v11 {
					v49 = v49 + v61
					v55 = v64
					continue
				} else {
					break
				}
				break
			}
		}
	}
	return
}
func F_removetraverse(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
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
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v94 int32
	_ = v94
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v167 int64
	_ = v167
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v11 = m.T0[v10].(func(*base.Module) int32)(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v11 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = int32(101)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	if v17 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v21 != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v19 = v17
	goto L8
L7:
	;
	v19 = int32(19)
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v19
	return
L9:
	;
	return
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = l1
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v23 == int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v31 = v23
	goto L12
L12:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	if v33 != 0 {
		goto L9
	} else {
		goto L14
	}
L13:
	;
	goto L9
L14:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	F_removetraverse(m, l0, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	if v38 != 0 {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	switch v40 - int32(76) {
	case 0, 18, 21, 38:
		goto L19
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 19, 20, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 35, 37, 39, 40, 41, 42, 43:
		goto L18
	case 34, 36, 44:
		goto L17
	default:
		goto L20
	}
L17:
	;
	if v39 != 0 {
		v31 = v39
		goto L12
	} else {
		goto L77
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+24)) = int32(101)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+12))
	if v179 != 0 {
		goto L74
	} else {
		goto L75
	}
L19:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_removetraverse[0]))
	if v47 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	if v40 != int32(36) {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	if v50 <= v51 {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	goto L24
L26:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	v107 = int32(*(*int16)(unsafe.Add(mBase, uint32(v31)+4)))
	if v107 < int32(0) {
		goto L49
	} else {
		goto L50
	}
L27:
	;
	F_createarc(m, l0, int32(110), int32(0), l1, v45)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L47
	}
L28:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v53 == int32(0) {
		goto L27
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v45)+16))
	if v69 == int32(0) {
		goto L27
	} else {
		goto L39
	}
L31:
	;
	v58 = v53
	goto L32
L32:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
	if v62 != v45 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L27
L34:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
	if v68 != 0 {
		v58 = v68
		goto L32
	} else {
		goto L38
	}
L35:
	;
	v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+4)))
	if v64 != 0 {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	if v65 == int32(110) {
		goto L26
	} else {
		goto L37
	}
L37:
	;
	goto L34
L38:
	;
	goto L33
L39:
	;
	v74 = v69
	goto L40
L40:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
	if v78 != l1 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L27
L42:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v74)+24))
	if v84 != 0 {
		v74 = v84
		goto L40
	} else {
		goto L46
	}
L43:
	;
	v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v74)+4)))
	if v80 != 0 {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	if v81 == int32(110) {
		goto L26
	} else {
		goto L45
	}
L45:
	;
	goto L42
L46:
	;
	goto L41
L47:
	;
	goto L26
L48:
	;
	goto L17
L49:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
	if v142 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L50:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v112 = v110 - int32(97)
	if base.B2i32(base.Ui32(int32(17)) < base.Ui32(v112))|base.B2i32(int32(1)<<(uint(v112)%32)&int32(_a_F_removetraverse_0) == int32(0)) != 0 {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v122 != 0 {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v31)+36))
	if v123 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	if v135 != 0 {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+20))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v31)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v127+v107*int32(24))+12)) = v131
	v135 = v131
	goto L53
L55:
	;
	goto L56
L56:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v31)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v123)+32)) = v133
	v135 = v133
	goto L53
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135)+36)) = v123
	goto L59
L58:
	;
	goto L59
L59:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v31)+32)) = int64(0)
	goto L49
L60:
	;
	if v141 != 0 {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v106)+20)) = v141
	goto L60
L62:
	;
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v142)+16)) = v141
	goto L60
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v141)+20)) = v142
	goto L66
L65:
	;
	goto L66
L66:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v106)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v106)+12)) = v148 - int32(1)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v31)+24))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v31)+28))
	if v153 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	if v152 != 0 {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105)+16)) = v152
	goto L67
L69:
	;
	goto L70
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153)+24)) = v152
	goto L67
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v152)+28)) = v153
	goto L73
L72:
	;
	goto L73
L73:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v105)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v105)+8)) = v159 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = int32(0)
	v166 = v31 + int32(8)
	v167 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v166)+16)) = v167
	*(*int64)(unsafe.Add(mBase, uint32(v166)+8)) = v167
	*(*int64)(unsafe.Add(mBase, uint32(v166))) = v167
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v31
	goto L48
L74:
	;
	v181 = v179
	goto L76
L75:
	;
	v181 = int32(15)
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v178)+12)) = v181
	goto L17
L77:
	;
	goto L13
}
func F_renameatt_check(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+119)))
	if l2 == int32(0) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
		if v13 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v83 = m.ExcPending
			if v83 != 0 {
				return
			} else {
				F_errcode(m, int32(151027844))
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_renameatt_check_0), int32(0))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_renameatt_check_1), int32(3835), int32(_a_F_renameatt_check_2))
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
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
			switch v10 - int32(73) {
			case 0, 26, 29, 32, 36, 39, 41, 45:
				v39 = *(*int32)(unsafe.Add(mBase, _c_F_renameatt_check[0]))
				v40 = F_object_ownercheck(m, int32(1259), l0, v39)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					if v40 == int32(0) {
						v45 = F_get_rel_relkind(m, l0)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							switch v45 - int32(73) {
							case 0, 32:
								v56 = int32(20)
								v58 = v56
							default:
								v56 = int32(42)
								v58 = v56
							case 10:
								v58 = int32(38)
							case 29:
								v58 = int32(18)
							case 36:
								v58 = int32(23)
							case 45:
								v58 = int32(52)
							}
							F_aclcheck_error(m, int32(2), v58, l1+int32(4))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return
							} else {
								v64 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_renameatt_check[1])))
								if v64 == int32(0) {
									v68 = int32(1)
									if base.Ui32(l0) < base.Ui32(int32(_a_F_renameatt_check_3)) {
										v76 = v68
									} else {
										v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
										if v71 == int32(99) {
											v76 = v68
										} else {
											v74 = F_isTempToastNamespace(m, v71)
											mBase = m.M
											v76 = v74
										}
									}
									if v76 != 0 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return
										} else {
											F_errcode(m, int32(16797828))
											mBase = m.M
											v102 = m.ExcPending
											if v102 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l1 + int32(4)
												F_errmsg(m, int32(_a_F_renameatt_check_4), v8+int32(16))
												mBase = m.M
												v110 = m.ExcPending
												if v110 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_renameatt_check_1), int32(3868), int32(_a_F_renameatt_check_2))
													mBase = m.M
													v115 = m.ExcPending
													if v115 != 0 {
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
										m.G0 = v8 + int32(32)
										return
									}
								} else {
									m.G0 = v8 + int32(32)
									return
								}
							}
						}
					} else {
						v64 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_renameatt_check[1])))
						if v64 == int32(0) {
							v68 = int32(1)
							if base.Ui32(l0) < base.Ui32(int32(_a_F_renameatt_check_3)) {
								v76 = v68
							} else {
								v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
								if v71 == int32(99) {
									v76 = v68
								} else {
									v74 = F_isTempToastNamespace(m, v71)
									mBase = m.M
									v76 = v74
								}
							}
							if v76 != 0 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return
								} else {
									F_errcode(m, int32(16797828))
									mBase = m.M
									v102 = m.ExcPending
									if v102 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l1 + int32(4)
										F_errmsg(m, int32(_a_F_renameatt_check_4), v8+int32(16))
										mBase = m.M
										v110 = m.ExcPending
										if v110 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_renameatt_check_1), int32(3868), int32(_a_F_renameatt_check_2))
											mBase = m.M
											v115 = m.ExcPending
											if v115 != 0 {
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
								m.G0 = v8 + int32(32)
								return
							}
						} else {
							m.G0 = v8 + int32(32)
							return
						}
					}
				}
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					F_errcode(m, int32(151027844))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1 + int32(4)
						F_errmsg(m, int32(_a_F_renameatt_check_5), v8)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return
						} else {
							F_errdetail_relkind_not_supported(m, base.I32_extend8_s(v10))
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_renameatt_check_1), int32(3856), int32(_a_F_renameatt_check_2))
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
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
		switch v10 - int32(73) {
		case 0, 26, 29, 32, 36, 39, 41, 45:
			v39 = *(*int32)(unsafe.Add(mBase, _c_F_renameatt_check[0]))
			v40 = F_object_ownercheck(m, int32(1259), l0, v39)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				if v40 == int32(0) {
					v45 = F_get_rel_relkind(m, l0)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						switch v45 - int32(73) {
						case 0, 32:
							v56 = int32(20)
							v58 = v56
						default:
							v56 = int32(42)
							v58 = v56
						case 10:
							v58 = int32(38)
						case 29:
							v58 = int32(18)
						case 36:
							v58 = int32(23)
						case 45:
							v58 = int32(52)
						}
						F_aclcheck_error(m, int32(2), v58, l1+int32(4))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return
						} else {
							v64 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_renameatt_check[1])))
							if v64 == int32(0) {
								v68 = int32(1)
								if base.Ui32(l0) < base.Ui32(int32(_a_F_renameatt_check_3)) {
									v76 = v68
								} else {
									v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
									if v71 == int32(99) {
										v76 = v68
									} else {
										v74 = F_isTempToastNamespace(m, v71)
										mBase = m.M
										v76 = v74
									}
								}
								if v76 != 0 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v99 = m.ExcPending
									if v99 != 0 {
										return
									} else {
										F_errcode(m, int32(16797828))
										mBase = m.M
										v102 = m.ExcPending
										if v102 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l1 + int32(4)
											F_errmsg(m, int32(_a_F_renameatt_check_4), v8+int32(16))
											mBase = m.M
											v110 = m.ExcPending
											if v110 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_renameatt_check_1), int32(3868), int32(_a_F_renameatt_check_2))
												mBase = m.M
												v115 = m.ExcPending
												if v115 != 0 {
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
									m.G0 = v8 + int32(32)
									return
								}
							} else {
								m.G0 = v8 + int32(32)
								return
							}
						}
					}
				} else {
					v64 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_renameatt_check[1])))
					if v64 == int32(0) {
						v68 = int32(1)
						if base.Ui32(l0) < base.Ui32(int32(_a_F_renameatt_check_3)) {
							v76 = v68
						} else {
							v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
							if v71 == int32(99) {
								v76 = v68
							} else {
								v74 = F_isTempToastNamespace(m, v71)
								mBase = m.M
								v76 = v74
							}
						}
						if v76 != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return
							} else {
								F_errcode(m, int32(16797828))
								mBase = m.M
								v102 = m.ExcPending
								if v102 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l1 + int32(4)
									F_errmsg(m, int32(_a_F_renameatt_check_4), v8+int32(16))
									mBase = m.M
									v110 = m.ExcPending
									if v110 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_renameatt_check_1), int32(3868), int32(_a_F_renameatt_check_2))
										mBase = m.M
										v115 = m.ExcPending
										if v115 != 0 {
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
							m.G0 = v8 + int32(32)
							return
						}
					} else {
						m.G0 = v8 + int32(32)
						return
					}
				}
			}
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				F_errcode(m, int32(151027844))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1 + int32(4)
					F_errmsg(m, int32(_a_F_renameatt_check_5), v8)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						F_errdetail_relkind_not_supported(m, base.I32_extend8_s(v10))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_renameatt_check_1), int32(3856), int32(_a_F_renameatt_check_2))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
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
func F_replace_s(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	v12 = l1 - l2 + l3
	if v12 == int32(0) {
		if l3 != 0 {
			v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			base.MemoryCopy(m, v64+l1, l4, l3)
		} else {
		}
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v15-int32(4))))
		v19 = v18 + v12
		v21 = v15 - int32(8)
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
		if v22 < v19 {
			v26 = F_repalloc(m, v21, v19+int32(29))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				if v26 == int32(0) {
					return int32(-1)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v26))) = v19 + int32(20)
					v38 = v26 + int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v38
					v40 = v38
					v41 = v18 - l2
					if v41 != 0 {
						v42 = l2 + v40
						base.MemoryCopy(m, v42+v12, v42, v41)
					} else {
					}
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v46-int32(4)))) = v19
					v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v50 + v12
					v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					if l2 <= v53 {
						v57 = v53 + v12
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v57
					} else {
						if v53 <= l1 {
						} else {
							v57 = l1
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v57
						}
					}
					if l3 != 0 {
						v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						base.MemoryCopy(m, v64+l1, l4, l3)
					} else {
					}
					return int32(0)
				}
			}
		} else {
			v40 = v15
			v41 = v18 - l2
			if v41 != 0 {
				v42 = l2 + v40
				base.MemoryCopy(m, v42+v12, v42, v41)
			} else {
			}
			v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(v46-int32(4)))) = v19
			v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v50 + v12
			v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if l2 <= v53 {
				v57 = v53 + v12
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v57
			} else {
				if v53 <= l1 {
				} else {
					v57 = l1
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v57
				}
			}
			if l3 != 0 {
				v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				base.MemoryCopy(m, v64+l1, l4, l3)
			} else {
			}
			return int32(0)
		}
	}
}
func F_reservoir_get_next_S(m *base.Module, l0 int32, l1 float64, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v4 float64
	_ = v4
	var v16 float64
	_ = v16
	var v21 int32
	_ = v21
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v62 float64
	_ = v62
	var v66 float64
	_ = v66
	var v68 float64
	_ = v68
	var v75 float64
	_ = v75
	var v76 float64
	_ = v76
	var v79 float64
	_ = v79
	var v87 float64
	_ = v87
	var v88 float64
	_ = v88
	var v90 float64
	_ = v90
	var v93 float64
	_ = v93
	var v95 float64
	_ = v95
	var v96 float64
	_ = v96
	var v97 float64
	_ = v97
	var v99 float64
	_ = v99
	var v100 float64
	_ = v100
	var v102 int32
	_ = v102
	var v103 float64
	_ = v103
	var v107 float64
	_ = v107
	var v121 int64
	_ = v121
	var v122 int64
	_ = v122
	var v123 int64
	_ = v123
	var v144 float64
	_ = v144
	var v149 float64
	_ = v149
	var v150 float64
	_ = v150
	var v151 float64
	_ = v151
	var v155 float64
	_ = v155
	var v157 float64
	_ = v157
	var v159 float64
	_ = v159
	var v162 float64
	_ = v162
	var v165 float64
	_ = v165
	var v171 float64
	_ = v171
	var v172 int32
	_ = v172
	var v173 float64
	_ = v173
	var v176 float64
	_ = v176
	var v180 float64
	_ = v180
	var v181 float64
	_ = v181
	var v185 float64
	_ = v185
	var v193 float64
	_ = v193
	var v194 float64
	_ = v194
	var v197 float64
	_ = v197
	var v207 float64
	_ = v207
	var v231 int64
	_ = v231
	var v232 int64
	_ = v232
	var v233 int64
	_ = v233
	var v254 float64
	_ = v254
	var v257 float64
	_ = v257
	var v259 float64
	_ = v259
	var v260 float64
	_ = v260
	var v263 float64
	_ = v263
	var v271 float64
	_ = v271
	var v291 float64
	_ = v291
	v4 = float64(0)
	v16 = base.F64_convert_i32_s(l2)
	if base.F64_ge(base.F64_mul(v16, float64(22)), l1) != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v291
L2:
	;
	v21 = l0 + int32(8)
	goto L5
L3:
	;
	goto L4
L4:
	;
	v95 = float64(1)
	v96 = base.F64_add(l1, v95)
	v97 = base.F64_sub(l1, v16)
	v99 = base.F64_add(v97, v95)
	v100 = base.F64_div(v96, v99)
	v102 = l0 + int32(8)
	v103 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v107 = v103
	goto L13
L5:
	;
	v39 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	v40 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
	v41 = v39 ^ v40
	*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = base.I64_rotl(v41, int64(37))
	*(*int64)(unsafe.Add(mBase, uint32(v21))) = v41<<(uint(int64(16))%64) ^ base.I64_rotl(v39, int64(24)) ^ v41
	v62 = F_scalbn(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v39*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
	mBase = m.M
	goto L7
L6:
	;
	v66 = base.F64_add(l1, float64(1))
	v68 = base.F64_div(base.F64_sub(v66, v16), v66)
	if base.F64_gt(v68, v62) == int32(0) {
		v291 = v4
		goto L1
	} else {
		goto L9
	}
L7:
	;
	if base.F64_eq(v62, float64(0)) != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	v75 = v66
	v76 = v68
	v79 = v4
	goto L10
L10:
	;
	v87 = float64(1)
	v88 = base.F64_add(v79, v87)
	v90 = base.F64_add(v75, v87)
	v93 = base.F64_mul(v76, base.F64_div(base.F64_sub(v90, v16), v90))
	if base.F64_gt(v93, v62) != 0 {
		v75 = v90
		v76 = v93
		v79 = v88
		goto L10
	} else {
		goto L12
	}
L11:
	;
	v291 = v88
	goto L1
L12:
	;
	goto L11
L13:
	;
	v121 = *(*int64)(unsafe.Add(mBase, uint32(v102)))
	v122 = *(*int64)(unsafe.Add(mBase, uint32(v102)+8))
	v123 = v121 ^ v122
	*(*int64)(unsafe.Add(mBase, uint32(v102)+8)) = base.I64_rotl(v123, int64(37))
	*(*int64)(unsafe.Add(mBase, uint32(v102))) = v123<<(uint(int64(16))%64) ^ base.I64_rotl(v121, int64(24)) ^ v123
	v144 = F_scalbn(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v121*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
	mBase = m.M
	goto L15
L14:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0))) = v271
	v291 = v150
	goto L1
L15:
	;
	if base.F64_eq(v144, float64(0)) != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v149 = base.F64_mul(l1, base.F64_add(v107, float64(-1)))
	v150 = base.F64_floor(v149)
	v151 = base.F64_add(v99, v150)
	v155 = base.F64_add(l1, v149)
	v157 = F_log(m, base.F64_div(base.F64_mul(v151, base.F64_mul(v100, base.F64_mul(v100, v144))), v155))
	mBase = m.M
	v159 = F_exp(m, base.F64_div(v157, v16))
	mBase = m.M
	v162 = base.F64_div(base.F64_mul(v99, base.F64_div(v155, v151)), l1)
	if base.F64_le(v159, v162) != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L14
L18:
	;
	v271 = base.F64_div(v162, v159)
	goto L17
L19:
	;
	goto L20
L20:
	;
	v165 = base.F64_add(l1, v150)
	v171 = base.F64_div(base.F64_mul(base.F64_add(v165, float64(1)), base.F64_div(base.F64_mul(v96, v144), v99)), v155)
	v172 = base.F64_lt(v16, v150)
	if v172 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v173 = v151
	goto L23
L22:
	;
	v173 = v96
	goto L23
L23:
	;
	if base.F64_le(v173, v165) != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	if v172 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v207 = v171
	goto L26
L26:
	;
	goto L33
L27:
	;
	v176 = l1
	goto L29
L28:
	;
	v176 = base.F64_add(v97, v150)
	goto L29
L29:
	;
	v180 = v165
	v181 = v176
	v185 = v171
	goto L30
L30:
	;
	v193 = base.F64_mul(v185, base.F64_div(v180, v181))
	v194 = float64(-1)
	v197 = base.F64_add(v180, v194)
	if base.F64_ge(v197, v173) != 0 {
		v180 = v197
		v181 = base.F64_add(v181, v194)
		v185 = v193
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v207 = v193
	goto L26
L32:
	;
	goto L31
L33:
	;
	v231 = *(*int64)(unsafe.Add(mBase, uint32(v102)))
	v232 = *(*int64)(unsafe.Add(mBase, uint32(v102)+8))
	v233 = v231 ^ v232
	*(*int64)(unsafe.Add(mBase, uint32(v102)+8)) = base.I64_rotl(v233, int64(37))
	*(*int64)(unsafe.Add(mBase, uint32(v102))) = v233<<(uint(int64(16))%64) ^ base.I64_rotl(v231, int64(24)) ^ v233
	v254 = F_scalbn(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v231*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
	mBase = m.M
	goto L35
L34:
	;
	v257 = F_log(m, v207)
	mBase = m.M
	v259 = F_exp(m, base.F64_div(v257, v16))
	mBase = m.M
	v260 = F_log(m, v254)
	mBase = m.M
	v263 = F_exp(m, base.F64_div(base.F64_neg(v260), v16))
	mBase = m.M
	if base.F64_le(v259, base.F64_div(v155, l1)) == int32(0) {
		v107 = v263
		goto L13
	} else {
		goto L37
	}
L35:
	;
	if base.F64_eq(v254, float64(0)) != 0 {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v271 = v263
	goto L17
}
func F_resolve_anyelement_from_others(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
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
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
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
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v9 != 0 {
		v10 = F_getBaseType(m, v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			v12 = F_get_element_type(m, v10)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				if v12 != 0 {
					v109 = v12
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v109
					m.G0 = v7 - int32(-64)
					return
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v17 = m.ExcPending
					if v17 != 0 {
						return
					} else {
						F_errcode(m, int32(67141764))
						mBase = m.M
						v20 = m.ExcPending
						if v20 != 0 {
							return
						} else {
							v21 = F_format_type_be(m, v10)
							mBase = m.M
							v22 = m.ExcPending
							if v22 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+52)) = v21
								*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = int32(_a_F_resolve_anyelement_from_others_0)
								F_errmsg(m, int32(_a_F_resolve_anyelement_from_others_1), v5+int32(-16))
								mBase = m.M
								v30 = m.ExcPending
								if v30 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_resolve_anyelement_from_others_2), int32(604), int32(_a_F_resolve_anyelement_from_others_3))
									mBase = m.M
									v35 = m.ExcPending
									if v35 != 0 {
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
	} else {
		v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v36 != 0 {
			v37 = F_getBaseType(m, v36)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return
			} else {
				v39 = F_get_range_subtype(m, v37)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					if v39 != 0 {
						v109 = v39
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v109
						m.G0 = v7 - int32(-64)
						return
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return
						} else {
							F_errcode(m, int32(67141764))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return
							} else {
								v48 = F_format_type_be(m, v37)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v48
									*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = int32(_a_F_resolve_anyelement_from_others_4)
									F_errmsg(m, int32(_a_F_resolve_anyelement_from_others_5), v5+int32(-32))
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_resolve_anyelement_from_others_2), int32(618), int32(_a_F_resolve_anyelement_from_others_3))
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
			}
		} else {
			v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v63 != 0 {
				v64 = F_getBaseType(m, v63)
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					v66 = F_get_multirange_range(m, v64)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return
					} else {
						if v66 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
								return
							} else {
								F_errcode(m, int32(67141764))
								mBase = m.M
								v121 = m.ExcPending
								if v121 != 0 {
									return
								} else {
									v122 = F_format_type_be(m, v64)
									mBase = m.M
									v123 = m.ExcPending
									if v123 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v122
										*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_resolve_anyelement_from_others_6)
										F_errmsg(m, int32(_a_F_resolve_anyelement_from_others_7), v7)
										mBase = m.M
										v129 = m.ExcPending
										if v129 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_resolve_anyelement_from_others_2), int32(636), int32(_a_F_resolve_anyelement_from_others_3))
											mBase = m.M
											v134 = m.ExcPending
											if v134 != 0 {
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
							v70 = F_getBaseType(m, v66)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return
							} else {
								v72 = F_get_range_subtype(m, v70)
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return
								} else {
									if v72 != 0 {
										v109 = v72
										*(*int32)(unsafe.Add(mBase, uint32(l0))) = v109
										m.G0 = v7 - int32(-64)
										return
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
											return
										} else {
											F_errcode(m, int32(67141764))
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return
											} else {
												v81 = F_format_type_be(m, v70)
												mBase = m.M
												v82 = m.ExcPending
												if v82 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v81
													*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(_a_F_resolve_anyelement_from_others_6)
													F_errmsg(m, int32(_a_F_resolve_anyelement_from_others_8), v5+int32(-48))
													mBase = m.M
													v90 = m.ExcPending
													if v90 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_resolve_anyelement_from_others_2), int32(646), int32(_a_F_resolve_anyelement_from_others_3))
														mBase = m.M
														v95 = m.ExcPending
														if v95 != 0 {
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
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v99 = m.ExcPending
				if v99 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(_a_F_resolve_anyelement_from_others_9), int32(0))
					mBase = m.M
					v103 = m.ExcPending
					if v103 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_resolve_anyelement_from_others_2), int32(650), int32(_a_F_resolve_anyelement_from_others_3))
						mBase = m.M
						v108 = m.ExcPending
						if v108 != 0 {
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
func F_rfree(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
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
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
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
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v8 != int32(_a_F_rfree_0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v11 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(0)
	if v13 == v11 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v18 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+72)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v18
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v13)+92))
	if v22 != v13+int32(180) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	F_pfree(m, v22)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v13)+96))
	if v28 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	return
L9:
	;
	goto L7
L10:
	;
	F_pfree(m, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L8
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v13)+160))
	if v31 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L12
L14:
	;
	F_pfree(m, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L8
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v13)+164))
	if v34 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	F_pfree(m, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L8
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v37 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	F_freesubre(m, int32(0), v37)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L8
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v13)+424))
	if v41 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L24
L26:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v13)+428))
	v44 = v42 - int32(1)
	if int32(0) < v44 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	if v82 != 0 {
		goto L42
	} else {
		goto L43
	}
L29:
	;
	v47 = v41
	v49 = v44
	goto L32
L30:
	;
	goto L31
L31:
	;
	F_pfree(m, v41)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L8
	} else {
		goto L41
	}
L32:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v47)+124))
	if v52 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L31
L34:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v47)+152))
	F_pfree(m, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L8
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v66 = int32(1)
	if v66 < v49 {
		v47 = v47 + int32(88)
		v49 = v49 - v66
		goto L32
	} else {
		goto L40
	}
L37:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v47)+156))
	F_pfree(m, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L8
	} else {
		goto L38
	}
L38:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v47)+160))
	F_pfree(m, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L8
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+124)) = int32(0)
	goto L36
L40:
	;
	goto L33
L41:
	;
	goto L28
L42:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
	F_pfree(m, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L8
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	F_pfree(m, v13)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L8
	} else {
		goto L48
	}
L45:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
	F_pfree(m, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L8
	} else {
		goto L46
	}
L46:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
	F_pfree(m, v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L8
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = int32(0)
	goto L44
L48:
	;
	goto L1
}
