package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_ScanKeyEntryInitialize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int64) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	v3 = l2
	v4 = l3
	*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l4
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v4)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v3)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	v16 = l0 + int32(16)
	if l6 != 0 {
		F_fmgr_info(m, l6, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			return
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = int32(0)
		v21 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v21
		*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v21
		*(*int64)(unsafe.Add(mBase, uint32(v16))) = v21
		return
	}
}
func F_SetHintBitsExt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v21 int64
	_ = v21
	var v24 int64
	_ = v24
	var v31 int32
	_ = v31
	var v34 int64
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int64
	_ = v40
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	if l4 != 0 {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
		if v7 == int32(1) {
			return
		} else {
			if l3 == int32(0) {
				if l4 == int32(0) {
					v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
					F_BufferSetHintBits16(m, l0+int32(20), v48|l2, l1)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						return
					}
				} else {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
					if v52 == int32(0) {
						v55 = F_BufferBeginSetHintBits(m, l1)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							if v55 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(1)
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(2)
								v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
								v64 = v63 | l2
								*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)) = uint16(v64)
								return
							}
						}
					} else {
						v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
						v64 = v63 | l2
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)) = uint16(v64)
						return
					}
				}
			} else {
				if int32(0) <= l1 {
					v15 = *(*int32)(unsafe.Add(mBase, _c_F_SetHintBitsExt[0]))
					v21 = int64(0)
					v24 = base.AtomicRmwCmpxchg64(m, v15+l1*int32(56)-int32(32), int32(0), v21, v21)
					v31 = base.I32_wrap_i64(int64(base.Ui64(v24&int64(2147483648)) >> (uint(int64(31)) % 64)))
				} else {
					v31 = int32(0)
				}
				if v31 == int32(0) {
					if l4 == int32(0) {
						v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
						F_BufferSetHintBits16(m, l0+int32(20), v48|l2, l1)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							return
						}
					} else {
						v52 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
						if v52 == int32(0) {
							v55 = F_BufferBeginSetHintBits(m, l1)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return
							} else {
								if v55 == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(1)
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(2)
									v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
									v64 = v63 | l2
									*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)) = uint16(v64)
									return
								}
							}
						} else {
							v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
							v64 = v63 | l2
							*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)) = uint16(v64)
							return
						}
					}
				} else {
					v34 = F_TransactionIdGetCommitLSN(m, l3)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						v36 = F_XLogNeedsFlush(m, v34)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							if v36 == int32(0) {
								if l4 == int32(0) {
									v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
									F_BufferSetHintBits16(m, l0+int32(20), v48|l2, l1)
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return
									} else {
										return
									}
								} else {
									v52 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
									if v52 == int32(0) {
										v55 = F_BufferBeginSetHintBits(m, l1)
										mBase = m.M
										v56 = m.ExcPending
										if v56 != 0 {
											return
										} else {
											if v55 == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(1)
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(2)
												v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
												v64 = v63 | l2
												*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)) = uint16(v64)
												return
											}
										}
									} else {
										v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
										v64 = v63 | l2
										*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)) = uint16(v64)
										return
									}
								}
							} else {
								v40 = F_BufferGetLSNAtomic(m, l1)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return
								} else {
									if base.Ui64(v40) < base.Ui64(v34) {
										return
									} else {
										if l4 == int32(0) {
											v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
											F_BufferSetHintBits16(m, l0+int32(20), v48|l2, l1)
											mBase = m.M
											v51 = m.ExcPending
											if v51 != 0 {
												return
											} else {
												return
											}
										} else {
											v52 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
											if v52 == int32(0) {
												v55 = F_BufferBeginSetHintBits(m, l1)
												mBase = m.M
												v56 = m.ExcPending
												if v56 != 0 {
													return
												} else {
													if v55 == int32(0) {
														*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(1)
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(2)
														v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
														v64 = v63 | l2
														*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)) = uint16(v64)
														return
													}
												}
											} else {
												v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
												v64 = v63 | l2
												*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)) = uint16(v64)
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
	} else {
		if l3 == int32(0) {
			if l4 == int32(0) {
				v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
				F_BufferSetHintBits16(m, l0+int32(20), v48|l2, l1)
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return
				} else {
					return
				}
			} else {
				v52 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
				if v52 == int32(0) {
					v55 = F_BufferBeginSetHintBits(m, l1)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						if v55 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(1)
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(2)
							v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
							v64 = v63 | l2
							*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)) = uint16(v64)
							return
						}
					}
				} else {
					v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
					v64 = v63 | l2
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)) = uint16(v64)
					return
				}
			}
		} else {
			if int32(0) <= l1 {
				v15 = *(*int32)(unsafe.Add(mBase, _c_F_SetHintBitsExt[0]))
				v21 = int64(0)
				v24 = base.AtomicRmwCmpxchg64(m, v15+l1*int32(56)-int32(32), int32(0), v21, v21)
				v31 = base.I32_wrap_i64(int64(base.Ui64(v24&int64(2147483648)) >> (uint(int64(31)) % 64)))
			} else {
				v31 = int32(0)
			}
			if v31 == int32(0) {
				if l4 == int32(0) {
					v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
					F_BufferSetHintBits16(m, l0+int32(20), v48|l2, l1)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						return
					}
				} else {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
					if v52 == int32(0) {
						v55 = F_BufferBeginSetHintBits(m, l1)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							if v55 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(1)
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(2)
								v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
								v64 = v63 | l2
								*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)) = uint16(v64)
								return
							}
						}
					} else {
						v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
						v64 = v63 | l2
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)) = uint16(v64)
						return
					}
				}
			} else {
				v34 = F_TransactionIdGetCommitLSN(m, l3)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					v36 = F_XLogNeedsFlush(m, v34)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						if v36 == int32(0) {
							if l4 == int32(0) {
								v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
								F_BufferSetHintBits16(m, l0+int32(20), v48|l2, l1)
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return
								} else {
									return
								}
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
								if v52 == int32(0) {
									v55 = F_BufferBeginSetHintBits(m, l1)
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return
									} else {
										if v55 == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(1)
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(2)
											v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
											v64 = v63 | l2
											*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)) = uint16(v64)
											return
										}
									}
								} else {
									v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
									v64 = v63 | l2
									*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)) = uint16(v64)
									return
								}
							}
						} else {
							v40 = F_BufferGetLSNAtomic(m, l1)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return
							} else {
								if base.Ui64(v40) < base.Ui64(v34) {
									return
								} else {
									if l4 == int32(0) {
										v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
										F_BufferSetHintBits16(m, l0+int32(20), v48|l2, l1)
										mBase = m.M
										v51 = m.ExcPending
										if v51 != 0 {
											return
										} else {
											return
										}
									} else {
										v52 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
										if v52 == int32(0) {
											v55 = F_BufferBeginSetHintBits(m, l1)
											mBase = m.M
											v56 = m.ExcPending
											if v56 != 0 {
												return
											} else {
												if v55 == int32(0) {
													*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(1)
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(2)
													v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
													v64 = v63 | l2
													*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)) = uint16(v64)
													return
												}
											}
										} else {
											v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
											v64 = v63 | l2
											*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)) = uint16(v64)
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
func F_SetPlannerInfoExtensionState(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+388))
	if v6 == int32(0) {
		v10 = int32(1)
		v13 = l1 + v10
		if l1&v13 != 0 {
			v18 = v10 << (uint(int32(32)-base.I32_clz(v13)) % 32)
		} else {
			v18 = v13
		}
		if base.Ui32(v18) <= base.Ui32(int32(4)) {
			v21 = int32(4)
		} else {
			v21 = v18
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+392)) = v21
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+300))
		v26 = F_MemoryContextAllocZero(m, v23, v21<<(uint(int32(2))%32))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+388)) = v26
			v29 = v26
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+392))
			if v30 <= l1 {
				v33 = F_mul_size(m, int32(4), v30)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					v36 = int32(1)
					v39 = l1 + v36
					if l1&v39 != 0 {
						v44 = v36 << (uint(int32(32)-base.I32_clz(v39)) % 32)
					} else {
						v44 = v39
					}
					v45 = F_mul_size(m, int32(4), v44)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						v47 = F_repalloc0(m, v29, v33, v45)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+392)) = v44
							*(*int32)(unsafe.Add(mBase, uint32(l0)+388)) = v47
							v51 = v47
							*(*int32)(unsafe.Add(mBase, uint32(v51+l1<<(uint(int32(2))%32)))) = l2
							return
						}
					}
				}
			} else {
				v51 = v29
				*(*int32)(unsafe.Add(mBase, uint32(v51+l1<<(uint(int32(2))%32)))) = l2
				return
			}
		}
	} else {
		v29 = v6
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+392))
		if v30 <= l1 {
			v33 = F_mul_size(m, int32(4), v30)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return
			} else {
				v36 = int32(1)
				v39 = l1 + v36
				if l1&v39 != 0 {
					v44 = v36 << (uint(int32(32)-base.I32_clz(v39)) % 32)
				} else {
					v44 = v39
				}
				v45 = F_mul_size(m, int32(4), v44)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return
				} else {
					v47 = F_repalloc0(m, v29, v33, v45)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+392)) = v44
						*(*int32)(unsafe.Add(mBase, uint32(l0)+388)) = v47
						v51 = v47
						*(*int32)(unsafe.Add(mBase, uint32(v51+l1<<(uint(int32(2))%32)))) = l2
						return
					}
				}
			}
		} else {
			v51 = v29
			*(*int32)(unsafe.Add(mBase, uint32(v51+l1<<(uint(int32(2))%32)))) = l2
			return
		}
	}
}
func F_SignalHandlerForConfigReload(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	*(*int32)(unsafe.Add(mBase, _c_F_SignalHandlerForConfigReload[0])) = int32(1)
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_SignalHandlerForConfigReload[1]))
	v8 = int32(0)
	v11 = base.AtomicRmwOr32(m, v8, int32(_a_F_SignalHandlerForConfigReload_0), v8)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if v12 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(1)
	v15 = int32(0)
	v18 = base.AtomicRmwOr32(m, v15, int32(_a_F_SignalHandlerForConfigReload_0), v15)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v19 == v15 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	if v22 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_SignalHandlerForConfigReload[2]))
	if v26 == v22 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v28 = m.G0
	v30 = v28 - int32(16)
	m.G0 = v30
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_SignalHandlerForConfigReload[3]))
	if v33 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v56 = F_pgmem_kill(m, v22, int32(23))
	mBase = m.M
	goto L2
L9:
	;
	m.G0 = v30 + int32(16)
	goto L1
L10:
	;
	v36 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+15)) = uint8(v36)
	goto L11
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_SignalHandlerForConfigReload[4]))
	v44 = F_write(m, v40, v30+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v44 {
		goto L9
	} else {
		goto L13
	}
L12:
	;
	goto L9
L13:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_SignalHandlerForConfigReload[5]))
	if v48 == int32(27) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
}
func F_SplitIdentifierString(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v25 int32
	_ = v25
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v143 int32
	_ = v143
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	v15 = l0
	goto L1
L1:
	;
	v25 = int32(*(*int8)(unsafe.Add(mBase, uint32(v15))))
	goto L3
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v15
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v37 == int32(0) {
		v143 = int32(1)
		goto L5
	} else {
		goto L6
	}
L3:
	;
	if base.B2i32(v25 == int32(32))|base.B2i32(base.Ui32((v25-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v15 = v15 + int32(1)
		goto L1
	} else {
		goto L4
	}
L4:
	;
	goto L2
L5:
	;
	m.G0 = v11 + int32(16)
	return v143
L6:
	;
	v41 = l1 & int32(255)
	goto L7
L7:
	;
	v54 = F_scan_identifier(m, v11+int32(8), v11+int32(12), l1)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v143 = int32(1)
	goto L5
L9:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v123 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v122))) = uint8(v123)
	v125 = F_strlen(m, v54)
	mBase = m.M
	F_truncate_identifier(m, v54, v125, v123)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L10
	} else {
		goto L27
	}
L10:
	;
	return int32(0)
L11:
	;
	if v54 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v59 = v58
	goto L15
L13:
	;
	goto L14
L14:
	;
	v143 = int32(0)
	goto L5
L15:
	;
	v69 = int32(*(*int8)(unsafe.Add(mBase, uint32(v59))))
	goto L17
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v59
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if v41 == v80 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	if base.B2i32(v69 == int32(32))|base.B2i32(base.Ui32((v69-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v59 = v59 + int32(1)
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v85 = v59
	goto L22
L20:
	;
	goto L21
L21:
	;
	if v80 == int32(0) {
		goto L9
	} else {
		goto L26
	}
L22:
	;
	v90 = int32(*(*int8)(unsafe.Add(mBase, uint32(v85)+1)))
	v92 = v85 + int32(1)
	goto L24
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v92
	goto L9
L24:
	;
	if base.B2i32(v90 == int32(32))|base.B2i32(base.Ui32((v90-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v85 = v92
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	goto L14
L27:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v130 = F_lappend(m, v129, v54)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v130
	if v80 == v41 {
		goto L7
	} else {
		goto L29
	}
L29:
	;
	goto L8
}
func F_StatementTimeoutHandler(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_StatementTimeoutHandler[0]))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StatementTimeoutHandler[1])))
	if v9 != 0 {
		v10 = int32(15)
	} else {
		v10 = int32(2)
	}
	v11 = F_pgmem_kill(m, int32(0)-v4, v10)
	mBase = m.M
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_StatementTimeoutHandler[0]))
	v14 = F_pgmem_kill(m, v13, v10)
	mBase = m.M
	return
}
func F_StrategyCtlShmemInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	v2 = int32(_a_F_StrategyCtlShmemInit_0)
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_StrategyCtlShmemInit[0]))
	v4 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v3))), uint32(v4))
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_StrategyCtlShmemInit[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v4
	return
}
func F___sin(m *base.Module, l0 float64, l1 float64, l2 int32) float64 {
	var v7 float64
	_ = v7
	var v22 float64
	_ = v22
	var v23 float64
	_ = v23
	v7 = base.F64_mul(l0, l0)
	v22 = base.F64_add(base.F64_mul(base.F64_mul(v7, base.F64_mul(v7, v7)), base.F64_add(base.F64_mul(v7, float64(1.58969099521155e-10)), float64(-2.5050760253406863e-08))), base.F64_add(base.F64_mul(v7, base.F64_add(base.F64_mul(v7, float64(2.7557313707070068e-06)), float64(-0.0001984126982985795))), float64(0.00833333333332249)))
	v23 = base.F64_mul(l0, v7)
	if l2 == int32(0) {
		return base.F64_add(base.F64_mul(v23, base.F64_add(base.F64_mul(v7, v22), float64(-0.16666666666666632))), l0)
	} else {
		return base.F64_sub(l0, base.F64_add(base.F64_sub(base.F64_mul(v7, base.F64_sub(base.F64_mul(l1, float64(0.5)), base.F64_mul(v23, v22))), l1), base.F64_mul(v23, float64(0.16666666666666632))))
	}
}
func F___stdio_read(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l1
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = l2 - base.B2i32(v13 != v4)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v18
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v27 = m.Wasi_snapshot_preview1.Fd_read(m, v21, v10+int32(16), int32(2), v10+int32(12))
	mBase = m.M
	if v27 == v4 {
		v34 = int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F___stdio_read[0])) = v27
		v34 = int32(-1)
	}
	if v34 != 0 {
		v43 = int32(32)
		v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v43 | v44
		v65 = v4
	} else {
		v36 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
		if int32(0) < v36 {
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
			if base.Ui32(v36) <= base.Ui32(v47) {
				v65 = v36
			} else {
				v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v49
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v49 + (v36 - v47)
				v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				if v54 != 0 {
					v55 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v49 + v55
					v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
					*(*uint8)(unsafe.Add(mBase, uint32(l1+l2-v55))) = uint8(v61)
				} else {
				}
				v65 = l2
			}
		} else {
			if v36 != 0 {
				v41 = int32(32)
			} else {
				v41 = int32(16)
			}
			v43 = v41
			v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v43 | v44
			v65 = v4
		}
	}
	m.G0 = v10 + int32(32)
	return v65
}
func F___stdio_seek(m *base.Module, l0 int32, l1 int64, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v5 = F___lseek(m, v4, l1, l2)
	mBase = m.M
	return v5
}
func F_s_lock_stuck(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l0
		if l2 != 0 {
			v16 = l2
		} else {
			v16 = int32(_a_F_s_lock_stuck_0)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = v16
		F_errmsg_internal(m, int32(_a_F_s_lock_stuck_1), v7)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			F_errfinish(m, int32(_a_F_s_lock_stuck_2), int32(90), int32(_a_F_s_lock_stuck_3))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_save_ps_display_args(m *base.Module, l0 int32, l1 int32) int32 {
	return l1
}
func F_sbrk(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int64
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_sbrk[0]))
	v12 = base.I64_extend_i32_u(v5) + (base.I64_extend_i32_u(l0)+int64(7))&int64(8589934584)
	if base.Ui64(v12) <= base.Ui64(int64(4294967295)) {
		v15 = base.I32_wrap_i64(v12)
		if base.Ui32(v15) <= base.Ui32(base.MemorySize(m)<<(uint(int32(16))%32)) {
			*(*int32)(unsafe.Add(mBase, _c_F_sbrk[0])) = v15
			return v5
		} else {
			v20 = m.Env.Emscripten_resize_heap(m, v15)
			mBase = m.M
			if v20 != 0 {
				*(*int32)(unsafe.Add(mBase, _c_F_sbrk[0])) = v15
				return v5
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_sbrk[1])) = int32(48)
				return int32(-1)
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_sbrk[1])) = int32(48)
		return int32(-1)
	}
}
func F_scalarltjoinsel(m *base.Module, l0 int32) int64 {
	return int64(4599676419421066581)
}
func F_scalarltsel(m *base.Module, l0 int32) int64 {
	var v2 int32
	_ = v2
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v2 = int32(0)
	v4 = F_scalarineqsel_wrapper(m, l0, v2, v2)
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_searchstoplist(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l1
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v11 == v3 {
		v28 = v3
		m.G0 = v7 + int32(16)
		return v28
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v14 <= int32(0) {
			v28 = v3
			m.G0 = v7 + int32(16)
			return v28
		} else {
			v21 = F_bsearch(m, v7+int32(12), v11, v14, int32(4), int32(1285))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v28 = base.B2i32(v21 != int32(0))
				m.G0 = v7 + int32(16)
				return v28
			}
		}
	}
}
func F_send_feedback(m *base.Module, l0 int64, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v21 int64
	_ = v21
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v31 int64
	_ = v31
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int64
	_ = v43
	var v50 int64
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v60 int64
	_ = v60
	var v65 int32
	_ = v65
	var v71 int64
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int64
	_ = v81
	var v89 int64
	_ = v89
	var v93 int32
	_ = v93
	var v94 int64
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v103 int64
	_ = v103
	var v108 int64
	_ = v108
	var v109 int64
	_ = v109
	var v116 int64
	_ = v116
	var v119 int64
	_ = v119
	var v121 int64
	_ = v121
	var v123 int64
	_ = v123
	var v125 int64
	_ = v125
	var v127 int64
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int64
	_ = v136
	var v137 int64
	_ = v137
	var v145 int64
	_ = v145
	var v147 int64
	_ = v147
	var v150 int64
	_ = v150
	var v153 int64
	_ = v153
	var v155 int32
	_ = v155
	var v168 int32
	_ = v168
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
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int64
	_ = v213
	var v215 int64
	_ = v215
	var v217 int64
	_ = v217
	var v220 int64
	_ = v220
	var v222 int64
	_ = v222
	var v224 int64
	_ = v224
	var v226 int64
	_ = v226
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int64
	_ = v260
	var v262 int64
	_ = v262
	var v264 int64
	_ = v264
	var v267 int64
	_ = v267
	var v269 int64
	_ = v269
	var v271 int64
	_ = v271
	var v273 int64
	_ = v273
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int64
	_ = v307
	var v309 int64
	_ = v309
	var v311 int64
	_ = v311
	var v314 int64
	_ = v314
	var v316 int64
	_ = v316
	var v318 int64
	_ = v318
	var v320 int64
	_ = v320
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int64
	_ = v354
	var v356 int64
	_ = v356
	var v358 int64
	_ = v358
	var v361 int64
	_ = v361
	var v363 int64
	_ = v363
	var v365 int64
	_ = v365
	var v367 int64
	_ = v367
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int64
	_ = v410
	var v411 int64
	_ = v411
	var v416 int64
	_ = v416
	var v420 int64
	_ = v420
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v442 int64
	_ = v442
	var v447 int64
	_ = v447
	var v452 int64
	_ = v452
	v3 = l2
	v7 = int64(0)
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	if l1 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(32)
	return
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_send_feedback[0]))
	if v17 <= int32(0) {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v21 = *(*int64)(unsafe.Add(mBase, _c_F_send_feedback[1]))
	if base.Ui64(v21) < base.Ui64(l0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L4
L6:
	;
	v23 = l0
	goto L8
L7:
	;
	v23 = v21
	goto L8
L8:
	;
	v24 = int32(0)
	v26 = int32(_a_F_send_feedback_0)
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_send_feedback[2]))
	v28 = int64(0)
	v31 = base.AtomicRmwCmpxchg64(m, v27, int32(272), v28, v28)
	*(*int64)(unsafe.Add(mBase, _c_F_send_feedback[3])) = v31
	v36 = base.AtomicRmwOr32(m, v24, int32(_a_F_send_feedback_1), v24)
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_send_feedback[2]))
	v43 = base.AtomicRmwCmpxchg64(m, v39, int32(264), v28, v28)
	*(*int64)(unsafe.Add(mBase, _c_F_send_feedback[4])) = v43
	goto L11
L9:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_send_feedback[5]))
	v53 = int32(0)
	if base.B2i32(v52 == v53)|base.B2i32(v52 == int32(_a_F_send_feedback_2)) == v53 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	goto L12
L12:
	;
	v50 = *(*int64)(unsafe.Add(mBase, _c_F_send_feedback[3]))
	goto L9
L13:
	;
	v121 = *(*int64)(unsafe.Add(mBase, _c_F_send_feedback[6]))
	if base.Ui64(v121) < base.Ui64(v116) {
		goto L33
	} else {
		goto L34
	}
L14:
	;
	v60 = *(*int64)(unsafe.Add(mBase, uint32(v52)+8))
	if base.Ui64(v60) <= base.Ui64(v50) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v100 = v52
	v103 = v7
	goto L16
L16:
	;
	if v100 != int32(_a_F_send_feedback_2) {
		goto L27
	} else {
		goto L28
	}
L17:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _c_F_send_feedback[5]))
	v100 = v96
	v103 = v71
	goto L16
L18:
	;
	v65 = v52
	goto L21
L19:
	;
	v89 = v7
	goto L20
L20:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_send_feedback[7]))
	v94 = *(*int64)(unsafe.Add(mBase, uint32(v93)+16))
	v116 = v89
	v119 = v94
	goto L13
L21:
	;
	v71 = *(*int64)(unsafe.Add(mBase, uint32(v65)+16))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v72)+4)) = v73
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = v75
	F_pfree(m, v65)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v89 = v71
	goto L20
L23:
	;
	return
L24:
	;
	if v73 == int32(_a_F_send_feedback_2) {
		goto L17
	} else {
		goto L25
	}
L25:
	;
	v81 = *(*int64)(unsafe.Add(mBase, uint32(v73)+8))
	if base.Ui64(v81) <= base.Ui64(v50) {
		v65 = v73
		goto L21
	} else {
		goto L26
	}
L26:
	;
	goto L22
L27:
	;
	v108 = v103
	goto L29
L28:
	;
	v108 = v23
	goto L29
L29:
	;
	if v100 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v109 = v108
	goto L32
L31:
	;
	v109 = v23
	goto L32
L32:
	;
	v116 = v109
	v119 = v109
	goto L13
L33:
	;
	v123 = v116
	goto L35
L34:
	;
	v123 = v121
	goto L35
L35:
	;
	v125 = *(*int64)(unsafe.Add(mBase, _c_F_send_feedback[8]))
	if base.Ui64(v125) < base.Ui64(v119) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v127 = v119
	goto L38
L37:
	;
	v127 = v125
	goto L38
L38:
	;
	v131 = m.G0
	v132 = int32(16)
	v133 = v131 - v132
	m.G0 = v133
	F_gettimeofday(m, v133)
	mBase = m.M
	v136 = *(*int64)(unsafe.Add(mBase, uint32(v133)))
	v137 = int64(*(*int32)(unsafe.Add(mBase, uint32(v133)+8)))
	m.G0 = v133 + v132
	v145 = v137 + v136*int64(1000000) - int64(946684800000000)
	goto L39
L39:
	;
	if l1 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_send_feedback[9])) = v145
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_send_feedback[10]))
	if v168 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L41:
	;
	v147 = *(*int64)(unsafe.Add(mBase, _c_F_send_feedback[8]))
	if v127 != v147 {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v150 = *(*int64)(unsafe.Add(mBase, _c_F_send_feedback[6]))
	if v123 != v150 {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v153 = *(*int64)(unsafe.Add(mBase, _c_F_send_feedback[9]))
	v155 = *(*int32)(unsafe.Add(mBase, _c_F_send_feedback[0]))
	goto L44
L44:
	;
	if base.B2i32(base.I64_extend_i32_s(v155*int32(1000))*int64(1000) <= v145-v153) == int32(0) {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	goto L40
L46:
	;
	F_enlargeStringInfo(m, v192, int32(1))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L23
	} else {
		goto L52
	}
L47:
	;
	v171 = int32(_a_F_send_feedback_3)
	v172 = *(*int32)(unsafe.Add(mBase, _c_F_send_feedback[11]))
	v175 = *(*int32)(unsafe.Add(mBase, _c_F_send_feedback[12]))
	*(*int32)(unsafe.Add(mBase, _c_F_send_feedback[11])) = v175
	v177 = F_makeStringInfo(m)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L23
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	v184 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v183))) = uint8(v184)
	*(*int32)(unsafe.Add(mBase, uint32(v168)+12)) = v184
	*(*int32)(unsafe.Add(mBase, uint32(v168)+4)) = v184
	goto L51
L50:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_send_feedback[11])) = v172
	*(*int32)(unsafe.Add(mBase, _c_F_send_feedback[10])) = v177
	v192 = v177
	goto L46
L51:
	;
	v191 = *(*int32)(unsafe.Add(mBase, _c_F_send_feedback[10]))
	v192 = v191
	goto L46
L52:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	v200 = int32(114)
	*(*uint8)(unsafe.Add(mBase, uint32(v197+v198))) = uint8(v200)
	*(*int32)(unsafe.Add(mBase, uint32(v192)+4)) = v197 + int32(1)
	v206 = *(*int32)(unsafe.Add(mBase, _c_F_send_feedback[10]))
	F_enlargeStringInfo(m, v206, int32(8))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L23
	} else {
		goto L53
	}
L53:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v206)+4))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	v213 = int64(56)
	v215 = int64(65280)
	v217 = int64(40)
	v220 = int64(16711680)
	v222 = int64(24)
	v224 = int64(4278190080)
	v226 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v210+v211))) = v23<<(uint(v213)%64) | v23&v215<<(uint(v217)%64) | (v23&v220<<(uint(v222)%64) | v23&v224<<(uint(v226)%64)) | (int64(base.Ui64(v23)>>(uint(v226)%64))&v224 | int64(base.Ui64(v23)>>(uint(v222)%64))&v220 | (int64(base.Ui64(v23)>>(uint(v217)%64))&v215 | int64(base.Ui64(v23)>>(uint(v213)%64))))
	v249 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v206)+4)) = v210 + v249
	v253 = *(*int32)(unsafe.Add(mBase, _c_F_send_feedback[10]))
	F_enlargeStringInfo(m, v253, v249)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L23
	} else {
		goto L54
	}
L54:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v253)+4))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v253)))
	v260 = int64(56)
	v262 = int64(65280)
	v264 = int64(40)
	v267 = int64(16711680)
	v269 = int64(24)
	v271 = int64(4278190080)
	v273 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v257+v258))) = v123<<(uint(v260)%64) | v123&v262<<(uint(v264)%64) | (v123&v267<<(uint(v269)%64) | v123&v271<<(uint(v273)%64)) | (int64(base.Ui64(v123)>>(uint(v273)%64))&v271 | int64(base.Ui64(v123)>>(uint(v269)%64))&v267 | (int64(base.Ui64(v123)>>(uint(v264)%64))&v262 | int64(base.Ui64(v123)>>(uint(v260)%64))))
	v296 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v253)+4)) = v257 + v296
	v300 = *(*int32)(unsafe.Add(mBase, _c_F_send_feedback[10]))
	F_enlargeStringInfo(m, v300, v296)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L23
	} else {
		goto L55
	}
L55:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v300)+4))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v300)))
	v307 = int64(56)
	v309 = int64(65280)
	v311 = int64(40)
	v314 = int64(16711680)
	v316 = int64(24)
	v318 = int64(4278190080)
	v320 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v304+v305))) = v127<<(uint(v307)%64) | v127&v309<<(uint(v311)%64) | (v127&v314<<(uint(v316)%64) | v127&v318<<(uint(v320)%64)) | (int64(base.Ui64(v127)>>(uint(v320)%64))&v318 | int64(base.Ui64(v127)>>(uint(v316)%64))&v314 | (int64(base.Ui64(v127)>>(uint(v311)%64))&v309 | int64(base.Ui64(v127)>>(uint(v307)%64))))
	v343 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v300)+4)) = v304 + v343
	v347 = *(*int32)(unsafe.Add(mBase, _c_F_send_feedback[10]))
	F_enlargeStringInfo(m, v347, v343)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L23
	} else {
		goto L56
	}
L56:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v347)+4))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v347)))
	v354 = int64(56)
	v356 = int64(65280)
	v358 = int64(40)
	v361 = int64(16711680)
	v363 = int64(24)
	v365 = int64(4278190080)
	v367 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v351+v352))) = v145<<(uint(v354)%64) | v145&v356<<(uint(v358)%64) | (v145&v361<<(uint(v363)%64) | v145&v365<<(uint(v367)%64)) | (int64(base.Ui64(v145)>>(uint(v367)%64))&v365 | int64(base.Ui64(v145)>>(uint(v363)%64))&v361 | (int64(base.Ui64(v145)>>(uint(v358)%64))&v356 | int64(base.Ui64(v145)>>(uint(v354)%64))))
	*(*int32)(unsafe.Add(mBase, uint32(v347)+4)) = v351 + int32(8)
	v394 = *(*int32)(unsafe.Add(mBase, _c_F_send_feedback[10]))
	F_enlargeStringInfo(m, v394, int32(1))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L23
	} else {
		goto L57
	}
L57:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v394)+4))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v394)))
	*(*uint8)(unsafe.Add(mBase, uint32(v398+v399))) = uint8(v3)
	*(*int32)(unsafe.Add(mBase, uint32(v394)+4)) = v398 + int32(1)
	v407 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L23
	} else {
		goto L58
	}
L58:
	;
	if v407 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+24)) = uint32(v123)
	v410 = int64(32)
	v411 = int64(base.Ui64(v123) >> (uint(v410) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+20)) = uint32(v411)
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+16)) = uint32(v127)
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
	v416 = int64(base.Ui64(v127) >> (uint(v410) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+12)) = uint32(v416)
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+8)) = uint32(v23)
	v420 = int64(base.Ui64(v23) >> (uint(v410) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+4)) = uint32(v420)
	F_errmsg_internal(m, int32(_a_F_send_feedback_4), v12)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L23
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v431 = *(*int32)(unsafe.Add(mBase, _c_F_send_feedback[13]))
	v433 = *(*int32)(unsafe.Add(mBase, _c_F_send_feedback[10]))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v433)))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v433)+4))
	v437 = *(*int32)(unsafe.Add(mBase, _c_F_send_feedback[14]))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v437)+44))
	m.T0[v438].(func(*base.Module, int32, int32, int32))(m, v431, v434, v435)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L23
	} else {
		goto L64
	}
L62:
	;
	F_errfinish(m, int32(_a_F_send_feedback_5), int32(_a_F_send_feedback_6), int32(_a_F_send_feedback_7))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L23
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v442 = *(*int64)(unsafe.Add(mBase, _c_F_send_feedback[1]))
	if base.Ui64(v442) < base.Ui64(v23) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_send_feedback[1])) = v23
	goto L67
L66:
	;
	goto L67
L67:
	;
	v447 = *(*int64)(unsafe.Add(mBase, _c_F_send_feedback[8]))
	if base.Ui64(v447) < base.Ui64(v127) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_send_feedback[8])) = v127
	goto L70
L69:
	;
	goto L70
L70:
	;
	v452 = *(*int64)(unsafe.Add(mBase, _c_F_send_feedback[6]))
	if base.Ui64(v123) <= base.Ui64(v452) {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_send_feedback[6])) = v123
	goto L1
}
func F_set_deparse_for_query(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
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
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v489 int32
	_ = v489
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v689 int32
	_ = v689
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v735 int32
	_ = v735
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v760 int32
	_ = v760
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v783 int32
	_ = v783
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v808 int32
	_ = v808
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v831 int32
	_ = v831
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
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v870 int32
	_ = v870
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v882 int32
	_ = v882
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v895 int32
	_ = v895
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v918 int32
	_ = v918
	var v920 int32
	_ = v920
	var v922 int32
	_ = v922
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v930 int32
	_ = v930
	var v933 int32
	_ = v933
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v955 int32
	_ = v955
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	v4 = int32(0)
	base.MemoryFill(m, l0, v4, int32(80))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v22
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v26
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+92))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v32
	F_set_rtable_names(m, l0, l2, v4)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v46 = v4
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v46
	if v46 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	if v69 != 0 {
		goto L16
	} else {
		goto L17
	}
L5:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v58 = v56
	goto L7
L6:
	;
	v58 = int32(0)
	goto L7
L7:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v59 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v62 = v60
	goto L10
L9:
	;
	v62 = int32(0)
	goto L10
L10:
	;
	if v58 < v62 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v65 = F_palloc0(m, int32(52))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	goto L4
L14:
	;
	v67 = F_lappend(m, v46, v65)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v46 = v67
	goto L3
L16:
	;
	v70 = F_has_dangerous_join_using(m, l0, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	v79 = v46
	v80 = v59
	goto L18
L18:
	;
	v93 = v4
	goto L21
L19:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)) = uint8(v70)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	F_set_using_names(m, l0, v73, int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v79 = v78
	v80 = v77
	goto L18
L21:
	;
	v99 = int32(0)
	if v80 == v99 {
		v109 = v99
		goto L23
	} else {
		goto L24
	}
L23:
	;
	if v79 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	if v103 <= v93 {
		v109 = int32(0)
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	v109 = v105 + v93<<(uint(int32(2))%32)
	goto L23
L26:
	;
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v136)+8))
	if int32(0) < v700 {
		goto L158
	} else {
		goto L159
	}
L27:
	;
	return
L28:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	if base.B2i32(v109 == int32(0))|base.B2i32(v114 <= v93) != 0 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	if v117 == int32(0) {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v120 = int32(2)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v117+v93<<(uint(v120)%32))))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	if v125 == v120 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v123)+32))
	v131 = int32(2)
	v134 = int32(4)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v129+v130<<(uint(v131)%32)-v134)))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v123)+28))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v129+v137<<(uint(v131)%32)-v134)))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v124)+8))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+8))
	if v145 != 0 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	F_set_relation_column_names(m, l0, v124, v123)
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L1
	} else {
		goto L157
	}
L34:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+4))
	v148 = v146
	goto L36
L35:
	;
	v148 = int32(0)
	goto L36
L36:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	if v149 < v148 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	if v151 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	goto L39
L39:
	;
	F_build_colinfo_names_hash(m, v123)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L48
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v123))) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v123)+4)) = v165
	goto L39
L41:
	;
	v155 = F_palloc0_mul(m, int32(4), v148)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v158 = F_mul_size(m, int32(4), v149)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L45
	}
