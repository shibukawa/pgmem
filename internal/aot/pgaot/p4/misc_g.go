package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
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
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
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
	v111 = m.ExcPending
	if v111 != 0 {
		goto L6
	} else {
		goto L30
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
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	if v72 == int32(0) {
		v91 = v71
		v92 = v72
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
	if v92-v91 != 0 {
		goto L26
	} else {
		goto L27
	}
L19:
	;
	goto L18
L20:
	;
	if v71 != v72 {
		v91 = v71
		v92 = v72
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v76 = v68
	v77 = v55
	goto L22
L22:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+1)))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+1)))
	if v81 == int32(0) {
		v91 = v80
		v92 = v81
		goto L19
	} else {
		goto L24
	}
L23:
	;
	v91 = v80
	v92 = v81
	goto L19
L24:
	;
	v84 = int32(1)
	if v80 == v81 {
		v76 = v76 + v84
		v77 = v77 + v84
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v95 = v59 + int32(1)
	if v54 != v95 {
		v59 = v95
		goto L16
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	goto L17
L29:
	;
	goto L10
L30:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v112
	F_errmsg_internal(m, int32(_a_F_GetCTEForRTE_3), v10)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_GetCTEForRTE_1), int32(586), int32(_a_F_GetCTEForRTE_2))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L6
	} else {
		goto L32
	}
L32:
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
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
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
				v18 = F__emscripten_memcpy_bulkmem(m, v15, v8, int32(184))
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(l0)+260)) = v18
				return v8
			}
		}
	} else {
		if l1 != 0 {
			v23 = F_palloc(m, int32(184))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
				v27 = F__emscripten_memcpy_bulkmem(m, v23, v25, int32(184))
				mBase = m.M
				v29 = v23
				return v29
			}
		} else {
			v29 = v4
			return v29
		}
	}
}
func F_GetFlushRecPtr(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	v4 = int32(_a_F_GetFlushRecPtr_0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_GetFlushRecPtr[0]))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)+280))
	*(*int64)(unsafe.Add(mBase, uint32(v5)+280)) = v6
	*(*int64)(unsafe.Add(mBase, _c_F_GetFlushRecPtr[1])) = v6
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_GetFlushRecPtr[0]))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)+272))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+272)) = v12
	*(*int64)(unsafe.Add(mBase, _c_F_GetFlushRecPtr[2])) = v12
	if l0 != 0 {
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_GetFlushRecPtr[0]))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+308))
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v18
	} else {
	}
	v21 = *(*int64)(unsafe.Add(mBase, _c_F_GetFlushRecPtr[1]))
	return v21
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int64
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int64
	_ = v69
	var v70 int32
	_ = v70
	var v74 int64
	_ = v74
	var v75 int32
	_ = v75
	var v78 int64
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int64
	_ = v87
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
		v28 = int32(_a_F_GetLatestLSN_0)
		v29 = *(*int32)(unsafe.Add(mBase, _c_F_GetLatestLSN[1]))
		v30 = *(*int64)(unsafe.Add(mBase, uint32(v29)+280))
		*(*int64)(unsafe.Add(mBase, uint32(v29)+280)) = v30
		*(*int64)(unsafe.Add(mBase, _c_F_GetLatestLSN[2])) = v30
		v35 = *(*int32)(unsafe.Add(mBase, _c_F_GetLatestLSN[1]))
		v36 = *(*int64)(unsafe.Add(mBase, uint32(v35)+272))
		*(*int64)(unsafe.Add(mBase, uint32(v35)+272)) = v36
		*(*int64)(unsafe.Add(mBase, _c_F_GetLatestLSN[3])) = v36
		if l0 != 0 {
			v41 = *(*int32)(unsafe.Add(mBase, _c_F_GetLatestLSN[1]))
			v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+308))
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v42
		} else {
		}
		v45 = *(*int64)(unsafe.Add(mBase, _c_F_GetLatestLSN[2]))
		v87 = v45
		m.G0 = v9 + int32(16)
		return v87
	} else {
		v47 = *(*int32)(unsafe.Add(mBase, _c_F_GetLatestLSN[1]))
		v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+440))
		*(*int32)(unsafe.Add(mBase, uint32(v47)+440)) = int32(1)
		if v48 != 0 {
			v52 = *(*int32)(unsafe.Add(mBase, _c_F_GetLatestLSN[1]))
			F_s_lock(m, v52+int32(440), int32(_a_F_GetLatestLSN_1), int32(_a_F_GetLatestLSN_2), int32(_a_F_GetLatestLSN_3))
			mBase = m.M
			v61 = m.ExcPending
			if v61 != 0 {
				return int64(0)
			} else {
				v63 = *(*int32)(unsafe.Add(mBase, _c_F_GetLatestLSN[1]))
				*(*int32)(unsafe.Add(mBase, uint32(v63)+440)) = int32(0)
				v66 = *(*int32)(unsafe.Add(mBase, uint32(v63)+308))
				if v66 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v66
					v69 = F_GetXLogReplayRecPtr(m, int32(0))
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return int64(0)
					} else {
						v87 = v69
						m.G0 = v9 + int32(16)
						return v87
					}
				} else {
					v74 = F_GetWalRcvFlushRecPtr(m, int32(0), v9+int32(12))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int64(0)
					} else {
						v78 = F_GetXLogReplayRecPtr(m, v9+int32(8))
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return int64(0)
						} else {
							if base.Ui64(v78) < base.Ui64(v74) {
								v81 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v81
								v87 = v74
							} else {
								v83 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v83
								v87 = v78
							}
							m.G0 = v9 + int32(16)
							return v87
						}
					}
				}
			}
		} else {
			v63 = *(*int32)(unsafe.Add(mBase, _c_F_GetLatestLSN[1]))
			*(*int32)(unsafe.Add(mBase, uint32(v63)+440)) = int32(0)
			v66 = *(*int32)(unsafe.Add(mBase, uint32(v63)+308))
			if v66 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v66
				v69 = F_GetXLogReplayRecPtr(m, int32(0))
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return int64(0)
				} else {
					v87 = v69
					m.G0 = v9 + int32(16)
					return v87
				}
			} else {
				v74 = F_GetWalRcvFlushRecPtr(m, int32(0), v9+int32(12))
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return int64(0)
				} else {
					v78 = F_GetXLogReplayRecPtr(m, v9+int32(8))
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return int64(0)
					} else {
						if base.Ui64(v78) < base.Ui64(v74) {
							v81 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v81
							v87 = v74
						} else {
							v83 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v83
							v87 = v78
						}
						m.G0 = v9 + int32(16)
						return v87
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
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
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
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int64
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v183 int64
	_ = v183
	var v189 int32
	_ = v189
	var v191 int64
	_ = v191
	var v197 int32
	_ = v197
	var v200 int64
	_ = v200
	var v205 int32
	_ = v205
	var v207 int64
	_ = v207
	var v213 int32
	_ = v213
	var v216 int64
	_ = v216
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v242 int32
	_ = v242
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v355 int32
	_ = v355
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v422 int32
	_ = v422
	var v431 int32
	_ = v431
	var v446 int32
	_ = v446
	var v452 int32
	_ = v452
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v505 int32
	_ = v505
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v539 int32
	_ = v539
	var v544 int32
	_ = v544
	var v550 int32
	_ = v550
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
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
	v559 = m.ExcPending
	if v559 != 0 {
		goto L13
	} else {
		goto L91
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L13
	} else {
		goto L88
	}
L5:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v24<<(uint(int32(2))%32))+uint32(_c_F_GetLockConflicts[0])))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	if v38 < l1 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[1]))
	if base.Ui32(int32(2)) <= base.Ui32(v42) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[2]))
	v77 = F_get_hash_value(m, v76, l0)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L13
	} else {
		goto L16
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[3])) = v73
	goto L7
L9:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[3]))
	if v46 != 0 {
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[4]))
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[5]))
	v71 = F_palloc0(m, (v63+v65)<<(uint(int32(3))%32)+int32(8))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L13
	} else {
		goto L15
	}
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[6]))
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[4]))
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[5]))
	v58 = F_MemoryContextAlloc(m, v48, (v50+v52)<<(uint(int32(3))%32)+int32(8))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	v73 = v58
	goto L8
L15:
	;
	v73 = v71
	goto L8
L16:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[7]))
	v87 = v80 + v77&int32(15)<<(uint(int32(7))%32) + int32(_a_F_GetLockConflicts_0)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v88+l1<<(uint(int32(2))%32))))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	if v93 != int32(1) {
		v288 = v4
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v305 = F_LWLockAcquire(m, v87, int32(1))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L13
	} else {
		goto L48
	}
L18:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
	if v96 != 0 {
		v288 = v4
		goto L17
	} else {
		goto L19
	}
L19:
	;
	if base.Ui32(l1) < base.Ui32(int32(5)) {
		v288 = v4
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v99 == int32(0) {
		v288 = v4
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[8]))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+16))
	if v104 == int32(0) {
		v288 = v4
		goto L17
	} else {
		goto L22
	}
L22:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[9]))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v114 = (v108 - int32(1)) & (v111 * int32(_a_F_GetLockConflicts_1))
	v117 = int32(3)
	v124 = v103
	v126 = v4
	v132 = v4
	goto L23
L23:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	v145 = v142 + v132*int32(640)
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[10]))
	if v145 != v147 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v288 = v263
	goto L17
L25:
	;
	v150 = v145 + int32(584)
	v152 = F_LWLockAcquire(m, v150, int32(1))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L13
	} else {
		goto L28
	}
L26:
	;
	v263 = v126
	goto L27
L27:
	;
	v280 = v132 + int32(1)
	v282 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[8]))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v282)+16))
	if base.Ui32(v280) < base.Ui32(v283) {
		v124 = v282
		v126 = v263
		v132 = v280
		goto L23
	} else {
		goto L47
	}
L28:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v145)+60))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v154 != v155 {
		v242 = v126
		goto L29
	} else {
		goto L30
	}
L29:
	;
	F_LWLockRelease(m, v150)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L13
	} else {
		goto L46
	}
L30:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v145)+600))
	v159 = *(*int64)(unsafe.Add(mBase, uint32(v157+v114<<(uint(v117)%32))))
	if v159 == int64(0) {
		v242 = v126
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v162 = v157 + v114&int32(268435455)<<(uint(v117)%32)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v145)+604))
	v164 = v163 + v114<<(uint(int32(6))%32)
	v183 = int64(0)
	goto L32
L32:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v164+base.I32_wrap_i64(v183)<<(uint(int32(2))%32))))
	if v189 == v111 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	if v219<<(uint(int32(1))%32)&v92 == int32(0) {
		v242 = v126
		goto L29
	} else {
		goto L44
	}
L34:
	;
	goto L33
L35:
	;
	v191 = *(*int64)(unsafe.Add(mBase, uint32(v162)))
	v197 = base.I32_wrap_i64(int64(base.Ui64(v191)>>(uint(v183*int64(3))%64))) & int32(7)
	if v197 != 0 {
		v219 = v197
		goto L34
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v200 = v183 | int64(1)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v164+base.I32_wrap_i64(v200)<<(uint(int32(2))%32))))
	if v205 == v111 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L37
L39:
	;
	v207 = *(*int64)(unsafe.Add(mBase, uint32(v162)))
	v213 = base.I32_wrap_i64(int64(base.Ui64(v207)>>(uint(v200*int64(3))%64))) & int32(7)
	if v213 != 0 {
		v219 = v213
		goto L34
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v216 = v183 + int64(2)
	if v216 != int64(16) {
		v183 = v216
		goto L32
	} else {
		goto L43
	}
L42:
	;
	goto L41
L43:
	;
	v242 = v126
	goto L29
L44:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v145)+56))
	if v226 == int32(0) {
		v242 = v126
		goto L29
	} else {
		goto L45
	}
L45:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v145)+52))
	v231 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[3]))
	v234 = v231 + v126<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v234)+4)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v234))) = v229
	v242 = v126 + int32(1)
	goto L29
L46:
	;
	v263 = v242
	goto L27
L47:
	;
	goto L24
L48:
	;
	v308 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[2]))
	v309 = int32(0)
	v311 = F_hash_search_with_hash_value(m, v308, l0, v77, v309, v309)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L13
	} else {
		goto L52
	}
L49:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L13
	} else {
		goto L85
	}
L50:
	;
	m.G0 = v22 + int32(32)
	return v505
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v487
	v505 = v485
	goto L50
L52:
	;
	if v311 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	F_LWLockRelease(m, v87)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L13
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v311)+28))
	if v324 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v318 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[3]))
	*(*int64)(unsafe.Add(mBase, uint32(v318+v288<<(uint(int32(3))%32)))) = int64(4294967295)
	if l2 != 0 {
		v485 = v318
		v487 = v288
		goto L51
	} else {
		goto L57
	}
L57:
	;
	v505 = v318
	goto L50
L58:
	;
	F_LWLockRelease(m, v87)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L13
	} else {
		goto L82
	}
L59:
	;
	v452 = v288
	goto L58
L60:
	;
	goto L61
L61:
	;
	v328 = v311 + int32(24)
	if v328 == v324 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v452 = v288
	goto L58
L63:
	;
	goto L64
L64:
	;
	v331 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[3]))
	v334 = v324
	v338 = v288
	goto L65
L65:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v334-int32(8))))
	if v355&v92 == int32(0) {
		v431 = v338
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v452 = v431
	goto L58
L67:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v334)+4))
	if v446 != v328 {
		v334 = v446
		v338 = v431
		goto L65
	} else {
		goto L81
	}
L68:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v334-int32(16))))
	v363 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[10]))
	if v361 == v363 {
		v431 = v338
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v361)+56))
	if v365 == int32(0) {
		v431 = v338
		goto L67
	} else {
		goto L70
	}
L70:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v361)+52))
	v369 = int32(0)
	if base.B2i32(v288 <= int32(0)) == v369 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v373 = v369
	goto L74
L72:
	;
	goto L73
L73:
	;
	v422 = v331 + v338<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v422)+4)) = v365
	*(*int32)(unsafe.Add(mBase, uint32(v422))) = v368
	v431 = v338 + int32(1)
	goto L67
L74:
	;
	v393 = v331 + v373<<(uint(int32(3))%32)
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v393)))
	if v368 == v394 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	goto L73
L76:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v393)+4))
	if v396 == v365 {
		v431 = v338
		goto L67
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v399 = v373 + int32(1)
	if v399 != v288 {
		v373 = v399
		goto L74
	} else {
		goto L80
	}
L79:
	;
	goto L78
L80:
	;
	goto L75
L81:
	;
	goto L66
L82:
	;
	v470 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[4]))
	v472 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[5]))
	if v470+v472 < v452 {
		goto L49
	} else {
		goto L83
	}
L83:
	;
	v476 = *(*int32)(unsafe.Add(mBase, _c_F_GetLockConflicts[3]))
	*(*int64)(unsafe.Add(mBase, uint32(v476+v452<<(uint(int32(3))%32)))) = int64(4294967295)
	if l2 == int32(0) {
		v505 = v476
		goto L50
	} else {
		goto L84
	}
L84:
	;
	v485 = v476
	v487 = v452
	goto L51
L85:
	;
	F_errmsg_internal(m, int32(_a_F_GetLockConflicts_2), int32(0))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L13
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(_a_F_GetLockConflicts_3), int32(3233), int32(_a_F_GetLockConflicts_4))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L13
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = l1
	F_errmsg_internal(m, int32(_a_F_GetLockConflicts_5), v22+int32(16))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L13
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_GetLockConflicts_3), int32(3055), int32(_a_F_GetLockConflicts_4))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L13
	} else {
		goto L90
	}
L90:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v24
	F_errmsg_internal(m, int32(_a_F_GetLockConflicts_6), v22)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L13
	} else {
		goto L92
	}
L92:
	;
	F_errfinish(m, int32(_a_F_GetLockConflicts_3), int32(3052), int32(_a_F_GetLockConflicts_4))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L13
	} else {
		goto L93
	}
L93:
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
	var v46 int32
	_ = v46
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
	var v67 int32
	_ = v67
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
	return v67
L2:
	;
	return int32(0)
L3:
	;
	if v10 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v67 = v2
	goto L1
L5:
	;
	goto L6
L6:
	;
	v17 = F_RelationGetIndexAttrBitmap(m, l0, int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	if v17 == int32(0) {
		v67 = v2
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v23 = F_table_open(m, int32(2606), int32(1))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_ScanKeyInit(m, v8, int32(13), int32(3), int32(184), v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	F_ScanKeyInit(m, v8+int32(48), int32(4), int32(3), int32(61), int32(102))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v39 = int32(0)
	v43 = F_systable_beginscan(m, v23, v39, int32(1), v39, int32(2), v8)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v46 = v2
	goto L13
L13:
	;
	v50 = F_systable_getnext(m, v43)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L2
	} else {
		goto L15
	}
L14:
	;
	F_systable_endscan(m, v43)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L2
	} else {
		goto L21
	}
L15:
	;
	if v50 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+22)))
	v54 = v52 + v53
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+92))
	if v55 == int32(0) {
		goto L13
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	goto L14
L19:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v59 = F_lappend_oid(m, v46, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v46 = v59
	goto L13
L21:
	;
	F_sequence_close(m, v23, int32(1))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v67 = v46
	goto L1
}
func F_GetRedoRecPtr(m *base.Module) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	var v31 int64
	_ = v31
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_GetRedoRecPtr[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+440)) = int32(1)
	if v7 != 0 {
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_GetRedoRecPtr[0]))
		F_s_lock(m, v11+int32(440), int32(_a_F_GetRedoRecPtr_0), int32(_a_F_GetRedoRecPtr_1), int32(_a_F_GetRedoRecPtr_2))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int64(0)
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_GetRedoRecPtr[0]))
			*(*int32)(unsafe.Add(mBase, uint32(v22)+440)) = int32(0)
			v26 = *(*int64)(unsafe.Add(mBase, _c_F_GetRedoRecPtr[1]))
			v27 = *(*int64)(unsafe.Add(mBase, uint32(v22)+200))
			if base.Ui64(v26) < base.Ui64(v27) {
				*(*int64)(unsafe.Add(mBase, _c_F_GetRedoRecPtr[1])) = v27
				v31 = v27
			} else {
				v31 = v26
			}
			return v31
		}
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, _c_F_GetRedoRecPtr[0]))
		*(*int32)(unsafe.Add(mBase, uint32(v22)+440)) = int32(0)
		v26 = *(*int64)(unsafe.Add(mBase, _c_F_GetRedoRecPtr[1]))
		v27 = *(*int64)(unsafe.Add(mBase, uint32(v22)+200))
		if base.Ui64(v26) < base.Ui64(v27) {
			*(*int64)(unsafe.Add(mBase, _c_F_GetRedoRecPtr[1])) = v27
			v31 = v27
		} else {
			v31 = v26
		}
		return v31
	}
}
func F_GetSafeSnapshotBlockingPids(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	v4 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_GetSafeSnapshotBlockingPids[0]))
	v13 = F_LWLockAcquire(m, v9+int32(3584), int32(1))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_GetSafeSnapshotBlockingPids[1]))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if v19 == int32(0) {
		v78 = v4
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_GetSafeSnapshotBlockingPids[0]))
	F_LWLockRelease(m, v82+int32(3584))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L22
	}
L4:
	;
	v23 = v18 + int32(8)
	if v19 == v23 {
		v78 = v4
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v28 = v19
	goto L7
L6:
	;
	if v28 == int32(64) {
		v78 = v4
		goto L3
	} else {
		goto L11
	}
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)+48))
	if v32 == l0 {
		goto L6
	} else {
		goto L9
	}
L8:
	;
	v78 = v4
	goto L3
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v34 != v23 {
		v28 = v34
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+44)))
	if v38&int32(64) == int32(0) {
		v78 = v4
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v28)+28))
	if v43 == int32(0) {
		v78 = v4
		goto L3
	} else {
		goto L13
	}
L13:
	;
	v47 = v28 + int32(24)
	if v43 == v47 {
		v78 = v4
		goto L3
	} else {
		goto L14
	}
L14:
	;
	v49 = int32(1)
	if l2 <= v49 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v52 = v49
	goto L17
L16:
	;
	v52 = l2
	goto L17
L17:
	;
	v56 = v43
	v59 = int32(0)
	goto L18
L18:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+112))
	*(*int32)(unsafe.Add(mBase, uint32(l1+v59<<(uint(int32(2))%32)))) = v67
	v70 = v59 + int32(1)
	if v59 == v52-int32(1) {
		v78 = v70
		goto L3
	} else {
		goto L20
	}
L19:
	;
	v78 = v70
	goto L3
L20:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v72 != v47 {
		v56 = v72
		v59 = v70
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	return v78
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
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
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
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
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
	v116 = v12 + int32(32)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v118-v106 < v117 {
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
	v70 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
	*(*int64)(unsafe.Add(mBase, uint32(v69))) = v70
	v72 = *(*int64)(unsafe.Add(mBase, uint32(v47)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v69)+8)) = v72
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v38-int32(8))))
	*(*int32)(unsafe.Add(mBase, uint32(v69)+16)) = v76
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v50)+92))
	if v47 == v79 {
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
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v50)+100))
	v82 = v81
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
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_GetSingleProcBlockerStatusData[0]))
	v123 = v122 + v118
	v124 = v106 + v117
	if v124 < v123 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	if v136 != 0 {
		goto L26
	} else {
		goto L27
	}
L21:
	;
	v126 = v123
	goto L23
L22:
	;
	v126 = v124
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v126
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v131 = F_repalloc(m, v128, v126<<(uint(int32(2))%32))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L12
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v131
	goto L20
L25:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v175 - v176
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v179 - v180
	goto L3
L26:
	;
	v137 = v136
	goto L28
L27:
	;
	v137 = v116
	goto L28
L28:
	;
	if v137 == v116 {
		goto L25
	} else {
		goto L29
	}
L29:
	;
	if l0 == v137 {
		goto L25
	} else {
		goto L30
	}
L30:
	;
	v142 = v137
	goto L31
L31:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v142)+44))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v152 + int32(1)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v156+v152<<(uint(int32(2))%32)))) = v151
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	if v161 == v116 {
		goto L25
	} else {
		goto L33
	}
L32:
	;
	goto L25
L33:
	;
	if l0 != v161 {
		v142 = v161
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
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
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v392 int32
	_ = v392
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v421 int32
	_ = v421
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v432 int64
	_ = v432
	var v433 int64
	_ = v433
	var v440 int32
	_ = v440
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
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v511 int64
	_ = v511
	var v513 int64
	_ = v513
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v553 int32
	_ = v553
	var v556 int64
	_ = v556
	var v562 int64
	_ = v562
	var v564 int32
	_ = v564
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
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
	return v337
L6:
	;
	goto L5
L7:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v312)+24))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v312)+20))
	v315 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v312)+24)) = (v313 + v315) & int32(-4194305)
	v321 = int32(_a_F_GetVictimBuffer_0)
	v322 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[1])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v322)+4)) = v315
	v329 = v314 + v315
	*(*int32)(unsafe.Add(mBase, uint32(v322))) = v329
	v332 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[0]))
	F_ResourceOwnerRemember(m, v332, v329, int32(_a_F_GetVictimBuffer_1))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L74
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
	v312 = v61
	goto L7
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+24)) = int32(-1)
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[4]))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	F_SetLatch(m, v88+v81*int32(640)+int32(20))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	v98 = v80
	goto L21
L21:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v98)+20)) = v99 + int32(1)
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[2]))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+8))
	if v105 < int32(0) {
		v168 = v104
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[2]))
	v98 = v97
	goto L21
L23:
	;
	v175 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[5]))
	v180 = v168
	v182 = v175
	goto L41
L24:
	;
	v112 = v104
	goto L25
L25:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	*(*int32)(unsafe.Add(mBase, uint32(v112))) = int32(1)
	if v118 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[2]))
	F_s_lock(m, v122, int32(_a_F_GetVictimBuffer_3), int32(273), int32(_a_F_GetVictimBuffer_4))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v129 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[2]))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+8))
	if v130 < int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L29
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v129))) = int32(0)
	v168 = v129
	goto L23
L32:
	;
	goto L33
L33:
	;
	v136 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[3]))
	v139 = v136 + v130<<(uint(int32(6))%32)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v129)+8)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v139)+32)) = int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v129))) = int32(0)
	v146 = F_LockBufHdr(m, v139)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	if v146&int32(_a_F_GetVictimBuffer_5) != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v139)+24)) = v146 & int32(-4194305)
	v154 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[2]))
	v112 = v154
	goto L25
L36:
	;
	if l0 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v139)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0+v155<<(uint(int32(2))%32))+12)) = v159 + int32(1)
	goto L40
L39:
	;
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v146
	v312 = v139
	goto L7
L41:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v180)+4))
	v188 = v186 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v180)+4)) = v188
	v191 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[5]))
	if base.Ui32(v186) < base.Ui32(v191) {
		v241 = v186
		goto L43
	} else {
		goto L44
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v253)+24)) = v254 & int32(-4194305)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L71
	}
L43:
	;
	v250 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[3]))
	v253 = v250 + v241<<(uint(int32(6))%32)
	v254 = F_LockBufHdr(m, v253)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L60
	}
L44:
	;
	v193 = base.I32_rem_u_s(v186, v191)
	if v193 != 0 {
		v241 = v193
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v195 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[2]))
	v199 = v188
	v200 = v195
	goto L46
L46:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v200)))
	*(*int32)(unsafe.Add(mBase, uint32(v200))) = int32(1)
	if v206 != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v230 = int32(0)
	v232 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v232))) = v230
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v232)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v232)+16)) = v235 + int32(1)
	v241 = v230
	goto L43
L48:
	;
	v210 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[2]))
	F_s_lock(m, v210, int32(_a_F_GetVictimBuffer_3), int32(151), int32(_a_F_GetVictimBuffer_6))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v217 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[2]))
	v219 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[5]))
	v220 = base.I32_rem_u_s(v199, v219)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v217)+4))
	if v221 == v199 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L50
L52:
	;
	v223 = v220
	goto L54
L53:
	;
	v223 = v221
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v217)+4)) = v223
	if v221 != v199 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v227 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v227))) = int32(0)
	v199 = v221
	v200 = v227
	goto L46
L56:
	;
	goto L57
L57:
	;
	goto L47
L58:
	;
	goto L42
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v253)+24)) = v279 & int32(-4194305)
	v285 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[2]))
	v180 = v285
	v182 = v280
	goto L41
L60:
	;
	if v254&int32(_a_F_GetVictimBuffer_7) == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	if v254&int32(_a_F_GetVictimBuffer_8) != 0 {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	goto L63
L63:
	;
	v276 = v182 - int32(1)
	if v276 == int32(0) {
		goto L58
	} else {
		goto L70
	}
L64:
	;
	v265 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[5]))
	v279 = v254 - int32(_a_F_GetVictimBuffer_9)
	v280 = v265
	goto L59
L65:
	;
	goto L66
L66:
	;
	if l0 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v253)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0+v266<<(uint(int32(2))%32))+12)) = v270 + int32(1)
	goto L69
L68:
	;
	goto L69
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v254
	v312 = v253
	goto L7
L70:
	;
	v279 = v254
	v280 = v276
	goto L59
L71:
	;
	F_errmsg_internal(m, int32(_a_F_GetVictimBuffer_10), int32(0))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_GetVictimBuffer_3), int32(353), int32(_a_F_GetVictimBuffer_4))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	v337 = v314 + int32(1)
	F_CheckBufferIsPinnedOnce(m, v337)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+6)))
	if v340&int32(128) == int32(0) {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v592 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[0]))
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v312)+20))
	F_ResourceOwnerForget(m, v592, v593+int32(1), int32(_a_F_GetVictimBuffer_1))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L1
	} else {
		goto L128
	}
L77:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v536&int32(16777216) != 0 {
		goto L118
	} else {
		goto L119
	}
L78:
	;
	v346 = v312 + int32(48)
	v348 = F_LWLockConditionalAcquire(m, v346, int32(1))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	if v348 == int32(0) {
		goto L76
	} else {
		goto L80
	}
L80:
	;
	if l0 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	F_FlushBuffer(m, v312, int32(0), l1)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L1
	} else {
		goto L109
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = int32(_a_F_GetVictimBuffer_11)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = int32(_a_F_GetVictimBuffer_12)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = int32(_a_F_GetVictimBuffer_13)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = int64(0)
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v312)+24))
	v365 = int32(_a_F_GetVictimBuffer_14)
	*(*int32)(unsafe.Add(mBase, uint32(v312)+24)) = v364 | v365
	if v364&v365 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	goto L86
L84:
	;
	v392 = v364
	goto L85
L85:
	;
	v403 = int32(_a_F_GetVictimBuffer_15)
	v404 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[6]))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(8))+8))
	if v406 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L86:
	;
	F_perform_spin_delay(m, v13+int32(8))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L88
	}
L87:
	;
	v392 = v384
	goto L85
L88:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v312)+24))
	v385 = int32(_a_F_GetVictimBuffer_14)
	*(*int32)(unsafe.Add(mBase, uint32(v312)+24)) = v384 | v385
	if v384&v385 != 0 {
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v392 | int32(_a_F_GetVictimBuffer_14)
	v427 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[7]))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v312)+20))
	v431 = v427 + v428<<(uint(int32(13))%32)
	v432 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v431)+4)))
	v433 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v431))))
	*(*int32)(unsafe.Add(mBase, uint32(v312)+24)) = v392 & int32(-4194305)
	v440 = F_XLogNeedsFlush(m, v432|v433<<(uint(int64(32))%64))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L101
	}
L91:
	;
	goto L90
L92:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[6])) = v421
	goto L91
L93:
	;
	if int32(999) < v404 {
		goto L91
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	if v404 < int32(11) {
		goto L91
	} else {
		goto L100
	}
L96:
	;
	v411 = int32(900)
	if v411 <= v404 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v414 = v411
	goto L99
L98:
	;
	v414 = v404
	goto L99
L99:
	;
	v421 = v414 + int32(100)
	goto L92
L100:
	;
	v421 = v404 - int32(1)
	goto L92
L101:
	;
	if v440 == int32(0) {
		goto L81
	} else {
		goto L102
	}
L102:
	;
	v444 = int32(0)
	v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+3)))
	if v445 == v444 {
		v465 = v444
		goto L103
	} else {
		goto L104
	}
L103:
	;
	if v465 == int32(0) {
		goto L81
	} else {
		goto L107
	}
L104:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v448 != int32(1) {
		v465 = v444
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v456 = l0 + v451<<(uint(int32(2))%32) + int32(12)
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v456)))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v312)+20))
	if v457 != v458+int32(1) {
		v465 = v444
		goto L103
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v456))) = int32(0)
	v465 = int32(1)
	goto L103
L107:
	;
	F_LWLockRelease(m, v346)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	goto L76
L109:
	;
	F_LWLockRelease(m, v346)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	v487 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetVictimBuffer[8])))
	if v487&int32(1) != 0 {
		goto L77
	} else {
		goto L111
	}
L111:
	;
	v491 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetVictimBuffer[9])))
	if v491 != int32(1) {
		goto L77
	} else {
		goto L112
	}
L112:
	;
	v495 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[10]))
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v495)))
	if int32(0) < v496 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v499 = int32(_a_F_GetVictimBuffer_16)
	v501 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[11])) = v501 + int32(1)
	v506 = v501 * int32(20)
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v312)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v506)+uint32(_c_F_GetVictimBuffer[12]))) = v509
	v511 = *(*int64)(unsafe.Add(mBase, uint32(v312)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v506)+uint32(_c_F_GetVictimBuffer[13]))) = v511
	v513 = *(*int64)(unsafe.Add(mBase, uint32(v312)))
	*(*int64)(unsafe.Add(mBase, uint32(v506)+uint32(_c_F_GetVictimBuffer[14]))) = v513
	v516 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[10]))
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v516)))
	v519 = v517
	goto L115
L114:
	;
	v519 = v496
	goto L115
L115:
	;
	v521 = *(*int32)(unsafe.Add(mBase, _c_F_GetVictimBuffer[11]))
	if v521 < v519 {
		goto L77
	} else {
		goto L116
	}
L116:
	;
	F_IssuePendingWritebacks(m, int32(_a_F_GetVictimBuffer_17), l1)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	goto L77
L118:
	;
	v539 = int32(0)
	v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+3)))
	if v542 != 0 {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	v574 = v536
	goto L120
L120:
	;
	if v574&int32(33554432) == int32(0) {
		goto L6
	} else {
		goto L125
	}
L121:
	;
	v543 = int32(3)
	goto L123
L122:
	;
	v543 = v539
	goto L123
L123:
	;
	v553 = int32(0) + l1<<(uint(int32(6))%32) + v543<<(uint(int32(3))%32)
	v556 = *(*int64)(unsafe.Add(mBase, uint32(v553)+uint32(_c_F_GetVictimBuffer[15])))
	*(*int64)(unsafe.Add(mBase, uint32(v553)+uint32(_c_F_GetVictimBuffer[15]))) = v556 + int64(1)
	v562 = *(*int64)(unsafe.Add(mBase, uint32(v553)+uint32(_c_F_GetVictimBuffer[16])))
	*(*int64)(unsafe.Add(mBase, uint32(v553)+uint32(_c_F_GetVictimBuffer[16]))) = v562
	v564 = int32(1)
	F_pgstat_count_backend_io_op(m, v539, l1, v543, v564, int64(0))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, _c_F_GetVictimBuffer[17])) = uint8(v564)
	*(*uint8)(unsafe.Add(mBase, _c_F_GetVictimBuffer[18])) = uint8(v564)
	goto L124
L124:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v574 = v573
	goto L120
L125:
	;
	v579 = F_InvalidateVictimBuffer(m, v312)
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	if v579 != 0 {
		goto L6
	} else {
		goto L127
	}
L127:
	;
	goto L76
