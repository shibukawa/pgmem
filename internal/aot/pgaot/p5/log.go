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
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	v4 = m.G0
	v6 = v4 + int32(-64)
	m.G0 = v6
	F_SetConfigOption(m, int32(403390), int32(531833), int32(5), int32(10))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, _consts[507]))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
		F_BackgroundWorkerInitializeConnectionByOid(m, v16, v17, int32(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			F_SetConfigOption(m, int32(336644), int32(790160), int32(5), int32(10))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, _consts[147]))
				v34 = F_AllocSetContextCreateInternal(m, v29, int32(67790), int32(0), int32(8192), int32(8388608))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[531])) = v34
					F_StartTransactionCommand(m)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						v39 = int32(4554128)
						v40 = *(*int32)(unsafe.Add(mBase, _consts[10]))
						v43 = *(*int32)(unsafe.Add(mBase, _consts[531]))
						*(*int32)(unsafe.Add(mBase, _consts[10])) = v43
						v47 = *(*int32)(unsafe.Add(mBase, _consts[507]))
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+32))
						F_LockSharedObject(m, int32(6100), v48, int32(1))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							v54 = *(*int32)(unsafe.Add(mBase, _consts[507]))
							v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+32))
							v57 = F_GetSubscription(m, v55, int32(1))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[527])) = v57
								if v57 == int32(0) {
									v64 = F_errstart(m, int32(15), int32(0))
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return
									} else {
										if v64 != 0 {
											v67 = *(*int32)(unsafe.Add(mBase, _consts[507]))
											v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+32))
											*(*int32)(unsafe.Add(mBase, uint32(v6))) = v68
											F_errmsg(m, int32(243132), v6)
											mBase = m.M
											v72 = m.ExcPending
											if v72 != 0 {
												return
											} else {
												F_errfinish(m, int32(518459), int32(4697), int32(232011))
												mBase = m.M
												v77 = m.ExcPending
												if v77 != 0 {
													return
												} else {
													v79 = *(*int32)(unsafe.Add(mBase, _consts[507]))
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
											v79 = *(*int32)(unsafe.Add(mBase, _consts[507]))
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
									*(*int32)(unsafe.Add(mBase, _consts[10])) = v40
									v92 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _consts[556])) = uint8(v92)
									v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+25)))
									if v96 != 0 {
										v127 = v57
										v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+44))
										F_SetConfigOption(m, int32(107232), v128, int32(4), int32(10))
										mBase = m.M
										v132 = m.ExcPending
										if v132 != 0 {
											return
										} else {
											F_CacheRegisterSyscacheCallback(m, int32(67), int32(1024), int32(0))
											mBase = m.M
											v137 = m.ExcPending
											if v137 != 0 {
												return
											} else {
												F_CacheRegisterSyscacheCallback(m, int32(11), int32(1024), int32(0))
												mBase = m.M
												v142 = m.ExcPending
												if v142 != 0 {
													return
												} else {
													v144 = *(*int32)(unsafe.Add(mBase, _consts[507]))
													v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+16)))
													if v145 != int32(1) {
														v183 = F_errstart(m, int32(15), int32(0))
														mBase = m.M
														v184 = m.ExcPending
														if v184 != 0 {
															return
														} else {
															if v183 == int32(0) {
																F_CommitTransactionCommand(m)
																mBase = m.M
																v204 = m.ExcPending
																if v204 != 0 {
																	return
																} else {
																	F_before_shmem_exit(m, int32(1025), int32(0))
																	mBase = m.M
																	v208 = m.ExcPending
																	if v208 != 0 {
																		return
																	} else {
																		m.G0 = v6 - int32(-64)
																		return
																	}
																}
															} else {
																v188 = *(*int32)(unsafe.Add(mBase, _consts[527]))
																v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)+16))
																*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = v189
																F_errmsg(m, int32(464682), v4+int32(-32))
																mBase = m.M
																v195 = m.ExcPending
																if v195 != 0 {
																	return
																} else {
																	v198 = int32(4742)
																	F_errfinish(m, int32(518459), v198, int32(232011))
																	mBase = m.M
																	v201 = m.ExcPending
																	if v201 != 0 {
																		return
																	} else {
																		F_CommitTransactionCommand(m)
																		mBase = m.M
																		v204 = m.ExcPending
																		if v204 != 0 {
																			return
																		} else {
																			F_before_shmem_exit(m, int32(1025), int32(0))
																			mBase = m.M
																			v208 = m.ExcPending
																			if v208 != 0 {
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
														v148 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
														if v148 != int32(1) {
															v183 = F_errstart(m, int32(15), int32(0))
															mBase = m.M
															v184 = m.ExcPending
															if v184 != 0 {
																return
															} else {
																if v183 == int32(0) {
																	F_CommitTransactionCommand(m)
																	mBase = m.M
																	v204 = m.ExcPending
																	if v204 != 0 {
																		return
																	} else {
																		F_before_shmem_exit(m, int32(1025), int32(0))
																		mBase = m.M
																		v208 = m.ExcPending
																		if v208 != 0 {
																			return
																		} else {
																			m.G0 = v6 - int32(-64)
																			return
																		}
																	}
																} else {
																	v188 = *(*int32)(unsafe.Add(mBase, _consts[527]))
																	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)+16))
																	*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = v189
																	F_errmsg(m, int32(464682), v4+int32(-32))
																	mBase = m.M
																	v195 = m.ExcPending
																	if v195 != 0 {
																		return
																	} else {
																		v198 = int32(4742)
																		F_errfinish(m, int32(518459), v198, int32(232011))
																		mBase = m.M
																		v201 = m.ExcPending
																		if v201 != 0 {
																			return
																		} else {
																			F_CommitTransactionCommand(m)
																			mBase = m.M
																			v204 = m.ExcPending
																			if v204 != 0 {
																				return
																			} else {
																				F_before_shmem_exit(m, int32(1025), int32(0))
																				mBase = m.M
																				v208 = m.ExcPending
																				if v208 != 0 {
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
															v153 = F_errstart(m, int32(15), int32(0))
															mBase = m.M
															v154 = m.ExcPending
															if v154 != 0 {
																return
															} else {
																if v153 == int32(0) {
																	F_CommitTransactionCommand(m)
																	mBase = m.M
																	v204 = m.ExcPending
																	if v204 != 0 {
																		return
																	} else {
																		F_before_shmem_exit(m, int32(1025), int32(0))
																		mBase = m.M
																		v208 = m.ExcPending
																		if v208 != 0 {
																			return
																		} else {
																			m.G0 = v6 - int32(-64)
																			return
																		}
																	}
																} else {
																	v158 = *(*int32)(unsafe.Add(mBase, _consts[527]))
																	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)+16))
																	v161 = *(*int32)(unsafe.Add(mBase, _consts[507]))
																	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)+36))
																	v163 = F_get_rel_name(m, v162)
																	mBase = m.M
																	v164 = m.ExcPending
																	if v164 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v163
																		*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v159
																		F_errmsg(m, int32(464749), v4+int32(-48))
																		mBase = m.M
																		v171 = m.ExcPending
																		if v171 != 0 {
																			return
																		} else {
																			v198 = int32(4738)
																			F_errfinish(m, int32(518459), v198, int32(232011))
																			mBase = m.M
																			v201 = m.ExcPending
																			if v201 != 0 {
																				return
																			} else {
																				F_CommitTransactionCommand(m)
																				mBase = m.M
																				v204 = m.ExcPending
																				if v204 != 0 {
																					return
																				} else {
																					F_before_shmem_exit(m, int32(1025), int32(0))
																					mBase = m.M
																					v208 = m.ExcPending
																					if v208 != 0 {
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
												v102 = *(*int32)(unsafe.Add(mBase, _consts[527]))
												v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+16))
												*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = v103
												F_errmsg(m, int32(243307), v4+int32(-16))
												mBase = m.M
												v109 = m.ExcPending
												if v109 != 0 {
													return
												} else {
													F_errfinish(m, int32(518459), int32(4713), int32(232011))
													mBase = m.M
													v114 = m.ExcPending
													if v114 != 0 {
														return
													} else {
														v116 = *(*int32)(unsafe.Add(mBase, _consts[507]))
														v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
														v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+16)))
														if v118 != int32(1) {
															if v117 == int32(2) {
																v175 = *(*int32)(unsafe.Add(mBase, uint32(v116)+32))
																F_ApplyLauncherForgetWorkerStartTime(m, v175)
																mBase = m.M
																v177 = m.ExcPending
																if v177 != 0 {
																	return
																} else {
																	F_proc_exit(m, int32(0))
																	mBase = m.M
																	v180 = m.ExcPending
																	if v180 != 0 {
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
																v180 = m.ExcPending
																if v180 != 0 {
																	return
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														} else {
															if v117 != int32(3) {
																if v117 == int32(2) {
																	v175 = *(*int32)(unsafe.Add(mBase, uint32(v116)+32))
																	F_ApplyLauncherForgetWorkerStartTime(m, v175)
																	mBase = m.M
																	v177 = m.ExcPending
																	if v177 != 0 {
																		return
																	} else {
																		F_proc_exit(m, int32(0))
																		mBase = m.M
																		v180 = m.ExcPending
																		if v180 != 0 {
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
																	v180 = m.ExcPending
																	if v180 != 0 {
																		return
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															} else {
																v124 = *(*int32)(unsafe.Add(mBase, _consts[527]))
																v127 = v124
																v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+44))
																F_SetConfigOption(m, int32(107232), v128, int32(4), int32(10))
																mBase = m.M
																v132 = m.ExcPending
																if v132 != 0 {
																	return
																} else {
																	F_CacheRegisterSyscacheCallback(m, int32(67), int32(1024), int32(0))
																	mBase = m.M
																	v137 = m.ExcPending
																	if v137 != 0 {
																		return
																	} else {
																		F_CacheRegisterSyscacheCallback(m, int32(11), int32(1024), int32(0))
																		mBase = m.M
																		v142 = m.ExcPending
																		if v142 != 0 {
																			return
																		} else {
																			v144 = *(*int32)(unsafe.Add(mBase, _consts[507]))
																			v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+16)))
																			if v145 != int32(1) {
																				v183 = F_errstart(m, int32(15), int32(0))
																				mBase = m.M
																				v184 = m.ExcPending
																				if v184 != 0 {
																					return
																				} else {
																					if v183 == int32(0) {
																						F_CommitTransactionCommand(m)
																						mBase = m.M
																						v204 = m.ExcPending
																						if v204 != 0 {
																							return
																						} else {
																							F_before_shmem_exit(m, int32(1025), int32(0))
																							mBase = m.M
																							v208 = m.ExcPending
																							if v208 != 0 {
																								return
																							} else {
																								m.G0 = v6 - int32(-64)
																								return
																							}
																						}
																					} else {
																						v188 = *(*int32)(unsafe.Add(mBase, _consts[527]))
																						v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)+16))
																						*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = v189
																						F_errmsg(m, int32(464682), v4+int32(-32))
																						mBase = m.M
																						v195 = m.ExcPending
																						if v195 != 0 {
																							return
																						} else {
																							v198 = int32(4742)
																							F_errfinish(m, int32(518459), v198, int32(232011))
																							mBase = m.M
																							v201 = m.ExcPending
																							if v201 != 0 {
																								return
																							} else {
																								F_CommitTransactionCommand(m)
																								mBase = m.M
																								v204 = m.ExcPending
																								if v204 != 0 {
																									return
																								} else {
																									F_before_shmem_exit(m, int32(1025), int32(0))
																									mBase = m.M
																									v208 = m.ExcPending
																									if v208 != 0 {
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
																				v148 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
																				if v148 != int32(1) {
																					v183 = F_errstart(m, int32(15), int32(0))
																					mBase = m.M
																					v184 = m.ExcPending
																					if v184 != 0 {
																						return
																					} else {
																						if v183 == int32(0) {
																							F_CommitTransactionCommand(m)
																							mBase = m.M
																							v204 = m.ExcPending
																							if v204 != 0 {
																								return
																							} else {
																								F_before_shmem_exit(m, int32(1025), int32(0))
																								mBase = m.M
																								v208 = m.ExcPending
																								if v208 != 0 {
																									return
																								} else {
																									m.G0 = v6 - int32(-64)
																									return
																								}
																							}
																						} else {
																							v188 = *(*int32)(unsafe.Add(mBase, _consts[527]))
																							v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)+16))
																							*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = v189
																							F_errmsg(m, int32(464682), v4+int32(-32))
																							mBase = m.M
																							v195 = m.ExcPending
																							if v195 != 0 {
																								return
																							} else {
																								v198 = int32(4742)
																								F_errfinish(m, int32(518459), v198, int32(232011))
																								mBase = m.M
																								v201 = m.ExcPending
																								if v201 != 0 {
																									return
																								} else {
																									F_CommitTransactionCommand(m)
																									mBase = m.M
																									v204 = m.ExcPending
																									if v204 != 0 {
																										return
																									} else {
																										F_before_shmem_exit(m, int32(1025), int32(0))
																										mBase = m.M
																										v208 = m.ExcPending
																										if v208 != 0 {
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
																					v153 = F_errstart(m, int32(15), int32(0))
																					mBase = m.M
																					v154 = m.ExcPending
																					if v154 != 0 {
																						return
																					} else {
																						if v153 == int32(0) {
																							F_CommitTransactionCommand(m)
																							mBase = m.M
																							v204 = m.ExcPending
																							if v204 != 0 {
																								return
																							} else {
																								F_before_shmem_exit(m, int32(1025), int32(0))
																								mBase = m.M
																								v208 = m.ExcPending
																								if v208 != 0 {
																									return
																								} else {
																									m.G0 = v6 - int32(-64)
																									return
																								}
																							}
																						} else {
																							v158 = *(*int32)(unsafe.Add(mBase, _consts[527]))
																							v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)+16))
																							v161 = *(*int32)(unsafe.Add(mBase, _consts[507]))
																							v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)+36))
																							v163 = F_get_rel_name(m, v162)
																							mBase = m.M
																							v164 = m.ExcPending
																							if v164 != 0 {
																								return
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v163
																								*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v159
																								F_errmsg(m, int32(464749), v4+int32(-48))
																								mBase = m.M
																								v171 = m.ExcPending
																								if v171 != 0 {
																									return
																								} else {
																									v198 = int32(4738)
																									F_errfinish(m, int32(518459), v198, int32(232011))
																									mBase = m.M
																									v201 = m.ExcPending
																									if v201 != 0 {
																										return
																									} else {
																										F_CommitTransactionCommand(m)
																										mBase = m.M
																										v204 = m.ExcPending
																										if v204 != 0 {
																											return
																										} else {
																											F_before_shmem_exit(m, int32(1025), int32(0))
																											mBase = m.M
																											v208 = m.ExcPending
																											if v208 != 0 {
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
											} else {
												v116 = *(*int32)(unsafe.Add(mBase, _consts[507]))
												v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
												v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+16)))
												if v118 != int32(1) {
													if v117 == int32(2) {
														v175 = *(*int32)(unsafe.Add(mBase, uint32(v116)+32))
														F_ApplyLauncherForgetWorkerStartTime(m, v175)
														mBase = m.M
														v177 = m.ExcPending
														if v177 != 0 {
															return
														} else {
															F_proc_exit(m, int32(0))
															mBase = m.M
															v180 = m.ExcPending
															if v180 != 0 {
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
														v180 = m.ExcPending
														if v180 != 0 {
															return
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												} else {
													if v117 != int32(3) {
														if v117 == int32(2) {
															v175 = *(*int32)(unsafe.Add(mBase, uint32(v116)+32))
															F_ApplyLauncherForgetWorkerStartTime(m, v175)
															mBase = m.M
															v177 = m.ExcPending
															if v177 != 0 {
																return
															} else {
																F_proc_exit(m, int32(0))
																mBase = m.M
																v180 = m.ExcPending
																if v180 != 0 {
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
															v180 = m.ExcPending
															if v180 != 0 {
																return
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													} else {
														v124 = *(*int32)(unsafe.Add(mBase, _consts[527]))
														v127 = v124
														v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+44))
														F_SetConfigOption(m, int32(107232), v128, int32(4), int32(10))
														mBase = m.M
														v132 = m.ExcPending
														if v132 != 0 {
															return
														} else {
															F_CacheRegisterSyscacheCallback(m, int32(67), int32(1024), int32(0))
															mBase = m.M
															v137 = m.ExcPending
															if v137 != 0 {
																return
															} else {
																F_CacheRegisterSyscacheCallback(m, int32(11), int32(1024), int32(0))
																mBase = m.M
																v142 = m.ExcPending
																if v142 != 0 {
																	return
																} else {
																	v144 = *(*int32)(unsafe.Add(mBase, _consts[507]))
																	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+16)))
																	if v145 != int32(1) {
																		v183 = F_errstart(m, int32(15), int32(0))
																		mBase = m.M
																		v184 = m.ExcPending
																		if v184 != 0 {
																			return
																		} else {
																			if v183 == int32(0) {
																				F_CommitTransactionCommand(m)
																				mBase = m.M
																				v204 = m.ExcPending
																				if v204 != 0 {
																					return
																				} else {
																					F_before_shmem_exit(m, int32(1025), int32(0))
																					mBase = m.M
																					v208 = m.ExcPending
																					if v208 != 0 {
																						return
																					} else {
																						m.G0 = v6 - int32(-64)
																						return
																					}
																				}
																			} else {
																				v188 = *(*int32)(unsafe.Add(mBase, _consts[527]))
																				v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)+16))
																				*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = v189
																				F_errmsg(m, int32(464682), v4+int32(-32))
																				mBase = m.M
																				v195 = m.ExcPending
																				if v195 != 0 {
																					return
																				} else {
																					v198 = int32(4742)
																					F_errfinish(m, int32(518459), v198, int32(232011))
																					mBase = m.M
																					v201 = m.ExcPending
																					if v201 != 0 {
																						return
																					} else {
																						F_CommitTransactionCommand(m)
																						mBase = m.M
																						v204 = m.ExcPending
																						if v204 != 0 {
																							return
																						} else {
																							F_before_shmem_exit(m, int32(1025), int32(0))
																							mBase = m.M
																							v208 = m.ExcPending
																							if v208 != 0 {
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
																		v148 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
																		if v148 != int32(1) {
																			v183 = F_errstart(m, int32(15), int32(0))
																			mBase = m.M
																			v184 = m.ExcPending
																			if v184 != 0 {
																				return
																			} else {
																				if v183 == int32(0) {
																					F_CommitTransactionCommand(m)
																					mBase = m.M
																					v204 = m.ExcPending
																					if v204 != 0 {
																						return
																					} else {
																						F_before_shmem_exit(m, int32(1025), int32(0))
																						mBase = m.M
																						v208 = m.ExcPending
																						if v208 != 0 {
																							return
																						} else {
																							m.G0 = v6 - int32(-64)
																							return
																						}
																					}
																				} else {
																					v188 = *(*int32)(unsafe.Add(mBase, _consts[527]))
																					v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)+16))
																					*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = v189
																					F_errmsg(m, int32(464682), v4+int32(-32))
																					mBase = m.M
																					v195 = m.ExcPending
																					if v195 != 0 {
																						return
																					} else {
																						v198 = int32(4742)
																						F_errfinish(m, int32(518459), v198, int32(232011))
																						mBase = m.M
																						v201 = m.ExcPending
																						if v201 != 0 {
																							return
																						} else {
																							F_CommitTransactionCommand(m)
																							mBase = m.M
																							v204 = m.ExcPending
																							if v204 != 0 {
																								return
																							} else {
																								F_before_shmem_exit(m, int32(1025), int32(0))
																								mBase = m.M
																								v208 = m.ExcPending
																								if v208 != 0 {
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
																			v153 = F_errstart(m, int32(15), int32(0))
																			mBase = m.M
																			v154 = m.ExcPending
																			if v154 != 0 {
																				return
																			} else {
																				if v153 == int32(0) {
																					F_CommitTransactionCommand(m)
																					mBase = m.M
																					v204 = m.ExcPending
																					if v204 != 0 {
																						return
																					} else {
																						F_before_shmem_exit(m, int32(1025), int32(0))
																						mBase = m.M
																						v208 = m.ExcPending
																						if v208 != 0 {
																							return
																						} else {
																							m.G0 = v6 - int32(-64)
																							return
																						}
																					}
																				} else {
																					v158 = *(*int32)(unsafe.Add(mBase, _consts[527]))
																					v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)+16))
																					v161 = *(*int32)(unsafe.Add(mBase, _consts[507]))
																					v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)+36))
																					v163 = F_get_rel_name(m, v162)
																					mBase = m.M
																					v164 = m.ExcPending
																					if v164 != 0 {
																						return
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v163
																						*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v159
																						F_errmsg(m, int32(464749), v4+int32(-48))
																						mBase = m.M
																						v171 = m.ExcPending
																						if v171 != 0 {
																							return
																						} else {
																							v198 = int32(4738)
																							F_errfinish(m, int32(518459), v198, int32(232011))
																							mBase = m.M
																							v201 = m.ExcPending
																							if v201 != 0 {
																								return
																							} else {
																								F_CommitTransactionCommand(m)
																								mBase = m.M
																								v204 = m.ExcPending
																								if v204 != 0 {
																									return
																								} else {
																									F_before_shmem_exit(m, int32(1025), int32(0))
																									mBase = m.M
																									v208 = m.ExcPending
																									if v208 != 0 {
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
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	v12 = m.G0
	v14 = v12 - int32(80)
	m.G0 = v14
	v22 = l2 - l1
	if v22 <= int64(0) {
		goto L3
	} else {
		goto L4
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
		v117 = v48
		goto L6
	} else {
		goto L7
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14+int32(76)))) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v14+int32(72)))) = v35
	goto L1
L3:
	;
	v34 = int32(0)
	v35 = int32(0)
	goto L2
L4:
	;
	goto L5
L5:
	;
	v26 = int64(1000000)
	v27 = base.I64_div_u_s(v22, v26)
	v34 = base.I32_wrap_i64(v27)
	v35 = base.I32_wrap_i64(v22 - v27*v26)
	goto L2
L6:
	;
	v122 = v40 + v45*v39
	v125 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L24
	} else {
		goto L28
	}
L7:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v51 == int32(0) {
		v117 = v48
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v57 = l3
	v60 = v48
	goto L9
L9:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v66 = int32(0)
	if v65 < v66 {
		v84 = v66
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v117 = v104
	goto L6
L11:
	;
	if v84 != 0 {
		goto L18
	} else {
		goto L19
	}
L12:
	;
	goto L11
L13:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _consts[102]))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+16))
	if base.Ui32(v73) <= base.Ui32(v65) {
		v84 = int32(0)
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v78 = v75 + v65*int32(640)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v78)+44))
	if v80 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v81 = v78
	goto L17
L16:
	;
	v81 = int32(0)
	goto L17
L17:
	;
	v84 = v81
	goto L12
L18:
	;
	if v60 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v104 = v60
	goto L20
L20:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v57+int32(12))))
	if v110 != 0 {
		v57 = v57 + int32(8)
		v60 = v104
		goto L9
	} else {
		goto L27
	}
L21:
	;
	F_initStringInfo(m, v14+int32(56))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v93 = int32(509837)
	goto L23
L23:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v84)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v94
	F_appendStringInfo(m, v14+int32(56), v93, v14+int32(48))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L24
	} else {
		goto L26
	}
L24:
	;
	return
L25:
	;
	v93 = int32(509851)
	goto L23
L26:
	;
	v104 = v60 + int32(1)
	goto L20
L27:
	;
	goto L10
L28:
	;
	if l4 != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	if int32(0) < v117 {
		goto L47
	} else {
		goto L48
	}
L30:
	;
	F_errfinish(m, int32(514522), v181, int32(117069))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L24
	} else {
		goto L46
	}
L31:
	;
	if v125 == int32(0) {
		goto L29
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	if v125 == int32(0) {
		goto L29
	} else {
		goto L41
	}
L34:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
	v132 = l0 - int32(7)
	if base.Ui32(v132) <= base.Ui32(int32(6)) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v132<<(uint(int32(2))%32))+uint32(_consts[695])))
	v140 = v139
	goto L37
L36:
	;
	v140 = int32(257673)
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v122
	F_errmsg(m, int32(210610), v14+int32(16))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L24
	} else {
		goto L38
	}
L38:
	;
	v149 = int32(334)
	if v117 <= int32(0) {
		v181 = v149
		goto L30
	} else {
		goto L39
	}
L39:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v152
	F_errdetail_log_plural(m, int32(633088), int32(633113), v117, v14)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L24
	} else {
		goto L40
	}
L40:
	;
	v181 = v149
	goto L30
L41:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
	v163 = l0 - int32(7)
	if base.Ui32(v163) <= base.Ui32(int32(6)) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v163<<(uint(int32(2))%32))+uint32(_consts[695])))
	v171 = v170
	goto L44
L43:
	;
	v171 = int32(257673)
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v161
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v122
	F_errmsg(m, int32(210655), v14+int32(32))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L24
	} else {
		goto L45
	}
L45:
	;
	v181 = int32(340)
	goto L30
L46:
	;
	goto L29
L47:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
	F_pfree(m, v193)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L24
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	m.G0 = v14 + int32(80)
	return
L50:
	;
	goto L49
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
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v374 int64
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	var v429 int64
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v451 int64
	_ = v451
	var v455 int32
	_ = v455
	var v459 int64
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v469 int64
	_ = v469
	var v475 int32
	_ = v475
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	v11 = m.G0
	v13 = v11 - int32(96)
	m.G0 = v13
	v17 = m.G0
	v19 = v17 - int32(32)
	m.G0 = v19
	v22 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v26 = F_LWLockAcquire(m, v22+int32(23296), int32(1))
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
	v31 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v35 = F_LWLockAcquire(m, v31+int32(23424), int32(1))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v42 = F_LWLockAcquire(m, v38+int32(23552), int32(1))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v49 = F_LWLockAcquire(m, v45+int32(23680), int32(1))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v56 = F_LWLockAcquire(m, v52+int32(23808), int32(1))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v63 = F_LWLockAcquire(m, v59+int32(23936), int32(1))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v70 = F_LWLockAcquire(m, v66+int32(24064), int32(1))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v77 = F_LWLockAcquire(m, v73+int32(24192), int32(1))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v84 = F_LWLockAcquire(m, v80+int32(24320), int32(1))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v91 = F_LWLockAcquire(m, v87+int32(24448), int32(1))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v94 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v98 = F_LWLockAcquire(m, v94+int32(24576), int32(1))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v105 = F_LWLockAcquire(m, v101+int32(24704), int32(1))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v112 = F_LWLockAcquire(m, v108+int32(24832), int32(1))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v119 = F_LWLockAcquire(m, v115+int32(24960), int32(1))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v126 = F_LWLockAcquire(m, v122+int32(25088), int32(1))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v129 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v133 = F_LWLockAcquire(m, v129+int32(25216), int32(1))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v136 = *(*int32)(unsafe.Add(mBase, _consts[696]))
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
	v208 = F_palloc(m, v205*int32(12))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
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
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v138-int32(-64))))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v138)+52))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v138)+40))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v138)+28))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v138)+16))
	v205 = v141 + (v142 + (v143 + (v144 + (v145 + (v146 + (v147 + (v148 + (v149 + (v150 + (v151 + (v152 + (v153 + (v154 + (v155 + (v156 + (v157 + (v158 + (v159 + (v160 + (v161 + (v162 + (v163 + (v164 + (v165 + (v166 + (v169 + (v170 + (v171 + (v172 + (v173 + v139))))))))))))))))))))))))))))))
	goto L21
