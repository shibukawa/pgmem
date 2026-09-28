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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+300))
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
									F_errfinish(m, int32(_a_F_ProcessWalSummarizerInterrupts_1), int32(968), int32(_a_F_ProcessWalSummarizerInterrupts_2))
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
								F_errfinish(m, int32(_a_F_ProcessWalSummarizerInterrupts_1), int32(968), int32(_a_F_ProcessWalSummarizerInterrupts_2))
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
								F_errfinish(m, int32(_a_F_ProcessWalSummarizerInterrupts_1), int32(968), int32(_a_F_ProcessWalSummarizerInterrupts_2))
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
							F_errfinish(m, int32(_a_F_ProcessWalSummarizerInterrupts_1), int32(968), int32(_a_F_ProcessWalSummarizerInterrupts_2))
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
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	v1 = l0
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_SetWalWriterSleeping[0]))
	v7 = base.AtomicRmwXchg32(m, v4, int32(440), int32(1))
	if v7 != 0 {
		F_s_lock(m, v4+int32(440), int32(_a_F_SetWalWriterSleeping_0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, _c_F_SetWalWriterSleeping[0]))
			*(*uint8)(unsafe.Add(mBase, uint32(v14)+313)) = uint8(v1)
			v16 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v14)+440)), uint32(v16))
			return
		}
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_SetWalWriterSleeping[0]))
		*(*uint8)(unsafe.Add(mBase, uint32(v14)+313)) = uint8(v1)
		v16 = int32(0)
		atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v14)+440)), uint32(v16))
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
func F_WalRcvRequestApplyReply(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_WalRcvRequestApplyReply[0]))
	v5 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4)+1472)) = v5
	v9 = base.AtomicRmwXchg32(m, v4, int32(1456), v5)
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_s_lock(m, v4+int32(1456), int32(_a_F_WalRcvRequestApplyReply_0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_WalRcvRequestApplyReply[0]))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v18 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v16)+1456)), uint32(v18))
	if v17 != int32(-1) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_WalRcvRequestApplyReply[1]))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v30 = v25 + v17*int32(768) + int32(316)
	v31 = int32(0)
	v34 = base.AtomicRmwOr32(m, v31, int32(_a_F_WalRcvRequestApplyReply_1), v31)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if v35 != 0 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	goto L8
L8:
	;
	return
L9:
	;
	goto L8
L10:
	;
	goto L9
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = int32(1)
	v38 = int32(0)
	v41 = base.AtomicRmwOr32(m, v38, int32(_a_F_WalRcvRequestApplyReply_1), v38)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v42 == v38 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	if v45 == int32(0) {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_WalRcvRequestApplyReply[2]))
	if v49 == v45 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v51 = m.G0
	v53 = v51 - int32(16)
	m.G0 = v53
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_WalRcvRequestApplyReply[3]))
	if v56 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v79 = F_pgmem_kill(m, v45, int32(23))
	mBase = m.M
	goto L10
L17:
	;
	m.G0 = v53 + int32(16)
	goto L9
L18:
	;
	v59 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+15)) = uint8(v59)
	goto L19
