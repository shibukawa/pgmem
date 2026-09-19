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
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_GetAdditionalLocalPinLimit[0]))
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_GetAdditionalLocalPinLimit[1]))
	return v2 - v4
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
	F_errfinish(m, int32(_a_F_GetCTEForRTE_1), int32(576), int32(_a_F_GetCTEForRTE_2))
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
	F_errfinish(m, int32(_a_F_GetCTEForRTE_1), int32(586), int32(_a_F_GetCTEForRTE_2))
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
			v15 = F_MemoryContextAlloc(m, v13, int32(184))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				base.MemoryCopy(m, v15, v8, int32(184))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+260)) = v15
				return v8
			}
		}
	} else {
		if l1 != 0 {
			v22 = F_palloc(m, int32(184))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
				base.MemoryCopy(m, v22, v24, int32(184))
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
	var v12 int32
	_ = v12
	var v16 int64
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int64
	_ = v23
	v3 = int32(_a_F_GetFlushRecPtr_0)
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_GetFlushRecPtr[0]))
	v5 = int64(0)
	v8 = base.AtomicRmwCmpxchg64(m, v4, int32(280), v5, v5)
	*(*int64)(unsafe.Add(mBase, _c_F_GetFlushRecPtr[1])) = v8
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_GetFlushRecPtr[0]))
	v16 = base.AtomicRmwCmpxchg64(m, v12, int32(272), v5, v5)
	*(*int64)(unsafe.Add(mBase, _c_F_GetFlushRecPtr[2])) = v16
	if l0 != 0 {
		v19 = *(*int32)(unsafe.Add(mBase, _c_F_GetFlushRecPtr[0]))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+308))
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v20
	} else {
	}
	v23 = *(*int64)(unsafe.Add(mBase, _c_F_GetFlushRecPtr[1]))
	return v23
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
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v32 int64
	_ = v32
	var v36 int32
	_ = v36
	var v40 int64
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int64
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int64
	_ = v72
	var v73 int32
	_ = v73
	var v77 int64
	_ = v77
	var v78 int32
	_ = v78
	var v81 int64
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int64
	_ = v90
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v13 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetLatestLSN[0])))
	if v13 == int32(1) {
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_GetLatestLSN[1]))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+316))
		v21 = base.B2i32(v19 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _c_F_GetLatestLSN[0])) = uint8(v21)
		v23 = v21
	} else {
		v23 = int32(0)
	}
	if v23 == int32(0) {
		v27 = int32(_a_F_GetLatestLSN_0)
		v28 = *(*int32)(unsafe.Add(mBase, _c_F_GetLatestLSN[1]))
		v29 = int64(0)
		v32 = base.AtomicRmwCmpxchg64(m, v28, int32(280), v29, v29)
		*(*int64)(unsafe.Add(mBase, _c_F_GetLatestLSN[2])) = v32
		v36 = *(*int32)(unsafe.Add(mBase, _c_F_GetLatestLSN[1]))
		v40 = base.AtomicRmwCmpxchg64(m, v36, int32(272), v29, v29)
		*(*int64)(unsafe.Add(mBase, _c_F_GetLatestLSN[3])) = v40
		if l0 != 0 {
			v43 = *(*int32)(unsafe.Add(mBase, _c_F_GetLatestLSN[1]))
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+308))
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v44
		} else {
		}
		v47 = *(*int64)(unsafe.Add(mBase, _c_F_GetLatestLSN[2]))
		v90 = v47
		m.G0 = v9 + int32(16)
		return v90
	} else {
		v49 = *(*int32)(unsafe.Add(mBase, _c_F_GetLatestLSN[1]))
		v52 = base.AtomicRmwXchg32(m, v49, int32(440), int32(1))
		if v52 != 0 {
			v54 = *(*int32)(unsafe.Add(mBase, _c_F_GetLatestLSN[1]))
			F_s_lock(m, v54+int32(440), int32(_a_F_GetLatestLSN_1), int32(_a_F_GetLatestLSN_2), int32(_a_F_GetLatestLSN_3))
			mBase = m.M
			v63 = m.ExcPending
			if v63 != 0 {
				return int64(0)
			} else {
				v65 = *(*int32)(unsafe.Add(mBase, _c_F_GetLatestLSN[1]))
				v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+308))
				v67 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v65)+440)), uint32(v67))
				if v66 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v66
					v72 = F_GetXLogReplayRecPtr(m, int32(0))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int64(0)
					} else {
						v90 = v72
						m.G0 = v9 + int32(16)
						return v90
					}
				} else {
					v77 = F_GetWalRcvFlushRecPtr(m, int32(0), v9+int32(12))
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return int64(0)
					} else {
						v81 = F_GetXLogReplayRecPtr(m, v9+int32(8))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return int64(0)
						} else {
							if base.Ui64(v81) < base.Ui64(v77) {
								v84 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v84
								v90 = v77
							} else {
								v86 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v86
								v90 = v81
							}
							m.G0 = v9 + int32(16)
							return v90
						}
					}
				}
			}
		} else {
			v65 = *(*int32)(unsafe.Add(mBase, _c_F_GetLatestLSN[1]))
			v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+308))
			v67 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v65)+440)), uint32(v67))
			if v66 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v66
				v72 = F_GetXLogReplayRecPtr(m, int32(0))
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return int64(0)
				} else {
					v90 = v72
					m.G0 = v9 + int32(16)
					return v90
				}
			} else {
				v77 = F_GetWalRcvFlushRecPtr(m, int32(0), v9+int32(12))
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return int64(0)
				} else {
					v81 = F_GetXLogReplayRecPtr(m, v9+int32(8))
					mBase = m.M
					v82 = m.ExcPending
					if v82 != 0 {
						return int64(0)
					} else {
						if base.Ui64(v81) < base.Ui64(v77) {
							v84 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v84
							v90 = v77
						} else {
							v86 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v86
							v90 = v81
						}
						m.G0 = v9 + int32(16)
						return v90
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
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
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
	var v158 int64
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v182 int64
	_ = v182
	var v188 int32
	_ = v188
	var v190 int64
	_ = v190
	var v196 int32
	_ = v196
	var v199 int64
	_ = v199
	var v204 int32
	_ = v204
	var v206 int64
	_ = v206
	var v212 int32
	_ = v212
	var v215 int64
	_ = v215
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v242 int32
	_ = v242
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v421 int32
	_ = v421
	var v429 int32
	_ = v429
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v504 int32
	_ = v504
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	var v543 int32
	_ = v543
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
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
	v558 = m.ExcPending
	if v558 != 0 {
		goto L13
	} else {
		goto L90
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
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
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[2]))
	v75 = F_get_hash_value(m, v74, l0)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L13
	} else {
		goto L16
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[3])) = v71
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
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[4]))
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[5]))
	v69 = F_palloc0(m, (v61+v63)<<(uint(int32(3))%32)+int32(8))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L13
	} else {
		goto L15
	}
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[6]))
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[4]))
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[5]))
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
	v71 = v56
	goto L8
L15:
	;
	v71 = v69
	goto L8
L16:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[7]))
	v85 = v78 + v75&int32(15)<<(uint(int32(7))%32) + int32(_a_F_GetLockConflicts_0)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v86+l1<<(uint(int32(2))%32))))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	if v91 != int32(1) {
		v288 = v4
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v304 = F_LWLockAcquire(m, v85, int32(1))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L13
	} else {
		goto L47
	}
L18:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
	if v94|base.B2i32(base.Ui32(l1) < base.Ui32(int32(5))) != 0 {
		v288 = v4
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v98 == int32(0) {
		v288 = v4
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[8]))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+16))
	if v103 == int32(0) {
		v288 = v4
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[9]))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v113 = (v107 - int32(1)) & (v110 * int32(_a_F_GetLockConflicts_1))
	v116 = int32(3)
	v123 = v102
	v126 = v4
	v130 = v4
	goto L22
L22:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	v144 = v141 + v130*int32(640)
	v146 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[10]))
	if v144 != v146 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v288 = v263
	goto L17
L24:
	;
	v149 = v144 + int32(584)
	v151 = F_LWLockAcquire(m, v149, int32(1))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L13
	} else {
		goto L27
	}
L25:
	;
	v263 = v126
	goto L26
L26:
	;
	v279 = v130 + int32(1)
	v281 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[8]))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)+16))
	if base.Ui32(v279) < base.Ui32(v282) {
		v123 = v281
		v126 = v263
		v130 = v279
		goto L22
	} else {
		goto L46
	}
L27:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v144)+60))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v153 != v154 {
		v242 = v126
		goto L28
	} else {
		goto L29
	}
L28:
	;
	F_LWLockRelease(m, v149)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L13
	} else {
		goto L45
	}
L29:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v144)+600))
	v158 = *(*int64)(unsafe.Add(mBase, uint32(v156+v113<<(uint(v116)%32))))
	if v158 == int64(0) {
		v242 = v126
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v161 = v156 + v113&int32(268435455)<<(uint(v116)%32)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v144)+604))
	v163 = v162 + v113<<(uint(int32(6))%32)
	v182 = int64(0)
	goto L31
L31:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v163+base.I32_wrap_i64(v182)<<(uint(int32(2))%32))))
	if v188 == v110 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	if v218<<(uint(int32(1))%32)&v90 == int32(0) {
		v242 = v126
		goto L28
	} else {
		goto L43
	}
L33:
	;
	goto L32
L34:
	;
	v190 = *(*int64)(unsafe.Add(mBase, uint32(v161)))
	v196 = base.I32_wrap_i64(int64(base.Ui64(v190)>>(uint(v182*int64(3))%64))) & int32(7)
	if v196 != 0 {
		v218 = v196
		goto L33
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v199 = v182 | int64(1)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v163+base.I32_wrap_i64(v199)<<(uint(int32(2))%32))))
	if v204 == v110 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L36
L38:
	;
	v206 = *(*int64)(unsafe.Add(mBase, uint32(v161)))
	v212 = base.I32_wrap_i64(int64(base.Ui64(v206)>>(uint(v199*int64(3))%64))) & int32(7)
	if v212 != 0 {
		v218 = v212
		goto L33
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v215 = v182 + int64(2)
	if v215 != int64(16) {
		v182 = v215
		goto L31
	} else {
		goto L42
	}
L41:
	;
	goto L40
L42:
	;
	v242 = v126
	goto L28
L43:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v144)+56))
	if v225 == int32(0) {
		v242 = v126
		goto L28
	} else {
		goto L44
	}
L44:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v144)+52))
	v230 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[3]))
	v233 = v230 + v126<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v233)+4)) = v225
	*(*int32)(unsafe.Add(mBase, uint32(v233))) = v228
	v242 = v126 + int32(1)
	goto L28
L45:
	;
	v263 = v242
	goto L26
L46:
	;
	goto L23
L47:
	;
	v307 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[2]))
	v308 = int32(0)
	v310 = F_hash_search_with_hash_value(m, v307, l0, v75, v308, v308)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L13
	} else {
		goto L51
	}
L48:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L13
	} else {
		goto L84
	}
L49:
	;
	m.G0 = v22 + int32(32)
	return v504
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v487
	v504 = v484
	goto L49
L51:
	;
	if v310 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	F_LWLockRelease(m, v85)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L13
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v310)+28))
	if v323 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	v317 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[3]))
	*(*int64)(unsafe.Add(mBase, uint32(v317+v288<<(uint(int32(3))%32)))) = int64(4294967295)
	if l2 != 0 {
		v484 = v317
		v487 = v288
		goto L50
	} else {
		goto L56
	}
L56:
	;
	v504 = v317
	goto L49
L57:
	;
	F_LWLockRelease(m, v85)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L13
	} else {
		goto L81
	}
L58:
	;
	v450 = v288
	goto L57
L59:
	;
	goto L60
L60:
	;
	v327 = v310 + int32(24)
	if v327 == v323 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v450 = v288
	goto L57
L62:
	;
	goto L63
L63:
	;
	v330 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[3]))
	v333 = v323
	v336 = v288
	goto L64
L64:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v333-int32(8))))
	if v354&v90 == int32(0) {
		v429 = v336
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v450 = v429
	goto L57
L66:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v333)+4))
	if v445 != v327 {
		v333 = v445
		v336 = v429
		goto L64
	} else {
		goto L80
	}
L67:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v333-int32(16))))
	v362 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[10]))
	if v360 == v362 {
		v429 = v336
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v360)+56))
	if v364 == int32(0) {
		v429 = v336
		goto L66
	} else {
		goto L69
	}
L69:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v360)+52))
	v368 = int32(0)
	if base.B2i32(v288 <= int32(0)) == v368 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v372 = v368
	goto L73
L71:
	;
	goto L72
L72:
	;
	v421 = v330 + v336<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v421)+4)) = v364
	*(*int32)(unsafe.Add(mBase, uint32(v421))) = v367
	v429 = v336 + int32(1)
	goto L66
L73:
	;
	v392 = v330 + v372<<(uint(int32(3))%32)
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v392)))
	if v367 == v393 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	goto L72
L75:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v392)+4))
	if v395 == v364 {
		v429 = v336
		goto L66
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v398 = v372 + int32(1)
	if v398 != v288 {
		v372 = v398
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
	v469 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[4]))
	v471 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[5]))
	if v469+v471 < v450 {
		goto L48
	} else {
		goto L82
	}
L82:
	;
	v475 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[3]))
	*(*int64)(unsafe.Add(mBase, uint32(v475+v450<<(uint(int32(3))%32)))) = int64(4294967295)
	if l2 == int32(0) {
		v504 = v475
		goto L49
	} else {
		goto L83
	}
L83:
	;
	v484 = v475
	v487 = v450
	goto L50
L84:
	;
	F_errmsg_internal(m, int32(_a_F_GetLockConflicts_2), int32(0))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L13
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_GetLockConflicts_3), int32(3233), int32(_a_F_GetLockConflicts_4))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
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
	v549 = m.ExcPending
	if v549 != 0 {
		goto L13
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(_a_F_GetLockConflicts_3), int32(3055), int32(_a_F_GetLockConflicts_4))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
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
	v562 = m.ExcPending
	if v562 != 0 {
		goto L13
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(_a_F_GetLockConflicts_3), int32(3052), int32(_a_F_GetLockConflicts_4))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
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
	var v28 int32
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
	v8 = v6 - int32(96)
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
	m.G0 = v8 + int32(96)
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
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
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
	F_ScanKeyInit(m, v8+int32(48), int32(4), int32(3), int32(61), int32(102))
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
	var v10 int32
	_ = v10
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v27 int64
	_ = v27
	var v31 int64
	_ = v31
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_GetRedoRecPtr[0]))
	v8 = base.AtomicRmwXchg32(m, v5, int32(440), int32(1))
	if v8 != 0 {
		v10 = *(*int32)(unsafe.Add(mBase, _c_F_GetRedoRecPtr[0]))
		F_s_lock(m, v10+int32(440), int32(_a_F_GetRedoRecPtr_0), int32(_a_F_GetRedoRecPtr_1), int32(_a_F_GetRedoRecPtr_2))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, _c_F_GetRedoRecPtr[0]))
			v22 = *(*int64)(unsafe.Add(mBase, uint32(v21)+200))
			v23 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v21)+440)), uint32(v23))
			v27 = *(*int64)(unsafe.Add(mBase, _c_F_GetRedoRecPtr[1]))
			if base.Ui64(v27) < base.Ui64(v22) {
				*(*int64)(unsafe.Add(mBase, _c_F_GetRedoRecPtr[1])) = v22
				v31 = v22
			} else {
				v31 = v27
			}
			return v31
		}
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, _c_F_GetRedoRecPtr[0]))
		v22 = *(*int64)(unsafe.Add(mBase, uint32(v21)+200))
		v23 = int32(0)
		atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v21)+440)), uint32(v23))
		v27 = *(*int64)(unsafe.Add(mBase, _c_F_GetRedoRecPtr[1]))
		if base.Ui64(v27) < base.Ui64(v22) {
			*(*int64)(unsafe.Add(mBase, _c_F_GetRedoRecPtr[1])) = v22
			v31 = v22
		} else {
			v31 = v27
		}
		return v31
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
	var v143 int32
	_ = v143
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
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
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
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
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v50)+92))
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
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v50)+100))
	v82 = v80
	goto L16
L15:
	;
	v82 = int32(0)
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69)+20)) = v82
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v50)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v69)+24)) = v84
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v50)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v69)+28)) = v86
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v50)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v69)+40)) = v88
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v38-int32(12))))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+44))
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
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v176 - v177
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v180 - v181
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
	if base.B2i32(v135 == v137)|base.B2i32(l0 == v137) != 0 {
		goto L25
	} else {
		goto L29
	}
L29:
	;
	v143 = v137
	goto L30
L30:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v143)+44))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v153 + int32(1)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v157+v153<<(uint(int32(2))%32)))) = v152
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	if v162 == v135 {
		goto L25
	} else {
		goto L32
	}
L31:
	;
	goto L25
L32:
	;
	if l0 != v162 {
		v143 = v162
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
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
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
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v10 = F_OidFunctionCall1Coll(m, l0, v2, v2)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 != 0 {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			if v14 == int32(440) {
				m.G0 = v6 + int32(16)
				return v10
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
					F_errmsg_internal(m, int32(_a_F_GetTsmRoutine_0), v6)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_GetTsmRoutine_1), int32(37), int32(_a_F_GetTsmRoutine_2))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
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
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg_internal(m, int32(_a_F_GetTsmRoutine_0), v6)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_GetTsmRoutine_1), int32(37), int32(_a_F_GetTsmRoutine_2))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v209 int32
	_ = v209
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v434 int32
	_ = v434
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v462 int32
	_ = v462
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v473 int64
	_ = v473
	var v474 int64
	_ = v474
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v554 int64
	_ = v554
	var v556 int64
	_ = v556
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v595 int32
	_ = v595
	var v596 int64
	_ = v596
	var v600 int64
	_ = v600
	var v602 int32
	_ = v602
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[0]))
	F_ResourceOwnerEnlarge(m, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	goto L4
L4:
	;
	v34 = v13 + int32(4)
	v36 = v13 + int32(3)
	v37 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v36))) = uint8(v37)
	if l0 == v37 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	m.G0 = v13 + int32(32)
	return v380
L6:
	;
	goto L5
L7:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v355)+24))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v355)+20))
	v358 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v355)+24)) = (v356 + v358) & int32(-4194305)
	v364 = int32(_a_F_GetVictimBuffer_0)
	v365 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[1])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v365)+4)) = v358
	v372 = v357 + v358
	*(*int32)(unsafe.Add(mBase, uint32(v365))) = v372
	v375 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[0]))
	F_ResourceOwnerRemember(m, v375, v372, int32(_a_F_GetVictimBuffer_1))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L86
	}
L8:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[2]))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+24))
	if v81 != int32(-1) {
		goto L19
	} else {
		goto L20
	}
L9:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v43 = v41 + int32(1)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v43 < v45 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v47 = v43
	goto L12
L11:
	;
	v47 = int32(0)
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v47
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0+v47<<(uint(int32(2))%32))+12))
	if v52 == int32(0) {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[3]))
	v59 = v56 + v52<<(uint(int32(6))%32)
	v61 = v59 + int32(-64)
	v62 = F_LockBufHdr(m, v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v62&int32(_a_F_GetVictimBuffer_2) != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59-int32(40)))) = v62 & int32(-4194305)
	goto L8
L16:
	;
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v62
	if v61 == int32(0) {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	v74 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v36))) = uint8(v74)
	v355 = v61
	goto L7
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+24)) = int32(-1)
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[4]))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	v93 = v88 + v81*int32(640) + int32(20)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	if v94 != 0 {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v139 = v80
	goto L21
L21:
	;
	v142 = base.AtomicRmwAdd32(m, v139, int32(20), int32(1))
	v144 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[2]))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+8))
	if v145 < int32(0) {
		v209 = v144
		goto L36
	} else {
		goto L37
	}
L22:
	;
	v138 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[2]))
	v139 = v138
	goto L21
L23:
	;
	goto L22
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = int32(1)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	if v97 == int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	if v100 == int32(0) {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[5]))
	if v104 == v100 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v106 = m.G0
	v108 = v106 - int32(16)
	m.G0 = v108
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[6]))
	if v111 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v134 = F_pgmem_kill(m, v100, int32(23))
	mBase = m.M
	goto L23
L30:
	;
	m.G0 = v108 + int32(16)
	goto L22
L31:
	;
	v114 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v108)+15)) = uint8(v114)
	goto L32
L32:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[7]))
	v122 = F_write(m, v118, v108+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v122 {
		goto L30
	} else {
		goto L34
	}
L33:
	;
	goto L30
L34:
	;
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[8]))
	if v126 == int32(27) {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v217 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[9]))
	v221 = v209
	v225 = v217
	goto L54
L37:
	;
	v151 = v144
	goto L38
L38:
	;
	v160 = base.AtomicRmwXchg32(m, v151, int32(0), int32(1))
	if v160 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v162 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[2]))
	F_s_lock(m, v162, int32(_a_F_GetVictimBuffer_3), int32(273), int32(_a_F_GetVictimBuffer_4))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[2]))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)+8))
	if v170 < int32(0) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L42
L44:
	;
	v173 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v169))), uint32(v173))
	v209 = v169
	goto L36
L45:
	;
	goto L46
L46:
	;
	v177 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[3]))
	v180 = v177 + v170<<(uint(int32(6))%32)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v169)+8)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v180)+32)) = int32(-2)
	v185 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v169))), uint32(v185))
	v188 = F_LockBufHdr(m, v180)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	if v188&int32(_a_F_GetVictimBuffer_5) != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v180)+24)) = v188 & int32(-4194305)
	v196 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[2]))
	v151 = v196
	goto L38
L49:
	;
	if l0 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v180)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0+v197<<(uint(int32(2))%32))+12)) = v201 + int32(1)
	goto L53
L52:
	;
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v188
	v355 = v180
	goto L7
L54:
	;
	v230 = base.AtomicRmwAdd32(m, v221, int32(4), int32(1))
	v232 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[9]))
	if base.Ui32(v230) < base.Ui32(v232) {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v296)+24)) = v297 & int32(-4194305)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L83
	}
L56:
	;
	v293 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[3]))
	v296 = v293 + v287<<(uint(int32(6))%32)
	v297 = F_LockBufHdr(m, v296)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L72
	}
L57:
	;
	v287 = v230
	goto L56
L58:
	;
	goto L59
L59:
	;
	v234 = base.I32_rem_u_s(v230, v232)
	if v234 != 0 {
		v287 = v234
		goto L56
	} else {
		goto L60
	}
L60:
	;
	v238 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[2]))
	v242 = v230 + int32(1)
	v244 = v238
	goto L61
L61:
	;
	v251 = base.AtomicRmwXchg32(m, v244, int32(0), int32(1))
	if v251 != 0 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v272 = int32(0)
	v274 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[2]))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v274)+16)) = v275 + int32(1)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v274))), uint32(v272))
	v287 = v272
	goto L56
L63:
	;
	v253 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[2]))
	F_s_lock(m, v253, int32(_a_F_GetVictimBuffer_3), int32(151), int32(_a_F_GetVictimBuffer_6))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v260 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[2]))
	v262 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[9]))
	v263 = base.I32_rem_u_s(v242, v262)
	v265 = base.AtomicRmwCmpxchg32(m, v260, int32(4), v242, v263)
	if v242 != v265 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	goto L65
L67:
	;
	v268 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[2]))
	v269 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v268))), uint32(v269))
	v242 = v265
	v244 = v268
	goto L61
L68:
	;
	goto L69
L69:
	;
	goto L62
L70:
	;
	goto L55
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v296)+24)) = v322 & int32(-4194305)
	v328 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[2]))
	v221 = v328
	v225 = v323
	goto L54
L72:
	;
	if v297&int32(_a_F_GetVictimBuffer_7) == int32(0) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	if v297&int32(_a_F_GetVictimBuffer_8) != 0 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L75
L75:
	;
	v319 = v225 - int32(1)
	if v319 == int32(0) {
		goto L70
	} else {
		goto L82
	}
L76:
	;
	v308 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[9]))
	v322 = v297 - int32(_a_F_GetVictimBuffer_9)
	v323 = v308
	goto L71
L77:
	;
	goto L78
L78:
	;
	if l0 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v296)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0+v309<<(uint(int32(2))%32))+12)) = v313 + int32(1)
	goto L81
L80:
	;
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v297
	v355 = v296
	goto L7
L82:
	;
	v322 = v297
	v323 = v319
	goto L71
L83:
	;
	F_errmsg_internal(m, int32(_a_F_GetVictimBuffer_10), int32(0))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_GetVictimBuffer_3), int32(353), int32(_a_F_GetVictimBuffer_4))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	v380 = v357 + int32(1)
	F_CheckBufferIsPinnedOnce(m, v380)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+6)))
	if v383&int32(128) == int32(0) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v630 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[0]))
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v355)+20))
	F_ResourceOwnerForget(m, v630, v631+int32(1), int32(_a_F_GetVictimBuffer_1))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L1
	} else {
		goto L140
	}
L89:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v579&int32(16777216) != 0 {
		goto L130
	} else {
		goto L131
	}
L90:
	;
	v389 = v355 + int32(48)
	v391 = F_LWLockConditionalAcquire(m, v389, int32(1))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	if v391 == int32(0) {
		goto L88
	} else {
		goto L92
	}
L92:
	;
	if l0 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	F_FlushBuffer(m, v355, int32(0), l1)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L1
	} else {
		goto L121
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = int32(_a_F_GetVictimBuffer_11)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = int32(_a_F_GetVictimBuffer_12)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = int32(_a_F_GetVictimBuffer_13)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = int64(0)
	v407 = int32(_a_F_GetVictimBuffer_14)
	v409 = base.AtomicRmwOr32(m, v355, int32(24), v407)
	if v409&v407 != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	goto L98
L96:
	;
	v434 = v409
	goto L97
L97:
	;
	v444 = int32(_a_F_GetVictimBuffer_15)
	v445 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[10]))
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(8))+8))
	if v447 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L98:
	;
	F_perform_spin_delay(m, v13+int32(8))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L100
	}
L99:
	;
	v434 = v428
	goto L97
L100:
	;
	v426 = int32(_a_F_GetVictimBuffer_14)
	v428 = base.AtomicRmwOr32(m, v355, int32(24), v426)
	if v428&v426 != 0 {
		goto L98
	} else {
		goto L101
	}
L101:
	;
	goto L99
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v434 | int32(_a_F_GetVictimBuffer_14)
	v468 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[11]))
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v355)+20))
	v472 = v468 + v469<<(uint(int32(13))%32)
	v473 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v472)+4)))
	v474 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v472))))
	*(*int32)(unsafe.Add(mBase, uint32(v355)+24)) = v434 & int32(-4194305)
	v481 = F_XLogNeedsFlush(m, v473|v474<<(uint(int64(32))%64))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L113
	}
L103:
	;
	goto L102
L104:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[10])) = v462
	goto L103
L105:
	;
	if int32(999) < v445 {
		goto L103
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	if v445 < int32(11) {
		goto L103
	} else {
		goto L112
	}
L108:
	;
	v452 = int32(900)
	if v452 <= v445 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v455 = v452
	goto L111
L110:
	;
	v455 = v445
	goto L111
L111:
	;
	v462 = v455 + int32(100)
	goto L104
L112:
	;
	v462 = v445 - int32(1)
	goto L104
L113:
	;
	if v481 == int32(0) {
		goto L93
	} else {
		goto L114
	}
L114:
	;
	v485 = int32(0)
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+3)))
	if v486 == v485 {
		v507 = v485
		goto L115
	} else {
		goto L116
	}
L115:
	;
	if v507 == int32(0) {
		goto L93
	} else {
		goto L119
	}
L116:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v489 != int32(1) {
		v507 = v485
		goto L115
	} else {
		goto L117
	}
L117:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v497 = l0 + v492<<(uint(int32(2))%32) + int32(12)
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v497)))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v355)+20))
	if v498 != v499+int32(1) {
		v507 = v485
		goto L115
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v497))) = int32(0)
	v507 = int32(1)
	goto L115
L119:
	;
	F_LWLockRelease(m, v389)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	goto L88
L121:
	;
	F_LWLockRelease(m, v389)
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	v528 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetVictimBuffer[12])))
	if v528&int32(1) != 0 {
		goto L89
	} else {
		goto L123
	}
L123:
	;
	v532 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetVictimBuffer[13])))
	if v532&int32(1) == int32(0) {
		goto L89
	} else {
		goto L124
	}
L124:
	;
	v538 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[14]))
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v538)))
	if int32(0) < v539 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v542 = int32(_a_F_GetVictimBuffer_16)
	v544 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[15]))
	*(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[15])) = v544 + int32(1)
	v549 = v544 * int32(20)
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v355)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v549)+uint32(_c_F_GetVictimBuffer[16]))) = v552
	v554 = *(*int64)(unsafe.Add(mBase, uint32(v355)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v549)+uint32(_c_F_GetVictimBuffer[17]))) = v554
	v556 = *(*int64)(unsafe.Add(mBase, uint32(v355)))
	*(*int64)(unsafe.Add(mBase, uint32(v549)+uint32(_c_F_GetVictimBuffer[18]))) = v556
	v559 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[14]))
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v559)))
	v562 = v560
	goto L127
L126:
	;
	v562 = v539
	goto L127
L127:
	;
	v564 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[15]))
	if v564 < v562 {
		goto L89
	} else {
		goto L128
	}
L128:
	;
	F_IssuePendingWritebacks(m, int32(_a_F_GetVictimBuffer_17), l1)
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	goto L89
L130:
	;
	v582 = int32(0)
	v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+3)))
	if v585 != 0 {
		goto L133
	} else {
		goto L134
	}
L131:
	;
	v612 = v579
	goto L132
L132:
	;
	if v612&int32(33554432) == int32(0) {
		goto L6
	} else {
		goto L137
	}
L133:
	;
	v586 = int32(3)
	goto L135
L134:
	;
	v586 = v582
	goto L135
L135:
	;
	v595 = int32(0) + l1<<(uint(int32(6))%32) + v586<<(uint(int32(3))%32)
	v596 = *(*int64)(unsafe.Add(mBase, uint32(v595)+uint32(_c_F_GetVictimBuffer[19])))
	*(*int64)(unsafe.Add(mBase, uint32(v595)+uint32(_c_F_GetVictimBuffer[19]))) = v596 + int64(1)
	v600 = *(*int64)(unsafe.Add(mBase, uint32(v595)+uint32(_c_F_GetVictimBuffer[20])))
	*(*int64)(unsafe.Add(mBase, uint32(v595)+uint32(_c_F_GetVictimBuffer[20]))) = v600
	v602 = int32(1)
	F_pgstat_count_backend_io_op(m, v582, l1, v586, v602, int64(0))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, _c_F_GetVictimBuffer[21])) = uint8(v602)
	*(*uint8)(unsafe.Add(mBase, _c_F_GetVictimBuffer[22])) = uint8(v602)
	goto L136
