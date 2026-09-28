package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_CheckLogicalSlotExists(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
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
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	v1 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_CheckLogicalSlotExists[0]))
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_CheckLogicalSlotExists[1]))
	if v1 < v7+v9 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_CheckLogicalSlotExists[2]))
	v18 = F_LWLockAcquire(m, v14+int32(_a_F_CheckLogicalSlotExists_0), int32(1))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v87 = v1
	goto L3
L3:
	;
	return v87
L4:
	;
	return int32(0)
L5:
	;
	v22 = int32(0)
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_CheckLogicalSlotExists[0]))
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_CheckLogicalSlotExists[1]))
	if v24+v26 <= v22 {
		v78 = v22
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_CheckLogicalSlotExists[2]))
	F_LWLockRelease(m, v80+int32(_a_F_CheckLogicalSlotExists_0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L4
	} else {
		goto L19
	}
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_CheckLogicalSlotExists[3]))
	v33 = v24
	v34 = v1
	v35 = v26
	v36 = v31
	goto L8
L8:
	;
	v39 = v36 + v34*int32(296)
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+4)))
	if v40 != int32(1) {
		v65 = v33
		v66 = v35
		v67 = v36
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v78 = int32(0)
	goto L6
L10:
	;
	v69 = v34 + int32(1)
	if v69 < v65+v66 {
		v33 = v65
		v34 = v69
		v35 = v66
		v36 = v67
		goto L8
	} else {
		goto L18
	}
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+88))
	if v43 == int32(0) {
		v65 = v33
		v66 = v35
		v67 = v36
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v48 = base.AtomicRmwXchg32(m, v39, int32(0), int32(1))
	if v48 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_s_lock(m, v39, int32(_a_F_CheckLogicalSlotExists_1))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v39)+112))
	v53 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v39))), uint32(v53))
	if v52 == v53 {
		v78 = int32(1)
		goto L6
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_CheckLogicalSlotExists[0]))
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_CheckLogicalSlotExists[1]))
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_CheckLogicalSlotExists[3]))
	v65 = v60
	v66 = v62
	v67 = v64
	goto L10
L18:
	;
	goto L9
