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
	F_SetConfigOption(m, int32(400938), int32(528190), int32(5), int32(10))
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
			F_SetConfigOption(m, int32(334407), int32(785340), int32(5), int32(10))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, _consts[147]))
				v34 = F_AllocSetContextCreateInternal(m, v29, int32(67212), int32(0), int32(8192), int32(8388608))
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
						v39 = int32(4548768)
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
											F_errmsg(m, int32(241566), v6)
											mBase = m.M
											v72 = m.ExcPending
											if v72 != 0 {
												return
											} else {
												F_errfinish(m, int32(515057), int32(4697), int32(230524))
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
										F_SetConfigOption(m, int32(106642), v128, int32(4), int32(10))
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
																F_errmsg(m, int32(461545), v4+int32(-32))
																mBase = m.M
																v195 = m.ExcPending
																if v195 != 0 {
																	return
																} else {
																	v198 = int32(4742)
																	F_errfinish(m, int32(515057), v198, int32(230524))
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
																	F_errmsg(m, int32(461545), v4+int32(-32))
																	mBase = m.M
																	v195 = m.ExcPending
																	if v195 != 0 {
																		return
																	} else {
																		v198 = int32(4742)
																		F_errfinish(m, int32(515057), v198, int32(230524))
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
																		F_errmsg(m, int32(461612), v4+int32(-48))
																		mBase = m.M
																		v171 = m.ExcPending
																		if v171 != 0 {
																			return
																		} else {
																			v198 = int32(4738)
																			F_errfinish(m, int32(515057), v198, int32(230524))
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
												F_errmsg(m, int32(241741), v4+int32(-16))
												mBase = m.M
												v109 = m.ExcPending
												if v109 != 0 {
													return
												} else {
													F_errfinish(m, int32(515057), int32(4713), int32(230524))
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
																F_SetConfigOption(m, int32(106642), v128, int32(4), int32(10))
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
																						F_errmsg(m, int32(461545), v4+int32(-32))
																						mBase = m.M
																						v195 = m.ExcPending
																						if v195 != 0 {
																							return
																						} else {
																							v198 = int32(4742)
																							F_errfinish(m, int32(515057), v198, int32(230524))
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
																							F_errmsg(m, int32(461545), v4+int32(-32))
																							mBase = m.M
																							v195 = m.ExcPending
																							if v195 != 0 {
																								return
																							} else {
																								v198 = int32(4742)
																								F_errfinish(m, int32(515057), v198, int32(230524))
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
																								F_errmsg(m, int32(461612), v4+int32(-48))
																								mBase = m.M
																								v171 = m.ExcPending
																								if v171 != 0 {
																									return
																								} else {
																									v198 = int32(4738)
																									F_errfinish(m, int32(515057), v198, int32(230524))
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
														F_SetConfigOption(m, int32(106642), v128, int32(4), int32(10))
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
																				F_errmsg(m, int32(461545), v4+int32(-32))
																				mBase = m.M
																				v195 = m.ExcPending
																				if v195 != 0 {
																					return
																				} else {
																					v198 = int32(4742)
																					F_errfinish(m, int32(515057), v198, int32(230524))
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
																					F_errmsg(m, int32(461545), v4+int32(-32))
																					mBase = m.M
																					v195 = m.ExcPending
																					if v195 != 0 {
																						return
																					} else {
																						v198 = int32(4742)
																						F_errfinish(m, int32(515057), v198, int32(230524))
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
																						F_errmsg(m, int32(461612), v4+int32(-48))
																						mBase = m.M
																						v171 = m.ExcPending
																						if v171 != 0 {
																							return
																						} else {
																							v198 = int32(4738)
																							F_errfinish(m, int32(515057), v198, int32(230524))
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
	v93 = int32(506545)
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
	v93 = int32(506559)
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
	F_errfinish(m, int32(511214), v181, int32(116386))
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
	v140 = int32(255843)
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v122
	F_errmsg(m, int32(209217), v14+int32(16))
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
	F_errdetail_log_plural(m, int32(628762), int32(628787), v117, v14)
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
	v171 = int32(255843)
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v161
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v122
	F_errmsg(m, int32(209262), v14+int32(32))
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
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
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
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
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v384 int64
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int64
	_ = v430
	var v431 int32
	_ = v431
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v467 int32
	_ = v467
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	var v602 int64
	_ = v602
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v674 int32
	_ = v674
	var v677 int64
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v699 int64
	_ = v699
	var v703 int32
	_ = v703
	var v707 int64
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v717 int64
	_ = v717
	var v723 int32
	_ = v723
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v754 int32
	_ = v754
	v16 = m.G0
	v18 = v16 - int32(96)
	m.G0 = v18
	v22 = m.G0
	v24 = v22 - int32(32)
	m.G0 = v24
	v27 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v31 = F_LWLockAcquire(m, v27+int32(23296), int32(1))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v40 = F_LWLockAcquire(m, v36+int32(23424), int32(1))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v47 = F_LWLockAcquire(m, v43+int32(23552), int32(1))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v54 = F_LWLockAcquire(m, v50+int32(23680), int32(1))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v61 = F_LWLockAcquire(m, v57+int32(23808), int32(1))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v68 = F_LWLockAcquire(m, v64+int32(23936), int32(1))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v75 = F_LWLockAcquire(m, v71+int32(24064), int32(1))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v82 = F_LWLockAcquire(m, v78+int32(24192), int32(1))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v89 = F_LWLockAcquire(m, v85+int32(24320), int32(1))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v96 = F_LWLockAcquire(m, v92+int32(24448), int32(1))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v103 = F_LWLockAcquire(m, v99+int32(24576), int32(1))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v110 = F_LWLockAcquire(m, v106+int32(24704), int32(1))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v117 = F_LWLockAcquire(m, v113+int32(24832), int32(1))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v124 = F_LWLockAcquire(m, v120+int32(24960), int32(1))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v131 = F_LWLockAcquire(m, v127+int32(25088), int32(1))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v138 = F_LWLockAcquire(m, v134+int32(25216), int32(1))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v141 = *(*int32)(unsafe.Add(mBase, _consts[696]))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v143)+412))
	if v145 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v213 = F_palloc(m, v210*int32(12))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L22
	}
L19:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v143)+376))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v143)+364))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v143)+352))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v143)+340))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v143)+328))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v143)+316))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v143)+304))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v143)+292))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v143)+280))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v143)+268))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v143)+256))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v143)+244))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v143)+232))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v143)+220))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v143)+208))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v143)+196))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v143)+184))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v143)+172))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v143)+160))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v143)+148))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v143)+136))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v143)+124))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v143)+112))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v143)+100))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v143)+88))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v143)+76))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v143-int32(-64))))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v143)+52))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v143)+40))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v143)+28))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v143)+16))
	v210 = v146 + (v147 + (v148 + (v149 + (v150 + (v151 + (v152 + (v153 + (v154 + (v155 + (v156 + (v157 + (v158 + (v159 + (v160 + (v161 + (v162 + (v163 + (v164 + (v165 + (v166 + (v167 + (v168 + (v169 + (v170 + (v171 + (v174 + (v175 + (v176 + (v177 + (v178 + v144))))))))))))))))))))))))))))))
	goto L21
