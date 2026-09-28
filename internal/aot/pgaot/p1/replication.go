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
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	var v26 int64
	_ = v26
	var v35 int64
	_ = v35
	var v38 int32
	_ = v38
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_GetReplicationTransferLatency[0]))
	v8 = base.AtomicRmwXchg32(m, v5, int32(1456), int32(1))
	if v8 != 0 {
		F_s_lock(m, v5+int32(1456), int32(_a_F_GetReplicationTransferLatency_0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = *(*int64)(unsafe.Add(mBase, uint32(v5)+80))
			v17 = *(*int64)(unsafe.Add(mBase, uint32(v5)+72))
			v18 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v5)+1456)), uint32(v18))
			if v16 <= v17 {
				v38 = v18
			} else {
				v26 = v16 - v17
				if base.B2i32(int64(0) < v17)^base.B2i32(v26 < v16)|base.B2i32(int64(2147483646000) < v26) != 0 {
					v38 = int32(2147483647)
				} else {
					v35 = base.I64_div_s(v26+int64(999), int64(1000))
					v38 = base.I32_wrap_i64(v35)
				}
			}
			return v38
		}
	} else {
		v16 = *(*int64)(unsafe.Add(mBase, uint32(v5)+80))
		v17 = *(*int64)(unsafe.Add(mBase, uint32(v5)+72))
		v18 = int32(0)
		atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v5)+1456)), uint32(v18))
		if v16 <= v17 {
			v38 = v18
		} else {
			v26 = v16 - v17
			if base.B2i32(int64(0) < v17)^base.B2i32(v26 < v16)|base.B2i32(int64(2147483646000) < v26) != 0 {
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
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
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
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+308))
			v22 = base.B2i32(v20 != int32(2))
			*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotDrop[0])) = uint8(v22)
			v24 = v22
		} else {
			v24 = int32(0)
		}
		v26 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDrop[2]))
		if v24 != 0 {
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+201)))
			if v27 == int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					F_errcode(m, int32(325))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
						F_errmsg(m, int32(_a_F_ReplicationSlotDrop_0), v7)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							v54 = F_errdetail(m, int32(_a_F_ReplicationSlotDrop_1), int32(0))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_ReplicationSlotDrop_2), int32(929), int32(_a_F_ReplicationSlotDrop_3))
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
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
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+88))
				*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDrop[2])) = int32(0)
				F_ReplicationSlotDropPtr(m, v26)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					if v30 != 0 {
						F_RequestDisableLogicalDecoding(m)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							m.G0 = v7 + int32(16)
							return
						}
					} else {
						m.G0 = v7 + int32(16)
						return
					}
				}
			}
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+88))
			*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDrop[2])) = int32(0)
			F_ReplicationSlotDropPtr(m, v26)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return
			} else {
				if v30 != 0 {
					F_RequestDisableLogicalDecoding(m)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						m.G0 = v7 + int32(16)
						return
					}
				} else {
					m.G0 = v7 + int32(16)
					return
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
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int64
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int64
	_ = v69
	var v72 int64
	_ = v72
	var v73 int64
	_ = v73
	var v79 int64
	_ = v79
	var v80 int32
	_ = v80
	var v84 int64
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v100 int64
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	v1 = int32(0)
	v6 = int64(0)
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeLogicalRestartLSN[0]))
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeLogicalRestartLSN[1]))
	if v10+v12 <= v1 {
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
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeLogicalRestartLSN[2]))
	v23 = F_LWLockAcquire(m, v19+int32(_a_F_ReplicationSlotsComputeLogicalRestartLSN_0), int32(1))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int64(0)
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeLogicalRestartLSN[0]))
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeLogicalRestartLSN[1]))
	if int32(0) < v28+v30 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeLogicalRestartLSN[3]))
	v37 = v35
	v38 = v1
	v43 = v6
	goto L9
L7:
	;
	v100 = v6
	goto L8
