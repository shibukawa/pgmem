package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_ReadyForQuery(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(2)
	if base.Ui32(l0-v9) <= base.Ui32(v9) {
		F_pq_beginmessage(m, v7, int32(90))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			v16 = m.G0
			v18 = v16 - int32(16)
			m.G0 = v18
			v21 = *(*int32)(unsafe.Add(mBase, _c_F_ReadyForQuery[0]))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
			if base.Ui32(int32(20)) <= base.Ui32(v22) {
				F_errstart_cold(m, int32(22), int32(0))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
					if base.Ui32(v29) <= base.Ui32(int32(19)) {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v29<<(uint(int32(2))%32))+uint32(_c_F_ReadyForQuery[1])))
						v36 = v34
					} else {
						v36 = int32(_a_F_ReadyForQuery_0)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v18))) = v36
					F_errmsg_internal(m, int32(_a_F_ReadyForQuery_1), v18)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_ReadyForQuery_2), int32(_a_F_ReadyForQuery_3), int32(_a_F_ReadyForQuery_4))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v46 = int32(*(*int8)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_ReadyForQuery[2]))))
				m.G0 = v18 + int32(16)
				F_enlargeStringInfo(m, v7, int32(1))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return
				} else {
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
					*(*uint8)(unsafe.Add(mBase, uint32(v53+v54))) = uint8(v46)
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v53 + int32(1)
					F_pq_endmessage(m, v7)
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return
					} else {
						v63 = *(*int32)(unsafe.Add(mBase, _c_F_ReadyForQuery[3]))
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
						v65 = m.T0[v64].(func(*base.Module) int32)(m)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return
						} else {
							m.G0 = v7 + int32(16)
							return
						}
					}
				}
			}
		}
	} else {
		m.G0 = v7 + int32(16)
		return
	}
}
func F_RecordKnownAssignedTransactionIds(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v10 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v10 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_RecordKnownAssignedTransactionIds[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v14
	F_errmsg_internal(m, int32(_a_F_RecordKnownAssignedTransactionIds_0), v6)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_RecordKnownAssignedTransactionIds[0]))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v25))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(l0)) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	F_errfinish(m, int32(_a_F_RecordKnownAssignedTransactionIds_1), int32(_a_F_RecordKnownAssignedTransactionIds_2), int32(_a_F_RecordKnownAssignedTransactionIds_3))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	m.G0 = v6 + int32(16)
	return
L9:
	;
	if v37 == int32(0) {
		goto L8
	} else {
		goto L13
	}
L10:
	;
	v37 = base.B2i32(base.Ui32(v25) < base.Ui32(l0))
	goto L9
L11:
	;
	goto L12
L12:
	;
	v37 = base.B2i32(int32(0) < l0-v25)
	goto L9
L13:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_RecordKnownAssignedTransactionIds[0]))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l0))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v41)) == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if v53 != 0 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v53 = base.B2i32(base.Ui32(v41) < base.Ui32(l0))
	goto L14
L16:
	;
	goto L17
L17:
	;
	v53 = int32(base.Ui32(v41-l0) >> (uint(int32(31)) % 32))
	goto L14
L18:
	;
	v55 = v41
	goto L21
L19:
	;
	goto L20
L20:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_RecordKnownAssignedTransactionIds[1]))
	if base.Ui32(v81) <= base.Ui32(int32(1)) {
		goto L32
	} else {
		goto L33
	}
L21:
	;
	v57 = int32(3)
	v59 = v55 + int32(1)
	if base.Ui32(v59) <= base.Ui32(v57) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L20
L23:
	;
	v62 = v57
	goto L25
L24:
	;
	v62 = v59
	goto L25
L25:
	;
	F_ExtendSUBTRANS(m, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l0))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v62)) == int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if v76 != 0 {
		v55 = v62
		goto L21
	} else {
		goto L31
	}
L28:
	;
	v76 = base.B2i32(base.Ui32(v62) < base.Ui32(l0))
	goto L27
L29:
	;
	goto L30
L30:
	;
	v76 = int32(base.Ui32(v62-l0) >> (uint(int32(31)) % 32))
	goto L27
L31:
	;
	goto L22
L32:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RecordKnownAssignedTransactionIds[0])) = l0
	goto L8
L33:
	;
	goto L34
L34:
	;
	v86 = int32(3)
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_RecordKnownAssignedTransactionIds[0]))
	v90 = v88 + int32(1)
	if base.Ui32(v90) <= base.Ui32(v86) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v93 = v86
	goto L37
L36:
	;
	v93 = v90
	goto L37
L37:
	;
	F_KnownAssignedXidsAdd(m, v93, l0, int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RecordKnownAssignedTransactionIds[0])) = l0
	F_AdvanceNextFullTransactionIdPastXid(m, l0)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	goto L8
}
func F_RegisterDynamicBackgroundWorker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	v3 = int32(0)
	v11 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RegisterDynamicBackgroundWorker[0])))
	if v11 != int32(1) {
		v142 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v142
L2:
	;
	v15 = F_SanityCheckBackgroundWorker(m, l0, int32(21))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	if v15 == int32(0) {
		v142 = v3
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterDynamicBackgroundWorker[1]))
	v27 = F_LWLockAcquire(m, v23+int32(_a_F_RegisterDynamicBackgroundWorker_0), int32(0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterDynamicBackgroundWorker[2]))
	v32 = v21 & int32(16)
	if v32 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if int32(0) < v49 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterDynamicBackgroundWorker[3]))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	if base.Ui32(v37-v38) < base.Ui32(v36) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterDynamicBackgroundWorker[1]))
	F_LWLockRelease(m, v42+int32(_a_F_RegisterDynamicBackgroundWorker_0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	return int32(0)
L11:
	;
	v58 = v3
	goto L14
L12:
	;
	goto L13
L13:
	;
	v135 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterDynamicBackgroundWorker[1]))
	F_LWLockRelease(m, v135+int32(_a_F_RegisterDynamicBackgroundWorker_0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L3
	} else {
		goto L30
	}
L14:
	;
	v65 = v30 + int32(16) + v58*int32(1480)
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v66 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L13
L16:
	;
	base.MemoryCopy(m, v65+int32(16), l0, int32(1460))
	*(*int32)(unsafe.Add(mBase, uint32(v65)+4)) = int32(-1)
	v75 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v65)+1)) = uint8(v75)
	v77 = *(*int64)(unsafe.Add(mBase, uint32(v65)+8))
	v79 = v77 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v65)+8)) = v79
	if v32 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v122 = v58 + int32(1)
	if v122 != v49 {
		v58 = v122
		goto L14
	} else {
		goto L29
	}
L19:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v81 + int32(1)
	goto L21
L20:
	;
	goto L21
L21:
	;
	v85 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v65))) = uint8(v85)
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterDynamicBackgroundWorker[1]))
	F_LWLockRelease(m, v89+int32(_a_F_RegisterDynamicBackgroundWorker_0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RegisterDynamicBackgroundWorker[0])))
	if v96 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if l1 == int32(0) {
		v142 = v85
		goto L1
	} else {
		goto L27
	}
L24:
	;
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterDynamicBackgroundWorker[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v100+int32(24)))) = int32(1)
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterDynamicBackgroundWorker[5]))
	v109 = F_pgmem_kill(m, v107, int32(10))
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
	v113 = F_palloc(m, int32(16))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v113))) = v58
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v117)+8)) = v79
	return int32(1)
L29:
	;
	goto L15
L30:
	;
	v142 = int32(0)
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
func F_RelfilenumberMapInvalidateCallback(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v10 = v7 + int32(12)
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_RelfilenumberMapInvalidateCallback[0]))
	F_hash_seq_init(m, v10, v12)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v15 = F_hash_seq_search(m, v10)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L19
	}
L4:
	;
	if v15 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v19 = v15
	goto L8
L6:
	;
	goto L7
L7:
	;
	m.G0 = v7 + int32(32)
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
	v39 = F_hash_seq_search(m, v7+int32(12))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L17
	}
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_RelfilenumberMapInvalidateCallback[0]))
	v32 = F_hash_search(m, v29, v19, int32(2), int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L15
	}
L12:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	if v23 == int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	if l1 != v23 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	if v32 == int32(0) {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	goto L10
L17:
	;
	if v39 != 0 {
		v19 = v39
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
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(_a_F_RelfilenumberMapInvalidateCallback_1), int32(76), int32(_a_F_RelfilenumberMapInvalidateCallback_2))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
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
	var v65 int32
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
	var v161 int32
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
	var v242 int32
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
	v16 = v14 - int32(208)
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
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
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
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
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
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
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
	m.G0 = v16 + int32(208)
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
func F_ReportApplyConflict(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
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
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
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
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v285 int32
	_ = v285
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
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int64
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v424 int64
	_ = v424
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
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
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	v19 = m.G0
	v21 = v19 - int32(320)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_initStringInfo(m, v21+int32(268))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
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
	v413 = int32(0)
	v415 = *(*int32)(unsafe.Add(mBase, _c_F_ReportApplyConflict[0]))
	v416 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v415))))
	v418 = F_pgstat_prep_pending_entry(m, int32(5), v413, v416, v413)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L1
	} else {
		goto L124
	}
L4:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v30 <= int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v49 = int32(0)
	goto L6
L6:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53+v49<<(uint(int32(2))%32))))
	v58 = *(*int64)(unsafe.Add(mBase, uint32(v57)+16))
	v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+12)))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	F_initStringInfo(m, v21+int32(288))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	goto L3
L8:
	;
	switch l3 {
	case 0, 2, 6:
		goto L17
	case 1:
		goto L16
	case 3:
		goto L15
	case 4:
		goto L14
	case 5:
		goto L13
	default:
		goto L10
	}
L9:
	;
	if v62 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L10:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)+52))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v232)+56))
	v236 = v21 + int32(304)
	F_initStringInfo(m, v236)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L65
	}
L11:
	;
	v218 = F_timestamptz_to_str(m, v58)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L63
	}
L12:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)+52))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v211)+56))
	F_initStringInfo(m, v21+int32(304))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L62
	}
L13:
	;
	F_appendStringInfoString(m, v21+int32(288), int32(_a_F_ReportApplyConflict_0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L61
	}
L14:
	;
	if v59 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L15:
	;
	F_appendStringInfoString(m, v21+int32(288), int32(_a_F_ReportApplyConflict_1))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L47
	}
L16:
	;
	if v59 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L17:
	;
	if v58 != int64(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	if v59 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v107 = F_get_rel_name(m, v61)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L32
	}
L21:
	;
	v71 = F_get_rel_name(m, v61)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v87 = F_replorigin_by_oid(m, v59, v21+int32(284))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L27
	}
L24:
	;
	v73 = F_timestamptz_to_str(m, v58)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+136)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v21)+132)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v21)+128)) = v71
	F_appendStringInfo(m, v21+int32(288), int32(_a_F_ReportApplyConflict_2), v21+int32(128))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	goto L10
L27:
	;
	v89 = F_get_rel_name(m, v61)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	if v87 == int32(0) {
		goto L11
	} else {
		goto L29
	}
L29:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v21)+284))
	v94 = F_timestamptz_to_str(m, v58)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+156)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v21)+152)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v21)+148)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v21)+144)) = v89
	F_appendStringInfo(m, v21+int32(288), int32(_a_F_ReportApplyConflict_3), v21+int32(144))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	goto L10
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+116)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v21)+112)) = v107
	F_appendStringInfo(m, v21+int32(288), int32(_a_F_ReportApplyConflict_4), v21+int32(112))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L10
L34:
	;
	v120 = F_timestamptz_to_str(m, v58)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v133 = F_replorigin_by_oid(m, v59, v21+int32(284))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+180)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v21)+176)) = v60
	F_appendStringInfo(m, v21+int32(288), int32(_a_F_ReportApplyConflict_5), v21+int32(176))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	goto L12
L39:
	;
	if v133 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v21)+284))
	v136 = F_timestamptz_to_str(m, v58)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v148 = F_timestamptz_to_str(m, v58)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+200)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v21)+196)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v21)+192)) = v135
	F_appendStringInfo(m, v21+int32(288), int32(_a_F_ReportApplyConflict_6), v21+int32(192))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	goto L12
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+212)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v21)+208)) = v60
	F_appendStringInfo(m, v21+int32(288), int32(_a_F_ReportApplyConflict_7), v21+int32(208))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	goto L12
L47:
	;
	goto L12
L48:
	;
	v166 = F_timestamptz_to_str(m, v58)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v179 = F_replorigin_by_oid(m, v59, v21+int32(284))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+228)) = v166
	*(*int32)(unsafe.Add(mBase, uint32(v21)+224)) = v60
	F_appendStringInfo(m, v21+int32(288), int32(_a_F_ReportApplyConflict_8), v21+int32(224))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	goto L12
