package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_GetWalRcvFlushRecPtr(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_GetWalRcvFlushRecPtr[0]))
	v8 = int32(1456)
	v9 = v7 + v8
	v12 = base.AtomicRmwXchg32(m, v7, v8, int32(1))
	if v12 != 0 {
		F_s_lock(m, v9, int32(_a_F_GetWalRcvFlushRecPtr_0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int64(0)
		} else {
			v18 = *(*int64)(unsafe.Add(mBase, uint32(v7)+48))
			if l0 != 0 {
				v19 = *(*int64)(unsafe.Add(mBase, uint32(v7)+64))
				*(*int64)(unsafe.Add(mBase, uint32(l0))) = v19
			} else {
			}
			if l1 != 0 {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v7)+56))
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v21
			} else {
			}
			v23 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v9))), uint32(v23))
			return v18
		}
	} else {
		v18 = *(*int64)(unsafe.Add(mBase, uint32(v7)+48))
		if l0 != 0 {
			v19 = *(*int64)(unsafe.Add(mBase, uint32(v7)+64))
			*(*int64)(unsafe.Add(mBase, uint32(l0))) = v19
		} else {
		}
		if l1 != 0 {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v7)+56))
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v21
		} else {
		}
		v23 = int32(0)
		atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v9))), uint32(v23))
		return v18
	}
}
func F_WALInsertLockRelease(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_WALInsertLockRelease[0]))
	v5 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WALInsertLockRelease[1])))
	if v5 != 0 {
		F_LWLockReleaseClearVar(m, v3, v3+int32(16))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, _c_F_WALInsertLockRelease[0]))
			F_LWLockReleaseClearVar(m, v11+int32(128), v11+int32(144))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, _c_F_WALInsertLockRelease[0]))
				F_LWLockReleaseClearVar(m, v19+int32(256), v19+int32(272))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, _c_F_WALInsertLockRelease[0]))
					F_LWLockReleaseClearVar(m, v27+int32(384), v27+int32(400))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, _c_F_WALInsertLockRelease[0]))
						F_LWLockReleaseClearVar(m, v35+int32(512), v35+int32(528))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							v43 = *(*int32)(unsafe.Add(mBase, _c_F_WALInsertLockRelease[0]))
							F_LWLockReleaseClearVar(m, v43+int32(640), v43+int32(656))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								v51 = *(*int32)(unsafe.Add(mBase, _c_F_WALInsertLockRelease[0]))
								F_LWLockReleaseClearVar(m, v51+int32(768), v51+int32(784))
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return
								} else {
									v59 = *(*int32)(unsafe.Add(mBase, _c_F_WALInsertLockRelease[0]))
									F_LWLockReleaseClearVar(m, v59+int32(896), v59+int32(912))
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return
									} else {
										v67 = int32(0)
										*(*uint8)(unsafe.Add(mBase, _c_F_WALInsertLockRelease[1])) = uint8(v67)
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
		v70 = *(*int32)(unsafe.Add(mBase, _c_F_WALInsertLockRelease[2]))
		v73 = v3 + v70<<(uint(int32(7))%32)
		F_LWLockReleaseClearVar(m, v73, v73+int32(16))
		mBase = m.M
		v77 = m.ExcPending
		if v77 != 0 {
			return
		} else {
			return
		}
	}
}
func F_WalRcvDie(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_WalRcvDie[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l1))))
	F_XLogWalRcvFlush(m, int32(1), v7)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v12 = base.AtomicRmwXchg32(m, v4, int32(1456), int32(1))
		if v12 != 0 {
			F_s_lock(m, v4+int32(1456), int32(_a_F_WalRcvDie_0))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				v18 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = v18
				*(*uint8)(unsafe.Add(mBase, uint32(v4)+1453)) = uint8(v18)
				*(*int64)(unsafe.Add(mBase, uint32(v4))) = int64(4294967295)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v4)+1456)), uint32(v18))
				F_ConditionVariableBroadcast(m, v4+int32(12))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, _c_F_WalRcvDie[1]))
					if v32 != 0 {
						v34 = *(*int32)(unsafe.Add(mBase, _c_F_WalRcvDie[2]))
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+64))
						m.T0[v35].(func(*base.Module, int32))(m, v32)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, _c_F_WalRcvDie[3]))
							F_SetLatch(m, v39+int32(4))
							mBase = m.M
							return
						}
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, _c_F_WalRcvDie[3]))
						F_SetLatch(m, v39+int32(4))
						mBase = m.M
						return
					}
				}
			}
		} else {
			v18 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = v18
			*(*uint8)(unsafe.Add(mBase, uint32(v4)+1453)) = uint8(v18)
			*(*int64)(unsafe.Add(mBase, uint32(v4))) = int64(4294967295)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v4)+1456)), uint32(v18))
			F_ConditionVariableBroadcast(m, v4+int32(12))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, _c_F_WalRcvDie[1]))
				if v32 != 0 {
					v34 = *(*int32)(unsafe.Add(mBase, _c_F_WalRcvDie[2]))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+64))
					m.T0[v35].(func(*base.Module, int32))(m, v32)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, _c_F_WalRcvDie[3]))
						F_SetLatch(m, v39+int32(4))
						mBase = m.M
						return
					}
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, _c_F_WalRcvDie[3]))
					F_SetLatch(m, v39+int32(4))
					mBase = m.M
					return
				}
			}
		}
	}
}
func F_WalRcvShmemInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_WalRcvShmemInit[0]))
	v4 = int32(0)
	base.MemoryFill(m, v3, v4, int32(1480))
	v8 = v3 + int32(12)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v8))), uint32(v4))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+4)) = int64(-1)
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_WalRcvShmemInit[0]))
	v16 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v15)+1456)), uint32(v16))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+1464)) = int64(0)
	return
}
func F_WalRcvStreaming(m *base.Module) int32 {
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
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_WalRcvStreaming[0]))
	v7 = int32(1456)
	v8 = v6 + v7
	v11 = base.AtomicRmwXchg32(m, v6, v7, int32(1))
	if v11 != 0 {
		F_s_lock(m, v8, int32(_a_F_WalRcvStreaming_0))
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
				return base.B2i32(v51 == int32(2)) | base.B2i32(v51&int32(-3) == int32(1)) | base.B2i32(v51 == int32(5))
			} else {
				v24 = int32(1)
				v25 = F_time(m)
				mBase = m.M
				if v25-v17 < int64(11) {
					v51 = v24
					return base.B2i32(v51 == int32(2)) | base.B2i32(v51&int32(-3) == int32(1)) | base.B2i32(v51 == int32(5))
				} else {
					v31 = base.AtomicRmwXchg32(m, v8, int32(0), int32(1))
					if v31 != 0 {
						F_s_lock(m, v8, int32(_a_F_WalRcvStreaming_0))
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
								return base.B2i32(v51 == int32(2)) | base.B2i32(v51&int32(-3) == int32(1)) | base.B2i32(v51 == int32(5))
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
									return base.B2i32(v51 == int32(2)) | base.B2i32(v51&int32(-3) == int32(1)) | base.B2i32(v51 == int32(5))
								}
							}
						}
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
						if v35 != int32(1) {
							v38 = int32(0)
							atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v8))), uint32(v38))
							v51 = v24
							return base.B2i32(v51 == int32(2)) | base.B2i32(v51&int32(-3) == int32(1)) | base.B2i32(v51 == int32(5))
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
								return base.B2i32(v51 == int32(2)) | base.B2i32(v51&int32(-3) == int32(1)) | base.B2i32(v51 == int32(5))
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
			return base.B2i32(v51 == int32(2)) | base.B2i32(v51&int32(-3) == int32(1)) | base.B2i32(v51 == int32(5))
		} else {
			v24 = int32(1)
			v25 = F_time(m)
			mBase = m.M
			if v25-v17 < int64(11) {
				v51 = v24
				return base.B2i32(v51 == int32(2)) | base.B2i32(v51&int32(-3) == int32(1)) | base.B2i32(v51 == int32(5))
			} else {
				v31 = base.AtomicRmwXchg32(m, v8, int32(0), int32(1))
				if v31 != 0 {
					F_s_lock(m, v8, int32(_a_F_WalRcvStreaming_0))
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
							return base.B2i32(v51 == int32(2)) | base.B2i32(v51&int32(-3) == int32(1)) | base.B2i32(v51 == int32(5))
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
								return base.B2i32(v51 == int32(2)) | base.B2i32(v51&int32(-3) == int32(1)) | base.B2i32(v51 == int32(5))
							}
						}
					}
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
					if v35 != int32(1) {
						v38 = int32(0)
						atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v8))), uint32(v38))
						v51 = v24
						return base.B2i32(v51 == int32(2)) | base.B2i32(v51&int32(-3) == int32(1)) | base.B2i32(v51 == int32(5))
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
							return base.B2i32(v51 == int32(2)) | base.B2i32(v51&int32(-3) == int32(1)) | base.B2i32(v51 == int32(5))
						}
					}
				}
			}
		}
	}
}
func F_WalSndKeepalive(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v42 int64
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v57 int64
	_ = v57
	var v59 int64
	_ = v59
	var v62 int64
	_ = v62
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v68 int64
	_ = v68
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int64
	_ = v103
	var v104 int64
	_ = v104
	var v112 int64
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int64
	_ = v122
	var v124 int64
	_ = v124
	var v126 int64
	_ = v126
	var v129 int64
	_ = v129
	var v131 int64
	_ = v131
	var v133 int64
	_ = v133
	var v135 int64
	_ = v135
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
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	v1 = l0
	v7 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		if v7 != 0 {
			F_errmsg_internal(m, int32(_a_F_WalSndKeepalive_0), int32(0))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_WalSndKeepalive_1), int32(_a_F_WalSndKeepalive_2), int32(_a_F_WalSndKeepalive_3))
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					v18 = int32(_a_F_WalSndKeepalive_4)
					v19 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[0]))
					v20 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v20)
					*(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[1])) = v20
					*(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[2])) = v20
					F_enlargeStringInfo(m, int32(_a_F_WalSndKeepalive_4), int32(1))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						v30 = int32(_a_F_WalSndKeepalive_5)
						v31 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[2]))
						v32 = int32(_a_F_WalSndKeepalive_4)
						v33 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[0]))
						v35 = int32(107)
						*(*uint8)(unsafe.Add(mBase, uint32(v31+v33))) = uint8(v35)
						*(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[2])) = v31 + int32(1)
						v42 = *(*int64)(unsafe.Add(mBase, _c_F_WalSndKeepalive[3]))
						F_enlargeStringInfo(m, v32, int32(8))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							v48 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[2]))
							v50 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[0]))
							if l1 == int64(0) {
								v54 = v42
							} else {
								v54 = l1
							}
							v55 = int64(56)
							v57 = int64(65280)
							v59 = int64(40)
							v62 = int64(16711680)
							v64 = int64(24)
							v66 = int64(4278190080)
							v68 = int64(8)
							*(*int64)(unsafe.Add(mBase, uint32(v48+v50))) = v54<<(uint(v55)%64) | v54&v57<<(uint(v59)%64) | (v54&v62<<(uint(v64)%64) | v54&v66<<(uint(v68)%64)) | (int64(base.Ui64(v54)>>(uint(v68)%64))&v66 | int64(base.Ui64(v54)>>(uint(v64)%64))&v62 | (int64(base.Ui64(v54)>>(uint(v59)%64))&v57 | int64(base.Ui64(v54)>>(uint(v55)%64))))
							*(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[2])) = v48 + int32(8)
							v98 = m.G0
							v99 = int32(16)
							v100 = v98 - v99
							m.G0 = v100
							F_gettimeofday(m, v100)
							mBase = m.M
							v103 = *(*int64)(unsafe.Add(mBase, uint32(v100)))
							v104 = int64(*(*int32)(unsafe.Add(mBase, uint32(v100)+8)))
							m.G0 = v100 + v99
							v112 = v104 + v103*int64(1000000) - int64(946684800000000)
							F_enlargeStringInfo(m, int32(_a_F_WalSndKeepalive_4), int32(8))
							mBase = m.M
							v116 = m.ExcPending
							if v116 != 0 {
								return
							} else {
								v117 = int32(_a_F_WalSndKeepalive_5)
								v118 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[2]))
								v119 = int32(_a_F_WalSndKeepalive_4)
								v120 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[0]))
								v122 = int64(56)
								v124 = int64(65280)
								v126 = int64(40)
								v129 = int64(16711680)
								v131 = int64(24)
								v133 = int64(4278190080)
								v135 = int64(8)
								*(*int64)(unsafe.Add(mBase, uint32(v118+v120))) = v112<<(uint(v122)%64) | v112&v124<<(uint(v126)%64) | (v112&v129<<(uint(v131)%64) | v112&v133<<(uint(v135)%64)) | (int64(base.Ui64(v112)>>(uint(v135)%64))&v133 | int64(base.Ui64(v112)>>(uint(v131)%64))&v129 | (int64(base.Ui64(v112)>>(uint(v126)%64))&v124 | int64(base.Ui64(v112)>>(uint(v122)%64))))
								*(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[2])) = v118 + int32(8)
								F_enlargeStringInfo(m, v119, int32(1))
								mBase = m.M
								v165 = m.ExcPending
								if v165 != 0 {
									return
								} else {
									v166 = int32(_a_F_WalSndKeepalive_5)
									v167 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[2]))
									v168 = int32(_a_F_WalSndKeepalive_4)
									v169 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[0]))
									*(*uint8)(unsafe.Add(mBase, uint32(v167+v169))) = uint8(v1)
									v174 = v167 + int32(1)
									*(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[2])) = v174
									v178 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[0]))
									v180 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[4]))
									v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+20))
									m.T0[v181].(func(*base.Module, int32, int32, int32))(m, int32(100), v178, v174)
									mBase = m.M
									v183 = m.ExcPending
									if v183 != 0 {
										return
									} else {
										if v1 != 0 {
											v185 = int32(1)
											*(*uint8)(unsafe.Add(mBase, _c_F_WalSndKeepalive[5])) = uint8(v185)
										} else {
										}
										return
									}
								}
							}
						}
					}
				}
			}
		} else {
			v18 = int32(_a_F_WalSndKeepalive_4)
			v19 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[0]))
			v20 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v20)
			*(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[1])) = v20
			*(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[2])) = v20
			F_enlargeStringInfo(m, int32(_a_F_WalSndKeepalive_4), int32(1))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				v30 = int32(_a_F_WalSndKeepalive_5)
				v31 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[2]))
				v32 = int32(_a_F_WalSndKeepalive_4)
				v33 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[0]))
				v35 = int32(107)
				*(*uint8)(unsafe.Add(mBase, uint32(v31+v33))) = uint8(v35)
				*(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[2])) = v31 + int32(1)
				v42 = *(*int64)(unsafe.Add(mBase, _c_F_WalSndKeepalive[3]))
				F_enlargeStringInfo(m, v32, int32(8))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[2]))
					v50 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[0]))
					if l1 == int64(0) {
						v54 = v42
					} else {
						v54 = l1
					}
					v55 = int64(56)
					v57 = int64(65280)
					v59 = int64(40)
					v62 = int64(16711680)
					v64 = int64(24)
					v66 = int64(4278190080)
					v68 = int64(8)
					*(*int64)(unsafe.Add(mBase, uint32(v48+v50))) = v54<<(uint(v55)%64) | v54&v57<<(uint(v59)%64) | (v54&v62<<(uint(v64)%64) | v54&v66<<(uint(v68)%64)) | (int64(base.Ui64(v54)>>(uint(v68)%64))&v66 | int64(base.Ui64(v54)>>(uint(v64)%64))&v62 | (int64(base.Ui64(v54)>>(uint(v59)%64))&v57 | int64(base.Ui64(v54)>>(uint(v55)%64))))
					*(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[2])) = v48 + int32(8)
					v98 = m.G0
					v99 = int32(16)
					v100 = v98 - v99
					m.G0 = v100
					F_gettimeofday(m, v100)
					mBase = m.M
					v103 = *(*int64)(unsafe.Add(mBase, uint32(v100)))
					v104 = int64(*(*int32)(unsafe.Add(mBase, uint32(v100)+8)))
					m.G0 = v100 + v99
					v112 = v104 + v103*int64(1000000) - int64(946684800000000)
					F_enlargeStringInfo(m, int32(_a_F_WalSndKeepalive_4), int32(8))
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return
					} else {
						v117 = int32(_a_F_WalSndKeepalive_5)
						v118 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[2]))
						v119 = int32(_a_F_WalSndKeepalive_4)
						v120 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[0]))
						v122 = int64(56)
						v124 = int64(65280)
						v126 = int64(40)
						v129 = int64(16711680)
						v131 = int64(24)
						v133 = int64(4278190080)
						v135 = int64(8)
						*(*int64)(unsafe.Add(mBase, uint32(v118+v120))) = v112<<(uint(v122)%64) | v112&v124<<(uint(v126)%64) | (v112&v129<<(uint(v131)%64) | v112&v133<<(uint(v135)%64)) | (int64(base.Ui64(v112)>>(uint(v135)%64))&v133 | int64(base.Ui64(v112)>>(uint(v131)%64))&v129 | (int64(base.Ui64(v112)>>(uint(v126)%64))&v124 | int64(base.Ui64(v112)>>(uint(v122)%64))))
						*(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[2])) = v118 + int32(8)
						F_enlargeStringInfo(m, v119, int32(1))
						mBase = m.M
						v165 = m.ExcPending
						if v165 != 0 {
							return
						} else {
							v166 = int32(_a_F_WalSndKeepalive_5)
							v167 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[2]))
							v168 = int32(_a_F_WalSndKeepalive_4)
							v169 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[0]))
							*(*uint8)(unsafe.Add(mBase, uint32(v167+v169))) = uint8(v1)
							v174 = v167 + int32(1)
							*(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[2])) = v174
							v178 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[0]))
							v180 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[4]))
							v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+20))
							m.T0[v181].(func(*base.Module, int32, int32, int32))(m, int32(100), v178, v174)
							mBase = m.M
							v183 = m.ExcPending
							if v183 != 0 {
								return
							} else {
								if v1 != 0 {
									v185 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _c_F_WalSndKeepalive[5])) = uint8(v185)
								} else {
								}
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_WalSndLoop(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int64
	_ = v23
	var v24 int64
	_ = v24
	var v35 int32
	_ = v35
	var v46 int64
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
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
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int64
	_ = v168
	var v171 int64
	_ = v171
	var v172 int64
	_ = v172
	var v174 int32
	_ = v174
	var v178 int64
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int64
	_ = v225
	var v226 int64
	_ = v226
	var v234 int64
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v241 int64
	_ = v241
	var v245 int32
	_ = v245
	var v254 int64
	_ = v254
	var v260 int64
	_ = v260
	var v269 int64
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int64
	_ = v276
	var v280 int32
	_ = v280
	var v286 int64
	_ = v286
	var v292 int64
	_ = v292
	var v301 int64
	_ = v301
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v344 int64
	_ = v344
	var v348 int32
	_ = v348
	var v352 int64
	_ = v352
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v383 int64
	_ = v383
	var v387 int32
	_ = v387
	var v389 int64
	_ = v389
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v414 int32
	_ = v414
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v436 int64
	_ = v436
	var v437 int64
	_ = v437
	var v445 int64
	_ = v445
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v452 int64
	_ = v452
	var v456 int32
	_ = v456
	var v465 int64
	_ = v465
	var v471 int64
	_ = v471
	var v480 int64
	_ = v480
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v490 int64
	_ = v490
	var v494 int32
	_ = v494
	var v500 int64
	_ = v500
	var v506 int64
	_ = v506
	var v515 int64
	_ = v515
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int64
	_ = v544
	var v547 int32
	_ = v547
	var v561 int32
	_ = v561
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v18 = m.G0
	v19 = int32(16)
	v20 = v18 - v19
	m.G0 = v20
	F_gettimeofday(m, v20)
	mBase = m.M
	v23 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	v24 = int64(*(*int32)(unsafe.Add(mBase, uint32(v20)+8)))
	m.G0 = v20 + v19
	goto L1
L1:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_WalSndLoop[0])) = v24 + v23*int64(1000000) - int64(946684800000000)
	v35 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_WalSndLoop[1])) = uint8(v35)
	v46 = int64(0)
	goto L3
