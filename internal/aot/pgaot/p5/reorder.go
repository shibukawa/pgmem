package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ReorderBufferAbort(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v103 int32
	_ = v103
	v5 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l1
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if base.B2i32(v15 == v5)|base.B2i32(l1 != v15) == v5 {
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
	*(*int64)(unsafe.Add(mBase, uint32(v45)+72)) = l3
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v48&int32(16) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v22 != 0 {
		v45 = v22
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v29 = F_hash_search(m, v23, v12+int32(12), int32(0), v12+int32(11))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L1
L7:
	;
	return
L8:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+11)))
	if v31 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v34
	goto L1
L10:
	;
	goto L11
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v38
	if v39 == int32(0) {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v45 = v39
	goto L2
L13:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v45)+24)) = l2
	F_ReorderBufferCleanupTXN(m, l0, v45)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L7
	} else {
		goto L29
	}
L14:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	m.T0[v53].(func(*base.Module, int32, int32, int64))(m, l0, v45, l2)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v45)+172))
	if v56 == int32(0) {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v45)+176))
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferAbort[0]))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+24))
	v64 = base.B2i32(v62 != int32(0))
	goto L17
L17:
	;
	if v62 != int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_BeginInternalSubTransaction(m, int32(_a_F_ReorderBufferAbort_0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L7
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v72 = int32(0)
	goto L23
L21:
	;
	F_AbortCurrentTransaction(m)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	F_LocalExecuteInvalidationMessage(m, v59+v72<<(uint(int32(4))%32))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L7
	} else {
		goto L25
	}
L24:
	;
	if v64 == int32(0) {
		goto L13
	} else {
		goto L27
	}
L25:
	;
	v86 = v72 + int32(1)
	if v86 != v56 {
		v72 = v86
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
	v91 = m.ExcPending
	if v91 != 0 {
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
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferCheckAndTruncateAbortedTXN[0]))
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
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int64
	_ = v59
	var v61 int64
	_ = v61
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v64 int64
	_ = v64
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v121 int32
	_ = v121
	v7 = l6
	v11 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = l1
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if base.B2i32(v21 == v11)|base.B2i32(l1 != v21) == v11 {
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
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v51)+72))
	v53 = *(*int64)(unsafe.Add(mBase, uint32(v51)+32))
	v54 = F_pstrdup(m, l8)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L7
	} else {
		goto L13
	}
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v28 != 0 {
		v51 = v28
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v35 = F_hash_search(m, v29, v18+int32(12), int32(0), v18+int32(11))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L1
L7:
	;
	return
L8:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+11)))
	if v37 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v40
	goto L1
L10:
	;
	goto L11
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v44
	if v45 == int32(0) {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v51 = v45
	goto L2
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = v54
	if l9 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v51)+72)) = l5
	*(*int64)(unsafe.Add(mBase, uint32(v51)+32)) = l3
	*(*int64)(unsafe.Add(mBase, uint32(v51)+24)) = l2
	*(*int64)(unsafe.Add(mBase, uint32(v51)+64)) = l7
	*(*uint16)(unsafe.Add(mBase, uint32(v51)+56)) = uint16(v7)
	if l9 != 0 {
		goto L19
	} else {
		goto L20
	}
L15:
	;
	v59 = *(*int64)(unsafe.Add(mBase, uint32(v51)+24))
	if base.Ui64(l4) <= base.Ui64(v59) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v61 = *(*int64)(unsafe.Add(mBase, uint32(v51)+32))
	v62 = *(*int64)(unsafe.Add(mBase, uint32(v51)+72))
	v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+56)))
	v64 = *(*int64)(unsafe.Add(mBase, uint32(v51)+64))
	F_ReorderBufferReplay(m, v51, l0, v59, v61, v62, v63, v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v51)+172))
	if v79 != 0 {
		goto L24
	} else {
		goto L25
	}