L19:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_WalRcvRequestApplyReply[4]))
	v67 = F_write(m, v63, v53+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v67 {
		goto L17
	} else {
		goto L21
	}
L20:
	;
	goto L17
L21:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_WalRcvRequestApplyReply[5]))
	if v71 == int32(27) {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
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
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_WalRcvRunning[0]))
	v7 = int32(1456)
	v8 = v6 + v7
	v11 = base.AtomicRmwXchg32(m, v6, v7, int32(1))
	if v11 != 0 {
		F_s_lock(m, v8, int32(_a_F_WalRcvRunning_0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*int64)(unsafe.Add(mBase, uint32(v6)+24))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
			v19 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v6)+1456)), uint32(v19))
			if v18 != int32(1) {
				v51 = v18
				return base.B2i32(v51 != int32(0))
			} else {
				v24 = int32(1)
				v25 = F_time(m)
				mBase = m.M
				if v25-v17 < int64(11) {
					v51 = v24
					return base.B2i32(v51 != int32(0))
				} else {
					v31 = base.AtomicRmwXchg32(m, v8, int32(0), int32(1))
					if v31 != 0 {
						F_s_lock(m, v8, int32(_a_F_WalRcvRunning_0))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
							if v35 != int32(1) {
								v38 = int32(0)
								atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v8))), uint32(v38))
								v51 = v24
								return base.B2i32(v51 != int32(0))
							} else {
								v41 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v41
								atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v6)+1456)), uint32(v41))
								F_ConditionVariableBroadcast(m, v6+int32(12))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int32(0)
								} else {
									v51 = v41
									return base.B2i32(v51 != int32(0))
								}
							}
						}
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
						if v35 != int32(1) {
							v38 = int32(0)
							atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v8))), uint32(v38))
							v51 = v24
							return base.B2i32(v51 != int32(0))
						} else {
							v41 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v41
							atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v6)+1456)), uint32(v41))
							F_ConditionVariableBroadcast(m, v6+int32(12))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								v51 = v41
								return base.B2i32(v51 != int32(0))
							}
						}
					}
				}
			}
		}
	} else {
		v17 = *(*int64)(unsafe.Add(mBase, uint32(v6)+24))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
		v19 = int32(0)
		atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v6)+1456)), uint32(v19))
		if v18 != int32(1) {
			v51 = v18
			return base.B2i32(v51 != int32(0))
		} else {
			v24 = int32(1)
			v25 = F_time(m)
			mBase = m.M
			if v25-v17 < int64(11) {
				v51 = v24
				return base.B2i32(v51 != int32(0))
			} else {
				v31 = base.AtomicRmwXchg32(m, v8, int32(0), int32(1))
				if v31 != 0 {
					F_s_lock(m, v8, int32(_a_F_WalRcvRunning_0))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
						if v35 != int32(1) {
							v38 = int32(0)
							atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v8))), uint32(v38))
							v51 = v24
							return base.B2i32(v51 != int32(0))
						} else {
							v41 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v41
							atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v6)+1456)), uint32(v41))
							F_ConditionVariableBroadcast(m, v6+int32(12))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								v51 = v41
								return base.B2i32(v51 != int32(0))
							}
						}
					}
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
					if v35 != int32(1) {
						v38 = int32(0)
						atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v8))), uint32(v38))
						v51 = v24
						return base.B2i32(v51 != int32(0))
					} else {
						v41 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v41
						atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v6)+1456)), uint32(v41))
						F_ConditionVariableBroadcast(m, v6+int32(12))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							v51 = v41
							return base.B2i32(v51 != int32(0))
						}
					}
				}
			}
		}
	}
}
func F_WalSndShmemRequest(m *base.Module, l0 int32) {
	var v8 int32
	_ = v8
	Fn14205(m, l0, int32(_a_F_WalSndShmemRequest_0), int32(_a_F_WalSndShmemRequest_1), int32(96), int32(_a_F_WalSndShmemRequest_2), int32(88))
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
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
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v59 int32
	_ = v59
	var v65 int64
	_ = v65
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v181 int32
	_ = v181
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v225 int32
	_ = v225
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v269 int32
	_ = v269
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v313 int32
	_ = v313
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v357 int32
	_ = v357
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v401 int32
	_ = v401
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v478 int32
	_ = v478
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v576 int32
	_ = v576
	var v583 int64
	_ = v583
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v604 int32
	_ = v604
	var v615 int64
	_ = v615
	var v617 int64
	_ = v617
	var v619 int64
	_ = v619
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v632 int64
	_ = v632
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v639 int64
	_ = v639
	var v645 int64
	_ = v645
	var v649 int32
	_ = v649
	var v651 int64
	_ = v651
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v665 int32
	_ = v665
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v702 int64
	_ = v702
	var v703 int32
	_ = v703
	var v706 int64
	_ = v706
	var v707 int64
	_ = v707
	var v712 int32
	_ = v712
	var v717 int32
	_ = v717
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v757 int64
	_ = v757
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v766 int64
	_ = v766
	var v769 int64
	_ = v769
	var v773 int64
	_ = v773
	var v774 int64
	_ = v774
	var v777 int64
	_ = v777
	var v780 int32
	_ = v780
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v807 int32
	_ = v807
	var v812 int32
	_ = v812
	var v813 int64
	_ = v813
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v827 int32
	_ = v827
	var v832 int32
	_ = v832
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v850 int32
	_ = v850
	var v855 int32
	_ = v855
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v875 int32
	_ = v875
	var v917 int64
	_ = v917
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v930 int32
	_ = v930
	var v934 int32
	_ = v934
	var v937 int32
	_ = v937
	var v953 int32
	_ = v953
	var v972 int32
	_ = v972
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v978 int32
	_ = v978
	var v980 int64
	_ = v980
	var v1020 int32
	_ = v1020
	var v1028 int32
	_ = v1028
	var v1035 int32
	_ = v1035
	var v1043 int32
	_ = v1043
	var v1051 int32
	_ = v1051
	var v1058 int32
	_ = v1058
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1076 int32
	_ = v1076
	var v1082 int32
	_ = v1082
	var v1104 int32
	_ = v1104
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1117 int32
	_ = v1117
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1128 int32
	_ = v1128
	var v1130 int32
	_ = v1130
	var v1138 int32
	_ = v1138
	var v1166 int32
	_ = v1166
	var v1169 int32
	_ = v1169
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1216 int32
	_ = v1216
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1225 int32
	_ = v1225
	var v1228 int32
	_ = v1228
	var v1233 int64
	_ = v1233
	var v1239 int32
	_ = v1239
	var v1246 int32
	_ = v1246
	var v1250 int32
	_ = v1250
	var v1264 int32
	_ = v1264
	var v1276 int64
	_ = v1276
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1289 int32
	_ = v1289
	var v1293 int32
	_ = v1293
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1301 int32
	_ = v1301
	var v1303 int32
	_ = v1303
	var v1305 int32
	_ = v1305
	var v1311 int32
	_ = v1311
	var v1315 int32
	_ = v1315
	var v1320 int32
	_ = v1320
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1345 int64
	_ = v1345
	var v1351 int64
	_ = v1351
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1378 int32
	_ = v1378
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1390 int32
	_ = v1390
	var v1395 int32
	_ = v1395
	var v1401 int32
	_ = v1401
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1414 int32
	_ = v1414
	var v1420 int32
	_ = v1420
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1430 int64
	_ = v1430
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1464 int64
	_ = v1464
	var v1472 int32
	_ = v1472
	var v1473 int32
	_ = v1473
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1481 int32
	_ = v1481
	var v1484 int32
	_ = v1484
	var v1488 int32
	_ = v1488
	var v1499 int64
	_ = v1499
	var v1500 int32
	_ = v1500
	var v1505 int32
	_ = v1505
	var v1533 int32
	_ = v1533
	var v1535 int32
	_ = v1535
	var v1539 int32
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1541 int32
	_ = v1541
	var v1574 int64
	_ = v1574
	var v1576 int32
	_ = v1576
	var v1578 int32
	_ = v1578
	var v1581 int32
	_ = v1581
	var v1624 int32
	_ = v1624
	var v1625 int32
	_ = v1625
	var v1628 int64
	_ = v1628
	var v1630 int32
	_ = v1630
	var v1635 int32
	_ = v1635
	var v1663 int32
	_ = v1663
	var v1665 int32
	_ = v1665
	var v1669 int32
	_ = v1669
	var v1670 int32
	_ = v1670
	var v1671 int32
	_ = v1671
	var v1704 int64
	_ = v1704
	var v1706 int32
	_ = v1706
	var v1708 int32
	_ = v1708
	var v1711 int32
	_ = v1711
	var v1723 int64
	_ = v1723
	var v1725 int32
	_ = v1725
	var v1726 int32
	_ = v1726
	var v1761 int32
	_ = v1761
	var v1767 int32
	_ = v1767
	var v1770 int32
	_ = v1770
	var v1771 int32
	_ = v1771
	var v1773 int32
	_ = v1773
	var v1809 int64
	_ = v1809
	var v1812 int32
	_ = v1812
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1821 int64
	_ = v1821
	var v1825 int64
	_ = v1825
	var v1826 int64
	_ = v1826
	var v1831 int64
	_ = v1831
	var v1837 int32
	_ = v1837
	var v1844 int32
	_ = v1844
	var v1846 int64
	_ = v1846
	var v1872 int64
	_ = v1872
	var v1874 int64
	_ = v1874
	var v1879 int64
	_ = v1879
	var v1883 int32
	_ = v1883
	var v1888 int32
	_ = v1888
	var v1889 int32
	_ = v1889
	var v1893 int32
	_ = v1893
	var v1897 int32
	_ = v1897
	var v1900 int32
	_ = v1900
	var v1901 int32
	_ = v1901
	var v1904 int32
	_ = v1904
	var v1912 int32
	_ = v1912
	var v1919 int32
	_ = v1919
	var v1928 int32
	_ = v1928
	var v1935 int32
	_ = v1935
	var v1937 int32
	_ = v1937
	var v1970 int32
	_ = v1970
	var v1971 int64
	_ = v1971
	var v1973 int32
	_ = v1973
	var v1974 int32
	_ = v1974
	var v1977 int32
	_ = v1977
	var v1979 int32
	_ = v1979
	var v1986 int32
	_ = v1986
	var v1993 int32
	_ = v1993
	var v1994 int32
	_ = v1994
	var v1997 int32
	_ = v1997
	var v2000 int32
	_ = v2000
	var v2001 int64
	_ = v2001
	var v2003 int32
	_ = v2003
	var v2005 int32
	_ = v2005
	var v2010 int32
	_ = v2010
	var v2015 int32
	_ = v2015
	var v2017 int32
	_ = v2017
	var v2020 int32
	_ = v2020
	var v2021 int32
	_ = v2021
	var v2023 int32
	_ = v2023
	var v2024 int32
	_ = v2024
	var v2034 int32
	_ = v2034
	var v2035 int32
	_ = v2035
	var v2036 int32
	_ = v2036
	var v2038 int32
	_ = v2038
	var v2039 int32
	_ = v2039
	var v2049 int32
	_ = v2049
	var v2050 int32
	_ = v2050
	var v2051 int32
	_ = v2051
	var v2052 int32
	_ = v2052
	var v2055 int32
	_ = v2055
	var v2065 int32
	_ = v2065
	var v2096 int32
	_ = v2096
	var v2102 int32
	_ = v2102
	var v2105 int32
	_ = v2105
	var v2107 int32
	_ = v2107
	var v2108 int32
	_ = v2108
	var v2110 int32
	_ = v2110
	var v2114 int32
	_ = v2114
	var v2117 int32
	_ = v2117
	var v2118 int32
	_ = v2118
	var v2125 int32
	_ = v2125
	var v2126 int32
	_ = v2126
	var v2127 int32
	_ = v2127
	var v2130 int32
	_ = v2130
	var v2137 int32
	_ = v2137
	var v2138 int32
	_ = v2138
	var v2140 int32
	_ = v2140
	var v2145 int32
	_ = v2145
	var v2156 int32
	_ = v2156
	var v2158 int32
	_ = v2158
	var v2159 int32
	_ = v2159
	var v2164 int32
	_ = v2164
	var v2168 int32
	_ = v2168
	var v2169 int32
	_ = v2169
	var v2175 int64
	_ = v2175
	var v2180 int32
	_ = v2180
	var v2184 int32
	_ = v2184
	var v2186 int32
	_ = v2186
	var v2192 int32
	_ = v2192
	var v2195 int32
	_ = v2195
	var v2197 int32
	_ = v2197
	var v2203 int32
	_ = v2203
	var v2207 int32
	_ = v2207
	var v2209 int32
	_ = v2209
	var v2212 int32
	_ = v2212
	var v2216 int32
	_ = v2216
	var v2221 int32
	_ = v2221
	var v2222 int32
	_ = v2222
	var v2223 int32
	_ = v2223
	var v2226 int32
	_ = v2226
	var v2230 int32
	_ = v2230
	var v2235 int32
	_ = v2235
	var v2236 int32
	_ = v2236
	var v2237 int32
	_ = v2237
	var v2240 int32
	_ = v2240
	var v2244 int32
	_ = v2244
	var v2251 int32
	_ = v2251
	var v2254 int32
	_ = v2254
	var v2262 int32
	_ = v2262
	var v2263 int32
	_ = v2263
	var v2267 int32
	_ = v2267
	var v2268 int32
	_ = v2268
	var v2269 int32
	_ = v2269
	var v2274 int64
	_ = v2274
	var v2275 int64
	_ = v2275
	var v2283 int32
	_ = v2283
	var v2284 int32
	_ = v2284
	var v2291 int32
	_ = v2291
	var v2322 int32
	_ = v2322
	var v2323 int32
	_ = v2323
	var v2325 int32
	_ = v2325
	var v2328 int32
	_ = v2328
	var v2331 int32
	_ = v2331
	var v2336 int32
	_ = v2336
	var v2339 int32
	_ = v2339
	var v2344 int32
	_ = v2344
	var v2346 int32
	_ = v2346
	var v2347 int32
	_ = v2347
	var v2349 int32
	_ = v2349
	var v2353 int32
	_ = v2353
	var v2354 int32
	_ = v2354
	var v2360 int64
	_ = v2360
	var v2365 int32
	_ = v2365
	var v2369 int32
	_ = v2369
	var v2371 int32
	_ = v2371
	var v2377 int32
	_ = v2377
	var v2380 int32
	_ = v2380
	var v2382 int32
	_ = v2382
	var v2388 int32
	_ = v2388
	var v2392 int32
	_ = v2392
	var v2394 int32
	_ = v2394
	var v2397 int32
	_ = v2397
	var v2401 int32
	_ = v2401
	var v2406 int32
	_ = v2406
	var v2407 int32
	_ = v2407
	var v2408 int32
	_ = v2408
	var v2411 int32
	_ = v2411
	var v2415 int32
	_ = v2415
	var v2422 int32
	_ = v2422
	var v2425 int32
	_ = v2425
	var v2433 int32
	_ = v2433
	var v2434 int32
	_ = v2434
	var v2438 int32
	_ = v2438
	var v2439 int32
	_ = v2439
	var v2440 int32
	_ = v2440
	var v2445 int64
	_ = v2445
	var v2446 int64
	_ = v2446
	var v2454 int32
	_ = v2454
	var v2455 int32
	_ = v2455
	var v2462 int32
	_ = v2462
	var v2493 int32
	_ = v2493
	var v2494 int32
	_ = v2494
	var v2496 int32
	_ = v2496
	var v2499 int32
	_ = v2499
	var v2502 int32
	_ = v2502
	var v2507 int32
	_ = v2507
	var v2510 int32
	_ = v2510
	var v2515 int32
	_ = v2515
	var v2517 int32
	_ = v2517
	var v2518 int32
	_ = v2518
	var v2552 int32
	_ = v2552
	var v2553 int32
	_ = v2553
	var v2554 int32
	_ = v2554
	var v2561 int32
	_ = v2561
	var v2592 int32
	_ = v2592
	var v2594 int32
	_ = v2594
	var v2596 int32
	_ = v2596
	var v2598 int32
	_ = v2598
	var v2599 int32
	_ = v2599
	var v2601 int32
	_ = v2601
	var v2602 int32
	_ = v2602
	var v2606 int32
	_ = v2606
	var v2607 int32
	_ = v2607
	var v2611 int32
	_ = v2611
	var v2612 int32
	_ = v2612
	var v2614 int64
	_ = v2614
	var v2616 int32
	_ = v2616
	var v2618 int32
	_ = v2618
	var v2626 int32
	_ = v2626
	var v2629 int32
	_ = v2629
	var v2634 int32
	_ = v2634
	var v2636 int32
	_ = v2636
	var v2639 int32
	_ = v2639
	var v2640 int32
	_ = v2640
	var v2641 int32
	_ = v2641
	var v2648 int32
	_ = v2648
	var v2676 int64
	_ = v2676
	var v2680 int32
	_ = v2680
	var v2684 int32
	_ = v2684
	var v2685 int32
	_ = v2685
	var v2687 int32
	_ = v2687
	var v2692 int32
	_ = v2692
	var v2696 int32
	_ = v2696
	var v2699 int64
	_ = v2699
	var v2704 int32
	_ = v2704
	var v2709 int32
	_ = v2709
	var v2710 int32
	_ = v2710
	var v2715 int32
	_ = v2715
	var v2743 int32
	_ = v2743
	var v2750 int32
	_ = v2750
	var v2751 int32
	_ = v2751
	var v2752 int64
	_ = v2752
	var v2753 int64
	_ = v2753
	var v2757 int64
	_ = v2757
	var v2758 int64
	_ = v2758
	var v2762 int64
	_ = v2762
	var v2769 int32
	_ = v2769
	var v2776 int32
	_ = v2776
	var v2779 int64
	_ = v2779
	var v2784 int32
	_ = v2784
	var v2806 int64
	_ = v2806
	var v2812 int32
	_ = v2812
	var v2838 int64
	_ = v2838
	var v2844 int32
	_ = v2844
	var v2848 int32
	_ = v2848
	var v2852 int32
	_ = v2852
	var v2853 int32
	_ = v2853
	var v2860 int32
	_ = v2860
	var v2864 int32
	_ = v2864
	var v2865 int32
	_ = v2865
	var v2868 int32
	_ = v2868
	var v2870 int64
	_ = v2870
	var v2872 int32
	_ = v2872
	var v2874 int32
	_ = v2874
	var v2878 int32
	_ = v2878
	var v2887 int32
	_ = v2887
	var v2888 int32
	_ = v2888
	var v2894 int32
	_ = v2894
	var v2895 int32
	_ = v2895
	var v2904 int32
	_ = v2904
	var v2908 int32
	_ = v2908
	var v2916 int32
	_ = v2916
	var v2923 int32
	_ = v2923
	var v2926 int32
	_ = v2926
	var v2930 int32
	_ = v2930
	var v2934 int32
	_ = v2934
	var v2935 int64
	_ = v2935
	var v2938 int32
	_ = v2938
	var v2939 int32
	_ = v2939
	var v2942 int32
	_ = v2942
	var v2951 int32
	_ = v2951
	var v2958 int32
	_ = v2958
	var v2968 int32
	_ = v2968
	var v2975 int32
	_ = v2975
	var v2978 int32
	_ = v2978
	var v2979 int32
	_ = v2979
	var v2981 int32
	_ = v2981
	var v2995 int32
	_ = v2995
	var v2999 int32
	_ = v2999
	var v3000 int32
	_ = v3000
	var v3002 int32
	_ = v3002
	var v3005 int32
	_ = v3005
	var v3007 int32
	_ = v3007
	var v3011 int32
	_ = v3011
	var v3012 int32
	_ = v3012
	var v3013 int32
	_ = v3013
	var v3014 int64
	_ = v3014
	var v3017 int32
	_ = v3017
	var v3026 int32
	_ = v3026
	var v3054 int32
	_ = v3054
	var v3058 int32
	_ = v3058
	var v3065 int32
	_ = v3065
	var v3094 int32
	_ = v3094
	var v3105 int32
	_ = v3105
	var v3106 int32
	_ = v3106
	var v3111 int32
	_ = v3111
	var v3113 int32
	_ = v3113
	var v3133 int32
	_ = v3133
	var v3135 int32
	_ = v3135
	var v3138 int32
	_ = v3138
	var v3162 int32
	_ = v3162
	var v3163 int32
	_ = v3163
	var v3164 int32
	_ = v3164
	var v3168 int32
	_ = v3168
	var v3169 int32
	_ = v3169
	var v3172 int32
	_ = v3172
	var v3173 int32
	_ = v3173
	var v3174 int32
	_ = v3174
	var v3175 int32
	_ = v3175
	var v3180 int32
	_ = v3180
	var v3181 int32
	_ = v3181
	var v3183 int64
	_ = v3183
	var v3185 int32
	_ = v3185
	var v3187 int32
	_ = v3187
	var v3189 int32
	_ = v3189
	var v3197 int32
	_ = v3197
	var v3225 int32
	_ = v3225
	var v3231 int32
	_ = v3231
	var v3236 int32
	_ = v3236
	var v3240 int32
	_ = v3240
	var v3241 int32
	_ = v3241
	var v3242 int32
	_ = v3242
	var v3252 int32
	_ = v3252
	var v3278 int32
	_ = v3278
	var v3279 int32
	_ = v3279
	var v3281 int32
	_ = v3281
	var v3283 int32
	_ = v3283
	var v3285 int32
	_ = v3285
	var v3290 int32
	_ = v3290
	var v3291 int32
	_ = v3291
	var v3292 int32
	_ = v3292
	var v3293 int32
	_ = v3293
	var v3294 int32
	_ = v3294
	var v3297 int32
	_ = v3297
	var v3298 int32
	_ = v3298
	var v3299 int64
	_ = v3299
	var v3301 int64
	_ = v3301
	var v3303 int64
	_ = v3303
	var v3305 int32
	_ = v3305
	var v3309 int32
	_ = v3309
	var v3310 int32
	_ = v3310
	var v3311 int32
	_ = v3311
	var v3313 int64
	_ = v3313
	var v3317 int32
	_ = v3317
	var v3318 int32
	_ = v3318
	var v3324 int32
	_ = v3324
	var v3331 int32
	_ = v3331
	var v3332 int32
	_ = v3332
	var v3333 int32
	_ = v3333
	var v3334 int32
	_ = v3334
	var v3335 int32
	_ = v3335
	var v3337 int32
	_ = v3337
	var v3338 int32
	_ = v3338
	var v3339 int32
	_ = v3339
	var v3341 int32
	_ = v3341
	var v3342 int32
	_ = v3342
	var v3344 int32
	_ = v3344
	var v3346 int32
	_ = v3346
	var v3350 int32
	_ = v3350
	var v3351 int32
	_ = v3351
	var v3352 int32
	_ = v3352
	var v3353 int32
	_ = v3353
	var v3357 int32
	_ = v3357
	var v3361 int32
	_ = v3361
	var v3365 int32
	_ = v3365
	var v3366 int32
	_ = v3366
	var v3367 int32
	_ = v3367
	var v3368 int32
	_ = v3368
	var v3372 int32
	_ = v3372
	var v3373 int32
	_ = v3373
	var v3374 int32
	_ = v3374
	var v3376 int32
	_ = v3376
	var v3387 int32
	_ = v3387
	var v3391 int32
	_ = v3391
	var v3392 int32
	_ = v3392
	var v3396 int32
	_ = v3396
	var v3397 int32
	_ = v3397
	var v3401 int32
	_ = v3401
	var v3402 int32
	_ = v3402
	var v3404 int32
	_ = v3404
	var v3406 int32
	_ = v3406
	var v3410 int32
	_ = v3410
	var v3411 int32
	_ = v3411
	var v3415 int32
	_ = v3415
	var v3416 int32
	_ = v3416
	var v3418 int32
	_ = v3418
	var v3419 int32
	_ = v3419
	var v3421 int32
	_ = v3421
	var v3425 int32
	_ = v3425
	var v3426 int32
	_ = v3426
	var v3430 int32
	_ = v3430
	var v3431 int32
	_ = v3431
	var v3433 int32
	_ = v3433
	var v3434 int32
	_ = v3434
	var v3435 int32
	_ = v3435
	var v3436 int32
	_ = v3436
	var v3437 int32
	_ = v3437
	var v3439 int32
	_ = v3439
	var v3440 int32
	_ = v3440
	var v3441 int32
	_ = v3441
	var v3443 int32
	_ = v3443
	var v3444 int32
	_ = v3444
	var v3446 int32
	_ = v3446
	var v3448 int32
	_ = v3448
	var v3452 int32
	_ = v3452
	var v3453 int32
	_ = v3453
	var v3454 int32
	_ = v3454
	var v3455 int32
	_ = v3455
	var v3459 int32
	_ = v3459
	var v3463 int32
	_ = v3463
	var v3467 int32
	_ = v3467
	var v3468 int32
	_ = v3468
	var v3469 int32
	_ = v3469
	var v3470 int32
	_ = v3470
	var v3474 int32
	_ = v3474
	var v3475 int32
	_ = v3475
	var v3476 int32
	_ = v3476
	var v3478 int32
	_ = v3478
	var v3489 int32
	_ = v3489
	var v3493 int32
	_ = v3493
	var v3494 int32
	_ = v3494
	var v3498 int32
	_ = v3498
	var v3499 int32
	_ = v3499
	var v3503 int32
	_ = v3503
	var v3504 int32
	_ = v3504
	var v3508 int32
	_ = v3508
	var v3509 int32
	_ = v3509
	var v3510 int32
	_ = v3510
	var v3514 int32
	_ = v3514
	var v3515 int32
	_ = v3515
	var v3516 int32
	_ = v3516
	var v3520 int32
	_ = v3520
	var v3521 int32
	_ = v3521
	var v3522 int32
	_ = v3522
	var v3524 int32
	_ = v3524
	var v3525 int32
	_ = v3525
	var v3526 int32
	_ = v3526
	var v3530 int32
	_ = v3530
	var v3531 int32
	_ = v3531
	var v3532 int32
	_ = v3532
	var v3533 int32
	_ = v3533
	var v3537 int32
	_ = v3537
	var v3538 int32
	_ = v3538
	var v3539 int32
	_ = v3539
	var v3540 int32
	_ = v3540
	var v3544 int32
	_ = v3544
	var v3545 int32
	_ = v3545
	var v3546 int32
	_ = v3546
	var v3547 int32
	_ = v3547
	var v3551 int32
	_ = v3551
	var v3552 int32
	_ = v3552
	var v3553 int32
	_ = v3553
	var v3556 int32
	_ = v3556
	var v3558 int32
	_ = v3558
	var v3562 int32
	_ = v3562
	var v3566 int32
	_ = v3566
	var v3570 int32
	_ = v3570
	var v3574 int32
	_ = v3574
	var v3578 int32
	_ = v3578
	var v3583 int32
	_ = v3583
	var v3584 int32
	_ = v3584
	var v3585 int32
	_ = v3585
	var v3588 int32
	_ = v3588
	var v3589 int32
	_ = v3589
	var v3600 int32
	_ = v3600
	var v3603 int32
	_ = v3603
	var v3625 int64
	_ = v3625
	var v3626 int64
	_ = v3626
	var v3628 int64
	_ = v3628
	var v3631 int64
	_ = v3631
	var v3638 int32
	_ = v3638
	var v3641 int32
	_ = v3641
	var v3642 int32
	_ = v3642
	var v3651 int32
	_ = v3651
	var v3676 int32
	_ = v3676
	var v3679 int32
	_ = v3679
	var v3680 int32
	_ = v3680
	var v3682 int32
	_ = v3682
	var v3683 int32
	_ = v3683
	var v3685 int32
	_ = v3685
	var v3689 int32
	_ = v3689
	var v3690 int32
	_ = v3690
	var v3691 int32
	_ = v3691
	var v3692 int32
	_ = v3692
	var v3693 int32
	_ = v3693
	var v3696 int32
	_ = v3696
	var v3699 int32
	_ = v3699
	var v3700 int32
	_ = v3700
	var v3701 int32
	_ = v3701
	var v3702 int32
	_ = v3702
	var v3705 int32
	_ = v3705
	var v3711 int32
	_ = v3711
	var v3723 int32
	_ = v3723
	var v3745 int32
	_ = v3745
	var v3749 int32
	_ = v3749
	var v3753 int32
	_ = v3753
	var v3754 int32
	_ = v3754
	var v3758 int32
	_ = v3758
	var v3760 int32
	_ = v3760
	var v3761 int32
	_ = v3761
	var v3763 int32
	_ = v3763
	var v3767 int32
	_ = v3767
	var v3768 int32
	_ = v3768
	var v3769 int32
	_ = v3769
	var v3770 int32
	_ = v3770
	var v3771 int32
	_ = v3771
	var v3774 int32
	_ = v3774
	var v3777 int32
	_ = v3777
	var v3778 int32
	_ = v3778
	var v3779 int32
	_ = v3779
	var v3780 int32
	_ = v3780
	var v3783 int32
	_ = v3783
	var v3790 int32
	_ = v3790
	var v3791 int32
	_ = v3791
	var v3826 int32
	_ = v3826
	var v3827 int32
	_ = v3827
	var v3828 int32
	_ = v3828
	var v3862 int32
	_ = v3862
	var v3864 int32
	_ = v3864
	var v3866 int64
	_ = v3866
	var v3873 int32
	_ = v3873
	var v3874 int32
	_ = v3874
	var v3877 int32
	_ = v3877
	var v3878 int32
	_ = v3878
	var v3880 int32
	_ = v3880
	var v3885 int32
	_ = v3885
	var v3888 int32
	_ = v3888
	var v3889 int32
	_ = v3889
	var v3890 int32
	_ = v3890
	var v3891 int32
	_ = v3891
	var v3894 int32
	_ = v3894
	var v3896 int32
	_ = v3896
	var v3897 int32
	_ = v3897
	var v3898 int64
	_ = v3898
	var v3900 int64
	_ = v3900
	var v3902 int64
	_ = v3902
	var v3904 int32
	_ = v3904
	var v3908 int32
	_ = v3908
	var v3912 int32
	_ = v3912
	var v3915 int32
	_ = v3915
	var v3917 int32
	_ = v3917
	var v3922 int32
	_ = v3922
	var v3923 int32
	_ = v3923
	var v3924 int32
	_ = v3924
	var v3925 int32
	_ = v3925
	var v3926 int32
	_ = v3926
	var v3929 int32
	_ = v3929
	var v3931 int32
	_ = v3931
	var v3933 int32
	_ = v3933
	var v3935 int32
	_ = v3935
	var v3937 int32
	_ = v3937
	var v3938 int32
	_ = v3938
	var v3939 int32
	_ = v3939
	var v3940 int32
	_ = v3940
	var v3951 int32
	_ = v3951
	var v3953 int32
	_ = v3953
	var v3958 int32
	_ = v3958
	var v3959 int32
	_ = v3959
	var v3973 int32
	_ = v3973
	var v3980 int32
	_ = v3980
	var v3988 int32
	_ = v3988
	var v3989 int32
	_ = v3989
	var v4031 int32
	_ = v4031
	var v4032 int32
	_ = v4032
	var v4038 int64
	_ = v4038
	var v4039 int64
	_ = v4039
	var v4044 int64
	_ = v4044
	var v4048 int32
	_ = v4048
	var v4055 int32
	_ = v4055
	var v4056 int32
	_ = v4056
	var v4061 int32
	_ = v4061
	var v4065 int32
	_ = v4065
	var v4066 int32
	_ = v4066
	var v4068 int32
	_ = v4068
	var v4070 int32
	_ = v4070
	var v4077 int32
	_ = v4077
	var v4081 int32
	_ = v4081
	var v4085 int32
	_ = v4085
	var v4089 int32
	_ = v4089
	var v4094 int32
	_ = v4094
	var v4097 int32
	_ = v4097
	var v4114 int32
	_ = v4114
	var v4127 int32
	_ = v4127
	var v4128 int64
	_ = v4128
	var v4132 int32
	_ = v4132
	var v4134 int32
	_ = v4134
	var v4135 int32
	_ = v4135
	var v4138 int32
	_ = v4138
	var v4140 int32
	_ = v4140
	var v4142 int32
	_ = v4142
	var v4143 int32
	_ = v4143
	var v4144 int64
	_ = v4144
	var v4146 int32
	_ = v4146
	v3 = int32(0)
	v33 = m.G0
	v35 = v33 - int32(3168)
	m.G0 = v35
	v42 = v35
	v43 = v3
	v44 = int32(-1)
	v49 = v3
	v59 = v35 + int32(128)
	v65 = int64(0)
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
	if v44 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v4127 = int32(m.ExcTag)
	v4128 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v4127 == int32(0) {
		goto L634
	} else {
		goto L635
	}
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v43
	F_AuxiliaryProcessMainCommon(m)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L10
	}