L136:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v612 = v611
	goto L132
L137:
	;
	v617 = F_InvalidateVictimBuffer(m, v355)
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	if v617 != 0 {
		goto L6
	} else {
		goto L139
	}
L139:
	;
	goto L88
L140:
	;
	F_UnpinBufferNoOwner(m, v355)
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v52 int32
	_ = v52
	var v53 int64
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int64
	_ = v66
	var v68 int32
	_ = v68
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	if l0 == v3 {
		v41 = v3
		v45 = v41
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+117)))
		if v15 != 0 {
			v41 = v3
			v45 = v41
		} else {
			v16 = F_RecoveryInProgress(m)
			mBase = m.M
			if v16 != 0 {
				v41 = v3
				v45 = v41
			} else {
				v17 = int32(1)
				v18 = F_IsCatalogRelation(m, l0)
				mBase = m.M
				if v18 != 0 {
					v41 = v17
					v45 = v41
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, _c_F_GlobalVisCheckRemovableFullXid[0]))
					if v20 < int32(2) {
						v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
						if v37 != 0 {
							v41 = int32(3)
							v45 = v41
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							if v38 != 0 {
								v41 = int32(3)
								v45 = v41
							} else {
								v45 = int32(2)
							}
						}
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+118)))
						if v24 != int32(112) {
							v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
							if v37 != 0 {
								v41 = int32(3)
								v45 = v41
							} else {
								v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								if v38 != 0 {
									v41 = int32(3)
									v45 = v41
								} else {
									v45 = int32(2)
								}
							}
						} else {
							v27 = F_IsCatalogRelation(m, l0)
							mBase = m.M
							if v27 != 0 {
								v41 = v17
								v45 = v41
							} else {
								v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
								if v28 == int32(0) {
									v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
									if v37 != 0 {
										v41 = int32(3)
										v45 = v41
									} else {
										v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
										if v38 != 0 {
											v41 = int32(3)
											v45 = v41
										} else {
											v45 = int32(2)
										}
									}
								} else {
									v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+119)))
									switch v32 - int32(109) {
									case 0, 5:
										v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+104)))
										if v35 != 0 {
											v41 = v17
											v45 = v41
										} else {
											v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
											if v37 != 0 {
												v41 = int32(3)
												v45 = v41
											} else {
												v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
												if v38 != 0 {
													v41 = int32(3)
													v45 = v41
												} else {
													v45 = int32(2)
												}
											}
										}
									default:
										v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
										if v37 != 0 {
											v41 = int32(3)
											v45 = v41
										} else {
											v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
											if v38 != 0 {
												v41 = int32(3)
												v45 = v41
											} else {
												v45 = int32(2)
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
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v45<<(uint(int32(2))%32))+uint32(_c_F_GlobalVisCheckRemovableFullXid[1])))
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v48)+8))
	if base.Ui64(l1) < base.Ui64(v49) {
		v68 = int32(1)
		m.G0 = v8 + int32(48)
		return v68
	} else {
		v52 = int32(0)
		v53 = *(*int64)(unsafe.Add(mBase, uint32(v48)))
		if base.Ui64(v53) <= base.Ui64(l1) {
			v68 = v52
			m.G0 = v8 + int32(48)
			return v68
		} else {
			v56 = *(*int32)(unsafe.Add(mBase, _c_F_GlobalVisCheckRemovableFullXid[2]))
			if v56 != 0 {
				v58 = *(*int32)(unsafe.Add(mBase, _c_F_GlobalVisCheckRemovableFullXid[3]))
				if v58 == v56 {
					v68 = v52
					m.G0 = v8 + int32(48)
					return v68
				} else {
					F_ComputeXidHorizons(m, v8+int32(8))
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int32(0)
					} else {
						v66 = *(*int64)(unsafe.Add(mBase, uint32(v48)+8))
						v68 = base.B2i32(base.Ui64(l1) < base.Ui64(v66))
						m.G0 = v8 + int32(48)
						return v68
					}
				}
			} else {
				F_ComputeXidHorizons(m, v8+int32(8))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return int32(0)
				} else {
					v66 = *(*int64)(unsafe.Add(mBase, uint32(v48)+8))
					v68 = base.B2i32(base.Ui64(l1) < base.Ui64(v66))
					m.G0 = v8 + int32(48)
					return v68
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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	v2 = int32(0)
	if l0 == v2 {
		v33 = v2
		v37 = v33
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+117)))
		if v7 != 0 {
			v33 = v2
			v37 = v33
		} else {
			v8 = F_RecoveryInProgress(m)
			mBase = m.M
			if v8 != 0 {
				v33 = v2
				v37 = v33
			} else {
				v9 = int32(1)
				v10 = F_IsCatalogRelation(m, l0)
				mBase = m.M
				if v10 != 0 {
					v33 = v9
					v37 = v33
				} else {
					v12 = *(*int32)(unsafe.Add(mBase, _c_F_GlobalVisTestFor[0]))
					if v12 < int32(2) {
						v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
						if v29 != 0 {
							v33 = int32(3)
							v37 = v33
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							if v30 != 0 {
								v33 = int32(3)
								v37 = v33
							} else {
								v37 = int32(2)
							}
						}
					} else {
						v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+118)))
						if v16 != int32(112) {
							v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
							if v29 != 0 {
								v33 = int32(3)
								v37 = v33
							} else {
								v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								if v30 != 0 {
									v33 = int32(3)
									v37 = v33
								} else {
									v37 = int32(2)
								}
							}
						} else {
							v19 = F_IsCatalogRelation(m, l0)
							mBase = m.M
							if v19 != 0 {
								v33 = v9
								v37 = v33
							} else {
								v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
								if v20 == int32(0) {
									v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
									if v29 != 0 {
										v33 = int32(3)
										v37 = v33
									} else {
										v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
										if v30 != 0 {
											v33 = int32(3)
											v37 = v33
										} else {
											v37 = int32(2)
										}
									}
								} else {
									v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+119)))
									switch v24 - int32(109) {
									case 0, 5:
										v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+104)))
										if v27 != 0 {
											v33 = v9
											v37 = v33
										} else {
											v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
											if v29 != 0 {
												v33 = int32(3)
												v37 = v33
											} else {
												v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
												if v30 != 0 {
													v33 = int32(3)
													v37 = v33
												} else {
													v37 = int32(2)
												}
											}
										}
									default:
										v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
										if v29 != 0 {
											v33 = int32(3)
											v37 = v33
										} else {
											v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
											if v30 != 0 {
												v33 = int32(3)
												v37 = v33
											} else {
												v37 = int32(2)
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
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37<<(uint(int32(2))%32))+uint32(_c_F_GlobalVisTestFor[1])))
	return v40
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
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 float32
	_ = v32
	var v40 float32
	_ = v40
	var v47 int32
	_ = v47
	var v55 float64
	_ = v55
	var v56 float64
	_ = v56
	var v57 float64
	_ = v57
	var v59 float64
	_ = v59
	var v69 float64
	_ = v69
	var v70 float64
	_ = v70
	var v71 float64
	_ = v71
	var v73 float64
	_ = v73
	var v84 float64
	_ = v84
	var v86 float64
	_ = v86
	var v88 float64
	_ = v88
	var v89 float64
	_ = v89
	var v105 float64
	_ = v105
	var v107 float64
	_ = v107
	var v112 float64
	_ = v112
	var v123 float32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 float32
	_ = v127
	var v130 float32
	_ = v130
	var v135 float32
	_ = v135
	var v136 float32
	_ = v136
	var v137 float32
	_ = v137
	var v140 float32
	_ = v140
	var v145 float64
	_ = v145
	var v153 int32
	_ = v153
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v20 = base.I32_div_s(v16+int32(1), int32(2))
	if l3 < v20 {
		v23 = base.I32_div_s(v16, int32(2))
		if l5 < v23 {
			v25 = l5
		} else {
			v25 = v23
		}
		v26 = v25
	} else {
		v26 = l3
	}
	v29 = v16 - v26
	if v26 < v29 {
		v31 = v26
	} else {
		v31 = v29
	}
	v32 = base.F32_convert_i32_s(v31)
	if base.B2i32(v16 == int32(0))&base.B2i32(base.Ui32(base.I32_reinterpret_f32(v32)&int32(2147483647)) <= base.Ui32(int32(2139095040))) != 0 {
		F_float_zero_divide_error(m)
		mBase = m.M
		v172 = m.ExcPending
		if v172 != 0 {
			return
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v40 = base.F32_div(v32, base.F32_convert_i32_s(v16))
		if base.F32_eq(base.F32_abs(v40), math.Float32frombits(uint32(0x7f800000))) != 0 {
			F_float_overflow_error(m)
			mBase = m.M
			v184 = m.ExcPending
			if v184 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			if v31 != 0 {
				v47 = base.F32_eq(v40, float32(0))
			} else {
				v47 = int32(0)
			}
			if v47 != 0 {
				F_float_underflow_error(m)
				mBase = m.M
				v178 = m.ExcPending
				if v178 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				if base.F64_gt(base.F64_promote_f32(v40), float64(0.3)) == int32(0) {
					return
				} else {
					if l1 == int32(0) {
						v55 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
						v56 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
						v57 = base.F64_sub(v55, v56)
						v59 = math.Float64frombits(uint64(0x7ff0000000000000))
						if base.F64_ne(base.F64_abs(v57), v59)|base.F64_eq(base.F64_abs(v55), v59)|base.F64_eq(base.F64_abs(v56), v59) != 0 {
							v84 = v57
							v86 = math.Float64frombits(uint64(0x7ff0000000000000))
							v88 = base.F64_sub(l4, l2)
							v89 = base.F64_abs(v88)
							if base.B2i32(base.F64_eq(base.F64_abs(l2), v86)|base.F64_ne(v89, v86) == int32(0))&base.F64_ne(base.F64_abs(l4), v86) != 0 {
								F_float_overflow_error(m)
								mBase = m.M
								v184 = m.ExcPending
								if v184 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								if base.B2i32(base.Ui64(base.I64_reinterpret_f64(v89)) <= base.Ui64(int64(9218868437227405312)))&base.F64_eq(v84, float64(0)) != 0 {
									F_float_zero_divide_error(m)
									mBase = m.M
									v172 = m.ExcPending
									if v172 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v105 = base.F64_div(v88, v84)
									v107 = math.Float64frombits(uint64(0x7ff0000000000000))
									if base.F64_eq(base.F64_abs(v105), v107)&base.F64_ne(v89, v107) != 0 {
										F_float_overflow_error(m)
										mBase = m.M
										v184 = m.ExcPending
										if v184 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v112 = float64(0)
										if base.B2i32(base.F64_eq(v88, v112)|base.F64_ne(v105, v112) == int32(0))&base.F64_ne(base.F64_abs(v84), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
											F_float_underflow_error(m)
											mBase = m.M
											v178 = m.ExcPending
											if v178 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											v123 = base.F32_demote_f64(v105)
											v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
											if v124 != 0 {
												*(*float64)(unsafe.Add(mBase, uint32(l0)+80)) = v84
												*(*float32)(unsafe.Add(mBase, uint32(l0)+64)) = v40
												v153 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)) = uint8(v153)
												*(*float32)(unsafe.Add(mBase, uint32(l0)+68)) = v123
												*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = l2
												*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = l1
												*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = l4
											} else {
												v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
												if l1 == v125 {
													v127 = *(*float32)(unsafe.Add(mBase, uint32(l0)+68))
													if base.F32_gt(v127, v123) != 0 {
														*(*float64)(unsafe.Add(mBase, uint32(l0)+80)) = v84
														*(*float32)(unsafe.Add(mBase, uint32(l0)+64)) = v40
														v153 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)) = uint8(v153)
														*(*float32)(unsafe.Add(mBase, uint32(l0)+68)) = v123
														*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = l2
														*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = l1
														*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = l4
													} else {
														if base.F32_ne(v123, v127) != 0 {
														} else {
															v130 = *(*float32)(unsafe.Add(mBase, uint32(l0)+64))
															if base.F32_gt(v40, v130) != 0 {
																*(*float64)(unsafe.Add(mBase, uint32(l0)+80)) = v84
																*(*float32)(unsafe.Add(mBase, uint32(l0)+64)) = v40
																v153 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)) = uint8(v153)
																*(*float32)(unsafe.Add(mBase, uint32(l0)+68)) = v123
																*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = l2
																*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = l1
																*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = l4
															} else {
															}
														}
													}
												} else {
													if base.F64_ge(v105, float64(-7.006492321624085e-46)) != 0 {
														v135 = v123
													} else {
														v135 = float32(0)
													}
													v136 = *(*float32)(unsafe.Add(mBase, uint32(l0)+68))
													v137 = float32(0)
													if base.F32_ge(v136, v137) != 0 {
														v140 = v136
													} else {
														v140 = v137
													}
													if base.F32_lt(v135, v140) != 0 {
														*(*float64)(unsafe.Add(mBase, uint32(l0)+80)) = v84
														*(*float32)(unsafe.Add(mBase, uint32(l0)+64)) = v40
														v153 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)) = uint8(v153)
														*(*float32)(unsafe.Add(mBase, uint32(l0)+68)) = v123
														*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = l2
														*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = l1
														*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = l4
													} else {
														if base.F32_le(v135, v140) == int32(0) {
														} else {
															v145 = *(*float64)(unsafe.Add(mBase, uint32(l0)+80))
															if base.F64_gt(v84, v145) == int32(0) {
															} else {
																*(*float64)(unsafe.Add(mBase, uint32(l0)+80)) = v84
																*(*float32)(unsafe.Add(mBase, uint32(l0)+64)) = v40
																v153 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)) = uint8(v153)
																*(*float32)(unsafe.Add(mBase, uint32(l0)+68)) = v123
																*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = l2
																*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = l1
																*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = l4
															}
														}
													}
												}
											}
											return
										}
									}
								}
							}
						} else {
							F_float_overflow_error(m)
							mBase = m.M
							v184 = m.ExcPending
							if v184 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					} else {
						v69 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
						v70 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
						v71 = base.F64_sub(v69, v70)
						v73 = math.Float64frombits(uint64(0x7ff0000000000000))
						if base.F64_ne(base.F64_abs(v71), v73)|base.F64_eq(base.F64_abs(v69), v73) != 0 {
							v84 = v71
							v86 = math.Float64frombits(uint64(0x7ff0000000000000))
							v88 = base.F64_sub(l4, l2)
							v89 = base.F64_abs(v88)
							if base.B2i32(base.F64_eq(base.F64_abs(l2), v86)|base.F64_ne(v89, v86) == int32(0))&base.F64_ne(base.F64_abs(l4), v86) != 0 {
								F_float_overflow_error(m)
								mBase = m.M
								v184 = m.ExcPending
								if v184 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								if base.B2i32(base.Ui64(base.I64_reinterpret_f64(v89)) <= base.Ui64(int64(9218868437227405312)))&base.F64_eq(v84, float64(0)) != 0 {
									F_float_zero_divide_error(m)
									mBase = m.M
									v172 = m.ExcPending
									if v172 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v105 = base.F64_div(v88, v84)
									v107 = math.Float64frombits(uint64(0x7ff0000000000000))
									if base.F64_eq(base.F64_abs(v105), v107)&base.F64_ne(v89, v107) != 0 {
										F_float_overflow_error(m)
										mBase = m.M
										v184 = m.ExcPending
										if v184 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v112 = float64(0)
										if base.B2i32(base.F64_eq(v88, v112)|base.F64_ne(v105, v112) == int32(0))&base.F64_ne(base.F64_abs(v84), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
											F_float_underflow_error(m)
											mBase = m.M
											v178 = m.ExcPending
											if v178 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											v123 = base.F32_demote_f64(v105)
											v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
											if v124 != 0 {
												*(*float64)(unsafe.Add(mBase, uint32(l0)+80)) = v84
												*(*float32)(unsafe.Add(mBase, uint32(l0)+64)) = v40
												v153 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)) = uint8(v153)
												*(*float32)(unsafe.Add(mBase, uint32(l0)+68)) = v123
												*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = l2
												*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = l1
												*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = l4
											} else {
												v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
												if l1 == v125 {
													v127 = *(*float32)(unsafe.Add(mBase, uint32(l0)+68))
													if base.F32_gt(v127, v123) != 0 {
														*(*float64)(unsafe.Add(mBase, uint32(l0)+80)) = v84
														*(*float32)(unsafe.Add(mBase, uint32(l0)+64)) = v40
														v153 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)) = uint8(v153)
														*(*float32)(unsafe.Add(mBase, uint32(l0)+68)) = v123
														*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = l2
														*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = l1
														*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = l4
													} else {
														if base.F32_ne(v123, v127) != 0 {
														} else {
															v130 = *(*float32)(unsafe.Add(mBase, uint32(l0)+64))
															if base.F32_gt(v40, v130) != 0 {
																*(*float64)(unsafe.Add(mBase, uint32(l0)+80)) = v84
																*(*float32)(unsafe.Add(mBase, uint32(l0)+64)) = v40
																v153 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)) = uint8(v153)
																*(*float32)(unsafe.Add(mBase, uint32(l0)+68)) = v123
																*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = l2
																*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = l1
																*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = l4
															} else {
															}
														}
													}
												} else {
													if base.F64_ge(v105, float64(-7.006492321624085e-46)) != 0 {
														v135 = v123
													} else {
														v135 = float32(0)
													}
													v136 = *(*float32)(unsafe.Add(mBase, uint32(l0)+68))
													v137 = float32(0)
													if base.F32_ge(v136, v137) != 0 {
														v140 = v136
													} else {
														v140 = v137
													}
													if base.F32_lt(v135, v140) != 0 {
														*(*float64)(unsafe.Add(mBase, uint32(l0)+80)) = v84
														*(*float32)(unsafe.Add(mBase, uint32(l0)+64)) = v40
														v153 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)) = uint8(v153)
														*(*float32)(unsafe.Add(mBase, uint32(l0)+68)) = v123
														*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = l2
														*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = l1
														*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = l4
													} else {
														if base.F32_le(v135, v140) == int32(0) {
														} else {
															v145 = *(*float64)(unsafe.Add(mBase, uint32(l0)+80))
															if base.F64_gt(v84, v145) == int32(0) {
															} else {
																*(*float64)(unsafe.Add(mBase, uint32(l0)+80)) = v84
																*(*float32)(unsafe.Add(mBase, uint32(l0)+64)) = v40
																v153 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)) = uint8(v153)
																*(*float32)(unsafe.Add(mBase, uint32(l0)+68)) = v123
																*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = l2
																*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = l1
																*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = l4
															}
														}
													}
												}
											}
											return
										}
									}
								}
							}
						} else {
							if base.F64_ne(base.F64_abs(v70), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								F_float_overflow_error(m)
								mBase = m.M
								v184 = m.ExcPending
								if v184 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v84 = v71
								v86 = math.Float64frombits(uint64(0x7ff0000000000000))
								v88 = base.F64_sub(l4, l2)
								v89 = base.F64_abs(v88)
								if base.B2i32(base.F64_eq(base.F64_abs(l2), v86)|base.F64_ne(v89, v86) == int32(0))&base.F64_ne(base.F64_abs(l4), v86) != 0 {
									F_float_overflow_error(m)
									mBase = m.M
									v184 = m.ExcPending
									if v184 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									if base.B2i32(base.Ui64(base.I64_reinterpret_f64(v89)) <= base.Ui64(int64(9218868437227405312)))&base.F64_eq(v84, float64(0)) != 0 {
										F_float_zero_divide_error(m)
										mBase = m.M
										v172 = m.ExcPending
										if v172 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v105 = base.F64_div(v88, v84)
										v107 = math.Float64frombits(uint64(0x7ff0000000000000))
										if base.F64_eq(base.F64_abs(v105), v107)&base.F64_ne(v89, v107) != 0 {
											F_float_overflow_error(m)
											mBase = m.M
											v184 = m.ExcPending
											if v184 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											v112 = float64(0)
											if base.B2i32(base.F64_eq(v88, v112)|base.F64_ne(v105, v112) == int32(0))&base.F64_ne(base.F64_abs(v84), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
												F_float_underflow_error(m)
												mBase = m.M
												v178 = m.ExcPending
												if v178 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												v123 = base.F32_demote_f64(v105)
												v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
												if v124 != 0 {
													*(*float64)(unsafe.Add(mBase, uint32(l0)+80)) = v84
													*(*float32)(unsafe.Add(mBase, uint32(l0)+64)) = v40
													v153 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)) = uint8(v153)
													*(*float32)(unsafe.Add(mBase, uint32(l0)+68)) = v123
													*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = l2
													*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = l1
													*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = l4
												} else {
													v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
													if l1 == v125 {
														v127 = *(*float32)(unsafe.Add(mBase, uint32(l0)+68))
														if base.F32_gt(v127, v123) != 0 {
															*(*float64)(unsafe.Add(mBase, uint32(l0)+80)) = v84
															*(*float32)(unsafe.Add(mBase, uint32(l0)+64)) = v40
															v153 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)) = uint8(v153)
															*(*float32)(unsafe.Add(mBase, uint32(l0)+68)) = v123
															*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = l2
															*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = l1
															*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = l4
														} else {
															if base.F32_ne(v123, v127) != 0 {
															} else {
																v130 = *(*float32)(unsafe.Add(mBase, uint32(l0)+64))
																if base.F32_gt(v40, v130) != 0 {
																	*(*float64)(unsafe.Add(mBase, uint32(l0)+80)) = v84
																	*(*float32)(unsafe.Add(mBase, uint32(l0)+64)) = v40
																	v153 = int32(0)
																	*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)) = uint8(v153)
																	*(*float32)(unsafe.Add(mBase, uint32(l0)+68)) = v123
																	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = l2
																	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = l1
																	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = l4
																} else {
																}
															}
														}
													} else {
														if base.F64_ge(v105, float64(-7.006492321624085e-46)) != 0 {
															v135 = v123
														} else {
															v135 = float32(0)
														}
														v136 = *(*float32)(unsafe.Add(mBase, uint32(l0)+68))
														v137 = float32(0)
														if base.F32_ge(v136, v137) != 0 {
															v140 = v136
														} else {
															v140 = v137
														}
														if base.F32_lt(v135, v140) != 0 {
															*(*float64)(unsafe.Add(mBase, uint32(l0)+80)) = v84
															*(*float32)(unsafe.Add(mBase, uint32(l0)+64)) = v40
															v153 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)) = uint8(v153)
															*(*float32)(unsafe.Add(mBase, uint32(l0)+68)) = v123
															*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = l2
															*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = l1
															*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = l4
														} else {
															if base.F32_le(v135, v140) == int32(0) {
															} else {
																v145 = *(*float64)(unsafe.Add(mBase, uint32(l0)+80))
																if base.F64_gt(v84, v145) == int32(0) {
																} else {
																	*(*float64)(unsafe.Add(mBase, uint32(l0)+80)) = v84
																	*(*float32)(unsafe.Add(mBase, uint32(l0)+64)) = v40
																	v153 = int32(0)
																	*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)) = uint8(v153)
																	*(*float32)(unsafe.Add(mBase, uint32(l0)+68)) = v123
																	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = l2
																	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = l1
																	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = l4
																}
															}
														}
													}
												}
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
	}
}
func F_g_intbig_same(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13906(m, l0, int32(2), int32(4), int32(252))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
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
func F_gbtreekey_in(m *base.Module, l0 int32) int32 {
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
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v19 = F_format_type_extended(m, v7, int32(-1), int32(2))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = v19
				F_errmsg(m, int32(_a_F_gbtreekey_in_0), v5)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_gbtreekey_in_1), int32(34), int32(_a_F_gbtreekey_in_2))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
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
func F_gen_partprune_steps_internal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
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
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
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
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
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
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
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
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
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
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v273 int32
	_ = v273
	var v291 int32
	_ = v291
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
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
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
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
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
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
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
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v676 int32
	_ = v676
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v719 int32
	_ = v719
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v739 int32
	_ = v739
	var v755 int32
	_ = v755
	var v761 int32
	_ = v761
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v782 int32
	_ = v782
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v871 int32
	_ = v871
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v941 int32
	_ = v941
	var v944 int32
	_ = v944
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v965 int32
	_ = v965
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v971 int32
	_ = v971
	var v973 int32
	_ = v973
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1005 int32
	_ = v1005
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1020 int32
	_ = v1020
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1028 int32
	_ = v1028
	var v1034 int32
	_ = v1034
	var v1041 int32
	_ = v1041
	var v1050 int32
	_ = v1050
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1083 int32
	_ = v1083
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1142 int32
	_ = v1142
	var v1144 int32
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1169 int32
	_ = v1169
	var v1171 int32
	_ = v1171
	var v1174 int32
	_ = v1174
	var v1177 int32
	_ = v1177
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1188 int32
	_ = v1188
	var v1192 int32
	_ = v1192
	var v1194 int32
	_ = v1194
	var v1200 int32
	_ = v1200
	var v1203 int32
	_ = v1203
	var v1205 int32
	_ = v1205
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1215 int32
	_ = v1215
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1220 int64
	_ = v1220
	var v1228 int32
	_ = v1228
	var v1241 int32
	_ = v1241
	var v1256 int32
	_ = v1256
	var v1260 int32
	_ = v1260
	var v1263 int32
	_ = v1263
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1282 int32
	_ = v1282
	var v1287 int32
	_ = v1287
	var v1304 int32
	_ = v1304
	var v1308 int32
	_ = v1308
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1317 int32
	_ = v1317
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1344 int32
	_ = v1344
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1352 int32
	_ = v1352
	var v1361 int32
	_ = v1361
	var v1362 int32
	_ = v1362
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1374 int32
	_ = v1374
	var v1378 int32
	_ = v1378
	var v1383 int32
	_ = v1383
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1394 int32
	_ = v1394
	var v1399 int32
	_ = v1399
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1450 int32
	_ = v1450
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1481 int32
	_ = v1481
	var v1485 int32
	_ = v1485
	var v1488 int32
	_ = v1488
	var v1490 int32
	_ = v1490
	var v1493 int32
	_ = v1493
	var v1494 int32
	_ = v1494
	var v1495 int32
	_ = v1495
	var v1499 int32
	_ = v1499
	var v1509 int32
	_ = v1509
	var v1512 int32
	_ = v1512
	var v1528 int32
	_ = v1528
	var v1531 int32
	_ = v1531
	var v1541 int32
	_ = v1541
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1549 int32
	_ = v1549
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1557 int32
	_ = v1557
	var v1558 int32
	_ = v1558
	var v1564 int32
	_ = v1564
	var v1565 int32
	_ = v1565
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1571 int32
	_ = v1571
	var v1575 int32
	_ = v1575
	var v1580 int32
	_ = v1580
	var v1583 int32
	_ = v1583
	var v1584 int32
	_ = v1584
	var v1585 int32
	_ = v1585
	var v1592 int32
	_ = v1592
	var v1593 int32
	_ = v1593
	var v1594 int32
	_ = v1594
	var v1596 int32
	_ = v1596
	var v1597 int32
	_ = v1597
	var v1598 int32
	_ = v1598
	var v1599 int32
	_ = v1599
	var v1600 int32
	_ = v1600
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1604 int32
	_ = v1604
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1612 int32
	_ = v1612
	var v1613 int32
	_ = v1613
	var v1615 int32
	_ = v1615
	var v1630 int32
	_ = v1630
	var v1633 int32
	_ = v1633
	var v1634 int32
	_ = v1634
	var v1636 int32
	_ = v1636
	var v1637 int32
	_ = v1637
	var v1639 int32
	_ = v1639
	var v1640 int32
	_ = v1640
	var v1641 int32
	_ = v1641
	var v1642 int32
	_ = v1642
	var v1644 int32
	_ = v1644
	var v1645 int32
	_ = v1645
	var v1646 int32
	_ = v1646
	var v1648 int32
	_ = v1648
	var v1649 int32
	_ = v1649
	var v1652 int32
	_ = v1652
	var v1657 int32
	_ = v1657
	var v1674 int32
	_ = v1674
	var v1677 int32
	_ = v1677
	var v1678 int32
	_ = v1678
	var v1679 int32
	_ = v1679
	var v1681 int32
	_ = v1681
	var v1682 int32
	_ = v1682
	var v1684 int32
	_ = v1684
	var v1685 int32
	_ = v1685
	var v1693 int32
	_ = v1693
	var v1698 int32
	_ = v1698
	var v1717 int32
	_ = v1717
	var v1719 int32
	_ = v1719
	var v1722 int32
	_ = v1722
	var v1739 int32
	_ = v1739
	var v1740 int32
	_ = v1740
	var v1742 int32
	_ = v1742
	var v1743 int32
	_ = v1743
	var v1746 int32
	_ = v1746
	var v1747 int32
	_ = v1747
	var v1748 int32
	_ = v1748
	var v1749 int32
	_ = v1749
	var v1751 int32
	_ = v1751
	var v1752 int32
	_ = v1752
	var v1753 int32
	_ = v1753
	var v1755 int32
	_ = v1755
	var v1756 int32
	_ = v1756
	var v1760 int32
	_ = v1760
	var v1765 int32
	_ = v1765
	var v1782 int32
	_ = v1782
	var v1785 int32
	_ = v1785
	var v1786 int32
	_ = v1786
	var v1787 int32
	_ = v1787
	var v1789 int32
	_ = v1789
	var v1790 int32
	_ = v1790
	var v1792 int32
	_ = v1792
	var v1793 int32
	_ = v1793
	var v1801 int32
	_ = v1801
	var v1802 int32
	_ = v1802
	var v1807 int32
	_ = v1807
	var v1823 int32
	_ = v1823
	var v1824 int32
	_ = v1824
	var v1826 int32
	_ = v1826
	var v1827 int32
	_ = v1827
	var v1829 int32
	_ = v1829
	var v1830 int32
	_ = v1830
	var v1831 int32
	_ = v1831
	var v1832 int32
	_ = v1832
	var v1834 int32
	_ = v1834
	var v1835 int32
	_ = v1835
	var v1837 int32
	_ = v1837
	var v1838 int32
	_ = v1838
	var v1850 int32
	_ = v1850
	var v1855 int32
	_ = v1855
	var v1867 int32
	_ = v1867
	var v1870 int32
	_ = v1870
	var v1871 int32
	_ = v1871
	var v1872 int32
	_ = v1872
	var v1874 int32
	_ = v1874
	var v1875 int32
	_ = v1875
	var v1877 int32
	_ = v1877
	var v1878 int32
	_ = v1878
	var v1882 int32
	_ = v1882
	var v1888 int32
	_ = v1888
	var v1892 int32
	_ = v1892
	var v1910 int32
	_ = v1910
	var v1911 int32
	_ = v1911
	var v1919 int32
	_ = v1919
	var v1936 int32
	_ = v1936
	var v1937 int32
	_ = v1937
	var v1938 int32
	_ = v1938
	var v1940 int32
	_ = v1940
	var v1941 int32
	_ = v1941
	var v1965 int32
	_ = v1965
	var v1966 int32
	_ = v1966
	var v1967 int32
	_ = v1967
	var v1969 int32
	_ = v1969
	var v1970 int32
	_ = v1970
	var v1985 int32
	_ = v1985
	var v1996 int32
	_ = v1996
	var v1999 int32
	_ = v1999
	var v2002 int32
	_ = v2002
	var v2003 int32
	_ = v2003
	var v2006 int32
	_ = v2006
	var v2012 int32
	_ = v2012
	var v2013 int32
	_ = v2013
	var v2018 int32
	_ = v2018
	var v2023 int32
	_ = v2023
	var v2038 int32
	_ = v2038
	var v2042 int32
	_ = v2042
	var v2043 int32
	_ = v2043
	var v2049 int32
	_ = v2049
	var v2059 int32
	_ = v2059
	var v2071 int32
	_ = v2071
	var v2075 int32
	_ = v2075
	var v2076 int32
	_ = v2076
	var v2077 int32
	_ = v2077
	var v2078 int32
	_ = v2078
	var v2079 int32
	_ = v2079
	var v2080 int32
	_ = v2080
	var v2081 int32
	_ = v2081
	var v2083 int32
	_ = v2083
	var v2084 int32
	_ = v2084
	var v2086 int32
	_ = v2086
	var v2087 int32
	_ = v2087
	var v2089 int32
	_ = v2089
	var v2090 int32
	_ = v2090
	var v2116 int32
	_ = v2116
	var v2120 int32
	_ = v2120
	var v2128 int32
	_ = v2128
	var v2129 int32
	_ = v2129
	var v2139 int32
	_ = v2139
	var v2140 int32
	_ = v2140
	var v2144 int32
	_ = v2144
	var v2145 int32
	_ = v2145
	var v2151 int32
	_ = v2151
	var v2156 int32
	_ = v2156
	var v2158 int32
	_ = v2158
	var v2159 int32
	_ = v2159
	var v2162 int32
	_ = v2162
	var v2169 int32
	_ = v2169
	var v2172 int32
	_ = v2172
	var v2173 int32
	_ = v2173
	var v2174 int32
	_ = v2174
	var v2176 int32
	_ = v2176
	var v2177 int32
	_ = v2177
	var v2190 int32
	_ = v2190
	var v2198 int32
	_ = v2198
	var v2201 int32
	_ = v2201
	var v2208 int32
	_ = v2208
	var v2209 int32
	_ = v2209
	var v2212 int32
	_ = v2212
	var v2216 int32
	_ = v2216
	var v2218 int32
	_ = v2218
	var v2224 int32
	_ = v2224
	var v2227 int32
	_ = v2227
	var v2229 int32
	_ = v2229
	var v2236 int32
	_ = v2236
	var v2237 int32
	_ = v2237
	var v2240 int32
	_ = v2240
	var v2241 int32
	_ = v2241
	var v2244 int32
	_ = v2244
	var v2248 int32
	_ = v2248
	var v2255 int32
	_ = v2255
	var v2256 int32
	_ = v2256
	var v2257 int32
	_ = v2257
	var v2259 int32
	_ = v2259
	var v2260 int32
	_ = v2260
	var v2261 int32
	_ = v2261
	var v2265 int32
	_ = v2265
	var v2273 int32
	_ = v2273
	var v2287 int32
	_ = v2287
	var v2290 int32
	_ = v2290
	var v2295 int32
	_ = v2295
	var v2298 int32
	_ = v2298
	var v2315 int32
	_ = v2315
	var v2319 int32
	_ = v2319
	var v2320 int32
	_ = v2320
	var v2321 int32
	_ = v2321
	var v2322 int32
	_ = v2322
	var v2324 int32
	_ = v2324
	var v2325 int32
	_ = v2325
	var v2328 int32
	_ = v2328
	var v2329 int32
	_ = v2329
	var v2332 int32
	_ = v2332
	var v2333 int32
	_ = v2333
	var v2340 int32
	_ = v2340
	var v2341 int32
	_ = v2341
	var v2342 int32
	_ = v2342
	var v2344 int32
	_ = v2344
	var v2345 int32
	_ = v2345
	var v2350 int32
	_ = v2350
	var v2358 int32
	_ = v2358
	v3 = int32(0)
	v24 = m.G0
	v26 = v24 - int32(288)
	m.G0 = v26
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+232))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v28)+240))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+32))
	if v31 == int32(-1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v2350 + int32(288)
	return v2358
L2:
	;
	v46 = int32(0)
	base.MemoryFill(m, v26+int32(80), v46, int32(128))
	if l1 == v46 {
		v2190 = v3
		v2198 = v3
		goto L8
	} else {
		goto L9
	}
L3:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)+248))
	v36 = F_predicate_refuted_by(m, v34, l1, int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	if v36 == int32(0) {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v42 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)) = uint8(v42)
	v2350 = v26
	v2358 = v3
	goto L1
