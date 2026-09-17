package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InitializeLogRepWorker(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	v4 = m.G0
	v6 = v4 + int32(-64)
	m.G0 = v6
	F_SetConfigOption(m, int32(_a_F_InitializeLogRepWorker_0), int32(_a_F_InitializeLogRepWorker_1), int32(5), int32(10))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[0]))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
		F_BackgroundWorkerInitializeConnectionByOid(m, v16, v17, int32(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			F_SetConfigOption(m, int32(_a_F_InitializeLogRepWorker_2), int32(_a_F_InitializeLogRepWorker_3), int32(5), int32(10))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[1]))
				v34 = F_AllocSetContextCreateInternal(m, v29, int32(_a_F_InitializeLogRepWorker_4), int32(0), int32(_a_F_InitializeLogRepWorker_5), int32(_a_F_InitializeLogRepWorker_6))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[2])) = v34
					F_StartTransactionCommand(m)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						v39 = int32(_a_F_InitializeLogRepWorker_7)
						v40 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[3]))
						v43 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[2]))
						*(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[3])) = v43
						v47 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[0]))
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+32))
						F_LockSharedObject(m, int32(_a_F_InitializeLogRepWorker_8), v48, int32(1))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							v54 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[0]))
							v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+32))
							v57 = F_GetSubscription(m, v55, int32(1))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[4])) = v57
								if v57 == int32(0) {
									v64 = F_errstart(m, int32(15), int32(0))
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return
									} else {
										if v64 != 0 {
											v67 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[0]))
											v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+32))
											*(*int32)(unsafe.Add(mBase, uint32(v6))) = v68
											F_errmsg(m, int32(_a_F_InitializeLogRepWorker_9), v6)
											mBase = m.M
											v72 = m.ExcPending
											if v72 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_InitializeLogRepWorker_10), int32(_a_F_InitializeLogRepWorker_11), int32(_a_F_InitializeLogRepWorker_12))
												mBase = m.M
												v77 = m.ExcPending
												if v77 != 0 {
													return
												} else {
													v79 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[0]))
													v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
													if v80 == int32(2) {
														v83 = *(*int32)(unsafe.Add(mBase, uint32(v79)+32))
														F_ApplyLauncherForgetWorkerStartTime(m, v83)
														mBase = m.M
														v85 = m.ExcPending
														if v85 != 0 {
															return
														} else {
															F_proc_exit(m, int32(0))
															mBase = m.M
															v88 = m.ExcPending
															if v88 != 0 {
																return
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													} else {
														F_proc_exit(m, int32(0))
														mBase = m.M
														v88 = m.ExcPending
														if v88 != 0 {
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
											v79 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[0]))
											v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
											if v80 == int32(2) {
												v83 = *(*int32)(unsafe.Add(mBase, uint32(v79)+32))
												F_ApplyLauncherForgetWorkerStartTime(m, v83)
												mBase = m.M
												v85 = m.ExcPending
												if v85 != 0 {
													return
												} else {
													F_proc_exit(m, int32(0))
													mBase = m.M
													v88 = m.ExcPending
													if v88 != 0 {
														return
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											} else {
												F_proc_exit(m, int32(0))
												mBase = m.M
												v88 = m.ExcPending
												if v88 != 0 {
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
									*(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[3])) = v40
									v92 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[5])) = uint8(v92)
									v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+25)))
									if v96 != 0 {
										v128 = v57
										v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+44))
										F_SetConfigOption(m, int32(_a_F_InitializeLogRepWorker_13), v129, int32(4), int32(10))
										mBase = m.M
										v133 = m.ExcPending
										if v133 != 0 {
											return
										} else {
											F_CacheRegisterSyscacheCallback(m, int32(67), int32(1024), int32(0))
											mBase = m.M
											v138 = m.ExcPending
											if v138 != 0 {
												return
											} else {
												F_CacheRegisterSyscacheCallback(m, int32(11), int32(1024), int32(0))
												mBase = m.M
												v143 = m.ExcPending
												if v143 != 0 {
													return
												} else {
													v145 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[0]))
													v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145)+16)))
													if v146 != int32(1) {
														v184 = F_errstart(m, int32(15), int32(0))
														mBase = m.M
														v185 = m.ExcPending
														if v185 != 0 {
															return
														} else {
															if v184 == int32(0) {
																F_CommitTransactionCommand(m)
																mBase = m.M
																v205 = m.ExcPending
																if v205 != 0 {
																	return
																} else {
																	F_before_shmem_exit(m, int32(1025), int32(0))
																	mBase = m.M
																	v209 = m.ExcPending
																	if v209 != 0 {
																		return
																	} else {
																		m.G0 = v6 - int32(-64)
																		return
																	}
																}
															} else {
																v189 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[4]))
																v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+16))
																*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = v190
																F_errmsg(m, int32(_a_F_InitializeLogRepWorker_14), v4+int32(-32))
																mBase = m.M
																v196 = m.ExcPending
																if v196 != 0 {
																	return
																} else {
																	v199 = int32(_a_F_InitializeLogRepWorker_15)
																	F_errfinish(m, int32(_a_F_InitializeLogRepWorker_10), v199, int32(_a_F_InitializeLogRepWorker_12))
																	mBase = m.M
																	v202 = m.ExcPending
																	if v202 != 0 {
																		return
																	} else {
																		F_CommitTransactionCommand(m)
																		mBase = m.M
																		v205 = m.ExcPending
																		if v205 != 0 {
																			return
																		} else {
																			F_before_shmem_exit(m, int32(1025), int32(0))
																			mBase = m.M
																			v209 = m.ExcPending
																			if v209 != 0 {
																				return
																			} else {
																				m.G0 = v6 - int32(-64)
																				return
																			}
																		}
																	}
																}
															}
														}
													} else {
														v149 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
														if v149 != int32(1) {
															v184 = F_errstart(m, int32(15), int32(0))
															mBase = m.M
															v185 = m.ExcPending
															if v185 != 0 {
																return
															} else {
																if v184 == int32(0) {
																	F_CommitTransactionCommand(m)
																	mBase = m.M
																	v205 = m.ExcPending
																	if v205 != 0 {
																		return
																	} else {
																		F_before_shmem_exit(m, int32(1025), int32(0))
																		mBase = m.M
																		v209 = m.ExcPending
																		if v209 != 0 {
																			return
																		} else {
																			m.G0 = v6 - int32(-64)
																			return
																		}
																	}
																} else {
																	v189 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[4]))
																	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+16))
																	*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = v190
																	F_errmsg(m, int32(_a_F_InitializeLogRepWorker_14), v4+int32(-32))
																	mBase = m.M
																	v196 = m.ExcPending
																	if v196 != 0 {
																		return
																	} else {
																		v199 = int32(_a_F_InitializeLogRepWorker_15)
																		F_errfinish(m, int32(_a_F_InitializeLogRepWorker_10), v199, int32(_a_F_InitializeLogRepWorker_12))
																		mBase = m.M
																		v202 = m.ExcPending
																		if v202 != 0 {
																			return
																		} else {
																			F_CommitTransactionCommand(m)
																			mBase = m.M
																			v205 = m.ExcPending
																			if v205 != 0 {
																				return
																			} else {
																				F_before_shmem_exit(m, int32(1025), int32(0))
																				mBase = m.M
																				v209 = m.ExcPending
																				if v209 != 0 {
																					return
																				} else {
																					m.G0 = v6 - int32(-64)
																					return
																				}
																			}
																		}
																	}
																}
															}
														} else {
															v154 = F_errstart(m, int32(15), int32(0))
															mBase = m.M
															v155 = m.ExcPending
															if v155 != 0 {
																return
															} else {
																if v154 == int32(0) {
																	F_CommitTransactionCommand(m)
																	mBase = m.M
																	v205 = m.ExcPending
																	if v205 != 0 {
																		return
																	} else {
																		F_before_shmem_exit(m, int32(1025), int32(0))
																		mBase = m.M
																		v209 = m.ExcPending
																		if v209 != 0 {
																			return
																		} else {
																			m.G0 = v6 - int32(-64)
																			return
																		}
																	}
																} else {
																	v159 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[4]))
																	v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)+16))
																	v162 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[0]))
																	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)+36))
																	v164 = F_get_rel_name(m, v163)
																	mBase = m.M
																	v165 = m.ExcPending
																	if v165 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v164
																		*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v160
																		F_errmsg(m, int32(_a_F_InitializeLogRepWorker_16), v4+int32(-48))
																		mBase = m.M
																		v172 = m.ExcPending
																		if v172 != 0 {
																			return
																		} else {
																			v199 = int32(_a_F_InitializeLogRepWorker_17)
																			F_errfinish(m, int32(_a_F_InitializeLogRepWorker_10), v199, int32(_a_F_InitializeLogRepWorker_12))
																			mBase = m.M
																			v202 = m.ExcPending
																			if v202 != 0 {
																				return
																			} else {
																				F_CommitTransactionCommand(m)
																				mBase = m.M
																				v205 = m.ExcPending
																				if v205 != 0 {
																					return
																				} else {
																					F_before_shmem_exit(m, int32(1025), int32(0))
																					mBase = m.M
																					v209 = m.ExcPending
																					if v209 != 0 {
																						return
																					} else {
																						m.G0 = v6 - int32(-64)
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
									} else {
										v99 = F_errstart(m, int32(15), int32(0))
										mBase = m.M
										v100 = m.ExcPending
										if v100 != 0 {
											return
										} else {
											if v99 != 0 {
												v102 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[4]))
												v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+16))
												*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = v103
												F_errmsg(m, int32(_a_F_InitializeLogRepWorker_18), v4+int32(-16))
												mBase = m.M
												v109 = m.ExcPending
												if v109 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_InitializeLogRepWorker_10), int32(_a_F_InitializeLogRepWorker_19), int32(_a_F_InitializeLogRepWorker_12))
													mBase = m.M
													v114 = m.ExcPending
													if v114 != 0 {
														return
													} else {
														v116 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[0]))
														v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
														v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+16)))
														if base.B2i32(v118 != int32(1))|base.B2i32(v117 != int32(3)) != 0 {
															if v117 == int32(2) {
																v176 = *(*int32)(unsafe.Add(mBase, uint32(v116)+32))
																F_ApplyLauncherForgetWorkerStartTime(m, v176)
																mBase = m.M
																v178 = m.ExcPending
																if v178 != 0 {
																	return
																} else {
																	F_proc_exit(m, int32(0))
																	mBase = m.M
																	v181 = m.ExcPending
																	if v181 != 0 {
																		return
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															} else {
																F_proc_exit(m, int32(0))
																mBase = m.M
																v181 = m.ExcPending
																if v181 != 0 {
																	return
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														} else {
															v125 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[4]))
															v128 = v125
															v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+44))
															F_SetConfigOption(m, int32(_a_F_InitializeLogRepWorker_13), v129, int32(4), int32(10))
															mBase = m.M
															v133 = m.ExcPending
															if v133 != 0 {
																return
															} else {
																F_CacheRegisterSyscacheCallback(m, int32(67), int32(1024), int32(0))
																mBase = m.M
																v138 = m.ExcPending
																if v138 != 0 {
																	return
																} else {
																	F_CacheRegisterSyscacheCallback(m, int32(11), int32(1024), int32(0))
																	mBase = m.M
																	v143 = m.ExcPending
																	if v143 != 0 {
																		return
																	} else {
																		v145 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[0]))
																		v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145)+16)))
																		if v146 != int32(1) {
																			v184 = F_errstart(m, int32(15), int32(0))
																			mBase = m.M
																			v185 = m.ExcPending
																			if v185 != 0 {
																				return
																			} else {
																				if v184 == int32(0) {
																					F_CommitTransactionCommand(m)
																					mBase = m.M
																					v205 = m.ExcPending
																					if v205 != 0 {
																						return
																					} else {
																						F_before_shmem_exit(m, int32(1025), int32(0))
																						mBase = m.M
																						v209 = m.ExcPending
																						if v209 != 0 {
																							return
																						} else {
																							m.G0 = v6 - int32(-64)
																							return
																						}
																					}
																				} else {
																					v189 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[4]))
																					v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+16))
																					*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = v190
																					F_errmsg(m, int32(_a_F_InitializeLogRepWorker_14), v4+int32(-32))
																					mBase = m.M
																					v196 = m.ExcPending
																					if v196 != 0 {
																						return
																					} else {
																						v199 = int32(_a_F_InitializeLogRepWorker_15)
																						F_errfinish(m, int32(_a_F_InitializeLogRepWorker_10), v199, int32(_a_F_InitializeLogRepWorker_12))
																						mBase = m.M
																						v202 = m.ExcPending
																						if v202 != 0 {
																							return
																						} else {
																							F_CommitTransactionCommand(m)
																							mBase = m.M
																							v205 = m.ExcPending
																							if v205 != 0 {
																								return
																							} else {
																								F_before_shmem_exit(m, int32(1025), int32(0))
																								mBase = m.M
																								v209 = m.ExcPending
																								if v209 != 0 {
																									return
																								} else {
																									m.G0 = v6 - int32(-64)
																									return
																								}
																							}
																						}
																					}
																				}
																			}
																		} else {
																			v149 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
																			if v149 != int32(1) {
																				v184 = F_errstart(m, int32(15), int32(0))
																				mBase = m.M
																				v185 = m.ExcPending
																				if v185 != 0 {
																					return
																				} else {
																					if v184 == int32(0) {
																						F_CommitTransactionCommand(m)
																						mBase = m.M
																						v205 = m.ExcPending
																						if v205 != 0 {
																							return
																						} else {
																							F_before_shmem_exit(m, int32(1025), int32(0))
																							mBase = m.M
																							v209 = m.ExcPending
																							if v209 != 0 {
																								return
																							} else {
																								m.G0 = v6 - int32(-64)
																								return
																							}
																						}
																					} else {
																						v189 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[4]))
																						v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+16))
																						*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = v190
																						F_errmsg(m, int32(_a_F_InitializeLogRepWorker_14), v4+int32(-32))
																						mBase = m.M
																						v196 = m.ExcPending
																						if v196 != 0 {
																							return
																						} else {
																							v199 = int32(_a_F_InitializeLogRepWorker_15)
																							F_errfinish(m, int32(_a_F_InitializeLogRepWorker_10), v199, int32(_a_F_InitializeLogRepWorker_12))
																							mBase = m.M
																							v202 = m.ExcPending
																							if v202 != 0 {
																								return
																							} else {
																								F_CommitTransactionCommand(m)
																								mBase = m.M
																								v205 = m.ExcPending
																								if v205 != 0 {
																									return
																								} else {
																									F_before_shmem_exit(m, int32(1025), int32(0))
																									mBase = m.M
																									v209 = m.ExcPending
																									if v209 != 0 {
																										return
																									} else {
																										m.G0 = v6 - int32(-64)
																										return
																									}
																								}
																							}
																						}
																					}
																				}
																			} else {
																				v154 = F_errstart(m, int32(15), int32(0))
																				mBase = m.M
																				v155 = m.ExcPending
																				if v155 != 0 {
																					return
																				} else {
																					if v154 == int32(0) {
																						F_CommitTransactionCommand(m)
																						mBase = m.M
																						v205 = m.ExcPending
																						if v205 != 0 {
																							return
																						} else {
																							F_before_shmem_exit(m, int32(1025), int32(0))
																							mBase = m.M
																							v209 = m.ExcPending
																							if v209 != 0 {
																								return
																							} else {
																								m.G0 = v6 - int32(-64)
																								return
																							}
																						}
																					} else {
																						v159 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[4]))
																						v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)+16))
																						v162 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[0]))
																						v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)+36))
																						v164 = F_get_rel_name(m, v163)
																						mBase = m.M
																						v165 = m.ExcPending
																						if v165 != 0 {
																							return
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v164
																							*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v160
																							F_errmsg(m, int32(_a_F_InitializeLogRepWorker_16), v4+int32(-48))
																							mBase = m.M
																							v172 = m.ExcPending
																							if v172 != 0 {
																								return
																							} else {
																								v199 = int32(_a_F_InitializeLogRepWorker_17)
																								F_errfinish(m, int32(_a_F_InitializeLogRepWorker_10), v199, int32(_a_F_InitializeLogRepWorker_12))
																								mBase = m.M
																								v202 = m.ExcPending
																								if v202 != 0 {
																									return
																								} else {
																									F_CommitTransactionCommand(m)
																									mBase = m.M
																									v205 = m.ExcPending
																									if v205 != 0 {
																										return
																									} else {
																										F_before_shmem_exit(m, int32(1025), int32(0))
																										mBase = m.M
																										v209 = m.ExcPending
																										if v209 != 0 {
																											return
																										} else {
																											m.G0 = v6 - int32(-64)
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
											} else {
												v116 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[0]))
												v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
												v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+16)))
												if base.B2i32(v118 != int32(1))|base.B2i32(v117 != int32(3)) != 0 {
													if v117 == int32(2) {
														v176 = *(*int32)(unsafe.Add(mBase, uint32(v116)+32))
														F_ApplyLauncherForgetWorkerStartTime(m, v176)
														mBase = m.M
														v178 = m.ExcPending
														if v178 != 0 {
															return
														} else {
															F_proc_exit(m, int32(0))
															mBase = m.M
															v181 = m.ExcPending
															if v181 != 0 {
																return
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													} else {
														F_proc_exit(m, int32(0))
														mBase = m.M
														v181 = m.ExcPending
														if v181 != 0 {
															return
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												} else {
													v125 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[4]))
													v128 = v125
													v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+44))
													F_SetConfigOption(m, int32(_a_F_InitializeLogRepWorker_13), v129, int32(4), int32(10))
													mBase = m.M
													v133 = m.ExcPending
													if v133 != 0 {
														return
													} else {
														F_CacheRegisterSyscacheCallback(m, int32(67), int32(1024), int32(0))
														mBase = m.M
														v138 = m.ExcPending
														if v138 != 0 {
															return
														} else {
															F_CacheRegisterSyscacheCallback(m, int32(11), int32(1024), int32(0))
															mBase = m.M
															v143 = m.ExcPending
															if v143 != 0 {
																return
															} else {
																v145 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[0]))
																v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145)+16)))
																if v146 != int32(1) {
																	v184 = F_errstart(m, int32(15), int32(0))
																	mBase = m.M
																	v185 = m.ExcPending
																	if v185 != 0 {
																		return
																	} else {
																		if v184 == int32(0) {
																			F_CommitTransactionCommand(m)
																			mBase = m.M
																			v205 = m.ExcPending
																			if v205 != 0 {
																				return
																			} else {
																				F_before_shmem_exit(m, int32(1025), int32(0))
																				mBase = m.M
																				v209 = m.ExcPending
																				if v209 != 0 {
																					return
																				} else {
																					m.G0 = v6 - int32(-64)
																					return
																				}
																			}
																		} else {
																			v189 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[4]))
																			v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+16))
																			*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = v190
																			F_errmsg(m, int32(_a_F_InitializeLogRepWorker_14), v4+int32(-32))
																			mBase = m.M
																			v196 = m.ExcPending
																			if v196 != 0 {
																				return
																			} else {
																				v199 = int32(_a_F_InitializeLogRepWorker_15)
																				F_errfinish(m, int32(_a_F_InitializeLogRepWorker_10), v199, int32(_a_F_InitializeLogRepWorker_12))
																				mBase = m.M
																				v202 = m.ExcPending
																				if v202 != 0 {
																					return
																				} else {
																					F_CommitTransactionCommand(m)
																					mBase = m.M
																					v205 = m.ExcPending
																					if v205 != 0 {
																						return
																					} else {
																						F_before_shmem_exit(m, int32(1025), int32(0))
																						mBase = m.M
																						v209 = m.ExcPending
																						if v209 != 0 {
																							return
																						} else {
																							m.G0 = v6 - int32(-64)
																							return
																						}
																					}
																				}
																			}
																		}
																	}
																} else {
																	v149 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
																	if v149 != int32(1) {
																		v184 = F_errstart(m, int32(15), int32(0))
																		mBase = m.M
																		v185 = m.ExcPending
																		if v185 != 0 {
																			return
																		} else {
																			if v184 == int32(0) {
																				F_CommitTransactionCommand(m)
																				mBase = m.M
																				v205 = m.ExcPending
																				if v205 != 0 {
																					return
																				} else {
																					F_before_shmem_exit(m, int32(1025), int32(0))
																					mBase = m.M
																					v209 = m.ExcPending
																					if v209 != 0 {
																						return
																					} else {
																						m.G0 = v6 - int32(-64)
																						return
																					}
																				}
																			} else {
																				v189 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[4]))
																				v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+16))
																				*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = v190
																				F_errmsg(m, int32(_a_F_InitializeLogRepWorker_14), v4+int32(-32))
																				mBase = m.M
																				v196 = m.ExcPending
																				if v196 != 0 {
																					return
																				} else {
																					v199 = int32(_a_F_InitializeLogRepWorker_15)
																					F_errfinish(m, int32(_a_F_InitializeLogRepWorker_10), v199, int32(_a_F_InitializeLogRepWorker_12))
																					mBase = m.M
																					v202 = m.ExcPending
																					if v202 != 0 {
																						return
																					} else {
																						F_CommitTransactionCommand(m)
																						mBase = m.M
																						v205 = m.ExcPending
																						if v205 != 0 {
																							return
																						} else {
																							F_before_shmem_exit(m, int32(1025), int32(0))
																							mBase = m.M
																							v209 = m.ExcPending
																							if v209 != 0 {
																								return
																							} else {
																								m.G0 = v6 - int32(-64)
																								return
																							}
																						}
																					}
																				}
																			}
																		}
																	} else {
																		v154 = F_errstart(m, int32(15), int32(0))
																		mBase = m.M
																		v155 = m.ExcPending
																		if v155 != 0 {
																			return
																		} else {
																			if v154 == int32(0) {
																				F_CommitTransactionCommand(m)
																				mBase = m.M
																				v205 = m.ExcPending
																				if v205 != 0 {
																					return
																				} else {
																					F_before_shmem_exit(m, int32(1025), int32(0))
																					mBase = m.M
																					v209 = m.ExcPending
																					if v209 != 0 {
																						return
																					} else {
																						m.G0 = v6 - int32(-64)
																						return
																					}
																				}
																			} else {
																				v159 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[4]))
																				v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)+16))
																				v162 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLogRepWorker[0]))
																				v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)+36))
																				v164 = F_get_rel_name(m, v163)
																				mBase = m.M
																				v165 = m.ExcPending
																				if v165 != 0 {
																					return
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v164
																					*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v160
																					F_errmsg(m, int32(_a_F_InitializeLogRepWorker_16), v4+int32(-48))
																					mBase = m.M
																					v172 = m.ExcPending
																					if v172 != 0 {
																						return
																					} else {
																						v199 = int32(_a_F_InitializeLogRepWorker_17)
																						F_errfinish(m, int32(_a_F_InitializeLogRepWorker_10), v199, int32(_a_F_InitializeLogRepWorker_12))
																						mBase = m.M
																						v202 = m.ExcPending
																						if v202 != 0 {
																							return
																						} else {
																							F_CommitTransactionCommand(m)
																							mBase = m.M
																							v205 = m.ExcPending
																							if v205 != 0 {
																								return
																							} else {
																								F_before_shmem_exit(m, int32(1025), int32(0))
																								mBase = m.M
																								v209 = m.ExcPending
																								if v209 != 0 {
																									return
																								} else {
																									m.G0 = v6 - int32(-64)
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
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_LogRecoveryConflict(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v22 int64
	_ = v22
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
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
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	v12 = m.G0
	v14 = v12 - int32(80)
	m.G0 = v14
	v22 = l2 - l1
	if v22 <= int64(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
	v39 = int32(1000)
	v40 = base.I32_div_s(v38, v39)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = v38 - v40*v39
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
	v48 = int32(0)
	if l3 == v48 {
		v115 = v48
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v34 = int32(0)
	v35 = int32(0)
	goto L4
L3:
	;
	v26 = int64(1000000)
	v27 = base.I64_div_u_s(v22, v26)
	v34 = base.I32_wrap_i64(v27)
	v35 = base.I32_wrap_i64(v22 - v27*v26)
	goto L4
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14+int32(76)))) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v14+int32(72)))) = v35
	goto L1
L5:
	;
	v120 = v40 + v45*v39
	v123 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L23
	} else {
		goto L27
	}
L6:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v51 == int32(0) {
		v115 = v48
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v57 = l3
	v60 = v48
	goto L8
L8:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v66 = int32(0)
	if v65 < v66 {
		v84 = v66
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v115 = v104
	goto L5
L10:
	;
	if v84 != 0 {
		goto L17
	} else {
		goto L18
	}
L11:
	;
	goto L10
L12:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_LogRecoveryConflict[0]))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+16))
	if base.Ui32(v73) <= base.Ui32(v65) {
		v84 = int32(0)
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v78 = v75 + v65*int32(640)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v78)+44))
	if v80 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v81 = v78
	goto L16
