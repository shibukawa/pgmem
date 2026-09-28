package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"sync/atomic"
	"unsafe"
)

func F_CombineRangeTables(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
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
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v9 == int32(0) {
		v56 = int32(0)
	} else {
		v13 = int32(0)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
		if base.B2i32(l2 == v13)|base.B2i32(v15 <= v13) != 0 {
			v56 = v9
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
			if int32(0) < v19 {
				v27 = int32(0)
				for {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v31+v27<<(uint(int32(2))%32))))
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+28))
					if v36 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v35)+28)) = v15 + v36
					} else {
					}
					v40 = v27 + int32(1)
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
					if v40 < v41 {
						v27 = v40
						continue
					} else {
						break
					}
					break
				}
			} else {
			}
			v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v56 = v51
		}
	}
	v60 = F_list_concat(m, v56, l3)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v60
		v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v64 = F_list_concat(m, v63, l2)
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v64
			return
		}
	}
}
func F_CompareCandidateDistancesOffset(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 float32
	_ = v7
	var v8 int32
	_ = v8
	var v9 float32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*float32)(unsafe.Add(mBase, uint32(v6)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v9 = *(*float32)(unsafe.Add(mBase, uint32(v8)+4))
	if base.F32_lt(v7, v9) != 0 {
		v27 = int32(1)
	} else {
		if base.F32_gt(v7, v9) != 0 {
			v27 = int32(-1)
		} else {
			v13 = int32(1)
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			v16 = v14 - v13
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
			v19 = v17 - v13
			if base.Ui32(v16) < base.Ui32(v19) {
				v27 = v13
			} else {
				if base.Ui32(v19) < base.Ui32(v16) {
					v24 = int32(-1)
				} else {
					v24 = int32(0)
				}
				v27 = v24
			}
		}
	}
	return v27
}
func F_CompareNearestCandidates(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 float64
	_ = v9
	var v10 float64
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v9 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	if base.F64_gt(v9, v10) != 0 {
		v12 = int32(-1)
	} else {
		v12 = int32(0)
	}
	if base.F64_lt(v9, v10) != 0 {
		v14 = int32(1)
	} else {
		v14 = v12
	}
	return v14
}
func F_ConditionVariableTimedSleep(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int64
	_ = v29
	var v32 int64
	_ = v32
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
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
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v114 int64
	_ = v114
	var v115 int64
	_ = v115
	var v124 int32
	_ = v124
	var v133 int32
	_ = v133
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableTimedSleep[0]))
	if l0 != v17 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v14 + int32(16)
	return v133
L2:
	;
	F_ConditionVariablePrepareToSleep(m, l0)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v24 = base.B2i32(l1 < int32(0))
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
	v133 = int32(0)
	goto L1
L7:
	;
	v46 = v37
	goto L11
L8:
	;
	v35 = int32(33)
	v36 = int64(0)
	v37 = int32(-1)
	goto L7
L9:
	;
	goto L10
L10:
	;
	F___clock_gettime(m, int32(1), v14)
	mBase = m.M
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	v32 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+8)))
	v35 = int32(41)
	v36 = v29*int64(-1000000000) - v32
	v37 = l1
	goto L7
L11:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableTimedSleep[1]))
	v51 = F_WaitLatch(m, v50, v35, v46, l2)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L5
	} else {
		goto L13
	}
L12:
	;
	v133 = int32(0)
	goto L1
L13:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableTimedSleep[1]))
	v55 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v55
	v60 = base.AtomicRmwOr32(m, v55, int32(_a_F_ConditionVariableTimedSleep_0), v55)
	goto L14
L14:
	;
	v63 = base.AtomicRmwXchg32(m, l0, int32(0), int32(1))
	if v63 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	F_s_lock(m, l0, int32(_a_F_ConditionVariableTimedSleep_1))
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
	v75 = v70 + v72*int32(768)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+360))
	if v76 != 0 {
		v96 = v67
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L17
L19:
	;
	v97 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v97))
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableTimedSleep[4]))
	if v101 != 0 {
		goto L26
	} else {
		goto L27
	}
L20:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v75)+356))
	if v77 != 0 {
		v96 = v67
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v78 == int32(-1) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v72
	v96 = int32(1)
	goto L19
L23:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v75)+356)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v72
	goto L22
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75)+360)) = v78
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableTimedSleep[2]))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	*(*int32)(unsafe.Add(mBase, uint32(v87+v78*int32(768))+356)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v75)+356)) = int32(-1)
	goto L22
L26:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L5
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableTimedSleep[0]))
	v107 = v96 | base.B2i32(l0 != v105)
	if v107|v24 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L28
L30:
	;
	v111 = int32(1)
	F___clock_gettime(m, v111, v14)
	mBase = m.M
	v114 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+8)))
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	v124 = l1 - base.I32_trunc_sat_f64_s(base.F64_div(base.F64_convert_i64_s(v114+(v115*int64(1000000000)+v36)), float64(1e+06)))
	if int32(0) < v124 {
		v46 = v124
		goto L11
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	if v107 == int32(0) {
		goto L11
	} else {
		goto L34
	}
L33:
	;
	v133 = v111
	goto L1
L34:
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
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int64
	_ = v85
	var v86 int64
	_ = v86
	var v87 int64
	_ = v87
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v5 {
	case 0:
		v7 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[0]))
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(167772178)
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v14 = F_fwrite(m, v10, v11, int32(1), v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			if v14 == int32(1) {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
				if int32(base.Ui32(v19)>>(uint(int32(5))%32))&int32(1) == int32(0) {
					v69 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[0]))
					*(*int32)(unsafe.Add(mBase, uint32(v69))) = int32(0)
					v85 = *(*int64)(unsafe.Add(mBase, uint32(l0)+208))
					v86 = int64(*(*int32)(unsafe.Add(mBase, uint32(v4)+4)))
					v87 = v85 + v86
					*(*int64)(unsafe.Add(mBase, uint32(l0)+208)) = v87
					v92 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[1]))
					if v92 == int32(0) {
					} else {
						v96 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[2])))
						if v96&int32(1) == int32(0) {
						} else {
							v101 = int32(_a_F_CopySendEndOfRow_0)
							v103 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[3]))
							v104 = int32(1)
							*(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[3])) = v103 + v104
							v107 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
							*(*int32)(unsafe.Add(mBase, uint32(v92))) = v107 + v104
							v111 = int32(0)
							v113 = int32(_a_F_CopySendEndOfRow_1)
							v114 = base.AtomicRmwOr32(m, v111, v113, v111)
							*(*int64)(unsafe.Add(mBase, uint32(v92+v111)+232)) = v87
							v122 = base.AtomicRmwOr32(m, v111, v113, v111)
							v123 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
							*(*int32)(unsafe.Add(mBase, uint32(v92))) = v123 + v104
							v129 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[3]))
							*(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[3])) = v129 - v104
						}
					}
					v133 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
					v134 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v133))) = uint8(v134)
					*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = v134
					*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = v134
					return
				} else {
					v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
					if v26 == int32(1) {
						v30 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[4]))
						if v30 == int32(64) {
							F_ClosePipeToProgram(m, l0)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[4])) = int32(64)
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return
								} else {
									F_errcode_for_file_access(m)
									mBase = m.M
									v43 = m.ExcPending
									if v43 != 0 {
										return
									} else {
										F_errmsg(m, int32(_a_F_CopySendEndOfRow_2), int32(0))
										mBase = m.M
										v47 = m.ExcPending
										if v47 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_CopySendEndOfRow_3), int32(636), int32(_a_F_CopySendEndOfRow_4))
											mBase = m.M
											v52 = m.ExcPending
											if v52 != 0 {
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
							v41 = m.ExcPending
							if v41 != 0 {
								return
							} else {
								F_errcode_for_file_access(m)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return
								} else {
									F_errmsg(m, int32(_a_F_CopySendEndOfRow_2), int32(0))
									mBase = m.M
									v47 = m.ExcPending
									if v47 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_CopySendEndOfRow_3), int32(636), int32(_a_F_CopySendEndOfRow_4))
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
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
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return
							} else {
								F_errmsg(m, int32(_a_F_CopySendEndOfRow_5), int32(0))
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_CopySendEndOfRow_3), int32(641), int32(_a_F_CopySendEndOfRow_4))
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
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
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
				if v26 == int32(1) {
					v30 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[4]))
					if v30 == int32(64) {
						F_ClosePipeToProgram(m, l0)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[4])) = int32(64)
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return
							} else {
								F_errcode_for_file_access(m)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return
								} else {
									F_errmsg(m, int32(_a_F_CopySendEndOfRow_2), int32(0))
									mBase = m.M
									v47 = m.ExcPending
									if v47 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_CopySendEndOfRow_3), int32(636), int32(_a_F_CopySendEndOfRow_4))
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
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
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return
							} else {
								F_errmsg(m, int32(_a_F_CopySendEndOfRow_2), int32(0))
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_CopySendEndOfRow_3), int32(636), int32(_a_F_CopySendEndOfRow_4))
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
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
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						F_errcode_for_file_access(m)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return
						} else {
							F_errmsg(m, int32(_a_F_CopySendEndOfRow_5), int32(0))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_CopySendEndOfRow_3), int32(641), int32(_a_F_CopySendEndOfRow_4))
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
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
		v73 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v74 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
		v76 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[5]))
		v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+16))
		v78 = m.T0[v77].(func(*base.Module, int32, int32, int32) int32)(m, int32(100), v73, v74)
		mBase = m.M
		v79 = m.ExcPending
		if v79 != 0 {
			return
		} else {
			v85 = *(*int64)(unsafe.Add(mBase, uint32(l0)+208))
			v86 = int64(*(*int32)(unsafe.Add(mBase, uint32(v4)+4)))
			v87 = v85 + v86
			*(*int64)(unsafe.Add(mBase, uint32(l0)+208)) = v87
			v92 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[1]))
			if v92 == int32(0) {
			} else {
				v96 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[2])))
				if v96&int32(1) == int32(0) {
				} else {
					v101 = int32(_a_F_CopySendEndOfRow_0)
					v103 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[3]))
					v104 = int32(1)
					*(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[3])) = v103 + v104
					v107 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
					*(*int32)(unsafe.Add(mBase, uint32(v92))) = v107 + v104
					v111 = int32(0)
					v113 = int32(_a_F_CopySendEndOfRow_1)
					v114 = base.AtomicRmwOr32(m, v111, v113, v111)
					*(*int64)(unsafe.Add(mBase, uint32(v92+v111)+232)) = v87
					v122 = base.AtomicRmwOr32(m, v111, v113, v111)
					v123 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
					*(*int32)(unsafe.Add(mBase, uint32(v92))) = v123 + v104
					v129 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[3]))
					*(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[3])) = v129 - v104
				}
			}
			v133 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
			v134 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v133))) = uint8(v134)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = v134
			*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = v134
			return
		}
	case 2:
		v80 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v81 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
		v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
		m.T0[v82].(func(*base.Module, int32, int32))(m, v80, v81)
		mBase = m.M
		v84 = m.ExcPending
		if v84 != 0 {
			return
		} else {
			v85 = *(*int64)(unsafe.Add(mBase, uint32(l0)+208))
			v86 = int64(*(*int32)(unsafe.Add(mBase, uint32(v4)+4)))
			v87 = v85 + v86
			*(*int64)(unsafe.Add(mBase, uint32(l0)+208)) = v87
			v92 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[1]))
			if v92 == int32(0) {
			} else {
				v96 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[2])))
				if v96&int32(1) == int32(0) {
				} else {
					v101 = int32(_a_F_CopySendEndOfRow_0)
					v103 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[3]))
					v104 = int32(1)
					*(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[3])) = v103 + v104
					v107 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
					*(*int32)(unsafe.Add(mBase, uint32(v92))) = v107 + v104
					v111 = int32(0)
					v113 = int32(_a_F_CopySendEndOfRow_1)
					v114 = base.AtomicRmwOr32(m, v111, v113, v111)
					*(*int64)(unsafe.Add(mBase, uint32(v92+v111)+232)) = v87
					v122 = base.AtomicRmwOr32(m, v111, v113, v111)
					v123 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
					*(*int32)(unsafe.Add(mBase, uint32(v92))) = v123 + v104
					v129 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[3]))
					*(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[3])) = v129 - v104
				}
			}
			v133 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
			v134 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v133))) = uint8(v134)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = v134
			*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = v134
			return
		}
	default:
		v85 = *(*int64)(unsafe.Add(mBase, uint32(l0)+208))
		v86 = int64(*(*int32)(unsafe.Add(mBase, uint32(v4)+4)))
		v87 = v85 + v86
		*(*int64)(unsafe.Add(mBase, uint32(l0)+208)) = v87
		v92 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[1]))
		if v92 == int32(0) {
		} else {
			v96 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[2])))
			if v96&int32(1) == int32(0) {
			} else {
				v101 = int32(_a_F_CopySendEndOfRow_0)
				v103 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[3]))
				v104 = int32(1)
				*(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[3])) = v103 + v104
				v107 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
				*(*int32)(unsafe.Add(mBase, uint32(v92))) = v107 + v104
				v111 = int32(0)
				v113 = int32(_a_F_CopySendEndOfRow_1)
				v114 = base.AtomicRmwOr32(m, v111, v113, v111)
				*(*int64)(unsafe.Add(mBase, uint32(v92+v111)+232)) = v87
				v122 = base.AtomicRmwOr32(m, v111, v113, v111)
				v123 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
				*(*int32)(unsafe.Add(mBase, uint32(v92))) = v123 + v104
				v129 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[3]))
				*(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[3])) = v129 - v104
			}
		}
		v133 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v134 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v133))) = uint8(v134)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = v134
		*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = v134
		return
	}
}
func F_collect_corrupt_items(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v381 int32
	_ = v381
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v434 int32
	_ = v434
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	v2 = l1
	v3 = l2
	v4 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(48)
	m.G0 = v23
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v4
	v28 = F_GetAccessStrategy(m, int32(1))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v33 = F_relation_open(m, l0, int32(1))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+48))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+119)))
	v38 = v36 - int32(109)
	v45 = int32(0)
	if base.B2i32(base.Ui32(int32(7)) < base.Ui32(v38))|base.B2i32(int32(1)<<(uint(v38)%32)&int32(161) == v45) == v45 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if v2 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L129
	}
L7:
	;
	v50 = F_GetStrictOldestNonRemovableTransactionId(m, v33)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	v52 = v4
	goto L9
L9:
	;
	v54 = F_palloc0(m, int32(12))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v52 = v50
	goto L9
