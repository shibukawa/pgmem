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
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_GetWalRcvFlushRecPtr[0]))
	v8 = int32(1456)
	v9 = v7 + v8
	v12 = base.AtomicRmwXchg32(m, v7, v8, int32(1))
	if v12 != 0 {
		F_s_lock(m, v9, int32(_a_F_GetWalRcvFlushRecPtr_0), int32(337), int32(_a_F_GetWalRcvFlushRecPtr_1))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			v20 = *(*int64)(unsafe.Add(mBase, uint32(v7)+48))
			if l0 != 0 {
				v21 = *(*int64)(unsafe.Add(mBase, uint32(v7)+64))
				*(*int64)(unsafe.Add(mBase, uint32(l0))) = v21
			} else {
			}
			if l1 != 0 {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v7)+56))
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v23
			} else {
			}
			v25 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v9))), uint32(v25))
			return v20
		}
	} else {
		v20 = *(*int64)(unsafe.Add(mBase, uint32(v7)+48))
		if l0 != 0 {
			v21 = *(*int64)(unsafe.Add(mBase, uint32(v7)+64))
			*(*int64)(unsafe.Add(mBase, uint32(l0))) = v21
		} else {
		}
		if l1 != 0 {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v7)+56))
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v23
		} else {
		}
		v25 = int32(0)
		atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v9))), uint32(v25))
		return v20
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
func F_WalRcvDie(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_WalRcvDie[0]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_XLogWalRcvFlush(m, int32(1), v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		v11 = base.AtomicRmwXchg32(m, v4, int32(1456), int32(1))
		if v11 != 0 {
			F_s_lock(m, v4+int32(1456), int32(_a_F_WalRcvDie_0), int32(792), int32(_a_F_WalRcvDie_1))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				v19 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = v19
				*(*uint8)(unsafe.Add(mBase, uint32(v4)+1453)) = uint8(v19)
				*(*int64)(unsafe.Add(mBase, uint32(v4))) = int64(4294967295)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v4)+1456)), uint32(v19))
				F_ConditionVariableBroadcast(m, v4+int32(12))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, _c_F_WalRcvDie[1]))
					if v33 != 0 {
						v35 = *(*int32)(unsafe.Add(mBase, _c_F_WalRcvDie[2]))
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+64))
						m.T0[v36].(func(*base.Module, int32))(m, v33)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							v40 = *(*int32)(unsafe.Add(mBase, _c_F_WalRcvDie[3]))
							F_SetLatch(m, v40+int32(4))
							mBase = m.M
							return
						}
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, _c_F_WalRcvDie[3]))
						F_SetLatch(m, v40+int32(4))
						mBase = m.M
						return
					}
				}
			}
		} else {
			v19 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = v19
			*(*uint8)(unsafe.Add(mBase, uint32(v4)+1453)) = uint8(v19)
			*(*int64)(unsafe.Add(mBase, uint32(v4))) = int64(4294967295)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v4)+1456)), uint32(v19))
			F_ConditionVariableBroadcast(m, v4+int32(12))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, _c_F_WalRcvDie[1]))
				if v33 != 0 {
					v35 = *(*int32)(unsafe.Add(mBase, _c_F_WalRcvDie[2]))
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+64))
					m.T0[v36].(func(*base.Module, int32))(m, v33)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, _c_F_WalRcvDie[3]))
						F_SetLatch(m, v40+int32(4))
						mBase = m.M
						return
					}
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, _c_F_WalRcvDie[3]))
					F_SetLatch(m, v40+int32(4))
					mBase = m.M
					return
				}
			}
		}
	}
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
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_WalRcvStreaming[0]))
	v7 = int32(1456)
	v8 = v6 + v7
	v11 = base.AtomicRmwXchg32(m, v6, v7, int32(1))
	if v11 != 0 {
		F_s_lock(m, v8, int32(_a_F_WalRcvStreaming_0), int32(133), int32(_a_F_WalRcvStreaming_1))
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
				return base.B2i32(v55 == int32(4)) | base.B2i32(base.Ui32(v55-int32(1)) < base.Ui32(int32(2)))
			} else {
				v26 = int32(1)
				v27 = F_time(m)
				mBase = m.M
				if v27-v19 < int64(11) {
					v55 = v26
					return base.B2i32(v55 == int32(4)) | base.B2i32(base.Ui32(v55-int32(1)) < base.Ui32(int32(2)))
				} else {
					v33 = base.AtomicRmwXchg32(m, v8, int32(0), int32(1))
					if v33 != 0 {
						F_s_lock(m, v8, int32(_a_F_WalRcvStreaming_0), int32(154), int32(_a_F_WalRcvStreaming_1))
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
								return base.B2i32(v55 == int32(4)) | base.B2i32(base.Ui32(v55-int32(1)) < base.Ui32(int32(2)))
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
									return base.B2i32(v55 == int32(4)) | base.B2i32(base.Ui32(v55-int32(1)) < base.Ui32(int32(2)))
								}
							}
						}
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
						if v39 != int32(1) {
							v42 = int32(0)
							atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v8))), uint32(v42))
							v55 = v26
							return base.B2i32(v55 == int32(4)) | base.B2i32(base.Ui32(v55-int32(1)) < base.Ui32(int32(2)))
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
								return base.B2i32(v55 == int32(4)) | base.B2i32(base.Ui32(v55-int32(1)) < base.Ui32(int32(2)))
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
			return base.B2i32(v55 == int32(4)) | base.B2i32(base.Ui32(v55-int32(1)) < base.Ui32(int32(2)))
		} else {
			v26 = int32(1)
			v27 = F_time(m)
			mBase = m.M
			if v27-v19 < int64(11) {
				v55 = v26
				return base.B2i32(v55 == int32(4)) | base.B2i32(base.Ui32(v55-int32(1)) < base.Ui32(int32(2)))
			} else {
				v33 = base.AtomicRmwXchg32(m, v8, int32(0), int32(1))
				if v33 != 0 {
					F_s_lock(m, v8, int32(_a_F_WalRcvStreaming_0), int32(154), int32(_a_F_WalRcvStreaming_1))
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
							return base.B2i32(v55 == int32(4)) | base.B2i32(base.Ui32(v55-int32(1)) < base.Ui32(int32(2)))
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
								return base.B2i32(v55 == int32(4)) | base.B2i32(base.Ui32(v55-int32(1)) < base.Ui32(int32(2)))
							}
						}
					}
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
					if v39 != int32(1) {
						v42 = int32(0)
						atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v8))), uint32(v42))
						v55 = v26
						return base.B2i32(v55 == int32(4)) | base.B2i32(base.Ui32(v55-int32(1)) < base.Ui32(int32(2)))
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
							return base.B2i32(v55 == int32(4)) | base.B2i32(base.Ui32(v55-int32(1)) < base.Ui32(int32(2)))
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
				F_errfinish(m, int32(_a_F_WalSndKeepalive_1), int32(4087), int32(_a_F_WalSndKeepalive_2))
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					v18 = int32(_a_F_WalSndKeepalive_3)
					v19 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[0]))
					v20 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v20)
					*(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[1])) = v20
					*(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[2])) = v20
					F_enlargeStringInfo(m, int32(_a_F_WalSndKeepalive_3), int32(1))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						v30 = int32(_a_F_WalSndKeepalive_4)
						v31 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[2]))
						v32 = int32(_a_F_WalSndKeepalive_3)
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
							F_enlargeStringInfo(m, int32(_a_F_WalSndKeepalive_3), int32(8))
							mBase = m.M
							v116 = m.ExcPending
							if v116 != 0 {
								return
							} else {
								v117 = int32(_a_F_WalSndKeepalive_4)
								v118 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[2]))
								v119 = int32(_a_F_WalSndKeepalive_3)
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
									v166 = int32(_a_F_WalSndKeepalive_4)
									v167 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[2]))
									v168 = int32(_a_F_WalSndKeepalive_3)
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
			v18 = int32(_a_F_WalSndKeepalive_3)
			v19 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[0]))
			v20 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v20)
			*(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[1])) = v20
			*(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[2])) = v20
			F_enlargeStringInfo(m, int32(_a_F_WalSndKeepalive_3), int32(1))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				v30 = int32(_a_F_WalSndKeepalive_4)
				v31 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[2]))
				v32 = int32(_a_F_WalSndKeepalive_3)
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
					F_enlargeStringInfo(m, int32(_a_F_WalSndKeepalive_3), int32(8))
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return
					} else {
						v117 = int32(_a_F_WalSndKeepalive_4)
						v118 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[2]))
						v119 = int32(_a_F_WalSndKeepalive_3)
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
							v166 = int32(_a_F_WalSndKeepalive_4)
							v167 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndKeepalive[2]))
							v168 = int32(_a_F_WalSndKeepalive_3)
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
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
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
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int64
	_ = v160
	var v163 int64
	_ = v163
	var v164 int64
	_ = v164
	var v166 int32
	_ = v166
	var v170 int64
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v188 int64
	_ = v188
	var v192 int32
	_ = v192
	var v196 int64
	_ = v196
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int64
	_ = v220
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
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
	var v270 int64
	_ = v270
	var v271 int64
	_ = v271
	var v279 int64
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v286 int64
	_ = v286
	var v290 int32
	_ = v290
	var v299 int64
	_ = v299
	var v305 int64
	_ = v305
	var v314 int64
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int64
	_ = v341
	var v344 int32
	_ = v344
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
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
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
	goto L4
