package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CheckSubscriptionRelkind(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l0&int32(-3) != int32(112) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			F_errcode(m, int32(151027844))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
				F_errmsg(m, int32(_a_F_CheckSubscriptionRelkind_0), v7)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					F_errdetail_relkind_not_supported(m, l0)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_CheckSubscriptionRelkind_1), int32(885), int32(_a_F_CheckSubscriptionRelkind_2))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
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
		m.G0 = v7 + int32(16)
		return
	}
}
func F_DisableSubscriptionAndExit(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int64
	_ = v86
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(_a_F_DisableSubscriptionAndExit_0)
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_DisableSubscriptionAndExit[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_DisableSubscriptionAndExit[0])) = v12 + int32(1)
	F_EmitErrorReport(m)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		F_AbortOutOfAnyTransaction(m)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			F_FlushErrorState(m)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				v22 = int32(_a_F_DisableSubscriptionAndExit_0)
				v24 = *(*int32)(unsafe.Add(mBase, _c_F_DisableSubscriptionAndExit[0]))
				v25 = int32(1)
				*(*int32)(unsafe.Add(mBase, _c_F_DisableSubscriptionAndExit[0])) = v24 - v25
				v29 = *(*int32)(unsafe.Add(mBase, _c_F_DisableSubscriptionAndExit[1]))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+16)))
				if v31 == v25 {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
					v38 = base.B2i32(v34 != int32(1))
				} else {
					v38 = int32(1)
				}
				F_pgstat_report_subscription_error(m, v30, v38)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					F_StartTransactionCommand(m)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						v43 = F_GetTransactionSnapshot(m)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return
						} else {
							F_PushActiveSnapshot(m, v43)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return
							} else {
								v48 = *(*int32)(unsafe.Add(mBase, _c_F_DisableSubscriptionAndExit[2]))
								v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
								v50 = m.G0
								v52 = v50 - int32(160)
								m.G0 = v52
								v56 = F_table_open(m, int32(_a_F_DisableSubscriptionAndExit_1), int32(3))
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return
								} else {
									v60 = F_SearchSysCacheCopy(m, int32(67), v49, int32(0))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return
									} else {
										if v60 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v67 = m.ExcPending
											if v67 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v52))) = v49
												F_errmsg_internal(m, int32(_a_F_DisableSubscriptionAndExit_2), v52)
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_DisableSubscriptionAndExit_3), int32(213), int32(_a_F_DisableSubscriptionAndExit_4))
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
										} else {
											F_LockSharedObject(m, int32(_a_F_DisableSubscriptionAndExit_1), v49, int32(1))
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return
											} else {
												v82 = v52 + int32(16)
												v83 = int32(0)
												base.MemoryFill(m, v82, v83, int32(72))
												v86 = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(v52)+96)) = v86
												*(*uint16)(unsafe.Add(mBase, uint32(v52)+144)) = uint16(v83)
												*(*int64)(unsafe.Add(mBase, uint32(v52)+136)) = v86
												*(*int64)(unsafe.Add(mBase, uint32(v52)+128)) = v86
												*(*int64)(unsafe.Add(mBase, uint32(v52)+104)) = v86
												*(*uint16)(unsafe.Add(mBase, uint32(v52)+112)) = uint16(v83)
												v98 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v52)+101)) = uint8(v98)
												v100 = *(*int32)(unsafe.Add(mBase, uint32(v56)+52))
												v105 = F_heap_modify_tuple(m, v60, v100, v82, v52+int32(128), v52+int32(96))
												mBase = m.M
												v106 = m.ExcPending
												if v106 != 0 {
													return
												} else {
													F_CatalogTupleUpdate(m, v56, v105+int32(4), v105)
													mBase = m.M
													v110 = m.ExcPending
													if v110 != 0 {
														return
													} else {
														F_pfree(m, v105)
														mBase = m.M
														v112 = m.ExcPending
														if v112 != 0 {
															return
														} else {
															F_relation_close(m, v56, int32(0))
															mBase = m.M
															v115 = m.ExcPending
															if v115 != 0 {
																return
															} else {
																m.G0 = v52 + int32(160)
																F_PopActiveSnapshot(m)
																mBase = m.M
																v120 = m.ExcPending
																if v120 != 0 {
																	return
																} else {
																	F_CommitTransactionCommand(m)
																	mBase = m.M
																	v122 = m.ExcPending
																	if v122 != 0 {
																		return
																	} else {
																		v124 = *(*int32)(unsafe.Add(mBase, _c_F_DisableSubscriptionAndExit[1]))
																		v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
																		if v125 == int32(2) {
																			v128 = *(*int32)(unsafe.Add(mBase, uint32(v124)+32))
																			F_ApplyLauncherForgetWorkerStartTime(m, v128)
																			mBase = m.M
																			v130 = m.ExcPending
																			if v130 != 0 {
																				return
																			} else {
																				v133 = F_errstart(m, int32(15), int32(0))
																				mBase = m.M
																				v134 = m.ExcPending
																				if v134 != 0 {
																					return
																				} else {
																					if v133 != 0 {
																						v136 = *(*int32)(unsafe.Add(mBase, _c_F_DisableSubscriptionAndExit[2]))
																						v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)+16))
																						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v137
																						F_errmsg(m, int32(_a_F_DisableSubscriptionAndExit_5), v8)
																						mBase = m.M
																						v141 = m.ExcPending
																						if v141 != 0 {
																							return
																						} else {
																							F_errfinish(m, int32(_a_F_DisableSubscriptionAndExit_6), int32(_a_F_DisableSubscriptionAndExit_7), int32(_a_F_DisableSubscriptionAndExit_8))
																							mBase = m.M
																							v146 = m.ExcPending
																							if v146 != 0 {
																								return
																							} else {
																								F_proc_exit(m, int32(0))
																								mBase = m.M
																								v149 = m.ExcPending
																								if v149 != 0 {
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
																						v149 = m.ExcPending
																						if v149 != 0 {
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
																			v133 = F_errstart(m, int32(15), int32(0))
																			mBase = m.M
																			v134 = m.ExcPending
																			if v134 != 0 {
																				return
																			} else {
																				if v133 != 0 {
																					v136 = *(*int32)(unsafe.Add(mBase, _c_F_DisableSubscriptionAndExit[2]))
																					v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)+16))
																					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v137
																					F_errmsg(m, int32(_a_F_DisableSubscriptionAndExit_5), v8)
																					mBase = m.M
																					v141 = m.ExcPending
																					if v141 != 0 {
																						return
																					} else {
																						F_errfinish(m, int32(_a_F_DisableSubscriptionAndExit_6), int32(_a_F_DisableSubscriptionAndExit_7), int32(_a_F_DisableSubscriptionAndExit_8))
																						mBase = m.M
																						v146 = m.ExcPending
																						if v146 != 0 {
																							return
																						} else {
																							F_proc_exit(m, int32(0))
																							mBase = m.M
																							v149 = m.ExcPending
																							if v149 != 0 {
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
																					v149 = m.ExcPending
																					if v149 != 0 {
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
func F_RemoveSubscriptionRel(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	v8 = m.G0
	v10 = v8 - int32(144)
	m.G0 = v10
	v14 = F_table_open(m, int32(_a_F_RemoveSubscriptionRel_0), int32(3))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if l0 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v16 = int32(1)
	F_ScanKeyInit(m, v10+int32(48), v16, int32(3), int32(184), l0)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v28 = int32(0)
	v29 = v10 + int32(48)
	goto L5
L5:
	;
	if l1 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v28 = v16
	v29 = v10 + int32(96)
	goto L5
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L28
	}
L8:
	;
	F_ScanKeyInit(m, v29, int32(2), int32(3), int32(184), l1)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	v37 = v28
	goto L10
L10:
	;
	v40 = F_table_beginscan_catalog(m, v14, v37, v10+int32(48))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	v37 = v28 + int32(1)
	goto L10
L12:
	;
	v42 = F_heap_getnext(m, v40)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	if v42 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v47 = v42
	goto L17
L15:
	;
	goto L16
L16:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+188))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
	m.T0[v75].(func(*base.Module, int32))(m, v40)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L26
	}
L17:
	;
	if l0 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L16
L19:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+22)))
	v55 = v53 + v54
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+8)))
	if v56 != int32(114) {
		goto L7
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	F_simple_heap_delete(m, v14, v47+int32(4))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L23
	}
L22:
	;
	goto L21
L23:
	;
	v64 = F_heap_getnext(m, v40)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if v64 != 0 {
		v47 = v64
		goto L17
	} else {
		goto L25
	}
L25:
	;
	goto L18
L26:
	;
	F_relation_close(m, v14, int32(3))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	m.G0 = v10 + int32(144)
	return
L28:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v93 = F_get_subscription_name(m, v91, int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v93
	F_errmsg(m, int32(_a_F_RemoveSubscriptionRel_1), v10+int32(32))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v101 = F_get_rel_name(m, l1)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v103 = int32(*(*int8)(unsafe.Add(mBase, uint32(v55)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v101
	F_errdetail(m, int32(_a_F_RemoveSubscriptionRel_2), v10+int32(16))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(_a_F_RemoveSubscriptionRel_3)
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_RemoveSubscriptionRel_4)
	F_errhint(m, int32(_a_F_RemoveSubscriptionRel_5), v10)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_RemoveSubscriptionRel_6), int32(495), int32(_a_F_RemoveSubscriptionRel_7))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_subscription_name(m *base.Module, l0 int32, l1 int32) int32 {
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v8 = Fn13924(m, l0, l1, int32(16), int32(_a_F_get_subscription_name_0), int32(3846), int32(_a_F_get_subscription_name_1), int32(67))
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
func F_parse_subscription_options_x2especialized_x2e1(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
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
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	v12 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l2)+16)) = v12
	*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = v12
	*(*int64)(unsafe.Add(mBase, uint32(l2)+32)) = v12
	*(*int64)(unsafe.Add(mBase, uint32(l2)+24)) = v12
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v12
	v22 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+16)) = uint8(v22)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+15)) = uint8(v22)
	if l1 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = int32(_a_F_parse_subscription_options_x2especialized_x2e1_0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = int32(_a_F_parse_subscription_options_x2especialized_x2e1_1)
	F_errmsg(m, int32(_a_F_parse_subscription_options_x2especialized_x2e1_2), v10+int32(48))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L26
	} else {
		goto L62
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(_a_F_parse_subscription_options_x2especialized_x2e1_3)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(_a_F_parse_subscription_options_x2especialized_x2e1_1)
	F_errmsg(m, int32(_a_F_parse_subscription_options_x2especialized_x2e1_2), v10+int32(16))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L26
	} else {
		goto L60
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L26
	} else {
		goto L56
	}
L4:
	;
	F_errorConflictingDefElem(m, v40, l0)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L26
	} else {
		goto L55
	}
