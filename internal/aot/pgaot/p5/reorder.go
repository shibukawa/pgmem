package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ReorderBufferAbort(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
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
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v100 int32
	_ = v100
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l1
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v15 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v42)+72)) = l3
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v45&int32(16) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v26 = F_hash_search(m, v20, v12+int32(12), int32(0), v12+int32(11))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	if l1 != v15 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v19 != 0 {
		v42 = v19
		goto L2
	} else {
		goto L6
	}
L6:
	;
	goto L1
L7:
	;
	return
L8:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+11)))
	if v28 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v31
	goto L1
L10:
	;
	goto L11
L11:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v35
	if v36 == int32(0) {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v42 = v36
	goto L2
L13:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v42)+24)) = l2
	F_ReorderBufferCleanupTXN(m, l0, v42)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L7
	} else {
		goto L29
	}
L14:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	m.T0[v50].(func(*base.Module, int32, int32, int64))(m, l0, v42, l2)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v42)+172))
	if v53 == int32(0) {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v42)+176))
	v58 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+24))
	v61 = base.B2i32(v59 != int32(0))
	goto L17
L17:
	;
	if v59 != int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_BeginInternalSubTransaction(m, int32(24803))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L7
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v69 = int32(0)
	goto L23
L21:
	;
	F_AbortCurrentTransaction(m)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	F_LocalExecuteInvalidationMessage(m, v56+v69<<(uint(int32(4))%32))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L7
	} else {
		goto L25
	}
L24:
	;
	if v61 == int32(0) {
		goto L13
	} else {
		goto L27
	}
L25:
	;
	v83 = v69 + int32(1)
	if v83 != v53 {
		v69 = v83
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	F_RollbackAndReleaseCurrentSubTransaction(m)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L7
	} else {
		goto L28
	}
L28:
	;
	goto L13
L29:
	;
	goto L1
}
func F_ReorderBufferCheckAndTruncateAbortedTXN(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
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
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, _consts[522]))
	if v6 == int32(1) {
		v42 = v3
		return v42
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		if v9&int32(1024) != 0 {
			v42 = v3
			return v42
		} else {
			if v9&int32(2048) != 0 {
				v42 = int32(1)
				return v42
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v17 = F_TransactionIdIsInProgress(m, v16)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					if v17 != 0 {
						v42 = int32(0)
						return v42
					} else {
						v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						v22 = F_TransactionIdDidCommit(m, v21)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int32(0)
						} else {
							v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							if v22 != 0 {
								v38 = v24 | int32(1024)
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = v38
								v42 = v22 ^ int32(1)
								return v42
							} else {
								F_ReorderBufferTruncateTXN(m, l0, l1, int32(base.Ui32(v24&int32(64))>>(uint(int32(6))%32)))
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return int32(0)
								} else {
									F_ReorderBufferToastReset(m, l0, l1)
									mBase = m.M
									v34 = m.ExcPending
									if v34 != 0 {
										return int32(0)
									} else {
										v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										v38 = v35 | int32(2048)
										*(*int32)(unsafe.Add(mBase, uint32(l1))) = v38
										v42 = v22 ^ int32(1)
										return v42
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
func F_ReorderBufferFinishPrepared(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32, l7 int64, l8 int32, l9 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int64
	_ = v57
	var v59 int64
	_ = v59
	var v60 int64
	_ = v60
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v119 int32
	_ = v119
	v7 = l6
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = l1
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v21 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v18 + int32(16)
	return
L2:
	;
	v50 = *(*int64)(unsafe.Add(mBase, uint32(v48)+72))
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v48)+32))
	v52 = F_pstrdup(m, l8)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L7
	} else {
		goto L13
	}
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v32 = F_hash_search(m, v26, v18+int32(12), int32(0), v18+int32(11))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	if l1 != v21 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v25 != 0 {
		v48 = v25
		goto L2
	} else {
		goto L6
	}
L6:
	;
	goto L1
L7:
	;
	return
L8:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+11)))
	if v34 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v37
	goto L1
L10:
	;
	goto L11
L11:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v41
	if v42 == int32(0) {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v48 = v42
	goto L2
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+12)) = v52
	if l9 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v48)+72)) = l5
	*(*int64)(unsafe.Add(mBase, uint32(v48)+32)) = l3
	*(*int64)(unsafe.Add(mBase, uint32(v48)+24)) = l2
	*(*int64)(unsafe.Add(mBase, uint32(v48)+64)) = l7
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+56)) = uint16(v7)
	if l9 != 0 {
		goto L19
	} else {
		goto L20
	}