L20:
	;
	v205 = v139
	goto L21
L21:
	;
	goto L18
L22:
	;
	v213 = *(*int32)(unsafe.Add(mBase, _consts[696]))
	F_hash_seq_init(m, v19+int32(12), v213)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
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
	v228 = F_hash_seq_search(m, v19+int32(12))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	v252 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v252+int32(25216))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L33
	}
L26:
	;
	if v228 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+13)))
	if v230&int32(1) == int32(0) {
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
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235)+14)))
	if v236 != 0 {
		goto L24
	} else {
		goto L31
	}
L31:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v228)+4))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v237)+36))
	if v238 == int32(0) {
		goto L24
	} else {
		goto L32
	}
L32:
	;
	v243 = v208 + v218*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v243))) = v238
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v235)))
	*(*int32)(unsafe.Add(mBase, uint32(v243)+4)) = v245
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v235)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v243)+8)) = v247
	v218 = v218 + int32(1)
	goto L24
L33:
	;
	v258 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v258+int32(25088))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v264 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v264+int32(24960))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v270 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v270+int32(24832))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v276 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v276+int32(24704))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v282 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v282+int32(24576))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v288 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v288+int32(24448))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v294 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v294+int32(24320))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v300 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v300+int32(24192))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v306 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v306+int32(24064))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v312 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v312+int32(23936))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v318 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v318+int32(23808))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v324 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v324+int32(23680))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v330 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v330+int32(23552))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v336 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v336+int32(23424))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v342 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v342+int32(23296))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(68)))) = v218
	m.G0 = v19 + int32(32)
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
	if int32(0) < v351 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+72)) = v351
	F_XLogBeginInsert(m)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	F_pfree(m, v208)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L57
	}