L2:
	;
	F_WalSndShutdown(m)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L9
	} else {
		goto L152
	}
L3:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[2]))
	v50 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = v50
	v55 = base.AtomicRmwOr32(m, v50, int32(_a_F_WalSndLoop_0), v50)
	goto L5
L4:
	;
	m.G0 = v12 + int32(32)
	return
L5:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[3]))
	if v57 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[4]))
	if v61 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	return
L10:
	;
	goto L8
L11:
	;
	F_ProcessRepliesIfAny(m)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L9
	} else {
		goto L17
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[4])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	F_SyncRepInitConfig(m)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalSndLoop[5])))
	if v73 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	F_SyncRepReleaseWaiters(m)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	goto L11
L17:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalSndLoop[6])))
	if v79 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L4
L19:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[7]))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+12))
	v98 = m.T0[v97].(func(*base.Module) int32)(m)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L9
	} else {
		goto L25
	}
L20:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalSndLoop[8])))
	if v83&int32(1) == int32(0) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[7]))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v91 = m.T0[v90].(func(*base.Module) int32)(m)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L9
	} else {
		goto L22
	}
L22:
	;
	if v91 == int32(0) {
		goto L18
	} else {
		goto L23
	}
L23:
	;
	goto L19
L24:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[7]))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+8))
	v110 = m.T0[v109].(func(*base.Module) int32)(m)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L9
	} else {
		goto L30
	}