L44:
	;
	v165 = v155
	goto L40
L45:
	;
	v161 = F_mul_size(m, int32(4), v148)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v163 = F_repalloc0(m, v151, v158, v161)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v165 = v163
	goto L40
L48:
	;
	v172 = int32(0)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v123)+44))
	if v174 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)+4))
	v176 = v175
	goto L51
L50:
	;
	v176 = v172
	goto L51
L51:
	;
	if v176 < v148 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v180 = v176
	v185 = v172
	goto L55
L53:
	;
	v315 = v174
	v318 = v172
	goto L54
L54:
	;
	v329 = int32(0)
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v136)+8))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v143)+8))
	if v315 != 0 {
		goto L93
	} else {
		goto L94
	}
L55:
	;
	v197 = v180 << (uint(int32(2)) % 32)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v123)+36))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v197+v198)))
	if int32(0) < v200 {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v123)+44))
	v315 = v310
	v318 = v306
	goto L54
L57:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	v230 = v229 + v197
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
	if v231 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L58:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	v228 = v203 + v200<<(uint(int32(2))%32) - int32(4)
	goto L57
L59:
	;
	goto L60
L60:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v123)+40))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v209+v197)))
	if int32(0) < v211 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	v228 = v214 + v211<<(uint(int32(2))%32) - int32(4)
	goto L57
L62:
	;
	goto L63
L63:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v124)+8))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+8))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+12))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v222+v197)))
	v228 = v224 + int32(4)
	goto L57
L64:
	;
	v308 = v180 + int32(1)
	if v308 != v148 {
		v180 = v308
		v185 = v306
		goto L55
	} else {
		goto L92
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v230))) = int32(0)
	v306 = v185
	goto L64
L66:
	;
	goto L67
L67:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	if v236 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v230))) = v231
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v123)+48))
	if v240 == int32(0) {
		v306 = v185
		goto L64
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
	if v247 != 0 {
		v271 = v247
		goto L73
	} else {
		goto L74
	}
L71:
	;
	v245 = F_hash_search(m, v240, v231, int32(1), int32(0))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v306 = v185
	goto L64
L73:
	;
	v272 = int32(1)
	if v185&v272 != 0 {
		v306 = v272
		goto L64
	} else {
		goto L84
	}
L74:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v236)+8))
	if v248 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v248)+4))
	v251 = v249
	goto L77
L76:
	;
	v251 = int32(0)
	goto L77
L77:
	;
	if v180 < v251 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v248)+12))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v253+v197)))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v255)+4))
	v257 = v256
	goto L80
L79:
	;
	v257 = v231
	goto L80
L80:
	;
	v258 = F_make_colname_unique(m, v257, l0, v123)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v260+v197))) = v258
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v123)+48))
	if v263 == int32(0) {
		v271 = v258
		goto L73
	} else {
		goto L82
	}
L82:
	;
	v268 = F_hash_search(m, v263, v258, int32(1), int32(0))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v271 = v258
	goto L73
L84:
	;
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271))))
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231))))
	if base.B2i32(v277 == int32(0))|base.B2i32(v277 != v280) != 0 {
		v298 = v277
		v299 = v280
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v306 = base.B2i32(v298-v299 != int32(0))
	goto L64
L86:
	;
	goto L85
L87:
	;
	v283 = v271
	v284 = v231
	goto L88
L88:
	;
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284)+1)))
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+1)))
	if v288 == int32(0) {
		v298 = v288
		v299 = v287
		goto L86
	} else {
		goto L90
	}
L89:
	;
	v298 = v288
	v299 = v287
	goto L86
L90:
	;
	v291 = int32(1)
	if v288 == v287 {
		v283 = v283 + v291
		v284 = v284 + v291
		goto L88
	} else {
		goto L91
	}
L91:
	;
	goto L89
L92:
	;
	goto L56
L93:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v315)+4))
	v335 = v333
	goto L95
L94:
	;
	v335 = int32(0)
	goto L95
L95:
	;
	v336 = v330 + v331 - v335
	*(*int32)(unsafe.Add(mBase, uint32(v123)+8)) = v336
	v340 = F_palloc0(m, v336<<(uint(int32(2))%32))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v123)+12)) = v340
	v343 = F_palloc0(m, v336)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v123)+16)) = v343
	if v148 <= int32(0) {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v143)+8))
	if v429 <= int32(0) {
		goto L115
	} else {
		goto L116
	}
L99:
	;
	v348 = int32(0)
	v412 = v348
	v422 = v348
	v427 = v329
	goto L98
L100:
	;
	goto L101
L101:
	;
	v350 = int32(0)
	v353 = v350
	v363 = v350
	v368 = v329
	goto L102
L102:
	;
	v371 = v353 << (uint(int32(2)) % 32)
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v123)+36))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v371+v372)))
	if v374 == int32(0) {
		v412 = v353
		v422 = v363
		v427 = v368
		goto L98
	} else {
		goto L104
	}
L103:
	;
	v412 = v148
	v422 = v399
	v427 = v407
	goto L98
L104:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v123)+40))
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v377+v371)))
	if v379 == int32(0) {
		v412 = v353
		v422 = v363
		v427 = v368
		goto L98
	} else {
		goto L105
	}
L105:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v384+v371)))
	*(*int32)(unsafe.Add(mBase, uint32(v382+v371))) = v386
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v123)+16))
	v390 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v388+v353))) = uint8(v390)
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v123)+36))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v392+v371)))
	if v390 < v394 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v397 = F_bms_add_member(m, v363, v394)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L109
	}
L107:
	;
	v399 = v363
	goto L108
L108:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v123)+40))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v400+v371)))
	if int32(0) < v402 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	v399 = v397
	goto L108
L110:
	;
	v405 = F_bms_add_member(m, v368, v402)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L113
	}
L111:
	;
	v407 = v368
	goto L112
L112:
	;
	v409 = v353 + int32(1)
	if v409 != v148 {
		v353 = v409
		v363 = v399
		v368 = v407
		goto L102
	} else {
		goto L114
	}
L113:
	;
	v407 = v405
	goto L112
L114:
	;
	goto L103
L115:
	;
	v683 = v412
	v684 = v412
	v689 = v318
	goto L26
L116:
	;
	goto L117
L117:
	;
	v432 = int32(0)
	v435 = v412
	v436 = v412
	v438 = v432
	v440 = v432
	v441 = v318
	goto L118
L118:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v143)+16))
	v454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v452+v440))))
	if v454 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L119:
	;
	v683 = v656
	v684 = v657
	v689 = v662
	goto L26
L120:
	;
	v674 = v440 + int32(1)
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v143)+8))
	if v674 < v675 {
		v435 = v656
		v436 = v657
		v438 = v659
		v440 = v674
		v441 = v662
		goto L118
	} else {
		goto L156
	}
L121:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v123)+16))
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v143)+16))
	v651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v649+v440))))
	*(*uint8)(unsafe.Add(mBase, uint32(v647+v435))) = uint8(v651)
	v656 = v435 + int32(1)
	v657 = v631
	v659 = v633
	v662 = v636
	goto L120
L122:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	if v457 <= v438 {
		v489 = v438
		goto L125
	} else {
		goto L126
	}
L123:
	;
	goto L124
L124:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v143)+12))
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v564+v440<<(uint(int32(2))%32))))
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	if v569 != 0 {
		goto L140
	} else {
		goto L141
	}
L125:
	;
	v504 = v489 + int32(1)
	v505 = F_bms_is_member(m, v504, v422)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L1
	} else {
		goto L131
	}
L126:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	v464 = v438
	goto L127
L127:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v459+v464<<(uint(int32(2))%32))))
	if v481 != 0 {
		v489 = v464
		goto L125
	} else {
		goto L129
	}
L128:
	;
	v489 = v457
	goto L125
L129:
	;
	v483 = v464 + int32(1)
	if v483 != v457 {
		v464 = v483
		goto L127
	} else {
		goto L130
	}
L130:
	;
	goto L128
L131:
	;
	if v505 != 0 {
		v656 = v435
		v657 = v436
		v659 = v504
		v662 = v441
		goto L120
	} else {
		goto L132
	}
L132:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	if v508 <= v436 {
		v537 = v436
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	v554 = int32(2)
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v507+v537<<(uint(v554)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v553+v435<<(uint(v554)%32)))) = v560
	v631 = v537 + int32(1)
	v633 = v504
	v636 = v441
	goto L121
L134:
	;
	v512 = v436
	goto L135
L135:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v507+v512<<(uint(int32(2))%32))))
	if v531 != 0 {
		v537 = v512
		goto L133
	} else {
		goto L137
	}
L136:
	;
	v537 = v508
	goto L133
L137:
	;
	v533 = v512 + int32(1)
	if v533 != v508 {
		v512 = v533
		goto L135
	} else {
		goto L138
	}
L138:
	;
	goto L136
L139:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v123)+48))
	if v619 != 0 {
		goto L152
	} else {
		goto L153
	}
L140:
	;
	v570 = F_make_colname_unique(m, v568, l0, v123)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L1
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v611+v435<<(uint(int32(2))%32)))) = v568
	v618 = v441
	goto L139
L143:
	;
	v573 = v435 << (uint(int32(2)) % 32)
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v573+v574))) = v570
	v577 = int32(1)
	if v441&v577 != 0 {
		v618 = v577
		goto L139
	} else {
		goto L144
	}
L144:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v580+v573)))
	v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v582))))
	v588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v568))))
	if base.B2i32(v585 == int32(0))|base.B2i32(v585 != v588) != 0 {
		v606 = v585
		v607 = v588
		goto L146
	} else {
		goto L147
	}
L145:
	;
	v618 = base.B2i32(v606-v607 != int32(0))
	goto L139
L146:
	;
	goto L145
L147:
	;
	v591 = v582
	v592 = v568
	goto L148
L148:
	;
	v595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v592)+1)))
	v596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v591)+1)))
	if v596 == int32(0) {
		v606 = v596
		v607 = v595
		goto L146
	} else {
		goto L150
	}
L149:
	;
	v606 = v596
	v607 = v595
	goto L146
L150:
	;
	v599 = int32(1)
	if v596 == v595 {
		v591 = v591 + v599
		v592 = v592 + v599
		goto L148
	} else {
		goto L151
	}
L151:
	;
	goto L149
L152:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v620+v435<<(uint(int32(2))%32))))
	v627 = F_hash_search(m, v619, v624, int32(1), int32(0))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L1
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	v631 = v436
	v633 = v438
	v636 = v618
	goto L121
L155:
	;
	goto L154
L156:
	;
	goto L119
L157:
	;
	v93 = v93 + int32(1)
	goto L21
L158:
	;
	v703 = int32(0)
	v706 = v683
	v707 = v684
	v709 = v703
	v712 = v689
	v713 = v703
	goto L161
L159:
	;
	v955 = v689
	goto L160
L160:
	;
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v123)+48))
	if v966 != 0 {
		goto L200
	} else {
		goto L201
	}
L161:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v136)+16))
	v725 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v723+v713))))
	if v725 == int32(0) {
		goto L165
	} else {
		goto L166
	}
L162:
	;
	v955 = v933
	goto L160
L163:
	;
	v945 = v713 + int32(1)
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v136)+8))
	if v945 < v946 {
		v706 = v927
		v707 = v928
		v709 = v930
		v712 = v933
		v713 = v945
		goto L161
	} else {
		goto L199
	}
L164:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v123)+16))
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v136)+16))
	v922 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v920+v713))))
	*(*uint8)(unsafe.Add(mBase, uint32(v918+v706))) = uint8(v922)
	v927 = v706 + int32(1)
	v928 = v902
	v930 = v904
	v933 = v907
	goto L163
L165:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	if v728 <= v709 {
		v760 = v709
		goto L168
	} else {
		goto L169
	}
L166:
	;
	goto L167
L167:
	;
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v136)+12))
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v835+v713<<(uint(int32(2))%32))))
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	if v840 != 0 {
		goto L183
	} else {
		goto L184
	}
L168:
	;
	v775 = v760 + int32(1)
	v776 = F_bms_is_member(m, v775, v427)
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L1
	} else {
		goto L174
	}
L169:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	v735 = v709
	goto L170
L170:
	;
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v730+v735<<(uint(int32(2))%32))))
	if v752 != 0 {
		v760 = v735
		goto L168
	} else {
		goto L172
	}
L171:
	;
	v760 = v728
	goto L168
L172:
	;
	v754 = v735 + int32(1)
	if v754 != v728 {
		v735 = v754
		goto L170
	} else {
		goto L173
	}
L173:
	;
	goto L171
L174:
	;
	if v776 != 0 {
		v927 = v706
		v928 = v707
		v930 = v775
		v933 = v712
		goto L163
	} else {
		goto L175
	}
L175:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	if v779 <= v707 {
		v808 = v707
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	v825 = int32(2)
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v778+v808<<(uint(v825)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v824+v706<<(uint(v825)%32)))) = v831
	v902 = v808 + int32(1)
	v904 = v775
	v907 = v712
	goto L164
L177:
	;
	v783 = v707
	goto L178
L178:
	;
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v778+v783<<(uint(int32(2))%32))))
	if v802 != 0 {
		v808 = v783
		goto L176
	} else {
		goto L180
	}
L179:
	;
	v808 = v779
	goto L176
L180:
	;
	v804 = v783 + int32(1)
	if v804 != v779 {
		v783 = v804
		goto L178
	} else {
		goto L181
	}
L181:
	;
	goto L179
L182:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v123)+48))
	if v890 != 0 {
		goto L195
	} else {
		goto L196
	}
L183:
	;
	v841 = F_make_colname_unique(m, v839, l0, v123)
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L1
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v882+v706<<(uint(int32(2))%32)))) = v839
	v889 = v712
	goto L182
L186:
	;
	v844 = v706 << (uint(int32(2)) % 32)
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v844+v845))) = v841
	v848 = int32(1)
	if v712&v848 != 0 {
		v889 = v848
		goto L182
	} else {
		goto L187
	}
L187:
	;
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v851+v844)))
	v856 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v853))))
	v859 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v839))))
	if base.B2i32(v856 == int32(0))|base.B2i32(v856 != v859) != 0 {
		v877 = v856
		v878 = v859
		goto L189
	} else {
		goto L190
	}
L188:
	;
	v889 = base.B2i32(v877-v878 != int32(0))
	goto L182
L189:
	;
	goto L188
L190:
	;
	v862 = v853
	v863 = v839
	goto L191
L191:
	;
	v866 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v863)+1)))
	v867 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v862)+1)))
	if v867 == int32(0) {
		v877 = v867
		v878 = v866
		goto L189
	} else {
		goto L193
	}
L192:
	;
	v877 = v867
	v878 = v866
	goto L189
L193:
	;
	v870 = int32(1)
	if v867 == v866 {
		v862 = v862 + v870
		v863 = v863 + v870
		goto L191
	} else {
		goto L194
	}
L194:
	;
	goto L192
L195:
	;
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v891+v706<<(uint(int32(2))%32))))
	v898 = F_hash_search(m, v890, v895, int32(1), int32(0))
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L1
	} else {
		goto L198
	}
L196:
	;
	goto L197
L197:
	;
	v902 = v707
	v904 = v709
	v907 = v889
	goto L164
L198:
	;
	goto L197
L199:
	;
	goto L162
L200:
	;
	F_hash_destroy(m, v966)
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L1
	} else {
		goto L203
	}
L201:
	;
	goto L202
L202:
	;
	v971 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	v974 = base.B2i32(v971 != int32(0)) & v955
	*(*uint8)(unsafe.Add(mBase, uint32(v123)+20)) = uint8(v974)
	v93 = v93 + int32(1)
	goto L21
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v123)+48)) = int32(0)
	goto L202
}
func F_set_pathtarget_cost_width(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int64
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int64
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 float64
	_ = v61
	var v62 float64
	_ = v62
	var v63 float64
	_ = v63
	var v66 float64
	_ = v66
	var v70 int64
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int64
	_ = v75
	var v78 int64
	_ = v78
	var v90 int32
	_ = v90
	v3 = int32(0)
	v8 = int64(0)
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v8
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = v8
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v20 == v3 {
		v90 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v90
	m.G0 = v13 + int32(32)
	return l1
L2:
	;
	v23 = int32(0)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v24 <= v23 {
		v90 = v23
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v28 = v13 + int32(16)
	v34 = v3
	v36 = v8
	goto L4
L4:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39+v34<<(uint(int32(2))%32))))
	v44 = F_get_expr_width(m, l0, v43)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v75 = int64(1073741823)
	if v75 <= v70 {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	return int32(0)
L7:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	if v49 != int32(6) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = l0
	v53 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v28))) = v53
	v59 = F_cost_qual_eval_walker(m, v43, v13+int32(8))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v70 = v36 + base.I64_extend_i32_s(v44)
	v72 = v34 + int32(1)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v72 < v73 {
		v34 = v72
		v36 = v70
		goto L4
	} else {
		goto L12
	}
L11:
	;
	v61 = *(*float64)(unsafe.Add(mBase, uint32(v13)+24))
	v62 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
	v63 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+16)) = base.F64_add(v62, v63)
	v66 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+24)) = base.F64_add(v61, v66)
	goto L10
L12:
	;
	goto L5
L13:
	;
	v78 = v75
	goto L15
L14:
	;
	v78 = v70
	goto L15
L15:
	;
	v90 = base.I32_wrap_i64(v78)
	goto L1
}
func F_set_rtable_names(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int64
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v236 int32
	_ = v236
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v299 int32
	_ = v299
	v4 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(80)
	m.G0 = v15
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v4
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+40)) = int64(292057776192)
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_set_rtable_names[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v23
	v26 = int64(*(*int32)(unsafe.Add(mBase, uint32(v19)+4)))
	v30 = F_hash_create(m, int32(_a_F_set_rtable_names_0), v26, v15+int32(32), int32(1048))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v15 + int32(80)
	return
L4:
	;
	return
L5:
	;
	if l1 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v117 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v34 <= int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v43 = v4
	goto L9
L9:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49+v43<<(uint(int32(2))%32))))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v54 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L6
L11:
	;
	v102 = v43 + int32(1)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v102 < v103 {
		v43 = v102
		goto L9
	} else {
		goto L21
	}
L12:
	;
	v57 = int32(0)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v58 <= v57 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v66 = v57
	goto L14
L14:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v54)+12))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v73+v66<<(uint(int32(2))%32))))
	if v77 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L11
L16:
	;
	v81 = F_hash_search(m, v30, v77, int32(1), v15+int32(31))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L4
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v86 = v66 + int32(1)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v86 < v87 {
		v66 = v86
		goto L14
	} else {
		goto L20
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v81)+64)) = int32(0)
	goto L18
L20:
	;
	goto L15
L21:
	;
	goto L10
L22:
	;
	F_hash_destroy(m, v30)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L4
	} else {
		goto L72
	}
L23:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	if v120 <= int32(0) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v133 = int32(1)
	v134 = v4
	goto L25
L25:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v117)+12))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v136+v134<<(uint(int32(2))%32))))
	v142 = *(*int32)(unsafe.Add(mBase, _c_F_set_rtable_names[1]))
	if v142 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L22
L27:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L4
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if l2 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L29
L31:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v277 = F_lappend(m, v276, v265)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L4
	} else {
		goto L70
	}
L32:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v140)+4))
	if v150 != 0 {
		goto L37
	} else {
		goto L38
	}
L33:
	;
	v147 = F_bms_is_member(m, v133, l2)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	if v147 != 0 {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v265 = int32(0)
	goto L31
L36:
	;
	if v160 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L37:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+4))
	v160 = v151
	goto L36
L38:
	;
	goto L39
L39:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v140)+12))
	switch v153 {
	case 0:
		goto L41
	default:
		goto L40
	case 2:
		v265 = int32(0)
		goto L31
	}
L40:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v140)+8))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	v160 = v158
	goto L36
L41:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
	v155 = F_get_rel_name(m, v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	v160 = v155
	goto L36
L43:
	;
	v265 = int32(0)
	goto L31
L44:
	;
	goto L45
L45:
	;
	v167 = F_hash_search(m, v30, v160, int32(1), v15+int32(31))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+31)))
	if v169 == int32(1) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v172 = F_strlen(m, v160)
	mBase = m.M
	v175 = F_palloc(m, v172+int32(16))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L4
	} else {
		goto L50
	}
L48:
	;
	v250 = v160
	v261 = v167
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v261)+64)) = int32(0)
	v265 = v250
	goto L31
L50:
	;
	v182 = v172
	goto L51
L51:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v167)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v167)+64)) = v189 + int32(1)
	if v182 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v250 = v175
	v261 = v246
	goto L49
L53:
	;
	base.MemoryCopy(m, v175, v160, v182)
	goto L55
L54:
	;
	goto L55
L55:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v167)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v194
	v200 = F_pg_sprintf(m, v175+v182, int32(_a_F_set_rtable_names_1), v15+int32(16))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	v202 = F_strlen(m, v175)
	mBase = m.M
	if base.Ui32(int32(64)) <= base.Ui32(v202) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v210 = v182
	goto L60
L58:
	;
	v236 = v182
	goto L59
L59:
	;
	v246 = F_hash_search(m, v30, v175, int32(1), v15+int32(31))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L4
	} else {
		goto L68
	}
L60:
	;
	v219 = F_pg_mbcliplen(m, v160, v210, v210-int32(1))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L4
	} else {
		goto L62
	}
L61:
	;
	v236 = v219
	goto L59
L62:
	;
	if v219 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	base.MemoryCopy(m, v175, v160, v219)
	goto L65
L64:
	;
	goto L65
L65:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v167)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v222
	v226 = F_pg_sprintf(m, v175+v219, int32(_a_F_set_rtable_names_1), v15)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L4
	} else {
		goto L66
	}
L66:
	;
	v228 = F_strlen(m, v175)
	mBase = m.M
	if base.Ui32(int32(63)) < base.Ui32(v228) {
		v210 = v219
		goto L60
	} else {
		goto L67
	}
L67:
	;
	goto L61
L68:
	;
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+31)))
	if v248 != 0 {
		v182 = v236
		goto L51
	} else {
		goto L69
	}
L69:
	;
	goto L52
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v277
	v280 = int32(1)
	v283 = v134 + v280
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	if v283 < v284 {
		v133 = v133 + v280
		v134 = v283
		goto L25
	} else {
		goto L71
	}
L71:
	;
	goto L26
L72:
	;
	goto L3
}
func F_set_using_names(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
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
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
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
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v325 int32
	_ = v325
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v644 int32
	_ = v644
	var v649 int32
	_ = v649
	v19 = m.G0
	v21 = v19 - int32(48)
	m.G0 = v21
	v24 = l1
	v25 = l2
	goto L1
L1:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v41 != int32(64) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L13
	} else {
		goto L147
	}
L3:
	;
	goto L2
L4:
	;
	switch v41 - int32(63) {
	case 0:
		goto L7
	default:
		goto L3
	case 2:
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	v106 = int32(4)
	v107 = v103<<(uint(int32(2))%32) - v106
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+12))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v107+v109)))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v113+v107)))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	switch v119 - int32(63) {
	case 0:
		v140 = v106
		goto L16
	case 1:
		goto L17
	default:
		goto L18
	}
L7:
	;
	m.G0 = v21 + int32(48)
	return
L8:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v46 == int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v49 <= int32(0) {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v57 = int32(0)
	goto L11
L11:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v71+v57<<(uint(int32(2))%32))))
	F_set_using_names(m, l0, v75, v25)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L7
L13:
	;
	return
L14:
	;
	v79 = v57 + int32(1)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v79 < v80 {
		v57 = v79
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v140+v118)))
	*(*int32)(unsafe.Add(mBase, uint32(v111)+28)) = v142
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	switch v145 - int32(63) {
	case 0:
		v166 = v106
		goto L22
	case 1:
		goto L23
	default:
		goto L24
	}
L17:
	;
	v140 = int32(36)
	goto L16
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v127
	F_errmsg_internal(m, int32(_a_F_set_using_names_0), v21+int32(16))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L13
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(_a_F_set_using_names_1), int32(_a_F_set_using_names_2), int32(_a_F_set_using_names_3))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L22:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v144+v166)))
	*(*int32)(unsafe.Add(mBase, uint32(v111)+32)) = v168
	v170 = int32(0)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v115)+52))
	if v172 != 0 {
		goto L28
	} else {
		goto L29
	}
L23:
	;
	v166 = int32(36)
	goto L22
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L13
	} else {
		goto L25
	}
L25:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v153
	F_errmsg_internal(m, int32(_a_F_set_using_names_0), v21+int32(32))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L13
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_set_using_names_1), int32(_a_F_set_using_names_4), int32(_a_F_set_using_names_3))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L13
	} else {
		goto L27
	}
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L28:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	v176 = v173 << (uint(int32(2)) % 32)
	goto L30
L29:
	;
	v176 = v170
	goto L30
L30:
	;
	v177 = F_palloc0(m, v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L13
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111)+36)) = v177
	v180 = F_palloc0(m, v176)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L13
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111)+40)) = v180
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v115)+56))
	if v183 == int32(0) {
		v223 = v170
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v115)+60))
	if v237 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L34:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v183)+4))
	if v186 <= int32(0) {
		v223 = v170
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v193 = v170
	goto L36
L36:
	;
	v208 = v193 << (uint(int32(2)) % 32)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v111)+36))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v183)+12))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v211+v208)))
	*(*int32)(unsafe.Add(mBase, uint32(v208+v209))) = v213
	v216 = v193 + int32(1)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v183)+4))
	if v216 < v217 {
		v193 = v216
		goto L36
	} else {
		goto L38
	}
L37:
	;
	v223 = v216
	goto L33
L38:
	;
	goto L37
L39:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v298)+12))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v111)+32))
	v301 = int32(2)
	v304 = int32(4)
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v299+v300<<(uint(v301)%32)-v304)))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v111)+28))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v299+v307<<(uint(v301)%32)-v304)))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v111)+40))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v111)+36))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v316 != 0 {
		goto L48
	} else {
		goto L49
	}
L40:
	;
	v240 = int32(0)
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v237)+4))
	if v241 <= v240 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v247 = v240
	v248 = v223
	goto L42
L42:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v111)+40))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v115)+48))
	v264 = base.B2i32(v263 <= v247)
	if v263 <= v247 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L39
L44:
	;
	v265 = v248
	goto L46
L45:
	;
	v265 = v247
	goto L46
L46:
	;
	v266 = int32(2)
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v237)+12))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v269+v247<<(uint(v266)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v262+v265<<(uint(v266)%32)))) = v273
	v277 = v247 + int32(1)
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v237)+4))
	if v277 < v278 {
		v247 = v277
		v248 = v248 + v264
		goto L42
	} else {
		goto L47
	}
L47:
	;
	goto L43
L48:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	if v439 == int32(0) {
		v614 = v25
		goto L82
	} else {
		goto L83
	}
L49:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	if v317 <= int32(0) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v325 = int32(0)
	goto L51
L51:
	;
	v340 = v325 << (uint(int32(2)) % 32)
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v340+v341)))
	if v343 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	goto L48
L53:
	;
	v418 = v325 + int32(1)
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	if v418 < v419 {
		v325 = v418
		goto L51
	} else {
		goto L81
	}
L54:
	;
	v346 = v340 + v315
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v346)))
	if int32(0) < v347 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v313)+4))
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v313)))
	if v351 < v347 {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	goto L57
L57:
	;
	v380 = v340 + v314
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v380)))
	if v381 <= int32(0) {
		goto L53
	} else {
		goto L69
	}
L58:
	;
	if v350 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	v370 = v350
	v371 = v347
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v371<<(uint(int32(2))%32)+v370-int32(4)))) = v343
	goto L57
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v313))) = v347
	*(*int32)(unsafe.Add(mBase, uint32(v313)+4)) = v366
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v346)))
	v370 = v366
	v371 = v369
	goto L60
L62:
	;
	v356 = F_palloc0_mul(m, int32(4), v347)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L13
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v359 = F_mul_size(m, int32(4), v351)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L13
	} else {
		goto L66
	}
L65:
	;
	v366 = v356
	goto L61
L66:
	;
	v362 = F_mul_size(m, int32(4), v347)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L13
	} else {
		goto L67
	}
L67:
	;
	v364 = F_repalloc0(m, v350, v359, v362)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L13
	} else {
		goto L68
	}
L68:
	;
	v366 = v364
	goto L61
L69:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v306)+4))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v306)))
	if v385 < v381 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	if v384 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L71:
	;
	v404 = v384
	v405 = v381
	goto L72
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v405<<(uint(int32(2))%32)+v404-int32(4)))) = v343
	goto L53
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v306))) = v381
	*(*int32)(unsafe.Add(mBase, uint32(v306)+4)) = v400
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v380)))
	v404 = v400
	v405 = v403
	goto L72
L74:
	;
	v390 = F_palloc0_mul(m, int32(4), v381)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L13
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v393 = F_mul_size(m, int32(4), v385)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L13
	} else {
		goto L78
	}
L77:
	;
	v400 = v390
	goto L73
L78:
	;
	v396 = F_mul_size(m, int32(4), v381)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L13
	} else {
		goto L79
	}
L79:
	;
	v398 = F_repalloc0(m, v384, v393, v396)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L13
	} else {
		goto L80
	}
L80:
	;
	v400 = v398
	goto L73
L81:
	;
	goto L52
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v313)+24)) = v614
	*(*int32)(unsafe.Add(mBase, uint32(v306)+24)) = v614
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	F_set_using_names(m, l0, v632, v614)
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L13
	} else {
		goto L146
	}
L83:
	;
	v442 = F_list_copy(m, v25)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L13
	} else {
		goto L84
	}
L84:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	if v444 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v444)+4))
	v447 = v445
	goto L87
L86:
	;
	v447 = int32(0)
	goto L87
L87:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	if v448 < v447 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	if v450 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L89:
	;
	v470 = v444
	goto L90
L90:
	;
	if v470 == int32(0) {
		v614 = v442
		goto L82
	} else {
		goto L99
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111))) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v111)+4)) = v464
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v470 = v467
	goto L90
L92:
	;
	v454 = F_palloc0_mul(m, int32(4), v447)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L13
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v457 = F_mul_size(m, int32(4), v448)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L13
	} else {
		goto L96
	}
L95:
	;
	v464 = v454
	goto L91
L96:
	;
	v460 = F_mul_size(m, int32(4), v447)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L13
	} else {
		goto L97
	}
L97:
	;
	v462 = F_repalloc0(m, v450, v457, v460)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L13
	} else {
		goto L98
	}
L98:
	;
	v464 = v462
	goto L91
L99:
	;
	v473 = int32(0)
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v470)+4))
	if v474 <= v473 {
		v614 = v442
		goto L82
	} else {
		goto L100
	}
L100:
	;
	v479 = v442
	v482 = v473
	goto L101
L101:
	;
	v496 = v482 << (uint(int32(2)) % 32)
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v496+v497)))
	if v499 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v614 = v538
	goto L82
L103:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v470)+12))
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v502+v496)))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v504)+4))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v506 == int32(0) {
		v519 = v505
		goto L106
	} else {
		goto L107
	}
L104:
	;
	v533 = v499
	goto L105
L105:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v111)+44))
	v535 = F_lappend(m, v534, v533)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L13
	} else {
		goto L115
	}
L106:
	;
	v520 = F_make_colname_unique(m, v519, l0, v111)
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L13
	} else {
		goto L110
	}
L107:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v506)+8))
	if v509 == int32(0) {
		v519 = v505
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v509)+4))
	if v512 <= v482 {
		v519 = v505
		goto L106
	} else {
		goto L109
	}
L109:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v509)+12))
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v514+v496)))
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v516)+4))
	v519 = v517
	goto L106
L110:
	;
	v522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v522 == int32(1) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v526 = F_lappend(m, v525, v520)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L13
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v529+v496))) = v520
	v533 = v520
	goto L105
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v526
	goto L113
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111)+44)) = v535
	v538 = F_lappend(m, v479, v533)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L13
	} else {
		goto L116
	}
L116:
	;
	v540 = v496 + v315
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v540)))
	if int32(0) < v541 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v313)+4))
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v313)))
	if v545 < v541 {
		goto L120
	} else {
		goto L121
	}
L118:
	;
	goto L119
L119:
	;
	v574 = v496 + v314
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v574)))
	if int32(0) < v575 {
		goto L131
	} else {
		goto L132
	}
L120:
	;
	if v544 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L121:
	;
	v564 = v544
	v565 = v541
	goto L122
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v565<<(uint(int32(2))%32)+v564-int32(4)))) = v533
	goto L119
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v313))) = v541
	*(*int32)(unsafe.Add(mBase, uint32(v313)+4)) = v560
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v540)))
	v564 = v560
	v565 = v563
	goto L122
L124:
	;
	v550 = F_palloc0_mul(m, int32(4), v541)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L13
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	v553 = F_mul_size(m, int32(4), v545)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L13
	} else {
		goto L128
	}
L127:
	;
	v560 = v550
	goto L123
L128:
	;
	v556 = F_mul_size(m, int32(4), v541)
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L13
	} else {
		goto L129
	}
L129:
	;
	v558 = F_repalloc0(m, v544, v553, v556)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L13
	} else {
		goto L130
	}
L130:
	;
	v560 = v558
	goto L123
L131:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v306)+4))
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v306)))
	if v579 < v575 {
		goto L134
	} else {
		goto L135
	}
L132:
	;
	goto L133
L133:
	;
	v609 = v482 + int32(1)
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v470)+4))
	if v609 < v610 {
		v479 = v538
		v482 = v609
		goto L101
	} else {
		goto L145
	}
L134:
	;
	if v578 == int32(0) {
		goto L138
	} else {
		goto L139
	}
L135:
	;
	v598 = v578
	v599 = v575
	goto L136
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v599<<(uint(int32(2))%32)+v598-int32(4)))) = v533
	goto L133
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v306))) = v575
	*(*int32)(unsafe.Add(mBase, uint32(v306)+4)) = v594
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v574)))
	v598 = v594
	v599 = v597
	goto L136
L138:
	;
	v584 = F_palloc0_mul(m, int32(4), v575)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L13
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	v587 = F_mul_size(m, int32(4), v579)
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L13
	} else {
		goto L142
	}
L141:
	;
	v594 = v584
	goto L137
L142:
	;
	v590 = F_mul_size(m, int32(4), v575)
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L13
	} else {
		goto L143
	}
L143:
	;
	v592 = F_repalloc0(m, v578, v587, v590)
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L13
	} else {
		goto L144
	}
L144:
	;
	v594 = v592
	goto L137
L145:
	;
	goto L102
L146:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v24 = v635
	v25 = v614
	goto L1
L147:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v640
	F_errmsg_internal(m, int32(_a_F_set_using_names_5), v21)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L13
	} else {
		goto L148
	}
L148:
	;
	F_errfinish(m, int32(_a_F_set_using_names_1), int32(_a_F_set_using_names_6), int32(_a_F_set_using_names_7))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L13
	} else {
		goto L149
	}
L149:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_setval_oid(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	F_SetSequence(m, v3, v4, int32(1))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_shdepDropDependency(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v30 int32
	_ = v30
	var v43 int32
	_ = v43
	var v60 int32
	_ = v60
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v96 int64
	_ = v96
	var v97 int64
	_ = v97
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v173 int32
	_ = v173
	v12 = m.G0
	v14 = v12 - int32(224)
	m.G0 = v14
	v18 = int32(1)
	if l1 <= int32(3591) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	v90 = int32(3)
	v96 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_shdepDropDependency[0])))
	if v89 != 0 {
		goto L26
	} else {
		goto L27
	}
L2:
	;
	goto L1
L3:
	;
	v89 = int32(0)
	goto L2