L19:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	m.T0[v73].(func(*base.Module, int32, int32, int64))(m, l0, v51, l2)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L7
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	m.T0[v76].(func(*base.Module, int32, int32, int64, int64))(m, l0, v51, v53, v52)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
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
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v51)+176))
	v83 = int32(0)
	goto L27
L25:
	;
	goto L26
L26:
	;
	F_ReorderBufferCleanupTXN(m, l0, v51)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L7
	} else {
		goto L31
	}
L27:
	;
	F_LocalExecuteInvalidationMessage(m, v80+v83<<(uint(int32(4))%32))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L7
	} else {
		goto L29
	}
L28:
	;
	goto L26
L29:
	;
	v103 = v83 + int32(1)
	if v103 != v79 {
		v83 = v103
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
		v6 = *(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferFree[0]))
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
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int64
	_ = v58
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	v8 = m.G0
	v10 = v8 - int32(1072)
	m.G0 = v10
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v14 = int64(*(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferRestoreCleanup[0])))
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
	v73 = m.ExcPending
	if v73 != 0 {
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
	v30 = int64(*(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferRestoreCleanup[0])))
	v31 = v26 * v30
	*(*uint32)(unsafe.Add(mBase, uint32(v10+int32(32)))) = uint32(v31)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(_a_F_ReorderBufferRestoreCleanup_0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v28
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferRestoreCleanup[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v37 + int32(24)
	v42 = int64(base.Ui64(v31) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v10)+28)) = uint32(v42)
	v45 = v10 + int32(48)
	v50 = F_pg_snprintf(m, v45, int32(1024), int32(_a_F_ReorderBufferRestoreCleanup_1), v10+int32(16))
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
	v52 = F_unlink(m, v45)
	mBase = m.M
	if v52 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferRestoreCleanup[2]))
	if v54 != int32(44) {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v58 = v26 + int64(1)
	if base.Ui64(v58) <= base.Ui64(v17) {
		v26 = v58
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
	v75 = m.ExcPending
	if v75 != 0 {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v10 + int32(48)
	F_errmsg(m, int32(_a_F_ReorderBufferRestoreCleanup_2), v10)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(_a_F_ReorderBufferRestoreCleanup_3), int32(_a_F_ReorderBufferRestoreCleanup_4), int32(_a_F_ReorderBufferRestoreCleanup_5))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int64
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int64
	_ = v132
	var v134 int64
	_ = v134
	var v135 int64
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int64
	_ = v140
	var v142 int64
	_ = v142
	var v143 int64
	_ = v143
	var v144 int64
	_ = v144
	var v150 int32
	_ = v150
	var v155 int64
	_ = v155
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v174 int64
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int64
	_ = v194
	var v196 int64
	_ = v196
	var v198 int64
	_ = v198
	var v200 int64
	_ = v200
	var v202 int64
	_ = v202
	var v204 int64
	_ = v204
	var v206 int64
	_ = v206
	var v208 int64
	_ = v208
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int64
	_ = v255
	var v257 int64
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v269 int64
	_ = v269
	var v271 int64
	_ = v271
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
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
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
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
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v458 int64
	_ = v458
	var v459 int64
	_ = v459
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v477 int32
	_ = v477
	var v485 int32
	_ = v485
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v522 int64
	_ = v522
	var v526 int64
	_ = v526
	var v530 int64
	_ = v530
	var v531 int32
	_ = v531
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v572 int32
	_ = v572
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v598 int32
	_ = v598
	var v603 int32
	_ = v603
	v3 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(1104)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+216))
	v28 = F_errstart(m, int32(13), v3)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v28 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v30 = *(*int64)(unsafe.Add(mBase, uint32(l1)+120))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v31
	*(*uint32)(unsafe.Add(mBase, uint32(v23)+64)) = uint32(v30)
	F_errmsg_internal(m, int32(_a_F_ReorderBufferSerializeTXN_0), v23-int32(-64))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+164))
	if v45 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	F_errfinish(m, int32(_a_F_ReorderBufferSerializeTXN_1), int32(3973), int32(_a_F_ReorderBufferSerializeTXN_2))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	v97 = int32(-1)
	v98 = int32(1)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
	if v99 == int32(0) {
		v477 = v98
		v485 = v97
		goto L17
	} else {
		goto L18
	}