L25:
	;
	if v98 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	m.T0[l0].(func(*base.Module))(m)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L9
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v105 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_WalSndLoop[9])) = uint8(v105)
	goto L24
L29:
	;
	goto L24
L30:
	;
	if v110 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalSndLoop[9])))
	if v113 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v344 = *(*int64)(unsafe.Add(mBase, _c_F_WalSndLoop[0]))
	if v344 <= int64(0) {
		goto L95
	} else {
		goto L96
	}
L33:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[7]))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+12))
	v119 = m.T0[v118].(func(*base.Module) int32)(m)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L9
	} else {
		goto L34
	}
L34:
	;
	if v119 != 0 {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[10]))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	if v123 != int32(2) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v161 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[11]))
	if v161 == int32(0) {
		goto L32
	} else {
		goto L49
	}
L37:
	;
	v128 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L9
	} else {
		goto L38
	}
L38:
	;
	if v128 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[12]))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v131
	F_errmsg_internal(m, int32(_a_F_WalSndLoop_1), v12)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L9
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v142 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[10]))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	if v143 == int32(3) {
		goto L36
	} else {
		goto L44
	}
L42:
	;
	F_errfinish(m, int32(_a_F_WalSndLoop_2), int32(3122), int32(_a_F_WalSndLoop_3))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L9
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v148 = base.AtomicRmwXchg32(m, v142, int32(76), int32(1))
	if v148 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	F_s_lock(m, v142+int32(76), int32(_a_F_WalSndLoop_4))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L9
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v142)+4)) = int32(3)
	v156 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v142)+76)), uint32(v156))
	goto L36