L128:
	;
	F_UnpinBufferNoOwner(m, v312)
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
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
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int64
	_ = v68
	var v70 int32
	_ = v70
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
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v45<<(uint(int32(2))%32))+uint32(_c_F_GlobalVisCheckRemovableFullXid[1])))
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v50)+8))
	if base.Ui64(l1) < base.Ui64(v51) {
		v70 = int32(1)
		m.G0 = v8 + int32(48)
		return v70
	} else {
		v54 = int32(0)
		v55 = *(*int64)(unsafe.Add(mBase, uint32(v50)))
		if base.Ui64(v55) <= base.Ui64(l1) {
			v70 = v54
			m.G0 = v8 + int32(48)
			return v70
		} else {
			v58 = *(*int32)(unsafe.Add(mBase, _c_F_GlobalVisCheckRemovableFullXid[2]))
			if v58 != 0 {
				v60 = *(*int32)(unsafe.Add(mBase, _c_F_GlobalVisCheckRemovableFullXid[3]))
				if v60 == v58 {
					v70 = v54
					m.G0 = v8 + int32(48)
					return v70
				} else {
					F_ComputeXidHorizons(m, v8+int32(8))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						v68 = *(*int64)(unsafe.Add(mBase, uint32(v50)+8))
						v70 = base.B2i32(base.Ui64(l1) < base.Ui64(v68))
						m.G0 = v8 + int32(48)
						return v70
					}
				}
			} else {
				F_ComputeXidHorizons(m, v8+int32(8))
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return int32(0)
				} else {
					v68 = *(*int64)(unsafe.Add(mBase, uint32(v50)+8))
					v70 = base.B2i32(base.Ui64(l1) < base.Ui64(v68))
					m.G0 = v8 + int32(48)
					return v70
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
	var v42 int32
	_ = v42
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
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v37<<(uint(int32(2))%32))+uint32(_c_F_GlobalVisTestFor[1])))
	return v42
}
func F___gettimeofday(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 float64
	_ = v5
	var v7 float64
	_ = v7
	var v11 int64
	_ = v11
	var v13 int64
	_ = v13
	var v20 float64
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	v5 = m.Env.Emscripten_date_now(m)
	mBase = m.M
	v7 = base.F64_div(v5, float64(1000))
	if base.F64_lt(base.F64_abs(v7), float64(9.223372036854776e+18)) != 0 {
		v11 = base.I64_trunc_f64_s(v7)
		v13 = v11
	} else {
		v13 = int64(-9223372036854775807 - 1)
	}
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v13
	v20 = base.F64_mul(base.F64_sub(v5, base.F64_convert_i64_s(v13*int64(1000))), float64(1000))
	if base.F64_lt(base.F64_abs(v20), float64(2.147483648e+09)) != 0 {
		v24 = base.I32_trunc_f64_s(v20)
		v26 = v24
	} else {
		v26 = int32(-2147483648)
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v26
	return
}
func F___gmtime_r(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	F_do_tzset(m)
	mBase = m.M
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	m.Env.X_gmtime_js(m, v4, l1)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = int32(_a_F___gmtime_r_0)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+32)) = int64(0)
	return l1
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
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 float32
	_ = v30
	var v39 float32
	_ = v39
	var v46 int32
	_ = v46
	var v54 float64
	_ = v54
	var v55 float64
	_ = v55
	var v56 float64
	_ = v56
	var v66 float64
	_ = v66
	var v67 float64
	_ = v67
	var v68 float64
	_ = v68
	var v80 float64
	_ = v80
	var v81 float64
	_ = v81
	var v82 float64
	_ = v82
	var v97 float64
	_ = v97
	var v99 float64
	_ = v99
	var v111 float32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 float32
	_ = v115
	var v118 float32
	_ = v118
	var v120 float32
	_ = v120
	var v123 float32
	_ = v123
	var v124 float32
	_ = v124
	var v125 float32
	_ = v125
	var v128 float32
	_ = v128
	var v133 float64
	_ = v133
	var v141 int32
	_ = v141
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v20 = base.I32_div_s(v16+int32(1), int32(2))
	if l3 < v20 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v23 = base.I32_div_s(v16, int32(2))
	if l5 < v23 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v26 = l3
	goto L3
L3:
	;
	v27 = v16 - v26
	if v26 < v27 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v25 = l5
	goto L6
L5:
	;
	v25 = v23
	goto L6
L6:
	;
	v26 = v25
	goto L3
L7:
	;
	v29 = v26
	goto L9
L8:
	;
	v29 = v27
	goto L9
L9:
	;
	v30 = base.F32_convert_i32_s(v29)
	if v16 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L61
	} else {
		goto L64
	}
L11:
	;
	F_float_underflow_error(m)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L61
	} else {
		goto L63
	}
L12:
	;
	F_float_zero_divide_error(m)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L61
	} else {
		goto L62
	}
L13:
	;
	if base.Ui32(base.I32_reinterpret_f32(v30)&int32(2147483647)) <= base.Ui32(int32(2139095040)) {
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v39 = base.F32_div(v30, base.F32_convert_i32_s(v16))
	if base.F32_eq(base.F32_abs(v39), math.Float32frombits(uint32(0x7f800000))) != 0 {
		goto L10
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	if v29 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v46 = base.F32_eq(v39, float32(0))
	goto L20
L19:
	;
	v46 = int32(0)
	goto L20
L20:
	;
	if v46 != 0 {
		goto L11
	} else {
		goto L21
	}
L21:
	;
	if base.F64_gt(base.F64_promote_f32(v39), float64(0.3)) == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	return
L23:
	;
	if l1 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v81 = base.F64_sub(l4, l2)
	v82 = base.F64_abs(v81)
	if base.F64_ne(v82, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L34
	} else {
		goto L35
	}
L25:
	;
	v54 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
	v55 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	v56 = base.F64_sub(v54, v55)
	if base.F64_ne(base.F64_abs(v56), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v80 = v56
		goto L24
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v66 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
	v67 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
	v68 = base.F64_sub(v66, v67)
	if base.F64_ne(base.F64_abs(v68), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v80 = v68
		goto L24
	} else {
		goto L31
	}
L28:
	;
	if base.F64_eq(base.F64_abs(v54), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v80 = v56
		goto L24
	} else {
		goto L29
	}
L29:
	;
	if base.F64_eq(base.F64_abs(v55), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v80 = v56
		goto L24
	} else {
		goto L30
	}
L30:
	;
	goto L10
L31:
	;
	if base.F64_eq(base.F64_abs(v66), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v80 = v68
		goto L24
	} else {
		goto L32
	}
L32:
	;
	if base.F64_ne(base.F64_abs(v67), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L10
	} else {
		goto L33
	}
L33:
	;
	v80 = v68
	goto L24
L34:
	;
	if base.B2i32(base.Ui64(base.I64_reinterpret_f64(v82)) <= base.Ui64(int64(9218868437227405312)))&base.F64_eq(v80, float64(0)) != 0 {
		goto L12
	} else {
		goto L38
	}
L35:
	;
	if base.F64_eq(base.F64_abs(l2), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	if base.F64_ne(base.F64_abs(l4), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L10
	} else {
		goto L37
	}
L37:
	;
	goto L34
L38:
	;
	v97 = base.F64_div(v81, v80)
	v99 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(base.F64_abs(v97), v99)&base.F64_ne(v82, v99) != 0 {
		goto L10
	} else {
		goto L39
	}
L39:
	;
	if base.F64_ne(v97, float64(0)) != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v111 = base.F32_demote_f64(v97)
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	if v112 != 0 {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	if base.F64_eq(v81, float64(0)) != 0 {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	if base.F64_ne(base.F64_abs(v80), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L11
	} else {
		goto L43
	}
L43:
	;
	goto L40
L44:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+80)) = v80
	*(*float32)(unsafe.Add(mBase, uint32(l0)+64)) = v39
	v141 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)) = uint8(v141)
	*(*float32)(unsafe.Add(mBase, uint32(l0)+68)) = v111
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = l1
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = l4
	goto L22
L45:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if l1 == v113 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v115 = *(*float32)(unsafe.Add(mBase, uint32(l0)+68))
	if base.F32_gt(v115, v111) != 0 {
		goto L44
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v120 = float32(0)
	if base.F32_ge(v111, v120) != 0 {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	if base.F32_ne(v111, v115) != 0 {
		goto L22
	} else {
		goto L50
	}
L50:
	;
	v118 = *(*float32)(unsafe.Add(mBase, uint32(l0)+64))
	if base.F32_gt(v39, v118) != 0 {
		goto L44
	} else {
		goto L51
	}
L51:
	;
	goto L22
L52:
	;
	v123 = v111
	goto L54
L53:
	;
	v123 = v120
	goto L54
L54:
	;
	v124 = *(*float32)(unsafe.Add(mBase, uint32(l0)+68))
	v125 = float32(0)
	if base.F32_ge(v124, v125) != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v128 = v124
	goto L57
L56:
	;
	v128 = v125
	goto L57
L57:
	;
	if base.F32_lt(v123, v128) != 0 {
		goto L44
	} else {
		goto L58
	}
L58:
	;
	if base.F32_ge(v128, v123) == int32(0) {
		goto L22
	} else {
		goto L59
	}
L59:
	;
	v133 = *(*float64)(unsafe.Add(mBase, uint32(l0)+80))
	if base.F64_gt(v80, v133) == int32(0) {
		goto L22
	} else {
		goto L60
	}
L60:
	;
	goto L44
L61:
	;
	return
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_g_intbig_same(m *base.Module, l0 int32) int32 {
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
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
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
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v78 int32
	_ = v78
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v10 == v2 {
		v27 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v27&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	if v14 == int32(0) {
		v27 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v17 != int32(7) {
		v27 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v20 != int32(17) {
		v27 = v2
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+24)))
	v27 = v23 ^ int32(1)
	goto L2
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v31 = F_get_fn_opclass_options(m, v30)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v36 = int32(252)
	goto L9
L9:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v38 = int32(4)
	v39 = v37 & v38
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+4)))
	if v40&v38 != 0 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	return int32(0)
L11:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v36 = v35
	goto L9
L12:
	;
	return v6
L13:
	;
	v78 = int32(base.Ui32(v39) >> (uint(int32(2)) % 32))
	goto L15
L14:
	;
	if v39 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(v78)
	goto L12
L16:
	;
	v78 = int32(0)
	goto L15
L17:
	;
	v45 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(v45)
	v47 = int32(0)
	if v36 <= v47 {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	v50 = int32(8)
	v54 = v47
	goto L19
L19:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54+(v8+v50)))))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54+(v7+v50)))))
	if v60 != v62 {
		goto L16
	} else {
		goto L21
	}
L20:
	;
	goto L12
L21:
	;
	v65 = v54 + int32(1)
	if v36 != v65 {
		v54 = v65
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
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
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
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
			v94 = v64
			v95 = int32(1)
			v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
			v100 = *(*int32)(unsafe.Add(mBase, uint32(v96+l1<<(uint(int32(2))%32))))
			v102 = F_ExecStoreMinimalTuple(m, v94, v100, v95)
			mBase = m.M
			v103 = m.ExcPending
			if v103 != 0 {
				return int32(0)
			} else {
				v104 = v95
				return v104
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
								v104 = v70
								return v104
							} else {
								v87 = F_heap_copy_minimal_tuple(m, v82, int32(0))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return int32(0)
								} else {
									if v87 == int32(0) {
										v104 = v70
										return v104
									} else {
										F_load_tuple_array(m, l0, l1)
										mBase = m.M
										v92 = m.ExcPending
										if v92 != 0 {
											return int32(0)
										} else {
											v94 = v87
											v95 = int32(1)
											v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
											v100 = *(*int32)(unsafe.Add(mBase, uint32(v96+l1<<(uint(int32(2))%32))))
											v102 = F_ExecStoreMinimalTuple(m, v94, v100, v95)
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
												return int32(0)
											} else {
												v104 = v95
												return v104
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
							v104 = v70
							return v104
						} else {
							v87 = F_heap_copy_minimal_tuple(m, v82, int32(0))
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return int32(0)
							} else {
								if v87 == int32(0) {
									v104 = v70
									return v104
								} else {
									F_load_tuple_array(m, l0, l1)
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int32(0)
									} else {
										v94 = v87
										v95 = int32(1)
										v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
										v100 = *(*int32)(unsafe.Add(mBase, uint32(v96+l1<<(uint(int32(2))%32))))
										v102 = F_ExecStoreMinimalTuple(m, v94, v100, v95)
										mBase = m.M
										v103 = m.ExcPending
										if v103 != 0 {
											return int32(0)
										} else {
											v104 = v95
											return v104
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
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
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
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_gbtreekey_in_1), int32(34), int32(_a_F_gbtreekey_in_2))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v235 int32
	_ = v235
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
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
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
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
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
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
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v578 int32
	_ = v578
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v643 int32
	_ = v643
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v746 int32
	_ = v746
	var v764 int32
	_ = v764
	var v770 int32
	_ = v770
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v793 int32
	_ = v793
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v874 int32
	_ = v874
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v947 int32
	_ = v947
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v970 int32
	_ = v970
	var v976 int32
	_ = v976
	var v982 int32
	_ = v982
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v997 int32
	_ = v997
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1037 int32
	_ = v1037
	var v1040 int32
	_ = v1040
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1058 int32
	_ = v1058
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1101 int32
	_ = v1101
	var v1105 int32
	_ = v1105
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1134 int32
	_ = v1134
	var v1137 int32
	_ = v1137
	var v1186 int32
	_ = v1186
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1196 int32
	_ = v1196
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1221 int32
	_ = v1221
	var v1224 int32
	_ = v1224
	var v1227 int32
	_ = v1227
	var v1232 int32
	_ = v1232
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1252 int32
	_ = v1252
	var v1255 int32
	_ = v1255
	var v1258 int32
	_ = v1258
	var v1263 int32
	_ = v1263
	var v1270 int32
	_ = v1270
	var v1273 int32
	_ = v1273
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1284 int32
	_ = v1284
	var v1288 int32
	_ = v1288
	var v1290 int32
	_ = v1290
	var v1296 int32
	_ = v1296
	var v1299 int32
	_ = v1299
	var v1301 int32
	_ = v1301
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1311 int32
	_ = v1311
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1316 int64
	_ = v1316
	var v1324 int32
	_ = v1324
	var v1337 int32
	_ = v1337
	var v1356 int32
	_ = v1356
	var v1360 int32
	_ = v1360
	var v1363 int32
	_ = v1363
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1385 int32
	_ = v1385
	var v1391 int32
	_ = v1391
	var v1408 int32
	_ = v1408
	var v1412 int32
	_ = v1412
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1421 int32
	_ = v1421
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1434 int32
	_ = v1434
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1443 int32
	_ = v1443
	var v1450 int32
	_ = v1450
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1458 int32
	_ = v1458
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1479 int32
	_ = v1479
	var v1483 int32
	_ = v1483
	var v1488 int32
	_ = v1488
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1499 int32
	_ = v1499
	var v1504 int32
	_ = v1504
	var v1533 int32
	_ = v1533
	var v1534 int32
	_ = v1534
	var v1563 int32
	_ = v1563
	var v1593 int32
	_ = v1593
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1598 int32
	_ = v1598
	var v1601 int32
	_ = v1601
	var v1604 int32
	_ = v1604
	var v1607 int32
	_ = v1607
	var v1609 int32
	_ = v1609
	var v1611 int32
	_ = v1611
	var v1612 int32
	_ = v1612
	var v1613 int32
	_ = v1613
	var v1630 int32
	_ = v1630
	var v1633 int32
	_ = v1633
	var v1653 int32
	_ = v1653
	var v1656 int32
	_ = v1656
	var v1666 int32
	_ = v1666
	var v1670 int32
	_ = v1670
	var v1671 int32
	_ = v1671
	var v1674 int32
	_ = v1674
	var v1675 int32
	_ = v1675
	var v1676 int32
	_ = v1676
	var v1682 int32
	_ = v1682
	var v1683 int32
	_ = v1683
	var v1689 int32
	_ = v1689
	var v1690 int32
	_ = v1690
	var v1692 int32
	_ = v1692
	var v1693 int32
	_ = v1693
	var v1696 int32
	_ = v1696
	var v1700 int32
	_ = v1700
	var v1705 int32
	_ = v1705
	var v1708 int32
	_ = v1708
	var v1709 int32
	_ = v1709
	var v1710 int32
	_ = v1710
	var v1717 int32
	_ = v1717
	var v1718 int32
	_ = v1718
	var v1719 int32
	_ = v1719
	var v1721 int32
	_ = v1721
	var v1722 int32
	_ = v1722
	var v1723 int32
	_ = v1723
	var v1724 int32
	_ = v1724
	var v1725 int32
	_ = v1725
	var v1727 int32
	_ = v1727
	var v1728 int32
	_ = v1728
	var v1729 int32
	_ = v1729
	var v1734 int32
	_ = v1734
	var v1736 int32
	_ = v1736
	var v1737 int32
	_ = v1737
	var v1740 int32
	_ = v1740
	var v1744 int32
	_ = v1744
	var v1759 int32
	_ = v1759
	var v1762 int32
	_ = v1762
	var v1763 int32
	_ = v1763
	var v1765 int32
	_ = v1765
	var v1766 int32
	_ = v1766
	var v1768 int32
	_ = v1768
	var v1769 int32
	_ = v1769
	var v1770 int32
	_ = v1770
	var v1771 int32
	_ = v1771
	var v1773 int32
	_ = v1773
	var v1774 int32
	_ = v1774
	var v1775 int32
	_ = v1775
	var v1777 int32
	_ = v1777
	var v1778 int32
	_ = v1778
	var v1784 int32
	_ = v1784
	var v1790 int32
	_ = v1790
	var v1807 int32
	_ = v1807
	var v1810 int32
	_ = v1810
	var v1811 int32
	_ = v1811
	var v1812 int32
	_ = v1812
	var v1814 int32
	_ = v1814
	var v1815 int32
	_ = v1815
	var v1817 int32
	_ = v1817
	var v1818 int32
	_ = v1818
	var v1824 int32
	_ = v1824
	var v1827 int32
	_ = v1827
	var v1852 int32
	_ = v1852
	var v1855 int32
	_ = v1855
	var v1856 int32
	_ = v1856
	var v1880 int32
	_ = v1880
	var v1881 int32
	_ = v1881
	var v1883 int32
	_ = v1883
	var v1884 int32
	_ = v1884
	var v1887 int32
	_ = v1887
	var v1888 int32
	_ = v1888
	var v1889 int32
	_ = v1889
	var v1890 int32
	_ = v1890
	var v1892 int32
	_ = v1892
	var v1893 int32
	_ = v1893
	var v1894 int32
	_ = v1894
	var v1896 int32
	_ = v1896
	var v1897 int32
	_ = v1897
	var v1904 int32
	_ = v1904
	var v1910 int32
	_ = v1910
	var v1927 int32
	_ = v1927
	var v1930 int32
	_ = v1930
	var v1931 int32
	_ = v1931
	var v1932 int32
	_ = v1932
	var v1934 int32
	_ = v1934
	var v1935 int32
	_ = v1935
	var v1937 int32
	_ = v1937
	var v1938 int32
	_ = v1938
	var v1945 int32
	_ = v1945
	var v1948 int32
	_ = v1948
	var v1953 int32
	_ = v1953
	var v1972 int32
	_ = v1972
	var v1973 int32
	_ = v1973
	var v1975 int32
	_ = v1975
	var v1976 int32
	_ = v1976
	var v1978 int32
	_ = v1978
	var v1979 int32
	_ = v1979
	var v1980 int32
	_ = v1980
	var v1981 int32
	_ = v1981
	var v1983 int32
	_ = v1983
	var v1984 int32
	_ = v1984
	var v1986 int32
	_ = v1986
	var v1987 int32
	_ = v1987
	var v1997 int32
	_ = v1997
	var v2000 int32
	_ = v2000
	var v2020 int32
	_ = v2020
	var v2023 int32
	_ = v2023
	var v2024 int32
	_ = v2024
	var v2025 int32
	_ = v2025
	var v2027 int32
	_ = v2027
	var v2028 int32
	_ = v2028
	var v2030 int32
	_ = v2030
	var v2031 int32
	_ = v2031
	var v2035 int32
	_ = v2035
	var v2043 int32
	_ = v2043
	var v2044 int32
	_ = v2044
	var v2067 int32
	_ = v2067
	var v2068 int32
	_ = v2068
	var v2074 int32
	_ = v2074
	var v2097 int32
	_ = v2097
	var v2098 int32
	_ = v2098
	var v2099 int32
	_ = v2099
	var v2101 int32
	_ = v2101
	var v2102 int32
	_ = v2102
	var v2130 int32
	_ = v2130
	var v2131 int32
	_ = v2131
	var v2132 int32
	_ = v2132
	var v2134 int32
	_ = v2134
	var v2135 int32
	_ = v2135
	var v2151 int32
	_ = v2151
	var v2165 int32
	_ = v2165
	var v2168 int32
	_ = v2168
	var v2169 int32
	_ = v2169
	var v2172 int32
	_ = v2172
	var v2175 int32
	_ = v2175
	var v2176 int32
	_ = v2176
	var v2177 int32
	_ = v2177
	var v2183 int32
	_ = v2183
	var v2184 int32
	_ = v2184
	var v2186 int32
	_ = v2186
	var v2188 int32
	_ = v2188
	var v2194 int32
	_ = v2194
	var v2197 int32
	_ = v2197
	var v2198 int32
	_ = v2198
	var v2217 int32
	_ = v2217
	var v2218 int32
	_ = v2218
	var v2220 int32
	_ = v2220
	var v2221 int32
	_ = v2221
	var v2223 int32
	_ = v2223
	var v2227 int32
	_ = v2227
	var v2228 int32
	_ = v2228
	var v2231 int32
	_ = v2231
	var v2232 int32
	_ = v2232
	var v2233 int32
	_ = v2233
	var v2243 int32
	_ = v2243
	var v2245 int32
	_ = v2245
	var v2264 int32
	_ = v2264
	var v2270 int32
	_ = v2270
	var v2280 int32
	_ = v2280
	var v2295 int32
	_ = v2295
	var v2299 int32
	_ = v2299
	var v2300 int32
	_ = v2300
	var v2301 int32
	_ = v2301
	var v2302 int32
	_ = v2302
	var v2303 int32
	_ = v2303
	var v2304 int32
	_ = v2304
	var v2305 int32
	_ = v2305
	var v2307 int32
	_ = v2307
	var v2308 int32
	_ = v2308
	var v2310 int32
	_ = v2310
	var v2313 int32
	_ = v2313
	var v2321 int32
	_ = v2321
	var v2324 int32
	_ = v2324
	var v2337 int32
	_ = v2337
	var v2338 int32
	_ = v2338
	var v2342 int32
	_ = v2342
	var v2343 int32
	_ = v2343
	var v2349 int32
	_ = v2349
	var v2354 int32
	_ = v2354
	var v2356 int32
	_ = v2356
	var v2357 int32
	_ = v2357
	var v2360 int32
	_ = v2360
	var v2367 int32
	_ = v2367
	var v2370 int32
	_ = v2370
	var v2371 int32
	_ = v2371
	var v2372 int32
	_ = v2372
	var v2374 int32
	_ = v2374
	var v2375 int32
	_ = v2375
	var v2387 int32
	_ = v2387
	var v2393 int32
	_ = v2393
	var v2403 int32
	_ = v2403
	var v2410 int32
	_ = v2410
	var v2411 int32
	_ = v2411
	var v2414 int32
	_ = v2414
	var v2418 int32
	_ = v2418
	var v2420 int32
	_ = v2420
	var v2426 int32
	_ = v2426
	var v2429 int32
	_ = v2429
	var v2431 int32
	_ = v2431
	var v2438 int32
	_ = v2438
	var v2439 int32
	_ = v2439
	var v2442 int32
	_ = v2442
	var v2443 int32
	_ = v2443
	var v2446 int32
	_ = v2446
	var v2450 int32
	_ = v2450
	var v2457 int32
	_ = v2457
	var v2458 int32
	_ = v2458
	var v2459 int32
	_ = v2459
	var v2461 int32
	_ = v2461
	var v2462 int32
	_ = v2462
	var v2463 int32
	_ = v2463
	var v2466 int32
	_ = v2466
	var v2474 int32
	_ = v2474
	var v2493 int32
	_ = v2493
	var v2496 int32
	_ = v2496
	var v2500 int32
	_ = v2500
	var v2502 int32
	_ = v2502
	var v2525 int32
	_ = v2525
	var v2529 int32
	_ = v2529
	var v2530 int32
	_ = v2530
	var v2531 int32
	_ = v2531
	var v2532 int32
	_ = v2532
	var v2534 int32
	_ = v2534
	var v2535 int32
	_ = v2535
	var v2538 int32
	_ = v2538
	var v2539 int32
	_ = v2539
	var v2542 int32
	_ = v2542
	var v2543 int32
	_ = v2543
	var v2550 int32
	_ = v2550
	var v2551 int32
	_ = v2551
	var v2552 int32
	_ = v2552
	var v2554 int32
	_ = v2554
	var v2555 int32
	_ = v2555
	var v2559 int32
	_ = v2559
	var v2567 int32
	_ = v2567
	v3 = int32(0)
	v28 = m.G0
	v30 = v28 - int32(288)
	m.G0 = v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+232))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)+240))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+32))
	if v35 == int32(-1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v2559 + int32(288)
	return v2567
L2:
	;
	v53 = F__emscripten_memset_bulkmem(m, v30+int32(80), base.I32_extend8_s(int32(0)), int32(128))
	mBase = m.M
	goto L7
L3:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v32)+248))
	v40 = F_predicate_refuted_by(m, v38, l1, int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	if v40 == int32(0) {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v46 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)) = uint8(v46)
	v2559 = v30
	v2567 = v3
	goto L1
L7:
	;
	if l1 == int32(0) {
		v2387 = v3
		v2393 = v3
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v2474 == int32(0) {
		goto L509
	} else {
		goto L510
	}
L9:
	;
	v2403 = int32(0)
	if v2393 == v2403 {
		goto L493
	} else {
		goto L494
	}
L10:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v56 <= int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v1255 == int32(0) {
		v1311 = v3
		goto L319
	} else {
		goto L320
	}
L12:
	;
	v1252 = v3
	v1255 = v3
	v1258 = v3
	v1263 = v3
	goto L11
L13:
	;
	goto L14
L14:
	;
	v70 = v3
	v73 = v3
	v76 = v3
	v79 = v3
	v81 = v3
	goto L15
L15:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v86+v79<<(uint(int32(2))%32))))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	if v91 == int32(318) {
		goto L24
	} else {
		goto L25
	}
L16:
	;
	v1252 = v1221
	v1255 = v1224
	v1258 = v1227
	v1263 = v1232
	goto L11
L17:
	;
	v1238 = v79 + int32(1)
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v1238 < v1239 {
		v70 = v1221
		v73 = v1224
		v76 = v1227
		v79 = v1238
		v81 = v1232
		goto L15
	} else {
		goto L317
	}
L18:
	;
	if v1134 == int32(0) {
		v1221 = v70
		v1224 = v73
		v1227 = v76
		v1232 = v81
		goto L17
	} else {
		goto L313
	}
L19:
	;
	v1186 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)) = uint8(v1186)
	v2559 = v30
	v2567 = int32(0)
	goto L1
L20:
	;
	if v1137&int32(1) == int32(0) {
		goto L18
	} else {
		goto L312
	}
L21:
	;
	v223 = int32(*(*int16)(unsafe.Add(mBase, uint32(v33)+2)))
	if v223 <= int32(0) {
		v1221 = v70
		v1224 = v73
		v1227 = v76
		v1232 = v81
		goto L17
	} else {
		goto L56
	}
L22:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	switch v107 {
	case 0:
		goto L31
	case 1:
		goto L32
	default:
		goto L21
	}
L23:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+24)))
	if v100 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v96 = v94
	v97 = v95
	goto L26
L25:
	;
	v96 = v90
	v97 = v91
	goto L26
L26:
	;
	switch v97 - int32(7) {
	case 0:
		goto L23
	default:
		goto L21
	case 14:
		goto L22
	}
L27:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v96)+20))
	if v103 != 0 {
		goto L21
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v104 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)) = uint8(v104)
	v2559 = v30
	v2567 = int32(0)
	goto L1
L30:
	;
	goto L29
L31:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
	v207 = F_gen_partprune_steps_internal(m, l0, v206)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L4
	} else {
		goto L50
	}
L32:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
	if v108 == int32(0) {
		goto L19
	} else {
		goto L33
	}
L33:
	;
	v111 = int32(0)
	v112 = int32(1)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	if v114 <= v111 {
		v1134 = v111
		v1137 = v112
		goto L20
	} else {
		goto L34
	}
L34:
	;
	v119 = v111
	v123 = v111
	v126 = v112
	goto L35
L35:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v108)+12))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v144+v119<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+72)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v30)+76)) = v148
	v154 = F_list_make1_impl(m, int32(1), v30+int32(72))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L4
	} else {
		goto L37
	}
L36:
	;
	v1134 = v199
	v1137 = v201
	goto L20
L37:
	;
	v156 = F_gen_partprune_steps_internal(m, l0, v154)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	v159 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)) = uint8(v159)
	if v158 == v159 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	if v156 != 0 {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v199 = v123
	v201 = v126
	goto L41
L41:
	;
	v203 = v119 + int32(1)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	if v203 < v204 {
		v119 = v203
		v123 = v199
		v126 = v201
		goto L35
	} else {
		goto L49
	}
L42:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
	v196 = F_lappend_int(m, v123, v195)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L4
	} else {
		goto L48
	}
L43:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v156)+12))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	v168 = int32(4)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v163+v164<<(uint(int32(2))%32)-v168)))
	v193 = v170 + v168
	goto L42
L44:
	;
	goto L45
L45:
	;
	v174 = F_palloc0(m, int32(16))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v174))) = int32(378)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v178 + int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v174)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v174)+4)) = v178
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v186 = F_lappend(m, v185, v174)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v186
	v193 = v174 + int32(4)
	goto L42
L48:
	;
	v199 = v196
	v201 = int32(0)
	goto L41
L49:
	;
	goto L36
L50:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	if v209 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v2559 = v30
	v2567 = int32(0)
	goto L1
L52:
	;
	goto L53
L53:
	;
	if v207 == int32(0) {
		v1221 = v70
		v1224 = v73
		v1227 = v76
		v1232 = v81
		goto L17
	} else {
		goto L54
	}
L54:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v207)+12))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v207)+4))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v213+v214<<(uint(int32(2))%32)-int32(4))))
	v221 = F_lappend(m, v70, v220)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	v1221 = v221
	v1224 = v73
	v1227 = v76
	v1232 = v81
	goto L17
L56:
	;
	v235 = int32(0)
	goto L60
L57:
	;
	v1115 = F_bms_is_member(m, v235, v73)
	mBase = m.M
	v1116 = m.ExcPending
	if v1116 != 0 {
		goto L4
	} else {
		goto L307
	}
L58:
	;
	v1101 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)) = uint8(v1101)
	v2559 = v30
	v2567 = int32(0)
	goto L1
L59:
	;
	v1072 = F_list_concat(m, v70, v1058)
	mBase = m.M
	v1073 = m.ExcPending
	if v1073 != 0 {
		goto L4
	} else {
		goto L306
	}
L60:
	;
	v255 = v235 << (uint(int32(2)) % 32)
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)+264))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v255+v257)))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v259)+12))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v256)+232))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v262)+12))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v263+v255)))
	v266 = int32(5)
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v262)+4))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v267+v255)))
	if base.B2i32(v269 != int32(2222))&base.B2i32(v269 != int32(424)) != 0 {
		v325 = v266
		goto L79
	} else {
		goto L80
	}
L61:
	;
	v1030 = F_bms_is_member(m, v235, v76)
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L4
	} else {
		goto L300
	}
L62:
	;
	goto L61
L63:
	;
	v1017 = v235 + int32(1)
	v1018 = int32(*(*int16)(unsafe.Add(mBase, uint32(v33)+2)))
	if v1017 < v1018 {
		v235 = v1017
		goto L60
	} else {
		goto L298
	}
L64:
	;
	if v992 == int32(0) {
		goto L63
	} else {
		goto L297
	}
L65:
	;
	v991 = int32(0)
	v992 = int32(5)
	v997 = v328
	goto L64
L66:
	;
	if v265 != 0 {
		goto L237
	} else {
		goto L238
	}
L67:
	;
	v811 = F_gen_partprune_steps_internal(m, l0, v793)
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L4
	} else {
		goto L234
	}
L68:
	;
	v764 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+20)))
	if v764 == int32(0) {
		v793 = v746
		goto L67
	} else {
		goto L227
	}
L69:
	;
	v730 = F_bms_is_member(m, v235, v73)
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L4
	} else {
		goto L222
	}
L70:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v713)))
	if v714 == int32(27) {
		goto L216
	} else {
		goto L217
	}
L71:
	;
	v681 = v666
	v682 = v666
	goto L211
L72:
	;
	if v330 == int32(52) {
		goto L70
	} else {
		goto L210
	}
L73:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v96)+28))
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v446)+12))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v447+int32(4))))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v96)+24))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v447)))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v453)))
	if v454 == int32(27) {
		goto L139
	} else {
		goto L140
	}
L74:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v96)+28))
	if v414 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L75:
	;
	v395 = F_makeBoolConst(m, v393, int32(0))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L4
	} else {
		goto L117
	}
L76:
	;
	v393 = int32(0)
	goto L75
L77:
	;
	v336 = F_makeBoolConst(m, v334, int32(0))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L4
	} else {
		goto L102
	}
L78:
	;
	v334 = int32(0)
	goto L77
L79:
	;
	v328 = int32(0)
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	switch v330 - int32(17) {
	case 0:
		goto L74
	case 1, 2:
		v991 = v328
		v992 = v325
		v997 = v328
		goto L64
	case 3:
		goto L73
	default:
		goto L72
	}
L80:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	if v275 != int32(21) {
		goto L84
	} else {
		goto L85
	}
L81:
	;
	v325 = int32(0)
	goto L79
L82:
	;
	v312 = F_equal(m, v309, v261)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L4
	} else {
		goto L95
	}
L83:
	;
	if v303 != int32(27) {
		v309 = v302
		v311 = v305
		goto L82
	} else {
		goto L94
	}
L84:
	;
	if v275 != int32(53) {
		v302 = v96
		v303 = v275
		v305 = int32(0)
		goto L83
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	if v294 != int32(2) {
		v309 = v96
		v311 = int32(0)
		goto L82
	} else {
		goto L93
	}
L87:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)))
	if v282 == int32(27) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v281)+4))
	v286 = v285
	goto L90
L89:
	;
	v286 = v281
	goto L90
L90:
	;
	v287 = F_equal(m, v286, v261)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L4
	} else {
		goto L91
	}
L91:
	;
	if v287 == int32(0) {
		goto L81
	} else {
		goto L92
	}
L92:
	;
	v291 = int32(1)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
	switch v292 {
	case 0:
		v393 = v291
		goto L75
	case 1:
		v334 = v291
		goto L77
	case 2:
		goto L76
	case 3:
		goto L78
	case 4:
		goto L62
	case 5:
		goto L69
	default:
		v325 = v266
		goto L79
	}
