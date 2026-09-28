package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BackgroundWorkerInitializeConnectionByOid(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerInitializeConnectionByOid[0]))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+192)))
	if v6&int32(2) != 0 {
		v9 = int32(0)
		F_InitPostgres(m, v9, l0, v9, l1, l2<<(uint(int32(1))%32)&int32(6), v9)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerInitializeConnectionByOid[1]))
			if v19 != int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_BackgroundWorkerInitializeConnectionByOid_0), int32(0))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_BackgroundWorkerInitializeConnectionByOid_1), int32(926), int32(_a_F_BackgroundWorkerInitializeConnectionByOid_2))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerInitializeConnectionByOid[1])) = int32(2)
				return
			}
		}
	} else {
		F_errstart_cold(m, int32(22), int32(0))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return
		} else {
			F_errcode(m, int32(261))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				F_errmsg(m, int32(_a_F_BackgroundWorkerInitializeConnectionByOid_3), int32(0))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_BackgroundWorkerInitializeConnectionByOid_1), int32(916), int32(_a_F_BackgroundWorkerInitializeConnectionByOid_2))
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
}
func F_BackgroundWorkerShmemRequest(m *base.Module, l0 int32) {
	var v8 int32
	_ = v8
	Fn14205(m, l0, int32(_a_F_BackgroundWorkerShmemRequest_0), int32(_a_F_BackgroundWorkerShmemRequest_1), int32(1488), int32(_a_F_BackgroundWorkerShmemRequest_2), int32(16))
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
	}
}
func F_RegisterBackgroundWorker(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	v4 = m.G0
	v6 = v4 - int32(96)
	m.G0 = v6
	v9 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RegisterBackgroundWorker[0])))
	if v9 == int32(0) {
		v13 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RegisterBackgroundWorker[1])))
		if v13&int32(1) != 0 {
			v39 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterBackgroundWorker[2]))
			if v39 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v181 = m.ExcPending
				if v181 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6)+64)) = l0
					F_errmsg_internal(m, int32(_a_F_RegisterBackgroundWorker_0), v6-int32(-64))
					mBase = m.M
					v187 = m.ExcPending
					if v187 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_RegisterBackgroundWorker_1), int32(990), int32(_a_F_RegisterBackgroundWorker_2))
						mBase = m.M
						v192 = m.ExcPending
						if v192 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v42 = F_errstart(m, int32(14), int32(0))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					if v42 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = l0
						F_errmsg_internal(m, int32(_a_F_RegisterBackgroundWorker_3), v6+int32(48))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_RegisterBackgroundWorker_1), int32(993), int32(_a_F_RegisterBackgroundWorker_2))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return
							} else {
								v56 = F_SanityCheckBackgroundWorker(m, l0, int32(15))
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return
								} else {
									if v56 == int32(0) {
										m.G0 = v6 + int32(96)
										return
									} else {
										v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1464))
										if v60 != 0 {
											v63 = F_errstart(m, int32(15), int32(0))
											mBase = m.M
											v64 = m.ExcPending
											if v64 != 0 {
												return
											} else {
												if v63 == int32(0) {
													m.G0 = v6 + int32(96)
													return
												} else {
													F_errcode(m, int32(1088))
													mBase = m.M
													v69 = m.ExcPending
													if v69 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = l0
														F_errmsg(m, int32(_a_F_RegisterBackgroundWorker_4), v6+int32(32))
														mBase = m.M
														v75 = m.ExcPending
														if v75 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_RegisterBackgroundWorker_1), int32(1003), int32(_a_F_RegisterBackgroundWorker_2))
															mBase = m.M
															v80 = m.ExcPending
															if v80 != 0 {
																return
															} else {
																m.G0 = v6 + int32(96)
																return
															}
														}
													}
												}
											}
										} else {
											v81 = int32(_a_F_RegisterBackgroundWorker_5)
											v83 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterBackgroundWorker[3]))
											v85 = v83 + int32(1)
											*(*int32)(unsafe.Add(mBase, _c_F_RegisterBackgroundWorker[3])) = v85
											v88 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterBackgroundWorker[4]))
											if v88 < v85 {
												v92 = F_errstart(m, int32(15), int32(0))
												mBase = m.M
												v93 = m.ExcPending
												if v93 != 0 {
													return
												} else {
													if v92 == int32(0) {
														m.G0 = v6 + int32(96)
														return
													} else {
														F_errcode(m, int32(_a_F_RegisterBackgroundWorker_6))
														mBase = m.M
														v98 = m.ExcPending
														if v98 != 0 {
															return
														} else {
															F_errmsg(m, int32(_a_F_RegisterBackgroundWorker_7), int32(0))
															mBase = m.M
															v102 = m.ExcPending
															if v102 != 0 {
																return
															} else {
																v104 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterBackgroundWorker[4]))
																*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v104
																F_errdetail_plural(m, int32(_a_F_RegisterBackgroundWorker_8), int32(_a_F_RegisterBackgroundWorker_9), v104, v6+int32(16))
																mBase = m.M
																v111 = m.ExcPending
																if v111 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a_F_RegisterBackgroundWorker_10)
																	F_errhint(m, int32(_a_F_RegisterBackgroundWorker_11), v6)
																	mBase = m.M
																	v116 = m.ExcPending
																	if v116 != 0 {
																		return
																	} else {
																		F_errfinish(m, int32(_a_F_RegisterBackgroundWorker_1), int32(1022), int32(_a_F_RegisterBackgroundWorker_2))
																		mBase = m.M
																		v121 = m.ExcPending
																		if v121 != 0 {
																			return
																		} else {
																			m.G0 = v6 + int32(96)
																			return
																		}
																	}
																}
															}
														}
													}
												}
											} else {
												v123 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterBackgroundWorker[5]))
												v126 = F_MemoryContextAllocExtended(m, v123, int32(1504), int32(2))
												mBase = m.M
												v127 = m.ExcPending
												if v127 != 0 {
													return
												} else {
													if v126 == int32(0) {
														v132 = F_errstart(m, int32(15), int32(0))
														mBase = m.M
														v133 = m.ExcPending
														if v133 != 0 {
															return
														} else {
															if v132 == int32(0) {
																m.G0 = v6 + int32(96)
																return
															} else {
																F_errcode(m, int32(_a_F_RegisterBackgroundWorker_12))
																mBase = m.M
																v138 = m.ExcPending
																if v138 != 0 {
																	return
																} else {
																	F_errmsg(m, int32(_a_F_RegisterBackgroundWorker_13), int32(0))
																	mBase = m.M
																	v142 = m.ExcPending
																	if v142 != 0 {
																		return
																	} else {
																		F_errfinish(m, int32(_a_F_RegisterBackgroundWorker_1), int32(1036), int32(_a_F_RegisterBackgroundWorker_2))
																		mBase = m.M
																		v147 = m.ExcPending
																		if v147 != 0 {
																			return
																		} else {
																			m.G0 = v6 + int32(96)
																			return
																		}
																	}
																}
															}
														}
													} else {
														base.MemoryCopy(m, v126, l0, int32(1472))
														v150 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v126)+1492)) = uint8(v150)
														*(*int64)(unsafe.Add(mBase, uint32(v126)+1480)) = int64(0)
														*(*int32)(unsafe.Add(mBase, uint32(v126)+1472)) = v150
														v157 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterBackgroundWorker[6]))
														if v157 == v150 {
															v160 = int32(_a_F_RegisterBackgroundWorker_14)
															*(*int32)(unsafe.Add(mBase, _c_F_RegisterBackgroundWorker[7])) = v160
															v164 = v160
														} else {
															v164 = v157
														}
														*(*int32)(unsafe.Add(mBase, uint32(v126)+1496)) = int32(_a_F_RegisterBackgroundWorker_14)
														*(*int32)(unsafe.Add(mBase, uint32(v126)+1500)) = v164
														v169 = v126 + int32(1496)
														*(*int32)(unsafe.Add(mBase, uint32(v164))) = v169
														*(*int32)(unsafe.Add(mBase, _c_F_RegisterBackgroundWorker[6])) = v169
														m.G0 = v6 + int32(96)
														return
													}
												}
											}
										}
									}
								}
							}
						}
					} else {
						v56 = F_SanityCheckBackgroundWorker(m, l0, int32(15))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return
						} else {
							if v56 == int32(0) {
								m.G0 = v6 + int32(96)
								return
							} else {
								v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1464))
								if v60 != 0 {
									v63 = F_errstart(m, int32(15), int32(0))
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return
									} else {
										if v63 == int32(0) {
											m.G0 = v6 + int32(96)
											return
										} else {
											F_errcode(m, int32(1088))
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = l0
												F_errmsg(m, int32(_a_F_RegisterBackgroundWorker_4), v6+int32(32))
												mBase = m.M
												v75 = m.ExcPending
												if v75 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_RegisterBackgroundWorker_1), int32(1003), int32(_a_F_RegisterBackgroundWorker_2))
													mBase = m.M
													v80 = m.ExcPending
													if v80 != 0 {
														return
													} else {
														m.G0 = v6 + int32(96)
														return
													}
												}
											}
										}
									}
								} else {
									v81 = int32(_a_F_RegisterBackgroundWorker_5)
									v83 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterBackgroundWorker[3]))
									v85 = v83 + int32(1)
									*(*int32)(unsafe.Add(mBase, _c_F_RegisterBackgroundWorker[3])) = v85
									v88 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterBackgroundWorker[4]))
									if v88 < v85 {
										v92 = F_errstart(m, int32(15), int32(0))
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return
										} else {
											if v92 == int32(0) {
												m.G0 = v6 + int32(96)
												return
											} else {
												F_errcode(m, int32(_a_F_RegisterBackgroundWorker_6))
												mBase = m.M
												v98 = m.ExcPending
												if v98 != 0 {
													return
												} else {
													F_errmsg(m, int32(_a_F_RegisterBackgroundWorker_7), int32(0))
													mBase = m.M
													v102 = m.ExcPending
													if v102 != 0 {
														return
													} else {
														v104 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterBackgroundWorker[4]))
														*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v104
														F_errdetail_plural(m, int32(_a_F_RegisterBackgroundWorker_8), int32(_a_F_RegisterBackgroundWorker_9), v104, v6+int32(16))
														mBase = m.M
														v111 = m.ExcPending
														if v111 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a_F_RegisterBackgroundWorker_10)
															F_errhint(m, int32(_a_F_RegisterBackgroundWorker_11), v6)
															mBase = m.M
															v116 = m.ExcPending
															if v116 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_RegisterBackgroundWorker_1), int32(1022), int32(_a_F_RegisterBackgroundWorker_2))
																mBase = m.M
																v121 = m.ExcPending
																if v121 != 0 {
																	return
																} else {
																	m.G0 = v6 + int32(96)
																	return
																}
															}
														}
													}
												}
											}
										}
									} else {
										v123 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterBackgroundWorker[5]))
										v126 = F_MemoryContextAllocExtended(m, v123, int32(1504), int32(2))
										mBase = m.M
										v127 = m.ExcPending
										if v127 != 0 {
											return
										} else {
											if v126 == int32(0) {
												v132 = F_errstart(m, int32(15), int32(0))
												mBase = m.M
												v133 = m.ExcPending
												if v133 != 0 {
													return
												} else {
													if v132 == int32(0) {
														m.G0 = v6 + int32(96)
														return
													} else {
														F_errcode(m, int32(_a_F_RegisterBackgroundWorker_12))
														mBase = m.M
														v138 = m.ExcPending
														if v138 != 0 {
															return
														} else {
															F_errmsg(m, int32(_a_F_RegisterBackgroundWorker_13), int32(0))
															mBase = m.M
															v142 = m.ExcPending
															if v142 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_RegisterBackgroundWorker_1), int32(1036), int32(_a_F_RegisterBackgroundWorker_2))
																mBase = m.M
																v147 = m.ExcPending
																if v147 != 0 {
																	return
																} else {
																	m.G0 = v6 + int32(96)
																	return
																}
															}
														}
													}
												}
											} else {
												base.MemoryCopy(m, v126, l0, int32(1472))
												v150 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v126)+1492)) = uint8(v150)
												*(*int64)(unsafe.Add(mBase, uint32(v126)+1480)) = int64(0)
												*(*int32)(unsafe.Add(mBase, uint32(v126)+1472)) = v150
												v157 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterBackgroundWorker[6]))
												if v157 == v150 {
													v160 = int32(_a_F_RegisterBackgroundWorker_14)
													*(*int32)(unsafe.Add(mBase, _c_F_RegisterBackgroundWorker[7])) = v160
													v164 = v160
												} else {
													v164 = v157
												}
												*(*int32)(unsafe.Add(mBase, uint32(v126)+1496)) = int32(_a_F_RegisterBackgroundWorker_14)
												*(*int32)(unsafe.Add(mBase, uint32(v126)+1500)) = v164
												v169 = v126 + int32(1496)
												*(*int32)(unsafe.Add(mBase, uint32(v164))) = v169
												*(*int32)(unsafe.Add(mBase, _c_F_RegisterBackgroundWorker[6])) = v169
												m.G0 = v6 + int32(96)
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
		} else {
			v17 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RegisterBackgroundWorker[8])))
			if v17 != 0 {
				m.G0 = v6 + int32(96)
				return
			} else {
				v20 = F_errstart(m, int32(15), int32(0))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					if v20 == int32(0) {
						m.G0 = v6 + int32(96)
						return
					} else {
						F_errcode(m, int32(1088))
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+80)) = l0
							F_errmsg(m, int32(_a_F_RegisterBackgroundWorker_15), v6+int32(80))
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_RegisterBackgroundWorker_1), int32(980), int32(_a_F_RegisterBackgroundWorker_2))
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return
								} else {
									m.G0 = v6 + int32(96)
									return
								}
							}
						}
					}
				}
			}
		}
	} else {
		v17 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RegisterBackgroundWorker[8])))
		if v17 != 0 {
			m.G0 = v6 + int32(96)
			return
		} else {
			v20 = F_errstart(m, int32(15), int32(0))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				if v20 == int32(0) {
					m.G0 = v6 + int32(96)
					return
				} else {
					F_errcode(m, int32(1088))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6)+80)) = l0
						F_errmsg(m, int32(_a_F_RegisterBackgroundWorker_15), v6+int32(80))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_RegisterBackgroundWorker_1), int32(980), int32(_a_F_RegisterBackgroundWorker_2))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return
							} else {
								m.G0 = v6 + int32(96)
								return
							}
						}
					}
				}
			}
		}
	}
}
