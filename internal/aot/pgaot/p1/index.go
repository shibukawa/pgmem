package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecIndexEvalRuntimeKeys(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = int32(_a_F_ExecIndexEvalRuntimeKeys_0)
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_ExecIndexEvalRuntimeKeys[0]))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecIndexEvalRuntimeKeys[0])) = v17
	if v4 < l2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v27 = v4
	goto L4
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecIndexEvalRuntimeKeys[0])) = v15
	m.G0 = v12 + int32(16)
	return
L4:
	;
	v32 = l1 + v27*int32(12)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
	v38 = m.T0[v37].(func(*base.Module, int32, int32, int32) int32)(m, v34, l0, v12+int32(15))
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
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
	if v40 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v57
	v60 = v27 + int32(1)
	if v60 != l2 {
		v27 = v60
		goto L4
	} else {
		goto L16
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+44)) = v38
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v57 = v44 | int32(1)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+8)))
	if v47 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v50 = F_pg_detoast_datum(m, v38)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L6
	} else {
		goto L15
	}
L13:
	;
	v52 = v38
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+44)) = v52
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v57 = v54 & int32(-2)
	goto L8
L15:
	;
	v52 = v50
	goto L14
L16:
	;
	goto L5
}
func F_IndexAmTranslateCompareType(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l1 != int32(403) {
		v19 = F_GetIndexAmRoutineByAmId(m, l1, int32(0))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+136))
			if v23 != 0 {
				v24 = m.T0[v23].(func(*base.Module, int32, int32) int32)(m, l0, l2)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = v24
					if l3 != 0 {
						v41 = v26
						m.G0 = v10 + int32(16)
						return v41 & int32(_a_F_IndexAmTranslateCompareType_0)
					} else {
						if v26 != 0 {
							v41 = v26
							m.G0 = v10 + int32(16)
							return v41 & int32(_a_F_IndexAmTranslateCompareType_0)
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l1
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
								F_errmsg_internal(m, int32(_a_F_IndexAmTranslateCompareType_1), v10)
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_IndexAmTranslateCompareType_2), int32(165), int32(_a_F_IndexAmTranslateCompareType_3))
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
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
			} else {
				v26 = int32(0)
				if l3 != 0 {
					v41 = v26
					m.G0 = v10 + int32(16)
					return v41 & int32(_a_F_IndexAmTranslateCompareType_0)
				} else {
					if v26 != 0 {
						v41 = v26
						m.G0 = v10 + int32(16)
						return v41 & int32(_a_F_IndexAmTranslateCompareType_0)
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
							F_errmsg_internal(m, int32(_a_F_IndexAmTranslateCompareType_1), v10)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_IndexAmTranslateCompareType_2), int32(165), int32(_a_F_IndexAmTranslateCompareType_3))
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
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
	} else {
		if base.Ui32(int32(4)) < base.Ui32(l0-int32(1)) {
			v19 = F_GetIndexAmRoutineByAmId(m, l1, int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+136))
				if v23 != 0 {
					v24 = m.T0[v23].(func(*base.Module, int32, int32) int32)(m, l0, l2)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = v24
						if l3 != 0 {
							v41 = v26
							m.G0 = v10 + int32(16)
							return v41 & int32(_a_F_IndexAmTranslateCompareType_0)
						} else {
							if v26 != 0 {
								v41 = v26
								m.G0 = v10 + int32(16)
								return v41 & int32(_a_F_IndexAmTranslateCompareType_0)
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v30 = m.ExcPending
								if v30 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l1
									*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
									F_errmsg_internal(m, int32(_a_F_IndexAmTranslateCompareType_1), v10)
									mBase = m.M
									v35 = m.ExcPending
									if v35 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_IndexAmTranslateCompareType_2), int32(165), int32(_a_F_IndexAmTranslateCompareType_3))
										mBase = m.M
										v40 = m.ExcPending
										if v40 != 0 {
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
				} else {
					v26 = int32(0)
					if l3 != 0 {
						v41 = v26
						m.G0 = v10 + int32(16)
						return v41 & int32(_a_F_IndexAmTranslateCompareType_0)
					} else {
						if v26 != 0 {
							v41 = v26
							m.G0 = v10 + int32(16)
							return v41 & int32(_a_F_IndexAmTranslateCompareType_0)
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l1
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
								F_errmsg_internal(m, int32(_a_F_IndexAmTranslateCompareType_1), v10)
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_IndexAmTranslateCompareType_2), int32(165), int32(_a_F_IndexAmTranslateCompareType_3))
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
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
		} else {
			v41 = l0
			m.G0 = v10 + int32(16)
			return v41 & int32(_a_F_IndexAmTranslateCompareType_0)
		}
	}
}
func F_IndexOnlyNext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
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
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 float64
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
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
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 float64
	_ = v239
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+100))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v23 = v20 * v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v26 != 0 {
		v53 = v26
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v54 = F_index_getnext_tid(m, v53, v23)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L3
	} else {
		goto L14
	}
L2:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v34 = F_index_beginscan(m, v27, v28, v29, l0+int32(160), v32, v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v34
	v39 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+28)) = uint8(v39)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = int32(0)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v43 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+144)))
	if v44 != int32(1) {
		v53 = v34
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	F_index_rescan(m, v34, v47, v48, v49, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L3
	} else {
		goto L9
	}
L8:
	;
	goto L7
L9:
	;
	v53 = v34
	goto L1
L10:
	;
	m.G0 = v17 + int32(16)
	return v24
L11:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
	if v305 <= int32(0) {
		goto L74
	} else {
		goto L75
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L3
	} else {
		goto L71
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L3
	} else {
		goto L68
	}
L14:
	;
	if v54 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v62 = v54
	goto L18
L16:
	;
	goto L17
L17:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v273)+12))
	m.T0[v274].(func(*base.Module, int32))(m, v24)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L3
	} else {
		goto L67
	}
L18:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_IndexOnlyNext[0]))
	if v73 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L17
L20:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L3
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v77 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62)+2)))
	v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62))))
	v82 = F_visibilitymap_get_status(m, v76, v77|v78<<(uint(int32(16))%32), l0+int32(176))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L3
	} else {
		goto L25
	}
L23:
	;
	goto L22
L24:
	;
	v257 = F_index_getnext_tid(m, v53, v23)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L3
	} else {
		goto L65
	}
L25:
	;
	v85 = v82 & int32(1)
	if v85 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v88 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v53)+52))
	if v107 != 0 {
		goto L37
	} else {
		goto L38
	}
L29:
	;
	v89 = *(*float64)(unsafe.Add(mBase, uint32(v88)+224))
	*(*float64)(unsafe.Add(mBase, uint32(v88)+224)) = base.F64_add(v89, float64(1))
	goto L31
L30:
	;
	goto L31
L31:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	v94 = F_index_fetch_heap(m, v53, v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L3
	} else {
		goto L32
	}
L32:
	;
	if v94 == int32(0) {
		goto L24
	} else {
		goto L33
	}
L33:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	m.T0[v100].(func(*base.Module, int32))(m, v98)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L3
	} else {
		goto L34
	}
L34:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+66)))
	if v103 == int32(1) {
		goto L13
	} else {
		goto L35
	}
L35:
	;
	goto L28
L36:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+72)))
	if v211 != int32(1) {
		goto L11
	} else {
		goto L56
	}
L37:
	;
	F_ExecForceStoreHeapTuple(m, v107, v24, int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L3
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v53)+44))
	if v111 == int32(0) {
		goto L12
	} else {
		goto L41
	}
L40:
	;
	goto L36
L41:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v53)+48))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	m.T0[v116].(func(*base.Module, int32))(m, v24)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L3
	} else {
		goto L42
	}
L42:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	F_index_deform_tuple(m, v111, v114, v119, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L3
	} else {
		goto L43
	}
L43:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	if v123 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v190 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
	v192 = v190 & int32(_a_F_IndexOnlyNext_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)) = uint16(v192)
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)))
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+6)) = uint16(v195)
	goto L55
L45:
	;
	v126 = int32(0)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	if v127 <= v126 {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v131 = v126
	goto L47
L47:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v148 = int32(*(*int16)(unsafe.Add(mBase, uint32(v144+v131<<(uint(int32(1))%32)))))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148+v149))))
	if v151 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L44
L49:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)+20))
	v157 = F_MemoryContextAlloc(m, v155, int32(64))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L3
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v174 = v131 + int32(1)
	if v174 != v127 {
		v131 = v174
		goto L47
	} else {
		goto L54
	}
L52:
	;
	v160 = v148 << (uint(int32(2)) % 32)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v160+v161)))
	v165 = F_strncpy(m, v157, v163, int32(64))
	mBase = m.M
	v166 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v165)+63)) = uint8(v166)
	goto L53
L53:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v168+v160))) = v157
	goto L51
L54:
	;
	goto L48