L52:
	;
	F_XLogRegisterData(m, v13+int32(72), int32(4))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_XLogRegisterData(m, v208, v351*int32(12))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v367 = int32(4449972)
	v369 = int32(*(*uint8)(unsafe.Add(mBase, _consts[44])))
	v370 = v369 | int32(2)
	*(*uint8)(unsafe.Add(mBase, _consts[44])) = uint8(v370)
	goto L55
L55:
	;
	v374 = F_XLogInsert(m, int32(8), int32(0))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	goto L51
L57:
	;
	v378 = F_GetRunningTransactionData(m)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v381 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if v381 <= int32(1) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v385 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v385+int32(512))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v378)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+72)) = v390
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v378)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+76)) = v392
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v378)+8))
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+80)) = uint8(base.B2i32(v394 != int32(0)))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v378)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v398
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v378)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v400
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v378)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+92)) = v402
	F_XLogBeginInsert(m)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L63
	}
L62:
	;
	goto L61
L63:
	;
	v407 = int32(4449972)
	v409 = int32(*(*uint8)(unsafe.Add(mBase, _consts[44])))
	v410 = v409 | int32(2)
	*(*uint8)(unsafe.Add(mBase, _consts[44])) = uint8(v410)
	goto L64
L64:
	;
	F_XLogRegisterData(m, v13+int32(72), int32(24))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
	if int32(0) < v417 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v378)+28))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
	F_XLogRegisterData(m, v420, (v421+v417)<<(uint(int32(2))%32))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v429 = F_XLogInsert(m, int32(8), int32(16))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L70
	}