L5:
	;
	m.G0 = v10 + int32(80)
	return
L6:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v26 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v136&int32(8) == int32(0) {
		goto L5
	} else {
		goto L40
	}
L9:
	;
	v35 = int32(0)
	goto L12
L10:
	;
	goto L11
L11:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v128 != 0 {
		goto L5
	} else {
		goto L39
	}
L12:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36+v35<<(uint(int32(2))%32))))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	v42 = int32(_a_F_parse_subscription_options_x2especialized_x2e1_4)
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_subscription_options_x2especialized_x2e1[0])))
	if base.B2i32(v45 == int32(0))|base.B2i32(v45 != v48) != 0 {
		v66 = v45
		v67 = v48
		goto L16
	} else {
		goto L17
	}
L13:
	;
	goto L11
L14:
	;
	v118 = v35 + int32(1)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v118 < v119 {
		v35 = v118
		goto L12
	} else {
		goto L38
	}
L15:
	;
	if v66-v67 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L16:
	;
	goto L15
L17:
	;
	v51 = v41
	v52 = v42
	goto L18
L18:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+1)))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
	if v56 == int32(0) {
		v66 = v56
		v67 = v55
		goto L16
	} else {
		goto L20
	}
L19:
	;
	v66 = v56
	v67 = v55
	goto L16