L8:
	;
	v497 = v43
	v498 = v49
	goto L9
L9:
	;
	if v498 != 0 {
		goto L105
	} else {
		goto L106
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v43
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v65
	v82 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L11
	}
L11:
	;
	if v82 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v43
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v65
	F_errmsg_internal(m, int32(_a_F_WalSummarizerMain_0), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v43
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v65
	v103 = m.G0
	v105 = v103 - int32(32)
	m.G0 = v105
	v108 = int32(967)
	switch v108 {
	case 0, 2:
		goto L18
	default:
		goto L19
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v43
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v65
	F_errfinish(m, int32(_a_F_WalSummarizerMain_1), int32(253), int32(_a_F_WalSummarizerMain_2))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v43
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v65
	v145 = int32(0)
	v147 = m.G0
	v149 = v147 - int32(32)
	m.G0 = v149
	switch v145 {
	case 0, 2:
		goto L28
	default:
		goto L29
	}
L18:
	;
	F_sigemptyset(m, v105+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v105)+24)) = int32(268435456)
	switch v108 {
	case 0:
		goto L23
	default:
		goto L21
	case 2:
		goto L22
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[0])) = int32(965)
	goto L18
L20:
	;
	goto L25
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v105)+12)) = int32(_a_F_WalSummarizerMain_3)
	goto L20
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105)+12)) = int32(0)
	goto L20
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105)+12)) = int32(-2)
	goto L20
L25:
	;
	goto L26
L26:
	;
	v137 = F___sigaction(m, int32(1), v105+int32(12), int32(0))
	mBase = m.M
	m.G0 = v105 + int32(32)
	goto L17
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v43
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v65
	v191 = m.G0
	v193 = v191 - int32(32)
	m.G0 = v193
	v196 = int32(969)
	switch v196 {
	case 0, 2:
		goto L38
	default:
		goto L39
	}
L28:
	;
	F_sigemptyset(m, v149+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v149)+24)) = int32(268435456)
	switch v145 {
	case 0:
		goto L33
	default:
		goto L31
	case 2:
		goto L32
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[1])) = int32(-2)
	goto L28
L30:
	;
	goto L35
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v149)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v149)+12)) = int32(_a_F_WalSummarizerMain_3)
	goto L30
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v149)+12)) = int32(0)
	goto L30
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v149)+12)) = int32(-2)
	goto L30
L35:
	;
	goto L36
L36:
	;
	v181 = F___sigaction(m, int32(2), v149+int32(12), int32(0))
	mBase = m.M
	m.G0 = v149 + int32(32)
	goto L27
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v43
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v65
	v233 = int32(0)
	v235 = m.G0
	v237 = v235 - int32(32)
	m.G0 = v237
	switch v233 {
	case 0, 2:
		goto L48
	default:
		goto L49
	}
L38:
	;
	F_sigemptyset(m, v193+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v193)+24)) = int32(268435456)
	switch v196 {
	case 0:
		goto L43
	default:
		goto L41
	case 2:
		goto L42
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[2])) = int32(967)
	goto L38
L40:
	;
	goto L45
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v193)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v193)+12)) = int32(_a_F_WalSummarizerMain_3)
	goto L40
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v193)+12)) = int32(0)
	goto L40
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v193)+12)) = int32(-2)
	goto L40
L45:
	;
	goto L46
L46:
	;
	v225 = F___sigaction(m, int32(15), v193+int32(12), int32(0))
	mBase = m.M
	m.G0 = v193 + int32(32)
	goto L37
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v43
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v65
	v277 = int32(0)
	v279 = m.G0
	v281 = v279 - int32(32)
	m.G0 = v281
	switch v277 {
	case 0, 2:
		goto L58
	default:
		goto L59
	}
L48:
	;
	F_sigemptyset(m, v237+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v237)+24)) = int32(268435456)
	switch v233 {
	case 0:
		goto L53
	default:
		goto L51
	case 2:
		goto L52
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[3])) = int32(-2)
	goto L48
L50:
	;
	goto L55
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v237)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v237)+12)) = int32(_a_F_WalSummarizerMain_3)
	goto L50
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v237)+12)) = int32(0)
	goto L50
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v237)+12)) = int32(-2)
	goto L50
L55:
	;
	goto L56
L56:
	;
	v269 = F___sigaction(m, int32(14), v237+int32(12), int32(0))
	mBase = m.M
	m.G0 = v237 + int32(32)
	goto L47
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v43
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v65
	v323 = m.G0
	v325 = v323 - int32(32)
	m.G0 = v325
	v328 = int32(970)
	switch v328 {
	case 0, 2:
		goto L68
	default:
		goto L69
	}
L58:
	;
	F_sigemptyset(m, v281+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v281)+24)) = int32(268435456)
	switch v277 {
	case 0:
		goto L63
	default:
		goto L61
	case 2:
		goto L62
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[4])) = int32(-2)
	goto L58
L60:
	;
	goto L65
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v281)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v281)+12)) = int32(_a_F_WalSummarizerMain_3)
	goto L60
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v281)+12)) = int32(0)
	goto L60
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v281)+12)) = int32(-2)
	goto L60
L65:
	;
	goto L66
L66:
	;
	v313 = F___sigaction(m, int32(13), v281+int32(12), int32(0))
	mBase = m.M
	m.G0 = v281 + int32(32)
	goto L57
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v43
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v65
	v365 = int32(0)
	v367 = m.G0
	v369 = v367 - int32(32)
	m.G0 = v369
	switch v365 {
	case 0, 2:
		goto L78
	default:
		goto L79
	}
L68:
	;
	F_sigemptyset(m, v325+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v325)+24)) = int32(268435456)
	switch v328 {
	case 0:
		goto L73
	default:
		goto L71
	case 2:
		goto L72
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[5])) = int32(968)
	goto L68
L70:
	;
	goto L75
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v325)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v325)+12)) = int32(_a_F_WalSummarizerMain_3)
	goto L70
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v325)+12)) = int32(0)
	goto L70
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v325)+12)) = int32(-2)
	goto L70
L75:
	;
	goto L76
L76:
	;
	v357 = F___sigaction(m, int32(10), v325+int32(12), int32(0))
	mBase = m.M
	m.G0 = v325 + int32(32)
	goto L67
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v43
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v65
	F_on_shmem_exit(m, int32(1027), int64(0))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L87
	}
L78:
	;
	F_sigemptyset(m, v369+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v369)+24)) = int32(268435456)
	switch v365 {
	case 0:
		goto L83
	default:
		goto L81
	case 2:
		goto L82
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[6])) = int32(-2)
	goto L78
L80:
	;
	goto L85
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v369)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v369)+12)) = int32(_a_F_WalSummarizerMain_3)
	goto L80
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v369)+12)) = int32(0)
	goto L80
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v369)+12)) = int32(-2)
	goto L80
L85:
	;
	goto L86
L86:
	;
	v401 = F___sigaction(m, int32(12), v369+int32(12), int32(0))
	mBase = m.M
	m.G0 = v369 + int32(32)
	goto L77
L87:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v43
	v414 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[7]))
	v418 = F_LWLockAcquire(m, v414+int32(_a_F_WalSummarizerMain_4), int32(0))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L88
	}
L88:
	;
	v421 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[8]))
	v423 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v421)+20)) = v423
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v43
	v428 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[7]))
	F_LWLockRelease(m, v428+int32(_a_F_WalSummarizerMain_4))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L89
	}
L89:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v43
	v437 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[10]))
	v442 = F_AllocSetContextCreateInternal(m, v437, int32(_a_F_WalSummarizerMain_5), int32(0), int32(_a_F_WalSummarizerMain_6), int32(_a_F_WalSummarizerMain_7))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[11])) = v442
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v442
	v451 = m.G0
	v453 = v451 - int32(32)
	m.G0 = v453
	v455 = int32(2)
	switch v455 {
	case 0, 2:
		goto L92
	default:
		goto L93
	}
L91:
	;
	goto L101
L92:
	;
	F_sigemptyset(m, v453+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v453)+24)) = int32(268435456)
	switch v455 {
	case 0:
		goto L97
	default:
		goto L95
	case 2:
		goto L96
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[12])) = int32(0)
	goto L92
L94:
	;
	goto L98
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v453)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v453)+12)) = int32(_a_F_WalSummarizerMain_3)
	v478 = int32(268435461)
	goto L94
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v453)+12)) = int32(0)
	v478 = int32(268435457)
	goto L94
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v453)+12)) = int32(-2)
	v478 = int32(268435457)
	goto L94
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v453)+24)) = v478
	goto L100
L100:
	;
	v485 = F___sigaction(m, int32(17), v453+int32(12), int32(0))
	mBase = m.M
	m.G0 = v453 + int32(32)
	goto L91
L101:
	;
	v490 = v42 + int32(304)
	*(*int32)(unsafe.Add(mBase, uint32(v490)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v490))) = v42 + int32(288)
	goto L104
L102:
	;
	v497 = v442
	v498 = int32(0)
	goto L9
L104:
	;
	goto L102
L105:
	;
	v499 = int32(_a_F_WalSummarizerMain_8)
	v501 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[13]))
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[13])) = v501 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[14])) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	F_EmitErrorReport(m)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v65
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[15])) = v42 + int32(304)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	F_pgmem_sigprocmask(m, int32(_a_F_WalSummarizerMain_9), int32(0))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L118
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v65
	F_LWLockReleaseAll(m)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v65
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L110
	}
L110:
	;
	v521 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[16]))
	*(*int32)(unsafe.Add(mBase, uint32(v521))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v65
	F_pgaio_error_cleanup(m)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L111
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v65
	F_ReleaseAuxProcessResources(m, int32(0))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v65
	F_AtEOXact_Files(m, int32(0))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v65
	F_AtEOXact_HashTables(m, int32(0))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[11])) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	F_FlushErrorState(m)
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v65
	F_MemoryContextReset(m, v497)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L116
	}
L116:
	;
	v553 = int32(_a_F_WalSummarizerMain_8)
	v555 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[13]))
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[13])) = v555 - int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	v565 = F_WaitLatch(m, int32(0), int32(40), int32(_a_F_WalSummarizerMain_10), int32(150994954))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L117
	}
L117:
	;
	goto L107
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v65
	v583 = F_GetOldestUnsummarizedLSN(m, v42+int32(300), v42+int32(299))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L119
	}
L119:
	;
	if v583 != int64(0) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v588 = int32(0)
	v591 = v588
	v604 = v588
	v615 = v65
	v617 = v583
	v619 = int64(0)
	goto L123
L121:
	;
	goto L122
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v65
	F_proc_exit(m, int32(0))
	mBase = m.M
	v4094 = m.ExcPending
	if v4094 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L633
	}
L123:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v615
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	F_MemoryContextReset(m, v497)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v615
	F_ProcessWalSummarizerInterrupts(m)
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v615
	v632 = F_GetRedoRecPtr(m)
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L127
	}
L127:
	;
	v635 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[17]))
	if v635 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v615
	v917 = F_GetLatestLSN(m, v42+int32(292))
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L179
	}
L129:
	;
	v639 = *(*int64)(unsafe.Add(mBase, _c_F_WalSummarizerMain[18]))
	if v632 == v639 {
		goto L128
	} else {
		goto L130
	}
L130:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_WalSummarizerMain[18])) = v632
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v615
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	v645 = F_time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v615
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	v649 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[17]))
	v651 = int64(0)
	v653 = F_GetWalSummaries(m, int32(0), v651, v651)
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L131
	}
L131:
	;
	if v653 == int32(0) {
		goto L128
	} else {
		goto L132
	}
L132:
	;
	v665 = v653
	goto L133
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v615
	F_ProcessWalSummarizerInterrupts(m)
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L135
	}
L134:
	;
	goto L128
L135:
	;
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v665)+12))
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v697)))
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v698)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v615
	v702 = F_XLogGetOldestSegno(m, v699)
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L136
	}
L136:
	;
	v706 = int64(*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[19])))
	v707 = v702 * v706
	v712 = v665
	v717 = int32(0)
	goto L137
L137:
	;
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v712)+4))
	if v717 < v740 {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	if v875 != 0 {
		v665 = v875
		goto L133
	} else {
		goto L178
	}
L139:
	;
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v712)+12))
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v742+v717<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v615
	F_ProcessWalSummarizerInterrupts(m)
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L142
	}
L140:
	;
	v875 = v712
	goto L141
L141:
	;
	goto L138
L142:
	;
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v746)+16))
	if v751 != v699 {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	if v869 != 0 {
		v712 = v869
		v717 = v870
		goto L137
	} else {
		goto L177
	}
L144:
	;
	v869 = v712
	v870 = v717 + int32(1)
	goto L143
L145:
	;
	goto L146
L146:
	;
	if v707 != int64(0) {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v615
	v862 = F_list_delete_nth_cell(m, v712, v717)
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L175
	}
L148:
	;
	v757 = *(*int64)(unsafe.Add(mBase, uint32(v746)+8))
	if base.Ui64(v707) < base.Ui64(v757) {
		goto L147
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v615
	v761 = m.G0
	v763 = v761 - int32(1200)
	m.G0 = v763
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v746)+16))
	v766 = *(*int64)(unsafe.Add(mBase, uint32(v746)))
	v769 = *(*int64)(unsafe.Add(mBase, uint32(v746)+8))
	*(*uint32)(unsafe.Add(mBase, uint32(v763-int32(-64)))) = uint32(v769)
	*(*uint32)(unsafe.Add(mBase, uint32(v763)+56)) = uint32(v766)
	*(*int32)(unsafe.Add(mBase, uint32(v763)+48)) = v765
	v773 = int64(32)
	v774 = int64(base.Ui64(v769) >> (uint(v773) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v763)+60)) = uint32(v774)
	v777 = int64(base.Ui64(v766) >> (uint(v773) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v763)+52)) = uint32(v777)
	v780 = v763 + int32(176)
	v785 = F_pg_snprintf(m, v780, int32(1024), int32(_a_F_WalSummarizerMain_11), v763+int32(48))
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L152
	}
L151:
	;
	goto L150
L152:
	;
	v791 = F___fstatat(m, int32(-100), v780, v763+int32(80), int32(256))
	mBase = m.M
	goto L156
L153:
	;
	goto L147
L154:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L171
	}
L155:
	;
	m.G0 = v763 + int32(1200)
	goto L153
L156:
	;
	if v791 != 0 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v793 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[20]))
	if v793 == int32(44) {
		goto L155
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	v813 = *(*int64)(unsafe.Add(mBase, uint32(v763)+136))
	if v645-base.I64_extend_i32_s(v649*int32(60)) <= v813 {
		goto L155
	} else {
		goto L165
	}
L160:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L161
	}
L161:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L162
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v763)+32)) = v780
	F_errmsg(m, int32(_a_F_WalSummarizerMain_12), v763+int32(32))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L163
	}
L163:
	;
	F_errfinish(m, int32(_a_F_WalSummarizerMain_13), int32(247), int32(_a_F_WalSummarizerMain_14))
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L164
	}
L164:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L165:
	;
	v816 = v763 + int32(176)
	v817 = F_unlink(m, v816)
	mBase = m.M
	if v817 != 0 {
		goto L154
	} else {
		goto L166
	}
L166:
	;
	v820 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L167
	}
L167:
	;
	if v820 == int32(0) {
		goto L155
	} else {
		goto L168
	}
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v763))) = v816
	F_errmsg_internal(m, int32(_a_F_WalSummarizerMain_15), v763)
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L169
	}
L169:
	;
	F_errfinish(m, int32(_a_F_WalSummarizerMain_13), int32(256), int32(_a_F_WalSummarizerMain_14))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L170
	}
L170:
	;
	goto L155
L171:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L172
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v763)+16)) = v763 + int32(176)
	F_errmsg(m, int32(_a_F_WalSummarizerMain_16), v763+int32(16))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L173
	}
L173:
	;
	F_errfinish(m, int32(_a_F_WalSummarizerMain_13), int32(254), int32(_a_F_WalSummarizerMain_14))
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L174
	}
L174:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v615
	F_pfree(m, v746)
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L176
	}
L176:
	;
	v869 = v862
	v870 = v717
	goto L143
L177:
	;
	v875 = v869
	goto L141
L178:
	;
	goto L134
L179:
	;
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v42)+300))
	if v619 != int64(0) {
		goto L183
	} else {
		goto L184
	}
L180:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v615
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	v1354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+299)))
	v1355 = F_CreateEmptyBlockRefTable(m)
	mBase = m.M
	v1356 = m.ExcPending
	if v1356 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L231
	}
L181:
	;
	v1320 = v591
	v1332 = v919
	v1333 = v604
	v1345 = v917
	v1351 = int64(0)
	goto L180
L182:
	;
	if base.Ui64(v617) < base.Ui64(v1276) {
		goto L225
	} else {
		goto L226
	}
L183:
	;
	if v619 == int64(0) {
		goto L181
	} else {
		goto L224
	}
L184:
	;
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v42)+292))
	if v919 == v922 {
		goto L183
	} else {
		goto L185
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v615
	v926 = F_readTimeLineHistory(m, v922)
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L186
	}
L186:
	;
	v930 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[11])) = v930
	if v926 == int32(0) {
		goto L188
	} else {
		goto L189
	}
L187:
	;
	if v953 == int32(0) {
		goto L201
	} else {
		goto L202
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v615
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L198
	}
L189:
	;
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v926)+4))
	if v934 <= int32(0) {
		goto L188
	} else {
		goto L190
	}
L190:
	;
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v926)+12))
	v953 = int32(0)
	goto L191
