package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
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
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	v1 = l0
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_SetWalWriterSleeping[0]))
	v7 = base.AtomicRmwXchg32(m, v4, int32(440), int32(1))
	if v7 != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_SetWalWriterSleeping[0]))
		F_s_lock(m, v9+int32(440), int32(_a_F_SetWalWriterSleeping_0), int32(_a_F_SetWalWriterSleeping_1), int32(_a_F_SetWalWriterSleeping_2))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, _c_F_SetWalWriterSleeping[0]))
			*(*uint8)(unsafe.Add(mBase, uint32(v18)+321)) = uint8(v1)
			v20 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v18)+440)), uint32(v20))
			return
		}
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_SetWalWriterSleeping[0]))
		*(*uint8)(unsafe.Add(mBase, uint32(v18)+321)) = uint8(v1)
		v20 = int32(0)
		atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v18)+440)), uint32(v20))
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
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_WalRcvRunning[0]))
	v7 = int32(1456)
	v8 = v6 + v7
	v11 = base.AtomicRmwXchg32(m, v6, v7, int32(1))
	if v11 != 0 {
		F_s_lock(m, v8, int32(_a_F_WalRcvRunning_0), int32(82), int32(_a_F_WalRcvRunning_1))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = *(*int64)(unsafe.Add(mBase, uint32(v6)+24))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
			v21 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v6)+1456)), uint32(v21))
			if v20 != int32(1) {
				v55 = v20
				return base.B2i32(v55 != int32(0))
			} else {
				v26 = int32(1)
				v27 = F_time(m)
				mBase = m.M
				if v27-v19 < int64(11) {
					v55 = v26
					return base.B2i32(v55 != int32(0))
				} else {
					v33 = base.AtomicRmwXchg32(m, v8, int32(0), int32(1))
					if v33 != 0 {
						F_s_lock(m, v8, int32(_a_F_WalRcvRunning_0), int32(103), int32(_a_F_WalRcvRunning_1))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
							if v39 != int32(1) {
								v42 = int32(0)
								atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v8))), uint32(v42))
								v55 = v26
								return base.B2i32(v55 != int32(0))
							} else {
								v45 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v45
								atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v6)+1456)), uint32(v45))
								F_ConditionVariableBroadcast(m, v6+int32(12))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									v55 = v45
									return base.B2i32(v55 != int32(0))
								}
							}
						}
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
						if v39 != int32(1) {
							v42 = int32(0)
							atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v8))), uint32(v42))
							v55 = v26
							return base.B2i32(v55 != int32(0))
						} else {
							v45 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v45
							atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v6)+1456)), uint32(v45))
							F_ConditionVariableBroadcast(m, v6+int32(12))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								v55 = v45
								return base.B2i32(v55 != int32(0))
							}
						}
					}
				}
			}
		}
	} else {
		v19 = *(*int64)(unsafe.Add(mBase, uint32(v6)+24))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
		v21 = int32(0)
		atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v6)+1456)), uint32(v21))
		if v20 != int32(1) {
			v55 = v20
			return base.B2i32(v55 != int32(0))
		} else {
			v26 = int32(1)
			v27 = F_time(m)
			mBase = m.M
			if v27-v19 < int64(11) {
				v55 = v26
				return base.B2i32(v55 != int32(0))
			} else {
				v33 = base.AtomicRmwXchg32(m, v8, int32(0), int32(1))
				if v33 != 0 {
					F_s_lock(m, v8, int32(_a_F_WalRcvRunning_0), int32(103), int32(_a_F_WalRcvRunning_1))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
						if v39 != int32(1) {
							v42 = int32(0)
							atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v8))), uint32(v42))
							v55 = v26
							return base.B2i32(v55 != int32(0))
						} else {
							v45 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v45
							atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v6)+1456)), uint32(v45))
							F_ConditionVariableBroadcast(m, v6+int32(12))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								v55 = v45
								return base.B2i32(v55 != int32(0))
							}
						}
					}
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
					if v39 != int32(1) {
						v42 = int32(0)
						atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v8))), uint32(v42))
						v55 = v26
						return base.B2i32(v55 != int32(0))
					} else {
						v45 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v45
						atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v6)+1456)), uint32(v45))
						F_ConditionVariableBroadcast(m, v6+int32(12))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							v55 = v45
							return base.B2i32(v55 != int32(0))
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
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v150 int32
	_ = v150
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v185 int32
	_ = v185
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v220 int32
	_ = v220
	var v234 int32
	_ = v234
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v255 int32
	_ = v255
	var v269 int32
	_ = v269
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v290 int32
	_ = v290
	var v304 int32
	_ = v304
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v325 int32
	_ = v325
	var v339 int32
	_ = v339
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v400 int32
	_ = v400
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v505 int32
	_ = v505
	var v512 int64
	_ = v512
	var v513 int32
	_ = v513
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v533 int32
	_ = v533
	var v540 int64
	_ = v540
	var v543 int64
	_ = v543
	var v544 int64
	_ = v544
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v557 int64
	_ = v557
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v564 int64
	_ = v564
	var v570 int64
	_ = v570
	var v574 int32
	_ = v574
	var v576 int64
	_ = v576
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v586 int32
	_ = v586
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v625 int64
	_ = v625
	var v626 int32
	_ = v626
	var v629 int64
	_ = v629
	var v630 int64
	_ = v630
	var v631 int32
	_ = v631
	var v636 int32
	_ = v636
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v678 int64
	_ = v678
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v687 int64
	_ = v687
	var v690 int64
	_ = v690
	var v694 int64
	_ = v694
	var v695 int64
	_ = v695
	var v698 int64
	_ = v698
	var v701 int32
	_ = v701
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v728 int32
	_ = v728
	var v733 int32
	_ = v733
	var v734 int64
	_ = v734
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v748 int32
	_ = v748
	var v753 int32
	_ = v753
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v771 int32
	_ = v771
	var v776 int32
	_ = v776
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v795 int32
	_ = v795
	var v836 int64
	_ = v836
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v841 int32
	_ = v841
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v851 int64
	_ = v851
	var v852 int32
	_ = v852
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v863 int32
	_ = v863
	var v868 int64
	_ = v868
	var v874 int32
	_ = v874
	var v881 int32
	_ = v881
	var v883 int64
	_ = v883
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v894 int32
	_ = v894
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v911 int32
	_ = v911
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v940 int32
	_ = v940
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v952 int32
	_ = v952
	var v957 int32
	_ = v957
	var v963 int32
	_ = v963
	var v969 int32
	_ = v969
	var v976 int32
	_ = v976
	var v982 int32
	_ = v982
	var v985 int32
	_ = v985
	var v987 int32
	_ = v987
	var v989 int32
	_ = v989
	var v992 int64
	_ = v992
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v1024 int64
	_ = v1024
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1041 int32
	_ = v1041
	var v1044 int32
	_ = v1044
	var v1048 int32
	_ = v1048
	var v1059 int64
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1091 int32
	_ = v1091
	var v1093 int32
	_ = v1093
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1130 int64
	_ = v1130
	var v1132 int32
	_ = v1132
	var v1134 int32
	_ = v1134
	var v1137 int32
	_ = v1137
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1184 int64
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1217 int32
	_ = v1217
	var v1219 int32
	_ = v1219
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1256 int64
	_ = v1256
	var v1258 int32
	_ = v1258
	var v1260 int32
	_ = v1260
	var v1263 int32
	_ = v1263
	var v1275 int64
	_ = v1275
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1311 int32
	_ = v1311
	var v1313 int64
	_ = v1313
	var v1342 int64
	_ = v1342
	var v1353 int32
	_ = v1353
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1362 int64
	_ = v1362
	var v1366 int64
	_ = v1366
	var v1367 int64
	_ = v1367
	var v1372 int64
	_ = v1372
	var v1378 int32
	_ = v1378
	var v1385 int32
	_ = v1385
	var v1387 int64
	_ = v1387
	var v1411 int64
	_ = v1411
	var v1414 int64
	_ = v1414
	var v1418 int64
	_ = v1418
	var v1422 int32
	_ = v1422
	var v1427 int32
	_ = v1427
	var v1428 int32
	_ = v1428
	var v1435 int32
	_ = v1435
	var v1440 int64
	_ = v1440
	var v1446 int32
	_ = v1446
	var v1453 int32
	_ = v1453
	var v1456 int32
	_ = v1456
	var v1486 int32
	_ = v1486
	var v1487 int64
	_ = v1487
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1493 int32
	_ = v1493
	var v1495 int32
	_ = v1495
	var v1502 int32
	_ = v1502
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1513 int32
	_ = v1513
	var v1516 int32
	_ = v1516
	var v1517 int64
	_ = v1517
	var v1519 int32
	_ = v1519
	var v1521 int32
	_ = v1521
	var v1526 int32
	_ = v1526
	var v1531 int32
	_ = v1531
	var v1533 int32
	_ = v1533
	var v1536 int32
	_ = v1536
	var v1537 int32
	_ = v1537
	var v1539 int32
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1552 int32
	_ = v1552
	var v1554 int32
	_ = v1554
	var v1555 int32
	_ = v1555
	var v1565 int32
	_ = v1565
	var v1566 int32
	_ = v1566
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1571 int32
	_ = v1571
	var v1577 int32
	_ = v1577
	var v1610 int32
	_ = v1610
	var v1616 int32
	_ = v1616
	var v1619 int32
	_ = v1619
	var v1621 int32
	_ = v1621
	var v1622 int32
	_ = v1622
	var v1624 int32
	_ = v1624
	var v1628 int32
	_ = v1628
	var v1631 int32
	_ = v1631
	var v1632 int32
	_ = v1632
	var v1639 int32
	_ = v1639
	var v1640 int32
	_ = v1640
	var v1641 int32
	_ = v1641
	var v1644 int32
	_ = v1644
	var v1651 int32
	_ = v1651
	var v1652 int32
	_ = v1652
	var v1654 int32
	_ = v1654
	var v1659 int32
	_ = v1659
	var v1666 int32
	_ = v1666
	var v1667 int32
	_ = v1667
	var v1672 int32
	_ = v1672
	var v1676 int32
	_ = v1676
	var v1677 int32
	_ = v1677
	var v1683 int64
	_ = v1683
	var v1688 int32
	_ = v1688
	var v1692 int32
	_ = v1692
	var v1694 int32
	_ = v1694
	var v1700 int32
	_ = v1700
	var v1703 int32
	_ = v1703
	var v1705 int32
	_ = v1705
	var v1711 int32
	_ = v1711
	var v1715 int32
	_ = v1715
	var v1717 int32
	_ = v1717
	var v1720 int32
	_ = v1720
	var v1724 int32
	_ = v1724
	var v1729 int32
	_ = v1729
	var v1730 int32
	_ = v1730
	var v1731 int32
	_ = v1731
	var v1734 int32
	_ = v1734
	var v1738 int32
	_ = v1738
	var v1743 int32
	_ = v1743
	var v1744 int32
	_ = v1744
	var v1745 int32
	_ = v1745
	var v1748 int32
	_ = v1748
	var v1752 int32
	_ = v1752
	var v1759 int32
	_ = v1759
	var v1762 int32
	_ = v1762
	var v1770 int32
	_ = v1770
	var v1771 int32
	_ = v1771
	var v1775 int32
	_ = v1775
	var v1776 int32
	_ = v1776
	var v1777 int32
	_ = v1777
	var v1782 int64
	_ = v1782
	var v1783 int64
	_ = v1783
	var v1791 int32
	_ = v1791
	var v1792 int32
	_ = v1792
	var v1795 int32
	_ = v1795
	var v1828 int32
	_ = v1828
	var v1829 int32
	_ = v1829
	var v1831 int32
	_ = v1831
	var v1834 int32
	_ = v1834
	var v1837 int32
	_ = v1837
	var v1842 int32
	_ = v1842
	var v1845 int32
	_ = v1845
	var v1850 int32
	_ = v1850
	var v1852 int32
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1855 int32
	_ = v1855
	var v1859 int32
	_ = v1859
	var v1860 int32
	_ = v1860
	var v1866 int64
	_ = v1866
	var v1871 int32
	_ = v1871
	var v1875 int32
	_ = v1875
	var v1877 int32
	_ = v1877
	var v1883 int32
	_ = v1883
	var v1886 int32
	_ = v1886
	var v1888 int32
	_ = v1888
	var v1894 int32
	_ = v1894
	var v1898 int32
	_ = v1898
	var v1900 int32
	_ = v1900
	var v1903 int32
	_ = v1903
	var v1907 int32
	_ = v1907
	var v1912 int32
	_ = v1912
	var v1913 int32
	_ = v1913
	var v1914 int32
	_ = v1914
	var v1917 int32
	_ = v1917
	var v1921 int32
	_ = v1921
	var v1928 int32
	_ = v1928
	var v1931 int32
	_ = v1931
	var v1939 int32
	_ = v1939
	var v1940 int32
	_ = v1940
	var v1944 int32
	_ = v1944
	var v1945 int32
	_ = v1945
	var v1946 int32
	_ = v1946
	var v1951 int64
	_ = v1951
	var v1952 int64
	_ = v1952
	var v1960 int32
	_ = v1960
	var v1961 int32
	_ = v1961
	var v1964 int32
	_ = v1964
	var v1997 int32
	_ = v1997
	var v1998 int32
	_ = v1998
	var v2000 int32
	_ = v2000
	var v2003 int32
	_ = v2003
	var v2006 int32
	_ = v2006
	var v2011 int32
	_ = v2011
	var v2014 int32
	_ = v2014
	var v2019 int32
	_ = v2019
	var v2021 int32
	_ = v2021
	var v2022 int32
	_ = v2022
	var v2054 int32
	_ = v2054
	var v2055 int32
	_ = v2055
	var v2056 int32
	_ = v2056
	var v2059 int32
	_ = v2059
	var v2092 int32
	_ = v2092
	var v2094 int32
	_ = v2094
	var v2096 int32
	_ = v2096
	var v2098 int32
	_ = v2098
	var v2099 int32
	_ = v2099
	var v2101 int32
	_ = v2101
	var v2102 int32
	_ = v2102
	var v2106 int32
	_ = v2106
	var v2107 int32
	_ = v2107
	var v2111 int32
	_ = v2111
	var v2112 int32
	_ = v2112
	var v2114 int64
	_ = v2114
	var v2116 int32
	_ = v2116
	var v2118 int32
	_ = v2118
	var v2126 int32
	_ = v2126
	var v2129 int32
	_ = v2129
	var v2134 int32
	_ = v2134
	var v2136 int32
	_ = v2136
	var v2139 int32
	_ = v2139
	var v2140 int32
	_ = v2140
	var v2141 int32
	_ = v2141
	var v2144 int32
	_ = v2144
	var v2174 int64
	_ = v2174
	var v2178 int32
	_ = v2178
	var v2182 int32
	_ = v2182
	var v2183 int32
	_ = v2183
	var v2185 int32
	_ = v2185
	var v2190 int32
	_ = v2190
	var v2194 int32
	_ = v2194
	var v2197 int64
	_ = v2197
	var v2202 int32
	_ = v2202
	var v2207 int32
	_ = v2207
	var v2208 int32
	_ = v2208
	var v2209 int32
	_ = v2209
	var v2239 int32
	_ = v2239
	var v2246 int32
	_ = v2246
	var v2247 int32
	_ = v2247
	var v2248 int64
	_ = v2248
	var v2249 int64
	_ = v2249
	var v2253 int64
	_ = v2253
	var v2254 int64
	_ = v2254
	var v2258 int64
	_ = v2258
	var v2265 int32
	_ = v2265
	var v2272 int32
	_ = v2272
	var v2275 int64
	_ = v2275
	var v2276 int32
	_ = v2276
	var v2301 int64
	_ = v2301
	var v2307 int32
	_ = v2307
	var v2331 int64
	_ = v2331
	var v2336 int32
	_ = v2336
	var v2340 int32
	_ = v2340
	var v2344 int32
	_ = v2344
	var v2345 int32
	_ = v2345
	var v2352 int32
	_ = v2352
	var v2356 int32
	_ = v2356
	var v2357 int32
	_ = v2357
	var v2360 int32
	_ = v2360
	var v2362 int64
	_ = v2362
	var v2364 int32
	_ = v2364
	var v2366 int32
	_ = v2366
	var v2370 int32
	_ = v2370
	var v2379 int32
	_ = v2379
	var v2380 int32
	_ = v2380
	var v2386 int32
	_ = v2386
	var v2387 int32
	_ = v2387
	var v2396 int32
	_ = v2396
	var v2400 int32
	_ = v2400
	var v2408 int32
	_ = v2408
	var v2415 int32
	_ = v2415
	var v2418 int32
	_ = v2418
	var v2422 int32
	_ = v2422
	var v2426 int32
	_ = v2426
	var v2427 int64
	_ = v2427
	var v2430 int32
	_ = v2430
	var v2431 int32
	_ = v2431
	var v2434 int32
	_ = v2434
	var v2443 int32
	_ = v2443
	var v2450 int32
	_ = v2450
	var v2460 int32
	_ = v2460
	var v2467 int32
	_ = v2467
	var v2470 int32
	_ = v2470
	var v2471 int32
	_ = v2471
	var v2473 int32
	_ = v2473
	var v2487 int32
	_ = v2487
	var v2491 int32
	_ = v2491
	var v2492 int32
	_ = v2492
	var v2494 int32
	_ = v2494
	var v2497 int32
	_ = v2497
	var v2499 int32
	_ = v2499
	var v2504 int32
	_ = v2504
	var v2505 int32
	_ = v2505
	var v2506 int32
	_ = v2506
	var v2507 int64
	_ = v2507
	var v2510 int32
	_ = v2510
	var v2517 int32
	_ = v2517
	var v2545 int32
	_ = v2545
	var v2549 int32
	_ = v2549
	var v2552 int32
	_ = v2552
	var v2583 int32
	_ = v2583
	var v2589 int32
	_ = v2589
	var v2591 int32
	_ = v2591
	var v2594 int32
	_ = v2594
	var v2599 int32
	_ = v2599
	var v2620 int32
	_ = v2620
	var v2621 int32
	_ = v2621
	var v2623 int32
	_ = v2623
	var v2647 int32
	_ = v2647
	var v2648 int32
	_ = v2648
	var v2649 int32
	_ = v2649
	var v2653 int32
	_ = v2653
	var v2654 int32
	_ = v2654
	var v2657 int32
	_ = v2657
	var v2658 int32
	_ = v2658
	var v2659 int32
	_ = v2659
	var v2660 int32
	_ = v2660
	var v2665 int32
	_ = v2665
	var v2666 int32
	_ = v2666
	var v2668 int64
	_ = v2668
	var v2670 int32
	_ = v2670
	var v2672 int32
	_ = v2672
	var v2674 int32
	_ = v2674
	var v2680 int32
	_ = v2680
	var v2708 int32
	_ = v2708
	var v2714 int32
	_ = v2714
	var v2719 int32
	_ = v2719
	var v2723 int32
	_ = v2723
	var v2724 int32
	_ = v2724
	var v2725 int32
	_ = v2725
	var v2733 int32
	_ = v2733
	var v2759 int32
	_ = v2759
	var v2760 int32
	_ = v2760
	var v2762 int32
	_ = v2762
	var v2764 int32
	_ = v2764
	var v2766 int32
	_ = v2766
	var v2771 int32
	_ = v2771
	var v2772 int32
	_ = v2772
	var v2773 int32
	_ = v2773
	var v2774 int32
	_ = v2774
	var v2775 int32
	_ = v2775
	var v2778 int32
	_ = v2778
	var v2779 int32
	_ = v2779
	var v2780 int64
	_ = v2780
	var v2782 int64
	_ = v2782
	var v2784 int64
	_ = v2784
	var v2786 int32
	_ = v2786
	var v2790 int32
	_ = v2790
	var v2791 int32
	_ = v2791
	var v2792 int32
	_ = v2792
	var v2794 int64
	_ = v2794
	var v2798 int32
	_ = v2798
	var v2799 int32
	_ = v2799
	var v2805 int32
	_ = v2805
	var v2812 int32
	_ = v2812
	var v2813 int32
	_ = v2813
	var v2814 int32
	_ = v2814
	var v2815 int32
	_ = v2815
	var v2816 int32
	_ = v2816
	var v2818 int32
	_ = v2818
	var v2819 int32
	_ = v2819
	var v2820 int32
	_ = v2820
	var v2822 int32
	_ = v2822
	var v2823 int32
	_ = v2823
	var v2825 int32
	_ = v2825
	var v2827 int32
	_ = v2827
	var v2831 int32
	_ = v2831
	var v2832 int32
	_ = v2832
	var v2833 int32
	_ = v2833
	var v2834 int32
	_ = v2834
	var v2838 int32
	_ = v2838
	var v2842 int32
	_ = v2842
	var v2846 int32
	_ = v2846
	var v2847 int32
	_ = v2847
	var v2848 int32
	_ = v2848
	var v2849 int32
	_ = v2849
	var v2853 int32
	_ = v2853
	var v2854 int32
	_ = v2854
	var v2855 int32
	_ = v2855
	var v2857 int32
	_ = v2857
	var v2868 int32
	_ = v2868
	var v2872 int32
	_ = v2872
	var v2873 int32
	_ = v2873
	var v2877 int32
	_ = v2877
	var v2878 int32
	_ = v2878
	var v2882 int32
	_ = v2882
	var v2883 int32
	_ = v2883
	var v2885 int32
	_ = v2885
	var v2887 int32
	_ = v2887
	var v2891 int32
	_ = v2891
	var v2892 int32
	_ = v2892
	var v2896 int32
	_ = v2896
	var v2897 int32
	_ = v2897
	var v2899 int32
	_ = v2899
	var v2900 int32
	_ = v2900
	var v2902 int32
	_ = v2902
	var v2906 int32
	_ = v2906
	var v2907 int32
	_ = v2907
	var v2911 int32
	_ = v2911
	var v2912 int32
	_ = v2912
	var v2914 int32
	_ = v2914
	var v2915 int32
	_ = v2915
	var v2916 int32
	_ = v2916
	var v2917 int32
	_ = v2917
	var v2918 int32
	_ = v2918
	var v2920 int32
	_ = v2920
	var v2921 int32
	_ = v2921
	var v2922 int32
	_ = v2922
	var v2924 int32
	_ = v2924
	var v2925 int32
	_ = v2925
	var v2927 int32
	_ = v2927
	var v2929 int32
	_ = v2929
	var v2933 int32
	_ = v2933
	var v2934 int32
	_ = v2934
	var v2935 int32
	_ = v2935
	var v2936 int32
	_ = v2936
	var v2940 int32
	_ = v2940
	var v2944 int32
	_ = v2944
	var v2948 int32
	_ = v2948
	var v2949 int32
	_ = v2949
	var v2950 int32
	_ = v2950
	var v2951 int32
	_ = v2951
	var v2955 int32
	_ = v2955
	var v2956 int32
	_ = v2956
	var v2957 int32
	_ = v2957
	var v2959 int32
	_ = v2959
	var v2970 int32
	_ = v2970
	var v2974 int32
	_ = v2974
	var v2975 int32
	_ = v2975
	var v2979 int32
	_ = v2979
	var v2980 int32
	_ = v2980
	var v2984 int32
	_ = v2984
	var v2985 int32
	_ = v2985
	var v2989 int32
	_ = v2989
	var v2990 int32
	_ = v2990
	var v2991 int32
	_ = v2991
	var v2995 int32
	_ = v2995
	var v2996 int32
	_ = v2996
	var v2997 int32
	_ = v2997
	var v3001 int32
	_ = v3001
	var v3002 int32
	_ = v3002
	var v3003 int32
	_ = v3003
	var v3005 int32
	_ = v3005
	var v3006 int32
	_ = v3006
	var v3007 int32
	_ = v3007
	var v3011 int32
	_ = v3011
	var v3012 int32
	_ = v3012
	var v3013 int32
	_ = v3013
	var v3014 int32
	_ = v3014
	var v3018 int32
	_ = v3018
	var v3019 int32
	_ = v3019
	var v3020 int32
	_ = v3020
	var v3021 int32
	_ = v3021
	var v3025 int32
	_ = v3025
	var v3026 int32
	_ = v3026
	var v3027 int32
	_ = v3027
	var v3028 int32
	_ = v3028
	var v3032 int32
	_ = v3032
	var v3033 int32
	_ = v3033
	var v3034 int32
	_ = v3034
	var v3037 int32
	_ = v3037
	var v3039 int32
	_ = v3039
	var v3043 int32
	_ = v3043
	var v3047 int32
	_ = v3047
	var v3051 int32
	_ = v3051
	var v3055 int32
	_ = v3055
	var v3059 int32
	_ = v3059
	var v3064 int32
	_ = v3064
	var v3065 int32
	_ = v3065
	var v3066 int32
	_ = v3066
	var v3069 int32
	_ = v3069
	var v3070 int32
	_ = v3070
	var v3079 int32
	_ = v3079
	var v3080 int32
	_ = v3080
	var v3104 int64
	_ = v3104
	var v3105 int64
	_ = v3105
	var v3107 int64
	_ = v3107
	var v3110 int64
	_ = v3110
	var v3117 int32
	_ = v3117
	var v3120 int32
	_ = v3120
	var v3121 int32
	_ = v3121
	var v3128 int32
	_ = v3128
	var v3153 int32
	_ = v3153
	var v3156 int32
	_ = v3156
	var v3157 int32
	_ = v3157
	var v3159 int32
	_ = v3159
	var v3160 int32
	_ = v3160
	var v3162 int32
	_ = v3162
	var v3166 int32
	_ = v3166
	var v3167 int32
	_ = v3167
	var v3168 int32
	_ = v3168
	var v3169 int32
	_ = v3169
	var v3170 int32
	_ = v3170
	var v3173 int32
	_ = v3173
	var v3176 int32
	_ = v3176
	var v3177 int32
	_ = v3177
	var v3178 int32
	_ = v3178
	var v3179 int32
	_ = v3179
	var v3182 int32
	_ = v3182
	var v3188 int32
	_ = v3188
	var v3196 int32
	_ = v3196
	var v3220 int32
	_ = v3220
	var v3224 int32
	_ = v3224
	var v3228 int32
	_ = v3228
	var v3229 int32
	_ = v3229
	var v3233 int32
	_ = v3233
	var v3235 int32
	_ = v3235
	var v3236 int32
	_ = v3236
	var v3238 int32
	_ = v3238
	var v3242 int32
	_ = v3242
	var v3243 int32
	_ = v3243
	var v3244 int32
	_ = v3244
	var v3245 int32
	_ = v3245
	var v3246 int32
	_ = v3246
	var v3249 int32
	_ = v3249
	var v3252 int32
	_ = v3252
	var v3253 int32
	_ = v3253
	var v3254 int32
	_ = v3254
	var v3255 int32
	_ = v3255
	var v3258 int32
	_ = v3258
	var v3265 int32
	_ = v3265
	var v3266 int32
	_ = v3266
	var v3299 int32
	_ = v3299
	var v3300 int32
	_ = v3300
	var v3301 int32
	_ = v3301
	var v3333 int32
	_ = v3333
	var v3335 int32
	_ = v3335
	var v3337 int64
	_ = v3337
	var v3344 int32
	_ = v3344
	var v3345 int32
	_ = v3345
	var v3348 int32
	_ = v3348
	var v3349 int32
	_ = v3349
	var v3351 int32
	_ = v3351
	var v3356 int32
	_ = v3356
	var v3359 int32
	_ = v3359
	var v3360 int32
	_ = v3360
	var v3361 int32
	_ = v3361
	var v3362 int32
	_ = v3362
	var v3365 int32
	_ = v3365
	var v3367 int32
	_ = v3367
	var v3368 int32
	_ = v3368
	var v3369 int64
	_ = v3369
	var v3371 int64
	_ = v3371
	var v3373 int64
	_ = v3373
	var v3375 int32
	_ = v3375
	var v3379 int32
	_ = v3379
	var v3383 int32
	_ = v3383
	var v3386 int32
	_ = v3386
	var v3388 int32
	_ = v3388
	var v3393 int32
	_ = v3393
	var v3394 int32
	_ = v3394
	var v3395 int32
	_ = v3395
	var v3396 int32
	_ = v3396
	var v3397 int32
	_ = v3397
	var v3400 int32
	_ = v3400
	var v3402 int32
	_ = v3402
	var v3404 int32
	_ = v3404
	var v3406 int32
	_ = v3406
	var v3408 int32
	_ = v3408
	var v3409 int32
	_ = v3409
	var v3410 int32
	_ = v3410
	var v3411 int32
	_ = v3411
	var v3422 int32
	_ = v3422
	var v3424 int32
	_ = v3424
	var v3429 int32
	_ = v3429
	var v3430 int32
	_ = v3430
	var v3444 int32
	_ = v3444
	var v3451 int32
	_ = v3451
	var v3459 int32
	_ = v3459
	var v3460 int32
	_ = v3460
	var v3500 int32
	_ = v3500
	var v3501 int32
	_ = v3501
	var v3507 int64
	_ = v3507
	var v3508 int64
	_ = v3508
	var v3513 int64
	_ = v3513
	var v3517 int32
	_ = v3517
	var v3524 int32
	_ = v3524
	var v3525 int32
	_ = v3525
	var v3530 int32
	_ = v3530
	var v3534 int32
	_ = v3534
	var v3535 int32
	_ = v3535
	var v3537 int32
	_ = v3537
	var v3539 int32
	_ = v3539
	var v3546 int32
	_ = v3546
	var v3550 int32
	_ = v3550
	var v3554 int32
	_ = v3554
	var v3558 int32
	_ = v3558
	var v3582 int64
	_ = v3582
	var v3583 int64
	_ = v3583
	var v3586 int64
	_ = v3586
	var v3593 int32
	_ = v3593
	var v3596 int32
	_ = v3596
	var v3610 int32
	_ = v3610
	var v3624 int32
	_ = v3624
	var v3625 int64
	_ = v3625
	var v3629 int32
	_ = v3629
	var v3631 int32
	_ = v3631
	var v3632 int32
	_ = v3632
	var v3635 int32
	_ = v3635
	var v3637 int32
	_ = v3637
	var v3639 int32
	_ = v3639
	var v3640 int32
	_ = v3640
	var v3641 int64
	_ = v3641
	var v3643 int32
	_ = v3643
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
	v3624 = int32(m.ExcTag)
	v3625 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v3624 == int32(0) {
		goto L555
	} else {
		goto L556
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
		v3596 = v40
		v3610 = v54
		goto L6
	} else {
		goto L10
	}
