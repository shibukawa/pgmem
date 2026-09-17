package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ReadDir(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_ReadDirExtended(m, l0, l1, int32(21))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_ReadNextFullTransactionId(m *base.Module) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_ReadNextFullTransactionId[0]))
	v7 = F_LWLockAcquire(m, v3+int32(384), int32(1))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_ReadNextFullTransactionId[1]))
		v13 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
		v15 = *(*int32)(unsafe.Add(mBase, _c_F_ReadNextFullTransactionId[0]))
		F_LWLockRelease(m, v15+int32(384))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			return v13
		}
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
						F_errfinish(m, int32(_a_F_RegisterBackgroundWorker_1), int32(977), int32(_a_F_RegisterBackgroundWorker_2))
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
							F_errfinish(m, int32(_a_F_RegisterBackgroundWorker_1), int32(980), int32(_a_F_RegisterBackgroundWorker_2))
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
										v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1456))
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
															F_errfinish(m, int32(_a_F_RegisterBackgroundWorker_1), int32(990), int32(_a_F_RegisterBackgroundWorker_2))
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
																		F_errfinish(m, int32(_a_F_RegisterBackgroundWorker_1), int32(1009), int32(_a_F_RegisterBackgroundWorker_2))
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
												v126 = F_MemoryContextAllocExtended(m, v123, int32(1488), int32(2))
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
																		F_errfinish(m, int32(_a_F_RegisterBackgroundWorker_1), int32(1023), int32(_a_F_RegisterBackgroundWorker_2))
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
														base.MemoryCopy(m, v126, l0, int32(1460))
														v150 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v126)+1476)) = uint8(v150)
														*(*int64)(unsafe.Add(mBase, uint32(v126)+1464)) = int64(0)
														*(*int32)(unsafe.Add(mBase, uint32(v126)+1460)) = v150
														v157 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterBackgroundWorker[6]))
														if v157 == v150 {
															v160 = int32(_a_F_RegisterBackgroundWorker_14)
															*(*int32)(unsafe.Add(mBase, _c_F_RegisterBackgroundWorker[7])) = v160
															v164 = v160
														} else {
															v164 = v157
														}
														*(*int32)(unsafe.Add(mBase, uint32(v126)+1480)) = int32(_a_F_RegisterBackgroundWorker_14)
														*(*int32)(unsafe.Add(mBase, uint32(v126)+1484)) = v164
														v169 = v126 + int32(1480)
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
								v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1456))
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
													F_errfinish(m, int32(_a_F_RegisterBackgroundWorker_1), int32(990), int32(_a_F_RegisterBackgroundWorker_2))
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
																F_errfinish(m, int32(_a_F_RegisterBackgroundWorker_1), int32(1009), int32(_a_F_RegisterBackgroundWorker_2))
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
										v126 = F_MemoryContextAllocExtended(m, v123, int32(1488), int32(2))
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
																F_errfinish(m, int32(_a_F_RegisterBackgroundWorker_1), int32(1023), int32(_a_F_RegisterBackgroundWorker_2))
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
												base.MemoryCopy(m, v126, l0, int32(1460))
												v150 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v126)+1476)) = uint8(v150)
												*(*int64)(unsafe.Add(mBase, uint32(v126)+1464)) = int64(0)
												*(*int32)(unsafe.Add(mBase, uint32(v126)+1460)) = v150
												v157 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterBackgroundWorker[6]))
												if v157 == v150 {
													v160 = int32(_a_F_RegisterBackgroundWorker_14)
													*(*int32)(unsafe.Add(mBase, _c_F_RegisterBackgroundWorker[7])) = v160
													v164 = v160
												} else {
													v164 = v157
												}
												*(*int32)(unsafe.Add(mBase, uint32(v126)+1480)) = int32(_a_F_RegisterBackgroundWorker_14)
												*(*int32)(unsafe.Add(mBase, uint32(v126)+1484)) = v164
												v169 = v126 + int32(1480)
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
								F_errfinish(m, int32(_a_F_RegisterBackgroundWorker_1), int32(967), int32(_a_F_RegisterBackgroundWorker_2))
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
							F_errfinish(m, int32(_a_F_RegisterBackgroundWorker_1), int32(967), int32(_a_F_RegisterBackgroundWorker_2))
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
func F_RegisterCatcacheInvalidation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	v1 = l0
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterCatcacheInvalidation[0]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterCatcacheInvalidation[1]))
	if v13 <= v11 {
		if v10 == int32(0) {
			v19 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterCatcacheInvalidation[2]))
			v21 = F_MemoryContextAlloc(m, v19, int32(512))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				v29 = int32(32)
				v30 = v21
				*(*int32)(unsafe.Add(mBase, _c_F_RegisterCatcacheInvalidation[1])) = v29
				*(*int32)(unsafe.Add(mBase, _c_F_RegisterCatcacheInvalidation[0])) = v30
				v35 = v30
				v39 = v35 + v11<<(uint(int32(4))%32)
				*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = l2
				*(*uint8)(unsafe.Add(mBase, uint32(v39))) = uint8(v1)
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v43 + int32(1)
				return
			}
		} else {
			v27 = F_repalloc(m, v10, v13<<(uint(int32(5))%32))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				v29 = v13 << (uint(int32(1)) % 32)
				v30 = v27
				*(*int32)(unsafe.Add(mBase, _c_F_RegisterCatcacheInvalidation[1])) = v29
				*(*int32)(unsafe.Add(mBase, _c_F_RegisterCatcacheInvalidation[0])) = v30
				v35 = v30
				v39 = v35 + v11<<(uint(int32(4))%32)
				*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = l2
				*(*uint8)(unsafe.Add(mBase, uint32(v39))) = uint8(v1)
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v43 + int32(1)
				return
			}
		}
	} else {
		v35 = v10
		v39 = v35 + v11<<(uint(int32(4))%32)
		*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = l2
		*(*uint8)(unsafe.Add(mBase, uint32(v39))) = uint8(v1)
		v43 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v43 + int32(1)
		return
	}
}
func F_ReleaseAuxProcessResources(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseAuxProcessResources[0]))
	v4 = int32(1)
	F_ResourceOwnerReleaseInternal(m, v3, v4, l0, v4)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseAuxProcessResources[0]))
		F_ResourceOwnerReleaseInternal(m, v9, int32(2), l0, int32(1))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseAuxProcessResources[0]))
			F_ResourceOwnerReleaseInternal(m, v15, int32(3), l0, int32(1))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseAuxProcessResources[0]))
				v22 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v21)+16)) = uint16(v22)
				return
			}
		}
	}
}
func F_RepairGraphElement(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v174 int32
	_ = v174
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+65)))
	v24 = F_add_size(m, v22, int32(2))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return
	} else {
		v26 = F_mul_size(m, v24, v19)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			v28 = F_mul_size(m, int32(6), v26)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				v30 = F_add_size(m, int32(4), v28)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					if l2 == int32(0) {
						v40 = int32(0)
						F_HnswInitNeighbors(m, v40, l1, v19, v40)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							v44 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l1)+64)) = uint8(v44)
							v48 = l0 + int32(24)
							F_HnswFindElementNeighbors(m, v44, l1, l2, v18, v48, v19, v17, int32(1))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return
							} else {
								v52 = int32(0)
								base.MemoryFill(m, v15, v52, int32(_a_F_RepairGraphElement_0))
								v67 = int32(2)
								*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v67)
								v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+65)))
								v79 = v71
								v83 = v52
								for {
									v98 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
									v102 = *(*int32)(unsafe.Add(mBase, uint32(v98+v79<<(uint(int32(2))%32))))
									if base.B2i32(v19 <= v52) == int32(0) {
										v112 = int32(0)
										v115 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
										v121 = v112
										v126 = v83
										for {
											v134 = v15 + int32(4) + v126*int32(6)
											if v121 < v115 {
												v141 = *(*int32)(unsafe.Add(mBase, uint32(v102+int32(8)+v121*int32(12))))
												v147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v141)+80)))
												v148 = *(*int32)(unsafe.Add(mBase, uint32(v141)+76))
												v150 = int32(base.Ui32(v148) >> (uint(int32(16)) % 32))
												*(*uint16)(unsafe.Add(mBase, uint32(v134))) = uint16(v150)
												v156 = v147
												v157 = v148
											} else {
												v152 = int32(_a_F_RepairGraphElement_1)
												*(*uint16)(unsafe.Add(mBase, uint32(v134))) = uint16(v152)
												v156 = int32(0)
												v157 = v152
											}
											v158 = int32(1)
											v159 = v126 + v158
											*(*uint16)(unsafe.Add(mBase, uint32(v134)+4)) = uint16(v156)
											*(*uint16)(unsafe.Add(mBase, uint32(v134)+2)) = uint16(v157)
											v163 = v121 + v158
											if v163 != v19<<(uint(base.B2i32(v79 == v112))%32) {
												v121 = v163
												v126 = v159
												continue
											} else {
												break
											}
											break
										}
										v174 = v159
									} else {
										v174 = v83
									}
									if int32(0) < v79 {
										v79 = v79 - int32(1)
										v83 = v174
										continue
									} else {
										break
									}
									break
								}
								*(*uint16)(unsafe.Add(mBase, uint32(v15)+2)) = uint16(v174)
								v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+67)))
								*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)) = uint8(v185)
								v187 = int32(0)
								v188 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
								v190 = F_ReadBufferExtended(m, v18, v187, v188, v187, v16)
								mBase = m.M
								v191 = m.ExcPending
								if v191 != 0 {
									return
								} else {
									F_LockBuffer(m, v190, int32(2))
									mBase = m.M
									v194 = m.ExcPending
									if v194 != 0 {
										return
									} else {
										v195 = F_GenericXLogStart(m, v18)
										mBase = m.M
										v196 = m.ExcPending
										if v196 != 0 {
											return
										} else {
											v198 = F_GenericXLogRegisterBuffer(m, v195, v190, int32(0))
											mBase = m.M
											v199 = m.ExcPending
											if v199 != 0 {
												return
											} else {
												v200 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+82)))
												v205 = F_PageIndexTupleOverwrite(m, v198, v200, v15, (v30+int32(7))&int32(-8))
												mBase = m.M
												v206 = m.ExcPending
												if v206 != 0 {
													return
												} else {
													if v205 == int32(0) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v226 = m.ExcPending
														if v226 != 0 {
															return
														} else {
															v227 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
															*(*int32)(unsafe.Add(mBase, uint32(v13))) = v227 + int32(4)
															F_errmsg_internal(m, int32(_a_F_RepairGraphElement_2), v13)
															mBase = m.M
															v233 = m.ExcPending
															if v233 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_RepairGraphElement_3), int32(266), int32(_a_F_RepairGraphElement_4))
																mBase = m.M
																v238 = m.ExcPending
																if v238 != 0 {
																	return
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													} else {
														F_GenericXLogFinish(m, v195)
														mBase = m.M
														v210 = m.ExcPending
														if v210 != 0 {
															return
														} else {
															F_UnlockReleaseBuffer(m, v190)
															mBase = m.M
															v212 = m.ExcPending
															if v212 != 0 {
																return
															} else {
																F_HnswUpdateNeighborsOnDisk(m, v18, v48, l1, v19, int32(1), int32(0))
																mBase = m.M
																v216 = m.ExcPending
																if v216 != 0 {
																	return
																} else {
																	m.G0 = v13 + int32(16)
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
					} else {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
						if v34 != v35 {
							v40 = int32(0)
							F_HnswInitNeighbors(m, v40, l1, v19, v40)
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return
							} else {
								v44 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l1)+64)) = uint8(v44)
								v48 = l0 + int32(24)
								F_HnswFindElementNeighbors(m, v44, l1, l2, v18, v48, v19, v17, int32(1))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return
								} else {
									v52 = int32(0)
									base.MemoryFill(m, v15, v52, int32(_a_F_RepairGraphElement_0))
									v67 = int32(2)
									*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v67)
									v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+65)))
									v79 = v71
									v83 = v52
									for {
										v98 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
										v102 = *(*int32)(unsafe.Add(mBase, uint32(v98+v79<<(uint(int32(2))%32))))
										if base.B2i32(v19 <= v52) == int32(0) {
											v112 = int32(0)
											v115 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
											v121 = v112
											v126 = v83
											for {
												v134 = v15 + int32(4) + v126*int32(6)
												if v121 < v115 {
													v141 = *(*int32)(unsafe.Add(mBase, uint32(v102+int32(8)+v121*int32(12))))
													v147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v141)+80)))
													v148 = *(*int32)(unsafe.Add(mBase, uint32(v141)+76))
													v150 = int32(base.Ui32(v148) >> (uint(int32(16)) % 32))
													*(*uint16)(unsafe.Add(mBase, uint32(v134))) = uint16(v150)
													v156 = v147
													v157 = v148
												} else {
													v152 = int32(_a_F_RepairGraphElement_1)
													*(*uint16)(unsafe.Add(mBase, uint32(v134))) = uint16(v152)
													v156 = int32(0)
													v157 = v152
												}
												v158 = int32(1)
												v159 = v126 + v158
												*(*uint16)(unsafe.Add(mBase, uint32(v134)+4)) = uint16(v156)
												*(*uint16)(unsafe.Add(mBase, uint32(v134)+2)) = uint16(v157)
												v163 = v121 + v158
												if v163 != v19<<(uint(base.B2i32(v79 == v112))%32) {
													v121 = v163
													v126 = v159
													continue
												} else {
													break
												}
												break
											}
											v174 = v159
										} else {
											v174 = v83
										}
										if int32(0) < v79 {
											v79 = v79 - int32(1)
											v83 = v174
											continue
										} else {
											break
										}
										break
									}
									*(*uint16)(unsafe.Add(mBase, uint32(v15)+2)) = uint16(v174)
									v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+67)))
									*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)) = uint8(v185)
									v187 = int32(0)
									v188 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
									v190 = F_ReadBufferExtended(m, v18, v187, v188, v187, v16)
									mBase = m.M
									v191 = m.ExcPending
									if v191 != 0 {
										return
									} else {
										F_LockBuffer(m, v190, int32(2))
										mBase = m.M
										v194 = m.ExcPending
										if v194 != 0 {
											return
										} else {
											v195 = F_GenericXLogStart(m, v18)
											mBase = m.M
											v196 = m.ExcPending
											if v196 != 0 {
												return
											} else {
												v198 = F_GenericXLogRegisterBuffer(m, v195, v190, int32(0))
												mBase = m.M
												v199 = m.ExcPending
												if v199 != 0 {
													return
												} else {
													v200 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+82)))
													v205 = F_PageIndexTupleOverwrite(m, v198, v200, v15, (v30+int32(7))&int32(-8))
													mBase = m.M
													v206 = m.ExcPending
													if v206 != 0 {
														return
													} else {
														if v205 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v226 = m.ExcPending
															if v226 != 0 {
																return
															} else {
																v227 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
																*(*int32)(unsafe.Add(mBase, uint32(v13))) = v227 + int32(4)
																F_errmsg_internal(m, int32(_a_F_RepairGraphElement_2), v13)
																mBase = m.M
																v233 = m.ExcPending
																if v233 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(_a_F_RepairGraphElement_3), int32(266), int32(_a_F_RepairGraphElement_4))
																	mBase = m.M
																	v238 = m.ExcPending
																	if v238 != 0 {
																		return
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														} else {
															F_GenericXLogFinish(m, v195)
															mBase = m.M
															v210 = m.ExcPending
															if v210 != 0 {
																return
															} else {
																F_UnlockReleaseBuffer(m, v190)
																mBase = m.M
																v212 = m.ExcPending
																if v212 != 0 {
																	return
																} else {
																	F_HnswUpdateNeighborsOnDisk(m, v18, v48, l1, v19, int32(1), int32(0))
																	mBase = m.M
																	v216 = m.ExcPending
																	if v216 != 0 {
																		return
																	} else {
																		m.G0 = v13 + int32(16)
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
						} else {
							v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+80)))
							v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+80)))
							if v37 == v38 {
								m.G0 = v13 + int32(16)
								return
							} else {
								v40 = int32(0)
								F_HnswInitNeighbors(m, v40, l1, v19, v40)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return
								} else {
									v44 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l1)+64)) = uint8(v44)
									v48 = l0 + int32(24)
									F_HnswFindElementNeighbors(m, v44, l1, l2, v18, v48, v19, v17, int32(1))
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return
									} else {
										v52 = int32(0)
										base.MemoryFill(m, v15, v52, int32(_a_F_RepairGraphElement_0))
										v67 = int32(2)
										*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v67)
										v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+65)))
										v79 = v71
										v83 = v52
										for {
											v98 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
											v102 = *(*int32)(unsafe.Add(mBase, uint32(v98+v79<<(uint(int32(2))%32))))
											if base.B2i32(v19 <= v52) == int32(0) {
												v112 = int32(0)
												v115 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
												v121 = v112
												v126 = v83
												for {
													v134 = v15 + int32(4) + v126*int32(6)
													if v121 < v115 {
														v141 = *(*int32)(unsafe.Add(mBase, uint32(v102+int32(8)+v121*int32(12))))
														v147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v141)+80)))
														v148 = *(*int32)(unsafe.Add(mBase, uint32(v141)+76))
														v150 = int32(base.Ui32(v148) >> (uint(int32(16)) % 32))
														*(*uint16)(unsafe.Add(mBase, uint32(v134))) = uint16(v150)
														v156 = v147
														v157 = v148
													} else {
														v152 = int32(_a_F_RepairGraphElement_1)
														*(*uint16)(unsafe.Add(mBase, uint32(v134))) = uint16(v152)
														v156 = int32(0)
														v157 = v152
													}
													v158 = int32(1)
													v159 = v126 + v158
													*(*uint16)(unsafe.Add(mBase, uint32(v134)+4)) = uint16(v156)
													*(*uint16)(unsafe.Add(mBase, uint32(v134)+2)) = uint16(v157)
													v163 = v121 + v158
													if v163 != v19<<(uint(base.B2i32(v79 == v112))%32) {
														v121 = v163
														v126 = v159
														continue
													} else {
														break
													}
													break
												}
												v174 = v159
											} else {
												v174 = v83
											}
											if int32(0) < v79 {
												v79 = v79 - int32(1)
												v83 = v174
												continue
											} else {
												break
											}
											break
										}
										*(*uint16)(unsafe.Add(mBase, uint32(v15)+2)) = uint16(v174)
										v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+67)))
										*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)) = uint8(v185)
										v187 = int32(0)
										v188 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
										v190 = F_ReadBufferExtended(m, v18, v187, v188, v187, v16)
										mBase = m.M
										v191 = m.ExcPending
										if v191 != 0 {
											return
										} else {
											F_LockBuffer(m, v190, int32(2))
											mBase = m.M
											v194 = m.ExcPending
											if v194 != 0 {
												return
											} else {
												v195 = F_GenericXLogStart(m, v18)
												mBase = m.M
												v196 = m.ExcPending
												if v196 != 0 {
													return
												} else {
													v198 = F_GenericXLogRegisterBuffer(m, v195, v190, int32(0))
													mBase = m.M
													v199 = m.ExcPending
													if v199 != 0 {
														return
													} else {
														v200 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+82)))
														v205 = F_PageIndexTupleOverwrite(m, v198, v200, v15, (v30+int32(7))&int32(-8))
														mBase = m.M
														v206 = m.ExcPending
														if v206 != 0 {
															return
														} else {
															if v205 == int32(0) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v226 = m.ExcPending
																if v226 != 0 {
																	return
																} else {
																	v227 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
																	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v227 + int32(4)
																	F_errmsg_internal(m, int32(_a_F_RepairGraphElement_2), v13)
																	mBase = m.M
																	v233 = m.ExcPending
																	if v233 != 0 {
																		return
																	} else {
																		F_errfinish(m, int32(_a_F_RepairGraphElement_3), int32(266), int32(_a_F_RepairGraphElement_4))
																		mBase = m.M
																		v238 = m.ExcPending
																		if v238 != 0 {
																			return
																		} else {
																			base.Wasm_trap_unreachable()
																			for {
																			}
																		}
																	}
																}
															} else {
																F_GenericXLogFinish(m, v195)
																mBase = m.M
																v210 = m.ExcPending
																if v210 != 0 {
																	return
																} else {
																	F_UnlockReleaseBuffer(m, v190)
																	mBase = m.M
																	v212 = m.ExcPending
																	if v212 != 0 {
																		return
																	} else {
																		F_HnswUpdateNeighborsOnDisk(m, v18, v48, l1, v19, int32(1), int32(0))
																		mBase = m.M
																		v216 = m.ExcPending
																		if v216 != 0 {
																			return
																		} else {
																			m.G0 = v13 + int32(16)
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
func F_ReplaceVarsFromTargetList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l2
	v23 = F_replace_rte_variables(m, l0, l1, int32(0), int32(1058), v12+int32(12), l7)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		return int32(0)
	} else {
		m.G0 = v12 + int32(32)
		return v23
	}
}
func F_ReplaceVarsFromTargetList_callback(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	v14 = F_ReplaceVarFromTargetList(m, l0, v9, v10, v11, v12, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v18 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v18
			v26 = F_query_or_expression_tree_walker_impl(m, v14, int32(1050), v6+int32(8), int32(16))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return v14
			}
		} else {
			m.G0 = v6 + int32(16)
			return v14
		}
	}
}
func F_ReportChangedGUCOptions(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReportChangedGUCOptions[0])))
	if v4 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v8 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReportChangedGUCOptions[1])))
	if v8 != int32(1) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_ReportChangedGUCOptions[2]))
	if v37 == int32(0) {
		goto L1
	} else {
		goto L12
	}
L4:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReportChangedGUCOptions[3])))
	if v13 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v23 != 0 {
		goto L3
	} else {
		goto L9
	}
L6:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_ReportChangedGUCOptions[4]))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+316))
	v21 = base.B2i32(v19 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_ReportChangedGUCOptions[3])) = uint8(v21)
	v23 = v21
	goto L8
L7:
	;
	v23 = int32(0)
	goto L8
L8:
	;
	goto L5
L9:
	;
	v25 = int32(0)
	v28 = int32(10)
	v34 = F_set_config_with_handle(m, int32(_a_F_ReportChangedGUCOptions_0), v25, int32(_a_F_ReportChangedGUCOptions_1), v25, v28, v28, v25, int32(1), v25, v25)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return
L11:
	;
	goto L3
L12:
	;
	v40 = v37
	goto L13
L13:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	F_ReportGUCOption(m, v40-int32(76))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L10
	} else {
		goto L15
	}
L14:
	;
	goto L1
L15:
	;
	v48 = v40 - int32(48)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v49 & int32(-5)
	*(*int32)(unsafe.Add(mBase, _c_F_ReportChangedGUCOptions[2])) = v42
	if v42 != 0 {
		v40 = v42
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
}
func F_RequestCheckpoint(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	v4 = F_CreateCheckPoint(m, l0|int32(4))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		F_smgrdestroyall(m)
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			return
		}
	}
}
func F_ReserveExternalFD(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_ReserveExternalFD[0]))
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_ReserveExternalFD[1]))
	if v6 <= int32(0) {
		v36 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReserveExternalFD[0])) = v36 + int32(1)
	return
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_ReserveExternalFD[2]))
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_ReserveExternalFD[3]))
	if v12+v6+v4 < v10 {
		v36 = v4
		goto L1
	} else {
		goto L3
	}