L15:
	;
	v81 = int32(0)
	goto L16
L16:
	;
	v84 = v81
	goto L11
L17:
	;
	if v60 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v104 = v60
	goto L19
L19:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	if v106 != 0 {
		v57 = v57 + int32(8)
		v60 = v104
		goto L8
	} else {
		goto L26
	}
L20:
	;
	F_initStringInfo(m, v14+int32(56))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v93 = int32(_a_F_LogRecoveryConflict_0)
	goto L22
L22:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v84)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v94
	F_appendStringInfo(m, v14+int32(56), v93, v14+int32(48))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L23
	} else {
		goto L25
	}
L23:
	;
	return
L24:
	;
	v93 = int32(_a_F_LogRecoveryConflict_1)
	goto L22
L25:
	;
	v104 = v60 + int32(1)
	goto L19
L26:
	;
	goto L9
L27:
	;
	if l4 != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	if int32(0) < v115 {
		goto L46
	} else {
		goto L47
	}
L29:
	;
	F_errfinish(m, int32(_a_F_LogRecoveryConflict_2), v176, int32(_a_F_LogRecoveryConflict_3))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L23
	} else {
		goto L45
	}
L30:
	;
	if v123 == int32(0) {
		goto L28
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	if v123 == int32(0) {
		goto L28
	} else {
		goto L40
	}
L33:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
	v129 = l0 - int32(7)
	if base.Ui32(v129) <= base.Ui32(int32(6)) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v129<<(uint(int32(2))%32))+uint32(_c_F_LogRecoveryConflict[1])))
	v136 = v134
	goto L36