L191:
	;
	v972 = v953 << (uint(int32(2)) % 32)
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v937+v972)))
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v974)))
	if v919 != v975 {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	v980 = *(*int64)(unsafe.Add(mBase, uint32(v974)+16))
	if v980 != int64(0) {
		goto L187
	} else {
		goto L197
	}
L193:
	;
	v978 = v953 + int32(1)
	if v934 != v978 {
		v953 = v978
		goto L191
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	goto L192
L196:
	;
	goto L188
L197:
	;
	goto L188
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v615
	*(*int32)(unsafe.Add(mBase, uint32(v42)+240)) = v919
	F_errmsg(m, int32(_a_F_WalSummarizerMain_17), v42+int32(240))
	mBase = m.M
	v1028 = m.ExcPending
	if v1028 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L199
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v615
	F_errfinish(m, int32(_a_F_WalSummarizerMain_1), int32(927), int32(_a_F_WalSummarizerMain_18))
	mBase = m.M
	v1035 = m.ExcPending
	if v1035 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L200
	}
L200:
	;
	goto L3
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v615
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L204
	}
L202:
	;
	goto L203
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v615
	v1062 = F_palloc_mul(m, int32(4), v953)
	mBase = m.M
	v1063 = m.ExcPending
	if v1063 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L207
	}
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v615
	*(*int32)(unsafe.Add(mBase, uint32(v42)+256)) = v919
	F_errmsg_internal(m, int32(_a_F_WalSummarizerMain_19), v42+int32(256))
	mBase = m.M
	v1051 = m.ExcPending
	if v1051 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L205
	}
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v615
	F_errfinish(m, int32(_a_F_WalSummarizerMain_1), int32(929), int32(_a_F_WalSummarizerMain_18))
	mBase = m.M
	v1058 = m.ExcPending
	if v1058 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L206
	}
L206:
	;
	goto L3
L207:
	;
	v1064 = int32(0)
	if v953 != int32(1) {
		goto L209
	} else {
		goto L210
	}
L208:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[11])) = v497
	if v591 != 0 {
		goto L216
	} else {
		goto L217
	}
L209:
	;
	v1076 = v1064
	v1082 = int32(0)
	goto L212
L210:
	;
	v1138 = v1064
	goto L211
L211:
	;
	v1166 = int32(2)
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(v926)+12))
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(v1169+v972+(v1138^int32(-1))<<(uint(v1166)%32))))
	v1177 = *(*int32)(unsafe.Add(mBase, uint32(v1176)))
	*(*int32)(unsafe.Add(mBase, uint32(v1062+v1138<<(uint(v1166)%32)))) = v1177
	goto L208
L212:
	;
	v1104 = int32(2)
	v1106 = v1062 + v1076<<(uint(v1104)%32)
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(v926)+12))
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(v1107+v972+(v1076^int32(-1))<<(uint(v1104)%32))))
	v1115 = *(*int32)(unsafe.Add(mBase, uint32(v1114)))
	*(*int32)(unsafe.Add(mBase, uint32(v1106))) = v1115
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v926)+12))
	v1124 = *(*int32)(unsafe.Add(mBase, uint32(v1117+v972+(v1076^int32(1073741822))<<(uint(v1104)%32))))
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(v1124)))
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+4)) = v1125
	v1128 = v1076 + v1104
	v1130 = v1082 + v1104
	if v1130 != v953&int32(2147483646) {
		v1076 = v1128
		v1082 = v1130
		goto L212
	} else {
		goto L214
	}
L213:
	;
	if v953&int32(1) == int32(0) {
		goto L208
	} else {
		goto L215
	}
L214:
	;
	goto L213
L215:
	;
	v1138 = v1128
	goto L211
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v615
	F_pfree(m, v591)
	mBase = m.M
	v1216 = m.ExcPending
	if v1216 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L219
	}
L217:
	;
	goto L218
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v615
	v1221 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1222 = m.ExcPending
	if v1222 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L220
	}
L219:
	;
	goto L218
L220:
	;
	if v1221 == int32(0) {
		v1250 = v1062
		v1264 = v953
		v1276 = v980
		goto L182
	} else {
		goto L221
	}
L221:
	;
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(v1062)))
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v615
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	v1228 = *(*int32)(unsafe.Add(mBase, uint32(v42)+300))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+276)) = v1225
	*(*int32)(unsafe.Add(mBase, uint32(v42)+272)) = v1228
	*(*uint32)(unsafe.Add(mBase, uint32(v42)+284)) = uint32(v980)
	v1233 = int64(base.Ui64(v980) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v42)+280)) = uint32(v1233)
	F_errmsg_internal(m, int32(_a_F_WalSummarizerMain_20), v42+int32(272))
	mBase = m.M
	v1239 = m.ExcPending
	if v1239 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L222
	}
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v615
	F_errfinish(m, int32(_a_F_WalSummarizerMain_1), int32(421), int32(_a_F_WalSummarizerMain_2))
	mBase = m.M
	v1246 = m.ExcPending
	if v1246 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L223
	}
L223:
	;
	v1250 = v1062
	v1264 = v953
	v1276 = v980
	goto L182
L224:
	;
	v1250 = v591
	v1264 = v604
	v1276 = v619
	goto L182
L225:
	;
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(v42)+300))
	v1320 = v1250
	v1332 = v1283
	v1333 = v1264
	v1345 = v1276
	v1351 = v1276
	goto L180
L226:
	;
	goto L227
L227:
	;
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(v1250)))
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v615
	*(*int32)(unsafe.Add(mBase, uint32(v42)+300)) = v1284
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	F_pfree(m, v1250)
	mBase = m.M
	v1289 = m.ExcPending
	if v1289 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L228
	}
L228:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v615
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	v1293 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[7]))
	v1297 = F_LWLockAcquire(m, v1293+int32(_a_F_WalSummarizerMain_4), int32(0))
	mBase = m.M
	v1298 = m.ExcPending
	if v1298 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L229
	}
L229:
	;
	v1301 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[8]))
	*(*int64)(unsafe.Add(mBase, uint32(v1301)+8)) = v1276
	v1303 = *(*int32)(unsafe.Add(mBase, uint32(v42)+300))
	*(*int64)(unsafe.Add(mBase, uint32(v1301)+24)) = v1276
	v1305 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1301)+16)) = uint8(v1305)
	*(*int32)(unsafe.Add(mBase, uint32(v1301)+4)) = v1303
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v615
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	v1311 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[7]))
	F_LWLockRelease(m, v1311+int32(_a_F_WalSummarizerMain_4))
	mBase = m.M
	v1315 = m.ExcPending
	if v1315 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L230
	}
L230:
	;
	v591 = int32(0)
	v604 = int32(0)
	v617 = v1276
	v619 = int64(0)
	goto L123
L231:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v615
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	v1360 = F_palloc0(m, int32(32))
	mBase = m.M
	v1361 = m.ExcPending
	if v1361 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L232
	}
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1360)+24)) = v1320
	*(*int32)(unsafe.Add(mBase, uint32(v1360)+20)) = v1333
	*(*int64)(unsafe.Add(mBase, uint32(v1360)+8)) = v1345
	*(*int32)(unsafe.Add(mBase, uint32(v1360))) = v1332
	*(*uint8)(unsafe.Add(mBase, uint32(v1360)+4)) = uint8(base.B2i32(v1351 != int64(0)))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+488)) = int32(414)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+484)) = int32(1028)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+480)) = int32(1029)
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v615
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	v1378 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[19]))
	v1381 = F_XLogReaderAllocate(m, v1378, v42+int32(480), v1360)
	mBase = m.M
	v1382 = m.ExcPending
	if v1382 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L233
	}
L233:
	;
	if v1381 == int32(0) {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v615
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1390 = m.ExcPending
	if v1390 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L237
	}
L235:
	;
	goto L236
L236:
	;
	if v1354&int32(1) != 0 {
		goto L251
	} else {
		goto L252
	}
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v615
	F_errcode(m, int32(_a_F_WalSummarizerMain_21))
	mBase = m.M
	v1395 = m.ExcPending
	if v1395 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L238
	}
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v615
	F_errmsg(m, int32(_a_F_WalSummarizerMain_22), int32(0))
	mBase = m.M
	v1401 = m.ExcPending
	if v1401 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L239
	}
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v615
	v1406 = F_errdetail(m, int32(_a_F_WalSummarizerMain_23), int32(0))
	mBase = m.M
	v1407 = m.ExcPending
	if v1407 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L240
	}
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v615
	F_errfinish(m, int32(_a_F_WalSummarizerMain_1), int32(1035), int32(_a_F_WalSummarizerMain_24))
	mBase = m.M
	v1414 = m.ExcPending
	if v1414 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L241
	}
L241:
	;
	goto L3
L242:
	;
	if (v2853|(v2812^int32(-1)))&int32(1) != 0 {
		goto L624
	} else {
		goto L625
	}
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	v2978 = int32(0)
	v2979 = m.G0
	v2981 = v2979 - int32(_a_F_WalSummarizerMain_25)
	m.G0 = v2981
	*(*int32)(unsafe.Add(mBase, uint32(v2981)+8)) = int32(1697321851)
	base.MemoryFill(m, v2981+int32(24), v2978, int32(_a_F_WalSummarizerMain_26))
	*(*int32)(unsafe.Add(mBase, uint32(v2981)+16)) = v42 + int32(496)
	*(*int32)(unsafe.Add(mBase, uint32(v2981)+12)) = int32(1030)
	v2995 = int32(-1)
	v2999 = int32(4)
	v3000 = m.Env.Pgmem_crc32c(m, v2995, v2981+int32(8), v2999)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v2981)+uint32(_c_F_WalSummarizerMain[21]))) = v3000
	v3002 = *(*int32)(unsafe.Add(mBase, uint32(v1355)))
	*(*int32)(unsafe.Add(mBase, uint32(v2981)+uint32(_c_F_WalSummarizerMain[22]))) = v2999
	v3005 = *(*int32)(unsafe.Add(mBase, uint32(v2981)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2981)+20)) = v3005
	v3007 = *(*int32)(unsafe.Add(mBase, uint32(v3002)+8))
	if v3007 == v2978 {
		goto L495
	} else {
		goto L496
	}
L244:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	v2926 = *(*int32)(unsafe.Add(mBase, uint32(v42)+492))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2930 = m.ExcPending
	if v2930 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L486
	}
L245:
	;
	v2844 = *(*int32)(unsafe.Add(mBase, uint32(v1381)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	F_pfree(m, v2844)
	mBase = m.M
	v2848 = m.ExcPending
	if v2848 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L475
	}
L246:
	;
	v2812 = v2784
	v2838 = v2806
	goto L245
L247:
	;
	v2743 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1360)+16)))
	if v2743 != int32(1) {
		goto L244
	} else {
		goto L468
	}
L248:
	;
	v1937 = int32(1)
	goto L326
L249:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1809
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	v1893 = *(*int32)(unsafe.Add(mBase, uint32(v42)+492))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1897 = m.ExcPending
	if v1897 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L318
	}
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	F_ProcessWalSummarizerInterrupts(m)
	mBase = m.M
	v1883 = m.ExcPending
	if v1883 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L315
	}
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v615
	F_XLogBeginRead(m, v1381, v617)
	mBase = m.M
	v1420 = m.ExcPending
	if v1420 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L254
	}
L252:
	;
	goto L253
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v615
	v1424 = v42 + int32(492)
	v1425 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1424))) = v1425
	*(*uint8)(unsafe.Add(mBase, uint32(v1381)+1257)) = uint8(v1425)
	v1430 = v617 & int64(-8192)
	v1434 = F_ReadPageInternal(m, v1381, v1430, base.I32_wrap_i64(v617)&int32(_a_F_WalSummarizerMain_27))
	mBase = m.M
	v1435 = m.ExcPending
	if v1435 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L257
	}
L254:
	;
	v1872 = v615
	v1874 = v617
	v1879 = v1351
	goto L250
L255:
	;
	if v1809 != int64(0) {
		goto L305
	} else {
		goto L306
	}
L256:
	;
	v1761 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1381)+1192)) = v1761
	*(*int64)(unsafe.Add(mBase, uint32(v1381)+1176)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1381)+132)) = v1761
	v1767 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1381)+1256)))
	if v1767 == int32(1) {
		goto L299
	} else {
		goto L300
	}
L257:
	;
	if v1434 < int32(0) {
		goto L256
	} else {
		goto L258
	}
L258:
	;
	v1464 = v1430
	goto L259
L259:
	;
	v1472 = *(*int32)(unsafe.Add(mBase, uint32(v1381)+128))
	v1473 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1472)+2)))
	if v1473&int32(2) != 0 {
		goto L261
	} else {
		goto L262
	}
L260:
	;
	goto L256
L261:
	;
	v1476 = int32(40)
	goto L263
L262:
	;
	v1476 = int32(24)
	goto L263
L263:
	;
	v1477 = F_ReadPageInternal(m, v1381, v1464, v1476)
	mBase = m.M
	v1478 = m.ExcPending
	if v1478 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L264
	}
L264:
	;
	if v1477 < int32(0) {
		goto L256
	} else {
		goto L265
	}
L265:
	;
	v1481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1472)+2)))
	if v1481&int32(1) != 0 {
		goto L268
	} else {
		goto L269
	}
L266:
	;
	v1723 = v1464 - int64(-8192)
	v1725 = F_ReadPageInternal(m, v1381, v1723, int32(0))
	mBase = m.M
	v1726 = m.ExcPending
	if v1726 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L297
	}
L267:
	;
	v1500 = *(*int32)(unsafe.Add(mBase, uint32(v1381)+120))
	if v1500 != 0 {
		goto L272
	} else {
		goto L273
	}
L268:
	;
	v1484 = *(*int32)(unsafe.Add(mBase, uint32(v1472)+16))
	v1488 = (v1484 + int32(7)) & int32(-8)
	if base.Ui32(int32(_a_F_WalSummarizerMain_6)-v1476) <= base.Ui32(v1488) {
		goto L266
	} else {
		goto L271
	}
L269:
	;
	goto L270
L270:
	;
	v1499 = v1464 | base.I64_extend_i32_u(v1476)
	goto L267
L271:
	;
	v1499 = base.I64_extend_i32_u(v1488) + (v1464 | base.I64_extend_i32_u(v1476))
	goto L267
L272:
	;
	v1505 = v1500
	goto L275
L273:
	;
	goto L274
L274:
	;
	v1574 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1381)+120)) = v1574
	v1576 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1381)+96)) = v1576
	v1578 = *(*int32)(unsafe.Add(mBase, uint32(v1381)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v1381)+116)) = v1578
	*(*int32)(unsafe.Add(mBase, uint32(v1381)+112)) = v1578
	v1581 = *(*int32)(unsafe.Add(mBase, uint32(v1381)+1252))
	*(*uint8)(unsafe.Add(mBase, uint32(v1581))) = uint8(v1576)
	*(*int64)(unsafe.Add(mBase, uint32(v1381)+80)) = v1499
	*(*int64)(unsafe.Add(mBase, uint32(v1381)+40)) = v1499
	*(*uint8)(unsafe.Add(mBase, uint32(v1381)+1256)) = uint8(v1576)
	*(*int64)(unsafe.Add(mBase, uint32(v1381)+72)) = v1574
	*(*int64)(unsafe.Add(mBase, uint32(v1381)+32)) = v1574
	goto L282
L275:
	;
	v1533 = *(*int32)(unsafe.Add(mBase, uint32(v1505)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1381)+120)) = v1533
	v1535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1505)+4)))
	if v1535 == int32(1) {
		goto L277
	} else {
		goto L278
	}
L276:
	;
	goto L274
L277:
	;
	F_pfree(m, v1505)
	mBase = m.M
	v1539 = m.ExcPending
	if v1539 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L280
	}
L278:
	;
	v1541 = v1533
	goto L279
L279:
	;
	if v1541 != 0 {
		v1505 = v1541
		goto L275
	} else {
		goto L281
	}
L280:
	;
	v1540 = *(*int32)(unsafe.Add(mBase, uint32(v1381)+120))
	v1541 = v1540
	goto L279
L281:
	;
	goto L276
L282:
	;
	v1624 = F_XLogReadRecord(m, v1381, v1424)
	mBase = m.M
	v1625 = m.ExcPending
	if v1625 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L284
	}
L283:
	;
	v1630 = *(*int32)(unsafe.Add(mBase, uint32(v1381)+120))
	if v1630 != 0 {
		goto L287
	} else {
		goto L288
	}
L284:
	;
	if v1624 == int32(0) {
		goto L256
	} else {
		goto L285
	}
L285:
	;
	v1628 = *(*int64)(unsafe.Add(mBase, uint32(v1381)+32))
	if base.Ui64(v1628) < base.Ui64(v617) {
		goto L282
	} else {
		goto L286
	}
L286:
	;
	goto L283
L287:
	;
	v1635 = v1630
	goto L290
L288:
	;
	goto L289
L289:
	;
	v1704 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1381)+120)) = v1704
	v1706 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1381)+96)) = v1706
	v1708 = *(*int32)(unsafe.Add(mBase, uint32(v1381)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v1381)+116)) = v1708
	*(*int32)(unsafe.Add(mBase, uint32(v1381)+112)) = v1708
	v1711 = *(*int32)(unsafe.Add(mBase, uint32(v1381)+1252))
	*(*uint8)(unsafe.Add(mBase, uint32(v1711))) = uint8(v1706)
	*(*int64)(unsafe.Add(mBase, uint32(v1381)+80)) = v1628
	*(*int64)(unsafe.Add(mBase, uint32(v1381)+40)) = v1628
	*(*uint8)(unsafe.Add(mBase, uint32(v1381)+1256)) = uint8(v1706)
	*(*int64)(unsafe.Add(mBase, uint32(v1381)+72)) = v1704
	*(*int64)(unsafe.Add(mBase, uint32(v1381)+32)) = v1704
	v1809 = v1628
	goto L255
L290:
	;
	v1663 = *(*int32)(unsafe.Add(mBase, uint32(v1635)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1381)+120)) = v1663
	v1665 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1635)+4)))
	if v1665 == int32(1) {
		goto L292
	} else {
		goto L293
	}