L7:
	;
	if v2273 == int32(0) {
		goto L502
	} else {
		goto L503
	}
L8:
	;
	v2201 = int32(0)
	if v2198 == v2201 {
		goto L486
	} else {
		goto L487
	}
L9:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v51 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v66 = v3
	v67 = v3
	v74 = v3
	v75 = v3
	v76 = v3
	goto L13
L11:
	;
	v1161 = v3
	v1162 = v3
	v1169 = v3
	v1171 = v3
	goto L12
L12:
	;
	if v1162 == int32(0) {
		v1215 = v3
		goto L318
	} else {
		goto L319
	}
L13:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v77+v75<<(uint(int32(2))%32))))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	if v82 == int32(318) {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	v1161 = v1134
	v1162 = v1135
	v1169 = v1142
	v1171 = v1144
	goto L12
L15:
	;
	v1146 = v75 + int32(1)
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v1146 < v1147 {
		v66 = v1134
		v67 = v1135
		v74 = v1142
		v75 = v1146
		v76 = v1144
		goto L13
	} else {
		goto L316
	}
L16:
	;
	v263 = int32(*(*int16)(unsafe.Add(mBase, uint32(v29)+2)))
	if v263 <= int32(0) {
		v1134 = v66
		v1135 = v67
		v1142 = v74
		v1144 = v76
		goto L15
	} else {
		goto L63
	}
L17:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
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
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v87 = v85
	v88 = v86
	goto L20
L19:
	;
	v87 = v81
	v88 = v82
	goto L20
L20:
	;
	v90 = v88 - int32(7)
	if v90 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	if v90 == int32(14) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+24)))
	if v93 == int32(0) {
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
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v87)+20))
	if v96 != 0 {
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
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)) = uint8(v97)
	v2350 = v26
	v2358 = int32(0)
	goto L1
L30:
	;
	goto L29
L31:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
	v247 = F_gen_partprune_steps_internal(m, l0, v246)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L4
	} else {
		goto L57
	}
L32:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
	if v101 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	if v188 == int32(0) {
		v1134 = v66
		v1135 = v67
		v1142 = v74
		v1144 = v76
		goto L15
	} else {
		goto L53
	}
L34:
	;
	v222 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)) = uint8(v222)
	v2350 = v26
	v2358 = int32(0)
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
	v113 = v105
	v117 = v105
	v119 = int32(1)
	goto L37
L37:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v101)+12))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v133+v113<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+72)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v26)+76)) = v137
	v143 = F_list_make1_impl(m, int32(1), v26+int32(72))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L4
	} else {
		goto L39
	}
L38:
	;
	if v189&int32(1) == int32(0) {
		goto L33
	} else {
		goto L52
	}
L39:
	;
	v145 = F_gen_partprune_steps_internal(m, l0, v143)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	v148 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)) = uint8(v148)
	if v147 == v148 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	if v145 != 0 {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	v188 = v117
	v189 = v119
	goto L43
L43:
	;
	v192 = v113 + int32(1)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	if v192 < v193 {
		v113 = v192
		v117 = v188
		v119 = v189
		goto L37
	} else {
		goto L51
	}
L44:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
	v185 = F_lappend_int(m, v117, v184)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L4
	} else {
		goto L50
	}
L45:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v145)+12))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v145)+4))
	v157 = int32(4)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v152+v153<<(uint(int32(2))%32)-v157)))
	v182 = v159 + v157
	goto L44
L46:
	;
	goto L47
L47:
	;
	v163 = F_palloc0(m, int32(16))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = int32(378)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v167 + int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v163)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v163)+4)) = v167
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v175 = F_lappend(m, v174, v163)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v175
	v182 = v163 + int32(4)
	goto L44
L50:
	;
	v188 = v185
	v189 = int32(0)
	goto L43
L51:
	;
	goto L38
L52:
	;
	goto L34
L53:
	;
	v228 = F_palloc0(m, int32(16))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v228))) = int32(378)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v232 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v228)+12)) = v188
	*(*int32)(unsafe.Add(mBase, uint32(v228)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v228)+4)) = v232
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v241 = F_lappend(m, v240, v228)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v241
	v244 = F_lappend(m, v66, v228)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	v1134 = v244
	v1135 = v67
	v1142 = v74
	v1144 = v76
	goto L15
L57:
	;
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	if v249 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v2350 = v26
	v2358 = int32(0)
	goto L1
L59:
	;
	goto L60
L60:
	;
	if v247 == int32(0) {
		v1134 = v66
		v1135 = v67
		v1142 = v74
		v1144 = v76
		goto L15
	} else {
		goto L61
	}
L61:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v247)+12))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v247)+4))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v253+v254<<(uint(int32(2))%32)-int32(4))))
	v261 = F_lappend(m, v66, v260)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	v1134 = v261
	v1135 = v67
	v1142 = v74
	v1144 = v76
	goto L15
L63:
	;
	v273 = int32(0)
	goto L65
L64:
	;
	v1109 = F_bms_is_member(m, v273, v67)
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L4
	} else {
		goto L311
	}
L65:
	;
	v291 = v273 << (uint(int32(2)) % 32)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v292)+264))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v291+v293)))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v295)+12))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v296)))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v292)+232))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v298)+12))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v299+v291)))
	v302 = int32(5)
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v298)+4))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v303+v291)))
	if base.B2i32(v305 != int32(2222))&base.B2i32(v305 != int32(424)) != 0 {
		v362 = v302
		goto L82
	} else {
		goto L83
	}
L66:
	;
	v1098 = F_list_concat(m, v66, v1083)
	mBase = m.M
	v1099 = m.ExcPending
	if v1099 != 0 {
		goto L4
	} else {
		goto L310
	}
L67:
	;
	goto L66
L68:
	;
	v1072 = v273 + int32(1)
	v1073 = int32(*(*int16)(unsafe.Add(mBase, uint32(v29)+2)))
	if v1072 < v1073 {
		v273 = v1072
		goto L65
	} else {
		goto L309
	}
L69:
	;
	if v1052 == int32(0) {
		goto L68
	} else {
		goto L308
	}
L70:
	;
	v1050 = int32(0)
	v1052 = int32(5)
	v1053 = v364
	goto L69
L71:
	;
	if v301 != 0 {
		goto L248
	} else {
		goto L249
	}
L72:
	;
	v871 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)) = uint8(v871)
	v2350 = v26
	v2358 = int32(0)
	goto L1
L73:
	;
	v841 = F_bms_is_member(m, v273, v67)
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L4
	} else {
		goto L243
	}
L74:
	;
	v822 = F_bms_is_member(m, v273, v74)
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L4
	} else {
		goto L237
	}
L75:
	;
	if v366 != int32(52) {
		v1050 = v364
		v1052 = v362
		v1053 = v364
		goto L69
	} else {
		goto L229
	}
L76:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v87)+28))
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v482)+12))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v483)))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v484)))
	if v485 == int32(27) {
		goto L142
	} else {
		goto L143
	}
L77:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v87)+28))
	if v450 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L78:
	;
	v431 = F_makeBoolConst(m, v429, int32(0))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L4
	} else {
		goto L120
	}
L79:
	;
	v429 = int32(0)
	goto L78
L80:
	;
	v372 = F_makeBoolConst(m, v370, int32(0))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L4
	} else {
		goto L105
	}
L81:
	;
	v370 = int32(0)
	goto L80
L82:
	;
	v364 = int32(0)
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	switch v366 - int32(17) {
	case 0:
		goto L77
	case 1, 2:
		v1050 = v364
		v1052 = v362
		v1053 = v364
		goto L69
	case 3:
		goto L76
	default:
		goto L75
	}
L83:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	if v311 != int32(21) {
		goto L87
	} else {
		goto L88
	}
L84:
	;
	v362 = int32(0)
	goto L82
L85:
	;
	v348 = F_equal(m, v345, v297)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L4
	} else {
		goto L98
	}
L86:
	;
	if v339 != int32(27) {
		v345 = v338
		v347 = v341
		goto L85
	} else {
		goto L97
	}
L87:
	;
	if v311 != int32(53) {
		v338 = v87
		v339 = v311
		v341 = int32(0)
		goto L86
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v330 != int32(2) {
		v345 = v87
		v347 = int32(0)
		goto L85
	} else {
		goto L96
	}
L90:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v317)))
	if v318 == int32(27) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v317)+4))
	v322 = v321
	goto L93
L92:
	;
	v322 = v317
	goto L93
L93:
	;
	v323 = F_equal(m, v322, v297)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L4
	} else {
		goto L94
	}
L94:
	;
	if v323 == int32(0) {
		goto L84
	} else {
		goto L95
	}
L95:
	;
	v327 = int32(1)
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
	switch v328 {
	case 0:
		v429 = v327
		goto L78
	case 1:
		v370 = v327
		goto L80
	case 2:
		goto L79
	case 3:
		goto L81
	case 4:
		goto L74
	case 5:
		goto L73
	default:
		v362 = v302
		goto L82
	}
L96:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v333)+12))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v334)))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v335)))
	v338 = v335
	v339 = v336
	v341 = int32(1)
	goto L86
L97:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v338)+4))
	v345 = v344
	v347 = v341
	goto L85
L98:
	;
	if v348 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v429 = v347 ^ int32(1)
	goto L78
L100:
	;
	goto L101
L101:
	;
	v352 = F_negate_clause(m, v345)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L4
	} else {
		goto L102
	}
L102:
	;
	v354 = F_equal(m, v352, v297)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L4
	} else {
		goto L103
	}
L103:
	;
	if v354 != 0 {
		v429 = v347
		goto L78
	} else {
		goto L104
	}
L104:
	;
	goto L84
L105:
	;
	v375 = F_copyObjectImpl(m, v87)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L4
	} else {
		goto L109
	}
L106:
	;
	v385 = F_palloc0(m, int32(20))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L4
	} else {
		goto L110
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v375)+8)) = v381
	goto L106
L108:
	;
	v381 = int32(0)
	goto L107
L109:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v375)+8))
	switch v377 - int32(1) {
	case 0:
		v381 = int32(2)
		goto L107
	default:
		goto L106
	case 2:
		goto L108
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v385))) = int32(52)
	v389 = F_copyObjectImpl(m, v297)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L4
	} else {
		goto L111
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v385)+16)) = int32(-1)
	v393 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v385)+12)) = uint8(v393)
	*(*int32)(unsafe.Add(mBase, uint32(v385)+8)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v385)+4)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v26)+284)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v26)+232)) = v375
	*(*int32)(unsafe.Add(mBase, uint32(v26)+68)) = v375
	*(*int32)(unsafe.Add(mBase, uint32(v26)+64)) = v385
	v408 = F_list_make2_impl(m, v26+int32(68), v26-int32(-64))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L4
	} else {
		goto L112
	}
L112:
	;
	v411 = F_makeBoolExpr(m, int32(1), v408, int32(-1))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L4
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v411
	*(*int32)(unsafe.Add(mBase, uint32(v26)+280)) = v411
	v418 = F_list_make1_impl(m, int32(1), v26+int32(60))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L4
	} else {
		goto L114
	}
L114:
	;
	v420 = F_gen_partprune_steps_internal(m, l0, v418)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L4
	} else {
		goto L115
	}
L115:
	;
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	if v422 != 0 {
		goto L72
	} else {
		goto L116
	}
L116:
	;
	if v420 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v425 = int32(3)
	goto L119
L118:
	;
	v425 = int32(5)
	goto L119
L119:
	;
	v1050 = v393
	v1052 = v425
	v1053 = v420
	goto L69
L120:
	;
	v434 = F_palloc(m, int32(24))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L4
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v434)+12)) = v431
	v437 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v434)+8)) = uint8(v437)
	*(*int32)(unsafe.Add(mBase, uint32(v434)+4)) = int32(91)
	*(*int32)(unsafe.Add(mBase, uint32(v434))) = v273
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v298)+24))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v442+v273*int32(28))+4))
	*(*int32)(unsafe.Add(mBase, uint32(v434)+20)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v434)+16)) = v446
	v1100 = v434
	goto L64
L122:
	;
	v1050 = int32(0)
	v1052 = v362
	v1053 = v364
	goto L69
L123:
	;
	goto L124
L124:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v450)+4))
	if v454 != int32(2) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v1050 = int32(0)
	v1052 = v362
	v1053 = v364
	goto L69
L126:
	;
	goto L127
L127:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v450)+12))
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v458)))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v459)))
	if v460 == int32(27) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v459)+4))
	v464 = v463
	goto L130
L129:
	;
	v464 = v459
	goto L130
L130:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v458)+4))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v465)))
	if v466 == int32(27) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v465)+4))
	v470 = v469
	goto L133
L132:
	;
	v470 = v465
	goto L133
L133:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	v472 = F_equal(m, v464, v297)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L4
	} else {
		goto L134
	}
L134:
	;
	if v472 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v875 = v470
	v877 = v471
	goto L71
L136:
	;
	goto L137
L137:
	;
	v474 = int32(0)
	v476 = F_equal(m, v470, v297)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L4
	} else {
		goto L138
	}
L138:
	;
	if v476 == int32(0) {
		v1050 = v474
		v1052 = v474
		v1053 = v364
		goto L69
	} else {
		goto L139
	}
L139:
	;
	v480 = F_get_commutator(m, v471)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L4
	} else {
		goto L140
	}
L140:
	;
	if v480 != 0 {
		v875 = v464
		v877 = v480
		goto L71
	} else {
		goto L141
	}
L141:
	;
	goto L70
L142:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v484)+4))
	v489 = v488
	goto L144
L143:
	;
	v489 = v484
	goto L144
L144:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v87)+24))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v483)+4))
	v493 = F_equal(m, v489, v297)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L4
	} else {
		goto L145
	}
L145:
	;
	if v493 == int32(0) {
		goto L68
	} else {
		goto L146
	}
L146:
	;
	if v301 != 0 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v87)+24))
	if v301 != v497 {
		goto L68
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	v499 = F_op_in_opfamily(m, v491, v305)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L4
	} else {
		goto L151
	}
L150:
	;
	goto L149
L151:
	;
	if v499 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298))))
	if v503 != int32(108) {
		goto L68
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	v527 = F_op_strict(m, v491)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L4
	} else {
		goto L162
	}
L155:
	;
	v506 = F_get_negator(m, v491)
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L4
	} else {
		goto L156
	}
L156:
	;
	if v506 == int32(0) {
		goto L68
	} else {
		goto L157
	}
L157:
	;
	v510 = F_op_in_opfamily(m, v506, v305)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L4
	} else {
		goto L158
	}
L158:
	;
	if v510 == int32(0) {
		goto L68
	} else {
		goto L159
	}
L159:
	;
	F_get_op_opfamily_properties(m, v506, v305, int32(0), v26+int32(240), v26+int32(276), v26+int32(228))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L4
	} else {
		goto L160
	}
L160:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v26)+240))
	if v523 != int32(3) {
		goto L68
	} else {
		goto L161
	}
L161:
	;
	goto L154
L162:
	;
	if v527 == int32(0) {
		v1134 = v66
		v1135 = v67
		v1142 = v74
		v1144 = v76
		goto L15
	} else {
		goto L163
	}
L163:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v492)))
	if v531 == int32(7) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v569 = F_op_volatile(m, v491)
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L4
	} else {
		goto L181
	}
L165:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v534 == int32(0) {
		v1134 = v66
		v1135 = v67
		v1142 = v74
		v1144 = v76
		goto L15
	} else {
		goto L166
	}
L166:
	;
	v537 = F_contain_var_clause(m, v492)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L4
	} else {
		goto L167
	}
L167:
	;
	if v537 != 0 {
		v1134 = v66
		v1135 = v67
		v1142 = v74
		v1144 = v76
		goto L15
	} else {
		goto L168
	}
L168:
	;
	v539 = F_contain_volatile_functions(m, v492)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L4
	} else {
		goto L169
	}
L169:
	;
	if v539 != 0 {
		v1134 = v66
		v1135 = v67
		v1142 = v74
		v1144 = v76
		goto L15
	} else {
		goto L170
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+240)) = int32(0)
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v492)))
	if v543 == int32(8) {
		goto L173
	} else {
		goto L174
	}
L171:
	;
	v567 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)) = uint8(v567)
	goto L164
L172:
	;
	if v559 == int32(0) {
		goto L171
	} else {
		goto L179
	}
L173:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v492)+4))
	if v546 != int32(1) {
		goto L171
	} else {
		goto L176
	}
L174:
	;
	goto L175
L175:
	;
	v556 = F_expression_tree_walker_impl(m, v492, int32(909), v26+int32(240))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L4
	} else {
		goto L178
	}
L176:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v492)+8))
	v551 = F_bms_add_member(m, int32(0), v550)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L4
	} else {
		goto L177
	}
L177:
	;
	v559 = v551
	goto L172
L178:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v26)+240))
	v559 = v558
	goto L172
L179:
	;
	v562 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)) = uint8(v562)
	v564 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v564 == int32(2) {
		goto L164
	} else {
		goto L180
	}
L180:
	;
	v1134 = v66
	v1135 = v67
	v1142 = v74
	v1144 = v76
	goto L15
L181:
	;
	if v569 != int32(105) {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v573 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)) = uint8(v573)
	v575 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v575 == int32(0) {
		v1134 = v66
		v1135 = v67
		v1142 = v74
		v1144 = v76
		goto L15
	} else {
		goto L185
	}
L183:
	;
	goto L184
L184:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v492)))
	if v578 != int32(35) {
		goto L187
	} else {
		goto L188
	}
L185:
	;
	goto L184
L186:
	;
	if v676 == int32(0) {
		goto L208
	} else {
		goto L209
	}
L187:
	;
	if v578 != int32(7) {
		v1134 = v66
		v1135 = v67
		v1142 = v74
		v1144 = v76
		goto L15
	} else {
		goto L190
	}
L188:
	;
	goto L189
L189:
	;
	v663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v492)+20)))
	if v663 != 0 {
		v1134 = v66
		v1135 = v67
		v1142 = v74
		v1144 = v76
		goto L15
	} else {
		goto L206
	}
L190:
	;
	v583 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v492)+24)))
	if v583 != 0 {
		goto L72
	} else {
		goto L191
	}
L191:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v492)+20))
	v585 = F_pg_detoast_datum(m, v584)
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L4
	} else {
		goto L192
	}
L192:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v585)+12))
	F_get_typlenbyvalalign(m, v587, v26+int32(224), v26+int32(223), v26+int32(222))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L4
	} else {
		goto L193
	}
L193:
	;
	v597 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+224)))
	v598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+223)))
	v599 = int32(*(*int8)(unsafe.Add(mBase, uint32(v26)+222)))
	F_deconstruct_array(m, v585, v597, v598, v599, v26+int32(240), v26+int32(276), v26+int32(228))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L4
	} else {
		goto L194
	}
L194:
	;
	v608 = int32(0)
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v26)+228))
	if v609 <= v608 {
		v676 = v608
		goto L186
	} else {
		goto L195
	}
L195:
	;
	v619 = int32(0)
	v622 = v609
	v624 = v608
	goto L196
L196:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v26)+276))
	v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v636+v619))))
	if v638 == int32(1) {
		goto L199
	} else {
		goto L200
	}
L197:
	;
	v676 = v659
	goto L186
L198:
	;
	v661 = v619 + int32(1)
	if v661 < v658 {
		v619 = v661
		v622 = v658
		v624 = v659
		goto L196
	} else {
		goto L205
	}
L199:
	;
	v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+20)))
	if v641 != 0 {
		v658 = v622
		v659 = v624
		goto L198
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v585)+12))
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v492)+12))
	v645 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+224)))
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v26)+240))
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v646+v619<<(uint(int32(2))%32))))
	v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+223)))
	v653 = F_makeConst(m, v642, int32(-1), v644, v645, v650, int32(0), v652)
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L4
	} else {
		goto L203
	}
L202:
	;
	goto L72
L203:
	;
	v655 = F_lappend(m, v624, v653)
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L4
	} else {
		goto L204
	}
L204:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v26)+228))
	v658 = v657
	v659 = v655
	goto L198
L205:
	;
	goto L197
L206:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v492)+16))
	v676 = v664
	goto L186
L207:
	;
	v798 = F_gen_partprune_steps_internal(m, l0, v782)
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L4
	} else {
		goto L226
	}
L208:
	;
	v782 = int32(0)
	goto L207
L209:
	;
	goto L210
L210:
	;
	v691 = int32(0)
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v676)+4))
	if v691 < v693 {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v702 = v691
	v703 = v691
	goto L214
L212:
	;
	v739 = v691
	goto L213
L213:
	;
	v755 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+20)))
	if v755 != int32(1) {
		v782 = v739
		goto L207
	} else {
		goto L219
	}
L214:
	;
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v676)+12))
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v719+v702<<(uint(int32(2))%32))))
	v724 = F_make_opclause(m, v491, v489, v723, v490)
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L4
	} else {
		goto L216
	}
L215:
	;
	v739 = v726
	goto L213
L216:
	;
	v726 = F_lappend(m, v703, v724)
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L4
	} else {
		goto L217
	}
L217:
	;
	v729 = v702 + int32(1)
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v676)+4))
	if v729 < v730 {
		v702 = v729
		v703 = v726
		goto L214
	} else {
		goto L218
	}
L218:
	;
	goto L215
L219:
	;
	if v739 == int32(0) {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	v782 = int32(0)
	goto L207
L221:
	;
	goto L222
L222:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v739)+4))
	if v761 < int32(2) {
		v782 = v739
		goto L207
	} else {
		goto L223
	}
L223:
	;
	v766 = F_makeBoolExpr(m, int32(1), v739, int32(-1))
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L4
	} else {
		goto L224
	}
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v766
	*(*int32)(unsafe.Add(mBase, uint32(v26)+216)) = v766
	v773 = F_list_make1_impl(m, int32(1), v26+int32(56))
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L4
	} else {
		goto L225
	}
L225:
	;
	v782 = v773
	goto L207
L226:
	;
	v800 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	if v800 != 0 {
		goto L72
	} else {
		goto L227
	}
L227:
	;
	if v798 == int32(0) {
		v1134 = v66
		v1135 = v67
		v1142 = v74
		v1144 = v76
		goto L15
	} else {
		goto L228
	}
L228:
	;
	v1083 = v798
	goto L67
L229:
	;
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v805)))
	if v806 == int32(27) {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v805)+4))
	v810 = v809
	goto L232
L231:
	;
	v810 = v805
	goto L232
L232:
	;
	v811 = F_equal(m, v810, v297)
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L4
	} else {
		goto L233
	}
L233:
	;
	if v811 == int32(0) {
		goto L68
	} else {
		goto L234
	}
L234:
	;
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
	if v815 == int32(1) {
		goto L73
	} else {
		goto L235
	}
L235:
	;
	goto L74
L236:
	;
	v835 = F_bms_add_member(m, v67, v273)
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L4
	} else {
		goto L242
	}
L237:
	;
	if v822 == int32(0) {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v26+int32(80)+v291)))
	if v829 == int32(0) {
		goto L236
	} else {
		goto L241
	}
L239:
	;
	goto L240
L240:
	;
	v832 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)) = uint8(v832)
	v2350 = v26
	v2358 = int32(0)
	goto L1
L241:
	;
	goto L240
L242:
	;
	v1134 = v66
	v1135 = v835
	v1142 = v74
	v1144 = v76
	goto L15
L243:
	;
	if v841 != 0 {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v843 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)) = uint8(v843)
	v2350 = v26
	v2358 = int32(0)
	goto L1
L245:
	;
	goto L246
L246:
	;
	v846 = F_bms_add_member(m, v74, v273)
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L4
	} else {
		goto L247
	}
L247:
	;
	v1134 = v66
	v1135 = v67
	v1142 = v846
	v1144 = v76
	goto L15
L248:
	;
	v878 = int32(0)
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v87)+24))
	if v301 != v880 {
		v1050 = v878
		v1052 = v878
		v1053 = v364
		goto L69
	} else {
		goto L251
	}
L249:
	;
	goto L250
L250:
	;
	v884 = F_op_in_opfamily(m, v877, v305)
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L4
	} else {
		goto L254
	}
L251:
	;
	goto L250
L252:
	;
	v1041 = int32(0)
	v1050 = v1041
	v1052 = v1041
	v1053 = v364
	goto L69
L253:
	;
	v923 = int32(5)
	v924 = int32(0)
	v925 = F_op_strict(m, v877)
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L4
	} else {
		goto L266
	}
L254:
	;
	if v884 != 0 {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	F_get_op_opfamily_properties(m, v877, v305, int32(0), v26+int32(224), v26+int32(276), v26+int32(228))
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L4
	} else {
		goto L258
	}
L256:
	;
	goto L257
L257:
	;
	v895 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298))))
	if v895 != int32(108) {
		goto L70
	} else {
		goto L259
	}
L258:
	;
	v922 = v877
	goto L253
L259:
	;
	v898 = F_get_negator(m, v877)
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L4
	} else {
		goto L260
	}
L260:
	;
	if v898 == int32(0) {
		goto L252
	} else {
		goto L261
	}
L261:
	;
	v902 = int32(0)
	v904 = F_op_in_opfamily(m, v898, v305)
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L4
	} else {
		goto L262
	}