L93:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v297)+12))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v298)))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v299)))
	v302 = v299
	v303 = v300
	v305 = int32(1)
	goto L83
L94:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v302)+4))
	v309 = v308
	v311 = v305
	goto L82
L95:
	;
	if v312 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v393 = v311 ^ int32(1)
	goto L75
L97:
	;
	goto L98
L98:
	;
	v316 = F_negate_clause(m, v309)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L4
	} else {
		goto L99
	}
L99:
	;
	v318 = F_equal(m, v316, v261)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L4
	} else {
		goto L100
	}
L100:
	;
	if v318 != 0 {
		v393 = v311
		goto L75
	} else {
		goto L101
	}
L101:
	;
	goto L81
L102:
	;
	v339 = F_copyObjectImpl(m, v96)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L4
	} else {
		goto L106
	}
L103:
	;
	v349 = F_palloc0(m, int32(20))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L4
	} else {
		goto L107
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v339)+8)) = v345
	goto L103
L105:
	;
	v345 = int32(0)
	goto L104
L106:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v339)+8))
	switch v341 - int32(1) {
	case 0:
		v345 = int32(2)
		goto L104
	default:
		goto L103
	case 2:
		goto L105
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v349))) = int32(52)
	v353 = F_copyObjectImpl(m, v261)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L4
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v349)+16)) = int32(-1)
	v357 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v349)+12)) = uint8(v357)
	*(*int32)(unsafe.Add(mBase, uint32(v349)+8)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v349)+4)) = v353
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v349
	*(*int32)(unsafe.Add(mBase, uint32(v30)+232)) = v339
	*(*int32)(unsafe.Add(mBase, uint32(v30)+68)) = v339
	*(*int32)(unsafe.Add(mBase, uint32(v30)+64)) = v349
	v372 = F_list_make2_impl(m, v30+int32(68), v30-int32(-64))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L4
	} else {
		goto L109
	}
L109:
	;
	v375 = F_makeBoolExpr(m, int32(1), v372, int32(-1))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L4
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+60)) = v375
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v375
	v382 = F_list_make1_impl(m, int32(1), v30+int32(60))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L4
	} else {
		goto L111
	}
L111:
	;
	v384 = F_gen_partprune_steps_internal(m, l0, v382)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L4
	} else {
		goto L112
	}
L112:
	;
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	if v386 != 0 {
		goto L58
	} else {
		goto L113
	}
L113:
	;
	if v384 != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v389 = int32(3)
	goto L116
L115:
	;
	v389 = int32(5)
	goto L116
L116:
	;
	v991 = v357
	v992 = v389
	v997 = v384
	goto L64
L117:
	;
	v398 = F_palloc(m, int32(24))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L4
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v398)+12)) = v395
	v401 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v398)+8)) = uint8(v401)
	*(*int32)(unsafe.Add(mBase, uint32(v398)+4)) = int32(91)
	*(*int32)(unsafe.Add(mBase, uint32(v398))) = v235
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v262)+24))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v406+v235*int32(28))+4))
	*(*int32)(unsafe.Add(mBase, uint32(v398)+20)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v398)+16)) = v410
	v1105 = v398
	goto L57
L119:
	;
	v991 = int32(0)
	v992 = v325
	v997 = v328
	goto L64
L120:
	;
	goto L121
L121:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v414)+4))
	if v418 != int32(2) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v991 = int32(0)
	v992 = v325
	v997 = v328
	goto L64
L123:
	;
	goto L124
L124:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v414)+12))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v422)))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v423)))
	if v424 == int32(27) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v423)+4))
	v428 = v427
	goto L127
L126:
	;
	v428 = v423
	goto L127
L127:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v422)+4))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v429)))
	if v430 == int32(27) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v429)+4))
	v434 = v433
	goto L130
L129:
	;
	v434 = v429
	goto L130
L130:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v436 = F_equal(m, v428, v261)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L4
	} else {
		goto L131
	}
L131:
	;
	if v436 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v818 = v434
	v819 = v435
	goto L66
L133:
	;
	goto L134
L134:
	;
	v438 = int32(0)
	v440 = F_equal(m, v434, v261)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L4
	} else {
		goto L135
	}
L135:
	;
	if v440 == int32(0) {
		v991 = v438
		v992 = v438
		v997 = v328
		goto L64
	} else {
		goto L136
	}
L136:
	;
	v444 = F_get_commutator(m, v435)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L4
	} else {
		goto L137
	}
L137:
	;
	if v444 != 0 {
		v818 = v428
		v819 = v444
		goto L66
	} else {
		goto L138
	}
L138:
	;
	goto L65
L139:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v453)+4))
	v458 = v457
	goto L141
L140:
	;
	v458 = v453
	goto L141
L141:
	;
	v459 = F_equal(m, v458, v261)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L4
	} else {
		goto L142
	}
L142:
	;
	if v459 == int32(0) {
		goto L63
	} else {
		goto L143
	}
L143:
	;
	if v265 != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v96)+24))
	if v265 != v463 {
		goto L63
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	v465 = F_op_in_opfamily(m, v452, v269)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L4
	} else {
		goto L148
	}
L147:
	;
	goto L146
L148:
	;
	if v465 == int32(0) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262))))
	if v469 != int32(108) {
		goto L63
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	v493 = F_op_strict(m, v452)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L4
	} else {
		goto L159
	}
L152:
	;
	v472 = F_get_negator(m, v452)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L4
	} else {
		goto L153
	}
L153:
	;
	if v472 == int32(0) {
		goto L63
	} else {
		goto L154
	}
L154:
	;
	v476 = F_op_in_opfamily(m, v472, v269)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L4
	} else {
		goto L155
	}
L155:
	;
	if v476 == int32(0) {
		goto L63
	} else {
		goto L156
	}
L156:
	;
	F_get_op_opfamily_properties(m, v472, v269, int32(0), v30+int32(240), v30+int32(276), v30+int32(228))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L4
	} else {
		goto L157
	}
L157:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v30)+240))
	if v489 != int32(3) {
		goto L63
	} else {
		goto L158
	}
L158:
	;
	goto L151
L159:
	;
	if v493 == int32(0) {
		v1221 = v70
		v1224 = v73
		v1227 = v76
		v1232 = v81
		goto L17
	} else {
		goto L160
	}
L160:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v450)))
	if v497 == int32(7) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v535 = F_op_volatile(m, v452)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L4
	} else {
		goto L178
	}
L162:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v500 == int32(0) {
		v1221 = v70
		v1224 = v73
		v1227 = v76
		v1232 = v81
		goto L17
	} else {
		goto L163
	}
L163:
	;
	v503 = F_contain_var_clause(m, v450)
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L4
	} else {
		goto L164
	}
L164:
	;
	if v503 != 0 {
		v1221 = v70
		v1224 = v73
		v1227 = v76
		v1232 = v81
		goto L17
	} else {
		goto L165
	}
L165:
	;
	v505 = F_contain_volatile_functions(m, v450)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L4
	} else {
		goto L166
	}
L166:
	;
	if v505 != 0 {
		v1221 = v70
		v1224 = v73
		v1227 = v76
		v1232 = v81
		goto L17
	} else {
		goto L167
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+240)) = int32(0)
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v450)))
	if v509 == int32(8) {
		goto L170
	} else {
		goto L171
	}
L168:
	;
	v533 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)) = uint8(v533)
	goto L161
L169:
	;
	if v525 == int32(0) {
		goto L168
	} else {
		goto L176
	}
L170:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v450)+4))
	if v512 != int32(1) {
		goto L168
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	v522 = F_expression_tree_walker_impl(m, v450, int32(909), v30+int32(240))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L4
	} else {
		goto L175
	}
L173:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v450)+8))
	v517 = F_bms_add_member(m, int32(0), v516)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L4
	} else {
		goto L174
	}
L174:
	;
	v525 = v517
	goto L169
L175:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v30)+240))
	v525 = v524
	goto L169
L176:
	;
	v528 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)) = uint8(v528)
	v530 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v530 == int32(2) {
		goto L161
	} else {
		goto L177
	}
L177:
	;
	v1221 = v70
	v1224 = v73
	v1227 = v76
	v1232 = v81
	goto L17
L178:
	;
	if v535 != int32(105) {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v539 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)) = uint8(v539)
	v541 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v541 == int32(0) {
		v1221 = v70
		v1224 = v73
		v1227 = v76
		v1232 = v81
		goto L17
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v450)))
	if v544 != int32(35) {
		goto L184
	} else {
		goto L185
	}
L182:
	;
	goto L181
L183:
	;
	if v643 == int32(0) {
		goto L206
	} else {
		goto L207
	}
L184:
	;
	if v544 != int32(7) {
		v1221 = v70
		v1224 = v73
		v1227 = v76
		v1232 = v81
		goto L17
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v450)+20)))
	if v634 != 0 {
		v1221 = v70
		v1224 = v73
		v1227 = v76
		v1232 = v81
		goto L17
	} else {
		goto L205
	}
L187:
	;
	v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v450)+24)))
	if v549 != 0 {
		goto L58
	} else {
		goto L188
	}
L188:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v450)+20))
	v551 = F_pg_detoast_datum(m, v550)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L4
	} else {
		goto L189
	}
L189:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v551)+12))
	F_get_typlenbyvalalign(m, v553, v30+int32(224), v30+int32(223), v30+int32(222))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L4
	} else {
		goto L190
	}
L190:
	;
	v563 = int32(*(*int16)(unsafe.Add(mBase, uint32(v30)+224)))
	v564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+223)))
	v565 = int32(*(*int8)(unsafe.Add(mBase, uint32(v30)+222)))
	F_deconstruct_array(m, v551, v563, v564, v565, v30+int32(240), v30+int32(276), v30+int32(228))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L4
	} else {
		goto L191
	}
L191:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v30)+228))
	if v574 <= int32(0) {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v643 = int32(0)
	goto L183
L193:
	;
	goto L194
L194:
	;
	v578 = int32(0)
	v587 = v578
	v588 = v578
	v589 = v574
	goto L195
L195:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v30)+276))
	v609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v607+v588))))
	if v609 == int32(1) {
		goto L198
	} else {
		goto L199
	}
L196:
	;
	v643 = v629
	goto L183
L197:
	;
	v632 = v588 + int32(1)
	if v632 < v630 {
		v587 = v629
		v588 = v632
		v589 = v630
		goto L195
	} else {
		goto L204
	}
L198:
	;
	v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+20)))
	if v612 != 0 {
		v629 = v587
		v630 = v589
		goto L197
	} else {
		goto L201
	}
L199:
	;
	goto L200
L200:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v551)+12))
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v450)+12))
	v616 = int32(*(*int16)(unsafe.Add(mBase, uint32(v30)+224)))
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v30)+240))
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v617+v588<<(uint(int32(2))%32))))
	v623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+223)))
	v624 = F_makeConst(m, v613, int32(-1), v615, v616, v621, int32(0), v623)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L4
	} else {
		goto L202
	}
L201:
	;
	goto L58
L202:
	;
	v626 = F_lappend(m, v587, v624)
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L4
	} else {
		goto L203
	}
L203:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v30)+228))
	v629 = v626
	v630 = v628
	goto L197
L204:
	;
	goto L196
L205:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v450)+16))
	v643 = v635
	goto L183
L206:
	;
	v793 = int32(0)
	goto L67
L207:
	;
	goto L208
L208:
	;
	v666 = int32(0)
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v643)+4))
	if v666 < v668 {
		goto L71
	} else {
		goto L209
	}
L209:
	;
	v746 = v666
	goto L68
L210:
	;
	v991 = v328
	v992 = v325
	v997 = v328
	goto L64
L211:
	;
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v643)+12))
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v700+v681<<(uint(int32(2))%32))))
	v705 = F_make_opclause(m, v452, v458, v704, v451)
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L4
	} else {
		goto L213
	}
L212:
	;
	v746 = v707
	goto L68
L213:
	;
	v707 = F_lappend(m, v682, v705)
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L4
	} else {
		goto L214
	}
L214:
	;
	v710 = v681 + int32(1)
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v643)+4))
	if v710 < v711 {
		v681 = v710
		v682 = v707
		goto L211
	} else {
		goto L215
	}
L215:
	;
	goto L212
L216:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v713)+4))
	v718 = v717
	goto L218
L217:
	;
	v718 = v713
	goto L218
L218:
	;
	v719 = F_equal(m, v718, v261)
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L4
	} else {
		goto L219
	}
L219:
	;
	if v719 == int32(0) {
		goto L63
	} else {
		goto L220
	}
L220:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
	if v723 != int32(1) {
		goto L62
	} else {
		goto L221
	}
L221:
	;
	goto L69
L222:
	;
	if v730 != 0 {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v732 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)) = uint8(v732)
	v2559 = v30
	v2567 = int32(0)
	goto L1
L224:
	;
	goto L225
L225:
	;
	v735 = F_bms_add_member(m, v76, v235)
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L4
	} else {
		goto L226
	}
L226:
	;
	v1221 = v70
	v1224 = v73
	v1227 = v735
	v1232 = v81
	goto L17
L227:
	;
	if v746 == int32(0) {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v793 = int32(0)
	goto L67
L229:
	;
	goto L230
L230:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v746)+4))
	if v770 < int32(2) {
		v793 = v746
		goto L67
	} else {
		goto L231
	}
L231:
	;
	v775 = F_makeBoolExpr(m, int32(1), v746, int32(-1))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L4
	} else {
		goto L232
	}
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+56)) = v775
	*(*int32)(unsafe.Add(mBase, uint32(v30)+216)) = v775
	v782 = F_list_make1_impl(m, int32(1), v30+int32(56))
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L4
	} else {
		goto L233
	}
L233:
	;
	v793 = v782
	goto L67
L234:
	;
	v813 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	if v813 != 0 {
		goto L58
	} else {
		goto L235
	}
L235:
	;
	if v811 == int32(0) {
		v1221 = v70
		v1224 = v73
		v1227 = v76
		v1232 = v81
		goto L17
	} else {
		goto L236
	}
L236:
	;
	v1058 = v811
	goto L59
L237:
	;
	v820 = int32(0)
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v96)+24))
	if v265 != v822 {
		v991 = v820
		v992 = v820
		v997 = v328
		goto L64
	} else {
		goto L240
	}
L238:
	;
	goto L239
L239:
	;
	v826 = F_op_in_opfamily(m, v819, v269)
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L4
	} else {
		goto L243
	}
L240:
	;
	goto L239
L241:
	;
	v982 = int32(0)
	v991 = v982
	v992 = v982
	v997 = v328
	goto L64
L242:
	;
	v865 = int32(5)
	v866 = int32(0)
	v867 = F_op_strict(m, v819)
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L4
	} else {
		goto L255
	}
L243:
	;
	if v826 != 0 {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	F_get_op_opfamily_properties(m, v819, v269, int32(0), v30+int32(224), v30+int32(276), v30+int32(228))
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L4
	} else {
		goto L247
	}
L245:
	;
	goto L246
L246:
	;
	v837 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262))))
	if v837 != int32(108) {
		goto L65
	} else {
		goto L248
	}
L247:
	;
	v864 = v819
	goto L242
L248:
	;
	v840 = F_get_negator(m, v819)
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L4
	} else {
		goto L249
	}
L249:
	;
	if v840 == int32(0) {
		goto L241
	} else {
		goto L250
	}
L250:
	;
	v844 = int32(0)
	v846 = F_op_in_opfamily(m, v840, v269)
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L4
	} else {
		goto L251
	}
L251:
	;
	if v846 == int32(0) {
		v991 = v844
		v992 = v844
		v997 = v328
		goto L64
	} else {
		goto L252
	}
L252:
	;
	F_get_op_opfamily_properties(m, v840, v269, int32(0), v30+int32(224), v30+int32(276), v30+int32(228))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L4
	} else {
		goto L253
	}
L253:
	;
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v30)+224))
	if v859 != int32(3) {
		v991 = v844
		v992 = v844
		v997 = v328
		goto L64
	} else {
		goto L254
	}
L254:
	;
	v864 = v840
	goto L242
L255:
	;
	if v867 == int32(0) {
		v991 = v866
		v992 = v865
		v997 = v328
		goto L64
	} else {
		goto L256
	}
L256:
	;
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v818)))
	if v871 == int32(7) {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	v909 = F_op_volatile(m, v819)
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L4
	} else {
		goto L274
	}
L258:
	;
	v874 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v874 == int32(0) {
		v991 = v866
		v992 = v865
		v997 = v328
		goto L64
	} else {
		goto L259
	}
L259:
	;
	v877 = F_contain_var_clause(m, v818)
	mBase = m.M
	v878 = m.ExcPending
	if v878 != 0 {
		goto L4
	} else {
		goto L260
	}
L260:
	;
	if v877 != 0 {
		v991 = v866
		v992 = v865
		v997 = v328
		goto L64
	} else {
		goto L261
	}
L261:
	;
	v879 = F_contain_volatile_functions(m, v818)
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L4
	} else {
		goto L262
	}
L262:
	;
	if v879 != 0 {
		v991 = v866
		v992 = v865
		v997 = v328
		goto L64
	} else {
		goto L263
	}
L263:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+240)) = int32(0)
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v818)))
	if v883 == int32(8) {
		goto L266
	} else {
		goto L267
	}
L264:
	;
	v907 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)) = uint8(v907)
	goto L257
L265:
	;
	if v899 == int32(0) {
		goto L264
	} else {
		goto L272
	}
L266:
	;
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v818)+4))
	if v886 != int32(1) {
		goto L264
	} else {
		goto L269
	}
L267:
	;
	goto L268
L268:
	;
	v896 = F_expression_tree_walker_impl(m, v818, int32(909), v30+int32(240))
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L4
	} else {
		goto L271
	}
L269:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v818)+8))
	v891 = F_bms_add_member(m, int32(0), v890)
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L4
	} else {
		goto L270
	}
L270:
	;
	v899 = v891
	goto L265
L271:
	;
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v30)+240))
	v899 = v898
	goto L265
L272:
	;
	v902 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)) = uint8(v902)
	v904 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v904 == int32(2) {
		goto L257
	} else {
		goto L273
	}
L273:
	;
	v991 = v866
	v992 = v865
	v997 = v328
	goto L64
L274:
	;
	if v909 != int32(105) {
		goto L275
	} else {
		goto L276
	}
L275:
	;
	v913 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)) = uint8(v913)
	v915 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v915 == int32(0) {
		v991 = v866
		v992 = v865
		v997 = v328
		goto L64
	} else {
		goto L278
	}
L276:
	;
	goto L277
L277:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v30)+228))
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v262)+8))
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v919+v255)))
	if v918 == v921 {
		goto L280
	} else {
		goto L281
	}
L278:
	;
	goto L277
L279:
	;
	v964 = F_palloc(m, int32(24))
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L4
	} else {
		goto L293
	}
L280:
	;
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v262)+24))
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v923+v235*int32(28))+4))
	v962 = v927
	goto L279
L281:
	;
	goto L282
L282:
	;
	v928 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262))))
	switch v928 - int32(104) {
	case 0:
		goto L286
	default:
		goto L285
	case 4, 10:
		goto L284
	}
L283:
	;
	if v959 == int32(0) {
		goto L241
	} else {
		goto L292
	}
L284:
	;
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v262)+4))
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v953+v255)))
	v957 = F_get_opfamily_proc(m, v955, v921, v918, int32(1))
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		goto L4
	} else {
		goto L291
	}
L285:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v940 = m.ExcPending
	if v940 != 0 {
		goto L4
	} else {
		goto L288
	}
L286:
	;
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v262)+4))
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v931+v255)))
	v935 = F_get_opfamily_proc(m, v933, v918, v918, int32(2))
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L4
	} else {
		goto L287
	}
L287:
	;
	v959 = v935
	goto L283
L288:
	;
	v941 = int32(*(*int8)(unsafe.Add(mBase, uint32(v262))))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+48)) = v941
	F_errmsg_internal(m, int32(_a_F_gen_partprune_steps_internal_0), v30+int32(48))
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L4
	} else {
		goto L289
	}
L289:
	;
	F_errfinish(m, int32(_a_F_gen_partprune_steps_internal_1), int32(2138), int32(_a_F_gen_partprune_steps_internal_2))
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L4
	} else {
		goto L290
	}
L290:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L291:
	;
	v959 = v957
	goto L283
L292:
	;
	v962 = v959
	goto L279
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v964))) = v235
	v967 = int32(1)
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v30)+224))
	v970 = v826 ^ v967
	*(*uint8)(unsafe.Add(mBase, uint32(v964)+8)) = uint8(v970)
	*(*int32)(unsafe.Add(mBase, uint32(v964)+4)) = v864
	*(*int32)(unsafe.Add(mBase, uint32(v964)+16)) = v962
	*(*int32)(unsafe.Add(mBase, uint32(v964)+12)) = v818
	if v826 != 0 {
		goto L294
	} else {
		goto L295
	}
L294:
	;
	v976 = v968
	goto L296
L295:
	;
	v976 = int32(0)
	goto L296
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v964)+20)) = v976
	v991 = v964
	v992 = v967
	v997 = v328
	goto L64
L297:
	;
	switch v992 - int32(1) {
	case 0:
		v1105 = v991
		goto L57
	case 1:
		goto L62
	case 2:
		v1058 = v997
		goto L59
	case 3:
		goto L58
	default:
		v1221 = v70
		v1224 = v73
		v1227 = v76
		v1232 = v81
		goto L17
	}
L298:
	;
	v1221 = v70
	v1224 = v73
	v1227 = v76
	v1232 = v81
	goto L17
L299:
	;
	v1043 = F_bms_add_member(m, v73, v235)
	mBase = m.M
	v1044 = m.ExcPending
	if v1044 != 0 {
		goto L4
	} else {
		goto L305
	}
L300:
	;
	if v1030 == int32(0) {
		goto L301
	} else {
		goto L302
	}
L301:
	;
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(v30+int32(80)+v255)))
	if v1037 == int32(0) {
		goto L299
	} else {
		goto L304
	}
L302:
	;
	goto L303
L303:
	;
	v1040 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)) = uint8(v1040)
	v2559 = v30
	v2567 = int32(0)
	goto L1
L304:
	;
	goto L303
L305:
	;
	v1221 = v70
	v1224 = v1043
	v1227 = v76
	v1232 = v81
	goto L17
L306:
	;
	v1221 = v1072
	v1224 = v73
	v1227 = v76
	v1232 = v81
	goto L17
L307:
	;
	if v1115 != 0 {
		goto L308
	} else {
		goto L309
	}
L308:
	;
	v1117 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)) = uint8(v1117)
	v2559 = v30
	v2567 = int32(0)
	goto L1
L309:
	;
	goto L310
L310:
	;
	v1122 = v30 + int32(80) + v255
	v1123 = *(*int32)(unsafe.Add(mBase, uint32(v1122)))
	v1124 = F_lappend(m, v1123, v1105)
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		goto L4
	} else {
		goto L311
	}
L311:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1122))) = v1124
	v1221 = v70
	v1224 = v73
	v1227 = v76
	v1232 = int32(1)
	goto L17
L312:
	;
	goto L19
L313:
	;
	v1192 = F_palloc0(m, int32(16))
	mBase = m.M
	v1193 = m.ExcPending
	if v1193 != 0 {
		goto L4
	} else {
		goto L314
	}
L314:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1192))) = int32(378)
	v1196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1196 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1192)+12)) = v1134
	*(*int32)(unsafe.Add(mBase, uint32(v1192)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1192)+4)) = v1196
	v1204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1205 = F_lappend(m, v1204, v1192)
	mBase = m.M
	v1206 = m.ExcPending
	if v1206 != 0 {
		goto L4
	} else {
		goto L315
	}
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1205
	v1208 = F_lappend(m, v70, v1192)
	mBase = m.M
	v1209 = m.ExcPending
	if v1209 != 0 {
		goto L4
	} else {
		goto L316
	}
L316:
	;
	v1221 = v1208
	v1224 = v73
	v1227 = v76
	v1232 = v81
	goto L17
L317:
	;
	goto L16
L318:
	;
	v2356 = F_palloc0(m, int32(24))
	mBase = m.M
	v2357 = m.ExcPending
	if v2357 != 0 {
		goto L4
	} else {
		goto L489
	}
L319:
	;
	if v1263 == int32(0) {
		v2387 = v1252
		v2393 = v1258
		goto L9
	} else {
		goto L336
	}
L320:
	;
	v1270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	switch v1270 - int32(104) {
	case 0:
		goto L321
	default:
		v1311 = v1255
		goto L319
	case 4, 10:
		goto L318
	}
L321:
	;
	v1273 = int32(0)
	if v1255 == v1273 {
		goto L323
	} else {
		goto L324
	}
L322:
	;
	v1309 = int32(*(*int16)(unsafe.Add(mBase, uint32(v33)+2)))
	if v1308 == v1309 {
		goto L318
	} else {
		goto L335
	}
L323:
	;
	v1308 = int32(0)
	goto L322
L324:
	;
	goto L325
L325:
	;
	v1280 = int32(1)
	v1281 = *(*int32)(unsafe.Add(mBase, uint32(v1255)+4))
	if v1281 <= v1280 {
		goto L326
	} else {
		goto L327
	}
L326:
	;
	v1284 = v1280
	goto L328
L327:
	;
	v1284 = v1281
	goto L328
L328:
	;
	v1288 = int32(0)
	v1290 = v1273
	goto L329
L329:
	;
	v1296 = *(*int32)(unsafe.Add(mBase, uint32(v1255+int32(8)+v1288<<(uint(int32(2))%32))))
	if v1296 != 0 {
		goto L331
	} else {
		goto L332
	}
L330:
	;
	v1308 = v1299
	goto L322
L331:
	;
	v1299 = v1290 + base.I32_popcnt(v1296)
	goto L333
L332:
	;
	v1299 = v1290
	goto L333
L333:
	;
	v1301 = v1288 + int32(1)
	if v1301 != v1284 {
		v1288 = v1301
		v1290 = v1299
		goto L329
	} else {
		goto L334
	}
L334:
	;
	goto L330
L335:
	;
	v1311 = v1255
	goto L319
L336:
	;
	v1314 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(v1314)+232))
	v1316 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v30)+256)) = v1316
	*(*int64)(unsafe.Add(mBase, uint32(v30)+240)) = v1316
	*(*int64)(unsafe.Add(mBase, uint32(v30)+248)) = v1316
	*(*int64)(unsafe.Add(mBase, uint32(v30)+232)) = v1316
	v1324 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1315)+2)))
	if v1324 <= int32(0) {
		goto L341
	} else {
		goto L342
	}
L337:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2342 = m.ExcPending
	if v2342 != 0 {
		goto L4
	} else {
		goto L486
	}
L338:
	;
	v2337 = F_list_concat(m, v2321, v2324)
	mBase = m.M
	v2338 = m.ExcPending
	if v2338 != 0 {
		goto L4
	} else {
		goto L485
	}
L339:
	;
	v2168 = int32(0)
	v2169 = *(*int32)(unsafe.Add(mBase, uint32(v30)+236))
	if v2169 == v2168 {
		v2310 = l0
		v2313 = v30
		v2321 = v1252
		v2324 = v2168
		goto L338
	} else {
		goto L465
	}
L340:
	;
	v1593 = *(*int32)(unsafe.Add(mBase, uint32(v30)+256))
	v1594 = *(*int32)(unsafe.Add(mBase, uint32(v30)+248))
	v1595 = *(*int32)(unsafe.Add(mBase, uint32(v30)+252))
	v1598 = l0
	v1601 = v30
	v1604 = v1595
	v1607 = v1593
	v1609 = v1252
	v1611 = v1594
	v1612 = int32(0)
	v1613 = int32(1)
	goto L380
L341:
	;
	v1563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1315))))
	switch v1563 - int32(104) {
	case 0:
		goto L339
	default:
		goto L337
	case 4, 10:
		goto L340
	}
L342:
	;
	v1337 = int32(0)
	goto L343
L343:
	;
	v1356 = v1337 << (uint(int32(2)) % 32)
	v1360 = *(*int32)(unsafe.Add(mBase, uint32(v1356+(v30+int32(80)))))
	v1363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1315))))
	if base.B2i32(v1360 == int32(0))&base.B2i32(v1363 == int32(114)) != 0 {
		goto L340
	} else {
		goto L345
	}
L344:
	;
	goto L341
L345:
	;
	if v1363 != int32(104) {
		goto L347
	} else {
		goto L348
	}
L346:
	;
	v1533 = v1337 + int32(1)
	v1534 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1315)+2)))
	if v1533 < v1534 {
		v1337 = v1533
		goto L343
	} else {
		goto L379
	}
L347:
	;
	if v1360 == int32(0) {
		goto L346
	} else {
		goto L353
	}
L348:
	;
	if v1360 != 0 {
		goto L347
	} else {
		goto L349
	}
L349:
	;
	v1369 = F_bms_is_member(m, v1337, v1311)
	mBase = m.M
	v1370 = m.ExcPending
	if v1370 != 0 {
		goto L4
	} else {
		goto L350
	}
L350:
	;
	if v1369 != 0 {
		goto L346
	} else {
		goto L351
	}
L351:
	;
	v1372 = F_list_concat(m, v1252, int32(0))
	mBase = m.M
	v1373 = m.ExcPending
	if v1373 != 0 {
		goto L4
	} else {
		goto L352
	}
L352:
	;
	v2463 = l0
	v2466 = v30
	v2474 = v1372
	goto L8
L353:
	;
	v1377 = int32(0)
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(v1360)+4))
	if v1378 <= v1377 {
		goto L346
	} else {
		goto L354
	}
L354:
	;
	v1385 = v1377
	v1391 = int32(1)
	goto L356
L355:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1492 = m.ExcPending
	if v1492 != 0 {
		goto L4
	} else {
		goto L376
	}
L356:
	;
	v1408 = *(*int32)(unsafe.Add(mBase, uint32(v1360)+12))
	v1412 = *(*int32)(unsafe.Add(mBase, uint32(v1408+v1385<<(uint(int32(2))%32))))
	v1414 = v1412 + int32(20)
	v1415 = *(*int32)(unsafe.Add(mBase, uint32(v1412)+20))
	if v1415 == int32(0) {
		goto L358
	} else {
		goto L359
	}
L357:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1479 = m.ExcPending
	if v1479 != 0 {
		goto L4
	} else {
		goto L373
	}
L358:
	;
	v1418 = *(*int32)(unsafe.Add(mBase, uint32(v1412)+4))
	v1419 = *(*int32)(unsafe.Add(mBase, uint32(v1315)+4))
	v1421 = *(*int32)(unsafe.Add(mBase, uint32(v1419+v1356)))
	F_get_op_opfamily_properties(m, v1418, v1421, int32(0), v1414, v30+int32(284), v30+int32(280))
	mBase = m.M
	v1428 = m.ExcPending
	if v1428 != 0 {
		goto L4
	} else {
		goto L361
	}
L359:
	;
	goto L360
L360:
	;
	v1429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1315))))
	switch v1429 - int32(104) {
	case 0:
		goto L365
	default:
		goto L355
	case 4, 10:
		goto L366
	}
L361:
	;
	goto L360
L362:
	;
	goto L357
L363:
	;
	v1473 = v1385 + int32(1)
	v1474 = *(*int32)(unsafe.Add(mBase, uint32(v1360)+4))
	if v1473 < v1474 {
		v1385 = v1473
		v1391 = int32(0)
		goto L356
	} else {
		goto L372
	}
L364:
	;
	v1466 = v1385 + int32(1)
	v1467 = *(*int32)(unsafe.Add(mBase, uint32(v1360)+4))
	if v1466 < v1467 {
		v1385 = v1466
		goto L356
	} else {
		goto L370
	}
L365:
	;
	v1450 = *(*int32)(unsafe.Add(mBase, uint32(v1414)))
	if v1450 != int32(1) {
		goto L362
	} else {
		goto L368
	}
L366:
	;
	v1434 = *(*int32)(unsafe.Add(mBase, uint32(v1412)+20))
	v1438 = *(*int32)(unsafe.Add(mBase, uint32(v30+int32(240)+v1434<<(uint(int32(2))%32))))
	v1439 = F_lappend(m, v1438, v1412)
	mBase = m.M
	v1440 = m.ExcPending
	if v1440 != 0 {
		goto L4
	} else {
		goto L367
	}
L367:
	;
	v1443 = *(*int32)(unsafe.Add(mBase, uint32(v1412)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v30+int32(240)+v1443<<(uint(int32(2))%32)))) = v1439
	switch v1443 - int32(1) {
	case 0, 4:
		goto L363
	default:
		goto L364
	}
L368:
	;
	v1453 = *(*int32)(unsafe.Add(mBase, uint32(v30)+236))
	v1454 = F_lappend(m, v1453, v1412)
	mBase = m.M
	v1455 = m.ExcPending
	if v1455 != 0 {
		goto L4
	} else {
		goto L369
	}
L369:
	;
	v1458 = *(*int32)(unsafe.Add(mBase, uint32(v1412)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v30+int32(232)+v1458<<(uint(int32(2))%32)))) = v1454
	goto L364
L370:
	;
	if v1391&int32(1) != 0 {
		goto L346
	} else {
		goto L371
	}
L371:
	;
	goto L341
L372:
	;
	goto L341
L373:
	;
	F_errmsg_internal(m, int32(_a_F_gen_partprune_steps_internal_3), int32(0))
	mBase = m.M
	v1483 = m.ExcPending
	if v1483 != 0 {
		goto L4
	} else {
		goto L374
	}
L374:
	;
	F_errfinish(m, int32(_a_F_gen_partprune_steps_internal_1), int32(1480), int32(_a_F_gen_partprune_steps_internal_4))
	mBase = m.M
	v1488 = m.ExcPending
	if v1488 != 0 {
		goto L4
	} else {
		goto L375
	}
L375:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L376:
	;
	v1493 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1315))))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+32)) = v1493
	F_errmsg_internal(m, int32(_a_F_gen_partprune_steps_internal_0), v30+int32(32))
	mBase = m.M
	v1499 = m.ExcPending
	if v1499 != 0 {
		goto L4
	} else {
		goto L377
	}
L377:
	;
	F_errfinish(m, int32(_a_F_gen_partprune_steps_internal_1), int32(1487), int32(_a_F_gen_partprune_steps_internal_4))
	mBase = m.M
	v1504 = m.ExcPending
	if v1504 != 0 {
		goto L4
	} else {
		goto L378
	}
L378:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L379:
	;
	goto L344
L380:
	;
	v1630 = *(*int32)(unsafe.Add(mBase, uint32(v1601+int32(240)+v1613<<(uint(int32(2))%32))))
	if v1630 == int32(0) {
		v2151 = v1612
		goto L382
	} else {
		goto L383
	}
