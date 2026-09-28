package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_FindAndDropRelationBuffers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int64
	_ = v63
	var v65 int64
	_ = v65
	var v80 int64
	_ = v80
	var v89 int64
	_ = v89
	var v109 int32
	_ = v109
	var v110 int64
	_ = v110
	var v113 int64
	_ = v113
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v152 int64
	_ = v152
	var v154 int64
	_ = v154
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v187 int64
	_ = v187
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	if base.Ui32(l3) < base.Ui32(l2) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v24 = v17
	v25 = l3
	goto L4
L2:
	;
	goto L3
L3:
	;
	m.G0 = v14 + int32(48)
	return
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v24
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v32
	v37 = v14 + int32(4)
	v38 = F_BufTableHashCode(m, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
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
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_FindAndDropRelationBuffers[0]))
	v48 = v41 + v38&int32(127)<<(uint(int32(7))%32) + int32(_a_F_FindAndDropRelationBuffers_0)
	v50 = F_LWLockAcquire(m, v48, int32(1))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v52 = F_BufTableLookup(m, v37, v38)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	F_LWLockRelease(m, v48)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	if v52 < int32(0) {
		v195 = v24
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v201 = v25 + int32(1)
	if v201 != l2 {
		v24 = v195
		v25 = v201
		goto L4
	} else {
		goto L44
	}
L12:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_FindAndDropRelationBuffers[1]))
	v62 = v59 + v52*int32(56)
	v63 = int64(4194304)
	v65 = base.AtomicRmwOr64(m, v62, int32(24), v63)
	if v65&v63 != int64(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v80 = v65
	goto L16
L14:
	;
	goto L15
L15:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v170 != v171 {
		goto L37
	} else {
		goto L38
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = int32(_a_F_FindAndDropRelationBuffers_1)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = int32(_a_F_FindAndDropRelationBuffers_2)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = int32(_a_F_FindAndDropRelationBuffers_3)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = int32(0)
	v89 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v89
	if v80&int64(4194304) != v89 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L15
L18:
	;
	goto L21
L19:
	;
	goto L20
L20:
	;
	v132 = int32(_a_F_FindAndDropRelationBuffers_4)
	v133 = *(*int32)(unsafe.Add(mBase, _c_F_FindAndDropRelationBuffers[2]))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(24))+8))
	if v135 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L21:
	;
	F_perform_spin_delay(m, v14+int32(24))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L6
	} else {
		goto L23
	}
L22:
	;
	goto L20
L23:
	;
	v110 = int64(0)
	v113 = base.AtomicRmwCmpxchg64(m, v62, int32(24), v110, v110)
	if v113&int64(4194304) != v110 {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v152 = int64(4194304)
	v154 = base.AtomicRmwOr64(m, v62, int32(24), v152)
	if v154&v152 != int64(0) {
		v80 = v154
		goto L16
	} else {
		goto L36
	}
L26:
	;
	goto L25
L27:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_FindAndDropRelationBuffers[2])) = v150
	goto L26
