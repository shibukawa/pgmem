package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgss_ExecutorEnd(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
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
	var v31 int32
	_ = v31
	var v32 float64
	_ = v32
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
	if v6 == int64(0) {
		v54 = *(*int32)(unsafe.Add(mBase, _consts[1611]))
		if v54 != 0 {
			m.T0[v54].(func(*base.Module, int32))(m, l0)
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return
			} else {
				return
			}
		} else {
			F_standard_ExecutorEnd(m, l0)
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		if v10 == int32(0) {
			v54 = *(*int32)(unsafe.Add(mBase, _consts[1611]))
			if v54 != 0 {
				m.T0[v54].(func(*base.Module, int32))(m, l0)
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return
				} else {
					return
				}
			} else {
				F_standard_ExecutorEnd(m, l0)
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, _consts[181]))
			if int32(0) <= v13 {
				v54 = *(*int32)(unsafe.Add(mBase, _consts[1611]))
				if v54 != 0 {
					m.T0[v54].(func(*base.Module, int32))(m, l0)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						return
					}
				} else {
					F_standard_ExecutorEnd(m, l0)
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, _consts[1612]))
				if v17 != int32(2) {
					if v17 != int32(1) {
						v54 = *(*int32)(unsafe.Add(mBase, _consts[1611]))
						if v54 != 0 {
							m.T0[v54].(func(*base.Module, int32))(m, l0)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return
							} else {
								return
							}
						} else {
							F_standard_ExecutorEnd(m, l0)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return
							} else {
								return
							}
						}
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, _consts[1609]))
						if v23 != 0 {
							v54 = *(*int32)(unsafe.Add(mBase, _consts[1611]))
							if v54 != 0 {
								m.T0[v54].(func(*base.Module, int32))(m, l0)
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return
								} else {
									return
								}
							} else {
								F_standard_ExecutorEnd(m, l0)
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return
								} else {
									return
								}
							}
						} else {
							F_InstrEndLoop(m, v10)
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return
							} else {
								v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+92))
								v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+96))
								v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
								v32 = *(*float64)(unsafe.Add(mBase, uint32(v31)+208))
								v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								v36 = *(*int64)(unsafe.Add(mBase, uint32(v35)+120))
								v41 = *(*int32)(unsafe.Add(mBase, uint32(v35)+180))
								if v41 != 0 {
									v45 = v41 + int32(8)
								} else {
									v45 = int32(0)
								}
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v35)+164))
								v48 = *(*int32)(unsafe.Add(mBase, uint32(v35)+168))
								F_pgss_store(m, v26, v6, v28, v29, int32(1), base.F64_mul(v32, float64(1000)), v36, v31+int32(256), v31+int32(384), v45, int32(0), v47, v48)
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return
								} else {
									v54 = *(*int32)(unsafe.Add(mBase, _consts[1611]))
									if v54 != 0 {
										m.T0[v54].(func(*base.Module, int32))(m, l0)
										mBase = m.M
										v56 = m.ExcPending
										if v56 != 0 {
											return
										} else {
											return
										}
									} else {
										F_standard_ExecutorEnd(m, l0)
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
											return
										} else {
											return
										}
									}
								}
							}
						}
					}
				} else {
					F_InstrEndLoop(m, v10)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+92))
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+96))
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						v32 = *(*float64)(unsafe.Add(mBase, uint32(v31)+208))
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						v36 = *(*int64)(unsafe.Add(mBase, uint32(v35)+120))
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v35)+180))
						if v41 != 0 {
							v45 = v41 + int32(8)
						} else {
							v45 = int32(0)
						}
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v35)+164))
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v35)+168))
						F_pgss_store(m, v26, v6, v28, v29, int32(1), base.F64_mul(v32, float64(1000)), v36, v31+int32(256), v31+int32(384), v45, int32(0), v47, v48)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							v54 = *(*int32)(unsafe.Add(mBase, _consts[1611]))
							if v54 != 0 {
								m.T0[v54].(func(*base.Module, int32))(m, l0)
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return
								} else {
									return
								}
							} else {
								F_standard_ExecutorEnd(m, l0)
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return
								} else {
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
func F_pgss_ExecutorRun(m *base.Module, l0 int32, l1 int32, l2 int64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int64
	_ = v78
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	v4 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v20 = v4
	v21 = v14
	v22 = int32(-1)
	v24 = v4
	v25 = v4
	v26 = v4
	goto L1
L1:
	;
	if v22 != int32(1) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, _consts[426])) = v51
	*(*int32)(unsafe.Add(mBase, _consts[394])) = v52
	v102 = int32(4735024)
	v103 = *(*int32)(unsafe.Add(mBase, _consts[1609]))
	*(*int32)(unsafe.Add(mBase, _consts[1609])) = v103 - int32(1)
	m.G0 = v14 + int32(16)
	return