L9:
	;
	v49 = l1 + int32(160)
	if v45 == v49 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v53 = v45
	goto L11
L11:
	;
	F_ReorderBufferSerializeTXN(m, l0, v53-int32(188))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	goto L8
L13:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v75 != v49 {
		v53 = v75
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v579 = *(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferSerializeTXN[0]))
	v580 = F_CloseTransientFile(m, v172)
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L1
	} else {
		goto L140
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L1
	} else {
		goto L136
	}
L17:
	;
	if v25 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L18:
	;
	v103 = l1 + int32(128)
	if v99 == v103 {
		v477 = v98
		v485 = v97
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v112 = v99
	v117 = v97
	v121 = v3
	v126 = int64(0)
	goto L20
L20:
	;
	v128 = v112 - int32(52)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	if v117 != int32(-1) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v477 = base.B2i32(v471 == int32(0))
	v485 = v172
	goto L17
L22:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if v175 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L23:
	;
	v132 = *(*int64)(unsafe.Add(mBase, uint32(v128)))
	v134 = int64(*(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferSerializeTXN[1])))
	v135 = base.I64_div_u_s(v132, v134)
	if v135 == v126 {
		v172 = v117
		v174 = v126
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v140 = *(*int64)(unsafe.Add(mBase, uint32(v128)))
	v142 = int64(*(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferSerializeTXN[1])))
	v143 = base.I64_div_u_s(v140, v142)
	v144 = v143 * v142
	*(*uint32)(unsafe.Add(mBase, uint32(v23+int32(48)))) = uint32(v144)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = int32(_a_F_ReorderBufferSerializeTXN_3)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v139
	v150 = *(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferSerializeTXN[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v150 + int32(24)
	v155 = int64(base.Ui64(v144) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v23)+44)) = uint32(v155)
	v158 = v23 + int32(80)
	v163 = F_pg_snprintf(m, v158, int32(1024), int32(_a_F_ReorderBufferSerializeTXN_4), v23+int32(32))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	v137 = F_CloseTransientFile(m, v117)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v166 = F_OpenTransientFile(m, v158, int32(1089))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	if v166 < int32(0) {
		goto L16
	} else {
		goto L30
	}
L30:
	;
	v172 = v166
	v174 = v143
	goto L22
L31:
	;
	v194 = *(*int64)(unsafe.Add(mBase, uint32(v128)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v193)+64)) = v194
	v196 = *(*int64)(unsafe.Add(mBase, uint32(v128)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v193)+56)) = v196
	v198 = *(*int64)(unsafe.Add(mBase, uint32(v128)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v193)+48)) = v198
	v200 = *(*int64)(unsafe.Add(mBase, uint32(v128)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v193)+40)) = v200
	v202 = *(*int64)(unsafe.Add(mBase, uint32(v128)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v193)+32)) = v202
	v204 = *(*int64)(unsafe.Add(mBase, uint32(v128)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v193)+24)) = v204
	v206 = *(*int64)(unsafe.Add(mBase, uint32(v128)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v193)+16)) = v206
	v208 = *(*int64)(unsafe.Add(mBase, uint32(v128)))
	*(*int64)(unsafe.Add(mBase, uint32(v193)+8)) = v208
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v112-int32(44))))
	switch v213 {
	case 0, 1, 2, 8:
		goto L44
	case 3:
		goto L43
	case 4:
		goto L42
	case 5:
		goto L41
	default:
		v434 = v193
		v436 = int32(72)
		goto L39
	case 11:
		goto L40
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = int32(72)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v189
	v193 = v189
	goto L31
L33:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v180 = F_MemoryContextAlloc(m, v178, int32(72))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if base.Ui32(int32(71)) < base.Ui32(v175) {
		v193 = v182
		goto L31
	} else {
		goto L37
	}
L36:
	;
	v189 = v180
	goto L32
L37:
	;
	v186 = F_repalloc(m, v182, int32(72))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v189 = v186
	goto L32
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v434))) = v436
	*(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferSerializeTXN[0])) = int32(0)
	v446 = *(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferSerializeTXN[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v446))) = int32(167772204)
	v449 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v434)))
	v451 = F_write(m, v172, v449, v450)
	mBase = m.M
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v434)))
	if v451 != v452 {
		goto L15
	} else {
		goto L114
	}