L69:
	;
	goto L68
L70:
	;
	v431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+80)))
	v434 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	if v431 == int32(1) {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	F_XLogSetAsyncXactLSN(m, v429)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L1
	} else {
		goto L82
	}
L73:
	;
	F_errfinish(m, int32(514522), v481, int32(133564))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L1
	} else {
		goto L81
	}
L74:
	;
	if v434 == int32(0) {
		goto L72
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	if v434 == int32(0) {
		goto L72
	} else {
		goto L79
	}
L77:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v378)))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v378)+16))
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v378)+24))
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v378)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v444
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v443
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v441
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+8)) = uint32(v429)
	v451 = int64(base.Ui64(v429) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+4)) = uint32(v451)
	F_errmsg_internal(m, int32(700889), v13)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v481 = int32(1384)
	goto L73
L79:
	;
	v459 = *(*int64)(unsafe.Add(mBase, uint32(v378)))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v378)+12))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v378)+24))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v378)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v461
	*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = v460
	*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v459
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+44)) = uint32(v429)
	v469 = int64(base.Ui64(v429) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+40)) = uint32(v469)
	F_errmsg_internal(m, int32(700790), v13+int32(32))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v481 = int32(1392)
	goto L73
L81:
	;
	goto L72
L82:
	;
	v492 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) <= v492 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v496 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v496+int32(512))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L1
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v502 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v502+int32(384))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
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
	return v429
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
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	v4 = int32(1)
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v5 != v4 {
		v30 = v4
		return v30
	} else {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, _consts[829])))
		if v9 != 0 {
			v18 = *(*int32)(unsafe.Add(mBase, _consts[87]))
			*(*int32)(unsafe.Add(mBase, _consts[88])) = v18
			v24 = F_format_elog_string(m, int32(655802), int32(0))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[89])) = v24
				v30 = int32(0)
				return v30
			}
		} else {
			v11 = int32(*(*uint8)(unsafe.Add(mBase, _consts[830])))
			if v11 != 0 {
				v18 = *(*int32)(unsafe.Add(mBase, _consts[87]))
				*(*int32)(unsafe.Add(mBase, _consts[88])) = v18
				v24 = F_format_elog_string(m, int32(655802), int32(0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[89])) = v24
					v30 = int32(0)
					return v30
				}
			} else {
				v12 = int32(1)
				v14 = int32(*(*uint8)(unsafe.Add(mBase, _consts[831])))
				if v14 != v12 {
					v30 = v12
					return v30
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, _consts[87]))
					*(*int32)(unsafe.Add(mBase, _consts[88])) = v18
					v24 = F_format_elog_string(m, int32(655802), int32(0))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[89])) = v24
						v30 = int32(0)
						return v30
					}
				}
			}
		}
	}
}
func F_log_status_format(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
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
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v256 int64
	_ = v256
	var v259 int32
	_ = v259
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v277 int32
	_ = v277
	var v281 int64
	_ = v281
	var v284 int32
	_ = v284
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v301 int32
	_ = v301
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v334 int32
	_ = v334
	var v342 int32
	_ = v342
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v363 int32
	_ = v363
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v384 int32
	_ = v384
	var v387 int64
	_ = v387
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v430 int64
	_ = v430
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v452 int32
	_ = v452
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v468 int32
	_ = v468
	var v471 int64
	_ = v471
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v494 int32
	_ = v494
	var v500 int32
	_ = v500
	var v504 int64
	_ = v504
	var v507 int32
	_ = v507
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v656 int32
	_ = v656
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v683 int32
	_ = v683
	var v687 int32
	_ = v687
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v737 int32
	_ = v737
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v752 int32
	_ = v752
	var v759 int32
	_ = v759
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v783 int32
	_ = v783
	var v791 int32
	_ = v791
	var v799 int32
	_ = v799
	var v807 int32
	_ = v807
	var v810 int32
	_ = v810
	var v819 int32
	_ = v819
	var v824 int32
	_ = v824
	var v829 int32
	_ = v829
	var v833 int64
	_ = v833
	var v834 int64
	_ = v834
	var v841 int32
	_ = v841
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	v9 = m.G0
	v11 = v9 - int32(688)
	m.G0 = v11
	v15 = *(*int32)(unsafe.Add(mBase, _consts[448]))
	v17 = *(*int32)(unsafe.Add(mBase, _consts[1195]))
	if v15 == v17 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1196])) = v29
	if l1 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _consts[1196]))
	v29 = v20 + int32(1)
	goto L1