L3:
	;
	v31 = v21 - int32(160)
	m.G0 = v31
	v33 = int32(4735024)
	v34 = *(*int32)(unsafe.Add(mBase, _consts[1609]))
	v35 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[1609])) = v34 + v35
	v39 = *(*int32)(unsafe.Add(mBase, _consts[426]))
	v41 = *(*int32)(unsafe.Add(mBase, _consts[394]))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v14 + int32(12)
	goto L6
L4:
	;
	v48 = v20
	v49 = v21
	v50 = v24
	v51 = v25
	v52 = v26
	goto L5
L5:
	;
	goto L8
L6:
	;
	v48 = int32(0)
	v49 = v31
	v50 = v31
	v51 = v39
	v52 = v41
	goto L5
L7:
	;
	goto L2
L8:
	;
	if v48 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	goto L7
L10:
	;
	v77 = int32(m.ExcTag)
	v78 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v77 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L11:
	;
	F_standard_ExecutorRun(m, l0, l1, l2)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L10
	} else {
		goto L18
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, _consts[426])) = v50
	v58 = *(*int32)(unsafe.Add(mBase, _consts[1610]))
	if v58 == int32(0) {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, _consts[394])) = v52
	*(*int32)(unsafe.Add(mBase, _consts[426])) = v51
	v67 = int32(4735024)
	v68 = *(*int32)(unsafe.Add(mBase, _consts[1609]))
	*(*int32)(unsafe.Add(mBase, _consts[1609])) = v68 - int32(1)
	F_pg_re_throw(m)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L10
	} else {
		goto L17
	}
L15:
	;
	m.T0[v58].(func(*base.Module, int32, int32, int64))(m, l0, l1, l2)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	goto L7
L17:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L18:
	;
	goto L9
L19:
	;
	v82 = int32(v78)
	m.G0 = v49
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	if v14+int32(12) == v89 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	m.ExcPending = 1
	goto L28
L21:
	;
	if v92 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	v92 = v91
	goto L24
L23:
	;
	v92 = int32(0)
	goto L24
L24:
	;
	goto L21
L25:
	;
	F___wasm_longjmp(m, v85, v84)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v20 = v84
	v21 = v49
	v22 = v92
	v24 = v50
	v25 = v51
	v26 = v52
	goto L1
L28:
	;
	return
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pgss_shmem_request(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	v4 = *(*int32)(unsafe.Add(mBase, _consts[1604]))
	if v4 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.T0[v4].(func(*base.Module))(m)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v9 = *(*int32)(unsafe.Add(mBase, _consts[1605]))
	v11 = F_hash_estimate_size(m, v9, int32(432))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L4
	} else {
		goto L6
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	v13 = F_add_size(m, int32(56), v11)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1606])))
	if v16 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v32 = int32(4478696)
	v34 = *(*int32)(unsafe.Add(mBase, _consts[1607]))
	v35 = F_add_size(m, v34, v13)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L14
	}
L11:
	;
	F_errmsg_internal(m, int32(331193), int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	F_errfinish(m, int32(523622), int32(77), int32(440847))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1607])) = v35
	v39 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1606])))
	if v39 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	return
L16:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _consts[977]))
	if v41 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L4
	} else {
		goto L62
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1608])) = int32(16)
	v49 = *(*int32)(unsafe.Add(mBase, _consts[36]))
	v51 = F_MemoryContextAlloc(m, v49, int32(1088))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L4
	} else {
		goto L22
	}
L20:
	;
	v54 = v41
	goto L21
L21:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _consts[976]))
	v58 = *(*int32)(unsafe.Add(mBase, _consts[1608]))
	if v58 <= v56 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, _consts[977])) = v51
	v54 = v51
	goto L21
L23:
	;
	v60 = int32(1)
	v63 = v56 + v60
	if v63&v56 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v79 = v54
	v81 = v56
	goto L25
