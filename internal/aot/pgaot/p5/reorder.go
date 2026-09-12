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
	F_BeginInternalSubTransaction(m, int32(26401))
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
	v6 = *(*int32)(unsafe.Add(mBase, _consts[521]))
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
		v6 = *(*int32)(unsafe.Add(mBase, _consts[517]))
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
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(84122)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v28
	v37 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v37 + int32(24)
	v42 = int64(base.Ui64(v31) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v10)+28)) = uint32(v42)
	v50 = F_pg_snprintf(m, v10+int32(48), int32(1024), int32(301268), v10+int32(16))
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
	F_errmsg(m, int32(296337), v10)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(490532), int32(4841), int32(230919))
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
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
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
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
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
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v469 int64
	_ = v469
	var v470 int64
	_ = v470
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v488 int32
	_ = v488
	var v496 int32
	_ = v496
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v534 int64
	_ = v534
	var v538 int64
	_ = v538
	var v542 int64
	_ = v542
	var v543 int32
	_ = v543
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v584 int32
	_ = v584
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v610 int32
	_ = v610
	var v615 int32
	_ = v615
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
	F_errmsg_internal(m, int32(311642), v24-int32(-64))
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
	F_errfinish(m, int32(490532), int32(3973), int32(522582))
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
		v488 = v101
		v496 = v100
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
	v591 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	v592 = F_CloseTransientFile(m, v177)
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L1
	} else {
		goto L162
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L1
	} else {
		goto L158
	}
L17:
	;
	if v26 == int32(0) {
		goto L142
	} else {
		goto L143
	}
L18:
	;
	v106 = l1 + int32(128)
	if v102 == v106 {
		v488 = v101
		v496 = v100
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
	v488 = base.B2i32(v482 == int32(0))
	v496 = v177
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
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = int32(84122)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = v143
	v154 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v154 + int32(24)
	v159 = int64(base.Ui64(v148) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v24)+44)) = uint32(v159)
	v167 = F_pg_snprintf(m, v24+int32(80), int32(1024), int32(301268), v24+int32(32))
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
		v444 = v198
		v445 = int32(72)
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
	*(*int32)(unsafe.Add(mBase, uint32(v444))) = v445
	*(*int32)(unsafe.Add(mBase, _consts[86])) = int32(0)
	v457 = *(*int32)(unsafe.Add(mBase, _consts[96]))
	*(*int32)(unsafe.Add(mBase, uint32(v457))) = int32(167772204)
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v444)))
	v462 = F_write(m, v177, v460, v461)
	mBase = m.M
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v444)))
	if v462 != v463 {
		goto L15
	} else {
		goto L136
	}
L40:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v117-int32(32))))
	v419 = v417 << (uint(int32(2)) % 32)
	v421 = v419 + int32(72)
	v422 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if v422 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L41:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v117-int32(32))))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v366)+24))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v366)+16))
	v373 = (v367+v368)<<(uint(int32(2))%32) + int32(144)
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if v374 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L42:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v117-int32(32))))
	v339 = v337 << (uint(int32(4)) % 32)
	v341 = v339 + int32(72)
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if v342 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L43:
	;
	v292 = v117 - int32(28)
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v292)))
	v295 = v117 - int32(32)
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v295)))
	v297 = F_strlen(m, v296)
	mBase = m.M
	v299 = v297 + int32(1)
	v302 = v293 + v299 + int32(80)
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if v303 == int32(0) {
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
		v444 = v257
		v445 = v241
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
	v444 = v257
	v445 = v241
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
	*(*int32)(unsafe.Add(mBase, uint32(v317)+72)) = v299
	v320 = v317 + int32(76)
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v295)))
	if v299 != 0 {
		goto L81
	} else {
		goto L82
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v302
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v314
	v317 = v314
	goto L72
L74:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v307 = F_MemoryContextAlloc(m, v306, v302)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if base.Ui32(v302) <= base.Ui32(v303) {
		v317 = v309
		goto L72
	} else {
		goto L78
	}
L77:
	;
	v314 = v307
	goto L73
L78:
	;
	v311 = F_repalloc(m, v309, v302)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v314 = v311
	goto L73
L80:
	;
	v324 = v323 + v299
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v292)))
	*(*int32)(unsafe.Add(mBase, uint32(v324))) = v325
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v117-int32(24))))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v292)))
	if v332 != 0 {
		goto L85
	} else {
		goto L86
	}
L81:
	;
	v322 = F__emscripten_memcpy_bulkmem(m, v320, v321, v299)
	mBase = m.M
	v323 = v322
	goto L83
L82:
	;
	v323 = v320
	goto L83
L83:
	;
	goto L80
L84:
	;
	v444 = v317
	v445 = v302
	goto L39
L85:
	;
	v333 = F__emscripten_memcpy_bulkmem(m, v324+int32(4), v331, v332)
	mBase = m.M
	goto L87
L86:
	;
	goto L87
L87:
	;
	goto L84
L88:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v117-int32(28))))
	if v339 != 0 {
		goto L97
	} else {
		goto L98
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v341
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v353
	v356 = v353
	goto L88
L90:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v346 = F_MemoryContextAlloc(m, v345, v341)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if base.Ui32(v341) <= base.Ui32(v342) {
		v356 = v348
		goto L88
	} else {
		goto L94
	}
L93:
	;
	v353 = v346
	goto L89
L94:
	;
	v350 = F_repalloc(m, v348, v341)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v353 = v350
	goto L89
L96:
	;
	v444 = v356
	v445 = v341
	goto L39
L97:
	;
	v362 = F__emscripten_memcpy_bulkmem(m, v356+int32(72), v361, v339)
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
	v389 = int32(72)
	goto L109
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v373
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v385
	v388 = v385
	goto L100
L102:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v378 = F_MemoryContextAlloc(m, v377, v373)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if base.Ui32(v373) <= base.Ui32(v374) {
		v388 = v380
		goto L100
	} else {
		goto L106
	}
L105:
	;
	v385 = v378
	goto L101
L106:
	;
	v382 = F_repalloc(m, v380, v373)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	v385 = v382
	goto L101
L108:
	;
	v395 = v388 + int32(144)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v366)+16))
	if v396 != 0 {
		goto L112
	} else {
		goto L113
	}