L8:
	;
	v426 = v41
	v427 = v43
	goto L9
L9:
	;
	if v427 != 0 {
		goto L73
	} else {
		goto L74
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v41
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	v83 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		v3596 = v40
		v3610 = v54
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
		v3596 = v40
		v3610 = v54
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
		v3596 = v40
		v3610 = v54
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
	v136 = int32(916)
	v138 = m.G0
	v140 = v138 - int32(32)
	m.G0 = v140
	switch int32(918) {
	case 0, 2:
		v150 = v136
		goto L24
	default:
		goto L25
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
	v129 = F___sigaction(m, int32(1), v105+int32(12), int32(0))
	mBase = m.M
	m.G0 = v105 + int32(32)
	goto L17
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v41
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	v171 = int32(916)
	v173 = m.G0
	v175 = v173 - int32(32)
	m.G0 = v175
	switch int32(918) {
	case 0, 2:
		v185 = v171
		goto L30
	default:
		goto L31
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v140)+12)) = v150
	F_sigemptyset(m, v140+int32(16))
	mBase = m.M
	goto L27
L25:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[2])) = v136
	v150 = int32(_a_F_WalSummarizerMain_3)
	goto L24
L27:
	;
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v140)+24)) = int32(268435456)
	v164 = F___sigaction(m, int32(2), v140+int32(12), int32(0))
	mBase = m.M
	m.G0 = v140 + int32(32)
	goto L23
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v41
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	v206 = int32(-2)
	v208 = m.G0
	v210 = v208 - int32(32)
	m.G0 = v210
	switch int32(0) {
	case 0, 2:
		v220 = v206
		goto L36
	default:
		goto L37
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v175)+12)) = v185
	F_sigemptyset(m, v175+int32(16))
	mBase = m.M
	goto L33
