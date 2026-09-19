package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_ExecSetExecProcNode(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(632)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(720)
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
	v36 = base.I32_div_s(l0-v16, int32(640))
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
	v121 = base.I32_div_s(l0-v118, int32(640))
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = v121
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v62
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayAdd[1]))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v126+v72))) = v128
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayAdd[1]))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+8))
	v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+276)))
	*(*uint16)(unsafe.Add(mBase, uint32(v132+v97))) = uint16(v134)
	v137 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayAdd[1]))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+124)))
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
	*(*int32)(unsafe.Add(mBase, uint32(v148+v163*int32(640))+48)) = v153
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
	F_errfinish(m, int32(_a_F_ProcArrayAdd_3), int32(488), int32(_a_F_ProcArrayAdd_4))
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
func F_ProcKill(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
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
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
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
	var v203 int32
	_ = v203
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
	v8 = m.Env.Pgmem_getpid(m)
	mBase = m.M
	if v7 == v8 {
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[0]))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+144))
		if v12 != 0 {
			v14 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[1]))
			v18 = F_LWLockAcquire(m, v14+int32(_a_F_ProcKill_0), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[0]))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+144))
				if v22 != 0 {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)+140))
					*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v22
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)+140))
					*(*int32)(unsafe.Add(mBase, uint32(v22))) = v25
					*(*int64)(unsafe.Add(mBase, uint32(v21)+140)) = int64(0)
				} else {
				}
				v30 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[1]))
				F_LWLockRelease(m, v30+int32(_a_F_ProcKill_0))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					F_LWLockReleaseAll(m)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						F_ConditionVariableCancelSleep(m)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							v42 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[0]))
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+616))
							if v43 != 0 {
								v45 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[1]))
								v47 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
								v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
								v51 = base.I32_div_s(v43-v48, int32(640))
								v53 = base.I32_rem_s(v51, int32(16))
								v58 = v45 + v53<<(uint(int32(7))%32) + int32(_a_F_ProcKill_1)
								v60 = F_LWLockAcquire(m, v58, int32(0))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return
								} else {
									v63 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[0]))
									v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+628))
									v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)+632))
									*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = v65
									v67 = *(*int32)(unsafe.Add(mBase, uint32(v63)+628))
									*(*int32)(unsafe.Add(mBase, uint32(v65))) = v67
									v69 = *(*int32)(unsafe.Add(mBase, uint32(v43)+624))
									if v69 != v43+int32(620) {
										v74 = v69
									} else {
										v74 = int32(0)
									}
									if v74 == int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(v43)+616)) = int32(0)
										v80 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[0]))
										if v43 == v80 {
											F_LWLockRelease(m, v58)
											mBase = m.M
											v117 = m.ExcPending
											if v117 != 0 {
												return
											} else {
												F_SwitchBackToLocalLatch(m)
												mBase = m.M
												v122 = m.ExcPending
												if v122 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_ProcKill[3])) = int32(_a_F_ProcKill_2)
													*(*int32)(unsafe.Add(mBase, _c_F_ProcKill[4])) = int32(-1)
													v129 = int32(_a_F_ProcKill_3)
													v130 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[0]))
													v132 = int32(0)
													*(*int32)(unsafe.Add(mBase, _c_F_ProcKill[0])) = v132
													*(*int32)(unsafe.Add(mBase, uint32(v130+int32(20))+12)) = v132
													*(*int64)(unsafe.Add(mBase, uint32(v130)+52)) = int64(4294967295)
													v140 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v130)+44)) = v140
													v142 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
													v144 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
													v147 = base.AtomicRmwXchg32(m, v144, v140, int32(1))
													if v147 != 0 {
														v149 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
														F_s_lock(m, v149, int32(_a_F_ProcKill_4), int32(1008), int32(_a_F_ProcKill_5))
														mBase = m.M
														v154 = m.ExcPending
														if v154 != 0 {
															return
														} else {
															v155 = *(*int32)(unsafe.Add(mBase, uint32(v130)+616))
															if v155 == int32(0) {
																v158 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
																if v158 == int32(0) {
																	*(*int32)(unsafe.Add(mBase, uint32(v142)+4)) = v142
																	*(*int32)(unsafe.Add(mBase, uint32(v142))) = v142
																} else {
																}
																*(*int32)(unsafe.Add(mBase, uint32(v130)+4)) = v142
																v164 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
																*(*int32)(unsafe.Add(mBase, uint32(v130))) = v164
																*(*int32)(unsafe.Add(mBase, uint32(v164)+4)) = v130
																*(*int32)(unsafe.Add(mBase, uint32(v142))) = v130
															} else {
															}
															v170 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
															v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+68))
															v173 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[6]))
															v178 = base.I32_div_s(v173+v171*int32(15), int32(16))
															v180 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
															*(*int32)(unsafe.Add(mBase, uint32(v180)+68)) = v178
															v183 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
															v184 = int32(0)
															atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v183))), uint32(v184))
															v188 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[7]))
															if v188 != 0 {
																v190 = F_pgmem_kill(m, v188, int32(12))
																mBase = m.M
															} else {
															}
															return
														}
													} else {
														v155 = *(*int32)(unsafe.Add(mBase, uint32(v130)+616))
														if v155 == int32(0) {
															v158 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
															if v158 == int32(0) {
																*(*int32)(unsafe.Add(mBase, uint32(v142)+4)) = v142
																*(*int32)(unsafe.Add(mBase, uint32(v142))) = v142
															} else {
															}
															*(*int32)(unsafe.Add(mBase, uint32(v130)+4)) = v142
															v164 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
															*(*int32)(unsafe.Add(mBase, uint32(v130))) = v164
															*(*int32)(unsafe.Add(mBase, uint32(v164)+4)) = v130
															*(*int32)(unsafe.Add(mBase, uint32(v142))) = v130
														} else {
														}
														v170 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
														v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+68))
														v173 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[6]))
														v178 = base.I32_div_s(v173+v171*int32(15), int32(16))
														v180 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
														*(*int32)(unsafe.Add(mBase, uint32(v180)+68)) = v178
														v183 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
														v184 = int32(0)
														atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v183))), uint32(v184))
														v188 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[7]))
														if v188 != 0 {
															v190 = F_pgmem_kill(m, v188, int32(12))
															mBase = m.M
														} else {
														}
														return
													}
												}
											}
										} else {
											v82 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
											v84 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
											v87 = base.AtomicRmwXchg32(m, v84, int32(0), int32(1))
											if v87 != 0 {
												v89 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
												F_s_lock(m, v89, int32(_a_F_ProcKill_4), int32(975), int32(_a_F_ProcKill_5))
												mBase = m.M
												v94 = m.ExcPending
												if v94 != 0 {
													return
												} else {
													v95 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
													if v95 == int32(0) {
														*(*int32)(unsafe.Add(mBase, uint32(v82))) = v82
														v99 = v82
													} else {
														v99 = v95
													}
													*(*int32)(unsafe.Add(mBase, uint32(v43))) = v82
													*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v99
													*(*int32)(unsafe.Add(mBase, uint32(v99))) = v43
													*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = v43
													v105 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
													v106 = int32(0)
													atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v105))), uint32(v106))
													F_LWLockRelease(m, v58)
													mBase = m.M
													v117 = m.ExcPending
													if v117 != 0 {
														return
													} else {
														F_SwitchBackToLocalLatch(m)
														mBase = m.M
														v122 = m.ExcPending
														if v122 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_ProcKill[3])) = int32(_a_F_ProcKill_2)
															*(*int32)(unsafe.Add(mBase, _c_F_ProcKill[4])) = int32(-1)
															v129 = int32(_a_F_ProcKill_3)
															v130 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[0]))
															v132 = int32(0)
															*(*int32)(unsafe.Add(mBase, _c_F_ProcKill[0])) = v132
															*(*int32)(unsafe.Add(mBase, uint32(v130+int32(20))+12)) = v132
															*(*int64)(unsafe.Add(mBase, uint32(v130)+52)) = int64(4294967295)
															v140 = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(v130)+44)) = v140
															v142 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
															v144 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
															v147 = base.AtomicRmwXchg32(m, v144, v140, int32(1))
															if v147 != 0 {
																v149 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
																F_s_lock(m, v149, int32(_a_F_ProcKill_4), int32(1008), int32(_a_F_ProcKill_5))
																mBase = m.M
																v154 = m.ExcPending
																if v154 != 0 {
																	return
																} else {
																	v155 = *(*int32)(unsafe.Add(mBase, uint32(v130)+616))
																	if v155 == int32(0) {
																		v158 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
																		if v158 == int32(0) {
																			*(*int32)(unsafe.Add(mBase, uint32(v142)+4)) = v142
																			*(*int32)(unsafe.Add(mBase, uint32(v142))) = v142
																		} else {
																		}
																		*(*int32)(unsafe.Add(mBase, uint32(v130)+4)) = v142
																		v164 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
																		*(*int32)(unsafe.Add(mBase, uint32(v130))) = v164
																		*(*int32)(unsafe.Add(mBase, uint32(v164)+4)) = v130
																		*(*int32)(unsafe.Add(mBase, uint32(v142))) = v130
																	} else {
																	}
																	v170 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
																	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+68))
																	v173 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[6]))
																	v178 = base.I32_div_s(v173+v171*int32(15), int32(16))
																	v180 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
																	*(*int32)(unsafe.Add(mBase, uint32(v180)+68)) = v178
																	v183 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
																	v184 = int32(0)
																	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v183))), uint32(v184))
																	v188 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[7]))
																	if v188 != 0 {
																		v190 = F_pgmem_kill(m, v188, int32(12))
																		mBase = m.M
																	} else {
																	}
																	return
																}
															} else {
																v155 = *(*int32)(unsafe.Add(mBase, uint32(v130)+616))
																if v155 == int32(0) {
																	v158 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
																	if v158 == int32(0) {
																		*(*int32)(unsafe.Add(mBase, uint32(v142)+4)) = v142
																		*(*int32)(unsafe.Add(mBase, uint32(v142))) = v142
																	} else {
																	}
																	*(*int32)(unsafe.Add(mBase, uint32(v130)+4)) = v142
																	v164 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
																	*(*int32)(unsafe.Add(mBase, uint32(v130))) = v164
																	*(*int32)(unsafe.Add(mBase, uint32(v164)+4)) = v130
																	*(*int32)(unsafe.Add(mBase, uint32(v142))) = v130
																} else {
																}
																v170 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
																v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+68))
																v173 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[6]))
																v178 = base.I32_div_s(v173+v171*int32(15), int32(16))
																v180 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
																*(*int32)(unsafe.Add(mBase, uint32(v180)+68)) = v178
																v183 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
																v184 = int32(0)
																atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v183))), uint32(v184))
																v188 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[7]))
																if v188 != 0 {
																	v190 = F_pgmem_kill(m, v188, int32(12))
																	mBase = m.M
																} else {
																}
																return
															}
														}
													}
												}
											} else {
												v95 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
												if v95 == int32(0) {
													*(*int32)(unsafe.Add(mBase, uint32(v82))) = v82
													v99 = v82
												} else {
													v99 = v95
												}
												*(*int32)(unsafe.Add(mBase, uint32(v43))) = v82
												*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v99
												*(*int32)(unsafe.Add(mBase, uint32(v99))) = v43
												*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = v43
												v105 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
												v106 = int32(0)
												atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v105))), uint32(v106))
												F_LWLockRelease(m, v58)
												mBase = m.M
												v117 = m.ExcPending
												if v117 != 0 {
													return
												} else {
													F_SwitchBackToLocalLatch(m)
													mBase = m.M
													v122 = m.ExcPending
													if v122 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_ProcKill[3])) = int32(_a_F_ProcKill_2)
														*(*int32)(unsafe.Add(mBase, _c_F_ProcKill[4])) = int32(-1)
														v129 = int32(_a_F_ProcKill_3)
														v130 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[0]))
														v132 = int32(0)
														*(*int32)(unsafe.Add(mBase, _c_F_ProcKill[0])) = v132
														*(*int32)(unsafe.Add(mBase, uint32(v130+int32(20))+12)) = v132
														*(*int64)(unsafe.Add(mBase, uint32(v130)+52)) = int64(4294967295)
														v140 = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(v130)+44)) = v140
														v142 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
														v144 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
														v147 = base.AtomicRmwXchg32(m, v144, v140, int32(1))
														if v147 != 0 {
															v149 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
															F_s_lock(m, v149, int32(_a_F_ProcKill_4), int32(1008), int32(_a_F_ProcKill_5))
															mBase = m.M
															v154 = m.ExcPending
															if v154 != 0 {
																return
															} else {
																v155 = *(*int32)(unsafe.Add(mBase, uint32(v130)+616))
																if v155 == int32(0) {
																	v158 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
																	if v158 == int32(0) {
																		*(*int32)(unsafe.Add(mBase, uint32(v142)+4)) = v142
																		*(*int32)(unsafe.Add(mBase, uint32(v142))) = v142
																	} else {
																	}
																	*(*int32)(unsafe.Add(mBase, uint32(v130)+4)) = v142
																	v164 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
																	*(*int32)(unsafe.Add(mBase, uint32(v130))) = v164
																	*(*int32)(unsafe.Add(mBase, uint32(v164)+4)) = v130
																	*(*int32)(unsafe.Add(mBase, uint32(v142))) = v130
																} else {
																}
																v170 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
																v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+68))
																v173 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[6]))
																v178 = base.I32_div_s(v173+v171*int32(15), int32(16))
																v180 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
																*(*int32)(unsafe.Add(mBase, uint32(v180)+68)) = v178
																v183 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
																v184 = int32(0)
																atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v183))), uint32(v184))
																v188 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[7]))
																if v188 != 0 {
																	v190 = F_pgmem_kill(m, v188, int32(12))
																	mBase = m.M
																} else {
																}
																return
															}
														} else {
															v155 = *(*int32)(unsafe.Add(mBase, uint32(v130)+616))
															if v155 == int32(0) {
																v158 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
																if v158 == int32(0) {
																	*(*int32)(unsafe.Add(mBase, uint32(v142)+4)) = v142
																	*(*int32)(unsafe.Add(mBase, uint32(v142))) = v142
																} else {
																}
																*(*int32)(unsafe.Add(mBase, uint32(v130)+4)) = v142
																v164 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
																*(*int32)(unsafe.Add(mBase, uint32(v130))) = v164
																*(*int32)(unsafe.Add(mBase, uint32(v164)+4)) = v130
																*(*int32)(unsafe.Add(mBase, uint32(v142))) = v130
															} else {
															}
															v170 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
															v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+68))
															v173 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[6]))
															v178 = base.I32_div_s(v173+v171*int32(15), int32(16))
															v180 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
															*(*int32)(unsafe.Add(mBase, uint32(v180)+68)) = v178
															v183 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
															v184 = int32(0)
															atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v183))), uint32(v184))
															v188 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[7]))
															if v188 != 0 {
																v190 = F_pgmem_kill(m, v188, int32(12))
																mBase = m.M
															} else {
															}
															return
														}
													}
												}
											}
										}
									} else {
										v110 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[0]))
										if v43 == v110 {
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v110)+616)) = int32(0)
										}
										F_LWLockRelease(m, v58)
										mBase = m.M
										v117 = m.ExcPending
										if v117 != 0 {
											return
										} else {
											F_SwitchBackToLocalLatch(m)
											mBase = m.M
											v122 = m.ExcPending
											if v122 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_ProcKill[3])) = int32(_a_F_ProcKill_2)
												*(*int32)(unsafe.Add(mBase, _c_F_ProcKill[4])) = int32(-1)
												v129 = int32(_a_F_ProcKill_3)
												v130 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[0]))
												v132 = int32(0)
												*(*int32)(unsafe.Add(mBase, _c_F_ProcKill[0])) = v132
												*(*int32)(unsafe.Add(mBase, uint32(v130+int32(20))+12)) = v132
												*(*int64)(unsafe.Add(mBase, uint32(v130)+52)) = int64(4294967295)
												v140 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v130)+44)) = v140
												v142 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
												v144 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
												v147 = base.AtomicRmwXchg32(m, v144, v140, int32(1))
												if v147 != 0 {
													v149 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
													F_s_lock(m, v149, int32(_a_F_ProcKill_4), int32(1008), int32(_a_F_ProcKill_5))
													mBase = m.M
													v154 = m.ExcPending
													if v154 != 0 {
														return
													} else {
														v155 = *(*int32)(unsafe.Add(mBase, uint32(v130)+616))
														if v155 == int32(0) {
															v158 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
															if v158 == int32(0) {
																*(*int32)(unsafe.Add(mBase, uint32(v142)+4)) = v142
																*(*int32)(unsafe.Add(mBase, uint32(v142))) = v142
															} else {
															}
															*(*int32)(unsafe.Add(mBase, uint32(v130)+4)) = v142
															v164 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
															*(*int32)(unsafe.Add(mBase, uint32(v130))) = v164
															*(*int32)(unsafe.Add(mBase, uint32(v164)+4)) = v130
															*(*int32)(unsafe.Add(mBase, uint32(v142))) = v130
														} else {
														}
														v170 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
														v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+68))
														v173 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[6]))
														v178 = base.I32_div_s(v173+v171*int32(15), int32(16))
														v180 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
														*(*int32)(unsafe.Add(mBase, uint32(v180)+68)) = v178
														v183 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
														v184 = int32(0)
														atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v183))), uint32(v184))
														v188 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[7]))
														if v188 != 0 {
															v190 = F_pgmem_kill(m, v188, int32(12))
															mBase = m.M
														} else {
														}
														return
													}
												} else {
													v155 = *(*int32)(unsafe.Add(mBase, uint32(v130)+616))
													if v155 == int32(0) {
														v158 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
														if v158 == int32(0) {
															*(*int32)(unsafe.Add(mBase, uint32(v142)+4)) = v142
															*(*int32)(unsafe.Add(mBase, uint32(v142))) = v142
														} else {
														}
														*(*int32)(unsafe.Add(mBase, uint32(v130)+4)) = v142
														v164 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
														*(*int32)(unsafe.Add(mBase, uint32(v130))) = v164
														*(*int32)(unsafe.Add(mBase, uint32(v164)+4)) = v130
														*(*int32)(unsafe.Add(mBase, uint32(v142))) = v130
													} else {
													}
													v170 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
													v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+68))
													v173 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[6]))
													v178 = base.I32_div_s(v173+v171*int32(15), int32(16))
													v180 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
													*(*int32)(unsafe.Add(mBase, uint32(v180)+68)) = v178
													v183 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
													v184 = int32(0)
													atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v183))), uint32(v184))
													v188 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[7]))
													if v188 != 0 {
														v190 = F_pgmem_kill(m, v188, int32(12))
														mBase = m.M
													} else {
													}
													return
												}
											}
										}
									}
								}
							} else {
								F_SwitchBackToLocalLatch(m)
								mBase = m.M
								v122 = m.ExcPending
								if v122 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_ProcKill[3])) = int32(_a_F_ProcKill_2)
									*(*int32)(unsafe.Add(mBase, _c_F_ProcKill[4])) = int32(-1)
									v129 = int32(_a_F_ProcKill_3)
									v130 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[0]))
									v132 = int32(0)
									*(*int32)(unsafe.Add(mBase, _c_F_ProcKill[0])) = v132
									*(*int32)(unsafe.Add(mBase, uint32(v130+int32(20))+12)) = v132
									*(*int64)(unsafe.Add(mBase, uint32(v130)+52)) = int64(4294967295)
									v140 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v130)+44)) = v140
									v142 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
									v144 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
									v147 = base.AtomicRmwXchg32(m, v144, v140, int32(1))
									if v147 != 0 {
										v149 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
										F_s_lock(m, v149, int32(_a_F_ProcKill_4), int32(1008), int32(_a_F_ProcKill_5))
										mBase = m.M
										v154 = m.ExcPending
										if v154 != 0 {
											return
										} else {
											v155 = *(*int32)(unsafe.Add(mBase, uint32(v130)+616))
											if v155 == int32(0) {
												v158 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
												if v158 == int32(0) {
													*(*int32)(unsafe.Add(mBase, uint32(v142)+4)) = v142
													*(*int32)(unsafe.Add(mBase, uint32(v142))) = v142
												} else {
												}
												*(*int32)(unsafe.Add(mBase, uint32(v130)+4)) = v142
												v164 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
												*(*int32)(unsafe.Add(mBase, uint32(v130))) = v164
												*(*int32)(unsafe.Add(mBase, uint32(v164)+4)) = v130
												*(*int32)(unsafe.Add(mBase, uint32(v142))) = v130
											} else {
											}
											v170 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
											v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+68))
											v173 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[6]))
											v178 = base.I32_div_s(v173+v171*int32(15), int32(16))
											v180 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
											*(*int32)(unsafe.Add(mBase, uint32(v180)+68)) = v178
											v183 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
											v184 = int32(0)
											atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v183))), uint32(v184))
											v188 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[7]))
											if v188 != 0 {
												v190 = F_pgmem_kill(m, v188, int32(12))
												mBase = m.M
											} else {
											}
											return
										}
									} else {
										v155 = *(*int32)(unsafe.Add(mBase, uint32(v130)+616))
										if v155 == int32(0) {
											v158 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
											if v158 == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(v142)+4)) = v142
												*(*int32)(unsafe.Add(mBase, uint32(v142))) = v142
											} else {
											}
											*(*int32)(unsafe.Add(mBase, uint32(v130)+4)) = v142
											v164 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
											*(*int32)(unsafe.Add(mBase, uint32(v130))) = v164
											*(*int32)(unsafe.Add(mBase, uint32(v164)+4)) = v130
											*(*int32)(unsafe.Add(mBase, uint32(v142))) = v130
										} else {
										}
										v170 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
										v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+68))
										v173 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[6]))
										v178 = base.I32_div_s(v173+v171*int32(15), int32(16))
										v180 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
										*(*int32)(unsafe.Add(mBase, uint32(v180)+68)) = v178
										v183 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
										v184 = int32(0)
										atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v183))), uint32(v184))
										v188 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[7]))
										if v188 != 0 {
											v190 = F_pgmem_kill(m, v188, int32(12))
											mBase = m.M
										} else {
										}
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
			v38 = m.ExcPending
			if v38 != 0 {
				return
			} else {
				F_ConditionVariableCancelSleep(m)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					v42 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[0]))
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+616))
					if v43 != 0 {
						v45 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[1]))
						v47 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
						v51 = base.I32_div_s(v43-v48, int32(640))
						v53 = base.I32_rem_s(v51, int32(16))
						v58 = v45 + v53<<(uint(int32(7))%32) + int32(_a_F_ProcKill_1)
						v60 = F_LWLockAcquire(m, v58, int32(0))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return
						} else {
							v63 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[0]))
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+628))
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)+632))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = v65
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v63)+628))
							*(*int32)(unsafe.Add(mBase, uint32(v65))) = v67
							v69 = *(*int32)(unsafe.Add(mBase, uint32(v43)+624))
							if v69 != v43+int32(620) {
								v74 = v69
							} else {
								v74 = int32(0)
							}
							if v74 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(v43)+616)) = int32(0)
								v80 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[0]))
								if v43 == v80 {
									F_LWLockRelease(m, v58)
									mBase = m.M
									v117 = m.ExcPending
									if v117 != 0 {
										return
									} else {
										F_SwitchBackToLocalLatch(m)
										mBase = m.M
										v122 = m.ExcPending
										if v122 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_ProcKill[3])) = int32(_a_F_ProcKill_2)
											*(*int32)(unsafe.Add(mBase, _c_F_ProcKill[4])) = int32(-1)
											v129 = int32(_a_F_ProcKill_3)
											v130 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[0]))
											v132 = int32(0)
											*(*int32)(unsafe.Add(mBase, _c_F_ProcKill[0])) = v132
											*(*int32)(unsafe.Add(mBase, uint32(v130+int32(20))+12)) = v132
											*(*int64)(unsafe.Add(mBase, uint32(v130)+52)) = int64(4294967295)
											v140 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v130)+44)) = v140
											v142 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
											v144 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
											v147 = base.AtomicRmwXchg32(m, v144, v140, int32(1))
											if v147 != 0 {
												v149 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
												F_s_lock(m, v149, int32(_a_F_ProcKill_4), int32(1008), int32(_a_F_ProcKill_5))
												mBase = m.M
												v154 = m.ExcPending
												if v154 != 0 {
													return
												} else {
													v155 = *(*int32)(unsafe.Add(mBase, uint32(v130)+616))
													if v155 == int32(0) {
														v158 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
														if v158 == int32(0) {
															*(*int32)(unsafe.Add(mBase, uint32(v142)+4)) = v142
															*(*int32)(unsafe.Add(mBase, uint32(v142))) = v142
														} else {
														}
														*(*int32)(unsafe.Add(mBase, uint32(v130)+4)) = v142
														v164 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
														*(*int32)(unsafe.Add(mBase, uint32(v130))) = v164
														*(*int32)(unsafe.Add(mBase, uint32(v164)+4)) = v130
														*(*int32)(unsafe.Add(mBase, uint32(v142))) = v130
													} else {
													}
													v170 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
													v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+68))
													v173 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[6]))
													v178 = base.I32_div_s(v173+v171*int32(15), int32(16))
													v180 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
													*(*int32)(unsafe.Add(mBase, uint32(v180)+68)) = v178
													v183 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
													v184 = int32(0)
													atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v183))), uint32(v184))
													v188 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[7]))
													if v188 != 0 {
														v190 = F_pgmem_kill(m, v188, int32(12))
														mBase = m.M
													} else {
													}
													return
												}
											} else {
												v155 = *(*int32)(unsafe.Add(mBase, uint32(v130)+616))
												if v155 == int32(0) {
													v158 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
													if v158 == int32(0) {
														*(*int32)(unsafe.Add(mBase, uint32(v142)+4)) = v142
														*(*int32)(unsafe.Add(mBase, uint32(v142))) = v142
													} else {
													}
													*(*int32)(unsafe.Add(mBase, uint32(v130)+4)) = v142
													v164 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
													*(*int32)(unsafe.Add(mBase, uint32(v130))) = v164
													*(*int32)(unsafe.Add(mBase, uint32(v164)+4)) = v130
													*(*int32)(unsafe.Add(mBase, uint32(v142))) = v130
												} else {
												}
												v170 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
												v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+68))
												v173 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[6]))
												v178 = base.I32_div_s(v173+v171*int32(15), int32(16))
												v180 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
												*(*int32)(unsafe.Add(mBase, uint32(v180)+68)) = v178
												v183 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
												v184 = int32(0)
												atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v183))), uint32(v184))
												v188 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[7]))
												if v188 != 0 {
													v190 = F_pgmem_kill(m, v188, int32(12))
													mBase = m.M
												} else {
												}
												return
											}
										}
									}
								} else {
									v82 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
									v84 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
									v87 = base.AtomicRmwXchg32(m, v84, int32(0), int32(1))
									if v87 != 0 {
										v89 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
										F_s_lock(m, v89, int32(_a_F_ProcKill_4), int32(975), int32(_a_F_ProcKill_5))
										mBase = m.M
										v94 = m.ExcPending
										if v94 != 0 {
											return
										} else {
											v95 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
											if v95 == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(v82))) = v82
												v99 = v82
											} else {
												v99 = v95
											}
											*(*int32)(unsafe.Add(mBase, uint32(v43))) = v82
											*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v99
											*(*int32)(unsafe.Add(mBase, uint32(v99))) = v43
											*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = v43
											v105 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
											v106 = int32(0)
											atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v105))), uint32(v106))
											F_LWLockRelease(m, v58)
											mBase = m.M
											v117 = m.ExcPending
											if v117 != 0 {
												return
											} else {
												F_SwitchBackToLocalLatch(m)
												mBase = m.M
												v122 = m.ExcPending
												if v122 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_ProcKill[3])) = int32(_a_F_ProcKill_2)
													*(*int32)(unsafe.Add(mBase, _c_F_ProcKill[4])) = int32(-1)
													v129 = int32(_a_F_ProcKill_3)
													v130 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[0]))
													v132 = int32(0)
													*(*int32)(unsafe.Add(mBase, _c_F_ProcKill[0])) = v132
													*(*int32)(unsafe.Add(mBase, uint32(v130+int32(20))+12)) = v132
													*(*int64)(unsafe.Add(mBase, uint32(v130)+52)) = int64(4294967295)
													v140 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v130)+44)) = v140
													v142 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
													v144 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
													v147 = base.AtomicRmwXchg32(m, v144, v140, int32(1))
													if v147 != 0 {
														v149 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
														F_s_lock(m, v149, int32(_a_F_ProcKill_4), int32(1008), int32(_a_F_ProcKill_5))
														mBase = m.M
														v154 = m.ExcPending
														if v154 != 0 {
															return
														} else {
															v155 = *(*int32)(unsafe.Add(mBase, uint32(v130)+616))
															if v155 == int32(0) {
																v158 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
																if v158 == int32(0) {
																	*(*int32)(unsafe.Add(mBase, uint32(v142)+4)) = v142
																	*(*int32)(unsafe.Add(mBase, uint32(v142))) = v142
																} else {
																}
																*(*int32)(unsafe.Add(mBase, uint32(v130)+4)) = v142
																v164 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
																*(*int32)(unsafe.Add(mBase, uint32(v130))) = v164
																*(*int32)(unsafe.Add(mBase, uint32(v164)+4)) = v130
																*(*int32)(unsafe.Add(mBase, uint32(v142))) = v130
															} else {
															}
															v170 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
															v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+68))
															v173 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[6]))
															v178 = base.I32_div_s(v173+v171*int32(15), int32(16))
															v180 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
															*(*int32)(unsafe.Add(mBase, uint32(v180)+68)) = v178
															v183 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
															v184 = int32(0)
															atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v183))), uint32(v184))
															v188 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[7]))
															if v188 != 0 {
																v190 = F_pgmem_kill(m, v188, int32(12))
																mBase = m.M
															} else {
															}
															return
														}
													} else {
														v155 = *(*int32)(unsafe.Add(mBase, uint32(v130)+616))
														if v155 == int32(0) {
															v158 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
															if v158 == int32(0) {
																*(*int32)(unsafe.Add(mBase, uint32(v142)+4)) = v142
																*(*int32)(unsafe.Add(mBase, uint32(v142))) = v142
															} else {
															}
															*(*int32)(unsafe.Add(mBase, uint32(v130)+4)) = v142
															v164 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
															*(*int32)(unsafe.Add(mBase, uint32(v130))) = v164
															*(*int32)(unsafe.Add(mBase, uint32(v164)+4)) = v130
															*(*int32)(unsafe.Add(mBase, uint32(v142))) = v130
														} else {
														}
														v170 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
														v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+68))
														v173 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[6]))
														v178 = base.I32_div_s(v173+v171*int32(15), int32(16))
														v180 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
														*(*int32)(unsafe.Add(mBase, uint32(v180)+68)) = v178
														v183 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
														v184 = int32(0)
														atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v183))), uint32(v184))
														v188 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[7]))
														if v188 != 0 {
															v190 = F_pgmem_kill(m, v188, int32(12))
															mBase = m.M
														} else {
														}
														return
													}
												}
											}
										}
									} else {
										v95 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
										if v95 == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(v82))) = v82
											v99 = v82
										} else {
											v99 = v95
										}
										*(*int32)(unsafe.Add(mBase, uint32(v43))) = v82
										*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v99
										*(*int32)(unsafe.Add(mBase, uint32(v99))) = v43
										*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = v43
										v105 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
										v106 = int32(0)
										atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v105))), uint32(v106))
										F_LWLockRelease(m, v58)
										mBase = m.M
										v117 = m.ExcPending
										if v117 != 0 {
											return
										} else {
											F_SwitchBackToLocalLatch(m)
											mBase = m.M
											v122 = m.ExcPending
											if v122 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_ProcKill[3])) = int32(_a_F_ProcKill_2)
												*(*int32)(unsafe.Add(mBase, _c_F_ProcKill[4])) = int32(-1)
												v129 = int32(_a_F_ProcKill_3)
												v130 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[0]))
												v132 = int32(0)
												*(*int32)(unsafe.Add(mBase, _c_F_ProcKill[0])) = v132
												*(*int32)(unsafe.Add(mBase, uint32(v130+int32(20))+12)) = v132
												*(*int64)(unsafe.Add(mBase, uint32(v130)+52)) = int64(4294967295)
												v140 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v130)+44)) = v140
												v142 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
												v144 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
												v147 = base.AtomicRmwXchg32(m, v144, v140, int32(1))
												if v147 != 0 {
													v149 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
													F_s_lock(m, v149, int32(_a_F_ProcKill_4), int32(1008), int32(_a_F_ProcKill_5))
													mBase = m.M
													v154 = m.ExcPending
													if v154 != 0 {
														return
													} else {
														v155 = *(*int32)(unsafe.Add(mBase, uint32(v130)+616))
														if v155 == int32(0) {
															v158 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
															if v158 == int32(0) {
																*(*int32)(unsafe.Add(mBase, uint32(v142)+4)) = v142
																*(*int32)(unsafe.Add(mBase, uint32(v142))) = v142
															} else {
															}
															*(*int32)(unsafe.Add(mBase, uint32(v130)+4)) = v142
															v164 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
															*(*int32)(unsafe.Add(mBase, uint32(v130))) = v164
															*(*int32)(unsafe.Add(mBase, uint32(v164)+4)) = v130
															*(*int32)(unsafe.Add(mBase, uint32(v142))) = v130
														} else {
														}
														v170 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
														v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+68))
														v173 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[6]))
														v178 = base.I32_div_s(v173+v171*int32(15), int32(16))
														v180 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
														*(*int32)(unsafe.Add(mBase, uint32(v180)+68)) = v178
														v183 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
														v184 = int32(0)
														atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v183))), uint32(v184))
														v188 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[7]))
														if v188 != 0 {
															v190 = F_pgmem_kill(m, v188, int32(12))
															mBase = m.M
														} else {
														}
														return
													}
												} else {
													v155 = *(*int32)(unsafe.Add(mBase, uint32(v130)+616))
													if v155 == int32(0) {
														v158 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
														if v158 == int32(0) {
															*(*int32)(unsafe.Add(mBase, uint32(v142)+4)) = v142
															*(*int32)(unsafe.Add(mBase, uint32(v142))) = v142
														} else {
														}
														*(*int32)(unsafe.Add(mBase, uint32(v130)+4)) = v142
														v164 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
														*(*int32)(unsafe.Add(mBase, uint32(v130))) = v164
														*(*int32)(unsafe.Add(mBase, uint32(v164)+4)) = v130
														*(*int32)(unsafe.Add(mBase, uint32(v142))) = v130
													} else {
													}
													v170 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
													v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+68))
													v173 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[6]))
													v178 = base.I32_div_s(v173+v171*int32(15), int32(16))
													v180 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
													*(*int32)(unsafe.Add(mBase, uint32(v180)+68)) = v178
													v183 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
													v184 = int32(0)
													atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v183))), uint32(v184))
													v188 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[7]))
													if v188 != 0 {
														v190 = F_pgmem_kill(m, v188, int32(12))
														mBase = m.M
													} else {
													}
													return
												}
											}
										}
									}
								}
							} else {
								v110 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[0]))
								if v43 == v110 {
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v110)+616)) = int32(0)
								}
								F_LWLockRelease(m, v58)
								mBase = m.M
								v117 = m.ExcPending
								if v117 != 0 {
									return
								} else {
									F_SwitchBackToLocalLatch(m)
									mBase = m.M
									v122 = m.ExcPending
									if v122 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_ProcKill[3])) = int32(_a_F_ProcKill_2)
										*(*int32)(unsafe.Add(mBase, _c_F_ProcKill[4])) = int32(-1)
										v129 = int32(_a_F_ProcKill_3)
										v130 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[0]))
										v132 = int32(0)
										*(*int32)(unsafe.Add(mBase, _c_F_ProcKill[0])) = v132
										*(*int32)(unsafe.Add(mBase, uint32(v130+int32(20))+12)) = v132
										*(*int64)(unsafe.Add(mBase, uint32(v130)+52)) = int64(4294967295)
										v140 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v130)+44)) = v140
										v142 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
										v144 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
										v147 = base.AtomicRmwXchg32(m, v144, v140, int32(1))
										if v147 != 0 {
											v149 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
											F_s_lock(m, v149, int32(_a_F_ProcKill_4), int32(1008), int32(_a_F_ProcKill_5))
											mBase = m.M
											v154 = m.ExcPending
											if v154 != 0 {
												return
											} else {
												v155 = *(*int32)(unsafe.Add(mBase, uint32(v130)+616))
												if v155 == int32(0) {
													v158 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
													if v158 == int32(0) {
														*(*int32)(unsafe.Add(mBase, uint32(v142)+4)) = v142
														*(*int32)(unsafe.Add(mBase, uint32(v142))) = v142
													} else {
													}
													*(*int32)(unsafe.Add(mBase, uint32(v130)+4)) = v142
													v164 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
													*(*int32)(unsafe.Add(mBase, uint32(v130))) = v164
													*(*int32)(unsafe.Add(mBase, uint32(v164)+4)) = v130
													*(*int32)(unsafe.Add(mBase, uint32(v142))) = v130
												} else {
												}
												v170 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
												v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+68))
												v173 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[6]))
												v178 = base.I32_div_s(v173+v171*int32(15), int32(16))
												v180 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
												*(*int32)(unsafe.Add(mBase, uint32(v180)+68)) = v178
												v183 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
												v184 = int32(0)
												atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v183))), uint32(v184))
												v188 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[7]))
												if v188 != 0 {
													v190 = F_pgmem_kill(m, v188, int32(12))
													mBase = m.M
												} else {
												}
												return
											}
										} else {
											v155 = *(*int32)(unsafe.Add(mBase, uint32(v130)+616))
											if v155 == int32(0) {
												v158 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
												if v158 == int32(0) {
													*(*int32)(unsafe.Add(mBase, uint32(v142)+4)) = v142
													*(*int32)(unsafe.Add(mBase, uint32(v142))) = v142
												} else {
												}
												*(*int32)(unsafe.Add(mBase, uint32(v130)+4)) = v142
												v164 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
												*(*int32)(unsafe.Add(mBase, uint32(v130))) = v164
												*(*int32)(unsafe.Add(mBase, uint32(v164)+4)) = v130
												*(*int32)(unsafe.Add(mBase, uint32(v142))) = v130
											} else {
											}
											v170 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
											v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+68))
											v173 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[6]))
											v178 = base.I32_div_s(v173+v171*int32(15), int32(16))
											v180 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
											*(*int32)(unsafe.Add(mBase, uint32(v180)+68)) = v178
											v183 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
											v184 = int32(0)
											atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v183))), uint32(v184))
											v188 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[7]))
											if v188 != 0 {
												v190 = F_pgmem_kill(m, v188, int32(12))
												mBase = m.M
											} else {
											}
											return
										}
									}
								}
							}
						}
					} else {
						F_SwitchBackToLocalLatch(m)
						mBase = m.M
						v122 = m.ExcPending
						if v122 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_ProcKill[3])) = int32(_a_F_ProcKill_2)
							*(*int32)(unsafe.Add(mBase, _c_F_ProcKill[4])) = int32(-1)
							v129 = int32(_a_F_ProcKill_3)
							v130 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[0]))
							v132 = int32(0)
							*(*int32)(unsafe.Add(mBase, _c_F_ProcKill[0])) = v132
							*(*int32)(unsafe.Add(mBase, uint32(v130+int32(20))+12)) = v132
							*(*int64)(unsafe.Add(mBase, uint32(v130)+52)) = int64(4294967295)
							v140 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v130)+44)) = v140
							v142 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
							v144 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
							v147 = base.AtomicRmwXchg32(m, v144, v140, int32(1))
							if v147 != 0 {
								v149 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
								F_s_lock(m, v149, int32(_a_F_ProcKill_4), int32(1008), int32(_a_F_ProcKill_5))
								mBase = m.M
								v154 = m.ExcPending
								if v154 != 0 {
									return
								} else {
									v155 = *(*int32)(unsafe.Add(mBase, uint32(v130)+616))
									if v155 == int32(0) {
										v158 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
										if v158 == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(v142)+4)) = v142
											*(*int32)(unsafe.Add(mBase, uint32(v142))) = v142
										} else {
										}
										*(*int32)(unsafe.Add(mBase, uint32(v130)+4)) = v142
										v164 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
										*(*int32)(unsafe.Add(mBase, uint32(v130))) = v164
										*(*int32)(unsafe.Add(mBase, uint32(v164)+4)) = v130
										*(*int32)(unsafe.Add(mBase, uint32(v142))) = v130
									} else {
									}
									v170 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
									v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+68))
									v173 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[6]))
									v178 = base.I32_div_s(v173+v171*int32(15), int32(16))
									v180 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
									*(*int32)(unsafe.Add(mBase, uint32(v180)+68)) = v178
									v183 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
									v184 = int32(0)
									atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v183))), uint32(v184))
									v188 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[7]))
									if v188 != 0 {
										v190 = F_pgmem_kill(m, v188, int32(12))
										mBase = m.M
									} else {
									}
									return
								}
							} else {
								v155 = *(*int32)(unsafe.Add(mBase, uint32(v130)+616))
								if v155 == int32(0) {
									v158 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
									if v158 == int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(v142)+4)) = v142
										*(*int32)(unsafe.Add(mBase, uint32(v142))) = v142
									} else {
									}
									*(*int32)(unsafe.Add(mBase, uint32(v130)+4)) = v142
									v164 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
									*(*int32)(unsafe.Add(mBase, uint32(v130))) = v164
									*(*int32)(unsafe.Add(mBase, uint32(v164)+4)) = v130
									*(*int32)(unsafe.Add(mBase, uint32(v142))) = v130
								} else {
								}
								v170 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
								v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+68))
								v173 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[6]))
								v178 = base.I32_div_s(v173+v171*int32(15), int32(16))
								v180 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[2]))
								*(*int32)(unsafe.Add(mBase, uint32(v180)+68)) = v178
								v183 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[5]))
								v184 = int32(0)
								atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v183))), uint32(v184))
								v188 = *(*int32)(unsafe.Add(mBase, _c_F_ProcKill[7]))
								if v188 != 0 {
									v190 = F_pgmem_kill(m, v188, int32(12))
									mBase = m.M
								} else {
								}
								return
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(23), int32(0))
		mBase = m.M
		v194 = m.ExcPending
		if v194 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_ProcKill_6), int32(0))
			mBase = m.M
			v198 = m.ExcPending
			if v198 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_ProcKill_4), int32(928), int32(_a_F_ProcKill_5))
				mBase = m.M
				v203 = m.ExcPending
				if v203 != 0 {
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