L15:
	;
	v57 = *(*int64)(unsafe.Add(mBase, uint32(v48)+24))
	if base.Ui64(l4) <= base.Ui64(v57) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v59 = *(*int64)(unsafe.Add(mBase, uint32(v48)+32))
	v60 = *(*int64)(unsafe.Add(mBase, uint32(v48)+72))
	v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+56)))
	v62 = *(*int64)(unsafe.Add(mBase, uint32(v48)+64))
	F_ReorderBufferReplay(m, v48, l0, v57, v59, v60, v61, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v48)+172))
	if v77 != 0 {
		goto L24
	} else {
		goto L25
	}
L19:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	m.T0[v71].(func(*base.Module, int32, int32, int64))(m, l0, v48, l2)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L7
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	m.T0[v74].(func(*base.Module, int32, int32, int64, int64))(m, l0, v48, v51, v50)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L7
	} else {
		goto L23
	}
L22:
	;
	goto L18
L23:
	;
	goto L18
L24:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v48)+176))
	v81 = int32(0)
	goto L27
L25:
	;
	goto L26
L26:
	;
	F_ReorderBufferCleanupTXN(m, l0, v48)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L7
	} else {
		goto L31
	}
L27:
	;
	F_LocalExecuteInvalidationMessage(m, v78+v81<<(uint(int32(4))%32))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L7
	} else {
		goto L29
	}
L28:
	;
	goto L26
L29:
	;
	v101 = v81 + int32(1)
	if v101 != v77 {
		v81 = v101
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	goto L1
}
func F_ReorderBufferFree(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	F_MemoryContextDelete(m, v2)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, _consts[518]))
		F_ReorderBufferCleanupSerializedTXNs(m, v6+int32(24))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			return
		}
	}
}
func F_ReorderBufferRestoreCleanup(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v26 int64
	_ = v26
	var v28 int32
	_ = v28
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v37 int32
	_ = v37
	var v42 int64
	_ = v42
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int64
	_ = v60
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	v8 = m.G0
	v10 = v8 - int32(1072)
	m.G0 = v10
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v14 = int64(*(*int32)(unsafe.Add(mBase, _consts[116])))
	v15 = base.I64_div_u_s(v12, v14)
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v17 = base.I64_div_u_s(v16, v14)
	if base.Ui64(v15) <= base.Ui64(v17) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L7
	} else {
		goto L14
	}
L2:
	;
	v26 = v15
	goto L5
L3:
	;
	goto L4
L4:
	;
	m.G0 = v10 + int32(1072)
	return
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v30 = int64(*(*int32)(unsafe.Add(mBase, _consts[116])))
	v31 = v26 * v30
	*(*uint32)(unsafe.Add(mBase, uint32(v10+int32(32)))) = uint32(v31)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(80540)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v28
	v37 = *(*int32)(unsafe.Add(mBase, _consts[518]))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v37 + int32(24)
	v42 = int64(base.Ui64(v31) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v10)+28)) = uint32(v42)
	v50 = F_pg_snprintf(m, v10+int32(48), int32(1024), int32(290489), v10+int32(16))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L4
L7:
	;
	return
L8:
	;
	v54 = F_unlink(m, v10+int32(48))
	mBase = m.M
	if v54 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	if v56 != int32(44) {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v60 = v26 + int64(1)
	if base.Ui64(v60) <= base.Ui64(v17) {
		v26 = v60
		goto L5
	} else {
		goto L13
	}
L12:
	;
	goto L11
L13:
	;
	goto L6
L14:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v10 + int32(48)
	F_errmsg(m, int32(285744), v10)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(474010), int32(4841), int32(221959))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ReorderBufferSerializeTXN(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v130 int64
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int64
	_ = v136
	var v138 int64
	_ = v138
	var v139 int64
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int64
	_ = v144
	var v146 int64
	_ = v146
	var v147 int64
	_ = v147
	var v148 int64
	_ = v148
	var v154 int32
	_ = v154
	var v159 int64
	_ = v159
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v179 int64
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int64
	_ = v199
	var v203 int64
	_ = v203
	var v205 int64
	_ = v205
	var v207 int64
	_ = v207
	var v209 int64
	_ = v209
	var v211 int64
	_ = v211
	var v213 int64
	_ = v213
	var v215 int64
	_ = v215
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int64
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int64
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v278 int64
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int64
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v353 int32
	_ = v353
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
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
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
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
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
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
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
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v525 int64
	_ = v525
	var v526 int64
	_ = v526
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v544 int32
	_ = v544
	var v552 int32
	_ = v552
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v590 int64
	_ = v590
	var v594 int64
	_ = v594
	var v598 int64
	_ = v598
	var v599 int32
	_ = v599
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v666 int32
	_ = v666
	var v671 int32
	_ = v671
	v3 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(1104)
	m.G0 = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+216))
	v29 = F_errstart(m, int32(13), v3)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v29 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v31 = *(*int64)(unsafe.Add(mBase, uint32(l1)+120))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+68)) = v32
	*(*uint32)(unsafe.Add(mBase, uint32(v24)+64)) = uint32(v31)
	F_errmsg_internal(m, int32(300371), v24-int32(-64))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+164))
	if v46 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	F_errfinish(m, int32(474010), int32(3973), int32(505508))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	v100 = int32(-1)
	v101 = int32(1)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
	if v102 == int32(0) {
		v544 = v101
		v552 = v100
		goto L17
	} else {
		goto L18
	}