L11:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v54))) = int64(274877906944)
	v59 = F_palloc(m, int32(384))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+8)) = v59
	v62 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v62
	v65 = F_RelationGetNumberOfBlocksInFork(m, v33, v62)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v67 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v65
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+25)) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+24)) = uint8(v3)
	v79 = F_read_stream_begin_relation(m, int32(4), v28, v33, v67, int32(_a_F_collect_corrupt_items_0), v23+int32(24), v67)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v82 = F_read_stream_next_buffer(m, v79, int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if v82 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v85 = v23 + int32(8)
	v93 = v52
	v95 = v82
	goto L19
L17:
	;
	goto L18
L18:
	;
	F_read_stream_end(m, v79)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L1
	} else {
		goto L119
	}
L19:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_collect_corrupt_items[0]))
	if v107 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L18
L21:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	F_LockBufferInternal(m, v95, int32(1))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	if v95 < int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v130)+12)))
	if v95 < int32(0) {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_collect_corrupt_items[1]))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v116+(v95^int32(-1))<<(uint(int32(2))%32))))
	v130 = v122
	goto L26
L28:
	;
	goto L29
L29:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_collect_corrupt_items[2]))
	v130 = v124 + v95<<(uint(int32(13))%32) + int32(-8192)
	goto L26
L30:
	;
	if v3 != 0 {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v135 = *(*int32)(unsafe.Add(mBase, _c_F_collect_corrupt_items[3]))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v135+(v95^int32(-1))*int32(56))+16))
	v150 = v141
	goto L30
L32:
	;
	goto L33
L33:
	;
	v143 = *(*int32)(unsafe.Add(mBase, _c_F_collect_corrupt_items[4]))
	v144 = int32(56)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v143+v95*v144-v144)+16))
	v150 = v149
	goto L30
L34:
	;
	v154 = F_visibilitymap_get_status(m, v33, v150, v23+int32(44))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	v160 = int32(0)
	goto L36
L36:
	;
	if v2 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v160 = int32(base.Ui32(v154&int32(2)) >> (uint(int32(1)) % 32))
	goto L36
L38:
	;
	v164 = F_visibilitymap_get_status(m, v33, v150, v23+int32(44))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	v168 = int32(0)
	goto L40
L40:
	;
	if v160|v168 == int32(0) {
		v434 = v93
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v168 = v164 & int32(1)
	goto L40
L42:
	;
	F_UnlockReleaseBuffer(m, v95)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L116
	}
L43:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v131) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v179 = int32(base.Ui32(v131+int32(_a_F_collect_corrupt_items_1)) >> (uint(int32(2)) % 32))
	goto L46
L45:
	;
	v179 = int32(0)
	goto L46
L46:
	;
	v181 = v179 & int32(_a_F_collect_corrupt_items_2)
	if v181 == int32(0) {
		v434 = v93
		goto L42
	} else {
		goto L47
	}
L47:
	;
	v185 = int32(base.Ui32(v150) >> (uint(int32(16)) % 32))
	v196 = v93
	v200 = int32(1)
	goto L48
L48:
	;
	v213 = v130 + int32(20) + v200&int32(_a_F_collect_corrupt_items_2)<<(uint(int32(2))%32)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	switch int32(base.Ui32(v214)>>(uint(int32(15))%32))&int32(3) - int32(1) {
	case 0:
		goto L52
	default:
		v421 = v196
		goto L50
	case 2:
		goto L53
	}
L49:
	;
	v434 = v421
	goto L42
L50:
	;
	v423 = v200 + int32(1)
	if base.Ui32(v423&int32(_a_F_collect_corrupt_items_2)) <= base.Ui32(v181) {
		v196 = v421
		v200 = v423
		goto L48
	} else {
		goto L115
	}
L51:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if base.Ui32(v393) < base.Ui32(v394) {
		goto L111
	} else {
		goto L112
	}
L52:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+12)) = uint16(v200)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+10)) = uint16(v150)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+8)) = uint16(v185)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v130 + v227&int32(_a_F_collect_corrupt_items_3)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = int32(base.Ui32(v232) >> (uint(int32(17)) % 32))
	if v168 == int32(0) {
		v355 = v196
		goto L54
	} else {
		goto L55
	}
L53:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+12)) = uint16(v200)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+10)) = uint16(v150)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+8)) = uint16(v185)
	v392 = v196
	goto L51
L54:
	;
	if v160 == int32(0) {
		v421 = v355
		goto L50
	} else {
		goto L95
	}
L55:
	;
	v241 = F_HeapTupleSatisfiesVacuum(m, v23+int32(4), v196, v95)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L57
	}
L56:
	;
	v266 = F_GetStrictOldestNonRemovableTransactionId(m, v33)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L69
	}
L57:
	;
	if v241 != int32(1) {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	v247 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v246)+20)))
	v248 = int32(768)
	if v247&v248 == v248 {
		v257 = int32(2)
		goto L60
	} else {
		goto L61
	}
L59:
	;
	if v252-v196 < int32(0) {
		v355 = v196
		goto L54
	} else {
		goto L65
	}
L60:
	;
	if base.Ui32(v196) <= base.Ui32(v257) {
		goto L56
	} else {
		goto L64
	}
L61:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v246)))
	if base.Ui32(v196) < base.Ui32(int32(3)) {
		v257 = v252
		goto L60
	} else {
		goto L62
	}
L62:
	;
	if base.Ui32(int32(2)) < base.Ui32(v252) {
		goto L59
	} else {
		goto L63
	}
L63:
	;
	v257 = v252
	goto L60
L64:
	;
	v355 = v196
	goto L54
L65:
	;
	goto L56
L66:
	;
	v349 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v85)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v347)+4)) = uint16(v349)
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	*(*int32)(unsafe.Add(mBase, uint32(v347))) = v351
	v355 = v348
	goto L54
L67:
	;
	v301 = F_HeapTupleSatisfiesVacuum(m, v23+int32(4), v266, v95)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L81
	}
L68:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if base.Ui32(v277) < base.Ui32(v278) {
		goto L76
	} else {
		goto L77
	}
L69:
	;
	if base.B2i32(base.Ui32(v196) < base.Ui32(int32(3)))|base.B2i32(base.Ui32(v266) < base.Ui32(int32(3))) == int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	if int32(0) <= v196-v266 {
		goto L68
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	if base.Ui32(v196) < base.Ui32(v266) {
		goto L67
	} else {
		goto L74
	}
L73:
	;
	goto L67
L74:
	;
	goto L68
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v291 + int32(1)
	v347 = v292 + v291*int32(6)
	v348 = v196
	goto L66
L76:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v291 = v277
	v292 = v280
	goto L75
L77:
	;
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = v278 << (uint(int32(1)) % 32)
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v287 = F_repalloc(m, v284, v278*int32(12))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+8)) = v287
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v291 = v290
	v292 = v287
	goto L75
L80:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if base.Ui32(v324) < base.Ui32(v325) {
		goto L91
	} else {
		goto L92
	}
L81:
	;
	if v301 != int32(1) {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	v307 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v306)+20)))
	v308 = int32(768)
	if v307&v308 == v308 {
		v317 = int32(2)
		goto L84
	} else {
		goto L85
	}
L83:
	;
	if int32(0) <= v312-v266 {
		goto L80
	} else {
		goto L89
	}
L84:
	;
	if base.Ui32(v266) <= base.Ui32(v317) {
		goto L80
	} else {
		goto L88
	}
L85:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v306)))
	if base.Ui32(v266) < base.Ui32(int32(3)) {
		v317 = v312
		goto L84
	} else {
		goto L86
	}
L86:
	;
	if base.Ui32(int32(2)) < base.Ui32(v312) {
		goto L83
	} else {
		goto L87
	}
L87:
	;
	v317 = v312
	goto L84
L88:
	;
	v355 = v266
	goto L54
L89:
	;
	v355 = v266
	goto L54
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v338 + int32(1)
	v347 = v339 + v338*int32(6)
	v348 = v266
	goto L66
L91:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v338 = v324
	v339 = v327
	goto L90
L92:
	;
	goto L93
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = v325 << (uint(int32(1)) % 32)
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v334 = F_repalloc(m, v331, v325*int32(12))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+8)) = v334
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v338 = v337
	v339 = v334
	goto L90
L95:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	v361 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v358)+20)))
	v362 = int32(768)
	if v361&v362 == v362 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	if v387 == int32(0) {
		v421 = v355
		goto L50
	} else {
		goto L109
	}
L97:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v358)+4))
	if v361&int32(_a_F_collect_corrupt_items_4) != 0 {
		goto L101
	} else {
		goto L102
	}
L98:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v358)))
	if base.Ui32(v366) <= base.Ui32(int32(2)) {
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v387 = int32(1)
	goto L96
L100:
	;
	if base.Ui32(v361) < base.Ui32(int32(_a_F_collect_corrupt_items_5)) {
		goto L106
	} else {
		goto L107
	}
L101:
	;
	if v370 == int32(0) {
		goto L100
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	if base.Ui32(v370) <= base.Ui32(int32(2)) {
		goto L100
	} else {
		goto L105
	}
L104:
	;
	v387 = int32(1)
	goto L96
L105:
	;
	v387 = int32(1)
	goto L96
L106:
	;
	v387 = int32(0)
	goto L96
L107:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v358)+8))
	if base.Ui32(v381) <= base.Ui32(int32(2)) {
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v387 = int32(1)
	goto L96
L109:
	;
	v392 = v355
	goto L51
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v407 + int32(1)
	v414 = v408 + v407*int32(6)
	v415 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v85)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v414)+4)) = uint16(v415)
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	*(*int32)(unsafe.Add(mBase, uint32(v414))) = v417
	v421 = v392
	goto L50
L111:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v407 = v393
	v408 = v396
	goto L110
L112:
	;
	goto L113
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = v394 << (uint(int32(1)) % 32)
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v403 = F_repalloc(m, v400, v394*int32(12))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+8)) = v403
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v407 = v406
	v408 = v403
	goto L110
L115:
	;
	goto L49
L116:
	;
	v450 = F_read_stream_next_buffer(m, v79, int32(0))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	if v450 != 0 {
		v93 = v434
		v95 = v450
		goto L19
	} else {
		goto L118
	}
L118:
	;
	goto L20
L119:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v23)+44))
	if v474 != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	F_ReleaseBuffer(m, v474)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L1
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v23)+40))
	if v477 != 0 {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	goto L122
L124:
	;
	F_ReleaseBuffer(m, v477)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L1
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	F_relation_close(m, v33, int32(1))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L128
	}
L127:
	;
	goto L126
L128:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = v483
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = int32(0)
	m.G0 = v23 + int32(48)
	return v54
L129:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v33)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v498 + int32(4)
	F_errmsg(m, int32(_a_F_collect_corrupt_items_6), v23)
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v33)+48))
	v506 = int32(*(*int8)(unsafe.Add(mBase, uint32(v505)+119)))
	F_errdetail_relkind_not_supported(m, v506)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_collect_corrupt_items_7), int32(932), int32(_a_F_collect_corrupt_items_8))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_collect_corrupt_items_read_stream_next_block(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	v5 = int32(-1)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.Ui32(v7) <= base.Ui32(v6) {
		v63 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v63
L2:
	;
	v10 = l1 + int32(16)
	goto L3
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_collect_corrupt_items_read_stream_next_block[0]))
	if v16 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v63 = v5
	goto L1
L5:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v21 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L8:
	;
	return int32(0)
L9:
	;
	goto L7
L10:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v57 = v55 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v57
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.Ui32(v57) < base.Ui32(v59) {
		goto L3
	} else {
		goto L23
	}
L11:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v50 + int32(1)
	v63 = v50
	goto L1
L12:
	;
	if v29 == int32(0) {
		goto L10
	} else {
		goto L22
	}
L13:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v42 = F_visibilitymap_get_status(m, v40, v41, v10)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L8
	} else {
		goto L20
	}
L14:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v26 = F_visibilitymap_get_status(m, v24, v25, v10)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L8
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v36 != int32(1) {
		goto L10
	} else {
		goto L19
	}
L17:
	;
	v29 = v26 & int32(2)
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v30 == int32(0) {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	v39 = int32(base.Ui32(v29) >> (uint(int32(1)) % 32))
	goto L13
L19:
	;
	v39 = int32(0)
	goto L13
L20:
	;
	if v42&int32(1)|v39 != 0 {
		goto L11
	} else {
		goto L21
	}
L21:
	;
	goto L10
L22:
	;
	goto L11
L23:
	;
	goto L4
}
func F_colorcomplement(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
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
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v148 int32
	_ = v148
	var v161 int32
	_ = v161
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v229 int32
	_ = v229
	var v239 int32
	_ = v239
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	if v12 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_colorcomplement[0]))
	if v173 != 0 {
		goto L55
	} else {
		goto L56
	}
L2:
	;
	v19 = v12
	goto L5
L3:
	;
	v59 = v11
	goto L4
L4:
	;
	v62 = int32(24)
	v66 = v11 + v10*v62 + v62
	if base.Ui32(v66) <= base.Ui32(v59) {
		goto L18
	} else {
		goto L19
	}
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v22 == int32(112) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v32 = v12
	goto L12
L7:
	;
	v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+4)))
	if v25 == int32(_a_F_colorcomplement_0) {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	if v28 != 0 {
		v19 = v28
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
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if v38 == int32(112) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v59 = v52
	goto L4
L14:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v42 = int32(*(*int16)(unsafe.Add(mBase, uint32(v32)+4)))
	v45 = v41 + v42*int32(24)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+20)) = v46 | int32(4)
	goto L16
L15:
	;
	goto L16
L16:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	if v51 != 0 {
		v32 = v51
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
	v75 = v59
	v76 = int32(0)
	goto L20
L20:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	if v79 != 0 {
		goto L18
	} else {
		goto L22
	}
L21:
	;
	goto L18
L22:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
	if v80&int32(4) != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v161 = v75 + int32(24)
	if base.Ui32(v161) < base.Ui32(v66) {
		v75 = v161
		v76 = v76 + int32(1)
		goto L20
	} else {
		goto L54
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75)+20)) = v80 & int32(-5)
	goto L23
L25:
	;
	goto L26
L26:
	;
	if v80&int32(3) != 0 {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_colorcomplement[0]))
	if v89 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	if v92 <= v93 {
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
	F_createarc(m, l0, l2, base.I32_extend16_s(v76), l4, l5)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L31
	} else {
		goto L53
	}
L34:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	if v95 == int32(0) {
		goto L33
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
	if v116 == int32(0) {
		goto L33
	} else {
		goto L45
	}
L37:
	;
	v101 = v95
	goto L38
L38:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v101)+12))
	if v107 != l5 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L33
L40:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v101)+16))
	if v115 != 0 {
		v101 = v115
		goto L38
	} else {
		goto L44
	}
L41:
	;
	v109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v101)+4)))
	if v109 != v76&int32(_a_F_colorcomplement_1) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	if v113 == l2 {
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
	v122 = v116
	goto L46
L46:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v122)+8))
	if v128 != l4 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L33