L20:
	;
	v210 = v144
	goto L21
L21:
	;
	goto L18
L22:
	;
	v218 = *(*int32)(unsafe.Add(mBase, _consts[696]))
	F_hash_seq_init(m, v24+int32(12), v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v224 = int32(0)
	goto L24
L24:
	;
	v238 = F_hash_seq_search(m, v24+int32(12))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	v262 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v262+int32(25216))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L33
	}
L26:
	;
	if v238 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238)+13)))
	if v240&int32(1) == int32(0) {
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
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v238)))
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245)+14)))
	if v246 != 0 {
		goto L24
	} else {
		goto L31
	}
L31:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v238)+4))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v247)+36))
	if v248 == int32(0) {
		goto L24
	} else {
		goto L32
	}
L32:
	;
	v253 = v213 + v224*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v253))) = v248
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v245)))
	*(*int32)(unsafe.Add(mBase, uint32(v253)+4)) = v255
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v245)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v253)+8)) = v257
	v224 = v224 + int32(1)
	goto L24
L33:
	;
	v268 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v268+int32(25088))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v274 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v274+int32(24960))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v280 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v280+int32(24832))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v286 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v286+int32(24704))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v292 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v292+int32(24576))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v298 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v298+int32(24448))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v304 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v304+int32(24320))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v310 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v310+int32(24192))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v316 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v316+int32(24064))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v322 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v322+int32(23936))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v328 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v328+int32(23808))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v334 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v334+int32(23680))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v340 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v340+int32(23552))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v346 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v346+int32(23424))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v352 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v352+int32(23296))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18+int32(68)))) = v224
	m.G0 = v24 + int32(32)
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v18)+68))
	if int32(0) < v361 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+72)) = v361
	F_XLogBeginInsert(m)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	F_pfree(m, v213)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L1
	} else {
		goto L57
	}
