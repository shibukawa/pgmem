package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_GetReplicationTransferLatency(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v20 int32
	_ = v20
	var v28 int64
	_ = v28
	var v37 int64
	_ = v37
	var v40 int32
	_ = v40
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_GetReplicationTransferLatency[0]))
	v8 = base.AtomicRmwXchg32(m, v5, int32(1456), int32(1))
	if v8 != 0 {
		F_s_lock(m, v5+int32(1456), int32(_a_F_GetReplicationTransferLatency_0), int32(401), int32(_a_F_GetReplicationTransferLatency_1))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = *(*int64)(unsafe.Add(mBase, uint32(v5)+80))
			v19 = *(*int64)(unsafe.Add(mBase, uint32(v5)+72))
			v20 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v5)+1456)), uint32(v20))
			if v18 <= v19 {
				v40 = v20
			} else {
				v28 = v18 - v19
				if base.B2i32(int64(0) < v19)^base.B2i32(v28 < v18)|base.B2i32(int64(2147483646000) < v28) != 0 {
					v40 = int32(2147483647)
				} else {
					v37 = base.I64_div_s(v28+int64(999), int64(1000))
					v40 = base.I32_wrap_i64(v37)
				}
			}
			return v40
		}
	} else {
		v18 = *(*int64)(unsafe.Add(mBase, uint32(v5)+80))
		v19 = *(*int64)(unsafe.Add(mBase, uint32(v5)+72))
		v20 = int32(0)
		atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v5)+1456)), uint32(v20))
		if v18 <= v19 {
			v40 = v20
		} else {
			v28 = v18 - v19
			if base.B2i32(int64(0) < v19)^base.B2i32(v28 < v18)|base.B2i32(int64(2147483646000) < v28) != 0 {
				v40 = int32(2147483647)
			} else {
				v37 = base.I64_div_s(v28+int64(999), int64(1000))
				v40 = base.I32_wrap_i64(v37)
			}
		}
		return v40
	}
}
func F_ReplicationSlotDrop(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	F_ReplicationSlotAcquire(m, l0, l1, int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		v14 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotDrop[0])))
		if v14 == int32(1) {
			v19 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDrop[1]))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+316))
			v22 = base.B2i32(v20 != int32(2))
			*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotDrop[0])) = uint8(v22)
			v24 = v22
		} else {
			v24 = int32(0)
		}
		v26 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDrop[2]))
		if v24 == int32(0) {
			*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDrop[2])) = int32(0)
			F_ReplicationSlotDropPtr(m, v26)
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return
			} else {
				m.G0 = v7 + int32(16)
				return
			}
		} else {
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+201)))
			if v29 == int32(0) {
				*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDrop[2])) = int32(0)
				F_ReplicationSlotDropPtr(m, v26)
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return
				} else {
					m.G0 = v7 + int32(16)
					return
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					F_errcode(m, int32(325))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
						F_errmsg(m, int32(_a_F_ReplicationSlotDrop_0), v7)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							F_errdetail(m, int32(_a_F_ReplicationSlotDrop_1), int32(0))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_ReplicationSlotDrop_2), int32(858), int32(_a_F_ReplicationSlotDrop_3))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
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
func F_ReplicationSlotNameForTablesync(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotNameForTablesync[0]))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v11
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
	v17 = F_pg_snprintf(m, l2, int32(64), int32(_a_F_ReplicationSlotNameForTablesync_0), v7)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		m.G0 = v7 + int32(16)
		return
	}
}
func F_ReplicationSlotSave(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v3 = m.G0
	v5 = v3 - int32(1040)
	m.G0 = v5
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(_a_F_ReplicationSlotSave_0)
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotSave[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v10 + int32(24)
	v15 = v5 + int32(16)
	v17 = F_pg_sprintf(m, v15, int32(_a_F_ReplicationSlotSave_1), v5)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotSave[0]))
		F_SaveSlotToPath(m, v20, v15, int32(21))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			m.G0 = v5 + int32(1040)
			return
		}
	}
}
func F_ReplicationSlotsComputeLogicalRestartLSN(m *base.Module) int64 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v6 int64
	_ = v6
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int64
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int64
	_ = v65
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v75 int64
	_ = v75
	var v76 int32
	_ = v76
	var v80 int64
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v93 int64
	_ = v93
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	v1 = int32(0)
	v6 = int64(0)
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeLogicalRestartLSN[0]))
	if v10 <= v1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	goto L3
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeLogicalRestartLSN[1]))
	v20 = F_LWLockAcquire(m, v16+int32(_a_F_ReplicationSlotsComputeLogicalRestartLSN_0), int32(1))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int64(0)
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeLogicalRestartLSN[0]))
	if int32(0) < v25 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeLogicalRestartLSN[2]))
	v31 = v29
	v32 = v1
	v37 = v6
	goto L9
