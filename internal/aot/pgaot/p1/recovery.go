package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_ExecuteRecoveryCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int64
	_ = v22
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	var v27 int64
	_ = v27
	var v28 int64
	_ = v28
	var v31 int64
	_ = v31
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v89 int32
	_ = v89
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	v10 = m.G0
	v12 = v10 - int32(1104)
	m.G0 = v12
	F_GetOldestRestartPoint(m, v12+int32(72), v12+int32(68))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v12)+68))
		*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v20
		v22 = *(*int64)(unsafe.Add(mBase, uint32(v12)+72))
		v24 = int64(*(*int32)(unsafe.Add(mBase, _c_F_ExecuteRecoveryCommand[0])))
		v25 = base.I64_div_u_s(v22, v24)
		v27 = base.I64_div_u_s(int64(4294967296), v24)
		v28 = base.I64_div_u_s(v25, v27)
		*(*uint32)(unsafe.Add(mBase, uint32(v12)+52)) = uint32(v28)
		v31 = v25 - v27*v28
		*(*uint32)(unsafe.Add(mBase, uint32(v12)+56)) = uint32(v31)
		v34 = v12 + int32(80)
		v39 = F_pg_snprintf(m, v34, int32(64), int32(_a_F_ExecuteRecoveryCommand_0), v12+int32(48))
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v34
			v45 = F_replace_percent_placeholders(m, l0, l1, int32(_a_F_ExecuteRecoveryCommand_1), v12+int32(32))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return
			} else {
				v49 = F_errstart(m, int32(12), int32(0))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return
				} else {
					if v49 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l1
						F_errmsg_internal(m, int32(_a_F_ExecuteRecoveryCommand_2), v12+int32(16))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_ExecuteRecoveryCommand_3), int32(324), int32(_a_F_ExecuteRecoveryCommand_4))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return
							} else {
								v64 = F_fflush(m, int32(0))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return
								} else {
									v67 = *(*int32)(unsafe.Add(mBase, _c_F_ExecuteRecoveryCommand[1]))
									*(*int32)(unsafe.Add(mBase, uint32(v67))) = l3
									v69 = F_pgl_system(m, v45)
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return
									} else {
										v72 = *(*int32)(unsafe.Add(mBase, _c_F_ExecuteRecoveryCommand[1]))
										*(*int32)(unsafe.Add(mBase, uint32(v72))) = int32(0)
										F_pfree(m, v45)
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return
										} else {
											if v69 == int32(0) {
												m.G0 = v12 + int32(1104)
												return
											} else {
												if l2 != 0 {
													v89 = int32(255)
													if base.B2i32(v69&int32(127) == int32(0))&base.B2i32(base.Ui32(int32(125)) < base.Ui32(int32(base.Ui32(v69)>>(uint(int32(8))%32))&v89))|base.B2i32(base.Ui32(v69&int32(_a_F_ExecuteRecoveryCommand_5)-int32(1)) < base.Ui32(v89)) != 0 {
														v101 = int32(22)
													} else {
														v101 = int32(19)
													}
													v103 = v101
												} else {
													v103 = int32(19)
												}
												v105 = F_errstart(m, v103, int32(0))
												mBase = m.M
												v106 = m.ExcPending
												if v106 != 0 {
													return
												} else {
													if v105 == int32(0) {
														m.G0 = v12 + int32(1104)
														return
													} else {
														v109 = F_wait_result_to_str(m, v69)
														mBase = m.M
														v110 = m.ExcPending
														if v110 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v109
															*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = l0
															*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
															F_errmsg(m, int32(_a_F_ExecuteRecoveryCommand_6), v12)
															mBase = m.M
															v116 = m.ExcPending
															if v116 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_ExecuteRecoveryCommand_3), int32(348), int32(_a_F_ExecuteRecoveryCommand_4))
																mBase = m.M
																v121 = m.ExcPending
																if v121 != 0 {
																	return
																} else {
																	m.G0 = v12 + int32(1104)
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
					} else {
						v64 = F_fflush(m, int32(0))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return
						} else {
							v67 = *(*int32)(unsafe.Add(mBase, _c_F_ExecuteRecoveryCommand[1]))
							*(*int32)(unsafe.Add(mBase, uint32(v67))) = l3
							v69 = F_pgl_system(m, v45)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return
							} else {
								v72 = *(*int32)(unsafe.Add(mBase, _c_F_ExecuteRecoveryCommand[1]))
								*(*int32)(unsafe.Add(mBase, uint32(v72))) = int32(0)
								F_pfree(m, v45)
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return
								} else {
									if v69 == int32(0) {
										m.G0 = v12 + int32(1104)
										return
									} else {
										if l2 != 0 {
											v89 = int32(255)
											if base.B2i32(v69&int32(127) == int32(0))&base.B2i32(base.Ui32(int32(125)) < base.Ui32(int32(base.Ui32(v69)>>(uint(int32(8))%32))&v89))|base.B2i32(base.Ui32(v69&int32(_a_F_ExecuteRecoveryCommand_5)-int32(1)) < base.Ui32(v89)) != 0 {
												v101 = int32(22)
											} else {
												v101 = int32(19)
											}
											v103 = v101
										} else {
											v103 = int32(19)
										}
										v105 = F_errstart(m, v103, int32(0))
										mBase = m.M
										v106 = m.ExcPending
										if v106 != 0 {
											return
										} else {
											if v105 == int32(0) {
												m.G0 = v12 + int32(1104)
												return
											} else {
												v109 = F_wait_result_to_str(m, v69)
												mBase = m.M
												v110 = m.ExcPending
												if v110 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v109
													*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = l0
													*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
													F_errmsg(m, int32(_a_F_ExecuteRecoveryCommand_6), v12)
													mBase = m.M
													v116 = m.ExcPending
													if v116 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_ExecuteRecoveryCommand_3), int32(348), int32(_a_F_ExecuteRecoveryCommand_4))
														mBase = m.M
														v121 = m.ExcPending
														if v121 != 0 {
															return
														} else {
															m.G0 = v12 + int32(1104)
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
func F_RecoveryRequiresIntParameter(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	if l1 < l2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RecoveryRequiresIntParameter[0])))
	if v14 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v10 + int32(48)
	return
L4:
	;
	v19 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
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
	v168 = m.ExcPending
	if v168 != 0 {
		goto L7
	} else {
		goto L56
	}
L7:
	;
	return
L8:
	;
	if v19 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	F_SetRecoveryPause(m, int32(1))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L7
	} else {
		goto L16
	}
L12:
	;
	F_errmsg(m, int32(_a_F_RecoveryRequiresIntParameter_0), int32(0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = l0
	v34 = F_errdetail(m, int32(_a_F_RecoveryRequiresIntParameter_1), v10+int32(32))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(_a_F_RecoveryRequiresIntParameter_2), int32(_a_F_RecoveryRequiresIntParameter_3), int32(_a_F_RecoveryRequiresIntParameter_4))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	goto L11
L16:
	;
	v46 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	if v46 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_errmsg(m, int32(_a_F_RecoveryRequiresIntParameter_5), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L7
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v70 = int32(0)
	goto L25
L21:
	;
	v54 = F_errdetail(m, int32(_a_F_RecoveryRequiresIntParameter_6), int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	F_errhint(m, int32(_a_F_RecoveryRequiresIntParameter_7), int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_RecoveryRequiresIntParameter_2), int32(_a_F_RecoveryRequiresIntParameter_8), int32(_a_F_RecoveryRequiresIntParameter_4))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L7
	} else {
		goto L24
	}
L24:
	;
	goto L20
L25:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_RecoveryRequiresIntParameter[1]))
	v76 = base.AtomicRmwXchg32(m, v73, int32(96), int32(1))
	if v76 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L7
	} else {
		goto L55
	}