L19:
	;
	v87 = v78
	goto L3
}
func F_EnableLogicalDecoding(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v95 int64
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_EnableLogicalDecoding[0]))
	v13 = F_LWLockAcquire(m, v9+int32(_a_F_EnableLogicalDecoding_0), int32(0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_EnableLogicalDecoding[1]))
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
		if v17 == int32(1) {
			v20 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v16)+2)) = uint8(v20)
			v23 = *(*int32)(unsafe.Add(mBase, _c_F_EnableLogicalDecoding[0]))
			F_LWLockRelease(m, v23+int32(_a_F_EnableLogicalDecoding_0))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				m.G0 = v6 + int32(16)
				return
			}
		} else {
			v28 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v28)
			v31 = *(*int32)(unsafe.Add(mBase, _c_F_EnableLogicalDecoding[0]))
			F_LWLockRelease(m, v31+int32(_a_F_EnableLogicalDecoding_0))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return
			} else {
				v37 = F_EmitProcSignalBarrier(m, int32(1))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					F_WaitForProcSignalBarrier(m, v37)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						v43 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_EnableLogicalDecoding[2])))
						if v43 == int32(1) {
							v48 = *(*int32)(unsafe.Add(mBase, _c_F_EnableLogicalDecoding[3]))
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+308))
							v51 = base.B2i32(v49 != int32(2))
							*(*uint8)(unsafe.Add(mBase, _c_F_EnableLogicalDecoding[2])) = uint8(v51)
							v53 = v51
						} else {
							v53 = int32(0)
						}
						v55 = *(*int32)(unsafe.Add(mBase, _c_F_EnableLogicalDecoding[0]))
						v59 = F_LWLockAcquire(m, v55+int32(_a_F_EnableLogicalDecoding_0), int32(0))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return
						} else {
							v62 = *(*int32)(unsafe.Add(mBase, _c_F_EnableLogicalDecoding[1]))
							v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)))
							if v63 == int32(1) {
								v66 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v62)+2)) = uint8(v66)
								v69 = *(*int32)(unsafe.Add(mBase, _c_F_EnableLogicalDecoding[0]))
								F_LWLockRelease(m, v69+int32(_a_F_EnableLogicalDecoding_0))
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return
								} else {
									m.G0 = v6 + int32(16)
									return
								}
							} else {
								v74 = int32(_a_F_EnableLogicalDecoding_1)
								v76 = *(*int32)(unsafe.Add(mBase, _c_F_EnableLogicalDecoding[4]))
								v77 = int32(1)
								*(*int32)(unsafe.Add(mBase, _c_F_EnableLogicalDecoding[4])) = v76 + v77
								*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)) = uint8(v77)
								if v53 == int32(0) {
									v84 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v6)+15)) = uint8(v84)
									F_XLogBeginInsert(m)
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return
									} else {
										F_XLogRegisterData(m, v6+int32(15), int32(1))
										mBase = m.M
										v92 = m.ExcPending
										if v92 != 0 {
											return
										} else {
											v95 = F_XLogInsert(m, int32(0), int32(240))
											mBase = m.M
											v96 = m.ExcPending
											if v96 != 0 {
												return
											} else {
												F_XLogFlush(m, v95)
												mBase = m.M
												v98 = m.ExcPending
												if v98 != 0 {
													return
												} else {
													v100 = *(*int32)(unsafe.Add(mBase, _c_F_EnableLogicalDecoding[1]))
													v101 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v100)+2)) = uint8(v101)
													v103 = int32(_a_F_EnableLogicalDecoding_1)
													v105 = *(*int32)(unsafe.Add(mBase, _c_F_EnableLogicalDecoding[4]))
													*(*int32)(unsafe.Add(mBase, _c_F_EnableLogicalDecoding[4])) = v105 - int32(1)
													v110 = *(*int32)(unsafe.Add(mBase, _c_F_EnableLogicalDecoding[0]))
													F_LWLockRelease(m, v110+int32(_a_F_EnableLogicalDecoding_0))
													mBase = m.M
													v114 = m.ExcPending
													if v114 != 0 {
														return
													} else {
														v117 = F_errstart(m, int32(15), int32(0))
														mBase = m.M
														v118 = m.ExcPending
														if v118 != 0 {
															return
														} else {
															if v117 == int32(0) {
																m.G0 = v6 + int32(16)
																return
															} else {
																F_errmsg(m, int32(_a_F_EnableLogicalDecoding_2), int32(0))
																mBase = m.M
																v124 = m.ExcPending
																if v124 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(_a_F_EnableLogicalDecoding_3), int32(442), int32(_a_F_EnableLogicalDecoding_4))
																	mBase = m.M
																	v129 = m.ExcPending
																	if v129 != 0 {
																		return
																	} else {
																		m.G0 = v6 + int32(16)
																		return
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
									v131 = *(*int32)(unsafe.Add(mBase, _c_F_EnableLogicalDecoding[1]))
									v132 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v131)+2)) = uint8(v132)
									v134 = int32(_a_F_EnableLogicalDecoding_1)
									v136 = *(*int32)(unsafe.Add(mBase, _c_F_EnableLogicalDecoding[4]))
									*(*int32)(unsafe.Add(mBase, _c_F_EnableLogicalDecoding[4])) = v136 - int32(1)
									v141 = *(*int32)(unsafe.Add(mBase, _c_F_EnableLogicalDecoding[0]))
									F_LWLockRelease(m, v141+int32(_a_F_EnableLogicalDecoding_0))
									mBase = m.M
									v145 = m.ExcPending
									if v145 != 0 {
										return
									} else {
										m.G0 = v6 + int32(16)
										return
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
func F_LogicalConfirmReceivedLocation(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v15 int64
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v29 int64
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int64
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int64
	_ = v54
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v61 int64
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int64
	_ = v121
	var v124 int32
	_ = v124
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_LogicalConfirmReceivedLocation[0]))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)+240))
	if v12 == int64(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v8 + int32(16)
	return
L2:
	;
	v115 = base.AtomicRmwXchg32(m, v11, int32(0), int32(1))
	if v115 != 0 {
		goto L42
	} else {
		goto L43
	}
L3:
	;
	v15 = *(*int64)(unsafe.Add(mBase, uint32(v11)+248))
	if v15 == int64(0) {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v20 = base.AtomicRmwXchg32(m, v11, int32(0), int32(1))
	if v20 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L5
L7:
	;
	F_s_lock(m, v11, int32(_a_F_LogicalConfirmReceivedLocation_0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_LogicalConfirmReceivedLocation[0]))
	v26 = *(*int64)(unsafe.Add(mBase, uint32(v25)+120))
	if base.Ui64(v26) < base.Ui64(l0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	return
L11:
	;
	goto L9
L12:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v25)+120)) = l0
	goto L14