L53:
	;
	if v179 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v21)+284))
	v182 = F_timestamptz_to_str(m, v58)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v194 = F_timestamptz_to_str(m, v58)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L59
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+248)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v21)+244)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v21)+240)) = v181
	F_appendStringInfo(m, v21+int32(288), int32(_a_F_ReportApplyConflict_9), v21+int32(240))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	goto L12
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+260)) = v194
	*(*int32)(unsafe.Add(mBase, uint32(v21)+256)) = v60
	F_appendStringInfo(m, v21+int32(288), int32(_a_F_ReportApplyConflict_10), v21+int32(256))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	goto L12
L61:
	;
	goto L12
L62:
	;
	v256 = v211
	v257 = v212
	v259 = v213
	goto L9
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+168)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v21)+164)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v21)+160)) = v89
	F_appendStringInfo(m, v21+int32(288), int32(_a_F_ReportApplyConflict_11), v21+int32(160))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	goto L10
L65:
	;
	if base.B2i32(base.Ui32(int32(6)) < base.Ui32(l3))|base.B2i32(int32(1)<<(uint(l3)%32)&int32(69) == int32(0)) != 0 {
		v256 = v232
		v257 = v233
		v259 = v234
		goto L9
	} else {
		goto L66
	}
L66:
	;
	v246 = F_build_index_value_desc(m, l0, v232, v62, v61)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	if v246 == int32(0) {
		v256 = v232
		v257 = v233
		v259 = v234
		goto L9
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+96)) = v246
	F_appendStringInfo(m, v236, int32(_a_F_ReportApplyConflict_12), v21+int32(96))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v256 = v232
	v257 = v233
	v259 = v234
	goto L9
L70:
	;
	if l5 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L71:
	;
	v264 = F_ExecBuildSlotValueDescription(m, v259, v62, v257, int32(0))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	if v264 == int32(0) {
		goto L70
	} else {
		goto L73
	}
L73:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v21)+308))
	if v268 <= int32(0) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v278 = int32(_a_F_ReportApplyConflict_13)
	goto L76
L75:
	;
	F_appendStringInfoString(m, v21+int32(304), int32(_a_F_ReportApplyConflict_14))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L77
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+80)) = v264
	F_appendStringInfo(m, v21+int32(304), v278, v21+int32(80))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L78
	}
L77:
	;
	v278 = int32(_a_F_ReportApplyConflict_15)
	goto L76
L78:
	;
	goto L70
L79:
	;
	if l4 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L80:
	;
	v290 = F_ExecGetInsertedCols(m, l1, l0)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v292 = F_ExecGetUpdatedCols(m, l1, l0)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v294 = F_bms_union(m, v290, v292)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v296 = F_ExecBuildSlotValueDescription(m, v259, l5, v257, v294)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	if v296 == int32(0) {
		goto L79
	} else {
		goto L85
	}
L85:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v21)+308))
	if v300 <= int32(0) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v310 = int32(_a_F_ReportApplyConflict_16)
	goto L88
L87:
	;
	F_appendStringInfoString(m, v21+int32(304), int32(_a_F_ReportApplyConflict_14))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L89
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = v296
	F_appendStringInfo(m, v21+int32(304), v310, v21-int32(-64))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L90
	}
L89:
	;
	v310 = int32(_a_F_ReportApplyConflict_17)
	goto L88
L90:
	;
	goto L79
L91:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v21)+308))
	if v357 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L92:
	;
	v322 = F_GetRelationIdentityOrPK(m, v256)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L94
	}
L93:
	;
	if v329 == int32(0) {
		goto L91
	} else {
		goto L100
	}
L94:
	;
	if v322 != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v324 = F_build_index_value_desc(m, l0, v256, l4, v322)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v327 = F_ExecBuildSlotValueDescription(m, v259, l4, v257, int32(0))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L99
	}
L98:
	;
	v329 = v324
	goto L93
L99:
	;
	v329 = v327
	goto L93
L100:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v21)+308))
	if int32(0) < v332 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = v329
	F_appendStringInfo(m, v21+int32(304), v346, v21+int32(48))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L112
	}
L102:
	;
	F_appendStringInfoString(m, v21+int32(304), int32(_a_F_ReportApplyConflict_14))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	if v322 != 0 {
		goto L109
	} else {
		goto L110
	}
L105:
	;
	if v322 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v342 = int32(_a_F_ReportApplyConflict_18)
	goto L108
L107:
	;
	v342 = int32(_a_F_ReportApplyConflict_19)
	goto L108
L108:
	;
	v346 = v342
	goto L101
L109:
	;
	v345 = int32(_a_F_ReportApplyConflict_20)
	goto L111
L110:
	;
	v345 = int32(_a_F_ReportApplyConflict_21)
	goto L111
L111:
	;
	v346 = v345
	goto L101
L112:
	;
	goto L91
L113:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v21)+272))
	if int32(0) < v377 {
		goto L118
	} else {
		goto L119
	}
L114:
	;
	F_appendStringInfoChar(m, v21+int32(304), int32(46))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v21)+304))
	if v365 == int32(0) {
		goto L113
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v365
	F_appendStringInfo(m, v21+int32(288), int32(_a_F_ReportApplyConflict_22), v21+int32(32))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	goto L113
L118:
	;
	F_appendStringInfoChar(m, v21+int32(268), int32(10))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v21)+288))
	F_appendStringInfoString(m, v21+int32(268), v387)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L122
	}
L121:
	;
	goto L120
L122:
	;
	v391 = v49 + int32(1)
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v391 < v392 {
		v49 = v391
		goto L6
	} else {
		goto L123
	}
L123:
	;
	goto L7
L124:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v418)+12))
	v423 = v420 + l3<<(uint(int32(3))%32)
	v424 = *(*int64)(unsafe.Add(mBase, uint32(v423)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v423)+16)) = v424 + int64(1)
	v429 = F_errstart(m, l2, int32(0))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	if v429 != 0 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	if base.Ui32(l3) <= base.Ui32(int32(6)) {
		goto L129
	} else {
		goto L130
	}
L127:
	;
	goto L128
L128:
	;
	m.G0 = v21 + int32(320)
	return
L129:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(l3<<(uint(int32(2))%32))+uint32(_c_F_ReportApplyConflict[1])))
	F_errcode(m, v435)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L1
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(l3<<(uint(int32(2))%32))+uint32(_c_F_ReportApplyConflict[2])))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v23)+48))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v441)+68))
	v443 = F_get_namespace_name(m, v442)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L1
	} else {
		goto L133
	}
L132:
	;
	goto L131
