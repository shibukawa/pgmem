package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
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
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
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
		v24 = int64(*(*int32)(unsafe.Add(mBase, _consts[138])))
		v25 = base.I64_div_u_s(v22, v24)
		v27 = base.I64_div_u_s(int64(4294967296), v24)
		v28 = base.I64_div_u_s(v25, v27)
		*(*uint32)(unsafe.Add(mBase, uint32(v12)+52)) = uint32(v28)
		v31 = v25 - v27*v28
		*(*uint32)(unsafe.Add(mBase, uint32(v12)+56)) = uint32(v31)
		v39 = F_pg_snprintf(m, v12+int32(80), int32(64), int32(511295), v12+int32(48))
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v12 + int32(80)
			v47 = F_replace_percent_placeholders(m, l0, l1, int32(231353), v12+int32(32))
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return
			} else {
				v51 = F_errstart(m, int32(12), int32(0))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return
				} else {
					if v51 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l1
						F_errmsg_internal(m, int32(701322), v12+int32(16))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return
						} else {
							F_errfinish(m, int32(499597), int32(323), int32(430114))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return
							} else {
								v66 = F_fflush(m, int32(0))
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return
								} else {
									v69 = *(*int32)(unsafe.Add(mBase, _consts[127]))
									*(*int32)(unsafe.Add(mBase, uint32(v69))) = l3
									v71 = F_pgl_system(m, v47)
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return
									} else {
										v74 = *(*int32)(unsafe.Add(mBase, _consts[127]))
										*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(0)
										F_pfree(m, v47)
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return
										} else {
											if v71 == int32(0) {
												m.G0 = v12 + int32(1104)
												return
											} else {
												if l2 != 0 {
													v85 = int32(1)
													if base.Ui32(v71&int32(65535)-v85) < base.Ui32(int32(255)) {
														v103 = v85
													} else {
														if v71&int32(127) == int32(0) {
															if base.Ui32(int32(125)) < base.Ui32(int32(base.Ui32(v71)>>(uint(int32(8))%32))&int32(255)) {
																v103 = v85
															} else {
																v103 = int32(0)
															}
														} else {
															v103 = int32(0)
														}
													}
													if v103 != 0 {
														v104 = int32(22)
													} else {
														v104 = int32(19)
													}
													v105 = v104
												} else {
													v105 = int32(19)
												}
												v107 = F_errstart(m, v105, int32(0))
												mBase = m.M
												v108 = m.ExcPending
												if v108 != 0 {
													return
												} else {
													if v107 == int32(0) {
														m.G0 = v12 + int32(1104)
														return
													} else {
														v111 = F_wait_result_to_str(m, v71)
														mBase = m.M
														v112 = m.ExcPending
														if v112 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v111
															*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = l0
															*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
															F_errmsg(m, int32(205850), v12)
															mBase = m.M
															v118 = m.ExcPending
															if v118 != 0 {
																return
															} else {
																F_errfinish(m, int32(499597), int32(347), int32(430114))
																mBase = m.M
																v123 = m.ExcPending
																if v123 != 0 {
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
						v66 = F_fflush(m, int32(0))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return
						} else {
							v69 = *(*int32)(unsafe.Add(mBase, _consts[127]))
							*(*int32)(unsafe.Add(mBase, uint32(v69))) = l3
							v71 = F_pgl_system(m, v47)
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return
							} else {
								v74 = *(*int32)(unsafe.Add(mBase, _consts[127]))
								*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(0)
								F_pfree(m, v47)
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return
								} else {
									if v71 == int32(0) {
										m.G0 = v12 + int32(1104)
										return
									} else {
										if l2 != 0 {
											v85 = int32(1)
											if base.Ui32(v71&int32(65535)-v85) < base.Ui32(int32(255)) {
												v103 = v85
											} else {
												if v71&int32(127) == int32(0) {
													if base.Ui32(int32(125)) < base.Ui32(int32(base.Ui32(v71)>>(uint(int32(8))%32))&int32(255)) {
														v103 = v85
													} else {
														v103 = int32(0)
													}
												} else {
													v103 = int32(0)
												}
											}
											if v103 != 0 {
												v104 = int32(22)
											} else {
												v104 = int32(19)
											}
											v105 = v104
										} else {
											v105 = int32(19)
										}
										v107 = F_errstart(m, v105, int32(0))
										mBase = m.M
										v108 = m.ExcPending
										if v108 != 0 {
											return
										} else {
											if v107 == int32(0) {
												m.G0 = v12 + int32(1104)
												return
											} else {
												v111 = F_wait_result_to_str(m, v71)
												mBase = m.M
												v112 = m.ExcPending
												if v112 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v111
													*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = l0
													*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
													F_errmsg(m, int32(205850), v12)
													mBase = m.M
													v118 = m.ExcPending
													if v118 != 0 {
														return
													} else {
														F_errfinish(m, int32(499597), int32(347), int32(430114))
														mBase = m.M
														v123 = m.ExcPending
														if v123 != 0 {
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
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
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
	v14 = int32(*(*uint8)(unsafe.Add(mBase, _consts[262])))
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
	v174 = m.ExcPending
	if v174 != 0 {
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
	F_errmsg(m, int32(156146), int32(0))
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
	F_errdetail(m, int32(654618), v10+int32(32))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(493460), int32(4715), int32(217324))
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
	F_errmsg(m, int32(450013), int32(0))
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
	F_errdetail(m, int32(614064), int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	F_errhint(m, int32(602864), int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(493460), int32(4722), int32(217324))
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
	v73 = *(*int32)(unsafe.Add(mBase, _consts[259]))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v73)+96)) = int32(1)
	if v74 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L7
	} else {
		goto L55
	}
L27:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _consts[259]))
	F_s_lock(m, v78+int32(96), int32(493460), int32(3096), int32(354902))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L7
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _consts[259]))
	*(*int32)(unsafe.Add(mBase, uint32(v87)+96)) = int32(0)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v87)+80))
	if v90 != 0 {
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
	v92 = m.ExcPending
	if v92 != 0 {
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
	v93 = F_CheckForStandbyTrigger(m)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L7
	} else {
		goto L36
	}
L35:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _consts[259]))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v134)+96)) = int32(1)
	if v135 != 0 {
		goto L47
	} else {
		goto L48
	}