L381:
	;
	v2310 = v1598
	v2313 = v1601
	v2321 = v1609
	v2324 = v2151
	goto L338
L382:
	;
	v2165 = v1613 + int32(1)
	if v2165 != int32(6) {
		v1612 = v2151
		v1613 = v2165
		goto L380
	} else {
		goto L464
	}
L383:
	;
	v1633 = *(*int32)(unsafe.Add(mBase, uint32(v1630)+4))
	if v1633 <= int32(0) {
		v2151 = v1612
		goto L382
	} else {
		goto L384
	}
L384:
	;
	v1653 = v1612
	v1656 = int32(0)
	goto L385
L385:
	;
	v1666 = *(*int32)(unsafe.Add(mBase, uint32(v1630)+12))
	v1670 = *(*int32)(unsafe.Add(mBase, uint32(v1666+v1656<<(uint(int32(2))%32))))
	v1671 = *(*int32)(unsafe.Add(mBase, uint32(v1670)))
	if v1671 == int32(0) {
		goto L388
	} else {
		goto L389
	}
L386:
	;
	v2151 = v2131
	goto L382
L387:
	;
	v2131 = F_list_concat(m, v1653, v2130)
	mBase = m.M
	v2132 = m.ExcPending
	if v2132 != 0 {
		goto L4
	} else {
		goto L462
	}
L388:
	;
	v1674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1670)+8)))
	v1675 = *(*int32)(unsafe.Add(mBase, uint32(v1670)+16))
	v1676 = *(*int32)(unsafe.Add(mBase, uint32(v1670)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1601)+12)) = v1676
	*(*int32)(unsafe.Add(mBase, uint32(v1601)+284)) = v1676
	v1682 = F_list_make1_impl(m, int32(1), v1601+int32(12))
	mBase = m.M
	v1683 = m.ExcPending
	if v1683 != 0 {
		goto L4
	} else {
		goto L391
	}
L389:
	;
	goto L390
L390:
	;
	v1719 = int32(0)
	if v1604 != 0 {
		goto L399
	} else {
		goto L400
	}
L391:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1601)+8)) = v1675
	*(*int32)(unsafe.Add(mBase, uint32(v1601)+280)) = v1675
	v1689 = F_list_make1_impl(m, int32(472), v1601+int32(8))
	mBase = m.M
	v1690 = m.ExcPending
	if v1690 != 0 {
		goto L4
	} else {
		goto L392
	}
L392:
	;
	v1692 = F_palloc0(m, int32(24))
	mBase = m.M
	v1693 = m.ExcPending
	if v1693 != 0 {
		goto L4
	} else {
		goto L393
	}
L393:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1692))) = int32(377)
	v1696 = *(*int32)(unsafe.Add(mBase, uint32(v1598)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1598)+16)) = v1696 + int32(1)
	v1700 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1692)+20)) = v1700
	*(*int32)(unsafe.Add(mBase, uint32(v1692)+16)) = v1689
	*(*int32)(unsafe.Add(mBase, uint32(v1692)+12)) = v1682
	if v1674 != 0 {
		goto L394
	} else {
		goto L395
	}
L394:
	;
	v1705 = v1700
	goto L396
L395:
	;
	v1705 = v1613
	goto L396
L396:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1692)+8)) = uint16(v1705)
	*(*int32)(unsafe.Add(mBase, uint32(v1692)+4)) = v1696
	v1708 = *(*int32)(unsafe.Add(mBase, uint32(v1598)+8))
	v1709 = F_lappend(m, v1708, v1692)
	mBase = m.M
	v1710 = m.ExcPending
	if v1710 != 0 {
		goto L4
	} else {
		goto L397
	}
L397:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1598)+8)) = v1709
	*(*int32)(unsafe.Add(mBase, uint32(v1601)+4)) = v1692
	*(*int32)(unsafe.Add(mBase, uint32(v1601)+276)) = v1692
	v1717 = F_list_make1_impl(m, int32(1), v1601+int32(4))
	mBase = m.M
	v1718 = m.ExcPending
	if v1718 != 0 {
		goto L4
	} else {
		goto L398
	}
L398:
	;
	v2130 = v1717
	goto L387
L399:
	;
	v1721 = *(*int32)(unsafe.Add(mBase, uint32(v1604)+12))
	v1722 = v1721
	goto L401
L400:
	;
	v1722 = v1719
	goto L401
L401:
	;
	if v1611 != 0 {
		goto L402
	} else {
		goto L403
	}
L402:
	;
	v1723 = *(*int32)(unsafe.Add(mBase, uint32(v1611)+12))
	v1724 = v1723
	goto L404
L403:
	;
	v1724 = v1719
	goto L404
L404:
	;
	v1725 = int32(0)
	if v1607 != 0 {
		goto L405
	} else {
		goto L406
	}
L405:
	;
	v1727 = *(*int32)(unsafe.Add(mBase, uint32(v1607)+12))
	v1728 = v1727
	goto L407
L406:
	;
	v1728 = v1725
	goto L407
L407:
	;
	v1729 = int32(0)
	if v1729 < v1671 {
		goto L408
	} else {
		goto L409
	}
L408:
	;
	v1734 = v1725
	v1736 = v1729
	v1737 = v1728
	v1740 = v1722
	v1744 = v1724
	goto L411
L409:
	;
	v2074 = v1729
	goto L410
L410:
	;
	v2097 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1670)+8)))
	v2098 = *(*int32)(unsafe.Add(mBase, uint32(v1670)+12))
	v2099 = *(*int32)(unsafe.Add(mBase, uint32(v1670)+16))
	v2101 = F_get_steps_using_prefix(m, v1598, v1613, v2097, v2098, v2099, int32(0), v2074)
	mBase = m.M
	v2102 = m.ExcPending
	if v2102 != 0 {
		goto L4
	} else {
		goto L461
	}
L411:
	;
	v1759 = int32(0)
	if v1740 == v1759 {
		v1824 = v1736
		v1827 = v1759
		goto L414
	} else {
		goto L415
	}
L412:
	;
	v2074 = v2043
	goto L410
L413:
	;
	if base.Ui32(int32(2)) < base.Ui32(v1613) {
		v1945 = v1852
		v1948 = v1855
		v1953 = v1744
		goto L425
	} else {
		goto L426
	}
L414:
	;
	v1852 = v1824
	v1855 = v1827
	v1856 = int32(0)
	goto L413
L415:
	;
	v1762 = *(*int32)(unsafe.Add(mBase, uint32(v1604)+12))
	v1763 = v1740 - v1762
	v1765 = v1763 >> (uint(int32(2)) % 32)
	v1766 = *(*int32)(unsafe.Add(mBase, uint32(v1604)+4))
	if v1766 <= v1765 {
		v1824 = v1736
		v1827 = v1759
		goto L414
	} else {
		goto L416
	}
L416:
	;
	v1768 = *(*int32)(unsafe.Add(mBase, uint32(v1604)+12))
	v1769 = v1768 + v1763
	v1770 = *(*int32)(unsafe.Add(mBase, uint32(v1769)))
	v1771 = *(*int32)(unsafe.Add(mBase, uint32(v1770)))
	if v1771 != v1734 {
		v1852 = v1736
		v1855 = v1759
		v1856 = v1769
		goto L413
	} else {
		goto L417
	}
L417:
	;
	v1773 = int32(1)
	v1774 = F_lappend(m, v1736, v1770)
	mBase = m.M
	v1775 = m.ExcPending
	if v1775 != 0 {
		goto L4
	} else {
		goto L418
	}
L418:
	;
	v1777 = v1765 + int32(1)
	v1778 = *(*int32)(unsafe.Add(mBase, uint32(v1604)+4))
	if v1778 <= v1777 {
		v1824 = v1774
		v1827 = v1773
		goto L414
	} else {
		goto L419
	}
L419:
	;
	v1784 = v1774
	v1790 = v1777
	goto L420
L420:
	;
	v1807 = *(*int32)(unsafe.Add(mBase, uint32(v1604)+12))
	v1810 = v1807 + v1790<<(uint(int32(2))%32)
	v1811 = *(*int32)(unsafe.Add(mBase, uint32(v1810)))
	v1812 = *(*int32)(unsafe.Add(mBase, uint32(v1811)))
	if v1812 != v1734 {
		v1852 = v1784
		v1855 = v1773
		v1856 = v1810
		goto L413
	} else {
		goto L422
	}
L421:
	;
	v1824 = v1814
	v1827 = v1773
	goto L414
L422:
	;
	v1814 = F_lappend(m, v1784, v1811)
	mBase = m.M
	v1815 = m.ExcPending
	if v1815 != 0 {
		goto L4
	} else {
		goto L423
	}
L423:
	;
	v1817 = v1790 + int32(1)
	v1818 = *(*int32)(unsafe.Add(mBase, uint32(v1604)+4))
	if v1817 < v1818 {
		v1784 = v1814
		v1790 = v1817
		goto L420
	} else {
		goto L424
	}
L424:
	;
	goto L421
L425:
	;
	if v1613&int32(6) != int32(4) {
		v2035 = v1737
		goto L444
	} else {
		goto L445
	}
L426:
	;
	if v1744 == int32(0) {
		goto L427
	} else {
		goto L428
	}
L427:
	;
	v1945 = v1852
	v1948 = v1855
	v1953 = int32(0)
	goto L425
L428:
	;
	goto L429
L429:
	;
	v1880 = *(*int32)(unsafe.Add(mBase, uint32(v1611)+12))
	v1881 = v1744 - v1880
	v1883 = v1881 >> (uint(int32(2)) % 32)
	v1884 = *(*int32)(unsafe.Add(mBase, uint32(v1611)+4))
	if v1884 <= v1883 {
		goto L430
	} else {
		goto L431
	}
L430:
	;
	v1945 = v1852
	v1948 = v1855
	v1953 = int32(0)
	goto L425
L431:
	;
	goto L432
L432:
	;
	v1887 = *(*int32)(unsafe.Add(mBase, uint32(v1611)+12))
	v1888 = v1887 + v1881
	v1889 = *(*int32)(unsafe.Add(mBase, uint32(v1888)))
	v1890 = *(*int32)(unsafe.Add(mBase, uint32(v1889)))
	if v1890 != v1734 {
		v1945 = v1852
		v1948 = v1855
		v1953 = v1888
		goto L425
	} else {
		goto L433
	}
L433:
	;
	v1892 = int32(1)
	v1893 = F_lappend(m, v1852, v1889)
	mBase = m.M
	v1894 = m.ExcPending
	if v1894 != 0 {
		goto L4
	} else {
		goto L434
	}
L434:
	;
	v1896 = v1883 + int32(1)
	v1897 = *(*int32)(unsafe.Add(mBase, uint32(v1611)+4))
	if v1897 <= v1896 {
		goto L435
	} else {
		goto L436
	}
L435:
	;
	v1945 = v1893
	v1948 = v1892
	v1953 = int32(0)
	goto L425
L436:
	;
	goto L437
L437:
	;
	v1904 = v1893
	v1910 = v1896
	goto L438
L438:
	;
	v1927 = *(*int32)(unsafe.Add(mBase, uint32(v1611)+12))
	v1930 = v1927 + v1910<<(uint(int32(2))%32)
	v1931 = *(*int32)(unsafe.Add(mBase, uint32(v1930)))
	v1932 = *(*int32)(unsafe.Add(mBase, uint32(v1931)))
	if v1932 != v1734 {
		v1945 = v1904
		v1948 = v1892
		v1953 = v1930
		goto L425
	} else {
		goto L440
	}
L439:
	;
	v1945 = v1934
	v1948 = v1892
	v1953 = int32(0)
	goto L425
L440:
	;
	v1934 = F_lappend(m, v1904, v1931)
	mBase = m.M
	v1935 = m.ExcPending
	if v1935 != 0 {
		goto L4
	} else {
		goto L441
	}
L441:
	;
	v1937 = v1910 + int32(1)
	v1938 = *(*int32)(unsafe.Add(mBase, uint32(v1611)+4))
	if v1937 < v1938 {
		v1904 = v1934
		v1910 = v1937
		goto L438
	} else {
		goto L442
	}
L442:
	;
	goto L439
L443:
	;
	v2067 = v1734 + int32(1)
	v2068 = *(*int32)(unsafe.Add(mBase, uint32(v1670)))
	if v2067 < v2068 {
		v1734 = v2067
		v1736 = v2043
		v1737 = v2044
		v1740 = v1856
		v1744 = v1953
		goto L411
	} else {
		goto L460
	}
L444:
	;
	if v1948 == int32(0) {
		v2151 = v1653
		goto L382
	} else {
		goto L459
	}
L445:
	;
	if v1737 == int32(0) {
		goto L447
	} else {
		goto L448
	}
L446:
	;
	v1997 = v1983
	v2000 = v1986
	goto L454
L447:
	;
	if v1948 != 0 {
		v2043 = v1945
		v2044 = int32(0)
		goto L443
	} else {
		goto L453
	}
L448:
	;
	v1972 = *(*int32)(unsafe.Add(mBase, uint32(v1607)+12))
	v1973 = v1737 - v1972
	v1975 = v1973 >> (uint(int32(2)) % 32)
	v1976 = *(*int32)(unsafe.Add(mBase, uint32(v1607)+4))
	if v1976 <= v1975 {
		goto L447
	} else {
		goto L449
	}
L449:
	;
	v1978 = *(*int32)(unsafe.Add(mBase, uint32(v1607)+12))
	v1979 = v1978 + v1973
	v1980 = *(*int32)(unsafe.Add(mBase, uint32(v1979)))
	v1981 = *(*int32)(unsafe.Add(mBase, uint32(v1980)))
	if v1981 != v1734 {
		v2035 = v1979
		goto L444
	} else {
		goto L450
	}
L450:
	;
	v1983 = F_lappend(m, v1945, v1980)
	mBase = m.M
	v1984 = m.ExcPending
	if v1984 != 0 {
		goto L4
	} else {
		goto L451
	}
L451:
	;
	v1986 = v1975 + int32(1)
	v1987 = *(*int32)(unsafe.Add(mBase, uint32(v1607)+4))
	if v1986 < v1987 {
		goto L446
	} else {
		goto L452
	}
L452:
	;
	v2043 = v1983
	v2044 = int32(0)
	goto L443
L453:
	;
	v2151 = v1653
	goto L382
L454:
	;
	v2020 = *(*int32)(unsafe.Add(mBase, uint32(v1607)+12))
	v2023 = v2020 + v2000<<(uint(int32(2))%32)
	v2024 = *(*int32)(unsafe.Add(mBase, uint32(v2023)))
	v2025 = *(*int32)(unsafe.Add(mBase, uint32(v2024)))
	if v2025 != v1734 {
		v2043 = v1997
		v2044 = v2023
		goto L443
	} else {
		goto L456
	}
L455:
	;
	v2043 = v2027
	v2044 = int32(0)
	goto L443
L456:
	;
	v2027 = F_lappend(m, v1997, v2024)
	mBase = m.M
	v2028 = m.ExcPending
	if v2028 != 0 {
		goto L4
	} else {
		goto L457
	}
L457:
	;
	v2030 = v2000 + int32(1)
	v2031 = *(*int32)(unsafe.Add(mBase, uint32(v1607)+4))
	if v2030 < v2031 {
		v1997 = v2027
		v2000 = v2030
		goto L454
	} else {
		goto L458
	}
L458:
	;
	goto L455
L459:
	;
	v2043 = v1945
	v2044 = v2035
	goto L443
L460:
	;
	goto L412
L461:
	;
	v2130 = v2101
	goto L387
L462:
	;
	v2134 = v1656 + int32(1)
	v2135 = *(*int32)(unsafe.Add(mBase, uint32(v1630)+4))
	if v2134 < v2135 {
		v1653 = v2131
		v1656 = v2134
		goto L385
	} else {
		goto L463
	}
L463:
	;
	goto L386
L464:
	;
	goto L381
L465:
	;
	v2172 = *(*int32)(unsafe.Add(mBase, uint32(v2169)+4))
	if v2172 <= int32(0) {
		v2310 = l0
		v2313 = v30
		v2321 = v1252
		v2324 = v2168
		goto L338
	} else {
		goto L466
	}
L466:
	;
	v2175 = *(*int32)(unsafe.Add(mBase, uint32(v2169)+12))
	v2176 = *(*int32)(unsafe.Add(mBase, uint32(v2175)))
	v2177 = *(*int32)(unsafe.Add(mBase, uint32(v2176)))
	v2183 = *(*int32)(unsafe.Add(mBase, uint32(v2175+v2172<<(uint(int32(2))%32)-int32(4))))
	v2184 = *(*int32)(unsafe.Add(mBase, uint32(v2183)))
	if v2177 == v2184 {
		goto L468
	} else {
		goto L469
	}
L467:
	;
	if v2243 <= v2264 {
		v2310 = l0
		v2313 = v30
		v2321 = v1252
		v2324 = v2168
		goto L338
	} else {
		goto L479
	}
L468:
	;
	v2186 = int32(0)
	v2243 = v2172
	v2245 = v2186
	v2264 = v2186
	goto L467
L469:
	;
	goto L470
L470:
	;
	v2188 = int32(0)
	v2194 = v2188
	v2197 = v2176
	v2198 = v2188
	goto L472
L471:
	;
	v2233 = int32(2)
	v2243 = v2221
	v2245 = v2217
	v2264 = v2220 << (uint(v2233) % 32) >> (uint(v2233) % 32)
	goto L467
L472:
	;
	v2217 = F_lappend(m, v2198, v2197)
	mBase = m.M
	v2218 = m.ExcPending
	if v2218 != 0 {
		goto L4
	} else {
		goto L475
	}
L473:
	;
	v2231 = F_list_concat(m, v1252, int32(0))
	mBase = m.M
	v2232 = m.ExcPending
	if v2232 != 0 {
		goto L4
	} else {
		goto L478
	}
L474:
	;
	goto L473
L475:
	;
	v2220 = v2194 + int32(1)
	v2221 = *(*int32)(unsafe.Add(mBase, uint32(v2169)+4))
	if v2221 <= v2220 {
		goto L474
	} else {
		goto L476
	}
L476:
	;
	v2223 = *(*int32)(unsafe.Add(mBase, uint32(v2169)+12))
	v2227 = *(*int32)(unsafe.Add(mBase, uint32(v2223+v2220<<(uint(int32(2))%32))))
	v2228 = *(*int32)(unsafe.Add(mBase, uint32(v2227)))
	if v2184 != v2228 {
		v2194 = v2220
		v2197 = v2227
		v2198 = v2217
		goto L472
	} else {
		goto L477
	}
L477:
	;
	goto L471
L478:
	;
	v2463 = l0
	v2466 = v30
	v2474 = v2231
	goto L8
L479:
	;
	v2270 = v2264
	v2280 = v2168
	goto L480
L480:
	;
	v2295 = *(*int32)(unsafe.Add(mBase, uint32(v2169)+12))
	v2299 = *(*int32)(unsafe.Add(mBase, uint32(v2295+v2270<<(uint(int32(2))%32))))
	v2300 = *(*int32)(unsafe.Add(mBase, uint32(v2299)+12))
	v2301 = *(*int32)(unsafe.Add(mBase, uint32(v2299)+16))
	v2302 = F_get_steps_using_prefix(m, l0, int32(1), int32(0), v2300, v2301, v1311, v2245)
	mBase = m.M
	v2303 = m.ExcPending
	if v2303 != 0 {
		goto L4
	} else {
		goto L482
	}
L481:
	;
	v2310 = l0
	v2313 = v30
	v2321 = v1252
	v2324 = v2304
	goto L338
L482:
	;
	v2304 = F_list_concat(m, v2280, v2302)
	mBase = m.M
	v2305 = m.ExcPending
	if v2305 != 0 {
		goto L4
	} else {
		goto L483
	}
L483:
	;
	v2307 = v2270 + int32(1)
	v2308 = *(*int32)(unsafe.Add(mBase, uint32(v2169)+4))
	if v2307 < v2308 {
		v2270 = v2307
		v2280 = v2304
		goto L480
	} else {
		goto L484
	}
L484:
	;
	goto L481
L485:
	;
	v2463 = v2310
	v2466 = v2313
	v2474 = v2337
	goto L8
L486:
	;
	v2343 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1315))))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v2343
	F_errmsg_internal(m, int32(_a_F_gen_partprune_steps_internal_0), v30+int32(16))
	mBase = m.M
	v2349 = m.ExcPending
	if v2349 != 0 {
		goto L4
	} else {
		goto L487
	}
L487:
	;
	F_errfinish(m, int32(_a_F_gen_partprune_steps_internal_1), int32(1768), int32(_a_F_gen_partprune_steps_internal_4))
	mBase = m.M
	v2354 = m.ExcPending
	if v2354 != 0 {
		goto L4
	} else {
		goto L488
	}
L488:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L489:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2356))) = int32(377)
	v2360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2360 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2356)+20)) = v1255
	*(*int64)(unsafe.Add(mBase, uint32(v2356)+12)) = int64(0)
	v2367 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v2356)+8)) = uint16(v2367)
	*(*int32)(unsafe.Add(mBase, uint32(v2356)+4)) = v2360
	v2370 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2371 = F_lappend(m, v2370, v2356)
	mBase = m.M
	v2372 = m.ExcPending
	if v2372 != 0 {
		goto L4
	} else {
		goto L490
	}
L490:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v2371
	v2374 = F_lappend(m, v1252, v2356)
	mBase = m.M
	v2375 = m.ExcPending
	if v2375 != 0 {
		goto L4
	} else {
		goto L491
	}
L491:
	;
	v2463 = l0
	v2466 = v30
	v2474 = v2374
	goto L8
L492:
	;
	v2439 = int32(*(*int16)(unsafe.Add(mBase, uint32(v33)+2)))
	if v2438 != v2439 {
		v2463 = l0
		v2466 = v30
		v2474 = v2387
		goto L8
	} else {
		goto L505
	}
L493:
	;
	v2438 = int32(0)
	goto L492
L494:
	;
	goto L495
L495:
	;
	v2410 = int32(1)
	v2411 = *(*int32)(unsafe.Add(mBase, uint32(v2393)+4))
	if v2411 <= v2410 {
		goto L496
	} else {
		goto L497
	}
L496:
	;
	v2414 = v2410
	goto L498
L497:
	;
	v2414 = v2411
	goto L498
L498:
	;
	v2418 = int32(0)
	v2420 = v2403
	goto L499
L499:
	;
	v2426 = *(*int32)(unsafe.Add(mBase, uint32(v2393+int32(8)+v2418<<(uint(int32(2))%32))))
	if v2426 != 0 {
		goto L501
	} else {
		goto L502
	}
L500:
	;
	v2438 = v2429
	goto L492
L501:
	;
	v2429 = v2420 + base.I32_popcnt(v2426)
	goto L503
L502:
	;
	v2429 = v2420
	goto L503
L503:
	;
	v2431 = v2418 + int32(1)
	if v2431 != v2414 {
		v2418 = v2431
		v2420 = v2429
		goto L499
	} else {
		goto L504
	}
L504:
	;
	goto L500
L505:
	;
	v2442 = F_palloc0(m, int32(24))
	mBase = m.M
	v2443 = m.ExcPending
	if v2443 != 0 {
		goto L4
	} else {
		goto L506
	}
L506:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2442))) = int32(377)
	v2446 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2446 + int32(1)
	v2450 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2442)+20)) = v2450
	*(*int64)(unsafe.Add(mBase, uint32(v2442)+12)) = int64(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v2442)+8)) = uint16(v2450)
	*(*int32)(unsafe.Add(mBase, uint32(v2442)+4)) = v2446
	v2457 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2458 = F_lappend(m, v2457, v2442)
	mBase = m.M
	v2459 = m.ExcPending
	if v2459 != 0 {
		goto L4
	} else {
		goto L507
	}
L507:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v2458
	v2461 = F_lappend(m, v2387, v2442)
	mBase = m.M
	v2462 = m.ExcPending
	if v2462 != 0 {
		goto L4
	} else {
		goto L508
	}
L508:
	;
	v2463 = l0
	v2466 = v30
	v2474 = v2461
	goto L8
L509:
	;
	v2559 = v2466
	v2567 = int32(0)
	goto L1
L510:
	;
	goto L511
L511:
	;
	v2493 = *(*int32)(unsafe.Add(mBase, uint32(v2474)+4))
	if v2493 < int32(2) {
		v2559 = v2466
		v2567 = v2474
		goto L1
	} else {
		goto L512
	}
L512:
	;
	v2496 = int32(0)
	v2500 = v2496
	v2502 = v2496
	goto L513
L513:
	;
	v2525 = *(*int32)(unsafe.Add(mBase, uint32(v2474)+12))
	v2529 = *(*int32)(unsafe.Add(mBase, uint32(v2525+v2500<<(uint(int32(2))%32))))
	v2530 = *(*int32)(unsafe.Add(mBase, uint32(v2529)+4))
	v2531 = F_lappend_int(m, v2502, v2530)
	mBase = m.M
	v2532 = m.ExcPending
	if v2532 != 0 {
		goto L4
	} else {
		goto L515
	}
L514:
	;
	v2538 = F_palloc0(m, int32(16))
	mBase = m.M
	v2539 = m.ExcPending
	if v2539 != 0 {
		goto L4
	} else {
		goto L517
	}
L515:
	;
	v2534 = v2500 + int32(1)
	v2535 = *(*int32)(unsafe.Add(mBase, uint32(v2474)+4))
	if v2534 < v2535 {
		v2500 = v2534
		v2502 = v2531
		goto L513
	} else {
		goto L516
	}
L516:
	;
	goto L514
L517:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2538))) = int32(378)
	v2542 = *(*int32)(unsafe.Add(mBase, uint32(v2463)+16))
	v2543 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2463)+16)) = v2542 + v2543
	*(*int32)(unsafe.Add(mBase, uint32(v2538)+12)) = v2531
	*(*int32)(unsafe.Add(mBase, uint32(v2538)+8)) = v2543
	*(*int32)(unsafe.Add(mBase, uint32(v2538)+4)) = v2542
	v2550 = *(*int32)(unsafe.Add(mBase, uint32(v2463)+8))
	v2551 = F_lappend(m, v2550, v2538)
	mBase = m.M
	v2552 = m.ExcPending
	if v2552 != 0 {
		goto L4
	} else {
		goto L518
	}
L518:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2463)+8)) = v2551
	v2554 = F_lappend(m, v2474, v2538)
	mBase = m.M
	v2555 = m.ExcPending
	if v2555 != 0 {
		goto L4
	} else {
		goto L519
	}
L519:
	;
	v2559 = v2466
	v2567 = v2554
	goto L1
}
func F_generate_setop_tlist(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
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
	var v77 int32
	_ = v77
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
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
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
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
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
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	v8 = int32(0)
	v17 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v17)
	v28 = v8
	v31 = v17
	v32 = v8
	goto L1
L1:
	;
	v36 = int32(0)
	if l0 == v36 {
		v46 = v36
		goto L3
	} else {
		goto L4
	}
L2:
	;
	return v32
L3:
	;
	v47 = int32(0)
	if l1 == v47 {
		v58 = v47
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v40 <= v28 {
		v46 = int32(0)
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v46 = v42 + v28<<(uint(int32(2))%32)
	goto L3
L6:
	;
	if l4 == int32(0) {
		v67 = v47
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v52 <= v28 {
		v58 = int32(0)
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v58 = v54 + v28<<(uint(int32(2))%32)
	goto L6
L9:
	;
	v68 = int32(0)
	if l5 == v68 {
		v77 = v68
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v61 <= v28 {
		v67 = v47
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v67 = v63 + v28<<(uint(int32(2))%32)
	goto L9
L12:
	;
	if v46 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v71 <= v28 {
		v77 = v68
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	v77 = v73 + v28<<(uint(int32(2))%32)
	goto L12
L15:
	;
	goto L2
L16:
	;
	if v58 == int32(0) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	if v67 == int32(0) {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	if v77 == int32(0) {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	if l3 == int32(0) {
		v99 = v90
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v115 = F_exprType(m, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L27
	} else {
		goto L32
	}
L21:
	;
	v100 = int32(*(*int16)(unsafe.Add(mBase, uint32(v89)+8)))
	v101 = F_exprType(m, v99)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L27
	} else {
		goto L28
	}
L22:
	;
	if v90 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v99 = int32(0)
	goto L21
L24:
	;
	goto L25
L25:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	if v96 == int32(7) {
		v114 = v90
		goto L20
	} else {
		goto L26
	}
L26:
	;
	v99 = v90
	goto L21
L27:
	;
	return int32(0)
L28:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	v106 = F_exprTypmod(m, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	v109 = F_exprCollation(m, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v112 = F_makeVar(m, l2, v100, v101, v106, v109, int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L27
	} else {
		goto L31
	}
L31:
	;
	v114 = v112
	goto L20
L32:
	;
	if v115 != v88 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v120 = F_coerce_to_common_type(m, int32(0), v114, v88, int32(_a_F_generate_setop_tlist_0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L27
	} else {
		goto L36
	}
L34:
	;
	v124 = v114
	goto L35
L35:
	;
	v125 = F_exprCollation(m, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L27
	} else {
		goto L37
	}
L36:
	;
	v122 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v122)
	v124 = v120
	goto L35
L37:
	;
	if v125 != v87 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v128 = F_exprType(m, v124)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L27
	} else {
		goto L41
	}
L39:
	;
	v139 = v124
	goto L40
L40:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v142 = F_pstrdup(m, v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L27
	} else {
		goto L44
	}
L41:
	;
	v130 = F_exprTypmod(m, v124)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L27
	} else {
		goto L42
	}
L42:
	;
	v135 = F_applyRelabelType(m, v124, v128, v130, v87, int32(2), int32(-1), int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L27
	} else {
		goto L43
	}
L43:
	;
	v137 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v137)
	v139 = v135
	goto L40
L44:
	;
	v145 = F_makeTargetEntry(m, v139, base.I32_extend16_s(v31), v142, int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L27
	} else {
		goto L45
	}
L45:
	;
	v147 = int32(*(*int16)(unsafe.Add(mBase, uint32(v145)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v145)+16)) = v147
	v149 = int32(1)
	v153 = F_lappend(m, v32, v145)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L27
	} else {
		goto L46
	}
L46:
	;
	v28 = v28 + v149
	v31 = v31 + v149
	v32 = v153
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
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
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
	var v114 int32
	_ = v114
	var v115 int64
	_ = v115
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
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
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = int32(2)
						v37 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v37)
						return int32(0)
					}
				} else {
					v42 = int32(0)
					if base.B2i32(base.Ui32(v17) <= base.Ui32(v27))&base.B2i32(v42 < v17) == v42 {
						F_end_MultiFuncCall(m, l0)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v49)+20)) = int32(2)
							v52 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v52)
							return int32(0)
						}
					} else {
						v56 = int32(_a_F_generate_subscripts_0)
						v57 = *(*int32)(unsafe.Add(mBase, _c_F_generate_subscripts[0]))
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
						*(*int32)(unsafe.Add(mBase, _c_F_generate_subscripts[0])) = v59
						v62 = F_palloc(m, int32(12))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
							if v64 == int32(-1) {
								v67 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
								v68 = *(*int32)(unsafe.Add(mBase, uint32(v13)+36))
								v75 = v67
								v76 = v68
							} else {
								v70 = v13 + int32(16)
								v71 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
								v75 = v70
								v76 = v70 + v71<<(uint(int32(2))%32)
							}
							v80 = v17<<(uint(int32(2))%32) - int32(4)
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v76+v80)))
							*(*int32)(unsafe.Add(mBase, uint32(v62))) = v82
							v85 = *(*int32)(unsafe.Add(mBase, uint32(v80+v75)))
							*(*int32)(unsafe.Add(mBase, uint32(v62)+4)) = v82 + v85 - int32(1)
							v91 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
							if int32(3) <= v91 {
								v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								v97 = base.B2i32(v94 != int32(0))
							} else {
								v97 = int32(0)
							}
							*(*uint8)(unsafe.Add(mBase, uint32(v62)+8)) = uint8(v97)
							*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v62
							*(*int32)(unsafe.Add(mBase, _c_F_generate_subscripts[0])) = v57
							v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+16))
							v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
							v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
							v112 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
							if v111 <= v112 {
								v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+8)))
								v115 = *(*int64)(unsafe.Add(mBase, uint32(v109)))
								*(*int64)(unsafe.Add(mBase, uint32(v109))) = v115 + int64(1)
								v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v119)+20)) = int32(1)
								if v114 == int32(0) {
									v124 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
									*(*int32)(unsafe.Add(mBase, uint32(v110))) = v124 + int32(1)
									return v124
								} else {
									v129 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v110)+4)) = v129 - int32(1)
									return v129
								}
							} else {
								F_end_MultiFuncCall(m, l0)
								mBase = m.M
								v135 = m.ExcPending
								if v135 != 0 {
									return int32(0)
								} else {
									v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v136)+20)) = int32(2)
									v139 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v139)
									return int32(0)
								}
							}
						}
					}
				}
			}
		}
	} else {
		v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+16))
		v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
		v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
		v112 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
		if v111 <= v112 {
			v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+8)))
			v115 = *(*int64)(unsafe.Add(mBase, uint32(v109)))
			*(*int64)(unsafe.Add(mBase, uint32(v109))) = v115 + int64(1)
			v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v119)+20)) = int32(1)
			if v114 == int32(0) {
				v124 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
				*(*int32)(unsafe.Add(mBase, uint32(v110))) = v124 + int32(1)
				return v124
			} else {
				v129 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v110)+4)) = v129 - int32(1)
				return v129
			}
		} else {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v135 = m.ExcPending
			if v135 != 0 {
				return int32(0)
			} else {
				v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v136)+20)) = int32(2)
				v139 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v139)
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
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v16 int32
	_ = v16
	var v20 int64
	_ = v20
	var v23 int64
	_ = v23
	var v26 int64
	_ = v26
	var v29 int64
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	v1 = l0
	v6 = F_palloc(m, int32(16))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v11 = int64(base.Ui64(v1) >> (uint(int64(40)) % 64))
		*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(v11)
		v16 = base.I32_div_u_s(l1<<(uint(int32(12))%32), int32(_a_F_generate_uuidv7_0))
		*(*uint8)(unsafe.Add(mBase, uint32(v6)+7)) = uint8(v16)
		*(*uint8)(unsafe.Add(mBase, uint32(v6)+5)) = uint8(v1)
		v20 = int64(base.Ui64(v1) >> (uint(int64(8)) % 64))
		*(*uint8)(unsafe.Add(mBase, uint32(v6)+4)) = uint8(v20)
		v23 = int64(base.Ui64(v1) >> (uint(int64(16)) % 64))
		*(*uint8)(unsafe.Add(mBase, uint32(v6)+3)) = uint8(v23)
		v26 = int64(base.Ui64(v1) >> (uint(int64(24)) % 64))
		*(*uint8)(unsafe.Add(mBase, uint32(v6)+2)) = uint8(v26)
		v29 = int64(base.Ui64(v1) >> (uint(int64(32)) % 64))
		*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)) = uint8(v29)
		v32 = v6 + int32(6)
		v33 = int32(8)
		v34 = int32(base.Ui32(v16) >> (uint(v33) % 32))
		*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v34)
		v37 = v6 + v33
		v39 = m.Env.Pgmem_random_bytes(m, v37, v33)
		mBase = m.M
		if v39 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(2600))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_generate_uuidv7_1), int32(0))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_generate_uuidv7_2), int32(638), int32(_a_F_generate_uuidv7_3))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
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
			v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
			v62 = v58&int32(15) | int32(112)
			*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v62)
			v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
			v68 = v64&int32(63) | int32(128)
			*(*uint8)(unsafe.Add(mBase, uint32(v37))) = uint8(v68)
			return v6
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
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
	var v69 int32
	_ = v69
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
	var v103 int32
	_ = v103
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
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
	var v155 int32
	_ = v155
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
	var v199 int32
	_ = v199
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v220 int32
	_ = v220
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
	var v402 float64
	_ = v402
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
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v25)+88))
	if v140 != 0 {
		goto L19
	} else {
		goto L20
	}