L4:
	;
	if base.B2i32(base.Ui32(l1-int32(2964)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(l1-int32(2846)) < base.Ui32(int32(2))) != 0 {
		v89 = v18
		goto L2
	} else {
		goto L25
	}
L5:
	;
	if l1 <= int32(2670) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	if l1 <= int32(_a_F_shdepDropDependency_0) {
		goto L15
	} else {
		goto L16
	}
L8:
	;
	switch l1 - int32(1213) {
	case 0, 1, 19, 20, 47, 48, 49:
		v89 = v18
		goto L2
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
		goto L3
	default:
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v30 = l1 - int32(2671)
	if base.B2i32(base.Ui32(int32(27)) < base.Ui32(v30))|base.B2i32(int32(1)<<(uint(v30)%32)&int32(226492515) == int32(0)) != 0 {
		goto L4
	} else {
		goto L13
	}
L11:
	;
	if base.Ui32(int32(2)) <= base.Ui32(l1-int32(2396)) {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v89 = v18
	goto L2
L13:
	;
	v89 = v18
	goto L2
L14:
	;
	if base.Ui32(l1-int32(3592)) < base.Ui32(int32(2)) {
		v89 = v18
		goto L2
	} else {
		goto L23
	}
L15:
	;
	v43 = l1 - int32(_a_F_shdepDropDependency_1)
	if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v43))|base.B2i32(int32(1)<<(uint(v43)%32)&int32(963) == int32(0)) != 0 {
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	switch l1 - int32(_a_F_shdepDropDependency_2) {
	case 0, 1, 2, 3, 4, 59, 60:
		v89 = v18
		goto L2
	case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
		goto L3
	default:
		goto L19
	}
L18:
	;
	v89 = v18
	goto L2
L19:
	;
	if base.Ui32(l1-int32(_a_F_shdepDropDependency_3)) < base.Ui32(int32(3)) {
		v89 = v18
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v60 = l1 - int32(_a_F_shdepDropDependency_4)
	if base.Ui32(int32(15)) < base.Ui32(v60) {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	if int32(1)<<(uint(v60)%32)&int32(_a_F_shdepDropDependency_5) != 0 {
		v89 = v18
		goto L2
	} else {
		goto L22
	}
L22:
	;
	goto L3
L23:
	;
	if base.Ui32(int32(2)) <= base.Ui32(l1-int32(4060)) {
		goto L3
	} else {
		goto L24
	}
L24:
	;
	v89 = v18
	goto L2
L25:
	;
	goto L3
L26:
	;
	v97 = int64(0)
	goto L28
L27:
	;
	v97 = v96
	goto L28
L28:
	;
	F_ScanKeyInit(m, v14, int32(1), v90, int32(184), v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	return
L30:
	;
	F_ScanKeyInit(m, v14+int32(56), int32(2), int32(3), int32(184), base.I64_extend_i32_u(l1))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v110 = int32(3)
	F_ScanKeyInit(m, v14+int32(112), v110, v110, int32(184), base.I64_extend_i32_u(l2))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	if l4 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v118 = int32(4)
	F_ScanKeyInit(m, v14+int32(168), v118, int32(3), int32(65), base.I64_extend_i32_s(l3))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L29
	} else {
		goto L36
	}
L34:
	;
	v127 = v90
	goto L35
L35:
	;
	v131 = F_systable_beginscan(m, l0, int32(1232), int32(1), int32(0), v127, v14)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L29
	} else {
		goto L37
	}
L36:
	;
	v127 = v118
	goto L35
L37:
	;
	v133 = F_systable_getnext(m, v131)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L29
	} else {
		goto L38
	}
L38:
	;
	if v133 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v136 = v133
	goto L42
L40:
	;
	goto L41
L41:
	;
	F_systable_endscan(m, v131)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L29
	} else {
		goto L60
	}
L42:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v136)+16))
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+22)))
	v148 = v146 + v147
	if l5 != 0 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L41
L44:
	;
	v159 = F_systable_getnext(m, v131)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L29
	} else {
		goto L58
	}
L45:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+16))
	if v149 != l5 {
		goto L44
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	if l6 != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L47
L49:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v148)+20))
	if v151 != l6 {
		goto L44
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	if l7 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	goto L51
L53:
	;
	v153 = int32(*(*int8)(unsafe.Add(mBase, uint32(v148)+24)))
	if l7 != v153 {
		goto L44
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	F_simple_heap_delete(m, l0, v136+int32(4))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L29
	} else {
		goto L57
	}
L56:
	;
	goto L55
L57:
	;
	goto L44
L58:
	;
	if v159 != 0 {
		v136 = v159
		goto L42
	} else {
		goto L59
	}
L59:
	;
	goto L43
L60:
	;
	m.G0 = v14 + int32(224)
	return
}
func F_shell_archive_init(m *base.Module) int32 {
	return int32(_a_F_shell_archive_init_0)
}
func F_shell_archive_shutdown(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	v4 = F_errstart(m, int32(14), int32(0))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		if v4 != 0 {
			F_errmsg_internal(m, int32(_a_F_shell_archive_shutdown_0), int32(0))
			v9 = m.ExcPending
			if v9 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_shell_archive_shutdown_1), int32(142), int32(_a_F_shell_archive_shutdown_2))
				v14 = m.ExcPending
				if v14 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			return
		}
	}
}
func F_shift_jis_2004_to_euc_jis_2004(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v112 int32
	_ = v112
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v139 int32
	_ = v139
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v168 int32
	_ = v168
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v283 int32
	_ = v283
	var v295 int32
	_ = v295
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+104))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	F_check_encoding_conversion_args(m, v15, v16, v17, int32(41), int32(5))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	if v17 <= int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	F_report_invalid_encoding(m, int32(41), v28, v26)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L109
	}
L4:
	;
	v283 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v276))) = uint8(v283)
	return base.I64_extend_i32_s(v274 - v14)
L5:
	;
	v274 = v14
	v276 = v13
	goto L4
L6:
	;
	goto L7
L7:
	;
	v26 = v17
	v28 = v14
	v30 = v13
	goto L8
L8:
	;
	v37 = int32(*(*int8)(unsafe.Add(mBase, uint32(v28))))
	if int32(0) <= v37 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v274 = v263
	v276 = v265
	goto L4
L10:
	;
	if int32(0) < v269 {
		v26 = v269
		v28 = v263
		v30 = v265
		goto L8
	} else {
		goto L108
	}
L11:
	;
	if v37 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v52 = F_pg_encoding_verifymbchar(m, int32(41), v28, v26)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L18
	}
L14:
	;
	if v12 != int64(0) {
		v274 = v28
		v276 = v30
		goto L4
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v30))) = uint8(v37)
	v45 = int32(1)
	v263 = v28 + v45
	v265 = v30 + v45
	v269 = v26 - v45
	goto L10
L17:
	;
	goto L3
L18:
	;
	if base.Ui32(v26) < base.Ui32(v52) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if v12 != int64(0) {
		v274 = v28
		v276 = v30
		goto L4
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if base.B2i32(v52 != int32(1))|base.B2i32(base.Ui32(int32(62)) < base.Ui32((v37+int32(95))&int32(255))) == int32(0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	goto L3
L23:
	;
	v263 = v28 + v52
	v265 = v257
	v269 = v26 - v52
	goto L10
L24:
	;
	v257 = v250 + int32(2)
	goto L23
L25:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+1)) = uint8(v37)
	v69 = int32(142)
	*(*uint8)(unsafe.Add(mBase, uint32(v30))) = uint8(v69)
	v250 = v30
	goto L24
L26:
	;
	goto L27
L27:
	;
	if v52 != int32(2) {
		v257 = v30
		goto L23
	} else {
		goto L28
	}
L28:
	;
	v74 = v37 & int32(255)
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
	v76 = base.I32_extend8_s(v75)
	if base.B2i32(v37 == int32(-128))|base.B2i32(base.Ui32(int32(-97)) < base.Ui32(v37)) == int32(0) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v242 = int32(96)
	v243 = v240 - v242
	*(*uint8)(unsafe.Add(mBase, uint32(v241)+1)) = uint8(v243)
	v246 = v239 - v242
	*(*uint8)(unsafe.Add(mBase, uint32(v241))) = uint8(v246)
	v250 = v241
	goto L24
L30:
	;
	v239 = v74<<(uint(int32(1))%32) + v232 - int32(256)
	v240 = v233
	v241 = v30
	goto L29
L31:
	;
	v85 = v75 + int32(-64)
	if base.Ui32(v85) <= base.Ui32(int32(62)) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	if v37&int32(-16) == int32(-32) {
		goto L45
	} else {
		goto L46
	}
L34:
	;
	v232 = int32(-1)
	v233 = v75 - int32(63)
	goto L30
L35:
	;
	goto L36
L36:
	;
	if v76 < int32(-97) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	if v12 != int64(0) {
		v274 = v28
		v276 = v30
		goto L4
	} else {
		goto L43
	}
L38:
	;
	v99 = v85
	v100 = int32(-1)
	goto L40
L39:
	;
	if int32(-4) < v76 {
		goto L37
	} else {
		goto L41
	}
L40:
	;
	if int32(0) <= v99 {
		v232 = v100
		v233 = v99
		goto L30
	} else {
		goto L42
	}
L41:
	;
	v99 = v75 - int32(158)
	v100 = int32(0)
	goto L40
L42:
	;
	goto L37
L43:
	;
	goto L3
L44:
	;
	v239 = v74<<(uint(int32(1))%32) + v225 - int32(384)
	v240 = v226
	v241 = v30
	goto L29
L45:
	;
	v112 = v75 + int32(-64)
	if base.Ui32(v112) <= base.Ui32(int32(62)) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L47
L47:
	;
	if v37&int32(-4) == int32(-16) {
		goto L60
	} else {
		goto L61
	}
L48:
	;
	v225 = int32(-1)
	v226 = v75 - int32(63)
	goto L44
L49:
	;
	goto L50
L50:
	;
	if v76 < int32(-97) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	if v12 != int64(0) {
		v274 = v28
		v276 = v30
		goto L4
	} else {
		goto L57
	}
L52:
	;
	v126 = v112
	v127 = int32(-1)
	goto L54
L53:
	;
	if int32(-4) < v76 {
		goto L51
	} else {
		goto L55
	}
L54:
	;
	if int32(0) <= v126 {
		v225 = v127
		v226 = v126
		goto L44
	} else {
		goto L56
	}
L55:
	;
	v126 = v75 - int32(158)
	v127 = int32(0)
	goto L54
L56:
	;
	goto L51
L57:
	;
	goto L3
L58:
	;
	v221 = int32(143)
	*(*uint8)(unsafe.Add(mBase, uint32(v30))) = uint8(v221)
	v239 = v220
	v240 = v219
	v241 = v30 + int32(1)
	goto L29
L59:
	;
	switch v74 - int32(240) {
	case 0:
		goto L92
	case 1:
		goto L95
	case 2:
		goto L94
	default:
		goto L93
	}
L60:
	;
	v139 = v75 + int32(-64)
	if base.Ui32(v139) <= base.Ui32(int32(62)) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L62
L62:
	;
	if base.Ui32((v37+int32(12))&int32(255)) <= base.Ui32(int32(8)) {
		goto L74
	} else {
		goto L75
	}
L63:
	;
	v202 = int32(1)
	v203 = v75 - int32(63)
	goto L59
L64:
	;
	goto L65
L65:
	;
	if v76 < int32(-97) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	if v12 != int64(0) {
		v274 = v28
		v276 = v30
		goto L4
	} else {
		goto L72
	}
L67:
	;
	v153 = v139
	v154 = int32(1)
	goto L69
L68:
	;
	if int32(-4) < v76 {
		goto L66
	} else {
		goto L70
	}
L69:
	;
	if int32(0) <= v153 {
		v202 = v154
		v203 = v153
		goto L59
	} else {
		goto L71
	}
L70:
	;
	v153 = v75 - int32(158)
	v154 = int32(0)
	goto L69
L71:
	;
	goto L66
L72:
	;
	goto L3
L73:
	;
	if v37 == int32(-12) {
		goto L88
	} else {
		goto L89
	}
L74:
	;
	v168 = v75 + int32(-64)
	if base.Ui32(v168) <= base.Ui32(int32(62)) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	if v12 != int64(0) {
		v274 = v28
		v276 = v30
		goto L4
	} else {
		goto L87
	}
L77:
	;
	v192 = int32(1)
	v193 = v75 - int32(63)
	goto L73
L78:
	;
	goto L79
L79:
	;
	if v76 < int32(-97) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	if v12 != int64(0) {
		v274 = v28
		v276 = v30
		goto L4
	} else {
		goto L86
	}
L81:
	;
	v182 = v168
	v183 = int32(1)
	goto L83
L82:
	;
	if int32(-4) < v76 {
		goto L80
	} else {
		goto L84
	}
L83:
	;
	if int32(0) <= v182 {
		v192 = v183
		v193 = v182
		goto L73
	} else {
		goto L85
	}
L84:
	;
	v182 = v75 - int32(158)
	v183 = int32(0)
	goto L83
L85:
	;
	goto L80
L86:
	;
	goto L3
L87:
	;
	goto L3
L88:
	;
	if v192 != 0 {
		v219 = v193
		v220 = int32(15)
		goto L58
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v219 = v193
	v220 = v74<<(uint(int32(1))%32) - v192 - int32(410)
	goto L58
L91:
	;
	goto L90
L92:
	;
	if v202 != 0 {
		goto L105
	} else {
		goto L106
	}
L93:
	;
	if v202 != 0 {
		goto L102
	} else {
		goto L103
	}
L94:
	;
	if v202 != 0 {
		goto L99
	} else {
		goto L100
	}
L95:
	;
	if v202 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v208 = int32(3)
	goto L98
L97:
	;
	v208 = int32(4)
	goto L98
L98:
	;
	v219 = v203
	v220 = v208
	goto L58
L99:
	;
	v211 = int32(5)
	goto L101
L100:
	;
	v211 = int32(12)
	goto L101
L101:
	;
	v219 = v203
	v220 = v211
	goto L58
L102:
	;
	v214 = int32(13)
	goto L104
L103:
	;
	v214 = int32(14)
	goto L104
L104:
	;
	v219 = v203
	v220 = v214
	goto L58
L105:
	;
	v217 = int32(1)
	goto L107
L106:
	;
	v217 = int32(8)
	goto L107
L107:
	;
	v219 = v203
	v220 = v217
	goto L58
L108:
	;
	goto L9
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_shortest(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v125 int32
	_ = v125
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v319 int32
	_ = v319
	var v326 int32
	_ = v326
	var v334 int32
	_ = v334
	var v345 int32
	_ = v345
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v359 int32
	_ = v359
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
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
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v557 int32
	_ = v557
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v640 int32
	_ = v640
	v8 = int32(0)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	if l5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = int32(0)
	goto L3
L2:
	;
	goto L3
L3:
	;
	if l6 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = int32(0)
	goto L6
L5:
	;
	goto L6
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	if v22 < int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	return v640
L8:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+8)))
	if v134&int32(2) != 0 {
		goto L47
	} else {
		goto L48
	}
L9:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v28 = v25 + v22<<(uint(int32(3))%32)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v29 < int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v32 = int32(0)
	v35 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+66)))
	v36 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+64)))
	if base.I32_extend16_s(v36) <= v35 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if base.B2i32(l5 == v32)|base.B2i32(v125 == int32(0)) != 0 {
		v640 = v125
		goto L7
	} else {
		goto L46
	}
L12:
	;
	v39 = l2
	goto L14
L13:
	;
	v39 = v32
	goto L14
L14:
	;
	if l2 == l3 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v42 = v39
	goto L17
L16:
	;
	v42 = int32(0)
	goto L17
L17:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v29 == v43 {
		v125 = v42
		goto L11
	} else {
		goto L18
	}
L18:
	;
	v45 = v43 - v29
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui32(l2) < base.Ui32(l3) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v51 = int32(1)
	v53 = base.I32_div_u_s((l3-l2)>>(uint(int32(2))%32)-v51, v45)
	v55 = v53 + v51
	if base.Ui32(v36) < base.Ui32(v55) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v58 = v36
	goto L21
L21:
	;
	v63 = base.I32_div_u_s((l4-l2)>>(uint(int32(2))%32), v45)
	if base.Ui32(v63) < base.Ui32(v35) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v57 = v55
	goto L24
L23:
	;
	v57 = v36
	goto L24
L24:
	;
	v58 = v57
	goto L21
L25:
	;
	v65 = v63
	goto L27
L26:
	;
	v65 = v35
	goto L27
L27:
	;
	if v35 == int32(256) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v68 = v63
	goto L30
L29:
	;
	v68 = v65
	goto L30
L30:
	;
	if base.Ui32(v68) < base.Ui32(v58) {
		v640 = v8
		goto L7
	} else {
		goto L31
	}
L31:
	;
	if v58 == int32(0) {
		v125 = l2
		goto L11
	} else {
		goto L32
	}
L32:
	;
	v72 = int32(2)
	v85 = int32(0)
	v87 = l2
	goto L34
L33:
	;
	if base.Ui32(v58) <= base.Ui32(v106) {
		goto L43
	} else {
		goto L44
	}
L34:
	;
	if v85 == v68 {
		v105 = v87
		v106 = v68
		goto L33
	} else {
		goto L36
	}
L35:
	;
	v105 = v100
	v106 = v58
	goto L33
L36:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+420))
	v96 = m.T0[v95].(func(*base.Module, int32, int32, int32) int32)(m, v46+v29<<(uint(v72)%32), v87, v45)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	return int32(0)
L38:
	;
	if v96 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v105 = v87
	v106 = v85
	goto L33
L40:
	;
	goto L41
L41:
	;
	v100 = v45<<(uint(v72)%32) + v87
	v102 = v85 + int32(1)
	if v102 != v58 {
		v85 = v102
		v87 = v100
		goto L34
	} else {
		goto L42
	}
L42:
	;
	goto L35
L43:
	;
	v109 = v105
	goto L45
L44:
	;
	v109 = int32(0)
	goto L45
L45:
	;
	v125 = v109
	goto L11
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = l2
	return v125
L47:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v133)+44))
	v142 = (l3 - l2) >> (uint(int32(2)) % 32)
	if base.B2i32(v137 != int32(256))&base.B2i32(base.Ui32(v137) < base.Ui32(v142)) != 0 {
		v640 = v8
		goto L7
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v159 = int32(0)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v159 < v167 {
		goto L59
	} else {
		goto L60
	}
L50:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v133)+40))
	if (l4-l2)>>(uint(int32(2))%32) < v145 {
		v640 = v8
		goto L7
	} else {
		goto L51
	}
L51:
	;
	if base.Ui32(v142) < base.Ui32(v145) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v154 = l2 + v145<<(uint(int32(2))%32)
	goto L54
L53:
	;
	v154 = l3
	goto L54
L54:
	;
	if l5 == int32(0) {
		v640 = v154
		goto L7
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = l2
	return v154
L56:
	;
	if v380 == int32(0) {
		v640 = v8
		goto L7
	} else {
		goto L93
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v359)+20)) = l2
	*(*int64)(unsafe.Add(mBase, uint32(l1)+48)) = int64(0)
	v380 = v359
	goto L56
L58:
	;
	v334 = int32(0)
	goto L90
L59:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+8)))
	if v171&int32(1) != 0 {
		v326 = v170
		goto L58
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v175 = F_getvacant(m, l0, l1, l2, l2)
	mBase = m.M
	if v175 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	goto L61
L63:
	;
	v380 = int32(0)
	goto L56
L64:
	;
	goto L65
L65:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if int32(0) < v179 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v183 = int32(0)
	goto L69
L67:
	;
	goto L68
L68:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)+12))
	v222 = v215 + int32(base.Ui32(v217)>>(uint(int32(3))%32))&int32(536870908)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	v224 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v222))) = v223 | v224<<(uint(v217)%32)
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v229 == v224 {
		goto L73
	} else {
		goto L74
	}
L69:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
	*(*int32)(unsafe.Add(mBase, uint32(v194+v183<<(uint(int32(2))%32)))) = int32(0)
	v201 = v183 + int32(1)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v201 < v202 {
		v183 = v201
		goto L69
	} else {
		goto L71
	}
L70:
	;
	goto L68
L71:
	;
	goto L70
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v175)+8)) = int32(13)
	*(*int32)(unsafe.Add(mBase, uint32(v175)+4)) = v308
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v319 <= int32(0) {
		v359 = v175
		goto L57
	} else {
		goto L89
	}
L73:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
	v308 = v232
	goto L72
L74:
	;
	goto L75
L75:
	;
	if v229 <= int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v308 = int32(0)
	goto L72
L77:
	;
	goto L78
L78:
	;
	v237 = v229 & int32(3)
	v238 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v229) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v244 = v238
	v247 = v238
	v253 = v159
	goto L82
L80:
	;
	v273 = v238
	v276 = v238
	goto L81
L81:
	;
	v284 = v273
	v287 = v276
	v294 = v159
	goto L86
L82:
	;
	v257 = v228 + v244<<(uint(int32(2))%32)
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v257)+12))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v257)+8))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v257)+4))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
	v265 = v258 ^ (v259 ^ (v260 ^ (v261 ^ v247)))
	v266 = int32(4)
	v267 = v244 + v266
	v269 = v253 + v266
	if v269 != v229&int32(2147483644) {
		v244 = v267
		v247 = v265
		v253 = v269
		goto L82
	} else {
		goto L84
	}
L83:
	;
	if v237 == int32(0) {
		v308 = v265
		goto L72
	} else {
		goto L85
	}
L84:
	;
	goto L83
L85:
	;
	v273 = v267
	v276 = v265
	goto L81
L86:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v228+v284<<(uint(int32(2))%32))))
	v299 = v298 ^ v287
	v300 = int32(1)
	v303 = v294 + v300
	if v303 != v237 {
		v284 = v284 + v300
		v287 = v299
		v294 = v303
		goto L86
	} else {
		goto L88
	}
L87:
	;
	v308 = v299
	goto L72
L88:
	;
	goto L87
L89:
	;
	v326 = v175
	goto L58
L90:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v345+v334<<(uint(int32(5))%32))+20)) = int32(0)
	v352 = v334 + int32(1)
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v352 < v353 {
		v334 = v352
		goto L90
	} else {
		goto L92
	}
L91:
	;
	v359 = v326
	goto L57
L92:
	;
	goto L91
L93:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v383 == l2 {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v410 = F_miss(m, l0, l1, v380, base.I32_extend16_s(v408), l2, l2)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L37
	} else {
		goto L102
	}
L95:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v389 = int32(1)
	v394 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v385+(v386^int32(-1))&v389<<(uint(v389)%32))+20)))
	v408 = v394
	goto L94
L96:
	;
	goto L97
L97:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l2-int32(4))))
	if base.Ui32(v397) <= base.Ui32(int32(2047)) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	v404 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v400+v397<<(uint(int32(1))%32)))))
	v408 = v404
	goto L94
L99:
	;
	goto L100
L100:
	;
	v405 = F_pg_reg_getcolor(m, v17, v397)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L37
	} else {
		goto L101
	}
L101:
	;
	v408 = v405
	goto L94
L102:
	;
	if v410 == int32(0) {
		v640 = v8
		goto L7
	} else {
		goto L103
	}
L103:
	;
	if l4 != v16 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v417 = int32(4)
	goto L106
L105:
	;
	v417 = int32(0)
	goto L106
L106:
	;
	if l3 != v16 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v422 = int32(4)
	goto L109
L108:
	;
	v422 = int32(0)
	goto L109
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v410)+20)) = l2
	v433 = v410
	v434 = l2
	goto L110
L110:
	;
	if base.Ui32(l4+v417) <= base.Ui32(v434) {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	if l5 != 0 {
		goto L127
	} else {
		goto L128
	}
L112:
	;
	goto L111
L113:
	;
	v477 = v433
	v479 = v434
	goto L112
L114:
	;
	goto L115
L115:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v434)))
	if base.Ui32(v441) <= base.Ui32(int32(2047)) {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v433)+24))
	v453 = base.I32_extend16_s(v451)
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v452+v453<<(uint(int32(2))%32))))
	if v457 == int32(0) {
		goto L121
	} else {
		goto L122
	}
L117:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	v448 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v444+v441<<(uint(int32(1))%32)))))
	v451 = v448
	goto L116
L118:
	;
	goto L119
L119:
	;
	v449 = F_pg_reg_getcolor(m, v17, v441)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L37
	} else {
		goto L120
	}
L120:
	;
	v451 = v449
	goto L116
L121:
	;
	v462 = F_miss(m, l0, l1, v433, v453, v434+int32(4), l2)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L37
	} else {
		goto L124
	}
L122:
	;
	v466 = v457
	goto L123
L123:
	;
	v468 = v434 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v466)+20)) = v468
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v466)+8)))
	if base.B2i32(v470&int32(2) == int32(0))|base.B2i32(base.Ui32(v468) < base.Ui32(l3+v422)) != 0 {
		v433 = v466
		v434 = v468
		goto L110
	} else {
		goto L126
	}
L124:
	;
	if v462 == int32(0) {
		v640 = v8
		goto L7
	} else {
		goto L125
	}
L125:
	;
	v466 = v462
	goto L123
L126:
	;
	v477 = v466
	v479 = v468
	goto L112
L127:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v481 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L128:
	;
	goto L129
L129:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v477)+8))
	v580 = v578 & int32(2)
	v581 = int32(0)
	if base.B2i32(v580 == v581)|base.B2i32(base.Ui32(v479) <= base.Ui32(l3)) == v581 {
		goto L160
	} else {
		goto L161
	}
L130:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v485 = v484
	goto L132
L131:
	;
	v485 = v481
	goto L132
L132:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v486 <= int32(0) {
		v557 = v485
		goto L133
	} else {
		goto L134
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v557
	goto L129
L134:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v486&int32(1) != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+8)))
	if v492&int32(8) != 0 {
		goto L138
	} else {
		goto L139
	}
L136:
	;
	v504 = v489
	v505 = v485
	v507 = v486
	goto L137
L137:
	;
	if v486 == int32(1) {
		v557 = v505
		goto L133
	} else {
		goto L144
	}
L138:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v489)+20))
	if base.Ui32(v485) < base.Ui32(v495) {
		goto L141
	} else {
		goto L142
	}
L139:
	;
	v498 = v485
	goto L140
L140:
	;
	v504 = v489 + int32(32)
	v505 = v498
	v507 = v486 - int32(1)
	goto L137
L141:
	;
	v497 = v495
	goto L143
L142:
	;
	v497 = v485
	goto L143
L143:
	;
	v498 = v497
	goto L140
L144:
	;
	v518 = v504
	v520 = v505
	v521 = v507
	goto L145
L145:
	;
	v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+8)))
	if v525&int32(8) != 0 {
		goto L147
	} else {
		goto L148
	}
L146:
	;
	v557 = v539
	goto L133
L147:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v518)+20))
	if base.Ui32(v520) < base.Ui32(v528) {
		goto L150
	} else {
		goto L151
	}
L148:
	;
	v531 = v520
	goto L149
L149:
	;
	v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+40)))
	if v533&int32(8) != 0 {
		goto L153
	} else {
		goto L154
	}
L150:
	;
	v530 = v528
	goto L152
L151:
	;
	v530 = v520
	goto L152
L152:
	;
	v531 = v530
	goto L149
L153:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v518)+52))
	if base.Ui32(v531) < base.Ui32(v536) {
		goto L156
	} else {
		goto L157
	}
L154:
	;
	v539 = v531
	goto L155
L155:
	;
	v543 = int32(2)
	if v543 < v521 {
		v518 = v518 - int32(-64)
		v520 = v539
		v521 = v521 - v543
		goto L145
	} else {
		goto L159
	}
L156:
	;
	v538 = v536
	goto L158
L157:
	;
	v538 = v531
	goto L158
L158:
	;
	v539 = v538
	goto L155
L159:
	;
	goto L146
L160:
	;
	return v479 - int32(4)
L161:
	;
	goto L162
L162:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if base.B2i32(v479 != v590)|base.B2i32(v590 != l4) != 0 {
		v618 = v580
		goto L164
	} else {
		goto L165
	}
L163:
	;
	if l6 == int32(0) {
		v640 = v8
		goto L7
	} else {
		goto L172
	}
L164:
	;
	if v618 != 0 {
		goto L169
	} else {
		goto L170
	}
L165:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v595 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v601 = int32(*(*int16)(unsafe.Add(mBase, uint32(v594+(v595^int32(-1))&int32(2))+24)))
	v602 = F_miss(m, l0, l1, v477, v601, v479, l2)
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L37
	} else {
		goto L166
	}
L166:
	;
	if v602 == int32(0) {
		goto L163
	} else {
		goto L167
	}
L167:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v602)+8))
	v608 = v606 & int32(2)
	if v608|base.B2i32(l6 == int32(0)) != 0 {
		v618 = v608
		goto L164
	} else {
		goto L168
	}
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = int32(1)
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v602)+8))
	v618 = v614 & int32(2)
	goto L164
L169:
	;
	v620 = v479
	goto L171
L170:
	;
	v620 = int32(0)
	goto L171
L171:
	;
	return v620
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = int32(1)
	v640 = v8
	goto L7
}
func F_show_incremental_sort_group_info(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
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
	var v38 int32
	_ = v38
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
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int64
	_ = v84
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v149 int64
	_ = v149
	var v152 int64
	_ = v152
	var v153 int64
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int64
	_ = v159
	var v162 int64
	_ = v162
	var v169 int32
	_ = v169
	var v173 int64
	_ = v173
	var v176 int64
	_ = v176
	var v177 int64
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int64
	_ = v183
	var v186 int64
	_ = v186
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int64
	_ = v209
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int64
	_ = v215
	var v218 int64
	_ = v218
	var v219 int64
	_ = v219
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v241 int64
	_ = v241
	var v243 int32
	_ = v243
	var v246 int64
	_ = v246
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v258 int64
	_ = v258
	var v261 int64
	_ = v261
	var v262 int64
	_ = v262
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v284 int64
	_ = v284
	var v286 int32
	_ = v286
	var v289 int64
	_ = v289
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v305 int32
	_ = v305
	v10 = m.G0
	v12 = v10 - int32(160)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v14&int32(1) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	goto L5
L2:
	;
	v29 = int32(0)
	v30 = v14
	goto L3
L3:
	;
	if v30&int32(2) != 0 {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	v26 = F_lappend(m, int32(0), v23)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_show_incremental_sort_group_info[0]))
	goto L7
L7:
	;
	goto L4
L8:
	;
	return
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v29 = v26
	v30 = v28
	goto L3
L10:
	;
	goto L14
L11:
	;
	v44 = v29
	v45 = v30
	goto L12
L12:
	;
	if v45&int32(4) != 0 {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	v41 = F_lappend(m, v29, v38)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L8
	} else {
		goto L17
	}
L14:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_show_incremental_sort_group_info[1]))
	goto L16
L16:
	;
	goto L13
L17:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v44 = v41
	v45 = v43
	goto L12
L18:
	;
	goto L22
L19:
	;
	v59 = v44
	v60 = v45
	goto L20
L20:
	;
	if v60&int32(8) != 0 {
		goto L26
	} else {
		goto L27
	}
L21:
	;
	v56 = F_lappend(m, v44, v53)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L8
	} else {
		goto L25
	}
L22:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_show_incremental_sort_group_info[2]))
	goto L24
L24:
	;
	goto L21
L25:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v59 = v56
	v60 = v58
	goto L20
L26:
	;
	goto L30
L27:
	;
	v73 = v59
	goto L28
L28:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	if v74 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L29:
	;
	v71 = F_lappend(m, v59, v68)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L8
	} else {
		goto L33
	}
L30:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_show_incremental_sort_group_info[3]))
	goto L32
L32:
	;
	goto L29
L33:
	;
	v73 = v71
	goto L28
L34:
	;
	m.G0 = v12 + int32(160)
	return
L35:
	;
	if l2 != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	v193 = v12 + int32(144)
	F_initStringInfo(m, v193)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L8
	} else {
		goto L75
	}
L38:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	F_appendStringInfoSpaces(m, v77, v78<<(uint(int32(1))%32))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L8
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v84 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+72)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = l1
	F_appendStringInfo(m, v83, int32(_a_F_show_incremental_sort_group_info_0), v12-int32(-64))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L8
	} else {
		goto L42
	}
L41:
	;
	goto L40
L42:
	;
	if v73 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v149 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if int64(0) < v149 {
		goto L61
	} else {
		goto L62
	}
L44:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	F_appendStringInfoString(m, v94, int32(_a_F_show_incremental_sort_group_info_1))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L8
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	if int32(1) < v101 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L43
L48:
	;
	v104 = int32(_a_F_show_incremental_sort_group_info_2)
	goto L50
L49:
	;
	v104 = int32(_a_F_show_incremental_sort_group_info_1)
	goto L50
L50:
	;
	F_appendStringInfoString(m, v98, v104)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L8
	} else {
		goto L51
	}
L51:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	if v107 <= int32(0) {
		goto L43
	} else {
		goto L52
	}
L52:
	;
	v117 = int32(0)
	goto L53
L53:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v121+v117<<(uint(int32(2))%32))))
	F_appendStringInfoString(m, v120, v125)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L8
	} else {
		goto L55
	}
L54:
	;
	goto L43
L55:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	if v117 < v128-int32(1) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	F_appendStringInfoString(m, v132, int32(_a_F_show_incremental_sort_group_info_3))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L8
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v137 = v117 + int32(1)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	if v137 < v138 {
		v117 = v137
		goto L53
	} else {
		goto L60
	}
L59:
	;
	goto L58
L60:
	;
	goto L54
L61:
	;
	v152 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v153 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v155 = int32(_a_F_show_incremental_sort_group_info_4)
	goto L65
L62:
	;
	goto L63
L63:
	;
	v173 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	if v173 <= int64(0) {
		goto L34
	} else {
		goto L69
	}
L64:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v159 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+56)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v155
	v162 = base.I64_div_s(v153, v152)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v155
	F_appendStringInfo(m, v158, int32(_a_F_show_incremental_sort_group_info_5), v12+int32(32))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L8
	} else {
		goto L68
	}
L65:
	;
	goto L67
L67:
	;
	goto L64
L68:
	;
	goto L63
L69:
	;
	v176 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v177 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v180 = int32(_a_F_show_incremental_sort_group_info_6)
	goto L72
L70:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v183 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v180
	v186 = base.I64_div_s(v177, v176)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v186
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v180
	F_appendStringInfo(m, v182, int32(_a_F_show_incremental_sort_group_info_5), v12)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L8
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	goto L70
L74:
	;
	goto L34
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+112)) = l1
	F_appendStringInfo(m, v193, int32(_a_F_show_incremental_sort_group_info_7), v12+int32(112))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L8
	} else {
		goto L76
	}
L76:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
	F_ExplainOpenGroup(m, int32(_a_F_show_incremental_sort_group_info_8), v203, int32(1), l3)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L8
	} else {
		goto L77
	}
L77:
	;
	v209 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	F_ExplainPropertyInteger(m, int32(_a_F_show_incremental_sort_group_info_9), int32(0), v209, l3)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L8
	} else {
		goto L78
	}
L78:
	;
	F_ExplainPropertyList(m, int32(_a_F_show_incremental_sort_group_info_10), v73, l3)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L8
	} else {
		goto L79
	}
L79:
	;
	v215 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if int64(0) < v215 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v218 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v219 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	goto L84
L81:
	;
	goto L82
L82:
	;
	v258 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	if int64(0) < v258 {
		goto L93
	} else {
		goto L94
	}
L83:
	;
	v225 = v12 + int32(128)
	F_initStringInfo(m, v225)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L8
	} else {
		goto L87
	}
L84:
	;
	goto L86
L86:
	;
	goto L83
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = int32(_a_F_show_incremental_sort_group_info_4)
	F_appendStringInfo(m, v225, int32(_a_F_show_incremental_sort_group_info_11), v12+int32(96))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L8
	} else {
		goto L88
	}
L88:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v12)+128))
	F_ExplainOpenGroup(m, int32(_a_F_show_incremental_sort_group_info_12), v235, int32(1), l3)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L8
	} else {
		goto L89
	}