L3:
	;
	goto L4
L4:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1195])) = v15
	v26 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[1183])) = uint8(v26)
	v29 = int32(1)
	goto L1
L5:
	;
	m.G0 = v11 + int32(688)
	return
L6:
	;
	v34 = l1
	goto L7
L7:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	if v41 != int32(37) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	if v41 == int32(0) {
		goto L5
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v52 = v34 + int32(1)
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	if v53 != int32(37) {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	F_appendStringInfoChar(m, l0, base.I32_extend8_s(v41))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return
L14:
	;
	v34 = v34 + int32(1)
	goto L7
L15:
	;
	v34 = v122 + int32(1)
	goto L7
L16:
	;
	if v53 == int32(0) {
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
	v854 = m.ExcPending
	if v854 != 0 {
		goto L13
	} else {
		goto L266
	}
L19:
	;
	v59 = base.I32_extend8_s(v53)
	if v59 <= int32(57) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if v59 != int32(45) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v122 = v52
	v123 = int32(0)
	v124 = v59
	goto L22
L22:
	;
	switch v124&int32(255) - int32(76) {
	case 0:
		goto L41
	default:
		goto L15
	case 4:
		goto L48
	case 5:
		goto L35
	case 21:
		goto L54
	case 22:
		goto L53
	case 23:
		goto L50
	case 24:
		goto L51
	case 25:
		goto L36
	case 28:
		goto L39
	case 29:
		goto L42
	case 32:
		goto L47
	case 33:
		goto L46
	case 34:
		goto L44
	case 36:
		goto L49
	case 37:
		goto L34
	case 38:
		goto L40
	case 39:
		goto L43
	case 40:
		goto L45
	case 41:
		goto L52
	case 42:
		goto L38
	case 44:
		goto L37
	}
L23:
	;
	v71 = v52
	v72 = v59
	v73 = int32(1)
	goto L25
L24:
	;
	v66 = v34 + int32(2)
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if v67 == int32(0) {
		goto L5
	} else {
		goto L26
	}
L25:
	;
	v74 = int32(0)
	if base.Ui32((v72-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v71 = v66
	v72 = v67
	v73 = int32(-1)
	goto L25
L27:
	;
	v82 = v74
	v85 = v71
	v87 = v72
	goto L30
L28:
	;
	v106 = v74
	v109 = v71
	v111 = v72
	goto L29
L29:
	;
	if v111&int32(255) == int32(0) {
		goto L5
	} else {
		goto L33
	}
L30:
	;
	v89 = int32(10)
	v91 = int32(48)
	v93 = int32(255)
	v95 = v82*v89 + (v87-v91)&v93
	v97 = v85 + int32(1)
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	if base.Ui32((v98-v91)&v93) < base.Ui32(v89) {
		v82 = v95
		v85 = v97
		v87 = v98
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v106 = v95
	v109 = v97
	v111 = v98
	goto L29
L32:
	;
	goto L31
L33:
	;
	v122 = v109
	v123 = v106 * v73
	v124 = v111
	goto L22
L34:
	;
	v851 = *(*int32)(unsafe.Add(mBase, _consts[378]))
	if v851 != 0 {
		goto L15
	} else {
		goto L265
	}
L35:
	;
	v829 = *(*int32)(unsafe.Add(mBase, _consts[49]))
	if v829 == int32(0) {
		goto L257
	} else {
		goto L258
	}
L36:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v772 = int32(63)
	v774 = int32(48)
	v775 = v771&v772 + v774
	*(*uint8)(unsafe.Add(mBase, _consts[1149])) = uint8(v775)
	v783 = int32(base.Ui32(v771)>>(uint(int32(24))%32))&v772 + v774
	*(*uint8)(unsafe.Add(mBase, _consts[1150])) = uint8(v783)
	v791 = int32(base.Ui32(v771)>>(uint(int32(18))%32))&v772 + v774
	*(*uint8)(unsafe.Add(mBase, _consts[1151])) = uint8(v791)
	v799 = int32(base.Ui32(v771)>>(uint(int32(12))%32))&v772 + v774
	*(*uint8)(unsafe.Add(mBase, _consts[1152])) = uint8(v799)
	v807 = int32(base.Ui32(v771)>>(uint(int32(6))%32))&v772 + v774
	*(*uint8)(unsafe.Add(mBase, _consts[1153])) = uint8(v807)
	v810 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[1154])) = uint8(v810)
	if v123 != 0 {
		goto L251
	} else {
		goto L252
	}
L37:
	;
	v752 = *(*int32)(unsafe.Add(mBase, _consts[38]))
	if v123 != 0 {
		goto L246
	} else {
		goto L247
	}
L38:
	;
	v702 = *(*int32)(unsafe.Add(mBase, _consts[293]))
	if v702 == int32(0) {
		goto L235
	} else {
		goto L236
	}
L39:
	;
	v671 = *(*int32)(unsafe.Add(mBase, _consts[378]))
	if v671 == int32(0) {
		goto L225
	} else {
		goto L226
	}
L40:
	;
	v602 = *(*int32)(unsafe.Add(mBase, _consts[378]))
	if v602 == int32(0) {
		goto L206
	} else {
		goto L207
	}
L41:
	;
	v568 = *(*int32)(unsafe.Add(mBase, _consts[378]))
	if v568 != 0 {
		goto L194
	} else {
		goto L195
	}
L42:
	;
	v537 = *(*int32)(unsafe.Add(mBase, _consts[378]))
	if v537 != 0 {
		goto L183
	} else {
		goto L184
	}
L43:
	;
	v504 = *(*int64)(unsafe.Add(mBase, _consts[447]))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+544)) = v504
	v507 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1183])))
	if v507 == int32(0) {
		goto L173
	} else {
		goto L174
	}