L55:
	;
	goto L36
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v24
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v215 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	F_MemoryContextReset(m, v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L3
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v221 = int32(_a_F_IndexOnlyNext_1)
	v222 = *(*int32)(unsafe.Add(mBase, _c_F_IndexOnlyNext[1]))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_IndexOnlyNext[1])) = v224
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v215)+20))
	v229 = m.T0[v228].(func(*base.Module, int32, int32, int32) int32)(m, v215, v25, v17+int32(15))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L3
	} else {
		goto L61
	}
L60:
	;
	goto L11
L61:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_IndexOnlyNext[1])) = v222
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	F_MemoryContextReset(m, v233)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L3
	} else {
		goto L62
	}
L62:
	;
	if v229 != 0 {
		goto L11
	} else {
		goto L63
	}
L63:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v236 == int32(0) {
		goto L24
	} else {
		goto L64
	}
L64:
	;
	v239 = *(*float64)(unsafe.Add(mBase, uint32(v236)+248))
	*(*float64)(unsafe.Add(mBase, uint32(v236)+248)) = base.F64_add(v239, float64(1))
	goto L24
L65:
	;
	if v257 != 0 {
		v62 = v257
		goto L18
	} else {
		goto L66
	}
L66:
	;
	goto L19
L67:
	;
	goto L10
L68:
	;
	F_errmsg_internal(m, int32(_a_F_IndexOnlyNext_2), int32(0))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L3
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_IndexOnlyNext_3), int32(181), int32(_a_F_IndexOnlyNext_4))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L3
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	F_errmsg_internal(m, int32(_a_F_IndexOnlyNext_5), int32(0))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L3
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_IndexOnlyNext_3), int32(213), int32(_a_F_IndexOnlyNext_4))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L3
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
	if v85 == int32(0) {
		goto L10
	} else {
		goto L81
	}
L75:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+84)))
	if v308 != int32(1) {
		goto L74
	} else {
		goto L76
	}
L76:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L3
	} else {
		goto L77
	}
L77:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L3
	} else {
		goto L78
	}
L78:
	;
	F_errmsg(m, int32(_a_F_IndexOnlyNext_6), int32(0))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L3
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(_a_F_IndexOnlyNext_3), int32(240), int32(_a_F_IndexOnlyNext_4))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L3
	} else {
		goto L80
	}
L80:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L81:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v330 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62)+2)))
	v331 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62))))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	F_PredicateLockPage(m, v329, v330|v331<<(uint(int32(16))%32), v335)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L3
	} else {
		goto L82
	}
L82:
	;
	goto L10
}
func F_build_index_pathkeys(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
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
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v198 int32
	_ = v198
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v222 int32
	_ = v222
	v4 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	if v12 == v4 {
		v222 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v222
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+92))
	if v15 == int32(0) {
		v222 = v4
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v18 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L6
L6:
	;
	v29 = v4
	v31 = v4
	goto L7
L7:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v36 <= v29 {
		v222 = v31
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v222 = v207
	goto L1
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+v29))))
	v42 = v29 << (uint(int32(2)) % 32)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v42+v43)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if base.B2i32(l2 != int32(-1)) == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v61+v42)))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v64+v42)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v67+v42)))
	v70 = int32(1)
	v74 = int32(0)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
	v78 = F_make_pathkey_from_sortinfo(m, l0, v46, v63, v66, v69, v59&v70, v60&v70, v74, v76, v74)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L15
	} else {
		goto L16
	}
L11:
	;
	v49 = int32(1)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51+v29))))
	v59 = v40 ^ v49
	v60 = v53 ^ v49
	goto L10
L12:
	;
	goto L13
L13:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+v29))))
	v59 = v40
	v60 = v58
	goto L10
L14:
	;
	v213 = v29 + int32(1)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v213 < v214 {
		v29 = v213
		v31 = v207
		goto L7
	} else {
		goto L50
	}
L15:
	;
	return int32(0)
L16:
	;
	if v78 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+40)))
	if v83 != 0 {
		v207 = v31
		goto L14
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v124+v29<<(uint(int32(2))%32))))
	if base.Ui32(v128) <= base.Ui32(int32(_a_F_build_index_pathkeys_0)) {
		goto L31
	} else {
		goto L32
	}
L20:
	;
	if v31 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v122 = F_lappend(m, v31, v78)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L15
	} else {
		goto L28
	}
L22:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v86 <= int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v93 = int32(0)
	goto L24
L24:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v89+v93<<(uint(int32(2))%32))))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
	if v82 == v106 {
		v207 = v31
		goto L14
	} else {
		goto L26
	}
L25:
	;
	goto L21
L26:
	;
	v109 = v93 + int32(1)
	if v86 != v109 {
		v93 = v109
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v207 = v122
	goto L14
L29:
	;
	if v198 == int32(0) {
		v222 = v31
		goto L1
	} else {
		goto L49
	}
L30:
	;
	v140 = int32(0)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+184))
	if v142 == v140 {
		v198 = v140
		goto L29
	} else {
		goto L38
	}
L31:
	;
	if v128 == int32(424) {
		goto L30
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v137 = F_op_in_opfamily(m, int32(91), v128)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L15
	} else {
		goto L36
	}
L34:
	;
	if v128 == int32(2222) {
		goto L30
	} else {
		goto L35
	}
L35:
	;
	v198 = int32(0)
	goto L29
L36:
	;
	if v137 != 0 {
		goto L30
	} else {
		goto L37
	}
L37:
	;
	v198 = int32(0)
	goto L29
L38:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	if int32(0) < v145 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v151 = int32(0)
	goto L42
L40:
	;
	goto L41
L41:
	;
	v198 = int32(0)
	goto L29
L42:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v160+v151<<(uint(int32(2))%32))))
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+10)))
	if v165 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L41
L44:
	;
	v172 = v151 + int32(1)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	if v172 < v173 {
		v151 = v172
		goto L42
	} else {
		goto L48
	}
L45:
	;
	v166 = F_match_boolean_index_clause(m, l0, v164, v29, l1)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L15
	} else {
		goto L46
	}
L46:
	;
	if v166 == int32(0) {
		goto L44
	} else {
		goto L47
	}
L47:
	;
	v198 = int32(1)
	goto L29
L48:
	;
	goto L43
L49:
	;
	v207 = v31
	goto L14
L50:
	;
	goto L8
}
func F_get_index_constraint(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
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
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(144)
	m.G0 = v8
	v12 = F_table_open(m, int32(2608), int32(1))
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
	F_ScanKeyInit(m, v8, int32(1), int32(3), int32(184), int32(1259))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_ScanKeyInit(m, v8+int32(48), int32(2), int32(3), int32(184), l0)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v31 = int32(3)
	F_ScanKeyInit(m, v8+int32(96), v31, v31, int32(65), int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v41 = F_systable_beginscan(m, v12, int32(2673), int32(1), int32(0), int32(3), v8)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	F_systable_endscan(m, v41)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L18
	}
L7:
	;
	v43 = F_systable_getnext(m, v41)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	if v43 == int32(0) {
		v68 = v2
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v47 = v43
	goto L10
L10:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+22)))
	v54 = v52 + v53
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+12))
	if v55 != int32(2606) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v68 = v2
	goto L6
L12:
	;
	v63 = F_systable_getnext(m, v41)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L16
	}
L13:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
	if v58 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+24)))
	if v59 != int32(105) {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	v68 = v62
	goto L6
L16:
	;
	if v63 != 0 {
		v47 = v63
		goto L10
	} else {
		goto L17
	}
L17:
	;
	goto L11
L18:
	;
	F_sequence_close(m, v12, int32(1))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	m.G0 = v8 + int32(144)
	return v68
}
func F_get_index_isclustered(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = F_SearchSysCache1(m, int32(34), l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg_internal(m, int32(_a_F_get_index_isclustered_0), v6)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_get_index_isclustered_1), int32(3749), int32(_a_F_get_index_isclustered_2))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+v29)+17)))
			F_ReleaseCatCache(m, v9)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return v31
			}
		}
	}
}
func F_index_am_handler_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(_a_F_index_am_handler_out_0)
			F_errmsg(m, int32(_a_F_index_am_handler_out_1), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_index_am_handler_out_2), int32(371), int32(_a_F_index_am_handler_out_3))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