L40:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v112-int32(32))))
	v408 = v406 << (uint(int32(2)) % 32)
	v410 = v408 + int32(72)
	v411 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if v411 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L41:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v112-int32(32))))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v355)+24))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v355)+16))
	v362 = (v356+v357)<<(uint(int32(2))%32) + int32(144)
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if v363 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L42:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v112-int32(32))))
	v327 = v325 << (uint(int32(4)) % 32)
	v329 = v327 + int32(72)
	v330 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if v330 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L43:
	;
	v280 = v112 - int32(28)
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v280)))
	v283 = v112 - int32(32)
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	v285 = F_strlen(m, v284)
	mBase = m.M
	v287 = v285 + int32(1)
	v290 = v281 + v287 + int32(80)
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if v291 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L44:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v112-int32(12))))
	v217 = int32(0)
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v112-int32(16))))
	if v220 == v217 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	if v216 != 0 {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	v228 = int32(0)
	v229 = int32(72)
	goto L45
L47:
	;
	goto L48
L48:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
	v228 = v225
	v229 = v225 + int32(92)
	goto L45
L49:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	v234 = v229 + v230 + int32(20)
	v235 = v230
	goto L51
L50:
	;
	v234 = v229
	v235 = v217
	goto L51
L51:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if v236 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v252 = v250 + int32(72)
	if v228 != 0 {
		goto L60
	} else {
		goto L61
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v234
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v247
	v250 = v247
	goto L52
L54:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v240 = F_MemoryContextAlloc(m, v239, v234)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if base.Ui32(v234) <= base.Ui32(v236) {
		v250 = v242
		goto L52
	} else {
		goto L58
	}
L57:
	;
	v247 = v240
	goto L53
L58:
	;
	v244 = F_repalloc(m, v242, v234)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v247 = v244
	goto L53
L60:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v220)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v252)+16)) = v253
	v255 = *(*int64)(unsafe.Add(mBase, uint32(v220)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v252)+8)) = v255
	v257 = *(*int64)(unsafe.Add(mBase, uint32(v220)))
	*(*int64)(unsafe.Add(mBase, uint32(v252))) = v257
	v260 = v250 + int32(92)
	if v228 != 0 {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	v264 = v252
	goto L62
L62:
	;
	if v235 == int32(0) {
		v434 = v250
		v436 = v234
		goto L39
	} else {
		goto L66
	}
L63:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v220)+16))
	base.MemoryCopy(m, v260, v261, v228)
	goto L65
L64:
	;
	goto L65
L65:
	;
	v264 = v260 + v228
	goto L62
L66:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v216)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v264)+16)) = v267
	v269 = *(*int64)(unsafe.Add(mBase, uint32(v216)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v264)+8)) = v269
	v271 = *(*int64)(unsafe.Add(mBase, uint32(v216)))
	*(*int64)(unsafe.Add(mBase, uint32(v264))) = v271
	if v235 == int32(0) {
		v434 = v250
		v436 = v234
		goto L39
	} else {
		goto L67
	}