L44:
	;
	v462 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1140])))
	if v462 == int32(0) {
		goto L164
	} else {
		goto L165
	}
L45:
	;
	v430 = F___time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v11)+680)) = v430
	v439 = *(*int32)(unsafe.Add(mBase, _consts[145]))
	v440 = F_pg_localtime(m, v11+int32(680), v439)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L13
	} else {
		goto L157
	}
L46:
	;
	v375 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[1141])) = uint8(v375)
	v378 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1140])))
	if v378 == v375 {
		goto L146
	} else {
		goto L147
	}
L47:
	;
	v356 = *(*int32)(unsafe.Add(mBase, _consts[1196]))
	if v123 != 0 {
		goto L141
	} else {
		goto L142
	}
L48:
	;
	v313 = *(*int32)(unsafe.Add(mBase, _consts[293]))
	if v313 != 0 {
		goto L125
	} else {
		goto L126
	}
L49:
	;
	v294 = *(*int32)(unsafe.Add(mBase, _consts[448]))
	if v123 != 0 {
		goto L120
	} else {
		goto L121
	}
L50:
	;
	if v123 != 0 {
		goto L114
	} else {
		goto L115
	}
L51:
	;
	v227 = *(*int32)(unsafe.Add(mBase, _consts[378]))
	if v227 != 0 {
		goto L99
	} else {
		goto L100
	}
L52:
	;
	v198 = *(*int32)(unsafe.Add(mBase, _consts[378]))
	if v198 != 0 {
		goto L84
	} else {
		goto L85
	}
L53:
	;
	v160 = *(*int32)(unsafe.Add(mBase, _consts[448]))
	v162 = *(*int32)(unsafe.Add(mBase, _consts[1185]))
	if v160 == v162 {
		v183 = int32(225967)
		goto L70
	} else {
		goto L71
	}
L54:
	;
	v131 = *(*int32)(unsafe.Add(mBase, _consts[378]))
	if v131 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v133 = *(*int32)(unsafe.Add(mBase, _consts[1184]))
	if v133 != 0 {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	goto L57
L57:
	;
	if v123 == int32(0) {
		goto L15
	} else {
		goto L68
	}
L58:
	;
	if v123 != 0 {
		goto L63
	} else {
		goto L64
	}
L59:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	if v134 != 0 {
		v136 = v133
		goto L58
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v136 = int32(532182)
	goto L58
L62:
	;
	goto L61
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v123
	F_appendStringInfo(m, l0, int32(185605), v11)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L13
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	F_appendStringInfoString(m, l0, v136)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L13
	} else {
		goto L67
	}
L66:
	;
	v34 = v122 + int32(1)
	goto L7
L67:
	;
	v34 = v122 + int32(1)
	goto L7
L68:
	;
	v151 = v123 >> (uint(int32(31)) % 32)
	F_appendStringInfoSpaces(m, l0, v123^v151-v151)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L13
	} else {
		goto L69
	}
L69:
	;
	v34 = v122 + int32(1)
	goto L7
L70:
	;
	if v123 != 0 {
		goto L79
	} else {
		goto L80
	}
L71:
	;
	v165 = *(*int32)(unsafe.Add(mBase, _consts[20]))
	if v165 == int32(5) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _consts[804]))
	v183 = v169 + int32(96)
	goto L70
L73:
	;
	goto L74
L74:
	;
	if base.Ui32(v165) <= base.Ui32(int32(17)) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v183 = v181
	goto L70
L76:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v165<<(uint(int32(2))%32))+uint32(_consts[468])))
	v181 = v180
	goto L78
L77:
	;
	v181 = int32(385464)
	goto L78
L78:
	;
	goto L75
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v123
	F_appendStringInfo(m, l0, int32(185605), v11+int32(16))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L13
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	F_appendStringInfoString(m, l0, v183)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L13
	} else {
		goto L83
	}
L82:
	;
	v34 = v122 + int32(1)
	goto L7
L83:
	;
	v34 = v122 + int32(1)
	goto L7
L84:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)+364))
	if v199 != 0 {
		goto L88
	} else {
		goto L89
	}
L85:
	;
	goto L86
L86:
	;
	if v123 == int32(0) {
		goto L15
	} else {
		goto L97
	}
L87:
	;
	if v123 != 0 {
		goto L92
	} else {
		goto L93
	}
L88:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199))))
	if v200 != 0 {
		v202 = v199
		goto L87
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v202 = int32(532182)
	goto L87
L91:
	;
	goto L90
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v202
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v123
	F_appendStringInfo(m, l0, int32(185605), v11+int32(32))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L13
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	F_appendStringInfoString(m, l0, v202)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L13
	} else {
		goto L96
	}
L95:
	;
	v34 = v122 + int32(1)
	goto L7
L96:
	;
	v34 = v122 + int32(1)
	goto L7
L97:
	;
	v219 = v123 >> (uint(int32(31)) % 32)
	F_appendStringInfoSpaces(m, l0, v123^v219-v219)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L13
	} else {
		goto L98
	}
L98:
	;
	v34 = v122 + int32(1)
	goto L7
L99:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)+360))
	if v228 != 0 {
		goto L103
	} else {
		goto L104
	}
L100:
	;
	goto L101
L101:
	;
	if v123 == int32(0) {
		goto L15
	} else {
		goto L112
	}
L102:
	;
	if v123 != 0 {
		goto L107
	} else {
		goto L108
	}
L103:
	;
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228))))
	if v229 != 0 {
		v231 = v228
		goto L102
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v231 = int32(532182)
	goto L102
L106:
	;
	goto L105
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v123
	F_appendStringInfo(m, l0, int32(185605), v11+int32(48))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L13
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	F_appendStringInfoString(m, l0, v231)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L13
	} else {
		goto L111
	}
L110:
	;
	v34 = v122 + int32(1)
	goto L7
L111:
	;
	v34 = v122 + int32(1)
	goto L7
L112:
	;
	v248 = v123 >> (uint(int32(31)) % 32)
	F_appendStringInfoSpaces(m, l0, v123^v248-v248)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L13
	} else {
		goto L113
	}
L113:
	;
	v34 = v122 + int32(1)
	goto L7
L114:
	;
	v256 = *(*int64)(unsafe.Add(mBase, _consts[447]))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+96)) = v256
	v259 = *(*int32)(unsafe.Add(mBase, _consts[448]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+104)) = v259
	v267 = F_pg_snprintf(m, v11+int32(544), int32(127), int32(31143), v11+int32(96))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L13
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	v281 = *(*int64)(unsafe.Add(mBase, _consts[447]))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+64)) = v281
	v284 = *(*int32)(unsafe.Add(mBase, _consts[448]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v284
	F_appendStringInfo(m, l0, int32(31143), v11-int32(-64))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L13
	} else {
		goto L119
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v11)+84)) = v11 + int32(544)
	F_appendStringInfo(m, l0, int32(185605), v11+int32(80))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L13
	} else {
		goto L118
	}