func F_index_beginscan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
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
	v7 = int32(0)
	v9 = F_index_beginscan_internal(m, l1, l4, l5, l2, v7, v7)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
		v18 = m.T0[v17].(func(*base.Module, int32) int32)(m, l0)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+68)) = v18
			return v9
		}
	}
}
func F_index_checkable(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+119)))
	if v11 != int32(105) {
		v84 = F_SearchSysCache1(m, int32(2), l1)
		mBase = m.M
		v85 = m.ExcPending
		if v85 != 0 {
			return int32(0)
		} else {
			v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+84))
			v89 = F_SearchSysCache1(m, int32(2), v88)
			mBase = m.M
			v90 = m.ExcPending
			if v90 != 0 {
				return int32(0)
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v94 = m.ExcPending
				if v94 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(1088))
					mBase = m.M
					v97 = m.ExcPending
					if v97 != 0 {
						return int32(0)
					} else {
						v98 = *(*int32)(unsafe.Add(mBase, uint32(v84)+16))
						v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+22)))
						*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v98 + v99 + int32(4)
						F_errmsg(m, int32(_a_F_index_checkable_0), v8-int32(-64))
						mBase = m.M
						v110 = m.ExcPending
						if v110 != 0 {
							return int32(0)
						} else {
							v111 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
							v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+22)))
							v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							v114 = int32(4)
							*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v113 + v114
							*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = v111 + v112 + v114
							F_errdetail(m, int32(_a_F_index_checkable_1), v8+int32(48))
							mBase = m.M
							v126 = m.ExcPending
							if v126 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_index_checkable_2), int32(173), int32(_a_F_index_checkable_3))
								mBase = m.M
								v133 = m.ExcPending
								if v133 != 0 {
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
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+84))
		if v14 != l1 {
			v84 = F_SearchSysCache1(m, int32(2), l1)
			mBase = m.M
			v85 = m.ExcPending
			if v85 != 0 {
				return int32(0)
			} else {
				v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+84))
				v89 = F_SearchSysCache1(m, int32(2), v88)
				mBase = m.M
				v90 = m.ExcPending
				if v90 != 0 {
					return int32(0)
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(1088))
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return int32(0)
						} else {
							v98 = *(*int32)(unsafe.Add(mBase, uint32(v84)+16))
							v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+22)))
							*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v98 + v99 + int32(4)
							F_errmsg(m, int32(_a_F_index_checkable_0), v8-int32(-64))
							mBase = m.M
							v110 = m.ExcPending
							if v110 != 0 {
								return int32(0)
							} else {
								v111 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
								v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+22)))
								v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
								v114 = int32(4)
								*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v113 + v114
								*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = v111 + v112 + v114
								F_errdetail(m, int32(_a_F_index_checkable_1), v8+int32(48))
								mBase = m.M
								v126 = m.ExcPending
								if v126 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_index_checkable_2), int32(173), int32(_a_F_index_checkable_3))
									mBase = m.M
									v133 = m.ExcPending
									if v133 != 0 {
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
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+118)))
			if v16 == int32(116) {
				v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
				if v19 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v137 = m.ExcPending
					if v137 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(1088))
						mBase = m.M
						v140 = m.ExcPending
						if v140 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_index_checkable_4), int32(0))
							mBase = m.M
							v146 = m.ExcPending
							if v146 != 0 {
								return int32(0)
							} else {
								v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
								*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v147 + int32(4)
								F_errdetail(m, int32(_a_F_index_checkable_5), v8+int32(16))
								mBase = m.M
								v156 = m.ExcPending
								if v156 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_index_checkable_2), int32(181), int32(_a_F_index_checkable_3))
									mBase = m.M
									v163 = m.ExcPending
									if v163 != 0 {
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
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
					v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+18)))
					if v24 != 0 {
						v78 = int32(1)
						m.G0 = v8 + int32(80)
						return v78
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v168 = m.ExcPending
						if v168 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(1088))
							mBase = m.M
							v171 = m.ExcPending
							if v171 != 0 {
								return int32(0)
							} else {
								v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v172 + int32(4)
								F_errmsg(m, int32(_a_F_index_checkable_6), v8)
								mBase = m.M
								v180 = m.ExcPending
								if v180 != 0 {
									return int32(0)
								} else {
									F_errdetail(m, int32(_a_F_index_checkable_7), int32(0))
									mBase = m.M
									v185 = m.ExcPending
									if v185 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_index_checkable_2), int32(188), int32(_a_F_index_checkable_3))
										mBase = m.M
										v192 = m.ExcPending
										if v192 != 0 {
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
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+18)))
				if v26 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v168 = m.ExcPending
					if v168 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(1088))
						mBase = m.M
						v171 = m.ExcPending
						if v171 != 0 {
							return int32(0)
						} else {
							v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v172 + int32(4)
							F_errmsg(m, int32(_a_F_index_checkable_6), v8)
							mBase = m.M
							v180 = m.ExcPending
							if v180 != 0 {
								return int32(0)
							} else {
								F_errdetail(m, int32(_a_F_index_checkable_7), int32(0))
								mBase = m.M
								v185 = m.ExcPending
								if v185 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_index_checkable_2), int32(188), int32(_a_F_index_checkable_3))
									mBase = m.M
									v192 = m.ExcPending
									if v192 != 0 {
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
					v29 = int32(1)
					if v16 != int32(117) {
						v78 = v29
						m.G0 = v8 + int32(80)
						return v78
					} else {
						v34 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_index_checkable[0])))
						if v34 == int32(1) {
							v39 = *(*int32)(unsafe.Add(mBase, _c_F_index_checkable[1]))
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+316))
							v42 = base.B2i32(v40 != int32(2))
							*(*uint8)(unsafe.Add(mBase, _c_F_index_checkable[0])) = uint8(v42)
							v44 = v42
						} else {
							v44 = int32(0)
						}
						if v44 == int32(0) {
							v78 = v29
							m.G0 = v8 + int32(80)
							return v78
						} else {
							v47 = int32(0)
							v50 = F_errstart(m, int32(18), v47)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								if v50 == int32(0) {
									v78 = v47
									m.G0 = v8 + int32(80)
									return v78
								} else {
									F_errcode(m, int32(100663618))
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return int32(0)
									} else {
										v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
										*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v59 + int32(4)
										F_errmsg(m, int32(_a_F_index_checkable_8), v8+int32(32))
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_index_checkable_2), int32(45), int32(_a_F_index_checkable_9))
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return int32(0)
											} else {
												v78 = v47
												m.G0 = v8 + int32(80)
												return v78
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
func F_index_deform_tuple_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v44 int32
	_ = v44
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	v7 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v7 < v19 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v25 = int32(0)
	v32 = v7
	v33 = v7
	goto L4
L2:
	;
	goto L3
L3:
	;
	m.G0 = v17 + int32(16)
	return
L4:
	;
	if l5 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v188 = v25 + int32(1)
	if v188 != v19 {
		v25 = v188
		v32 = v186
		v33 = v183
		goto L4
	} else {
		goto L56
	}
L7:
	;
	v61 = l0 + int32(20) + v25<<(uint(int32(4))%32)
	v63 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v25+l2))) = uint8(v63)
	if v32&int32(1) == v63 {
		goto L14
	} else {
		goto L15
	}
L8:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4+int32(base.Ui32(v25)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v44)>>(uint(v25&int32(7))%32))&int32(1) != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1+v25<<(uint(int32(2))%32)))) = int32(0)
	v56 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v25+l2))) = uint8(v56)
	v183 = v33
	v186 = v56
	goto L6
L10:
	;
	v112 = l3 + v110
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+6)))
	if v116 == int32(1) {
		goto L25
	} else {
		goto L26
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v79
	v110 = v79
	v111 = v69
	goto L10
L12:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+12)))
	v102 = int32(1)
	v110 = (v33 + v100 - v102) & (int32(0) - v100)
	v111 = v102
	goto L10
L13:
	;
	v90 = int32(1)
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v33))))
	if v92 != 0 {
		v110 = v33
		v111 = v90
		goto L10
	} else {
		goto L23
	}
L14:
	;
	v69 = int32(0)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	if v69 <= v70 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61)+4)))
	if v85 != int32(_a_F_index_deform_tuple_internal_0) {
		goto L12
	} else {
		goto L22
	}
L17:
	;
	v110 = v70
	v111 = v69
	goto L10
L18:
	;
	goto L19
L19:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+12)))
	v79 = (v33 + v73 - int32(1)) & (int32(0) - v73)
	v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61)+4)))
	if v80 != int32(_a_F_index_deform_tuple_internal_0) {
		goto L11
	} else {
		goto L20
	}
L20:
	;
	if v79 != v33 {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v33
	v110 = v33
	v111 = v69
	goto L10
L22:
	;
	goto L13
L23:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+12)))
	v110 = (v33 + v93 - int32(1)) & (int32(0) - v93)
	v111 = v90
	goto L10
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1+v25<<(uint(int32(2))%32)))) = v140
	v142 = int32(*(*int16)(unsafe.Add(mBase, uint32(v61)+4)))
	v143 = int32(0)
	v144 = base.B2i32(v142 <= v143)
	if v144 == v143 {
		v178 = v142
		goto L36
	} else {
		goto L37
	}
L25:
	;
	v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61)+4)))
	switch v119 - int32(1) {
	case 0:
		goto L31
	case 1:
		goto L30
	default:
		goto L28
	case 3:
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v140 = v112
	goto L24
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	v140 = v124
	goto L24
L30:
	;
	v123 = int32(*(*int16)(unsafe.Add(mBase, uint32(v112))))
	v140 = v123
	goto L24
L31:
	;
	v122 = int32(*(*int8)(unsafe.Add(mBase, uint32(v112))))
	v140 = v122
	goto L24
L32:
	;
	return
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = base.I32_extend16_s(v119)
	F_errmsg_internal(m, int32(_a_F_index_deform_tuple_internal_1), v17)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_index_deform_tuple_internal_2), int32(70), int32(_a_F_index_deform_tuple_internal_3))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L32
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
	v183 = v178 + v110
	v186 = v111 | v144
	goto L6
