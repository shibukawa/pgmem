package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetReplicationTransferLatency(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v26 int64
	_ = v26
	var v35 int64
	_ = v35
	var v38 int32
	_ = v38
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_GetReplicationTransferLatency[0]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+1456))
	*(*int32)(unsafe.Add(mBase, uint32(v4)+1456)) = int32(1)
	if v5 != 0 {
		F_s_lock(m, v4+int32(1456), int32(_a_F_GetReplicationTransferLatency_0), int32(401), int32(_a_F_GetReplicationTransferLatency_1))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+1456)) = v17
			v19 = *(*int64)(unsafe.Add(mBase, uint32(v4)+72))
			v20 = *(*int64)(unsafe.Add(mBase, uint32(v4)+80))
			if v20 <= v19 {
				v38 = v17
			} else {
				v26 = v20 - v19
				if base.B2i32(int64(0) < v19)^base.B2i32(v26 < v20)|base.B2i32(int64(2147483646000) < v26) != 0 {
					v38 = int32(2147483647)
				} else {
					v35 = base.I64_div_s(v26+int64(999), int64(1000))
					v38 = base.I32_wrap_i64(v35)
				}
			}
			return v38
		}
	} else {
		v17 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+1456)) = v17
		v19 = *(*int64)(unsafe.Add(mBase, uint32(v4)+72))
		v20 = *(*int64)(unsafe.Add(mBase, uint32(v4)+80))
		if v20 <= v19 {
			v38 = v17
		} else {
			v26 = v20 - v19
			if base.B2i32(int64(0) < v19)^base.B2i32(v26 < v20)|base.B2i32(int64(2147483646000) < v26) != 0 {
				v38 = int32(2147483647)
			} else {
				v35 = base.I64_div_s(v26+int64(999), int64(1000))
				v38 = base.I32_wrap_i64(v35)
			}
		}
		return v38
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
	var v4 int64
	_ = v4
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int64
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v59 int64
	_ = v59
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v64 int64
	_ = v64
	var v70 int64
	_ = v70
	var v71 int32
	_ = v71
	var v73 int64
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v84 int64
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	v1 = int32(0)
	v4 = int64(0)
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeLogicalRestartLSN[0]))
	if v8 <= v1 {
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
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeLogicalRestartLSN[1]))
	v18 = F_LWLockAcquire(m, v14+int32(_a_F_ReplicationSlotsComputeLogicalRestartLSN_0), int32(1))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int64(0)
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeLogicalRestartLSN[0]))
	if int32(0) < v23 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeLogicalRestartLSN[2]))
	v29 = v27
	v30 = v1
	v32 = v4
	goto L9
L7:
	;
	v84 = v4
	goto L8
L8:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeLogicalRestartLSN[1]))
	F_LWLockRelease(m, v87+int32(_a_F_ReplicationSlotsComputeLogicalRestartLSN_0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L4
	} else {
		goto L33
	}
L9:
	;
	v36 = v29 + v30*int32(288)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+4)))
	if v37 != int32(1) {
		v71 = v29
		v73 = v32
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v84 = v73
	goto L8
L11:
	;
	v76 = v30 + int32(1)
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeLogicalRestartLSN[0]))
	if v76 < v78 {
		v29 = v71
		v30 = v76
		v32 = v73
		goto L9
	} else {
		goto L32
	}
L12:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36)+88))
	if v40 == int32(0) {
		v71 = v29
		v73 = v32
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = int32(1)
	if v43 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	F_s_lock(m, v36, int32(_a_F_ReplicationSlotsComputeLogicalRestartLSN_1), int32(1324), int32(_a_F_ReplicationSlotsComputeLogicalRestartLSN_2))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = int32(0)
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeLogicalRestartLSN[2]))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v36)+112))
	if v55 != 0 {
		v71 = v54
		v73 = v32
		goto L11
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v36)+104))
	v57 = *(*int64)(unsafe.Add(mBase, uint32(v36)+280))
	if base.Ui64(v56) < base.Ui64(v57) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v59 = v56
	goto L21
L20:
	;
	v59 = v57
	goto L21
L21:
	;
	if v57 != int64(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v62 = v59
	goto L24
L23:
	;
	v62 = v56
	goto L24
L24:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v36)+92))
	if v63 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v64 = v56
	goto L27
L26:
	;
	v64 = v62
	goto L27
L27:
	;
	if v64 == int64(0) {
		v71 = v54
		v73 = v32
		goto L11
	} else {
		goto L28
	}
L28:
	;
	if base.Ui64(v32-int64(1)) < base.Ui64(v64) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v70 = v32
	goto L31
L30:
	;
	v70 = v64
	goto L31
L31:
	;
	v71 = v54
	v73 = v70
	goto L11
L32:
	;
	goto L10