L2:
	;
	v124 = v5
	goto L1
L3:
	;
	goto L4
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v29 <= int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v124 = v5
	goto L1
L6:
	;
	goto L7
L7:
	;
	v37 = v5
	v39 = v29
	v40 = v5
	goto L8
L8:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v52+v40<<(uint(int32(2))%32))))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	if v57 == int32(0) {
		v101 = v37
		v103 = v39
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v124 = v101
	goto L1
L10:
	;
	v117 = v40 + int32(1)
	if v117 < v103 {
		v37 = v101
		v39 = v103
		v40 = v117
		goto L8
	} else {
		goto L18
	}
L11:
	;
	v60 = int32(0)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v61 <= v60 {
		v101 = v37
		v103 = v39
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v68 = v60
	v69 = v37
	goto L13
L13:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v84+v68<<(uint(int32(2))%32))))
	v89 = F_lappend(m, v69, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v101 = v89
	v103 = v95
	goto L10
L15:
	;
	return
L16:
	;
	v92 = v68 + int32(1)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v92 < v93 {
		v68 = v92
		v69 = v89
		goto L13
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	goto L9
L19:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)+4))
	if v141 <= int32(0) {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v220 = v124
	goto L21
L21:
	;
	v233 = *(*float64)(unsafe.Add(mBase, uint32(l3)+56))
	if base.F64_lt(v233, float64(1)) == int32(0) {
		v302 = v233
		goto L36
	} else {
		goto L37
	}
L22:
	;
	v211 = F_list_concat(m, v199, v124)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L15
	} else {
		goto L35
	}
L23:
	;
	v199 = int32(0)
	goto L22
L24:
	;
	goto L25
L25:
	;
	v145 = int32(0)
	v151 = v145
	v155 = v145
	goto L26
L26:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v140)+12))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v167+v151<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v171
	v177 = F_list_make1_impl(m, int32(1), v23+int32(4))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L15
	} else {
		goto L28
	}
L27:
	;
	v199 = v186
	goto L22
L28:
	;
	v180 = F_predicate_implied_by(m, v177, v124, int32(0))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L15
	} else {
		goto L29
	}
L29:
	;
	if v180 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v184 = F_list_concat(m, v155, v177)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L15
	} else {
		goto L33
	}
L31:
	;
	v186 = v155
	goto L32
L32:
	;
	v188 = v151 + int32(1)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v140)+4))
	if v188 < v189 {
		v151 = v188
		v155 = v186
		goto L26
	} else {
		goto L34
	}
L33:
	;
	v186 = v184
	goto L32
L34:
	;
	goto L27
L35:
	;
	v220 = v211
	goto L21
L36:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v309)+68))
	v311 = int32(0)
	v313 = F_clauselist_selectivity(m, l0, v220, v310, v311, v311)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L15
	} else {
		goto L50
	}
L37:
	;
	v238 = float64(1)
	if v124 == int32(0) {
		v302 = v238
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	if v241 <= int32(0) {
		v302 = v238
		goto L36
	} else {
		goto L39
	}
L39:
	;
	v249 = int32(0)
	v258 = v238
	goto L40
L40:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v265+v249<<(uint(int32(2))%32))))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v269)+4))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v270)))
	if v271 == int32(20) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v302 = v284
	goto L36
L42:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v270)+28))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)+12))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v275)+4))
	v277 = F_estimate_array_length(m, l0, v276)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L15
	} else {
		goto L45
	}
L43:
	;
	v284 = v258
	goto L44
L44:
	;
	v286 = v249 + int32(1)
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	if v286 < v287 {
		v249 = v286
		v258 = v284
		goto L40
	} else {
		goto L49
	}
L45:
	;
	if base.F64_gt(v277, float64(1)) != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v282 = base.F64_mul(v258, v277)
	goto L48
L47:
	;
	v282 = v258
	goto L48
L48:
	;
	v284 = v282
	goto L44
L49:
	;
	goto L41
L50:
	;
	v315 = *(*float64)(unsafe.Add(mBase, uint32(l3)+40))
	if base.F64_le(v315, float64(0)) != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v319 = *(*float64)(unsafe.Add(mBase, uint32(v318)+120))
	v323 = base.F64_nearest(base.F64_div(base.F64_mul(v313, v319), v302))
	goto L53
L52:
	;
	v323 = v315
	goto L53
L53:
	;
	v325 = *(*float64)(unsafe.Add(mBase, uint32(v25)+24))
	if base.F64_gt(v323, v325) != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v327 = v325
	goto L56
L55:
	;
	v327 = v323
	goto L56
L56:
	;
	if base.F64_lt(v327, float64(1)) != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v330 = float64(1)
	goto L59
L58:
	;
	v330 = v327
	goto L59
L59:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	if base.Ui32(v332) < base.Ui32(int32(2)) {
		v344 = float64(1)
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	F_get_tablespace_page_costs(m, v345, v23+int32(8), int32(0))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L15
	} else {
		goto L63
	}
L61:
	;
	v335 = float64(1)
	if base.F64_gt(v325, v335) == int32(0) {
		v344 = v335
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v344 = base.F64_ceil(base.F64_div(base.F64_mul(v330, base.F64_convert_i32_u(v332)), v325))
	goto L60
L63:
	;
	v351 = base.F64_mul(l2, v302)
	if base.F64_gt(v351, float64(1)) != 0 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v417 = F_index_other_operands_eval_cost(m, l0, v124)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L15
	} else {
		goto L86
	}
L65:
	;
	v354 = base.F64_mul(v351, v344)
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v360 = int32(1)
	if base.Ui32(v355) <= base.Ui32(v360) {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	goto L67
L67:
	;
	v412 = *(*float64)(unsafe.Add(mBase, uint32(v23)+8))
	v415 = base.F64_mul(v344, v412)
	goto L64
L68:
	;
	v409 = *(*float64)(unsafe.Add(mBase, uint32(v23)+8))
	v415 = base.F64_div(base.F64_mul(v408, v409), l2)
	goto L64
L69:
	;
	v363 = v360
	goto L71
L70:
	;
	v363 = v355
	goto L71
L71:
	;
	v364 = base.F64_convert_i32_u(v363)
	v365 = base.F64_add(v364, v364)
	v366 = float64(1)
	v368 = *(*int32)(unsafe.Add(mBase, _c_F_genericcostestimate[0]))
	v371 = *(*float64)(unsafe.Add(mBase, uint32(l0)+288))
	v372 = base.F64_add(base.F64_convert_i32_u(v355), v371)
	if base.F64_gt(v372, v366) != 0 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v408 = v402
	goto L68
L73:
	;
	v376 = v372
	goto L75
L74:
	;
	v376 = v366
	goto L75
L75:
	;
	v377 = base.F64_div(base.F64_mul(v364, base.F64_convert_i32_s(v368)), v376)
	if base.F64_le(v377, float64(1)) != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v381 = v366
	goto L78
L77:
	;
	v381 = base.F64_ceil(v377)
	goto L78
L78:
	;
	if base.F64_le(v364, v381) != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v385 = base.F64_div(base.F64_mul(v354, v365), base.F64_add(v365, v354))
	if base.F64_ge(v385, v364) != 0 {
		v402 = v364
		goto L72
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v390 = base.F64_div(base.F64_mul(v365, v381), base.F64_sub(v365, v381))
	if base.F64_ge(v390, v354) != 0 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v408 = base.F64_ceil(v385)
	goto L68
L83:
	;
	v400 = base.F64_div(base.F64_mul(v354, v365), base.F64_add(v365, v354))
	goto L85
L84:
	;
	v400 = base.F64_add(v381, base.F64_div(base.F64_mul(base.F64_sub(v364, v381), base.F64_sub(v354, v390)), v364))
	goto L85
L85:
	;
	v402 = base.F64_ceil(v400)
	goto L72
L86:
	;
	v419 = F_index_other_operands_eval_cost(m, l0, v139)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L15
	} else {
		goto L87
	}
L87:
	;
	if v124 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	v423 = v422
	goto L90
L89:
	;
	v423 = int32(0)
	goto L90
L90:
	;
	v425 = *(*float64)(unsafe.Add(mBase, _c_F_genericcostestimate[1]))
	if v139 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	v427 = v426
	goto L93
L92:
	;
	v427 = int32(0)
	goto L93
L93:
	;
	v429 = *(*float64)(unsafe.Add(mBase, _c_F_genericcostestimate[2]))
	*(*float64)(unsafe.Add(mBase, uint32(l3)+40)) = v330
	*(*float64)(unsafe.Add(mBase, uint32(l3)+32)) = v344
	*(*int64)(unsafe.Add(mBase, uint32(l3)+24)) = int64(0)
	*(*float64)(unsafe.Add(mBase, uint32(l3)+16)) = v313
	v435 = base.F64_add(v417, v419)
	*(*float64)(unsafe.Add(mBase, uint32(l3))) = v435
	*(*float64)(unsafe.Add(mBase, uint32(l3)+8)) = base.F64_add(base.F64_mul(base.F64_mul(v302, v330), base.F64_add(v429, base.F64_mul(v425, base.F64_convert_i32_s(v427+v423)))), base.F64_add(v415, v435))
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
	var v67 int32
	_ = v67
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
			v67 = v60
		} else {
			v67 = v4
		}
		if v13 == int32(0) {
		} else {
			v78 = v67
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	v10 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(256)
	m.G0 = v19
	*(*int32)(unsafe.Add(mBase, uint32(v19)+180)) = v10
	*(*int32)(unsafe.Add(mBase, uint32(v19)+184)) = int32(6)
	v25 = F_GlobalVisHorizonKindForRel(m, l0)
	mBase = m.M
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v25<<(uint(int32(2))%32))+uint32(_c_F_get_actual_variable_endpoint[0])))
	goto L1
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+224)) = v30
	v34 = int32(0)
	v37 = F_index_beginscan(m, l0, l1, v19+int32(184), v34, int32(1), v34)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	return int32(0)
L3:
	;
	v41 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v37)+28)) = uint8(v41)
	v44 = int32(0)
	F_index_rescan(m, v37, l3, v41, v44, v44)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v60 = int32(-1)
	v61 = v10
	goto L9
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L2
	} else {
		goto L38
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L2
	} else {
		goto L35
	}
L7:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v19)+180))
	if v126 != 0 {
		goto L30
	} else {
		goto L31
	}
L8:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v37)+44))
	if v95 == int32(0) {
		goto L6
	} else {
		goto L25
	}
L9:
	;
	v65 = F_index_getnext_tid(m, v37, l2)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L2
	} else {
		goto L11
	}
L10:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	m.T0[v92].(func(*base.Module, int32))(m, l6)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L2
	} else {
		goto L24
	}
L11:
	;
	if v65 == int32(0) {
		v123 = v10
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65)+2)))
	v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65))))
	v73 = v69 | v70<<(uint(int32(16))%32)
	v76 = F_visibilitymap_get_status(m, l0, v73, v19+int32(180))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	if v76&int32(1) != 0 {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	v80 = F_index_fetch_heap(m, v37, l6)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	if v80 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v85 = v61 + int32(1)
	if v73 != v60 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	goto L10
L19:
	;
	v87 = v85
	goto L21
L20:
	;
	v87 = v61
	goto L21
L21:
	;
	if v73 == v60 {
		v60 = v73
		v61 = v87
		goto L9
	} else {
		goto L22
	}
L22:
	;
	if v85 <= int32(100) {
		v60 = v73
		v61 = v87
		goto L9
	} else {
		goto L23
	}
L23:
	;
	v123 = v10
	goto L7
L24:
	;
	goto L8
L25:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+72)))
	if v98 != 0 {
		v123 = v10
		goto L7
	} else {
		goto L26
	}
L26:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v37)+48))
	F_index_deform_tuple(m, v95, v99, v19+int32(48), v19+int32(16))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L2
	} else {
		goto L27
	}
L27:
	;
	v106 = int32(1)
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+16)))
	if v107 == v106 {
		goto L5
	} else {
		goto L28
	}
L28:
	;
	v110 = int32(_a_F_get_actual_variable_endpoint_0)
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_get_actual_variable_endpoint[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_get_actual_variable_endpoint[1])) = l7
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
	v115 = F_datumCopy(m, v114, l5, l4)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L2
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = v115
	*(*int32)(unsafe.Add(mBase, _c_F_get_actual_variable_endpoint[1])) = v111
	v123 = v106
	goto L7
L30:
	;
	F_ReleaseBuffer(m, v126)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L2
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	F_index_endscan(m, v37)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L2
	} else {
		goto L34
	}
L33:
	;
	goto L32
L34:
	;
	m.G0 = v19 + int32(256)
	return v123
L35:
	;
	F_errmsg_internal(m, int32(_a_F_get_actual_variable_endpoint_1), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_get_actual_variable_endpoint_2), int32(_a_F_get_actual_variable_endpoint_3), int32(_a_F_get_actual_variable_endpoint_4))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L2
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
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v152 + int32(4)
	F_errmsg_internal(m, int32(_a_F_get_actual_variable_endpoint_5), v19)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_get_actual_variable_endpoint_2), int32(_a_F_get_actual_variable_endpoint_6), int32(_a_F_get_actual_variable_endpoint_4))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L2
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
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
			v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11+v12)+74)))
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
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v12+v13)+92))
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
func F_get_controlfile(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	v8 = m.G0
	v10 = v8 - int32(1040)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(_a_F_get_controlfile_0)
	v19 = F_pg_snprintf(m, v10+int32(16), int32(1024), int32(_a_F_get_controlfile_1), v10)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v24 = v10 + int32(16)
		v25 = m.G0
		v27 = v25 + int32(-64)
		m.G0 = v27
		v30 = F_palloc(m, int32(296))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			v33 = F_OpenTransientFile(m, v24, int32(0))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				if v33 != int32(-1) {
					v37 = int32(296)
					v38 = F_read(m, v33, v30, v37)
					mBase = m.M
					if v38 != v37 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							if v38 < int32(0) {
								F_errcode_for_file_access(m)
								mBase = m.M
								v98 = m.ExcPending
								if v98 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v24
									F_errmsg(m, int32(_a_F_get_controlfile_2), v25+int32(-32))
									mBase = m.M
									v104 = m.ExcPending
									if v104 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_get_controlfile_3), int32(108), int32(_a_F_get_controlfile_4))
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
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
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v27)+56)) = int32(296)
									*(*int32)(unsafe.Add(mBase, uint32(v27)+52)) = v38
									*(*int32)(unsafe.Add(mBase, uint32(v27)+48)) = v24
									F_errmsg(m, int32(_a_F_get_controlfile_5), v25+int32(-16))
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_get_controlfile_3), int32(117), int32(_a_F_get_controlfile_4))
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
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
						v64 = F_CloseTransientFile(m, v33)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							if v64 != 0 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v113 = m.ExcPending
								if v113 != 0 {
									return int32(0)
								} else {
									F_errcode_for_file_access(m)
									mBase = m.M
									v115 = m.ExcPending
									if v115 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v24
										F_errmsg(m, int32(_a_F_get_controlfile_6), v25+int32(-48))
										mBase = m.M
										v121 = m.ExcPending
										if v121 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_get_controlfile_3), int32(129), int32(_a_F_get_controlfile_4))
											mBase = m.M
											v126 = m.ExcPending
											if v126 != 0 {
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
								v66 = int32(-1)
								v68 = m.Env.Pgmem_crc32c(m, v66, v30, int32(292))
								mBase = m.M
								v69 = *(*int32)(unsafe.Add(mBase, uint32(v30)+292))
								*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(base.B2i32(v68^v69 == v66))
								v75 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
								if v75&int32(_a_F_get_controlfile_7) != 0 {
									v78 = int32(0)
								} else {
									v78 = v75
								}
								if v78 != 0 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v130 = m.ExcPending
									if v130 != 0 {
										return int32(0)
									} else {
										F_errmsg_internal(m, int32(_a_F_get_controlfile_8), int32(0))
										mBase = m.M
										v134 = m.ExcPending
										if v134 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_get_controlfile_3), int32(168), int32(_a_F_get_controlfile_4))
											mBase = m.M
											v139 = m.ExcPending
											if v139 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									m.G0 = v27 - int32(-64)
									m.G0 = v10 + int32(1040)
									return v30
								}
							}
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return int32(0)
					} else {
						F_errcode_for_file_access(m)
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v27))) = v24
							F_errmsg(m, int32(_a_F_get_controlfile_9), v27)
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_get_controlfile_3), int32(94), int32(_a_F_get_controlfile_4))
								mBase = m.M
								v96 = m.ExcPending
								if v96 != 0 {
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
	var v63 int32
	_ = v63
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
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
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
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
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
	F_ScanKeyInit(m, v24+int32(-48), int32(2), int32(3), int32(62), l0)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	F_sequence_close(m, v30, int32(1))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L94
	}
L5:
	;
	goto L4
L6:
	;
	v65 = int32(1)
	v70 = F_systable_beginscan(m, v30, int32(2671), v65, int32(0), v65, v24+int32(-48))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v72 = F_systable_getnext(m, v70)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	if v72 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	F_systable_endscan(m, v70)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v72)+16))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+22)))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v78+v79)))
	F_systable_endscan(m, v70)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	goto L5
L13:
	;
	F_LockSharedObject(m, int32(1262), v81, l1)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v88 = F_SearchSysCache1(m, int32(21), v81)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if v88 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+22)))
	v92 = v90 + v91
	v94 = v92 + int32(4)
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v98 == int32(0) {
		v117 = v97
		v118 = v98
		goto L20
	} else {
		goto L21
	}
L17:
	;
	goto L18
L18:
	;
	F_UnlockSharedObject(m, int32(1262), v81, l1)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L93
	}
L19:
	;
	if v118-v117 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L20:
	;
	goto L19
L21:
	;
	if v97 != v98 {
		v117 = v97
		v118 = v98
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v102 = l0
	v103 = v94
	goto L23
L23:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+1)))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+1)))
	if v107 == int32(0) {
		v117 = v106
		v118 = v107
		goto L20
	} else {
		goto L25
	}
L24:
	;
	v117 = v106
	v118 = v107
	goto L20
L25:
	;
	v110 = int32(1)
	if v106 == v107 {
		v102 = v102 + v110
		v103 = v103 + v110
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v81
	if l3 != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	F_ReleaseCatCache(m, v88)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L92
	}
L30:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v92)+68))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v123
	goto L32
L31:
	;
	goto L32
L32:
	;
	if l4 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v92)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v125
	goto L35
L34:
	;
	goto L35
L35:
	;
	if l5 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+77)))
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v127)
	goto L38
L37:
	;
	goto L38
L38:
	;
	if l7 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+79)))
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v129)
	goto L41
L40:
	;
	goto L41
L41:
	;
	if l6 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+78)))
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v131)
	goto L44
L43:
	;
	goto L44
L44:
	;
	if l8 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v92)+84))
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = v133
	goto L47
L46:
	;
	goto L47
L47:
	;
	if l9 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v92)+88))
	*(*int32)(unsafe.Add(mBase, uint32(l9))) = v135
	goto L50
L49:
	;
	goto L50
L50:
	;
	if l10 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v92)+92))
	*(*int32)(unsafe.Add(mBase, uint32(l10))) = v137
	goto L53
L52:
	;
	goto L53
L53:
	;
	if l15 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+76)))
	*(*uint8)(unsafe.Add(mBase, uint32(l15))) = uint8(v139)
	goto L56
L55:
	;
	goto L56
L56:
	;
	if l11 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v143 = F_SysCacheGetAttrNotNull(m, int32(21), v88, int32(13))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	if l12 != 0 {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	v145 = F_text_to_cstring(m, v143)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l11))) = v145
	goto L59
L62:
	;
	v150 = F_SysCacheGetAttrNotNull(m, int32(21), v88, int32(14))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	if l13 != 0 {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	v152 = F_text_to_cstring(m, v150)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l12))) = v152
	goto L64
L67:
	;
	v159 = F_SysCacheGetAttr(m, int32(21), v88, int32(15), v24+int32(-49))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	if l14 != 0 {
		goto L75
	} else {
		goto L76
	}
L70:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+15)))
	if v161 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v165 = int32(0)
	goto L73
L72:
	;
	v163 = F_text_to_cstring(m, v159)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L74
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l13))) = v165
	goto L69
L74:
	;
	v165 = v163
	goto L73
L75:
	;
	v172 = F_SysCacheGetAttr(m, int32(21), v88, int32(16), v24+int32(-49))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	if l16 != 0 {
		goto L83
	} else {
		goto L84
	}
L78:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+15)))
	if v174 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v178 = int32(0)
	goto L81
L80:
	;
	v176 = F_text_to_cstring(m, v172)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L82
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l14))) = v178
	goto L77
L82:
	;
	v178 = v176
	goto L81
L83:
	;
	v185 = F_SysCacheGetAttr(m, int32(21), v88, int32(17), v24+int32(-49))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	F_ReleaseCatCache(m, v88)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L91
	}
L86:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+15)))
	if v187 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v191 = int32(0)
	goto L89
L88:
	;
	v189 = F_text_to_cstring(m, v185)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L90
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l16))) = v191
	goto L85
L90:
	;
	v191 = v189
	goto L89
L91:
	;
	goto L5
L92:
	;
	goto L18
L93:
	;
	goto L3
L94:
	;
	m.G0 = v26 - int32(-64)
	return base.B2i32(v72 != int32(0))
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
	var v32 int32
	_ = v32
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
			v32 = int32(0)
			m.G0 = v7 + int32(16)
			return v32
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
					v32 = v24
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
					*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(base.B2i32(v28 == int32(5)))
					v32 = v24
				}
				m.G0 = v7 + int32(16)
				return v32
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
	var v91 int32
	_ = v91
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
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int64
	_ = v180
	var v182 int64
	_ = v182
	var v184 int32
	_ = v184
	var v186 int64
	_ = v186
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
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
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
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
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v368 int64
	_ = v368
	var v369 int32
	_ = v369
	var v370 int64
	_ = v370
	var v371 int64
	_ = v371
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v408 int32
	_ = v408
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
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
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v591 int32
	_ = v591
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v606 int32
	_ = v606
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v640 int32
	_ = v640
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v678 int32
	_ = v678
	var v686 int32
	_ = v686
	var v702 int32
	_ = v702
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v752 int32
	_ = v752
	var v765 int32
	_ = v765
	var v775 int32
	_ = v775
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v800 int32
	_ = v800
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v816 int32
	_ = v816
	var v826 int32
	_ = v826
	var v839 int32
	_ = v839
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v870 int32
	_ = v870
	var v876 int32
	_ = v876
	var v881 int32
	_ = v881
	var v888 int32
	_ = v888
	var v897 int32
	_ = v897
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v920 int32
	_ = v920
	var v925 int32
	_ = v925
	var v929 int32
	_ = v929
	var v931 int32
	_ = v931
	var v934 int32
	_ = v934
	var v936 int32
	_ = v936
	var v938 int32
	_ = v938
	var v944 int32
	_ = v944
	var v948 int32
	_ = v948
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v965 int32
	_ = v965
	var v970 int32
	_ = v970
	var v974 int32
	_ = v974
	var v980 int32
	_ = v980
	var v985 int32
	_ = v985
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v994 int32
	_ = v994
	var v999 int32
	_ = v999
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1019 int32
	_ = v1019
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1067 int32
	_ = v1067
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1082 int32
	_ = v1082
	var v1085 int32
	_ = v1085
	var v1087 int32
	_ = v1087
	var v1090 int32
	_ = v1090
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1106 int32
	_ = v1106
	var v1120 int32
	_ = v1120
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1136 int32
	_ = v1136
	var v1139 int32
	_ = v1139
	var v1141 int32
	_ = v1141
	var v1144 int32
	_ = v1144
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1153 int32
	_ = v1153
	var v1170 int32
	_ = v1170
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1182 int32
	_ = v1182
	var v1186 int32
	_ = v1186
	var v1191 int32
	_ = v1191
	var v1215 int32
	_ = v1215
	var v1219 int32
	_ = v1219
	var v1224 int32
	_ = v1224
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1265 int32
	_ = v1265
	var v1269 int32
	_ = v1269
	var v1272 int32
	_ = v1272
	var v1274 int32
	_ = v1274
	var v1277 int32
	_ = v1277
	var v1284 int32
	_ = v1284
	var v1286 int32
	_ = v1286
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1308 int32
	_ = v1308
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1319 int32
	_ = v1319
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1336 int32
	_ = v1336
	var v1339 int32
	_ = v1339
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1354 int32
	_ = v1354
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1360 int32
	_ = v1360
	var v1364 int32
	_ = v1364
	var v1367 int32
	_ = v1367
	var v1369 int32
	_ = v1369
	var v1372 int32
	_ = v1372
	var v1379 int32
	_ = v1379
	var v1381 int32
	_ = v1381
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1403 int32
	_ = v1403
	var v1410 int32
	_ = v1410
	var v1414 int32
	_ = v1414
	var v1426 int32
	_ = v1426
	var v1429 int32
	_ = v1429
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1462 int32
	_ = v1462
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
	return v1462
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
	v1462 = v32
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
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(v38+v37-int32(4))))
	v1249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1248)+4)))
	v1250 = int32(0)
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(v1248)))
	if v1251 == v1250 {
		goto L283
	} else {
		goto L284
	}
L11:
	;
	v60 = int32(0)
	goto L13
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1215 = m.ExcPending
	if v1215 != 0 {
		goto L7
	} else {
		goto L278
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
	v1182 = m.ExcPending
	if v1182 != 0 {
		goto L7
	} else {
		goto L275
	}
L15:
	;
	goto L14
L16:
	;
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v38+v1170<<(uint(int32(2))%32)))) = v1153
	v1176 = v60 + int32(1)
	v1177 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v1176 < v1177 {
		v60 = v1176
		goto L13
	} else {
		goto L274
	}
L17:
	;
	v1001 = F_palloc0(m, int32(8))
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L7
	} else {
		goto L240
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L7
	} else {
		goto L237
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
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	v262 = v257 + v258*v242*int32(28)
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	switch v263 - int32(104) {
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
	v242 = v79
	v243 = int32(0)
	goto L26
L28:
	;
	goto L29
L29:
	;
	v83 = int32(0)
	v87 = v75
	v88 = v83
	v89 = v78
	v91 = v83
	goto L30
L30:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	v106 = F_bms_is_member(m, v88, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L7
	} else {
		goto L33
	}
L31:
	;
	v242 = v235
	v243 = v228
	goto L26
L32:
	;
	v234 = v88 + int32(1)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v234 < v235 {
		v87 = v225
		v88 = v234
		v89 = v226
		v91 = v228
		goto L30
	} else {
		goto L63
	}
L33:
	;
	if v106 != 0 {
		v225 = v87
		v226 = v89
		v228 = v91
		goto L32
	} else {
		goto L34
	}
L34:
	;
	if v88 <= v91 {
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
	v242 = v112
	v243 = v91
	goto L26
L38:
	;
	v225 = int32(0)
	v226 = v89
	v228 = v91
	goto L32
L39:
	;
	goto L40
L40:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v119 = v116*v117 + v88
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	if v121 == int32(7) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v167 = v164 + v119*int32(28)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	if v163 == v168 {
		goto L50
	} else {
		goto L51
	}
L42:
	;
	v154 = F_palloc(m, int32(8))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
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
	v161 = v124
	goto L41
L47:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_get_matching_partitions[0])) = v133
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+192)))
	if v145&int32(1) == int32(0) {
		v161 = v141
		goto L41
	} else {
		goto L48
	}
L48:
	;
	goto L42
L49:
	;
	v156 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v154)+4)) = uint16(v156)
	*(*int32)(unsafe.Add(mBase, uint32(v154))) = v156
	v1153 = v154
	goto L16
L50:
	;
	v197 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v23-int32(-64)+v88<<(uint(v197)%32)))) = v161
	v202 = v89 + int32(4)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)+12))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	if base.Ui32(v202) < base.Ui32(v205+v206<<(uint(v197)%32)) {
		goto L57
	} else {
		goto L58
	}
L51:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v174 = v171 + v88*int32(28)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)+4))
	if v175 == v163 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v179 = v167 + int32(16)
	v180 = *(*int64)(unsafe.Add(mBase, uint32(v174)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v179))) = v180
	v182 = *(*int64)(unsafe.Add(mBase, uint32(v174)))
	*(*int64)(unsafe.Add(mBase, uint32(v167))) = v182
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v174)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v167)+24)) = v184
	v186 = *(*int64)(unsafe.Add(mBase, uint32(v174)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v167)+8)) = v186
	*(*int32)(unsafe.Add(mBase, uint32(v167)+20)) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v179))) = int32(0)
	goto L55
L53:
	;
	goto L54
L54:
	;
	F_fmgr_info_cxt(m, v163, v167, v170)
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
	v213 = v87 + int32(4)
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
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
	v225 = v222
	v226 = v211
	v228 = v91 + int32(1)
	goto L32
L63:
	;
	goto L31
L64:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L7
	} else {
		goto L234
	}
L65:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		goto L7
	} else {
		goto L231
	}
L66:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	v561 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+8)))
	v563 = F_palloc0(m, int32(8))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L7
	} else {
		goto L142
	}
L67:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	v433 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+8)))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v436 = F_palloc0(m, int32(8))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L7
	} else {
		goto L99
	}
L68:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	v268 = F_palloc0(m, int32(8))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L7
	} else {
		goto L69
	}
L69:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v270)+24))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v274 = int32(0)
	if v266 == v274 {
		goto L73
	} else {
		goto L74
	}
L70:
	;
	v430 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v268)+4)) = uint16(v430)
	v1153 = v268
	goto L16
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v268))) = v408
	goto L70
L72:
	;
	if v273 == v309+v243 {
		goto L85
	} else {
		goto L86
	}
L73:
	;
	v309 = int32(0)
	goto L72
L74:
	;
	goto L75
L75:
	;
	v281 = int32(1)
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v266)+4))
	if v282 <= v281 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v285 = v281
	goto L78
L77:
	;
	v285 = v282
	goto L78
L78:
	;
	v289 = int32(0)
	v291 = v274
	goto L79
L79:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v266+int32(8)+v289<<(uint(int32(2))%32))))
	if v297 != 0 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v309 = v300
	goto L72
L81:
	;
	v300 = v291 + base.I32_popcnt(v297)
	goto L83
L82:
	;
	v300 = v291
	goto L83
L83:
	;
	v302 = v289 + int32(1)
	if v302 != v285 {
		v289 = v302
		v291 = v300
		goto L79
	} else {
		goto L84
	}
L84:
	;
	goto L80
L85:
	;
	v312 = int32(0)
	if v312 < v273 {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	goto L87
L87:
	;
	v381 = int32(0)
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
	v386 = F_bms_add_range(m, v381, v381, v383-int32(1))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L7
	} else {
		goto L98
	}
L88:
	;
	v317 = v312
	goto L91
L89:
	;
	goto L90
L90:
	;
	v368 = F_compute_partition_hash_value(m, v273, v262, v272, v23-int32(-64), v23+int32(192))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L7
	} else {
		goto L95
	}
L91:
	;
	v338 = F_bms_is_member(m, v317, v266)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L7
	} else {
		goto L93
	}
L92:
	;
	goto L90
L93:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v23+int32(192)+v317))) = uint8(v338)
	v342 = v317 + int32(1)
	if v342 != v273 {
		v317 = v342
		goto L91
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	v370 = int64(*(*int32)(unsafe.Add(mBase, uint32(v270)+20)))
	v371 = base.I64_rem_u_s(v368, v370)
	v372 = base.I32_wrap_i64(v371)
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v271+v372<<(uint(int32(2))%32))))
	if v376 < int32(0) {
		goto L70
	} else {
		goto L96
	}
L96:
	;
	v379 = F_bms_make_singleton(m, v372)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L7
	} else {
		goto L97
	}
L97:
	;
	v408 = v379
	goto L71
L98:
	;
	v408 = v386
	goto L71
L99:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v440 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v436)+4)) = uint16(v440)
	if v432 != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v438)+28))
	if v442 != int32(-1) {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	goto L102
L102:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v438)+4))
	if v451 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L103:
	;
	v445 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v436)+5)) = uint8(v445)
	v1153 = v436
	goto L16
L104:
	;
	goto L105
L105:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v438)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v436)+4)) = uint8(base.B2i32(v447 != int32(-1)))
	v1153 = v436
	goto L16
L106:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v438)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v436)+4)) = uint8(base.B2i32(v454 != int32(-1)))
	v1153 = v436
	goto L16
L107:
	;
	goto L108
L108:
	;
	v459 = v451 - int32(1)
	if v243 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v462 = int32(0)
	v464 = F_bms_add_range(m, v462, v462, v459)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L7
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	switch v433 {
	case 0:
		goto L120
	default:
		goto L119
	case 3:
		goto L118
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v436))) = v464
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v438)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v436)+4)) = uint8(base.B2i32(v467 != int32(-1)))
	v1153 = v436
	goto L16