L48:
	;
	goto L47
L49:
	;
	m.T0[l0].(func(*base.Module))(m)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L9
	} else {
		goto L50
	}
L50:
	;
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[10]))
	v168 = *(*int64)(unsafe.Add(mBase, uint32(v167)+32))
	if v168 == int64(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v171 = *(*int64)(unsafe.Add(mBase, uint32(v167)+24))
	v172 = v171
	goto L53
L52:
	;
	v172 = v168
	goto L53
L53:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalSndLoop[9])))
	if v174 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v336 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalSndLoop[1])))
	if v336 != 0 {
		goto L32
	} else {
		goto L93
	}
L55:
	;
	v178 = *(*int64)(unsafe.Add(mBase, _c_F_WalSndLoop[13]))
	if v178 != v172 {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v181 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[7]))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
	v183 = m.T0[v182].(func(*base.Module) int32)(m)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L9
	} else {
		goto L57
	}
L57:
	;
	if v183 != 0 {
		goto L54
	} else {
		goto L58
	}
L58:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(56)
	F_EndCommandExtended(m, v12+int32(16))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L9
	} else {
		goto L59
	}
L59:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_WalSndLoop[0])) = int64(0)
	v197 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_WalSndLoop[14])) = uint8(v197)
	goto L60
L60:
	;
	F_WalSndCheckShutdownTimeout(m)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L9
	} else {
		goto L63
	}
