package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CheckForStandbyTrigger(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	v4 = m.G0
	v6 = v4 - int32(96)
	m.G0 = v6
	v10 = int32(*(*uint8)(unsafe.Add(mBase, _consts[184])))
	if v10 != 0 {
		v75 = int32(1)
		m.G0 = v6 + int32(96)
		return v75
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, _consts[185]))
		if v12 == int32(0) {
			v75 = int32(0)
			m.G0 = v6 + int32(96)
			return v75
		} else {
			v16 = int32(0)
			v20 = F___fstatat(m, int32(-100), int32(347095), v6, v16)
			mBase = m.M
			if v20 != 0 {
				v75 = v16
				m.G0 = v6 + int32(96)
				return v75
			} else {
				v23 = F_errstart(m, int32(15), int32(0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					if v23 != 0 {
						F_errmsg(m, int32(76782), int32(0))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(489255), int32(4482), int32(223333))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								v37 = F_unlink(m, int32(347095))
								mBase = m.M
								*(*int32)(unsafe.Add(mBase, _consts[185])) = int32(0)
								v42 = *(*int32)(unsafe.Add(mBase, _consts[186]))
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+96))
								*(*int32)(unsafe.Add(mBase, uint32(v42)+96)) = int32(1)
								if v43 != 0 {
									v47 = *(*int32)(unsafe.Add(mBase, _consts[186]))
									F_s_lock(m, v47+int32(96), int32(489255), int32(4456), int32(448496))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										v55 = int32(4382964)
										v56 = *(*int32)(unsafe.Add(mBase, _consts[186]))
										v57 = int32(1)
										*(*int32)(unsafe.Add(mBase, uint32(v56)+96)) = v57
										*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)) = uint8(v57)
										v62 = *(*int32)(unsafe.Add(mBase, _consts[186]))
										v63 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v62)+96)) = v63
										*(*int32)(unsafe.Add(mBase, uint32(v62)+80)) = v63
										F_ConditionVariableBroadcast(m, v62+int32(84))
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return int32(0)
										} else {
											v71 = int32(1)
											*(*uint8)(unsafe.Add(mBase, _consts[184])) = uint8(v71)
											v75 = v71
											m.G0 = v6 + int32(96)
											return v75
										}
									}
								} else {
									v55 = int32(4382964)
									v56 = *(*int32)(unsafe.Add(mBase, _consts[186]))
									v57 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v56)+96)) = v57
									*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)) = uint8(v57)
									v62 = *(*int32)(unsafe.Add(mBase, _consts[186]))
									v63 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v62)+96)) = v63
									*(*int32)(unsafe.Add(mBase, uint32(v62)+80)) = v63
									F_ConditionVariableBroadcast(m, v62+int32(84))
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										v71 = int32(1)
										*(*uint8)(unsafe.Add(mBase, _consts[184])) = uint8(v71)
										v75 = v71
										m.G0 = v6 + int32(96)
										return v75
									}
								}
							}
						}
					} else {
						v37 = F_unlink(m, int32(347095))
						mBase = m.M
						*(*int32)(unsafe.Add(mBase, _consts[185])) = int32(0)
						v42 = *(*int32)(unsafe.Add(mBase, _consts[186]))
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+96))
						*(*int32)(unsafe.Add(mBase, uint32(v42)+96)) = int32(1)
						if v43 != 0 {
							v47 = *(*int32)(unsafe.Add(mBase, _consts[186]))
							F_s_lock(m, v47+int32(96), int32(489255), int32(4456), int32(448496))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								v55 = int32(4382964)
								v56 = *(*int32)(unsafe.Add(mBase, _consts[186]))
								v57 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v56)+96)) = v57
								*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)) = uint8(v57)
								v62 = *(*int32)(unsafe.Add(mBase, _consts[186]))
								v63 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v62)+96)) = v63
								*(*int32)(unsafe.Add(mBase, uint32(v62)+80)) = v63
								F_ConditionVariableBroadcast(m, v62+int32(84))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									v71 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _consts[184])) = uint8(v71)
									v75 = v71
									m.G0 = v6 + int32(96)
									return v75
								}
							}
						} else {
							v55 = int32(4382964)
							v56 = *(*int32)(unsafe.Add(mBase, _consts[186]))
							v57 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v56)+96)) = v57
							*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)) = uint8(v57)
							v62 = *(*int32)(unsafe.Add(mBase, _consts[186]))
							v63 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v62)+96)) = v63
							*(*int32)(unsafe.Add(mBase, uint32(v62)+80)) = v63
							F_ConditionVariableBroadcast(m, v62+int32(84))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								v71 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _consts[184])) = uint8(v71)
								v75 = v71
								m.G0 = v6 + int32(96)
								return v75
							}
						}
					}
				}
			}
		}
	}
}
func F_WaitForParallelWorkersToFinish(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int64
	_ = v121
	var v123 int64
	_ = v123
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	goto L2
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L7
	} else {
		goto L37
	}
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v20 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v114 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L4:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v23 <= int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	return
L8:
	;
	goto L6