L113:
	;
	v557 = F_bms_add_range(m, int32(0), v554, v553)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L7
	} else {
		goto L141
	}
L114:
	;
	v540 = F_partition_list_bsearch(m, v262, v439, v438, v434, v23+int32(192))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L7
	} else {
		goto L137
	}
L115:
	;
	v537 = int32(0)
	goto L114
L116:
	;
	v521 = F_partition_list_bsearch(m, v262, v439, v438, v434, v23+int32(192))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L7
	} else {
		goto L132
	}
L117:
	;
	v518 = int32(1)
	goto L116
L118:
	;
	v503 = F_partition_list_bsearch(m, v262, v439, v438, v434, v23+int32(192))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L7
	} else {
		goto L128
	}
L119:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v438)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v436)+4)) = uint8(base.B2i32(v493 != int32(-1)))
	v497 = int32(0)
	v498 = int32(1)
	switch v433 - v498 {
	case 0:
		v537 = v498
		goto L114
	case 1:
		goto L115
	default:
		goto L64
	case 3:
		goto L117
	case 4:
		v518 = v497
		goto L116
	}
L120:
	;
	v471 = int32(0)
	v473 = F_bms_add_range(m, v471, v471, v459)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L7
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v436))) = v473
	v478 = F_partition_list_bsearch(m, v262, v439, v438, v434, v23+int32(192))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L7
	} else {
		goto L123
	}
L122:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v438)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v436)+4)) = uint8(base.B2i32(v489 != int32(-1)))
	v1153 = v436
	goto L16
L123:
	;
	if v478 < int32(0) {
		goto L122
	} else {
		goto L124
	}
L124:
	;
	v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+192)))
	if v482 != int32(1) {
		goto L122
	} else {
		goto L125
	}
L125:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v436)))
	v486 = F_bms_del_member(m, v485, v478)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L7
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v436))) = v486
	goto L122
L127:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v438)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v436)+4)) = uint8(base.B2i32(v513 != int32(-1)))
	v1153 = v436
	goto L16
L128:
	;
	if v503 < int32(0) {
		goto L127
	} else {
		goto L129
	}
L129:
	;
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+192)))
	if v507 != int32(1) {
		goto L127
	} else {
		goto L130
	}
L130:
	;
	v510 = F_bms_make_singleton(m, v503)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L7
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v436))) = v510
	v1153 = v436
	goto L16
L132:
	;
	v523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+192)))
	v528 = int32(0)
	if v528 <= v521 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v531 = v521 + (v518&v523 ^ int32(1))
	goto L135
L134:
	;
	v531 = v528
	goto L135
L135:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v438)+4))
	if v531 <= v532-int32(1) {
		v553 = v459
		v554 = v531
		goto L113
	} else {
		goto L136
	}
L136:
	;
	v1153 = v436
	goto L16
L137:
	;
	if v540 < int32(0) {
		v1153 = v436
		goto L16
	} else {
		goto L138
	}
L138:
	;
	v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+192)))
	if v537&v544 != int32(1) {
		v553 = v540
		v554 = v497
		goto L113
	} else {
		goto L139
	}
L139:
	;
	if v540 == int32(0) {
		v1153 = v436
		goto L16
	} else {
		goto L140
	}
L140:
	;
	v553 = v540 - int32(1)
	v554 = v497
	goto L113
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v436))) = v557
	v1153 = v436
	goto L16
L142:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v566 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v567 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v567)+24))
	v569 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v563)+4)) = uint16(v569)
	if v560 == v569 {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	if v243 == int32(0) {
		goto L148
	} else {
		goto L149
	}
L144:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v567)+4))
	if v573 != 0 {
		goto L143
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v567)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v563)+4)) = uint8(base.B2i32(v575 != int32(-1)))
	v1153 = v563
	goto L16
L147:
	;
	goto L146
L148:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v568+v573<<(uint(int32(2))%32))))
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v568)))
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v567)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v563)+4)) = uint8(base.B2i32(v586 != int32(-1)))
	v591 = int32(31)
	v596 = F_bms_add_range(m, int32(0), int32(base.Ui32(v585)>>(uint(v591)%32)), v584>>(uint(v591)%32)+v573)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L7
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	v599 = base.B2i32(v565 <= v243)
	if v599 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v563))) = v596
	v1153 = v563
	goto L16
L152:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v567)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v563)+4)) = uint8(base.B2i32(v602 != int32(-1)))
	goto L154
L153:
	;
	goto L154
L154:
	;
	v606 = int32(0)
	switch v561 - int32(1) {
	case 0:
		v800 = v606
		goto L158
	case 1:
		goto L159
	case 2:
		goto L162
	case 3:
		goto L161
	case 4:
		v736 = v606
		goto L160
	default:
		goto L157
	}
L155:
	;
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v567)+4))
	if v904 <= v897 {
		v925 = v897
		goto L223
	} else {
		goto L224
	}
L156:
	;
	v888 = v573
	v897 = v741 + int32(1)
	goto L155
L157:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L7
	} else {
		goto L220
	}
L158:
	;
	v805 = F_partition_range_datum_bsearch(m, v262, v566, v567, v243, v23-int32(-64), v23+int32(192))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L7
	} else {
		goto L203
	}
L159:
	;
	v800 = int32(1)
	goto L158
L160:
	;
	v741 = F_partition_range_datum_bsearch(m, v262, v566, v567, v243, v23-int32(-64), v23+int32(192))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L7
	} else {
		goto L189
	}
L161:
	;
	v736 = int32(1)
	goto L160
L162:
	;
	v615 = F_partition_range_datum_bsearch(m, v262, v566, v567, v243, v23-int32(-64), v23+int32(192))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L7
	} else {
		goto L164
	}
L163:
	;
	v732 = F_bms_make_singleton(m, v615+int32(1))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L7
	} else {
		goto L188
	}
L164:
	;
	if v615 < int32(0) {
		goto L163
	} else {
		goto L165
	}
L165:
	;
	v619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+192)))
	if v619 != int32(1) {
		goto L163
	} else {
		goto L166
	}
L166:
	;
	if v243 == v565 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v625 = F_bms_make_singleton(m, v615+int32(1))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L7
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	v640 = v615
	goto L171
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v563))) = v625
	v1153 = v563
	goto L16
L171:
	;
	if v640 <= int32(0) {
		goto L174
	} else {
		goto L175
	}
L172:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v567)+12))
	v671 = int32(2)
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v670+v668<<(uint(v671)%32))))
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v674+v243<<(uint(v671)%32))))
	v686 = v615
	goto L179
L173:
	;
	goto L172
L174:
	;
	v668 = int32(0)
	goto L173
L175:
	;
	goto L176
L176:
	;
	v652 = v640 - int32(1)
	v654 = v652 << (uint(int32(2)) % 32)
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v567)+8))
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v654+v655)))
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v567)+12))
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v658+v654)))
	v663 = F_partition_rbound_datum_cmp(m, v262, v566, v657, v660, v23-int32(-64), v243)
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L7
	} else {
		goto L177
	}
L177:
	;
	if v663 == int32(0) {
		v640 = v652
		goto L171
	} else {
		goto L178
	}
L178:
	;
	v668 = v640
	goto L173
L179:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v567)+4))
	if v702-int32(1) <= v686 {
		goto L182
	} else {
		goto L183
	}
L180:
	;
	v727 = F_bms_add_range(m, int32(0), v668+base.B2i32(v678 == int32(-1)), v724)
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L7
	} else {
		goto L187
	}
L181:
	;
	goto L180
L182:
	;
	v724 = v686 + int32(1)
	goto L181
L183:
	;
	goto L184
L184:
	;
	v709 = v686 + int32(1)
	v711 = v709 << (uint(int32(2)) % 32)
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v567)+8))
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v711+v712)))
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v567)+12))
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v715+v711)))
	v720 = F_partition_rbound_datum_cmp(m, v262, v566, v714, v717, v23-int32(-64), v243)
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L7
	} else {
		goto L185
	}
L185:
	;
	if v720 == int32(0) {
		v686 = v709
		goto L179
	} else {
		goto L186
	}
L186:
	;
	v724 = v709
	goto L181
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v563))) = v727
	v1153 = v563
	goto L16
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v563))) = v732
	v1153 = v563
	goto L16
L189:
	;
	if v741 < int32(0) {
		v888 = v573
		v897 = v606
		goto L155
	} else {
		goto L190
	}
L190:
	;
	if v565 <= v243 {
		goto L156
	} else {
		goto L191
	}
L191:
	;
	v745 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+192)))
	if v745&int32(1) == int32(0) {
		goto L156
	} else {
		goto L192
	}
L192:
	;
	if v736 != 0 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v752 = int32(-1)
	goto L195
L194:
	;
	v752 = int32(1)
	goto L195
L195:
	;
	v765 = v741
	goto L196
L196:
	;
	if v765 <= int32(0) {
		goto L198
	} else {
		goto L199
	}
L197:
	;
	v888 = v573
	v897 = v765 + base.B2i32(v736 == int32(0))
	goto L155
L198:
	;
	goto L197
L199:
	;
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v567)+4))
	if v775-int32(1) <= v765 {
		goto L198
	} else {
		goto L200
	}
L200:
	;
	v779 = v765 + v752
	v781 = v779 << (uint(int32(2)) % 32)
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v567)+8))
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v781+v782)))
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v567)+12))
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v785+v781)))
	v790 = F_partition_rbound_datum_cmp(m, v262, v566, v784, v787, v23-int32(-64), v243)
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L7
	} else {
		goto L201
	}
L201:
	;
	if v790 == int32(0) {
		v765 = v779
		goto L196
	} else {
		goto L202
	}
L202:
	;
	goto L198
L203:
	;
	if int32(0) <= v805 {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v809 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+192)))
	v811 = v809 ^ int32(1)
	if v565 <= v243 {
		goto L207
	} else {
		goto L208
	}
L205:
	;
	goto L206
L206:
	;
	v888 = v805 + int32(1)
	v897 = v606
	goto L155
L207:
	;
	v888 = v805 + (v811|v800)&int32(1)
	v897 = v606
	goto L155
L208:
	;
	if v811&int32(1) != 0 {
		goto L207
	} else {
		goto L209
	}
L209:
	;
	if v800 != 0 {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v816 = int32(1)
	goto L212
L211:
	;
	v816 = int32(-1)
	goto L212
L212:
	;
	v826 = v805
	goto L213
L213:
	;
	if v826 <= int32(0) {
		goto L215
	} else {
		goto L216
	}
L214:
	;
	v888 = v826 + v800
	v897 = v606
	goto L155
L215:
	;
	goto L214
L216:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v567)+4))
	if v839-int32(1) <= v826 {
		goto L215
	} else {
		goto L217
	}
L217:
	;
	v843 = v826 + v816
	v845 = v843 << (uint(int32(2)) % 32)
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v567)+8))
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v845+v846)))
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v567)+12))
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v849+v845)))
	v854 = F_partition_rbound_datum_cmp(m, v262, v566, v848, v851, v23-int32(-64), v243)
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L7
	} else {
		goto L218
	}
L218:
	;
	if v854 == int32(0) {
		v826 = v843
		goto L213
	} else {
		goto L219
	}
L219:
	;
	goto L215
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v561
	F_errmsg_internal(m, int32(_a_F_get_matching_partitions_1), v23+int32(48))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L7
	} else {
		goto L221
	}
L221:
	;
	F_errfinish(m, int32(_a_F_get_matching_partitions_2), int32(3318), int32(_a_F_get_matching_partitions_3))
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L7
	} else {
		goto L222
	}
L222:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L223:
	;
	if v888 <= int32(0) {
		v948 = v888
		goto L226
	} else {
		goto L227
	}
L224:
	;
	v907 = v897 << (uint(int32(2)) % 32)
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v568+v907)))
	if int32(0) <= v909 {
		v925 = v897
		goto L223
	} else {
		goto L225
	}
L225:
	;
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v567)+12))
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v912+v907)))
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v914+v243<<(uint(int32(2))%32)-int32(4))))
	v925 = v897 + base.B2i32(v920 == int32(-1))
	goto L223
L226:
	;
	if v948 < v925 {
		v1153 = v563
		goto L16
	} else {
		goto L229
	}
L227:
	;
	v929 = v888 << (uint(int32(2)) % 32)
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v568+v929)))
	if int32(0) <= v931 {
		v948 = v888
		goto L226
	} else {
		goto L228
	}
L228:
	;
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v567)+12))
	v936 = int32(4)
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v934+v929-v936)))
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v938+v243<<(uint(int32(2))%32)-v936)))
	v948 = v888 - base.B2i32(v944 == int32(1))
	goto L226
L229:
	;
	v952 = F_bms_add_range(m, int32(0), v925, v948)
	mBase = m.M
	v953 = m.ExcPending
	if v953 != 0 {
		goto L7
	} else {
		goto L230
	}
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v563))) = v952
	v1153 = v563
	goto L16
L231:
	;
	v959 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v959
	F_errmsg_internal(m, int32(_a_F_get_matching_partitions_4), v23+int32(16))
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L7
	} else {
		goto L232
	}
L232:
	;
	F_errfinish(m, int32(_a_F_get_matching_partitions_2), int32(3575), int32(_a_F_get_matching_partitions_5))
	mBase = m.M
	v970 = m.ExcPending
	if v970 != 0 {
		goto L7
	} else {
		goto L233
	}
L233:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v433
	F_errmsg_internal(m, int32(_a_F_get_matching_partitions_1), v23+int32(32))
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L7
	} else {
		goto L235
	}
L235:
	;
	F_errfinish(m, int32(_a_F_get_matching_partitions_2), int32(2941), int32(_a_F_get_matching_partitions_6))
	mBase = m.M
	v985 = m.ExcPending
	if v985 != 0 {
		goto L7
	} else {
		goto L236
	}
L236:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L237:
	;
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v990
	F_errmsg_internal(m, int32(_a_F_get_matching_partitions_7), v23)
	mBase = m.M
	v994 = m.ExcPending
	if v994 != 0 {
		goto L7
	} else {
		goto L238
	}
L238:
	;
	F_errfinish(m, int32(_a_F_get_matching_partitions_2), int32(893), int32(_a_F_get_matching_partitions_8))
	mBase = m.M
	v999 = m.ExcPending
	if v999 != 0 {
		goto L7
	} else {
		goto L239
	}
L239:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L240:
	;
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	if v1003 == int32(0) {
		goto L241
	} else {
		goto L242
	}
L241:
	;
	v1006 = int32(0)
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v1008)+20))
	v1012 = F_bms_add_range(m, v1006, v1006, v1009-int32(1))
	mBase = m.M
	v1013 = m.ExcPending
	if v1013 != 0 {
		goto L7
	} else {
		goto L244
	}
L242:
	;
	goto L243
L243:
	;
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
	switch v1023 {
	case 0:
		goto L245
	case 1:
		goto L246
	default:
		v1153 = v1001
		goto L16
	}
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1001))) = v1012
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(v1008)+32))
	v1016 = int32(-1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1001)+4)) = uint8(base.B2i32(v1015 != v1016))
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(v1008)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v1001)+5)) = uint8(base.B2i32(v1019 != v1016))
	v1153 = v1001
	goto L16
L245:
	;
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v1003)+4))
	if v1096 <= int32(0) {
		v1153 = v1001
		goto L16
	} else {
		goto L262
	}
L246:
	;
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(v1003)+4))
	if v1024 <= int32(0) {
		v1153 = v1001
		goto L16
	} else {
		goto L247
	}
L247:
	;
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(v1003)+12))
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(v1027)))
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	if v1029 <= v1028 {
		goto L12
	} else {
		goto L248
	}
L248:
	;
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v38+v1028<<(uint(int32(2))%32))))
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(v1034)))
	v1036 = F_bms_copy(m, v1035)
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L7
	} else {
		goto L249
	}
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1001))) = v1036
	v1039 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1034)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1001)+5)) = uint8(v1039)
	v1041 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1034)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1001)+4)) = uint8(v1041)
	v1043 = int32(1)
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(v1003)+4))
	if v1044 <= v1043 {
		v1153 = v1001
		goto L16
	} else {
		goto L250
	}
L250:
	;
	v1052 = v1036
	v1053 = v1043
	goto L251
L251:
	;
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(v1003)+12))
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v1067+v1053<<(uint(int32(2))%32))))
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	if v1072 <= v1071 {
		goto L12
	} else {
		goto L253
	}
L252:
	;
	v1153 = v1001
	goto L16
L253:
	;
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v38+v1071<<(uint(int32(2))%32))))
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v1077)))
	v1079 = F_bms_int_members(m, v1052, v1078)
	mBase = m.M
	v1080 = m.ExcPending
	if v1080 != 0 {
		goto L7
	} else {
		goto L254
	}
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1001))) = v1079
	v1082 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1001)+5)))
	if v1082 == int32(1) {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	v1085 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1077)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1001)+5)) = uint8(v1085)
	goto L257
L256:
	;
	goto L257
L257:
	;
	v1087 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1001)+4)))
	if v1087 == int32(1) {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	v1090 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1077)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1001)+4)) = uint8(v1090)
	goto L260
L259:
	;
	goto L260
L260:
	;
	v1093 = v1053 + int32(1)
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v1003)+4))
	if v1093 < v1094 {
		v1052 = v1079
		v1053 = v1093
		goto L251
	} else {
		goto L261
	}
L261:
	;
	goto L252
L262:
	;
	v1106 = int32(0)
	goto L263
L263:
	;
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(v1003)+12))
	v1124 = *(*int32)(unsafe.Add(mBase, uint32(v1120+v1106<<(uint(int32(2))%32))))
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	if v1125 <= v1124 {
		goto L15
	} else {
		goto L265
	}
L264:
	;
	v1153 = v1001
	goto L16
L265:
	;
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(v1001)))
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(v38+v1124<<(uint(int32(2))%32))))
	v1132 = *(*int32)(unsafe.Add(mBase, uint32(v1131)))
	v1133 = F_bms_add_members(m, v1127, v1132)
	mBase = m.M
	v1134 = m.ExcPending
	if v1134 != 0 {
		goto L7
	} else {
		goto L266
	}
L266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1001))) = v1133
	v1136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1001)+5)))
	if v1136 == int32(0) {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	v1139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1131)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1001)+5)) = uint8(v1139)
	goto L269
L268:
	;
	goto L269
L269:
	;
	v1141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1001)+4)))
	if v1141 == int32(0) {
		goto L270
	} else {
		goto L271
	}
L270:
	;
	v1144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1131)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1001)+4)) = uint8(v1144)
	goto L272
L271:
	;
	goto L272
L272:
	;
	v1147 = v1106 + int32(1)
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(v1003)+4))
	if v1147 < v1148 {
		v1106 = v1147
		goto L263
	} else {
		goto L273
	}
L273:
	;
	goto L264
L274:
	;
	goto L10
L275:
	;
	F_errmsg_internal(m, int32(_a_F_get_matching_partitions_9), int32(0))
	mBase = m.M
	v1186 = m.ExcPending
	if v1186 != 0 {
		goto L7
	} else {
		goto L276
	}
L276:
	;
	F_errfinish(m, int32(_a_F_get_matching_partitions_2), int32(3630), int32(_a_F_get_matching_partitions_10))
	mBase = m.M
	v1191 = m.ExcPending
	if v1191 != 0 {
		goto L7
	} else {
		goto L277
	}
L277:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L278:
	;
	F_errmsg_internal(m, int32(_a_F_get_matching_partitions_9), int32(0))
	mBase = m.M
	v1219 = m.ExcPending
	if v1219 != 0 {
		goto L7
	} else {
		goto L279
	}
L279:
	;
	F_errfinish(m, int32(_a_F_get_matching_partitions_2), int32(3654), int32(_a_F_get_matching_partitions_10))
	mBase = m.M
	v1224 = m.ExcPending
	if v1224 != 0 {
		goto L7
	} else {
		goto L280
	}
L280:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L281:
	;
	if int32(0) <= v1308 {
		goto L292
	} else {
		goto L293
	}
L282:
	;
	v1308 = base.I32_ctz(v1294) | v1295<<(uint(int32(5))%32)
	goto L281
L283:
	;
	v1308 = int32(-2)
	goto L281
L284:
	;
	v1261 = base.I32_div_s(int32(0), int32(32))
	v1262 = *(*int32)(unsafe.Add(mBase, uint32(v1251)+4))
	if v1262 <= v1261 {
		goto L283
	} else {
		goto L285
	}
L285:
	;
	v1265 = v1251 + int32(8)
	v1269 = *(*int32)(unsafe.Add(mBase, uint32(v1265+v1261<<(uint(int32(2))%32))))
	v1272 = v1269 & int32(-1)
	if v1272 != 0 {
		v1294 = v1272
		v1295 = v1261
		goto L282
	} else {
		goto L286
	}
L286:
	;
	v1274 = v1261 + int32(1)
	if v1274 == v1262 {
		goto L283
	} else {
		goto L287
	}
L287:
	;
	v1277 = v1274
	goto L288
L288:
	;
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(v1265+v1277<<(uint(int32(2))%32))))
	if v1284 != 0 {
		v1294 = v1284
		v1295 = v1277
		goto L282
	} else {
		goto L290
	}
L289:
	;
	goto L283
L290:
	;
	v1286 = v1277 + int32(1)
	if v1286 != v1262 {
		v1277 = v1286
		goto L288
	} else {
		goto L291
	}
L291:
	;
	goto L289
L292:
	;
	v1314 = v1308
	v1315 = v1249
	v1319 = v1250
	goto L295
L293:
	;
	v1410 = v1249
	v1414 = v1250
	goto L294
L294:
	;
	v1426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1248)+5)))
	if v1426 == int32(1) {
		goto L314
	} else {
		goto L315
	}
L295:
	;
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(v1331)+24))
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(v1332+v1314<<(uint(int32(2))%32))))
	if v1336 < int32(0) {
		goto L298
	} else {
		goto L299
	}
L296:
	;
	v1410 = v1345
	v1414 = v1346
	goto L294
L297:
	;
	v1347 = *(*int32)(unsafe.Add(mBase, uint32(v1248)))
	if v1347 == int32(0) {
		goto L304
	} else {
		goto L305
	}
L298:
	;
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(v1331)+32))
	v1345 = v1315 | base.B2i32(v1339 != int32(-1))
	v1346 = v1319
	goto L297
L299:
	;
	goto L300
L300:
	;
	v1343 = F_bms_add_member(m, v1319, v1336)
	mBase = m.M
	v1344 = m.ExcPending
	if v1344 != 0 {
		goto L7
	} else {
		goto L301
	}
L301:
	;
	v1345 = v1315
	v1346 = v1343
	goto L297
L302:
	;
	if int32(0) <= v1403 {
		v1314 = v1403
		v1315 = v1345
		v1319 = v1346
		goto L295
	} else {
		goto L313
	}
L303:
	;
	v1403 = base.I32_ctz(v1389) | v1390<<(uint(int32(5))%32)
	goto L302
L304:
	;
	v1403 = int32(-2)
	goto L302
L305:
	;
	v1354 = v1314 + int32(1)
	v1356 = base.I32_div_s(v1354, int32(32))
	v1357 = *(*int32)(unsafe.Add(mBase, uint32(v1347)+4))
	if v1357 <= v1356 {
		goto L304
	} else {
		goto L306
	}
L306:
	;
	v1360 = v1347 + int32(8)
	v1364 = *(*int32)(unsafe.Add(mBase, uint32(v1360+v1356<<(uint(int32(2))%32))))
	v1367 = v1364 & (int32(-1) << (uint(v1354) % 32))
	if v1367 != 0 {
		v1389 = v1367
		v1390 = v1356
		goto L303
	} else {
		goto L307
	}
L307:
	;
	v1369 = v1356 + int32(1)
	if v1369 == v1357 {
		goto L304
	} else {
		goto L308
	}
L308:
	;
	v1372 = v1369
	goto L309
L309:
	;
	v1379 = *(*int32)(unsafe.Add(mBase, uint32(v1360+v1372<<(uint(int32(2))%32))))
	if v1379 != 0 {
		v1389 = v1379
		v1390 = v1372
		goto L303
	} else {
		goto L311
	}
L310:
	;
	goto L304
L311:
	;
	v1381 = v1372 + int32(1)
	if v1381 != v1357 {
		v1372 = v1381
		goto L309
	} else {
		goto L312
	}
L312:
	;
	goto L310
L313:
	;
	goto L296
L314:
	;
	v1429 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1430 = *(*int32)(unsafe.Add(mBase, uint32(v1429)+28))
	v1431 = F_bms_add_member(m, v1414, v1430)
	mBase = m.M
	v1432 = m.ExcPending
	if v1432 != 0 {
		goto L7
	} else {
		goto L317
	}
L315:
	;
	v1433 = v1414
	goto L316
L316:
	;
	if v1410&int32(1) == int32(0) {
		v1462 = v1433
		goto L1
	} else {
		goto L318
	}
L317:
	;
	v1433 = v1431
	goto L316
L318:
	;
	v1438 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1439 = *(*int32)(unsafe.Add(mBase, uint32(v1438)+32))
	v1440 = F_bms_add_member(m, v1433, v1439)
	mBase = m.M
	v1441 = m.ExcPending
	if v1441 != 0 {
		goto L7
	} else {
		goto L319
	}
L319:
	;
	v1462 = v1440
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
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int64
	_ = v132
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v153 int64
	_ = v153
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
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
	v209 = m.ExcPending
	if v209 != 0 {
		goto L3
	} else {
		goto L61
	}
L2:
	;
	m.G0 = v13 + int32(176)
	return v201
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
	v201 = v16
	goto L2
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v30 == int32(0) {
		v49 = v29
		v50 = v30
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if l2 != l6 {
		goto L18
	} else {
		goto L19
	}
L11:
	;
	goto L10
L12:
	;
	if v29 != v30 {
		v49 = v29
		v50 = v30
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v34 = v26
	v35 = l3
	goto L14
L14:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+1)))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+1)))
	if v39 == int32(0) {
		v49 = v38
		v50 = v39
		goto L11
	} else {
		goto L16
	}
L15:
	;
	v49 = v38
	v50 = v39
	goto L11
L16:
	;
	v42 = int32(1)
	if v38 == v39 {
		v34 = v34 + v42
		v35 = v35 + v42
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_get_other_operator[0]))
	v59 = F_object_aclcheck(m, int32(2615), v24, v57, int64(512))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L3
	} else {
		goto L23
	}
L19:
	;
	if l1 != l5 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	if v50-v49 != 0 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	if v24 == l4 {
		v201 = int32(0)
		goto L2
	} else {
		goto L22
	}
L22:
	;
	goto L18
L23:
	;
	if v59 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v62 = F_get_namespace_name(m, v24)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L3
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	v67 = int32(0)
	v70 = F_strlen(m, v66)
	mBase = m.M
	if base.Ui32(v70+int32(-64)) < base.Ui32(int32(-63)) {
		v123 = v67
		goto L30
	} else {
		goto L31
	}
L27:
	;
	F_aclcheck_error(m, v59, int32(36), v62)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	if v123 == int32(0) {
		goto L1
	} else {
		goto L44
	}
L30:
	;
	goto L29
L31:
	;
	v76 = F_strspn(m, v66, int32(_a_F_get_other_operator_0))
	mBase = m.M
	if v76 != v70 {
		v123 = v67
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v79 = F_strstr(m, v66, int32(_a_F_get_other_operator_1))
	mBase = m.M
	if v79 != 0 {
		v123 = v67
		goto L30
	} else {
		goto L33
	}
L33:
	;
	v81 = F_strstr(m, v66, int32(_a_F_get_other_operator_2))
	mBase = m.M
	if v81 != 0 {
		v123 = v67
		goto L30
	} else {
		goto L34
	}
L34:
	;
	if base.Ui32(v70) < base.Ui32(int32(2)) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v111 = int32(1)
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if v112 != int32(33) {
		v123 = v111
		goto L30
	} else {
		goto L42
	}
L36:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66+v70-int32(1)))))
	switch v87 - int32(43) {
	case 0, 2:
		goto L37
	default:
		goto L35
	}
L37:
	;
	v93 = v70 - int32(2)
	goto L38
L38:
	;
	v98 = int32(*(*int8)(unsafe.Add(mBase, uint32(v66+v93))))
	v100 = F_memchr(m, int32(_a_F_get_other_operator_3), v98, int32(11))
	mBase = m.M
	if v100 != 0 {
		goto L35
	} else {
		goto L40
	}
L39:
	;
	v123 = v67
	goto L30
L40:
	;
	v101 = int32(0)
	if base.B2i32(v93 <= v101) == v101 {
		v93 = v93 - int32(1)
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
	if v115 != int32(61) {
		v123 = v111
		goto L30
	} else {
		goto L43
	}
L43:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+2)))
	v123 = base.B2i32(v118 != int32(0))
	goto L30
L44:
	;
	v129 = F_table_open(m, int32(2617), int32(3))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L3
	} else {
		goto L45
	}
L45:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)+52))
	v132 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+103)) = v132
	*(*int64)(unsafe.Add(mBase, uint32(v13)+96)) = v132
	v138 = F_GetNewOidWithIndex(m, v129, int32(2688), int32(1))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L3
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+112)) = v138
	v144 = F_strncpy(m, v13+int32(32), v66, int32(64))
	mBase = m.M
	v145 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v144)+63)) = uint8(v145)
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+120)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v13)+116)) = v13 + int32(32)
	v152 = *(*int32)(unsafe.Add(mBase, _c_F_get_other_operator[0]))
	v153 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+156)) = v153
	*(*int64)(unsafe.Add(mBase, uint32(v13)+164)) = v153
	*(*int64)(unsafe.Add(mBase, uint32(v13)+148)) = v153
	*(*int32)(unsafe.Add(mBase, uint32(v13)+144)) = l2
	*(*int64)(unsafe.Add(mBase, uint32(v13)+132)) = v153
	*(*int32)(unsafe.Add(mBase, uint32(v13)+124)) = v152
	*(*int32)(unsafe.Add(mBase, uint32(v13)+140)) = l1
	if l1 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v166 = int32(98)
	goto L50
L49:
	;
	v166 = int32(108)
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+128)) = v166
	v172 = F_heap_form_tuple(m, v131, v13+int32(112), v13+int32(96))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L3
	} else {
		goto L51
	}
L51:
	;
	F_CatalogTupleInsert(m, v129, v172)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L3
	} else {
		goto L52
	}
L52:
	;
	F_makeOperatorDependencies(m, v13+int32(20), v172, int32(1), int32(0))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L3
	} else {
		goto L53
	}
L53:
	;
	F_pfree(m, v172)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L3
	} else {
		goto L54
	}
L54:
	;
	v185 = *(*int32)(unsafe.Add(mBase, _c_F_get_other_operator[1]))
	if v185 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v187 = int32(0)
	F_RunObjectPostCreateHook(m, int32(2617), v138, v187, v187)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L3
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L3
	} else {
		goto L59
	}
L58:
	;
	goto L57
L59:
	;
	F_sequence_close(m, v129, int32(3))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L3
	} else {
		goto L60
	}
L60:
	;
	v201 = v138
	goto L2
L61:
	;
	F_errcode(m, int32(33579140))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L3
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v66
	F_errmsg(m, int32(_a_F_get_other_operator_4), v13)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L3
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_get_other_operator_5), int32(214), int32(_a_F_get_other_operator_6))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L3
	} else {
		goto L64
	}
L64:
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
		return int32(20)
	default:
		v15 = int32(41)
		return v15
	case 10:
		return int32(37)
	case 29:
		v15 = int32(18)
		return v15
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
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int64
	_ = v47
	var v49 int64
	_ = v49
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	v7 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	if l1 == v7 {
		v68 = v7
		m.G0 = v12 + int32(32)
		return v68
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		if v16 != int32(2) {
			v68 = v7
			m.G0 = v12 + int32(32)
			return v68
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
					if v29 == int32(0) {
						if v29 != 0 {
							v56 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
							if v56 != 0 {
								v57 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
								m.T0[v57].(func(*base.Module, int32))(m, v56)
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int32(0)
								} else {
									v60 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
									if v60 == int32(0) {
										v68 = v7
										m.G0 = v12 + int32(32)
										return v68
									} else {
										v63 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
										m.T0[v63].(func(*base.Module, int32))(m, v60)
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return int32(0)
										} else {
											v68 = v7
											m.G0 = v12 + int32(32)
											return v68
										}
									}
								}
							} else {
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
								if v60 == int32(0) {
									v68 = v7
									m.G0 = v12 + int32(32)
									return v68
								} else {
									v63 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
									m.T0[v63].(func(*base.Module, int32))(m, v60)
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return int32(0)
									} else {
										v68 = v7
										m.G0 = v12 + int32(32)
										return v68
									}
								}
							}
						} else {
							if v28 == int32(0) {
								v56 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
								if v56 != 0 {
									v57 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
									m.T0[v57].(func(*base.Module, int32))(m, v56)
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return int32(0)
									} else {
										v60 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
										if v60 == int32(0) {
											v68 = v7
											m.G0 = v12 + int32(32)
											return v68
										} else {
											v63 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
											m.T0[v63].(func(*base.Module, int32))(m, v60)
											mBase = m.M
											v65 = m.ExcPending
											if v65 != 0 {
												return int32(0)
											} else {
												v68 = v7
												m.G0 = v12 + int32(32)
												return v68
											}
										}
									}
								} else {
									v60 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
									if v60 == int32(0) {
										v68 = v7
										m.G0 = v12 + int32(32)
										return v68
									} else {
										v63 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
										m.T0[v63].(func(*base.Module, int32))(m, v60)
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return int32(0)
										} else {
											v68 = v7
											m.G0 = v12 + int32(32)
											return v68
										}
									}
								}
							} else {
								v41 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v41)
								v43 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
								v44 = F_estimate_expression_value(m, l0, v43)
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l4))) = v44
									v47 = *(*int64)(unsafe.Add(mBase, uint32(v12)+24))
									*(*int64)(unsafe.Add(mBase, uint32(l3)+24)) = v47
									v49 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
									*(*int64)(unsafe.Add(mBase, uint32(l3)+16)) = v49
									v51 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
									*(*int64)(unsafe.Add(mBase, uint32(l3)+8)) = v51
									v53 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
									*(*int64)(unsafe.Add(mBase, uint32(l3))) = v53
									v68 = int32(1)
									m.G0 = v12 + int32(32)
									return v68
								}
							}
						}
					} else {
						if v28 != 0 {
							if v29 != 0 {
								v56 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
								if v56 != 0 {
									v57 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
									m.T0[v57].(func(*base.Module, int32))(m, v56)
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return int32(0)
									} else {
										v60 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
										if v60 == int32(0) {
											v68 = v7
											m.G0 = v12 + int32(32)
											return v68
										} else {
											v63 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
											m.T0[v63].(func(*base.Module, int32))(m, v60)
											mBase = m.M
											v65 = m.ExcPending
											if v65 != 0 {
												return int32(0)
											} else {
												v68 = v7
												m.G0 = v12 + int32(32)
												return v68
											}
										}
									}
								} else {
									v60 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
									if v60 == int32(0) {
										v68 = v7
										m.G0 = v12 + int32(32)
										return v68
									} else {
										v63 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
										m.T0[v63].(func(*base.Module, int32))(m, v60)
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return int32(0)
										} else {
											v68 = v7
											m.G0 = v12 + int32(32)
											return v68
										}
									}
								}
							} else {
								if v28 == int32(0) {
									v56 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
									if v56 != 0 {
										v57 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
										m.T0[v57].(func(*base.Module, int32))(m, v56)
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return int32(0)
										} else {
											v60 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
											if v60 == int32(0) {
												v68 = v7
												m.G0 = v12 + int32(32)
												return v68
											} else {
												v63 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
												m.T0[v63].(func(*base.Module, int32))(m, v60)
												mBase = m.M
												v65 = m.ExcPending
												if v65 != 0 {
													return int32(0)
												} else {
													v68 = v7
													m.G0 = v12 + int32(32)
													return v68
												}
											}
										}
									} else {
										v60 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
										if v60 == int32(0) {
											v68 = v7
											m.G0 = v12 + int32(32)
											return v68
										} else {
											v63 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
											m.T0[v63].(func(*base.Module, int32))(m, v60)
											mBase = m.M
											v65 = m.ExcPending
											if v65 != 0 {
												return int32(0)
											} else {
												v68 = v7
												m.G0 = v12 + int32(32)
												return v68
											}
										}
									}
								} else {
									v41 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v41)
									v43 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
									v44 = F_estimate_expression_value(m, l0, v43)
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l4))) = v44
										v47 = *(*int64)(unsafe.Add(mBase, uint32(v12)+24))
										*(*int64)(unsafe.Add(mBase, uint32(l3)+24)) = v47
										v49 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
										*(*int64)(unsafe.Add(mBase, uint32(l3)+16)) = v49
										v51 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
										*(*int64)(unsafe.Add(mBase, uint32(l3)+8)) = v51
										v53 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
										*(*int64)(unsafe.Add(mBase, uint32(l3))) = v53
										v68 = int32(1)
										m.G0 = v12 + int32(32)
										return v68
									}
								}
							}
						} else {
							v32 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v32)
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
							v36 = F_estimate_expression_value(m, l0, v35)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l4))) = v36
								v68 = v32
								m.G0 = v12 + int32(32)
								return v68
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
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v13)+79)))
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
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
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
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
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
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
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
	v30 = int32(0)
	goto L3