L61:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L9
	} else {
		goto L92
	}
L62:
	;
	goto L61
L63:
	;
	v211 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[7]))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)+12))
	v213 = m.T0[v212].(func(*base.Module) int32)(m)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L9
	} else {
		goto L64
	}
L64:
	;
	if v213 == int32(0) {
		goto L62
	} else {
		goto L65
	}
L65:
	;
	v220 = m.G0
	v221 = int32(16)
	v222 = v220 - v221
	m.G0 = v222
	F_gettimeofday(m, v222)
	mBase = m.M
	v225 = *(*int64)(unsafe.Add(mBase, uint32(v222)))
	v226 = int64(*(*int32)(unsafe.Add(mBase, uint32(v222)+8)))
	m.G0 = v222 + v221
	v234 = v226 + v225*int64(1000000) - int64(946684800000000)
	goto L66
L66:
	;
	v235 = int32(_a_F_WalSndLoop_5)
	v237 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[15]))
	if v237 <= int32(0) {
		v273 = v235
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v276 = *(*int64)(unsafe.Add(mBase, _c_F_WalSndLoop[16]))
	if v276 == int64(0) {
		v308 = v273
		goto L74
	} else {
		goto L75
	}
L68:
	;
	v241 = *(*int64)(unsafe.Add(mBase, _c_F_WalSndLoop[0]))
	if v241 <= int64(0) {
		v273 = v235
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalSndLoop[1])))
	v254 = base.I64_extend_i32_u(int32(base.Ui32(v237)>>(uint((v245^int32(-1))&int32(1))%32)))*int64(1000) + v241
	if v254 <= v234 {
		v272 = int32(0)
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v273 = v272
	goto L67
L71:
	;
	goto L70
L72:
	;
	v260 = v254 - v234
	if base.B2i32(int64(0) < v234)^base.B2i32(v260 < v254)|base.B2i32(int64(2147483646000) < v260) != 0 {
		v272 = int32(2147483647)
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v269 = base.I64_div_s(v260+int64(999), int64(1000))
	v272 = base.I32_wrap_i64(v269)
	goto L71
L74:
	;
	F_WalSndWait(m, int32(4), v308, int32(100663304))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L9
	} else {
		goto L84
	}
L75:
	;
	v280 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[17]))
	if v280 <= int32(0) {
		v308 = v273
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v286 = base.I64_extend_i32_u(v280)*int64(1000) + v276
	if v286 <= v234 {
		v304 = int32(0)
		goto L78
	} else {
		goto L79
	}
L77:
	;
	if v304 < v273 {
		goto L81
	} else {
		goto L82
	}
L78:
	;
	goto L77
L79:
	;
	v292 = v286 - v234
	if base.B2i32(int64(0) < v234)^base.B2i32(v292 < v286)|base.B2i32(int64(2147483646000) < v292) != 0 {
		v304 = int32(2147483647)
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v301 = base.I64_div_s(v292+int64(999), int64(1000))
	v304 = base.I32_wrap_i64(v301)
	goto L78
L81:
	;
	v306 = v304
	goto L83
L82:
	;
	v306 = v273
	goto L83
L83:
	;
	v308 = v306
	goto L74
L84:
	;
	v314 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[2]))
	v315 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v314))) = v315
	v320 = base.AtomicRmwOr32(m, v315, int32(_a_F_WalSndLoop_0), v315)
	goto L85