L118:
	;
	v34 = v122 + int32(1)
	goto L7
L119:
	;
	v34 = v122 + int32(1)
	goto L7
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+132)) = v294
	*(*int32)(unsafe.Add(mBase, uint32(v11)+128)) = v123
	F_appendStringInfo(m, l0, int32(486670), v11+int32(128))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L13
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+112)) = v294
	F_appendStringInfo(m, l0, int32(509851), v11+int32(112))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L13
	} else {
		goto L124
	}
L123:
	;
	v34 = v122 + int32(1)
	goto L7
L124:
	;
	v34 = v122 + int32(1)
	goto L7
L125:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v313)+616))
	if v314 != 0 {
		goto L129
	} else {
		goto L130
	}
L126:
	;
	goto L127
L127:
	;
	if v123 == int32(0) {
		goto L15
	} else {
		goto L139
	}
L128:
	;
	if v123 != 0 {
		goto L134
	} else {
		goto L135
	}
L129:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v314)+44))
	v317 = *(*int32)(unsafe.Add(mBase, _consts[448]))
	if v315 != v317 {
		goto L128
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v321 = v123 >> (uint(int32(31)) % 32)
	F_appendStringInfoSpaces(m, l0, v123^v321-v321)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L13
	} else {
		goto L133
	}
L132:
	;
	goto L131
L133:
	;
	v34 = v122 + int32(1)
	goto L7
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+164)) = v315
	*(*int32)(unsafe.Add(mBase, uint32(v11)+160)) = v123
	F_appendStringInfo(m, l0, int32(486670), v11+int32(160))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L13
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+144)) = v315
	F_appendStringInfo(m, l0, int32(509851), v11+int32(144))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L13
	} else {
		goto L138
	}
L137:
	;
	v34 = v122 + int32(1)
	goto L7
L138:
	;
	v34 = v122 + int32(1)
	goto L7
L139:
	;
	v348 = v123 >> (uint(int32(31)) % 32)
	F_appendStringInfoSpaces(m, l0, v123^v348-v348)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L13
	} else {
		goto L140
	}
L140:
	;
	v34 = v122 + int32(1)
	goto L7
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+196)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v11)+192)) = v123
	F_appendStringInfo(m, l0, int32(451863), v11+int32(192))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L13
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+176)) = v356
	F_appendStringInfo(m, l0, int32(452018), v11+int32(176))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L13
	} else {
		goto L145
	}
L144:
	;
	v34 = v122 + int32(1)
	goto L7
L145:
	;
	v34 = v122 + int32(1)
	goto L7
L146:
	;
	F___gettimeofday(m, int32(4547712))
	mBase = m.M
	v384 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[1140])) = uint8(v384)
	goto L148
L147:
	;
	goto L148
L148:
	;
	v387 = *(*int64)(unsafe.Add(mBase, _consts[1197]))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+680)) = v387
	v395 = *(*int32)(unsafe.Add(mBase, _consts[145]))
	v396 = F_pg_localtime(m, v11+int32(680), v395)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L13
	} else {
		goto L149
	}
L149:
	;
	v398 = F_pg_strftime(m, int32(4547568), int32(128), int32(532800), v396)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L13
	} else {
		goto L150
	}
L150:
	;
	v401 = *(*int32)(unsafe.Add(mBase, _consts[1198]))
	v403 = base.I32_div_s(v401, int32(1000))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+224)) = v403
	v410 = F_pg_sprintf(m, v11+int32(544), int32(486535), v11+int32(224))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L13
	} else {
		goto L151
	}
L151:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v11)+544))
	*(*int32)(unsafe.Add(mBase, _consts[1199])) = v413
	if v123 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+212)) = int32(4547568)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+208)) = v123
	F_appendStringInfo(m, l0, int32(185605), v11+int32(208))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L13
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	F_appendStringInfoString(m, l0, int32(4547568))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L13
	} else {
		goto L156
	}
L155:
	;
	v34 = v122 + int32(1)
	goto L7
L156:
	;
	v34 = v122 + int32(1)
	goto L7
L157:
	;
	v442 = F_pg_strftime(m, v11+int32(544), int32(128), int32(532779), v440)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L13
	} else {
		goto L158
	}
L158:
	;
	if v123 != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+240)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v11)+244)) = v11 + int32(544)
	F_appendStringInfo(m, l0, int32(185605), v11+int32(240))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L13
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	F_appendStringInfoString(m, l0, v11+int32(544))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L13
	} else {
		goto L163
	}
L162:
	;
	v34 = v122 + int32(1)
	goto L7
L163:
	;
	v34 = v122 + int32(1)
	goto L7
L164:
	;
	F___gettimeofday(m, int32(4547712))
	mBase = m.M
	v468 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[1140])) = uint8(v468)
	goto L166
L165:
	;
	goto L166
L166:
	;
	v471 = *(*int64)(unsafe.Add(mBase, _consts[1197]))
	*(*uint32)(unsafe.Add(mBase, uint32(v11)+272)) = uint32(v471)
	v474 = *(*int32)(unsafe.Add(mBase, _consts[1198]))
	v476 = base.I32_div_s(v474, int32(1000))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+276)) = v476
	v484 = F_pg_snprintf(m, v11+int32(544), int32(128), int32(486532), v11+int32(272))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L13
	} else {
		goto L167
	}
L167:
	;
	if v123 != 0 {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+256)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v11)+260)) = v11 + int32(544)
	F_appendStringInfo(m, l0, int32(185605), v11+int32(256))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L13
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	F_appendStringInfoString(m, l0, v11+int32(544))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L13
	} else {
		goto L172
	}
L171:
	;
	v34 = v122 + int32(1)
	goto L7
L172:
	;
	v34 = v122 + int32(1)
	goto L7
L173:
	;
	v516 = *(*int32)(unsafe.Add(mBase, _consts[145]))
	v517 = F_pg_localtime(m, v11+int32(544), v516)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L13
	} else {
		goto L176
	}
L174:
	;
	goto L175
L175:
	;
	if v123 != 0 {
		goto L178
	} else {
		goto L179
	}
L176:
	;
	v519 = F_pg_strftime(m, int32(4547728), int32(128), int32(532779), v517)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L13
	} else {
		goto L177
	}
L177:
	;
	goto L175
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+292)) = int32(4547728)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+288)) = v123
	F_appendStringInfo(m, l0, int32(185605), v11+int32(288))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L13
	} else {
		goto L181
	}
L179:
	;
	goto L180
L180:
	;
	F_appendStringInfoString(m, l0, int32(4547728))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L13
	} else {
		goto L182
	}
L181:
	;
	v34 = v122 + int32(1)
	goto L7
L182:
	;
	v34 = v122 + int32(1)
	goto L7
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11+int32(544)))) = int32(0)
	v542 = int32(790160)
	goto L186
L184:
	;
	goto L185
L185:
	;
	if v123 == int32(0) {
		goto L15
	} else {
		goto L192
	}
L186:
	;
	if v123 != 0 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+308)) = v542
	*(*int32)(unsafe.Add(mBase, uint32(v11)+304)) = v123
	F_appendStringInfo(m, l0, int32(185605), v11+int32(304))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L13
	} else {
		goto L190
	}
L188:
	;
	goto L189
L189:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v11)+544))
	F_appendBinaryStringInfo(m, l0, v542, v552)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L13
	} else {
		goto L191
	}
L190:
	;
	v34 = v122 + int32(1)
	goto L7
L191:
	;
	v34 = v122 + int32(1)
	goto L7
L192:
	;
	v560 = v123 >> (uint(int32(31)) % 32)
	F_appendStringInfoSpaces(m, l0, v123^v560-v560)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L13
	} else {
		goto L193
	}