L9:
	;
	v50 = l1 + int32(160)
	if v46 == v50 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v54 = v46
	goto L11
L11:
	;
	F_ReorderBufferSerializeTXN(m, l0, v54-int32(188))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	goto L8
L13:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v77 != v50 {
		v54 = v77
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v647 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	v648 = F_CloseTransientFile(m, v177)
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L1
	} else {
		goto L179
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L1
	} else {
		goto L175
	}
L17:
	;
	if v26 == int32(0) {
		goto L159
	} else {
		goto L160
	}
L18:
	;
	v106 = l1 + int32(128)
	if v102 == v106 {
		v544 = v101
		v552 = v100
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v117 = v102
	v120 = v100
	v124 = v3
	v130 = int64(0)
	goto L20
L20:
	;
	v132 = v117 - int32(52)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	if v120 != int32(-1) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v544 = base.B2i32(v538 == int32(0))
	v552 = v177
	goto L17
L22:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if v180 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L23:
	;
	v136 = *(*int64)(unsafe.Add(mBase, uint32(v132)))
	v138 = int64(*(*int32)(unsafe.Add(mBase, _consts[116])))
	v139 = base.I64_div_u_s(v136, v138)
	if v139 == v130 {
		v177 = v120
		v179 = v130
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v144 = *(*int64)(unsafe.Add(mBase, uint32(v132)))
	v146 = int64(*(*int32)(unsafe.Add(mBase, _consts[116])))
	v147 = base.I64_div_u_s(v144, v146)
	v148 = v147 * v146
	*(*uint32)(unsafe.Add(mBase, uint32(v24+int32(48)))) = uint32(v148)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = int32(80540)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = v143
	v154 = *(*int32)(unsafe.Add(mBase, _consts[518]))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v154 + int32(24)
	v159 = int64(base.Ui64(v148) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v24)+44)) = uint32(v159)
	v167 = F_pg_snprintf(m, v24+int32(80), int32(1024), int32(290489), v24+int32(32))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	v141 = F_CloseTransientFile(m, v120)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v172 = F_OpenTransientFile(m, v24+int32(80), int32(1089))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	if v172 < int32(0) {
		goto L16
	} else {
		goto L30
	}
L30:
	;
	v177 = v172
	v179 = v147
	goto L22
L31:
	;
	v199 = *(*int64)(unsafe.Add(mBase, uint32(v132)))
	*(*int64)(unsafe.Add(mBase, uint32(v198)+8)) = v199
	v203 = *(*int64)(unsafe.Add(mBase, uint32(v132)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v198-int32(-64)))) = v203
	v205 = *(*int64)(unsafe.Add(mBase, uint32(v132)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v198)+56)) = v205
	v207 = *(*int64)(unsafe.Add(mBase, uint32(v132)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v198)+48)) = v207
	v209 = *(*int64)(unsafe.Add(mBase, uint32(v132)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v198)+40)) = v209
	v211 = *(*int64)(unsafe.Add(mBase, uint32(v132)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v198)+32)) = v211
	v213 = *(*int64)(unsafe.Add(mBase, uint32(v132)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v198)+24)) = v213
	v215 = *(*int64)(unsafe.Add(mBase, uint32(v132)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v198)+16)) = v215
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v117-int32(44))))
	switch v220 {
	case 0, 1, 2, 8:
		goto L44
	case 3:
		goto L43
	case 4:
		goto L42
	case 5:
		goto L41
	default:
		v500 = v198
		v501 = int32(72)
		goto L39
	case 11:
		goto L40
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = int32(72)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v194
	v198 = v194
	goto L31
L33:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v185 = F_MemoryContextAlloc(m, v183, int32(72))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if base.Ui32(int32(71)) < base.Ui32(v180) {
		v198 = v187
		goto L31
	} else {
		goto L37
	}
L36:
	;
	v194 = v185
	goto L32
L37:
	;
	v191 = F_repalloc(m, v187, int32(72))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v194 = v191
	goto L32
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v500))) = v501
	*(*int32)(unsafe.Add(mBase, _consts[86])) = int32(0)
	v513 = *(*int32)(unsafe.Add(mBase, _consts[96]))
	*(*int32)(unsafe.Add(mBase, uint32(v513))) = int32(167772204)
	v516 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v500)))
	v518 = F_write(m, v177, v516, v517)
	mBase = m.M
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v500)))
	if v518 != v519 {
		goto L15
	} else {
		goto L153
	}