L89:
	;
	v241 = base.I64_div_s(v219, v218)
	F_ExplainPropertyInteger(m, int32(_a_F_show_incremental_sort_group_info_13), int32(_a_F_show_incremental_sort_group_info_14), v241, l3)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L8
	} else {
		goto L90
	}
L90:
	;
	v246 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	F_ExplainPropertyInteger(m, int32(_a_F_show_incremental_sort_group_info_15), int32(_a_F_show_incremental_sort_group_info_14), v246, l3)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L8
	} else {
		goto L91
	}
L91:
	;
	F_ExplainCloseGroup(m, int32(_a_F_show_incremental_sort_group_info_12), int32(1), l3)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L8
	} else {
		goto L92
	}
L92:
	;
	goto L82
L93:
	;
	v261 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v262 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	goto L98
L94:
	;
	goto L95
L95:
	;
	F_ExplainCloseGroup(m, int32(_a_F_show_incremental_sort_group_info_8), int32(1), l3)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L8
	} else {
		goto L106
	}
L96:
	;
	v268 = v12 + int32(128)
	F_initStringInfo(m, v268)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L8
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	goto L96
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = int32(_a_F_show_incremental_sort_group_info_6)
	F_appendStringInfo(m, v268, int32(_a_F_show_incremental_sort_group_info_11), v12+int32(80))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L8
	} else {
		goto L101
	}
L101:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v12)+128))
	F_ExplainOpenGroup(m, int32(_a_F_show_incremental_sort_group_info_12), v278, int32(1), l3)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L8
	} else {
		goto L102
	}
L102:
	;
	v284 = base.I64_div_s(v262, v261)
	F_ExplainPropertyInteger(m, int32(_a_F_show_incremental_sort_group_info_13), int32(_a_F_show_incremental_sort_group_info_14), v284, l3)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L8
	} else {
		goto L103
	}
L103:
	;
	v289 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	F_ExplainPropertyInteger(m, int32(_a_F_show_incremental_sort_group_info_15), int32(_a_F_show_incremental_sort_group_info_14), v289, l3)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L8
	} else {
		goto L104
	}
L104:
	;
	F_ExplainCloseGroup(m, int32(_a_F_show_incremental_sort_group_info_12), int32(1), l3)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L8
	} else {
		goto L105
	}
L105:
	;
	goto L95
L106:
	;
	goto L34
}
func F_sigemptyset(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	return
}
func F_slice_to(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
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
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	v9 = int32(-1)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v10 < int32(0) {
		v57 = v9
		return v57
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v13 < v10 {
			v57 = v9
			return v57
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v15 < v13 {
				v57 = v9
				return v57
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v17-int32(4))))
				if v20 < v15 {
					v57 = v9
					return v57
				} else {
					v22 = v13 - v10
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v25 = v23 - int32(8)
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
					if v26 < v22 {
						v30 = F_repalloc(m, v25, v22+int32(29))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							if v30 == int32(0) {
								v57 = v9
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v30))) = v22 + int32(20)
								v40 = v30 + int32(8)
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = v40
								v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								v44 = v43
								v45 = v40
								v46 = v42
								if v22 != 0 {
									base.MemoryCopy(m, v45, v44+v46, v22)
								} else {
								}
								v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								*(*int32)(unsafe.Add(mBase, uint32(v49-int32(4)))) = v22
								v57 = int32(0)
							}
							return v57
						}
					} else {
						v44 = v10
						v45 = v23
						v46 = v17
						if v22 != 0 {
							base.MemoryCopy(m, v45, v44+v46, v22)
						} else {
						}
						v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						*(*int32)(unsafe.Add(mBase, uint32(v49-int32(4)))) = v22
						v57 = int32(0)
						return v57
					}
				}
			}
		}
	}
}
func F_slotsync_reread_config(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_reread_config[0]))
	v15 = F_pstrdup(m, v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_reread_config[1]))
	v19 = F_pstrdup(m, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_slotsync_reread_config[2])) = int32(0)
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_reread_config[3]))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_slotsync_reread_config[4])))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_slotsync_reread_config[5])))
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_reread_config[0]))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	if base.B2i32(v37 == int32(0))|base.B2i32(v37 != v40) != 0 {
		v58 = v37
		v59 = v40
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_reread_config[1]))
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	if base.B2i32(v65 == int32(0))|base.B2i32(v65 != v68) != 0 {
		v86 = v65
		v87 = v68
		goto L13
	} else {
		goto L14
	}
L6:
	;
	goto L5
L7:
	;
	v43 = v15
	v44 = v34
	goto L8
L8:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+1)))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+1)))
	if v48 == int32(0) {
		v58 = v48
		v59 = v47
		goto L6
	} else {
		goto L10
	}
L9:
	;
	v58 = v48
	v59 = v47
	goto L6
L10:
	;
	v51 = int32(1)
	if v48 == v47 {
		v43 = v43 + v51
		v44 = v44 + v51
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	F_pfree(m, v15)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L19
	}
L13:
	;
	goto L12
L14:
	;
	v71 = v19
	v72 = v62
	goto L15
L15:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+1)))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
	if v76 == int32(0) {
		v86 = v76
		v87 = v75
		goto L13
	} else {
		goto L17
	}
L16:
	;
	v86 = v76
	v87 = v75
	goto L13
L17:
	;
	v79 = int32(1)
	if v76 == v75 {
		v71 = v71 + v79
		v72 = v72 + v79
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	F_pfree(m, v19)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_slotsync_reread_config[5])))
	if v94 != v29 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	m.G0 = v11 + int32(16)
	return
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L46
	}
L23:
	;
	if v25 != int32(7) {
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if v58-v59|(v86-v87) == int32(0) {
		goto L34
	} else {
		goto L35
	}
L26:
	;
	v100 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if v100 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(_a_F_slotsync_reread_config_0)
	F_errmsg(m, int32(_a_F_slotsync_reread_config_1), v11)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	F_errfinish(m, int32(_a_F_slotsync_reread_config_2), int32(1322), int32(_a_F_slotsync_reread_config_3))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L34:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_slotsync_reread_config[4])))
	if v27 == v119 {
		goto L21
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	if v25 != int32(7) {
		goto L22
	} else {
		goto L38
	}
L37:
	;
	goto L36
L38:
	;
	v125 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	if v125 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	F_errmsg(m, int32(_a_F_slotsync_reread_config_4), int32(0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v137 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_reread_config[6]))
	*(*int64)(unsafe.Add(mBase, uint32(v137)+8)) = int64(0)
	F_proc_exit(m, int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	F_errfinish(m, int32(_a_F_slotsync_reread_config_2), int32(1339), int32(_a_F_slotsync_reread_config_3))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L46:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_errmsg(m, int32(_a_F_slotsync_reread_config_5), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_slotsync_reread_config_2), int32(1364), int32(_a_F_slotsync_reread_config_3))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_slotsync_worker_onexit(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_worker_onexit[0]))
	if v4 != 0 {
		F_ReplicationSlotRelease(m)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			F_ReplicationSlotCleanup(m, int32(0))
			mBase = m.M
			v9 = m.ExcPending
			if v9 != 0 {
				return
			} else {
				v11 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_worker_onexit[1]))
				v14 = base.AtomicRmwXchg32(m, v11, int32(16), int32(1))
				if v14 != 0 {
					F_s_lock(m, v11+int32(16), int32(_a_F_slotsync_worker_onexit_0))
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return
					} else {
						v20 = int32(_a_F_slotsync_worker_onexit_1)
						v21 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_worker_onexit[1]))
						*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(-1)
						v25 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_worker_onexit[1]))
						v27 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_slotsync_worker_onexit[2])))
						if v27 != 0 {
							v28 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v25)+5)) = uint8(v28)
							*(*uint8)(unsafe.Add(mBase, _c_F_slotsync_worker_onexit[2])) = uint8(v28)
						} else {
						}
						v33 = int32(0)
						atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v25)+16)), uint32(v33))
						return
					}
				} else {
					v20 = int32(_a_F_slotsync_worker_onexit_1)
					v21 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_worker_onexit[1]))
					*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(-1)
					v25 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_worker_onexit[1]))
					v27 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_slotsync_worker_onexit[2])))
					if v27 != 0 {
						v28 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v25)+5)) = uint8(v28)
						*(*uint8)(unsafe.Add(mBase, _c_F_slotsync_worker_onexit[2])) = uint8(v28)
					} else {
					}
					v33 = int32(0)
					atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v25)+16)), uint32(v33))
					return
				}
			}
		}
	} else {
		F_ReplicationSlotCleanup(m, int32(0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_worker_onexit[1]))
			v14 = base.AtomicRmwXchg32(m, v11, int32(16), int32(1))
			if v14 != 0 {
				F_s_lock(m, v11+int32(16), int32(_a_F_slotsync_worker_onexit_0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					v20 = int32(_a_F_slotsync_worker_onexit_1)
					v21 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_worker_onexit[1]))
					*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(-1)
					v25 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_worker_onexit[1]))
					v27 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_slotsync_worker_onexit[2])))
					if v27 != 0 {
						v28 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v25)+5)) = uint8(v28)
						*(*uint8)(unsafe.Add(mBase, _c_F_slotsync_worker_onexit[2])) = uint8(v28)
					} else {
					}
					v33 = int32(0)
					atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v25)+16)), uint32(v33))
					return
				}
			} else {
				v20 = int32(_a_F_slotsync_worker_onexit_1)
				v21 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_worker_onexit[1]))
				*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(-1)
				v25 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_worker_onexit[1]))
				v27 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_slotsync_worker_onexit[2])))
				if v27 != 0 {
					v28 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v25)+5)) = uint8(v28)
					*(*uint8)(unsafe.Add(mBase, _c_F_slotsync_worker_onexit[2])) = uint8(v28)
				} else {
				}
				v33 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v25)+16)), uint32(v33))
				return
			}
		}
	}
}
func F_smgrdestroyall(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
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
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	v1 = int32(0)
	v5 = int32(_a_F_smgrdestroyall_0)
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdestroyall[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrdestroyall[0])) = v7 + int32(1)
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdestroyall[1]))
	if base.B2i32(v12 == v1)|base.B2i32(v12 == int32(_a_F_smgrdestroyall_1)) == v1 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L7
	} else {
		goto L15
	}
L2:
	;
	v20 = v12
	goto L5
L3:
	;
	goto L4
L4:
	;
	v70 = int32(_a_F_smgrdestroyall_0)
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdestroyall[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrdestroyall[0])) = v72 - int32(1)
	return
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v25 = int32(_a_F_smgrdestroyall_0)
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdestroyall[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrdestroyall[0])) = v27 + int32(1)
	v32 = v20 - int32(76)
	F_mdclose(m, v32, int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L4
L7:
	;
	return
L8:
	;
	F_mdclose(m, v32, int32(1))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	F_mdclose(m, v32, int32(2))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	F_mdclose(m, v32, int32(3))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = v46
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	*(*int32)(unsafe.Add(mBase, uint32(v46))) = v48
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdestroyall[2]))
	v54 = F_hash_search(m, v51, v32, int32(2), int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	if v54 == int32(0) {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v58 = int32(_a_F_smgrdestroyall_0)
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdestroyall[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrdestroyall[0])) = v60 - int32(1)
	if v24 != int32(_a_F_smgrdestroyall_1) {
		v20 = v24
		goto L5
	} else {
		goto L14
	}
L14:
	;
	goto L6
L15:
	;
	F_errmsg_internal(m, int32(_a_F_smgrdestroyall_2), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(_a_F_smgrdestroyall_3), int32(339), int32(_a_F_smgrdestroyall_4))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_smgrnblocks_cached(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_smgrnblocks_cached[0])))
	if v4 == int32(1) {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0+l1<<(uint(int32(2))%32))+20))
		if v10 != int32(-1) {
			v15 = v10
		} else {
			v15 = int32(-1)
		}
	} else {
		v15 = int32(-1)
	}
	return v15
}
func F_smgrprefetch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v105 int32
	_ = v105
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	v11 = int32(_a_F_smgrprefetch_0)
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_smgrprefetch[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrprefetch[0])) = v13 + int32(1)
	if base.Ui64(int64(4294967295)) < base.Ui64(base.I64_extend_i32_s(l3)+base.I64_extend_i32_u(l2)) {
		v114 = int32(0)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v118 = int32(_a_F_smgrprefetch_0)
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_smgrprefetch[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrprefetch[0])) = v120 - int32(1)
	return v114
L2:
	;
	if l3 <= int32(0) {
		v114 = int32(1)
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = l2
	v28 = l3
	goto L4
L4:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_smgrprefetch[1])))
	if v39 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v114 = v46
	goto L1
L6:
	;
	v40 = int32(2)
	goto L8
L7:
	;
	v40 = int32(1)
	goto L8
L8:
	;
	v41 = F__mdfd_getseg(m, l0, l1, v27, int32(0), v40)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	v45 = int32(0)
	v46 = base.B2i32(v41 != v45)
	if v41 == v45 {
		v114 = v46
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v50 = v27 & int32(_a_F_smgrprefetch_1)
	v55 = int32(_a_F_smgrprefetch_2) - v50
	if base.Ui32(v28) < base.Ui32(v55) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v57 = v28
	goto L14
L13:
	;
	v57 = v55
	goto L14
L14:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v62 = F_FileAccess(m, v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	if int32(0) <= v62 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	goto L19
L17:
	;
	goto L18
L18:
	;
	v105 = v28 - v57
	if int32(0) < v105 {
		v27 = v27 + v57
		v28 = v105
		goto L4
	} else {
		goto L22
	}
L19:
	;
	v78 = int32(_a_F_smgrprefetch_3)
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_smgrprefetch[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v79))) = int32(167772182)
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_smgrprefetch[3]))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v83+v61*int32(48))))
	v87 = F_posix_fadvise(m, v85, base.I64_extend_i32_u(v50<<(uint(int32(13))%32)), base.I64_extend_i32_u(v57<<(uint(int32(13))%32)), int32(3))
	mBase = m.M
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_smgrprefetch[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v89))) = int32(0)
	if v87 == int32(27) {
		goto L19
	} else {
		goto L21
	}
L20:
	;
	goto L18
L21:
	;
	goto L20
L22:
	;
	goto L5
}
func F_sort(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
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
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum_copy(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = int32(1)
		v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
		if v18 != int32(2) {
			v66 = int32(0)
			v67 = v17
			v69 = v2
			v70 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
			if v70 != 0 {
				v71 = F_array_contains_nulls(m, v13)
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return int64(0)
				} else {
					if v71 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v148 = m.ExcPending
						if v148 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(67108994))
							mBase = m.M
							v151 = m.ExcPending
							if v151 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_F_sort_0), int32(0))
								mBase = m.M
								v155 = m.ExcPending
								if v155 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_sort_1), int32(207), int32(_a_F_sort_2))
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v73 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
						v75 = v13 + int32(16)
						v76 = F_ArrayGetNItemsSafe(m, v73, v75)
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return int64(0)
						} else {
							if int32(2) <= v76 {
								v80 = int32(1)
								if v67 != 0 {
									v119 = v80
									v120 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
									v121 = F_ArrayGetNItemsSafe(m, v120, v75)
									mBase = m.M
									v122 = m.ExcPending
									if v122 != 0 {
										return int64(0)
									} else {
										*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v119)
										v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
										if v124 != 0 {
											v132 = v124
										} else {
											v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
											v132 = (v125<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										}
										F_isort(m, v132+v13, v121, v10+int32(15))
										mBase = m.M
										m.G0 = v10 + int32(16)
										return base.I64_extend_i32_u(v13)
									}
								} else {
									switch v69 - int32(3) {
									case 0:
										v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
										if v83|int32(32) != int32(97) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v165 = m.ExcPending
											if v165 != 0 {
												return int64(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v168 = m.ExcPending
												if v168 != 0 {
													return int64(0)
												} else {
													F_errmsg(m, int32(_a_F_sort_3), int32(0))
													mBase = m.M
													v172 = m.ExcPending
													if v172 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_sort_1), int32(225), int32(_a_F_sort_2))
														mBase = m.M
														v177 = m.ExcPending
														if v177 != 0 {
															return int64(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										} else {
											v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
											if v88|int32(32) != int32(115) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v165 = m.ExcPending
												if v165 != 0 {
													return int64(0)
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v168 = m.ExcPending
													if v168 != 0 {
														return int64(0)
													} else {
														F_errmsg(m, int32(_a_F_sort_3), int32(0))
														mBase = m.M
														v172 = m.ExcPending
														if v172 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(_a_F_sort_1), int32(225), int32(_a_F_sort_2))
															mBase = m.M
															v177 = m.ExcPending
															if v177 != 0 {
																return int64(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+2)))
												if v93|int32(32) == int32(99) {
													v119 = v80
													v120 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
													v121 = F_ArrayGetNItemsSafe(m, v120, v75)
													mBase = m.M
													v122 = m.ExcPending
													if v122 != 0 {
														return int64(0)
													} else {
														*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v119)
														v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
														if v124 != 0 {
															v132 = v124
														} else {
															v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
															v132 = (v125<<(uint(int32(3))%32) + int32(23)) & int32(-8)
														}
														F_isort(m, v132+v13, v121, v10+int32(15))
														mBase = m.M
														m.G0 = v10 + int32(16)
														return base.I64_extend_i32_u(v13)
													}
												} else {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v165 = m.ExcPending
													if v165 != 0 {
														return int64(0)
													} else {
														F_errcode(m, int32(50856066))
														mBase = m.M
														v168 = m.ExcPending
														if v168 != 0 {
															return int64(0)
														} else {
															F_errmsg(m, int32(_a_F_sort_3), int32(0))
															mBase = m.M
															v172 = m.ExcPending
															if v172 != 0 {
																return int64(0)
															} else {
																F_errfinish(m, int32(_a_F_sort_1), int32(225), int32(_a_F_sort_2))
																mBase = m.M
																v177 = m.ExcPending
																if v177 != 0 {
																	return int64(0)
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
									case 1:
										v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
										if v98|int32(32) != int32(100) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v165 = m.ExcPending
											if v165 != 0 {
												return int64(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v168 = m.ExcPending
												if v168 != 0 {
													return int64(0)
												} else {
													F_errmsg(m, int32(_a_F_sort_3), int32(0))
													mBase = m.M
													v172 = m.ExcPending
													if v172 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_sort_1), int32(225), int32(_a_F_sort_2))
														mBase = m.M
														v177 = m.ExcPending
														if v177 != 0 {
															return int64(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										} else {
											v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
											if v103|int32(32) != int32(101) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v165 = m.ExcPending
												if v165 != 0 {
													return int64(0)
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v168 = m.ExcPending
													if v168 != 0 {
														return int64(0)
													} else {
														F_errmsg(m, int32(_a_F_sort_3), int32(0))
														mBase = m.M
														v172 = m.ExcPending
														if v172 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(_a_F_sort_1), int32(225), int32(_a_F_sort_2))
															mBase = m.M
															v177 = m.ExcPending
															if v177 != 0 {
																return int64(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+2)))
												if v108|int32(32) != int32(115) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v165 = m.ExcPending
													if v165 != 0 {
														return int64(0)
													} else {
														F_errcode(m, int32(50856066))
														mBase = m.M
														v168 = m.ExcPending
														if v168 != 0 {
															return int64(0)
														} else {
															F_errmsg(m, int32(_a_F_sort_3), int32(0))
															mBase = m.M
															v172 = m.ExcPending
															if v172 != 0 {
																return int64(0)
															} else {
																F_errfinish(m, int32(_a_F_sort_1), int32(225), int32(_a_F_sort_2))
																mBase = m.M
																v177 = m.ExcPending
																if v177 != 0 {
																	return int64(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													}
												} else {
													v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+3)))
													if v114|int32(32) != int32(99) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v165 = m.ExcPending
														if v165 != 0 {
															return int64(0)
														} else {
															F_errcode(m, int32(50856066))
															mBase = m.M
															v168 = m.ExcPending
															if v168 != 0 {
																return int64(0)
															} else {
																F_errmsg(m, int32(_a_F_sort_3), int32(0))
																mBase = m.M
																v172 = m.ExcPending
																if v172 != 0 {
																	return int64(0)
																} else {
																	F_errfinish(m, int32(_a_F_sort_1), int32(225), int32(_a_F_sort_2))
																	mBase = m.M
																	v177 = m.ExcPending
																	if v177 != 0 {
																		return int64(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														}
													} else {
														v119 = int32(0)
														v120 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
														v121 = F_ArrayGetNItemsSafe(m, v120, v75)
														mBase = m.M
														v122 = m.ExcPending
														if v122 != 0 {
															return int64(0)
														} else {
															*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v119)
															v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
															if v124 != 0 {
																v132 = v124
															} else {
																v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
																v132 = (v125<<(uint(int32(3))%32) + int32(23)) & int32(-8)
															}
															F_isort(m, v132+v13, v121, v10+int32(15))
															mBase = m.M
															m.G0 = v10 + int32(16)
															return base.I64_extend_i32_u(v13)
														}
													}
												}
											}
										}
									default:
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v165 = m.ExcPending
										if v165 != 0 {
											return int64(0)
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v168 = m.ExcPending
											if v168 != 0 {
												return int64(0)
											} else {
												F_errmsg(m, int32(_a_F_sort_3), int32(0))
												mBase = m.M
												v172 = m.ExcPending
												if v172 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(_a_F_sort_1), int32(225), int32(_a_F_sort_2))
													mBase = m.M
													v177 = m.ExcPending
													if v177 != 0 {
														return int64(0)
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
								m.G0 = v10 + int32(16)
								return base.I64_extend_i32_u(v13)
							}
						}
					}
				}
			} else {
				v73 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
				v75 = v13 + int32(16)
				v76 = F_ArrayGetNItemsSafe(m, v73, v75)
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return int64(0)
				} else {
					if int32(2) <= v76 {
						v80 = int32(1)
						if v67 != 0 {
							v119 = v80
							v120 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
							v121 = F_ArrayGetNItemsSafe(m, v120, v75)
							mBase = m.M
							v122 = m.ExcPending
							if v122 != 0 {
								return int64(0)
							} else {
								*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v119)
								v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
								if v124 != 0 {
									v132 = v124
								} else {
									v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
									v132 = (v125<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								}
								F_isort(m, v132+v13, v121, v10+int32(15))
								mBase = m.M
								m.G0 = v10 + int32(16)
								return base.I64_extend_i32_u(v13)
							}
						} else {
							switch v69 - int32(3) {
							case 0:
								v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
								if v83|int32(32) != int32(97) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v165 = m.ExcPending
									if v165 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v168 = m.ExcPending
										if v168 != 0 {
											return int64(0)
										} else {
											F_errmsg(m, int32(_a_F_sort_3), int32(0))
											mBase = m.M
											v172 = m.ExcPending
											if v172 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_sort_1), int32(225), int32(_a_F_sort_2))
												mBase = m.M
												v177 = m.ExcPending
												if v177 != 0 {
													return int64(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
									if v88|int32(32) != int32(115) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v165 = m.ExcPending
										if v165 != 0 {
											return int64(0)
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v168 = m.ExcPending
											if v168 != 0 {
												return int64(0)
											} else {
												F_errmsg(m, int32(_a_F_sort_3), int32(0))
												mBase = m.M
												v172 = m.ExcPending
												if v172 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(_a_F_sort_1), int32(225), int32(_a_F_sort_2))
													mBase = m.M
													v177 = m.ExcPending
													if v177 != 0 {
														return int64(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									} else {
										v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+2)))
										if v93|int32(32) == int32(99) {
											v119 = v80
											v120 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
											v121 = F_ArrayGetNItemsSafe(m, v120, v75)
											mBase = m.M
											v122 = m.ExcPending
											if v122 != 0 {
												return int64(0)
											} else {
												*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v119)
												v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
												if v124 != 0 {
													v132 = v124
												} else {
													v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
													v132 = (v125<<(uint(int32(3))%32) + int32(23)) & int32(-8)
												}
												F_isort(m, v132+v13, v121, v10+int32(15))
												mBase = m.M
												m.G0 = v10 + int32(16)
												return base.I64_extend_i32_u(v13)
											}
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v165 = m.ExcPending
											if v165 != 0 {
												return int64(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v168 = m.ExcPending
												if v168 != 0 {
													return int64(0)
												} else {
													F_errmsg(m, int32(_a_F_sort_3), int32(0))
													mBase = m.M
													v172 = m.ExcPending
													if v172 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_sort_1), int32(225), int32(_a_F_sort_2))
														mBase = m.M
														v177 = m.ExcPending
														if v177 != 0 {
															return int64(0)
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
							case 1:
								v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
								if v98|int32(32) != int32(100) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v165 = m.ExcPending
									if v165 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v168 = m.ExcPending
										if v168 != 0 {
											return int64(0)
										} else {
											F_errmsg(m, int32(_a_F_sort_3), int32(0))
											mBase = m.M
											v172 = m.ExcPending
											if v172 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_sort_1), int32(225), int32(_a_F_sort_2))
												mBase = m.M
												v177 = m.ExcPending
												if v177 != 0 {
													return int64(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
									if v103|int32(32) != int32(101) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v165 = m.ExcPending
										if v165 != 0 {
											return int64(0)
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v168 = m.ExcPending
											if v168 != 0 {
												return int64(0)
											} else {
												F_errmsg(m, int32(_a_F_sort_3), int32(0))
												mBase = m.M
												v172 = m.ExcPending
												if v172 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(_a_F_sort_1), int32(225), int32(_a_F_sort_2))
													mBase = m.M
													v177 = m.ExcPending
													if v177 != 0 {
														return int64(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									} else {
										v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+2)))
										if v108|int32(32) != int32(115) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v165 = m.ExcPending
											if v165 != 0 {
												return int64(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v168 = m.ExcPending
												if v168 != 0 {
													return int64(0)
												} else {
													F_errmsg(m, int32(_a_F_sort_3), int32(0))
													mBase = m.M
													v172 = m.ExcPending
													if v172 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_sort_1), int32(225), int32(_a_F_sort_2))
														mBase = m.M
														v177 = m.ExcPending
														if v177 != 0 {
															return int64(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										} else {
											v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+3)))
											if v114|int32(32) != int32(99) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v165 = m.ExcPending
												if v165 != 0 {
													return int64(0)
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v168 = m.ExcPending
													if v168 != 0 {
														return int64(0)
													} else {
														F_errmsg(m, int32(_a_F_sort_3), int32(0))
														mBase = m.M
														v172 = m.ExcPending
														if v172 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(_a_F_sort_1), int32(225), int32(_a_F_sort_2))
															mBase = m.M
															v177 = m.ExcPending
															if v177 != 0 {
																return int64(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												v119 = int32(0)
												v120 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
												v121 = F_ArrayGetNItemsSafe(m, v120, v75)
												mBase = m.M
												v122 = m.ExcPending
												if v122 != 0 {
													return int64(0)
												} else {
													*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v119)
													v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
													if v124 != 0 {
														v132 = v124
													} else {
														v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
														v132 = (v125<<(uint(int32(3))%32) + int32(23)) & int32(-8)
													}
													F_isort(m, v132+v13, v121, v10+int32(15))
													mBase = m.M
													m.G0 = v10 + int32(16)
													return base.I64_extend_i32_u(v13)
												}
											}
										}
									}
								}
							default:
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v165 = m.ExcPending
								if v165 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v168 = m.ExcPending
									if v168 != 0 {
										return int64(0)
									} else {
										F_errmsg(m, int32(_a_F_sort_3), int32(0))
										mBase = m.M
										v172 = m.ExcPending
										if v172 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_sort_1), int32(225), int32(_a_F_sort_2))
											mBase = m.M
											v177 = m.ExcPending
											if v177 != 0 {
												return int64(0)
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
						m.G0 = v10 + int32(16)
						return base.I64_extend_i32_u(v13)
					}
				}
			}
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			v23 = F_pg_detoast_datum_packed(m, v22)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int64(0)
			} else {
				if v23 == int32(0) {
					v66 = int32(0)
					v67 = v17
					v69 = v2
				} else {
					v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
					if v28 == int32(1) {
						v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
						if v34 == int32(18) {
							v37 = int32(16)
						} else {
							v37 = int32(0)
						}
						if base.Ui32((v34-int32(1))&int32(255)) < base.Ui32(int32(3)) {
							v44 = int32(4)
						} else {
							v44 = v37
						}
						v57 = v17
						v58 = v44
					} else {
						if v28&int32(1) != 0 {
							v47 = int32(1)
							v57 = v28
							v58 = int32(base.Ui32(v28)>>(uint(v47)%32)) - v47
						} else {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
							v57 = v51
							v58 = int32(base.Ui32(v51)>>(uint(int32(2))%32)) - int32(4)
						}
					}
					v59 = int32(1)
					if v57&v59 != 0 {
						v63 = v59
					} else {
						v63 = int32(4)
					}
					v66 = v23 + v63
					v67 = int32(0)
					v69 = v58
				}
				v70 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
				if v70 != 0 {
					v71 = F_array_contains_nulls(m, v13)
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return int64(0)
					} else {
						if v71 != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v148 = m.ExcPending
							if v148 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(67108994))
								mBase = m.M
								v151 = m.ExcPending
								if v151 != 0 {
									return int64(0)
								} else {
									F_errmsg(m, int32(_a_F_sort_0), int32(0))
									mBase = m.M
									v155 = m.ExcPending
									if v155 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_sort_1), int32(207), int32(_a_F_sort_2))
										mBase = m.M
										v160 = m.ExcPending
										if v160 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v73 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
							v75 = v13 + int32(16)
							v76 = F_ArrayGetNItemsSafe(m, v73, v75)
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return int64(0)
							} else {
								if int32(2) <= v76 {
									v80 = int32(1)
									if v67 != 0 {
										v119 = v80
										v120 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
										v121 = F_ArrayGetNItemsSafe(m, v120, v75)
										mBase = m.M
										v122 = m.ExcPending
										if v122 != 0 {
											return int64(0)
										} else {
											*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v119)
											v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
											if v124 != 0 {
												v132 = v124
											} else {
												v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
												v132 = (v125<<(uint(int32(3))%32) + int32(23)) & int32(-8)
											}
											F_isort(m, v132+v13, v121, v10+int32(15))
											mBase = m.M
											m.G0 = v10 + int32(16)
											return base.I64_extend_i32_u(v13)
										}
									} else {
										switch v69 - int32(3) {
										case 0:
											v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
											if v83|int32(32) != int32(97) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v165 = m.ExcPending
												if v165 != 0 {
													return int64(0)
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v168 = m.ExcPending
													if v168 != 0 {
														return int64(0)
													} else {
														F_errmsg(m, int32(_a_F_sort_3), int32(0))
														mBase = m.M
														v172 = m.ExcPending
														if v172 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(_a_F_sort_1), int32(225), int32(_a_F_sort_2))
															mBase = m.M
															v177 = m.ExcPending
															if v177 != 0 {
																return int64(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
												if v88|int32(32) != int32(115) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v165 = m.ExcPending
													if v165 != 0 {
														return int64(0)
													} else {
														F_errcode(m, int32(50856066))
														mBase = m.M
														v168 = m.ExcPending
														if v168 != 0 {
															return int64(0)
														} else {
															F_errmsg(m, int32(_a_F_sort_3), int32(0))
															mBase = m.M
															v172 = m.ExcPending
															if v172 != 0 {
																return int64(0)
															} else {
																F_errfinish(m, int32(_a_F_sort_1), int32(225), int32(_a_F_sort_2))
																mBase = m.M
																v177 = m.ExcPending
																if v177 != 0 {
																	return int64(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													}
												} else {
													v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+2)))
													if v93|int32(32) == int32(99) {
														v119 = v80
														v120 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
														v121 = F_ArrayGetNItemsSafe(m, v120, v75)
														mBase = m.M
														v122 = m.ExcPending
														if v122 != 0 {
															return int64(0)
														} else {
															*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v119)
															v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
															if v124 != 0 {
																v132 = v124
															} else {
																v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
																v132 = (v125<<(uint(int32(3))%32) + int32(23)) & int32(-8)
															}
															F_isort(m, v132+v13, v121, v10+int32(15))
															mBase = m.M
															m.G0 = v10 + int32(16)
															return base.I64_extend_i32_u(v13)
														}
													} else {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v165 = m.ExcPending
														if v165 != 0 {
															return int64(0)
														} else {
															F_errcode(m, int32(50856066))
															mBase = m.M
															v168 = m.ExcPending
															if v168 != 0 {
																return int64(0)
															} else {
																F_errmsg(m, int32(_a_F_sort_3), int32(0))
																mBase = m.M
																v172 = m.ExcPending
																if v172 != 0 {
																	return int64(0)
																} else {
																	F_errfinish(m, int32(_a_F_sort_1), int32(225), int32(_a_F_sort_2))
																	mBase = m.M
																	v177 = m.ExcPending
																	if v177 != 0 {
																		return int64(0)
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
										case 1:
											v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
											if v98|int32(32) != int32(100) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v165 = m.ExcPending
												if v165 != 0 {
													return int64(0)
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v168 = m.ExcPending
													if v168 != 0 {
														return int64(0)
													} else {
														F_errmsg(m, int32(_a_F_sort_3), int32(0))
														mBase = m.M
														v172 = m.ExcPending
														if v172 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(_a_F_sort_1), int32(225), int32(_a_F_sort_2))
															mBase = m.M
															v177 = m.ExcPending
															if v177 != 0 {
																return int64(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
												if v103|int32(32) != int32(101) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v165 = m.ExcPending
													if v165 != 0 {
														return int64(0)
													} else {
														F_errcode(m, int32(50856066))
														mBase = m.M
														v168 = m.ExcPending
														if v168 != 0 {
															return int64(0)
														} else {
															F_errmsg(m, int32(_a_F_sort_3), int32(0))
															mBase = m.M
															v172 = m.ExcPending
															if v172 != 0 {
																return int64(0)
															} else {
																F_errfinish(m, int32(_a_F_sort_1), int32(225), int32(_a_F_sort_2))
																mBase = m.M
																v177 = m.ExcPending
																if v177 != 0 {
																	return int64(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													}
												} else {
													v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+2)))
													if v108|int32(32) != int32(115) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v165 = m.ExcPending
														if v165 != 0 {
															return int64(0)
														} else {
															F_errcode(m, int32(50856066))
															mBase = m.M
															v168 = m.ExcPending
															if v168 != 0 {
																return int64(0)
															} else {
																F_errmsg(m, int32(_a_F_sort_3), int32(0))
																mBase = m.M
																v172 = m.ExcPending
																if v172 != 0 {
																	return int64(0)
																} else {
																	F_errfinish(m, int32(_a_F_sort_1), int32(225), int32(_a_F_sort_2))
																	mBase = m.M
																	v177 = m.ExcPending
																	if v177 != 0 {
																		return int64(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														}
													} else {
														v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+3)))
														if v114|int32(32) != int32(99) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v165 = m.ExcPending
															if v165 != 0 {
																return int64(0)
															} else {
																F_errcode(m, int32(50856066))
																mBase = m.M
																v168 = m.ExcPending
																if v168 != 0 {
																	return int64(0)
																} else {
																	F_errmsg(m, int32(_a_F_sort_3), int32(0))
																	mBase = m.M
																	v172 = m.ExcPending
																	if v172 != 0 {
																		return int64(0)
																	} else {
																		F_errfinish(m, int32(_a_F_sort_1), int32(225), int32(_a_F_sort_2))
																		mBase = m.M
																		v177 = m.ExcPending
																		if v177 != 0 {
																			return int64(0)
																		} else {
																			base.Wasm_trap_unreachable()
																			for {
																			}
																		}
																	}
																}
															}
														} else {
															v119 = int32(0)
															v120 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
															v121 = F_ArrayGetNItemsSafe(m, v120, v75)
															mBase = m.M
															v122 = m.ExcPending
															if v122 != 0 {
																return int64(0)
															} else {
																*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v119)
																v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
																if v124 != 0 {
																	v132 = v124
																} else {
																	v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
																	v132 = (v125<<(uint(int32(3))%32) + int32(23)) & int32(-8)
																}
																F_isort(m, v132+v13, v121, v10+int32(15))
																mBase = m.M
																m.G0 = v10 + int32(16)
																return base.I64_extend_i32_u(v13)
															}
														}
													}
												}
											}
										default:
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v165 = m.ExcPending
											if v165 != 0 {
												return int64(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v168 = m.ExcPending
												if v168 != 0 {
													return int64(0)
												} else {
													F_errmsg(m, int32(_a_F_sort_3), int32(0))
													mBase = m.M
													v172 = m.ExcPending
													if v172 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_sort_1), int32(225), int32(_a_F_sort_2))
														mBase = m.M
														v177 = m.ExcPending
														if v177 != 0 {
															return int64(0)
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
									m.G0 = v10 + int32(16)
									return base.I64_extend_i32_u(v13)
								}
							}
						}
					}
				} else {
					v73 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
					v75 = v13 + int32(16)
					v76 = F_ArrayGetNItemsSafe(m, v73, v75)
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return int64(0)
					} else {
						if int32(2) <= v76 {
							v80 = int32(1)
							if v67 != 0 {
								v119 = v80
								v120 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
								v121 = F_ArrayGetNItemsSafe(m, v120, v75)
								mBase = m.M
								v122 = m.ExcPending
								if v122 != 0 {
									return int64(0)
								} else {
									*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v119)
									v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
									if v124 != 0 {
										v132 = v124
									} else {
										v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
										v132 = (v125<<(uint(int32(3))%32) + int32(23)) & int32(-8)
									}
									F_isort(m, v132+v13, v121, v10+int32(15))
									mBase = m.M
									m.G0 = v10 + int32(16)
									return base.I64_extend_i32_u(v13)
								}
							} else {
								switch v69 - int32(3) {
								case 0:
									v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
									if v83|int32(32) != int32(97) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v165 = m.ExcPending
										if v165 != 0 {
											return int64(0)
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v168 = m.ExcPending
											if v168 != 0 {
												return int64(0)
											} else {
												F_errmsg(m, int32(_a_F_sort_3), int32(0))
												mBase = m.M
												v172 = m.ExcPending
												if v172 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(_a_F_sort_1), int32(225), int32(_a_F_sort_2))
													mBase = m.M
													v177 = m.ExcPending
													if v177 != 0 {
														return int64(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									} else {
										v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
										if v88|int32(32) != int32(115) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v165 = m.ExcPending
											if v165 != 0 {
												return int64(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v168 = m.ExcPending
												if v168 != 0 {
													return int64(0)
												} else {
													F_errmsg(m, int32(_a_F_sort_3), int32(0))
													mBase = m.M
													v172 = m.ExcPending
													if v172 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_sort_1), int32(225), int32(_a_F_sort_2))
														mBase = m.M
														v177 = m.ExcPending
														if v177 != 0 {
															return int64(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										} else {
											v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+2)))
											if v93|int32(32) == int32(99) {
												v119 = v80
												v120 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
												v121 = F_ArrayGetNItemsSafe(m, v120, v75)
												mBase = m.M
												v122 = m.ExcPending
												if v122 != 0 {
													return int64(0)
												} else {
													*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v119)
													v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
													if v124 != 0 {
														v132 = v124
													} else {
														v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
														v132 = (v125<<(uint(int32(3))%32) + int32(23)) & int32(-8)
													}
													F_isort(m, v132+v13, v121, v10+int32(15))
													mBase = m.M
													m.G0 = v10 + int32(16)
													return base.I64_extend_i32_u(v13)
												}
											} else {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v165 = m.ExcPending
												if v165 != 0 {
													return int64(0)
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v168 = m.ExcPending
													if v168 != 0 {
														return int64(0)
													} else {
														F_errmsg(m, int32(_a_F_sort_3), int32(0))
														mBase = m.M
														v172 = m.ExcPending
														if v172 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(_a_F_sort_1), int32(225), int32(_a_F_sort_2))
															mBase = m.M
															v177 = m.ExcPending
															if v177 != 0 {
																return int64(0)
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
								case 1:
									v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
									if v98|int32(32) != int32(100) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v165 = m.ExcPending
										if v165 != 0 {
											return int64(0)
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v168 = m.ExcPending
											if v168 != 0 {
												return int64(0)
											} else {
												F_errmsg(m, int32(_a_F_sort_3), int32(0))
												mBase = m.M
												v172 = m.ExcPending
												if v172 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(_a_F_sort_1), int32(225), int32(_a_F_sort_2))
													mBase = m.M
													v177 = m.ExcPending
													if v177 != 0 {
														return int64(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									} else {
										v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
										if v103|int32(32) != int32(101) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v165 = m.ExcPending
											if v165 != 0 {
												return int64(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v168 = m.ExcPending
												if v168 != 0 {
													return int64(0)
												} else {
													F_errmsg(m, int32(_a_F_sort_3), int32(0))
													mBase = m.M
													v172 = m.ExcPending
													if v172 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_sort_1), int32(225), int32(_a_F_sort_2))
														mBase = m.M
														v177 = m.ExcPending
														if v177 != 0 {
															return int64(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										} else {
											v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+2)))
											if v108|int32(32) != int32(115) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v165 = m.ExcPending
												if v165 != 0 {
													return int64(0)
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v168 = m.ExcPending
													if v168 != 0 {
														return int64(0)
													} else {
														F_errmsg(m, int32(_a_F_sort_3), int32(0))
														mBase = m.M
														v172 = m.ExcPending
														if v172 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(_a_F_sort_1), int32(225), int32(_a_F_sort_2))
															mBase = m.M
															v177 = m.ExcPending
															if v177 != 0 {
																return int64(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+3)))
												if v114|int32(32) != int32(99) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v165 = m.ExcPending
													if v165 != 0 {
														return int64(0)
													} else {
														F_errcode(m, int32(50856066))
														mBase = m.M
														v168 = m.ExcPending
														if v168 != 0 {
															return int64(0)
														} else {
															F_errmsg(m, int32(_a_F_sort_3), int32(0))
															mBase = m.M
															v172 = m.ExcPending
															if v172 != 0 {
																return int64(0)
															} else {
																F_errfinish(m, int32(_a_F_sort_1), int32(225), int32(_a_F_sort_2))
																mBase = m.M
																v177 = m.ExcPending
																if v177 != 0 {
																	return int64(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													}
												} else {
													v119 = int32(0)
													v120 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
													v121 = F_ArrayGetNItemsSafe(m, v120, v75)
													mBase = m.M
													v122 = m.ExcPending
													if v122 != 0 {
														return int64(0)
													} else {
														*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v119)
														v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
														if v124 != 0 {
															v132 = v124
														} else {
															v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
															v132 = (v125<<(uint(int32(3))%32) + int32(23)) & int32(-8)
														}
														F_isort(m, v132+v13, v121, v10+int32(15))
														mBase = m.M
														m.G0 = v10 + int32(16)
														return base.I64_extend_i32_u(v13)
													}
												}
											}
										}
									}
								default:
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v165 = m.ExcPending
									if v165 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v168 = m.ExcPending
										if v168 != 0 {
											return int64(0)
										} else {
											F_errmsg(m, int32(_a_F_sort_3), int32(0))
											mBase = m.M
											v172 = m.ExcPending
											if v172 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_sort_1), int32(225), int32(_a_F_sort_2))
												mBase = m.M
												v177 = m.ExcPending
												if v177 != 0 {
													return int64(0)
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
							m.G0 = v10 + int32(16)
							return base.I64_extend_i32_u(v13)
						}
					}
				}
			}
		}
	}
}
func F_sort_pending_writebacks(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v625 int32
	_ = v625
	var v627 int64
	_ = v627
	var v629 int64
	_ = v629
	var v631 int32
	_ = v631
	var v633 int64
	_ = v633
	var v635 int64
	_ = v635
	var v637 int32
	_ = v637
	var v639 int64
	_ = v639
	var v641 int64
	_ = v641
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v694 int64
	_ = v694
	var v696 int64
	_ = v696
	var v698 int32
	_ = v698
	var v700 int64
	_ = v700
	var v702 int64
	_ = v702
	var v704 int32
	_ = v704
	var v706 int64
	_ = v706
	var v708 int64
	_ = v708
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v734 int32
	_ = v734
	var v739 int32
	_ = v739
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
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
	var v767 int64
	_ = v767
	var v769 int64
	_ = v769
	var v771 int32
	_ = v771
	var v773 int64
	_ = v773
	var v775 int64
	_ = v775
	var v777 int32
	_ = v777
	var v779 int64
	_ = v779
	var v781 int64
	_ = v781
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v793 int32
	_ = v793
	var v798 int32
	_ = v798
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v821 int32
	_ = v821
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v834 int64
	_ = v834
	var v836 int64
	_ = v836
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v841 int64
	_ = v841
	var v843 int64
	_ = v843
	var v845 int32
	_ = v845
	var v847 int64
	_ = v847
	var v849 int64
	_ = v849
	var v852 int32
	_ = v852
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v887 int32
	_ = v887
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v899 int64
	_ = v899
	var v901 int64
	_ = v901
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v906 int64
	_ = v906
	var v908 int64
	_ = v908
	var v910 int32
	_ = v910
	var v912 int64
	_ = v912
	var v914 int64
	_ = v914
	var v917 int32
	_ = v917
	var v936 int32
	_ = v936
	var v947 int32
	_ = v947
	var v949 int64
	_ = v949
	var v951 int64
	_ = v951
	var v953 int32
	_ = v953
	var v955 int64
	_ = v955
	var v957 int64
	_ = v957
	var v959 int32
	_ = v959
	var v961 int64
	_ = v961
	var v963 int64
	_ = v963
	var v965 int32
	_ = v965
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v982 int32
	_ = v982
	var v984 int32
	_ = v984
	var v997 int32
	_ = v997
	var v1007 int32
	_ = v1007
	var v1019 int32
	_ = v1019
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1035 int32
	_ = v1035
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1060 int32
	_ = v1060
	var v1062 int64
	_ = v1062
	var v1064 int64
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1068 int64
	_ = v1068
	var v1070 int64
	_ = v1070
	var v1072 int32
	_ = v1072
	var v1074 int64
	_ = v1074
	var v1076 int64
	_ = v1076
	var v1093 int32
	_ = v1093
	v14 = m.G0
	v16 = v14 - int32(32)
	m.G0 = v16
	if base.Ui32(l1) < base.Ui32(int32(7)) {
		v969 = l0
		v970 = l1
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v16 + int32(32)
	return
L2:
	;
	if base.Ui32(v984) < base.Ui32(int32(2)) {
		goto L1
	} else {
		goto L318
	}
L3:
	;
	v982 = v969
	v984 = v970
	goto L2
L4:
	;
	v20 = l0
	v21 = l1
	goto L5
L5:
	;
	v34 = v20 + int32(20)
	v36 = v21
	goto L7
L7:
	;
	if base.Ui32(v36) < base.Ui32(int32(2)) {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v52 = v20 + v36*int32(20)
	v55 = v34
	goto L10
L10:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v55-int32(12))))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	if base.Ui32(v68) < base.Ui32(v69) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v106 = v20 + int32(base.Ui32(v36)>>(uint(int32(1))%32))*int32(20)
	if v36 != int32(7) {
		goto L24
	} else {
		goto L25
	}
L12:
	;
	goto L11
L13:
	;
	v98 = v55 + int32(20)
	if base.Ui32(v98) < base.Ui32(v52) {
		v55 = v98
		goto L10
	} else {
		goto L23
	}
L14:
	;
	if base.Ui32(v69) < base.Ui32(v68) {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v55-int32(16))))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if base.Ui32(v74) < base.Ui32(v75) {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	if base.Ui32(v75) < base.Ui32(v74) {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v55-int32(20))))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	if base.Ui32(v80) < base.Ui32(v81) {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	if base.Ui32(v81) < base.Ui32(v80) {
		goto L12
	} else {
		goto L19
	}
L19:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v55-int32(8))))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v55)+12))
	if v86 < v87 {
		goto L13
	} else {
		goto L20
	}
L20:
	;
	if v87 < v86 {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v55-int32(4))))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
	if base.Ui32(v93) < base.Ui32(v92) {
		goto L12
	} else {
		goto L22
	}
L22:
	;
	goto L13
L23:
	;
	goto L1
L24:
	;
	v110 = v52 - int32(20)
	if base.Ui32(v36) < base.Ui32(int32(41)) {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	v620 = v106
	goto L26
L26:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v625
	v627 = *(*int64)(unsafe.Add(mBase, uint32(v20)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v627
	v629 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v629
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v620)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v631
	v633 = *(*int64)(unsafe.Add(mBase, uint32(v620)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+8)) = v633
	v635 = *(*int64)(unsafe.Add(mBase, uint32(v620)))
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v635
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v620)+16)) = v637
	v639 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v620)+8)) = v639
	v641 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v620))) = v641
	v644 = v52 - int32(20)
	v647 = v644
	v648 = v34
	v650 = v34
	v652 = v644
	goto L259
L27:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v493)+4))
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v493)))
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v492)+4))
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v492)))
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v492)+8))
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v493)+8))
	if base.Ui32(v511) < base.Ui32(v512) {
		goto L207
	} else {
		goto L208
	}