L3:
	;
	goto L4
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_ReserveExternalFD[4]))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	F_LruDelete(m, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v36 = v24
	goto L1
L6:
	;
	return
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_ReserveExternalFD[0]))
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_ReserveExternalFD[1]))
	if v26 <= int32(0) {
		v36 = v24
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_ReserveExternalFD[2]))
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_ReserveExternalFD[3]))
	if v30 <= v32+v26+v24 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	goto L5
}
func F_r_LONG_1(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_find_among_b(m, l0, int32(_a_F_r_LONG_1_0), int32(7))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v4 != int32(0))
	}
}
func F_r_SUFFIX_AN_OK(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
	return base.B2i32(v3 != int32(1))
}
func F_r_Step_5b(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v7 <= v9 {
		v47 = v2
		return v47
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+v7-int32(1)))))
		if v15 != int32(108) {
			v47 = v2
			return v47
		} else {
			v19 = v7 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v19
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v19
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
			if base.B2i32(v19 <= v9)|base.B2i32(v7 <= v24) != 0 {
				v47 = v2
				return v47
			} else {
				v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19+v11-int32(1)))))
				if v30 != int32(108) {
					v47 = v2
					return v47
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7 - int32(2)
					v37 = F_slice_del(m, l0)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						if int32(0) <= v37 {
							v43 = int32(1)
						} else {
							v43 = v37
						}
						v47 = v43
						return v47
					}
				}
			}
		}
	}
}
func F_r_VI_2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v4 <= v5 {
		v148 = v2
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7+v4-int32(1)))))
		if v11 != int32(105) {
			v148 = v2
		} else {
			v15 = v4 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v15
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v15 <= v30 {
				v138 = int32(-1)
				v145 = v138
			} else {
				v47 = int32(1)
				v48 = v15 - v47
				v50 = int32(*(*int8)(unsafe.Add(mBase, uint32(v31+v48))))
				v52 = v50 & int32(255)
				if base.B2i32(v48 == v30)|base.B2i32(int32(0) <= v50) != 0 {
					v110 = v52
					v114 = v47
				} else {
					v59 = v52 & int32(63)
					v61 = v15 - int32(2)
					v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+v61))))
					v65 = v63 << (uint(int32(6)) % 32)
					if base.B2i32(v61 != v30)&base.B2i32(base.Ui32(v63) < base.Ui32(int32(192))) == int32(0) {
						v110 = v65&int32(1984) | v59
						v114 = int32(2)
					} else {
						v78 = v65&int32(4032) | v59
						v80 = v15 - int32(3)
						v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+v80))))
						if base.B2i32(v80 != v30)&base.B2i32(base.Ui32(v82) < base.Ui32(int32(224))) == int32(0) {
							v110 = v82<<(uint(int32(12))%32)&int32(_a_F_r_VI_2_0) | v78
							v114 = int32(3)
						} else {
							v100 = int32(4)
							v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+v31-v100))))
							v110 = v82<<(uint(int32(12))%32)&int32(_a_F_r_VI_2_1) | v102&int32(7)<<(uint(int32(18))%32) | v78
							v114 = v100
						}
					}
				}
				if int32(246) < v110 {
					v145 = v114
				} else {
					v116 = v110 - int32(97)
					if v116 < int32(0) {
						v145 = v114
					} else {
						v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v116)>>(uint(int32(3))%32)))+uint32(_c_F_r_VI_2[0]))))
						if int32(base.Ui32(v122)>>(uint(v116&int32(7))%32))&int32(1) == int32(0) {
							v145 = v114
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v15 - v114
							v138 = int32(0)
							v145 = v138
						}
					}
				}
			}
			v148 = base.B2i32(v145 == int32(0))
		}
	}
	return v148
}
func F_r_VOWEL_2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v74 int32
	_ = v74
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v14 <= v13 {
		v119 = int32(-1)
	} else {
		v30 = int32(1)
		v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v15))))
		if base.Ui32(v32) < base.Ui32(int32(192)) {
			v89 = v32
			v90 = v30
		} else {
			v36 = v13 + int32(1)
			if v36 == v14 {
				v89 = v32
				v90 = v30
			} else {
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+v15))))
				v41 = v39 & int32(63)
				if base.Ui32(int32(224)) <= base.Ui32(v32) {
					v45 = v13 + int32(2)
					if v45 != v14 {
						v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+v15))))
						v57 = v55 & int32(63)
						if base.Ui32(int32(240)) <= base.Ui32(v32) {
							v61 = v13 + int32(3)
							if v61 != v14 {
								v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+v61))))
								v89 = v74&int32(63) | (v32<<(uint(int32(18))%32)&int32(_a_F_r_VOWEL_2_0) | v41<<(uint(int32(12))%32) | v57<<(uint(int32(6))%32))
								v90 = int32(4)
							} else {
								v89 = v32<<(uint(int32(12))%32)&int32(_a_F_r_VOWEL_2_1) | v41<<(uint(int32(6))%32) | v57
								v90 = int32(3)
							}
						} else {
							v89 = v32<<(uint(int32(12))%32)&int32(_a_F_r_VOWEL_2_1) | v41<<(uint(int32(6))%32) | v57
							v90 = int32(3)
						}
					} else {
						v89 = v32<<(uint(int32(6))%32)&int32(1984) | v41
						v90 = int32(2)
					}
				} else {
					v89 = v32<<(uint(int32(6))%32)&int32(1984) | v41
					v90 = int32(2)
				}
			}
		}
		if int32(117) < v89 {
			v112 = v90
		} else {
			v94 = v89 - int32(97)
			if v94 < int32(0) {
				v112 = v90
			} else {
				v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v94)>>(uint(int32(3))%32)))+uint32(_c_F_r_VOWEL_2[0]))))
				if int32(base.Ui32(v100)>>(uint(v94&int32(7))%32))&int32(1) == int32(0) {
					v112 = v90
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90 + v13
					v112 = int32(0)
				}
			}
		}
		v119 = v112
	}
	return base.B2i32(v119 == int32(0))
}
func F_r_e_ending_2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v2
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v9 <= v11 {
		v174 = v2
		return v174
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v9-int32(1)))))
		if v17 != int32(101) {
			v174 = v2
			return v174
		} else {
			v21 = v9 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v21
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v21
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
			if v9 <= v24 {
				v174 = v2
				return v174
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v39 <= v40 {
					v149 = int32(-1)
					v156 = v149
				} else {
					v57 = int32(1)
					v58 = v39 - v57
					v60 = int32(*(*int8)(unsafe.Add(mBase, uint32(v41+v58))))
					v62 = v60 & int32(255)
					if base.B2i32(v58 == v40)|base.B2i32(int32(0) <= v60) != 0 {
						v120 = v62
						v124 = v57
					} else {
						v69 = v62 & int32(63)
						v71 = v39 - int32(2)
						v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+v71))))
						v75 = v73 << (uint(int32(6)) % 32)
						if base.B2i32(v71 != v40)&base.B2i32(base.Ui32(v73) < base.Ui32(int32(192))) == int32(0) {
							v120 = v75&int32(1984) | v69
							v124 = int32(2)
						} else {
							v88 = v75&int32(4032) | v69
							v90 = v39 - int32(3)
							v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+v90))))
							if base.B2i32(v90 != v40)&base.B2i32(base.Ui32(v92) < base.Ui32(int32(224))) == int32(0) {
								v120 = v92<<(uint(int32(12))%32)&int32(_a_F_r_e_ending_2_0) | v88
								v124 = int32(3)
							} else {
								v110 = int32(4)
								v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+v41-v110))))
								v120 = v92<<(uint(int32(12))%32)&int32(_a_F_r_e_ending_2_1) | v112&int32(7)<<(uint(int32(18))%32) | v88
								v124 = v110
							}
						}
					}
					if int32(232) < v120 {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v39 - v124
						v149 = int32(0)
						v156 = v149
					} else {
						v126 = v120 - int32(97)
						if v126 < int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v39 - v124
							v149 = int32(0)
							v156 = v149
						} else {
							v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v126)>>(uint(int32(3))%32)))+uint32(_c_F_r_e_ending_2[0]))))
							if int32(base.Ui32(v132)>>(uint(v126&int32(7))%32))&int32(1) == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v39 - v124
								v149 = int32(0)
								v156 = v149
							} else {
								v156 = v124
							}
						}
					}
				}
				if v156 != 0 {
					v174 = v2
					return v174
				} else {
					v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v157 + (v21 - v26)
					v161 = F_slice_del(m, l0)
					mBase = m.M
					v164 = m.ExcPending
					if v164 != 0 {
						return int32(0)
					} else {
						if v161 < int32(0) {
							v174 = v161
							return v174
						} else {
							v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							*(*int32)(unsafe.Add(mBase, uint32(v167)+12)) = int32(1)
							v170 = F_r_undouble_3(m, l0)
							mBase = m.M
							v171 = m.ExcPending
							if v171 != 0 {
								return int32(0)
							} else {
								v174 = v170
								return v174
							}
						}
					}
				}
			}
		}
	}
}
func F_r_remove_suffix_1(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13984(m, l0, int32(_a_F_r_remove_suffix_1_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_r_stem_suffix_chain_before_ki(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v279 int32
	_ = v279
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v398 int32
	_ = v398
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v462 int32
	_ = v462
	var v470 int32
	_ = v470
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v534 int32
	_ = v534
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v6
	v8 = int32(2)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v6-v13 < v8 {
		v23 = v2
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v534
L2:
	;
	if v23 == int32(0) {
		v534 = v2
		goto L1
	} else {
		goto L6
	}
L3:
	;
	goto L2
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v19 = F_memcmp(m, v16+v6-v8, int32(_a_F_r_stem_suffix_chain_before_ki_0), v8)
	mBase = m.M
	if v19 != 0 {
		v23 = v2
		goto L3
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6 - v8
	v23 = int32(1)
	goto L3
L6:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v28 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v28 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v534 = int32(1)
	goto L1
L8:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v143 = v27 - v26
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v142 - v143
	v146 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v146 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L9:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v33 = v31 - int32(1)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v33 <= v34 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+v33))))
	switch v38 - int32(97) {
	case 0, 4:
		goto L11
	default:
		goto L8
	}
L11:
	;
	v43 = F_find_among_b(m, l0, int32(_a_F_r_stem_suffix_chain_before_ki_6), int32(4))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	if v43 == int32(0) {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v49
	v51 = F_slice_del(m, l0)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	if v51 < int32(0) {
		v534 = v51
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v55
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v58 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v58 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v98 = v55 - v57
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v98 + v99
	v102 = F_r_mark_possessives(m, l0)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L12
	} else {
		goto L30
	}
L18:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v62-int32(2) <= v61 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66+v62-int32(1)))))
	if v70 != int32(114) {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v75 = F_find_among_b(m, l0, int32(_a_F_r_stem_suffix_chain_before_ki_7), int32(2))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	if v75 == int32(0) {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v79
	v81 = F_slice_del(m, l0)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L12
	} else {
		goto L23
	}
L23:
	;
	if v81 < int32(0) {
		v534 = v81
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v87 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L12
	} else {
		goto L25
	}
L25:
	;
	if v87 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v91 + (v85 - v86)
	goto L7
L27:
	;
	goto L28
L28:
	;
	if int32(0) <= v87 {
		goto L7
	} else {
		goto L29
	}
L29:
	;
	v534 = v87
	goto L1
L30:
	;
	if v102 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v106 + v98
	goto L7
L32:
	;
	goto L33
L33:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v109
	v111 = F_slice_del(m, l0)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L12
	} else {
		goto L34
	}
L34:
	;
	if v111 < int32(0) {
		v534 = v111
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v115
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v118 = v117 - v115
	v119 = F_r_mark_lAr(m, l0)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L12
	} else {
		goto L36
	}
L36:
	;
	if v119 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v123 - v118
	goto L7
L38:
	;
	goto L39
L39:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v126
	v128 = F_slice_del(m, l0)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L12
	} else {
		goto L40
	}
L40:
	;
	if v128 < int32(0) {
		v534 = v128
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v132 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L12
	} else {
		goto L42
	}
L42:
	;
	if v132 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v136 - v118
	goto L7
L44:
	;
	goto L45
L45:
	;
	if int32(0) <= v132 {
		goto L7
	} else {
		goto L46
	}
L46:
	;
	v534 = v132
	goto L1
L47:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v523
	v525 = F_slice_del(m, l0)
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L12
	} else {
		goto L164
	}
L48:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v517
	v519 = F_slice_del(m, l0)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L12
	} else {
		goto L162
	}
L49:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v343 - v143
	v346 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v346 == int32(0) {
		v534 = v2
		goto L1
	} else {
		goto L111
	}
L50:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v151 = v149 - int32(1)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v151 <= v152 {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154+v151))))
	if v156 != int32(110) {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v161 = F_find_among_b(m, l0, int32(_a_F_r_stem_suffix_chain_before_ki_5), int32(4))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L12
	} else {
		goto L53
	}
L53:
	;
	if v161 == int32(0) {
		goto L49
	} else {
		goto L54
	}
L54:
	;
	v166 = Fn13981(m, l0, int32(110))
	mBase = m.M
	goto L55
L55:
	;
	if v166 == int32(0) {
		goto L49
	} else {
		goto L56
	}
L56:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v169
	v171 = F_slice_del(m, l0)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L12
	} else {
		goto L57
	}
L57:
	;
	if v171 < int32(0) {
		v534 = v171
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v175
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v175-int32(3) <= v178 {
		v197 = v177
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v198 = v177 - v175
	v199 = v197 - v198
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v199
	v202 = F_r_mark_possessives(m, l0)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L12
	} else {
		goto L68
	}
L60:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182+v175-int32(1)))))
	if v186 != int32(177) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	if v186 != int32(105) {
		v197 = v177
		goto L59
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v193 = F_find_among_b(m, l0, int32(_a_F_r_stem_suffix_chain_before_ki_4), int32(2))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L12
	} else {
		goto L65
	}
L64:
	;
	goto L63
L65:
	;
	if v193 != 0 {
		goto L48
	} else {
		goto L66
	}
L66:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v197 = v195
	goto L59
L67:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v329 - v198
	v332 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L12
	} else {
		goto L106
	}
L68:
	;
	if v202 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v206 - v198
	v209 = int32(0)
	v215 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v215 == v209 {
		goto L73
	} else {
		goto L74
	}
L70:
	;
	goto L71
L71:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v297
	v299 = F_slice_del(m, l0)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L12
	} else {
		goto L93
	}
L72:
	;
	if v294 == int32(0) {
		goto L67
	} else {
		goto L92
	}
L73:
	;
	v294 = int32(0)
	goto L72
L74:
	;
	goto L75
L75:
	;
	v223 = F_in_grouping_b_U(m, l0, int32(_a_F_r_stem_suffix_chain_before_ki_2), int32(105), int32(305), int32(0))
	mBase = m.M
	if v223 != 0 {
		v287 = v209
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v294 = v287
	goto L72
L77:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v226 <= v227 {
		goto L81
	} else {
		goto L82
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v279
	v287 = int32(1)
	goto L76
L79:
	;
	v279 = v236 - v224 + v243
	goto L78
L80:
	;
	v251 = v226 - v224
	v252 = v248 + v251
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v252
	if v250 < v252 {
		goto L86
	} else {
		goto L87
	}
L81:
	;
	v248 = v224
	v249 = v225
	v250 = v227
	goto L80
L82:
	;
	goto L83
L83:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225+v226-int32(1)))))
	if v232 != int32(115) {
		v248 = v224
		v249 = v225
		v250 = v227
		goto L80
	} else {
		goto L84
	}
L84:
	;
	v236 = v226 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v236
	v241 = int32(0)
	v242 = F_in_grouping_b_U(m, l0, int32(_a_F_r_stem_suffix_chain_before_ki_3), int32(97), int32(305), v241)
	mBase = m.M
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v242 == v241 {
		goto L79
	} else {
		goto L85
	}
L85:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v248 = v243
	v249 = v246
	v250 = v247
	goto L80
L86:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252+v249-int32(1)))))
	if v258 == int32(115) {
		v287 = v209
		goto L76
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v262 = F_skip_b_utf8(m, v249, v252, v250, int32(1))
	mBase = m.M
	if v262 < int32(0) {
		v287 = v209
		goto L76
	} else {
		goto L90
	}
L89:
	;
	goto L88
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v262
	v270 = F_in_grouping_b_U(m, l0, int32(_a_F_r_stem_suffix_chain_before_ki_3), int32(97), int32(305), int32(0))
	mBase = m.M
	if v270 != 0 {
		v287 = v209
		goto L76
	} else {
		goto L91
	}
L91:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v279 = v271 + v251
	goto L78
L92:
	;
	goto L71
L93:
	;
	if v299 < int32(0) {
		v534 = v299
		goto L1
	} else {
		goto L94
	}
L94:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v303
	v305 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v306 = v305 - v303
	v307 = F_r_mark_lAr(m, l0)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L12
	} else {
		goto L95
	}
L95:
	;
	if v307 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v311 - v306
	goto L7
L97:
	;
	goto L98
L98:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v314
	v316 = F_slice_del(m, l0)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L12
	} else {
		goto L99
	}
L99:
	;
	if v316 < int32(0) {
		v534 = v316
		goto L1
	} else {
		goto L100
	}
L100:
	;
	v320 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L12
	} else {
		goto L101
	}
L101:
	;
	if v320 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v324 - v306
	goto L7
L103:
	;
	goto L104
L104:
	;
	if int32(0) <= v320 {
		goto L7
	} else {
		goto L105
	}
L105:
	;
	v534 = v320
	goto L1
L106:
	;
	if v332 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v336 + (v175 - v177)
	goto L7
L108:
	;
	goto L109
L109:
	;
	if int32(0) <= v332 {
		goto L7
	} else {
		goto L110
	}
L110:
	;
	v534 = v332
	goto L1
L111:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v350-int32(2) <= v349 {
		v534 = v2
		goto L1
	} else {
		goto L112
	}
L112:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354+v350-int32(1)))))
	switch v358 - int32(97) {
	case 0, 4:
		goto L113
	default:
		v534 = v2
		goto L1
	}
L113:
	;
	v363 = F_find_among_b(m, l0, int32(_a_F_r_stem_suffix_chain_before_ki_1), int32(2))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L12
	} else {
		goto L114
	}
L114:
	;
	if v363 == int32(0) {
		v534 = v2
		goto L1
	} else {
		goto L115
	}
L115:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v369-int32(3) <= v368 {
		v388 = v367
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v389 = v367 - v369
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v388 - v389
	v392 = int32(0)
	v398 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v398 == v392 {
		goto L125
	} else {
		goto L126
	}
L117:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373+v369-int32(1)))))
	if v377 != int32(177) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	if v377 != int32(105) {
		v388 = v367
		goto L116
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v384 = F_find_among_b(m, l0, int32(_a_F_r_stem_suffix_chain_before_ki_4), int32(2))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L12
	} else {
		goto L122
	}
L121:
	;
	goto L120
L122:
	;
	if v384 != 0 {
		goto L47
	} else {
		goto L123
	}
L123:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v388 = v386
	goto L116
L124:
	;
	if v477 != 0 {
		goto L144
	} else {
		goto L145
	}
L125:
	;
	v477 = int32(0)
	goto L124
L126:
	;
	goto L127
L127:
	;
	v406 = F_in_grouping_b_U(m, l0, int32(_a_F_r_stem_suffix_chain_before_ki_2), int32(105), int32(305), int32(0))
	mBase = m.M
	if v406 != 0 {
		v470 = v392
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v477 = v470
	goto L124
L129:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v409 <= v410 {
		goto L133
	} else {
		goto L134
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v462
	v470 = int32(1)
	goto L128
L131:
	;
	v462 = v419 - v407 + v426
	goto L130
L132:
	;
	v434 = v409 - v407
	v435 = v431 + v434
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v435
	if v433 < v435 {
		goto L138
	} else {
		goto L139
	}
L133:
	;
	v431 = v407
	v432 = v408
	v433 = v410
	goto L132
L134:
	;
	goto L135
L135:
	;
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408+v409-int32(1)))))
	if v415 != int32(115) {
		v431 = v407
		v432 = v408
		v433 = v410
		goto L132
	} else {
		goto L136
	}
L136:
	;
	v419 = v409 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v419
	v424 = int32(0)
	v425 = F_in_grouping_b_U(m, l0, int32(_a_F_r_stem_suffix_chain_before_ki_3), int32(97), int32(305), v424)
	mBase = m.M
	v426 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v425 == v424 {
		goto L131
	} else {
		goto L137
	}
L137:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v431 = v426
	v432 = v429
	v433 = v430
	goto L132
L138:
	;
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v435+v432-int32(1)))))
	if v441 == int32(115) {
		v470 = v392
		goto L128
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	v445 = F_skip_b_utf8(m, v432, v435, v433, int32(1))
	mBase = m.M
	if v445 < int32(0) {
		v470 = v392
		goto L128
	} else {
		goto L142
	}
L141:
	;
	goto L140
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v445
	v453 = F_in_grouping_b_U(m, l0, int32(_a_F_r_stem_suffix_chain_before_ki_3), int32(97), int32(305), int32(0))
	mBase = m.M
	if v453 != 0 {
		v470 = v392
		goto L128
	} else {
		goto L143
	}
L143:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v462 = v454 + v434
	goto L130
L144:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v478
	v480 = F_slice_del(m, l0)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L12
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v510 - v389
	v513 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L12
	} else {
		goto L160
	}
L147:
	;
	if v480 < int32(0) {
		v534 = v480
		goto L1
	} else {
		goto L148
	}
L148:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v484
	v486 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v487 = v486 - v484
	v488 = F_r_mark_lAr(m, l0)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L12
	} else {
		goto L149
	}
L149:
	;
	if v488 == int32(0) {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v492 - v487
	goto L7
L151:
	;
	goto L152
L152:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v495
	v497 = F_slice_del(m, l0)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L12
	} else {
		goto L153
	}
L153:
	;
	if v497 < int32(0) {
		v534 = v497
		goto L1
	} else {
		goto L154
	}
L154:
	;
	v501 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L12
	} else {
		goto L155
	}
L155:
	;
	if v501 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v505 - v487
	goto L7
L157:
	;
	goto L158
L158:
	;
	if int32(0) <= v501 {
		goto L7
	} else {
		goto L159
	}
L159:
	;
	v534 = v501
	goto L1
L160:
	;
	if v513 <= int32(0) {
		v534 = v513
		goto L1
	} else {
		goto L161
	}
L161:
	;
	goto L7
L162:
	;
	if int32(0) <= v519 {
		goto L7
	} else {
		goto L163
	}
L163:
	;
	v534 = v519
	goto L1
L164:
	;
	if v525 < int32(0) {
		v534 = v525
		goto L1
	} else {
		goto L165
	}
L165:
	;
	goto L7
}
func F_r_undouble_3(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = v5 - int32(1)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v7 <= v8 {
		v108 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v108
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v7))))
	if base.B2i32(v12&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v12)%32)&int32(_a_F_r_undouble_3_0) == int32(0)) != 0 {
		v108 = v2
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v27 = F_find_among_b(m, l0, int32(_a_F_r_undouble_3_1), int32(3))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	if v27 == int32(0) {
		v108 = v2
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v35 = v33 + (v5 - v24)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v35
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L9
L7:
	;
	if v91 < int32(0) {
		v108 = v2
		goto L1
	} else {
		goto L26
	}
L9:
	;
	goto L10
L10:
	;
	goto L11
L11:
	;
	v46 = v35
	v48 = int32(1)
	goto L14
L13:
	;
	v91 = v73
	goto L7
L14:
	;
	if v46 <= v39 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L13
L16:
	;
	v91 = int32(-1)
	goto L7
L17:
	;
	goto L18
L18:
	;
	v53 = v46 - int32(1)
	v55 = int32(*(*int8)(unsafe.Add(mBase, uint32(v38+v53))))
	if base.B2i32(int32(0) <= v55)|base.B2i32(v53 <= v39) != 0 {
		v73 = v53
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v77 = int32(1)
	if v77 < v48 {
		v46 = v73
		v48 = v48 - v77
		goto L14
	} else {
		goto L25
	}
L20:
	;
	v61 = v53
	goto L21
L21:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+v61))))
	if base.Ui32(int32(191)) < base.Ui32(v66) {
		v73 = v61
		goto L19
	} else {
		goto L23
	}
L22:
	;
	v73 = v39
	goto L19
L23:
	;
	v70 = v61 - int32(1)
	if v39 < v70 {
		v61 = v70
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	goto L15
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v91
	v97 = F_slice_del(m, l0)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	if int32(0) <= v97 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v104 = int32(1)
	goto L30
L29:
	;
	v104 = v97 >> (uint(int32(31)) % 32) & v97
	goto L30
L30:
	;
	v108 = v104
	goto L1
}
func F_rangeTableEntry_used(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	v3 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = l1
	v15 = F_query_or_expression_tree_walker_impl(m, l0, int32(1051), v6+int32(8), v3)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		m.G0 = v6 + int32(16)
		return v15
	}
}
func F_rboolop(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_DirectFunctionCall2Coll(m, int32(_a_F_rboolop_0), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_read_gucstate(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(v6) < base.Ui32(l1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = v6
	goto L5
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L9
	} else {
		goto L13
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v14
	return v6
L5:
	;
	v14 = v10 + int32(1)
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v15 == int32(0) {
		goto L4
	} else {
		goto L7
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	if v14 != l1 {
		v10 = v14
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	return int32(0)
L10:
	;
	F_errmsg_internal(m, int32(_a_F_read_gucstate_0), int32(0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(_a_F_read_gucstate_1), int32(_a_F_read_gucstate_2), int32(_a_F_read_gucstate_3))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L13:
	;
	F_errmsg_internal(m, int32(_a_F_read_gucstate_4), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(_a_F_read_gucstate_1), int32(_a_F_read_gucstate_5), int32(_a_F_read_gucstate_3))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_readdir(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int64
	_ = v33
	var v38 int32
	_ = v38
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v5 <= v4 {
		v7 = int32(0)
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v12 = m.Env.X__syscall_getdents64(m, v8, l0+int32(24), int32(2048))
		mBase = m.M
		if v12 <= v7 {
			if base.B2i32(v12 == int32(0))|base.B2i32(v12 == int32(-44)) != 0 {
				v38 = v7
				return v38
			} else {
				v21 = int32(0)
				*(*int32)(unsafe.Add(mBase, _c_F_readdir[0])) = v21 - v12
				return v21
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v12
			v28 = v7
			v29 = l0 + v28
			v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+40)))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v28 + v30
			v33 = *(*int64)(unsafe.Add(mBase, uint32(v29)+32))
			*(*int64)(unsafe.Add(mBase, uint32(l0))) = v33
			v38 = v29 + int32(24)
			return v38
		}
	} else {
		v28 = v4
		v29 = l0 + v28
		v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+40)))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v28 + v30
		v33 = *(*int64)(unsafe.Add(mBase, uint32(v29)+32))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v33
		v38 = v29 + int32(24)
		return v38
	}
}
func F_record_larger(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v4 = F_record_cmp(m, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if int32(0) < v4 {
			v10 = int32(20)
		} else {
			v10 = int32(28)
		}
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0+v10)))
		return v12
	}
}
func F_record_le(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_record_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v2 <= int32(0))
	}
}
func F_record_ne(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_record_eq(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2 ^ int32(1)
	}
}
func F_reduce_expanded_ranges(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v45 int32
	_ = v45
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
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v138 int32
	_ = v138
	var v150 int32
	_ = v150
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v18 = base.I32_div_s(l3, int32(2))
	v20 = l1 - int32(1)
	if v18 <= v20 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = l5
	v26 = F_palloc(m, l3<<(uint(int32(2))%32))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v211 = l1
	goto L3
L3:
	;
	m.G0 = v15 + int32(16)
	return v211
L4:
	;
	return int32(0)
L5:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v30
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0+v20*int32(12))+4))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v35
	if l3 <= int32(3) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v170 = int32(0)
	goto L23
