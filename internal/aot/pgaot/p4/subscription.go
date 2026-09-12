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
				F_errmsg(m, int32(108320), v7)
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
						F_errfinish(m, int32(496305), int32(885), int32(425312))
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
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int64
	_ = v91
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(4510044)
	v12 = *(*int32)(unsafe.Add(mBase, _consts[171]))
	*(*int32)(unsafe.Add(mBase, _consts[171])) = v12 + int32(1)
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
				v22 = int32(4510044)
				v24 = *(*int32)(unsafe.Add(mBase, _consts[171]))
				v25 = int32(1)
				*(*int32)(unsafe.Add(mBase, _consts[171])) = v24 - v25
				v29 = *(*int32)(unsafe.Add(mBase, _consts[828]))
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
								v48 = *(*int32)(unsafe.Add(mBase, _consts[832]))
								v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
								v50 = m.G0
								v52 = v50 - int32(160)
								m.G0 = v52
								v56 = F_table_open(m, int32(6100), int32(3))
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
												F_errmsg_internal(m, int32(44593), v52)
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return
												} else {
													F_errfinish(m, int32(496196), int32(213), int32(247478))
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
											F_LockSharedObject(m, int32(6100), v49, int32(1))
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return
											} else {
												v86 = F__emscripten_memset_bulkmem(m, v52+int32(16), base.I32_extend8_s(int32(0)), int32(72))
												mBase = m.M
												v87 = int32(0)
												*(*uint16)(unsafe.Add(mBase, uint32(v52)+144)) = uint16(v87)
												*(*uint16)(unsafe.Add(mBase, uint32(v52)+112)) = uint16(v87)
												v91 = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(v52)+96)) = v91
												*(*int64)(unsafe.Add(mBase, uint32(v52)+136)) = v91
												*(*int64)(unsafe.Add(mBase, uint32(v52)+128)) = v91
												*(*int64)(unsafe.Add(mBase, uint32(v52)+104)) = v91
												v99 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v52)+101)) = uint8(v99)
												v101 = *(*int32)(unsafe.Add(mBase, uint32(v56)+52))
												v108 = F_heap_modify_tuple(m, v60, v101, v52+int32(16), v52+int32(128), v52+int32(96))
												mBase = m.M
												v109 = m.ExcPending
												if v109 != 0 {
													return
												} else {
													F_CatalogTupleUpdate(m, v56, v108+int32(4), v108)
													mBase = m.M
													v113 = m.ExcPending
													if v113 != 0 {
														return
													} else {
														F_pfree(m, v108)
														mBase = m.M
														v115 = m.ExcPending
														if v115 != 0 {
															return
														} else {
															F_sequence_close(m, v56, int32(0))
															mBase = m.M
															v118 = m.ExcPending
															if v118 != 0 {
																return
															} else {
																m.G0 = v52 + int32(160)
																F_PopActiveSnapshot(m)
																mBase = m.M
																v123 = m.ExcPending
																if v123 != 0 {
																	return
																} else {
																	F_CommitTransactionCommand(m)
																	mBase = m.M
																	v125 = m.ExcPending
																	if v125 != 0 {
																		return
																	} else {
																		v127 = *(*int32)(unsafe.Add(mBase, _consts[828]))
																		v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
																		if v128 == int32(2) {
																			v131 = *(*int32)(unsafe.Add(mBase, uint32(v127)+32))
																			F_ApplyLauncherForgetWorkerStartTime(m, v131)
																			mBase = m.M
																			v133 = m.ExcPending
																			if v133 != 0 {
																				return
																			} else {
																				v136 = F_errstart(m, int32(15), int32(0))
																				mBase = m.M
																				v137 = m.ExcPending
																				if v137 != 0 {
																					return
																				} else {
																					if v136 != 0 {
																						v139 = *(*int32)(unsafe.Add(mBase, _consts[832]))
																						v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+16))
																						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v140
																						F_errmsg(m, int32(212273), v8)
																						mBase = m.M
																						v144 = m.ExcPending
																						if v144 != 0 {
																							return
																						} else {
																							F_errfinish(m, int32(495381), int32(4876), int32(99635))
																							mBase = m.M
																							v149 = m.ExcPending
																							if v149 != 0 {
																								return
																							} else {
																								F_proc_exit(m, int32(0))
																								mBase = m.M
																								v152 = m.ExcPending
																								if v152 != 0 {
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
																						v152 = m.ExcPending
																						if v152 != 0 {
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
																			v136 = F_errstart(m, int32(15), int32(0))
																			mBase = m.M
																			v137 = m.ExcPending
																			if v137 != 0 {
																				return
																			} else {
																				if v136 != 0 {
																					v139 = *(*int32)(unsafe.Add(mBase, _consts[832]))
																					v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+16))
																					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v140
																					F_errmsg(m, int32(212273), v8)
																					mBase = m.M
																					v144 = m.ExcPending
																					if v144 != 0 {
																						return
																					} else {
																						F_errfinish(m, int32(495381), int32(4876), int32(99635))
																						mBase = m.M
																						v149 = m.ExcPending
																						if v149 != 0 {
																							return
																						} else {
																							F_proc_exit(m, int32(0))
																							mBase = m.M
																							v152 = m.ExcPending
																							if v152 != 0 {
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
																					v152 = m.ExcPending
																					if v152 != 0 {
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
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
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
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	v8 = m.G0
	v10 = v8 - int32(144)
	m.G0 = v10
	v14 = F_table_open(m, int32(6102), int32(3))
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
	if l0 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if l1 != 0 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	v30 = int32(0)
	v31 = v10 + int32(48)
	goto L3
L5:
	;
	goto L6
L6:
	;
	v20 = int32(1)
	F_ScanKeyInit(m, v10+int32(48), v20, int32(3), int32(184), l0)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v30 = v20
	v31 = v10 + int32(96)
	goto L3
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L29
	}
L9:
	;
	F_ScanKeyInit(m, v31, int32(2), int32(3), int32(184), l1)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	v39 = v30
	goto L11
L11:
	;
	v42 = F_table_beginscan_catalog(m, v14, v39, v10+int32(48))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	v39 = v30 + int32(1)
	goto L11
L13:
	;
	v44 = F_heap_getnext(m, v42)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v44 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v49 = v44
	goto L18
L16:
	;
	goto L17
L17:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+188))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
	m.T0[v77].(func(*base.Module, int32))(m, v42)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L27
	}
L18:
	;
	if l0 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L17
L20:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+22)))
	v57 = v55 + v56
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+8)))
	if v58 != int32(114) {
		goto L8
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	F_CatalogTupleDelete(m, v14, v49+int32(4))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	v66 = F_heap_getnext(m, v42)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	if v66 != 0 {
		v49 = v66
		goto L18
	} else {
		goto L26
	}
L26:
	;
	goto L19
L27:
	;
	F_sequence_close(m, v14, int32(3))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	m.G0 = v10 + int32(144)
	return
L29:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v95 = F_get_subscription_name(m, v93, int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v95
	F_errmsg(m, int32(702759), v10+int32(32))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v103 = F_get_rel_name(m, l1)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v105 = int32(*(*int8)(unsafe.Add(mBase, uint32(v57)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v103
	F_errdetail(m, int32(669079), v10+int32(16))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(660389)
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(541052)
	F_errhint(m, int32(613303), v10)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(496196), int32(495), int32(308325))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_subscription_name(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = F_SearchSysCache1(m, int32(67), l0)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 == int32(0) {
			if l1 != 0 {
				v40 = int32(0)
				m.G0 = v8 + int32(16)
				return v40
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
					F_errmsg_internal(m, int32(44593), v8)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(499145), int32(3846), int32(378762))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
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
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+22)))
			v35 = F_pstrdup(m, v30+v31+int32(16))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				F_ReleaseCatCache(m, v11)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					v40 = v35
					m.G0 = v8 + int32(16)
					return v40
				}
			}
		}
	}
}