L40:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v117-int32(32))))
	v475 = v473 << (uint(int32(2)) % 32)
	v477 = v475 + int32(72)
	v478 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if v478 == int32(0) {
		goto L143
	} else {
		goto L144
	}
L41:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v117-int32(32))))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v422)+24))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v422)+16))
	v429 = (v423+v424)<<(uint(int32(2))%32) + int32(144)
	v430 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if v430 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L42:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v117-int32(32))))
	v395 = v393 << (uint(int32(4)) % 32)
	v397 = v395 + int32(72)
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if v398 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L43:
	;
	v292 = v117 - int32(28)
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v292)))
	v295 = v117 - int32(32)
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v295)))
	if v296&int32(3) == int32(0) {
		v320 = v296
		goto L74
	} else {
		goto L75
	}
L44:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v117-int32(12))))
	v224 = int32(0)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v117-int32(16))))
	if v227 == v224 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	if v223 != 0 {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	v235 = int32(0)
	v236 = int32(72)
	goto L45
L47:
	;
	goto L48
L48:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v227)))
	v235 = v232
	v236 = v232 + int32(92)
	goto L45
L49:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	v241 = v236 + v237 + int32(20)
	v242 = v237
	goto L51
L50:
	;
	v241 = v236
	v242 = v224
	goto L51
L51:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if v243 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v259 = v257 + int32(72)
	if v235 != 0 {
		goto L60
	} else {
		goto L61
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v241
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v254
	v257 = v254
	goto L52
L54:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v247 = F_MemoryContextAlloc(m, v246, v241)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if base.Ui32(v241) <= base.Ui32(v243) {
		v257 = v249
		goto L52
	} else {
		goto L58
	}
L57:
	;
	v254 = v247
	goto L53
L58:
	;
	v251 = F_repalloc(m, v249, v241)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v254 = v251
	goto L53
L60:
	;
	v260 = *(*int64)(unsafe.Add(mBase, uint32(v227)))
	*(*int64)(unsafe.Add(mBase, uint32(v259))) = v260
	v263 = v227 + int32(16)
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v263)))
	*(*int32)(unsafe.Add(mBase, uint32(v259)+16)) = v264
	v266 = *(*int64)(unsafe.Add(mBase, uint32(v227)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v259)+8)) = v266
	v269 = v257 + int32(92)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v263)))
	if v235 != 0 {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	v274 = v259
	goto L62
L62:
	;
	if v242 == int32(0) {
		v500 = v257
		v501 = v241
		goto L39
	} else {
		goto L67
	}
L63:
	;
	v274 = v272 + v235
	goto L62
L64:
	;
	v271 = F__emscripten_memcpy_bulkmem(m, v269, v270, v235)
	mBase = m.M
	v272 = v271
	goto L66
L65:
	;
	v272 = v269
	goto L66
L66:
	;
	goto L63
L67:
	;
	v278 = *(*int64)(unsafe.Add(mBase, uint32(v223)))
	*(*int64)(unsafe.Add(mBase, uint32(v274))) = v278
	v281 = v223 + int32(16)
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)))
	*(*int32)(unsafe.Add(mBase, uint32(v274)+16)) = v282
	v284 = *(*int64)(unsafe.Add(mBase, uint32(v223)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v274)+8)) = v284
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v281)))
	if v242 != 0 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v500 = v257
	v501 = v241
	goto L39
L69:
	;
	v289 = F__emscripten_memcpy_bulkmem(m, v274+int32(20), v288, v242)
	mBase = m.M
	goto L71
L70:
	;
	goto L71
L71:
	;
	goto L68
L72:
	;
	v355 = v353 + int32(1)
	v358 = v293 + v355 + int32(80)
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if v359 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L73:
	;
	v353 = v345 - v296
	goto L72
L74:
	;
	v324 = v320
	goto L83
L75:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296))))
	if v304 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v353 = int32(0)
	goto L72
L77:
	;
	goto L78
L78:
	;
	v309 = v296
	goto L79