L133:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v23)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v443
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v445 + int32(4)
	F_errmsg(m, int32(_a_F_ReportApplyConflict_23), v21+int32(16))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v21)+268))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v456
	F_errdetail_internal(m, int32(_a_F_ReportApplyConflict_24), v21)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	F_errfinish(m, int32(_a_F_ReportApplyConflict_25), int32(130), int32(_a_F_ReportApplyConflict_26))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	goto L128
}
func F_ReservedPLKeywords_hash_func(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	v3 = int32(0)
	if l1 == v3 {
		v83 = int32(1)
		v91 = int32(0)
	} else {
		v13 = int32(1)
		if l1 != v13 {
			v21 = l0
			v22 = int32(0)
			v23 = v13
			v24 = v3
			for {
				v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
				v30 = int32(32)
				v31 = v29 | v30
				v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
				v34 = v32 | v30
				v35 = int32(_a_F_ReservedPLKeywords_hash_func_0)
				v40 = v31 + (v34+v23*v35)*v35
				v41 = int32(257)
				v46 = (v24*v41+v34)*v41 + v31
				v47 = int32(2)
				v48 = v21 + v47
				v50 = v22 + v47
				if v50 != l1&int32(-2) {
					v21 = v48
					v22 = v50
					v23 = v40
					v24 = v46
					continue
				} else {
					break
				}
				break
			}
			if l1&int32(1) == int32(0) {
				v73 = v40
				v74 = v46
			} else {
				v54 = v48
				v56 = v40
				v57 = v46
				v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
				v64 = v62 | int32(32)
				v73 = v64 + v56*int32(_a_F_ReservedPLKeywords_hash_func_0)
				v74 = v57*int32(257) + v64
			}
		} else {
			v54 = l0
			v56 = v13
			v57 = v3
			v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
			v64 = v62 | int32(32)
			v73 = v64 + v56*int32(_a_F_ReservedPLKeywords_hash_func_0)
			v74 = v57*int32(257) + v64
		}
		v79 = int32(49)
		v80 = base.I32_rem_u_s(v73, v79)
		v82 = base.I32_rem_u_s(v74, v79)
		v83 = v80
		v91 = v82
	}
	v94 = int32(*(*int8)(unsafe.Add(mBase, uint32(v83)+uint32(_c_F_ReservedPLKeywords_hash_func[0]))))
	v97 = int32(*(*int8)(unsafe.Add(mBase, uint32(v91)+uint32(_c_F_ReservedPLKeywords_hash_func[0]))))
	return v94 + v97
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
	F_errfinish(m, int32(_a_F__readBitmapset_1), int32(234), int32(_a_F__readBitmapset_2))
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
	F_errfinish(m, int32(_a_F__readBitmapset_1), int32(217), int32(_a_F__readBitmapset_2))
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
	F_errfinish(m, int32(_a_F__readBitmapset_1), int32(219), int32(_a_F__readBitmapset_2))
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
	F_errfinish(m, int32(_a_F__readBitmapset_1), int32(223), int32(_a_F__readBitmapset_2))
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
	F_errfinish(m, int32(_a_F__readBitmapset_1), int32(239), int32(_a_F__readBitmapset_2))
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
	F_errfinish(m, int32(_a_F__readBitmapset_1), int32(225), int32(_a_F__readBitmapset_2))
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
func F_r_SUFFIX_KAN_OK(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
	return base.B2i32(v3&int32(-2) != int32(2))
}
func F_r_VI_1(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v26 int32
	_ = v26
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v4 <= v5 {
		v72 = v2
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7+v4-int32(1)))))
		if v11 != int32(105) {
			v72 = v2
		} else {
			v15 = v4 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v15
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v15 <= v26 {
				v69 = int32(-1)
			} else {
				v38 = int32(1)
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+v15-v38))))
				if int32(246) < v43 {
					v65 = v38
				} else {
					v45 = v43 - int32(97)
					if v45 < int32(0) {
						v65 = v38
					} else {
						v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v45)>>(uint(int32(3))%32)))+uint32(_c_F_r_VI_1[0]))))
						if int32(base.Ui32(v51)>>(uint(v45&int32(7))%32))&int32(1) == int32(0) {
							v65 = v38
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v15 - int32(1)
							v65 = int32(0)
						}
					}
				}
				v69 = v65
			}
			v72 = base.B2i32(v69 == int32(0))
		}
	}
	return v72
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
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v4
	v7 = v4 - int32(1)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v7 <= v8 {
		v45 = v2
		return v45
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v7))))
		if base.B2i32(v12 != int32(190))&base.B2i32(v12 != int32(141)) != 0 {
			v45 = v2
			return v45
		} else {
			v20 = F_find_among_b(m, l0, int32(_a_F_r_fix_chdz_0), int32(2))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				if v20 == int32(0) {
					v45 = v2
					return v45
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v26
					switch v20 - int32(1) {
					case 0:
						v32 = F_slice_from_s(m, l0, int32(1), int32(_a_F_r_fix_chdz_1))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							if int32(0) <= v32 {
								v45 = int32(1)
							} else {
								v45 = v32
							}
							return v45
						}
					case 1:
						v38 = F_slice_from_s(m, l0, int32(1), int32(_a_F_r_fix_chdz_2))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							if v38 < int32(0) {
								v45 = v38
							} else {
								v45 = int32(1)
							}
							return v45
						}
					default:
						v45 = int32(1)
						return v45
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
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
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
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
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
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v7-int32(4))))
	if v15 == v2 {
		v89 = int32(0)
	} else {
		v20 = v15 & int32(3)
		if base.Ui32(v15) < base.Ui32(int32(4)) {
			v56 = v7
			v57 = int32(0)
			v62 = v56
			v63 = v57
			v67 = v2
			for {
				v68 = int32(*(*int8)(unsafe.Add(mBase, uint32(v62))))
				v71 = v63 + base.B2i32(int32(-65) < v68)
				v72 = int32(1)
				v75 = v67 + v72
				if v75 != v20 {
					v62 = v62 + v72
					v63 = v71
					v67 = v75
					continue
				} else {
					break
				}
				break
			}
			v78 = v71
		} else {
			v27 = v7
			v28 = int32(0)
			v31 = v2
			for {
				v33 = int32(*(*int8)(unsafe.Add(mBase, uint32(v27))))
				v34 = int32(-65)
				v37 = int32(*(*int8)(unsafe.Add(mBase, uint32(v27)+1)))
				v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v27)+2)))
				v45 = int32(*(*int8)(unsafe.Add(mBase, uint32(v27)+3)))
				v48 = v28 + base.B2i32(v34 < v33) + base.B2i32(v34 < v37) + base.B2i32(v34 < v41) + base.B2i32(v34 < v45)
				v49 = int32(4)
				v50 = v27 + v49
				v52 = v31 + v49
				if v52 != v15&int32(-4) {
					v27 = v50
					v28 = v48
					v31 = v52
					continue
				} else {
					break
				}
				break
			}
			if v20 == int32(0) {
				v78 = v48
			} else {
				v56 = v50
				v57 = v48
				v62 = v56
				v63 = v57
				v67 = v2
				for {
					v68 = int32(*(*int8)(unsafe.Add(mBase, uint32(v62))))
					v71 = v63 + base.B2i32(int32(-65) < v68)
					v72 = int32(1)
					v75 = v67 + v72
					if v75 != v20 {
						v62 = v62 + v72
						v63 = v71
						v67 = v75
						continue
					} else {
						break
					}
					break
				}
				v78 = v71
			}
		}
		v89 = v78
	}
	if v89 < int32(4) {
		v347 = v2
		return v347
	} else {
		v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v92
		v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v94
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v94
		v99 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_0), int32(17))
		mBase = m.M
		v102 = m.ExcPending
		if v102 != 0 {
			return int32(0)
		} else {
			if v99 == int32(0) {
				v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v225
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v225
				v228 = int32(3)
				v230 = int32(0)
				v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v225-v233 < v228 {
					v243 = v230
				} else {
					v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v239 = F_memcmp(m, v236+v225-v228, int32(_a_F_r_fix_ending_1), v228)
					mBase = m.M
					if v239 != 0 {
						v243 = v230
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v225 - v228
						v243 = int32(1)
					}
				}
				if v243 == int32(0) {
					v347 = v2
					return v347
				} else {
					v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v250 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_2), int32(6))
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
								v282 = v273 - v254
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v282
								v284 = v282
								*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v284
								v286 = F_slice_del(m, l0)
								mBase = m.M
								v287 = m.ExcPending
								if v287 != 0 {
									return int32(0)
								} else {
									if int32(0) <= v286 {
										v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
										v347 = int32(1)
									} else {
										v347 = v286
									}
									return v347
								}
							} else {
								v277 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_4), int32(6))
								mBase = m.M
								v278 = m.ExcPending
								if v278 != 0 {
									return int32(0)
								} else {
									if v277 != 0 {
										v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										v284 = v279
									} else {
										v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v282 = v280 - v254
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v282
										v284 = v282
									}
									*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v284
									v286 = F_slice_del(m, l0)
									mBase = m.M
									v287 = m.ExcPending
									if v287 != 0 {
										return int32(0)
									} else {
										if int32(0) <= v286 {
											v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
											v347 = int32(1)
										} else {
											v347 = v286
										}
										return v347
									}
								}
							}
						} else {
							v290 = v247 - v246
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v252 - v290
							v295 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_5), int32(11))
							mBase = m.M
							v296 = m.ExcPending
							if v296 != 0 {
								return int32(0)
							} else {
								if v295 == int32(0) {
									v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v323 - v290
									v328 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_6), int32(9))
									mBase = m.M
									v329 = m.ExcPending
									if v329 != 0 {
										return int32(0)
									} else {
										if v328 == int32(0) {
											v347 = v2
											return v347
										} else {
											v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v333 = v332 - v290
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v333
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v333
											v336 = F_slice_del(m, l0)
											mBase = m.M
											v337 = m.ExcPending
											if v337 != 0 {
												return int32(0)
											} else {
												if v336 < int32(0) {
													v347 = v336
												} else {
													v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
													v347 = int32(1)
												}
												return v347
											}
										}
									}
								} else {
									v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v299
									v301 = int32(3)
									v303 = int32(0)
									v306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v299-v306 < v301 {
										v316 = v303
									} else {
										v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v312 = F_memcmp(m, v309+v299-v301, int32(_a_F_r_fix_ending_7), v301)
										mBase = m.M
										if v312 != 0 {
											v316 = v303
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v299 - v301
											v316 = int32(1)
										}
									}
									if v316 == int32(0) {
										v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v323 - v290
										v328 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_6), int32(9))
										mBase = m.M
										v329 = m.ExcPending
										if v329 != 0 {
											return int32(0)
										} else {
											if v328 == int32(0) {
												v347 = v2
												return v347
											} else {
												v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												v333 = v332 - v290
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v333
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v333
												v336 = F_slice_del(m, l0)
												mBase = m.M
												v337 = m.ExcPending
												if v337 != 0 {
													return int32(0)
												} else {
													if v336 < int32(0) {
														v347 = v336
													} else {
														v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
														v347 = int32(1)
													}
													return v347
												}
											}
										}
									} else {
										v319 = F_slice_del(m, l0)
										mBase = m.M
										v320 = m.ExcPending
										if v320 != 0 {
											return int32(0)
										} else {
											if int32(0) <= v319 {
												v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
												v347 = int32(1)
											} else {
												v347 = v319
											}
											return v347
										}
									}
								}
							}
						}
					}
				}
			} else {
				v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v105
				switch v99 - int32(1) {
				case 0:
					v109 = F_slice_del(m, l0)
					mBase = m.M
					v110 = m.ExcPending
					if v110 != 0 {
						return int32(0)
					} else {
						if int32(0) <= v109 {
							v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
							v347 = int32(1)
						} else {
							v347 = v109
						}
						return v347
					}
				case 1:
					v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v116 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_8), int32(3))
					mBase = m.M
					v117 = m.ExcPending
					if v117 != 0 {
						return int32(0)
					} else {
						if v116 == int32(0) {
							v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v225
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v225
							v228 = int32(3)
							v230 = int32(0)
							v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v225-v233 < v228 {
								v243 = v230
							} else {
								v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v239 = F_memcmp(m, v236+v225-v228, int32(_a_F_r_fix_ending_1), v228)
								mBase = m.M
								if v239 != 0 {
									v243 = v230
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v225 - v228
									v243 = int32(1)
								}
							}
							if v243 == int32(0) {
								v347 = v2
								return v347
							} else {
								v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v250 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_2), int32(6))
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
											v282 = v273 - v254
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v282
											v284 = v282
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v284
											v286 = F_slice_del(m, l0)
											mBase = m.M
											v287 = m.ExcPending
											if v287 != 0 {
												return int32(0)
											} else {
												if int32(0) <= v286 {
													v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
													v347 = int32(1)
												} else {
													v347 = v286
												}
												return v347
											}
										} else {
											v277 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_4), int32(6))
											mBase = m.M
											v278 = m.ExcPending
											if v278 != 0 {
												return int32(0)
											} else {
												if v277 != 0 {
													v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
													v284 = v279
												} else {
													v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													v282 = v280 - v254
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v282
													v284 = v282
												}
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v284
												v286 = F_slice_del(m, l0)
												mBase = m.M
												v287 = m.ExcPending
												if v287 != 0 {
													return int32(0)
												} else {
													if int32(0) <= v286 {
														v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
														v347 = int32(1)
													} else {
														v347 = v286
													}
													return v347
												}
											}
										}
									} else {
										v290 = v247 - v246
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v252 - v290
										v295 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_5), int32(11))
										mBase = m.M
										v296 = m.ExcPending
										if v296 != 0 {
											return int32(0)
										} else {
											if v295 == int32(0) {
												v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v323 - v290
												v328 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_6), int32(9))
												mBase = m.M
												v329 = m.ExcPending
												if v329 != 0 {
													return int32(0)
												} else {
													if v328 == int32(0) {
														v347 = v2
														return v347
													} else {
														v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														v333 = v332 - v290
														*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v333
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v333
														v336 = F_slice_del(m, l0)
														mBase = m.M
														v337 = m.ExcPending
														if v337 != 0 {
															return int32(0)
														} else {
															if v336 < int32(0) {
																v347 = v336
															} else {
																v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
																v347 = int32(1)
															}
															return v347
														}
													}
												}
											} else {
												v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v299
												v301 = int32(3)
												v303 = int32(0)
												v306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												if v299-v306 < v301 {
													v316 = v303
												} else {
													v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													v312 = F_memcmp(m, v309+v299-v301, int32(_a_F_r_fix_ending_7), v301)
													mBase = m.M
													if v312 != 0 {
														v316 = v303
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v299 - v301
														v316 = int32(1)
													}
												}
												if v316 == int32(0) {
													v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v323 - v290
													v328 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_6), int32(9))
													mBase = m.M
													v329 = m.ExcPending
													if v329 != 0 {
														return int32(0)
													} else {
														if v328 == int32(0) {
															v347 = v2
															return v347
														} else {
															v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															v333 = v332 - v290
															*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v333
															*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v333
															v336 = F_slice_del(m, l0)
															mBase = m.M
															v337 = m.ExcPending
															if v337 != 0 {
																return int32(0)
															} else {
																if v336 < int32(0) {
																	v347 = v336
																} else {
																	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
																	v347 = int32(1)
																}
																return v347
															}
														}
													}
												} else {
													v319 = F_slice_del(m, l0)
													mBase = m.M
													v320 = m.ExcPending
													if v320 != 0 {
														return int32(0)
													} else {
														if int32(0) <= v319 {
															v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
															*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
															v347 = int32(1)
														} else {
															v347 = v319
														}
														return v347
													}
												}
											}
										}
									}
								}
							}
						} else {
							v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v120 + (v105 - v113)
							v124 = F_slice_del(m, l0)
							mBase = m.M
							v125 = m.ExcPending
							if v125 != 0 {
								return int32(0)
							} else {
								if int32(0) <= v124 {
									v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
									v347 = int32(1)
								} else {
									v347 = v124
								}
								return v347
							}
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
					v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
					if v147 == int32(0) {
						v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v225
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v225
						v228 = int32(3)
						v230 = int32(0)
						v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v225-v233 < v228 {
							v243 = v230
						} else {
							v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v239 = F_memcmp(m, v236+v225-v228, int32(_a_F_r_fix_ending_1), v228)
							mBase = m.M
							if v239 != 0 {
								v243 = v230
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v225 - v228
								v243 = int32(1)
							}
						}
						if v243 == int32(0) {
							v347 = v2
							return v347
						} else {
							v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v250 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_2), int32(6))
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
										v282 = v273 - v254
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v282
										v284 = v282
										*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v284
										v286 = F_slice_del(m, l0)
										mBase = m.M
										v287 = m.ExcPending
										if v287 != 0 {
											return int32(0)
										} else {
											if int32(0) <= v286 {
												v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
												v347 = int32(1)
											} else {
												v347 = v286
											}
											return v347
										}
									} else {
										v277 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_4), int32(6))
										mBase = m.M
										v278 = m.ExcPending
										if v278 != 0 {
											return int32(0)
										} else {
											if v277 != 0 {
												v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												v284 = v279
											} else {
												v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												v282 = v280 - v254
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v282
												v284 = v282
											}
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v284
											v286 = F_slice_del(m, l0)
											mBase = m.M
											v287 = m.ExcPending
											if v287 != 0 {
												return int32(0)
											} else {
												if int32(0) <= v286 {
													v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
													v347 = int32(1)
												} else {
													v347 = v286
												}
												return v347
											}
										}
									}
								} else {
									v290 = v247 - v246
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v252 - v290
									v295 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_5), int32(11))
									mBase = m.M
									v296 = m.ExcPending
									if v296 != 0 {
										return int32(0)
									} else {
										if v295 == int32(0) {
											v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v323 - v290
											v328 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_6), int32(9))
											mBase = m.M
											v329 = m.ExcPending
											if v329 != 0 {
												return int32(0)
											} else {
												if v328 == int32(0) {
													v347 = v2
													return v347
												} else {
													v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													v333 = v332 - v290
													*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v333
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v333
													v336 = F_slice_del(m, l0)
													mBase = m.M
													v337 = m.ExcPending
													if v337 != 0 {
														return int32(0)
													} else {
														if v336 < int32(0) {
															v347 = v336
														} else {
															v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
															*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
															v347 = int32(1)
														}
														return v347
													}
												}
											}
										} else {
											v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v299
											v301 = int32(3)
											v303 = int32(0)
											v306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v299-v306 < v301 {
												v316 = v303
											} else {
												v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v312 = F_memcmp(m, v309+v299-v301, int32(_a_F_r_fix_ending_7), v301)
												mBase = m.M
												if v312 != 0 {
													v316 = v303
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v299 - v301
													v316 = int32(1)
												}
											}
											if v316 == int32(0) {
												v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v323 - v290
												v328 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_6), int32(9))
												mBase = m.M
												v329 = m.ExcPending
												if v329 != 0 {
													return int32(0)
												} else {
													if v328 == int32(0) {
														v347 = v2
														return v347
													} else {
														v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														v333 = v332 - v290
														*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v333
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v333
														v336 = F_slice_del(m, l0)
														mBase = m.M
														v337 = m.ExcPending
														if v337 != 0 {
															return int32(0)
														} else {
															if v336 < int32(0) {
																v347 = v336
															} else {
																v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
																v347 = int32(1)
															}
															return v347
														}
													}
												}
											} else {
												v319 = F_slice_del(m, l0)
												mBase = m.M
												v320 = m.ExcPending
												if v320 != 0 {
													return int32(0)
												} else {
													if int32(0) <= v319 {
														v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
														v347 = int32(1)
													} else {
														v347 = v319
													}
													return v347
												}
											}
										}
									}
								}
							}
						}
					} else {
						v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v151 = int32(3)
						v153 = int32(0)
						v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v155-v156 < v151 {
							v166 = v153
						} else {
							v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v162 = F_memcmp(m, v159+v155-v151, int32(_a_F_r_fix_ending_12), v151)
							mBase = m.M
							if v162 != 0 {
								v166 = v153
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v155 - v151
								v166 = int32(1)
							}
						}
						if v166 != 0 {
							v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v225
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v225
							v228 = int32(3)
							v230 = int32(0)
							v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v225-v233 < v228 {
								v243 = v230
							} else {
								v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v239 = F_memcmp(m, v236+v225-v228, int32(_a_F_r_fix_ending_1), v228)
								mBase = m.M
								if v239 != 0 {
									v243 = v230
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v225 - v228
									v243 = int32(1)
								}
							}
							if v243 == int32(0) {
								v347 = v2
								return v347
							} else {
								v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v250 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_2), int32(6))
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
											v282 = v273 - v254
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v282
											v284 = v282
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v284
											v286 = F_slice_del(m, l0)
											mBase = m.M
											v287 = m.ExcPending
											if v287 != 0 {
												return int32(0)
											} else {
												if int32(0) <= v286 {
													v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
													v347 = int32(1)
												} else {
													v347 = v286
												}
												return v347
											}
										} else {
											v277 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_4), int32(6))
											mBase = m.M
											v278 = m.ExcPending
											if v278 != 0 {
												return int32(0)
											} else {
												if v277 != 0 {
													v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
													v284 = v279
												} else {
													v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													v282 = v280 - v254
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v282
													v284 = v282
												}
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v284
												v286 = F_slice_del(m, l0)
												mBase = m.M
												v287 = m.ExcPending
												if v287 != 0 {
													return int32(0)
												} else {
													if int32(0) <= v286 {
														v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
														v347 = int32(1)
													} else {
														v347 = v286
													}
													return v347
												}
											}
										}
									} else {
										v290 = v247 - v246
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v252 - v290
										v295 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_5), int32(11))
										mBase = m.M
										v296 = m.ExcPending
										if v296 != 0 {
											return int32(0)
										} else {
											if v295 == int32(0) {
												v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v323 - v290
												v328 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_6), int32(9))
												mBase = m.M
												v329 = m.ExcPending
												if v329 != 0 {
													return int32(0)
												} else {
													if v328 == int32(0) {
														v347 = v2
														return v347
													} else {
														v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														v333 = v332 - v290
														*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v333
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v333
														v336 = F_slice_del(m, l0)
														mBase = m.M
														v337 = m.ExcPending
														if v337 != 0 {
															return int32(0)
														} else {
															if v336 < int32(0) {
																v347 = v336
															} else {
																v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
																v347 = int32(1)
															}
															return v347
														}
													}
												}
											} else {
												v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v299
												v301 = int32(3)
												v303 = int32(0)
												v306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												if v299-v306 < v301 {
													v316 = v303
												} else {
													v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													v312 = F_memcmp(m, v309+v299-v301, int32(_a_F_r_fix_ending_7), v301)
													mBase = m.M
													if v312 != 0 {
														v316 = v303
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v299 - v301
														v316 = int32(1)
													}
												}
												if v316 == int32(0) {
													v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v323 - v290
													v328 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_6), int32(9))
													mBase = m.M
													v329 = m.ExcPending
													if v329 != 0 {
														return int32(0)
													} else {
														if v328 == int32(0) {
															v347 = v2
															return v347
														} else {
															v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															v333 = v332 - v290
															*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v333
															*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v333
															v336 = F_slice_del(m, l0)
															mBase = m.M
															v337 = m.ExcPending
															if v337 != 0 {
																return int32(0)
															} else {
																if v336 < int32(0) {
																	v347 = v336
																} else {
																	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
																	v347 = int32(1)
																}
																return v347
															}
														}
													}
												} else {
													v319 = F_slice_del(m, l0)
													mBase = m.M
													v320 = m.ExcPending
													if v320 != 0 {
														return int32(0)
													} else {
														if int32(0) <= v319 {
															v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
															*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
															v347 = int32(1)
														} else {
															v347 = v319
														}
														return v347
													}
												}
											}
										}
									}
								}
							}
						} else {
							v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v167 + (v105 - v150)
							v173 = F_slice_from_s(m, l0, int32(6), int32(_a_F_r_fix_ending_13))
							mBase = m.M
							v174 = m.ExcPending
							if v174 != 0 {
								return int32(0)
							} else {
								if int32(0) <= v173 {
									v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
									v347 = int32(1)
								} else {
									v347 = v173
								}
								return v347
							}
						}
					}
				case 6:
					v179 = F_slice_from_s(m, l0, int32(3), int32(_a_F_r_fix_ending_14))
					mBase = m.M
					v180 = m.ExcPending
					if v180 != 0 {
						return int32(0)
					} else {
						if int32(0) <= v179 {
							v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
							v347 = int32(1)
						} else {
							v347 = v179
						}
						return v347
					}
				case 7:
					v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v186 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_15), int32(8))
					mBase = m.M
					v187 = m.ExcPending
					if v187 != 0 {
						return int32(0)
					} else {
						if v186 != 0 {
							v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v225
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v225
							v228 = int32(3)
							v230 = int32(0)
							v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v225-v233 < v228 {
								v243 = v230
							} else {
								v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v239 = F_memcmp(m, v236+v225-v228, int32(_a_F_r_fix_ending_1), v228)
								mBase = m.M
								if v239 != 0 {
									v243 = v230
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v225 - v228
									v243 = int32(1)
								}
							}
							if v243 == int32(0) {
								v347 = v2
								return v347
							} else {
								v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v250 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_2), int32(6))
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
											v282 = v273 - v254
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v282
											v284 = v282
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v284
											v286 = F_slice_del(m, l0)
											mBase = m.M
											v287 = m.ExcPending
											if v287 != 0 {
												return int32(0)
											} else {
												if int32(0) <= v286 {
													v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
													v347 = int32(1)
												} else {
													v347 = v286
												}
												return v347
											}
										} else {
											v277 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_4), int32(6))
											mBase = m.M
											v278 = m.ExcPending
											if v278 != 0 {
												return int32(0)
											} else {
												if v277 != 0 {
													v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
													v284 = v279
												} else {
													v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													v282 = v280 - v254
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v282
													v284 = v282
												}
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v284
												v286 = F_slice_del(m, l0)
												mBase = m.M
												v287 = m.ExcPending
												if v287 != 0 {
													return int32(0)
												} else {
													if int32(0) <= v286 {
														v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
														v347 = int32(1)
													} else {
														v347 = v286
													}
													return v347
												}
											}
										}
									} else {
										v290 = v247 - v246
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v252 - v290
										v295 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_5), int32(11))
										mBase = m.M
										v296 = m.ExcPending
										if v296 != 0 {
											return int32(0)
										} else {
											if v295 == int32(0) {
												v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v323 - v290
												v328 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_6), int32(9))
												mBase = m.M
												v329 = m.ExcPending
												if v329 != 0 {
													return int32(0)
												} else {
													if v328 == int32(0) {
														v347 = v2
														return v347
													} else {
														v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														v333 = v332 - v290
														*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v333
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v333
														v336 = F_slice_del(m, l0)
														mBase = m.M
														v337 = m.ExcPending
														if v337 != 0 {
															return int32(0)
														} else {
															if v336 < int32(0) {
																v347 = v336
															} else {
																v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
																v347 = int32(1)
															}
															return v347
														}
													}
												}
											} else {
												v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v299
												v301 = int32(3)
												v303 = int32(0)
												v306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												if v299-v306 < v301 {
													v316 = v303
												} else {
													v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													v312 = F_memcmp(m, v309+v299-v301, int32(_a_F_r_fix_ending_7), v301)
													mBase = m.M
													if v312 != 0 {
														v316 = v303
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v299 - v301
														v316 = int32(1)
													}
												}
												if v316 == int32(0) {
													v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v323 - v290
													v328 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_6), int32(9))
													mBase = m.M
													v329 = m.ExcPending
													if v329 != 0 {
														return int32(0)
													} else {
														if v328 == int32(0) {
															v347 = v2
															return v347
														} else {
															v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															v333 = v332 - v290
															*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v333
															*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v333
															v336 = F_slice_del(m, l0)
															mBase = m.M
															v337 = m.ExcPending
															if v337 != 0 {
																return int32(0)
															} else {
																if v336 < int32(0) {
																	v347 = v336
																} else {
																	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
																	v347 = int32(1)
																}
																return v347
															}
														}
													}
												} else {
													v319 = F_slice_del(m, l0)
													mBase = m.M
													v320 = m.ExcPending
													if v320 != 0 {
														return int32(0)
													} else {
														if int32(0) <= v319 {
															v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
															*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
															v347 = int32(1)
														} else {
															v347 = v319
														}
														return v347
													}
												}
											}
										}
									}
								}
							}
						} else {
							v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v188 + (v105 - v183)
							v192 = F_slice_del(m, l0)
							mBase = m.M
							v193 = m.ExcPending
							if v193 != 0 {
								return int32(0)
							} else {
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
					}
				case 8:
					v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v105-int32(2) <= v196 {
						v219 = F_slice_from_s(m, l0, int32(6), int32(_a_F_r_fix_ending_16))
						mBase = m.M
						v220 = m.ExcPending
						if v220 != 0 {
							return int32(0)
						} else {
							if int32(0) <= v219 {
								v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
								v347 = int32(1)
							} else {
								v347 = v219
							}
							return v347
						}
					} else {
						v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200+v105-int32(1)))))
						switch v204 - int32(136) {
						case 0, 5:
							v209 = F_find_among_b(m, l0, int32(_a_F_r_fix_ending_17), int32(3))
							mBase = m.M
							v210 = m.ExcPending
							if v210 != 0 {
								return int32(0)
							} else {
								switch v209 - int32(1) {
								case 0:
									v213 = F_slice_del(m, l0)
									mBase = m.M
									v214 = m.ExcPending
									if v214 != 0 {
										return int32(0)
									} else {
										if int32(0) <= v213 {
											v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
											v347 = int32(1)
										} else {
											v347 = v213
										}
										return v347
									}
								case 1:
									v219 = F_slice_from_s(m, l0, int32(6), int32(_a_F_r_fix_ending_16))
									mBase = m.M
									v220 = m.ExcPending
									if v220 != 0 {
										return int32(0)
									} else {
										if int32(0) <= v219 {
											v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
											v347 = int32(1)
										} else {
											v347 = v219
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
							v219 = F_slice_from_s(m, l0, int32(6), int32(_a_F_r_fix_ending_16))
							mBase = m.M
							v220 = m.ExcPending
							if v220 != 0 {
								return int32(0)
							} else {
								if int32(0) <= v219 {
									v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344
									v347 = int32(1)
								} else {
									v347 = v219
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
func F_radians(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v7 float64
	_ = v7
	var v9 float64
	_ = v9
	var v17 float64
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	v7 = base.F64_mul(v5, float64(0.017453292519943295))
	v9 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(base.F64_abs(v7), v9)&base.F64_ne(base.F64_abs(v5), v9) == int32(0) {
		v17 = float64(0)
		if base.F64_eq(v7, v17)&base.F64_ne(v5, v17) != 0 {
			F_float_underflow_error(m)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v22 = F_Float8GetDatum(m, v7)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				return v22
			}
		}
	} else {
		F_float_overflow_error(m)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
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
			v60 = F_expression_tree_walker_impl(m, l0, int32(1051), l1)
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
				v60 = F_expression_tree_walker_impl(m, l0, int32(1051), l1)
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
					v60 = F_expression_tree_walker_impl(m, l0, int32(1051), l1)
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
			v52 = F_query_tree_walker_impl(m, l0, int32(1051), l1, int32(0))
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
				v60 = F_expression_tree_walker_impl(m, l0, int32(1051), l1)
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
	var v58 int32
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
										v58 = F_heap_getattr_1(m, v16, v54, v55, l1+int32(8))
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v58
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
								v58 = F_heap_getattr_1(m, v16, v54, v55, l1+int32(8))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v58
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
	var v50 int32
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
							v50 = F_heap_getattr_1(m, v10+v37, v46, v47, l1+v37)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v50
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
					v50 = F_heap_getattr_1(m, v10+v37, v46, v47, l1+v37)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v50
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
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
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
	if v72 != l0+int32(72) {
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
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
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
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
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
	var v87 int32
	_ = v87
	var v88 int64
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
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
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v366 int32
	_ = v366
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	v22 = m.G0
	v24 = v22 - int32(112)
	m.G0 = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v27 = F_pg_detoast_datum(m, v26)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v32 = F_pg_detoast_datum(m, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v38 = F_lookup_rowtype_tupdesc(m, v36, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v43 = F_lookup_rowtype_tupdesc(m, v41, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+108)) = v27
	v48 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+104)) = v48
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+100)) = uint16(v48)
	v52 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+96)) = v52
	v54 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+92)) = int32(base.Ui32(v46) >> (uint(v54) % 32))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+88)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v24)+84)) = v48
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+80)) = uint16(v48)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+76)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v24)+72)) = int32(base.Ui32(v57) >> (uint(v54) % 32))
	if v45 < v40 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v69 = v40
	goto L9