L28:
	;
	if int32(999) < v133 {
		goto L26
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	if v133 < int32(11) {
		goto L26
	} else {
		goto L35
	}
L31:
	;
	v140 = int32(900)
	if v140 <= v133 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v143 = v140
	goto L34
L33:
	;
	v143 = v133
	goto L34
L34:
	;
	v150 = v143 + int32(100)
	goto L27
L35:
	;
	v150 = v133 - int32(1)
	goto L27
L36:
	;
	goto L17
L37:
	;
	v187 = base.AtomicRmwSub64(m, v62, int32(24), int64(4194304))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v195 = v188
	goto L11
L38:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v173 != v174 {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v176 != v177 {
		goto L37
	} else {
		goto L40
	}
L40:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	if v179 != l1 {
		goto L37
	} else {
		goto L41
	}
L41:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
	if base.Ui32(v181) < base.Ui32(l3) {
		goto L37
	} else {
		goto L42
	}
L42:
	;
	F_InvalidateBuffer(m, v62)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L6
	} else {
		goto L43
	}
L43:
	;
	v195 = v170
	goto L11
L44:
	;
	goto L5
}
func F_check_and_set_sync_info(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_check_and_set_sync_info[0]))
	v7 = base.AtomicRmwXchg32(m, v4, int32(16), int32(1))
	if v7 != 0 {
		F_s_lock(m, v4+int32(16), int32(_a_F_check_and_set_sync_info_0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, _c_F_check_and_set_sync_info[0]))
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+4)))
			if v15 == int32(1) {
				v18 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v14)+16)), uint32(v18))
				v22 = *(*int32)(unsafe.Add(mBase, _c_F_check_and_set_sync_info[1]))
				if v22 == int32(7) {
					v27 = F_errstart(m, int32(14), int32(0))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						if v27 != 0 {
							F_errmsg(m, int32(_a_F_check_and_set_sync_info_1), int32(0))
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_check_and_set_sync_info_2), int32(1534), int32(_a_F_check_and_set_sync_info_3))
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return
								} else {
									F_proc_exit(m, int32(0))
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							F_proc_exit(m, int32(0))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
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
							F_errmsg(m, int32(_a_F_check_and_set_sync_info_4), int32(0))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_check_and_set_sync_info_2), int32(1546), int32(_a_F_check_and_set_sync_info_3))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
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
				v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+5)))
				if v57 == int32(1) {
					v60 = int32(0)
					atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v14)+16)), uint32(v60))
					F_errstart_cold(m, int32(21), v60)
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return
					} else {
						F_errcode(m, int32(325))
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return
						} else {
							F_errmsg(m, int32(_a_F_check_and_set_sync_info_5), int32(0))
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_check_and_set_sync_info_2), int32(1555), int32(_a_F_check_and_set_sync_info_3))
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v14))) = l0
					v80 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v14)+5)) = uint8(v80)
					v83 = *(*int32)(unsafe.Add(mBase, _c_F_check_and_set_sync_info[0]))
					v84 = int32(0)
					atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v83)+16)), uint32(v84))
					*(*uint8)(unsafe.Add(mBase, _c_F_check_and_set_sync_info[2])) = uint8(v80)
					return
				}
			}
		}
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_check_and_set_sync_info[0]))
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+4)))
		if v15 == int32(1) {
			v18 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v14)+16)), uint32(v18))
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_check_and_set_sync_info[1]))
			if v22 == int32(7) {
				v27 = F_errstart(m, int32(14), int32(0))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					if v27 != 0 {
						F_errmsg(m, int32(_a_F_check_and_set_sync_info_1), int32(0))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_check_and_set_sync_info_2), int32(1534), int32(_a_F_check_and_set_sync_info_3))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return
							} else {
								F_proc_exit(m, int32(0))
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						F_proc_exit(m, int32(0))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
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
						F_errmsg(m, int32(_a_F_check_and_set_sync_info_4), int32(0))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_check_and_set_sync_info_2), int32(1546), int32(_a_F_check_and_set_sync_info_3))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
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
			v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+5)))
			if v57 == int32(1) {
				v60 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v14)+16)), uint32(v60))
				F_errstart_cold(m, int32(21), v60)
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return
				} else {
					F_errcode(m, int32(325))
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_check_and_set_sync_info_5), int32(0))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_check_and_set_sync_info_2), int32(1555), int32(_a_F_check_and_set_sync_info_3))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = l0
				v80 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v14)+5)) = uint8(v80)
				v83 = *(*int32)(unsafe.Add(mBase, _c_F_check_and_set_sync_info[0]))
				v84 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v83)+16)), uint32(v84))
				*(*uint8)(unsafe.Add(mBase, _c_F_check_and_set_sync_info[2])) = uint8(v80)
				return
			}
		}
	}
}
func F_parse_and_validate_value(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
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
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v164 int32
	_ = v164
	var v165 float64
	_ = v165
	var v166 float64
	_ = v166
	var v170 float64
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v185 float64
	_ = v185
	var v186 int32
	_ = v186
	var v187 float64
	_ = v187
	var v188 float64
	_ = v188
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v276 int32
	_ = v276
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v317 int32
	_ = v317
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v394 int32
	_ = v394
	v7 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(224)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	switch v17 {
	case 0:
		goto L8
	case 1:
		goto L7
	case 2:
		goto L6
	case 3:
		goto L5
	case 4:
		goto L4
	default:
		goto L2
	}
L1:
	;
	m.G0 = v15 + int32(224)
	return v394
L2:
	;
	v394 = int32(1)
	goto L1
L3:
	;
	v394 = int32(0)
	goto L1
L4:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v241 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L5:
	;
	v222 = F_guc_strdup(m, l3, l1)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L13
	} else {
		goto L81
	}