L7:
	;
	F_qsort_arg(m, v26, int32(2), int32(4), int32(21), v15+int32(8))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v47 = int32(0)
	v48 = int32(2)
	if int32(6) <= l3 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v165 = int32(1)
	goto L6
L11:
	;
	F_qsort_arg(m, v26, v138, int32(4), int32(21), v15+int32(8))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L4
	} else {
		goto L22
	}
L12:
	;
	v51 = int32(2)
	if v18 <= v51 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v109 = v47
	v110 = v48
	goto L14
L14:
	;
	v117 = int32(2)
	v119 = v26 + v110<<(uint(v117)%32)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l2+v109<<(uint(int32(4))%32))))
	v126 = l0 + v123*int32(12)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v119))) = v127
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v126)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v119)+4)) = v129
	v138 = v110 + v117
	goto L11
L15:
	;
	v54 = v51
	goto L17
L16:
	;
	v54 = v18
	goto L17
L17:
	;
	v55 = int32(1)
	v56 = v54 - v55
	v65 = int32(0)
	v66 = v47
	v67 = v48
	goto L18
L18:
	;
	v74 = int32(2)
	v76 = v26 + v67<<(uint(v74)%32)
	v77 = int32(4)
	v79 = l2 + v66<<(uint(v77)%32)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v81 = int32(12)
	v83 = l0 + v80*v81
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v76))) = v84
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v76)+4)) = v86
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
	v91 = l0 + v88*v81
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v76)+8)) = v92
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v76)+12)) = v94
	v97 = v66 + v74
	v99 = v67 + v77
	v101 = v65 + v74
	if v101 != v56&int32(-2) {
		v65 = v101
		v66 = v97
		v67 = v99
		goto L18
	} else {
		goto L20
	}
L19:
	;
	if v56&v55 == int32(0) {
		v138 = v99
		goto L11
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	v109 = v97
	v110 = v99
	goto L14
L22:
	;
	v165 = int32(base.Ui32(v138) >> (uint(int32(1)) % 32))
	goto L6
L23:
	;
	v181 = l0 + v170*int32(12)
	v184 = v26 + v170<<(uint(int32(3))%32)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	*(*int32)(unsafe.Add(mBase, uint32(v181))) = v185
	v188 = v184 + int32(4)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	*(*int32)(unsafe.Add(mBase, uint32(v181)+4)) = v189
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	v194 = F_FunctionCall2Coll(m, v191, v192, v193, v189)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L4
	} else {
		goto L25
	}
L24:
	;
	v211 = v165
	goto L3
L25:
	;
	if v194 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v203 = int32(1)
	goto L28
L27:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	v201 = F_FunctionCall2Coll(m, v197, v198, v199, v200)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L4
	} else {
		goto L29
	}
L28:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v181)+8)) = uint8(base.B2i32(v203 == int32(0)))
	v208 = v170 + int32(1)
	if v208 != v165 {
		v170 = v208
		goto L23
	} else {
		goto L30
	}
L29:
	;
	v203 = v201
	goto L28
L30:
	;
	goto L24
}
func F_reduce_outer_joins_pass1(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = F_palloc(m, int32(12))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v15 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v15
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+4)) = uint8(v15)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v15
	if l0 == v15 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L27
	}
L4:
	;
	m.G0 = v8 + int32(16)
	return v11
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v23 - int32(63) {
	case 0:
		goto L6
	case 1:
		goto L7
	case 2:
		goto L8
	default:
		goto L3
	}
L6:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v101 = F_bms_make_singleton(m, v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L26
	}
L7:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(1)<<(uint(v62)%32)&int32(174) != 0 {
		goto L17
	} else {
		goto L18
	}
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v26 == int32(0) {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v29 <= int32(0) {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v36 = int32(0)
	goto L11
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37+v36<<(uint(int32(2))%32))))
	v42 = F_reduce_outer_joins_pass1(m, v41)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	goto L4
L13:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v46 = F_bms_add_members(m, v44, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v46
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+4)))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+4)))
	v51 = v49 | v50
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+4)) = uint8(v51)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v54 = F_lappend(m, v53, v42)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v54
	v58 = v36 + int32(1)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v58 < v59 {
		v36 = v58
		goto L11
	} else {
		goto L16
	}
L16:
	;
	goto L12
L17:
	;
	v66 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+4)) = uint8(v66)
	goto L19
L18:
	;
	goto L19
L19:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v69 = F_reduce_outer_joins_pass1(m, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v73 = F_bms_add_members(m, v71, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v73
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+4)))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+4)))
	v78 = v76 | v77
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+4)) = uint8(v78)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v81 = F_lappend(m, v80, v69)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v81
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v85 = F_reduce_outer_joins_pass1(m, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v89 = F_bms_add_members(m, v87, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v89
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+4)))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+4)))
	v94 = v92 | v93
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+4)) = uint8(v94)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v97 = F_lappend(m, v96, v85)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v97
	goto L4
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v101
	goto L4
L27:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v117
	F_errmsg_internal(m, int32(_a_F_reduce_outer_joins_pass1_0), v8)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_reduce_outer_joins_pass1_1), int32(3232), int32(_a_F_reduce_outer_joins_pass1_2))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_regclassin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
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
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
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
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v12 == int32(45) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L30
	} else {
		goto L46
	}
L2:
	;
	m.G0 = v8 + int32(16)
	return v158
L3:
	;
	v158 = int32(0)
	goto L2
L4:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_regclassin[0]))
	if v120 == int32(0) {
		goto L1
	} else {
		goto L32
	}
L5:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v15 == int32(0) {
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	if base.Ui32(int32(9)) < base.Ui32((v12-int32(48))&int32(255)) {
		goto L4
	} else {
		goto L9
	}
L8:
	;
	goto L4
L9:
	;
	v24 = int32(_a_F_regclassin_0)
	v28 = m.G0
	v30 = v28 - int32(32)
	v31 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v30)+24)) = v31
	*(*int64)(unsafe.Add(mBase, uint32(v30)+16)) = v31
	*(*int64)(unsafe.Add(mBase, uint32(v30)+8)) = v31
	*(*int64)(unsafe.Add(mBase, uint32(v30))) = v31
	v39 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_regclassin[1])))
	if v39 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v108 = F_strlen(m, v11)
	mBase = m.M
	if v107 != v108 {
		goto L4
	} else {
		goto L29
	}
L11:
	;
	v107 = int32(0)
	goto L10
L12:
	;
	goto L13
L13:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_regclassin[2])))
	if v43 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v47 = v11
	goto L17
L15:
	;
	goto L16
L16:
	;
	v57 = v24
	v58 = v39
	goto L20
L17:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	if v53 == v39 {
		v47 = v47 + int32(1)
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v107 = v47 - v11
	goto L10
L19:
	;
	goto L18
L20:
	;
	v65 = v30 + int32(base.Ui32(v58)>>(uint(int32(3))%32))&int32(28)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v67 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v66 | v67<<(uint(v58)%32)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+1)))
	if v71 != 0 {
		v57 = v57 + v67
		v58 = v71
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v74 == int32(0) {
		v97 = v11
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L21
L23:
	;
	v107 = v97 - v11
	goto L10
L24:
	;
	v78 = v11
	v79 = v74
	goto L25
L25:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v30+int32(base.Ui32(v79)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v87)>>(uint(v79)%32))&int32(1) == int32(0) {
		v97 = v78
		goto L23
	} else {
		goto L27
	}
L26:
	;
	v97 = v95
	goto L23
L27:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+1)))
	v95 = v78 + int32(1)
	if v93 != 0 {
		v78 = v95
		v79 = v93
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v114 = F_DirectInputFunctionCallSafe(m, int32(547), v11, int32(-1), v10, v8+int32(12))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	return int32(0)
L31:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v158 = v118
	goto L2
L32:
	;
	v123 = F_stringToQualifiedNameList(m, v11, v10)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	if v123 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v127 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v127)
	goto L3
L35:
	;
	goto L36
L36:
	;
	v129 = F_makeRangeVarFromNameList(m, v123)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L30
	} else {
		goto L37
	}
L37:
	;
	v131 = int32(0)
	v135 = F_RangeVarGetRelidExtended(m, v129, v131, int32(1), v131, v131)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L30
	} else {
		goto L38
	}
L38:
	;
	if v135 != 0 {
		v158 = v135
		goto L2
	} else {
		goto L39
	}
L39:
	;
	v137 = F_errsave_start(m, v10)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L30
	} else {
		goto L40
	}
L40:
	;
	if v137 == int32(0) {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	F_errcode(m, int32(16908420))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L30
	} else {
		goto L42
	}
L42:
	;
	v144 = F_NameListToString(m, v123)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L30
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v144
	F_errmsg(m, int32(_a_F_regclassin_1), v8)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L30
	} else {
		goto L44
	}
L44:
	;
	F_errsave_finish(m, v10, int32(_a_F_regclassin_2), int32(914), int32(_a_F_regclassin_3))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L30
	} else {
		goto L45
	}
L45:
	;
	goto L3
L46:
	;
	F_errmsg_internal(m, int32(_a_F_regclassin_4), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L30
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_regclassin_2), int32(897), int32(_a_F_regclassin_3))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L30
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_regcollationout(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v10 == int32(0) {
		v14 = F_pstrdup(m, int32(_a_F_regcollationout_0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v53 = v14
			m.G0 = v8 + int32(16)
			return v53
		}
	} else {
		v19 = F_SearchSysCache1(m, int32(16), v10)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			if v19 != 0 {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+22)))
				v23 = v21 + v22
				v25 = v23 + int32(4)
				v27 = *(*int32)(unsafe.Add(mBase, _c_F_regcollationout[0]))
				if v27 == int32(0) {
					v30 = F_pstrdup(m, v25)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						F_ReleaseCatCache(m, v19)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							v53 = v30
							m.G0 = v8 + int32(16)
							return v53
						}
					}
				} else {
					v34 = F_CollationIsVisible(m, v10)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						if v34 != 0 {
							v40 = int32(0)
							v41 = F_quote_qualified_identifier(m, v40, v25)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								F_ReleaseCatCache(m, v19)
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return int32(0)
								} else {
									v53 = v41
									m.G0 = v8 + int32(16)
									return v53
								}
							}
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v23)+68))
							v38 = F_get_namespace_name(m, v37)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								v40 = v38
								v41 = F_quote_qualified_identifier(m, v40, v25)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									F_ReleaseCatCache(m, v19)
									mBase = m.M
									v44 = m.ExcPending
									if v44 != 0 {
										return int32(0)
									} else {
										v53 = v41
										m.G0 = v8 + int32(16)
										return v53
									}
								}
							}
						}
					}
				}
			} else {
				v46 = F_palloc(m, int32(64))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
					v51 = F_pg_snprintf(m, v46, int32(64), int32(_a_F_regcollationout_1), v8)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						v53 = v46
						m.G0 = v8 + int32(16)
						return v53
					}
				}
			}
		}
	}
}
func F_register_ENR(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = F_lappend(m, v3, l1)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v4
		return
	}
}
func F_regoperatorin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	v5 = m.G0
	v7 = v5 - int32(432)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if base.Ui32(int32(9)) < base.Ui32((v11-int32(48))&int32(255)) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L25
	} else {
		goto L54
	}
L2:
	;
	m.G0 = v7 + int32(432)
	return v198
L3:
	;
	v114 = *(*int32)(unsafe.Add(mBase, _c_F_regoperatorin[0]))
	if v114 == int32(0) {
		goto L1
	} else {
		goto L27
	}
L4:
	;
	v18 = int32(_a_F_regoperatorin_0)
	v22 = m.G0
	v24 = v22 - int32(32)
	v25 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v24)+24)) = v25
	*(*int64)(unsafe.Add(mBase, uint32(v24)+16)) = v25
	*(*int64)(unsafe.Add(mBase, uint32(v24)+8)) = v25
	*(*int64)(unsafe.Add(mBase, uint32(v24))) = v25
	v33 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_regoperatorin[1])))
	if v33 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v102 = F_strlen(m, v10)
	mBase = m.M
	if v101 != v102 {
		goto L3
	} else {
		goto L24
	}
L6:
	;
	v101 = int32(0)
	goto L5
L7:
	;
	goto L8
L8:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_regoperatorin[2])))
	if v37 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v41 = v10
	goto L12
L10:
	;
	goto L11
L11:
	;
	v51 = v18
	v52 = v33
	goto L15
L12:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v47 == v33 {
		v41 = v41 + int32(1)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v101 = v41 - v10
	goto L5
L14:
	;
	goto L13
L15:
	;
	v59 = v24 + int32(base.Ui32(v52)>>(uint(int32(3))%32))&int32(28)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v61 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v59))) = v60 | v61<<(uint(v52)%32)
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
	if v65 != 0 {
		v51 = v51 + v61
		v52 = v65
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v68 == int32(0) {
		v91 = v10
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	v101 = v91 - v10
	goto L5
L19:
	;
	v72 = v10
	v73 = v68
	goto L20
L20:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(base.Ui32(v73)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v81)>>(uint(v73)%32))&int32(1) == int32(0) {
		v91 = v72
		goto L18
	} else {
		goto L22
	}
L21:
	;
	v91 = v89
	goto L18
L22:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+1)))
	v89 = v72 + int32(1)
	if v87 != 0 {
		v72 = v89
		v73 = v87
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v108 = F_DirectInputFunctionCallSafe(m, int32(547), v10, int32(-1), v9, v7+int32(16))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	return int32(0)
L26:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	v198 = v112
	goto L2
L27:
	;
	v124 = F_parseNameAndArgTypes(m, v10, int32(1), v7+int32(428), v7+int32(424), v7+int32(16), v9)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	if v124 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v128 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v128)
	v198 = int32(0)
	goto L2
L30:
	;
	goto L31
L31:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v7)+424))
	switch v131 - int32(1) {
	case 0:
		goto L34
	case 1:
		goto L32
	default:
		goto L33
	}
L32:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v7)+428))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	v179 = F_OpernameGetOprid(m, v176, v177, v178)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L25
	} else {
		goto L47
	}
L33:
	;
	v155 = int32(0)
	v156 = F_errsave_start(m, v9)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L25
	} else {
		goto L41
	}
L34:
	;
	v134 = int32(0)
	v135 = F_errsave_start(m, v9)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L25
	} else {
		goto L35
	}
L35:
	;
	if v135 == int32(0) {
		v198 = v134
		goto L2
	} else {
		goto L36
	}
L36:
	;
	F_errcode(m, int32(33685636))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L25
	} else {
		goto L37
	}
L37:
	;
	F_errmsg(m, int32(_a_F_regoperatorin_1), int32(0))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L25
	} else {
		goto L38
	}
L38:
	;
	F_errhint(m, int32(_a_F_regoperatorin_2), int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L25
	} else {
		goto L39
	}
L39:
	;
	F_errsave_finish(m, v9, int32(_a_F_regoperatorin_3), int32(671), int32(_a_F_regoperatorin_4))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L25
	} else {
		goto L40
	}
L40:
	;
	v198 = v134
	goto L2
L41:
	;
	if v156 == int32(0) {
		v198 = v155
		goto L2
	} else {
		goto L42
	}
L42:
	;
	F_errcode(m, int32(50856197))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L25
	} else {
		goto L43
	}
L43:
	;
	F_errmsg(m, int32(_a_F_regoperatorin_5), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L25
	} else {
		goto L44
	}
L44:
	;
	F_errhint(m, int32(_a_F_regoperatorin_6), int32(0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L25
	} else {
		goto L45
	}
L45:
	;
	F_errsave_finish(m, v9, int32(_a_F_regoperatorin_3), int32(676), int32(_a_F_regoperatorin_4))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L25
	} else {
		goto L46
	}
L46:
	;
	v198 = v155
	goto L2
L47:
	;
	if v179 != 0 {
		v198 = v179
		goto L2
	} else {
		goto L48
	}
L48:
	;
	v181 = int32(0)
	v182 = F_errsave_start(m, v9)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L25
	} else {
		goto L49
	}
L49:
	;
	if v182 == int32(0) {
		v198 = v181
		goto L2
	} else {
		goto L50
	}
L50:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L25
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v10
	F_errmsg(m, int32(_a_F_regoperatorin_7), v7)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L25
	} else {
		goto L52
	}
L52:
	;
	F_errsave_finish(m, v9, int32(_a_F_regoperatorin_3), int32(683), int32(_a_F_regoperatorin_4))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L25
	} else {
		goto L53
	}
L53:
	;
	v198 = v181
	goto L2
L54:
	;
	F_errmsg_internal(m, int32(_a_F_regoperatorin_8), int32(0))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L25
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_regoperatorin_3), int32(654), int32(_a_F_regoperatorin_4))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L25
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_regoperin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
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
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if base.Ui32(int32(9)) < base.Ui32((v12-int32(48))&int32(255)) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L25
	} else {
		goto L49
	}
L2:
	;
	m.G0 = v8 + int32(32)
	return v168
L3:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_regoperin[0]))
	if v115 == int32(0) {
		goto L1
	} else {
		goto L27
	}
L4:
	;
	v19 = int32(_a_F_regoperin_0)
	v23 = m.G0
	v25 = v23 - int32(32)
	v26 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v25)+24)) = v26
	*(*int64)(unsafe.Add(mBase, uint32(v25)+16)) = v26
	*(*int64)(unsafe.Add(mBase, uint32(v25)+8)) = v26
	*(*int64)(unsafe.Add(mBase, uint32(v25))) = v26
	v34 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_regoperin[1])))
	if v34 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v103 = F_strlen(m, v11)
	mBase = m.M
	if v102 != v103 {
		goto L3
	} else {
		goto L24
	}
L6:
	;
	v102 = int32(0)
	goto L5
L7:
	;
	goto L8
L8:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_regoperin[2])))
	if v38 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v42 = v11
	goto L12
L10:
	;
	goto L11
L11:
	;
	v52 = v19
	v53 = v34
	goto L15
L12:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v48 == v34 {
		v42 = v42 + int32(1)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v102 = v42 - v11
	goto L5
L14:
	;
	goto L13
L15:
	;
	v60 = v25 + int32(base.Ui32(v53)>>(uint(int32(3))%32))&int32(28)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v62 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = v61 | v62<<(uint(v53)%32)
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+1)))
	if v66 != 0 {
		v52 = v52 + v62
		v53 = v66
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v69 == int32(0) {
		v92 = v11
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	v102 = v92 - v11
	goto L5
L19:
	;
	v73 = v11
	v74 = v69
	goto L20
L20:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v25+int32(base.Ui32(v74)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v82)>>(uint(v74)%32))&int32(1) == int32(0) {
		v92 = v73
		goto L18
	} else {
		goto L22
	}
L21:
	;
	v92 = v90
	goto L18
L22:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+1)))
	v90 = v73 + int32(1)
	if v88 != 0 {
		v73 = v90
		v74 = v88
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v109 = F_DirectInputFunctionCallSafe(m, int32(547), v11, int32(-1), v10, v8+int32(28))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	return int32(0)
L26:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
	v168 = v113
	goto L2
L27:
	;
	v118 = F_stringToQualifiedNameList(m, v11, v10)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	if v118 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v122 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v122)
	v168 = int32(0)
	goto L2
L30:
	;
	goto L31
L31:
	;
	v125 = int32(0)
	v128 = F_OpernameGetCandidates(m, v118, v125, int32(1))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L25
	} else {
		goto L32
	}
L32:
	;
	if v128 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v132 = F_errsave_start(m, v10)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L25
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	if v148 != 0 {
		goto L41
	} else {
		goto L42
	}
L36:
	;
	if v132 == int32(0) {
		v168 = v125
		goto L2
	} else {
		goto L37
	}
L37:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L25
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v11
	F_errmsg(m, int32(_a_F_regoperin_1), v8)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L25
	} else {
		goto L39
	}
L39:
	;
	F_errsave_finish(m, v10, int32(_a_F_regoperin_2), int32(509), int32(_a_F_regoperin_3))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L25
	} else {
		goto L40
	}
L40:
	;
	v168 = v125
	goto L2
L41:
	;
	v149 = F_errsave_start(m, v10)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L25
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
	v168 = v167
	goto L2
L44:
	;
	if v149 == int32(0) {
		v168 = v125
		goto L2
	} else {
		goto L45
	}
L45:
	;
	F_errcode(m, int32(84439172))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L25
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v11
	F_errmsg(m, int32(_a_F_regoperin_4), v8+int32(16))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L25
	} else {
		goto L47
	}
L47:
	;
	F_errsave_finish(m, v10, int32(_a_F_regoperin_2), int32(514), int32(_a_F_regoperin_3))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L25
	} else {
		goto L48
	}
L48:
	;
	v168 = v125
	goto L2
L49:
	;
	F_errmsg_internal(m, int32(_a_F_regoperin_5), int32(0))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L25
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_regoperin_2), int32(494), int32(_a_F_regoperin_3))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L25
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_regprocedureout(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v2 == int32(0) {
		v6 = F_pstrdup(m, int32(_a_F_regprocedureout_0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v6
		}
	} else {
		v12 = F_format_procedure_extended(m, v2, int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			return v12
		}
	}
}
func F_regprocin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v12 == int32(45) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L29
	} else {
		goto L53
	}
L2:
	;
	m.G0 = v8 + int32(32)
	return v176
L3:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_regprocin[0]))
	if v119 == int32(0) {
		goto L1
	} else {
		goto L31
	}
L4:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v15 != 0 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if base.Ui32(int32(9)) < base.Ui32((v12-int32(48))&int32(255)) {
		goto L3
	} else {
		goto L8
	}
L7:
	;
	v176 = int32(0)
	goto L2
L8:
	;
	v23 = int32(_a_F_regprocin_0)
	v27 = m.G0
	v29 = v27 - int32(32)
	v30 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+24)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29)+16)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29))) = v30
	v38 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_regprocin[1])))
	if v38 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v107 = F_strlen(m, v11)
	mBase = m.M
	if v106 != v107 {
		goto L3
	} else {
		goto L28
	}
L10:
	;
	v106 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_regprocin[2])))
	if v42 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v46 = v11
	goto L16
L14:
	;
	goto L15
L15:
	;
	v56 = v23
	v57 = v38
	goto L19
L16:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v52 == v38 {
		v46 = v46 + int32(1)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v106 = v46 - v11
	goto L9
L18:
	;
	goto L17
L19:
	;
	v64 = v29 + int32(base.Ui32(v57)>>(uint(int32(3))%32))&int32(28)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v66 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v65 | v66<<(uint(v57)%32)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	if v70 != 0 {
		v56 = v56 + v66
		v57 = v70
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v73 == int32(0) {
		v96 = v11
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v106 = v96 - v11
	goto L9
L23:
	;
	v77 = v11
	v78 = v73
	goto L24
L24:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(base.Ui32(v78)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v86)>>(uint(v78)%32))&int32(1) == int32(0) {
		v96 = v77
		goto L22
	} else {
		goto L26
	}
L25:
	;
	v96 = v94
	goto L22
L26:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+1)))
	v94 = v77 + int32(1)
	if v92 != 0 {
		v77 = v94
		v78 = v92
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v113 = F_DirectInputFunctionCallSafe(m, int32(547), v11, int32(-1), v10, v8+int32(28))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	return int32(0)
L30:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
	v176 = v117
	goto L2
L31:
	;
	v122 = F_stringToQualifiedNameList(m, v11, v10)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	if v122 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v126 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v126)
	v176 = int32(0)
	goto L2
L34:
	;
	goto L35
L35:
	;
	v129 = int32(0)
	v136 = F_FuncnameGetCandidates(m, v122, int32(-1), v129, v129, v129, v129, int32(1))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L29
	} else {
		goto L36
	}
L36:
	;
	if v136 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v140 = F_errsave_start(m, v10)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L29
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	if v156 != 0 {
		goto L45
	} else {
		goto L46
	}
L40:
	;
	if v140 == int32(0) {
		v176 = v129
		goto L2
	} else {
		goto L41
	}
L41:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L29
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v11
	F_errmsg(m, int32(_a_F_regprocin_1), v8)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L29
	} else {
		goto L43
	}
L43:
	;
	F_errsave_finish(m, v10, int32(_a_F_regprocin_2), int32(100), int32(_a_F_regprocin_3))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L29
	} else {
		goto L44
	}
L44:
	;
	v176 = v129
	goto L2