L7:
	;
	v93 = v6
	goto L8
L8:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeLogicalRestartLSN[1]))
	F_LWLockRelease(m, v95+int32(_a_F_ReplicationSlotsComputeLogicalRestartLSN_0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L4
	} else {
		goto L33
	}
L9:
	;
	v40 = v31 + v32*int32(288)
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+4)))
	if v41 != int32(1) {
		v76 = v31
		v80 = v37
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v93 = v80
	goto L8
L11:
	;
	v82 = v32 + int32(1)
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeLogicalRestartLSN[0]))
	if v82 < v84 {
		v31 = v76
		v32 = v82
		v37 = v80
		goto L9
	} else {
		goto L32
	}
L12:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40)+88))
	if v44 == int32(0) {
		v76 = v31
		v80 = v37
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v49 = base.AtomicRmwXchg32(m, v40, int32(0), int32(1))
	if v49 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	F_s_lock(m, v40, int32(_a_F_ReplicationSlotsComputeLogicalRestartLSN_1), int32(1324), int32(_a_F_ReplicationSlotsComputeLogicalRestartLSN_2))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v55 = *(*int64)(unsafe.Add(mBase, uint32(v40)+280))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v40)+112))
	v57 = *(*int64)(unsafe.Add(mBase, uint32(v40)+104))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v40)+92))
	v59 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v40))), uint32(v59))
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeLogicalRestartLSN[2]))
	if v56 != 0 {
		v76 = v63
		v80 = v37
		goto L11
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	if base.Ui64(v57) < base.Ui64(v55) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v65 = v57
	goto L21
L20:
	;
	v65 = v55
	goto L21
L21:
	;
	if v55 != int64(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v68 = v65
	goto L24
L23:
	;
	v68 = v57
	goto L24
L24:
	;
	if v58 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v69 = v57
	goto L27
L26:
	;
	v69 = v68
	goto L27
L27:
	;
	if v69 == int64(0) {
		v76 = v63
		v80 = v37
		goto L11
	} else {
		goto L28
	}
L28:
	;
	if base.Ui64(v37-int64(1)) < base.Ui64(v69) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v75 = v37
	goto L31
L30:
	;
	v75 = v69
	goto L31
L31:
	;
	v76 = v63
	v80 = v75
	goto L11
L32:
	;
	goto L10
L33:
	;
	return v93
}
func F_ReplicationSlotsComputeRequiredXmin(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
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
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	v2 = int32(0)
	if l0 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredXmin[0]))
	v16 = F_LWLockAcquire(m, v12+int32(_a_F_ReplicationSlotsComputeRequiredXmin_0), int32(1))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredXmin[1]))
	if int32(0) < v19 {
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
	v26 = v2
	v27 = v2
	v28 = v2
	goto L9
L7:
	;
	v98 = v2
	v99 = v2
	goto L8
L8:
	;
	v102 = m.G0
	v104 = v102 - int32(16)
	m.G0 = v104
	if l0 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L9:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredXmin[2]))
	v34 = v31 + v28*int32(288)
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+4)))
	if v35 != int32(1) {
		v87 = v26
		v88 = v27
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v98 = v87
	v99 = v88
	goto L8
L11:
	;
	v90 = v28 + int32(1)
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredXmin[1]))
	if v90 < v92 {
		v26 = v87
		v27 = v88
		v28 = v90
		goto L9
	} else {
		goto L37
	}