L37:
	;
	if v142 == int32(-1) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	if v149 == int32(1) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	v175 = F_strlen(m, v112)
	mBase = m.M
	v178 = v175 + int32(1)
	goto L36
L41:
	;
	v152 = int32(6)
	v154 = int32(18)
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+1)))
	if v156 == v154 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	if v149&int32(1) != 0 {
		goto L53
	} else {
		goto L54
	}
L44:
	;
	v159 = v154
	goto L46
L45:
	;
	v159 = int32(2)
	goto L46
L46:
	;
	if v156&int32(254) == int32(2) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v164 = v152
	goto L49
L48:
	;
	v164 = v159
	goto L49
L49:
	;
	if v156 == int32(1) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v167 = v152
	goto L52
L51:
	;
	v167 = v164
	goto L52
L52:
	;
	v178 = v167
	goto L36
L53:
	;
	v178 = int32(base.Ui32(v149) >> (uint(int32(1)) % 32))
	goto L36
L54:
	;
	goto L55
L55:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	v178 = int32(base.Ui32(v172) >> (uint(int32(2)) % 32))
	goto L36
L56:
	;
	goto L5
}
func F_index_endscan(m *base.Module, l0 int32) {
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+204))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+108))
	if v11 != 0 {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
		if v12 != 0 {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+188))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
			m.T0[v15].(func(*base.Module, int32))(m, v12)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = int32(0)
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+204))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+108))
				v23 = v22
				m.T0[v23].(func(*base.Module, int32))(m, l0)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					F_RelationDecrementReferenceCount(m, v26)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
						if v29 == int32(1) {
							v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							F_UnregisterSnapshot(m, v32)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return
							} else {
								v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								if v35 != 0 {
									F_pfree(m, v35)
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return
									} else {
										v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										if v38 != 0 {
											F_pfree(m, v38)
											mBase = m.M
											v40 = m.ExcPending
											if v40 != 0 {
												return
											} else {
												F_pfree(m, l0)
												mBase = m.M
												v42 = m.ExcPending
												if v42 != 0 {
													return
												} else {
													m.G0 = v7 + int32(16)
													return
												}
											}
										} else {
											F_pfree(m, l0)
											mBase = m.M
											v42 = m.ExcPending
											if v42 != 0 {
												return
											} else {
												m.G0 = v7 + int32(16)
												return
											}
										}
									}
								} else {
									v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									if v38 != 0 {
										F_pfree(m, v38)
										mBase = m.M
										v40 = m.ExcPending
										if v40 != 0 {
											return
										} else {
											F_pfree(m, l0)
											mBase = m.M
											v42 = m.ExcPending
											if v42 != 0 {
												return
											} else {
												m.G0 = v7 + int32(16)
												return
											}
										}
									} else {
										F_pfree(m, l0)
										mBase = m.M
										v42 = m.ExcPending
										if v42 != 0 {
											return
										} else {
											m.G0 = v7 + int32(16)
											return
										}
									}
								}
							}
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							if v35 != 0 {
								F_pfree(m, v35)
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return
								} else {
									v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									if v38 != 0 {
										F_pfree(m, v38)
										mBase = m.M
										v40 = m.ExcPending
										if v40 != 0 {
											return
										} else {
											F_pfree(m, l0)
											mBase = m.M
											v42 = m.ExcPending
											if v42 != 0 {
												return
											} else {
												m.G0 = v7 + int32(16)
												return
											}
										}
									} else {
										F_pfree(m, l0)
										mBase = m.M
										v42 = m.ExcPending
										if v42 != 0 {
											return
										} else {
											m.G0 = v7 + int32(16)
											return
										}
									}
								}
							} else {
								v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								if v38 != 0 {
									F_pfree(m, v38)
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return
									} else {
										F_pfree(m, l0)
										mBase = m.M
										v42 = m.ExcPending
										if v42 != 0 {
											return
										} else {
											m.G0 = v7 + int32(16)
											return
										}
									}
								} else {
									F_pfree(m, l0)
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
										return
									} else {
										m.G0 = v7 + int32(16)
										return
									}
								}
							}
						}
					}
				}
			}
		} else {
			v23 = v11
			m.T0[v23].(func(*base.Module, int32))(m, l0)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				F_RelationDecrementReferenceCount(m, v26)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
					if v29 == int32(1) {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						F_UnregisterSnapshot(m, v32)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							if v35 != 0 {
								F_pfree(m, v35)
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return
								} else {
									v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									if v38 != 0 {
										F_pfree(m, v38)
										mBase = m.M
										v40 = m.ExcPending
										if v40 != 0 {
											return
										} else {
											F_pfree(m, l0)
											mBase = m.M
											v42 = m.ExcPending
											if v42 != 0 {
												return
											} else {
												m.G0 = v7 + int32(16)
												return
											}
										}
									} else {
										F_pfree(m, l0)
										mBase = m.M
										v42 = m.ExcPending
										if v42 != 0 {
											return
										} else {
											m.G0 = v7 + int32(16)
											return
										}
									}
								}
							} else {
								v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								if v38 != 0 {
									F_pfree(m, v38)
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return
									} else {
										F_pfree(m, l0)
										mBase = m.M
										v42 = m.ExcPending
										if v42 != 0 {
											return
										} else {
											m.G0 = v7 + int32(16)
											return
										}
									}
								} else {
									F_pfree(m, l0)
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
										return
									} else {
										m.G0 = v7 + int32(16)
										return
									}
								}
							}
						}
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v35 != 0 {
							F_pfree(m, v35)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return
							} else {
								v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								if v38 != 0 {
									F_pfree(m, v38)
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return
									} else {
										F_pfree(m, l0)
										mBase = m.M
										v42 = m.ExcPending
										if v42 != 0 {
											return
										} else {
											m.G0 = v7 + int32(16)
											return
										}
									}
								} else {
									F_pfree(m, l0)
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
										return
									} else {
										m.G0 = v7 + int32(16)
										return
									}
								}
							}
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							if v38 != 0 {
								F_pfree(m, v38)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return
								} else {
									F_pfree(m, l0)
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
										return
									} else {
										m.G0 = v7 + int32(16)
										return
									}
								}
							} else {
								F_pfree(m, l0)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return
								} else {
									m.G0 = v7 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v49 = m.ExcPending
		if v49 != 0 {
			return
		} else {
			v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+48))
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_index_endscan_0)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v51 + int32(4)
			F_errmsg_internal(m, int32(_a_F_index_endscan_1), v7)
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_index_endscan_2), int32(385), int32(_a_F_index_endscan_3))
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
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
func F_index_getprocinfo(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
	v14 = int32(1)
	v15 = l1 - v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+6)))
	v21 = l2 + v15*v17 - v14
	v24 = v13 + v21*int32(28)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v25 != 0 {
		m.G0 = v11 + int32(16)
		return v24
	} else {
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v26+v21<<(uint(int32(2))%32))))
		if v30 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return int32(0)
			} else {
				v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v69 + int32(4)
				F_errmsg_internal(m, int32(_a_F_index_getprocinfo_0), v11)
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_index_getprocinfo_1), int32(947), int32(_a_F_index_getprocinfo_2))
					mBase = m.M
					v82 = m.ExcPending
					if v82 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+8)))
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
			F_fmgr_info_cxt(m, v30, v24, v34)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				if l2 == v33&int32(_a_F_index_getprocinfo_3) {
					m.G0 = v11 + int32(16)
					return v24
				} else {
					v43 = F_RelationGetIndexAttOptions(m, l0, int32(0))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						v45 = int32(_a_F_index_getprocinfo_4)
						v46 = *(*int32)(unsafe.Add(mBase, _c_F_index_getprocinfo[0]))
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
						*(*int32)(unsafe.Add(mBase, _c_F_index_getprocinfo[0])) = v48
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v43+v15<<(uint(int32(2))%32))))
						F_set_fn_opclass_options(m, v24, v53)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_index_getprocinfo[0])) = v46
							m.G0 = v11 + int32(16)
							return v24
						}
					}
				}
			}
		}
	}
}
func F_index_opclass_options(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v338 int32
	_ = v338
	var v351 int32
	_ = v351
	v5 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+8)))
	if v18 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v15 + int32(16)
	return v351
L2:
	;
	v79 = v15 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v79))) = int64(0)
	goto L16
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+6)))
	v24 = int32(2)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v19+v20*(l1-int32(1))<<(uint(v24)%32)+v18<<(uint(v24)%32)-int32(4))))
	if v32 != 0 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v33 = int32(0)
	if l2 == v33 {
		v351 = v33
		goto L1
	} else {
		goto L7
	}
