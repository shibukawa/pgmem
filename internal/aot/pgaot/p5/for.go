package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CheckForSerializableConflictIn(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_CheckForSerializableConflictIn[0]))
	if v12 == int32(0) {
		m.G0 = v9 + int32(16)
		return
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		if base.Ui32(v15) < base.Ui32(int32(_a_F_CheckForSerializableConflictIn_0)) {
			m.G0 = v9 + int32(16)
			return
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+118)))
			if v19 == int32(116) {
				m.G0 = v9 + int32(16)
				return
			} else {
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+108)))
				if v22&int32(8) != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return
					} else {
						F_errcode(m, int32(16777220))
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return
						} else {
							F_errmsg(m, int32(_a_F_CheckForSerializableConflictIn_1), int32(0))
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return
							} else {
								F_errdetail_internal(m, int32(_a_F_CheckForSerializableConflictIn_2), int32(0))
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return
								} else {
									F_errhint(m, int32(_a_F_CheckForSerializableConflictIn_3), int32(0))
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_CheckForSerializableConflictIn_4), int32(_a_F_CheckForSerializableConflictIn_5), int32(_a_F_CheckForSerializableConflictIn_6))
										mBase = m.M
										v94 = m.ExcPending
										if v94 != 0 {
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
				} else {
					v26 = int32(1)
					*(*uint8)(unsafe.Add(mBase, _c_F_CheckForSerializableConflictIn[1])) = uint8(v26)
					if l1 != 0 {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v15
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v28
						v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
						v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v31 | v32<<(uint(int32(16))%32)
						v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v37
						F_CheckTargetForConflictsIn(m, v9)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
							v42 = v41
							if l2 != int32(-1) {
								v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l2
								*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v42
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v46
								F_CheckTargetForConflictsIn(m, v9)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return
								} else {
									v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
									v56 = v54
									v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = int64(4294967295)
									*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v56
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = v57
									F_CheckTargetForConflictsIn(m, v9)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return
									} else {
										m.G0 = v9 + int32(16)
										return
									}
								}
							} else {
								v56 = v42
								v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = int64(4294967295)
								*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v56
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v57
								F_CheckTargetForConflictsIn(m, v9)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									m.G0 = v9 + int32(16)
									return
								}
							}
						}
					} else {
						v42 = v15
						if l2 != int32(-1) {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l2
							*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v42
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v46
							F_CheckTargetForConflictsIn(m, v9)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return
							} else {
								v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
								v56 = v54
								v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = int64(4294967295)
								*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v56
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v57
								F_CheckTargetForConflictsIn(m, v9)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									m.G0 = v9 + int32(16)
									return
								}
							}
						} else {
							v56 = v42
							v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = int64(4294967295)
							*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v56
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v57
							F_CheckTargetForConflictsIn(m, v9)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								m.G0 = v9 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_CheckForSerializableConflictOutNeeded(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_CheckForSerializableConflictOutNeeded[0]))
	if v6 == v3 {
		v56 = v3
		return v56
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		switch v9 {
		case 0, 5:
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+108))
			if v10&int32(128) != 0 {
				F_ReleasePredicateLocks(m, int32(0), int32(1))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					v56 = v3
					return v56
				}
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				if base.Ui32(v19) < base.Ui32(int32(_a_F_CheckForSerializableConflictOutNeeded_0)) {
					v56 = v3
					return v56
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+118)))
					if v23 == int32(116) {
						v56 = v3
						return v56
					} else {
						if v10&int32(8) == int32(0) {
							v56 = int32(1)
							return v56
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(16777220))
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_CheckForSerializableConflictOutNeeded_1), int32(0))
									mBase = m.M
									v41 = m.ExcPending
									if v41 != 0 {
										return int32(0)
									} else {
										F_errdetail_internal(m, int32(_a_F_CheckForSerializableConflictOutNeeded_2), int32(0))
										mBase = m.M
										v45 = m.ExcPending
										if v45 != 0 {
											return int32(0)
										} else {
											F_errhint(m, int32(_a_F_CheckForSerializableConflictOutNeeded_3), int32(0))
											mBase = m.M
											v49 = m.ExcPending
											if v49 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_CheckForSerializableConflictOutNeeded_4), int32(4003), int32(_a_F_CheckForSerializableConflictOutNeeded_5))
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
								}
							}
						}
					}
				}
			}
		default:
			v56 = v3
			return v56
		}
	}
}
func F_WaitForLockersMultiple(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v58 int64
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int64
	_ = v155
	var v158 int32
	_ = v158
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
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v199 int64
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v251 int64
	_ = v251
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v286 int32
	_ = v286
	var v289 int64
	_ = v289
	var v291 int64
	_ = v291
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v444 int64
	_ = v444
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v484 int32
	_ = v484
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	if l0 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v14 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v12 - int32(-64)
	return
L4:
	;
	v54 = v4
	v58 = int64(0)
	goto L6
L5:
	;
	if l2 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if l2 != 0 {
		goto L19
	} else {
		goto L20
	}
L7:
	;
	v21 = v10 + int32(-48)
	goto L9
L8:
	;
	v21 = int32(0)
	goto L9
L9:
	;
	v26 = v4
	v27 = v4
	v28 = v4
	goto L10
L10:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31+v26<<(uint(int32(2))%32))))
	v36 = F_GetLockConflicts(m, v35, l1, v21)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v54 = v38
	v58 = base.I64_extend_i32_s(v43)
	goto L6