L79:
	;
	v313 = v309 + int32(1)
	if v313&int32(3) == int32(0) {
		v320 = v313
		goto L74
	} else {
		goto L81
	}
L80:
	;
	v345 = v313
	goto L73
L81:
	;
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v313))))
	if v318 != 0 {
		v309 = v313
		goto L79
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v324)))
	v333 = int32(-2139062144)
	if (int32(16843008)-v330|v330)&v333 == v333 {
		v324 = v324 + int32(4)
		goto L83
	} else {
		goto L85
	}
L84:
	;
	v339 = v324
	goto L86
L85:
	;
	goto L84
L86:
	;
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v339))))
	if v343 != 0 {
		v339 = v339 + int32(1)
		goto L86
	} else {
		goto L88
	}
L87:
	;
	v345 = v339
	goto L73
L88:
	;
	goto L87
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v373)+72)) = v355
	v376 = v373 + int32(76)
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v295)))
	if v355 != 0 {
		goto L98
	} else {
		goto L99
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v358
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v370
	v373 = v370
	goto L89
L91:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v363 = F_MemoryContextAlloc(m, v362, v358)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if base.Ui32(v358) <= base.Ui32(v359) {
		v373 = v365
		goto L89
	} else {
		goto L95
	}
L94:
	;
	v370 = v363
	goto L90
L95:
	;
	v367 = F_repalloc(m, v365, v358)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v370 = v367
	goto L90
L97:
	;
	v380 = v379 + v355
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v292)))
	*(*int32)(unsafe.Add(mBase, uint32(v380))) = v381
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v117-int32(24))))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v292)))
	if v388 != 0 {
		goto L102
	} else {
		goto L103
	}
L98:
	;
	v378 = F__emscripten_memcpy_bulkmem(m, v376, v377, v355)
	mBase = m.M
	v379 = v378
	goto L100
L99:
	;
	v379 = v376
	goto L100
L100:
	;
	goto L97
L101:
	;
	v500 = v373
	v501 = v358
	goto L39
L102:
	;
	v389 = F__emscripten_memcpy_bulkmem(m, v380+int32(4), v387, v388)
	mBase = m.M
	goto L104
L103:
	;
	goto L104
L104:
	;
	goto L101
L105:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v117-int32(28))))
	if v395 != 0 {
		goto L114
	} else {
		goto L115
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v409
	v412 = v409
	goto L105
L107:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v402 = F_MemoryContextAlloc(m, v401, v397)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if base.Ui32(v397) <= base.Ui32(v398) {
		v412 = v404
		goto L105
	} else {
		goto L111
	}
L110:
	;
	v409 = v402
	goto L106
L111:
	;
	v406 = F_repalloc(m, v404, v397)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	v409 = v406
	goto L106
L113:
	;
	v500 = v412
	v501 = v397
	goto L39
L114:
	;
	v418 = F__emscripten_memcpy_bulkmem(m, v412+int32(72), v417, v395)
	mBase = m.M
	goto L116
L115:
	;
	goto L116
L116:
	;
	goto L113
L117:
	;
	v445 = int32(72)
	goto L126
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v429
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v441
	v444 = v441
	goto L117
L119:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v434 = F_MemoryContextAlloc(m, v433, v429)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if base.Ui32(v429) <= base.Ui32(v430) {
		v444 = v436
		goto L117
	} else {
		goto L123
	}
L122:
	;
	v441 = v434
	goto L118
L123:
	;
	v438 = F_repalloc(m, v436, v429)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	v441 = v438
	goto L118
L125:
	;
	v451 = v444 + int32(144)
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v422)+16))
	if v452 != 0 {
		goto L129
	} else {
		goto L130
	}
L126:
	;
	v448 = F__emscripten_memcpy_bulkmem(m, v444+v445, v422, v445)
	mBase = m.M
	goto L128
L128:
	;
	goto L125
L129:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v422)+12))
	v455 = v452 << (uint(int32(2)) % 32)
	if v455 != 0 {
		goto L133
	} else {
		goto L134
	}
L130:
	;
	v462 = v451
	goto L131
L131:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v422)+24))
	if v463 == int32(0) {
		v500 = v444
		v501 = v429
		goto L39
	} else {
		goto L136
	}
L132:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v422)+16))
	v462 = v457 + v458<<(uint(int32(2))%32)
	goto L131
L133:
	;
	v456 = F__emscripten_memcpy_bulkmem(m, v451, v453, v455)
	mBase = m.M
	v457 = v456
	goto L135
L134:
	;
	v457 = v451
	goto L135
L135:
	;
	goto L132
