package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_CombineRangeTables(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v9 == int32(0) {
		v55 = int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
		if v13 <= int32(0) {
			v55 = v9
		} else {
			if l2 == int32(0) {
				v55 = v9
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
				if int32(0) < v18 {
					v26 = int32(0)
					for {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v30+v26<<(uint(int32(2))%32))))
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+28))
						if v35 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v34)+28)) = v13 + v35
						} else {
						}
						v39 = v26 + int32(1)
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
						if v39 < v40 {
							v26 = v39
							continue
						} else {
							break
						}
						break
					}
				} else {
				}
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v55 = v50
			}
		}
	}
	v59 = F_list_concat(m, v55, l3)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v59
		v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v63 = F_list_concat(m, v62, l2)
		mBase = m.M
		v64 = m.ExcPending
		if v64 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v63
			return
		}
	}
}
func F_ConditionVariableTimedSleep(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int64
	_ = v30
	var v33 int64
	_ = v33
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int64
	_ = v113
	var v114 int64
	_ = v114
	var v121 float64
	_ = v121
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableTimedSleep[0]))
	if l0 != v18 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v15 + int32(16)
	return v137
L2:
	;
	F_ConditionVariablePrepareToSleep(m, l0)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v25 = base.B2i32(l1 < int32(0))
	if l1 < int32(0) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	return int32(0)
L6:
	;
	v137 = int32(0)
	goto L1
L7:
	;
	v47 = v38
	goto L11
L8:
	;
	v36 = int32(33)
	v37 = int64(0)
	v38 = int32(-1)
	goto L7
L9:
	;
	goto L10
L10:
	;
	F___clock_gettime(m, int32(1), v15)
	mBase = m.M
	v30 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
	v33 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15)+8)))
	v36 = int32(41)
	v37 = v30*int64(-1000000000) - v33
	v38 = l1
	goto L7
L11:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableTimedSleep[1]))
	v53 = F_WaitLatch(m, v52, v36, v47, l2)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L5
	} else {
		goto L13
	}
L12:
	;
	v137 = int32(0)
	goto L1
L13:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableTimedSleep[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = int32(0)
	goto L14
L14:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
	if v59 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	F_s_lock(m, l0, int32(_a_F_ConditionVariableTimedSleep_0), int32(183), int32(_a_F_ConditionVariableTimedSleep_1))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L5
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v67 = int32(0)
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableTimedSleep[2]))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableTimedSleep[3]))
	v75 = v70 + v72*int32(640)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+88))
	if v76 != 0 {
		v98 = v67
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L17
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableTimedSleep[4]))
	if v103 != 0 {
		goto L26
	} else {
		goto L27
	}
L20:
	;
	v78 = v75 + int32(84)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	if v79 != 0 {
		v98 = v67
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v80 == int32(-1) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v72
	v98 = int32(1)
	goto L19
L23:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v78))) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v72
	goto L22
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78)+4)) = v80
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableTimedSleep[2]))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	*(*int32)(unsafe.Add(mBase, uint32(v89+v80*int32(640))+84)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = int32(-1)
	goto L22
L26:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L5
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableTimedSleep[0]))
	v109 = v98 | base.B2i32(l0 != v107)
	if l1 < int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L28
L30:
	;
	if v109 == int32(0) {
		goto L11
	} else {
		goto L38
	}
L31:
	;
	if v109 != 0 {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v110 = int32(1)
	F___clock_gettime(m, v110, v15)
	mBase = m.M
	v113 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15)+8)))
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
	v121 = base.F64_div(base.F64_convert_i64_s(v113+(v114*int64(1000000000)+v37)), float64(1e+06))
	if base.F64_lt(base.F64_abs(v121), float64(2.147483648e+09)) != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v128 = l1 - v127
	if int32(0) < v128 {
		v47 = v128
		goto L11
	} else {
		goto L37
	}
L34:
	;
	v125 = base.I32_trunc_f64_s(v121)
	v127 = v125
	goto L33
L35:
	;
	goto L36
L36:
	;
	v127 = int32(-2147483648)
	goto L33
L37:
	;
	v137 = v110
	goto L1
L38:
	;
	goto L12
}
func F_CopySendEndOfRow(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
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
	var v81 int32
	_ = v81
	var v82 int64
	_ = v82
	var v83 int64
	_ = v83
	var v84 int64
	_ = v84
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v5 {
	case 0:
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v10 = F_fwrite(m, v6, v7, int32(1), v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			if v10 == int32(1) {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
				if v15 < int32(0) {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					v20 = v18
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					v20 = v19
				}
				if int32(base.Ui32(v20)>>(uint(int32(5))%32))&int32(1) == int32(0) {
					v82 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
					v83 = int64(*(*int32)(unsafe.Add(mBase, uint32(v4)+4)))
					v84 = v82 + v83
					*(*int64)(unsafe.Add(mBase, uint32(l0)+176)) = v84
					v89 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[0]))
					if v89 == int32(0) {
					} else {
						v93 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[1])))
						if v93 != int32(1) {
						} else {
							v96 = int32(_a_F_CopySendEndOfRow_0)
							v98 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[2]))
							v99 = int32(1)
							*(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[2])) = v98 + v99
							v102 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
							*(*int32)(unsafe.Add(mBase, uint32(v89))) = v102 + v99
							*(*int64)(unsafe.Add(mBase, uint32(v89+int32(0))+232)) = v84
							v110 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
							*(*int32)(unsafe.Add(mBase, uint32(v89))) = v110 + v99
							v116 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[2]))
							*(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[2])) = v116 - v99
						}
					}
					v120 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
					v121 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v120))) = uint8(v121)
					*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = v121
					*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = v121
					return
				} else {
					v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
					if v27 == int32(1) {
						v31 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[3]))
						if v31 == int32(64) {
							F_ClosePipeToProgram(m, l0)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[3])) = int32(64)
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return
								} else {
									F_errcode_for_file_access(m)
									mBase = m.M
									v44 = m.ExcPending
									if v44 != 0 {
										return
									} else {
										F_errmsg(m, int32(_a_F_CopySendEndOfRow_1), int32(0))
										mBase = m.M
										v48 = m.ExcPending
										if v48 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_CopySendEndOfRow_2), int32(477), int32(_a_F_CopySendEndOfRow_3))
											mBase = m.M
											v53 = m.ExcPending
											if v53 != 0 {
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
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return
							} else {
								F_errcode_for_file_access(m)
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return
								} else {
									F_errmsg(m, int32(_a_F_CopySendEndOfRow_1), int32(0))
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_CopySendEndOfRow_2), int32(477), int32(_a_F_CopySendEndOfRow_3))
										mBase = m.M
										v53 = m.ExcPending
										if v53 != 0 {
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
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return
							} else {
								F_errmsg(m, int32(_a_F_CopySendEndOfRow_4), int32(0))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_CopySendEndOfRow_2), int32(482), int32(_a_F_CopySendEndOfRow_3))
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
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
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
				if v27 == int32(1) {
					v31 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[3]))
					if v31 == int32(64) {
						F_ClosePipeToProgram(m, l0)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[3])) = int32(64)
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return
							} else {
								F_errcode_for_file_access(m)
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return
								} else {
									F_errmsg(m, int32(_a_F_CopySendEndOfRow_1), int32(0))
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_CopySendEndOfRow_2), int32(477), int32(_a_F_CopySendEndOfRow_3))
										mBase = m.M
										v53 = m.ExcPending
										if v53 != 0 {
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
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return
							} else {
								F_errmsg(m, int32(_a_F_CopySendEndOfRow_1), int32(0))
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_CopySendEndOfRow_2), int32(477), int32(_a_F_CopySendEndOfRow_3))
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
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
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return
					} else {
						F_errcode_for_file_access(m)
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return
						} else {
							F_errmsg(m, int32(_a_F_CopySendEndOfRow_4), int32(0))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_CopySendEndOfRow_2), int32(482), int32(_a_F_CopySendEndOfRow_3))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
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
		}
	case 1:
		v70 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v71 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
		v73 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[4]))
		v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+16))
		v75 = m.T0[v74].(func(*base.Module, int32, int32, int32) int32)(m, int32(100), v70, v71)
		mBase = m.M
		v76 = m.ExcPending
		if v76 != 0 {
			return
		} else {
			v82 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
			v83 = int64(*(*int32)(unsafe.Add(mBase, uint32(v4)+4)))
			v84 = v82 + v83
			*(*int64)(unsafe.Add(mBase, uint32(l0)+176)) = v84
			v89 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[0]))
			if v89 == int32(0) {
			} else {
				v93 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[1])))
				if v93 != int32(1) {
				} else {
					v96 = int32(_a_F_CopySendEndOfRow_0)
					v98 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[2]))
					v99 = int32(1)
					*(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[2])) = v98 + v99
					v102 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
					*(*int32)(unsafe.Add(mBase, uint32(v89))) = v102 + v99
					*(*int64)(unsafe.Add(mBase, uint32(v89+int32(0))+232)) = v84
					v110 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
					*(*int32)(unsafe.Add(mBase, uint32(v89))) = v110 + v99
					v116 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[2]))
					*(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[2])) = v116 - v99
				}
			}
			v120 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
			v121 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v120))) = uint8(v121)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = v121
			*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = v121
			return
		}
	case 2:
		v77 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v78 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
		v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		m.T0[v79].(func(*base.Module, int32, int32))(m, v77, v78)
		mBase = m.M
		v81 = m.ExcPending
		if v81 != 0 {
			return
		} else {
			v82 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
			v83 = int64(*(*int32)(unsafe.Add(mBase, uint32(v4)+4)))
			v84 = v82 + v83
			*(*int64)(unsafe.Add(mBase, uint32(l0)+176)) = v84
			v89 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[0]))
			if v89 == int32(0) {
			} else {
				v93 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[1])))
				if v93 != int32(1) {
				} else {
					v96 = int32(_a_F_CopySendEndOfRow_0)
					v98 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[2]))
					v99 = int32(1)
					*(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[2])) = v98 + v99
					v102 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
					*(*int32)(unsafe.Add(mBase, uint32(v89))) = v102 + v99
					*(*int64)(unsafe.Add(mBase, uint32(v89+int32(0))+232)) = v84
					v110 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
					*(*int32)(unsafe.Add(mBase, uint32(v89))) = v110 + v99
					v116 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[2]))
					*(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[2])) = v116 - v99
				}
			}
			v120 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
			v121 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v120))) = uint8(v121)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = v121
			*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = v121
			return
		}
	default:
		v82 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
		v83 = int64(*(*int32)(unsafe.Add(mBase, uint32(v4)+4)))
		v84 = v82 + v83
		*(*int64)(unsafe.Add(mBase, uint32(l0)+176)) = v84
		v89 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[0]))
		if v89 == int32(0) {
		} else {
			v93 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[1])))
			if v93 != int32(1) {
			} else {
				v96 = int32(_a_F_CopySendEndOfRow_0)
				v98 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[2]))
				v99 = int32(1)
				*(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[2])) = v98 + v99
				v102 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
				*(*int32)(unsafe.Add(mBase, uint32(v89))) = v102 + v99
				*(*int64)(unsafe.Add(mBase, uint32(v89+int32(0))+232)) = v84
				v110 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
				*(*int32)(unsafe.Add(mBase, uint32(v89))) = v110 + v99
				v116 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[2]))
				*(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[2])) = v116 - v99
			}
		}
		v120 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v121 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v120))) = uint8(v121)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = v121
		*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = v121
		return
	}
}
func F_colorcomplement(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
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
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v158 int32
	_ = v158
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v244 int32
	_ = v244
	var v255 int32
	_ = v255
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	if v13 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v185 = *(*int32)(unsafe.Add(mBase, _c_F_colorcomplement[0]))
	if v185 != 0 {
		goto L55
	} else {
		goto L56
	}
L2:
	;
	v20 = v13
	goto L5
L3:
	;
	v66 = v12
	goto L4
L4:
	;
	v68 = int32(24)
	v70 = v12 + v11*v68
	if base.Ui32(v70+v68) <= base.Ui32(v66) {
		goto L18
	} else {
		goto L19
	}
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v24 == int32(112) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v34 = v13
	goto L12
L7:
	;
	v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+4)))
	if v27 == int32(_a_F_colorcomplement_0) {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	if v30 != 0 {
		v20 = v30
		goto L5
	} else {
		goto L11
	}
L10:
	;
	goto L9
L11:
	;
	goto L6
L12:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	if v41 == int32(112) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v66 = v57
	goto L4
L14:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v45 = int32(*(*int16)(unsafe.Add(mBase, uint32(v34)+4)))
	v50 = v44 + v45*int32(24) + int32(20)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v51 | int32(4)
	goto L16
L15:
	;
	goto L16
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	if v56 != 0 {
		v34 = v56
		goto L12
	} else {
		goto L17
	}
L17:
	;
	goto L13
L18:
	;
	return
L19:
	;
	v82 = int32(0)
	v83 = v66
	goto L20
L20:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
	if v86 != 0 {
		goto L18
	} else {
		goto L22
	}
L21:
	;
	goto L18
L22:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v83)+20))
	if v87&int32(4) != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if base.Ui32(v83) < base.Ui32(v70) {
		v82 = v82 + int32(1)
		v83 = v83 + int32(24)
		goto L20
	} else {
		goto L54
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+20)) = v87 & int32(-5)
	goto L23
L25:
	;
	goto L26
L26:
	;
	if v87&int32(3) != 0 {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _c_F_colorcomplement[0]))
	if v96 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	if v99 <= v100 {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	return
L32:
	;
	goto L30
L33:
	;
	F_createarc(m, l0, l2, base.I32_extend16_s(v82), l4, l5)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L31
	} else {
		goto L53
	}
L34:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	if v102 == int32(0) {
		goto L33
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
	if v124 == int32(0) {
		goto L33
	} else {
		goto L45
	}
L37:
	;
	v108 = v102
	goto L38
L38:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v108)+12))
	if v115 != l5 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L33
L40:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v108)+16))
	if v123 != 0 {
		v108 = v123
		goto L38
	} else {
		goto L44
	}
L41:
	;
	v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v108)+4)))
	if v117 != v82&int32(_a_F_colorcomplement_1) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	if v121 == l2 {
		goto L23
	} else {
		goto L43
	}
L43:
	;
	goto L40
L44:
	;
	goto L39
L45:
	;
	v130 = v124
	goto L46
L46:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
	if v137 != l4 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L33
L48:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v130)+24))
	if v145 != 0 {
		v130 = v145
		goto L46
	} else {
		goto L52
	}
L49:
	;
	v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v130)+4)))
	if v139 != v82&int32(_a_F_colorcomplement_1) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	if v143 == l2 {
		goto L23
	} else {
		goto L51
	}
L51:
	;
	goto L48
L52:
	;
	goto L47
L53:
	;
	goto L23
L54:
	;
	goto L21
L55:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L31
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	if v188 <= v189 {
		goto L61
	} else {
		goto L62
	}