L291:
	;
	goto L289
L292:
	;
	F_pfree(m, v1635)
	mBase = m.M
	v1669 = m.ExcPending
	if v1669 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L295
	}
L293:
	;
	v1671 = v1663
	goto L294
L294:
	;
	if v1671 != 0 {
		v1635 = v1671
		goto L290
	} else {
		goto L296
	}
L295:
	;
	v1670 = *(*int32)(unsafe.Add(mBase, uint32(v1381)+120))
	v1671 = v1670
	goto L294
L296:
	;
	goto L291
L297:
	;
	if int32(0) <= v1725 {
		v1464 = v1723
		goto L259
	} else {
		goto L298
	}
L298:
	;
	goto L260
L299:
	;
	v1770 = *(*int32)(unsafe.Add(mBase, uint32(v1381)+1252))
	v1771 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1770))))
	if v1771 != 0 {
		goto L302
	} else {
		goto L303
	}
L300:
	;
	goto L301
L301:
	;
	v1809 = int64(0)
	goto L255
L302:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1424))) = v1770
	goto L304
L303:
	;
	goto L304
L304:
	;
	v1773 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1381)+1256)) = uint8(v1773)
	goto L301
L305:
	;
	v1872 = v1809
	v1874 = v1809
	v1879 = v1351
	goto L250
L306:
	;
	goto L307
L307:
	;
	v1812 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1360)+16)))
	if v1812 != int32(1) {
		goto L249
	} else {
		goto L308
	}
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1809
	v1819 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1820 = m.ExcPending
	if v1820 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L309
	}
L309:
	;
	if v1819 != 0 {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	v1821 = *(*int64)(unsafe.Add(mBase, uint32(v1360)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1809
	*(*uint32)(unsafe.Add(mBase, uint32(v42)+192)) = uint32(v1821)
	v1825 = int64(32)
	v1826 = int64(base.Ui64(v1821) >> (uint(v1825) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v42)+188)) = uint32(v1826)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+176)) = v1332
	*(*uint32)(unsafe.Add(mBase, uint32(v42)+184)) = uint32(v617)
	v1831 = int64(base.Ui64(v617) >> (uint(v1825) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v42)+180)) = uint32(v1831)
	F_errmsg_internal(m, int32(_a_F_WalSummarizerMain_28), v42+int32(176))
	mBase = m.M
	v1837 = m.ExcPending
	if v1837 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L313
	}
L311:
	;
	goto L312
L312:
	;
	v1846 = *(*int64)(unsafe.Add(mBase, uint32(v1381)+40))
	v1872 = v1809
	v1874 = v617
	v1879 = v1846
	goto L250
L313:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1809
	F_errfinish(m, int32(_a_F_WalSummarizerMain_1), int32(1083), int32(_a_F_WalSummarizerMain_24))
	mBase = m.M
	v1844 = m.ExcPending
	if v1844 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L314
	}
L314:
	;
	goto L312
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	v1888 = F_XLogReadRecord(m, v1381, v42+int32(492))
	mBase = m.M
	v1889 = m.ExcPending
	if v1889 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L316
	}
L316:
	;
	if v1888 != 0 {
		goto L248
	} else {
		goto L317
	}
L317:
	;
	v2715 = int32(1)
	goto L247
L318:
	;
	v1900 = base.I32_wrap_i64(int64(base.Ui64(v617) >> (uint(int64(32)) % 64)))
	v1901 = base.I32_wrap_i64(v617)
	if v1893 != 0 {
		goto L319
	} else {
		goto L320
	}
L319:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1809
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	v1904 = *(*int32)(unsafe.Add(mBase, uint32(v42)+492))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+232)) = v1904
	*(*int32)(unsafe.Add(mBase, uint32(v42)+228)) = v1901
	*(*int32)(unsafe.Add(mBase, uint32(v42)+224)) = v1900
	F_errmsg(m, int32(_a_F_WalSummarizerMain_29), v42+int32(224))
	mBase = m.M
	v1912 = m.ExcPending
	if v1912 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L322
	}
L320:
	;
	goto L321
L321:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1809
	*(*int32)(unsafe.Add(mBase, uint32(v42)+212)) = v1901
	*(*int32)(unsafe.Add(mBase, uint32(v42)+208)) = v1900
	F_errmsg(m, int32(_a_F_WalSummarizerMain_30), v42+int32(208))
	mBase = m.M
	v1928 = m.ExcPending
	if v1928 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L324
	}
L322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1809
	F_errfinish(m, int32(_a_F_WalSummarizerMain_1), int32(1102), int32(_a_F_WalSummarizerMain_24))
	mBase = m.M
	v1919 = m.ExcPending
	if v1919 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L323
	}
L323:
	;
	goto L3
L324:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1809
	F_errfinish(m, int32(_a_F_WalSummarizerMain_1), int32(1106), int32(_a_F_WalSummarizerMain_24))
	mBase = m.M
	v1935 = m.ExcPending
	if v1935 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L325
	}
L325:
	;
	goto L3
L326:
	;
	v1970 = base.B2i32(v1879 == int64(0))
	if v1879 == int64(0) {
		goto L328
	} else {
		goto L329
	}
L327:
	;
	v2715 = v2648
	goto L247
L328:
	;
	v1973 = *(*int32)(unsafe.Add(mBase, uint32(v1381)+96))
	v1974 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1973)+49)))
	if v1974 == int32(0) {
		goto L333
	} else {
		goto L334
	}
L329:
	;
	v1971 = *(*int64)(unsafe.Add(mBase, uint32(v1381)+32))
	if base.Ui64(v1971) < base.Ui64(v1879) {
		goto L328
	} else {
		goto L330
	}
L330:
	;
	v2812 = v1937
	v2838 = v1879
	goto L245
L331:
	;
	v2676 = *(*int64)(unsafe.Add(mBase, uint32(v1381)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	v2680 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[7]))
	v2684 = F_LWLockAcquire(m, v2680+int32(_a_F_WalSummarizerMain_4), int32(0))
	mBase = m.M
	v2685 = m.ExcPending
	if v2685 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L459
	}
L332:
	;
	v2552 = int32(0)
	v2553 = *(*int32)(unsafe.Add(mBase, uint32(v1381)+96))
	v2554 = *(*int32)(unsafe.Add(mBase, uint32(v2553)+72))
	if v2554 < v2552 {
		v2648 = v2552
		goto L331
	} else {
		goto L437
	}
L333:
	;
	v1977 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1973)+48)))
	v1979 = v1977 & int32(240)
	if base.Ui32(v1979) <= base.Ui32(int32(143)) {
		goto L340
	} else {
		goto L341
	}
L334:
	;
	goto L335
L335:
	;
	v2010 = int32(1)
	if v1937&v2010 != 0 {
		v2648 = v2010
		goto L331
	} else {
		goto L350
	}
L336:
	;
	v2005 = int32(1)
	if v1937&v2005 == int32(0) {
		goto L332
	} else {
		goto L349
	}
L337:
	;
	v2001 = *(*int64)(unsafe.Add(mBase, uint32(v1381)+32))
	if base.Ui64(v1874) < base.Ui64(v2001) {
		v2812 = v1937
		v2838 = v2001
		goto L245
	} else {
		goto L347
	}
L338:
	;
	v1997 = *(*int32)(unsafe.Add(mBase, uint32(v1973)+64))
	v2000 = v1997 + int32(16)
	goto L337
L339:
	;
	v1994 = *(*int32)(unsafe.Add(mBase, uint32(v1973)+64))
	v2000 = v1994 + int32(20)
	goto L337
L340:
	;
	if v1979 == int32(0) {
		goto L339
	} else {
		goto L343
	}
L341:
	;
	goto L342
L342:
	;
	if v1979 == int32(144) {
		goto L338
	} else {
		goto L345
	}
L343:
	;
	if v1979 != int32(96) {
		goto L336
	} else {
		goto L344
	}
L344:
	;
	v1986 = *(*int32)(unsafe.Add(mBase, uint32(v1973)+64))
	v2000 = v1986 + int32(20)
	goto L337
L345:
	;
	if v1979 != int32(224) {
		goto L336
	} else {
		goto L346
	}
L346:
	;
	v1993 = *(*int32)(unsafe.Add(mBase, uint32(v1973)+64))
	v2000 = v1993
	goto L337
L347:
	;
	v2003 = *(*int32)(unsafe.Add(mBase, uint32(v2000)))
	if v2003 != 0 {
		goto L332
	} else {
		goto L348
	}
L348:
	;
	v2648 = int32(1)
	goto L331
L349:
	;
	v2648 = v2005
	goto L331
L350:
	;
	switch v1974 - int32(1) {
	case 0:
		goto L351
	case 1:
		goto L352
	default:
		goto L332
	case 3:
		goto L353
	}
L351:
	;
	v2159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1973)+48)))
	switch int32(base.Ui32(v2159)>>(uint(int32(4))%32)) & int32(7) {
	case 0, 3:
		goto L381
	default:
		goto L332
	case 2, 4:
		goto L380
	}
L352:
	;
	v2110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1973)+48)))
	v2114 = v2110&int32(240) - int32(16)
	if v2114 != 0 {
		goto L366
	} else {
		goto L367
	}
L353:
	;
	v2015 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1973)+48)))
	v2017 = v2015 & int32(240)
	switch v2017 - int32(16) {
	case 0:
		goto L355
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
		goto L332
	case 16:
		goto L354
	default:
		goto L356
	}
L354:
	;
	v2050 = *(*int32)(unsafe.Add(mBase, uint32(v1973)+64))
	v2051 = *(*int32)(unsafe.Add(mBase, uint32(v2050)))
	v2052 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+2572)) = v2052
	*(*int32)(unsafe.Add(mBase, uint32(v42)+2568)) = v2051
	v2055 = *(*int32)(unsafe.Add(mBase, uint32(v2050)+4))
	if v2055 <= v2052 {
		goto L332
	} else {
		goto L360
	}
L355:
	;
	v2035 = *(*int32)(unsafe.Add(mBase, uint32(v1973)+64))
	v2036 = *(*int32)(unsafe.Add(mBase, uint32(v2035)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+2576)) = v2036
	v2038 = *(*int32)(unsafe.Add(mBase, uint32(v2035)))
	v2039 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+2584)) = v2039
	*(*int32)(unsafe.Add(mBase, uint32(v42)+2580)) = v2038
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	F_BlockRefTableSetLimitBlock(m, v1355, v42+int32(2576), v2039, v2039)
	mBase = m.M
	v2049 = m.ExcPending
	if v2049 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L359
	}
L356:
	;
	if v2017 != 0 {
		goto L332
	} else {
		goto L357
	}
L357:
	;
	v2020 = *(*int32)(unsafe.Add(mBase, uint32(v1973)+64))
	v2021 = *(*int32)(unsafe.Add(mBase, uint32(v2020)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+2588)) = v2021
	v2023 = *(*int32)(unsafe.Add(mBase, uint32(v2020)))
	v2024 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+2596)) = v2024
	*(*int32)(unsafe.Add(mBase, uint32(v42)+2592)) = v2023
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	F_BlockRefTableSetLimitBlock(m, v1355, v42+int32(2588), v2024, v2024)
	mBase = m.M
	v2034 = m.ExcPending
	if v2034 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L358
	}
L358:
	;
	goto L332
L359:
	;
	goto L332
L360:
	;
	v2065 = int32(0)
	goto L361
L361:
	;
	v2096 = *(*int32)(unsafe.Add(mBase, uint32(v2050+int32(8)+v2065<<(uint(int32(2))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	*(*int32)(unsafe.Add(mBase, uint32(v42)+2564)) = v2096
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	v2102 = int32(0)
	F_BlockRefTableSetLimitBlock(m, v1355, v42+int32(2564), v2102, v2102)
	mBase = m.M
	v2105 = m.ExcPending
	if v2105 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L363
	}
L362:
	;
	goto L332
L363:
	;
	v2107 = v2065 + int32(1)
	v2108 = *(*int32)(unsafe.Add(mBase, uint32(v2050)+4))
	if v2107 < v2108 {
		v2065 = v2107
		goto L361
	} else {
		goto L364
	}
L364:
	;
	goto L362
L365:
	;
	v2126 = *(*int32)(unsafe.Add(mBase, uint32(v1973)+64))
	v2127 = *(*int32)(unsafe.Add(mBase, uint32(v2126)+16))
	if v2127&int32(1) != 0 {
		goto L374
	} else {
		goto L375
	}
L366:
	;
	if v2114 == int32(16) {
		goto L369
	} else {
		goto L370
	}
L367:
	;
	goto L368
L368:
	;
	v2117 = *(*int32)(unsafe.Add(mBase, uint32(v1973)+64))
	v2118 = *(*int32)(unsafe.Add(mBase, uint32(v2117)+12))
	if v2118 == int32(1) {
		goto L332
	} else {
		goto L372
	}
L369:
	;
	goto L365
L370:
	;
	goto L332
L372:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	F_BlockRefTableSetLimitBlock(m, v1355, v2117, v2118, int32(0))
	mBase = m.M
	v2125 = m.ExcPending
	if v2125 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L373
	}
L373:
	;
	goto L332
L374:
	;
	v2130 = *(*int32)(unsafe.Add(mBase, uint32(v2126)))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	F_BlockRefTableSetLimitBlock(m, v1355, v2126+int32(4), int32(0), v2130)
	mBase = m.M
	v2137 = m.ExcPending
	if v2137 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L377
	}
L375:
	;
	v2140 = v2127
	goto L376
L376:
	;
	if v2140&int32(2) == int32(0) {
		goto L332
	} else {
		goto L378
	}
L377:
	;
	v2138 = *(*int32)(unsafe.Add(mBase, uint32(v2126)+16))
	v2140 = v2138
	goto L376
L378:
	;
	v2145 = *(*int32)(unsafe.Add(mBase, uint32(v2126)))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	v2156 = base.I32_div_u_s(v2145+int32(_a_F_WalSummarizerMain_31), int32(_a_F_WalSummarizerMain_32))
	F_BlockRefTableSetLimitBlock(m, v1355, v2126+int32(4), int32(2), v2156)
	mBase = m.M
	v2158 = m.ExcPending
	if v2158 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L379
	}
L379:
	;
	goto L332
L380:
	;
	v2349 = *(*int32)(unsafe.Add(mBase, uint32(v1973)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	v2353 = v42 + int32(2600)
	v2354 = int32(0)
	base.MemoryFill(m, v2353, v2354, int32(264))
	v2360 = *(*int64)(unsafe.Add(mBase, uint32(v2349)))
	*(*int64)(unsafe.Add(mBase, uint32(v2353))) = v2360
	if v2354 <= base.I32_extend8_s(v2159) {
		goto L412
	} else {
		goto L413
	}
L381:
	;
	v2164 = *(*int32)(unsafe.Add(mBase, uint32(v1973)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	v2168 = v42 + int32(2864)
	v2169 = int32(0)
	base.MemoryFill(m, v2168, v2169, int32(288))
	v2175 = *(*int64)(unsafe.Add(mBase, uint32(v2164)))
	*(*int64)(unsafe.Add(mBase, uint32(v2168))) = v2175
	if v2169 <= base.I32_extend8_s(v2159) {
		goto L383
	} else {
		goto L384
	}
L382:
	;
	v2283 = int32(0)
	v2284 = *(*int32)(unsafe.Add(mBase, uint32(v42)+2892))
	if v2284 <= v2283 {
		goto L332
	} else {
		goto L404
	}
L383:
	;
	goto L382
L384:
	;
	v2180 = *(*int32)(unsafe.Add(mBase, uint32(v2164)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2168)+8)) = v2180
	if v2180&int32(1) != 0 {
		goto L385
	} else {
		goto L386
	}
L385:
	;
	v2184 = *(*int32)(unsafe.Add(mBase, uint32(v2164)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v2168)+12)) = v2184
	v2186 = *(*int32)(unsafe.Add(mBase, uint32(v2164)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v2168)+16)) = v2186
	v2192 = v2164 + int32(20)
	goto L387
L386:
	;
	v2192 = v2164 + int32(12)
	goto L387
L387:
	;
	if v2180&int32(2) != 0 {
		goto L388
	} else {
		goto L389
	}
L388:
	;
	v2195 = *(*int32)(unsafe.Add(mBase, uint32(v2192)))
	v2197 = v2192 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v2168)+24)) = v2197
	*(*int32)(unsafe.Add(mBase, uint32(v2168)+20)) = v2195
	v2203 = v2197 + v2195<<(uint(int32(2))%32)
	goto L390
L389:
	;
	v2203 = v2192
	goto L390
L390:
	;
	if v2180&int32(4) != 0 {
		goto L391
	} else {
		goto L392
	}
L391:
	;
	v2207 = *(*int32)(unsafe.Add(mBase, uint32(v2203)))
	v2209 = v2203 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v2168)+32)) = v2209
	*(*int32)(unsafe.Add(mBase, uint32(v2168)+28)) = v2207
	v2212 = *(*int32)(unsafe.Add(mBase, uint32(v2203)))
	v2216 = v2209 + v2212*int32(12)
	goto L393
L392:
	;
	v2216 = v2203
	goto L393
L393:
	;
	if v2180&int32(256) != 0 {
		goto L394
	} else {
		goto L395
	}
L394:
	;
	v2221 = *(*int32)(unsafe.Add(mBase, uint32(v2216)))
	v2222 = int32(4)
	v2223 = v2216 + v2222
	*(*int32)(unsafe.Add(mBase, uint32(v2168)+40)) = v2223
	*(*int32)(unsafe.Add(mBase, uint32(v2168)+36)) = v2221
	v2226 = *(*int32)(unsafe.Add(mBase, uint32(v2216)))
	v2230 = v2223 + v2226<<(uint(v2222)%32)
	goto L396
L395:
	;
	v2230 = v2216
	goto L396
L396:
	;
	if v2180&int32(8) != 0 {
		goto L397
	} else {
		goto L398
	}