L36:
	;
	if (v93^int32(-1)|v70)&int32(1) != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v132 = v93 | v70
	goto L35
L38:
	;
	goto L39
L39:
	;
	v101 = int32(1)
	v104 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L7
	} else {
		goto L40
	}
L40:
	;
	if v104 == int32(0) {
		v132 = v101
		goto L35
	} else {
		goto L41
	}
L41:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L7
	} else {
		goto L42
	}
L42:
	;
	F_errmsg(m, int32(156217), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L7
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l0
	F_errdetail(m, int32(654618), v10+int32(16))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L7
	} else {
		goto L44
	}
L44:
	;
	F_errhint(m, int32(603023), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L7
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(493460), int32(4743), int32(217324))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L7
	} else {
		goto L46
	}
L46:
	;
	v132 = v101
	goto L35
L47:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _consts[259]))
	F_s_lock(m, v139+int32(96), int32(493460), int32(3135), int32(450033))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L7
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _consts[259]))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+80))
	if v149 == int32(1) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L49
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v148)+80)) = int32(2)
	goto L53
L52:
	;
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v148)+96)) = int32(0)
	v160 = F_ConditionVariableTimedSleep(m, v148+int32(84), int32(1000), int32(134217775))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L7
	} else {
		goto L54
	}
L54:
	;
	v70 = v132
	goto L25
L55:
	;
	goto L6
L56:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L7
	} else {
		goto L57
	}
L57:
	;
	F_errmsg(m, int32(156286), int32(0))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L7
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
	F_errdetail(m, int32(654618), v10)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L7
	} else {
		goto L59
	}
L59:
	;
	F_errhint(m, int32(602946), int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L7
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(493460), int32(4773), int32(217324))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
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
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
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
					F_ResolveRecoveryConflictWithVirtualXIDs(m, v14, int32(10), int32(134217772), int32(1))
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return
					} else {
						if l1 == int32(0) {
							return
						} else {
							v24 = *(*int32)(unsafe.Add(mBase, _consts[2]))
							if v24 < int32(2) {
								return
							} else {
								v29 = F_InvalidateObsoleteReplicationSlots(m, int32(2), int64(0), v13, v10)
								mBase = m.M
								v30 = m.ExcPending
								if v30 != 0 {
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
	v4 = *(*int32)(unsafe.Add(mBase, _consts[263]))
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
		*(*int32)(unsafe.Add(mBase, _consts[263])) = v11
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
	v4 = *(*int32)(unsafe.Add(mBase, _consts[263]))
	if base.Ui32(v4) < base.Ui32(int32(2)) {
		if l0 == int32(0) {
			*(*int32)(unsafe.Add(mBase, _consts[263])) = int32(0)
			return
		} else {
			v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if v9 == int32(0) {
				*(*int32)(unsafe.Add(mBase, _consts[263])) = int32(0)
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[263])) = int32(1)
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				*(*int32)(unsafe.Add(mBase, _consts[264])) = v16
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