L45:
	;
	v157 = F_errsave_start(m, v10)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L29
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v136)+8))
	v176 = v175
	goto L2
L48:
	;
	if v157 == int32(0) {
		v176 = v129
		goto L2
	} else {
		goto L49
	}
L49:
	;
	F_errcode(m, int32(84439172))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L29
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v11
	F_errmsg(m, int32(_a_F_regprocin_4), v8+int32(16))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L29
	} else {
		goto L51
	}
L51:
	;
	F_errsave_finish(m, v10, int32(_a_F_regprocin_2), int32(105), int32(_a_F_regprocin_3))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L29
	} else {
		goto L52
	}
L52:
	;
	v176 = v129
	goto L2
L53:
	;
	F_errmsg_internal(m, int32(_a_F_regprocin_5), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L29
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_regprocin_2), int32(85), int32(_a_F_regprocin_3))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L29
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_remove_self_joins_recurse(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v139 int32
	_ = v139
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v346 int32
	_ = v346
	var v380 int32
	_ = v380
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v672 int32
	_ = v672
	var v683 int32
	_ = v683
	var v703 int32
	_ = v703
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v717 int32
	_ = v717
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v809 int32
	_ = v809
	var v822 int32
	_ = v822
	var v826 int32
	_ = v826
	var v835 int32
	_ = v835
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v883 int32
	_ = v883
	var v887 int32
	_ = v887
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v906 int32
	_ = v906
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v940 int32
	_ = v940
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v960 int32
	_ = v960
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v972 int32
	_ = v972
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v984 int32
	_ = v984
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v1000 int32
	_ = v1000
	var v1004 int32
	_ = v1004
	var v1006 int32
	_ = v1006
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1020 int32
	_ = v1020
	var v1022 int32
	_ = v1022
	var v1025 int32
	_ = v1025
	var v1033 int32
	_ = v1033
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1049 int32
	_ = v1049
	var v1053 int32
	_ = v1053
	var v1055 int32
	_ = v1055
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1069 int32
	_ = v1069
	var v1071 int32
	_ = v1071
	var v1074 int32
	_ = v1074
	var v1082 int32
	_ = v1082
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1089 int32
	_ = v1089
	var v1092 int32
	_ = v1092
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1104 int32
	_ = v1104
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1111 int32
	_ = v1111
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1178 int32
	_ = v1178
	var v1180 int32
	_ = v1180
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1189 int32
	_ = v1189
	var v1192 int32
	_ = v1192
	var v1195 int32
	_ = v1195
	var v1210 int32
	_ = v1210
	var v1225 int32
	_ = v1225
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1243 int32
	_ = v1243
	var v1245 int32
	_ = v1245
	var v1247 int32
	_ = v1247
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1262 int32
	_ = v1262
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1269 int32
	_ = v1269
	var v1276 int32
	_ = v1276
	var v1301 int32
	_ = v1301
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1316 int32
	_ = v1316
	var v1318 int32
	_ = v1318
	var v1320 int32
	_ = v1320
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1326 int32
	_ = v1326
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1390 int32
	_ = v1390
	var v1392 int32
	_ = v1392
	var v1395 int32
	_ = v1395
	var v1399 int32
	_ = v1399
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1426 int32
	_ = v1426
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1448 int32
	_ = v1448
	var v1449 int32
	_ = v1449
	var v1452 int32
	_ = v1452
	var v1456 int32
	_ = v1456
	var v1458 int32
	_ = v1458
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1472 int32
	_ = v1472
	var v1474 int32
	_ = v1474
	var v1477 int32
	_ = v1477
	var v1485 int32
	_ = v1485
	var v1488 int32
	_ = v1488
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1495 int32
	_ = v1495
	var v1496 int32
	_ = v1496
	var v1506 int32
	_ = v1506
	var v1507 int32
	_ = v1507
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1532 int32
	_ = v1532
	var v1533 int32
	_ = v1533
	var v1539 int32
	_ = v1539
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1564 int32
	_ = v1564
	var v1568 int32
	_ = v1568
	var v1569 int32
	_ = v1569
	var v1570 int32
	_ = v1570
	var v1573 int32
	_ = v1573
	var v1574 int32
	_ = v1574
	var v1575 int32
	_ = v1575
	var v1583 int32
	_ = v1583
	var v1584 int32
	_ = v1584
	var v1587 int32
	_ = v1587
	var v1591 int32
	_ = v1591
	var v1593 int32
	_ = v1593
	var v1600 int32
	_ = v1600
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1607 int32
	_ = v1607
	var v1609 int32
	_ = v1609
	var v1612 int32
	_ = v1612
	var v1620 int32
	_ = v1620
	var v1623 int32
	_ = v1623
	var v1624 int32
	_ = v1624
	var v1625 int32
	_ = v1625
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1628 int32
	_ = v1628
	var v1630 int32
	_ = v1630
	var v1631 int32
	_ = v1631
	var v1641 int32
	_ = v1641
	var v1642 int32
	_ = v1642
	var v1664 int32
	_ = v1664
	var v1668 int32
	_ = v1668
	var v1670 int32
	_ = v1670
	var v1672 int32
	_ = v1672
	var v1673 int32
	_ = v1673
	var v1683 int32
	_ = v1683
	var v1684 int32
	_ = v1684
	var v1687 int32
	_ = v1687
	var v1691 int32
	_ = v1691
	var v1694 int32
	_ = v1694
	var v1696 int32
	_ = v1696
	var v1699 int32
	_ = v1699
	var v1706 int32
	_ = v1706
	var v1708 int32
	_ = v1708
	var v1716 int32
	_ = v1716
	var v1717 int32
	_ = v1717
	var v1730 int32
	_ = v1730
	var v1742 int32
	_ = v1742
	var v1761 int32
	_ = v1761
	var v1762 int32
	_ = v1762
	var v1763 int32
	_ = v1763
	var v1766 int32
	_ = v1766
	var v1767 int32
	_ = v1767
	var v1771 int32
	_ = v1771
	var v1772 int32
	_ = v1772
	var v1775 int32
	_ = v1775
	var v1776 int32
	_ = v1776
	var v1780 int32
	_ = v1780
	var v1793 int32
	_ = v1793
	var v1807 int32
	_ = v1807
	var v1811 int32
	_ = v1811
	var v1812 int32
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1814 int32
	_ = v1814
	var v1817 int32
	_ = v1817
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1821 int32
	_ = v1821
	var v1822 int32
	_ = v1822
	var v1823 int32
	_ = v1823
	var v1824 int32
	_ = v1824
	var v1825 int32
	_ = v1825
	var v1827 int32
	_ = v1827
	var v1830 int32
	_ = v1830
	var v1833 int32
	_ = v1833
	var v1834 int32
	_ = v1834
	var v1840 int32
	_ = v1840
	var v1865 int32
	_ = v1865
	var v1866 int32
	_ = v1866
	var v1870 int32
	_ = v1870
	var v1871 int32
	_ = v1871
	var v1872 int32
	_ = v1872
	var v1873 int32
	_ = v1873
	var v1874 int32
	_ = v1874
	var v1875 int32
	_ = v1875
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1879 int32
	_ = v1879
	var v1880 int32
	_ = v1880
	var v1910 int32
	_ = v1910
	var v1911 int32
	_ = v1911
	var v1913 int32
	_ = v1913
	var v1941 int32
	_ = v1941
	var v1942 int32
	_ = v1942
	var v1944 int32
	_ = v1944
	var v1946 int32
	_ = v1946
	var v1973 int32
	_ = v1973
	var v1975 int32
	_ = v1975
	var v1978 int32
	_ = v1978
	var v1979 int32
	_ = v1979
	var v1980 int32
	_ = v1980
	var v1983 int32
	_ = v1983
	var v1984 int32
	_ = v1984
	var v1988 int32
	_ = v1988
	var v2001 int32
	_ = v2001
	var v2015 int32
	_ = v2015
	var v2019 int32
	_ = v2019
	var v2020 int32
	_ = v2020
	var v2021 int32
	_ = v2021
	var v2022 int32
	_ = v2022
	var v2027 int32
	_ = v2027
	var v2030 int32
	_ = v2030
	var v2031 int32
	_ = v2031
	var v2037 int32
	_ = v2037
	var v2062 int32
	_ = v2062
	var v2063 int32
	_ = v2063
	var v2067 int32
	_ = v2067
	var v2068 int32
	_ = v2068
	var v2069 int32
	_ = v2069
	var v2070 int32
	_ = v2070
	var v2071 int32
	_ = v2071
	var v2072 int32
	_ = v2072
	var v2073 int32
	_ = v2073
	var v2074 int32
	_ = v2074
	var v2076 int32
	_ = v2076
	var v2077 int32
	_ = v2077
	var v2107 int32
	_ = v2107
	var v2108 int32
	_ = v2108
	var v2110 int32
	_ = v2110
	var v2138 int32
	_ = v2138
	var v2139 int32
	_ = v2139
	var v2141 int32
	_ = v2141
	var v2143 int32
	_ = v2143
	var v2150 int32
	_ = v2150
	var v2171 int32
	_ = v2171
	var v2173 int32
	_ = v2173
	var v2174 int32
	_ = v2174
	var v2175 int32
	_ = v2175
	var v2177 int32
	_ = v2177
	var v2178 int32
	_ = v2178
	var v2179 int32
	_ = v2179
	var v2181 int32
	_ = v2181
	var v2188 int32
	_ = v2188
	var v2190 int32
	_ = v2190
	var v2191 int32
	_ = v2191
	var v2194 int32
	_ = v2194
	var v2198 int32
	_ = v2198
	var v2201 int32
	_ = v2201
	var v2203 int32
	_ = v2203
	var v2206 int32
	_ = v2206
	var v2213 int32
	_ = v2213
	var v2215 int32
	_ = v2215
	var v2223 int32
	_ = v2223
	var v2224 int32
	_ = v2224
	var v2237 int32
	_ = v2237
	var v2268 int32
	_ = v2268
	var v2269 int32
	_ = v2269
	var v2272 int32
	_ = v2272
	var v2273 int32
	_ = v2273
	var v2279 int32
	_ = v2279
	var v2304 int32
	_ = v2304
	var v2308 int32
	_ = v2308
	var v2309 int32
	_ = v2309
	var v2310 int32
	_ = v2310
	var v2313 int32
	_ = v2313
	var v2314 int32
	_ = v2314
	var v2315 int32
	_ = v2315
	var v2316 int32
	_ = v2316
	var v2317 int32
	_ = v2317
	var v2320 int32
	_ = v2320
	var v2321 int32
	_ = v2321
	var v2322 int32
	_ = v2322
	var v2323 int32
	_ = v2323
	var v2324 int32
	_ = v2324
	var v2328 int32
	_ = v2328
	var v2329 int32
	_ = v2329
	var v2359 int32
	_ = v2359
	var v2360 int32
	_ = v2360
	var v2363 int32
	_ = v2363
	var v2390 int32
	_ = v2390
	var v2393 int32
	_ = v2393
	var v2394 int32
	_ = v2394
	var v2396 int32
	_ = v2396
	var v2397 int32
	_ = v2397
	var v2398 int32
	_ = v2398
	var v2399 int32
	_ = v2399
	var v2400 int32
	_ = v2400
	var v2401 int32
	_ = v2401
	var v2404 int32
	_ = v2404
	var v2406 int32
	_ = v2406
	var v2407 int32
	_ = v2407
	var v2409 int32
	_ = v2409
	var v2410 int32
	_ = v2410
	var v2411 int32
	_ = v2411
	var v2412 int32
	_ = v2412
	var v2415 int32
	_ = v2415
	var v2449 int32
	_ = v2449
	var v2450 int32
	_ = v2450
	var v2451 int32
	_ = v2451
	var v2453 int32
	_ = v2453
	var v2457 int32
	_ = v2457
	var v2458 int32
	_ = v2458
	var v2459 int32
	_ = v2459
	var v2462 int32
	_ = v2462
	var v2463 int32
	_ = v2463
	var v2464 int32
	_ = v2464
	var v2467 int32
	_ = v2467
	var v2468 int32
	_ = v2468
	var v2469 int32
	_ = v2469
	var v2470 int32
	_ = v2470
	var v2473 int32
	_ = v2473
	var v2474 int32
	_ = v2474
	var v2475 int32
	_ = v2475
	var v2476 int32
	_ = v2476
	var v2479 int32
	_ = v2479
	var v2480 int32
	_ = v2480
	var v2481 int32
	_ = v2481
	var v2482 int32
	_ = v2482
	var v2483 int32
	_ = v2483
	var v2484 int32
	_ = v2484
	var v2485 int32
	_ = v2485
	var v2486 int32
	_ = v2486
	var v2487 int32
	_ = v2487
	var v2488 int32
	_ = v2488
	var v2489 int32
	_ = v2489
	var v2490 int32
	_ = v2490
	var v2491 int32
	_ = v2491
	var v2492 int32
	_ = v2492
	var v2495 int32
	_ = v2495
	var v2497 int32
	_ = v2497
	var v2498 int32
	_ = v2498
	var v2505 int32
	_ = v2505
	var v2507 int32
	_ = v2507
	var v2509 int32
	_ = v2509
	var v2511 int32
	_ = v2511
	var v2513 int32
	_ = v2513
	var v2514 int32
	_ = v2514
	var v2515 int32
	_ = v2515
	var v2550 int32
	_ = v2550
	var v2552 int32
	_ = v2552
	var v2553 int32
	_ = v2553
	var v2556 int32
	_ = v2556
	var v2560 int32
	_ = v2560
	var v2563 int32
	_ = v2563
	var v2565 int32
	_ = v2565
	var v2568 int32
	_ = v2568
	var v2575 int32
	_ = v2575
	var v2577 int32
	_ = v2577
	var v2585 int32
	_ = v2585
	var v2586 int32
	_ = v2586
	var v2599 int32
	_ = v2599
	var v2625 int32
	_ = v2625
	var v2636 int32
	_ = v2636
	var v2638 int32
	_ = v2638
	var v2639 int32
	_ = v2639
	var v2642 int32
	_ = v2642
	var v2646 int32
	_ = v2646
	var v2649 int32
	_ = v2649
	var v2651 int32
	_ = v2651
	var v2654 int32
	_ = v2654
	var v2661 int32
	_ = v2661
	var v2663 int32
	_ = v2663
	var v2671 int32
	_ = v2671
	var v2672 int32
	_ = v2672
	var v2685 int32
	_ = v2685
	var v2688 int32
	_ = v2688
	var v2690 int32
	_ = v2690
	var v2699 int32
	_ = v2699
	var v2703 int32
	_ = v2703
	var v2705 int32
	_ = v2705
	var v2707 int32
	_ = v2707
	var v2709 int32
	_ = v2709
	var v2710 int32
	_ = v2710
	var v2711 int32
	_ = v2711
	var v2712 int32
	_ = v2712
	var v2716 int32
	_ = v2716
	var v2717 int32
	_ = v2717
	var v2718 int32
	_ = v2718
	var v2719 int32
	_ = v2719
	var v2720 int32
	_ = v2720
	var v2728 int32
	_ = v2728
	var v2729 int32
	_ = v2729
	var v2732 int32
	_ = v2732
	var v2736 int32
	_ = v2736
	var v2738 int32
	_ = v2738
	var v2745 int32
	_ = v2745
	var v2746 int32
	_ = v2746
	var v2747 int32
	_ = v2747
	var v2752 int32
	_ = v2752
	var v2754 int32
	_ = v2754
	var v2757 int32
	_ = v2757
	var v2765 int32
	_ = v2765
	var v2769 int32
	_ = v2769
	var v2771 int32
	_ = v2771
	var v2775 int32
	_ = v2775
	var v2776 int32
	_ = v2776
	var v2777 int32
	_ = v2777
	var v2778 int32
	_ = v2778
	var v2780 int32
	_ = v2780
	var v2789 int32
	_ = v2789
	var v2795 int32
	_ = v2795
	var v2797 int32
	_ = v2797
	var v2799 int32
	_ = v2799
	var v2800 int32
	_ = v2800
	var v2802 int32
	_ = v2802
	var v2811 int32
	_ = v2811
	var v2830 int32
	_ = v2830
	v4 = int32(0)
	v29 = m.G0
	v31 = v29 - int32(16)
	m.G0 = v31
	if l1 == v4 {
		v122 = l2
		v139 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v148 = int32(0)
	if v139 == v148 {
		goto L27
	} else {
		goto L28
	}
L2:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v35 <= int32(0) {
		v122 = l2
		v139 = v4
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v40 = l2
	v41 = v4
	v57 = v4
	goto L4
L4:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v66+v41<<(uint(int32(2))%32))))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	if v71 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v122 = v112
	v139 = v115
	goto L1
L6:
	;
	v117 = v41 + int32(1)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v117 < v118 {
		v40 = v112
		v41 = v117
		v57 = v115
		goto L4
	} else {
		goto L24
	}
L7:
	;
	if v71 == int32(63) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v110 = F_remove_self_joins_recurse(m, l0, v70, v40)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L18
	} else {
		goto L23
	}
L10:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v76+v77<<(uint(int32(2))%32))))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
	if v82 != 0 {
		v112 = v40
		v115 = v57
		goto L6
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L18
	} else {
		goto L20
	}
L13:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+21)))
	if v83 != int32(114) {
		v112 = v40
		v115 = v57
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v81)+32))
	if v86 != 0 {
		v112 = v40
		v115 = v57
		goto L6
	} else {
		goto L15
	}
L15:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+32))
	if v77 == v88 {
		v112 = v40
		v115 = v57
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v87)+68))
	if v77 == v90 {
		v112 = v40
		v115 = v57
		goto L6
	} else {
		goto L17
	}
L17:
	;
	v92 = F_bms_add_member(m, v57, v77)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return int32(0)
L19:
	;
	v112 = v40
	v115 = v92
	goto L6
L20:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v100
	F_errmsg_internal(m, int32(_a_F_remove_self_joins_recurse_0), v31)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_remove_self_joins_recurse_1), int32(2352), int32(_a_F_remove_self_joins_recurse_2))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	v112 = v110
	v115 = v57
	goto L6
L24:
	;
	goto L5
L25:
	;
	m.G0 = v2830 + int32(16)
	return v2811
L26:
	;
	if v183 < int32(2) {
		v2811 = v122
		v2830 = v31
		goto L25
	} else {
		goto L39
	}
L27:
	;
	v183 = int32(0)
	goto L26
L28:
	;
	goto L29
L29:
	;
	v155 = int32(1)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	if v156 <= v155 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v159 = v155
	goto L32
L31:
	;
	v159 = v156
	goto L32
L32:
	;
	v163 = int32(0)
	v165 = v148
	goto L33
L33:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v139+int32(8)+v163<<(uint(int32(2))%32))))
	if v171 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v183 = v174
	goto L26
L35:
	;
	v174 = v165 + base.I32_popcnt(v171)
	goto L37
L36:
	;
	v174 = v165
	goto L37
L37:
	;
	v176 = v163 + int32(1)
	if v176 != v159 {
		v163 = v176
		v165 = v174
		goto L33
	} else {
		goto L38
	}
L38:
	;
	goto L34
L39:
	;
	v189 = F_palloc(m, v183<<(uint(int32(3))%32))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L18
	} else {
		goto L40
	}
L40:
	;
	if v139 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	if int32(0) <= v247 {
		goto L52
	} else {
		goto L53
	}
L42:
	;
	v247 = base.I32_ctz(v233) | v234<<(uint(int32(5))%32)
	goto L41
L43:
	;
	v247 = int32(-2)
	goto L41
L44:
	;
	v200 = base.I32_div_s(int32(0), int32(32))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	if v201 <= v200 {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v204 = v139 + int32(8)
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v204+v200<<(uint(int32(2))%32))))
	v211 = v208 & int32(-1)
	if v211 != 0 {
		v233 = v211
		v234 = v200
		goto L42
	} else {
		goto L46
	}
L46:
	;
	v213 = v200 + int32(1)
	if v213 == v201 {
		goto L43
	} else {
		goto L47
	}
L47:
	;
	v216 = v213
	goto L48
L48:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v204+v216<<(uint(int32(2))%32))))
	if v223 != 0 {
		v233 = v223
		v234 = v216
		goto L42
	} else {
		goto L50
	}
L49:
	;
	goto L43
L50:
	;
	v225 = v216 + int32(1)
	if v225 != v201 {
		v216 = v225
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v251 = int32(0)
	v253 = v247
	goto L55
L53:
	;
	goto L54
L54:
	;
	F_pg_qsort(m, v189, v183, int32(8), int32(829))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L18
	} else {
		goto L69
	}
L55:
	;
	v280 = v189 + v251<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v280))) = v253
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v282+v253<<(uint(int32(2))%32))))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v286)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v280)+4)) = v287
	if v139 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	goto L54
L57:
	;
	if int32(0) <= v346 {
		v251 = v251 + int32(1)
		v253 = v346
		goto L55
	} else {
		goto L68
	}
L58:
	;
	v346 = base.I32_ctz(v332) | v333<<(uint(int32(5))%32)
	goto L57
L59:
	;
	v346 = int32(-2)
	goto L57
L60:
	;
	v297 = v253 + int32(1)
	v299 = base.I32_div_s(v297, int32(32))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	if v300 <= v299 {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v303 = v139 + int32(8)
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v303+v299<<(uint(int32(2))%32))))
	v310 = v307 & (int32(-1) << (uint(v297) % 32))
	if v310 != 0 {
		v332 = v310
		v333 = v299
		goto L58
	} else {
		goto L62
	}
L62:
	;
	v312 = v299 + int32(1)
	if v312 == v300 {
		goto L59
	} else {
		goto L63
	}
L63:
	;
	v315 = v312
	goto L64
L64:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v303+v315<<(uint(int32(2))%32))))
	if v322 != 0 {
		v332 = v322
		v333 = v315
		goto L58
	} else {
		goto L66
	}
L65:
	;
	goto L59
L66:
	;
	v324 = v315 + int32(1)
	if v324 != v300 {
		v315 = v324
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	goto L56
L69:
	;
	if v183+int32(1) < int32(2) {
		v2811 = v122
		v2830 = v31
		goto L25
	} else {
		goto L70
	}
L70:
	;
	v386 = l0
	v388 = v122
	v389 = int32(1)
	v403 = v4
	v405 = v139
	v407 = v31
	v408 = v189
	v410 = v183
	goto L71
L71:
	;
	if v410 != v389 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v2811 = v2780
	v2830 = v2799
	goto L25
L73:
	;
	if v2789 != v2802 {
		v386 = v2778
		v388 = v2780
		v389 = v2789 + int32(1)
		v403 = v2795
		v405 = v2797
		v407 = v2799
		v408 = v2800
		v410 = v2802
		goto L71
	} else {
		goto L529
	}
L74:
	;
	v415 = int32(3)
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v408+v389<<(uint(v415)%32))+4))
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v408+v403<<(uint(v415)%32))+4))
	if v418 == v422 {
		v2778 = v386
		v2780 = v388
		v2789 = v389
		v2795 = v403
		v2797 = v405
		v2799 = v407
		v2800 = v408
		v2802 = v410
		goto L73
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	if int32(2) <= v389-v403 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	goto L76
L78:
	;
	v427 = int32(0)
	if v403 < v389 {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L80
L80:
	;
	v2775 = *(*int32)(unsafe.Add(mBase, uint32(v408+v403<<(uint(int32(3))%32))))
	v2776 = F_bms_del_member(m, v405, v2775)
	mBase = m.M
	v2777 = m.ExcPending
	if v2777 != 0 {
		goto L18
	} else {
		goto L528
	}
L81:
	;
	v444 = v427
	v446 = v403
	goto L84
L82:
	;
	v481 = v427
	v483 = v403
	goto L83
L83:
	;
	v494 = F_bms_del_members(m, v405, v481)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L18
	} else {
		goto L88
	}
L84:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v408+v446<<(uint(int32(3))%32))))
	v461 = F_bms_add_member(m, v444, v460)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L18
	} else {
		goto L86
	}
L85:
	;
	v481 = v461
	v483 = v389
	goto L83
L86:
	;
	v464 = v446 + int32(1)
	if v464 != v389 {
		v444 = v461
		v446 = v464
		goto L84
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	v496 = v386
	v498 = v388
	v507 = v389
	v511 = v481
	v513 = v483
	v515 = v494
	v517 = v407
	v518 = v408
	v520 = v410
	goto L89
L89:
	;
	v524 = int32(0)
	if v511 == v524 {
		goto L93
	} else {
		goto L94
	}
L90:
	;
	F_bms_free(m, v2711)
	mBase = m.M
	v2769 = m.ExcPending
	if v2769 != 0 {
		goto L18
	} else {
		goto L526
	}
L91:
	;
	if int32(0) < v581 {
		goto L102
	} else {
		goto L103
	}
L92:
	;
	v581 = base.I32_ctz(v567) | v568<<(uint(int32(5))%32)
	goto L91
L93:
	;
	v581 = int32(-2)
	goto L91
L94:
	;
	v534 = base.I32_div_s(int32(0), int32(32))
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v511)+4))
	if v535 <= v534 {
		goto L93
	} else {
		goto L95
	}