L397:
	;
	v2235 = *(*int32)(unsafe.Add(mBase, uint32(v2230)))
	v2236 = int32(4)
	v2237 = v2230 + v2236
	*(*int32)(unsafe.Add(mBase, uint32(v2168)+48)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v2168)+44)) = v2235
	v2240 = *(*int32)(unsafe.Add(mBase, uint32(v2230)))
	v2244 = v2237 + v2240<<(uint(v2236)%32)
	goto L399
L398:
	;
	v2244 = v2230
	goto L399
L399:
	;
	if v2180&int32(16) == int32(0) {
		v2268 = v2180
		v2269 = v2244
		goto L400
	} else {
		goto L401
	}
L400:
	;
	if v2268&int32(32) == int32(0) {
		goto L383
	} else {
		goto L403
	}
L401:
	;
	v2251 = *(*int32)(unsafe.Add(mBase, uint32(v2244)))
	*(*int32)(unsafe.Add(mBase, uint32(v2168)+52)) = v2251
	v2254 = v2244 + int32(4)
	if v2180&int32(128) == int32(0) {
		v2268 = v2180
		v2269 = v2254
		goto L400
	} else {
		goto L402
	}
L402:
	;
	v2262 = F_strlcpy(m, v42+int32(2920), v2254, int32(200))
	mBase = m.M
	v2263 = F_strlen(m, v2254)
	mBase = m.M
	v2267 = *(*int32)(unsafe.Add(mBase, uint32(v2168)+8))
	v2268 = v2267
	v2269 = v2263 + v2254 + int32(1)
	goto L400
L403:
	;
	v2274 = *(*int64)(unsafe.Add(mBase, uint32(v2269)))
	v2275 = *(*int64)(unsafe.Add(mBase, uint32(v2269)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2168)+280)) = v2275
	*(*int64)(unsafe.Add(mBase, uint32(v2168)+272)) = v2274
	goto L383
L404:
	;
	v2291 = v2283
	goto L405
L405:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	v2322 = v2291 * int32(12)
	v2323 = *(*int32)(unsafe.Add(mBase, uint32(v42)+2896))
	v2325 = int32(0)
	F_BlockRefTableSetLimitBlock(m, v1355, v2322+v2323, v2325, v2325)
	mBase = m.M
	v2328 = m.ExcPending
	if v2328 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L407
	}
L406:
	;
	goto L332
L407:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	v2331 = *(*int32)(unsafe.Add(mBase, uint32(v42)+2896))
	F_BlockRefTableSetLimitBlock(m, v1355, v2331+v2322, int32(2), int32(0))
	mBase = m.M
	v2336 = m.ExcPending
	if v2336 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L408
	}
L408:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	v2339 = *(*int32)(unsafe.Add(mBase, uint32(v42)+2896))
	F_BlockRefTableSetLimitBlock(m, v1355, v2339+v2322, int32(3), int32(0))
	mBase = m.M
	v2344 = m.ExcPending
	if v2344 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L409
	}
L409:
	;
	v2346 = v2291 + int32(1)
	v2347 = *(*int32)(unsafe.Add(mBase, uint32(v42)+2892))
	if v2346 < v2347 {
		v2291 = v2346
		goto L405
	} else {
		goto L410
	}
L410:
	;
	goto L406
L411:
	;
	v2454 = int32(0)
	v2455 = *(*int32)(unsafe.Add(mBase, uint32(v42)+2628))
	if v2455 <= v2454 {
		goto L332
	} else {
		goto L430
	}
L412:
	;
	goto L411
L413:
	;
	v2365 = *(*int32)(unsafe.Add(mBase, uint32(v2349)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2353)+8)) = v2365
	if v2365&int32(1) != 0 {
		goto L414
	} else {
		goto L415
	}
L414:
	;
	v2369 = *(*int32)(unsafe.Add(mBase, uint32(v2349)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v2353)+12)) = v2369
	v2371 = *(*int32)(unsafe.Add(mBase, uint32(v2349)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v2353)+16)) = v2371
	v2377 = v2349 + int32(20)
	goto L416
L415:
	;
	v2377 = v2349 + int32(12)
	goto L416
L416:
	;
	if v2365&int32(2) != 0 {
		goto L417
	} else {
		goto L418
	}
L417:
	;
	v2380 = *(*int32)(unsafe.Add(mBase, uint32(v2377)))
	v2382 = v2377 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v2353)+24)) = v2382
	*(*int32)(unsafe.Add(mBase, uint32(v2353)+20)) = v2380
	v2388 = v2382 + v2380<<(uint(int32(2))%32)
	goto L419
L418:
	;
	v2388 = v2377
	goto L419
L419:
	;
	if v2365&int32(4) != 0 {
		goto L420
	} else {
		goto L421
	}
L420:
	;
	v2392 = *(*int32)(unsafe.Add(mBase, uint32(v2388)))
	v2394 = v2388 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v2353)+32)) = v2394
	*(*int32)(unsafe.Add(mBase, uint32(v2353)+28)) = v2392
	v2397 = *(*int32)(unsafe.Add(mBase, uint32(v2388)))
	v2401 = v2394 + v2397*int32(12)
	goto L422
L421:
	;
	v2401 = v2388
	goto L422
L422:
	;
	if v2365&int32(256) != 0 {
		goto L423
	} else {
		goto L424
	}
L423:
	;
	v2406 = *(*int32)(unsafe.Add(mBase, uint32(v2401)))
	v2407 = int32(4)
	v2408 = v2401 + v2407
	*(*int32)(unsafe.Add(mBase, uint32(v2353)+40)) = v2408
	*(*int32)(unsafe.Add(mBase, uint32(v2353)+36)) = v2406
	v2411 = *(*int32)(unsafe.Add(mBase, uint32(v2401)))
	v2415 = v2408 + v2411<<(uint(v2407)%32)
	goto L425
L424:
	;
	v2415 = v2401
	goto L425
L425:
	;
	if v2365&int32(16) == int32(0) {
		v2439 = v2365
		v2440 = v2415
		goto L426
	} else {
		goto L427
	}
L426:
	;
	if v2439&int32(32) == int32(0) {
		goto L412
	} else {
		goto L429
	}
L427:
	;
	v2422 = *(*int32)(unsafe.Add(mBase, uint32(v2415)))
	*(*int32)(unsafe.Add(mBase, uint32(v2353)+44)) = v2422
	v2425 = v2415 + int32(4)
	if v2365&int32(128) == int32(0) {
		v2439 = v2365
		v2440 = v2425
		goto L426
	} else {
		goto L428
	}
L428:
	;
	v2433 = F_strlcpy(m, v42+int32(2648), v2425, int32(200))
	mBase = m.M
	v2434 = F_strlen(m, v2425)
	mBase = m.M
	v2438 = *(*int32)(unsafe.Add(mBase, uint32(v2353)+8))
	v2439 = v2438
	v2440 = v2434 + v2425 + int32(1)
	goto L426
L429:
	;
	v2445 = *(*int64)(unsafe.Add(mBase, uint32(v2440)))
	v2446 = *(*int64)(unsafe.Add(mBase, uint32(v2440)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2353)+256)) = v2446
	*(*int64)(unsafe.Add(mBase, uint32(v2353)+248)) = v2445
	goto L412
L430:
	;
	v2462 = v2454
	goto L431
L431:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	v2493 = v2462 * int32(12)
	v2494 = *(*int32)(unsafe.Add(mBase, uint32(v42)+2632))
	v2496 = int32(0)
	F_BlockRefTableSetLimitBlock(m, v1355, v2493+v2494, v2496, v2496)
	mBase = m.M
	v2499 = m.ExcPending
	if v2499 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L433
	}
L432:
	;
	goto L332
L433:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	v2502 = *(*int32)(unsafe.Add(mBase, uint32(v42)+2632))
	F_BlockRefTableSetLimitBlock(m, v1355, v2502+v2493, int32(2), int32(0))
	mBase = m.M
	v2507 = m.ExcPending
	if v2507 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L434
	}
L434:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	v2510 = *(*int32)(unsafe.Add(mBase, uint32(v42)+2632))
	F_BlockRefTableSetLimitBlock(m, v1355, v2510+v2493, int32(3), int32(0))
	mBase = m.M
	v2515 = m.ExcPending
	if v2515 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L435
	}
L435:
	;
	v2517 = v2462 + int32(1)
	v2518 = *(*int32)(unsafe.Add(mBase, uint32(v42)+2628))
	if v2517 < v2518 {
		v2462 = v2517
		goto L431
	} else {
		goto L436
	}
L436:
	;
	goto L432
L437:
	;
	v2561 = v2552
	goto L438
L438:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	v2592 = v2561 & int32(255)
	v2594 = v42 + int32(468)
	v2596 = v42 + int32(464)
	v2598 = v42 + int32(460)
	v2599 = int32(0)
	v2601 = *(*int32)(unsafe.Add(mBase, uint32(v1381)+96))
	v2602 = *(*int32)(unsafe.Add(mBase, uint32(v2601)+72))
	if v2602 < v2592 {
		v2626 = v2599
		goto L442
	} else {
		goto L443
	}
L439:
	;
	v2648 = int32(0)
	goto L331
L440:
	;
	v2639 = v2561 + int32(1)
	v2640 = *(*int32)(unsafe.Add(mBase, uint32(v1381)+96))
	v2641 = *(*int32)(unsafe.Add(mBase, uint32(v2640)+72))
	if v2639 <= v2641 {
		v2561 = v2639
		goto L438
	} else {
		goto L458
	}
L441:
	;
	if v2626 == int32(0) {
		goto L440
	} else {
		goto L455
	}
L442:
	;
	goto L441
L443:
	;
	v2606 = v2601 + v2592*int32(52)
	v2607 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2606)+76)))
	if v2607 != int32(1) {
		v2626 = v2599
		goto L442
	} else {
		goto L444
	}
L444:
	;
	v2611 = v2606 + int32(76)
	if v2594 != 0 {
		goto L445
	} else {
		goto L446
	}
L445:
	;
	v2612 = *(*int32)(unsafe.Add(mBase, uint32(v2611)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v2594)+8)) = v2612
	v2614 = *(*int64)(unsafe.Add(mBase, uint32(v2611)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v2594))) = v2614
	goto L447
L446:
	;
	goto L447
L447:
	;
	if v2596 != 0 {
		goto L448
	} else {
		goto L449
	}
L448:
	;
	v2616 = *(*int32)(unsafe.Add(mBase, uint32(v2611)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v2596))) = v2616
	goto L450
L449:
	;
	goto L450
L450:
	;
	if v2598 != 0 {
		goto L451
	} else {
		goto L452
	}
L451:
	;
	v2618 = *(*int32)(unsafe.Add(mBase, uint32(v2611)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2598))) = v2618
	goto L453
L452:
	;
	goto L453
L453:
	;
	v2626 = int32(1)
	goto L442
L455:
	;
	v2629 = *(*int32)(unsafe.Add(mBase, uint32(v42)+464))
	if v2629 == int32(1) {
		goto L440
	} else {
		goto L456
	}
L456:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	v2634 = *(*int32)(unsafe.Add(mBase, uint32(v42)+460))
	F_BlockRefTableMarkBlockModified(m, v1355, v2594, v2629, v2634)
	mBase = m.M
	v2636 = m.ExcPending
	if v2636 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L457
	}
L457:
	;
	goto L440
L458:
	;
	goto L439
L459:
	;
	v2687 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[8]))
	*(*int64)(unsafe.Add(mBase, uint32(v2687)+24)) = v2676
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	v2692 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[7]))
	F_LWLockRelease(m, v2692+int32(_a_F_WalSummarizerMain_4))
	mBase = m.M
	v2696 = m.ExcPending
	if v2696 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L460
	}
L460:
	;
	if v1970 == int32(0) {
		goto L461
	} else {
		goto L462
	}
L461:
	;
	v2699 = *(*int64)(unsafe.Add(mBase, uint32(v1381)+40))
	if base.Ui64(v1879) <= base.Ui64(v2699) {
		v2784 = v2648
		v2806 = v2676
		goto L246
	} else {
		goto L464
	}
L462:
	;
	goto L463
L463:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	F_ProcessWalSummarizerInterrupts(m)
	mBase = m.M
	v2704 = m.ExcPending
	if v2704 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L465
	}
L464:
	;
	goto L463
L465:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	v2709 = F_XLogReadRecord(m, v1381, v42+int32(492))
	mBase = m.M
	v2710 = m.ExcPending
	if v2710 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L466
	}
L466:
	;
	if v2709 != 0 {
		v1937 = v2648
		goto L326
	} else {
		goto L467
	}
L467:
	;
	goto L327
L468:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	v2750 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2751 = m.ExcPending
	if v2751 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L469
	}
L469:
	;
	if v2750 != 0 {
		goto L470
	} else {
		goto L471
	}
L470:
	;
	v2752 = *(*int64)(unsafe.Add(mBase, uint32(v1381)+40))
	v2753 = *(*int64)(unsafe.Add(mBase, uint32(v1360)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	*(*uint32)(unsafe.Add(mBase, uint32(v59))) = uint32(v2753)
	v2757 = int64(32)
	v2758 = int64(base.Ui64(v2753) >> (uint(v2757) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v42)+124)) = uint32(v2758)
	*(*uint32)(unsafe.Add(mBase, uint32(v42)+120)) = uint32(v2752)
	v2762 = int64(base.Ui64(v2752) >> (uint(v2757) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v42)+116)) = uint32(v2762)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+112)) = v1332
	F_errmsg_internal(m, int32(_a_F_WalSummarizerMain_28), v42+int32(112))
	mBase = m.M
	v2769 = m.ExcPending
	if v2769 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L473
	}
L471:
	;
	goto L472
L472:
	;
	v2779 = *(*int64)(unsafe.Add(mBase, uint32(v1360)+8))
	v2784 = v2715
	v2806 = v2779
	goto L246
L473:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	F_errfinish(m, int32(_a_F_WalSummarizerMain_1), int32(1142), int32(_a_F_WalSummarizerMain_24))
	mBase = m.M
	v2776 = m.ExcPending
	if v2776 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L474
	}
L474:
	;
	goto L472
L475:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	F_XLogReaderFree(m, v1381)
	mBase = m.M
	v2852 = m.ExcPending
	if v2852 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L476
	}
L476:
	;
	v2853 = base.B2i32(base.Ui64(v2838) <= base.Ui64(v1874))
	if (v2853|v2812)&int32(1) != 0 {
		goto L242
	} else {
		goto L477
	}
L477:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	v2860 = v42 + int32(1536)
	v2864 = F_pg_snprintf(m, v2860, int32(1024), int32(_a_F_WalSummarizerMain_33), int32(0))
	mBase = m.M
	v2865 = m.ExcPending
	if v2865 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L478
	}
L478:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	v2868 = base.I32_wrap_i64(v2838)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+96)) = v2868
	v2870 = int64(32)
	v2872 = base.I32_wrap_i64(int64(base.Ui64(v2838) >> (uint(v2870) % 64)))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+92)) = v2872
	v2874 = base.I32_wrap_i64(v1874)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+88)) = v2874
	v2878 = base.I32_wrap_i64(int64(base.Ui64(v1874) >> (uint(v2870) % 64)))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+84)) = v2878
	*(*int32)(unsafe.Add(mBase, uint32(v42)+80)) = v1332
	v2887 = F_pg_snprintf(m, v42+int32(512), int32(1024), int32(_a_F_WalSummarizerMain_11), v42+int32(80))
	mBase = m.M
	v2888 = m.ExcPending
	if v2888 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L479
	}
L479:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	*(*int64)(unsafe.Add(mBase, uint32(v42)+504)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	v2894 = F_PathNameOpenFile(m, v2860, int32(577))
	mBase = m.M
	v2895 = m.ExcPending
	if v2895 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L480
	}
L480:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+496)) = v2894
	if int32(0) <= v2894 {
		goto L243
	} else {
		goto L481
	}
L481:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2904 = m.ExcPending
	if v2904 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L482
	}
L482:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	F_errcode_for_file_access(m)
	mBase = m.M
	v2908 = m.ExcPending
	if v2908 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L483
	}
L483:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	*(*int32)(unsafe.Add(mBase, uint32(v42)+32)) = v2860
	F_errmsg(m, int32(_a_F_WalSummarizerMain_34), v42+int32(32))
	mBase = m.M
	v2916 = m.ExcPending
	if v2916 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L484
	}
L484:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	F_errfinish(m, int32(_a_F_WalSummarizerMain_1), int32(1317), int32(_a_F_WalSummarizerMain_24))
	mBase = m.M
	v2923 = m.ExcPending
	if v2923 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L485
	}
L485:
	;
	goto L3
L486:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	F_errcode_for_file_access(m)
	mBase = m.M
	v2934 = m.ExcPending
	if v2934 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L487
	}
L487:
	;
	v2935 = *(*int64)(unsafe.Add(mBase, uint32(v1381)+40))
	v2938 = base.I32_wrap_i64(int64(base.Ui64(v2935) >> (uint(int64(32)) % 64)))
	v2939 = base.I32_wrap_i64(v2935)
	if v2926 != 0 {
		goto L488
	} else {
		goto L489
	}
L488:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	v2942 = *(*int32)(unsafe.Add(mBase, uint32(v42)+492))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+172)) = v2942
	*(*int32)(unsafe.Add(mBase, uint32(v42)+168)) = v2939
	*(*int32)(unsafe.Add(mBase, uint32(v42)+164)) = v2938
	*(*int32)(unsafe.Add(mBase, uint32(v42)+160)) = v1332
	F_errmsg(m, int32(_a_F_WalSummarizerMain_35), v42+int32(160))
	mBase = m.M
	v2951 = m.ExcPending
	if v2951 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L491
	}
L489:
	;
	goto L490