L13:
	;
	goto L14
L14:
	;
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v25)+240))
	if base.Ui64(l0) <= base.Ui64(v29-int64(1)) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v69 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v25))), uint32(v69))
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L10
	} else {
		goto L24
	}
L16:
	;
	v59 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v25)+248)) = v59
	v61 = *(*int64)(unsafe.Add(mBase, uint32(v25)+256))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+256)) = v59
	*(*int64)(unsafe.Add(mBase, uint32(v25)+104)) = v61
	v67 = v58
	v68 = int32(1)
	goto L15
L17:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v25)+240)) = int64(0)
	v49 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+236)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v25)+100)) = v33
	v52 = int32(1)
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v25)+248))
	if base.Ui64(l0) <= base.Ui64(v54-int64(1)) {
		v67 = v52
		v68 = v49
		goto L15
	} else {
		goto L23
	}
L18:
	;
	v40 = *(*int64)(unsafe.Add(mBase, uint32(v25)+248))
	if base.Ui64(v40-int64(1)) < base.Ui64(l0) {
		v58 = int32(0)
		goto L16
	} else {
		goto L22
	}
L19:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v25)+236))
	if v33 == int32(0) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v25)+100))
	if v36 != v33 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	goto L18
L22:
	;
	v44 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v25))), uint32(v44))
	goto L1
L23:
	;
	v58 = v52
	goto L16
L24:
	;
	F_ReplicationSlotSave(m)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L10
	} else {
		goto L25
	}
L25:
	;
	v78 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L10
	} else {
		goto L26
	}
L26:
	;
	if v78 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v67
	F_errmsg_internal(m, int32(_a_F_LogicalConfirmReceivedLocation_1), v8)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L10
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if v67 != 0 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	F_errfinish(m, int32(_a_F_LogicalConfirmReceivedLocation_2), int32(1986), int32(_a_F_LogicalConfirmReceivedLocation_3))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L10
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_LogicalConfirmReceivedLocation[0]))
	v94 = base.AtomicRmwXchg32(m, v91, int32(0), int32(1))
	if v94 != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	if v68 == int32(0) {
		goto L1
	} else {
		goto L40
	}
L35:
	;
	F_s_lock(m, v91, int32(_a_F_LogicalConfirmReceivedLocation_0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L10
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_LogicalConfirmReceivedLocation[0]))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v99)+20)) = v100
	v102 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v99))), uint32(v102))
	F_ReplicationSlotsComputeRequiredXmin(m, v102)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L10
	} else {
		goto L39
	}
L38:
	;
	goto L37
L39:
	;
	goto L34
L40:
	;
	F_ReplicationSlotsComputeRequiredLSN(m)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L10
	} else {
		goto L41
	}
L41:
	;
	goto L1
L42:
	;
	F_s_lock(m, v11, int32(_a_F_LogicalConfirmReceivedLocation_0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L10
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_LogicalConfirmReceivedLocation[0]))
	v121 = *(*int64)(unsafe.Add(mBase, uint32(v120)+120))
	if base.Ui64(v121) < base.Ui64(l0) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	goto L44
L46:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v120)+120)) = l0
	goto L48
L47:
	;
	goto L48
L48:
	;
	v124 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v120))), uint32(v124))
	goto L1
}
func F_LogicalTapeSetBlocks(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int64
	_ = v3
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	return v2 - v3
}
func F_abort_logical_decoding_activation(m *base.Module, l0 int32, l1 int64) {
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v5 = F_errstart(m, int32(14), int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		if v5 != 0 {
			F_errmsg_internal(m, int32(_a_F_abort_logical_decoding_activation_0), int32(0))
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_abort_logical_decoding_activation_1), int32(270), int32(_a_F_abort_logical_decoding_activation_2))
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					F_RequestDisableLogicalDecoding(m)
					v17 = m.ExcPending
					if v17 != 0 {
						return
					} else {
						return
					}
				}
			}
		} else {
			F_RequestDisableLogicalDecoding(m)
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				return
			}
		}
	}
}
