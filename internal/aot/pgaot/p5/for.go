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
	v12 = *(*int32)(unsafe.Add(mBase, _consts[788]))
	if v12 == int32(0) {
		m.G0 = v9 + int32(16)
		return
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		if base.Ui32(v15) < base.Ui32(int32(12000)) {
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
							F_errmsg(m, int32(141355), int32(0))
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return
							} else {
								F_errdetail_internal(m, int32(608882), int32(0))
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return
								} else {
									F_errhint(m, int32(632079), int32(0))
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return
									} else {
										F_errfinish(m, int32(495507), int32(4349), int32(283908))
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
					*(*uint8)(unsafe.Add(mBase, _consts[792])) = uint8(v26)
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
	v6 = *(*int32)(unsafe.Add(mBase, _consts[788]))
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
				if base.Ui32(v19) < base.Ui32(int32(12000)) {
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
									F_errmsg(m, int32(141355), int32(0))
									mBase = m.M
									v41 = m.ExcPending
									if v41 != 0 {
										return int32(0)
									} else {
										F_errdetail_internal(m, int32(608800), int32(0))
										mBase = m.M
										v45 = m.ExcPending
										if v45 != 0 {
											return int32(0)
										} else {
											F_errhint(m, int32(632079), int32(0))
											mBase = m.M
											v49 = m.ExcPending
											if v49 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(495507), int32(4003), int32(459994))
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
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v145 int64
	_ = v145
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v179 int64
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v221 int64
	_ = v221
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v258 int32
	_ = v258
	var v261 int64
	_ = v261
	var v263 int64
	_ = v263
	var v274 int32
	_ = v274
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v412 int64
	_ = v412
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v448 int32
	_ = v448
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
		goto L5
	} else {
		goto L6
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
	if l2 != 0 {
		goto L20
	} else {
		goto L21
	}
L5:
	;
	v54 = v4
	v58 = int64(0)
	goto L4
L6:
	;
	goto L7
L7:
	;
	if l2 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v21 = v10 + int32(-48)
	goto L10
L9:
	;
	v21 = int32(0)
	goto L10
L10:
	;
	v26 = v4
	v27 = v4
	v28 = v4
	goto L11
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31+v26<<(uint(int32(2))%32))))
	v36 = F_GetLockConflicts(m, v35, l1, v21)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v54 = v38
	v58 = base.I64_extend_i32_s(v43)
	goto L4
L13:
	;
	return
L14:
	;
	v38 = F_lappend(m, v27, v36)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	if l2 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v42 = v40
	goto L18
L17:
	;
	v42 = int32(0)
	goto L18
L18:
	;
	v43 = v42 + v28
	v45 = v26 + int32(1)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v45 < v46 {
		v26 = v45
		v27 = v38
		v28 = v43
		goto L11
	} else {
		goto L19
	}
L19:
	;
	goto L12
L20:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _consts[48]))
	if v62 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	goto L22
L22:
	;
	if v54 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	goto L22
L24:
	;
	goto L23
L25:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, _consts[49])))
	if v66 != int32(1) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v69 = int32(4481700)
	v71 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v72 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[13])) = v71 + v72
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v75 + v72
	*(*int64)(unsafe.Add(mBase, uint32(v62+int32(24))+232)) = v58
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v83 + v72
	v89 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	*(*int32)(unsafe.Add(mBase, _consts[13])) = v89 - v72
	goto L24
L27:
	;
	if l2 != 0 {
		goto L63
	} else {
		goto L64
	}
L28:
	;
	v95 = int32(0)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v96 <= v95 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v101 = int32(0)
	v106 = v95
	goto L30
L30:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v54)+12))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v109+v101<<(uint(int32(2))%32))))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
	if v114 != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L27
L32:
	;
	v119 = v113
	v121 = v106
	goto L35
L33:
	;
	v241 = v106
	goto L34
L34:
	;
	v245 = v101 + int32(1)
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v245 < v246 {
		v101 = v245
		v106 = v241
		goto L30
	} else {
		goto L62
	}
L35:
	;
	if l2 != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v241 = v229
	goto L34
L37:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v119+int32(12))))
	if v234 != 0 {
		v119 = v119 + int32(8)
		v121 = v229
		goto L35
	} else {
		goto L61
	}
L38:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v125 = int32(0)
	if v124 < v125 {
		v143 = v125
		goto L42
	} else {
		goto L43
	}
L39:
	;
	goto L40