L52:
	;
	F_XLogRegisterData(m, v18+int32(72), int32(4))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_XLogRegisterData(m, v213, v361*int32(12))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v377 = int32(4444612)
	v379 = int32(*(*uint8)(unsafe.Add(mBase, _consts[44])))
	v380 = v379 | int32(2)
	*(*uint8)(unsafe.Add(mBase, _consts[44])) = uint8(v380)
	goto L55
L55:
	;
	v384 = F_XLogInsert(m, int32(8), int32(0))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	goto L51
L57:
	;
	v388 = int32(0)
	v392 = *(*int32)(unsafe.Add(mBase, _consts[513]))
	v394 = *(*int32)(unsafe.Add(mBase, _consts[102]))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v394)+4))
	v397 = *(*int32)(unsafe.Add(mBase, _consts[697]))
	if v397 == v388 {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	v629 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if v629 <= int32(1) {
		goto L113
	} else {
		goto L114
	}
L59:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L1
	} else {
		goto L109
	}
L60:
	;
	v402 = *(*int32)(unsafe.Add(mBase, _consts[103]))
	v404 = *(*int32)(unsafe.Add(mBase, _consts[657]))
	v408 = F_emscripten_builtin_malloc(m, (v402+v404)*int32(260))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _consts[697])) = v408
	if v408 == int32(0) {
		goto L59
	} else {
		goto L63
	}
L61:
	;
	v412 = v397
	goto L62
L62:
	;
	v414 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v418 = F_LWLockAcquire(m, v414+int32(512), int32(1))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L1
	} else {
		goto L64
	}
L63:
	;
	v412 = v408
	goto L62
L64:
	;
	v421 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v425 = F_LWLockAcquire(m, v421+int32(384), int32(1))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v428 = *(*int32)(unsafe.Add(mBase, _consts[69]))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v428)+8))
	v430 = *(*int64)(unsafe.Add(mBase, uint32(v428)+48))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v392)))
	if v431 <= int32(0) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, _consts[698])) = v580
	*(*int32)(unsafe.Add(mBase, _consts[699])) = v583
	v597 = int32(4465496)
	*(*int32)(unsafe.Add(mBase, _consts[700])) = v585 - v583
	v601 = *(*int32)(unsafe.Add(mBase, _consts[69]))
	v602 = *(*int64)(unsafe.Add(mBase, uint32(v601)+8))
	*(*uint32)(unsafe.Add(mBase, _consts[701])) = uint32(v430)
	*(*int32)(unsafe.Add(mBase, _consts[702])) = v581
	*(*int32)(unsafe.Add(mBase, _consts[703])) = v578
	*(*uint32)(unsafe.Add(mBase, _consts[704])) = uint32(v602)
	goto L58
L67:
	;
	v578 = v429
	v580 = v388
	v581 = v429
	v583 = v388
	v585 = v388
	goto L66
L68:
	;
	goto L69
L69:
	;
	v437 = v429
	v440 = v429
	v441 = int32(0)
	v442 = v388
	v443 = v431
	v444 = v388
	goto L70
L70:
	;
	v453 = v441 << (uint(int32(2)) % 32)
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v395+v453)))
	if v455 != 0 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	if v511&int32(1) != 0 {
		goto L93
	} else {
		goto L94
	}