L95:
	;
	v538 = v511 + int32(8)
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v538+v534<<(uint(int32(2))%32))))
	v545 = v542 & int32(-1)
	if v545 != 0 {
		v567 = v545
		v568 = v534
		goto L92
	} else {
		goto L96
	}
L96:
	;
	v547 = v534 + int32(1)
	if v547 == v535 {
		goto L93
	} else {
		goto L97
	}
L97:
	;
	v550 = v547
	goto L98
L98:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v538+v550<<(uint(int32(2))%32))))
	if v557 != 0 {
		v567 = v557
		v568 = v550
		goto L92
	} else {
		goto L100
	}
L99:
	;
	goto L93
L100:
	;
	v559 = v550 + int32(1)
	if v559 != v535 {
		v550 = v559
		goto L98
	} else {
		goto L101
	}
L101:
	;
	goto L99
L102:
	;
	v584 = v496
	v586 = v498
	v595 = v507
	v599 = v511
	v601 = v513
	v602 = v581
	v603 = v515
	v605 = v517
	v606 = v518
	v607 = v524
	v608 = v520
	goto L105
L103:
	;
	v2688 = v496
	v2690 = v498
	v2699 = v507
	v2703 = v511
	v2705 = v513
	v2707 = v515
	v2709 = v517
	v2710 = v518
	v2711 = v524
	v2712 = v520
	goto L104
L104:
	;
	v2716 = F_bms_add_members(m, v2690, v2711)
	mBase = m.M
	v2717 = m.ExcPending
	if v2717 != 0 {
		goto L18
	} else {
		goto L504
	}
L105:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v584)+28))
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v612+v602<<(uint(int32(2))%32))))
	if v599 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L106:
	;
	v2688 = v584
	v2690 = v586
	v2699 = v595
	v2703 = v599
	v2705 = v601
	v2707 = v603
	v2709 = v605
	v2710 = v606
	v2711 = v2625
	v2712 = v608
	goto L104
L107:
	;
	if v599 == int32(0) {
		goto L494
	} else {
		goto L495
	}
L108:
	;
	if v672 <= int32(0) {
		v2625 = v607
		goto L107
	} else {
		goto L119
	}
L109:
	;
	v672 = base.I32_ctz(v658) | v659<<(uint(int32(5))%32)
	goto L108
L110:
	;
	v672 = int32(-2)
	goto L108
L111:
	;
	v623 = v602 + int32(1)
	v625 = base.I32_div_s(v623, int32(32))
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v599)+4))
	if v626 <= v625 {
		goto L110
	} else {
		goto L112
	}
L112:
	;
	v629 = v599 + int32(8)
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v629+v625<<(uint(int32(2))%32))))
	v636 = v633 & (int32(-1) << (uint(v623) % 32))
	if v636 != 0 {
		v658 = v636
		v659 = v625
		goto L109
	} else {
		goto L113
	}
L113:
	;
	v638 = v625 + int32(1)
	if v638 == v626 {
		goto L110
	} else {
		goto L114
	}
L114:
	;
	v641 = v638
	goto L115
L115:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v629+v641<<(uint(int32(2))%32))))
	if v648 != 0 {
		v658 = v648
		v659 = v641
		goto L109
	} else {
		goto L117
	}
L116:
	;
	goto L110
L117:
	;
	v650 = v641 + int32(1)
	if v650 != v626 {
		v641 = v650
		goto L115
	} else {
		goto L118
	}
L118:
	;
	goto L116
L119:
	;
	v683 = v672
	goto L120
L120:
	;
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v584)+28))
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v703+v683<<(uint(int32(2))%32))))
	v708 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v605)+12)) = v708
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v584)+112))
	if v710 == v708 {
		goto L123
	} else {
		goto L124
	}
L121:
	;
	v2625 = v607
	goto L107
L122:
	;
	if v599 == int32(0) {
		goto L482
	} else {
		goto L483
	}
L123:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v584)+136))
	if v793 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L124:
	;
	v717 = int32(0)
	goto L125
L125:
	;
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v710)+4))
	if v742 <= v717 {
		goto L123
	} else {
		goto L127
	}
L126:
	;
	goto L122
L127:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v710)+12))
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v744+v717<<(uint(int32(2))%32))))
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v748)+12))
	v750 = F_bms_is_member(m, v683, v749)
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L18
	} else {
		goto L128
	}
L128:
	;
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v748)+12))
	v753 = F_bms_is_member(m, v602, v752)
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L18
	} else {
		goto L129
	}
L129:
	;
	if v750 != v753 {
		goto L122
	} else {
		goto L130
	}
L130:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v748)+16))
	v759 = F_bms_is_member(m, v683, v758)
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L18
	} else {
		goto L131
	}
L131:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v748)+16))
	v762 = F_bms_is_member(m, v602, v761)
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L18
	} else {
		goto L132
	}
L132:
	;
	if v759 == v762 {
		v717 = v717 + int32(1)
		goto L125
	} else {
		goto L133
	}
L133:
	;
	goto L126
L134:
	;
	v896 = F_bms_add_member(m, int32(0), v602)
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L18
	} else {
		goto L158
	}
L135:
	;
	v796 = int32(0)
	v883 = v796
	v887 = v796
	goto L134
L136:
	;
	goto L137
L137:
	;
	v798 = int32(0)
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v793)+4))
	if v798 < v799 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v803 = v799
	goto L140
L139:
	;
	v803 = v798
	goto L140
L140:
	;
	v804 = int32(0)
	v809 = v804
	v822 = v798
	v826 = v804
	goto L141
L141:
	;
	if v809 != v803 {
		goto L143
	} else {
		goto L144
	}
L142:
	;
	v859 = int32(0)
	if base.B2i32(v858 == v859)|base.B2i32(v857 == v859) != 0 {
		v883 = v857
		v887 = v858
		goto L134
	} else {
		goto L156
	}
L143:
	;
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v793)+12))
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v835+v809<<(uint(int32(2))%32))))
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v839)+4))
	v841 = base.B2i32(v840 == v602)
	if v840 == v602 {
		goto L146
	} else {
		goto L147
	}
L144:
	;
	v857 = v822
	v858 = v826
	goto L145
L145:
	;
	goto L142
L146:
	;
	v842 = v839
	goto L148
L147:
	;
	v842 = v822
	goto L148
L148:
	;
	if v683 == v840 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v846 = v839
	goto L151
L150:
	;
	v846 = v826
	goto L151
L151:
	;
	if v840 == v602 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v847 = v826
	goto L154
L153:
	;
	v847 = v846
	goto L154
L154:
	;
	v848 = int32(0)
	if base.B2i32(v847 == v848)|base.B2i32(v842 == v848) != 0 {
		v809 = v809 + int32(1)
		v822 = v842
		v826 = v847
		goto L141
	} else {
		goto L155
	}
L155:
	;
	v857 = v842
	v858 = v847
	goto L145
L156:
	;
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v858)+16))
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v857)+16))
	if v864 != v865 {
		goto L122
	} else {
		goto L157
	}
L157:
	;
	v883 = v857
	v887 = v858
	goto L134
L158:
	;
	v898 = F_bms_add_member(m, v896, v683)
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L18
	} else {
		goto L159
	}
L159:
	;
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v616)+8))
	v902 = F_generate_join_implied_equalities(m, v584, v898, v900, v707, int32(0))
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L18
	} else {
		goto L160
	}
L160:
	;
	if v902 == int32(0) {
		goto L122
	} else {
		goto L161
	}
L161:
	;
	v906 = int32(0)
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v902)+4))
	if v906 < v909 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v913 = v906
	v921 = v906
	v922 = v906
	goto L165
L163:
	;
	v1154 = v906
	v1155 = v906
	goto L164
L164:
	;
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(v707)+184))
	v1174 = F_list_concat(m, v1155, v1173)
	mBase = m.M
	v1175 = m.ExcPending
	if v1175 != 0 {
		goto L18
	} else {
		goto L238
	}
L165:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v902)+12))
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v940+v913<<(uint(int32(2))%32))))
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v944)+96))
	if v945 == int32(0) {
		goto L168
	} else {
		goto L169
	}
L166:
	;
	v1154 = v1138
	v1155 = v1139
	goto L164
L167:
	;
	v1142 = v913 + int32(1)
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(v902)+4))
	if v1142 < v1143 {
		v913 = v1142
		v921 = v1138
		v922 = v1139
		goto L165
	} else {
		goto L237
	}
L168:
	;
	v1134 = F_lappend(m, v921, v944)
	mBase = m.M
	v1135 = m.ExcPending
	if v1135 != 0 {
		goto L18
	} else {
		goto L236
	}
L169:
	;
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v944)+28))
	v949 = int32(0)
	if v948 == v949 {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	if v984 != int32(2) {
		goto L168
	} else {
		goto L183
	}
L171:
	;
	v984 = int32(0)
	goto L170
L172:
	;
	goto L173
L173:
	;
	v956 = int32(1)
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v948)+4))
	if v957 <= v956 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v960 = v956
	goto L176
L175:
	;
	v960 = v957
	goto L176
L176:
	;
	v964 = int32(0)
	v966 = v949
	goto L177
L177:
	;
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v948+int32(8)+v964<<(uint(int32(2))%32))))
	if v972 != 0 {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	v984 = v975
	goto L170
L179:
	;
	v975 = v966 + base.I32_popcnt(v972)
	goto L181
L180:
	;
	v975 = v966
	goto L181
L181:
	;
	v977 = v964 + int32(1)
	if v977 != v960 {
		v964 = v977
		v966 = v975
		goto L177
	} else {
		goto L182
	}
L182:
	;
	goto L178
L183:
	;
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v944)+44))
	v988 = int32(0)
	if v987 == v988 {
		goto L185
	} else {
		goto L186
	}
L184:
	;
	if v1033 != int32(1) {
		goto L168
	} else {
		goto L200
	}
L185:
	;
	v1033 = int32(0)
	goto L184
L186:
	;
	goto L187
L187:
	;
	v996 = int32(1)
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v987)+4))
	if v997 <= v996 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v1000 = v996
	goto L190
L189:
	;
	v1000 = v997
	goto L190
L190:
	;
	v1004 = int32(0)
	v1006 = v988
	goto L191
L191:
	;
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(v987+int32(8)+v1004<<(uint(int32(2))%32))))
	if v1013 != 0 {
		goto L194
	} else {
		goto L195
	}
L192:
	;
	v1033 = v1025
	goto L184
L193:
	;
	goto L192
L194:
	;
	v1014 = int32(2)
	if v1006 != 0 {
		v1025 = v1014
		goto L193
	} else {
		goto L197
	}
L195:
	;
	v1020 = v1006
	goto L196
L196:
	;
	v1022 = v1004 + int32(1)
	if v1022 != v1000 {
		v1004 = v1022
		v1006 = v1020
		goto L191
	} else {
		goto L199
	}
L197:
	;
	v1015 = int32(1)
	if base.Ui32(v1015) < base.Ui32(base.I32_popcnt(v1013)) {
		v1025 = v1014
		goto L193
	} else {
		goto L198
	}
L198:
	;
	v1020 = v1015
	goto L196
L199:
	;
	v1025 = v1020
	goto L193
L200:
	;
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v944)+48))
	v1037 = int32(0)
	if v1036 == v1037 {
		goto L202
	} else {
		goto L203
	}
L201:
	;
	if v1082 != int32(1) {
		goto L168
	} else {
		goto L217
	}
L202:
	;
	v1082 = int32(0)
	goto L201
L203:
	;
	goto L204
L204:
	;
	v1045 = int32(1)
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v1036)+4))
	if v1046 <= v1045 {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v1049 = v1045
	goto L207
L206:
	;
	v1049 = v1046
	goto L207
L207:
	;
	v1053 = int32(0)
	v1055 = v1037
	goto L208
L208:
	;
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v1036+int32(8)+v1053<<(uint(int32(2))%32))))
	if v1062 != 0 {
		goto L211
	} else {
		goto L212
	}
L209:
	;
	v1082 = v1074
	goto L201
L210:
	;
	goto L209
L211:
	;
	v1063 = int32(2)
	if v1055 != 0 {
		v1074 = v1063
		goto L210
	} else {
		goto L214
	}
L212:
	;
	v1069 = v1055
	goto L213
L213:
	;
	v1071 = v1053 + int32(1)
	if v1071 != v1049 {
		v1053 = v1071
		v1055 = v1069
		goto L208
	} else {
		goto L216
	}
L214:
	;
	v1064 = int32(1)
	if base.Ui32(v1064) < base.Ui32(base.I32_popcnt(v1062)) {
		v1074 = v1063
		goto L210
	} else {
		goto L215
	}
L215:
	;
	v1069 = v1064
	goto L213
L216:
	;
	v1074 = v1069
	goto L210
L217:
	;
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(v944)+4))
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v1085)))
	if v1086 != int32(17) {
		goto L168
	} else {
		goto L218
	}
L218:
	;
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v1085)+28))
	if v1089 == int32(0) {
		goto L168
	} else {
		goto L219
	}
L219:
	;
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v1089)+4))
	if v1092 != int32(2) {
		goto L168
	} else {
		goto L220
	}
L220:
	;
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v1089)+12))
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v1095)))
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v1095)+4))
	v1099 = F_copyObjectImpl(m, v1098)
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		goto L18
	} else {
		goto L221
	}
L221:
	;
	v1101 = int32(0)
	if v1096 == v1101 {
		v1108 = v1101
		goto L222
	} else {
		goto L223
	}
L222:
	;
	if v1099 == int32(0) {
		v1115 = int32(0)
		goto L225
	} else {
		goto L226
	}
L223:
	;
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(v1096)))
	if v1104 != int32(27) {
		v1108 = v1096
		goto L222
	} else {
		goto L224
	}
L224:
	;
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(v1096)+4))
	v1108 = v1107
	goto L222
L225:
	;
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(v944)+48))
	v1117 = F_bms_singleton_member(m, v1116)
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L18
	} else {
		goto L230
	}
L226:
	;
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(v1099)))
	if v1111 != int32(27) {
		goto L227
	} else {
		goto L228
	}
L227:
	;
	v1115 = v1099
	goto L225
L228:
	;
	goto L229
L229:
	;
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(v1099)+4))
	v1115 = v1114
	goto L225
L230:
	;
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(v944)+44))
	v1120 = F_bms_singleton_member(m, v1119)
	mBase = m.M
	v1121 = m.ExcPending
	if v1121 != 0 {
		goto L18
	} else {
		goto L231
	}
L231:
	;
	F_ChangeVarNodesExtended(m, v1115, v1117, v1120, int32(828))
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
		goto L18
	} else {
		goto L232
	}
L232:
	;
	v1125 = F_equal(m, v1108, v1115)
	mBase = m.M
	v1126 = m.ExcPending
	if v1126 != 0 {
		goto L18
	} else {
		goto L233
	}
L233:
	;
	if v1125 == int32(0) {
		goto L168
	} else {
		goto L234
	}
L234:
	;
	v1129 = F_lappend(m, v922, v944)
	mBase = m.M
	v1130 = m.ExcPending
	if v1130 != 0 {
		goto L18
	} else {
		goto L235
	}
L235:
	;
	v1138 = v921
	v1139 = v1129
	goto L167
L236:
	;
	v1138 = v1134
	v1139 = v922
	goto L167
L237:
	;
	goto L166
L238:
	;
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(v616)+8))
	if v1154 != 0 {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v1178 = *(*int32)(unsafe.Add(mBase, uint32(v1154)+4))
	v1180 = v1178
	goto L241
L240:
	;
	v1180 = int32(0)
	goto L241
L241:
	;
	v1185 = F_innerrel_is_unique_ext(m, v584, v898, v1176, v707, int32(0), v1174, base.B2i32(v1180 == int32(0)), v605+int32(12))
	mBase = m.M
	v1186 = m.ExcPending
	if v1186 != 0 {
		goto L18
	} else {
		goto L242
	}
L242:
	;
	if v1185 == int32(0) {
		goto L122
	} else {
		goto L243
	}
L243:
	;
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v605)+12))
	if v1189 == int32(0) {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(v616)+212))
	v1386 = F_list_copy(m, v1385)
	mBase = m.M
	v1387 = m.ExcPending
	if v1387 != 0 {
		goto L18
	} else {
		goto L293
	}
L245:
	;
	v1192 = *(*int32)(unsafe.Add(mBase, uint32(v1189)+4))
	if v1192 <= int32(0) {
		goto L244
	} else {
		goto L246
	}
L246:
	;
	v1195 = *(*int32)(unsafe.Add(mBase, uint32(v707)+68))
	v1210 = int32(0)
	goto L247
L247:
	;
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(v1189)+12))
	v1229 = *(*int32)(unsafe.Add(mBase, uint32(v1225+v1210<<(uint(int32(2))%32))))
	v1230 = *(*int32)(unsafe.Add(mBase, uint32(v1229)+4))
	v1231 = F_copyObjectImpl(m, v1230)
	mBase = m.M
	v1232 = m.ExcPending
	if v1232 != 0 {
		goto L18
	} else {
		goto L249
	}
L248:
	;
	goto L244
L249:
	;
	v1233 = *(*int32)(unsafe.Add(mBase, uint32(v616)+68))
	F_ChangeVarNodesExtended(m, v1231, v1195, v1233, int32(828))
	mBase = m.M
	v1236 = m.ExcPending
	if v1236 != 0 {
		goto L18
	} else {
		goto L250
	}
L250:
	;
	v1237 = *(*int32)(unsafe.Add(mBase, uint32(v1231)+28))
	v1238 = *(*int32)(unsafe.Add(mBase, uint32(v1229)+44))
	if v1238 == int32(0) {
		goto L252
	} else {
		goto L253
	}
L251:
	;
	v1266 = *(*int32)(unsafe.Add(mBase, uint32(v616)+184))
	if v1266 == int32(0) {
		goto L122
	} else {
		goto L265
	}
L252:
	;
	if v1237 == int32(0) {
		goto L255
	} else {
		goto L256
	}
L253:
	;
	goto L254
L254:
	;
	v1253 = int32(0)
	if v1237 == v1253 {
		goto L261
	} else {
		goto L262
	}
L255:
	;
	v1243 = int32(0)
	v1264 = v1243
	v1265 = v1243
	goto L251
L256:
	;
	goto L257
L257:
	;
	v1245 = *(*int32)(unsafe.Add(mBase, uint32(v1237)+12))
	v1247 = *(*int32)(unsafe.Add(mBase, uint32(v1237)+4))
	if int32(2) <= v1247 {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	v1250 = *(*int32)(unsafe.Add(mBase, uint32(v1245)+4))
	v1251 = v1250
	goto L260
L259:
	;
	v1251 = int32(0)
	goto L260
L260:
	;
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(v1245)))
	v1264 = v1252
	v1265 = v1251
	goto L251
L261:
	;
	v1264 = v1253
	v1265 = int32(0)
	goto L251
L262:
	;
	goto L263
L263:
	;
	v1257 = *(*int32)(unsafe.Add(mBase, uint32(v1237)+12))
	v1258 = *(*int32)(unsafe.Add(mBase, uint32(v1257)))
	v1259 = *(*int32)(unsafe.Add(mBase, uint32(v1237)+4))
	if v1259 < int32(2) {
		v1264 = v1253
		v1265 = v1258
		goto L251
	} else {
		goto L264
	}
L264:
	;
	v1262 = *(*int32)(unsafe.Add(mBase, uint32(v1257)+4))
	v1264 = v1262
	v1265 = v1258
	goto L251
L265:
	;
	v1269 = *(*int32)(unsafe.Add(mBase, uint32(v1266)+4))
	if v1269 <= int32(0) {
		goto L122
	} else {
		goto L266
	}
L266:
	;
	v1276 = int32(0)
	goto L267
L267:
	;
	v1301 = *(*int32)(unsafe.Add(mBase, uint32(v1266)+12))
	v1305 = *(*int32)(unsafe.Add(mBase, uint32(v1301+v1276<<(uint(int32(2))%32))))
	v1306 = *(*int32)(unsafe.Add(mBase, uint32(v1305)+96))
	if v1306 == int32(0) {
		goto L270
	} else {
		goto L271
	}
L268:
	;
	v1354 = v1210 + int32(1)
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(v1189)+4))
	if v1354 < v1355 {
		v1210 = v1354
		goto L247
	} else {
		goto L291
	}
L269:
	;
	goto L268
L270:
	;
	v1350 = v1276 + int32(1)
	v1351 = *(*int32)(unsafe.Add(mBase, uint32(v1266)+4))
	if v1350 < v1351 {
		v1276 = v1350
		goto L267
	} else {
		goto L290
	}
L271:
	;
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(v1305)+4))
	v1310 = *(*int32)(unsafe.Add(mBase, uint32(v1309)+28))
	v1311 = *(*int32)(unsafe.Add(mBase, uint32(v1305)+44))
	if v1311 == int32(0) {
		goto L273
	} else {
		goto L274
	}
L272:
	;
	v1339 = F_equal(m, v1265, v1336)
	mBase = m.M
	v1340 = m.ExcPending
	if v1340 != 0 {
		goto L18
	} else {
		goto L286
	}
L273:
	;
	if v1310 == int32(0) {
		goto L276
	} else {
		goto L277
	}
L274:
	;
	goto L275
L275:
	;
	v1326 = int32(0)
	if v1310 == v1326 {
		goto L282
	} else {
		goto L283
	}
L276:
	;
	v1316 = int32(0)
	v1336 = v1316
	v1337 = v1316
	goto L272
L277:
	;
	goto L278
L278:
	;
	v1318 = *(*int32)(unsafe.Add(mBase, uint32(v1310)+12))
	v1320 = *(*int32)(unsafe.Add(mBase, uint32(v1310)+4))
	if int32(2) <= v1320 {
		goto L279
	} else {
		goto L280
	}
L279:
	;
	v1323 = *(*int32)(unsafe.Add(mBase, uint32(v1318)+4))
	v1324 = v1323
	goto L281
L280:
	;
	v1324 = int32(0)
	goto L281
L281:
	;
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(v1318)))
	v1336 = v1324
	v1337 = v1325
	goto L272
L282:
	;
	v1336 = int32(0)
	v1337 = v1326
	goto L272
L283:
	;
	goto L284
L284:
	;
	v1330 = *(*int32)(unsafe.Add(mBase, uint32(v1310)+12))
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(v1330)))
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(v1310)+4))
	if v1332 < int32(2) {
		v1336 = v1331
		v1337 = v1326
		goto L272
	} else {
		goto L285
	}
L285:
	;
	v1335 = *(*int32)(unsafe.Add(mBase, uint32(v1330)+4))
	v1336 = v1331
	v1337 = v1335
	goto L272
L286:
	;
	if v1339 == int32(0) {
		goto L270
	} else {
		goto L287
	}
L287:
	;
	v1343 = F_equal(m, v1264, v1337)
	mBase = m.M
	v1344 = m.ExcPending
	if v1344 != 0 {
		goto L18
	} else {
		goto L288
	}
L288:
	;
	if v1343 != 0 {
		goto L269
	} else {
		goto L289
	}
L289:
	;
	goto L270
L290:
	;
	goto L122
L291:
	;
	goto L248
L292:
	;
	v1526 = *(*int32)(unsafe.Add(mBase, uint32(v616)+184))
	v1527 = F_list_concat(m, v1526, v902)
	mBase = m.M
	v1528 = m.ExcPending
	if v1528 != 0 {
		goto L18
	} else {
		goto L325
	}
L293:
	;
	if v1386 == int32(0) {
		goto L294
	} else {
		goto L295
	}
L294:
	;
	v1390 = int32(0)
	v1506 = v1390
	v1507 = v1390
	goto L292
L295:
	;
	goto L296
L296:
	;
	v1392 = int32(0)
	v1395 = *(*int32)(unsafe.Add(mBase, uint32(v1386)+4))
	if v1395 <= v1392 {
		v1506 = v1392
		v1507 = v1392
		goto L292
	} else {
		goto L297
	}
L297:
	;
	v1399 = v1392
	v1406 = v1392
	v1407 = v1392
	goto L298
L298:
	;
	v1426 = *(*int32)(unsafe.Add(mBase, uint32(v1386)+12))
	v1430 = *(*int32)(unsafe.Add(mBase, uint32(v1426+v1399<<(uint(int32(2))%32))))
	v1431 = *(*int32)(unsafe.Add(mBase, uint32(v1430)+32))
	F_remove_join_clause_from_rels(m, v584, v1430, v1431)
	mBase = m.M
	v1433 = m.ExcPending
	if v1433 != 0 {
		goto L18
	} else {
		goto L300
	}
L299:
	;
	v1506 = v1492
	v1507 = v1493
	goto L292
L300:
	;
	v1434 = *(*int32)(unsafe.Add(mBase, uint32(v616)+68))
	v1435 = *(*int32)(unsafe.Add(mBase, uint32(v707)+68))
	F_ChangeVarNodesExtended(m, v1430, v1434, v1435, int32(828))
	mBase = m.M
	v1438 = m.ExcPending
	if v1438 != 0 {
		goto L18
	} else {
		goto L301
	}
L301:
	;
	v1439 = *(*int32)(unsafe.Add(mBase, uint32(v1430)+32))
	v1440 = int32(0)
	if v1439 == v1440 {
		goto L304
	} else {
		goto L305
	}