L31:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[3])) = v171
	v185 = int32(_a_F_WalSummarizerMain_3)
	goto L30
L33:
	;
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v175)+24)) = int32(268435456)
	v199 = F___sigaction(m, int32(15), v175+int32(12), int32(0))
	mBase = m.M
	m.G0 = v175 + int32(32)
	goto L29
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v41
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	v241 = int32(-2)
	v243 = m.G0
	v245 = v243 - int32(32)
	m.G0 = v245
	switch int32(0) {
	case 0, 2:
		v255 = v241
		goto L42
	default:
		goto L43
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v210)+12)) = v220
	F_sigemptyset(m, v210+int32(16))
	mBase = m.M
	goto L39
L37:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[4])) = v206
	v220 = int32(_a_F_WalSummarizerMain_3)
	goto L36
L39:
	;
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v210)+24)) = int32(268435456)
	v234 = F___sigaction(m, int32(14), v210+int32(12), int32(0))
	mBase = m.M
	m.G0 = v210 + int32(32)
	goto L35
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v41
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	v276 = int32(917)
	v278 = m.G0
	v280 = v278 - int32(32)
	m.G0 = v280
	switch int32(919) {
	case 0, 2:
		v290 = v276
		goto L48
	default:
		goto L49
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v245)+12)) = v255
	F_sigemptyset(m, v245+int32(16))
	mBase = m.M
	goto L45
L43:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[5])) = v241
	v255 = int32(_a_F_WalSummarizerMain_3)
	goto L42
L45:
	;
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v245)+24)) = int32(268435456)
	v269 = F___sigaction(m, int32(13), v245+int32(12), int32(0))
	mBase = m.M
	m.G0 = v245 + int32(32)
	goto L41
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v41
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	v311 = int32(-2)
	v313 = m.G0
	v315 = v313 - int32(32)
	m.G0 = v315
	switch int32(0) {
	case 0, 2:
		v325 = v311
		goto L54
	default:
		goto L55
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v280)+12)) = v290
	F_sigemptyset(m, v280+int32(16))
	mBase = m.M
	goto L51
L49:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[6])) = v276
	v290 = int32(_a_F_WalSummarizerMain_3)
	goto L48
L51:
	;
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v280)+24)) = int32(268435456)
	v304 = F___sigaction(m, int32(10), v280+int32(12), int32(0))
	mBase = m.M
	m.G0 = v280 + int32(32)
	goto L47
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v41
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	F_on_shmem_exit(m, int32(965), int32(0))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		v3596 = v40
		v3610 = v54
		goto L6
	} else {
		goto L59
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v315)+12)) = v325
	F_sigemptyset(m, v315+int32(16))
	mBase = m.M
	goto L57
L55:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[7])) = v311
	v325 = int32(_a_F_WalSummarizerMain_3)
	goto L54
L57:
	;
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v315)+24)) = int32(268435456)
	v339 = F___sigaction(m, int32(12), v315+int32(12), int32(0))
	mBase = m.M
	m.G0 = v315 + int32(32)
	goto L53
L59:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v41
	v352 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[8]))
	v356 = F_LWLockAcquire(m, v352+int32(_a_F_WalSummarizerMain_4), int32(0))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		v3596 = v40
		v3610 = v54
		goto L6
	} else {
		goto L60
	}
L60:
	;
	v359 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[9]))
	v361 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[10]))
	*(*int32)(unsafe.Add(mBase, uint32(v359)+20)) = v361
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v41
	v366 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[8]))
	F_LWLockRelease(m, v366+int32(_a_F_WalSummarizerMain_4))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		v3596 = v40
		v3610 = v54
		goto L6
	} else {
		goto L61
	}
L61:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v41
	v375 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[11]))
	v380 = F_AllocSetContextCreateInternal(m, v375, int32(_a_F_WalSummarizerMain_5), int32(0), int32(_a_F_WalSummarizerMain_6), int32(_a_F_WalSummarizerMain_7))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		v3596 = v40
		v3610 = v54
		goto L6
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[12])) = v380
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v380
	v386 = int32(0)
	v388 = m.G0
	v390 = v388 - int32(32)
	m.G0 = v390
	switch int32(2) {
	case 0, 2:
		v400 = v386
		goto L64
	default:
		goto L65
	}
L63:
	;
	goto L69
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v390)+12)) = v400
	F_sigemptyset(m, v390+int32(16))
	mBase = m.M
	goto L66
L65:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[13])) = v386
	v400 = int32(_a_F_WalSummarizerMain_3)
	goto L64
L66:
	;
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v390)+24)) = int32(268435457)
	v414 = F___sigaction(m, int32(17), v390+int32(12), int32(0))
	mBase = m.M
	m.G0 = v390 + int32(32)
	goto L63
L69:
	;
	v419 = v40 + int32(272)
	*(*int32)(unsafe.Add(mBase, uint32(v419)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v419))) = v40 + int32(252)
	goto L72
L70:
	;
	v426 = v380
	v427 = int32(0)
	goto L9
L72:
	;
	goto L70
L73:
	;
	v428 = int32(_a_F_WalSummarizerMain_8)
	v430 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[14]))
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[14])) = v430 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[15])) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v426
	F_EmitErrorReport(m)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		v3596 = v40
		v3610 = v54
		goto L6
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[16])) = v40 + int32(272)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v426
	F_pgmem_sigprocmask(m, int32(_a_F_WalSummarizerMain_9), int32(0))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		v3596 = v40
		v3610 = v54
		goto L6
	} else {
		goto L86
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v426
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	F_LWLockReleaseAll(m)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		v3596 = v40
		v3610 = v54
		goto L6
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v426
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		v3596 = v40
		v3610 = v54
		goto L6
	} else {
		goto L78
	}
L78:
	;
	v450 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[17]))
	*(*int32)(unsafe.Add(mBase, uint32(v450))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v426
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	F_pgaio_error_cleanup(m)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		v3596 = v40
		v3610 = v54
		goto L6
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v426
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	F_ReleaseAuxProcessResources(m, int32(0))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		v3596 = v40
		v3610 = v54
		goto L6
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v426
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	F_AtEOXact_Files(m, int32(0))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		v3596 = v40
		v3610 = v54
		goto L6
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v426
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	F_AtEOXact_HashTables(m, int32(0))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		v3596 = v40
		v3610 = v54
		goto L6
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[12])) = v426
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v426
	F_FlushErrorState(m)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		v3596 = v40
		v3610 = v54
		goto L6
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v426
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	F_MemoryContextReset(m, v426)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		v3596 = v40
		v3610 = v54
		goto L6
	} else {
		goto L84
	}
L84:
	;
	v482 = int32(_a_F_WalSummarizerMain_8)
	v484 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[14]))
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[14])) = v484 - int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v426
	v494 = F_WaitLatch(m, int32(0), int32(40), int32(_a_F_WalSummarizerMain_10), int32(150994953))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		v3596 = v40
		v3610 = v54
		goto L6
	} else {
		goto L85
	}
L85:
	;
	goto L75
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v426
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	v512 = F_GetOldestUnsummarizedLSN(m, v40+int32(268), v40+int32(267))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		v3596 = v40
		v3610 = v54
		goto L6
	} else {
		goto L87
	}
L87:
	;
	if v512 != int64(0) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v519 = v40
	v520 = v426
	v533 = v54
	v540 = v61
	v543 = v512
	v544 = int64(0)
	goto L91
L89:
	;
	goto L90
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+3132)) = v426
	*(*int64)(unsafe.Add(mBase, uint32(v40)+3120)) = v61
	F_proc_exit(m, int32(0))
	mBase = m.M
	v3593 = m.ExcPending
	if v3593 != 0 {
		v3596 = v40
		v3610 = v54
		goto L6
	} else {
		goto L554
	}
L91:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v540
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	F_MemoryContextReset(m, v520)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v540
	F_ProcessWalSummarizerInterrupts(m)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v540
	v557 = F_GetRedoRecPtr(m)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L95
	}
L95:
	;
	v560 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[18]))
	if v560 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v540
	v836 = F_GetLatestLSN(m, v519+int32(256))
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L147
	}
L97:
	;
	v564 = *(*int64)(unsafe.Add(mBase, _c_F_WalSummarizerMain[19]))
	if v557 == v564 {
		goto L96
	} else {
		goto L98
	}
L98:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_WalSummarizerMain[19])) = v557
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v540
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	v570 = F_time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v540
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	v574 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[18]))
	v576 = int64(0)
	v578 = F_GetWalSummaries(m, int32(0), v576, v576)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L99
	}
L99:
	;
	if v578 == int32(0) {
		goto L96
	} else {
		goto L100
	}
L100:
	;
	v586 = v578
	goto L101
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v540
	F_ProcessWalSummarizerInterrupts(m)
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L103
	}
L102:
	;
	goto L96
L103:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v586)+12))
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v620)))
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v621)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v540
	v625 = F_XLogGetOldestSegno(m, v622)
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L104
	}
L104:
	;
	v629 = int64(*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[20])))
	v630 = v625 * v629
	v631 = v586
	v636 = int32(0)
	goto L105
L105:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v631)+4))
	if v636 < v661 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	if v795 != 0 {
		v586 = v795
		goto L101
	} else {
		goto L146
	}
L107:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v631)+12))
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v663+v636<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v540
	F_ProcessWalSummarizerInterrupts(m)
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L110
	}
L108:
	;
	v795 = v631
	goto L109
L109:
	;
	goto L106
L110:
	;
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v667)+16))
	if v672 != v622 {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	if v789 != 0 {
		v631 = v789
		v636 = v791
		goto L105
	} else {
		goto L145
	}
L112:
	;
	v789 = v631
	v791 = v636 + int32(1)
	goto L111
L113:
	;
	goto L114
L114:
	;
	if v630 != int64(0) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v540
	v783 = F_list_delete_nth_cell(m, v631, v636)
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L143
	}
L116:
	;
	v678 = *(*int64)(unsafe.Add(mBase, uint32(v667)+8))
	if base.Ui64(v630) < base.Ui64(v678) {
		goto L115
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v540
	v682 = m.G0
	v684 = v682 - int32(1200)
	m.G0 = v684
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v667)+16))
	v687 = *(*int64)(unsafe.Add(mBase, uint32(v667)))
	v690 = *(*int64)(unsafe.Add(mBase, uint32(v667)+8))
	*(*uint32)(unsafe.Add(mBase, uint32(v684-int32(-64)))) = uint32(v690)
	*(*uint32)(unsafe.Add(mBase, uint32(v684)+56)) = uint32(v687)
	*(*int32)(unsafe.Add(mBase, uint32(v684)+48)) = v686
	v694 = int64(32)
	v695 = int64(base.Ui64(v690) >> (uint(v694) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v684)+60)) = uint32(v695)
	v698 = int64(base.Ui64(v687) >> (uint(v694) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v684)+52)) = uint32(v698)
	v701 = v684 + int32(176)
	v706 = F_pg_snprintf(m, v701, int32(1024), int32(_a_F_WalSummarizerMain_11), v684+int32(48))
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L120
	}
L119:
	;
	goto L118
L120:
	;
	v712 = F___fstatat(m, int32(-100), v701, v684+int32(80), int32(256))
	mBase = m.M
	goto L124
L121:
	;
	goto L115
L122:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L139
	}
L123:
	;
	m.G0 = v684 + int32(1200)
	goto L121
L124:
	;
	if v712 != 0 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v714 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[21]))
	if v714 == int32(44) {
		goto L123
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	v734 = *(*int64)(unsafe.Add(mBase, uint32(v684)+136))
	if v570-base.I64_extend_i32_s(v574*int32(60)) <= v734 {
		goto L123
	} else {
		goto L133
	}
L128:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L129
	}
L129:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L130
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v684)+32)) = v701
	F_errmsg(m, int32(_a_F_WalSummarizerMain_12), v684+int32(32))
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(_a_F_WalSummarizerMain_13), int32(247), int32(_a_F_WalSummarizerMain_14))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L132
	}
L132:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L133:
	;
	v737 = v684 + int32(176)
	v738 = F_unlink(m, v737)
	mBase = m.M
	if v738 != 0 {
		goto L122
	} else {
		goto L134
	}
L134:
	;
	v741 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L135
	}
L135:
	;
	if v741 == int32(0) {
		goto L123
	} else {
		goto L136
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v684))) = v737
	F_errmsg_internal(m, int32(_a_F_WalSummarizerMain_15), v684)
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(_a_F_WalSummarizerMain_13), int32(256), int32(_a_F_WalSummarizerMain_14))
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L138
	}
L138:
	;
	goto L123
L139:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v684)+16)) = v684 + int32(176)
	F_errmsg(m, int32(_a_F_WalSummarizerMain_16), v684+int32(16))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(_a_F_WalSummarizerMain_13), int32(254), int32(_a_F_WalSummarizerMain_14))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L142
	}
L142:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v540
	F_pfree(m, v667)
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L144
	}
L144:
	;
	v789 = v783
	v791 = v636
	goto L111
L145:
	;
	v795 = v789
	goto L109
L146:
	;
	goto L102
L147:
	;
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v519)+268))
	if v544 != int64(0) {
		v883 = v544
		goto L148
	} else {
		goto L149
	}
L148:
	;
	if base.Ui64(v883-int64(1)) < base.Ui64(v543) {
		goto L158
	} else {
		goto L159
	}
L149:
	;
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v519)+256))
	if v838 == v841 {
		v883 = v544
		goto L148
	} else {
		goto L150
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v540
	v845 = F_readTimeLineHistory(m, v841)
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L151
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v540
	v851 = F_tliSwitchPoint(m, v838, v845, v519+int32(260))
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v540
	v857 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L153
	}
L153:
	;
	if v857 == int32(0) {
		v883 = v851
		goto L148
	} else {
		goto L154
	}
L154:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v540
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v519)+260))
	*(*int32)(unsafe.Add(mBase, uint32(v519)+228)) = v863
	*(*int32)(unsafe.Add(mBase, uint32(v519)+224)) = v838
	*(*uint32)(unsafe.Add(mBase, uint32(v519)+236)) = uint32(v851)
	v868 = int64(base.Ui64(v851) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v519)+232)) = uint32(v868)
	F_errmsg_internal(m, int32(_a_F_WalSummarizerMain_17), v519+int32(224))
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L155
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v540
	F_errfinish(m, int32(_a_F_WalSummarizerMain_1), int32(389), int32(_a_F_WalSummarizerMain_2))
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L156
	}
L156:
	;
	v883 = v851
	goto L148
L157:
	;
	v540 = v3582
	v543 = v3586
	v544 = v3583
	goto L91
L158:
	;
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v519)+260))
	*(*int32)(unsafe.Add(mBase, uint32(v519)+268)) = v887
	v889 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v519)+260)) = v889
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v540
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	v894 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[8]))
	v898 = F_LWLockAcquire(m, v894+int32(_a_F_WalSummarizerMain_4), v889)
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L161
	}
L159:
	;
	goto L160
L160:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v540
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	v918 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519)+267)))
	v919 = F_CreateEmptyBlockRefTable(m)
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L163
	}
L161:
	;
	v901 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[9]))
	*(*int64)(unsafe.Add(mBase, uint32(v901)+24)) = v883
	v903 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v901)+16)) = uint8(v903)
	*(*int32)(unsafe.Add(mBase, uint32(v901)+4)) = v887
	*(*int64)(unsafe.Add(mBase, uint32(v901)+8)) = v883
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v540
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	v911 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[8]))
	F_LWLockRelease(m, v911+int32(_a_F_WalSummarizerMain_4))
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L162
	}
L162:
	;
	v3582 = v540
	v3583 = int64(0)
	v3586 = v883
	goto L157
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v540
	v924 = F_palloc0(m, int32(24))
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L164
	}