L109:
	;
	v392 = F__emscripten_memcpy_bulkmem(m, v388+v389, v366, v389)
	mBase = m.M
	goto L111
L111:
	;
	goto L108
L112:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v366)+12))
	v399 = v396 << (uint(int32(2)) % 32)
	if v399 != 0 {
		goto L116
	} else {
		goto L117
	}
L113:
	;
	v406 = v395
	goto L114
L114:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v366)+24))
	if v407 == int32(0) {
		v444 = v388
		v445 = v373
		goto L39
	} else {
		goto L119
	}
L115:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v366)+16))
	v406 = v401 + v402<<(uint(int32(2))%32)
	goto L114
L116:
	;
	v400 = F__emscripten_memcpy_bulkmem(m, v395, v397, v399)
	mBase = m.M
	v401 = v400
	goto L118
L117:
	;
	v401 = v395
	goto L118
L118:
	;
	goto L115
L119:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v366)+20))
	v412 = v407 << (uint(int32(2)) % 32)
	if v412 != 0 {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v444 = v388
	v445 = v373
	goto L39
L121:
	;
	v413 = F__emscripten_memcpy_bulkmem(m, v406, v410, v412)
	mBase = m.M
	goto L123
L122:
	;
	goto L123
L123:
	;
	goto L120
L124:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v117-int32(24))))
	if v419 != 0 {
		goto L133
	} else {
		goto L134
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v421
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v433
	v436 = v433
	goto L124
L126:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v426 = F_MemoryContextAlloc(m, v425, v421)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L1
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if base.Ui32(v421) <= base.Ui32(v422) {
		v436 = v428
		goto L124
	} else {
		goto L130
	}
L129:
	;
	v433 = v426
	goto L125
L130:
	;
	v430 = F_repalloc(m, v428, v421)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	v433 = v430
	goto L125
L132:
	;
	v444 = v436
	v445 = v421
	goto L39
L133:
	;
	v442 = F__emscripten_memcpy_bulkmem(m, v436+int32(72), v441, v419)
	mBase = m.M
	goto L135
L134:
	;
	goto L135
L135:
	;
	goto L132
L136:
	;
	v466 = *(*int32)(unsafe.Add(mBase, _consts[96]))
	*(*int32)(unsafe.Add(mBase, uint32(v466))) = int32(0)
	v469 = *(*int64)(unsafe.Add(mBase, uint32(v132)))
	v470 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	if base.Ui64(v470) < base.Ui64(v469) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v469
	goto L139
L138:
	;
	goto L139
L139:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v473)+4)) = v474
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	*(*int32)(unsafe.Add(mBase, uint32(v474))) = v476
	F_ReorderBufferFreeChange(m, l0, v132, int32(0))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	v482 = v124 + int32(1)
	if v133 != v106 {
		v117 = v133
		v120 = v177
		v124 = v482
		v130 = v179
		goto L20
	} else {
		goto L141
	}
L141:
	;
	goto L21
L142:
	;
	if v488 == int32(0) {
		goto L150
	} else {
		goto L151
	}
L143:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(l1)+216))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+216)) = v509 - v26
	v512 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v513 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v513 - v26
	if v512 != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v516 = v512
	goto L146
L145:
	;
	v516 = l1
	goto L146
L146:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v516)+220))
	*(*int32)(unsafe.Add(mBase, uint32(v516)+220)) = v517 - v26
	v520 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v522 = l1 + int32(204)
	F_pairingheap_remove(m, v520, v522)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l1)+216))
	if v525 == int32(0) {
		goto L142
	} else {
		goto L148
	}
L148:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	F_pairingheap_add(m, v528, v522)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	goto L142
L150:
	;
	v534 = *(*int64)(unsafe.Add(mBase, uint32(l0)+168))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = v534 + int64(1)
	v538 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+176)) = v538 + base.I64_extend_i32_u(v26)
	v542 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
	v543 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v542 + base.I64_extend_i32_u(base.B2i32(v543&int32(4) == int32(0))&int32(base.Ui32(v543^int32(-1))>>(uint(int32(3))%32)))
	v556 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	F_UpdateDecodingStats(m, v556)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L1
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1)+120)) = int64(0)
	v562 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v562 | int32(4)
	if v496 != int32(-1) {
		goto L154
	} else {
		goto L155
	}
L153:
	;
	goto L152
L154:
	;
	v568 = F_CloseTransientFile(m, v496)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L1
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	m.G0 = v24 + int32(1104)
	return
L157:
	;
	goto L156
L158:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v24 + int32(80)
	F_errmsg(m, int32(295783), v24)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	F_errfinish(m, int32(490532), int32(4019), int32(522582))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L162:
	;
	if v591 != 0 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v596 = v591
	goto L165
L164:
	;
	v596 = int32(51)
	goto L165
L165:
	;
	*(*int32)(unsafe.Add(mBase, _consts[86])) = v596
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v604
	F_errmsg(m, int32(289578), v24+int32(16))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	F_errfinish(m, int32(490532), int32(4253), int32(399293))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
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
