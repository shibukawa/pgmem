package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"sync/atomic"
	"unsafe"
)

func F_GetAdditionalLocalPinLimit(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_GetAdditionalLocalPinLimit[0]))
	v6 = base.I32_div_s(v4, int32(4))
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_GetAdditionalLocalPinLimit[1]))
	v9 = v6 - v8
	if base.Ui32(v9) <= base.Ui32(v6) {
		v12 = v9
	} else {
		v12 = int32(0)
	}
	return v12
}
func F_GetCTEForRTE(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v14 = l0
	v16 = v12 + l2
	goto L2
L1:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	if v44 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	if v16 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v25 != 0 {
		v14 = v25
		v16 = v16 - int32(1)
		goto L2
	} else {
		goto L5
	}
L5:
	;
	goto L3
L6:
	;
	return int32(0)
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v32
	F_errmsg_internal(m, int32(_a_F_GetCTEForRTE_0), v10+int32(16))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	F_errfinish(m, int32(_a_F_GetCTEForRTE_1), int32(604), int32(_a_F_GetCTEForRTE_2))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L6
	} else {
		goto L29
	}
L11:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v47 <= int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v50 = int32(0)
	if v50 < v47 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v54 = v47
	goto L15
L14:
	;
	v54 = v50
	goto L15
L15:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	v59 = v50
	goto L16
L16:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v56+v59<<(uint(int32(2))%32))))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if base.B2i32(v71 == int32(0))|base.B2i32(v71 != v74) != 0 {
		v92 = v71
		v93 = v74
		goto L19
	} else {
		goto L20
	}
L17:
	;
	m.G0 = v10 + int32(32)
	return v67
L18:
	;
	if v92-v93 != 0 {
		goto L25
	} else {
		goto L26
	}
L19:
	;
	goto L18
L20:
	;
	v77 = v68
	v78 = v55
	goto L21
L21:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+1)))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+1)))
	if v82 == int32(0) {
		v92 = v82
		v93 = v81
		goto L19
	} else {
		goto L23
	}
L22:
	;
	v92 = v82
	v93 = v81
	goto L19
L23:
	;
	v85 = int32(1)
	if v82 == v81 {
		v77 = v77 + v85
		v78 = v78 + v85
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v96 = v59 + int32(1)
	if v54 != v96 {
		v59 = v96
		goto L16
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	goto L17
L28:
	;
	goto L10
L29:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v113
	F_errmsg_internal(m, int32(_a_F_GetCTEForRTE_3), v10)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_GetCTEForRTE_1), int32(614), int32(_a_F_GetCTEForRTE_2))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_GetFdwRoutineForRelation(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	if v4 == int32(0) {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		v8 = F_GetFdwRoutineByRelId(m, v7)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, _c_F_GetFdwRoutineForRelation[0]))
			v15 = F_MemoryContextAlloc(m, v13, int32(188))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				base.MemoryCopy(m, v15, v8, int32(188))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+260)) = v15
				return v8
			}
		}
	} else {
		if l1 != 0 {
			v22 = F_palloc(m, int32(188))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
				base.MemoryCopy(m, v22, v24, int32(188))
				v27 = v22
				return v27
			}
		} else {
			v27 = v4
			return v27
		}
	}
}
func F_GetFlushRecPtr(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v8 int64
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int64
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int64
	_ = v27
	v3 = int32(_a_F_GetFlushRecPtr_0)
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_GetFlushRecPtr[0]))
	v5 = int64(0)
	v8 = base.AtomicRmwCmpxchg64(m, v4, int32(272), v5, v5)
	*(*int64)(unsafe.Add(mBase, _c_F_GetFlushRecPtr[1])) = v8
	v10 = int32(0)
	v13 = base.AtomicRmwOr32(m, v10, int32(_a_F_GetFlushRecPtr_1), v10)
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_GetFlushRecPtr[0]))
	v20 = base.AtomicRmwCmpxchg64(m, v16, int32(264), v5, v5)
	*(*int64)(unsafe.Add(mBase, _c_F_GetFlushRecPtr[2])) = v20
	if l0 != 0 {
		v23 = *(*int32)(unsafe.Add(mBase, _c_F_GetFlushRecPtr[0]))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+300))
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v24
	} else {
	}
	v27 = *(*int64)(unsafe.Add(mBase, _c_F_GetFlushRecPtr[1]))
	return v27
}
func F_GetLastImportantRecPtr(m *base.Module) int64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
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
	var v31 int64
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int64
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int64
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int64
	_ = v101
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int64
	_ = v115
	var v119 int32
	_ = v119
	var v121 int64
	_ = v121
	var v123 int64
	_ = v123
	var v125 int64
	_ = v125
	var v127 int64
	_ = v127
	var v129 int64
	_ = v129
	var v131 int64
	_ = v131
	var v133 int64
	_ = v133
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_GetLastImportantRecPtr[0]))
	v13 = F_LWLockAcquire(m, v11, int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_GetLastImportantRecPtr[0]))
		v19 = *(*int64)(unsafe.Add(mBase, uint32(v18)+24))
		F_LWLockRelease(m, v18)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int64(0)
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, _c_F_GetLastImportantRecPtr[0]))
			v27 = F_LWLockAcquire(m, v23+int32(128), int32(0))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int64(0)
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, _c_F_GetLastImportantRecPtr[0]))
				v31 = *(*int64)(unsafe.Add(mBase, uint32(v30)+152))
				F_LWLockRelease(m, v30+int32(128))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int64(0)
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, _c_F_GetLastImportantRecPtr[0]))
					v41 = F_LWLockAcquire(m, v37+int32(256), int32(0))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int64(0)
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, _c_F_GetLastImportantRecPtr[0]))
						v45 = *(*int64)(unsafe.Add(mBase, uint32(v44)+280))
						F_LWLockRelease(m, v44+int32(256))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int64(0)
						} else {
							v51 = *(*int32)(unsafe.Add(mBase, _c_F_GetLastImportantRecPtr[0]))
							v55 = F_LWLockAcquire(m, v51+int32(384), int32(0))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int64(0)
							} else {
								v58 = *(*int32)(unsafe.Add(mBase, _c_F_GetLastImportantRecPtr[0]))
								v59 = *(*int64)(unsafe.Add(mBase, uint32(v58)+408))
								F_LWLockRelease(m, v58+int32(384))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int64(0)
								} else {
									v65 = *(*int32)(unsafe.Add(mBase, _c_F_GetLastImportantRecPtr[0]))
									v69 = F_LWLockAcquire(m, v65+int32(512), int32(0))
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int64(0)
									} else {
										v72 = *(*int32)(unsafe.Add(mBase, _c_F_GetLastImportantRecPtr[0]))
										v73 = *(*int64)(unsafe.Add(mBase, uint32(v72)+536))
										F_LWLockRelease(m, v72+int32(512))
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
											return int64(0)
										} else {
											v79 = *(*int32)(unsafe.Add(mBase, _c_F_GetLastImportantRecPtr[0]))
											v83 = F_LWLockAcquire(m, v79+int32(640), int32(0))
											mBase = m.M
											v84 = m.ExcPending
											if v84 != 0 {
												return int64(0)
											} else {
												v86 = *(*int32)(unsafe.Add(mBase, _c_F_GetLastImportantRecPtr[0]))
												v87 = *(*int64)(unsafe.Add(mBase, uint32(v86)+664))
												F_LWLockRelease(m, v86+int32(640))
												mBase = m.M
												v91 = m.ExcPending
												if v91 != 0 {
													return int64(0)
												} else {
													v93 = *(*int32)(unsafe.Add(mBase, _c_F_GetLastImportantRecPtr[0]))
													v97 = F_LWLockAcquire(m, v93+int32(768), int32(0))
													mBase = m.M
													v98 = m.ExcPending
													if v98 != 0 {
														return int64(0)
													} else {
														v100 = *(*int32)(unsafe.Add(mBase, _c_F_GetLastImportantRecPtr[0]))
														v101 = *(*int64)(unsafe.Add(mBase, uint32(v100)+792))
														F_LWLockRelease(m, v100+int32(768))
														mBase = m.M
														v105 = m.ExcPending
														if v105 != 0 {
															return int64(0)
														} else {
															v107 = *(*int32)(unsafe.Add(mBase, _c_F_GetLastImportantRecPtr[0]))
															v111 = F_LWLockAcquire(m, v107+int32(896), int32(0))
															mBase = m.M
															v112 = m.ExcPending
															if v112 != 0 {
																return int64(0)
															} else {
																v114 = *(*int32)(unsafe.Add(mBase, _c_F_GetLastImportantRecPtr[0]))
																v115 = *(*int64)(unsafe.Add(mBase, uint32(v114)+920))
																F_LWLockRelease(m, v114+int32(896))
																mBase = m.M
																v119 = m.ExcPending
																if v119 != 0 {
																	return int64(0)
																} else {
																	if base.Ui64(v31) < base.Ui64(v19) {
																		v121 = v19
																	} else {
																		v121 = v31
																	}
																	if base.Ui64(v45) < base.Ui64(v121) {
																		v123 = v121
																	} else {
																		v123 = v45
																	}
																	if base.Ui64(v59) < base.Ui64(v123) {
																		v125 = v123
																	} else {
																		v125 = v59
																	}
																	if base.Ui64(v73) < base.Ui64(v125) {
																		v127 = v125
																	} else {
																		v127 = v73
																	}
																	if base.Ui64(v87) < base.Ui64(v127) {
																		v129 = v127
																	} else {
																		v129 = v87
																	}
																	if base.Ui64(v101) < base.Ui64(v129) {
																		v131 = v129
																	} else {
																		v131 = v101
																	}
																	if base.Ui64(v115) < base.Ui64(v131) {
																		v133 = v131
																	} else {
																		v133 = v115
																	}
																	return v133
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
				}
			}
		}
	}
}
func F_GetLatestLSN(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v31 int64
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int64
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int64
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int64
	_ = v57
	var v58 int32
	_ = v58
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v66 int64
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int64
	_ = v74
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetLatestLSN[0])))
	if v12 == int32(1) {
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_GetLatestLSN[1]))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+308))
		v20 = base.B2i32(v18 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _c_F_GetLatestLSN[0])) = uint8(v20)
		v22 = v20
	} else {
		v22 = int32(0)
	}
	if v22 == int32(0) {
		v26 = int32(_a_F_GetLatestLSN_0)
		v27 = *(*int32)(unsafe.Add(mBase, _c_F_GetLatestLSN[1]))
		v28 = int64(0)
		v31 = base.AtomicRmwCmpxchg64(m, v27, int32(272), v28, v28)
		*(*int64)(unsafe.Add(mBase, _c_F_GetLatestLSN[2])) = v31
		v33 = int32(0)
		v36 = base.AtomicRmwOr32(m, v33, int32(_a_F_GetLatestLSN_1), v33)
		v39 = *(*int32)(unsafe.Add(mBase, _c_F_GetLatestLSN[1]))
		v43 = base.AtomicRmwCmpxchg64(m, v39, int32(264), v28, v28)
		*(*int64)(unsafe.Add(mBase, _c_F_GetLatestLSN[3])) = v43
		if l0 != 0 {
			v46 = *(*int32)(unsafe.Add(mBase, _c_F_GetLatestLSN[1]))
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+300))
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v47
		} else {
		}
		v50 = *(*int64)(unsafe.Add(mBase, _c_F_GetLatestLSN[2]))
		v74 = v50
		m.G0 = v8 + int32(16)
		return v74
	} else {
		v51 = F_GetWALInsertionTimeLineIfSet(m)
		mBase = m.M
		v54 = m.ExcPending
		if v54 != 0 {
			return int64(0)
		} else {
			if v51 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v51
				v57 = F_GetXLogReplayRecPtr(m, int32(0))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return int64(0)
				} else {
					v74 = v57
					m.G0 = v8 + int32(16)
					return v74
				}
			} else {
				v62 = F_GetWalRcvFlushRecPtr(m, int32(0), v8+int32(12))
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return int64(0)
				} else {
					v66 = F_GetXLogReplayRecPtr(m, v8+int32(8))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int64(0)
					} else {
						if base.Ui64(v66) < base.Ui64(v62) {
							v69 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v69
							v74 = v62
						} else {
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v71
							v74 = v66
						}
						m.G0 = v8 + int32(16)
						return v74
					}
				}
			}
		}
	}
}
func F_GetLockConflicts(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
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
	var v157 int64
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v181 int64
	_ = v181
	var v187 int32
	_ = v187
	var v189 int64
	_ = v189
	var v195 int32
	_ = v195
	var v198 int64
	_ = v198
	var v203 int32
	_ = v203
	var v205 int64
	_ = v205
	var v211 int32
	_ = v211
	var v214 int64
	_ = v214
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v241 int32
	_ = v241
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v353 int32
	_ = v353
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v420 int32
	_ = v420
	var v428 int32
	_ = v428
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v503 int32
	_ = v503
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v548 int32
	_ = v548
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v566 int32
	_ = v566
	v4 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(32)
	m.G0 = v22
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	if base.Ui32(int32(253)) < base.Ui32((v24-int32(3))&int32(255)) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l1 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L13
	} else {
		goto L90
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L13
	} else {
		goto L87
	}
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v24<<(uint(int32(2))%32))+uint32(_c_F_GetLockConflicts[0])))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v36 < l1 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[1]))
	if base.Ui32(int32(2)) <= base.Ui32(v40) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[2]))
	v74 = F_get_hash_value(m, v73, l0)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L13
	} else {
		goto L16
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[3])) = v70
	goto L7
L9:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[3]))
	if v44 != 0 {
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[4]))
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[5]))
	v68 = F_palloc0_mul(m, int32(8), v62+v64+int32(1))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L13
	} else {
		goto L15
	}
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[6]))
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[5]))
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[4]))
	v56 = F_MemoryContextAlloc(m, v46, (v48+v50)<<(uint(int32(3))%32)+int32(8))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	v70 = v56
	goto L8
L15:
	;
	v70 = v68
	goto L8
L16:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[7]))
	v84 = v77 + v74&int32(15)<<(uint(int32(7))%32) + int32(_a_F_GetLockConflicts_0)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v85+l1<<(uint(int32(2))%32))))
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	if v90 != int32(1) {
		v287 = v4
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v303 = F_LWLockAcquire(m, v84, int32(1))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L13
	} else {
		goto L47
	}
L18:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
	if v93|base.B2i32(base.Ui32(l1) < base.Ui32(int32(5))) != 0 {
		v287 = v4
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v97 == int32(0) {
		v287 = v4
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[8]))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+16))
	if v102 == int32(0) {
		v287 = v4
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[9]))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v112 = (v106 - int32(1)) & (v109 * int32(_a_F_GetLockConflicts_1))
	v115 = int32(3)
	v122 = v101
	v125 = v4
	v129 = v4
	goto L22
L22:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v143 = v140 + v129*int32(768)
	v145 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[10]))
	if v143 != v145 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v287 = v262
	goto L17
L24:
	;
	v148 = v143 + int32(548)
	v150 = F_LWLockAcquire(m, v148, int32(1))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L13
	} else {
		goto L27
	}
L25:
	;
	v262 = v125
	goto L26
L26:
	;
	v278 = v129 + int32(1)
	v280 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[8]))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v280)+16))
	if base.Ui32(v278) < base.Ui32(v281) {
		v122 = v280
		v125 = v262
		v129 = v278
		goto L22
	} else {
		goto L46
	}
L27:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v143)+20))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v152 != v153 {
		v241 = v125
		goto L28
	} else {
		goto L29
	}
L28:
	;
	F_LWLockRelease(m, v148)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L13
	} else {
		goto L45
	}
L29:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v143)+564))
	v157 = *(*int64)(unsafe.Add(mBase, uint32(v155+v112<<(uint(v115)%32))))
	if v157 == int64(0) {
		v241 = v125
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v160 = v155 + v112&int32(268435455)<<(uint(v115)%32)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v143)+568))
	v162 = v161 + v112<<(uint(int32(6))%32)
	v181 = int64(0)
	goto L31
L31:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v162+base.I32_wrap_i64(v181)<<(uint(int32(2))%32))))
	if v187 == v109 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	if v217<<(uint(int32(1))%32)&v89 == int32(0) {
		v241 = v125
		goto L28
	} else {
		goto L43
	}
L33:
	;
	goto L32
L34:
	;
	v189 = *(*int64)(unsafe.Add(mBase, uint32(v160)))
	v195 = base.I32_wrap_i64(int64(base.Ui64(v189)>>(uint(v181*int64(3))%64))) & int32(7)
	if v195 != 0 {
		v217 = v195
		goto L33
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v198 = v181 | int64(1)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v162+base.I32_wrap_i64(v198)<<(uint(int32(2))%32))))
	if v203 == v109 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L36
L38:
	;
	v205 = *(*int64)(unsafe.Add(mBase, uint32(v160)))
	v211 = base.I32_wrap_i64(int64(base.Ui64(v205)>>(uint(v198*int64(3))%64))) & int32(7)
	if v211 != 0 {
		v217 = v211
		goto L33
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v214 = v181 + int64(2)
	if v214 != int64(16) {
		v181 = v214
		goto L31
	} else {
		goto L42
	}
L41:
	;
	goto L40
L42:
	;
	v241 = v125
	goto L28
L43:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v143)+44))
	if v224 == int32(0) {
		v241 = v125
		goto L28
	} else {
		goto L44
	}
L44:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v143)+40))
	v229 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[3]))
	v232 = v229 + v125<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v232)+4)) = v224
	*(*int32)(unsafe.Add(mBase, uint32(v232))) = v227
	v241 = v125 + int32(1)
	goto L28
L45:
	;
	v262 = v241
	goto L26
L46:
	;
	goto L23
L47:
	;
	v306 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[2]))
	v307 = int32(0)
	v309 = F_hash_search_with_hash_value(m, v306, l0, v74, v307, v307)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L13
	} else {
		goto L51
	}
L48:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L13
	} else {
		goto L84
	}
L49:
	;
	m.G0 = v22 + int32(32)
	return v503
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v486
	v503 = v483
	goto L49
L51:
	;
	if v309 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	F_LWLockRelease(m, v84)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L13
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v309)+28))
	if v322 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	v316 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[3]))
	*(*int64)(unsafe.Add(mBase, uint32(v316+v287<<(uint(int32(3))%32)))) = int64(4294967295)
	if l2 != 0 {
		v483 = v316
		v486 = v287
		goto L50
	} else {
		goto L56
	}
L56:
	;
	v503 = v316
	goto L49
L57:
	;
	F_LWLockRelease(m, v84)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L13
	} else {
		goto L81
	}
L58:
	;
	v449 = v287
	goto L57
L59:
	;
	goto L60
L60:
	;
	v326 = v309 + int32(24)
	if v326 == v322 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v449 = v287
	goto L57
L62:
	;
	goto L63
L63:
	;
	v329 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[3]))
	v332 = v322
	v335 = v287
	goto L64
L64:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v332-int32(8))))
	if v353&v89 == int32(0) {
		v428 = v335
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v449 = v428
	goto L57
L66:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v332)+4))
	if v444 != v326 {
		v332 = v444
		v335 = v428
		goto L64
	} else {
		goto L80
	}
L67:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v332-int32(16))))
	v361 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[10]))
	if v359 == v361 {
		v428 = v335
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v359)+44))
	if v363 == int32(0) {
		v428 = v335
		goto L66
	} else {
		goto L69
	}
L69:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v359)+40))
	v367 = int32(0)
	if base.B2i32(v287 <= int32(0)) == v367 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v371 = v367
	goto L73
L71:
	;
	goto L72
L72:
	;
	v420 = v329 + v335<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v420)+4)) = v363
	*(*int32)(unsafe.Add(mBase, uint32(v420))) = v366
	v428 = v335 + int32(1)
	goto L66
L73:
	;
	v391 = v329 + v371<<(uint(int32(3))%32)
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	if v366 == v392 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	goto L72
L75:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v391)+4))
	if v394 == v363 {
		v428 = v335
		goto L66
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v397 = v371 + int32(1)
	if v397 != v287 {
		v371 = v397
		goto L73
	} else {
		goto L79
	}
L78:
	;
	goto L77
L79:
	;
	goto L74
L80:
	;
	goto L65
L81:
	;
	v468 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[5]))
	v470 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[4]))
	if v468+v470 < v449 {
		goto L48
	} else {
		goto L82
	}
L82:
	;
	v474 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[3]))
	*(*int64)(unsafe.Add(mBase, uint32(v474+v449<<(uint(int32(3))%32)))) = int64(4294967295)
	if l2 == int32(0) {
		v503 = v474
		goto L49
	} else {
		goto L83
	}
L83:
	;
	v483 = v474
	v486 = v449
	goto L50
L84:
	;
	F_errmsg_internal(m, int32(_a_F_GetLockConflicts_2), int32(0))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L13
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_GetLockConflicts_3), int32(3303), int32(_a_F_GetLockConflicts_4))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L13
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = l1
	F_errmsg_internal(m, int32(_a_F_GetLockConflicts_5), v22+int32(16))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L13
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(_a_F_GetLockConflicts_3), int32(3127), int32(_a_F_GetLockConflicts_4))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L13
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v24
	F_errmsg_internal(m, int32(_a_F_GetLockConflicts_6), v22)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L13
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(_a_F_GetLockConflicts_3), int32(3124), int32(_a_F_GetLockConflicts_4))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L13
	} else {
		goto L92
	}
L92:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_GetNSItemByVar(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
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
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	v3 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v7 <= v3 {
		v52 = l0
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v52)+28))
	if v58 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L2:
	;
	v11 = v7 & int32(7)
	if v11 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if base.Ui32(v7) < base.Ui32(int32(8)) {
		v52 = v26
		goto L1
	} else {
		goto L10
	}
L4:
	;
	v26 = l0
	v28 = v7
	goto L3
L5:
	;
	goto L6
L6:
	;
	v14 = l0
	v16 = v7
	v19 = v3
	goto L7
L7:
	;
	v20 = int32(1)
	v21 = v16 - v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v24 = v19 + v20
	if v24 != v11 {
		v14 = v22
		v16 = v21
		v19 = v24
		goto L7
	} else {
		goto L9
	}
L8:
	;
	v26 = v22
	v28 = v21
	goto L3
L9:
	;
	goto L8
L10:
	;
	v34 = v26
	v36 = v28
	goto L11
L11:
	;
	v40 = int32(8)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if v40 < v36 {
		v34 = v49
		v36 = v36 - v40
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v52 = v49
	goto L1
L13:
	;
	goto L12
L14:
	;
	return v80
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L28
	} else {
		goto L29
	}
L16:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	if v61 <= int32(0) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v64 = int32(0)
	if v64 < v61 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v68 = v61
	goto L20
L19:
	;
	v68 = v64
	goto L20
L20:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
	v73 = v64
	goto L21
L21:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v70+v73<<(uint(int32(2))%32))))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
	if v69 == v81 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L15
L23:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v80)+24))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v83 == v84 {
		goto L14
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v87 = v73 + int32(1)
	if v87 != v68 {
		v73 = v87
		goto L21
	} else {
		goto L27
	}
L26:
	;
	goto L25
L27:
	;
	goto L22
L28:
	;
	return int32(0)
L29:
	;
	F_errmsg_internal(m, int32(_a_F_GetNSItemByVar_0), int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_GetNSItemByVar_1), int32(564), int32(_a_F_GetNSItemByVar_2))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_GetParentedForeignKeyRefs(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int64
	_ = v28
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
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
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(112)
	m.G0 = v8
	v10 = F_RelationGetIndexList(m, l0)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v8 + int32(112)
	return v70
L2:
	;
	return int32(0)
L3:
	;
	if v10 == int32(0) {
		v70 = v2
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v17 = F_RelationGetIndexAttrBitmap(m, l0, int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	if v17 == int32(0) {
		v70 = v2
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v23 = F_table_open(m, int32(2606), int32(1))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v28 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
	F_ScanKeyInit(m, v8, int32(13), int32(3), int32(184), v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	F_ScanKeyInit(m, v8+int32(56), int32(4), int32(3), int32(61), int64(102))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v39 = int32(0)
	v43 = F_systable_beginscan(m, v23, v39, int32(1), v39, int32(2), v8)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v49 = v2
	goto L11
L11:
	;
	v50 = F_systable_getnext(m, v43)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L2
	} else {
		goto L13
	}
L12:
	;
	F_systable_endscan(m, v43)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L2
	} else {
		goto L19
	}
L13:
	;
	if v50 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+22)))
	v54 = v52 + v53
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+92))
	if v55 == int32(0) {
		goto L11
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	goto L12
L17:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v59 = F_lappend_oid(m, v49, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	v49 = v59
	goto L11
L19:
	;
	F_relation_close(m, v23, int32(1))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v70 = v49
	goto L1
}
func F_GetRedoRecPtr(m *base.Module) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v23 int64
	_ = v23
	var v27 int64
	_ = v27
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_GetRedoRecPtr[0]))
	v8 = base.AtomicRmwXchg32(m, v5, int32(440), int32(1))
	if v8 != 0 {
		F_s_lock(m, v5+int32(440), int32(_a_F_GetRedoRecPtr_0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, _c_F_GetRedoRecPtr[0]))
			v18 = *(*int64)(unsafe.Add(mBase, uint32(v17)+200))
			v19 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v17)+440)), uint32(v19))
			v23 = *(*int64)(unsafe.Add(mBase, _c_F_GetRedoRecPtr[1]))
			if base.Ui64(v23) < base.Ui64(v18) {
				*(*int64)(unsafe.Add(mBase, _c_F_GetRedoRecPtr[1])) = v18
				v27 = v18
			} else {
				v27 = v23
			}
			return v27
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_GetRedoRecPtr[0]))
		v18 = *(*int64)(unsafe.Add(mBase, uint32(v17)+200))
		v19 = int32(0)
		atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v17)+440)), uint32(v19))
		v23 = *(*int64)(unsafe.Add(mBase, _c_F_GetRedoRecPtr[1]))
		if base.Ui64(v23) < base.Ui64(v18) {
			*(*int64)(unsafe.Add(mBase, _c_F_GetRedoRecPtr[1])) = v18
			v27 = v18
		} else {
			v27 = v23
		}
		return v27
	}
}
func F_GetSafeSnapshotBlockingPids(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	v4 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_GetSafeSnapshotBlockingPids[0]))
	v12 = F_LWLockAcquire(m, v8+int32(3584), int32(1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_GetSafeSnapshotBlockingPids[1]))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if v18 == int32(0) {
		v74 = v4
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_GetSafeSnapshotBlockingPids[0]))
	F_LWLockRelease(m, v76+int32(3584))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L21
	}
L4:
	;
	v22 = v17 + int32(8)
	if v18 == v22 {
		v74 = v4
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v27 = v18
	goto L7
L6:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+44)))
	if v34&int32(64) == int32(0) {
		v74 = v4
		goto L3
	} else {
		goto L11
	}
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+48))
	if v30 == l0 {
		goto L6
	} else {
		goto L9
	}
L8:
	;
	v74 = v4
	goto L3
L9:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v32 != v22 {
		v27 = v32
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v27)+28))
	if v39 == int32(0) {
		v74 = v4
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v43 = v27 + int32(24)
	if v39 == v43 {
		v74 = v4
		goto L3
	} else {
		goto L13
	}
L13:
	;
	v45 = int32(1)
	if l2 <= v45 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v48 = v45
	goto L16
L15:
	;
	v48 = l2
	goto L16
L16:
	;
	v52 = v39
	v55 = int32(0)
	goto L17
L17:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+112))
	*(*int32)(unsafe.Add(mBase, uint32(l1+v55<<(uint(int32(2))%32)))) = v62
	v65 = v55 + int32(1)
	if v48-int32(1) == v55 {
		v74 = v65
		goto L3
	} else {
		goto L19
	}
L18:
	;
	v74 = v65
	goto L3
L19:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v67 != v43 {
		v52 = v67
		v55 = v65
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	return v74
}
func F_GetSingleProcBlockerStatusData(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
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
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int64
	_ = v70
	var v72 int64
	_ = v72
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+384))
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v13 + int32(1)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v20 = v17 + v13*int32(20)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	if v27 == int32(0) {
		v106 = v25
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v116-v106 < v115 {
		goto L18
	} else {
		goto L19
	}
L5:
	;
	v31 = v12 + int32(24)
	if v27 == v31 {
		v106 = v25
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v36 = v33
	v38 = v27
	goto L7
L7:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v38-int32(20))))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v38-int32(16))))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v52 <= v36 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v106 = v103
	goto L4
L9:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_GetSingleProcBlockerStatusData[0]))
	v56 = v55 + v52
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v56
	v60 = F_repalloc(m, v51, v56*int32(56))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v65 = v51
	v66 = v36
	goto L11
L11:
	;
	v69 = v66*int32(56) + v65
	v70 = *(*int64)(unsafe.Add(mBase, uint32(v47)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v69)+8)) = v70
	v72 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
	*(*int64)(unsafe.Add(mBase, uint32(v69))) = v72
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v38-int32(8))))
	*(*int32)(unsafe.Add(mBase, uint32(v69)+16)) = v76
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v50)+384))
	if v47 == v78 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	return
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v60
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v65 = v60
	v66 = v63
	goto L11
L14:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v50)+400))
	v82 = v80
	goto L16
L15:
	;
	v82 = int32(0)
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69)+20)) = v82
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v50)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v69)+24)) = v84
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v50)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v69)+28)) = v86
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v69)+40)) = v88
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v38-int32(12))))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+12))
	v94 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v69)+48)) = uint8(v94)
	*(*int32)(unsafe.Add(mBase, uint32(v69)+44)) = v93
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v99 = v97 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v99
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v101 != v31 {
		v36 = v99
		v38 = v101
		goto L7
	} else {
		goto L17
	}
L17:
	;
	goto L8
L18:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_GetSingleProcBlockerStatusData[0]))
	v121 = v120 + v116
	v122 = v106 + v115
	if v122 < v121 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v135 = v12 + int32(32)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	if v136 != 0 {
		goto L26
	} else {
		goto L27
	}
L21:
	;
	v124 = v121
	goto L23
L22:
	;
	v124 = v122
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v124
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v129 = F_repalloc(m, v126, v124<<(uint(int32(2))%32))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L12
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v129
	goto L20
L25:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v182 - v183
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v186 - v187
	goto L3
L26:
	;
	v137 = v136
	goto L28
L27:
	;
	v137 = v135
	goto L28
L28:
	;
	if base.B2i32(v135 == v137)|base.B2i32(v137-int32(388) == l0) != 0 {
		goto L25
	} else {
		goto L29
	}
L29:
	;
	v145 = v137
	goto L30
L30:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v145-int32(376))))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v157 + int32(1)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v161+v157<<(uint(int32(2))%32)))) = v156
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v145)+4))
	if v166 == v135 {
		goto L25
	} else {
		goto L32
	}
L31:
	;
	goto L25
L32:
	;
	if v166-int32(388) != l0 {
		v145 = v166
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
}
func F_GetTsmRoutine(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int64
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v10 = F_OidFunctionCall1Coll(m, l0, int32(0), int64(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = base.I32_wrap_i64(v10)
		if v14 != 0 {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
			if v15 == int32(446) {
				m.G0 = v6 + int32(16)
				return v14
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
					F_errmsg_internal(m, int32(_a_F_GetTsmRoutine_0), v6)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_GetTsmRoutine_1), int32(37), int32(_a_F_GetTsmRoutine_2))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg_internal(m, int32(_a_F_GetTsmRoutine_0), v6)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_GetTsmRoutine_1), int32(37), int32(_a_F_GetTsmRoutine_2))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
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
func F_GetVictimBuffer(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int64
	_ = v64
	var v67 int64
	_ = v67
	var v73 int32
	_ = v73
	var v82 int64
	_ = v82
	var v89 int64
	_ = v89
	var v90 int32
	_ = v90
	var v92 int64
	_ = v92
	var v94 int64
	_ = v94
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int64
	_ = v108
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
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
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v267 int32
	_ = v267
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int64
	_ = v280
	var v283 int64
	_ = v283
	var v296 int64
	_ = v296
	var v303 int64
	_ = v303
	var v304 int32
	_ = v304
	var v312 int64
	_ = v312
	var v315 int32
	_ = v315
	var v317 int64
	_ = v317
	var v319 int64
	_ = v319
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v335 int64
	_ = v335
	var v339 int64
	_ = v339
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v416 int64
	_ = v416
	var v419 int64
	_ = v419
	var v428 int64
	_ = v428
	var v432 int64
	_ = v432
	var v436 int32
	_ = v436
	var v437 int64
	_ = v437
	var v440 int64
	_ = v440
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v490 int64
	_ = v490
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v567 int64
	_ = v567
	var v569 int64
	_ = v569
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v593 int64
	_ = v593
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v611 int32
	_ = v611
	var v612 int64
	_ = v612
	var v616 int64
	_ = v616
	var v618 int32
	_ = v618
	var v627 int64
	_ = v627
	var v628 int64
	_ = v628
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[0]))
	F_ResourceOwnerEnlarge(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	goto L4
L4:
	;
	v36 = v14 + int32(8)
	v38 = v14 + int32(7)
	v39 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v38))) = uint8(v39)
	if l0 == v39 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	m.G0 = v14 + int32(16)
	return v384
L6:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v381)+20))
	v384 = v382 + int32(1)
	F_CheckBufferIsPinnedOnce(m, v384)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L85
	}
L7:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[1]))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+16))
	if v128 != int32(-1) {
		goto L27
	} else {
		goto L28
	}
L8:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v45 = v43 + int32(1)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v45 < v47 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v49 = v45
	goto L11
L10:
	;
	v49 = int32(0)
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v49
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0+v49<<(uint(int32(2))%32))+12))
	if v54 == int32(0) {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[2]))
	v61 = v58 + v54*int32(56)
	v63 = v61 - int32(32)
	v64 = int64(0)
	v67 = base.AtomicRmwCmpxchg64(m, v63, int32(0), v64, v64)
	if v67&int64(3932159) != v64 {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v73 = v61 - int32(56)
	v82 = v67
	goto L14
L14:
	;
	if v82&int64(4194304) != int64(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L7
L16:
	;
	if v108&int64(3932159) == int64(0) {
		v82 = v108
		goto L14
	} else {
		goto L26
	}
L17:
	;
	v89 = F_WaitBufHdrUnlocked(m, v73)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v92 = v82 | int64(1)
	v94 = base.AtomicRmwCmpxchg64(m, v63, int32(0), v82, v92)
	if v82 != v94 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v108 = v89
	goto L16
L21:
	;
	v108 = v94
	goto L16
L22:
	;
	goto L23
L23:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v36))) = v92
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v61-int32(36))))
	F_TrackNewBufferPin(m, v99+int32(1))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if v73 == int32(0) {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	v106 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v38))) = uint8(v106)
	v381 = v73
	goto L6
L26:
	;
	goto L15
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v127)+16)) = int32(-1)
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[3]))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	v140 = v135 + v128*int32(768) + int32(316)
	v141 = int32(0)
	v144 = base.AtomicRmwOr32(m, v141, int32(_a_F_GetVictimBuffer_0), v141)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	if v145 != 0 {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v194 = v127
	goto L29
L29:
	;
	v197 = base.AtomicRmwAdd32(m, v194, int32(12), int32(1))
	v199 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[4]))
	v204 = v199
	goto L44
L30:
	;
	v193 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[1]))
	v194 = v193
	goto L29
L31:
	;
	goto L30
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v140))) = int32(1)
	v148 = int32(0)
	v151 = base.AtomicRmwOr32(m, v148, int32(_a_F_GetVictimBuffer_0), v148)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v140)+4))
	if v152 == v148 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v140)+12))
	if v155 == int32(0) {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[5]))
	if v159 == v155 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v161 = m.G0
	v163 = v161 - int32(16)
	m.G0 = v163
	v166 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[6]))
	if v166 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	v189 = F_pgmem_kill(m, v155, int32(23))
	mBase = m.M
	goto L31
L38:
	;
	m.G0 = v163 + int32(16)
	goto L30
L39:
	;
	v169 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v163)+15)) = uint8(v169)
	goto L40
L40:
	;
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[7]))
	v177 = F_write(m, v173, v163+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v177 {
		goto L38
	} else {
		goto L42
	}
L41:
	;
	goto L38
L42:
	;
	v181 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[8]))
	if v181 == int32(27) {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v212 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[1]))
	v215 = base.AtomicRmwAdd32(m, v212, int32(4), int32(1))
	v217 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[4]))
	if base.Ui32(v215) < base.Ui32(v217) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L82
	}
L46:
	;
	v276 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[2]))
	v279 = v276 + v267*int32(56)
	v280 = int64(0)
	v283 = base.AtomicRmwCmpxchg64(m, v279, int32(24), v280, v280)
	if v283&int64(262143) == v280 {
		goto L60
	} else {
		goto L61
	}
L47:
	;
	v267 = v215
	goto L46
L48:
	;
	goto L49
L49:
	;
	v219 = base.I32_rem_u_s(v215, v217)
	if v219 != 0 {
		v267 = v219
		goto L46
	} else {
		goto L50
	}
L50:
	;
	v224 = v215 + int32(1)
	goto L51
L51:
	;
	v234 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[1]))
	v237 = base.AtomicRmwXchg32(m, v234, int32(0), int32(1))
	if v237 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v254 = int32(0)
	v256 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[1]))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v256)+8)) = v257 + int32(1)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v256))), uint32(v254))
	v267 = v254
	goto L46
L53:
	;
	F_s_lock(m, v234, int32(_a_F_GetVictimBuffer_1))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v242 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[1]))
	v244 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[4]))
	v245 = base.I32_rem_u_s(v224, v244)
	v247 = base.AtomicRmwCmpxchg32(m, v242, int32(4), v224, v245)
	if v224 != v247 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L55
L57:
	;
	v250 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[1]))
	v251 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v250))), uint32(v251))
	v224 = v247
	goto L51
L58:
	;
	goto L59
L59:
	;
	goto L52
L60:
	;
	v296 = v283
	goto L63
L61:
	;
	goto L62
L62:
	;
	v356 = v204 - int32(1)
	if v356 != 0 {
		v204 = v356
		goto L44
	} else {
		goto L81
	}
L63:
	;
	if v296&int64(4194304) != int64(0) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L62
L65:
	;
	if v339&int64(262143) == int64(0) {
		v296 = v339
		goto L63
	} else {
		goto L80
	}
L66:
	;
	v303 = F_WaitBufHdrUnlocked(m, v279)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	if v296&int64(3932160) != int64(0) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v339 = v303
	goto L65
L70:
	;
	v339 = v335
	goto L65
L71:
	;
	v312 = base.AtomicRmwCmpxchg64(m, v279, int32(24), v296, v296-int64(262144))
	if v296 != v312 {
		v335 = v312
		goto L70
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v317 = v296 | int64(1)
	v319 = base.AtomicRmwCmpxchg64(m, v279, int32(24), v296, v317)
	if v296 != v319 {
		v335 = v319
		goto L70
	} else {
		goto L75
	}
L74:
	;
	v315 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[4]))
	v204 = v315
	goto L44
L75:
	;
	if l0 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v279)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0+v321<<(uint(int32(2))%32))+12)) = v325 + int32(1)
	goto L78
L77:
	;
	goto L78
L78:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v36))) = v317
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v279)+20))
	F_TrackNewBufferPin(m, v330+int32(1))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v381 = v279
	goto L6
L80:
	;
	goto L64
L81:
	;
	goto L45
L82:
	;
	F_errmsg_internal(m, int32(_a_F_GetVictimBuffer_2), int32(0))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_GetVictimBuffer_3), int32(274), int32(_a_F_GetVictimBuffer_4))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L85:
	;
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+10)))
	if v387&int32(128) == int32(0) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v593 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
	if v593&int64(16777216) != int64(0) {
		goto L132
	} else {
		goto L133
	}
L87:
	;
	v393 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[9]))
	if v393 != int32(-1) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v406)+12))
	if v407 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L89:
	;
	v397 = v393 << (uint(int32(4)) % 32)
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v397)+uint32(_c_F_GetVictimBuffer[10])))
	if v400 == v384 {
		v406 = v397 + int32(_a_F_GetVictimBuffer_5)
		goto L88
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v404 = F_GetPrivateRefCountEntrySlow(m, v384, int32(1))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L93
	}
L92:
	;
	goto L91
L93:
	;
	v406 = v404
	goto L88
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v406)+12)) = int32(2)
	if l0 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L95:
	;
	v410 = int32(_a_F_GetVictimBuffer_6)
	v412 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[11])) = v412 + int32(1)
	v416 = int64(0)
	v419 = base.AtomicRmwCmpxchg64(m, v381, int32(24), v416, v416)
	v428 = v419
	goto L98
L96:
	;
	goto L97
L97:
	;
	v460 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[0]))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v381)+20))
	F_ResourceOwnerForget(m, v460, base.I64_extend_i32_s(v461+int32(1)), int32(_a_F_GetVictimBuffer_7))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L1
	} else {
		goto L105
	}
L98:
	;
	v432 = int64(0)
	v436 = base.B2i32(v428&int64(13510798882111488) == v432)
	if v428&int64(13510798882111488) == v432 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	if v428&int64(13510798882111488) == v432 {
		goto L94
	} else {
		goto L104
	}
L100:
	;
	v437 = int64(4503599627370496)
	goto L102
L101:
	;
	v437 = v432
	goto L102
L102:
	;
	v440 = base.AtomicRmwCmpxchg64(m, v381, int32(24), v428, v437|v428)
	if v428 != v440 {
		v428 = v440
		goto L98
	} else {
		goto L103
	}
L103:
	;
	goto L99
L104:
	;
	v442 = int32(_a_F_GetVictimBuffer_6)
	v444 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[11])) = v444 - int32(1)
	goto L97
L105:
	;
	F_UnpinBufferNoOwner(m, v381)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	goto L4
L107:
	;
	F_FlushBuffer(m, v381, int32(0), l1)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L1
	} else {
		goto L119
	}
L108:
	;
	v474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+7)))
	if v474&int32(1) == int32(0) {
		goto L107
	} else {
		goto L109
	}
L109:
	;
	v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+11)))
	if v479&int32(128) == int32(0) {
		goto L107
	} else {
		goto L110
	}
L110:
	;
	v485 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[12]))
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v381)+20))
	v490 = *(*int64)(unsafe.Add(mBase, uint32(v485+v486<<(uint(int32(13))%32))))
	v493 = F_XLogNeedsFlush(m, base.I64_rotl(v490, int64(32)))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	if v493 == int32(0) {
		goto L107
	} else {
		goto L112
	}
L112:
	;
	v497 = int32(0)
	v498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+7)))
	if v498 == v497 {
		v518 = v497
		goto L113
	} else {
		goto L114
	}
L113:
	;
	if v518 == int32(0) {
		goto L107
	} else {
		goto L117
	}
L114:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v501 != int32(1) {
		v518 = v497
		goto L113
	} else {
		goto L115
	}
L115:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v509 = l0 + v504<<(uint(int32(2))%32) + int32(12)
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v509)))
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v381)+20))
	if v510 != v511+int32(1) {
		v518 = v497
		goto L113
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v509))) = int32(0)
	v518 = int32(1)
	goto L113
L117:
	;
	F_UnlockReleaseBuffer(m, v384)
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	goto L4
L119:
	;
	if int32(0) <= v384 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v532 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[2]))
	v533 = int32(56)
	F_BufferLockUnlock(m, v384, v532+v384*v533-v533)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L1
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v541 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetVictimBuffer[13])))
	if v541&int32(1) != 0 {
		goto L86
	} else {
		goto L124
	}
L123:
	;
	goto L122
L124:
	;
	v545 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetVictimBuffer[14])))
	if v545&int32(1) == int32(0) {
		goto L86
	} else {
		goto L125
	}
L125:
	;
	v551 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[15]))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v551)))
	if int32(0) < v552 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v555 = int32(_a_F_GetVictimBuffer_8)
	v557 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[16]))
	*(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[16])) = v557 + int32(1)
	v562 = v557 * int32(20)
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v381)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v562)+uint32(_c_F_GetVictimBuffer[17]))) = v565
	v567 = *(*int64)(unsafe.Add(mBase, uint32(v381)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v562)+uint32(_c_F_GetVictimBuffer[18]))) = v567
	v569 = *(*int64)(unsafe.Add(mBase, uint32(v381)))
	*(*int64)(unsafe.Add(mBase, uint32(v562)+uint32(_c_F_GetVictimBuffer[19]))) = v569
	v572 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[15]))
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v572)))
	v575 = v573
	goto L128
L127:
	;
	v575 = v552
	goto L128
L128:
	;
	v577 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[16]))
	if v577 < v575 {
		goto L86
	} else {
		goto L129
	}
L129:
	;
	F_IssuePendingWritebacks(m, int32(_a_F_GetVictimBuffer_9), l1)
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	goto L86
L131:
	;
	goto L5
L132:
	;
	v598 = int32(0)
	v601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+7)))
	if v601 != 0 {
		goto L135
	} else {
		goto L136
	}
L133:
	;
	v628 = v593
	goto L134
L134:
	;
	if v628&int64(33554432) == int64(0) {
		goto L131
	} else {
		goto L139
	}
L135:
	;
	v602 = int32(3)
	goto L137
L136:
	;
	v602 = v598
	goto L137
L137:
	;
	v611 = int32(0) + l1<<(uint(int32(6))%32) + v602<<(uint(int32(3))%32)
	v612 = *(*int64)(unsafe.Add(mBase, uint32(v611)+uint32(_c_F_GetVictimBuffer[20])))
	*(*int64)(unsafe.Add(mBase, uint32(v611)+uint32(_c_F_GetVictimBuffer[20]))) = v612 + int64(1)
	v616 = *(*int64)(unsafe.Add(mBase, uint32(v611)+uint32(_c_F_GetVictimBuffer[21])))
	*(*int64)(unsafe.Add(mBase, uint32(v611)+uint32(_c_F_GetVictimBuffer[21]))) = v616
	v618 = int32(1)
	F_pgstat_count_backend_io_op(m, v598, l1, v602, v618, int64(0))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, _c_F_GetVictimBuffer[22])) = uint8(v618)
	*(*uint8)(unsafe.Add(mBase, _c_F_GetVictimBuffer[23])) = uint8(v618)
	goto L138
L138:
	;
	v627 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
	v628 = v627
	goto L134
L139:
	;
	v633 = F_InvalidateVictimBuffer(m, v381)
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	if v633 != 0 {
		goto L131
	} else {
		goto L141
	}
L141:
	;
	v636 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[0]))
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v381)+20))
	F_ResourceOwnerForget(m, v636, base.I64_extend_i32_s(v637+int32(1)), int32(_a_F_GetVictimBuffer_7))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	F_UnpinBufferNoOwner(m, v381)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	goto L4
}
func F_GlobalVisCheckRemovableFullXid(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v64 int32
	_ = v64
	var v65 int64
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int64
	_ = v78
	var v80 int32
	_ = v80
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	if l0 == v3 {
		v53 = v3
		v57 = v53
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+117)))
		if v15 != 0 {
			v53 = v3
			v57 = v53
		} else {
			v16 = F_RecoveryInProgress(m)
			mBase = m.M
			if v16 != 0 {
				v53 = v3
				v57 = v53
			} else {
				v17 = F_IsCatalogRelation(m, l0)
				mBase = m.M
				if v17 != 0 {
					v57 = int32(1)
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, _c_F_GlobalVisCheckRemovableFullXid[0]))
					if v20 <= int32(1) {
						v24 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GlobalVisCheckRemovableFullXid[1])))
						if v24&int32(1) == int32(0) {
							v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
							if v49 != 0 {
								v53 = int32(3)
								v57 = v53
							} else {
								v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								if v50 != 0 {
									v53 = int32(3)
									v57 = v53
								} else {
									v57 = int32(2)
								}
							}
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+118)))
							if v30 != int32(112) {
								v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
								if v49 != 0 {
									v53 = int32(3)
									v57 = v53
								} else {
									v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
									if v50 != 0 {
										v53 = int32(3)
										v57 = v53
									} else {
										v57 = int32(2)
									}
								}
							} else {
								if v20 <= int32(0) {
									v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
									if v35 != 0 {
										v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
										if v49 != 0 {
											v53 = int32(3)
											v57 = v53
										} else {
											v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
											if v50 != 0 {
												v53 = int32(3)
												v57 = v53
											} else {
												v57 = int32(2)
											}
										}
									} else {
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
										if v36 != 0 {
											v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
											if v49 != 0 {
												v53 = int32(3)
												v57 = v53
											} else {
												v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
												if v50 != 0 {
													v53 = int32(3)
													v57 = v53
												} else {
													v57 = int32(2)
												}
											}
										} else {
											v37 = int32(1)
											v38 = F_IsCatalogRelation(m, l0)
											mBase = m.M
											if v38 != 0 {
												v53 = v37
												v57 = v53
											} else {
												v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
												if v39 == int32(0) {
													v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
													if v49 != 0 {
														v53 = int32(3)
														v57 = v53
													} else {
														v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
														if v50 != 0 {
															v53 = int32(3)
															v57 = v53
														} else {
															v57 = int32(2)
														}
													}
												} else {
													v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
													v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+119)))
													switch v43 - int32(109) {
													case 0, 5:
														v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+112)))
														if v46 != 0 {
															v53 = v37
															v57 = v53
														} else {
															v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
															if v49 != 0 {
																v53 = int32(3)
																v57 = v53
															} else {
																v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
																if v50 != 0 {
																	v53 = int32(3)
																	v57 = v53
																} else {
																	v57 = int32(2)
																}
															}
														}
													default:
														v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
														if v49 != 0 {
															v53 = int32(3)
															v57 = v53
														} else {
															v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
															if v50 != 0 {
																v53 = int32(3)
																v57 = v53
															} else {
																v57 = int32(2)
															}
														}
													}
												}
											}
										}
									}
								} else {
									v37 = int32(1)
									v38 = F_IsCatalogRelation(m, l0)
									mBase = m.M
									if v38 != 0 {
										v53 = v37
										v57 = v53
									} else {
										v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
										if v39 == int32(0) {
											v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
											if v49 != 0 {
												v53 = int32(3)
												v57 = v53
											} else {
												v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
												if v50 != 0 {
													v53 = int32(3)
													v57 = v53
												} else {
													v57 = int32(2)
												}
											}
										} else {
											v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
											v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+119)))
											switch v43 - int32(109) {
											case 0, 5:
												v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+112)))
												if v46 != 0 {
													v53 = v37
													v57 = v53
												} else {
													v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
													if v49 != 0 {
														v53 = int32(3)
														v57 = v53
													} else {
														v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
														if v50 != 0 {
															v53 = int32(3)
															v57 = v53
														} else {
															v57 = int32(2)
														}
													}
												}
											default:
												v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
												if v49 != 0 {
													v53 = int32(3)
													v57 = v53
												} else {
													v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
													if v50 != 0 {
														v53 = int32(3)
														v57 = v53
													} else {
														v57 = int32(2)
													}
												}
											}
										}
									}
								}
							}
						}
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+118)))
						if v30 != int32(112) {
							v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
							if v49 != 0 {
								v53 = int32(3)
								v57 = v53
							} else {
								v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								if v50 != 0 {
									v53 = int32(3)
									v57 = v53
								} else {
									v57 = int32(2)
								}
							}
						} else {
							if v20 <= int32(0) {
								v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								if v35 != 0 {
									v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
									if v49 != 0 {
										v53 = int32(3)
										v57 = v53
									} else {
										v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
										if v50 != 0 {
											v53 = int32(3)
											v57 = v53
										} else {
											v57 = int32(2)
										}
									}
								} else {
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									if v36 != 0 {
										v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
										if v49 != 0 {
											v53 = int32(3)
											v57 = v53
										} else {
											v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
											if v50 != 0 {
												v53 = int32(3)
												v57 = v53
											} else {
												v57 = int32(2)
											}
										}
									} else {
										v37 = int32(1)
										v38 = F_IsCatalogRelation(m, l0)
										mBase = m.M
										if v38 != 0 {
											v53 = v37
											v57 = v53
										} else {
											v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
											if v39 == int32(0) {
												v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
												if v49 != 0 {
													v53 = int32(3)
													v57 = v53
												} else {
													v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
													if v50 != 0 {
														v53 = int32(3)
														v57 = v53
													} else {
														v57 = int32(2)
													}
												}
											} else {
												v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
												v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+119)))
												switch v43 - int32(109) {
												case 0, 5:
													v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+112)))
													if v46 != 0 {
														v53 = v37
														v57 = v53
													} else {
														v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
														if v49 != 0 {
															v53 = int32(3)
															v57 = v53
														} else {
															v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
															if v50 != 0 {
																v53 = int32(3)
																v57 = v53
															} else {
																v57 = int32(2)
															}
														}
													}
												default:
													v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
													if v49 != 0 {
														v53 = int32(3)
														v57 = v53
													} else {
														v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
														if v50 != 0 {
															v53 = int32(3)
															v57 = v53
														} else {
															v57 = int32(2)
														}
													}
												}
											}
										}
									}
								}
							} else {
								v37 = int32(1)
								v38 = F_IsCatalogRelation(m, l0)
								mBase = m.M
								if v38 != 0 {
									v53 = v37
									v57 = v53
								} else {
									v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
									if v39 == int32(0) {
										v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
										if v49 != 0 {
											v53 = int32(3)
											v57 = v53
										} else {
											v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
											if v50 != 0 {
												v53 = int32(3)
												v57 = v53
											} else {
												v57 = int32(2)
											}
										}
									} else {
										v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
										v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+119)))
										switch v43 - int32(109) {
										case 0, 5:
											v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+112)))
											if v46 != 0 {
												v53 = v37
												v57 = v53
											} else {
												v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
												if v49 != 0 {
													v53 = int32(3)
													v57 = v53
												} else {
													v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
													if v50 != 0 {
														v53 = int32(3)
														v57 = v53
													} else {
														v57 = int32(2)
													}
												}
											}
										default:
											v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
											if v49 != 0 {
												v53 = int32(3)
												v57 = v53
											} else {
												v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
												if v50 != 0 {
													v53 = int32(3)
													v57 = v53
												} else {
													v57 = int32(2)
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
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v57<<(uint(int32(2))%32))+uint32(_c_F_GlobalVisCheckRemovableFullXid[2])))
	v61 = *(*int64)(unsafe.Add(mBase, uint32(v60)+8))
	if base.Ui64(l1) < base.Ui64(v61) {
		v80 = int32(1)
		m.G0 = v8 + int32(48)
		return v80
	} else {
		v64 = int32(0)
		v65 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
		if base.Ui64(v65) <= base.Ui64(l1) {
			v80 = v64
			m.G0 = v8 + int32(48)
			return v80
		} else {
			v68 = *(*int32)(unsafe.Add(mBase, _c_F_GlobalVisCheckRemovableFullXid[3]))
			if v68 != 0 {
				v70 = *(*int32)(unsafe.Add(mBase, _c_F_GlobalVisCheckRemovableFullXid[4]))
				if v70 == v68 {
					v80 = v64
					m.G0 = v8 + int32(48)
					return v80
				} else {
					F_ComputeXidHorizons(m, v8+int32(8))
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return int32(0)
					} else {
						v78 = *(*int64)(unsafe.Add(mBase, uint32(v60)+8))
						v80 = base.B2i32(base.Ui64(l1) < base.Ui64(v78))
						m.G0 = v8 + int32(48)
						return v80
					}
				}
			} else {
				F_ComputeXidHorizons(m, v8+int32(8))
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return int32(0)
				} else {
					v78 = *(*int64)(unsafe.Add(mBase, uint32(v60)+8))
					v80 = base.B2i32(base.Ui64(l1) < base.Ui64(v78))
					m.G0 = v8 + int32(48)
					return v80
				}
			}
		}
	}
}
func F_GlobalVisTestFor(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	v2 = int32(0)
	if l0 == v2 {
		v45 = v2
		v49 = v45
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+117)))
		if v7 != 0 {
			v45 = v2
			v49 = v45
		} else {
			v8 = F_RecoveryInProgress(m)
			mBase = m.M
			if v8 != 0 {
				v45 = v2
				v49 = v45
			} else {
				v9 = F_IsCatalogRelation(m, l0)
				mBase = m.M
				if v9 != 0 {
					v49 = int32(1)
				} else {
					v12 = *(*int32)(unsafe.Add(mBase, _c_F_GlobalVisTestFor[0]))
					if v12 <= int32(1) {
						v16 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GlobalVisTestFor[1])))
						if v16&int32(1) == int32(0) {
							v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
							if v41 != 0 {
								v45 = int32(3)
								v49 = v45
							} else {
								v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								if v42 != 0 {
									v45 = int32(3)
									v49 = v45
								} else {
									v49 = int32(2)
								}
							}
						} else {
							v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+118)))
							if v22 != int32(112) {
								v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
								if v41 != 0 {
									v45 = int32(3)
									v49 = v45
								} else {
									v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
									if v42 != 0 {
										v45 = int32(3)
										v49 = v45
									} else {
										v49 = int32(2)
									}
								}
							} else {
								if v12 <= int32(0) {
									v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
									if v27 != 0 {
										v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
										if v41 != 0 {
											v45 = int32(3)
											v49 = v45
										} else {
											v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
											if v42 != 0 {
												v45 = int32(3)
												v49 = v45
											} else {
												v49 = int32(2)
											}
										}
									} else {
										v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
										if v28 != 0 {
											v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
											if v41 != 0 {
												v45 = int32(3)
												v49 = v45
											} else {
												v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
												if v42 != 0 {
													v45 = int32(3)
													v49 = v45
												} else {
													v49 = int32(2)
												}
											}
										} else {
											v29 = int32(1)
											v30 = F_IsCatalogRelation(m, l0)
											mBase = m.M
											if v30 != 0 {
												v45 = v29
												v49 = v45
											} else {
												v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
												if v31 == int32(0) {
													v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
													if v41 != 0 {
														v45 = int32(3)
														v49 = v45
													} else {
														v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
														if v42 != 0 {
															v45 = int32(3)
															v49 = v45
														} else {
															v49 = int32(2)
														}
													}
												} else {
													v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
													v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+119)))
													switch v35 - int32(109) {
													case 0, 5:
														v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+112)))
														if v38 != 0 {
															v45 = v29
															v49 = v45
														} else {
															v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
															if v41 != 0 {
																v45 = int32(3)
																v49 = v45
															} else {
																v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
																if v42 != 0 {
																	v45 = int32(3)
																	v49 = v45
																} else {
																	v49 = int32(2)
																}
															}
														}
													default:
														v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
														if v41 != 0 {
															v45 = int32(3)
															v49 = v45
														} else {
															v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
															if v42 != 0 {
																v45 = int32(3)
																v49 = v45
															} else {
																v49 = int32(2)
															}
														}
													}
												}
											}
										}
									}
								} else {
									v29 = int32(1)
									v30 = F_IsCatalogRelation(m, l0)
									mBase = m.M
									if v30 != 0 {
										v45 = v29
										v49 = v45
									} else {
										v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
										if v31 == int32(0) {
											v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
											if v41 != 0 {
												v45 = int32(3)
												v49 = v45
											} else {
												v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
												if v42 != 0 {
													v45 = int32(3)
													v49 = v45
												} else {
													v49 = int32(2)
												}
											}
										} else {
											v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
											v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+119)))
											switch v35 - int32(109) {
											case 0, 5:
												v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+112)))
												if v38 != 0 {
													v45 = v29
													v49 = v45
												} else {
													v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
													if v41 != 0 {
														v45 = int32(3)
														v49 = v45
													} else {
														v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
														if v42 != 0 {
															v45 = int32(3)
															v49 = v45
														} else {
															v49 = int32(2)
														}
													}
												}
											default:
												v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
												if v41 != 0 {
													v45 = int32(3)
													v49 = v45
												} else {
													v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
													if v42 != 0 {
														v45 = int32(3)
														v49 = v45
													} else {
														v49 = int32(2)
													}
												}
											}
										}
									}
								}
							}
						}
					} else {
						v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+118)))
						if v22 != int32(112) {
							v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
							if v41 != 0 {
								v45 = int32(3)
								v49 = v45
							} else {
								v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								if v42 != 0 {
									v45 = int32(3)
									v49 = v45
								} else {
									v49 = int32(2)
								}
							}
						} else {
							if v12 <= int32(0) {
								v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								if v27 != 0 {
									v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
									if v41 != 0 {
										v45 = int32(3)
										v49 = v45
									} else {
										v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
										if v42 != 0 {
											v45 = int32(3)
											v49 = v45
										} else {
											v49 = int32(2)
										}
									}
								} else {
									v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									if v28 != 0 {
										v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
										if v41 != 0 {
											v45 = int32(3)
											v49 = v45
										} else {
											v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
											if v42 != 0 {
												v45 = int32(3)
												v49 = v45
											} else {
												v49 = int32(2)
											}
										}
									} else {
										v29 = int32(1)
										v30 = F_IsCatalogRelation(m, l0)
										mBase = m.M
										if v30 != 0 {
											v45 = v29
											v49 = v45
										} else {
											v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
											if v31 == int32(0) {
												v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
												if v41 != 0 {
													v45 = int32(3)
													v49 = v45
												} else {
													v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
													if v42 != 0 {
														v45 = int32(3)
														v49 = v45
													} else {
														v49 = int32(2)
													}
												}
											} else {
												v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
												v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+119)))
												switch v35 - int32(109) {
												case 0, 5:
													v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+112)))
													if v38 != 0 {
														v45 = v29
														v49 = v45
													} else {
														v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
														if v41 != 0 {
															v45 = int32(3)
															v49 = v45
														} else {
															v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
															if v42 != 0 {
																v45 = int32(3)
																v49 = v45
															} else {
																v49 = int32(2)
															}
														}
													}
												default:
													v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
													if v41 != 0 {
														v45 = int32(3)
														v49 = v45
													} else {
														v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
														if v42 != 0 {
															v45 = int32(3)
															v49 = v45
														} else {
															v49 = int32(2)
														}
													}
												}
											}
										}
									}
								}
							} else {
								v29 = int32(1)
								v30 = F_IsCatalogRelation(m, l0)
								mBase = m.M
								if v30 != 0 {
									v45 = v29
									v49 = v45
								} else {
									v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
									if v31 == int32(0) {
										v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
										if v41 != 0 {
											v45 = int32(3)
											v49 = v45
										} else {
											v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
											if v42 != 0 {
												v45 = int32(3)
												v49 = v45
											} else {
												v49 = int32(2)
											}
										}
									} else {
										v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
										v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+119)))
										switch v35 - int32(109) {
										case 0, 5:
											v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+112)))
											if v38 != 0 {
												v45 = v29
												v49 = v45
											} else {
												v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
												if v41 != 0 {
													v45 = int32(3)
													v49 = v45
												} else {
													v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
													if v42 != 0 {
														v45 = int32(3)
														v49 = v45
													} else {
														v49 = int32(2)
													}
												}
											}
										default:
											v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
											if v41 != 0 {
												v45 = int32(3)
												v49 = v45
											} else {
												v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
												if v42 != 0 {
													v45 = int32(3)
													v49 = v45
												} else {
													v49 = int32(2)
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
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v49<<(uint(int32(2))%32))+uint32(_c_F_GlobalVisTestFor[2])))
	return v52
}
func F___gmtime_r(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	F_do_tzset(m)
	mBase = m.M
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v5 = m.Env.X_gmtime_js(m, v4, l1)
	mBase = m.M
	if v5 != 0 {
		v11 = int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = int32(_a_F___gmtime_r_0)
		*(*int64)(unsafe.Add(mBase, uint32(l1)+32)) = int64(0)
		v11 = l1
	}
	return v11
}
func F_g_box_consider_split(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 float64, l5 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 float32
	_ = v31
	var v41 float32
	_ = v41
	var v48 int32
	_ = v48
	var v56 float64
	_ = v56
	var v57 float64
	_ = v57
	var v58 float64
	_ = v58
	var v60 float64
	_ = v60
	var v69 float64
	_ = v69
	var v70 float64
	_ = v70
	var v71 float64
	_ = v71
	var v73 float64
	_ = v73
	var v87 float64
	_ = v87
	var v88 int32
	_ = v88
	var v91 float64
	_ = v91
	var v93 float64
	_ = v93
	var v95 float64
	_ = v95
	var v107 float64
	_ = v107
	var v108 int32
	_ = v108
	var v109 float64
	_ = v109
	var v121 float64
	_ = v121
	var v122 int32
	_ = v122
	var v124 float64
	_ = v124
	var v126 float64
	_ = v126
	var v134 float64
	_ = v134
	var v135 int32
	_ = v135
	var v139 float64
	_ = v139
	var v146 float64
	_ = v146
	var v147 int32
	_ = v147
	var v148 float64
	_ = v148
	var v149 float32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 float32
	_ = v153
	var v156 float32
	_ = v156
	var v161 float32
	_ = v161
	var v162 float32
	_ = v162
	var v163 float32
	_ = v163
	var v166 float32
	_ = v166
	var v171 float64
	_ = v171
	var v179 int32
	_ = v179
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v19 = base.I32_div_s(v15+int32(1), int32(2))
	if l3 < v19 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = base.I32_div_s(v15, int32(2))
	if l5 < v22 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v25 = l3
	goto L3
L3:
	;
	v28 = v15 - v25
	if v25 < v28 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	v24 = l5
	goto L6
L5:
	;
	v24 = v22
	goto L6
L6:
	;
	v25 = v24
	goto L3
L7:
	;
	F_float_underflow_error(m)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L30
	} else {
		goto L66
	}
L8:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L30
	} else {
		goto L65
	}
L9:
	;
	v30 = v25
	goto L11
L10:
	;
	v30 = v28
	goto L11
L11:
	;
	v31 = base.F32_convert_i32_s(v30)
	if base.B2i32(v15 == int32(0))&base.B2i32(base.Ui32(base.I32_reinterpret_f32(v31)&int32(2147483647)) <= base.Ui32(int32(2139095040))) == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v41 = base.F32_div(v31, base.F32_convert_i32_s(v15))
	if base.F32_eq(base.F32_abs(v41), math.Float32frombits(uint32(0x7f800000))) != 0 {
		goto L8
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	F_float_zero_divide_error(m)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L30
	} else {
		goto L64
	}
L15:
	;
	if v30 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v48 = base.F32_eq(v41, float32(0))
	goto L18
L17:
	;
	v48 = int32(0)
	goto L18
L18:
	;
	if v48 != 0 {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	if base.F64_gt(base.F64_promote_f32(v41), float64(0.3)) == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return
L21:
	;
	if l1 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v93 = math.Float64frombits(uint64(0x7ff0000000000000))
	v95 = base.F64_sub(l4, l2)
	if base.F64_eq(base.F64_abs(l2), v93)|base.F64_ne(base.F64_abs(v95), v93)|base.F64_eq(base.F64_abs(l4), v93) == int32(0) {
		goto L32
	} else {
		goto L33
	}
L23:
	;
	v87 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L30
	} else {
		goto L31
	}
L24:
	;
	v56 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
	v57 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	v58 = base.F64_sub(v56, v57)
	v60 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v58), v60)|base.F64_eq(base.F64_abs(v56), v60) != 0 {
		v91 = v58
		goto L22
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v69 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
	v70 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
	v71 = base.F64_sub(v69, v70)
	v73 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v71), v73)|base.F64_eq(base.F64_abs(v69), v73)|base.F64_eq(base.F64_abs(v70), v73) != 0 {
		v91 = v71
		goto L22
	} else {
		goto L29
	}
L27:
	;
	if base.F64_ne(base.F64_abs(v57), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L23
	} else {
		goto L28
	}
L28:
	;
	v91 = v58
	goto L22
L29:
	;
	goto L23
L30:
	;
	return
L31:
	;
	v91 = v87
	goto L22
L32:
	;
	v107 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L30
	} else {
		goto L35
	}
L33:
	;
	v109 = v95
	goto L34
L34:
	;
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v109)&int64(9223372036854775807)))|base.F64_ne(v91, float64(0)) == int32(0) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v109 = v107
	goto L34
L36:
	;
	v149 = base.F32_demote_f64(v148)
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	if v150 != 0 {
		goto L47
	} else {
		goto L48
	}
L37:
	;
	v121 = F_float_zero_divide_error_ext(m, int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L30
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v124 = math.Float64frombits(uint64(0x7ff0000000000000))
	v126 = base.F64_div(v109, v91)
	if base.F64_eq(base.F64_abs(v109), v124)|base.F64_ne(base.F64_abs(v126), v124) == int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v148 = v121
	goto L36
L41:
	;
	v134 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L30
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v139 = float64(0)
	if base.F64_eq(base.F64_abs(v91), math.Float64frombits(uint64(0x7ff0000000000000)))|base.F64_ne(v126, v139)|base.F64_eq(v109, v139) != 0 {
		v148 = v126
		goto L36
	} else {
		goto L45
	}
L44:
	;
	v148 = v134
	goto L36
L45:
	;
	v146 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L30
	} else {
		goto L46
	}
L46:
	;
	v148 = v146
	goto L36
L47:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+80)) = v91
	*(*float32)(unsafe.Add(mBase, uint32(l0)+64)) = v41
	v179 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)) = uint8(v179)
	*(*float32)(unsafe.Add(mBase, uint32(l0)+68)) = v149
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = l1
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = l4
	goto L20
L48:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if l1 == v151 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v153 = *(*float32)(unsafe.Add(mBase, uint32(l0)+68))
	if base.F32_gt(v153, v149) != 0 {
		goto L47
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	if base.F64_ge(v148, float64(-7.006492321624085e-46)) != 0 {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	if base.F32_ne(v149, v153) != 0 {
		goto L20
	} else {
		goto L53
	}
L53:
	;
	v156 = *(*float32)(unsafe.Add(mBase, uint32(l0)+64))
	if base.F32_gt(v41, v156) != 0 {
		goto L47
	} else {
		goto L54
	}
L54:
	;
	goto L20
L55:
	;
	v161 = v149
	goto L57
L56:
	;
	v161 = float32(0)
	goto L57
L57:
	;
	v162 = *(*float32)(unsafe.Add(mBase, uint32(l0)+68))
	v163 = float32(0)
	if base.F32_ge(v162, v163) != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v166 = v162
	goto L60
L59:
	;
	v166 = v163
	goto L60
L60:
	;
	if base.F32_lt(v161, v166) != 0 {
		goto L47
	} else {
		goto L61
	}
L61:
	;
	if base.F32_le(v161, v166) == int32(0) {
		goto L20
	} else {
		goto L62
	}
L62:
	;
	v171 = *(*float64)(unsafe.Add(mBase, uint32(l0)+80))
	if base.F64_gt(v91, v171) == int32(0) {
		goto L20
	} else {
		goto L63
	}
L63:
	;
	goto L47
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_g_intbig_same(m *base.Module, l0 int32) int64 {
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14295(m, l0, int32(2), int32(4), int32(252))
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return v5
	}
}
func F_gather_merge_readnext(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
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
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	if l1 == int32(0) {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+106)))
		if v9 != int32(1) {
			return int32(0)
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
			if v16 != 0 {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
				v19 = v17
			} else {
				v19 = int32(0)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v15)+172)) = v19
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
			if v21 != 0 {
				F_ExecReScan(m, v14)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
					v27 = m.T0[v26].(func(*base.Module, int32) int32)(m, v14)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						v29 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v15)+172)) = v29
						if v27 == v29 {
							v40 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+106)) = uint8(v40)
							return v40
						} else {
							v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+4)))
							if v33&int32(2) != 0 {
								v40 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+106)) = uint8(v40)
								return v40
							} else {
								v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
								*(*int32)(unsafe.Add(mBase, uint32(v36))) = v27
								return int32(1)
							}
						}
					}
				}
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
				v27 = m.T0[v26].(func(*base.Module, int32) int32)(m, v14)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v29 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v15)+172)) = v29
					if v27 == v29 {
						v40 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+106)) = uint8(v40)
						return v40
					} else {
						v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+4)))
						if v33&int32(2) != 0 {
							v40 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+106)) = uint8(v40)
							return v40
						} else {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
							*(*int32)(unsafe.Add(mBase, uint32(v36))) = v27
							return int32(1)
						}
					}
				}
			}
		}
	} else {
		v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
		v47 = v44 + l1<<(uint(int32(4))%32)
		v50 = *(*int32)(unsafe.Add(mBase, uint32(v47-int32(12))))
		v52 = v47 - int32(8)
		v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
		if v53 < v50 {
			v57 = *(*int32)(unsafe.Add(mBase, uint32(v47-int32(16))))
			*(*int32)(unsafe.Add(mBase, uint32(v52))) = v53 + int32(1)
			v64 = *(*int32)(unsafe.Add(mBase, uint32(v57+v53<<(uint(int32(2))%32))))
			v93 = v64
			v96 = int32(1)
			v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
			v101 = *(*int32)(unsafe.Add(mBase, uint32(v97+l1<<(uint(int32(2))%32))))
			v103 = F_ExecStoreMinimalTuple(m, v93, v101, v96)
			mBase = m.M
			v104 = m.ExcPending
			if v104 != 0 {
				return int32(0)
			} else {
				v106 = v96
				return v106
			}
		} else {
			v66 = v47 - int32(4)
			v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
			if v67 != 0 {
				return int32(0)
			} else {
				v70 = int32(0)
				v72 = *(*int32)(unsafe.Add(mBase, _c_F_gather_merge_readnext[0]))
				if v72 != 0 {
					F_ProcessInterrupts(m)
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
						v81 = *(*int32)(unsafe.Add(mBase, uint32(v75+l1<<(uint(int32(2))%32)-int32(4))))
						v82 = F_TupleQueueReaderNext(m, v81, l2, v66)
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return int32(0)
						} else {
							if v82 == int32(0) {
								v106 = v70
								return v106
							} else {
								v87 = F_heap_copy_minimal_tuple(m, v82, int32(0))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return int32(0)
								} else {
									if v87 == int32(0) {
										v106 = v70
										return v106
									} else {
										F_load_tuple_array(m, l0, l1)
										mBase = m.M
										v92 = m.ExcPending
										if v92 != 0 {
											return int32(0)
										} else {
											v93 = v87
											v96 = int32(1)
											v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
											v101 = *(*int32)(unsafe.Add(mBase, uint32(v97+l1<<(uint(int32(2))%32))))
											v103 = F_ExecStoreMinimalTuple(m, v93, v101, v96)
											mBase = m.M
											v104 = m.ExcPending
											if v104 != 0 {
												return int32(0)
											} else {
												v106 = v96
												return v106
											}
										}
									}
								}
							}
						}
					}
				} else {
					v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v75+l1<<(uint(int32(2))%32)-int32(4))))
					v82 = F_TupleQueueReaderNext(m, v81, l2, v66)
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return int32(0)
					} else {
						if v82 == int32(0) {
							v106 = v70
							return v106
						} else {
							v87 = F_heap_copy_minimal_tuple(m, v82, int32(0))
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return int32(0)
							} else {
								if v87 == int32(0) {
									v106 = v70
									return v106
								} else {
									F_load_tuple_array(m, l0, l1)
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int32(0)
									} else {
										v93 = v87
										v96 = int32(1)
										v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
										v101 = *(*int32)(unsafe.Add(mBase, uint32(v97+l1<<(uint(int32(2))%32))))
										v103 = F_ExecStoreMinimalTuple(m, v93, v101, v96)
										mBase = m.M
										v104 = m.ExcPending
										if v104 != 0 {
											return int32(0)
										} else {
											v106 = v96
											return v106
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
func F_gbtreekey_in(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			v19 = F_format_type_extended(m, v7, int32(-1), int32(2))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = v19
				F_errmsg(m, int32(_a_F_gbtreekey_in_0), v5)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_gbtreekey_in_1), int32(34), int32(_a_F_gbtreekey_in_2))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int64(0)
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
func F_gen_partprune_steps_internal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int64
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v272 int32
	_ = v272
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
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
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v435 int32
	_ = v435
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
	var v476 int32
	_ = v476
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v560 int32
	_ = v560
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v608 int32
	_ = v608
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v634 int32
	_ = v634
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
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
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v699 int32
	_ = v699
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
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
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v740 int32
	_ = v740
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v766 int32
	_ = v766
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v910 int32
	_ = v910
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v928 int32
	_ = v928
	var v930 int32
	_ = v930
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v942 int64
	_ = v942
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v964 int32
	_ = v964
	var v982 int32
	_ = v982
	var v984 int32
	_ = v984
	var v991 int32
	_ = v991
	var v995 int32
	_ = v995
	var v1009 int32
	_ = v1009
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1026 int32
	_ = v1026
	var v1044 int32
	_ = v1044
	var v1050 int32
	_ = v1050
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1068 int32
	_ = v1068
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1101 int32
	_ = v1101
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1127 int32
	_ = v1127
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1151 int32
	_ = v1151
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1183 int32
	_ = v1183
	var v1186 int32
	_ = v1186
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1242 int32
	_ = v1242
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1249 int32
	_ = v1249
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1287 int32
	_ = v1287
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1294 int32
	_ = v1294
	var v1297 int32
	_ = v1297
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1306 int32
	_ = v1306
	var v1309 int32
	_ = v1309
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1325 int32
	_ = v1325
	var v1327 int32
	_ = v1327
	var v1330 int32
	_ = v1330
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1336 int32
	_ = v1336
	var v1338 int32
	_ = v1338
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1344 int32
	_ = v1344
	var v1346 int32
	_ = v1346
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1354 int32
	_ = v1354
	var v1356 int32
	_ = v1356
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1370 int32
	_ = v1370
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1378 int32
	_ = v1378
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1385 int32
	_ = v1385
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1393 int32
	_ = v1393
	var v1399 int32
	_ = v1399
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1413 int32
	_ = v1413
	var v1423 int32
	_ = v1423
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1475 int32
	_ = v1475
	var v1495 int32
	_ = v1495
	var v1496 int32
	_ = v1496
	var v1506 int32
	_ = v1506
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1543 int32
	_ = v1543
	var v1544 int32
	_ = v1544
	var v1550 int32
	_ = v1550
	var v1553 int32
	_ = v1553
	var v1555 int32
	_ = v1555
	var v1556 int32
	_ = v1556
	var v1558 int32
	_ = v1558
	var v1561 int32
	_ = v1561
	var v1569 int32
	_ = v1569
	var v1570 int32
	_ = v1570
	var v1575 int32
	_ = v1575
	var v1576 int32
	_ = v1576
	var v1577 int32
	_ = v1577
	var v1579 int32
	_ = v1579
	var v1582 int32
	_ = v1582
	var v1586 int64
	_ = v1586
	var v1591 int32
	_ = v1591
	var v1592 int32
	_ = v1592
	var v1595 int32
	_ = v1595
	var v1598 int32
	_ = v1598
	var v1603 int32
	_ = v1603
	var v1604 int32
	_ = v1604
	var v1605 int64
	_ = v1605
	var v1606 int32
	_ = v1606
	var v1607 int64
	_ = v1607
	var v1608 int32
	_ = v1608
	var v1609 int64
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1611 int64
	_ = v1611
	var v1612 int32
	_ = v1612
	var v1613 int64
	_ = v1613
	var v1617 int64
	_ = v1617
	var v1618 int32
	_ = v1618
	var v1621 int32
	_ = v1621
	var v1622 int64
	_ = v1622
	var v1625 int64
	_ = v1625
	var v1630 int32
	_ = v1630
	var v1631 int32
	_ = v1631
	var v1633 int32
	_ = v1633
	var v1636 int32
	_ = v1636
	var v1637 int32
	_ = v1637
	var v1638 int64
	_ = v1638
	var v1646 int32
	_ = v1646
	var v1657 int32
	_ = v1657
	var v1673 int32
	_ = v1673
	var v1677 int32
	_ = v1677
	var v1680 int32
	_ = v1680
	var v1689 int32
	_ = v1689
	var v1690 int32
	_ = v1690
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1702 int32
	_ = v1702
	var v1703 int32
	_ = v1703
	var v1720 int32
	_ = v1720
	var v1724 int32
	_ = v1724
	var v1726 int32
	_ = v1726
	var v1727 int32
	_ = v1727
	var v1730 int32
	_ = v1730
	var v1731 int32
	_ = v1731
	var v1733 int32
	_ = v1733
	var v1740 int32
	_ = v1740
	var v1741 int32
	_ = v1741
	var v1745 int32
	_ = v1745
	var v1746 int32
	_ = v1746
	var v1750 int32
	_ = v1750
	var v1751 int32
	_ = v1751
	var v1752 int32
	_ = v1752
	var v1753 int32
	_ = v1753
	var v1760 int32
	_ = v1760
	var v1763 int32
	_ = v1763
	var v1764 int32
	_ = v1764
	var v1765 int32
	_ = v1765
	var v1768 int32
	_ = v1768
	var v1777 int32
	_ = v1777
	var v1778 int32
	_ = v1778
	var v1784 int32
	_ = v1784
	var v1785 int32
	_ = v1785
	var v1790 int32
	_ = v1790
	var v1794 int32
	_ = v1794
	var v1799 int32
	_ = v1799
	var v1803 int32
	_ = v1803
	var v1804 int32
	_ = v1804
	var v1810 int32
	_ = v1810
	var v1815 int32
	_ = v1815
	var v1839 int32
	_ = v1839
	var v1840 int32
	_ = v1840
	var v1864 int32
	_ = v1864
	var v1889 int32
	_ = v1889
	var v1890 int32
	_ = v1890
	var v1891 int32
	_ = v1891
	var v1894 int32
	_ = v1894
	var v1897 int32
	_ = v1897
	var v1901 int32
	_ = v1901
	var v1905 int32
	_ = v1905
	var v1906 int32
	_ = v1906
	var v1907 int32
	_ = v1907
	var v1911 int32
	_ = v1911
	var v1912 int32
	_ = v1912
	var v1921 int32
	_ = v1921
	var v1924 int32
	_ = v1924
	var v1936 int32
	_ = v1936
	var v1948 int32
	_ = v1948
	var v1952 int32
	_ = v1952
	var v1956 int32
	_ = v1956
	var v1957 int32
	_ = v1957
	var v1960 int32
	_ = v1960
	var v1961 int32
	_ = v1961
	var v1962 int32
	_ = v1962
	var v1968 int32
	_ = v1968
	var v1969 int32
	_ = v1969
	var v1975 int32
	_ = v1975
	var v1976 int32
	_ = v1976
	var v1978 int32
	_ = v1978
	var v1979 int32
	_ = v1979
	var v1982 int32
	_ = v1982
	var v1986 int32
	_ = v1986
	var v1991 int32
	_ = v1991
	var v1994 int32
	_ = v1994
	var v1995 int32
	_ = v1995
	var v1996 int32
	_ = v1996
	var v2003 int32
	_ = v2003
	var v2004 int32
	_ = v2004
	var v2005 int32
	_ = v2005
	var v2007 int32
	_ = v2007
	var v2008 int32
	_ = v2008
	var v2009 int32
	_ = v2009
	var v2010 int32
	_ = v2010
	var v2011 int32
	_ = v2011
	var v2013 int32
	_ = v2013
	var v2014 int32
	_ = v2014
	var v2015 int32
	_ = v2015
	var v2019 int32
	_ = v2019
	var v2020 int32
	_ = v2020
	var v2023 int32
	_ = v2023
	var v2026 int32
	_ = v2026
	var v2028 int32
	_ = v2028
	var v2040 int32
	_ = v2040
	var v2043 int32
	_ = v2043
	var v2044 int32
	_ = v2044
	var v2046 int32
	_ = v2046
	var v2047 int32
	_ = v2047
	var v2049 int32
	_ = v2049
	var v2050 int32
	_ = v2050
	var v2051 int32
	_ = v2051
	var v2052 int32
	_ = v2052
	var v2054 int32
	_ = v2054
	var v2055 int32
	_ = v2055
	var v2056 int32
	_ = v2056
	var v2058 int32
	_ = v2058
	var v2059 int32
	_ = v2059
	var v2065 int32
	_ = v2065
	var v2066 int32
	_ = v2066
	var v2083 int32
	_ = v2083
	var v2086 int32
	_ = v2086
	var v2087 int32
	_ = v2087
	var v2088 int32
	_ = v2088
	var v2090 int32
	_ = v2090
	var v2091 int32
	_ = v2091
	var v2093 int32
	_ = v2093
	var v2094 int32
	_ = v2094
	var v2101 int32
	_ = v2101
	var v2105 int32
	_ = v2105
	var v2124 int32
	_ = v2124
	var v2127 int32
	_ = v2127
	var v2128 int32
	_ = v2128
	var v2146 int32
	_ = v2146
	var v2147 int32
	_ = v2147
	var v2149 int32
	_ = v2149
	var v2150 int32
	_ = v2150
	var v2153 int32
	_ = v2153
	var v2154 int32
	_ = v2154
	var v2155 int32
	_ = v2155
	var v2156 int32
	_ = v2156
	var v2158 int32
	_ = v2158
	var v2159 int32
	_ = v2159
	var v2160 int32
	_ = v2160
	var v2162 int32
	_ = v2162
	var v2163 int32
	_ = v2163
	var v2170 int32
	_ = v2170
	var v2171 int32
	_ = v2171
	var v2188 int32
	_ = v2188
	var v2191 int32
	_ = v2191
	var v2192 int32
	_ = v2192
	var v2193 int32
	_ = v2193
	var v2195 int32
	_ = v2195
	var v2196 int32
	_ = v2196
	var v2198 int32
	_ = v2198
	var v2199 int32
	_ = v2199
	var v2203 int32
	_ = v2203
	var v2207 int32
	_ = v2207
	var v2211 int32
	_ = v2211
	var v2228 int32
	_ = v2228
	var v2229 int32
	_ = v2229
	var v2231 int32
	_ = v2231
	var v2232 int32
	_ = v2232
	var v2234 int32
	_ = v2234
	var v2235 int32
	_ = v2235
	var v2236 int32
	_ = v2236
	var v2237 int32
	_ = v2237
	var v2239 int32
	_ = v2239
	var v2240 int32
	_ = v2240
	var v2242 int32
	_ = v2242
	var v2243 int32
	_ = v2243
	var v2254 int32
	_ = v2254
	var v2258 int32
	_ = v2258
	var v2271 int32
	_ = v2271
	var v2274 int32
	_ = v2274
	var v2275 int32
	_ = v2275
	var v2276 int32
	_ = v2276
	var v2278 int32
	_ = v2278
	var v2279 int32
	_ = v2279
	var v2281 int32
	_ = v2281
	var v2282 int32
	_ = v2282
	var v2286 int32
	_ = v2286
	var v2295 int32
	_ = v2295
	var v2300 int32
	_ = v2300
	var v2313 int32
	_ = v2313
	var v2314 int32
	_ = v2314
	var v2321 int32
	_ = v2321
	var v2338 int32
	_ = v2338
	var v2339 int32
	_ = v2339
	var v2340 int32
	_ = v2340
	var v2342 int32
	_ = v2342
	var v2343 int32
	_ = v2343
	var v2366 int32
	_ = v2366
	var v2367 int32
	_ = v2367
	var v2368 int32
	_ = v2368
	var v2370 int32
	_ = v2370
	var v2371 int32
	_ = v2371
	var v2391 int32
	_ = v2391
	var v2396 int32
	_ = v2396
	var v2399 int32
	_ = v2399
	var v2402 int32
	_ = v2402
	var v2403 int32
	_ = v2403
	var v2406 int32
	_ = v2406
	var v2412 int32
	_ = v2412
	var v2413 int32
	_ = v2413
	var v2417 int32
	_ = v2417
	var v2423 int32
	_ = v2423
	var v2437 int32
	_ = v2437
	var v2441 int32
	_ = v2441
	var v2442 int32
	_ = v2442
	var v2447 int32
	_ = v2447
	var v2463 int32
	_ = v2463
	var v2469 int32
	_ = v2469
	var v2473 int32
	_ = v2473
	var v2474 int32
	_ = v2474
	var v2475 int32
	_ = v2475
	var v2476 int32
	_ = v2476
	var v2477 int32
	_ = v2477
	var v2478 int32
	_ = v2478
	var v2479 int32
	_ = v2479
	var v2481 int32
	_ = v2481
	var v2482 int32
	_ = v2482
	var v2484 int32
	_ = v2484
	var v2485 int32
	_ = v2485
	var v2487 int32
	_ = v2487
	var v2488 int32
	_ = v2488
	var v2513 int32
	_ = v2513
	var v2516 int32
	_ = v2516
	var v2524 int32
	_ = v2524
	var v2531 int32
	_ = v2531
	var v2535 int32
	_ = v2535
	var v2536 int32
	_ = v2536
	var v2540 int32
	_ = v2540
	var v2541 int32
	_ = v2541
	var v2547 int32
	_ = v2547
	var v2552 int32
	_ = v2552
	var v2554 int32
	_ = v2554
	var v2555 int32
	_ = v2555
	var v2558 int32
	_ = v2558
	var v2565 int32
	_ = v2565
	var v2568 int32
	_ = v2568
	var v2569 int32
	_ = v2569
	var v2570 int32
	_ = v2570
	var v2572 int32
	_ = v2572
	var v2573 int32
	_ = v2573
	var v2574 int32
	_ = v2574
	var v2577 int32
	_ = v2577
	var v2585 int32
	_ = v2585
	var v2592 int32
	_ = v2592
	var v2593 int32
	_ = v2593
	var v2597 int64
	_ = v2597
	var v2602 int32
	_ = v2602
	var v2603 int32
	_ = v2603
	var v2606 int32
	_ = v2606
	var v2609 int32
	_ = v2609
	var v2614 int32
	_ = v2614
	var v2615 int32
	_ = v2615
	var v2616 int64
	_ = v2616
	var v2617 int32
	_ = v2617
	var v2618 int64
	_ = v2618
	var v2619 int32
	_ = v2619
	var v2620 int64
	_ = v2620
	var v2621 int32
	_ = v2621
	var v2622 int64
	_ = v2622
	var v2623 int32
	_ = v2623
	var v2624 int64
	_ = v2624
	var v2628 int64
	_ = v2628
	var v2629 int32
	_ = v2629
	var v2632 int32
	_ = v2632
	var v2633 int64
	_ = v2633
	var v2636 int64
	_ = v2636
	var v2641 int32
	_ = v2641
	var v2642 int32
	_ = v2642
	var v2645 int32
	_ = v2645
	var v2646 int32
	_ = v2646
	var v2649 int32
	_ = v2649
	var v2653 int32
	_ = v2653
	var v2660 int32
	_ = v2660
	var v2661 int32
	_ = v2661
	var v2662 int32
	_ = v2662
	var v2664 int32
	_ = v2664
	var v2665 int32
	_ = v2665
	var v2666 int32
	_ = v2666
	var v2669 int32
	_ = v2669
	var v2677 int32
	_ = v2677
	var v2691 int32
	_ = v2691
	var v2694 int32
	_ = v2694
	var v2698 int32
	_ = v2698
	var v2701 int32
	_ = v2701
	var v2718 int32
	_ = v2718
	var v2722 int32
	_ = v2722
	var v2723 int32
	_ = v2723
	var v2724 int32
	_ = v2724
	var v2725 int32
	_ = v2725
	var v2727 int32
	_ = v2727
	var v2728 int32
	_ = v2728
	var v2731 int32
	_ = v2731
	var v2732 int32
	_ = v2732
	var v2735 int32
	_ = v2735
	var v2736 int32
	_ = v2736
	var v2743 int32
	_ = v2743
	var v2744 int32
	_ = v2744
	var v2745 int32
	_ = v2745
	var v2747 int32
	_ = v2747
	var v2748 int32
	_ = v2748
	var v2752 int32
	_ = v2752
	var v2760 int32
	_ = v2760
	v3 = int32(0)
	v23 = m.G0
	v25 = v23 - int32(288)
	m.G0 = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+256))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+264))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	if v30 == int32(-1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v2752 + int32(288)
	return v2760
L2:
	;
	v45 = int32(0)
	base.MemoryFill(m, v25+int32(80), v45, int32(128))
	if l1 == v45 {
		v2574 = l0
		v2577 = v25
		v2585 = v3
		v2592 = v3
		v2593 = v28
		goto L8
	} else {
		goto L9
	}
L3:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v27)+272))
	v35 = F_predicate_refuted_by(m, v33, l1, int32(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	if v35 == int32(0) {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v41 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)) = uint8(v41)
	v2752 = v25
	v2760 = v3
	goto L1
L7:
	;
	if v2677 == int32(0) {
		goto L531
	} else {
		goto L532
	}
L8:
	;
	v2597 = int64(0)
	if v2592 == int32(0) {
		goto L513
	} else {
		goto L514
	}
L9:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v50 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v53 = l0
	v54 = l1
	v56 = v25
	v64 = v3
	v65 = v3
	v70 = v3
	v71 = v3
	v72 = v28
	v73 = v3
	v74 = v3
	goto L13
L11:
	;
	v1558 = l0
	v1561 = v25
	v1569 = v3
	v1570 = v3
	v1575 = v3
	v1576 = v3
	v1577 = v28
	v1579 = v3
	goto L12
L12:
	;
	if v1570 == int32(0) {
		v1633 = v1575
		goto L343
	} else {
		goto L344
	}
L13:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v54)+12))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v75+v73<<(uint(int32(2))%32))))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	if v80 == int32(320) {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	v1558 = v53
	v1561 = v56
	v1569 = v1543
	v1570 = v1544
	v1575 = v70
	v1576 = v1550
	v1577 = v72
	v1579 = v1553
	goto L12
L15:
	;
	v1555 = v73 + int32(1)
	v1556 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v1555 < v1556 {
		v64 = v1543
		v65 = v1544
		v71 = v1550
		v73 = v1555
		v74 = v1553
		goto L13
	} else {
		goto L341
	}
L16:
	;
	v260 = int32(*(*int16)(unsafe.Add(mBase, uint32(v72)+2)))
	if v260 <= int32(0) {
		v1543 = v64
		v1544 = v65
		v1550 = v71
		v1553 = v74
		goto L15
	} else {
		goto L63
	}
L17:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	switch v100 {
	case 0:
		goto L31
	case 1:
		goto L32
	default:
		goto L16
	}
L18:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v85 = v83
	v86 = v84
	goto L20
L19:
	;
	v85 = v79
	v86 = v80
	goto L20
L20:
	;
	v88 = v86 - int32(7)
	if v88 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	if v88 == int32(14) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+32)))
	if v91 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	goto L17
L25:
	;
	goto L16
L27:
	;
	v94 = *(*int64)(unsafe.Add(mBase, uint32(v85)+24))
	if v94 != int64(0) {
		goto L16
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v97 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+15)) = uint8(v97)
	v2752 = v56
	v2760 = int32(0)
	goto L1
L30:
	;
	goto L29
L31:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
	v244 = F_gen_partprune_steps_internal(m, v53, v243)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L4
	} else {
		goto L57
	}
L32:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
	if v101 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	if v186 == int32(0) {
		v1543 = v64
		v1544 = v65
		v1550 = v71
		v1553 = v74
		goto L15
	} else {
		goto L53
	}
L34:
	;
	v219 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+15)) = uint8(v219)
	v2752 = v56
	v2760 = int32(0)
	goto L1
L35:
	;
	v105 = int32(0)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	if v107 <= v105 {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v112 = v105
	v114 = v105
	v117 = int32(1)
	goto L37
L37:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v101)+12))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v132+v112<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+72)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v56)+76)) = v136
	v142 = F_list_make1_impl(m, int32(1), v56+int32(72))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L4
	} else {
		goto L39
	}
L38:
	;
	if v188&int32(1) == int32(0) {
		goto L33
	} else {
		goto L52
	}
L39:
	;
	v144 = F_gen_partprune_steps_internal(m, v53, v142)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+15)))
	v147 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+15)) = uint8(v147)
	if v146 == v147 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	if v144 != 0 {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	v186 = v114
	v188 = v117
	goto L43
L43:
	;
	v190 = v112 + int32(1)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	if v190 < v191 {
		v112 = v190
		v114 = v186
		v117 = v188
		goto L37
	} else {
		goto L51
	}
L44:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	v184 = F_lappend_int(m, v114, v183)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L4
	} else {
		goto L50
	}
L45:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v144)+12))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v144)+4))
	v156 = int32(4)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v151+v152<<(uint(int32(2))%32)-v156)))
	v181 = v158 + v156
	goto L44
L46:
	;
	goto L47
L47:
	;
	v162 = F_palloc0(m, int32(16))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v162))) = int32(382)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+16)) = v166 + int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v162)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v162)+4)) = v166
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	v174 = F_lappend(m, v173, v162)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+8)) = v174
	v181 = v162 + int32(4)
	goto L44
L50:
	;
	v186 = v184
	v188 = int32(0)
	goto L43
L51:
	;
	goto L38
L52:
	;
	goto L34
L53:
	;
	v225 = F_palloc0(m, int32(16))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v225))) = int32(382)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+16)) = v229 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v225)+12)) = v186
	*(*int32)(unsafe.Add(mBase, uint32(v225)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v225)+4)) = v229
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	v238 = F_lappend(m, v237, v225)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+8)) = v238
	v241 = F_lappend(m, v64, v225)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	v1543 = v241
	v1544 = v65
	v1550 = v71
	v1553 = v74
	goto L15
L57:
	;
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+15)))
	if v246 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v2752 = v56
	v2760 = int32(0)
	goto L1
L59:
	;
	goto L60
L60:
	;
	if v244 == int32(0) {
		v1543 = v64
		v1544 = v65
		v1550 = v71
		v1553 = v74
		goto L15
	} else {
		goto L61
	}
L61:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v244)+12))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v250+v251<<(uint(int32(2))%32)-int32(4))))
	v258 = F_lappend(m, v64, v257)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	v1543 = v258
	v1544 = v65
	v1550 = v71
	v1553 = v74
	goto L15
L63:
	;
	v272 = int32(0)
	goto L65
L64:
	;
	v1519 = F_bms_is_member(m, v272, v65)
	mBase = m.M
	v1520 = m.ExcPending
	if v1520 != 0 {
		goto L4
	} else {
		goto L336
	}
L65:
	;
	v287 = v272 << (uint(int32(2)) % 32)
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)+288))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v287+v289)))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v291)+12))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v292)))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v288)+256))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v294)+12))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v295+v287)))
	v298 = int32(5)
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v294)+4))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v299+v287)))
	if base.B2i32(v301 != int32(2222))&base.B2i32(v301 != int32(424)) != 0 {
		v476 = v298
		goto L84
	} else {
		goto L85
	}
L66:
	;
	v1495 = F_list_concat(m, v64, v1475)
	mBase = m.M
	v1496 = m.ExcPending
	if v1496 != 0 {
		goto L4
	} else {
		goto L335
	}
L67:
	;
	goto L66
L68:
	;
	v1470 = v272 + int32(1)
	v1471 = int32(*(*int16)(unsafe.Add(mBase, uint32(v72)+2)))
	if v1470 < v1471 {
		v272 = v1470
		goto L65
	} else {
		goto L334
	}
L69:
	;
	if v1431 == int32(0) {
		goto L68
	} else {
		goto L333
	}
L70:
	;
	v1423 = int32(0)
	v1430 = v488
	v1431 = int32(5)
	goto L69
L71:
	;
	v1413 = int32(0)
	v1423 = v1413
	v1430 = v488
	v1431 = v1413
	goto L69
L72:
	;
	v1423 = int32(0)
	v1430 = v1405
	v1431 = v1406
	goto L69
L73:
	;
	if v297 != 0 {
		goto L274
	} else {
		goto L275
	}
L74:
	;
	v1242 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+15)) = uint8(v1242)
	v2752 = v56
	v2760 = int32(0)
	goto L1
L75:
	;
	v1213 = F_bms_is_member(m, v272, v65)
	mBase = m.M
	v1214 = m.ExcPending
	if v1214 != 0 {
		goto L4
	} else {
		goto L269
	}
L76:
	;
	v1176 = F_bms_is_member(m, v272, v71)
	mBase = m.M
	v1177 = m.ExcPending
	if v1177 != 0 {
		goto L4
	} else {
		goto L263
	}
L77:
	;
	if v490 != int32(52) {
		v1423 = v488
		v1430 = v488
		v1431 = v476
		goto L69
	} else {
		goto L251
	}
L78:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v85)+24))
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v85)+28))
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v729)+12))
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v730)+4))
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v730)))
	v733 = F_strip_noop_phvs(m, v732)
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L4
	} else {
		goto L160
	}
L79:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v85)+28))
	if v593 == int32(0) {
		v1423 = v488
		v1430 = v488
		v1431 = v476
		goto L69
	} else {
		goto L133
	}
L80:
	;
	v574 = F_makeBoolConst(m, v560, int32(0))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L4
	} else {
		goto L131
	}
L81:
	;
	v560 = int32(0)
	goto L80
L82:
	;
	v496 = F_makeBoolConst(m, v494, int32(0))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L4
	} else {
		goto L116
	}
L83:
	;
	v494 = int32(0)
	goto L82
L84:
	;
	v488 = int32(0)
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	switch v490 - int32(17) {
	case 0:
		goto L79
	case 1, 2:
		v1423 = v488
		v1430 = v488
		v1431 = v476
		goto L69
	case 3:
		goto L78
	default:
		goto L77
	}
L85:
	;
	v307 = int32(0)
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	if v308 != int32(21) {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	v476 = int32(0)
	goto L84
L87:
	;
	v382 = F_strip_noop_phvs(m, v380)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L4
	} else {
		goto L102
	}
L88:
	;
	if v308 != int32(53) {
		v380 = v85
		v381 = v307
		goto L87
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	if v373 != int32(2) {
		v380 = v85
		v381 = v307
		goto L87
	} else {
		goto L101
	}
L91:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	v314 = F_strip_noop_phvs(m, v313)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L4
	} else {
		goto L92
	}
L92:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v314)))
	if v316 == int32(27) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v321 = v314
	goto L96
L94:
	;
	v347 = v314
	goto L95
L95:
	;
	v367 = F_equal(m, v347, v293)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L4
	} else {
		goto L99
	}
L96:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v321)+4))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v341)))
	if v342 == int32(27) {
		v321 = v341
		goto L96
	} else {
		goto L98
	}
L97:
	;
	v347 = v341
	goto L95
L98:
	;
	goto L97
L99:
	;
	if v367 == int32(0) {
		goto L86
	} else {
		goto L100
	}
L100:
	;
	v371 = int32(1)
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
	switch v372 {
	case 0:
		v560 = v371
		goto L80
	case 1:
		v494 = v371
		goto L82
	case 2:
		goto L81
	case 3:
		goto L83
	case 4:
		goto L76
	case 5:
		goto L75
	default:
		v476 = v298
		goto L84
	}
L101:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v376)+12))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)))
	v380 = v378
	v381 = int32(1)
	goto L87
L102:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v382)))
	if v384 == int32(27) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v389 = v382
	goto L106
L104:
	;
	v415 = v382
	goto L105
L105:
	;
	v435 = F_equal(m, v415, v293)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L4
	} else {
		goto L109
	}
L106:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v389)+4))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v409)))
	if v410 == int32(27) {
		v389 = v409
		goto L106
	} else {
		goto L108
	}
L107:
	;
	v415 = v409
	goto L105
L108:
	;
	goto L107
L109:
	;
	if v435 != 0 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v560 = v381 ^ int32(1)
	goto L80
L111:
	;
	goto L112
L112:
	;
	v439 = F_negate_clause(m, v415)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L4
	} else {
		goto L113
	}
L113:
	;
	v441 = F_equal(m, v439, v293)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L4
	} else {
		goto L114
	}
L114:
	;
	if v441 != 0 {
		v560 = v381
		goto L80
	} else {
		goto L115
	}
L115:
	;
	goto L86
L116:
	;
	v499 = F_copyObjectImpl(m, v85)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L4
	} else {
		goto L120
	}
L117:
	;
	v509 = F_palloc0(m, int32(20))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L4
	} else {
		goto L121
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v499)+8)) = v505
	goto L117
L119:
	;
	v505 = int32(0)
	goto L118
L120:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v499)+8))
	switch v501 - int32(1) {
	case 0:
		v505 = int32(2)
		goto L118
	default:
		goto L117
	case 2:
		goto L119
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v509))) = int32(52)
	v513 = F_copyObjectImpl(m, v293)
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L4
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v509)+16)) = int32(-1)
	v517 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v509)+12)) = uint8(v517)
	*(*int32)(unsafe.Add(mBase, uint32(v509)+8)) = v517
	*(*int32)(unsafe.Add(mBase, uint32(v509)+4)) = v513
	*(*int32)(unsafe.Add(mBase, uint32(v56)+284)) = v509
	*(*int32)(unsafe.Add(mBase, uint32(v56)+232)) = v499
	*(*int32)(unsafe.Add(mBase, uint32(v56)+68)) = v499
	*(*int32)(unsafe.Add(mBase, uint32(v56)+64)) = v509
	v532 = F_list_make2_impl(m, v56+int32(68), v56-int32(-64))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L4
	} else {
		goto L123
	}
L123:
	;
	v535 = F_makeBoolExpr(m, int32(1), v532, int32(-1))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L4
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+60)) = v535
	*(*int32)(unsafe.Add(mBase, uint32(v56)+280)) = v535
	v542 = F_list_make1_impl(m, int32(1), v56+int32(60))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L4
	} else {
		goto L125
	}
L125:
	;
	v544 = F_gen_partprune_steps_internal(m, v53, v542)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L4
	} else {
		goto L126
	}
L126:
	;
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+15)))
	if v546 != 0 {
		goto L74
	} else {
		goto L127
	}
L127:
	;
	if v544 != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v549 = int32(3)
	goto L130
L129:
	;
	v549 = int32(5)
	goto L130
L130:
	;
	v1423 = v544
	v1430 = v517
	v1431 = v549
	goto L69
L131:
	;
	v577 = F_palloc(m, int32(24))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L4
	} else {
		goto L132
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v577)+12)) = v574
	v580 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v577)+8)) = uint8(v580)
	*(*int32)(unsafe.Add(mBase, uint32(v577)+4)) = int32(91)
	*(*int32)(unsafe.Add(mBase, uint32(v577))) = v272
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v294)+24))
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v585+v272*int32(28))+4))
	*(*int32)(unsafe.Add(mBase, uint32(v577)+20)) = v580
	*(*int32)(unsafe.Add(mBase, uint32(v577)+16)) = v589
	v1506 = v577
	goto L64
L133:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v593)+4))
	if v596 != int32(2) {
		v1423 = v488
		v1430 = v488
		v1431 = v476
		goto L69
	} else {
		goto L134
	}
L134:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v593)+12))
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v599)))
	v601 = F_strip_noop_phvs(m, v600)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L4
	} else {
		goto L135
	}
L135:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v601)))
	if v603 == int32(27) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v608 = v601
	goto L139
L137:
	;
	v634 = v601
	goto L138
L138:
	;
	v654 = int32(0)
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v85)+28))
	if v655 == v654 {
		v663 = v654
		goto L142
	} else {
		goto L143
	}
L139:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v608)+4))
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v628)))
	if v629 == int32(27) {
		v608 = v628
		goto L139
	} else {
		goto L141
	}
L140:
	;
	v634 = v628
	goto L138
L141:
	;
	goto L140
L142:
	;
	v664 = F_strip_noop_phvs(m, v663)
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L4
	} else {
		goto L145
	}
L143:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v655)+4))
	if v658 < int32(2) {
		v663 = v654
		goto L142
	} else {
		goto L144
	}
L144:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v655)+12))
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v661)+4))
	v663 = v662
	goto L142
L145:
	;
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v664)))
	if v666 == int32(27) {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v673 = v664
	goto L149
L147:
	;
	v699 = v664
	goto L148
L148:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	v718 = F_equal(m, v634, v293)
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L4
	} else {
		goto L152
	}
L149:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v673)+4))
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v691)))
	if v692 == int32(27) {
		v673 = v691
		goto L149
	} else {
		goto L151
	}
L150:
	;
	v699 = v691
	goto L148
L151:
	;
	goto L150
L152:
	;
	if v718 != 0 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v1245 = v699
	v1246 = v717
	goto L73
L154:
	;
	goto L155
L155:
	;
	v721 = F_equal(m, v699, v293)
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L4
	} else {
		goto L156
	}
L156:
	;
	if v721 == int32(0) {
		v1405 = v488
		v1406 = int32(0)
		goto L72
	} else {
		goto L157
	}
L157:
	;
	v725 = F_get_commutator(m, v717)
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L4
	} else {
		goto L158
	}
L158:
	;
	if v725 != 0 {
		v1245 = v634
		v1246 = v725
		goto L73
	} else {
		goto L159
	}
L159:
	;
	goto L70
L160:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v733)))
	if v735 == int32(27) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v740 = v733
	goto L164
L162:
	;
	v766 = v733
	goto L163
L163:
	;
	v786 = F_equal(m, v766, v293)
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L4
	} else {
		goto L167
	}
L164:
	;
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v740)+4))
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v760)))
	if v761 == int32(27) {
		v740 = v760
		goto L164
	} else {
		goto L166
	}
L165:
	;
	v766 = v760
	goto L163
L166:
	;
	goto L165
L167:
	;
	if v786 == int32(0) {
		goto L68
	} else {
		goto L168
	}
L168:
	;
	if v297 != 0 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v85)+24))
	if v297 != v790 {
		goto L68
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	v792 = F_op_in_opfamily(m, v728, v301)
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L4
	} else {
		goto L173
	}
L172:
	;
	goto L171
L173:
	;
	if v792 == int32(0) {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v796 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294))))
	if v796 != int32(108) {
		goto L68
	} else {
		goto L177
	}
L175:
	;
	goto L176
L176:
	;
	v820 = F_op_strict(m, v728)
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L4
	} else {
		goto L184
	}
L177:
	;
	v799 = F_get_negator(m, v728)
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L4
	} else {
		goto L178
	}
L178:
	;
	if v799 == int32(0) {
		goto L68
	} else {
		goto L179
	}
L179:
	;
	v803 = F_op_in_opfamily(m, v799, v301)
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L4
	} else {
		goto L180
	}
L180:
	;
	if v803 == int32(0) {
		goto L68
	} else {
		goto L181
	}
L181:
	;
	F_get_op_opfamily_properties(m, v799, v301, int32(0), v56+int32(240), v56+int32(276), v56+int32(228))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L4
	} else {
		goto L182
	}
L182:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v56)+240))
	if v816 != int32(3) {
		goto L68
	} else {
		goto L183
	}
L183:
	;
	goto L176
L184:
	;
	if v820 == int32(0) {
		v1543 = v64
		v1544 = v65
		v1550 = v71
		v1553 = v74
		goto L15
	} else {
		goto L185
	}
L185:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v731)))
	if v824 == int32(7) {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v862 = F_op_volatile(m, v728)
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L4
	} else {
		goto L203
	}
L187:
	;
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v827 == int32(0) {
		v1543 = v64
		v1544 = v65
		v1550 = v71
		v1553 = v74
		goto L15
	} else {
		goto L188
	}
L188:
	;
	v830 = F_contain_var_clause(m, v731)
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L4
	} else {
		goto L189
	}
L189:
	;
	if v830 != 0 {
		v1543 = v64
		v1544 = v65
		v1550 = v71
		v1553 = v74
		goto L15
	} else {
		goto L190
	}
L190:
	;
	v832 = F_contain_volatile_functions(m, v731)
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L4
	} else {
		goto L191
	}
L191:
	;
	if v832 != 0 {
		v1543 = v64
		v1544 = v65
		v1550 = v71
		v1553 = v74
		goto L15
	} else {
		goto L192
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+240)) = int32(0)
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v731)))
	if v836 == int32(8) {
		goto L195
	} else {
		goto L196
	}
L193:
	;
	v860 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+13)) = uint8(v860)
	goto L186
L194:
	;
	if v852 == int32(0) {
		goto L193
	} else {
		goto L201
	}
L195:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v731)+4))
	if v839 != int32(1) {
		goto L193
	} else {
		goto L198
	}
L196:
	;
	goto L197
L197:
	;
	v849 = F_expression_tree_walker_impl(m, v731, int32(960), v56+int32(240))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L4
	} else {
		goto L200
	}
L198:
	;
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v731)+8))
	v844 = F_bms_add_member(m, int32(0), v843)
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L4
	} else {
		goto L199
	}
L199:
	;
	v852 = v844
	goto L194
L200:
	;
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v56)+240))
	v852 = v851
	goto L194
L201:
	;
	v855 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+14)) = uint8(v855)
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v857 == int32(2) {
		goto L186
	} else {
		goto L202
	}
L202:
	;
	v1543 = v64
	v1544 = v65
	v1550 = v71
	v1553 = v74
	goto L15
L203:
	;
	if v862 != int32(105) {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v866 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+12)) = uint8(v866)
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v868 == int32(0) {
		v1543 = v64
		v1544 = v65
		v1550 = v71
		v1553 = v74
		goto L15
	} else {
		goto L207
	}
L205:
	;
	goto L206
L206:
	;
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v731)))
	if v871 != int32(35) {
		goto L209
	} else {
		goto L210
	}
L207:
	;
	goto L206
L208:
	;
	if v964 == int32(0) {
		goto L230
	} else {
		goto L231
	}
L209:
	;
	if v871 != int32(7) {
		v1543 = v64
		v1544 = v65
		v1550 = v71
		v1553 = v74
		goto L15
	} else {
		goto L212
	}
L210:
	;
	goto L211
L211:
	;
	v955 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v731)+20)))
	if v955 != 0 {
		v1543 = v64
		v1544 = v65
		v1550 = v71
		v1553 = v74
		goto L15
	} else {
		goto L228
	}
L212:
	;
	v876 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v731)+32)))
	if v876 != 0 {
		goto L74
	} else {
		goto L213
	}
L213:
	;
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v731)+24))
	v878 = F_pg_detoast_datum(m, v877)
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L4
	} else {
		goto L214
	}
L214:
	;
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v878)+12))
	F_get_typlenbyvalalign(m, v880, v56+int32(224), v56+int32(223), v56+int32(222))
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L4
	} else {
		goto L215
	}
L215:
	;
	v890 = int32(*(*int16)(unsafe.Add(mBase, uint32(v56)+224)))
	v891 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+223)))
	v892 = int32(*(*int8)(unsafe.Add(mBase, uint32(v56)+222)))
	F_deconstruct_array(m, v878, v890, v891, v892, v56+int32(240), v56+int32(276), v56+int32(228))
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L4
	} else {
		goto L216
	}
L216:
	;
	v901 = int32(0)
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v56)+228))
	if v902 <= v901 {
		v964 = v901
		goto L208
	} else {
		goto L217
	}
L217:
	;
	v910 = v902
	v913 = v901
	v914 = int32(0)
	goto L218
L218:
	;
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v56)+276))
	v930 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v928+v914))))
	if v930 == int32(1) {
		goto L221
	} else {
		goto L222
	}
L219:
	;
	v964 = v951
	goto L208
L220:
	;
	v953 = v914 + int32(1)
	if v953 < v950 {
		v910 = v950
		v913 = v951
		v914 = v953
		goto L218
	} else {
		goto L227
	}
L221:
	;
	v933 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+20)))
	if v933 != 0 {
		v950 = v910
		v951 = v913
		goto L220
	} else {
		goto L224
	}
L222:
	;
	goto L223
L223:
	;
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v878)+12))
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v731)+12))
	v937 = int32(*(*int16)(unsafe.Add(mBase, uint32(v56)+224)))
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v56)+240))
	v942 = *(*int64)(unsafe.Add(mBase, uint32(v938+v914<<(uint(int32(3))%32))))
	v944 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+223)))
	v945 = F_makeConst(m, v934, int32(-1), v936, v937, v942, int32(0), v944)
	mBase = m.M
	v946 = m.ExcPending
	if v946 != 0 {
		goto L4
	} else {
		goto L225
	}
L224:
	;
	goto L74
L225:
	;
	v947 = F_lappend(m, v913, v945)
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L4
	} else {
		goto L226
	}
L226:
	;
	v949 = *(*int32)(unsafe.Add(mBase, uint32(v56)+228))
	v950 = v949
	v951 = v947
	goto L220
L227:
	;
	goto L219
L228:
	;
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v731)+16))
	v964 = v956
	goto L208
L229:
	;
	v1086 = F_gen_partprune_steps_internal(m, v53, v1068)
	mBase = m.M
	v1087 = m.ExcPending
	if v1087 != 0 {
		goto L4
	} else {
		goto L248
	}
L230:
	;
	v1068 = int32(0)
	goto L229
L231:
	;
	goto L232
L232:
	;
	v982 = int32(0)
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v964)+4))
	if v982 < v984 {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	v991 = v982
	v995 = v982
	goto L236
L234:
	;
	v1026 = v982
	goto L235
L235:
	;
	v1044 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+20)))
	if v1044 != int32(1) {
		v1068 = v1026
		goto L229
	} else {
		goto L241
	}
L236:
	;
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v964)+12))
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(v1009+v995<<(uint(int32(2))%32))))
	v1014 = F_make_opclause(m, v728, v766, v1013, v727)
	mBase = m.M
	v1015 = m.ExcPending
	if v1015 != 0 {
		goto L4
	} else {
		goto L238
	}
L237:
	;
	v1026 = v1016
	goto L235
L238:
	;
	v1016 = F_lappend(m, v991, v1014)
	mBase = m.M
	v1017 = m.ExcPending
	if v1017 != 0 {
		goto L4
	} else {
		goto L239
	}
L239:
	;
	v1019 = v995 + int32(1)
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(v964)+4))
	if v1019 < v1020 {
		v991 = v1016
		v995 = v1019
		goto L236
	} else {
		goto L240
	}
L240:
	;
	goto L237
L241:
	;
	if v1026 == int32(0) {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	v1068 = int32(0)
	goto L229
L243:
	;
	goto L244
L244:
	;
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v1026)+4))
	if v1050 < int32(2) {
		v1068 = v1026
		goto L229
	} else {
		goto L245
	}
L245:
	;
	v1055 = F_makeBoolExpr(m, int32(1), v1026, int32(-1))
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L4
	} else {
		goto L246
	}
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+56)) = v1055
	*(*int32)(unsafe.Add(mBase, uint32(v56)+216)) = v1055
	v1062 = F_list_make1_impl(m, int32(1), v56+int32(56))
	mBase = m.M
	v1063 = m.ExcPending
	if v1063 != 0 {
		goto L4
	} else {
		goto L247
	}
L247:
	;
	v1068 = v1062
	goto L229
L248:
	;
	v1088 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+15)))
	if v1088 != 0 {
		goto L74
	} else {
		goto L249
	}
L249:
	;
	if v1086 == int32(0) {
		v1543 = v64
		v1544 = v65
		v1550 = v71
		v1553 = v74
		goto L15
	} else {
		goto L250
	}
L250:
	;
	v1475 = v1086
	goto L67
L251:
	;
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	v1094 = F_strip_noop_phvs(m, v1093)
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L4
	} else {
		goto L252
	}
L252:
	;
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v1094)))
	if v1096 == int32(27) {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v1101 = v1094
	goto L256
L254:
	;
	v1127 = v1094
	goto L255
L255:
	;
	v1147 = F_equal(m, v1127, v293)
	mBase = m.M
	v1148 = m.ExcPending
	if v1148 != 0 {
		goto L4
	} else {
		goto L259
	}
L256:
	;
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+4))
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(v1121)))
	if v1122 == int32(27) {
		v1101 = v1121
		goto L256
	} else {
		goto L258
	}
L257:
	;
	v1127 = v1121
	goto L255
L258:
	;
	goto L257
L259:
	;
	if v1147 == int32(0) {
		goto L68
	} else {
		goto L260
	}
L260:
	;
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
	if v1151 == int32(1) {
		goto L75
	} else {
		goto L261
	}
L261:
	;
	goto L76
L262:
	;
	v1189 = F_bms_add_member(m, v65, v272)
	mBase = m.M
	v1190 = m.ExcPending
	if v1190 != 0 {
		goto L4
	} else {
		goto L268
	}
L263:
	;
	if v1176 == int32(0) {
		goto L264
	} else {
		goto L265
	}
L264:
	;
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(v56+int32(80)+v287)))
	if v1183 == int32(0) {
		goto L262
	} else {
		goto L267
	}
L265:
	;
	goto L266
L266:
	;
	v1186 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+15)) = uint8(v1186)
	v2752 = v56
	v2760 = int32(0)
	goto L1
L267:
	;
	goto L266
L268:
	;
	v1543 = v64
	v1544 = v1189
	v1550 = v71
	v1553 = v74
	goto L15
L269:
	;
	if v1213 != 0 {
		goto L270
	} else {
		goto L271
	}
L270:
	;
	v1215 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+15)) = uint8(v1215)
	v2752 = v56
	v2760 = int32(0)
	goto L1
L271:
	;
	goto L272
L272:
	;
	v1218 = F_bms_add_member(m, v71, v272)
	mBase = m.M
	v1219 = m.ExcPending
	if v1219 != 0 {
		goto L4
	} else {
		goto L273
	}
L273:
	;
	v1543 = v64
	v1544 = v65
	v1550 = v1218
	v1553 = v74
	goto L15
L274:
	;
	v1249 = *(*int32)(unsafe.Add(mBase, uint32(v85)+24))
	if v297 != v1249 {
		v1405 = v488
		v1406 = int32(0)
		goto L72
	} else {
		goto L277
	}
L275:
	;
	goto L276
L276:
	;
	v1252 = F_op_in_opfamily(m, v1246, v301)
	mBase = m.M
	v1253 = m.ExcPending
	if v1253 != 0 {
		goto L4
	} else {
		goto L279
	}
L277:
	;
	goto L276
L278:
	;
	v1289 = int32(5)
	v1290 = F_op_strict(m, v1246)
	mBase = m.M
	v1291 = m.ExcPending
	if v1291 != 0 {
		goto L4
	} else {
		goto L291
	}
L279:
	;
	if v1252 != 0 {
		goto L280
	} else {
		goto L281
	}
L280:
	;
	F_get_op_opfamily_properties(m, v1246, v301, int32(0), v56+int32(224), v56+int32(276), v56+int32(228))
	mBase = m.M
	v1262 = m.ExcPending
	if v1262 != 0 {
		goto L4
	} else {
		goto L283
	}
L281:
	;
	goto L282
L282:
	;
	v1263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294))))
	if v1263 != int32(108) {
		goto L70
	} else {
		goto L284
	}
L283:
	;
	v1287 = v1246
	goto L278
L284:
	;
	v1266 = F_get_negator(m, v1246)
	mBase = m.M
	v1267 = m.ExcPending
	if v1267 != 0 {
		goto L4
	} else {
		goto L285
	}
L285:
	;
	if v1266 == int32(0) {
		goto L71
	} else {
		goto L286
	}
L286:
	;
	v1270 = int32(0)
	v1271 = F_op_in_opfamily(m, v1266, v301)
	mBase = m.M
	v1272 = m.ExcPending
	if v1272 != 0 {
		goto L4
	} else {
		goto L287
	}
L287:
	;
	if v1271 == int32(0) {
		v1405 = v488
		v1406 = v1270
		goto L72
	} else {
		goto L288
	}
L288:
	;
	F_get_op_opfamily_properties(m, v1266, v301, int32(0), v56+int32(224), v56+int32(276), v56+int32(228))
	mBase = m.M
	v1283 = m.ExcPending
	if v1283 != 0 {
		goto L4
	} else {
		goto L289
	}
L289:
	;
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(v56)+224))
	if v1284 != int32(3) {
		v1405 = v488
		v1406 = v1270
		goto L72
	} else {
		goto L290
	}
L290:
	;
	v1287 = v1266
	goto L278
L291:
	;
	if v1290 == int32(0) {
		v1405 = v488
		v1406 = v1289
		goto L72
	} else {
		goto L292
	}
L292:
	;
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(v1245)))
	if v1294 == int32(7) {
		goto L293
	} else {
		goto L294
	}
L293:
	;
	v1332 = F_op_volatile(m, v1246)
	mBase = m.M
	v1333 = m.ExcPending
	if v1333 != 0 {
		goto L4
	} else {
		goto L310
	}
L294:
	;
	v1297 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v1297 == int32(0) {
		v1405 = v488
		v1406 = v1289
		goto L72
	} else {
		goto L295
	}
L295:
	;
	v1300 = F_contain_var_clause(m, v1245)
	mBase = m.M
	v1301 = m.ExcPending
	if v1301 != 0 {
		goto L4
	} else {
		goto L296
	}
L296:
	;
	if v1300 != 0 {
		v1405 = v488
		v1406 = v1289
		goto L72
	} else {
		goto L297
	}
L297:
	;
	v1302 = F_contain_volatile_functions(m, v1245)
	mBase = m.M
	v1303 = m.ExcPending
	if v1303 != 0 {
		goto L4
	} else {
		goto L298
	}
L298:
	;
	if v1302 != 0 {
		v1405 = v488
		v1406 = v1289
		goto L72
	} else {
		goto L299
	}
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+240)) = int32(0)
	v1306 = *(*int32)(unsafe.Add(mBase, uint32(v1245)))
	if v1306 == int32(8) {
		goto L302
	} else {
		goto L303
	}
L300:
	;
	v1330 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+13)) = uint8(v1330)
	goto L293
L301:
	;
	if v1322 == int32(0) {
		goto L300
	} else {
		goto L308
	}
L302:
	;
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(v1245)+4))
	if v1309 != int32(1) {
		goto L300
	} else {
		goto L305
	}
L303:
	;
	goto L304
L304:
	;
	v1319 = F_expression_tree_walker_impl(m, v1245, int32(960), v56+int32(240))
	mBase = m.M
	v1320 = m.ExcPending
	if v1320 != 0 {
		goto L4
	} else {
		goto L307
	}
L305:
	;
	v1313 = *(*int32)(unsafe.Add(mBase, uint32(v1245)+8))
	v1314 = F_bms_add_member(m, int32(0), v1313)
	mBase = m.M
	v1315 = m.ExcPending
	if v1315 != 0 {
		goto L4
	} else {
		goto L306
	}
L306:
	;
	v1322 = v1314
	goto L301
L307:
	;
	v1321 = *(*int32)(unsafe.Add(mBase, uint32(v56)+240))
	v1322 = v1321
	goto L301
L308:
	;
	v1325 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+14)) = uint8(v1325)
	v1327 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v1327 == int32(2) {
		goto L293
	} else {
		goto L309
	}
L309:
	;
	v1405 = v488
	v1406 = v1289
	goto L72
L310:
	;
	if v1332 != int32(105) {
		goto L311
	} else {
		goto L312
	}
L311:
	;
	v1336 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+12)) = uint8(v1336)
	v1338 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v1338 == int32(0) {
		v1405 = v488
		v1406 = v1289
		goto L72
	} else {
		goto L314
	}
L312:
	;
	goto L313
L313:
	;
	v1341 = *(*int32)(unsafe.Add(mBase, uint32(v56)+228))
	v1342 = *(*int32)(unsafe.Add(mBase, uint32(v294)+8))
	v1344 = *(*int32)(unsafe.Add(mBase, uint32(v1342+v287)))
	if v1341 == v1344 {
		goto L316
	} else {
		goto L317
	}
L314:
	;
	goto L313
L315:
	;
	v1387 = F_palloc(m, int32(24))
	mBase = m.M
	v1388 = m.ExcPending
	if v1388 != 0 {
		goto L4
	} else {
		goto L329
	}
L316:
	;
	v1346 = *(*int32)(unsafe.Add(mBase, uint32(v294)+24))
	v1350 = *(*int32)(unsafe.Add(mBase, uint32(v1346+v272*int32(28))+4))
	v1385 = v1350
	goto L315
L317:
	;
	goto L318
L318:
	;
	v1351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294))))
	switch v1351 - int32(104) {
	case 0:
		goto L322
	default:
		goto L321
	case 4, 10:
		goto L320
	}
L319:
	;
	if v1382 == int32(0) {
		goto L71
	} else {
		goto L328
	}
L320:
	;
	v1376 = *(*int32)(unsafe.Add(mBase, uint32(v294)+4))
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(v1376+v287)))
	v1380 = F_get_opfamily_proc(m, v1378, v1344, v1341, int32(1))
	mBase = m.M
	v1381 = m.ExcPending
	if v1381 != 0 {
		goto L4
	} else {
		goto L327
	}
L321:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1363 = m.ExcPending
	if v1363 != 0 {
		goto L4
	} else {
		goto L324
	}
L322:
	;
	v1354 = *(*int32)(unsafe.Add(mBase, uint32(v294)+4))
	v1356 = *(*int32)(unsafe.Add(mBase, uint32(v1354+v287)))
	v1358 = F_get_opfamily_proc(m, v1356, v1341, v1341, int32(2))
	mBase = m.M
	v1359 = m.ExcPending
	if v1359 != 0 {
		goto L4
	} else {
		goto L323
	}
L323:
	;
	v1382 = v1358
	goto L319
L324:
	;
	v1364 = int32(*(*int8)(unsafe.Add(mBase, uint32(v294))))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+48)) = v1364
	F_errmsg_internal(m, int32(_a_F_gen_partprune_steps_internal_0), v56+int32(48))
	mBase = m.M
	v1370 = m.ExcPending
	if v1370 != 0 {
		goto L4
	} else {
		goto L325
	}
L325:
	;
	F_errfinish(m, int32(_a_F_gen_partprune_steps_internal_1), int32(2149), int32(_a_F_gen_partprune_steps_internal_2))
	mBase = m.M
	v1375 = m.ExcPending
	if v1375 != 0 {
		goto L4
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
	v1382 = v1380
	goto L319
L328:
	;
	v1385 = v1382
	goto L315
L329:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1387))) = v272
	v1390 = int32(1)
	v1391 = *(*int32)(unsafe.Add(mBase, uint32(v56)+224))
	v1393 = v1252 ^ v1390
	*(*uint8)(unsafe.Add(mBase, uint32(v1387)+8)) = uint8(v1393)
	*(*int32)(unsafe.Add(mBase, uint32(v1387)+4)) = v1287
	*(*int32)(unsafe.Add(mBase, uint32(v1387)+16)) = v1385
	*(*int32)(unsafe.Add(mBase, uint32(v1387)+12)) = v1245
	if v1252 != 0 {
		goto L330
	} else {
		goto L331
	}
L330:
	;
	v1399 = v1391
	goto L332
L331:
	;
	v1399 = int32(0)
	goto L332
L332:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1387)+20)) = v1399
	v1405 = v1387
	v1406 = v1390
	goto L72
L333:
	;
	switch v1431 - int32(1) {
	case 0:
		v1506 = v1430
		goto L64
	default:
		v1543 = v64
		v1544 = v65
		v1550 = v71
		v1553 = v74
		goto L15
	case 2:
		v1475 = v1423
		goto L67
	}
L334:
	;
	v1543 = v64
	v1544 = v65
	v1550 = v71
	v1553 = v74
	goto L15
L335:
	;
	v1543 = v1495
	v1544 = v65
	v1550 = v71
	v1553 = v74
	goto L15
L336:
	;
	if v1519 != 0 {
		goto L337
	} else {
		goto L338
	}
L337:
	;
	v1521 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+15)) = uint8(v1521)
	v2752 = v56
	v2760 = int32(0)
	goto L1
L338:
	;
	goto L339
L339:
	;
	v1526 = v56 + int32(80) + v287
	v1527 = *(*int32)(unsafe.Add(mBase, uint32(v1526)))
	v1528 = F_lappend(m, v1527, v1506)
	mBase = m.M
	v1529 = m.ExcPending
	if v1529 != 0 {
		goto L4
	} else {
		goto L340
	}
L340:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1526))) = v1528
	v1543 = v64
	v1544 = v65
	v1550 = v71
	v1553 = int32(1)
	goto L15
L341:
	;
	goto L14
L342:
	;
	v2554 = F_palloc0(m, int32(24))
	mBase = m.M
	v2555 = m.ExcPending
	if v2555 != 0 {
		goto L4
	} else {
		goto L509
	}
L343:
	;
	if v1579 == int32(0) {
		v2574 = v1558
		v2577 = v1561
		v2585 = v1569
		v2592 = v1576
		v2593 = v1577
		goto L8
	} else {
		goto L362
	}
L344:
	;
	v1582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1577))))
	switch v1582 - int32(104) {
	case 0:
		goto L345
	default:
		v1633 = v1570
		goto L343
	case 4, 10:
		goto L342
	}
L345:
	;
	v1586 = int64(0)
	if v1570 == int32(0) {
		goto L347
	} else {
		goto L348
	}
L346:
	;
	v1631 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1577)+2)))
	if v1630 == v1631 {
		goto L342
	} else {
		goto L361
	}
L347:
	;
	v1630 = int32(0)
	goto L346
L348:
	;
	goto L349
L349:
	;
	v1591 = v1570 + int32(8)
	v1592 = *(*int32)(unsafe.Add(mBase, uint32(v1570)+4))
	if v1592 == int32(1) {
		goto L350
	} else {
		goto L351
	}
L350:
	;
	v1595 = *(*int32)(unsafe.Add(mBase, uint32(v1591)))
	v1630 = base.I32_popcnt(v1595)
	goto L346
L351:
	;
	goto L352
L352:
	;
	v1598 = v1592 << (uint(int32(2)) % 32)
	if v1598 <= int32(7) {
		goto L354
	} else {
		goto L355
	}
L353:
	;
	v1630 = base.I32_wrap_i64(v1625)
	goto L346
L354:
	;
	if v1598 == int32(0) {
		v1625 = v1586
		goto L353
	} else {
		goto L357
	}
L355:
	;
	goto L356
L356:
	;
	v1622 = F_pg_popcount_optimized(m, v1591, v1598)
	mBase = m.M
	v1625 = v1622
	goto L353
L357:
	;
	v1603 = v1598
	v1604 = v1591
	v1605 = v1586
	goto L358
L358:
	;
	v1606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1604)+3)))
	v1607 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1606)+uint32(_c_F_gen_partprune_steps_internal[0]))))
	v1608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1604)+2)))
	v1609 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1608)+uint32(_c_F_gen_partprune_steps_internal[0]))))
	v1610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1604)+1)))
	v1611 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1610)+uint32(_c_F_gen_partprune_steps_internal[0]))))
	v1612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1604))))
	v1613 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1612)+uint32(_c_F_gen_partprune_steps_internal[0]))))
	v1617 = v1607 + (v1609 + (v1611 + (v1605 + v1613)))
	v1618 = int32(4)
	v1621 = v1603 - v1618
	if v1621 != 0 {
		v1603 = v1621
		v1604 = v1604 + v1618
		v1605 = v1617
		goto L358
	} else {
		goto L360
	}
L359:
	;
	v1625 = v1617
	goto L353
L360:
	;
	goto L359
L361:
	;
	v1633 = v1570
	goto L343
L362:
	;
	v1636 = *(*int32)(unsafe.Add(mBase, uint32(v1558)))
	v1637 = *(*int32)(unsafe.Add(mBase, uint32(v1636)+256))
	v1638 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1561)+240)) = v1638
	*(*int64)(unsafe.Add(mBase, uint32(v1561)+248)) = v1638
	*(*int64)(unsafe.Add(mBase, uint32(v1561)+256)) = v1638
	*(*int64)(unsafe.Add(mBase, uint32(v1561)+232)) = v1638
	v1646 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1637)+2)))
	if v1646 <= int32(0) {
		goto L368
	} else {
		goto L369
	}
L363:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2540 = m.ExcPending
	if v2540 != 0 {
		goto L4
	} else {
		goto L506
	}
L364:
	;
	v2535 = F_list_concat(m, v2524, v2531)
	mBase = m.M
	v2536 = m.ExcPending
	if v2536 != 0 {
		goto L4
	} else {
		goto L505
	}
L365:
	;
	v2513 = v1558
	v2516 = v1561
	v2524 = v1569
	v2531 = int32(0)
	goto L364
L366:
	;
	v2399 = *(*int32)(unsafe.Add(mBase, uint32(v1561)+236))
	if v2399 == int32(0) {
		goto L365
	} else {
		goto L491
	}
L367:
	;
	v1889 = *(*int32)(unsafe.Add(mBase, uint32(v1561)+256))
	v1890 = *(*int32)(unsafe.Add(mBase, uint32(v1561)+248))
	v1891 = *(*int32)(unsafe.Add(mBase, uint32(v1561)+252))
	v1894 = v1558
	v1897 = v1561
	v1901 = v1889
	v1905 = v1569
	v1906 = v1890
	v1907 = int32(1)
	v1911 = v1891
	v1912 = int32(0)
	goto L406
L368:
	;
	v1864 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1637))))
	switch v1864 - int32(104) {
	case 0:
		goto L366
	default:
		goto L363
	case 4, 10:
		goto L367
	}
L369:
	;
	v1657 = int32(0)
	goto L370
L370:
	;
	v1673 = v1657 << (uint(int32(2)) % 32)
	v1677 = *(*int32)(unsafe.Add(mBase, uint32(v1673+(v1561+int32(80)))))
	v1680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1637))))
	if base.B2i32(v1677 == int32(0))&base.B2i32(v1680 == int32(114)) != 0 {
		goto L367
	} else {
		goto L372
	}
L371:
	;
	goto L368
L372:
	;
	if base.B2i32(v1680 != int32(104))|v1677 == int32(0) {
		goto L374
	} else {
		goto L375
	}
L373:
	;
	v1839 = v1657 + int32(1)
	v1840 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1637)+2)))
	if v1839 < v1840 {
		v1657 = v1839
		goto L370
	} else {
		goto L405
	}
L374:
	;
	v1689 = F_bms_is_member(m, v1657, v1633)
	mBase = m.M
	v1690 = m.ExcPending
	if v1690 != 0 {
		goto L4
	} else {
		goto L377
	}
L375:
	;
	goto L376
L376:
	;
	if v1677 == int32(0) {
		goto L373
	} else {
		goto L379
	}
L377:
	;
	if v1689 != 0 {
		goto L373
	} else {
		goto L378
	}
L378:
	;
	goto L365
L379:
	;
	v1694 = int32(0)
	v1695 = *(*int32)(unsafe.Add(mBase, uint32(v1677)+4))
	if v1695 <= v1694 {
		goto L373
	} else {
		goto L380
	}
L380:
	;
	v1702 = int32(1)
	v1703 = v1694
	goto L382
L381:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1803 = m.ExcPending
	if v1803 != 0 {
		goto L4
	} else {
		goto L402
	}
L382:
	;
	v1720 = *(*int32)(unsafe.Add(mBase, uint32(v1677)+12))
	v1724 = *(*int32)(unsafe.Add(mBase, uint32(v1720+v1703<<(uint(int32(2))%32))))
	v1726 = v1724 + int32(20)
	v1727 = *(*int32)(unsafe.Add(mBase, uint32(v1724)+20))
	if v1727 == int32(0) {
		goto L384
	} else {
		goto L385
	}
L383:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1790 = m.ExcPending
	if v1790 != 0 {
		goto L4
	} else {
		goto L399
	}
L384:
	;
	v1730 = *(*int32)(unsafe.Add(mBase, uint32(v1724)+4))
	v1731 = *(*int32)(unsafe.Add(mBase, uint32(v1637)+4))
	v1733 = *(*int32)(unsafe.Add(mBase, uint32(v1731+v1673)))
	F_get_op_opfamily_properties(m, v1730, v1733, int32(0), v1726, v1561+int32(284), v1561+int32(280))
	mBase = m.M
	v1740 = m.ExcPending
	if v1740 != 0 {
		goto L4
	} else {
		goto L387
	}
L385:
	;
	goto L386
L386:
	;
	v1741 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1637))))
	switch v1741 - int32(104) {
	case 0:
		goto L391
	default:
		goto L381
	case 4, 10:
		goto L392
	}
L387:
	;
	goto L386
L388:
	;
	goto L383
L389:
	;
	v1784 = v1703 + int32(1)
	v1785 = *(*int32)(unsafe.Add(mBase, uint32(v1677)+4))
	if v1784 < v1785 {
		v1702 = int32(0)
		v1703 = v1784
		goto L382
	} else {
		goto L398
	}
L390:
	;
	v1777 = v1703 + int32(1)
	v1778 = *(*int32)(unsafe.Add(mBase, uint32(v1677)+4))
	if v1777 < v1778 {
		v1703 = v1777
		goto L382
	} else {
		goto L396
	}
L391:
	;
	v1760 = *(*int32)(unsafe.Add(mBase, uint32(v1726)))
	if v1760 != int32(1) {
		goto L388
	} else {
		goto L394
	}
L392:
	;
	v1745 = v1561 + int32(240)
	v1746 = *(*int32)(unsafe.Add(mBase, uint32(v1724)+20))
	v1750 = *(*int32)(unsafe.Add(mBase, uint32(v1745+v1746<<(uint(int32(2))%32))))
	v1751 = F_lappend(m, v1750, v1724)
	mBase = m.M
	v1752 = m.ExcPending
	if v1752 != 0 {
		goto L4
	} else {
		goto L393
	}
L393:
	;
	v1753 = *(*int32)(unsafe.Add(mBase, uint32(v1724)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1753<<(uint(int32(2))%32)+v1745))) = v1751
	switch v1753 - int32(1) {
	case 0, 4:
		goto L389
	default:
		goto L390
	}
L394:
	;
	v1763 = *(*int32)(unsafe.Add(mBase, uint32(v1561)+236))
	v1764 = F_lappend(m, v1763, v1724)
	mBase = m.M
	v1765 = m.ExcPending
	if v1765 != 0 {
		goto L4
	} else {
		goto L395
	}
L395:
	;
	v1768 = *(*int32)(unsafe.Add(mBase, uint32(v1724)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1561+int32(232)+v1768<<(uint(int32(2))%32)))) = v1764
	goto L390
L396:
	;
	if v1702&int32(1) != 0 {
		goto L373
	} else {
		goto L397
	}
L397:
	;
	goto L368
L398:
	;
	goto L368
L399:
	;
	F_errmsg_internal(m, int32(_a_F_gen_partprune_steps_internal_3), int32(0))
	mBase = m.M
	v1794 = m.ExcPending
	if v1794 != 0 {
		goto L4
	} else {
		goto L400
	}
L400:
	;
	F_errfinish(m, int32(_a_F_gen_partprune_steps_internal_1), int32(1480), int32(_a_F_gen_partprune_steps_internal_4))
	mBase = m.M
	v1799 = m.ExcPending
	if v1799 != 0 {
		goto L4
	} else {
		goto L401
	}
L401:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L402:
	;
	v1804 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1637))))
	*(*int32)(unsafe.Add(mBase, uint32(v1561)+32)) = v1804
	F_errmsg_internal(m, int32(_a_F_gen_partprune_steps_internal_0), v1561+int32(32))
	mBase = m.M
	v1810 = m.ExcPending
	if v1810 != 0 {
		goto L4
	} else {
		goto L403
	}
L403:
	;
	F_errfinish(m, int32(_a_F_gen_partprune_steps_internal_1), int32(1487), int32(_a_F_gen_partprune_steps_internal_4))
	mBase = m.M
	v1815 = m.ExcPending
	if v1815 != 0 {
		goto L4
	} else {
		goto L404
	}
L404:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L405:
	;
	goto L371
L406:
	;
	v1921 = *(*int32)(unsafe.Add(mBase, uint32(v1897+int32(240)+v1907<<(uint(int32(2))%32))))
	if v1921 == int32(0) {
		v2391 = v1912
		goto L408
	} else {
		goto L409
	}
L407:
	;
	v2513 = v1894
	v2516 = v1897
	v2524 = v1905
	v2531 = v2391
	goto L364
L408:
	;
	v2396 = v1907 + int32(1)
	if v2396 != int32(6) {
		v1907 = v2396
		v1912 = v2391
		goto L406
	} else {
		goto L490
	}
L409:
	;
	v1924 = *(*int32)(unsafe.Add(mBase, uint32(v1921)+4))
	if v1924 <= int32(0) {
		v2391 = v1912
		goto L408
	} else {
		goto L410
	}
L410:
	;
	v1936 = int32(0)
	v1948 = v1912
	goto L411
L411:
	;
	v1952 = *(*int32)(unsafe.Add(mBase, uint32(v1921)+12))
	v1956 = *(*int32)(unsafe.Add(mBase, uint32(v1952+v1936<<(uint(int32(2))%32))))
	v1957 = *(*int32)(unsafe.Add(mBase, uint32(v1956)))
	if v1957 == int32(0) {
		goto L414
	} else {
		goto L415
	}
L412:
	;
	v2391 = v2367
	goto L408
L413:
	;
	v2367 = F_list_concat(m, v1948, v2366)
	mBase = m.M
	v2368 = m.ExcPending
	if v2368 != 0 {
		goto L4
	} else {
		goto L488
	}
L414:
	;
	v1960 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1956)+8)))
	v1961 = *(*int32)(unsafe.Add(mBase, uint32(v1956)+16))
	v1962 = *(*int32)(unsafe.Add(mBase, uint32(v1956)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1897)+12)) = v1962
	*(*int32)(unsafe.Add(mBase, uint32(v1897)+284)) = v1962
	v1968 = F_list_make1_impl(m, int32(1), v1897+int32(12))
	mBase = m.M
	v1969 = m.ExcPending
	if v1969 != 0 {
		goto L4
	} else {
		goto L417
	}
L415:
	;
	goto L416
L416:
	;
	v2005 = int32(0)
	if v1911 != 0 {
		goto L425
	} else {
		goto L426
	}
L417:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1897)+8)) = v1961
	*(*int32)(unsafe.Add(mBase, uint32(v1897)+280)) = v1961
	v1975 = F_list_make1_impl(m, int32(480), v1897+int32(8))
	mBase = m.M
	v1976 = m.ExcPending
	if v1976 != 0 {
		goto L4
	} else {
		goto L418
	}
L418:
	;
	v1978 = F_palloc0(m, int32(24))
	mBase = m.M
	v1979 = m.ExcPending
	if v1979 != 0 {
		goto L4
	} else {
		goto L419
	}
L419:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1978))) = int32(381)
	v1982 = *(*int32)(unsafe.Add(mBase, uint32(v1894)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1894)+16)) = v1982 + int32(1)
	v1986 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1978)+20)) = v1986
	*(*int32)(unsafe.Add(mBase, uint32(v1978)+16)) = v1975
	*(*int32)(unsafe.Add(mBase, uint32(v1978)+12)) = v1968
	if v1960 != 0 {
		goto L420
	} else {
		goto L421
	}
L420:
	;
	v1991 = v1986
	goto L422
L421:
	;
	v1991 = v1907
	goto L422
L422:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1978)+8)) = uint16(v1991)
	*(*int32)(unsafe.Add(mBase, uint32(v1978)+4)) = v1982
	v1994 = *(*int32)(unsafe.Add(mBase, uint32(v1894)+8))
	v1995 = F_lappend(m, v1994, v1978)
	mBase = m.M
	v1996 = m.ExcPending
	if v1996 != 0 {
		goto L4
	} else {
		goto L423
	}
L423:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1894)+8)) = v1995
	*(*int32)(unsafe.Add(mBase, uint32(v1897)+4)) = v1978
	*(*int32)(unsafe.Add(mBase, uint32(v1897)+276)) = v1978
	v2003 = F_list_make1_impl(m, int32(1), v1897+int32(4))
	mBase = m.M
	v2004 = m.ExcPending
	if v2004 != 0 {
		goto L4
	} else {
		goto L424
	}
L424:
	;
	v2366 = v2003
	goto L413
L425:
	;
	v2007 = *(*int32)(unsafe.Add(mBase, uint32(v1911)+12))
	v2008 = v2007
	goto L427
L426:
	;
	v2008 = v2005
	goto L427
L427:
	;
	if v1906 != 0 {
		goto L428
	} else {
		goto L429
	}
L428:
	;
	v2009 = *(*int32)(unsafe.Add(mBase, uint32(v1906)+12))
	v2010 = v2009
	goto L430
L429:
	;
	v2010 = v2005
	goto L430
L430:
	;
	v2011 = int32(0)
	if v1901 != 0 {
		goto L431
	} else {
		goto L432
	}
L431:
	;
	v2013 = *(*int32)(unsafe.Add(mBase, uint32(v1901)+12))
	v2014 = v2013
	goto L433
L432:
	;
	v2014 = v2011
	goto L433
L433:
	;
	v2015 = int32(0)
	if v2015 < v1957 {
		goto L434
	} else {
		goto L435
	}
L434:
	;
	v2019 = v2010
	v2020 = v2015
	v2023 = v2011
	v2026 = v2008
	v2028 = v2014
	goto L437
L435:
	;
	v2321 = v2011
	goto L436
L436:
	;
	v2338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1956)+8)))
	v2339 = *(*int32)(unsafe.Add(mBase, uint32(v1956)+12))
	v2340 = *(*int32)(unsafe.Add(mBase, uint32(v1956)+16))
	v2342 = F_get_steps_using_prefix(m, v1894, v1907, v2338, v2339, v2340, int32(0), v2321)
	mBase = m.M
	v2343 = m.ExcPending
	if v2343 != 0 {
		goto L4
	} else {
		goto L487
	}
L437:
	;
	v2040 = int32(0)
	if v2026 == v2040 {
		v2101 = v2023
		v2105 = v2040
		goto L440
	} else {
		goto L441
	}
L438:
	;
	v2321 = v2295
	goto L436
L439:
	;
	if base.Ui32(int32(2)) < base.Ui32(v1907) {
		v2203 = v2019
		v2207 = v2124
		v2211 = v2128
		goto L451
	} else {
		goto L452
	}
L440:
	;
	v2124 = v2101
	v2127 = int32(0)
	v2128 = v2105
	goto L439
L441:
	;
	v2043 = *(*int32)(unsafe.Add(mBase, uint32(v1911)+12))
	v2044 = v2026 - v2043
	v2046 = v2044 >> (uint(int32(2)) % 32)
	v2047 = *(*int32)(unsafe.Add(mBase, uint32(v1911)+4))
	if v2047 <= v2046 {
		v2101 = v2023
		v2105 = v2040
		goto L440
	} else {
		goto L442
	}
L442:
	;
	v2049 = *(*int32)(unsafe.Add(mBase, uint32(v1911)+12))
	v2050 = v2049 + v2044
	v2051 = *(*int32)(unsafe.Add(mBase, uint32(v2050)))
	v2052 = *(*int32)(unsafe.Add(mBase, uint32(v2051)))
	if v2052 != v2020 {
		v2124 = v2023
		v2127 = v2050
		v2128 = v2040
		goto L439
	} else {
		goto L443
	}
L443:
	;
	v2054 = int32(1)
	v2055 = F_lappend(m, v2023, v2051)
	mBase = m.M
	v2056 = m.ExcPending
	if v2056 != 0 {
		goto L4
	} else {
		goto L444
	}
L444:
	;
	v2058 = v2046 + int32(1)
	v2059 = *(*int32)(unsafe.Add(mBase, uint32(v1911)+4))
	if v2059 <= v2058 {
		v2101 = v2055
		v2105 = v2054
		goto L440
	} else {
		goto L445
	}
L445:
	;
	v2065 = v2058
	v2066 = v2055
	goto L446
L446:
	;
	v2083 = *(*int32)(unsafe.Add(mBase, uint32(v1911)+12))
	v2086 = v2083 + v2065<<(uint(int32(2))%32)
	v2087 = *(*int32)(unsafe.Add(mBase, uint32(v2086)))
	v2088 = *(*int32)(unsafe.Add(mBase, uint32(v2087)))
	if v2088 != v2020 {
		v2124 = v2066
		v2127 = v2086
		v2128 = v2054
		goto L439
	} else {
		goto L448
	}
L447:
	;
	v2101 = v2090
	v2105 = v2054
	goto L440
L448:
	;
	v2090 = F_lappend(m, v2066, v2087)
	mBase = m.M
	v2091 = m.ExcPending
	if v2091 != 0 {
		goto L4
	} else {
		goto L449
	}
L449:
	;
	v2093 = v2065 + int32(1)
	v2094 = *(*int32)(unsafe.Add(mBase, uint32(v1911)+4))
	if v2093 < v2094 {
		v2065 = v2093
		v2066 = v2090
		goto L446
	} else {
		goto L450
	}
L450:
	;
	goto L447
L451:
	;
	if v1907&int32(6) != int32(4) {
		v2286 = v2028
		goto L470
	} else {
		goto L471
	}
L452:
	;
	if v2019 == int32(0) {
		goto L453
	} else {
		goto L454
	}
L453:
	;
	v2203 = int32(0)
	v2207 = v2124
	v2211 = v2128
	goto L451
L454:
	;
	goto L455
L455:
	;
	v2146 = *(*int32)(unsafe.Add(mBase, uint32(v1906)+12))
	v2147 = v2019 - v2146
	v2149 = v2147 >> (uint(int32(2)) % 32)
	v2150 = *(*int32)(unsafe.Add(mBase, uint32(v1906)+4))
	if v2150 <= v2149 {
		goto L456
	} else {
		goto L457
	}
L456:
	;
	v2203 = int32(0)
	v2207 = v2124
	v2211 = v2128
	goto L451
L457:
	;
	goto L458
L458:
	;
	v2153 = *(*int32)(unsafe.Add(mBase, uint32(v1906)+12))
	v2154 = v2153 + v2147
	v2155 = *(*int32)(unsafe.Add(mBase, uint32(v2154)))
	v2156 = *(*int32)(unsafe.Add(mBase, uint32(v2155)))
	if v2156 != v2020 {
		v2203 = v2154
		v2207 = v2124
		v2211 = v2128
		goto L451
	} else {
		goto L459
	}
L459:
	;
	v2158 = int32(1)
	v2159 = F_lappend(m, v2124, v2155)
	mBase = m.M
	v2160 = m.ExcPending
	if v2160 != 0 {
		goto L4
	} else {
		goto L460
	}
L460:
	;
	v2162 = v2149 + int32(1)
	v2163 = *(*int32)(unsafe.Add(mBase, uint32(v1906)+4))
	if v2163 <= v2162 {
		goto L461
	} else {
		goto L462
	}
L461:
	;
	v2203 = int32(0)
	v2207 = v2159
	v2211 = v2158
	goto L451
L462:
	;
	goto L463
L463:
	;
	v2170 = v2162
	v2171 = v2159
	goto L464
L464:
	;
	v2188 = *(*int32)(unsafe.Add(mBase, uint32(v1906)+12))
	v2191 = v2188 + v2170<<(uint(int32(2))%32)
	v2192 = *(*int32)(unsafe.Add(mBase, uint32(v2191)))
	v2193 = *(*int32)(unsafe.Add(mBase, uint32(v2192)))
	if v2193 != v2020 {
		v2203 = v2191
		v2207 = v2171
		v2211 = v2158
		goto L451
	} else {
		goto L466
	}
L465:
	;
	v2203 = int32(0)
	v2207 = v2195
	v2211 = v2158
	goto L451
L466:
	;
	v2195 = F_lappend(m, v2171, v2192)
	mBase = m.M
	v2196 = m.ExcPending
	if v2196 != 0 {
		goto L4
	} else {
		goto L467
	}
L467:
	;
	v2198 = v2170 + int32(1)
	v2199 = *(*int32)(unsafe.Add(mBase, uint32(v1906)+4))
	if v2198 < v2199 {
		v2170 = v2198
		v2171 = v2195
		goto L464
	} else {
		goto L468
	}
L468:
	;
	goto L465
L469:
	;
	v2313 = v2020 + int32(1)
	v2314 = *(*int32)(unsafe.Add(mBase, uint32(v1956)))
	if v2313 < v2314 {
		v2019 = v2203
		v2020 = v2313
		v2023 = v2295
		v2026 = v2127
		v2028 = v2300
		goto L437
	} else {
		goto L486
	}
L470:
	;
	if v2211 == int32(0) {
		v2391 = v1948
		goto L408
	} else {
		goto L485
	}
L471:
	;
	if v2028 == int32(0) {
		goto L473
	} else {
		goto L474
	}
L472:
	;
	v2254 = v2239
	v2258 = v2242
	goto L480
L473:
	;
	if v2211 != 0 {
		v2295 = v2207
		v2300 = int32(0)
		goto L469
	} else {
		goto L479
	}
L474:
	;
	v2228 = *(*int32)(unsafe.Add(mBase, uint32(v1901)+12))
	v2229 = v2028 - v2228
	v2231 = v2229 >> (uint(int32(2)) % 32)
	v2232 = *(*int32)(unsafe.Add(mBase, uint32(v1901)+4))
	if v2232 <= v2231 {
		goto L473
	} else {
		goto L475
	}
L475:
	;
	v2234 = *(*int32)(unsafe.Add(mBase, uint32(v1901)+12))
	v2235 = v2234 + v2229
	v2236 = *(*int32)(unsafe.Add(mBase, uint32(v2235)))
	v2237 = *(*int32)(unsafe.Add(mBase, uint32(v2236)))
	if v2237 != v2020 {
		v2286 = v2235
		goto L470
	} else {
		goto L476
	}
L476:
	;
	v2239 = F_lappend(m, v2207, v2236)
	mBase = m.M
	v2240 = m.ExcPending
	if v2240 != 0 {
		goto L4
	} else {
		goto L477
	}
L477:
	;
	v2242 = v2231 + int32(1)
	v2243 = *(*int32)(unsafe.Add(mBase, uint32(v1901)+4))
	if v2242 < v2243 {
		goto L472
	} else {
		goto L478
	}
L478:
	;
	v2295 = v2239
	v2300 = int32(0)
	goto L469
L479:
	;
	v2391 = v1948
	goto L408
L480:
	;
	v2271 = *(*int32)(unsafe.Add(mBase, uint32(v1901)+12))
	v2274 = v2271 + v2258<<(uint(int32(2))%32)
	v2275 = *(*int32)(unsafe.Add(mBase, uint32(v2274)))
	v2276 = *(*int32)(unsafe.Add(mBase, uint32(v2275)))
	if v2276 != v2020 {
		v2295 = v2254
		v2300 = v2274
		goto L469
	} else {
		goto L482
	}
L481:
	;
	v2295 = v2278
	v2300 = int32(0)
	goto L469
L482:
	;
	v2278 = F_lappend(m, v2254, v2275)
	mBase = m.M
	v2279 = m.ExcPending
	if v2279 != 0 {
		goto L4
	} else {
		goto L483
	}
L483:
	;
	v2281 = v2258 + int32(1)
	v2282 = *(*int32)(unsafe.Add(mBase, uint32(v1901)+4))
	if v2281 < v2282 {
		v2254 = v2278
		v2258 = v2281
		goto L480
	} else {
		goto L484
	}
L484:
	;
	goto L481
L485:
	;
	v2295 = v2207
	v2300 = v2286
	goto L469
L486:
	;
	goto L438
L487:
	;
	v2366 = v2342
	goto L413
L488:
	;
	v2370 = v1936 + int32(1)
	v2371 = *(*int32)(unsafe.Add(mBase, uint32(v1921)+4))
	if v2370 < v2371 {
		v1936 = v2370
		v1948 = v2367
		goto L411
	} else {
		goto L489
	}
L489:
	;
	goto L412
L490:
	;
	goto L407
L491:
	;
	v2402 = int32(0)
	v2403 = *(*int32)(unsafe.Add(mBase, uint32(v2399)+4))
	if v2403 <= v2402 {
		goto L365
	} else {
		goto L492
	}
L492:
	;
	v2406 = *(*int32)(unsafe.Add(mBase, uint32(v2399)+12))
	v2412 = *(*int32)(unsafe.Add(mBase, uint32(v2406+v2403<<(uint(int32(2))%32)-int32(4))))
	v2413 = *(*int32)(unsafe.Add(mBase, uint32(v2412)))
	v2417 = v2402
	v2423 = int32(0)
	goto L493
L493:
	;
	v2437 = *(*int32)(unsafe.Add(mBase, uint32(v2399)+12))
	v2441 = *(*int32)(unsafe.Add(mBase, uint32(v2437+v2417<<(uint(int32(2))%32))))
	v2442 = *(*int32)(unsafe.Add(mBase, uint32(v2441)))
	if v2413 == v2442 {
		goto L495
	} else {
		goto L496
	}
L494:
	;
	goto L365
L495:
	;
	v2447 = v2417
	v2463 = int32(0)
	goto L498
L496:
	;
	goto L497
L497:
	;
	v2484 = F_lappend(m, v2423, v2441)
	mBase = m.M
	v2485 = m.ExcPending
	if v2485 != 0 {
		goto L4
	} else {
		goto L503
	}
L498:
	;
	v2469 = *(*int32)(unsafe.Add(mBase, uint32(v2399)+12))
	v2473 = *(*int32)(unsafe.Add(mBase, uint32(v2469+v2447<<(uint(int32(2))%32))))
	v2474 = *(*int32)(unsafe.Add(mBase, uint32(v2473)+12))
	v2475 = *(*int32)(unsafe.Add(mBase, uint32(v2473)+16))
	v2476 = F_get_steps_using_prefix(m, v1558, int32(1), int32(0), v2474, v2475, v1633, v2423)
	mBase = m.M
	v2477 = m.ExcPending
	if v2477 != 0 {
		goto L4
	} else {
		goto L500
	}
L499:
	;
	v2513 = v1558
	v2516 = v1561
	v2524 = v1569
	v2531 = v2478
	goto L364
L500:
	;
	v2478 = F_list_concat(m, v2463, v2476)
	mBase = m.M
	v2479 = m.ExcPending
	if v2479 != 0 {
		goto L4
	} else {
		goto L501
	}
L501:
	;
	v2481 = v2447 + int32(1)
	v2482 = *(*int32)(unsafe.Add(mBase, uint32(v2399)+4))
	if v2481 < v2482 {
		v2447 = v2481
		v2463 = v2478
		goto L498
	} else {
		goto L502
	}
L502:
	;
	goto L499
L503:
	;
	v2487 = v2417 + int32(1)
	v2488 = *(*int32)(unsafe.Add(mBase, uint32(v2399)+4))
	if v2487 < v2488 {
		v2417 = v2487
		v2423 = v2484
		goto L493
	} else {
		goto L504
	}
L504:
	;
	goto L494
L505:
	;
	v2666 = v2513
	v2669 = v2516
	v2677 = v2535
	goto L7
L506:
	;
	v2541 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1637))))
	*(*int32)(unsafe.Add(mBase, uint32(v1561)+16)) = v2541
	F_errmsg_internal(m, int32(_a_F_gen_partprune_steps_internal_0), v1561+int32(16))
	mBase = m.M
	v2547 = m.ExcPending
	if v2547 != 0 {
		goto L4
	} else {
		goto L507
	}
L507:
	;
	F_errfinish(m, int32(_a_F_gen_partprune_steps_internal_1), int32(1768), int32(_a_F_gen_partprune_steps_internal_4))
	mBase = m.M
	v2552 = m.ExcPending
	if v2552 != 0 {
		goto L4
	} else {
		goto L508
	}
L508:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L509:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2554))) = int32(381)
	v2558 = *(*int32)(unsafe.Add(mBase, uint32(v1558)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1558)+16)) = v2558 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2554)+20)) = v1570
	*(*int64)(unsafe.Add(mBase, uint32(v2554)+12)) = int64(0)
	v2565 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v2554)+8)) = uint16(v2565)
	*(*int32)(unsafe.Add(mBase, uint32(v2554)+4)) = v2558
	v2568 = *(*int32)(unsafe.Add(mBase, uint32(v1558)+8))
	v2569 = F_lappend(m, v2568, v2554)
	mBase = m.M
	v2570 = m.ExcPending
	if v2570 != 0 {
		goto L4
	} else {
		goto L510
	}
L510:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1558)+8)) = v2569
	v2572 = F_lappend(m, v1569, v2554)
	mBase = m.M
	v2573 = m.ExcPending
	if v2573 != 0 {
		goto L4
	} else {
		goto L511
	}
L511:
	;
	v2666 = v1558
	v2669 = v1561
	v2677 = v2572
	goto L7
L512:
	;
	v2642 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2593)+2)))
	if v2641 != v2642 {
		v2666 = v2574
		v2669 = v2577
		v2677 = v2585
		goto L7
	} else {
		goto L527
	}
L513:
	;
	v2641 = int32(0)
	goto L512
L514:
	;
	goto L515
L515:
	;
	v2602 = v2592 + int32(8)
	v2603 = *(*int32)(unsafe.Add(mBase, uint32(v2592)+4))
	if v2603 == int32(1) {
		goto L516
	} else {
		goto L517
	}
L516:
	;
	v2606 = *(*int32)(unsafe.Add(mBase, uint32(v2602)))
	v2641 = base.I32_popcnt(v2606)
	goto L512
L517:
	;
	goto L518
L518:
	;
	v2609 = v2603 << (uint(int32(2)) % 32)
	if v2609 <= int32(7) {
		goto L520
	} else {
		goto L521
	}
L519:
	;
	v2641 = base.I32_wrap_i64(v2636)
	goto L512
L520:
	;
	if v2609 == int32(0) {
		v2636 = v2597
		goto L519
	} else {
		goto L523
	}
L521:
	;
	goto L522
L522:
	;
	v2633 = F_pg_popcount_optimized(m, v2602, v2609)
	mBase = m.M
	v2636 = v2633
	goto L519
L523:
	;
	v2614 = v2609
	v2615 = v2602
	v2616 = v2597
	goto L524
L524:
	;
	v2617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2615)+3)))
	v2618 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v2617)+uint32(_c_F_gen_partprune_steps_internal[0]))))
	v2619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2615)+2)))
	v2620 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v2619)+uint32(_c_F_gen_partprune_steps_internal[0]))))
	v2621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2615)+1)))
	v2622 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v2621)+uint32(_c_F_gen_partprune_steps_internal[0]))))
	v2623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2615))))
	v2624 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v2623)+uint32(_c_F_gen_partprune_steps_internal[0]))))
	v2628 = v2618 + (v2620 + (v2622 + (v2616 + v2624)))
	v2629 = int32(4)
	v2632 = v2614 - v2629
	if v2632 != 0 {
		v2614 = v2632
		v2615 = v2615 + v2629
		v2616 = v2628
		goto L524
	} else {
		goto L526
	}
L525:
	;
	v2636 = v2628
	goto L519
L526:
	;
	goto L525
L527:
	;
	v2645 = F_palloc0(m, int32(24))
	mBase = m.M
	v2646 = m.ExcPending
	if v2646 != 0 {
		goto L4
	} else {
		goto L528
	}
L528:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2645))) = int32(381)
	v2649 = *(*int32)(unsafe.Add(mBase, uint32(v2574)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v2574)+16)) = v2649 + int32(1)
	v2653 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2645)+20)) = v2653
	*(*int64)(unsafe.Add(mBase, uint32(v2645)+12)) = int64(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v2645)+8)) = uint16(v2653)
	*(*int32)(unsafe.Add(mBase, uint32(v2645)+4)) = v2649
	v2660 = *(*int32)(unsafe.Add(mBase, uint32(v2574)+8))
	v2661 = F_lappend(m, v2660, v2645)
	mBase = m.M
	v2662 = m.ExcPending
	if v2662 != 0 {
		goto L4
	} else {
		goto L529
	}
L529:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2574)+8)) = v2661
	v2664 = F_lappend(m, v2585, v2645)
	mBase = m.M
	v2665 = m.ExcPending
	if v2665 != 0 {
		goto L4
	} else {
		goto L530
	}
L530:
	;
	v2666 = v2574
	v2669 = v2577
	v2677 = v2664
	goto L7
L531:
	;
	v2752 = v2669
	v2760 = int32(0)
	goto L1
L532:
	;
	goto L533
L533:
	;
	v2691 = *(*int32)(unsafe.Add(mBase, uint32(v2677)+4))
	if v2691 < int32(2) {
		v2752 = v2669
		v2760 = v2677
		goto L1
	} else {
		goto L534
	}
L534:
	;
	v2694 = int32(0)
	v2698 = v2694
	v2701 = v2694
	goto L535
L535:
	;
	v2718 = *(*int32)(unsafe.Add(mBase, uint32(v2677)+12))
	v2722 = *(*int32)(unsafe.Add(mBase, uint32(v2718+v2698<<(uint(int32(2))%32))))
	v2723 = *(*int32)(unsafe.Add(mBase, uint32(v2722)+4))
	v2724 = F_lappend_int(m, v2701, v2723)
	mBase = m.M
	v2725 = m.ExcPending
	if v2725 != 0 {
		goto L4
	} else {
		goto L537
	}
L536:
	;
	v2731 = F_palloc0(m, int32(16))
	mBase = m.M
	v2732 = m.ExcPending
	if v2732 != 0 {
		goto L4
	} else {
		goto L539
	}
L537:
	;
	v2727 = v2698 + int32(1)
	v2728 = *(*int32)(unsafe.Add(mBase, uint32(v2677)+4))
	if v2727 < v2728 {
		v2698 = v2727
		v2701 = v2724
		goto L535
	} else {
		goto L538
	}
L538:
	;
	goto L536
L539:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2731))) = int32(382)
	v2735 = *(*int32)(unsafe.Add(mBase, uint32(v2666)+16))
	v2736 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2666)+16)) = v2735 + v2736
	*(*int32)(unsafe.Add(mBase, uint32(v2731)+12)) = v2724
	*(*int32)(unsafe.Add(mBase, uint32(v2731)+8)) = v2736
	*(*int32)(unsafe.Add(mBase, uint32(v2731)+4)) = v2735
	v2743 = *(*int32)(unsafe.Add(mBase, uint32(v2666)+8))
	v2744 = F_lappend(m, v2743, v2731)
	mBase = m.M
	v2745 = m.ExcPending
	if v2745 != 0 {
		goto L4
	} else {
		goto L540
	}
L540:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2666)+8)) = v2744
	v2747 = F_lappend(m, v2677, v2731)
	mBase = m.M
	v2748 = m.ExcPending
	if v2748 != 0 {
		goto L4
	} else {
		goto L541
	}
L541:
	;
	v2752 = v2669
	v2760 = v2747
	goto L1
}
func F_generate_subscripts(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v75 int64
	_ = v75
	var v79 int32
	_ = v79
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int64
	_ = v97
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	if v9 == int32(0) {
		v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		v13 = F_DatumGetAnyArrayP(m, v12)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			v18 = F_init_MultiFuncCall(m, l0)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int64(0)
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				if v22 == int32(-1) {
					v25 = int32(28)
				} else {
					v25 = int32(4)
				}
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v13+v25)))
				if base.Ui32(v27-int32(7)) <= base.Ui32(int32(-7)) {
					F_end_MultiFuncCall(m, l0)
					mBase = m.M
					v128 = m.ExcPending
					if v128 != 0 {
						return int64(0)
					} else {
						v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v129)+20)) = int32(2)
						v132 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v132)
						return int64(0)
					}
				} else {
					v33 = int32(0)
					if base.B2i32(base.Ui32(v17) <= base.Ui32(v27))&base.B2i32(v33 < v17) == v33 {
						F_end_MultiFuncCall(m, l0)
						mBase = m.M
						v128 = m.ExcPending
						if v128 != 0 {
							return int64(0)
						} else {
							v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v129)+20)) = int32(2)
							v132 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v132)
							return int64(0)
						}
					} else {
						v38 = int32(_a_F_generate_subscripts_0)
						v39 = *(*int32)(unsafe.Add(mBase, _c_F_generate_subscripts[0]))
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
						*(*int32)(unsafe.Add(mBase, _c_F_generate_subscripts[0])) = v41
						v44 = F_palloc(m, int32(12))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int64(0)
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
							if v46 == int32(-1) {
								v49 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v13)+36))
								v57 = v49
								v58 = v50
							} else {
								v52 = v13 + int32(16)
								v53 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
								v57 = v52
								v58 = v52 + v53<<(uint(int32(2))%32)
							}
							v62 = v17<<(uint(int32(2))%32) - int32(4)
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v58+v62)))
							*(*int32)(unsafe.Add(mBase, uint32(v44))) = v64
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v62+v57)))
							*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v64 + v67 - int32(1)
							v72 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
							if int32(3) <= v72 {
								v75 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
								v79 = base.B2i32(v75 != int64(0))
							} else {
								v79 = int32(0)
							}
							*(*uint8)(unsafe.Add(mBase, uint32(v44)+8)) = uint8(v79)
							*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v44
							*(*int32)(unsafe.Add(mBase, _c_F_generate_subscripts[0])) = v39
							v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
							v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
							v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
							v94 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
							if v93 <= v94 {
								v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+8)))
								v97 = *(*int64)(unsafe.Add(mBase, uint32(v91)))
								*(*int64)(unsafe.Add(mBase, uint32(v91))) = v97 + int64(1)
								v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v101)+20)) = int32(1)
								if v96 == int32(0) {
									v106 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
									*(*int32)(unsafe.Add(mBase, uint32(v92))) = v106 + int32(1)
									return base.I64_extend_i32_s(v106)
								} else {
									v112 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v112 - int32(1)
									return base.I64_extend_i32_s(v112)
								}
							} else {
								F_end_MultiFuncCall(m, l0)
								mBase = m.M
								v119 = m.ExcPending
								if v119 != 0 {
									return int64(0)
								} else {
									v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v120)+20)) = int32(2)
									v123 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v123)
									return int64(0)
								}
							}
						}
					}
				}
			}
		}
	} else {
		v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
		v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
		v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
		v94 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
		if v93 <= v94 {
			v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+8)))
			v97 = *(*int64)(unsafe.Add(mBase, uint32(v91)))
			*(*int64)(unsafe.Add(mBase, uint32(v91))) = v97 + int64(1)
			v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v101)+20)) = int32(1)
			if v96 == int32(0) {
				v106 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
				*(*int32)(unsafe.Add(mBase, uint32(v92))) = v106 + int32(1)
				return base.I64_extend_i32_s(v106)
			} else {
				v112 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v112 - int32(1)
				return base.I64_extend_i32_s(v112)
			}
		} else {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v119 = m.ExcPending
			if v119 != 0 {
				return int64(0)
			} else {
				v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v120)+20)) = int32(2)
				v123 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v123)
				return int64(0)
			}
		}
	}
}
func F_generate_uuidv7(m *base.Module, l0 int64, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v1 int64
	_ = v1
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int64
	_ = v11
	var v14 int64
	_ = v14
	var v17 int64
	_ = v17
	var v20 int64
	_ = v20
	var v23 int64
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	v1 = l0
	v5 = F_palloc(m, int32(16))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+5)) = uint8(v1)
		v11 = int64(base.Ui64(v1) >> (uint(int64(8)) % 64))
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+4)) = uint8(v11)
		v14 = int64(base.Ui64(v1) >> (uint(int64(16)) % 64))
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+3)) = uint8(v14)
		v17 = int64(base.Ui64(v1) >> (uint(int64(24)) % 64))
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+2)) = uint8(v17)
		v20 = int64(base.Ui64(v1) >> (uint(int64(32)) % 64))
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+1)) = uint8(v20)
		v23 = int64(base.Ui64(v1) >> (uint(int64(40)) % 64))
		*(*uint8)(unsafe.Add(mBase, uint32(v5))) = uint8(v23)
		v28 = base.I32_div_u_s(l1<<(uint(int32(12))%32), int32(_a_F_generate_uuidv7_0))
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+7)) = uint8(v28)
		v30 = int32(8)
		v31 = int32(base.Ui32(v28) >> (uint(v30) % 32))
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+6)) = uint8(v31)
		v36 = m.Env.Pgmem_random_bytes(m, v5+v30, v30)
		mBase = m.M
		if v36 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(2600))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_generate_uuidv7_1), int32(0))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_generate_uuidv7_2), int32(651), int32(_a_F_generate_uuidv7_3))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+6)))
			v59 = v55&int32(15) | int32(112)
			*(*uint8)(unsafe.Add(mBase, uint32(v5)+6)) = uint8(v59)
			v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+8)))
			v65 = v61&int32(63) | int32(128)
			*(*uint8)(unsafe.Add(mBase, uint32(v5)+8)) = uint8(v65)
			return v5
		}
	}
}
func F_genericcostestimate(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v198 int32
	_ = v198
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v233 float64
	_ = v233
	var v238 float64
	_ = v238
	var v241 int32
	_ = v241
	var v249 int32
	_ = v249
	var v258 float64
	_ = v258
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 float64
	_ = v277
	var v278 int32
	_ = v278
	var v282 float64
	_ = v282
	var v284 float64
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v302 float64
	_ = v302
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 float64
	_ = v313
	var v314 int32
	_ = v314
	var v315 float64
	_ = v315
	var v318 int32
	_ = v318
	var v319 float64
	_ = v319
	var v323 float64
	_ = v323
	var v325 float64
	_ = v325
	var v327 float64
	_ = v327
	var v330 float64
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 float64
	_ = v335
	var v345 float64
	_ = v345
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v352 float64
	_ = v352
	var v355 float64
	_ = v355
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v365 float64
	_ = v365
	var v366 float64
	_ = v366
	var v367 float64
	_ = v367
	var v369 int32
	_ = v369
	var v372 float64
	_ = v372
	var v373 float64
	_ = v373
	var v377 float64
	_ = v377
	var v378 float64
	_ = v378
	var v382 float64
	_ = v382
	var v386 float64
	_ = v386
	var v391 float64
	_ = v391
	var v401 float64
	_ = v401
	var v404 float64
	_ = v404
	var v409 float64
	_ = v409
	var v410 float64
	_ = v410
	var v413 float64
	_ = v413
	var v416 float64
	_ = v416
	var v418 float64
	_ = v418
	var v419 int32
	_ = v419
	var v420 float64
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v426 float64
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 float64
	_ = v430
	var v436 float64
	_ = v436
	var v447 float64
	_ = v447
	v5 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(16)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v26 == v5 {
		v125 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v25)+88))
	if v140 != 0 {
		goto L15
	} else {
		goto L16
	}
L2:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v29 <= int32(0) {
		v125 = v5
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v37 = v29
	v38 = v5
	v39 = v5
	goto L4
L4:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v52+v39<<(uint(int32(2))%32))))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	if v57 == int32(0) {
		v101 = v37
		v102 = v38
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v125 = v102
	goto L1
L6:
	;
	v117 = v39 + int32(1)
	if v117 < v101 {
		v37 = v101
		v38 = v102
		v39 = v117
		goto L4
	} else {
		goto L14
	}
L7:
	;
	v60 = int32(0)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v61 <= v60 {
		v101 = v37
		v102 = v38
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v68 = v60
	v70 = v38
	goto L9
L9:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v84+v68<<(uint(int32(2))%32))))
	v89 = F_lappend(m, v70, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v101 = v95
	v102 = v89
	goto L6
L11:
	;
	return
L12:
	;
	v92 = v68 + int32(1)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v92 < v93 {
		v68 = v92
		v70 = v89
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	goto L5
L15:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)+4))
	if v141 <= int32(0) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v218 = v125
	goto L17
L17:
	;
	v233 = *(*float64)(unsafe.Add(mBase, uint32(l3)+56))
	if base.F64_lt(v233, float64(1)) == int32(0) {
		v302 = v233
		goto L32
	} else {
		goto L33
	}
L18:
	;
	v211 = F_list_concat(m, v198, v125)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L11
	} else {
		goto L31
	}
L19:
	;
	v198 = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v145 = int32(0)
	v151 = v145
	v154 = v145
	goto L22
L22:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v140)+12))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v167+v151<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v171
	v177 = F_list_make1_impl(m, int32(1), v23+int32(4))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L11
	} else {
		goto L24
	}
L23:
	;
	v198 = v186
	goto L18
L24:
	;
	v180 = F_predicate_implied_by(m, v177, v125, int32(0))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L11
	} else {
		goto L25
	}
L25:
	;
	if v180 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v184 = F_list_concat(m, v154, v177)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L11
	} else {
		goto L29
	}
L27:
	;
	v186 = v154
	goto L28
L28:
	;
	v188 = v151 + int32(1)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v140)+4))
	if v188 < v189 {
		v151 = v188
		v154 = v186
		goto L22
	} else {
		goto L30
	}
L29:
	;
	v186 = v184
	goto L28
L30:
	;
	goto L23
L31:
	;
	v218 = v211
	goto L17
L32:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v309)+76))
	v311 = int32(0)
	v313 = F_clauselist_selectivity(m, l0, v218, v310, v311, v311)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L11
	} else {
		goto L46
	}
L33:
	;
	v238 = float64(1)
	if v125 == int32(0) {
		v302 = v238
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	if v241 <= int32(0) {
		v302 = v238
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v249 = int32(0)
	v258 = v238
	goto L36
L36:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v265+v249<<(uint(int32(2))%32))))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v269)+4))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v270)))
	if v271 == int32(20) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v302 = v284
	goto L32
L38:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v270)+28))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)+12))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v275)+4))
	v277 = F_estimate_array_length(m, l0, v276)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L11
	} else {
		goto L41
	}
L39:
	;
	v284 = v258
	goto L40
L40:
	;
	v286 = v249 + int32(1)
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	if v286 < v287 {
		v249 = v286
		v258 = v284
		goto L36
	} else {
		goto L45
	}
L41:
	;
	if base.F64_gt(v277, float64(1)) != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v282 = base.F64_mul(v258, v277)
	goto L44
L43:
	;
	v282 = v258
	goto L44
L44:
	;
	v284 = v282
	goto L40
L45:
	;
	goto L37
L46:
	;
	v315 = *(*float64)(unsafe.Add(mBase, uint32(l3)+40))
	if base.F64_le(v315, float64(0)) != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v319 = *(*float64)(unsafe.Add(mBase, uint32(v318)+128))
	v323 = base.F64_nearest(base.F64_div(base.F64_mul(v313, v319), v302))
	goto L49
L48:
	;
	v323 = v315
	goto L49
L49:
	;
	v325 = *(*float64)(unsafe.Add(mBase, uint32(v25)+24))
	if base.F64_gt(v323, v325) != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v327 = v325
	goto L52
L51:
	;
	v327 = v323
	goto L52
L52:
	;
	if base.F64_lt(v327, float64(1)) != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v330 = float64(1)
	goto L55
L54:
	;
	v330 = v327
	goto L55
L55:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l3)+64))
	if base.Ui32(v332) <= base.Ui32(v333) {
		v345 = float64(1)
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	F_get_tablespace_page_costs(m, v346, v23+int32(8), int32(0))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L11
	} else {
		goto L59
	}
L57:
	;
	v335 = float64(1)
	if base.F64_gt(v325, v335) == int32(0) {
		v345 = v335
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v345 = base.F64_ceil(base.F64_div(base.F64_mul(v330, base.F64_convert_i32_u(v332-v333)), v325))
	goto L56
L59:
	;
	v352 = base.F64_mul(l2, v302)
	if base.F64_gt(v352, float64(1)) != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v418 = F_index_other_operands_eval_cost(m, l0, v125)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L11
	} else {
		goto L82
	}
L61:
	;
	v355 = base.F64_mul(v352, v345)
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v361 = int32(1)
	if base.Ui32(v356) <= base.Ui32(v361) {
		goto L65
	} else {
		goto L66
	}
L62:
	;
	goto L63
L63:
	;
	v413 = *(*float64)(unsafe.Add(mBase, uint32(v23)+8))
	v416 = base.F64_mul(v345, v413)
	goto L60
L64:
	;
	v410 = *(*float64)(unsafe.Add(mBase, uint32(v23)+8))
	v416 = base.F64_div(base.F64_mul(v409, v410), l2)
	goto L60
L65:
	;
	v364 = v361
	goto L67
L66:
	;
	v364 = v356
	goto L67
L67:
	;
	v365 = base.F64_convert_i32_u(v364)
	v366 = base.F64_add(v365, v365)
	v367 = float64(1)
	v369 = *(*int32)(unsafe.Add(mBase, _c_F_genericcostestimate[0]))
	v372 = *(*float64)(unsafe.Add(mBase, uint32(l0)+304))
	v373 = base.F64_add(base.F64_convert_i32_u(v356), v372)
	if base.F64_gt(v373, v367) != 0 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v409 = v404
	goto L64
L69:
	;
	v377 = v373
	goto L71
L70:
	;
	v377 = v367
	goto L71
L71:
	;
	v378 = base.F64_div(base.F64_mul(v365, base.F64_convert_i32_s(v369)), v377)
	if base.F64_le(v378, float64(1)) != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v382 = v367
	goto L74
L73:
	;
	v382 = base.F64_ceil(v378)
	goto L74
L74:
	;
	if base.F64_le(v365, v382) != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v386 = base.F64_div(base.F64_mul(v355, v366), base.F64_add(v366, v355))
	if base.F64_ge(v386, v365) != 0 {
		v404 = v365
		goto L68
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v391 = base.F64_div(base.F64_mul(v366, v382), base.F64_sub(v366, v382))
	if base.F64_ge(v391, v355) != 0 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v409 = base.F64_ceil(v386)
	goto L64
L79:
	;
	v401 = base.F64_div(base.F64_mul(v355, v366), base.F64_add(v366, v355))
	goto L81
L80:
	;
	v401 = base.F64_add(v382, base.F64_div(base.F64_mul(base.F64_sub(v365, v382), base.F64_sub(v355, v391)), v365))
	goto L81
L81:
	;
	v404 = base.F64_ceil(v401)
	goto L68
L82:
	;
	v420 = F_index_other_operands_eval_cost(m, l0, v139)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L11
	} else {
		goto L83
	}
L83:
	;
	if v125 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	v424 = v423
	goto L86
L85:
	;
	v424 = int32(0)
	goto L86
L86:
	;
	v426 = *(*float64)(unsafe.Add(mBase, _c_F_genericcostestimate[1]))
	if v139 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	v428 = v427
	goto L89
L88:
	;
	v428 = int32(0)
	goto L89
L89:
	;
	v430 = *(*float64)(unsafe.Add(mBase, _c_F_genericcostestimate[2]))
	*(*float64)(unsafe.Add(mBase, uint32(l3)+40)) = v330
	*(*float64)(unsafe.Add(mBase, uint32(l3)+32)) = v345
	*(*int64)(unsafe.Add(mBase, uint32(l3)+24)) = int64(0)
	*(*float64)(unsafe.Add(mBase, uint32(l3)+16)) = v313
	v436 = base.F64_add(v418, v420)
	*(*float64)(unsafe.Add(mBase, uint32(l3))) = v436
	*(*float64)(unsafe.Add(mBase, uint32(l3)+8)) = base.F64_add(base.F64_mul(base.F64_mul(v302, v330), base.F64_add(v430, base.F64_mul(v426, base.F64_convert_i32_s(v424+v428)))), base.F64_add(v416, v436))
	v447 = *(*float64)(unsafe.Add(mBase, uint32(v23)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l3)+56)) = v302
	*(*float64)(unsafe.Add(mBase, uint32(l3)+48)) = v447
	m.G0 = v23 + int32(16)
	return
}
func F_geqo_copy(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
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
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v106 int64
	_ = v106
	var v108 int64
	_ = v108
	v4 = int32(0)
	if l2 <= v4 {
	} else {
		v13 = l2 & int32(3)
		if base.Ui32(int32(4)) <= base.Ui32(l2) {
			v21 = v4
			v25 = v4
			for {
				v28 = v21 << (uint(int32(2)) % 32)
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v31+v28)))
				*(*int32)(unsafe.Add(mBase, uint32(v28+v29))) = v33
				v35 = int32(4)
				v36 = v28 | v35
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v39+v36)))
				*(*int32)(unsafe.Add(mBase, uint32(v36+v37))) = v41
				v44 = v28 | int32(8)
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v49 = *(*int32)(unsafe.Add(mBase, uint32(v47+v44)))
				*(*int32)(unsafe.Add(mBase, uint32(v44+v45))) = v49
				v52 = v28 | int32(12)
				v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v55+v52)))
				*(*int32)(unsafe.Add(mBase, uint32(v52+v53))) = v57
				v60 = v21 + v35
				v62 = v25 + v35
				if v62 != l2&int32(2147483644) {
					v21 = v60
					v25 = v62
					continue
				} else {
					break
				}
				break
			}
			if v13 == int32(0) {
			} else {
				v69 = v60
				v78 = v69
				v83 = v4
				for {
					v85 = v78 << (uint(int32(2)) % 32)
					v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v90 = *(*int32)(unsafe.Add(mBase, uint32(v88+v85)))
					*(*int32)(unsafe.Add(mBase, uint32(v85+v86))) = v90
					v92 = int32(1)
					v95 = v83 + v92
					if v95 != v13 {
						v78 = v78 + v92
						v83 = v95
						continue
					} else {
						break
					}
					break
				}
			}
		} else {
			v69 = v4
			v78 = v69
			v83 = v4
			for {
				v85 = v78 << (uint(int32(2)) % 32)
				v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v90 = *(*int32)(unsafe.Add(mBase, uint32(v88+v85)))
				*(*int32)(unsafe.Add(mBase, uint32(v85+v86))) = v90
				v92 = int32(1)
				v95 = v83 + v92
				if v95 != v13 {
					v78 = v78 + v92
					v83 = v95
					continue
				} else {
					break
				}
				break
			}
		}
	}
	v106 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v106
	v108 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v108
	return
}
func F_getQuadrant(m *base.Module, l0 int32, l1 int32) int32 {
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v30 int64
	_ = v30
	var v31 int32
	_ = v31
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int64
	_ = v44
	var v45 int32
	_ = v45
	var v50 int64
	_ = v50
	var v51 int32
	_ = v51
	var v57 int64
	_ = v57
	var v58 int32
	_ = v58
	var v63 int64
	_ = v63
	var v64 int32
	_ = v64
	var v69 int64
	_ = v69
	var v70 int32
	_ = v70
	var v77 int64
	_ = v77
	var v78 int32
	_ = v78
	var v83 int64
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	v7 = base.I64_extend_i32_u(l1)
	v8 = base.I64_extend_i32_u(l0)
	v9 = F_DirectFunctionCall2Coll(m, int32(256), int32(0), v7, v8)
	v12 = m.ExcPending
	if v12 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L36
	}
L2:
	;
	return v88
L3:
	;
	v37 = F_DirectFunctionCall2Coll(m, int32(260), int32(0), v7, v8)
	v38 = m.ExcPending
	if v38 != 0 {
		goto L4
	} else {
		goto L15
	}
L4:
	;
	return int32(0)
L5:
	;
	if v9 == int64(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v17 = F_DirectFunctionCall2Coll(m, int32(257), int32(0), v7, v8)
	v18 = m.ExcPending
	if v18 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v21 = int32(1)
	v24 = F_DirectFunctionCall2Coll(m, int32(258), int32(0), v7, v8)
	v25 = m.ExcPending
	if v25 != 0 {
		goto L4
	} else {
		goto L11
	}
L9:
	;
	if v17 == int64(0) {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	if v24 != int64(0) {
		v88 = v21
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v30 = F_DirectFunctionCall2Coll(m, int32(259), int32(0), v7, v8)
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	if v30 != int64(0) {
		v88 = v21
		goto L2
	} else {
		goto L14
	}
L14:
	;
	goto L3
L15:
	;
	if v37 != int64(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v41 = int32(2)
	v44 = F_DirectFunctionCall2Coll(m, int32(258), int32(0), v7, v8)
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v57 = F_DirectFunctionCall2Coll(m, int32(260), int32(0), v7, v8)
	v58 = m.ExcPending
	if v58 != 0 {
		goto L4
	} else {
		goto L24
	}
L19:
	;
	if v44 != int64(0) {
		v88 = v41
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v50 = F_DirectFunctionCall2Coll(m, int32(259), int32(0), v7, v8)
	v51 = m.ExcPending
	if v51 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	if v50 != int64(0) {
		v88 = v41
		goto L2
	} else {
		goto L22
	}
L22:
	;
	goto L18
L23:
	;
	v77 = F_DirectFunctionCall2Coll(m, int32(256), int32(0), v7, v8)
	v78 = m.ExcPending
	if v78 != 0 {
		goto L4
	} else {
		goto L32
	}
L24:
	;
	if v57 == int64(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v63 = F_DirectFunctionCall2Coll(m, int32(257), int32(0), v7, v8)
	v64 = m.ExcPending
	if v64 != 0 {
		goto L4
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v69 = F_DirectFunctionCall2Coll(m, int32(261), int32(0), v7, v8)
	v70 = m.ExcPending
	if v70 != 0 {
		goto L4
	} else {
		goto L30
	}
L28:
	;
	if v63 == int64(0) {
		goto L23
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	if v69 == int64(0) {
		goto L23
	} else {
		goto L31
	}
L31:
	;
	return int32(3)
L32:
	;
	if v77 == int64(0) {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v83 = F_DirectFunctionCall2Coll(m, int32(261), int32(0), v7, v8)
	v84 = m.ExcPending
	if v84 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	if v83 == int64(0) {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v88 = int32(4)
	goto L2
L36:
	;
	F_errmsg_internal(m, int32(_a_F_getQuadrant_0), int32(0))
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_getQuadrant_1), int32(79), int32(_a_F_getQuadrant_2))
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_actual_variable_endpoint(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int64
	_ = v118
	var v119 int64
	_ = v119
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	v10 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(384)
	m.G0 = v16
	*(*int32)(unsafe.Add(mBase, uint32(v16)+308)) = v10
	*(*int32)(unsafe.Add(mBase, uint32(v16)+312)) = int32(6)
	v22 = F_GlobalVisHorizonKindForRel(m, l0)
	mBase = m.M
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22<<(uint(int32(2))%32))+uint32(_c_F_get_actual_variable_endpoint[0])))
	goto L1
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+352)) = v25
	v29 = int32(0)
	v33 = F_index_beginscan(m, l0, l1, v16+int32(312), v29, int32(1), v29, v29)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	return int32(0)
L3:
	;
	v37 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+28)) = uint8(v37)
	v40 = int32(0)
	F_index_rescan(m, v33, l3, v37, v40, v40)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v44 = F_index_getnext_tid(m, v33, l2)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L2
	} else {
		goto L6
	}
L5:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v16)+308))
	if v171 != 0 {
		goto L42
	} else {
		goto L43
	}
L6:
	;
	if v44 == int32(0) {
		v169 = v10
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v52 = v44
	v60 = v10
	v61 = int32(-1)
	goto L10
L8:
	;
	v169 = int32(0)
	goto L5
L9:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v33)+44))
	if v90 != 0 {
		goto L26
	} else {
		goto L27
	}
L10:
	;
	v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52)+2)))
	v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52))))
	v66 = v62 | v63<<(uint(int32(16))%32)
	v69 = F_visibilitymap_get_status(m, l0, v66, v16+int32(308))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L2
	} else {
		goto L12
	}
L11:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	m.T0[v87].(func(*base.Module, int32))(m, l6)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L2
	} else {
		goto L24
	}
L12:
	;
	if v69&int32(1) != 0 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v73 = F_index_fetch_heap(m, v33, l6)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	if v73 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	if v66 != v61 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	goto L11
L18:
	;
	v79 = v60 + int32(1)
	if int32(100) < v79 {
		goto L8
	} else {
		goto L21
	}
L19:
	;
	v82 = v60
	v83 = v61
	goto L20
L20:
	;
	v84 = F_index_getnext_tid(m, v33, l2)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L2
	} else {
		goto L22
	}
L21:
	;
	v82 = v79
	v83 = v66
	goto L20
L22:
	;
	if v84 != 0 {
		v52 = v84
		v60 = v82
		v61 = v83
		goto L10
	} else {
		goto L23
	}
L23:
	;
	goto L8
L24:
	;
	goto L9
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L2
	} else {
		goto L39
	}
L26:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+72)))
	if v91 != 0 {
		goto L8
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L2
	} else {
		goto L36
	}
L29:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v33)+48))
	v95 = int32(16)
	v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v90)+6)))
	if int32(0) <= base.I32_extend16_s(v99) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v110 = int32(1)
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+16)))
	if v111 == v110 {
		goto L25
	} else {
		goto L34
	}
L31:
	;
	v103 = int32(8)
	goto L33
L32:
	;
	v103 = v95
	goto L33
L33:
	;
	F_index_deform_tuple_internal(m, v92, v16+int32(48), v16+v95, v90+v103, v90+int32(8), int32(base.Ui32(v99)>>(uint(int32(15))%32)))
	mBase = m.M
	goto L30
L34:
	;
	v114 = int32(_a_F_get_actual_variable_endpoint_0)
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_get_actual_variable_endpoint[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_get_actual_variable_endpoint[1])) = l7
	v118 = *(*int64)(unsafe.Add(mBase, uint32(v16)+48))
	v119 = F_datumCopy(m, v118, l5, l4)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l8))) = v119
	*(*int32)(unsafe.Add(mBase, _c_F_get_actual_variable_endpoint[1])) = v115
	v169 = v110
	goto L5
L36:
	;
	F_errmsg_internal(m, int32(_a_F_get_actual_variable_endpoint_1), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_get_actual_variable_endpoint_2), int32(_a_F_get_actual_variable_endpoint_3), int32(_a_F_get_actual_variable_endpoint_4))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L2
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L39:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v141 + int32(4)
	F_errmsg_internal(m, int32(_a_F_get_actual_variable_endpoint_5), v16)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L2
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_get_actual_variable_endpoint_2), int32(_a_F_get_actual_variable_endpoint_6), int32(_a_F_get_actual_variable_endpoint_4))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	F_ReleaseBuffer(m, v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L2
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	F_index_endscan(m, v33)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L2
	} else {
		goto L46
	}
L45:
	;
	goto L44
L46:
	;
	m.G0 = v16 + int32(384)
	return v169
}
func F_get_aggregate_argtypes(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v10 == int32(0) {
		v37 = int32(0)
		m.G0 = v8 + int32(16)
		return v37
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		if int32(100) <= v14 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50856197))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					v55 = int32(99)
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v55
					F_errmsg_plural(m, int32(_a_F_get_aggregate_argtypes_0), int32(_a_F_get_aggregate_argtypes_1), v55, v8)
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_get_aggregate_argtypes_2), int32(2122), int32(_a_F_get_aggregate_argtypes_3))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v17 = int32(0)
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
			if v18 <= v17 {
				v37 = v17
			} else {
				v21 = v17
				for {
					v27 = v21 << (uint(int32(2)) % 32)
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v29+v27)))
					*(*int32)(unsafe.Add(mBase, uint32(l1+v27))) = v31
					v34 = v21 + int32(1)
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
					if v34 < v35 {
						v21 = v34
						continue
					} else {
						break
					}
					break
				}
				v37 = v34
			}
			m.G0 = v8 + int32(16)
			return v37
		}
	}
}
func F_get_attnum(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	v3 = F_SearchSysCacheAttName(m, l0, l1)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			return int32(0)
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v3)+16))
			v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)))
			v14 = int32(*(*int16)(unsafe.Add(mBase, uint32(v11+v12)+74)))
			F_ReleaseCatCache(m, v3)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				return base.I32_extend16_s(v14)
			}
		}
	}
}
func F_get_attoptions(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int64
	_ = v29
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v14 = F_SearchSysCache2(m, int32(7), base.I64_extend_i32_u(l0), base.I64_extend_i32_s(l1))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		if v14 != 0 {
			v22 = F_SysCacheGetAttr(m, int32(6), v14, int32(23), v9+int32(15))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int64(0)
			} else {
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
				if v24 == int32(0) {
					v29 = F_datumCopy(m, v22, int32(0), int32(-1))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int64(0)
					} else {
						v31 = v29
						F_ReleaseCatCache(m, v14)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int64(0)
						} else {
							m.G0 = v9 + int32(16)
							return v31
						}
					}
				} else {
					v31 = int64(0)
					F_ReleaseCatCache(m, v14)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int64(0)
					} else {
						m.G0 = v9 + int32(16)
						return v31
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l0
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
				F_errmsg_internal(m, int32(_a_F_get_attoptions_0), v9)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_get_attoptions_1), int32(1207), int32(_a_F_get_attoptions_2))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int64(0)
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
func F_get_commutator(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14287(m, l0, int32(40))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_get_controlfile(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	v8 = m.G0
	v10 = v8 - int32(1040)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(_a_F_get_controlfile_0)
	v16 = v10 + int32(16)
	v19 = F_pg_snprintf(m, v16, int32(1024), int32(_a_F_get_controlfile_1), v10)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v23 = m.G0
		v25 = v23 + int32(-64)
		m.G0 = v25
		v28 = F_palloc(m, int32(312))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			v31 = F_OpenTransientFile(m, v16, int32(0))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				if v31 != int32(-1) {
					v35 = int32(312)
					v36 = F_read(m, v31, v28, v35)
					mBase = m.M
					if v36 != v35 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							if v36 < int32(0) {
								F_errcode_for_file_access(m)
								mBase = m.M
								v96 = m.ExcPending
								if v96 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = v16
									F_errmsg(m, int32(_a_F_get_controlfile_2), v23+int32(-32))
									mBase = m.M
									v102 = m.ExcPending
									if v102 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_get_controlfile_3), int32(109), int32(_a_F_get_controlfile_4))
										mBase = m.M
										v107 = m.ExcPending
										if v107 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								F_errcode(m, int32(16779816))
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v25)+56)) = int32(312)
									*(*int32)(unsafe.Add(mBase, uint32(v25)+52)) = v36
									*(*int32)(unsafe.Add(mBase, uint32(v25)+48)) = v16
									F_errmsg(m, int32(_a_F_get_controlfile_5), v23+int32(-16))
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_get_controlfile_3), int32(118), int32(_a_F_get_controlfile_4))
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return int32(0)
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
						v62 = F_CloseTransientFile(m, v31)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							if v62 != 0 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v111 = m.ExcPending
								if v111 != 0 {
									return int32(0)
								} else {
									F_errcode_for_file_access(m)
									mBase = m.M
									v113 = m.ExcPending
									if v113 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v16
										F_errmsg(m, int32(_a_F_get_controlfile_6), v23+int32(-48))
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_get_controlfile_3), int32(130), int32(_a_F_get_controlfile_4))
											mBase = m.M
											v124 = m.ExcPending
											if v124 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v64 = int32(-1)
								v66 = m.Env.Pgmem_crc32c(m, v64, v28, int32(308))
								mBase = m.M
								v67 = *(*int32)(unsafe.Add(mBase, uint32(v28)+308))
								*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(base.B2i32(v66^v67 == v64))
								v73 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
								if v73&int32(_a_F_get_controlfile_7) != 0 {
									v76 = int32(0)
								} else {
									v76 = v73
								}
								if v76 != 0 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v128 = m.ExcPending
									if v128 != 0 {
										return int32(0)
									} else {
										F_errmsg_internal(m, int32(_a_F_get_controlfile_8), int32(0))
										mBase = m.M
										v132 = m.ExcPending
										if v132 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_get_controlfile_3), int32(169), int32(_a_F_get_controlfile_4))
											mBase = m.M
											v137 = m.ExcPending
											if v137 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									m.G0 = v25 - int32(-64)
									m.G0 = v10 + int32(1040)
									return v28
								}
							}
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return int32(0)
					} else {
						F_errcode_for_file_access(m)
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v25))) = v16
							F_errmsg(m, int32(_a_F_get_controlfile_9), v25)
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_get_controlfile_3), int32(95), int32(_a_F_get_controlfile_4))
								mBase = m.M
								v94 = m.ExcPending
								if v94 != 0 {
									return int32(0)
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
func F_get_db_info(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32, l15 int32, l16 int32) int32 {
	mBase := m.M
	_ = mBase
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
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
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
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
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v146 int64
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int64
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v164 int64
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v178 int64
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v192 int64
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	v25 = m.G0
	v27 = v25 + int32(-64)
	m.G0 = v27
	v32 = F_table_open(m, int32(1262), int32(1))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v61 = v25 + int32(-56)
	F_ScanKeyInit(m, v61, int32(2), int32(3), int32(62), base.I64_extend_i32_u(l0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	F_relation_close(m, v32, int32(1))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L93
	}
L5:
	;
	goto L4
L6:
	;
	v68 = int32(1)
	v71 = F_systable_beginscan(m, v32, int32(2671), v68, int32(0), v68, v61)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v73 = F_systable_getnext(m, v71)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	if v73 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	F_systable_endscan(m, v71)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v73)+16))
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+22)))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v79+v80)))
	F_systable_endscan(m, v71)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	goto L5
L13:
	;
	F_LockSharedObject(m, int32(1262), v82, l1)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v90 = F_SearchSysCache1(m, int32(21), base.I64_extend_i32_u(v82))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if v90 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+22)))
	v94 = v92 + v93
	v96 = v94 + int32(4)
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	if base.B2i32(v99 == int32(0))|base.B2i32(v99 != v102) != 0 {
		v120 = v99
		v121 = v102
		goto L20
	} else {
		goto L21
	}
L17:
	;
	goto L18
L18:
	;
	F_UnlockSharedObject(m, int32(1262), v82, l1)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L92
	}
L19:
	;
	if v120-v121 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L20:
	;
	goto L19
L21:
	;
	v105 = l0
	v106 = v96
	goto L22
L22:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+1)))
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+1)))
	if v110 == int32(0) {
		v120 = v110
		v121 = v109
		goto L20
	} else {
		goto L24
	}
L23:
	;
	v120 = v110
	v121 = v109
	goto L20
L24:
	;
	v113 = int32(1)
	if v110 == v109 {
		v105 = v105 + v113
		v106 = v106 + v113
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v82
	if l3 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	F_ReleaseCatCache(m, v90)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L91
	}
L29:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v94)+68))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v126
	goto L31
L30:
	;
	goto L31
L31:
	;
	if l4 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v94)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v128
	goto L34
L33:
	;
	goto L34
L34:
	;
	if l5 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+77)))
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v130)
	goto L37
L36:
	;
	goto L37
L37:
	;
	if l7 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+79)))
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v132)
	goto L40
L39:
	;
	goto L40
L40:
	;
	if l6 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+78)))
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v134)
	goto L43
L42:
	;
	goto L43
L43:
	;
	if l8 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v94)+84))
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = v136
	goto L46
L45:
	;
	goto L46
L46:
	;
	if l9 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v94)+88))
	*(*int32)(unsafe.Add(mBase, uint32(l9))) = v138
	goto L49
L48:
	;
	goto L49
L49:
	;
	if l10 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v94)+92))
	*(*int32)(unsafe.Add(mBase, uint32(l10))) = v140
	goto L52
L51:
	;
	goto L52
L52:
	;
	if l15 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+76)))
	*(*uint8)(unsafe.Add(mBase, uint32(l15))) = uint8(v142)
	goto L55
L54:
	;
	goto L55
L55:
	;
	if l11 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v146 = F_SysCacheGetAttrNotNull(m, int32(21), v90, int32(13))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	if l12 != 0 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	v149 = F_text_to_cstring(m, base.I32_wrap_i64(v146))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l11))) = v149
	goto L58
L61:
	;
	v154 = F_SysCacheGetAttrNotNull(m, int32(21), v90, int32(14))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	if l13 != 0 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	v157 = F_text_to_cstring(m, base.I32_wrap_i64(v154))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l12))) = v157
	goto L63
L66:
	;
	v164 = F_SysCacheGetAttr(m, int32(21), v90, int32(15), v25+int32(-57))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	if l14 != 0 {
		goto L74
	} else {
		goto L75
	}
L69:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+7)))
	if v166 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v171 = int32(0)
	goto L72
L71:
	;
	v169 = F_text_to_cstring(m, base.I32_wrap_i64(v164))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L73
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l13))) = v171
	goto L68
L73:
	;
	v171 = v169
	goto L72
L74:
	;
	v178 = F_SysCacheGetAttr(m, int32(21), v90, int32(16), v25+int32(-57))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	if l16 != 0 {
		goto L82
	} else {
		goto L83
	}
L77:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+7)))
	if v180 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v185 = int32(0)
	goto L80
L79:
	;
	v183 = F_text_to_cstring(m, base.I32_wrap_i64(v178))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L81
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l14))) = v185
	goto L76
L81:
	;
	v185 = v183
	goto L80
L82:
	;
	v192 = F_SysCacheGetAttr(m, int32(21), v90, int32(17), v25+int32(-57))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	F_ReleaseCatCache(m, v90)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L90
	}
L85:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+7)))
	if v194 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v199 = int32(0)
	goto L88
L87:
	;
	v197 = F_text_to_cstring(m, base.I32_wrap_i64(v192))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L89
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l16))) = v199
	goto L84
L89:
	;
	v199 = v197
	goto L88
L90:
	;
	goto L5
L91:
	;
	goto L18
L92:
	;
	goto L3
L93:
	;
	m.G0 = v27 - int32(-64)
	return base.B2i32(v73 != int32(0))
}
func F_get_equality_op_for_ordering_op(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v15 = F_get_ordering_op_properties(m, l0, v7+int32(12), v7+int32(8), v7+int32(4))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		if v15 == int32(0) {
			v33 = int32(0)
			m.G0 = v7 + int32(16)
			return v33
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
			v24 = F_get_opfamily_member_for_cmptype(m, v21, v22, v22, int32(3))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				if l1 == int32(0) {
					v33 = v24
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
					*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(base.B2i32(v28 == int32(5)))
					v33 = v24
				}
				m.G0 = v7 + int32(16)
				return v33
			}
		}
	}
}
func F_get_generated_columns(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
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
	var v77 int32
	_ = v77
	v4 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	if v10 == v4 {
		v77 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v77
L2:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+18)))
	if v13 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	if l2 == int32(0) {
		v77 = v4
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	if v21 <= int32(0) {
		v77 = v4
		goto L1
	} else {
		goto L8
	}
L6:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+17)))
	if v18 != int32(1) {
		v77 = v4
		goto L1
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	v28 = int32(0)
	v29 = v21
	v32 = v4
	goto L9
L9:
	;
	v36 = v28 + int32(1)
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28*int32(100)+(v9+v29<<(uint(int32(3))%32)))+118)))
	v47 = int32(0)
	if base.B2i32(v41 != int32(118))&base.B2i32(l2&base.B2i32(v41 == int32(115)) == v47) == v47 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v77 = v68
	goto L1
L11:
	;
	v52 = F_build_generation_expression(m, l0, v36)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v67 = v29
	v68 = v32
	goto L13
L13:
	;
	if v36 < v67 {
		v28 = v36
		v29 = v67
		v32 = v68
		goto L9
	} else {
		goto L19
	}
L14:
	;
	return int32(0)
L15:
	;
	F_ChangeVarNodes(m, v52, int32(1), l1)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v60 = int32(0)
	v62 = F_makeTargetEntry(m, v52, base.I32_extend16_s(v36), v60, v60)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v64 = F_lappend(m, v32, v62)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v67 = v66
	v68 = v64
	goto L13
L19:
	;
	goto L10
}
func F_get_matching_partitions(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
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
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int64
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int64
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v165 int64
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
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int64
	_ = v180
	var v182 int32
	_ = v182
	var v184 int64
	_ = v184
	var v186 int64
	_ = v186
	var v192 int32
	_ = v192
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int64
	_ = v276
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int64
	_ = v295
	var v296 int32
	_ = v296
	var v297 int64
	_ = v297
	var v298 int32
	_ = v298
	var v299 int64
	_ = v299
	var v300 int32
	_ = v300
	var v301 int64
	_ = v301
	var v302 int32
	_ = v302
	var v303 int64
	_ = v303
	var v307 int64
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int64
	_ = v312
	var v315 int64
	_ = v315
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v381 int64
	_ = v381
	var v382 int32
	_ = v382
	var v383 int64
	_ = v383
	var v384 int64
	_ = v384
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v422 int32
	_ = v422
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int64
	_ = v449
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v644 int32
	_ = v644
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v698 int32
	_ = v698
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
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
	var v769 int32
	_ = v769
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v798 int32
	_ = v798
	var v801 int32
	_ = v801
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v822 int32
	_ = v822
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v863 int32
	_ = v863
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v895 int32
	_ = v895
	var v901 int32
	_ = v901
	var v906 int32
	_ = v906
	var v914 int32
	_ = v914
	var v918 int32
	_ = v918
	var v932 int32
	_ = v932
	var v935 int32
	_ = v935
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v948 int32
	_ = v948
	var v952 int32
	_ = v952
	var v954 int32
	_ = v954
	var v957 int32
	_ = v957
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v984 int32
	_ = v984
	var v989 int32
	_ = v989
	var v993 int32
	_ = v993
	var v999 int32
	_ = v999
	var v1004 int32
	_ = v1004
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1013 int32
	_ = v1013
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1025 int32
	_ = v1025
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1038 int32
	_ = v1038
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1058 int32
	_ = v1058
	var v1060 int32
	_ = v1060
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1071 int32
	_ = v1071
	var v1074 int32
	_ = v1074
	var v1087 int32
	_ = v1087
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1102 int32
	_ = v1102
	var v1105 int32
	_ = v1105
	var v1107 int32
	_ = v1107
	var v1110 int32
	_ = v1110
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1116 int32
	_ = v1116
	var v1128 int32
	_ = v1128
	var v1141 int32
	_ = v1141
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1148 int32
	_ = v1148
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1157 int32
	_ = v1157
	var v1160 int32
	_ = v1160
	var v1162 int32
	_ = v1162
	var v1165 int32
	_ = v1165
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1175 int32
	_ = v1175
	var v1192 int32
	_ = v1192
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1204 int32
	_ = v1204
	var v1208 int32
	_ = v1208
	var v1213 int32
	_ = v1213
	var v1238 int32
	_ = v1238
	var v1242 int32
	_ = v1242
	var v1247 int32
	_ = v1247
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1283 int32
	_ = v1283
	var v1286 int32
	_ = v1286
	var v1289 int32
	_ = v1289
	var v1293 int32
	_ = v1293
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1301 int32
	_ = v1301
	var v1308 int32
	_ = v1308
	var v1310 int32
	_ = v1310
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1332 int32
	_ = v1332
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1339 int32
	_ = v1339
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1361 int32
	_ = v1361
	var v1364 int32
	_ = v1364
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1379 int32
	_ = v1379
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1385 int32
	_ = v1385
	var v1389 int32
	_ = v1389
	var v1392 int32
	_ = v1392
	var v1394 int32
	_ = v1394
	var v1397 int32
	_ = v1397
	var v1404 int32
	_ = v1404
	var v1406 int32
	_ = v1406
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1428 int32
	_ = v1428
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1452 int32
	_ = v1452
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1489 int32
	_ = v1489
	v22 = m.G0
	v24 = v22 - int32(352)
	m.G0 = v24
	if l1 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v24 + int32(352)
	return v1489
L2:
	;
	v38 = v26 << (uint(int32(2)) % 32)
	v39 = F_palloc0(m, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L7
	} else {
		goto L9
	}
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v26 != 0 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v28 = int32(0)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v33 = F_bms_add_range(m, v28, v28, v30-int32(1))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L5
L7:
	;
	return int32(0)
L8:
	;
	v1489 = v33
	goto L1
L9:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v41 <= int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(v39+v38-int32(4))))
	v1273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1272)+4)))
	v1274 = int32(0)
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(v1272)))
	if v1275 == v1274 {
		goto L285
	} else {
		goto L286
	}
L11:
	;
	v60 = int32(0)
	goto L13
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1238 = m.ExcPending
	if v1238 != 0 {
		goto L7
	} else {
		goto L280
	}
L13:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v65+v60<<(uint(int32(2))%32))))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	switch v70 - int32(381) {
	case 0:
		goto L19
	case 1:
		goto L17
	default:
		goto L18
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1204 = m.ExcPending
	if v1204 != 0 {
		goto L7
	} else {
		goto L277
	}
L15:
	;
	goto L14
L16:
	;
	v1192 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v39+v1192<<(uint(int32(2))%32)))) = v1175
	v1198 = v60 + int32(1)
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v1198 < v1199 {
		v60 = v1198
		goto L13
	} else {
		goto L276
	}
L17:
	;
	v1020 = F_palloc0(m, int32(8))
	mBase = m.M
	v1021 = m.ExcPending
	if v1021 != 0 {
		goto L7
	} else {
		goto L242
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1008 = m.ExcPending
	if v1008 != 0 {
		goto L7
	} else {
		goto L239
	}
L19:
	;
	v73 = int32(0)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	if v75 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
	v77 = v76
	goto L22
L21:
	;
	v77 = v73
	goto L22
L22:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
	if v78 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	v80 = v79
	goto L25
L24:
	;
	v80 = v73
	goto L25
L25:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v81 <= int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v263 = v258 + v259*v242*int32(28)
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	switch v264 - int32(104) {
	case 0:
		goto L68
	default:
		goto L65
	case 4:
		goto L67
	case 10:
		goto L66
	}
L27:
	;
	v242 = v81
	v245 = int32(0)
	goto L26
L28:
	;
	goto L29
L29:
	;
	v85 = int32(0)
	v89 = v80
	v90 = v77
	v91 = v85
	v95 = v85
	goto L30
L30:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v69)+20))
	v109 = F_bms_is_member(m, v91, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L7
	} else {
		goto L33
	}
L31:
	;
	v242 = v235
	v245 = v229
	goto L26
L32:
	;
	v234 = v91 + int32(1)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v234 < v235 {
		v89 = v225
		v90 = v226
		v91 = v234
		v95 = v229
		goto L30
	} else {
		goto L63
	}
L33:
	;
	if v109 != 0 {
		v225 = v89
		v226 = v90
		v229 = v95
		goto L32
	} else {
		goto L34
	}
L34:
	;
	if v91 <= v95 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	if v90 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v112 != int32(114) {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v242 = v115
	v245 = v95
	goto L26
L38:
	;
	v225 = v89
	v226 = int32(0)
	v229 = v95
	goto L32
L39:
	;
	goto L40
L40:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v122 = v119*v120 + v91
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	if v124 == int32(7) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v170 = v167 + v122*int32(28)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+4))
	if v166 == v171 {
		goto L50
	} else {
		goto L51
	}
L42:
	;
	v156 = F_palloc(m, int32(8))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L7
	} else {
		goto L49
	}
L43:
	;
	v127 = *(*int64)(unsafe.Add(mBase, uint32(v123)+24))
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+32)))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+320)) = uint8(v128)
	if v128 != 0 {
		goto L42
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v130+v122<<(uint(int32(2))%32))))
	v135 = int32(_a_F_get_matching_partitions_0)
	v136 = *(*int32)(unsafe.Add(mBase, _c_F_get_matching_partitions[0]))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_get_matching_partitions[0])) = v139
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v134)+24))
	v144 = m.T0[v143].(func(*base.Module, int32, int32, int32) int64)(m, v134, v138, v24+int32(320))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L7
	} else {
		goto L47
	}
L46:
	;
	v165 = v127
	goto L41
L47:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_get_matching_partitions[0])) = v136
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+320)))
	if v148 != int32(1) {
		v165 = v144
		goto L41
	} else {
		goto L48
	}
L48:
	;
	goto L42
L49:
	;
	v158 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v156)+4)) = uint16(v158)
	*(*int32)(unsafe.Add(mBase, uint32(v156))) = v158
	v1175 = v156
	goto L16
L50:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v24-int32(-64)+v91<<(uint(int32(3))%32)))) = v165
	v202 = v89 + int32(4)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)+12))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	if base.Ui32(v202) < base.Ui32(v205+v206<<(uint(int32(2))%32)) {
		goto L57
	} else {
		goto L58
	}
L51:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v177 = v174 + v91*int32(28)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v177)+4))
	if v178 == v166 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v180 = *(*int64)(unsafe.Add(mBase, uint32(v177)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v170)+16)) = v180
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v177)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v170)+24)) = v182
	v184 = *(*int64)(unsafe.Add(mBase, uint32(v177)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v170)+8)) = v184
	v186 = *(*int64)(unsafe.Add(mBase, uint32(v177)))
	*(*int64)(unsafe.Add(mBase, uint32(v170))) = v186
	*(*int32)(unsafe.Add(mBase, uint32(v170)+20)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v170)+16)) = int32(0)
	goto L55
L53:
	;
	goto L54
L54:
	;
	F_fmgr_info_cxt(m, v166, v170, v173)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L7
	} else {
		goto L56
	}
L55:
	;
	goto L50
L56:
	;
	goto L50
L57:
	;
	v211 = v202
	goto L59
L58:
	;
	v211 = int32(0)
	goto L59
L59:
	;
	v213 = v90 + int32(4)
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)+12))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	if base.Ui32(v213) < base.Ui32(v216+v217<<(uint(int32(2))%32)) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v222 = v213
	goto L62
L61:
	;
	v222 = int32(0)
	goto L62
L62:
	;
	v225 = v211
	v226 = v222
	v229 = v95 + int32(1)
	goto L32
L63:
	;
	goto L31
L64:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L7
	} else {
		goto L236
	}
L65:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L7
	} else {
		goto L233
	}
L66:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v69)+20))
	v581 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+8)))
	v583 = F_palloc0(m, int32(8))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L7
	} else {
		goto L145
	}
L67:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v69)+20))
	v448 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+8)))
	v449 = *(*int64)(unsafe.Add(mBase, uint32(v24)+64))
	v451 = F_palloc0(m, int32(8))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L7
	} else {
		goto L101
	}
L68:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v69)+20))
	v269 = F_palloc0(m, int32(8))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L7
	} else {
		goto L69
	}
L69:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v271)+24))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v276 = int64(0)
	if v267 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L70:
	;
	v445 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v269)+4)) = uint16(v445)
	v1175 = v269
	goto L16
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v269))) = v422
	goto L70
L72:
	;
	if v274 == v320+v245 {
		goto L87
	} else {
		goto L88
	}
L73:
	;
	v320 = int32(0)
	goto L72
L74:
	;
	goto L75
L75:
	;
	v281 = v267 + int32(8)
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v267)+4))
	if v282 == int32(1) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v281)))
	v320 = base.I32_popcnt(v285)
	goto L72
L77:
	;
	goto L78
L78:
	;
	v288 = v282 << (uint(int32(2)) % 32)
	if v288 <= int32(7) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v320 = base.I32_wrap_i64(v315)
	goto L72
L80:
	;
	if v288 == int32(0) {
		v315 = v276
		goto L79
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v312 = F_pg_popcount_optimized(m, v281, v288)
	mBase = m.M
	v315 = v312
	goto L79
L83:
	;
	v293 = v288
	v294 = v281
	v295 = v276
	goto L84
L84:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294)+3)))
	v297 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v296)+uint32(_c_F_get_matching_partitions[1]))))
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294)+2)))
	v299 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v298)+uint32(_c_F_get_matching_partitions[1]))))
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294)+1)))
	v301 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v300)+uint32(_c_F_get_matching_partitions[1]))))
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294))))
	v303 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v302)+uint32(_c_F_get_matching_partitions[1]))))
	v307 = v297 + (v299 + (v301 + (v295 + v303)))
	v308 = int32(4)
	v311 = v293 - v308
	if v311 != 0 {
		v293 = v311
		v294 = v294 + v308
		v295 = v307
		goto L84
	} else {
		goto L86
	}
L85:
	;
	v315 = v307
	goto L79
L86:
	;
	goto L85
L87:
	;
	v323 = int32(0)
	if v323 < v274 {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	goto L89
L89:
	;
	v394 = int32(0)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v271)+20))
	v399 = F_bms_add_range(m, v394, v394, v396-int32(1))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L7
	} else {
		goto L100
	}
L90:
	;
	v329 = v323
	goto L93
L91:
	;
	goto L92
L92:
	;
	v381 = F_compute_partition_hash_value(m, v274, v263, v273, v24-int32(-64), v24+int32(320))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L7
	} else {
		goto L97
	}
L93:
	;
	v350 = F_bms_is_member(m, v329, v267)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L7
	} else {
		goto L95
	}
L94:
	;
	goto L92
L95:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v24+int32(320)+v329))) = uint8(v350)
	v354 = v329 + int32(1)
	if v354 != v274 {
		v329 = v354
		goto L93
	} else {
		goto L96
	}
L96:
	;
	goto L94
L97:
	;
	v383 = int64(*(*int32)(unsafe.Add(mBase, uint32(v271)+20)))
	v384 = base.I64_rem_u_s(v381, v383)
	v385 = base.I32_wrap_i64(v384)
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v272+v385<<(uint(int32(2))%32))))
	if v389 < int32(0) {
		goto L70
	} else {
		goto L98
	}
L98:
	;
	v392 = F_bms_make_singleton(m, v385)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L7
	} else {
		goto L99
	}
L99:
	;
	v422 = v392
	goto L71
L100:
	;
	v422 = v399
	goto L71
L101:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v455 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v451)+4)) = uint16(v455)
	if v447 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v453)+28))
	if v457 != int32(-1) {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	goto L104
L104:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v453)+4))
	if v466 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L105:
	;
	v460 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v451)+5)) = uint8(v460)
	v1175 = v451
	goto L16
L106:
	;
	goto L107
L107:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v453)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v451)+4)) = uint8(base.B2i32(v462 != int32(-1)))
	v1175 = v451
	goto L16
L108:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v453)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v451)+4)) = uint8(base.B2i32(v469 != int32(-1)))
	v1175 = v451
	goto L16
L109:
	;
	goto L110
L110:
	;
	v474 = v466 - int32(1)
	if v245 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v477 = int32(0)
	v479 = F_bms_add_range(m, v477, v477, v474)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L7
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	switch v448 {
	case 0:
		goto L122
	default:
		goto L121
	case 3:
		goto L120
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v451))) = v479
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v453)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v451)+4)) = uint8(base.B2i32(v482 != int32(-1)))
	v1175 = v451
	goto L16
L115:
	;
	v577 = F_bms_add_range(m, int32(0), v571, v573)
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L7
	} else {
		goto L144
	}
L116:
	;
	v560 = F_partition_list_bsearch(m, v263, v454, v453, v449, v24+int32(320))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L7
	} else {
		goto L139
	}
L117:
	;
	v556 = int32(0)
	goto L116
L118:
	;
	v540 = F_partition_list_bsearch(m, v263, v454, v453, v449, v24+int32(320))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L7
	} else {
		goto L134
	}
L119:
	;
	v537 = int32(1)
	goto L118
L120:
	;
	v520 = F_partition_list_bsearch(m, v263, v454, v453, v449, v24+int32(320))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L7
	} else {
		goto L130
	}
L121:
	;
	v510 = int32(-1)
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v453)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v451)+4)) = uint8(base.B2i32(v511 != v510))
	switch v448 - int32(1) {
	case 0:
		v556 = v510
		goto L116
	case 1:
		goto L117
	default:
		goto L64
	case 3:
		goto L119
	case 4:
		v537 = int32(0)
		goto L118
	}
L122:
	;
	v486 = int32(0)
	v488 = F_bms_add_range(m, v486, v486, v474)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L7
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v451))) = v488
	v493 = F_partition_list_bsearch(m, v263, v454, v453, v449, v24+int32(320))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L7
	} else {
		goto L125
	}
L124:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v453)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v451)+4)) = uint8(base.B2i32(v506 != int32(-1)))
	v1175 = v451
	goto L16
L125:
	;
	if v493 < int32(0) {
		goto L124
	} else {
		goto L126
	}
L126:
	;
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+320)))
	if v497&int32(1) == int32(0) {
		goto L124
	} else {
		goto L127
	}
L127:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v451)))
	v503 = F_bms_del_member(m, v502, v493)
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L7
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v451))) = v503
	goto L124
L129:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v453)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v451)+4)) = uint8(base.B2i32(v532 != int32(-1)))
	v1175 = v451
	goto L16
L130:
	;
	if v520 < int32(0) {
		goto L129
	} else {
		goto L131
	}
L131:
	;
	v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+320)))
	if v524&int32(1) == int32(0) {
		goto L129
	} else {
		goto L132
	}
L132:
	;
	v529 = F_bms_make_singleton(m, v520)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L7
	} else {
		goto L133
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v451))) = v529
	v1175 = v451
	goto L16
L134:
	;
	v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+320)))
	v544 = int32(0)
	if v544 <= v540 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v550 = v540 + base.B2i32(v537&v542 == v544)
	goto L137
L136:
	;
	v550 = v544
	goto L137
L137:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v453)+4))
	if v550 <= v551-int32(1) {
		v571 = v550
		v573 = v474
		goto L115
	} else {
		goto L138
	}
L138:
	;
	v1175 = v451
	goto L16
L139:
	;
	v562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+320)))
	v563 = int32(0)
	if v563 <= v560 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v566 = v562
	goto L142
L141:
	;
	v566 = v563
	goto L142
L142:
	;
	v568 = v560 - v566&v556
	if v568 < int32(0) {
		v1175 = v451
		goto L16
	} else {
		goto L143
	}
L143:
	;
	v571 = int32(0)
	v573 = v568
	goto L115
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v451))) = v577
	v1175 = v451
	goto L16
L145:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v586 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v587 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v587)+24))
	v589 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v583)+4)) = uint16(v589)
	if v580 == v589 {
		goto L147
	} else {
		goto L148
	}
L146:
	;
	if v245 == int32(0) {
		goto L151
	} else {
		goto L152
	}
L147:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v587)+4))
	if v593 != 0 {
		goto L146
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v587)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v583)+4)) = uint8(base.B2i32(v595 != int32(-1)))
	v1175 = v583
	goto L16
L150:
	;
	goto L149
L151:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v587)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v583)+4)) = uint8(base.B2i32(v601 != int32(-1)))
	v605 = int32(0)
	v607 = F_bms_add_range(m, v605, v605, v593)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L7
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	v610 = base.B2i32(v585 <= v245)
	if v610 == int32(0) {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v583))) = v607
	v1175 = v583
	goto L16
L155:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v587)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v583)+4)) = uint8(base.B2i32(v613 != int32(-1)))
	goto L157
L156:
	;
	goto L157
L157:
	;
	v617 = int32(1)
	v618 = int32(0)
	switch v581 - v617 {
	case 0:
		v822 = v617
		goto L161
	case 1:
		goto L162
	case 2:
		goto L165
	case 3:
		goto L164
	case 4:
		v760 = v618
		goto L163
	default:
		goto L160
	}
L158:
	;
	if v245 != int32(1) {
		v967 = v914
		v968 = v918
		goto L224
	} else {
		goto L225
	}
L159:
	;
	v914 = v593
	v918 = v765 + int32(1)
	goto L158
L160:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L7
	} else {
		goto L221
	}
L161:
	;
	v827 = F_partition_range_datum_bsearch(m, v263, v586, v587, v245, v24-int32(-64), v24+int32(320))
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L7
	} else {
		goto L204
	}
L162:
	;
	v822 = int32(0)
	goto L161
L163:
	;
	v765 = F_partition_range_datum_bsearch(m, v263, v586, v587, v245, v24-int32(-64), v24+int32(320))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L7
	} else {
		goto L190
	}
L164:
	;
	v760 = int32(1)
	goto L163
L165:
	;
	v626 = F_partition_range_datum_bsearch(m, v263, v586, v587, v245, v24-int32(-64), v24+int32(320))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L7
	} else {
		goto L167
	}
L166:
	;
	v756 = F_bms_make_singleton(m, v626+int32(1))
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L7
	} else {
		goto L189
	}
L167:
	;
	if v626 < int32(0) {
		goto L166
	} else {
		goto L168
	}
L168:
	;
	v630 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+320)))
	if v630&int32(1) == int32(0) {
		goto L166
	} else {
		goto L169
	}
L169:
	;
	if v245 == v585 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v638 = F_bms_make_singleton(m, v626+int32(1))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L7
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	v644 = v626
	goto L174
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v583))) = v638
	v1175 = v583
	goto L16
L174:
	;
	if v644 <= int32(0) {
		goto L177
	} else {
		goto L178
	}
L175:
	;
	v684 = int32(2)
	v685 = v245 << (uint(v684) % 32)
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v587)+12))
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v686+v682<<(uint(v684)%32))))
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v685+v690)))
	v698 = v626
	goto L183
L176:
	;
	goto L175
L177:
	;
	v682 = int32(0)
	goto L176
L178:
	;
	goto L179
L179:
	;
	v666 = v644 - int32(1)
	v668 = v666 << (uint(int32(2)) % 32)
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v587)+8))
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v668+v669)))
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v587)+12))
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v672+v668)))
	v677 = F_partition_rbound_datum_cmp(m, v263, v586, v671, v674, v24-int32(-64), v245)
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L7
	} else {
		goto L180
	}
L180:
	;
	if v677 == int32(0) {
		v644 = v666
		goto L174
	} else {
		goto L181
	}
L181:
	;
	v682 = v644
	goto L176
L182:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v740+v698<<(uint(int32(2))%32))))
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v745+v685)))
	v751 = F_bms_add_range(m, int32(0), v682+base.B2i32(v692 == int32(-1)), v698+base.B2i32(v747 != int32(1)))
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L7
	} else {
		goto L188
	}
L183:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v587)+12))
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v587)+4))
	if v718-int32(1) <= v698 {
		v740 = v717
		goto L182
	} else {
		goto L185
	}
L184:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v587)+12))
	v740 = v737
	goto L182
L185:
	;
	v723 = v698 + int32(1)
	v725 = v723 << (uint(int32(2)) % 32)
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v587)+8))
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v725+v726)))
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v717+v725)))
	v733 = F_partition_rbound_datum_cmp(m, v263, v586, v728, v730, v24-int32(-64), v245)
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L7
	} else {
		goto L186
	}
L186:
	;
	if v733 == int32(0) {
		v698 = v723
		goto L183
	} else {
		goto L187
	}
L187:
	;
	goto L184
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v583))) = v751
	v1175 = v583
	goto L16
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v583))) = v756
	v1175 = v583
	goto L16
L190:
	;
	if v765 < int32(0) {
		v914 = v593
		v918 = v618
		goto L158
	} else {
		goto L191
	}
L191:
	;
	if v585 <= v245 {
		goto L159
	} else {
		goto L192
	}
L192:
	;
	v769 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+320)))
	if v769&int32(1) == int32(0) {
		goto L159
	} else {
		goto L193
	}
L193:
	;
	if v760 != 0 {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v776 = int32(-1)
	goto L196
L195:
	;
	v776 = int32(1)
	goto L196
L196:
	;
	v779 = v765
	goto L197
L197:
	;
	v798 = v779 + v776
	if v798 < int32(0) {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	v914 = v593
	v918 = v779 + base.B2i32(v760 == int32(0))
	goto L158
L199:
	;
	goto L198
L200:
	;
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v587)+4))
	if v801 <= v798 {
		goto L199
	} else {
		goto L201
	}
L201:
	;
	v804 = v798 << (uint(int32(2)) % 32)
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v587)+8))
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v804+v805)))
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v587)+12))
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v808+v804)))
	v813 = F_partition_rbound_datum_cmp(m, v263, v586, v807, v810, v24-int32(-64), v245)
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L7
	} else {
		goto L202
	}
L202:
	;
	if v813 == int32(0) {
		v779 = v798
		goto L197
	} else {
		goto L203
	}
L203:
	;
	goto L199
L204:
	;
	if int32(0) <= v827 {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v831 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+320)))
	v834 = int32(0)
	if v610|base.B2i32(v831&int32(1) == v834) == v834 {
		goto L208
	} else {
		goto L209
	}
L206:
	;
	goto L207
L207:
	;
	v914 = v827 + int32(1)
	v918 = v618
	goto L158
L208:
	;
	if v822 != 0 {
		goto L211
	} else {
		goto L212
	}
L209:
	;
	goto L210
L210:
	;
	v914 = v827 + base.B2i32(v831&v822 == int32(0))
	v918 = v618
	goto L158
L211:
	;
	v841 = int32(-1)
	goto L213
L212:
	;
	v841 = int32(1)
	goto L213
L213:
	;
	v844 = v827
	goto L214
L214:
	;
	v863 = v844 + v841
	if v863 < int32(0) {
		goto L216
	} else {
		goto L217
	}
L215:
	;
	v914 = v844 + base.B2i32(v822 == int32(0))
	v918 = v618
	goto L158
L216:
	;
	goto L215
L217:
	;
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v587)+4))
	if v866 <= v863 {
		goto L216
	} else {
		goto L218
	}
L218:
	;
	v869 = v863 << (uint(int32(2)) % 32)
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v587)+8))
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v869+v870)))
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v587)+12))
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v873+v869)))
	v878 = F_partition_rbound_datum_cmp(m, v263, v586, v872, v875, v24-int32(-64), v245)
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L7
	} else {
		goto L219
	}
L219:
	;
	if v878 == int32(0) {
		v844 = v863
		goto L214
	} else {
		goto L220
	}
L220:
	;
	goto L216
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v581
	F_errmsg_internal(m, int32(_a_F_get_matching_partitions_1), v24+int32(48))
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L7
	} else {
		goto L222
	}
L222:
	;
	F_errfinish(m, int32(_a_F_get_matching_partitions_2), int32(3343), int32(_a_F_get_matching_partitions_3))
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L7
	} else {
		goto L223
	}
L223:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L224:
	;
	if v967 < v968 {
		v1175 = v583
		goto L16
	} else {
		goto L231
	}
L225:
	;
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v587)+4))
	if v932 <= v918 {
		v948 = v918
		goto L226
	} else {
		goto L227
	}
L226:
	;
	if v914 <= int32(0) {
		v967 = v914
		v968 = v948
		goto L224
	} else {
		goto L229
	}
L227:
	;
	v935 = v918 << (uint(int32(2)) % 32)
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v588+v935)))
	if int32(0) <= v937 {
		v948 = v918
		goto L226
	} else {
		goto L228
	}
L228:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v587)+12))
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v940+v935)))
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v942)))
	v948 = v918 + base.B2i32(v943 == int32(-1))
	goto L226
L229:
	;
	v952 = v914 << (uint(int32(2)) % 32)
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v588+v952)))
	if int32(0) <= v954 {
		v967 = v914
		v968 = v948
		goto L224
	} else {
		goto L230
	}
L230:
	;
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v587)+12))
	v961 = *(*int32)(unsafe.Add(mBase, uint32(v957+v952-int32(4))))
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v961)))
	v967 = v914 - base.B2i32(v962 == int32(1))
	v968 = v948
	goto L224
L231:
	;
	v971 = F_bms_add_range(m, int32(0), v968, v967)
	mBase = m.M
	v972 = m.ExcPending
	if v972 != 0 {
		goto L7
	} else {
		goto L232
	}
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v583))) = v971
	v1175 = v583
	goto L16
L233:
	;
	v978 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v978
	F_errmsg_internal(m, int32(_a_F_get_matching_partitions_4), v24+int32(16))
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L7
	} else {
		goto L234
	}
L234:
	;
	F_errfinish(m, int32(_a_F_get_matching_partitions_2), int32(3602), int32(_a_F_get_matching_partitions_5))
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L7
	} else {
		goto L235
	}
L235:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v448
	F_errmsg_internal(m, int32(_a_F_get_matching_partitions_1), v24+int32(32))
	mBase = m.M
	v999 = m.ExcPending
	if v999 != 0 {
		goto L7
	} else {
		goto L237
	}
L237:
	;
	F_errfinish(m, int32(_a_F_get_matching_partitions_2), int32(2956), int32(_a_F_get_matching_partitions_6))
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L7
	} else {
		goto L238
	}
L238:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L239:
	;
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v1009
	F_errmsg_internal(m, int32(_a_F_get_matching_partitions_7), v24)
	mBase = m.M
	v1013 = m.ExcPending
	if v1013 != 0 {
		goto L7
	} else {
		goto L240
	}
L240:
	;
	F_errfinish(m, int32(_a_F_get_matching_partitions_2), int32(893), int32(_a_F_get_matching_partitions_8))
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L7
	} else {
		goto L241
	}
L241:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L242:
	;
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	if v1022 == int32(0) {
		goto L243
	} else {
		goto L244
	}
L243:
	;
	v1025 = int32(0)
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(v1027)+20))
	v1031 = F_bms_add_range(m, v1025, v1025, v1028-int32(1))
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		goto L7
	} else {
		goto L246
	}
L244:
	;
	goto L245
L245:
	;
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
	switch v1042 {
	case 0:
		goto L247
	case 1:
		goto L248
	default:
		v1175 = v1020
		goto L16
	}
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1020))) = v1031
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v1027)+32))
	v1035 = int32(-1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1020)+4)) = uint8(base.B2i32(v1034 != v1035))
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v1027)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v1020)+5)) = uint8(base.B2i32(v1038 != v1035))
	v1175 = v1020
	goto L16
L247:
	;
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(v1022)+4))
	if v1116 <= int32(0) {
		v1175 = v1020
		goto L16
	} else {
		goto L264
	}
L248:
	;
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v1022)+4))
	if v1043 <= int32(0) {
		v1175 = v1020
		goto L16
	} else {
		goto L249
	}
L249:
	;
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v1022)+12))
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v1046)))
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if v1048 <= v1047 {
		goto L12
	} else {
		goto L250
	}
L250:
	;
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v39+v1047<<(uint(int32(2))%32))))
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(v1053)))
	v1055 = F_bms_copy(m, v1054)
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L7
	} else {
		goto L251
	}
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1020))) = v1055
	v1058 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1053)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1020)+5)) = uint8(v1058)
	v1060 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1053)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1020)+4)) = uint8(v1060)
	v1062 = int32(1)
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(v1022)+4))
	if v1063 <= v1062 {
		v1175 = v1020
		goto L16
	} else {
		goto L252
	}
L252:
	;
	v1071 = v1055
	v1074 = v1062
	goto L253
L253:
	;
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v1022)+12))
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v1087+v1074<<(uint(int32(2))%32))))
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if v1092 <= v1091 {
		goto L12
	} else {
		goto L255
	}
L254:
	;
	v1175 = v1020
	goto L16
L255:
	;
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v39+v1091<<(uint(int32(2))%32))))
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v1097)))
	v1099 = F_bms_int_members(m, v1071, v1098)
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		goto L7
	} else {
		goto L256
	}
L256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1020))) = v1099
	v1102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1020)+5)))
	if v1102 == int32(1) {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	v1105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1097)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1020)+5)) = uint8(v1105)
	goto L259
L258:
	;
	goto L259
L259:
	;
	v1107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1020)+4)))
	if v1107 == int32(1) {
		goto L260
	} else {
		goto L261
	}
L260:
	;
	v1110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1097)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1020)+4)) = uint8(v1110)
	goto L262
L261:
	;
	goto L262
L262:
	;
	v1113 = v1074 + int32(1)
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(v1022)+4))
	if v1113 < v1114 {
		v1071 = v1099
		v1074 = v1113
		goto L253
	} else {
		goto L263
	}
L263:
	;
	goto L254
L264:
	;
	v1128 = int32(0)
	goto L265
L265:
	;
	v1141 = *(*int32)(unsafe.Add(mBase, uint32(v1022)+12))
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v1141+v1128<<(uint(int32(2))%32))))
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if v1146 <= v1145 {
		goto L15
	} else {
		goto L267
	}
L266:
	;
	v1175 = v1020
	goto L16
L267:
	;
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(v1020)))
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v39+v1145<<(uint(int32(2))%32))))
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v1152)))
	v1154 = F_bms_add_members(m, v1148, v1153)
	mBase = m.M
	v1155 = m.ExcPending
	if v1155 != 0 {
		goto L7
	} else {
		goto L268
	}
L268:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1020))) = v1154
	v1157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1020)+5)))
	if v1157 == int32(0) {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	v1160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1152)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1020)+5)) = uint8(v1160)
	goto L271
L270:
	;
	goto L271
L271:
	;
	v1162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1020)+4)))
	if v1162 == int32(0) {
		goto L272
	} else {
		goto L273
	}
L272:
	;
	v1165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1152)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1020)+4)) = uint8(v1165)
	goto L274
L273:
	;
	goto L274
L274:
	;
	v1168 = v1128 + int32(1)
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(v1022)+4))
	if v1168 < v1169 {
		v1128 = v1168
		goto L265
	} else {
		goto L275
	}
L275:
	;
	goto L266
L276:
	;
	goto L10
L277:
	;
	F_errmsg_internal(m, int32(_a_F_get_matching_partitions_9), int32(0))
	mBase = m.M
	v1208 = m.ExcPending
	if v1208 != 0 {
		goto L7
	} else {
		goto L278
	}
L278:
	;
	F_errfinish(m, int32(_a_F_get_matching_partitions_2), int32(3657), int32(_a_F_get_matching_partitions_10))
	mBase = m.M
	v1213 = m.ExcPending
	if v1213 != 0 {
		goto L7
	} else {
		goto L279
	}
L279:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L280:
	;
	F_errmsg_internal(m, int32(_a_F_get_matching_partitions_9), int32(0))
	mBase = m.M
	v1242 = m.ExcPending
	if v1242 != 0 {
		goto L7
	} else {
		goto L281
	}
L281:
	;
	F_errfinish(m, int32(_a_F_get_matching_partitions_2), int32(3681), int32(_a_F_get_matching_partitions_10))
	mBase = m.M
	v1247 = m.ExcPending
	if v1247 != 0 {
		goto L7
	} else {
		goto L282
	}
L282:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L283:
	;
	if int32(0) <= v1332 {
		goto L294
	} else {
		goto L295
	}
L284:
	;
	v1332 = base.I32_ctz(v1318) | v1319<<(uint(int32(5))%32)
	goto L283
L285:
	;
	v1332 = int32(-2)
	goto L283
L286:
	;
	v1283 = int32(0)
	v1286 = *(*int32)(unsafe.Add(mBase, uint32(v1275)+4))
	if v1286 <= v1283 {
		goto L285
	} else {
		goto L287
	}
L287:
	;
	v1289 = v1275 + int32(8)
	v1293 = *(*int32)(unsafe.Add(mBase, uint32(v1289)))
	v1296 = v1293 & int32(-1)
	if v1296 != 0 {
		v1318 = v1296
		v1319 = v1283
		goto L284
	} else {
		goto L288
	}
L288:
	;
	v1297 = int32(1)
	if v1297 == v1286 {
		goto L285
	} else {
		goto L289
	}
L289:
	;
	v1301 = v1297
	goto L290
L290:
	;
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(v1289+v1301<<(uint(int32(2))%32))))
	if v1308 != 0 {
		v1318 = v1308
		v1319 = v1301
		goto L284
	} else {
		goto L292
	}
L291:
	;
	goto L285
L292:
	;
	v1310 = v1301 + int32(1)
	if v1310 != v1286 {
		v1301 = v1310
		goto L290
	} else {
		goto L293
	}
L293:
	;
	goto L291
L294:
	;
	v1336 = v1274
	v1337 = v1273
	v1339 = v1332
	goto L297
L295:
	;
	v1432 = v1274
	v1433 = v1273
	goto L296
L296:
	;
	v1452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1272)+5)))
	if v1452 == int32(1) {
		goto L316
	} else {
		goto L317
	}
L297:
	;
	v1356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1357 = *(*int32)(unsafe.Add(mBase, uint32(v1356)+24))
	v1361 = *(*int32)(unsafe.Add(mBase, uint32(v1357+v1339<<(uint(int32(2))%32))))
	if v1361 < int32(0) {
		goto L300
	} else {
		goto L301
	}
L298:
	;
	v1432 = v1370
	v1433 = v1371
	goto L296
L299:
	;
	v1372 = *(*int32)(unsafe.Add(mBase, uint32(v1272)))
	if v1372 == int32(0) {
		goto L306
	} else {
		goto L307
	}
L300:
	;
	v1364 = *(*int32)(unsafe.Add(mBase, uint32(v1356)+32))
	v1370 = v1336
	v1371 = v1337 | base.B2i32(v1364 != int32(-1))
	goto L299
L301:
	;
	goto L302
L302:
	;
	v1368 = F_bms_add_member(m, v1336, v1361)
	mBase = m.M
	v1369 = m.ExcPending
	if v1369 != 0 {
		goto L7
	} else {
		goto L303
	}
L303:
	;
	v1370 = v1368
	v1371 = v1337
	goto L299
L304:
	;
	if int32(0) <= v1428 {
		v1336 = v1370
		v1337 = v1371
		v1339 = v1428
		goto L297
	} else {
		goto L315
	}
L305:
	;
	v1428 = base.I32_ctz(v1414) | v1415<<(uint(int32(5))%32)
	goto L304
L306:
	;
	v1428 = int32(-2)
	goto L304
L307:
	;
	v1379 = v1339 + int32(1)
	v1381 = int32(base.Ui32(v1379) >> (uint(int32(5)) % 32))
	v1382 = *(*int32)(unsafe.Add(mBase, uint32(v1372)+4))
	if v1382 <= v1381 {
		goto L306
	} else {
		goto L308
	}
L308:
	;
	v1385 = v1372 + int32(8)
	v1389 = *(*int32)(unsafe.Add(mBase, uint32(v1385+v1381<<(uint(int32(2))%32))))
	v1392 = v1389 & (int32(-1) << (uint(v1379) % 32))
	if v1392 != 0 {
		v1414 = v1392
		v1415 = v1381
		goto L305
	} else {
		goto L309
	}
L309:
	;
	v1394 = v1381 + int32(1)
	if v1394 == v1382 {
		goto L306
	} else {
		goto L310
	}
L310:
	;
	v1397 = v1394
	goto L311
L311:
	;
	v1404 = *(*int32)(unsafe.Add(mBase, uint32(v1385+v1397<<(uint(int32(2))%32))))
	if v1404 != 0 {
		v1414 = v1404
		v1415 = v1397
		goto L305
	} else {
		goto L313
	}
L312:
	;
	goto L306
L313:
	;
	v1406 = v1397 + int32(1)
	if v1406 != v1382 {
		v1397 = v1406
		goto L311
	} else {
		goto L314
	}
L314:
	;
	goto L312
L315:
	;
	goto L298
L316:
	;
	v1455 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1456 = *(*int32)(unsafe.Add(mBase, uint32(v1455)+28))
	v1457 = F_bms_add_member(m, v1432, v1456)
	mBase = m.M
	v1458 = m.ExcPending
	if v1458 != 0 {
		goto L7
	} else {
		goto L319
	}
L317:
	;
	v1459 = v1432
	goto L318
L318:
	;
	if v1433&int32(1) == int32(0) {
		v1489 = v1459
		goto L1
	} else {
		goto L320
	}
L319:
	;
	v1459 = v1457
	goto L318
L320:
	;
	v1464 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1465 = *(*int32)(unsafe.Add(mBase, uint32(v1464)+32))
	v1466 = F_bms_add_member(m, v1459, v1465)
	mBase = m.M
	v1467 = m.ExcPending
	if v1467 != 0 {
		goto L7
	} else {
		goto L321
	}
L321:
	;
	v1489 = v1466
	goto L1
}
func F_get_notnull_info(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int64
	_ = v20
	var v25 int64
	_ = v25
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v52 int64
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int64
	_ = v56
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v86 int32
	_ = v86
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if v7 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	if v11 <= l1 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v86 = int32(0)
	goto L3
L3:
	;
	return v86
L4:
	;
	v13 = int32(_a_F_get_notnull_info_0)
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_get_notnull_info[0]))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+64))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_get_notnull_info[0])) = v18
	v20 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	v25 = v20
	goto L7
L5:
	;
	goto L6
L6:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66+base.I32_wrap_i64(int64(base.Ui64(l1)>>(uint(int64(2))%64)))))))
	v86 = int32(base.Ui32(v71)>>(uint(base.I32_wrap_i64(l1)<<(uint(int32(1))%32)&int32(6))%32)) & int32(3)
	goto L3
L7:
	;
	v27 = base.I64_div_s(v25, int64(4))
	v28 = base.I32_wrap_i64(v27)
	if v28 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_get_notnull_info[0])) = v14
	goto L6
L9:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int64)(unsafe.Add(mBase, uint32(v53))) = v52
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v55)))
	if v56 <= l1 {
		v25 = v56
		goto L7
	} else {
		goto L16
	}
L10:
	;
	v32 = F_palloc0(m, int32(32))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v43 = F_repalloc0(m, v40, v28, v28<<(uint(int32(1))%32))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L13
	} else {
		goto L15
	}
L13:
	;
	return int32(0)
L14:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v32
	v52 = int64(128)
	goto L9
L15:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = v43
	v52 = v27 << (uint(int64(3)) % 64) & int64(4294967288)
	goto L9
L16:
	;
	goto L8
}
func F_get_oprjoin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v5 = F_SearchSysCache1(m, int32(40), base.I64_extend_i32_u(l0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 == int32(0) {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+22)))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v13+v14)+108))
			F_ReleaseCatCache(m, v5)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				return v16
			}
		}
	}
}
func F_get_other_operator(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
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
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int64
	_ = v138
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v164 int64
	_ = v164
	var v166 int64
	_ = v166
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	v11 = m.G0
	v13 = v11 - int32(240)
	m.G0 = v13
	v16 = F_LookupOperName(m, l0, l1, l2, int32(1))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L3
	} else {
		goto L56
	}
L2:
	;
	m.G0 = v13 + int32(240)
	return v219
L3:
	;
	return int32(0)
L4:
	;
	if v16 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v20 = F_get_opcode(m, v16)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v24 = F_QualifiedNameGetCreationNamespace(m, l0, v13+int32(16))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L3
	} else {
		goto L9
	}
L8:
	;
	v219 = v16
	goto L2
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if base.B2i32(v29 == int32(0))|base.B2i32(v29 != v32) != 0 {
		v50 = v29
		v51 = v32
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if base.B2i32(v50-v51|(base.B2i32(l2 != l6)|base.B2i32(l1 != l5)) == int32(0))&base.B2i32(v24 == l4) != 0 {
		v219 = int32(0)
		goto L2
	} else {
		goto L17
	}
L11:
	;
	goto L10
L12:
	;
	v35 = v26
	v36 = l3
	goto L13
L13:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+1)))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+1)))
	if v40 == int32(0) {
		v50 = v40
		v51 = v39
		goto L11
	} else {
		goto L15
	}
L14:
	;
	v50 = v40
	v51 = v39
	goto L11
L15:
	;
	v43 = int32(1)
	if v40 == v39 {
		v35 = v35 + v43
		v36 = v36 + v43
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_get_other_operator[0]))
	v65 = F_object_aclcheck(m, int32(2615), v24, v63, int64(512))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L3
	} else {
		goto L18
	}
L18:
	;
	if v65 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v68 = F_get_namespace_name(m, v24)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L3
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	v73 = int32(0)
	v76 = F_strlen(m, v72)
	mBase = m.M
	if base.Ui32(v76+int32(-64)) < base.Ui32(int32(-63)) {
		v129 = v73
		goto L25
	} else {
		goto L26
	}
L22:
	;
	F_aclcheck_error(m, v65, int32(37), v68)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	if v129 == int32(0) {
		goto L1
	} else {
		goto L39
	}
L25:
	;
	goto L24
L26:
	;
	v82 = F_strspn(m, v72, int32(_a_F_get_other_operator_0))
	mBase = m.M
	if v82 != v76 {
		v129 = v73
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v85 = F_strstr(m, v72, int32(_a_F_get_other_operator_1))
	mBase = m.M
	if v85 != 0 {
		v129 = v73
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v87 = F_strstr(m, v72, int32(_a_F_get_other_operator_2))
	mBase = m.M
	if v87 != 0 {
		v129 = v73
		goto L25
	} else {
		goto L29
	}
L29:
	;
	if base.Ui32(v76) < base.Ui32(int32(2)) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v117 = int32(1)
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if v118 != int32(33) {
		v129 = v117
		goto L25
	} else {
		goto L37
	}
L31:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72+v76-int32(1)))))
	switch v93 - int32(43) {
	case 0, 2:
		goto L32
	default:
		goto L30
	}
L32:
	;
	v99 = v76 - int32(2)
	goto L33
L33:
	;
	v104 = int32(*(*int8)(unsafe.Add(mBase, uint32(v72+v99))))
	v106 = F_memchr(m, int32(_a_F_get_other_operator_3), v104, int32(11))
	mBase = m.M
	if v106 != 0 {
		goto L30
	} else {
		goto L35
	}
L34:
	;
	v129 = v73
	goto L25
L35:
	;
	v107 = int32(0)
	if base.B2i32(v99 <= v107) == v107 {
		v99 = v99 - int32(1)
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+1)))
	if v121 != int32(61) {
		v129 = v117
		goto L25
	} else {
		goto L38
	}
L38:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+2)))
	v129 = base.B2i32(v124 != int32(0))
	goto L25
L39:
	;
	v135 = F_table_open(m, int32(2617), int32(3))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L3
	} else {
		goto L40
	}
L40:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v135)+52))
	v138 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+103)) = v138
	*(*int64)(unsafe.Add(mBase, uint32(v13)+96)) = v138
	v144 = F_GetNewOidWithIndex(m, v135, int32(2688), int32(1))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+112)) = base.I64_extend_i32_u(v144)
	v149 = v13 + int32(32)
	v151 = F_strncpy(m, v149, v72, int32(64))
	mBase = m.M
	v152 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v151)+63)) = uint8(v152)
	goto L42
L42:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+128)) = base.I64_extend_i32_u(v24)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+120)) = base.I64_extend_i32_u(v149)
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_get_other_operator[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+152)) = int64(0)
	if l1 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v164 = int64(98)
	goto L45
L44:
	;
	v164 = int64(108)
	goto L45
L45:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+144)) = v164
	v166 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+160)) = v166
	*(*int64)(unsafe.Add(mBase, uint32(v13)+184)) = v166
	*(*int64)(unsafe.Add(mBase, uint32(v13)+176)) = base.I64_extend_i32_u(l2)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+168)) = base.I64_extend_i32_u(l1)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+192)) = v166
	*(*int64)(unsafe.Add(mBase, uint32(v13)+200)) = v166
	*(*int64)(unsafe.Add(mBase, uint32(v13)+208)) = v166
	*(*int64)(unsafe.Add(mBase, uint32(v13)+216)) = v166
	*(*int64)(unsafe.Add(mBase, uint32(v13)+224)) = v166
	*(*int64)(unsafe.Add(mBase, uint32(v13)+136)) = base.I64_extend_i32_u(v159)
	v190 = F_heap_form_tuple(m, v137, v13+int32(112), v13+int32(96))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L3
	} else {
		goto L46
	}
L46:
	;
	F_CatalogTupleInsert(m, v135, v190)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L3
	} else {
		goto L47
	}
L47:
	;
	F_makeOperatorDependencies(m, v13+int32(20), v190, int32(1), int32(0))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L3
	} else {
		goto L48
	}
L48:
	;
	F_pfree(m, v190)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L3
	} else {
		goto L49
	}
L49:
	;
	v203 = *(*int32)(unsafe.Add(mBase, _c_F_get_other_operator[1]))
	if v203 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v205 = int32(0)
	F_RunObjectPostCreateHook(m, int32(2617), v144, v205, v205)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L3
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L3
	} else {
		goto L54
	}
L53:
	;
	goto L52
L54:
	;
	F_relation_close(m, v135, int32(3))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L3
	} else {
		goto L55
	}
L55:
	;
	v219 = v144
	goto L2
L56:
	;
	F_errcode(m, int32(33579140))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L3
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v72
	F_errmsg(m, int32(_a_F_get_other_operator_4), v13)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L3
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_get_other_operator_5), int32(214), int32(_a_F_get_other_operator_6))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L3
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_parse_rowmark(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v39 int32
	_ = v39
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v6 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v39
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	if v7 <= int32(0) {
		v39 = int32(0)
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v39 = int32(0)
	goto L1
L5:
	;
	v10 = int32(0)
	if v10 < v7 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v13 = v7
	goto L8
L7:
	;
	v13 = v10
	goto L8
L8:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	v16 = int32(0)
	goto L9
L9:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v14+v16<<(uint(int32(2))%32))))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v25 == l1 {
		v39 = v24
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L4
L11:
	;
	v28 = v16 + int32(1)
	if v28 != v13 {
		v16 = v28
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
}
func F_get_pkglib_path(m *base.Module, l0 int32) {
	var v5 int32
	_ = v5
	F_make_relative_path(m, l0, int32(_a_F_get_pkglib_path_0), int32(_a_F_get_pkglib_path_1))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_get_plan_rowmark(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v38 int32
	_ = v38
	if l0 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v38
L2:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v6 <= int32(0) {
		v38 = int32(0)
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v38 = int32(0)
	goto L1
L5:
	;
	v9 = int32(0)
	if v9 < v6 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v12 = v6
	goto L8
L7:
	;
	v12 = v9
	goto L8
L8:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v15 = int32(0)
	goto L9
L9:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v13+v15<<(uint(int32(2))%32))))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v24 == l1 {
		v38 = v23
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L4
L11:
	;
	v27 = v15 + int32(1)
	if v27 != v12 {
		v15 = v27
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
}
func F_get_relkind_objtype(m *base.Module, l0 int32) int32 {
	var v15 int32
	_ = v15
	switch l0 - int32(73) {
	case 0, 32:
		v15 = int32(20)
		return v15
	default:
		v15 = int32(42)
		return v15
	case 10:
		return int32(38)
	case 29:
		return int32(18)
	case 36:
		return int32(23)
	case 45:
		return int32(52)
	}
}
func F_get_restriction_variable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int64
	_ = v53
	var v55 int64
	_ = v55
	var v57 int64
	_ = v57
	var v59 int64
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	v7 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	if l1 == v7 {
		v75 = v7
		m.G0 = v12 + int32(32)
		return v75
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		if v16 != int32(2) {
			v75 = v7
			m.G0 = v12 + int32(32)
			return v75
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			F_examine_variable(m, l0, v21, l2, l3)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				F_examine_variable(m, l0, v20, l2, v12)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
					v30 = int32(0)
					if v28|base.B2i32(v29 == v30) == v30 {
						v35 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v35)
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
						v39 = F_estimate_expression_value(m, l0, v38)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = v39
							v75 = v35
							m.G0 = v12 + int32(32)
							return v75
						}
					} else {
						v42 = int32(0)
						if v29|base.B2i32(v28 == v42) == v42 {
							v47 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v47)
							v49 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
							v50 = F_estimate_expression_value(m, l0, v49)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l4))) = v50
								v53 = *(*int64)(unsafe.Add(mBase, uint32(v12)+24))
								*(*int64)(unsafe.Add(mBase, uint32(l3)+24)) = v53
								v55 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
								*(*int64)(unsafe.Add(mBase, uint32(l3)+16)) = v55
								v57 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
								*(*int64)(unsafe.Add(mBase, uint32(l3)+8)) = v57
								v59 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
								*(*int64)(unsafe.Add(mBase, uint32(l3))) = v59
								v75 = int32(1)
								m.G0 = v12 + int32(32)
								return v75
							}
						} else {
							v62 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
							if v62 != 0 {
								v63 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
								m.T0[v63].(func(*base.Module, int32))(m, v62)
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int32(0)
								} else {
									v66 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
									if v66 == int32(0) {
										v75 = v7
										m.G0 = v12 + int32(32)
										return v75
									} else {
										v69 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
										m.T0[v69].(func(*base.Module, int32))(m, v66)
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return int32(0)
										} else {
											v75 = v7
											m.G0 = v12 + int32(32)
											return v75
										}
									}
								}
							} else {
								v66 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
								if v66 == int32(0) {
									v75 = v7
									m.G0 = v12 + int32(32)
									return v75
								} else {
									v69 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
									m.T0[v69].(func(*base.Module, int32))(m, v66)
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int32(0)
									} else {
										v75 = v7
										m.G0 = v12 + int32(32)
										return v75
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
func F_get_typ_typrelid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v5 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(l0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 == int32(0) {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+22)))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v13+v14)+84))
			F_ReleaseCatCache(m, v5)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				return v16
			}
		}
	}
}
func F_get_typbyval(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v5 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(l0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 != 0 {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
			v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
			v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+v10)+78)))
			F_ReleaseCatCache(m, v5)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				v15 = v12
				return v15 & int32(1)
			}
		} else {
			v15 = int32(0)
			return v15 & int32(1)
		}
	}
}
func F_get_typlenbyvalalign(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
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
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v13 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(l0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		if v13 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
				F_errmsg_internal(m, int32(_a_F_get_typlenbyvalalign_0), v9)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_get_typlenbyvalalign_1), int32(2591), int32(_a_F_get_typlenbyvalalign_2))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+22)))
			v32 = v30 + v31
			v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+76)))
			*(*uint16)(unsafe.Add(mBase, uint32(l1))) = uint16(v33)
			v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+78)))
			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v35)
			v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+128)))
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v37)
			F_ReleaseCatCache(m, v13)
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return
			} else {
				m.G0 = v9 + int32(16)
				return
			}
		}
	}
}
func F_get_typtype(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v5 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(l0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 == int32(0) {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+22)))
			v16 = int32(*(*int8)(unsafe.Add(mBase, uint32(v13+v14)+79)))
			F_ReleaseCatCache(m, v5)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				return base.I32_extend8_s(v16)
			}
		}
	}
}
func F_get_val(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
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
	var v156 int32
	_ = v156
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
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
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
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
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v318 int32
	_ = v318
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v14
	v17 = F_palloc(m, v14)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v17
	v23 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v23)
	v28 = int32(0)
	goto L3
L3:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	v36 = base.I32_extend8_s(v35)
	switch v28 {
	case 0:
		goto L12
	case 1:
		goto L11
	case 2:
		goto L10
	case 3:
		goto L9
	default:
		goto L8
	}
L4:
	;
	m.G0 = v12 + int32(16)
	return v318
L5:
	;
	goto L4
L6:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v311 + int32(1)
	v28 = v310
	goto L3
L7:
	;
	v310 = int32(2)
	goto L6
L8:
	;
	if v36 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L9:
	;
	if v36 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L10:
	;
	if v35 == int32(92) {
		v310 = int32(4)
		goto L6
	} else {
		goto L46
	}
L11:
	;
	if v36 == int32(92) {
		v310 = int32(3)
		goto L6
	} else {
		goto L30
	}
L12:
	;
	if v35 != int32(34) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if l1|base.B2i32(v36 != int32(61)) == int32(0) {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	if v35 != 0 {
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v40 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v40)
	goto L7
L17:
	;
	v318 = int32(0)
	goto L5
L18:
	;
	v47 = int32(0)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v49 = F_errsave_start(m, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if v36 == int32(92) {
		v310 = int32(3)
		goto L6
	} else {
		goto L27
	}
L21:
	;
	if v49 == int32(0) {
		v318 = v47
		goto L5
	} else {
		goto L22
	}
L22:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v57 = F_pg_mblen_cstr(m, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v60 - v59
	F_errmsg(m, int32(_a_F_get_val_0), v12)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_errsave_finish(m, v48, int32(_a_F_get_val_1), int32(71), int32(_a_F_get_val_2))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v318 = v47
	goto L5
L27:
	;
	goto L28
L28:
	;
	if base.B2i32(v36 == int32(32))|base.B2i32(base.Ui32((v36-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v310 = int32(0)
		goto L6
	} else {
		goto L29
	}
L29:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	*(*uint8)(unsafe.Add(mBase, uint32(v86))) = uint8(v88)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v91 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v90 + v91
	v310 = v91
	goto L6
L30:
	;
	if l1|base.B2i32(v36 != int32(61)) == int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v103 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v34 - v103
	v318 = v103
	goto L5
L32:
	;
	goto L33
L33:
	;
	v107 = int32(0)
	if base.B2i32(l1 == v107)|base.B2i32(v36 != int32(44)) == v107 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v114 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v34 - v114
	v318 = v114
	goto L5
L35:
	;
	goto L36
L36:
	;
	v118 = int32(1)
	goto L37
L37:
	;
	if base.B2i32(v36 == int32(32))|base.B2i32(base.Ui32((v36-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v318 = v118
		goto L5
	} else {
		goto L38
	}
L38:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	if v129 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v128 - int32(1)
	v318 = v118
	goto L5
L40:
	;
	goto L41
L41:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v138 = v136 - v137
	if v135 <= v138+int32(1) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v143 = v135 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v143
	v145 = F_repalloc(m, v137, v143)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	v152 = v136
	v153 = v129
	goto L44
L44:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v152))) = uint8(v153)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v156 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v155 + v156
	v310 = v156
	goto L6
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v145
	v148 = v145 + v138
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v148
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150))))
	v152 = v148
	v153 = v151
	goto L44
L46:
	;
	if v35 == int32(34) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v318 = int32(1)
	goto L5
L48:
	;
	goto L49
L49:
	;
	if v35 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v168 = int32(0)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v170 = F_errsave_start(m, v169)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v189 = v187 - v188
	if v186 <= v189+int32(1) {
		goto L58
	} else {
		goto L59
	}
L53:
	;
	if v170 == int32(0) {
		v318 = v168
		goto L5
	} else {
		goto L54
	}
L54:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_errmsg(m, int32(_a_F_get_val_3), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_errsave_finish(m, v169, int32(_a_F_get_val_1), int32(83), int32(_a_F_get_val_4))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v318 = v168
	goto L5
L58:
	;
	v194 = v186 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v194
	v196 = F_repalloc(m, v188, v194)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	v203 = v187
	v204 = v36
	goto L60
L60:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v203))) = uint8(v204)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v206 + int32(1)
	goto L7
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v196
	v199 = v196 + v189
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v199
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201))))
	v203 = v199
	v204 = v202
	goto L60
L62:
	;
	v212 = int32(0)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v214 = F_errsave_start(m, v213)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v233 = v231 - v232
	if v230 <= v233+int32(1) {
		goto L70
	} else {
		goto L71
	}
L65:
	;
	if v214 == int32(0) {
		v318 = v212
		goto L5
	} else {
		goto L66
	}
L66:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errmsg(m, int32(_a_F_get_val_3), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_errsave_finish(m, v213, int32(_a_F_get_val_1), int32(83), int32(_a_F_get_val_4))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v318 = v212
	goto L5
L70:
	;
	v238 = v230 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v238
	v240 = F_repalloc(m, v232, v238)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L73
	}
L71:
	;
	v247 = v231
	v248 = v36
	goto L72
L72:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v247))) = uint8(v248)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v251 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v250 + v251
	v310 = v251
	goto L6
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v240
	v243 = v240 + v233
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v243
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245))))
	v247 = v243
	v248 = v246
	goto L72
L74:
	;
	v257 = int32(0)
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v259 = F_errsave_start(m, v258)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v278 = v276 - v277
	if v275 <= v278+int32(1) {
		goto L82
	} else {
		goto L83
	}
L77:
	;
	if v259 == int32(0) {
		v318 = v257
		goto L5
	} else {
		goto L78
	}
L78:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	F_errmsg(m, int32(_a_F_get_val_3), int32(0))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	F_errsave_finish(m, v258, int32(_a_F_get_val_1), int32(83), int32(_a_F_get_val_4))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v318 = v257
	goto L5
L82:
	;
	v283 = v275 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v283
	v285 = F_repalloc(m, v277, v283)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L85
	}
L83:
	;
	v292 = v276
	v293 = v36
	goto L84
L84:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v292))) = uint8(v293)
	v295 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v295 + int32(1)
	goto L7
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v285
	v288 = v285 + v278
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v288
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290))))
	v292 = v288
	v293 = v291
	goto L84
}
func F_getrule(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v342 int32
	_ = v342
	var v351 int32
	_ = v351
	v3 = int32(0)
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	switch v7 - int32(74) {
	case 0:
		goto L5
	default:
		goto L3
	case 3:
		goto L4
	}
L1:
	;
	return v351
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v183
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184))))
	if v187 == int32(47) {
		goto L39
	} else {
		goto L40
	}
L3:
	;
	if base.Ui32(int32(9)) < base.Ui32(base.I32_extend8_s(v7)-int32(48)) {
		v351 = v3
		goto L1
	} else {
		goto L32
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(2)
	v44 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+1)))
	if base.Ui32(int32(9)) < base.Ui32(v44-int32(48)) {
		v351 = v3
		goto L1
	} else {
		goto L12
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	v12 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+1)))
	if base.Ui32(int32(9)) < base.Ui32(v12-int32(48)) {
		v351 = v3
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v19 = l0 + int32(1)
	v21 = v12
	v22 = v3
	goto L7
L7:
	;
	v30 = base.I32_extend8_s(v21) + v22*int32(10) - int32(48)
	if int32(365) < v30 {
		v351 = v3
		goto L1
	} else {
		goto L9
	}
L8:
	;
	if v30 <= int32(0) {
		v351 = v3
		goto L1
	} else {
		goto L11
	}
L9:
	;
	v33 = int32(*(*int8)(unsafe.Add(mBase, uint32(v19)+1)))
	v35 = v19 + int32(1)
	if base.Ui32(v33-int32(48)) < base.Ui32(int32(10)) {
		v19 = v35
		v21 = v33
		v22 = v30
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v183 = v30
	v184 = v35
	goto L2
L12:
	;
	v52 = int32(0)
	v55 = v44
	v56 = l0 + int32(1)
	goto L13
L13:
	;
	v63 = base.I32_extend8_s(v55) + v52*int32(10) - int32(48)
	if int32(12) < v63 {
		v351 = v3
		goto L1
	} else {
		goto L15
	}
L14:
	;
	if v63 <= int32(0) {
		v351 = v3
		goto L1
	} else {
		goto L17
	}
L15:
	;
	v67 = v56 + int32(1)
	v68 = int32(*(*int8)(unsafe.Add(mBase, uint32(v56)+1)))
	if base.Ui32(v68-int32(48)) < base.Ui32(int32(10)) {
		v52 = v63
		v55 = v68
		v56 = v67
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v63
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	if v76 != int32(46) {
		v351 = v3
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v79 = int32(*(*int8)(unsafe.Add(mBase, uint32(v56)+2)))
	if base.Ui32(int32(9)) < base.Ui32(v79-int32(48)) {
		v351 = v3
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v87 = int32(0)
	v90 = v79
	v91 = v56 + int32(2)
	goto L20
L20:
	;
	v98 = base.I32_extend8_s(v90) + v87*int32(10) - int32(48)
	if int32(5) < v98 {
		v351 = v3
		goto L1
	} else {
		goto L22
	}
L21:
	;
	if v98 <= int32(0) {
		v351 = v3
		goto L1
	} else {
		goto L24
	}
L22:
	;
	v102 = v91 + int32(1)
	v103 = int32(*(*int8)(unsafe.Add(mBase, uint32(v91)+1)))
	if base.Ui32(v103-int32(48)) < base.Ui32(int32(10)) {
		v87 = v98
		v90 = v103
		v91 = v102
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v98
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	if v111 != int32(46) {
		v351 = v3
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v114 = int32(*(*int8)(unsafe.Add(mBase, uint32(v91)+2)))
	if base.Ui32(int32(9)) < base.Ui32(v114-int32(48)) {
		v351 = v3
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v122 = v114
	v124 = v91 + int32(2)
	v125 = int32(0)
	goto L27
L27:
	;
	v133 = base.I32_extend8_s(v122) + v125*int32(10) - int32(48)
	if int32(6) < v133 {
		v351 = v3
		goto L1
	} else {
		goto L29
	}
L28:
	;
	if v133 < int32(0) {
		v351 = v3
		goto L1
	} else {
		goto L31
	}
L29:
	;
	v136 = int32(*(*int8)(unsafe.Add(mBase, uint32(v124)+1)))
	v138 = v124 + int32(1)
	if base.Ui32(v136-int32(48)) < base.Ui32(int32(10)) {
		v122 = v136
		v124 = v138
		v125 = v133
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v183 = v133
	v184 = v138
	goto L2
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1)
	v152 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
	if base.Ui32(int32(9)) < base.Ui32(v152-int32(48)) {
		v351 = v3
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v157 = l0
	v159 = v152
	v160 = v3
	goto L34
L34:
	;
	v168 = base.I32_extend8_s(v159) + v160*int32(10) - int32(48)
	if int32(365) < v168 {
		v351 = v3
		goto L1
	} else {
		goto L36
	}
L35:
	;
	if v168 < int32(0) {
		v351 = v3
		goto L1
	} else {
		goto L38
	}
L36:
	;
	v171 = int32(*(*int8)(unsafe.Add(mBase, uint32(v157)+1)))
	v173 = v157 + int32(1)
	if base.Ui32(v171-int32(48)) < base.Ui32(int32(10)) {
		v157 = v173
		v159 = v171
		v160 = v168
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v183 = v168
	v184 = v173
	goto L2
L39:
	;
	v190 = int32(1)
	v191 = v184 + v190
	v193 = l1 + int32(16)
	v194 = int32(0)
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191))))
	switch v201 - int32(43) {
	case 0:
		goto L44
	default:
		v209 = v191
		v210 = v190
		goto L43
	case 2:
		goto L45
	}
L40:
	;
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = int32(_a_F_getrule_0)
	v351 = v184
	goto L1
L42:
	;
	return v342
L43:
	;
	v211 = int32(*(*int8)(unsafe.Add(mBase, uint32(v209))))
	if base.Ui32(int32(9)) < base.Ui32(v211-int32(48)) {
		v342 = v194
		goto L46
	} else {
		goto L47
	}
L44:
	;
	v209 = v184 + int32(2)
	v210 = v190
	goto L43
L45:
	;
	v209 = v184 + int32(2)
	v210 = int32(0)
	goto L43
L46:
	;
	goto L42
L47:
	;
	v216 = v209
	v218 = v194
	v219 = v211
	goto L48
L48:
	;
	v229 = base.I32_extend8_s(v219) + v218*int32(10) - int32(48)
	if int32(167) < v229 {
		v342 = v194
		goto L46
	} else {
		goto L50
	}
L49:
	;
	if v229 < int32(0) {
		v342 = v194
		goto L46
	} else {
		goto L52
	}
L50:
	;
	v233 = v216 + int32(1)
	v234 = int32(*(*int8)(unsafe.Add(mBase, uint32(v216)+1)))
	if base.Ui32(v234-int32(48)) < base.Ui32(int32(10)) {
		v216 = v233
		v218 = v229
		v219 = v234
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v242 = v229 * int32(3600)
	*(*int32)(unsafe.Add(mBase, uint32(v193))) = v242
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233))))
	if v244 != int32(58) {
		v322 = v233
		v327 = v242
		goto L53
	} else {
		goto L54
	}
L53:
	;
	if v210 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L54:
	;
	v247 = int32(*(*int8)(unsafe.Add(mBase, uint32(v216)+2)))
	if base.Ui32(int32(9)) < base.Ui32(v247-int32(48)) {
		v342 = v194
		goto L46
	} else {
		goto L55
	}
L55:
	;
	v255 = v216 + int32(2)
	v257 = int32(0)
	v258 = v247
	goto L56
L56:
	;
	v268 = base.I32_extend8_s(v258) + v257*int32(10) - int32(48)
	if int32(59) < v268 {
		v342 = v194
		goto L46
	} else {
		goto L58
	}
L57:
	;
	if v268 < int32(0) {
		v342 = v194
		goto L46
	} else {
		goto L60
	}
L58:
	;
	v272 = v255 + int32(1)
	v273 = int32(*(*int8)(unsafe.Add(mBase, uint32(v255)+1)))
	if base.Ui32(v273-int32(48)) < base.Ui32(int32(10)) {
		v255 = v272
		v257 = v268
		v258 = v273
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	v282 = v268*int32(60) + v242
	*(*int32)(unsafe.Add(mBase, uint32(v193))) = v282
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272))))
	if v284 != int32(58) {
		v322 = v272
		v327 = v282
		goto L53
	} else {
		goto L61
	}
L61:
	;
	v287 = int32(*(*int8)(unsafe.Add(mBase, uint32(v255)+2)))
	if base.Ui32(int32(9)) < base.Ui32(v287-int32(48)) {
		v342 = v194
		goto L46
	} else {
		goto L62
	}
L62:
	;
	v295 = v255 + int32(2)
	v297 = v287
	v298 = int32(0)
	goto L63
L63:
	;
	v308 = base.I32_extend8_s(v297) + v298*int32(10) - int32(48)
	if int32(60) < v308 {
		v342 = v194
		goto L46
	} else {
		goto L65
	}
L64:
	;
	if v308 < int32(0) {
		v342 = v194
		goto L46
	} else {
		goto L67
	}
L65:
	;
	v311 = int32(*(*int8)(unsafe.Add(mBase, uint32(v295)+1)))
	v313 = v295 + int32(1)
	if base.Ui32(v311-int32(48)) < base.Ui32(int32(10)) {
		v295 = v313
		v297 = v311
		v298 = v308
		goto L63
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	v320 = v308 + v282
	*(*int32)(unsafe.Add(mBase, uint32(v193))) = v320
	v322 = v313
	v327 = v320
	goto L53
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v193))) = int32(0) - v327
	goto L70
L69:
	;
	goto L70
L70:
	;
	v342 = v322
	goto L46
}
func F_ginarrayextract(m *base.Module, l0 int32) int64 {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int64
	_ = v40
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum_copy(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
		F_get_typlenbyvalalign(m, v17, v8+int32(14), v8+int32(13), v8+int32(12))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int64(0)
		} else {
			v27 = int32(*(*int16)(unsafe.Add(mBase, uint32(v8)+14)))
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+13)))
			v29 = int32(*(*int8)(unsafe.Add(mBase, uint32(v8)+12)))
			F_deconstruct_array(m, v11, v27, v28, v29, v8+int32(8), v8+int32(4), v8)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int64(0)
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				*(*int32)(unsafe.Add(mBase, uint32(v16))) = v36
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v15))) = v38
				v40 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v8)+8)))
				m.G0 = v8 + int32(16)
				return v40
			}
		}
	}
}
func F_ginbuild(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int64
	_ = v39
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int64
	_ = v91
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
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
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v368 int64
	_ = v368
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
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
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
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
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
	var v457 int32
	_ = v457
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
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
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
	var v524 float64
	_ = v524
	var v526 float64
	_ = v526
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v533 float64
	_ = v533
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v586 int64
	_ = v586
	var v589 int64
	_ = v589
	var v593 int64
	_ = v593
	var v595 float64
	_ = v595
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v642 int32
	_ = v642
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v667 int64
	_ = v667
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v681 int64
	_ = v681
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v695 int64
	_ = v695
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v709 int64
	_ = v709
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v763 int32
	_ = v763
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v806 float64
	_ = v806
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v825 int64
	_ = v825
	var v826 int32
	_ = v826
	var v829 int64
	_ = v829
	var v831 int64
	_ = v831
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v864 int32
	_ = v864
	var v865 int64
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	var v880 int64
	_ = v880
	var v882 int32
	_ = v882
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v906 int32
	_ = v906
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v930 int32
	_ = v930
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v949 int32
	_ = v949
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v965 int32
	_ = v965
	var v969 int32
	_ = v969
	var v970 int64
	_ = v970
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v975 int32
	_ = v975
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v992 int32
	_ = v992
	var v995 int32
	_ = v995
	var v1002 int32
	_ = v1002
	var v1005 float64
	_ = v1005
	var v1009 int32
	_ = v1009
	var v1013 int32
	_ = v1013
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1024 int32
	_ = v1024
	var v1028 int32
	_ = v1028
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1046 int32
	_ = v1046
	var v1050 int32
	_ = v1050
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1076 int64
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1080 int32
	_ = v1080
	var v1081 int64
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1085 int32
	_ = v1085
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1092 int32
	_ = v1092
	var v1093 int64
	_ = v1093
	var v1095 int32
	_ = v1095
	var v1106 int32
	_ = v1106
	var v1110 int32
	_ = v1110
	var v1115 int32
	_ = v1115
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1121 int32
	_ = v1121
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1143 int32
	_ = v1143
	var v1148 int32
	_ = v1148
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1158 int32
	_ = v1158
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1178 float64
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1183 int32
	_ = v1183
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1197 int32
	_ = v1197
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1234 int32
	_ = v1234
	var v1236 int32
	_ = v1236
	var v1239 int32
	_ = v1239
	var v1240 int64
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1244 int32
	_ = v1244
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1288 float64
	_ = v1288
	var v1291 int32
	_ = v1291
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1296 int32
	_ = v1296
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1309 int32
	_ = v1309
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1320 int32
	_ = v1320
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1325 float64
	_ = v1325
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1341 int32
	_ = v1341
	var v1346 int32
	_ = v1346
	v18 = m.G0
	v20 = v18 - int32(_a_F_ginbuild_0)
	m.G0 = v20
	v23 = F_RelationGetNumberOfBlocksInFork(m, l1, int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v23 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v30 = v20 + int32(16)
	F_initGinState(m, v30, l1)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1334 = m.ExcPending
	if v1334 != 0 {
		goto L1
	} else {
		goto L252
	}
L6:
	;
	v33 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[0]))) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[1]))) = v33
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[2]))) = uint16(v33)
	v39 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[3]))) = v39
	*(*int64)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[4]))) = v39
	*(*int64)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[5]))) = v39
	*(*int64)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[6]))) = v39
	*(*int64)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[7]))) = v39
	*(*int64)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[8]))) = v39
	*(*int64)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[9]))) = v39
	v53 = F_GinNewBuffer(m, l1)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v55 = F_GinNewBuffer(m, l1)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v57 = int32(_a_F_ginbuild_1)
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10])) = v59 + int32(1)
	if v53 < int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	F_MarkBufferDirty(m, v53)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L14
	}
L10:
	;
	v83 = int32(8)
	F_PageInit(m, v81, int32(_a_F_ginbuild_2), v83)
	mBase = m.M
	v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v81)+16)))
	v86 = v81 + v85
	*(*int32)(unsafe.Add(mBase, uint32(v86))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v86)+6)) = uint16(v83)
	v91 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v81)+64)) = v91
	*(*int64)(unsafe.Add(mBase, uint32(v81)+24)) = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v81)+32)) = v91
	*(*int64)(unsafe.Add(mBase, uint32(v81)+40)) = v91
	*(*int64)(unsafe.Add(mBase, uint32(v81)+48)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v81)+56)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v81)+72)) = int32(2)
	v105 = int32(80)
	*(*uint16)(unsafe.Add(mBase, uint32(v81)+12)) = uint16(v105)
	goto L9
L11:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[11]))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v67+(v53^int32(-1))<<(uint(int32(2))%32))))
	v81 = v73
	goto L10
L12:
	;
	goto L13
L13:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[12]))
	v81 = v75 + v53<<(uint(int32(13))%32) + int32(-8192)
	goto L10
L14:
	;
	v109 = int32(2)
	if v55 < int32(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	F_MarkBufferDirty(m, v55)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L20
	}
L16:
	;
	F_PageInit(m, v127, int32(_a_F_ginbuild_2), int32(8))
	mBase = m.M
	v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v127)+16)))
	v132 = v127 + v131
	*(*int32)(unsafe.Add(mBase, uint32(v132))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v132)+6)) = uint16(v109)
	goto L15
L17:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[11]))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v113+(v55^int32(-1))<<(uint(int32(2))%32))))
	v127 = v119
	goto L16
L18:
	;
	goto L19
L19:
	;
	v121 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[12]))
	v127 = v121 + v55<<(uint(int32(13))%32) + int32(-8192)
	goto L16
L20:
	;
	F_UnlockReleaseBuffer(m, v53)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_UnlockReleaseBuffer(m, v55)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v142 = int32(_a_F_ginbuild_1)
	v144 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10]))
	v145 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10])) = v144 - v145
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[5])))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[5]))) = v148 + v145
	v153 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[13]))
	v158 = F_AllocSetContextCreateInternal(m, v153, int32(_a_F_ginbuild_3), int32(0), int32(_a_F_ginbuild_2), int32(_a_F_ginbuild_4))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[14]))) = v158
	v162 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[13]))
	v167 = F_AllocSetContextCreateInternal(m, v162, int32(_a_F_ginbuild_5), int32(0), int32(_a_F_ginbuild_2), int32(_a_F_ginbuild_4))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[15]))) = v167
	*(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[16]))) = v30
	v172 = v20 + int32(_a_F_ginbuild_6)
	F_ginInitBA(m, v172)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v179 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[17]))
	if v179 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l2)+128))
	if v220 <= int32(0) {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	goto L26
L28:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ginbuild[18])))
	if v183&int32(1) == int32(0) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v188 = int32(_a_F_ginbuild_1)
	v190 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10]))
	v191 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10])) = v190 + v191
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	*(*int32)(unsafe.Add(mBase, uint32(v179))) = v194 + v191
	v198 = int32(0)
	v200 = int32(_a_F_ginbuild_7)
	v201 = base.AtomicRmwOr32(m, v198, v200, v198)
	*(*int64)(unsafe.Add(mBase, uint32(v179+int32(80))+232)) = int64(2)
	v209 = base.AtomicRmwOr32(m, v198, v200, v198)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	*(*int32)(unsafe.Add(mBase, uint32(v179))) = v210 + v191
	v216 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10])) = v216 - v191
	goto L27
L30:
	;
	v468 = v20 + int32(_a_F_ginbuild_8)
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[1])))
	if v469 != 0 {
		goto L96
	} else {
		goto L97
	}
L31:
	;
	v224 = v220 + int32(1)
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+121)))
	v227 = F_palloc0(m, int32(28))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v231 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[19]))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v231)+72)) = v232 + int32(1)
	goto L33
L33:
	;
	v239 = F_CreateParallelContext(m, int32(_a_F_ginbuild_9), int32(_a_F_ginbuild_10), v220)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	if v225 == int32(1) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v243 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	v247 = int32(_a_F_ginbuild_11)
	goto L37
L37:
	;
	v249 = F_table_parallelscan_estimate(m, l0, v247)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	v245 = F_RegisterSnapshot(m, v243)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v247 = v245
	goto L37
L40:
	;
	v251 = F_add_size(m, int32(64), v249)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v239)+36))
	v258 = F_add_size(m, v253, (v251+int32(31))&int32(-32))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v239)+36)) = v258
	v261 = F_tuplesort_estimate_shared(m, v224)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v239)+36))
	v268 = F_add_size(m, v263, (v261+int32(31))&int32(-32))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v239)+36)) = v268
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v239)+40))
	v273 = F_add_size(m, v271, int32(2))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v239)+40)) = v273
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v239)+36))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v239)+12))
	v279 = F_mul_size(m, int32(40), v278)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v285 = F_add_size(m, v276, (v279+int32(31))&int32(-32))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v239)+36)) = v285
	v288 = int32(1)
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v239)+40))
	v291 = F_add_size(m, v289, v288)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v239)+40)) = v291
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v239)+36))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v239)+12))
	v297 = F_mul_size(m, int32(128), v296)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v303 = F_add_size(m, v294, (v297+int32(31))&int32(-32))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v239)+36)) = v303
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v239)+40))
	v308 = F_add_size(m, v306, int32(1))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v239)+40)) = v308
	v312 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[20]))
	if v312 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v239)+36))
	v314 = F_strlen(m, v312)
	mBase = m.M
	v319 = F_add_size(m, v313, v314&int32(-32)+int32(32))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L55
	}
L53:
	;
	v329 = v288
	goto L54
L54:
	;
	F_InitializeParallelDSM(m, v239)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L57
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v239)+36)) = v319
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v239)+40))
	v324 = F_add_size(m, v322, int32(1))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v239)+40)) = v324
	v329 = v314 + int32(1)
	goto L54
L57:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v239)+44))
	if v332 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v247)))
	if v335 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	goto L60
L60:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v239)+52))
	v350 = F_shm_toc_allocate(m, v349, v251)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L67
	}
L61:
	;
	F_UnregisterSnapshot(m, v247)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	F_DestroyParallelContext(m, v239)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L65
	}
L64:
	;
	goto L63
L65:
	;
	v344 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[19]))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v344)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v344)+72)) = v345 - int32(1)
	goto L66
L66:
	;
	goto L30
L67:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v350))) = v352
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v350)+12)) = v224
	*(*uint8)(unsafe.Add(mBase, uint32(v350)+8)) = uint8(v225)
	*(*int32)(unsafe.Add(mBase, uint32(v350)+4)) = v354
	v359 = v350 + int32(16)
	v360 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v359))), uint32(v360))
	*(*int64)(unsafe.Add(mBase, uint32(v359)+4)) = int64(-1)
	goto L68
L68:
	;
	v365 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v350)+28)), uint32(v365))
	v368 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v350)+40)) = v368
	*(*int32)(unsafe.Add(mBase, uint32(v350)+32)) = v365
	*(*int64)(unsafe.Add(mBase, uint32(v350)+48)) = v368
	F_table_parallelscan_initialize(m, l0, v350-int32(-64), v247)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v239)+52))
	v379 = F_shm_toc_allocate(m, v378, v261)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v239)+44))
	F_tuplesort_initialize_shared(m, v379, v224, v381)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v239)+52))
	F_shm_toc_insert(m, v384, int64(-5764607523034234879), v350)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v239)+52))
	F_shm_toc_insert(m, v388, int64(-5764607523034234878), v379)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v393 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[20]))
	if v393 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v239)+52))
	v395 = F_shm_toc_allocate(m, v394, v329)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v239)+52))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v239)+12))
	v408 = F_mul_size(m, int32(40), v407)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L82
	}
L77:
	;
	if v329 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v398 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[20]))
	base.MemoryCopy(m, v395, v398, v329)
	goto L80
L79:
	;
	goto L80
L80:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v239)+52))
	F_shm_toc_insert(m, v400, int64(-5764607523034234877), v395)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	goto L76
L82:
	;
	v410 = F_shm_toc_allocate(m, v405, v408)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v239)+52))
	F_shm_toc_insert(m, v412, int64(-5764607523034234876), v410)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v239)+52))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v239)+12))
	v419 = F_mul_size(m, int32(128), v418)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v421 = F_shm_toc_allocate(m, v416, v419)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v239)+52))
	F_shm_toc_insert(m, v423, int64(-5764607523034234875), v421)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	F_LaunchParallelWorkers(m, v239)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v227))) = v239
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v239)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v227)+24)) = v421
	*(*int32)(unsafe.Add(mBase, uint32(v227)+20)) = v410
	*(*int32)(unsafe.Add(mBase, uint32(v227)+16)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v227)+12)) = v379
	*(*int32)(unsafe.Add(mBase, uint32(v227)+8)) = v350
	*(*int32)(unsafe.Add(mBase, uint32(v227)+4)) = v430 + int32(1)
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v239)+20))
	if v439 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	F__brin_end_parallel(m, v227)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L1
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v445 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[21]))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[1]))) = v227
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v227)+8))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v227)+12))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v227)+4))
	v452 = base.I32_div_s(v445, v451)
	F__gin_parallel_scan_and_build(m, v20+int32(16), v449, v450, l0, l1, v452, int32(1))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L1
	} else {
		goto L93
	}
L92:
	;
	goto L30
L93:
	;
	F_WaitForParallelWorkersToAttach(m, v239)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	goto L30
L95:
	;
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[15])))
	F_MemoryContextDelete(m, v1291)
	mBase = m.M
	v1293 = m.ExcPending
	if v1293 != 0 {
		goto L1
	} else {
		goto L238
	}
L96:
	;
	v471 = F_palloc0(m, int32(12))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L1
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v1167 = int32(0)
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v1177 = *(*int32)(unsafe.Add(mBase, uint32(v1176)+140))
	v1178 = m.T0[v1177].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, l0, l1, l2, v1167, v1167, int32(1), v1167, int32(-1), int32(56), v20+int32(16), v1167)
	mBase = m.M
	v1179 = m.ExcPending
	if v1179 != 0 {
		goto L1
	} else {
		goto L223
	}
L99:
	;
	v473 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v471))) = uint8(v473)
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[1])))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v475)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v471)+4)) = v476
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[1])))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v478)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v471)+8)) = v479
	v482 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[21]))
	v483 = F_tuplesort_begin_index_gin(m, l1, v482, v471)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[22]))) = v483
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[1])))
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v486)+8))
	v491 = v487 + int32(28)
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v486)+4))
	goto L101
L101:
	;
	v512 = base.AtomicRmwXchg32(m, v491, int32(0), int32(1))
	if v512 != 0 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v524 = *(*float64)(unsafe.Add(mBase, uint32(v487)+40))
	*(*float64)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[9]))) = v524
	v526 = *(*float64)(unsafe.Add(mBase, uint32(v487)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[8]))) = v526
	v528 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v487)+28)), uint32(v528))
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L1
	} else {
		goto L111
	}
L103:
	;
	F_s_lock(m, v491, int32(_a_F_ginbuild_12))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L1
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v487)+32))
	if v492 != v516 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	goto L105
L107:
	;
	v518 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v491))), uint32(v518))
	F_ConditionVariableSleep(m, v487+int32(16), int32(134217767))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L1
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	goto L102
L110:
	;
	goto L101
L111:
	;
	v533 = *(*float64)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[9])))
	v538 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[17]))
	if v538 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[22])))
	F_tuplesort_performsort(m, v579)
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L1
	} else {
		goto L116
	}
L113:
	;
	goto L112
L114:
	;
	v542 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ginbuild[18])))
	if v542&int32(1) == int32(0) {
		goto L113
	} else {
		goto L115
	}
L115:
	;
	v547 = int32(_a_F_ginbuild_1)
	v549 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10]))
	v550 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10])) = v549 + v550
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v538)))
	*(*int32)(unsafe.Add(mBase, uint32(v538))) = v553 + v550
	v557 = int32(0)
	v559 = int32(_a_F_ginbuild_7)
	v560 = base.AtomicRmwOr32(m, v557, v559, v557)
	*(*int64)(unsafe.Add(mBase, uint32(v538+int32(80))+232)) = int64(5)
	v568 = base.AtomicRmwOr32(m, v557, v559, v557)
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v538)))
	*(*int32)(unsafe.Add(mBase, uint32(v538))) = v569 + v550
	v575 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10])) = v575 - v550
	goto L113
L116:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v583 = F_GinBufferInit(m, v582)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	v586 = *(*int64)(unsafe.Add(mBase, _c_F_ginbuild[23]))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[24]))) = v586
	v589 = *(*int64)(unsafe.Add(mBase, _c_F_ginbuild[25]))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[26]))) = v589
	*(*int64)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[27]))) = int64(6)
	v593 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[28]))) = v593
	v595 = *(*float64)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[8])))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[29]))) = base.I64_trunc_sat_f64_s(v595)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[30]))) = v593
	v602 = v20 + int32(_a_F_ginbuild_13)
	v604 = v20 + int32(_a_F_ginbuild_14)
	goto L120
L118:
	;
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[22])))
	v787 = F_tuplesort_getgintuple(m, v786, v604)
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L1
	} else {
		goto L135
	}
L119:
	;
	goto L118
L120:
	;
	v614 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[17]))
	if v614 == int32(0) {
		goto L119
	} else {
		goto L121
	}
L121:
	;
	v618 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ginbuild[18])))
	if v618&int32(1) == int32(0) {
		goto L119
	} else {
		goto L122
	}
L122:
	;
	v623 = int32(_a_F_ginbuild_1)
	v625 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10]))
	v626 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10])) = v625 + v626
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v614)))
	*(*int32)(unsafe.Add(mBase, uint32(v614))) = v629 + v626
	v633 = int32(0)
	v636 = base.AtomicRmwOr32(m, v633, int32(_a_F_ginbuild_7), v633)
	goto L124
L123:
	;
	v763 = int32(0)
	v766 = base.AtomicRmwOr32(m, v763, int32(_a_F_ginbuild_7), v763)
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v614)))
	v768 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v614))) = v767 + v768
	v771 = int32(_a_F_ginbuild_1)
	v773 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10])) = v773 - v768
	goto L119
L124:
	;
	v642 = v614 + int32(232)
	goto L125
L125:
	;
	v648 = int32(0)
	v651 = int32(0)
	goto L128
L128:
	;
	v657 = int32(2)
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v602+v651<<(uint(v657)%32))))
	v661 = int32(3)
	v667 = *(*int64)(unsafe.Add(mBase, uint32(v604+v651<<(uint(v661)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v642+v660<<(uint(v661)%32)))) = v667
	v670 = v651 | int32(1)
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v602+v670<<(uint(v657)%32))))
	v681 = *(*int64)(unsafe.Add(mBase, uint32(v604+v670<<(uint(v661)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v642+v674<<(uint(v661)%32)))) = v681
	v684 = v651 | v657
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v602+v684<<(uint(v657)%32))))
	v695 = *(*int64)(unsafe.Add(mBase, uint32(v604+v684<<(uint(v661)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v642+v688<<(uint(v661)%32)))) = v695
	v698 = v651 | v661
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v602+v698<<(uint(v657)%32))))
	v709 = *(*int64)(unsafe.Add(mBase, uint32(v604+v698<<(uint(v661)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v642+v702<<(uint(v661)%32)))) = v709
	v711 = int32(4)
	v714 = v648 + v711
	if v714 != int32(4) {
		v648 = v714
		v651 = v651 + v711
		goto L128
	} else {
		goto L130
	}
L129:
	;
	goto L123
L130:
	;
	goto L129
L135:
	;
	if v787 != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v790 = v583 + int32(8)
	v793 = v787
	v806 = float64(0)
	goto L139
L137:
	;
	v1076 = int64(1)
	goto L138
L138:
	;
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v583)+28))
	if v1077 != 0 {
		goto L199
	} else {
		goto L200
	}
L139:
	;
	v809 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[31]))
	if v809 != 0 {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	v1076 = base.I64_trunc_sat_f64_s(base.F64_add(v1005, float64(1)))
	goto L138
L141:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L1
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v583)+28))
	if v812 == int32(0) {
		goto L145
	} else {
		goto L146
	}
L144:
	;
	goto L143
L145:
	;
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v583)+28))
	if v897 <= int32(0) {
		v915 = v897
		goto L170
	} else {
		goto L171
	}
L146:
	;
	v815 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v793)+4)))
	v816 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v583))))
	if v815 != v816 {
		v854 = v812
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v857 = int32(_a_F_ginbuild_15)
	v858 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[13]))
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[14])))
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[13])) = v860
	v864 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v583))))
	v865 = *(*int64)(unsafe.Add(mBase, uint32(v583)+8))
	v866 = int32(*(*int8)(unsafe.Add(mBase, uint32(v583)+2)))
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v583)+40))
	F_ginEntryInsert(m, v20+int32(16), v864, v865, v866, v867, v854, v468)
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L1
	} else {
		goto L163
	}
L148:
	;
	v818 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v793)+15)))
	v819 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v583)+2)))
	if v818 != v819 {
		v854 = v812
		goto L147
	} else {
		goto L149
	}
L149:
	;
	if v818 != 0 {
		goto L145
	} else {
		goto L150
	}
L150:
	;
	v822 = v793 + int32(24)
	v823 = int32(1)
	v825 = *(*int64)(unsafe.Add(mBase, uint32(v583)+8))
	v826 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v583)+22)))
	if v826 == v823 {
		goto L152
	} else {
		goto L153
	}
L151:
	;
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v583)+36))
	v833 = int32(36)
	v835 = v832 + v815*v833
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v835-int32(20))))
	v841 = m.T0[v840].(func(*base.Module, int64, int64, int32) int32)(m, v825, v831, v835-v833)
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L1
	} else {
		goto L155
	}
L152:
	;
	v829 = *(*int64)(unsafe.Add(mBase, uint32(v822)))
	v831 = v829
	goto L151
L153:
	;
	goto L154
L154:
	;
	v831 = base.I64_extend_i32_u(v822)
	goto L151
L155:
	;
	if v841 < int32(0) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v846 = v823
	goto L158
L157:
	;
	v846 = int32(0) - v841
	goto L158
L158:
	;
	v849 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v835-int32(28)))))
	if v849 != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v850 = v846
	goto L161
L160:
	;
	v850 = v841
	goto L161
L161:
	;
	if v850 == int32(0) {
		goto L145
	} else {
		goto L162
	}
L162:
	;
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v583)+28))
	v854 = v853
	goto L147
L163:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[13])) = v858
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[14])))
	F_MemoryContextReset(m, v872)
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	v875 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v583)+2)))
	if v875 != 0 {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v880 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v583)+28)) = v880
	v882 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v583)+2)) = uint8(v882)
	*(*uint16)(unsafe.Add(mBase, uint32(v583))) = uint16(v882)
	*(*int64)(unsafe.Add(mBase, uint32(v790)+7)) = v880
	*(*int64)(unsafe.Add(mBase, uint32(v790))) = v880
	goto L145
L166:
	;
	v876 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v583)+22)))
	if v876 != 0 {
		goto L165
	} else {
		goto L167
	}
L167:
	;
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v790)))
	F_pfree(m, v877)
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	goto L165
L169:
	;
	if v961 != 0 {
		goto L184
	} else {
		goto L185
	}
L170:
	;
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v583)+32))
	if v915 <= v916 {
		goto L174
	} else {
		goto L175
	}
L171:
	;
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v583)+40))
	v901 = int32(6)
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v793)+8))
	v912 = F_itemptr_comparator(m, v900+v897*v901-v901, (v793+v906+int32(25))&int32(-2))
	mBase = m.M
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v583)+28))
	if v912 != 0 {
		v915 = v913
		goto L170
	} else {
		goto L172
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v583)+32)) = v913
	v915 = v913
	goto L170
L173:
	;
	if int32(1024) <= v949 {
		goto L181
	} else {
		goto L182
	}
L174:
	;
	v949 = v916
	goto L173
L175:
	;
	goto L176
L176:
	;
	v923 = v916
	goto L177
L177:
	;
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v583)+40))
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v793)+8))
	v936 = F_itemptr_comparator(m, v926+v923*int32(6), (v793+int32(24)+v930+int32(1))&int32(-2))
	mBase = m.M
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v583)+32))
	if int32(0) < v936 {
		v949 = v937
		goto L173
	} else {
		goto L179
	}
L178:
	;
	v949 = v941
	goto L173
L179:
	;
	v940 = int32(1)
	v941 = v937 + v940
	*(*int32)(unsafe.Add(mBase, uint32(v583)+32)) = v941
	v944 = v923 + v940
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v583)+28))
	if v944 < v945 {
		v923 = v944
		goto L177
	} else {
		goto L180
	}
L180:
	;
	goto L178
L181:
	;
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v583)+24))
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v793)+16))
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v583)+28))
	v961 = base.B2i32(v955 <= v956+v957)
	goto L183
L182:
	;
	v961 = int32(0)
	goto L183
L183:
	;
	goto L169
L184:
	;
	v962 = int32(_a_F_ginbuild_15)
	v963 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[13]))
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[14])))
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[13])) = v965
	v969 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v583))))
	v970 = *(*int64)(unsafe.Add(mBase, uint32(v583)+8))
	v971 = int32(*(*int8)(unsafe.Add(mBase, uint32(v583)+2)))
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v583)+40))
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v583)+32))
	F_ginEntryInsert(m, v20+int32(16), v969, v970, v971, v972, v973, v468)
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L1
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	F_GinBufferStoreTuple(m, v583, v793)
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L1
	} else {
		goto L192
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[13])) = v963
	v978 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[14])))
	F_MemoryContextReset(m, v978)
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L1
	} else {
		goto L188
	}
L188:
	;
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v583)+28))
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v583)+32))
	v985 = (v981 - v982) * int32(6)
	if v985 != 0 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v583)+40))
	base.MemoryCopy(m, v986, v986+v982*int32(6), v985)
	goto L191
L190:
	;
	goto L191
L191:
	;
	v992 = *(*int32)(unsafe.Add(mBase, uint32(v583)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v583)+32)) = int32(0)
	v995 = *(*int32)(unsafe.Add(mBase, uint32(v583)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v583)+28)) = v995 - v992
	goto L186
L192:
	;
	v1005 = base.F64_add(v806, float64(1))
	v1009 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[17]))
	if v1009 == int32(0) {
		goto L194
	} else {
		goto L195
	}
L193:
	;
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[22])))
	v1053 = F_tuplesort_getgintuple(m, v1050, v20+int32(_a_F_ginbuild_14))
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L1
	} else {
		goto L197
	}
L194:
	;
	goto L193
L195:
	;
	v1013 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ginbuild[18])))
	if v1013&int32(1) == int32(0) {
		goto L194
	} else {
		goto L196
	}
L196:
	;
	v1018 = int32(_a_F_ginbuild_1)
	v1020 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10]))
	v1021 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10])) = v1020 + v1021
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(v1009)))
	*(*int32)(unsafe.Add(mBase, uint32(v1009))) = v1024 + v1021
	v1028 = int32(0)
	v1030 = int32(_a_F_ginbuild_7)
	v1031 = base.AtomicRmwOr32(m, v1028, v1030, v1028)
	*(*int64)(unsafe.Add(mBase, uint32(v1009+int32(96))+232)) = base.I64_trunc_sat_f64_s(v1005)
	v1039 = base.AtomicRmwOr32(m, v1028, v1030, v1028)
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v1009)))
	*(*int32)(unsafe.Add(mBase, uint32(v1009))) = v1040 + v1021
	v1046 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10])) = v1046 - v1021
	goto L194
L197:
	;
	if v1053 != 0 {
		v793 = v1053
		v806 = v1005
		goto L139
	} else {
		goto L198
	}
L198:
	;
	goto L140
L199:
	;
	v1080 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v583))))
	v1081 = *(*int64)(unsafe.Add(mBase, uint32(v583)+8))
	v1082 = int32(*(*int8)(unsafe.Add(mBase, uint32(v583)+2)))
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(v583)+40))
	F_ginEntryInsert(m, v20+int32(16), v1080, v1081, v1082, v1083, v1077, v468)
	mBase = m.M
	v1085 = m.ExcPending
	if v1085 != 0 {
		goto L1
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(v583)+40))
	if v1148 != 0 {
		goto L211
	} else {
		goto L212
	}
L202:
	;
	v1087 = v583 + int32(8)
	v1088 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v583)+2)))
	if v1088 != 0 {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v1093 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v583)+28)) = v1093
	v1095 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v583)+2)) = uint8(v1095)
	*(*uint16)(unsafe.Add(mBase, uint32(v583))) = uint16(v1095)
	*(*int64)(unsafe.Add(mBase, uint32(v1087)+7)) = v1093
	*(*int64)(unsafe.Add(mBase, uint32(v1087))) = v1093
	v1106 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[17]))
	if v1106 == v1095 {
		goto L208
	} else {
		goto L209
	}
L204:
	;
	v1089 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v583)+22)))
	if v1089 != 0 {
		goto L203
	} else {
		goto L205
	}
L205:
	;
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(v1087)))
	F_pfree(m, v1090)
	mBase = m.M
	v1092 = m.ExcPending
	if v1092 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	goto L203
L207:
	;
	goto L201
L208:
	;
	goto L207
L209:
	;
	v1110 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ginbuild[18])))
	if v1110&int32(1) == int32(0) {
		goto L208
	} else {
		goto L210
	}
L210:
	;
	v1115 = int32(_a_F_ginbuild_1)
	v1117 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10]))
	v1118 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10])) = v1117 + v1118
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(v1106)))
	*(*int32)(unsafe.Add(mBase, uint32(v1106))) = v1121 + v1118
	v1125 = int32(0)
	v1127 = int32(_a_F_ginbuild_7)
	v1128 = base.AtomicRmwOr32(m, v1125, v1127, v1125)
	*(*int64)(unsafe.Add(mBase, uint32(v1106+int32(96))+232)) = v1076
	v1136 = base.AtomicRmwOr32(m, v1125, v1127, v1125)
	v1137 = *(*int32)(unsafe.Add(mBase, uint32(v1106)))
	*(*int32)(unsafe.Add(mBase, uint32(v1106))) = v1137 + v1118
	v1143 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10])) = v1143 - v1118
	goto L208
L211:
	;
	F_pfree(m, v1148)
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L1
	} else {
		goto L214
	}
L212:
	;
	goto L213
L213:
	;
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(v583)+28))
	if v1151 == int32(0) {
		goto L215
	} else {
		goto L216
	}
L214:
	;
	goto L213
L215:
	;
	F_pfree(m, v583)
	mBase = m.M
	v1160 = m.ExcPending
	if v1160 != 0 {
		goto L1
	} else {
		goto L220
	}
L216:
	;
	v1154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v583)+2)))
	if v1154 != 0 {
		goto L215
	} else {
		goto L217
	}
L217:
	;
	v1155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v583)+22)))
	if v1155 != 0 {
		goto L215
	} else {
		goto L218
	}
L218:
	;
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v583)+8))
	F_pfree(m, v1156)
	mBase = m.M
	v1158 = m.ExcPending
	if v1158 != 0 {
		goto L1
	} else {
		goto L219
	}
L219:
	;
	goto L215
L220:
	;
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[22])))
	F_tuplesort_end(m, v1161)
	mBase = m.M
	v1163 = m.ExcPending
	if v1163 != 0 {
		goto L1
	} else {
		goto L221
	}
L221:
	;
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[1])))
	F__brin_end_parallel(m, v1164)
	mBase = m.M
	v1166 = m.ExcPending
	if v1166 != 0 {
		goto L1
	} else {
		goto L222
	}
L222:
	;
	v1288 = v533
	goto L95
L223:
	;
	v1180 = int32(_a_F_ginbuild_15)
	v1181 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[13]))
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[14])))
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[13])) = v1183
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[32])))
	v1188 = m.G0
	v1189 = int32(16)
	v1190 = v1188 - v1189
	m.G0 = v1190
	*(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[33]))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[34]))) = v1187
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(v1187)))
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[35]))) = uint8(base.B2i32(v1197 == int32(_a_F_ginbuild_16)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[36]))) = int32(836)
	m.G0 = v1190 + v1189
	goto L224
L224:
	;
	v1214 = F_ginGetBAEntry(m, v172, v20+int32(12), v20+int32(_a_F_ginbuild_14), v20+int32(15), v20+int32(_a_F_ginbuild_13))
	mBase = m.M
	v1215 = m.ExcPending
	if v1215 != 0 {
		goto L1
	} else {
		goto L225
	}
L225:
	;
	if v1214 != 0 {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	v1216 = v1214
	goto L229
L227:
	;
	goto L228
L228:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[13])) = v1181
	v1288 = v1178
	goto L95
L229:
	;
	v1234 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[31]))
	if v1234 != 0 {
		goto L231
	} else {
		goto L232
	}
L230:
	;
	goto L228
L231:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1236 = m.ExcPending
	if v1236 != 0 {
		goto L1
	} else {
		goto L234
	}
L232:
	;
	goto L233
L233:
	;
	v1239 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+12)))
	v1240 = *(*int64)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[27])))
	v1241 = int32(*(*int8)(unsafe.Add(mBase, uint32(v20)+15)))
	v1242 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[26])))
	F_ginEntryInsert(m, v20+int32(16), v1239, v1240, v1241, v1216, v1242, v468)
	mBase = m.M
	v1244 = m.ExcPending
	if v1244 != 0 {
		goto L1
	} else {
		goto L235
	}
L234:
	;
	goto L233
L235:
	;
	v1253 = F_ginGetBAEntry(m, v172, v20+int32(12), v20+int32(_a_F_ginbuild_14), v20+int32(15), v20+int32(_a_F_ginbuild_13))
	mBase = m.M
	v1254 = m.ExcPending
	if v1254 != 0 {
		goto L1
	} else {
		goto L236
	}
L236:
	;
	if v1253 != 0 {
		v1216 = v1253
		goto L229
	} else {
		goto L237
	}
L237:
	;
	goto L230
L238:
	;
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[14])))
	F_MemoryContextDelete(m, v1294)
	mBase = m.M
	v1296 = m.ExcPending
	if v1296 != 0 {
		goto L1
	} else {
		goto L239
	}
L239:
	;
	v1298 = F_RelationGetNumberOfBlocksInFork(m, l1, int32(0))
	mBase = m.M
	v1299 = m.ExcPending
	if v1299 != 0 {
		goto L1
	} else {
		goto L240
	}
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[37]))) = v1298
	F_ginUpdateStats(m, l1, v468, int32(1))
	mBase = m.M
	v1303 = m.ExcPending
	if v1303 != 0 {
		goto L1
	} else {
		goto L241
	}
L241:
	;
	v1304 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v1305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1304)+118)))
	if v1305 != int32(112) {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	v1322 = F_palloc(m, int32(16))
	mBase = m.M
	v1323 = m.ExcPending
	if v1323 != 0 {
		goto L1
	} else {
		goto L251
	}
L243:
	;
	v1309 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[38]))
	if v1309 <= int32(0) {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v1312 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v1312 != 0 {
		goto L242
	} else {
		goto L247
	}
L245:
	;
	goto L246
L246:
	;
	v1314 = int32(0)
	v1316 = F_RelationGetNumberOfBlocksInFork(m, l1, v1314)
	mBase = m.M
	v1317 = m.ExcPending
	if v1317 != 0 {
		goto L1
	} else {
		goto L249
	}
L247:
	;
	v1313 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v1313 != 0 {
		goto L242
	} else {
		goto L248
	}
L248:
	;
	goto L246
L249:
	;
	F_log_newpage_range(m, l1, v1314, v1316, int32(1))
	mBase = m.M
	v1320 = m.ExcPending
	if v1320 != 0 {
		goto L1
	} else {
		goto L250
	}
L250:
	;
	goto L242
L251:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1322))) = v1288
	v1325 = *(*float64)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[3])))
	*(*float64)(unsafe.Add(mBase, uint32(v1322)+8)) = v1325
	m.G0 = v20 + int32(_a_F_ginbuild_0)
	return v1322
L252:
	;
	v1335 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v1335 + int32(4)
	F_errmsg_internal(m, int32(_a_F_ginbuild_17), v20)
	mBase = m.M
	v1341 = m.ExcPending
	if v1341 != 0 {
		goto L1
	} else {
		goto L253
	}
L253:
	;
	F_errfinish(m, int32(_a_F_ginbuild_18), int32(635), int32(_a_F_ginbuild_19))
	mBase = m.M
	v1346 = m.ExcPending
	if v1346 != 0 {
		goto L1
	} else {
		goto L254
	}
L254:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gincost_pattern(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
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
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int64
	_ = v67
	var v80 int64
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v93 float64
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v105 float64
	_ = v105
	var v109 int32
	_ = v109
	var v112 float64
	_ = v112
	var v116 float64
	_ = v116
	var v121 float64
	_ = v121
	var v124 int32
	_ = v124
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 float64
	_ = v142
	var v143 float64
	_ = v143
	var v146 float64
	_ = v146
	var v151 int32
	_ = v151
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	v6 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(80)
	m.G0 = v13
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v6
	v26 = l1 << (uint(int32(2)) % 32)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26+v27)))
	F_get_op_opfamily_properties(m, l2, v29, v6, v13+int32(48), v13+int32(44), v13+int32(40))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		return int32(0)
	} else {
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v43 = *(*int32)(unsafe.Add(mBase, uint32(v41+v26)))
		v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		v46 = *(*int32)(unsafe.Add(mBase, uint32(v44+v26)))
		v48 = F_get_opfamily_proc(m, v43, v46, v46, int32(3))
		mBase = m.M
		v49 = m.ExcPending
		if v49 != 0 {
			return int32(0)
		} else {
			if v48 != 0 {
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v50+v26)))
				v54 = v13 + int32(52)
				F_fmgr_info(m, v48, v54)
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v57+v26)))
					F_set_fn_opclass_options(m, v54, v59)
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						if v52 != 0 {
							v63 = v52
						} else {
							v63 = int32(100)
						}
						v67 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v13)+48)))
						v80 = F_FunctionCall7Coll(m, v54, v63, l3, base.I64_extend_i32_u(v13+int32(36)), v67, base.I64_extend_i32_u(v13+int32(32)), base.I64_extend_i32_u(v13+int32(28)), base.I64_extend_i32_u(v13+int32(24)), base.I64_extend_i32_u(v13+int32(20)))
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return int32(0)
						} else {
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
							v83 = int32(0)
							v85 = *(*int32)(unsafe.Add(mBase, uint32(v13)+36))
							v88 = base.B2i32(v82 != v83) | base.B2i32(v83 < v85)
							if v88 == v83 {
							} else {
								if int32(0) < v85 {
									v93 = *(*float64)(unsafe.Add(mBase, uint32(l4)+80))
									v95 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
									v96 = int32(0)
									v105 = v93
									for {
										if v95 == int32(0) {
											v116 = *(*float64)(unsafe.Add(mBase, uint32(l4)+72))
											*(*float64)(unsafe.Add(mBase, uint32(l4)+72)) = base.F64_add(v116, float64(1))
										} else {
											v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96+v95))))
											if v109 != int32(1) {
												v116 = *(*float64)(unsafe.Add(mBase, uint32(l4)+72))
												*(*float64)(unsafe.Add(mBase, uint32(l4)+72)) = base.F64_add(v116, float64(1))
											} else {
												v112 = *(*float64)(unsafe.Add(mBase, uint32(l4)+64))
												*(*float64)(unsafe.Add(mBase, uint32(l4)+64)) = base.F64_add(v112, float64(100))
											}
										}
										v121 = base.F64_add(v105, float64(1))
										*(*float64)(unsafe.Add(mBase, uint32(l4)+80)) = v121
										v124 = v96 + int32(1)
										if v124 != v85 {
											v96 = v124
											v105 = v121
											continue
										} else {
											break
										}
										break
									}
								} else {
								}
								switch v82 {
								case 0:
									v137 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l1+l4)+32)) = uint8(v137)
								case 1:
									v140 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l1+l4)+32)) = uint8(v140)
									v142 = *(*float64)(unsafe.Add(mBase, uint32(l4)+72))
									v143 = float64(1)
									*(*float64)(unsafe.Add(mBase, uint32(l4)+72)) = base.F64_add(v142, v143)
									v146 = *(*float64)(unsafe.Add(mBase, uint32(l4)+80))
									*(*float64)(unsafe.Add(mBase, uint32(l4)+80)) = base.F64_add(v146, v143)
								default:
									v151 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l1+l4))) = uint8(v151)
								}
							}
							m.G0 = v13 + int32(80)
							return v88
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v170 = m.ExcPending
				if v170 != 0 {
					return int32(0)
				} else {
					v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v172 = F_get_rel_name(m, v171)
					mBase = m.M
					v173 = m.ExcPending
					if v173 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v172
						*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = l1 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(3)
						F_errmsg_internal(m, int32(_a_F_gincost_pattern_0), v13)
						mBase = m.M
						v182 = m.ExcPending
						if v182 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_gincost_pattern_1), int32(_a_F_gincost_pattern_2), int32(_a_F_gincost_pattern_3))
							mBase = m.M
							v187 = m.ExcPending
							if v187 != 0 {
								return int32(0)
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
func F_ginoptions(m *base.Module, l0 int64, l1 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_build_reloptions(m, l0, l1, int32(16), int32(12), int32(_a_F_ginoptions_0), int32(2))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_gistbuildempty(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int64
	_ = v10
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	v3 = m.G0
	v5 = v3 - int32(32)
	m.G0 = v5
	*(*int64)(unsafe.Add(mBase, uint32(v5)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+20)) = l0
	v10 = *(*int64)(unsafe.Add(mBase, uint32(v5)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v5)+8)) = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v5)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = v12
	v19 = F_ExtendBufferedRel(m, v5+int32(8), int32(3), int32(0), int32(9))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return
	} else {
		v21 = int32(_a_F_gistbuildempty_0)
		v23 = *(*int32)(unsafe.Add(mBase, _c_F_gistbuildempty[0]))
		v24 = int32(1)
		*(*int32)(unsafe.Add(mBase, _c_F_gistbuildempty[0])) = v23 + v24
		if v19 < int32(0) {
			v31 = *(*int32)(unsafe.Add(mBase, _c_F_gistbuildempty[1]))
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v31+(v19^int32(-1))<<(uint(int32(2))%32))))
			v45 = v37
		} else {
			v39 = *(*int32)(unsafe.Add(mBase, _c_F_gistbuildempty[2]))
			v45 = v39 + v19<<(uint(int32(13))%32) + int32(-8192)
		}
		F_PageInit(m, v45, int32(_a_F_gistbuildempty_1), int32(16))
		mBase = m.M
		v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+16)))
		v50 = v45 + v49
		v51 = int32(_a_F_gistbuildempty_2)
		*(*uint16)(unsafe.Add(mBase, uint32(v50)+14)) = uint16(v51)
		*(*uint16)(unsafe.Add(mBase, uint32(v50)+12)) = uint16(v24)
		*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = int32(-1)
		F_MarkBufferDirty(m, v19)
		mBase = m.M
		v57 = m.ExcPending
		if v57 != 0 {
			return
		} else {
			F_log_newpage_buffer(m, v19, int32(1))
			mBase = m.M
			v60 = m.ExcPending
			if v60 != 0 {
				return
			} else {
				v61 = int32(_a_F_gistbuildempty_0)
				v63 = *(*int32)(unsafe.Add(mBase, _c_F_gistbuildempty[0]))
				*(*int32)(unsafe.Add(mBase, _c_F_gistbuildempty[0])) = v63 - int32(1)
				F_UnlockReleaseBuffer(m, v19)
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return
				} else {
					m.G0 = v5 + int32(32)
					return
				}
			}
		}
	}
}
func F_gistinitpage(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	v2 = l1
	v3 = int32(_a_F_gistinitpage_0)
	v5 = int32(0)
	if v5|(l0&int32(3)|int32(1)) == v5 {
		v21 = l0 + v3
		v23 = l0 + int32(4)
		if base.Ui32(v23) < base.Ui32(v21) {
			v25 = v21
		} else {
			v25 = v23
		}
		v30 = (l0^int32(-1)+v25)&int32(-4) + int32(4)
		if v30 == int32(0) {
		} else {
			base.MemoryFill(m, l0, int32(0), v30)
		}
	} else {
		base.MemoryFill(m, l0, int32(0), v3)
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+10)) = int32(_a_F_gistinitpage_1)
	v44 = int32(_a_F_gistinitpage_2)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)) = uint16(v44)
	v50 = int32(_a_F_gistinitpage_3)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v50)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v50)
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	v54 = l0 + v53
	v55 = int32(_a_F_gistinitpage_4)
	*(*uint16)(unsafe.Add(mBase, uint32(v54)+14)) = uint16(v55)
	*(*uint16)(unsafe.Add(mBase, uint32(v54)+12)) = uint16(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v54)+8)) = int32(-1)
	return
}
func F_gistinsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
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
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_gistinsert[0]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l7)+136))
	if v11 == int32(0) {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l7)+140))
		*(*int32)(unsafe.Add(mBase, _c_F_gistinsert[0])) = v15
		v17 = F_initGISTstate(m, l0)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_gistinsert[0]))
			v27 = F_AllocSetContextCreateInternal(m, v22, int32(_a_F_gistinsert_0), int32(0), int32(_a_F_gistinsert_1), int32(_a_F_gistinsert_2))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v27
				*(*int32)(unsafe.Add(mBase, uint32(l7)+136)) = v17
				v31 = v17
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
				*(*int32)(unsafe.Add(mBase, _c_F_gistinsert[0])) = v33
				v36 = F_gistFormTuple(m, v31, l0, l1, l2, int32(1))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)))
					*(*uint16)(unsafe.Add(mBase, uint32(v36)+4)) = uint16(v38)
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
					*(*int32)(unsafe.Add(mBase, uint32(v36))) = v40
					v42 = int32(0)
					F_gistdoinsert(m, l0, v36, v42, v31, l4, v42)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_gistinsert[0])) = v10
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
						F_MemoryContextReset(m, v48)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							return int32(0)
						}
					}
				}
			}
		}
	} else {
		v31 = v11
		v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
		*(*int32)(unsafe.Add(mBase, _c_F_gistinsert[0])) = v33
		v36 = F_gistFormTuple(m, v31, l0, l1, l2, int32(1))
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return int32(0)
		} else {
			v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)))
			*(*uint16)(unsafe.Add(mBase, uint32(v36)+4)) = uint16(v38)
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
			*(*int32)(unsafe.Add(mBase, uint32(v36))) = v40
			v42 = int32(0)
			F_gistdoinsert(m, l0, v36, v42, v31, l4, v42)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_gistinsert[0])) = v10
				v48 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
				F_MemoryContextReset(m, v48)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					return int32(0)
				}
			}
		}
	}
}
func F_gistpenalty(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) float32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v38 int64
	_ = v38
	var v41 int32
	_ = v41
	var v42 float32
	_ = v42
	var v44 float32
	_ = v44
	var v47 float32
	_ = v47
	var v53 float32
	_ = v53
	var v56 float32
	_ = v56
	var v58 float32
	_ = v58
	var v60 float32
	_ = v60
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(0)
	v18 = l0 + l1*int32(28)
	if l3|l5 == int32(1) {
		v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+int32(3614)))))
		if v24 != 0 {
			if l5 != 0 {
				v56 = float32(0)
			} else {
				v56 = math.Float32frombits(uint32(0x7f800000))
			}
			if l3 != 0 {
				v58 = v56
			} else {
				v58 = math.Float32frombits(uint32(0x7f800000))
			}
			v60 = v58
			m.G0 = v12 + int32(16)
			return v60
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0+l1<<(uint(int32(2))%32))+uint32(_c_F_gistpenalty[0])))
			v38 = F_FunctionCall3Coll(m, v18+int32(3604), v32, base.I64_extend_i32_u(l2), base.I64_extend_i32_u(l4), base.I64_extend_i32_u(v12+int32(12)))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return float32(0)
			} else {
				v42 = float32(0)
				v44 = *(*float32)(unsafe.Add(mBase, uint32(v12)+12))
				if base.F32_lt(v44, v42) != 0 {
					v47 = v42
				} else {
					v47 = v44
				}
				if base.Ui32(int32(2139095040)) < base.Ui32(base.I32_reinterpret_f32(v44)&int32(2147483647)) {
					v53 = v42
				} else {
					v53 = v47
				}
				v60 = v53
				m.G0 = v12 + int32(16)
				return v60
			}
		}
	} else {
		v32 = *(*int32)(unsafe.Add(mBase, uint32(l0+l1<<(uint(int32(2))%32))+uint32(_c_F_gistpenalty[0])))
		v38 = F_FunctionCall3Coll(m, v18+int32(3604), v32, base.I64_extend_i32_u(l2), base.I64_extend_i32_u(l4), base.I64_extend_i32_u(v12+int32(12)))
		mBase = m.M
		v41 = m.ExcPending
		if v41 != 0 {
			return float32(0)
		} else {
			v42 = float32(0)
			v44 = *(*float32)(unsafe.Add(mBase, uint32(v12)+12))
			if base.F32_lt(v44, v42) != 0 {
				v47 = v42
			} else {
				v47 = v44
			}
			if base.Ui32(int32(2139095040)) < base.Ui32(base.I32_reinterpret_f32(v44)&int32(2147483647)) {
				v53 = v42
			} else {
				v53 = v47
			}
			v60 = v53
			m.G0 = v12 + int32(16)
			return v60
		}
	}
}
func F_gistproperty(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	v9 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	if l1 == v9 {
		v66 = v9
		m.G0 = v13 + int32(16)
		return v66
	} else {
		switch l2 - int32(6) {
		case 0:
			v21 = int64(8)
			v22 = int32(1)
			v24 = F_get_index_column_opclass(m, l0, l1)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				if v24 == int32(0) {
					v60 = v22
					*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v60)
					v66 = v22
					m.G0 = v13 + int32(16)
					return v66
				} else {
					v34 = F_get_opclass_opfamily_and_input_type(m, v24, v13+int32(12), v13+int32(8))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						if v34 == int32(0) {
							v60 = v22
							*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v60)
							v66 = v22
							m.G0 = v13 + int32(16)
							return v66
						} else {
							v39 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v13)+12)))
							v40 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v13)+8)))
							v41 = F_SearchSysCacheExists(m, int32(5), v39, v40, v40, v21)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v41)
								v44 = int32(0)
								if base.B2i32(l2 != int32(7))|v41 != 0 {
									v60 = v44
									*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v60)
									v66 = v22
									m.G0 = v13 + int32(16)
									return v66
								} else {
									v49 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v13)+12)))
									v50 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v13)+8)))
									v52 = F_SearchSysCacheExists(m, int32(5), v49, v50, v50, int64(3))
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int32(0)
									} else {
										v55 = v52 ^ int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v55)
										v60 = v44
										*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v60)
										v66 = v22
										m.G0 = v13 + int32(16)
										return v66
									}
								}
							}
						}
					}
				}
			}
		case 1:
			v21 = int64(9)
			v22 = int32(1)
			v24 = F_get_index_column_opclass(m, l0, l1)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				if v24 == int32(0) {
					v60 = v22
					*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v60)
					v66 = v22
					m.G0 = v13 + int32(16)
					return v66
				} else {
					v34 = F_get_opclass_opfamily_and_input_type(m, v24, v13+int32(12), v13+int32(8))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						if v34 == int32(0) {
							v60 = v22
							*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v60)
							v66 = v22
							m.G0 = v13 + int32(16)
							return v66
						} else {
							v39 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v13)+12)))
							v40 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v13)+8)))
							v41 = F_SearchSysCacheExists(m, int32(5), v39, v40, v40, v21)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v41)
								v44 = int32(0)
								if base.B2i32(l2 != int32(7))|v41 != 0 {
									v60 = v44
									*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v60)
									v66 = v22
									m.G0 = v13 + int32(16)
									return v66
								} else {
									v49 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v13)+12)))
									v50 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v13)+8)))
									v52 = F_SearchSysCacheExists(m, int32(5), v49, v50, v50, int64(3))
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int32(0)
									} else {
										v55 = v52 ^ int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v55)
										v60 = v44
										*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v60)
										v66 = v22
										m.G0 = v13 + int32(16)
										return v66
									}
								}
							}
						}
					}
				}
			}
		default:
			v66 = v9
			m.G0 = v13 + int32(16)
			return v66
		}
	}
}
func F_gistvacuumcleanup(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
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
	var v18 float64
	_ = v18
	var v19 float64
	_ = v19
	var v24 int32
	_ = v24
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v4 != 0 {
		v24 = l1
		return v24
	} else {
		if l1 == int32(0) {
			v8 = F_palloc0(m, int32(40))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				v12 = int32(0)
				F_gistvacuumscan(m, l0, v8, v12, v12)
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return int32(0)
				} else {
					v16 = v8
					v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)))
					if v17 != 0 {
						v24 = v16
					} else {
						v18 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
						v19 = *(*float64)(unsafe.Add(mBase, uint32(v16)+8))
						if base.F64_lt(v18, v19) == int32(0) {
							v24 = v16
						} else {
							*(*float64)(unsafe.Add(mBase, uint32(v16)+8)) = v18
							v24 = v16
						}
					}
					return v24
				}
			}
		} else {
			v16 = l1
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)))
			if v17 != 0 {
				v24 = v16
			} else {
				v18 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
				v19 = *(*float64)(unsafe.Add(mBase, uint32(v16)+8))
				if base.F64_lt(v18, v19) == int32(0) {
					v24 = v16
				} else {
					*(*float64)(unsafe.Add(mBase, uint32(v16)+8)) = v18
					v24 = v16
				}
			}
			return v24
		}
	}
}
func F_gseg_union(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v22 int64
	_ = v22
	var v29 int64
	_ = v29
	var v30 int64
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v43 int64
	_ = v43
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(v7)+8))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(12)
	if int32(2) <= v9 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = int32(1)
	v22 = v8
	goto L4
L2:
	;
	v43 = int64(0)
	goto L3
L3:
	;
	return v43
L4:
	;
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v7+int32(8)+v18*int32(24))))
	v30 = F_DirectFunctionCall2Coll(m, int32(_a_F_gseg_union_0), int32(0), v22, v29)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v43 = v30
	goto L3
L6:
	;
	return int64(0)
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(12)
	v37 = v18 + int32(1)
	if v37 != v9 {
		v18 = v37
		v22 = v30
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
}
func F_gtsquery_penalty(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v18 int32
	_ = v18
	var v41 int32
	_ = v41
	var v43 int64
	_ = v43
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
	v9 = v6 ^ v8
	v10 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
	v12 = int32(0)
	v13 = int64(0)
	for {
		v18 = int32(1)
		v41 = base.I32_wrap_i64(int64(base.Ui64(v9)>>(uint(v13)%64)))&v18 + v12 + base.I32_wrap_i64(int64(base.Ui64(v9)>>(uint(v13|int64(1))%64)))&v18 + base.I32_wrap_i64(int64(base.Ui64(v9)>>(uint(v13|int64(2))%64)))&v18 + base.I32_wrap_i64(int64(base.Ui64(v9)>>(uint(v13|int64(3))%64)))&v18
		v43 = v13 + int64(4)
		if v43 != int64(64) {
			v12 = v41
			v13 = v43
			continue
		} else {
			break
		}
		break
	}
	*(*float32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v10)))) = base.F32_convert_i32_s(v41)
	return v10
}
func F_gtsquery_union(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int64
	_ = v8
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int64
	_ = v31
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v58 int64
	_ = v58
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int64
	_ = v67
	var v72 int64
	_ = v72
	var v73 int64
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v86 int64
	_ = v86
	v2 = int32(0)
	v8 = int64(0)
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v12 <= v2 {
		v86 = v8
	} else {
		v16 = v12 & int32(3)
		v18 = v11 + int32(8)
		v19 = int32(0)
		if base.Ui32(int32(4)) <= base.Ui32(v12) {
			v24 = v19
			v29 = v2
			v31 = v8
			for {
				v35 = v18 + v24*int32(24)
				v36 = *(*int64)(unsafe.Add(mBase, uint32(v35)+72))
				v37 = *(*int64)(unsafe.Add(mBase, uint32(v35)+48))
				v38 = *(*int64)(unsafe.Add(mBase, uint32(v35)+24))
				v39 = *(*int64)(unsafe.Add(mBase, uint32(v35)))
				v43 = v36 | (v37 | (v38 | (v39 | v31)))
				v44 = int32(4)
				v45 = v24 + v44
				v47 = v29 + v44
				if v47 != v12&int32(2147483644) {
					v24 = v45
					v29 = v47
					v31 = v43
					continue
				} else {
					break
				}
				break
			}
			if v16 == int32(0) {
				v86 = v43
			} else {
				v51 = v45
				v58 = v43
				v60 = v51
				v66 = v2
				v67 = v58
				for {
					v72 = *(*int64)(unsafe.Add(mBase, uint32(v18+v60*int32(24))))
					v73 = v72 | v67
					v74 = int32(1)
					v77 = v66 + v74
					if v77 != v16 {
						v60 = v60 + v74
						v66 = v77
						v67 = v73
						continue
					} else {
						break
					}
					break
				}
				v86 = v73
			}
		} else {
			v51 = v19
			v58 = v8
			v60 = v51
			v66 = v2
			v67 = v58
			for {
				v72 = *(*int64)(unsafe.Add(mBase, uint32(v18+v60*int32(24))))
				v73 = v72 | v67
				v74 = int32(1)
				v77 = v66 + v74
				if v77 != v16 {
					v60 = v60 + v74
					v66 = v77
					v67 = v73
					continue
				} else {
					break
				}
				break
			}
			v86 = v73
		}
	}
	*(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v10)))) = int32(8)
	return v86
}
