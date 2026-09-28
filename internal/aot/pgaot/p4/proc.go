package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_ExecSetExecProcNode(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(678)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(766)
	return
}
func F_ProcArrayAdd(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	v2 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayAdd[0]))
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayAdd[1]))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayAdd[2]))
	v22 = F_LWLockAcquire(m, v18+int32(512), v2)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayAdd[2]))
	v29 = F_LWLockAcquire(m, v25+int32(384), int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v31 < v32 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v36 = base.I32_div_s(l0-v16, int32(768))
	v38 = v13 + int32(36)
	if v31 <= int32(0) {
		v62 = v2
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L33
	}
L7:
	;
	v71 = int32(2)
	v72 = v62 << (uint(v71) % 32)
	v73 = v38 + v72
	v75 = v62 + int32(1)
	v77 = v75 << (uint(v71) % 32)
	v78 = v31 - v62
	v80 = v78 << (uint(v71) % 32)
	v81 = int32(0)
	v82 = base.B2i32(v80 == v81)
	if v82 == v81 {
		goto L13
	} else {
		goto L14
	}
L8:
	;
	v43 = v2
	goto L9
L9:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v38+v43<<(uint(int32(2))%32))))
	if v36 < v55 {
		v62 = v43
		goto L7
	} else {
		goto L11
	}
L10:
	;
	v62 = v31
	goto L7
L11:
	;
	v58 = v43 + int32(1)
	if v58 != v31 {
		v43 = v58
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	base.MemoryCopy(m, v77+v38, v73, v80)
	goto L15
L14:
	;
	goto L15
L15:
	;
	if v82 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayAdd[1]))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	base.MemoryCopy(m, v91+v77, v91+v72, v80)
	goto L18
L17:
	;
	goto L18
L18:
	;
	v96 = int32(1)
	v97 = v62 << (uint(v96) % 32)
	v99 = v78 << (uint(v96) % 32)
	if v99 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayAdd[1]))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+8))
	base.MemoryCopy(m, v102+v75<<(uint(int32(1))%32), v102+v97, v99)
	goto L21
L20:
	;
	goto L21
L21:
	;
	if v78 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayAdd[1]))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+12))
	base.MemoryCopy(m, v111+v75, v111+v62, v78)
	goto L24
L23:
	;
	goto L24
L24:
	;
	v116 = int32(_a_F_ProcArrayAdd_0)
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayAdd[1]))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	v121 = base.I32_div_s(l0-v118, int32(768))
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = v121
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v62
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayAdd[1]))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v126+v72))) = v128
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayAdd[1]))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+8))
	v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+56)))
	*(*uint16)(unsafe.Add(mBase, uint32(v132+v97))) = uint16(v134)
	v137 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayAdd[1]))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	*(*uint8)(unsafe.Add(mBase, uint32(v138+v62))) = uint8(v140)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v144 = v142 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v144
	if v75 < v144 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayAdd[3]))
	v153 = v75
	goto L28
L26:
	;
	goto L27
L27:
	;
	v184 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayAdd[2]))
	F_LWLockRelease(m, v184+int32(384))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L31
	}
L28:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v38+v153<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v148+v163*int32(768))+32)) = v153
	v169 = v153 + int32(1)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v169 < v170 {
		v153 = v169
		goto L28
	} else {
		goto L30
	}
L29:
	;
	goto L27
L30:
	;
	goto L29
L31:
	;
	v190 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayAdd[2]))
	F_LWLockRelease(m, v190+int32(512))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	return