L58:
	;
	goto L57
L59:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v255 | int32(4)
	return
L60:
	;
	F_createarc(m, l0, int32(120), int32(0), l4, l5)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L31
	} else {
		goto L80
	}
L61:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	if v191 == int32(0) {
		goto L60
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
	if v211 == int32(0) {
		goto L60
	} else {
		goto L72
	}
L64:
	;
	v197 = v191
	goto L65
L65:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v197)+12))
	if v204 != l5 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	goto L60
L67:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v197)+16))
	if v210 != 0 {
		v197 = v210
		goto L65
	} else {
		goto L71
	}
L68:
	;
	v206 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v197)+4)))
	if v206 != 0 {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
	if v207 == int32(120) {
		goto L59
	} else {
		goto L70
	}
L70:
	;
	goto L67
L71:
	;
	goto L66
L72:
	;
	v217 = v211
	goto L73
L73:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v217)+8))
	if v224 != l4 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	goto L60
L75:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v217)+24))
	if v230 != 0 {
		v217 = v230
		goto L73
	} else {
		goto L79
	}
L76:
	;
	v226 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217)+4)))
	if v226 != 0 {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
	if v227 == int32(120) {
		goto L59
	} else {
		goto L78
	}
L78:
	;
	goto L75
L79:
	;
	goto L74
L80:
	;
	goto L59
}
func F_combine(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
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
	var v88 int32
	_ = v88
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v10 = v6 | v7<<(uint(int32(8))%32)
	if v10 <= int32(_a_F_combine_0) {
		if v10 <= int32(_a_F_combine_1) {
			v15 = int32(1)
			switch v10 - int32(_a_F_combine_2) {
			case 0, 18:
				v114 = int32(3)
				return v114
			case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17:
				v114 = v15
				return v114
			default:
				if v10 == int32(_a_F_combine_3) {
					v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
					v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
					if v69 == v70 {
						v72 = int32(2)
					} else {
						v72 = int32(1)
					}
					return v72
				} else {
					if v10 == int32(_a_F_combine_4) {
						v114 = int32(3)
					} else {
						v114 = v15
					}
					return v114
				}
			}
		} else {
			v22 = int32(1)
			switch v10 - int32(_a_F_combine_5) {
			case 0, 21:
				v114 = int32(3)
				return v114
			case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 19, 20:
				v114 = v22
				return v114
			case 18:
				v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
				v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
				if v69 == v70 {
					v72 = int32(2)
				} else {
					v72 = int32(1)
				}
				return v72
			default:
				if v10 != int32(_a_F_combine_6) {
					v114 = v22
				} else {
					v114 = int32(3)
				}
				return v114
			}
		}
	} else {
		v25 = int32(1)
		switch v10 - int32(_a_F_combine_7) {
		case 0, 18, 38:
			v114 = int32(3)
			return v114
		case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 19, 20, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 37:
			v114 = v25
			return v114
		case 21:
			v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
			v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
			if v74 == v75 {
				return int32(2)
			} else {
				if v74 == int32(_a_F_combine_8) {
					v81 = int32(2)
					v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
					v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
					v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+base.I32_extend16_s(v75)*int32(24))+20)))
					if v88&v81 == int32(0) {
						v114 = v81
						return v114
					} else {
						return int32(1)
					}
				} else {
					if v75 != int32(_a_F_combine_8) {
						return int32(1)
					} else {
						v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+20))
						v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98+base.I32_extend16_s(v74)*int32(24))+20)))
						if v103&int32(2) != 0 {
							v106 = int32(1)
						} else {
							v106 = int32(4)
						}
						return v106
					}
				}
			}
		case 36:
			v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
			v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
			if v32 == v33 {
				return int32(2)
			} else {
				if v32 == int32(_a_F_combine_8) {
					v39 = int32(2)
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+20))
					v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+base.I32_extend16_s(v33)*int32(24))+20)))
					if v46&v39 == int32(0) {
						v114 = v39
						return v114
					} else {
						return int32(1)
					}
				} else {
					if v33 != int32(_a_F_combine_8) {
						return int32(1)
					} else {
						v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+20))
						v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+base.I32_extend16_s(v32)*int32(24))+20)))
						if v61&int32(2) != 0 {
							v64 = int32(1)
						} else {
							v64 = int32(4)
						}
						return v64
					}
				}
			}
		default:
			switch v10 - int32(_a_F_combine_9) {
			case 0, 21:
				v114 = int32(3)
				return v114
			case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 37:
				v114 = v25
				return v114
			case 36:
				v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
				v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
				if v32 == v33 {
					return int32(2)
				} else {
					if v32 == int32(_a_F_combine_8) {
						v39 = int32(2)
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+20))
						v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+base.I32_extend16_s(v33)*int32(24))+20)))
						if v46&v39 == int32(0) {
							v114 = v39
							return v114
						} else {
							return int32(1)
						}
					} else {
						if v33 != int32(_a_F_combine_8) {
							return int32(1)
						} else {
							v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+20))
							v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+base.I32_extend16_s(v32)*int32(24))+20)))
							if v61&int32(2) != 0 {
								v64 = int32(1)
							} else {
								v64 = int32(4)
							}
							return v64
						}
					}
				}
			case 38:
				v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
				v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
				if v74 == v75 {
					return int32(2)
				} else {
					if v74 == int32(_a_F_combine_8) {
						v81 = int32(2)
						v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
						v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+base.I32_extend16_s(v75)*int32(24))+20)))
						if v88&v81 == int32(0) {
							v114 = v81
							return v114
						} else {
							return int32(1)
						}
					} else {
						if v75 != int32(_a_F_combine_8) {
							return int32(1)
						} else {
							v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
							v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+20))
							v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98+base.I32_extend16_s(v74)*int32(24))+20)))
							if v103&int32(2) != 0 {
								v106 = int32(1)
							} else {
								v106 = int32(4)
							}
							return v106
						}
					}
				}
			default:
				if v10 == int32(_a_F_combine_10) {
					v114 = int32(3)
				} else {
					v114 = v25
				}
				return v114
			}
		}
	}
}
func F_compactify_tuples(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
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
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v300 int32
	_ = v300
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v404 int32
	_ = v404
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v598 int32
	_ = v598
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v609 int32
	_ = v609
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v621 int32
	_ = v621
	v5 = int32(0)
	v14 = m.G0
	v16 = v14 + int32(-8192)
	m.G0 = v16
	if l3 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l2)+14)) = uint16(v621)
	m.G0 = v16 - int32(-8192)
	return
L2:
	;
	v18 = int32(1)
	if l1 <= v18 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v404 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v404) {
		goto L118
	} else {
		goto L119
	}
L5:
	;
	v21 = v18
	goto L7
L6:
	;
	v21 = l1
	goto L7
L7:
	;
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+16)))
	v26 = v22
	v28 = v5
	goto L9
L8:
	;
	v257 = l2 + v247
	v258 = l2 + v248
	v259 = v251 - v248
	if v257 == v258 {
		goto L72
	} else {
		goto L73
	}
L9:
	;
	v38 = l0 + v28*int32(6)
	v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+4)))
	v40 = int32(*(*int16)(unsafe.Add(mBase, uint32(v38)+2)))
	v41 = v39 + v40
	if v26 == v41 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if l1 <= v28 {
		goto L15
	} else {
		goto L16
	}
L11:
	;
	v43 = v26 - v39
	v45 = v28 + int32(1)
	if v45 != v21 {
		v26 = v43
		v28 = v45
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	goto L10
L14:
	;
	v247 = v43
	v248 = v26
	v251 = v26
	goto L8
L15:
	;
	v247 = v26
	v248 = v41
	v251 = v41
	goto L8
L16:
	;
	goto L17
L17:
	;
	v53 = v26
	v54 = v41
	v55 = v28
	v57 = v41
	goto L18
L18:
	;
	v65 = l0 + v55*int32(6)
	v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65))))
	v75 = (v66+int32(1))&int32(_a_F_compactify_tuples_0)<<(uint(int32(2))%32) + (l2 + int32(24)) - int32(4)
	v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65)+4)))
	v77 = int32(*(*int16)(unsafe.Add(mBase, uint32(v65)+2)))
	if v76+v77 == v54 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v247 = v236
	v248 = v230
	v251 = v231
	goto L8
L20:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v236 = v53 - v232
	*(*int32)(unsafe.Add(mBase, uint32(v75))) = v233&int32(-32768) | v236&int32(_a_F_compactify_tuples_1)
	v242 = v55 + int32(1)
	if v242 != l1 {
		v53 = v236
		v54 = v230
		v55 = v242
		v57 = v231
		goto L18
	} else {
		goto L70
	}
L21:
	;
	v230 = v77
	v231 = v57
	v232 = v76
	goto L20
L22:
	;
	goto L23
L23:
	;
	v80 = l2 + v53
	v81 = l2 + v54
	v82 = v57 - v54
	if v80 == v81 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v227 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65)+4)))
	v228 = int32(*(*int16)(unsafe.Add(mBase, uint32(v65)+2)))
	v230 = v228
	v231 = v227 + v228
	v232 = v227
	goto L20
L25:
	;
	goto L24
L26:
	;
	v86 = v80 + v82
	if base.Ui32(v81-v86) <= base.Ui32(int32(0)-v82<<(uint(int32(1))%32)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v93 = F___memcpy(m, v80, v81, v82)
	mBase = m.M
	goto L24
L28:
	;
	goto L29
L29:
	;
	v96 = (v80 ^ v81) & int32(3)
	if base.Ui32(v80) < base.Ui32(v81) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	if v198 == int32(0) {
		goto L25
	} else {
		goto L66
	}
L31:
	;
	if base.Ui32(v176) <= base.Ui32(int32(3)) {
		v197 = v175
		v198 = v176
		v199 = v177
		goto L30
	} else {
		goto L62
	}
L32:
	;
	if v96 != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	if v96 != 0 {
		v158 = v82
		goto L45
	} else {
		goto L46
	}
L35:
	;
	v197 = v81
	v198 = v82
	v199 = v80
	goto L30
L36:
	;
	goto L37
L37:
	;
	if v80&int32(3) == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v175 = v81
	v176 = v82
	v177 = v80
	goto L31
L39:
	;
	goto L40
L40:
	;
	v103 = v81
	v104 = v82
	v105 = v80
	goto L41
L41:
	;
	if v104 == int32(0) {
		goto L25
	} else {
		goto L43
	}
L42:
	;
	v175 = v112
	v176 = v114
	v177 = v116
	goto L31
L43:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	*(*uint8)(unsafe.Add(mBase, uint32(v105))) = uint8(v109)
	v111 = int32(1)
	v112 = v103 + v111
	v114 = v104 - v111
	v116 = v105 + v111
	if v116&int32(3) != 0 {
		v103 = v112
		v104 = v114
		v105 = v116
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	if v158 == int32(0) {
		goto L25
	} else {
		goto L58
	}
L46:
	;
	if v86&int32(3) != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v123 = v82
	goto L50
L48:
	;
	v138 = v82
	goto L49
L49:
	;
	if base.Ui32(v138) <= base.Ui32(int32(3)) {
		v158 = v138
		goto L45
	} else {
		goto L54
	}
L50:
	;
	if v123 == int32(0) {
		goto L25
	} else {
		goto L52
	}
L51:
	;
	v138 = v129
	goto L49
L52:
	;
	v129 = v123 - int32(1)
	v130 = v80 + v129
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81+v129))))
	*(*uint8)(unsafe.Add(mBase, uint32(v130))) = uint8(v132)
	if v130&int32(3) != 0 {
		v123 = v129
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v145 = v138
	goto L55
L55:
	;
	v149 = v145 - int32(4)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v81+v149)))
	*(*int32)(unsafe.Add(mBase, uint32(v80+v149))) = v152
	if base.Ui32(int32(3)) < base.Ui32(v149) {
		v145 = v149
		goto L55
	} else {
		goto L57
	}
L56:
	;
	v158 = v149
	goto L45
L57:
	;
	goto L56
L58:
	;
	v165 = v158
	goto L59
L59:
	;
	v169 = v165 - int32(1)
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81+v169))))
	*(*uint8)(unsafe.Add(mBase, uint32(v80+v169))) = uint8(v172)
	if v169 != 0 {
		v165 = v169
		goto L59
	} else {
		goto L61
	}
L60:
	;
	goto L25
L61:
	;
	goto L60
L62:
	;
	v182 = v175
	v183 = v176
	v184 = v177
	goto L63
L63:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
	*(*int32)(unsafe.Add(mBase, uint32(v184))) = v186
	v188 = int32(4)
	v189 = v182 + v188
	v191 = v184 + v188
	v193 = v183 - v188
	if base.Ui32(int32(3)) < base.Ui32(v193) {
		v182 = v189
		v183 = v193
		v184 = v191
		goto L63
	} else {
		goto L65
	}
L64:
	;
	v197 = v189
	v198 = v193
	v199 = v191
	goto L30
L65:
	;
	goto L64
L66:
	;
	v204 = v197
	v205 = v198
	v206 = v199
	goto L67
L67:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	*(*uint8)(unsafe.Add(mBase, uint32(v206))) = uint8(v208)
	v210 = int32(1)
	v215 = v205 - v210
	if v215 != 0 {
		v204 = v204 + v210
		v205 = v215
		v206 = v206 + v210
		goto L67
	} else {
		goto L69
	}
L68:
	;
	goto L25
L69:
	;
	goto L68
L70:
	;
	goto L19
L71:
	;
	v621 = v247
	goto L1
L72:
	;
	goto L71
L73:
	;
	v263 = v257 + v259
	if base.Ui32(v258-v263) <= base.Ui32(int32(0)-v259<<(uint(int32(1))%32)) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v270 = F___memcpy(m, v257, v258, v259)
	mBase = m.M
	goto L71
L75:
	;
	goto L76
L76:
	;
	v273 = (v257 ^ v258) & int32(3)
	if base.Ui32(v257) < base.Ui32(v258) {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	if v375 == int32(0) {
		goto L72
	} else {
		goto L113
	}
L78:
	;
	if base.Ui32(v353) <= base.Ui32(int32(3)) {
		v374 = v352
		v375 = v353
		v376 = v354
		goto L77
	} else {
		goto L109
	}
L79:
	;
	if v273 != 0 {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	goto L81
L81:
	;
	if v273 != 0 {
		v335 = v259
		goto L92
	} else {
		goto L93
	}
L82:
	;
	v374 = v258
	v375 = v259
	v376 = v257
	goto L77
L83:
	;
	goto L84
L84:
	;
	if v257&int32(3) == int32(0) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v352 = v258
	v353 = v259
	v354 = v257
	goto L78
L86:
	;
	goto L87
L87:
	;
	v280 = v258
	v281 = v259
	v282 = v257
	goto L88
L88:
	;
	if v281 == int32(0) {
		goto L72
	} else {
		goto L90
	}
L89:
	;
	v352 = v289
	v353 = v291
	v354 = v293
	goto L78
L90:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280))))
	*(*uint8)(unsafe.Add(mBase, uint32(v282))) = uint8(v286)
	v288 = int32(1)
	v289 = v280 + v288
	v291 = v281 - v288
	v293 = v282 + v288
	if v293&int32(3) != 0 {
		v280 = v289
		v281 = v291
		v282 = v293
		goto L88
	} else {
		goto L91
	}