L193:
	;
	v34 = v122 + int32(1)
	goto L7
L194:
	;
	v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v568)+296)))
	if v569 != 0 {
		goto L197
	} else {
		goto L198
	}
L195:
	;
	v587 = int32(532200)
	goto L196
L196:
	;
	if v123 != 0 {
		goto L201
	} else {
		goto L202
	}
L197:
	;
	v583 = v568
	goto L199
L198:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v568)+140))
	v576 = int32(0)
	v579 = F_pg_getnameinfo_all(m, v568+int32(12), v572, v568+int32(296), int32(64), v576, v576, int32(3))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L13
	} else {
		goto L200
	}
L199:
	;
	v587 = v583 + int32(296)
	goto L196
L200:
	;
	v582 = *(*int32)(unsafe.Add(mBase, _consts[378]))
	v583 = v582
	goto L199
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+324)) = v587
	*(*int32)(unsafe.Add(mBase, uint32(v11)+320)) = v123
	F_appendStringInfo(m, l0, int32(185605), v11+int32(320))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L13
	} else {
		goto L204
	}
L202:
	;
	goto L203
L203:
	;
	F_appendStringInfoString(m, l0, v587)
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L13
	} else {
		goto L205
	}
L204:
	;
	v34 = v122 + int32(1)
	goto L7
L205:
	;
	v34 = v122 + int32(1)
	goto L7
L206:
	;
	if v123 == int32(0) {
		goto L15
	} else {
		goto L223
	}
L207:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v602)+276))
	if v605 == int32(0) {
		goto L206
	} else {
		goto L208
	}
L208:
	;
	if v123 != 0 {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v602)+292))
	if v608 == int32(0) {
		goto L212
	} else {
		goto L213
	}
L210:
	;
	goto L211
L211:
	;
	F_appendStringInfoString(m, l0, v605)
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L13
	} else {
		goto L219
	}
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+356)) = v605
	*(*int32)(unsafe.Add(mBase, uint32(v11)+352)) = v123
	F_appendStringInfo(m, l0, int32(185605), v11+int32(352))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L13
	} else {
		goto L218
	}
L213:
	;
	v611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v608))))
	if v611 == int32(0) {
		goto L212
	} else {
		goto L214
	}
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+388)) = v608
	*(*int32)(unsafe.Add(mBase, uint32(v11)+384)) = v605
	v619 = F_psprintf(m, int32(704542), v11+int32(384))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L13
	} else {
		goto L215
	}
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+372)) = v619
	*(*int32)(unsafe.Add(mBase, uint32(v11)+368)) = v123
	F_appendStringInfo(m, l0, int32(185605), v11+int32(368))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L13
	} else {
		goto L216
	}
L216:
	;
	F_pfree(m, v619)
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L13
	} else {
		goto L217
	}
L217:
	;
	v34 = v122 + int32(1)
	goto L7
L218:
	;
	v34 = v122 + int32(1)
	goto L7
L219:
	;
	v644 = *(*int32)(unsafe.Add(mBase, _consts[378]))
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v644)+292))
	if v645 == int32(0) {
		goto L15
	} else {
		goto L220
	}
L220:
	;
	v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v645))))
	if v648 == int32(0) {
		goto L15
	} else {
		goto L221
	}
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+336)) = v645
	F_appendStringInfo(m, l0, int32(705414), v11+int32(336))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L13
	} else {
		goto L222
	}
L222:
	;
	v34 = v122 + int32(1)
	goto L7
L223:
	;
	v663 = v123 >> (uint(int32(31)) % 32)
	F_appendStringInfoSpaces(m, l0, v123^v663-v663)
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L13
	} else {
		goto L224
	}
L224:
	;
	v34 = v122 + int32(1)
	goto L7
L225:
	;
	if v123 == int32(0) {
		goto L15
	} else {
		goto L233
	}
L226:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v671)+276))
	if v674 == int32(0) {
		goto L225
	} else {
		goto L227
	}
L227:
	;
	if v123 != 0 {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+404)) = v674
	*(*int32)(unsafe.Add(mBase, uint32(v11)+400)) = v123
	F_appendStringInfo(m, l0, int32(185605), v11+int32(400))
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L13
	} else {
		goto L231
	}
L229:
	;
	goto L230
L230:
	;
	F_appendStringInfoString(m, l0, v674)
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L13
	} else {
		goto L232
	}
L231:
	;
	v34 = v122 + int32(1)
	goto L7
L232:
	;
	v34 = v122 + int32(1)
	goto L7
L233:
	;
	v694 = v123 >> (uint(int32(31)) % 32)
	F_appendStringInfoSpaces(m, l0, v123^v694-v694)
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L13
	} else {
		goto L234
	}
L234:
	;
	v34 = v122 + int32(1)
	goto L7
L235:
	;
	if v123 == int32(0) {
		goto L15
	} else {
		goto L244
	}
L236:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v702)+52))
	if v705 == int32(-1) {
		goto L235
	} else {
		goto L237
	}
L237:
	;
	if v123 != 0 {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v702)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+452)) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v11)+448)) = v705
	v717 = F_pg_snprintf(m, v11+int32(544), int32(127), int32(41907), v11+int32(448))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L13
	} else {
		goto L241
	}
L239:
	;
	goto L240
L240:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v702)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+420)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v11)+416)) = v705
	F_appendStringInfo(m, l0, int32(41907), v11+int32(416))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L13
	} else {
		goto L243
	}
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+432)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v11)+436)) = v11 + int32(544)
	F_appendStringInfo(m, l0, int32(185605), v11+int32(432))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L13
	} else {
		goto L242
	}
L242:
	;
	v34 = v122 + int32(1)
	goto L7
L243:
	;
	v34 = v122 + int32(1)
	goto L7
L244:
	;
	v744 = v123 >> (uint(int32(31)) % 32)
	F_appendStringInfoSpaces(m, l0, v123^v744-v744)
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L13
	} else {
		goto L245
	}
L245:
	;
	v34 = v122 + int32(1)
	goto L7
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+484)) = v752
	*(*int32)(unsafe.Add(mBase, uint32(v11)+480)) = v123
	F_appendStringInfo(m, l0, int32(39991), v11+int32(480))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L13
	} else {
		goto L249
	}
L247:
	;
	goto L248
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+464)) = v752
	F_appendStringInfo(m, l0, int32(64823), v11+int32(464))
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L13
	} else {
		goto L250
	}
L249:
	;
	v34 = v122 + int32(1)
	goto L7
L250:
	;
	v34 = v122 + int32(1)
	goto L7
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+500)) = int32(4547864)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+496)) = v123
	F_appendStringInfo(m, l0, int32(185605), v11+int32(496))
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L13
	} else {
		goto L254
	}
L252:
	;
	goto L253
L253:
	;
	F_appendStringInfoString(m, l0, int32(4547864))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L13
	} else {
		goto L255
	}
L254:
	;
	v34 = v122 + int32(1)
	goto L7
L255:
	;
	v34 = v122 + int32(1)
	goto L7
L256:
	;
	if v123 != 0 {
		goto L260
	} else {
		goto L261
	}
L257:
	;
	v834 = int64(0)
	goto L256
L258:
	;
	goto L259
L259:
	;
	v833 = *(*int64)(unsafe.Add(mBase, uint32(v829)+392))
	v834 = v833
	goto L256
L260:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+536)) = v834
	*(*int32)(unsafe.Add(mBase, uint32(v11)+528)) = v123
	F_appendStringInfo(m, l0, int32(449636), v11+int32(528))
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L13
	} else {
		goto L263
	}
L261:
	;
	goto L262
L262:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+512)) = v834
	F_appendStringInfo(m, l0, int32(450270), v11+int32(512))
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L13
	} else {
		goto L264
	}
L263:
	;
	v34 = v122 + int32(1)
	goto L7
L264:
	;
	goto L15
L265:
	;
	goto L5
L266:
	;
	v34 = v34 + int32(2)
	goto L7
}