L33:
	;
	F_errcode(m, int32(_a_F_ProcArrayAdd_1))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_errmsg(m, int32(_a_F_ProcArrayAdd_2), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_ProcArrayAdd_3), int32(484), int32(_a_F_ProcArrayAdd_4))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ProcKill(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
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
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
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
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	v3 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[0]))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v13 = m.Env.Pgmem_getpid(m)
	mBase = m.M
	if v12 == v13 {
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[0]))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+600))
		if v17 != 0 {
			v19 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[1]))
			v23 = F_LWLockAcquire(m, v19+int32(_a_F_ProcKill_0), int32(0))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[0]))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+600))
				if v27 != 0 {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)+596))
					*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = v27
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+596))
					*(*int32)(unsafe.Add(mBase, uint32(v27))) = v30
					*(*int64)(unsafe.Add(mBase, uint32(v26)+596)) = int64(0)
				} else {
				}
				v35 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[1]))
				F_LWLockRelease(m, v35+int32(_a_F_ProcKill_0))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					F_LWLockReleaseAll(m)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						F_ConditionVariableCancelSleep(m)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							F_SwitchBackToLocalLatch(m)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return
							} else {
								v49 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[0]))
								*(*int32)(unsafe.Add(mBase, uint32(v49+int32(316))+12)) = int32(0)
								v55 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[0]))
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
								v57 = int32(1)
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)+364))
								if v58 != 0 {
									v60 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[1]))
									v62 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
									v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
									v66 = base.I32_div_s(v58-v63, int32(768))
									v68 = base.I32_rem_s(v66, int32(16))
									v73 = v60 + v68<<(uint(int32(7))%32) + int32(_a_F_ProcKill_1)
									v75 = F_LWLockAcquire(m, v73, int32(0))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return
									} else {
										v77 = *(*int32)(unsafe.Add(mBase, uint32(v55)+376))
										v78 = *(*int32)(unsafe.Add(mBase, uint32(v55)+380))
										*(*int32)(unsafe.Add(mBase, uint32(v77)+4)) = v78
										v80 = *(*int32)(unsafe.Add(mBase, uint32(v55)+376))
										*(*int32)(unsafe.Add(mBase, uint32(v78))) = v80
										v82 = *(*int32)(unsafe.Add(mBase, uint32(v58)+372))
										v88 = base.B2i32(v82 == int32(0)) | base.B2i32(v82 == v58+int32(368))
										if v88 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(v58)+364)) = int32(0)
											if v55 != v58 {
												*(*int32)(unsafe.Add(mBase, uint32(v55)+364)) = int32(0)
												v98 = v88
												v99 = int32(1)
											} else {
												v98 = v3
												v99 = v57
											}
										} else {
											if v55 == v58 {
												v98 = v3
												v99 = int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v55)+364)) = int32(0)
												v98 = v88
												v99 = int32(1)
											}
										}
										F_LWLockRelease(m, v73)
										mBase = m.M
										v101 = m.ExcPending
										if v101 != 0 {
											return
										} else {
											v102 = v98
											v104 = v99
											*(*int32)(unsafe.Add(mBase, _c_F_ProcKill[3])) = int32(_a_F_ProcKill_2)
											*(*int32)(unsafe.Add(mBase, _c_F_ProcKill[4])) = int32(-1)
											v113 = int32(0)
											*(*int32)(unsafe.Add(mBase, _c_F_ProcKill[0])) = v113
											*(*int64)(unsafe.Add(mBase, uint32(v55)+40)) = int64(4294967295)
											*(*int32)(unsafe.Add(mBase, uint32(v55)+12)) = v113
											v120 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
											v123 = base.AtomicRmwXchg32(m, v120, int32(20), int32(1))
											if v123 != 0 {
												F_s_lock(m, v120+int32(20), int32(_a_F_ProcKill_3))
												mBase = m.M
												v128 = m.ExcPending
												if v128 != 0 {
													return
												} else {
													if v102 != 0 {
														v129 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
														v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
														if v130 == int32(0) {
															*(*int32)(unsafe.Add(mBase, uint32(v129))) = v129
															v134 = v129
														} else {
															v134 = v130
														}
														*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v129
														*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = v134
														v138 = v58 + int32(4)
														*(*int32)(unsafe.Add(mBase, uint32(v134))) = v138
														*(*int32)(unsafe.Add(mBase, uint32(v129)+4)) = v138
													} else {
													}
													if v104 != 0 {
														v143 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
														if v143 == int32(0) {
															*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = v56
															*(*int32)(unsafe.Add(mBase, uint32(v56))) = v56
														} else {
														}
														*(*int32)(unsafe.Add(mBase, uint32(v55)+8)) = v56
														v149 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
														*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = v149
														v152 = v55 + int32(4)
														*(*int32)(unsafe.Add(mBase, uint32(v149)+4)) = v152
														*(*int32)(unsafe.Add(mBase, uint32(v56))) = v152
													} else {
													}
													v157 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
													v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+72))
													v160 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
													v165 = base.I32_div_s(v160+v158*int32(15), int32(16))
													v167 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
													*(*int32)(unsafe.Add(mBase, uint32(v167)+72)) = v165
													v169 = int32(0)
													atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v167)+20)), uint32(v169))
													return
												}
											} else {
												if v102 != 0 {
													v129 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
													v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
													if v130 == int32(0) {
														*(*int32)(unsafe.Add(mBase, uint32(v129))) = v129
														v134 = v129
													} else {
														v134 = v130
													}
													*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v129
													*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = v134
													v138 = v58 + int32(4)
													*(*int32)(unsafe.Add(mBase, uint32(v134))) = v138
													*(*int32)(unsafe.Add(mBase, uint32(v129)+4)) = v138
												} else {
												}
												if v104 != 0 {
													v143 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
													if v143 == int32(0) {
														*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = v56
														*(*int32)(unsafe.Add(mBase, uint32(v56))) = v56
													} else {
													}
													*(*int32)(unsafe.Add(mBase, uint32(v55)+8)) = v56
													v149 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
													*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = v149
													v152 = v55 + int32(4)
													*(*int32)(unsafe.Add(mBase, uint32(v149)+4)) = v152
													*(*int32)(unsafe.Add(mBase, uint32(v56))) = v152
												} else {
												}
												v157 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
												v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+72))
												v160 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
												v165 = base.I32_div_s(v160+v158*int32(15), int32(16))
												v167 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
												*(*int32)(unsafe.Add(mBase, uint32(v167)+72)) = v165
												v169 = int32(0)
												atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v167)+20)), uint32(v169))
												return
											}
										}
									}
								} else {
									v102 = v3
									v104 = v57
									*(*int32)(unsafe.Add(mBase, _c_F_ProcKill[3])) = int32(_a_F_ProcKill_2)
									*(*int32)(unsafe.Add(mBase, _c_F_ProcKill[4])) = int32(-1)
									v113 = int32(0)
									*(*int32)(unsafe.Add(mBase, _c_F_ProcKill[0])) = v113
									*(*int64)(unsafe.Add(mBase, uint32(v55)+40)) = int64(4294967295)
									*(*int32)(unsafe.Add(mBase, uint32(v55)+12)) = v113
									v120 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
									v123 = base.AtomicRmwXchg32(m, v120, int32(20), int32(1))
									if v123 != 0 {
										F_s_lock(m, v120+int32(20), int32(_a_F_ProcKill_3))
										mBase = m.M
										v128 = m.ExcPending
										if v128 != 0 {
											return
										} else {
											if v102 != 0 {
												v129 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
												v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
												if v130 == int32(0) {
													*(*int32)(unsafe.Add(mBase, uint32(v129))) = v129
													v134 = v129
												} else {
													v134 = v130
												}
												*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v129
												*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = v134
												v138 = v58 + int32(4)
												*(*int32)(unsafe.Add(mBase, uint32(v134))) = v138
												*(*int32)(unsafe.Add(mBase, uint32(v129)+4)) = v138
											} else {
											}
											if v104 != 0 {
												v143 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
												if v143 == int32(0) {
													*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = v56
													*(*int32)(unsafe.Add(mBase, uint32(v56))) = v56
												} else {
												}
												*(*int32)(unsafe.Add(mBase, uint32(v55)+8)) = v56
												v149 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
												*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = v149
												v152 = v55 + int32(4)
												*(*int32)(unsafe.Add(mBase, uint32(v149)+4)) = v152
												*(*int32)(unsafe.Add(mBase, uint32(v56))) = v152
											} else {
											}
											v157 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
											v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+72))
											v160 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
											v165 = base.I32_div_s(v160+v158*int32(15), int32(16))
											v167 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
											*(*int32)(unsafe.Add(mBase, uint32(v167)+72)) = v165
											v169 = int32(0)
											atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v167)+20)), uint32(v169))
											return
										}
									} else {
										if v102 != 0 {
											v129 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
											v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
											if v130 == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(v129))) = v129
												v134 = v129
											} else {
												v134 = v130
											}
											*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v129
											*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = v134
											v138 = v58 + int32(4)
											*(*int32)(unsafe.Add(mBase, uint32(v134))) = v138
											*(*int32)(unsafe.Add(mBase, uint32(v129)+4)) = v138
										} else {
										}
										if v104 != 0 {
											v143 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
											if v143 == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = v56
												*(*int32)(unsafe.Add(mBase, uint32(v56))) = v56
											} else {
											}
											*(*int32)(unsafe.Add(mBase, uint32(v55)+8)) = v56
											v149 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
											*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = v149
											v152 = v55 + int32(4)
											*(*int32)(unsafe.Add(mBase, uint32(v149)+4)) = v152
											*(*int32)(unsafe.Add(mBase, uint32(v56))) = v152
										} else {
										}
										v157 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
										v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+72))
										v160 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
										v165 = base.I32_div_s(v160+v158*int32(15), int32(16))
										v167 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
										*(*int32)(unsafe.Add(mBase, uint32(v167)+72)) = v165
										v169 = int32(0)
										atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v167)+20)), uint32(v169))
										return
									}
								}
							}
						}
					}
				}
			}
		} else {
			F_LWLockReleaseAll(m)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return
			} else {
				F_ConditionVariableCancelSleep(m)
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return
				} else {
					F_SwitchBackToLocalLatch(m)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[0]))
						*(*int32)(unsafe.Add(mBase, uint32(v49+int32(316))+12)) = int32(0)
						v55 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[0]))
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
						v57 = int32(1)
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)+364))
						if v58 != 0 {
							v60 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[1]))
							v62 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
							v66 = base.I32_div_s(v58-v63, int32(768))
							v68 = base.I32_rem_s(v66, int32(16))
							v73 = v60 + v68<<(uint(int32(7))%32) + int32(_a_F_ProcKill_1)
							v75 = F_LWLockAcquire(m, v73, int32(0))
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return
							} else {
								v77 = *(*int32)(unsafe.Add(mBase, uint32(v55)+376))
								v78 = *(*int32)(unsafe.Add(mBase, uint32(v55)+380))
								*(*int32)(unsafe.Add(mBase, uint32(v77)+4)) = v78
								v80 = *(*int32)(unsafe.Add(mBase, uint32(v55)+376))
								*(*int32)(unsafe.Add(mBase, uint32(v78))) = v80
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v58)+372))
								v88 = base.B2i32(v82 == int32(0)) | base.B2i32(v82 == v58+int32(368))
								if v88 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v58)+364)) = int32(0)
									if v55 != v58 {
										*(*int32)(unsafe.Add(mBase, uint32(v55)+364)) = int32(0)
										v98 = v88
										v99 = int32(1)
									} else {
										v98 = v3
										v99 = v57
									}
								} else {
									if v55 == v58 {
										v98 = v3
										v99 = int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v55)+364)) = int32(0)
										v98 = v88
										v99 = int32(1)
									}
								}
								F_LWLockRelease(m, v73)
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return
								} else {
									v102 = v98
									v104 = v99
									*(*int32)(unsafe.Add(mBase, _c_F_ProcKill[3])) = int32(_a_F_ProcKill_2)
									*(*int32)(unsafe.Add(mBase, _c_F_ProcKill[4])) = int32(-1)
									v113 = int32(0)
									*(*int32)(unsafe.Add(mBase, _c_F_ProcKill[0])) = v113
									*(*int64)(unsafe.Add(mBase, uint32(v55)+40)) = int64(4294967295)
									*(*int32)(unsafe.Add(mBase, uint32(v55)+12)) = v113
									v120 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
									v123 = base.AtomicRmwXchg32(m, v120, int32(20), int32(1))
									if v123 != 0 {
										F_s_lock(m, v120+int32(20), int32(_a_F_ProcKill_3))
										mBase = m.M
										v128 = m.ExcPending
										if v128 != 0 {
											return
										} else {
											if v102 != 0 {
												v129 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
												v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
												if v130 == int32(0) {
													*(*int32)(unsafe.Add(mBase, uint32(v129))) = v129
													v134 = v129
												} else {
													v134 = v130
												}
												*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v129
												*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = v134
												v138 = v58 + int32(4)
												*(*int32)(unsafe.Add(mBase, uint32(v134))) = v138
												*(*int32)(unsafe.Add(mBase, uint32(v129)+4)) = v138
											} else {
											}
											if v104 != 0 {
												v143 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
												if v143 == int32(0) {
													*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = v56
													*(*int32)(unsafe.Add(mBase, uint32(v56))) = v56
												} else {
												}
												*(*int32)(unsafe.Add(mBase, uint32(v55)+8)) = v56
												v149 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
												*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = v149
												v152 = v55 + int32(4)
												*(*int32)(unsafe.Add(mBase, uint32(v149)+4)) = v152
												*(*int32)(unsafe.Add(mBase, uint32(v56))) = v152
											} else {
											}
											v157 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
											v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+72))
											v160 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
											v165 = base.I32_div_s(v160+v158*int32(15), int32(16))
											v167 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
											*(*int32)(unsafe.Add(mBase, uint32(v167)+72)) = v165
											v169 = int32(0)
											atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v167)+20)), uint32(v169))
											return
										}
									} else {
										if v102 != 0 {
											v129 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
											v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
											if v130 == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(v129))) = v129
												v134 = v129
											} else {
												v134 = v130
											}
											*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v129
											*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = v134
											v138 = v58 + int32(4)
											*(*int32)(unsafe.Add(mBase, uint32(v134))) = v138
											*(*int32)(unsafe.Add(mBase, uint32(v129)+4)) = v138
										} else {
										}
										if v104 != 0 {
											v143 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
											if v143 == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = v56
												*(*int32)(unsafe.Add(mBase, uint32(v56))) = v56
											} else {
											}
											*(*int32)(unsafe.Add(mBase, uint32(v55)+8)) = v56
											v149 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
											*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = v149
											v152 = v55 + int32(4)
											*(*int32)(unsafe.Add(mBase, uint32(v149)+4)) = v152
											*(*int32)(unsafe.Add(mBase, uint32(v56))) = v152
										} else {
										}
										v157 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
										v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+72))
										v160 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
										v165 = base.I32_div_s(v160+v158*int32(15), int32(16))
										v167 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
										*(*int32)(unsafe.Add(mBase, uint32(v167)+72)) = v165
										v169 = int32(0)
										atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v167)+20)), uint32(v169))
										return
									}
								}
							}
						} else {
							v102 = v3
							v104 = v57
							*(*int32)(unsafe.Add(mBase, _c_F_ProcKill[3])) = int32(_a_F_ProcKill_2)
							*(*int32)(unsafe.Add(mBase, _c_F_ProcKill[4])) = int32(-1)
							v113 = int32(0)
							*(*int32)(unsafe.Add(mBase, _c_F_ProcKill[0])) = v113
							*(*int64)(unsafe.Add(mBase, uint32(v55)+40)) = int64(4294967295)
							*(*int32)(unsafe.Add(mBase, uint32(v55)+12)) = v113
							v120 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
							v123 = base.AtomicRmwXchg32(m, v120, int32(20), int32(1))
							if v123 != 0 {
								F_s_lock(m, v120+int32(20), int32(_a_F_ProcKill_3))
								mBase = m.M
								v128 = m.ExcPending
								if v128 != 0 {
									return
								} else {
									if v102 != 0 {
										v129 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
										v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
										if v130 == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(v129))) = v129
											v134 = v129
										} else {
											v134 = v130
										}
										*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v129
										*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = v134
										v138 = v58 + int32(4)
										*(*int32)(unsafe.Add(mBase, uint32(v134))) = v138
										*(*int32)(unsafe.Add(mBase, uint32(v129)+4)) = v138
									} else {
									}
									if v104 != 0 {
										v143 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
										if v143 == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = v56
											*(*int32)(unsafe.Add(mBase, uint32(v56))) = v56
										} else {
										}
										*(*int32)(unsafe.Add(mBase, uint32(v55)+8)) = v56
										v149 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
										*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = v149
										v152 = v55 + int32(4)
										*(*int32)(unsafe.Add(mBase, uint32(v149)+4)) = v152
										*(*int32)(unsafe.Add(mBase, uint32(v56))) = v152
									} else {
									}
									v157 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
									v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+72))
									v160 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
									v165 = base.I32_div_s(v160+v158*int32(15), int32(16))
									v167 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
									*(*int32)(unsafe.Add(mBase, uint32(v167)+72)) = v165
									v169 = int32(0)
									atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v167)+20)), uint32(v169))
									return
								}
							} else {
								if v102 != 0 {
									v129 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
									v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
									if v130 == int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(v129))) = v129
										v134 = v129
									} else {
										v134 = v130
									}
									*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v129
									*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = v134
									v138 = v58 + int32(4)
									*(*int32)(unsafe.Add(mBase, uint32(v134))) = v138
									*(*int32)(unsafe.Add(mBase, uint32(v129)+4)) = v138
								} else {
								}
								if v104 != 0 {
									v143 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
									if v143 == int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = v56
										*(*int32)(unsafe.Add(mBase, uint32(v56))) = v56
									} else {
									}
									*(*int32)(unsafe.Add(mBase, uint32(v55)+8)) = v56
									v149 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
									*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = v149
									v152 = v55 + int32(4)
									*(*int32)(unsafe.Add(mBase, uint32(v149)+4)) = v152
									*(*int32)(unsafe.Add(mBase, uint32(v56))) = v152
								} else {
								}
								v157 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
								v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+72))
								v160 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
								v165 = base.I32_div_s(v160+v158*int32(15), int32(16))
								v167 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
								*(*int32)(unsafe.Add(mBase, uint32(v167)+72)) = v165
								v169 = int32(0)
								atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v167)+20)), uint32(v169))
								return
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(24), int32(0))
		mBase = m.M
		v175 = m.ExcPending
		if v175 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_ProcKill_4), int32(0))
			mBase = m.M
			v179 = m.ExcPending
			if v179 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_ProcKill_5), int32(935), int32(_a_F_ProcKill_6))
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
		}
	}
}
