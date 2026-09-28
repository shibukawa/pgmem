package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_ReplicationOriginShmemAttach(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationOriginShmemAttach[0]))
	if v3 != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationOriginShmemAttach[1]))
		*(*int32)(unsafe.Add(mBase, _c_F_ReplicationOriginShmemAttach[2])) = v6 + int32(8)
	} else {
	}
	return
}
func F_ReplicationSlotDropAcquired(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v3 = int32(_a_F_ReplicationSlotDropAcquired_0)
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropAcquired[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropAcquired[0])) = int32(0)
	F_ReplicationSlotDropPtr(m, v4)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		if l0 != 0 {
			F_RequestDisableLogicalDecoding(m)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				return
			}
		} else {
			return
		}
	}
}
func F_ReplicationSlotDropPtr(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v166 int64
	_ = v166
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	v5 = m.G0
	v7 = v5 - int32(2112)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[0]))
	v14 = F_LWLockAcquire(m, v10+int32(_a_F_ReplicationSlotDropPtr_0), int32(0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v17 = l0 + int32(24)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+52)) = v17
		*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = int32(_a_F_ReplicationSlotDropPtr_1)
		v22 = v7 + int32(1088)
		v26 = F_pg_sprintf(m, v22, int32(_a_F_ReplicationSlotDropPtr_2), v7+int32(48))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v17
			*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = int32(_a_F_ReplicationSlotDropPtr_1)
			v32 = v7 - int32(-64)
			v36 = F_pg_sprintf(m, v32, int32(_a_F_ReplicationSlotDropPtr_3), v7+int32(32))
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return
			} else {
				v38 = F_rename(m, v22, v32)
				mBase = m.M
				if v38 == int32(0) {
					v41 = int32(_a_F_ReplicationSlotDropPtr_4)
					v43 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[1]))
					v44 = int32(1)
					*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[1])) = v43 + v44
					F_fsync_fname(m, v32, v44)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						F_fsync_fname(m, int32(_a_F_ReplicationSlotDropPtr_1), int32(1))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return
						} else {
							v54 = int32(_a_F_ReplicationSlotDropPtr_4)
							v56 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[1]))
							*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[1])) = v56 - int32(1)
							v104 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[0]))
							v108 = F_LWLockAcquire(m, v104+int32(_a_F_ReplicationSlotDropPtr_5), int32(0))
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return
							} else {
								v110 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v110)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(-1)
								v115 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[0]))
								F_LWLockRelease(m, v115+int32(_a_F_ReplicationSlotDropPtr_5))
								mBase = m.M
								v119 = m.ExcPending
								if v119 != 0 {
									return
								} else {
									F_ConditionVariableBroadcast(m, l0+int32(224))
									mBase = m.M
									v123 = m.ExcPending
									if v123 != 0 {
										return
									} else {
										F_ReplicationSlotsComputeRequiredXmin(m, int32(0))
										mBase = m.M
										v126 = m.ExcPending
										if v126 != 0 {
											return
										} else {
											F_ReplicationSlotsComputeRequiredLSN(m)
											mBase = m.M
											v128 = m.ExcPending
											if v128 != 0 {
												return
											} else {
												v130 = v7 - int32(-64)
												v131 = F_rmtree(m, v130)
												mBase = m.M
												v132 = m.ExcPending
												if v132 != 0 {
													return
												} else {
													if v131 != 0 {
														v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
														if v148 != 0 {
															v152 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[2]))
															v155 = base.I32_div_s(l0-v152, int32(296))
															v158 = F_pgstat_drop_entry(m, int32(4), int32(0), base.I64_extend_i32_s(v155), int32(0))
															mBase = m.M
															v159 = m.ExcPending
															if v159 != 0 {
																return
															} else {
																if v158 == int32(0) {
																	v163 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[3]))
																	v166 = base.AtomicRmwAdd64(m, v163, int32(16), int64(1))
																} else {
																}
																v168 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[0]))
																F_LWLockRelease(m, v168+int32(_a_F_ReplicationSlotDropPtr_0))
																mBase = m.M
																v172 = m.ExcPending
																if v172 != 0 {
																	return
																} else {
																	m.G0 = v7 + int32(2112)
																	return
																}
															}
														} else {
															v168 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[0]))
															F_LWLockRelease(m, v168+int32(_a_F_ReplicationSlotDropPtr_0))
															mBase = m.M
															v172 = m.ExcPending
															if v172 != 0 {
																return
															} else {
																m.G0 = v7 + int32(2112)
																return
															}
														}
													} else {
														v135 = F_errstart(m, int32(19), int32(0))
														mBase = m.M
														v136 = m.ExcPending
														if v136 != 0 {
															return
														} else {
															if v135 == int32(0) {
																v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
																if v148 != 0 {
																	v152 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[2]))
																	v155 = base.I32_div_s(l0-v152, int32(296))
																	v158 = F_pgstat_drop_entry(m, int32(4), int32(0), base.I64_extend_i32_s(v155), int32(0))
																	mBase = m.M
																	v159 = m.ExcPending
																	if v159 != 0 {
																		return
																	} else {
																		if v158 == int32(0) {
																			v163 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[3]))
																			v166 = base.AtomicRmwAdd64(m, v163, int32(16), int64(1))
																		} else {
																		}
																		v168 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[0]))
																		F_LWLockRelease(m, v168+int32(_a_F_ReplicationSlotDropPtr_0))
																		mBase = m.M
																		v172 = m.ExcPending
																		if v172 != 0 {
																			return
																		} else {
																			m.G0 = v7 + int32(2112)
																			return
																		}
																	}
																} else {
																	v168 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[0]))
																	F_LWLockRelease(m, v168+int32(_a_F_ReplicationSlotDropPtr_0))
																	mBase = m.M
																	v172 = m.ExcPending
																	if v172 != 0 {
																		return
																	} else {
																		m.G0 = v7 + int32(2112)
																		return
																	}
																}
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v7))) = v130
																F_errmsg(m, int32(_a_F_ReplicationSlotDropPtr_6), v7)
																mBase = m.M
																v142 = m.ExcPending
																if v142 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(_a_F_ReplicationSlotDropPtr_7), int32(1141), int32(_a_F_ReplicationSlotDropPtr_8))
																	mBase = m.M
																	v147 = m.ExcPending
																	if v147 != 0 {
																		return
																	} else {
																		v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
																		if v148 != 0 {
																			v152 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[2]))
																			v155 = base.I32_div_s(l0-v152, int32(296))
																			v158 = F_pgstat_drop_entry(m, int32(4), int32(0), base.I64_extend_i32_s(v155), int32(0))
																			mBase = m.M
																			v159 = m.ExcPending
																			if v159 != 0 {
																				return
																			} else {
																				if v158 == int32(0) {
																					v163 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[3]))
																					v166 = base.AtomicRmwAdd64(m, v163, int32(16), int64(1))
																				} else {
																				}
																				v168 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[0]))
																				F_LWLockRelease(m, v168+int32(_a_F_ReplicationSlotDropPtr_0))
																				mBase = m.M
																				v172 = m.ExcPending
																				if v172 != 0 {
																					return
																				} else {
																					m.G0 = v7 + int32(2112)
																					return
																				}
																			}
																		} else {
																			v168 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[0]))
																			F_LWLockRelease(m, v168+int32(_a_F_ReplicationSlotDropPtr_0))
																			mBase = m.M
																			v172 = m.ExcPending
																			if v172 != 0 {
																				return
																			} else {
																				m.G0 = v7 + int32(2112)
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
					v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
					v63 = base.AtomicRmwXchg32(m, l0, int32(0), int32(1))
					if v63 != 0 {
						F_s_lock(m, l0, int32(_a_F_ReplicationSlotDropPtr_9))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(-1)
							v69 = int32(0)
							atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v69))
							F_ConditionVariableBroadcast(m, l0+int32(224))
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return
							} else {
								if v60 != 0 {
									v78 = int32(19)
								} else {
									v78 = int32(21)
								}
								v80 = F_errstart(m, v78, int32(0))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return
								} else {
									if v80 == int32(0) {
										v104 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[0]))
										v108 = F_LWLockAcquire(m, v104+int32(_a_F_ReplicationSlotDropPtr_5), int32(0))
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
											return
										} else {
											v110 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v110)
											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(-1)
											v115 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[0]))
											F_LWLockRelease(m, v115+int32(_a_F_ReplicationSlotDropPtr_5))
											mBase = m.M
											v119 = m.ExcPending
											if v119 != 0 {
												return
											} else {
												F_ConditionVariableBroadcast(m, l0+int32(224))
												mBase = m.M
												v123 = m.ExcPending
												if v123 != 0 {
													return
												} else {
													F_ReplicationSlotsComputeRequiredXmin(m, int32(0))
													mBase = m.M
													v126 = m.ExcPending
													if v126 != 0 {
														return
													} else {
														F_ReplicationSlotsComputeRequiredLSN(m)
														mBase = m.M
														v128 = m.ExcPending
														if v128 != 0 {
															return
														} else {
															v130 = v7 - int32(-64)
															v131 = F_rmtree(m, v130)
															mBase = m.M
															v132 = m.ExcPending
															if v132 != 0 {
																return
															} else {
																if v131 != 0 {
																	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
																	if v148 != 0 {
																		v152 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[2]))
																		v155 = base.I32_div_s(l0-v152, int32(296))
																		v158 = F_pgstat_drop_entry(m, int32(4), int32(0), base.I64_extend_i32_s(v155), int32(0))
																		mBase = m.M
																		v159 = m.ExcPending
																		if v159 != 0 {
																			return
																		} else {
																			if v158 == int32(0) {
																				v163 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[3]))
																				v166 = base.AtomicRmwAdd64(m, v163, int32(16), int64(1))
																			} else {
																			}
																			v168 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[0]))
																			F_LWLockRelease(m, v168+int32(_a_F_ReplicationSlotDropPtr_0))
																			mBase = m.M
																			v172 = m.ExcPending
																			if v172 != 0 {
																				return
																			} else {
																				m.G0 = v7 + int32(2112)
																				return
																			}
																		}
																	} else {
																		v168 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[0]))
																		F_LWLockRelease(m, v168+int32(_a_F_ReplicationSlotDropPtr_0))
																		mBase = m.M
																		v172 = m.ExcPending
																		if v172 != 0 {
																			return
																		} else {
																			m.G0 = v7 + int32(2112)
																			return
																		}
																	}
																} else {
																	v135 = F_errstart(m, int32(19), int32(0))
																	mBase = m.M
																	v136 = m.ExcPending
																	if v136 != 0 {
																		return
																	} else {
																		if v135 == int32(0) {
																			v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
																			if v148 != 0 {
																				v152 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[2]))
																				v155 = base.I32_div_s(l0-v152, int32(296))
																				v158 = F_pgstat_drop_entry(m, int32(4), int32(0), base.I64_extend_i32_s(v155), int32(0))
																				mBase = m.M
																				v159 = m.ExcPending
																				if v159 != 0 {
																					return
																				} else {
																					if v158 == int32(0) {
																						v163 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[3]))
																						v166 = base.AtomicRmwAdd64(m, v163, int32(16), int64(1))
																					} else {
																					}
																					v168 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[0]))
																					F_LWLockRelease(m, v168+int32(_a_F_ReplicationSlotDropPtr_0))
																					mBase = m.M
																					v172 = m.ExcPending
																					if v172 != 0 {
																						return
																					} else {
																						m.G0 = v7 + int32(2112)
																						return
																					}
																				}
																			} else {
																				v168 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[0]))
																				F_LWLockRelease(m, v168+int32(_a_F_ReplicationSlotDropPtr_0))
																				mBase = m.M
																				v172 = m.ExcPending
																				if v172 != 0 {
																					return
																				} else {
																					m.G0 = v7 + int32(2112)
																					return
																				}
																			}
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v7))) = v130
																			F_errmsg(m, int32(_a_F_ReplicationSlotDropPtr_6), v7)
																			mBase = m.M
																			v142 = m.ExcPending
																			if v142 != 0 {
																				return
																			} else {
																				F_errfinish(m, int32(_a_F_ReplicationSlotDropPtr_7), int32(1141), int32(_a_F_ReplicationSlotDropPtr_8))
																				mBase = m.M
																				v147 = m.ExcPending
																				if v147 != 0 {
																					return
																				} else {
																					v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
																					if v148 != 0 {
																						v152 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[2]))
																						v155 = base.I32_div_s(l0-v152, int32(296))
																						v158 = F_pgstat_drop_entry(m, int32(4), int32(0), base.I64_extend_i32_s(v155), int32(0))
																						mBase = m.M
																						v159 = m.ExcPending
																						if v159 != 0 {
																							return
																						} else {
																							if v158 == int32(0) {
																								v163 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[3]))
																								v166 = base.AtomicRmwAdd64(m, v163, int32(16), int64(1))
																							} else {
																							}
																							v168 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[0]))
																							F_LWLockRelease(m, v168+int32(_a_F_ReplicationSlotDropPtr_0))
																							mBase = m.M
																							v172 = m.ExcPending
																							if v172 != 0 {
																								return
																							} else {
																								m.G0 = v7 + int32(2112)
																								return
																							}
																						}
																					} else {
																						v168 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[0]))
																						F_LWLockRelease(m, v168+int32(_a_F_ReplicationSlotDropPtr_0))
																						mBase = m.M
																						v172 = m.ExcPending
																						if v172 != 0 {
																							return
																						} else {
																							m.G0 = v7 + int32(2112)
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
									} else {
										F_errcode_for_file_access(m)
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v7 - int32(-64)
											*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v7 + int32(1088)
											F_errmsg(m, int32(_a_F_ReplicationSlotDropPtr_10), v7+int32(16))
											mBase = m.M
											v96 = m.ExcPending
											if v96 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_ReplicationSlotDropPtr_7), int32(1109), int32(_a_F_ReplicationSlotDropPtr_8))
												mBase = m.M
												v101 = m.ExcPending
												if v101 != 0 {
													return
												} else {
													v104 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[0]))
													v108 = F_LWLockAcquire(m, v104+int32(_a_F_ReplicationSlotDropPtr_5), int32(0))
													mBase = m.M
													v109 = m.ExcPending
													if v109 != 0 {
														return
													} else {
														v110 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v110)
														*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(-1)
														v115 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[0]))
														F_LWLockRelease(m, v115+int32(_a_F_ReplicationSlotDropPtr_5))
														mBase = m.M
														v119 = m.ExcPending
														if v119 != 0 {
															return
														} else {
															F_ConditionVariableBroadcast(m, l0+int32(224))
															mBase = m.M
															v123 = m.ExcPending
															if v123 != 0 {
																return
															} else {
																F_ReplicationSlotsComputeRequiredXmin(m, int32(0))
																mBase = m.M
																v126 = m.ExcPending
																if v126 != 0 {
																	return
																} else {
																	F_ReplicationSlotsComputeRequiredLSN(m)
																	mBase = m.M
																	v128 = m.ExcPending
																	if v128 != 0 {
																		return
																	} else {
																		v130 = v7 - int32(-64)
																		v131 = F_rmtree(m, v130)
																		mBase = m.M
																		v132 = m.ExcPending
																		if v132 != 0 {
																			return
																		} else {
																			if v131 != 0 {
																				v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
																				if v148 != 0 {
																					v152 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[2]))
																					v155 = base.I32_div_s(l0-v152, int32(296))
																					v158 = F_pgstat_drop_entry(m, int32(4), int32(0), base.I64_extend_i32_s(v155), int32(0))
																					mBase = m.M
																					v159 = m.ExcPending
																					if v159 != 0 {
																						return
																					} else {
																						if v158 == int32(0) {
																							v163 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[3]))
																							v166 = base.AtomicRmwAdd64(m, v163, int32(16), int64(1))
																						} else {
																						}
																						v168 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[0]))
																						F_LWLockRelease(m, v168+int32(_a_F_ReplicationSlotDropPtr_0))
																						mBase = m.M
																						v172 = m.ExcPending
																						if v172 != 0 {
																							return
																						} else {
																							m.G0 = v7 + int32(2112)
																							return
																						}
																					}
																				} else {
																					v168 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[0]))
																					F_LWLockRelease(m, v168+int32(_a_F_ReplicationSlotDropPtr_0))
																					mBase = m.M
																					v172 = m.ExcPending
																					if v172 != 0 {
																						return
																					} else {
																						m.G0 = v7 + int32(2112)
																						return
																					}
																				}
																			} else {
																				v135 = F_errstart(m, int32(19), int32(0))
																				mBase = m.M
																				v136 = m.ExcPending
																				if v136 != 0 {
																					return
																				} else {
																					if v135 == int32(0) {
																						v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
																						if v148 != 0 {
																							v152 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[2]))
																							v155 = base.I32_div_s(l0-v152, int32(296))
																							v158 = F_pgstat_drop_entry(m, int32(4), int32(0), base.I64_extend_i32_s(v155), int32(0))
																							mBase = m.M
																							v159 = m.ExcPending
																							if v159 != 0 {
																								return
																							} else {
																								if v158 == int32(0) {
																									v163 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[3]))
																									v166 = base.AtomicRmwAdd64(m, v163, int32(16), int64(1))
																								} else {
																								}
																								v168 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[0]))
																								F_LWLockRelease(m, v168+int32(_a_F_ReplicationSlotDropPtr_0))
																								mBase = m.M
																								v172 = m.ExcPending
																								if v172 != 0 {
																									return
																								} else {
																									m.G0 = v7 + int32(2112)
																									return
																								}
																							}
																						} else {
																							v168 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[0]))
																							F_LWLockRelease(m, v168+int32(_a_F_ReplicationSlotDropPtr_0))
																							mBase = m.M
																							v172 = m.ExcPending
																							if v172 != 0 {
																								return
																							} else {
																								m.G0 = v7 + int32(2112)
																								return
																							}
																						}
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v130
																						F_errmsg(m, int32(_a_F_ReplicationSlotDropPtr_6), v7)
																						mBase = m.M
																						v142 = m.ExcPending
																						if v142 != 0 {
																							return
																						} else {
																							F_errfinish(m, int32(_a_F_ReplicationSlotDropPtr_7), int32(1141), int32(_a_F_ReplicationSlotDropPtr_8))
																							mBase = m.M
																							v147 = m.ExcPending
																							if v147 != 0 {
																								return
																							} else {
																								v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
																								if v148 != 0 {
																									v152 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[2]))
																									v155 = base.I32_div_s(l0-v152, int32(296))
																									v158 = F_pgstat_drop_entry(m, int32(4), int32(0), base.I64_extend_i32_s(v155), int32(0))
																									mBase = m.M
																									v159 = m.ExcPending
																									if v159 != 0 {
																										return
																									} else {
																										if v158 == int32(0) {
																											v163 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[3]))
																											v166 = base.AtomicRmwAdd64(m, v163, int32(16), int64(1))
																										} else {
																										}
																										v168 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[0]))
																										F_LWLockRelease(m, v168+int32(_a_F_ReplicationSlotDropPtr_0))
																										mBase = m.M
																										v172 = m.ExcPending
																										if v172 != 0 {
																											return
																										} else {
																											m.G0 = v7 + int32(2112)
																											return
																										}
																									}
																								} else {
																									v168 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[0]))
																									F_LWLockRelease(m, v168+int32(_a_F_ReplicationSlotDropPtr_0))
																									mBase = m.M
																									v172 = m.ExcPending
																									if v172 != 0 {
																										return
																									} else {
																										m.G0 = v7 + int32(2112)
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
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(-1)
						v69 = int32(0)
						atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v69))
						F_ConditionVariableBroadcast(m, l0+int32(224))
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return
						} else {
							if v60 != 0 {
								v78 = int32(19)
							} else {
								v78 = int32(21)
							}
							v80 = F_errstart(m, v78, int32(0))
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return
							} else {
								if v80 == int32(0) {
									v104 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[0]))
									v108 = F_LWLockAcquire(m, v104+int32(_a_F_ReplicationSlotDropPtr_5), int32(0))
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
										return
									} else {
										v110 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v110)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(-1)
										v115 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[0]))
										F_LWLockRelease(m, v115+int32(_a_F_ReplicationSlotDropPtr_5))
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return
										} else {
											F_ConditionVariableBroadcast(m, l0+int32(224))
											mBase = m.M
											v123 = m.ExcPending
											if v123 != 0 {
												return
											} else {
												F_ReplicationSlotsComputeRequiredXmin(m, int32(0))
												mBase = m.M
												v126 = m.ExcPending
												if v126 != 0 {
													return
												} else {
													F_ReplicationSlotsComputeRequiredLSN(m)
													mBase = m.M
													v128 = m.ExcPending
													if v128 != 0 {
														return
													} else {
														v130 = v7 - int32(-64)
														v131 = F_rmtree(m, v130)
														mBase = m.M
														v132 = m.ExcPending
														if v132 != 0 {
															return
														} else {
															if v131 != 0 {
																v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
																if v148 != 0 {
																	v152 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[2]))
																	v155 = base.I32_div_s(l0-v152, int32(296))
																	v158 = F_pgstat_drop_entry(m, int32(4), int32(0), base.I64_extend_i32_s(v155), int32(0))
																	mBase = m.M
																	v159 = m.ExcPending
																	if v159 != 0 {
																		return
																	} else {
																		if v158 == int32(0) {
																			v163 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[3]))
																			v166 = base.AtomicRmwAdd64(m, v163, int32(16), int64(1))
																		} else {
																		}
																		v168 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[0]))
																		F_LWLockRelease(m, v168+int32(_a_F_ReplicationSlotDropPtr_0))
																		mBase = m.M
																		v172 = m.ExcPending
																		if v172 != 0 {
																			return
																		} else {
																			m.G0 = v7 + int32(2112)
																			return
																		}
																	}
																} else {
																	v168 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[0]))
																	F_LWLockRelease(m, v168+int32(_a_F_ReplicationSlotDropPtr_0))
																	mBase = m.M
																	v172 = m.ExcPending
																	if v172 != 0 {
																		return
																	} else {
																		m.G0 = v7 + int32(2112)
																		return
																	}
																}
															} else {
																v135 = F_errstart(m, int32(19), int32(0))
																mBase = m.M
																v136 = m.ExcPending
																if v136 != 0 {
																	return
																} else {
																	if v135 == int32(0) {
																		v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
																		if v148 != 0 {
																			v152 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[2]))
																			v155 = base.I32_div_s(l0-v152, int32(296))
																			v158 = F_pgstat_drop_entry(m, int32(4), int32(0), base.I64_extend_i32_s(v155), int32(0))
																			mBase = m.M
																			v159 = m.ExcPending
																			if v159 != 0 {
																				return
																			} else {
																				if v158 == int32(0) {
																					v163 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[3]))
																					v166 = base.AtomicRmwAdd64(m, v163, int32(16), int64(1))
																				} else {
																				}
																				v168 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[0]))
																				F_LWLockRelease(m, v168+int32(_a_F_ReplicationSlotDropPtr_0))
																				mBase = m.M
																				v172 = m.ExcPending
																				if v172 != 0 {
																					return
																				} else {
																					m.G0 = v7 + int32(2112)
																					return
																				}
																			}
																		} else {
																			v168 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[0]))
																			F_LWLockRelease(m, v168+int32(_a_F_ReplicationSlotDropPtr_0))
																			mBase = m.M
																			v172 = m.ExcPending
																			if v172 != 0 {
																				return
																			} else {
																				m.G0 = v7 + int32(2112)
																				return
																			}
																		}
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v7))) = v130
																		F_errmsg(m, int32(_a_F_ReplicationSlotDropPtr_6), v7)
																		mBase = m.M
																		v142 = m.ExcPending
																		if v142 != 0 {
																			return
																		} else {
																			F_errfinish(m, int32(_a_F_ReplicationSlotDropPtr_7), int32(1141), int32(_a_F_ReplicationSlotDropPtr_8))
																			mBase = m.M
																			v147 = m.ExcPending
																			if v147 != 0 {
																				return
																			} else {
																				v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
																				if v148 != 0 {
																					v152 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[2]))
																					v155 = base.I32_div_s(l0-v152, int32(296))
																					v158 = F_pgstat_drop_entry(m, int32(4), int32(0), base.I64_extend_i32_s(v155), int32(0))
																					mBase = m.M
																					v159 = m.ExcPending
																					if v159 != 0 {
																						return
																					} else {
																						if v158 == int32(0) {
																							v163 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[3]))
																							v166 = base.AtomicRmwAdd64(m, v163, int32(16), int64(1))
																						} else {
																						}
																						v168 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[0]))
																						F_LWLockRelease(m, v168+int32(_a_F_ReplicationSlotDropPtr_0))
																						mBase = m.M
																						v172 = m.ExcPending
																						if v172 != 0 {
																							return
																						} else {
																							m.G0 = v7 + int32(2112)
																							return
																						}
																					}
																				} else {
																					v168 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[0]))
																					F_LWLockRelease(m, v168+int32(_a_F_ReplicationSlotDropPtr_0))
																					mBase = m.M
																					v172 = m.ExcPending
																					if v172 != 0 {
																						return
																					} else {
																						m.G0 = v7 + int32(2112)
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
								} else {
									F_errcode_for_file_access(m)
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v7 - int32(-64)
										*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v7 + int32(1088)
										F_errmsg(m, int32(_a_F_ReplicationSlotDropPtr_10), v7+int32(16))
										mBase = m.M
										v96 = m.ExcPending
										if v96 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_ReplicationSlotDropPtr_7), int32(1109), int32(_a_F_ReplicationSlotDropPtr_8))
											mBase = m.M
											v101 = m.ExcPending
											if v101 != 0 {
												return
											} else {
												v104 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[0]))
												v108 = F_LWLockAcquire(m, v104+int32(_a_F_ReplicationSlotDropPtr_5), int32(0))
												mBase = m.M
												v109 = m.ExcPending
												if v109 != 0 {
													return
												} else {
													v110 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v110)
													*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(-1)
													v115 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[0]))
													F_LWLockRelease(m, v115+int32(_a_F_ReplicationSlotDropPtr_5))
													mBase = m.M
													v119 = m.ExcPending
													if v119 != 0 {
														return
													} else {
														F_ConditionVariableBroadcast(m, l0+int32(224))
														mBase = m.M
														v123 = m.ExcPending
														if v123 != 0 {
															return
														} else {
															F_ReplicationSlotsComputeRequiredXmin(m, int32(0))
															mBase = m.M
															v126 = m.ExcPending
															if v126 != 0 {
																return
															} else {
																F_ReplicationSlotsComputeRequiredLSN(m)
																mBase = m.M
																v128 = m.ExcPending
																if v128 != 0 {
																	return
																} else {
																	v130 = v7 - int32(-64)
																	v131 = F_rmtree(m, v130)
																	mBase = m.M
																	v132 = m.ExcPending
																	if v132 != 0 {
																		return
																	} else {
																		if v131 != 0 {
																			v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
																			if v148 != 0 {
																				v152 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[2]))
																				v155 = base.I32_div_s(l0-v152, int32(296))
																				v158 = F_pgstat_drop_entry(m, int32(4), int32(0), base.I64_extend_i32_s(v155), int32(0))
																				mBase = m.M
																				v159 = m.ExcPending
																				if v159 != 0 {
																					return
																				} else {
																					if v158 == int32(0) {
																						v163 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[3]))
																						v166 = base.AtomicRmwAdd64(m, v163, int32(16), int64(1))
																					} else {
																					}
																					v168 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[0]))
																					F_LWLockRelease(m, v168+int32(_a_F_ReplicationSlotDropPtr_0))
																					mBase = m.M
																					v172 = m.ExcPending
																					if v172 != 0 {
																						return
																					} else {
																						m.G0 = v7 + int32(2112)
																						return
																					}
																				}
																			} else {
																				v168 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[0]))
																				F_LWLockRelease(m, v168+int32(_a_F_ReplicationSlotDropPtr_0))
																				mBase = m.M
																				v172 = m.ExcPending
																				if v172 != 0 {
																					return
																				} else {
																					m.G0 = v7 + int32(2112)
																					return
																				}
																			}
																		} else {
																			v135 = F_errstart(m, int32(19), int32(0))
																			mBase = m.M
																			v136 = m.ExcPending
																			if v136 != 0 {
																				return
																			} else {
																				if v135 == int32(0) {
																					v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
																					if v148 != 0 {
																						v152 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[2]))
																						v155 = base.I32_div_s(l0-v152, int32(296))
																						v158 = F_pgstat_drop_entry(m, int32(4), int32(0), base.I64_extend_i32_s(v155), int32(0))
																						mBase = m.M
																						v159 = m.ExcPending
																						if v159 != 0 {
																							return
																						} else {
																							if v158 == int32(0) {
																								v163 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[3]))
																								v166 = base.AtomicRmwAdd64(m, v163, int32(16), int64(1))
																							} else {
																							}
																							v168 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[0]))
																							F_LWLockRelease(m, v168+int32(_a_F_ReplicationSlotDropPtr_0))
																							mBase = m.M
																							v172 = m.ExcPending
																							if v172 != 0 {
																								return
																							} else {
																								m.G0 = v7 + int32(2112)
																								return
																							}
																						}
																					} else {
																						v168 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[0]))
																						F_LWLockRelease(m, v168+int32(_a_F_ReplicationSlotDropPtr_0))
																						mBase = m.M
																						v172 = m.ExcPending
																						if v172 != 0 {
																							return
																						} else {
																							m.G0 = v7 + int32(2112)
																							return
																						}
																					}
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v130
																					F_errmsg(m, int32(_a_F_ReplicationSlotDropPtr_6), v7)
																					mBase = m.M
																					v142 = m.ExcPending
																					if v142 != 0 {
																						return
																					} else {
																						F_errfinish(m, int32(_a_F_ReplicationSlotDropPtr_7), int32(1141), int32(_a_F_ReplicationSlotDropPtr_8))
																						mBase = m.M
																						v147 = m.ExcPending
																						if v147 != 0 {
																							return
																						} else {
																							v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
																							if v148 != 0 {
																								v152 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[2]))
																								v155 = base.I32_div_s(l0-v152, int32(296))
																								v158 = F_pgstat_drop_entry(m, int32(4), int32(0), base.I64_extend_i32_s(v155), int32(0))
																								mBase = m.M
																								v159 = m.ExcPending
																								if v159 != 0 {
																									return
																								} else {
																									if v158 == int32(0) {
																										v163 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[3]))
																										v166 = base.AtomicRmwAdd64(m, v163, int32(16), int64(1))
																									} else {
																									}
																									v168 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[0]))
																									F_LWLockRelease(m, v168+int32(_a_F_ReplicationSlotDropPtr_0))
																									mBase = m.M
																									v172 = m.ExcPending
																									if v172 != 0 {
																										return
																									} else {
																										m.G0 = v7 + int32(2112)
																										return
																									}
																								}
																							} else {
																								v168 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropPtr[0]))
																								F_LWLockRelease(m, v168+int32(_a_F_ReplicationSlotDropPtr_0))
																								mBase = m.M
																								v172 = m.ExcPending
																								if v172 != 0 {
																									return
																								} else {
																									m.G0 = v7 + int32(2112)
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
func F_ReplicationSlotMarkDirty(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotMarkDirty[0]))
	v6 = base.AtomicRmwXchg32(m, v3, int32(0), int32(1))
	if v6 != 0 {
		F_s_lock(m, v3, int32(_a_F_ReplicationSlotMarkDirty_0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotMarkDirty[0]))
			v12 = int32(257)
			*(*uint16)(unsafe.Add(mBase, uint32(v11)+12)) = uint16(v12)
			v14 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v3))), uint32(v14))
			return
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotMarkDirty[0]))
		v12 = int32(257)
		*(*uint16)(unsafe.Add(mBase, uint32(v11)+12)) = uint16(v12)
		v14 = int32(0)
		atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v3))), uint32(v14))
		return
	}
}
func F_ReplicationSlotValidateName(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	v4 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v4
	v19 = F_ReplicationSlotValidateNameInternal(m, l0, l1, v7+int32(28), v7+int32(24), v7+int32(20))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		if v19 != 0 {
			m.G0 = v7 + int32(32)
			return v19
		} else {
			v24 = F_errstart(m, l2, int32(0))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				if v24 != 0 {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
					F_errcode(m, v26)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v29
						F_errmsg_internal(m, int32(_a_F_ReplicationSlotValidateName_0), v7+int32(16))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
							if v36 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v36
								F_errhint_internal(m, int32(_a_F_ReplicationSlotValidateName_0), v7)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_ReplicationSlotValidateName_1), int32(284), int32(_a_F_ReplicationSlotValidateName_2))
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return int32(0)
									} else {
										v47 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
										F_pfree(m, v47)
										mBase = m.M
										v49 = m.ExcPending
										if v49 != 0 {
											return int32(0)
										} else {
											v50 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
											if v50 == int32(0) {
												m.G0 = v7 + int32(32)
												return v19
											} else {
												F_pfree(m, v50)
												mBase = m.M
												v54 = m.ExcPending
												if v54 != 0 {
													return int32(0)
												} else {
													m.G0 = v7 + int32(32)
													return v19
												}
											}
										}
									}
								}
							} else {
								F_errfinish(m, int32(_a_F_ReplicationSlotValidateName_1), int32(284), int32(_a_F_ReplicationSlotValidateName_2))
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return int32(0)
								} else {
									v47 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
									F_pfree(m, v47)
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return int32(0)
									} else {
										v50 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
										if v50 == int32(0) {
											m.G0 = v7 + int32(32)
											return v19
										} else {
											F_pfree(m, v50)
											mBase = m.M
											v54 = m.ExcPending
											if v54 != 0 {
												return int32(0)
											} else {
												m.G0 = v7 + int32(32)
												return v19
											}
										}
									}
								}
							}
						}
					}
				} else {
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
					F_pfree(m, v47)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
						if v50 == int32(0) {
							m.G0 = v7 + int32(32)
							return v19
						} else {
							F_pfree(m, v50)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								m.G0 = v7 + int32(32)
								return v19
							}
						}
					}
				}
			}
		}
	}
}
func F_replication_yylex_destroy(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int64
	_ = v112
	var v125 int32
	_ = v125
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v5 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v107 != 0 {
		goto L29
	} else {
		goto L30
	}
L2:
	;
	v9 = v5
	goto L3
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v15 = v9 + v12<<(uint(int32(2))%32)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v16 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_pfree(m, v9)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L11
	} else {
		goto L28
	}