L91:
	;
	goto L89
L92:
	;
	if v335 == int32(0) {
		goto L72
	} else {
		goto L105
	}
L93:
	;
	if v263&int32(3) != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v300 = v259
	goto L97
L95:
	;
	v315 = v259
	goto L96
L96:
	;
	if base.Ui32(v315) <= base.Ui32(int32(3)) {
		v335 = v315
		goto L92
	} else {
		goto L101
	}
L97:
	;
	if v300 == int32(0) {
		goto L72
	} else {
		goto L99
	}
L98:
	;
	v315 = v306
	goto L96
L99:
	;
	v306 = v300 - int32(1)
	v307 = v257 + v306
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258+v306))))
	*(*uint8)(unsafe.Add(mBase, uint32(v307))) = uint8(v309)
	if v307&int32(3) != 0 {
		v300 = v306
		goto L97
	} else {
		goto L100
	}
L100:
	;
	goto L98
L101:
	;
	v322 = v315
	goto L102
L102:
	;
	v326 = v322 - int32(4)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v258+v326)))
	*(*int32)(unsafe.Add(mBase, uint32(v257+v326))) = v329
	if base.Ui32(int32(3)) < base.Ui32(v326) {
		v322 = v326
		goto L102
	} else {
		goto L104
	}
L103:
	;
	v335 = v326
	goto L92
L104:
	;
	goto L103
L105:
	;
	v342 = v335
	goto L106
L106:
	;
	v346 = v342 - int32(1)
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258+v346))))
	*(*uint8)(unsafe.Add(mBase, uint32(v257+v346))) = uint8(v349)
	if v346 != 0 {
		v342 = v346
		goto L106
	} else {
		goto L108
	}
L107:
	;
	goto L72
L108:
	;
	goto L107
L109:
	;
	v359 = v352
	v360 = v353
	v361 = v354
	goto L110
L110:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v359)))
	*(*int32)(unsafe.Add(mBase, uint32(v361))) = v363
	v365 = int32(4)
	v366 = v359 + v365
	v368 = v361 + v365
	v370 = v360 - v365
	if base.Ui32(int32(3)) < base.Ui32(v370) {
		v359 = v366
		v360 = v370
		v361 = v368
		goto L110
	} else {
		goto L112
	}
L111:
	;
	v374 = v366
	v375 = v370
	v376 = v368
	goto L77
L112:
	;
	goto L111
L113:
	;
	v381 = v374
	v382 = v375
	v383 = v376
	goto L114
L114:
	;
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381))))
	*(*uint8)(unsafe.Add(mBase, uint32(v383))) = uint8(v385)
	v387 = int32(1)
	v392 = v382 - v387
	if v392 != 0 {
		v381 = v381 + v387
		v382 = v392
		v383 = v383 + v387
		goto L114
	} else {
		goto L116
	}
L115:
	;
	goto L72
L116:
	;
	goto L115
L117:
	;
	if l1 <= v537 {
		goto L161
	} else {
		goto L162
	}
L118:
	;
	v414 = int32(base.Ui32(v404+int32(_a_F_compactify_tuples_2))>>(uint(int32(4))%32)) & int32(_a_F_compactify_tuples_3)
	goto L120
L119:
	;
	v414 = int32(0)
	goto L120
L120:
	;
	if l1 < v414 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v416 = int32(1)
	if l1 <= v416 {
		goto L124
	} else {
		goto L125
	}
L122:
	;
	goto L123
L123:
	;
	v495 = int32(1)
	if base.Ui32(l1) <= base.Ui32(v495) {
		goto L148
	} else {
		goto L149
	}
L124:
	;
	v419 = v416
	goto L126
L125:
	;
	v419 = l1
	goto L126
L126:
	;
	v422 = int32(0)
	if int32(2) <= l1 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v431 = v422
	v433 = int32(0)
	goto L130
L128:
	;
	v469 = v422
	goto L129
L129:
	;
	if v419&int32(1) != 0 {
		goto L141
	} else {
		goto L142
	}
L130:
	;
	v443 = l0 + v431*int32(6)
	v444 = int32(*(*int16)(unsafe.Add(mBase, uint32(v443)+2)))
	v447 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v443)+4)))
	if v447 != 0 {
		goto L133
	} else {
		goto L134
	}
L131:
	;
	v469 = v462
	goto L129
L132:
	;
	v454 = l0 + (v431|int32(1))*int32(6)
	v455 = int32(*(*int16)(unsafe.Add(mBase, uint32(v454)+2)))
	v458 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v454)+4)))
	if v458 != 0 {
		goto L137
	} else {
		goto L138
	}
L133:
	;
	v448 = F__emscripten_memcpy_bulkmem(m, v16+v444, l2+v444, v447)
	mBase = m.M
	goto L135
L134:
	;
	goto L135
L135:
	;
	goto L132
L136:
	;
	v461 = int32(2)
	v462 = v431 + v461
	v464 = v433 + v461
	if v464 != v419&int32(2147483646) {
		v431 = v462
		v433 = v464
		goto L130
	} else {
		goto L140
	}
L137:
	;
	v459 = F__emscripten_memcpy_bulkmem(m, v16+v455, l2+v455, v458)
	mBase = m.M
	goto L139
L138:
	;
	goto L139
L139:
	;
	goto L136
L140:
	;
	goto L131
L141:
	;
	v481 = l0 + v469*int32(6)
	v482 = int32(*(*int16)(unsafe.Add(mBase, uint32(v481)+2)))
	v485 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v481)+4)))
	if v485 != 0 {
		goto L145
	} else {
		goto L146
	}
L142:
	;
	goto L143
L143:
	;
	v490 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v491 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
	v493 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+16)))
	v535 = v493
	v537 = int32(0)
	v541 = v490 + v491
	goto L117
L144:
	;
	goto L143
L145:
	;
	v486 = F__emscripten_memcpy_bulkmem(m, v16+v482, l2+v482, v485)
	mBase = m.M
	goto L147
L146:
	;
	goto L147
L147:
	;
	goto L144
L148:
	;
	v498 = v495
	goto L150
L149:
	;
	v498 = l1
	goto L150
L150:
	;
	v499 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+16)))
	v503 = v499
	v505 = v5
	goto L152
L151:
	;
	v526 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+14)))
	v529 = v524 - v526
	if v529 != 0 {
		goto L157
	} else {
		goto L158
	}
L152:
	;
	v515 = l0 + v505*int32(6)
	v516 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v515)+4)))
	v517 = int32(*(*int16)(unsafe.Add(mBase, uint32(v515)+2)))
	v518 = v516 + v517
	if v503 != v518 {
		v524 = v503
		v525 = v505
		goto L151
	} else {
		goto L154
	}
L153:
	;
	v524 = v520
	v525 = v498
	goto L151
L154:
	;
	v520 = v503 - v516
	v522 = v505 + int32(1)
	if v522 != v498 {
		v503 = v520
		v505 = v522
		goto L152
	} else {
		goto L155
	}
L155:
	;
	goto L153
L156:
	;
	v535 = v524
	v537 = v525
	v541 = v518
	goto L117
L157:
	;
	v530 = F__emscripten_memcpy_bulkmem(m, v16+v526, l2+v526, v529)
	mBase = m.M
	goto L159
L158:
	;
	goto L159
L159:
	;
	goto L156
L160:
	;
	v615 = v609 - v604
	if v615 != 0 {
		goto L176
	} else {
		goto L177
	}
L161:
	;
	v603 = v535
	v604 = v541
	v609 = v541
	goto L160
L162:
	;
	goto L163
L163:
	;
	v551 = v535
	v552 = v541
	v553 = v537
	v557 = v541
	goto L164
L164:
	;
	v563 = l0 + v553*int32(6)
	v564 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v563))))
	v573 = (v564+int32(1))&int32(_a_F_compactify_tuples_0)<<(uint(int32(2))%32) + (l2 + int32(24)) - int32(4)
	v574 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v563)+4)))
	v575 = int32(*(*int16)(unsafe.Add(mBase, uint32(v563)+2)))
	if v574+v575 == v552 {
		goto L167
	} else {
		goto L168
	}
L165:
	;
	v603 = v592
	v604 = v586
	v609 = v588
	goto L160
L166:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v573)))
	v592 = v551 - v587
	*(*int32)(unsafe.Add(mBase, uint32(v573))) = v589&int32(-32768) | v592&int32(_a_F_compactify_tuples_1)
	v598 = v553 + int32(1)
	if v598 != l1 {
		v551 = v592
		v552 = v586
		v553 = v598
		v557 = v588
		goto L164
	} else {
		goto L174
	}
L167:
	;
	v586 = v575
	v587 = v574
	v588 = v557
	goto L166
L168:
	;
	goto L169
L169:
	;
	v580 = v557 - v552
	if v580 != 0 {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	v583 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v563)+4)))
	v584 = int32(*(*int16)(unsafe.Add(mBase, uint32(v563)+2)))
	v586 = v584
	v587 = v583
	v588 = v583 + v584
	goto L166
L171:
	;
	v581 = F__emscripten_memcpy_bulkmem(m, l2+v551, v552+v16, v580)
	mBase = m.M
	goto L173
L172:
	;
	goto L173
L173:
	;
	goto L170
L174:
	;
	goto L165
L175:
	;
	v621 = v603
	goto L1
L176:
	;
	v616 = F__emscripten_memcpy_bulkmem(m, l2+v603, v604+v16, v615)
	mBase = m.M
	goto L178
L177:
	;
	goto L178