L8:
	;
	v69 = v45
	goto L9
L9:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
	if v71 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v36 != v94 {
		goto L16
	} else {
		goto L17
	}
L11:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v70)+20))
	v82 = F_MemoryContextAlloc(m, v77, v69<<(uint(int32(2))%32)+int32(20))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	if v74 < v69 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	v93 = v71
	v94 = v76
	goto L10
L14:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v84)+16)) = v82
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+16))
	v88 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v87)+4)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v87))) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v87)+12)) = v88
	v93 = v87
	v94 = int32(0)
	goto L10
L15:
	;
	v150 = F_palloc(m, v40<<(uint(int32(2))%32))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L31
	}
L16:
	;
	v103 = v93 + int32(20)
	v107 = v69 << (uint(int32(2)) % 32)
	if v103&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v107)) == int32(0) {
		goto L22
	} else {
		goto L23
	}
L17:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v93)+8))
	if v96 != v37 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	if v98 != v41 {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v93)+16))
	if v100 == v42 {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	goto L16
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+16)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v93)+12)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v93)+8)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v93)+4)) = v36
	goto L15
L22:
	;
	if v107 == int32(0) {
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if v107 == int32(0) {
		goto L21
	} else {
		goto L30
	}
L25:
	;
	v117 = v93 + v107 + int32(20)
	v119 = v93 + int32(24)
	if base.Ui32(v119) < base.Ui32(v117) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v121 = v117
	goto L28
L27:
	;
	v121 = v119
	goto L28
L28:
	;
	v128 = (v121-v93-int32(21))&int32(-4) + int32(4)
	if v128 == int32(0) {
		goto L21
	} else {
		goto L29
	}
L29:
	;
	base.MemoryFill(m, v103, int32(0), v128)
	goto L21
L30:
	;
	base.MemoryFill(m, v103, int32(0), v107)
	goto L21
L31:
	;
	v152 = F_palloc(m, v40)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_heap_deform_tuple(m, v24+int32(92), v38, v150, v152)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v160 = F_palloc(m, v45<<(uint(int32(2))%32))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v162 = F_palloc(m, v45)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_heap_deform_tuple(m, v24+int32(72), v43, v160, v162)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v166 = int32(0)
	v172 = base.B2i32(v166 < v40)
	if v166 < v40 {
		goto L43
	} else {
		goto L44
	}
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L113
	}
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L108
	}
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L1
	} else {
		goto L102
	}