L262:
	;
	if v904 == int32(0) {
		v1050 = v902
		v1052 = v902
		v1053 = v364
		goto L69
	} else {
		goto L263
	}
L263:
	;
	F_get_op_opfamily_properties(m, v898, v305, int32(0), v26+int32(224), v26+int32(276), v26+int32(228))
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L4
	} else {
		goto L264
	}
L264:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v26)+224))
	if v917 != int32(3) {
		v1050 = v902
		v1052 = v902
		v1053 = v364
		goto L69
	} else {
		goto L265
	}
L265:
	;
	v922 = v898
	goto L253
L266:
	;
	if v925 == int32(0) {
		v1050 = v924
		v1052 = v923
		v1053 = v364
		goto L69
	} else {
		goto L267
	}
L267:
	;
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v875)))
	if v929 == int32(7) {
		goto L268
	} else {
		goto L269
	}
L268:
	;
	v967 = F_op_volatile(m, v877)
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L4
	} else {
		goto L285
	}
L269:
	;
	v932 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v932 == int32(0) {
		v1050 = v924
		v1052 = v923
		v1053 = v364
		goto L69
	} else {
		goto L270
	}
L270:
	;
	v935 = F_contain_var_clause(m, v875)
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L4
	} else {
		goto L271
	}
L271:
	;
	if v935 != 0 {
		v1050 = v924
		v1052 = v923
		v1053 = v364
		goto L69
	} else {
		goto L272
	}
L272:
	;
	v937 = F_contain_volatile_functions(m, v875)
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		goto L4
	} else {
		goto L273
	}
L273:
	;
	if v937 != 0 {
		v1050 = v924
		v1052 = v923
		v1053 = v364
		goto L69
	} else {
		goto L274
	}
L274:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+240)) = int32(0)
	v941 = *(*int32)(unsafe.Add(mBase, uint32(v875)))
	if v941 == int32(8) {
		goto L277
	} else {
		goto L278
	}
L275:
	;
	v965 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)) = uint8(v965)
	goto L268
L276:
	;
	if v957 == int32(0) {
		goto L275
	} else {
		goto L283
	}
L277:
	;
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v875)+4))
	if v944 != int32(1) {
		goto L275
	} else {
		goto L280
	}
L278:
	;
	goto L279
L279:
	;
	v954 = F_expression_tree_walker_impl(m, v875, int32(909), v26+int32(240))
	mBase = m.M
	v955 = m.ExcPending
	if v955 != 0 {
		goto L4
	} else {
		goto L282
	}
L280:
	;
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v875)+8))
	v949 = F_bms_add_member(m, int32(0), v948)
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L4
	} else {
		goto L281
	}
L281:
	;
	v957 = v949
	goto L276
L282:
	;
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v26)+240))
	v957 = v956
	goto L276
L283:
	;
	v960 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)) = uint8(v960)
	v962 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v962 == int32(2) {
		goto L268
	} else {
		goto L284
	}
L284:
	;
	v1050 = v924
	v1052 = v923
	v1053 = v364
	goto L69
L285:
	;
	if v967 != int32(105) {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	v971 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)) = uint8(v971)
	v973 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v973 == int32(0) {
		v1050 = v924
		v1052 = v923
		v1053 = v364
		goto L69
	} else {
		goto L289
	}
L287:
	;
	goto L288
L288:
	;
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v26)+228))
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v298)+8))
	v979 = *(*int32)(unsafe.Add(mBase, uint32(v977+v291)))
	if v976 == v979 {
		goto L291
	} else {
		goto L292
	}
L289:
	;
	goto L288
L290:
	;
	v1022 = F_palloc(m, int32(24))
	mBase = m.M
	v1023 = m.ExcPending
	if v1023 != 0 {
		goto L4
	} else {
		goto L304
	}
L291:
	;
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v298)+24))
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v981+v273*int32(28))+4))
	v1020 = v985
	goto L290
L292:
	;
	goto L293
L293:
	;
	v986 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298))))
	switch v986 - int32(104) {
	case 0:
		goto L297
	default:
		goto L296
	case 4, 10:
		goto L295
	}
L294:
	;
	if v1017 == int32(0) {
		goto L252
	} else {
		goto L303
	}
L295:
	;
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(v298)+4))
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(v1011+v291)))
	v1015 = F_get_opfamily_proc(m, v1013, v979, v976, int32(1))
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L4
	} else {
		goto L302
	}
L296:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L4
	} else {
		goto L299
	}
L297:
	;
	v989 = *(*int32)(unsafe.Add(mBase, uint32(v298)+4))
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v989+v291)))
	v993 = F_get_opfamily_proc(m, v991, v976, v976, int32(2))
	mBase = m.M
	v994 = m.ExcPending
	if v994 != 0 {
		goto L4
	} else {
		goto L298
	}
L298:
	;
	v1017 = v993
	goto L294
L299:
	;
	v999 = int32(*(*int8)(unsafe.Add(mBase, uint32(v298))))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v999
	F_errmsg_internal(m, int32(_a_F_gen_partprune_steps_internal_0), v26+int32(48))
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L4
	} else {
		goto L300
	}
L300:
	;
	F_errfinish(m, int32(_a_F_gen_partprune_steps_internal_1), int32(2138), int32(_a_F_gen_partprune_steps_internal_2))
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L4
	} else {
		goto L301
	}
L301:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L302:
	;
	v1017 = v1015
	goto L294
L303:
	;
	v1020 = v1017
	goto L290
L304:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1022))) = v273
	v1025 = int32(1)
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(v26)+224))
	v1028 = v884 ^ v1025
	*(*uint8)(unsafe.Add(mBase, uint32(v1022)+8)) = uint8(v1028)
	*(*int32)(unsafe.Add(mBase, uint32(v1022)+4)) = v922
	*(*int32)(unsafe.Add(mBase, uint32(v1022)+16)) = v1020
	*(*int32)(unsafe.Add(mBase, uint32(v1022)+12)) = v875
	if v884 != 0 {
		goto L305
	} else {
		goto L306
	}
L305:
	;
	v1034 = v1026
	goto L307
L306:
	;
	v1034 = int32(0)
	goto L307
L307:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1022)+20)) = v1034
	v1050 = v1022
	v1052 = v1025
	v1053 = v364
	goto L69
L308:
	;
	switch v1052 - int32(1) {
	case 0:
		v1100 = v1050
		goto L64
	default:
		v1134 = v66
		v1135 = v67
		v1142 = v74
		v1144 = v76
		goto L15
	case 2:
		v1083 = v1053
		goto L67
	}
L309:
	;
	v1134 = v66
	v1135 = v67
	v1142 = v74
	v1144 = v76
	goto L15
L310:
	;
	v1134 = v1098
	v1135 = v67
	v1142 = v74
	v1144 = v76
	goto L15
L311:
	;
	if v1109 != 0 {
		goto L312
	} else {
		goto L313
	}
L312:
	;
	v1111 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)) = uint8(v1111)
	v2350 = v26
	v2358 = int32(0)
	goto L1
L313:
	;
	goto L314
L314:
	;
	v1116 = v26 + int32(80) + v291
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v1116)))
	v1118 = F_lappend(m, v1117, v1100)
	mBase = m.M
	v1119 = m.ExcPending
	if v1119 != 0 {
		goto L4
	} else {
		goto L315
	}
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1116))) = v1118
	v1134 = v66
	v1135 = v67
	v1142 = v74
	v1144 = int32(1)
	goto L15
L316:
	;
	goto L14
L317:
	;
	v2158 = F_palloc0(m, int32(24))
	mBase = m.M
	v2159 = m.ExcPending
	if v2159 != 0 {
		goto L4
	} else {
		goto L482
	}
L318:
	;
	if v1171 == int32(0) {
		v2190 = v1161
		v2198 = v1169
		goto L8
	} else {
		goto L335
	}
L319:
	;
	v1174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	switch v1174 - int32(104) {
	case 0:
		goto L320
	default:
		v1215 = v1162
		goto L318
	case 4, 10:
		goto L317
	}
L320:
	;
	v1177 = int32(0)
	if v1162 == v1177 {
		goto L322
	} else {
		goto L323
	}
L321:
	;
	v1213 = int32(*(*int16)(unsafe.Add(mBase, uint32(v29)+2)))
	if v1212 == v1213 {
		goto L317
	} else {
		goto L334
	}
L322:
	;
	v1212 = int32(0)
	goto L321
L323:
	;
	goto L324
L324:
	;
	v1184 = int32(1)
	v1185 = *(*int32)(unsafe.Add(mBase, uint32(v1162)+4))
	if v1185 <= v1184 {
		goto L325
	} else {
		goto L326
	}
L325:
	;
	v1188 = v1184
	goto L327
L326:
	;
	v1188 = v1185
	goto L327
L327:
	;
	v1192 = int32(0)
	v1194 = v1177
	goto L328
L328:
	;
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(v1162+int32(8)+v1192<<(uint(int32(2))%32))))
	if v1200 != 0 {
		goto L330
	} else {
		goto L331
	}
L329:
	;
	v1212 = v1203
	goto L321
L330:
	;
	v1203 = v1194 + base.I32_popcnt(v1200)
	goto L332
L331:
	;
	v1203 = v1194
	goto L332
L332:
	;
	v1205 = v1192 + int32(1)
	if v1205 != v1188 {
		v1192 = v1205
		v1194 = v1203
		goto L328
	} else {
		goto L333
	}
L333:
	;
	goto L329
L334:
	;
	v1215 = v1162
	goto L318
L335:
	;
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(v1218)+232))
	v1220 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v1220
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v1220
	*(*int64)(unsafe.Add(mBase, uint32(v26)+256)) = v1220
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v1220
	v1228 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1219)+2)))
	if v1228 <= int32(0) {
		goto L341
	} else {
		goto L342
	}
L336:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2144 = m.ExcPending
	if v2144 != 0 {
		goto L4
	} else {
		goto L479
	}
L337:
	;
	v2139 = F_list_concat(m, v2128, v2129)
	mBase = m.M
	v2140 = m.ExcPending
	if v2140 != 0 {
		goto L4
	} else {
		goto L478
	}
L338:
	;
	v2116 = l0
	v2120 = v26
	v2128 = v1161
	v2129 = int32(0)
	goto L337
L339:
	;
	v1999 = *(*int32)(unsafe.Add(mBase, uint32(v26)+236))
	if v1999 == int32(0) {
		goto L338
	} else {
		goto L464
	}
L340:
	;
	v1476 = *(*int32)(unsafe.Add(mBase, uint32(v26)+256))
	v1477 = *(*int32)(unsafe.Add(mBase, uint32(v26)+248))
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(v26)+252))
	v1481 = l0
	v1485 = v26
	v1488 = v1477
	v1490 = v1478
	v1493 = v1161
	v1494 = int32(0)
	v1495 = int32(1)
	v1499 = v1476
	goto L379
L341:
	;
	v1450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1219))))
	switch v1450 - int32(104) {
	case 0:
		goto L339
	default:
		goto L336
	case 4, 10:
		goto L340
	}
L342:
	;
	v1241 = int32(0)
	goto L343
L343:
	;
	v1256 = v1241 << (uint(int32(2)) % 32)
	v1260 = *(*int32)(unsafe.Add(mBase, uint32(v1256+(v26+int32(80)))))
	v1263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1219))))
	if base.B2i32(v1260 == int32(0))&base.B2i32(v1263 == int32(114)) != 0 {
		goto L340
	} else {
		goto L345
	}
L344:
	;
	goto L341
L345:
	;
	if base.B2i32(v1263 != int32(104))|v1260 == int32(0) {
		goto L347
	} else {
		goto L348
	}
L346:
	;
	v1424 = v1241 + int32(1)
	v1425 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1219)+2)))
	if v1424 < v1425 {
		v1241 = v1424
		goto L343
	} else {
		goto L378
	}
L347:
	;
	v1272 = F_bms_is_member(m, v1241, v1215)
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
		goto L4
	} else {
		goto L350
	}
L348:
	;
	goto L349
L349:
	;
	if v1260 == int32(0) {
		goto L346
	} else {
		goto L352
	}
L350:
	;
	if v1272 != 0 {
		goto L346
	} else {
		goto L351
	}
L351:
	;
	goto L338
L352:
	;
	v1277 = int32(0)
	v1278 = *(*int32)(unsafe.Add(mBase, uint32(v1260)+4))
	if v1278 <= v1277 {
		goto L346
	} else {
		goto L353
	}
L353:
	;
	v1282 = int32(1)
	v1287 = v1277
	goto L355
L354:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1387 = m.ExcPending
	if v1387 != 0 {
		goto L4
	} else {
		goto L375
	}
L355:
	;
	v1304 = *(*int32)(unsafe.Add(mBase, uint32(v1260)+12))
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(v1304+v1287<<(uint(int32(2))%32))))
	v1310 = v1308 + int32(20)
	v1311 = *(*int32)(unsafe.Add(mBase, uint32(v1308)+20))
	if v1311 == int32(0) {
		goto L357
	} else {
		goto L358
	}
L356:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1374 = m.ExcPending
	if v1374 != 0 {
		goto L4
	} else {
		goto L372
	}
L357:
	;
	v1314 = *(*int32)(unsafe.Add(mBase, uint32(v1308)+4))
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(v1219)+4))
	v1317 = *(*int32)(unsafe.Add(mBase, uint32(v1315+v1256)))
	F_get_op_opfamily_properties(m, v1314, v1317, int32(0), v1310, v26+int32(284), v26+int32(280))
	mBase = m.M
	v1324 = m.ExcPending
	if v1324 != 0 {
		goto L4
	} else {
		goto L360
	}
L358:
	;
	goto L359
L359:
	;
	v1325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1219))))
	switch v1325 - int32(104) {
	case 0:
		goto L364
	default:
		goto L354
	case 4, 10:
		goto L365
	}
L360:
	;
	goto L359
L361:
	;
	goto L356
L362:
	;
	v1368 = v1287 + int32(1)
	v1369 = *(*int32)(unsafe.Add(mBase, uint32(v1260)+4))
	if v1368 < v1369 {
		v1282 = int32(0)
		v1287 = v1368
		goto L355
	} else {
		goto L371
	}
L363:
	;
	v1361 = v1287 + int32(1)
	v1362 = *(*int32)(unsafe.Add(mBase, uint32(v1260)+4))
	if v1361 < v1362 {
		v1287 = v1361
		goto L355
	} else {
		goto L369
	}
L364:
	;
	v1344 = *(*int32)(unsafe.Add(mBase, uint32(v1310)))
	if v1344 != int32(1) {
		goto L361
	} else {
		goto L367
	}
L365:
	;
	v1329 = v26 + int32(240)
	v1330 = *(*int32)(unsafe.Add(mBase, uint32(v1308)+20))
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(v1329+v1330<<(uint(int32(2))%32))))
	v1335 = F_lappend(m, v1334, v1308)
	mBase = m.M
	v1336 = m.ExcPending
	if v1336 != 0 {
		goto L4
	} else {
		goto L366
	}
L366:
	;
	v1337 = *(*int32)(unsafe.Add(mBase, uint32(v1308)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1337<<(uint(int32(2))%32)+v1329))) = v1335
	switch v1337 - int32(1) {
	case 0, 4:
		goto L362
	default:
		goto L363
	}
L367:
	;
	v1347 = *(*int32)(unsafe.Add(mBase, uint32(v26)+236))
	v1348 = F_lappend(m, v1347, v1308)
	mBase = m.M
	v1349 = m.ExcPending
	if v1349 != 0 {
		goto L4
	} else {
		goto L368
	}
L368:
	;
	v1352 = *(*int32)(unsafe.Add(mBase, uint32(v1308)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v26+int32(232)+v1352<<(uint(int32(2))%32)))) = v1348
	goto L363
L369:
	;
	if v1282&int32(1) != 0 {
		goto L346
	} else {
		goto L370
	}
L370:
	;
	goto L341
L371:
	;
	goto L341
L372:
	;
	F_errmsg_internal(m, int32(_a_F_gen_partprune_steps_internal_3), int32(0))
	mBase = m.M
	v1378 = m.ExcPending
	if v1378 != 0 {
		goto L4
	} else {
		goto L373
	}
L373:
	;
	F_errfinish(m, int32(_a_F_gen_partprune_steps_internal_1), int32(1480), int32(_a_F_gen_partprune_steps_internal_4))
	mBase = m.M
	v1383 = m.ExcPending
	if v1383 != 0 {
		goto L4
	} else {
		goto L374
	}
L374:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L375:
	;
	v1388 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1219))))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v1388
	F_errmsg_internal(m, int32(_a_F_gen_partprune_steps_internal_0), v26+int32(32))
	mBase = m.M
	v1394 = m.ExcPending
	if v1394 != 0 {
		goto L4
	} else {
		goto L376
	}
L376:
	;
	F_errfinish(m, int32(_a_F_gen_partprune_steps_internal_1), int32(1487), int32(_a_F_gen_partprune_steps_internal_4))
	mBase = m.M
	v1399 = m.ExcPending
	if v1399 != 0 {
		goto L4
	} else {
		goto L377
	}
L377:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L378:
	;
	goto L344
L379:
	;
	v1509 = *(*int32)(unsafe.Add(mBase, uint32(v1485+int32(240)+v1495<<(uint(int32(2))%32))))
	if v1509 == int32(0) {
		v1985 = v1494
		goto L381
	} else {
		goto L382
	}
L380:
	;
	v2116 = v1481
	v2120 = v1485
	v2128 = v1493
	v2129 = v1985
	goto L337
L381:
	;
	v1996 = v1495 + int32(1)
	if v1996 != int32(6) {
		v1494 = v1985
		v1495 = v1996
		goto L379
	} else {
		goto L463
	}
L382:
	;
	v1512 = *(*int32)(unsafe.Add(mBase, uint32(v1509)+4))
	if v1512 <= int32(0) {
		v1985 = v1494
		goto L381
	} else {
		goto L383
	}
L383:
	;
	v1528 = int32(0)
	v1531 = v1494
	goto L384
L384:
	;
	v1541 = *(*int32)(unsafe.Add(mBase, uint32(v1509)+12))
	v1545 = *(*int32)(unsafe.Add(mBase, uint32(v1541+v1528<<(uint(int32(2))%32))))
	v1546 = *(*int32)(unsafe.Add(mBase, uint32(v1545)))
	if v1546 == int32(0) {
		goto L387
	} else {
		goto L388
	}
L385:
	;
	v1985 = v1966
	goto L381
L386:
	;
	v1966 = F_list_concat(m, v1531, v1965)
	mBase = m.M
	v1967 = m.ExcPending
	if v1967 != 0 {
		goto L4
	} else {
		goto L461
	}
L387:
	;
	v1549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1545)+8)))
	v1550 = *(*int32)(unsafe.Add(mBase, uint32(v1545)+16))
	v1551 = *(*int32)(unsafe.Add(mBase, uint32(v1545)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1485)+12)) = v1551
	*(*int32)(unsafe.Add(mBase, uint32(v1485)+284)) = v1551
	v1557 = F_list_make1_impl(m, int32(1), v1485+int32(12))
	mBase = m.M
	v1558 = m.ExcPending
	if v1558 != 0 {
		goto L4
	} else {
		goto L390
	}
L388:
	;
	goto L389
L389:
	;
	v1594 = int32(0)
	if v1490 != 0 {
		goto L398
	} else {
		goto L399
	}
L390:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1485)+8)) = v1550
	*(*int32)(unsafe.Add(mBase, uint32(v1485)+280)) = v1550
	v1564 = F_list_make1_impl(m, int32(472), v1485+int32(8))
	mBase = m.M
	v1565 = m.ExcPending
	if v1565 != 0 {
		goto L4
	} else {
		goto L391
	}
L391:
	;
	v1567 = F_palloc0(m, int32(24))
	mBase = m.M
	v1568 = m.ExcPending
	if v1568 != 0 {
		goto L4
	} else {
		goto L392
	}
L392:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1567))) = int32(377)
	v1571 = *(*int32)(unsafe.Add(mBase, uint32(v1481)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1481)+16)) = v1571 + int32(1)
	v1575 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1567)+20)) = v1575
	*(*int32)(unsafe.Add(mBase, uint32(v1567)+16)) = v1564
	*(*int32)(unsafe.Add(mBase, uint32(v1567)+12)) = v1557
	if v1549 != 0 {
		goto L393
	} else {
		goto L394
	}
L393:
	;
	v1580 = v1575
	goto L395
L394:
	;
	v1580 = v1495
	goto L395
L395:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1567)+8)) = uint16(v1580)
	*(*int32)(unsafe.Add(mBase, uint32(v1567)+4)) = v1571
	v1583 = *(*int32)(unsafe.Add(mBase, uint32(v1481)+8))
	v1584 = F_lappend(m, v1583, v1567)
	mBase = m.M
	v1585 = m.ExcPending
	if v1585 != 0 {
		goto L4
	} else {
		goto L396
	}
L396:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1481)+8)) = v1584
	*(*int32)(unsafe.Add(mBase, uint32(v1485)+4)) = v1567
	*(*int32)(unsafe.Add(mBase, uint32(v1485)+276)) = v1567
	v1592 = F_list_make1_impl(m, int32(1), v1485+int32(4))
	mBase = m.M
	v1593 = m.ExcPending
	if v1593 != 0 {
		goto L4
	} else {
		goto L397
	}
L397:
	;
	v1965 = v1592
	goto L386
L398:
	;
	v1596 = *(*int32)(unsafe.Add(mBase, uint32(v1490)+12))
	v1597 = v1596
	goto L400
L399:
	;
	v1597 = v1594
	goto L400
L400:
	;
	if v1488 != 0 {
		goto L401
	} else {
		goto L402
	}
L401:
	;
	v1598 = *(*int32)(unsafe.Add(mBase, uint32(v1488)+12))
	v1599 = v1598
	goto L403
L402:
	;
	v1599 = v1594
	goto L403
L403:
	;
	v1600 = int32(0)
	if v1499 != 0 {
		goto L404
	} else {
		goto L405
	}
L404:
	;
	v1602 = *(*int32)(unsafe.Add(mBase, uint32(v1499)+12))
	v1603 = v1602
	goto L406
L405:
	;
	v1603 = v1600
	goto L406
L406:
	;
	v1604 = int32(0)
	if v1604 < v1546 {
		goto L407
	} else {
		goto L408
	}
L407:
	;
	v1609 = v1603
	v1610 = v1604
	v1612 = v1599
	v1613 = v1600
	v1615 = v1597
	goto L410
L408:
	;
	v1919 = v1600
	goto L409
L409:
	;
	v1936 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1545)+8)))
	v1937 = *(*int32)(unsafe.Add(mBase, uint32(v1545)+12))
	v1938 = *(*int32)(unsafe.Add(mBase, uint32(v1545)+16))
	v1940 = F_get_steps_using_prefix(m, v1481, v1495, v1936, v1937, v1938, int32(0), v1919)
	mBase = m.M
	v1941 = m.ExcPending
	if v1941 != 0 {
		goto L4
	} else {
		goto L460
	}
L410:
	;
	v1630 = int32(0)
	if v1615 == v1630 {
		v1693 = v1613
		v1698 = v1630
		goto L413
	} else {
		goto L414
	}
L411:
	;
	v1919 = v1892
	goto L409
L412:
	;
	if base.Ui32(int32(2)) < base.Ui32(v1495) {
		v1801 = v1612
		v1802 = v1717
		v1807 = v1722
		goto L424
	} else {
		goto L425
	}
L413:
	;
	v1717 = v1693
	v1719 = int32(0)
	v1722 = v1698
	goto L412
L414:
	;
	v1633 = *(*int32)(unsafe.Add(mBase, uint32(v1490)+12))
	v1634 = v1615 - v1633
	v1636 = v1634 >> (uint(int32(2)) % 32)
	v1637 = *(*int32)(unsafe.Add(mBase, uint32(v1490)+4))
	if v1637 <= v1636 {
		v1693 = v1613
		v1698 = v1630
		goto L413
	} else {
		goto L415
	}
L415:
	;
	v1639 = *(*int32)(unsafe.Add(mBase, uint32(v1490)+12))
	v1640 = v1639 + v1634
	v1641 = *(*int32)(unsafe.Add(mBase, uint32(v1640)))
	v1642 = *(*int32)(unsafe.Add(mBase, uint32(v1641)))
	if v1642 != v1610 {
		v1717 = v1613
		v1719 = v1640
		v1722 = v1630
		goto L412
	} else {
		goto L416
	}
L416:
	;
	v1644 = int32(1)
	v1645 = F_lappend(m, v1613, v1641)
	mBase = m.M
	v1646 = m.ExcPending
	if v1646 != 0 {
		goto L4
	} else {
		goto L417
	}
L417:
	;
	v1648 = v1636 + int32(1)
	v1649 = *(*int32)(unsafe.Add(mBase, uint32(v1490)+4))
	if v1649 <= v1648 {
		v1693 = v1645
		v1698 = v1644
		goto L413
	} else {
		goto L418
	}
L418:
	;
	v1652 = v1648
	v1657 = v1645
	goto L419
L419:
	;
	v1674 = *(*int32)(unsafe.Add(mBase, uint32(v1490)+12))
	v1677 = v1674 + v1652<<(uint(int32(2))%32)
	v1678 = *(*int32)(unsafe.Add(mBase, uint32(v1677)))
	v1679 = *(*int32)(unsafe.Add(mBase, uint32(v1678)))
	if v1679 != v1610 {
		v1717 = v1657
		v1719 = v1677
		v1722 = v1644
		goto L412
	} else {
		goto L421
	}
L420:
	;
	v1693 = v1681
	v1698 = v1644
	goto L413
L421:
	;
	v1681 = F_lappend(m, v1657, v1678)
	mBase = m.M
	v1682 = m.ExcPending
	if v1682 != 0 {
		goto L4
	} else {
		goto L422
	}
L422:
	;
	v1684 = v1652 + int32(1)
	v1685 = *(*int32)(unsafe.Add(mBase, uint32(v1490)+4))
	if v1684 < v1685 {
		v1652 = v1684
		v1657 = v1681
		goto L419
	} else {
		goto L423
	}
L423:
	;
	goto L420
L424:
	;
	if v1495&int32(6) != int32(4) {
		v1882 = v1609
		goto L443
	} else {
		goto L444
	}
L425:
	;
	if v1612 == int32(0) {
		goto L426
	} else {
		goto L427
	}
L426:
	;
	v1801 = int32(0)
	v1802 = v1717
	v1807 = v1722
	goto L424
L427:
	;
	goto L428
L428:
	;
	v1739 = *(*int32)(unsafe.Add(mBase, uint32(v1488)+12))
	v1740 = v1612 - v1739
	v1742 = v1740 >> (uint(int32(2)) % 32)
	v1743 = *(*int32)(unsafe.Add(mBase, uint32(v1488)+4))
	if v1743 <= v1742 {
		goto L429
	} else {
		goto L430
	}
L429:
	;
	v1801 = int32(0)
	v1802 = v1717
	v1807 = v1722
	goto L424
L430:
	;
	goto L431
L431:
	;
	v1746 = *(*int32)(unsafe.Add(mBase, uint32(v1488)+12))
	v1747 = v1746 + v1740
	v1748 = *(*int32)(unsafe.Add(mBase, uint32(v1747)))
	v1749 = *(*int32)(unsafe.Add(mBase, uint32(v1748)))
	if v1749 != v1610 {
		v1801 = v1747
		v1802 = v1717
		v1807 = v1722
		goto L424
	} else {
		goto L432
	}
L432:
	;
	v1751 = int32(1)
	v1752 = F_lappend(m, v1717, v1748)
	mBase = m.M
	v1753 = m.ExcPending
	if v1753 != 0 {
		goto L4
	} else {
		goto L433
	}
L433:
	;
	v1755 = v1742 + int32(1)
	v1756 = *(*int32)(unsafe.Add(mBase, uint32(v1488)+4))
	if v1756 <= v1755 {
		goto L434
	} else {
		goto L435
	}
L434:
	;
	v1801 = int32(0)
	v1802 = v1752
	v1807 = v1751
	goto L424
L435:
	;
	goto L436
L436:
	;
	v1760 = v1755
	v1765 = v1752
	goto L437
L437:
	;
	v1782 = *(*int32)(unsafe.Add(mBase, uint32(v1488)+12))
	v1785 = v1782 + v1760<<(uint(int32(2))%32)
	v1786 = *(*int32)(unsafe.Add(mBase, uint32(v1785)))
	v1787 = *(*int32)(unsafe.Add(mBase, uint32(v1786)))
	if v1787 != v1610 {
		v1801 = v1785
		v1802 = v1765
		v1807 = v1751
		goto L424
	} else {
		goto L439
	}
L438:
	;
	v1801 = int32(0)
	v1802 = v1789
	v1807 = v1751
	goto L424
L439:
	;
	v1789 = F_lappend(m, v1765, v1786)
	mBase = m.M
	v1790 = m.ExcPending
	if v1790 != 0 {
		goto L4
	} else {
		goto L440
	}
L440:
	;
	v1792 = v1760 + int32(1)
	v1793 = *(*int32)(unsafe.Add(mBase, uint32(v1488)+4))
	if v1792 < v1793 {
		v1760 = v1792
		v1765 = v1789
		goto L437
	} else {
		goto L441
	}
L441:
	;
	goto L438
L442:
	;
	v1910 = v1610 + int32(1)
	v1911 = *(*int32)(unsafe.Add(mBase, uint32(v1545)))
	if v1910 < v1911 {
		v1609 = v1888
		v1610 = v1910
		v1612 = v1801
		v1613 = v1892
		v1615 = v1719
		goto L410
	} else {
		goto L459
	}
L443:
	;
	if v1807 == int32(0) {
		v1985 = v1531
		goto L381
	} else {
		goto L458
	}
L444:
	;
	if v1609 == int32(0) {
		goto L446
	} else {
		goto L447
	}
L445:
	;
	v1850 = v1834
	v1855 = v1837
	goto L453
L446:
	;
	if v1807 != 0 {
		v1888 = int32(0)
		v1892 = v1802
		goto L442
	} else {
		goto L452
	}
L447:
	;
	v1823 = *(*int32)(unsafe.Add(mBase, uint32(v1499)+12))
	v1824 = v1609 - v1823
	v1826 = v1824 >> (uint(int32(2)) % 32)
	v1827 = *(*int32)(unsafe.Add(mBase, uint32(v1499)+4))
	if v1827 <= v1826 {
		goto L446
	} else {
		goto L448
	}