L6:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v135 = F_parse_real(m, l1, l4, v132, v15+int32(220))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L13
	} else {
		goto L53
	}
L7:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v47 = F_parse_int(m, l1, l4, v44, v15+int32(220))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L13
	} else {
		goto L23
	}
L8:
	;
	v18 = F_strlen(m, l1)
	mBase = m.M
	v19 = F_parse_bool_with_len(m, l1, v18, l4)
	mBase = m.M
	goto L9
L9:
	;
	if v19 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v23 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v42 = F_call_bool_check_hook(m, l0, l4, l5, l2, l3)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L13
	} else {
		goto L19
	}
L13:
	;
	return int32(0)
L14:
	;
	if v23 == int32(0) {
		v394 = v7
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v32
	F_errmsg(m, int32(_a_F_parse_and_validate_value_0), v15)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(_a_F_parse_and_validate_value_1), int32(3044), int32(_a_F_parse_and_validate_value_2))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	v394 = v7
	goto L1
L19:
	;
	if v42 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	goto L3
L21:
	;
	v128 = F_call_int_check_hook(m, l0, l4, l5, l2, l3)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L13
	} else {
		goto L49
	}
L22:
	;
	F_errfinish(m, int32(_a_F_parse_and_validate_value_1), v118, int32(_a_F_parse_and_validate_value_2))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L13
	} else {
		goto L48
	}
L23:
	;
	if v47 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v52 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L13
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v78 <= v77 {
		goto L33
	} else {
		goto L34
	}
L27:
	;
	if v52 == int32(0) {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L13
	} else {
		goto L29
	}
L29:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+84)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v59
	F_errmsg(m, int32(_a_F_parse_and_validate_value_3), v15+int32(80))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L13
	} else {
		goto L30
	}
L30:
	;
	v67 = int32(3065)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v15)+220))
	if v68 == int32(0) {
		v118 = v67
		goto L22
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v68
	F_errhint(m, int32(_a_F_parse_and_validate_value_4), v15-int32(-64))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L13
	} else {
		goto L32
	}
L32:
	;
	v118 = v67
	goto L22
L33:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v77 <= v80 {
		goto L21
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v83 = F_get_config_unit_name(m, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L13
	} else {
		goto L37
	}
L36:
	;
	goto L35
L37:
	;
	v86 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L13
	} else {
		goto L38
	}
L38:
	;
	if v86 == int32(0) {
		goto L3
	} else {
		goto L39
	}
L39:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L13
	} else {
		goto L40
	}
L40:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v83 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v98 = v83
	goto L43
L42:
	;
	v98 = int32(_a_F_parse_and_validate_value_5)
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v98
	if v83 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v102 = int32(_a_F_parse_and_validate_value_6)
	goto L46
L45:
	;
	v102 = int32(_a_F_parse_and_validate_value_5)
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v95
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v93
	F_errmsg(m, int32(_a_F_parse_and_validate_value_7), v15+int32(16))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L13
	} else {
		goto L47
	}
L47:
	;
	v118 = int32(3085)
	goto L22
L48:
	;
	goto L3
L49:
	;
	if v128 == int32(0) {
		goto L3
	} else {
		goto L50
	}
L50:
	;
	goto L2
L51:
	;
	v220 = F_call_real_check_hook(m, l0, l4, l5, l2, l3)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L13
	} else {
		goto L79
	}
L52:
	;
	F_errfinish(m, int32(_a_F_parse_and_validate_value_1), v210, int32(_a_F_parse_and_validate_value_2))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L13
	} else {
		goto L78
	}
L53:
	;
	if v135 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v140 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L13
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v165 = *(*float64)(unsafe.Add(mBase, uint32(l4)))
	v166 = *(*float64)(unsafe.Add(mBase, uint32(l0)+112))
	if base.F64_lt(v165, v166) == int32(0) {
		goto L63
	} else {
		goto L64
	}