L85:
	;
	v322 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[3]))
	if v322 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L9
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v326 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[7]))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v326)+8))
	v328 = m.T0[v327].(func(*base.Module) int32)(m)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L9
	} else {
		goto L90
	}
L89:
	;
	goto L88
L90:
	;
	if v328 == int32(0) {
		goto L60
	} else {
		goto L91
	}
L91:
	;
	goto L2
L92:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L93:
	;
	F_WalSndKeepalive(m, int32(1), int64(0))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L9
	} else {
		goto L94
	}
L94:
	;
	goto L32
L95:
	;
	F_WalSndCheckShutdownTimeout(m)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L9
	} else {
		goto L104
	}
L96:
	;
	v348 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[15]))
	if v348 <= int32(0) {
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v352 = *(*int64)(unsafe.Add(mBase, _c_F_WalSndLoop[18]))
	if v352 < base.I64_extend_i32_u(v348)*int64(1000)+v344 {
		goto L95
	} else {
		goto L98
	}
L98:
	;
	v360 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L9
	} else {
		goto L99
	}
L99:
	;
	if v360 == int32(0) {
		goto L2
	} else {
		goto L100
	}
L100:
	;
	F_errmsg(m, int32(_a_F_WalSndLoop_6), int32(0))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L9
	} else {
		goto L101
	}
L101:
	;
	F_errfinish(m, int32(_a_F_WalSndLoop_2), int32(3008), int32(_a_F_WalSndLoop_7))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L9
	} else {
		goto L102
	}
L102:
	;
	F_WalSndShutdown(m)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L9
	} else {
		goto L103
	}
L103:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L104:
	;
	v379 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[15]))
	if v379 <= int32(0) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	if l0 == int32(1112) {
		goto L114
	} else {
		goto L115
	}
L106:
	;
	v383 = *(*int64)(unsafe.Add(mBase, _c_F_WalSndLoop[0]))
	if v383 <= int64(0) {
		goto L105
	} else {
		goto L107
	}
L107:
	;
	v387 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalSndLoop[1])))
	if v387 != 0 {
		goto L105
	} else {
		goto L108
	}
L108:
	;
	v389 = *(*int64)(unsafe.Add(mBase, _c_F_WalSndLoop[18]))
	if v389 < base.I64_extend_i32_u(int32(base.Ui32(v379)>>(uint(int32(1))%32)))*int64(1000)+v383 {
		goto L105
	} else {
		goto L109
	}
L109:
	;
	F_WalSndKeepalive(m, int32(1), int64(0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L9
	} else {
		goto L110
	}
L110:
	;
	v402 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[7]))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v402)+8))
	v404 = m.T0[v403].(func(*base.Module) int32)(m)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L9
	} else {
		goto L111
	}
L111:
	;
	if v404 != 0 {
		goto L2
	} else {
		goto L112
	}
L112:
	;
	goto L105
L113:
	;
	v427 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalSndLoop[6])))
	v431 = m.G0
	v432 = int32(16)
	v433 = v431 - v432
	m.G0 = v433
	F_gettimeofday(m, v433)
	mBase = m.M
	v436 = *(*int64)(unsafe.Add(mBase, uint32(v433)))
	v437 = int64(*(*int32)(unsafe.Add(mBase, uint32(v433)+8)))
	m.G0 = v433 + v432
	v445 = v437 + v436*int64(1000000) - int64(946684800000000)
	goto L120
L114:
	;
	v420 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[7]))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v420)+12))
	v422 = m.T0[v421].(func(*base.Module) int32)(m)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L9
	} else {
		goto L118
	}
L115:
	;
	v408 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalSndLoop[9])))
	if v408&int32(1) == int32(0) {
		goto L114
	} else {
		goto L116
	}
L116:
	;
	v414 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalSndLoop[8])))
	if v414&int32(1) == int32(0) {
		goto L113
	} else {
		goto L117
	}
L117:
	;
	goto L114
L118:
	;
	if v422 == int32(0) {
		goto L3
	} else {
		goto L119
	}
L119:
	;
	goto L113
L120:
	;
	v446 = int32(_a_F_WalSndLoop_5)
	v448 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[15]))
	if v448 <= int32(0) {
		v484 = v446
		goto L121
	} else {
		goto L122
	}
L121:
	;
	if v427 != 0 {
		goto L128
	} else {
		goto L129
	}