L448:
	;
	v1829 = *(*int32)(unsafe.Add(mBase, uint32(v1499)+12))
	v1830 = v1829 + v1824
	v1831 = *(*int32)(unsafe.Add(mBase, uint32(v1830)))
	v1832 = *(*int32)(unsafe.Add(mBase, uint32(v1831)))
	if v1832 != v1610 {
		v1882 = v1830
		goto L443
	} else {
		goto L449
	}
L449:
	;
	v1834 = F_lappend(m, v1802, v1831)
	mBase = m.M
	v1835 = m.ExcPending
	if v1835 != 0 {
		goto L4
	} else {
		goto L450
	}
L450:
	;
	v1837 = v1826 + int32(1)
	v1838 = *(*int32)(unsafe.Add(mBase, uint32(v1499)+4))
	if v1837 < v1838 {
		goto L445
	} else {
		goto L451
	}
L451:
	;
	v1888 = int32(0)
	v1892 = v1834
	goto L442
L452:
	;
	v1985 = v1531
	goto L381
L453:
	;
	v1867 = *(*int32)(unsafe.Add(mBase, uint32(v1499)+12))
	v1870 = v1867 + v1855<<(uint(int32(2))%32)
	v1871 = *(*int32)(unsafe.Add(mBase, uint32(v1870)))
	v1872 = *(*int32)(unsafe.Add(mBase, uint32(v1871)))
	if v1872 != v1610 {
		v1888 = v1870
		v1892 = v1850
		goto L442
	} else {
		goto L455
	}
L454:
	;
	v1888 = int32(0)
	v1892 = v1874
	goto L442
L455:
	;
	v1874 = F_lappend(m, v1850, v1871)
	mBase = m.M
	v1875 = m.ExcPending
	if v1875 != 0 {
		goto L4
	} else {
		goto L456
	}
L456:
	;
	v1877 = v1855 + int32(1)
	v1878 = *(*int32)(unsafe.Add(mBase, uint32(v1499)+4))
	if v1877 < v1878 {
		v1850 = v1874
		v1855 = v1877
		goto L453
	} else {
		goto L457
	}
L457:
	;
	goto L454
L458:
	;
	v1888 = v1882
	v1892 = v1802
	goto L442
L459:
	;
	goto L411
L460:
	;
	v1965 = v1940
	goto L386
L461:
	;
	v1969 = v1528 + int32(1)
	v1970 = *(*int32)(unsafe.Add(mBase, uint32(v1509)+4))
	if v1969 < v1970 {
		v1528 = v1969
		v1531 = v1966
		goto L384
	} else {
		goto L462
	}
L462:
	;
	goto L385
L463:
	;
	goto L380
L464:
	;
	v2002 = int32(0)
	v2003 = *(*int32)(unsafe.Add(mBase, uint32(v1999)+4))
	if v2003 <= v2002 {
		goto L338
	} else {
		goto L465
	}
L465:
	;
	v2006 = *(*int32)(unsafe.Add(mBase, uint32(v1999)+12))
	v2012 = *(*int32)(unsafe.Add(mBase, uint32(v2006+v2003<<(uint(int32(2))%32)-int32(4))))
	v2013 = *(*int32)(unsafe.Add(mBase, uint32(v2012)))
	v2018 = v2002
	v2023 = int32(0)
	goto L466
L466:
	;
	v2038 = *(*int32)(unsafe.Add(mBase, uint32(v1999)+12))
	v2042 = *(*int32)(unsafe.Add(mBase, uint32(v2038+v2018<<(uint(int32(2))%32))))
	v2043 = *(*int32)(unsafe.Add(mBase, uint32(v2042)))
	if v2013 == v2043 {
		goto L468
	} else {
		goto L469
	}
L467:
	;
	goto L338
L468:
	;
	v2049 = v2018
	v2059 = int32(0)
	goto L471
L469:
	;
	goto L470
L470:
	;
	v2086 = F_lappend(m, v2023, v2042)
	mBase = m.M
	v2087 = m.ExcPending
	if v2087 != 0 {
		goto L4
	} else {
		goto L476
	}
L471:
	;
	v2071 = *(*int32)(unsafe.Add(mBase, uint32(v1999)+12))
	v2075 = *(*int32)(unsafe.Add(mBase, uint32(v2071+v2049<<(uint(int32(2))%32))))
	v2076 = *(*int32)(unsafe.Add(mBase, uint32(v2075)+12))
	v2077 = *(*int32)(unsafe.Add(mBase, uint32(v2075)+16))
	v2078 = F_get_steps_using_prefix(m, l0, int32(1), int32(0), v2076, v2077, v1215, v2023)
	mBase = m.M
	v2079 = m.ExcPending
	if v2079 != 0 {
		goto L4
	} else {
		goto L473
	}
L472:
	;
	v2116 = l0
	v2120 = v26
	v2128 = v1161
	v2129 = v2080
	goto L337
L473:
	;
	v2080 = F_list_concat(m, v2059, v2078)
	mBase = m.M
	v2081 = m.ExcPending
	if v2081 != 0 {
		goto L4
	} else {
		goto L474
	}
L474:
	;
	v2083 = v2049 + int32(1)
	v2084 = *(*int32)(unsafe.Add(mBase, uint32(v1999)+4))
	if v2083 < v2084 {
		v2049 = v2083
		v2059 = v2080
		goto L471
	} else {
		goto L475
	}
L475:
	;
	goto L472
L476:
	;
	v2089 = v2018 + int32(1)
	v2090 = *(*int32)(unsafe.Add(mBase, uint32(v1999)+4))
	if v2089 < v2090 {
		v2018 = v2089
		v2023 = v2086
		goto L466
	} else {
		goto L477
	}
L477:
	;
	goto L467
L478:
	;
	v2261 = v2116
	v2265 = v2120
	v2273 = v2139
	goto L7
L479:
	;
	v2145 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1219))))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v2145
	F_errmsg_internal(m, int32(_a_F_gen_partprune_steps_internal_0), v26+int32(16))
	mBase = m.M
	v2151 = m.ExcPending
	if v2151 != 0 {
		goto L4
	} else {
		goto L480
	}
L480:
	;
	F_errfinish(m, int32(_a_F_gen_partprune_steps_internal_1), int32(1768), int32(_a_F_gen_partprune_steps_internal_4))
	mBase = m.M
	v2156 = m.ExcPending
	if v2156 != 0 {
		goto L4
	} else {
		goto L481
	}
L481:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L482:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2158))) = int32(377)
	v2162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2162 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2158)+20)) = v1162
	*(*int64)(unsafe.Add(mBase, uint32(v2158)+12)) = int64(0)
	v2169 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v2158)+8)) = uint16(v2169)
	*(*int32)(unsafe.Add(mBase, uint32(v2158)+4)) = v2162
	v2172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2173 = F_lappend(m, v2172, v2158)
	mBase = m.M
	v2174 = m.ExcPending
	if v2174 != 0 {
		goto L4
	} else {
		goto L483
	}
L483:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v2173
	v2176 = F_lappend(m, v1161, v2158)
	mBase = m.M
	v2177 = m.ExcPending
	if v2177 != 0 {
		goto L4
	} else {
		goto L484
	}
L484:
	;
	v2261 = l0
	v2265 = v26
	v2273 = v2176
	goto L7
L485:
	;
	v2237 = int32(*(*int16)(unsafe.Add(mBase, uint32(v29)+2)))
	if v2236 != v2237 {
		v2261 = l0
		v2265 = v26
		v2273 = v2190
		goto L7
	} else {
		goto L498
	}
L486:
	;
	v2236 = int32(0)
	goto L485
L487:
	;
	goto L488
L488:
	;
	v2208 = int32(1)
	v2209 = *(*int32)(unsafe.Add(mBase, uint32(v2198)+4))
	if v2209 <= v2208 {
		goto L489
	} else {
		goto L490
	}
L489:
	;
	v2212 = v2208
	goto L491
L490:
	;
	v2212 = v2209
	goto L491
L491:
	;
	v2216 = int32(0)
	v2218 = v2201
	goto L492
L492:
	;
	v2224 = *(*int32)(unsafe.Add(mBase, uint32(v2198+int32(8)+v2216<<(uint(int32(2))%32))))
	if v2224 != 0 {
		goto L494
	} else {
		goto L495
	}
L493:
	;
	v2236 = v2227
	goto L485
L494:
	;
	v2227 = v2218 + base.I32_popcnt(v2224)
	goto L496
L495:
	;
	v2227 = v2218
	goto L496
L496:
	;
	v2229 = v2216 + int32(1)
	if v2229 != v2212 {
		v2216 = v2229
		v2218 = v2227
		goto L492
	} else {
		goto L497
	}
L497:
	;
	goto L493
L498:
	;
	v2240 = F_palloc0(m, int32(24))
	mBase = m.M
	v2241 = m.ExcPending
	if v2241 != 0 {
		goto L4
	} else {
		goto L499
	}
L499:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2240))) = int32(377)
	v2244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2244 + int32(1)
	v2248 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2240)+20)) = v2248
	*(*int64)(unsafe.Add(mBase, uint32(v2240)+12)) = int64(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v2240)+8)) = uint16(v2248)
	*(*int32)(unsafe.Add(mBase, uint32(v2240)+4)) = v2244
	v2255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2256 = F_lappend(m, v2255, v2240)
	mBase = m.M
	v2257 = m.ExcPending
	if v2257 != 0 {
		goto L4
	} else {
		goto L500
	}
L500:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v2256
	v2259 = F_lappend(m, v2190, v2240)
	mBase = m.M
	v2260 = m.ExcPending
	if v2260 != 0 {
		goto L4
	} else {
		goto L501
	}
L501:
	;
	v2261 = l0
	v2265 = v26
	v2273 = v2259
	goto L7
L502:
	;
	v2350 = v2265
	v2358 = int32(0)
	goto L1
L503:
	;
	goto L504
L504:
	;
	v2287 = *(*int32)(unsafe.Add(mBase, uint32(v2273)+4))
	if v2287 < int32(2) {
		v2350 = v2265
		v2358 = v2273
		goto L1
	} else {
		goto L505
	}
L505:
	;
	v2290 = int32(0)
	v2295 = v2290
	v2298 = v2290
	goto L506
L506:
	;
	v2315 = *(*int32)(unsafe.Add(mBase, uint32(v2273)+12))
	v2319 = *(*int32)(unsafe.Add(mBase, uint32(v2315+v2295<<(uint(int32(2))%32))))
	v2320 = *(*int32)(unsafe.Add(mBase, uint32(v2319)+4))
	v2321 = F_lappend_int(m, v2298, v2320)
	mBase = m.M
	v2322 = m.ExcPending
	if v2322 != 0 {
		goto L4
	} else {
		goto L508
	}
L507:
	;
	v2328 = F_palloc0(m, int32(16))
	mBase = m.M
	v2329 = m.ExcPending
	if v2329 != 0 {
		goto L4
	} else {
		goto L510
	}
L508:
	;
	v2324 = v2295 + int32(1)
	v2325 = *(*int32)(unsafe.Add(mBase, uint32(v2273)+4))
	if v2324 < v2325 {
		v2295 = v2324
		v2298 = v2321
		goto L506
	} else {
		goto L509
	}
L509:
	;
	goto L507
L510:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2328))) = int32(378)
	v2332 = *(*int32)(unsafe.Add(mBase, uint32(v2261)+16))
	v2333 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2261)+16)) = v2332 + v2333
	*(*int32)(unsafe.Add(mBase, uint32(v2328)+12)) = v2321
	*(*int32)(unsafe.Add(mBase, uint32(v2328)+8)) = v2333
	*(*int32)(unsafe.Add(mBase, uint32(v2328)+4)) = v2332
	v2340 = *(*int32)(unsafe.Add(mBase, uint32(v2261)+8))
	v2341 = F_lappend(m, v2340, v2328)
	mBase = m.M
	v2342 = m.ExcPending
	if v2342 != 0 {
		goto L4
	} else {
		goto L511
	}
L511:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2261)+8)) = v2341
	v2344 = F_lappend(m, v2273, v2328)
	mBase = m.M
	v2345 = m.ExcPending
	if v2345 != 0 {
		goto L4
	} else {
		goto L512
	}
L512:
	;
	v2350 = v2265
	v2358 = v2344
	goto L1
}
func F_generate_setop_tlist(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
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
	var v76 int32
	_ = v76
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
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	v8 = int32(0)
	v16 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v16)
	v27 = v8
	v31 = v16
	v32 = v8
	goto L1
L1:
	;
	v34 = int32(0)
	if l0 == v34 {
		v44 = v34
		goto L3
	} else {
		goto L4
	}
L2:
	;
	return v32
L3:
	;
	v45 = int32(0)
	if l1 == v45 {
		v56 = v45
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v38 <= v27 {
		v44 = int32(0)
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v44 = v40 + v27<<(uint(int32(2))%32)
	goto L3
L6:
	;
	if l4 == int32(0) {
		v65 = v45
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v50 <= v27 {
		v56 = int32(0)
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v56 = v52 + v27<<(uint(int32(2))%32)
	goto L6
L9:
	;
	v66 = int32(0)
	if l5 == v66 {
		v75 = v66
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v59 <= v27 {
		v65 = v45
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v65 = v61 + v27<<(uint(int32(2))%32)
	goto L9
L12:
	;
	v76 = int32(0)
	if base.B2i32(v44 == v76)|base.B2i32(v56 == v76)|(base.B2i32(v65 == v76)|base.B2i32(v75 == v76)) == v76 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v69 <= v27 {
		v75 = v66
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	v75 = v71 + v27<<(uint(int32(2))%32)
	goto L12
L15:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	if l3 == int32(0) {
		v102 = v93
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	goto L2
L18:
	;
	v118 = F_exprType(m, v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L25
	} else {
		goto L30
	}
L19:
	;
	v103 = int32(*(*int16)(unsafe.Add(mBase, uint32(v92)+8)))
	v104 = F_exprType(m, v102)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L25
	} else {
		goto L26
	}
L20:
	;
	if v93 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v102 = int32(0)
	goto L19
L22:
	;
	goto L23
L23:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	if v99 == int32(7) {
		v117 = v93
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v102 = v93
	goto L19
L25:
	;
	return int32(0)
L26:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	v109 = F_exprTypmod(m, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	v112 = F_exprCollation(m, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v115 = F_makeVar(m, l2, v103, v104, v109, v112, int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L25
	} else {
		goto L29
	}
L29:
	;
	v117 = v115
	goto L18
L30:
	;
	if v118 != v91 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v123 = F_coerce_to_common_type(m, int32(0), v117, v91, int32(_a_F_generate_setop_tlist_0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L25
	} else {
		goto L34
	}
L32:
	;
	v127 = v117
	goto L33
L33:
	;
	v128 = F_exprCollation(m, v127)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L25
	} else {
		goto L35
	}
L34:
	;
	v125 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v125)
	v127 = v123
	goto L33
L35:
	;
	if v128 != v90 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v131 = F_exprType(m, v127)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L25
	} else {
		goto L39
	}
L37:
	;
	v142 = v127
	goto L38
L38:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v145 = F_pstrdup(m, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L25
	} else {
		goto L42
	}
L39:
	;
	v133 = F_exprTypmod(m, v127)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L25
	} else {
		goto L40
	}
L40:
	;
	v138 = F_applyRelabelType(m, v127, v131, v133, v90, int32(2), int32(-1), int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L25
	} else {
		goto L41
	}
L41:
	;
	v140 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v140)
	v142 = v138
	goto L38
L42:
	;
	v148 = F_makeTargetEntry(m, v142, base.I32_extend16_s(v31), v145, int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L25
	} else {
		goto L43
	}
L43:
	;
	v150 = int32(*(*int16)(unsafe.Add(mBase, uint32(v148)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v148)+16)) = v150
	v152 = int32(1)
	v156 = F_lappend(m, v32, v148)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L25
	} else {
		goto L44
	}
L44:
	;
	v27 = v27 + v152
	v31 = v31 + v152
	v32 = v156
	goto L1
}
func F_generate_subscripts(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
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
	var v75 int32
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
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	if v9 == int32(0) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v13 = F_DatumGetAnyArrayP(m, v12)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v18 = F_init_MultiFuncCall(m, l0)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
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
					v126 = m.ExcPending
					if v126 != 0 {
						return int32(0)
					} else {
						v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v127)+20)) = int32(2)
						v130 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v130)
						return int32(0)
					}
				} else {
					v33 = int32(0)
					if base.B2i32(base.Ui32(v17) <= base.Ui32(v27))&base.B2i32(v33 < v17) == v33 {
						F_end_MultiFuncCall(m, l0)
						mBase = m.M
						v126 = m.ExcPending
						if v126 != 0 {
							return int32(0)
						} else {
							v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v127)+20)) = int32(2)
							v130 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v130)
							return int32(0)
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
							return int32(0)
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
								v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								v79 = base.B2i32(v75 != int32(0))
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
									return v106
								} else {
									v111 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v111 - int32(1)
									return v111
								}
							} else {
								F_end_MultiFuncCall(m, l0)
								mBase = m.M
								v117 = m.ExcPending
								if v117 != 0 {
									return int32(0)
								} else {
									v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v118)+20)) = int32(2)
									v121 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v121)
									return int32(0)
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
				return v106
			} else {
				v111 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v111 - int32(1)
				return v111
			}
		} else {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v117 = m.ExcPending
			if v117 != 0 {
				return int32(0)
			} else {
				v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v118)+20)) = int32(2)
				v121 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v121)
				return int32(0)
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
						F_errfinish(m, int32(_a_F_generate_uuidv7_2), int32(638), int32(_a_F_generate_uuidv7_3))
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
	var v36 int32
	_ = v36
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
	var v69 int32
	_ = v69
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
	var v100 int32
	_ = v100
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
	var v152 int32
	_ = v152
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
	var v217 int32
	_ = v217
	var v233 float64
	_ = v233
	var v238 float64
	_ = v238
	var v241 int32
	_ = v241
	var v250 int32
	_ = v250
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
	var v335 float64
	_ = v335
	var v344 float64
	_ = v344
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v351 float64
	_ = v351
	var v354 float64
	_ = v354
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 float64
	_ = v364
	var v365 float64
	_ = v365
	var v366 float64
	_ = v366
	var v368 int32
	_ = v368
	var v371 float64
	_ = v371
	var v372 float64
	_ = v372
	var v376 float64
	_ = v376
	var v377 float64
	_ = v377
	var v381 float64
	_ = v381
	var v385 float64
	_ = v385
	var v390 float64
	_ = v390
	var v400 float64
	_ = v400
	var v403 float64
	_ = v403
	var v408 float64
	_ = v408
	var v409 float64
	_ = v409
	var v412 float64
	_ = v412
	var v415 float64
	_ = v415
	var v417 float64
	_ = v417
	var v418 int32
	_ = v418
	var v419 float64
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v425 float64
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v429 float64
	_ = v429
	var v435 float64
	_ = v435
	var v446 float64
	_ = v446
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
	v36 = v29
	v38 = v5
	v39 = v5
	goto L4
L4:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v52+v39<<(uint(int32(2))%32))))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	if v57 == int32(0) {
		v100 = v36
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
	if v117 < v100 {
		v36 = v100
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
		v100 = v36
		v102 = v38
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v69 = v60
	v70 = v38
	goto L9
L9:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v84+v69<<(uint(int32(2))%32))))
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
	v100 = v95
	v102 = v89
	goto L6
L11:
	;
	return
L12:
	;
	v92 = v69 + int32(1)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v92 < v93 {
		v69 = v92
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
	v217 = v125
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
	v152 = v145
	v154 = v145
	goto L22
L22:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v140)+12))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v167+v152<<(uint(int32(2))%32))))
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
	v188 = v152 + int32(1)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v140)+4))
	if v188 < v189 {
		v152 = v188
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
	v217 = v211
	goto L17
L32:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v309)+68))
	v311 = int32(0)
	v313 = F_clauselist_selectivity(m, l0, v217, v310, v311, v311)
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
	v250 = int32(0)
	v258 = v238
	goto L36
L36:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v265+v250<<(uint(int32(2))%32))))
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
	v286 = v250 + int32(1)
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	if v286 < v287 {
		v250 = v286
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
	v319 = *(*float64)(unsafe.Add(mBase, uint32(v318)+120))
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
	if base.Ui32(v332) < base.Ui32(int32(2)) {
		v344 = float64(1)
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	F_get_tablespace_page_costs(m, v345, v23+int32(8), int32(0))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L11
	} else {
		goto L59
	}
L57:
	;
	v335 = float64(1)
	if base.F64_gt(v325, v335) == int32(0) {
		v344 = v335
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v344 = base.F64_ceil(base.F64_div(base.F64_mul(v330, base.F64_convert_i32_u(v332)), v325))
	goto L56
L59:
	;
	v351 = base.F64_mul(l2, v302)
	if base.F64_gt(v351, float64(1)) != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v417 = F_index_other_operands_eval_cost(m, l0, v125)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L11
	} else {
		goto L82
	}
L61:
	;
	v354 = base.F64_mul(v351, v344)
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v360 = int32(1)
	if base.Ui32(v355) <= base.Ui32(v360) {
		goto L65
	} else {
		goto L66
	}
L62:
	;
	goto L63
L63:
	;
	v412 = *(*float64)(unsafe.Add(mBase, uint32(v23)+8))
	v415 = base.F64_mul(v344, v412)
	goto L60
L64:
	;
	v409 = *(*float64)(unsafe.Add(mBase, uint32(v23)+8))
	v415 = base.F64_div(base.F64_mul(v408, v409), l2)
	goto L60
L65:
	;
	v363 = v360
	goto L67
L66:
	;
	v363 = v355
	goto L67
L67:
	;
	v364 = base.F64_convert_i32_u(v363)
	v365 = base.F64_add(v364, v364)
	v366 = float64(1)
	v368 = *(*int32)(unsafe.Add(mBase, _c_F_genericcostestimate[0]))
	v371 = *(*float64)(unsafe.Add(mBase, uint32(l0)+288))
	v372 = base.F64_add(base.F64_convert_i32_u(v355), v371)
	if base.F64_gt(v372, v366) != 0 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v408 = v403
	goto L64
L69:
	;
	v376 = v372
	goto L71
L70:
	;
	v376 = v366
	goto L71
L71:
	;
	v377 = base.F64_div(base.F64_mul(v364, base.F64_convert_i32_s(v368)), v376)
	if base.F64_le(v377, float64(1)) != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v381 = v366
	goto L74
L73:
	;
	v381 = base.F64_ceil(v377)
	goto L74
L74:
	;
	if base.F64_le(v364, v381) != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v385 = base.F64_div(base.F64_mul(v354, v365), base.F64_add(v365, v354))
	if base.F64_ge(v385, v364) != 0 {
		v403 = v364
		goto L68
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v390 = base.F64_div(base.F64_mul(v365, v381), base.F64_sub(v365, v381))
	if base.F64_ge(v390, v354) != 0 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v408 = base.F64_ceil(v385)
	goto L64
L79:
	;
	v400 = base.F64_div(base.F64_mul(v354, v365), base.F64_add(v365, v354))
	goto L81
L80:
	;
	v400 = base.F64_add(v381, base.F64_div(base.F64_mul(base.F64_sub(v364, v381), base.F64_sub(v354, v390)), v364))
	goto L81
L81:
	;
	v403 = base.F64_ceil(v400)
	goto L68
L82:
	;
	v419 = F_index_other_operands_eval_cost(m, l0, v139)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
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
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	v423 = v422
	goto L86
L85:
	;
	v423 = int32(0)
	goto L86
L86:
	;
	v425 = *(*float64)(unsafe.Add(mBase, _c_F_genericcostestimate[1]))
	if v139 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	v427 = v426
	goto L89
L88:
	;
	v427 = int32(0)
	goto L89
L89:
	;
	v429 = *(*float64)(unsafe.Add(mBase, _c_F_genericcostestimate[2]))
	*(*float64)(unsafe.Add(mBase, uint32(l3)+40)) = v330
	*(*float64)(unsafe.Add(mBase, uint32(l3)+32)) = v344
	*(*int64)(unsafe.Add(mBase, uint32(l3)+24)) = int64(0)
	*(*float64)(unsafe.Add(mBase, uint32(l3)+16)) = v313
	v435 = base.F64_add(v417, v419)
	*(*float64)(unsafe.Add(mBase, uint32(l3))) = v435
	*(*float64)(unsafe.Add(mBase, uint32(l3)+8)) = base.F64_add(base.F64_mul(base.F64_mul(v302, v330), base.F64_add(v429, base.F64_mul(v425, base.F64_convert_i32_s(v423+v427)))), base.F64_add(v415, v435))
	v446 = *(*float64)(unsafe.Add(mBase, uint32(v23)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l3)+56)) = v302
	*(*float64)(unsafe.Add(mBase, uint32(l3)+48)) = v446
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
	var v106 float64
	_ = v106
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
	v106 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+8)) = v106
	return
}
func F_getQuadrant(m *base.Module, l0 int32, l1 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	v6 = F_DirectFunctionCall2Coll(m, int32(248), int32(0), l1, l0)
	v9 = m.ExcPending
	if v9 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L36
	}
L2:
	;
	return v75
L3:
	;
	v30 = F_DirectFunctionCall2Coll(m, int32(252), int32(0), l1, l0)
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L15
	}
L4:
	;
	return int32(0)
L5:
	;
	if v6 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v14 = F_DirectFunctionCall2Coll(m, int32(249), int32(0), l1, l0)
	v15 = m.ExcPending
	if v15 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v18 = int32(1)
	v21 = F_DirectFunctionCall2Coll(m, int32(250), int32(0), l1, l0)
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L11
	}
L9:
	;
	if v14 == int32(0) {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	if v21 != 0 {
		v75 = v18
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v25 = F_DirectFunctionCall2Coll(m, int32(251), int32(0), l1, l0)
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	if v25 != 0 {
		v75 = v18
		goto L2
	} else {
		goto L14
	}
L14:
	;
	goto L3
L15:
	;
	if v30 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v32 = int32(2)
	v35 = F_DirectFunctionCall2Coll(m, int32(250), int32(0), l1, l0)
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v44 = F_DirectFunctionCall2Coll(m, int32(252), int32(0), l1, l0)
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L24
	}
L19:
	;
	if v35 != 0 {
		v75 = v32
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v39 = F_DirectFunctionCall2Coll(m, int32(251), int32(0), l1, l0)
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	if v39 != 0 {
		v75 = v32
		goto L2
	} else {
		goto L22
	}
L22:
	;
	goto L18
L23:
	;
	v64 = F_DirectFunctionCall2Coll(m, int32(248), int32(0), l1, l0)
	v65 = m.ExcPending
	if v65 != 0 {
		goto L4
	} else {
		goto L32
	}
L24:
	;
	if v44 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v50 = F_DirectFunctionCall2Coll(m, int32(249), int32(0), l1, l0)
	v51 = m.ExcPending
	if v51 != 0 {
		goto L4
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v56 = F_DirectFunctionCall2Coll(m, int32(253), int32(0), l1, l0)
	v57 = m.ExcPending
	if v57 != 0 {
		goto L4
	} else {
		goto L30
	}
L28:
	;
	if v50 == int32(0) {
		goto L23
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	if v56 == int32(0) {
		goto L23
	} else {
		goto L31
	}
L31:
	;
	return int32(3)
L32:
	;
	if v64 == int32(0) {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v70 = F_DirectFunctionCall2Coll(m, int32(253), int32(0), l1, l0)
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	if v70 == int32(0) {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v75 = int32(4)
	goto L2
L36:
	;
	F_errmsg_internal(m, int32(_a_F_getQuadrant_0), int32(0))
	v84 = m.ExcPending
	if v84 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_getQuadrant_1), int32(77), int32(_a_F_getQuadrant_2))
	v89 = m.ExcPending
	if v89 != 0 {
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
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
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
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	v10 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(256)
	m.G0 = v16
	*(*int32)(unsafe.Add(mBase, uint32(v16)+180)) = v10
	*(*int32)(unsafe.Add(mBase, uint32(v16)+184)) = int32(6)
	v22 = F_GlobalVisHorizonKindForRel(m, l0)
	mBase = m.M
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22<<(uint(int32(2))%32))+uint32(_c_F_get_actual_variable_endpoint[0])))
	goto L1
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+224)) = v25
	v29 = int32(0)
	v32 = F_index_beginscan(m, l0, l1, v16+int32(184), v29, int32(1), v29)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	return int32(0)
L3:
	;
	v36 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+28)) = uint8(v36)
	v39 = int32(0)
	F_index_rescan(m, v32, l3, v36, v39, v39)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v43 = F_index_getnext_tid(m, v32, l2)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L2
	} else {
		goto L6
	}
L5:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v16)+180))
	if v159 != 0 {
		goto L39
	} else {
		goto L40
	}
L6:
	;
	if v43 == int32(0) {
		v157 = v10
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v51 = v43
	v59 = v10
	v60 = int32(-1)
	goto L10
L8:
	;
	v157 = int32(0)
	goto L5
L9:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v32)+44))
	if v89 != 0 {
		goto L26
	} else {
		goto L27
	}
L10:
	;
	v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+2)))
	v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51))))
	v65 = v61 | v62<<(uint(int32(16))%32)
	v68 = F_visibilitymap_get_status(m, l0, v65, v16+int32(180))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L2
	} else {
		goto L12
	}
L11:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
	m.T0[v86].(func(*base.Module, int32))(m, l6)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L2
	} else {
		goto L24
	}
L12:
	;
	if v68&int32(1) != 0 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v72 = F_index_fetch_heap(m, v32, l6)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	if v72 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	if v65 != v60 {
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
	v78 = v59 + int32(1)
	if int32(100) < v78 {
		goto L8
	} else {
		goto L21
	}
L19:
	;
	v81 = v59
	v82 = v60
	goto L20
L20:
	;
	v83 = F_index_getnext_tid(m, v32, l2)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L2
	} else {
		goto L22
	}
L21:
	;
	v81 = v78
	v82 = v65
	goto L20
L22:
	;
	if v83 != 0 {
		v51 = v83
		v59 = v81
		v60 = v82
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
	v128 = m.ExcPending
	if v128 != 0 {
		goto L2
	} else {
		goto L36
	}
L26:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+72)))
	if v90 != 0 {
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
	v115 = m.ExcPending
	if v115 != 0 {
		goto L2
	} else {
		goto L33
	}
L29:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v32)+48))
	F_index_deform_tuple(m, v89, v91, v16+int32(48), v16+int32(16))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L2
	} else {
		goto L30
	}
L30:
	;
	v98 = int32(1)
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+16)))
	if v99 == v98 {
		goto L25
	} else {
		goto L31
	}