L164:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v924)+8)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v924))) = v838
	*(*uint8)(unsafe.Add(mBase, uint32(v924)+4)) = uint8(base.B2i32(v883 != int64(0)))
	*(*int32)(unsafe.Add(mBase, uint32(v519)+460)) = int32(394)
	*(*int32)(unsafe.Add(mBase, uint32(v519)+456)) = int32(395)
	*(*int32)(unsafe.Add(mBase, uint32(v519)+452)) = int32(966)
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v540
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	v940 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[20]))
	v943 = F_XLogReaderAllocate(m, v940, v519+int32(452), v924)
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L165
	}
L165:
	;
	if v943 == int32(0) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v540
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	if v918&int32(1) != 0 {
		goto L183
	} else {
		goto L184
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v540
	F_errcode(m, int32(_a_F_WalSummarizerMain_18))
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L170
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v540
	F_errmsg(m, int32(_a_F_WalSummarizerMain_19), int32(0))
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L171
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v540
	F_errdetail(m, int32(_a_F_WalSummarizerMain_20), int32(0))
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L172
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v540
	F_errfinish(m, int32(_a_F_WalSummarizerMain_1), int32(939), int32(_a_F_WalSummarizerMain_21))
	mBase = m.M
	v976 = m.ExcPending
	if v976 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L173
	}
L173:
	;
	goto L3
L174:
	;
	if (v2307^int32(-1)|v2345)&int32(1) != 0 {
		goto L545
	} else {
		goto L546
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	v2470 = int32(0)
	v2471 = m.G0
	v2473 = v2471 - int32(_a_F_WalSummarizerMain_22)
	m.G0 = v2473
	*(*int32)(unsafe.Add(mBase, uint32(v2473)+8)) = int32(1697321851)
	base.MemoryFill(m, v2473+int32(24), v2470, int32(_a_F_WalSummarizerMain_23))
	*(*int32)(unsafe.Add(mBase, uint32(v2473)+16)) = v519 + int32(464)
	*(*int32)(unsafe.Add(mBase, uint32(v2473)+12)) = int32(967)
	v2487 = int32(-1)
	v2491 = int32(4)
	v2492 = m.Env.Pgmem_crc32c(m, v2487, v2473+int32(8), v2491)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v2473)+uint32(_c_F_WalSummarizerMain[22]))) = v2492
	v2494 = *(*int32)(unsafe.Add(mBase, uint32(v919)))
	*(*int32)(unsafe.Add(mBase, uint32(v2473)+uint32(_c_F_WalSummarizerMain[23]))) = v2491
	v2497 = *(*int32)(unsafe.Add(mBase, uint32(v2473)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2473)+20)) = v2497
	v2499 = *(*int32)(unsafe.Add(mBase, uint32(v2494)+8))
	if v2499 == v2470 {
		goto L416
	} else {
		goto L417
	}
L176:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	v2418 = *(*int32)(unsafe.Add(mBase, uint32(v519)+448))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2422 = m.ExcPending
	if v2422 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L407
	}
L177:
	;
	v2336 = *(*int32)(unsafe.Add(mBase, uint32(v943)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	F_pfree(m, v2336)
	mBase = m.M
	v2340 = m.ExcPending
	if v2340 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L396
	}
L178:
	;
	v2307 = v2276
	v2331 = v2301
	goto L177
L179:
	;
	v2239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v924)+16)))
	if v2239 != int32(1) {
		goto L176
	} else {
		goto L389
	}
L180:
	;
	v1456 = int32(1)
	goto L247
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1342
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1435 = m.ExcPending
	if v1435 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L244
	}
L182:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	F_ProcessWalSummarizerInterrupts(m)
	mBase = m.M
	v1422 = m.ExcPending
	if v1422 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L241
	}
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v540
	F_XLogBeginRead(m, v943, v543)
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v540
	v985 = m.G0
	v987 = v985 - int32(16)
	m.G0 = v987
	v989 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v943)+1257)) = uint8(v989)
	v992 = v543 & int64(-8192)
	v996 = F_ReadPageInternal(m, v943, v992, base.I32_wrap_i64(v543)&int32(_a_F_WalSummarizerMain_24))
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L189
	}
L186:
	;
	v1411 = v540
	v1414 = v543
	v1418 = v883
	goto L182
L187:
	;
	m.G0 = v987 + int32(16)
	if v1342 != int64(0) {
		goto L231
	} else {
		goto L232
	}
L188:
	;
	v1311 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v943)+1192)) = v1311
	v1313 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v943)+1176)) = v1313
	*(*int32)(unsafe.Add(mBase, uint32(v943)+132)) = v1311
	v1342 = v1313
	goto L187
L189:
	;
	if v996 < int32(0) {
		goto L188
	} else {
		goto L190
	}
L190:
	;
	v1024 = v992
	goto L191
L191:
	;
	v1032 = *(*int32)(unsafe.Add(mBase, uint32(v943)+128))
	v1033 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1032)+2)))
	if v1033&int32(2) != 0 {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	goto L188
L193:
	;
	v1036 = int32(40)
	goto L195
L194:
	;
	v1036 = int32(24)
	goto L195
L195:
	;
	v1037 = F_ReadPageInternal(m, v943, v1024, v1036)
	mBase = m.M
	v1038 = m.ExcPending
	if v1038 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L196
	}
L196:
	;
	if v1037 < int32(0) {
		goto L188
	} else {
		goto L197
	}
L197:
	;
	v1041 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1032)+2)))
	if v1041&int32(1) != 0 {
		goto L200
	} else {
		goto L201
	}
L198:
	;
	v1275 = v1024 - int64(-8192)
	v1277 = F_ReadPageInternal(m, v943, v1275, int32(0))
	mBase = m.M
	v1278 = m.ExcPending
	if v1278 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L229
	}
L199:
	;
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(v943)+120))
	if v1060 != 0 {
		goto L204
	} else {
		goto L205
	}
L200:
	;
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(v1032)+16))
	v1048 = (v1044 + int32(7)) & int32(-8)
	if base.Ui32(int32(_a_F_WalSummarizerMain_6)-v1036) <= base.Ui32(v1048) {
		goto L198
	} else {
		goto L203
	}
L201:
	;
	goto L202
L202:
	;
	v1059 = v1024 | base.I64_extend_i32_u(v1036)
	goto L199
L203:
	;
	v1059 = base.I64_extend_i32_u(v1048) + (v1024 | base.I64_extend_i32_u(v1036))
	goto L199
L204:
	;
	v1061 = v1060
	goto L207
L205:
	;
	goto L206
L206:
	;
	v1130 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v943)+120)) = v1130
	v1132 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v943)+96)) = v1132
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v943)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v943)+116)) = v1134
	*(*int32)(unsafe.Add(mBase, uint32(v943)+112)) = v1134
	v1137 = *(*int32)(unsafe.Add(mBase, uint32(v943)+1252))
	*(*uint8)(unsafe.Add(mBase, uint32(v1137))) = uint8(v1132)
	*(*int64)(unsafe.Add(mBase, uint32(v943)+80)) = v1059
	*(*int64)(unsafe.Add(mBase, uint32(v943)+40)) = v1059
	*(*uint8)(unsafe.Add(mBase, uint32(v943)+1256)) = uint8(v1132)
	*(*int64)(unsafe.Add(mBase, uint32(v943)+72)) = v1130
	*(*int64)(unsafe.Add(mBase, uint32(v943)+32)) = v1130
	goto L214
L207:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v1061)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v943)+120)) = v1091
	v1093 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1061)+4)))
	if v1093 == int32(1) {
		goto L209
	} else {
		goto L210
	}
L208:
	;
	goto L206
L209:
	;
	F_pfree(m, v1061)
	mBase = m.M
	v1097 = m.ExcPending
	if v1097 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L212
	}
L210:
	;
	v1099 = v1091
	goto L211
L211:
	;
	if v1099 != 0 {
		v1061 = v1099
		goto L207
	} else {
		goto L213
	}
L212:
	;
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v943)+120))
	v1099 = v1098
	goto L211
L213:
	;
	goto L208
L214:
	;
	v1180 = F_XLogReadRecord(m, v943, v987+int32(12))
	mBase = m.M
	v1181 = m.ExcPending
	if v1181 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L216
	}
L215:
	;
	v1186 = *(*int32)(unsafe.Add(mBase, uint32(v943)+120))
	if v1186 != 0 {
		goto L219
	} else {
		goto L220
	}
L216:
	;
	if v1180 == int32(0) {
		goto L188
	} else {
		goto L217
	}
L217:
	;
	v1184 = *(*int64)(unsafe.Add(mBase, uint32(v943)+32))
	if base.Ui64(v1184) < base.Ui64(v543) {
		goto L214
	} else {
		goto L218
	}
L218:
	;
	goto L215
L219:
	;
	v1187 = v1186
	goto L222
L220:
	;
	goto L221
L221:
	;
	v1256 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v943)+120)) = v1256
	v1258 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v943)+96)) = v1258
	v1260 = *(*int32)(unsafe.Add(mBase, uint32(v943)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v943)+116)) = v1260
	*(*int32)(unsafe.Add(mBase, uint32(v943)+112)) = v1260
	v1263 = *(*int32)(unsafe.Add(mBase, uint32(v943)+1252))
	*(*uint8)(unsafe.Add(mBase, uint32(v1263))) = uint8(v1258)
	*(*int64)(unsafe.Add(mBase, uint32(v943)+80)) = v1184
	*(*int64)(unsafe.Add(mBase, uint32(v943)+40)) = v1184
	*(*uint8)(unsafe.Add(mBase, uint32(v943)+1256)) = uint8(v1258)
	*(*int64)(unsafe.Add(mBase, uint32(v943)+72)) = v1256
	*(*int64)(unsafe.Add(mBase, uint32(v943)+32)) = v1256
	v1342 = v1184
	goto L187
L222:
	;
	v1217 = *(*int32)(unsafe.Add(mBase, uint32(v1187)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v943)+120)) = v1217
	v1219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1187)+4)))
	if v1219 == int32(1) {
		goto L224
	} else {
		goto L225
	}
L223:
	;
	goto L221
L224:
	;
	F_pfree(m, v1187)
	mBase = m.M
	v1223 = m.ExcPending
	if v1223 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L227
	}
L225:
	;
	v1225 = v1217
	goto L226
L226:
	;
	if v1225 != 0 {
		v1187 = v1225
		goto L222
	} else {
		goto L228
	}
L227:
	;
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(v943)+120))
	v1225 = v1224
	goto L226
L228:
	;
	goto L223
L229:
	;
	if int32(0) <= v1277 {
		v1024 = v1275
		goto L191
	} else {
		goto L230
	}
L230:
	;
	goto L192
L231:
	;
	v1411 = v1342
	v1414 = v1342
	v1418 = v883
	goto L182
L232:
	;
	goto L233
L233:
	;
	v1353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v924)+16)))
	if v1353 != int32(1) {
		goto L181
	} else {
		goto L234
	}
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1342
	v1360 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1361 = m.ExcPending
	if v1361 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L235
	}
L235:
	;
	if v1360 != 0 {
		goto L236
	} else {
		goto L237
	}
L236:
	;
	v1362 = *(*int64)(unsafe.Add(mBase, uint32(v924)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1342
	*(*uint32)(unsafe.Add(mBase, uint32(v519)+192)) = uint32(v1362)
	v1366 = int64(32)
	v1367 = int64(base.Ui64(v1362) >> (uint(v1366) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v519)+188)) = uint32(v1367)
	*(*int32)(unsafe.Add(mBase, uint32(v519)+176)) = v838
	*(*uint32)(unsafe.Add(mBase, uint32(v519)+184)) = uint32(v543)
	v1372 = int64(base.Ui64(v543) >> (uint(v1366) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v519)+180)) = uint32(v1372)
	F_errmsg_internal(m, int32(_a_F_WalSummarizerMain_25), v519+int32(176))
	mBase = m.M
	v1378 = m.ExcPending
	if v1378 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L239
	}
L237:
	;
	goto L238
L238:
	;
	v1387 = *(*int64)(unsafe.Add(mBase, uint32(v943)+40))
	v1411 = v1342
	v1414 = v543
	v1418 = v1387
	goto L182
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1342
	F_errfinish(m, int32(_a_F_WalSummarizerMain_1), int32(987), int32(_a_F_WalSummarizerMain_21))
	mBase = m.M
	v1385 = m.ExcPending
	if v1385 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L240
	}
L240:
	;
	goto L238
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	v1427 = F_XLogReadRecord(m, v943, v519+int32(448))
	mBase = m.M
	v1428 = m.ExcPending
	if v1428 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L242
	}
L242:
	;
	if v1427 != 0 {
		goto L180
	} else {
		goto L243
	}
L243:
	;
	v2209 = int32(1)
	goto L179
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1342
	*(*uint32)(unsafe.Add(mBase, uint32(v519)+212)) = uint32(v543)
	v1440 = int64(base.Ui64(v543) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v519)+208)) = uint32(v1440)
	F_errmsg(m, int32(_a_F_WalSummarizerMain_26), v519+int32(208))
	mBase = m.M
	v1446 = m.ExcPending
	if v1446 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L245
	}
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1342
	F_errfinish(m, int32(_a_F_WalSummarizerMain_1), int32(1004), int32(_a_F_WalSummarizerMain_21))
	mBase = m.M
	v1453 = m.ExcPending
	if v1453 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L246
	}
L246:
	;
	goto L3
L247:
	;
	v1486 = base.B2i32(v1418 == int64(0))
	if v1418 == int64(0) {
		goto L249
	} else {
		goto L250
	}
L248:
	;
	v2209 = v2144
	goto L179
L249:
	;
	v1489 = *(*int32)(unsafe.Add(mBase, uint32(v943)+96))
	v1490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1489)+49)))
	if v1490 == int32(0) {
		goto L254
	} else {
		goto L255
	}
L250:
	;
	v1487 = *(*int64)(unsafe.Add(mBase, uint32(v943)+32))
	if base.Ui64(v1487) < base.Ui64(v1418) {
		goto L249
	} else {
		goto L251
	}
L251:
	;
	v2307 = v1456
	v2331 = v1418
	goto L177
L252:
	;
	v2174 = *(*int64)(unsafe.Add(mBase, uint32(v943)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	v2178 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[8]))
	v2182 = F_LWLockAcquire(m, v2178+int32(_a_F_WalSummarizerMain_4), int32(0))
	mBase = m.M
	v2183 = m.ExcPending
	if v2183 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L380
	}
L253:
	;
	v2054 = int32(0)
	v2055 = *(*int32)(unsafe.Add(mBase, uint32(v943)+96))
	v2056 = *(*int32)(unsafe.Add(mBase, uint32(v2055)+72))
	if v2056 < v2054 {
		v2144 = v2054
		goto L252
	} else {
		goto L358
	}
L254:
	;
	v1493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1489)+48)))
	v1495 = v1493 & int32(240)
	if base.Ui32(v1495) <= base.Ui32(int32(143)) {
		goto L261
	} else {
		goto L262
	}
L255:
	;
	goto L256
L256:
	;
	v1526 = int32(1)
	if v1456&v1526 != 0 {
		v2144 = v1526
		goto L252
	} else {
		goto L271
	}
L257:
	;
	v1521 = int32(1)
	if v1456&v1521 == int32(0) {
		goto L253
	} else {
		goto L270
	}
L258:
	;
	v1517 = *(*int64)(unsafe.Add(mBase, uint32(v943)+32))
	if base.Ui64(v1414) < base.Ui64(v1517) {
		v2307 = v1456
		v2331 = v1517
		goto L177
	} else {
		goto L268
	}
L259:
	;
	v1513 = *(*int32)(unsafe.Add(mBase, uint32(v1489)+64))
	v1516 = v1513 + int32(16)
	goto L258
L260:
	;
	v1510 = *(*int32)(unsafe.Add(mBase, uint32(v1489)+64))
	v1516 = v1510 + int32(20)
	goto L258
L261:
	;
	if v1495 == int32(0) {
		goto L260
	} else {
		goto L264
	}
L262:
	;
	goto L263
L263:
	;
	if v1495 == int32(144) {
		goto L259
	} else {
		goto L266
	}