L28:
	;
	v492 = v20
	v493 = v106
	v494 = v110
	goto L27
L29:
	;
	goto L30
L30:
	;
	v114 = int32(base.Ui32(v36) >> (uint(int32(3)) % 32))
	v116 = v114 * int32(20)
	v117 = v20 + v116
	v120 = v20 + v114*int32(40)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v117)+8))
	if base.Ui32(v134) < base.Ui32(v135) {
		goto L36
	} else {
		goto L37
	}
L31:
	;
	v243 = v114 * int32(-20)
	v244 = v106 + v243
	v245 = v106 + v116
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v244)))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v244)+8))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v106)+8))
	if base.Ui32(v259) < base.Ui32(v260) {
		goto L93
	} else {
		goto L94
	}
L32:
	;
	v241 = v20
	goto L31
L33:
	;
	v241 = v120
	goto L31
L34:
	;
	v241 = v217
	goto L31
L35:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
	if base.Ui32(v135) < base.Ui32(v187) {
		goto L67
	} else {
		goto L68
	}
L36:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
	if base.Ui32(v135) < base.Ui32(v151) {
		v217 = v117
		goto L34
	} else {
		goto L46
	}
L37:
	;
	if base.Ui32(v135) < base.Ui32(v134) {
		goto L35
	} else {
		goto L38
	}
L38:
	;
	if base.Ui32(v132) < base.Ui32(v130) {
		goto L36
	} else {
		goto L39
	}
L39:
	;
	if base.Ui32(v130) < base.Ui32(v132) {
		goto L35
	} else {
		goto L40
	}
L40:
	;
	if base.Ui32(v133) < base.Ui32(v131) {
		goto L36
	} else {
		goto L41
	}
L41:
	;
	if base.Ui32(v131) < base.Ui32(v133) {
		goto L35
	} else {
		goto L42
	}
L42:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v117)+12))
	if v142 < v143 {
		goto L36
	} else {
		goto L43
	}
L43:
	;
	if v143 < v142 {
		goto L35
	} else {
		goto L44
	}
L44:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v117)+16))
	if base.Ui32(v147) <= base.Ui32(v146) {
		goto L35
	} else {
		goto L45
	}
L45:
	;
	goto L36
L46:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	if base.Ui32(v151) < base.Ui32(v135) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	if base.Ui32(v134) < base.Ui32(v151) {
		goto L33
	} else {
		goto L56
	}
L48:
	;
	if base.Ui32(v130) < base.Ui32(v153) {
		v217 = v117
		goto L34
	} else {
		goto L49
	}
L49:
	;
	if base.Ui32(v153) < base.Ui32(v130) {
		goto L47
	} else {
		goto L50
	}
L50:
	;
	if base.Ui32(v131) < base.Ui32(v154) {
		v217 = v117
		goto L34
	} else {
		goto L51
	}
L51:
	;
	if base.Ui32(v154) < base.Ui32(v131) {
		goto L47
	} else {
		goto L52
	}
L52:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v117)+12))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	if v160 < v161 {
		v217 = v117
		goto L34
	} else {
		goto L53
	}
L53:
	;
	if v161 < v160 {
		goto L47
	} else {
		goto L54
	}
L54:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v117)+16))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v120)+16))
	if base.Ui32(v164) < base.Ui32(v165) {
		v217 = v117
		goto L34
	} else {
		goto L55
	}
L55:
	;
	goto L47
L56:
	;
	if base.Ui32(v151) < base.Ui32(v134) {
		goto L32
	} else {
		goto L57
	}
L57:
	;
	if base.Ui32(v132) < base.Ui32(v153) {
		goto L33
	} else {
		goto L58
	}
L58:
	;
	if base.Ui32(v153) < base.Ui32(v132) {
		goto L32
	} else {
		goto L59
	}
L59:
	;
	if base.Ui32(v133) < base.Ui32(v154) {
		goto L33
	} else {
		goto L60
	}
L60:
	;
	if base.Ui32(v154) < base.Ui32(v133) {
		goto L32
	} else {
		goto L61
	}
L61:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	if v175 < v176 {
		goto L33
	} else {
		goto L62
	}
L62:
	;
	if v176 < v175 {
		goto L32
	} else {
		goto L63
	}
L63:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v120)+16))
	if base.Ui32(v179) < base.Ui32(v180) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v182 = v120
	goto L66
L65:
	;
	v182 = v20
	goto L66
L66:
	;
	v241 = v182
	goto L31
L67:
	;
	if base.Ui32(v134) < base.Ui32(v187) {
		goto L32
	} else {
		goto L77
	}
L68:
	;
	if base.Ui32(v187) < base.Ui32(v135) {
		v217 = v117
		goto L34
	} else {
		goto L69
	}
L69:
	;
	if base.Ui32(v130) < base.Ui32(v185) {
		goto L67
	} else {
		goto L70
	}
L70:
	;
	if base.Ui32(v185) < base.Ui32(v130) {
		v217 = v117
		goto L34
	} else {
		goto L71
	}
L71:
	;
	if base.Ui32(v131) < base.Ui32(v186) {
		goto L67
	} else {
		goto L72
	}
L72:
	;
	if base.Ui32(v186) < base.Ui32(v131) {
		v217 = v117
		goto L34
	} else {
		goto L73
	}
L73:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v117)+12))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	if v194 < v195 {
		goto L67
	} else {
		goto L74
	}
L74:
	;
	if v195 < v194 {
		v217 = v117
		goto L34
	} else {
		goto L75
	}
L75:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v117)+16))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v120)+16))
	if base.Ui32(v199) < base.Ui32(v198) {
		v217 = v117
		goto L34
	} else {
		goto L76
	}
L76:
	;
	goto L67
L77:
	;
	if base.Ui32(v187) < base.Ui32(v134) {
		goto L33
	} else {
		goto L78
	}
L78:
	;
	if base.Ui32(v132) < base.Ui32(v185) {
		goto L32
	} else {
		goto L79
	}
L79:
	;
	if base.Ui32(v185) < base.Ui32(v132) {
		goto L33
	} else {
		goto L80
	}
L80:
	;
	if base.Ui32(v133) < base.Ui32(v186) {
		goto L32
	} else {
		goto L81
	}
L81:
	;
	if base.Ui32(v186) < base.Ui32(v133) {
		goto L33
	} else {
		goto L82
	}
L82:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	if v209 < v210 {
		goto L32
	} else {
		goto L83
	}
L83:
	;
	if v210 < v209 {
		goto L33
	} else {
		goto L84
	}
L84:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v120)+16))
	if base.Ui32(v213) < base.Ui32(v214) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v216 = v20
	goto L87
L86:
	;
	v216 = v120
	goto L87
L87:
	;
	v217 = v216
	goto L34
L88:
	;
	v369 = v110 + v114*int32(-40)
	v370 = v110 + v243
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v370)+4))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v370)))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v369)+4))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v369)))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v369)+8))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v370)+8))
	if base.Ui32(v384) < base.Ui32(v385) {
		goto L150
	} else {
		goto L151
	}
L89:
	;
	v366 = v244
	goto L88
L90:
	;
	v366 = v245
	goto L88
L91:
	;
	v366 = v342
	goto L88
L92:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v245)+4))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v245)))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v245)+8))
	if base.Ui32(v260) < base.Ui32(v312) {
		goto L124
	} else {
		goto L125
	}
L93:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v245)+8))
	if base.Ui32(v260) < base.Ui32(v276) {
		v342 = v106
		goto L91
	} else {
		goto L103
	}
L94:
	;
	if base.Ui32(v260) < base.Ui32(v259) {
		goto L92
	} else {
		goto L95
	}
L95:
	;
	if base.Ui32(v257) < base.Ui32(v255) {
		goto L93
	} else {
		goto L96
	}
L96:
	;
	if base.Ui32(v255) < base.Ui32(v257) {
		goto L92
	} else {
		goto L97
	}
L97:
	;
	if base.Ui32(v258) < base.Ui32(v256) {
		goto L93
	} else {
		goto L98
	}
L98:
	;
	if base.Ui32(v256) < base.Ui32(v258) {
		goto L92
	} else {
		goto L99
	}
L99:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v244)+12))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v106)+12))
	if v267 < v268 {
		goto L93
	} else {
		goto L100
	}
L100:
	;
	if v268 < v267 {
		goto L92
	} else {
		goto L101
	}
L101:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v244)+16))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v106)+16))
	if base.Ui32(v272) <= base.Ui32(v271) {
		goto L92
	} else {
		goto L102
	}
L102:
	;
	goto L93
L103:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v245)+4))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v245)))
	if base.Ui32(v276) < base.Ui32(v260) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	if base.Ui32(v259) < base.Ui32(v276) {
		goto L90
	} else {
		goto L113
	}
L105:
	;
	if base.Ui32(v255) < base.Ui32(v278) {
		v342 = v106
		goto L91
	} else {
		goto L106
	}
L106:
	;
	if base.Ui32(v278) < base.Ui32(v255) {
		goto L104
	} else {
		goto L107
	}
L107:
	;
	if base.Ui32(v256) < base.Ui32(v279) {
		v342 = v106
		goto L91
	} else {
		goto L108
	}
L108:
	;
	if base.Ui32(v279) < base.Ui32(v256) {
		goto L104
	} else {
		goto L109
	}
L109:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v106)+12))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v245)+12))
	if v285 < v286 {
		v342 = v106
		goto L91
	} else {
		goto L110
	}
L110:
	;
	if v286 < v285 {
		goto L104
	} else {
		goto L111
	}
L111:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v106)+16))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v245)+16))
	if base.Ui32(v289) < base.Ui32(v290) {
		v342 = v106
		goto L91
	} else {
		goto L112
	}
L112:
	;
	goto L104
L113:
	;
	if base.Ui32(v276) < base.Ui32(v259) {
		goto L89
	} else {
		goto L114
	}
L114:
	;
	if base.Ui32(v257) < base.Ui32(v278) {
		goto L90
	} else {
		goto L115
	}
L115:
	;
	if base.Ui32(v278) < base.Ui32(v257) {
		goto L89
	} else {
		goto L116
	}
L116:
	;
	if base.Ui32(v258) < base.Ui32(v279) {
		goto L90
	} else {
		goto L117
	}
L117:
	;
	if base.Ui32(v279) < base.Ui32(v258) {
		goto L89
	} else {
		goto L118
	}
L118:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v244)+12))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v245)+12))
	if v300 < v301 {
		goto L90
	} else {
		goto L119
	}
L119:
	;
	if v301 < v300 {
		goto L89
	} else {
		goto L120
	}
L120:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v244)+16))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v245)+16))
	if base.Ui32(v304) < base.Ui32(v305) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v307 = v245
	goto L123
L122:
	;
	v307 = v244
	goto L123
L123:
	;
	v366 = v307
	goto L88
L124:
	;
	if base.Ui32(v259) < base.Ui32(v312) {
		goto L89
	} else {
		goto L134
	}
L125:
	;
	if base.Ui32(v312) < base.Ui32(v260) {
		v342 = v106
		goto L91
	} else {
		goto L126
	}
L126:
	;
	if base.Ui32(v255) < base.Ui32(v310) {
		goto L124
	} else {
		goto L127
	}
L127:
	;
	if base.Ui32(v310) < base.Ui32(v255) {
		v342 = v106
		goto L91
	} else {
		goto L128
	}
L128:
	;
	if base.Ui32(v256) < base.Ui32(v311) {
		goto L124
	} else {
		goto L129
	}
L129:
	;
	if base.Ui32(v311) < base.Ui32(v256) {
		v342 = v106
		goto L91
	} else {
		goto L130
	}
L130:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v106)+12))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v245)+12))
	if v319 < v320 {
		goto L124
	} else {
		goto L131
	}
L131:
	;
	if v320 < v319 {
		v342 = v106
		goto L91
	} else {
		goto L132
	}
L132:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v106)+16))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v245)+16))
	if base.Ui32(v324) < base.Ui32(v323) {
		v342 = v106
		goto L91
	} else {
		goto L133
	}
L133:
	;
	goto L124
L134:
	;
	if base.Ui32(v312) < base.Ui32(v259) {
		goto L90
	} else {
		goto L135
	}
L135:
	;
	if base.Ui32(v257) < base.Ui32(v310) {
		goto L89
	} else {
		goto L136
	}
L136:
	;
	if base.Ui32(v310) < base.Ui32(v257) {
		goto L90
	} else {
		goto L137
	}
L137:
	;
	if base.Ui32(v258) < base.Ui32(v311) {
		goto L89
	} else {
		goto L138
	}
L138:
	;
	if base.Ui32(v311) < base.Ui32(v258) {
		goto L90
	} else {
		goto L139
	}
L139:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v244)+12))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v245)+12))
	if v334 < v335 {
		goto L89
	} else {
		goto L140
	}
L140:
	;
	if v335 < v334 {
		goto L90
	} else {
		goto L141
	}
L141:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v244)+16))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v245)+16))
	if base.Ui32(v338) < base.Ui32(v339) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v341 = v244
	goto L144
L143:
	;
	v341 = v245
	goto L144
L144:
	;
	v342 = v341
	goto L91
L145:
	;
	v492 = v241
	v493 = v366
	v494 = v491
	goto L27
L146:
	;
	v491 = v369
	goto L145
L147:
	;
	v491 = v110
	goto L145
L148:
	;
	v491 = v467
	goto L145
L149:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v110)+8))
	if base.Ui32(v385) < base.Ui32(v437) {
		goto L181
	} else {
		goto L182
	}
L150:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v110)+8))
	if base.Ui32(v385) < base.Ui32(v401) {
		v467 = v370
		goto L148
	} else {
		goto L160
	}
L151:
	;
	if base.Ui32(v385) < base.Ui32(v384) {
		goto L149
	} else {
		goto L152
	}
L152:
	;
	if base.Ui32(v382) < base.Ui32(v380) {
		goto L150
	} else {
		goto L153
	}
L153:
	;
	if base.Ui32(v380) < base.Ui32(v382) {
		goto L149
	} else {
		goto L154
	}
L154:
	;
	if base.Ui32(v383) < base.Ui32(v381) {
		goto L150
	} else {
		goto L155
	}
L155:
	;
	if base.Ui32(v381) < base.Ui32(v383) {
		goto L149
	} else {
		goto L156
	}
L156:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v369)+12))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v370)+12))
	if v392 < v393 {
		goto L150
	} else {
		goto L157
	}
L157:
	;
	if v393 < v392 {
		goto L149
	} else {
		goto L158
	}
L158:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v369)+16))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v370)+16))
	if base.Ui32(v397) <= base.Ui32(v396) {
		goto L149
	} else {
		goto L159
	}
L159:
	;
	goto L150
L160:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	if base.Ui32(v401) < base.Ui32(v385) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	if base.Ui32(v384) < base.Ui32(v401) {
		goto L147
	} else {
		goto L170
	}
L162:
	;
	if base.Ui32(v380) < base.Ui32(v403) {
		v467 = v370
		goto L148
	} else {
		goto L163
	}
L163:
	;
	if base.Ui32(v403) < base.Ui32(v380) {
		goto L161
	} else {
		goto L164
	}
L164:
	;
	if base.Ui32(v381) < base.Ui32(v404) {
		v467 = v370
		goto L148
	} else {
		goto L165
	}
L165:
	;
	if base.Ui32(v404) < base.Ui32(v381) {
		goto L161
	} else {
		goto L166
	}
L166:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v370)+12))
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v110)+12))
	if v410 < v411 {
		v467 = v370
		goto L148
	} else {
		goto L167
	}
L167:
	;
	if v411 < v410 {
		goto L161
	} else {
		goto L168
	}
L168:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v370)+16))
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v110)+16))
	if base.Ui32(v414) < base.Ui32(v415) {
		v467 = v370
		goto L148
	} else {
		goto L169
	}
L169:
	;
	goto L161
L170:
	;
	if base.Ui32(v401) < base.Ui32(v384) {
		goto L146
	} else {
		goto L171
	}
L171:
	;
	if base.Ui32(v382) < base.Ui32(v403) {
		goto L147
	} else {
		goto L172
	}
L172:
	;
	if base.Ui32(v403) < base.Ui32(v382) {
		goto L146
	} else {
		goto L173
	}
L173:
	;
	if base.Ui32(v383) < base.Ui32(v404) {
		goto L147
	} else {
		goto L174
	}
L174:
	;
	if base.Ui32(v404) < base.Ui32(v383) {
		goto L146
	} else {
		goto L175
	}
L175:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v369)+12))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v110)+12))
	if v425 < v426 {
		goto L147
	} else {
		goto L176
	}
L176:
	;
	if v426 < v425 {
		goto L146
	} else {
		goto L177
	}
L177:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v369)+16))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v110)+16))
	if base.Ui32(v429) < base.Ui32(v430) {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v432 = v110
	goto L180
L179:
	;
	v432 = v369
	goto L180
L180:
	;
	v491 = v432
	goto L145
L181:
	;
	if base.Ui32(v384) < base.Ui32(v437) {
		goto L146
	} else {
		goto L191
	}
L182:
	;
	if base.Ui32(v437) < base.Ui32(v385) {
		v467 = v370
		goto L148
	} else {
		goto L183
	}
L183:
	;
	if base.Ui32(v380) < base.Ui32(v435) {
		goto L181
	} else {
		goto L184
	}
L184:
	;
	if base.Ui32(v435) < base.Ui32(v380) {
		v467 = v370
		goto L148
	} else {
		goto L185
	}
L185:
	;
	if base.Ui32(v381) < base.Ui32(v436) {
		goto L181
	} else {
		goto L186
	}
L186:
	;
	if base.Ui32(v436) < base.Ui32(v381) {
		v467 = v370
		goto L148
	} else {
		goto L187
	}
L187:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v370)+12))
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v110)+12))
	if v444 < v445 {
		goto L181
	} else {
		goto L188
	}
L188:
	;
	if v445 < v444 {
		v467 = v370
		goto L148
	} else {
		goto L189
	}
L189:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v370)+16))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v110)+16))
	if base.Ui32(v449) < base.Ui32(v448) {
		v467 = v370
		goto L148
	} else {
		goto L190
	}
L190:
	;
	goto L181
L191:
	;
	if base.Ui32(v437) < base.Ui32(v384) {
		goto L147
	} else {
		goto L192
	}
L192:
	;
	if base.Ui32(v382) < base.Ui32(v435) {
		goto L146
	} else {
		goto L193
	}
L193:
	;
	if base.Ui32(v435) < base.Ui32(v382) {
		goto L147
	} else {
		goto L194
	}
L194:
	;
	if base.Ui32(v383) < base.Ui32(v436) {
		goto L146
	} else {
		goto L195
	}
L195:
	;
	if base.Ui32(v436) < base.Ui32(v383) {
		goto L147
	} else {
		goto L196
	}
L196:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v369)+12))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v110)+12))
	if v459 < v460 {
		goto L146
	} else {
		goto L197
	}
L197:
	;
	if v460 < v459 {
		goto L147
	} else {
		goto L198
	}
L198:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v369)+16))
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v110)+16))
	if base.Ui32(v463) < base.Ui32(v464) {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v466 = v369
	goto L201
L200:
	;
	v466 = v110
	goto L201
L201:
	;
	v467 = v466
	goto L148
L202:
	;
	v620 = v618
	goto L26
L203:
	;
	v618 = v492
	goto L202
L204:
	;
	v618 = v494
	goto L202
L205:
	;
	v618 = v594
	goto L202
L206:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v494)+4))
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v494)))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v494)+8))
	if base.Ui32(v512) < base.Ui32(v564) {
		goto L238
	} else {
		goto L239
	}
L207:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v494)+8))
	if base.Ui32(v512) < base.Ui32(v528) {
		v594 = v493
		goto L205
	} else {
		goto L217
	}
L208:
	;
	if base.Ui32(v512) < base.Ui32(v511) {
		goto L206
	} else {
		goto L209
	}
L209:
	;
	if base.Ui32(v509) < base.Ui32(v507) {
		goto L207
	} else {
		goto L210
	}
L210:
	;
	if base.Ui32(v507) < base.Ui32(v509) {
		goto L206
	} else {
		goto L211
	}
L211:
	;
	if base.Ui32(v510) < base.Ui32(v508) {
		goto L207
	} else {
		goto L212
	}
L212:
	;
	if base.Ui32(v508) < base.Ui32(v510) {
		goto L206
	} else {
		goto L213
	}
L213:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v492)+12))
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v493)+12))
	if v519 < v520 {
		goto L207
	} else {
		goto L214
	}
L214:
	;
	if v520 < v519 {
		goto L206
	} else {
		goto L215
	}
L215:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v492)+16))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v493)+16))
	if base.Ui32(v524) <= base.Ui32(v523) {
		goto L206
	} else {
		goto L216
	}
L216:
	;
	goto L207
L217:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v494)+4))
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v494)))
	if base.Ui32(v528) < base.Ui32(v512) {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	if base.Ui32(v511) < base.Ui32(v528) {
		goto L204
	} else {
		goto L227
	}
L219:
	;
	if base.Ui32(v507) < base.Ui32(v530) {
		v594 = v493
		goto L205
	} else {
		goto L220
	}
L220:
	;
	if base.Ui32(v530) < base.Ui32(v507) {
		goto L218
	} else {
		goto L221
	}
L221:
	;
	if base.Ui32(v508) < base.Ui32(v531) {
		v594 = v493
		goto L205
	} else {
		goto L222
	}
L222:
	;
	if base.Ui32(v531) < base.Ui32(v508) {
		goto L218
	} else {
		goto L223
	}
L223:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v493)+12))
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v494)+12))
	if v537 < v538 {
		v594 = v493
		goto L205
	} else {
		goto L224
	}
L224:
	;
	if v538 < v537 {
		goto L218
	} else {
		goto L225
	}
L225:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v493)+16))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v494)+16))
	if base.Ui32(v541) < base.Ui32(v542) {
		v594 = v493
		goto L205
	} else {
		goto L226
	}
L226:
	;
	goto L218
L227:
	;
	if base.Ui32(v528) < base.Ui32(v511) {
		goto L203
	} else {
		goto L228
	}
L228:
	;
	if base.Ui32(v509) < base.Ui32(v530) {
		goto L204
	} else {
		goto L229
	}
L229:
	;
	if base.Ui32(v530) < base.Ui32(v509) {
		goto L203
	} else {
		goto L230
	}
L230:
	;
	if base.Ui32(v510) < base.Ui32(v531) {
		goto L204
	} else {
		goto L231
	}
L231:
	;
	if base.Ui32(v531) < base.Ui32(v510) {
		goto L203
	} else {
		goto L232
	}
L232:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v492)+12))
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v494)+12))
	if v552 < v553 {
		goto L204
	} else {
		goto L233
	}
L233:
	;
	if v553 < v552 {
		goto L203
	} else {
		goto L234
	}