L2:
	;
	F_WalSndShutdown(m)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L10
	} else {
		goto L108
	}
L3:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(56)
	F_EndCommand(m, v12+int32(16), int32(2))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L10
	} else {
		goto L105
	}
L4:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = int32(0)
	goto L6
L5:
	;
	m.G0 = v12 + int32(32)
	return
L6:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[3]))
	if v53 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[4]))
	if v57 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	return
L11:
	;
	goto L9
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[4])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L10
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	F_ProcessRepliesIfAny(m)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L10
	} else {
		goto L17
	}
L15:
	;
	F_SyncRepInitConfig(m)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalSndLoop[5])))
	if v69 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L5
L19:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[6]))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v88 = m.T0[v87].(func(*base.Module) int32)(m)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L10
	} else {
		goto L25
	}
L20:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalSndLoop[7])))
	if v73&int32(1) == int32(0) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[6]))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	v81 = m.T0[v80].(func(*base.Module) int32)(m)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L10
	} else {
		goto L22
	}
L22:
	;
	if v81 == int32(0) {
		goto L18
	} else {
		goto L23
	}
L23:
	;
	goto L19
L24:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[6]))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
	v100 = m.T0[v99].(func(*base.Module) int32)(m)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L10
	} else {
		goto L30
	}
L25:
	;
	if v88 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	m.T0[l0].(func(*base.Module))(m)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L10
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v95 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_WalSndLoop[8])) = uint8(v95)
	goto L24