L27:
	;
	F_s_lock(m, v73+int32(96), int32(_a_F_RecoveryRequiresIntParameter_9))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L7
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_RecoveryRequiresIntParameter[1]))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+80))
	v85 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v83)+96)), uint32(v85))
	if v84 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L29
L31:
	;
	F_ProcessStartupProcInterrupts(m)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L7
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	goto L26
L34:
	;
	v90 = F_CheckForStandbyTrigger(m)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L7
	} else {
		goto L36
	}
L35:
	;
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_RecoveryRequiresIntParameter[1]))
	v134 = base.AtomicRmwXchg32(m, v131, int32(96), int32(1))
	if v134 != 0 {
		goto L47
	} else {
		goto L48
	}
L36:
	;
	if (v90^int32(-1)|v70)&int32(1) != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v129 = v90 | v70
	goto L35
L38:
	;
	goto L39
L39:
	;
	v98 = int32(1)
	v101 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L7
	} else {
		goto L40
	}
L40:
	;
	if v101 == int32(0) {
		v129 = v98
		goto L35
	} else {
		goto L41
	}
L41:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L7
	} else {
		goto L42
	}
L42:
	;
	F_errmsg(m, int32(_a_F_RecoveryRequiresIntParameter_10), int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L7
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l0
	v118 = F_errdetail(m, int32(_a_F_RecoveryRequiresIntParameter_1), v10+int32(16))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L7
	} else {
		goto L44
	}
