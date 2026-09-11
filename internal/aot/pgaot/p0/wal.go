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
	v2 = *(*int32)(unsafe.Add(mBase, _consts[3]))
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
	v3 = int32(*(*uint8)(unsafe.Add(mBase, _consts[179])))
	if v3 != 0 {
		v5 = int32(0)
		v8 = F_find_option(m, int32(314762), v5, v5, int32(21))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			v11 = int32(0)
			*(*uint8)(unsafe.Add(mBase, _consts[179])) = uint8(v11)
			v15 = *(*int32)(unsafe.Add(mBase, _consts[180]))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
			F_set_config_option_ext(m, int32(314762), v15, v16, v17, v18, v11, int32(21))
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
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	v2 = *(*int32)(unsafe.Add(mBase, _consts[492]))
	if v2 != 0 {
		F_ProcessProcSignalBarrier(m)
		mBase = m.M
		v4 = m.ExcPending
		if v4 != 0 {
			return
		} else {
			v6 = *(*int32)(unsafe.Add(mBase, _consts[305]))
			if v6 != 0 {
				*(*int32)(unsafe.Add(mBase, _consts[305])) = int32(0)
				F_ProcessConfigFile(m, int32(2))
				mBase = m.M
				v12 = m.ExcPending
				if v12 != 0 {
					return
				} else {
					v14 = *(*int32)(unsafe.Add(mBase, _consts[441]))
					if v14 == int32(0) {
						v18 = int32(*(*uint8)(unsafe.Add(mBase, _consts[493])))
						if v18 != 0 {
							v36 = *(*int32)(unsafe.Add(mBase, _consts[494]))
							if v36 != 0 {
								F_ProcessLogMemoryContextInterrupt(m)
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return
								} else {
									return
								}
							} else {
								return
							}
						} else {
							v21 = F_errstart(m, int32(14), int32(0))
							mBase = m.M
							v22 = m.ExcPending
							if v22 != 0 {
								return
							} else {
								if v21 != 0 {
									F_errmsg_internal(m, int32(229060), int32(0))
									mBase = m.M
									v26 = m.ExcPending
									if v26 != 0 {
										return
									} else {
										F_errfinish(m, int32(464783), int32(875), int32(108847))
										mBase = m.M
										v31 = m.ExcPending
										if v31 != 0 {
											return
										} else {
											F_proc_exit(m, int32(0))
											mBase = m.M
											v34 = m.ExcPending
											if v34 != 0 {
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
									v34 = m.ExcPending
									if v34 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v21 = F_errstart(m, int32(14), int32(0))
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return
						} else {
							if v21 != 0 {
								F_errmsg_internal(m, int32(229060), int32(0))
								mBase = m.M
								v26 = m.ExcPending
								if v26 != 0 {
									return
								} else {
									F_errfinish(m, int32(464783), int32(875), int32(108847))
									mBase = m.M
									v31 = m.ExcPending
									if v31 != 0 {
										return
									} else {
										F_proc_exit(m, int32(0))
										mBase = m.M
										v34 = m.ExcPending
										if v34 != 0 {
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
								v34 = m.ExcPending
								if v34 != 0 {
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
				v14 = *(*int32)(unsafe.Add(mBase, _consts[441]))
				if v14 == int32(0) {
					v18 = int32(*(*uint8)(unsafe.Add(mBase, _consts[493])))
					if v18 != 0 {
						v36 = *(*int32)(unsafe.Add(mBase, _consts[494]))
						if v36 != 0 {
							F_ProcessLogMemoryContextInterrupt(m)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return
							} else {
								return
							}
						} else {
							return
						}
					} else {
						v21 = F_errstart(m, int32(14), int32(0))
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return
						} else {
							if v21 != 0 {
								F_errmsg_internal(m, int32(229060), int32(0))
								mBase = m.M
								v26 = m.ExcPending
								if v26 != 0 {
									return
								} else {
									F_errfinish(m, int32(464783), int32(875), int32(108847))
									mBase = m.M
									v31 = m.ExcPending
									if v31 != 0 {
										return
									} else {
										F_proc_exit(m, int32(0))
										mBase = m.M
										v34 = m.ExcPending
										if v34 != 0 {
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
								v34 = m.ExcPending
								if v34 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v21 = F_errstart(m, int32(14), int32(0))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						if v21 != 0 {
							F_errmsg_internal(m, int32(229060), int32(0))
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return
							} else {
								F_errfinish(m, int32(464783), int32(875), int32(108847))
								mBase = m.M
								v31 = m.ExcPending
								if v31 != 0 {
									return
								} else {
									F_proc_exit(m, int32(0))
									mBase = m.M
									v34 = m.ExcPending
									if v34 != 0 {
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
							v34 = m.ExcPending
							if v34 != 0 {
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
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, _consts[305]))
		if v6 != 0 {
			*(*int32)(unsafe.Add(mBase, _consts[305])) = int32(0)
			F_ProcessConfigFile(m, int32(2))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, _consts[441]))
				if v14 == int32(0) {
					v18 = int32(*(*uint8)(unsafe.Add(mBase, _consts[493])))
					if v18 != 0 {
						v36 = *(*int32)(unsafe.Add(mBase, _consts[494]))
						if v36 != 0 {
							F_ProcessLogMemoryContextInterrupt(m)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return
							} else {
								return
							}
						} else {
							return
						}
					} else {
						v21 = F_errstart(m, int32(14), int32(0))
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return
						} else {
							if v21 != 0 {
								F_errmsg_internal(m, int32(229060), int32(0))
								mBase = m.M
								v26 = m.ExcPending
								if v26 != 0 {
									return
								} else {
									F_errfinish(m, int32(464783), int32(875), int32(108847))
									mBase = m.M
									v31 = m.ExcPending
									if v31 != 0 {
										return
									} else {
										F_proc_exit(m, int32(0))
										mBase = m.M
										v34 = m.ExcPending
										if v34 != 0 {
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
								v34 = m.ExcPending
								if v34 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v21 = F_errstart(m, int32(14), int32(0))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						if v21 != 0 {
							F_errmsg_internal(m, int32(229060), int32(0))
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return
							} else {
								F_errfinish(m, int32(464783), int32(875), int32(108847))
								mBase = m.M
								v31 = m.ExcPending
								if v31 != 0 {
									return
								} else {
									F_proc_exit(m, int32(0))
									mBase = m.M
									v34 = m.ExcPending
									if v34 != 0 {
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
							v34 = m.ExcPending
							if v34 != 0 {
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
			v14 = *(*int32)(unsafe.Add(mBase, _consts[441]))
			if v14 == int32(0) {
				v18 = int32(*(*uint8)(unsafe.Add(mBase, _consts[493])))
				if v18 != 0 {
					v36 = *(*int32)(unsafe.Add(mBase, _consts[494]))
					if v36 != 0 {
						F_ProcessLogMemoryContextInterrupt(m)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							return
						}
					} else {
						return
					}
				} else {
					v21 = F_errstart(m, int32(14), int32(0))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						if v21 != 0 {
							F_errmsg_internal(m, int32(229060), int32(0))
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return
							} else {
								F_errfinish(m, int32(464783), int32(875), int32(108847))
								mBase = m.M
								v31 = m.ExcPending
								if v31 != 0 {
									return
								} else {
									F_proc_exit(m, int32(0))
									mBase = m.M
									v34 = m.ExcPending
									if v34 != 0 {
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
							v34 = m.ExcPending
							if v34 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v21 = F_errstart(m, int32(14), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					if v21 != 0 {
						F_errmsg_internal(m, int32(229060), int32(0))
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return
						} else {
							F_errfinish(m, int32(464783), int32(875), int32(108847))
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return
							} else {
								F_proc_exit(m, int32(0))
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
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
						v34 = m.ExcPending
						if v34 != 0 {
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
	v5 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v5)+440)) = int32(1)
	if v6 != 0 {
		v10 = *(*int32)(unsafe.Add(mBase, _consts[3]))
		F_s_lock(m, v10+int32(440), int32(467473), int32(9567), int32(313602))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, _consts[3]))
			*(*int32)(unsafe.Add(mBase, uint32(v19)+440)) = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v19)+321)) = uint8(v1)
			return
		}
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, _consts[3]))
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
	v3 = *(*int32)(unsafe.Add(mBase, _consts[160]))
	v5 = F_LWLockAcquire(m, v3, int32(0))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, _consts[160]))
		F_LWLockUpdateVar(m, v8, v8+int32(16), int64(-1))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, _consts[160]))
			v19 = F_LWLockAcquire(m, v15+int32(128), int32(0))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, _consts[160]))
				F_LWLockUpdateVar(m, v22+int32(128), v22+int32(144), int64(-1))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, _consts[160]))
					v35 = F_LWLockAcquire(m, v31+int32(256), int32(0))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, _consts[160]))
						F_LWLockUpdateVar(m, v38+int32(256), v38+int32(272), int64(-1))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							v47 = *(*int32)(unsafe.Add(mBase, _consts[160]))
							v51 = F_LWLockAcquire(m, v47+int32(384), int32(0))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								v54 = *(*int32)(unsafe.Add(mBase, _consts[160]))
								F_LWLockUpdateVar(m, v54+int32(384), v54+int32(400), int64(-1))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return
								} else {
									v63 = *(*int32)(unsafe.Add(mBase, _consts[160]))
									v67 = F_LWLockAcquire(m, v63+int32(512), int32(0))
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return
									} else {
										v70 = *(*int32)(unsafe.Add(mBase, _consts[160]))
										F_LWLockUpdateVar(m, v70+int32(512), v70+int32(528), int64(-1))
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
											return
										} else {
											v79 = *(*int32)(unsafe.Add(mBase, _consts[160]))
											v83 = F_LWLockAcquire(m, v79+int32(640), int32(0))
											mBase = m.M
											v84 = m.ExcPending
											if v84 != 0 {
												return
											} else {
												v86 = *(*int32)(unsafe.Add(mBase, _consts[160]))
												F_LWLockUpdateVar(m, v86+int32(640), v86+int32(656), int64(-1))
												mBase = m.M
												v93 = m.ExcPending
												if v93 != 0 {
													return
												} else {
													v95 = *(*int32)(unsafe.Add(mBase, _consts[160]))
													v99 = F_LWLockAcquire(m, v95+int32(768), int32(0))
													mBase = m.M
													v100 = m.ExcPending
													if v100 != 0 {
														return
													} else {
														v102 = *(*int32)(unsafe.Add(mBase, _consts[160]))
														F_LWLockUpdateVar(m, v102+int32(768), v102+int32(784), int64(-1))
														mBase = m.M
														v109 = m.ExcPending
														if v109 != 0 {
															return
														} else {
															v111 = *(*int32)(unsafe.Add(mBase, _consts[160]))
															v115 = F_LWLockAcquire(m, v111+int32(896), int32(0))
															mBase = m.M
															v116 = m.ExcPending
															if v116 != 0 {
																return
															} else {
																v118 = int32(1)
																*(*uint8)(unsafe.Add(mBase, _consts[161])) = uint8(v118)
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
	v6 = *(*int32)(unsafe.Add(mBase, _consts[550]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+1456))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+1456)) = int32(1)
	v11 = v6 + int32(1456)
	if v7 != 0 {
		F_s_lock(m, v11, int32(464298), int32(82), int32(314190))
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
				v26 = F___time(m)
				mBase = m.M
				if v26-v24 < int64(11) {
					v53 = int32(1)
					return base.B2i32(v53 != int32(0))
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(1)
					if v30 != 0 {
						F_s_lock(m, v11, int32(464298), int32(103), int32(314190))
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
			v26 = F___time(m)
			mBase = m.M
			if v26-v24 < int64(11) {
				v53 = int32(1)
				return base.B2i32(v53 != int32(0))
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(1)
				if v30 != 0 {
					F_s_lock(m, v11, int32(464298), int32(103), int32(314190))
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
		v4 = *(*int32)(unsafe.Add(mBase, _consts[547]))
		F_ConditionVariableBroadcast(m, v4+int32(52))
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			if l1 != 0 {
				v10 = *(*int32)(unsafe.Add(mBase, _consts[547]))
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
			v10 = *(*int32)(unsafe.Add(mBase, _consts[547]))
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