L122:
	;
	v452 = *(*int64)(unsafe.Add(mBase, _c_F_WalSndLoop[0]))
	if v452 <= int64(0) {
		v484 = v446
		goto L121
	} else {
		goto L123
	}
L123:
	;
	v456 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalSndLoop[1])))
	v465 = base.I64_extend_i32_u(int32(base.Ui32(v448)>>(uint((v456^int32(-1))&int32(1))%32)))*int64(1000) + v452
	if v465 <= v445 {
		v483 = int32(0)
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v484 = v483
	goto L121
L125:
	;
	goto L124
L126:
	;
	v471 = v465 - v445
	if base.B2i32(int64(0) < v445)^base.B2i32(v471 < v465)|base.B2i32(int64(2147483646000) < v471) != 0 {
		v483 = int32(2147483647)
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v480 = base.I64_div_s(v471+int64(999), int64(1000))
	v483 = base.I32_wrap_i64(v480)
	goto L125
L128:
	;
	v488 = int32(0)
	goto L130
L129:
	;
	v488 = int32(2)
	goto L130
L130:
	;
	v490 = *(*int64)(unsafe.Add(mBase, _c_F_WalSndLoop[16]))
	if v490 == int64(0) {
		v521 = v484
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v526 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[7]))
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v526)+12))
	v528 = m.T0[v527].(func(*base.Module) int32)(m)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L9
	} else {
		goto L141
	}
L132:
	;
	v494 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[17]))
	if v494 <= int32(0) {
		v521 = v484
		goto L131
	} else {
		goto L133
	}
L133:
	;
	v500 = base.I64_extend_i32_u(v494)*int64(1000) + v490
	if v500 <= v445 {
		v518 = int32(0)
		goto L135
	} else {
		goto L136
	}
L134:
	;
	if v518 < v484 {
		goto L138
	} else {
		goto L139
	}
L135:
	;
	goto L134
L136:
	;
	v506 = v500 - v445
	if base.B2i32(int64(0) < v445)^base.B2i32(v506 < v500)|base.B2i32(int64(2147483646000) < v506) != 0 {
		v518 = int32(2147483647)
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v515 = base.I64_div_s(v506+int64(999), int64(1000))
	v518 = base.I32_wrap_i64(v515)
	goto L135
L138:
	;
	v520 = v518
	goto L140
L139:
	;
	v520 = v484
	goto L140
L140:
	;
	v521 = v520
	goto L131
L141:
	;
	if v528 != 0 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v530 = v488 | int32(4)
	goto L144
L143:
	;
	v530 = v488
	goto L144
L144:
	;
	goto L145
L145:
	;
	if base.I64_extend_i32_s(int32(1000))*int64(1000) <= v445-v46 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	F_pgstat_flush_io(m, int32(0))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L9
	} else {
		goto L149
	}
L147:
	;
	v544 = v46
	goto L148
L148:
	;
	F_WalSndWait(m, v530, v521, int32(83886095))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L9
	} else {
		goto L151
	}
L149:
	;
	v542 = F_pgstat_flush_backend(m, int32(0), int32(1))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L9
	} else {
		goto L150
	}
L150:
	;
	v544 = v445
	goto L148
L151:
	;
	v46 = v544
	goto L3
