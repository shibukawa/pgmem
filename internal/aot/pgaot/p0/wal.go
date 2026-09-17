package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetWALInsertionTimeLine(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_GetWALInsertionTimeLine[0]))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+308))
	return v3
}
func F_InitializeWalConsistencyChecking(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	v3 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitializeWalConsistencyChecking[0])))
	if v3 != 0 {
		v5 = int32(0)
		v8 = F_find_option(m, int32(_a_F_InitializeWalConsistencyChecking_0), v5, v5, int32(21))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			v11 = int32(0)
			*(*uint8)(unsafe.Add(mBase, _c_F_InitializeWalConsistencyChecking[0])) = uint8(v11)
			v15 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeWalConsistencyChecking[1]))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
			F_set_config_option_ext(m, int32(_a_F_InitializeWalConsistencyChecking_0), v15, v16, v17, v18, v11, int32(21))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		return
	}
}
func F_ProcessWalSummarizerInterrupts(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessWalSummarizerInterrupts[0]))
	if v2 != 0 {
		F_ProcessProcSignalBarrier(m)
		mBase = m.M
		v4 = m.ExcPending
		if v4 != 0 {
			return
		} else {
			v6 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessWalSummarizerInterrupts[1]))
			if v6 != 0 {
				*(*int32)(unsafe.Add(mBase, _c_F_ProcessWalSummarizerInterrupts[1])) = int32(0)
				F_ProcessConfigFile(m, int32(2))
				mBase = m.M
				v12 = m.ExcPending
				if v12 != 0 {
					return
				} else {
					v14 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessWalSummarizerInterrupts[2]))
					v15 = int32(0)
					v18 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessWalSummarizerInterrupts[3])))
					if base.B2i32(v14 == v15)&(v18&int32(1)) == v15 {
						v26 = F_errstart(m, int32(14), int32(0))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							if v26 != 0 {
								F_errmsg_internal(m, int32(_a_F_ProcessWalSummarizerInterrupts_0), int32(0))
								mBase = m.M
								v31 = m.ExcPending
								if v31 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_ProcessWalSummarizerInterrupts_1), int32(875), int32(_a_F_ProcessWalSummarizerInterrupts_2))
									mBase = m.M
									v36 = m.ExcPending
									if v36 != 0 {
										return
									} else {
										F_proc_exit(m, int32(0))
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								F_proc_exit(m, int32(0))
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessWalSummarizerInterrupts[4]))
						if v41 != 0 {
							F_ProcessLogMemoryContextInterrupt(m)
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return
							} else {
								return
							}
						} else {
							return
						}
					}
				}
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessWalSummarizerInterrupts[2]))
				v15 = int32(0)
				v18 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessWalSummarizerInterrupts[3])))
				if base.B2i32(v14 == v15)&(v18&int32(1)) == v15 {
					v26 = F_errstart(m, int32(14), int32(0))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						if v26 != 0 {
							F_errmsg_internal(m, int32(_a_F_ProcessWalSummarizerInterrupts_0), int32(0))
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_ProcessWalSummarizerInterrupts_1), int32(875), int32(_a_F_ProcessWalSummarizerInterrupts_2))
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return
								} else {
									F_proc_exit(m, int32(0))
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							F_proc_exit(m, int32(0))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessWalSummarizerInterrupts[4]))
					if v41 != 0 {
						F_ProcessLogMemoryContextInterrupt(m)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							return
						}
					} else {
						return
					}
				}
			}
		}
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessWalSummarizerInterrupts[1]))
		if v6 != 0 {
			*(*int32)(unsafe.Add(mBase, _c_F_ProcessWalSummarizerInterrupts[1])) = int32(0)
			F_ProcessConfigFile(m, int32(2))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessWalSummarizerInterrupts[2]))
				v15 = int32(0)
				v18 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessWalSummarizerInterrupts[3])))
				if base.B2i32(v14 == v15)&(v18&int32(1)) == v15 {
					v26 = F_errstart(m, int32(14), int32(0))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						if v26 != 0 {
							F_errmsg_internal(m, int32(_a_F_ProcessWalSummarizerInterrupts_0), int32(0))
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_ProcessWalSummarizerInterrupts_1), int32(875), int32(_a_F_ProcessWalSummarizerInterrupts_2))
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return
								} else {
									F_proc_exit(m, int32(0))
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							F_proc_exit(m, int32(0))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessWalSummarizerInterrupts[4]))
					if v41 != 0 {
						F_ProcessLogMemoryContextInterrupt(m)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							return
						}
					} else {
						return
					}
				}
			}
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessWalSummarizerInterrupts[2]))
			v15 = int32(0)
			v18 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessWalSummarizerInterrupts[3])))
			if base.B2i32(v14 == v15)&(v18&int32(1)) == v15 {
				v26 = F_errstart(m, int32(14), int32(0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					if v26 != 0 {
						F_errmsg_internal(m, int32(_a_F_ProcessWalSummarizerInterrupts_0), int32(0))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_ProcessWalSummarizerInterrupts_1), int32(875), int32(_a_F_ProcessWalSummarizerInterrupts_2))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return
							} else {
								F_proc_exit(m, int32(0))
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						F_proc_exit(m, int32(0))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessWalSummarizerInterrupts[4]))
				if v41 != 0 {
					F_ProcessLogMemoryContextInterrupt(m)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						return
					}
				} else {
					return
				}
			}
		}
	}
}
func F_SetWalWriterSleeping(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	v1 = l0
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_SetWalWriterSleeping[0]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v5)+440)) = int32(1)
	if v6 != 0 {
		v10 = *(*int32)(unsafe.Add(mBase, _c_F_SetWalWriterSleeping[0]))
		F_s_lock(m, v10+int32(440), int32(_a_F_SetWalWriterSleeping_0), int32(_a_F_SetWalWriterSleeping_1), int32(_a_F_SetWalWriterSleeping_2))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, _c_F_SetWalWriterSleeping[0]))
			*(*int32)(unsafe.Add(mBase, uint32(v19)+440)) = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v19)+321)) = uint8(v1)
			return
		}
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, _c_F_SetWalWriterSleeping[0]))
		*(*int32)(unsafe.Add(mBase, uint32(v19)+440)) = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v19)+321)) = uint8(v1)
		return
	}
}
func F_WALInsertLockAcquireExclusive(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
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
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_WALInsertLockAcquireExclusive[0]))
	v5 = F_LWLockAcquire(m, v3, int32(0))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, _c_F_WALInsertLockAcquireExclusive[0]))
		F_LWLockUpdateVar(m, v8, v8+int32(16), int64(-1))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, _c_F_WALInsertLockAcquireExclusive[0]))
			v19 = F_LWLockAcquire(m, v15+int32(128), int32(0))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, _c_F_WALInsertLockAcquireExclusive[0]))
				F_LWLockUpdateVar(m, v22+int32(128), v22+int32(144), int64(-1))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, _c_F_WALInsertLockAcquireExclusive[0]))
					v35 = F_LWLockAcquire(m, v31+int32(256), int32(0))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, _c_F_WALInsertLockAcquireExclusive[0]))
						F_LWLockUpdateVar(m, v38+int32(256), v38+int32(272), int64(-1))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							v47 = *(*int32)(unsafe.Add(mBase, _c_F_WALInsertLockAcquireExclusive[0]))
							v51 = F_LWLockAcquire(m, v47+int32(384), int32(0))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								v54 = *(*int32)(unsafe.Add(mBase, _c_F_WALInsertLockAcquireExclusive[0]))
								F_LWLockUpdateVar(m, v54+int32(384), v54+int32(400), int64(-1))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return
								} else {
									v63 = *(*int32)(unsafe.Add(mBase, _c_F_WALInsertLockAcquireExclusive[0]))
									v67 = F_LWLockAcquire(m, v63+int32(512), int32(0))
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return
									} else {
										v70 = *(*int32)(unsafe.Add(mBase, _c_F_WALInsertLockAcquireExclusive[0]))
										F_LWLockUpdateVar(m, v70+int32(512), v70+int32(528), int64(-1))
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
											return
										} else {
											v79 = *(*int32)(unsafe.Add(mBase, _c_F_WALInsertLockAcquireExclusive[0]))
											v83 = F_LWLockAcquire(m, v79+int32(640), int32(0))
											mBase = m.M
											v84 = m.ExcPending
											if v84 != 0 {
												return
											} else {
												v86 = *(*int32)(unsafe.Add(mBase, _c_F_WALInsertLockAcquireExclusive[0]))
												F_LWLockUpdateVar(m, v86+int32(640), v86+int32(656), int64(-1))
												mBase = m.M
												v93 = m.ExcPending
												if v93 != 0 {
													return
												} else {
													v95 = *(*int32)(unsafe.Add(mBase, _c_F_WALInsertLockAcquireExclusive[0]))
													v99 = F_LWLockAcquire(m, v95+int32(768), int32(0))
													mBase = m.M
													v100 = m.ExcPending
													if v100 != 0 {
														return
													} else {
														v102 = *(*int32)(unsafe.Add(mBase, _c_F_WALInsertLockAcquireExclusive[0]))
														F_LWLockUpdateVar(m, v102+int32(768), v102+int32(784), int64(-1))
														mBase = m.M
														v109 = m.ExcPending
														if v109 != 0 {
															return
														} else {
															v111 = *(*int32)(unsafe.Add(mBase, _c_F_WALInsertLockAcquireExclusive[0]))
															v115 = F_LWLockAcquire(m, v111+int32(896), int32(0))
															mBase = m.M
															v116 = m.ExcPending
															if v116 != 0 {
																return
															} else {
																v118 = int32(1)
																*(*uint8)(unsafe.Add(mBase, _c_F_WALInsertLockAcquireExclusive[1])) = uint8(v118)
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
func F_WalRcvRunning(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int64
	_ = v24
	var v26 int64
	_ = v26
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_WalRcvRunning[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+1456))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+1456)) = int32(1)
	v11 = v6 + int32(1456)
	if v7 != 0 {
		F_s_lock(m, v11, int32(_a_F_WalRcvRunning_0), int32(82), int32(_a_F_WalRcvRunning_1))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6)+1456)) = int32(0)
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
			if v21 != int32(1) {
				v53 = v21
				return base.B2i32(v53 != int32(0))
			} else {
				v24 = *(*int64)(unsafe.Add(mBase, uint32(v6)+24))
				v26 = F_time(m)
				mBase = m.M
				if v26-v24 < int64(11) {
					v53 = int32(1)
					return base.B2i32(v53 != int32(0))
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(1)
					if v30 != 0 {
						F_s_lock(m, v11, int32(_a_F_WalRcvRunning_0), int32(103), int32(_a_F_WalRcvRunning_1))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
							if v38 != int32(1) {
								*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(0)
								v53 = int32(1)
								return base.B2i32(v53 != int32(0))
							} else {
								v44 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v6)+1456)) = v44
								*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v44
								F_ConditionVariableBroadcast(m, v6+int32(12))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									v53 = v44
									return base.B2i32(v53 != int32(0))
								}
							}
						}
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
						if v38 != int32(1) {
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(0)
							v53 = int32(1)
							return base.B2i32(v53 != int32(0))
						} else {
							v44 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+1456)) = v44
							*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v44
							F_ConditionVariableBroadcast(m, v6+int32(12))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								v53 = v44
								return base.B2i32(v53 != int32(0))
							}
						}
					}
				}
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+1456)) = int32(0)
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
		if v21 != int32(1) {
			v53 = v21
			return base.B2i32(v53 != int32(0))
		} else {
			v24 = *(*int64)(unsafe.Add(mBase, uint32(v6)+24))
			v26 = F_time(m)
			mBase = m.M
			if v26-v24 < int64(11) {
				v53 = int32(1)
				return base.B2i32(v53 != int32(0))
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(1)
				if v30 != 0 {
					F_s_lock(m, v11, int32(_a_F_WalRcvRunning_0), int32(103), int32(_a_F_WalRcvRunning_1))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
						if v38 != int32(1) {
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(0)
							v53 = int32(1)
							return base.B2i32(v53 != int32(0))
						} else {
							v44 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+1456)) = v44
							*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v44
							F_ConditionVariableBroadcast(m, v6+int32(12))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								v53 = v44
								return base.B2i32(v53 != int32(0))
							}
						}
					}
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
					if v38 != int32(1) {
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(0)
						v53 = int32(1)
						return base.B2i32(v53 != int32(0))
					} else {
						v44 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v6)+1456)) = v44
						*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v44
						F_ConditionVariableBroadcast(m, v6+int32(12))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							v53 = v44
							return base.B2i32(v53 != int32(0))
						}
					}
				}
			}
		}
	}
}
func F_WalSndWakeup(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	if l0 != 0 {
		v4 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndWakeup[0]))
		F_ConditionVariableBroadcast(m, v4+int32(52))
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			if l1 != 0 {
				v10 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndWakeup[0]))
				F_ConditionVariableBroadcast(m, v10-int32(-64))
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
					return
				} else {
					return
				}
			} else {
				return
			}
		}
	} else {
		if l1 != 0 {
			v10 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndWakeup[0]))
			F_ConditionVariableBroadcast(m, v10-int32(-64))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				return
			}
		} else {
			return
		}
	}
}
func F_WalSummarizerMain(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v54 int32
	_ = v54
	var v61 int64
	_ = v61
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v115 int32
	_ = v115
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int64
	_ = v137
	var v139 int64
	_ = v139
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v163 int32
	_ = v163
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int64
	_ = v185
	var v187 int64
	_ = v187
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v211 int32
	_ = v211
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int64
	_ = v233
	var v235 int64
	_ = v235
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v259 int32
	_ = v259
	var v271 int32
	_ = v271
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int64
	_ = v281
	var v283 int64
	_ = v283
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v307 int32
	_ = v307
	var v319 int32
	_ = v319
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int64
	_ = v329
	var v331 int64
	_ = v331
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v355 int32
	_ = v355
	var v367 int32
	_ = v367
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int64
	_ = v377
	var v379 int64
	_ = v379
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v403 int32
	_ = v403
	var v415 int32
	_ = v415
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v425 int64
	_ = v425
	var v427 int64
	_ = v427
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v491 int32
	_ = v491
	var v503 int32
	_ = v503
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v513 int64
	_ = v513
	var v515 int64
	_ = v515
	var v523 int32
	_ = v523
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v570 int32
	_ = v570
	var v575 int32
	_ = v575
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v609 int32
	_ = v609
	var v616 int64
	_ = v616
	var v617 int32
	_ = v617
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v637 int32
	_ = v637
	var v644 int64
	_ = v644
	var v647 int64
	_ = v647
	var v648 int64
	_ = v648
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v661 int64
	_ = v661
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v668 int64
	_ = v668
	var v674 int64
	_ = v674
	var v678 int32
	_ = v678
	var v680 int64
	_ = v680
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v690 int32
	_ = v690
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v729 int64
	_ = v729
	var v730 int32
	_ = v730
	var v733 int64
	_ = v733
	var v734 int64
	_ = v734
	var v735 int32
	_ = v735
	var v740 int32
	_ = v740
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v782 int64
	_ = v782
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v791 int64
	_ = v791
	var v794 int64
	_ = v794
	var v798 int64
	_ = v798
	var v799 int64
	_ = v799
	var v802 int64
	_ = v802
	var v805 int32
	_ = v805
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v832 int32
	_ = v832
	var v837 int32
	_ = v837
	var v838 int64
	_ = v838
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v852 int32
	_ = v852
	var v857 int32
	_ = v857
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v875 int32
	_ = v875
	var v880 int32
	_ = v880
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v899 int32
	_ = v899
	var v940 int64
	_ = v940
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v945 int32
	_ = v945
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v955 int64
	_ = v955
	var v956 int32
	_ = v956
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v967 int32
	_ = v967
	var v972 int64
	_ = v972
	var v978 int32
	_ = v978
	var v985 int32
	_ = v985
	var v987 int64
	_ = v987
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v998 int32
	_ = v998
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1007 int32
	_ = v1007
	var v1015 int32
	_ = v1015
	var v1019 int32
	_ = v1019
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1044 int32
	_ = v1044
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1056 int32
	_ = v1056
	var v1061 int32
	_ = v1061
	var v1067 int32
	_ = v1067
	var v1073 int32
	_ = v1073
	var v1080 int32
	_ = v1080
	var v1086 int32
	_ = v1086
	var v1089 int32
	_ = v1089
	var v1091 int32
	_ = v1091
	var v1093 int32
	_ = v1093
	var v1096 int64
	_ = v1096
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1128 int64
	_ = v1128
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1145 int32
	_ = v1145
	var v1148 int32
	_ = v1148
	var v1152 int32
	_ = v1152
	var v1163 int64
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1195 int32
	_ = v1195
	var v1197 int32
	_ = v1197
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1234 int64
	_ = v1234
	var v1236 int32
	_ = v1236
	var v1238 int32
	_ = v1238
	var v1241 int32
	_ = v1241
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1288 int64
	_ = v1288
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1321 int32
	_ = v1321
	var v1323 int32
	_ = v1323
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1360 int64
	_ = v1360
	var v1362 int32
	_ = v1362
	var v1364 int32
	_ = v1364
	var v1367 int32
	_ = v1367
	var v1379 int64
	_ = v1379
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1415 int32
	_ = v1415
	var v1417 int64
	_ = v1417
	var v1446 int64
	_ = v1446
	var v1457 int32
	_ = v1457
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1466 int64
	_ = v1466
	var v1470 int64
	_ = v1470
	var v1471 int64
	_ = v1471
	var v1476 int64
	_ = v1476
	var v1482 int32
	_ = v1482
	var v1489 int32
	_ = v1489
	var v1491 int64
	_ = v1491
	var v1515 int64
	_ = v1515
	var v1518 int64
	_ = v1518
	var v1522 int64
	_ = v1522
	var v1526 int32
	_ = v1526
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1539 int32
	_ = v1539
	var v1544 int64
	_ = v1544
	var v1550 int32
	_ = v1550
	var v1557 int32
	_ = v1557
	var v1560 int32
	_ = v1560
	var v1590 int32
	_ = v1590
	var v1591 int64
	_ = v1591
	var v1593 int32
	_ = v1593
	var v1594 int32
	_ = v1594
	var v1597 int32
	_ = v1597
	var v1599 int32
	_ = v1599
	var v1606 int32
	_ = v1606
	var v1613 int32
	_ = v1613
	var v1614 int32
	_ = v1614
	var v1617 int32
	_ = v1617
	var v1620 int32
	_ = v1620
	var v1621 int64
	_ = v1621
	var v1623 int32
	_ = v1623
	var v1625 int32
	_ = v1625
	var v1630 int32
	_ = v1630
	var v1635 int32
	_ = v1635
	var v1637 int32
	_ = v1637
	var v1640 int32
	_ = v1640
	var v1641 int32
	_ = v1641
	var v1643 int32
	_ = v1643
	var v1644 int32
	_ = v1644
	var v1654 int32
	_ = v1654
	var v1655 int32
	_ = v1655
	var v1656 int32
	_ = v1656
	var v1658 int32
	_ = v1658
	var v1659 int32
	_ = v1659
	var v1669 int32
	_ = v1669
	var v1670 int32
	_ = v1670
	var v1671 int32
	_ = v1671
	var v1672 int32
	_ = v1672
	var v1675 int32
	_ = v1675
	var v1681 int32
	_ = v1681
	var v1714 int32
	_ = v1714
	var v1720 int32
	_ = v1720
	var v1723 int32
	_ = v1723
	var v1725 int32
	_ = v1725
	var v1726 int32
	_ = v1726
	var v1728 int32
	_ = v1728
	var v1732 int32
	_ = v1732
	var v1735 int32
	_ = v1735
	var v1736 int32
	_ = v1736
	var v1743 int32
	_ = v1743
	var v1744 int32
	_ = v1744
	var v1745 int32
	_ = v1745
	var v1748 int32
	_ = v1748
	var v1755 int32
	_ = v1755
	var v1756 int32
	_ = v1756
	var v1758 int32
	_ = v1758
	var v1763 int32
	_ = v1763
	var v1770 int32
	_ = v1770
	var v1771 int32
	_ = v1771
	var v1776 int32
	_ = v1776
	var v1780 int32
	_ = v1780
	var v1781 int32
	_ = v1781
	var v1787 int64
	_ = v1787
	var v1792 int32
	_ = v1792
	var v1796 int32
	_ = v1796
	var v1798 int32
	_ = v1798
	var v1804 int32
	_ = v1804
	var v1807 int32
	_ = v1807
	var v1809 int32
	_ = v1809
	var v1815 int32
	_ = v1815
	var v1819 int32
	_ = v1819
	var v1821 int32
	_ = v1821
	var v1824 int32
	_ = v1824
	var v1828 int32
	_ = v1828
	var v1833 int32
	_ = v1833
	var v1834 int32
	_ = v1834
	var v1835 int32
	_ = v1835
	var v1838 int32
	_ = v1838
	var v1842 int32
	_ = v1842
	var v1847 int32
	_ = v1847
	var v1848 int32
	_ = v1848
	var v1849 int32
	_ = v1849
	var v1852 int32
	_ = v1852
	var v1856 int32
	_ = v1856
	var v1863 int32
	_ = v1863
	var v1866 int32
	_ = v1866
	var v1874 int32
	_ = v1874
	var v1875 int32
	_ = v1875
	var v1879 int32
	_ = v1879
	var v1880 int32
	_ = v1880
	var v1881 int32
	_ = v1881
	var v1886 int64
	_ = v1886
	var v1887 int64
	_ = v1887
	var v1895 int32
	_ = v1895
	var v1896 int32
	_ = v1896
	var v1899 int32
	_ = v1899
	var v1932 int32
	_ = v1932
	var v1933 int32
	_ = v1933
	var v1935 int32
	_ = v1935
	var v1938 int32
	_ = v1938
	var v1941 int32
	_ = v1941
	var v1946 int32
	_ = v1946
	var v1949 int32
	_ = v1949
	var v1954 int32
	_ = v1954
	var v1956 int32
	_ = v1956
	var v1957 int32
	_ = v1957
	var v1959 int32
	_ = v1959
	var v1963 int32
	_ = v1963
	var v1964 int32
	_ = v1964
	var v1970 int64
	_ = v1970
	var v1975 int32
	_ = v1975
	var v1979 int32
	_ = v1979
	var v1981 int32
	_ = v1981
	var v1987 int32
	_ = v1987
	var v1990 int32
	_ = v1990
	var v1992 int32
	_ = v1992
	var v1998 int32
	_ = v1998
	var v2002 int32
	_ = v2002
	var v2004 int32
	_ = v2004
	var v2007 int32
	_ = v2007
	var v2011 int32
	_ = v2011
	var v2016 int32
	_ = v2016
	var v2017 int32
	_ = v2017
	var v2018 int32
	_ = v2018
	var v2021 int32
	_ = v2021
	var v2025 int32
	_ = v2025
	var v2032 int32
	_ = v2032
	var v2035 int32
	_ = v2035
	var v2043 int32
	_ = v2043
	var v2044 int32
	_ = v2044
	var v2048 int32
	_ = v2048
	var v2049 int32
	_ = v2049
	var v2050 int32
	_ = v2050
	var v2055 int64
	_ = v2055
	var v2056 int64
	_ = v2056
	var v2064 int32
	_ = v2064
	var v2065 int32
	_ = v2065
	var v2068 int32
	_ = v2068
	var v2101 int32
	_ = v2101
	var v2102 int32
	_ = v2102
	var v2104 int32
	_ = v2104
	var v2107 int32
	_ = v2107
	var v2110 int32
	_ = v2110
	var v2115 int32
	_ = v2115
	var v2118 int32
	_ = v2118
	var v2123 int32
	_ = v2123
	var v2125 int32
	_ = v2125
	var v2126 int32
	_ = v2126
	var v2158 int32
	_ = v2158
	var v2159 int32
	_ = v2159
	var v2160 int32
	_ = v2160
	var v2163 int32
	_ = v2163
	var v2196 int32
	_ = v2196
	var v2198 int32
	_ = v2198
	var v2200 int32
	_ = v2200
	var v2202 int32
	_ = v2202
	var v2203 int32
	_ = v2203
	var v2205 int32
	_ = v2205
	var v2206 int32
	_ = v2206
	var v2210 int32
	_ = v2210
	var v2211 int32
	_ = v2211
	var v2215 int32
	_ = v2215
	var v2216 int32
	_ = v2216
	var v2218 int64
	_ = v2218
	var v2220 int32
	_ = v2220
	var v2222 int32
	_ = v2222
	var v2230 int32
	_ = v2230
	var v2233 int32
	_ = v2233
	var v2238 int32
	_ = v2238
	var v2240 int32
	_ = v2240
	var v2243 int32
	_ = v2243
	var v2244 int32
	_ = v2244
	var v2245 int32
	_ = v2245
	var v2248 int32
	_ = v2248
	var v2278 int64
	_ = v2278
	var v2282 int32
	_ = v2282
	var v2286 int32
	_ = v2286
	var v2287 int32
	_ = v2287
	var v2289 int32
	_ = v2289
	var v2294 int32
	_ = v2294
	var v2298 int32
	_ = v2298
	var v2301 int64
	_ = v2301
	var v2306 int32
	_ = v2306
	var v2311 int32
	_ = v2311
	var v2312 int32
	_ = v2312
	var v2313 int32
	_ = v2313
	var v2343 int32
	_ = v2343
	var v2350 int32
	_ = v2350
	var v2351 int32
	_ = v2351
	var v2352 int64
	_ = v2352
	var v2353 int64
	_ = v2353
	var v2357 int64
	_ = v2357
	var v2358 int64
	_ = v2358
	var v2362 int64
	_ = v2362
	var v2369 int32
	_ = v2369
	var v2376 int32
	_ = v2376
	var v2379 int64
	_ = v2379
	var v2380 int32
	_ = v2380
	var v2405 int64
	_ = v2405
	var v2411 int32
	_ = v2411
	var v2435 int64
	_ = v2435
	var v2440 int32
	_ = v2440
	var v2444 int32
	_ = v2444
	var v2448 int32
	_ = v2448
	var v2449 int32
	_ = v2449
	var v2456 int32
	_ = v2456
	var v2460 int32
	_ = v2460
	var v2461 int32
	_ = v2461
	var v2464 int32
	_ = v2464
	var v2466 int64
	_ = v2466
	var v2468 int32
	_ = v2468
	var v2470 int32
	_ = v2470
	var v2474 int32
	_ = v2474
	var v2483 int32
	_ = v2483
	var v2484 int32
	_ = v2484
	var v2490 int32
	_ = v2490
	var v2491 int32
	_ = v2491
	var v2500 int32
	_ = v2500
	var v2504 int32
	_ = v2504
	var v2512 int32
	_ = v2512
	var v2519 int32
	_ = v2519
	var v2522 int32
	_ = v2522
	var v2526 int32
	_ = v2526
	var v2530 int32
	_ = v2530
	var v2531 int64
	_ = v2531
	var v2534 int32
	_ = v2534
	var v2535 int32
	_ = v2535
	var v2538 int32
	_ = v2538
	var v2547 int32
	_ = v2547
	var v2554 int32
	_ = v2554
	var v2564 int32
	_ = v2564
	var v2571 int32
	_ = v2571
	var v2574 int32
	_ = v2574
	var v2575 int32
	_ = v2575
	var v2577 int32
	_ = v2577
	var v2591 int32
	_ = v2591
	var v2595 int32
	_ = v2595
	var v2596 int32
	_ = v2596
	var v2598 int32
	_ = v2598
	var v2601 int32
	_ = v2601
	var v2603 int32
	_ = v2603
	var v2608 int32
	_ = v2608
	var v2609 int32
	_ = v2609
	var v2610 int32
	_ = v2610
	var v2611 int64
	_ = v2611
	var v2614 int32
	_ = v2614
	var v2621 int32
	_ = v2621
	var v2649 int32
	_ = v2649
	var v2653 int32
	_ = v2653
	var v2656 int32
	_ = v2656
	var v2687 int32
	_ = v2687
	var v2693 int32
	_ = v2693
	var v2695 int32
	_ = v2695
	var v2698 int32
	_ = v2698
	var v2703 int32
	_ = v2703
	var v2724 int32
	_ = v2724
	var v2725 int32
	_ = v2725
	var v2727 int32
	_ = v2727
	var v2751 int32
	_ = v2751
	var v2752 int32
	_ = v2752
	var v2753 int32
	_ = v2753
	var v2757 int32
	_ = v2757
	var v2758 int32
	_ = v2758
	var v2761 int32
	_ = v2761
	var v2762 int32
	_ = v2762
	var v2763 int32
	_ = v2763
	var v2764 int32
	_ = v2764
	var v2769 int32
	_ = v2769
	var v2770 int32
	_ = v2770
	var v2772 int64
	_ = v2772
	var v2774 int32
	_ = v2774
	var v2776 int32
	_ = v2776
	var v2778 int32
	_ = v2778
	var v2784 int32
	_ = v2784
	var v2812 int32
	_ = v2812
	var v2818 int32
	_ = v2818
	var v2823 int32
	_ = v2823
	var v2827 int32
	_ = v2827
	var v2828 int32
	_ = v2828
	var v2829 int32
	_ = v2829
	var v2837 int32
	_ = v2837
	var v2863 int32
	_ = v2863
	var v2864 int32
	_ = v2864
	var v2866 int32
	_ = v2866
	var v2868 int32
	_ = v2868
	var v2870 int32
	_ = v2870
	var v2875 int32
	_ = v2875
	var v2876 int32
	_ = v2876
	var v2877 int32
	_ = v2877
	var v2878 int32
	_ = v2878
	var v2879 int32
	_ = v2879
	var v2882 int32
	_ = v2882
	var v2883 int32
	_ = v2883
	var v2884 int64
	_ = v2884
	var v2886 int64
	_ = v2886
	var v2888 int64
	_ = v2888
	var v2890 int32
	_ = v2890
	var v2894 int32
	_ = v2894
	var v2895 int32
	_ = v2895
	var v2896 int32
	_ = v2896
	var v2898 int64
	_ = v2898
	var v2902 int32
	_ = v2902
	var v2903 int32
	_ = v2903
	var v2909 int32
	_ = v2909
	var v2916 int32
	_ = v2916
	var v2917 int32
	_ = v2917
	var v2918 int32
	_ = v2918
	var v2919 int32
	_ = v2919
	var v2920 int32
	_ = v2920
	var v2922 int32
	_ = v2922
	var v2923 int32
	_ = v2923
	var v2924 int32
	_ = v2924
	var v2926 int32
	_ = v2926
	var v2927 int32
	_ = v2927
	var v2929 int32
	_ = v2929
	var v2931 int32
	_ = v2931
	var v2935 int32
	_ = v2935
	var v2936 int32
	_ = v2936
	var v2937 int32
	_ = v2937
	var v2938 int32
	_ = v2938
	var v2942 int32
	_ = v2942
	var v2946 int32
	_ = v2946
	var v2950 int32
	_ = v2950
	var v2951 int32
	_ = v2951
	var v2952 int32
	_ = v2952
	var v2953 int32
	_ = v2953
	var v2957 int32
	_ = v2957
	var v2958 int32
	_ = v2958
	var v2959 int32
	_ = v2959
	var v2961 int32
	_ = v2961
	var v2972 int32
	_ = v2972
	var v2976 int32
	_ = v2976
	var v2977 int32
	_ = v2977
	var v2981 int32
	_ = v2981
	var v2982 int32
	_ = v2982
	var v2986 int32
	_ = v2986
	var v2987 int32
	_ = v2987
	var v2989 int32
	_ = v2989
	var v2991 int32
	_ = v2991
	var v2995 int32
	_ = v2995
	var v2996 int32
	_ = v2996
	var v3000 int32
	_ = v3000
	var v3001 int32
	_ = v3001
	var v3003 int32
	_ = v3003
	var v3004 int32
	_ = v3004
	var v3006 int32
	_ = v3006
	var v3010 int32
	_ = v3010
	var v3011 int32
	_ = v3011
	var v3015 int32
	_ = v3015
	var v3016 int32
	_ = v3016
	var v3018 int32
	_ = v3018
	var v3019 int32
	_ = v3019
	var v3020 int32
	_ = v3020
	var v3021 int32
	_ = v3021
	var v3022 int32
	_ = v3022
	var v3024 int32
	_ = v3024
	var v3025 int32
	_ = v3025
	var v3026 int32
	_ = v3026
	var v3028 int32
	_ = v3028
	var v3029 int32
	_ = v3029
	var v3031 int32
	_ = v3031
	var v3033 int32
	_ = v3033
	var v3037 int32
	_ = v3037
	var v3038 int32
	_ = v3038
	var v3039 int32
	_ = v3039
	var v3040 int32
	_ = v3040
	var v3044 int32
	_ = v3044
	var v3048 int32
	_ = v3048
	var v3052 int32
	_ = v3052
	var v3053 int32
	_ = v3053
	var v3054 int32
	_ = v3054
	var v3055 int32
	_ = v3055
	var v3059 int32
	_ = v3059
	var v3060 int32
	_ = v3060
	var v3061 int32
	_ = v3061
	var v3063 int32
	_ = v3063
	var v3074 int32
	_ = v3074
	var v3078 int32
	_ = v3078
	var v3079 int32
	_ = v3079
	var v3083 int32
	_ = v3083
	var v3084 int32
	_ = v3084
	var v3088 int32
	_ = v3088
	var v3089 int32
	_ = v3089
	var v3093 int32
	_ = v3093
	var v3094 int32
	_ = v3094
	var v3095 int32
	_ = v3095
	var v3099 int32
	_ = v3099
	var v3100 int32
	_ = v3100
	var v3101 int32
	_ = v3101
	var v3105 int32
	_ = v3105
	var v3106 int32
	_ = v3106
	var v3107 int32
	_ = v3107
	var v3109 int32
	_ = v3109
	var v3110 int32
	_ = v3110
	var v3111 int32
	_ = v3111
	var v3115 int32
	_ = v3115
	var v3116 int32
	_ = v3116
	var v3117 int32
	_ = v3117
	var v3118 int32
	_ = v3118
	var v3122 int32
	_ = v3122
	var v3123 int32
	_ = v3123
	var v3124 int32
	_ = v3124
	var v3125 int32
	_ = v3125
	var v3129 int32
	_ = v3129
	var v3130 int32
	_ = v3130
	var v3131 int32
	_ = v3131
	var v3132 int32
	_ = v3132
	var v3136 int32
	_ = v3136
	var v3137 int32
	_ = v3137
	var v3138 int32
	_ = v3138
	var v3141 int32
	_ = v3141
	var v3143 int32
	_ = v3143
	var v3147 int32
	_ = v3147
	var v3151 int32
	_ = v3151
	var v3155 int32
	_ = v3155
	var v3159 int32
	_ = v3159
	var v3163 int32
	_ = v3163
	var v3168 int32
	_ = v3168
	var v3169 int32
	_ = v3169
	var v3170 int32
	_ = v3170
	var v3173 int32
	_ = v3173
	var v3174 int32
	_ = v3174
	var v3183 int32
	_ = v3183
	var v3184 int32
	_ = v3184
	var v3208 int64
	_ = v3208
	var v3209 int64
	_ = v3209
	var v3211 int64
	_ = v3211
	var v3214 int64
	_ = v3214
	var v3221 int32
	_ = v3221
	var v3224 int32
	_ = v3224
	var v3225 int32
	_ = v3225
	var v3232 int32
	_ = v3232
	var v3257 int32
	_ = v3257
	var v3260 int32
	_ = v3260
	var v3261 int32
	_ = v3261
	var v3263 int32
	_ = v3263
	var v3264 int32
	_ = v3264
	var v3266 int32
	_ = v3266
	var v3270 int32
	_ = v3270
	var v3271 int32
	_ = v3271
	var v3272 int32
	_ = v3272
	var v3273 int32
	_ = v3273
	var v3274 int32
	_ = v3274
	var v3277 int32
	_ = v3277
	var v3280 int32
	_ = v3280
	var v3281 int32
	_ = v3281
	var v3282 int32
	_ = v3282
	var v3283 int32
	_ = v3283
	var v3286 int32
	_ = v3286
	var v3292 int32
	_ = v3292
	var v3300 int32
	_ = v3300
	var v3324 int32
	_ = v3324
	var v3328 int32
	_ = v3328
	var v3332 int32
	_ = v3332
	var v3333 int32
	_ = v3333
	var v3337 int32
	_ = v3337
	var v3339 int32
	_ = v3339
	var v3340 int32
	_ = v3340
	var v3342 int32
	_ = v3342
	var v3346 int32
	_ = v3346
	var v3347 int32
	_ = v3347
	var v3348 int32
	_ = v3348
	var v3349 int32
	_ = v3349
	var v3350 int32
	_ = v3350
	var v3353 int32
	_ = v3353
	var v3356 int32
	_ = v3356
	var v3357 int32
	_ = v3357
	var v3358 int32
	_ = v3358
	var v3359 int32
	_ = v3359
	var v3362 int32
	_ = v3362
	var v3369 int32
	_ = v3369
	var v3370 int32
	_ = v3370
	var v3403 int32
	_ = v3403
	var v3404 int32
	_ = v3404
	var v3405 int32
	_ = v3405
	var v3437 int32
	_ = v3437
	var v3439 int32
	_ = v3439
	var v3441 int64
	_ = v3441
	var v3448 int32
	_ = v3448
	var v3449 int32
	_ = v3449
	var v3452 int32
	_ = v3452
	var v3453 int32
	_ = v3453
	var v3455 int32
	_ = v3455
	var v3460 int32
	_ = v3460
	var v3463 int32
	_ = v3463
	var v3464 int32
	_ = v3464
	var v3465 int32
	_ = v3465
	var v3466 int32
	_ = v3466
	var v3469 int32
	_ = v3469
	var v3471 int32
	_ = v3471
	var v3472 int32
	_ = v3472
	var v3473 int64
	_ = v3473
	var v3475 int64
	_ = v3475
	var v3477 int64
	_ = v3477
	var v3479 int32
	_ = v3479
	var v3483 int32
	_ = v3483
	var v3487 int32
	_ = v3487
	var v3490 int32
	_ = v3490
	var v3492 int32
	_ = v3492
	var v3497 int32
	_ = v3497
	var v3498 int32
	_ = v3498
	var v3499 int32
	_ = v3499
	var v3500 int32
	_ = v3500
	var v3501 int32
	_ = v3501
	var v3504 int32
	_ = v3504
	var v3506 int32
	_ = v3506
	var v3508 int32
	_ = v3508
	var v3510 int32
	_ = v3510
	var v3512 int32
	_ = v3512
	var v3513 int32
	_ = v3513
	var v3514 int32
	_ = v3514
	var v3515 int32
	_ = v3515
	var v3526 int32
	_ = v3526
	var v3528 int32
	_ = v3528
	var v3533 int32
	_ = v3533
	var v3534 int32
	_ = v3534
	var v3548 int32
	_ = v3548
	var v3555 int32
	_ = v3555
	var v3563 int32
	_ = v3563
	var v3564 int32
	_ = v3564
	var v3604 int32
	_ = v3604
	var v3605 int32
	_ = v3605
	var v3611 int64
	_ = v3611
	var v3612 int64
	_ = v3612
	var v3617 int64
	_ = v3617
	var v3621 int32
	_ = v3621
	var v3628 int32
	_ = v3628
	var v3629 int32
	_ = v3629
	var v3634 int32
	_ = v3634
	var v3638 int32
	_ = v3638
	var v3639 int32
	_ = v3639
	var v3641 int32
	_ = v3641
	var v3643 int32
	_ = v3643
	var v3650 int32
	_ = v3650
	var v3654 int32
	_ = v3654
	var v3658 int32
	_ = v3658
	var v3662 int32
	_ = v3662
	var v3686 int64
	_ = v3686
	var v3687 int64
	_ = v3687
	var v3690 int64
	_ = v3690
	var v3697 int32
	_ = v3697
	var v3700 int32
	_ = v3700
	var v3714 int32
	_ = v3714
	var v3728 int32
	_ = v3728
	var v3729 int64
	_ = v3729
	var v3733 int32
	_ = v3733
	var v3735 int32
	_ = v3735
	var v3736 int32
	_ = v3736
	var v3739 int32
	_ = v3739
	var v3741 int32
	_ = v3741
	var v3743 int32
	_ = v3743
	var v3744 int32
	_ = v3744
	var v3745 int64
	_ = v3745
	var v3747 int32
	_ = v3747
	v3 = int32(0)
	v31 = m.G0
	v33 = v31 - int32(3136)
	m.G0 = v33
	v38 = int32(-1)
	v40 = v33
	v41 = v3
	v43 = v3
	v54 = v33 + int32(128)
	v61 = int64(0)
	goto L1
L1:
	;
	goto L4
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	goto L2
L4:
	;
	if v38 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v3728 = int32(m.ExcTag)
	v3729 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v3728 == int32(0) {
		goto L611
	} else {
		goto L612
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[0])) = int32(15)
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v40)+260)) = int32(0)
	F_AuxiliaryProcessMainCommon(m)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		v3700 = v40
		v3714 = v54
		goto L6
	} else {
		goto L10
	}