L302:
	;
	v1495 = v1399 + int32(1)
	v1496 = *(*int32)(unsafe.Add(mBase, uint32(v1386)+4))
	if v1495 < v1496 {
		v1399 = v1495
		v1406 = v1492
		v1407 = v1493
		goto L298
	} else {
		goto L324
	}
L303:
	;
	if v1485 == int32(2) {
		goto L319
	} else {
		goto L320
	}
L304:
	;
	v1485 = int32(0)
	goto L303
L305:
	;
	goto L306
L306:
	;
	v1448 = int32(1)
	v1449 = *(*int32)(unsafe.Add(mBase, uint32(v1439)+4))
	if v1449 <= v1448 {
		goto L307
	} else {
		goto L308
	}
L307:
	;
	v1452 = v1448
	goto L309
L308:
	;
	v1452 = v1449
	goto L309
L309:
	;
	v1456 = int32(0)
	v1458 = v1440
	goto L310
L310:
	;
	v1465 = *(*int32)(unsafe.Add(mBase, uint32(v1439+int32(8)+v1456<<(uint(int32(2))%32))))
	if v1465 != 0 {
		goto L313
	} else {
		goto L314
	}
L311:
	;
	v1485 = v1477
	goto L303
L312:
	;
	goto L311
L313:
	;
	v1466 = int32(2)
	if v1458 != 0 {
		v1477 = v1466
		goto L312
	} else {
		goto L316
	}
L314:
	;
	v1472 = v1458
	goto L315
L315:
	;
	v1474 = v1456 + int32(1)
	if v1474 != v1452 {
		v1456 = v1474
		v1458 = v1472
		goto L310
	} else {
		goto L318
	}
L316:
	;
	v1467 = int32(1)
	if base.Ui32(v1467) < base.Ui32(base.I32_popcnt(v1465)) {
		v1477 = v1466
		goto L312
	} else {
		goto L317
	}
L317:
	;
	v1472 = v1467
	goto L315
L318:
	;
	v1477 = v1472
	goto L312
L319:
	;
	v1488 = F_lappend(m, v1407, v1430)
	mBase = m.M
	v1489 = m.ExcPending
	if v1489 != 0 {
		goto L18
	} else {
		goto L322
	}
L320:
	;
	goto L321
L321:
	;
	v1490 = F_lappend(m, v1406, v1430)
	mBase = m.M
	v1491 = m.ExcPending
	if v1491 != 0 {
		goto L18
	} else {
		goto L323
	}
L322:
	;
	v1492 = v1406
	v1493 = v1488
	goto L302
L323:
	;
	v1492 = v1490
	v1493 = v1407
	goto L302
L324:
	;
	goto L299
L325:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v616)+184)) = v1527
	if v1527 == int32(0) {
		v1641 = v1506
		v1642 = v1507
		goto L326
	} else {
		goto L327
	}
L326:
	;
	F_add_non_redundant_clauses(m, v584, v1641, v707+int32(184))
	mBase = m.M
	v1664 = m.ExcPending
	if v1664 != 0 {
		goto L18
	} else {
		goto L355
	}
L327:
	;
	v1532 = int32(0)
	v1533 = *(*int32)(unsafe.Add(mBase, uint32(v1527)+4))
	if v1533 <= v1532 {
		v1641 = v1506
		v1642 = v1507
		goto L326
	} else {
		goto L328
	}
L328:
	;
	v1539 = v1532
	v1544 = v1506
	v1545 = v1507
	goto L329
L329:
	;
	v1564 = *(*int32)(unsafe.Add(mBase, uint32(v1527)+12))
	v1568 = *(*int32)(unsafe.Add(mBase, uint32(v1564+v1539<<(uint(int32(2))%32))))
	v1569 = *(*int32)(unsafe.Add(mBase, uint32(v616)+68))
	v1570 = *(*int32)(unsafe.Add(mBase, uint32(v707)+68))
	F_ChangeVarNodesExtended(m, v1568, v1569, v1570, int32(828))
	mBase = m.M
	v1573 = m.ExcPending
	if v1573 != 0 {
		goto L18
	} else {
		goto L331
	}
L330:
	;
	v1641 = v1627
	v1642 = v1628
	goto L326
L331:
	;
	v1574 = *(*int32)(unsafe.Add(mBase, uint32(v1568)+32))
	v1575 = int32(0)
	if v1574 == v1575 {
		goto L334
	} else {
		goto L335
	}
L332:
	;
	v1630 = v1539 + int32(1)
	v1631 = *(*int32)(unsafe.Add(mBase, uint32(v1527)+4))
	if v1630 < v1631 {
		v1539 = v1630
		v1544 = v1627
		v1545 = v1628
		goto L329
	} else {
		goto L354
	}
L333:
	;
	if v1620 == int32(2) {
		goto L349
	} else {
		goto L350
	}
L334:
	;
	v1620 = int32(0)
	goto L333
L335:
	;
	goto L336
L336:
	;
	v1583 = int32(1)
	v1584 = *(*int32)(unsafe.Add(mBase, uint32(v1574)+4))
	if v1584 <= v1583 {
		goto L337
	} else {
		goto L338
	}
L337:
	;
	v1587 = v1583
	goto L339
L338:
	;
	v1587 = v1584
	goto L339
L339:
	;
	v1591 = int32(0)
	v1593 = v1575
	goto L340
L340:
	;
	v1600 = *(*int32)(unsafe.Add(mBase, uint32(v1574+int32(8)+v1591<<(uint(int32(2))%32))))
	if v1600 != 0 {
		goto L343
	} else {
		goto L344
	}
L341:
	;
	v1620 = v1612
	goto L333
L342:
	;
	goto L341
L343:
	;
	v1601 = int32(2)
	if v1593 != 0 {
		v1612 = v1601
		goto L342
	} else {
		goto L346
	}
L344:
	;
	v1607 = v1593
	goto L345
L345:
	;
	v1609 = v1591 + int32(1)
	if v1609 != v1587 {
		v1591 = v1609
		v1593 = v1607
		goto L340
	} else {
		goto L348
	}
L346:
	;
	v1602 = int32(1)
	if base.Ui32(v1602) < base.Ui32(base.I32_popcnt(v1600)) {
		v1612 = v1601
		goto L342
	} else {
		goto L347
	}
L347:
	;
	v1607 = v1602
	goto L345
L348:
	;
	v1612 = v1607
	goto L342
L349:
	;
	v1623 = F_lappend(m, v1545, v1568)
	mBase = m.M
	v1624 = m.ExcPending
	if v1624 != 0 {
		goto L18
	} else {
		goto L352
	}
L350:
	;
	goto L351
L351:
	;
	v1625 = F_lappend(m, v1544, v1568)
	mBase = m.M
	v1626 = m.ExcPending
	if v1626 != 0 {
		goto L18
	} else {
		goto L353
	}
L352:
	;
	v1627 = v1544
	v1628 = v1623
	goto L332
L353:
	;
	v1627 = v1625
	v1628 = v1545
	goto L332
L354:
	;
	goto L330
L355:
	;
	F_add_non_redundant_clauses(m, v584, v1642, v707+int32(212))
	mBase = m.M
	v1668 = m.ExcPending
	if v1668 != 0 {
		goto L18
	} else {
		goto L356
	}
L356:
	;
	F_list_free(m, v1641)
	mBase = m.M
	v1670 = m.ExcPending
	if v1670 != 0 {
		goto L18
	} else {
		goto L357
	}
L357:
	;
	F_list_free(m, v1642)
	mBase = m.M
	v1672 = m.ExcPending
	if v1672 != 0 {
		goto L18
	} else {
		goto L358
	}
L358:
	;
	v1673 = *(*int32)(unsafe.Add(mBase, uint32(v616)+136))
	if v1673 == int32(0) {
		goto L361
	} else {
		goto L362
	}
L359:
	;
	if int32(0) <= v1730 {
		goto L370
	} else {
		goto L371
	}
L360:
	;
	v1730 = base.I32_ctz(v1716) | v1717<<(uint(int32(5))%32)
	goto L359
L361:
	;
	v1730 = int32(-2)
	goto L359
L362:
	;
	v1683 = base.I32_div_s(int32(0), int32(32))
	v1684 = *(*int32)(unsafe.Add(mBase, uint32(v1673)+4))
	if v1684 <= v1683 {
		goto L361
	} else {
		goto L363
	}
L363:
	;
	v1687 = v1673 + int32(8)
	v1691 = *(*int32)(unsafe.Add(mBase, uint32(v1687+v1683<<(uint(int32(2))%32))))
	v1694 = v1691 & int32(-1)
	if v1694 != 0 {
		v1716 = v1694
		v1717 = v1683
		goto L360
	} else {
		goto L364
	}
L364:
	;
	v1696 = v1683 + int32(1)
	if v1696 == v1684 {
		goto L361
	} else {
		goto L365
	}
L365:
	;
	v1699 = v1696
	goto L366
L366:
	;
	v1706 = *(*int32)(unsafe.Add(mBase, uint32(v1687+v1699<<(uint(int32(2))%32))))
	if v1706 != 0 {
		v1716 = v1706
		v1717 = v1699
		goto L360
	} else {
		goto L368
	}
L367:
	;
	goto L361
L368:
	;
	v1708 = v1699 + int32(1)
	if v1708 != v1684 {
		v1699 = v1708
		goto L366
	} else {
		goto L369
	}
L369:
	;
	goto L367
L370:
	;
	v1742 = v1730
	goto L373
L371:
	;
	goto L372
L372:
	;
	v2268 = *(*int32)(unsafe.Add(mBase, uint32(v616)+28))
	v2269 = *(*int32)(unsafe.Add(mBase, uint32(v2268)+4))
	if v2269 == int32(0) {
		goto L442
	} else {
		goto L443
	}
L373:
	;
	v1761 = *(*int32)(unsafe.Add(mBase, uint32(v707)+68))
	v1762 = *(*int32)(unsafe.Add(mBase, uint32(v616)+68))
	v1763 = int32(0)
	v1766 = *(*int32)(unsafe.Add(mBase, uint32(v584)+88))
	v1767 = *(*int32)(unsafe.Add(mBase, uint32(v1766)+12))
	v1771 = *(*int32)(unsafe.Add(mBase, uint32(v1767+v1742<<(uint(int32(2))%32))))
	v1772 = *(*int32)(unsafe.Add(mBase, uint32(v1771)+16))
	if v1772 == v1763 {
		v1946 = v1763
		v1973 = v1763
		goto L375
	} else {
		goto L376
	}
L374:
	;
	goto L372
L375:
	;
	F_list_free(m, v1973)
	mBase = m.M
	v1975 = m.ExcPending
	if v1975 != 0 {
		goto L18
	} else {
		goto L400
	}
L376:
	;
	v1775 = int32(0)
	v1776 = *(*int32)(unsafe.Add(mBase, uint32(v1772)+4))
	if v1776 <= v1775 {
		v1946 = v1763
		v1973 = v1772
		goto L375
	} else {
		goto L377
	}
L377:
	;
	v1780 = v1763
	v1793 = v1775
	goto L378
L378:
	;
	v1807 = *(*int32)(unsafe.Add(mBase, uint32(v1772)+12))
	v1811 = *(*int32)(unsafe.Add(mBase, uint32(v1807+v1793<<(uint(int32(2))%32))))
	v1812 = *(*int32)(unsafe.Add(mBase, uint32(v1811)+8))
	v1813 = F_bms_is_member(m, v1762, v1812)
	mBase = m.M
	v1814 = m.ExcPending
	if v1814 != 0 {
		goto L18
	} else {
		goto L382
	}
L379:
	;
	v1944 = *(*int32)(unsafe.Add(mBase, uint32(v1771)+16))
	v1946 = v1913
	v1973 = v1944
	goto L375
L380:
	;
	v1941 = v1793 + int32(1)
	v1942 = *(*int32)(unsafe.Add(mBase, uint32(v1772)+4))
	if v1941 < v1942 {
		v1780 = v1913
		v1793 = v1941
		goto L378
	} else {
		goto L399
	}
L381:
	;
	v1910 = F_lappend(m, v1780, v1811)
	mBase = m.M
	v1911 = m.ExcPending
	if v1911 != 0 {
		goto L18
	} else {
		goto L398
	}
L382:
	;
	if v1813 == int32(0) {
		goto L381
	} else {
		goto L383
	}
L383:
	;
	v1817 = *(*int32)(unsafe.Add(mBase, uint32(v1811)+8))
	v1818 = F_adjust_relid_set(m, v1817, v1762, v1761)
	mBase = m.M
	v1819 = m.ExcPending
	if v1819 != 0 {
		goto L18
	} else {
		goto L384
	}
L384:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1811)+8)) = v1818
	v1821 = *(*int32)(unsafe.Add(mBase, uint32(v1811)+20))
	v1822 = *(*int32)(unsafe.Add(mBase, uint32(v1821)+4))
	v1823 = F_adjust_relid_set(m, v1822, v1762, v1761)
	mBase = m.M
	v1824 = m.ExcPending
	if v1824 != 0 {
		goto L18
	} else {
		goto L385
	}
L385:
	;
	v1825 = *(*int32)(unsafe.Add(mBase, uint32(v1811)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+4)) = v1823
	v1827 = *(*int32)(unsafe.Add(mBase, uint32(v1811)+4))
	F_ChangeVarNodesExtended(m, v1827, v1762, v1761, int32(828))
	mBase = m.M
	v1830 = m.ExcPending
	if v1830 != 0 {
		goto L18
	} else {
		goto L386
	}
L386:
	;
	if v1780 == int32(0) {
		goto L381
	} else {
		goto L387
	}
L387:
	;
	v1833 = int32(0)
	v1834 = *(*int32)(unsafe.Add(mBase, uint32(v1780)+4))
	if v1834 <= v1833 {
		goto L381
	} else {
		goto L388
	}
L388:
	;
	v1840 = v1833
	goto L389
L389:
	;
	v1865 = *(*int32)(unsafe.Add(mBase, uint32(v1811)+8))
	v1866 = *(*int32)(unsafe.Add(mBase, uint32(v1780)+12))
	v1870 = *(*int32)(unsafe.Add(mBase, uint32(v1866+v1840<<(uint(int32(2))%32))))
	v1871 = *(*int32)(unsafe.Add(mBase, uint32(v1870)+8))
	v1872 = F_equal(m, v1865, v1871)
	mBase = m.M
	v1873 = m.ExcPending
	if v1873 != 0 {
		goto L18
	} else {
		goto L391
	}
L390:
	;
	goto L381
L391:
	;
	if v1872 != 0 {
		goto L392
	} else {
		goto L393
	}
L392:
	;
	v1874 = *(*int32)(unsafe.Add(mBase, uint32(v1811)+4))
	v1875 = *(*int32)(unsafe.Add(mBase, uint32(v1870)+4))
	v1876 = F_equal(m, v1874, v1875)
	mBase = m.M
	v1877 = m.ExcPending
	if v1877 != 0 {
		goto L18
	} else {
		goto L395
	}
L393:
	;
	goto L394
L394:
	;
	v1879 = v1840 + int32(1)
	v1880 = *(*int32)(unsafe.Add(mBase, uint32(v1780)+4))
	if v1879 < v1880 {
		v1840 = v1879
		goto L389
	} else {
		goto L397
	}
L395:
	;
	if v1876 != 0 {
		v1913 = v1780
		goto L380
	} else {
		goto L396
	}
L396:
	;
	goto L394
L397:
	;
	goto L390
L398:
	;
	v1913 = v1910
	goto L380
L399:
	;
	goto L379
L400:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1771)+16)) = v1946
	F_ec_clear_derived_clauses(m, v1771)
	mBase = m.M
	v1978 = m.ExcPending
	if v1978 != 0 {
		goto L18
	} else {
		goto L401
	}
L401:
	;
	v1979 = int32(0)
	v1980 = *(*int32)(unsafe.Add(mBase, uint32(v1771)+24))
	if v1980 == v1979 {
		v2143 = v1979
		v2150 = v1763
		goto L402
	} else {
		goto L403
	}
L402:
	;
	F_list_free(m, v2150)
	mBase = m.M
	v2171 = m.ExcPending
	if v2171 != 0 {
		goto L18
	} else {
		goto L427
	}
L403:
	;
	v1983 = int32(0)
	v1984 = *(*int32)(unsafe.Add(mBase, uint32(v1980)+4))
	if v1984 <= v1983 {
		goto L404
	} else {
		goto L405
	}
L404:
	;
	v2143 = v1979
	v2150 = v1980
	goto L402
L405:
	;
	goto L406
L406:
	;
	v1988 = v1979
	v2001 = v1983
	goto L407
L407:
	;
	v2015 = *(*int32)(unsafe.Add(mBase, uint32(v1980)+12))
	v2019 = *(*int32)(unsafe.Add(mBase, uint32(v2015+v2001<<(uint(int32(2))%32))))
	v2020 = *(*int32)(unsafe.Add(mBase, uint32(v2019)+32))
	v2021 = F_bms_is_member(m, v1762, v2020)
	mBase = m.M
	v2022 = m.ExcPending
	if v2022 != 0 {
		goto L18
	} else {
		goto L411
	}
L408:
	;
	v2141 = *(*int32)(unsafe.Add(mBase, uint32(v1771)+24))
	v2143 = v2110
	v2150 = v2141
	goto L402
L409:
	;
	v2138 = v2001 + int32(1)
	v2139 = *(*int32)(unsafe.Add(mBase, uint32(v1980)+4))
	if v2138 < v2139 {
		v1988 = v2110
		v2001 = v2138
		goto L407
	} else {
		goto L426
	}
L410:
	;
	v2107 = F_lappend(m, v1988, v2019)
	mBase = m.M
	v2108 = m.ExcPending
	if v2108 != 0 {
		goto L18
	} else {
		goto L425
	}
L411:
	;
	if v2021 == int32(0) {
		goto L410
	} else {
		goto L412
	}
L412:
	;
	F_ChangeVarNodesExtended(m, v2019, v1762, v1761, int32(828))
	mBase = m.M
	v2027 = m.ExcPending
	if v2027 != 0 {
		goto L18
	} else {
		goto L413
	}
L413:
	;
	if v1988 == int32(0) {
		goto L410
	} else {
		goto L414
	}
L414:
	;
	v2030 = int32(0)
	v2031 = *(*int32)(unsafe.Add(mBase, uint32(v1988)+4))
	if v2031 <= v2030 {
		goto L410
	} else {
		goto L415
	}
L415:
	;
	v2037 = v2030
	goto L416
L416:
	;
	v2062 = *(*int32)(unsafe.Add(mBase, uint32(v2019)+28))
	v2063 = *(*int32)(unsafe.Add(mBase, uint32(v1988)+12))
	v2067 = *(*int32)(unsafe.Add(mBase, uint32(v2063+v2037<<(uint(int32(2))%32))))
	v2068 = *(*int32)(unsafe.Add(mBase, uint32(v2067)+28))
	v2069 = F_equal(m, v2062, v2068)
	mBase = m.M
	v2070 = m.ExcPending
	if v2070 != 0 {
		goto L18
	} else {
		goto L418
	}
L417:
	;
	goto L410
L418:
	;
	if v2069 != 0 {
		goto L419
	} else {
		goto L420
	}
L419:
	;
	v2071 = *(*int32)(unsafe.Add(mBase, uint32(v2019)+4))
	v2072 = *(*int32)(unsafe.Add(mBase, uint32(v2067)+4))
	v2073 = F_equal(m, v2071, v2072)
	mBase = m.M
	v2074 = m.ExcPending
	if v2074 != 0 {
		goto L18
	} else {
		goto L422
	}
L420:
	;
	goto L421
L421:
	;
	v2076 = v2037 + int32(1)
	v2077 = *(*int32)(unsafe.Add(mBase, uint32(v1988)+4))
	if v2076 < v2077 {
		v2037 = v2076
		goto L416
	} else {
		goto L424
	}
L422:
	;
	if v2073 != 0 {
		v2110 = v1988
		goto L409
	} else {
		goto L423
	}
L423:
	;
	goto L421
L424:
	;
	goto L417
L425:
	;
	v2110 = v2107
	goto L409
L426:
	;
	goto L408
L427:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1771)+24)) = v2143
	v2173 = *(*int32)(unsafe.Add(mBase, uint32(v1771)+36))
	v2174 = F_adjust_relid_set(m, v2173, v1762, v1761)
	mBase = m.M
	v2175 = m.ExcPending
	if v2175 != 0 {
		goto L18
	} else {
		goto L428
	}
L428:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1771)+36)) = v2174
	v2177 = *(*int32)(unsafe.Add(mBase, uint32(v707)+136))
	v2178 = F_bms_add_member(m, v2177, v1742)
	mBase = m.M
	v2179 = m.ExcPending
	if v2179 != 0 {
		goto L18
	} else {
		goto L429
	}
L429:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v707)+136)) = v2178
	v2181 = *(*int32)(unsafe.Add(mBase, uint32(v616)+136))
	if v2181 == int32(0) {
		goto L432
	} else {
		goto L433
	}
L430:
	;
	if int32(0) <= v2237 {
		v1742 = v2237
		goto L373
	} else {
		goto L441
	}
L431:
	;
	v2237 = base.I32_ctz(v2223) | v2224<<(uint(int32(5))%32)
	goto L430
L432:
	;
	v2237 = int32(-2)
	goto L430
L433:
	;
	v2188 = v1742 + int32(1)
	v2190 = base.I32_div_s(v2188, int32(32))
	v2191 = *(*int32)(unsafe.Add(mBase, uint32(v2181)+4))
	if v2191 <= v2190 {
		goto L432
	} else {
		goto L434
	}
L434:
	;
	v2194 = v2181 + int32(8)
	v2198 = *(*int32)(unsafe.Add(mBase, uint32(v2194+v2190<<(uint(int32(2))%32))))
	v2201 = v2198 & (int32(-1) << (uint(v2188) % 32))
	if v2201 != 0 {
		v2223 = v2201
		v2224 = v2190
		goto L431
	} else {
		goto L435
	}
L435:
	;
	v2203 = v2190 + int32(1)
	if v2203 == v2191 {
		goto L432
	} else {
		goto L436
	}
L436:
	;
	v2206 = v2203
	goto L437
L437:
	;
	v2213 = *(*int32)(unsafe.Add(mBase, uint32(v2194+v2206<<(uint(int32(2))%32))))
	if v2213 != 0 {
		v2223 = v2213
		v2224 = v2206
		goto L431
	} else {
		goto L439
	}
L438:
	;
	goto L432
L439:
	;
	v2215 = v2206 + int32(1)
	if v2215 != v2191 {
		v2206 = v2215
		goto L437
	} else {
		goto L440
	}
L440:
	;
	goto L438
L441:
	;
	goto L374
L442:
	;
	v2359 = int32(*(*int16)(unsafe.Add(mBase, uint32(v707)+80)))
	v2360 = int32(*(*int16)(unsafe.Add(mBase, uint32(v707)+82)))
	if v2359 <= v2360 {
		goto L454
	} else {
		goto L455
	}
L443:
	;
	v2272 = int32(0)
	v2273 = *(*int32)(unsafe.Add(mBase, uint32(v2269)+4))
	if v2273 <= v2272 {
		goto L442
	} else {
		goto L444
	}
L444:
	;
	v2279 = v2272
	goto L445
L445:
	;
	v2304 = *(*int32)(unsafe.Add(mBase, uint32(v2269)+12))
	v2308 = *(*int32)(unsafe.Add(mBase, uint32(v2304+v2279<<(uint(int32(2))%32))))
	v2309 = *(*int32)(unsafe.Add(mBase, uint32(v616)+68))
	v2310 = *(*int32)(unsafe.Add(mBase, uint32(v707)+68))
	F_ChangeVarNodesExtended(m, v2308, v2309, v2310, int32(828))
	mBase = m.M
	v2313 = m.ExcPending
	if v2313 != 0 {
		goto L18
	} else {
		goto L447
	}
L446:
	;
	goto L442
L447:
	;
	v2314 = *(*int32)(unsafe.Add(mBase, uint32(v707)+28))
	v2315 = *(*int32)(unsafe.Add(mBase, uint32(v2314)+4))
	v2316 = F_list_member(m, v2315, v2308)
	mBase = m.M
	v2317 = m.ExcPending
	if v2317 != 0 {
		goto L18
	} else {
		goto L448
	}
L448:
	;
	if v2316 == int32(0) {
		goto L449
	} else {
		goto L450
	}
L449:
	;
	v2320 = *(*int32)(unsafe.Add(mBase, uint32(v707)+28))
	v2321 = *(*int32)(unsafe.Add(mBase, uint32(v2320)+4))
	v2322 = F_lappend(m, v2321, v2308)
	mBase = m.M
	v2323 = m.ExcPending
	if v2323 != 0 {
		goto L18
	} else {
		goto L452
	}
L450:
	;
	goto L451
L451:
	;
	v2328 = v2279 + int32(1)
	v2329 = *(*int32)(unsafe.Add(mBase, uint32(v2269)+4))
	if v2328 < v2329 {
		v2279 = v2328
		goto L445
	} else {
		goto L453
	}
L452:
	;
	v2324 = *(*int32)(unsafe.Add(mBase, uint32(v707)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v2324)+4)) = v2322
	goto L451