L9:
	;
	goto L3
L10:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v27 = int32(0)
	v30 = v27
	v33 = v27
	goto L12
L11:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _consts[80]))
	v101 = F_WaitLatch(m, v97, int32(33), int32(-1), int32(134217768))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L7
	} else {
		goto L31
	}
L12:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v26+v30<<(uint(int32(3))%32))+4))
	if v39 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	if v23 <= v47 {
		goto L9
	} else {
		goto L20
	}
L14:
	;
	v49 = v30 + int32(1)
	if v49 != v23 {
		v30 = v49
		v33 = v47
		goto L12
	} else {
		goto L19
	}
L15:
	;
	v47 = v33 + int32(1)
	goto L14
L16:
	;
	goto L17
L17:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+v30))))
	if v46 != 0 {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	v47 = v33
	goto L14
L19:
	;
	goto L13
L20:
	;
	v54 = int32(0)
	goto L21
L21:
	;
	v61 = v54 << (uint(int32(3)) % 32)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v63 = v61 + v62
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if v64 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L11
L23:
	;
	v86 = v54 + int32(1)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v86 < v87 {
		v54 = v86
		goto L21
	} else {
		goto L30
	}
L24:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	if v67 == int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v72 = F_GetBackgroundWorkerPid(m, v67, v10+int32(12))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L7
	} else {
		goto L26
	}
L26:
	;
	if v72 != int32(2) {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v76+v61)+4))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v80 = F_shm_mq_get_sender(m, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L7
	} else {
		goto L28
	}
L28:
	;
	if v80 == int32(0) {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	goto L23
L30:
	;
	goto L22
L31:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _consts[80]))
	*(*int32)(unsafe.Add(mBase, uint32(v104))) = int32(0)
	goto L32
L32:
	;
	goto L2
L33:
	;
	m.G0 = v10 + int32(16)
	return
L34:
	;
	v119 = F_shm_toc_lookup(m, v114, int64(-65535), int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L7
	} else {
		goto L35
	}
L35:
	;
	v121 = *(*int64)(unsafe.Add(mBase, uint32(v119)+72))
	v123 = *(*int64)(unsafe.Add(mBase, _consts[81]))
	if base.Ui64(v121) <= base.Ui64(v123) {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	*(*int64)(unsafe.Add(mBase, _consts[81])) = v121
	goto L33
L37:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L7
	} else {
		goto L38
	}
L38:
	;
	F_errmsg(m, int32(339567), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L7
	} else {
		goto L39
	}
L39:
	;
	F_errhint(m, int32(607288), int32(0))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L7
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(494337), int32(879), int32(320648))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L7
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_WaitForStandbyConfirmation(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	v3 = *(*int32)(unsafe.Add(mBase, _consts[512]))
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+202)))
	if v4 != int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, _consts[545]))
	if v8 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _consts[547]))
	F_ConditionVariablePrepareToSleep(m, v12+int32(76))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	goto L6
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v19 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L4
	} else {
		goto L21
	}
L8:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _consts[305]))
	if v23 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L10
L12:
	;
	*(*int32)(unsafe.Add(mBase, _consts[305])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L4
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v31 = F_StandbySlotsHaveCaughtup(m, l0, int32(19))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L16
	}
L15:
	;
	goto L14
L16:
	;
	if v31 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _consts[547]))
	v41 = F_ConditionVariableTimedSleep(m, v36+int32(76), int32(1000), int32(100663302))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	goto L7
L20:
	;
	goto L6
L21:
	;
	goto L1
}