L264:
	;
	if v1495 != int32(96) {
		goto L257
	} else {
		goto L265
	}
L265:
	;
	v1502 = *(*int32)(unsafe.Add(mBase, uint32(v1489)+64))
	v1516 = v1502 + int32(20)
	goto L258
L266:
	;
	if v1495 != int32(224) {
		goto L257
	} else {
		goto L267
	}
L267:
	;
	v1509 = *(*int32)(unsafe.Add(mBase, uint32(v1489)+64))
	v1516 = v1509
	goto L258
L268:
	;
	v1519 = *(*int32)(unsafe.Add(mBase, uint32(v1516)))
	if v1519 != 0 {
		goto L253
	} else {
		goto L269
	}
L269:
	;
	v2144 = int32(1)
	goto L252
L270:
	;
	v2144 = v1521
	goto L252
L271:
	;
	switch v1490 - int32(1) {
	case 0:
		goto L272
	case 1:
		goto L273
	default:
		goto L253
	case 3:
		goto L274
	}
L272:
	;
	v1667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1489)+48)))
	switch int32(base.Ui32(v1667)>>(uint(int32(4))%32)) & int32(7) {
	case 0, 3:
		goto L302
	default:
		goto L253
	case 2, 4:
		goto L301
	}
L273:
	;
	v1624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1489)+48)))
	v1628 = v1624&int32(240) - int32(16)
	if v1628 != 0 {
		goto L287
	} else {
		goto L288
	}
L274:
	;
	v1531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1489)+48)))
	v1533 = v1531 & int32(240)
	switch v1533 - int32(16) {
	case 0:
		goto L276
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
		goto L253
	case 16:
		goto L275
	default:
		goto L277
	}
L275:
	;
	v1566 = *(*int32)(unsafe.Add(mBase, uint32(v1489)+64))
	v1567 = *(*int32)(unsafe.Add(mBase, uint32(v1566)))
	v1568 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v519)+2540)) = v1568
	*(*int32)(unsafe.Add(mBase, uint32(v519)+2536)) = v1567
	v1571 = *(*int32)(unsafe.Add(mBase, uint32(v1566)+4))
	if v1571 <= v1568 {
		goto L253
	} else {
		goto L281
	}
L276:
	;
	v1551 = *(*int32)(unsafe.Add(mBase, uint32(v1489)+64))
	v1552 = *(*int32)(unsafe.Add(mBase, uint32(v1551)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v519)+2544)) = v1552
	v1554 = *(*int32)(unsafe.Add(mBase, uint32(v1551)))
	v1555 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v519)+2552)) = v1555
	*(*int32)(unsafe.Add(mBase, uint32(v519)+2548)) = v1554
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	F_BlockRefTableSetLimitBlock(m, v919, v519+int32(2544), v1555, v1555)
	mBase = m.M
	v1565 = m.ExcPending
	if v1565 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L280
	}
L277:
	;
	if v1533 != 0 {
		goto L253
	} else {
		goto L278
	}
L278:
	;
	v1536 = *(*int32)(unsafe.Add(mBase, uint32(v1489)+64))
	v1537 = *(*int32)(unsafe.Add(mBase, uint32(v1536)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v519)+2556)) = v1537
	v1539 = *(*int32)(unsafe.Add(mBase, uint32(v1536)))
	v1540 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v519)+2564)) = v1540
	*(*int32)(unsafe.Add(mBase, uint32(v519)+2560)) = v1539
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	F_BlockRefTableSetLimitBlock(m, v919, v519+int32(2556), v1540, v1540)
	mBase = m.M
	v1550 = m.ExcPending
	if v1550 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L279
	}
L279:
	;
	goto L253
L280:
	;
	goto L253
L281:
	;
	v1577 = int32(0)
	goto L282
L282:
	;
	v1610 = *(*int32)(unsafe.Add(mBase, uint32(v1566+int32(8)+v1577<<(uint(int32(2))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	*(*int32)(unsafe.Add(mBase, uint32(v519)+2532)) = v1610
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	v1616 = int32(0)
	F_BlockRefTableSetLimitBlock(m, v919, v519+int32(2532), v1616, v1616)
	mBase = m.M
	v1619 = m.ExcPending
	if v1619 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L284
	}
L283:
	;
	goto L253
L284:
	;
	v1621 = v1577 + int32(1)
	v1622 = *(*int32)(unsafe.Add(mBase, uint32(v1566)+4))
	if v1621 < v1622 {
		v1577 = v1621
		goto L282
	} else {
		goto L285
	}
L285:
	;
	goto L283
L286:
	;
	v1640 = *(*int32)(unsafe.Add(mBase, uint32(v1489)+64))
	v1641 = *(*int32)(unsafe.Add(mBase, uint32(v1640)+16))
	if v1641&int32(1) != 0 {
		goto L295
	} else {
		goto L296
	}
L287:
	;
	if v1628 == int32(16) {
		goto L290
	} else {
		goto L291
	}
L288:
	;
	goto L289
L289:
	;
	v1631 = *(*int32)(unsafe.Add(mBase, uint32(v1489)+64))
	v1632 = *(*int32)(unsafe.Add(mBase, uint32(v1631)+12))
	if v1632 == int32(1) {
		goto L253
	} else {
		goto L293
	}
L290:
	;
	goto L286
L291:
	;
	goto L253
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	F_BlockRefTableSetLimitBlock(m, v919, v1631, v1632, int32(0))
	mBase = m.M
	v1639 = m.ExcPending
	if v1639 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L294
	}
L294:
	;
	goto L253
L295:
	;
	v1644 = *(*int32)(unsafe.Add(mBase, uint32(v1640)))
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	F_BlockRefTableSetLimitBlock(m, v919, v1640+int32(4), int32(0), v1644)
	mBase = m.M
	v1651 = m.ExcPending
	if v1651 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L298
	}
L296:
	;
	v1654 = v1641
	goto L297
L297:
	;
	if v1654&int32(2) == int32(0) {
		goto L253
	} else {
		goto L299
	}
L298:
	;
	v1652 = *(*int32)(unsafe.Add(mBase, uint32(v1640)+16))
	v1654 = v1652
	goto L297
L299:
	;
	v1659 = *(*int32)(unsafe.Add(mBase, uint32(v1640)))
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	F_BlockRefTableSetLimitBlock(m, v919, v1640+int32(4), int32(2), v1659)
	mBase = m.M
	v1666 = m.ExcPending
	if v1666 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L300
	}
L300:
	;
	goto L253
L301:
	;
	v1855 = *(*int32)(unsafe.Add(mBase, uint32(v1489)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	v1859 = v519 + int32(2568)
	v1860 = int32(0)
	base.MemoryFill(m, v1859, v1860, int32(264))
	v1866 = *(*int64)(unsafe.Add(mBase, uint32(v1855)))
	*(*int64)(unsafe.Add(mBase, uint32(v1859))) = v1866
	if v1860 <= base.I32_extend8_s(v1667) {
		goto L333
	} else {
		goto L334
	}
L302:
	;
	v1672 = *(*int32)(unsafe.Add(mBase, uint32(v1489)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	v1676 = v519 + int32(2832)
	v1677 = int32(0)
	base.MemoryFill(m, v1676, v1677, int32(288))
	v1683 = *(*int64)(unsafe.Add(mBase, uint32(v1672)))
	*(*int64)(unsafe.Add(mBase, uint32(v1676))) = v1683
	if v1677 <= base.I32_extend8_s(v1667) {
		goto L304
	} else {
		goto L305
	}
L303:
	;
	v1791 = int32(0)
	v1792 = *(*int32)(unsafe.Add(mBase, uint32(v519)+2860))
	if v1792 <= v1791 {
		goto L253
	} else {
		goto L325
	}
L304:
	;
	goto L303
L305:
	;
	v1688 = *(*int32)(unsafe.Add(mBase, uint32(v1672)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1676)+8)) = v1688
	if v1688&int32(1) != 0 {
		goto L306
	} else {
		goto L307
	}
L306:
	;
	v1692 = *(*int32)(unsafe.Add(mBase, uint32(v1672)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1676)+12)) = v1692
	v1694 = *(*int32)(unsafe.Add(mBase, uint32(v1672)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1676)+16)) = v1694
	v1700 = v1672 + int32(20)
	goto L308
L307:
	;
	v1700 = v1672 + int32(12)
	goto L308
L308:
	;
	if v1688&int32(2) != 0 {
		goto L309
	} else {
		goto L310
	}
L309:
	;
	v1703 = *(*int32)(unsafe.Add(mBase, uint32(v1700)))
	v1705 = v1700 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1676)+24)) = v1705
	*(*int32)(unsafe.Add(mBase, uint32(v1676)+20)) = v1703
	v1711 = v1705 + v1703<<(uint(int32(2))%32)
	goto L311
L310:
	;
	v1711 = v1700
	goto L311
L311:
	;
	if v1688&int32(4) != 0 {
		goto L312
	} else {
		goto L313
	}
L312:
	;
	v1715 = *(*int32)(unsafe.Add(mBase, uint32(v1711)))
	v1717 = v1711 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1676)+32)) = v1717
	*(*int32)(unsafe.Add(mBase, uint32(v1676)+28)) = v1715
	v1720 = *(*int32)(unsafe.Add(mBase, uint32(v1711)))
	v1724 = v1717 + v1720*int32(12)
	goto L314
L313:
	;
	v1724 = v1711
	goto L314
L314:
	;
	if v1688&int32(256) != 0 {
		goto L315
	} else {
		goto L316
	}
L315:
	;
	v1729 = *(*int32)(unsafe.Add(mBase, uint32(v1724)))
	v1730 = int32(4)
	v1731 = v1724 + v1730
	*(*int32)(unsafe.Add(mBase, uint32(v1676)+40)) = v1731
	*(*int32)(unsafe.Add(mBase, uint32(v1676)+36)) = v1729
	v1734 = *(*int32)(unsafe.Add(mBase, uint32(v1724)))
	v1738 = v1731 + v1734<<(uint(v1730)%32)
	goto L317
L316:
	;
	v1738 = v1724
	goto L317
L317:
	;
	if v1688&int32(8) != 0 {
		goto L318
	} else {
		goto L319
	}
L318:
	;
	v1743 = *(*int32)(unsafe.Add(mBase, uint32(v1738)))
	v1744 = int32(4)
	v1745 = v1738 + v1744
	*(*int32)(unsafe.Add(mBase, uint32(v1676)+48)) = v1745
	*(*int32)(unsafe.Add(mBase, uint32(v1676)+44)) = v1743
	v1748 = *(*int32)(unsafe.Add(mBase, uint32(v1738)))
	v1752 = v1745 + v1748<<(uint(v1744)%32)
	goto L320
L319:
	;
	v1752 = v1738
	goto L320
L320:
	;
	if v1688&int32(16) == int32(0) {
		v1776 = v1688
		v1777 = v1752
		goto L321
	} else {
		goto L322
	}
L321:
	;
	if v1776&int32(32) == int32(0) {
		goto L304
	} else {
		goto L324
	}
L322:
	;
	v1759 = *(*int32)(unsafe.Add(mBase, uint32(v1752)))
	*(*int32)(unsafe.Add(mBase, uint32(v1676)+52)) = v1759
	v1762 = v1752 + int32(4)
	if v1688&int32(128) == int32(0) {
		v1776 = v1688
		v1777 = v1762
		goto L321
	} else {
		goto L323
	}
L323:
	;
	v1770 = F_strlcpy(m, v519+int32(2888), v1762, int32(200))
	mBase = m.M
	v1771 = F_strlen(m, v1762)
	mBase = m.M
	v1775 = *(*int32)(unsafe.Add(mBase, uint32(v1676)+8))
	v1776 = v1775
	v1777 = v1771 + v1762 + int32(1)
	goto L321
L324:
	;
	v1782 = *(*int64)(unsafe.Add(mBase, uint32(v1777)))
	v1783 = *(*int64)(unsafe.Add(mBase, uint32(v1777)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1676)+280)) = v1783
	*(*int64)(unsafe.Add(mBase, uint32(v1676)+272)) = v1782
	goto L304
L325:
	;
	v1795 = v1791
	goto L326
L326:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	v1828 = v1795 * int32(12)
	v1829 = *(*int32)(unsafe.Add(mBase, uint32(v519)+2864))
	v1831 = int32(0)
	F_BlockRefTableSetLimitBlock(m, v919, v1828+v1829, v1831, v1831)
	mBase = m.M
	v1834 = m.ExcPending
	if v1834 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L328
	}
L327:
	;
	goto L253
L328:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	v1837 = *(*int32)(unsafe.Add(mBase, uint32(v519)+2864))
	F_BlockRefTableSetLimitBlock(m, v919, v1837+v1828, int32(2), int32(0))
	mBase = m.M
	v1842 = m.ExcPending
	if v1842 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L329
	}
L329:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	v1845 = *(*int32)(unsafe.Add(mBase, uint32(v519)+2864))
	F_BlockRefTableSetLimitBlock(m, v919, v1845+v1828, int32(3), int32(0))
	mBase = m.M
	v1850 = m.ExcPending
	if v1850 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L330
	}
L330:
	;
	v1852 = v1795 + int32(1)
	v1853 = *(*int32)(unsafe.Add(mBase, uint32(v519)+2860))
	if v1852 < v1853 {
		v1795 = v1852
		goto L326
	} else {
		goto L331
	}
L331:
	;
	goto L327
L332:
	;
	v1960 = int32(0)
	v1961 = *(*int32)(unsafe.Add(mBase, uint32(v519)+2596))
	if v1961 <= v1960 {
		goto L253
	} else {
		goto L351
	}
L333:
	;
	goto L332
L334:
	;
	v1871 = *(*int32)(unsafe.Add(mBase, uint32(v1855)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1859)+8)) = v1871
	if v1871&int32(1) != 0 {
		goto L335
	} else {
		goto L336
	}
L335:
	;
	v1875 = *(*int32)(unsafe.Add(mBase, uint32(v1855)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1859)+12)) = v1875
	v1877 = *(*int32)(unsafe.Add(mBase, uint32(v1855)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1859)+16)) = v1877
	v1883 = v1855 + int32(20)
	goto L337
L336:
	;
	v1883 = v1855 + int32(12)
	goto L337
L337:
	;
	if v1871&int32(2) != 0 {
		goto L338
	} else {
		goto L339
	}
L338:
	;
	v1886 = *(*int32)(unsafe.Add(mBase, uint32(v1883)))
	v1888 = v1883 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1859)+24)) = v1888
	*(*int32)(unsafe.Add(mBase, uint32(v1859)+20)) = v1886
	v1894 = v1888 + v1886<<(uint(int32(2))%32)
	goto L340
L339:
	;
	v1894 = v1883
	goto L340
L340:
	;
	if v1871&int32(4) != 0 {
		goto L341
	} else {
		goto L342
	}
L341:
	;
	v1898 = *(*int32)(unsafe.Add(mBase, uint32(v1894)))
	v1900 = v1894 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1859)+32)) = v1900
	*(*int32)(unsafe.Add(mBase, uint32(v1859)+28)) = v1898
	v1903 = *(*int32)(unsafe.Add(mBase, uint32(v1894)))
	v1907 = v1900 + v1903*int32(12)
	goto L343
L342:
	;
	v1907 = v1894
	goto L343
L343:
	;
	if v1871&int32(256) != 0 {
		goto L344
	} else {
		goto L345
	}
L344:
	;
	v1912 = *(*int32)(unsafe.Add(mBase, uint32(v1907)))
	v1913 = int32(4)
	v1914 = v1907 + v1913
	*(*int32)(unsafe.Add(mBase, uint32(v1859)+40)) = v1914
	*(*int32)(unsafe.Add(mBase, uint32(v1859)+36)) = v1912
	v1917 = *(*int32)(unsafe.Add(mBase, uint32(v1907)))
	v1921 = v1914 + v1917<<(uint(v1913)%32)
	goto L346
L345:
	;
	v1921 = v1907
	goto L346
L346:
	;
	if v1871&int32(16) == int32(0) {
		v1945 = v1871
		v1946 = v1921
		goto L347
	} else {
		goto L348
	}