L31:
	;
	v102 = int32(_a_F_get_actual_variable_endpoint_0)
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_get_actual_variable_endpoint[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_get_actual_variable_endpoint[1])) = l7
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v16)+48))
	v107 = F_datumCopy(m, v106, l5, l4)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = v107
	*(*int32)(unsafe.Add(mBase, _c_F_get_actual_variable_endpoint[1])) = v103
	v157 = v98
	goto L5
L33:
	;
	F_errmsg_internal(m, int32(_a_F_get_actual_variable_endpoint_1), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L2
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_get_actual_variable_endpoint_2), int32(_a_F_get_actual_variable_endpoint_3), int32(_a_F_get_actual_variable_endpoint_4))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v129 + int32(4)
	F_errmsg_internal(m, int32(_a_F_get_actual_variable_endpoint_5), v16)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_get_actual_variable_endpoint_2), int32(_a_F_get_actual_variable_endpoint_6), int32(_a_F_get_actual_variable_endpoint_4))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
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
	F_ReleaseBuffer(m, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L2
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	F_index_endscan(m, v32)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L2
	} else {
		goto L43
	}
L42:
	;
	goto L41
L43:
	;
	m.G0 = v16 + int32(256)
	return v157
}
func F_get_aggregate_argtypes(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v5 == int32(0) {
		return int32(0)
	} else {
		v10 = int32(0)
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
		if v10 < v11 {
			v14 = v10
			for {
				v19 = v14 << (uint(int32(2)) % 32)
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v21+v19)))
				*(*int32)(unsafe.Add(mBase, uint32(l1+v19))) = v23
				v26 = v14 + int32(1)
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
				if v26 < v27 {
					v14 = v26
					continue
				} else {
					break
				}
				break
			}
			v29 = v26
		} else {
			v29 = v10
		}
		return v29
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
func F_get_attoptions(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = F_SearchSysCache2(m, int32(7), l0, l1)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 != 0 {
			v19 = F_SysCacheGetAttr(m, int32(6), v10, int32(23), v7+int32(15))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
				if v21 == int32(0) {
					v26 = F_datumCopy(m, v19, int32(0), int32(-1))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v28 = v26
						F_ReleaseCatCache(m, v10)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							m.G0 = v7 + int32(16)
							return v28
						}
					}
				} else {
					v28 = int32(0)
					F_ReleaseCatCache(m, v10)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						m.G0 = v7 + int32(16)
						return v28
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l0
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
				F_errmsg_internal(m, int32(_a_F_get_attoptions_0), v7)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_get_attoptions_1), int32(1075), int32(_a_F_get_attoptions_2))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
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
func F_get_commutator(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13898(m, l0, int32(40))
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
		v28 = F_palloc(m, int32(296))
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
					v35 = int32(296)
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
										F_errfinish(m, int32(_a_F_get_controlfile_3), int32(108), int32(_a_F_get_controlfile_4))
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
									*(*int32)(unsafe.Add(mBase, uint32(v25)+56)) = int32(296)
									*(*int32)(unsafe.Add(mBase, uint32(v25)+52)) = v36
									*(*int32)(unsafe.Add(mBase, uint32(v25)+48)) = v16
									F_errmsg(m, int32(_a_F_get_controlfile_5), v23+int32(-16))
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_get_controlfile_3), int32(117), int32(_a_F_get_controlfile_4))
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
											F_errfinish(m, int32(_a_F_get_controlfile_3), int32(129), int32(_a_F_get_controlfile_4))
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
								v66 = m.Env.Pgmem_crc32c(m, v64, v28, int32(292))
								mBase = m.M
								v67 = *(*int32)(unsafe.Add(mBase, uint32(v28)+292))
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
											F_errfinish(m, int32(_a_F_get_controlfile_3), int32(168), int32(_a_F_get_controlfile_4))
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
								F_errfinish(m, int32(_a_F_get_controlfile_3), int32(94), int32(_a_F_get_controlfile_4))
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
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
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
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
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
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
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	v24 = m.G0
	v26 = v24 + int32(-64)
	m.G0 = v26
	v30 = F_table_open(m, int32(1262), int32(1))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
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
	v58 = v24 + int32(-48)
	F_ScanKeyInit(m, v58, int32(2), int32(3), int32(62), l0)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	F_relation_close(m, v30, int32(1))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L93
	}
L5:
	;
	goto L4
L6:
	;
	v65 = int32(1)
	v68 = F_systable_beginscan(m, v30, int32(2671), v65, int32(0), v65, v58)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v70 = F_systable_getnext(m, v68)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	if v70 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	F_systable_endscan(m, v68)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+22)))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v76+v77)))
	F_systable_endscan(m, v68)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	goto L5
L13:
	;
	F_LockSharedObject(m, int32(1262), v79, l1)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v86 = F_SearchSysCache1(m, int32(21), v79)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if v86 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v86)+16))
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+22)))
	v90 = v88 + v89
	v92 = v90 + int32(4)
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	if base.B2i32(v95 == int32(0))|base.B2i32(v95 != v98) != 0 {
		v116 = v95
		v117 = v98
		goto L20
	} else {
		goto L21
	}
L17:
	;
	goto L18
L18:
	;
	F_UnlockSharedObject(m, int32(1262), v79, l1)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L92
	}
L19:
	;
	if v116-v117 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L20:
	;
	goto L19
L21:
	;
	v101 = l0
	v102 = v92
	goto L22
L22:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+1)))
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+1)))
	if v106 == int32(0) {
		v116 = v106
		v117 = v105
		goto L20
	} else {
		goto L24
	}
L23:
	;
	v116 = v106
	v117 = v105
	goto L20
L24:
	;
	v109 = int32(1)
	if v106 == v105 {
		v101 = v101 + v109
		v102 = v102 + v109
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v79
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
	F_ReleaseCatCache(m, v86)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L91
	}
L29:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v90)+68))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v122
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
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v90)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v124
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
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+77)))
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v126)
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
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+79)))
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v128)
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
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+78)))
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v130)
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
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v90)+84))
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = v132
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
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v90)+88))
	*(*int32)(unsafe.Add(mBase, uint32(l9))) = v134
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
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v90)+92))
	*(*int32)(unsafe.Add(mBase, uint32(l10))) = v136
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
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+76)))
	*(*uint8)(unsafe.Add(mBase, uint32(l15))) = uint8(v138)
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
	v142 = F_SysCacheGetAttrNotNull(m, int32(21), v86, int32(13))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
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
	v144 = F_text_to_cstring(m, v142)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l11))) = v144
	goto L58
L61:
	;
	v149 = F_SysCacheGetAttrNotNull(m, int32(21), v86, int32(14))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
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
	v151 = F_text_to_cstring(m, v149)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l12))) = v151
	goto L63
L66:
	;
	v158 = F_SysCacheGetAttr(m, int32(21), v86, int32(15), v24+int32(-49))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
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
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+15)))
	if v160 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v164 = int32(0)
	goto L72
L71:
	;
	v162 = F_text_to_cstring(m, v158)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L73
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l13))) = v164
	goto L68
L73:
	;
	v164 = v162
	goto L72
L74:
	;
	v171 = F_SysCacheGetAttr(m, int32(21), v86, int32(16), v24+int32(-49))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
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
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+15)))
	if v173 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v177 = int32(0)
	goto L80
L79:
	;
	v175 = F_text_to_cstring(m, v171)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L81
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l14))) = v177
	goto L76
L81:
	;
	v177 = v175
	goto L80
L82:
	;
	v184 = F_SysCacheGetAttr(m, int32(21), v86, int32(17), v24+int32(-49))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	F_ReleaseCatCache(m, v86)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L90
	}
L85:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+15)))
	if v186 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v190 = int32(0)
	goto L88
L87:
	;
	v188 = F_text_to_cstring(m, v184)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L89
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l16))) = v190
	goto L84
L89:
	;
	v190 = v188
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
	m.G0 = v26 - int32(-64)
	return base.B2i32(v70 != int32(0))
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
func F_get_matching_partitions(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
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
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int64
	_ = v175
	var v177 int32
	_ = v177
	var v179 int64
	_ = v179
	var v181 int64
	_ = v181
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v363 int64
	_ = v363
	var v364 int32
	_ = v364
	var v365 int64
	_ = v365
	var v366 int64
	_ = v366
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v403 int32
	_ = v403
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v590 int32
	_ = v590
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v686 int32
	_ = v686
	var v703 int32
	_ = v703
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v776 int32
	_ = v776
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v801 int32
	_ = v801
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v843 int32
	_ = v843
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v855 int32
	_ = v855
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v876 int32
	_ = v876
	var v882 int32
	_ = v882
	var v887 int32
	_ = v887
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v910 int32
	_ = v910
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v920 int32
	_ = v920
	var v926 int32
	_ = v926
	var v931 int32
	_ = v931
	var v935 int32
	_ = v935
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v942 int32
	_ = v942
	var v944 int32
	_ = v944
	var v950 int32
	_ = v950
	var v955 int32
	_ = v955
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v971 int32
	_ = v971
	var v976 int32
	_ = v976
	var v980 int32
	_ = v980
	var v986 int32
	_ = v986
	var v991 int32
	_ = v991
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v1000 int32
	_ = v1000
	var v1005 int32
	_ = v1005
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1012 int32
	_ = v1012
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1025 int32
	_ = v1025
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1059 int32
	_ = v1059
	var v1061 int32
	_ = v1061
	var v1073 int32
	_ = v1073
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1088 int32
	_ = v1088
	var v1091 int32
	_ = v1091
	var v1093 int32
	_ = v1093
	var v1096 int32
	_ = v1096
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1102 int32
	_ = v1102
	var v1114 int32
	_ = v1114
	var v1126 int32
	_ = v1126
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1133 int32
	_ = v1133
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1142 int32
	_ = v1142
	var v1145 int32
	_ = v1145
	var v1147 int32
	_ = v1147
	var v1150 int32
	_ = v1150
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1160 int32
	_ = v1160
	var v1176 int32
	_ = v1176
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1188 int32
	_ = v1188
	var v1192 int32
	_ = v1192
	var v1197 int32
	_ = v1197
	var v1221 int32
	_ = v1221
	var v1225 int32
	_ = v1225
	var v1230 int32
	_ = v1230
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1271 int32
	_ = v1271
	var v1275 int32
	_ = v1275
	var v1278 int32
	_ = v1278
	var v1280 int32
	_ = v1280
	var v1283 int32
	_ = v1283
	var v1290 int32
	_ = v1290
	var v1292 int32
	_ = v1292
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1314 int32
	_ = v1314
	var v1318 int32
	_ = v1318
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1342 int32
	_ = v1342
	var v1345 int32
	_ = v1345
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1360 int32
	_ = v1360
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1366 int32
	_ = v1366
	var v1370 int32
	_ = v1370
	var v1373 int32
	_ = v1373
	var v1375 int32
	_ = v1375
	var v1378 int32
	_ = v1378
	var v1385 int32
	_ = v1385
	var v1387 int32
	_ = v1387
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1409 int32
	_ = v1409
	var v1413 int32
	_ = v1413
	var v1415 int32
	_ = v1415
	var v1432 int32
	_ = v1432
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1444 int32
	_ = v1444
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1468 int32
	_ = v1468
	v21 = m.G0
	v23 = v21 - int32(224)
	m.G0 = v23
	if l1 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v23 + int32(224)
	return v1468
L2:
	;
	v37 = v25 << (uint(int32(2)) % 32)
	v38 = F_palloc0(m, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L7
	} else {
		goto L9
	}
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v25 != 0 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v27 = int32(0)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v32 = F_bms_add_range(m, v27, v27, v29-int32(1))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
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
	v1468 = v32
	goto L1
L9:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v40 <= int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(v38+v37-int32(4))))
	v1255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1254)+4)))
	v1256 = int32(0)
	v1257 = *(*int32)(unsafe.Add(mBase, uint32(v1254)))
	if v1257 == v1256 {
		goto L284
	} else {
		goto L285
	}
L11:
	;
	v60 = int32(0)
	goto L13
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1221 = m.ExcPending
	if v1221 != 0 {
		goto L7
	} else {
		goto L279
	}
L13:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63+v60<<(uint(int32(2))%32))))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	switch v68 - int32(377) {
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
	v1188 = m.ExcPending
	if v1188 != 0 {
		goto L7
	} else {
		goto L276
	}
L15:
	;
	goto L14
L16:
	;
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v38+v1176<<(uint(int32(2))%32)))) = v1160
	v1182 = v60 + int32(1)
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v1182 < v1183 {
		v60 = v1182
		goto L13
	} else {
		goto L275
	}
L17:
	;
	v1007 = F_palloc0(m, int32(8))
	mBase = m.M
	v1008 = m.ExcPending
	if v1008 != 0 {
		goto L7
	} else {
		goto L241
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v995 = m.ExcPending
	if v995 != 0 {
		goto L7
	} else {
		goto L238
	}
L19:
	;
	v71 = int32(0)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	if v73 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
	v75 = v74
	goto L22
L21:
	;
	v75 = v71
	goto L22
L22:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
	if v76 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
	v78 = v77
	goto L25
L24:
	;
	v78 = v71
	goto L25
L25:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v79 <= int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	v257 = v252 + v253*v238*int32(28)
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	switch v258 - int32(104) {
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
	v238 = v79
	v240 = int32(0)
	goto L26
L28:
	;
	goto L29
L29:
	;
	v83 = int32(0)
	v87 = v75
	v88 = v78
	v89 = v83
	v93 = v83
	goto L30
L30:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	v106 = F_bms_is_member(m, v89, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L7
	} else {
		goto L33
	}
L31:
	;
	v238 = v230
	v240 = v224
	goto L26
L32:
	;
	v229 = v89 + int32(1)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v229 < v230 {
		v87 = v220
		v88 = v221
		v89 = v229
		v93 = v224
		goto L30
	} else {
		goto L63
	}
L33:
	;
	if v106 != 0 {
		v220 = v87
		v221 = v88
		v224 = v93
		goto L32
	} else {
		goto L34
	}
L34:
	;
	if v89 <= v93 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	if v87 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v109 != int32(114) {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v238 = v112
	v240 = v93
	goto L26
L38:
	;
	v220 = int32(0)
	v221 = v88
	v224 = v93
	goto L32
L39:
	;
	goto L40
L40:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v119 = v116*v117 + v89
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	if v121 == int32(7) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v165 = v162 + v119*int32(28)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	if v161 == v166 {
		goto L50
	} else {
		goto L51
	}
L42:
	;
	v152 = F_palloc(m, int32(8))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L7
	} else {
		goto L49
	}
L43:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v120)+20))
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+192)) = uint8(v125)
	if v125 != 0 {
		goto L42
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v127+v119<<(uint(int32(2))%32))))
	v132 = int32(_a_F_get_matching_partitions_0)
	v133 = *(*int32)(unsafe.Add(mBase, _c_F_get_matching_partitions[0]))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_get_matching_partitions[0])) = v136
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v131)+20))
	v141 = m.T0[v140].(func(*base.Module, int32, int32, int32) int32)(m, v131, v135, v23+int32(192))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L7
	} else {
		goto L47
	}
L46:
	;
	v158 = v124
	goto L41
L47:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_get_matching_partitions[0])) = v133
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+192)))
	if v145 != int32(1) {
		v158 = v141
		goto L41
	} else {
		goto L48
	}
L48:
	;
	goto L42
L49:
	;
	v154 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v152)+4)) = uint16(v154)
	*(*int32)(unsafe.Add(mBase, uint32(v152))) = v154
	v1160 = v152
	goto L16
L50:
	;
	v192 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v23-int32(-64)+v89<<(uint(v192)%32)))) = v158
	v197 = v88 + int32(4)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v199)+12))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v199)+4))
	if base.Ui32(v197) < base.Ui32(v200+v201<<(uint(v192)%32)) {
		goto L57
	} else {
		goto L58
	}
L51:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v172 = v169 + v89*int32(28)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	if v173 == v161 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v175 = *(*int64)(unsafe.Add(mBase, uint32(v172)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v165)+16)) = v175
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v172)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v165)+24)) = v177
	v179 = *(*int64)(unsafe.Add(mBase, uint32(v172)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v165)+8)) = v179
	v181 = *(*int64)(unsafe.Add(mBase, uint32(v172)))
	*(*int64)(unsafe.Add(mBase, uint32(v165))) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v165)+20)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v165)+16)) = int32(0)
	goto L55
L53:
	;
	goto L54
L54:
	;
	F_fmgr_info_cxt(m, v161, v165, v168)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
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
	v206 = v197
	goto L59
L58:
	;
	v206 = int32(0)
	goto L59
L59:
	;
	v208 = v87 + int32(4)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+12))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v210)+4))
	if base.Ui32(v208) < base.Ui32(v211+v212<<(uint(int32(2))%32)) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v217 = v208
	goto L62
L61:
	;
	v217 = int32(0)
	goto L62
L62:
	;
	v220 = v217
	v221 = v206
	v224 = v93 + int32(1)
	goto L32
L63:
	;
	goto L31
L64:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L7
	} else {
		goto L235
	}
L65:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
		goto L7
	} else {
		goto L232
	}
L66:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	v560 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+8)))
	v562 = F_palloc0(m, int32(8))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L7
	} else {
		goto L143
	}
L67:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	v428 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+8)))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v431 = F_palloc0(m, int32(8))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L7
	} else {
		goto L99
	}
L68:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	v263 = F_palloc0(m, int32(8))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L7
	} else {
		goto L69
	}
L69:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)+24))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v269 = int32(0)
	if v261 == v269 {
		goto L73
	} else {
		goto L74
	}
L70:
	;
	v425 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v263)+4)) = uint16(v425)
	v1160 = v263
	goto L16
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v263))) = v403
	goto L70
L72:
	;
	if v268 == v304+v240 {
		goto L85
	} else {
		goto L86
	}
L73:
	;
	v304 = int32(0)
	goto L72
L74:
	;
	goto L75
L75:
	;
	v276 = int32(1)
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v261)+4))
	if v277 <= v276 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v280 = v276
	goto L78
L77:
	;
	v280 = v277
	goto L78
L78:
	;
	v284 = int32(0)
	v286 = v269
	goto L79
L79:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v261+int32(8)+v284<<(uint(int32(2))%32))))
	if v292 != 0 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v304 = v295
	goto L72
L81:
	;
	v295 = v286 + base.I32_popcnt(v292)
	goto L83
L82:
	;
	v295 = v286
	goto L83
L83:
	;
	v297 = v284 + int32(1)
	if v297 != v280 {
		v284 = v297
		v286 = v295
		goto L79
	} else {
		goto L84
	}
L84:
	;
	goto L80
L85:
	;
	v307 = int32(0)
	if v307 < v268 {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	goto L87
L87:
	;
	v376 = int32(0)
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v265)+20))
	v381 = F_bms_add_range(m, v376, v376, v378-int32(1))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L7
	} else {
		goto L98
	}
L88:
	;
	v312 = v307
	goto L91
L89:
	;
	goto L90
L90:
	;
	v363 = F_compute_partition_hash_value(m, v268, v257, v267, v23-int32(-64), v23+int32(192))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L7
	} else {
		goto L95
	}
L91:
	;
	v333 = F_bms_is_member(m, v312, v261)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L7
	} else {
		goto L93
	}
L92:
	;
	goto L90
L93:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v23+int32(192)+v312))) = uint8(v333)
	v337 = v312 + int32(1)
	if v337 != v268 {
		v312 = v337
		goto L91
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	v365 = int64(*(*int32)(unsafe.Add(mBase, uint32(v265)+20)))
	v366 = base.I64_rem_u_s(v363, v365)
	v367 = base.I32_wrap_i64(v366)
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v266+v367<<(uint(int32(2))%32))))
	if v371 < int32(0) {
		goto L70
	} else {
		goto L96
	}
L96:
	;
	v374 = F_bms_make_singleton(m, v367)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L7
	} else {
		goto L97
	}
L97:
	;
	v403 = v374
	goto L71
L98:
	;
	v403 = v381
	goto L71
L99:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v435 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v431)+4)) = uint16(v435)
	if v427 != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v433)+28))
	if v437 != int32(-1) {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	goto L102
L102:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v433)+4))
	if v446 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L103:
	;
	v440 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v431)+5)) = uint8(v440)
	v1160 = v431
	goto L16
L104:
	;
	goto L105
L105:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v433)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v431)+4)) = uint8(base.B2i32(v442 != int32(-1)))
	v1160 = v431
	goto L16
L106:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v433)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v431)+4)) = uint8(base.B2i32(v449 != int32(-1)))
	v1160 = v431
	goto L16
L107:
	;
	goto L108
L108:
	;
	v454 = v446 - int32(1)
	if v240 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v457 = int32(0)
	v459 = F_bms_add_range(m, v457, v457, v454)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L7
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	switch v428 {
	case 0:
		goto L120
	default:
		goto L119
	case 3:
		goto L118
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v431))) = v459
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v433)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v431)+4)) = uint8(base.B2i32(v462 != int32(-1)))
	v1160 = v431
	goto L16
L113:
	;
	v556 = F_bms_add_range(m, int32(0), v552, v553)
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L7
	} else {
		goto L142
	}
L114:
	;
	v540 = F_partition_list_bsearch(m, v257, v434, v433, v429, v23+int32(192))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L7
	} else {
		goto L137
	}
L115:
	;
	v536 = int32(0)
	goto L114
L116:
	;
	v520 = F_partition_list_bsearch(m, v257, v434, v433, v429, v23+int32(192))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L7
	} else {
		goto L132
	}
L117:
	;
	v517 = int32(1)
	goto L116
L118:
	;
	v500 = F_partition_list_bsearch(m, v257, v434, v433, v429, v23+int32(192))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L7
	} else {
		goto L128
	}
L119:
	;
	v490 = int32(-1)
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v433)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v431)+4)) = uint8(base.B2i32(v491 != v490))
	switch v428 - int32(1) {
	case 0:
		v536 = v490
		goto L114
	case 1:
		goto L115
	default:
		goto L64
	case 3:
		goto L117
	case 4:
		v517 = int32(0)
		goto L116
	}
L120:
	;
	v466 = int32(0)
	v468 = F_bms_add_range(m, v466, v466, v454)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L7
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v431))) = v468
	v473 = F_partition_list_bsearch(m, v257, v434, v433, v429, v23+int32(192))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L7
	} else {
		goto L123
	}
L122:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v433)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v431)+4)) = uint8(base.B2i32(v486 != int32(-1)))
	v1160 = v431
	goto L16
L123:
	;
	if v473 < int32(0) {
		goto L122
	} else {
		goto L124
	}
L124:
	;
	v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+192)))
	if v477&int32(1) == int32(0) {
		goto L122
	} else {
		goto L125
	}
L125:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v431)))
	v483 = F_bms_del_member(m, v482, v473)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L7
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v431))) = v483
	goto L122
L127:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v433)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v431)+4)) = uint8(base.B2i32(v512 != int32(-1)))
	v1160 = v431
	goto L16
L128:
	;
	if v500 < int32(0) {
		goto L127
	} else {
		goto L129
	}
L129:
	;
	v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+192)))
	if v504&int32(1) == int32(0) {
		goto L127
	} else {
		goto L130
	}
L130:
	;
	v509 = F_bms_make_singleton(m, v500)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L7
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v431))) = v509
	v1160 = v431
	goto L16
L132:
	;
	v522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+192)))
	v527 = int32(0)
	if v527 <= v520 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v530 = v520 + (v517&v522 ^ int32(1))
	goto L135
L134:
	;
	v530 = v527
	goto L135
L135:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v433)+4))
	if v530 <= v531-int32(1) {
		v552 = v530
		v553 = v454
		goto L113
	} else {
		goto L136
	}
L136:
	;
	v1160 = v431
	goto L16
L137:
	;
	v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+192)))
	v543 = int32(0)
	if v543 <= v540 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v546 = v542
	goto L140
L139:
	;
	v546 = v543
	goto L140
L140:
	;
	v548 = v540 - v546&v536
	if v548 < int32(0) {
		v1160 = v431
		goto L16
	} else {
		goto L141
	}
L141:
	;
	v552 = int32(0)
	v553 = v548
	goto L113
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v431))) = v556
	v1160 = v431
	goto L16
L143:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v565 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v566 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v566)+24))
	v568 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v562)+4)) = uint16(v568)
	if v559 == v568 {
		goto L145
	} else {
		goto L146
	}
L144:
	;
	if v240 == int32(0) {
		goto L149
	} else {
		goto L150
	}
L145:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v566)+4))
	if v572 != 0 {
		goto L144
	} else {
		goto L148
	}
L146:
	;
	goto L147
L147:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v566)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v562)+4)) = uint8(base.B2i32(v574 != int32(-1)))
	v1160 = v562
	goto L16
L148:
	;
	goto L147
L149:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v567+v572<<(uint(int32(2))%32))))
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v567)))
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v566)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v562)+4)) = uint8(base.B2i32(v585 != int32(-1)))
	v590 = int32(31)
	v595 = F_bms_add_range(m, int32(0), int32(base.Ui32(v584)>>(uint(v590)%32)), v583>>(uint(v590)%32)+v572)
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L7
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	v598 = base.B2i32(v564 <= v240)
	if v598 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v562))) = v595
	v1160 = v562
	goto L16
L153:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v566)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v562)+4)) = uint8(base.B2i32(v601 != int32(-1)))
	goto L155
L154:
	;
	goto L155
L155:
	;
	v605 = int32(1)
	v606 = int32(0)
	switch v560 - v605 {
	case 0:
		v801 = v605
		goto L159
	case 1:
		goto L160
	case 2:
		goto L163
	case 3:
		goto L162
	case 4:
		v737 = v606
		goto L161
	default:
		goto L158
	}
L156:
	;
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v566)+4))
	if v910 <= v895 {
		v931 = v895
		goto L224
	} else {
		goto L225
	}
L157:
	;
	v893 = v572
	v895 = v742 + int32(1)
	goto L156
L158:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L7
	} else {
		goto L221
	}
L159:
	;
	v806 = F_partition_range_datum_bsearch(m, v257, v565, v566, v240, v23-int32(-64), v23+int32(192))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L7
	} else {
		goto L204
	}
L160:
	;
	v801 = int32(0)
	goto L159
L161:
	;
	v742 = F_partition_range_datum_bsearch(m, v257, v565, v566, v240, v23-int32(-64), v23+int32(192))
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L7
	} else {
		goto L190
	}
L162:
	;
	v737 = int32(1)
	goto L161
L163:
	;
	v614 = F_partition_range_datum_bsearch(m, v257, v565, v566, v240, v23-int32(-64), v23+int32(192))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L7
	} else {
		goto L165
	}
L164:
	;
	v733 = F_bms_make_singleton(m, v614+int32(1))
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L7
	} else {
		goto L189
	}
L165:
	;
	if v614 < int32(0) {
		goto L164
	} else {
		goto L166
	}
L166:
	;
	v618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+192)))
	if v618&int32(1) == int32(0) {
		goto L164
	} else {
		goto L167
	}
L167:
	;
	if v240 == v564 {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v626 = F_bms_make_singleton(m, v614+int32(1))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L7
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	v631 = v614
	goto L172
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v562))) = v626
	v1160 = v562
	goto L16
L172:
	;
	if v631 <= int32(0) {
		goto L175
	} else {
		goto L176
	}
L173:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v566)+12))
	v672 = int32(2)
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v671+v669<<(uint(v672)%32))))
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v675+v240<<(uint(v672)%32))))
	v686 = v614
	goto L180
L174:
	;
	goto L173
L175:
	;
	v669 = int32(0)
	goto L174
L176:
	;
	goto L177
L177:
	;
	v653 = v631 - int32(1)
	v655 = v653 << (uint(int32(2)) % 32)
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v566)+8))
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v655+v656)))
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v566)+12))
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v659+v655)))
	v664 = F_partition_rbound_datum_cmp(m, v257, v565, v658, v661, v23-int32(-64), v240)
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L7
	} else {
		goto L178
	}
L178:
	;
	if v664 == int32(0) {
		v631 = v653
		goto L172
	} else {
		goto L179
	}
L179:
	;
	v669 = v631
	goto L174
L180:
	;
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v566)+4))
	if v703-int32(1) <= v686 {
		goto L183
	} else {
		goto L184
	}
L181:
	;
	v728 = F_bms_add_range(m, int32(0), v669+base.B2i32(v679 == int32(-1)), v725)
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L7
	} else {
		goto L188
	}
L182:
	;
	goto L181
L183:
	;
	v725 = v686 + int32(1)
	goto L182
L184:
	;
	goto L185
L185:
	;
	v710 = v686 + int32(1)
	v712 = v710 << (uint(int32(2)) % 32)
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v566)+8))
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v712+v713)))
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v566)+12))
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v716+v712)))
	v721 = F_partition_rbound_datum_cmp(m, v257, v565, v715, v718, v23-int32(-64), v240)
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L7
	} else {
		goto L186
	}
L186:
	;
	if v721 == int32(0) {
		v686 = v710
		goto L180
	} else {
		goto L187
	}
L187:
	;
	v725 = v710
	goto L182
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v562))) = v728
	v1160 = v562
	goto L16
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v562))) = v733
	v1160 = v562
	goto L16
L190:
	;
	if v742 < int32(0) {
		v893 = v572
		v895 = v606
		goto L156
	} else {
		goto L191
	}
L191:
	;
	if v564 <= v240 {
		goto L157
	} else {
		goto L192
	}
L192:
	;
	v746 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+192)))
	if v746&int32(1) == int32(0) {
		goto L157
	} else {
		goto L193
	}
L193:
	;
	if v737 != 0 {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v753 = int32(-1)
	goto L196
L195:
	;
	v753 = int32(1)
	goto L196
L196:
	;
	v756 = v742
	goto L197
L197:
	;
	if v756 <= int32(0) {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	v893 = v572
	v895 = v756 + base.B2i32(v737 == int32(0))
	goto L156
L199:
	;
	goto L198
L200:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v566)+4))
	if v776-int32(1) <= v756 {
		goto L199
	} else {
		goto L201
	}
L201:
	;
	v780 = v756 + v753
	v782 = v780 << (uint(int32(2)) % 32)
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v566)+8))
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v782+v783)))
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v566)+12))
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v786+v782)))
	v791 = F_partition_rbound_datum_cmp(m, v257, v565, v785, v788, v23-int32(-64), v240)
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L7
	} else {
		goto L202
	}
L202:
	;
	if v791 == int32(0) {
		v756 = v780
		goto L197
	} else {
		goto L203
	}
L203:
	;
	goto L199
L204:
	;
	if int32(0) <= v806 {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+192)))
	v813 = int32(0)
	if v598|base.B2i32(v810&int32(1) == v813) == v813 {
		goto L208
	} else {
		goto L209
	}