L29:
	;
	goto L24
L30:
	;
	if v100 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalSndLoop[8])))
	if v103 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v188 = *(*int64)(unsafe.Add(mBase, _c_F_WalSndLoop[0]))
	if v188 <= int64(0) {
		goto L61
	} else {
		goto L62
	}
L33:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[6]))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+12))
	v109 = m.T0[v108].(func(*base.Module) int32)(m)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L10
	} else {
		goto L34
	}
L34:
	;
	if v109 != 0 {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[9]))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	if v113 != int32(2) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v153 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[10]))
	if v153 == int32(0) {
		goto L32
	} else {
		goto L49
	}
L37:
	;
	v118 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L10
	} else {
		goto L38
	}
L38:
	;
	if v118 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v121 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[11]))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v121
	F_errmsg_internal(m, int32(_a_F_WalSndLoop_0), v12)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L10
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v132 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[9]))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	if v133 == int32(3) {
		goto L36
	} else {
		goto L44
	}
L42:
	;
	F_errfinish(m, int32(_a_F_WalSndLoop_1), int32(2869), int32(_a_F_WalSndLoop_2))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L10
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v138 = base.AtomicRmwXchg32(m, v132, int32(76), int32(1))
	if v138 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	F_s_lock(m, v132+int32(76), int32(_a_F_WalSndLoop_1), int32(3869), int32(_a_F_WalSndLoop_3))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L10
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v132)+4)) = int32(3)
	v148 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v132)+76)), uint32(v148))
	goto L36
L48:
	;
	goto L47
L49:
	;
	m.T0[l0].(func(*base.Module))(m)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L10
	} else {
		goto L50
	}
L50:
	;
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[9]))
	v160 = *(*int64)(unsafe.Add(mBase, uint32(v159)+32))
	if v160 == int64(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v163 = *(*int64)(unsafe.Add(mBase, uint32(v159)+24))
	v164 = v163
	goto L53
L52:
	;
	v164 = v160
	goto L53
L53:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalSndLoop[8])))
	if v166 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalSndLoop[1])))
	if v180 != 0 {
		goto L32
	} else {
		goto L59
	}
L55:
	;
	v170 = *(*int64)(unsafe.Add(mBase, _c_F_WalSndLoop[12]))
	if v170 != v164 {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[6]))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+12))
	v175 = m.T0[v174].(func(*base.Module) int32)(m)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L10
	} else {
		goto L57
	}
L57:
	;
	if v175 == int32(0) {
		goto L3
	} else {
		goto L58
	}
L58:
	;
	goto L54
L59:
	;
	F_WalSndKeepalive(m, int32(1), int64(0))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L10
	} else {
		goto L60
	}
L60:
	;
	goto L32
L61:
	;
	if l0 == int32(1036) {
		goto L77
	} else {
		goto L78
	}
L62:
	;
	v192 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[13]))
	if v192 <= int32(0) {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v196 = *(*int64)(unsafe.Add(mBase, _c_F_WalSndLoop[14]))
	if base.I64_extend_i32_u(v192)*int64(1000)+v188 <= v196 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v204 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L10
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalSndLoop[1])))
	if v218 != 0 {
		goto L61
	} else {
		goto L71
	}