L3:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	v36 = base.I32_extend8_s(v35)
	switch v30 - int32(1) {
	case 0:
		goto L12
	case 1:
		goto L11
	case 2:
		goto L10
	case 3:
		goto L9
	default:
		goto L13
	}
L5:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v329 + int32(1)
	v30 = v328
	goto L3
L6:
	;
	v328 = int32(2)
	goto L5
L7:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v296 = v294 - v295
	if v293 <= v296+int32(1) {
		goto L84
	} else {
		goto L85
	}
L8:
	;
	m.G0 = v12 + int32(16)
	return v286
L9:
	;
	if v36 != 0 {
		goto L7
	} else {
		goto L78
	}
L10:
	;
	if v36 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L11:
	;
	if v35 == int32(92) {
		v328 = int32(4)
		goto L5
	} else {
		goto L50
	}
L12:
	;
	if v35 == int32(92) {
		v328 = int32(3)
		goto L5
	} else {
		goto L33
	}
L13:
	;
	if v35 <= int32(60) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	goto L31
L15:
	;
	if l1 != 0 {
		goto L14
	} else {
		goto L24
	}
L16:
	;
	v49 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v49)
	goto L6
L17:
	;
	if v35 == int32(34) {
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if v35 == int32(61) {
		goto L15
	} else {
		goto L22
	}
L20:
	;
	if v35 != 0 {
		goto L14
	} else {
		goto L21
	}
L21:
	;
	v286 = int32(0)
	goto L8
L22:
	;
	if v35 != int32(92) {
		goto L14
	} else {
		goto L23
	}
L23:
	;
	v328 = int32(3)
	goto L5
L24:
	;
	v51 = int32(0)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v53 = F_errsave_start(m, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	if v53 == int32(0) {
		v286 = v51
		goto L8
	} else {
		goto L26
	}
L26:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v61 = F_pg_mblen_cstr(m, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v64 - v63
	F_errmsg(m, int32(_a_F_get_val_0), v12)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_errsave_finish(m, v52, int32(_a_F_get_val_1), int32(71), int32(_a_F_get_val_2))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v286 = v51
	goto L8
L31:
	;
	if base.B2i32(v36 == int32(32))|base.B2i32(base.Ui32((v36-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v328 = int32(0)
		goto L5
	} else {
		goto L32
	}
L32:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	*(*uint8)(unsafe.Add(mBase, uint32(v91))) = uint8(v93)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v96 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v95 + v96
	v328 = v96
	goto L5
L33:
	;
	if v35 == int32(61) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v117 = int32(1)
	goto L41
L35:
	;
	if l1 != 0 {
		goto L34
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	if l1 == int32(0) {
		goto L34
	} else {
		goto L39
	}
L38:
	;
	v105 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v34 - v105
	v286 = v105
	goto L8
L39:
	;
	if v36 != int32(44) {
		goto L34
	} else {
		goto L40
	}
L40:
	;
	v113 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v34 - v113
	v286 = v113
	goto L8
L41:
	;
	if base.B2i32(v36 == int32(32))|base.B2i32(base.Ui32((v36-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v286 = v117
		goto L8
	} else {
		goto L42
	}
L42:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	if v128 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v127 - int32(1)
	v286 = v117
	goto L8
L44:
	;
	goto L45
L45:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v137 = v135 - v136
	if v134 <= v137+int32(1) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v142 = v134 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v142
	v144 = F_repalloc(m, v136, v142)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L49
	}
L47:
	;
	v151 = v128
	v152 = v135
	goto L48
L48:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v152))) = uint8(v151)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v155 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v154 + v155
	v328 = v155
	goto L5
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v144
	v147 = v144 + v137
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v147
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	v151 = v150
	v152 = v147
	goto L48
L50:
	;
	if v35 == int32(34) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v286 = int32(1)
	goto L8
L52:
	;
	goto L53
L53:
	;
	if v35 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v167 = int32(0)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v169 = F_errsave_start(m, v168)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v192 = v190 - v191
	if v189 <= v192+int32(1) {
		goto L62
	} else {
		goto L63
	}
L57:
	;
	if v169 == int32(0) {
		v286 = v167
		goto L8
	} else {
		goto L58
	}
L58:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errmsg(m, int32(_a_F_get_val_3), int32(0))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_errsave_finish(m, v168, int32(_a_F_get_val_1), int32(83), int32(_a_F_get_val_4))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v286 = v167
	goto L8
L62:
	;
	v197 = v189 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v197
	v199 = F_repalloc(m, v191, v197)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L65
	}
L63:
	;
	v206 = v36
	v207 = v190
	goto L64
L64:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v207))) = uint8(v206)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v209 + int32(1)
	goto L6
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v199
	v202 = v199 + v192
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v202
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	v206 = v205
	v207 = v202
	goto L64
L66:
	;
	v215 = int32(0)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v217 = F_errsave_start(m, v216)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v240 = v238 - v239
	if v237 <= v240+int32(1) {
		goto L74
	} else {
		goto L75
	}
L69:
	;
	if v217 == int32(0) {
		v286 = v215
		goto L8
	} else {
		goto L70
	}
L70:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_errmsg(m, int32(_a_F_get_val_3), int32(0))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	F_errsave_finish(m, v216, int32(_a_F_get_val_1), int32(83), int32(_a_F_get_val_4))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v286 = v215
	goto L8
L74:
	;
	v245 = v237 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v245
	v247 = F_repalloc(m, v239, v245)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	v254 = v36
	v255 = v238
	goto L76
L76:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v255))) = uint8(v254)
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v258 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v257 + v258
	v328 = v258
	goto L5
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v247
	v250 = v247 + v240
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v250
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252))))
	v254 = v253
	v255 = v250
	goto L76
L78:
	;
	v262 = int32(0)
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v264 = F_errsave_start(m, v263)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	if v264 == int32(0) {
		v286 = v262
		goto L8
	} else {
		goto L80
	}
L80:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_errmsg(m, int32(_a_F_get_val_3), int32(0))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	F_errsave_finish(m, v263, int32(_a_F_get_val_1), int32(83), int32(_a_F_get_val_4))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v286 = v262
	goto L8
L84:
	;
	v301 = v293 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v301
	v303 = F_repalloc(m, v295, v301)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L87
	}
L85:
	;
	v310 = v36
	v311 = v294
	goto L86
L86:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v311))) = uint8(v310)
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v313 + int32(1)
	goto L6
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v303
	v306 = v303 + v296
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v306
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v308))))
	v310 = v309
	v311 = v306
	goto L86
}
func F_getrule(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v357 int32
	_ = v357
	var v368 int32
	_ = v368
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
	return v368
L2:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	if v198 == int32(47) {
		goto L47
	} else {
		goto L48
	}
L3:
	;
	if base.Ui32(int32(9)) < base.Ui32(base.I32_extend8_s(v7)-int32(48)) {
		goto L38
	} else {
		goto L39
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(2)
	v48 = l0 + int32(1)
	if v48 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L5:
	;
	v10 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v10
	v13 = l0 + int32(1)
	if v13 == v10 {
		v368 = v3
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v16 = int32(*(*int8)(unsafe.Add(mBase, uint32(v13))))
	if base.Ui32(int32(9)) < base.Ui32(v16-int32(48)) {
		v368 = v3
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v22 = v13
	v24 = int32(0)
	v25 = v16
	goto L8
L8:
	;
	v32 = v25 + v24*int32(10) - int32(48)
	if int32(365) < v32 {
		v368 = v3
		goto L1
	} else {
		goto L10
	}
L9:
	;
	if v32 <= int32(0) {
		v368 = v3
		goto L1
	} else {
		goto L12
	}
L10:
	;
	v36 = v22 + int32(1)
	v37 = int32(*(*int8)(unsafe.Add(mBase, uint32(v36))))
	if base.Ui32(v37-int32(48)) < base.Ui32(int32(10)) {
		v22 = v36
		v24 = v32
		v25 = v37
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v32
	v192 = v36
	goto L2
L13:
	;
	return int32(0)
L14:
	;
	goto L15
L15:
	;
	v53 = int32(*(*int8)(unsafe.Add(mBase, uint32(v48))))
	if base.Ui32(int32(9)) < base.Ui32(v53-int32(48)) {
		v368 = v3
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v59 = v48
	v61 = int32(0)
	v62 = v53
	goto L17
L17:
	;
	v69 = v62 + v61*int32(10) - int32(48)
	if int32(12) < v69 {
		v368 = v3
		goto L1
	} else {
		goto L19
	}
L18:
	;
	if v69 <= int32(0) {
		v368 = v3
		goto L1
	} else {
		goto L21
	}
L19:
	;
	v73 = v59 + int32(1)
	v74 = int32(*(*int8)(unsafe.Add(mBase, uint32(v73))))
	if base.Ui32(v74-int32(48)) < base.Ui32(int32(10)) {
		v59 = v73
		v61 = v69
		v62 = v74
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v69
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if v82 != int32(46) {
		v368 = v3
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v86 = v59 + int32(2)
	if v86 == int32(0) {
		v368 = v3
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v89 = int32(*(*int8)(unsafe.Add(mBase, uint32(v86))))
	if base.Ui32(int32(9)) < base.Ui32(v89-int32(48)) {
		v368 = v3
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v95 = v86
	v97 = int32(0)
	v98 = v89
	goto L25
L25:
	;
	v105 = v98 + v97*int32(10) - int32(48)
	if int32(5) < v105 {
		v368 = v3
		goto L1
	} else {
		goto L27
	}
L26:
	;
	if v105 <= int32(0) {
		v368 = v3
		goto L1
	} else {
		goto L29
	}
L27:
	;
	v109 = v95 + int32(1)
	v110 = int32(*(*int8)(unsafe.Add(mBase, uint32(v109))))
	if base.Ui32(v110-int32(48)) < base.Ui32(int32(10)) {
		v95 = v109
		v97 = v105
		v98 = v110
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v105
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	if v118 != int32(46) {
		v368 = v3
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v122 = v95 + int32(2)
	if v122 == int32(0) {
		v368 = v3
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v125 = int32(*(*int8)(unsafe.Add(mBase, uint32(v122))))
	if base.Ui32(int32(9)) < base.Ui32(v125-int32(48)) {
		v368 = v3
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v131 = v122
	v133 = int32(0)
	v134 = v125
	goto L33
L33:
	;
	v141 = v134 + v133*int32(10) - int32(48)
	if int32(6) < v141 {
		v368 = v3
		goto L1
	} else {
		goto L35
	}
L34:
	;
	if v141 < int32(0) {
		v368 = v3
		goto L1
	} else {
		goto L37
	}
L35:
	;
	v145 = v131 + int32(1)
	v146 = int32(*(*int8)(unsafe.Add(mBase, uint32(v145))))
	if base.Ui32(v146-int32(48)) < base.Ui32(int32(10)) {
		v131 = v145
		v133 = v141
		v134 = v146
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v141
	v192 = v145
	goto L2
L38:
	;
	return int32(0)
L39:
	;
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1)
	v163 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
	if base.Ui32(int32(9)) < base.Ui32(v163-int32(48)) {
		v368 = v3
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v169 = l0
	v171 = int32(0)
	v172 = v163
	goto L42
L42:
	;
	v179 = v172 + v171*int32(10) - int32(48)
	if int32(365) < v179 {
		v368 = v3
		goto L1
	} else {
		goto L44
	}
L43:
	;
	if v179 < int32(0) {
		v368 = v3
		goto L1
	} else {
		goto L46
	}
L44:
	;
	v183 = v169 + int32(1)
	v184 = int32(*(*int8)(unsafe.Add(mBase, uint32(v183))))
	if base.Ui32(v184-int32(48)) < base.Ui32(int32(10)) {
		v169 = v183
		v171 = v179
		v172 = v184
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v179
	v192 = v183
	goto L2
L47:
	;
	v201 = int32(1)
	v202 = v192 + v201
	v204 = l1 + int32(16)
	v205 = int32(0)
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202))))
	switch v212 - int32(43) {
	case 0:
		goto L54
	default:
		v222 = v202
		v223 = v212
		v225 = v205
		goto L52
	case 2:
		v216 = v201
		goto L53
	}
L48:
	;
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = int32(_a_F_getrule_0)
	v368 = v192
	goto L1
L50:
	;
	return v357
L51:
	;
	goto L50
L52:
	;
	if base.Ui32(int32(9)) < base.Ui32(base.I32_extend8_s(v223)-int32(48)) {
		v357 = v205
		goto L51
	} else {
		goto L56
	}
L53:
	;
	v218 = v192 + int32(2)
	if v218 == int32(0) {
		v357 = v205
		goto L51
	} else {
		goto L55
	}
L54:
	;
	v216 = int32(0)
	goto L53
L55:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218))))
	v222 = v218
	v223 = v221
	v225 = v216
	goto L52
L56:
	;
	v232 = v222
	v234 = v223
	v235 = int32(0)
	goto L57
L57:
	;
	v245 = base.I32_extend8_s(v234) + v235*int32(10) - int32(48)
	if int32(167) < v245 {
		v357 = v205
		goto L51
	} else {
		goto L59
	}
L58:
	;
	if v245 < int32(0) {
		v357 = v205
		goto L51
	} else {
		goto L61
	}
L59:
	;
	v249 = v232 + int32(1)
	v250 = int32(*(*int8)(unsafe.Add(mBase, uint32(v249))))
	if base.Ui32(v250-int32(48)) < base.Ui32(int32(10)) {
		v232 = v249
		v234 = v250
		v235 = v245
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v258 = v245 * int32(3600)
	*(*int32)(unsafe.Add(mBase, uint32(v204))) = v258
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249))))
	if v260 != int32(58) {
		v342 = v249
		v347 = v258
		goto L62
	} else {
		goto L63
	}
L62:
	;
	if v225 != 0 {
		goto L79
	} else {
		goto L80
	}
L63:
	;
	v264 = v232 + int32(2)
	if v264 == int32(0) {
		v357 = v205
		goto L51
	} else {
		goto L64
	}
L64:
	;
	v267 = int32(*(*int8)(unsafe.Add(mBase, uint32(v264))))
	if base.Ui32(int32(9)) < base.Ui32(v267-int32(48)) {
		v357 = v205
		goto L51
	} else {
		goto L65
	}
L65:
	;
	v273 = v264
	v275 = int32(0)
	v276 = v267
	goto L66
L66:
	;
	v286 = base.I32_extend8_s(v276) + v275*int32(10) - int32(48)
	if int32(59) < v286 {
		v357 = v205
		goto L51
	} else {
		goto L68
	}
L67:
	;
	if v286 < int32(0) {
		v357 = v205
		goto L51
	} else {
		goto L70
	}
L68:
	;
	v290 = v273 + int32(1)
	v291 = int32(*(*int8)(unsafe.Add(mBase, uint32(v290))))
	if base.Ui32(v291-int32(48)) < base.Ui32(int32(10)) {
		v273 = v290
		v275 = v286
		v276 = v291
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	v300 = v286*int32(60) + v258
	*(*int32)(unsafe.Add(mBase, uint32(v204))) = v300
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290))))
	if v302 != int32(58) {
		v342 = v290
		v347 = v300
		goto L62
	} else {
		goto L71
	}
L71:
	;
	v306 = v273 + int32(2)
	if v306 == int32(0) {
		v357 = v205
		goto L51
	} else {
		goto L72
	}
L72:
	;
	v309 = int32(*(*int8)(unsafe.Add(mBase, uint32(v306))))
	if base.Ui32(int32(9)) < base.Ui32(v309-int32(48)) {
		v357 = v205
		goto L51
	} else {
		goto L73
	}
L73:
	;
	v315 = v306
	v317 = int32(0)
	v318 = v309
	goto L74
L74:
	;
	v328 = base.I32_extend8_s(v318) + v317*int32(10) - int32(48)
	if int32(60) < v328 {
		v357 = v205
		goto L51
	} else {
		goto L76
	}
L75:
	;
	if v328 < int32(0) {
		v357 = v205
		goto L51
	} else {
		goto L78
	}
L76:
	;
	v332 = v315 + int32(1)
	v333 = int32(*(*int8)(unsafe.Add(mBase, uint32(v332))))
	if base.Ui32(v333-int32(48)) < base.Ui32(int32(10)) {
		v315 = v332
		v317 = v328
		v318 = v333
		goto L74
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	v340 = v328 + v300
	*(*int32)(unsafe.Add(mBase, uint32(v204))) = v340
	v342 = v332
	v347 = v340
	goto L62
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v204))) = int32(0) - v347
	goto L81
L80:
	;
	goto L81
L81:
	;
	v357 = v342
	goto L51
}
func F_ginarrayextract(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pg_detoast_datum_copy(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
		F_get_typlenbyvalalign(m, v20, v11+int32(14), v11+int32(13), v11+int32(12))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			v30 = int32(*(*int16)(unsafe.Add(mBase, uint32(v11)+14)))
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+13)))
			v32 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11)+12)))
			F_deconstruct_array(m, v14, v30, v31, v32, v11+int32(8), v11+int32(4), v11)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				*(*int32)(unsafe.Add(mBase, uint32(v19))) = v39
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v18))) = v41
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
				m.G0 = v11 + int32(16)
				return v43
			}
		}
	}
}
func F_ginbuild(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v106 int64
	_ = v106
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
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
	var v256 int32
	_ = v256
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
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v369 int64
	_ = v369
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
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
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
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
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
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
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
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
	var v511 int32
	_ = v511
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v525 int32
	_ = v525
	var v526 float64
	_ = v526
	var v528 float64
	_ = v528
	var v533 int32
	_ = v533
	var v534 float64
	_ = v534
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v560 int32
	_ = v560
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v577 int64
	_ = v577
	var v580 int64
	_ = v580
	var v584 int64
	_ = v584
	var v590 float64
	_ = v590
	var v594 int64
	_ = v594
	var v596 int64
	_ = v596
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v610 int32
	_ = v610
	var v616 int32
	_ = v616
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v636 int32
	_ = v636
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
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
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v703 int64
	_ = v703
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v786 int32
	_ = v786
	var v796 int32
	_ = v796
	var v804 float64
	_ = v804
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v844 int32
	_ = v844
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v878 int64
	_ = v878
	var v880 int32
	_ = v880
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v925 int32
	_ = v925
	var v929 int32
	_ = v929
	var v936 int32
	_ = v936
	var v939 int32
	_ = v939
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v952 int32
	_ = v952
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v957 int32
	_ = v957
	var v959 int32
	_ = v959
	var v966 int32
	_ = v966
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v975 int32
	_ = v975
	var v981 int32
	_ = v981
	var v988 int32
	_ = v988
	var v992 int32
	_ = v992
	var v995 int32
	_ = v995
	var v1001 int32
	_ = v1001
	var v1008 int32
	_ = v1008
	var v1012 int32
	_ = v1012
	var v1015 int32
	_ = v1015
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1029 int32
	_ = v1029
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1034 int32
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1053 int32
	_ = v1053
	var v1058 int32
	_ = v1058
	var v1070 int32
	_ = v1070
	var v1073 int32
	_ = v1073
	var v1080 int32
	_ = v1080
	var v1083 float64
	_ = v1083
	var v1087 int64
	_ = v1087
	var v1089 int64
	_ = v1089
	var v1092 int32
	_ = v1092
	var v1096 int32
	_ = v1096
	var v1099 int32
	_ = v1099
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1105 int32
	_ = v1105
	var v1113 int32
	_ = v1113
	var v1119 int32
	_ = v1119
	var v1123 int32
	_ = v1123
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1129 float64
	_ = v1129
	var v1133 int64
	_ = v1133
	var v1153 int64
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1162 int32
	_ = v1162
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1169 int32
	_ = v1169
	var v1170 int64
	_ = v1170
	var v1172 int32
	_ = v1172
	var v1183 int32
	_ = v1183
	var v1187 int32
	_ = v1187
	var v1190 int32
	_ = v1190
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1196 int32
	_ = v1196
	var v1204 int32
	_ = v1204
	var v1210 int32
	_ = v1210
	var v1215 int32
	_ = v1215
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1225 int32
	_ = v1225
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1230 int32
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1245 float64
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1250 int32
	_ = v1250
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1264 int32
	_ = v1264
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1286 int32
	_ = v1286
	var v1302 int32
	_ = v1302
	var v1304 int32
	_ = v1304
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1312 int32
	_ = v1312
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1359 float64
	_ = v1359
	var v1361 int32
	_ = v1361
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1366 int32
	_ = v1366
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1379 int32
	_ = v1379
	var v1382 int32
	_ = v1382
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1390 int32
	_ = v1390
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1395 float64
	_ = v1395
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1411 int32
	_ = v1411
	var v1416 int32
	_ = v1416
	v19 = m.G0
	v21 = v19 - int32(_a_F_ginbuild_0)
	m.G0 = v21
	v24 = F_RelationGetNumberOfBlocksInFork(m, l1, int32(0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v24 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_initGinState(m, v21+int32(16), l1)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
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
	v1404 = m.ExcPending
	if v1404 != 0 {
		goto L1
	} else {
		goto L289
	}
L6:
	;
	v36 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[0]))) = uint16(v36)
	v39 = v21 + int32(_a_F_ginbuild_1)
	v40 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[1]))) = v40
	*(*int64)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[2]))) = v40
	*(*int64)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[3]))) = v40
	*(*int64)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[4]))) = v40
	*(*int64)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[5]))) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[6]))) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[7]))) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[8]))) = v40
	*(*int64)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[9]))) = v40
	v66 = F_GinNewBuffer(m, l1)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v68 = F_GinNewBuffer(m, l1)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v70 = int32(_a_F_ginbuild_2)
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10])) = v72 + int32(1)
	if v66 < int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	F_MarkBufferDirty(m, v66)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L14
	}
L10:
	;
	v96 = int32(8)
	F_PageInit(m, v94, int32(_a_F_ginbuild_3), v96)
	mBase = m.M
	v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v94)+16)))
	v99 = v94 + v98
	*(*int32)(unsafe.Add(mBase, uint32(v99))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v99)+6)) = uint16(v96)
	v106 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v94-int32(-64)))) = v106
	*(*int64)(unsafe.Add(mBase, uint32(v94)+24)) = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v94)+32)) = v106
	*(*int64)(unsafe.Add(mBase, uint32(v94)+40)) = v106
	*(*int64)(unsafe.Add(mBase, uint32(v94)+48)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v94)+56)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v94)+72)) = int32(2)
	v120 = int32(80)
	*(*uint16)(unsafe.Add(mBase, uint32(v94)+12)) = uint16(v120)
	goto L9
L11:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[11]))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v80+(v66^int32(-1))<<(uint(int32(2))%32))))
	v94 = v86
	goto L10
L12:
	;
	goto L13
L13:
	;
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[12]))
	v94 = v88 + v66<<(uint(int32(13))%32) + int32(-8192)
	goto L10
L14:
	;
	v124 = int32(2)
	if v68 < int32(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	F_MarkBufferDirty(m, v68)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L20
	}
L16:
	;
	F_PageInit(m, v142, int32(_a_F_ginbuild_3), int32(8))
	mBase = m.M
	v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v142)+16)))
	v147 = v142 + v146
	*(*int32)(unsafe.Add(mBase, uint32(v147))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v147)+6)) = uint16(v124)
	goto L15
L17:
	;
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[11]))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v128+(v68^int32(-1))<<(uint(int32(2))%32))))
	v142 = v134
	goto L16
L18:
	;
	goto L19
L19:
	;
	v136 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[12]))
	v142 = v136 + v68<<(uint(int32(13))%32) + int32(-8192)
	goto L16
L20:
	;
	F_UnlockReleaseBuffer(m, v66)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_UnlockReleaseBuffer(m, v68)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v157 = int32(_a_F_ginbuild_2)
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10]))
	v160 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10])) = v159 - v160
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[2])))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[2]))) = v163 + v160
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[13]))
	v173 = F_AllocSetContextCreateInternal(m, v168, int32(_a_F_ginbuild_4), int32(0), int32(_a_F_ginbuild_3), int32(_a_F_ginbuild_5))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[14]))) = v173
	v177 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[13]))
	v182 = F_AllocSetContextCreateInternal(m, v177, int32(_a_F_ginbuild_6), int32(0), int32(_a_F_ginbuild_3), int32(_a_F_ginbuild_5))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[15]))) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[16]))) = v21 + int32(16)
	v189 = v21 + int32(_a_F_ginbuild_7)
	F_ginInitBA(m, v189)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v196 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[17]))
	if v196 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l2)+128))
	if v227 <= int32(0) {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	goto L26
L28:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ginbuild[18])))
	if v200 != int32(1) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v203 = int32(_a_F_ginbuild_2)
	v205 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10]))
	v206 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10])) = v205 + v206
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
	*(*int32)(unsafe.Add(mBase, uint32(v196))) = v209 + v206
	*(*int64)(unsafe.Add(mBase, uint32(v196+int32(80))+232)) = int64(2)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
	*(*int32)(unsafe.Add(mBase, uint32(v196))) = v217 + v206
	v223 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10])) = v223 - v206
	goto L27
L30:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[7])))
	if v469 != 0 {
		goto L96
	} else {
		goto L97
	}
L31:
	;
	v231 = v227 + int32(1)
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+121)))
	v234 = F_palloc0(m, int32(28))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v238 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[19]))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v238)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v238)+72)) = v239 + int32(1)
	goto L33
L33:
	;
	v246 = F_CreateParallelContext(m, int32(_a_F_ginbuild_8), int32(_a_F_ginbuild_9), v227)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	if v232 == int32(1) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v250 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	v254 = int32(_a_F_ginbuild_10)
	goto L37
L37:
	;
	v256 = F_table_parallelscan_estimate(m, l0, v254)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	v252 = F_RegisterSnapshot(m, v250)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v254 = v252
	goto L37
L40:
	;
	v258 = F_add_size(m, int32(64), v256)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v246)+36))
	v265 = F_add_size(m, v260, (v258+int32(31))&int32(-32))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v246)+36)) = v265
	v268 = F_tuplesort_estimate_shared(m, v231)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v246)+36))
	v275 = F_add_size(m, v270, (v268+int32(31))&int32(-32))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v246)+36)) = v275
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v246)+40))
	v280 = F_add_size(m, v278, int32(2))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v246)+40)) = v280
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v246)+36))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v246)+12))
	v286 = F_mul_size(m, int32(32), v285)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v292 = F_add_size(m, v283, (v286+int32(31))&int32(-32))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v246)+36)) = v292
	v295 = int32(1)
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v246)+40))
	v298 = F_add_size(m, v296, v295)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v246)+40)) = v298
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v246)+36))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v246)+12))
	v304 = F_mul_size(m, int32(128), v303)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v310 = F_add_size(m, v301, (v304+int32(31))&int32(-32))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v246)+36)) = v310
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v246)+40))
	v315 = F_add_size(m, v313, int32(1))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v246)+40)) = v315
	v319 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[20]))
	if v319 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v246)+36))
	v321 = F_strlen(m, v319)
	mBase = m.M
	v326 = F_add_size(m, v320, v321&int32(-32)+int32(32))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L55
	}
L53:
	;
	v336 = v295
	goto L54
L54:
	;
	F_InitializeParallelDSM(m, v246)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L1
	} else {
		goto L57
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v246)+36)) = v326
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v246)+40))
	v331 = F_add_size(m, v329, int32(1))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v246)+40)) = v331
	v336 = v321 + int32(1)
	goto L54
L57:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v246)+44))
	if v339 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v254)))
	switch v342 {
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
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v246)+52))
	v355 = F_shm_toc_allocate(m, v354, v258)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L66
	}
L61:
	;
	F_DestroyParallelContext(m, v246)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L1
	} else {
		goto L64
	}
L62:
	;
	F_UnregisterSnapshot(m, v254)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v349 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[19]))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v349)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v349)+72)) = v350 - int32(1)
	goto L65
L65:
	;
	goto L30
L66:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v355))) = v357
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v355)+12)) = v231
	*(*uint8)(unsafe.Add(mBase, uint32(v355)+8)) = uint8(v232)
	*(*int32)(unsafe.Add(mBase, uint32(v355)+4)) = v359
	v364 = v355 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v364)+8)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v364))) = int64(-4294967296)
	goto L67
L67:
	;
	v369 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v355)+40)) = v369
	*(*int64)(unsafe.Add(mBase, uint32(v355)+28)) = v369
	*(*int64)(unsafe.Add(mBase, uint32(v355)+48)) = v369
	F_table_parallelscan_initialize(m, l0, v355-int32(-64), v254)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v246)+52))
	v380 = F_shm_toc_allocate(m, v379, v268)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v246)+44))
	F_tuplesort_initialize_shared(m, v380, v231, v382)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v246)+52))
	F_shm_toc_insert(m, v385, int64(-5764607523034234879), v355)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v246)+52))
	F_shm_toc_insert(m, v389, int64(-5764607523034234878), v380)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v394 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[20]))
	if v394 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v246)+52))
	v396 = F_shm_toc_allocate(m, v395, v336)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v246)+52))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v246)+12))
	v410 = F_mul_size(m, int32(32), v409)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L82
	}
L76:
	;
	v399 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[20]))
	if v336 != 0 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v246)+52))
	F_shm_toc_insert(m, v402, int64(-5764607523034234877), v401)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L81
	}
L78:
	;
	v400 = F__emscripten_memcpy_bulkmem(m, v396, v399, v336)
	mBase = m.M
	v401 = v400
	goto L80
L79:
	;
	v401 = v396
	goto L80
L80:
	;
	goto L77
L81:
	;
	goto L75
L82:
	;
	v412 = F_shm_toc_allocate(m, v407, v410)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v246)+52))
	F_shm_toc_insert(m, v414, int64(-5764607523034234876), v412)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v246)+52))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v246)+12))
	v421 = F_mul_size(m, int32(128), v420)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v423 = F_shm_toc_allocate(m, v418, v421)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v246)+52))
	F_shm_toc_insert(m, v425, int64(-5764607523034234875), v423)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	F_LaunchParallelWorkers(m, v246)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234))) = v246
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v246)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v234)+24)) = v423
	*(*int32)(unsafe.Add(mBase, uint32(v234)+20)) = v412
	*(*int32)(unsafe.Add(mBase, uint32(v234)+16)) = v254
	*(*int32)(unsafe.Add(mBase, uint32(v234)+12)) = v380
	*(*int32)(unsafe.Add(mBase, uint32(v234)+8)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v234)+4)) = v432 + int32(1)
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v246)+20))
	if v441 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	F__brin_end_parallel(m, v234)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v447 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[21]))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[7]))) = v234
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v234)+8))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v234)+12))
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
	v454 = base.I32_div_s(v447, v453)
	F__gin_parallel_scan_and_build(m, v21+int32(16), v451, v452, l0, l1, v454, int32(1))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L1
	} else {
		goto L93
	}
L92:
	;
	goto L30
L93:
	;
	F_WaitForParallelWorkersToAttach(m, v246)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	goto L30
L95:
	;
	v1361 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[15])))
	F_MemoryContextDelete(m, v1361)
	mBase = m.M
	v1363 = m.ExcPending
	if v1363 != 0 {
		goto L1
	} else {
		goto L275
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
	v1234 = int32(0)
	v1243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(v1243)+140))
	v1245 = m.T0[v1244].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, l0, l1, l2, v1234, v1234, int32(1), v1234, int32(-1), int32(53), v21+int32(16), v1234)
	mBase = m.M
	v1246 = m.ExcPending
	if v1246 != 0 {
		goto L1
	} else {
		goto L260
	}
L99:
	;
	v473 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v471))) = uint8(v473)
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[7])))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v475)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v471)+4)) = v476
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[7])))
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[22]))) = v483
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[7])))
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v486)+8))
	v491 = v487 + int32(28)
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v486)+4))
	goto L101