L12:
	;
	v40 = base.AtomicRmwXchg32(m, v34, int32(0), int32(1))
	if v40 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_s_lock(m, v34, int32(_a_F_ReplicationSlotsComputeRequiredXmin_1), int32(1188), int32(_a_F_ReplicationSlotsComputeRequiredXmin_2))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v34)+112))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	v49 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v34))), uint32(v49))
	if v46 != 0 {
		v87 = v26
		v88 = v27
		goto L11
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	if v48 == int32(0) {
		v68 = v26
		goto L18
	} else {
		goto L19
	}
L18:
	;
	if v47 == int32(0) {
		v87 = v68
		v88 = v27
		goto L11
	} else {
		goto L28
	}
L19:
	;
	if v26 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v26))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v48)) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	goto L22
L22:
	;
	v68 = v48
	goto L18
L23:
	;
	if v65 == int32(0) {
		v68 = v26
		goto L18
	} else {
		goto L27
	}
L24:
	;
	v65 = base.B2i32(base.Ui32(v48) < base.Ui32(v26))
	goto L23
L25:
	;
	goto L26
L26:
	;
	v65 = int32(base.Ui32(v48-v26) >> (uint(int32(31)) % 32))
	goto L23
L27:
	;
	goto L22
L28:
	;
	if v27 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v27))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v47)) == int32(0) {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	goto L31
L31:
	;
	v87 = v68
	v88 = v47
	goto L11
L32:
	;
	if v82 == int32(0) {
		v87 = v68
		v88 = v27
		goto L11
	} else {
		goto L36
	}
L33:
	;
	v82 = base.B2i32(base.Ui32(v47) < base.Ui32(v27))
	goto L32
L34:
	;
	goto L35
L35:
	;
	v82 = int32(base.Ui32(v47-v27) >> (uint(int32(31)) % 32))
	goto L32
L36:
	;
	goto L31
L37:
	;
	goto L10
L38:
	;
	v132 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L4
	} else {
		goto L44
	}
L39:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredXmin[0]))
	v113 = F_LWLockAcquire(m, v109+int32(512), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L4
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredXmin[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v126)+32)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v126)+28)) = v98
	goto L38
L42:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredXmin[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v116)+32)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v116)+28)) = v98
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredXmin[0]))
	F_LWLockRelease(m, v120+int32(512))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	goto L38
L44:
	;
	if v132 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v104)+4)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v104))) = v98
	F_errmsg_internal(m, int32(_a_F_ReplicationSlotsComputeRequiredXmin_3), v104)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L4
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	m.G0 = v104 + int32(16)
	if l0 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	F_errfinish(m, int32(_a_F_ReplicationSlotsComputeRequiredXmin_4), int32(3958), int32(_a_F_ReplicationSlotsComputeRequiredXmin_5))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v150 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredXmin[0]))
	F_LWLockRelease(m, v150+int32(_a_F_ReplicationSlotsComputeRequiredXmin_0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L4
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	return
L53:
	;
	goto L52
}
func F_ReplicationSlotsDropDBSlots(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsDropDBSlots[0]))
	if int32(0) < v13 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v123 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v46))), uint32(v123))
	F_errstart_cold(m, int32(21), v123)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L5
	} else {
		goto L27
	}
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsDropDBSlots[1]))
	v21 = F_LWLockAcquire(m, v17+int32(_a_F_ReplicationSlotsDropDBSlots_0), int32(1))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	m.G0 = v10 + int32(16)
	return