L48:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v122)+24))
	if v136 != 0 {
		v122 = v136
		goto L46
	} else {
		goto L52
	}
L49:
	;
	v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122)+4)))
	if v130 != v76&int32(_a_F_colorcomplement_1) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	if v134 == l2 {
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
	v175 = m.ExcPending
	if v175 != 0 {
		goto L31
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	if v176 <= v177 {
		goto L61
	} else {
		goto L62
	}
L58:
	;
	goto L57
L59:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v239 | int32(4)
	return
L60:
	;
	F_createarc(m, l0, int32(120), int32(0), l4, l5)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L31
	} else {
		goto L80
	}
L61:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	if v179 == int32(0) {
		goto L60
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
	if v198 == int32(0) {
		goto L60
	} else {
		goto L72
	}
L64:
	;
	v185 = v179
	goto L65
L65:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v185)+12))
	if v191 != l5 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	goto L60
L67:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v185)+16))
	if v197 != 0 {
		v185 = v197
		goto L65
	} else {
		goto L71
	}
L68:
	;
	v193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v185)+4)))
	if v193 != 0 {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v185)))
	if v194 == int32(120) {
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
	v204 = v198
	goto L73
L73:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v204)+8))
	if v210 != l4 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	goto L60
L75:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v204)+24))
	if v216 != 0 {
		v204 = v216
		goto L73
	} else {
		goto L79
	}
L76:
	;
	v212 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v204)+4)))
	if v212 != 0 {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
	if v213 == int32(120) {
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
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	v6 = int32(3)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v11 = v7 | v8<<(uint(int32(8))%32)
	if v11 <= int32(_a_F_combine_0) {
		if v11 <= int32(_a_F_combine_1) {
			switch v11 - int32(_a_F_combine_2) {
			case 0, 18:
				v104 = v6
				return v104
			case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17:
				return int32(1)
			default:
				if v11 == int32(_a_F_combine_3) {
					v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
					v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
					if v66 == v67 {
						v69 = int32(2)
					} else {
						v69 = int32(1)
					}
					return v69
				} else {
					if v11 != int32(_a_F_combine_4) {
						return int32(1)
					} else {
						v104 = v6
						return v104
					}
				}
			}
		} else {
			switch v11 - int32(_a_F_combine_5) {
			case 0, 21:
				v104 = v6
				return v104
			case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 19, 20:
				return int32(1)
			case 18:
				v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
				v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
				if v66 == v67 {
					v69 = int32(2)
				} else {
					v69 = int32(1)
				}
				return v69
			default:
				if v11 != int32(_a_F_combine_6) {
					return int32(1)
				} else {
					v104 = v6
					return v104
				}
			}
		}
	} else {
		switch v11 - int32(_a_F_combine_7) {
		case 0, 18, 38:
			v104 = v6
			return v104
		case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 19, 20, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 37:
			return int32(1)
		case 21:
			v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
			v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
			if v71 == v72 {
				return int32(2)
			} else {
				if v71 == int32(_a_F_combine_8) {
					v78 = int32(2)
					v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+20))
					v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+base.I32_extend16_s(v72)*int32(24))+20)))
					if v85&v78 != 0 {
						return int32(1)
					} else {
						v104 = v78
						return v104
					}
				} else {
					if v72 != int32(_a_F_combine_8) {
						return int32(1)
					} else {
						v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+20))
						v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93+base.I32_extend16_s(v71)*int32(24))+20)))
						if v98&int32(2) != 0 {
							v101 = int32(1)
						} else {
							v101 = int32(4)
						}
						v104 = v101
						return v104
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
					if v46&v39 != 0 {
						return int32(1)
					} else {
						v104 = v39
						return v104
					}
				} else {
					if v33 != int32(_a_F_combine_8) {
						return int32(1)
					} else {
						v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+20))
						v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54+base.I32_extend16_s(v32)*int32(24))+20)))
						if v59&int32(2) != 0 {
							v62 = int32(1)
						} else {
							v62 = int32(4)
						}
						return v62
					}
				}
			}
		default:
			switch v11 - int32(_a_F_combine_9) {
			case 0, 21:
				v104 = v6
				return v104
			case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 37:
				return int32(1)
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
						if v46&v39 != 0 {
							return int32(1)
						} else {
							v104 = v39
							return v104
						}
					} else {
						if v33 != int32(_a_F_combine_8) {
							return int32(1)
						} else {
							v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+20))
							v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54+base.I32_extend16_s(v32)*int32(24))+20)))
							if v59&int32(2) != 0 {
								v62 = int32(1)
							} else {
								v62 = int32(4)
							}
							return v62
						}
					}
				}
			case 38:
				v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
				v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
				if v71 == v72 {
					return int32(2)
				} else {
					if v71 == int32(_a_F_combine_8) {
						v78 = int32(2)
						v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+20))
						v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+base.I32_extend16_s(v72)*int32(24))+20)))
						if v85&v78 != 0 {
							return int32(1)
						} else {
							v104 = v78
							return v104
						}
					} else {
						if v72 != int32(_a_F_combine_8) {
							return int32(1)
						} else {
							v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
							v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+20))
							v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93+base.I32_extend16_s(v71)*int32(24))+20)))
							if v98&int32(2) != 0 {
								v101 = int32(1)
							} else {
								v101 = int32(4)
							}
							v104 = v101
							return v104
						}
					}
				}
			default:
				if v11 == int32(_a_F_combine_10) {
					v104 = v6
					return v104
				} else {
					return int32(1)
				}
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
	var v29 int32
	_ = v29
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
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
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v334 int32
	_ = v334
	var v343 int32
	_ = v343
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
	*(*uint16)(unsafe.Add(mBase, uint32(l2)+14)) = uint16(v343)
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
	v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v118) {
		goto L30
	} else {
		goto L31
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
	v29 = v5
	goto L9
L8:
	;
	v112 = v106 - v103
	if v112 == int32(0) {
		v343 = v102
		goto L1
	} else {
		goto L28
	}
L9:
	;
	v38 = l0 + v29*int32(6)
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
	if l1 <= v29 {
		goto L15
	} else {
		goto L16
	}
L11:
	;
	v43 = v26 - v39
	v45 = v29 + int32(1)
	if v45 != v21 {
		v26 = v43
		v29 = v45
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
	v102 = v43
	v103 = v26
	v106 = v26
	goto L8
L15:
	;
	v102 = v26
	v103 = v41
	v106 = v41
	goto L8
L16:
	;
	goto L17
L17:
	;
	v53 = v26
	v54 = v41
	v56 = v29
	v57 = v41
	goto L18
L18:
	;
	v65 = l0 + v56*int32(6)
	v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65))))
	v73 = l2 + int32(20) + (v66+int32(1))&int32(_a_F_compactify_tuples_0)<<(uint(int32(2))%32)
	v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65)+4)))
	v75 = int32(*(*int16)(unsafe.Add(mBase, uint32(v65)+2)))
	if v74+v75 == v54 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v102 = v91
	v103 = v85
	v106 = v87
	goto L8
L20:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v91 = v53 - v86
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = v88&int32(-32768) | v91&int32(_a_F_compactify_tuples_1)
	v97 = v56 + int32(1)
	if v97 != l1 {
		v53 = v91
		v54 = v85
		v56 = v97
		v57 = v87
		goto L18
	} else {
		goto L27
	}
L21:
	;
	v85 = v75
	v86 = v74
	v87 = v57
	goto L20
L22:
	;
	goto L23
L23:
	;
	v78 = v57 - v54
	if v78 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	base.MemoryCopy(m, l2+v53, l2+v54, v78)
	goto L26
L25:
	;
	goto L26
L26:
	;
	v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65)+4)))
	v83 = int32(*(*int16)(unsafe.Add(mBase, uint32(v65)+2)))
	v85 = v83
	v86 = v82
	v87 = v82 + v83
	goto L20
L27:
	;
	goto L19
L28:
	;
	base.MemoryCopy(m, l2+v102, l2+v103, v112)
	v343 = v102
	goto L1
L29:
	;
	if l1 <= v262 {
		goto L64
	} else {
		goto L65
	}
L30:
	;
	v128 = int32(base.Ui32(v118+int32(_a_F_compactify_tuples_2))>>(uint(int32(4))%32)) & int32(_a_F_compactify_tuples_3)
	goto L32
L31:
	;
	v128 = int32(0)
	goto L32
L32:
	;
	if l1 < v128 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	if int32(2) <= l1 {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	goto L35
L35:
	;
	v218 = int32(1)
	if base.Ui32(l1) <= base.Ui32(v218) {
		goto L54
	} else {
		goto L55
	}
L36:
	;
	v213 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v214 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
	v216 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+16)))
	v259 = v216
	v261 = v213 + v214
	v262 = int32(0)
	goto L29
L37:
	;
	v133 = int32(1)
	if l1 <= v133 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v180 = int32(0)
	goto L39
L39:
	;
	v192 = l0 + v180*int32(6)
	v193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v192)+4)))
	if v193 == int32(0) {
		goto L36
	} else {
		goto L53
	}
L40:
	;
	v136 = v133
	goto L42
L41:
	;
	v136 = l1
	goto L42
L42:
	;
	v145 = int32(0)
	v151 = v5
	goto L43
L43:
	;
	v157 = l0 + v145*int32(6)
	v158 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v157)+4)))
	if v158 != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	if v136&int32(1) == int32(0) {
		goto L36
	} else {
		goto L52
	}
L45:
	;
	v159 = int32(*(*int16)(unsafe.Add(mBase, uint32(v157)+2)))
	base.MemoryCopy(m, v16+v159, l2+v159, v158)
	goto L47
L46:
	;
	goto L47
L47:
	;
	v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v157)+10)))
	if v164 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v165 = int32(*(*int16)(unsafe.Add(mBase, uint32(v157)+8)))
	base.MemoryCopy(m, v16+v165, l2+v165, v164)
	goto L50
L49:
	;
	goto L50
L50:
	;
	v170 = int32(2)
	v171 = v145 + v170
	v173 = v151 + v170
	if v173 != v136&int32(2147483646) {
		v145 = v171
		v151 = v173
		goto L43
	} else {
		goto L51
	}
L51:
	;
	goto L44
L52:
	;
	v180 = v171
	goto L39
L53:
	;
	v196 = int32(*(*int16)(unsafe.Add(mBase, uint32(v192)+2)))
	base.MemoryCopy(m, v16+v196, l2+v196, v193)
	goto L36
L54:
	;
	v221 = v218
	goto L56
L55:
	;
	v221 = l1
	goto L56
L56:
	;
	v222 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+16)))
	v226 = v222
	v229 = v5
	goto L58
L57:
	;
	v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+14)))
	v250 = v247 - v249
	if v250 == int32(0) {
		v259 = v247
		v261 = v241
		v262 = v248
		goto L29
	} else {
		goto L62
	}
L58:
	;
	v238 = l0 + v229*int32(6)
	v239 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v238)+4)))
	v240 = int32(*(*int16)(unsafe.Add(mBase, uint32(v238)+2)))
	v241 = v239 + v240
	if v226 != v241 {
		v247 = v226
		v248 = v229
		goto L57
	} else {
		goto L60
	}
L59:
	;
	v247 = v243
	v248 = v221
	goto L57
L60:
	;
	v243 = v226 - v239
	v245 = v229 + int32(1)
	if v245 != v221 {
		v226 = v243
		v229 = v245
		goto L58
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	base.MemoryCopy(m, v249+v16, l2+v249, v250)
	v259 = v247
	v261 = v241
	v262 = v248
	goto L29
L63:
	;
	v334 = v326 - v325
	if v334 == int32(0) {
		v343 = v324
		goto L1
	} else {
		goto L77
	}
L64:
	;
	v324 = v259
	v325 = v261
	v326 = v261
	goto L63
L65:
	;
	goto L66
L66:
	;
	v275 = v259
	v276 = v261
	v277 = v261
	v278 = v262
	goto L67
L67:
	;
	v287 = l0 + v278*int32(6)
	v288 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v287))))
	v295 = l2 + int32(20) + (v288+int32(1))&int32(_a_F_compactify_tuples_0)<<(uint(int32(2))%32)
	v296 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v287)+4)))
	v297 = int32(*(*int16)(unsafe.Add(mBase, uint32(v287)+2)))
	if v296+v297 == v276 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	v324 = v313
	v325 = v307
	v326 = v308
	goto L63
L69:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v295)))
	v313 = v275 - v309
	*(*int32)(unsafe.Add(mBase, uint32(v295))) = v310&int32(-32768) | v313&int32(_a_F_compactify_tuples_1)
	v319 = v278 + int32(1)
	if v319 != l1 {
		v275 = v313
		v276 = v307
		v277 = v308
		v278 = v319
		goto L67
	} else {
		goto L76
	}
L70:
	;
	v307 = v297
	v308 = v277
	v309 = v296
	goto L69
L71:
	;
	goto L72
L72:
	;
	v300 = v277 - v276
	if v300 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	base.MemoryCopy(m, l2+v275, v276+v16, v300)
	goto L75
L74:
	;
	goto L75
L75:
	;
	v304 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v287)+4)))
	v305 = int32(*(*int16)(unsafe.Add(mBase, uint32(v287)+2)))
	v307 = v305
	v308 = v304 + v305
	v309 = v304
	goto L69
L76:
	;
	goto L68