L40:
	;
	F_pfree(m, v150)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L82
	}
L41:
	;
	if base.B2i32(v342 != v40)|base.B2i32(v339 != v45) != 0 {
		goto L37
	} else {
		goto L81
	}
L42:
	;
	v182 = v179
	v184 = base.B2i32(v166 < v45)
	v185 = v166
	v186 = v172
	v187 = v180
	goto L47
L43:
	;
	v173 = int32(0)
	v179 = v173
	v180 = v173
	goto L42
L44:
	;
	goto L45
L45:
	;
	v175 = int32(0)
	if v45 <= v175 {
		v339 = v175
		v342 = v166
		goto L41
	} else {
		goto L46
	}
L46:
	;
	v179 = v175
	v180 = v175
	goto L42
L47:
	;
	if v186&int32(1) == int32(0) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v339 = v327
	v342 = v329
	goto L41
L49:
	;
	v335 = base.B2i32(v327 < v45)
	v336 = base.B2i32(v329 < v40)
	if v335|v336 != 0 {
		v182 = v327
		v184 = v335
		v185 = v329
		v186 = v336
		v187 = v331
		goto L47
	} else {
		goto L80
	}
L50:
	;
	if v184&int32(1) == int32(0) {
		v339 = v182
		v342 = v185
		goto L41
	} else {
		goto L53
	}
L51:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+v206<<(uint(int32(4))%32)+v185*int32(100))+111)))
	if v213 != int32(1) {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v327 = v182
	v329 = v185 + int32(1)
	v331 = v187
	goto L49
L53:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v225 = v43 + v222<<(uint(int32(4))%32)
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225+v182*int32(100))+111)))
	if v229 == int32(1) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v327 = v182 + int32(1)
	v329 = v185
	v331 = v187
	goto L49
L55:
	;
	goto L56
L56:
	;
	if v186&int32(1) == int32(0) {
		v339 = v182
		v342 = v185
		goto L41
	} else {
		goto L57
	}
L57:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v242 = int32(100)
	v244 = v38 + v238<<(uint(int32(4))%32) + v185*v242
	v245 = int32(20)
	v246 = v244 + v245
	v249 = v225 + v182*v242
	v251 = v249 + v245
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v244)+88))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v249)+88))
	if v252 != v253 {
		goto L39
	} else {
		goto L58
	}
L58:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v251)+96))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v246)+96))
	v259 = v93 + int32(20) + v187<<(uint(int32(2))%32)
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
	if v260 != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185+v152))))
	if v272 == int32(1) {
		goto L67
	} else {
		goto L68
	}
L60:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	if v261 == v252 {
		v270 = v260
		goto L59
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v264 = F_lookup_type_cache(m, v252, int32(64))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L64
	}
L63:
	;
	goto L62
L64:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v264)+108))
	if v266 == int32(0) {
		goto L38
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v259))) = v264
	v270 = v264
	goto L59
L66:
	;
	v321 = int32(1)
	v327 = v182 + v321
	v329 = v185 + v321
	v331 = v187 + v321
	goto L49
L67:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182+v162))))
	if v276 != 0 {
		goto L66
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182+v162))))
	if v279 != 0 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v366 = int32(1)
	goto L40
L71:
	;
	v366 = int32(-1)
	goto L40
L72:
	;
	goto L73
L73:
	;
	v281 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+52)) = uint8(v281)
	if v256 == v255 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v285 = v256
	goto L76
L75:
	;
	v285 = v281
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v285
	*(*int64)(unsafe.Add(mBase, uint32(v24)+40)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v270 + int32(104)
	v292 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+54)) = uint16(v292)
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v150+v185<<(uint(v292)%32))))
	v298 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+60)) = uint8(v298)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+56)) = v297
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v160+v182<<(uint(v292)%32))))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+68)) = uint8(v298)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+64)) = v304
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v270)+104))
	v312 = m.T0[v311].(func(*base.Module, int32) int32)(m, v24+int32(36))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	if v312 < int32(0) {
		v366 = int32(-1)
		goto L40
	} else {
		goto L78
	}
L78:
	;
	if v312 == int32(0) {
		goto L66
	} else {
		goto L79
	}
L79:
	;
	v366 = int32(1)
	goto L40
L80:
	;
	goto L48
L81:
	;
	v366 = int32(0)
	goto L40
L82:
	;
	F_pfree(m, v152)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_pfree(m, v160)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	F_pfree(m, v162)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	if int32(0) <= v392 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	F_DecrTupleDescRefCount(m, v38)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	if int32(0) <= v397 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	goto L88
L90:
	;
	F_DecrTupleDescRefCount(m, v43)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v402 != v27 {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	goto L92
L94:
	;
	F_pfree(m, v27)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v406 != v32 {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	goto L96
L98:
	;
	F_pfree(m, v32)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	m.G0 = v24 + int32(112)
	return v366
L101:
	;
	goto L100
L102:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v246)+68))
	v422 = F_format_type_be(m, v421)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v251)+68))
	v425 = F_format_type_be(m, v424)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = v187 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v425
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v422
	F_errmsg(m, int32(_a_F_record_cmp_0), v24+int32(16))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(_a_F_record_cmp_1), int32(952), int32(_a_F_record_cmp_2))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
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
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v264)))
	v450 = F_format_type_be(m, v449)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v450
	F_errmsg(m, int32(_a_F_record_cmp_3), v24)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_record_cmp_1), int32(975), int32(_a_F_record_cmp_2))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
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
	v467 = m.ExcPending
	if v467 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	F_errmsg(m, int32(_a_F_record_cmp_4), int32(0))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_record_cmp_1), int32(1040), int32(_a_F_record_cmp_2))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
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
func F_record_image_eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
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
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int64
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
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
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v311 int32
	_ = v311
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
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
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	v20 = m.G0
	v22 = v20 + int32(-64)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v25 = F_pg_detoast_datum(m, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v30 = F_pg_detoast_datum(m, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v34 = F_lookup_rowtype_tupdesc(m, v32, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
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
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+60)) = v25
	v44 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+56)) = v44
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+52)) = uint16(v44)
	v48 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v48
	v50 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+44)) = int32(base.Ui32(v42) >> (uint(v50) % 32))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = v44
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+32)) = uint16(v44)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = int32(base.Ui32(v53) >> (uint(v50) % 32))
	if v41 < v36 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v65 = v36
	goto L8
L7:
	;
	v65 = v41
	goto L8
L8:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+16))
	if v67 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v32 != v90 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v66)+20))
	v78 = F_MemoryContextAlloc(m, v73, v65<<(uint(int32(2))%32)+int32(20))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	if v70 < v65 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	v89 = v67
	v90 = v72
	goto L9
L13:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v80)+16)) = v78
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+16))
	v84 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v83)+4)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = v65
	*(*int64)(unsafe.Add(mBase, uint32(v83)+12)) = v84
	v89 = v83
	v90 = int32(0)
	goto L9
L14:
	;
	v146 = F_palloc(m, v36<<(uint(int32(2))%32))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L30
	}
L15:
	;
	v99 = v89 + int32(20)
	v103 = v65 << (uint(int32(2)) % 32)
	if v99&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v103)) == int32(0) {
		goto L21
	} else {
		goto L22
	}
L16:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	if v92 != v33 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if v94 != v37 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	if v96 == v38 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L15
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+16)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v89)+12)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v89)+8)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v89)+4)) = v32
	goto L14
L21:
	;
	if v103 == int32(0) {
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	if v103 == int32(0) {
		goto L20
	} else {
		goto L29
	}
L24:
	;
	v113 = v89 + v103 + int32(20)
	v115 = v89 + int32(24)
	if base.Ui32(v115) < base.Ui32(v113) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v117 = v113
	goto L27
L26:
	;
	v117 = v115
	goto L27
L27:
	;
	v124 = (v117-v89-int32(21))&int32(-4) + int32(4)
	if v124 == int32(0) {
		goto L20
	} else {
		goto L28
	}
L28:
	;
	base.MemoryFill(m, v99, int32(0), v124)
	goto L20
L29:
	;
	base.MemoryFill(m, v99, int32(0), v103)
	goto L20
L30:
	;
	v148 = F_palloc(m, v36)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_heap_deform_tuple(m, v20+int32(-20), v34, v146, v148)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v156 = F_palloc(m, v41<<(uint(int32(2))%32))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v158 = F_palloc(m, v41)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_heap_deform_tuple(m, v20+int32(-40), v39, v156, v158)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v162 = int32(0)
	v166 = base.B2i32(v162 < v36)
	if v162 < v36 {
		goto L41
	} else {
		goto L42
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L93
	}
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L87
	}
L38:
	;
	F_pfree(m, v146)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L67
	}
L39:
	;
	if base.B2i32(v286 != v36)|base.B2i32(v285 != v41) != 0 {
		goto L36
	} else {
		goto L66
	}
L40:
	;
	v176 = v173
	v177 = v162
	v179 = base.B2i32(v162 < v41)
	v180 = v166
	v181 = v174
	goto L45
L41:
	;
	v167 = int32(0)
	v173 = v167
	v174 = v167
	goto L40
L42:
	;
	goto L43
L43:
	;
	v169 = int32(0)
	if v41 <= v169 {
		v285 = v169
		v286 = v162
		goto L39
	} else {
		goto L44
	}