L57:
	;
	if v140 == int32(0) {
		goto L3
	} else {
		goto L58
	}
L58:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L13
	} else {
		goto L59
	}
L59:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+180)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = v147
	F_errmsg(m, int32(_a_F_parse_and_validate_value_3), v15+int32(176))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L13
	} else {
		goto L60
	}
L60:
	;
	v155 = int32(3106)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v15)+220))
	if v156 == int32(0) {
		v210 = v155
		goto L52
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v156
	F_errhint(m, int32(_a_F_parse_and_validate_value_4), v15+int32(160))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L13
	} else {
		goto L62
	}
L62:
	;
	v210 = v155
	goto L52
L63:
	;
	v170 = *(*float64)(unsafe.Add(mBase, uint32(l0)+120))
	if base.F64_gt(v165, v170) == int32(0) {
		goto L51
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v175 = F_get_config_unit_name(m, v174)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L13
	} else {
		goto L67
	}
L66:
	;
	goto L65
L67:
	;
	v178 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L13
	} else {
		goto L68
	}
L68:
	;
	if v178 == int32(0) {
		goto L3
	} else {
		goto L69
	}
L69:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L13
	} else {
		goto L70
	}
L70:
	;
	v185 = *(*float64)(unsafe.Add(mBase, uint32(l4)))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v187 = *(*float64)(unsafe.Add(mBase, uint32(l0)+112))
	v188 = *(*float64)(unsafe.Add(mBase, uint32(l0)+120))
	if v175 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v190 = v175
	goto L73
L72:
	;
	v190 = int32(_a_F_parse_and_validate_value_5)
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+148)) = v190
	if v175 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v194 = int32(_a_F_parse_and_validate_value_6)
	goto L76
L75:
	;
	v194 = int32(_a_F_parse_and_validate_value_5)
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+144)) = v194
	*(*float64)(unsafe.Add(mBase, uint32(v15)+136)) = v188
	*(*int32)(unsafe.Add(mBase, uint32(v15)+132)) = v190
	*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = v194
	*(*float64)(unsafe.Add(mBase, uint32(v15)+120)) = v187
	*(*int32)(unsafe.Add(mBase, uint32(v15)+112)) = v186
	*(*int32)(unsafe.Add(mBase, uint32(v15)+108)) = v190
	*(*int32)(unsafe.Add(mBase, uint32(v15)+104)) = v194
	*(*float64)(unsafe.Add(mBase, uint32(v15)+96)) = v185
	F_errmsg(m, int32(_a_F_parse_and_validate_value_8), v15+int32(96))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L13
	} else {
		goto L77
	}
L77:
	;
	v210 = int32(3126)
	goto L52
L78:
	;
	goto L3
L79:
	;
	if v220 != 0 {
		goto L2
	} else {
		goto L80
	}
L80:
	;
	goto L3
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v222
	if v222 == int32(0) {
		goto L3
	} else {
		goto L82
	}
L82:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
	if v227&int32(8) != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v230 = F_strlen(m, v222)
	mBase = m.M
	F_truncate_identifier(m, v222, v230, int32(1))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L13
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v234 = F_call_string_check_hook(m, l0, l4, l5, l2, l3)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L13
	} else {
		goto L87
	}
L86:
	;
	goto L85
L87:
	;
	if v234 != 0 {
		goto L2
	} else {
		goto L88
	}
L88:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v236 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	F_pfree(m, v236)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L13
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
	v394 = v7
	goto L1
L92:
	;
	goto L91
L93:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v254)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v357
	v359 = F_call_enum_check_hook(m, l0, l4, l5, l2, l3)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L13
	} else {
		goto L128
	}
L94:
	;
	v317 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v317
	v325 = F_config_enum_get_options(m, l0+int32(96), int32(_a_F_parse_and_validate_value_9), int32(_a_F_parse_and_validate_value_10), int32(_a_F_parse_and_validate_value_11))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L13
	} else {
		goto L114
	}
L95:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v241)))
	if v244 == int32(0) {
		goto L94
	} else {
		goto L96
	}
L96:
	;
	v254 = v241
	v255 = v244
	goto L97
L97:
	;
	v261 = l1
	v262 = v255
	goto L100
L98:
	;
	goto L94