L347:
	;
	if v1945&int32(32) == int32(0) {
		goto L333
	} else {
		goto L350
	}
L348:
	;
	v1928 = *(*int32)(unsafe.Add(mBase, uint32(v1921)))
	*(*int32)(unsafe.Add(mBase, uint32(v1859)+44)) = v1928
	v1931 = v1921 + int32(4)
	if v1871&int32(128) == int32(0) {
		v1945 = v1871
		v1946 = v1931
		goto L347
	} else {
		goto L349
	}
L349:
	;
	v1939 = F_strlcpy(m, v519+int32(2616), v1931, int32(200))
	mBase = m.M
	v1940 = F_strlen(m, v1931)
	mBase = m.M
	v1944 = *(*int32)(unsafe.Add(mBase, uint32(v1859)+8))
	v1945 = v1944
	v1946 = v1940 + v1931 + int32(1)
	goto L347
L350:
	;
	v1951 = *(*int64)(unsafe.Add(mBase, uint32(v1946)))
	v1952 = *(*int64)(unsafe.Add(mBase, uint32(v1946)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1859)+256)) = v1952
	*(*int64)(unsafe.Add(mBase, uint32(v1859)+248)) = v1951
	goto L333
L351:
	;
	v1964 = v1960
	goto L352
L352:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	v1997 = v1964 * int32(12)
	v1998 = *(*int32)(unsafe.Add(mBase, uint32(v519)+2600))
	v2000 = int32(0)
	F_BlockRefTableSetLimitBlock(m, v919, v1997+v1998, v2000, v2000)
	mBase = m.M
	v2003 = m.ExcPending
	if v2003 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L354
	}
L353:
	;
	goto L253
L354:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	v2006 = *(*int32)(unsafe.Add(mBase, uint32(v519)+2600))
	F_BlockRefTableSetLimitBlock(m, v919, v2006+v1997, int32(2), int32(0))
	mBase = m.M
	v2011 = m.ExcPending
	if v2011 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L355
	}
L355:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	v2014 = *(*int32)(unsafe.Add(mBase, uint32(v519)+2600))
	F_BlockRefTableSetLimitBlock(m, v919, v2014+v1997, int32(3), int32(0))
	mBase = m.M
	v2019 = m.ExcPending
	if v2019 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L356
	}
L356:
	;
	v2021 = v1964 + int32(1)
	v2022 = *(*int32)(unsafe.Add(mBase, uint32(v519)+2596))
	if v2021 < v2022 {
		v1964 = v2021
		goto L352
	} else {
		goto L357
	}
L357:
	;
	goto L353
L358:
	;
	v2059 = v2054
	goto L359
L359:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	v2092 = v2059 & int32(255)
	v2094 = v519 + int32(436)
	v2096 = v519 + int32(432)
	v2098 = v519 + int32(428)
	v2099 = int32(0)
	v2101 = *(*int32)(unsafe.Add(mBase, uint32(v943)+96))
	v2102 = *(*int32)(unsafe.Add(mBase, uint32(v2101)+72))
	if v2102 < v2092 {
		v2126 = v2099
		goto L363
	} else {
		goto L364
	}
L360:
	;
	v2144 = int32(0)
	goto L252
L361:
	;
	v2139 = v2059 + int32(1)
	v2140 = *(*int32)(unsafe.Add(mBase, uint32(v943)+96))
	v2141 = *(*int32)(unsafe.Add(mBase, uint32(v2140)+72))
	if v2139 <= v2141 {
		v2059 = v2139
		goto L359
	} else {
		goto L379
	}
L362:
	;
	if v2126 == int32(0) {
		goto L361
	} else {
		goto L376
	}
L363:
	;
	goto L362
L364:
	;
	v2106 = v2101 + v2092*int32(52)
	v2107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2106)+76)))
	if v2107 != int32(1) {
		v2126 = v2099
		goto L363
	} else {
		goto L365
	}
L365:
	;
	v2111 = v2106 + int32(76)
	if v2094 != 0 {
		goto L366
	} else {
		goto L367
	}
L366:
	;
	v2112 = *(*int32)(unsafe.Add(mBase, uint32(v2111)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v2094)+8)) = v2112
	v2114 = *(*int64)(unsafe.Add(mBase, uint32(v2111)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v2094))) = v2114
	goto L368
L367:
	;
	goto L368
L368:
	;
	if v2096 != 0 {
		goto L369
	} else {
		goto L370
	}
L369:
	;
	v2116 = *(*int32)(unsafe.Add(mBase, uint32(v2111)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v2096))) = v2116
	goto L371
L370:
	;
	goto L371
L371:
	;
	if v2098 != 0 {
		goto L372
	} else {
		goto L373
	}
L372:
	;
	v2118 = *(*int32)(unsafe.Add(mBase, uint32(v2111)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2098))) = v2118
	goto L374
L373:
	;
	goto L374
L374:
	;
	v2126 = int32(1)
	goto L363
L376:
	;
	v2129 = *(*int32)(unsafe.Add(mBase, uint32(v519)+432))
	if v2129 == int32(1) {
		goto L361
	} else {
		goto L377
	}
L377:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	v2134 = *(*int32)(unsafe.Add(mBase, uint32(v519)+428))
	F_BlockRefTableMarkBlockModified(m, v919, v2094, v2129, v2134)
	mBase = m.M
	v2136 = m.ExcPending
	if v2136 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L378
	}
L378:
	;
	goto L361
L379:
	;
	goto L360
L380:
	;
	v2185 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[9]))
	*(*int64)(unsafe.Add(mBase, uint32(v2185)+24)) = v2174
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	v2190 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[8]))
	F_LWLockRelease(m, v2190+int32(_a_F_WalSummarizerMain_4))
	mBase = m.M
	v2194 = m.ExcPending
	if v2194 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L381
	}
L381:
	;
	if v1486 == int32(0) {
		goto L382
	} else {
		goto L383
	}
L382:
	;
	v2197 = *(*int64)(unsafe.Add(mBase, uint32(v943)+40))
	if base.Ui64(v1418) <= base.Ui64(v2197) {
		v2276 = v2144
		v2301 = v2174
		goto L178
	} else {
		goto L385
	}
L383:
	;
	goto L384
L384:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	F_ProcessWalSummarizerInterrupts(m)
	mBase = m.M
	v2202 = m.ExcPending
	if v2202 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L386
	}
L385:
	;
	goto L384
L386:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	v2207 = F_XLogReadRecord(m, v943, v519+int32(448))
	mBase = m.M
	v2208 = m.ExcPending
	if v2208 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L387
	}
L387:
	;
	if v2207 != 0 {
		v1456 = v2144
		goto L247
	} else {
		goto L388
	}
L388:
	;
	goto L248
L389:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	v2246 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2247 = m.ExcPending
	if v2247 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L390
	}
L390:
	;
	if v2246 != 0 {
		goto L391
	} else {
		goto L392
	}
L391:
	;
	v2248 = *(*int64)(unsafe.Add(mBase, uint32(v943)+40))
	v2249 = *(*int64)(unsafe.Add(mBase, uint32(v924)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	*(*uint32)(unsafe.Add(mBase, uint32(v533))) = uint32(v2249)
	v2253 = int64(32)
	v2254 = int64(base.Ui64(v2249) >> (uint(v2253) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v519)+124)) = uint32(v2254)
	*(*uint32)(unsafe.Add(mBase, uint32(v519)+120)) = uint32(v2248)
	v2258 = int64(base.Ui64(v2248) >> (uint(v2253) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v519)+116)) = uint32(v2258)
	*(*int32)(unsafe.Add(mBase, uint32(v519)+112)) = v838
	F_errmsg_internal(m, int32(_a_F_WalSummarizerMain_25), v519+int32(112))
	mBase = m.M
	v2265 = m.ExcPending
	if v2265 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L394
	}
L392:
	;
	goto L393
L393:
	;
	v2275 = *(*int64)(unsafe.Add(mBase, uint32(v924)+8))
	v2276 = v2209
	v2301 = v2275
	goto L178
L394:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	F_errfinish(m, int32(_a_F_WalSummarizerMain_1), int32(1040), int32(_a_F_WalSummarizerMain_21))
	mBase = m.M
	v2272 = m.ExcPending
	if v2272 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L395
	}
L395:
	;
	goto L393
L396:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	F_XLogReaderFree(m, v943)
	mBase = m.M
	v2344 = m.ExcPending
	if v2344 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L397
	}
L397:
	;
	v2345 = base.B2i32(base.Ui64(v2331) <= base.Ui64(v1414))
	if (v2307|v2345)&int32(1) != 0 {
		goto L174
	} else {
		goto L398
	}
L398:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	v2352 = v519 + int32(1504)
	v2356 = F_pg_snprintf(m, v2352, int32(1024), int32(_a_F_WalSummarizerMain_27), int32(0))
	mBase = m.M
	v2357 = m.ExcPending
	if v2357 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L399
	}
L399:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	v2360 = base.I32_wrap_i64(v2331)
	*(*int32)(unsafe.Add(mBase, uint32(v519)+96)) = v2360
	v2362 = int64(32)
	v2364 = base.I32_wrap_i64(int64(base.Ui64(v2331) >> (uint(v2362) % 64)))
	*(*int32)(unsafe.Add(mBase, uint32(v519)+92)) = v2364
	v2366 = base.I32_wrap_i64(v1414)
	*(*int32)(unsafe.Add(mBase, uint32(v519)+88)) = v2366
	v2370 = base.I32_wrap_i64(int64(base.Ui64(v1414) >> (uint(v2362) % 64)))
	*(*int32)(unsafe.Add(mBase, uint32(v519)+84)) = v2370
	*(*int32)(unsafe.Add(mBase, uint32(v519)+80)) = v838
	v2379 = F_pg_snprintf(m, v519+int32(480), int32(1024), int32(_a_F_WalSummarizerMain_11), v519+int32(80))
	mBase = m.M
	v2380 = m.ExcPending
	if v2380 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L400
	}
L400:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	*(*int64)(unsafe.Add(mBase, uint32(v519)+472)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	v2386 = F_PathNameOpenFile(m, v2352, int32(577))
	mBase = m.M
	v2387 = m.ExcPending
	if v2387 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L401
	}
L401:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+464)) = v2386
	if int32(0) <= v2386 {
		goto L175
	} else {
		goto L402
	}
L402:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2396 = m.ExcPending
	if v2396 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L403
	}
L403:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	F_errcode_for_file_access(m)
	mBase = m.M
	v2400 = m.ExcPending
	if v2400 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L404
	}
L404:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	*(*int32)(unsafe.Add(mBase, uint32(v519)+32)) = v2352
	F_errmsg(m, int32(_a_F_WalSummarizerMain_28), v519+int32(32))
	mBase = m.M
	v2408 = m.ExcPending
	if v2408 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L405
	}
L405:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	F_errfinish(m, int32(_a_F_WalSummarizerMain_1), int32(1215), int32(_a_F_WalSummarizerMain_21))
	mBase = m.M
	v2415 = m.ExcPending
	if v2415 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L406
	}
L406:
	;
	goto L3
L407:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	F_errcode_for_file_access(m)
	mBase = m.M
	v2426 = m.ExcPending
	if v2426 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L408
	}
L408:
	;
	v2427 = *(*int64)(unsafe.Add(mBase, uint32(v943)+40))
	v2430 = base.I32_wrap_i64(int64(base.Ui64(v2427) >> (uint(int64(32)) % 64)))
	v2431 = base.I32_wrap_i64(v2427)
	if v2418 != 0 {
		goto L409
	} else {
		goto L410
	}
L409:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	v2434 = *(*int32)(unsafe.Add(mBase, uint32(v519)+448))
	*(*int32)(unsafe.Add(mBase, uint32(v519)+172)) = v2434
	*(*int32)(unsafe.Add(mBase, uint32(v519)+168)) = v2431
	*(*int32)(unsafe.Add(mBase, uint32(v519)+164)) = v2430
	*(*int32)(unsafe.Add(mBase, uint32(v519)+160)) = v838
	F_errmsg(m, int32(_a_F_WalSummarizerMain_29), v519+int32(160))
	mBase = m.M
	v2443 = m.ExcPending
	if v2443 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L412
	}
L410:
	;
	goto L411
L411:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	*(*int32)(unsafe.Add(mBase, uint32(v519)+152)) = v2431
	*(*int32)(unsafe.Add(mBase, uint32(v519)+148)) = v2430
	*(*int32)(unsafe.Add(mBase, uint32(v519)+144)) = v838
	F_errmsg(m, int32(_a_F_WalSummarizerMain_30), v519+int32(144))
	mBase = m.M
	v2460 = m.ExcPending
	if v2460 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L414
	}
L412:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	F_errfinish(m, int32(_a_F_WalSummarizerMain_1), int32(1050), int32(_a_F_WalSummarizerMain_21))
	mBase = m.M
	v2450 = m.ExcPending
	if v2450 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L413
	}
L413:
	;
	goto L3
L414:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	F_errfinish(m, int32(_a_F_WalSummarizerMain_1), int32(1055), int32(_a_F_WalSummarizerMain_21))
	mBase = m.M
	v2467 = m.ExcPending
	if v2467 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L415
	}
L415:
	;
	goto L3
L416:
	;
	v3333 = m.G0
	v3335 = v3333 - int32(32)
	m.G0 = v3335
	v3337 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3335)+24)) = v3337
	*(*int64)(unsafe.Add(mBase, uint32(v3335)+16)) = v3337
	*(*int64)(unsafe.Add(mBase, uint32(v3335)+8)) = v3337
	v3344 = v2473 + int32(12)
	v3345 = *(*int32)(unsafe.Add(mBase, uint32(v2473)+uint32(_c_F_WalSummarizerMain[22])))
	v3348 = int32(24)
	v3349 = m.Env.Pgmem_crc32c(m, v3345, v3335+int32(8), v3348)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v2473)+uint32(_c_F_WalSummarizerMain[22]))) = v3349
	v3351 = *(*int32)(unsafe.Add(mBase, uint32(v2473)+uint32(_c_F_WalSummarizerMain[23])))
	if int32(_a_F_WalSummarizerMain_31) <= v3351+v3348 {
		goto L528
	} else {
		goto L529
	}
L417:
	;
	v2504 = F_palloc(m, v2499*int32(24))
	mBase = m.M
	v2505 = m.ExcPending
	if v2505 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L418
	}
L418:
	;
	v2506 = *(*int32)(unsafe.Add(mBase, uint32(v919)))
	v2507 = *(*int64)(unsafe.Add(mBase, uint32(v2506)))
	if v2507 == int64(0) {
		v2552 = v2487
		goto L419
	} else {
		goto L420
	}
L419:
	;
	v2583 = v2473 + int32(20)
	v2589 = v2552
	v2591 = int32(0)
	v2594 = v2506
	v2599 = v2470
	goto L427
L420:
	;
	v2510 = *(*int32)(unsafe.Add(mBase, uint32(v2506)+20))
	v2517 = int32(0)
	goto L421
L421:
	;
	v2545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2510+v2517*int32(40))+20)))
	if v2545 != int32(1) {
		goto L423
	} else {
		goto L424
	}
L422:
	;
	v2552 = v2487
	goto L419
L423:
	;
	v2552 = v2517
	goto L419
L424:
	;
	goto L425
L425:
	;
	v2549 = v2517 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v2549)) < base.Ui64(v2507) {
		v2517 = v2549
		goto L421
	} else {
		goto L426
	}
L426:
	;
	goto L422
L427:
	;
	v2620 = v2589
	v2621 = v2591
	v2623 = v2591
	goto L430
L428:
	;
	F_pg_qsort(m, v2504, v2599, int32(24), int32(_a_F_WalSummarizerMain_32))
	mBase = m.M
	v2723 = m.ExcPending
	if v2723 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L439
	}
L429:
	;
	goto L428
L430:
	;
	if v2623&int32(1) != 0 {
		goto L429
	} else {
		goto L432
	}