L72:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v437))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v455)) == int32(0) {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	v509 = v437
	v510 = v440
	v511 = v442
	v512 = v443
	v513 = v444
	goto L74
L74:
	;
	v515 = v441 + int32(1)
	if v515 < v512 {
		v437 = v509
		v440 = v510
		v441 = v515
		v442 = v511
		v443 = v512
		v444 = v513
		goto L70
	} else {
		goto L92
	}
L75:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v440))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v455)) == int32(0) {
		goto L80
	} else {
		goto L81
	}
L76:
	;
	v467 = base.B2i32(base.Ui32(v455) < base.Ui32(v437))
	goto L75
L77:
	;
	goto L78
L78:
	;
	v467 = int32(base.Ui32(v455-v437) >> (uint(int32(31)) % 32))
	goto L75
L79:
	;
	if v479 != 0 {
		goto L83
	} else {
		goto L84
	}
L80:
	;
	v479 = base.B2i32(base.Ui32(v455) < base.Ui32(v440))
	goto L79
L81:
	;
	goto L82
L82:
	;
	v479 = int32(base.Ui32(v455-v440) >> (uint(int32(31)) % 32))
	goto L79
L83:
	;
	v481 = *(*int32)(unsafe.Add(mBase, _consts[514]))
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v392+int32(36)+v453)))
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v481+v483*int32(640))+60))
	v489 = *(*int32)(unsafe.Add(mBase, _consts[108]))
	if v487 == v489 {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	v492 = v440
	goto L85
L85:
	;
	if v467 != 0 {
		goto L89
	} else {
		goto L90
	}
L86:
	;
	v491 = v455
	goto L88
L87:
	;
	v491 = v440
	goto L88
L88:
	;
	v492 = v491
	goto L85
L89:
	;
	v493 = v455
	goto L91
L90:
	;
	v493 = v437
	goto L91
L91:
	;
	v495 = *(*int32)(unsafe.Add(mBase, _consts[102]))
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v495)+8))
	v497 = int32(1)
	v500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496+v441<<(uint(v497)%32))+1)))
	*(*int32)(unsafe.Add(mBase, uint32(v412+v444<<(uint(int32(2))%32)))) = v455
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v392)))
	v509 = v493
	v510 = v492
	v511 = v442 | v500
	v512 = v508
	v513 = v444 + v497
	goto L74
L92:
	;
	goto L71
L93:
	;
	v578 = v509
	v580 = int32(2)
	v581 = v510
	v583 = int32(0)
	v585 = v513
	goto L66
L94:
	;
	goto L95
L95:
	;
	v521 = int32(0)
	if v512 <= v521 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v578 = v509
	v580 = v521
	v581 = v510
	v583 = int32(0)
	v585 = v513
	goto L66
L97:
	;
	goto L98
L98:
	;
	v526 = *(*int32)(unsafe.Add(mBase, _consts[102]))
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v526)+8))
	v533 = v521
	v536 = int32(0)
	v537 = v512
	v538 = v513
	goto L99
L99:
	;
	v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v527+v533<<(uint(int32(1))%32)))))
	if v549 != 0 {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v578 = v509
	v580 = int32(0)
	v581 = v510
	v583 = v571
	v585 = v573
	goto L66
L101:
	;
	v550 = int32(2)
	v554 = *(*int32)(unsafe.Add(mBase, _consts[514]))
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v392+int32(36)+v533<<(uint(v550)%32))))
	v565 = v549 << (uint(v550) % 32)
	if v565 != 0 {
		goto L105
	} else {
		goto L106
	}
L102:
	;
	v571 = v536
	v572 = v537
	v573 = v538
	goto L103
L103:
	;
	v575 = v533 + int32(1)
	if v575 < v572 {
		v533 = v575
		v536 = v571
		v537 = v572
		v538 = v573
		goto L99
	} else {
		goto L108
	}
L104:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v392)))
	v571 = v549 + v536
	v572 = v569
	v573 = v549 + v538
	goto L103
L105:
	;
	v566 = F__emscripten_memcpy_bulkmem(m, v412+v538<<(uint(v550)%32), v554+v558*int32(640)+int32(280), v565)
	mBase = m.M
	goto L107
L106:
	;
	goto L107