L44:
	;
	v173 = v169
	v174 = v169
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
	v285 = v274
	v286 = v275
	goto L39
L47:
	;
	v281 = base.B2i32(v274 < v41)
	v282 = base.B2i32(v275 < v36)
	if v281|v282 != 0 {
		v176 = v274
		v177 = v275
		v179 = v281
		v180 = v282
		v181 = v278
		goto L45
	} else {
		goto L65
	}
L48:
	;
	if v179 == int32(0) {
		v285 = v176
		v286 = v177
		goto L39
	} else {
		goto L51
	}
L49:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34+v196<<(uint(int32(4))%32)+v177*int32(100))+111)))
	if v203 != int32(1) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v274 = v176
	v275 = v177 + int32(1)
	v278 = v181
	goto L47
L51:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v213 = v39 + v210<<(uint(int32(4))%32)
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213+v176*int32(100))+111)))
	if v217 == int32(1) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v274 = v176 + int32(1)
	v275 = v177
	v278 = v181
	goto L47
L53:
	;
	goto L54
L54:
	;
	if v180 == int32(0) {
		v285 = v176
		v286 = v177
		goto L39
	} else {
		goto L55
	}
L55:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v228 = int32(100)
	v230 = v34 + v224<<(uint(int32(4))%32) + v177*v228
	v231 = int32(20)
	v232 = v230 + v231
	v235 = v213 + v176*v228
	v237 = v235 + v231
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v230)+88))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v235)+88))
	if v238 != v239 {
		goto L37
	} else {
		goto L56
	}
L56:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176+v158))))
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177+v148))))
	if v244 == int32(1) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v268 = int32(1)
	v274 = v176 + v268
	v275 = v177 + v268
	v278 = v181 + v268
	goto L47
L58:
	;
	if v242&int32(1) != 0 {
		goto L57
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v250 = int32(0)
	if v242&int32(1) != 0 {
		v311 = v250
		goto L38
	} else {
		goto L62
	}
L61:
	;
	v311 = int32(0)
	goto L38
L62:
	;
	v253 = int32(2)
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v146+v177<<(uint(v253)%32))))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v156+v176<<(uint(v253)%32))))
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232)+82)))
	v262 = int32(*(*int16)(unsafe.Add(mBase, uint32(v237)+72)))
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
		v311 = v250
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
	v311 = int32(1)
	goto L38
L67:
	;
	F_pfree(m, v148)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_pfree(m, v156)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	F_pfree(m, v158)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
	if int32(0) <= v334 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	F_DecrTupleDescRefCount(m, v34)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L1
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	if int32(0) <= v339 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	goto L73
L75:
	;
	F_DecrTupleDescRefCount(m, v39)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v344 != v25 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	goto L77
L79:
	;
	F_pfree(m, v25)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v348 != v30 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	goto L81
L83:
	;
	F_pfree(m, v30)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	m.G0 = v22 - int32(-64)
	return v311
L86:
	;
	goto L85
L87:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v232)+68))
	v364 = F_format_type_be(m, v363)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v237)+68))
	v367 = F_format_type_be(m, v366)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v181 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v367
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v364
	F_errmsg(m, int32(_a_F_record_image_eq_0), v22)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(_a_F_record_image_eq_1), int32(1720), int32(_a_F_record_image_eq_2))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
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
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	F_errmsg(m, int32(_a_F_record_image_eq_3), int32(0))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_record_image_eq_1), int32(1753), int32(_a_F_record_image_eq_2))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
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
func F_record_image_ge(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_record_image_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return int32(base.Ui32(v2^int32(-1)) >> (uint(int32(31)) % 32))
	}
}
func F_record_send(m *base.Module, l0 int32) int32 {
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
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
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
	var v345 int32
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
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
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
	v115 = F_palloc(m, v31<<(uint(int32(2))%32))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
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
	v117 = F_palloc(m, v31)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_heap_deform_tuple(m, v18+int32(28), v29, v115, v117)
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
	F_pfree(m, v115)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L1
	} else {
		goto L64
	}
L30:
	;
	v128 = v31 & int32(3)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v130 = int32(4)
	v132 = v29 + v129<<(uint(v130)%32)
	v133 = int32(0)
	if base.Ui32(v130) <= base.Ui32(v31) {
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
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+111)))
	v158 = int32(1)
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+211)))
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+311)))
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+411)))
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
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132+v197*int32(100))+111)))
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
	v282 = v29 + v276<<(uint(int32(4))%32) + v263*int32(100)
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282)+111)))
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
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v282)+88))
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
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v115+v263<<(uint(int32(2))%32))))
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
	return v430
}
func F_record_smaller(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v4 = F_record_cmp(m, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 < int32(0) {
			v10 = int32(20)
		} else {
			v10 = int32(28)
		}
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0+v10)))
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
func F_recv(m *base.Module, l0 int32, l1 int32) int32 {
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	v5 = int32(0)
	v7 = F_recvfrom(m, l0, l1, int32(_a_F_recv_0), int32(64), v5, v5)
	return v7
}
func F_reduce_outer_joins_pass2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
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
	var v102 int32
	_ = v102
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
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
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v334 int64
	_ = v334
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
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
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v585 int32
	_ = v585
	var v593 int32
	_ = v593
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v637 int32
	_ = v637
	v15 = m.G0
	v17 = v15 - int32(32)
	m.G0 = v17
	if l0 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L10
	} else {
		goto L189
	}
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v19 - int32(63) {
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
	v614 = m.ExcPending
	if v614 != 0 {
		goto L10
	} else {
		goto L186
	}
L5:
	;
	m.G0 = v17 + int32(32)
	return
L6:
	;
	F_bms_free(m, v585)
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L10
	} else {
		goto L185
	}
L7:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+12))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v100 {
	case 0, 4, 5:
		v502 = v100
		v504 = v98
		v505 = v99
		goto L30
	case 1:
		goto L37
	case 2:
		goto L35
	case 3:
		goto L36
	default:
		goto L34
	}
L8:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v36 = F_find_nonnullable_rels(m, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L10
	} else {
		goto L14
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
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
	v29 = m.ExcPending
	if v29 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	F_errfinish(m, int32(_a_F_reduce_outer_joins_pass2_1), int32(3267), int32(_a_F_reduce_outer_joins_pass2_2))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
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
	v38 = F_bms_add_members(m, v36, l4)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v41 = F_find_forced_null_vars(m, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	v43 = F_mbms_add_members(m, v41, l5)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v48 = int32(0)
	goto L18
L18:
	;
	v62 = int32(0)
	if v46 == v62 {
		v72 = v62
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if v45 == int32(0) {
		v585 = v38
		goto L6
	} else {
		goto L23
	}
L21:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v66 <= v48 {
		v72 = int32(0)
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	v72 = v68 + v48<<(uint(int32(2))%32)
	goto L20
L23:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if base.B2i32(v72 == int32(0))|base.B2i32(v77 <= v48) != 0 {
		v585 = v38
		goto L6
	} else {
		goto L24
	}
L24:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	if v80 == int32(0) {
		v585 = v38
		goto L6
	} else {
		goto L25
	}
L25:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v80+v48<<(uint(int32(2))%32))))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+4)))
	if v87 == int32(1) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	F_reduce_outer_joins_pass2(m, v90, v86, l2, l3, v38, v43)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L10
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v48 = v48 + int32(1)
	goto L18
L29:
	;
	goto L28
L30:
	;
	if v95 == int32(0) {
		goto L151
	} else {
		goto L152
	}
L31:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v348 = F_find_nonnullable_vars_walker(m, v346, int32(1))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L10
	} else {
		goto L106
	}
L32:
	;
	v334 = *(*int64)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = base.I64_rotl(v334, int64(32))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v338)+12))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v339)+4))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v339)))
	v344 = v341
	v345 = v340
	goto L31
L33:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v324 = F_palloc(m, int32(8))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L10
	} else {
		goto L104
	}
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L10
	} else {
		goto L101
	}
L35:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v202 = int32(0)
	if base.B2i32(l4 == v202)|base.B2i32(v201 == v202) != 0 {
		v247 = v202
		goto L67
	} else {
		goto L68
	}
L36:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v152 = int32(0)
	if base.B2i32(l4 == v152)|base.B2i32(v151 == v152) != 0 {
		v197 = v152
		goto L53
	} else {
		goto L54
	}
L37:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v102 = int32(0)
	if base.B2i32(l4 == v102)|base.B2i32(v101 == v102) != 0 {
		v147 = v102
		goto L39
	} else {
		goto L40
	}
L38:
	;
	if v147 == int32(0) {
		v344 = v98
		v345 = v99
		goto L31
	} else {
		goto L51
	}
L39:
	;
	goto L38
L40:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	if v112 < v113 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v115 = v112
	goto L43
L42:
	;
	v115 = v113
	goto L43
L43:
	;
	if v115 <= int32(1) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v118 = int32(1)
	goto L46
L45:
	;
	v118 = v115
	goto L46
L46:
	;
	v119 = int32(8)
	v124 = int32(0)
	goto L47
L47:
	;
	v131 = v124 << (uint(int32(2)) % 32)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v101+v119+v131)))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l4+v119+v131)))
	v136 = v133 & v135
	v138 = base.B2i32(v136 != int32(0))
	if v136 != 0 {
		v147 = v138
		goto L39
	} else {
		goto L49
	}
L48:
	;
	v147 = v138
	goto L39
L49:
	;
	v140 = v124 + int32(1)
	if v140 != v118 {
		v124 = v140
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v502 = int32(0)
	v504 = v98
	v505 = v99
	goto L30
L52:
	;
	if v197 == int32(0) {
		goto L32
	} else {
		goto L65
	}
L53:
	;
	goto L52
L54:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v151)+4))
	if v162 < v163 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v165 = v162
	goto L57
L56:
	;
	v165 = v163
	goto L57
L57:
	;
	if v165 <= int32(1) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v168 = int32(1)
	goto L60
L59:
	;
	v168 = v165
	goto L60
L60:
	;
	v169 = int32(8)
	v174 = int32(0)
	goto L61
L61:
	;
	v181 = v174 << (uint(int32(2)) % 32)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v151+v169+v181)))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l4+v169+v181)))
	v186 = v183 & v185
	v188 = base.B2i32(v186 != int32(0))
	if v186 != 0 {
		v197 = v188
		goto L53
	} else {
		goto L63
	}
L62:
	;
	v197 = v188
	goto L53
L63:
	;
	v190 = v174 + int32(1)
	if v190 != v168 {
		v174 = v190
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	v502 = int32(0)
	v504 = v98
	v505 = v99
	goto L30
L66:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v249 = int32(0)
	if base.B2i32(l4 == v249)|base.B2i32(v248 == v249) != 0 {
		v294 = v249
		goto L80
	} else {
		goto L81
	}
L67:
	;
	goto L66
L68:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v201)+4))
	if v212 < v213 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v215 = v212
	goto L71
L70:
	;
	v215 = v213
	goto L71
L71:
	;
	if v215 <= int32(1) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v218 = int32(1)
	goto L74
L73:
	;
	v218 = v215
	goto L74
L74:
	;
	v219 = int32(8)
	v224 = int32(0)
	goto L75
L75:
	;
	v231 = v224 << (uint(int32(2)) % 32)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v201+v219+v231)))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l4+v219+v231)))
	v236 = v233 & v235
	v238 = base.B2i32(v236 != int32(0))
	if v236 != 0 {
		v247 = v238
		goto L67
	} else {
		goto L77
	}
L76:
	;
	v247 = v238
	goto L67
L77:
	;
	v240 = v224 + int32(1)
	if v240 != v218 {
		v224 = v240
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	if v247 != 0 {
		goto L92
	} else {
		goto L93
	}
L80:
	;
	goto L79
L81:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v248)+4))
	if v259 < v260 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v262 = v259
	goto L84
L83:
	;
	v262 = v260
	goto L84
L84:
	;
	if v262 <= int32(1) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v265 = int32(1)
	goto L87
L86:
	;
	v265 = v262
	goto L87
L87:
	;
	v266 = int32(8)
	v271 = int32(0)
	goto L88
L88:
	;
	v278 = v271 << (uint(int32(2)) % 32)
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v248+v266+v278)))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l4+v266+v278)))
	v283 = v280 & v282
	v285 = base.B2i32(v283 != int32(0))
	if v283 != 0 {
		v294 = v285
		goto L80
	} else {
		goto L90
	}
L89:
	;
	v294 = v285
	goto L80
L90:
	;
	v287 = v271 + int32(1)
	if v287 != v265 {
		v271 = v287
		goto L88
	} else {
		goto L91
	}
L91:
	;
	goto L89