L136:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v422)+20))
	v468 = v463 << (uint(int32(2)) % 32)
	if v468 != 0 {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	v500 = v444
	v501 = v429
	goto L39
L138:
	;
	v469 = F__emscripten_memcpy_bulkmem(m, v462, v466, v468)
	mBase = m.M
	goto L140
L139:
	;
	goto L140
L140:
	;
	goto L137
L141:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v117-int32(24))))
	if v475 != 0 {
		goto L150
	} else {
		goto L151
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v477
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v489
	v492 = v489
	goto L141
L143:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v482 = F_MemoryContextAlloc(m, v481, v477)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L1
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if base.Ui32(v477) <= base.Ui32(v478) {
		v492 = v484
		goto L141
	} else {
		goto L147
	}
L146:
	;
	v489 = v482
	goto L142
L147:
	;
	v486 = F_repalloc(m, v484, v477)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	v489 = v486
	goto L142
L149:
	;
	v500 = v492
	v501 = v477
	goto L39
L150:
	;
	v498 = F__emscripten_memcpy_bulkmem(m, v492+int32(72), v497, v475)
	mBase = m.M
	goto L152
L151:
	;
	goto L152
L152:
	;
	goto L149
L153:
	;
	v522 = *(*int32)(unsafe.Add(mBase, _consts[96]))
	*(*int32)(unsafe.Add(mBase, uint32(v522))) = int32(0)
	v525 = *(*int64)(unsafe.Add(mBase, uint32(v132)))
	v526 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	if base.Ui64(v526) < base.Ui64(v525) {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v525
	goto L156
L155:
	;
	goto L156
L156:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v529)+4)) = v530
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	*(*int32)(unsafe.Add(mBase, uint32(v530))) = v532
	F_ReorderBufferFreeChange(m, l0, v132, int32(0))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	v538 = v124 + int32(1)
	if v133 != v106 {
		v117 = v133
		v120 = v177
		v124 = v538
		v130 = v179
		goto L20
	} else {
		goto L158
	}
L158:
	;
	goto L21
L159:
	;
	if v544 == int32(0) {
		goto L167
	} else {
		goto L168
	}
L160:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(l1)+216))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+216)) = v565 - v26
	v568 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v569 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v569 - v26
	if v568 != 0 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v572 = v568
	goto L163
L162:
	;
	v572 = l1
	goto L163
L163:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v572)+220))
	*(*int32)(unsafe.Add(mBase, uint32(v572)+220)) = v573 - v26
	v576 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v578 = l1 + int32(204)
	F_pairingheap_remove(m, v576, v578)
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(l1)+216))
	if v581 == int32(0) {
		goto L159
	} else {
		goto L165
	}
L165:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	F_pairingheap_add(m, v584, v578)
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	goto L159
L167:
	;
	v590 = *(*int64)(unsafe.Add(mBase, uint32(l0)+168))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = v590 + int64(1)
	v594 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+176)) = v594 + base.I64_extend_i32_u(v26)
	v598 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
	v599 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v598 + base.I64_extend_i32_u(base.B2i32(v599&int32(4) == int32(0))&int32(base.Ui32(v599^int32(-1))>>(uint(int32(3))%32)))
	v612 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	F_UpdateDecodingStats(m, v612)
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L1
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1)+120)) = int64(0)
	v618 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v618 | int32(4)
	if v552 != int32(-1) {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	goto L169
L171:
	;
	v624 = F_CloseTransientFile(m, v552)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L1
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	m.G0 = v24 + int32(1104)
	return
L174:
	;
	goto L173
L175:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v24 + int32(80)
	F_errmsg(m, int32(285190), v24)
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	F_errfinish(m, int32(474010), int32(4019), int32(505508))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L1
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
	if v647 != 0 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v652 = v647
	goto L182
L181:
	;
	v652 = int32(51)
	goto L182
L182:
	;
	*(*int32)(unsafe.Add(mBase, _consts[86])) = v652
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L1
	} else {
		goto L184
	}
L184:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v660
	F_errmsg(m, int32(279023), v24+int32(16))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L1
	} else {
		goto L185
	}