L44:
	;
	F_errhint(m, int32(_a_F_RecoveryRequiresIntParameter_11), int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L7
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_RecoveryRequiresIntParameter_2), int32(_a_F_RecoveryRequiresIntParameter_12), int32(_a_F_RecoveryRequiresIntParameter_4))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L7
	} else {
		goto L46
	}
L46:
	;
	v129 = v98
	goto L35
L47:
	;
	F_s_lock(m, v131+int32(96), int32(_a_F_RecoveryRequiresIntParameter_9))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L7
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v141 = *(*int32)(unsafe.Add(mBase, _c_F_RecoveryRequiresIntParameter[1]))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+80))
	if v142 == int32(1) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L49
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v141)+80)) = int32(2)
	goto L53
L52:
	;
	goto L53
L53:
	;
	v147 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v141)+96)), uint32(v147))
	v154 = F_ConditionVariableTimedSleep(m, v141+int32(84), int32(1000), int32(134217775))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L7
	} else {
		goto L54
	}
L54:
	;
	v70 = v129
	goto L25
L55:
	;
	goto L6
L56:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L7
	} else {
		goto L57
	}
L57:
	;
	F_errmsg(m, int32(_a_F_RecoveryRequiresIntParameter_13), int32(0))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L7
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
	v180 = F_errdetail(m, int32(_a_F_RecoveryRequiresIntParameter_1), v10)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L7
	} else {
		goto L59
	}
L59:
	;
	F_errhint(m, int32(_a_F_RecoveryRequiresIntParameter_14), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L7
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_RecoveryRequiresIntParameter_2), int32(_a_F_RecoveryRequiresIntParameter_15), int32(_a_F_RecoveryRequiresIntParameter_4))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L7
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ResolveRecoveryConflictWithSnapshotFullXid(m *base.Module, l0 int64, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v5 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		if base.Ui64(int64(2147483646)) < base.Ui64(v5-l0) {
			return
		} else {
			v10 = base.I32_wrap_i64(l0)
			if v10 == int32(0) {
				return
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
				v14 = F_GetConflictingVirtualXIDs(m, v10, v13)
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					F_ResolveRecoveryConflictWithVirtualXIDs(m, v14, int32(3), int32(134217772), int32(1))
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return
					} else {
						v23 = F_IsLogicalDecodingEnabled(m)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return
						} else {
							if base.B2i32(l1 == int32(0))|base.B2i32(v23 == int32(0)) != 0 {
								return
							} else {
								v30 = F_InvalidateObsoleteReplicationSlots(m, int32(2), int64(0), v13, v10)
								mBase = m.M
								v31 = m.ExcPending
								if v31 != 0 {
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
func F_SignalRecoveryConflictWithVirtualXID(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_SignalRecoveryConflictWithVirtualXID[0]))
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_SignalRecoveryConflictWithVirtualXID[1]))
	v17 = F_LWLockAcquire(m, v13+int32(512), int32(1))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v21 <= int32(0) {
		v71 = int32(0)
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_SignalRecoveryConflictWithVirtualXID[1]))
	F_LWLockRelease(m, v78+int32(512))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L14
	}
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_SignalRecoveryConflictWithVirtualXID[2]))
	v33 = int32(0)
	goto L5