L67:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v216)+16))
	base.MemoryCopy(m, v264+int32(20), v277, v235)
	v434 = v250
	v436 = v234
	goto L39
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v305)+72)) = v287
	v308 = v305 + int32(76)
	if v287 != 0 {
		goto L76
	} else {
		goto L77
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v290
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v302
	v305 = v302
	goto L68
L70:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v295 = F_MemoryContextAlloc(m, v294, v290)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if base.Ui32(v290) <= base.Ui32(v291) {
		v305 = v297
		goto L68
	} else {
		goto L74
	}
L73:
	;
	v302 = v295
	goto L69
L74:
	;
	v299 = F_repalloc(m, v297, v290)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v302 = v299
	goto L69
L76:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	base.MemoryCopy(m, v308, v309, v287)
	goto L78
L77:
	;
	goto L78
L78:
	;
	v311 = v287 + v308
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v280)))
	*(*int32)(unsafe.Add(mBase, uint32(v311))) = v312
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v280)))
	if v314 == int32(0) {
		v434 = v305
		v436 = v290
		goto L39
	} else {
		goto L79
	}
L79:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v112-int32(24))))
	base.MemoryCopy(m, v311+int32(4), v321, v314)
	v434 = v305
	v436 = v290
	goto L39
L80:
	;
	if v327 == int32(0) {
		v434 = v344
		v436 = v329
		goto L39
	} else {
		goto L88
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v329
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v341
	v344 = v341
	goto L80
L82:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v334 = F_MemoryContextAlloc(m, v333, v329)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if base.Ui32(v329) <= base.Ui32(v330) {
		v344 = v336
		goto L80
	} else {
		goto L86
	}
L85:
	;
	v341 = v334
	goto L81
L86:
	;
	v338 = F_repalloc(m, v336, v329)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v341 = v338
	goto L81
L88:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v112-int32(28))))
	base.MemoryCopy(m, v344+int32(72), v351, v327)
	v434 = v344
	v436 = v329
	goto L39
L89:
	;
	v378 = int32(72)
	base.MemoryCopy(m, v377+v378, v355, v378)
	v383 = v377 + int32(144)
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v355)+16))
	if v384 != 0 {
		goto L97
	} else {
		goto L98
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v362
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v374
	v377 = v374
	goto L89
L91:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v367 = F_MemoryContextAlloc(m, v366, v362)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if base.Ui32(v362) <= base.Ui32(v363) {
		v377 = v369
		goto L89
	} else {
		goto L95
	}
L94:
	;
	v374 = v367
	goto L90
L95:
	;
	v371 = F_repalloc(m, v369, v362)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v374 = v371
	goto L90
L97:
	;
	v386 = v384 << (uint(int32(2)) % 32)
	if v386 != 0 {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	v393 = v383
	goto L99
L99:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v355)+24))
	if v395 == int32(0) {
		v434 = v377
		v436 = v362
		goto L39
	} else {
		goto L103
	}
L100:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v355)+12))
	base.MemoryCopy(m, v383, v387, v386)
	goto L102
L101:
	;
	goto L102
L102:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v355)+16))
	v393 = v383 + v389<<(uint(int32(2))%32)
	goto L99
L103:
	;
	v399 = v395 << (uint(int32(2)) % 32)
	if v399 == int32(0) {
		v434 = v377
		v436 = v362
		goto L39
	} else {
		goto L104
	}
L104:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v355)+20))
	base.MemoryCopy(m, v393, v402, v399)
	v434 = v377
	v436 = v362
	goto L39
L105:
	;
	if v408 == int32(0) {
		v434 = v425
		v436 = v410
		goto L39
	} else {
		goto L113
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v410
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v422
	v425 = v422
	goto L105
L107:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v415 = F_MemoryContextAlloc(m, v414, v410)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if base.Ui32(v410) <= base.Ui32(v411) {
		v425 = v417
		goto L105
	} else {
		goto L111
	}
L110:
	;
	v422 = v415
	goto L106
L111:
	;
	v419 = F_repalloc(m, v417, v410)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	v422 = v419
	goto L106
L113:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v112-int32(24))))
	base.MemoryCopy(m, v425+int32(72), v432, v408)
	v434 = v425
	v436 = v410
	goto L39