L178:
	;
	goto L175
}
func F_compareDoubles(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 float64
	_ = v7
	var v8 float64
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v7 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	if base.F64_gt(v7, v8) != 0 {
		v10 = int32(1)
	} else {
		v10 = int32(-1)
	}
	if base.F64_ne(v7, v8) != 0 {
		v13 = v10
	} else {
		v13 = int32(0)
	}
	return v13
}
func F_compareentry(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = int32(12)
	v8 = int32(1)
	v10 = int32(2047)
	v11 = int32(base.Ui32(v4)>>(uint(v8)%32)) & v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v19 = int32(base.Ui32(v12)>>(uint(v8)%32)) & v10
	if v11 == int32(0) {
		v25 = int32(0)
		if v25 < v19 {
			v28 = int32(-1)
		} else {
			v28 = v25
		}
		v45 = v28
	} else {
		if v19 == int32(0) {
			v45 = base.B2i32(int32(0) < v11)
		} else {
			if base.Ui32(v11) < base.Ui32(v19) {
				v34 = v11
			} else {
				v34 = v19
			}
			v35 = F_memcmp(m, l2+int32(base.Ui32(v4)>>(uint(v5)%32)), l2+int32(base.Ui32(v12)>>(uint(v5)%32)), v34)
			mBase = m.M
			if v35 != 0 {
				v43 = v35
				v45 = v43
			} else {
				if v11 == v19 {
					v45 = int32(0)
				} else {
					if v11 < v19 {
						v42 = int32(-1)
					} else {
						v42 = int32(1)
					}
					v43 = v42
					v45 = v43
				}
			}
		}
	}
	return v45
}
func F_comparison_shim(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	v4 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = l0
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+44)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l1
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v18 = m.T0[v17].(func(*base.Module, int32) int32)(m, v9+int32(28))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+44)))
		if v22 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v29
				F_errmsg_internal(m, int32(_a_F_comparison_shim_0), v7)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_comparison_shim_1), int32(58), int32(_a_F_comparison_shim_2))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			m.G0 = v7 + int32(16)
			return v18
		}
	}
}
func F_compress_init(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
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
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+64)))
	if base.Ui32(v6-int32(3)) < base.Ui32(int32(-2)) {
		return int32(-102)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
		v16 = m.Env.Pgmem_deflate_create(m, base.B2i32(v6 == int32(1)), v15)
		mBase = m.M
		if v16 <= int32(0) {
			return int32(-105)
		} else {
			v22 = F_palloc0(m, int32(_a_F_compress_init_0))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = int32(_a_F_compress_init_1)
				v28 = int32(_a_F_compress_init_2)
				v29 = *(*int32)(unsafe.Add(mBase, _c_F_compress_init[0]))
				F_ResourceOwnerEnlarge(m, v29)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, _c_F_compress_init[1]))
					v35 = F_MemoryContextAlloc(m, v33, int32(8))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v35))) = v16
						v38 = *(*int32)(unsafe.Add(mBase, _c_F_compress_init[0]))
						*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v38
						F_ResourceOwnerRemember(m, v38, v35, int32(_a_F_compress_init_3))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v22))) = v35
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v22
							return int32(_a_F_compress_init_1)
						}
					}
				}
			}
		}
	}
}
func F_compute_gather_rows(m *base.Module, l0 int32) float64 {
	mBase := m.M
	_ = mBase
	var v5 float64
	_ = v5
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v9 int32
	_ = v9
	var v15 float64
	_ = v15
	var v19 float64
	_ = v19
	var v21 float64
	_ = v21
	var v23 float64
	_ = v23
	var v24 float64
	_ = v24
	var v32 float64
	_ = v32
	var v36 float64
	_ = v36
	v5 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = base.F64_convert_i32_s(v6)
	v9 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_compute_gather_rows[0])))
	if v9 == int32(1) {
		v15 = base.F64_add(base.F64_mul(v7, float64(-0.3)), float64(1))
		if base.F64_gt(v15, float64(0)) != 0 {
			v19 = v15
		} else {
			v19 = math.Float64frombits(uint64(0x8000000000000000))
		}
		v21 = base.F64_add(v19, v7)
	} else {
		v21 = v7
	}
	v23 = float64(1e+100)
	v24 = base.F64_mul(v5, v21)
	if base.F64_gt(v24, v23) != 0 {
		v36 = v23
	} else {
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v24)&int64(9223372036854775807)) {
			v36 = v23
		} else {
			v32 = float64(1)
			if base.F64_le(v24, v32) != 0 {
				v36 = v32
			} else {
				v36 = base.F64_nearest(v24)
			}
		}
	}
	return v36
}
func F_connectby_text(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
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
	var v27 int32
	_ = v27
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
	var v33 int32
	_ = v33
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
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
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
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
	var v73 int32
	_ = v73
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
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v21 = F_pg_detoast_datum_packed(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		v25 = F_text_to_cstring(m, v21)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v28 = F_pg_detoast_datum_packed(m, v27)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				v30 = F_text_to_cstring(m, v28)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					v33 = F_pg_detoast_datum_packed(m, v32)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						v35 = F_text_to_cstring(m, v33)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
							v38 = F_pg_detoast_datum_packed(m, v37)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								v40 = F_text_to_cstring(m, v38)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int32(0)
								} else {
									v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									if v42 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v129 = m.ExcPending
										if v129 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(1088))
											mBase = m.M
											v132 = m.ExcPending
											if v132 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(_a_F_connectby_text_0), int32(0))
												mBase = m.M
												v136 = m.ExcPending
												if v136 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_connectby_text_1), int32(997), int32(_a_F_connectby_text_2))
													mBase = m.M
													v141 = m.ExcPending
													if v141 != 0 {
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
										v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
										if v45 != int32(383) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v129 = m.ExcPending
											if v129 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(1088))
												mBase = m.M
												v132 = m.ExcPending
												if v132 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(_a_F_connectby_text_0), int32(0))
													mBase = m.M
													v136 = m.ExcPending
													if v136 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_connectby_text_1), int32(997), int32(_a_F_connectby_text_2))
														mBase = m.M
														v141 = m.ExcPending
														if v141 != 0 {
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
											v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+12)))
											if v48&int32(2) == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v145 = m.ExcPending
												if v145 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(1088))
													mBase = m.M
													v148 = m.ExcPending
													if v148 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(_a_F_connectby_text_3), int32(0))
														mBase = m.M
														v152 = m.ExcPending
														if v152 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_connectby_text_1), int32(1002), int32(_a_F_connectby_text_2))
															mBase = m.M
															v157 = m.ExcPending
															if v157 != 0 {
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
												v53 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
												if v53 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v145 = m.ExcPending
													if v145 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(1088))
														mBase = m.M
														v148 = m.ExcPending
														if v148 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(_a_F_connectby_text_3), int32(0))
															mBase = m.M
															v152 = m.ExcPending
															if v152 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_connectby_text_1), int32(1002), int32(_a_F_connectby_text_2))
																mBase = m.M
																v157 = m.ExcPending
																if v157 != 0 {
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
													v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
													v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
													if v57 == int32(6) {
														v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
														v61 = F_pg_detoast_datum_packed(m, v60)
														mBase = m.M
														v62 = m.ExcPending
														if v62 != 0 {
															return int32(0)
														} else {
															v63 = F_text_to_cstring(m, v61)
															mBase = m.M
															v64 = m.ExcPending
															if v64 != 0 {
																return int32(0)
															} else {
																v68 = v63
																v69 = int32(_a_F_connectby_text_4)
																v70 = *(*int32)(unsafe.Add(mBase, _c_F_connectby_text[0]))
																v72 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
																v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+16))
																*(*int32)(unsafe.Add(mBase, _c_F_connectby_text[0])) = v73
																v75 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
																v76 = F_CreateTupleDescCopy(m, v75)
																mBase = m.M
																v77 = m.ExcPending
																if v77 != 0 {
																	return int32(0)
																} else {
																	v79 = base.B2i32(v57 == int32(6))
																	F_validateConnectbyTupleDesc(m, v76, v79, int32(0))
																	mBase = m.M
																	v82 = m.ExcPending
																	if v82 != 0 {
																		return int32(0)
																	} else {
																		v83 = F_TupleDescGetAttInMetadata(m, v76)
																		mBase = m.M
																		v84 = m.ExcPending
																		if v84 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v42)+16)) = int32(2)
																			v87 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
																			*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = int32(1)
																			F_SPI_connect_ext(m, int32(0))
																			mBase = m.M
																			v92 = m.ExcPending
																			if v92 != 0 {
																				return int32(0)
																			} else {
																				v93 = int32(_a_F_connectby_text_4)
																				v94 = *(*int32)(unsafe.Add(mBase, _c_F_connectby_text[0]))
																				*(*int32)(unsafe.Add(mBase, _c_F_connectby_text[0])) = v73
																				v103 = *(*int32)(unsafe.Add(mBase, _c_F_connectby_text[1]))
																				v104 = F_tuplestore_begin_heap(m, int32(base.Ui32(v87&int32(4))>>(uint(int32(2))%32)), int32(0), v103)
																				mBase = m.M
																				v105 = m.ExcPending
																				if v105 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, _c_F_connectby_text[0])) = v94
																					v108 = int32(0)
																					F_build_tuplestore_recursively(m, v30, v35, v25, v108, v68, v40, v40, v108, v18+int32(12), v56, v79, v108, v83, v104)
																					mBase = m.M
																					v114 = m.ExcPending
																					if v114 != 0 {
																						return int32(0)
																					} else {
																						v115 = F_SPI_finish(m)
																						mBase = m.M
																						v116 = m.ExcPending
																						if v116 != 0 {
																							return int32(0)
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(v42)+28)) = v76
																							*(*int32)(unsafe.Add(mBase, uint32(v42)+24)) = v104
																							*(*int32)(unsafe.Add(mBase, _c_F_connectby_text[0])) = v70
																							m.G0 = v18 + int32(16)
																							return int32(0)
																						}
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
													} else {
														v66 = F_pstrdup(m, int32(_a_F_connectby_text_5))
														mBase = m.M
														v67 = m.ExcPending
														if v67 != 0 {
															return int32(0)
														} else {
															v68 = v66
															v69 = int32(_a_F_connectby_text_4)
															v70 = *(*int32)(unsafe.Add(mBase, _c_F_connectby_text[0]))
															v72 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
															v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+16))
															*(*int32)(unsafe.Add(mBase, _c_F_connectby_text[0])) = v73
															v75 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
															v76 = F_CreateTupleDescCopy(m, v75)
															mBase = m.M
															v77 = m.ExcPending
															if v77 != 0 {
																return int32(0)
															} else {
																v79 = base.B2i32(v57 == int32(6))
																F_validateConnectbyTupleDesc(m, v76, v79, int32(0))
																mBase = m.M
																v82 = m.ExcPending
																if v82 != 0 {
																	return int32(0)
																} else {
																	v83 = F_TupleDescGetAttInMetadata(m, v76)
																	mBase = m.M
																	v84 = m.ExcPending
																	if v84 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v42)+16)) = int32(2)
																		v87 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
																		*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = int32(1)
																		F_SPI_connect_ext(m, int32(0))
																		mBase = m.M
																		v92 = m.ExcPending
																		if v92 != 0 {
																			return int32(0)
																		} else {
																			v93 = int32(_a_F_connectby_text_4)
																			v94 = *(*int32)(unsafe.Add(mBase, _c_F_connectby_text[0]))
																			*(*int32)(unsafe.Add(mBase, _c_F_connectby_text[0])) = v73
																			v103 = *(*int32)(unsafe.Add(mBase, _c_F_connectby_text[1]))
																			v104 = F_tuplestore_begin_heap(m, int32(base.Ui32(v87&int32(4))>>(uint(int32(2))%32)), int32(0), v103)
																			mBase = m.M
																			v105 = m.ExcPending
																			if v105 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, _c_F_connectby_text[0])) = v94
																				v108 = int32(0)
																				F_build_tuplestore_recursively(m, v30, v35, v25, v108, v68, v40, v40, v108, v18+int32(12), v56, v79, v108, v83, v104)
																				mBase = m.M
																				v114 = m.ExcPending
																				if v114 != 0 {
																					return int32(0)
																				} else {
																					v115 = F_SPI_finish(m)
																					mBase = m.M
																					v116 = m.ExcPending
																					if v116 != 0 {
																						return int32(0)
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v42)+28)) = v76
																						*(*int32)(unsafe.Add(mBase, uint32(v42)+24)) = v104
																						*(*int32)(unsafe.Add(mBase, _c_F_connectby_text[0])) = v70
																						m.G0 = v18 + int32(16)
																						return int32(0)
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
				}
			}
		}
	}
}
func F_consider_groupingsets_paths(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 float64) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v31 float64
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 float64
	_ = v40
	var v42 int32
	_ = v42
	var v46 float64
	_ = v46
	var v47 float64
	_ = v47
	var v50 float64
	_ = v50
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 float64
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 float64
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
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
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v232 int32
	_ = v232
	var v254 int32
	_ = v254
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v369 int32
	_ = v369
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v417 int32
	_ = v417
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v543 int32
	_ = v543
	var v548 int32
	_ = v548
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v590 int32
	_ = v590
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v640 int32
	_ = v640
	var v658 float64
	_ = v658
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v711 float64
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v722 float64
	_ = v722
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v736 float64
	_ = v736
	var v737 float64
	_ = v737
	var v740 float64
	_ = v740
	var v742 float64
	_ = v742
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v763 float64
	_ = v763
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v796 int32
	_ = v796
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v807 float64
	_ = v807
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v819 float64
	_ = v819
	var v821 float64
	_ = v821
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v846 int32
	_ = v846
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v885 int32
	_ = v885
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v900 int32
	_ = v900
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v993 int32
	_ = v993
	var v1011 int32
	_ = v1011
	var v1017 int32
	_ = v1017
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1048 float64
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1052 int32
	_ = v1052
	var v1053 float64
	_ = v1053
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1070 int32
	_ = v1070
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1082 int32
	_ = v1082
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1107 int32
	_ = v1107
	var v1132 int32
	_ = v1132
	var v1135 int32
	_ = v1135
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1174 int32
	_ = v1174
	var v1187 int32
	_ = v1187
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1211 float64
	_ = v1211
	var v1248 int32
	_ = v1248
	var v1283 int32
	_ = v1283
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1326 int32
	_ = v1326
	var v1328 int32
	_ = v1328
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1336 int32
	_ = v1336
	var v1340 int32
	_ = v1340
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1345 int32
	_ = v1345
	var v1353 int32
	_ = v1353
	var v1357 int32
	_ = v1357
	var v1363 int32
	_ = v1363
	var v1365 int32
	_ = v1365
	var v1381 int32
	_ = v1381
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1411 int32
	_ = v1411
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1423 int32
	_ = v1423
	var v1425 int32
	_ = v1425
	var v1428 int32
	_ = v1428
	var v1432 int32
	_ = v1432
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1445 int32
	_ = v1445
	var v1448 int32
	_ = v1448
	var v1449 int32
	_ = v1449
	var v1466 int32
	_ = v1466
	var v1473 int32
	_ = v1473
	var v1484 int32
	_ = v1484
	var v1488 int32
	_ = v1488
	var v1490 int32
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1494 int32
	_ = v1494
	var v1495 int32
	_ = v1495
	var v1496 int32
	_ = v1496
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1506 int32
	_ = v1506
	var v1507 int32
	_ = v1507
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1522 int32
	_ = v1522
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1557 int32
	_ = v1557
	var v1558 int32
	_ = v1558
	var v1595 int32
	_ = v1595
	var v1597 int32
	_ = v1597
	var v1611 int32
	_ = v1611
	var v1615 int32
	_ = v1615
	var v1632 int32
	_ = v1632
	var v1633 int32
	_ = v1633
	var v1637 int32
	_ = v1637
	var v1638 int32
	_ = v1638
	var v1641 int32
	_ = v1641
	var v1642 int32
	_ = v1642
	var v1648 int32
	_ = v1648
	var v1653 int32
	_ = v1653
	var v1677 int32
	_ = v1677
	var v1678 int32
	_ = v1678
	var v1681 int32
	_ = v1681
	var v1685 int32
	_ = v1685
	var v1686 int32
	_ = v1686
	var v1687 int32
	_ = v1687
	var v1689 int32
	_ = v1689
	var v1690 int32
	_ = v1690
	var v1695 int32
	_ = v1695
	var v1724 int32
	_ = v1724
	var v1725 int32
	_ = v1725
	var v1727 int32
	_ = v1727
	var v1728 int32
	_ = v1728
	var v1745 int32
	_ = v1745
	var v1763 float64
	_ = v1763
	var v1764 int32
	_ = v1764
	var v1767 int32
	_ = v1767
	var v1768 int32
	_ = v1768
	var v1770 int32
	_ = v1770
	var v1771 int32
	_ = v1771
	var v1787 int32
	_ = v1787
	var v1807 int32
	_ = v1807
	var v1809 int32
	_ = v1809
	var v1810 int32
	_ = v1810
	var v1812 int32
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1814 int32
	_ = v1814
	var v1815 int32
	_ = v1815
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1832 int32
	_ = v1832
	var v1836 int32
	_ = v1836
	var v1845 int32
	_ = v1845
	var v1846 int32
	_ = v1846
	var v1848 int32
	_ = v1848
	var v1849 int32
	_ = v1849
	var v1850 int32
	_ = v1850
	var v1867 int32
	_ = v1867
	var v1871 int32
	_ = v1871
	var v1873 int32
	_ = v1873
	var v1891 int32
	_ = v1891
	var v1892 int32
	_ = v1892
	var v1897 int32
	_ = v1897
	var v1898 int32
	_ = v1898
	var v1903 int32
	_ = v1903
	var v1905 int32
	_ = v1905
	var v1906 int32
	_ = v1906
	var v1907 int32
	_ = v1907
	var v1908 int32
	_ = v1908
	var v1910 int32
	_ = v1910
	var v1911 int32
	_ = v1911
	var v1912 int32
	_ = v1912
	var v1932 int32
	_ = v1932
	var v1945 int32
	_ = v1945
	var v1947 int32
	_ = v1947
	var v1967 int32
	_ = v1967
	v9 = int32(0)
	v31 = float64(0)
	v33 = m.G0
	v35 = v33 - int32(32)
	m.G0 = v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v40 = *(*float64)(unsafe.Add(mBase, _c_F_consider_groupingsets_paths[0]))
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_consider_groupingsets_paths[1]))
	v46 = base.F64_mul(base.F64_mul(v40, base.F64_convert_i32_s(v42)), float64(1024))
	v47 = float64(4.294967295e+09)
	if base.F64_lt(v46, v47) != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if l3 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L2:
	;
	v50 = v46
	goto L4
L3:
	;
	v50 = v47
	goto L4
L4:
	;
	if base.F64_lt(v50, float64(4.294967296e+09))&base.F64_ge(v50, float64(0)) != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v56 = base.I32_trunc_f64_u(v50)
	v58 = v56
	goto L1
L6:
	;
	goto L7
L7:
	;
	v58 = int32(0)
	goto L1
L8:
	;
	m.G0 = v1967 + int32(32)
	return
L9:
	;
	F_add_path(m, l1, v1945)
	mBase = m.M
	v1947 = m.ExcPending
	if v1947 != 0 {
		goto L47
	} else {
		goto L273
	}
L10:
	;
	if v1867 == int32(0) {
		v1967 = v35
		goto L8
	} else {
		goto L262
	}
L11:
	;
	v62 = int32(0)
	if v59 == v62 {
		v142 = v9
		v143 = v31
		v144 = v62
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	if v59 == int32(0) {
		v1967 = v35
		goto L8
	} else {
		goto L106
	}
L14:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	if v146 != 0 {
		goto L40
	} else {
		goto L41
	}
L15:
	;
	v65 = int32(0)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	if v66 == v65 {
		v142 = v9
		v143 = v31
		v144 = v65
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l2)+64))
	if v69 == v70 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v123 == int32(0) {
		v142 = v9
		v143 = v31
		v144 = v66
		goto L14
	} else {
		goto L35
	}