L490:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	*(*int32)(unsafe.Add(mBase, uint32(v42)+152)) = v2939
	*(*int32)(unsafe.Add(mBase, uint32(v42)+148)) = v2938
	*(*int32)(unsafe.Add(mBase, uint32(v42)+144)) = v1332
	F_errmsg(m, int32(_a_F_WalSummarizerMain_36), v42+int32(144))
	mBase = m.M
	v2968 = m.ExcPending
	if v2968 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L493
	}
L491:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	F_errfinish(m, int32(_a_F_WalSummarizerMain_1), int32(1152), int32(_a_F_WalSummarizerMain_24))
	mBase = m.M
	v2958 = m.ExcPending
	if v2958 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L492
	}
L492:
	;
	goto L3
L493:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	F_errfinish(m, int32(_a_F_WalSummarizerMain_1), int32(1157), int32(_a_F_WalSummarizerMain_24))
	mBase = m.M
	v2975 = m.ExcPending
	if v2975 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L494
	}
L494:
	;
	goto L3
L495:
	;
	v3862 = m.G0
	v3864 = v3862 - int32(32)
	m.G0 = v3864
	v3866 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3864)+24)) = v3866
	*(*int64)(unsafe.Add(mBase, uint32(v3864)+16)) = v3866
	*(*int64)(unsafe.Add(mBase, uint32(v3864)+8)) = v3866
	v3873 = v2981 + int32(12)
	v3874 = *(*int32)(unsafe.Add(mBase, uint32(v2981)+uint32(_c_F_WalSummarizerMain[21])))
	v3877 = int32(24)
	v3878 = m.Env.Pgmem_crc32c(m, v3874, v3864+int32(8), v3877)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v2981)+uint32(_c_F_WalSummarizerMain[21]))) = v3878
	v3880 = *(*int32)(unsafe.Add(mBase, uint32(v2981)+uint32(_c_F_WalSummarizerMain[22])))
	if int32(_a_F_WalSummarizerMain_37) <= v3880+v3877 {
		goto L607
	} else {
		goto L608
	}
L496:
	;
	v3011 = F_palloc_mul(m, int32(24), v3007)
	mBase = m.M
	v3012 = m.ExcPending
	if v3012 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L497
	}
L497:
	;
	v3013 = *(*int32)(unsafe.Add(mBase, uint32(v1355)))
	v3014 = *(*int64)(unsafe.Add(mBase, uint32(v3013)))
	if v3014 == int64(0) {
		v3065 = v2995
		goto L498
	} else {
		goto L499
	}
L498:
	;
	v3094 = v2981 + int32(20)
	v3105 = v3065
	v3106 = int32(0)
	v3111 = v3013
	v3113 = v2978
	goto L506
L499:
	;
	v3017 = *(*int32)(unsafe.Add(mBase, uint32(v3013)+20))
	v3026 = int32(0)
	goto L500
L500:
	;
	v3054 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3017+v3026*int32(40))+20)))
	if v3054 != int32(1) {
		goto L502
	} else {
		goto L503
	}
L501:
	;
	v3065 = v2995
	goto L498
L502:
	;
	v3065 = v3026
	goto L498
L503:
	;
	goto L504
L504:
	;
	v3058 = v3026 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v3058)) < base.Ui64(v3014) {
		v3026 = v3058
		goto L500
	} else {
		goto L505
	}
L505:
	;
	goto L501
L506:
	;
	v3133 = v3106
	v3135 = v3105
	v3138 = v3106
	goto L509
L507:
	;
	F_pg_qsort(m, v3011, v3113, int32(24), int32(_a_F_WalSummarizerMain_38))
	mBase = m.M
	v3240 = m.ExcPending
	if v3240 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L518
	}
L508:
	;
	goto L507
L509:
	;
	if v3133&int32(1) != 0 {
		goto L508
	} else {
		goto L511
	}
L510:
	;
	v3180 = v3011 + v3113*int32(24)
	v3181 = *(*int32)(unsafe.Add(mBase, uint32(v3174)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v3180)+8)) = v3181
	v3183 = *(*int64)(unsafe.Add(mBase, uint32(v3174)))
	*(*int64)(unsafe.Add(mBase, uint32(v3180))) = v3183
	v3185 = *(*int32)(unsafe.Add(mBase, uint32(v3174)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v3180)+12)) = v3185
	v3187 = *(*int32)(unsafe.Add(mBase, uint32(v3174)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v3180)+16)) = v3187
	v3189 = *(*int32)(unsafe.Add(mBase, uint32(v3174)+24))
	v3197 = v3189
	goto L513
L511:
	;
	v3162 = *(*int32)(unsafe.Add(mBase, uint32(v3111)+12))
	v3163 = int32(1)
	v3164 = v3135 - v3163
	v3168 = base.B2i32(v3162&(v3164^v3065) == int32(0))
	v3169 = v3168 | v3138
	v3172 = v3164 & v3162
	v3173 = *(*int32)(unsafe.Add(mBase, uint32(v3111)+20))
	v3174 = v3135*int32(40) + v3173
	v3175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3174)+20)))
	if v3175 != v3163 {
		v3133 = v3168
		v3135 = v3172
		v3138 = v3169
		goto L509
	} else {
		goto L512
	}
L512:
	;
	goto L510
L513:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3180)+20)) = v3197
	if v3197 == int32(0) {
		goto L515
	} else {
		goto L516
	}
L514:
	;
	v3236 = *(*int32)(unsafe.Add(mBase, uint32(v1355)))
	v3105 = v3172
	v3106 = v3169
	v3111 = v3236
	v3113 = v3113 + int32(1)
	goto L506
L515:
	;
	goto L514
L516:
	;
	v3225 = *(*int32)(unsafe.Add(mBase, uint32(v3174)+32))
	v3231 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3225+v3197<<(uint(int32(1))%32)-int32(2)))))
	if v3231 != 0 {
		goto L515
	} else {
		goto L517
	}
L517:
	;
	v3197 = v3197 - int32(1)
	goto L513
L518:
	;
	v3241 = *(*int32)(unsafe.Add(mBase, uint32(v1355)))
	v3242 = *(*int32)(unsafe.Add(mBase, uint32(v3241)+8))
	if v3242 == int32(0) {
		goto L495
	} else {
		goto L519
	}
L519:
	;
	v3252 = int32(0)
	goto L520
L520:
	;
	v3278 = *(*int32)(unsafe.Add(mBase, uint32(v2981)+uint32(_c_F_WalSummarizerMain[21])))
	v3279 = int32(24)
	v3281 = v3011 + v3252*v3279
	v3283 = m.Env.Pgmem_crc32c(m, v3278, v3281, v3279)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v2981)+uint32(_c_F_WalSummarizerMain[21]))) = v3283
	v3285 = *(*int32)(unsafe.Add(mBase, uint32(v2981)+uint32(_c_F_WalSummarizerMain[22])))
	if int32(_a_F_WalSummarizerMain_37) <= v3285+v3279 {
		goto L522
	} else {
		goto L523
	}
L521:
	;
	goto L495
L522:
	;
	v3290 = *(*int32)(unsafe.Add(mBase, uint32(v2981)+16))
	v3291 = *(*int32)(unsafe.Add(mBase, uint32(v2981)+12))
	v3292 = m.T0[v3291].(func(*base.Module, int32, int32, int32) int32)(m, v3290, v3094, v3285)
	mBase = m.M
	v3293 = m.ExcPending
	if v3293 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L525
	}
L523:
	;
	v3297 = v3285
	goto L524
L524:
	;
	v3298 = v3297 + v3094
	v3299 = *(*int64)(unsafe.Add(mBase, uint32(v3281)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v3298)+16)) = v3299
	v3301 = *(*int64)(unsafe.Add(mBase, uint32(v3281)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v3298)+8)) = v3301
	v3303 = *(*int64)(unsafe.Add(mBase, uint32(v3281)))
	*(*int64)(unsafe.Add(mBase, uint32(v3298))) = v3303
	v3305 = *(*int32)(unsafe.Add(mBase, uint32(v2981)+uint32(_c_F_WalSummarizerMain[22])))
	*(*int32)(unsafe.Add(mBase, uint32(v2981)+uint32(_c_F_WalSummarizerMain[22]))) = v3305 + int32(24)
	v3309 = *(*int32)(unsafe.Add(mBase, uint32(v1355)))
	v3310 = *(*int32)(unsafe.Add(mBase, uint32(v3281)+12))
	v3311 = *(*int32)(unsafe.Add(mBase, uint32(v3281)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2981)+uint32(_c_F_WalSummarizerMain[23]))) = v3311
	v3313 = *(*int64)(unsafe.Add(mBase, uint32(v3281)))
	*(*int64)(unsafe.Add(mBase, uint32(v2981)+uint32(_c_F_WalSummarizerMain[24]))) = v3313
	*(*int32)(unsafe.Add(mBase, uint32(v2981)+uint32(_c_F_WalSummarizerMain[25]))) = v3310
	v3317 = v2981 + int32(_a_F_WalSummarizerMain_39)
	v3318 = int32(16)
	v3324 = int32(-1636608416)
	if v3317&int32(3) != 0 {
		goto L530
	} else {
		goto L531
	}
L525:
	;
	v3294 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2981)+uint32(_c_F_WalSummarizerMain[22]))) = v3294
	v3297 = v3294
	goto L524
L526:
	;
	v3583 = *(*int32)(unsafe.Add(mBase, uint32(v3309)+20))
	v3584 = *(*int32)(unsafe.Add(mBase, uint32(v3309)+12))
	v3585 = (v3578 ^ v3570 - base.I32_rotl(v3578, int32(24))) & v3584
	v3588 = v3583 + v3585*int32(40)
	v3589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3588)+20)))
	if v3589 == int32(0) {
		goto L567
	} else {
		goto L568
	}
L527:
	;
	v3556 = int32(14)
	v3558 = v3552 ^ v3553 - base.I32_rotl(v3552, v3556)
	v3562 = v3558 ^ v3551 - base.I32_rotl(v3558, int32(11))
	v3566 = v3562 ^ v3552 - base.I32_rotl(v3562, int32(25))
	v3570 = v3566 ^ v3558 - base.I32_rotl(v3566, int32(16))
	v3574 = v3570 ^ v3562 - base.I32_rotl(v3570, int32(4))
	v3578 = v3574 ^ v3566 - base.I32_rotl(v3574, v3556)
	goto L526
L528:
	;
	switch v3478 - int32(1) {
	case 0:
		v3544 = v3469
		v3545 = v3470
		v3546 = v3474
		goto L555
	case 1:
		v3537 = v3469
		v3538 = v3470
		v3539 = v3474
		goto L556
	case 2:
		v3530 = v3469
		v3531 = v3470
		v3532 = v3474
		goto L557
	case 3:
		v3524 = v3470
		v3525 = v3474
		goto L558
	case 4:
		v3520 = v3470
		v3521 = v3474
		goto L559
	case 5:
		v3514 = v3470
		v3515 = v3474
		goto L560
	case 6:
		v3508 = v3470
		v3509 = v3474
		goto L561
	case 7:
		v3503 = v3474
		goto L562
	case 8:
		v3498 = v3474
		goto L563
	case 9:
		v3493 = v3474
		goto L564
	case 10:
		goto L565
	default:
		v3551 = v3469
		v3552 = v3470
		v3553 = v3474
		goto L527
	}
L529:
	;
	v3433 = v3317
	v3434 = v3318
	v3435 = v3324
	v3436 = v3324
	v3437 = v3324
	goto L552
L530:
	;
	goto L529
L531:
	;
	goto L532
L532:
	;
	goto L536
L534:
	;
	switch v3376 - int32(1) {
	case 0:
		v3430 = v3367
		goto L541
	case 1:
		v3425 = v3367
		goto L542
	case 2:
		goto L543
	case 3:
		v3418 = v3368
		goto L544
	case 4:
		v3415 = v3368
		goto L545
	case 5:
		v3410 = v3368
		goto L546
	case 6:
		goto L547
	case 7:
		v3401 = v3372
		goto L548
	case 8:
		v3396 = v3372
		goto L549
	case 9:
		v3391 = v3372
		goto L550
	case 10:
		goto L551
	default:
		v3551 = v3367
		v3552 = v3368
		v3553 = v3372
		goto L527
	}
L536:
	;
	goto L537
L537:
	;
	v3331 = v3317
	v3332 = v3318
	v3333 = v3324
	v3334 = v3324
	v3335 = v3324
	goto L538
L538:
	;
	v3337 = *(*int32)(unsafe.Add(mBase, uint32(v3331)+4))
	v3338 = v3337 + v3334
	v3339 = *(*int32)(unsafe.Add(mBase, uint32(v3331)))
	v3341 = *(*int32)(unsafe.Add(mBase, uint32(v3331)+8))
	v3342 = v3341 + v3335
	v3344 = int32(4)
	v3346 = v3339 + v3333 - v3342 ^ base.I32_rotl(v3342, v3344)
	v3350 = v3338 - v3346 ^ base.I32_rotl(v3346, int32(6))
	v3351 = v3342 + v3338
	v3352 = v3346 + v3351
	v3353 = v3350 + v3352
	v3357 = v3351 - v3350 ^ base.I32_rotl(v3350, int32(8))
	v3361 = v3352 - v3357 ^ base.I32_rotl(v3357, int32(16))
	v3365 = v3353 - v3361 ^ base.I32_rotl(v3361, int32(19))
	v3366 = v3357 + v3353
	v3367 = v3361 + v3366
	v3368 = v3365 + v3367
	v3372 = v3366 - v3365 ^ base.I32_rotl(v3365, v3344)
	v3373 = int32(12)
	v3374 = v3331 + v3373
	v3376 = v3332 - v3373
	if base.Ui32(int32(11)) < base.Ui32(v3376) {
		v3331 = v3374
		v3332 = v3376
		v3333 = v3367
		v3334 = v3368
		v3335 = v3372
		goto L538
	} else {
		goto L540
	}
L539:
	;
	goto L534
L540:
	;
	goto L539
L541:
	;
	v3431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3374))))
	v3551 = v3430 + v3431
	v3552 = v3368
	v3553 = v3372
	goto L527
L542:
	;
	v3426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3374)+1)))
	v3430 = v3426<<(uint(int32(8))%32) + v3425
	goto L541
L543:
	;
	v3421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3374)+2)))
	v3425 = v3421<<(uint(int32(16))%32) + v3367
	goto L542
L544:
	;
	v3419 = *(*int32)(unsafe.Add(mBase, uint32(v3374)))
	v3551 = v3419 + v3367
	v3552 = v3418
	v3553 = v3372
	goto L527
L545:
	;
	v3416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3374)+4)))
	v3418 = v3415 + v3416
	goto L544
L546:
	;
	v3411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3374)+5)))
	v3415 = v3411<<(uint(int32(8))%32) + v3410
	goto L545
L547:
	;
	v3406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3374)+6)))
	v3410 = v3406<<(uint(int32(16))%32) + v3368
	goto L546
L548:
	;
	v3402 = *(*int32)(unsafe.Add(mBase, uint32(v3374)))
	v3404 = *(*int32)(unsafe.Add(mBase, uint32(v3374)+4))
	v3551 = v3402 + v3367
	v3552 = v3404 + v3368
	v3553 = v3401
	goto L527
L549:
	;
	v3397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3374)+8)))
	v3401 = v3397<<(uint(int32(8))%32) + v3396
	goto L548
L550:
	;
	v3392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3374)+9)))
	v3396 = v3392<<(uint(int32(16))%32) + v3391
	goto L549
L551:
	;
	v3387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3374)+10)))
	v3391 = v3387<<(uint(int32(24))%32) + v3372
	goto L550
L552:
	;
	v3439 = *(*int32)(unsafe.Add(mBase, uint32(v3433)+4))
	v3440 = v3439 + v3436
	v3441 = *(*int32)(unsafe.Add(mBase, uint32(v3433)))
	v3443 = *(*int32)(unsafe.Add(mBase, uint32(v3433)+8))
	v3444 = v3443 + v3437
	v3446 = int32(4)
	v3448 = v3441 + v3435 - v3444 ^ base.I32_rotl(v3444, v3446)
	v3452 = v3440 - v3448 ^ base.I32_rotl(v3448, int32(6))
	v3453 = v3444 + v3440
	v3454 = v3448 + v3453
	v3455 = v3452 + v3454
	v3459 = v3453 - v3452 ^ base.I32_rotl(v3452, int32(8))
	v3463 = v3454 - v3459 ^ base.I32_rotl(v3459, int32(16))
	v3467 = v3455 - v3463 ^ base.I32_rotl(v3463, int32(19))
	v3468 = v3459 + v3455
	v3469 = v3463 + v3468
	v3470 = v3467 + v3469
	v3474 = v3468 - v3467 ^ base.I32_rotl(v3467, v3446)
	v3475 = int32(12)
	v3476 = v3433 + v3475
	v3478 = v3434 - v3475
	if base.Ui32(int32(11)) < base.Ui32(v3478) {
		v3433 = v3476
		v3434 = v3478
		v3435 = v3469
		v3436 = v3470
		v3437 = v3474
		goto L552
	} else {
		goto L554
	}
L553:
	;
	goto L528
L554:
	;
	goto L553
L555:
	;
	v3547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3476))))
	v3551 = v3544 + v3547
	v3552 = v3545
	v3553 = v3546
	goto L527
L556:
	;
	v3540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3476)+1)))
	v3544 = v3540<<(uint(int32(8))%32) + v3537
	v3545 = v3538
	v3546 = v3539
	goto L555
L557:
	;
	v3533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3476)+2)))
	v3537 = v3533<<(uint(int32(16))%32) + v3530
	v3538 = v3531
	v3539 = v3532
	goto L556
L558:
	;
	v3526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3476)+3)))
	v3530 = v3526<<(uint(int32(24))%32) + v3469
	v3531 = v3524
	v3532 = v3525
	goto L557
L559:
	;
	v3522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3476)+4)))
	v3524 = v3520 + v3522
	v3525 = v3521
	goto L558