L35:
	;
	v136 = int32(_a_F_LogRecoveryConflict_4)
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v120
	F_errmsg(m, int32(_a_F_LogRecoveryConflict_5), v14+int32(16))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L23
	} else {
		goto L37
	}
L37:
	;
	v145 = int32(334)
	if v115 <= int32(0) {
		v176 = v145
		goto L29
	} else {
		goto L38
	}
L38:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v148
	F_errdetail_log_plural(m, int32(_a_F_LogRecoveryConflict_6), int32(_a_F_LogRecoveryConflict_7), v115, v14)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L23
	} else {
		goto L39
	}
L39:
	;
	v176 = v145
	goto L29
L40:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
	v158 = l0 - int32(7)
	if base.Ui32(v158) <= base.Ui32(int32(6)) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v158<<(uint(int32(2))%32))+uint32(_c_F_LogRecoveryConflict[1])))
	v165 = v163
	goto L43
L42:
	;
	v165 = int32(_a_F_LogRecoveryConflict_4)
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v165
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v120
	F_errmsg(m, int32(_a_F_LogRecoveryConflict_8), v14+int32(32))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L23
	} else {
		goto L44
	}
L44:
	;
	v176 = int32(340)
	goto L29
L45:
	;
	goto L28
L46:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
	F_pfree(m, v187)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L23
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	m.G0 = v14 + int32(80)
	return