L8:
	;
	v530 = v41
	v531 = v43
	goto L9
L9:
	;
	if v531 != 0 {
		goto L129
	} else {
		goto L130
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v41
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	v83 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		v3700 = v40
		v3714 = v54
		goto L6
	} else {
		goto L11
	}
L11:
	;
	if v83 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v41
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	F_errmsg_internal(m, int32(_a_F_WalSummarizerMain_0), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		v3700 = v40
		v3714 = v54
		goto L6
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v41
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	v101 = int32(914)
	v103 = m.G0
	v105 = v103 - int32(32)
	m.G0 = v105
	switch int32(916) {
	case 0, 2:
		v115 = v101
		goto L18
	default:
		goto L19
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v41
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	F_errfinish(m, int32(_a_F_WalSummarizerMain_1), int32(241), int32(_a_F_WalSummarizerMain_2))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		v3700 = v40
		v3714 = v54
		goto L6
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v41
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	v149 = int32(916)
	v151 = m.G0
	v153 = v151 - int32(32)
	m.G0 = v153
	switch int32(918) {
	case 0, 2:
		v163 = v149
		goto L31
	default:
		goto L32
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105)+12)) = v115
	F_sigemptyset(m, v105+int32(16))
	mBase = m.M
	goto L21
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[1])) = v101
	v115 = int32(_a_F_WalSummarizerMain_3)
	goto L18