L5:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(36)+v33<<(uint(int32(2))%32))))
	v46 = v30 + v43*int32(768)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+40))
	if v47 != v27 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v71 = int32(0)
	goto L3
L7:
	;
	goto L6
L8:
	;
	v64 = v33 + int32(1)
	if v64 != v21 {
		v33 = v64
		goto L5
	} else {
		goto L13
	}
L9:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)+44))
	if v49 != v26 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	if v51 == int32(0) {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v54 = int32(1)
	v58 = base.AtomicRmwOr32(m, v46, int32(340), v54<<(uint(l1)%32))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v61 = F_SendProcSignal(m, v51, int32(9), v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v71 = v54
	goto L3
L13:
	;
	goto L7
L14:
	;
	return v71
}
func F_assign_recovery_target(m *base.Module, l0 int32, l1 int32) {
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
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_assign_recovery_target[0]))
	switch v4 {
	case 0, 5:
		if l0 != 0 {
			v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if v9 != 0 {
				v11 = int32(5)
			} else {
				v11 = int32(0)
			}
		} else {
			v11 = int32(0)
		}
		*(*int32)(unsafe.Add(mBase, _c_F_assign_recovery_target[0])) = v11
		return
	default:
		F_error_multiple_recovery_targets(m)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_assign_recovery_target_xid(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_assign_recovery_target_xid[0]))
	if base.Ui32(v4) < base.Ui32(int32(2)) {
		if l0 == int32(0) {
			*(*int32)(unsafe.Add(mBase, _c_F_assign_recovery_target_xid[0])) = int32(0)
			return
		} else {
			v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if v9 == int32(0) {
				*(*int32)(unsafe.Add(mBase, _c_F_assign_recovery_target_xid[0])) = int32(0)
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_assign_recovery_target_xid[0])) = int32(1)
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				*(*int32)(unsafe.Add(mBase, _c_F_assign_recovery_target_xid[1])) = v16
				return
			}
		}
	} else {
		F_error_multiple_recovery_targets(m)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_report_recovery_conflict(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	switch l0 {
	case 0:
		F_pgstat_report_recovery_conflict(m, int32(0))
		mBase = m.M
		v4 = m.ExcPending
		if v4 != 0 {
			return
		} else {
			F_errstart_cold(m, int32(22), int32(0))
			mBase = m.M
			v8 = m.ExcPending
			if v8 != 0 {
				return
			} else {
				F_errcode(m, int32(67240389))
				mBase = m.M
				v11 = m.ExcPending
				if v11 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_report_recovery_conflict_0), int32(0))
					mBase = m.M
					v15 = m.ExcPending
					if v15 != 0 {
						return
					} else {
						F_errdetail_recovery_conflict(m, int32(0))
						mBase = m.M
						v18 = m.ExcPending
						if v18 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_report_recovery_conflict_1), int32(3324), int32(_a_F_report_recovery_conflict_2))
							mBase = m.M
							v23 = m.ExcPending
							if v23 != 0 {
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
	default:
		v25 = *(*int32)(unsafe.Add(mBase, _c_F_report_recovery_conflict[0]))
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
		if int32(1) < v26 {
			F_pgstat_report_recovery_conflict(m, l0)
			mBase = m.M
			v78 = m.ExcPending
			if v78 != 0 {
				return
			} else {
				F_errstart_cold(m, int32(22), int32(0))
				mBase = m.M
				v82 = m.ExcPending
				if v82 != 0 {
					return
				} else {
					F_errcode(m, int32(16777220))
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_report_recovery_conflict_0), int32(0))
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return
						} else {
							F_errdetail_recovery_conflict(m, l0)
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return
							} else {
								F_errhint(m, int32(_a_F_report_recovery_conflict_3), int32(0))
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_report_recovery_conflict_1), int32(3412), int32(_a_F_report_recovery_conflict_2))
									mBase = m.M
									v100 = m.ExcPending
									if v100 != 0 {
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
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, _c_F_report_recovery_conflict[0]))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+24))
			if base.B2i32((v31-int32(7))&int32(-9) == int32(0)) == int32(0) {
				v41 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_report_recovery_conflict[1])))
				if v41 != 0 {
					F_pgstat_report_recovery_conflict(m, l0)
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return
					} else {
						F_errstart_cold(m, int32(22), int32(0))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return
						} else {
							F_errcode(m, int32(16777220))
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return
							} else {
								F_errmsg(m, int32(_a_F_report_recovery_conflict_0), int32(0))
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return
								} else {
									F_errdetail_recovery_conflict(m, l0)
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return
									} else {
										F_errhint(m, int32(_a_F_report_recovery_conflict_3), int32(0))
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_report_recovery_conflict_1), int32(3412), int32(_a_F_report_recovery_conflict_2))
											mBase = m.M
											v100 = m.ExcPending
											if v100 != 0 {
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
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, _c_F_report_recovery_conflict[2]))
					if v43 == int32(0) {
						F_LockErrorCleanup(m)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							F_pgstat_report_recovery_conflict(m, l0)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return
								} else {
									F_errcode(m, int32(16777220))
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return
									} else {
										F_errmsg(m, int32(_a_F_report_recovery_conflict_4), int32(0))
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return
										} else {
											F_errdetail_recovery_conflict(m, l0)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_report_recovery_conflict_1), int32(3398), int32(_a_F_report_recovery_conflict_2))
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
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
					} else {
						v47 = *(*int32)(unsafe.Add(mBase, _c_F_report_recovery_conflict[3]))
						v48 = int32(1)
						v51 = base.AtomicRmwOr32(m, v47, int32(340), v48<<(uint(l0)%32))
						*(*int32)(unsafe.Add(mBase, _c_F_report_recovery_conflict[4])) = v48
						return
					}
				}
			} else {
				return
			}
		}
	case 4:
		v30 = *(*int32)(unsafe.Add(mBase, _c_F_report_recovery_conflict[0]))
		v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+24))
		if base.B2i32((v31-int32(7))&int32(-9) == int32(0)) == int32(0) {
			v41 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_report_recovery_conflict[1])))
			if v41 != 0 {
				F_pgstat_report_recovery_conflict(m, l0)
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return
				} else {
					F_errstart_cold(m, int32(22), int32(0))
					mBase = m.M
					v82 = m.ExcPending
					if v82 != 0 {
						return
					} else {
						F_errcode(m, int32(16777220))
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return
						} else {
							F_errmsg(m, int32(_a_F_report_recovery_conflict_0), int32(0))
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return
							} else {
								F_errdetail_recovery_conflict(m, l0)
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return
								} else {
									F_errhint(m, int32(_a_F_report_recovery_conflict_3), int32(0))
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_report_recovery_conflict_1), int32(3412), int32(_a_F_report_recovery_conflict_2))
										mBase = m.M
										v100 = m.ExcPending
										if v100 != 0 {
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
			} else {
				v43 = *(*int32)(unsafe.Add(mBase, _c_F_report_recovery_conflict[2]))
				if v43 == int32(0) {
					F_LockErrorCleanup(m)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						F_pgstat_report_recovery_conflict(m, l0)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return
							} else {
								F_errcode(m, int32(16777220))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return
								} else {
									F_errmsg(m, int32(_a_F_report_recovery_conflict_4), int32(0))
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return
									} else {
										F_errdetail_recovery_conflict(m, l0)
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_report_recovery_conflict_1), int32(3398), int32(_a_F_report_recovery_conflict_2))
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
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
				} else {
					v47 = *(*int32)(unsafe.Add(mBase, _c_F_report_recovery_conflict[3]))
					v48 = int32(1)
					v51 = base.AtomicRmwOr32(m, v47, int32(340), v48<<(uint(l0)%32))
					*(*int32)(unsafe.Add(mBase, _c_F_report_recovery_conflict[4])) = v48
					return
				}
			}
		} else {
			return
		}
	}
}