L77:
	;
	base.MemoryCopy(m, l2+v324, v325+v16, v334)
	v343 = v324
	goto L1
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
func F_comparison_shim(m *base.Module, l0 int64, l1 int64, l2 int32) int32 {
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
	var v18 int64
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
	*(*int64)(unsafe.Add(mBase, uint32(v9)+72)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = l0
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+48)) = uint8(v4)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v18 = m.T0[v17].(func(*base.Module, int32) int64)(m, v9+int32(32))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+48)))
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
			return base.I32_wrap_i64(v18)
		}
	}
}
func F_compress_init(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+64)))
	if base.Ui32(v5-int32(3)) < base.Ui32(int32(-2)) {
		return int32(-102)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
		v15 = m.Env.Pgmem_deflate_create(m, base.B2i32(v5 == int32(1)), v14)
		mBase = m.M
		if v15 <= int32(0) {
			return int32(-105)
		} else {
			v21 = F_palloc0(m, int32(_a_F_compress_init_0))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = int32(_a_F_compress_init_1)
				v28 = *(*int32)(unsafe.Add(mBase, _c_F_compress_init[0]))
				F_ResourceOwnerEnlarge(m, v28)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, _c_F_compress_init[1]))
					v34 = F_MemoryContextAlloc(m, v32, int32(8))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v34))) = v15
						v38 = *(*int32)(unsafe.Add(mBase, _c_F_compress_init[0]))
						*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v38
						F_ResourceOwnerRemember(m, v38, base.I64_extend_i32_u(v34), int32(_a_F_compress_init_2))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v21))) = v34
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v21
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
	var v33 float64
	_ = v33
	var v37 float64
	_ = v37
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
	if base.F64_gt(v24, v23)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v24)&int64(9223372036854775807))) != 0 {
		v37 = v23
	} else {
		v33 = float64(1)
		if base.F64_le(v24, v33) != 0 {
			v37 = v33
		} else {
			v37 = base.F64_nearest(v24)
		}
	}
	return v37
}
func F_connectby_text(m *base.Module, l0 int32) int64 {
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
	var v56 int64
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
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v21 = F_pg_detoast_datum_packed(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int64(0)
	} else {
		v25 = F_text_to_cstring(m, v21)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int64(0)
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			v28 = F_pg_detoast_datum_packed(m, v27)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int64(0)
			} else {
				v30 = F_text_to_cstring(m, v28)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int64(0)
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
					v33 = F_pg_detoast_datum_packed(m, v32)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int64(0)
					} else {
						v35 = F_text_to_cstring(m, v33)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int64(0)
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
							v38 = F_pg_detoast_datum_packed(m, v37)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int64(0)
							} else {
								v40 = F_text_to_cstring(m, v38)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int64(0)
								} else {
									v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									if v42 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v130 = m.ExcPending
										if v130 != 0 {
											return int64(0)
										} else {
											F_errcode(m, int32(1088))
											mBase = m.M
											v133 = m.ExcPending
											if v133 != 0 {
												return int64(0)
											} else {
												F_errmsg(m, int32(_a_F_connectby_text_0), int32(0))
												mBase = m.M
												v137 = m.ExcPending
												if v137 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(_a_F_connectby_text_1), int32(999), int32(_a_F_connectby_text_2))
													mBase = m.M
													v142 = m.ExcPending
													if v142 != 0 {
														return int64(0)
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
										if v45 != int32(389) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v130 = m.ExcPending
											if v130 != 0 {
												return int64(0)
											} else {
												F_errcode(m, int32(1088))
												mBase = m.M
												v133 = m.ExcPending
												if v133 != 0 {
													return int64(0)
												} else {
													F_errmsg(m, int32(_a_F_connectby_text_0), int32(0))
													mBase = m.M
													v137 = m.ExcPending
													if v137 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_connectby_text_1), int32(999), int32(_a_F_connectby_text_2))
														mBase = m.M
														v142 = m.ExcPending
														if v142 != 0 {
															return int64(0)
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
												v146 = m.ExcPending
												if v146 != 0 {
													return int64(0)
												} else {
													F_errcode(m, int32(1088))
													mBase = m.M
													v149 = m.ExcPending
													if v149 != 0 {
														return int64(0)
													} else {
														F_errmsg(m, int32(_a_F_connectby_text_3), int32(0))
														mBase = m.M
														v153 = m.ExcPending
														if v153 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(_a_F_connectby_text_1), int32(1004), int32(_a_F_connectby_text_2))
															mBase = m.M
															v158 = m.ExcPending
															if v158 != 0 {
																return int64(0)
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
													v146 = m.ExcPending
													if v146 != 0 {
														return int64(0)
													} else {
														F_errcode(m, int32(1088))
														mBase = m.M
														v149 = m.ExcPending
														if v149 != 0 {
															return int64(0)
														} else {
															F_errmsg(m, int32(_a_F_connectby_text_3), int32(0))
															mBase = m.M
															v153 = m.ExcPending
															if v153 != 0 {
																return int64(0)
															} else {
																F_errfinish(m, int32(_a_F_connectby_text_1), int32(1004), int32(_a_F_connectby_text_2))
																mBase = m.M
																v158 = m.ExcPending
																if v158 != 0 {
																	return int64(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													}
												} else {
													v56 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
													v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
													if v57 == int32(6) {
														v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
														v61 = F_pg_detoast_datum_packed(m, v60)
														mBase = m.M
														v62 = m.ExcPending
														if v62 != 0 {
															return int64(0)
														} else {
															v63 = F_text_to_cstring(m, v61)
															mBase = m.M
															v64 = m.ExcPending
															if v64 != 0 {
																return int64(0)
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
																	return int64(0)
																} else {
																	v79 = base.B2i32(v57 == int32(6))
																	F_validateConnectbyTupleDesc(m, v76, v79, int32(0))
																	mBase = m.M
																	v82 = m.ExcPending
																	if v82 != 0 {
																		return int64(0)
																	} else {
																		v83 = F_TupleDescGetAttInMetadata(m, v76)
																		mBase = m.M
																		v84 = m.ExcPending
																		if v84 != 0 {
																			return int64(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v42)+16)) = int32(2)
																			v87 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
																			*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = int32(1)
																			F_SPI_connect_ext(m, int32(0))
																			mBase = m.M
																			v92 = m.ExcPending
																			if v92 != 0 {
																				return int64(0)
																			} else {
																				v93 = int32(_a_F_connectby_text_4)
																				v94 = *(*int32)(unsafe.Add(mBase, _c_F_connectby_text[0]))
																				*(*int32)(unsafe.Add(mBase, _c_F_connectby_text[0])) = v73
																				v103 = *(*int32)(unsafe.Add(mBase, _c_F_connectby_text[1]))
																				v104 = F_tuplestore_begin_heap(m, int32(base.Ui32(v87&int32(4))>>(uint(int32(2))%32)), int32(0), v103)
																				mBase = m.M
																				v105 = m.ExcPending
																				if v105 != 0 {
																					return int64(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, _c_F_connectby_text[0])) = v94
																					v108 = int32(0)
																					F_build_tuplestore_recursively(m, v30, v35, v25, v108, v68, v40, v40, v108, v18+int32(12), base.I32_wrap_i64(v56), v79, v108, v83, v104)
																					mBase = m.M
																					v115 = m.ExcPending
																					if v115 != 0 {
																						return int64(0)
																					} else {
																						v116 = F_SPI_finish(m)
																						mBase = m.M
																						v117 = m.ExcPending
																						if v117 != 0 {
																							return int64(0)
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(v42)+28)) = v76
																							*(*int32)(unsafe.Add(mBase, uint32(v42)+24)) = v104
																							*(*int32)(unsafe.Add(mBase, _c_F_connectby_text[0])) = v70
																							m.G0 = v18 + int32(16)
																							return int64(0)
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
															return int64(0)
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
																return int64(0)
															} else {
																v79 = base.B2i32(v57 == int32(6))
																F_validateConnectbyTupleDesc(m, v76, v79, int32(0))
																mBase = m.M
																v82 = m.ExcPending
																if v82 != 0 {
																	return int64(0)
																} else {
																	v83 = F_TupleDescGetAttInMetadata(m, v76)
																	mBase = m.M
																	v84 = m.ExcPending
																	if v84 != 0 {
																		return int64(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v42)+16)) = int32(2)
																		v87 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
																		*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = int32(1)
																		F_SPI_connect_ext(m, int32(0))
																		mBase = m.M
																		v92 = m.ExcPending
																		if v92 != 0 {
																			return int64(0)
																		} else {
																			v93 = int32(_a_F_connectby_text_4)
																			v94 = *(*int32)(unsafe.Add(mBase, _c_F_connectby_text[0]))
																			*(*int32)(unsafe.Add(mBase, _c_F_connectby_text[0])) = v73
																			v103 = *(*int32)(unsafe.Add(mBase, _c_F_connectby_text[1]))
																			v104 = F_tuplestore_begin_heap(m, int32(base.Ui32(v87&int32(4))>>(uint(int32(2))%32)), int32(0), v103)
																			mBase = m.M
																			v105 = m.ExcPending
																			if v105 != 0 {
																				return int64(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, _c_F_connectby_text[0])) = v94
																				v108 = int32(0)
																				F_build_tuplestore_recursively(m, v30, v35, v25, v108, v68, v40, v40, v108, v18+int32(12), base.I32_wrap_i64(v56), v79, v108, v83, v104)
																				mBase = m.M
																				v115 = m.ExcPending
																				if v115 != 0 {
																					return int64(0)
																				} else {
																					v116 = F_SPI_finish(m)
																					mBase = m.M
																					v117 = m.ExcPending
																					if v117 != 0 {
																						return int64(0)
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v42)+28)) = v76
																						*(*int32)(unsafe.Add(mBase, uint32(v42)+24)) = v104
																						*(*int32)(unsafe.Add(mBase, _c_F_connectby_text[0])) = v70
																						m.G0 = v18 + int32(16)
																						return int64(0)
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
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
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
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 float64
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 float64
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v225 int32
	_ = v225
	var v246 int32
	_ = v246
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v361 int32
	_ = v361
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v409 int32
	_ = v409
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v491 int32
	_ = v491
	var v497 int32
	_ = v497
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v535 int32
	_ = v535
	var v540 int32
	_ = v540
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v582 int32
	_ = v582
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v621 int32
	_ = v621
	var v650 float64
	_ = v650
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v766 float64
	_ = v766
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v777 float64
	_ = v777
	var v782 int32
	_ = v782
	var v785 int32
	_ = v785
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v798 float64
	_ = v798
	var v799 float64
	_ = v799
	var v802 float64
	_ = v802
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v811 float64
	_ = v811
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v845 int32
	_ = v845
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v854 float64
	_ = v854
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v866 float64
	_ = v866
	var v868 float64
	_ = v868
	var v873 int32
	_ = v873
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v888 int32
	_ = v888
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v927 int32
	_ = v927
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v942 int32
	_ = v942
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v982 int32
	_ = v982
	var v1030 int32
	_ = v1030
	var v1053 int32
	_ = v1053
	var v1059 int32
	_ = v1059
	var v1087 int32
	_ = v1087
	var v1089 int32
	_ = v1089
	var v1090 float64
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1094 int32
	_ = v1094
	var v1095 float64
	_ = v1095
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1106 int32
	_ = v1106
	var v1109 int32
	_ = v1109
	var v1112 int32
	_ = v1112
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1150 int32
	_ = v1150
	var v1175 int32
	_ = v1175
	var v1178 int32
	_ = v1178
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1195 int32
	_ = v1195
	var v1249 int32
	_ = v1249
	var v1261 int32
	_ = v1261
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1286 float64
	_ = v1286
	var v1323 int32
	_ = v1323
	var v1358 int32
	_ = v1358
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1362 int32
	_ = v1362
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1377 int32
	_ = v1377
	var v1382 int32
	_ = v1382
	var v1384 int32
	_ = v1384
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1403 int32
	_ = v1403
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1411 int32
	_ = v1411
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1420 int32
	_ = v1420
	var v1428 int32
	_ = v1428
	var v1432 int32
	_ = v1432
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1456 int32
	_ = v1456
	var v1460 int32
	_ = v1460
	var v1461 int32
	_ = v1461
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1471 int32
	_ = v1471
	var v1472 int32
	_ = v1472
	var v1475 int32
	_ = v1475
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1501 int32
	_ = v1501
	var v1506 int32
	_ = v1506
	var v1508 int32
	_ = v1508
	var v1516 int32
	_ = v1516
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1544 int32
	_ = v1544
	var v1554 int32
	_ = v1554
	var v1562 int32
	_ = v1562
	var v1566 int32
	_ = v1566
	var v1568 int32
	_ = v1568
	var v1569 int32
	_ = v1569
	var v1572 int32
	_ = v1572
	var v1573 int32
	_ = v1573
	var v1574 int32
	_ = v1574
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1584 int32
	_ = v1584
	var v1585 int32
	_ = v1585
	var v1588 int32
	_ = v1588
	var v1589 int32
	_ = v1589
	var v1600 int32
	_ = v1600
	var v1624 int32
	_ = v1624
	var v1625 int32
	_ = v1625
	var v1628 int32
	_ = v1628
	var v1629 int32
	_ = v1629
	var v1635 int32
	_ = v1635
	var v1636 int32
	_ = v1636
	var v1673 int32
	_ = v1673
	var v1675 int32
	_ = v1675
	var v1682 int32
	_ = v1682
	var v1688 int32
	_ = v1688
	var v1710 int32
	_ = v1710
	var v1711 int32
	_ = v1711
	var v1715 int32
	_ = v1715
	var v1716 int32
	_ = v1716
	var v1719 int32
	_ = v1719
	var v1720 int32
	_ = v1720
	var v1726 int32
	_ = v1726
	var v1731 int32
	_ = v1731
	var v1755 int32
	_ = v1755
	var v1756 int32
	_ = v1756
	var v1759 int32
	_ = v1759
	var v1763 int32
	_ = v1763
	var v1764 int32
	_ = v1764
	var v1765 int32
	_ = v1765
	var v1767 int32
	_ = v1767
	var v1768 int32
	_ = v1768
	var v1773 int32
	_ = v1773
	var v1802 int32
	_ = v1802
	var v1803 int32
	_ = v1803
	var v1805 int32
	_ = v1805
	var v1806 int32
	_ = v1806
	var v1812 int32
	_ = v1812
	var v1841 float64
	_ = v1841
	var v1842 int32
	_ = v1842
	var v1845 int32
	_ = v1845
	var v1846 int32
	_ = v1846
	var v1848 int32
	_ = v1848
	var v1849 int32
	_ = v1849
	var v1865 int32
	_ = v1865
	var v1885 int32
	_ = v1885
	var v1887 int32
	_ = v1887
	var v1888 int32
	_ = v1888
	var v1890 int32
	_ = v1890
	var v1891 int32
	_ = v1891
	var v1892 int32
	_ = v1892
	var v1893 int32
	_ = v1893
	var v1896 int32
	_ = v1896
	var v1897 int32
	_ = v1897
	var v1908 int32
	_ = v1908
	var v1913 int32
	_ = v1913
	var v1923 int32
	_ = v1923
	var v1924 int32
	_ = v1924
	var v1926 int32
	_ = v1926
	var v1927 int32
	_ = v1927
	var v1928 int32
	_ = v1928
	var v1946 int32
	_ = v1946
	var v1961 int32
	_ = v1961
	var v1963 int32
	_ = v1963
	var v1981 int32
	_ = v1981
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
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if l3 == int32(0) {
		goto L7
	} else {
		goto L8
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
	v51 = base.I32_trunc_sat_f64_u(v50)
	goto L1
L5:
	;
	m.G0 = v1981 + int32(32)
	return
L6:
	;
	F_add_path(m, l1, v1961)
	mBase = m.M
	v1963 = m.ExcPending
	if v1963 != 0 {
		goto L43
	} else {
		goto L261
	}
L7:
	;
	v55 = int32(0)
	if v52 == v55 {
		v134 = v9
		v135 = v31
		v136 = v55
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	if v52 == int32(0) {
		v1981 = v35
		goto L5
	} else {
		goto L115
	}
L10:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+348))
	if v138 != 0 {
		goto L36
	} else {
		goto L37
	}
L11:
	;
	v58 = int32(0)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	if v59 == v58 {
		v134 = v9
		v135 = v31
		v136 = v58
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l2)+64))
	if v62 == v63 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v116 == int32(0) {
		v134 = v9
		v135 = v31
		v136 = v59
		goto L10
	} else {
		goto L31
	}
L14:
	;
	v116 = int32(1)
	goto L13
L15:
	;
	goto L16
L16:
	;
	v72 = int32(0)
	goto L18
L17:
	;
	v116 = v108
	goto L13
L18:
	;
	v76 = int32(0)
	if v62 == v76 {
		v86 = v76
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v108 = int32(0)
	goto L17
L20:
	;
	if v63 != 0 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	if v80 <= v72 {
		v86 = int32(0)
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	v86 = v82 + v72<<(uint(int32(2))%32)
	goto L20
L23:
	;
	v92 = base.B2i32(v86 == int32(0))
	if v86 == int32(0) {
		v108 = v92
		goto L17
	} else {
		goto L28
	}
L24:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if v72 < v87 {
		goto L23
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v116 = base.B2i32(v86 == int32(0))
	goto L13
L27:
	;
	goto L26
L28:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	if v95 == int32(0) {
		v108 = v92
		goto L17
	} else {
		goto L29
	}
L29:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v72<<(uint(int32(2))%32)+v95)))
	if v102 == v104 {
		v72 = v72 + int32(1)
		goto L18
	} else {
		goto L30
	}
L30:
	;
	goto L19
L31:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v120 = *(*float64)(unsafe.Add(mBase, uint32(v119)+16))
	v122 = v59 + int32(4)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	if base.Ui32(v122) < base.Ui32(v125+v126<<(uint(int32(2))%32)) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v131 = v122
	goto L34
L33:
	;
	v131 = int32(0)
	goto L34
L34:
	;
	v134 = v119
	v135 = v120
	v136 = v131
	goto L10
L35:
	;
	if base.F64_gt(base.F64_mul(base.F64_sub(l7, v135), base.F64_convert_i32_u(v145)), base.F64_convert_i32_u(v51)) != 0 {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	v141 = v139
	goto L38
L37:
	;
	v141 = int32(0)
	goto L38
L38:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+32))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l6)+32))
	v145 = F_hash_agg_entry_size(m, v141, v143, v144)
	mBase = m.M
	goto L35
L39:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v150 != 0 {
		v1981 = v35
		goto L5
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l5)+28))
	v152 = F_list_copy(m, v151)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L41
L43:
	;
	return
L44:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v136 != 0 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	if v225 == int32(0) {
		v1981 = v35
		goto L5
	} else {
		goto L57
	}
L46:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v154)+4))
	if v163 <= v162 {
		v225 = v152
		goto L45
	} else {
		goto L51
	}
L47:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)+12))
	v162 = (v136 - v155) >> (uint(int32(2)) % 32)
	goto L46
L48:
	;
	goto L49
L49:
	;
	if v154 == int32(0) {
		v225 = v152
		goto L45
	} else {
		goto L50
	}
L50:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v154)+4))
	v162 = v161
	goto L46