L21:
	;
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105)+24)) = int32(268435456)
	v127 = v105 + int32(12)
	goto L25
L23:
	;
	m.G0 = v105 + int32(32)
	goto L17
L25:
	;
	goto L26
L26:
	;
	if v127 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v133 = int32(20)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v127)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[2])) = v135
	v137 = *(*int64)(unsafe.Add(mBase, uint32(v127)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_WalSummarizerMain[3])) = v137
	v139 = *(*int64)(unsafe.Add(mBase, uint32(v127)))
	*(*int64)(unsafe.Add(mBase, _c_F_WalSummarizerMain[4])) = v139
	goto L29
L28:
	;
	goto L29
L29:
	;
	goto L23
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v41
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	v197 = int32(916)
	v199 = m.G0
	v201 = v199 - int32(32)
	m.G0 = v201
	switch int32(918) {
	case 0, 2:
		v211 = v197
		goto L44
	default:
		goto L45
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153)+12)) = v163
	F_sigemptyset(m, v153+int32(16))
	mBase = m.M
	goto L34
L32:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[5])) = v149
	v163 = int32(_a_F_WalSummarizerMain_3)
	goto L31
L34:
	;
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153)+24)) = int32(268435456)
	v175 = v153 + int32(12)
	goto L38
L36:
	;
	m.G0 = v153 + int32(32)
	goto L30
L38:
	;
	goto L39
L39:
	;
	if v175 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v182 = int32(40)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v175)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[6])) = v183
	v185 = *(*int64)(unsafe.Add(mBase, uint32(v175)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_WalSummarizerMain[7])) = v185
	v187 = *(*int64)(unsafe.Add(mBase, uint32(v175)))
	*(*int64)(unsafe.Add(mBase, _c_F_WalSummarizerMain[8])) = v187
	goto L42
L41:
	;
	goto L42
L42:
	;
	goto L36
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v41
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	v245 = int32(-2)
	v247 = m.G0
	v249 = v247 - int32(32)
	m.G0 = v249
	switch int32(0) {
	case 0, 2:
		v259 = v245
		goto L57
	default:
		goto L58
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v201)+12)) = v211
	F_sigemptyset(m, v201+int32(16))
	mBase = m.M
	goto L47
L45:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[9])) = v197
	v211 = int32(_a_F_WalSummarizerMain_3)
	goto L44
L47:
	;
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v201)+24)) = int32(268435456)
	v223 = v201 + int32(12)
	goto L51
L49:
	;
	m.G0 = v201 + int32(32)
	goto L43
L51:
	;
	goto L52
L52:
	;
	if v223 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v230 = int32(300)
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v223)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[10])) = v231
	v233 = *(*int64)(unsafe.Add(mBase, uint32(v223)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_WalSummarizerMain[11])) = v233
	v235 = *(*int64)(unsafe.Add(mBase, uint32(v223)))
	*(*int64)(unsafe.Add(mBase, _c_F_WalSummarizerMain[12])) = v235
	goto L55
L54:
	;
	goto L55
L55:
	;
	goto L49
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v41
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	v293 = int32(-2)
	v295 = m.G0
	v297 = v295 - int32(32)
	m.G0 = v297
	switch int32(0) {
	case 0, 2:
		v307 = v293
		goto L70
	default:
		goto L71
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v249)+12)) = v259
	F_sigemptyset(m, v249+int32(16))
	mBase = m.M
	goto L60
L58:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[13])) = v245
	v259 = int32(_a_F_WalSummarizerMain_3)
	goto L57
L60:
	;
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v249)+24)) = int32(268435456)
	v271 = v249 + int32(12)
	goto L64
L62:
	;
	m.G0 = v249 + int32(32)
	goto L56
L64:
	;
	goto L65
L65:
	;
	if v271 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v278 = int32(280)
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v271)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[14])) = v279
	v281 = *(*int64)(unsafe.Add(mBase, uint32(v271)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_WalSummarizerMain[15])) = v281
	v283 = *(*int64)(unsafe.Add(mBase, uint32(v271)))
	*(*int64)(unsafe.Add(mBase, _c_F_WalSummarizerMain[16])) = v283
	goto L68
L67:
	;
	goto L68
L68:
	;
	goto L62
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v41
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	v341 = int32(917)
	v343 = m.G0
	v345 = v343 - int32(32)
	m.G0 = v345
	switch int32(919) {
	case 0, 2:
		v355 = v341
		goto L83
	default:
		goto L84
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v297)+12)) = v307
	F_sigemptyset(m, v297+int32(16))
	mBase = m.M
	goto L73
L71:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[17])) = v293
	v307 = int32(_a_F_WalSummarizerMain_3)
	goto L70
L73:
	;
	goto L74
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v297)+24)) = int32(268435456)
	v319 = v297 + int32(12)
	goto L77
L75:
	;
	m.G0 = v297 + int32(32)
	goto L69
L77:
	;
	goto L78
L78:
	;
	if v319 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v326 = int32(260)
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v319)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[18])) = v327
	v329 = *(*int64)(unsafe.Add(mBase, uint32(v319)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_WalSummarizerMain[19])) = v329
	v331 = *(*int64)(unsafe.Add(mBase, uint32(v319)))
	*(*int64)(unsafe.Add(mBase, _c_F_WalSummarizerMain[20])) = v331
	goto L81
L80:
	;
	goto L81
L81:
	;
	goto L75
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v41
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	v389 = int32(-2)
	v391 = m.G0
	v393 = v391 - int32(32)
	m.G0 = v393
	switch int32(0) {
	case 0, 2:
		v403 = v389
		goto L96
	default:
		goto L97
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v345)+12)) = v355
	F_sigemptyset(m, v345+int32(16))
	mBase = m.M
	goto L86
L84:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[21])) = v341
	v355 = int32(_a_F_WalSummarizerMain_3)
	goto L83
L86:
	;
	goto L87
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v345)+24)) = int32(268435456)
	v367 = v345 + int32(12)
	goto L90
L88:
	;
	m.G0 = v345 + int32(32)
	goto L82
L90:
	;
	goto L91
L91:
	;
	if v367 != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v374 = int32(200)
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v367)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[22])) = v375
	v377 = *(*int64)(unsafe.Add(mBase, uint32(v367)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_WalSummarizerMain[23])) = v377
	v379 = *(*int64)(unsafe.Add(mBase, uint32(v367)))
	*(*int64)(unsafe.Add(mBase, _c_F_WalSummarizerMain[24])) = v379
	goto L94
L93:
	;
	goto L94
L94:
	;
	goto L88
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v41
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	F_on_shmem_exit(m, int32(965), int32(0))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		v3700 = v40
		v3714 = v54
		goto L6
	} else {
		goto L108
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v393)+12)) = v403
	F_sigemptyset(m, v393+int32(16))
	mBase = m.M
	goto L99
L97:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[25])) = v389
	v403 = int32(_a_F_WalSummarizerMain_3)
	goto L96
L99:
	;
	goto L100
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v393)+24)) = int32(268435456)
	v415 = v393 + int32(12)
	goto L103
L101:
	;
	m.G0 = v393 + int32(32)
	goto L95
L103:
	;
	goto L104
L104:
	;
	if v415 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v422 = int32(240)
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v415)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[26])) = v423
	v425 = *(*int64)(unsafe.Add(mBase, uint32(v415)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_WalSummarizerMain[27])) = v425
	v427 = *(*int64)(unsafe.Add(mBase, uint32(v415)))
	*(*int64)(unsafe.Add(mBase, _c_F_WalSummarizerMain[28])) = v427
	goto L107
L106:
	;
	goto L107
L107:
	;
	goto L101
L108:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v41
	v443 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[29]))
	v447 = F_LWLockAcquire(m, v443+int32(_a_F_WalSummarizerMain_4), int32(0))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		v3700 = v40
		v3714 = v54
		goto L6
	} else {
		goto L109
	}
L109:
	;
	v450 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[30]))
	v452 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[31]))
	*(*int32)(unsafe.Add(mBase, uint32(v450)+20)) = v452
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v41
	v457 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[29]))
	F_LWLockRelease(m, v457+int32(_a_F_WalSummarizerMain_4))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		v3700 = v40
		v3714 = v54
		goto L6
	} else {
		goto L110
	}
L110:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v41
	v466 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[32]))
	v471 = F_AllocSetContextCreateInternal(m, v466, int32(_a_F_WalSummarizerMain_5), int32(0), int32(_a_F_WalSummarizerMain_6), int32(_a_F_WalSummarizerMain_7))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		v3700 = v40
		v3714 = v54
		goto L6
	} else {
		goto L111
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[33])) = v471
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v471
	v477 = int32(0)
	v479 = m.G0
	v481 = v479 - int32(32)
	m.G0 = v481
	switch int32(2) {
	case 0, 2:
		v491 = v477
		goto L113
	default:
		goto L114
	}
L112:
	;
	goto L125
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v481)+12)) = v491
	F_sigemptyset(m, v481+int32(16))
	mBase = m.M
	goto L115
L114:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[34])) = v477
	v491 = int32(_a_F_WalSummarizerMain_3)
	goto L113
L115:
	;
	goto L117
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v481)+24)) = int32(268435457)
	v503 = v481 + int32(12)
	goto L120
L118:
	;
	m.G0 = v481 + int32(32)
	goto L112
L120:
	;
	goto L121
L121:
	;
	if v503 != 0 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v510 = int32(340)
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v503)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[35])) = v511
	v513 = *(*int64)(unsafe.Add(mBase, uint32(v503)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_WalSummarizerMain[36])) = v513
	v515 = *(*int64)(unsafe.Add(mBase, uint32(v503)))
	*(*int64)(unsafe.Add(mBase, _c_F_WalSummarizerMain[37])) = v515
	goto L124
L123:
	;
	goto L124
L124:
	;
	goto L118
L125:
	;
	v523 = v40 + int32(272)
	*(*int32)(unsafe.Add(mBase, uint32(v523)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v523))) = v40 + int32(252)
	goto L128
L126:
	;
	v530 = v471
	v531 = int32(0)
	goto L9
L128:
	;
	goto L126
L129:
	;
	v532 = int32(_a_F_WalSummarizerMain_8)
	v534 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[38]))
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[38])) = v534 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[39])) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v530
	F_EmitErrorReport(m)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		v3700 = v40
		v3714 = v54
		goto L6
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[40])) = v40 + int32(272)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v530
	F_sigprocmask(m, int32(_a_F_WalSummarizerMain_9), int32(0))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		v3700 = v40
		v3714 = v54
		goto L6
	} else {
		goto L142
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v530
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	F_LWLockReleaseAll(m)
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		v3700 = v40
		v3714 = v54
		goto L6
	} else {
		goto L133
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v530
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		v3700 = v40
		v3714 = v54
		goto L6
	} else {
		goto L134
	}
L134:
	;
	v554 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[41]))
	*(*int32)(unsafe.Add(mBase, uint32(v554))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v530
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	F_pgaio_error_cleanup(m)
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		v3700 = v40
		v3714 = v54
		goto L6
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v530
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	F_ReleaseAuxProcessResources(m, int32(0))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		v3700 = v40
		v3714 = v54
		goto L6
	} else {
		goto L136
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v530
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	F_AtEOXact_Files(m, int32(0))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		v3700 = v40
		v3714 = v54
		goto L6
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v530
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	F_AtEOXact_HashTables(m, int32(0))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		v3700 = v40
		v3714 = v54
		goto L6
	} else {
		goto L138
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[33])) = v530
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v530
	F_FlushErrorState(m)
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		v3700 = v40
		v3714 = v54
		goto L6
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v530
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	F_MemoryContextReset(m, v530)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		v3700 = v40
		v3714 = v54
		goto L6
	} else {
		goto L140
	}