L6:
	;
	goto L5
L7:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	v39 = F_SysCacheGetAttrNotNull(m, int32(34), v37, int32(18))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v39+l1<<(uint(int32(2))%32))+20))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v54 = m.G0
	v56 = v54 - int32(16)
	m.G0 = v56
	F_initStringInfo(m, v56)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	F_get_opclass_name(m, v46, int32(0), v56)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	m.G0 = v56 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v63 + int32(1)
	F_errmsg(m, int32(_a_F_index_opclass_options_0), v15)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(_a_F_index_opclass_options_1), int32(1076), int32(_a_F_index_opclass_options_2))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L16:
	;
	v84 = F_index_getprocinfo(m, l0, l1, v18)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	v89 = F_FunctionCall1Coll(m, v84, int32(0), v15+int32(4))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	v92 = v15 + int32(4)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	if v93 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v95 = v94
	goto L21
L20:
	;
	v95 = v5
	goto L21
L21:
	;
	v98 = F_palloc(m, v95<<(uint(int32(4))%32))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L8
	} else {
		goto L22
	}
L22:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	if v100 == int32(0) {
		v165 = v5
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v171 = F_palloc(m, v165<<(uint(int32(4))%32))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L8
	} else {
		goto L32
	}
L24:
	;
	v103 = int32(0)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	if v103 < v104 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v107 = v103
	goto L28
L26:
	;
	goto L27
L27:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	if v153 == int32(0) {
		v165 = v5
		goto L23
	} else {
		goto L31
	}
L28:
	;
	v121 = v98 + v107<<(uint(int32(4))%32)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v122+v107<<(uint(int32(2))%32))))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	*(*int32)(unsafe.Add(mBase, uint32(v121))) = v128
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v121)+4)) = v131
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v121)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v121)+8)) = v133
	v138 = v107 + int32(1)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	if v138 < v139 {
		v107 = v138
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
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	v165 = v156
	goto L23
L32:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	if v173 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	if l2 != 0 {
		goto L39
	} else {
		goto L40
	}
L34:
	;
	v176 = int32(0)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
	if v177 <= v176 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v180 = v176
	goto L36
L36:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v173)+12))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v192+v180<<(uint(int32(2))%32))))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
	v200 = v171 + v180<<(uint(int32(4))%32)
	v201 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v200)+4)) = uint8(v201)
	*(*int32)(unsafe.Add(mBase, uint32(v200))) = v197
	v205 = v180 + int32(1)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
	if v205 < v206 {
		v180 = v205
		goto L36
	} else {
		goto L38
	}
L37:
	;
	goto L33
L38:
	;
	goto L37
L39:
	;
	F_parseRelOptionsInternal(m, l2, l3, v171, v165)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L8
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v92)+8))
	if int32(0) < v95 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L41
L43:
	;
	v226 = int32(0)
	v236 = v222
	goto L46
L44:
	;
	v286 = v222
	goto L45
L45:
	;
	v288 = F_palloc0(m, v286)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L8
	} else {
		goto L68
	}
L46:
	;
	v238 = int32(4)
	v240 = v171 + v226<<(uint(v238)%32)
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v240)))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v241)+20))
	if v242 == v238 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v286 = v272
	goto L45
L48:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240)+4)))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v241)+36))
	if v246 != 0 {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v272 = v236
	goto L50
L50:
	;
	v274 = v226 + int32(1)
	if v274 != v95 {
		v226 = v274
		v236 = v272
		goto L46
	} else {
		goto L67
	}
L51:
	;
	v272 = v268 + v236
	goto L50
L52:
	;
	if v245&int32(1) != 0 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	if v245&int32(1) != 0 {
		goto L64
	} else {
		goto L65
	}
L55:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v240)+8))
	v251 = m.T0[v246].(func(*base.Module, int32, int32) int32)(m, v249, int32(0))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L8
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241)+28)))
	if v253 != 0 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v268 = v251
	goto L51
L59:
	;
	v256 = int32(0)
	goto L61
L60:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v241)+40))
	v256 = v255
	goto L61
L61:
	;
	v258 = m.T0[v246].(func(*base.Module, int32, int32) int32)(m, v256, int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L8
	} else {
		goto L62
	}
L62:
	;
	v268 = v258
	goto L51
L63:
	;
	v268 = v265 + int32(1)
	goto L51
L64:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v240)+8))
	v263 = F_strlen(m, v262)
	mBase = m.M
	v265 = v263
	goto L63
L65:
	;
	goto L66
L66:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v241)+24))
	v265 = v264
	goto L63
L67:
	;
	goto L47
L68:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v92)+8))
	F_fillRelOptions(m, v288, v290, v171, v95, l3, v98, v95)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L8
	} else {
		goto L69
	}
L69:
	;
	if l3 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	if v98 != 0 {
		goto L78
	} else {
		goto L79
	}
L71:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	if v295 == int32(0) {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v295)+4))
	if v298 <= int32(0) {
		goto L70
	} else {
		goto L73
	}
L73:
	;
	v302 = int32(0)
	goto L74
L74:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v295)+12))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v314+v302<<(uint(int32(2))%32))))
	m.T0[v318].(func(*base.Module, int32, int32, int32))(m, v288, v171, v95)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L8
	} else {
		goto L76
	}
L75:
	;
	goto L70
L76:
	;
	v322 = v302 + int32(1)
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v295)+4))
	if v322 < v323 {
		v302 = v322
		goto L74
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	F_pfree(m, v98)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L8
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v351 = v288
	goto L1
L81:
	;
	goto L80
}
func F_index_open(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = F_relation_open(m, l0, l1)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+119)))
		if v13|int32(32) != int32(105) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(151027844))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = v25 + int32(4)
					F_errmsg(m, int32(_a_F_index_open_0), v6)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_index_open_1), int32(204), int32(_a_F_index_open_2))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
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
			m.G0 = v6 + int32(16)
			return v8
		}
	}
}
func F_index_other_operands_eval_cost(m *base.Module, l0 int32, l1 int32) float64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 float64
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v25 float64
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
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
	var v68 int32
	_ = v68
	var v69 float64
	_ = v69
	var v70 float64
	_ = v70
	var v72 float64
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v84 float64
	_ = v84
	v3 = int32(0)
	v8 = float64(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	if l1 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(32)
	return v84
L2:
	;
	v84 = v8
	goto L1
L3:
	;
	goto L4
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v15 <= int32(0) {
		v84 = v8
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v23 = v3
	v25 = v8
	goto L6
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27+v23<<(uint(int32(2))%32))))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v32 == int32(318) {
		goto L13
	} else {
		goto L14
	}
L7:
	;
	v84 = v72
	goto L1
L8:
	;
	F_cost_qual_eval_node(m, v11+int32(16), v64, l0)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L16
	} else {
		goto L20
	}
L9:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v37)+28))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	v64 = v63
	goto L8
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L16
	} else {
		goto L17
	}
L11:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v37)+28))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v64 = v44
	goto L8
L12:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	v64 = v41
	goto L8
L13:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v37 = v35
	v38 = v36
	goto L15
L14:
	;
	v37 = v31
	v38 = v32
	goto L15
L15:
	;
	switch v38 - int32(17) {
	case 0:
		goto L9
	default:
		goto L10
	case 3:
		goto L11
	case 20:
		goto L12
	case 35:
		v64 = int32(0)
		goto L8
	}
L16:
	;
	return float64(0)
L17:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v51
	F_errmsg_internal(m, int32(_a_F_index_other_operands_eval_cost_0), v11)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(_a_F_index_other_operands_eval_cost_1), int32(_a_F_index_other_operands_eval_cost_2), int32(_a_F_index_other_operands_eval_cost_3))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L20:
	;
	v69 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
	v70 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
	v72 = base.F64_add(v25, base.F64_add(v69, v70))
	v74 = v23 + int32(1)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v74 < v75 {
		v23 = v74
		v25 = v72
		goto L6
	} else {
		goto L21
	}