L51:
	;
	v173 = v162
	v178 = v152
	goto L52
L52:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v154)+12))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v197+v173<<(uint(int32(2))%32))))
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+24)))
	if v202 != int32(1) {
		v1981 = v35
		goto L5
	} else {
		goto L54
	}
L53:
	;
	v225 = v206
	goto L45
L54:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v201)+12))
	v206 = F_list_concat(m, v178, v205)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L43
	} else {
		goto L55
	}
L55:
	;
	v209 = v173 + int32(1)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v154)+4))
	if v209 < v210 {
		v173 = v209
		v178 = v206
		goto L52
	} else {
		goto L56
	}
L56:
	;
	goto L53
L57:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v225)+4))
	if int32(0) < v246 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v263 = v9
	v264 = v9
	v267 = v9
	v269 = v9
	goto L61
L59:
	;
	v707 = v9
	v710 = v9
	v712 = v9
	goto L60
L60:
	;
	if v707 == int32(0) {
		v1981 = v35
		goto L5
	} else {
		goto L104
	}
L61:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v225)+12))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v281+v263<<(uint(int32(2))%32))))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v285)+4))
	if v286 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	v707 = v671
	v710 = v674
	v712 = v676
	goto L60
L63:
	;
	v689 = v263 + int32(1)
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v225)+4))
	if v689 < v690 {
		v263 = v689
		v264 = v671
		v267 = v674
		v269 = v676
		goto L61
	} else {
		goto L103
	}
L64:
	;
	v289 = F_lappend(m, v269, v285)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L43
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v295 = F_palloc0(m, int32(32))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L43
	} else {
		goto L69
	}
L67:
	;
	v292 = F_lappend(m, v267, int32(0))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L43
	} else {
		goto L68
	}
L68:
	;
	v671 = v264
	v674 = v292
	v676 = v289
	goto L63
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v295))) = int32(311)
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v286)+4))
	if v299 <= int32(0) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v295)+4)) = v361
	*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v285
	*(*int32)(unsafe.Add(mBase, uint32(v35)+28)) = v285
	v390 = F_list_make1_impl(m, int32(1), v35+int32(16))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L43
	} else {
		goto L79
	}
L71:
	;
	v361 = int32(0)
	goto L70
L72:
	;
	goto L73
L73:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v304 = int32(0)
	v314 = v304
	v315 = v304
	goto L74
L74:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v286)+12))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v338+v314<<(uint(int32(2))%32))))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v303)+100))
	v344 = F_get_sortgroupref_clause(m, v342, v343)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L43
	} else {
		goto L76
	}
L75:
	;
	v361 = v346
	goto L70
L76:
	;
	v346 = F_lappend(m, v315, v344)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L43
	} else {
		goto L77
	}
L77:
	;
	v349 = v314 + int32(1)
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v286)+4))
	if v349 < v350 {
		v314 = v349
		v315 = v346
		goto L74
	} else {
		goto L78
	}
L78:
	;
	goto L75
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v295)+12)) = v390
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l5)+32))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v295)+4))
	if v394 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	if v390 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L81:
	;
	v397 = int32(0)
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v394)+4))
	if v398 <= v397 {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v409 = v397
	goto L83
L83:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v394)+12))
	v434 = int32(2)
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v433+v409<<(uint(v434)%32))))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v437)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v393+v438<<(uint(v434)%32)))) = v409
	v444 = v409 + int32(1)
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v394)+4))
	if v444 < v445 {
		v409 = v444
		goto L83
	} else {
		goto L85
	}
L84:
	;
	goto L80
L85:
	;
	goto L84
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v295)+8)) = v621
	v650 = *(*float64)(unsafe.Add(mBase, uint32(v285)+8))
	v651 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v295)+24)) = uint16(v651)
	*(*float64)(unsafe.Add(mBase, uint32(v295)+16)) = v650
	v654 = F_lappend(m, v264, v295)
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L43
	} else {
		goto L102
	}
L87:
	;
	v621 = int32(0)
	goto L86
L88:
	;
	goto L89
L89:
	;
	v482 = int32(0)
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v390)+4))
	if v484 <= v482 {
		v621 = v482
		goto L86
	} else {
		goto L90
	}
L90:
	;
	v491 = v482
	v497 = v482
	goto L91
L91:
	;
	v519 = int32(0)
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v390)+12))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v520+v497<<(uint(int32(2))%32))))
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v524)+4))
	if v525 == v519 {
		v582 = v519
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v621 = v611
	goto L86
L93:
	;
	v611 = F_lappend(m, v491, v582)
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L43
	} else {
		goto L100
	}
L94:
	;
	v528 = int32(0)
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v525)+4))
	if v529 <= v528 {
		v582 = v519
		goto L93
	} else {
		goto L95
	}
L95:
	;
	v535 = v519
	v540 = v528
	goto L96
L96:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v525)+12))
	v565 = int32(2)
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v564+v540<<(uint(v565)%32))))
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v393+v568<<(uint(v565)%32))))
	v573 = F_lappend_int(m, v535, v572)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L43
	} else {
		goto L98
	}
L97:
	;
	v582 = v573
	goto L93
L98:
	;
	v576 = v540 + int32(1)
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v525)+4))
	if v576 < v577 {
		v535 = v573
		v540 = v576
		goto L96
	} else {
		goto L99
	}
L99:
	;
	goto L97
L100:
	;
	v614 = v497 + int32(1)
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v390)+4))
	if v614 < v615 {
		v491 = v611
		v497 = v614
		goto L91
	} else {
		goto L101
	}
L101:
	;
	goto L92
L102:
	;
	v671 = v654
	v674 = v267
	v676 = v269
	goto L63
L103:
	;
	goto L62
L104:
	;
	if v134 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v37)+112))
	v753 = F_create_groupingsets_path(m, l0, l1, l2, v752, v749, v751, l6)
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L43
	} else {
		goto L114
	}
L106:
	;
	if v710 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L107:
	;
	v744 = v134
	goto L108
L108:
	;
	v746 = F_lappend(m, v707, v744)
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L43
	} else {
		goto L113
	}
L109:
	;
	v749 = int32(2)
	v751 = v707
	goto L105
L110:
	;
	goto L111
L111:
	;
	v732 = F_palloc0(m, int32(32))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L43
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v732)+12)) = v712
	*(*int64)(unsafe.Add(mBase, uint32(v732))) = int64(311)
	*(*int32)(unsafe.Add(mBase, uint32(v732)+8)) = v710
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v710)+4))
	v739 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v732)+24)) = uint16(v739)
	*(*float64)(unsafe.Add(mBase, uint32(v732)+16)) = base.F64_convert_i32_s(v738)
	v744 = v732
	goto L108
L113:
	;
	v749 = int32(3)
	v751 = v746
	goto L105
L114:
	;
	v1946 = v35
	v1961 = v753
	goto L6
L115:
	;
	if l4 == int32(0) {
		v1891 = l0
		v1892 = l1
		v1893 = l2
		v1896 = l5
		v1897 = l6
		v1908 = v35
		v1913 = v37
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v1923 = *(*int32)(unsafe.Add(mBase, uint32(v1896)+28))
	if v1923 != 0 {
		v1981 = v1908
		goto L5
	} else {
		goto L259
	}
L117:
	;
	v759 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+16)))
	if v759 != int32(1) {
		v1891 = l0
		v1892 = l1
		v1893 = l2
		v1896 = l5
		v1897 = l6
		v1908 = v35
		v1913 = v37
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v762 = *(*int32)(unsafe.Add(mBase, uint32(l5)+28))
	v763 = F_list_copy(m, v762)
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L43
	} else {
		goto L119
	}
L119:
	;
	v766 = *(*float64)(unsafe.Add(mBase, uint32(l5)+8))
	v767 = *(*int32)(unsafe.Add(mBase, uint32(l0)+348))
	if v767 != 0 {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	v1516 = int32(0)
	if v1498|base.B2i32(v1499 == v1516) == v1516 {
		goto L221
	} else {
		goto L222
	}
L121:
	;
	v777 = base.F64_sub(base.F64_convert_i32_u(v51), base.F64_mul(v766, base.F64_convert_i32_u(v774)))
	if base.F64_gt(v777, float64(0)) == int32(0) {
		v1484 = l0
		v1485 = l1
		v1486 = l2
		v1489 = l5
		v1490 = l6
		v1498 = v9
		v1499 = v763
		v1501 = v35
		v1506 = v37
		v1508 = v9
		goto L120
	} else {
		goto L125
	}
L122:
	;
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v767)+4))
	v770 = v768
	goto L124
L123:
	;
	v770 = int32(0)
	goto L124
L124:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v771)+32))
	v773 = *(*int32)(unsafe.Add(mBase, uint32(l6)+32))
	v774 = F_hash_agg_entry_size(m, v770, v772, v773)
	mBase = m.M
	goto L121
L125:
	;
	v782 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v782 == int32(0) {
		v1484 = l0
		v1485 = l1
		v1486 = l2
		v1489 = l5
		v1490 = l6
		v1498 = v9
		v1499 = v763
		v1501 = v35
		v1506 = v37
		v1508 = v9
		goto L120
	} else {
		goto L126
	}
L126:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v782)+4))
	if v785 < int32(2) {
		v1484 = l0
		v1485 = l1
		v1486 = l2
		v1489 = l5
		v1490 = l6
		v1498 = v9
		v1499 = v763
		v1501 = v35
		v1506 = v37
		v1508 = v9
		goto L120
	} else {
		goto L127
	}
L127:
	;
	v790 = F_palloc(m, v785<<(uint(int32(2))%32))
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L43
	} else {
		goto L128
	}
L128:
	;
	v792 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v792 == int32(0) {
		v1484 = l0
		v1485 = l1
		v1486 = l2
		v1489 = l5
		v1490 = l6
		v1498 = v9
		v1499 = v763
		v1501 = v35
		v1506 = v37
		v1508 = v9
		goto L120
	} else {
		goto L129
	}
L129:
	;
	v798 = base.F64_div(v777, base.F64_mul(base.F64_convert_i32_u(v785), float64(20)))
	v799 = float64(1)
	if base.F64_gt(v798, v799) != 0 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v802 = v798
	goto L132
L131:
	;
	v802 = v799
	goto L132
L132:
	;
	v805 = base.I32_trunc_sat_f64_s(base.F64_floor(base.F64_div(v777, v802)))
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v792)+4))
	if int32(2) <= v806 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v811 = base.F64_add(base.F64_convert_i32_s(v805), float64(1))
	v821 = int32(1)
	v822 = v9
	goto L136
L134:
	;
	v888 = v9
	goto L135
L135:
	;
	if v888 <= int32(0) {
		v1484 = l0
		v1485 = l1
		v1486 = l2
		v1489 = l5
		v1490 = l6
		v1498 = v9
		v1499 = v763
		v1501 = v35
		v1506 = v37
		v1508 = v9
		goto L120
	} else {
		goto L149
	}
L136:
	;
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v792)+12))
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v845+v821<<(uint(int32(2))%32))))
	v850 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v849)+24)))
	if v850 != 0 {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	v888 = v873
	goto L135
L138:
	;
	v854 = *(*float64)(unsafe.Add(mBase, uint32(v849)+16))
	v855 = *(*int32)(unsafe.Add(mBase, uint32(l0)+348))
	if v855 != 0 {
		goto L142
	} else {
		goto L143
	}
L139:
	;
	v873 = v822
	goto L140
L140:
	;
	v876 = v821 + int32(1)
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v792)+4))
	if v876 < v877 {
		v821 = v876
		v822 = v873
		goto L136
	} else {
		goto L148
	}