L49:
	;
	goto L48
}
func F_LogStandbySnapshot(m *base.Module) int64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
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
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
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
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
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
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
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
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v372 int64
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v427 int64
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v449 int64
	_ = v449
	var v453 int32
	_ = v453
	var v457 int64
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v467 int64
	_ = v467
	var v473 int32
	_ = v473
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	v11 = m.G0
	v13 = v11 - int32(96)
	m.G0 = v13
	v17 = m.G0
	v19 = v17 - int32(32)
	m.G0 = v19
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	v26 = F_LWLockAcquire(m, v22+int32(_a_F_LogStandbySnapshot_0), int32(1))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	v35 = F_LWLockAcquire(m, v31+int32(_a_F_LogStandbySnapshot_1), int32(1))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	v42 = F_LWLockAcquire(m, v38+int32(_a_F_LogStandbySnapshot_2), int32(1))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	v49 = F_LWLockAcquire(m, v45+int32(_a_F_LogStandbySnapshot_3), int32(1))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	v56 = F_LWLockAcquire(m, v52+int32(_a_F_LogStandbySnapshot_4), int32(1))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	v63 = F_LWLockAcquire(m, v59+int32(_a_F_LogStandbySnapshot_5), int32(1))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	v70 = F_LWLockAcquire(m, v66+int32(_a_F_LogStandbySnapshot_6), int32(1))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	v77 = F_LWLockAcquire(m, v73+int32(_a_F_LogStandbySnapshot_7), int32(1))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	v84 = F_LWLockAcquire(m, v80+int32(_a_F_LogStandbySnapshot_8), int32(1))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	v91 = F_LWLockAcquire(m, v87+int32(_a_F_LogStandbySnapshot_9), int32(1))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	v98 = F_LWLockAcquire(m, v94+int32(_a_F_LogStandbySnapshot_10), int32(1))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	v105 = F_LWLockAcquire(m, v101+int32(_a_F_LogStandbySnapshot_11), int32(1))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	v112 = F_LWLockAcquire(m, v108+int32(_a_F_LogStandbySnapshot_12), int32(1))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	v119 = F_LWLockAcquire(m, v115+int32(_a_F_LogStandbySnapshot_13), int32(1))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	v126 = F_LWLockAcquire(m, v122+int32(_a_F_LogStandbySnapshot_14), int32(1))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v129 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	v133 = F_LWLockAcquire(m, v129+int32(_a_F_LogStandbySnapshot_15), int32(1))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v136 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[1]))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v138)+412))
	if v140 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v206 = F_palloc(m, v203*int32(12))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L22
	}
L19:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v138)+376))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v138)+364))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v138)+352))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v138)+340))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v138)+328))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v138)+316))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v138)+304))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v138)+292))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v138)+280))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v138)+268))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v138)+256))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v138)+244))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v138)+232))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v138)+220))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v138)+208))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v138)+196))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v138)+184))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v138)+172))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v138)+160))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v138)+148))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v138)+136))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v138)+124))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v138)+112))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v138)+100))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v138)+88))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v138)+76))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v138)+64))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v138)+52))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v138)+40))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v138)+28))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v138)+16))
	v203 = v141 + (v142 + (v143 + (v144 + (v145 + (v146 + (v147 + (v148 + (v149 + (v150 + (v151 + (v152 + (v153 + (v154 + (v155 + (v156 + (v157 + (v158 + (v159 + (v160 + (v161 + (v162 + (v163 + (v164 + (v165 + (v166 + (v167 + (v168 + (v169 + (v170 + (v171 + v139))))))))))))))))))))))))))))))
	goto L21
L20:
	;
	v203 = v139
	goto L21
L21:
	;
	goto L18
L22:
	;
	v211 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[1]))
	F_hash_seq_init(m, v19+int32(12), v211)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v218 = int32(0)
	goto L24
L24:
	;
	v226 = F_hash_seq_search(m, v19+int32(12))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	v250 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	F_LWLockRelease(m, v250+int32(_a_F_LogStandbySnapshot_15))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L33
	}
L26:
	;
	if v226 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+13)))
	if v228&int32(1) == int32(0) {
		goto L24
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	goto L25
L30:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v226)))
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+14)))
	if v234 != 0 {
		goto L24
	} else {
		goto L31
	}
L31:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v226)+4))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v235)+36))
	if v236 == int32(0) {
		goto L24
	} else {
		goto L32
	}
L32:
	;
	v241 = v206 + v218*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v241))) = v236
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
	*(*int32)(unsafe.Add(mBase, uint32(v241)+4)) = v243
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v233)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v241)+8)) = v245
	v218 = v218 + int32(1)
	goto L24
