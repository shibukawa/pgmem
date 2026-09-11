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
	var v52 int32
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
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int64
	_ = v73
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v117 int64
	_ = v117
	var v118 int64
	_ = v118
	var v122 int64
	_ = v122
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
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
	v8 = m.G0
	v10 = v8 - int32(192)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _consts[662]))
	v14 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
	if v14 == int64(0) {
		m.G0 = v10 + int32(192)
		return
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, _consts[648]))
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+16)))
		if v19 == int32(1) {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
			if v22 == int32(3) {
				m.G0 = v10 + int32(192)
				return
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, _consts[4]))
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
								v40 = *(*int32)(unsafe.Add(mBase, _consts[662]))
								v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
								F_LockSharedObject(m, int32(6100), v41, int32(1))
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return
								} else {
									v47 = F_table_open(m, int32(6100), int32(3))
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return
									} else {
										v51 = *(*int32)(unsafe.Add(mBase, _consts[662]))
										v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
										v54 = F_SearchSysCacheCopy(m, int32(67), v52, int32(0))
										mBase = m.M
										v55 = m.ExcPending
										if v55 != 0 {
											return
										} else {
											if v54 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v155 = m.ExcPending
												if v155 != 0 {
													return
												} else {
													v157 = *(*int32)(unsafe.Add(mBase, _consts[662]))
													v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+16))
													*(*int32)(unsafe.Add(mBase, uint32(v10))) = v158
													F_errmsg_internal(m, int32(67257), v10)
													mBase = m.M
													v162 = m.ExcPending
													if v162 != 0 {
														return
													} else {
														F_errfinish(m, int32(465060), int32(4990), int32(229343))
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
												v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
												v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+22)))
												v61 = *(*int64)(unsafe.Add(mBase, uint32(v58+v59)+8))
												if v61 != v14 {
													v134 = v54
													F_pfree(m, v134)
													mBase = m.M
													v137 = m.ExcPending
													if v137 != 0 {
														return
													} else {
														F_sequence_close(m, v47, int32(0))
														mBase = m.M
														v140 = m.ExcPending
														if v140 != 0 {
															return
														} else {
															F_PopActiveSnapshot(m)
															mBase = m.M
															v142 = m.ExcPending
															if v142 != 0 {
																return
															} else {
																if v27 == int32(2) {
																	m.G0 = v10 + int32(192)
																	return
																} else {
																	F_CommitTransactionCommand(m)
																	mBase = m.M
																	v144 = m.ExcPending
																	if v144 != 0 {
																		return
																	} else {
																		m.G0 = v10 + int32(192)
																		return
																	}
																}
															}
														}
													}
												} else {
													v68 = F__emscripten_memset_bulkmem(m, v10+int32(48), base.I32_extend8_s(int32(0)), int32(72))
													mBase = m.M
													v69 = int32(0)
													*(*uint16)(unsafe.Add(mBase, uint32(v10)+176)) = uint16(v69)
													*(*uint16)(unsafe.Add(mBase, uint32(v10)+144)) = uint16(v69)
													v73 = int64(0)
													*(*int64)(unsafe.Add(mBase, uint32(v10)+168)) = v73
													*(*int64)(unsafe.Add(mBase, uint32(v10)+160)) = v73
													*(*int64)(unsafe.Add(mBase, uint32(v10)+128)) = v73
													*(*int64)(unsafe.Add(mBase, uint32(v10)+136)) = v73
													v82 = F_Int64GetDatum(m, v73)
													mBase = m.M
													v83 = m.ExcPending
													if v83 != 0 {
														return
													} else {
														v84 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v10)+130)) = uint8(v84)
														*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = v82
														v87 = *(*int32)(unsafe.Add(mBase, uint32(v47)+52))
														v94 = F_heap_modify_tuple(m, v54, v87, v10+int32(48), v10+int32(160), v10+int32(128))
														mBase = m.M
														v95 = m.ExcPending
														if v95 != 0 {
															return
														} else {
															F_CatalogTupleUpdate(m, v47, v94+int32(4), v94)
															mBase = m.M
															v99 = m.ExcPending
															if v99 != 0 {
																return
															} else {
																if v1 == v14 {
																	v134 = v94
																	F_pfree(m, v134)
																	mBase = m.M
																	v137 = m.ExcPending
																	if v137 != 0 {
																		return
																	} else {
																		F_sequence_close(m, v47, int32(0))
																		mBase = m.M
																		v140 = m.ExcPending
																		if v140 != 0 {
																			return
																		} else {
																			F_PopActiveSnapshot(m)
																			mBase = m.M
																			v142 = m.ExcPending
																			if v142 != 0 {
																				return
																			} else {
																				if v27 == int32(2) {
																					m.G0 = v10 + int32(192)
																					return
																				} else {
																					F_CommitTransactionCommand(m)
																					mBase = m.M
																					v144 = m.ExcPending
																					if v144 != 0 {
																						return
																					} else {
																						m.G0 = v10 + int32(192)
																						return
																					}
																				}
																			}
																		}
																	}
																} else {
																	v103 = F_errstart(m, int32(19), int32(0))
																	mBase = m.M
																	v104 = m.ExcPending
																	if v104 != 0 {
																		return
																	} else {
																		if v103 == int32(0) {
																			v134 = v94
																			F_pfree(m, v134)
																			mBase = m.M
																			v137 = m.ExcPending
																			if v137 != 0 {
																				return
																			} else {
																				F_sequence_close(m, v47, int32(0))
																				mBase = m.M
																				v140 = m.ExcPending
																				if v140 != 0 {
																					return
																				} else {
																					F_PopActiveSnapshot(m)
																					mBase = m.M
																					v142 = m.ExcPending
																					if v142 != 0 {
																						return
																					} else {
																						if v27 == int32(2) {
																							m.G0 = v10 + int32(192)
																							return
																						} else {
																							F_CommitTransactionCommand(m)
																							mBase = m.M
																							v144 = m.ExcPending
																							if v144 != 0 {
																								return
																							} else {
																								m.G0 = v10 + int32(192)
																								return
																							}
																						}
																					}
																				}
																			}
																		} else {
																			v108 = *(*int32)(unsafe.Add(mBase, _consts[662]))
																			v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+16))
																			*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v109
																			F_errmsg(m, int32(424639), v10+int32(32))
																			mBase = m.M
																			v115 = m.ExcPending
																			if v115 != 0 {
																				return
																			} else {
																				*(*uint32)(unsafe.Add(mBase, uint32(v10)+28)) = uint32(v14)
																				v117 = int64(32)
																				v118 = int64(base.Ui64(v14) >> (uint(v117) % 64))
																				*(*uint32)(unsafe.Add(mBase, uint32(v10)+24)) = uint32(v118)
																				*(*uint32)(unsafe.Add(mBase, uint32(v10)+20)) = uint32(v1)
																				v122 = int64(base.Ui64(v1) >> (uint(v117) % 64))
																				*(*uint32)(unsafe.Add(mBase, uint32(v10)+16)) = uint32(v122)
																				F_errdetail(m, int32(607087), v10+int32(16))
																				mBase = m.M
																				v128 = m.ExcPending
																				if v128 != 0 {
																					return
																				} else {
																					F_errfinish(m, int32(465060), int32(5028), int32(229343))
																					mBase = m.M
																					v133 = m.ExcPending
																					if v133 != 0 {
																						return
																					} else {
																						v134 = v94
																						F_pfree(m, v134)
																						mBase = m.M
																						v137 = m.ExcPending
																						if v137 != 0 {
																							return
																						} else {
																							F_sequence_close(m, v47, int32(0))
																							mBase = m.M
																							v140 = m.ExcPending
																							if v140 != 0 {
																								return
																							} else {
																								F_PopActiveSnapshot(m)
																								mBase = m.M
																								v142 = m.ExcPending
																								if v142 != 0 {
																									return
																								} else {
																									if v27 == int32(2) {
																										m.G0 = v10 + int32(192)
																										return
																									} else {
																										F_CommitTransactionCommand(m)
																										mBase = m.M
																										v144 = m.ExcPending
																										if v144 != 0 {
																											return
																										} else {
																											m.G0 = v10 + int32(192)
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
							v40 = *(*int32)(unsafe.Add(mBase, _consts[662]))
							v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
							F_LockSharedObject(m, int32(6100), v41, int32(1))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return
							} else {
								v47 = F_table_open(m, int32(6100), int32(3))
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return
								} else {
									v51 = *(*int32)(unsafe.Add(mBase, _consts[662]))
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
									v54 = F_SearchSysCacheCopy(m, int32(67), v52, int32(0))
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return
									} else {
										if v54 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v155 = m.ExcPending
											if v155 != 0 {
												return
											} else {
												v157 = *(*int32)(unsafe.Add(mBase, _consts[662]))
												v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+16))
												*(*int32)(unsafe.Add(mBase, uint32(v10))) = v158
												F_errmsg_internal(m, int32(67257), v10)
												mBase = m.M
												v162 = m.ExcPending
												if v162 != 0 {
													return
												} else {
													F_errfinish(m, int32(465060), int32(4990), int32(229343))
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
											v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
											v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+22)))
											v61 = *(*int64)(unsafe.Add(mBase, uint32(v58+v59)+8))
											if v61 != v14 {
												v134 = v54
												F_pfree(m, v134)
												mBase = m.M
												v137 = m.ExcPending
												if v137 != 0 {
													return
												} else {
													F_sequence_close(m, v47, int32(0))
													mBase = m.M
													v140 = m.ExcPending
													if v140 != 0 {
														return
													} else {
														F_PopActiveSnapshot(m)
														mBase = m.M
														v142 = m.ExcPending
														if v142 != 0 {
															return
														} else {
															if v27 == int32(2) {
																m.G0 = v10 + int32(192)
																return
															} else {
																F_CommitTransactionCommand(m)
																mBase = m.M
																v144 = m.ExcPending
																if v144 != 0 {
																	return
																} else {
																	m.G0 = v10 + int32(192)
																	return
																}
															}
														}
													}
												}
											} else {
												v68 = F__emscripten_memset_bulkmem(m, v10+int32(48), base.I32_extend8_s(int32(0)), int32(72))
												mBase = m.M
												v69 = int32(0)
												*(*uint16)(unsafe.Add(mBase, uint32(v10)+176)) = uint16(v69)
												*(*uint16)(unsafe.Add(mBase, uint32(v10)+144)) = uint16(v69)
												v73 = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(v10)+168)) = v73
												*(*int64)(unsafe.Add(mBase, uint32(v10)+160)) = v73
												*(*int64)(unsafe.Add(mBase, uint32(v10)+128)) = v73
												*(*int64)(unsafe.Add(mBase, uint32(v10)+136)) = v73
												v82 = F_Int64GetDatum(m, v73)
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
													return
												} else {
													v84 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v10)+130)) = uint8(v84)
													*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = v82
													v87 = *(*int32)(unsafe.Add(mBase, uint32(v47)+52))
													v94 = F_heap_modify_tuple(m, v54, v87, v10+int32(48), v10+int32(160), v10+int32(128))
													mBase = m.M
													v95 = m.ExcPending
													if v95 != 0 {
														return
													} else {
														F_CatalogTupleUpdate(m, v47, v94+int32(4), v94)
														mBase = m.M
														v99 = m.ExcPending
														if v99 != 0 {
															return
														} else {
															if v1 == v14 {
																v134 = v94
																F_pfree(m, v134)
																mBase = m.M
																v137 = m.ExcPending
																if v137 != 0 {
																	return
																} else {
																	F_sequence_close(m, v47, int32(0))
																	mBase = m.M
																	v140 = m.ExcPending
																	if v140 != 0 {
																		return
																	} else {
																		F_PopActiveSnapshot(m)
																		mBase = m.M
																		v142 = m.ExcPending
																		if v142 != 0 {
																			return
																		} else {
																			if v27 == int32(2) {
																				m.G0 = v10 + int32(192)
																				return
																			} else {
																				F_CommitTransactionCommand(m)
																				mBase = m.M
																				v144 = m.ExcPending
																				if v144 != 0 {
																					return
																				} else {
																					m.G0 = v10 + int32(192)
																					return
																				}
																			}
																		}
																	}
																}
															} else {
																v103 = F_errstart(m, int32(19), int32(0))
																mBase = m.M
																v104 = m.ExcPending
																if v104 != 0 {
																	return
																} else {
																	if v103 == int32(0) {
																		v134 = v94
																		F_pfree(m, v134)
																		mBase = m.M
																		v137 = m.ExcPending
																		if v137 != 0 {
																			return
																		} else {
																			F_sequence_close(m, v47, int32(0))
																			mBase = m.M
																			v140 = m.ExcPending
																			if v140 != 0 {
																				return
																			} else {
																				F_PopActiveSnapshot(m)
																				mBase = m.M
																				v142 = m.ExcPending
																				if v142 != 0 {
																					return
																				} else {
																					if v27 == int32(2) {
																						m.G0 = v10 + int32(192)
																						return
																					} else {
																						F_CommitTransactionCommand(m)
																						mBase = m.M
																						v144 = m.ExcPending
																						if v144 != 0 {
																							return
																						} else {
																							m.G0 = v10 + int32(192)
																							return
																						}
																					}
																				}
																			}
																		}
																	} else {
																		v108 = *(*int32)(unsafe.Add(mBase, _consts[662]))
																		v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+16))
																		*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v109
																		F_errmsg(m, int32(424639), v10+int32(32))
																		mBase = m.M
																		v115 = m.ExcPending
																		if v115 != 0 {
																			return
																		} else {
																			*(*uint32)(unsafe.Add(mBase, uint32(v10)+28)) = uint32(v14)
																			v117 = int64(32)
																			v118 = int64(base.Ui64(v14) >> (uint(v117) % 64))
																			*(*uint32)(unsafe.Add(mBase, uint32(v10)+24)) = uint32(v118)
																			*(*uint32)(unsafe.Add(mBase, uint32(v10)+20)) = uint32(v1)
																			v122 = int64(base.Ui64(v1) >> (uint(v117) % 64))
																			*(*uint32)(unsafe.Add(mBase, uint32(v10)+16)) = uint32(v122)
																			F_errdetail(m, int32(607087), v10+int32(16))
																			mBase = m.M
																			v128 = m.ExcPending
																			if v128 != 0 {
																				return
																			} else {
																				F_errfinish(m, int32(465060), int32(5028), int32(229343))
																				mBase = m.M
																				v133 = m.ExcPending
																				if v133 != 0 {
																					return
																				} else {
																					v134 = v94
																					F_pfree(m, v134)
																					mBase = m.M
																					v137 = m.ExcPending
																					if v137 != 0 {
																						return
																					} else {
																						F_sequence_close(m, v47, int32(0))
																						mBase = m.M
																						v140 = m.ExcPending
																						if v140 != 0 {
																							return
																						} else {
																							F_PopActiveSnapshot(m)
																							mBase = m.M
																							v142 = m.ExcPending
																							if v142 != 0 {
																								return
																							} else {
																								if v27 == int32(2) {
																									m.G0 = v10 + int32(192)
																									return
																								} else {
																									F_CommitTransactionCommand(m)
																									mBase = m.M
																									v144 = m.ExcPending
																									if v144 != 0 {
																										return
																									} else {
																										m.G0 = v10 + int32(192)
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
			v26 = *(*int32)(unsafe.Add(mBase, _consts[4]))
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
							v40 = *(*int32)(unsafe.Add(mBase, _consts[662]))
							v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
							F_LockSharedObject(m, int32(6100), v41, int32(1))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return
							} else {
								v47 = F_table_open(m, int32(6100), int32(3))
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return
								} else {
									v51 = *(*int32)(unsafe.Add(mBase, _consts[662]))
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
									v54 = F_SearchSysCacheCopy(m, int32(67), v52, int32(0))
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return
									} else {
										if v54 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v155 = m.ExcPending
											if v155 != 0 {
												return
											} else {
												v157 = *(*int32)(unsafe.Add(mBase, _consts[662]))
												v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+16))
												*(*int32)(unsafe.Add(mBase, uint32(v10))) = v158
												F_errmsg_internal(m, int32(67257), v10)
												mBase = m.M
												v162 = m.ExcPending
												if v162 != 0 {
													return
												} else {
													F_errfinish(m, int32(465060), int32(4990), int32(229343))
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
											v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
											v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+22)))
											v61 = *(*int64)(unsafe.Add(mBase, uint32(v58+v59)+8))
											if v61 != v14 {
												v134 = v54
												F_pfree(m, v134)
												mBase = m.M
												v137 = m.ExcPending
												if v137 != 0 {
													return
												} else {
													F_sequence_close(m, v47, int32(0))
													mBase = m.M
													v140 = m.ExcPending
													if v140 != 0 {
														return
													} else {
														F_PopActiveSnapshot(m)
														mBase = m.M
														v142 = m.ExcPending
														if v142 != 0 {
															return
														} else {
															if v27 == int32(2) {
																m.G0 = v10 + int32(192)
																return
															} else {
																F_CommitTransactionCommand(m)
																mBase = m.M
																v144 = m.ExcPending
																if v144 != 0 {
																	return
																} else {
																	m.G0 = v10 + int32(192)
																	return
																}
															}
														}
													}
												}
											} else {
												v68 = F__emscripten_memset_bulkmem(m, v10+int32(48), base.I32_extend8_s(int32(0)), int32(72))
												mBase = m.M
												v69 = int32(0)
												*(*uint16)(unsafe.Add(mBase, uint32(v10)+176)) = uint16(v69)
												*(*uint16)(unsafe.Add(mBase, uint32(v10)+144)) = uint16(v69)
												v73 = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(v10)+168)) = v73
												*(*int64)(unsafe.Add(mBase, uint32(v10)+160)) = v73
												*(*int64)(unsafe.Add(mBase, uint32(v10)+128)) = v73
												*(*int64)(unsafe.Add(mBase, uint32(v10)+136)) = v73
												v82 = F_Int64GetDatum(m, v73)
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
													return
												} else {
													v84 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v10)+130)) = uint8(v84)
													*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = v82
													v87 = *(*int32)(unsafe.Add(mBase, uint32(v47)+52))
													v94 = F_heap_modify_tuple(m, v54, v87, v10+int32(48), v10+int32(160), v10+int32(128))
													mBase = m.M
													v95 = m.ExcPending
													if v95 != 0 {
														return
													} else {
														F_CatalogTupleUpdate(m, v47, v94+int32(4), v94)
														mBase = m.M
														v99 = m.ExcPending
														if v99 != 0 {
															return
														} else {
															if v1 == v14 {
																v134 = v94
																F_pfree(m, v134)
																mBase = m.M
																v137 = m.ExcPending
																if v137 != 0 {
																	return
																} else {
																	F_sequence_close(m, v47, int32(0))
																	mBase = m.M
																	v140 = m.ExcPending
																	if v140 != 0 {
																		return
																	} else {
																		F_PopActiveSnapshot(m)
																		mBase = m.M
																		v142 = m.ExcPending
																		if v142 != 0 {
																			return
																		} else {
																			if v27 == int32(2) {
																				m.G0 = v10 + int32(192)
																				return
																			} else {
																				F_CommitTransactionCommand(m)
																				mBase = m.M
																				v144 = m.ExcPending
																				if v144 != 0 {
																					return
																				} else {
																					m.G0 = v10 + int32(192)
																					return
																				}
																			}
																		}
																	}
																}
															} else {
																v103 = F_errstart(m, int32(19), int32(0))
																mBase = m.M
																v104 = m.ExcPending
																if v104 != 0 {
																	return
																} else {
																	if v103 == int32(0) {
																		v134 = v94
																		F_pfree(m, v134)
																		mBase = m.M
																		v137 = m.ExcPending
																		if v137 != 0 {
																			return
																		} else {
																			F_sequence_close(m, v47, int32(0))
																			mBase = m.M
																			v140 = m.ExcPending
																			if v140 != 0 {
																				return
																			} else {
																				F_PopActiveSnapshot(m)
																				mBase = m.M
																				v142 = m.ExcPending
																				if v142 != 0 {
																					return
																				} else {
																					if v27 == int32(2) {
																						m.G0 = v10 + int32(192)
																						return
																					} else {
																						F_CommitTransactionCommand(m)
																						mBase = m.M
																						v144 = m.ExcPending
																						if v144 != 0 {
																							return
																						} else {
																							m.G0 = v10 + int32(192)
																							return
																						}
																					}
																				}
																			}
																		}
																	} else {
																		v108 = *(*int32)(unsafe.Add(mBase, _consts[662]))
																		v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+16))
																		*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v109
																		F_errmsg(m, int32(424639), v10+int32(32))
																		mBase = m.M
																		v115 = m.ExcPending
																		if v115 != 0 {
																			return
																		} else {
																			*(*uint32)(unsafe.Add(mBase, uint32(v10)+28)) = uint32(v14)
																			v117 = int64(32)
																			v118 = int64(base.Ui64(v14) >> (uint(v117) % 64))
																			*(*uint32)(unsafe.Add(mBase, uint32(v10)+24)) = uint32(v118)
																			*(*uint32)(unsafe.Add(mBase, uint32(v10)+20)) = uint32(v1)
																			v122 = int64(base.Ui64(v1) >> (uint(v117) % 64))
																			*(*uint32)(unsafe.Add(mBase, uint32(v10)+16)) = uint32(v122)
																			F_errdetail(m, int32(607087), v10+int32(16))
																			mBase = m.M
																			v128 = m.ExcPending
																			if v128 != 0 {
																				return
																			} else {
																				F_errfinish(m, int32(465060), int32(5028), int32(229343))
																				mBase = m.M
																				v133 = m.ExcPending
																				if v133 != 0 {
																					return
																				} else {
																					v134 = v94
																					F_pfree(m, v134)
																					mBase = m.M
																					v137 = m.ExcPending
																					if v137 != 0 {
																						return
																					} else {
																						F_sequence_close(m, v47, int32(0))
																						mBase = m.M
																						v140 = m.ExcPending
																						if v140 != 0 {
																							return
																						} else {
																							F_PopActiveSnapshot(m)
																							mBase = m.M
																							v142 = m.ExcPending
																							if v142 != 0 {
																								return
																							} else {
																								if v27 == int32(2) {
																									m.G0 = v10 + int32(192)
																									return
																								} else {
																									F_CommitTransactionCommand(m)
																									mBase = m.M
																									v144 = m.ExcPending
																									if v144 != 0 {
																										return
																									} else {
																										m.G0 = v10 + int32(192)
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
						v40 = *(*int32)(unsafe.Add(mBase, _consts[662]))
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
						F_LockSharedObject(m, int32(6100), v41, int32(1))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return
						} else {
							v47 = F_table_open(m, int32(6100), int32(3))
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return
							} else {
								v51 = *(*int32)(unsafe.Add(mBase, _consts[662]))
								v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
								v54 = F_SearchSysCacheCopy(m, int32(67), v52, int32(0))
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return
								} else {
									if v54 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v155 = m.ExcPending
										if v155 != 0 {
											return
										} else {
											v157 = *(*int32)(unsafe.Add(mBase, _consts[662]))
											v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+16))
											*(*int32)(unsafe.Add(mBase, uint32(v10))) = v158
											F_errmsg_internal(m, int32(67257), v10)
											mBase = m.M
											v162 = m.ExcPending
											if v162 != 0 {
												return
											} else {
												F_errfinish(m, int32(465060), int32(4990), int32(229343))
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
										v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
										v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+22)))
										v61 = *(*int64)(unsafe.Add(mBase, uint32(v58+v59)+8))
										if v61 != v14 {
											v134 = v54
											F_pfree(m, v134)
											mBase = m.M
											v137 = m.ExcPending
											if v137 != 0 {
												return
											} else {
												F_sequence_close(m, v47, int32(0))
												mBase = m.M
												v140 = m.ExcPending
												if v140 != 0 {
													return
												} else {
													F_PopActiveSnapshot(m)
													mBase = m.M
													v142 = m.ExcPending
													if v142 != 0 {
														return
													} else {
														if v27 == int32(2) {
															m.G0 = v10 + int32(192)
															return
														} else {
															F_CommitTransactionCommand(m)
															mBase = m.M
															v144 = m.ExcPending
															if v144 != 0 {
																return
															} else {
																m.G0 = v10 + int32(192)
																return
															}
														}
													}
												}
											}
										} else {
											v68 = F__emscripten_memset_bulkmem(m, v10+int32(48), base.I32_extend8_s(int32(0)), int32(72))
											mBase = m.M
											v69 = int32(0)
											*(*uint16)(unsafe.Add(mBase, uint32(v10)+176)) = uint16(v69)
											*(*uint16)(unsafe.Add(mBase, uint32(v10)+144)) = uint16(v69)
											v73 = int64(0)
											*(*int64)(unsafe.Add(mBase, uint32(v10)+168)) = v73
											*(*int64)(unsafe.Add(mBase, uint32(v10)+160)) = v73
											*(*int64)(unsafe.Add(mBase, uint32(v10)+128)) = v73
											*(*int64)(unsafe.Add(mBase, uint32(v10)+136)) = v73
											v82 = F_Int64GetDatum(m, v73)
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
												return
											} else {
												v84 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v10)+130)) = uint8(v84)
												*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = v82
												v87 = *(*int32)(unsafe.Add(mBase, uint32(v47)+52))
												v94 = F_heap_modify_tuple(m, v54, v87, v10+int32(48), v10+int32(160), v10+int32(128))
												mBase = m.M
												v95 = m.ExcPending
												if v95 != 0 {
													return
												} else {
													F_CatalogTupleUpdate(m, v47, v94+int32(4), v94)
													mBase = m.M
													v99 = m.ExcPending
													if v99 != 0 {
														return
													} else {
														if v1 == v14 {
															v134 = v94
															F_pfree(m, v134)
															mBase = m.M
															v137 = m.ExcPending
															if v137 != 0 {
																return
															} else {
																F_sequence_close(m, v47, int32(0))
																mBase = m.M
																v140 = m.ExcPending
																if v140 != 0 {
																	return
																} else {
																	F_PopActiveSnapshot(m)
																	mBase = m.M
																	v142 = m.ExcPending
																	if v142 != 0 {
																		return
																	} else {
																		if v27 == int32(2) {
																			m.G0 = v10 + int32(192)
																			return
																		} else {
																			F_CommitTransactionCommand(m)
																			mBase = m.M
																			v144 = m.ExcPending
																			if v144 != 0 {
																				return
																			} else {
																				m.G0 = v10 + int32(192)
																				return
																			}
																		}
																	}
																}
															}
														} else {
															v103 = F_errstart(m, int32(19), int32(0))
															mBase = m.M
															v104 = m.ExcPending
															if v104 != 0 {
																return
															} else {
																if v103 == int32(0) {
																	v134 = v94
																	F_pfree(m, v134)
																	mBase = m.M
																	v137 = m.ExcPending
																	if v137 != 0 {
																		return
																	} else {
																		F_sequence_close(m, v47, int32(0))
																		mBase = m.M
																		v140 = m.ExcPending
																		if v140 != 0 {
																			return
																		} else {
																			F_PopActiveSnapshot(m)
																			mBase = m.M
																			v142 = m.ExcPending
																			if v142 != 0 {
																				return
																			} else {
																				if v27 == int32(2) {
																					m.G0 = v10 + int32(192)
																					return
																				} else {
																					F_CommitTransactionCommand(m)
																					mBase = m.M
																					v144 = m.ExcPending
																					if v144 != 0 {
																						return
																					} else {
																						m.G0 = v10 + int32(192)
																						return
																					}
																				}
																			}
																		}
																	}
																} else {
																	v108 = *(*int32)(unsafe.Add(mBase, _consts[662]))
																	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+16))
																	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v109
																	F_errmsg(m, int32(424639), v10+int32(32))
																	mBase = m.M
																	v115 = m.ExcPending
																	if v115 != 0 {
																		return
																	} else {
																		*(*uint32)(unsafe.Add(mBase, uint32(v10)+28)) = uint32(v14)
																		v117 = int64(32)
																		v118 = int64(base.Ui64(v14) >> (uint(v117) % 64))
																		*(*uint32)(unsafe.Add(mBase, uint32(v10)+24)) = uint32(v118)
																		*(*uint32)(unsafe.Add(mBase, uint32(v10)+20)) = uint32(v1)
																		v122 = int64(base.Ui64(v1) >> (uint(v117) % 64))
																		*(*uint32)(unsafe.Add(mBase, uint32(v10)+16)) = uint32(v122)
																		F_errdetail(m, int32(607087), v10+int32(16))
																		mBase = m.M
																		v128 = m.ExcPending
																		if v128 != 0 {
																			return
																		} else {
																			F_errfinish(m, int32(465060), int32(5028), int32(229343))
																			mBase = m.M
																			v133 = m.ExcPending
																			if v133 != 0 {
																				return
																			} else {
																				v134 = v94
																				F_pfree(m, v134)
																				mBase = m.M
																				v137 = m.ExcPending
																				if v137 != 0 {
																					return
																				} else {
																					F_sequence_close(m, v47, int32(0))
																					mBase = m.M
																					v140 = m.ExcPending
																					if v140 != 0 {
																						return
																					} else {
																						F_PopActiveSnapshot(m)
																						mBase = m.M
																						v142 = m.ExcPending
																						if v142 != 0 {
																							return
																						} else {
																							if v27 == int32(2) {
																								m.G0 = v10 + int32(192)
																								return
																							} else {
																								F_CommitTransactionCommand(m)
																								mBase = m.M
																								v144 = m.ExcPending
																								if v144 != 0 {
																									return
																								} else {
																									m.G0 = v10 + int32(192)
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