L431:
	;
	v2665 = v2504 + v2599*int32(24)
	v2666 = *(*int32)(unsafe.Add(mBase, uint32(v2659)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2665)+8)) = v2666
	v2668 = *(*int64)(unsafe.Add(mBase, uint32(v2659)))
	*(*int64)(unsafe.Add(mBase, uint32(v2665))) = v2668
	v2670 = *(*int32)(unsafe.Add(mBase, uint32(v2659)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v2665)+12)) = v2670
	v2672 = *(*int32)(unsafe.Add(mBase, uint32(v2659)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v2665)+16)) = v2672
	v2674 = *(*int32)(unsafe.Add(mBase, uint32(v2659)+24))
	v2680 = v2674
	goto L434
L432:
	;
	v2647 = *(*int32)(unsafe.Add(mBase, uint32(v2594)+12))
	v2648 = int32(1)
	v2649 = v2620 - v2648
	v2653 = base.B2i32(v2647&(v2649^v2552) == int32(0))
	v2654 = v2653 | v2621
	v2657 = v2647 & v2649
	v2658 = *(*int32)(unsafe.Add(mBase, uint32(v2594)+20))
	v2659 = v2620*int32(40) + v2658
	v2660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2659)+20)))
	if v2660 != v2648 {
		v2620 = v2657
		v2621 = v2654
		v2623 = v2653
		goto L430
	} else {
		goto L433
	}
L433:
	;
	goto L431
L434:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2665)+20)) = v2680
	if v2680 == int32(0) {
		goto L436
	} else {
		goto L437
	}
L435:
	;
	v2719 = *(*int32)(unsafe.Add(mBase, uint32(v919)))
	v2589 = v2657
	v2591 = v2654
	v2594 = v2719
	v2599 = v2599 + int32(1)
	goto L427
L436:
	;
	goto L435
L437:
	;
	v2708 = *(*int32)(unsafe.Add(mBase, uint32(v2659)+32))
	v2714 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2708+v2680<<(uint(int32(1))%32)-int32(2)))))
	if v2714 != 0 {
		goto L436
	} else {
		goto L438
	}
L438:
	;
	v2680 = v2680 - int32(1)
	goto L434
L439:
	;
	v2724 = *(*int32)(unsafe.Add(mBase, uint32(v919)))
	v2725 = *(*int32)(unsafe.Add(mBase, uint32(v2724)+8))
	if v2725 == int32(0) {
		goto L416
	} else {
		goto L440
	}
L440:
	;
	v2733 = int32(0)
	goto L441
L441:
	;
	v2759 = *(*int32)(unsafe.Add(mBase, uint32(v2473)+uint32(_c_F_WalSummarizerMain[22])))
	v2760 = int32(24)
	v2762 = v2504 + v2733*v2760
	v2764 = m.Env.Pgmem_crc32c(m, v2759, v2762, v2760)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v2473)+uint32(_c_F_WalSummarizerMain[22]))) = v2764
	v2766 = *(*int32)(unsafe.Add(mBase, uint32(v2473)+uint32(_c_F_WalSummarizerMain[23])))
	if int32(_a_F_WalSummarizerMain_31) <= v2766+v2760 {
		goto L443
	} else {
		goto L444
	}
L442:
	;
	goto L416
L443:
	;
	v2771 = *(*int32)(unsafe.Add(mBase, uint32(v2473)+16))
	v2772 = *(*int32)(unsafe.Add(mBase, uint32(v2473)+12))
	v2773 = m.T0[v2772].(func(*base.Module, int32, int32, int32) int32)(m, v2771, v2583, v2766)
	mBase = m.M
	v2774 = m.ExcPending
	if v2774 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L446
	}
L444:
	;
	v2778 = v2766
	goto L445
L445:
	;
	v2779 = v2778 + v2583
	v2780 = *(*int64)(unsafe.Add(mBase, uint32(v2762)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v2779)+16)) = v2780
	v2782 = *(*int64)(unsafe.Add(mBase, uint32(v2762)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2779)+8)) = v2782
	v2784 = *(*int64)(unsafe.Add(mBase, uint32(v2762)))
	*(*int64)(unsafe.Add(mBase, uint32(v2779))) = v2784
	v2786 = *(*int32)(unsafe.Add(mBase, uint32(v2473)+uint32(_c_F_WalSummarizerMain[23])))
	*(*int32)(unsafe.Add(mBase, uint32(v2473)+uint32(_c_F_WalSummarizerMain[23]))) = v2786 + int32(24)
	v2790 = *(*int32)(unsafe.Add(mBase, uint32(v919)))
	v2791 = *(*int32)(unsafe.Add(mBase, uint32(v2762)+12))
	v2792 = *(*int32)(unsafe.Add(mBase, uint32(v2762)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2473)+uint32(_c_F_WalSummarizerMain[24]))) = v2792
	v2794 = *(*int64)(unsafe.Add(mBase, uint32(v2762)))
	*(*int64)(unsafe.Add(mBase, uint32(v2473)+uint32(_c_F_WalSummarizerMain[25]))) = v2794
	*(*int32)(unsafe.Add(mBase, uint32(v2473)+uint32(_c_F_WalSummarizerMain[26]))) = v2791
	v2798 = v2473 + int32(_a_F_WalSummarizerMain_33)
	v2799 = int32(16)
	v2805 = int32(-1636608416)
	if v2798&int32(3) != 0 {
		goto L451
	} else {
		goto L452
	}
L446:
	;
	v2775 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2473)+uint32(_c_F_WalSummarizerMain[23]))) = v2775
	v2778 = v2775
	goto L445
L447:
	;
	v3064 = *(*int32)(unsafe.Add(mBase, uint32(v2790)+20))
	v3065 = *(*int32)(unsafe.Add(mBase, uint32(v2790)+12))
	v3066 = (v3059 ^ v3051 - base.I32_rotl(v3059, int32(24))) & v3065
	v3069 = v3064 + v3066*int32(40)
	v3070 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3069)+20)))
	if v3070 == int32(0) {
		goto L488
	} else {
		goto L489
	}
L448:
	;
	v3037 = int32(14)
	v3039 = v3033 ^ v3034 - base.I32_rotl(v3033, v3037)
	v3043 = v3039 ^ v3032 - base.I32_rotl(v3039, int32(11))
	v3047 = v3043 ^ v3033 - base.I32_rotl(v3043, int32(25))
	v3051 = v3047 ^ v3039 - base.I32_rotl(v3047, int32(16))
	v3055 = v3051 ^ v3043 - base.I32_rotl(v3051, int32(4))
	v3059 = v3055 ^ v3047 - base.I32_rotl(v3055, v3037)
	goto L447
L449:
	;
	switch v2959 - int32(1) {
	case 0:
		v3025 = v2950
		v3026 = v2951
		v3027 = v2955
		goto L476
	case 1:
		v3018 = v2950
		v3019 = v2951
		v3020 = v2955
		goto L477
	case 2:
		v3011 = v2950
		v3012 = v2951
		v3013 = v2955
		goto L478
	case 3:
		v3005 = v2951
		v3006 = v2955
		goto L479
	case 4:
		v3001 = v2951
		v3002 = v2955
		goto L480
	case 5:
		v2995 = v2951
		v2996 = v2955
		goto L481
	case 6:
		v2989 = v2951
		v2990 = v2955
		goto L482
	case 7:
		v2984 = v2955
		goto L483
	case 8:
		v2979 = v2955
		goto L484
	case 9:
		v2974 = v2955
		goto L485
	case 10:
		goto L486
	default:
		v3032 = v2950
		v3033 = v2951
		v3034 = v2955
		goto L448
	}
L450:
	;
	v2914 = v2798
	v2915 = v2799
	v2916 = v2805
	v2917 = v2805
	v2918 = v2805
	goto L473
L451:
	;
	goto L450
L452:
	;
	goto L453
L453:
	;
	goto L457
L455:
	;
	switch v2857 - int32(1) {
	case 0:
		v2911 = v2848
		goto L462
	case 1:
		v2906 = v2848
		goto L463
	case 2:
		goto L464
	case 3:
		v2899 = v2849
		goto L465
	case 4:
		v2896 = v2849
		goto L466
	case 5:
		v2891 = v2849
		goto L467
	case 6:
		goto L468
	case 7:
		v2882 = v2853
		goto L469
	case 8:
		v2877 = v2853
		goto L470
	case 9:
		v2872 = v2853
		goto L471
	case 10:
		goto L472
	default:
		v3032 = v2848
		v3033 = v2849
		v3034 = v2853
		goto L448
	}
L457:
	;
	goto L458
L458:
	;
	v2812 = v2798
	v2813 = v2799
	v2814 = v2805
	v2815 = v2805
	v2816 = v2805
	goto L459
L459:
	;
	v2818 = *(*int32)(unsafe.Add(mBase, uint32(v2812)+4))
	v2819 = v2818 + v2815
	v2820 = *(*int32)(unsafe.Add(mBase, uint32(v2812)))
	v2822 = *(*int32)(unsafe.Add(mBase, uint32(v2812)+8))
	v2823 = v2822 + v2816
	v2825 = int32(4)
	v2827 = v2820 + v2814 - v2823 ^ base.I32_rotl(v2823, v2825)
	v2831 = v2819 - v2827 ^ base.I32_rotl(v2827, int32(6))
	v2832 = v2823 + v2819
	v2833 = v2827 + v2832
	v2834 = v2831 + v2833
	v2838 = v2832 - v2831 ^ base.I32_rotl(v2831, int32(8))
	v2842 = v2833 - v2838 ^ base.I32_rotl(v2838, int32(16))
	v2846 = v2834 - v2842 ^ base.I32_rotl(v2842, int32(19))
	v2847 = v2838 + v2834
	v2848 = v2842 + v2847
	v2849 = v2846 + v2848
	v2853 = v2847 - v2846 ^ base.I32_rotl(v2846, v2825)
	v2854 = int32(12)
	v2855 = v2812 + v2854
	v2857 = v2813 - v2854
	if base.Ui32(int32(11)) < base.Ui32(v2857) {
		v2812 = v2855
		v2813 = v2857
		v2814 = v2848
		v2815 = v2849
		v2816 = v2853
		goto L459
	} else {
		goto L461
	}
L460:
	;
	goto L455
L461:
	;
	goto L460
L462:
	;
	v2912 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2855))))
	v3032 = v2911 + v2912
	v3033 = v2849
	v3034 = v2853
	goto L448
L463:
	;
	v2907 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2855)+1)))
	v2911 = v2907<<(uint(int32(8))%32) + v2906
	goto L462
L464:
	;
	v2902 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2855)+2)))
	v2906 = v2902<<(uint(int32(16))%32) + v2848
	goto L463
L465:
	;
	v2900 = *(*int32)(unsafe.Add(mBase, uint32(v2855)))
	v3032 = v2900 + v2848
	v3033 = v2899
	v3034 = v2853
	goto L448
L466:
	;
	v2897 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2855)+4)))
	v2899 = v2896 + v2897
	goto L465
L467:
	;
	v2892 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2855)+5)))
	v2896 = v2892<<(uint(int32(8))%32) + v2891
	goto L466
L468:
	;
	v2887 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2855)+6)))
	v2891 = v2887<<(uint(int32(16))%32) + v2849
	goto L467
L469:
	;
	v2883 = *(*int32)(unsafe.Add(mBase, uint32(v2855)))
	v2885 = *(*int32)(unsafe.Add(mBase, uint32(v2855)+4))
	v3032 = v2883 + v2848
	v3033 = v2885 + v2849
	v3034 = v2882
	goto L448
L470:
	;
	v2878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2855)+8)))
	v2882 = v2878<<(uint(int32(8))%32) + v2877
	goto L469
L471:
	;
	v2873 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2855)+9)))
	v2877 = v2873<<(uint(int32(16))%32) + v2872
	goto L470
L472:
	;
	v2868 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2855)+10)))
	v2872 = v2868<<(uint(int32(24))%32) + v2853
	goto L471
L473:
	;
	v2920 = *(*int32)(unsafe.Add(mBase, uint32(v2914)+4))
	v2921 = v2920 + v2917
	v2922 = *(*int32)(unsafe.Add(mBase, uint32(v2914)))
	v2924 = *(*int32)(unsafe.Add(mBase, uint32(v2914)+8))
	v2925 = v2924 + v2918
	v2927 = int32(4)
	v2929 = v2922 + v2916 - v2925 ^ base.I32_rotl(v2925, v2927)
	v2933 = v2921 - v2929 ^ base.I32_rotl(v2929, int32(6))
	v2934 = v2925 + v2921
	v2935 = v2929 + v2934
	v2936 = v2933 + v2935
	v2940 = v2934 - v2933 ^ base.I32_rotl(v2933, int32(8))
	v2944 = v2935 - v2940 ^ base.I32_rotl(v2940, int32(16))
	v2948 = v2936 - v2944 ^ base.I32_rotl(v2944, int32(19))
	v2949 = v2940 + v2936
	v2950 = v2944 + v2949
	v2951 = v2948 + v2950
	v2955 = v2949 - v2948 ^ base.I32_rotl(v2948, v2927)
	v2956 = int32(12)
	v2957 = v2914 + v2956
	v2959 = v2915 - v2956
	if base.Ui32(int32(11)) < base.Ui32(v2959) {
		v2914 = v2957
		v2915 = v2959
		v2916 = v2950
		v2917 = v2951
		v2918 = v2955
		goto L473
	} else {
		goto L475
	}
L474:
	;
	goto L449
L475:
	;
	goto L474
L476:
	;
	v3028 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2957))))
	v3032 = v3025 + v3028
	v3033 = v3026
	v3034 = v3027
	goto L448
L477:
	;
	v3021 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2957)+1)))
	v3025 = v3021<<(uint(int32(8))%32) + v3018
	v3026 = v3019
	v3027 = v3020
	goto L476
L478:
	;
	v3014 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2957)+2)))
	v3018 = v3014<<(uint(int32(16))%32) + v3011
	v3019 = v3012
	v3020 = v3013
	goto L477
L479:
	;
	v3007 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2957)+3)))
	v3011 = v3007<<(uint(int32(24))%32) + v2950
	v3012 = v3005
	v3013 = v3006
	goto L478
L480:
	;
	v3003 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2957)+4)))
	v3005 = v3001 + v3003
	v3006 = v3002
	goto L479
L481:
	;
	v2997 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2957)+5)))
	v3001 = v2997<<(uint(int32(8))%32) + v2995
	v3002 = v2996
	goto L480
L482:
	;
	v2991 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2957)+6)))
	v2995 = v2991<<(uint(int32(16))%32) + v2989
	v2996 = v2990
	goto L481
L483:
	;
	v2985 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2957)+7)))
	v2989 = v2985<<(uint(int32(24))%32) + v2951
	v2990 = v2984
	goto L482
L484:
	;
	v2980 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2957)+8)))
	v2984 = v2980<<(uint(int32(8))%32) + v2979
	goto L483
L485:
	;
	v2975 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2957)+9)))
	v2979 = v2975<<(uint(int32(16))%32) + v2974
	goto L484
L486:
	;
	v2970 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2957)+10)))
	v2974 = v2970<<(uint(int32(24))%32) + v2955
	goto L485
L487:
	;
	v3153 = *(*int32)(unsafe.Add(mBase, uint32(v2762)+20))
	if v3153 == int32(0) {
		goto L495
	} else {
		goto L496
	}
L488:
	;
	v3128 = int32(0)
	goto L487
L489:
	;
	goto L490
L490:
	;
	v3079 = v3069
	v3080 = v3066
	goto L491
L491:
	;
	v3104 = *(*int64)(unsafe.Add(mBase, uint32(v3079)))
	v3105 = *(*int64)(unsafe.Add(mBase, uint32(v2473)+uint32(_c_F_WalSummarizerMain[25])))
	v3107 = *(*int64)(unsafe.Add(mBase, uint32(v3079)+8))
	v3110 = *(*int64)(unsafe.Add(mBase, uint32(v2473)+uint32(_c_F_WalSummarizerMain[24])))
	if v3104^v3105|(v3107^v3110) == int64(0) {
		v3128 = v3079
		goto L487
	} else {
		goto L493
	}
L492:
	;
	v3128 = int32(0)
	goto L487