L8:
	;
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeLogicalRestartLSN[2]))
	F_LWLockRelease(m, v102+int32(_a_F_ReplicationSlotsComputeLogicalRestartLSN_0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L33
	}
L9:
	;
	v46 = v37 + v38*int32(296)
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+4)))
	if v47 != int32(1) {
		v80 = v37
		v84 = v43
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v100 = v84
	goto L8
L11:
	;
	v86 = v38 + int32(1)
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeLogicalRestartLSN[0]))
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeLogicalRestartLSN[1]))
	if v86 < v88+v90 {
		v37 = v80
		v38 = v86
		v43 = v84
		goto L9
	} else {
		goto L32
	}
L12:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+88))
	if v50 == int32(0) {
		v80 = v37
		v84 = v43
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v55 = base.AtomicRmwXchg32(m, v46, int32(0), int32(1))
	if v55 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	F_s_lock(m, v46, int32(_a_F_ReplicationSlotsComputeLogicalRestartLSN_1))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v59 = *(*int64)(unsafe.Add(mBase, uint32(v46)+280))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v46)+112))
	v61 = *(*int64)(unsafe.Add(mBase, uint32(v46)+104))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v46)+92))
	v63 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v46))), uint32(v63))
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeLogicalRestartLSN[3]))
	if v60 != 0 {
		v80 = v67
		v84 = v43
		goto L11
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	if base.Ui64(v61) < base.Ui64(v59) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v69 = v61
	goto L21
L20:
	;
	v69 = v59
	goto L21
L21:
	;
	if v59 != int64(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v72 = v69
	goto L24
L23:
	;
	v72 = v61
	goto L24
L24:
	;
	if v62 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v73 = v61
	goto L27
L26:
	;
	v73 = v72
	goto L27
L27:
	;
	if v73 == int64(0) {
		v80 = v67
		v84 = v43
		goto L11
	} else {
		goto L28
	}
L28:
	;
	if base.Ui64(v43-int64(1)) < base.Ui64(v73) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v79 = v43
	goto L31
L30:
	;
	v79 = v73
	goto L31
L31:
	;
	v80 = v67
	v84 = v79
	goto L11
L32:
	;
	goto L10
L33:
	;
	return v100
}
func F_ReplicationSlotsComputeRequiredXmin(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	v2 = int32(0)
	if l0 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredXmin[0]))
	v17 = F_LWLockAcquire(m, v13+int32(_a_F_ReplicationSlotsComputeRequiredXmin_0), int32(1))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredXmin[1]))
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredXmin[2]))
	if int32(0) < v20+v22 {
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
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredXmin[3]))
	v31 = v2
	v32 = v2
	v34 = v27
	v35 = v2
	goto L9
L7:
	;
	v104 = v2
	v105 = v2
	goto L8
L8:
	;
	v110 = m.G0
	v112 = v110 - int32(16)
	m.G0 = v112
	if l0 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L9:
	;
	v39 = v34 + v35*int32(296)
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+4)))
	if v40 != int32(1) {
		v90 = v31
		v91 = v32
		v92 = v34
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v104 = v90
	v105 = v91
	goto L8
L11:
	;
	v94 = v35 + int32(1)
	v96 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredXmin[1]))
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredXmin[2]))
	if v94 < v96+v98 {
		v31 = v90
		v32 = v91
		v34 = v92
		v35 = v94
		goto L9
	} else {
		goto L35
	}