L453:
	;
	goto L446
L454:
	;
	v2363 = v2359
	goto L457
L455:
	;
	goto L456
L456:
	;
	if v883 == int32(0) {
		goto L462
	} else {
		goto L463
	}
L457:
	;
	v2390 = int32(*(*int16)(unsafe.Add(mBase, uint32(v707)+80)))
	v2393 = (v2363 - v2390) << (uint(int32(2)) % 32)
	v2394 = *(*int32)(unsafe.Add(mBase, uint32(v616)+84))
	v2396 = *(*int32)(unsafe.Add(mBase, uint32(v2393+v2394)))
	v2397 = *(*int32)(unsafe.Add(mBase, uint32(v616)+68))
	v2398 = *(*int32)(unsafe.Add(mBase, uint32(v707)+68))
	v2399 = F_adjust_relid_set(m, v2396, v2397, v2398)
	mBase = m.M
	v2400 = m.ExcPending
	if v2400 != 0 {
		goto L18
	} else {
		goto L459
	}
L458:
	;
	goto L456
L459:
	;
	v2401 = *(*int32)(unsafe.Add(mBase, uint32(v616)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v2401+v2393))) = v2399
	v2404 = *(*int32)(unsafe.Add(mBase, uint32(v707)+84))
	v2406 = *(*int32)(unsafe.Add(mBase, uint32(v2404+v2393)))
	v2407 = *(*int32)(unsafe.Add(mBase, uint32(v616)+84))
	v2409 = *(*int32)(unsafe.Add(mBase, uint32(v2407+v2393)))
	v2410 = F_bms_add_members(m, v2406, v2409)
	mBase = m.M
	v2411 = m.ExcPending
	if v2411 != 0 {
		goto L18
	} else {
		goto L460
	}
L460:
	;
	v2412 = *(*int32)(unsafe.Add(mBase, uint32(v707)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v2412+v2393))) = v2410
	v2415 = int32(*(*int16)(unsafe.Add(mBase, uint32(v707)+82)))
	if v2363 < v2415 {
		v2363 = v2363 + int32(1)
		goto L457
	} else {
		goto L461
	}
L461:
	;
	goto L458
L462:
	;
	v2457 = *(*int32)(unsafe.Add(mBase, uint32(v584)+4))
	v2458 = *(*int32)(unsafe.Add(mBase, uint32(v616)+68))
	v2459 = *(*int32)(unsafe.Add(mBase, uint32(v707)+68))
	F_ChangeVarNodesExtended(m, v2457, v2458, v2459, int32(828))
	mBase = m.M
	v2462 = m.ExcPending
	if v2462 != 0 {
		goto L18
	} else {
		goto L468
	}
L463:
	;
	if v887 != 0 {
		goto L464
	} else {
		goto L465
	}
L464:
	;
	v2449 = *(*int32)(unsafe.Add(mBase, uint32(v584)+136))
	v2450 = F_list_delete_ptr(m, v2449, v883)
	mBase = m.M
	v2451 = m.ExcPending
	if v2451 != 0 {
		goto L18
	} else {
		goto L467
	}
L465:
	;
	goto L466
L466:
	;
	v2453 = *(*int32)(unsafe.Add(mBase, uint32(v707)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v883)+4)) = v2453
	*(*int32)(unsafe.Add(mBase, uint32(v883)+8)) = v2453
	goto L462
L467:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v584)+136)) = v2450
	goto L462
L468:
	;
	v2463 = *(*int32)(unsafe.Add(mBase, uint32(v707)+68))
	v2464 = int32(0)
	F_remove_rel_from_query(m, v584, v616, v2463, v2464, v2464)
	mBase = m.M
	v2467 = m.ExcPending
	if v2467 != 0 {
		goto L18
	} else {
		goto L469
	}
L469:
	;
	v2468 = *(*int32)(unsafe.Add(mBase, uint32(v584)+264))
	v2469 = *(*int32)(unsafe.Add(mBase, uint32(v616)+68))
	v2470 = *(*int32)(unsafe.Add(mBase, uint32(v707)+68))
	F_ChangeVarNodesExtended(m, v2468, v2469, v2470, int32(828))
	mBase = m.M
	v2473 = m.ExcPending
	if v2473 != 0 {
		goto L18
	} else {
		goto L470
	}
L470:
	;
	v2474 = *(*int32)(unsafe.Add(mBase, uint32(v584)+256))
	v2475 = *(*int32)(unsafe.Add(mBase, uint32(v616)+68))
	v2476 = *(*int32)(unsafe.Add(mBase, uint32(v707)+68))
	F_ChangeVarNodesExtended(m, v2474, v2475, v2476, int32(828))
	mBase = m.M
	v2479 = m.ExcPending
	if v2479 != 0 {
		goto L18
	} else {
		goto L471
	}
L471:
	;
	v2480 = *(*int32)(unsafe.Add(mBase, uint32(v584)+120))
	v2481 = *(*int32)(unsafe.Add(mBase, uint32(v616)+68))
	v2482 = *(*int32)(unsafe.Add(mBase, uint32(v707)+68))
	v2483 = F_adjust_relid_set(m, v2480, v2481, v2482)
	mBase = m.M
	v2484 = m.ExcPending
	if v2484 != 0 {
		goto L18
	} else {
		goto L472
	}
L472:
	;
	v2485 = *(*int32)(unsafe.Add(mBase, uint32(v584)+124))
	v2486 = *(*int32)(unsafe.Add(mBase, uint32(v616)+68))
	v2487 = *(*int32)(unsafe.Add(mBase, uint32(v707)+68))
	v2488 = F_adjust_relid_set(m, v2485, v2486, v2487)
	mBase = m.M
	v2489 = m.ExcPending
	if v2489 != 0 {
		goto L18
	} else {
		goto L473
	}
L473:
	;
	v2490 = *(*int32)(unsafe.Add(mBase, uint32(v584)+28))
	v2491 = *(*int32)(unsafe.Add(mBase, uint32(v616)+68))
	v2492 = int32(2)
	v2495 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2490+v2491<<(uint(v2492)%32)))) = v2495
	v2497 = *(*int32)(unsafe.Add(mBase, uint32(v584)+36))
	v2498 = *(*int32)(unsafe.Add(mBase, uint32(v616)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v2497+v2498<<(uint(v2492)%32)))) = v2495
	F_pfree(m, v616)
	mBase = m.M
	v2505 = m.ExcPending
	if v2505 != 0 {
		goto L18
	} else {
		goto L474
	}
L474:
	;
	F_rebuild_placeholder_attr_needed(m, v584)
	mBase = m.M
	v2507 = m.ExcPending
	if v2507 != 0 {
		goto L18
	} else {
		goto L475
	}
L475:
	;
	F_rebuild_joinclause_attr_needed(m, v584)
	mBase = m.M
	v2509 = m.ExcPending
	if v2509 != 0 {
		goto L18
	} else {
		goto L476
	}
L476:
	;
	F_rebuild_eclass_attr_needed(m, v584)
	mBase = m.M
	v2511 = m.ExcPending
	if v2511 != 0 {
		goto L18
	} else {
		goto L477
	}
L477:
	;
	F_rebuild_lateral_attr_needed(m, v584)
	mBase = m.M
	v2513 = m.ExcPending
	if v2513 != 0 {
		goto L18
	} else {
		goto L478
	}
L478:
	;
	v2514 = F_bms_add_member(m, v607, v602)
	mBase = m.M
	v2515 = m.ExcPending
	if v2515 != 0 {
		goto L18
	} else {
		goto L479
	}
L479:
	;
	v2625 = v2514
	goto L107
L480:
	;
	if int32(0) < v2599 {
		v683 = v2599
		goto L120
	} else {
		goto L491
	}
L481:
	;
	v2599 = base.I32_ctz(v2585) | v2586<<(uint(int32(5))%32)
	goto L480
L482:
	;
	v2599 = int32(-2)
	goto L480
L483:
	;
	v2550 = v683 + int32(1)
	v2552 = base.I32_div_s(v2550, int32(32))
	v2553 = *(*int32)(unsafe.Add(mBase, uint32(v599)+4))
	if v2553 <= v2552 {
		goto L482
	} else {
		goto L484
	}
L484:
	;
	v2556 = v599 + int32(8)
	v2560 = *(*int32)(unsafe.Add(mBase, uint32(v2556+v2552<<(uint(int32(2))%32))))
	v2563 = v2560 & (int32(-1) << (uint(v2550) % 32))
	if v2563 != 0 {
		v2585 = v2563
		v2586 = v2552
		goto L481
	} else {
		goto L485
	}
L485:
	;
	v2565 = v2552 + int32(1)
	if v2565 == v2553 {
		goto L482
	} else {
		goto L486
	}
L486:
	;
	v2568 = v2565
	goto L487
L487:
	;
	v2575 = *(*int32)(unsafe.Add(mBase, uint32(v2556+v2568<<(uint(int32(2))%32))))
	if v2575 != 0 {
		v2585 = v2575
		v2586 = v2568
		goto L481
	} else {
		goto L489
	}
L488:
	;
	goto L482
L489:
	;
	v2577 = v2568 + int32(1)
	if v2577 != v2553 {
		v2568 = v2577
		goto L487
	} else {
		goto L490
	}
L490:
	;
	goto L488
L491:
	;
	goto L121
L492:
	;
	if int32(0) < v2685 {
		v602 = v2685
		v607 = v2625
		goto L105
	} else {
		goto L503
	}
L493:
	;
	v2685 = base.I32_ctz(v2671) | v2672<<(uint(int32(5))%32)
	goto L492
L494:
	;
	v2685 = int32(-2)
	goto L492
L495:
	;
	v2636 = v602 + int32(1)
	v2638 = base.I32_div_s(v2636, int32(32))
	v2639 = *(*int32)(unsafe.Add(mBase, uint32(v599)+4))
	if v2639 <= v2638 {
		goto L494
	} else {
		goto L496
	}
L496:
	;
	v2642 = v599 + int32(8)
	v2646 = *(*int32)(unsafe.Add(mBase, uint32(v2642+v2638<<(uint(int32(2))%32))))
	v2649 = v2646 & (int32(-1) << (uint(v2636) % 32))
	if v2649 != 0 {
		v2671 = v2649
		v2672 = v2638
		goto L493
	} else {
		goto L497
	}
L497:
	;
	v2651 = v2638 + int32(1)
	if v2651 == v2639 {
		goto L494
	} else {
		goto L498
	}
L498:
	;
	v2654 = v2651
	goto L499
L499:
	;
	v2661 = *(*int32)(unsafe.Add(mBase, uint32(v2642+v2654<<(uint(int32(2))%32))))
	if v2661 != 0 {
		v2671 = v2661
		v2672 = v2654
		goto L493
	} else {
		goto L501
	}
L500:
	;
	goto L494
L501:
	;
	v2663 = v2654 + int32(1)
	if v2663 != v2639 {
		v2654 = v2663
		goto L499
	} else {
		goto L502
	}
L502:
	;
	goto L500
L503:
	;
	goto L106
L504:
	;
	v2718 = F_bms_del_members(m, v2703, v2711)
	mBase = m.M
	v2719 = m.ExcPending
	if v2719 != 0 {
		goto L18
	} else {
		goto L505
	}
L505:
	;
	if v2711 != 0 {
		goto L506
	} else {
		goto L507
	}
L506:
	;
	v2720 = int32(0)
	if v2718 == v2720 {
		goto L510
	} else {
		goto L511
	}
L507:
	;
	goto L508
L508:
	;
	goto L90
L509:
	;
	if v2765 == int32(2) {
		v496 = v2688
		v498 = v2716
		v507 = v2699
		v511 = v2718
		v513 = v2705
		v515 = v2707
		v517 = v2709
		v518 = v2710
		v520 = v2712
		goto L89
	} else {
		goto L525
	}
L510:
	;
	v2765 = int32(0)
	goto L509
L511:
	;
	goto L512
L512:
	;
	v2728 = int32(1)
	v2729 = *(*int32)(unsafe.Add(mBase, uint32(v2718)+4))
	if v2729 <= v2728 {
		goto L513
	} else {
		goto L514
	}
L513:
	;
	v2732 = v2728
	goto L515
L514:
	;
	v2732 = v2729
	goto L515
L515:
	;
	v2736 = int32(0)
	v2738 = v2720
	goto L516
L516:
	;
	v2745 = *(*int32)(unsafe.Add(mBase, uint32(v2718+int32(8)+v2736<<(uint(int32(2))%32))))
	if v2745 != 0 {
		goto L519
	} else {
		goto L520
	}
L517:
	;
	v2765 = v2757
	goto L509
L518:
	;
	goto L517
L519:
	;
	v2746 = int32(2)
	if v2738 != 0 {
		v2757 = v2746
		goto L518
	} else {
		goto L522
	}
L520:
	;
	v2752 = v2738
	goto L521
L521:
	;
	v2754 = v2736 + int32(1)
	if v2754 != v2732 {
		v2736 = v2754
		v2738 = v2752
		goto L516
	} else {
		goto L524
	}
L522:
	;
	v2747 = int32(1)
	if base.Ui32(v2747) < base.Ui32(base.I32_popcnt(v2745)) {
		v2757 = v2746
		goto L518
	} else {
		goto L523
	}
L523:
	;
	v2752 = v2747
	goto L521
L524:
	;
	v2757 = v2752
	goto L518
L525:
	;
	goto L508
L526:
	;
	F_bms_free(m, v2718)
	mBase = m.M
	v2771 = m.ExcPending
	if v2771 != 0 {
		goto L18
	} else {
		goto L527
	}
L527:
	;
	v2778 = v2688
	v2780 = v2716
	v2789 = v2699
	v2795 = v2705
	v2797 = v2707
	v2799 = v2709
	v2800 = v2710
	v2802 = v2712
	goto L73
L528:
	;
	v2778 = v386
	v2780 = v388
	v2789 = v389
	v2795 = v389
	v2797 = v2776
	v2799 = v407
	v2800 = v408
	v2802 = v410
	goto L73
L529:
	;
	goto L72
}
func F_removeabbrev_cluster(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	v4 = int32(0)
	if v4 < l2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v12 = v4
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v16 = l1 + v12<<(uint(int32(4))%32)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v19 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+12)))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v23 = F_heap_getattr_1(m, v17, v19, v20, v16+int32(8))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	return
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v23
	v27 = v12 + int32(1)
	if v27 != l2 {
		v12 = v27
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
}
func F_removecaptures(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v5&int32(32) == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(0)
	v13 = v5 & int32(215)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)) = uint8(v13)
	v15 = v13
	goto L3
L2:
	;
	v15 = v5
	goto L3
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v16 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v20 = v16
	goto L7
L5:
	;
	v34 = v15
	goto L6
L6:
	;
	if v34&int32(24) == int32(0) {
		goto L15
	} else {
		goto L16
	}
L7:
	;
	F_removecaptures(m, l0, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	v34 = v31
	goto L6
L9:
	;
	return
L10:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	if v23&int32(8) != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	v28 = v26 | int32(8)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)) = uint8(v28)
	goto L13
L12:
	;
	goto L13
L13:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	if v30 != 0 {
		v20 = v30
		goto L7
	} else {
		goto L14
	}
L14:
	;
	goto L8
L15:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v40 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	return
L18:
	;
	v44 = v40
	goto L21
L19:
	;
	v51 = v34
	goto L20
L20:
	;
	v53 = int32(61)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v53)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = int32(0)
	v58 = v51 & int32(251)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)) = uint8(v58)
	goto L17
L21:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+24))
	F_freesubre(m, l0, v44)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L9
	} else {
		goto L23
	}
L22:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	v51 = v48
	goto L20
L23:
	;
	if v45 != 0 {
		v44 = v45
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
}
func F_renameatt_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	v13 = m.G0
	v15 = v13 + int32(-64)
	m.G0 = v15
	v18 = F_relation_open(m, l0, int32(8))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
	F_renameatt_check(m, l0, v22, l4)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if l3 != 0 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L64
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L60
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L56
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L52
	}
L8:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+119)))
	if v93 != int32(99) {
		goto L28
	} else {
		goto L29
	}
L9:
	;
	v28 = F_find_all_inheritors(m, l0, int32(8), v13+int32(-4))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	if l5 != 0 {
		goto L8
	} else {
		goto L25
	}
L12:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
	v36 = int32(0)
	goto L13
L13:
	;
	v44 = int32(0)
	if v28 == v44 {
		v54 = v44
		goto L15
	} else {
		goto L16
	}
L15:
	;
	if v30 == int32(0) {
		goto L8
	} else {
		goto L18
	}
L16:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v48 <= v36 {
		v54 = int32(0)
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v54 = v50 + v36<<(uint(int32(2))%32)
	goto L15
L18:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if base.B2i32(v54 == int32(0))|base.B2i32(v59 <= v36) != 0 {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	if v62 == int32(0) {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	if l0 != v65 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v62+v36<<(uint(int32(2))%32))))
	v73 = F_renameatt_internal(m, v65, l1, l2, int32(0), int32(1), v72, l6)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v36 = v36 + int32(1)
	goto L13
L24:
	;
	goto L23
L25:
	;
	v78 = F_find_inheritance_children(m, l0, int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	if v78 != 0 {
		goto L7
	} else {
		goto L27
	}
L27:
	;
	goto L8
L28:
	;
	v147 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L37
	}
L29:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v92)+72))
	v99 = F_find_typed_table_dependencies(m, v96, v92+int32(4), l6)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	if v99 == int32(0) {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	if v103 <= int32(0) {
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v111 = int32(0)
	goto L33
L33:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v119+v111<<(uint(int32(2))%32))))
	v124 = int32(1)
	v127 = F_renameatt_internal(m, v123, l1, l2, v124, v124, int32(0), l6)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	goto L28
L35:
	;
	v130 = v111 + int32(1)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	if v130 < v131 {
		v111 = v130
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v149 = F_SearchSysCacheCopyAttName(m, l0, l1)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if v149 == int32(0) {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v149)+16))
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+22)))
	v155 = v153 + v154
	v156 = int32(*(*int16)(unsafe.Add(mBase, uint32(v155)+74)))
	if v156 <= int32(0) {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	v159 = int32(*(*int16)(unsafe.Add(mBase, uint32(v155)+94)))
	if l5 < v159 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	v162 = F_check_for_column_name_collision(m, v18, l2, int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v167 = F_strncpy(m, v155+int32(4), l2, int32(64))
	mBase = m.M
	v168 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v167)+63)) = uint8(v168)
	goto L43
L43:
	;
	F_CatalogTupleUpdate(m, v147, v149+int32(4), v149)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v175 = *(*int32)(unsafe.Add(mBase, _c_F_renameatt_internal[0]))
	if v175 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v177 = int32(0)
	F_RunObjectPostAlterHook(m, int32(1259), l0, v156, v177, v177)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	F_pfree(m, v149)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L49
	}
L48:
	;
	goto L47
L49:
	;
	F_relation_close(m, v147, int32(3))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_relation_close(m, v18, int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	m.G0 = v15 - int32(-64)
	return v156
L52:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = l1
	F_errmsg(m, int32(_a_F_renameatt_internal_0), v13+int32(-16))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_renameatt_internal_1), int32(3917), int32(_a_F_renameatt_internal_2))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l1
	F_errmsg(m, int32(_a_F_renameatt_internal_3), v15)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_renameatt_internal_1), int32(3941), int32(_a_F_renameatt_internal_2))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
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
	F_errcode(m, int32(1088))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l1
	F_errmsg(m, int32(_a_F_renameatt_internal_4), v13+int32(-48))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_renameatt_internal_1), int32(3949), int32(_a_F_renameatt_internal_2))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = l1
	F_errmsg(m, int32(_a_F_renameatt_internal_5), v13+int32(-32))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_renameatt_internal_1), int32(3964), int32(_a_F_renameatt_internal_2))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_renametrig_partition(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
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
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v93 int32
	_ = v93
	var v104 int32
	_ = v104
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	F_ScanKeyInit(m, v12, int32(2), int32(3), int32(184), l1)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v20 = int32(1)
	v23 = F_systable_beginscan(m, l0, int32(2701), v20, int32(0), v20, v12)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	goto L5
L4:
	;
	F_systable_endscan(m, v23)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L21
	}
L5:
	;
	v34 = F_systable_getnext(m, v23)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	v44 = F_table_open(m, l1, int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L10
	}
L7:
	;
	if v34 == int32(0) {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+22)))
	v40 = v38 + v39
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	if v41 != l2 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	goto L6
L10:
	;
	F_renametrig_internal(m, l0, v44, v34, l3, l4)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v44)+48))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+119)))
	if v49 != int32(112) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	F_relation_close(m, v44, int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L20
	}
L13:
	;
	v53 = F_RelationGetPartitionDesc(m, v44, int32(1))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	if v55 <= int32(0) {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v63 = int32(0)
	goto L16
L16:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v70+v63<<(uint(int32(2))%32))))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	F_renametrig_partition(m, l0, v74, v75, l3, v40+int32(12))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L18
	}
L17:
	;
	goto L12
L18:
	;
	v79 = v63 + int32(1)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	if v79 < v80 {
		v63 = v79
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	goto L4
L21:
	;
	m.G0 = v12 + int32(48)
	return
}
func F_replace_correlation_vars_mutator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
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
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
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
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v655 int32
	_ = v655
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	v3 = int32(0)
	if l0 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v12 - int32(6) {
	case 0:
		goto L9
	case 1, 2, 5, 6:
		goto L4
	case 3:
		goto L8
	case 4:
		goto L7
	case 7:
		goto L6
	default:
		goto L10
	}
L4:
	;
	v781 = F_expression_tree_mutator_impl(m, l0, int32(846), l1)
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L39
	} else {
		goto L195
	}
L5:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v671 <= int32(0) {
		goto L4
	} else {
		goto L169
	}
L6:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v601)+4))
	if v602 == int32(5) {
		goto L4
	} else {
		goto L150
	}
L7:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v500 == int32(0) {
		goto L4
	} else {
		goto L126
	}
L8:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v397 == int32(0) {
		goto L4
	} else {
		goto L103
	}
L9:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v180 == int32(0) {
		goto L4
	} else {
		goto L53
	}
L10:
	;
	if v12 == int32(61) {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	if v12 != int32(319) {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v19 == int32(0) {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v22 == int32(0) {
		v70 = l1
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v70)+20))
	if v75 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L15:
	;
	v26 = v22 & int32(7)
	if v26 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if base.Ui32(v22) < base.Ui32(int32(8)) {
		v70 = v44
		goto L14
	} else {
		goto L23
	}
L17:
	;
	v42 = v22
	v44 = l1
	goto L16
L18:
	;
	goto L19
L19:
	;
	v29 = v22
	v31 = l1
	v33 = v3
	goto L20
L20:
	;
	v36 = int32(1)
	v37 = v29 - v36
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	v40 = v33 + v36
	if v40 != v26 {
		v29 = v37
		v31 = v38
		v33 = v40
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v42 = v37
	v44 = v38
	goto L16
L22:
	;
	goto L21
L23:
	;
	v51 = v42
	v53 = v44
	goto L24
L24:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
	v67 = v51 - int32(8)
	if v67 != 0 {
		v51 = v67
		v53 = v65
		goto L24
	} else {
		goto L26
	}
L25:
	;
	v70 = v65
	goto L14
L26:
	;
	goto L25
L27:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	v160 = F_palloc0(m, int32(28))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L39
	} else {
		goto L49
	}
L28:
	;
	v157 = v93 + int32(8)
	goto L27
L29:
	;
	v111 = F_copyObjectImpl(m, l0)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L39
	} else {
		goto L40
	}
L30:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	if v78 <= int32(0) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
	v83 = int32(0)
	goto L32
L32:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v81+v83<<(uint(int32(2))%32))))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	if v95 == int32(319) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L29
L34:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v94)+16))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v98 == v99 {
		goto L28
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v102 = v83 + int32(1)
	if v78 != v102 {
		v83 = v102
		goto L32
	} else {
		goto L38
	}
L37:
	;
	goto L36
L38:
	;
	goto L33
L39:
	;
	return int32(0)
L40:
	;
	v115 = int32(0)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	F_IncrementVarSublevelsUp(m, v111, v115-v116, v115)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v122 = F_palloc0(m, int32(12))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L39
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v122)+4)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v122))) = int32(326)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+64))
	if v128 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	v131 = v129
	goto L45
L44:
	;
	v131 = int32(0)
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v122)+8)) = v131
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)+64))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v136 = F_exprType(m, v135)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L39
	} else {
		goto L46
	}
L46:
	;
	v138 = F_lappend_oid(m, v134, v136)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L39
	} else {
		goto L47
	}
L47:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v140)+64)) = v138
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v70)+20))
	v143 = F_lappend(m, v142, v122)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L39
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+20)) = v143
	v157 = v122 + int32(8)
	goto L27
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160)+8)) = v158
	*(*int64)(unsafe.Add(mBase, uint32(v160))) = int64(4294967304)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v166 = F_exprType(m, v165)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L39
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160)+12)) = v166
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v170 = F_exprTypmod(m, v169)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L39
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160)+16)) = v170
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v174 = F_exprCollation(m, v173)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L39
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160)+24)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v160)+20)) = v174
	return v160