L101:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v491)))
	*(*int32)(unsafe.Add(mBase, uint32(v491))) = int32(1)
	if v511 != 0 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v526 = *(*float64)(unsafe.Add(mBase, uint32(v487)+40))
	*(*float64)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[5]))) = v526
	v528 = *(*float64)(unsafe.Add(mBase, uint32(v487)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[9]))) = v528
	*(*int32)(unsafe.Add(mBase, uint32(v487)+28)) = int32(0)
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L1
	} else {
		goto L111
	}
L103:
	;
	F_s_lock(m, v491, int32(_a_F_ginbuild_11), int32(1141), int32(_a_F_ginbuild_12))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L1
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v487)+32))
	if v492 != v519 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	goto L105
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v491))) = int32(0)
	F_ConditionVariableSleep(m, v487+int32(16), int32(134217767))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
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
	v534 = *(*float64)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[5])))
	v539 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[17]))
	if v539 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[22])))
	F_tuplesort_performsort(m, v570)
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L1
	} else {
		goto L116
	}
L113:
	;
	goto L112
L114:
	;
	v543 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ginbuild[18])))
	if v543 != int32(1) {
		goto L113
	} else {
		goto L115
	}
L115:
	;
	v546 = int32(_a_F_ginbuild_2)
	v548 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10]))
	v549 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10])) = v548 + v549
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v539)))
	*(*int32)(unsafe.Add(mBase, uint32(v539))) = v552 + v549
	*(*int64)(unsafe.Add(mBase, uint32(v539+int32(80))+232)) = int64(5)
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v539)))
	*(*int32)(unsafe.Add(mBase, uint32(v539))) = v560 + v549
	v566 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10])) = v566 - v549
	goto L113
L116:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	v574 = F_GinBufferInit(m, v573)
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	v577 = *(*int64)(unsafe.Add(mBase, _c_F_ginbuild[23]))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[24]))) = v577
	v580 = *(*int64)(unsafe.Add(mBase, _c_F_ginbuild[25]))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[26]))) = v580
	v584 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[27]))) = v584
	*(*int64)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[28]))) = int64(6)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[29]))) = v584
	v590 = *(*float64)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[9])))
	if base.F64_lt(base.F64_abs(v590), float64(9.223372036854776e+18)) != 0 {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[30]))) = v596
	v600 = v21 + int32(_a_F_ginbuild_13)
	v602 = v21 + int32(_a_F_ginbuild_14)
	v603 = int32(0)
	v610 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[17]))
	if v610 == v603 {
		goto L123
	} else {
		goto L124
	}
L119:
	;
	v594 = base.I64_trunc_f64_s(v590)
	v596 = v594
	goto L118
L120:
	;
	goto L121
L121:
	;
	v596 = int64(-9223372036854775807 - 1)
	goto L118
L122:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[22])))
	v780 = F_tuplesort_getgintuple(m, v777, v21+int32(_a_F_ginbuild_14))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L1
	} else {
		goto L140
	}
L123:
	;
	goto L122
L124:
	;
	goto L125
L125:
	;
	v616 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ginbuild[18])))
	if v616&int32(1) == int32(0) {
		goto L123
	} else {
		goto L126
	}
L126:
	;
	v621 = int32(_a_F_ginbuild_2)
	v623 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10]))
	v624 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10])) = v623 + v624
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v610)))
	*(*int32)(unsafe.Add(mBase, uint32(v610))) = v627 + v624
	goto L128
L127:
	;
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v610)))
	v758 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v610))) = v757 + v758
	v761 = int32(_a_F_ginbuild_2)
	v763 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10])) = v763 - v758
	goto L123
L128:
	;
	v636 = v610 + int32(232)
	goto L129
L129:
	;
	v642 = int32(0)
	v645 = v603
	goto L132
L131:
	;
	goto L127
L132:
	;
	v651 = int32(2)
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v600+v645<<(uint(v651)%32))))
	v655 = int32(3)
	v661 = *(*int64)(unsafe.Add(mBase, uint32(v602+v645<<(uint(v655)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v636+v654<<(uint(v655)%32)))) = v661
	v664 = v645 | int32(1)
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v600+v664<<(uint(v651)%32))))
	v675 = *(*int64)(unsafe.Add(mBase, uint32(v602+v664<<(uint(v655)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v636+v668<<(uint(v655)%32)))) = v675
	v678 = v645 | v651
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v600+v678<<(uint(v651)%32))))
	v689 = *(*int64)(unsafe.Add(mBase, uint32(v602+v678<<(uint(v655)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v636+v682<<(uint(v655)%32)))) = v689
	v692 = v645 | v655
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v600+v692<<(uint(v651)%32))))
	v703 = *(*int64)(unsafe.Add(mBase, uint32(v602+v692<<(uint(v655)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v636+v696<<(uint(v655)%32)))) = v703
	v705 = int32(4)
	v708 = v642 + v705
	if v708 != int32(4) {
		v642 = v708
		v645 = v645 + v705
		goto L132
	} else {
		goto L134
	}
L133:
	;
	goto L131
L134:
	;
	goto L133
L139:
	;
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v574)+20))
	if v1154 != 0 {
		goto L236
	} else {
		goto L237
	}
L140:
	;
	if v780 == int32(0) {
		v1153 = int64(1)
		goto L139
	} else {
		goto L141
	}
L141:
	;
	v786 = v574 + int32(4)
	v796 = v780
	v804 = float64(0)
	goto L142
L142:
	;
	v808 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[31]))
	if v808 != 0 {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	v1129 = base.F64_add(v1083, float64(1))
	if base.F64_lt(base.F64_abs(v1129), float64(9.223372036854776e+18)) != 0 {
		goto L233
	} else {
		goto L234
	}
L144:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L1
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v574)+20))
	if v811 == int32(0) {
		goto L149
	} else {
		goto L150
	}
L147:
	;
	goto L146
L148:
	;
	F_GinBufferStoreTuple(m, v574, v796)
	mBase = m.M
	v1080 = m.ExcPending
	if v1080 != 0 {
		goto L1
	} else {
		goto L222
	}
L149:
	;
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v574)+24))
	if v891 < int32(1024) {
		goto L148
	} else {
		goto L172
	}
L150:
	;
	v814 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v796)+4)))
	v815 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v574))))
	if v814 != v815 {
		v852 = v811
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v855 = int32(_a_F_ginbuild_15)
	v856 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[13]))
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[14])))
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[13])) = v858
	v862 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v574))))
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v574)+4))
	v864 = int32(*(*int8)(unsafe.Add(mBase, uint32(v574)+2)))
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v574)+32))
	F_ginEntryInsert(m, v21+int32(16), v862, v863, v864, v865, v852, v39)
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L1
	} else {
		goto L166
	}
L152:
	;
	v817 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v796)+11)))
	v818 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574)+2)))
	if v817 != v818 {
		v852 = v811
		goto L151
	} else {
		goto L153
	}
L153:
	;
	if v817 != 0 {
		goto L149
	} else {
		goto L154
	}
L154:
	;
	v821 = v796 + int32(16)
	v822 = int32(1)
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v574)+4))
	v825 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574)+14)))
	if v825 == v822 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v821)))
	v829 = v828
	goto L157
L156:
	;
	v829 = v821
	goto L157
L157:
	;
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v574)+28))
	v831 = int32(36)
	v833 = v830 + v814*v831
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v833-int32(20))))
	v839 = m.T0[v838].(func(*base.Module, int32, int32, int32) int32)(m, v824, v829, v833-v831)
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	if v839 < int32(0) {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v844 = v822
	goto L161
L160:
	;
	v844 = int32(0) - v839
	goto L161
L161:
	;
	v847 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v833-int32(28)))))
	if v847 != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v848 = v844
	goto L164
L163:
	;
	v848 = v839
	goto L164
L164:
	;
	if v848 == int32(0) {
		goto L149
	} else {
		goto L165
	}
L165:
	;
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v574)+20))
	v852 = v851
	goto L151
L166:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[13])) = v856
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[14])))
	F_MemoryContextReset(m, v870)
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	v873 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574)+2)))
	if v873 != 0 {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v878 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v574)+20)) = v878
	v880 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v574)+2)) = uint8(v880)
	*(*uint16)(unsafe.Add(mBase, uint32(v574))) = uint16(v880)
	*(*int32)(unsafe.Add(mBase, uint32(v574+int32(11)))) = v880
	*(*int64)(unsafe.Add(mBase, uint32(v786))) = v878
	goto L148
L169:
	;
	v874 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574)+14)))
	if v874 != 0 {
		goto L168
	} else {
		goto L170
	}
L170:
	;
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v786)))
	F_pfree(m, v875)
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	goto L168
L172:
	;
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v574)+16))
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v796)+12))
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v574)+20))
	if v895+v896 < v894 {
		goto L148
	} else {
		goto L173
	}
L173:
	;
	v899 = int32(_a_F_ginbuild_15)
	v900 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[13]))
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[14])))
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[13])) = v902
	v906 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v574))))
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v574)+4))
	v908 = int32(*(*int8)(unsafe.Add(mBase, uint32(v574)+2)))
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v574)+32))
	F_ginEntryInsert(m, v21+int32(16), v906, v907, v908, v909, v891, v39)
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[13])) = v900
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[14])))
	F_MemoryContextReset(m, v914)
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v574)+32))
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v574)+24))
	v919 = int32(6)
	v921 = v917 + v918*v919
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v574)+20))
	v925 = (v922 - v918) * v919
	if v917 == v921 {
		goto L177
	} else {
		goto L178
	}
L176:
	;
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v574)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v574)+24)) = int32(0)
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(v574)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v574)+20)) = v1073 - v1070
	goto L148
L177:
	;
	goto L176
L178:
	;
	v929 = v917 + v925
	if base.Ui32(v921-v929) <= base.Ui32(int32(0)-v925<<(uint(int32(1))%32)) {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v936 = F___memcpy(m, v917, v921, v925)
	mBase = m.M
	goto L176
L180:
	;
	goto L181
L181:
	;
	v939 = (v917 ^ v921) & int32(3)
	if base.Ui32(v917) < base.Ui32(v921) {
		goto L184
	} else {
		goto L185
	}
L182:
	;
	if v1041 == int32(0) {
		goto L177
	} else {
		goto L218
	}
L183:
	;
	if base.Ui32(v1019) <= base.Ui32(int32(3)) {
		v1040 = v1018
		v1041 = v1019
		v1042 = v1020
		goto L182
	} else {
		goto L214
	}
L184:
	;
	if v939 != 0 {
		goto L187
	} else {
		goto L188
	}
L185:
	;
	goto L186
L186:
	;
	if v939 != 0 {
		v1001 = v925
		goto L197
	} else {
		goto L198
	}
L187:
	;
	v1040 = v921
	v1041 = v925
	v1042 = v917
	goto L182
L188:
	;
	goto L189
L189:
	;
	if v917&int32(3) == int32(0) {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v1018 = v921
	v1019 = v925
	v1020 = v917
	goto L183
L191:
	;
	goto L192
L192:
	;
	v946 = v921
	v947 = v925
	v948 = v917
	goto L193
L193:
	;
	if v947 == int32(0) {
		goto L177
	} else {
		goto L195
	}
L194:
	;
	v1018 = v955
	v1019 = v957
	v1020 = v959
	goto L183
L195:
	;
	v952 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v946))))
	*(*uint8)(unsafe.Add(mBase, uint32(v948))) = uint8(v952)
	v954 = int32(1)
	v955 = v946 + v954
	v957 = v947 - v954
	v959 = v948 + v954
	if v959&int32(3) != 0 {
		v946 = v955
		v947 = v957
		v948 = v959
		goto L193
	} else {
		goto L196
	}
L196:
	;
	goto L194
L197:
	;
	if v1001 == int32(0) {
		goto L177
	} else {
		goto L210
	}
L198:
	;
	if v929&int32(3) != 0 {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v966 = v925
	goto L202
L200:
	;
	v981 = v925
	goto L201
L201:
	;
	if base.Ui32(v981) <= base.Ui32(int32(3)) {
		v1001 = v981
		goto L197
	} else {
		goto L206
	}
L202:
	;
	if v966 == int32(0) {
		goto L177
	} else {
		goto L204
	}
L203:
	;
	v981 = v972
	goto L201
L204:
	;
	v972 = v966 - int32(1)
	v973 = v917 + v972
	v975 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v921+v972))))
	*(*uint8)(unsafe.Add(mBase, uint32(v973))) = uint8(v975)
	if v973&int32(3) != 0 {
		v966 = v972
		goto L202
	} else {
		goto L205
	}
L205:
	;
	goto L203
L206:
	;
	v988 = v981
	goto L207
L207:
	;
	v992 = v988 - int32(4)
	v995 = *(*int32)(unsafe.Add(mBase, uint32(v921+v992)))
	*(*int32)(unsafe.Add(mBase, uint32(v917+v992))) = v995
	if base.Ui32(int32(3)) < base.Ui32(v992) {
		v988 = v992
		goto L207
	} else {
		goto L209
	}
L208:
	;
	v1001 = v992
	goto L197
L209:
	;
	goto L208
L210:
	;
	v1008 = v1001
	goto L211
L211:
	;
	v1012 = v1008 - int32(1)
	v1015 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v921+v1012))))
	*(*uint8)(unsafe.Add(mBase, uint32(v917+v1012))) = uint8(v1015)
	if v1012 != 0 {
		v1008 = v1012
		goto L211
	} else {
		goto L213
	}
L212:
	;
	goto L177
L213:
	;
	goto L212
L214:
	;
	v1025 = v1018
	v1026 = v1019
	v1027 = v1020
	goto L215
L215:
	;
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(v1025)))
	*(*int32)(unsafe.Add(mBase, uint32(v1027))) = v1029
	v1031 = int32(4)
	v1032 = v1025 + v1031
	v1034 = v1027 + v1031
	v1036 = v1026 - v1031
	if base.Ui32(int32(3)) < base.Ui32(v1036) {
		v1025 = v1032
		v1026 = v1036
		v1027 = v1034
		goto L215
	} else {
		goto L217
	}
L216:
	;
	v1040 = v1032
	v1041 = v1036
	v1042 = v1034
	goto L182
L217:
	;
	goto L216
L218:
	;
	v1047 = v1040
	v1048 = v1041
	v1049 = v1042
	goto L219
L219:
	;
	v1051 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1049))) = uint8(v1051)
	v1053 = int32(1)
	v1058 = v1048 - v1053
	if v1058 != 0 {
		v1047 = v1047 + v1053
		v1048 = v1058
		v1049 = v1049 + v1053
		goto L219
	} else {
		goto L221
	}
L220:
	;
	goto L177
L221:
	;
	goto L220
L222:
	;
	v1083 = base.F64_add(v804, float64(1))
	if base.F64_lt(base.F64_abs(v1083), float64(9.223372036854776e+18)) != 0 {
		goto L224
	} else {
		goto L225
	}
L223:
	;
	v1092 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[17]))
	if v1092 == int32(0) {
		goto L228
	} else {
		goto L229
	}
L224:
	;
	v1087 = base.I64_trunc_f64_s(v1083)
	v1089 = v1087
	goto L223
L225:
	;
	goto L226
L226:
	;
	v1089 = int64(-9223372036854775807 - 1)
	goto L223
L227:
	;
	v1123 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[22])))
	v1126 = F_tuplesort_getgintuple(m, v1123, v21+int32(_a_F_ginbuild_14))
	mBase = m.M
	v1127 = m.ExcPending
	if v1127 != 0 {
		goto L1
	} else {
		goto L231
	}
L228:
	;
	goto L227
L229:
	;
	v1096 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ginbuild[18])))
	if v1096 != int32(1) {
		goto L228
	} else {
		goto L230
	}
L230:
	;
	v1099 = int32(_a_F_ginbuild_2)
	v1101 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10]))
	v1102 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10])) = v1101 + v1102
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(v1092)))
	*(*int32)(unsafe.Add(mBase, uint32(v1092))) = v1105 + v1102
	*(*int64)(unsafe.Add(mBase, uint32(v1092+int32(96))+232)) = v1089
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(v1092)))
	*(*int32)(unsafe.Add(mBase, uint32(v1092))) = v1113 + v1102
	v1119 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10])) = v1119 - v1102
	goto L228
L231:
	;
	if v1126 != 0 {
		v796 = v1126
		v804 = v1083
		goto L142
	} else {
		goto L232
	}
L232:
	;
	goto L143
L233:
	;
	v1133 = base.I64_trunc_f64_s(v1129)
	v1153 = v1133
	goto L139
L234:
	;
	goto L235
L235:
	;
	v1153 = int64(-9223372036854775807 - 1)
	goto L139
L236:
	;
	v1157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v574))))
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(v574)+4))
	v1159 = int32(*(*int8)(unsafe.Add(mBase, uint32(v574)+2)))
	v1160 = *(*int32)(unsafe.Add(mBase, uint32(v574)+32))
	F_ginEntryInsert(m, v21+int32(16), v1157, v1158, v1159, v1160, v1154, v39)
	mBase = m.M
	v1162 = m.ExcPending
	if v1162 != 0 {
		goto L1
	} else {
		goto L239
	}
L237:
	;
	goto L238
L238:
	;
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(v574)+32))
	if v1215 != 0 {
		goto L248
	} else {
		goto L249
	}
L239:
	;
	v1164 = v574 + int32(4)
	v1165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574)+2)))
	if v1165 != 0 {
		goto L240
	} else {
		goto L241
	}
L240:
	;
	v1170 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v574)+20)) = v1170
	v1172 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v574)+2)) = uint8(v1172)
	*(*uint16)(unsafe.Add(mBase, uint32(v574))) = uint16(v1172)
	*(*int32)(unsafe.Add(mBase, uint32(v1164)+7)) = v1172
	*(*int64)(unsafe.Add(mBase, uint32(v1164))) = v1170
	v1183 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[17]))
	if v1183 == v1172 {
		goto L245
	} else {
		goto L246
	}
L241:
	;
	v1166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574)+14)))
	if v1166 != 0 {
		goto L240
	} else {
		goto L242
	}
L242:
	;
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(v1164)))
	F_pfree(m, v1167)
	mBase = m.M
	v1169 = m.ExcPending
	if v1169 != 0 {
		goto L1
	} else {
		goto L243
	}
L243:
	;
	goto L240
L244:
	;
	goto L238
L245:
	;
	goto L244
L246:
	;
	v1187 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ginbuild[18])))
	if v1187 != int32(1) {
		goto L245
	} else {
		goto L247
	}
L247:
	;
	v1190 = int32(_a_F_ginbuild_2)
	v1192 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10]))
	v1193 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10])) = v1192 + v1193
	v1196 = *(*int32)(unsafe.Add(mBase, uint32(v1183)))
	*(*int32)(unsafe.Add(mBase, uint32(v1183))) = v1196 + v1193
	*(*int64)(unsafe.Add(mBase, uint32(v1183+int32(96))+232)) = v1153
	v1204 = *(*int32)(unsafe.Add(mBase, uint32(v1183)))
	*(*int32)(unsafe.Add(mBase, uint32(v1183))) = v1204 + v1193
	v1210 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[10])) = v1210 - v1193
	goto L245
L248:
	;
	F_pfree(m, v1215)
	mBase = m.M
	v1217 = m.ExcPending
	if v1217 != 0 {
		goto L1
	} else {
		goto L251
	}
L249:
	;
	goto L250
L250:
	;
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(v574)+20))
	if v1218 == int32(0) {
		goto L252
	} else {
		goto L253
	}
L251:
	;
	goto L250
L252:
	;
	F_pfree(m, v574)
	mBase = m.M
	v1227 = m.ExcPending
	if v1227 != 0 {
		goto L1
	} else {
		goto L257
	}
L253:
	;
	v1221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574)+2)))
	if v1221 != 0 {
		goto L252
	} else {
		goto L254
	}
L254:
	;
	v1222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574)+14)))
	if v1222 != 0 {
		goto L252
	} else {
		goto L255
	}
L255:
	;
	v1223 = *(*int32)(unsafe.Add(mBase, uint32(v574)+4))
	F_pfree(m, v1223)
	mBase = m.M
	v1225 = m.ExcPending
	if v1225 != 0 {
		goto L1
	} else {
		goto L256
	}
L256:
	;
	goto L252
L257:
	;
	v1228 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[22])))
	F_tuplesort_end(m, v1228)
	mBase = m.M
	v1230 = m.ExcPending
	if v1230 != 0 {
		goto L1
	} else {
		goto L258
	}
L258:
	;
	v1231 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[7])))
	F__brin_end_parallel(m, v1231)
	mBase = m.M
	v1233 = m.ExcPending
	if v1233 != 0 {
		goto L1
	} else {
		goto L259
	}
L259:
	;
	v1359 = v534
	goto L95
L260:
	;
	v1247 = int32(_a_F_ginbuild_15)
	v1248 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[13]))
	v1250 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[14])))
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[13])) = v1250
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[32])))
	v1255 = m.G0
	v1256 = int32(16)
	v1257 = v1255 - v1256
	m.G0 = v1257
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[33]))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[34]))) = v1254
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(v1254)))
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[35]))) = uint8(base.B2i32(v1264 == int32(_a_F_ginbuild_16)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[36]))) = int32(790)
	m.G0 = v1257 + v1256
	goto L261
L261:
	;
	v1281 = F_ginGetBAEntry(m, v189, v21+int32(12), v21+int32(_a_F_ginbuild_14), v21+int32(15), v21+int32(_a_F_ginbuild_13))
	mBase = m.M
	v1282 = m.ExcPending
	if v1282 != 0 {
		goto L1
	} else {
		goto L262
	}
L262:
	;
	if v1281 != 0 {
		goto L263
	} else {
		goto L264
	}
L263:
	;
	v1286 = v1281
	goto L266
L264:
	;
	goto L265
L265:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ginbuild[13])) = v1248
	v1359 = v1245
	goto L95
L266:
	;
	v1302 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[31]))
	if v1302 != 0 {
		goto L268
	} else {
		goto L269
	}
L267:
	;
	goto L265
L268:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1304 = m.ExcPending
	if v1304 != 0 {
		goto L1
	} else {
		goto L271
	}
L269:
	;
	goto L270
L270:
	;
	v1307 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+12)))
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[28])))
	v1309 = int32(*(*int8)(unsafe.Add(mBase, uint32(v21)+15)))
	v1310 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[26])))
	F_ginEntryInsert(m, v21+int32(16), v1307, v1308, v1309, v1286, v1310, v39)
	mBase = m.M
	v1312 = m.ExcPending
	if v1312 != 0 {
		goto L1
	} else {
		goto L272
	}
L271:
	;
	goto L270
L272:
	;
	v1321 = F_ginGetBAEntry(m, v189, v21+int32(12), v21+int32(_a_F_ginbuild_14), v21+int32(15), v21+int32(_a_F_ginbuild_13))
	mBase = m.M
	v1322 = m.ExcPending
	if v1322 != 0 {
		goto L1
	} else {
		goto L273
	}
L273:
	;
	if v1321 != 0 {
		v1286 = v1321
		goto L266
	} else {
		goto L274
	}
L274:
	;
	goto L267
L275:
	;
	v1364 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[14])))
	F_MemoryContextDelete(m, v1364)
	mBase = m.M
	v1366 = m.ExcPending
	if v1366 != 0 {
		goto L1
	} else {
		goto L276
	}
L276:
	;
	v1368 = F_RelationGetNumberOfBlocksInFork(m, l1, int32(0))
	mBase = m.M
	v1369 = m.ExcPending
	if v1369 != 0 {
		goto L1
	} else {
		goto L277
	}
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[37]))) = v1368
	F_ginUpdateStats(m, l1, v39, int32(1))
	mBase = m.M
	v1373 = m.ExcPending
	if v1373 != 0 {
		goto L1
	} else {
		goto L278
	}
L278:
	;
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v1375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1374)+118)))
	if v1375 != int32(112) {
		goto L279
	} else {
		goto L280
	}
L279:
	;
	v1392 = F_palloc(m, int32(16))
	mBase = m.M
	v1393 = m.ExcPending
	if v1393 != 0 {
		goto L1
	} else {
		goto L288
	}
L280:
	;
	v1379 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuild[38]))
	if v1379 <= int32(0) {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v1382 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v1382 != 0 {
		goto L279
	} else {
		goto L284
	}
L282:
	;
	goto L283
L283:
	;
	v1384 = int32(0)
	v1386 = F_RelationGetNumberOfBlocksInFork(m, l1, v1384)
	mBase = m.M
	v1387 = m.ExcPending
	if v1387 != 0 {
		goto L1
	} else {
		goto L286
	}
L284:
	;
	v1383 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v1383 != 0 {
		goto L279
	} else {
		goto L285
	}
L285:
	;
	goto L283
L286:
	;
	F_log_newpage_range(m, l1, v1384, v1386, int32(1))
	mBase = m.M
	v1390 = m.ExcPending
	if v1390 != 0 {
		goto L1
	} else {
		goto L287
	}
L287:
	;
	goto L279
L288:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1392))) = v1359
	v1395 = *(*float64)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ginbuild[8])))
	*(*float64)(unsafe.Add(mBase, uint32(v1392)+8)) = v1395
	m.G0 = v21 + int32(_a_F_ginbuild_0)
	return v1392
L289:
	;
	v1405 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v1405 + int32(4)
	F_errmsg_internal(m, int32(_a_F_ginbuild_17), v21)
	mBase = m.M
	v1411 = m.ExcPending
	if v1411 != 0 {
		goto L1
	} else {
		goto L290
	}
L290:
	;
	F_errfinish(m, int32(_a_F_ginbuild_11), int32(625), int32(_a_F_ginbuild_18))
	mBase = m.M
	v1416 = m.ExcPending
	if v1416 != 0 {
		goto L1
	} else {
		goto L291
	}
L291:
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
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v91 float64
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v102 float64
	_ = v102
	var v106 int32
	_ = v106
	var v109 float64
	_ = v109
	var v113 float64
	_ = v113
	var v118 float64
	_ = v118
	var v121 int32
	_ = v121
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 float64
	_ = v138
	var v139 float64
	_ = v139
	var v142 float64
	_ = v142
	var v147 int32
	_ = v147
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
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
				F_fmgr_info(m, v47, v12+int32(52))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v58+v25)))
					F_set_fn_opclass_options(m, v12+int32(52), v60)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						if v51 != 0 {
							v66 = v51
						} else {
							v66 = int32(100)
						}
						v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+48)))
						v78 = F_FunctionCall7Coll(m, v12+int32(52), v66, l3, v12+int32(36), v69, v12+int32(32), v12+int32(28), v12+int32(24), v12+int32(20))
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return int32(0)
						} else {
							v80 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
							v81 = int32(0)
							v83 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
							v86 = base.B2i32(v80 != v81) | base.B2i32(v81 < v83)
							if v86 == v81 {
							} else {
								if int32(0) < v83 {
									v91 = *(*float64)(unsafe.Add(mBase, uint32(l4)+80))
									v93 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
									v94 = int32(0)
									v102 = v91
									for {
										if v93 == int32(0) {
											v113 = *(*float64)(unsafe.Add(mBase, uint32(l4)+72))
											*(*float64)(unsafe.Add(mBase, uint32(l4)+72)) = base.F64_add(v113, float64(1))
										} else {
											v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94+v93))))
											if v106 != int32(1) {
												v113 = *(*float64)(unsafe.Add(mBase, uint32(l4)+72))
												*(*float64)(unsafe.Add(mBase, uint32(l4)+72)) = base.F64_add(v113, float64(1))
											} else {
												v109 = *(*float64)(unsafe.Add(mBase, uint32(l4)+64))
												*(*float64)(unsafe.Add(mBase, uint32(l4)+64)) = base.F64_add(v109, float64(100))
											}
										}
										v118 = base.F64_add(v102, float64(1))
										*(*float64)(unsafe.Add(mBase, uint32(l4)+80)) = v118
										v121 = v94 + int32(1)
										if v121 != v83 {
											v94 = v121
											v102 = v118
											continue
										} else {
											break
										}
										break
									}
								} else {
								}
								switch v80 {
								case 0:
									v133 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l1+l4)+32)) = uint8(v133)
								case 1:
									v136 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l1+l4)+32)) = uint8(v136)
									v138 = *(*float64)(unsafe.Add(mBase, uint32(l4)+72))
									v139 = float64(1)
									*(*float64)(unsafe.Add(mBase, uint32(l4)+72)) = base.F64_add(v138, v139)
									v142 = *(*float64)(unsafe.Add(mBase, uint32(l4)+80))
									*(*float64)(unsafe.Add(mBase, uint32(l4)+80)) = base.F64_add(v142, v139)
								default:
									v147 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l1+l4))) = uint8(v147)
								}
							}
							m.G0 = v12 + int32(80)
							return v86
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v165 = m.ExcPending
				if v165 != 0 {
					return int32(0)
				} else {
					v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v167 = F_get_rel_name(m, v166)
					mBase = m.M
					v168 = m.ExcPending
					if v168 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v167
						*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = l1 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(3)
						F_errmsg_internal(m, int32(_a_F_gincost_pattern_0), v12)
						mBase = m.M
						v177 = m.ExcPending
						if v177 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_gincost_pattern_1), int32(_a_F_gincost_pattern_2), int32(_a_F_gincost_pattern_3))
							mBase = m.M
							v182 = m.ExcPending
							if v182 != 0 {
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
	var v9 int32
	_ = v9
	var v12 int64
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
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = v9
	*(*int32)(unsafe.Add(mBase, uint32(v5)+20)) = l0
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v5)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v5)+8)) = v12
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
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	v2 = l1
	if l0&int32(3) != 0 {
	} else {
	}
	v29 = F___memset(m, l0, int32(0), int32(_a_F_gistinitpage_0))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+10)) = int32(_a_F_gistinitpage_1)
	v35 = int32(_a_F_gistinitpage_2)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)) = uint16(v35)
	v41 = int32(_a_F_gistinitpage_3)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v41)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v41)
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	v45 = l0 + v44
	v46 = int32(_a_F_gistinitpage_4)
	*(*uint16)(unsafe.Add(mBase, uint32(v45)+14)) = uint16(v46)
	*(*uint16)(unsafe.Add(mBase, uint32(v45)+12)) = uint16(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+8)) = int32(-1)
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	v7 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	if l1 == v7 {
		v60 = v7
		m.G0 = v12 + int32(16)
		return v60
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
					v56 = v21
					*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v56)
					v60 = v21
					m.G0 = v12 + int32(16)
					return v60
				} else {
					v33 = F_get_opclass_opfamily_and_input_type(m, v23, v12+int32(12), v12+int32(8))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						if v33 == int32(0) {
							v56 = v21
							*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v56)
							v60 = v21
							m.G0 = v12 + int32(16)
							return v60
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
								if l2 != int32(7) {
									v56 = v43
									*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v56)
									v60 = v21
									m.G0 = v12 + int32(16)
									return v60
								} else {
									if v40 != 0 {
										v56 = v43
										*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v56)
										v60 = v21
										m.G0 = v12 + int32(16)
										return v60
									} else {
										v47 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
										v48 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
										v50 = F_SearchSysCacheExists(m, int32(5), v47, v48, v48, int32(3))
										mBase = m.M
										v51 = m.ExcPending
										if v51 != 0 {
											return int32(0)
										} else {
											v53 = v50 ^ int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v53)
											v56 = v43
											*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v56)
											v60 = v21
											m.G0 = v12 + int32(16)
											return v60
										}
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
					v56 = v21
					*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v56)
					v60 = v21
					m.G0 = v12 + int32(16)
					return v60
				} else {
					v33 = F_get_opclass_opfamily_and_input_type(m, v23, v12+int32(12), v12+int32(8))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						if v33 == int32(0) {
							v56 = v21
							*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v56)
							v60 = v21
							m.G0 = v12 + int32(16)
							return v60
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
								if l2 != int32(7) {
									v56 = v43
									*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v56)
									v60 = v21
									m.G0 = v12 + int32(16)
									return v60
								} else {
									if v40 != 0 {
										v56 = v43
										*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v56)
										v60 = v21
										m.G0 = v12 + int32(16)
										return v60
									} else {
										v47 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
										v48 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
										v50 = F_SearchSysCacheExists(m, int32(5), v47, v48, v48, int32(3))
										mBase = m.M
										v51 = m.ExcPending
										if v51 != 0 {
											return int32(0)
										} else {
											v53 = v50 ^ int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v53)
											v56 = v43
											*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v56)
											v60 = v21
											m.G0 = v12 + int32(16)
											return v60
										}
									}
								}
							}
						}
					}
				}
			}
		default:
			v60 = v7
			m.G0 = v12 + int32(16)
			return v60
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
	var v2 int64
	_ = v2
	var v3 int32
	_ = v3
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
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v48 int64
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v65 int32
	_ = v65
	var v66 int64
	_ = v66
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int64
	_ = v86
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	v2 = int64(0)
	v3 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v12 <= v3 {
		v86 = v2
	} else {
		v16 = v12 & int32(3)
		v17 = int32(4)
		v18 = v11 + v17
		if base.Ui32(v12) < base.Ui32(v17) {
			v54 = int32(0)
			v55 = v2
		} else {
			v25 = int32(0)
			v26 = v2
			v31 = v3
			for {
				v34 = int32(4)
				v36 = v18 + v25<<(uint(v34)%32)
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+48))
				v38 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v36)+32))
				v40 = *(*int64)(unsafe.Add(mBase, uint32(v39)))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
				v42 = *(*int64)(unsafe.Add(mBase, uint32(v41)))
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
				v44 = *(*int64)(unsafe.Add(mBase, uint32(v43)))
				v48 = v38 | (v40 | (v42 | (v44 | v26)))
				v50 = v25 + v34
				v52 = v31 + v34
				if v52 != v12&int32(2147483644) {
					v25 = v50
					v26 = v48
					v31 = v52
					continue
				} else {
					break
				}
				break
			}
			v54 = v50
			v55 = v48
		}
		if v16 == int32(0) {
			v86 = v55
		} else {
			v65 = v54
			v66 = v55
			v70 = v3
			for {
				v77 = *(*int32)(unsafe.Add(mBase, uint32(v18+v65<<(uint(int32(4))%32))))
				v78 = *(*int64)(unsafe.Add(mBase, uint32(v77)))
				v79 = v78 | v66
				v80 = int32(1)
				v83 = v70 + v80
				if v83 != v16 {
					v65 = v65 + v80
					v66 = v79
					v70 = v83
					continue
				} else {
					break
				}
				break
			}
			v86 = v79
		}
	}
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(8)
	v96 = F_Int64GetDatum(m, v86)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		return int32(0)
	} else {
		return v96
	}
}