L40:
	;
	v221 = *(*int64)(unsafe.Add(mBase, uint32(v119)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v221
	v226 = F_VirtualXactLock(m, v10+int32(-56), int32(1))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L13
	} else {
		goto L60
	}
L41:
	;
	if v143 != 0 {
		goto L48
	} else {
		goto L49
	}
L42:
	;
	goto L41
L43:
	;
	v131 = *(*int32)(unsafe.Add(mBase, _consts[101]))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+16))
	if base.Ui32(v132) <= base.Ui32(v124) {
		v143 = int32(0)
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v137 = v134 + v124*int32(640)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)+44))
	if v139 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v140 = v137
	goto L47
L46:
	;
	v140 = int32(0)
	goto L47
L47:
	;
	v143 = v140
	goto L42
L48:
	;
	v145 = int64(*(*int32)(unsafe.Add(mBase, uint32(v143)+44)))
	v148 = *(*int32)(unsafe.Add(mBase, _consts[48]))
	if v148 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	goto L50
L50:
	;
	v179 = *(*int64)(unsafe.Add(mBase, uint32(v119)))
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v179
	v182 = F_VirtualXactLock(m, v12, int32(1))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L13
	} else {
		goto L55
	}
L51:
	;
	goto L50
L52:
	;
	goto L51
L53:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, _consts[49])))
	if v152 != int32(1) {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v155 = int32(4481700)
	v157 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v158 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[13])) = v157 + v158
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v148))) = v161 + v158
	*(*int64)(unsafe.Add(mBase, uint32(v148+int32(40))+232)) = v145
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v148))) = v169 + v158
	v175 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	*(*int32)(unsafe.Add(mBase, _consts[13])) = v175 - v158
	goto L52
L55:
	;
	v186 = v121 + int32(1)
	v190 = *(*int32)(unsafe.Add(mBase, _consts[48]))
	if v190 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v229 = v186
	goto L37
L57:
	;
	goto L56
L58:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, _consts[49])))
	if v194 != int32(1) {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v197 = int32(4481700)
	v199 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v200 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[13])) = v199 + v200
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
	*(*int32)(unsafe.Add(mBase, uint32(v190))) = v203 + v200
	*(*int64)(unsafe.Add(mBase, uint32(v190+int32(32))+232)) = base.I64_extend_i32_s(v186)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
	*(*int32)(unsafe.Add(mBase, uint32(v190))) = v211 + v200
	v217 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	*(*int32)(unsafe.Add(mBase, _consts[13])) = v217 - v200
	goto L57
L60:
	;
	v229 = v121
	goto L37
L61:
	;
	goto L36
L62:
	;
	goto L31
L63:
	;
	v258 = *(*int32)(unsafe.Add(mBase, _consts[715]))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v258
	v261 = *(*int64)(unsafe.Add(mBase, _consts[716]))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+48)) = v261
	v263 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v263
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v263
	*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v263
	v274 = int32(0)
	v281 = *(*int32)(unsafe.Add(mBase, _consts[48]))
	if v281 == v274 {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	goto L65
L65:
	;
	F_list_free_deep(m, v54)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L13
	} else {
		goto L83
	}
L66:
	;
	goto L65
L67:
	;
	goto L66
L68:
	;
	goto L69
L69:
	;
	v287 = int32(*(*uint8)(unsafe.Add(mBase, _consts[49])))
	if v287&int32(1) == int32(0) {
		goto L67
	} else {
		goto L70
	}
L70:
	;
	v292 = int32(4481700)
	v294 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v295 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[13])) = v294 + v295
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v281)))
	*(*int32)(unsafe.Add(mBase, uint32(v281))) = v298 + v295
	goto L72
L71:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v281)))
	v429 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v281))) = v428 + v429
	v432 = int32(4481700)
	v434 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	*(*int32)(unsafe.Add(mBase, _consts[13])) = v434 - v429
	goto L67
L72:
	;
	goto L74
L74:
	;
	goto L75
L75:
	;
	goto L79
L79:
	;
	v393 = int32(0)
	v396 = v274
	goto L80
L80:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(-16)+v396<<(uint(int32(2))%32))))
	v406 = int32(3)
	v412 = *(*int64)(unsafe.Add(mBase, uint32(v10+int32(-48)+v396<<(uint(v406)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v281+int32(232)+v405<<(uint(v406)%32)))) = v412
	v414 = int32(1)
	v417 = v393 + v414
	if v417 != int32(3) {
		v393 = v417
		v396 = v396 + v414
		goto L80
	} else {
		goto L82
	}
L81:
	;
	goto L71
L82:
	;
	goto L81
L83:
	;
	goto L3
}
func F_WaitForParallelWorkersToExit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v2 < v7 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L16
	} else {
		goto L33
	}