L18:
	;
	v123 = int32(1)
	goto L17
L19:
	;
	goto L20
L20:
	;
	v79 = int32(0)
	goto L22
L21:
	;
	v123 = v115
	goto L17
L22:
	;
	v83 = int32(0)
	if v69 == v83 {
		v93 = v83
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v115 = int32(0)
	goto L21
L24:
	;
	if v70 != 0 {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if v87 <= v79 {
		v93 = int32(0)
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	v93 = v89 + v79<<(uint(int32(2))%32)
	goto L24
L27:
	;
	v99 = base.B2i32(v93 == int32(0))
	if v93 == int32(0) {
		v115 = v99
		goto L21
	} else {
		goto L32
	}
L28:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	if v79 < v94 {
		goto L27
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v123 = base.B2i32(v93 == int32(0))
	goto L17
L31:
	;
	goto L30
L32:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
	v105 = v102 + v79<<(uint(int32(2))%32)
	if v105 == int32(0) {
		v115 = v99
		goto L21
	} else {
		goto L33
	}
L33:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	if v110 == v111 {
		v79 = v79 + int32(1)
		goto L22
	} else {
		goto L34
	}
L34:
	;
	goto L23
L35:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v127 = *(*float64)(unsafe.Add(mBase, uint32(v126)+16))
	v129 = v66 + int32(4)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+12))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	if base.Ui32(v129) < base.Ui32(v132+v133<<(uint(int32(2))%32)) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v138 = v129
	goto L38
L37:
	;
	v138 = int32(0)
	goto L38
L38:
	;
	v142 = v126
	v143 = v127
	v144 = v138
	goto L14
L39:
	;
	if base.F64_gt(base.F64_mul(base.F64_sub(l7, v143), base.F64_convert_i32_u(v153)), base.F64_convert_i32_u(v58)) != 0 {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)+4))
	v149 = v147
	goto L42
L41:
	;
	v149 = int32(0)
	goto L42
L42:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+32))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l6)+32))
	v153 = F_hash_agg_entry_size(m, v149, v151, v152)
	mBase = m.M
	goto L39
L43:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v158 != 0 {
		v1967 = v35
		goto L8
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l5)+28))
	v160 = F_list_copy(m, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L45
L47:
	;
	return
L48:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v144 != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	if v232 == int32(0) {
		v1967 = v35
		goto L8
	} else {
		goto L61
	}
L50:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v162)+4))
	if v171 <= v170 {
		v232 = v160
		goto L49
	} else {
		goto L55
	}
L51:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)+12))
	v170 = (v144 - v163) >> (uint(int32(2)) % 32)
	goto L50
L52:
	;
	goto L53
L53:
	;
	if v162 == int32(0) {
		v232 = v160
		goto L49
	} else {
		goto L54
	}
L54:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v162)+4))
	v170 = v169
	goto L50
L55:
	;
	v181 = v170
	v185 = v160
	goto L56
L56:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v162)+12))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v205+v181<<(uint(int32(2))%32))))
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+24)))
	if v210 != int32(1) {
		v1967 = v35
		goto L8
	} else {
		goto L58
	}
L57:
	;
	v232 = v214
	goto L49
L58:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v209)+12))
	v214 = F_list_concat(m, v185, v213)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L47
	} else {
		goto L59
	}
L59:
	;
	v217 = v181 + int32(1)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v162)+4))
	if v217 < v218 {
		v181 = v217
		v185 = v214
		goto L56
	} else {
		goto L60
	}
L60:
	;
	goto L57
L61:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v232)+4))
	if v254 <= int32(0) {
		v1867 = v9
		v1871 = v9
		v1873 = v9
		goto L10
	} else {
		goto L62
	}
L62:
	;
	v271 = v9
	v273 = v9
	v277 = v9
	v279 = v9
	goto L63
L63:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v232)+12))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v289+v271<<(uint(int32(2))%32))))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v293)+4))
	if v294 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	v1867 = v680
	v1871 = v684
	v1873 = v686
	goto L10
L65:
	;
	v697 = v271 + int32(1)
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v232)+4))
	if v697 < v698 {
		v271 = v697
		v273 = v680
		v277 = v684
		v279 = v686
		goto L63
	} else {
		goto L105
	}
L66:
	;
	v297 = F_lappend(m, v279, v293)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L47
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v303 = F_palloc0(m, int32(32))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L47
	} else {
		goto L71
	}
L69:
	;
	v300 = F_lappend(m, v277, int32(0))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L47
	} else {
		goto L70
	}
L70:
	;
	v680 = v273
	v684 = v300
	v686 = v297
	goto L65
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v303))) = int32(309)
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v294)+4))
	if v307 <= int32(0) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v303)+4)) = v369
	*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v293
	*(*int32)(unsafe.Add(mBase, uint32(v35)+28)) = v293
	v398 = F_list_make1_impl(m, int32(1), v35+int32(16))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L47
	} else {
		goto L81
	}
L73:
	;
	v369 = int32(0)
	goto L72
L74:
	;
	goto L75
L75:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v312 = int32(0)
	v322 = v312
	v323 = v312
	goto L76
L76:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v294)+12))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v346+v322<<(uint(int32(2))%32))))
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v311)+100))
	v352 = F_get_sortgroupref_clause(m, v350, v351)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L47
	} else {
		goto L78
	}
L77:
	;
	v369 = v354
	goto L72
L78:
	;
	v354 = F_lappend(m, v323, v352)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L47
	} else {
		goto L79
	}
L79:
	;
	v357 = v322 + int32(1)
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v294)+4))
	if v357 < v358 {
		v322 = v357
		v323 = v354
		goto L76
	} else {
		goto L80
	}
L80:
	;
	goto L77
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v303)+12)) = v398
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l5)+32))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v303)+4))
	if v402 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	if v398 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L83:
	;
	v405 = int32(0)
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v402)+4))
	if v406 <= v405 {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v417 = v405
	goto L85
L85:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v402)+12))
	v442 = int32(2)
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v441+v417<<(uint(v442)%32))))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v445)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v401+v446<<(uint(v442)%32)))) = v417
	v452 = v417 + int32(1)
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v402)+4))
	if v452 < v453 {
		v417 = v452
		goto L85
	} else {
		goto L87
	}
L86:
	;
	goto L82
L87:
	;
	goto L86
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v303)+8)) = v640
	v658 = *(*float64)(unsafe.Add(mBase, uint32(v293)+8))
	v659 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v303)+24)) = uint16(v659)
	*(*float64)(unsafe.Add(mBase, uint32(v303)+16)) = v658
	v662 = F_lappend(m, v273, v303)
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L47
	} else {
		goto L104
	}
L89:
	;
	v640 = int32(0)
	goto L88
L90:
	;
	goto L91
L91:
	;
	v490 = int32(0)
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v398)+4))
	if v492 <= v490 {
		v640 = v490
		goto L88
	} else {
		goto L92
	}
L92:
	;
	v506 = v490
	v510 = v490
	goto L93
L93:
	;
	v527 = int32(0)
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v398)+12))
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v528+v506<<(uint(int32(2))%32))))
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v532)+4))
	if v533 == v527 {
		v590 = v527
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v640 = v619
	goto L88
L95:
	;
	v619 = F_lappend(m, v510, v590)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L47
	} else {
		goto L102
	}
L96:
	;
	v536 = int32(0)
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v533)+4))
	if v537 <= v536 {
		v590 = v527
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v543 = v527
	v548 = v536
	goto L98
L98:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v533)+12))
	v573 = int32(2)
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v572+v548<<(uint(v573)%32))))
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v401+v576<<(uint(v573)%32))))
	v581 = F_lappend_int(m, v543, v580)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L47
	} else {
		goto L100
	}
L99:
	;
	v590 = v581
	goto L95
L100:
	;
	v584 = v548 + int32(1)
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v533)+4))
	if v584 < v585 {
		v543 = v581
		v548 = v584
		goto L98
	} else {
		goto L101
	}
L101:
	;
	goto L99
L102:
	;
	v622 = v506 + int32(1)
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v398)+4))
	if v622 < v623 {
		v506 = v622
		v510 = v619
		goto L93
	} else {
		goto L103
	}
L103:
	;
	goto L94
L104:
	;
	v680 = v662
	v684 = v277
	v686 = v279
	goto L65
L105:
	;
	goto L64
L106:
	;
	if l4 == int32(0) {
		v1813 = l0
		v1814 = l1
		v1815 = l2
		v1818 = l5
		v1819 = l6
		v1832 = v35
		v1836 = v37
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v1845 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+28))
	if v1845 != 0 {
		v1967 = v1832
		goto L8
	} else {
		goto L260
	}
L108:
	;
	v704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+16)))
	if v704 != int32(1) {
		v1813 = l0
		v1814 = l1
		v1815 = l2
		v1818 = l5
		v1819 = l6
		v1832 = v35
		v1836 = v37
		goto L107
	} else {
		goto L109
	}
L109:
	;
	v707 = *(*int32)(unsafe.Add(mBase, uint32(l5)+28))
	v708 = F_list_copy(m, v707)
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L47
	} else {
		goto L110
	}
L110:
	;
	v711 = *(*float64)(unsafe.Add(mBase, uint32(l5)+8))
	v712 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	if v712 != 0 {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	if v1423 != 0 {
		goto L222
	} else {
		goto L223
	}
L112:
	;
	v722 = base.F64_sub(base.F64_convert_i32_u(v58), base.F64_mul(v711, base.F64_convert_i32_u(v719)))
	if base.F64_gt(v722, float64(0)) == int32(0) {
		v1409 = l0
		v1410 = l1
		v1411 = l2
		v1414 = l5
		v1415 = l6
		v1423 = v9
		v1425 = v708
		v1428 = v35
		v1432 = v37
		goto L111
	} else {
		goto L116
	}
L113:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v712)+4))
	v715 = v713
	goto L115
L114:
	;
	v715 = int32(0)
	goto L115
L115:
	;
	v716 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v716)+32))
	v718 = *(*int32)(unsafe.Add(mBase, uint32(l6)+32))
	v719 = F_hash_agg_entry_size(m, v715, v717, v718)
	mBase = m.M
	goto L112
L116:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v727 == int32(0) {
		v1409 = l0
		v1410 = l1
		v1411 = l2
		v1414 = l5
		v1415 = l6
		v1423 = v9
		v1425 = v708
		v1428 = v35
		v1432 = v37
		goto L111
	} else {
		goto L117
	}
L117:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v727)+4))
	if v730 < int32(2) {
		v1409 = l0
		v1410 = l1
		v1411 = l2
		v1414 = l5
		v1415 = l6
		v1423 = v9
		v1425 = v708
		v1428 = v35
		v1432 = v37
		goto L111
	} else {
		goto L118
	}
L118:
	;
	v736 = base.F64_div(v722, base.F64_mul(base.F64_convert_i32_u(v730), float64(20)))
	v737 = float64(1)
	if base.F64_gt(v736, v737) != 0 {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v751 = F_palloc(m, v730<<(uint(int32(2))%32))
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L47
	} else {
		goto L126
	}
L120:
	;
	v740 = v736
	goto L122
L121:
	;
	v740 = v737
	goto L122
L122:
	;
	v742 = base.F64_floor(base.F64_div(v722, v740))
	if base.F64_lt(base.F64_abs(v742), float64(2.147483648e+09)) != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v746 = base.I32_trunc_f64_s(v742)
	v748 = v746
	goto L119
L124:
	;
	goto L125
L125:
	;
	v748 = int32(-2147483648)
	goto L119
L126:
	;
	v753 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v753 == int32(0) {
		v1409 = l0
		v1410 = l1
		v1411 = l2
		v1414 = l5
		v1415 = l6
		v1423 = v9
		v1425 = v708
		v1428 = v35
		v1432 = v37
		goto L111
	} else {
		goto L127
	}
L127:
	;
	v756 = int32(1)
	v757 = int32(0)
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v753)+4))
	if v756 < v758 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v763 = base.F64_add(base.F64_convert_i32_s(v748), float64(1))
	v772 = v756
	v773 = v757
	goto L131
L129:
	;
	v846 = v757
	goto L130
L130:
	;
	if v846 <= int32(0) {
		v1409 = l0
		v1410 = l1
		v1411 = l2
		v1414 = l5
		v1415 = l6
		v1423 = v9
		v1425 = v708
		v1428 = v35
		v1432 = v37
		goto L111
	} else {
		goto L148
	}
L131:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v753)+12))
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v796+v772<<(uint(int32(2))%32))))
	v801 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v800)+24)))
	if v801 == int32(1) {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	v846 = v831
	goto L130
L133:
	;
	v807 = *(*float64)(unsafe.Add(mBase, uint32(v800)+16))
	v808 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	if v808 != 0 {
		goto L138
	} else {
		goto L139
	}
L134:
	;
	v831 = v773
	goto L135
L135:
	;
	v834 = v772 + int32(1)
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v753)+4))
	if v834 < v835 {
		v772 = v834
		v773 = v831
		goto L131
	} else {
		goto L147
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v751+v773<<(uint(int32(2))%32)))) = v827
	v831 = v773 + int32(1)
	goto L135
L137:
	;
	v819 = base.F64_floor(base.F64_div(base.F64_mul(v807, base.F64_convert_i32_u(v815)), v740))
	if base.F64_gt(v763, v819) != 0 {
		goto L141
	} else {
		goto L142
	}
L138:
	;
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v808)+4))
	v811 = v809
	goto L140
L139:
	;
	v811 = int32(0)
	goto L140
L140:
	;
	v812 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v812)+32))
	v814 = *(*int32)(unsafe.Add(mBase, uint32(l6)+32))
	v815 = F_hash_agg_entry_size(m, v811, v813, v814)
	mBase = m.M
	goto L137
L141:
	;
	v821 = v819
	goto L143
L142:
	;
	v821 = v763
	goto L143
L143:
	;
	if base.F64_lt(base.F64_abs(v821), float64(2.147483648e+09)) != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v825 = base.I32_trunc_f64_s(v821)
	v827 = v825
	goto L136
L145:
	;
	goto L146
L146:
	;
	v827 = int32(-2147483648)
	goto L136
L147:
	;
	goto L132
L148:
	;
	v871 = int32(0)
	v873 = *(*int32)(unsafe.Add(mBase, _c_F_consider_groupingsets_paths[2]))
	v878 = F_AllocSetContextCreateInternal(m, v873, int32(_a_F_consider_groupingsets_paths_0), v871, int32(1024), int32(_a_F_consider_groupingsets_paths_1))
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L47
	} else {
		goto L149
	}
L149:
	;
	v880 = int32(_a_F_consider_groupingsets_paths_2)
	v881 = *(*int32)(unsafe.Add(mBase, _c_F_consider_groupingsets_paths[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_consider_groupingsets_paths[2])) = v878
	v885 = v748 + int32(1)
	v888 = F_palloc(m, v885<<(uint(int32(3))%32))
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L47
	} else {
		goto L150
	}
