package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SubscriptionConninfo(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
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
	var v33 int64
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
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
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
	v12 = F_SearchSysCache1(m, int32(67), v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 != 0 {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+22)))
			v18 = v16 + v17
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+104))
			if v19 != 0 {
				v20 = F_GetForeignServer(m, v19)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)+104))
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v26 = F_object_aclcheck(m, int32(1417), v23, v24, int64(256))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						if v26 != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(16797828))
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
									v69 = F_GetUserNameFromId(m, v67, int32(0))
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										v71 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
										*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v71
										*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v69
										F_errmsg(m, int32(_a_F_SubscriptionConninfo_0), v8+int32(16))
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_SubscriptionConninfo_1), int32(220), int32(_a_F_SubscriptionConninfo_2))
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
												return int32(0)
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
							v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							v29 = F_ForeignServerConnectionString(m, v28, v20)
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return int32(0)
							} else {
								v39 = v29
								F_ReleaseCatCache(m, v12)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int32(0)
								} else {
									m.G0 = v8 + int32(32)
									return v39
								}
							}
						}
					}
				}
			} else {
				v33 = F_SysCacheGetAttrNotNull(m, int32(67), v12, int32(18))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v36 = F_text_to_cstring(m, base.I32_wrap_i64(v33))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						v39 = v36
						F_ReleaseCatCache(m, v12)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(32)
							return v39
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return int32(0)
			} else {
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v50
				F_errmsg_internal(m, int32(_a_F_SubscriptionConninfo_3), v8)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_SubscriptionConninfo_1), int32(201), int32(_a_F_SubscriptionConninfo_2))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
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
func F_clear_subscription_skip_lsn(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v1 int64
	_ = v1
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int64
	_ = v61
	var v64 int32
	_ = v64
	var v68 int64
	_ = v68
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v110 int64
	_ = v110
	var v111 int64
	_ = v111
	var v115 int64
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	v1 = l0
	v8 = m.G0
	v10 = v8 - int32(304)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_clear_subscription_skip_lsn[0]))
	v14 = *(*int64)(unsafe.Add(mBase, uint32(v13)+16))
	if v14 == int64(0) {
		m.G0 = v10 + int32(304)
		return
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_clear_subscription_skip_lsn[1]))
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+16)))
		if v19 == int32(1) {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
			if v22 == int32(4) {
				m.G0 = v10 + int32(304)
				return
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, _c_F_clear_subscription_skip_lsn[2]))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
				v29 = base.B2i32(v27 == int32(2))
				if v29 == int32(0) {
					F_StartTransactionCommand(m)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						v34 = F_GetTransactionSnapshot(m)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return
						} else {
							F_PushActiveSnapshot(m, v34)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return
							} else {
								v40 = *(*int32)(unsafe.Add(mBase, _c_F_clear_subscription_skip_lsn[0]))
								v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
								F_LockSharedObject(m, int32(_a_F_clear_subscription_skip_lsn_0), v41, int32(1))
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return
								} else {
									v47 = F_table_open(m, int32(_a_F_clear_subscription_skip_lsn_0), int32(3))
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return
									} else {
										v51 = *(*int32)(unsafe.Add(mBase, _c_F_clear_subscription_skip_lsn[0]))
										v52 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v51)+4)))
										v54 = F_SearchSysCacheCopy(m, int32(67), v52, int64(0))
										mBase = m.M
										v55 = m.ExcPending
										if v55 != 0 {
											return
										} else {
											if v54 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v148 = m.ExcPending
												if v148 != 0 {
													return
												} else {
													v150 = *(*int32)(unsafe.Add(mBase, _c_F_clear_subscription_skip_lsn[0]))
													v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+24))
													*(*int32)(unsafe.Add(mBase, uint32(v10))) = v151
													F_errmsg_internal(m, int32(_a_F_clear_subscription_skip_lsn_1), v10)
													mBase = m.M
													v155 = m.ExcPending
													if v155 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_clear_subscription_skip_lsn_2), int32(_a_F_clear_subscription_skip_lsn_3), int32(_a_F_clear_subscription_skip_lsn_4))
														mBase = m.M
														v160 = m.ExcPending
														if v160 != 0 {
															return
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
												v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+22)))
												v61 = *(*int64)(unsafe.Add(mBase, uint32(v58+v59)+8))
												if v61 != v14 {
													v127 = v54
													F_pfree(m, v127)
													mBase = m.M
													v130 = m.ExcPending
													if v130 != 0 {
														return
													} else {
														F_relation_close(m, v47, int32(0))
														mBase = m.M
														v133 = m.ExcPending
														if v133 != 0 {
															return
														} else {
															F_PopActiveSnapshot(m)
															mBase = m.M
															v135 = m.ExcPending
															if v135 != 0 {
																return
															} else {
																if v27 == int32(2) {
																	m.G0 = v10 + int32(304)
																	return
																} else {
																	F_CommitTransactionCommand(m)
																	mBase = m.M
																	v137 = m.ExcPending
																	if v137 != 0 {
																		return
																	} else {
																		m.G0 = v10 + int32(304)
																		return
																	}
																}
															}
														}
													}
												} else {
													v64 = v10 + int32(48)
													base.MemoryFill(m, v64, int32(0), int32(184))
													v68 = int64(0)
													*(*int64)(unsafe.Add(mBase, uint32(v10)+240)) = v68
													*(*int64)(unsafe.Add(mBase, uint32(v10)+287)) = v68
													*(*int64)(unsafe.Add(mBase, uint32(v10)+280)) = v68
													*(*int64)(unsafe.Add(mBase, uint32(v10)+272)) = v68
													*(*int64)(unsafe.Add(mBase, uint32(v10)+248)) = v68
													*(*int64)(unsafe.Add(mBase, uint32(v10)+255)) = v68
													v80 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v10)+242)) = uint8(v80)
													v82 = *(*int32)(unsafe.Add(mBase, uint32(v47)+52))
													v87 = F_heap_modify_tuple(m, v54, v82, v64, v10+int32(272), v10+int32(240))
													mBase = m.M
													v88 = m.ExcPending
													if v88 != 0 {
														return
													} else {
														F_CatalogTupleUpdate(m, v47, v87+int32(4), v87)
														mBase = m.M
														v92 = m.ExcPending
														if v92 != 0 {
															return
														} else {
															if v1 == v14 {
																v127 = v87
																F_pfree(m, v127)
																mBase = m.M
																v130 = m.ExcPending
																if v130 != 0 {
																	return
																} else {
																	F_relation_close(m, v47, int32(0))
																	mBase = m.M
																	v133 = m.ExcPending
																	if v133 != 0 {
																		return
																	} else {
																		F_PopActiveSnapshot(m)
																		mBase = m.M
																		v135 = m.ExcPending
																		if v135 != 0 {
																			return
																		} else {
																			if v27 == int32(2) {
																				m.G0 = v10 + int32(304)
																				return
																			} else {
																				F_CommitTransactionCommand(m)
																				mBase = m.M
																				v137 = m.ExcPending
																				if v137 != 0 {
																					return
																				} else {
																					m.G0 = v10 + int32(304)
																					return
																				}
																			}
																		}
																	}
																}
															} else {
																v96 = F_errstart(m, int32(19), int32(0))
																mBase = m.M
																v97 = m.ExcPending
																if v97 != 0 {
																	return
																} else {
																	if v96 == int32(0) {
																		v127 = v87
																		F_pfree(m, v127)
																		mBase = m.M
																		v130 = m.ExcPending
																		if v130 != 0 {
																			return
																		} else {
																			F_relation_close(m, v47, int32(0))
																			mBase = m.M
																			v133 = m.ExcPending
																			if v133 != 0 {
																				return
																			} else {
																				F_PopActiveSnapshot(m)
																				mBase = m.M
																				v135 = m.ExcPending
																				if v135 != 0 {
																					return
																				} else {
																					if v27 == int32(2) {
																						m.G0 = v10 + int32(304)
																						return
																					} else {
																						F_CommitTransactionCommand(m)
																						mBase = m.M
																						v137 = m.ExcPending
																						if v137 != 0 {
																							return
																						} else {
																							m.G0 = v10 + int32(304)
																							return
																						}
																					}
																				}
																			}
																		}
																	} else {
																		v101 = *(*int32)(unsafe.Add(mBase, _c_F_clear_subscription_skip_lsn[0]))
																		v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+24))
																		*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v102
																		F_errmsg(m, int32(_a_F_clear_subscription_skip_lsn_5), v10+int32(32))
																		mBase = m.M
																		v108 = m.ExcPending
																		if v108 != 0 {
																			return
																		} else {
																			*(*uint32)(unsafe.Add(mBase, uint32(v10)+28)) = uint32(v14)
																			v110 = int64(32)
																			v111 = int64(base.Ui64(v14) >> (uint(v110) % 64))
																			*(*uint32)(unsafe.Add(mBase, uint32(v10)+24)) = uint32(v111)
																			*(*uint32)(unsafe.Add(mBase, uint32(v10)+20)) = uint32(v1)
																			v115 = int64(base.Ui64(v1) >> (uint(v110) % 64))
																			*(*uint32)(unsafe.Add(mBase, uint32(v10)+16)) = uint32(v115)
																			v120 = F_errdetail(m, int32(_a_F_clear_subscription_skip_lsn_6), v10+int32(16))
																			mBase = m.M
																			v121 = m.ExcPending
																			if v121 != 0 {
																				return
																			} else {
																				F_errfinish(m, int32(_a_F_clear_subscription_skip_lsn_2), int32(_a_F_clear_subscription_skip_lsn_7), int32(_a_F_clear_subscription_skip_lsn_4))
																				mBase = m.M
																				v126 = m.ExcPending
																				if v126 != 0 {
																					return
																				} else {
																					v127 = v87
																					F_pfree(m, v127)
																					mBase = m.M
																					v130 = m.ExcPending
																					if v130 != 0 {
																						return
																					} else {
																						F_relation_close(m, v47, int32(0))
																						mBase = m.M
																						v133 = m.ExcPending
																						if v133 != 0 {
																							return
																						} else {
																							F_PopActiveSnapshot(m)
																							mBase = m.M
																							v135 = m.ExcPending
																							if v135 != 0 {
																								return
																							} else {
																								if v27 == int32(2) {
																									m.G0 = v10 + int32(304)
																									return
																								} else {
																									F_CommitTransactionCommand(m)
																									mBase = m.M
																									v137 = m.ExcPending
																									if v137 != 0 {
																										return
																									} else {
																										m.G0 = v10 + int32(304)
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
				} else {
					v34 = F_GetTransactionSnapshot(m)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						F_PushActiveSnapshot(m, v34)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							v40 = *(*int32)(unsafe.Add(mBase, _c_F_clear_subscription_skip_lsn[0]))
							v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
							F_LockSharedObject(m, int32(_a_F_clear_subscription_skip_lsn_0), v41, int32(1))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return
							} else {
								v47 = F_table_open(m, int32(_a_F_clear_subscription_skip_lsn_0), int32(3))
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return
								} else {
									v51 = *(*int32)(unsafe.Add(mBase, _c_F_clear_subscription_skip_lsn[0]))
									v52 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v51)+4)))
									v54 = F_SearchSysCacheCopy(m, int32(67), v52, int64(0))
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return
									} else {
										if v54 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v148 = m.ExcPending
											if v148 != 0 {
												return
											} else {
												v150 = *(*int32)(unsafe.Add(mBase, _c_F_clear_subscription_skip_lsn[0]))
												v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+24))
												*(*int32)(unsafe.Add(mBase, uint32(v10))) = v151
												F_errmsg_internal(m, int32(_a_F_clear_subscription_skip_lsn_1), v10)
												mBase = m.M
												v155 = m.ExcPending
												if v155 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_clear_subscription_skip_lsn_2), int32(_a_F_clear_subscription_skip_lsn_3), int32(_a_F_clear_subscription_skip_lsn_4))
													mBase = m.M
													v160 = m.ExcPending
													if v160 != 0 {
														return
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										} else {
											v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
											v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+22)))
											v61 = *(*int64)(unsafe.Add(mBase, uint32(v58+v59)+8))
											if v61 != v14 {
												v127 = v54
												F_pfree(m, v127)
												mBase = m.M
												v130 = m.ExcPending
												if v130 != 0 {
													return
												} else {
													F_relation_close(m, v47, int32(0))
													mBase = m.M
													v133 = m.ExcPending
													if v133 != 0 {
														return
													} else {
														F_PopActiveSnapshot(m)
														mBase = m.M
														v135 = m.ExcPending
														if v135 != 0 {
															return
														} else {
															if v27 == int32(2) {
																m.G0 = v10 + int32(304)
																return
															} else {
																F_CommitTransactionCommand(m)
																mBase = m.M
																v137 = m.ExcPending
																if v137 != 0 {
																	return
																} else {
																	m.G0 = v10 + int32(304)
																	return
																}
															}
														}
													}
												}
											} else {
												v64 = v10 + int32(48)
												base.MemoryFill(m, v64, int32(0), int32(184))
												v68 = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(v10)+240)) = v68
												*(*int64)(unsafe.Add(mBase, uint32(v10)+287)) = v68
												*(*int64)(unsafe.Add(mBase, uint32(v10)+280)) = v68
												*(*int64)(unsafe.Add(mBase, uint32(v10)+272)) = v68
												*(*int64)(unsafe.Add(mBase, uint32(v10)+248)) = v68
												*(*int64)(unsafe.Add(mBase, uint32(v10)+255)) = v68
												v80 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v10)+242)) = uint8(v80)
												v82 = *(*int32)(unsafe.Add(mBase, uint32(v47)+52))
												v87 = F_heap_modify_tuple(m, v54, v82, v64, v10+int32(272), v10+int32(240))
												mBase = m.M
												v88 = m.ExcPending
												if v88 != 0 {
													return
												} else {
													F_CatalogTupleUpdate(m, v47, v87+int32(4), v87)
													mBase = m.M
													v92 = m.ExcPending
													if v92 != 0 {
														return
													} else {
														if v1 == v14 {
															v127 = v87
															F_pfree(m, v127)
															mBase = m.M
															v130 = m.ExcPending
															if v130 != 0 {
																return
															} else {
																F_relation_close(m, v47, int32(0))
																mBase = m.M
																v133 = m.ExcPending
																if v133 != 0 {
																	return
																} else {
																	F_PopActiveSnapshot(m)
																	mBase = m.M
																	v135 = m.ExcPending
																	if v135 != 0 {
																		return
																	} else {
																		if v27 == int32(2) {
																			m.G0 = v10 + int32(304)
																			return
																		} else {
																			F_CommitTransactionCommand(m)
																			mBase = m.M
																			v137 = m.ExcPending
																			if v137 != 0 {
																				return
																			} else {
																				m.G0 = v10 + int32(304)
																				return
																			}
																		}
																	}
																}
															}
														} else {
															v96 = F_errstart(m, int32(19), int32(0))
															mBase = m.M
															v97 = m.ExcPending
															if v97 != 0 {
																return
															} else {
																if v96 == int32(0) {
																	v127 = v87
																	F_pfree(m, v127)
																	mBase = m.M
																	v130 = m.ExcPending
																	if v130 != 0 {
																		return
																	} else {
																		F_relation_close(m, v47, int32(0))
																		mBase = m.M
																		v133 = m.ExcPending
																		if v133 != 0 {
																			return
																		} else {
																			F_PopActiveSnapshot(m)
																			mBase = m.M
																			v135 = m.ExcPending
																			if v135 != 0 {
																				return
																			} else {
																				if v27 == int32(2) {
																					m.G0 = v10 + int32(304)
																					return
																				} else {
																					F_CommitTransactionCommand(m)
																					mBase = m.M
																					v137 = m.ExcPending
																					if v137 != 0 {
																						return
																					} else {
																						m.G0 = v10 + int32(304)
																						return
																					}
																				}
																			}
																		}
																	}
																} else {
																	v101 = *(*int32)(unsafe.Add(mBase, _c_F_clear_subscription_skip_lsn[0]))
																	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+24))
																	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v102
																	F_errmsg(m, int32(_a_F_clear_subscription_skip_lsn_5), v10+int32(32))
																	mBase = m.M
																	v108 = m.ExcPending
																	if v108 != 0 {
																		return
																	} else {
																		*(*uint32)(unsafe.Add(mBase, uint32(v10)+28)) = uint32(v14)
																		v110 = int64(32)
																		v111 = int64(base.Ui64(v14) >> (uint(v110) % 64))
																		*(*uint32)(unsafe.Add(mBase, uint32(v10)+24)) = uint32(v111)
																		*(*uint32)(unsafe.Add(mBase, uint32(v10)+20)) = uint32(v1)
																		v115 = int64(base.Ui64(v1) >> (uint(v110) % 64))
																		*(*uint32)(unsafe.Add(mBase, uint32(v10)+16)) = uint32(v115)
																		v120 = F_errdetail(m, int32(_a_F_clear_subscription_skip_lsn_6), v10+int32(16))
																		mBase = m.M
																		v121 = m.ExcPending
																		if v121 != 0 {
																			return
																		} else {
																			F_errfinish(m, int32(_a_F_clear_subscription_skip_lsn_2), int32(_a_F_clear_subscription_skip_lsn_7), int32(_a_F_clear_subscription_skip_lsn_4))
																			mBase = m.M
																			v126 = m.ExcPending
																			if v126 != 0 {
																				return
																			} else {
																				v127 = v87
																				F_pfree(m, v127)
																				mBase = m.M
																				v130 = m.ExcPending
																				if v130 != 0 {
																					return
																				} else {
																					F_relation_close(m, v47, int32(0))
																					mBase = m.M
																					v133 = m.ExcPending
																					if v133 != 0 {
																						return
																					} else {
																						F_PopActiveSnapshot(m)
																						mBase = m.M
																						v135 = m.ExcPending
																						if v135 != 0 {
																							return
																						} else {
																							if v27 == int32(2) {
																								m.G0 = v10 + int32(304)
																								return
																							} else {
																								F_CommitTransactionCommand(m)
																								mBase = m.M
																								v137 = m.ExcPending
																								if v137 != 0 {
																									return
																								} else {
																									m.G0 = v10 + int32(304)
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
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, _c_F_clear_subscription_skip_lsn[2]))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
			v29 = base.B2i32(v27 == int32(2))
			if v29 == int32(0) {
				F_StartTransactionCommand(m)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					v34 = F_GetTransactionSnapshot(m)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						F_PushActiveSnapshot(m, v34)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							v40 = *(*int32)(unsafe.Add(mBase, _c_F_clear_subscription_skip_lsn[0]))
							v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
							F_LockSharedObject(m, int32(_a_F_clear_subscription_skip_lsn_0), v41, int32(1))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return
							} else {
								v47 = F_table_open(m, int32(_a_F_clear_subscription_skip_lsn_0), int32(3))
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return
								} else {
									v51 = *(*int32)(unsafe.Add(mBase, _c_F_clear_subscription_skip_lsn[0]))
									v52 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v51)+4)))
									v54 = F_SearchSysCacheCopy(m, int32(67), v52, int64(0))
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return
									} else {
										if v54 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v148 = m.ExcPending
											if v148 != 0 {
												return
											} else {
												v150 = *(*int32)(unsafe.Add(mBase, _c_F_clear_subscription_skip_lsn[0]))
												v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+24))
												*(*int32)(unsafe.Add(mBase, uint32(v10))) = v151
												F_errmsg_internal(m, int32(_a_F_clear_subscription_skip_lsn_1), v10)
												mBase = m.M
												v155 = m.ExcPending
												if v155 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_clear_subscription_skip_lsn_2), int32(_a_F_clear_subscription_skip_lsn_3), int32(_a_F_clear_subscription_skip_lsn_4))
													mBase = m.M
													v160 = m.ExcPending
													if v160 != 0 {
														return
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										} else {
											v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
											v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+22)))
											v61 = *(*int64)(unsafe.Add(mBase, uint32(v58+v59)+8))
											if v61 != v14 {
												v127 = v54
												F_pfree(m, v127)
												mBase = m.M
												v130 = m.ExcPending
												if v130 != 0 {
													return
												} else {
													F_relation_close(m, v47, int32(0))
													mBase = m.M
													v133 = m.ExcPending
													if v133 != 0 {
														return
													} else {
														F_PopActiveSnapshot(m)
														mBase = m.M
														v135 = m.ExcPending
														if v135 != 0 {
															return
														} else {
															if v27 == int32(2) {
																m.G0 = v10 + int32(304)
																return
															} else {
																F_CommitTransactionCommand(m)
																mBase = m.M
																v137 = m.ExcPending
																if v137 != 0 {
																	return
																} else {
																	m.G0 = v10 + int32(304)
																	return
																}
															}
														}
													}
												}
											} else {
												v64 = v10 + int32(48)
												base.MemoryFill(m, v64, int32(0), int32(184))
												v68 = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(v10)+240)) = v68
												*(*int64)(unsafe.Add(mBase, uint32(v10)+287)) = v68
												*(*int64)(unsafe.Add(mBase, uint32(v10)+280)) = v68
												*(*int64)(unsafe.Add(mBase, uint32(v10)+272)) = v68
												*(*int64)(unsafe.Add(mBase, uint32(v10)+248)) = v68
												*(*int64)(unsafe.Add(mBase, uint32(v10)+255)) = v68
												v80 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v10)+242)) = uint8(v80)
												v82 = *(*int32)(unsafe.Add(mBase, uint32(v47)+52))
												v87 = F_heap_modify_tuple(m, v54, v82, v64, v10+int32(272), v10+int32(240))
												mBase = m.M
												v88 = m.ExcPending
												if v88 != 0 {
													return
												} else {
													F_CatalogTupleUpdate(m, v47, v87+int32(4), v87)
													mBase = m.M
													v92 = m.ExcPending
													if v92 != 0 {
														return
													} else {
														if v1 == v14 {
															v127 = v87
															F_pfree(m, v127)
															mBase = m.M
															v130 = m.ExcPending
															if v130 != 0 {
																return
															} else {
																F_relation_close(m, v47, int32(0))
																mBase = m.M
																v133 = m.ExcPending
																if v133 != 0 {
																	return
																} else {
																	F_PopActiveSnapshot(m)
																	mBase = m.M
																	v135 = m.ExcPending
																	if v135 != 0 {
																		return
																	} else {
																		if v27 == int32(2) {
																			m.G0 = v10 + int32(304)
																			return
																		} else {
																			F_CommitTransactionCommand(m)
																			mBase = m.M
																			v137 = m.ExcPending
																			if v137 != 0 {
																				return
																			} else {
																				m.G0 = v10 + int32(304)
																				return
																			}
																		}
																	}
																}
															}
														} else {
															v96 = F_errstart(m, int32(19), int32(0))
															mBase = m.M
															v97 = m.ExcPending
															if v97 != 0 {
																return
															} else {
																if v96 == int32(0) {
																	v127 = v87
																	F_pfree(m, v127)
																	mBase = m.M
																	v130 = m.ExcPending
																	if v130 != 0 {
																		return
																	} else {
																		F_relation_close(m, v47, int32(0))
																		mBase = m.M
																		v133 = m.ExcPending
																		if v133 != 0 {
																			return
																		} else {
																			F_PopActiveSnapshot(m)
																			mBase = m.M
																			v135 = m.ExcPending
																			if v135 != 0 {
																				return
																			} else {
																				if v27 == int32(2) {
																					m.G0 = v10 + int32(304)
																					return
																				} else {
																					F_CommitTransactionCommand(m)
																					mBase = m.M
																					v137 = m.ExcPending
																					if v137 != 0 {
																						return
																					} else {
																						m.G0 = v10 + int32(304)
																						return
																					}
																				}
																			}
																		}
																	}
																} else {
																	v101 = *(*int32)(unsafe.Add(mBase, _c_F_clear_subscription_skip_lsn[0]))
																	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+24))
																	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v102
																	F_errmsg(m, int32(_a_F_clear_subscription_skip_lsn_5), v10+int32(32))
																	mBase = m.M
																	v108 = m.ExcPending
																	if v108 != 0 {
																		return
																	} else {
																		*(*uint32)(unsafe.Add(mBase, uint32(v10)+28)) = uint32(v14)
																		v110 = int64(32)
																		v111 = int64(base.Ui64(v14) >> (uint(v110) % 64))
																		*(*uint32)(unsafe.Add(mBase, uint32(v10)+24)) = uint32(v111)
																		*(*uint32)(unsafe.Add(mBase, uint32(v10)+20)) = uint32(v1)
																		v115 = int64(base.Ui64(v1) >> (uint(v110) % 64))
																		*(*uint32)(unsafe.Add(mBase, uint32(v10)+16)) = uint32(v115)
																		v120 = F_errdetail(m, int32(_a_F_clear_subscription_skip_lsn_6), v10+int32(16))
																		mBase = m.M
																		v121 = m.ExcPending
																		if v121 != 0 {
																			return
																		} else {
																			F_errfinish(m, int32(_a_F_clear_subscription_skip_lsn_2), int32(_a_F_clear_subscription_skip_lsn_7), int32(_a_F_clear_subscription_skip_lsn_4))
																			mBase = m.M
																			v126 = m.ExcPending
																			if v126 != 0 {
																				return
																			} else {
																				v127 = v87
																				F_pfree(m, v127)
																				mBase = m.M
																				v130 = m.ExcPending
																				if v130 != 0 {
																					return
																				} else {
																					F_relation_close(m, v47, int32(0))
																					mBase = m.M
																					v133 = m.ExcPending
																					if v133 != 0 {
																						return
																					} else {
																						F_PopActiveSnapshot(m)
																						mBase = m.M
																						v135 = m.ExcPending
																						if v135 != 0 {
																							return
																						} else {
																							if v27 == int32(2) {
																								m.G0 = v10 + int32(304)
																								return
																							} else {
																								F_CommitTransactionCommand(m)
																								mBase = m.M
																								v137 = m.ExcPending
																								if v137 != 0 {
																									return
																								} else {
																									m.G0 = v10 + int32(304)
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
			} else {
				v34 = F_GetTransactionSnapshot(m)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					F_PushActiveSnapshot(m, v34)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, _c_F_clear_subscription_skip_lsn[0]))
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
						F_LockSharedObject(m, int32(_a_F_clear_subscription_skip_lsn_0), v41, int32(1))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return
						} else {
							v47 = F_table_open(m, int32(_a_F_clear_subscription_skip_lsn_0), int32(3))
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return
							} else {
								v51 = *(*int32)(unsafe.Add(mBase, _c_F_clear_subscription_skip_lsn[0]))
								v52 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v51)+4)))
								v54 = F_SearchSysCacheCopy(m, int32(67), v52, int64(0))
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return
								} else {
									if v54 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v148 = m.ExcPending
										if v148 != 0 {
											return
										} else {
											v150 = *(*int32)(unsafe.Add(mBase, _c_F_clear_subscription_skip_lsn[0]))
											v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+24))
											*(*int32)(unsafe.Add(mBase, uint32(v10))) = v151
											F_errmsg_internal(m, int32(_a_F_clear_subscription_skip_lsn_1), v10)
											mBase = m.M
											v155 = m.ExcPending
											if v155 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_clear_subscription_skip_lsn_2), int32(_a_F_clear_subscription_skip_lsn_3), int32(_a_F_clear_subscription_skip_lsn_4))
												mBase = m.M
												v160 = m.ExcPending
												if v160 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
										v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+22)))
										v61 = *(*int64)(unsafe.Add(mBase, uint32(v58+v59)+8))
										if v61 != v14 {
											v127 = v54
											F_pfree(m, v127)
											mBase = m.M
											v130 = m.ExcPending
											if v130 != 0 {
												return
											} else {
												F_relation_close(m, v47, int32(0))
												mBase = m.M
												v133 = m.ExcPending
												if v133 != 0 {
													return
												} else {
													F_PopActiveSnapshot(m)
													mBase = m.M
													v135 = m.ExcPending
													if v135 != 0 {
														return
													} else {
														if v27 == int32(2) {
															m.G0 = v10 + int32(304)
															return
														} else {
															F_CommitTransactionCommand(m)
															mBase = m.M
															v137 = m.ExcPending
															if v137 != 0 {
																return
															} else {
																m.G0 = v10 + int32(304)
																return
															}
														}
													}
												}
											}
										} else {
											v64 = v10 + int32(48)
											base.MemoryFill(m, v64, int32(0), int32(184))
											v68 = int64(0)
											*(*int64)(unsafe.Add(mBase, uint32(v10)+240)) = v68
											*(*int64)(unsafe.Add(mBase, uint32(v10)+287)) = v68
											*(*int64)(unsafe.Add(mBase, uint32(v10)+280)) = v68
											*(*int64)(unsafe.Add(mBase, uint32(v10)+272)) = v68
											*(*int64)(unsafe.Add(mBase, uint32(v10)+248)) = v68
											*(*int64)(unsafe.Add(mBase, uint32(v10)+255)) = v68
											v80 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v10)+242)) = uint8(v80)
											v82 = *(*int32)(unsafe.Add(mBase, uint32(v47)+52))
											v87 = F_heap_modify_tuple(m, v54, v82, v64, v10+int32(272), v10+int32(240))
											mBase = m.M
											v88 = m.ExcPending
											if v88 != 0 {
												return
											} else {
												F_CatalogTupleUpdate(m, v47, v87+int32(4), v87)
												mBase = m.M
												v92 = m.ExcPending
												if v92 != 0 {
													return
												} else {
													if v1 == v14 {
														v127 = v87
														F_pfree(m, v127)
														mBase = m.M
														v130 = m.ExcPending
														if v130 != 0 {
															return
														} else {
															F_relation_close(m, v47, int32(0))
															mBase = m.M
															v133 = m.ExcPending
															if v133 != 0 {
																return
															} else {
																F_PopActiveSnapshot(m)
																mBase = m.M
																v135 = m.ExcPending
																if v135 != 0 {
																	return
																} else {
																	if v27 == int32(2) {
																		m.G0 = v10 + int32(304)
																		return
																	} else {
																		F_CommitTransactionCommand(m)
																		mBase = m.M
																		v137 = m.ExcPending
																		if v137 != 0 {
																			return
																		} else {
																			m.G0 = v10 + int32(304)
																			return
																		}
																	}
																}
															}
														}
													} else {
														v96 = F_errstart(m, int32(19), int32(0))
														mBase = m.M
														v97 = m.ExcPending
														if v97 != 0 {
															return
														} else {
															if v96 == int32(0) {
																v127 = v87
																F_pfree(m, v127)
																mBase = m.M
																v130 = m.ExcPending
																if v130 != 0 {
																	return
																} else {
																	F_relation_close(m, v47, int32(0))
																	mBase = m.M
																	v133 = m.ExcPending
																	if v133 != 0 {
																		return
																	} else {
																		F_PopActiveSnapshot(m)
																		mBase = m.M
																		v135 = m.ExcPending
																		if v135 != 0 {
																			return
																		} else {
																			if v27 == int32(2) {
																				m.G0 = v10 + int32(304)
																				return
																			} else {
																				F_CommitTransactionCommand(m)
																				mBase = m.M
																				v137 = m.ExcPending
																				if v137 != 0 {
																					return
																				} else {
																					m.G0 = v10 + int32(304)
																					return
																				}
																			}
																		}
																	}
																}
															} else {
																v101 = *(*int32)(unsafe.Add(mBase, _c_F_clear_subscription_skip_lsn[0]))
																v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+24))
																*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v102
																F_errmsg(m, int32(_a_F_clear_subscription_skip_lsn_5), v10+int32(32))
																mBase = m.M
																v108 = m.ExcPending
																if v108 != 0 {
																	return
																} else {
																	*(*uint32)(unsafe.Add(mBase, uint32(v10)+28)) = uint32(v14)
																	v110 = int64(32)
																	v111 = int64(base.Ui64(v14) >> (uint(v110) % 64))
																	*(*uint32)(unsafe.Add(mBase, uint32(v10)+24)) = uint32(v111)
																	*(*uint32)(unsafe.Add(mBase, uint32(v10)+20)) = uint32(v1)
																	v115 = int64(base.Ui64(v1) >> (uint(v110) % 64))
																	*(*uint32)(unsafe.Add(mBase, uint32(v10)+16)) = uint32(v115)
																	v120 = F_errdetail(m, int32(_a_F_clear_subscription_skip_lsn_6), v10+int32(16))
																	mBase = m.M
																	v121 = m.ExcPending
																	if v121 != 0 {
																		return
																	} else {
																		F_errfinish(m, int32(_a_F_clear_subscription_skip_lsn_2), int32(_a_F_clear_subscription_skip_lsn_7), int32(_a_F_clear_subscription_skip_lsn_4))
																		mBase = m.M
																		v126 = m.ExcPending
																		if v126 != 0 {
																			return
																		} else {
																			v127 = v87
																			F_pfree(m, v127)
																			mBase = m.M
																			v130 = m.ExcPending
																			if v130 != 0 {
																				return
																			} else {
																				F_relation_close(m, v47, int32(0))
																				mBase = m.M
																				v133 = m.ExcPending
																				if v133 != 0 {
																					return
																				} else {
																					F_PopActiveSnapshot(m)
																					mBase = m.M
																					v135 = m.ExcPending
																					if v135 != 0 {
																						return
																					} else {
																						if v27 == int32(2) {
																							m.G0 = v10 + int32(304)
																							return
																						} else {
																							F_CommitTransactionCommand(m)
																							mBase = m.M
																							v137 = m.ExcPending
																							if v137 != 0 {
																								return
																							} else {
																								m.G0 = v10 + int32(304)
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