L234:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v492)+16))
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v494)+16))
	if base.Ui32(v556) < base.Ui32(v557) {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v559 = v494
	goto L237
L236:
	;
	v559 = v492
	goto L237
L237:
	;
	v618 = v559
	goto L202
L238:
	;
	if base.Ui32(v511) < base.Ui32(v564) {
		goto L203
	} else {
		goto L248
	}
L239:
	;
	if base.Ui32(v564) < base.Ui32(v512) {
		v594 = v493
		goto L205
	} else {
		goto L240
	}
L240:
	;
	if base.Ui32(v507) < base.Ui32(v562) {
		goto L238
	} else {
		goto L241
	}
L241:
	;
	if base.Ui32(v562) < base.Ui32(v507) {
		v594 = v493
		goto L205
	} else {
		goto L242
	}
L242:
	;
	if base.Ui32(v508) < base.Ui32(v563) {
		goto L238
	} else {
		goto L243
	}
L243:
	;
	if base.Ui32(v563) < base.Ui32(v508) {
		v594 = v493
		goto L205
	} else {
		goto L244
	}
L244:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v493)+12))
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v494)+12))
	if v571 < v572 {
		goto L238
	} else {
		goto L245
	}
L245:
	;
	if v572 < v571 {
		v594 = v493
		goto L205
	} else {
		goto L246
	}
L246:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v493)+16))
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v494)+16))
	if base.Ui32(v576) < base.Ui32(v575) {
		v594 = v493
		goto L205
	} else {
		goto L247
	}
L247:
	;
	goto L238
L248:
	;
	if base.Ui32(v564) < base.Ui32(v511) {
		goto L204
	} else {
		goto L249
	}
L249:
	;
	if base.Ui32(v509) < base.Ui32(v562) {
		goto L203
	} else {
		goto L250
	}
L250:
	;
	if base.Ui32(v562) < base.Ui32(v509) {
		goto L204
	} else {
		goto L251
	}
L251:
	;
	if base.Ui32(v510) < base.Ui32(v563) {
		goto L203
	} else {
		goto L252
	}
L252:
	;
	if base.Ui32(v563) < base.Ui32(v510) {
		goto L204
	} else {
		goto L253
	}
L253:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v492)+12))
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v494)+12))
	if v586 < v587 {
		goto L203
	} else {
		goto L254
	}
L254:
	;
	if v587 < v586 {
		goto L204
	} else {
		goto L255
	}
L255:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v492)+16))
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v494)+16))
	if base.Ui32(v590) < base.Ui32(v591) {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	v593 = v492
	goto L258
L257:
	;
	v593 = v494
	goto L258
L258:
	;
	v594 = v593
	goto L205
L259:
	;
	if base.Ui32(v647) < base.Ui32(v648) {
		v721 = v648
		v723 = v650
		goto L261
	} else {
		goto L262
	}
L261:
	;
	if base.Ui32(v721) <= base.Ui32(v647) {
		goto L278
	} else {
		goto L279
	}
L262:
	;
	v662 = v648
	v664 = v650
	goto L263
L263:
	;
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v662)+8))
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	if base.Ui32(v672) < base.Ui32(v673) {
		v713 = v664
		goto L265
	} else {
		goto L266
	}
L264:
	;
	v721 = v716
	v723 = v713
	goto L261
L265:
	;
	v716 = v662 + int32(20)
	if base.Ui32(v716) <= base.Ui32(v647) {
		v662 = v716
		v664 = v713
		goto L263
	} else {
		goto L276
	}
L266:
	;
	if base.Ui32(v673) < base.Ui32(v672) {
		v721 = v662
		v723 = v664
		goto L261
	} else {
		goto L267
	}
L267:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v662)+4))
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if base.Ui32(v676) < base.Ui32(v677) {
		v713 = v664
		goto L265
	} else {
		goto L268
	}
L268:
	;
	if base.Ui32(v677) < base.Ui32(v676) {
		v721 = v662
		v723 = v664
		goto L261
	} else {
		goto L269
	}
L269:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v662)))
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if base.Ui32(v680) < base.Ui32(v681) {
		v713 = v664
		goto L265
	} else {
		goto L270
	}
L270:
	;
	if base.Ui32(v681) < base.Ui32(v680) {
		v721 = v662
		v723 = v664
		goto L261
	} else {
		goto L271
	}
L271:
	;
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v662)+12))
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	if v684 < v685 {
		v713 = v664
		goto L265
	} else {
		goto L272
	}
L272:
	;
	if v685 < v684 {
		v721 = v662
		v723 = v664
		goto L261
	} else {
		goto L273
	}
L273:
	;
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v662)+16))
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	if base.Ui32(v688) < base.Ui32(v689) {
		v713 = v664
		goto L265
	} else {
		goto L274
	}
L274:
	;
	if base.Ui32(v689) < base.Ui32(v688) {
		v721 = v662
		v723 = v664
		goto L261
	} else {
		goto L275
	}
L275:
	;
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v664)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v692
	v694 = *(*int64)(unsafe.Add(mBase, uint32(v664)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v694
	v696 = *(*int64)(unsafe.Add(mBase, uint32(v664)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v696
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v662)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v664)+16)) = v698
	v700 = *(*int64)(unsafe.Add(mBase, uint32(v662)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v664)+8)) = v700
	v702 = *(*int64)(unsafe.Add(mBase, uint32(v662)))
	*(*int64)(unsafe.Add(mBase, uint32(v664))) = v702
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v662)+16)) = v704
	v706 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v662)+8)) = v706
	v708 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v662))) = v708
	v713 = v664 + int32(20)
	goto L265
L276:
	;
	goto L264
L277:
	;
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v721)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v947
	v949 = *(*int64)(unsafe.Add(mBase, uint32(v721)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v949
	v951 = *(*int64)(unsafe.Add(mBase, uint32(v721)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v951
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v734)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v721)+16)) = v953
	v955 = *(*int64)(unsafe.Add(mBase, uint32(v734)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v721)+8)) = v955
	v957 = *(*int64)(unsafe.Add(mBase, uint32(v734)))
	*(*int64)(unsafe.Add(mBase, uint32(v721))) = v957
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v734)+16)) = v959
	v961 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v734)+8)) = v961
	v963 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v734))) = v963
	v965 = int32(20)
	v647 = v734 - v965
	v648 = v721 + v965
	v650 = v723
	v652 = v739
	goto L259
L278:
	;
	v734 = v647
	v739 = v652
	goto L281
L279:
	;
	v793 = v647
	v798 = v652
	goto L280
L280:
	;
	v805 = int32(20)
	v806 = base.I32_div_s(v723-v20, v805)
	v809 = base.I32_div_s(v721-v723, v805)
	if v806 < v809 {
		goto L295
	} else {
		goto L296
	}
L281:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v734)+8))
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	if base.Ui32(v745) < base.Ui32(v746) {
		goto L277
	} else {
		goto L283
	}
L282:
	;
	v793 = v789
	v798 = v787
	goto L280
L283:
	;
	if base.Ui32(v746) < base.Ui32(v745) {
		v787 = v739
		goto L284
	} else {
		goto L285
	}
L284:
	;
	v789 = v734 - int32(20)
	if base.Ui32(v721) <= base.Ui32(v789) {
		v734 = v789
		v739 = v787
		goto L281
	} else {
		goto L294
	}
L285:
	;
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v734)+4))
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if base.Ui32(v749) < base.Ui32(v750) {
		goto L277
	} else {
		goto L286
	}
L286:
	;
	if base.Ui32(v750) < base.Ui32(v749) {
		v787 = v739
		goto L284
	} else {
		goto L287
	}
L287:
	;
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v734)))
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if base.Ui32(v753) < base.Ui32(v754) {
		goto L277
	} else {
		goto L288
	}
L288:
	;
	if base.Ui32(v754) < base.Ui32(v753) {
		v787 = v739
		goto L284
	} else {
		goto L289
	}
L289:
	;
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v734)+12))
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	if v757 < v758 {
		goto L277
	} else {
		goto L290
	}
L290:
	;
	if v758 < v757 {
		v787 = v739
		goto L284
	} else {
		goto L291
	}
L291:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v734)+16))
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	if base.Ui32(v761) < base.Ui32(v762) {
		goto L277
	} else {
		goto L292
	}
L292:
	;
	if base.Ui32(v762) < base.Ui32(v761) {
		v787 = v739
		goto L284
	} else {
		goto L293
	}
L293:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v734)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v765
	v767 = *(*int64)(unsafe.Add(mBase, uint32(v734)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v767
	v769 = *(*int64)(unsafe.Add(mBase, uint32(v734)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v769
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v739)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v734)+16)) = v771
	v773 = *(*int64)(unsafe.Add(mBase, uint32(v739)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v734)+8)) = v773
	v775 = *(*int64)(unsafe.Add(mBase, uint32(v739)))
	*(*int64)(unsafe.Add(mBase, uint32(v734))) = v775
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v739)+16)) = v777
	v779 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v739)+8)) = v779
	v781 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v739))) = v781
	v787 = v739 - int32(20)
	goto L284
L294:
	;
	goto L282
L295:
	;
	v811 = v806
	goto L297
L296:
	;
	v811 = v809
	goto L297
L297:
	;
	if v811 != 0 {
		goto L298
	} else {
		goto L299
	}
L298:
	;
	v821 = int32(0)
	goto L301
L299:
	;
	goto L300
L300:
	;
	v868 = int32(20)
	v869 = base.I32_div_s(v798-v793, v868)
	v872 = base.I32_div_s(v52-v798, v868)
	v874 = v872 - int32(1)
	if v869 < v874 {
		goto L304
	} else {
		goto L305
	}
L301:
	;
	v830 = v821 * int32(20)
	v831 = v20 + v830
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v831)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v832
	v834 = *(*int64)(unsafe.Add(mBase, uint32(v831)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v834
	v836 = *(*int64)(unsafe.Add(mBase, uint32(v831)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v836
	v838 = v830 + (v721 + v811*int32(-20))
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v838)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v831)+16)) = v839
	v841 = *(*int64)(unsafe.Add(mBase, uint32(v838)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v831)+8)) = v841
	v843 = *(*int64)(unsafe.Add(mBase, uint32(v838)))
	*(*int64)(unsafe.Add(mBase, uint32(v831))) = v843
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v838)+16)) = v845
	v847 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v838)+8)) = v847
	v849 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v838))) = v849
	v852 = v821 + int32(1)
	if v852 != v811 {
		v821 = v852
		goto L301
	} else {
		goto L303
	}
L302:
	;
	goto L300
L303:
	;
	goto L302
L304:
	;
	v876 = v869
	goto L306
L305:
	;
	v876 = v874
	goto L306
L306:
	;
	if v876 != 0 {
		goto L307
	} else {
		goto L308
	}
L307:
	;
	v887 = int32(0)
	goto L310
L308:
	;
	goto L309
L309:
	;
	if base.Ui32(v809) <= base.Ui32(v869) {
		goto L313
	} else {
		goto L314
	}
L310:
	;
	v895 = v887 * int32(20)
	v896 = v721 + v895
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v896)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v897
	v899 = *(*int64)(unsafe.Add(mBase, uint32(v896)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v899
	v901 = *(*int64)(unsafe.Add(mBase, uint32(v896)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v901
	v903 = v895 + (v52 + v876*int32(-20))
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v903)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v896)+16)) = v904
	v906 = *(*int64)(unsafe.Add(mBase, uint32(v903)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v896)+8)) = v906
	v908 = *(*int64)(unsafe.Add(mBase, uint32(v903)))
	*(*int64)(unsafe.Add(mBase, uint32(v896))) = v908
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v903)+16)) = v910
	v912 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v903)+8)) = v912
	v914 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v903))) = v914
	v917 = v887 + int32(1)
	if v917 != v876 {
		v887 = v917
		goto L310
	} else {
		goto L312
	}
L311:
	;
	goto L309
L312:
	;
	goto L311
L313:
	;
	F_sort_pending_writebacks(m, v20, v809)
	mBase = m.M
	v936 = v52 + v869*int32(-20)
	if base.Ui32(v869) < base.Ui32(int32(7)) {
		v982 = v936
		v984 = v869
		goto L2
	} else {
		goto L316
	}
L314:
	;
	goto L315
L315:
	;
	F_sort_pending_writebacks(m, v52+v869*int32(-20), v869)
	mBase = m.M
	if base.Ui32(v809) < base.Ui32(int32(7)) {
		v969 = v20
		v970 = v809
		goto L3
	} else {
		goto L317
	}
L316:
	;
	v20 = v936
	v21 = v869
	goto L5
L317:
	;
	v36 = v809
	goto L7
L318:
	;
	v997 = int32(20)
	v1007 = v982 + v997
	goto L319
L319:
	;
	if base.Ui32(v1007) <= base.Ui32(v982) {
		goto L321
	} else {
		goto L322
	}
L320:
	;
	goto L1
L321:
	;
	v1093 = v1007 + int32(20)
	if base.Ui32(v1093) < base.Ui32(v982+v984*v997) {
		v1007 = v1093
		goto L319
	} else {
		goto L336
	}
L322:
	;
	v1019 = v1007
	goto L323
L323:
	;
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(v1019-int32(12))))
	v1032 = *(*int32)(unsafe.Add(mBase, uint32(v1019)+8))
	if base.Ui32(v1031) < base.Ui32(v1032) {
		goto L321
	} else {
		goto L325
	}
L324:
	;
	goto L321
L325:
	;
	v1035 = v1019 - int32(20)
	if base.Ui32(v1032) < base.Ui32(v1031) {
		goto L326
	} else {
		goto L327
	}
L326:
	;
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(v1019)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v1060
	v1062 = *(*int64)(unsafe.Add(mBase, uint32(v1019)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v1062
	v1064 = *(*int64)(unsafe.Add(mBase, uint32(v1019)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v1064
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(v1035)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1019)+16)) = v1066
	v1068 = *(*int64)(unsafe.Add(mBase, uint32(v1035)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1019)+8)) = v1068
	v1070 = *(*int64)(unsafe.Add(mBase, uint32(v1035)))
	*(*int64)(unsafe.Add(mBase, uint32(v1019))) = v1070
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1035)+16)) = v1072
	v1074 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1035)+8)) = v1074
	v1076 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1035))) = v1076
	if base.Ui32(v982) < base.Ui32(v1035) {
		v1019 = v1035
		goto L323
	} else {
		goto L335
	}
L327:
	;
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(v1019-int32(16))))
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v1019)+4))
	if base.Ui32(v1039) < base.Ui32(v1040) {
		goto L321
	} else {
		goto L328
	}
L328:
	;
	if base.Ui32(v1040) < base.Ui32(v1039) {
		goto L326
	} else {
		goto L329
	}
L329:
	;
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v1035)))
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(v1019)))
	if base.Ui32(v1043) < base.Ui32(v1044) {
		goto L321
	} else {
		goto L330
	}
L330:
	;
	if base.Ui32(v1044) < base.Ui32(v1043) {
		goto L326
	} else {
		goto L331
	}
L331:
	;
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v1019-int32(8))))
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v1019)+12))
	if v1049 < v1050 {
		goto L321
	} else {
		goto L332
	}
L332:
	;
	if v1050 < v1049 {
		goto L326
	} else {
		goto L333
	}
L333:
	;
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v1019-int32(4))))
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v1019)+16))
	if base.Ui32(v1055) <= base.Ui32(v1056) {
		goto L321
	} else {
		goto L334
	}
L334:
	;
	goto L326
L335:
	;
	goto L324
L336:
	;
	goto L320
}
func F_sortouts_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if v8 < v11 {
		return int32(-1)
	} else {
		v15 = int32(1)
		if v11 < v8 {
			v28 = v15
			return v28
		} else {
			v17 = int32(*(*int16)(unsafe.Add(mBase, uint32(v6)+4)))
			v18 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9)+4)))
			if v17 < v18 {
				return int32(-1)
			} else {
				if v18 < v17 {
					v28 = v15
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					if v24 < v25 {
						v28 = int32(-1)
					} else {
						v28 = base.B2i32(v25 < v24)
					}
				}
				return v28
			}
		}
	}
}
func F_soundex(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var __phi64 int32
	_ = __phi64
	var v65 int32
	_ = v65
	var __phi65 int32
	_ = __phi65
	var v67 int32
	_ = v67
	var __phi67 int32
	_ = __phi67
	var v68 int32
	_ = v68
	var __phi68 int32
	_ = __phi68
	var v69 int32
	_ = v69
	var __phi69 int32
	_ = __phi69
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v12 = F_text_to_cstring(m, v8)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v15 = v5 + int32(11)
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v20 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v182 = F_cstring_to_text(m, v15)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L49
	}
L5:
	;
	if base.Ui32((v23-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	v21 = v12
	v23 = v20
	goto L9
L7:
	;
	goto L8
L8:
	;
	v44 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v44)
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v44
	goto L4
L9:
	;
	if base.Ui32((v23&int32(223)-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L5
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if v35 != 0 {
		v21 = v21 + int32(1)
		v23 = v35
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v56 = v23 - int32(32)
	goto L15
L14:
	;
	v56 = v23
	goto L15
L15:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v56)
	v58 = int32(1)
	v60 = v5 + int32(12)
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if v61 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v174 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v171))) = uint8(v174)
	goto L4
L17:
	;
	__phi64 = v21
	__phi65 = v61
	__phi67 = v60
	__phi68 = v58
	__phi69 = v21 + int32(1)
	v64 = __phi64
	v65 = __phi65
	v67 = __phi67
	v68 = __phi68
	v69 = __phi69
	goto L20
L18:
	;
	v160 = v60
	v161 = v58
	goto L19
L19:
	;
	v164 = int32(4) - v161
	if v164 != 0 {
		goto L46
	} else {
		goto L47
	}
L20:
	;
	if base.Ui32(int32(25)) < base.Ui32((v65&int32(223)-int32(65))&int32(255)) {
		v145 = v67
		v146 = v68
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if int32(3) < v146 {
		v171 = v145
		goto L16
	} else {
		goto L45
	}
L22:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+1)))
	if v148 != 0 {
		goto L41
	} else {
		goto L42
	}
L23:
	;
	if base.Ui32((v65-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v86 = v65 - int32(32)
	goto L26
L25:
	;
	v86 = v65
	goto L26
L26:
	;
	v92 = base.B2i32(base.Ui32(int32(25)) < base.Ui32((v86-int32(65))&int32(255)))
	if base.Ui32(int32(25)) < base.Ui32((v86-int32(65))&int32(255)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v98 = v86
	goto L29
L28:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86&int32(255))+uint32(_c_F_soundex[0]))))
	v98 = v97
	goto L29
L29:
	;
	v99 = int32(255)
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	if base.Ui32((v101-int32(97))&v99) < base.Ui32(int32(26)) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v110 = v101 - int32(32)
	goto L32
L31:
	;
	v110 = v101
	goto L32
L32:
	;
	if base.Ui32((v110-int32(65))&int32(255)) <= base.Ui32(int32(25)) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110&int32(255))+uint32(_c_F_soundex[0]))))
	v122 = v121
	goto L35
L34:
	;
	v122 = v110
	goto L35
L35:
	;
	if v98&v99 == v122&int32(255) {
		v145 = v67
		v146 = v68
		goto L22
	} else {
		goto L36
	}
L36:
	;
	if v92 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86&int32(255))+uint32(_c_F_soundex[0]))))
	v133 = v132
	goto L39
L38:
	;
	v133 = v86
	goto L39
L39:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v67))) = uint8(v133)
	if v133&int32(255) == int32(48) {
		v145 = v67
		v146 = v68
		goto L22
	} else {
		goto L40
	}
L40:
	;
	v139 = int32(1)
	v145 = v67 + v139
	v146 = v68 + v139
	goto L22
L41:
	;
	if v146 < int32(4) {
		__phi64 = v69
		__phi65 = v148
		__phi67 = v145
		__phi68 = v146
		__phi69 = v69 + int32(1)
		v64 = __phi64
		v65 = __phi65
		v67 = __phi67
		v68 = __phi68
		v69 = __phi69
		goto L20
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	goto L21
L44:
	;
	goto L43
L45:
	;
	v160 = v145
	v161 = v146
	goto L19
L46:
	;
	base.MemoryFill(m, v160, int32(48), v164)
	goto L48
L47:
	;
	goto L48
L48:
	;
	v171 = v164 + v160
	goto L16
L49:
	;
	m.G0 = v5 + int32(16)
	return base.I64_extend_i32_u(v182)
}
func F_spgbuild(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int64
	_ = v60
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 float64
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int64
	_ = v220
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	v9 = m.G0
	v11 = v9 - int32(112)
	m.G0 = v11
	v14 = F_RelationGetNumberOfBlocksInFork(m, l1, int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v14 == int32(0) {
			v20 = F_SpGistNewBuffer(m, l1)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v22 = F_SpGistNewBuffer(m, l1)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v24 = F_SpGistNewBuffer(m, l1)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = int32(_a_F_spgbuild_0)
						v28 = *(*int32)(unsafe.Add(mBase, _c_F_spgbuild[0]))
						*(*int32)(unsafe.Add(mBase, _c_F_spgbuild[0])) = v28 + int32(1)
						if v20 < int32(0) {
							v35 = *(*int32)(unsafe.Add(mBase, _c_F_spgbuild[1]))
							v41 = *(*int32)(unsafe.Add(mBase, uint32(v35+(v20^int32(-1))<<(uint(int32(2))%32))))
							v49 = v41
						} else {
							v43 = *(*int32)(unsafe.Add(mBase, _c_F_spgbuild[2]))
							v49 = v43 + v20<<(uint(int32(13))%32) + int32(-8192)
						}
						F_PageInit(m, v49, int32(_a_F_spgbuild_1), int32(8))
						mBase = m.M
						v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v49)+16)))
						v55 = v49 + v54
						v56 = int32(_a_F_spgbuild_2)
						*(*uint16)(unsafe.Add(mBase, uint32(v55)+6)) = uint16(v56)
						v58 = int32(1)
						*(*uint16)(unsafe.Add(mBase, uint32(v55))) = uint16(v58)
						v60 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v49)+80)) = v60
						*(*int64)(unsafe.Add(mBase, uint32(v49)+72)) = v60
						*(*int64)(unsafe.Add(mBase, uint32(v49)+64)) = v60
						*(*int64)(unsafe.Add(mBase, uint32(v49)+56)) = v60
						*(*int64)(unsafe.Add(mBase, uint32(v49)+48)) = v60
						*(*int64)(unsafe.Add(mBase, uint32(v49)+40)) = v60
						*(*int64)(unsafe.Add(mBase, uint32(v49)+32)) = v60
						*(*int32)(unsafe.Add(mBase, uint32(v49)+88)) = int32(0)
						*(*int64)(unsafe.Add(mBase, uint32(v49)+24)) = int64(-1173640210)
						v78 = int32(92)
						*(*uint16)(unsafe.Add(mBase, uint32(v49)+12)) = uint16(v78)
						v80 = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v49)+84)) = v80
						*(*int32)(unsafe.Add(mBase, uint32(v49)+76)) = v80
						*(*int32)(unsafe.Add(mBase, uint32(v49)+68)) = v80
						*(*int32)(unsafe.Add(mBase, uint32(v49)+60)) = v80
						*(*int32)(unsafe.Add(mBase, uint32(v49)+52)) = v80
						*(*int32)(unsafe.Add(mBase, uint32(v49)+44)) = v80
						*(*int32)(unsafe.Add(mBase, uint32(v49)+36)) = v80
						F_MarkBufferDirty(m, v20)
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return int32(0)
						} else {
							v96 = int32(4)
							if v22 < int32(0) {
								v100 = *(*int32)(unsafe.Add(mBase, _c_F_spgbuild[1]))
								v106 = *(*int32)(unsafe.Add(mBase, uint32(v100+(v22^int32(-1))<<(uint(int32(2))%32))))
								v114 = v106
							} else {
								v108 = *(*int32)(unsafe.Add(mBase, _c_F_spgbuild[2]))
								v114 = v108 + v22<<(uint(int32(13))%32) + int32(-8192)
							}
							F_PageInit(m, v114, int32(_a_F_spgbuild_1), int32(8))
							mBase = m.M
							v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114)+16)))
							v119 = v114 + v118
							v120 = int32(_a_F_spgbuild_2)
							*(*uint16)(unsafe.Add(mBase, uint32(v119)+6)) = uint16(v120)
							*(*uint16)(unsafe.Add(mBase, uint32(v119))) = uint16(v96)
							F_MarkBufferDirty(m, v22)
							mBase = m.M
							v124 = m.ExcPending
							if v124 != 0 {
								return int32(0)
							} else {
								v125 = int32(12)
								if v24 < int32(0) {
									v129 = *(*int32)(unsafe.Add(mBase, _c_F_spgbuild[1]))
									v135 = *(*int32)(unsafe.Add(mBase, uint32(v129+(v24^int32(-1))<<(uint(int32(2))%32))))
									v143 = v135
								} else {
									v137 = *(*int32)(unsafe.Add(mBase, _c_F_spgbuild[2]))
									v143 = v137 + v24<<(uint(int32(13))%32) + int32(-8192)
								}
								F_PageInit(m, v143, int32(_a_F_spgbuild_1), int32(8))
								mBase = m.M
								v147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v143)+16)))
								v148 = v143 + v147
								v149 = int32(_a_F_spgbuild_2)
								*(*uint16)(unsafe.Add(mBase, uint32(v148)+6)) = uint16(v149)
								*(*uint16)(unsafe.Add(mBase, uint32(v148))) = uint16(v125)
								F_MarkBufferDirty(m, v24)
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return int32(0)
								} else {
									v154 = int32(_a_F_spgbuild_0)
									v156 = *(*int32)(unsafe.Add(mBase, _c_F_spgbuild[0]))
									*(*int32)(unsafe.Add(mBase, _c_F_spgbuild[0])) = v156 - int32(1)
									F_UnlockReleaseBuffer(m, v20)
									mBase = m.M
									v161 = m.ExcPending
									if v161 != 0 {
										return int32(0)
									} else {
										F_UnlockReleaseBuffer(m, v22)
										mBase = m.M
										v163 = m.ExcPending
										if v163 != 0 {
											return int32(0)
										} else {
											F_UnlockReleaseBuffer(m, v24)
											mBase = m.M
											v165 = m.ExcPending
											if v165 != 0 {
												return int32(0)
											} else {
												v167 = v11 + int32(8)
												F_initSpGistState(m, v167, l1)
												mBase = m.M
												v169 = m.ExcPending
												if v169 != 0 {
													return int32(0)
												} else {
													*(*int64)(unsafe.Add(mBase, uint32(v11)+96)) = int64(0)
													v172 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+88)) = uint8(v172)
													v175 = *(*int32)(unsafe.Add(mBase, _c_F_spgbuild[3]))
													v180 = F_AllocSetContextCreateInternal(m, v175, int32(_a_F_spgbuild_3), int32(0), int32(_a_F_spgbuild_1), int32(_a_F_spgbuild_4))
													mBase = m.M
													v181 = m.ExcPending
													if v181 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v11)+104)) = v180
														v183 = int32(1)
														v184 = int32(0)
														v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
														v191 = *(*int32)(unsafe.Add(mBase, uint32(v190)+140))
														v192 = m.T0[v191].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, l0, l1, l2, v183, v184, v183, v184, int32(-1), int32(253), v167, v184)
														mBase = m.M
														v193 = m.ExcPending
														if v193 != 0 {
															return int32(0)
														} else {
															v194 = *(*int32)(unsafe.Add(mBase, uint32(v11)+104))
															F_MemoryContextDelete(m, v194)
															mBase = m.M
															v196 = m.ExcPending
															if v196 != 0 {
																return int32(0)
															} else {
																F_SpGistUpdateMetaPage(m, l1)
																mBase = m.M
																v198 = m.ExcPending
																if v198 != 0 {
																	return int32(0)
																} else {
																	v199 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
																	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199)+118)))
																	if v200 != int32(112) {
																		v217 = F_palloc0(m, int32(16))
																		mBase = m.M
																		v218 = m.ExcPending
																		if v218 != 0 {
																			return int32(0)
																		} else {
																			*(*float64)(unsafe.Add(mBase, uint32(v217))) = v192
																			v220 = *(*int64)(unsafe.Add(mBase, uint32(v11)+96))
																			*(*float64)(unsafe.Add(mBase, uint32(v217)+8)) = base.F64_convert_i64_s(v220)
																			m.G0 = v11 + int32(112)
																			return v217
																		}
																	} else {
																		v204 = *(*int32)(unsafe.Add(mBase, _c_F_spgbuild[4]))
																		if v204 <= int32(0) {
																			v207 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
																			if v207 != 0 {
																				v217 = F_palloc0(m, int32(16))
																				mBase = m.M
																				v218 = m.ExcPending
																				if v218 != 0 {
																					return int32(0)
																				} else {
																					*(*float64)(unsafe.Add(mBase, uint32(v217))) = v192
																					v220 = *(*int64)(unsafe.Add(mBase, uint32(v11)+96))
																					*(*float64)(unsafe.Add(mBase, uint32(v217)+8)) = base.F64_convert_i64_s(v220)
																					m.G0 = v11 + int32(112)
																					return v217
																				}
																			} else {
																				v208 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
																				if v208 != 0 {
																					v217 = F_palloc0(m, int32(16))
																					mBase = m.M
																					v218 = m.ExcPending
																					if v218 != 0 {
																						return int32(0)
																					} else {
																						*(*float64)(unsafe.Add(mBase, uint32(v217))) = v192
																						v220 = *(*int64)(unsafe.Add(mBase, uint32(v11)+96))
																						*(*float64)(unsafe.Add(mBase, uint32(v217)+8)) = base.F64_convert_i64_s(v220)
																						m.G0 = v11 + int32(112)
																						return v217
																					}
																				} else {
																					v209 = int32(0)
																					v211 = F_RelationGetNumberOfBlocksInFork(m, l1, v209)
																					mBase = m.M
																					v212 = m.ExcPending
																					if v212 != 0 {
																						return int32(0)
																					} else {
																						F_log_newpage_range(m, l1, v209, v211, int32(1))
																						mBase = m.M
																						v215 = m.ExcPending
																						if v215 != 0 {
																							return int32(0)
																						} else {
																							v217 = F_palloc0(m, int32(16))
																							mBase = m.M
																							v218 = m.ExcPending
																							if v218 != 0 {
																								return int32(0)
																							} else {
																								*(*float64)(unsafe.Add(mBase, uint32(v217))) = v192
																								v220 = *(*int64)(unsafe.Add(mBase, uint32(v11)+96))
																								*(*float64)(unsafe.Add(mBase, uint32(v217)+8)) = base.F64_convert_i64_s(v220)
																								m.G0 = v11 + int32(112)
																								return v217
																							}
																						}
																					}
																				}
																			}
																		} else {
																			v209 = int32(0)
																			v211 = F_RelationGetNumberOfBlocksInFork(m, l1, v209)
																			mBase = m.M
																			v212 = m.ExcPending
																			if v212 != 0 {
																				return int32(0)
																			} else {
																				F_log_newpage_range(m, l1, v209, v211, int32(1))
																				mBase = m.M
																				v215 = m.ExcPending
																				if v215 != 0 {
																					return int32(0)
																				} else {
																					v217 = F_palloc0(m, int32(16))
																					mBase = m.M
																					v218 = m.ExcPending
																					if v218 != 0 {
																						return int32(0)
																					} else {
																						*(*float64)(unsafe.Add(mBase, uint32(v217))) = v192
																						v220 = *(*int64)(unsafe.Add(mBase, uint32(v11)+96))
																						*(*float64)(unsafe.Add(mBase, uint32(v217)+8)) = base.F64_convert_i64_s(v220)
																						m.G0 = v11 + int32(112)
																						return v217
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
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v230 = m.ExcPending
			if v230 != 0 {
				return int32(0)
			} else {
				v231 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v231 + int32(4)
				F_errmsg_internal(m, int32(_a_F_spgbuild_5), v11)
				mBase = m.M
				v237 = m.ExcPending
				if v237 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_spgbuild_6), int32(84), int32(_a_F_spgbuild_7))
					mBase = m.M
					v242 = m.ExcPending
					if v242 != 0 {
						return int32(0)
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
func F_spgcanreturn(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	if l1 <= int32(1) {
		v5 = F_spgGetCache(m, l0)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+12)))
			v11 = v9
			return v11 & int32(1)
		}
	} else {
		v11 = int32(1)
		return v11 & int32(1)
	}
}
func F_spgendscan(m *base.Module, l0 int32) {
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
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
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
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+88))
	F_MemoryContextDelete(m, v5)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+92))
	F_MemoryContextDelete(m, v8)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v4)+104))
	if v11 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	F_pfree(m, v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v4)+68))
	if v14 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L6
L8:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v4)+72))
	if v22 != 0 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+52))
	if v14 == v18 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	F_FreeTupleDesc(m, v14)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	F_pfree(m, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if int32(0) < v25 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L14
L16:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v4)+120))
	F_pfree(m, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	F_pfree(m, v4)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L25
	}
L19:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v4)+124))
	F_pfree(m, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v4)+188))
	F_pfree(m, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v4)+192))
	F_pfree(m, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	F_pfree(m, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	F_pfree(m, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	goto L18
L25:
	;
	return
}
func F_spggettuple(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	if l1 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+208)) = uint8(v13)
	v16 = v12 + int32(3488)
	v18 = v12 + int32(_a_F_spggettuple_0)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+216))
	v21 = v19
	goto L4
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L12
	} else {
		goto L33
	}
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v12)+220))
	if v29 < v21 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	return base.B2i32(v29 < v21)
L6:
	;
	v33 = v12 + v29*int32(6)
	v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33)+228)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+64)) = uint16(v34)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+224))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v36
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v12)+220))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v38+int32(2672)))))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)) = uint8(v42)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v12)+220))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v16+v44<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v48
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v12)+108))
	if int32(0) < v50 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v12)+108))
	if v74 <= int32(0) {
		v105 = v21
		goto L14
	} else {
		goto L15
	}
L9:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v12)+220))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v18+v54<<(uint(int32(2))%32))))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54+v12+int32(3080)))))
	F_index_store_float8_orderby_distances(m, l0, v53, v58, v62)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v12)+220))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+220)) = v68 + int32(1)
	return base.B2i32(v29 < v21)
L12:
	;
	return int32(0)
L13:
	;
	goto L11
L14:
	;
	if v105 <= int32(0) {
		goto L24
	} else {
		goto L25
	}
L15:
	;
	v77 = int32(0)
	if v21 <= v77 {
		v105 = v21
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v81 = v77
	v85 = v21
	goto L17
L17:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v18+v81<<(uint(int32(2))%32))))
	if v92 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v105 = v96
	goto L14
L19:
	;
	F_pfree(m, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L12
	} else {
		goto L22
	}
L20:
	;
	v96 = v85
	goto L21
L21:
	;
	v98 = v81 + int32(1)
	if v98 < v96 {
		v81 = v98
		v85 = v96
		goto L17
	} else {
		goto L23
	}
L22:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v12)+216))
	v96 = v95
	goto L21
L23:
	;
	goto L18
L24:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+216)) = int64(0)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_spgWalk(m, v147, v12, int32(0), int32(266))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L12
	} else {
		goto L31
	}
L25:
	;
	v111 = int32(0)
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+208)))
	if v112&int32(1) == v111 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v118 = v111
	goto L27
L27:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v16+v118<<(uint(int32(2))%32))))
	F_pfree(m, v129)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L12
	} else {
		goto L29
	}
L28:
	;
	goto L24
L29:
	;
	v133 = v118 + int32(1)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v12)+216))
	if v133 < v134 {
		v118 = v133
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v12)+216))
	if v152 != 0 {
		v21 = v152
		goto L4
	} else {
		goto L32
	}
L32:
	;
	goto L5
L33:
	;
	F_errmsg_internal(m, int32(_a_F_spggettuple_1), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L12
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_spggettuple_2), int32(1026), int32(_a_F_spggettuple_3))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L12
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_split_pathtarget_at_srfs_extended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
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
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
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
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v390 int32
	_ = v390
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v507 int32
	_ = v507
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v608 int32
	_ = v608
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	v6 = l5
	v7 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(96)
	m.G0 = v20
	if l1 == l2 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v20 + int32(96)
	return