L99:
	;
	if v299 == int32(0) {
		goto L93
	} else {
		goto L112
	}
L100:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261))))
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262))))
	if v265 == v266 {
		v288 = v265
		goto L102
	} else {
		goto L103
	}
L101:
	;
	v299 = int32(0)
	goto L99
L102:
	;
	v290 = int32(1)
	if v288 != 0 {
		v261 = v261 + v290
		v262 = v262 + v290
		goto L100
	} else {
		goto L111
	}
L103:
	;
	if base.Ui32((v265-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v276 = v265 | int32(32)
	goto L106
L105:
	;
	v276 = v265
	goto L106
L106:
	;
	if base.Ui32((v266-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v285 = v266 | int32(32)
	goto L109
L108:
	;
	v285 = v266
	goto L109
L109:
	;
	if v276 == v285 {
		v288 = v276
		goto L102
	} else {
		goto L110
	}
L110:
	;
	v299 = v276 - v285
	goto L99
L111:
	;
	goto L101
L112:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v254)+12))
	if v302 != 0 {
		v254 = v254 + int32(12)
		v255 = v302
		goto L97
	} else {
		goto L113
	}
L113:
	;
	goto L98
L114:
	;
	v328 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L13
	} else {
		goto L115
	}
L115:
	;
	if v328 != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L13
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	if v325 == int32(0) {
		v394 = v317
		goto L1
	} else {
		goto L126
	}
L119:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+212)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v15)+208)) = v333
	F_errmsg(m, int32(_a_F_parse_and_validate_value_3), v15+int32(208))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L13
	} else {
		goto L120
	}
L120:
	;
	if v325 != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+192)) = v325
	F_errhint(m, int32(_a_F_parse_and_validate_value_4), v15+int32(192))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L13
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	F_errfinish(m, int32(_a_F_parse_and_validate_value_1), int32(3190), int32(_a_F_parse_and_validate_value_2))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L13
	} else {
		goto L125
	}
L124:
	;
	goto L123
L125:
	;
	goto L118
L126:
	;
	F_pfree(m, v325)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L13
	} else {
		goto L127
	}
L127:
	;
	v394 = v317
	goto L1
L128:
	;
	if v359 != 0 {
		goto L2
	} else {
		goto L129
	}
L129:
	;
	goto L3
}
func F_update_and_persist_local_synced_slot(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	v4 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_update_and_persist_local_synced_slot[0]))
	v13 = F_update_local_synced_slot(m, l0, l1)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)+288))
		if v17 != 0 {
			if l2 == int32(0) {
				v69 = v4
			} else {
				v20 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v20)
				v69 = v4
			}
			m.G0 = v9 + int32(32)
			return v69
		} else {
			v22 = F_IsLogicalDecodingEnabled(m)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				if v22 == int32(0) {
					v28 = F_errstart(m, int32(15), int32(0))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						if v28 != 0 {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v30
							F_errmsg(m, int32(_a_F_update_and_persist_local_synced_slot_0), v9+int32(16))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								v39 = F_errdetail(m, int32(_a_F_update_and_persist_local_synced_slot_1), int32(0))
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_update_and_persist_local_synced_slot_2), int32(739), int32(_a_F_update_and_persist_local_synced_slot_3))
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return int32(0)
									} else {
										if l2 == int32(0) {
											v69 = v4
										} else {
											v48 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v48)
											v69 = v4
										}
										m.G0 = v9 + int32(32)
										return v69
									}
								}
							}
						} else {
							if l2 == int32(0) {
								v69 = v4
							} else {
								v48 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v48)
								v69 = v4
							}
							m.G0 = v9 + int32(32)
							return v69
						}
					}
				} else {
					F_ReplicationSlotPersist(m)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						v52 = int32(1)
						v55 = F_errstart(m, int32(15), int32(0))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							if v55 == int32(0) {
								v69 = v52
								m.G0 = v9 + int32(32)
								return v69
							} else {
								v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v59
								F_errmsg(m, int32(_a_F_update_and_persist_local_synced_slot_4), v9)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_update_and_persist_local_synced_slot_2), int32(751), int32(_a_F_update_and_persist_local_synced_slot_3))
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return int32(0)
									} else {
										v69 = v52
										m.G0 = v9 + int32(32)
										return v69
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