L33:
	;
	return v84
}
func F_ReplicationSlotsComputeRequiredXmin(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	v2 = int32(0)
	if l0 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredXmin[0]))
	v14 = F_LWLockAcquire(m, v10+int32(_a_F_ReplicationSlotsComputeRequiredXmin_0), int32(1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredXmin[1]))
	if int32(0) < v17 {
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
	v23 = v2
	v24 = v2
	v25 = v2
	goto L9
L7:
	;
	v92 = v2
	v93 = v2
	goto L8
L8:
	;
	v95 = m.G0
	v97 = v95 - int32(16)
	m.G0 = v97
	if l0 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredXmin[2]))
	v30 = v27 + v25*int32(288)
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+4)))
	if v31 != int32(1) {
		v82 = v23
		v83 = v24
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v92 = v82
	v93 = v83
	goto L8
L11:
	;
	v85 = v25 + int32(1)
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredXmin[1]))
	if v85 < v87 {
		v23 = v82
		v24 = v83
		v25 = v85
		goto L9
	} else {
		goto L37
	}
L12:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = int32(1)
	if v34 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_s_lock(m, v30, int32(_a_F_ReplicationSlotsComputeRequiredXmin_1), int32(1188), int32(_a_F_ReplicationSlotsComputeRequiredXmin_2))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = int32(0)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v30)+112))
	if v44 != 0 {
		v82 = v23
		v83 = v24
		goto L11
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	if v46 == int32(0) {
		v63 = v23
		goto L18
	} else {
		goto L19
	}
L18:
	;
	if v45 == int32(0) {
		v82 = v63
		v83 = v24
		goto L11
	} else {
		goto L28
	}
L19:
	;
	if v23 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v23))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v46)) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	goto L22
L22:
	;
	v63 = v46
	goto L18
L23:
	;
	if v60 == int32(0) {
		v63 = v23
		goto L18
	} else {
		goto L27
	}
L24:
	;
	v60 = base.B2i32(base.Ui32(v46) < base.Ui32(v23))
	goto L23
L25:
	;
	goto L26
L26:
	;
	v60 = int32(base.Ui32(v46-v23) >> (uint(int32(31)) % 32))
	goto L23
L27:
	;
	goto L22
L28:
	;
	if v24 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v24))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v45)) == int32(0) {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	goto L31
L31:
	;
	v82 = v63
	v83 = v45
	goto L11
L32:
	;
	if v77 == int32(0) {
		v82 = v63
		v83 = v24
		goto L11
	} else {
		goto L36
	}
L33:
	;
	v77 = base.B2i32(base.Ui32(v45) < base.Ui32(v24))
	goto L32
L34:
	;
	goto L35
L35:
	;
	v77 = int32(base.Ui32(v45-v24) >> (uint(int32(31)) % 32))
	goto L32
L36:
	;
	goto L31
L37:
	;
	goto L10
L38:
	;
	v125 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L4
	} else {
		goto L44
	}
L39:
	;
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredXmin[0]))
	v106 = F_LWLockAcquire(m, v102+int32(512), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredXmin[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v119)+32)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v119)+28)) = v92
	goto L38
L42:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredXmin[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v109)+32)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v109)+28)) = v92
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredXmin[0]))
	F_LWLockRelease(m, v113+int32(512))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	goto L38
L44:
	;
	if v125 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+4)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = v92
	F_errmsg_internal(m, int32(_a_F_ReplicationSlotsComputeRequiredXmin_3), v97)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L4
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	m.G0 = v97 + int32(16)
	if l0 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	F_errfinish(m, int32(_a_F_ReplicationSlotsComputeRequiredXmin_4), int32(3958), int32(_a_F_ReplicationSlotsComputeRequiredXmin_5))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v143 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredXmin[0]))
	F_LWLockRelease(m, v143+int32(_a_F_ReplicationSlotsComputeRequiredXmin_0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
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
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
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
	v122 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v46))) = v122
	F_errstart_cold(m, int32(21), v122)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
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
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsDropDBSlots[1]))
	F_LWLockRelease(m, v107+int32(_a_F_ReplicationSlotsDropDBSlots_0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
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
	v97 = v39 + int32(1)
	if v97 != v31 {
		v39 = v97
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
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	*(*int32)(unsafe.Add(mBase, uint32(v46))) = int32(1)
	if v55 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v46))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+8)) = v67
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsDropDBSlots[1]))
	F_LWLockRelease(m, v72+int32(_a_F_ReplicationSlotsDropDBSlots_0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	v77 = int32(_a_F_ReplicationSlotsDropDBSlots_3)
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsDropDBSlots[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsDropDBSlots[3])) = int32(0)
	F_ReplicationSlotDropPtr(m, v78)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsDropDBSlots[1]))
	v89 = F_LWLockAcquire(m, v85+int32(_a_F_ReplicationSlotsDropDBSlots_0), int32(1))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsDropDBSlots[0]))
	if int32(0) < v92 {
		v31 = v92
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
	v130 = m.ExcPending
	if v130 != 0 {
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
	v137 = m.ExcPending
	if v137 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_ReplicationSlotsDropDBSlots_1), int32(1500), int32(_a_F_ReplicationSlotsDropDBSlots_2))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
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