L5:
	;
	return
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsDropDBSlots[0]))
	if v24 <= int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsDropDBSlots[1]))
	F_LWLockRelease(m, v108+int32(_a_F_ReplicationSlotsDropDBSlots_0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L5
	} else {
		goto L26
	}
L8:
	;
	v31 = v24
	goto L9
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsDropDBSlots[2]))
	v39 = int32(0)
	goto L11
L10:
	;
	goto L7
L11:
	;
	v46 = v35 + v39*int32(288)
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+4)))
	if v47 != int32(1) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L10
L13:
	;
	v98 = v39 + int32(1)
	if v98 != v31 {
		v39 = v98
		goto L11
	} else {
		goto L25
	}
L14:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+88))
	if base.B2i32(v50 == int32(0))|base.B2i32(l0 != v50) != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v57 = base.AtomicRmwXchg32(m, v46, int32(0), int32(1))
	if v57 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	F_s_lock(m, v46, int32(_a_F_ReplicationSlotsDropDBSlots_1), int32(1464), int32(_a_F_ReplicationSlotsDropDBSlots_2))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L5
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	if v63 != 0 {
		goto L1
	} else {
		goto L20
	}
L19:
	;
	goto L18
L20:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsDropDBSlots[3])) = v46
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsDropDBSlots[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+8)) = v67
	v69 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v46))), uint32(v69))
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsDropDBSlots[1]))
	F_LWLockRelease(m, v73+int32(_a_F_ReplicationSlotsDropDBSlots_0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	v78 = int32(_a_F_ReplicationSlotsDropDBSlots_3)
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsDropDBSlots[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsDropDBSlots[3])) = int32(0)
	F_ReplicationSlotDropPtr(m, v79)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsDropDBSlots[1]))
	v90 = F_LWLockAcquire(m, v86+int32(_a_F_ReplicationSlotsDropDBSlots_0), int32(1))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsDropDBSlots[0]))
	if int32(0) < v93 {
		v31 = v93
		goto L9
	} else {
		goto L24
	}
L24:
	;
	goto L7
L25:
	;
	goto L12
L26:
	;
	goto L4
L27:
	;
	F_errcode(m, int32(100663621))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L5
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v46 + int32(24)
	F_errmsg(m, int32(_a_F_ReplicationSlotsDropDBSlots_4), v10)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_ReplicationSlotsDropDBSlots_1), int32(1500), int32(_a_F_ReplicationSlotsDropDBSlots_2))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_replication_scanner_finish(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_pfree(m, v2)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		F_replication_yylex_destroy(m, l0)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			return
		}
	}
}
func F_replication_yy_create_buffer(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	v8 = F_palloc(m, int32(48))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if v8 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(_a_F_replication_yy_create_buffer_0)
			v15 = F_palloc(m, int32(_a_F_replication_yy_create_buffer_1))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v15
				if v15 == int32(0) {
					F_yy_fatal_error_3(m, int32(_a_F_replication_yy_create_buffer_2))
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v20 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v20
					v23 = *(*int32)(unsafe.Add(mBase, _c_F_replication_yy_create_buffer[0]))
					v24 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v24
					*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v24)
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
					*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)) = uint8(v24)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = v24
					*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v20
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v35
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					if v37 == v24 {
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						v43 = v37 + v40<<(uint(int32(2))%32)
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
						if v8 != v44 {
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v46
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v49
							*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v49
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v53
							v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
							*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v55)
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
					v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					if v62 != 0 {
						v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v62+v63<<(uint(int32(2))%32))))
						if v8 == v67 {
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = int64(1)
						}
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = int64(1)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = int32(0)
					*(*int32)(unsafe.Add(mBase, _c_F_replication_yy_create_buffer[0])) = v23
					return v8
				}
			}
		} else {
			F_yy_fatal_error_3(m, int32(_a_F_replication_yy_create_buffer_2))
			mBase = m.M
			v78 = m.ExcPending
			if v78 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