L2:
	;
	if v149 != 0 {
		goto L47
	} else {
		goto L48
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v205
	goto L1
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = l1
	v28 = F_list_make1_impl(m, int32(1), v20+int32(4))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+60)) = uint8(v6)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = l0
	if l2 != 0 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	return
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v28
	v31 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v31
	v36 = F_list_make1_impl(m, int32(479), v20)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v205 = v36
	goto L3
L10:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v42 = v41
	goto L12
L11:
	;
	v42 = int32(0)
	goto L12
L12:
	;
	v43 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v43
	v51 = F_list_make1_impl(m, int32(1), v20+int32(24))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v53 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v20)+68)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v53
	v61 = F_list_make1_impl(m, int32(1), v20+int32(20))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v63 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v20)+72)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v63
	v71 = F_list_make1_impl(m, int32(1), v20+int32(16))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+80)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+76)) = v71
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v76 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	if v77 <= int32(0) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = l1
	v176 = F_list_make1_impl(m, int32(1), v20+int32(12))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L7
	} else {
		goto L44
	}
L19:
	;
	if v142 != 0 {
		goto L2
	} else {
		goto L43
	}
L20:
	;
	v142 = int32(0)
	v149 = v7
	goto L19
L21:
	;
	goto L22
L22:
	;
	v87 = int32(0)
	v91 = v7
	v94 = v7
	goto L23
L23:
	;
	v100 = v91 << (uint(int32(2)) % 32)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v100+v101)))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v105 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v142 = v130
	v149 = v132
	goto L19
L25:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v100+v105)))
	v108 = v107
	goto L27
L26:
	;
	v108 = int32(0)
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+88)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+92)) = v108
	v114 = F_split_pathtarget_walker(m, v103, v20+int32(56))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L7
	} else {
		goto L28
	}
L28:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v20)+88))
	if v116 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v117 = base.B2i32(v116 < v87)
	if v116 < v87 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v130 = v87
	v132 = v94
	goto L31
L31:
	;
	v134 = v91 + int32(1)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	if v134 < v135 {
		v87 = v130
		v91 = v134
		v94 = v132
		goto L23
	} else {
		goto L42
	}
L32:
	;
	v118 = v87
	goto L34
L33:
	;
	v118 = v116
	goto L34
L34:
	;
	v120 = base.B2i32(v116 <= v87) & v94
	if v116 < v87 {
		v129 = v120
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v130 = v118
	v132 = v129
	goto L31
L36:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	switch v121 - int32(15) {
	case 0:
		goto L39
	default:
		goto L37
	case 2:
		goto L38
	}
L37:
	;
	v129 = int32(1)
	goto L35
L38:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+16)))
	if v127 != 0 {
		v129 = v120
		goto L35
	} else {
		goto L41
	}
L39:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+12)))
	if v124 == int32(0) {
		goto L37
	} else {
		goto L40
	}
L40:
	;
	v129 = v120
	goto L35
L41:
	;
	goto L37
L42:
	;
	goto L24
L43:
	;
	goto L18
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v176
	v179 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v179
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v179
	v186 = F_list_make1_impl(m, int32(479), v20+int32(8))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L7
	} else {
		goto L45
	}
L45:
	;
	v205 = v186
	goto L3
L46:
	;
	v244 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v244
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v244
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v20)+72))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v20)+68))
	v263 = v7
	v266 = v7
	goto L55
L47:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v20)+68))
	v209 = F_lappend(m, v207, int32(0))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L7
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v223 = v142 << (uint(int32(2)) % 32)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v20)+72))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+12))
	v226 = v223 + v225
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v20)+80))
	v229 = F_list_concat(m, v227, v228)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L7
	} else {
		goto L53
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+68)) = v209
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v20)+72))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v20)+80))
	v214 = F_lappend(m, v212, v213)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L7
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+72)) = v214
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v20)+76))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v20)+84))
	v219 = F_lappend(m, v217, v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L7
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+76)) = v219
	v243 = v219
	goto L46
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v226))) = v229
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v20)+76))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)+12))
	v234 = v233 + v223
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v20)+84))
	v237 = F_list_concat(m, v235, v236)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L7
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234))) = v237
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v20)+76))
	v243 = v240
	goto L46
L55:
	;
	v267 = int32(0)
	if v249 == v267 {
		v277 = v267
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v278 = int32(0)
	if v248 == v278 {
		v287 = v278
		goto L60
	} else {
		goto L61
	}
L58:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v249)+4))
	if v271 <= v263 {
		v277 = int32(0)
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v249)+12))
	v277 = v273 + v263<<(uint(int32(2))%32)
	goto L57
L60:
	;
	if v243 == int32(0) {
		goto L1
	} else {
		goto L63
	}
L61:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v248)+4))
	if v281 <= v263 {
		v287 = v278
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v248)+12))
	v287 = v283 + v263<<(uint(int32(2))%32)
	goto L60
L63:
	;
	v290 = int32(0)
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v243)+4))
	if base.B2i32(v287 == v290)|(base.B2i32(v277 == v290)|base.B2i32(v294 <= v263)) != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v243)+12))
	if v298 == int32(0) {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v277)))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v20)+68))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v304)+12))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v304)+4))
	if base.Ui32(v277+int32(4)) < base.Ui32(v305+v306<<(uint(int32(2))%32)) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v312 = F_palloc0(m, int32(40))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L7
	} else {
		goto L69
	}
L67:
	;
	v608 = l1
	goto L68
L68:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v617 = F_lappend(m, v616, v608)
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L7
	} else {
		goto L120
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v312))) = int32(280)
	if v301 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v20)+72))
	if v367 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L71:
	;
	v318 = int32(0)
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v301)+4))
	if v319 <= v318 {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v324 = v318
	goto L73
L73:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v301)+12))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v339+v324<<(uint(int32(2))%32))))
	F_add_sp_item_to_pathtarget(m, v312, v343)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L7
	} else {
		goto L75
	}
L74:
	;
	goto L70
L75:
	;
	v347 = v324 + int32(1)
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v301)+4))
	if v347 < v348 {
		v324 = v347
		goto L73
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v20)+76))
	if v479 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L78:
	;
	v371 = v287 + int32(4)
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v367)+12))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v367)+4))
	v378 = base.B2i32(base.Ui32(v371) < base.Ui32(v373+v374<<(uint(int32(2))%32)))
	if base.Ui32(v371) < base.Ui32(v373+v374<<(uint(int32(2))%32)) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v379 = v371
	goto L81
L80:
	;
	v379 = int32(0)
	goto L81
L81:
	;
	if base.Ui32(v371) < base.Ui32(v373+v374<<(uint(int32(2))%32)) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v383 = (v379 - v373) >> (uint(int32(2)) % 32)
	goto L84
L83:
	;
	v383 = v374
	goto L84
L84:
	;
	if v374 <= v383 {
		goto L77
	} else {
		goto L85
	}
L85:
	;
	v390 = v383
	goto L86
L86:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v367)+12))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v402+v390<<(uint(int32(2))%32))))
	if v406 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	goto L77
L88:
	;
	v459 = v390 + int32(1)
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v367)+4))
	if v459 < v460 {
		v390 = v459
		goto L86
	} else {
		goto L95
	}
L89:
	;
	v409 = int32(0)
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v406)+4))
	if v410 <= v409 {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v415 = v409
	goto L91
L91:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v406)+12))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v430+v415<<(uint(int32(2))%32))))
	F_add_sp_item_to_pathtarget(m, v312, v434)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L7
	} else {
		goto L93
	}
L92:
	;
	goto L88
L93:
	;
	v438 = v415 + int32(1)
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v406)+4))
	if v438 < v439 {
		v415 = v438
		goto L91
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	goto L87
L96:
	;
	v597 = F_set_pathtarget_cost_width(m, l0, v312)
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L7
	} else {
		goto L119
	}
L97:
	;
	v482 = int32(2)
	v486 = v298 + v263<<(uint(v482)%32) + int32(4)
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v479)+12))
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v479)+4))
	v493 = base.B2i32(base.Ui32(v486) < base.Ui32(v488+v489<<(uint(v482)%32)))
	if base.Ui32(v486) < base.Ui32(v488+v489<<(uint(v482)%32)) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v494 = v486
	goto L100
L99:
	;
	v494 = int32(0)
	goto L100
L100:
	;
	if base.Ui32(v486) < base.Ui32(v488+v489<<(uint(v482)%32)) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v498 = (v494 - v488) >> (uint(int32(2)) % 32)
	goto L103
L102:
	;
	v498 = v489
	goto L103
L103:
	;
	if v489 <= v498 {
		goto L96
	} else {
		goto L104
	}
L104:
	;
	v507 = v498
	goto L105
L105:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v479)+12))
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v517+v507<<(uint(int32(2))%32))))
	if v521 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	goto L96
L107:
	;
	v577 = v507 + int32(1)
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v479)+4))
	if v577 < v578 {
		v507 = v577
		goto L105
	} else {
		goto L118
	}
L108:
	;
	v524 = int32(0)
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v521)+4))
	if v525 <= v524 {
		goto L107
	} else {
		goto L109
	}
L109:
	;
	v530 = v524
	goto L110
L110:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v521)+12))
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v545+v530<<(uint(int32(2))%32))))
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v549)))
	v551 = F_list_member(m, v266, v550)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L7
	} else {
		goto L112
	}
L111:
	;
	goto L107
L112:
	;
	if v551 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	F_add_sp_item_to_pathtarget(m, v312, v549)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L7
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v556 = v530 + int32(1)
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v521)+4))
	if v556 < v557 {
		v530 = v556
		goto L110
	} else {
		goto L117
	}
L116:
	;
	goto L115
L117:
	;
	goto L111
L118:
	;
	goto L106
L119:
	;
	v608 = v312
	goto L68
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v617
	v620 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v623 = F_lappend_int(m, v620, base.B2i32(v301 != int32(0)))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L7
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v623
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v608)+4))
	v263 = v263 + int32(1)
	v266 = v628
	goto L55
}
func F_stop_repack_decoding_worker_cb(m *base.Module, l0 int32, l1 int64) {
	var v4 int32
	_ = v4
	F_stop_repack_decoding_worker(m)
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_str_tolower(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l0 == v4 {
		v64 = v4
		m.G0 = v10 + int32(16)
		return v64
	} else {
		if l2 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v75 = m.ExcPending
			if v75 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(34209924))
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_str_tolower_0)
					F_errmsg(m, int32(_a_F_str_tolower_1), v10)
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return int32(0)
					} else {
						F_errhint(m, int32(_a_F_str_tolower_2), int32(0))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_str_tolower_3), int32(1639), int32(_a_F_str_tolower_4))
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return int32(0)
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
			v16 = F_pg_newlocale_from_collation(m, l2)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+2)))
				if v20 == int32(1) {
					v23 = F_pnstrdup(m, l0, l1)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
						if v25 == int32(0) {
							v64 = v23
						} else {
							v28 = v25
							v30 = v23
							for {
								if base.Ui32((v28-int32(65))&int32(255)) < base.Ui32(int32(26)) {
									v43 = v28 | int32(32)
								} else {
									v43 = v28
								}
								*(*uint8)(unsafe.Add(mBase, uint32(v30))) = uint8(v43)
								v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+1)))
								if v45 != 0 {
									v28 = v45
									v30 = v30 + int32(1)
									continue
								} else {
									break
								}
								break
							}
							v64 = v23
						}
						m.G0 = v10 + int32(16)
						return v64
					}
				} else {
					v49 = l1 + int32(1)
					v50 = F_palloc(m, v49)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						v52 = F_pg_strlower(m, v50, v49, l0, l1, v16)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							v55 = v52 + int32(1)
							if base.Ui32(v55) <= base.Ui32(v49) {
								v64 = v50
								m.G0 = v10 + int32(16)
								return v64
							} else {
								v57 = F_repalloc(m, v50, v55)
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return int32(0)
								} else {
									v59 = F_pg_strlower(m, v57, v55, l0, l1, v16)
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return int32(0)
									} else {
										v64 = v57
										m.G0 = v10 + int32(16)
										return v64
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
func F_strcoll_libc(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if base.B2i32(v7 == int32(0))|base.B2i32(v7 != v10) != 0 {
		v28 = v7
		v29 = v10
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v28 - v29
L2:
	;
	goto L1
L3:
	;
	v13 = l0
	v14 = l1
	goto L4
L4:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	if v18 == int32(0) {
		v28 = v18
		v29 = v17
		goto L2
	} else {
		goto L6
	}
L5:
	;
	v28 = v18
	v29 = v17
	goto L2
L6:
	;
	v21 = int32(1)
	if v18 == v17 {
		v13 = v13 + v21
		v14 = v14 + v21
		goto L4
	} else {
		goto L7
	}
L7:
	;
	goto L5
}
func F_strict_word_similarity(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14406(m, l0, int32(2))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_strict_word_similarity_op(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14410(m, l0, int32(_a_F_strict_word_similarity_op_0), int32(3))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_strip_implicit_coercions(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	if l0 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v40
L2:
	;
	v2 = l0
	goto L5
L3:
	;
	goto L4
L4:
	;
	v40 = int32(0)
	goto L1
L5:
	;
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
	switch v3 - int32(15) {
	case 0:
		goto L13
	default:
		v40 = v2
		goto L1
	case 12:
		goto L12
	case 13:
		goto L11
	case 14:
		goto L10
	case 15:
		goto L9
	case 40:
		goto L8
	}
L6:
	;
	goto L4
L7:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	if v37 != 0 {
		v2 = v37
		goto L5
	} else {
		goto L20
	}
L8:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v2)+20))
	if v31 != int32(2) {
		v40 = v2
		goto L1
	} else {
		goto L19
	}
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v2)+12))
	if v26 != int32(2) {
		v40 = v2
		goto L1
	} else {
		goto L18
	}
L10:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v2)+24))
	if v21 != int32(2) {
		v40 = v2
		goto L1
	} else {
		goto L17
	}
L11:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v2)+16))
	if v16 != int32(2) {
		v40 = v2
		goto L1
	} else {
		goto L16
	}
L12:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v2)+20))
	if v11 != int32(2) {
		v40 = v2
		goto L1
	} else {
		goto L15
	}
L13:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v2)+16))
	if v6 != int32(2) {
		v40 = v2
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v2)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v36 = v10
	goto L7
L15:
	;
	v36 = v2 + int32(4)
	goto L7
L16:
	;
	v36 = v2 + int32(4)
	goto L7
L17:
	;
	v36 = v2 + int32(4)
	goto L7
L18:
	;
	v36 = v2 + int32(4)
	goto L7
L19:
	;
	v36 = v2 + int32(4)
	goto L7
L20:
	;
	goto L6
}
func F_strip_noop_phvs_mutator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	v3 = int32(0)
	if l0 == v3 {
		v21 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v21
L2:
	;
	v6 = l0
	goto L3
L3:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	if v9 != int32(321) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v15 = F_expression_tree_mutator_impl(m, v6, int32(932), l1)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	goto L4
L6:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	if v12 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	if v13 != 0 {
		v6 = v13
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v21 = v3
	goto L1
L9:
	;
	return int32(0)
L10:
	;
	v21 = v15
	goto L1
}
func F_strlen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	if l0&int32(3) == int32(0) {
		v26 = l0
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v51 - l0
L2:
	;
	v30 = v26
	goto L11
L3:
	;
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v9 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L6
L6:
	;
	v15 = l0
	goto L7
L7:
	;
	v19 = v15 + int32(1)
	if v19&int32(3) == int32(0) {
		v26 = v19
		goto L2
	} else {
		goto L9
	}
L8:
	;
	v51 = v19
	goto L1
L9:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v24 != 0 {
		v15 = v19
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v39 = int32(-2139062144)
	if (int32(16843008)-v36|v36)&v39 == v39 {
		v30 = v30 + int32(4)
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v45 = v30
	goto L14
L13:
	;
	goto L12
L14:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v49 != 0 {
		v45 = v45 + int32(1)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v51 = v45
	goto L1
L16:
	;
	goto L15
}
func F_strpbrk(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1))))
	if v10 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v69 = v61 - l0 + l0
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v71 != 0 {
		goto L19
	} else {
		goto L20
	}
L2:
	;
	m.G0 = v8 + int32(32)
	goto L1
L3:
	;
	F___memset(m, v8, int32(0), int32(32))
	mBase = m.M
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v16 != 0 {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v11 != 0 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v12 = F___strchrnul(m, l0, v10)
	mBase = m.M
	v61 = v12
	goto L2
L7:
	;
	goto L6
L8:
	;
	v18 = l1
	v19 = v16
	goto L11
L9:
	;
	goto L10
L10:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v40 == int32(0) {
		v61 = l0
		goto L2
	} else {
		goto L14
	}
L11:
	;
	v26 = v8 + int32(base.Ui32(v19)>>(uint(int32(3))%32))&int32(28)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v28 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v27 | v28<<(uint(v19)%32)
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	if v32 != 0 {
		v18 = v18 + v28
		v19 = v32
		goto L11
	} else {
		goto L13
	}
L12:
	;
	goto L10
L13:
	;
	goto L12
L14:
	;
	v44 = l0
	v45 = v40
	goto L15
L15:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(base.Ui32(v45)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v53)>>(uint(v45)%32))&int32(1) != 0 {
		v61 = v44
		goto L2
	} else {
		goto L17
	}
L16:
	;
	v61 = v59
	goto L2
L17:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+1)))
	v59 = v44 + int32(1)
	if v57 != 0 {
		v44 = v59
		v45 = v57
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v72 = v69
	goto L21
L20:
	;
	v72 = int32(0)
	goto L21
L21:
	;
	return v72
}
func F_strtod(m *base.Module, l0 int32, l1 int32) float64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int64
	_ = v28
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
	var v42 int64
	_ = v42
	var v47 int64
	_ = v47
	var v57 int64
	_ = v57
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int64
	_ = v92
	var v96 int64
	_ = v96
	var v104 int64
	_ = v104
	var v105 int64
	_ = v105
	var v109 int32
	_ = v109
	var v111 int64
	_ = v111
	var v114 int64
	_ = v114
	var v117 int64
	_ = v117
	var v121 int64
	_ = v121
	var v131 int64
	_ = v131
	var v135 int32
	_ = v135
	var v136 int64
	_ = v136
	var v138 int64
	_ = v138
	var v145 int64
	_ = v145
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	F_strtox_1(m, v7, l0, l1, int32(1))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return float64(0)
	} else {
		v14 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
		v15 = *(*int64)(unsafe.Add(mBase, uint32(v7)+8))
		v23 = m.G0
		v25 = v23 - int32(32)
		m.G0 = v25
		v28 = v15 & int64(281474976710655)
		v32 = int64(base.Ui64(v15)>>(uint(int64(48))%64)) & int64(32767)
		v33 = base.I32_wrap_i64(v32)
		if base.Ui32(v33-int32(_a_F_strtod_0)) <= base.Ui32(int32(2045)) {
			v42 = v28<<(uint(int64(4))%64) | int64(base.Ui64(v14)>>(uint(int64(60))%64))
			v47 = v14 & int64(1152921504606846975)
			if base.Ui64(int64(576460752303423489)) <= base.Ui64(v47) {
				v57 = v42 + int64(1)
			} else {
				if v47 != int64(576460752303423488) {
					v57 = v42
				} else {
					v57 = v42&int64(1) + v42
				}
			}
			v60 = base.B2i32(base.Ui64(int64(4503599627370495)) < base.Ui64(v57))
			if base.Ui64(int64(4503599627370495)) < base.Ui64(v57) {
				v61 = int64(0)
			} else {
				v61 = v57
			}
			v138 = v61
			v145 = base.I64_extend_i32_u(v60) + base.I64_extend_i32_u(v33-int32(_a_F_strtod_1))
		} else {
			if base.B2i32(v14|v28 == int64(0))|base.B2i32(v32 != int64(32767)) == int32(0) {
				v138 = v28<<(uint(int64(4))%64) | int64(base.Ui64(v14)>>(uint(int64(60))%64)) | int64(2251799813685248)
				v145 = int64(2047)
			} else {
				if base.Ui32(int32(_a_F_strtod_2)) < base.Ui32(v33) {
					v138 = int64(0)
					v145 = int64(2047)
				} else {
					v87 = base.B2i32(v32 == int64(0))
					if v32 == int64(0) {
						v88 = int32(_a_F_strtod_1)
					} else {
						v88 = int32(_a_F_strtod_0)
					}
					v89 = v88 - v33
					if int32(112) < v89 {
						v92 = int64(0)
						v138 = v92
						v145 = v92
					} else {
						if v32 == int64(0) {
							v96 = v28
						} else {
							v96 = v28 | int64(281474976710656)
						}
						if v33 != v88 {
							F___ashlti3(m, v25+int32(16), v14, v96, int32(128)-v89)
							mBase = m.M
							v104 = *(*int64)(unsafe.Add(mBase, uint32(v25)+16))
							v105 = *(*int64)(unsafe.Add(mBase, uint32(v25)+24))
							v109 = base.B2i32(v104|v105 != int64(0))
						} else {
							v109 = int32(0)
						}
						F___lshrti3(m, v25, v14, v96, v89)
						mBase = m.M
						v111 = *(*int64)(unsafe.Add(mBase, uint32(v25)+8))
						v114 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
						v117 = v111<<(uint(int64(4))%64) | int64(base.Ui64(v114)>>(uint(int64(60))%64))
						v121 = base.I64_extend_i32_u(v109) | v114&int64(1152921504606846975)
						if base.Ui64(int64(576460752303423489)) <= base.Ui64(v121) {
							v131 = v117 + int64(1)
						} else {
							if v121 != int64(576460752303423488) {
								v131 = v117
							} else {
								v131 = v117&int64(1) + v117
							}
						}
						v135 = base.B2i32(base.Ui64(int64(4503599627370495)) < base.Ui64(v131))
						if base.Ui64(int64(4503599627370495)) < base.Ui64(v131) {
							v136 = v131 ^ int64(4503599627370496)
						} else {
							v136 = v131
						}
						v138 = v136
						v145 = base.I64_extend_i32_u(v135)
					}
				}
			}
		}
		m.G0 = v25 + int32(32)
		m.G0 = v7 + int32(16)
		return base.F64_reinterpret_i64(v15&int64(-9223372036854775807-1) | v145<<(uint(int64(52))%64) | v138)
	}
}
func F_strtok_r(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	if l0 != 0 {
		v7 = l0
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = int32(_a_F_strtok_r_0)
	v12 = m.G0
	v14 = v12 - int32(32)
	v15 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v14))) = v15
	v23 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_strtok_r[0])))
	if v23 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v4 != 0 {
		v7 = v4
		goto L1
	} else {
		goto L3
	}
L3:
	;
	return int32(0)
L4:
	;
	v92 = v91 + v7
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	if v93 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L5:
	;
	v91 = int32(0)
	goto L4
L6:
	;
	goto L7
L7:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_strtok_r[1])))
	if v27 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v31 = v7
	goto L11
L9:
	;
	goto L10
L10:
	;
	v41 = v8
	v42 = v23
	goto L14
L11:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if v37 == v23 {
		v31 = v31 + int32(1)
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v91 = v31 - v7
	goto L4
L13:
	;
	goto L12
L14:
	;
	v49 = v14 + int32(base.Ui32(v42)>>(uint(int32(3))%32))&int32(28)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v51 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = v50 | v51<<(uint(v42)%32)
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
	if v55 != 0 {
		v41 = v41 + v51
		v42 = v55
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v58 == int32(0) {
		v81 = v7
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L15
L17:
	;
	v91 = v81 - v7
	goto L4
L18:
	;
	v62 = v7
	v63 = v58
	goto L19
L19:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(base.Ui32(v63)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v71)>>(uint(v63)%32))&int32(1) == int32(0) {
		v81 = v62
		goto L17
	} else {
		goto L21
	}
L20:
	;
	v81 = v79
	goto L17
L21:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)))
	v79 = v62 + int32(1)
	if v77 != 0 {
		v62 = v79
		v63 = v77
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v96 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v96
	return v96
L24:
	;
	goto L25
L25:
	;
	v100 = int32(_a_F_strtok_r_0)
	v104 = m.G0
	v106 = v104 - int32(32)
	m.G0 = v106
	v108 = int32(*(*int8)(unsafe.Add(mBase, _c_F_strtok_r[0])))
	if v108 != 0 {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v167 = v159 - v92 + v92
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
	if v168 != 0 {
		goto L44
	} else {
		goto L45
	}
L27:
	;
	m.G0 = v106 + int32(32)
	goto L26
L28:
	;
	F___memset(m, v106, int32(0), int32(32))
	mBase = m.M
	v114 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_strtok_r[0])))
	if v114 != 0 {
		goto L33
	} else {
		goto L34
	}
L29:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_strtok_r[1])))
	if v109 != 0 {
		goto L28
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v110 = F___strchrnul(m, v92, v108)
	mBase = m.M
	v159 = v110
	goto L27
L32:
	;
	goto L31
L33:
	;
	v116 = v100
	v117 = v114
	goto L36
L34:
	;
	goto L35
L35:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	if v138 == int32(0) {
		v159 = v92
		goto L27
	} else {
		goto L39
	}
L36:
	;
	v124 = v106 + int32(base.Ui32(v117)>>(uint(int32(3))%32))&int32(28)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	v126 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v124))) = v125 | v126<<(uint(v117)%32)
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+1)))
	if v130 != 0 {
		v116 = v116 + v126
		v117 = v130
		goto L36
	} else {
		goto L38
	}
L37:
	;
	goto L35
L38:
	;
	goto L37
L39:
	;
	v142 = v92
	v143 = v138
	goto L40
L40:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v106+int32(base.Ui32(v143)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v151)>>(uint(v143)%32))&int32(1) != 0 {
		v159 = v142
		goto L27
	} else {
		goto L42
	}
L41:
	;
	v159 = v157
	goto L27
L42:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+1)))
	v157 = v142 + int32(1)
	if v155 != 0 {
		v142 = v157
		v143 = v155
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v167 + int32(1)
	v172 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v167))) = uint8(v172)
	return v92
L45:
	;
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	return v92
}
func F_strtoll(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	var v5 int64
	_ = v5
	v5 = F_strtox_2(m, l0, l1, l2, int64(-9223372036854775807-1))
	return v5
}
func F_strtox_1(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	v7 = int64(0)
	v9 = m.G0
	v11 = v9 - int32(160)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = int32(-1)
	v18 = v11 + int32(16)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+112)) = v7
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+44))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+120)) = base.I64_extend_i32_s(v22 - v23)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if int32(1)|base.B2i32(base.I64_extend_i32_s(v29-v23) <= v7) != 0 {
		v36 = v29
	} else {
		v36 = v23 + base.I32_wrap_i64(v7)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v18)+104)) = v36
	F___floatscan(m, v11, v18, l3, int32(1))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		return
	} else {
		v41 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
		v42 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
		if l2 != 0 {
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v11)+136))
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
			v45 = *(*int32)(unsafe.Add(mBase, uint32(v11)+60))
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v43 + (l1 + (v44 - v45))
		} else {
		}
		*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v41
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v42
		m.G0 = v11 + int32(160)
		return
	}
}
func F_subltree(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
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
	var v16 int32
	_ = v16
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_pg_detoast_datum(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		v11 = F_inner_subltree(m, v5, v9, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v13 != v5 {
				F_pfree(m, v5)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int64(0)
				} else {
					return base.I64_extend_i32_u(v11)
				}
			} else {
				return base.I64_extend_i32_u(v11)
			}
		}
	}
}
func F_substitute_actual_parameters_in_from_mutator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l0 == int32(0) {
		v54 = int32(0)
		m.G0 = v7 + int32(16)
		return v54
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v12 != int32(8) {
			if v12 != int32(67) {
				v52 = F_expression_tree_mutator_impl(m, l0, int32(923), l1)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					v54 = v52
					m.G0 = v7 + int32(16)
					return v54
				}
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v17 + int32(1)
				v23 = F_query_tree_mutator_impl(m, l0, int32(923), l1, int32(0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v27 - int32(1)
					v54 = v23
					m.G0 = v7 + int32(16)
					return v54
				}
			}
		} else {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v31 != 0 {
				v52 = F_expression_tree_mutator_impl(m, l0, int32(923), l1)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					v54 = v52
					m.G0 = v7 + int32(16)
					return v54
				}
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v32 <= int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v64
						F_errmsg_internal(m, int32(_a_F_substitute_actual_parameters_in_from_mutator_0), v7)
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_substitute_actual_parameters_in_from_mutator_1), int32(_a_F_substitute_actual_parameters_in_from_mutator_2), int32(_a_F_substitute_actual_parameters_in_from_mutator_3))
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					if v35 < v32 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v64
							F_errmsg_internal(m, int32(_a_F_substitute_actual_parameters_in_from_mutator_0), v7)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_substitute_actual_parameters_in_from_mutator_1), int32(_a_F_substitute_actual_parameters_in_from_mutator_2), int32(_a_F_substitute_actual_parameters_in_from_mutator_3))
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v38+v32<<(uint(int32(2))%32)-int32(4))))
						v45 = F_copyObjectImpl(m, v44)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							F_IncrementVarSublevelsUp(m, v45, v47, int32(0))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								v54 = v45
								m.G0 = v7 + int32(16)
								return v54
							}
						}
					}
				}
			}
		}
	}
}
func F_substitute_grouped_columns_mutator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v412 int32
	_ = v412
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	if l0 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(32)
	return v466
L2:
	;
	v466 = int32(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v16 - int32(9) {
	case 0:
		goto L8
	case 1:
		goto L7
	default:
		goto L6
	}
L5:
	;
	v456 = F_copyObjectImpl(m, l0)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L22
	} else {
		goto L127
	}
L6:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+28)))
	if v28 != int32(1) {
		v155 = v16
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v24 <= v23 {
		v466 = l0
		goto L1
	} else {
		goto L11
	}
L8:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v19 == v20 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	if v20 < v19 {
		v466 = l0
		goto L1
	} else {
		goto L10
	}
L10:
	;
	goto L6
L11:
	;
	goto L6
L12:
	;
	switch v155 - int32(6) {
	case 0:
		goto L45
	case 1, 2:
		v466 = l0
		goto L1
	default:
		goto L44
	}
L13:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v31 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v41 = v33 + v31<<(uint(int32(2))%32) - int32(4)
	goto L16
L15:
	;
	v41 = l1 + int32(12)
	goto L16
L16:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if v42 == int32(0) {
		v155 = v16
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v48 = int32(0)
	goto L19
L18:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v155 = v152
	goto L12
L19:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v54 <= v48 {
		goto L18
	} else {
		goto L21
	}
L20:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+56))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+16))
	v76 = v73 + v59<<(uint(int32(5))%32)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v76-int32(32))))
	v82 = int32(*(*int16)(unsafe.Add(mBase, uint32(v76-int32(28)))))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v76-int32(24))))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v76-int32(20))))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v76-int32(16))))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v93 = F_makeVar(m, v79, v82, v85, v88, v91, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L22
	} else {
		goto L25
	}
L21:
	;
	v59 = v48 + int32(1)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v48<<(uint(int32(2))%32)+v60)))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	v64 = F_equal(m, l0, v63)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	return int32(0)
L23:
	;
	if v64 == int32(0) {
		v48 = v59
		goto L19
	} else {
		goto L24
	}
L24:
	;
	goto L20
L25:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v76-int32(8))))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+36)) = v97
	v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v76-int32(4)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v93)+40)) = uint16(v101)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+108))
	if v104 == int32(0) {
		v466 = v93
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v108 = int32(0)
	if v107 == v108 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if v146 != 0 {
		v466 = v93
		goto L1
	} else {
		goto L40
	}
L28:
	;
	v146 = int32(0)
	goto L27
L29:
	;
	goto L30
L30:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	if v114 <= int32(0) {
		v140 = v108
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v146 = v140
	goto L27
L32:
	;
	v117 = int32(0)
	if v117 < v114 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v120 = v114
	goto L35
L34:
	;
	v120 = v117
	goto L35
L35:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v107)+12))
	v123 = int32(0)
	goto L36
L36:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v121+v123<<(uint(int32(2))%32))))
	v132 = base.B2i32(v131 == v70)
	if v131 == v70 {
		v140 = v132
		goto L31
	} else {
		goto L38
	}
L37:
	;
	v140 = v132
	goto L31
L38:
	;
	v134 = v123 + int32(1)
	if v134 != v120 {
		v123 = v134
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v93)+24))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
	v149 = F_bms_add_member(m, v147, v148)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L22
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+24)) = v149
	v466 = v93
	goto L1
L42:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v425 = int32(1)
	v426 = v424 + v425
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v426
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+28)))
	if v428 != v425 {
		goto L117
	} else {
		goto L118
	}
L43:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v192)+16))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)+56))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v344)+16))
	v348 = v345 + v188<<(uint(int32(5))%32)
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v348-int32(32))))
	v354 = int32(*(*int16)(unsafe.Add(mBase, uint32(v348-int32(28)))))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v348-int32(24))))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v348-int32(20))))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v348-int32(16))))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v365 = F_makeVar(m, v351, v354, v357, v360, v363, v364)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L22
	} else {
		goto L99
	}
L44:
	;
	if v155 == int32(67) {
		goto L42
	} else {
		goto L97
	}
L45:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v163 != v164 {
		v466 = l0
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+28)))
	if v166 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v218 = int32(0)
	if v216 == v218 {
		goto L63
	} else {
		goto L64
	}
L48:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v167 == int32(0) {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	if v170 <= int32(0) {
		goto L47
	} else {
		goto L50
	}
L50:
	;
	v173 = int32(0)
	if v173 < v170 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v177 = v170
	goto L53
L52:
	;
	v177 = v173
	goto L53
L53:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v167)+12))
	v181 = v173
	goto L54
L54:
	;
	v188 = v181 + int32(1)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v178+v181<<(uint(int32(2))%32))))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
	if v194 != int32(6) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L47
L56:
	;
	if v188 != v177 {
		v181 = v188
		goto L54
	} else {
		goto L61
	}
L57:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v197 != v198 {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v200 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v193)+8)))
	v201 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
	if v200 != v201 {
		goto L56
	} else {
		goto L59
	}
L59:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v193)+28))
	if v203 == int32(0) {
		goto L43
	} else {
		goto L60
	}
L60:
	;
	goto L56
L61:
	;
	goto L55
L62:
	;
	if v256 != 0 {
		v466 = l0
		goto L1
	} else {
		goto L75
	}
L63:
	;
	v256 = int32(0)
	goto L62
L64:
	;
	goto L65
L65:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
	if v224 <= int32(0) {
		v250 = v218
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v256 = v250
	goto L62
L67:
	;
	v227 = int32(0)
	if v227 < v224 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v230 = v224
	goto L70
L69:
	;
	v230 = v227
	goto L70
L70:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v216)+12))
	v233 = int32(0)
	goto L71
L71:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v231+v233<<(uint(int32(2))%32))))
	v242 = base.B2i32(v241 == v217)
	if v241 == v217 {
		v250 = v242
		goto L66
	} else {
		goto L73
	}
L72:
	;
	v250 = v242
	goto L66
L73:
	;
	v244 = v233 + int32(1)
	if v244 != v230 {
		v233 = v244
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v257)+8))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v258)+12))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v259+v260<<(uint(int32(2))%32)-int32(4))))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v266)+12))
	if v267 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v285 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
	v286 = F_get_rte_attribute_name(m, v266, v285)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L22
	} else {
		goto L81
	}
L77:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v266)+16))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v274 = F_check_functional_grouping(m, v268, v260, int32(0), v270, v271+int32(148))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L22
	} else {
		goto L78
	}
L78:
	;
	if v274 == int32(0) {
		goto L76
	} else {
		goto L79
	}
L79:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v278)))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v281 = F_lappend_int(m, v279, v280)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L22
	} else {
		goto L80
	}
L80:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v283))) = v281
	v466 = l0
	goto L1
