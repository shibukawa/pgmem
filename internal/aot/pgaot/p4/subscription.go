package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CheckSubscriptionRelkind(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	switch l0 - int32(112) {
	case 0, 2:
		v14 = int32(83)
		v15 = base.B2i32(l0 == v14)
		v17 = base.B2i32(l1 == v14)
		if v15 != v17 {
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
					if l0 == v14 {
						v50 = int32(_a_F_CheckSubscriptionRelkind_0)
					} else {
						v50 = int32(_a_F_CheckSubscriptionRelkind_1)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v50
					if l1 == v14 {
						v54 = int32(_a_F_CheckSubscriptionRelkind_0)
					} else {
						v54 = int32(_a_F_CheckSubscriptionRelkind_1)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v54
					*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l3
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l2
					F_errmsg(m, int32(_a_F_CheckSubscriptionRelkind_2), v8+int32(16))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_CheckSubscriptionRelkind_3), int32(1165), int32(_a_F_CheckSubscriptionRelkind_4))
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
		} else {
			m.G0 = v8 + int32(32)
			return
		}
	case 1:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			F_errcode(m, int32(151027844))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = l2
				F_errmsg(m, int32(_a_F_CheckSubscriptionRelkind_5), v8)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					F_errdetail_relkind_not_supported(m, l0)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_CheckSubscriptionRelkind_3), int32(1150), int32(_a_F_CheckSubscriptionRelkind_4))
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
			}
		}
	default:
		if l0 != int32(83) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				F_errcode(m, int32(151027844))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l3
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l2
					F_errmsg(m, int32(_a_F_CheckSubscriptionRelkind_5), v8)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						F_errdetail_relkind_not_supported(m, l0)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_CheckSubscriptionRelkind_3), int32(1150), int32(_a_F_CheckSubscriptionRelkind_4))
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
				}
			}
		} else {
			v14 = int32(83)
			v15 = base.B2i32(l0 == v14)
			v17 = base.B2i32(l1 == v14)
			if v15 != v17 {
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
						if l0 == v14 {
							v50 = int32(_a_F_CheckSubscriptionRelkind_0)
						} else {
							v50 = int32(_a_F_CheckSubscriptionRelkind_1)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v50
						if l1 == v14 {
							v54 = int32(_a_F_CheckSubscriptionRelkind_0)
						} else {
							v54 = int32(_a_F_CheckSubscriptionRelkind_1)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v54
						*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l3
						*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l2
						F_errmsg(m, int32(_a_F_CheckSubscriptionRelkind_2), v8+int32(16))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_CheckSubscriptionRelkind_3), int32(1165), int32(_a_F_CheckSubscriptionRelkind_4))
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
			} else {
				m.G0 = v8 + int32(32)
				return
			}
		}
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int64
	_ = v79
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
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
				*(*int32)(unsafe.Add(mBase, _c_F_DisableSubscriptionAndExit[0])) = v24 - int32(1)
				v29 = *(*int32)(unsafe.Add(mBase, _c_F_DisableSubscriptionAndExit[1]))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
				F_pgstat_report_subscription_error(m, v30)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					F_StartTransactionCommand(m)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						v35 = F_GetTransactionSnapshot(m)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							F_PushActiveSnapshot(m, v35)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return
							} else {
								v40 = *(*int32)(unsafe.Add(mBase, _c_F_DisableSubscriptionAndExit[2]))
								v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
								v42 = m.G0
								v44 = v42 - int32(272)
								m.G0 = v44
								v48 = F_table_open(m, int32(_a_F_DisableSubscriptionAndExit_1), int32(3))
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return
								} else {
									v53 = F_SearchSysCacheCopy(m, int32(67), base.I64_extend_i32_u(v41), int64(0))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return
									} else {
										if v53 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v60 = m.ExcPending
											if v60 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v44))) = v41
												F_errmsg_internal(m, int32(_a_F_DisableSubscriptionAndExit_2), v44)
												mBase = m.M
												v64 = m.ExcPending
												if v64 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_DisableSubscriptionAndExit_3), int32(286), int32(_a_F_DisableSubscriptionAndExit_4))
													mBase = m.M
													v69 = m.ExcPending
													if v69 != 0 {
														return
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										} else {
											F_LockSharedObject(m, int32(_a_F_DisableSubscriptionAndExit_1), v41, int32(1))
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return
											} else {
												v75 = v44 + int32(16)
												base.MemoryFill(m, v75, int32(0), int32(184))
												v79 = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(v44)+208)) = v79
												*(*int64)(unsafe.Add(mBase, uint32(v44)+255)) = v79
												*(*int64)(unsafe.Add(mBase, uint32(v44)+248)) = v79
												*(*int64)(unsafe.Add(mBase, uint32(v44)+240)) = v79
												*(*int64)(unsafe.Add(mBase, uint32(v44)+216)) = v79
												*(*int64)(unsafe.Add(mBase, uint32(v44)+223)) = v79
												v91 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v44)+213)) = uint8(v91)
												v93 = *(*int32)(unsafe.Add(mBase, uint32(v48)+52))
												v98 = F_heap_modify_tuple(m, v53, v93, v75, v44+int32(240), v44+int32(208))
												mBase = m.M
												v99 = m.ExcPending
												if v99 != 0 {
													return
												} else {
													F_CatalogTupleUpdate(m, v48, v98+int32(4), v98)
													mBase = m.M
													v103 = m.ExcPending
													if v103 != 0 {
														return
													} else {
														F_pfree(m, v98)
														mBase = m.M
														v105 = m.ExcPending
														if v105 != 0 {
															return
														} else {
															F_relation_close(m, v48, int32(0))
															mBase = m.M
															v108 = m.ExcPending
															if v108 != 0 {
																return
															} else {
																m.G0 = v44 + int32(272)
																F_PopActiveSnapshot(m)
																mBase = m.M
																v113 = m.ExcPending
																if v113 != 0 {
																	return
																} else {
																	F_CommitTransactionCommand(m)
																	mBase = m.M
																	v115 = m.ExcPending
																	if v115 != 0 {
																		return
																	} else {
																		v117 = *(*int32)(unsafe.Add(mBase, _c_F_DisableSubscriptionAndExit[1]))
																		v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
																		if v118 == int32(3) {
																			v121 = *(*int32)(unsafe.Add(mBase, uint32(v117)+32))
																			F_ApplyLauncherForgetWorkerStartTime(m, v121)
																			mBase = m.M
																			v123 = m.ExcPending
																			if v123 != 0 {
																				return
																			} else {
																				v126 = F_errstart(m, int32(15), int32(0))
																				mBase = m.M
																				v127 = m.ExcPending
																				if v127 != 0 {
																					return
																				} else {
																					if v126 != 0 {
																						v129 = *(*int32)(unsafe.Add(mBase, _c_F_DisableSubscriptionAndExit[2]))
																						v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+24))
																						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v130
																						F_errmsg(m, int32(_a_F_DisableSubscriptionAndExit_5), v8)
																						mBase = m.M
																						v134 = m.ExcPending
																						if v134 != 0 {
																							return
																						} else {
																							F_errfinish(m, int32(_a_F_DisableSubscriptionAndExit_6), int32(_a_F_DisableSubscriptionAndExit_7), int32(_a_F_DisableSubscriptionAndExit_8))
																							mBase = m.M
																							v139 = m.ExcPending
																							if v139 != 0 {
																								return
																							} else {
																								v140 = int32(0)
																								v144 = *(*int32)(unsafe.Add(mBase, _c_F_DisableSubscriptionAndExit[2]))
																								v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+41)))
																								v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+48)))
																								F_CheckSubDeadTupleRetention(m, v140, int32(1), int32(19), v145, v146, v140)
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
																						}
																					} else {
																						v140 = int32(0)
																						v144 = *(*int32)(unsafe.Add(mBase, _c_F_DisableSubscriptionAndExit[2]))
																						v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+41)))
																						v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+48)))
																						F_CheckSubDeadTupleRetention(m, v140, int32(1), int32(19), v145, v146, v140)
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
																				}
																			}
																		} else {
																			v126 = F_errstart(m, int32(15), int32(0))
																			mBase = m.M
																			v127 = m.ExcPending
																			if v127 != 0 {
																				return
																			} else {
																				if v126 != 0 {
																					v129 = *(*int32)(unsafe.Add(mBase, _c_F_DisableSubscriptionAndExit[2]))
																					v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+24))
																					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v130
																					F_errmsg(m, int32(_a_F_DisableSubscriptionAndExit_5), v8)
																					mBase = m.M
																					v134 = m.ExcPending
																					if v134 != 0 {
																						return
																					} else {
																						F_errfinish(m, int32(_a_F_DisableSubscriptionAndExit_6), int32(_a_F_DisableSubscriptionAndExit_7), int32(_a_F_DisableSubscriptionAndExit_8))
																						mBase = m.M
																						v139 = m.ExcPending
																						if v139 != 0 {
																							return
																						} else {
																							v140 = int32(0)
																							v144 = *(*int32)(unsafe.Add(mBase, _c_F_DisableSubscriptionAndExit[2]))
																							v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+41)))
																							v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+48)))
																							F_CheckSubDeadTupleRetention(m, v140, int32(1), int32(19), v145, v146, v140)
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
																					}
																				} else {
																					v140 = int32(0)
																					v144 = *(*int32)(unsafe.Add(mBase, _c_F_DisableSubscriptionAndExit[2]))
																					v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+41)))
																					v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+48)))
																					F_CheckSubDeadTupleRetention(m, v140, int32(1), int32(19), v145, v146, v140)
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
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	v8 = m.G0
	v10 = v8 - int32(160)
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
	F_ScanKeyInit(m, v10+int32(48), v16, int32(3), int32(184), base.I64_extend_i32_u(l0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v29 = int32(0)
	v30 = v10 + int32(48)
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
	v29 = v16
	v30 = v10 + int32(104)
	goto L5
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L29
	}
L8:
	;
	F_ScanKeyInit(m, v30, int32(2), int32(3), int32(184), base.I64_extend_i32_u(l1))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	v39 = v29
	goto L10
L10:
	;
	v42 = F_table_beginscan_catalog(m, v14, v39, v10+int32(48))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	v39 = v29 + int32(1)
	goto L10
L12:
	;
	v44 = F_heap_getnext(m, v42)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	if v44 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v49 = v44
	goto L17
L15:
	;
	goto L16
L16:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+188))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	m.T0[v80].(func(*base.Module, int32))(m, v42)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L27
	}