L107:
	;
	goto L104
L108:
	;
	goto L100
L109:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	F_errmsg(m, int32(13961), int32(0))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(511224), int32(2727), int32(526208))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L113:
	;
	v633 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v633+int32(512))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L1
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v638 = *(*int32)(unsafe.Add(mBase, _consts[700]))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+72)) = v638
	v640 = *(*int32)(unsafe.Add(mBase, _consts[699]))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+76)) = v640
	v642 = *(*int32)(unsafe.Add(mBase, _consts[698]))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+80)) = uint8(base.B2i32(v642 != int32(0)))
	v646 = *(*int32)(unsafe.Add(mBase, _consts[704]))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+84)) = v646
	v648 = *(*int32)(unsafe.Add(mBase, _consts[703]))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+88)) = v648
	v650 = *(*int32)(unsafe.Add(mBase, _consts[701]))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+92)) = v650
	F_XLogBeginInsert(m)
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L1
	} else {
		goto L117
	}
L116:
	;
	goto L115
L117:
	;
	v655 = int32(4444612)
	v657 = int32(*(*uint8)(unsafe.Add(mBase, _consts[44])))
	v658 = v657 | int32(2)
	*(*uint8)(unsafe.Add(mBase, _consts[44])) = uint8(v658)
	goto L118
L118:
	;
	F_XLogRegisterData(m, v18+int32(72), int32(24))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v18)+72))
	if int32(0) < v665 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v668 = *(*int32)(unsafe.Add(mBase, _consts[697]))
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v18)+76))
	F_XLogRegisterData(m, v668, (v669+v665)<<(uint(int32(2))%32))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L1
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v677 = F_XLogInsert(m, int32(8), int32(16))
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L1
	} else {
		goto L124
	}
L123:
	;
	goto L122
L124:
	;
	v679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+80)))
	v682 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	if v679 == int32(1) {
		goto L128
	} else {
		goto L129
	}
L126:
	;
	F_XLogSetAsyncXactLSN(m, v677)
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L1
	} else {
		goto L136
	}
L127:
	;
	F_errfinish(m, int32(511214), v729, int32(132717))
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L1
	} else {
		goto L135
	}
L128:
	;
	if v682 == int32(0) {
		goto L126
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	if v682 == int32(0) {
		goto L126
	} else {
		goto L133
	}
L131:
	;
	v689 = *(*int32)(unsafe.Add(mBase, _consts[700]))
	v690 = *(*int32)(unsafe.Add(mBase, _consts[703]))
	v691 = *(*int32)(unsafe.Add(mBase, _consts[701]))
	v692 = *(*int32)(unsafe.Add(mBase, _consts[704]))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v692
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v691
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v690
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v689
	*(*uint32)(unsafe.Add(mBase, uint32(v18)+8)) = uint32(v677)
	v699 = int64(base.Ui64(v677) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v18)+4)) = uint32(v699)
	F_errmsg_internal(m, int32(696384), v18)
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	v729 = int32(1384)
	goto L127
L133:
	;
	v707 = *(*int64)(unsafe.Add(mBase, _consts[700]))
	v708 = *(*int32)(unsafe.Add(mBase, _consts[704]))
	v709 = *(*int32)(unsafe.Add(mBase, _consts[701]))
	v710 = *(*int32)(unsafe.Add(mBase, _consts[703]))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v710
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v709
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v708
	*(*int64)(unsafe.Add(mBase, uint32(v18)+32)) = v707
	*(*uint32)(unsafe.Add(mBase, uint32(v18)+44)) = uint32(v677)
	v717 = int64(base.Ui64(v677) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v18)+40)) = uint32(v717)
	F_errmsg_internal(m, int32(696285), v18+int32(32))
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	v729 = int32(1392)
	goto L127
L135:
	;
	goto L126
L136:
	;
	v740 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) <= v740 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v744 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v744+int32(512))
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L1
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v750 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v750+int32(384))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L1
	} else {
		goto L141
	}
L140:
	;
	goto L139