L81:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L22
	} else {
		goto L82
	}
L82:
	;
	F_errcode(m, int32(50364548))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L22
	} else {
		goto L83
	}
L83:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v266)+8))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v296)+4))
	if v288 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v286
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v297
	F_errmsg(m, int32(_a_F_substitute_grouped_columns_mutator_0), v11)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L22
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v286
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v297
	F_errmsg(m, int32(_a_F_substitute_grouped_columns_mutator_1), v11+int32(16))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L22
	} else {
		goto L94
	}
L87:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+40)))
	if v305 == int32(1) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v310 = F_errdetail(m, int32(_a_F_substitute_grouped_columns_mutator_2), int32(0))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L22
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	F_parser_errposition(m, v312, v313)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L22
	} else {
		goto L92
	}
L91:
	;
	goto L90
L92:
	;
	F_errfinish(m, int32(_a_F_substitute_grouped_columns_mutator_3), int32(1557), int32(_a_F_substitute_grouped_columns_mutator_4))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L22
	} else {
		goto L93
	}
L93:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L94:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	F_parser_errposition(m, v328, v329)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L22
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_substitute_grouped_columns_mutator_3), int32(1563), int32(_a_F_substitute_grouped_columns_mutator_4))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L22
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L97:
	;
	v340 = F_expression_tree_mutator_impl(m, l0, int32(516), l1)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L22
	} else {
		goto L98
	}
L98:
	;
	v466 = v340
	goto L1
L99:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v348-int32(8))))
	*(*int32)(unsafe.Add(mBase, uint32(v365)+36)) = v369
	v373 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v348-int32(4)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v365)+40)) = uint16(v373)
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v375)+108))
	if v376 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v466 = v365
	goto L1
L101:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v380 = int32(0)
	if v379 == v380 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	if v418 != 0 {
		goto L100
	} else {
		goto L115
	}
L103:
	;
	v418 = int32(0)
	goto L102
L104:
	;
	goto L105
L105:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v379)+4))
	if v386 <= int32(0) {
		v412 = v380
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v418 = v412
	goto L102
L107:
	;
	v389 = int32(0)
	if v389 < v386 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v392 = v386
	goto L110
L109:
	;
	v392 = v389
	goto L110
L110:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v379)+12))
	v395 = int32(0)
	goto L111
L111:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v393+v395<<(uint(int32(2))%32))))
	v404 = base.B2i32(v403 == v342)
	if v403 == v342 {
		v412 = v404
		goto L106
	} else {
		goto L113
	}
L112:
	;
	v412 = v404
	goto L106
L113:
	;
	v406 = v395 + int32(1)
	if v406 != v392 {
		v395 = v406
		goto L111
	} else {
		goto L114
	}
L114:
	;
	goto L112
L115:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v365)+24))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v344)+8))
	v421 = F_bms_add_member(m, v419, v420)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L22
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v365)+24)) = v421
	goto L100
L117:
	;
	v450 = F_query_tree_mutator_impl(m, l0, int32(516), l1, int32(0))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L22
	} else {
		goto L126
	}
L118:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v431 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v431)+4))
	v434 = v432
	goto L121
L120:
	;
	v434 = int32(0)
	goto L121
L121:
	;
	if v426 <= v434 {
		goto L117
	} else {
		goto L122
	}
L122:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v437 = F_copyObjectImpl(m, v436)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L22
	} else {
		goto L123
	}
L123:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	F_IncrementVarSublevelsUp(m, v437, v439, int32(0))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L22
	} else {
		goto L124
	}
L124:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v444 = F_lappend(m, v443, v437)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L22
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v444
	goto L117
L126:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v452 - int32(1)
	v466 = v450
	goto L1
L127:
	;
	v458 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+40)) = uint8(v458)
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v456)+28))
	v461 = F_substitute_grouped_columns_mutator(m, v460, l1)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L22
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v456)+28)) = v461
	v464 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+40)) = uint8(v464)
	v466 = v456
	goto L1
}
func F_switchToPresortedPrefixMode(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int64
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int64
	_ = v134
	var v135 int64
	_ = v135
	var v147 int64
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int64
	_ = v177
	var v178 int64
	_ = v178
	var v187 int64
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int64
	_ = v193
	var v200 int64
	_ = v200
	var v201 int64
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int64
	_ = v238
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v249 int64
	_ = v249
	var v250 int64
	_ = v250
	var v252 int32
	_ = v252
	var v253 int64
	_ = v253
	var v255 int64
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int64
	_ = v264
	var v270 int64
	_ = v270
	var v274 int32
	_ = v274
	var v284 int32
	_ = v284
	var v287 int64
	_ = v287
	var v291 int64
	_ = v291
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int64
	_ = v310
	var v311 int64
	_ = v311
	var v315 int32
	_ = v315
	var v316 int64
	_ = v316
	var v320 int32
	_ = v320
	var v321 int64
	_ = v321
	var v322 int64
	_ = v322
	var v326 int32
	_ = v326
	var v327 int64
	_ = v327
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int64
	_ = v339
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v350 int64
	_ = v350
	var v351 int64
	_ = v351
	var v353 int32
	_ = v353
	var v354 int64
	_ = v354
	var v356 int64
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int64
	_ = v365
	var v371 int64
	_ = v371
	var v375 int32
	_ = v375
	var v385 int32
	_ = v385
	var v388 int64
	_ = v388
	var v392 int64
	_ = v392
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int64
	_ = v409
	var v410 int64
	_ = v410
	var v413 int64
	_ = v413
	var v416 int64
	_ = v416
	var v417 int64
	_ = v417
	var v420 int64
	_ = v420
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v432 int32
	_ = v432
	var v435 int64
	_ = v435
	var v436 int64
	_ = v436
	var v437 int64
	_ = v437
	var v439 int64
	_ = v439
	v7 = int64(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+56))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	if v18 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)))
	if v51 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v13)+96))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
	v25 = int32(1)
	v28 = int32(2)
	v29 = v22 << (uint(v28) % 32)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v13)+80))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v13)+84))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v13)+88))
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_switchToPresortedPrefixMode[0]))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)))
	v44 = F_tuplesort_begin_heap(m, v17, v21-v22, v24+v22<<(uint(v25)%32), v29+v30, v32+v29, v34+v22, v37, int32(0), v39<<(uint(v25)%32)&v28)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	F_tuplesort_reset(m, v18)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L5
	} else {
		goto L7
	}
L5:
	;
	return
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v44
	goto L1
L7:
	;
	goto L1
L8:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v55 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	v56 = *(*int64)(unsafe.Add(mBase, uint32(l0)+136))
	v57 = v55 - v56
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v54)+236))
	if v60 != 0 {
		goto L14
	} else {
		goto L15
	}
L9:
	;
	goto L10
L10:
	;
	v86 = *(*int64)(unsafe.Add(mBase, uint32(l0)+152))
	if v86 <= int64(0) {
		v200 = v86
		v201 = v7
		goto L23
	} else {
		goto L24
	}
L11:
	;
	goto L10
L12:
	;
	goto L11
L13:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v54)+72)) = uint32(v57)
	v69 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v54)+68)) = uint8(v69)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v54)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v71)+24)) = int32(0)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v54)+44))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+32))
	if v75 != 0 {
		goto L20
	} else {
		goto L21
	}
L14:
	;
	if int64(1073741823) < v57 {
		goto L12
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	if int64(1073741823) < v57 {
		goto L12
	} else {
		goto L19
	}
L17:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v54)+232))
	if v63 != int32(-1) {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	goto L12
L19:
	;
	goto L13
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74)+16)) = v75
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v54)+44))
	v78 = v77
	goto L22
L21:
	;
	v78 = v74
	goto L22
L22:
	;
	v79 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v78)+28)) = v79
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v54)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v81)+32)) = v79
	goto L12
L23:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+152)) = v200 - v201
	if v200 == v201 {
		goto L58
	} else {
		goto L59
	}
L24:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	if v89 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)+8))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+12))
	m.T0[v190].(func(*base.Module, int32))(m, v188)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L5
	} else {
		goto L56
	}
L26:
	;
	v134 = int64(1)
	v135 = *(*int64)(unsafe.Add(mBase, uint32(l0)+152))
	if v135 < int64(2) {
		v200 = v135
		v201 = v134
		goto L23
	} else {
		goto L42
	}
L27:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v107 = int32(0)
	v109 = F_tuplesort_gettupleslot(m, v104, base.B2i32(v15 == int32(1)), v107, v89, v107)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L5
	} else {
		goto L32
	}
L28:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+4)))
	if v92&int32(2) != 0 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	F_tuplesort_puttupleslot(m, v95, v89)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+32))
	m.T0[v101].(func(*base.Module, int32, int32))(m, v98, v99)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	goto L26
L32:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v111 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	v125 = F_isCurrentGroup(m, l0, v123, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L5
	} else {
		goto L39
	}
L34:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+4)))
	if v112&int32(2) == int32(0) {
		v123 = v111
		goto L33
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v111)+8))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+32))
	m.T0[v119].(func(*base.Module, int32, int32))(m, v111, v117)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L5
	} else {
		goto L38
	}
L37:
	;
	goto L36
L38:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v123 = v122
	goto L33
L39:
	;
	if v125 == int32(0) {
		v187 = v7
		goto L25
	} else {
		goto L40
	}
L40:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	F_tuplesort_puttupleslot(m, v129, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L5
	} else {
		goto L41
	}
L41:
	;
	goto L26
L42:
	;
	v147 = v134
	goto L43
L43:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v149 = int32(0)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	v152 = F_tuplesort_gettupleslot(m, v148, base.B2i32(v15 == int32(1)), v149, v150, v149)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L5
	} else {
		goto L45
	}
L44:
	;
	v200 = v178
	v201 = v177
	goto L23
L45:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v154 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	v168 = F_isCurrentGroup(m, l0, v166, v167)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L5
	} else {
		goto L52
	}
L47:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+4)))
	if v155&int32(2) == int32(0) {
		v166 = v154
		goto L46
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v154)+8))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)+32))
	m.T0[v162].(func(*base.Module, int32, int32))(m, v154, v160)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L5
	} else {
		goto L51
	}
L50:
	;
	goto L49
L51:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v166 = v165
	goto L46
L52:
	;
	if v168 == int32(0) {
		v187 = v147
		goto L25
	} else {
		goto L53
	}
L53:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	F_tuplesort_puttupleslot(m, v172, v173)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L5
	} else {
		goto L54
	}
L54:
	;
	v177 = v147 + int64(1)
	v178 = *(*int64)(unsafe.Add(mBase, uint32(l0)+152))
	if v177 < v178 {
		v147 = v177
		goto L43
	} else {
		goto L55
	}
L55:
	;
	goto L44
L56:
	;
	v193 = *(*int64)(unsafe.Add(mBase, uint32(l0)+152))
	v200 = v193
	v201 = v187
	goto L23
L57:
	;
	m.G0 = v11 + int32(16)
	return
L58:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v205)+8))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+32))
	m.T0[v208].(func(*base.Module, int32, int32))(m, v205, v206)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L5
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	F_tuplesort_performsort(m, v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L5
	} else {
		goto L63
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = int32(1)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)+8))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v214)+12))
	m.T0[v215].(func(*base.Module, int32))(m, v213)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L5
	} else {
		goto L62
	}
L62:
	;
	goto L57
L63:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v221 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)))
	if v432 == int32(1) {
		goto L119
	} else {
		goto L120
	}
L65:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	if v224 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v339 = *(*int64)(unsafe.Add(mBase, uint32(l0)+224))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+224)) = v339 + int64(1)
	v343 = int32(0)
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v338)+128))
	if v347 == v343 {
		goto L97
	} else {
		goto L98
	}
L67:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+280)))
	if v227 != int32(1) {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v232 = *(*int32)(unsafe.Add(mBase, _c_F_switchToPresortedPrefixMode[1]))
	v235 = v224 + v232*int32(96)
	v237 = v235 + int32(56)
	v238 = *(*int64)(unsafe.Add(mBase, uint32(v237)))
	*(*int64)(unsafe.Add(mBase, uint32(v237))) = v238 + int64(1)
	v242 = int32(0)
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v230)+128))
	if v246 == v242 {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	switch v307 {
	case 0:
		goto L91
	case 1:
		goto L90
	default:
		goto L89
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v284
	v287 = *(*int64)(unsafe.Add(mBase, uint32(v230)+112))
	v291 = base.I64_div_s(v287+int64(1023), int64(1024))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v291
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v230)+124))
	switch v293 - int32(3) {
	case 0:
		goto L85
	case 1:
		v304 = v293
		goto L82
	case 2:
		goto L84
	default:
		goto L83
	}
L71:
	;
	if v263&int32(255) != base.B2i32(v246 != int32(0)) {
		goto L77
	} else {
		goto L78
	}
L72:
	;
	v249 = *(*int64)(unsafe.Add(mBase, uint32(v230)+96))
	v250 = *(*int64)(unsafe.Add(mBase, uint32(v230)+88))
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+120)))
	v263 = v252
	v264 = v249 - v250
	goto L71
L73:
	;
	goto L74
L74:
	;
	v253 = F_LogicalTapeSetBlocks(m, v246)
	mBase = m.M
	v255 = v253 << (uint(int64(13)) % 64)
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+120)))
	if v257 != 0 {
		v263 = int32(1)
		v264 = v255
		goto L71
	} else {
		goto L75
	}
L75:
	;
	v258 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v230)+120)) = uint8(v258)
	*(*int64)(unsafe.Add(mBase, uint32(v230)+112)) = v255
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v230)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v230)+124)) = v261
	v284 = v242
	goto L70
L76:
	;
	v284 = int32(1)
	goto L70
L77:
	;
	if v263&int32(1) != 0 {
		v284 = v242
		goto L70
	} else {
		goto L81
	}
L78:
	;
	v270 = *(*int64)(unsafe.Add(mBase, uint32(v230)+112))
	if v264 <= v270 {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v230)+120)) = uint8(v263)
	*(*int64)(unsafe.Add(mBase, uint32(v230)+112)) = v264
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v230)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v230)+124)) = v274
	if v263&int32(1) == int32(0) {
		goto L76
	} else {
		goto L80
	}
L80:
	;
	v284 = v242
	goto L70
L81:
	;
	goto L76
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v304
	goto L69
L83:
	;
	v304 = int32(0)
	goto L82
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(8)
	goto L69
L85:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+69)))
	if v298 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v299 = int32(1)
	goto L88
L87:
	;
	v299 = int32(2)
	goto L88
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v299
	goto L69
L89:
	;
	v333 = v235 + int32(96)
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v333)))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	*(*int32)(unsafe.Add(mBase, uint32(v333))) = v334 | v335
	goto L64
L90:
	;
	v320 = v235 + int32(88)
	v321 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	v322 = *(*int64)(unsafe.Add(mBase, uint32(v320)))
	*(*int64)(unsafe.Add(mBase, uint32(v320))) = v321 + v322
	v326 = v235 + int32(80)
	v327 = *(*int64)(unsafe.Add(mBase, uint32(v326)))
	if v321 <= v327 {
		goto L89
	} else {
		goto L93
	}
L91:
	;
	v309 = v235 + int32(72)
	v310 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	v311 = *(*int64)(unsafe.Add(mBase, uint32(v309)))
	*(*int64)(unsafe.Add(mBase, uint32(v309))) = v310 + v311
	v315 = v235 - int32(-64)
	v316 = *(*int64)(unsafe.Add(mBase, uint32(v315)))
	if v310 <= v316 {
		goto L89
	} else {
		goto L92
	}
L92:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v315))) = v310
	goto L89
L93:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v326))) = v321
	goto L89
L94:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	switch v408 {
	case 0:
		goto L116
	case 1:
		goto L115
	default:
		goto L114
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v385
	v388 = *(*int64)(unsafe.Add(mBase, uint32(v338)+112))
	v392 = base.I64_div_s(v388+int64(1023), int64(1024))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v392
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v338)+124))
	switch v394 - int32(3) {
	case 0:
		goto L110
	case 1:
		v405 = v394
		goto L107
	case 2:
		goto L109
	default:
		goto L108
	}
L96:
	;
	if v364&int32(255) != base.B2i32(v347 != int32(0)) {
		goto L102
	} else {
		goto L103
	}
L97:
	;
	v350 = *(*int64)(unsafe.Add(mBase, uint32(v338)+96))
	v351 = *(*int64)(unsafe.Add(mBase, uint32(v338)+88))
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338)+120)))
	v364 = v353
	v365 = v350 - v351
	goto L96
L98:
	;
	goto L99
L99:
	;
	v354 = F_LogicalTapeSetBlocks(m, v347)
	mBase = m.M
	v356 = v354 << (uint(int64(13)) % 64)
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338)+120)))
	if v358 != 0 {
		v364 = int32(1)
		v365 = v356
		goto L96
	} else {
		goto L100
	}
L100:
	;
	v359 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v338)+120)) = uint8(v359)
	*(*int64)(unsafe.Add(mBase, uint32(v338)+112)) = v356
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v338)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v338)+124)) = v362
	v385 = v343
	goto L95
L101:
	;
	v385 = int32(1)
	goto L95
L102:
	;
	if v364&int32(1) != 0 {
		v385 = v343
		goto L95
	} else {
		goto L106
	}
L103:
	;
	v371 = *(*int64)(unsafe.Add(mBase, uint32(v338)+112))
	if v365 <= v371 {
		goto L102
	} else {
		goto L104
	}
L104:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v338)+120)) = uint8(v364)
	*(*int64)(unsafe.Add(mBase, uint32(v338)+112)) = v365
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v338)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v338)+124)) = v375
	if v364&int32(1) == int32(0) {
		goto L101
	} else {
		goto L105
	}
L105:
	;
	v385 = v343
	goto L95
L106:
	;
	goto L101
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v405
	goto L94
L108:
	;
	v405 = int32(0)
	goto L107
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(8)
	goto L94
L110:
	;
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338)+69)))
	if v399 != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v400 = int32(1)
	goto L113
L112:
	;
	v400 = int32(2)
	goto L113
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v400
	goto L94
L114:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+264)) = v424 | v425
	goto L64
L115:
	;
	v416 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	v417 = *(*int64)(unsafe.Add(mBase, uint32(l0)+256))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+256)) = v416 + v417
	v420 = *(*int64)(unsafe.Add(mBase, uint32(l0)+248))
	if v416 <= v420 {
		goto L114
	} else {
		goto L118
	}
L116:
	;
	v409 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	v410 = *(*int64)(unsafe.Add(mBase, uint32(l0)+240))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+240)) = v409 + v410
	v413 = *(*int64)(unsafe.Add(mBase, uint32(l0)+232))
	if v409 <= v413 {
		goto L114
	} else {
		goto L117
	}
L117:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+232)) = v409
	goto L114
L118:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+248)) = v416
	goto L114
L119:
	;
	v435 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	v436 = *(*int64)(unsafe.Add(mBase, uint32(l0)+136))
	v437 = v436 + v201
	if v435 < v437 {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	goto L121
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = int32(3)
	goto L57
L122:
	;
	v439 = v435
	goto L124
L123:
	;
	v439 = v437
	goto L124
L124:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+136)) = v439
	goto L121
}
func F_synchronize_slots(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v227 int64
	_ = v227
	var v229 int64
	_ = v229
	var v231 int64
	_ = v231
	var v233 int64
	_ = v233
	var v235 int64
	_ = v235
	var v237 int64
	_ = v237
	var v239 int64
	_ = v239
	var v241 int64
	_ = v241
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v277 int32
	_ = v277
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int64
	_ = v419
	var v420 int64
	_ = v420
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v454 int64
	_ = v454
	var v456 int64
	_ = v456
	var v458 int64
	_ = v458
	var v460 int64
	_ = v460
	var v462 int64
	_ = v462
	var v464 int64
	_ = v464
	var v466 int64
	_ = v466
	var v468 int64
	_ = v468
	var v470 int32
	_ = v470
	var v473 int64
	_ = v473
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int64
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v497 int64
	_ = v497
	var v498 int32
	_ = v498
	var v502 int64
	_ = v502
	var v505 int64
	_ = v505
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v513 int64
	_ = v513
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v521 int64
	_ = v521
	var v522 int64
	_ = v522
	var v523 int64
	_ = v523
	var v524 int32
	_ = v524
	var v525 int64
	_ = v525
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v604 int32
	_ = v604
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v631 int32
	_ = v631
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v647 int32
	_ = v647
	var v648 int64
	_ = v648
	var v649 int64
	_ = v649
	var v652 int64
	_ = v652
	var v653 int64
	_ = v653
	var v656 int64
	_ = v656
	var v662 int32
	_ = v662
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v677 int32
	_ = v677
	var v682 int32
	_ = v682
	v3 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(144)
	m.G0 = v17
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[0]))
	v24 = F_LWLockAcquire(m, v20+int32(_a_F_synchronize_slots_0), int32(1))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[1]))
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[2]))
	if int32(0) < v29+v31 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[3]))
	v39 = v3
	v40 = v31
	v42 = v3
	v43 = v36
	v44 = v29
	goto L6
L4:
	;
	v81 = v3
	goto L5
L5:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[0]))
	F_LWLockRelease(m, v91+int32(_a_F_synchronize_slots_0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L13
	}
L6:
	;
	v53 = v43 + v39*int32(296)
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+4)))
	if v54 != int32(1) {
		v68 = v40
		v69 = v42
		v70 = v43
		v71 = v44
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v81 = v69
	goto L5
L8:
	;
	v73 = v39 + int32(1)
	if v73 < v68+v71 {
		v39 = v73
		v40 = v68
		v42 = v69
		v43 = v70
		v44 = v71
		goto L6
	} else {
		goto L12
	}
L9:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+201)))
	if v57 != int32(1) {
		v68 = v40
		v69 = v42
		v70 = v43
		v71 = v44
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v60 = F_lappend(m, v42, v53)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[1]))
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[2]))
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[3]))
	v68 = v65
	v69 = v60
	v70 = v67
	v71 = v63
	goto L8
L12:
	;
	goto L7
L13:
	;
	if v81 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if l0 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L15:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	if v98 <= int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v110 = v3
	goto L17
L17:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v115+v110<<(uint(int32(2))%32))))
	v121 = v119 + int32(24)
	if l0 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L14
L19:
	;
	v293 = v110 + int32(1)
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	if v293 < v294 {
		v110 = v293
		goto L17
	} else {
		goto L61
	}
L20:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v119)+88))
	F_LockSharedObject(m, int32(1262), v208, int32(1))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L44
	}
L21:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v124 <= int32(0) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v131 = int32(0)
	goto L23
L23:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v127+v131<<(uint(int32(2))%32))))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
	if base.B2i32(v150 == int32(0))|base.B2i32(v150 != v153) != 0 {
		v171 = v150
		v172 = v153
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v179 = base.AtomicRmwXchg32(m, v119, int32(0), int32(1))
	if v179 != 0 {
		goto L36
	} else {
		goto L37
	}
L25:
	;
	if v171-v172 != 0 {
		goto L32
	} else {
		goto L33
	}
L26:
	;
	goto L25
L27:
	;
	v156 = v147
	v157 = v121
	goto L28
L28:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+1)))
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+1)))
	if v161 == int32(0) {
		v171 = v161
		v172 = v160
		goto L26
	} else {
		goto L30
	}
L29:
	;
	v171 = v161
	v172 = v160
	goto L26
L30:
	;
	v164 = int32(1)
	if v161 == v160 {
		v156 = v156 + v164
		v157 = v157 + v164
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v175 = v131 + int32(1)
	if v175 != v124 {
		v131 = v175
		goto L23
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	goto L24
L35:
	;
	goto L20
L36:
	;
	F_s_lock(m, v119, int32(_a_F_synchronize_slots_1))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v146)+44))
	if v183 != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L38
L40:
	;
	v184 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v119))), uint32(v184))
	goto L19
L41:
	;
	goto L42
L42:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v119)+112))
	v188 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v119))), uint32(v188))
	if v187 == v188 {
		goto L19
	} else {
		goto L43
	}
L43:
	;
	goto L20
L44:
	;
	v214 = base.AtomicRmwXchg32(m, v119, int32(0), int32(1))
	if v214 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	F_s_lock(m, v119, int32(_a_F_synchronize_slots_1))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+4)))
	if v218 == int32(1) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L47
L49:
	;
	F_UnlockSharedObject(m, int32(1262), v208, int32(1))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L60
	}
L50:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+201)))
	v222 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v119))), uint32(v222))
	if v221 != int32(1) {
		goto L49
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v270 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v119))), uint32(v270))
	goto L49
L53:
	;
	v227 = *(*int64)(unsafe.Add(mBase, uint32(v121)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+136)) = v227
	v229 = *(*int64)(unsafe.Add(mBase, uint32(v121)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+128)) = v229
	v231 = *(*int64)(unsafe.Add(mBase, uint32(v121)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+120)) = v231
	v233 = *(*int64)(unsafe.Add(mBase, uint32(v121)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+112)) = v233
	v235 = *(*int64)(unsafe.Add(mBase, uint32(v121)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+104)) = v235
	v237 = *(*int64)(unsafe.Add(mBase, uint32(v121)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+96)) = v237
	v239 = *(*int64)(unsafe.Add(mBase, uint32(v121)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+88)) = v239
	v241 = *(*int64)(unsafe.Add(mBase, uint32(v121)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+80)) = v241
	v244 = v17 + int32(80)
	F_ReplicationSlotAcquire(m, v244, int32(1), int32(0))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_ReplicationSlotDropAcquired(m, int32(0))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v254 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	if v254 == int32(0) {
		goto L49
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+68)) = v208
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = v244
	F_errmsg(m, int32(_a_F_synchronize_slots_2), v17-int32(-64))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_synchronize_slots_3), int32(585), int32(_a_F_synchronize_slots_4))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	goto L49
L60:
	;
	goto L19
L61:
	;
	goto L18
L62:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L1
	} else {
		goto L168
	}
L63:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L1
	} else {
		goto L164
	}
L64:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L1
	} else {
		goto L160
	}
L65:
	;
	m.G0 = v17 + int32(144)
	return v604 & int32(1)
L66:
	;
	v604 = int32(0)
	goto L65
L67:
	;
	goto L68
L68:
	;
	v313 = int32(0)
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v314 <= v313 {
		v604 = v313
		goto L65
	} else {
		goto L69
	}
L69:
	;
	v324 = v313
	v325 = int32(0)
	goto L70
L70:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v333+v325<<(uint(int32(2))%32))))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v337)+8))
	v340 = F_get_database_oid(m, v338, int32(0))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L72
	}
L71:
	;
	v604 = v593
	goto L65
L72:
	;
	F_LockSharedObject(m, int32(1262), v340, int32(1))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v337)))
	v347 = F_SearchNamedReplicationSlot(m, v345, int32(1))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L75
	}
L74:
	;
	F_UnlockSharedObject(m, int32(1262), v340, int32(1))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L1
	} else {
		goto L158
	}
L75:
	;
	if v347 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v351 = base.AtomicRmwXchg32(m, v347, int32(0), int32(1))
	if v351 != 0 {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	goto L78
L78:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v337)+44))
	if v426 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L79:
	;
	F_s_lock(m, v347, int32(_a_F_synchronize_slots_1))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347)+201)))
	v356 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v347))), uint32(v356))
	if v355 == v356 {
		goto L64
	} else {
		goto L83
	}
L82:
	;
	goto L81
L83:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v337)))
	F_ReplicationSlotAcquire(m, v361, int32(1), int32(0))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v347)+112))
	if v366 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v347)+92))
	if v412 == int32(2) {
		goto L106
	} else {
		goto L107
	}
L86:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v337)+44))
	if v369 == int32(0) {
		goto L85
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v391 = *(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[4]))
	F_pgstat_report_replslotsync(m, v391)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L97
	}
L89:
	;
	v374 = base.AtomicRmwXchg32(m, v347, int32(0), int32(1))
	if v374 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	F_s_lock(m, v347, int32(_a_F_synchronize_slots_1))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v337)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v347)+112)) = v378
	v380 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v347))), uint32(v380))
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L94
	}
L93:
	;
	goto L92
L94:
	;
	F_ReplicationSlotSave(m)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v347)+112))
	if v387 == int32(0) {
		goto L85
	} else {
		goto L96
	}
L96:
	;
	goto L88
L97:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v391)+288))
	if v394 != int32(4) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v399 = base.AtomicRmwXchg32(m, v391, int32(0), int32(1))
	if v399 != 0 {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	goto L100
L100:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L105
	}
L101:
	;
	F_s_lock(m, v391, int32(_a_F_synchronize_slots_1))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v391)+288)) = int32(4)
	v405 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v391))), uint32(v405))
	goto L100
L104:
	;
	goto L103
L105:
	;
	v582 = base.B2i32(v366 == int32(0))
	goto L74
L106:
	;
	v415 = F_update_and_persist_local_synced_slot(m, v337, v340, l1)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	v419 = *(*int64)(unsafe.Add(mBase, uint32(v337)+24))
	v420 = *(*int64)(unsafe.Add(mBase, uint32(v347)+120))
	if base.Ui64(v419) < base.Ui64(v420) {
		goto L63
	} else {
		goto L111
	}
L109:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	v582 = v415
	goto L74
L111:
	;
	v422 = F_update_local_synced_slot(m, v337, v340)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	v582 = v422
	goto L74
L114:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v337)))
	v430 = int32(1)
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337)+12)))
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337)+13)))
	F_ReplicationSlotCreate(m, v429, v430, int32(2), v432, int32(0), v434, v430)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L1
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	v582 = int32(0)
	goto L74
L117:
	;
	v439 = *(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[4]))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v337)+4))
	v444 = F_strncpy(m, v17+int32(80), v442, int32(64))
	mBase = m.M
	v445 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v444)+63)) = uint8(v445)
	goto L118
L118:
	;
	v449 = base.AtomicRmwXchg32(m, v439, int32(0), int32(1))
	if v449 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	F_s_lock(m, v439, int32(_a_F_synchronize_slots_1))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L1
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v439)+88)) = v340
	v454 = *(*int64)(unsafe.Add(mBase, uint32(v17)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v439)+137)) = v454
	v456 = *(*int64)(unsafe.Add(mBase, uint32(v17)+88))
	*(*int64)(unsafe.Add(mBase, uint32(v439)+145)) = v456
	v458 = *(*int64)(unsafe.Add(mBase, uint32(v17)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v439)+153)) = v458
	v460 = *(*int64)(unsafe.Add(mBase, uint32(v17)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v439)+161)) = v460
	v462 = *(*int64)(unsafe.Add(mBase, uint32(v17)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v439)+169)) = v462
	v464 = *(*int64)(unsafe.Add(mBase, uint32(v17)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v439)+177)) = v464
	v466 = *(*int64)(unsafe.Add(mBase, uint32(v17)+128))
	*(*int64)(unsafe.Add(mBase, uint32(v439)+185)) = v466
	v468 = *(*int64)(unsafe.Add(mBase, uint32(v17)+136))
	*(*int64)(unsafe.Add(mBase, uint32(v439)+193)) = v468
	v470 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v439))), uint32(v470))
	v473 = *(*int64)(unsafe.Add(mBase, uint32(v337)+16))
	v475 = *(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[4]))
	v477 = *(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[0]))
	v481 = F_LWLockAcquire(m, v477+int32(_a_F_synchronize_slots_5), v470)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L123
	}
L122:
	;
	goto L121
L123:
	;
	v483 = F_GetRedoRecPtr(m)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	v486 = *(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[5]))
	v489 = base.AtomicRmwXchg32(m, v486, int32(440), int32(1))
	if v489 != 0 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	F_s_lock(m, v486+int32(440), int32(_a_F_synchronize_slots_1))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	v496 = *(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[5]))
	v497 = *(*int64)(unsafe.Add(mBase, uint32(v496)+216))
	v498 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v496)+440)), uint32(v498))
	if base.Ui64(v483) < base.Ui64(v497) {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	goto L127
L129:
	;
	v502 = v483
	goto L131
L130:
	;
	v502 = v497
	goto L131
L131:
	;
	if v497 == int64(0) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v505 = v483
	goto L134
L133:
	;
	v505 = v502
	goto L134
L134:
	;
	v508 = base.AtomicRmwXchg32(m, v475, int32(0), int32(1))
	if v508 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	F_s_lock(m, v475, int32(_a_F_synchronize_slots_1))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L1
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	if base.Ui64(v505) < base.Ui64(v473) {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	goto L137
L139:
	;
	v513 = v473
	goto L141
L140:
	;
	v513 = v505
	goto L141
L141:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v475)+104)) = v513
	v515 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v475))), uint32(v515))
	F_ReplicationSlotsComputeRequiredLSN(m)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	v521 = int64(*(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[6])))
	v522 = *(*int64)(unsafe.Add(mBase, uint32(v475)+104))
	v523 = F_XLogGetLastRemovedSegno(m)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	v525 = base.I64_div_u_s(v522, v521)
	if base.Ui64(v525) <= base.Ui64(v523) {
		goto L62
	} else {
		goto L144
	}
L144:
	;
	v528 = *(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[0]))
	F_LWLockRelease(m, v528+int32(_a_F_synchronize_slots_5))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	v534 = *(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[0]))
	v538 = F_LWLockAcquire(m, v534+int32(_a_F_synchronize_slots_0), int32(0))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	v541 = *(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[0]))
	v545 = F_LWLockAcquire(m, v541+int32(512), int32(0))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	v548 = F_GetOldestSafeDecodingTransactionId(m, int32(1))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	v552 = base.AtomicRmwXchg32(m, v439, int32(0), int32(1))
	if v552 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	F_s_lock(m, v439, int32(_a_F_synchronize_slots_1))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L1
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v439)+100)) = v548
	*(*int32)(unsafe.Add(mBase, uint32(v439)+20)) = v548
	v558 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v439))), uint32(v558))
	v561 = int32(1)
	F_ReplicationSlotsComputeRequiredXmin(m, v561)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L1
	} else {
		goto L153
	}
L152:
	;
	goto L151
L153:
	;
	v566 = *(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[0]))
	F_LWLockRelease(m, v566+int32(512))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	v572 = *(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[0]))
	F_LWLockRelease(m, v572+int32(_a_F_synchronize_slots_0))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	v577 = F_update_and_persist_local_synced_slot(m, v337, v340, l1)
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	v582 = v561
	goto L74
L158:
	;
	v593 = v582 | v324
	v595 = v325 + int32(1)
	v596 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v595 < v596 {
		v324 = v593
		v325 = v595
		goto L70
	} else {
		goto L159
	}
L159:
	;
	goto L71
L160:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v337)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v625
	F_errmsg(m, int32(_a_F_synchronize_slots_6), v17+int32(48))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	F_errfinish(m, int32(_a_F_synchronize_slots_3), int32(794), int32(_a_F_synchronize_slots_7))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L164:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v337)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v641
	F_errmsg_internal(m, int32(_a_F_synchronize_slots_8), v17+int32(32))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	v648 = *(*int64)(unsafe.Add(mBase, uint32(v347)+120))
	v649 = *(*int64)(unsafe.Add(mBase, uint32(v337)+24))
	*(*uint32)(unsafe.Add(mBase, uint32(v17)+28)) = uint32(v649)
	*(*uint32)(unsafe.Add(mBase, uint32(v17)+20)) = uint32(v648)
	v652 = int64(32)
	v653 = int64(base.Ui64(v649) >> (uint(v652) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v17)+24)) = uint32(v653)
	v656 = int64(base.Ui64(v648) >> (uint(v652) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v17)+16)) = uint32(v656)
	F_errdetail_internal(m, int32(_a_F_synchronize_slots_9), v17+int32(16))
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	F_errfinish(m, int32(_a_F_synchronize_slots_3), int32(867), int32(_a_F_synchronize_slots_7))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v475 + int32(24)
	F_errmsg_internal(m, int32(_a_F_synchronize_slots_10), v17)
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	F_errfinish(m, int32(_a_F_synchronize_slots_3), int32(660), int32(_a_F_synchronize_slots_11))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