L140:
	;
	v586 = int32(_a_F_WalSummarizerMain_8)
	v588 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[38]))
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[38])) = v588 - int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v530
	v598 = F_WaitLatch(m, int32(0), int32(40), int32(_a_F_WalSummarizerMain_10), int32(150994953))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		v3700 = v40
		v3714 = v54
		goto L6
	} else {
		goto L141
	}
L141:
	;
	goto L131
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v530
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	v616 = F_GetOldestUnsummarizedLSN(m, v40+int32(268), v40+int32(267))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		v3700 = v40
		v3714 = v54
		goto L6
	} else {
		goto L143
	}
L143:
	;
	if v616 != int64(0) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v623 = v40
	v624 = v530
	v637 = v54
	v644 = v61
	v647 = v616
	v648 = int64(0)
	goto L147
L145:
	;
	goto L146
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v530
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	F_proc_exit(m, int32(0))
	mBase = m.M
	v3697 = m.ExcPending
	if v3697 != 0 {
		v3700 = v40
		v3714 = v54
		goto L6
	} else {
		goto L610
	}
L147:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v644
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	F_MemoryContextReset(m, v624)
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L149
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v644
	F_ProcessWalSummarizerInterrupts(m)
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L150
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v644
	v661 = F_GetRedoRecPtr(m)
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L151
	}
L151:
	;
	v664 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[42]))
	if v664 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v644
	v940 = F_GetLatestLSN(m, v623+int32(256))
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L203
	}
L153:
	;
	v668 = *(*int64)(unsafe.Add(mBase, _c_F_WalSummarizerMain[43]))
	if v661 == v668 {
		goto L152
	} else {
		goto L154
	}
L154:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_WalSummarizerMain[43])) = v661
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v644
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	v674 = F_time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v644
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	v678 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[42]))
	v680 = int64(0)
	v682 = F_GetWalSummaries(m, int32(0), v680, v680)
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L155
	}
L155:
	;
	if v682 == int32(0) {
		goto L152
	} else {
		goto L156
	}
L156:
	;
	v690 = v682
	goto L157
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v644
	F_ProcessWalSummarizerInterrupts(m)
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L159
	}
L158:
	;
	goto L152
L159:
	;
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v690)+12))
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v724)))
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v725)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v644
	v729 = F_XLogGetOldestSegno(m, v726)
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L160
	}
L160:
	;
	v733 = int64(*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[44])))
	v734 = v729 * v733
	v735 = v690
	v740 = int32(0)
	goto L161
L161:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v735)+4))
	if v740 < v765 {
		goto L163
	} else {
		goto L164
	}
L162:
	;
	if v899 != 0 {
		v690 = v899
		goto L157
	} else {
		goto L202
	}
L163:
	;
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v735)+12))
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v767+v740<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v644
	F_ProcessWalSummarizerInterrupts(m)
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L166
	}
L164:
	;
	v899 = v735
	goto L165
L165:
	;
	goto L162
L166:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v771)+16))
	if v776 != v726 {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	if v893 != 0 {
		v735 = v893
		v740 = v895
		goto L161
	} else {
		goto L201
	}
L168:
	;
	v893 = v735
	v895 = v740 + int32(1)
	goto L167
L169:
	;
	goto L170
L170:
	;
	if v734 != int64(0) {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v644
	v887 = F_list_delete_nth_cell(m, v735, v740)
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L199
	}
L172:
	;
	v782 = *(*int64)(unsafe.Add(mBase, uint32(v771)+8))
	if base.Ui64(v734) < base.Ui64(v782) {
		goto L171
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v644
	v786 = m.G0
	v788 = v786 - int32(1200)
	m.G0 = v788
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v771)+16))
	v791 = *(*int64)(unsafe.Add(mBase, uint32(v771)))
	v794 = *(*int64)(unsafe.Add(mBase, uint32(v771)+8))
	*(*uint32)(unsafe.Add(mBase, uint32(v788-int32(-64)))) = uint32(v794)
	*(*uint32)(unsafe.Add(mBase, uint32(v788)+56)) = uint32(v791)
	*(*int32)(unsafe.Add(mBase, uint32(v788)+48)) = v790
	v798 = int64(32)
	v799 = int64(base.Ui64(v794) >> (uint(v798) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v788)+60)) = uint32(v799)
	v802 = int64(base.Ui64(v791) >> (uint(v798) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v788)+52)) = uint32(v802)
	v805 = v788 + int32(176)
	v810 = F_pg_snprintf(m, v805, int32(1024), int32(_a_F_WalSummarizerMain_11), v788+int32(48))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L176
	}
L175:
	;
	goto L174
L176:
	;
	v816 = F___fstatat(m, int32(-100), v805, v788+int32(80), int32(256))
	mBase = m.M
	goto L180
L177:
	;
	goto L171
L178:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L195
	}
L179:
	;
	m.G0 = v788 + int32(1200)
	goto L177
L180:
	;
	if v816 != 0 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v818 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[45]))
	if v818 == int32(44) {
		goto L179
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	v838 = *(*int64)(unsafe.Add(mBase, uint32(v788)+136))
	if v674-base.I64_extend_i32_s(v678*int32(60)) <= v838 {
		goto L179
	} else {
		goto L189
	}
L184:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L185
	}
L185:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L186
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v788)+32)) = v805
	F_errmsg(m, int32(_a_F_WalSummarizerMain_12), v788+int32(32))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L187
	}
L187:
	;
	F_errfinish(m, int32(_a_F_WalSummarizerMain_13), int32(247), int32(_a_F_WalSummarizerMain_14))
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L188
	}
L188:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L189:
	;
	v841 = v788 + int32(176)
	v842 = F_unlink(m, v841)
	mBase = m.M
	if v842 != 0 {
		goto L178
	} else {
		goto L190
	}
L190:
	;
	v845 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L191
	}
L191:
	;
	if v845 == int32(0) {
		goto L179
	} else {
		goto L192
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v788))) = v841
	F_errmsg_internal(m, int32(_a_F_WalSummarizerMain_15), v788)
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L193
	}
L193:
	;
	F_errfinish(m, int32(_a_F_WalSummarizerMain_13), int32(256), int32(_a_F_WalSummarizerMain_14))
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L194
	}
L194:
	;
	goto L179
L195:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L196
	}
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v788)+16)) = v788 + int32(176)
	F_errmsg(m, int32(_a_F_WalSummarizerMain_16), v788+int32(16))
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L197
	}
L197:
	;
	F_errfinish(m, int32(_a_F_WalSummarizerMain_13), int32(254), int32(_a_F_WalSummarizerMain_14))
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L198
	}
L198:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v644
	F_pfree(m, v771)
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L200
	}
L200:
	;
	v893 = v887
	v895 = v740
	goto L167
L201:
	;
	v899 = v893
	goto L165
L202:
	;
	goto L158
L203:
	;
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v623)+268))
	if v648 != int64(0) {
		v987 = v648
		goto L204
	} else {
		goto L205
	}
L204:
	;
	if base.Ui64(v987-int64(1)) < base.Ui64(v647) {
		goto L214
	} else {
		goto L215
	}
L205:
	;
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v623)+256))
	if v942 == v945 {
		v987 = v648
		goto L204
	} else {
		goto L206
	}
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v644
	v949 = F_readTimeLineHistory(m, v945)
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L207
	}
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v644
	v955 = F_tliSwitchPoint(m, v942, v949, v623+int32(260))
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L208
	}
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v644
	v961 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L209
	}
L209:
	;
	if v961 == int32(0) {
		v987 = v955
		goto L204
	} else {
		goto L210
	}
L210:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v644
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v623)+260))
	*(*int32)(unsafe.Add(mBase, uint32(v623)+228)) = v967
	*(*int32)(unsafe.Add(mBase, uint32(v623)+224)) = v942
	*(*uint32)(unsafe.Add(mBase, uint32(v623)+236)) = uint32(v955)
	v972 = int64(base.Ui64(v955) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v623)+232)) = uint32(v972)
	F_errmsg_internal(m, int32(_a_F_WalSummarizerMain_17), v623+int32(224))
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L211
	}
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v644
	F_errfinish(m, int32(_a_F_WalSummarizerMain_1), int32(389), int32(_a_F_WalSummarizerMain_2))
	mBase = m.M
	v985 = m.ExcPending
	if v985 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L212
	}
L212:
	;
	v987 = v955
	goto L204
L213:
	;
	v644 = v3686
	v647 = v3690
	v648 = v3687
	goto L147
L214:
	;
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v623)+260))
	*(*int32)(unsafe.Add(mBase, uint32(v623)+268)) = v991
	v993 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v623)+260)) = v993
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v644
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	v998 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[29]))
	v1002 = F_LWLockAcquire(m, v998+int32(_a_F_WalSummarizerMain_4), v993)
	mBase = m.M
	v1003 = m.ExcPending
	if v1003 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L217
	}
L215:
	;
	goto L216
L216:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v644
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	v1022 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v623)+267)))
	v1023 = F_CreateEmptyBlockRefTable(m)
	mBase = m.M
	v1024 = m.ExcPending
	if v1024 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L219
	}
L217:
	;
	v1005 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[30]))
	*(*int64)(unsafe.Add(mBase, uint32(v1005)+24)) = v987
	v1007 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1005)+16)) = uint8(v1007)
	*(*int32)(unsafe.Add(mBase, uint32(v1005)+4)) = v991
	*(*int64)(unsafe.Add(mBase, uint32(v1005)+8)) = v987
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v644
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	v1015 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[29]))
	F_LWLockRelease(m, v1015+int32(_a_F_WalSummarizerMain_4))
	mBase = m.M
	v1019 = m.ExcPending
	if v1019 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L218
	}
L218:
	;
	v3686 = v644
	v3687 = int64(0)
	v3690 = v987
	goto L213
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v644
	v1028 = F_palloc0(m, int32(24))
	mBase = m.M
	v1029 = m.ExcPending
	if v1029 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L220
	}
L220:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1028)+8)) = v940
	*(*int32)(unsafe.Add(mBase, uint32(v1028))) = v942
	*(*uint8)(unsafe.Add(mBase, uint32(v1028)+4)) = uint8(base.B2i32(v987 != int64(0)))
	*(*int32)(unsafe.Add(mBase, uint32(v623)+460)) = int32(394)
	*(*int32)(unsafe.Add(mBase, uint32(v623)+456)) = int32(395)
	*(*int32)(unsafe.Add(mBase, uint32(v623)+452)) = int32(966)
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v644
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	v1044 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[44]))
	v1047 = F_XLogReaderAllocate(m, v1044, v623+int32(452), v1028)
	mBase = m.M
	v1048 = m.ExcPending
	if v1048 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L221
	}
L221:
	;
	if v1047 == int32(0) {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v644
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L225
	}
L223:
	;
	goto L224
L224:
	;
	if v1022&int32(1) != 0 {
		goto L239
	} else {
		goto L240
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v644
	F_errcode(m, int32(_a_F_WalSummarizerMain_18))
	mBase = m.M
	v1061 = m.ExcPending
	if v1061 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L226
	}
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v644
	F_errmsg(m, int32(_a_F_WalSummarizerMain_19), int32(0))
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L227
	}
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v644
	F_errdetail(m, int32(_a_F_WalSummarizerMain_20), int32(0))
	mBase = m.M
	v1073 = m.ExcPending
	if v1073 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L228
	}
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v644
	F_errfinish(m, int32(_a_F_WalSummarizerMain_1), int32(939), int32(_a_F_WalSummarizerMain_21))
	mBase = m.M
	v1080 = m.ExcPending
	if v1080 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L229
	}
L229:
	;
	goto L3
L230:
	;
	if (v2411^int32(-1)|v2449)&int32(1) != 0 {
		goto L601
	} else {
		goto L602
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	v2574 = int32(0)
	v2575 = m.G0
	v2577 = v2575 - int32(_a_F_WalSummarizerMain_22)
	m.G0 = v2577
	*(*int32)(unsafe.Add(mBase, uint32(v2577)+8)) = int32(1697321851)
	base.MemoryFill(m, v2577+int32(24), v2574, int32(_a_F_WalSummarizerMain_23))
	*(*int32)(unsafe.Add(mBase, uint32(v2577)+16)) = v623 + int32(464)
	*(*int32)(unsafe.Add(mBase, uint32(v2577)+12)) = int32(967)
	v2591 = int32(-1)
	v2595 = int32(4)
	v2596 = m.Env.Pgmem_crc32c(m, v2591, v2577+int32(8), v2595)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v2577)+uint32(_c_F_WalSummarizerMain[46]))) = v2596
	v2598 = *(*int32)(unsafe.Add(mBase, uint32(v1023)))
	*(*int32)(unsafe.Add(mBase, uint32(v2577)+uint32(_c_F_WalSummarizerMain[47]))) = v2595
	v2601 = *(*int32)(unsafe.Add(mBase, uint32(v2577)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2577)+20)) = v2601
	v2603 = *(*int32)(unsafe.Add(mBase, uint32(v2598)+8))
	if v2603 == v2574 {
		goto L472
	} else {
		goto L473
	}
L232:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	v2522 = *(*int32)(unsafe.Add(mBase, uint32(v623)+448))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2526 = m.ExcPending
	if v2526 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L463
	}
L233:
	;
	v2440 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	F_pfree(m, v2440)
	mBase = m.M
	v2444 = m.ExcPending
	if v2444 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L452
	}
L234:
	;
	v2411 = v2380
	v2435 = v2405
	goto L233
L235:
	;
	v2343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1028)+16)))
	if v2343 != int32(1) {
		goto L232
	} else {
		goto L445
	}
L236:
	;
	v1560 = int32(1)
	goto L303
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1446
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1539 = m.ExcPending
	if v1539 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L300
	}
L238:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	F_ProcessWalSummarizerInterrupts(m)
	mBase = m.M
	v1526 = m.ExcPending
	if v1526 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L297
	}
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v644
	F_XLogBeginRead(m, v1047, v647)
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L242
	}
L240:
	;
	goto L241
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v644
	v1089 = m.G0
	v1091 = v1089 - int32(16)
	m.G0 = v1091
	v1093 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1047)+1257)) = uint8(v1093)
	v1096 = v647 & int64(-8192)
	v1100 = F_ReadPageInternal(m, v1047, v1096, base.I32_wrap_i64(v647)&int32(_a_F_WalSummarizerMain_24))
	mBase = m.M
	v1101 = m.ExcPending
	if v1101 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L245
	}
L242:
	;
	v1515 = v644
	v1518 = v647
	v1522 = v987
	goto L238
L243:
	;
	m.G0 = v1091 + int32(16)
	if v1446 != int64(0) {
		goto L287
	} else {
		goto L288
	}
L244:
	;
	v1415 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1047)+1192)) = v1415
	v1417 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1047)+1176)) = v1417
	*(*int32)(unsafe.Add(mBase, uint32(v1047)+132)) = v1415
	v1446 = v1417
	goto L243
L245:
	;
	if v1100 < int32(0) {
		goto L244
	} else {
		goto L246
	}
L246:
	;
	v1128 = v1096
	goto L247
L247:
	;
	v1136 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+128))
	v1137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1136)+2)))
	if v1137&int32(2) != 0 {
		goto L249
	} else {
		goto L250
	}
L248:
	;
	goto L244
L249:
	;
	v1140 = int32(40)
	goto L251
L250:
	;
	v1140 = int32(24)
	goto L251
L251:
	;
	v1141 = F_ReadPageInternal(m, v1047, v1128, v1140)
	mBase = m.M
	v1142 = m.ExcPending
	if v1142 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L252
	}
L252:
	;
	if v1141 < int32(0) {
		goto L244
	} else {
		goto L253
	}
L253:
	;
	v1145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1136)+2)))
	if v1145&int32(1) != 0 {
		goto L256
	} else {
		goto L257
	}
L254:
	;
	v1379 = v1128 - int64(-8192)
	v1381 = F_ReadPageInternal(m, v1047, v1379, int32(0))
	mBase = m.M
	v1382 = m.ExcPending
	if v1382 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L285
	}
L255:
	;
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+120))
	if v1164 != 0 {
		goto L260
	} else {
		goto L261
	}
L256:
	;
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(v1136)+16))
	v1152 = (v1148 + int32(7)) & int32(-8)
	if base.Ui32(int32(_a_F_WalSummarizerMain_6)-v1140) <= base.Ui32(v1152) {
		goto L254
	} else {
		goto L259
	}
L257:
	;
	goto L258
L258:
	;
	v1163 = v1128 | base.I64_extend_i32_u(v1140)
	goto L255
L259:
	;
	v1163 = base.I64_extend_i32_u(v1152) + (v1128 | base.I64_extend_i32_u(v1140))
	goto L255
L260:
	;
	v1165 = v1164
	goto L263
L261:
	;
	goto L262
L262:
	;
	v1234 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1047)+120)) = v1234
	v1236 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1047)+96)) = v1236
	v1238 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v1047)+116)) = v1238
	*(*int32)(unsafe.Add(mBase, uint32(v1047)+112)) = v1238
	v1241 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+1252))
	*(*uint8)(unsafe.Add(mBase, uint32(v1241))) = uint8(v1236)
	*(*int64)(unsafe.Add(mBase, uint32(v1047)+80)) = v1163
	*(*int64)(unsafe.Add(mBase, uint32(v1047)+40)) = v1163
	*(*uint8)(unsafe.Add(mBase, uint32(v1047)+1256)) = uint8(v1236)
	*(*int64)(unsafe.Add(mBase, uint32(v1047)+72)) = v1234
	*(*int64)(unsafe.Add(mBase, uint32(v1047)+32)) = v1234
	goto L270
L263:
	;
	v1195 = *(*int32)(unsafe.Add(mBase, uint32(v1165)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1047)+120)) = v1195
	v1197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1165)+4)))
	if v1197 == int32(1) {
		goto L265
	} else {
		goto L266
	}
L264:
	;
	goto L262
L265:
	;
	F_pfree(m, v1165)
	mBase = m.M
	v1201 = m.ExcPending
	if v1201 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L268
	}