L141:
	;
	v866 = base.F64_floor(base.F64_div(base.F64_mul(v854, base.F64_convert_i32_u(v862)), v802))
	if base.F64_gt(v811, v866) != 0 {
		goto L145
	} else {
		goto L146
	}
L142:
	;
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v855)+4))
	v858 = v856
	goto L144
L143:
	;
	v858 = int32(0)
	goto L144
L144:
	;
	v859 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v859)+32))
	v861 = *(*int32)(unsafe.Add(mBase, uint32(l6)+32))
	v862 = F_hash_agg_entry_size(m, v858, v860, v861)
	mBase = m.M
	goto L141
L145:
	;
	v868 = v866
	goto L147
L146:
	;
	v868 = v811
	goto L147
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v790+v822<<(uint(int32(2))%32)))) = base.I32_trunc_sat_f64_s(v868)
	v873 = v822 + int32(1)
	goto L140
L148:
	;
	goto L137
L149:
	;
	v913 = int32(0)
	v915 = *(*int32)(unsafe.Add(mBase, _c_F_consider_groupingsets_paths[2]))
	v920 = F_AllocSetContextCreateInternal(m, v915, int32(_a_F_consider_groupingsets_paths_0), v913, int32(1024), int32(_a_F_consider_groupingsets_paths_1))
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L43
	} else {
		goto L150
	}
L150:
	;
	v922 = int32(_a_F_consider_groupingsets_paths_2)
	v923 = *(*int32)(unsafe.Add(mBase, _c_F_consider_groupingsets_paths[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_consider_groupingsets_paths[2])) = v920
	v927 = v805 + int32(1)
	v930 = F_palloc(m, v927<<(uint(int32(3))%32))
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L43
	} else {
		goto L151
	}
L151:
	;
	v934 = F_palloc(m, v927<<(uint(int32(2))%32))
	mBase = m.M
	v935 = m.ExcPending
	if v935 != 0 {
		goto L43
	} else {
		goto L152
	}
L152:
	;
	if int32(0) <= v805 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v942 = v913
	goto L156
L154:
	;
	goto L155
L155:
	;
	if int32(0) < v888 {
		goto L160
	} else {
		goto L161
	}
L156:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v930+v942<<(uint(int32(3))%32)))) = int64(0)
	v978 = F_bms_make_singleton(m, v888)
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L43
	} else {
		goto L158
	}
L157:
	;
	goto L155
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v934+v942<<(uint(int32(2))%32)))) = v978
	v982 = v942 + int32(1)
	if v982 <= v805 {
		v942 = v982
		goto L156
	} else {
		goto L159
	}
L159:
	;
	goto L157
L160:
	;
	v1030 = v9
	goto L163
L161:
	;
	v1360 = l0
	v1361 = l1
	v1362 = l2
	v1365 = l5
	v1366 = l6
	v1374 = v9
	v1375 = v763
	v1377 = v35
	v1382 = v37
	v1384 = v9
	goto L162
L162:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_consider_groupingsets_paths[2])) = v923
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(v934+v805<<(uint(int32(2))%32))))
	v1398 = F_bms_copy(m, v1397)
	mBase = m.M
	v1399 = m.ExcPending
	if v1399 != 0 {
		goto L43
	} else {
		goto L198
	}
L163:
	;
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v790+v1030<<(uint(int32(2))%32))))
	if v1053 <= v805 {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	v1360 = l0
	v1361 = l1
	v1362 = l2
	v1365 = l5
	v1366 = l6
	v1374 = v9
	v1375 = v763
	v1377 = v35
	v1382 = v37
	v1384 = v9
	goto L162
L165:
	;
	v1059 = v805
	goto L168
L166:
	;
	goto L167
L167:
	;
	v1358 = v1030 + int32(1)
	if v1358 != v888 {
		v1030 = v1358
		goto L163
	} else {
		goto L197
	}
L168:
	;
	v1087 = int32(3)
	v1089 = v930 + v1059<<(uint(v1087)%32)
	v1090 = *(*float64)(unsafe.Add(mBase, uint32(v1089)))
	v1091 = v1059 - v1053
	v1094 = v930 + v1091<<(uint(v1087)%32)
	v1095 = *(*float64)(unsafe.Add(mBase, uint32(v1094)))
	if base.F64_le(v1090, base.F64_add(v1095, float64(1))) != 0 {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	goto L167
L170:
	;
	v1101 = v934 + v1059<<(uint(int32(2))%32)
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(v1101)))
	if v1053 != 0 {
		goto L173
	} else {
		goto L174
	}
L171:
	;
	goto L172
L172:
	;
	v1323 = v1059 - int32(1)
	if v1053 <= v1323 {
		v1059 = v1323
		goto L168
	} else {
		goto L196
	}
L173:
	;
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(v934+v1091<<(uint(int32(2))%32))))
	if v1102 == int32(0) {
		goto L178
	} else {
		goto L179
	}
L174:
	;
	v1261 = v1102
	goto L175
L175:
	;
	v1283 = F_bms_add_member(m, v1261, v1030)
	mBase = m.M
	v1284 = m.ExcPending
	if v1284 != 0 {
		goto L43
	} else {
		goto L195
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1101))) = v1249
	v1261 = v1249
	goto L175
L177:
	;
	v1249 = v1195
	goto L176
L178:
	;
	v1109 = int32(0)
	if v1106 == v1109 {
		v1249 = v1109
		goto L176
	} else {
		goto L181
	}
L179:
	;
	goto L180
L180:
	;
	if v1106 == int32(0) {
		goto L184
	} else {
		goto L185
	}
L181:
	;
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v1106)+4))
	v1116 = v1112<<(uint(int32(2))%32) + int32(8)
	v1117 = F_palloc(m, v1116)
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L43
	} else {
		goto L182
	}
L182:
	;
	if v1116 == int32(0) {
		v1195 = v1117
		goto L177
	} else {
		goto L183
	}
L183:
	;
	base.MemoryCopy(m, v1117, v1106, v1116)
	v1249 = v1117
	goto L176
L184:
	;
	F_pfree(m, v1102)
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		goto L43
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(v1106)+4))
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(v1102)+4))
	if v1128 < v1127 {
		goto L188
	} else {
		goto L189
	}
L187:
	;
	v1249 = int32(0)
	goto L176
L188:
	;
	v1134 = F_repalloc(m, v1102, v1127<<(uint(int32(2))%32)+int32(8))
	mBase = m.M
	v1135 = m.ExcPending
	if v1135 != 0 {
		goto L43
	} else {
		goto L191
	}
L189:
	;
	v1136 = v1102
	goto L190
L190:
	;
	v1137 = int32(8)
	v1150 = int32(0)
	goto L192
L191:
	;
	v1136 = v1134
	goto L190
L192:
	;
	v1175 = v1150 << (uint(int32(2)) % 32)
	v1178 = *(*int32)(unsafe.Add(mBase, uint32(v1175+(v1106+v1137))))
	*(*int32)(unsafe.Add(mBase, uint32(v1136+v1137+v1175))) = v1178
	v1181 = v1150 + int32(1)
	v1182 = *(*int32)(unsafe.Add(mBase, uint32(v1106)+4))
	if v1181 < v1182 {
		v1150 = v1181
		goto L192
	} else {
		goto L194
	}
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1136)+4)) = v1182
	v1195 = v1136
	goto L177
L194:
	;
	goto L193
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1101))) = v1283
	v1286 = *(*float64)(unsafe.Add(mBase, uint32(v1094)))
	*(*float64)(unsafe.Add(mBase, uint32(v1089))) = base.F64_add(v1286, float64(1))
	goto L172
L196:
	;
	goto L169
L197:
	;
	goto L164
L198:
	;
	v1400 = F_bms_del_member(m, v1398, v888)
	mBase = m.M
	v1401 = m.ExcPending
	if v1401 != 0 {
		goto L43
	} else {
		goto L199
	}
L199:
	;
	F_MemoryContextDelete(m, v920)
	mBase = m.M
	v1403 = m.ExcPending
	if v1403 != 0 {
		goto L43
	} else {
		goto L200
	}
L200:
	;
	if v1400 == int32(0) {
		v1484 = v1360
		v1485 = v1361
		v1486 = v1362
		v1489 = v1365
		v1490 = v1366
		v1498 = v1374
		v1499 = v1375
		v1501 = v1377
		v1506 = v1382
		v1508 = v1384
		goto L120
	} else {
		goto L201
	}
L201:
	;
	v1406 = *(*int32)(unsafe.Add(mBase, uint32(v1365)))
	v1407 = *(*int32)(unsafe.Add(mBase, uint32(v1406)+12))
	v1408 = *(*int32)(unsafe.Add(mBase, uint32(v1407)))
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+12)) = v1408
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+24)) = v1408
	v1411 = int32(1)
	v1415 = F_list_make1_impl(m, v1411, v1377+int32(12))
	mBase = m.M
	v1416 = m.ExcPending
	if v1416 != 0 {
		goto L43
	} else {
		goto L202
	}
L202:
	;
	v1417 = *(*int32)(unsafe.Add(mBase, uint32(v1365)))
	if v1417 == int32(0) {
		v1484 = v1360
		v1485 = v1361
		v1486 = v1362
		v1489 = v1365
		v1490 = v1366
		v1498 = v1415
		v1499 = v1375
		v1501 = v1377
		v1506 = v1382
		v1508 = v1384
		goto L120
	} else {
		goto L203
	}
L203:
	;
	v1420 = *(*int32)(unsafe.Add(mBase, uint32(v1417)+4))
	if v1420 < int32(2) {
		v1484 = v1360
		v1485 = v1361
		v1486 = v1362
		v1489 = v1365
		v1490 = v1366
		v1498 = v1415
		v1499 = v1375
		v1501 = v1377
		v1506 = v1382
		v1508 = v1384
		goto L120
	} else {
		goto L204
	}
L204:
	;
	v1428 = int32(0)
	v1432 = v1411
	v1438 = v1415
	v1439 = v1375
	goto L205
L205:
	;
	v1456 = *(*int32)(unsafe.Add(mBase, uint32(v1417)+12))
	v1460 = *(*int32)(unsafe.Add(mBase, uint32(v1456+v1432<<(uint(int32(2))%32))))
	v1461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1460)+24)))
	if v1461 == int32(1) {
		goto L208
	} else {
		goto L209
	}
L206:
	;
	v1484 = v1360
	v1485 = v1361
	v1486 = v1362
	v1489 = v1365
	v1490 = v1366
	v1498 = v1478
	v1499 = v1479
	v1501 = v1377
	v1506 = v1382
	v1508 = v1384
	goto L120
L207:
	;
	v1481 = v1432 + int32(1)
	v1482 = *(*int32)(unsafe.Add(mBase, uint32(v1417)+4))
	if v1481 < v1482 {
		v1428 = v1477
		v1432 = v1481
		v1438 = v1478
		v1439 = v1479
		goto L205
	} else {
		goto L218
	}
L208:
	;
	v1464 = F_bms_is_member(m, v1428, v1400)
	mBase = m.M
	v1465 = m.ExcPending
	if v1465 != 0 {
		goto L43
	} else {
		goto L211
	}
L209:
	;
	goto L210
L210:
	;
	v1475 = F_lappend(m, v1438, v1460)
	mBase = m.M
	v1476 = m.ExcPending
	if v1476 != 0 {
		goto L43
	} else {
		goto L217
	}
L211:
	;
	if v1464 != 0 {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v1466 = *(*int32)(unsafe.Add(mBase, uint32(v1460)+12))
	v1467 = F_list_concat(m, v1439, v1466)
	mBase = m.M
	v1468 = m.ExcPending
	if v1468 != 0 {
		goto L43
	} else {
		goto L215
	}
L213:
	;
	goto L214
L214:
	;
	v1471 = F_lappend(m, v1438, v1460)
	mBase = m.M
	v1472 = m.ExcPending
	if v1472 != 0 {
		goto L43
	} else {
		goto L216
	}
L215:
	;
	v1477 = v1428 + int32(1)
	v1478 = v1438
	v1479 = v1467
	goto L207
L216:
	;
	v1477 = v1428 + int32(1)
	v1478 = v1471
	v1479 = v1439
	goto L207
L217:
	;
	v1477 = v1428
	v1478 = v1475
	v1479 = v1439
	goto L207
L218:
	;
	goto L206
L219:
	;
	if v1865 == int32(0) {
		v1891 = v1484
		v1892 = v1485
		v1893 = v1486
		v1896 = v1489
		v1897 = v1490
		v1908 = v1501
		v1913 = v1506
		goto L116
	} else {
		goto L256
	}
L220:
	;
	v1527 = *(*int32)(unsafe.Add(mBase, uint32(v1499)+4))
	if v1527 <= int32(0) {
		v1865 = v1526
		goto L219
	} else {
		goto L226
	}
L221:
	;
	v1521 = *(*int32)(unsafe.Add(mBase, uint32(v1489)))
	v1522 = F_list_copy(m, v1521)
	mBase = m.M
	v1523 = m.ExcPending
	if v1523 != 0 {
		goto L43
	} else {
		goto L224
	}
L222:
	;
	goto L223
L223:
	;
	if v1499 == int32(0) {
		v1865 = v1498
		goto L219
	} else {
		goto L225
	}
L224:
	;
	v1526 = v1522
	goto L220
L225:
	;
	v1526 = v1498
	goto L220
L226:
	;
	v1544 = v1526
	v1554 = v1508
	goto L227
L227:
	;
	v1562 = *(*int32)(unsafe.Add(mBase, uint32(v1499)+12))
	v1566 = *(*int32)(unsafe.Add(mBase, uint32(v1562+v1554<<(uint(int32(2))%32))))
	v1568 = F_palloc0(m, int32(32))
	mBase = m.M
	v1569 = m.ExcPending
	if v1569 != 0 {
		goto L43
	} else {
		goto L229
	}
L228:
	;
	v1865 = v1845
	goto L219
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1568))) = int32(311)
	v1572 = *(*int32)(unsafe.Add(mBase, uint32(v1566)+4))
	v1573 = F_preprocess_groupclause(m, v1484, v1572)
	mBase = m.M
	v1574 = m.ExcPending
	if v1574 != 0 {
		goto L43
	} else {
		goto L230
	}
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1568)+4)) = v1573
	*(*int32)(unsafe.Add(mBase, uint32(v1501)+8)) = v1566
	*(*int32)(unsafe.Add(mBase, uint32(v1501)+20)) = v1566
	v1581 = F_list_make1_impl(m, int32(1), v1501+int32(8))
	mBase = m.M
	v1582 = m.ExcPending
	if v1582 != 0 {
		goto L43
	} else {
		goto L231
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1568)+12)) = v1581
	v1584 = *(*int32)(unsafe.Add(mBase, uint32(v1489)+32))
	v1585 = *(*int32)(unsafe.Add(mBase, uint32(v1568)+4))
	if v1585 == int32(0) {
		goto L232
	} else {
		goto L233
	}