L20:
	;
	v59 = int32(1)
	if v56 == v55 {
		v51 = v51 + v59
		v52 = v52 + v59
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v71&int32(16) != 0 {
		goto L4
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v80 = int32(_a_F_parse_subscription_options_x2especialized_x2e1_5)
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_subscription_options_x2especialized_x2e1[1])))
	if base.B2i32(v83 == int32(0))|base.B2i32(v83 != v86) != 0 {
		v104 = v83
		v105 = v86
		goto L29
	} else {
		goto L30
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v71 | int32(16)
	v77 = F_defGetBoolean(m, v40)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	return
L27:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+15)) = uint8(v77)
	goto L14
L28:
	;
	if v104-v105 != 0 {
		goto L3
	} else {
		goto L35
	}
L29:
	;
	goto L28
L30:
	;
	v89 = v41
	v90 = v80
	goto L31
L31:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+1)))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+1)))
	if v94 == int32(0) {
		v104 = v94
		v105 = v93
		goto L29
	} else {
		goto L33
	}
L32:
	;
	v104 = v94
	v105 = v93
	goto L29
L33:
	;
	v97 = int32(1)
	if v94 == v93 {
		v89 = v89 + v97
		v90 = v90 + v97
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v107&int32(64) != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v107 | int32(64)
	v113 = F_defGetBoolean(m, v40)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L26
	} else {
		goto L37
	}