L150:
	;
	v892 = F_palloc(m, v885<<(uint(int32(2))%32))
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L47
	} else {
		goto L151
	}
L151:
	;
	if int32(0) <= v748 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v900 = v871
	goto L155
L153:
	;
	goto L154
L154:
	;
	if int32(0) < v846 {
		goto L159
	} else {
		goto L160
	}
L155:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v888+v900<<(uint(int32(3))%32)))) = int64(0)
	v936 = F_bms_make_singleton(m, v846)
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L47
	} else {
		goto L157
	}
L156:
	;
	goto L154
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v892+v900<<(uint(int32(2))%32)))) = v936
	v940 = v900 + int32(1)
	if v940 <= v748 {
		v900 = v940
		goto L155
	} else {
		goto L158
	}
L158:
	;
	goto L156
L159:
	;
	v993 = v9
	goto L162
L160:
	;
	goto L161
L161:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_consider_groupingsets_paths[2])) = v881
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(v892+v748<<(uint(int32(2))%32))))
	v1323 = F_bms_copy(m, v1322)
	mBase = m.M
	v1324 = m.ExcPending
	if v1324 != 0 {
		goto L47
	} else {
		goto L199
	}
L162:
	;
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(v751+v993<<(uint(int32(2))%32))))
	if v1011 <= v748 {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	goto L161
L164:
	;
	v1017 = v748
	goto L167
L165:
	;
	goto L166
L166:
	;
	v1283 = v993 + int32(1)
	if v1283 != v846 {
		v993 = v1283
		goto L162
	} else {
		goto L198
	}
L167:
	;
	v1045 = int32(3)
	v1047 = v888 + v1017<<(uint(v1045)%32)
	v1048 = *(*float64)(unsafe.Add(mBase, uint32(v1047)))
	v1049 = v1017 - v1011
	v1052 = v888 + v1049<<(uint(v1045)%32)
	v1053 = *(*float64)(unsafe.Add(mBase, uint32(v1052)))
	if base.F64_le(v1048, base.F64_add(v1053, float64(1))) != 0 {
		goto L169
	} else {
		goto L170
	}
L168:
	;
	goto L166
L169:
	;
	v1059 = v892 + v1017<<(uint(int32(2))%32)
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(v1059)))
	if v1011 != 0 {
		goto L172
	} else {
		goto L173
	}
L170:
	;
	goto L171
L171:
	;
	v1248 = v1017 - int32(1)
	if v1011 <= v1248 {
		v1017 = v1248
		goto L167
	} else {
		goto L197
	}
L172:
	;
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(v892+v1049<<(uint(int32(2))%32))))
	if v1060 == int32(0) {
		goto L176
	} else {
		goto L177
	}
L173:
	;
	v1187 = v1060
	goto L174
L174:
	;
	v1208 = F_bms_add_member(m, v1187, v993)
	mBase = m.M
	v1209 = m.ExcPending
	if v1209 != 0 {
		goto L47
	} else {
		goto L196
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1059))) = v1174
	v1187 = v1174
	goto L174
L176:
	;
	v1067 = int32(0)
	if v1064 == v1067 {
		v1174 = v1067
		goto L175
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	if v1064 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L179:
	;
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+4))
	v1074 = v1070<<(uint(int32(2))%32) + int32(8)
	v1075 = F_palloc(m, v1074)
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L47
	} else {
		goto L180
	}
L180:
	;
	if v1074 != 0 {
		goto L182
	} else {
		goto L183
	}
L181:
	;
	v1174 = v1078
	goto L175
L182:
	;
	v1077 = F__emscripten_memcpy_bulkmem(m, v1075, v1064, v1074)
	mBase = m.M
	v1078 = v1077
	goto L184
L183:
	;
	v1078 = v1075
	goto L184
L184:
	;
	goto L181
L185:
	;
	F_pfree(m, v1060)
	mBase = m.M
	v1082 = m.ExcPending
	if v1082 != 0 {
		goto L47
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+4))
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(v1060)+4))
	if v1085 < v1084 {
		goto L189
	} else {
		goto L190
	}
L188:
	;
	v1174 = int32(0)
	goto L175
L189:
	;
	v1091 = F_repalloc(m, v1060, v1084<<(uint(int32(2))%32)+int32(8))
	mBase = m.M
	v1092 = m.ExcPending
	if v1092 != 0 {
		goto L47
	} else {
		goto L192
	}
L190:
	;
	v1093 = v1060
	goto L191
L191:
	;
	v1094 = int32(8)
	v1107 = int32(0)
	goto L193
L192:
	;
	v1093 = v1091
	goto L191
L193:
	;
	v1132 = v1107 << (uint(int32(2)) % 32)
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(v1132+(v1064+v1094))))
	*(*int32)(unsafe.Add(mBase, uint32(v1093+v1094+v1132))) = v1135
	v1138 = v1107 + int32(1)
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+4))
	if v1138 < v1139 {
		v1107 = v1138
		goto L193
	} else {
		goto L195
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1093)+4)) = v1139
	v1174 = v1093
	goto L175
L195:
	;
	goto L194
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1059))) = v1208
	v1211 = *(*float64)(unsafe.Add(mBase, uint32(v1052)))
	*(*float64)(unsafe.Add(mBase, uint32(v1047))) = base.F64_add(v1211, float64(1))
	goto L171
L197:
	;
	goto L168
L198:
	;
	goto L163
L199:
	;
	v1325 = F_bms_del_member(m, v1323, v846)
	mBase = m.M
	v1326 = m.ExcPending
	if v1326 != 0 {
		goto L47
	} else {
		goto L200
	}
L200:
	;
	F_MemoryContextDelete(m, v878)
	mBase = m.M
	v1328 = m.ExcPending
	if v1328 != 0 {
		goto L47
	} else {
		goto L201
	}
L201:
	;
	if v1325 == int32(0) {
		v1409 = l0
		v1410 = l1
		v1411 = l2
		v1414 = l5
		v1415 = l6
		v1423 = v9
		v1425 = v708
		v1428 = v35
		v1432 = v37
		goto L111
	} else {
		goto L202
	}
L202:
	;
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(v1331)+12))
	v1333 = *(*int32)(unsafe.Add(mBase, uint32(v1332)))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = v1333
	v1336 = int32(1)
	v1340 = F_list_make1_impl(m, v1336, v35+int32(12))
	mBase = m.M
	v1341 = m.ExcPending
	if v1341 != 0 {
		goto L47
	} else {
		goto L203
	}
L203:
	;
	v1342 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v1342 == int32(0) {
		v1409 = l0
		v1410 = l1
		v1411 = l2
		v1414 = l5
		v1415 = l6
		v1423 = v1340
		v1425 = v708
		v1428 = v35
		v1432 = v37
		goto L111
	} else {
		goto L204
	}
L204:
	;
	v1345 = *(*int32)(unsafe.Add(mBase, uint32(v1342)+4))
	if v1345 < int32(2) {
		v1409 = l0
		v1410 = l1
		v1411 = l2
		v1414 = l5
		v1415 = l6
		v1423 = v1340
		v1425 = v708
		v1428 = v35
		v1432 = v37
		goto L111
	} else {
		goto L205
	}
L205:
	;
	v1353 = int32(0)
	v1357 = v1336
	v1363 = v1340
	v1365 = v708
	goto L206
L206:
	;
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(v1342)+12))
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(v1381+v1357<<(uint(int32(2))%32))))
	v1386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1385)+24)))
	if v1386 == int32(1) {
		goto L209
	} else {
		goto L210
	}
L207:
	;
	v1409 = l0
	v1410 = l1
	v1411 = l2
	v1414 = l5
	v1415 = l6
	v1423 = v1403
	v1425 = v1404
	v1428 = v35
	v1432 = v37
	goto L111
L208:
	;
	v1406 = v1357 + int32(1)
	v1407 = *(*int32)(unsafe.Add(mBase, uint32(v1342)+4))
	if v1406 < v1407 {
		v1353 = v1402
		v1357 = v1406
		v1363 = v1403
		v1365 = v1404
		goto L206
	} else {
		goto L219
	}
L209:
	;
	v1389 = F_bms_is_member(m, v1353, v1325)
	mBase = m.M
	v1390 = m.ExcPending
	if v1390 != 0 {
		goto L47
	} else {
		goto L212
	}
L210:
	;
	goto L211
L211:
	;
	v1400 = F_lappend(m, v1363, v1385)
	mBase = m.M
	v1401 = m.ExcPending
	if v1401 != 0 {
		goto L47
	} else {
		goto L218
	}
L212:
	;
	if v1389 != 0 {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v1391 = *(*int32)(unsafe.Add(mBase, uint32(v1385)+12))
	v1392 = F_list_concat(m, v1365, v1391)
	mBase = m.M
	v1393 = m.ExcPending
	if v1393 != 0 {
		goto L47
	} else {
		goto L216
	}
L214:
	;
	goto L215
L215:
	;
	v1396 = F_lappend(m, v1363, v1385)
	mBase = m.M
	v1397 = m.ExcPending
	if v1397 != 0 {
		goto L47
	} else {
		goto L217
	}
L216:
	;
	v1402 = v1353 + int32(1)
	v1403 = v1363
	v1404 = v1392
	goto L208
L217:
	;
	v1402 = v1353 + int32(1)
	v1403 = v1396
	v1404 = v1365
	goto L208
L218:
	;
	v1402 = v1353
	v1403 = v1400
	v1404 = v1365
	goto L208
L219:
	;
	goto L207
L220:
	;
	if v1787 == int32(0) {
		v1813 = v1409
		v1814 = v1410
		v1815 = v1411
		v1818 = v1414
		v1819 = v1415
		v1832 = v1428
		v1836 = v1432
		goto L107
	} else {
		goto L257
	}
L221:
	;
	v1449 = *(*int32)(unsafe.Add(mBase, uint32(v1425)+4))
	if v1449 <= int32(0) {
		v1787 = v1448
		goto L220
	} else {
		goto L227
	}
L222:
	;
	if v1425 == int32(0) {
		v1787 = v1423
		goto L220
	} else {
		goto L226
	}
L223:
	;
	if v1425 == int32(0) {
		goto L222
	} else {
		goto L224
	}
L224:
	;
	v1443 = *(*int32)(unsafe.Add(mBase, uint32(v1414)))
	v1444 = F_list_copy(m, v1443)
	mBase = m.M
	v1445 = m.ExcPending
	if v1445 != 0 {
		goto L47
	} else {
		goto L225
	}
L225:
	;
	v1448 = v1444
	goto L221
L226:
	;
	v1448 = v1423
	goto L221
L227:
	;
	v1466 = v1448
	v1473 = v9
	goto L228
L228:
	;
	v1484 = *(*int32)(unsafe.Add(mBase, uint32(v1425)+12))
	v1488 = *(*int32)(unsafe.Add(mBase, uint32(v1484+v1473<<(uint(int32(2))%32))))
	v1490 = F_palloc0(m, int32(32))
	mBase = m.M
	v1491 = m.ExcPending
	if v1491 != 0 {
		goto L47
	} else {
		goto L230
	}
L229:
	;
	v1787 = v1767
	goto L220
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1490))) = int32(309)
	v1494 = *(*int32)(unsafe.Add(mBase, uint32(v1488)+4))
	v1495 = F_preprocess_groupclause(m, v1409, v1494)
	mBase = m.M
	v1496 = m.ExcPending
	if v1496 != 0 {
		goto L47
	} else {
		goto L231
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1490)+4)) = v1495
	*(*int32)(unsafe.Add(mBase, uint32(v1428)+8)) = v1488
	*(*int32)(unsafe.Add(mBase, uint32(v1428)+20)) = v1488
	v1503 = F_list_make1_impl(m, int32(1), v1428+int32(8))
	mBase = m.M
	v1504 = m.ExcPending
	if v1504 != 0 {
		goto L47
	} else {
		goto L232
	}
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1490)+12)) = v1503
	v1506 = *(*int32)(unsafe.Add(mBase, uint32(v1414)+32))
	v1507 = *(*int32)(unsafe.Add(mBase, uint32(v1490)+4))
	if v1507 == int32(0) {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	if v1503 == int32(0) {
		goto L240
	} else {
		goto L241
	}
L234:
	;
	v1510 = int32(0)
	v1511 = *(*int32)(unsafe.Add(mBase, uint32(v1507)+4))
	if v1511 <= v1510 {
		goto L233
	} else {
		goto L235
	}
L235:
	;
	v1522 = v1510
	goto L236
L236:
	;
	v1546 = *(*int32)(unsafe.Add(mBase, uint32(v1507)+12))
	v1547 = int32(2)
	v1550 = *(*int32)(unsafe.Add(mBase, uint32(v1546+v1522<<(uint(v1547)%32))))
	v1551 = *(*int32)(unsafe.Add(mBase, uint32(v1550)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1506+v1551<<(uint(v1547)%32)))) = v1522
	v1557 = v1522 + int32(1)
	v1558 = *(*int32)(unsafe.Add(mBase, uint32(v1507)+4))
	if v1557 < v1558 {
		v1522 = v1557
		goto L236
	} else {
		goto L238
	}
L237:
	;
	goto L233
L238:
	;
	goto L237
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1490)+8)) = v1745
	v1763 = *(*float64)(unsafe.Add(mBase, uint32(v1488)+8))
	v1764 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v1490)+24)) = uint16(v1764)
	*(*float64)(unsafe.Add(mBase, uint32(v1490)+16)) = v1763
	v1767 = F_lcons(m, v1490, v1466)
	mBase = m.M
	v1768 = m.ExcPending
	if v1768 != 0 {
		goto L47
	} else {
		goto L255
	}
L240:
	;
	v1745 = int32(0)
	goto L239
L241:
	;
	goto L242
L242:
	;
	v1595 = int32(0)
	v1597 = *(*int32)(unsafe.Add(mBase, uint32(v1503)+4))
	if v1597 <= v1595 {
		v1745 = v1595
		goto L239
	} else {
		goto L243
	}
L243:
	;
	v1611 = v1595
	v1615 = v1595
	goto L244
L244:
	;
	v1632 = int32(0)
	v1633 = *(*int32)(unsafe.Add(mBase, uint32(v1503)+12))
	v1637 = *(*int32)(unsafe.Add(mBase, uint32(v1633+v1611<<(uint(int32(2))%32))))
	v1638 = *(*int32)(unsafe.Add(mBase, uint32(v1637)+4))
	if v1638 == v1632 {
		v1695 = v1632
		goto L246
	} else {
		goto L247
	}
L245:
	;
	v1745 = v1724
	goto L239
L246:
	;
	v1724 = F_lappend(m, v1615, v1695)
	mBase = m.M
	v1725 = m.ExcPending
	if v1725 != 0 {
		goto L47
	} else {
		goto L253
	}
L247:
	;
	v1641 = int32(0)
	v1642 = *(*int32)(unsafe.Add(mBase, uint32(v1638)+4))
	if v1642 <= v1641 {
		v1695 = v1632
		goto L246
	} else {
		goto L248
	}