L141:
	;
	m.G0 = v18 + int32(96)
	return v677
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
		v9 = int32(*(*uint8)(unsafe.Add(mBase, _consts[837])))
		if v9 != 0 {
			v18 = *(*int32)(unsafe.Add(mBase, _consts[87]))
			*(*int32)(unsafe.Add(mBase, _consts[88])) = v18
			v24 = F_format_elog_string(m, int32(651394), int32(0))
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
			v11 = int32(*(*uint8)(unsafe.Add(mBase, _consts[838])))
			if v11 != 0 {
				v18 = *(*int32)(unsafe.Add(mBase, _consts[87]))
				*(*int32)(unsafe.Add(mBase, _consts[88])) = v18
				v24 = F_format_elog_string(m, int32(651394), int32(0))
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
				v14 = int32(*(*uint8)(unsafe.Add(mBase, _consts[839])))
				if v14 != v12 {
					v30 = v12
					return v30
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, _consts[87]))
					*(*int32)(unsafe.Add(mBase, _consts[88])) = v18
					v24 = F_format_elog_string(m, int32(651394), int32(0))
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
	v17 = *(*int32)(unsafe.Add(mBase, _consts[1202]))
	if v15 == v17 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1203])) = v29
	if l1 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _consts[1203]))
	v29 = v20 + int32(1)
	goto L1
L3:
	;
	goto L4
L4:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1202])) = v15
	v26 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[1190])) = uint8(v26)
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
	*(*uint8)(unsafe.Add(mBase, _consts[1156])) = uint8(v775)
	v783 = int32(base.Ui32(v771)>>(uint(int32(24))%32))&v772 + v774
	*(*uint8)(unsafe.Add(mBase, _consts[1157])) = uint8(v783)
	v791 = int32(base.Ui32(v771)>>(uint(int32(18))%32))&v772 + v774
	*(*uint8)(unsafe.Add(mBase, _consts[1158])) = uint8(v791)
	v799 = int32(base.Ui32(v771)>>(uint(int32(12))%32))&v772 + v774
	*(*uint8)(unsafe.Add(mBase, _consts[1159])) = uint8(v799)
	v807 = int32(base.Ui32(v771)>>(uint(int32(6))%32))&v772 + v774
	*(*uint8)(unsafe.Add(mBase, _consts[1160])) = uint8(v807)
	v810 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[1161])) = uint8(v810)
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
	v507 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1190])))
	if v507 == int32(0) {
		goto L173
	} else {
		goto L174
	}
L44:
	;
	v462 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1147])))
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
	*(*uint8)(unsafe.Add(mBase, _consts[1148])) = uint8(v375)
	v378 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1147])))
	if v378 == v375 {
		goto L146
	} else {
		goto L147
	}
L47:
	;
	v356 = *(*int32)(unsafe.Add(mBase, _consts[1203]))
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
	v162 = *(*int32)(unsafe.Add(mBase, _consts[1192]))
	if v160 == v162 {
		v183 = int32(224480)
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
	v133 = *(*int32)(unsafe.Add(mBase, _consts[1191]))
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
	v136 = int32(528539)
	goto L58
L62:
	;
	goto L61
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v123
	F_appendStringInfo(m, l0, int32(184212), v11)
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
	v169 = *(*int32)(unsafe.Add(mBase, _consts[812]))
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
	v181 = int32(383012)
	goto L78
L78:
	;
	goto L75
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v123
	F_appendStringInfo(m, l0, int32(184212), v11+int32(16))
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
	v202 = int32(528539)
	goto L87
L91:
	;
	goto L90
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v202
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v123
	F_appendStringInfo(m, l0, int32(184212), v11+int32(32))
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
	v231 = int32(528539)
	goto L102
L106:
	;
	goto L105
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v123
	F_appendStringInfo(m, l0, int32(184212), v11+int32(48))
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
	v267 = F_pg_snprintf(m, v11+int32(544), int32(127), int32(30845), v11+int32(96))
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
	F_appendStringInfo(m, l0, int32(30845), v11-int32(-64))
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
	F_appendStringInfo(m, l0, int32(184212), v11+int32(80))
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
	F_appendStringInfo(m, l0, int32(483440), v11+int32(128))
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
	F_appendStringInfo(m, l0, int32(506559), v11+int32(112))
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
	F_appendStringInfo(m, l0, int32(483440), v11+int32(160))
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
	F_appendStringInfo(m, l0, int32(506559), v11+int32(144))
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
	F_appendStringInfo(m, l0, int32(448735), v11+int32(192))
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
	F_appendStringInfo(m, l0, int32(448890), v11+int32(176))
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
	F___gettimeofday(m, int32(4542352))
	mBase = m.M
	v384 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[1147])) = uint8(v384)
	goto L148