L232:
	;
	if v1581 == int32(0) {
		goto L239
	} else {
		goto L240
	}
L233:
	;
	v1588 = int32(0)
	v1589 = *(*int32)(unsafe.Add(mBase, uint32(v1585)+4))
	if v1589 <= v1588 {
		goto L232
	} else {
		goto L234
	}
L234:
	;
	v1600 = v1588
	goto L235
L235:
	;
	v1624 = *(*int32)(unsafe.Add(mBase, uint32(v1585)+12))
	v1625 = int32(2)
	v1628 = *(*int32)(unsafe.Add(mBase, uint32(v1624+v1600<<(uint(v1625)%32))))
	v1629 = *(*int32)(unsafe.Add(mBase, uint32(v1628)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1584+v1629<<(uint(v1625)%32)))) = v1600
	v1635 = v1600 + int32(1)
	v1636 = *(*int32)(unsafe.Add(mBase, uint32(v1585)+4))
	if v1635 < v1636 {
		v1600 = v1635
		goto L235
	} else {
		goto L237
	}
L236:
	;
	goto L232
L237:
	;
	goto L236
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1568)+8)) = v1812
	v1841 = *(*float64)(unsafe.Add(mBase, uint32(v1566)+8))
	v1842 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v1568)+24)) = uint16(v1842)
	*(*float64)(unsafe.Add(mBase, uint32(v1568)+16)) = v1841
	v1845 = F_lcons(m, v1568, v1544)
	mBase = m.M
	v1846 = m.ExcPending
	if v1846 != 0 {
		goto L43
	} else {
		goto L254
	}
L239:
	;
	v1812 = int32(0)
	goto L238
L240:
	;
	goto L241
L241:
	;
	v1673 = int32(0)
	v1675 = *(*int32)(unsafe.Add(mBase, uint32(v1581)+4))
	if v1675 <= v1673 {
		v1812 = v1673
		goto L238
	} else {
		goto L242
	}
L242:
	;
	v1682 = v1673
	v1688 = v1673
	goto L243
L243:
	;
	v1710 = int32(0)
	v1711 = *(*int32)(unsafe.Add(mBase, uint32(v1581)+12))
	v1715 = *(*int32)(unsafe.Add(mBase, uint32(v1711+v1688<<(uint(int32(2))%32))))
	v1716 = *(*int32)(unsafe.Add(mBase, uint32(v1715)+4))
	if v1716 == v1710 {
		v1773 = v1710
		goto L245
	} else {
		goto L246
	}
L244:
	;
	v1812 = v1802
	goto L238
L245:
	;
	v1802 = F_lappend(m, v1682, v1773)
	mBase = m.M
	v1803 = m.ExcPending
	if v1803 != 0 {
		goto L43
	} else {
		goto L252
	}
L246:
	;
	v1719 = int32(0)
	v1720 = *(*int32)(unsafe.Add(mBase, uint32(v1716)+4))
	if v1720 <= v1719 {
		v1773 = v1710
		goto L245
	} else {
		goto L247
	}
L247:
	;
	v1726 = v1710
	v1731 = v1719
	goto L248
L248:
	;
	v1755 = *(*int32)(unsafe.Add(mBase, uint32(v1716)+12))
	v1756 = int32(2)
	v1759 = *(*int32)(unsafe.Add(mBase, uint32(v1755+v1731<<(uint(v1756)%32))))
	v1763 = *(*int32)(unsafe.Add(mBase, uint32(v1584+v1759<<(uint(v1756)%32))))
	v1764 = F_lappend_int(m, v1726, v1763)
	mBase = m.M
	v1765 = m.ExcPending
	if v1765 != 0 {
		goto L43
	} else {
		goto L250
	}
L249:
	;
	v1773 = v1764
	goto L245
L250:
	;
	v1767 = v1731 + int32(1)
	v1768 = *(*int32)(unsafe.Add(mBase, uint32(v1716)+4))
	if v1767 < v1768 {
		v1726 = v1764
		v1731 = v1767
		goto L248
	} else {
		goto L251
	}
L251:
	;
	goto L249
L252:
	;
	v1805 = v1688 + int32(1)
	v1806 = *(*int32)(unsafe.Add(mBase, uint32(v1581)+4))
	if v1805 < v1806 {
		v1682 = v1802
		v1688 = v1805
		goto L243
	} else {
		goto L253
	}
L253:
	;
	goto L244
L254:
	;
	v1848 = v1554 + int32(1)
	v1849 = *(*int32)(unsafe.Add(mBase, uint32(v1499)+4))
	if v1848 < v1849 {
		v1544 = v1845
		v1554 = v1848
		goto L227
	} else {
		goto L255
	}
L255:
	;
	goto L228
L256:
	;
	v1885 = *(*int32)(unsafe.Add(mBase, uint32(v1506)+112))
	v1887 = F_create_groupingsets_path(m, v1484, v1485, v1486, v1885, int32(3), v1865, v1490)
	mBase = m.M
	v1888 = m.ExcPending
	if v1888 != 0 {
		goto L43
	} else {
		goto L257
	}
L257:
	;
	F_add_path(m, v1485, v1887)
	mBase = m.M
	v1890 = m.ExcPending
	if v1890 != 0 {
		goto L43
	} else {
		goto L258
	}
L258:
	;
	v1891 = v1484
	v1892 = v1485
	v1893 = v1486
	v1896 = v1489
	v1897 = v1490
	v1908 = v1501
	v1913 = v1506
	goto L116
L259:
	;
	v1924 = *(*int32)(unsafe.Add(mBase, uint32(v1913)+112))
	v1926 = *(*int32)(unsafe.Add(mBase, uint32(v1896)))
	v1927 = F_create_groupingsets_path(m, v1891, v1892, v1893, v1924, int32(1), v1926, v1897)
	mBase = m.M
	v1928 = m.ExcPending
	if v1928 != 0 {
		goto L43
	} else {
		goto L260
	}
L260:
	;
	v1946 = v1908
	v1961 = v1927
	goto L6