L248:
	;
	v1648 = v1632
	v1653 = v1641
	goto L249
L249:
	;
	v1677 = *(*int32)(unsafe.Add(mBase, uint32(v1638)+12))
	v1678 = int32(2)
	v1681 = *(*int32)(unsafe.Add(mBase, uint32(v1677+v1653<<(uint(v1678)%32))))
	v1685 = *(*int32)(unsafe.Add(mBase, uint32(v1506+v1681<<(uint(v1678)%32))))
	v1686 = F_lappend_int(m, v1648, v1685)
	mBase = m.M
	v1687 = m.ExcPending
	if v1687 != 0 {
		goto L47
	} else {
		goto L251
	}
L250:
	;
	v1695 = v1686
	goto L246
L251:
	;
	v1689 = v1653 + int32(1)
	v1690 = *(*int32)(unsafe.Add(mBase, uint32(v1638)+4))
	if v1689 < v1690 {
		v1648 = v1686
		v1653 = v1689
		goto L249
	} else {
		goto L252
	}
L252:
	;
	goto L250
L253:
	;
	v1727 = v1611 + int32(1)
	v1728 = *(*int32)(unsafe.Add(mBase, uint32(v1503)+4))
	if v1727 < v1728 {
		v1611 = v1727
		v1615 = v1724
		goto L244
	} else {
		goto L254
	}
L254:
	;
	goto L245
L255:
	;
	v1770 = v1473 + int32(1)
	v1771 = *(*int32)(unsafe.Add(mBase, uint32(v1425)+4))
	if v1770 < v1771 {
		v1466 = v1767
		v1473 = v1770
		goto L228
	} else {
		goto L256
	}
L256:
	;
	goto L229
L257:
	;
	v1807 = *(*int32)(unsafe.Add(mBase, uint32(v1432)+112))
	v1809 = F_create_groupingsets_path(m, v1409, v1410, v1411, v1807, int32(3), v1787, v1415)
	mBase = m.M
	v1810 = m.ExcPending
	if v1810 != 0 {
		goto L47
	} else {
		goto L258
	}
L258:
	;
	F_add_path(m, v1410, v1809)
	mBase = m.M
	v1812 = m.ExcPending
	if v1812 != 0 {
		goto L47
	} else {
		goto L259
	}
L259:
	;
	v1813 = v1409
	v1814 = v1410
	v1815 = v1411
	v1818 = v1414
	v1819 = v1415
	v1832 = v1428
	v1836 = v1432
	goto L107
L260:
	;
	v1846 = *(*int32)(unsafe.Add(mBase, uint32(v1836)+112))
	v1848 = *(*int32)(unsafe.Add(mBase, uint32(v1818)))
	v1849 = F_create_groupingsets_path(m, v1813, v1814, v1815, v1846, int32(1), v1848, v1819)
	mBase = m.M
	v1850 = m.ExcPending
	if v1850 != 0 {
		goto L47
	} else {
		goto L261
	}
L261:
	;
	v1932 = v1832
	v1945 = v1849
	goto L9
L262:
	;
	if v142 == int32(0) {
		goto L264
	} else {
		goto L265
	}
L263:
	;
	v1910 = *(*int32)(unsafe.Add(mBase, uint32(v37)+112))
	v1911 = F_create_groupingsets_path(m, l0, l1, l2, v1910, v1907, v1908, l6)
	mBase = m.M
	v1912 = m.ExcPending
	if v1912 != 0 {
		goto L47
	} else {
		goto L272
	}
L264:
	;
	if v1871 == int32(0) {
		goto L267
	} else {
		goto L268
	}
L265:
	;
	v1903 = v142
	goto L266
L266:
	;
	v1905 = F_lappend(m, v1867, v1903)
	mBase = m.M
	v1906 = m.ExcPending
	if v1906 != 0 {
		goto L47
	} else {
		goto L271
	}
L267:
	;
	v1907 = int32(2)
	v1908 = v1867
	goto L263
L268:
	;
	goto L269
L269:
	;
	v1891 = F_palloc0(m, int32(32))
	mBase = m.M
	v1892 = m.ExcPending
	if v1892 != 0 {
		goto L47
	} else {
		goto L270
	}
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1891)+12)) = v1873
	*(*int64)(unsafe.Add(mBase, uint32(v1891))) = int64(309)
	*(*int32)(unsafe.Add(mBase, uint32(v1891)+8)) = v1871
	v1897 = *(*int32)(unsafe.Add(mBase, uint32(v1871)+4))
	v1898 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1891)+24)) = uint16(v1898)
	*(*float64)(unsafe.Add(mBase, uint32(v1891)+16)) = base.F64_convert_i32_s(v1897)
	v1903 = v1891
	goto L266
L271:
	;
	v1907 = int32(3)
	v1908 = v1905
	goto L263
L272:
	;
	v1932 = v35
	v1945 = v1911
	goto L9