L33:
	;
	v256 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	F_LWLockRelease(m, v256+int32(_a_F_LogStandbySnapshot_14))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v262 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	F_LWLockRelease(m, v262+int32(_a_F_LogStandbySnapshot_13))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v268 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	F_LWLockRelease(m, v268+int32(_a_F_LogStandbySnapshot_12))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v274 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	F_LWLockRelease(m, v274+int32(_a_F_LogStandbySnapshot_11))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v280 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	F_LWLockRelease(m, v280+int32(_a_F_LogStandbySnapshot_10))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v286 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	F_LWLockRelease(m, v286+int32(_a_F_LogStandbySnapshot_9))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v292 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	F_LWLockRelease(m, v292+int32(_a_F_LogStandbySnapshot_8))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v298 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	F_LWLockRelease(m, v298+int32(_a_F_LogStandbySnapshot_7))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v304 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	F_LWLockRelease(m, v304+int32(_a_F_LogStandbySnapshot_6))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v310 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	F_LWLockRelease(m, v310+int32(_a_F_LogStandbySnapshot_5))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v316 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	F_LWLockRelease(m, v316+int32(_a_F_LogStandbySnapshot_4))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v322 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	F_LWLockRelease(m, v322+int32(_a_F_LogStandbySnapshot_3))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v328 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	F_LWLockRelease(m, v328+int32(_a_F_LogStandbySnapshot_2))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v334 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	F_LWLockRelease(m, v334+int32(_a_F_LogStandbySnapshot_1))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v340 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	F_LWLockRelease(m, v340+int32(_a_F_LogStandbySnapshot_0))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(68)))) = v218
	m.G0 = v19 + int32(32)
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
	if int32(0) < v349 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+72)) = v349
	F_XLogBeginInsert(m)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	F_pfree(m, v206)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L1
	} else {
		goto L57
	}
L52:
	;
	F_XLogRegisterData(m, v13+int32(72), int32(4))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_XLogRegisterData(m, v206, v349*int32(12))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v365 = int32(_a_F_LogStandbySnapshot_16)
	v367 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[2])))
	v368 = v367 | int32(2)
	*(*uint8)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[2])) = uint8(v368)
	goto L55
L55:
	;
	v372 = F_XLogInsert(m, int32(8), int32(0))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	goto L51
L57:
	;
	v376 = F_GetRunningTransactionData(m)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v379 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[3]))
	if v379 <= int32(1) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v383 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	F_LWLockRelease(m, v383+int32(512))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L1
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v376)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+72)) = v388
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v376)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+76)) = v390
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v376)+8))
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+80)) = uint8(base.B2i32(v392 != int32(0)))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v376)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v396
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v376)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v398
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v376)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+92)) = v400
	F_XLogBeginInsert(m)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L63
	}
L62:
	;
	goto L61
L63:
	;
	v405 = int32(_a_F_LogStandbySnapshot_16)
	v407 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[2])))
	v408 = v407 | int32(2)
	*(*uint8)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[2])) = uint8(v408)
	goto L64
L64:
	;
	F_XLogRegisterData(m, v13+int32(72), int32(24))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
	if int32(0) < v415 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v376)+28))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
	F_XLogRegisterData(m, v418, (v419+v415)<<(uint(int32(2))%32))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v427 = F_XLogInsert(m, int32(8), int32(16))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L1
	} else {
		goto L70
	}
L69:
	;
	goto L68
L70:
	;
	v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+80)))
	v432 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	if v429 == int32(1) {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	F_XLogSetAsyncXactLSN(m, v427)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L1
	} else {
		goto L82
	}
L73:
	;
	F_errfinish(m, int32(_a_F_LogStandbySnapshot_17), v479, int32(_a_F_LogStandbySnapshot_18))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L81
	}
L74:
	;
	if v432 == int32(0) {
		goto L72
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	if v432 == int32(0) {
		goto L72
	} else {
		goto L79
	}
L77:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v376)))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v376)+16))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v376)+24))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v376)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v439
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+8)) = uint32(v427)
	v449 = int64(base.Ui64(v427) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+4)) = uint32(v449)
	F_errmsg_internal(m, int32(_a_F_LogStandbySnapshot_19), v13)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v479 = int32(1384)
	goto L73
L79:
	;
	v457 = *(*int64)(unsafe.Add(mBase, uint32(v376)))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v376)+12))
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v376)+24))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v376)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v460
	*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = v458
	*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v457
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+44)) = uint32(v427)
	v467 = int64(base.Ui64(v427) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+40)) = uint32(v467)
	F_errmsg_internal(m, int32(_a_F_LogStandbySnapshot_20), v13+int32(32))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v479 = int32(1392)
	goto L73
L81:
	;
	goto L72
L82:
	;
	v490 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[3]))
	if int32(2) <= v490 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v494 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	F_LWLockRelease(m, v494+int32(512))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L1
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v500 = *(*int32)(unsafe.Add(mBase, _c_F_LogStandbySnapshot[0]))
	F_LWLockRelease(m, v500+int32(384))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L1
	} else {
		goto L87
	}
L86:
	;
	goto L85
L87:
	;
	m.G0 = v13 + int32(96)
	return v427
}
func F_check_log_stats(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	v4 = int32(1)
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v5 != v4 {
		v33 = v4
		return v33
	} else {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_log_stats[0])))
		if v9 != 0 {
			v21 = *(*int32)(unsafe.Add(mBase, _c_F_check_log_stats[1]))
			*(*int32)(unsafe.Add(mBase, _c_F_check_log_stats[2])) = v21
			v27 = F_format_elog_string(m, int32(_a_F_check_log_stats_0), int32(0))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_check_log_stats[3])) = v27
				v33 = int32(0)
				return v33
			}
		} else {
			v11 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_log_stats[4])))
			if v11&int32(1) != 0 {
				v21 = *(*int32)(unsafe.Add(mBase, _c_F_check_log_stats[1]))
				*(*int32)(unsafe.Add(mBase, _c_F_check_log_stats[2])) = v21
				v27 = F_format_elog_string(m, int32(_a_F_check_log_stats_0), int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_check_log_stats[3])) = v27
					v33 = int32(0)
					return v33
				}
			} else {
				v15 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_log_stats[5])))
				if v15&int32(1) == int32(0) {
					v33 = v4
					return v33
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, _c_F_check_log_stats[1]))
					*(*int32)(unsafe.Add(mBase, _c_F_check_log_stats[2])) = v21
					v27 = F_format_elog_string(m, int32(_a_F_check_log_stats_0), int32(0))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_check_log_stats[3])) = v27
						v33 = int32(0)
						return v33
					}
				}
			}
		}
	}
}
func F_log_status_format(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v257 int64
	_ = v257
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v276 int32
	_ = v276
	var v280 int64
	_ = v280
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v300 int32
	_ = v300
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v333 int32
	_ = v333
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v362 int32
	_ = v362
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v383 int32
	_ = v383
	var v386 int64
	_ = v386
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	var v429 int64
	_ = v429
	var v432 int32
	_ = v432
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v449 int32
	_ = v449
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v465 int32
	_ = v465
	var v468 int64
	_ = v468
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v489 int32
	_ = v489
	var v495 int32
	_ = v495
	var v499 int64
	_ = v499
	var v502 int32
	_ = v502
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v523 int32
	_ = v523
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v651 int32
	_ = v651
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v689 int32
	_ = v689
	var v693 int32
	_ = v693
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v707 int32
	_ = v707
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v730 int32
	_ = v730
	var v737 int32
	_ = v737
	var v741 int32
	_ = v741
	var v745 int32
	_ = v745
	var v752 int32
	_ = v752
	var v760 int32
	_ = v760
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v776 int32
	_ = v776
	var v784 int32
	_ = v784
	var v792 int32
	_ = v792
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v812 int32
	_ = v812
	var v817 int32
	_ = v817
	var v822 int32
	_ = v822
	var v826 int64
	_ = v826
	var v827 int64
	_ = v827
	var v834 int32
	_ = v834
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v847 int32
	_ = v847
	v10 = m.G0
	v12 = v10 - int32(688)
	m.G0 = v12
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[0]))
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[1]))
	if v16 == v18 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_log_status_format[2])) = v30
	if l1 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[2]))
	v30 = v21 + int32(1)
	goto L1
L3:
	;
	goto L4
L4:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_log_status_format[1])) = v16
	v27 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_log_status_format[3])) = uint8(v27)
	v30 = int32(1)
	goto L1
L5:
	;
	m.G0 = v12 + int32(688)
	return
L6:
	;
	v35 = l1
	goto L7