L12:
	;
	return
L13:
	;
	v38 = F_lappend(m, v27, v36)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	if l2 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v42 = v40
	goto L17
L16:
	;
	v42 = int32(0)
	goto L17
L17:
	;
	v43 = v42 + v28
	v45 = v26 + int32(1)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v45 < v46 {
		v26 = v45
		v27 = v38
		v28 = v43
		goto L10
	} else {
		goto L18
	}
L18:
	;
	goto L11
L19:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[0]))
	if v62 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	goto L21
L21:
	;
	if v54 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	goto L21
L23:
	;
	goto L22
L24:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[1])))
	if v66&int32(1) == int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v71 = int32(_a_F_WaitForLockersMultiple_0)
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[2]))
	v74 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[2])) = v73 + v74
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v77 + v74
	v81 = int32(0)
	v83 = int32(_a_F_WaitForLockersMultiple_1)
	v84 = base.AtomicRmwOr32(m, v81, v83, v81)
	*(*int64)(unsafe.Add(mBase, uint32(v62+int32(24))+232)) = v58
	v92 = base.AtomicRmwOr32(m, v81, v83, v81)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v93 + v74
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[2])) = v99 - v74
	goto L23
L26:
	;
	if l2 != 0 {
		goto L62
	} else {
		goto L63
	}
L27:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v105 <= int32(0) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v108 = int32(0)
	v111 = v108
	v116 = v108
	goto L29
L29:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v54)+12))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v119+v111<<(uint(int32(2))%32))))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	if v124 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L26
L31:
	;
	v129 = v123
	v131 = v116
	goto L34
L32:
	;
	v269 = v116
	goto L33
L33:
	;
	v273 = v111 + int32(1)
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v273 < v274 {
		v111 = v273
		v116 = v269
		goto L29
	} else {
		goto L61
	}
L34:
	;
	if l2 != 0 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v269 = v259
	goto L33
L36:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v129)+12))
	if v260 != 0 {
		v129 = v129 + int32(8)
		v131 = v259
		goto L34
	} else {
		goto L60
	}
L37:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v135 = int32(0)
	if v134 < v135 {
		v153 = v135
		goto L41
	} else {
		goto L42
	}
L38:
	;
	goto L39
L39:
	;
	v251 = *(*int64)(unsafe.Add(mBase, uint32(v129)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v251
	v256 = F_VirtualXactLock(m, v10+int32(-56), int32(1))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L12
	} else {
		goto L59
	}
L40:
	;
	if v153 != 0 {
		goto L47
	} else {
		goto L48
	}
L41:
	;
	goto L40
L42:
	;
	v141 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[3]))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+16))
	if base.Ui32(v142) <= base.Ui32(v134) {
		v153 = int32(0)
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	v147 = v144 + v134*int32(640)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v147)+44))
	if v149 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v150 = v147
	goto L46
L45:
	;
	v150 = int32(0)
	goto L46
