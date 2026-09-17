package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_clear_subscription_skip_lsn(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v1 int64
	_ = v1
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int64
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int64
	_ = v71
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
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v115 int64
	_ = v115
	var v116 int64
	_ = v116
	var v120 int64
	_ = v120
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	v1 = l0
	v9 = m.G0
	v11 = v9 - int32(192)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_clear_subscription_skip_lsn[0]))
	v15 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
	if v15 == int64(0) {
		m.G0 = v11 + int32(192)
		return
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, _c_F_clear_subscription_skip_lsn[1]))
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+16)))
		if v20 == int32(1) {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			if v23 == int32(3) {
				m.G0 = v11 + int32(192)
				return
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, _c_F_clear_subscription_skip_lsn[2]))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
				v30 = base.B2i32(v28 == int32(2))
				if v30 == int32(0) {
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
								v41 = *(*int32)(unsafe.Add(mBase, _c_F_clear_subscription_skip_lsn[0]))
								v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
								F_LockSharedObject(m, int32(_a_F_clear_subscription_skip_lsn_0), v42, int32(1))
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return
								} else {
									v48 = F_table_open(m, int32(_a_F_clear_subscription_skip_lsn_0), int32(3))
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return
									} else {
										v52 = *(*int32)(unsafe.Add(mBase, _c_F_clear_subscription_skip_lsn[0]))
										v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
										v55 = F_SearchSysCacheCopy(m, int32(67), v53, int32(0))
										mBase = m.M
										v56 = m.ExcPending
										if v56 != 0 {
											return
										} else {
											if v55 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v155 = m.ExcPending
												if v155 != 0 {
													return
												} else {
													v157 = *(*int32)(unsafe.Add(mBase, _c_F_clear_subscription_skip_lsn[0]))
													v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+16))
													*(*int32)(unsafe.Add(mBase, uint32(v11))) = v158
													F_errmsg_internal(m, int32(_a_F_clear_subscription_skip_lsn_1), v11)
													mBase = m.M
													v162 = m.ExcPending
													if v162 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_clear_subscription_skip_lsn_2), int32(_a_F_clear_subscription_skip_lsn_3), int32(_a_F_clear_subscription_skip_lsn_4))
														mBase = m.M
														v167 = m.ExcPending
														if v167 != 0 {
															return
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
												v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+22)))
												v62 = *(*int64)(unsafe.Add(mBase, uint32(v59+v60)+8))
												if v62 != v15 {
													v132 = v55
													F_pfree(m, v132)
													mBase = m.M
													v136 = m.ExcPending
													if v136 != 0 {
														return
													} else {
														F_relation_close(m, v48, int32(0))
														mBase = m.M
														v139 = m.ExcPending
														if v139 != 0 {
															return
														} else {
															F_PopActiveSnapshot(m)
															mBase = m.M
															v141 = m.ExcPending
															if v141 != 0 {
																return
															} else {
																if v28 == int32(2) {
																	m.G0 = v11 + int32(192)
																	return
																} else {
																	F_CommitTransactionCommand(m)
																	mBase = m.M
																	v143 = m.ExcPending
																	if v143 != 0 {
																		return
																	} else {
																		m.G0 = v11 + int32(192)
																		return
																	}
																}
															}
														}
													}
												} else {
													v65 = v11 + int32(48)
													v66 = int32(0)
													base.MemoryFill(m, v65, v66, int32(72))
													*(*uint16)(unsafe.Add(mBase, uint32(v11)+176)) = uint16(v66)
													v71 = int64(0)
													*(*int64)(unsafe.Add(mBase, uint32(v11)+168)) = v71
													*(*int64)(unsafe.Add(mBase, uint32(v11)+160)) = v71
													*(*int64)(unsafe.Add(mBase, uint32(v11)+128)) = v71
													*(*int64)(unsafe.Add(mBase, uint32(v11)+136)) = v71
													*(*uint16)(unsafe.Add(mBase, uint32(v11)+144)) = uint16(v66)
													v82 = F_Int64GetDatum(m, v71)
													mBase = m.M
													v83 = m.ExcPending
													if v83 != 0 {
														return
													} else {
														v84 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+130)) = uint8(v84)
														*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v82
														v87 = *(*int32)(unsafe.Add(mBase, uint32(v48)+52))
														v92 = F_heap_modify_tuple(m, v55, v87, v65, v11+int32(160), v11+int32(128))
														mBase = m.M
														v93 = m.ExcPending
														if v93 != 0 {
															return
														} else {
															F_CatalogTupleUpdate(m, v48, v92+int32(4), v92)
															mBase = m.M
															v97 = m.ExcPending
															if v97 != 0 {
																return
															} else {
																if v1 == v15 {
																	v132 = v92
																	F_pfree(m, v132)
																	mBase = m.M
																	v136 = m.ExcPending
																	if v136 != 0 {
																		return
																	} else {
																		F_relation_close(m, v48, int32(0))
																		mBase = m.M
																		v139 = m.ExcPending
																		if v139 != 0 {
																			return
																		} else {
																			F_PopActiveSnapshot(m)
																			mBase = m.M
																			v141 = m.ExcPending
																			if v141 != 0 {
																				return
																			} else {
																				if v28 == int32(2) {
																					m.G0 = v11 + int32(192)
																					return
																				} else {
																					F_CommitTransactionCommand(m)
																					mBase = m.M
																					v143 = m.ExcPending
																					if v143 != 0 {
																						return
																					} else {
																						m.G0 = v11 + int32(192)
																						return
																					}
																				}
																			}
																		}
																	}
																} else {
																	v101 = F_errstart(m, int32(19), int32(0))
																	mBase = m.M
																	v102 = m.ExcPending
																	if v102 != 0 {
																		return
																	} else {
																		if v101 == int32(0) {
																			v132 = v92
																			F_pfree(m, v132)
																			mBase = m.M
																			v136 = m.ExcPending
																			if v136 != 0 {
																				return
																			} else {
																				F_relation_close(m, v48, int32(0))
																				mBase = m.M
																				v139 = m.ExcPending
																				if v139 != 0 {
																					return
																				} else {
																					F_PopActiveSnapshot(m)
																					mBase = m.M
																					v141 = m.ExcPending
																					if v141 != 0 {
																						return
																					} else {
																						if v28 == int32(2) {
																							m.G0 = v11 + int32(192)
																							return
																						} else {
																							F_CommitTransactionCommand(m)
																							mBase = m.M
																							v143 = m.ExcPending
																							if v143 != 0 {
																								return
																							} else {
																								m.G0 = v11 + int32(192)
																								return
																							}
																						}
																					}
																				}
																			}
																		} else {
																			v106 = *(*int32)(unsafe.Add(mBase, _c_F_clear_subscription_skip_lsn[0]))
																			v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+16))
																			*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v107
																			F_errmsg(m, int32(_a_F_clear_subscription_skip_lsn_5), v11+int32(32))
																			mBase = m.M
																			v113 = m.ExcPending
																			if v113 != 0 {
																				return
																			} else {
																				*(*uint32)(unsafe.Add(mBase, uint32(v11)+28)) = uint32(v15)
																				v115 = int64(32)
																				v116 = int64(base.Ui64(v15) >> (uint(v115) % 64))
																				*(*uint32)(unsafe.Add(mBase, uint32(v11)+24)) = uint32(v116)
																				*(*uint32)(unsafe.Add(mBase, uint32(v11)+20)) = uint32(v1)
																				v120 = int64(base.Ui64(v1) >> (uint(v115) % 64))
																				*(*uint32)(unsafe.Add(mBase, uint32(v11)+16)) = uint32(v120)
																				F_errdetail(m, int32(_a_F_clear_subscription_skip_lsn_6), v11+int32(16))
																				mBase = m.M
																				v126 = m.ExcPending
																				if v126 != 0 {
																					return
																				} else {
																					F_errfinish(m, int32(_a_F_clear_subscription_skip_lsn_2), int32(_a_F_clear_subscription_skip_lsn_7), int32(_a_F_clear_subscription_skip_lsn_4))
																					mBase = m.M
																					v131 = m.ExcPending
																					if v131 != 0 {
																						return
																					} else {
																						v132 = v92
																						F_pfree(m, v132)
																						mBase = m.M
																						v136 = m.ExcPending
																						if v136 != 0 {
																							return
																						} else {
																							F_relation_close(m, v48, int32(0))
																							mBase = m.M
																							v139 = m.ExcPending
																							if v139 != 0 {
																								return
																							} else {
																								F_PopActiveSnapshot(m)
																								mBase = m.M
																								v141 = m.ExcPending
																								if v141 != 0 {
																									return
																								} else {
																									if v28 == int32(2) {
																										m.G0 = v11 + int32(192)
																										return
																									} else {
																										F_CommitTransactionCommand(m)
																										mBase = m.M
																										v143 = m.ExcPending
																										if v143 != 0 {
																											return
																										} else {
																											m.G0 = v11 + int32(192)
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
							v41 = *(*int32)(unsafe.Add(mBase, _c_F_clear_subscription_skip_lsn[0]))
							v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
							F_LockSharedObject(m, int32(_a_F_clear_subscription_skip_lsn_0), v42, int32(1))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
							} else {
								v48 = F_table_open(m, int32(_a_F_clear_subscription_skip_lsn_0), int32(3))
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return
								} else {
									v52 = *(*int32)(unsafe.Add(mBase, _c_F_clear_subscription_skip_lsn[0]))
									v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
									v55 = F_SearchSysCacheCopy(m, int32(67), v53, int32(0))
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return
									} else {
										if v55 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v155 = m.ExcPending
											if v155 != 0 {
												return
											} else {
												v157 = *(*int32)(unsafe.Add(mBase, _c_F_clear_subscription_skip_lsn[0]))
												v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+16))
												*(*int32)(unsafe.Add(mBase, uint32(v11))) = v158
												F_errmsg_internal(m, int32(_a_F_clear_subscription_skip_lsn_1), v11)
												mBase = m.M
												v162 = m.ExcPending
												if v162 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_clear_subscription_skip_lsn_2), int32(_a_F_clear_subscription_skip_lsn_3), int32(_a_F_clear_subscription_skip_lsn_4))
													mBase = m.M
													v167 = m.ExcPending
													if v167 != 0 {
														return
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										} else {
											v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
											v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+22)))
											v62 = *(*int64)(unsafe.Add(mBase, uint32(v59+v60)+8))
											if v62 != v15 {
												v132 = v55
												F_pfree(m, v132)
												mBase = m.M
												v136 = m.ExcPending
												if v136 != 0 {
													return
												} else {
													F_relation_close(m, v48, int32(0))
													mBase = m.M
													v139 = m.ExcPending
													if v139 != 0 {
														return
													} else {
														F_PopActiveSnapshot(m)
														mBase = m.M
														v141 = m.ExcPending
														if v141 != 0 {
															return
														} else {
															if v28 == int32(2) {
																m.G0 = v11 + int32(192)
																return
															} else {
																F_CommitTransactionCommand(m)
																mBase = m.M
																v143 = m.ExcPending
																if v143 != 0 {
																	return
																} else {
																	m.G0 = v11 + int32(192)
																	return
																}
															}
														}
													}
												}
											} else {
												v65 = v11 + int32(48)
												v66 = int32(0)
												base.MemoryFill(m, v65, v66, int32(72))
												*(*uint16)(unsafe.Add(mBase, uint32(v11)+176)) = uint16(v66)
												v71 = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(v11)+168)) = v71
												*(*int64)(unsafe.Add(mBase, uint32(v11)+160)) = v71
												*(*int64)(unsafe.Add(mBase, uint32(v11)+128)) = v71
												*(*int64)(unsafe.Add(mBase, uint32(v11)+136)) = v71
												*(*uint16)(unsafe.Add(mBase, uint32(v11)+144)) = uint16(v66)
												v82 = F_Int64GetDatum(m, v71)
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
													return
												} else {
													v84 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+130)) = uint8(v84)
													*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v82
													v87 = *(*int32)(unsafe.Add(mBase, uint32(v48)+52))
													v92 = F_heap_modify_tuple(m, v55, v87, v65, v11+int32(160), v11+int32(128))
													mBase = m.M
													v93 = m.ExcPending
													if v93 != 0 {
														return
													} else {
														F_CatalogTupleUpdate(m, v48, v92+int32(4), v92)
														mBase = m.M
														v97 = m.ExcPending
														if v97 != 0 {
															return
														} else {
															if v1 == v15 {
																v132 = v92
																F_pfree(m, v132)
																mBase = m.M
																v136 = m.ExcPending
																if v136 != 0 {
																	return
																} else {
																	F_relation_close(m, v48, int32(0))
																	mBase = m.M
																	v139 = m.ExcPending
																	if v139 != 0 {
																		return
																	} else {
																		F_PopActiveSnapshot(m)
																		mBase = m.M
																		v141 = m.ExcPending
																		if v141 != 0 {
																			return
																		} else {
																			if v28 == int32(2) {
																				m.G0 = v11 + int32(192)
																				return
																			} else {
																				F_CommitTransactionCommand(m)
																				mBase = m.M
																				v143 = m.ExcPending
																				if v143 != 0 {
																					return
																				} else {
																					m.G0 = v11 + int32(192)
																					return
																				}
																			}
																		}
																	}
																}
															} else {
																v101 = F_errstart(m, int32(19), int32(0))
																mBase = m.M
																v102 = m.ExcPending
																if v102 != 0 {
																	return
																} else {
																	if v101 == int32(0) {
																		v132 = v92
																		F_pfree(m, v132)
																		mBase = m.M
																		v136 = m.ExcPending
																		if v136 != 0 {
																			return
																		} else {
																			F_relation_close(m, v48, int32(0))
																			mBase = m.M
																			v139 = m.ExcPending
																			if v139 != 0 {
																				return
																			} else {
																				F_PopActiveSnapshot(m)
																				mBase = m.M
																				v141 = m.ExcPending
																				if v141 != 0 {
																					return
																				} else {
																					if v28 == int32(2) {
																						m.G0 = v11 + int32(192)
																						return
																					} else {
																						F_CommitTransactionCommand(m)
																						mBase = m.M
																						v143 = m.ExcPending
																						if v143 != 0 {
																							return
																						} else {
																							m.G0 = v11 + int32(192)
																							return
																						}
																					}
																				}
																			}
																		}
																	} else {
																		v106 = *(*int32)(unsafe.Add(mBase, _c_F_clear_subscription_skip_lsn[0]))
																		v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+16))
																		*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v107
																		F_errmsg(m, int32(_a_F_clear_subscription_skip_lsn_5), v11+int32(32))
																		mBase = m.M
																		v113 = m.ExcPending
																		if v113 != 0 {
																			return
																		} else {
																			*(*uint32)(unsafe.Add(mBase, uint32(v11)+28)) = uint32(v15)
																			v115 = int64(32)
																			v116 = int64(base.Ui64(v15) >> (uint(v115) % 64))
																			*(*uint32)(unsafe.Add(mBase, uint32(v11)+24)) = uint32(v116)
																			*(*uint32)(unsafe.Add(mBase, uint32(v11)+20)) = uint32(v1)
																			v120 = int64(base.Ui64(v1) >> (uint(v115) % 64))
																			*(*uint32)(unsafe.Add(mBase, uint32(v11)+16)) = uint32(v120)
																			F_errdetail(m, int32(_a_F_clear_subscription_skip_lsn_6), v11+int32(16))
																			mBase = m.M
																			v126 = m.ExcPending
																			if v126 != 0 {
																				return
																			} else {
																				F_errfinish(m, int32(_a_F_clear_subscription_skip_lsn_2), int32(_a_F_clear_subscription_skip_lsn_7), int32(_a_F_clear_subscription_skip_lsn_4))
																				mBase = m.M
																				v131 = m.ExcPending
																				if v131 != 0 {
																					return
																				} else {
																					v132 = v92
																					F_pfree(m, v132)
																					mBase = m.M
																					v136 = m.ExcPending
																					if v136 != 0 {
																						return
																					} else {
																						F_relation_close(m, v48, int32(0))
																						mBase = m.M
																						v139 = m.ExcPending
																						if v139 != 0 {
																							return
																						} else {
																							F_PopActiveSnapshot(m)
																							mBase = m.M
																							v141 = m.ExcPending
																							if v141 != 0 {
																								return
																							} else {
																								if v28 == int32(2) {
																									m.G0 = v11 + int32(192)
																									return
																								} else {
																									F_CommitTransactionCommand(m)
																									mBase = m.M
																									v143 = m.ExcPending
																									if v143 != 0 {
																										return
																									} else {
																										m.G0 = v11 + int32(192)
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
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, _c_F_clear_subscription_skip_lsn[2]))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
			v30 = base.B2i32(v28 == int32(2))
			if v30 == int32(0) {
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
							v41 = *(*int32)(unsafe.Add(mBase, _c_F_clear_subscription_skip_lsn[0]))
							v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
							F_LockSharedObject(m, int32(_a_F_clear_subscription_skip_lsn_0), v42, int32(1))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
							} else {
								v48 = F_table_open(m, int32(_a_F_clear_subscription_skip_lsn_0), int32(3))
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return
								} else {
									v52 = *(*int32)(unsafe.Add(mBase, _c_F_clear_subscription_skip_lsn[0]))
									v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
									v55 = F_SearchSysCacheCopy(m, int32(67), v53, int32(0))
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return
									} else {
										if v55 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v155 = m.ExcPending
											if v155 != 0 {
												return
											} else {
												v157 = *(*int32)(unsafe.Add(mBase, _c_F_clear_subscription_skip_lsn[0]))
												v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+16))
												*(*int32)(unsafe.Add(mBase, uint32(v11))) = v158
												F_errmsg_internal(m, int32(_a_F_clear_subscription_skip_lsn_1), v11)
												mBase = m.M
												v162 = m.ExcPending
												if v162 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_clear_subscription_skip_lsn_2), int32(_a_F_clear_subscription_skip_lsn_3), int32(_a_F_clear_subscription_skip_lsn_4))
													mBase = m.M
													v167 = m.ExcPending
													if v167 != 0 {
														return
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										} else {
											v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
											v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+22)))
											v62 = *(*int64)(unsafe.Add(mBase, uint32(v59+v60)+8))
											if v62 != v15 {
												v132 = v55
												F_pfree(m, v132)
												mBase = m.M
												v136 = m.ExcPending
												if v136 != 0 {
													return
												} else {
													F_relation_close(m, v48, int32(0))
													mBase = m.M
													v139 = m.ExcPending
													if v139 != 0 {
														return
													} else {
														F_PopActiveSnapshot(m)
														mBase = m.M
														v141 = m.ExcPending
														if v141 != 0 {
															return
														} else {
															if v28 == int32(2) {
																m.G0 = v11 + int32(192)
																return
															} else {
																F_CommitTransactionCommand(m)
																mBase = m.M
																v143 = m.ExcPending
																if v143 != 0 {
																	return
																} else {
																	m.G0 = v11 + int32(192)
																	return
																}
															}
														}
													}
												}
											} else {
												v65 = v11 + int32(48)
												v66 = int32(0)
												base.MemoryFill(m, v65, v66, int32(72))
												*(*uint16)(unsafe.Add(mBase, uint32(v11)+176)) = uint16(v66)
												v71 = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(v11)+168)) = v71
												*(*int64)(unsafe.Add(mBase, uint32(v11)+160)) = v71
												*(*int64)(unsafe.Add(mBase, uint32(v11)+128)) = v71
												*(*int64)(unsafe.Add(mBase, uint32(v11)+136)) = v71
												*(*uint16)(unsafe.Add(mBase, uint32(v11)+144)) = uint16(v66)
												v82 = F_Int64GetDatum(m, v71)
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
													return
												} else {
													v84 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+130)) = uint8(v84)
													*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v82
													v87 = *(*int32)(unsafe.Add(mBase, uint32(v48)+52))
													v92 = F_heap_modify_tuple(m, v55, v87, v65, v11+int32(160), v11+int32(128))
													mBase = m.M
													v93 = m.ExcPending
													if v93 != 0 {
														return
													} else {
														F_CatalogTupleUpdate(m, v48, v92+int32(4), v92)
														mBase = m.M
														v97 = m.ExcPending
														if v97 != 0 {
															return
														} else {
															if v1 == v15 {
																v132 = v92
																F_pfree(m, v132)
																mBase = m.M
																v136 = m.ExcPending
																if v136 != 0 {
																	return
																} else {
																	F_relation_close(m, v48, int32(0))
																	mBase = m.M
																	v139 = m.ExcPending
																	if v139 != 0 {
																		return
																	} else {
																		F_PopActiveSnapshot(m)
																		mBase = m.M
																		v141 = m.ExcPending
																		if v141 != 0 {
																			return
																		} else {
																			if v28 == int32(2) {
																				m.G0 = v11 + int32(192)
																				return
																			} else {
																				F_CommitTransactionCommand(m)
																				mBase = m.M
																				v143 = m.ExcPending
																				if v143 != 0 {
																					return
																				} else {
																					m.G0 = v11 + int32(192)
																					return
																				}
																			}
																		}
																	}
																}
															} else {
																v101 = F_errstart(m, int32(19), int32(0))
																mBase = m.M
																v102 = m.ExcPending
																if v102 != 0 {
																	return
																} else {
																	if v101 == int32(0) {
																		v132 = v92
																		F_pfree(m, v132)
																		mBase = m.M
																		v136 = m.ExcPending
																		if v136 != 0 {
																			return
																		} else {
																			F_relation_close(m, v48, int32(0))
																			mBase = m.M
																			v139 = m.ExcPending
																			if v139 != 0 {
																				return
																			} else {
																				F_PopActiveSnapshot(m)
																				mBase = m.M
																				v141 = m.ExcPending
																				if v141 != 0 {
																					return
																				} else {
																					if v28 == int32(2) {
																						m.G0 = v11 + int32(192)
																						return
																					} else {
																						F_CommitTransactionCommand(m)
																						mBase = m.M
																						v143 = m.ExcPending
																						if v143 != 0 {
																							return
																						} else {
																							m.G0 = v11 + int32(192)
																							return
																						}
																					}
																				}
																			}
																		}
																	} else {
																		v106 = *(*int32)(unsafe.Add(mBase, _c_F_clear_subscription_skip_lsn[0]))
																		v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+16))
																		*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v107
																		F_errmsg(m, int32(_a_F_clear_subscription_skip_lsn_5), v11+int32(32))
																		mBase = m.M
																		v113 = m.ExcPending
																		if v113 != 0 {
																			return
																		} else {
																			*(*uint32)(unsafe.Add(mBase, uint32(v11)+28)) = uint32(v15)
																			v115 = int64(32)
																			v116 = int64(base.Ui64(v15) >> (uint(v115) % 64))
																			*(*uint32)(unsafe.Add(mBase, uint32(v11)+24)) = uint32(v116)
																			*(*uint32)(unsafe.Add(mBase, uint32(v11)+20)) = uint32(v1)
																			v120 = int64(base.Ui64(v1) >> (uint(v115) % 64))
																			*(*uint32)(unsafe.Add(mBase, uint32(v11)+16)) = uint32(v120)
																			F_errdetail(m, int32(_a_F_clear_subscription_skip_lsn_6), v11+int32(16))
																			mBase = m.M
																			v126 = m.ExcPending
																			if v126 != 0 {
																				return
																			} else {
																				F_errfinish(m, int32(_a_F_clear_subscription_skip_lsn_2), int32(_a_F_clear_subscription_skip_lsn_7), int32(_a_F_clear_subscription_skip_lsn_4))
																				mBase = m.M
																				v131 = m.ExcPending
																				if v131 != 0 {
																					return
																				} else {
																					v132 = v92
																					F_pfree(m, v132)
																					mBase = m.M
																					v136 = m.ExcPending
																					if v136 != 0 {
																						return
																					} else {
																						F_relation_close(m, v48, int32(0))
																						mBase = m.M
																						v139 = m.ExcPending
																						if v139 != 0 {
																							return
																						} else {
																							F_PopActiveSnapshot(m)
																							mBase = m.M
																							v141 = m.ExcPending
																							if v141 != 0 {
																								return
																							} else {
																								if v28 == int32(2) {
																									m.G0 = v11 + int32(192)
																									return
																								} else {
																									F_CommitTransactionCommand(m)
																									mBase = m.M
																									v143 = m.ExcPending
																									if v143 != 0 {
																										return
																									} else {
																										m.G0 = v11 + int32(192)
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
						v41 = *(*int32)(unsafe.Add(mBase, _c_F_clear_subscription_skip_lsn[0]))
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
						F_LockSharedObject(m, int32(_a_F_clear_subscription_skip_lsn_0), v42, int32(1))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							v48 = F_table_open(m, int32(_a_F_clear_subscription_skip_lsn_0), int32(3))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, _c_F_clear_subscription_skip_lsn[0]))
								v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
								v55 = F_SearchSysCacheCopy(m, int32(67), v53, int32(0))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return
								} else {
									if v55 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v155 = m.ExcPending
										if v155 != 0 {
											return
										} else {
											v157 = *(*int32)(unsafe.Add(mBase, _c_F_clear_subscription_skip_lsn[0]))
											v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+16))
											*(*int32)(unsafe.Add(mBase, uint32(v11))) = v158
											F_errmsg_internal(m, int32(_a_F_clear_subscription_skip_lsn_1), v11)
											mBase = m.M
											v162 = m.ExcPending
											if v162 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_clear_subscription_skip_lsn_2), int32(_a_F_clear_subscription_skip_lsn_3), int32(_a_F_clear_subscription_skip_lsn_4))
												mBase = m.M
												v167 = m.ExcPending
												if v167 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
										v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+22)))
										v62 = *(*int64)(unsafe.Add(mBase, uint32(v59+v60)+8))
										if v62 != v15 {
											v132 = v55
											F_pfree(m, v132)
											mBase = m.M
											v136 = m.ExcPending
											if v136 != 0 {
												return
											} else {
												F_relation_close(m, v48, int32(0))
												mBase = m.M
												v139 = m.ExcPending
												if v139 != 0 {
													return
												} else {
													F_PopActiveSnapshot(m)
													mBase = m.M
													v141 = m.ExcPending
													if v141 != 0 {
														return
													} else {
														if v28 == int32(2) {
															m.G0 = v11 + int32(192)
															return
														} else {
															F_CommitTransactionCommand(m)
															mBase = m.M
															v143 = m.ExcPending
															if v143 != 0 {
																return
															} else {
																m.G0 = v11 + int32(192)
																return
															}
														}
													}
												}
											}
										} else {
											v65 = v11 + int32(48)
											v66 = int32(0)
											base.MemoryFill(m, v65, v66, int32(72))
											*(*uint16)(unsafe.Add(mBase, uint32(v11)+176)) = uint16(v66)
											v71 = int64(0)
											*(*int64)(unsafe.Add(mBase, uint32(v11)+168)) = v71
											*(*int64)(unsafe.Add(mBase, uint32(v11)+160)) = v71
											*(*int64)(unsafe.Add(mBase, uint32(v11)+128)) = v71
											*(*int64)(unsafe.Add(mBase, uint32(v11)+136)) = v71
											*(*uint16)(unsafe.Add(mBase, uint32(v11)+144)) = uint16(v66)
											v82 = F_Int64GetDatum(m, v71)
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
												return
											} else {
												v84 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v11)+130)) = uint8(v84)
												*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v82
												v87 = *(*int32)(unsafe.Add(mBase, uint32(v48)+52))
												v92 = F_heap_modify_tuple(m, v55, v87, v65, v11+int32(160), v11+int32(128))
												mBase = m.M
												v93 = m.ExcPending
												if v93 != 0 {
													return
												} else {
													F_CatalogTupleUpdate(m, v48, v92+int32(4), v92)
													mBase = m.M
													v97 = m.ExcPending
													if v97 != 0 {
														return
													} else {
														if v1 == v15 {
															v132 = v92
															F_pfree(m, v132)
															mBase = m.M
															v136 = m.ExcPending
															if v136 != 0 {
																return
															} else {
																F_relation_close(m, v48, int32(0))
																mBase = m.M
																v139 = m.ExcPending
																if v139 != 0 {
																	return
																} else {
																	F_PopActiveSnapshot(m)
																	mBase = m.M
																	v141 = m.ExcPending
																	if v141 != 0 {
																		return
																	} else {
																		if v28 == int32(2) {
																			m.G0 = v11 + int32(192)
																			return
																		} else {
																			F_CommitTransactionCommand(m)
																			mBase = m.M
																			v143 = m.ExcPending
																			if v143 != 0 {
																				return
																			} else {
																				m.G0 = v11 + int32(192)
																				return
																			}
																		}
																	}
																}
															}
														} else {
															v101 = F_errstart(m, int32(19), int32(0))
															mBase = m.M
															v102 = m.ExcPending
															if v102 != 0 {
																return
															} else {
																if v101 == int32(0) {
																	v132 = v92
																	F_pfree(m, v132)
																	mBase = m.M
																	v136 = m.ExcPending
																	if v136 != 0 {
																		return
																	} else {
																		F_relation_close(m, v48, int32(0))
																		mBase = m.M
																		v139 = m.ExcPending
																		if v139 != 0 {
																			return
																		} else {
																			F_PopActiveSnapshot(m)
																			mBase = m.M
																			v141 = m.ExcPending
																			if v141 != 0 {
																				return
																			} else {
																				if v28 == int32(2) {
																					m.G0 = v11 + int32(192)
																					return
																				} else {
																					F_CommitTransactionCommand(m)
																					mBase = m.M
																					v143 = m.ExcPending
																					if v143 != 0 {
																						return
																					} else {
																						m.G0 = v11 + int32(192)
																						return
																					}
																				}
																			}
																		}
																	}
																} else {
																	v106 = *(*int32)(unsafe.Add(mBase, _c_F_clear_subscription_skip_lsn[0]))
																	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+16))
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v107
																	F_errmsg(m, int32(_a_F_clear_subscription_skip_lsn_5), v11+int32(32))
																	mBase = m.M
																	v113 = m.ExcPending
																	if v113 != 0 {
																		return
																	} else {
																		*(*uint32)(unsafe.Add(mBase, uint32(v11)+28)) = uint32(v15)
																		v115 = int64(32)
																		v116 = int64(base.Ui64(v15) >> (uint(v115) % 64))
																		*(*uint32)(unsafe.Add(mBase, uint32(v11)+24)) = uint32(v116)
																		*(*uint32)(unsafe.Add(mBase, uint32(v11)+20)) = uint32(v1)
																		v120 = int64(base.Ui64(v1) >> (uint(v115) % 64))
																		*(*uint32)(unsafe.Add(mBase, uint32(v11)+16)) = uint32(v120)
																		F_errdetail(m, int32(_a_F_clear_subscription_skip_lsn_6), v11+int32(16))
																		mBase = m.M
																		v126 = m.ExcPending
																		if v126 != 0 {
																			return
																		} else {
																			F_errfinish(m, int32(_a_F_clear_subscription_skip_lsn_2), int32(_a_F_clear_subscription_skip_lsn_7), int32(_a_F_clear_subscription_skip_lsn_4))
																			mBase = m.M
																			v131 = m.ExcPending
																			if v131 != 0 {
																				return
																			} else {
																				v132 = v92
																				F_pfree(m, v132)
																				mBase = m.M
																				v136 = m.ExcPending
																				if v136 != 0 {
																					return
																				} else {
																					F_relation_close(m, v48, int32(0))
																					mBase = m.M
																					v139 = m.ExcPending
																					if v139 != 0 {
																						return
																					} else {
																						F_PopActiveSnapshot(m)
																						mBase = m.M
																						v141 = m.ExcPending
																						if v141 != 0 {
																							return
																						} else {
																							if v28 == int32(2) {
																								m.G0 = v11 + int32(192)
																								return
																							} else {
																								F_CommitTransactionCommand(m)
																								mBase = m.M
																								v143 = m.ExcPending
																								if v143 != 0 {
																									return
																								} else {
																									m.G0 = v11 + int32(192)
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