L2:
	;
	v11 = v7
	v13 = v2
	goto L5
L3:
	;
	goto L4
L4:
	;
	return
L5:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v16 == int32(0) {
		v101 = v11
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L4
L7:
	;
	v107 = v13 + int32(1)
	if v107 < v101 {
		v11 = v101
		v13 = v107
		goto L5
	} else {
		goto L32
	}
L8:
	;
	v20 = v13 << (uint(int32(3)) % 32)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v16+v20)))
	if v22 == int32(0) {
		v101 = v11
		goto L7
	} else {
		goto L9
	}
L9:
	;
	goto L10
L10:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	if v32 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	if v87 == int32(3) {
		goto L1
	} else {
		goto L30
	}
L12:
	;
	goto L11
L13:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _consts[82]))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v39 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v43 = F_LWLockAcquire(m, v39+int32(4224), int32(1))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L16
	} else {
		goto L18
	}
L16:
	;
	return
L17:
	;
	goto L15
L18:
	;
	v45 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
	v48 = v36 + v37*int32(1480)
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v48)+24))
	if v45 == v49 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v64 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v64+int32(4224))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L16
	} else {
		goto L25
	}
L20:
	;
	v52 = v48 + int32(16)
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	if v53 != 0 {
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v56+int32(4224))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L16
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	v87 = int32(2)
	goto L12
L25:
	;
	if v62 == int32(0) {
		v87 = int32(2)
		goto L12
	} else {
		goto L26
	}
L26:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _consts[83]))
	v78 = F_WaitLatch(m, v74, int32(17), int32(0), int32(134217733))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L16
	} else {
		goto L27
	}
L27:
	;
	if v78&int32(16) != 0 {
		v87 = int32(3)
		goto L12
	} else {
		goto L28
	}
L28:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _consts[83]))
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = int32(0)
	goto L29
L29:
	;
	goto L10
L30:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v90+v20)))
	F_pfree(m, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L16
	} else {
		goto L31
	}
L31:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v95+v20))) = int32(0)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v101 = v99
	goto L7
L32:
	;
	goto L6
L33:
	;
	F_errcode(m, int32(16908741))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L16
	} else {
		goto L34
	}
L34:
	;
	F_errmsg(m, int32(255713), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L16
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(494337), int32(940), int32(99180))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L16
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_WaitForProcSignalBarrier(m *base.Module, l0 int64) {
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
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v49 int32
	_ = v49
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int64
	_ = v82
	var v94 int32
	_ = v94
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v13 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v13 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = l0
	F_errmsg_internal(m, int32(38087), v9+int32(32))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _consts[659]))
	v29 = v27 + int32(37)
	if int32(0) <= v29 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	F_errfinish(m, int32(494449), int32(431), int32(220847))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	v35 = v29
	goto L11
L9:
	;
	goto L10
L10:
	;
	v105 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L28
	}
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _consts[679]))
	v42 = v39 + v35<<(uint(int32(7))%32)
	v44 = v42 + int32(112)
	v45 = *(*int64)(unsafe.Add(mBase, uint32(v44)))
	*(*int64)(unsafe.Add(mBase, uint32(v44))) = v45
	if base.Ui64(v45) < base.Ui64(l0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L10
L13:
	;
	v49 = v42 + int32(8)
	goto L16
L14:
	;
	goto L15
L15:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L26
	}
L16:
	;
	v60 = F_ConditionVariableTimedSleep(m, v42+int32(124), int32(5000), int32(134217770))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	goto L15
L18:
	;
	v82 = *(*int64)(unsafe.Add(mBase, uint32(v49)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v49)+104)) = v82
	if base.Ui64(v82) < base.Ui64(l0) {
		goto L16
	} else {
		goto L25
	}
L19:
	;
	if v60 == int32(0) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v66 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	if v66 == int32(0) {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v70
	F_errmsg(m, int32(220872), v9+int32(16))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(494449), int32(452), int32(220847))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
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
	if int32(0) < v35 {
		v35 = v35 - int32(1)
		goto L11
	} else {
		goto L27
	}
L27:
	;
	goto L12
L28:
	;
	if v105 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9))) = l0
	F_errmsg_internal(m, int32(38078), v9)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	m.G0 = v9 + int32(48)
	return
L32:
	;
	F_errfinish(m, int32(494449), int32(461), int32(220847))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L31
}