L273:
	;
	v1967 = v1932
	goto L8
}
func F_conv_18030_to_utf8(m *base.Module, l0 int32) int32 {
	var v13 int32
	_ = v13
	var v24 int32
	_ = v24
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v157 int32
	_ = v157
	var v165 int32
	_ = v165
	var v191 int32
	_ = v191
	var v202 int32
	_ = v202
	var v224 int32
	_ = v224
	var v232 int32
	_ = v232
	var v254 int32
	_ = v254
	var v262 int32
	_ = v262
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v324 int32
	_ = v324
	var v335 int32
	_ = v335
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v393 int32
	_ = v393
	var v401 int32
	_ = v401
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v461 int32
	_ = v461
	var v479 int32
	_ = v479
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v539 int32
	_ = v539
	var v557 int32
	_ = v557
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v619 int32
	_ = v619
	var v630 int32
	_ = v630
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v696 int32
	_ = v696
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v756 int32
	_ = v756
	var v774 int32
	_ = v774
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	if base.Ui32(l0+int32(2127506640)) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_0)) {
		v13 = int32(255)
		v24 = int32(base.Ui32(l0)>>(uint(int32(16))%32))&int32(55)*int32(1260) + l0&v13 + int32(base.Ui32(l0)>>(uint(int32(8))%32))&v13*int32(10) - int32(_a_F_conv_18030_to_utf8_1)
		if base.Ui32(v24) < base.Ui32(int32(128)) {
			v826 = v24
			return v826
		} else {
			if base.Ui32(v24) <= base.Ui32(int32(2047)) {
				return v24&int32(63) | v24<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_2) | int32(_a_F_conv_18030_to_utf8_3)
			} else {
				if base.Ui32(v24) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_4)) {
					return v24<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_5) | (v24&int32(63) | v24<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_6)) | int32(14712960)
				} else {
					return v24<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_6) | (v24<<(uint(int32(6))%32)&int32(117440512) | (v24&int32(63) | v24<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_7))) | int32(-260013952)
				}
			}
		}
	} else {
		if base.Ui32(l0+int32(2127058887)) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_8)) {
			v86 = int32(255)
			v95 = int32(base.Ui32(l0)>>(uint(int32(16))%32))&int32(63)*int32(1260) + l0&v86 + int32(base.Ui32(l0)>>(uint(int32(8))%32))&v86*int32(10)
			v97 = v95 - int32(_a_F_conv_18030_to_utf8_9)
			if base.Ui32(v97) < base.Ui32(int32(128)) {
				v826 = v97
				return v826
			} else {
				v101 = v95 - int32(_a_F_conv_18030_to_utf8_10)
				if base.Ui32(v97) <= base.Ui32(int32(2047)) {
					return v101&int32(63) | v97<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_2) | int32(_a_F_conv_18030_to_utf8_3)
				} else {
					if base.Ui32(v97) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_4)) {
						return v97<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_5) | (v101&int32(63) | v97<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_6)) | int32(14712960)
					} else {
						return v97<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_6) | (v97<<(uint(int32(6))%32)&int32(117440512) | (v101&int32(63) | v97<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_7))) | int32(-260013952)
					}
				}
			}
		} else {
			if base.Ui32(l0+int32(2110740941)) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_11)) {
				v157 = int32(255)
				v165 = int32(base.Ui32(l0)>>(uint(int32(8))%32))&v157*int32(10) + l0&v157 + int32(_a_F_conv_18030_to_utf8_12)
				return v165&int32(63) | v165<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_6) | v165<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_13) | int32(14712960)
			} else {
				if base.Ui32(l0+int32(2110663624)) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_14)) {
					v191 = int32(255)
					v202 = int32(base.Ui32(l0)>>(uint(int32(16))%32))&int32(51)*int32(1260) + l0&v191 + int32(base.Ui32(l0)>>(uint(int32(8))%32))&v191*int32(10) - int32(_a_F_conv_18030_to_utf8_15)
					return v202&int32(63) | v202<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_6) | v202<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_5) | int32(14712960)
				} else {
					if base.Ui32(l0+int32(2110600905)) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_16)) {
						v224 = int32(255)
						v232 = int32(base.Ui32(l0)>>(uint(int32(8))%32))&v224*int32(10) + l0&v224 + int32(_a_F_conv_18030_to_utf8_17)
						return v232&int32(63) | v232<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_6) | v232<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_13) | int32(14712960)
					} else {
						if base.Ui32(l0+int32(2110545095)) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_18)) {
							v254 = int32(255)
							v262 = int32(base.Ui32(l0)>>(uint(int32(8))%32))&v254*int32(10) + l0&v254 + int32(_a_F_conv_18030_to_utf8_19)
							if base.Ui32(int32(128)) <= base.Ui32(v262) {
								if base.Ui32(v262) <= base.Ui32(int32(2047)) {
									v312 = v262&int32(63) | v262<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_2) | int32(_a_F_conv_18030_to_utf8_3)
								} else {
									if base.Ui32(v262) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_4)) {
										v312 = v262<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_5) | (v262&int32(63) | v262<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_6)) | int32(14712960)
									} else {
										v311 = v262<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_6) | (v262<<(uint(int32(6))%32)&int32(117440512) | (v262&int32(63) | v262<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_7))) | int32(-260013952)
										v312 = v311
									}
								}
							} else {
								v311 = v262
								v312 = v311
							}
							return v312
						} else {
							if base.Ui32(l0+int32(2110527432)) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_20)) {
								v324 = int32(255)
								v335 = int32(base.Ui32(l0)>>(uint(int32(16))%32))&int32(55)*int32(1260) + l0&v324 + int32(base.Ui32(l0)>>(uint(int32(8))%32))&v324*int32(10) - int32(_a_F_conv_18030_to_utf8_21)
								if base.Ui32(int32(128)) <= base.Ui32(v335) {
									if base.Ui32(v335) <= base.Ui32(int32(2047)) {
										v385 = v335&int32(63) | v335<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_2) | int32(_a_F_conv_18030_to_utf8_3)
									} else {
										if base.Ui32(v335) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_4)) {
											v385 = v335<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_5) | (v335&int32(63) | v335<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_6)) | int32(14712960)
										} else {
											v384 = v335<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_6) | (v335<<(uint(int32(6))%32)&int32(117440512) | (v335&int32(63) | v335<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_7))) | int32(-260013952)
											v385 = v384
										}
									}
								} else {
									v384 = v335
									v385 = v384
								}
								return v385
							} else {
								if base.Ui32(l0+int32(2110480079)) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_22)) {
									v393 = int32(255)
									v401 = int32(base.Ui32(l0)>>(uint(int32(8))%32))&v393*int32(10) + l0&v393 + int32(_a_F_conv_18030_to_utf8_23)
									if base.Ui32(int32(128)) <= base.Ui32(v401) {
										if base.Ui32(v401) <= base.Ui32(int32(2047)) {
											v451 = v401&int32(63) | v401<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_2) | int32(_a_F_conv_18030_to_utf8_3)
										} else {
											if base.Ui32(v401) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_4)) {
												v451 = v401<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_5) | (v401&int32(63) | v401<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_6)) | int32(14712960)
											} else {
												v450 = v401<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_6) | (v401<<(uint(int32(6))%32)&int32(117440512) | (v401&int32(63) | v401<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_7))) | int32(-260013952)
												v451 = v450
											}
										}
									} else {
										v450 = v401
										v451 = v450
									}
									return v451
								} else {
									if base.Ui32(l0+int32(2110419149)) <= base.Ui32(int32(16857093)) {
										v461 = int32(255)
										v479 = int32(base.Ui32(l0)>>(uint(int32(24))%32))*int32(_a_F_conv_18030_to_utf8_24) + l0&v461 + int32(base.Ui32(l0)>>(uint(int32(16))%32))&v461*int32(1260) + int32(base.Ui32(l0)>>(uint(int32(8))%32))&v461*int32(10) - int32(_a_F_conv_18030_to_utf8_25)
										if base.Ui32(int32(128)) <= base.Ui32(v479) {
											if base.Ui32(v479) <= base.Ui32(int32(2047)) {
												v529 = v479&int32(63) | v479<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_2) | int32(_a_F_conv_18030_to_utf8_3)
											} else {
												if base.Ui32(v479) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_4)) {
													v529 = v479<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_5) | (v479&int32(63) | v479<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_6)) | int32(14712960)
												} else {
													v528 = v479<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_6) | (v479<<(uint(int32(6))%32)&int32(117440512) | (v479&int32(63) | v479<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_7))) | int32(-260013952)
													v529 = v528
												}
											}
										} else {
											v528 = v479
											v529 = v528
										}
										return v529
									} else {
										if base.Ui32(l0+int32(2093559760)) <= base.Ui32(int32(16364804)) {
											v539 = int32(255)
											v557 = int32(base.Ui32(l0)>>(uint(int32(24))%32))*int32(_a_F_conv_18030_to_utf8_24) + l0&v539 + int32(base.Ui32(l0)>>(uint(int32(16))%32))&v539*int32(1260) + int32(base.Ui32(l0)>>(uint(int32(8))%32))&v539*int32(10) - int32(_a_F_conv_18030_to_utf8_26)
											if base.Ui32(int32(128)) <= base.Ui32(v557) {
												if base.Ui32(v557) <= base.Ui32(int32(2047)) {
													v607 = v557&int32(63) | v557<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_2) | int32(_a_F_conv_18030_to_utf8_3)
												} else {
													if base.Ui32(v557) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_4)) {
														v607 = v557<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_5) | (v557&int32(63) | v557<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_6)) | int32(14712960)
													} else {
														v606 = v557<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_6) | (v557<<(uint(int32(6))%32)&int32(117440512) | (v557&int32(63) | v557<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_7))) | int32(-260013952)
														v607 = v606
													}
												}
											} else {
												v606 = v557
												v607 = v606
											}
											return v607
										} else {
											if base.Ui32(l0+int32(2077189064)) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_27)) {
												v619 = int32(255)
												v630 = int32(base.Ui32(l0)>>(uint(int32(16))%32))&int32(49)*int32(1260) + l0&v619 + int32(base.Ui32(l0)>>(uint(int32(8))%32))&v619*int32(10) + int32(1946)
												if base.Ui32(int32(128)) <= base.Ui32(v630) {
													if base.Ui32(v630) <= base.Ui32(int32(2047)) {
														v680 = v630&int32(63) | v630<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_2) | int32(_a_F_conv_18030_to_utf8_3)
													} else {
														if base.Ui32(v630) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_4)) {
															v680 = v630<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_5) | (v630&int32(63) | v630<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_6)) | int32(14712960)
														} else {
															v679 = v630<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_6) | (v630<<(uint(int32(6))%32)&int32(117440512) | (v630&int32(63) | v630<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_7))) | int32(-260013952)
															v680 = v679
														}
													}
												} else {
													v679 = v630
													v680 = v679
												}
												return v680
											} else {
												if base.Ui32(l0+int32(2077121996)) <= base.Ui32(int32(517)) {
													v696 = int32(base.Ui32(l0)>>(uint(int32(8))%32))&int32(167)*int32(10) + l0&int32(255) + int32(_a_F_conv_18030_to_utf8_28)
													if base.Ui32(int32(128)) <= base.Ui32(v696) {
														if base.Ui32(v696) <= base.Ui32(int32(2047)) {
															v746 = v696&int32(63) | v696<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_2) | int32(_a_F_conv_18030_to_utf8_3)
														} else {
															if base.Ui32(v696) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_4)) {
																v746 = v696<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_5) | (v696&int32(63) | v696<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_6)) | int32(14712960)
															} else {
																v745 = v696<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_6) | (v696<<(uint(int32(6))%32)&int32(117440512) | (v696&int32(63) | v696<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_7))) | int32(-260013952)
																v746 = v745
															}
														}
													} else {
														v745 = v696
														v746 = v745
													}
													return v746
												} else {
													if base.Ui32(int32(1392646405)) < base.Ui32(l0+int32(1875869392)) {
														v826 = int32(0)
													} else {
														v756 = int32(255)
														v774 = int32(base.Ui32(l0)>>(uint(int32(24))%32))*int32(_a_F_conv_18030_to_utf8_24) + l0&v756 + int32(base.Ui32(l0)>>(uint(int32(16))%32))&v756*int32(1260) + int32(base.Ui32(l0)>>(uint(int32(8))%32))&v756*int32(10) - int32(_a_F_conv_18030_to_utf8_29)
														if base.Ui32(int32(128)) <= base.Ui32(v774) {
															if base.Ui32(v774) <= base.Ui32(int32(2047)) {
																v824 = v774&int32(63) | v774<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_2) | int32(_a_F_conv_18030_to_utf8_3)
															} else {
																if base.Ui32(v774) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_4)) {
																	v824 = v774<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_5) | (v774&int32(63) | v774<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_6)) | int32(14712960)
																} else {
																	v823 = v774<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_6) | (v774<<(uint(int32(6))%32)&int32(117440512) | (v774&int32(63) | v774<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_7))) | int32(-260013952)
																	v824 = v823
																}
															}
														} else {
															v823 = v774
															v824 = v823
														}
														v826 = v824
													}
													return v826
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
func F_convert_saop_to_hashed_saop(m *base.Module, l0 int32) {
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = F_convert_saop_to_hashed_saop_walker(m, l0, l0)
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		return
	}
}
func F_convert_tuples_by_name_attrmap(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
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
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8 = F_palloc(m, int32(28))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
		v17 = F_palloc(m, v6<<(uint(int32(2))%32))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v17
			v20 = F_palloc(m, v6)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v20
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v25 = v23 + int32(1)
				v28 = F_palloc(m, v25<<(uint(int32(2))%32))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v28
					v31 = F_palloc(m, v25)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v31
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v34))) = int32(0)
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
						v38 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v37))) = uint8(v38)
						return v8
					}
				}
			}
		}
	}
}
func F_cost_ctescan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 float64
	_ = v7
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 float64
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 float64
	_ = v24
	var v25 int64
	_ = v25
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 float64
	_ = v62
	var v63 float64
	_ = v63
	var v64 float64
	_ = v64
	var v66 float64
	_ = v66
	var v73 float64
	_ = v73
	var v74 float64
	_ = v74
	var v76 float64
	_ = v76
	var v81 float64
	_ = v81
	var v82 float64
	_ = v82
	var v84 float64
	_ = v84
	var v86 float64
	_ = v86
	var v88 float64
	_ = v88
	var v90 float64
	_ = v90
	var v91 float64
	_ = v91
	var v99 float64
	_ = v99
	var v100 float64
	_ = v100
	var v101 float64
	_ = v101
	var v102 float64
	_ = v102
	var v106 float64
	_ = v106
	var v107 float64
	_ = v107
	var v108 int32
	_ = v108
	var v109 float64
	_ = v109
	var v110 float64
	_ = v110
	var v113 float64
	_ = v113
	var v115 float64
	_ = v115
	v7 = float64(0)
	v15 = m.G0
	v17 = v15 - int32(32)
	m.G0 = v17
	if l3 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v107 = *(*float64)(unsafe.Add(mBase, uint32(l2)+120))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v109 = *(*float64)(unsafe.Add(mBase, uint32(v108)+24))
	v110 = *(*float64)(unsafe.Add(mBase, uint32(v108)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
	v113 = float64(0)
	v115 = base.F64_add(v110, base.F64_add(v106, v113))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = v115
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = base.F64_add(v115, base.F64_add(base.F64_mul(v109, v101), base.F64_add(base.F64_mul(v107, base.F64_add(v100, base.F64_add(v99, v102))), v113)))
	m.G0 = v17 + int32(32)
	return
L2:
	;
	v19 = *(*float64)(unsafe.Add(mBase, uint32(l3)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v22 = int32(0)
	v24 = *(*float64)(unsafe.Add(mBase, _c_F_cost_ctescan[0]))
	v25 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = v25
	*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = l1
	if v21 == v22 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	goto L4
L4:
	;
	v86 = *(*float64)(unsafe.Add(mBase, uint32(l2)+16))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v86
	v88 = *(*float64)(unsafe.Add(mBase, uint32(l2)+200))
	v90 = *(*float64)(unsafe.Add(mBase, _c_F_cost_ctescan[0]))
	v91 = *(*float64)(unsafe.Add(mBase, uint32(l2)+192))
	v99 = v88
	v100 = v90
	v101 = v86
	v102 = v90
	v106 = v91
	goto L1
L5:
	;
	v82 = *(*float64)(unsafe.Add(mBase, uint32(l2)+200))
	v84 = *(*float64)(unsafe.Add(mBase, uint32(l2)+192))
	v99 = base.F64_add(v74, v82)
	v100 = v24
	v101 = v76
	v102 = v81
	v106 = base.F64_add(v73, v84)
	goto L1
L6:
	;
	v73 = v7
	v74 = v7
	v76 = v19
	v81 = v24
	goto L5
L7:
	;
	goto L8
L8:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v32 <= int32(0) {
		v73 = v7
		v74 = v7
		v76 = v19
		v81 = v24
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v38 = v22
	goto L10
L10:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49+v38<<(uint(int32(2))%32))))
	v56 = F_cost_qual_eval_walker(m, v53, v17+int32(8))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v62 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
	v63 = *(*float64)(unsafe.Add(mBase, uint32(v17)+24))
	v64 = *(*float64)(unsafe.Add(mBase, uint32(v17)+16))
	v66 = *(*float64)(unsafe.Add(mBase, _c_F_cost_ctescan[0]))
	v73 = v64
	v74 = v63
	v76 = v62
	v81 = v66
	goto L5
L12:
	;
	return
L13:
	;
	v59 = v38 + int32(1)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v59 < v60 {
		v38 = v59
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
}
func F_cost_incremental_sort(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 float64, l6 float64, l7 float64, l8 int32, l9 int32, l10 float64) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 float64
	_ = v23
	var v26 float64
	_ = v26
	var v29 int32
	_ = v29
	var v32 float64
	_ = v32
	var v35 float64
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
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
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
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
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v103 float64
	_ = v103
	var v104 int32
	_ = v104
	var v120 float64
	_ = v120
	var v124 int32
	_ = v124
	var v125 float64
	_ = v125
	var v131 float64
	_ = v131
	var v134 float64
	_ = v134
	var v136 float64
	_ = v136
	var v139 float64
	_ = v139
	var v146 float64
	_ = v146
	var v148 float64
	_ = v148
	var v152 int32
	_ = v152
	var v153 float64
	_ = v153
	var v156 int64
	_ = v156
	var v157 float64
	_ = v157
	var v159 float64
	_ = v159
	var v160 int32
	_ = v160
	var v163 float64
	_ = v163
	var v168 float64
	_ = v168
	var v170 float64
	_ = v170
	var v171 float64
	_ = v171
	var v173 float64
	_ = v173
	var v174 float64
	_ = v174
	var v177 float64
	_ = v177
	var v180 float64
	_ = v180
	var v184 float64
	_ = v184
	var v191 float64
	_ = v191
	var v192 float64
	_ = v192
	var v195 float64
	_ = v195
	var v199 float64
	_ = v199
	var v209 float64
	_ = v209
	var v212 float64
	_ = v212
	var v216 float64
	_ = v216
	var v217 float64
	_ = v217
	var v218 float64
	_ = v218
	var v222 float64
	_ = v222
	var v224 float64
	_ = v224
	var v232 float64
	_ = v232
	v12 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v23 = float64(2)
	if base.F64_lt(l7, v23) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v26 = v23
	goto L3
L2:
	;
	v26 = l7
	goto L3
L3:
	;
	if l2 == int32(0) {
		v96 = v12
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v124 = v21 + int32(8)
	v125 = base.F64_div(v26, v120)
	v131 = float64(2)
	if base.F64_lt(v125, v131) != 0 {
		goto L25
	} else {
		goto L26
	}
L5:
	;
	v101 = int32(0)
	v103 = F_estimate_num_groups(m, l1, v96, v26, v101, v101)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L16
	} else {
		goto L23
	}
L6:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v29 <= int32(0) {
		v96 = v12
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v32 = float64(200)
	if base.F64_lt(v26, v32) != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v35 = v26
	goto L10
L9:
	;
	v35 = v32
	goto L10
L10:
	;
	v36 = int32(1)
	if l3 <= v36 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v39 = v36
	goto L13
L12:
	;
	v39 = l3
	goto L13
L13:
	;
	v54 = v12
	v55 = v12
	goto L14
L14:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v61+v54<<(uint(int32(2))%32))))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+16))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v71 = F_pull_varnos(m, l1, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v96 = v76
	goto L5
L16:
	;
	return
L17:
	;
	v73 = F_bms_is_member(m, int32(0), v71)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	if v73 != 0 {
		v120 = v35
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v76 = F_lappend(m, v55, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L16
	} else {
		goto L20
	}
L20:
	;
	if v54 == v39-int32(1) {
		v96 = v76
		goto L5
	} else {
		goto L21
	}
L21:
	;
	v80 = v54 + int32(1)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v80 < v81 {
		v54 = v80
		v55 = v76
		goto L14
	} else {
		goto L22
	}
L22:
	;
	goto L15
L23:
	;
	v120 = v103
	goto L4
L24:
	;
	v216 = *(*float64)(unsafe.Add(mBase, _c_F_cost_incremental_sort[0]))
	v217 = *(*float64)(unsafe.Add(mBase, uint32(v21)))
	v218 = *(*float64)(unsafe.Add(mBase, uint32(v21)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l4
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v26
	v222 = base.F64_div(base.F64_sub(l6, l5), v120)
	v224 = base.F64_add(v222, base.F64_add(l5, v218))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = v224
	v232 = base.F64_add(v120, float64(-1))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = base.F64_add(v224, base.F64_add(base.F64_mul(base.F64_add(v216, v216), v120), base.F64_add(base.F64_mul(base.F64_add(v216, float64(0)), v26), base.F64_add(base.F64_mul(v222, v232), base.F64_add(v217, base.F64_mul(base.F64_add(v218, v217), v232))))))
	m.G0 = v21 + int32(16)
	return
L25:
	;
	v134 = v131
	goto L27
L26:
	;
	v134 = v125
	goto L27
L27:
	;
	v136 = *(*float64)(unsafe.Add(mBase, _c_F_cost_incremental_sort[1]))
	v139 = base.F64_mul(v134, base.F64_add(base.F64_add(v136, v136), float64(0)))
	v146 = base.F64_convert_i32_u((l8+int32(7))&int32(-8) + int32(24))
	v148 = base.F64_mul(v125, v146)
	v152 = base.F64_lt(l10, v134) & base.F64_gt(l10, float64(0))
	if v152 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v124))) = v209
	v212 = *(*float64)(unsafe.Add(mBase, _c_F_cost_incremental_sort[1]))
	*(*float64)(unsafe.Add(mBase, uint32(v21))) = base.F64_mul(v134, v212)
	goto L24
L29:
	;
	v153 = base.F64_mul(l10, v146)
	goto L31
L30:
	;
	v153 = v148
	goto L31
L31:
	;
	v156 = base.I64_extend_i32_s(l9) << (uint(int64(10)) % 64)
	v157 = base.F64_convert_i64_s(v156)
	if base.F64_gt(v153, v157) != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v159 = F_log(m, v134)
	mBase = m.M
	v160 = F_tuplesort_merge_order(m, v156)
	mBase = m.M
	v163 = base.F64_mul(base.F64_div(v159, float64(0.693147180559945)), v139)
	*(*float64)(unsafe.Add(mBase, uint32(v124))) = v163
	v168 = base.F64_ceil(base.F64_mul(v148, float64(0.0001220703125)))
	v170 = base.F64_div(v148, v157)
	v171 = base.F64_convert_i32_s(v160)
	if base.F64_gt(v170, v171) != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	if v152 != 0 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v173 = F_log(m, v170)
	mBase = m.M
	v174 = F_log(m, v171)
	mBase = m.M
	v177 = base.F64_ceil(base.F64_div(v173, v174))
	goto L37
L36:
	;
	v177 = float64(1)
	goto L37
L37:
	;
	v180 = *(*float64)(unsafe.Add(mBase, _c_F_cost_incremental_sort[2]))
	v184 = *(*float64)(unsafe.Add(mBase, _c_F_cost_incremental_sort[3]))
	v209 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v168, v168), v177), base.F64_add(base.F64_mul(v180, float64(0.75)), base.F64_mul(v184, float64(0.25)))), v163)
	goto L28
L38:
	;
	v191 = l10
	goto L40
L39:
	;
	v191 = v134
	goto L40
L40:
	;
	v192 = base.F64_add(v191, v191)
	if base.F64_gt(v148, v157)|base.F64_gt(v134, v192) != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v195 = F_log(m, v192)
	mBase = m.M
	v209 = base.F64_mul(base.F64_div(v195, float64(0.693147180559945)), v139)
	goto L28
L42:
	;
	goto L43
L43:
	;
	v199 = F_log(m, v134)
	mBase = m.M
	v209 = base.F64_mul(base.F64_div(v199, float64(0.693147180559945)), v139)
	goto L28
}