L261:
	;
	v1981 = v1946
	goto L5
}
func F_conv_18030_to_utf8(m *base.Module, l0 int32) int32 {
	var v15 int32
	_ = v15
	var v26 int32
	_ = v26
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v139 int32
	_ = v139
	var v150 int32
	_ = v150
	var v172 int32
	_ = v172
	var v180 int32
	_ = v180
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v272 int32
	_ = v272
	var v283 int32
	_ = v283
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v341 int32
	_ = v341
	var v349 int32
	_ = v349
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v409 int32
	_ = v409
	var v427 int32
	_ = v427
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v487 int32
	_ = v487
	var v505 int32
	_ = v505
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v567 int32
	_ = v567
	var v578 int32
	_ = v578
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v644 int32
	_ = v644
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v704 int32
	_ = v704
	var v722 int32
	_ = v722
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	if base.Ui32(l0+int32(2127506640)) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_0)) {
		v15 = int32(255)
		v26 = int32(base.Ui32(l0)>>(uint(int32(16))%32))&int32(55)*int32(1260) + l0&v15 + int32(base.Ui32(l0)>>(uint(int32(8))%32))&v15*int32(10) - int32(_a_F_conv_18030_to_utf8_1)
		if base.Ui32(v26) < base.Ui32(int32(128)) {
			v774 = v26
			return v774
		} else {
			if base.Ui32(v26) <= base.Ui32(int32(2047)) {
				return v26&int32(63) | v26<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_2) | int32(_a_F_conv_18030_to_utf8_3)
			} else {
				v42 = v26 & int32(63)
				v44 = v26 << (uint(int32(4)) % 32)
				v48 = v26 << (uint(int32(2)) % 32) & int32(_a_F_conv_18030_to_utf8_4)
				if base.Ui32(v26) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_5)) {
					v776 = v42
					v778 = v44
					v779 = v48
					return v779 | v778&int32(_a_F_conv_18030_to_utf8_6) | v776 | int32(14712960)
				} else {
					v787 = v42
					v789 = v44
					v790 = v48
					return v787 | v789&int32(_a_F_conv_18030_to_utf8_7) | v790 | int32(-142573440)
				}
			}
		}
	} else {
		if base.Ui32(l0+int32(2127058887)) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_8)) {
			v61 = int32(255)
			v70 = int32(base.Ui32(l0)>>(uint(int32(16))%32))&int32(63)*int32(1260) + l0&v61 + int32(base.Ui32(l0)>>(uint(int32(8))%32))&v61*int32(10)
			v72 = v70 - int32(_a_F_conv_18030_to_utf8_9)
			if base.Ui32(v72) < base.Ui32(int32(128)) {
				v774 = v72
				return v774
			} else {
				v76 = v70 - int32(_a_F_conv_18030_to_utf8_10)
				if base.Ui32(v72) <= base.Ui32(int32(2047)) {
					return v76&int32(63) | v72<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_2) | int32(_a_F_conv_18030_to_utf8_3)
				} else {
					v90 = v76 & int32(63)
					v92 = v72 << (uint(int32(4)) % 32)
					v96 = v72 << (uint(int32(2)) % 32) & int32(_a_F_conv_18030_to_utf8_4)
					if base.Ui32(v72) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_5)) {
						v776 = v90
						v778 = v92
						v779 = v96
						return v779 | v778&int32(_a_F_conv_18030_to_utf8_6) | v776 | int32(14712960)
					} else {
						v787 = v90
						v789 = v92
						v790 = v96
						return v787 | v789&int32(_a_F_conv_18030_to_utf8_7) | v790 | int32(-142573440)
					}
				}
			}
		} else {
			if base.Ui32(l0+int32(2110740941)) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_11)) {
				v105 = int32(255)
				v113 = int32(base.Ui32(l0)>>(uint(int32(8))%32))&v105*int32(10) + l0&v105 + int32(_a_F_conv_18030_to_utf8_12)
				return v113&int32(63) | v113<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_4) | v113<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_13) | int32(14712960)
			} else {
				if base.Ui32(l0+int32(2110663624)) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_14)) {
					v139 = int32(255)
					v150 = int32(base.Ui32(l0)>>(uint(int32(16))%32))&int32(51)*int32(1260) + l0&v139 + int32(base.Ui32(l0)>>(uint(int32(8))%32))&v139*int32(10) - int32(_a_F_conv_18030_to_utf8_15)
					return v150&int32(63) | v150<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_4) | v150<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_6) | int32(14712960)
				} else {
					if base.Ui32(l0+int32(2110600905)) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_16)) {
						v172 = int32(255)
						v180 = int32(base.Ui32(l0)>>(uint(int32(8))%32))&v172*int32(10) + l0&v172 + int32(_a_F_conv_18030_to_utf8_17)
						return v180&int32(63) | v180<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_4) | v180<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_13) | int32(14712960)
					} else {
						if base.Ui32(l0+int32(2110545095)) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_18)) {
							v202 = int32(255)
							v210 = int32(base.Ui32(l0)>>(uint(int32(8))%32))&v202*int32(10) + l0&v202 + int32(_a_F_conv_18030_to_utf8_19)
							if base.Ui32(int32(128)) <= base.Ui32(v210) {
								if base.Ui32(v210) <= base.Ui32(int32(2047)) {
									v260 = v210&int32(63) | v210<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_2) | int32(_a_F_conv_18030_to_utf8_3)
								} else {
									if base.Ui32(v210) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_5)) {
										v260 = v210<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_6) | (v210&int32(63) | v210<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_4)) | int32(14712960)
									} else {
										v259 = v210<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_4) | (v210<<(uint(int32(6))%32)&int32(117440512) | (v210&int32(63) | v210<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_7))) | int32(-260013952)
										v260 = v259
									}
								}
							} else {
								v259 = v210
								v260 = v259
							}
							return v260
						} else {
							if base.Ui32(l0+int32(2110527432)) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_20)) {
								v272 = int32(255)
								v283 = int32(base.Ui32(l0)>>(uint(int32(16))%32))&int32(55)*int32(1260) + l0&v272 + int32(base.Ui32(l0)>>(uint(int32(8))%32))&v272*int32(10) - int32(_a_F_conv_18030_to_utf8_21)
								if base.Ui32(int32(128)) <= base.Ui32(v283) {
									if base.Ui32(v283) <= base.Ui32(int32(2047)) {
										v333 = v283&int32(63) | v283<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_2) | int32(_a_F_conv_18030_to_utf8_3)
									} else {
										if base.Ui32(v283) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_5)) {
											v333 = v283<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_6) | (v283&int32(63) | v283<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_4)) | int32(14712960)
										} else {
											v332 = v283<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_4) | (v283<<(uint(int32(6))%32)&int32(117440512) | (v283&int32(63) | v283<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_7))) | int32(-260013952)
											v333 = v332
										}
									}
								} else {
									v332 = v283
									v333 = v332
								}
								return v333
							} else {
								if base.Ui32(l0+int32(2110480079)) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_22)) {
									v341 = int32(255)
									v349 = int32(base.Ui32(l0)>>(uint(int32(8))%32))&v341*int32(10) + l0&v341 + int32(_a_F_conv_18030_to_utf8_23)
									if base.Ui32(int32(128)) <= base.Ui32(v349) {
										if base.Ui32(v349) <= base.Ui32(int32(2047)) {
											v399 = v349&int32(63) | v349<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_2) | int32(_a_F_conv_18030_to_utf8_3)
										} else {
											if base.Ui32(v349) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_5)) {
												v399 = v349<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_6) | (v349&int32(63) | v349<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_4)) | int32(14712960)
											} else {
												v398 = v349<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_4) | (v349<<(uint(int32(6))%32)&int32(117440512) | (v349&int32(63) | v349<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_7))) | int32(-260013952)
												v399 = v398
											}
										}
									} else {
										v398 = v349
										v399 = v398
									}
									return v399
								} else {
									if base.Ui32(l0+int32(2110419149)) <= base.Ui32(int32(16857093)) {
										v409 = int32(255)
										v427 = int32(base.Ui32(l0)>>(uint(int32(24))%32))*int32(_a_F_conv_18030_to_utf8_24) + l0&v409 + int32(base.Ui32(l0)>>(uint(int32(16))%32))&v409*int32(1260) + int32(base.Ui32(l0)>>(uint(int32(8))%32))&v409*int32(10) - int32(_a_F_conv_18030_to_utf8_25)
										if base.Ui32(int32(128)) <= base.Ui32(v427) {
											if base.Ui32(v427) <= base.Ui32(int32(2047)) {
												v477 = v427&int32(63) | v427<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_2) | int32(_a_F_conv_18030_to_utf8_3)
											} else {
												if base.Ui32(v427) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_5)) {
													v477 = v427<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_6) | (v427&int32(63) | v427<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_4)) | int32(14712960)
												} else {
													v476 = v427<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_4) | (v427<<(uint(int32(6))%32)&int32(117440512) | (v427&int32(63) | v427<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_7))) | int32(-260013952)
													v477 = v476
												}
											}
										} else {
											v476 = v427
											v477 = v476
										}
										return v477
									} else {
										if base.Ui32(l0+int32(2093559760)) <= base.Ui32(int32(16364804)) {
											v487 = int32(255)
											v505 = int32(base.Ui32(l0)>>(uint(int32(24))%32))*int32(_a_F_conv_18030_to_utf8_24) + l0&v487 + int32(base.Ui32(l0)>>(uint(int32(16))%32))&v487*int32(1260) + int32(base.Ui32(l0)>>(uint(int32(8))%32))&v487*int32(10) - int32(_a_F_conv_18030_to_utf8_26)
											if base.Ui32(int32(128)) <= base.Ui32(v505) {
												if base.Ui32(v505) <= base.Ui32(int32(2047)) {
													v555 = v505&int32(63) | v505<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_2) | int32(_a_F_conv_18030_to_utf8_3)
												} else {
													if base.Ui32(v505) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_5)) {
														v555 = v505<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_6) | (v505&int32(63) | v505<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_4)) | int32(14712960)
													} else {
														v554 = v505<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_4) | (v505<<(uint(int32(6))%32)&int32(117440512) | (v505&int32(63) | v505<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_7))) | int32(-260013952)
														v555 = v554
													}
												}
											} else {
												v554 = v505
												v555 = v554
											}
											return v555
										} else {
											if base.Ui32(l0+int32(2077189064)) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_27)) {
												v567 = int32(255)
												v578 = int32(base.Ui32(l0)>>(uint(int32(16))%32))&int32(49)*int32(1260) + l0&v567 + int32(base.Ui32(l0)>>(uint(int32(8))%32))&v567*int32(10) + int32(1946)
												if base.Ui32(int32(128)) <= base.Ui32(v578) {
													if base.Ui32(v578) <= base.Ui32(int32(2047)) {
														v628 = v578&int32(63) | v578<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_2) | int32(_a_F_conv_18030_to_utf8_3)
													} else {
														if base.Ui32(v578) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_5)) {
															v628 = v578<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_6) | (v578&int32(63) | v578<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_4)) | int32(14712960)
														} else {
															v627 = v578<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_4) | (v578<<(uint(int32(6))%32)&int32(117440512) | (v578&int32(63) | v578<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_7))) | int32(-260013952)
															v628 = v627
														}
													}
												} else {
													v627 = v578
													v628 = v627
												}
												return v628
											} else {
												if base.Ui32(l0+int32(2077121996)) <= base.Ui32(int32(517)) {
													v644 = int32(base.Ui32(l0)>>(uint(int32(8))%32))&int32(167)*int32(10) + l0&int32(255) + int32(_a_F_conv_18030_to_utf8_28)
													if base.Ui32(int32(128)) <= base.Ui32(v644) {
														if base.Ui32(v644) <= base.Ui32(int32(2047)) {
															v694 = v644&int32(63) | v644<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_2) | int32(_a_F_conv_18030_to_utf8_3)
														} else {
															if base.Ui32(v644) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_5)) {
																v694 = v644<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_6) | (v644&int32(63) | v644<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_4)) | int32(14712960)
															} else {
																v693 = v644<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_4) | (v644<<(uint(int32(6))%32)&int32(117440512) | (v644&int32(63) | v644<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_7))) | int32(-260013952)
																v694 = v693
															}
														}
													} else {
														v693 = v644
														v694 = v693
													}
													return v694
												} else {
													if base.Ui32(int32(1392646405)) < base.Ui32(l0+int32(1875869392)) {
														v774 = int32(0)
													} else {
														v704 = int32(255)
														v722 = int32(base.Ui32(l0)>>(uint(int32(24))%32))*int32(_a_F_conv_18030_to_utf8_24) + l0&v704 + int32(base.Ui32(l0)>>(uint(int32(16))%32))&v704*int32(1260) + int32(base.Ui32(l0)>>(uint(int32(8))%32))&v704*int32(10) - int32(_a_F_conv_18030_to_utf8_29)
														if base.Ui32(int32(128)) <= base.Ui32(v722) {
															if base.Ui32(v722) <= base.Ui32(int32(2047)) {
																v772 = v722&int32(63) | v722<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_2) | int32(_a_F_conv_18030_to_utf8_3)
															} else {
																if base.Ui32(v722) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_5)) {
																	v772 = v722<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_6) | (v722&int32(63) | v722<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_4)) | int32(14712960)
																} else {
																	v771 = v722<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_4) | (v722<<(uint(int32(6))%32)&int32(117440512) | (v722&int32(63) | v722<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_7))) | int32(-260013952)
																	v772 = v771
																}
															}
														} else {
															v771 = v722
															v772 = v771
														}
														v774 = v772
													}
													return v774
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
func F_convert_tuples_by_name_attrmap(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
		v16 = F_palloc_mul(m, int32(8), v6)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v16
			v20 = F_palloc_mul(m, int32(1), v6)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v20
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v26 = v24 + int32(1)
				v27 = F_palloc_mul(m, int32(8), v26)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v27
					v31 = F_palloc_mul(m, int32(1), v26)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v31
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
						*(*int64)(unsafe.Add(mBase, uint32(v34))) = int64(0)
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
	var v75 float64
	_ = v75
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
	var v110 int64
	_ = v110
	var v113 float64
	_ = v113
	var v114 float64
	_ = v114
	var v118 int32
	_ = v118
	var v119 int64
	_ = v119
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
	v107 = *(*float64)(unsafe.Add(mBase, uint32(l2)+128))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v109 = *(*float64)(unsafe.Add(mBase, uint32(v108)+24))
	v110 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
	v113 = *(*float64)(unsafe.Add(mBase, uint32(v108)+16))
	v114 = base.F64_add(base.F64_add(v106, float64(0)), v113)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = v114
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v118 != 0 {
		goto L13
	} else {
		goto L14
	}
L2:
	;
	v19 = *(*float64)(unsafe.Add(mBase, uint32(l3)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v22 = int32(0)
	v24 = *(*float64)(unsafe.Add(mBase, _c_F_cost_ctescan[0]))
	v25 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = v25
	if v21 == v22 {
		v73 = v7
		v75 = v7
		v76 = v19
		v81 = v24
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v86 = *(*float64)(unsafe.Add(mBase, uint32(l2)+16))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v86
	v88 = *(*float64)(unsafe.Add(mBase, uint32(l2)+216))
	v90 = *(*float64)(unsafe.Add(mBase, _c_F_cost_ctescan[0]))
	v91 = *(*float64)(unsafe.Add(mBase, uint32(l2)+208))
	v99 = v90
	v100 = v88
	v101 = v86
	v102 = v90
	v106 = v91
	goto L1
L5:
	;
	v82 = *(*float64)(unsafe.Add(mBase, uint32(l2)+216))
	v84 = *(*float64)(unsafe.Add(mBase, uint32(l2)+208))
	v99 = v24
	v100 = base.F64_add(v75, v82)
	v101 = v76
	v102 = v81
	v106 = base.F64_add(v73, v84)
	goto L1
L6:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v32 <= int32(0) {
		v73 = v7
		v75 = v7
		v76 = v19
		v81 = v24
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v38 = v22
	goto L8
L8:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49+v38<<(uint(int32(2))%32))))
	v56 = F_cost_qual_eval_walker(m, v53, v17+int32(8))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v62 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
	v63 = *(*float64)(unsafe.Add(mBase, uint32(v17)+24))
	v64 = *(*float64)(unsafe.Add(mBase, uint32(v17)+16))
	v66 = *(*float64)(unsafe.Add(mBase, _c_F_cost_ctescan[0]))
	v73 = v64
	v75 = v63
	v76 = v62
	v81 = v66
	goto L5
L10:
	;
	return
L11:
	;
	v59 = v38 + int32(1)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v59 < v60 {
		v38 = v59
		goto L8
	} else {
		goto L12
	}
L12:
	;
	goto L9
L13:
	;
	v119 = int64(-1)
	goto L15
L14:
	;
	v119 = int64(-262145)
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = base.B2i32(v110|v119 != int64(-1))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = base.F64_add(v114, base.F64_add(base.F64_mul(v109, v101), base.F64_add(base.F64_mul(v107, base.F64_add(v99, base.F64_add(v100, v102))), float64(0))))
	m.G0 = v17 + int32(32)
	return
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
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
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
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v104 float64
	_ = v104
	var v105 int32
	_ = v105
	var v121 float64
	_ = v121
	var v125 int32
	_ = v125
	var v126 float64
	_ = v126
	var v132 float64
	_ = v132
	var v135 float64
	_ = v135
	var v137 float64
	_ = v137
	var v140 float64
	_ = v140
	var v143 int64
	_ = v143
	var v144 float64
	_ = v144
	var v151 float64
	_ = v151
	var v153 float64
	_ = v153
	var v157 int32
	_ = v157
	var v158 float64
	_ = v158
	var v160 float64
	_ = v160
	var v161 int32
	_ = v161
	var v164 float64
	_ = v164
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
	var v178 float64
	_ = v178
	var v181 float64
	_ = v181
	var v185 float64
	_ = v185
	var v191 float64
	_ = v191
	var v192 float64
	_ = v192
	var v196 float64
	_ = v196
	var v200 float64
	_ = v200
	var v208 float64
	_ = v208
	var v211 float64
	_ = v211
	var v215 float64
	_ = v215
	var v216 float64
	_ = v216
	var v217 float64
	_ = v217
	var v221 float64
	_ = v221
	var v223 float64
	_ = v223
	var v231 float64
	_ = v231
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
	v125 = v21 + int32(8)
	v126 = base.F64_div(v26, v121)
	v132 = float64(2)
	if base.F64_lt(v126, v132) != 0 {
		goto L25
	} else {
		goto L26
	}
L5:
	;
	v102 = int32(0)
	v104 = F_estimate_num_groups(m, l1, v96, v26, v102, v102)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
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
	v46 = int32(0)
	v55 = v12
	goto L14
L14:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v62+v46<<(uint(int32(2))%32))))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	v72 = F_pull_varnos(m, l1, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v96 = v77
	goto L5
L16:
	;
	return
L17:
	;
	v74 = F_bms_is_member(m, int32(0), v72)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	if v74 != 0 {
		v121 = v35
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	v77 = F_lappend(m, v55, v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L16
	} else {
		goto L20
	}
L20:
	;
	if v46 == v39-int32(1) {
		v96 = v77
		goto L5
	} else {
		goto L21
	}
L21:
	;
	v81 = v46 + int32(1)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v81 < v82 {
		v46 = v81
		v55 = v77
		goto L14
	} else {
		goto L22
	}
L22:
	;
	goto L15
L23:
	;
	v121 = v104
	goto L4
L24:
	;
	v215 = *(*float64)(unsafe.Add(mBase, _c_F_cost_incremental_sort[0]))
	v216 = *(*float64)(unsafe.Add(mBase, uint32(v21)))
	v217 = *(*float64)(unsafe.Add(mBase, uint32(v21)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l4
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v26
	v221 = base.F64_div(base.F64_sub(l6, l5), v121)
	v223 = base.F64_add(v221, base.F64_add(l5, v217))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = v223
	v231 = base.F64_add(v121, float64(-1))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = base.F64_add(v223, base.F64_add(base.F64_mul(base.F64_add(v215, v215), v121), base.F64_add(base.F64_mul(base.F64_add(v215, float64(0)), v26), base.F64_add(base.F64_mul(v221, v231), base.F64_add(v216, base.F64_mul(base.F64_add(v217, v216), v231))))))
	m.G0 = v21 + int32(16)
	return
L25:
	;
	v135 = v132
	goto L27
L26:
	;
	v135 = v126
	goto L27
L27:
	;
	v137 = *(*float64)(unsafe.Add(mBase, _c_F_cost_incremental_sort[1]))
	v140 = base.F64_mul(v135, base.F64_add(base.F64_add(v137, v137), float64(0)))
	v143 = base.I64_extend_i32_s(l9) << (uint(int64(10)) % 64)
	v144 = base.F64_convert_i64_s(v143)
	v151 = base.F64_convert_i32_u((l8+int32(7))&int32(-8) + int32(24))
	v153 = base.F64_mul(v126, v151)
	v157 = base.F64_lt(l10, v135) & base.F64_gt(l10, float64(0))
	if v157 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v125))) = v208
	v211 = *(*float64)(unsafe.Add(mBase, _c_F_cost_incremental_sort[1]))
	*(*float64)(unsafe.Add(mBase, uint32(v21))) = base.F64_mul(v135, v211)
	goto L24
L29:
	;
	v158 = base.F64_mul(l10, v151)
	goto L31
L30:
	;
	v158 = v153
	goto L31
L31:
	;
	if base.F64_lt(v144, v158) != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v160 = F_log(m, v135)
	mBase = m.M
	v161 = F_tuplesort_merge_order(m, v143)
	mBase = m.M
	v164 = base.F64_mul(base.F64_div(v160, float64(0.693147180559945)), v140)
	*(*float64)(unsafe.Add(mBase, uint32(v125))) = v164
	v168 = base.F64_ceil(base.F64_mul(v153, float64(0.0001220703125)))
	v170 = base.F64_div(v153, v144)
	v171 = base.F64_convert_i32_s(v161)
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
	if v157 != 0 {
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
	v178 = base.F64_ceil(base.F64_div(v173, v174))
	goto L37
L36:
	;
	v178 = float64(1)
	goto L37
L37:
	;
	v181 = *(*float64)(unsafe.Add(mBase, _c_F_cost_incremental_sort[2]))
	v185 = *(*float64)(unsafe.Add(mBase, _c_F_cost_incremental_sort[3]))
	v208 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v168, v168), v178), base.F64_add(base.F64_mul(v181, float64(0.75)), base.F64_mul(v185, float64(0.25)))), v164)
	goto L28
L38:
	;
	v191 = l10
	goto L40
L39:
	;
	v191 = v135
	goto L40
L40:
	;
	v192 = base.F64_add(v191, v191)
	if base.F64_gt(v135, v192)|base.F64_gt(v153, v144) != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v196 = F_log(m, v192)
	mBase = m.M
	v208 = base.F64_mul(base.F64_div(v196, float64(0.693147180559945)), v140)
	goto L28
L42:
	;
	goto L43
L43:
	;
	v200 = F_log(m, v135)
	mBase = m.M
	v208 = base.F64_mul(base.F64_div(v200, float64(0.693147180559945)), v140)
	goto L28
}