L266:
	;
	v1203 = v1195
	goto L267
L267:
	;
	if v1203 != 0 {
		v1165 = v1203
		goto L263
	} else {
		goto L269
	}
L268:
	;
	v1202 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+120))
	v1203 = v1202
	goto L267
L269:
	;
	goto L264
L270:
	;
	v1284 = F_XLogReadRecord(m, v1047, v1091+int32(12))
	mBase = m.M
	v1285 = m.ExcPending
	if v1285 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L272
	}
L271:
	;
	v1290 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+120))
	if v1290 != 0 {
		goto L275
	} else {
		goto L276
	}
L272:
	;
	if v1284 == int32(0) {
		goto L244
	} else {
		goto L273
	}
L273:
	;
	v1288 = *(*int64)(unsafe.Add(mBase, uint32(v1047)+32))
	if base.Ui64(v1288) < base.Ui64(v647) {
		goto L270
	} else {
		goto L274
	}
L274:
	;
	goto L271
L275:
	;
	v1291 = v1290
	goto L278
L276:
	;
	goto L277
L277:
	;
	v1360 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1047)+120)) = v1360
	v1362 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1047)+96)) = v1362
	v1364 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v1047)+116)) = v1364
	*(*int32)(unsafe.Add(mBase, uint32(v1047)+112)) = v1364
	v1367 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+1252))
	*(*uint8)(unsafe.Add(mBase, uint32(v1367))) = uint8(v1362)
	*(*int64)(unsafe.Add(mBase, uint32(v1047)+80)) = v1288
	*(*int64)(unsafe.Add(mBase, uint32(v1047)+40)) = v1288
	*(*uint8)(unsafe.Add(mBase, uint32(v1047)+1256)) = uint8(v1362)
	*(*int64)(unsafe.Add(mBase, uint32(v1047)+72)) = v1360
	*(*int64)(unsafe.Add(mBase, uint32(v1047)+32)) = v1360
	v1446 = v1288
	goto L243
L278:
	;
	v1321 = *(*int32)(unsafe.Add(mBase, uint32(v1291)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1047)+120)) = v1321
	v1323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1291)+4)))
	if v1323 == int32(1) {
		goto L280
	} else {
		goto L281
	}
L279:
	;
	goto L277
L280:
	;
	F_pfree(m, v1291)
	mBase = m.M
	v1327 = m.ExcPending
	if v1327 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L283
	}
L281:
	;
	v1329 = v1321
	goto L282
L282:
	;
	if v1329 != 0 {
		v1291 = v1329
		goto L278
	} else {
		goto L284
	}
L283:
	;
	v1328 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+120))
	v1329 = v1328
	goto L282
L284:
	;
	goto L279
L285:
	;
	if int32(0) <= v1381 {
		v1128 = v1379
		goto L247
	} else {
		goto L286
	}
L286:
	;
	goto L248
L287:
	;
	v1515 = v1446
	v1518 = v1446
	v1522 = v987
	goto L238
L288:
	;
	goto L289
L289:
	;
	v1457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1028)+16)))
	if v1457 != int32(1) {
		goto L237
	} else {
		goto L290
	}
L290:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1446
	v1464 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1465 = m.ExcPending
	if v1465 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L291
	}
L291:
	;
	if v1464 != 0 {
		goto L292
	} else {
		goto L293
	}
L292:
	;
	v1466 = *(*int64)(unsafe.Add(mBase, uint32(v1028)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1446
	*(*uint32)(unsafe.Add(mBase, uint32(v623)+192)) = uint32(v1466)
	v1470 = int64(32)
	v1471 = int64(base.Ui64(v1466) >> (uint(v1470) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v623)+188)) = uint32(v1471)
	*(*int32)(unsafe.Add(mBase, uint32(v623)+176)) = v942
	*(*uint32)(unsafe.Add(mBase, uint32(v623)+184)) = uint32(v647)
	v1476 = int64(base.Ui64(v647) >> (uint(v1470) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v623)+180)) = uint32(v1476)
	F_errmsg_internal(m, int32(_a_F_WalSummarizerMain_25), v623+int32(176))
	mBase = m.M
	v1482 = m.ExcPending
	if v1482 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L295
	}
L293:
	;
	goto L294
L294:
	;
	v1491 = *(*int64)(unsafe.Add(mBase, uint32(v1047)+40))
	v1515 = v1446
	v1518 = v647
	v1522 = v1491
	goto L238
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1446
	F_errfinish(m, int32(_a_F_WalSummarizerMain_1), int32(987), int32(_a_F_WalSummarizerMain_21))
	mBase = m.M
	v1489 = m.ExcPending
	if v1489 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L296
	}
L296:
	;
	goto L294
L297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	v1531 = F_XLogReadRecord(m, v1047, v623+int32(448))
	mBase = m.M
	v1532 = m.ExcPending
	if v1532 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L298
	}
L298:
	;
	if v1531 != 0 {
		goto L236
	} else {
		goto L299
	}
L299:
	;
	v2313 = int32(1)
	goto L235
L300:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1446
	*(*uint32)(unsafe.Add(mBase, uint32(v623)+212)) = uint32(v647)
	v1544 = int64(base.Ui64(v647) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v623)+208)) = uint32(v1544)
	F_errmsg(m, int32(_a_F_WalSummarizerMain_26), v623+int32(208))
	mBase = m.M
	v1550 = m.ExcPending
	if v1550 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L301
	}
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1446
	F_errfinish(m, int32(_a_F_WalSummarizerMain_1), int32(1004), int32(_a_F_WalSummarizerMain_21))
	mBase = m.M
	v1557 = m.ExcPending
	if v1557 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L302
	}
L302:
	;
	goto L3
L303:
	;
	v1590 = base.B2i32(v1522 == int64(0))
	if v1522 == int64(0) {
		goto L305
	} else {
		goto L306
	}
L304:
	;
	v2313 = v2248
	goto L235
L305:
	;
	v1593 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+96))
	v1594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1593)+49)))
	if v1594 == int32(0) {
		goto L310
	} else {
		goto L311
	}
L306:
	;
	v1591 = *(*int64)(unsafe.Add(mBase, uint32(v1047)+32))
	if base.Ui64(v1591) < base.Ui64(v1522) {
		goto L305
	} else {
		goto L307
	}
L307:
	;
	v2411 = v1560
	v2435 = v1522
	goto L233
L308:
	;
	v2278 = *(*int64)(unsafe.Add(mBase, uint32(v1047)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	v2282 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[29]))
	v2286 = F_LWLockAcquire(m, v2282+int32(_a_F_WalSummarizerMain_4), int32(0))
	mBase = m.M
	v2287 = m.ExcPending
	if v2287 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L436
	}
L309:
	;
	v2158 = int32(0)
	v2159 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+96))
	v2160 = *(*int32)(unsafe.Add(mBase, uint32(v2159)+72))
	if v2160 < v2158 {
		v2248 = v2158
		goto L308
	} else {
		goto L414
	}
L310:
	;
	v1597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1593)+48)))
	v1599 = v1597 & int32(240)
	if base.Ui32(v1599) <= base.Ui32(int32(143)) {
		goto L317
	} else {
		goto L318
	}
L311:
	;
	goto L312
L312:
	;
	v1630 = int32(1)
	if v1560&v1630 != 0 {
		v2248 = v1630
		goto L308
	} else {
		goto L327
	}
L313:
	;
	v1625 = int32(1)
	if v1560&v1625 == int32(0) {
		goto L309
	} else {
		goto L326
	}
L314:
	;
	v1621 = *(*int64)(unsafe.Add(mBase, uint32(v1047)+32))
	if base.Ui64(v1518) < base.Ui64(v1621) {
		v2411 = v1560
		v2435 = v1621
		goto L233
	} else {
		goto L324
	}
L315:
	;
	v1617 = *(*int32)(unsafe.Add(mBase, uint32(v1593)+64))
	v1620 = v1617 + int32(16)
	goto L314
L316:
	;
	v1614 = *(*int32)(unsafe.Add(mBase, uint32(v1593)+64))
	v1620 = v1614 + int32(20)
	goto L314
L317:
	;
	if v1599 == int32(0) {
		goto L316
	} else {
		goto L320
	}
L318:
	;
	goto L319
L319:
	;
	if v1599 == int32(144) {
		goto L315
	} else {
		goto L322
	}
L320:
	;
	if v1599 != int32(96) {
		goto L313
	} else {
		goto L321
	}
L321:
	;
	v1606 = *(*int32)(unsafe.Add(mBase, uint32(v1593)+64))
	v1620 = v1606 + int32(20)
	goto L314
L322:
	;
	if v1599 != int32(224) {
		goto L313
	} else {
		goto L323
	}
L323:
	;
	v1613 = *(*int32)(unsafe.Add(mBase, uint32(v1593)+64))
	v1620 = v1613
	goto L314
L324:
	;
	v1623 = *(*int32)(unsafe.Add(mBase, uint32(v1620)))
	if v1623 != 0 {
		goto L309
	} else {
		goto L325
	}
L325:
	;
	v2248 = int32(1)
	goto L308
L326:
	;
	v2248 = v1625
	goto L308
L327:
	;
	switch v1594 - int32(1) {
	case 0:
		goto L328
	case 1:
		goto L329
	default:
		goto L309
	case 3:
		goto L330
	}
L328:
	;
	v1771 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1593)+48)))
	switch int32(base.Ui32(v1771)>>(uint(int32(4))%32)) & int32(7) {
	case 0, 3:
		goto L358
	default:
		goto L309
	case 2, 4:
		goto L357
	}
L329:
	;
	v1728 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1593)+48)))
	v1732 = v1728&int32(240) - int32(16)
	if v1732 != 0 {
		goto L343
	} else {
		goto L344
	}
L330:
	;
	v1635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1593)+48)))
	v1637 = v1635 & int32(240)
	switch v1637 - int32(16) {
	case 0:
		goto L332
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
		goto L309
	case 16:
		goto L331
	default:
		goto L333
	}
L331:
	;
	v1670 = *(*int32)(unsafe.Add(mBase, uint32(v1593)+64))
	v1671 = *(*int32)(unsafe.Add(mBase, uint32(v1670)))
	v1672 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v623)+2540)) = v1672
	*(*int32)(unsafe.Add(mBase, uint32(v623)+2536)) = v1671
	v1675 = *(*int32)(unsafe.Add(mBase, uint32(v1670)+4))
	if v1675 <= v1672 {
		goto L309
	} else {
		goto L337
	}
L332:
	;
	v1655 = *(*int32)(unsafe.Add(mBase, uint32(v1593)+64))
	v1656 = *(*int32)(unsafe.Add(mBase, uint32(v1655)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v623)+2544)) = v1656
	v1658 = *(*int32)(unsafe.Add(mBase, uint32(v1655)))
	v1659 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v623)+2552)) = v1659
	*(*int32)(unsafe.Add(mBase, uint32(v623)+2548)) = v1658
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	F_BlockRefTableSetLimitBlock(m, v1023, v623+int32(2544), v1659, v1659)
	mBase = m.M
	v1669 = m.ExcPending
	if v1669 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L336
	}
L333:
	;
	if v1637 != 0 {
		goto L309
	} else {
		goto L334
	}
L334:
	;
	v1640 = *(*int32)(unsafe.Add(mBase, uint32(v1593)+64))
	v1641 = *(*int32)(unsafe.Add(mBase, uint32(v1640)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v623)+2556)) = v1641
	v1643 = *(*int32)(unsafe.Add(mBase, uint32(v1640)))
	v1644 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v623)+2564)) = v1644
	*(*int32)(unsafe.Add(mBase, uint32(v623)+2560)) = v1643
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	F_BlockRefTableSetLimitBlock(m, v1023, v623+int32(2556), v1644, v1644)
	mBase = m.M
	v1654 = m.ExcPending
	if v1654 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L335
	}
L335:
	;
	goto L309
L336:
	;
	goto L309
L337:
	;
	v1681 = int32(0)
	goto L338
L338:
	;
	v1714 = *(*int32)(unsafe.Add(mBase, uint32(v1670+int32(8)+v1681<<(uint(int32(2))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	*(*int32)(unsafe.Add(mBase, uint32(v623)+2532)) = v1714
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	v1720 = int32(0)
	F_BlockRefTableSetLimitBlock(m, v1023, v623+int32(2532), v1720, v1720)
	mBase = m.M
	v1723 = m.ExcPending
	if v1723 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L340
	}
L339:
	;
	goto L309
L340:
	;
	v1725 = v1681 + int32(1)
	v1726 = *(*int32)(unsafe.Add(mBase, uint32(v1670)+4))
	if v1725 < v1726 {
		v1681 = v1725
		goto L338
	} else {
		goto L341
	}
L341:
	;
	goto L339
L342:
	;
	v1744 = *(*int32)(unsafe.Add(mBase, uint32(v1593)+64))
	v1745 = *(*int32)(unsafe.Add(mBase, uint32(v1744)+16))
	if v1745&int32(1) != 0 {
		goto L351
	} else {
		goto L352
	}
L343:
	;
	if v1732 == int32(16) {
		goto L346
	} else {
		goto L347
	}
L344:
	;
	goto L345
L345:
	;
	v1735 = *(*int32)(unsafe.Add(mBase, uint32(v1593)+64))
	v1736 = *(*int32)(unsafe.Add(mBase, uint32(v1735)+12))
	if v1736 == int32(1) {
		goto L309
	} else {
		goto L349
	}
L346:
	;
	goto L342
L347:
	;
	goto L309
L349:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	F_BlockRefTableSetLimitBlock(m, v1023, v1735, v1736, int32(0))
	mBase = m.M
	v1743 = m.ExcPending
	if v1743 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L350
	}
L350:
	;
	goto L309
L351:
	;
	v1748 = *(*int32)(unsafe.Add(mBase, uint32(v1744)))
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	F_BlockRefTableSetLimitBlock(m, v1023, v1744+int32(4), int32(0), v1748)
	mBase = m.M
	v1755 = m.ExcPending
	if v1755 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L354
	}
L352:
	;
	v1758 = v1745
	goto L353
L353:
	;
	if v1758&int32(2) == int32(0) {
		goto L309
	} else {
		goto L355
	}
L354:
	;
	v1756 = *(*int32)(unsafe.Add(mBase, uint32(v1744)+16))
	v1758 = v1756
	goto L353
L355:
	;
	v1763 = *(*int32)(unsafe.Add(mBase, uint32(v1744)))
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	F_BlockRefTableSetLimitBlock(m, v1023, v1744+int32(4), int32(2), v1763)
	mBase = m.M
	v1770 = m.ExcPending
	if v1770 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L356
	}
L356:
	;
	goto L309
L357:
	;
	v1959 = *(*int32)(unsafe.Add(mBase, uint32(v1593)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	v1963 = v623 + int32(2568)
	v1964 = int32(0)
	base.MemoryFill(m, v1963, v1964, int32(264))
	v1970 = *(*int64)(unsafe.Add(mBase, uint32(v1959)))
	*(*int64)(unsafe.Add(mBase, uint32(v1963))) = v1970
	if v1964 <= base.I32_extend8_s(v1771) {
		goto L389
	} else {
		goto L390
	}
L358:
	;
	v1776 = *(*int32)(unsafe.Add(mBase, uint32(v1593)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	v1780 = v623 + int32(2832)
	v1781 = int32(0)
	base.MemoryFill(m, v1780, v1781, int32(288))
	v1787 = *(*int64)(unsafe.Add(mBase, uint32(v1776)))
	*(*int64)(unsafe.Add(mBase, uint32(v1780))) = v1787
	if v1781 <= base.I32_extend8_s(v1771) {
		goto L360
	} else {
		goto L361
	}
L359:
	;
	v1895 = int32(0)
	v1896 = *(*int32)(unsafe.Add(mBase, uint32(v623)+2860))
	if v1896 <= v1895 {
		goto L309
	} else {
		goto L381
	}
L360:
	;
	goto L359
L361:
	;
	v1792 = *(*int32)(unsafe.Add(mBase, uint32(v1776)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1780)+8)) = v1792
	if v1792&int32(1) != 0 {
		goto L362
	} else {
		goto L363
	}
L362:
	;
	v1796 = *(*int32)(unsafe.Add(mBase, uint32(v1776)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1780)+12)) = v1796
	v1798 = *(*int32)(unsafe.Add(mBase, uint32(v1776)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1780)+16)) = v1798
	v1804 = v1776 + int32(20)
	goto L364
L363:
	;
	v1804 = v1776 + int32(12)
	goto L364
L364:
	;
	if v1792&int32(2) != 0 {
		goto L365
	} else {
		goto L366
	}
L365:
	;
	v1807 = *(*int32)(unsafe.Add(mBase, uint32(v1804)))
	v1809 = v1804 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1780)+24)) = v1809
	*(*int32)(unsafe.Add(mBase, uint32(v1780)+20)) = v1807
	v1815 = v1809 + v1807<<(uint(int32(2))%32)
	goto L367
L366:
	;
	v1815 = v1804
	goto L367
L367:
	;
	if v1792&int32(4) != 0 {
		goto L368
	} else {
		goto L369
	}
L368:
	;
	v1819 = *(*int32)(unsafe.Add(mBase, uint32(v1815)))
	v1821 = v1815 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1780)+32)) = v1821
	*(*int32)(unsafe.Add(mBase, uint32(v1780)+28)) = v1819
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(v1815)))
	v1828 = v1821 + v1824*int32(12)
	goto L370
L369:
	;
	v1828 = v1815
	goto L370
L370:
	;
	if v1792&int32(256) != 0 {
		goto L371
	} else {
		goto L372
	}
L371:
	;
	v1833 = *(*int32)(unsafe.Add(mBase, uint32(v1828)))
	v1834 = int32(4)
	v1835 = v1828 + v1834
	*(*int32)(unsafe.Add(mBase, uint32(v1780)+40)) = v1835
	*(*int32)(unsafe.Add(mBase, uint32(v1780)+36)) = v1833
	v1838 = *(*int32)(unsafe.Add(mBase, uint32(v1828)))
	v1842 = v1835 + v1838<<(uint(v1834)%32)
	goto L373