L53:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v183 == int32(0) {
		v231 = l1
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v231)+20))
	if v236 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L55:
	;
	v187 = v183 & int32(7)
	if v187 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	if base.Ui32(v183) < base.Ui32(int32(8)) {
		v231 = v205
		goto L54
	} else {
		goto L63
	}
L57:
	;
	v203 = v183
	v205 = l1
	goto L56
L58:
	;
	goto L59
L59:
	;
	v190 = v183
	v192 = l1
	v194 = v3
	goto L60
L60:
	;
	v197 = int32(1)
	v198 = v190 - v197
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v192)+16))
	v201 = v194 + v197
	if v201 != v187 {
		v190 = v198
		v192 = v199
		v194 = v201
		goto L60
	} else {
		goto L62
	}
L61:
	;
	v203 = v198
	v205 = v199
	goto L56
L62:
	;
	goto L61
L63:
	;
	v212 = v203
	v214 = v205
	goto L64
L64:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v214)+16))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)+16))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+16))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)+16))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+16))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+16))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)+16))
	v228 = v212 - int32(8)
	if v228 != 0 {
		v212 = v228
		v214 = v226
		goto L64
	} else {
		goto L66
	}
L65:
	;
	v231 = v226
	goto L54
L66:
	;
	goto L65
L67:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v380)))
	v383 = F_palloc0(m, int32(28))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L39
	} else {
		goto L102
	}
L68:
	;
	v380 = v254 + int32(8)
	goto L67
L69:
	;
	v342 = F_copyObjectImpl(m, l0)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L39
	} else {
		goto L95
	}
L70:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v236)+4))
	if v239 <= int32(0) {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v243 = int32(0)
	v244 = v239
	goto L72
L72:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v236)+12))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v250+v243<<(uint(int32(2))%32))))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v254)+4))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	if v256 != int32(6) {
		v331 = v244
		goto L74
	} else {
		goto L75
	}
L73:
	;
	goto L69
L74:
	;
	v333 = v243 + int32(1)
	if v333 < v331 {
		v243 = v333
		v244 = v331
		goto L72
	} else {
		goto L94
	}
L75:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v255)+4))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v259 != v260 {
		v331 = v244
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v262 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v255)+8)))
	v263 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
	if v262 != v263 {
		v331 = v244
		goto L74
	} else {
		goto L77
	}
L77:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v255)+12))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v265 != v266 {
		v331 = v244
		goto L74
	} else {
		goto L78
	}
L78:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v255)+16))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v268 != v269 {
		v331 = v244
		goto L74
	} else {
		goto L79
	}
L79:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v255)+20))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v271 != v272 {
		v331 = v244
		goto L74
	} else {
		goto L80
	}
L80:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v255)+32))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v274 != v275 {
		v331 = v244
		goto L74
	} else {
		goto L81
	}
L81:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v255)+24))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v279 = int32(0)
	if base.B2i32(v277 == v279)|base.B2i32(v278 == v279) != 0 {
		v325 = base.B2i32(v277|v278 == v279)
		goto L83
	} else {
		goto L84
	}
L82:
	;
	if v325 != 0 {
		goto L68
	} else {
		goto L93
	}
L83:
	;
	goto L82
L84:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v277)+4))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v278)+4))
	if v293 != v294 {
		v325 = int32(0)
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v296 = int32(1)
	if v293 <= v296 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v299 = v296
	goto L88
L87:
	;
	v299 = v293
	goto L88
L88:
	;
	v300 = int32(8)
	v305 = int32(0)
	goto L89
L89:
	;
	v313 = v305 << (uint(int32(2)) % 32)
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v277+v300+v313)))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v278+v300+v313)))
	v318 = base.B2i32(v315 == v317)
	if v315 != v317 {
		v325 = v318
		goto L83
	} else {
		goto L91
	}
L90:
	;
	v325 = v318
	goto L83
L91:
	;
	v321 = v305 + int32(1)
	if v321 != v299 {
		v305 = v321
		goto L89
	} else {
		goto L92
	}
L92:
	;
	goto L90
L93:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v236)+4))
	v331 = v330
	goto L74
L94:
	;
	goto L73
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v342)+28)) = int32(0)
	v347 = F_palloc0(m, int32(12))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L39
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v347)+4)) = v342
	*(*int32)(unsafe.Add(mBase, uint32(v347))) = int32(326)
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v231)+8))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v352)+64))
	if v353 != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v353)+4))
	v356 = v354
	goto L99
L98:
	;
	v356 = int32(0)
	goto L99
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v347)+8)) = v356
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v231)+8))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v358)+64))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v342)+12))
	v361 = F_lappend_oid(m, v359, v360)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L39
	} else {
		goto L100
	}
L100:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v231)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v363)+64)) = v361
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v231)+20))
	v366 = F_lappend(m, v365, v347)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L39
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+20)) = v366
	v380 = v347 + int32(8)
	goto L67
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v383)+8)) = v381
	*(*int64)(unsafe.Add(mBase, uint32(v383))) = int64(4294967304)
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v383)+12)) = v388
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v383)+16)) = v390
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v383)+20)) = v392
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v383)+24)) = v394
	return v383
L103:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v400 == int32(0) {
		v448 = l1
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v453 = F_copyObjectImpl(m, l0)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L39
	} else {
		goto L117
	}
L105:
	;
	v404 = v400 & int32(7)
	if v404 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	if base.Ui32(v400) < base.Ui32(int32(8)) {
		v448 = v422
		goto L104
	} else {
		goto L113
	}
L107:
	;
	v420 = v400
	v422 = l1
	goto L106
L108:
	;
	goto L109
L109:
	;
	v407 = v400
	v409 = l1
	v411 = v3
	goto L110
L110:
	;
	v414 = int32(1)
	v415 = v407 - v414
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v409)+16))
	v418 = v411 + v414
	if v418 != v404 {
		v407 = v415
		v409 = v416
		v411 = v418
		goto L110
	} else {
		goto L112
	}
L111:
	;
	v420 = v415
	v422 = v416
	goto L106
L112:
	;
	goto L111
L113:
	;
	v429 = v420
	v431 = v422
	goto L114
L114:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v431)+16))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v436)+16))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v437)+16))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v438)+16))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v439)+16))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v440)+16))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v441)+16))
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v442)+16))
	v445 = v429 - int32(8)
	if v445 != 0 {
		v429 = v445
		v431 = v443
		goto L114
	} else {
		goto L116
	}
L115:
	;
	v448 = v443
	goto L104
L116:
	;
	goto L115
L117:
	;
	v455 = int32(0)
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v453)+52))
	F_IncrementVarSublevelsUp(m, v453, v455-v456, v455)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L39
	} else {
		goto L118
	}
L118:
	;
	v462 = F_palloc0(m, int32(12))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L39
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v462)+4)) = v453
	*(*int32)(unsafe.Add(mBase, uint32(v462))) = int32(326)
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v448)+8))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v467)+64))
	if v468 != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v468)+4))
	v471 = v469
	goto L122
L121:
	;
	v471 = int32(0)
	goto L122
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v462)+8)) = v471
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v448)+8))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v473)+64))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v453)+8))
	v476 = F_lappend_oid(m, v474, v475)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L39
	} else {
		goto L123
	}
L123:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v448)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v478)+64)) = v476
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v448)+20))
	v481 = F_lappend(m, v480, v462)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L39
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v448)+20)) = v481
	v485 = F_palloc0(m, int32(28))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L39
	} else {
		goto L125
	}
L125:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v485))) = int64(4294967304)
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v462)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v485)+8)) = v489
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v453)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v485)+16)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v485)+12)) = v491
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v453)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v485)+20)) = v495
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v453)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v485)+24)) = v497
	return v485
L126:
	;
	v503 = F_exprType(m, l0)
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L39
	} else {
		goto L127
	}
L127:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v505 == int32(0) {
		v553 = l1
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v558 = F_copyObjectImpl(m, l0)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L39
	} else {
		goto L141
	}
L129:
	;
	v509 = v505 & int32(7)
	if v509 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	if base.Ui32(v505) < base.Ui32(int32(8)) {
		v553 = v527
		goto L128
	} else {
		goto L137
	}
L131:
	;
	v525 = v505
	v527 = l1
	goto L130
L132:
	;
	goto L133
L133:
	;
	v512 = v505
	v514 = l1
	v516 = v3
	goto L134
L134:
	;
	v519 = int32(1)
	v520 = v512 - v519
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v514)+16))
	v523 = v516 + v519
	if v523 != v509 {
		v512 = v520
		v514 = v521
		v516 = v523
		goto L134
	} else {
		goto L136
	}
L135:
	;
	v525 = v520
	v527 = v521
	goto L130
L136:
	;
	goto L135
L137:
	;
	v534 = v525
	v536 = v527
	goto L138
L138:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v536)+16))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v541)+16))
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v542)+16))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v543)+16))
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v544)+16))
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v545)+16))
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v546)+16))
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v547)+16))
	v550 = v534 - int32(8)
	if v550 != 0 {
		v534 = v550
		v536 = v548
		goto L138
	} else {
		goto L140
	}
L139:
	;
	v553 = v548
	goto L128
L140:
	;
	goto L139
L141:
	;
	v560 = int32(0)
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v558)+16))
	F_IncrementVarSublevelsUp(m, v558, v560-v561, v560)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L39
	} else {
		goto L142
	}
L142:
	;
	v567 = F_palloc0(m, int32(12))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L39
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v567)+4)) = v558
	*(*int32)(unsafe.Add(mBase, uint32(v567))) = int32(326)
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v553)+8))
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v572)+64))
	if v573 != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v573)+4))
	v576 = v574
	goto L146
L145:
	;
	v576 = int32(0)
	goto L146
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v567)+8)) = v576
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v553)+8))
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v578)+64))
	v580 = F_lappend_oid(m, v579, v503)
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L39
	} else {
		goto L147
	}
L147:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v553)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v582)+64)) = v580
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v553)+20))
	v585 = F_lappend(m, v584, v567)
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L39
	} else {
		goto L148
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v553)+20)) = v585
	v589 = F_palloc0(m, int32(28))
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L39
	} else {
		goto L149
	}
L149:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v589))) = int64(4294967304)
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v567)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v589)+16)) = int64(4294967295)
	*(*int32)(unsafe.Add(mBase, uint32(v589)+12)) = v503
	*(*int32)(unsafe.Add(mBase, uint32(v589)+8)) = v593
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v558)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v589)+24)) = v598
	return v589
L150:
	;
	v605 = F_exprType(m, l0)
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L39
	} else {
		goto L152
	}
L151:
	;
	return v646
L152:
	;
	v608 = l1
	goto L154
L153:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L39
	} else {
		goto L166
	}
L154:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v608)+16))
	if v614 == int32(0) {
		goto L153
	} else {
		goto L156
	}
L155:
	;
	v621 = F_copyObjectImpl(m, l0)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L39
	} else {
		goto L158
	}
L156:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v614)+4))
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v617)+4))
	if v618 != int32(5) {
		v608 = v614
		goto L154
	} else {
		goto L157
	}
L157:
	;
	goto L155
L158:
	;
	v624 = F_palloc0(m, int32(12))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L39
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v624)+4)) = v621
	*(*int32)(unsafe.Add(mBase, uint32(v624))) = int32(326)
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v614)+8))
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v629)+64))
	if v630 != 0 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v630)+4))
	v633 = v631
	goto L162
L161:
	;
	v633 = int32(0)
	goto L162
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v624)+8)) = v633
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v614)+8))
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v635)+64))
	v637 = F_lappend_oid(m, v636, v605)
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L39
	} else {
		goto L163
	}
L163:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v614)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v639)+64)) = v637
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v614)+20))
	v642 = F_lappend(m, v641, v624)
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L39
	} else {
		goto L164
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v614)+20)) = v642
	v646 = F_palloc0(m, int32(28))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L39
	} else {
		goto L165
	}
L165:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v646))) = int64(4294967304)
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v624)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v646)+16)) = int64(4294967295)
	*(*int32)(unsafe.Add(mBase, uint32(v646)+12)) = v605
	*(*int32)(unsafe.Add(mBase, uint32(v646)+8)) = v650
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v621)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v646)+24)) = v655
	goto L151
L166:
	;
	F_errmsg_internal(m, int32(_a_F_replace_correlation_vars_mutator_0), int32(0))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L39
	} else {
		goto L167
	}
L167:
	;
	F_errfinish(m, int32(_a_F_replace_correlation_vars_mutator_1), int32(334), int32(_a_F_replace_correlation_vars_mutator_2))
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L39
	} else {
		goto L168
	}
L168:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L169:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v675 = F_exprType(m, v674)
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L39
	} else {
		goto L170
	}
L170:
	;
	v677 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v677 == int32(0) {
		v725 = l1
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v730 = F_copyObjectImpl(m, l0)
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L39
	} else {
		goto L184
	}
L172:
	;
	v681 = v677 & int32(7)
	if v681 == int32(0) {
		goto L174
	} else {
		goto L175
	}
L173:
	;
	if base.Ui32(v677) < base.Ui32(int32(8)) {
		v725 = v699
		goto L171
	} else {
		goto L180
	}
L174:
	;
	v697 = v677
	v699 = l1
	goto L173
L175:
	;
	goto L176
L176:
	;
	v684 = v677
	v686 = l1
	v688 = v3
	goto L177
L177:
	;
	v691 = int32(1)
	v692 = v684 - v691
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v686)+16))
	v695 = v688 + v691
	if v695 != v681 {
		v684 = v692
		v686 = v693
		v688 = v695
		goto L177
	} else {
		goto L179
	}
L178:
	;
	v697 = v692
	v699 = v693
	goto L173
L179:
	;
	goto L178
L180:
	;
	v706 = v697
	v708 = v699
	goto L181
L181:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v708)+16))
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v713)+16))
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v714)+16))
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v715)+16))
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v716)+16))
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v717)+16))
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v718)+16))
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v719)+16))
	v722 = v706 - int32(8)
	if v722 != 0 {
		v706 = v722
		v708 = v720
		goto L181
	} else {
		goto L183
	}
L182:
	;
	v725 = v720
	goto L171
L183:
	;
	goto L182
L184:
	;
	v732 = int32(0)
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v730)+4))
	F_IncrementVarSublevelsUp(m, v730, v732-v733, v732)
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L39
	} else {
		goto L185
	}
L185:
	;
	v739 = F_palloc0(m, int32(12))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L39
	} else {
		goto L186
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v739)+4)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v739))) = int32(326)
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v725)+8))
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v744)+64))
	if v745 != 0 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v745)+4))
	v748 = v746
	goto L189
L188:
	;
	v748 = int32(0)
	goto L189
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v739)+8)) = v748
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v725)+8))
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v750)+64))
	v752 = F_lappend_oid(m, v751, v675)
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L39
	} else {
		goto L190
	}
L190:
	;
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v725)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v754)+64)) = v752
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v725)+20))
	v757 = F_lappend(m, v756, v739)
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L39
	} else {
		goto L191
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v725)+20)) = v757
	v761 = F_palloc0(m, int32(28))
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L39
	} else {
		goto L192
	}
L192:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v761))) = int64(4294967304)
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v739)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v761)+12)) = v675
	*(*int32)(unsafe.Add(mBase, uint32(v761)+8)) = v765
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v730)+12))
	v769 = F_exprTypmod(m, v768)
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L39
	} else {
		goto L193
	}
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v761)+16)) = v769
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v730)+12))
	v773 = F_exprCollation(m, v772)
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L39
	} else {
		goto L194
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v761)+20)) = v773
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v730)+12))
	v777 = F_exprLocation(m, v776)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v761)+24)) = v777
	return v761
L195:
	;
	return v781
}
func F_replace_nestloop_params_mutator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int64
	_ = v194
	var v196 int64
	_ = v196
	var v198 int64
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v12 != int32(319) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v316 = F_expression_tree_mutator_impl(m, l0, int32(830), l1)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L12
	} else {
		goto L83
	}
L5:
	;
	if v12 != int32(6) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v129 = F_find_placeholder_info(m, l1, l0)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L12
	} else {
		goto L37
	}
L8:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v17 < int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return l0
L10:
	;
	goto L11
L11:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+352))
	v22 = F_bms_is_member(m, v17, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	if v22 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return l0
L15:
	;
	goto L16
L16:
	;
	v29 = int32(0)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+356))
	if v30 == v29 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	return v127
L18:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v83 = F_palloc0(m, int32(28))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L12
	} else {
		goto L29
	}
L19:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v33 <= int32(0) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v38 = v29
	goto L21
L21:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43+v38<<(uint(int32(2))%32))))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	v49 = F_equal(m, l0, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L12
	} else {
		goto L23
	}
L22:
	;
	v58 = F_palloc0(m, int32(28))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L12
	} else {
		goto L28
	}
L23:
	;
	if v49 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v54 = v38 + int32(1)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v54 < v55 {
		v38 = v54
		goto L21
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	goto L22
L27:
	;
	goto L18
L28:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v58))) = int64(4294967304)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = v62
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+12)) = v64
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+16)) = v66
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+20)) = v68
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+24)) = v70
	v127 = v58
	goto L17
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = int64(4294967304)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+64))
	if v88 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	v91 = v89
	goto L32
L31:
	;
	v91 = int32(0)
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = v91
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+64))
	v95 = F_lappend_oid(m, v94, v81)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L12
	} else {
		goto L33
	}
L33:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+64)) = v95
	*(*int32)(unsafe.Add(mBase, uint32(v83)+24)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v83)+20)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v83)+16)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v83)+12)) = v81
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v83)+24)) = v104
	v107 = F_palloc0(m, int32(12))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L12
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = int32(357)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v107)+4)) = v111
	v113 = F_copyObjectImpl(m, l0)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L12
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107)+8)) = v113
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)+356))
	v117 = F_lappend(m, v116, v107)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L12
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+356)) = v117
	v127 = v83
	goto L17
L37:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)+12))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l1)+352))
	v133 = int32(0)
	if v131 == v133 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	if v186 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L39:
	;
	v186 = int32(1)
	goto L38
L40:
	;
	goto L41
L41:
	;
	if v132 == int32(0) {
		v179 = v133
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v186 = v179
	goto L38
L43:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	if v143 < v142 {
		v179 = v133
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v145 = int32(1)
	if v142 <= v145 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v148 = v145
	goto L47
L46:
	;
	v148 = v142
	goto L47
L47:
	;
	v149 = int32(8)
	v154 = int32(0)
	goto L48
L48:
	;
	v161 = v154 << (uint(int32(2)) % 32)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v131+v149+v161)))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v132+v149+v161)))
	v168 = v163 & (v165 ^ int32(-1))
	v170 = base.B2i32(v168 == int32(0))
	if v168 != 0 {
		v179 = v170
		goto L42
	} else {
		goto L50
	}
L49:
	;
	v179 = v170
	goto L42
L50:
	;
	v172 = v154 + int32(1)
	if v172 != v148 {
		v154 = v172
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v190 = F_palloc0(m, int32(24))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L12
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v205 = int32(0)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l1)+356))
	if v206 == v205 {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v190))) = int32(319)
	v194 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v190)+8)) = v194
	v196 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v190)+16)) = v196
	v198 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v190))) = v198
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v201 = F_replace_nestloop_params_mutator(m, v200, l1)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L12
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v190)+4)) = v201
	return v190
L57:
	;
	return v313
L58:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v262 = F_exprType(m, v261)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L12
	} else {
		goto L72
	}
L59:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v206)+4))
	if v209 <= int32(0) {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v214 = v205
	goto L61
L61:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v206)+12))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v219+v214<<(uint(int32(2))%32))))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+8))
	v225 = F_equal(m, l0, v224)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L12
	} else {
		goto L63
	}
L62:
	;
	v234 = F_palloc0(m, int32(28))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L12
	} else {
		goto L68
	}
L63:
	;
	if v225 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v230 = v214 + int32(1)
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v206)+4))
	if v230 < v231 {
		v214 = v230
		goto L61
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	goto L62
L67:
	;
	goto L58
L68:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v234))) = int64(4294967304)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v234)+8)) = v238
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v241 = F_exprType(m, v240)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L12
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234)+12)) = v241
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v245 = F_exprTypmod(m, v244)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L12
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234)+16)) = v245
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v249 = F_exprCollation(m, v248)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L12
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234)+24)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v234)+20)) = v249
	v313 = v234
	goto L57
L72:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v265 = F_exprTypmod(m, v264)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L12
	} else {
		goto L73
	}
L73:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v268 = F_exprCollation(m, v267)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L12
	} else {
		goto L74
	}
L74:
	;
	v271 = F_palloc0(m, int32(28))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L12
	} else {
		goto L75
	}
L75:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v271))) = int64(4294967304)
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v275)+64))
	if v276 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v276)+4))
	v279 = v277
	goto L78
L77:
	;
	v279 = int32(0)
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v271)+8)) = v279
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)+64))
	v283 = F_lappend_oid(m, v282, v262)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L12
	} else {
		goto L79
	}
L79:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v285)+64)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v271)+24)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v271)+20)) = v268
	*(*int32)(unsafe.Add(mBase, uint32(v271)+16)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v271)+12)) = v262
	v293 = F_palloc0(m, int32(12))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L12
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v293))) = int32(357)
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v271)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v293)+4)) = v297
	v299 = F_copyObjectImpl(m, l0)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L12
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v293)+8)) = v299
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l1)+356))
	v303 = F_lappend(m, v302, v293)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L12
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+356)) = v303
	v313 = v271
	goto L57
L83:
	;
	return v316
}
func F_report_invalid_encoding_db(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_report_invalid_encoding_db[0]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	F_report_invalid_encoding_int(m, v6, l0, l1, l2)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		base.Wasm_trap_unreachable()
		for {
		}
	}
}
func F_reset_formatted_start_time(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_reset_formatted_start_time[0])) = uint8(v2)
	return
}
func F_resolve_anyrange_from_others(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v9 != 0 {
		v10 = F_getBaseType(m, v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			v12 = F_get_multirange_range(m, v10)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				if v12 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						F_errcode(m, int32(67141764))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return
						} else {
							v40 = F_format_type_be(m, v10)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v40
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_resolve_anyrange_from_others_0)
								F_errmsg(m, int32(_a_F_resolve_anyrange_from_others_1), v7)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_resolve_anyrange_from_others_2), int32(699), int32(_a_F_resolve_anyrange_from_others_3))
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
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
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v12
					m.G0 = v7 + int32(16)
					return
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_resolve_anyrange_from_others_4), int32(0))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_resolve_anyrange_from_others_2), int32(703), int32(_a_F_resolve_anyrange_from_others_3))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
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
func F_rowtype_field_matches(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	if l0 == int32(2249) {
		v55 = int32(1)
		return v55
	} else {
		v11 = int32(0)
		v13 = F_lookup_rowtype_tupdesc_domain(m, l0, int32(-1))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			if int32(0) < l1 {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				if l1 <= v19 {
					v30 = v13 + v19<<(uint(int32(4))%32) + l1*int32(100)
					v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+11)))
					if v31 != 0 {
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
						if int32(0) <= v41 {
							v49 = v11
							F_DecrTupleDescRefCount(m, v13)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								v55 = v49
								return v55
							}
						} else {
							v55 = v11
							return v55
						}
					} else {
						v33 = v30 - int32(80)
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+68))
						if v34 != l2 {
							v41 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
							if int32(0) <= v41 {
								v49 = v11
								F_DecrTupleDescRefCount(m, v13)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									v55 = v49
									return v55
								}
							} else {
								v55 = v11
								return v55
							}
						} else {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+76))
							if v36 != l3 {
								v41 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
								if int32(0) <= v41 {
									v49 = v11
									F_DecrTupleDescRefCount(m, v13)
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return int32(0)
									} else {
										v55 = v49
										return v55
									}
								} else {
									v55 = v11
									return v55
								}
							} else {
								v38 = *(*int32)(unsafe.Add(mBase, uint32(v33)+96))
								if v38 == l4 {
									v44 = int32(1)
									v45 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
									if v45 < int32(0) {
										v55 = v44
										return v55
									} else {
										v49 = v44
										F_DecrTupleDescRefCount(m, v13)
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return int32(0)
										} else {
											v55 = v49
											return v55
										}
									}
								} else {
									v41 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
									if int32(0) <= v41 {
										v49 = v11
										F_DecrTupleDescRefCount(m, v13)
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return int32(0)
										} else {
											v55 = v49
											return v55
										}
									} else {
										v55 = v11
										return v55
									}
								}
							}
						}
					}
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
					if int32(0) <= v22 {
						v49 = v11
						F_DecrTupleDescRefCount(m, v13)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							v55 = v49
							return v55
						}
					} else {
						v55 = v11
						return v55
					}
				}
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
				if int32(0) <= v22 {
					v49 = v11
					F_DecrTupleDescRefCount(m, v13)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						v55 = v49
						return v55
					}
				} else {
					v55 = v11
					return v55
				}
			}
		}
	}
}
func F_rt__int_size(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = F_ArrayGetNItemsSafe(m, v3, l0+int32(16))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		*(*float32)(unsafe.Add(mBase, uint32(l1))) = base.F32_convert_i32_s(v6)
		return
	}
}
func F_rtrim1(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13856(m, l0, int32(1), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