L46:
	;
	v153 = v150
	goto L41
L47:
	;
	v155 = int64(*(*int32)(unsafe.Add(mBase, uint32(v153)+44)))
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[0]))
	if v158 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L48:
	;
	goto L49
L49:
	;
	v199 = *(*int64)(unsafe.Add(mBase, uint32(v129)))
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v199
	v202 = F_VirtualXactLock(m, v12, int32(1))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L12
	} else {
		goto L54
	}
L50:
	;
	goto L49
L51:
	;
	goto L50
L52:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[1])))
	if v162&int32(1) == int32(0) {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v167 = int32(_a_F_WaitForLockersMultiple_0)
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[2]))
	v170 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[2])) = v169 + v170
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	*(*int32)(unsafe.Add(mBase, uint32(v158))) = v173 + v170
	v177 = int32(0)
	v179 = int32(_a_F_WaitForLockersMultiple_1)
	v180 = base.AtomicRmwOr32(m, v177, v179, v177)
	*(*int64)(unsafe.Add(mBase, uint32(v158+int32(40))+232)) = v155
	v188 = base.AtomicRmwOr32(m, v177, v179, v177)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	*(*int32)(unsafe.Add(mBase, uint32(v158))) = v189 + v170
	v195 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[2])) = v195 - v170
	goto L51
L54:
	;
	v206 = v131 + int32(1)
	v210 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[0]))
	if v210 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v259 = v206
	goto L36
L56:
	;
	goto L55
L57:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[1])))
	if v214&int32(1) == int32(0) {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v219 = int32(_a_F_WaitForLockersMultiple_0)
	v221 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[2]))
	v222 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[2])) = v221 + v222
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	*(*int32)(unsafe.Add(mBase, uint32(v210))) = v225 + v222
	v229 = int32(0)
	v231 = int32(_a_F_WaitForLockersMultiple_1)
	v232 = base.AtomicRmwOr32(m, v229, v231, v229)
	*(*int64)(unsafe.Add(mBase, uint32(v210+int32(32))+232)) = base.I64_extend_i32_s(v206)
	v240 = base.AtomicRmwOr32(m, v229, v231, v229)
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	*(*int32)(unsafe.Add(mBase, uint32(v210))) = v241 + v222
	v247 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[2])) = v247 - v222
	goto L56
L59:
	;
	v259 = v131
	goto L36
L60:
	;
	goto L35
L61:
	;
	goto L30
L62:
	;
	v286 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v286
	v289 = *(*int64)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[5]))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+48)) = v289
	v291 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v291
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v291
	*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v291
	goto L67
L63:
	;
	goto L64
L64:
	;
	F_list_free_deep(m, v54)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L12
	} else {
		goto L82
	}
L65:
	;
	goto L64
L66:
	;
	goto L65
L67:
	;
	v311 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[0]))
	if v311 == int32(0) {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v315 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[1])))
	if v315&int32(1) == int32(0) {
		goto L66
	} else {
		goto L69
	}
L69:
	;
	v320 = int32(_a_F_WaitForLockersMultiple_0)
	v322 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[2]))
	v323 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[2])) = v322 + v323
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v311)))
	*(*int32)(unsafe.Add(mBase, uint32(v311))) = v326 + v323
	v330 = int32(0)
	v333 = base.AtomicRmwOr32(m, v330, int32(_a_F_WaitForLockersMultiple_1), v330)
	goto L71
L70:
	;
	v460 = int32(0)
	v463 = base.AtomicRmwOr32(m, v460, int32(_a_F_WaitForLockersMultiple_1), v460)
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v311)))
	v465 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v311))) = v464 + v465
	v468 = int32(_a_F_WaitForLockersMultiple_0)
	v470 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[2])) = v470 - v465
	goto L66
L71:
	;
	goto L73
L73:
	;
	goto L74
L74:
	;
	v425 = int32(0)
	v428 = int32(0)
	goto L79