L493:
	;
	v3117 = (v3080 + int32(1)) & v3065
	v3120 = v3064 + v3117*int32(40)
	v3121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3120)+20)))
	if v3121 != 0 {
		v3079 = v3120
		v3080 = v3117
		goto L491
	} else {
		goto L494
	}
L494:
	;
	goto L492
L495:
	;
	v3188 = *(*int32)(unsafe.Add(mBase, uint32(v3128)+24))
	if v3188 != 0 {
		goto L508
	} else {
		goto L509
	}
L496:
	;
	v3156 = *(*int32)(unsafe.Add(mBase, uint32(v2473)+uint32(_c_F_WalSummarizerMain[22])))
	v3157 = *(*int32)(unsafe.Add(mBase, uint32(v3128)+32))
	v3159 = v3153 << (uint(int32(1)) % 32)
	v3160 = m.Env.Pgmem_crc32c(m, v3156, v3157, v3159)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v2473)+uint32(_c_F_WalSummarizerMain[22]))) = v3160
	v3162 = *(*int32)(unsafe.Add(mBase, uint32(v2473)+uint32(_c_F_WalSummarizerMain[23])))
	if int32(_a_F_WalSummarizerMain_31) <= v3162+v3159 {
		goto L497
	} else {
		goto L498
	}
L497:
	;
	v3166 = *(*int32)(unsafe.Add(mBase, uint32(v2473)+16))
	v3167 = *(*int32)(unsafe.Add(mBase, uint32(v2473)+12))
	v3168 = m.T0[v3167].(func(*base.Module, int32, int32, int32) int32)(m, v3166, v2583, v3162)
	mBase = m.M
	v3169 = m.ExcPending
	if v3169 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L500
	}
L498:
	;
	v3173 = v3162
	goto L499
L499:
	;
	if int32(_a_F_WalSummarizerMain_34) <= v3159 {
		goto L501
	} else {
		goto L502
	}
L500:
	;
	v3170 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2473)+uint32(_c_F_WalSummarizerMain[23]))) = v3170
	v3173 = v3170
	goto L499
L501:
	;
	v3176 = *(*int32)(unsafe.Add(mBase, uint32(v2473)+16))
	v3177 = *(*int32)(unsafe.Add(mBase, uint32(v2473)+12))
	v3178 = m.T0[v3177].(func(*base.Module, int32, int32, int32) int32)(m, v3176, v3157, v3159)
	mBase = m.M
	v3179 = m.ExcPending
	if v3179 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L504
	}
L502:
	;
	goto L503
L503:
	;
	if v3159 != 0 {
		goto L505
	} else {
		goto L506
	}
L504:
	;
	goto L495
L505:
	;
	base.MemoryCopy(m, v3173+v2583, v3157, v3159)
	goto L507
L506:
	;
	goto L507
L507:
	;
	v3182 = *(*int32)(unsafe.Add(mBase, uint32(v2473)+uint32(_c_F_WalSummarizerMain[23])))
	*(*int32)(unsafe.Add(mBase, uint32(v2473)+uint32(_c_F_WalSummarizerMain[23]))) = v3182 + v3159
	goto L495
L508:
	;
	v3196 = int32(0)
	goto L511
L509:
	;
	goto L510
L510:
	;
	v3299 = v2733 + int32(1)
	v3300 = *(*int32)(unsafe.Add(mBase, uint32(v919)))
	v3301 = *(*int32)(unsafe.Add(mBase, uint32(v3300)+8))
	if base.Ui32(v3299) < base.Ui32(v3301) {
		v2733 = v3299
		goto L441
	} else {
		goto L527
	}
L511:
	;
	v3220 = *(*int32)(unsafe.Add(mBase, uint32(v3128)+32))
	v3224 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3220+v3196<<(uint(int32(1))%32)))))
	if v3224 == int32(0) {
		goto L513
	} else {
		goto L514
	}
L512:
	;
	goto L510
L513:
	;
	v3265 = v3196 + int32(1)
	v3266 = *(*int32)(unsafe.Add(mBase, uint32(v3128)+24))
	if base.Ui32(v3265) < base.Ui32(v3266) {
		v3196 = v3265
		goto L511
	} else {
		goto L526
	}
L514:
	;
	v3228 = *(*int32)(unsafe.Add(mBase, uint32(v2473)+uint32(_c_F_WalSummarizerMain[22])))
	v3229 = *(*int32)(unsafe.Add(mBase, uint32(v3128)+36))
	v3233 = *(*int32)(unsafe.Add(mBase, uint32(v3229+v3196<<(uint(int32(2))%32))))
	v3235 = v3224 << (uint(int32(1)) % 32)
	v3236 = m.Env.Pgmem_crc32c(m, v3228, v3233, v3235)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v2473)+uint32(_c_F_WalSummarizerMain[22]))) = v3236
	v3238 = *(*int32)(unsafe.Add(mBase, uint32(v2473)+uint32(_c_F_WalSummarizerMain[23])))
	if int32(_a_F_WalSummarizerMain_31) <= v3238+v3235 {
		goto L515
	} else {
		goto L516
	}
L515:
	;
	v3242 = *(*int32)(unsafe.Add(mBase, uint32(v2473)+16))
	v3243 = *(*int32)(unsafe.Add(mBase, uint32(v2473)+12))
	v3244 = m.T0[v3243].(func(*base.Module, int32, int32, int32) int32)(m, v3242, v2583, v3238)
	mBase = m.M
	v3245 = m.ExcPending
	if v3245 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L518
	}
L516:
	;
	v3249 = v3238
	goto L517
L517:
	;
	if base.I32_extend16_s(v3224) < int32(0) {
		goto L519
	} else {
		goto L520
	}
L518:
	;
	v3246 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2473)+uint32(_c_F_WalSummarizerMain[23]))) = v3246
	v3249 = v3246
	goto L517
L519:
	;
	v3252 = *(*int32)(unsafe.Add(mBase, uint32(v2473)+16))
	v3253 = *(*int32)(unsafe.Add(mBase, uint32(v2473)+12))
	v3254 = m.T0[v3253].(func(*base.Module, int32, int32, int32) int32)(m, v3252, v3233, v3235)
	mBase = m.M
	v3255 = m.ExcPending
	if v3255 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L522
	}
L520:
	;
	goto L521
L521:
	;
	if v3235 != 0 {
		goto L523
	} else {
		goto L524
	}
L522:
	;
	goto L513
L523:
	;
	base.MemoryCopy(m, v3249+v2583, v3233, v3235)
	goto L525
L524:
	;
	goto L525
L525:
	;
	v3258 = *(*int32)(unsafe.Add(mBase, uint32(v2473)+uint32(_c_F_WalSummarizerMain[23])))
	*(*int32)(unsafe.Add(mBase, uint32(v2473)+uint32(_c_F_WalSummarizerMain[23]))) = v3258 + v3235
	goto L513
L526:
	;
	goto L512
L527:
	;
	goto L442
L528:
	;
	v3356 = *(*int32)(unsafe.Add(mBase, uint32(v3344)+4))
	v3359 = *(*int32)(unsafe.Add(mBase, uint32(v3344)))
	v3360 = m.T0[v3359].(func(*base.Module, int32, int32, int32) int32)(m, v3356, v2473+int32(20), v3351)
	mBase = m.M
	v3361 = m.ExcPending
	if v3361 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L531
	}
L529:
	;
	v3365 = v3351
	goto L530
L530:
	;
	v3367 = v2473 + int32(20)
	v3368 = v3365 + v3367
	v3369 = *(*int64)(unsafe.Add(mBase, uint32(v3335)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v3368)+16)) = v3369
	v3371 = *(*int64)(unsafe.Add(mBase, uint32(v3335)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v3368)+8)) = v3371
	v3373 = *(*int64)(unsafe.Add(mBase, uint32(v3335)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v3368))) = v3373
	v3375 = *(*int32)(unsafe.Add(mBase, uint32(v2473)+uint32(_c_F_WalSummarizerMain[23])))
	*(*int32)(unsafe.Add(mBase, uint32(v2473)+uint32(_c_F_WalSummarizerMain[23]))) = v3375 + int32(24)
	v3379 = *(*int32)(unsafe.Add(mBase, uint32(v2473)+uint32(_c_F_WalSummarizerMain[22])))
	*(*int32)(unsafe.Add(mBase, uint32(v3335)+4)) = v3379 ^ int32(-1)
	v3383 = int32(4)
	v3386 = m.Env.Pgmem_crc32c(m, v3379, v3335+v3383, v3383)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v2473)+uint32(_c_F_WalSummarizerMain[22]))) = v3386
	v3388 = *(*int32)(unsafe.Add(mBase, uint32(v2473)+uint32(_c_F_WalSummarizerMain[23])))
	if int32(_a_F_WalSummarizerMain_31) <= v3388+v3383 {
		goto L532
	} else {
		goto L533
	}
L531:
	;
	v3362 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2473)+uint32(_c_F_WalSummarizerMain[23]))) = v3362
	v3365 = v3362
	goto L530
L532:
	;
	v3393 = *(*int32)(unsafe.Add(mBase, uint32(v3344)+4))
	v3394 = *(*int32)(unsafe.Add(mBase, uint32(v3344)))
	v3395 = m.T0[v3394].(func(*base.Module, int32, int32, int32) int32)(m, v3393, v3367, v3388)
	mBase = m.M
	v3396 = m.ExcPending
	if v3396 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L535
	}
L533:
	;
	v3400 = v3388
	goto L534
L534:
	;
	v3402 = *(*int32)(unsafe.Add(mBase, uint32(v3335)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v3400+v3367))) = v3402
	v3404 = *(*int32)(unsafe.Add(mBase, uint32(v2473)+uint32(_c_F_WalSummarizerMain[23])))
	v3406 = v3404 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v2473)+uint32(_c_F_WalSummarizerMain[23]))) = v3406
	v3408 = *(*int32)(unsafe.Add(mBase, uint32(v3344)+4))
	v3409 = *(*int32)(unsafe.Add(mBase, uint32(v3344)))
	v3410 = m.T0[v3409].(func(*base.Module, int32, int32, int32) int32)(m, v3408, v3367, v3406)
	mBase = m.M
	v3411 = m.ExcPending
	if v3411 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L536
	}
L535:
	;
	v3397 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2473)+uint32(_c_F_WalSummarizerMain[23]))) = v3397
	v3400 = v3397
	goto L534
L536:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2473)+uint32(_c_F_WalSummarizerMain[23]))) = int32(0)
	m.G0 = v3335 + int32(32)
	m.G0 = v2473 + int32(_a_F_WalSummarizerMain_22)
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	v3422 = *(*int32)(unsafe.Add(mBase, uint32(v519)+464))
	F_FileClose(m, v3422)
	mBase = m.M
	v3424 = m.ExcPending
	if v3424 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L537
	}
L537:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	v3429 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v3430 = m.ExcPending
	if v3430 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L538
	}
L538:
	;
	if v3429 != 0 {
		goto L539
	} else {
		goto L540
	}
L539:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	*(*int32)(unsafe.Add(mBase, uint32(v519-int32(-64)))) = v2360
	*(*int32)(unsafe.Add(mBase, uint32(v519)+60)) = v2364
	*(*int32)(unsafe.Add(mBase, uint32(v519)+56)) = v2366
	*(*int32)(unsafe.Add(mBase, uint32(v519)+52)) = v2370
	*(*int32)(unsafe.Add(mBase, uint32(v519)+48)) = v838
	F_errmsg_internal(m, int32(_a_F_WalSummarizerMain_35), v519+int32(48))
	mBase = m.M
	v3444 = m.ExcPending
	if v3444 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L542
	}
L540:
	;
	goto L541
L541:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	v3459 = F_durable_rename(m, v519+int32(1504), v519+int32(480), int32(21))
	mBase = m.M
	v3460 = m.ExcPending
	if v3460 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L544
	}
L542:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	F_errfinish(m, int32(_a_F_WalSummarizerMain_1), int32(1228), int32(_a_F_WalSummarizerMain_21))
	mBase = m.M
	v3451 = m.ExcPending
	if v3451 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L543
	}
L543:
	;
	goto L541
L544:
	;
	goto L174
L545:
	;
	v3525 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v519)+267)) = uint8(v3525)
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	v3530 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[8]))
	v3534 = F_LWLockAcquire(m, v3530+int32(_a_F_WalSummarizerMain_4), int32(0))
	mBase = m.M
	v3535 = m.ExcPending
	if v3535 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L551
	}
L546:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	v3500 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v3501 = m.ExcPending
	if v3501 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L547
	}
L547:
	;
	if v3500 == int32(0) {
		goto L545
	} else {
		goto L548
	}
L548:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	*(*uint32)(unsafe.Add(mBase, uint32(v519)+16)) = uint32(v2331)
	v3507 = int64(32)
	v3508 = int64(base.Ui64(v2331) >> (uint(v3507) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v519)+12)) = uint32(v3508)
	*(*int32)(unsafe.Add(mBase, uint32(v519))) = v838
	*(*uint32)(unsafe.Add(mBase, uint32(v519)+8)) = uint32(v1414)
	v3513 = int64(base.Ui64(v1414) >> (uint(v3507) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v519)+4)) = uint32(v3513)
	F_errmsg_internal(m, int32(_a_F_WalSummarizerMain_36), v519)
	mBase = m.M
	v3517 = m.ExcPending
	if v3517 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L549
	}
L549:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	F_errfinish(m, int32(_a_F_WalSummarizerMain_1), int32(1240), int32(_a_F_WalSummarizerMain_21))
	mBase = m.M
	v3524 = m.ExcPending
	if v3524 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L550
	}
L550:
	;
	goto L545
L551:
	;
	v3537 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[9]))
	*(*int64)(unsafe.Add(mBase, uint32(v3537)+24)) = v2331
	v3539 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3537)+16)) = uint8(v3539)
	*(*int32)(unsafe.Add(mBase, uint32(v3537)+4)) = v838
	*(*int64)(unsafe.Add(mBase, uint32(v3537)+8)) = v2331
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	v3546 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[8]))
	F_LWLockRelease(m, v3546+int32(_a_F_WalSummarizerMain_4))
	mBase = m.M
	v3550 = m.ExcPending
	if v3550 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L552
	}
L552:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v519)+3120)) = v1411
	*(*int32)(unsafe.Add(mBase, uint32(v519)+3132)) = v520
	v3554 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[9]))
	F_ConditionVariableBroadcast(m, v3554+int32(32))
	mBase = m.M
	v3558 = m.ExcPending
	if v3558 != 0 {
		v3596 = v519
		v3610 = v533
		goto L6
	} else {
		goto L553
	}
L553:
	;
	v3582 = v1411
	v3583 = v883
	v3586 = v2331
	goto L157
L554:
	;
	goto L5
L555:
	;
	v3629 = int32(v3625)
	m.G0 = v3596
	v3631 = *(*int32)(unsafe.Add(mBase, uint32(v3629)+4))
	v3632 = *(*int32)(unsafe.Add(mBase, uint32(v3629)))
	v3635 = *(*int32)(unsafe.Add(mBase, uint32(v3632)))
	if v3596+int32(252) == v3635 {
		goto L558
	} else {
		goto L559
	}
L556:
	;
	m.ExcPending = 1
	goto L564
L557:
	;
	if v3639 != 0 {
		goto L561
	} else {
		goto L562
	}
L558:
	;
	v3637 = *(*int32)(unsafe.Add(mBase, uint32(v3632)+4))
	v3639 = v3637
	goto L560
L559:
	;
	v3639 = int32(0)
	goto L560
L560:
	;
	goto L557
L561:
	;
	v3640 = *(*int32)(unsafe.Add(mBase, uint32(v3596)+3132))
	v3641 = *(*int64)(unsafe.Add(mBase, uint32(v3596)+3120))
	v38 = v3639
	v40 = v3596
	v41 = v3640
	v43 = v3631
	v54 = v3610
	v61 = v3641
	goto L1
L562:
	;
	goto L563
L563:
	;
	F___wasm_longjmp(m, v3632, v3631)
	mBase = m.M
	v3643 = m.ExcPending
	if v3643 != 0 {
		goto L564
	} else {
		goto L565
	}
L564:
	;
	return
L565:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