L7:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v43 != int32(37) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	if v43 == int32(0) {
		goto L5
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v54 = v35 + int32(1)
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+1)))
	if v55 != int32(37) {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	F_appendStringInfoChar(m, l0, base.I32_extend8_s(v43))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return
L14:
	;
	v35 = v35 + int32(1)
	goto L7
L15:
	;
	v35 = v122 + int32(1)
	goto L7
L16:
	;
	if v55 == int32(0) {
		goto L5
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	F_appendStringInfoChar(m, l0, int32(37))
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L13
	} else {
		goto L268
	}
L19:
	;
	v60 = int32(0)
	v61 = base.I32_extend8_s(v55)
	if int32(57) < v61 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	switch v127&int32(255) - int32(76) {
	case 0:
		goto L43
	default:
		goto L15
	case 4:
		goto L50
	case 5:
		goto L37
	case 21:
		goto L56
	case 22:
		goto L55
	case 23:
		goto L52
	case 24:
		goto L53
	case 25:
		goto L38
	case 28:
		goto L41
	case 29:
		goto L44
	case 32:
		goto L49
	case 33:
		goto L48
	case 34:
		goto L46
	case 36:
		goto L51
	case 37:
		goto L36
	case 38:
		goto L42
	case 39:
		goto L45
	case 40:
		goto L47
	case 41:
		goto L54
	case 42:
		goto L40
	case 44:
		goto L39
	}
L21:
	;
	v122 = v54
	v126 = v60
	v127 = v61
	goto L20
L22:
	;
	goto L23
L23:
	;
	if v61 != int32(45) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	if base.Ui32((v73-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	v73 = v61
	v74 = int32(1)
	v75 = v54
	goto L24
L26:
	;
	goto L27
L27:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+2)))
	if v67 == int32(0) {
		goto L5
	} else {
		goto L28
	}
L28:
	;
	v73 = v67
	v74 = int32(-1)
	v75 = v35 + int32(2)
	goto L24
L29:
	;
	v83 = v75
	v87 = v60
	v88 = v73
	goto L32
L30:
	;
	v109 = v75
	v113 = v60
	v114 = v73
	goto L31
L31:
	;
	if v114&int32(255) == int32(0) {
		goto L5
	} else {
		goto L35
	}
L32:
	;
	v91 = int32(10)
	v93 = int32(48)
	v95 = int32(255)
	v97 = v87*v91 + (v88-v93)&v95
	v99 = v83 + int32(1)
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+1)))
	if base.Ui32((v100-v93)&v95) < base.Ui32(v91) {
		v83 = v99
		v87 = v97
		v88 = v100
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v109 = v99
	v113 = v97 * v74
	v114 = v100
	goto L31
L34:
	;
	goto L33
L35:
	;
	v122 = v109
	v126 = v113
	v127 = v114
	goto L20
L36:
	;
	v844 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[4]))
	if v844 != 0 {
		goto L15
	} else {
		goto L267
	}
L37:
	;
	v822 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[5]))
	if v822 == int32(0) {
		goto L259
	} else {
		goto L260
	}
L38:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v765 = int32(63)
	v767 = int32(48)
	v768 = v764&v765 + v767
	*(*uint8)(unsafe.Add(mBase, _c_F_log_status_format[6])) = uint8(v768)
	v776 = int32(base.Ui32(v764)>>(uint(int32(24))%32))&v765 + v767
	*(*uint8)(unsafe.Add(mBase, _c_F_log_status_format[7])) = uint8(v776)
	v784 = int32(base.Ui32(v764)>>(uint(int32(18))%32))&v765 + v767
	*(*uint8)(unsafe.Add(mBase, _c_F_log_status_format[8])) = uint8(v784)
	v792 = int32(base.Ui32(v764)>>(uint(int32(12))%32))&v765 + v767
	*(*uint8)(unsafe.Add(mBase, _c_F_log_status_format[9])) = uint8(v792)
	v800 = int32(base.Ui32(v764)>>(uint(int32(6))%32))&v765 + v767
	*(*uint8)(unsafe.Add(mBase, _c_F_log_status_format[10])) = uint8(v800)
	v803 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_log_status_format[11])) = uint8(v803)
	if v126 != 0 {
		goto L253
	} else {
		goto L254
	}
L39:
	;
	v745 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[12]))
	if v126 != 0 {
		goto L248
	} else {
		goto L249
	}
L40:
	;
	v697 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[13]))
	if v697 == int32(0) {
		goto L237
	} else {
		goto L238
	}
L41:
	;
	v666 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[4]))
	if v666 == int32(0) {
		goto L227
	} else {
		goto L228
	}
L42:
	;
	v597 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[4]))
	if v597 == int32(0) {
		goto L208
	} else {
		goto L209
	}
L43:
	;
	v563 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[4]))
	if v563 != 0 {
		goto L196
	} else {
		goto L197
	}
L44:
	;
	v532 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[4]))
	if v532 != 0 {
		goto L185
	} else {
		goto L186
	}
L45:
	;
	v499 = *(*int64)(unsafe.Add(mBase, _c_F_log_status_format[14]))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+544)) = v499
	v502 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_log_status_format[3])))
	if v502 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L46:
	;
	v459 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_log_status_format[15])))
	if v459 == int32(0) {
		goto L166
	} else {
		goto L167
	}
L47:
	;
	v429 = F_time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v12)+680)) = v429
	v432 = v12 + int32(544)
	v438 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[16]))
	v439 = F_pg_localtime(m, v12+int32(680), v438)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L13
	} else {
		goto L159
	}
L48:
	;
	v374 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_log_status_format[17])) = uint8(v374)
	v377 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_log_status_format[15])))
	if v377 == v374 {
		goto L148
	} else {
		goto L149
	}
L49:
	;
	v355 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[2]))
	if v126 != 0 {
		goto L143
	} else {
		goto L144
	}
L50:
	;
	v312 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[13]))
	if v312 != 0 {
		goto L127
	} else {
		goto L128
	}
L51:
	;
	v293 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[0]))
	if v126 != 0 {
		goto L122
	} else {
		goto L123
	}
L52:
	;
	if v126 != 0 {
		goto L116
	} else {
		goto L117
	}
L53:
	;
	v228 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[4]))
	if v228 != 0 {
		goto L101
	} else {
		goto L102
	}
L54:
	;
	v199 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[4]))
	if v199 != 0 {
		goto L86
	} else {
		goto L87
	}
L55:
	;
	v164 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[0]))
	v166 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[18]))
	if v164 == v166 {
		v184 = int32(_a_F_log_status_format_0)
		goto L72
	} else {
		goto L73
	}
L56:
	;
	v135 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[4]))
	if v135 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v137 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[19]))
	if v137 != 0 {
		goto L61
	} else {
		goto L62
	}
L58:
	;
	goto L59
L59:
	;
	if v126 == int32(0) {
		goto L15
	} else {
		goto L70
	}
L60:
	;
	if v126 != 0 {
		goto L65
	} else {
		goto L66
	}
L61:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
	if v138 != 0 {
		v140 = v137
		goto L60
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v140 = int32(_a_F_log_status_format_1)
	goto L60
L64:
	;
	goto L63
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v126
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_2), v12)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L13
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	F_appendStringInfoString(m, l0, v140)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L13
	} else {
		goto L69
	}
L68:
	;
	v35 = v122 + int32(1)
	goto L7
L69:
	;
	v35 = v122 + int32(1)
	goto L7
L70:
	;
	v155 = v126 >> (uint(int32(31)) % 32)
	F_appendStringInfoSpaces(m, l0, v126^v155-v155)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L13
	} else {
		goto L71
	}
L71:
	;
	v35 = v122 + int32(1)
	goto L7
L72:
	;
	if v126 != 0 {
		goto L81
	} else {
		goto L82
	}
L73:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[20]))
	if v169 == int32(5) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[21]))
	v184 = v173 + int32(96)
	goto L72
L75:
	;
	goto L76
L76:
	;
	if base.Ui32(v169) <= base.Ui32(int32(17)) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v184 = v182
	goto L72
L78:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v169<<(uint(int32(2))%32))+uint32(_c_F_log_status_format[22])))
	v182 = v180
	goto L80
L79:
	;
	v182 = int32(_a_F_log_status_format_3)
	goto L80
L80:
	;
	goto L77
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v184
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v126
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_2), v12+int32(16))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L13
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	F_appendStringInfoString(m, l0, v184)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L13
	} else {
		goto L85
	}
L84:
	;
	v35 = v122 + int32(1)
	goto L7
L85:
	;
	v35 = v122 + int32(1)
	goto L7
L86:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v199)+364))
	if v200 != 0 {
		goto L90
	} else {
		goto L91
	}
L87:
	;
	goto L88
L88:
	;
	if v126 == int32(0) {
		goto L15
	} else {
		goto L99
	}
L89:
	;
	if v126 != 0 {
		goto L94
	} else {
		goto L95
	}