L25:
	;
	v84 = v81*int32(68) + v79
	v85 = int32(131133)
	goto L33
L26:
	;
	v68 = v60 << (uint(int32(32)-base.I32_clz(v63)) % 32)
	goto L28
L27:
	;
	v68 = v63
	goto L28
L28:
	;
	v71 = F_repalloc(m, v54, v68*int32(68))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1608])) = v68
	*(*int32)(unsafe.Add(mBase, _consts[977])) = v71
	v78 = *(*int32)(unsafe.Add(mBase, _consts[976]))
	v79 = v71
	v81 = v78
	goto L25
L30:
	;
	v201 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v84)+64)) = v201
	v203 = int32(4483720)
	v205 = *(*int32)(unsafe.Add(mBase, _consts[976]))
	*(*int32)(unsafe.Add(mBase, _consts[976])) = v205 + v201
	goto L15
L31:
	;
	v198 = F_strlen(m, v187)
	mBase = m.M
	goto L30
L33:
	;
	goto L34
L34:
	;
	v92 = int32(63)
	if (v84^v85)&int32(3) != 0 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v191 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v188))) = uint8(v191)
	goto L31
L36:
	;
	v172 = v167
	v173 = v168
	v174 = v169
	goto L58
L37:
	;
	if v162 == int32(0) {
		v187 = v160
		v188 = v161
		goto L35
	} else {
		goto L57
	}
L38:
	;
	v160 = v85
	v161 = v84
	v162 = v92
	goto L37
L39:
	;
	goto L40
L40:
	;
	goto L43
L41:
	;
	if base.B2i32(v116 != v117) == int32(0) {
		v187 = v120
		v188 = v114
		goto L35
	} else {
		goto L50
	}
L43:
	;
	goto L44
L44:
	;
	v105 = v85
	v106 = v84
	v107 = v92
	goto L45
L45:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
	*(*uint8)(unsafe.Add(mBase, uint32(v106))) = uint8(v109)
	if v109 == int32(0) {
		v167 = v105
		v168 = v106
		v169 = v107
		goto L36
	} else {
		goto L47
	}
L46:
	;
	goto L41
L47:
	;
	v113 = int32(1)
	v114 = v106 + v113
	v116 = v107 - v113
	v117 = int32(0)
	v120 = v105 + v113
	if v120&int32(3) == v117 {
		goto L41
	} else {
		goto L48
	}
L48:
	;
	if v116 != 0 {
		v105 = v120
		v106 = v114
		v107 = v116
		goto L45
	} else {
		goto L49
	}
L49:
	;
	goto L46
L50:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
	if v132 == int32(0) {
		v160 = v120
		v161 = v114
		v162 = v116
		goto L37
	} else {
		goto L51
	}
L51:
	;
	if base.Ui32(v116) < base.Ui32(int32(4)) {
		v160 = v120
		v161 = v114
		v162 = v116
		goto L37
	} else {
		goto L52
	}
L52:
	;
	v138 = v120
	v139 = v114
	v140 = v116
	goto L53
L53:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v146 = int32(-2139062144)
	if (int32(16843008)-v143|v143)&v146 != v146 {
		v167 = v138
		v168 = v139
		v169 = v140
		goto L36
	} else {
		goto L55
	}
L54:
	;
	v160 = v154
	v161 = v152
	v162 = v156
	goto L37
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v139))) = v143
	v151 = int32(4)
	v152 = v139 + v151
	v154 = v138 + v151
	v156 = v140 - v151
	if base.Ui32(int32(3)) < base.Ui32(v156) {
		v138 = v154
		v139 = v152
		v140 = v156
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v167 = v160
	v168 = v161
	v169 = v162
	goto L36
L58:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
	*(*uint8)(unsafe.Add(mBase, uint32(v173))) = uint8(v176)
	if v176 == int32(0) {
		v187 = v172
		v188 = v173
		goto L35
	} else {
		goto L60
	}
L59:
	;
	v187 = v183
	v188 = v181
	goto L35
L60:
	;
	v180 = int32(1)
	v181 = v173 + v180
	v183 = v172 + v180
	v185 = v174 - v180
	if v185 != 0 {
		v172 = v183
		v173 = v181
		v174 = v185
		goto L58
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	F_errmsg_internal(m, int32(331260), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(523489), int32(687), int32(418747))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