L114:
	;
	v455 = *(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferSerializeTXN[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v455))) = int32(0)
	v458 = *(*int64)(unsafe.Add(mBase, uint32(v128)))
	v459 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	if base.Ui64(v459) < base.Ui64(v458) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v458
	goto L117
L116:
	;
	goto L117
L117:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v462)+4)) = v463
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	*(*int32)(unsafe.Add(mBase, uint32(v463))) = v465
	F_ReorderBufferFreeChange(m, l0, v128, int32(0))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	v471 = v121 + int32(1)
	if v129 != v103 {
		v112 = v129
		v117 = v172
		v121 = v471
		v126 = v174
		goto L20
	} else {
		goto L119
	}
L119:
	;
	goto L21
L120:
	;
	if v477 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L121:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(l1)+216))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+216)) = v497 - v25
	v500 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v501 - v25
	if v500 != 0 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v504 = v500
	goto L124
L123:
	;
	v504 = l1
	goto L124
L124:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v504)+220))
	*(*int32)(unsafe.Add(mBase, uint32(v504)+220)) = v505 - v25
	v508 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v510 = l1 + int32(204)
	F_pairingheap_remove(m, v508, v510)
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(l1)+216))
	if v513 == int32(0) {
		goto L120
	} else {
		goto L126
	}
L126:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	F_pairingheap_add(m, v516, v510)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	goto L120
L128:
	;
	v522 = *(*int64)(unsafe.Add(mBase, uint32(l0)+168))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = v522 + int64(1)
	v526 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+176)) = v526 + base.I64_extend_i32_u(v25)
	v530 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
	v531 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v530 + base.I64_extend_i32_u(base.B2i32(v531&int32(4) == int32(0))&int32(base.Ui32(v531^int32(-1))>>(uint(int32(3))%32)))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	F_UpdateDecodingStats(m, v544)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L1
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1)+120)) = int64(0)
	v550 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v550 | int32(4)
	if v485 != int32(-1) {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	goto L130
L132:
	;
	v556 = F_CloseTransientFile(m, v485)
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L1
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	m.G0 = v23 + int32(1104)
	return
L135:
	;
	goto L134
L136:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v23 + int32(80)
	F_errmsg(m, int32(_a_F_ReorderBufferSerializeTXN_5), v23)
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(_a_F_ReorderBufferSerializeTXN_1), int32(4019), int32(_a_F_ReorderBufferSerializeTXN_2))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L140:
	;
	if v579 != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v584 = v579
	goto L143
L142:
	;
	v584 = int32(51)
	goto L143