L90:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200))))
	if v201 != 0 {
		v203 = v200
		goto L89
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v203 = int32(_a_F_log_status_format_1)
	goto L89
L93:
	;
	goto L92
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v203
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v126
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_2), v12+int32(32))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L13
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	F_appendStringInfoString(m, l0, v203)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L13
	} else {
		goto L98
	}
L97:
	;
	v35 = v122 + int32(1)
	goto L7
L98:
	;
	v35 = v122 + int32(1)
	goto L7
L99:
	;
	v220 = v126 >> (uint(int32(31)) % 32)
	F_appendStringInfoSpaces(m, l0, v126^v220-v220)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L13
	} else {
		goto L100
	}
L100:
	;
	v35 = v122 + int32(1)
	goto L7
L101:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)+360))
	if v229 != 0 {
		goto L105
	} else {
		goto L106
	}
L102:
	;
	goto L103
L103:
	;
	if v126 == int32(0) {
		goto L15
	} else {
		goto L114
	}
L104:
	;
	if v126 != 0 {
		goto L109
	} else {
		goto L110
	}
L105:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229))))
	if v230 != 0 {
		v232 = v229
		goto L104
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v232 = int32(_a_F_log_status_format_1)
	goto L104
L108:
	;
	goto L107
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v126
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_2), v12+int32(48))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L13
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	F_appendStringInfoString(m, l0, v232)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L13
	} else {
		goto L113
	}
L112:
	;
	v35 = v122 + int32(1)
	goto L7
L113:
	;
	v35 = v122 + int32(1)
	goto L7
L114:
	;
	v249 = v126 >> (uint(int32(31)) % 32)
	F_appendStringInfoSpaces(m, l0, v126^v249-v249)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L13
	} else {
		goto L115
	}
L115:
	;
	v35 = v122 + int32(1)
	goto L7
L116:
	;
	v257 = *(*int64)(unsafe.Add(mBase, _c_F_log_status_format[14]))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+96)) = v257
	v260 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+104)) = v260
	v263 = v12 + int32(544)
	v268 = F_pg_snprintf(m, v263, int32(127), int32(_a_F_log_status_format_4), v12+int32(96))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L13
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	v280 = *(*int64)(unsafe.Add(mBase, _c_F_log_status_format[14]))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+64)) = v280
	v283 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = v283
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_4), v12-int32(-64))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L13
	} else {
		goto L121
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v12)+84)) = v263
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_2), v12+int32(80))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L13
	} else {
		goto L120
	}
L120:
	;
	v35 = v122 + int32(1)
	goto L7
L121:
	;
	v35 = v122 + int32(1)
	goto L7
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+132)) = v293
	*(*int32)(unsafe.Add(mBase, uint32(v12)+128)) = v126
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_5), v12+int32(128))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L13
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+112)) = v293
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_6), v12+int32(112))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L13
	} else {
		goto L126
	}
L125:
	;
	v35 = v122 + int32(1)
	goto L7
L126:
	;
	v35 = v122 + int32(1)
	goto L7
L127:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v312)+616))
	if v313 != 0 {
		goto L131
	} else {
		goto L132
	}
L128:
	;
	goto L129
L129:
	;
	if v126 == int32(0) {
		goto L15
	} else {
		goto L141
	}
L130:
	;
	if v126 != 0 {
		goto L136
	} else {
		goto L137
	}
L131:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v313)+44))
	v316 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[0]))
	if v314 != v316 {
		goto L130
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v320 = v126 >> (uint(int32(31)) % 32)
	F_appendStringInfoSpaces(m, l0, v126^v320-v320)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L13
	} else {
		goto L135
	}
L134:
	;
	goto L133
L135:
	;
	v35 = v122 + int32(1)
	goto L7
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+164)) = v314
	*(*int32)(unsafe.Add(mBase, uint32(v12)+160)) = v126
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_5), v12+int32(160))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L13
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+144)) = v314
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_6), v12+int32(144))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L13
	} else {
		goto L140
	}
L139:
	;
	v35 = v122 + int32(1)
	goto L7
L140:
	;
	v35 = v122 + int32(1)
	goto L7
L141:
	;
	v347 = v126 >> (uint(int32(31)) % 32)
	F_appendStringInfoSpaces(m, l0, v126^v347-v347)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L13
	} else {
		goto L142
	}
L142:
	;
	v35 = v122 + int32(1)
	goto L7
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+196)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v12)+192)) = v126
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_7), v12+int32(192))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L13
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+176)) = v355
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_8), v12+int32(176))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L13
	} else {
		goto L147
	}
L146:
	;
	v35 = v122 + int32(1)
	goto L7
L147:
	;
	v35 = v122 + int32(1)
	goto L7
L148:
	;
	F_gettimeofday(m, int32(_a_F_log_status_format_9))
	mBase = m.M
	v383 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_log_status_format[15])) = uint8(v383)
	goto L150
L149:
	;
	goto L150
L150:
	;
	v386 = *(*int64)(unsafe.Add(mBase, _c_F_log_status_format[23]))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+680)) = v386
	v394 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[16]))
	v395 = F_pg_localtime(m, v12+int32(680), v394)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L13
	} else {
		goto L151
	}
L151:
	;
	v397 = F_pg_strftime(m, int32(_a_F_log_status_format_10), int32(128), int32(_a_F_log_status_format_11), v395)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L13
	} else {
		goto L152
	}
L152:
	;
	v400 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[24]))
	v402 = base.I32_div_s(v400, int32(1000))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+224)) = v402
	v409 = F_pg_sprintf(m, v12+int32(544), int32(_a_F_log_status_format_12), v12+int32(224))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L13
	} else {
		goto L153
	}
L153:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v12)+544))
	*(*int32)(unsafe.Add(mBase, _c_F_log_status_format[25])) = v412
	if v126 != 0 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+212)) = int32(_a_F_log_status_format_10)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+208)) = v126
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_2), v12+int32(208))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L13
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_log_status_format_10))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L13
	} else {
		goto L158
	}
L157:
	;
	v35 = v122 + int32(1)
	goto L7
L158:
	;
	v35 = v122 + int32(1)
	goto L7
L159:
	;
	v441 = F_pg_strftime(m, v432, int32(128), int32(_a_F_log_status_format_13), v439)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L13
	} else {
		goto L160
	}
L160:
	;
	if v126 != 0 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+240)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v12)+244)) = v432
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_2), v12+int32(240))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L13
	} else {
		goto L164
	}
L162:
	;
	goto L163
L163:
	;
	F_appendStringInfoString(m, l0, v12+int32(544))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L13
	} else {
		goto L165
	}
L164:
	;
	v35 = v122 + int32(1)
	goto L7
L165:
	;
	v35 = v122 + int32(1)
	goto L7
L166:
	;
	F_gettimeofday(m, int32(_a_F_log_status_format_9))
	mBase = m.M
	v465 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_log_status_format[15])) = uint8(v465)
	goto L168
L167:
	;
	goto L168
L168:
	;
	v468 = *(*int64)(unsafe.Add(mBase, _c_F_log_status_format[23]))
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+272)) = uint32(v468)
	v471 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[24]))
	v473 = base.I32_div_s(v471, int32(1000))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+276)) = v473
	v476 = v12 + int32(544)
	v481 = F_pg_snprintf(m, v476, int32(128), int32(_a_F_log_status_format_14), v12+int32(272))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L13
	} else {
		goto L169
	}
L169:
	;
	if v126 != 0 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+256)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v12)+260)) = v476
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_2), v12+int32(256))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L13
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	F_appendStringInfoString(m, l0, v12+int32(544))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L13
	} else {
		goto L174
	}
L173:
	;
	v35 = v122 + int32(1)
	goto L7
L174:
	;
	v35 = v122 + int32(1)
	goto L7
L175:
	;
	v511 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[16]))
	v512 = F_pg_localtime(m, v12+int32(544), v511)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L13
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	if v126 != 0 {
		goto L180
	} else {
		goto L181
	}
L178:
	;
	v514 = F_pg_strftime(m, int32(_a_F_log_status_format_15), int32(128), int32(_a_F_log_status_format_13), v512)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L13
	} else {
		goto L179
	}
L179:
	;
	goto L177
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+292)) = int32(_a_F_log_status_format_15)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+288)) = v126
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_2), v12+int32(288))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L13
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_log_status_format_15))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L13
	} else {
		goto L184
	}
L183:
	;
	v35 = v122 + int32(1)
	goto L7
L184:
	;
	v35 = v122 + int32(1)
	goto L7
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12+int32(544)))) = int32(0)
	v537 = int32(_a_F_log_status_format_16)
	goto L188