L147:
	;
	goto L148
L148:
	;
	v387 = *(*int64)(unsafe.Add(mBase, _consts[1204]))
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
	v398 = F_pg_strftime(m, int32(4542208), int32(128), int32(529157), v396)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L13
	} else {
		goto L150
	}
L150:
	;
	v401 = *(*int32)(unsafe.Add(mBase, _consts[1205]))
	v403 = base.I32_div_s(v401, int32(1000))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+224)) = v403
	v410 = F_pg_sprintf(m, v11+int32(544), int32(483305), v11+int32(224))
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
	*(*int32)(unsafe.Add(mBase, _consts[1206])) = v413
	if v123 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+212)) = int32(4542208)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+208)) = v123
	F_appendStringInfo(m, l0, int32(184212), v11+int32(208))
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
	F_appendStringInfoString(m, l0, int32(4542208))
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
	v442 = F_pg_strftime(m, v11+int32(544), int32(128), int32(529136), v440)
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
	F_appendStringInfo(m, l0, int32(184212), v11+int32(240))
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
	F___gettimeofday(m, int32(4542352))
	mBase = m.M
	v468 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[1147])) = uint8(v468)
	goto L166
L165:
	;
	goto L166
L166:
	;
	v471 = *(*int64)(unsafe.Add(mBase, _consts[1204]))
	*(*uint32)(unsafe.Add(mBase, uint32(v11)+272)) = uint32(v471)
	v474 = *(*int32)(unsafe.Add(mBase, _consts[1205]))
	v476 = base.I32_div_s(v474, int32(1000))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+276)) = v476
	v484 = F_pg_snprintf(m, v11+int32(544), int32(128), int32(483302), v11+int32(272))
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
	F_appendStringInfo(m, l0, int32(184212), v11+int32(256))
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
	v519 = F_pg_strftime(m, int32(4542368), int32(128), int32(529136), v517)
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
	*(*int32)(unsafe.Add(mBase, uint32(v11)+292)) = int32(4542368)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+288)) = v123
	F_appendStringInfo(m, l0, int32(184212), v11+int32(288))
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
	F_appendStringInfoString(m, l0, int32(4542368))
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
	v542 = int32(785340)
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
	F_appendStringInfo(m, l0, int32(184212), v11+int32(304))
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
	v587 = int32(528557)
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
	F_appendStringInfo(m, l0, int32(184212), v11+int32(320))
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
	F_appendStringInfo(m, l0, int32(184212), v11+int32(352))
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
	v619 = F_psprintf(m, int32(700005), v11+int32(384))
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
	F_appendStringInfo(m, l0, int32(184212), v11+int32(368))
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
	F_appendStringInfo(m, l0, int32(700877), v11+int32(336))
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
	F_appendStringInfo(m, l0, int32(184212), v11+int32(400))
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
	v717 = F_pg_snprintf(m, v11+int32(544), int32(127), int32(41425), v11+int32(448))
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
	F_appendStringInfo(m, l0, int32(41425), v11+int32(416))
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
	F_appendStringInfo(m, l0, int32(184212), v11+int32(432))
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
	F_appendStringInfo(m, l0, int32(39554), v11+int32(480))
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
	F_appendStringInfo(m, l0, int32(64245), v11+int32(464))
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
	*(*int32)(unsafe.Add(mBase, uint32(v11)+500)) = int32(4542504)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+496)) = v123
	F_appendStringInfo(m, l0, int32(184212), v11+int32(496))
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
	F_appendStringInfoString(m, l0, int32(4542504))
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
	F_appendStringInfo(m, l0, int32(446534), v11+int32(528))
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
	F_appendStringInfo(m, l0, int32(447142), v11+int32(512))
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