L560:
	;
	v3516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3476)+5)))
	v3520 = v3516<<(uint(int32(8))%32) + v3514
	v3521 = v3515
	goto L559
L561:
	;
	v3510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3476)+6)))
	v3514 = v3510<<(uint(int32(16))%32) + v3508
	v3515 = v3509
	goto L560
L562:
	;
	v3504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3476)+7)))
	v3508 = v3504<<(uint(int32(24))%32) + v3470
	v3509 = v3503
	goto L561
L563:
	;
	v3499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3476)+8)))
	v3503 = v3499<<(uint(int32(8))%32) + v3498
	goto L562
L564:
	;
	v3494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3476)+9)))
	v3498 = v3494<<(uint(int32(16))%32) + v3493
	goto L563
L565:
	;
	v3489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3476)+10)))
	v3493 = v3489<<(uint(int32(24))%32) + v3474
	goto L564
L566:
	;
	v3676 = *(*int32)(unsafe.Add(mBase, uint32(v3281)+20))
	if v3676 == int32(0) {
		goto L574
	} else {
		goto L575
	}
L567:
	;
	v3651 = int32(0)
	goto L566
L568:
	;
	goto L569
L569:
	;
	v3600 = v3588
	v3603 = v3585
	goto L570
L570:
	;
	v3625 = *(*int64)(unsafe.Add(mBase, uint32(v3600)))
	v3626 = *(*int64)(unsafe.Add(mBase, uint32(v2981)+uint32(_c_F_WalSummarizerMain[24])))
	v3628 = *(*int64)(unsafe.Add(mBase, uint32(v3600)+8))
	v3631 = *(*int64)(unsafe.Add(mBase, uint32(v2981)+uint32(_c_F_WalSummarizerMain[23])))
	if v3625^v3626|(v3628^v3631) == int64(0) {
		v3651 = v3600
		goto L566
	} else {
		goto L572
	}
L571:
	;
	v3651 = int32(0)
	goto L566
L572:
	;
	v3638 = (v3603 + int32(1)) & v3584
	v3641 = v3583 + v3638*int32(40)
	v3642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3641)+20)))
	if v3642 != 0 {
		v3600 = v3641
		v3603 = v3638
		goto L570
	} else {
		goto L573
	}
L573:
	;
	goto L571
L574:
	;
	v3711 = *(*int32)(unsafe.Add(mBase, uint32(v3651)+24))
	if v3711 != 0 {
		goto L587
	} else {
		goto L588
	}
L575:
	;
	v3679 = *(*int32)(unsafe.Add(mBase, uint32(v2981)+uint32(_c_F_WalSummarizerMain[21])))
	v3680 = *(*int32)(unsafe.Add(mBase, uint32(v3651)+32))
	v3682 = v3676 << (uint(int32(1)) % 32)
	v3683 = m.Env.Pgmem_crc32c(m, v3679, v3680, v3682)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v2981)+uint32(_c_F_WalSummarizerMain[21]))) = v3683
	v3685 = *(*int32)(unsafe.Add(mBase, uint32(v2981)+uint32(_c_F_WalSummarizerMain[22])))
	if int32(_a_F_WalSummarizerMain_37) <= v3685+v3682 {
		goto L576
	} else {
		goto L577
	}
L576:
	;
	v3689 = *(*int32)(unsafe.Add(mBase, uint32(v2981)+16))
	v3690 = *(*int32)(unsafe.Add(mBase, uint32(v2981)+12))
	v3691 = m.T0[v3690].(func(*base.Module, int32, int32, int32) int32)(m, v3689, v3094, v3685)
	mBase = m.M
	v3692 = m.ExcPending
	if v3692 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L579
	}
L577:
	;
	v3696 = v3685
	goto L578
L578:
	;
	if int32(_a_F_WalSummarizerMain_40) <= v3682 {
		goto L580
	} else {
		goto L581
	}
L579:
	;
	v3693 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2981)+uint32(_c_F_WalSummarizerMain[22]))) = v3693
	v3696 = v3693
	goto L578
L580:
	;
	v3699 = *(*int32)(unsafe.Add(mBase, uint32(v2981)+16))
	v3700 = *(*int32)(unsafe.Add(mBase, uint32(v2981)+12))
	v3701 = m.T0[v3700].(func(*base.Module, int32, int32, int32) int32)(m, v3699, v3680, v3682)
	mBase = m.M
	v3702 = m.ExcPending
	if v3702 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L583
	}
L581:
	;
	goto L582
L582:
	;
	if v3682 != 0 {
		goto L584
	} else {
		goto L585
	}
L583:
	;
	goto L574
L584:
	;
	base.MemoryCopy(m, v3696+v3094, v3680, v3682)
	goto L586
L585:
	;
	goto L586
L586:
	;
	v3705 = *(*int32)(unsafe.Add(mBase, uint32(v2981)+uint32(_c_F_WalSummarizerMain[22])))
	*(*int32)(unsafe.Add(mBase, uint32(v2981)+uint32(_c_F_WalSummarizerMain[22]))) = v3705 + v3682
	goto L574
L587:
	;
	v3723 = int32(0)
	goto L590
L588:
	;
	goto L589
L589:
	;
	v3826 = v3252 + int32(1)
	v3827 = *(*int32)(unsafe.Add(mBase, uint32(v1355)))
	v3828 = *(*int32)(unsafe.Add(mBase, uint32(v3827)+8))
	if base.Ui32(v3826) < base.Ui32(v3828) {
		v3252 = v3826
		goto L520
	} else {
		goto L606
	}
L590:
	;
	v3745 = *(*int32)(unsafe.Add(mBase, uint32(v3651)+32))
	v3749 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3745+v3723<<(uint(int32(1))%32)))))
	if v3749 == int32(0) {
		goto L592
	} else {
		goto L593
	}
L591:
	;
	goto L589
L592:
	;
	v3790 = v3723 + int32(1)
	v3791 = *(*int32)(unsafe.Add(mBase, uint32(v3651)+24))
	if base.Ui32(v3790) < base.Ui32(v3791) {
		v3723 = v3790
		goto L590
	} else {
		goto L605
	}
L593:
	;
	v3753 = *(*int32)(unsafe.Add(mBase, uint32(v2981)+uint32(_c_F_WalSummarizerMain[21])))
	v3754 = *(*int32)(unsafe.Add(mBase, uint32(v3651)+36))
	v3758 = *(*int32)(unsafe.Add(mBase, uint32(v3754+v3723<<(uint(int32(2))%32))))
	v3760 = v3749 << (uint(int32(1)) % 32)
	v3761 = m.Env.Pgmem_crc32c(m, v3753, v3758, v3760)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v2981)+uint32(_c_F_WalSummarizerMain[21]))) = v3761
	v3763 = *(*int32)(unsafe.Add(mBase, uint32(v2981)+uint32(_c_F_WalSummarizerMain[22])))
	if int32(_a_F_WalSummarizerMain_37) <= v3763+v3760 {
		goto L594
	} else {
		goto L595
	}
L594:
	;
	v3767 = *(*int32)(unsafe.Add(mBase, uint32(v2981)+16))
	v3768 = *(*int32)(unsafe.Add(mBase, uint32(v2981)+12))
	v3769 = m.T0[v3768].(func(*base.Module, int32, int32, int32) int32)(m, v3767, v3094, v3763)
	mBase = m.M
	v3770 = m.ExcPending
	if v3770 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L597
	}
L595:
	;
	v3774 = v3763
	goto L596
L596:
	;
	if base.I32_extend16_s(v3749) < int32(0) {
		goto L598
	} else {
		goto L599
	}
L597:
	;
	v3771 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2981)+uint32(_c_F_WalSummarizerMain[22]))) = v3771
	v3774 = v3771
	goto L596
L598:
	;
	v3777 = *(*int32)(unsafe.Add(mBase, uint32(v2981)+16))
	v3778 = *(*int32)(unsafe.Add(mBase, uint32(v2981)+12))
	v3779 = m.T0[v3778].(func(*base.Module, int32, int32, int32) int32)(m, v3777, v3758, v3760)
	mBase = m.M
	v3780 = m.ExcPending
	if v3780 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L601
	}
L599:
	;
	goto L600
L600:
	;
	if v3760 != 0 {
		goto L602
	} else {
		goto L603
	}
L601:
	;
	goto L592
L602:
	;
	base.MemoryCopy(m, v3774+v3094, v3758, v3760)
	goto L604
L603:
	;
	goto L604
L604:
	;
	v3783 = *(*int32)(unsafe.Add(mBase, uint32(v2981)+uint32(_c_F_WalSummarizerMain[22])))
	*(*int32)(unsafe.Add(mBase, uint32(v2981)+uint32(_c_F_WalSummarizerMain[22]))) = v3783 + v3760
	goto L592
L605:
	;
	goto L591
L606:
	;
	goto L521
L607:
	;
	v3885 = *(*int32)(unsafe.Add(mBase, uint32(v3873)+4))
	v3888 = *(*int32)(unsafe.Add(mBase, uint32(v3873)))
	v3889 = m.T0[v3888].(func(*base.Module, int32, int32, int32) int32)(m, v3885, v2981+int32(20), v3880)
	mBase = m.M
	v3890 = m.ExcPending
	if v3890 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L610
	}
L608:
	;
	v3894 = v3880
	goto L609
L609:
	;
	v3896 = v2981 + int32(20)
	v3897 = v3894 + v3896
	v3898 = *(*int64)(unsafe.Add(mBase, uint32(v3864)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v3897)+16)) = v3898
	v3900 = *(*int64)(unsafe.Add(mBase, uint32(v3864)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v3897)+8)) = v3900
	v3902 = *(*int64)(unsafe.Add(mBase, uint32(v3864)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v3897))) = v3902
	v3904 = *(*int32)(unsafe.Add(mBase, uint32(v2981)+uint32(_c_F_WalSummarizerMain[22])))
	*(*int32)(unsafe.Add(mBase, uint32(v2981)+uint32(_c_F_WalSummarizerMain[22]))) = v3904 + int32(24)
	v3908 = *(*int32)(unsafe.Add(mBase, uint32(v2981)+uint32(_c_F_WalSummarizerMain[21])))
	*(*int32)(unsafe.Add(mBase, uint32(v3864)+4)) = v3908 ^ int32(-1)
	v3912 = int32(4)
	v3915 = m.Env.Pgmem_crc32c(m, v3908, v3864+v3912, v3912)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v2981)+uint32(_c_F_WalSummarizerMain[21]))) = v3915
	v3917 = *(*int32)(unsafe.Add(mBase, uint32(v2981)+uint32(_c_F_WalSummarizerMain[22])))
	if int32(_a_F_WalSummarizerMain_37) <= v3917+v3912 {
		goto L611
	} else {
		goto L612
	}
L610:
	;
	v3891 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2981)+uint32(_c_F_WalSummarizerMain[22]))) = v3891
	v3894 = v3891
	goto L609
L611:
	;
	v3922 = *(*int32)(unsafe.Add(mBase, uint32(v3873)+4))
	v3923 = *(*int32)(unsafe.Add(mBase, uint32(v3873)))
	v3924 = m.T0[v3923].(func(*base.Module, int32, int32, int32) int32)(m, v3922, v3896, v3917)
	mBase = m.M
	v3925 = m.ExcPending
	if v3925 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L614
	}
L612:
	;
	v3929 = v3917
	goto L613
L613:
	;
	v3931 = *(*int32)(unsafe.Add(mBase, uint32(v3864)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v3929+v3896))) = v3931
	v3933 = *(*int32)(unsafe.Add(mBase, uint32(v2981)+uint32(_c_F_WalSummarizerMain[22])))
	v3935 = v3933 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v2981)+uint32(_c_F_WalSummarizerMain[22]))) = v3935
	v3937 = *(*int32)(unsafe.Add(mBase, uint32(v3873)+4))
	v3938 = *(*int32)(unsafe.Add(mBase, uint32(v3873)))
	v3939 = m.T0[v3938].(func(*base.Module, int32, int32, int32) int32)(m, v3937, v3896, v3935)
	mBase = m.M
	v3940 = m.ExcPending
	if v3940 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L615
	}
L614:
	;
	v3926 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2981)+uint32(_c_F_WalSummarizerMain[22]))) = v3926
	v3929 = v3926
	goto L613
L615:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2981)+uint32(_c_F_WalSummarizerMain[22]))) = int32(0)
	m.G0 = v3864 + int32(32)
	m.G0 = v2981 + int32(_a_F_WalSummarizerMain_25)
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	v3951 = *(*int32)(unsafe.Add(mBase, uint32(v42)+496))
	F_FileClose(m, v3951)
	mBase = m.M
	v3953 = m.ExcPending
	if v3953 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L616
	}
L616:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	v3958 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v3959 = m.ExcPending
	if v3959 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L617
	}
L617:
	;
	if v3958 != 0 {
		goto L618
	} else {
		goto L619
	}
L618:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	*(*int32)(unsafe.Add(mBase, uint32(v42-int32(-64)))) = v2868
	*(*int32)(unsafe.Add(mBase, uint32(v42)+60)) = v2872
	*(*int32)(unsafe.Add(mBase, uint32(v42)+56)) = v2874
	*(*int32)(unsafe.Add(mBase, uint32(v42)+52)) = v2878
	*(*int32)(unsafe.Add(mBase, uint32(v42)+48)) = v1332
	F_errmsg_internal(m, int32(_a_F_WalSummarizerMain_41), v42+int32(48))
	mBase = m.M
	v3973 = m.ExcPending
	if v3973 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L621
	}
L619:
	;
	goto L620
L620:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	v3988 = F_durable_rename(m, v42+int32(1536), v42+int32(512), int32(21))
	mBase = m.M
	v3989 = m.ExcPending
	if v3989 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L623
	}
L621:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	F_errfinish(m, int32(_a_F_WalSummarizerMain_1), int32(1330), int32(_a_F_WalSummarizerMain_24))
	mBase = m.M
	v3980 = m.ExcPending
	if v3980 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L622
	}
L622:
	;
	goto L620
L623:
	;
	goto L242
L624:
	;
	v4056 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v42)+299)) = uint8(v4056)
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	v4061 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[7]))
	v4065 = F_LWLockAcquire(m, v4061+int32(_a_F_WalSummarizerMain_4), int32(0))
	mBase = m.M
	v4066 = m.ExcPending
	if v4066 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L630
	}
L625:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	v4031 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v4032 = m.ExcPending
	if v4032 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L626
	}
L626:
	;
	if v4031 == int32(0) {
		goto L624
	} else {
		goto L627
	}
L627:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	*(*uint32)(unsafe.Add(mBase, uint32(v42)+16)) = uint32(v2838)
	v4038 = int64(32)
	v4039 = int64(base.Ui64(v2838) >> (uint(v4038) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v42)+12)) = uint32(v4039)
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = v1332
	*(*uint32)(unsafe.Add(mBase, uint32(v42)+8)) = uint32(v1874)
	v4044 = int64(base.Ui64(v1874) >> (uint(v4038) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v42)+4)) = uint32(v4044)
	F_errmsg_internal(m, int32(_a_F_WalSummarizerMain_42), v42)
	mBase = m.M
	v4048 = m.ExcPending
	if v4048 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L628
	}
L628:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	F_errfinish(m, int32(_a_F_WalSummarizerMain_1), int32(1342), int32(_a_F_WalSummarizerMain_24))
	mBase = m.M
	v4055 = m.ExcPending
	if v4055 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L629
	}
L629:
	;
	goto L624
L630:
	;
	v4068 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[8]))
	*(*int64)(unsafe.Add(mBase, uint32(v4068)+24)) = v2838
	v4070 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v4068)+16)) = uint8(v4070)
	*(*int32)(unsafe.Add(mBase, uint32(v4068)+4)) = v1332
	*(*int64)(unsafe.Add(mBase, uint32(v4068)+8)) = v2838
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	v4077 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[7]))
	F_LWLockRelease(m, v4077+int32(_a_F_WalSummarizerMain_4))
	mBase = m.M
	v4081 = m.ExcPending
	if v4081 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L631
	}
L631:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v42)+3152)) = v1872
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3164)) = v497
	v4085 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerMain[8]))
	F_ConditionVariableBroadcast(m, v4085+int32(32))
	mBase = m.M
	v4089 = m.ExcPending
	if v4089 != 0 {
		v4097 = v42
		v4114 = v59
		goto L6
	} else {
		goto L632
	}
L632:
	;
	v591 = v1320
	v604 = v1333
	v615 = v1872
	v617 = v2838
	v619 = v1351
	goto L123
L633:
	;
	goto L5
L634:
	;
	v4132 = int32(v4128)
	m.G0 = v4097
	v4134 = *(*int32)(unsafe.Add(mBase, uint32(v4132)+4))
	v4135 = *(*int32)(unsafe.Add(mBase, uint32(v4132)))
	v4138 = *(*int32)(unsafe.Add(mBase, uint32(v4135)))
	if v4097+int32(288) == v4138 {
		goto L637
	} else {
		goto L638
	}
L635:
	;
	m.ExcPending = 1
	goto L643
L636:
	;
	if v4142 != 0 {
		goto L640
	} else {
		goto L641
	}
L637:
	;
	v4140 = *(*int32)(unsafe.Add(mBase, uint32(v4135)+4))
	v4142 = v4140
	goto L639
L638:
	;
	v4142 = int32(0)
	goto L639
L639:
	;
	goto L636
L640:
	;
	v4143 = *(*int32)(unsafe.Add(mBase, uint32(v4097)+3164))
	v4144 = *(*int64)(unsafe.Add(mBase, uint32(v4097)+3152))
	v42 = v4097
	v43 = v4143
	v44 = v4142
	v49 = v4134
	v59 = v4114
	v65 = v4144
	goto L1
L641:
	;
	goto L642
L642:
	;
	F___wasm_longjmp(m, v4135, v4134)
	mBase = m.M
	v4146 = m.ExcPending
	if v4146 != 0 {
		goto L643
	} else {
		goto L644
	}
L643:
	;
	return
L644:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