L17:
	;
	if l0 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L16
L19:
	;
	F_simple_heap_delete(m, v14, v49+int32(4))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L24
	}
L20:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+22)))
	v55 = v53 + v54
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+8)))
	if v56 == int32(114) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	v60 = F_get_rel_relkind(m, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	if v60 != int32(83) {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	goto L19
L24:
	;
	v69 = F_heap_getnext(m, v42)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	if v69 != 0 {
		v49 = v69
		goto L17
	} else {
		goto L26
	}
L26:
	;
	goto L18
L27:
	;
	F_relation_close(m, v14, int32(3))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	m.G0 = v10 + int32(160)
	return
L29:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v98 = F_get_subscription_name(m, v96, int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v98
	F_errmsg(m, int32(_a_F_RemoveSubscriptionRel_1), v10+int32(32))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v106 = F_get_rel_name(m, l1)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v108 = int32(*(*int8)(unsafe.Add(mBase, uint32(v55)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v106
	v114 = F_errdetail(m, int32(_a_F_RemoveSubscriptionRel_2), v10+int32(16))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(_a_F_RemoveSubscriptionRel_3)
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_RemoveSubscriptionRel_4)
	F_errhint(m, int32(_a_F_RemoveSubscriptionRel_5), v10)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_RemoveSubscriptionRel_6), int32(575), int32(_a_F_RemoveSubscriptionRel_7))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
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
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v8 = Fn14293(m, l0, l1, int32(16), int32(_a_F_get_subscription_name_0), int32(4040), int32(_a_F_get_subscription_name_1), int32(67))
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