L67:
	;
	if v204 == int32(0) {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	F_errmsg(m, int32(_a_F_WalSndLoop_4), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L10
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_WalSndLoop_1), int32(2789), int32(_a_F_WalSndLoop_5))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L10
	} else {
		goto L70
	}
L70:
	;
	goto L2
L71:
	;
	v220 = *(*int64)(unsafe.Add(mBase, _c_F_WalSndLoop[14]))
	if v220 < base.I64_extend_i32_u(int32(base.Ui32(v192)>>(uint(int32(1))%32)))*int64(1000)+v188 {
		goto L61
	} else {
		goto L72
	}
L72:
	;
	F_WalSndKeepalive(m, int32(1), int64(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L10
	} else {
		goto L73
	}
L73:
	;
	v233 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[6]))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)+8))
	v235 = m.T0[v234].(func(*base.Module) int32)(m)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L10
	} else {
		goto L74
	}
L74:
	;
	if v235 != 0 {
		goto L2
	} else {
		goto L75
	}
L75:
	;
	goto L61
L76:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalSndLoop[5])))
	if v260 != 0 {
		goto L83
	} else {
		goto L84
	}
L77:
	;
	v251 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[6]))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)+12))
	v253 = m.T0[v252].(func(*base.Module) int32)(m)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L10
	} else {
		goto L81
	}
L78:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalSndLoop[8])))
	if v239&int32(1) == int32(0) {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalSndLoop[7])))
	if v245&int32(1) == int32(0) {
		goto L76
	} else {
		goto L80
	}
L80:
	;
	goto L77
L81:
	;
	if v253 == int32(0) {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	goto L76
L83:
	;
	v261 = int32(0)
	goto L85
L84:
	;
	v261 = int32(2)
	goto L85
L85:
	;
	v265 = m.G0
	v266 = int32(16)
	v267 = v265 - v266
	m.G0 = v267
	F_gettimeofday(m, v267)
	mBase = m.M
	v270 = *(*int64)(unsafe.Add(mBase, uint32(v267)))
	v271 = int64(*(*int32)(unsafe.Add(mBase, uint32(v267)+8)))
	m.G0 = v267 + v266
	v279 = v271 + v270*int64(1000000) - int64(946684800000000)
	goto L86
L86:
	;
	v280 = int32(_a_F_WalSndLoop_6)
	v282 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[13]))
	if v282 <= int32(0) {
		v318 = v280
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v323 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[6]))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v323)+12))
	v325 = m.T0[v324].(func(*base.Module) int32)(m)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L10
	} else {
		goto L94
	}
L88:
	;
	v286 = *(*int64)(unsafe.Add(mBase, _c_F_WalSndLoop[0]))
	if v286 <= int64(0) {
		v318 = v280
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalSndLoop[1])))
	v299 = base.I64_extend_i32_u(int32(base.Ui32(v282)>>(uint((v290^int32(-1))&int32(1))%32)))*int64(1000) + v286
	if v299 <= v279 {
		v317 = int32(0)
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v318 = v317
	goto L87
L91:
	;
	goto L90
L92:
	;
	v305 = v299 - v279
	if base.B2i32(int64(0) < v279)^base.B2i32(v305 < v299)|base.B2i32(int64(2147483646000) < v305) != 0 {
		v317 = int32(2147483647)
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v314 = base.I64_div_s(v305+int64(999), int64(1000))
	v317 = base.I32_wrap_i64(v314)
	goto L91
L94:
	;
	if v325 != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v327 = v261 | int32(4)
	goto L97
L96:
	;
	v327 = v261
	goto L97
L97:
	;
	goto L98
L98:
	;
	if base.I64_extend_i32_s(int32(1000))*int64(1000) <= v279-v46 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	F_pgstat_flush_io(m, int32(0))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L10
	} else {
		goto L102
	}
L100:
	;
	v341 = v46
	goto L101
L101:
	;
	F_WalSndWait(m, v327, v318, int32(83886095))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L10
	} else {
		goto L104
	}
L102:
	;
	v339 = F_pgstat_flush_backend(m, int32(0), int32(1))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L10
	} else {
		goto L103
	}
L103:
	;
	v341 = v279
	goto L101
L104:
	;
	v46 = v341
	goto L4
L105:
	;
	v358 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndLoop[6]))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v358)+4))
	v360 = m.T0[v359].(func(*base.Module) int32)(m)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L10
	} else {
		goto L106
	}
L106:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L10
	} else {
		goto L107
	}
L107:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L108:
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