L12:
	;
	v45 = base.AtomicRmwXchg32(m, v39, int32(0), int32(1))
	if v45 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_s_lock(m, v39, int32(_a_F_ReplicationSlotsComputeRequiredXmin_1))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v39)+112))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
	v52 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v39))), uint32(v52))
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredXmin[3]))
	if v49 != 0 {
		v90 = v31
		v91 = v32
		v92 = v56
		goto L11
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	if v51 == int32(0) {
		v72 = v31
		goto L18
	} else {
		goto L19
	}
L18:
	;
	if v50 == int32(0) {
		v90 = v72
		v91 = v32
		v92 = v56
		goto L11
	} else {
		goto L27
	}
L19:
	;
	if v31 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v72 = v51
	goto L18
L21:
	;
	v61 = int32(3)
	if base.B2i32(base.Ui32(v31) < base.Ui32(v61))|base.B2i32(base.Ui32(v51) < base.Ui32(v61)) == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	if v51-v31 < int32(0) {
		goto L20
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if base.Ui32(v31) <= base.Ui32(v51) {
		v72 = v31
		goto L18
	} else {
		goto L26
	}
L25:
	;
	v72 = v31
	goto L18
L26:
	;
	goto L20
L27:
	;
	if v32 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v90 = v72
	v91 = v50
	v92 = v56
	goto L11
L29:
	;
	v77 = int32(3)
	if base.B2i32(base.Ui32(v32) < base.Ui32(v77))|base.B2i32(base.Ui32(v50) < base.Ui32(v77)) == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if v50-v32 < int32(0) {
		goto L28
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	if base.Ui32(v32) <= base.Ui32(v50) {
		v90 = v72
		v91 = v32
		v92 = v56
		goto L11
	} else {
		goto L34
	}
L33:
	;
	v90 = v72
	v91 = v32
	v92 = v56
	goto L11
L34:
	;
	goto L28
L35:
	;
	goto L10
L36:
	;
	v140 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L4
	} else {
		goto L42
	}
L37:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredXmin[0]))
	v121 = F_LWLockAcquire(m, v117+int32(512), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L4
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredXmin[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v134)+32)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v134)+28)) = v104
	goto L36
L40:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredXmin[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v124)+32)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v124)+28)) = v104
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredXmin[0]))
	F_LWLockRelease(m, v128+int32(512))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	goto L36
L42:
	;
	if v140 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v112)+4)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v112))) = v104
	F_errmsg_internal(m, int32(_a_F_ReplicationSlotsComputeRequiredXmin_2), v112)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L4
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	m.G0 = v112 + int32(16)
	if l0 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	F_errfinish(m, int32(_a_F_ReplicationSlotsComputeRequiredXmin_3), int32(3970), int32(_a_F_ReplicationSlotsComputeRequiredXmin_4))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsComputeRequiredXmin[0]))
	F_LWLockRelease(m, v158+int32(_a_F_ReplicationSlotsComputeRequiredXmin_0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L4
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	return
L51:
	;
	goto L50
}
func F_ReplicationSlotsDropDBSlots(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
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
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsDropDBSlots[0]))
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsDropDBSlots[1]))
	if v16+v18 <= int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsDropDBSlots[2]))
	v27 = F_LWLockAcquire(m, v23+int32(_a_F_ReplicationSlotsDropDBSlots_0), int32(1))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsDropDBSlots[0]))
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsDropDBSlots[1]))
	if int32(0) < v30+v32 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v172 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v60))), uint32(v172))
	F_errstart_cold(m, int32(21), v172)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L3
	} else {
		goto L38
	}
L6:
	;
	v38 = v30
	v40 = v32
	v43 = int32(1)
	goto L10
L7:
	;
	goto L8
L8:
	;
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsDropDBSlots[2]))
	F_LWLockRelease(m, v167+int32(_a_F_ReplicationSlotsDropDBSlots_0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L3
	} else {
		goto L37
	}
L9:
	;
	F_RequestDisableLogicalDecoding(m)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L3
	} else {
		goto L36
	}
L10:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsDropDBSlots[3]))
	v47 = int32(0)
	v51 = v38
	v53 = v40
	v54 = v47
	v55 = v47
	v57 = v46
	goto L13