L206:
	;
	goto L207
L207:
	;
	v893 = v806 + int32(1)
	v895 = v606
	goto L156
L208:
	;
	if v801 != 0 {
		goto L211
	} else {
		goto L212
	}
L209:
	;
	goto L210
L210:
	;
	v893 = v806 + base.B2i32(v810&v801 == int32(0))
	v895 = v606
	goto L156
L211:
	;
	v820 = int32(-1)
	goto L213
L212:
	;
	v820 = int32(1)
	goto L213
L213:
	;
	v823 = v806
	goto L214
L214:
	;
	if v823 <= int32(0) {
		goto L216
	} else {
		goto L217
	}
L215:
	;
	v893 = v823 + base.B2i32(v801 == int32(0))
	v895 = v606
	goto L156
L216:
	;
	goto L215
L217:
	;
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v566)+4))
	if v843-int32(1) <= v823 {
		goto L216
	} else {
		goto L218
	}
L218:
	;
	v847 = v823 + v820
	v849 = v847 << (uint(int32(2)) % 32)
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v566)+8))
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v849+v850)))
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v566)+12))
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v853+v849)))
	v858 = F_partition_rbound_datum_cmp(m, v257, v565, v852, v855, v23-int32(-64), v240)
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L7
	} else {
		goto L219
	}
L219:
	;
	if v858 == int32(0) {
		v823 = v847
		goto L214
	} else {
		goto L220
	}
L220:
	;
	goto L216
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v560
	F_errmsg_internal(m, int32(_a_F_get_matching_partitions_1), v23+int32(48))
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L7
	} else {
		goto L222
	}
L222:
	;
	F_errfinish(m, int32(_a_F_get_matching_partitions_2), int32(3318), int32(_a_F_get_matching_partitions_3))
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
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
	if v893 <= int32(0) {
		v955 = v893
		goto L227
	} else {
		goto L228
	}
L225:
	;
	v913 = v895 << (uint(int32(2)) % 32)
	v915 = *(*int32)(unsafe.Add(mBase, uint32(v567+v913)))
	if int32(0) <= v915 {
		v931 = v895
		goto L224
	} else {
		goto L226
	}
L226:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v566)+12))
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v918+v913)))
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v920+v240<<(uint(int32(2))%32)-int32(4))))
	v931 = v895 + base.B2i32(v926 == int32(-1))
	goto L224
L227:
	;
	if v955 < v931 {
		v1160 = v562
		goto L16
	} else {
		goto L230
	}
L228:
	;
	v935 = v893 << (uint(int32(2)) % 32)
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v567+v935)))
	if int32(0) <= v937 {
		v955 = v893
		goto L227
	} else {
		goto L229
	}
L229:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v566)+12))
	v942 = int32(4)
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v940+v935-v942)))
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v944+v240<<(uint(int32(2))%32)-v942)))
	v955 = v893 - base.B2i32(v950 == int32(1))
	goto L227
L230:
	;
	v958 = F_bms_add_range(m, int32(0), v931, v955)
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L7
	} else {
		goto L231
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v562))) = v958
	v1160 = v562
	goto L16
L232:
	;
	v965 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v965
	F_errmsg_internal(m, int32(_a_F_get_matching_partitions_4), v23+int32(16))
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		goto L7
	} else {
		goto L233
	}
L233:
	;
	F_errfinish(m, int32(_a_F_get_matching_partitions_2), int32(3575), int32(_a_F_get_matching_partitions_5))
	mBase = m.M
	v976 = m.ExcPending
	if v976 != 0 {
		goto L7
	} else {
		goto L234
	}
L234:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v428
	F_errmsg_internal(m, int32(_a_F_get_matching_partitions_1), v23+int32(32))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L7
	} else {
		goto L236
	}
L236:
	;
	F_errfinish(m, int32(_a_F_get_matching_partitions_2), int32(2941), int32(_a_F_get_matching_partitions_6))
	mBase = m.M
	v991 = m.ExcPending
	if v991 != 0 {
		goto L7
	} else {
		goto L237
	}
L237:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L238:
	;
	v996 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v996
	F_errmsg_internal(m, int32(_a_F_get_matching_partitions_7), v23)
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L7
	} else {
		goto L239
	}
L239:
	;
	F_errfinish(m, int32(_a_F_get_matching_partitions_2), int32(893), int32(_a_F_get_matching_partitions_8))
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L7
	} else {
		goto L240
	}
L240:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L241:
	;
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	if v1009 == int32(0) {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	v1012 = int32(0)
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(v1014)+20))
	v1018 = F_bms_add_range(m, v1012, v1012, v1015-int32(1))
	mBase = m.M
	v1019 = m.ExcPending
	if v1019 != 0 {
		goto L7
	} else {
		goto L245
	}
L243:
	;
	goto L244
L244:
	;
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
	switch v1029 {
	case 0:
		goto L246
	case 1:
		goto L247
	default:
		v1160 = v1007
		goto L16
	}
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1007))) = v1018
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(v1014)+32))
	v1022 = int32(-1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1007)+4)) = uint8(base.B2i32(v1021 != v1022))
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v1014)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v1007)+5)) = uint8(base.B2i32(v1025 != v1022))
	v1160 = v1007
	goto L16
L246:
	;
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(v1009)+4))
	if v1102 <= int32(0) {
		v1160 = v1007
		goto L16
	} else {
		goto L263
	}
L247:
	;
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v1009)+4))
	if v1030 <= int32(0) {
		v1160 = v1007
		goto L16
	} else {
		goto L248
	}
L248:
	;
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(v1009)+12))
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v1033)))
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	if v1035 <= v1034 {
		goto L12
	} else {
		goto L249
	}
L249:
	;
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v38+v1034<<(uint(int32(2))%32))))
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v1040)))
	v1042 = F_bms_copy(m, v1041)
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L7
	} else {
		goto L250
	}
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1007))) = v1042
	v1045 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1040)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1007)+5)) = uint8(v1045)
	v1047 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1040)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1007)+4)) = uint8(v1047)
	v1049 = int32(1)
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v1009)+4))
	if v1050 <= v1049 {
		v1160 = v1007
		goto L16
	} else {
		goto L251
	}
L251:
	;
	v1059 = v1042
	v1061 = v1049
	goto L252
L252:
	;
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(v1009)+12))
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v1073+v1061<<(uint(int32(2))%32))))
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	if v1078 <= v1077 {
		goto L12
	} else {
		goto L254
	}
L253:
	;
	v1160 = v1007
	goto L16
L254:
	;
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(v38+v1077<<(uint(int32(2))%32))))
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(v1083)))
	v1085 = F_bms_int_members(m, v1059, v1084)
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		goto L7
	} else {
		goto L255
	}
L255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1007))) = v1085
	v1088 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1007)+5)))
	if v1088 == int32(1) {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	v1091 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1083)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1007)+5)) = uint8(v1091)
	goto L258
L257:
	;
	goto L258
L258:
	;
	v1093 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1007)+4)))
	if v1093 == int32(1) {
		goto L259
	} else {
		goto L260
	}
L259:
	;
	v1096 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1083)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1007)+4)) = uint8(v1096)
	goto L261
L260:
	;
	goto L261
L261:
	;
	v1099 = v1061 + int32(1)
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(v1009)+4))
	if v1099 < v1100 {
		v1059 = v1085
		v1061 = v1099
		goto L252
	} else {
		goto L262
	}
L262:
	;
	goto L253
L263:
	;
	v1114 = int32(0)
	goto L264
L264:
	;
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(v1009)+12))
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(v1126+v1114<<(uint(int32(2))%32))))
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	if v1131 <= v1130 {
		goto L15
	} else {
		goto L266
	}
L265:
	;
	v1160 = v1007
	goto L16
L266:
	;
	v1133 = *(*int32)(unsafe.Add(mBase, uint32(v1007)))
	v1137 = *(*int32)(unsafe.Add(mBase, uint32(v38+v1130<<(uint(int32(2))%32))))
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(v1137)))
	v1139 = F_bms_add_members(m, v1133, v1138)
	mBase = m.M
	v1140 = m.ExcPending
	if v1140 != 0 {
		goto L7
	} else {
		goto L267
	}
L267:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1007))) = v1139
	v1142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1007)+5)))
	if v1142 == int32(0) {
		goto L268
	} else {
		goto L269
	}
L268:
	;
	v1145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1137)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1007)+5)) = uint8(v1145)
	goto L270
L269:
	;
	goto L270
L270:
	;
	v1147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1007)+4)))
	if v1147 == int32(0) {
		goto L271
	} else {
		goto L272
	}
L271:
	;
	v1150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1137)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1007)+4)) = uint8(v1150)
	goto L273
L272:
	;
	goto L273
L273:
	;
	v1153 = v1114 + int32(1)
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v1009)+4))
	if v1153 < v1154 {
		v1114 = v1153
		goto L264
	} else {
		goto L274
	}
L274:
	;
	goto L265
L275:
	;
	goto L10
L276:
	;
	F_errmsg_internal(m, int32(_a_F_get_matching_partitions_9), int32(0))
	mBase = m.M
	v1192 = m.ExcPending
	if v1192 != 0 {
		goto L7
	} else {
		goto L277
	}
L277:
	;
	F_errfinish(m, int32(_a_F_get_matching_partitions_2), int32(3630), int32(_a_F_get_matching_partitions_10))
	mBase = m.M
	v1197 = m.ExcPending
	if v1197 != 0 {
		goto L7
	} else {
		goto L278
	}
L278:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L279:
	;
	F_errmsg_internal(m, int32(_a_F_get_matching_partitions_9), int32(0))
	mBase = m.M
	v1225 = m.ExcPending
	if v1225 != 0 {
		goto L7
	} else {
		goto L280
	}
L280:
	;
	F_errfinish(m, int32(_a_F_get_matching_partitions_2), int32(3654), int32(_a_F_get_matching_partitions_10))
	mBase = m.M
	v1230 = m.ExcPending
	if v1230 != 0 {
		goto L7
	} else {
		goto L281
	}
L281:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L282:
	;
	if int32(0) <= v1314 {
		goto L293
	} else {
		goto L294
	}
L283:
	;
	v1314 = base.I32_ctz(v1300) | v1301<<(uint(int32(5))%32)
	goto L282
L284:
	;
	v1314 = int32(-2)
	goto L282
L285:
	;
	v1267 = base.I32_div_s(int32(0), int32(32))
	v1268 = *(*int32)(unsafe.Add(mBase, uint32(v1257)+4))
	if v1268 <= v1267 {
		goto L284
	} else {
		goto L286
	}
L286:
	;
	v1271 = v1257 + int32(8)
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(v1271+v1267<<(uint(int32(2))%32))))
	v1278 = v1275 & int32(-1)
	if v1278 != 0 {
		v1300 = v1278
		v1301 = v1267
		goto L283
	} else {
		goto L287
	}
L287:
	;
	v1280 = v1267 + int32(1)
	if v1280 == v1268 {
		goto L284
	} else {
		goto L288
	}
L288:
	;
	v1283 = v1280
	goto L289
L289:
	;
	v1290 = *(*int32)(unsafe.Add(mBase, uint32(v1271+v1283<<(uint(int32(2))%32))))
	if v1290 != 0 {
		v1300 = v1290
		v1301 = v1283
		goto L283
	} else {
		goto L291
	}
L290:
	;
	goto L284
L291:
	;
	v1292 = v1283 + int32(1)
	if v1292 != v1268 {
		v1283 = v1292
		goto L289
	} else {
		goto L292
	}
L292:
	;
	goto L290
L293:
	;
	v1318 = v1256
	v1320 = v1255
	v1321 = v1314
	goto L296
L294:
	;
	v1413 = v1256
	v1415 = v1255
	goto L295
L295:
	;
	v1432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1254)+5)))
	if v1432 == int32(1) {
		goto L315
	} else {
		goto L316
	}
L296:
	;
	v1337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1338 = *(*int32)(unsafe.Add(mBase, uint32(v1337)+24))
	v1342 = *(*int32)(unsafe.Add(mBase, uint32(v1338+v1321<<(uint(int32(2))%32))))
	if v1342 < int32(0) {
		goto L299
	} else {
		goto L300
	}
L297:
	;
	v1413 = v1351
	v1415 = v1352
	goto L295
L298:
	;
	v1353 = *(*int32)(unsafe.Add(mBase, uint32(v1254)))
	if v1353 == int32(0) {
		goto L305
	} else {
		goto L306
	}
L299:
	;
	v1345 = *(*int32)(unsafe.Add(mBase, uint32(v1337)+32))
	v1351 = v1318
	v1352 = v1320 | base.B2i32(v1345 != int32(-1))
	goto L298
L300:
	;
	goto L301
L301:
	;
	v1349 = F_bms_add_member(m, v1318, v1342)
	mBase = m.M
	v1350 = m.ExcPending
	if v1350 != 0 {
		goto L7
	} else {
		goto L302
	}
L302:
	;
	v1351 = v1349
	v1352 = v1320
	goto L298
L303:
	;
	if int32(0) <= v1409 {
		v1318 = v1351
		v1320 = v1352
		v1321 = v1409
		goto L296
	} else {
		goto L314
	}
L304:
	;
	v1409 = base.I32_ctz(v1395) | v1396<<(uint(int32(5))%32)
	goto L303
L305:
	;
	v1409 = int32(-2)
	goto L303
L306:
	;
	v1360 = v1321 + int32(1)
	v1362 = base.I32_div_s(v1360, int32(32))
	v1363 = *(*int32)(unsafe.Add(mBase, uint32(v1353)+4))
	if v1363 <= v1362 {
		goto L305
	} else {
		goto L307
	}
L307:
	;
	v1366 = v1353 + int32(8)
	v1370 = *(*int32)(unsafe.Add(mBase, uint32(v1366+v1362<<(uint(int32(2))%32))))
	v1373 = v1370 & (int32(-1) << (uint(v1360) % 32))
	if v1373 != 0 {
		v1395 = v1373
		v1396 = v1362
		goto L304
	} else {
		goto L308
	}
L308:
	;
	v1375 = v1362 + int32(1)
	if v1375 == v1363 {
		goto L305
	} else {
		goto L309
	}
L309:
	;
	v1378 = v1375
	goto L310
L310:
	;
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(v1366+v1378<<(uint(int32(2))%32))))
	if v1385 != 0 {
		v1395 = v1385
		v1396 = v1378
		goto L304
	} else {
		goto L312
	}
L311:
	;
	goto L305
L312:
	;
	v1387 = v1378 + int32(1)
	if v1387 != v1363 {
		v1378 = v1387
		goto L310
	} else {
		goto L313
	}
L313:
	;
	goto L311
L314:
	;
	goto L297
L315:
	;
	v1435 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(v1435)+28))
	v1437 = F_bms_add_member(m, v1413, v1436)
	mBase = m.M
	v1438 = m.ExcPending
	if v1438 != 0 {
		goto L7
	} else {
		goto L318
	}
L316:
	;
	v1439 = v1413
	goto L317
L317:
	;
	if v1415&int32(1) == int32(0) {
		v1468 = v1439
		goto L1
	} else {
		goto L319
	}
L318:
	;
	v1439 = v1437
	goto L317
L319:
	;
	v1444 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1445 = *(*int32)(unsafe.Add(mBase, uint32(v1444)+32))
	v1446 = F_bms_add_member(m, v1439, v1445)
	mBase = m.M
	v1447 = m.ExcPending
	if v1447 != 0 {
		goto L7
	} else {
		goto L320
	}
L320:
	;
	v1468 = v1446
	goto L1
}
func F_get_oprjoin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v4 = F_SearchSysCache1(m, int32(40), l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v12+v13)+108))
			F_ReleaseCatCache(m, v4)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				return v15
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
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int64
	_ = v157
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	v11 = m.G0
	v13 = v11 - int32(176)
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
	v213 = m.ExcPending
	if v213 != 0 {
		goto L3
	} else {
		goto L56
	}
L2:
	;
	m.G0 = v13 + int32(176)
	return v205
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
	v205 = v16
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
		v205 = int32(0)
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
	F_aclcheck_error(m, v65, int32(36), v68)
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
	*(*int32)(unsafe.Add(mBase, uint32(v13)+112)) = v144
	v148 = v13 + int32(32)
	v150 = F_strncpy(m, v148, v72, int32(64))
	mBase = m.M
	v151 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v150)+63)) = uint8(v151)
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+120)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v13)+116)) = v148
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_get_other_operator[0]))
	v157 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+148)) = v157
	*(*int32)(unsafe.Add(mBase, uint32(v13)+144)) = l2
	*(*int64)(unsafe.Add(mBase, uint32(v13)+132)) = v157
	*(*int32)(unsafe.Add(mBase, uint32(v13)+124)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v13)+156)) = v157
	*(*int64)(unsafe.Add(mBase, uint32(v13)+164)) = v157
	*(*int32)(unsafe.Add(mBase, uint32(v13)+140)) = l1
	if l1 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v170 = int32(98)
	goto L45
L44:
	;
	v170 = int32(108)
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+128)) = v170
	v176 = F_heap_form_tuple(m, v137, v13+int32(112), v13+int32(96))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L3
	} else {
		goto L46
	}
L46:
	;
	F_CatalogTupleInsert(m, v135, v176)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L3
	} else {
		goto L47
	}
L47:
	;
	F_makeOperatorDependencies(m, v13+int32(20), v176, int32(1), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L3
	} else {
		goto L48
	}
L48:
	;
	F_pfree(m, v176)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L3
	} else {
		goto L49
	}
L49:
	;
	v189 = *(*int32)(unsafe.Add(mBase, _c_F_get_other_operator[1]))
	if v189 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v191 = int32(0)
	F_RunObjectPostCreateHook(m, int32(2617), v144, v191, v191)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
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
	v196 = m.ExcPending
	if v196 != 0 {
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
	v199 = m.ExcPending
	if v199 != 0 {
		goto L3
	} else {
		goto L55
	}
L55:
	;
	v205 = v144
	goto L2
L56:
	;
	F_errcode(m, int32(33579140))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L3
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v72
	F_errmsg(m, int32(_a_F_get_other_operator_4), v13)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L3
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_get_other_operator_5), int32(214), int32(_a_F_get_other_operator_6))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
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
		v15 = int32(41)
		return v15
	case 10:
		return int32(37)
	case 29:
		return int32(18)
	case 36:
		return int32(23)
	case 45:
		return int32(51)
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
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v4 = F_SearchSysCache1(m, int32(82), l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v12+v13)+84))
			F_ReleaseCatCache(m, v4)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	}
}
func F_get_typbyval(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v4 = F_SearchSysCache1(m, int32(82), l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 != 0 {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
			v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+22)))
			v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+v9)+78)))
			F_ReleaseCatCache(m, v4)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				v14 = v11
				return v14 & int32(1)
			}
		} else {
			v14 = int32(0)
			return v14 & int32(1)
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
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
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
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = F_SearchSysCache1(m, int32(82), l0)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		if v12 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
				F_errmsg_internal(m, int32(_a_F_get_typlenbyvalalign_0), v9)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_get_typlenbyvalalign_1), int32(2419), int32(_a_F_get_typlenbyvalalign_2))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+22)))
			v31 = v29 + v30
			v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31)+76)))
			*(*uint16)(unsafe.Add(mBase, uint32(l1))) = uint16(v32)
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+78)))
			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v34)
			v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+128)))
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v36)
			F_ReleaseCatCache(m, v12)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
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
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v4 = F_SearchSysCache1(m, int32(82), l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
			v15 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12+v13)+79)))
			F_ReleaseCatCache(m, v4)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				return base.I32_extend8_s(v15)
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
func F_ginarrayextract(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_copy(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
		F_get_typlenbyvalalign(m, v16, v7+int32(14), v7+int32(13), v7+int32(12))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v26 = int32(*(*int16)(unsafe.Add(mBase, uint32(v7)+14)))
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+13)))
			v28 = int32(*(*int8)(unsafe.Add(mBase, uint32(v7)+12)))
			F_deconstruct_array(m, v10, v26, v27, v28, v7+int32(8), v7+int32(4), v7)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				*(*int32)(unsafe.Add(mBase, uint32(v15))) = v35
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = v37
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
				m.G0 = v7 + int32(16)
				return v39
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
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
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
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v358 int64
	_ = v358
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
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
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v515 int32
	_ = v515
	var v516 float64
	_ = v516
	var v518 float64
	_ = v518
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v525 float64
	_ = v525
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v553 int32
	_ = v553
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v570 int64
	_ = v570
	var v573 int64
	_ = v573
	var v577 int64
	_ = v577
	var v579 float64
	_ = v579
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v622 int32
	_ = v622
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v647 int64
	_ = v647
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v661 int64
	_ = v661
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v675 int64
	_ = v675
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v689 int64
	_ = v689
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v782 float64
	_ = v782
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v855 int64
	_ = v855
	var v857 int32
	_ = v857
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v905 int32
	_ = v905
	var v908 int32
	_ = v908
	var v915 int32
	_ = v915
	var v918 float64
	_ = v918
	var v922 int32
	_ = v922
	var v926 int32
	_ = v926
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v937 int32
	_ = v937
	var v945 int32
	_ = v945
	var v951 int32
	_ = v951
	var v955 int32
	_ = v955
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v981 int64
	_ = v981
	var v982 int32
	_ = v982
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v997 int32
	_ = v997
	var v998 int64
	_ = v998
	var v1000 int32
	_ = v1000
	var v1011 int32
	_ = v1011
	var v1015 int32
	_ = v1015
	var v1020 int32
	_ = v1020
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1026 int32
	_ = v1026
	var v1034 int32
	_ = v1034
	var v1040 int32
	_ = v1040
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1055 int32
	_ = v1055
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1075 float64
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1080 int32
	_ = v1080
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1094 int32
	_ = v1094
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1131 int32
	_ = v1131
	var v1133 int32
	_ = v1133
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1141 int32
	_ = v1141
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1185 float64
	_ = v1185
	var v1188 int32
	_ = v1188
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1193 int32
	_ = v1193
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1206 int32
	_ = v1206
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1217 int32
	_ = v1217
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1222 float64
	_ = v1222
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1238 int32
	_ = v1238
	var v1243 int32
	_ = v1243
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
	v1231 = m.ExcPending
	if v1231 != 0 {
		goto L1
	} else {
		goto L235
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
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l2)+128))
	if v212 <= int32(0) {
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
	*(*int64)(unsafe.Add(mBase, uint32(v179+int32(80))+232)) = int64(2)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	*(*int32)(unsafe.Add(mBase, uint32(v179))) = v202 + v191
	v208 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10])) = v208 - v191
	goto L27
L30:
	;
	v458 = v20 + int32(_a_F_ginbuild_7)
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[1])))
	if v459 != 0 {
		goto L95
	} else {
		goto L96
	}
L31:
	;
	v216 = v212 + int32(1)
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+121)))
	v219 = F_palloc0(m, int32(28))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v223 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[19]))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v223)+72)) = v224 + int32(1)
	goto L33
L33:
	;
	v231 = F_CreateParallelContext(m, int32(_a_F_ginbuild_8), int32(_a_F_ginbuild_9), v212)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	if v217 == int32(1) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v235 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	v239 = int32(_a_F_ginbuild_10)
	goto L37
L37:
	;
	v241 = F_table_parallelscan_estimate(m, l0, v239)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	v237 = F_RegisterSnapshot(m, v235)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v239 = v237
	goto L37
L40:
	;
	v243 = F_add_size(m, int32(64), v241)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v231)+36))
	v250 = F_add_size(m, v245, (v243+int32(31))&int32(-32))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+36)) = v250
	v253 = F_tuplesort_estimate_shared(m, v216)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v231)+36))
	v260 = F_add_size(m, v255, (v253+int32(31))&int32(-32))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+36)) = v260
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v231)+40))
	v265 = F_add_size(m, v263, int32(2))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+40)) = v265
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v231)+36))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v231)+12))
	v271 = F_mul_size(m, int32(32), v270)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v277 = F_add_size(m, v268, (v271+int32(31))&int32(-32))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+36)) = v277
	v280 = int32(1)
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v231)+40))
	v283 = F_add_size(m, v281, v280)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+40)) = v283
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v231)+36))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v231)+12))
	v289 = F_mul_size(m, int32(128), v288)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v295 = F_add_size(m, v286, (v289+int32(31))&int32(-32))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+36)) = v295
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v231)+40))
	v300 = F_add_size(m, v298, int32(1))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+40)) = v300
	v304 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[20]))
	if v304 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v231)+36))
	v306 = F_strlen(m, v304)
	mBase = m.M
	v311 = F_add_size(m, v305, v306&int32(-32)+int32(32))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L55
	}
L53:
	;
	v321 = v280
	goto L54
L54:
	;
	F_InitializeParallelDSM(m, v231)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L57
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+36)) = v311
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v231)+40))
	v316 = F_add_size(m, v314, int32(1))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+40)) = v316
	v321 = v306 + int32(1)
	goto L54
L57:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v231)+44))
	if v324 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v239)))
	switch v327 {
	case 0, 5:
		goto L62
	default:
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	v340 = F_shm_toc_allocate(m, v339, v243)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L66
	}
L61:
	;
	F_DestroyParallelContext(m, v231)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L64
	}
L62:
	;
	F_UnregisterSnapshot(m, v239)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v334 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[19]))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v334)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v334)+72)) = v335 - int32(1)
	goto L65
L65:
	;
	goto L30
L66:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v340))) = v342
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v340)+12)) = v216
	*(*uint8)(unsafe.Add(mBase, uint32(v340)+8)) = uint8(v217)
	*(*int32)(unsafe.Add(mBase, uint32(v340)+4)) = v344
	v349 = v340 + int32(16)
	v350 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v349))), uint32(v350))
	*(*int64)(unsafe.Add(mBase, uint32(v349)+4)) = int64(-1)
	goto L67
L67:
	;
	v355 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v340)+28)), uint32(v355))
	v358 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v340)+40)) = v358
	*(*int32)(unsafe.Add(mBase, uint32(v340)+32)) = v355
	*(*int64)(unsafe.Add(mBase, uint32(v340)+48)) = v358
	F_table_parallelscan_initialize(m, l0, v340-int32(-64), v239)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	v369 = F_shm_toc_allocate(m, v368, v253)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v231)+44))
	F_tuplesort_initialize_shared(m, v369, v216, v371)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	F_shm_toc_insert(m, v374, int64(-5764607523034234879), v340)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	F_shm_toc_insert(m, v378, int64(-5764607523034234878), v369)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v383 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[20]))
	if v383 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	v385 = F_shm_toc_allocate(m, v384, v321)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v231)+12))
	v398 = F_mul_size(m, int32(32), v397)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L81
	}
L76:
	;
	if v321 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v388 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[20]))
	base.MemoryCopy(m, v385, v388, v321)
	goto L79
L78:
	;
	goto L79
L79:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	F_shm_toc_insert(m, v390, int64(-5764607523034234877), v385)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	goto L75
L81:
	;
	v400 = F_shm_toc_allocate(m, v395, v398)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	F_shm_toc_insert(m, v402, int64(-5764607523034234876), v400)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v231)+12))
	v409 = F_mul_size(m, int32(128), v408)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v411 = F_shm_toc_allocate(m, v406, v409)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	F_shm_toc_insert(m, v413, int64(-5764607523034234875), v411)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	F_LaunchParallelWorkers(m, v231)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v219))) = v231
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v231)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v219)+24)) = v411
	*(*int32)(unsafe.Add(mBase, uint32(v219)+20)) = v400
	*(*int32)(unsafe.Add(mBase, uint32(v219)+16)) = v239
	*(*int32)(unsafe.Add(mBase, uint32(v219)+12)) = v369
	*(*int32)(unsafe.Add(mBase, uint32(v219)+8)) = v340
	*(*int32)(unsafe.Add(mBase, uint32(v219)+4)) = v420 + int32(1)
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v231)+20))
	if v429 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	F__brin_end_parallel(m, v219)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v435 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[21]))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[1]))) = v219
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v219)+8))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v219)+12))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v219)+4))
	v442 = base.I32_div_s(v435, v441)
	F__gin_parallel_scan_and_build(m, v20+int32(16), v439, v440, l0, l1, v442, int32(1))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L92
	}
L91:
	;
	goto L30
L92:
	;
	F_WaitForParallelWorkersToAttach(m, v231)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	goto L30
L94:
	;
	v1188 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[15])))
	F_MemoryContextDelete(m, v1188)
	mBase = m.M
	v1190 = m.ExcPending
	if v1190 != 0 {
		goto L1
	} else {
		goto L221
	}
L95:
	;
	v461 = F_palloc0(m, int32(12))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L1
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v1064 = int32(0)
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(v1073)+140))
	v1075 = m.T0[v1074].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, l0, l1, l2, v1064, v1064, int32(1), v1064, int32(-1), int32(53), v20+int32(16), v1064)
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L1
	} else {
		goto L206
	}
L98:
	;
	v463 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v461))) = uint8(v463)
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[1])))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v465)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v461)+4)) = v466
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[1])))
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v468)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v461)+8)) = v469
	v472 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[21]))
	v473 = F_tuplesort_begin_index_gin(m, l1, v472, v461)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[22]))) = v473
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[1])))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v476)+8))
	v481 = v477 + int32(28)
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v476)+4))
	goto L100
L100:
	;
	v502 = base.AtomicRmwXchg32(m, v481, int32(0), int32(1))
	if v502 != 0 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	v516 = *(*float64)(unsafe.Add(mBase, uint32(v477)+40))
	*(*float64)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[9]))) = v516
	v518 = *(*float64)(unsafe.Add(mBase, uint32(v477)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[8]))) = v518
	v520 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v477)+28)), uint32(v520))
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L1
	} else {
		goto L110
	}
L102:
	;
	F_s_lock(m, v481, int32(_a_F_ginbuild_11), int32(1141), int32(_a_F_ginbuild_12))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L1
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v477)+32))
	if v482 != v508 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	goto L104
L106:
	;
	v510 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v481))), uint32(v510))
	F_ConditionVariableSleep(m, v477+int32(16), int32(134217767))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L1
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	goto L101
L109:
	;
	goto L100
L110:
	;
	v525 = *(*float64)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[9])))
	v530 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[17]))
	if v530 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[22])))
	F_tuplesort_performsort(m, v563)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L1
	} else {
		goto L115
	}
L112:
	;
	goto L111
L113:
	;
	v534 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ginbuild[18])))
	if v534&int32(1) == int32(0) {
		goto L112
	} else {
		goto L114
	}