L37:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+16)) = uint8(v113)
	goto L14
L38:
	;
	goto L13
L39:
	;
	goto L8
L40:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+13)))
	if v141 == int32(1) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L26
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+14)))
	if v165 != int32(1) {
		goto L5
	} else {
		goto L49
	}
L44:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L26
	} else {
		goto L45
	}
L45:
	;
	if v136&int32(2) != 0 {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(_a_F_parse_subscription_options_x2especialized_x2e1_6)
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_parse_subscription_options_x2especialized_x2e1_1)
	F_errmsg(m, int32(_a_F_parse_subscription_options_x2especialized_x2e1_7), v10)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L26
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_parse_subscription_options_x2especialized_x2e1_8), int32(421), int32(_a_F_parse_subscription_options_x2especialized_x2e1_9))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L26
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L26
	} else {
		goto L50
	}
L50:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L26
	} else {
		goto L51
	}
L51:
	;
	if v136&int32(4) != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = int32(_a_F_parse_subscription_options_x2especialized_x2e1_10)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = int32(_a_F_parse_subscription_options_x2especialized_x2e1_1)
	F_errmsg(m, int32(_a_F_parse_subscription_options_x2especialized_x2e1_7), v10+int32(32))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L26
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_parse_subscription_options_x2especialized_x2e1_8), int32(437), int32(_a_F_parse_subscription_options_x2especialized_x2e1_9))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L26
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L26
	} else {
		goto L57
	}
L57:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v211
	F_errmsg(m, int32(_a_F_parse_subscription_options_x2especialized_x2e1_11), v10-int32(-64))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L26
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_parse_subscription_options_x2especialized_x2e1_8), int32(363), int32(_a_F_parse_subscription_options_x2especialized_x2e1_9))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L26
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L60:
	;
	F_errfinish(m, int32(_a_F_parse_subscription_options_x2especialized_x2e1_8), int32(415), int32(_a_F_parse_subscription_options_x2especialized_x2e1_9))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L26
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	F_errfinish(m, int32(_a_F_parse_subscription_options_x2especialized_x2e1_8), int32(431), int32(_a_F_parse_subscription_options_x2especialized_x2e1_9))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L26
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