L11:
	;
	v153 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsDropDBSlots[2]))
	F_LWLockRelease(m, v153+int32(_a_F_ReplicationSlotsDropDBSlots_0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L3
	} else {
		goto L35
	}
L12:
	;
	v109 = base.AtomicRmwXchg32(m, v60, int32(0), int32(1))
	if v109 != 0 {
		goto L26
	} else {
		goto L27
	}
L13:
	;
	v60 = v57 + v54*int32(296)
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+4)))
	if v61 != int32(1) {
		v88 = v51
		v89 = v53
		v90 = v55
		v91 = v57
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsDropDBSlots[2]))
	F_LWLockRelease(m, v97+int32(_a_F_ReplicationSlotsDropDBSlots_0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L3
	} else {
		goto L24
	}
L15:
	;
	v93 = v54 + int32(1)
	if v93 < v88+v89 {
		v51 = v88
		v53 = v89
		v54 = v93
		v55 = v90
		v57 = v91
		goto L13
	} else {
		goto L23
	}
L16:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v60)+88))
	if v64 == int32(0) {
		v88 = v51
		v89 = v53
		v90 = v55
		v91 = v57
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v69 = base.AtomicRmwXchg32(m, v60, int32(0), int32(1))
	if v69 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_s_lock(m, v60, int32(_a_F_ReplicationSlotsDropDBSlots_1))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L3
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v60)+112))
	v74 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v60))), uint32(v74))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v60)+88))
	if v77 == l0 {
		goto L12
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsDropDBSlots[0]))
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsDropDBSlots[1]))
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsDropDBSlots[3]))
	v88 = v83
	v89 = v85
	v90 = v55 | base.B2i32(v73 == int32(0))
	v91 = v87
	goto L15
L23:
	;
	goto L14
L24:
	;
	if (v90|v43)&int32(1) == int32(0) {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	goto L1
L26:
	;
	F_s_lock(m, v60, int32(_a_F_ReplicationSlotsDropDBSlots_1))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L3
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	if v113 != int32(-1) {
		goto L5
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	v116 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsDropDBSlots[4])) = v60
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsDropDBSlots[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+8)) = v120
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v60))), uint32(v116))
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsDropDBSlots[2]))
	F_LWLockRelease(m, v126+int32(_a_F_ReplicationSlotsDropDBSlots_0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L3
	} else {
		goto L31
	}
L31:
	;
	v131 = int32(_a_F_ReplicationSlotsDropDBSlots_2)
	v132 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsDropDBSlots[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsDropDBSlots[4])) = int32(0)
	F_ReplicationSlotDropPtr(m, v132)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L3
	} else {
		goto L32
	}
L32:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsDropDBSlots[2]))
	v143 = F_LWLockAcquire(m, v139+int32(_a_F_ReplicationSlotsDropDBSlots_0), int32(1))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	v146 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsDropDBSlots[0]))
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsDropDBSlots[1]))
	if int32(0) < v146+v148 {
		v38 = v146
		v40 = v148
		v43 = v116
		goto L10
	} else {
		goto L34
	}
L34:
	;
	goto L11
L35:
	;
	goto L9
L36:
	;
	goto L1
L37:
	;
	goto L1
L38:
	;
	F_errcode(m, int32(100663621))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L3
	} else {
		goto L39
	}
L39:
	;
	v183 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsDropDBSlots[6]))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v184+v113*int32(768))+12))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v60 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v188
	F_errmsg(m, int32(_a_F_ReplicationSlotsDropDBSlots_3), v12)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L3
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_ReplicationSlotsDropDBSlots_4), int32(1595), int32(_a_F_ReplicationSlotsDropDBSlots_5))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ReplicationSlotsShmemInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsShmemInit[0]))
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsShmemInit[1]))
	if v2 < v5+v7 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v11 = v2
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsShmemInit[2]))
	v17 = v14 + v11*int32(296)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = int32(-1)
	v20 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v17))), uint32(v20))
	F_LWLockInitialize(m, v17+int32(208), int32(68))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
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
	v29 = v17 + int32(224)
	v30 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v29))), uint32(v30))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+4)) = int64(-1)
	goto L8
L8:
	;
	v36 = v11 + int32(1)
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsShmemInit[0]))
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsShmemInit[1]))
	if v36 < v38+v40 {
		v11 = v36
		goto L4
	} else {
		goto L9
	}
L9:
	;
	goto L5
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