L372:
	;
	v1842 = v1828
	goto L373
L373:
	;
	if v1792&int32(8) != 0 {
		goto L374
	} else {
		goto L375
	}
L374:
	;
	v1847 = *(*int32)(unsafe.Add(mBase, uint32(v1842)))
	v1848 = int32(4)
	v1849 = v1842 + v1848
	*(*int32)(unsafe.Add(mBase, uint32(v1780)+48)) = v1849
	*(*int32)(unsafe.Add(mBase, uint32(v1780)+44)) = v1847
	v1852 = *(*int32)(unsafe.Add(mBase, uint32(v1842)))
	v1856 = v1849 + v1852<<(uint(v1848)%32)
	goto L376
L375:
	;
	v1856 = v1842
	goto L376
L376:
	;
	if v1792&int32(16) == int32(0) {
		v1880 = v1792
		v1881 = v1856
		goto L377
	} else {
		goto L378
	}
L377:
	;
	if v1880&int32(32) == int32(0) {
		goto L360
	} else {
		goto L380
	}
L378:
	;
	v1863 = *(*int32)(unsafe.Add(mBase, uint32(v1856)))
	*(*int32)(unsafe.Add(mBase, uint32(v1780)+52)) = v1863
	v1866 = v1856 + int32(4)
	if v1792&int32(128) == int32(0) {
		v1880 = v1792
		v1881 = v1866
		goto L377
	} else {
		goto L379
	}
L379:
	;
	v1874 = F_strlcpy(m, v623+int32(2888), v1866, int32(200))
	mBase = m.M
	v1875 = F_strlen(m, v1866)
	mBase = m.M
	v1879 = *(*int32)(unsafe.Add(mBase, uint32(v1780)+8))
	v1880 = v1879
	v1881 = v1875 + v1866 + int32(1)
	goto L377
L380:
	;
	v1886 = *(*int64)(unsafe.Add(mBase, uint32(v1881)))
	v1887 = *(*int64)(unsafe.Add(mBase, uint32(v1881)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1780)+280)) = v1887
	*(*int64)(unsafe.Add(mBase, uint32(v1780)+272)) = v1886
	goto L360
L381:
	;
	v1899 = v1895
	goto L382
L382:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	v1932 = v1899 * int32(12)
	v1933 = *(*int32)(unsafe.Add(mBase, uint32(v623)+2864))
	v1935 = int32(0)
	F_BlockRefTableSetLimitBlock(m, v1023, v1932+v1933, v1935, v1935)
	mBase = m.M
	v1938 = m.ExcPending
	if v1938 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L384
	}
L383:
	;
	goto L309
L384:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	v1941 = *(*int32)(unsafe.Add(mBase, uint32(v623)+2864))
	F_BlockRefTableSetLimitBlock(m, v1023, v1941+v1932, int32(2), int32(0))
	mBase = m.M
	v1946 = m.ExcPending
	if v1946 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L385
	}
L385:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	v1949 = *(*int32)(unsafe.Add(mBase, uint32(v623)+2864))
	F_BlockRefTableSetLimitBlock(m, v1023, v1949+v1932, int32(3), int32(0))
	mBase = m.M
	v1954 = m.ExcPending
	if v1954 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L386
	}
L386:
	;
	v1956 = v1899 + int32(1)
	v1957 = *(*int32)(unsafe.Add(mBase, uint32(v623)+2860))
	if v1956 < v1957 {
		v1899 = v1956
		goto L382
	} else {
		goto L387
	}
L387:
	;
	goto L383
L388:
	;
	v2064 = int32(0)
	v2065 = *(*int32)(unsafe.Add(mBase, uint32(v623)+2596))
	if v2065 <= v2064 {
		goto L309
	} else {
		goto L407
	}
L389:
	;
	goto L388
L390:
	;
	v1975 = *(*int32)(unsafe.Add(mBase, uint32(v1959)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1963)+8)) = v1975
	if v1975&int32(1) != 0 {
		goto L391
	} else {
		goto L392
	}
L391:
	;
	v1979 = *(*int32)(unsafe.Add(mBase, uint32(v1959)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1963)+12)) = v1979
	v1981 = *(*int32)(unsafe.Add(mBase, uint32(v1959)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1963)+16)) = v1981
	v1987 = v1959 + int32(20)
	goto L393
L392:
	;
	v1987 = v1959 + int32(12)
	goto L393
L393:
	;
	if v1975&int32(2) != 0 {
		goto L394
	} else {
		goto L395
	}
L394:
	;
	v1990 = *(*int32)(unsafe.Add(mBase, uint32(v1987)))
	v1992 = v1987 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1963)+24)) = v1992
	*(*int32)(unsafe.Add(mBase, uint32(v1963)+20)) = v1990
	v1998 = v1992 + v1990<<(uint(int32(2))%32)
	goto L396
L395:
	;
	v1998 = v1987
	goto L396
L396:
	;
	if v1975&int32(4) != 0 {
		goto L397
	} else {
		goto L398
	}
L397:
	;
	v2002 = *(*int32)(unsafe.Add(mBase, uint32(v1998)))
	v2004 = v1998 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1963)+32)) = v2004
	*(*int32)(unsafe.Add(mBase, uint32(v1963)+28)) = v2002
	v2007 = *(*int32)(unsafe.Add(mBase, uint32(v1998)))
	v2011 = v2004 + v2007*int32(12)
	goto L399
L398:
	;
	v2011 = v1998
	goto L399
L399:
	;
	if v1975&int32(256) != 0 {
		goto L400
	} else {
		goto L401
	}
L400:
	;
	v2016 = *(*int32)(unsafe.Add(mBase, uint32(v2011)))
	v2017 = int32(4)
	v2018 = v2011 + v2017
	*(*int32)(unsafe.Add(mBase, uint32(v1963)+40)) = v2018
	*(*int32)(unsafe.Add(mBase, uint32(v1963)+36)) = v2016
	v2021 = *(*int32)(unsafe.Add(mBase, uint32(v2011)))
	v2025 = v2018 + v2021<<(uint(v2017)%32)
	goto L402
L401:
	;
	v2025 = v2011
	goto L402
L402:
	;
	if v1975&int32(16) == int32(0) {
		v2049 = v1975
		v2050 = v2025
		goto L403
	} else {
		goto L404
	}
L403:
	;
	if v2049&int32(32) == int32(0) {
		goto L389
	} else {
		goto L406
	}
L404:
	;
	v2032 = *(*int32)(unsafe.Add(mBase, uint32(v2025)))
	*(*int32)(unsafe.Add(mBase, uint32(v1963)+44)) = v2032
	v2035 = v2025 + int32(4)
	if v1975&int32(128) == int32(0) {
		v2049 = v1975
		v2050 = v2035
		goto L403
	} else {
		goto L405
	}
L405:
	;
	v2043 = F_strlcpy(m, v623+int32(2616), v2035, int32(200))
	mBase = m.M
	v2044 = F_strlen(m, v2035)
	mBase = m.M
	v2048 = *(*int32)(unsafe.Add(mBase, uint32(v1963)+8))
	v2049 = v2048
	v2050 = v2044 + v2035 + int32(1)
	goto L403
L406:
	;
	v2055 = *(*int64)(unsafe.Add(mBase, uint32(v2050)))
	v2056 = *(*int64)(unsafe.Add(mBase, uint32(v2050)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1963)+256)) = v2056
	*(*int64)(unsafe.Add(mBase, uint32(v1963)+248)) = v2055
	goto L389
L407:
	;
	v2068 = v2064
	goto L408
L408:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	v2101 = v2068 * int32(12)
	v2102 = *(*int32)(unsafe.Add(mBase, uint32(v623)+2600))
	v2104 = int32(0)
	F_BlockRefTableSetLimitBlock(m, v1023, v2101+v2102, v2104, v2104)
	mBase = m.M
	v2107 = m.ExcPending
	if v2107 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L410
	}
L409:
	;
	goto L309
L410:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	v2110 = *(*int32)(unsafe.Add(mBase, uint32(v623)+2600))
	F_BlockRefTableSetLimitBlock(m, v1023, v2110+v2101, int32(2), int32(0))
	mBase = m.M
	v2115 = m.ExcPending
	if v2115 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L411
	}
L411:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	v2118 = *(*int32)(unsafe.Add(mBase, uint32(v623)+2600))
	F_BlockRefTableSetLimitBlock(m, v1023, v2118+v2101, int32(3), int32(0))
	mBase = m.M
	v2123 = m.ExcPending
	if v2123 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L412
	}
L412:
	;
	v2125 = v2068 + int32(1)
	v2126 = *(*int32)(unsafe.Add(mBase, uint32(v623)+2596))
	if v2125 < v2126 {
		v2068 = v2125
		goto L408
	} else {
		goto L413
	}
L413:
	;
	goto L409
L414:
	;
	v2163 = v2158
	goto L415
L415:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	v2196 = v2163 & int32(255)
	v2198 = v623 + int32(436)
	v2200 = v623 + int32(432)
	v2202 = v623 + int32(428)
	v2203 = int32(0)
	v2205 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+96))
	v2206 = *(*int32)(unsafe.Add(mBase, uint32(v2205)+72))
	if v2206 < v2196 {
		v2230 = v2203
		goto L419
	} else {
		goto L420
	}
L416:
	;
	v2248 = int32(0)
	goto L308
L417:
	;
	v2243 = v2163 + int32(1)
	v2244 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+96))
	v2245 = *(*int32)(unsafe.Add(mBase, uint32(v2244)+72))
	if v2243 <= v2245 {
		v2163 = v2243
		goto L415
	} else {
		goto L435
	}
L418:
	;
	if v2230 == int32(0) {
		goto L417
	} else {
		goto L432
	}
L419:
	;
	goto L418
L420:
	;
	v2210 = v2205 + v2196*int32(52)
	v2211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2210)+76)))
	if v2211 != int32(1) {
		v2230 = v2203
		goto L419
	} else {
		goto L421
	}
L421:
	;
	v2215 = v2210 + int32(76)
	if v2198 != 0 {
		goto L422
	} else {
		goto L423
	}
L422:
	;
	v2216 = *(*int32)(unsafe.Add(mBase, uint32(v2215)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v2198)+8)) = v2216
	v2218 = *(*int64)(unsafe.Add(mBase, uint32(v2215)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v2198))) = v2218
	goto L424
L423:
	;
	goto L424
L424:
	;
	if v2200 != 0 {
		goto L425
	} else {
		goto L426
	}
L425:
	;
	v2220 = *(*int32)(unsafe.Add(mBase, uint32(v2215)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v2200))) = v2220
	goto L427
L426:
	;
	goto L427
L427:
	;
	if v2202 != 0 {
		goto L428
	} else {
		goto L429
	}
L428:
	;
	v2222 = *(*int32)(unsafe.Add(mBase, uint32(v2215)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2202))) = v2222
	goto L430
L429:
	;
	goto L430
L430:
	;
	v2230 = int32(1)
	goto L419
L432:
	;
	v2233 = *(*int32)(unsafe.Add(mBase, uint32(v623)+432))
	if v2233 == int32(1) {
		goto L417
	} else {
		goto L433
	}
L433:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	v2238 = *(*int32)(unsafe.Add(mBase, uint32(v623)+428))
	F_BlockRefTableMarkBlockModified(m, v1023, v2198, v2233, v2238)
	mBase = m.M
	v2240 = m.ExcPending
	if v2240 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L434
	}
L434:
	;
	goto L417
L435:
	;
	goto L416
L436:
	;
	v2289 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[30]))
	*(*int64)(unsafe.Add(mBase, uint32(v2289)+24)) = v2278
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	v2294 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[29]))
	F_LWLockRelease(m, v2294+int32(_a_F_WalSummarizerMain_4))
	mBase = m.M
	v2298 = m.ExcPending
	if v2298 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L437
	}
L437:
	;
	if v1590 == int32(0) {
		goto L438
	} else {
		goto L439
	}
L438:
	;
	v2301 = *(*int64)(unsafe.Add(mBase, uint32(v1047)+40))
	if base.Ui64(v1522) <= base.Ui64(v2301) {
		v2380 = v2248
		v2405 = v2278
		goto L234
	} else {
		goto L441
	}
L439:
	;
	goto L440
L440:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	F_ProcessWalSummarizerInterrupts(m)
	mBase = m.M
	v2306 = m.ExcPending
	if v2306 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L442
	}
L441:
	;
	goto L440
L442:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	v2311 = F_XLogReadRecord(m, v1047, v623+int32(448))
	mBase = m.M
	v2312 = m.ExcPending
	if v2312 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L443
	}
L443:
	;
	if v2311 != 0 {
		v1560 = v2248
		goto L303
	} else {
		goto L444
	}
L444:
	;
	goto L304
L445:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	v2350 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2351 = m.ExcPending
	if v2351 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L446
	}
L446:
	;
	if v2350 != 0 {
		goto L447
	} else {
		goto L448
	}
L447:
	;
	v2352 = *(*int64)(unsafe.Add(mBase, uint32(v1047)+40))
	v2353 = *(*int64)(unsafe.Add(mBase, uint32(v1028)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	*(*uint32)(unsafe.Add(mBase, uint32(v637))) = uint32(v2353)
	v2357 = int64(32)
	v2358 = int64(base.Ui64(v2353) >> (uint(v2357) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v623)+124)) = uint32(v2358)
	*(*uint32)(unsafe.Add(mBase, uint32(v623)+120)) = uint32(v2352)
	v2362 = int64(base.Ui64(v2352) >> (uint(v2357) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v623)+116)) = uint32(v2362)
	*(*int32)(unsafe.Add(mBase, uint32(v623)+112)) = v942
	F_errmsg_internal(m, int32(_a_F_WalSummarizerMain_25), v623+int32(112))
	mBase = m.M
	v2369 = m.ExcPending
	if v2369 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L450
	}
L448:
	;
	goto L449
L449:
	;
	v2379 = *(*int64)(unsafe.Add(mBase, uint32(v1028)+8))
	v2380 = v2313
	v2405 = v2379
	goto L234
L450:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	F_errfinish(m, int32(_a_F_WalSummarizerMain_1), int32(1040), int32(_a_F_WalSummarizerMain_21))
	mBase = m.M
	v2376 = m.ExcPending
	if v2376 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L451
	}
L451:
	;
	goto L449
L452:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	F_XLogReaderFree(m, v1047)
	mBase = m.M
	v2448 = m.ExcPending
	if v2448 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L453
	}
L453:
	;
	v2449 = base.B2i32(base.Ui64(v2435) <= base.Ui64(v1518))
	if (v2411|v2449)&int32(1) != 0 {
		goto L230
	} else {
		goto L454
	}
L454:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	v2456 = v623 + int32(1504)
	v2460 = F_pg_snprintf(m, v2456, int32(1024), int32(_a_F_WalSummarizerMain_27), int32(0))
	mBase = m.M
	v2461 = m.ExcPending
	if v2461 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L455
	}
L455:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	v2464 = base.I32_wrap_i64(v2435)
	*(*int32)(unsafe.Add(mBase, uint32(v623)+96)) = v2464
	v2466 = int64(32)
	v2468 = base.I32_wrap_i64(int64(base.Ui64(v2435) >> (uint(v2466) % 64)))
	*(*int32)(unsafe.Add(mBase, uint32(v623)+92)) = v2468
	v2470 = base.I32_wrap_i64(v1518)
	*(*int32)(unsafe.Add(mBase, uint32(v623)+88)) = v2470
	v2474 = base.I32_wrap_i64(int64(base.Ui64(v1518) >> (uint(v2466) % 64)))
	*(*int32)(unsafe.Add(mBase, uint32(v623)+84)) = v2474
	*(*int32)(unsafe.Add(mBase, uint32(v623)+80)) = v942
	v2483 = F_pg_snprintf(m, v623+int32(480), int32(1024), int32(_a_F_WalSummarizerMain_11), v623+int32(80))
	mBase = m.M
	v2484 = m.ExcPending
	if v2484 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L456
	}
L456:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	*(*int64)(unsafe.Add(mBase, uint32(v623)+472)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	v2490 = F_PathNameOpenFile(m, v2456, int32(577))
	mBase = m.M
	v2491 = m.ExcPending
	if v2491 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L457
	}
L457:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+464)) = v2490
	if int32(0) <= v2490 {
		goto L231
	} else {
		goto L458
	}
L458:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2500 = m.ExcPending
	if v2500 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L459
	}
L459:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	F_errcode_for_file_access(m)
	mBase = m.M
	v2504 = m.ExcPending
	if v2504 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L460
	}
L460:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	*(*int32)(unsafe.Add(mBase, uint32(v623)+32)) = v2456
	F_errmsg(m, int32(_a_F_WalSummarizerMain_28), v623+int32(32))
	mBase = m.M
	v2512 = m.ExcPending
	if v2512 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L461
	}
L461:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	F_errfinish(m, int32(_a_F_WalSummarizerMain_1), int32(1215), int32(_a_F_WalSummarizerMain_21))
	mBase = m.M
	v2519 = m.ExcPending
	if v2519 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L462
	}
L462:
	;
	goto L3
L463:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	F_errcode_for_file_access(m)
	mBase = m.M
	v2530 = m.ExcPending
	if v2530 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L464
	}
L464:
	;
	v2531 = *(*int64)(unsafe.Add(mBase, uint32(v1047)+40))
	v2534 = base.I32_wrap_i64(int64(base.Ui64(v2531) >> (uint(int64(32)) % 64)))
	v2535 = base.I32_wrap_i64(v2531)
	if v2522 != 0 {
		goto L465
	} else {
		goto L466
	}
L465:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	v2538 = *(*int32)(unsafe.Add(mBase, uint32(v623)+448))
	*(*int32)(unsafe.Add(mBase, uint32(v623)+172)) = v2538
	*(*int32)(unsafe.Add(mBase, uint32(v623)+168)) = v2535
	*(*int32)(unsafe.Add(mBase, uint32(v623)+164)) = v2534
	*(*int32)(unsafe.Add(mBase, uint32(v623)+160)) = v942
	F_errmsg(m, int32(_a_F_WalSummarizerMain_29), v623+int32(160))
	mBase = m.M
	v2547 = m.ExcPending
	if v2547 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L468
	}