L114:
	;
	v539 = int32(_a_F_ginbuild_1)
	v541 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10]))
	v542 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10])) = v541 + v542
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v530)))
	*(*int32)(unsafe.Add(mBase, uint32(v530))) = v545 + v542
	*(*int64)(unsafe.Add(mBase, uint32(v530+int32(80))+232)) = int64(5)
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v530)))
	*(*int32)(unsafe.Add(mBase, uint32(v530))) = v553 + v542
	v559 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10])) = v559 - v542
	goto L112
L115:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v567 = F_GinBufferInit(m, v566)
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v570 = *(*int64)(unsafe.Add(mBase, _c_F_ginbuild[23]))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[24]))) = v570
	v573 = *(*int64)(unsafe.Add(mBase, _c_F_ginbuild[25]))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[26]))) = v573
	*(*int64)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[27]))) = int64(6)
	v577 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[28]))) = v577
	v579 = *(*float64)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[8])))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[29]))) = base.I64_trunc_sat_f64_s(v579)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[30]))) = v577
	v586 = v20 + int32(_a_F_ginbuild_13)
	v588 = v20 + int32(_a_F_ginbuild_14)
	goto L119
L117:
	;
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[22])))
	v763 = F_tuplesort_getgintuple(m, v762, v588)
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L1
	} else {
		goto L134
	}
L118:
	;
	goto L117
L119:
	;
	v598 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[17]))
	if v598 == int32(0) {
		goto L118
	} else {
		goto L120
	}
L120:
	;
	v602 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ginbuild[18])))
	if v602&int32(1) == int32(0) {
		goto L118
	} else {
		goto L121
	}
L121:
	;
	v607 = int32(_a_F_ginbuild_1)
	v609 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10]))
	v610 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10])) = v609 + v610
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v598)))
	*(*int32)(unsafe.Add(mBase, uint32(v598))) = v613 + v610
	goto L123
L122:
	;
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v598)))
	v744 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v598))) = v743 + v744
	v747 = int32(_a_F_ginbuild_1)
	v749 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10])) = v749 - v744
	goto L118
L123:
	;
	v622 = v598 + int32(232)
	goto L124
L124:
	;
	v628 = int32(0)
	v631 = int32(0)
	goto L127
L127:
	;
	v637 = int32(2)
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v586+v631<<(uint(v637)%32))))
	v641 = int32(3)
	v647 = *(*int64)(unsafe.Add(mBase, uint32(v588+v631<<(uint(v641)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v622+v640<<(uint(v641)%32)))) = v647
	v650 = v631 | int32(1)
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v586+v650<<(uint(v637)%32))))
	v661 = *(*int64)(unsafe.Add(mBase, uint32(v588+v650<<(uint(v641)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v622+v654<<(uint(v641)%32)))) = v661
	v664 = v631 | v637
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v586+v664<<(uint(v637)%32))))
	v675 = *(*int64)(unsafe.Add(mBase, uint32(v588+v664<<(uint(v641)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v622+v668<<(uint(v641)%32)))) = v675
	v678 = v631 | v641
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v586+v678<<(uint(v637)%32))))
	v689 = *(*int64)(unsafe.Add(mBase, uint32(v588+v678<<(uint(v641)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v622+v682<<(uint(v641)%32)))) = v689
	v691 = int32(4)
	v694 = v628 + v691
	if v694 != int32(4) {
		v628 = v694
		v631 = v631 + v691
		goto L127
	} else {
		goto L129
	}
L128:
	;
	goto L122
L129:
	;
	goto L128
L134:
	;
	if v763 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v766 = v567 + int32(4)
	v769 = v763
	v782 = float64(0)
	goto L138
L136:
	;
	v981 = int64(1)
	goto L137
L137:
	;
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v567)+20))
	if v982 != 0 {
		goto L182
	} else {
		goto L183
	}
L138:
	;
	v785 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[31]))
	if v785 != 0 {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	v981 = base.I64_trunc_sat_f64_s(base.F64_add(v918, float64(1)))
	goto L137
L140:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L1
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v567)+20))
	if v788 == int32(0) {
		goto L145
	} else {
		goto L146
	}
L143:
	;
	goto L142
L144:
	;
	F_GinBufferStoreTuple(m, v567, v769)
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L1
	} else {
		goto L175
	}
L145:
	;
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v567)+24))
	if v868 < int32(1024) {
		goto L144
	} else {
		goto L168
	}
L146:
	;
	v791 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v769)+4)))
	v792 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v567))))
	if v791 != v792 {
		v829 = v788
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v832 = int32(_a_F_ginbuild_15)
	v833 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[13]))
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[14])))
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[13])) = v835
	v839 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v567))))
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v567)+4))
	v841 = int32(*(*int8)(unsafe.Add(mBase, uint32(v567)+2)))
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v567)+32))
	F_ginEntryInsert(m, v20+int32(16), v839, v840, v841, v842, v829, v458)
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L1
	} else {
		goto L162
	}
L148:
	;
	v794 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v769)+11)))
	v795 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567)+2)))
	if v794 != v795 {
		v829 = v788
		goto L147
	} else {
		goto L149
	}
L149:
	;
	if v794 != 0 {
		goto L145
	} else {
		goto L150
	}
L150:
	;
	v798 = v769 + int32(16)
	v799 = int32(1)
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v567)+4))
	v802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567)+14)))
	if v802 == v799 {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v798)))
	v806 = v805
	goto L153
L152:
	;
	v806 = v798
	goto L153
L153:
	;
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v567)+28))
	v808 = int32(36)
	v810 = v807 + v791*v808
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v810-int32(20))))
	v816 = m.T0[v815].(func(*base.Module, int32, int32, int32) int32)(m, v801, v806, v810-v808)
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	if v816 < int32(0) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v821 = v799
	goto L157
L156:
	;
	v821 = int32(0) - v816
	goto L157
L157:
	;
	v824 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v810-int32(28)))))
	if v824 != 0 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v825 = v821
	goto L160
L159:
	;
	v825 = v816
	goto L160
L160:
	;
	if v825 == int32(0) {
		goto L145
	} else {
		goto L161
	}
L161:
	;
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v567)+20))
	v829 = v828
	goto L147
L162:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[13])) = v833
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[14])))
	F_MemoryContextReset(m, v847)
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	v850 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567)+2)))
	if v850 != 0 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v855 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v567)+20)) = v855
	v857 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v567)+2)) = uint8(v857)
	*(*uint16)(unsafe.Add(mBase, uint32(v567))) = uint16(v857)
	*(*int32)(unsafe.Add(mBase, uint32(v766)+7)) = v857
	*(*int64)(unsafe.Add(mBase, uint32(v766))) = v855
	goto L144
L165:
	;
	v851 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567)+14)))
	if v851 != 0 {
		goto L164
	} else {
		goto L166
	}
L166:
	;
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v766)))
	F_pfree(m, v852)
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	goto L164
L168:
	;
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v567)+16))
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v769)+12))
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v567)+20))
	if v872+v873 < v871 {
		goto L144
	} else {
		goto L169
	}
L169:
	;
	v876 = int32(_a_F_ginbuild_15)
	v877 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[13]))
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[14])))
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[13])) = v879
	v883 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v567))))
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v567)+4))
	v885 = int32(*(*int8)(unsafe.Add(mBase, uint32(v567)+2)))
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v567)+32))
	F_ginEntryInsert(m, v20+int32(16), v883, v884, v885, v886, v868, v458)
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[13])) = v877
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[14])))
	F_MemoryContextReset(m, v891)
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v567)+20))
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v567)+24))
	v898 = (v894 - v895) * int32(6)
	if v898 != 0 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v567)+32))
	base.MemoryCopy(m, v899, v899+v895*int32(6), v898)
	goto L174
L173:
	;
	goto L174
L174:
	;
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v567)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v567)+24)) = int32(0)
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v567)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v567)+20)) = v908 - v905
	goto L144
L175:
	;
	v918 = base.F64_add(v782, float64(1))
	v922 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[17]))
	if v922 == int32(0) {
		goto L177
	} else {
		goto L178
	}
L176:
	;
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[22])))
	v958 = F_tuplesort_getgintuple(m, v955, v20+int32(_a_F_ginbuild_14))
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L1
	} else {
		goto L180
	}
L177:
	;
	goto L176
L178:
	;
	v926 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ginbuild[18])))
	if v926&int32(1) == int32(0) {
		goto L177
	} else {
		goto L179
	}
L179:
	;
	v931 = int32(_a_F_ginbuild_1)
	v933 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10]))
	v934 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10])) = v933 + v934
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v922)))
	*(*int32)(unsafe.Add(mBase, uint32(v922))) = v937 + v934
	*(*int64)(unsafe.Add(mBase, uint32(v922+int32(96))+232)) = base.I64_trunc_sat_f64_s(v918)
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v922)))
	*(*int32)(unsafe.Add(mBase, uint32(v922))) = v945 + v934
	v951 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10])) = v951 - v934
	goto L177
L180:
	;
	if v958 != 0 {
		v769 = v958
		v782 = v918
		goto L138
	} else {
		goto L181
	}
L181:
	;
	goto L139
L182:
	;
	v985 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v567))))
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v567)+4))
	v987 = int32(*(*int8)(unsafe.Add(mBase, uint32(v567)+2)))
	v988 = *(*int32)(unsafe.Add(mBase, uint32(v567)+32))
	F_ginEntryInsert(m, v20+int32(16), v985, v986, v987, v988, v982, v458)
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L1
	} else {
		goto L185
	}
L183:
	;
	goto L184
L184:
	;
	v1045 = *(*int32)(unsafe.Add(mBase, uint32(v567)+32))
	if v1045 != 0 {
		goto L194
	} else {
		goto L195
	}
L185:
	;
	v992 = v567 + int32(4)
	v993 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567)+2)))
	if v993 != 0 {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v998 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v567)+20)) = v998
	v1000 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v567)+2)) = uint8(v1000)
	*(*uint16)(unsafe.Add(mBase, uint32(v567))) = uint16(v1000)
	*(*int32)(unsafe.Add(mBase, uint32(v992)+7)) = v1000
	*(*int64)(unsafe.Add(mBase, uint32(v992))) = v998
	v1011 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[17]))
	if v1011 == v1000 {
		goto L191
	} else {
		goto L192
	}
L187:
	;
	v994 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567)+14)))
	if v994 != 0 {
		goto L186
	} else {
		goto L188
	}
L188:
	;
	v995 = *(*int32)(unsafe.Add(mBase, uint32(v992)))
	F_pfree(m, v995)
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	goto L186
L190:
	;
	goto L184
L191:
	;
	goto L190
L192:
	;
	v1015 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ginbuild[18])))
	if v1015&int32(1) == int32(0) {
		goto L191
	} else {
		goto L193
	}
L193:
	;
	v1020 = int32(_a_F_ginbuild_1)
	v1022 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10]))
	v1023 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10])) = v1022 + v1023
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(v1011)))
	*(*int32)(unsafe.Add(mBase, uint32(v1011))) = v1026 + v1023
	*(*int64)(unsafe.Add(mBase, uint32(v1011+int32(96))+232)) = v981
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v1011)))
	*(*int32)(unsafe.Add(mBase, uint32(v1011))) = v1034 + v1023
	v1040 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10])) = v1040 - v1023
	goto L191
L194:
	;
	F_pfree(m, v1045)
	mBase = m.M
	v1047 = m.ExcPending
	if v1047 != 0 {
		goto L1
	} else {
		goto L197
	}
L195:
	;
	goto L196
L196:
	;
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(v567)+20))
	if v1048 == int32(0) {
		goto L198
	} else {
		goto L199
	}
L197:
	;
	goto L196
L198:
	;
	F_pfree(m, v567)
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L1
	} else {
		goto L203
	}
L199:
	;
	v1051 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567)+2)))
	if v1051 != 0 {
		goto L198
	} else {
		goto L200
	}
L200:
	;
	v1052 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567)+14)))
	if v1052 != 0 {
		goto L198
	} else {
		goto L201
	}
L201:
	;
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v567)+4))
	F_pfree(m, v1053)
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		goto L1
	} else {
		goto L202
	}
L202:
	;
	goto L198
L203:
	;
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[22])))
	F_tuplesort_end(m, v1058)
	mBase = m.M
	v1060 = m.ExcPending
	if v1060 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[1])))
	F__brin_end_parallel(m, v1061)
	mBase = m.M
	v1063 = m.ExcPending
	if v1063 != 0 {
		goto L1
	} else {
		goto L205
	}
L205:
	;
	v1185 = v525
	goto L94
L206:
	;
	v1077 = int32(_a_F_ginbuild_15)
	v1078 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[13]))
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[14])))
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[13])) = v1080
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[32])))
	v1085 = m.G0
	v1086 = int32(16)
	v1087 = v1085 - v1086
	m.G0 = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[33]))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[34]))) = v1084
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v1084)))
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[35]))) = uint8(base.B2i32(v1094 == int32(_a_F_ginbuild_16)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[36]))) = int32(790)
	m.G0 = v1087 + v1086
	goto L207
L207:
	;
	v1111 = F_ginGetBAEntry(m, v172, v20+int32(12), v20+int32(_a_F_ginbuild_14), v20+int32(15), v20+int32(_a_F_ginbuild_13))
	mBase = m.M
	v1112 = m.ExcPending
	if v1112 != 0 {
		goto L1
	} else {
		goto L208
	}
L208:
	;
	if v1111 != 0 {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v1113 = v1111
	goto L212
L210:
	;
	goto L211
L211:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[13])) = v1078
	v1185 = v1075
	goto L94
L212:
	;
	v1131 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[31]))
	if v1131 != 0 {
		goto L214
	} else {
		goto L215
	}
L213:
	;
	goto L211
L214:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1133 = m.ExcPending
	if v1133 != 0 {
		goto L1
	} else {
		goto L217
	}
L215:
	;
	goto L216
L216:
	;
	v1136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+12)))
	v1137 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[27])))
	v1138 = int32(*(*int8)(unsafe.Add(mBase, uint32(v20)+15)))
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[26])))
	F_ginEntryInsert(m, v20+int32(16), v1136, v1137, v1138, v1113, v1139, v458)
	mBase = m.M
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L1
	} else {
		goto L218
	}
L217:
	;
	goto L216
L218:
	;
	v1150 = F_ginGetBAEntry(m, v172, v20+int32(12), v20+int32(_a_F_ginbuild_14), v20+int32(15), v20+int32(_a_F_ginbuild_13))
	mBase = m.M
	v1151 = m.ExcPending
	if v1151 != 0 {
		goto L1
	} else {
		goto L219
	}
L219:
	;
	if v1150 != 0 {
		v1113 = v1150
		goto L212
	} else {
		goto L220
	}
L220:
	;
	goto L213
L221:
	;
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[14])))
	F_MemoryContextDelete(m, v1191)
	mBase = m.M
	v1193 = m.ExcPending
	if v1193 != 0 {
		goto L1
	} else {
		goto L222
	}
L222:
	;
	v1195 = F_RelationGetNumberOfBlocksInFork(m, l1, int32(0))
	mBase = m.M
	v1196 = m.ExcPending
	if v1196 != 0 {
		goto L1
	} else {
		goto L223
	}
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[37]))) = v1195
	F_ginUpdateStats(m, l1, v458, int32(1))
	mBase = m.M
	v1200 = m.ExcPending
	if v1200 != 0 {
		goto L1
	} else {
		goto L224
	}
L224:
	;
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v1202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1201)+118)))
	if v1202 != int32(112) {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v1219 = F_palloc(m, int32(16))
	mBase = m.M
	v1220 = m.ExcPending
	if v1220 != 0 {
		goto L1
	} else {
		goto L234
	}
L226:
	;
	v1206 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[38]))
	if v1206 <= int32(0) {
		goto L227
	} else {
		goto L228
	}
L227:
	;
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v1209 != 0 {
		goto L225
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	v1211 = int32(0)
	v1213 = F_RelationGetNumberOfBlocksInFork(m, l1, v1211)
	mBase = m.M
	v1214 = m.ExcPending
	if v1214 != 0 {
		goto L1
	} else {
		goto L232
	}
L230:
	;
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v1210 != 0 {
		goto L225
	} else {
		goto L231
	}
L231:
	;
	goto L229
L232:
	;
	F_log_newpage_range(m, l1, v1211, v1213, int32(1))
	mBase = m.M
	v1217 = m.ExcPending
	if v1217 != 0 {
		goto L1
	} else {
		goto L233
	}
L233:
	;
	goto L225
L234:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1219))) = v1185
	v1222 = *(*float64)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_ginbuild[3])))
	*(*float64)(unsafe.Add(mBase, uint32(v1219)+8)) = v1222
	m.G0 = v20 + int32(_a_F_ginbuild_0)
	return v1219
L235:
	;
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v1232 + int32(4)
	F_errmsg_internal(m, int32(_a_F_ginbuild_17), v20)
	mBase = m.M
	v1238 = m.ExcPending
	if v1238 != 0 {
		goto L1
	} else {
		goto L236
	}
L236:
	;
	F_errfinish(m, int32(_a_F_ginbuild_11), int32(625), int32(_a_F_ginbuild_18))
	mBase = m.M
	v1243 = m.ExcPending
	if v1243 != 0 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gincost_pattern(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
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
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v87 float64
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v98 float64
	_ = v98
	var v102 int32
	_ = v102
	var v105 float64
	_ = v105
	var v109 float64
	_ = v109
	var v114 float64
	_ = v114
	var v117 int32
	_ = v117
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 float64
	_ = v134
	var v135 float64
	_ = v135
	var v138 float64
	_ = v138
	var v143 int32
	_ = v143
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	v6 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v6
	v25 = l1 << (uint(int32(2)) % 32)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25+v26)))
	F_get_op_opfamily_properties(m, l2, v28, v6, v12+int32(48), v12+int32(44), v12+int32(40))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		return int32(0)
	} else {
		v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v42 = *(*int32)(unsafe.Add(mBase, uint32(v40+v25)))
		v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		v45 = *(*int32)(unsafe.Add(mBase, uint32(v43+v25)))
		v47 = F_get_opfamily_proc(m, v42, v45, v45, int32(3))
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return int32(0)
		} else {
			if v47 != 0 {
				v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v49+v25)))
				v53 = v12 + int32(52)
				F_fmgr_info(m, v47, v53)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v56+v25)))
					F_set_fn_opclass_options(m, v53, v58)
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int32(0)
					} else {
						if v51 != 0 {
							v62 = v51
						} else {
							v62 = int32(100)
						}
						v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+48)))
						v74 = F_FunctionCall7Coll(m, v53, v62, l3, v12+int32(36), v65, v12+int32(32), v12+int32(28), v12+int32(24), v12+int32(20))
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return int32(0)
						} else {
							v76 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
							v77 = int32(0)
							v79 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
							v82 = base.B2i32(v76 != v77) | base.B2i32(v77 < v79)
							if v82 == v77 {
							} else {
								if int32(0) < v79 {
									v87 = *(*float64)(unsafe.Add(mBase, uint32(l4)+80))
									v89 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
									v90 = int32(0)
									v98 = v87
									for {
										if v89 == int32(0) {
											v109 = *(*float64)(unsafe.Add(mBase, uint32(l4)+72))
											*(*float64)(unsafe.Add(mBase, uint32(l4)+72)) = base.F64_add(v109, float64(1))
										} else {
											v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90+v89))))
											if v102 != int32(1) {
												v109 = *(*float64)(unsafe.Add(mBase, uint32(l4)+72))
												*(*float64)(unsafe.Add(mBase, uint32(l4)+72)) = base.F64_add(v109, float64(1))
											} else {
												v105 = *(*float64)(unsafe.Add(mBase, uint32(l4)+64))
												*(*float64)(unsafe.Add(mBase, uint32(l4)+64)) = base.F64_add(v105, float64(100))
											}
										}
										v114 = base.F64_add(v98, float64(1))
										*(*float64)(unsafe.Add(mBase, uint32(l4)+80)) = v114
										v117 = v90 + int32(1)
										if v117 != v79 {
											v90 = v117
											v98 = v114
											continue
										} else {
											break
										}
										break
									}
								} else {
								}
								switch v76 {
								case 0:
									v129 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l1+l4)+32)) = uint8(v129)
								case 1:
									v132 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l1+l4)+32)) = uint8(v132)
									v134 = *(*float64)(unsafe.Add(mBase, uint32(l4)+72))
									v135 = float64(1)
									*(*float64)(unsafe.Add(mBase, uint32(l4)+72)) = base.F64_add(v134, v135)
									v138 = *(*float64)(unsafe.Add(mBase, uint32(l4)+80))
									*(*float64)(unsafe.Add(mBase, uint32(l4)+80)) = base.F64_add(v138, v135)
								default:
									v143 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l1+l4))) = uint8(v143)
								}
							}
							m.G0 = v12 + int32(80)
							return v82
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v161 = m.ExcPending
				if v161 != 0 {
					return int32(0)
				} else {
					v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v163 = F_get_rel_name(m, v162)
					mBase = m.M
					v164 = m.ExcPending
					if v164 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v163
						*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = l1 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(3)
						F_errmsg_internal(m, int32(_a_F_gincost_pattern_0), v12)
						mBase = m.M
						v173 = m.ExcPending
						if v173 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_gincost_pattern_1), int32(_a_F_gincost_pattern_2), int32(_a_F_gincost_pattern_3))
							mBase = m.M
							v178 = m.ExcPending
							if v178 != 0 {
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
func F_ginoptions(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 float32
	_ = v39
	var v41 float32
	_ = v41
	var v44 float32
	_ = v44
	var v50 float32
	_ = v50
	var v53 float32
	_ = v53
	var v55 float32
	_ = v55
	var v57 float32
	_ = v57
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(0)
	v18 = l0 + l1*int32(28)
	if l3|l5 == int32(1) {
		v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+int32(3614)))))
		if v24 != 0 {
			if l5 != 0 {
				v53 = float32(0)
			} else {
				v53 = math.Float32frombits(uint32(0x7f800000))
			}
			if l3 != 0 {
				v55 = v53
			} else {
				v55 = math.Float32frombits(uint32(0x7f800000))
			}
			v57 = v55
			m.G0 = v12 + int32(16)
			return v57
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0+l1<<(uint(int32(2))%32))+uint32(_c_F_gistpenalty[0])))
			v35 = F_FunctionCall3Coll(m, v18+int32(3604), v32, l2, l4, v12+int32(12))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return float32(0)
			} else {
				v39 = float32(0)
				v41 = *(*float32)(unsafe.Add(mBase, uint32(v12)+12))
				if base.F32_lt(v41, v39) != 0 {
					v44 = v39
				} else {
					v44 = v41
				}
				if base.Ui32(int32(2139095040)) < base.Ui32(base.I32_reinterpret_f32(v41)&int32(2147483647)) {
					v50 = v39
				} else {
					v50 = v44
				}
				v57 = v50
				m.G0 = v12 + int32(16)
				return v57
			}
		}
	} else {
		v32 = *(*int32)(unsafe.Add(mBase, uint32(l0+l1<<(uint(int32(2))%32))+uint32(_c_F_gistpenalty[0])))
		v35 = F_FunctionCall3Coll(m, v18+int32(3604), v32, l2, l4, v12+int32(12))
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return float32(0)
		} else {
			v39 = float32(0)
			v41 = *(*float32)(unsafe.Add(mBase, uint32(v12)+12))
			if base.F32_lt(v41, v39) != 0 {
				v44 = v39
			} else {
				v44 = v41
			}
			if base.Ui32(int32(2139095040)) < base.Ui32(base.I32_reinterpret_f32(v41)&int32(2147483647)) {
				v50 = v39
			} else {
				v50 = v44
			}
			v57 = v50
			m.G0 = v12 + int32(16)
			return v57
		}
	}
}
func F_gistproperty(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	v7 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	if l1 == v7 {
		v62 = v7
		m.G0 = v12 + int32(16)
		return v62
	} else {
		switch l2 - int32(6) {
		case 0:
			v20 = int32(8)
			v21 = int32(1)
			v23 = F_get_index_column_opclass(m, l0, l1)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				if v23 == int32(0) {
					v57 = v21
					*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v57)
					v62 = v21
					m.G0 = v12 + int32(16)
					return v62
				} else {
					v33 = F_get_opclass_opfamily_and_input_type(m, v23, v12+int32(12), v12+int32(8))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						if v33 == int32(0) {
							v57 = v21
							*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v57)
							v62 = v21
							m.G0 = v12 + int32(16)
							return v62
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
							v40 = F_SearchSysCacheExists(m, int32(5), v38, v39, v39, v20)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v40)
								v43 = int32(0)
								if base.B2i32(l2 != int32(7))|v40 != 0 {
									v57 = v43
									*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v57)
									v62 = v21
									m.G0 = v12 + int32(16)
									return v62
								} else {
									v48 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
									v49 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
									v51 = F_SearchSysCacheExists(m, int32(5), v48, v49, v49, int32(3))
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return int32(0)
									} else {
										v54 = v51 ^ int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v54)
										v57 = v43
										*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v57)
										v62 = v21
										m.G0 = v12 + int32(16)
										return v62
									}
								}
							}
						}
					}
				}
			}
		case 1:
			v20 = int32(9)
			v21 = int32(1)
			v23 = F_get_index_column_opclass(m, l0, l1)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				if v23 == int32(0) {
					v57 = v21
					*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v57)
					v62 = v21
					m.G0 = v12 + int32(16)
					return v62
				} else {
					v33 = F_get_opclass_opfamily_and_input_type(m, v23, v12+int32(12), v12+int32(8))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						if v33 == int32(0) {
							v57 = v21
							*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v57)
							v62 = v21
							m.G0 = v12 + int32(16)
							return v62
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
							v40 = F_SearchSysCacheExists(m, int32(5), v38, v39, v39, v20)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v40)
								v43 = int32(0)
								if base.B2i32(l2 != int32(7))|v40 != 0 {
									v57 = v43
									*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v57)
									v62 = v21
									m.G0 = v12 + int32(16)
									return v62
								} else {
									v48 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
									v49 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
									v51 = F_SearchSysCacheExists(m, int32(5), v48, v49, v49, int32(3))
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return int32(0)
									} else {
										v54 = v51 ^ int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v54)
										v57 = v43
										*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v57)
										v62 = v21
										m.G0 = v12 + int32(16)
										return v62
									}
								}
							}
						}
					}
				}
			}
		default:
			v62 = v7
			m.G0 = v12 + int32(16)
			return v62
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
func F_gseg_union(m *base.Module, l0 int32) int32 {
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(12)
	if int32(2) <= v9 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = int32(1)
	v20 = v8
	goto L4
L2:
	;
	v41 = int32(0)
	goto L3
L3:
	;
	return v41
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v7+int32(4)+v18<<(uint(int32(4))%32))))
	v30 = F_DirectFunctionCall2Coll(m, int32(_a_F_gseg_union_0), int32(0), v20, v29)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v41 = v30
	goto L3
L6:
	;
	return int32(0)
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(12)
	v37 = v18 + int32(1)
	if v37 != v9 {
		v18 = v37
		v20 = v30
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
}
func F_gtsquery_penalty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v20 int32
	_ = v20
	var v43 int32
	_ = v43
	var v45 int64
	_ = v45
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
	v11 = v7 ^ v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = int32(0)
	v15 = int64(0)
	for {
		v20 = int32(1)
		v43 = base.I32_wrap_i64(int64(base.Ui64(v11)>>(uint(v15)%64)))&v20 + v14 + base.I32_wrap_i64(int64(base.Ui64(v11)>>(uint(v15|int64(1))%64)))&v20 + base.I32_wrap_i64(int64(base.Ui64(v11)>>(uint(v15|int64(2))%64)))&v20 + base.I32_wrap_i64(int64(base.Ui64(v11)>>(uint(v15|int64(3))%64)))&v20
		v45 = v15 + int64(4)
		if v45 != int64(64) {
			v14 = v43
			v15 = v45
			continue
		} else {
			break
		}
		break
	}
	*(*float32)(unsafe.Add(mBase, uint32(v12))) = base.F32_convert_i32_s(v43)
	return v12
}
func F_gtsquery_union(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v47 int64
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v63 int64
	_ = v63
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int64
	_ = v72
	var v76 int32
	_ = v76
	var v77 int64
	_ = v77
	var v78 int64
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v92 int64
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	v2 = int32(0)
	v9 = int64(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v12 <= v2 {
		v92 = v9
	} else {
		v16 = v12 & int32(3)
		v17 = int32(4)
		v18 = v11 + v17
		v19 = int32(0)
		if base.Ui32(v17) <= base.Ui32(v12) {
			v24 = v19
			v30 = v2
			v32 = v9
			for {
				v33 = int32(4)
				v35 = v18 + v24<<(uint(v33)%32)
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+48))
				v37 = *(*int64)(unsafe.Add(mBase, uint32(v36)))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+32))
				v39 = *(*int64)(unsafe.Add(mBase, uint32(v38)))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
				v41 = *(*int64)(unsafe.Add(mBase, uint32(v40)))
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
				v43 = *(*int64)(unsafe.Add(mBase, uint32(v42)))
				v47 = v37 | (v39 | (v41 | (v43 | v32)))
				v49 = v24 + v33
				v51 = v30 + v33
				if v51 != v12&int32(2147483644) {
					v24 = v49
					v30 = v51
					v32 = v47
					continue
				} else {
					break
				}
				break
			}
			if v16 == int32(0) {
				v92 = v47
			} else {
				v55 = v49
				v63 = v47
				v64 = v55
				v71 = v2
				v72 = v63
				for {
					v76 = *(*int32)(unsafe.Add(mBase, uint32(v18+v64<<(uint(int32(4))%32))))
					v77 = *(*int64)(unsafe.Add(mBase, uint32(v76)))
					v78 = v77 | v72
					v79 = int32(1)
					v82 = v71 + v79
					if v82 != v16 {
						v64 = v64 + v79
						v71 = v82
						v72 = v78
						continue
					} else {
						break
					}
					break
				}
				v92 = v78
			}
		} else {
			v55 = v19
			v63 = v9
			v64 = v55
			v71 = v2
			v72 = v63
			for {
				v76 = *(*int32)(unsafe.Add(mBase, uint32(v18+v64<<(uint(int32(4))%32))))
				v77 = *(*int64)(unsafe.Add(mBase, uint32(v76)))
				v78 = v77 | v72
				v79 = int32(1)
				v82 = v71 + v79
				if v82 != v16 {
					v64 = v64 + v79
					v71 = v82
					v72 = v78
					continue
				} else {
					break
				}
				break
			}
			v92 = v78
		}
	}
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(8)
	v95 = F_Int64GetDatum(m, v92)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		return int32(0)
	} else {
		return v95
	}
}