L92:
	;
	if v294 != 0 {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	goto L94
L94:
	;
	if v294 != 0 {
		goto L33
	} else {
		goto L100
	}
L95:
	;
	v502 = int32(0)
	v504 = v98
	v505 = v99
	goto L30
L96:
	;
	goto L97
L97:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v298 = F_palloc(m, int32(8))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L10
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v298)+4)) = v296
	*(*int32)(unsafe.Add(mBase, uint32(v298))) = v95
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v303 = F_lappend(m, v302, v298)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L10
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v303
	v344 = v98
	v345 = v99
	goto L31
L100:
	;
	v502 = int32(2)
	v504 = v98
	v505 = v99
	goto L30
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v100
	F_errmsg_internal(m, int32(_a_F_reduce_outer_joins_pass2_3), v17+int32(16))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L10
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(_a_F_reduce_outer_joins_pass2_1), int32(3355), int32(_a_F_reduce_outer_joins_pass2_2))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v324)+4)) = v322
	*(*int32)(unsafe.Add(mBase, uint32(v324))) = v95
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v329 = F_lappend(m, v328, v324)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L10
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v329
	goto L32
L106:
	;
	v350 = int32(0)
	v353 = v350
	v358 = v350
	goto L107
L107:
	;
	v366 = int32(0)
	if v348 == v366 {
		v376 = v366
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v344)))
	v449 = int32(0)
	if base.B2i32(v444 == v449)|base.B2i32(v448 == v449) != 0 {
		v494 = v449
		goto L136
	} else {
		goto L137
	}
L109:
	;
	if l5 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L110:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v348)+4))
	if v370 <= v353 {
		v376 = int32(0)
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v348)+12))
	v376 = v372 + v353<<(uint(int32(2))%32)
	goto L109
L112:
	;
	goto L108
L113:
	;
	v444 = int32(0)
	goto L112
L114:
	;
	goto L115
L115:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if base.B2i32(v376 == int32(0))|base.B2i32(v382 <= v353) != 0 {
		v444 = v358
		goto L112
	} else {
		goto L116
	}
L116:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	if v385 == int32(0) {
		v444 = v358
		goto L112
	} else {
		goto L117
	}
L117:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v376)))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v385+v353<<(uint(int32(2))%32))))
	v393 = int32(0)
	if base.B2i32(v388 == v393)|base.B2i32(v392 == v393) != 0 {
		v438 = v393
		goto L119
	} else {
		goto L120
	}
L118:
	;
	if v438 != 0 {
		goto L131
	} else {
		goto L132
	}
L119:
	;
	goto L118
L120:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v388)+4))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v392)+4))
	if v403 < v404 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v406 = v403
	goto L123
L122:
	;
	v406 = v404
	goto L123
L123:
	;
	if v406 <= int32(1) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v409 = int32(1)
	goto L126
L125:
	;
	v409 = v406
	goto L126
L126:
	;
	v410 = int32(8)
	v415 = int32(0)
	goto L127
L127:
	;
	v422 = v415 << (uint(int32(2)) % 32)
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v392+v410+v422)))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v388+v410+v422)))
	v427 = v424 & v426
	v429 = base.B2i32(v427 != int32(0))
	if v427 != 0 {
		v438 = v429
		goto L119
	} else {
		goto L129
	}
L128:
	;
	v438 = v429
	goto L119
L129:
	;
	v431 = v415 + int32(1)
	if v431 != v409 {
		v415 = v431
		goto L127
	} else {
		goto L130
	}
L130:
	;
	goto L128
L131:
	;
	v439 = F_bms_add_member(m, v358, v353)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L10
	} else {
		goto L134
	}
L132:
	;
	v441 = v358
	goto L133
L133:
	;
	v353 = v353 + int32(1)
	v358 = v441
	goto L107
L134:
	;
	v441 = v439
	goto L133
L135:
	;
	if v494 != 0 {
		goto L148
	} else {
		goto L149
	}
L136:
	;
	goto L135
L137:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v444)+4))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v448)+4))
	if v459 < v460 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v462 = v459
	goto L140
L139:
	;
	v462 = v460
	goto L140
L140:
	;
	if v462 <= int32(1) {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v465 = int32(1)
	goto L143
L142:
	;
	v465 = v462
	goto L143
L143:
	;
	v466 = int32(8)
	v471 = int32(0)
	goto L144
L144:
	;
	v478 = v471 << (uint(int32(2)) % 32)
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v448+v466+v478)))
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v444+v466+v478)))
	v483 = v480 & v482
	v485 = base.B2i32(v483 != int32(0))
	if v483 != 0 {
		v494 = v485
		goto L136
	} else {
		goto L146
	}
L145:
	;
	v494 = v485
	goto L136
L146:
	;
	v487 = v471 + int32(1)
	if v487 != v465 {
		v471 = v487
		goto L144
	} else {
		goto L147
	}
L147:
	;
	goto L145
L148:
	;
	v495 = int32(5)
	goto L150
L149:
	;
	v495 = int32(1)
	goto L150
L150:
	;
	v502 = v495
	v504 = v344
	v505 = v345
	goto L30
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v502
	v529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v505)+4)))
	if v529 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L152:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v502 == v512 {
		goto L151
	} else {
		goto L153
	}
L153:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v514)+52))
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v515)+12))
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v516+v95<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v522)+44)) = v502
	if v502 != 0 {
		goto L151
	} else {
		goto L154
	}
L154:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v525 = F_bms_add_member(m, v524, v95)
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L10
	} else {
		goto L155
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v525
	goto L151
L156:
	;
	v532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v504)+4)))
	if v532 != int32(1) {
		goto L5
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	v535 = int32(0)
	if v502 == int32(2) {
		v552 = v535
		v553 = v535
		goto L160
	} else {
		goto L161
	}
L159:
	;
	goto L158
L160:
	;
	v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v505)+4)))
	if v554 == int32(1) {
		goto L167
	} else {
		goto L168
	}
L161:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v540 = F_find_nonnullable_rels(m, v539)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L10
	} else {
		goto L162
	}
L162:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v543 = F_find_forced_null_vars(m, v542)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L10
	} else {
		goto L163
	}
L163:
	;
	if v502&int32(-5) != 0 {
		v552 = v540
		v553 = v543
		goto L160
	} else {
		goto L164
	}
L164:
	;
	v547 = F_bms_add_members(m, v540, l4)
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L10
	} else {
		goto L165
	}
L165:
	;
	v549 = F_mbms_add_members(m, v543, l5)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L10
	} else {
		goto L166
	}
L166:
	;
	v552 = v547
	v553 = v549
	goto L160
L167:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v560 = base.B2i32(v502 == int32(2))
	if v502 == int32(2) {
		goto L170
	} else {
		goto L171
	}
L168:
	;
	goto L169
L169:
	;
	v572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v504)+4)))
	if v572 != int32(1) {
		v585 = v552
		goto L6
	} else {
		goto L183
	}
L170:
	;
	v561 = int32(0)
	goto L172
L171:
	;
	v561 = l4
	goto L172
L172:
	;
	v563 = v502 & int32(-5)
	if v563 != 0 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v564 = v561
	goto L175
L174:
	;
	v564 = v552
	goto L175
L175:
	;
	if v502 == int32(2) {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v566 = int32(0)
	goto L178
L177:
	;
	v566 = l5
	goto L178
L178:
	;
	if v563 != 0 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v567 = v566
	goto L181
L180:
	;
	v567 = v553
	goto L181
L181:
	;
	F_reduce_outer_joins_pass2(m, v557, v505, l2, l3, v564, v567)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L10
	} else {
		goto L182
	}
L182:
	;
	goto L169
L183:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	F_reduce_outer_joins_pass2(m, v575, v504, l2, l3, v552, v553)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L10
	} else {
		goto L184
	}
L184:
	;
	v585 = v552
	goto L6
L185:
	;
	goto L5
L186:
	;
	F_errmsg_internal(m, int32(_a_F_reduce_outer_joins_pass2_4), int32(0))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L10
	} else {
		goto L187
	}
L187:
	;
	F_errfinish(m, int32(_a_F_reduce_outer_joins_pass2_1), int32(3265), int32(_a_F_reduce_outer_joins_pass2_2))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L10
	} else {
		goto L188
	}
L188:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L189:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v628
	F_errmsg_internal(m, int32(_a_F_reduce_outer_joins_pass2_5), v17)
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L10
	} else {
		goto L190
	}
L190:
	;
	F_errfinish(m, int32(_a_F_reduce_outer_joins_pass2_1), int32(3522), int32(_a_F_reduce_outer_joins_pass2_2))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L10
	} else {
		goto L191
	}
L191:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_regcollationin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
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
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
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
	var v155 int32
	_ = v155
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v12 == int32(45) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L29
	} else {
		goto L45
	}
L2:
	;
	m.G0 = v8 + int32(16)
	return v155
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
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v15 != 0 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if base.Ui32(int32(9)) < base.Ui32((v12-int32(48))&int32(255)) {
		goto L3
	} else {
		goto L8
	}
L7:
	;
	v155 = int32(0)
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
	v107 = F_strlen(m, v11)
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
	v46 = v11
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
	v106 = v46 - v11
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
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v73 == int32(0) {
		v96 = v11
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v106 = v96 - v11
	goto L9
L23:
	;
	v77 = v11
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
	v113 = F_DirectInputFunctionCallSafe(m, int32(547), v11, int32(-1), v10, v8+int32(12))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	return int32(0)
L30:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v155 = v117
	goto L2
L31:
	;
	v122 = F_stringToQualifiedNameList(m, v11, v10)
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
	v155 = int32(0)
	goto L2
L34:
	;
	goto L35
L35:
	;
	v130 = F_get_collation_oid(m, v122, int32(1))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L29
	} else {
		goto L36
	}
L36:
	;
	if v130 != 0 {
		v155 = v130
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v132 = int32(0)
	v133 = F_errsave_start(m, v10)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L29
	} else {
		goto L38
	}
L38:
	;
	if v133 == int32(0) {
		v155 = v132
		goto L2
	} else {
		goto L39
	}
L39:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L29
	} else {
		goto L40
	}
L40:
	;
	v140 = F_NameListToString(m, v122)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L29
	} else {
		goto L41
	}
L41:
	;
	v143 = *(*int32)(unsafe.Add(mBase, _c_F_regcollationin[3]))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v140
	F_errmsg(m, int32(_a_F_regcollationin_1), v8)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L29
	} else {
		goto L43
	}
L43:
	;
	F_errsave_finish(m, v10, int32(_a_F_regcollationin_2), int32(1057), int32(_a_F_regcollationin_3))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L29
	} else {
		goto L44
	}
L44:
	;
	v155 = v132
	goto L2
L45:
	;
	F_errmsg_internal(m, int32(_a_F_regcollationin_4), int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L29
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_regcollationin_2), int32(1041), int32(_a_F_regcollationin_3))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L29
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_regconfigout(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
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
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v10 == int32(0) {
		v14 = F_pstrdup(m, int32(_a_F_regconfigout_0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v45 = v14
			m.G0 = v8 + int32(16)
			return v45
		}
	} else {
		v19 = F_SearchSysCache1(m, int32(74), v10)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			if v19 != 0 {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+22)))
				v23 = v21 + v22
				v26 = F_TSConfigIsVisible(m, v10)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					if v26 != 0 {
						v32 = int32(0)
						v33 = F_quote_qualified_identifier(m, v32, v23+int32(4))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							F_ReleaseCatCache(m, v19)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								v45 = v33
								m.G0 = v8 + int32(16)
								return v45
							}
						}
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v23)+68))
						v30 = F_get_namespace_name(m, v29)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							v32 = v30
							v33 = F_quote_qualified_identifier(m, v32, v23+int32(4))
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								F_ReleaseCatCache(m, v19)
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return int32(0)
								} else {
									v45 = v33
									m.G0 = v8 + int32(16)
									return v45
								}
							}
						}
					}
				}
			} else {
				v38 = F_palloc(m, int32(64))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
					v43 = F_pg_snprintf(m, v38, int32(64), int32(_a_F_regconfigout_1), v8)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						v45 = v38
						m.G0 = v8 + int32(16)
						return v45
					}
				}
			}
		}
	}
}
func F_regnamespacein(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
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
	var v117 int32
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
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
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
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v12 == int32(45) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L29
	} else {
		goto L51
	}
L2:
	;
	m.G0 = v8 + int32(16)
	return v175
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
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v15 != 0 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if base.Ui32(int32(9)) < base.Ui32((v12-int32(48))&int32(255)) {
		goto L3
	} else {
		goto L8
	}