L143:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferSerializeTXN[0])) = v584
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v592
	F_errmsg(m, int32(_a_F_ReorderBufferSerializeTXN_6), v23+int32(16))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(_a_F_ReorderBufferSerializeTXN_1), int32(_a_F_ReorderBufferSerializeTXN_7), int32(_a_F_ReorderBufferSerializeTXN_8))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ReorderBufferTXNByXid(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int64
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	v5 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if base.B2i32(v13 == v5)|base.B2i32(l1 != v13) == v5 {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		if v20 != 0 {
			if l2 != 0 {
				v92 = v20
				v95 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v95)
				v97 = v92
			} else {
				v97 = v20
			}
			m.G0 = v10 + int32(16)
			return v97
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v29 = F_hash_search(m, v23, v10+int32(12), int32(1), v10+int32(11))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+11)))
				if v33 == int32(1) {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
					v79 = v36
					v82 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v79
					*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v82
					if l2 == int32(0) {
						v97 = v79
					} else {
						v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+11)))
						v92 = v79
						v95 = (v87 ^ int32(-1)) & int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v95)
						v97 = v92
					}
					m.G0 = v10 + int32(16)
					return v97
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
					v39 = F_MemoryContextAlloc(m, v37, int32(232))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						v41 = int32(0)
						base.MemoryFill(m, v39, v41, int32(232))
						v45 = v39 + int32(160)
						*(*int32)(unsafe.Add(mBase, uint32(v39)+164)) = v45
						*(*int32)(unsafe.Add(mBase, uint32(v39)+160)) = v45
						v49 = v39 + int32(136)
						*(*int32)(unsafe.Add(mBase, uint32(v39)+140)) = v49
						*(*int32)(unsafe.Add(mBase, uint32(v39)+136)) = v49
						v53 = v39 + int32(128)
						*(*int32)(unsafe.Add(mBase, uint32(v39)+132)) = v53
						*(*int32)(unsafe.Add(mBase, uint32(v39)+128)) = v53
						*(*int32)(unsafe.Add(mBase, uint32(v39)+108)) = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v39
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = v59
						v61 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
						*(*int64)(unsafe.Add(mBase, uint32(v61)+16)) = l3
						v63 = *(*int64)(unsafe.Add(mBase, uint32(l0)+136))
						*(*int64)(unsafe.Add(mBase, uint32(v61)+48)) = v63
						v66 = l0 + int32(4)
						v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						if v67 == v41 {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v66
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v66
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(v61)+192)) = v66
						v73 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
						*(*int32)(unsafe.Add(mBase, uint32(v61)+188)) = v73
						v76 = v61 + int32(188)
						*(*int32)(unsafe.Add(mBase, uint32(v73)+4)) = v76
						*(*int32)(unsafe.Add(mBase, uint32(v66))) = v76
						v79 = v61
						v82 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v79
						*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v82
						if l2 == int32(0) {
							v97 = v79
						} else {
							v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+11)))
							v92 = v79
							v95 = (v87 ^ int32(-1)) & int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v95)
							v97 = v92
						}
						m.G0 = v10 + int32(16)
						return v97
					}
				}
			}
		}
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v29 = F_hash_search(m, v23, v10+int32(12), int32(1), v10+int32(11))
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+11)))
			if v33 == int32(1) {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
				v79 = v36
				v82 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v79
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v82
				if l2 == int32(0) {
					v97 = v79
				} else {
					v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+11)))
					v92 = v79
					v95 = (v87 ^ int32(-1)) & int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v95)
					v97 = v92
				}
				m.G0 = v10 + int32(16)
				return v97
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v39 = F_MemoryContextAlloc(m, v37, int32(232))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					v41 = int32(0)
					base.MemoryFill(m, v39, v41, int32(232))
					v45 = v39 + int32(160)
					*(*int32)(unsafe.Add(mBase, uint32(v39)+164)) = v45
					*(*int32)(unsafe.Add(mBase, uint32(v39)+160)) = v45
					v49 = v39 + int32(136)
					*(*int32)(unsafe.Add(mBase, uint32(v39)+140)) = v49
					*(*int32)(unsafe.Add(mBase, uint32(v39)+136)) = v49
					v53 = v39 + int32(128)
					*(*int32)(unsafe.Add(mBase, uint32(v39)+132)) = v53
					*(*int32)(unsafe.Add(mBase, uint32(v39)+128)) = v53
					*(*int32)(unsafe.Add(mBase, uint32(v39)+108)) = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v39
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = v59
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
					*(*int64)(unsafe.Add(mBase, uint32(v61)+16)) = l3
					v63 = *(*int64)(unsafe.Add(mBase, uint32(l0)+136))
					*(*int64)(unsafe.Add(mBase, uint32(v61)+48)) = v63
					v66 = l0 + int32(4)
					v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if v67 == v41 {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v66
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v66
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(v61)+192)) = v66
					v73 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
					*(*int32)(unsafe.Add(mBase, uint32(v61)+188)) = v73
					v76 = v61 + int32(188)
					*(*int32)(unsafe.Add(mBase, uint32(v73)+4)) = v76
					*(*int32)(unsafe.Add(mBase, uint32(v66))) = v76
					v79 = v61
					v82 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v79
					*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v82
					if l2 == int32(0) {
						v97 = v79
					} else {
						v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+11)))
						v92 = v79
						v95 = (v87 ^ int32(-1)) & int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v95)
						v97 = v92
					}
					m.G0 = v10 + int32(16)
					return v97
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