L185:
	;
	F_errfinish(m, int32(474010), int32(4253), int32(385360))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ReorderBufferTXNByXid(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
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
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int64
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = l1
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v14 == int32(0) {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v27 = F_hash_search(m, v21, v11+int32(12), int32(1), v11+int32(11))
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int32(0)
		} else {
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+11)))
			if v31 == int32(1) {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
				v78 = v34
				v82 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v78
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v82
				if l2 == int32(0) {
					v98 = v78
				} else {
					v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+11)))
					v92 = v78
					v96 = (v87 ^ int32(-1)) & int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v96)
					v98 = v92
				}
				m.G0 = v11 + int32(16)
				return v98
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v37 = F_MemoryContextAlloc(m, v35, int32(232))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					v42 = F__emscripten_memset_bulkmem(m, v37, base.I32_extend8_s(int32(0)), int32(232))
					mBase = m.M
					v44 = v42 + int32(160)
					*(*int32)(unsafe.Add(mBase, uint32(v42)+164)) = v44
					*(*int32)(unsafe.Add(mBase, uint32(v42)+160)) = v44
					v48 = v42 + int32(136)
					*(*int32)(unsafe.Add(mBase, uint32(v42)+140)) = v48
					*(*int32)(unsafe.Add(mBase, uint32(v42)+136)) = v48
					v52 = v42 + int32(128)
					*(*int32)(unsafe.Add(mBase, uint32(v42)+132)) = v52
					*(*int32)(unsafe.Add(mBase, uint32(v42)+128)) = v52
					*(*int32)(unsafe.Add(mBase, uint32(v42)+108)) = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v42
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v42)+4)) = v58
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
					*(*int64)(unsafe.Add(mBase, uint32(v60)+16)) = l3
					v62 = *(*int64)(unsafe.Add(mBase, uint32(l0)+136))
					*(*int64)(unsafe.Add(mBase, uint32(v60)+48)) = v62
					v65 = v60 + int32(188)
					v67 = l0 + int32(4)
					v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if v68 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v67
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v67
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(v60)+192)) = v67
					v74 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
					*(*int32)(unsafe.Add(mBase, uint32(v60)+188)) = v74
					*(*int32)(unsafe.Add(mBase, uint32(v74)+4)) = v65
					*(*int32)(unsafe.Add(mBase, uint32(v67))) = v65
					v78 = v60
					v82 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v78
					*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v82
					if l2 == int32(0) {
						v98 = v78
					} else {
						v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+11)))
						v92 = v78
						v96 = (v87 ^ int32(-1)) & int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v96)
						v98 = v92
					}
					m.G0 = v11 + int32(16)
					return v98
				}
			}
		}
	} else {
		if l1 != v14 {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v27 = F_hash_search(m, v21, v11+int32(12), int32(1), v11+int32(11))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+11)))
				if v31 == int32(1) {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
					v78 = v34
					v82 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v78
					*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v82
					if l2 == int32(0) {
						v98 = v78
					} else {
						v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+11)))
						v92 = v78
						v96 = (v87 ^ int32(-1)) & int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v96)
						v98 = v92
					}
					m.G0 = v11 + int32(16)
					return v98
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
					v37 = F_MemoryContextAlloc(m, v35, int32(232))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						v42 = F__emscripten_memset_bulkmem(m, v37, base.I32_extend8_s(int32(0)), int32(232))
						mBase = m.M
						v44 = v42 + int32(160)
						*(*int32)(unsafe.Add(mBase, uint32(v42)+164)) = v44
						*(*int32)(unsafe.Add(mBase, uint32(v42)+160)) = v44
						v48 = v42 + int32(136)
						*(*int32)(unsafe.Add(mBase, uint32(v42)+140)) = v48
						*(*int32)(unsafe.Add(mBase, uint32(v42)+136)) = v48
						v52 = v42 + int32(128)
						*(*int32)(unsafe.Add(mBase, uint32(v42)+132)) = v52
						*(*int32)(unsafe.Add(mBase, uint32(v42)+128)) = v52
						*(*int32)(unsafe.Add(mBase, uint32(v42)+108)) = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v42
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v42)+4)) = v58
						v60 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
						*(*int64)(unsafe.Add(mBase, uint32(v60)+16)) = l3
						v62 = *(*int64)(unsafe.Add(mBase, uint32(l0)+136))
						*(*int64)(unsafe.Add(mBase, uint32(v60)+48)) = v62
						v65 = v60 + int32(188)
						v67 = l0 + int32(4)
						v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						if v68 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v67
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v67
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(v60)+192)) = v67
						v74 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
						*(*int32)(unsafe.Add(mBase, uint32(v60)+188)) = v74
						*(*int32)(unsafe.Add(mBase, uint32(v74)+4)) = v65
						*(*int32)(unsafe.Add(mBase, uint32(v67))) = v65
						v78 = v60
						v82 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v78
						*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v82
						if l2 == int32(0) {
							v98 = v78
						} else {
							v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+11)))
							v92 = v78
							v96 = (v87 ^ int32(-1)) & int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v96)
							v98 = v92
						}
						m.G0 = v11 + int32(16)
						return v98
					}
				}
			}
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			if v18 != 0 {
				if l2 != 0 {
					v92 = v18
					v96 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v96)
					v98 = v92
				} else {
					v98 = v18
				}
				m.G0 = v11 + int32(16)
				return v98
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v27 = F_hash_search(m, v21, v11+int32(12), int32(1), v11+int32(11))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+11)))
					if v31 == int32(1) {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
						v78 = v34
						v82 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v78
						*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v82
						if l2 == int32(0) {
							v98 = v78
						} else {
							v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+11)))
							v92 = v78
							v96 = (v87 ^ int32(-1)) & int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v96)
							v98 = v92
						}
						m.G0 = v11 + int32(16)
						return v98
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
						v37 = F_MemoryContextAlloc(m, v35, int32(232))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v42 = F__emscripten_memset_bulkmem(m, v37, base.I32_extend8_s(int32(0)), int32(232))
							mBase = m.M
							v44 = v42 + int32(160)
							*(*int32)(unsafe.Add(mBase, uint32(v42)+164)) = v44
							*(*int32)(unsafe.Add(mBase, uint32(v42)+160)) = v44
							v48 = v42 + int32(136)
							*(*int32)(unsafe.Add(mBase, uint32(v42)+140)) = v48
							*(*int32)(unsafe.Add(mBase, uint32(v42)+136)) = v48
							v52 = v42 + int32(128)
							*(*int32)(unsafe.Add(mBase, uint32(v42)+132)) = v52
							*(*int32)(unsafe.Add(mBase, uint32(v42)+128)) = v52
							*(*int32)(unsafe.Add(mBase, uint32(v42)+108)) = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v42
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v42)+4)) = v58
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
							*(*int64)(unsafe.Add(mBase, uint32(v60)+16)) = l3
							v62 = *(*int64)(unsafe.Add(mBase, uint32(l0)+136))
							*(*int64)(unsafe.Add(mBase, uint32(v60)+48)) = v62
							v65 = v60 + int32(188)
							v67 = l0 + int32(4)
							v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							if v68 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v67
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v67
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(v60)+192)) = v67
							v74 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
							*(*int32)(unsafe.Add(mBase, uint32(v60)+188)) = v74
							*(*int32)(unsafe.Add(mBase, uint32(v74)+4)) = v65
							*(*int32)(unsafe.Add(mBase, uint32(v67))) = v65
							v78 = v60
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v78
							*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v82
							if l2 == int32(0) {
								v98 = v78
							} else {
								v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+11)))
								v92 = v78
								v96 = (v87 ^ int32(-1)) & int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v96)
								v98 = v92
							}
							m.G0 = v11 + int32(16)
							return v98
						}
					}
				}
			}
		}
	}
}
func F_ReorderBufferTransferSnapToParent(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int64
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	if v5 != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
		if v6 != 0 {
			v7 = *(*int64)(unsafe.Add(mBase, uint32(l1)+88))
			v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
			if base.Ui64(v8) <= base.Ui64(v7) {
				F_SnapBuildSnapDecRefcount(m, v5)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
					*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v45
					v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
					*(*int32)(unsafe.Add(mBase, uint32(v45))) = v47
					*(*int64)(unsafe.Add(mBase, uint32(l1)+88)) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = int32(0)
					return
				}
			} else {
				F_SnapBuildSnapDecRefcount(m, v6)
				mBase = m.M
				v11 = m.ExcPending
				if v11 != 0 {
					return
				} else {
					v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
					v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
					*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v13
					v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
					*(*int32)(unsafe.Add(mBase, uint32(v13))) = v15
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
					v19 = v17
					*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v19
					v21 = *(*int64)(unsafe.Add(mBase, uint32(l1)+88))
					*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v21
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
					v24 = int32(96)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = l1 + v24
					*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v23
					v29 = l0 + v24
					*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = v29
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
					*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v29
					*(*int64)(unsafe.Add(mBase, uint32(l1)+88)) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = int32(0)
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
					*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = v38
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
					*(*int32)(unsafe.Add(mBase, uint32(v38))) = v40
					return
				}
			}
		} else {
			v19 = v5
			*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v19
			v21 = *(*int64)(unsafe.Add(mBase, uint32(l1)+88))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v21
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
			v24 = int32(96)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = l1 + v24
			*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v23
			v29 = l0 + v24
			*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = v29
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
			*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v29
			*(*int64)(unsafe.Add(mBase, uint32(l1)+88)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = int32(0)
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
			v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
			*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = v38
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
			*(*int32)(unsafe.Add(mBase, uint32(v38))) = v40
			return
		}
	} else {
		return
	}
}