L7:
	;
	v175 = int32(0)
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
	v107 = F_strlen(m, v11)
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
	v46 = v11
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
	v106 = v46 - v11
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
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v73 == int32(0) {
		v96 = v11
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v106 = v96 - v11
	goto L9
L23:
	;
	v77 = v11
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
	v113 = F_DirectInputFunctionCallSafe(m, int32(547), v11, int32(-1), v10, v8+int32(12))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	return int32(0)
L30:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v175 = v117
	goto L2
L31:
	;
	v122 = F_stringToQualifiedNameList(m, v11, v10)
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
	v175 = int32(0)
	goto L2
L34:
	;
	goto L35
L35:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	if v129 != int32(1) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v132 = int32(0)
	v133 = F_errsave_start(m, v10)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L29
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+4))
	v153 = F_get_namespace_oid(m, v151, int32(1))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L29
	} else {
		goto L44
	}
L39:
	;
	if v133 == int32(0) {
		v175 = v132
		goto L2
	} else {
		goto L40
	}
L40:
	;
	F_errcode(m, int32(33579140))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L29
	} else {
		goto L41
	}
L41:
	;
	F_errmsg(m, int32(_a_F_regnamespacein_1), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L29
	} else {
		goto L42
	}
L42:
	;
	F_errsave_finish(m, v10, int32(_a_F_regnamespacein_2), int32(1681), int32(_a_F_regnamespacein_3))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L29
	} else {
		goto L43
	}
L43:
	;
	v175 = v132
	goto L2
L44:
	;
	if v153 != 0 {
		v175 = v153
		goto L2
	} else {
		goto L45
	}
L45:
	;
	v155 = int32(0)
	v156 = F_errsave_start(m, v10)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L29
	} else {
		goto L46
	}
L46:
	;
	if v156 == int32(0) {
		v175 = v155
		goto L2
	} else {
		goto L47
	}
L47:
	;
	F_errcode(m, int32(1411))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L29
	} else {
		goto L48
	}
L48:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v165
	F_errmsg(m, int32(_a_F_regnamespacein_4), v8)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L29
	} else {
		goto L49
	}
L49:
	;
	F_errsave_finish(m, v10, int32(_a_F_regnamespacein_2), int32(1689), int32(_a_F_regnamespacein_3))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L29
	} else {
		goto L50
	}
L50:
	;
	v175 = v155
	goto L2
L51:
	;
	F_errmsg_internal(m, int32(_a_F_regnamespacein_5), int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L29
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_regnamespacein_2), int32(1671), int32(_a_F_regnamespacein_3))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L29
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_regprocedurein(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v117 int32
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
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
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
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(432)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	v257 = m.ExcPending
	if v257 != 0 {
		goto L29
	} else {
		goto L68
	}
L2:
	;
	m.G0 = v9 + int32(432)
	return v249
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
	v249 = v2
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
	v113 = F_DirectInputFunctionCallSafe(m, int32(547), v12, int32(-1), v11, v9+int32(16))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	return int32(0)
L30:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	v249 = v117
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
	v249 = v2
	goto L2
L34:
	;
	goto L35
L35:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v9)+428))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v9)+424))
	v137 = int32(0)
	v142 = F_FuncnameGetCandidates(m, v135, v136, v137, v137, v137, v137, int32(1))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L29
	} else {
		goto L37
	}
L36:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v146)+8))
	v249 = v243
	goto L2
L37:
	;
	if v142 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v145 = v136 << (uint(int32(2)) % 32)
	v146 = v142
	goto L41
L39:
	;
	goto L40
L40:
	;
	v227 = F_errsave_start(m, v11)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L29
	} else {
		goto L63
	}
L41:
	;
	v153 = v146 + int32(32)
	v155 = v9 + int32(16)
	if base.Ui32(int32(4)) <= base.Ui32(v145) {
		goto L46
	} else {
		goto L47
	}
L42:
	;
	goto L40
L43:
	;
	if v217 == int32(0) {
		goto L36
	} else {
		goto L61
	}
L44:
	;
	v217 = int32(0)
	goto L43
L45:
	;
	v191 = v186
	v192 = v187
	v193 = v188
	goto L55
L46:
	;
	if (v153|v155)&int32(3) != 0 {
		v186 = v153
		v187 = v155
		v188 = v145
		goto L45
	} else {
		goto L49
	}
L47:
	;
	v179 = v153
	v180 = v155
	v181 = v145
	goto L48
L48:
	;
	if v181 == int32(0) {
		goto L44
	} else {
		goto L54
	}
L49:
	;
	v163 = v153
	v164 = v155
	v165 = v145
	goto L50
L50:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
	if v168 != v169 {
		v186 = v163
		v187 = v164
		v188 = v165
		goto L45
	} else {
		goto L52
	}
L51:
	;
	v179 = v174
	v180 = v172
	v181 = v176
	goto L48
L52:
	;
	v171 = int32(4)
	v172 = v164 + v171
	v174 = v163 + v171
	v176 = v165 - v171
	if base.Ui32(int32(3)) < base.Ui32(v176) {
		v163 = v174
		v164 = v172
		v165 = v176
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v186 = v179
	v187 = v180
	v188 = v181
	goto L45
L55:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191))))
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	if v196 == v197 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v217 = v196 - v197
	goto L43
L57:
	;
	v199 = int32(1)
	v204 = v193 - v199
	if v204 != 0 {
		v191 = v191 + v199
		v192 = v192 + v199
		v193 = v204
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
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	if v220 != 0 {
		v146 = v220
		goto L41
	} else {
		goto L62
	}
L62:
	;
	goto L42
L63:
	;
	if v227 == int32(0) {
		v249 = v2
		goto L2
	} else {
		goto L64
	}
L64:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L29
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v12
	F_errmsg(m, int32(_a_F_regprocedurein_1), v9)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L29
	} else {
		goto L66
	}
L66:
	;
	F_errsave_finish(m, v11, int32(_a_F_regprocedurein_2), int32(265), int32(_a_F_regprocedurein_3))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L29
	} else {
		goto L67
	}
L67:
	;
	v249 = v2
	goto L2
L68:
	;
	F_errmsg_internal(m, int32(_a_F_regprocedurein_4), int32(0))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L29
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_regprocedurein_2), int32(240), int32(_a_F_regprocedurein_3))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
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
func F_regprocsend(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_date_send(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
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
	var v59 int32
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
				v24 = int32(4)
				v26 = l1 + v17<<(uint(v24)%32)
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
				*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v27
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v29
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
				*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v31
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v26)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v33
				v36 = v17 + v24
				v38 = v22 + v24
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
					v58 = l1 + v49<<(uint(int32(4))%32)
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
					*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v59
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
				v58 = l1 + v49<<(uint(int32(4))%32)
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
				*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v59
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
						F_errfinish(m, int32(_a_F_renameatt_check_1), int32(3802), int32(_a_F_renameatt_check_2))
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
								v56 = int32(41)
								v58 = v56
							case 10:
								v58 = int32(37)
							case 29:
								v58 = int32(18)
							case 36:
								v58 = int32(23)
							case 45:
								v58 = int32(51)
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
													F_errfinish(m, int32(_a_F_renameatt_check_1), int32(3835), int32(_a_F_renameatt_check_2))
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
											F_errfinish(m, int32(_a_F_renameatt_check_1), int32(3835), int32(_a_F_renameatt_check_2))
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
								F_errfinish(m, int32(_a_F_renameatt_check_1), int32(3823), int32(_a_F_renameatt_check_2))
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
							v56 = int32(41)
							v58 = v56
						case 10:
							v58 = int32(37)
						case 29:
							v58 = int32(18)
						case 36:
							v58 = int32(23)
						case 45:
							v58 = int32(51)
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
												F_errfinish(m, int32(_a_F_renameatt_check_1), int32(3835), int32(_a_F_renameatt_check_2))
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
										F_errfinish(m, int32(_a_F_renameatt_check_1), int32(3835), int32(_a_F_renameatt_check_2))
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
							F_errfinish(m, int32(_a_F_renameatt_check_1), int32(3823), int32(_a_F_renameatt_check_2))
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
func F_replace_s(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
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
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v12 == int32(0) {
		v16 = F_palloc(m, int32(10))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			if v16 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
				return int32(-1)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v16))) = int64(1)
				v29 = v16 + int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v29
				v31 = v29
				v33 = l1 - l2 + l3
				if v33 == int32(0) {
					v82 = v31
					if l3 != 0 {
						base.MemoryCopy(m, l1+v82, l4, l3)
					} else {
					}
					if l5 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(l5))) = v33
					} else {
					}
					return int32(0)
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v31-int32(4))))
					v39 = v38 + v33
					v41 = v31 - int32(8)
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
					if v42 < v39 {
						v46 = F_repalloc(m, v41, v39+int32(29))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							if v46 == int32(0) {
								F_pfree(m, v41)
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
									return int32(-1)
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v46))) = v39 + int32(20)
								v60 = v46 + int32(8)
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v60
								v62 = v60
								v63 = v38 - l2
								if v63 != 0 {
									v64 = l2 + v62
									base.MemoryCopy(m, v64+v33, v64, v63)
								} else {
								}
								v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v68-int32(4)))) = v39
								v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v72 + v33
								v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								if l2 <= v75 {
									v79 = v75 + v33
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v79
									v82 = v68
								} else {
									if v75 <= l1 {
										v82 = v68
									} else {
										v79 = l1
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v79
										v82 = v68
									}
								}
								if l3 != 0 {
									base.MemoryCopy(m, l1+v82, l4, l3)
								} else {
								}
								if l5 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(l5))) = v33
								} else {
								}
								return int32(0)
							}
						}
					} else {
						v62 = v31
						v63 = v38 - l2
						if v63 != 0 {
							v64 = l2 + v62
							base.MemoryCopy(m, v64+v33, v64, v63)
						} else {
						}
						v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v68-int32(4)))) = v39
						v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v72 + v33
						v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						if l2 <= v75 {
							v79 = v75 + v33
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v79
							v82 = v68
						} else {
							if v75 <= l1 {
								v82 = v68
							} else {
								v79 = l1
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v79
								v82 = v68
							}
						}
						if l3 != 0 {
							base.MemoryCopy(m, l1+v82, l4, l3)
						} else {
						}
						if l5 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(l5))) = v33
						} else {
						}
						return int32(0)
					}
				}
			}
		}
	} else {
		v31 = v12
		v33 = l1 - l2 + l3
		if v33 == int32(0) {
			v82 = v31
			if l3 != 0 {
				base.MemoryCopy(m, l1+v82, l4, l3)
			} else {
			}
			if l5 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(l5))) = v33
			} else {
			}
			return int32(0)
		} else {
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v31-int32(4))))
			v39 = v38 + v33
			v41 = v31 - int32(8)
			v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
			if v42 < v39 {
				v46 = F_repalloc(m, v41, v39+int32(29))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					if v46 == int32(0) {
						F_pfree(m, v41)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
							return int32(-1)
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v46))) = v39 + int32(20)
						v60 = v46 + int32(8)
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v60
						v62 = v60
						v63 = v38 - l2
						if v63 != 0 {
							v64 = l2 + v62
							base.MemoryCopy(m, v64+v33, v64, v63)
						} else {
						}
						v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v68-int32(4)))) = v39
						v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v72 + v33
						v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						if l2 <= v75 {
							v79 = v75 + v33
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v79
							v82 = v68
						} else {
							if v75 <= l1 {
								v82 = v68
							} else {
								v79 = l1
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v79
								v82 = v68
							}
						}
						if l3 != 0 {
							base.MemoryCopy(m, l1+v82, l4, l3)
						} else {
						}
						if l5 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(l5))) = v33
						} else {
						}
						return int32(0)
					}
				}
			} else {
				v62 = v31
				v63 = v38 - l2
				if v63 != 0 {
					v64 = l2 + v62
					base.MemoryCopy(m, v64+v33, v64, v63)
				} else {
				}
				v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(v68-int32(4)))) = v39
				v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v72 + v33
				v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if l2 <= v75 {
					v79 = v75 + v33
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v79
					v82 = v68
				} else {
					if v75 <= l1 {
						v82 = v68
					} else {
						v79 = l1
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v79
						v82 = v68
					}
				}
				if l3 != 0 {
					base.MemoryCopy(m, l1+v82, l4, l3)
				} else {
				}
				if l5 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(l5))) = v33
				} else {
				}
				return int32(0)
			}
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
									F_errfinish(m, int32(_a_F_resolve_anyelement_from_others_2), int32(602), int32(_a_F_resolve_anyelement_from_others_3))
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
										F_errfinish(m, int32(_a_F_resolve_anyelement_from_others_2), int32(616), int32(_a_F_resolve_anyelement_from_others_3))
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
											F_errfinish(m, int32(_a_F_resolve_anyelement_from_others_2), int32(634), int32(_a_F_resolve_anyelement_from_others_3))
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
														F_errfinish(m, int32(_a_F_resolve_anyelement_from_others_2), int32(644), int32(_a_F_resolve_anyelement_from_others_3))
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
						F_errfinish(m, int32(_a_F_resolve_anyelement_from_others_2), int32(648), int32(_a_F_resolve_anyelement_from_others_3))
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