L5:
	;
	v17 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	if v19 == v17 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	goto L4
L8:
	;
	F_pfree(m, v16)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L11
	} else {
		goto L13
	}
L9:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v22 == int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	F_pfree(m, v22)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return
L12:
	;
	goto L8
L13:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v35 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v30+v31<<(uint(int32(2))%32)))) = v35
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v38 == v35 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v98 != 0 {
		v9 = v98
		goto L3
	} else {
		goto L27
	}
L15:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v44 = v38 + v41<<(uint(int32(2))%32)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	if v45 == int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v48 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v48
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v45)+20))
	if v50 == v48 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	F_pfree(m, v45)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L11
	} else {
		goto L21
	}
L18:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v53 == int32(0) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	F_pfree(m, v53)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L11
	} else {
		goto L20
	}
L20:
	;
	goto L17
L21:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v61+v62<<(uint(int32(2))%32)))) = int32(0)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v68 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v70 = v68 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v70
	v72 = v70
	goto L24
L23:
	;
	v72 = v35
	goto L24
L24:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v73 == int32(0) {
		goto L14
	} else {
		goto L25
	}
L25:
	;
	v78 = v73 + v72<<(uint(int32(2))%32)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	if v79 == int32(0) {
		goto L14
	} else {
		goto L26
	}
L26:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v82
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v85
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v89
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)) = uint8(v91)
	goto L14
L27:
	;
	goto L1
L28:
	;
	goto L1
L29:
	;
	F_pfree(m, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L11
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v110 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v110
	v112 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = v112
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v112
	*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v110
	*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v110
	F_pfree(m, l0)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L11
	} else {
		goto L33
	}
L32:
	;
	goto L31
L33:
	;
	return
}