L152:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_WalSndShutdown(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndShutdown[0]))
	if v2 == int32(2) {
		*(*int32)(unsafe.Add(mBase, _c_F_WalSndShutdown[0])) = int32(0)
	} else {
	}
	F_proc_exit(m, int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		base.Wasm_trap_unreachable()
		for {
		}
	}
}
func F_WalSndUpdateProgress(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	var v32 int64
	_ = v32
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v64 int64
	_ = v64
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int64
	_ = v74
	var v76 int64
	_ = v76
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int64
	_ = v86
	var v88 int64
	_ = v88
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int64
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	v13 = m.G0
	v14 = int32(16)
	v15 = v13 - v14
	m.G0 = v15
	F_gettimeofday(m, v15)
	mBase = m.M
	v18 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
	v19 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15)+8)))
	m.G0 = v15 + v14
	v27 = v19 + v18*int64(1000000) - int64(946684800000000)
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+164)))
	if v28 != int32(1) {
	} else {
		v32 = *(*int64)(unsafe.Add(mBase, _c_F_WalSndUpdateProgress[0]))
		if base.B2i32(base.I64_extend_i32_s(int32(1000))*int64(1000) <= v27-v32) == int32(0) {
		} else {
			v42 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalSndUpdateProgress[1])))
			if v42 != int32(1) {
			} else {
				v46 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndUpdateProgress[2]))
				v47 = *(*int64)(unsafe.Add(mBase, uint32(v46)))
				if v47 == l1 {
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v46))) = l1
					v51 = v46 + int32(8)
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_WalSndUpdateProgress[3])))
					v56 = base.I32_rem_s(v52+int32(1), int32(_a_F_WalSndUpdateProgress_0))
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_WalSndUpdateProgress[4])))
					if v56 == v57 {
						v61 = v51 + v56<<(uint(int32(4))%32)
						v62 = *(*int64)(unsafe.Add(mBase, uint32(v61)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_WalSndUpdateProgress[5]))) = v62
						v64 = *(*int64)(unsafe.Add(mBase, uint32(v61)))
						*(*int64)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_WalSndUpdateProgress[6]))) = v64
						*(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_WalSndUpdateProgress[4]))) = int32(-1)
					} else {
					}
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_WalSndUpdateProgress[7])))
					if v69 == v56 {
						v73 = v51 + v56<<(uint(int32(4))%32)
						v74 = *(*int64)(unsafe.Add(mBase, uint32(v73)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_WalSndUpdateProgress[8]))) = v74
						v76 = *(*int64)(unsafe.Add(mBase, uint32(v73)))
						*(*int64)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_WalSndUpdateProgress[9]))) = v76
						*(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_WalSndUpdateProgress[7]))) = int32(-1)
					} else {
					}
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_WalSndUpdateProgress[10])))
					if v81 == v56 {
						v85 = v51 + v56<<(uint(int32(4))%32)
						v86 = *(*int64)(unsafe.Add(mBase, uint32(v85)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_WalSndUpdateProgress[11]))) = v86
						v88 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
						*(*int64)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_WalSndUpdateProgress[12]))) = v88
						*(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_WalSndUpdateProgress[10]))) = int32(-1)
					} else {
					}
					v93 = int32(4)
					*(*int64)(unsafe.Add(mBase, uint32(v51+v52<<(uint(v93)%32)))) = l1
					v97 = *(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_WalSndUpdateProgress[3])))
					*(*int64)(unsafe.Add(mBase, uint32(v46+v97<<(uint(v93)%32))+16)) = v27
					*(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_WalSndUpdateProgress[3]))) = v56
				}
			}
			*(*int64)(unsafe.Add(mBase, _c_F_WalSndUpdateProgress[0])) = v27
		}
	}
	if l3 == int32(0) {
		if v28 != 0 {
			return
		} else {
			v146 = *(*int64)(unsafe.Add(mBase, _c_F_WalSndUpdateProgress[13]))
			v148 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndUpdateProgress[14]))
			v150 = base.I32_div_s(v148, int32(2))
			if v27 < v146+base.I64_extend_i32_s(v150)*int64(1000) {
				return
			} else {
				F_ProcessPendingWrites(m)
				mBase = m.M
				v157 = m.ExcPending
				if v157 != 0 {
					return
				} else {
					return
				}
			}
		}
	} else {
		v118 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndUpdateProgress[15]))
		if v118 <= int32(0) {
			if v28 != 0 {
				return
			} else {
				v146 = *(*int64)(unsafe.Add(mBase, _c_F_WalSndUpdateProgress[13]))
				v148 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndUpdateProgress[14]))
				v150 = base.I32_div_s(v148, int32(2))
				if v27 < v146+base.I64_extend_i32_s(v150)*int64(1000) {
					return
				} else {
					F_ProcessPendingWrites(m)
					mBase = m.M
					v157 = m.ExcPending
					if v157 != 0 {
						return
					} else {
						return
					}
				}
			}
		} else {
			v122 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndUpdateProgress[16]))
			if v122 < int32(2) {
				if v28 != 0 {
					return
				} else {
					v146 = *(*int64)(unsafe.Add(mBase, _c_F_WalSndUpdateProgress[13]))
					v148 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndUpdateProgress[14]))
					v150 = base.I32_div_s(v148, int32(2))
					if v27 < v146+base.I64_extend_i32_s(v150)*int64(1000) {
						return
					} else {
						F_ProcessPendingWrites(m)
						mBase = m.M
						v157 = m.ExcPending
						if v157 != 0 {
							return
						} else {
							return
						}
					}
				}
			} else {
				v126 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndUpdateProgress[17]))
				v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+48)))
				if v127&int32(2) == int32(0) {
					if v28 != 0 {
						return
					} else {
						v146 = *(*int64)(unsafe.Add(mBase, _c_F_WalSndUpdateProgress[13]))
						v148 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndUpdateProgress[14]))
						v150 = base.I32_div_s(v148, int32(2))
						if v27 < v146+base.I64_extend_i32_s(v150)*int64(1000) {
							return
						} else {
							F_ProcessPendingWrites(m)
							mBase = m.M
							v157 = m.ExcPending
							if v157 != 0 {
								return
							} else {
								return
							}
						}
					}
				} else {
					F_WalSndKeepalive(m, int32(0), l1)
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						v136 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndUpdateProgress[18]))
						v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)+8))
						v138 = m.T0[v137].(func(*base.Module) int32)(m)
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							if v138 != 0 {
								F_WalSndShutdown(m)
								mBase = m.M
								v159 = m.ExcPending
								if v159 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v141 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndUpdateProgress[18]))
								v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+12))
								v143 = m.T0[v142].(func(*base.Module) int32)(m)
								mBase = m.M
								v144 = m.ExcPending
								if v144 != 0 {
									return
								} else {
									if v143 != 0 {
										F_ProcessPendingWrites(m)
										mBase = m.M
										v157 = m.ExcPending
										if v157 != 0 {
											return
										} else {
											return
										}
									} else {
										if v28 != 0 {
											return
										} else {
											v146 = *(*int64)(unsafe.Add(mBase, _c_F_WalSndUpdateProgress[13]))
											v148 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndUpdateProgress[14]))
											v150 = base.I32_div_s(v148, int32(2))
											if v27 < v146+base.I64_extend_i32_s(v150)*int64(1000) {
												return
											} else {
												F_ProcessPendingWrites(m)
												mBase = m.M
												v157 = m.ExcPending
												if v157 != 0 {
													return
												} else {
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
func F_wal_segment_close(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1168))
	v3 = F_close(m, v2)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+1168)) = int32(-1)
	return
}