L79:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(-16)+v428<<(uint(int32(2))%32))))
	v438 = int32(3)
	v444 = *(*int64)(unsafe.Add(mBase, uint32(v10+int32(-48)+v428<<(uint(v438)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v311+int32(232)+v437<<(uint(v438)%32)))) = v444
	v446 = int32(1)
	v449 = v425 + v446
	if v449 != int32(3) {
		v425 = v449
		v428 = v428 + v446
		goto L79
	} else {
		goto L81
	}
L80:
	;
	goto L70
L81:
	;
	goto L80
L82:
	;
	goto L3
}
func F_WaitForParallelWorkersToExit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
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
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v2 < v6 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L10
	} else {
		goto L15
	}
L2:
	;
	v11 = v6
	v12 = v2
	goto L5
L3:
	;
	goto L4
L4:
	;
	return
L5:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v14 == int32(0) {
		v38 = v11
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L4
L7:
	;
	v41 = v12 + int32(1)
	if v41 < v38 {
		v11 = v38
		v12 = v41
		goto L5
	} else {
		goto L14
	}
L8:
	;
	v18 = v12 << (uint(int32(3)) % 32)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v14+v18)))
	if v20 == int32(0) {
		v38 = v11
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v23 = F_WaitForBackgroundWorkerShutdown(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return
L11:
	;
	if v23 == int32(3) {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v27+v18)))
	F_pfree(m, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v32+v18))) = int32(0)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v38 = v36
	goto L7
L14:
	;
	goto L6
L15:
	;
	F_errcode(m, int32(16908741))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	F_errmsg(m, int32(_a_F_WaitForParallelWorkersToExit_0), int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(_a_F_WaitForParallelWorkersToExit_1), int32(940), int32(_a_F_WaitForParallelWorkersToExit_2))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_WaitForProcSignalBarrier(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v44 int64
	_ = v44
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int64
	_ = v79
	var v82 int64
	_ = v82
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v12 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v12 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = l0
	F_errmsg_internal(m, int32(_a_F_WaitForProcSignalBarrier_0), v8+int32(32))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForProcSignalBarrier[0]))
	v28 = v26 + int32(37)
	if int32(0) <= v28 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	F_errfinish(m, int32(_a_F_WaitForProcSignalBarrier_1), int32(431), int32(_a_F_WaitForProcSignalBarrier_2))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	v34 = v28
	goto L11
L9:
	;
	goto L10
L10:
	;
	v102 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L28
	}
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForProcSignalBarrier[1]))
	v40 = v37 + v34<<(uint(int32(7))%32)
	v41 = int64(0)
	v44 = base.AtomicRmwCmpxchg64(m, v40, int32(112), v41, v41)
	if base.Ui64(v44) < base.Ui64(l0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L10
L13:
	;
	v47 = v40 + int32(8)
	goto L16
L14:
	;
	goto L15
L15:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L26
	}
L16:
	;
	v57 = F_ConditionVariableTimedSleep(m, v40+int32(124), int32(_a_F_WaitForProcSignalBarrier_3), int32(134217770))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	goto L15
L18:
	;
	v79 = int64(0)
	v82 = base.AtomicRmwCmpxchg64(m, v47, int32(104), v79, v79)
	if base.Ui64(v82) < base.Ui64(l0) {
		goto L16
	} else {
		goto L25
	}
L19:
	;
	if v57 == int32(0) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v63 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	if v63 == int32(0) {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v67
	F_errmsg(m, int32(_a_F_WaitForProcSignalBarrier_4), v8+int32(16))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_WaitForProcSignalBarrier_1), int32(452), int32(_a_F_WaitForProcSignalBarrier_2))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	goto L18
L25:
	;
	goto L17
L26:
	;
	if int32(0) < v34 {
		v34 = v34 - int32(1)
		goto L11
	} else {
		goto L27
	}
L27:
	;
	goto L12
L28:
	;
	if v102 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8))) = l0
	F_errmsg_internal(m, int32(_a_F_WaitForProcSignalBarrier_5), v8)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v113 = int32(0)
	v116 = base.AtomicRmwOr32(m, v113, int32(_a_F_WaitForProcSignalBarrier_6), v113)
	m.G0 = v8 + int32(48)
	return
L32:
	;
	F_errfinish(m, int32(_a_F_WaitForProcSignalBarrier_1), int32(461), int32(_a_F_WaitForProcSignalBarrier_2))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L31
}
