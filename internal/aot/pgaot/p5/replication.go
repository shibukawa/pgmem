package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int64
	_ = v174
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	v5 = m.G0
	v7 = v5 - int32(2112)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v14 = F_LWLockAcquire(m, v10+int32(4608), int32(0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v17 = l0 + int32(24)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+52)) = v17
		*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = int32(83530)
		v26 = F_pg_sprintf(m, v7+int32(1088), int32(174083), v7+int32(48))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v17
			*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = int32(83530)
			v36 = F_pg_sprintf(m, v7-int32(-64), int32(231633), v7+int32(32))
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return
			} else {
				v42 = F_rename(m, v7+int32(1088), v7-int32(-64))
				mBase = m.M
				if v42 == int32(0) {
					v45 = int32(4465060)
					v47 = *(*int32)(unsafe.Add(mBase, _consts[13]))
					v48 = int32(1)
					*(*int32)(unsafe.Add(mBase, _consts[13])) = v47 + v48
					F_fsync_fname(m, v7-int32(-64), v48)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return
					} else {
						F_fsync_fname(m, int32(83530), int32(1))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return
						} else {
							v60 = int32(4465060)
							v62 = *(*int32)(unsafe.Add(mBase, _consts[13]))
							*(*int32)(unsafe.Add(mBase, _consts[13])) = v62 - int32(1)
							v112 = *(*int32)(unsafe.Add(mBase, _consts[7]))
							v116 = F_LWLockAcquire(m, v112+int32(4736), int32(0))
							mBase = m.M
							v117 = m.ExcPending
							if v117 != 0 {
								return
							} else {
								v118 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v118)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v118
								v123 = *(*int32)(unsafe.Add(mBase, _consts[7]))
								F_LWLockRelease(m, v123+int32(4736))
								mBase = m.M
								v127 = m.ExcPending
								if v127 != 0 {
									return
								} else {
									F_ConditionVariableBroadcast(m, l0+int32(224))
									mBase = m.M
									v131 = m.ExcPending
									if v131 != 0 {
										return
									} else {
										F_ReplicationSlotsComputeRequiredXmin(m, int32(0))
										mBase = m.M
										v134 = m.ExcPending
										if v134 != 0 {
											return
										} else {
											F_ReplicationSlotsComputeRequiredLSN(m)
											mBase = m.M
											v136 = m.ExcPending
											if v136 != 0 {
												return
											} else {
												v139 = F_rmtree(m, v7-int32(-64))
												mBase = m.M
												v140 = m.ExcPending
												if v140 != 0 {
													return
												} else {
													if v139 != 0 {
														v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
														if v158 != 0 {
															v162 = *(*int32)(unsafe.Add(mBase, _consts[562]))
															v165 = base.I32_div_s(l0-v162, int32(288))
															v167 = F_pgstat_drop_entry(m, int32(4), int32(0), base.I64_extend_i32_s(v165))
															mBase = m.M
															v168 = m.ExcPending
															if v168 != 0 {
																return
															} else {
																if v167 == int32(0) {
																	v173 = *(*int32)(unsafe.Add(mBase, _consts[563]))
																	v174 = *(*int64)(unsafe.Add(mBase, uint32(v173)+16))
																	*(*int64)(unsafe.Add(mBase, uint32(v173)+16)) = v174 + int64(1)
																} else {
																}
																v179 = *(*int32)(unsafe.Add(mBase, _consts[7]))
																F_LWLockRelease(m, v179+int32(4608))
																mBase = m.M
																v183 = m.ExcPending
																if v183 != 0 {
																	return
																} else {
																	m.G0 = v7 + int32(2112)
																	return
																}
															}
														} else {
															v179 = *(*int32)(unsafe.Add(mBase, _consts[7]))
															F_LWLockRelease(m, v179+int32(4608))
															mBase = m.M
															v183 = m.ExcPending
															if v183 != 0 {
																return
															} else {
																m.G0 = v7 + int32(2112)
																return
															}
														}
													} else {
														v143 = F_errstart(m, int32(19), int32(0))
														mBase = m.M
														v144 = m.ExcPending
														if v144 != 0 {
															return
														} else {
															if v143 == int32(0) {
																v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
																if v158 != 0 {
																	v162 = *(*int32)(unsafe.Add(mBase, _consts[562]))
																	v165 = base.I32_div_s(l0-v162, int32(288))
																	v167 = F_pgstat_drop_entry(m, int32(4), int32(0), base.I64_extend_i32_s(v165))
																	mBase = m.M
																	v168 = m.ExcPending
																	if v168 != 0 {
																		return
																	} else {
																		if v167 == int32(0) {
																			v173 = *(*int32)(unsafe.Add(mBase, _consts[563]))
																			v174 = *(*int64)(unsafe.Add(mBase, uint32(v173)+16))
																			*(*int64)(unsafe.Add(mBase, uint32(v173)+16)) = v174 + int64(1)
																		} else {
																		}
																		v179 = *(*int32)(unsafe.Add(mBase, _consts[7]))
																		F_LWLockRelease(m, v179+int32(4608))
																		mBase = m.M
																		v183 = m.ExcPending
																		if v183 != 0 {
																			return
																		} else {
																			m.G0 = v7 + int32(2112)
																			return
																		}
																	}
																} else {
																	v179 = *(*int32)(unsafe.Add(mBase, _consts[7]))
																	F_LWLockRelease(m, v179+int32(4608))
																	mBase = m.M
																	v183 = m.ExcPending
																	if v183 != 0 {
																		return
																	} else {
																		m.G0 = v7 + int32(2112)
																		return
																	}
																}
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v7))) = v7 - int32(-64)
																F_errmsg(m, int32(666400), v7)
																mBase = m.M
																v152 = m.ExcPending
																if v152 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(484206), int32(1060), int32(203377))
																	mBase = m.M
																	v157 = m.ExcPending
																	if v157 != 0 {
																		return
																	} else {
																		v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
																		if v158 != 0 {
																			v162 = *(*int32)(unsafe.Add(mBase, _consts[562]))
																			v165 = base.I32_div_s(l0-v162, int32(288))
																			v167 = F_pgstat_drop_entry(m, int32(4), int32(0), base.I64_extend_i32_s(v165))
																			mBase = m.M
																			v168 = m.ExcPending
																			if v168 != 0 {
																				return
																			} else {
																				if v167 == int32(0) {
																					v173 = *(*int32)(unsafe.Add(mBase, _consts[563]))
																					v174 = *(*int64)(unsafe.Add(mBase, uint32(v173)+16))
																					*(*int64)(unsafe.Add(mBase, uint32(v173)+16)) = v174 + int64(1)
																				} else {
																				}
																				v179 = *(*int32)(unsafe.Add(mBase, _consts[7]))
																				F_LWLockRelease(m, v179+int32(4608))
																				mBase = m.M
																				v183 = m.ExcPending
																				if v183 != 0 {
																					return
																				} else {
																					m.G0 = v7 + int32(2112)
																					return
																				}
																			}
																		} else {
																			v179 = *(*int32)(unsafe.Add(mBase, _consts[7]))
																			F_LWLockRelease(m, v179+int32(4608))
																			mBase = m.M
																			v183 = m.ExcPending
																			if v183 != 0 {
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
					v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
					v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
					if v66 != 0 {
						F_s_lock(m, l0, int32(484206), int32(1018), int32(203377))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return
						} else {
							v75 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v75
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v75
							F_ConditionVariableBroadcast(m, l0+int32(224))
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return
							} else {
								if v69 != 0 {
									v85 = int32(19)
								} else {
									v85 = int32(21)
								}
								v87 = F_errstart(m, v85, int32(0))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return
								} else {
									if v87 == int32(0) {
										v112 = *(*int32)(unsafe.Add(mBase, _consts[7]))
										v116 = F_LWLockAcquire(m, v112+int32(4736), int32(0))
										mBase = m.M
										v117 = m.ExcPending
										if v117 != 0 {
											return
										} else {
											v118 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v118)
											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v118
											v123 = *(*int32)(unsafe.Add(mBase, _consts[7]))
											F_LWLockRelease(m, v123+int32(4736))
											mBase = m.M
											v127 = m.ExcPending
											if v127 != 0 {
												return
											} else {
												F_ConditionVariableBroadcast(m, l0+int32(224))
												mBase = m.M
												v131 = m.ExcPending
												if v131 != 0 {
													return
												} else {
													F_ReplicationSlotsComputeRequiredXmin(m, int32(0))
													mBase = m.M
													v134 = m.ExcPending
													if v134 != 0 {
														return
													} else {
														F_ReplicationSlotsComputeRequiredLSN(m)
														mBase = m.M
														v136 = m.ExcPending
														if v136 != 0 {
															return
														} else {
															v139 = F_rmtree(m, v7-int32(-64))
															mBase = m.M
															v140 = m.ExcPending
															if v140 != 0 {
																return
															} else {
																if v139 != 0 {
																	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
																	if v158 != 0 {
																		v162 = *(*int32)(unsafe.Add(mBase, _consts[562]))
																		v165 = base.I32_div_s(l0-v162, int32(288))
																		v167 = F_pgstat_drop_entry(m, int32(4), int32(0), base.I64_extend_i32_s(v165))
																		mBase = m.M
																		v168 = m.ExcPending
																		if v168 != 0 {
																			return
																		} else {
																			if v167 == int32(0) {
																				v173 = *(*int32)(unsafe.Add(mBase, _consts[563]))
																				v174 = *(*int64)(unsafe.Add(mBase, uint32(v173)+16))
																				*(*int64)(unsafe.Add(mBase, uint32(v173)+16)) = v174 + int64(1)
																			} else {
																			}
																			v179 = *(*int32)(unsafe.Add(mBase, _consts[7]))
																			F_LWLockRelease(m, v179+int32(4608))
																			mBase = m.M
																			v183 = m.ExcPending
																			if v183 != 0 {
																				return
																			} else {
																				m.G0 = v7 + int32(2112)
																				return
																			}
																		}
																	} else {
																		v179 = *(*int32)(unsafe.Add(mBase, _consts[7]))
																		F_LWLockRelease(m, v179+int32(4608))
																		mBase = m.M
																		v183 = m.ExcPending
																		if v183 != 0 {
																			return
																		} else {
																			m.G0 = v7 + int32(2112)
																			return
																		}
																	}
																} else {
																	v143 = F_errstart(m, int32(19), int32(0))
																	mBase = m.M
																	v144 = m.ExcPending
																	if v144 != 0 {
																		return
																	} else {
																		if v143 == int32(0) {
																			v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
																			if v158 != 0 {
																				v162 = *(*int32)(unsafe.Add(mBase, _consts[562]))
																				v165 = base.I32_div_s(l0-v162, int32(288))
																				v167 = F_pgstat_drop_entry(m, int32(4), int32(0), base.I64_extend_i32_s(v165))
																				mBase = m.M
																				v168 = m.ExcPending
																				if v168 != 0 {
																					return
																				} else {
																					if v167 == int32(0) {
																						v173 = *(*int32)(unsafe.Add(mBase, _consts[563]))
																						v174 = *(*int64)(unsafe.Add(mBase, uint32(v173)+16))
																						*(*int64)(unsafe.Add(mBase, uint32(v173)+16)) = v174 + int64(1)
																					} else {
																					}
																					v179 = *(*int32)(unsafe.Add(mBase, _consts[7]))
																					F_LWLockRelease(m, v179+int32(4608))
																					mBase = m.M
																					v183 = m.ExcPending
																					if v183 != 0 {
																						return
																					} else {
																						m.G0 = v7 + int32(2112)
																						return
																					}
																				}
																			} else {
																				v179 = *(*int32)(unsafe.Add(mBase, _consts[7]))
																				F_LWLockRelease(m, v179+int32(4608))
																				mBase = m.M
																				v183 = m.ExcPending
																				if v183 != 0 {
																					return
																				} else {
																					m.G0 = v7 + int32(2112)
																					return
																				}
																			}
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v7))) = v7 - int32(-64)
																			F_errmsg(m, int32(666400), v7)
																			mBase = m.M
																			v152 = m.ExcPending
																			if v152 != 0 {
																				return
																			} else {
																				F_errfinish(m, int32(484206), int32(1060), int32(203377))
																				mBase = m.M
																				v157 = m.ExcPending
																				if v157 != 0 {
																					return
																				} else {
																					v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
																					if v158 != 0 {
																						v162 = *(*int32)(unsafe.Add(mBase, _consts[562]))
																						v165 = base.I32_div_s(l0-v162, int32(288))
																						v167 = F_pgstat_drop_entry(m, int32(4), int32(0), base.I64_extend_i32_s(v165))
																						mBase = m.M
																						v168 = m.ExcPending
																						if v168 != 0 {
																							return
																						} else {
																							if v167 == int32(0) {
																								v173 = *(*int32)(unsafe.Add(mBase, _consts[563]))
																								v174 = *(*int64)(unsafe.Add(mBase, uint32(v173)+16))
																								*(*int64)(unsafe.Add(mBase, uint32(v173)+16)) = v174 + int64(1)
																							} else {
																							}
																							v179 = *(*int32)(unsafe.Add(mBase, _consts[7]))
																							F_LWLockRelease(m, v179+int32(4608))
																							mBase = m.M
																							v183 = m.ExcPending
																							if v183 != 0 {
																								return
																							} else {
																								m.G0 = v7 + int32(2112)
																								return
																							}
																						}
																					} else {
																						v179 = *(*int32)(unsafe.Add(mBase, _consts[7]))
																						F_LWLockRelease(m, v179+int32(4608))
																						mBase = m.M
																						v183 = m.ExcPending
																						if v183 != 0 {
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
										v92 = m.ExcPending
										if v92 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v7 - int32(-64)
											*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v7 + int32(1088)
											F_errmsg(m, int32(292340), v7+int32(16))
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
												return
											} else {
												F_errfinish(m, int32(484206), int32(1028), int32(203377))
												mBase = m.M
												v108 = m.ExcPending
												if v108 != 0 {
													return
												} else {
													v112 = *(*int32)(unsafe.Add(mBase, _consts[7]))
													v116 = F_LWLockAcquire(m, v112+int32(4736), int32(0))
													mBase = m.M
													v117 = m.ExcPending
													if v117 != 0 {
														return
													} else {
														v118 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v118)
														*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v118
														v123 = *(*int32)(unsafe.Add(mBase, _consts[7]))
														F_LWLockRelease(m, v123+int32(4736))
														mBase = m.M
														v127 = m.ExcPending
														if v127 != 0 {
															return
														} else {
															F_ConditionVariableBroadcast(m, l0+int32(224))
															mBase = m.M
															v131 = m.ExcPending
															if v131 != 0 {
																return
															} else {
																F_ReplicationSlotsComputeRequiredXmin(m, int32(0))
																mBase = m.M
																v134 = m.ExcPending
																if v134 != 0 {
																	return
																} else {
																	F_ReplicationSlotsComputeRequiredLSN(m)
																	mBase = m.M
																	v136 = m.ExcPending
																	if v136 != 0 {
																		return
																	} else {
																		v139 = F_rmtree(m, v7-int32(-64))
																		mBase = m.M
																		v140 = m.ExcPending
																		if v140 != 0 {
																			return
																		} else {
																			if v139 != 0 {
																				v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
																				if v158 != 0 {
																					v162 = *(*int32)(unsafe.Add(mBase, _consts[562]))
																					v165 = base.I32_div_s(l0-v162, int32(288))
																					v167 = F_pgstat_drop_entry(m, int32(4), int32(0), base.I64_extend_i32_s(v165))
																					mBase = m.M
																					v168 = m.ExcPending
																					if v168 != 0 {
																						return
																					} else {
																						if v167 == int32(0) {
																							v173 = *(*int32)(unsafe.Add(mBase, _consts[563]))
																							v174 = *(*int64)(unsafe.Add(mBase, uint32(v173)+16))
																							*(*int64)(unsafe.Add(mBase, uint32(v173)+16)) = v174 + int64(1)
																						} else {
																						}
																						v179 = *(*int32)(unsafe.Add(mBase, _consts[7]))
																						F_LWLockRelease(m, v179+int32(4608))
																						mBase = m.M
																						v183 = m.ExcPending
																						if v183 != 0 {
																							return
																						} else {
																							m.G0 = v7 + int32(2112)
																							return
																						}
																					}
																				} else {
																					v179 = *(*int32)(unsafe.Add(mBase, _consts[7]))
																					F_LWLockRelease(m, v179+int32(4608))
																					mBase = m.M
																					v183 = m.ExcPending
																					if v183 != 0 {
																						return
																					} else {
																						m.G0 = v7 + int32(2112)
																						return
																					}
																				}
																			} else {
																				v143 = F_errstart(m, int32(19), int32(0))
																				mBase = m.M
																				v144 = m.ExcPending
																				if v144 != 0 {
																					return
																				} else {
																					if v143 == int32(0) {
																						v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
																						if v158 != 0 {
																							v162 = *(*int32)(unsafe.Add(mBase, _consts[562]))
																							v165 = base.I32_div_s(l0-v162, int32(288))
																							v167 = F_pgstat_drop_entry(m, int32(4), int32(0), base.I64_extend_i32_s(v165))
																							mBase = m.M
																							v168 = m.ExcPending
																							if v168 != 0 {
																								return
																							} else {
																								if v167 == int32(0) {
																									v173 = *(*int32)(unsafe.Add(mBase, _consts[563]))
																									v174 = *(*int64)(unsafe.Add(mBase, uint32(v173)+16))
																									*(*int64)(unsafe.Add(mBase, uint32(v173)+16)) = v174 + int64(1)
																								} else {
																								}
																								v179 = *(*int32)(unsafe.Add(mBase, _consts[7]))
																								F_LWLockRelease(m, v179+int32(4608))
																								mBase = m.M
																								v183 = m.ExcPending
																								if v183 != 0 {
																									return
																								} else {
																									m.G0 = v7 + int32(2112)
																									return
																								}
																							}
																						} else {
																							v179 = *(*int32)(unsafe.Add(mBase, _consts[7]))
																							F_LWLockRelease(m, v179+int32(4608))
																							mBase = m.M
																							v183 = m.ExcPending
																							if v183 != 0 {
																								return
																							} else {
																								m.G0 = v7 + int32(2112)
																								return
																							}
																						}
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v7 - int32(-64)
																						F_errmsg(m, int32(666400), v7)
																						mBase = m.M
																						v152 = m.ExcPending
																						if v152 != 0 {
																							return
																						} else {
																							F_errfinish(m, int32(484206), int32(1060), int32(203377))
																							mBase = m.M
																							v157 = m.ExcPending
																							if v157 != 0 {
																								return
																							} else {
																								v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
																								if v158 != 0 {
																									v162 = *(*int32)(unsafe.Add(mBase, _consts[562]))
																									v165 = base.I32_div_s(l0-v162, int32(288))
																									v167 = F_pgstat_drop_entry(m, int32(4), int32(0), base.I64_extend_i32_s(v165))
																									mBase = m.M
																									v168 = m.ExcPending
																									if v168 != 0 {
																										return
																									} else {
																										if v167 == int32(0) {
																											v173 = *(*int32)(unsafe.Add(mBase, _consts[563]))
																											v174 = *(*int64)(unsafe.Add(mBase, uint32(v173)+16))
																											*(*int64)(unsafe.Add(mBase, uint32(v173)+16)) = v174 + int64(1)
																										} else {
																										}
																										v179 = *(*int32)(unsafe.Add(mBase, _consts[7]))
																										F_LWLockRelease(m, v179+int32(4608))
																										mBase = m.M
																										v183 = m.ExcPending
																										if v183 != 0 {
																											return
																										} else {
																											m.G0 = v7 + int32(2112)
																											return
																										}
																									}
																								} else {
																									v179 = *(*int32)(unsafe.Add(mBase, _consts[7]))
																									F_LWLockRelease(m, v179+int32(4608))
																									mBase = m.M
																									v183 = m.ExcPending
																									if v183 != 0 {
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
						v75 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v75
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v75
						F_ConditionVariableBroadcast(m, l0+int32(224))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return
						} else {
							if v69 != 0 {
								v85 = int32(19)
							} else {
								v85 = int32(21)
							}
							v87 = F_errstart(m, v85, int32(0))
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return
							} else {
								if v87 == int32(0) {
									v112 = *(*int32)(unsafe.Add(mBase, _consts[7]))
									v116 = F_LWLockAcquire(m, v112+int32(4736), int32(0))
									mBase = m.M
									v117 = m.ExcPending
									if v117 != 0 {
										return
									} else {
										v118 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v118)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v118
										v123 = *(*int32)(unsafe.Add(mBase, _consts[7]))
										F_LWLockRelease(m, v123+int32(4736))
										mBase = m.M
										v127 = m.ExcPending
										if v127 != 0 {
											return
										} else {
											F_ConditionVariableBroadcast(m, l0+int32(224))
											mBase = m.M
											v131 = m.ExcPending
											if v131 != 0 {
												return
											} else {
												F_ReplicationSlotsComputeRequiredXmin(m, int32(0))
												mBase = m.M
												v134 = m.ExcPending
												if v134 != 0 {
													return
												} else {
													F_ReplicationSlotsComputeRequiredLSN(m)
													mBase = m.M
													v136 = m.ExcPending
													if v136 != 0 {
														return
													} else {
														v139 = F_rmtree(m, v7-int32(-64))
														mBase = m.M
														v140 = m.ExcPending
														if v140 != 0 {
															return
														} else {
															if v139 != 0 {
																v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
																if v158 != 0 {
																	v162 = *(*int32)(unsafe.Add(mBase, _consts[562]))
																	v165 = base.I32_div_s(l0-v162, int32(288))
																	v167 = F_pgstat_drop_entry(m, int32(4), int32(0), base.I64_extend_i32_s(v165))
																	mBase = m.M
																	v168 = m.ExcPending
																	if v168 != 0 {
																		return
																	} else {
																		if v167 == int32(0) {
																			v173 = *(*int32)(unsafe.Add(mBase, _consts[563]))
																			v174 = *(*int64)(unsafe.Add(mBase, uint32(v173)+16))
																			*(*int64)(unsafe.Add(mBase, uint32(v173)+16)) = v174 + int64(1)
																		} else {
																		}
																		v179 = *(*int32)(unsafe.Add(mBase, _consts[7]))
																		F_LWLockRelease(m, v179+int32(4608))
																		mBase = m.M
																		v183 = m.ExcPending
																		if v183 != 0 {
																			return
																		} else {
																			m.G0 = v7 + int32(2112)
																			return
																		}
																	}
																} else {
																	v179 = *(*int32)(unsafe.Add(mBase, _consts[7]))
																	F_LWLockRelease(m, v179+int32(4608))
																	mBase = m.M
																	v183 = m.ExcPending
																	if v183 != 0 {
																		return
																	} else {
																		m.G0 = v7 + int32(2112)
																		return
																	}
																}
															} else {
																v143 = F_errstart(m, int32(19), int32(0))
																mBase = m.M
																v144 = m.ExcPending
																if v144 != 0 {
																	return
																} else {
																	if v143 == int32(0) {
																		v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
																		if v158 != 0 {
																			v162 = *(*int32)(unsafe.Add(mBase, _consts[562]))
																			v165 = base.I32_div_s(l0-v162, int32(288))
																			v167 = F_pgstat_drop_entry(m, int32(4), int32(0), base.I64_extend_i32_s(v165))
																			mBase = m.M
																			v168 = m.ExcPending
																			if v168 != 0 {
																				return
																			} else {
																				if v167 == int32(0) {
																					v173 = *(*int32)(unsafe.Add(mBase, _consts[563]))
																					v174 = *(*int64)(unsafe.Add(mBase, uint32(v173)+16))
																					*(*int64)(unsafe.Add(mBase, uint32(v173)+16)) = v174 + int64(1)
																				} else {
																				}
																				v179 = *(*int32)(unsafe.Add(mBase, _consts[7]))
																				F_LWLockRelease(m, v179+int32(4608))
																				mBase = m.M
																				v183 = m.ExcPending
																				if v183 != 0 {
																					return
																				} else {
																					m.G0 = v7 + int32(2112)
																					return
																				}
																			}
																		} else {
																			v179 = *(*int32)(unsafe.Add(mBase, _consts[7]))
																			F_LWLockRelease(m, v179+int32(4608))
																			mBase = m.M
																			v183 = m.ExcPending
																			if v183 != 0 {
																				return
																			} else {
																				m.G0 = v7 + int32(2112)
																				return
																			}
																		}
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v7))) = v7 - int32(-64)
																		F_errmsg(m, int32(666400), v7)
																		mBase = m.M
																		v152 = m.ExcPending
																		if v152 != 0 {
																			return
																		} else {
																			F_errfinish(m, int32(484206), int32(1060), int32(203377))
																			mBase = m.M
																			v157 = m.ExcPending
																			if v157 != 0 {
																				return
																			} else {
																				v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
																				if v158 != 0 {
																					v162 = *(*int32)(unsafe.Add(mBase, _consts[562]))
																					v165 = base.I32_div_s(l0-v162, int32(288))
																					v167 = F_pgstat_drop_entry(m, int32(4), int32(0), base.I64_extend_i32_s(v165))
																					mBase = m.M
																					v168 = m.ExcPending
																					if v168 != 0 {
																						return
																					} else {
																						if v167 == int32(0) {
																							v173 = *(*int32)(unsafe.Add(mBase, _consts[563]))
																							v174 = *(*int64)(unsafe.Add(mBase, uint32(v173)+16))
																							*(*int64)(unsafe.Add(mBase, uint32(v173)+16)) = v174 + int64(1)
																						} else {
																						}
																						v179 = *(*int32)(unsafe.Add(mBase, _consts[7]))
																						F_LWLockRelease(m, v179+int32(4608))
																						mBase = m.M
																						v183 = m.ExcPending
																						if v183 != 0 {
																							return
																						} else {
																							m.G0 = v7 + int32(2112)
																							return
																						}
																					}
																				} else {
																					v179 = *(*int32)(unsafe.Add(mBase, _consts[7]))
																					F_LWLockRelease(m, v179+int32(4608))
																					mBase = m.M
																					v183 = m.ExcPending
																					if v183 != 0 {
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
									v92 = m.ExcPending
									if v92 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v7 - int32(-64)
										*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v7 + int32(1088)
										F_errmsg(m, int32(292340), v7+int32(16))
										mBase = m.M
										v103 = m.ExcPending
										if v103 != 0 {
											return
										} else {
											F_errfinish(m, int32(484206), int32(1028), int32(203377))
											mBase = m.M
											v108 = m.ExcPending
											if v108 != 0 {
												return
											} else {
												v112 = *(*int32)(unsafe.Add(mBase, _consts[7]))
												v116 = F_LWLockAcquire(m, v112+int32(4736), int32(0))
												mBase = m.M
												v117 = m.ExcPending
												if v117 != 0 {
													return
												} else {
													v118 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v118)
													*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v118
													v123 = *(*int32)(unsafe.Add(mBase, _consts[7]))
													F_LWLockRelease(m, v123+int32(4736))
													mBase = m.M
													v127 = m.ExcPending
													if v127 != 0 {
														return
													} else {
														F_ConditionVariableBroadcast(m, l0+int32(224))
														mBase = m.M
														v131 = m.ExcPending
														if v131 != 0 {
															return
														} else {
															F_ReplicationSlotsComputeRequiredXmin(m, int32(0))
															mBase = m.M
															v134 = m.ExcPending
															if v134 != 0 {
																return
															} else {
																F_ReplicationSlotsComputeRequiredLSN(m)
																mBase = m.M
																v136 = m.ExcPending
																if v136 != 0 {
																	return
																} else {
																	v139 = F_rmtree(m, v7-int32(-64))
																	mBase = m.M
																	v140 = m.ExcPending
																	if v140 != 0 {
																		return
																	} else {
																		if v139 != 0 {
																			v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
																			if v158 != 0 {
																				v162 = *(*int32)(unsafe.Add(mBase, _consts[562]))
																				v165 = base.I32_div_s(l0-v162, int32(288))
																				v167 = F_pgstat_drop_entry(m, int32(4), int32(0), base.I64_extend_i32_s(v165))
																				mBase = m.M
																				v168 = m.ExcPending
																				if v168 != 0 {
																					return
																				} else {
																					if v167 == int32(0) {
																						v173 = *(*int32)(unsafe.Add(mBase, _consts[563]))
																						v174 = *(*int64)(unsafe.Add(mBase, uint32(v173)+16))
																						*(*int64)(unsafe.Add(mBase, uint32(v173)+16)) = v174 + int64(1)
																					} else {
																					}
																					v179 = *(*int32)(unsafe.Add(mBase, _consts[7]))
																					F_LWLockRelease(m, v179+int32(4608))
																					mBase = m.M
																					v183 = m.ExcPending
																					if v183 != 0 {
																						return
																					} else {
																						m.G0 = v7 + int32(2112)
																						return
																					}
																				}
																			} else {
																				v179 = *(*int32)(unsafe.Add(mBase, _consts[7]))
																				F_LWLockRelease(m, v179+int32(4608))
																				mBase = m.M
																				v183 = m.ExcPending
																				if v183 != 0 {
																					return
																				} else {
																					m.G0 = v7 + int32(2112)
																					return
																				}
																			}
																		} else {
																			v143 = F_errstart(m, int32(19), int32(0))
																			mBase = m.M
																			v144 = m.ExcPending
																			if v144 != 0 {
																				return
																			} else {
																				if v143 == int32(0) {
																					v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
																					if v158 != 0 {
																						v162 = *(*int32)(unsafe.Add(mBase, _consts[562]))
																						v165 = base.I32_div_s(l0-v162, int32(288))
																						v167 = F_pgstat_drop_entry(m, int32(4), int32(0), base.I64_extend_i32_s(v165))
																						mBase = m.M
																						v168 = m.ExcPending
																						if v168 != 0 {
																							return
																						} else {
																							if v167 == int32(0) {
																								v173 = *(*int32)(unsafe.Add(mBase, _consts[563]))
																								v174 = *(*int64)(unsafe.Add(mBase, uint32(v173)+16))
																								*(*int64)(unsafe.Add(mBase, uint32(v173)+16)) = v174 + int64(1)
																							} else {
																							}
																							v179 = *(*int32)(unsafe.Add(mBase, _consts[7]))
																							F_LWLockRelease(m, v179+int32(4608))
																							mBase = m.M
																							v183 = m.ExcPending
																							if v183 != 0 {
																								return
																							} else {
																								m.G0 = v7 + int32(2112)
																								return
																							}
																						}
																					} else {
																						v179 = *(*int32)(unsafe.Add(mBase, _consts[7]))
																						F_LWLockRelease(m, v179+int32(4608))
																						mBase = m.M
																						v183 = m.ExcPending
																						if v183 != 0 {
																							return
																						} else {
																							m.G0 = v7 + int32(2112)
																							return
																						}
																					}
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v7 - int32(-64)
																					F_errmsg(m, int32(666400), v7)
																					mBase = m.M
																					v152 = m.ExcPending
																					if v152 != 0 {
																						return
																					} else {
																						F_errfinish(m, int32(484206), int32(1060), int32(203377))
																						mBase = m.M
																						v157 = m.ExcPending
																						if v157 != 0 {
																							return
																						} else {
																							v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
																							if v158 != 0 {
																								v162 = *(*int32)(unsafe.Add(mBase, _consts[562]))
																								v165 = base.I32_div_s(l0-v162, int32(288))
																								v167 = F_pgstat_drop_entry(m, int32(4), int32(0), base.I64_extend_i32_s(v165))
																								mBase = m.M
																								v168 = m.ExcPending
																								if v168 != 0 {
																									return
																								} else {
																									if v167 == int32(0) {
																										v173 = *(*int32)(unsafe.Add(mBase, _consts[563]))
																										v174 = *(*int64)(unsafe.Add(mBase, uint32(v173)+16))
																										*(*int64)(unsafe.Add(mBase, uint32(v173)+16)) = v174 + int64(1)
																									} else {
																									}
																									v179 = *(*int32)(unsafe.Add(mBase, _consts[7]))
																									F_LWLockRelease(m, v179+int32(4608))
																									mBase = m.M
																									v183 = m.ExcPending
																									if v183 != 0 {
																										return
																									} else {
																										m.G0 = v7 + int32(2112)
																										return
																									}
																								}
																							} else {
																								v179 = *(*int32)(unsafe.Add(mBase, _consts[7]))
																								F_LWLockRelease(m, v179+int32(4608))
																								mBase = m.M
																								v183 = m.ExcPending
																								if v183 != 0 {
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
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v4 = *(*int32)(unsafe.Add(mBase, _consts[519]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(1)
	if v5 != 0 {
		F_s_lock(m, v4, int32(484206), int32(1107), int32(8422))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, _consts[519]))
			v15 = int32(257)
			*(*uint16)(unsafe.Add(mBase, uint32(v14)+12)) = uint16(v15)
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(0)
			return
		}
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, _consts[519]))
		v15 = int32(257)
		*(*uint16)(unsafe.Add(mBase, uint32(v14)+12)) = uint16(v15)
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(0)
		return
	}
}
func F_ReplicationSlotValidateName(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	v3 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v3
	v18 = F_ReplicationSlotValidateNameInternal(m, l0, v6+int32(28), v6+int32(24), v6+int32(20))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		if v18 != 0 {
			m.G0 = v6 + int32(32)
			return v18
		} else {
			v23 = F_errstart(m, l1, int32(0))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				if v23 != 0 {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
					F_errcode(m, v25)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v6)+24))
						*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v28
						F_errmsg_internal(m, int32(202950), v6+int32(16))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
							if v35 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v6))) = v35
								F_errhint_internal(m, int32(202950), v6)
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(484206), int32(282), int32(375337))
									mBase = m.M
									v44 = m.ExcPending
									if v44 != 0 {
										return int32(0)
									} else {
										v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+24))
										F_pfree(m, v46)
										mBase = m.M
										v48 = m.ExcPending
										if v48 != 0 {
											return int32(0)
										} else {
											v49 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
											if v49 == int32(0) {
												m.G0 = v6 + int32(32)
												return v18
											} else {
												F_pfree(m, v49)
												mBase = m.M
												v53 = m.ExcPending
												if v53 != 0 {
													return int32(0)
												} else {
													m.G0 = v6 + int32(32)
													return v18
												}
											}
										}
									}
								}
							} else {
								F_errfinish(m, int32(484206), int32(282), int32(375337))
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return int32(0)
								} else {
									v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+24))
									F_pfree(m, v46)
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return int32(0)
									} else {
										v49 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
										if v49 == int32(0) {
											m.G0 = v6 + int32(32)
											return v18
										} else {
											F_pfree(m, v49)
											mBase = m.M
											v53 = m.ExcPending
											if v53 != 0 {
												return int32(0)
											} else {
												m.G0 = v6 + int32(32)
												return v18
											}
										}
									}
								}
							}
						}
					}
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+24))
					F_pfree(m, v46)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
						if v49 == int32(0) {
							m.G0 = v6 + int32(32)
							return v18
						} else {
							F_pfree(m, v49)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								m.G0 = v6 + int32(32)
								return v18
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
	*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = v112
	*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v110
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