L466:
	;
	goto L467
L467:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	*(*int32)(unsafe.Add(mBase, uint32(v623)+152)) = v2535
	*(*int32)(unsafe.Add(mBase, uint32(v623)+148)) = v2534
	*(*int32)(unsafe.Add(mBase, uint32(v623)+144)) = v942
	F_errmsg(m, int32(_a_F_WalSummarizerMain_30), v623+int32(144))
	mBase = m.M
	v2564 = m.ExcPending
	if v2564 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L470
	}
L468:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	F_errfinish(m, int32(_a_F_WalSummarizerMain_1), int32(1050), int32(_a_F_WalSummarizerMain_21))
	mBase = m.M
	v2554 = m.ExcPending
	if v2554 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L469
	}
L469:
	;
	goto L3
L470:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	F_errfinish(m, int32(_a_F_WalSummarizerMain_1), int32(1055), int32(_a_F_WalSummarizerMain_21))
	mBase = m.M
	v2571 = m.ExcPending
	if v2571 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L471
	}
L471:
	;
	goto L3
L472:
	;
	v3437 = m.G0
	v3439 = v3437 - int32(32)
	m.G0 = v3439
	v3441 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3439)+24)) = v3441
	*(*int64)(unsafe.Add(mBase, uint32(v3439)+16)) = v3441
	*(*int64)(unsafe.Add(mBase, uint32(v3439)+8)) = v3441
	v3448 = v2577 + int32(12)
	v3449 = *(*int32)(unsafe.Add(mBase, uint32(v2577)+uint32(_c_F_WalSummarizerMain[46])))
	v3452 = int32(24)
	v3453 = m.Env.Pgmem_crc32c(m, v3449, v3439+int32(8), v3452)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v2577)+uint32(_c_F_WalSummarizerMain[46]))) = v3453
	v3455 = *(*int32)(unsafe.Add(mBase, uint32(v2577)+uint32(_c_F_WalSummarizerMain[47])))
	if int32(_a_F_WalSummarizerMain_31) <= v3455+v3452 {
		goto L584
	} else {
		goto L585
	}
L473:
	;
	v2608 = F_palloc(m, v2603*int32(24))
	mBase = m.M
	v2609 = m.ExcPending
	if v2609 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L474
	}
L474:
	;
	v2610 = *(*int32)(unsafe.Add(mBase, uint32(v1023)))
	v2611 = *(*int64)(unsafe.Add(mBase, uint32(v2610)))
	if v2611 == int64(0) {
		v2656 = v2591
		goto L475
	} else {
		goto L476
	}
L475:
	;
	v2687 = v2577 + int32(20)
	v2693 = v2656
	v2695 = int32(0)
	v2698 = v2610
	v2703 = v2574
	goto L483
L476:
	;
	v2614 = *(*int32)(unsafe.Add(mBase, uint32(v2610)+20))
	v2621 = int32(0)
	goto L477
L477:
	;
	v2649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2614+v2621*int32(40))+20)))
	if v2649 != int32(1) {
		goto L479
	} else {
		goto L480
	}
L478:
	;
	v2656 = v2591
	goto L475
L479:
	;
	v2656 = v2621
	goto L475
L480:
	;
	goto L481
L481:
	;
	v2653 = v2621 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v2653)) < base.Ui64(v2611) {
		v2621 = v2653
		goto L477
	} else {
		goto L482
	}
L482:
	;
	goto L478
L483:
	;
	v2724 = v2693
	v2725 = v2695
	v2727 = v2695
	goto L486
L484:
	;
	F_pg_qsort(m, v2608, v2703, int32(24), int32(_a_F_WalSummarizerMain_32))
	mBase = m.M
	v2827 = m.ExcPending
	if v2827 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L495
	}
L485:
	;
	goto L484
L486:
	;
	if v2727&int32(1) != 0 {
		goto L485
	} else {
		goto L488
	}
L487:
	;
	v2769 = v2608 + v2703*int32(24)
	v2770 = *(*int32)(unsafe.Add(mBase, uint32(v2763)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2769)+8)) = v2770
	v2772 = *(*int64)(unsafe.Add(mBase, uint32(v2763)))
	*(*int64)(unsafe.Add(mBase, uint32(v2769))) = v2772
	v2774 = *(*int32)(unsafe.Add(mBase, uint32(v2763)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v2769)+12)) = v2774
	v2776 = *(*int32)(unsafe.Add(mBase, uint32(v2763)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v2769)+16)) = v2776
	v2778 = *(*int32)(unsafe.Add(mBase, uint32(v2763)+24))
	v2784 = v2778
	goto L490
L488:
	;
	v2751 = *(*int32)(unsafe.Add(mBase, uint32(v2698)+12))
	v2752 = int32(1)
	v2753 = v2724 - v2752
	v2757 = base.B2i32(v2751&(v2753^v2656) == int32(0))
	v2758 = v2757 | v2725
	v2761 = v2751 & v2753
	v2762 = *(*int32)(unsafe.Add(mBase, uint32(v2698)+20))
	v2763 = v2724*int32(40) + v2762
	v2764 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2763)+20)))
	if v2764 != v2752 {
		v2724 = v2761
		v2725 = v2758
		v2727 = v2757
		goto L486
	} else {
		goto L489
	}
L489:
	;
	goto L487
L490:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2769)+20)) = v2784
	if v2784 == int32(0) {
		goto L492
	} else {
		goto L493
	}
L491:
	;
	v2823 = *(*int32)(unsafe.Add(mBase, uint32(v1023)))
	v2693 = v2761
	v2695 = v2758
	v2698 = v2823
	v2703 = v2703 + int32(1)
	goto L483
L492:
	;
	goto L491
L493:
	;
	v2812 = *(*int32)(unsafe.Add(mBase, uint32(v2763)+32))
	v2818 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2812+v2784<<(uint(int32(1))%32)-int32(2)))))
	if v2818 != 0 {
		goto L492
	} else {
		goto L494
	}
L494:
	;
	v2784 = v2784 - int32(1)
	goto L490
L495:
	;
	v2828 = *(*int32)(unsafe.Add(mBase, uint32(v1023)))
	v2829 = *(*int32)(unsafe.Add(mBase, uint32(v2828)+8))
	if v2829 == int32(0) {
		goto L472
	} else {
		goto L496
	}
L496:
	;
	v2837 = int32(0)
	goto L497
L497:
	;
	v2863 = *(*int32)(unsafe.Add(mBase, uint32(v2577)+uint32(_c_F_WalSummarizerMain[46])))
	v2864 = int32(24)
	v2866 = v2608 + v2837*v2864
	v2868 = m.Env.Pgmem_crc32c(m, v2863, v2866, v2864)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v2577)+uint32(_c_F_WalSummarizerMain[46]))) = v2868
	v2870 = *(*int32)(unsafe.Add(mBase, uint32(v2577)+uint32(_c_F_WalSummarizerMain[47])))
	if int32(_a_F_WalSummarizerMain_31) <= v2870+v2864 {
		goto L499
	} else {
		goto L500
	}
L498:
	;
	goto L472
L499:
	;
	v2875 = *(*int32)(unsafe.Add(mBase, uint32(v2577)+16))
	v2876 = *(*int32)(unsafe.Add(mBase, uint32(v2577)+12))
	v2877 = m.T0[v2876].(func(*base.Module, int32, int32, int32) int32)(m, v2875, v2687, v2870)
	mBase = m.M
	v2878 = m.ExcPending
	if v2878 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L502
	}
L500:
	;
	v2882 = v2870
	goto L501
L501:
	;
	v2883 = v2882 + v2687
	v2884 = *(*int64)(unsafe.Add(mBase, uint32(v2866)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v2883)+16)) = v2884
	v2886 = *(*int64)(unsafe.Add(mBase, uint32(v2866)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2883)+8)) = v2886
	v2888 = *(*int64)(unsafe.Add(mBase, uint32(v2866)))
	*(*int64)(unsafe.Add(mBase, uint32(v2883))) = v2888
	v2890 = *(*int32)(unsafe.Add(mBase, uint32(v2577)+uint32(_c_F_WalSummarizerMain[47])))
	*(*int32)(unsafe.Add(mBase, uint32(v2577)+uint32(_c_F_WalSummarizerMain[47]))) = v2890 + int32(24)
	v2894 = *(*int32)(unsafe.Add(mBase, uint32(v1023)))
	v2895 = *(*int32)(unsafe.Add(mBase, uint32(v2866)+12))
	v2896 = *(*int32)(unsafe.Add(mBase, uint32(v2866)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2577)+uint32(_c_F_WalSummarizerMain[48]))) = v2896
	v2898 = *(*int64)(unsafe.Add(mBase, uint32(v2866)))
	*(*int64)(unsafe.Add(mBase, uint32(v2577)+uint32(_c_F_WalSummarizerMain[49]))) = v2898
	*(*int32)(unsafe.Add(mBase, uint32(v2577)+uint32(_c_F_WalSummarizerMain[50]))) = v2895
	v2902 = v2577 + int32(_a_F_WalSummarizerMain_33)
	v2903 = int32(16)
	v2909 = int32(-1636608416)
	if v2902&int32(3) != 0 {
		goto L507
	} else {
		goto L508
	}
L502:
	;
	v2879 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2577)+uint32(_c_F_WalSummarizerMain[47]))) = v2879
	v2882 = v2879
	goto L501
L503:
	;
	v3168 = *(*int32)(unsafe.Add(mBase, uint32(v2894)+20))
	v3169 = *(*int32)(unsafe.Add(mBase, uint32(v2894)+12))
	v3170 = (v3163 ^ v3155 - base.I32_rotl(v3163, int32(24))) & v3169
	v3173 = v3168 + v3170*int32(40)
	v3174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3173)+20)))
	if v3174 == int32(0) {
		goto L544
	} else {
		goto L545
	}
L504:
	;
	v3141 = int32(14)
	v3143 = v3137 ^ v3138 - base.I32_rotl(v3137, v3141)
	v3147 = v3143 ^ v3136 - base.I32_rotl(v3143, int32(11))
	v3151 = v3147 ^ v3137 - base.I32_rotl(v3147, int32(25))
	v3155 = v3151 ^ v3143 - base.I32_rotl(v3151, int32(16))
	v3159 = v3155 ^ v3147 - base.I32_rotl(v3155, int32(4))
	v3163 = v3159 ^ v3151 - base.I32_rotl(v3159, v3141)
	goto L503
L505:
	;
	switch v3063 - int32(1) {
	case 0:
		v3129 = v3054
		v3130 = v3055
		v3131 = v3059
		goto L532
	case 1:
		v3122 = v3054
		v3123 = v3055
		v3124 = v3059
		goto L533
	case 2:
		v3115 = v3054
		v3116 = v3055
		v3117 = v3059
		goto L534
	case 3:
		v3109 = v3055
		v3110 = v3059
		goto L535
	case 4:
		v3105 = v3055
		v3106 = v3059
		goto L536
	case 5:
		v3099 = v3055
		v3100 = v3059
		goto L537
	case 6:
		v3093 = v3055
		v3094 = v3059
		goto L538
	case 7:
		v3088 = v3059
		goto L539
	case 8:
		v3083 = v3059
		goto L540
	case 9:
		v3078 = v3059
		goto L541
	case 10:
		goto L542
	default:
		v3136 = v3054
		v3137 = v3055
		v3138 = v3059
		goto L504
	}
L506:
	;
	v3018 = v2902
	v3019 = v2903
	v3020 = v2909
	v3021 = v2909
	v3022 = v2909
	goto L529
L507:
	;
	goto L506
L508:
	;
	goto L509
L509:
	;
	goto L513
L511:
	;
	switch v2961 - int32(1) {
	case 0:
		v3015 = v2952
		goto L518
	case 1:
		v3010 = v2952
		goto L519
	case 2:
		goto L520
	case 3:
		v3003 = v2953
		goto L521
	case 4:
		v3000 = v2953
		goto L522
	case 5:
		v2995 = v2953
		goto L523
	case 6:
		goto L524
	case 7:
		v2986 = v2957
		goto L525
	case 8:
		v2981 = v2957
		goto L526
	case 9:
		v2976 = v2957
		goto L527
	case 10:
		goto L528
	default:
		v3136 = v2952
		v3137 = v2953
		v3138 = v2957
		goto L504
	}
L513:
	;
	goto L514
L514:
	;
	v2916 = v2902
	v2917 = v2903
	v2918 = v2909
	v2919 = v2909
	v2920 = v2909
	goto L515
L515:
	;
	v2922 = *(*int32)(unsafe.Add(mBase, uint32(v2916)+4))
	v2923 = v2922 + v2919
	v2924 = *(*int32)(unsafe.Add(mBase, uint32(v2916)))
	v2926 = *(*int32)(unsafe.Add(mBase, uint32(v2916)+8))
	v2927 = v2926 + v2920
	v2929 = int32(4)
	v2931 = v2924 + v2918 - v2927 ^ base.I32_rotl(v2927, v2929)
	v2935 = v2923 - v2931 ^ base.I32_rotl(v2931, int32(6))
	v2936 = v2927 + v2923
	v2937 = v2931 + v2936
	v2938 = v2935 + v2937
	v2942 = v2936 - v2935 ^ base.I32_rotl(v2935, int32(8))
	v2946 = v2937 - v2942 ^ base.I32_rotl(v2942, int32(16))
	v2950 = v2938 - v2946 ^ base.I32_rotl(v2946, int32(19))
	v2951 = v2942 + v2938
	v2952 = v2946 + v2951
	v2953 = v2950 + v2952
	v2957 = v2951 - v2950 ^ base.I32_rotl(v2950, v2929)
	v2958 = int32(12)
	v2959 = v2916 + v2958
	v2961 = v2917 - v2958
	if base.Ui32(int32(11)) < base.Ui32(v2961) {
		v2916 = v2959
		v2917 = v2961
		v2918 = v2952
		v2919 = v2953
		v2920 = v2957
		goto L515
	} else {
		goto L517
	}
L516:
	;
	goto L511
L517:
	;
	goto L516
L518:
	;
	v3016 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2959))))
	v3136 = v3015 + v3016
	v3137 = v2953
	v3138 = v2957
	goto L504
L519:
	;
	v3011 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2959)+1)))
	v3015 = v3011<<(uint(int32(8))%32) + v3010
	goto L518
L520:
	;
	v3006 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2959)+2)))
	v3010 = v3006<<(uint(int32(16))%32) + v2952
	goto L519
L521:
	;
	v3004 = *(*int32)(unsafe.Add(mBase, uint32(v2959)))
	v3136 = v3004 + v2952
	v3137 = v3003
	v3138 = v2957
	goto L504
L522:
	;
	v3001 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2959)+4)))
	v3003 = v3000 + v3001
	goto L521
L523:
	;
	v2996 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2959)+5)))
	v3000 = v2996<<(uint(int32(8))%32) + v2995
	goto L522
L524:
	;
	v2991 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2959)+6)))
	v2995 = v2991<<(uint(int32(16))%32) + v2953
	goto L523
L525:
	;
	v2987 = *(*int32)(unsafe.Add(mBase, uint32(v2959)))
	v2989 = *(*int32)(unsafe.Add(mBase, uint32(v2959)+4))
	v3136 = v2987 + v2952
	v3137 = v2989 + v2953
	v3138 = v2986
	goto L504
L526:
	;
	v2982 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2959)+8)))
	v2986 = v2982<<(uint(int32(8))%32) + v2981
	goto L525
L527:
	;
	v2977 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2959)+9)))
	v2981 = v2977<<(uint(int32(16))%32) + v2976
	goto L526
L528:
	;
	v2972 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2959)+10)))
	v2976 = v2972<<(uint(int32(24))%32) + v2957
	goto L527
L529:
	;
	v3024 = *(*int32)(unsafe.Add(mBase, uint32(v3018)+4))
	v3025 = v3024 + v3021
	v3026 = *(*int32)(unsafe.Add(mBase, uint32(v3018)))
	v3028 = *(*int32)(unsafe.Add(mBase, uint32(v3018)+8))
	v3029 = v3028 + v3022
	v3031 = int32(4)
	v3033 = v3026 + v3020 - v3029 ^ base.I32_rotl(v3029, v3031)
	v3037 = v3025 - v3033 ^ base.I32_rotl(v3033, int32(6))
	v3038 = v3029 + v3025
	v3039 = v3033 + v3038
	v3040 = v3037 + v3039
	v3044 = v3038 - v3037 ^ base.I32_rotl(v3037, int32(8))
	v3048 = v3039 - v3044 ^ base.I32_rotl(v3044, int32(16))
	v3052 = v3040 - v3048 ^ base.I32_rotl(v3048, int32(19))
	v3053 = v3044 + v3040
	v3054 = v3048 + v3053
	v3055 = v3052 + v3054
	v3059 = v3053 - v3052 ^ base.I32_rotl(v3052, v3031)
	v3060 = int32(12)
	v3061 = v3018 + v3060
	v3063 = v3019 - v3060
	if base.Ui32(int32(11)) < base.Ui32(v3063) {
		v3018 = v3061
		v3019 = v3063
		v3020 = v3054
		v3021 = v3055
		v3022 = v3059
		goto L529
	} else {
		goto L531
	}
L530:
	;
	goto L505
L531:
	;
	goto L530
L532:
	;
	v3132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3061))))
	v3136 = v3129 + v3132
	v3137 = v3130
	v3138 = v3131
	goto L504
L533:
	;
	v3125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3061)+1)))
	v3129 = v3125<<(uint(int32(8))%32) + v3122
	v3130 = v3123
	v3131 = v3124
	goto L532
L534:
	;
	v3118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3061)+2)))
	v3122 = v3118<<(uint(int32(16))%32) + v3115
	v3123 = v3116
	v3124 = v3117
	goto L533
L535:
	;
	v3111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3061)+3)))
	v3115 = v3111<<(uint(int32(24))%32) + v3054
	v3116 = v3109
	v3117 = v3110
	goto L534
L536:
	;
	v3107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3061)+4)))
	v3109 = v3105 + v3107
	v3110 = v3106
	goto L535