L21:
	;
	goto L7
}
func F_index_pages_fetched(m *base.Module, l0 float64, l1 int32, l2 float64, l3 int32) float64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v13 float64
	_ = v13
	var v14 float64
	_ = v14
	var v16 int32
	_ = v16
	var v19 float64
	_ = v19
	var v20 float64
	_ = v20
	var v24 float64
	_ = v24
	var v25 float64
	_ = v25
	var v29 float64
	_ = v29
	var v33 float64
	_ = v33
	var v39 float64
	_ = v39
	var v49 float64
	_ = v49
	var v51 float64
	_ = v51
	v8 = int32(1)
	if base.Ui32(l1) <= base.Ui32(v8) {
		v11 = v8
	} else {
		v11 = l1
	}
	v12 = base.F64_convert_i32_u(v11)
	v13 = base.F64_add(v12, v12)
	v14 = float64(1)
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_index_pages_fetched[0]))
	v19 = *(*float64)(unsafe.Add(mBase, uint32(l3)+288))
	v20 = base.F64_add(l2, v19)
	if base.F64_gt(v20, v14) != 0 {
		v24 = v20
	} else {
		v24 = v14
	}
	v25 = base.F64_div(base.F64_mul(v12, base.F64_convert_i32_s(v16)), v24)
	if base.F64_le(v25, float64(1)) != 0 {
		v29 = v14
	} else {
		v29 = base.F64_ceil(v25)
	}
	if base.F64_le(v12, v29) != 0 {
		v33 = base.F64_div(base.F64_mul(l0, v13), base.F64_add(v13, l0))
		if base.F64_ge(v33, v12) != 0 {
			v51 = v12
			return v51
		} else {
			return base.F64_ceil(v33)
		}
	} else {
		v39 = base.F64_div(base.F64_mul(v13, v29), base.F64_sub(v13, v29))
		if base.F64_ge(v39, l0) != 0 {
			v49 = base.F64_div(base.F64_mul(l0, v13), base.F64_add(v13, l0))
		} else {
			v49 = base.F64_add(v29, base.F64_div(base.F64_mul(base.F64_sub(v12, v29), base.F64_sub(l0, v39)), v12))
		}
		v51 = base.F64_ceil(v49)
		return v51
	}
}
func F_index_restrpos(m *base.Module, l0 int32) {
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+204))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+116))
	if v11 != 0 {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
		if v12 != 0 {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+188))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
			m.T0[v15].(func(*base.Module, int32))(m, v12)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v19 = v18
				v20 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+66)) = uint8(v20)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)) = uint8(v20)
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v19)+204))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+116))
				m.T0[v25].(func(*base.Module, int32))(m, l0)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					m.G0 = v7 + int32(16)
					return
				}
			}
		} else {
			v19 = v9
			v20 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+66)) = uint8(v20)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)) = uint8(v20)
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v19)+204))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+116))
			m.T0[v25].(func(*base.Module, int32))(m, l0)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				m.G0 = v7 + int32(16)
				return
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+48))
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_index_restrpos_0)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v36 + int32(4)
			F_errmsg_internal(m, int32(_a_F_index_restrpos_1), v7)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_index_restrpos_2), int32(441), int32(_a_F_index_restrpos_3))
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
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
func F_index_strategy_get_limit(m *base.Module, l0 int32) float64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 float64
	_ = v34
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	switch l0 - int32(1) {
	case 0:
		v33 = int32(_a_F_index_strategy_get_limit_0)
		v34 = *(*float64)(unsafe.Add(mBase, uint32(v33)))
		m.G0 = v7 + int32(16)
		return v34
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return float64(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
			F_errmsg_internal(m, int32(_a_F_index_strategy_get_limit_1), v7)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return float64(0)
			} else {
				F_errfinish(m, int32(_a_F_index_strategy_get_limit_2), int32(267), int32(_a_F_index_strategy_get_limit_3))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return float64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 6:
		v33 = int32(_a_F_index_strategy_get_limit_4)
		v34 = *(*float64)(unsafe.Add(mBase, uint32(v33)))
		m.G0 = v7 + int32(16)
		return v34
	case 8:
		v33 = int32(_a_F_index_strategy_get_limit_5)
		v34 = *(*float64)(unsafe.Add(mBase, uint32(v33)))
		m.G0 = v7 + int32(16)
		return v34
	}
}
func F_index_vacuum_cleanup(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
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
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+56))
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_index_vacuum_cleanup[0]))
	if v15 != v11 {
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_index_vacuum_cleanup[1]))
		v19 = F_list_member_ptr(m, v18, v11)
		mBase = m.M
		v20 = v19
	} else {
		v20 = int32(1)
	}
	if v20 == int32(0) {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v10)+204))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+56))
		if v24 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return int32(0)
			} else {
				v58 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = int32(_a_F_index_vacuum_cleanup_0)
				*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v58 + int32(4)
				F_errmsg_internal(m, int32(_a_F_index_vacuum_cleanup_1), v8+int32(16))
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_index_vacuum_cleanup_2), int32(822), int32(_a_F_index_vacuum_cleanup_3))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v27 = m.T0[v24].(func(*base.Module, int32, int32) int32)(m, l0, l1)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				m.G0 = v8 + int32(32)
				return v27
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v42 + int32(4)
				F_errmsg(m, int32(_a_F_index_vacuum_cleanup_4), v8)
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_index_vacuum_cleanup_2), int32(821), int32(_a_F_index_vacuum_cleanup_3))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
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
func F_plan_create_index_workers(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
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
	var v96 int32
	_ = v96
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
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 float64
	_ = v131
	var v132 float64
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v13 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plan_create_index_workers[0])))
	if v13 != int32(1) {
		v272 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v10 + int32(48)
	return v272
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_plan_create_index_workers[1]))
	if v17 == int32(0) {
		v272 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = F_palloc0(m, int32(168))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21))) = int64(4294967363)
	v28 = F_palloc0(m, int32(92))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = int32(266)
	v33 = F_palloc0(m, int32(384))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+12)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = int32(267)
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_plan_create_index_workers[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+344)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+280)) = v42
	v47 = F_palloc0(m, int32(8))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = int32(272)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v47
	v56 = F_list_make1_impl(m, int32(1), v10+int32(12))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v56
	v60 = F_palloc0(m, int32(136))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v62 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v60)+24)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v60)+16)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v60)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = int32(101)
	v69 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+124)) = uint16(v69)
	v71 = int32(_a_F_plan_create_index_workers_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+20)) = uint16(v71)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v60
	v78 = F_list_make1_impl(m, v62, v10+int32(8))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+52)) = v78
	v83 = F_addRTEPermissionInfo(m, v21+int32(56), v60)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	F_setup_simple_rel_arrays(m, v33)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v89 = F_build_simple_rel(m, v33, int32(1), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v92 = F_table_open(m, l0, int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v95 = F_index_open(m, l1, int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v97 = int32(0)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v92)+48))
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+118)))
	if v99 == int32(116) {
		v259 = v97
		goto L17
	} else {
		goto L18
	}
L17:
	;
	F_relation_close(m, v95, int32(0))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L4
	} else {
		goto L81
	}
L18:
	;
	v102 = F_RelationGetIndexExpressions(m, v95)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v104 = F_is_parallel_safe(m, v33, v102)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	if v104 == int32(0) {
		v259 = v97
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v108 = F_RelationGetIndexPredicate(m, v95)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	v110 = F_is_parallel_safe(m, v33, v108)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	if v110 == int32(0) {
		v259 = v97
		goto L17
	} else {
		goto L24
	}
L24:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v89)+148))
	if v114 != int32(-1) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_plan_create_index_workers[1]))
	if v114 < v118 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	F_estimate_rel_size(m, v92, int32(0), v10+int32(44), v10+int32(32), v10+int32(24))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L4
	} else {
		goto L31
	}
L28:
	;
	v120 = v114
	goto L30
L29:
	;
	v120 = v118
	goto L30
L30:
	;
	v259 = v120
	goto L17
L31:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
	v131 = base.F64_convert_i32_u(v130)
	v132 = float64(-1)
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_plan_create_index_workers[1]))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v89)+148))
	if v137 != int32(-1) {
		v226 = v137
		goto L34
	} else {
		goto L35
	}
L32:
	;
	if v236 <= int32(0) {
		goto L72
	} else {
		goto L73
	}
L33:
	;
	goto L32
L34:
	;
	if v226 < v134 {
		goto L69
	} else {
		goto L70
	}
L35:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	if v140 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v157 = int32(0)
	if base.F64_ge(v131, float64(0)) == v157 {
		v189 = v157
		goto L44
	} else {
		goto L45
	}