L186:
	;
	goto L187
L187:
	;
	if v126 == int32(0) {
		goto L15
	} else {
		goto L194
	}
L188:
	;
	if v126 != 0 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+308)) = v537
	*(*int32)(unsafe.Add(mBase, uint32(v12)+304)) = v126
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_2), v12+int32(304))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L13
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v12)+544))
	F_appendBinaryStringInfo(m, l0, v537, v547)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L13
	} else {
		goto L193
	}
L192:
	;
	v35 = v122 + int32(1)
	goto L7
L193:
	;
	v35 = v122 + int32(1)
	goto L7
L194:
	;
	v555 = v126 >> (uint(int32(31)) % 32)
	F_appendStringInfoSpaces(m, l0, v126^v555-v555)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L13
	} else {
		goto L195
	}
L195:
	;
	v35 = v122 + int32(1)
	goto L7
L196:
	;
	v564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v563)+296)))
	if v564 != 0 {
		goto L199
	} else {
		goto L200
	}
L197:
	;
	v582 = int32(_a_F_log_status_format_17)
	goto L198
L198:
	;
	if v126 != 0 {
		goto L203
	} else {
		goto L204
	}
L199:
	;
	v578 = v563
	goto L201
L200:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v563)+140))
	v571 = int32(0)
	v574 = F_pg_getnameinfo_all(m, v563+int32(12), v567, v563+int32(296), int32(64), v571, v571, int32(3))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L13
	} else {
		goto L202
	}
L201:
	;
	v582 = v578 + int32(296)
	goto L198
L202:
	;
	v577 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[4]))
	v578 = v577
	goto L201
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+324)) = v582
	*(*int32)(unsafe.Add(mBase, uint32(v12)+320)) = v126
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_2), v12+int32(320))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L13
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	F_appendStringInfoString(m, l0, v582)
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L13
	} else {
		goto L207
	}
L206:
	;
	v35 = v122 + int32(1)
	goto L7
L207:
	;
	v35 = v122 + int32(1)
	goto L7
L208:
	;
	if v126 == int32(0) {
		goto L15
	} else {
		goto L225
	}
L209:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v597)+276))
	if v600 == int32(0) {
		goto L208
	} else {
		goto L210
	}
L210:
	;
	if v126 != 0 {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v597)+292))
	if v603 == int32(0) {
		goto L214
	} else {
		goto L215
	}
L212:
	;
	goto L213
L213:
	;
	F_appendStringInfoString(m, l0, v600)
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L13
	} else {
		goto L221
	}
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+356)) = v600
	*(*int32)(unsafe.Add(mBase, uint32(v12)+352)) = v126
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_2), v12+int32(352))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L13
	} else {
		goto L220
	}
L215:
	;
	v606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v603))))
	if v606 == int32(0) {
		goto L214
	} else {
		goto L216
	}
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+388)) = v603
	*(*int32)(unsafe.Add(mBase, uint32(v12)+384)) = v600
	v614 = F_psprintf(m, int32(_a_F_log_status_format_18), v12+int32(384))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L13
	} else {
		goto L217
	}
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+372)) = v614
	*(*int32)(unsafe.Add(mBase, uint32(v12)+368)) = v126
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_2), v12+int32(368))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L13
	} else {
		goto L218
	}
L218:
	;
	F_pfree(m, v614)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L13
	} else {
		goto L219
	}
L219:
	;
	v35 = v122 + int32(1)
	goto L7
L220:
	;
	v35 = v122 + int32(1)
	goto L7
L221:
	;
	v639 = *(*int32)(unsafe.Add(mBase, _c_F_log_status_format[4]))
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v639)+292))
	if v640 == int32(0) {
		goto L15
	} else {
		goto L222
	}
L222:
	;
	v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v640))))
	if v643 == int32(0) {
		goto L15
	} else {
		goto L223
	}
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+336)) = v640
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_19), v12+int32(336))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L13
	} else {
		goto L224
	}
L224:
	;
	v35 = v122 + int32(1)
	goto L7
L225:
	;
	v658 = v126 >> (uint(int32(31)) % 32)
	F_appendStringInfoSpaces(m, l0, v126^v658-v658)
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L13
	} else {
		goto L226
	}
L226:
	;
	v35 = v122 + int32(1)
	goto L7
L227:
	;
	if v126 == int32(0) {
		goto L15
	} else {
		goto L235
	}
L228:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v666)+276))
	if v669 == int32(0) {
		goto L227
	} else {
		goto L229
	}
L229:
	;
	if v126 != 0 {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+404)) = v669
	*(*int32)(unsafe.Add(mBase, uint32(v12)+400)) = v126
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_2), v12+int32(400))
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L13
	} else {
		goto L233
	}
L231:
	;
	goto L232
L232:
	;
	F_appendStringInfoString(m, l0, v669)
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L13
	} else {
		goto L234
	}
L233:
	;
	v35 = v122 + int32(1)
	goto L7
L234:
	;
	v35 = v122 + int32(1)
	goto L7
L235:
	;
	v689 = v126 >> (uint(int32(31)) % 32)
	F_appendStringInfoSpaces(m, l0, v126^v689-v689)
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L13
	} else {
		goto L236
	}
L236:
	;
	v35 = v122 + int32(1)
	goto L7
L237:
	;
	if v126 == int32(0) {
		goto L15
	} else {
		goto L246
	}
L238:
	;
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v697)+52))
	if v700 == int32(-1) {
		goto L237
	} else {
		goto L239
	}
L239:
	;
	if v126 != 0 {
		goto L240
	} else {
		goto L241
	}
L240:
	;
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v697)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+452)) = v703
	*(*int32)(unsafe.Add(mBase, uint32(v12)+448)) = v700
	v707 = v12 + int32(544)
	v712 = F_pg_snprintf(m, v707, int32(127), int32(_a_F_log_status_format_20), v12+int32(448))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L13
	} else {
		goto L243
	}
L241:
	;
	goto L242
L242:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v697)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+420)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v12)+416)) = v700
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_20), v12+int32(416))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L13
	} else {
		goto L245
	}
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+432)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v12)+436)) = v707
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_2), v12+int32(432))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L13
	} else {
		goto L244
	}
L244:
	;
	v35 = v122 + int32(1)
	goto L7
L245:
	;
	v35 = v122 + int32(1)
	goto L7
L246:
	;
	v737 = v126 >> (uint(int32(31)) % 32)
	F_appendStringInfoSpaces(m, l0, v126^v737-v737)
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L13
	} else {
		goto L247
	}
L247:
	;
	v35 = v122 + int32(1)
	goto L7
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+484)) = v745
	*(*int32)(unsafe.Add(mBase, uint32(v12)+480)) = v126
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_21), v12+int32(480))
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L13
	} else {
		goto L251
	}
L249:
	;
	goto L250
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+464)) = v745
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_22), v12+int32(464))
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L13
	} else {
		goto L252
	}
L251:
	;
	v35 = v122 + int32(1)
	goto L7
L252:
	;
	v35 = v122 + int32(1)
	goto L7
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+500)) = int32(_a_F_log_status_format_23)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+496)) = v126
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_2), v12+int32(496))
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L13
	} else {
		goto L256
	}
L254:
	;
	goto L255
L255:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_log_status_format_23))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L13
	} else {
		goto L257
	}
L256:
	;
	v35 = v122 + int32(1)
	goto L7
L257:
	;
	v35 = v122 + int32(1)
	goto L7
L258:
	;
	if v126 != 0 {
		goto L262
	} else {
		goto L263
	}
L259:
	;
	v827 = int64(0)
	goto L258
L260:
	;
	goto L261
L261:
	;
	v826 = *(*int64)(unsafe.Add(mBase, uint32(v822)+392))
	v827 = v826
	goto L258
L262:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+536)) = v827
	*(*int32)(unsafe.Add(mBase, uint32(v12)+528)) = v126
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_24), v12+int32(528))
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L13
	} else {
		goto L265
	}
L263:
	;
	goto L264
L264:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+512)) = v827
	F_appendStringInfo(m, l0, int32(_a_F_log_status_format_25), v12+int32(512))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L13
	} else {
		goto L266
	}
L265:
	;
	v35 = v122 + int32(1)
	goto L7
L266:
	;
	goto L15
L267:
	;
	goto L5
L268:
	;
	v35 = v35 + int32(2)
	goto L7
}