L537:
	;
	v3101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3061)+5)))
	v3105 = v3101<<(uint(int32(8))%32) + v3099
	v3106 = v3100
	goto L536
L538:
	;
	v3095 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3061)+6)))
	v3099 = v3095<<(uint(int32(16))%32) + v3093
	v3100 = v3094
	goto L537
L539:
	;
	v3089 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3061)+7)))
	v3093 = v3089<<(uint(int32(24))%32) + v3055
	v3094 = v3088
	goto L538
L540:
	;
	v3084 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3061)+8)))
	v3088 = v3084<<(uint(int32(8))%32) + v3083
	goto L539
L541:
	;
	v3079 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3061)+9)))
	v3083 = v3079<<(uint(int32(16))%32) + v3078
	goto L540
L542:
	;
	v3074 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3061)+10)))
	v3078 = v3074<<(uint(int32(24))%32) + v3059
	goto L541
L543:
	;
	v3257 = *(*int32)(unsafe.Add(mBase, uint32(v2866)+20))
	if v3257 == int32(0) {
		goto L551
	} else {
		goto L552
	}
L544:
	;
	v3232 = int32(0)
	goto L543
L545:
	;
	goto L546
L546:
	;
	v3183 = v3173
	v3184 = v3170
	goto L547
L547:
	;
	v3208 = *(*int64)(unsafe.Add(mBase, uint32(v3183)))
	v3209 = *(*int64)(unsafe.Add(mBase, uint32(v2577)+uint32(_c_F_WalSummarizerMain[49])))
	v3211 = *(*int64)(unsafe.Add(mBase, uint32(v3183)+8))
	v3214 = *(*int64)(unsafe.Add(mBase, uint32(v2577)+uint32(_c_F_WalSummarizerMain[48])))
	if v3208^v3209|(v3211^v3214) == int64(0) {
		v3232 = v3183
		goto L543
	} else {
		goto L549
	}
L548:
	;
	v3232 = int32(0)
	goto L543
L549:
	;
	v3221 = (v3184 + int32(1)) & v3169
	v3224 = v3168 + v3221*int32(40)
	v3225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3224)+20)))
	if v3225 != 0 {
		v3183 = v3224
		v3184 = v3221
		goto L547
	} else {
		goto L550
	}
L550:
	;
	goto L548
L551:
	;
	v3292 = *(*int32)(unsafe.Add(mBase, uint32(v3232)+24))
	if v3292 != 0 {
		goto L564
	} else {
		goto L565
	}
L552:
	;
	v3260 = *(*int32)(unsafe.Add(mBase, uint32(v2577)+uint32(_c_F_WalSummarizerMain[46])))
	v3261 = *(*int32)(unsafe.Add(mBase, uint32(v3232)+32))
	v3263 = v3257 << (uint(int32(1)) % 32)
	v3264 = m.Env.Pgmem_crc32c(m, v3260, v3261, v3263)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v2577)+uint32(_c_F_WalSummarizerMain[46]))) = v3264
	v3266 = *(*int32)(unsafe.Add(mBase, uint32(v2577)+uint32(_c_F_WalSummarizerMain[47])))
	if int32(_a_F_WalSummarizerMain_31) <= v3266+v3263 {
		goto L553
	} else {
		goto L554
	}
L553:
	;
	v3270 = *(*int32)(unsafe.Add(mBase, uint32(v2577)+16))
	v3271 = *(*int32)(unsafe.Add(mBase, uint32(v2577)+12))
	v3272 = m.T0[v3271].(func(*base.Module, int32, int32, int32) int32)(m, v3270, v2687, v3266)
	mBase = m.M
	v3273 = m.ExcPending
	if v3273 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L556
	}
L554:
	;
	v3277 = v3266
	goto L555
L555:
	;
	if int32(_a_F_WalSummarizerMain_34) <= v3263 {
		goto L557
	} else {
		goto L558
	}
L556:
	;
	v3274 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2577)+uint32(_c_F_WalSummarizerMain[47]))) = v3274
	v3277 = v3274
	goto L555
L557:
	;
	v3280 = *(*int32)(unsafe.Add(mBase, uint32(v2577)+16))
	v3281 = *(*int32)(unsafe.Add(mBase, uint32(v2577)+12))
	v3282 = m.T0[v3281].(func(*base.Module, int32, int32, int32) int32)(m, v3280, v3261, v3263)
	mBase = m.M
	v3283 = m.ExcPending
	if v3283 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L560
	}
L558:
	;
	goto L559
L559:
	;
	if v3263 != 0 {
		goto L561
	} else {
		goto L562
	}
L560:
	;
	goto L551
L561:
	;
	base.MemoryCopy(m, v3277+v2687, v3261, v3263)
	goto L563
L562:
	;
	goto L563
L563:
	;
	v3286 = *(*int32)(unsafe.Add(mBase, uint32(v2577)+uint32(_c_F_WalSummarizerMain[47])))
	*(*int32)(unsafe.Add(mBase, uint32(v2577)+uint32(_c_F_WalSummarizerMain[47]))) = v3286 + v3263
	goto L551
L564:
	;
	v3300 = int32(0)
	goto L567
L565:
	;
	goto L566
L566:
	;
	v3403 = v2837 + int32(1)
	v3404 = *(*int32)(unsafe.Add(mBase, uint32(v1023)))
	v3405 = *(*int32)(unsafe.Add(mBase, uint32(v3404)+8))
	if base.Ui32(v3403) < base.Ui32(v3405) {
		v2837 = v3403
		goto L497
	} else {
		goto L583
	}
L567:
	;
	v3324 = *(*int32)(unsafe.Add(mBase, uint32(v3232)+32))
	v3328 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3324+v3300<<(uint(int32(1))%32)))))
	if v3328 == int32(0) {
		goto L569
	} else {
		goto L570
	}
L568:
	;
	goto L566
L569:
	;
	v3369 = v3300 + int32(1)
	v3370 = *(*int32)(unsafe.Add(mBase, uint32(v3232)+24))
	if base.Ui32(v3369) < base.Ui32(v3370) {
		v3300 = v3369
		goto L567
	} else {
		goto L582
	}
L570:
	;
	v3332 = *(*int32)(unsafe.Add(mBase, uint32(v2577)+uint32(_c_F_WalSummarizerMain[46])))
	v3333 = *(*int32)(unsafe.Add(mBase, uint32(v3232)+36))
	v3337 = *(*int32)(unsafe.Add(mBase, uint32(v3333+v3300<<(uint(int32(2))%32))))
	v3339 = v3328 << (uint(int32(1)) % 32)
	v3340 = m.Env.Pgmem_crc32c(m, v3332, v3337, v3339)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v2577)+uint32(_c_F_WalSummarizerMain[46]))) = v3340
	v3342 = *(*int32)(unsafe.Add(mBase, uint32(v2577)+uint32(_c_F_WalSummarizerMain[47])))
	if int32(_a_F_WalSummarizerMain_31) <= v3342+v3339 {
		goto L571
	} else {
		goto L572
	}
L571:
	;
	v3346 = *(*int32)(unsafe.Add(mBase, uint32(v2577)+16))
	v3347 = *(*int32)(unsafe.Add(mBase, uint32(v2577)+12))
	v3348 = m.T0[v3347].(func(*base.Module, int32, int32, int32) int32)(m, v3346, v2687, v3342)
	mBase = m.M
	v3349 = m.ExcPending
	if v3349 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L574
	}
L572:
	;
	v3353 = v3342
	goto L573
L573:
	;
	if base.I32_extend16_s(v3328) < int32(0) {
		goto L575
	} else {
		goto L576
	}
L574:
	;
	v3350 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2577)+uint32(_c_F_WalSummarizerMain[47]))) = v3350
	v3353 = v3350
	goto L573
L575:
	;
	v3356 = *(*int32)(unsafe.Add(mBase, uint32(v2577)+16))
	v3357 = *(*int32)(unsafe.Add(mBase, uint32(v2577)+12))
	v3358 = m.T0[v3357].(func(*base.Module, int32, int32, int32) int32)(m, v3356, v3337, v3339)
	mBase = m.M
	v3359 = m.ExcPending
	if v3359 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L578
	}
L576:
	;
	goto L577
L577:
	;
	if v3339 != 0 {
		goto L579
	} else {
		goto L580
	}
L578:
	;
	goto L569
L579:
	;
	base.MemoryCopy(m, v3353+v2687, v3337, v3339)
	goto L581
L580:
	;
	goto L581
L581:
	;
	v3362 = *(*int32)(unsafe.Add(mBase, uint32(v2577)+uint32(_c_F_WalSummarizerMain[47])))
	*(*int32)(unsafe.Add(mBase, uint32(v2577)+uint32(_c_F_WalSummarizerMain[47]))) = v3362 + v3339
	goto L569
L582:
	;
	goto L568
L583:
	;
	goto L498
L584:
	;
	v3460 = *(*int32)(unsafe.Add(mBase, uint32(v3448)+4))
	v3463 = *(*int32)(unsafe.Add(mBase, uint32(v3448)))
	v3464 = m.T0[v3463].(func(*base.Module, int32, int32, int32) int32)(m, v3460, v2577+int32(20), v3455)
	mBase = m.M
	v3465 = m.ExcPending
	if v3465 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L587
	}
L585:
	;
	v3469 = v3455
	goto L586
L586:
	;
	v3471 = v2577 + int32(20)
	v3472 = v3469 + v3471
	v3473 = *(*int64)(unsafe.Add(mBase, uint32(v3439)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v3472)+16)) = v3473
	v3475 = *(*int64)(unsafe.Add(mBase, uint32(v3439)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v3472)+8)) = v3475
	v3477 = *(*int64)(unsafe.Add(mBase, uint32(v3439)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v3472))) = v3477
	v3479 = *(*int32)(unsafe.Add(mBase, uint32(v2577)+uint32(_c_F_WalSummarizerMain[47])))
	*(*int32)(unsafe.Add(mBase, uint32(v2577)+uint32(_c_F_WalSummarizerMain[47]))) = v3479 + int32(24)
	v3483 = *(*int32)(unsafe.Add(mBase, uint32(v2577)+uint32(_c_F_WalSummarizerMain[46])))
	*(*int32)(unsafe.Add(mBase, uint32(v3439)+4)) = v3483 ^ int32(-1)
	v3487 = int32(4)
	v3490 = m.Env.Pgmem_crc32c(m, v3483, v3439+v3487, v3487)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v2577)+uint32(_c_F_WalSummarizerMain[46]))) = v3490
	v3492 = *(*int32)(unsafe.Add(mBase, uint32(v2577)+uint32(_c_F_WalSummarizerMain[47])))
	if int32(_a_F_WalSummarizerMain_31) <= v3492+v3487 {
		goto L588
	} else {
		goto L589
	}
L587:
	;
	v3466 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2577)+uint32(_c_F_WalSummarizerMain[47]))) = v3466
	v3469 = v3466
	goto L586
L588:
	;
	v3497 = *(*int32)(unsafe.Add(mBase, uint32(v3448)+4))
	v3498 = *(*int32)(unsafe.Add(mBase, uint32(v3448)))
	v3499 = m.T0[v3498].(func(*base.Module, int32, int32, int32) int32)(m, v3497, v3471, v3492)
	mBase = m.M
	v3500 = m.ExcPending
	if v3500 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L591
	}
L589:
	;
	v3504 = v3492
	goto L590
L590:
	;
	v3506 = *(*int32)(unsafe.Add(mBase, uint32(v3439)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v3504+v3471))) = v3506
	v3508 = *(*int32)(unsafe.Add(mBase, uint32(v2577)+uint32(_c_F_WalSummarizerMain[47])))
	v3510 = v3508 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v2577)+uint32(_c_F_WalSummarizerMain[47]))) = v3510
	v3512 = *(*int32)(unsafe.Add(mBase, uint32(v3448)+4))
	v3513 = *(*int32)(unsafe.Add(mBase, uint32(v3448)))
	v3514 = m.T0[v3513].(func(*base.Module, int32, int32, int32) int32)(m, v3512, v3471, v3510)
	mBase = m.M
	v3515 = m.ExcPending
	if v3515 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L592
	}
L591:
	;
	v3501 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2577)+uint32(_c_F_WalSummarizerMain[47]))) = v3501
	v3504 = v3501
	goto L590
L592:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2577)+uint32(_c_F_WalSummarizerMain[47]))) = int32(0)
	m.G0 = v3439 + int32(32)
	m.G0 = v2577 + int32(_a_F_WalSummarizerMain_22)
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	v3526 = *(*int32)(unsafe.Add(mBase, uint32(v623)+464))
	F_FileClose(m, v3526)
	mBase = m.M
	v3528 = m.ExcPending
	if v3528 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L593
	}
L593:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	v3533 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v3534 = m.ExcPending
	if v3534 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L594
	}
L594:
	;
	if v3533 != 0 {
		goto L595
	} else {
		goto L596
	}
L595:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	*(*int32)(unsafe.Add(mBase, uint32(v623-int32(-64)))) = v2464
	*(*int32)(unsafe.Add(mBase, uint32(v623)+60)) = v2468
	*(*int32)(unsafe.Add(mBase, uint32(v623)+56)) = v2470
	*(*int32)(unsafe.Add(mBase, uint32(v623)+52)) = v2474
	*(*int32)(unsafe.Add(mBase, uint32(v623)+48)) = v942
	F_errmsg_internal(m, int32(_a_F_WalSummarizerMain_35), v623+int32(48))
	mBase = m.M
	v3548 = m.ExcPending
	if v3548 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L598
	}
L596:
	;
	goto L597
L597:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	v3563 = F_durable_rename(m, v623+int32(1504), v623+int32(480), int32(21))
	mBase = m.M
	v3564 = m.ExcPending
	if v3564 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L600
	}
L598:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	F_errfinish(m, int32(_a_F_WalSummarizerMain_1), int32(1228), int32(_a_F_WalSummarizerMain_21))
	mBase = m.M
	v3555 = m.ExcPending
	if v3555 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L599
	}
L599:
	;
	goto L597
L600:
	;
	goto L230
L601:
	;
	v3629 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v623)+267)) = uint8(v3629)
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	v3634 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[29]))
	v3638 = F_LWLockAcquire(m, v3634+int32(_a_F_WalSummarizerMain_4), int32(0))
	mBase = m.M
	v3639 = m.ExcPending
	if v3639 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L607
	}
L602:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	v3604 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v3605 = m.ExcPending
	if v3605 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L603
	}
L603:
	;
	if v3604 == int32(0) {
		goto L601
	} else {
		goto L604
	}
L604:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	*(*uint32)(unsafe.Add(mBase, uint32(v623)+16)) = uint32(v2435)
	v3611 = int64(32)
	v3612 = int64(base.Ui64(v2435) >> (uint(v3611) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v623)+12)) = uint32(v3612)
	*(*int32)(unsafe.Add(mBase, uint32(v623))) = v942
	*(*uint32)(unsafe.Add(mBase, uint32(v623)+8)) = uint32(v1518)
	v3617 = int64(base.Ui64(v1518) >> (uint(v3611) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v623)+4)) = uint32(v3617)
	F_errmsg_internal(m, int32(_a_F_WalSummarizerMain_36), v623)
	mBase = m.M
	v3621 = m.ExcPending
	if v3621 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L605
	}
L605:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	F_errfinish(m, int32(_a_F_WalSummarizerMain_1), int32(1240), int32(_a_F_WalSummarizerMain_21))
	mBase = m.M
	v3628 = m.ExcPending
	if v3628 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L606
	}
L606:
	;
	goto L601
L607:
	;
	v3641 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[30]))
	*(*int64)(unsafe.Add(mBase, uint32(v3641)+24)) = v2435
	v3643 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3641)+16)) = uint8(v3643)
	*(*int32)(unsafe.Add(mBase, uint32(v3641)+4)) = v942
	*(*int64)(unsafe.Add(mBase, uint32(v3641)+8)) = v2435
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	v3650 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[29]))
	F_LWLockRelease(m, v3650+int32(_a_F_WalSummarizerMain_4))
	mBase = m.M
	v3654 = m.ExcPending
	if v3654 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L608
	}
L608:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v623)+3120)) = v1515
	*(*int32)(unsafe.Add(mBase, uint32(v623)+3132)) = v624
	v3658 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[30]))
	F_ConditionVariableBroadcast(m, v3658+int32(32))
	mBase = m.M
	v3662 = m.ExcPending
	if v3662 != 0 {
		v3700 = v623
		v3714 = v637
		goto L6
	} else {
		goto L609
	}
L609:
	;
	v3686 = v1515
	v3687 = v987
	v3690 = v2435
	goto L213
L610:
	;
	goto L5
L611:
	;
	v3733 = int32(v3729)
	m.G0 = v3700
	v3735 = *(*int32)(unsafe.Add(mBase, uint32(v3733)+4))
	v3736 = *(*int32)(unsafe.Add(mBase, uint32(v3733)))
	v3739 = *(*int32)(unsafe.Add(mBase, uint32(v3736)))
	if v3700+int32(252) == v3739 {
		goto L614
	} else {
		goto L615
	}
L612:
	;
	m.ExcPending = 1
	goto L620
L613:
	;
	if v3743 != 0 {
		goto L617
	} else {
		goto L618
	}
L614:
	;
	v3741 = *(*int32)(unsafe.Add(mBase, uint32(v3736)+4))
	v3743 = v3741
	goto L616
L615:
	;
	v3743 = int32(0)
	goto L616
L616:
	;
	goto L613
L617:
	;
	v3744 = *(*int32)(unsafe.Add(mBase, uint32(v3700)+3132))
	v3745 = *(*int64)(unsafe.Add(mBase, uint32(v3700)+3120))
	v38 = v3743
	v40 = v3700
	v41 = v3744
	v43 = v3735
	v54 = v3714
	v61 = v3745
	goto L1
L618:
	;
	goto L619
L619:
	;
	F___wasm_longjmp(m, v3736, v3735)
	mBase = m.M
	v3747 = m.ExcPending
	if v3747 != 0 {
		goto L620
	} else {
		goto L621
	}
L620:
	;
	return
L621:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