L37:
	;
	if base.F64_ge(v131, float64(0)) != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v145 = *(*int32)(unsafe.Add(mBase, _c_F_plan_create_index_workers[3]))
	if base.F64_lt(v131, base.F64_convert_i32_s(v145)) != 0 {
		v236 = int32(0)
		goto L33
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	if base.F64_ge(v132, float64(0)) == int32(0) {
		goto L36
	} else {
		goto L42
	}
L41:
	;
	goto L40
L42:
	;
	v154 = *(*int32)(unsafe.Add(mBase, _c_F_plan_create_index_workers[4]))
	if base.F64_lt(v132, base.F64_convert_i32_s(v154)) != 0 {
		v236 = int32(0)
		goto L33
	} else {
		goto L43
	}
L43:
	;
	goto L36
L44:
	;
	if base.F64_ge(v132, float64(0)) == int32(0) {
		v226 = v189
		goto L34
	} else {
		goto L53
	}
L45:
	;
	v162 = int32(1)
	v164 = *(*int32)(unsafe.Add(mBase, _c_F_plan_create_index_workers[3]))
	if v164 <= v162 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v167 = v162
	goto L48
L47:
	;
	v167 = v164
	goto L48
L48:
	;
	v169 = v167
	v173 = int32(1)
	goto L49
L49:
	;
	v176 = v169 * int32(3)
	if base.F64_ge(v131, base.F64_convert_i32_u(v176)) == int32(0) {
		v189 = v173
		goto L44
	} else {
		goto L51
	}
L50:
	;
	v189 = v182
	goto L44
L51:
	;
	v182 = v173 + int32(1)
	if v176 < int32(715827883) {
		v169 = v176
		v173 = v182
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v195 = int32(1)
	v197 = *(*int32)(unsafe.Add(mBase, _c_F_plan_create_index_workers[4]))
	if v197 <= v195 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v200 = v195
	goto L56
L55:
	;
	v200 = v197
	goto L56
L56:
	;
	v202 = v200
	v207 = int32(1)
	goto L57
L57:
	;
	v209 = v202 * int32(3)
	if base.F64_le(base.F64_convert_i32_u(v209), v132) != 0 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	if v189 < v216 {
		goto L63
	} else {
		goto L64
	}
L59:
	;
	v213 = v207 + int32(1)
	if v209 < int32(715827883) {
		v202 = v209
		v207 = v213
		goto L57
	} else {
		goto L62
	}
L60:
	;
	v216 = v207
	goto L61
L61:
	;
	goto L58
L62:
	;
	v216 = v213
	goto L61
L63:
	;
	v218 = v189
	goto L65
L64:
	;
	v218 = v216
	goto L65
L65:
	;
	if int32(0) < v189 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v221 = v218
	goto L68
L67:
	;
	v221 = v216
	goto L68
L68:
	;
	v226 = v221
	goto L34
L69:
	;
	v229 = v226
	goto L71
L70:
	;
	v229 = v134
	goto L71
L71:
	;
	v236 = v229
	goto L33
L72:
	;
	v259 = v236
	goto L17
L73:
	;
	goto L74
L74:
	;
	v240 = *(*int32)(unsafe.Add(mBase, _c_F_plan_create_index_workers[5]))
	v244 = v236
	goto L75
L75:
	;
	v250 = base.I32_div_s(v240, v244+int32(1))
	if int32(_a_F_plan_create_index_workers_1) < v250 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v259 = v97
	goto L17
L77:
	;
	v259 = v244
	goto L17
L78:
	;
	goto L79
L79:
	;
	v253 = int32(1)
	if v253 < v244 {
		v244 = v244 - v253
		goto L75
	} else {
		goto L80
	}
L80:
	;
	goto L76
L81:
	;
	F_sequence_close(m, v92, int32(0))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	v272 = v259
	goto L1
}
func F_validate_index(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
	var v23 int64
	_ = v23
	var v25 int64
	_ = v25
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v98 int64
	_ = v98
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v112 int64
	_ = v112
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v126 int64
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v140 int64
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v178 int64
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
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
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 float32
	_ = v260
	var v261 int32
	_ = v261
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int64
	_ = v274
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v293 int64
	_ = v293
	var v296 int64
	_ = v296
	var v299 int64
	_ = v299
	var v302 int64
	_ = v302
	var v309 int32
	_ = v309
	var v316 int32
	_ = v316
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v447 int64
	_ = v447
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v510 int32
	_ = v510
	var v516 int32
	_ = v516
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 float64
	_ = v535
	var v537 float64
	_ = v537
	var v539 float64
	_ = v539
	var v543 int32
	_ = v543
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(160)
	m.G0 = v12
	v15 = v12 + int32(144)
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_validate_index[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v17
	v20 = *(*int64)(unsafe.Add(mBase, _c_F_validate_index[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+136)) = v20
	v23 = *(*int64)(unsafe.Add(mBase, _c_F_validate_index[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+128)) = v23
	v25 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+112)) = v25
	*(*int64)(unsafe.Add(mBase, uint32(v12)+104)) = v25
	*(*int64)(unsafe.Add(mBase, uint32(v12)+96)) = v25
	*(*int64)(unsafe.Add(mBase, uint32(v12)+88)) = v25
	*(*int64)(unsafe.Add(mBase, uint32(v12)+80)) = int64(4)
	v37 = v12 + int32(128)
	v39 = v12 + int32(80)
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_validate_index[3]))
	if v47 == v4 {
	} else {
		v53 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_validate_index[4])))
		if v53&int32(1) == int32(0) {
		} else {
			v58 = int32(_a_F_validate_index_0)
			v60 = *(*int32)(unsafe.Add(mBase, _c_F_validate_index[5]))
			v61 = int32(1)
			*(*int32)(unsafe.Add(mBase, _c_F_validate_index[5])) = v60 + v61
			v64 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
			*(*int32)(unsafe.Add(mBase, uint32(v47))) = v64 + v61
			v73 = v47 + int32(232)
			v79 = int32(0)
			v82 = v4
			for {
				v88 = int32(2)
				v91 = *(*int32)(unsafe.Add(mBase, uint32(v37+v82<<(uint(v88)%32))))
				v92 = int32(3)
				v98 = *(*int64)(unsafe.Add(mBase, uint32(v39+v82<<(uint(v92)%32))))
				*(*int64)(unsafe.Add(mBase, uint32(v73+v91<<(uint(v92)%32)))) = v98
				v101 = v82 | int32(1)
				v105 = *(*int32)(unsafe.Add(mBase, uint32(v37+v101<<(uint(v88)%32))))
				v112 = *(*int64)(unsafe.Add(mBase, uint32(v39+v101<<(uint(v92)%32))))
				*(*int64)(unsafe.Add(mBase, uint32(v73+v105<<(uint(v92)%32)))) = v112
				v115 = v82 | v88
				v119 = *(*int32)(unsafe.Add(mBase, uint32(v37+v115<<(uint(v88)%32))))
				v126 = *(*int64)(unsafe.Add(mBase, uint32(v39+v115<<(uint(v92)%32))))
				*(*int64)(unsafe.Add(mBase, uint32(v73+v119<<(uint(v92)%32)))) = v126
				v129 = v82 | v92
				v133 = *(*int32)(unsafe.Add(mBase, uint32(v37+v129<<(uint(v88)%32))))
				v140 = *(*int64)(unsafe.Add(mBase, uint32(v39+v129<<(uint(v92)%32))))
				*(*int64)(unsafe.Add(mBase, uint32(v73+v133<<(uint(v92)%32)))) = v140
				v142 = int32(4)
				v143 = v82 + v142
				v145 = v79 + v142
				if v145 != int32(4) {
					v79 = v145
					v82 = v143
					continue
				} else {
					break
				}
				break
			}
			v159 = int32(0)
			v162 = v143
			for {
				v171 = *(*int32)(unsafe.Add(mBase, uint32(v37+v162<<(uint(int32(2))%32))))
				v172 = int32(3)
				v178 = *(*int64)(unsafe.Add(mBase, uint32(v39+v162<<(uint(v172)%32))))
				*(*int64)(unsafe.Add(mBase, uint32(v73+v171<<(uint(v172)%32)))) = v178
				v180 = int32(1)
				v183 = v159 + v180
				if v183 != int32(1) {
					v159 = v183
					v162 = v162 + v180
					continue
				} else {
					break
				}
				break
			}
			v194 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
			v195 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v47))) = v194 + v195
			v198 = int32(_a_F_validate_index_0)
			v200 = *(*int32)(unsafe.Add(mBase, _c_F_validate_index[5]))
			*(*int32)(unsafe.Add(mBase, _c_F_validate_index[5])) = v200 - v195
		}
	}
	v214 = F_table_open(m, l0, int32(4))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		return
	} else {
		v221 = *(*int32)(unsafe.Add(mBase, _c_F_validate_index[6]))
		*(*int32)(unsafe.Add(mBase, uint32(v12+int32(124)))) = v221
		v224 = *(*int32)(unsafe.Add(mBase, _c_F_validate_index[7]))
		*(*int32)(unsafe.Add(mBase, uint32(v12+int32(120)))) = v224
		v226 = *(*int32)(unsafe.Add(mBase, uint32(v214)+48))
		v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+80))
		v228 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
		*(*int32)(unsafe.Add(mBase, _c_F_validate_index[7])) = v228 | int32(2)
		*(*int32)(unsafe.Add(mBase, _c_F_validate_index[6])) = v227
		v236 = int32(_a_F_validate_index_1)
		v238 = *(*int32)(unsafe.Add(mBase, _c_F_validate_index[8]))
		v240 = v238 + int32(1)
		*(*int32)(unsafe.Add(mBase, _c_F_validate_index[8])) = v240
		F_RestrictSearchPath(m)
		mBase = m.M
		v243 = m.ExcPending
		if v243 != 0 {
			return
		} else {
			v245 = F_index_open(m, l1, int32(3))
			mBase = m.M
			v246 = m.ExcPending
			if v246 != 0 {
				return
			} else {
				v247 = F_BuildIndexInfo(m, v245)
				mBase = m.M
				v248 = m.ExcPending
				if v248 != 0 {
					return
				} else {
					v249 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v247)+121)) = uint8(v249)
					*(*int32)(unsafe.Add(mBase, uint32(v12)+92)) = int32(13)
					*(*uint8)(unsafe.Add(mBase, uint32(v12)+90)) = uint8(v249)
					v255 = int32(256)
					*(*uint16)(unsafe.Add(mBase, uint32(v12)+88)) = uint16(v255)
					*(*int32)(unsafe.Add(mBase, uint32(v12)+84)) = v214
					*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v245
					v259 = *(*int32)(unsafe.Add(mBase, uint32(v214)+48))
					v260 = *(*float32)(unsafe.Add(mBase, uint32(v259)+100))
					v261 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v12)+104)) = v261
					*(*float64)(unsafe.Add(mBase, uint32(v12)+96)) = base.F64_promote_f32(v260)
					v270 = *(*int32)(unsafe.Add(mBase, _c_F_validate_index[9]))
					v272 = F_tuplesort_begin_datum(m, int32(20), int32(412), v261, v261, v270, v261)
					mBase = m.M
					v273 = m.ExcPending
					if v273 != 0 {
						return
					} else {
						v274 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v15))) = v274
						*(*int64)(unsafe.Add(mBase, uint32(v12)+152)) = v274
						*(*int64)(unsafe.Add(mBase, uint32(v12)+136)) = v274
						*(*int32)(unsafe.Add(mBase, uint32(v12)+128)) = v272
						v287 = F_index_bulk_delete(m, v12+int32(80), int32(0), int32(468), v12+int32(128))
						mBase = m.M
						v288 = m.ExcPending
						if v288 != 0 {
							return
						} else {
							v290 = *(*int32)(unsafe.Add(mBase, _c_F_validate_index[10]))
							*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = v290
							v293 = *(*int64)(unsafe.Add(mBase, _c_F_validate_index[11]))
							*(*int64)(unsafe.Add(mBase, uint32(v12)+64)) = v293
							v296 = *(*int64)(unsafe.Add(mBase, _c_F_validate_index[12]))
							*(*int64)(unsafe.Add(mBase, uint32(v12)+48)) = v296
							v299 = *(*int64)(unsafe.Add(mBase, _c_F_validate_index[13]))
							*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = v299
							v302 = *(*int64)(unsafe.Add(mBase, _c_F_validate_index[14]))
							*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v302
							v309 = int32(0)
							v316 = *(*int32)(unsafe.Add(mBase, _c_F_validate_index[3]))
							if v316 == v309 {
							} else {
								v322 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_validate_index[4])))
								if v322&int32(1) == int32(0) {
								} else {
									v327 = int32(_a_F_validate_index_0)
									v329 = *(*int32)(unsafe.Add(mBase, _c_F_validate_index[5]))
									v330 = int32(1)
									*(*int32)(unsafe.Add(mBase, _c_F_validate_index[5])) = v329 + v330
									v333 = *(*int32)(unsafe.Add(mBase, uint32(v316)))
									*(*int32)(unsafe.Add(mBase, uint32(v316))) = v333 + v330
									v428 = int32(0)
									v431 = v309
									for {
										v440 = *(*int32)(unsafe.Add(mBase, uint32(v12-int32(-64)+v431<<(uint(int32(2))%32))))
										v441 = int32(3)
										v447 = *(*int64)(unsafe.Add(mBase, uint32(v12+int32(32)+v431<<(uint(v441)%32))))
										*(*int64)(unsafe.Add(mBase, uint32(v316+int32(232)+v440<<(uint(v441)%32)))) = v447
										v449 = int32(1)
										v452 = v428 + v449
										if v452 != int32(3) {
											v428 = v452
											v431 = v431 + v449
											continue
										} else {
											break
										}
										break
									}
									v463 = *(*int32)(unsafe.Add(mBase, uint32(v316)))
									v464 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v316))) = v463 + v464
									v467 = int32(_a_F_validate_index_0)
									v469 = *(*int32)(unsafe.Add(mBase, _c_F_validate_index[5]))
									*(*int32)(unsafe.Add(mBase, _c_F_validate_index[5])) = v469 - v464
								}
							}
							v482 = *(*int32)(unsafe.Add(mBase, uint32(v12)+128))
							F_tuplesort_performsort(m, v482)
							mBase = m.M
							v484 = m.ExcPending
							if v484 != 0 {
								return
							} else {
								v489 = *(*int32)(unsafe.Add(mBase, _c_F_validate_index[3]))
								if v489 == int32(0) {
								} else {
									v493 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_validate_index[4])))
									if v493 != int32(1) {
									} else {
										v496 = int32(_a_F_validate_index_0)
										v498 = *(*int32)(unsafe.Add(mBase, _c_F_validate_index[5]))
										v499 = int32(1)
										*(*int32)(unsafe.Add(mBase, _c_F_validate_index[5])) = v498 + v499
										v502 = *(*int32)(unsafe.Add(mBase, uint32(v489)))
										*(*int32)(unsafe.Add(mBase, uint32(v489))) = v502 + v499
										*(*int64)(unsafe.Add(mBase, uint32(v489+int32(72))+232)) = int64(6)
										v510 = *(*int32)(unsafe.Add(mBase, uint32(v489)))
										*(*int32)(unsafe.Add(mBase, uint32(v489))) = v510 + v499
										v516 = *(*int32)(unsafe.Add(mBase, _c_F_validate_index[5]))
										*(*int32)(unsafe.Add(mBase, _c_F_validate_index[5])) = v516 - v499
									}
								}
								v522 = *(*int32)(unsafe.Add(mBase, uint32(v214)+188))
								v523 = *(*int32)(unsafe.Add(mBase, uint32(v522)+144))
								m.T0[v523].(func(*base.Module, int32, int32, int32, int32, int32))(m, v214, v245, v247, l2, v12+int32(128))
								mBase = m.M
								v525 = m.ExcPending
								if v525 != 0 {
									return
								} else {
									v526 = *(*int32)(unsafe.Add(mBase, uint32(v12)+128))
									F_tuplesort_end(m, v526)
									mBase = m.M
									v528 = m.ExcPending
									if v528 != 0 {
										return
									} else {
										F_index_insert_cleanup(m, v245, v247)
										mBase = m.M
										v530 = m.ExcPending
										if v530 != 0 {
											return
										} else {
											v533 = F_errstart(m, int32(13), int32(0))
											mBase = m.M
											v534 = m.ExcPending
											if v534 != 0 {
												return
											} else {
												if v533 != 0 {
													v535 = *(*float64)(unsafe.Add(mBase, uint32(v12)+152))
													*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v535
													v537 = *(*float64)(unsafe.Add(mBase, uint32(v12)+136))
													*(*float64)(unsafe.Add(mBase, uint32(v12))) = v537
													v539 = *(*float64)(unsafe.Add(mBase, uint32(v12)+144))
													*(*float64)(unsafe.Add(mBase, uint32(v12)+8)) = v539
													F_errmsg_internal(m, int32(_a_F_validate_index_2), v12)
													mBase = m.M
													v543 = m.ExcPending
													if v543 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_validate_index_3), int32(3466), int32(_a_F_validate_index_4))
														mBase = m.M
														v548 = m.ExcPending
														if v548 != 0 {
															return
														} else {
															F_AtEOXact_GUC(m, int32(0), v240)
															mBase = m.M
															v551 = m.ExcPending
															if v551 != 0 {
																return
															} else {
																v552 = *(*int32)(unsafe.Add(mBase, uint32(v12)+124))
																v553 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
																*(*int32)(unsafe.Add(mBase, _c_F_validate_index[7])) = v553
																*(*int32)(unsafe.Add(mBase, _c_F_validate_index[6])) = v552
																F_relation_close(m, v245, int32(0))
																mBase = m.M
																v560 = m.ExcPending
																if v560 != 0 {
																	return
																} else {
																	F_sequence_close(m, v214, int32(0))
																	mBase = m.M
																	v563 = m.ExcPending
																	if v563 != 0 {
																		return
																	} else {
																		m.G0 = v12 + int32(160)
																		return
																	}
																}
															}
														}
													}
												} else {
													F_AtEOXact_GUC(m, int32(0), v240)
													mBase = m.M
													v551 = m.ExcPending
													if v551 != 0 {
														return
													} else {
														v552 = *(*int32)(unsafe.Add(mBase, uint32(v12)+124))
														v553 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
														*(*int32)(unsafe.Add(mBase, _c_F_validate_index[7])) = v553
														*(*int32)(unsafe.Add(mBase, _c_F_validate_index[6])) = v552
														F_relation_close(m, v245, int32(0))
														mBase = m.M
														v560 = m.ExcPending
														if v560 != 0 {
															return
														} else {
															F_sequence_close(m, v214, int32(0))
															mBase = m.M
															v563 = m.ExcPending
															if v563 != 0 {
																return
															} else {
																m.G0 = v12 + int32(160)
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
			}
		}
	}
}
