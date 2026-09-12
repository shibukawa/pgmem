package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetWalRcvFlushRecPtr(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v24 int32
	_ = v24
	v8 = *(*int32)(unsafe.Add(mBase, _consts[253]))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+1456))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+1456)) = int32(1)
	v13 = v8 + int32(1456)
	if v9 != 0 {
		F_s_lock(m, v13, int32(520223), int32(337), int32(218005))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int64(0)
		} else {
			v21 = *(*int64)(unsafe.Add(mBase, uint32(v8)+48))
			if l0 != 0 {
				v22 = *(*int64)(unsafe.Add(mBase, uint32(v8)+64))
				*(*int64)(unsafe.Add(mBase, uint32(l0))) = v22
			} else {
			}
			if l1 != 0 {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v8)+56))
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v24
			} else {
			}
			*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(0)
			return v21
		}
	} else {
		v21 = *(*int64)(unsafe.Add(mBase, uint32(v8)+48))
		if l0 != 0 {
			v22 = *(*int64)(unsafe.Add(mBase, uint32(v8)+64))
			*(*int64)(unsafe.Add(mBase, uint32(l0))) = v22
		} else {
		}
		if l1 != 0 {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v8)+56))
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v24
		} else {
		}
		*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(0)
		return v21
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
	v3 = *(*int32)(unsafe.Add(mBase, _consts[218]))
	v5 = int32(*(*uint8)(unsafe.Add(mBase, _consts[219])))
	if v5 != 0 {
		F_LWLockReleaseClearVar(m, v3, v3+int32(16))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, _consts[218]))
			F_LWLockReleaseClearVar(m, v11+int32(128), v11+int32(144))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, _consts[218]))
				F_LWLockReleaseClearVar(m, v19+int32(256), v19+int32(272))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, _consts[218]))
					F_LWLockReleaseClearVar(m, v27+int32(384), v27+int32(400))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, _consts[218]))
						F_LWLockReleaseClearVar(m, v35+int32(512), v35+int32(528))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							v43 = *(*int32)(unsafe.Add(mBase, _consts[218]))
							F_LWLockReleaseClearVar(m, v43+int32(640), v43+int32(656))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								v51 = *(*int32)(unsafe.Add(mBase, _consts[218]))
								F_LWLockReleaseClearVar(m, v51+int32(768), v51+int32(784))
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return
								} else {
									v59 = *(*int32)(unsafe.Add(mBase, _consts[218]))
									F_LWLockReleaseClearVar(m, v59+int32(896), v59+int32(912))
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return
									} else {
										v67 = int32(0)
										*(*uint8)(unsafe.Add(mBase, _consts[219])) = uint8(v67)
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
		v70 = *(*int32)(unsafe.Add(mBase, _consts[220]))
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
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	v4 = *(*int32)(unsafe.Add(mBase, _consts[253]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_XLogWalRcvFlush(m, int32(1), v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v4)+1456))
		*(*int32)(unsafe.Add(mBase, uint32(v4)+1456)) = int32(1)
		if v9 != 0 {
			F_s_lock(m, v4+int32(1456), int32(520975), int32(792), int32(419127))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				v19 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = v19
				*(*int32)(unsafe.Add(mBase, uint32(v4)+1456)) = v19
				*(*uint8)(unsafe.Add(mBase, uint32(v4)+1453)) = uint8(v19)
				*(*int64)(unsafe.Add(mBase, uint32(v4))) = int64(4294967295)
				F_ConditionVariableBroadcast(m, v4+int32(12))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, _consts[678]))
					if v32 != 0 {
						v34 = *(*int32)(unsafe.Add(mBase, _consts[662]))
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+64))
						m.T0[v35].(func(*base.Module, int32))(m, v32)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							F_WakeupRecovery(m)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return
							} else {
								return
							}
						}
					} else {
						F_WakeupRecovery(m)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		} else {
			v19 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = v19
			*(*int32)(unsafe.Add(mBase, uint32(v4)+1456)) = v19
			*(*uint8)(unsafe.Add(mBase, uint32(v4)+1453)) = uint8(v19)
			*(*int64)(unsafe.Add(mBase, uint32(v4))) = int64(4294967295)
			F_ConditionVariableBroadcast(m, v4+int32(12))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, _consts[678]))
				if v32 != 0 {
					v34 = *(*int32)(unsafe.Add(mBase, _consts[662]))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+64))
					m.T0[v35].(func(*base.Module, int32))(m, v32)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						F_WakeupRecovery(m)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return
						} else {
							return
						}
					}
				} else {
					F_WakeupRecovery(m)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						return
					}
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
	v6 = *(*int32)(unsafe.Add(mBase, _consts[253]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+1456))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+1456)) = int32(1)
	v11 = v6 + int32(1456)
	if v7 != 0 {
		F_s_lock(m, v11, int32(520223), int32(133), int32(353595))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6)+1456)) = int32(0)
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
			if v21 != int32(1) {
				v53 = v21
				return base.B2i32(v53 == int32(4)) | base.B2i32(base.Ui32(v53-int32(1)) < base.Ui32(int32(2)))
			} else {
				v24 = *(*int64)(unsafe.Add(mBase, uint32(v6)+24))
				v26 = F___time(m)
				mBase = m.M
				if v26-v24 < int64(11) {
					v53 = int32(1)
					return base.B2i32(v53 == int32(4)) | base.B2i32(base.Ui32(v53-int32(1)) < base.Ui32(int32(2)))
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(1)
					if v30 != 0 {
						F_s_lock(m, v11, int32(520223), int32(154), int32(353595))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
							if v38 != int32(1) {
								*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(0)
								v53 = int32(1)
								return base.B2i32(v53 == int32(4)) | base.B2i32(base.Ui32(v53-int32(1)) < base.Ui32(int32(2)))
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
									return base.B2i32(v53 == int32(4)) | base.B2i32(base.Ui32(v53-int32(1)) < base.Ui32(int32(2)))
								}
							}
						}
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
						if v38 != int32(1) {
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(0)
							v53 = int32(1)
							return base.B2i32(v53 == int32(4)) | base.B2i32(base.Ui32(v53-int32(1)) < base.Ui32(int32(2)))
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
								return base.B2i32(v53 == int32(4)) | base.B2i32(base.Ui32(v53-int32(1)) < base.Ui32(int32(2)))
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
			return base.B2i32(v53 == int32(4)) | base.B2i32(base.Ui32(v53-int32(1)) < base.Ui32(int32(2)))
		} else {
			v24 = *(*int64)(unsafe.Add(mBase, uint32(v6)+24))
			v26 = F___time(m)
			mBase = m.M
			if v26-v24 < int64(11) {
				v53 = int32(1)
				return base.B2i32(v53 == int32(4)) | base.B2i32(base.Ui32(v53-int32(1)) < base.Ui32(int32(2)))
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(1)
				if v30 != 0 {
					F_s_lock(m, v11, int32(520223), int32(154), int32(353595))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
						if v38 != int32(1) {
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(0)
							v53 = int32(1)
							return base.B2i32(v53 == int32(4)) | base.B2i32(base.Ui32(v53-int32(1)) < base.Ui32(int32(2)))
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
								return base.B2i32(v53 == int32(4)) | base.B2i32(base.Ui32(v53-int32(1)) < base.Ui32(int32(2)))
							}
						}
					}
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
					if v38 != int32(1) {
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(0)
						v53 = int32(1)
						return base.B2i32(v53 == int32(4)) | base.B2i32(base.Ui32(v53-int32(1)) < base.Ui32(int32(2)))
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
							return base.B2i32(v53 == int32(4)) | base.B2i32(base.Ui32(v53-int32(1)) < base.Ui32(int32(2)))
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
			F_errmsg_internal(m, int32(362470), int32(0))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				F_errfinish(m, int32(521307), int32(4087), int32(362500))
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					v18 = int32(4472512)
					v19 = *(*int32)(unsafe.Add(mBase, _consts[718]))
					v20 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v20)
					*(*int32)(unsafe.Add(mBase, _consts[719])) = v20
					*(*int32)(unsafe.Add(mBase, _consts[720])) = v20
					F_enlargeStringInfo(m, int32(4472512), int32(1))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						v30 = int32(4472516)
						v31 = *(*int32)(unsafe.Add(mBase, _consts[720]))
						v32 = int32(4472512)
						v33 = *(*int32)(unsafe.Add(mBase, _consts[718]))
						v35 = int32(107)
						*(*uint8)(unsafe.Add(mBase, uint32(v31+v33))) = uint8(v35)
						*(*int32)(unsafe.Add(mBase, _consts[720])) = v31 + int32(1)
						v42 = *(*int64)(unsafe.Add(mBase, _consts[716]))
						F_enlargeStringInfo(m, v32, int32(8))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							v48 = *(*int32)(unsafe.Add(mBase, _consts[720]))
							v50 = *(*int32)(unsafe.Add(mBase, _consts[718]))
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
							*(*int32)(unsafe.Add(mBase, _consts[720])) = v48 + int32(8)
							v98 = m.G0
							v99 = int32(16)
							v100 = v98 - v99
							m.G0 = v100
							F___gettimeofday(m, v100)
							mBase = m.M
							v103 = *(*int64)(unsafe.Add(mBase, uint32(v100)))
							v104 = int64(*(*int32)(unsafe.Add(mBase, uint32(v100)+8)))
							m.G0 = v100 + v99
							v112 = v104 + v103*int64(1000000) - int64(946684800000000)
							F_enlargeStringInfo(m, int32(4472512), int32(8))
							mBase = m.M
							v116 = m.ExcPending
							if v116 != 0 {
								return
							} else {
								v117 = int32(4472516)
								v118 = *(*int32)(unsafe.Add(mBase, _consts[720]))
								v119 = int32(4472512)
								v120 = *(*int32)(unsafe.Add(mBase, _consts[718]))
								v122 = int64(56)
								v124 = int64(65280)
								v126 = int64(40)
								v129 = int64(16711680)
								v131 = int64(24)
								v133 = int64(4278190080)
								v135 = int64(8)
								*(*int64)(unsafe.Add(mBase, uint32(v118+v120))) = v112<<(uint(v122)%64) | v112&v124<<(uint(v126)%64) | (v112&v129<<(uint(v131)%64) | v112&v133<<(uint(v135)%64)) | (int64(base.Ui64(v112)>>(uint(v135)%64))&v133 | int64(base.Ui64(v112)>>(uint(v131)%64))&v129 | (int64(base.Ui64(v112)>>(uint(v126)%64))&v124 | int64(base.Ui64(v112)>>(uint(v122)%64))))
								*(*int32)(unsafe.Add(mBase, _consts[720])) = v118 + int32(8)
								F_enlargeStringInfo(m, v119, int32(1))
								mBase = m.M
								v165 = m.ExcPending
								if v165 != 0 {
									return
								} else {
									v166 = int32(4472516)
									v167 = *(*int32)(unsafe.Add(mBase, _consts[720]))
									v168 = int32(4472512)
									v169 = *(*int32)(unsafe.Add(mBase, _consts[718]))
									*(*uint8)(unsafe.Add(mBase, uint32(v167+v169))) = uint8(v1)
									v174 = v167 + int32(1)
									*(*int32)(unsafe.Add(mBase, _consts[720])) = v174
									v178 = *(*int32)(unsafe.Add(mBase, _consts[718]))
									v180 = *(*int32)(unsafe.Add(mBase, _consts[707]))
									v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+20))
									m.T0[v181].(func(*base.Module, int32, int32, int32))(m, int32(100), v178, v174)
									mBase = m.M
									v183 = m.ExcPending
									if v183 != 0 {
										return
									} else {
										if v1 != 0 {
											v185 = int32(1)
											*(*uint8)(unsafe.Add(mBase, _consts[708])) = uint8(v185)
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
			v18 = int32(4472512)
			v19 = *(*int32)(unsafe.Add(mBase, _consts[718]))
			v20 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v20)
			*(*int32)(unsafe.Add(mBase, _consts[719])) = v20
			*(*int32)(unsafe.Add(mBase, _consts[720])) = v20
			F_enlargeStringInfo(m, int32(4472512), int32(1))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				v30 = int32(4472516)
				v31 = *(*int32)(unsafe.Add(mBase, _consts[720]))
				v32 = int32(4472512)
				v33 = *(*int32)(unsafe.Add(mBase, _consts[718]))
				v35 = int32(107)
				*(*uint8)(unsafe.Add(mBase, uint32(v31+v33))) = uint8(v35)
				*(*int32)(unsafe.Add(mBase, _consts[720])) = v31 + int32(1)
				v42 = *(*int64)(unsafe.Add(mBase, _consts[716]))
				F_enlargeStringInfo(m, v32, int32(8))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, _consts[720]))
					v50 = *(*int32)(unsafe.Add(mBase, _consts[718]))
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
					*(*int32)(unsafe.Add(mBase, _consts[720])) = v48 + int32(8)
					v98 = m.G0
					v99 = int32(16)
					v100 = v98 - v99
					m.G0 = v100
					F___gettimeofday(m, v100)
					mBase = m.M
					v103 = *(*int64)(unsafe.Add(mBase, uint32(v100)))
					v104 = int64(*(*int32)(unsafe.Add(mBase, uint32(v100)+8)))
					m.G0 = v100 + v99
					v112 = v104 + v103*int64(1000000) - int64(946684800000000)
					F_enlargeStringInfo(m, int32(4472512), int32(8))
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return
					} else {
						v117 = int32(4472516)
						v118 = *(*int32)(unsafe.Add(mBase, _consts[720]))
						v119 = int32(4472512)
						v120 = *(*int32)(unsafe.Add(mBase, _consts[718]))
						v122 = int64(56)
						v124 = int64(65280)
						v126 = int64(40)
						v129 = int64(16711680)
						v131 = int64(24)
						v133 = int64(4278190080)
						v135 = int64(8)
						*(*int64)(unsafe.Add(mBase, uint32(v118+v120))) = v112<<(uint(v122)%64) | v112&v124<<(uint(v126)%64) | (v112&v129<<(uint(v131)%64) | v112&v133<<(uint(v135)%64)) | (int64(base.Ui64(v112)>>(uint(v135)%64))&v133 | int64(base.Ui64(v112)>>(uint(v131)%64))&v129 | (int64(base.Ui64(v112)>>(uint(v126)%64))&v124 | int64(base.Ui64(v112)>>(uint(v122)%64))))
						*(*int32)(unsafe.Add(mBase, _consts[720])) = v118 + int32(8)
						F_enlargeStringInfo(m, v119, int32(1))
						mBase = m.M
						v165 = m.ExcPending
						if v165 != 0 {
							return
						} else {
							v166 = int32(4472516)
							v167 = *(*int32)(unsafe.Add(mBase, _consts[720]))
							v168 = int32(4472512)
							v169 = *(*int32)(unsafe.Add(mBase, _consts[718]))
							*(*uint8)(unsafe.Add(mBase, uint32(v167+v169))) = uint8(v1)
							v174 = v167 + int32(1)
							*(*int32)(unsafe.Add(mBase, _consts[720])) = v174
							v178 = *(*int32)(unsafe.Add(mBase, _consts[718]))
							v180 = *(*int32)(unsafe.Add(mBase, _consts[707]))
							v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+20))
							m.T0[v181].(func(*base.Module, int32, int32, int32))(m, int32(100), v178, v174)
							mBase = m.M
							v183 = m.ExcPending
							if v183 != 0 {
								return
							} else {
								if v1 != 0 {
									v185 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _consts[708])) = uint8(v185)
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
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int64
	_ = v158
	var v161 int64
	_ = v161
	var v162 int64
	_ = v162
	var v164 int32
	_ = v164
	var v168 int64
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v187 int64
	_ = v187
	var v191 int32
	_ = v191
	var v195 int64
	_ = v195
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int64
	_ = v219
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
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int64
	_ = v267
	var v268 int64
	_ = v268
	var v276 int64
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v283 int64
	_ = v283
	var v287 int32
	_ = v287
	var v296 int64
	_ = v296
	var v299 int32
	_ = v299
	var v302 int64
	_ = v302
	var v310 int64
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int64
	_ = v337
	var v340 int32
	_ = v340
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v18 = m.G0
	v19 = int32(16)
	v20 = v18 - v19
	m.G0 = v20
	F___gettimeofday(m, v20)
	mBase = m.M
	v23 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	v24 = int64(*(*int32)(unsafe.Add(mBase, uint32(v20)+8)))
	m.G0 = v20 + v19
	goto L1
L1:
	;
	*(*int64)(unsafe.Add(mBase, _consts[704])) = v24 + v23*int64(1000000) - int64(946684800000000)
	v35 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[708])) = uint8(v35)
	v46 = int64(0)
	goto L4
L2:
	;
	F_WalSndShutdown(m)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L10
	} else {
		goto L109
	}
L3:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(56)
	F_EndCommand(m, v12+int32(16), int32(2))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L10
	} else {
		goto L106
	}
L4:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _consts[519]))
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = int32(0)
	goto L6
L5:
	;
	m.G0 = v12 + int32(32)
	return
L6:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _consts[48]))
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
	v57 = *(*int32)(unsafe.Add(mBase, _consts[709]))
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
	*(*int32)(unsafe.Add(mBase, _consts[709])) = int32(0)
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
	v69 = int32(*(*uint8)(unsafe.Add(mBase, _consts[710])))
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
	v84 = *(*int32)(unsafe.Add(mBase, _consts[707]))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
	v86 = m.T0[v85].(func(*base.Module) int32)(m)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L10
	} else {
		goto L25
	}
L20:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, _consts[711])))
	if v73 == int32(0) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _consts[707]))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	v79 = m.T0[v78].(func(*base.Module) int32)(m)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L10
	} else {
		goto L22
	}
L22:
	;
	if v79 == int32(0) {
		goto L18
	} else {
		goto L23
	}
L23:
	;
	goto L19
L24:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _consts[707]))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
	v98 = m.T0[v97].(func(*base.Module) int32)(m)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L10
	} else {
		goto L30
	}
L25:
	;
	if v86 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	m.T0[l0].(func(*base.Module))(m)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L10
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v93 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[712])) = uint8(v93)
	goto L24
L29:
	;
	goto L24
L30:
	;
	if v98 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, _consts[712])))
	if v101 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v187 = *(*int64)(unsafe.Add(mBase, _consts[704]))
	if v187 <= int64(0) {
		goto L61
	} else {
		goto L62
	}
L33:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _consts[707]))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+12))
	v107 = m.T0[v106].(func(*base.Module) int32)(m)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L10
	} else {
		goto L34
	}
L34:
	;
	if v107 != 0 {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v110 = *(*int32)(unsafe.Add(mBase, _consts[713]))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	if v111 != int32(2) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v151 = *(*int32)(unsafe.Add(mBase, _consts[714]))
	if v151 == int32(0) {
		goto L32
	} else {
		goto L49
	}
L37:
	;
	v116 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L10
	} else {
		goto L38
	}
L38:
	;
	if v116 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _consts[715]))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v119
	F_errmsg_internal(m, int32(226233), v12)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L10
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v130 = *(*int32)(unsafe.Add(mBase, _consts[713]))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	if v131 == int32(3) {
		goto L36
	} else {
		goto L44
	}
L42:
	;
	F_errfinish(m, int32(521307), int32(2869), int32(247347))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L10
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v130)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v130)+76)) = int32(1)
	if v134 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	F_s_lock(m, v130+int32(76), int32(521307), int32(3869), int32(372501))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L10
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v130)+76)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v130)+4)) = int32(3)
	goto L36
L48:
	;
	goto L47
L49:
	;
	m.T0[l0].(func(*base.Module))(m)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L10
	} else {
		goto L50
	}
L50:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _consts[713]))
	v158 = *(*int64)(unsafe.Add(mBase, uint32(v157)+32))
	if v158 == int64(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v161 = *(*int64)(unsafe.Add(mBase, uint32(v157)+24))
	v162 = v161
	goto L53
L52:
	;
	v162 = v158
	goto L53
L53:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, _consts[712])))
	if v164 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, _consts[708])))
	if v178 != 0 {
		goto L32
	} else {
		goto L59
	}
L55:
	;
	v168 = *(*int64)(unsafe.Add(mBase, _consts[716]))
	if v168 != v162 {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v171 = *(*int32)(unsafe.Add(mBase, _consts[707]))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)+12))
	v173 = m.T0[v172].(func(*base.Module) int32)(m)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L10
	} else {
		goto L57
	}
L57:
	;
	if v173 == int32(0) {
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
	v182 = m.ExcPending
	if v182 != 0 {
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
	v191 = *(*int32)(unsafe.Add(mBase, _consts[705]))
	if v191 <= int32(0) {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v195 = *(*int64)(unsafe.Add(mBase, _consts[717]))
	if base.I64_extend_i32_u(v191)*int64(1000)+v187 <= v195 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v203 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L10
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, _consts[708])))
	if v217 != 0 {
		goto L61
	} else {
		goto L71
	}
L67:
	;
	if v203 == int32(0) {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	F_errmsg(m, int32(72020), int32(0))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L10
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(521307), int32(2789), int32(73813))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L10
	} else {
		goto L70
	}
L70:
	;
	goto L2
L71:
	;
	v219 = *(*int64)(unsafe.Add(mBase, _consts[717]))
	if v219 < base.I64_extend_i32_u(int32(base.Ui32(v191)>>(uint(int32(1))%32)))*int64(1000)+v187 {
		goto L61
	} else {
		goto L72
	}
L72:
	;
	F_WalSndKeepalive(m, int32(1), int64(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L10
	} else {
		goto L73
	}
L73:
	;
	v232 = *(*int32)(unsafe.Add(mBase, _consts[707]))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)+8))
	v234 = m.T0[v233].(func(*base.Module) int32)(m)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L10
	} else {
		goto L74
	}
L74:
	;
	if v234 != 0 {
		goto L2
	} else {
		goto L75
	}
L75:
	;
	goto L61
L76:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, _consts[710])))
	if v257 != 0 {
		goto L83
	} else {
		goto L84
	}
L77:
	;
	v248 = *(*int32)(unsafe.Add(mBase, _consts[707]))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v248)+12))
	v250 = m.T0[v249].(func(*base.Module) int32)(m)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L10
	} else {
		goto L81
	}
L78:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, _consts[712])))
	if v238&int32(1) == int32(0) {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, _consts[711])))
	if v244 == int32(0) {
		goto L76
	} else {
		goto L80
	}
L80:
	;
	goto L77
L81:
	;
	if v250 == int32(0) {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	goto L76
L83:
	;
	v258 = int32(0)
	goto L85
L84:
	;
	v258 = int32(2)
	goto L85
L85:
	;
	v262 = m.G0
	v263 = int32(16)
	v264 = v262 - v263
	m.G0 = v264
	F___gettimeofday(m, v264)
	mBase = m.M
	v267 = *(*int64)(unsafe.Add(mBase, uint32(v264)))
	v268 = int64(*(*int32)(unsafe.Add(mBase, uint32(v264)+8)))
	m.G0 = v264 + v263
	v276 = v268 + v267*int64(1000000) - int64(946684800000000)
	goto L86
L86:
	;
	v277 = int32(10000)
	v279 = *(*int32)(unsafe.Add(mBase, _consts[705]))
	if v279 <= int32(0) {
		v314 = v277
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v319 = *(*int32)(unsafe.Add(mBase, _consts[707]))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v319)+12))
	v321 = m.T0[v320].(func(*base.Module) int32)(m)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L10
	} else {
		goto L95
	}
L88:
	;
	v283 = *(*int64)(unsafe.Add(mBase, _consts[704]))
	if v283 <= int64(0) {
		v314 = v277
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v287 = int32(*(*uint8)(unsafe.Add(mBase, _consts[708])))
	v296 = base.I64_extend_i32_u(int32(base.Ui32(v279)>>(uint((v287^int32(-1))&int32(1))%32)))*int64(1000) + v283
	if v296 <= v276 {
		v313 = int32(0)
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v314 = v313
	goto L87
L91:
	;
	goto L90
L92:
	;
	v299 = int32(2147483647)
	v302 = v296 - v276
	if base.B2i32(int64(0) < v276)^base.B2i32(v302 < v296) != 0 {
		v313 = v299
		goto L91
	} else {
		goto L93
	}
L93:
	;
	if int64(2147483646000) < v302 {
		v313 = v299
		goto L91
	} else {
		goto L94
	}
L94:
	;
	v310 = base.I64_div_s(v302+int64(999), int64(1000))
	v313 = base.I32_wrap_i64(v310)
	goto L91
L95:
	;
	if v321 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v323 = v258 | int32(4)
	goto L98
L97:
	;
	v323 = v258
	goto L98
L98:
	;
	goto L99
L99:
	;
	if base.I64_extend_i32_s(int32(1000))*int64(1000) <= v276-v46 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	F_pgstat_flush_io(m, int32(0))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L10
	} else {
		goto L103
	}
L101:
	;
	v337 = v46
	goto L102
L102:
	;
	F_WalSndWait(m, v323, v314, int32(83886095))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L10
	} else {
		goto L105
	}
L103:
	;
	v335 = F_pgstat_flush_backend(m, int32(0), int32(1))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L10
	} else {
		goto L104
	}
L104:
	;
	v337 = v276
	goto L102
L105:
	;
	v46 = v337
	goto L4
L106:
	;
	v354 = *(*int32)(unsafe.Add(mBase, _consts[707]))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v354)+4))
	v356 = m.T0[v355].(func(*base.Module) int32)(m)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L10
	} else {
		goto L107
	}
L107:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L10
	} else {
		goto L108
	}
L108:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L109:
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
	v2 = *(*int32)(unsafe.Add(mBase, _consts[146]))
	if v2 == int32(2) {
		*(*int32)(unsafe.Add(mBase, _consts[146])) = int32(0)
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v33 int64
	_ = v33
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int64
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int64
	_ = v63
	var v67 int64
	_ = v67
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v81 int64
	_ = v81
	var v85 int64
	_ = v85
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v99 int64
	_ = v99
	var v103 int64
	_ = v103
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
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
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int64
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	v14 = m.G0
	v15 = int32(16)
	v16 = v14 - v15
	m.G0 = v16
	F___gettimeofday(m, v16)
	mBase = m.M
	v19 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
	v20 = int64(*(*int32)(unsafe.Add(mBase, uint32(v16)+8)))
	m.G0 = v16 + v15
	v28 = v20 + v19*int64(1000000) - int64(946684800000000)
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+164)))
	if v29 != int32(1) {
	} else {
		v33 = *(*int64)(unsafe.Add(mBase, _consts[691]))
		if base.B2i32(base.I64_extend_i32_s(int32(1000))*int64(1000) <= v28-v33) == int32(0) {
		} else {
			v43 = int32(*(*uint8)(unsafe.Add(mBase, _consts[692])))
			if v43 != int32(1) {
			} else {
				v47 = *(*int32)(unsafe.Add(mBase, _consts[693]))
				v48 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
				if v48 == l1 {
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v47))) = l1
					v52 = v47 + int32(8)
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v47)+uint32(_consts[694])))
					v57 = base.I32_rem_s(v53+int32(1), int32(8192))
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v47)+uint32(_consts[695])))
					if v57 == v58 {
						v62 = v52 + v57<<(uint(int32(4))%32)
						v63 = *(*int64)(unsafe.Add(mBase, uint32(v62)))
						*(*int64)(unsafe.Add(mBase, uint32(v47)+uint32(_consts[696]))) = v63
						v67 = *(*int64)(unsafe.Add(mBase, uint32(v62)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v47)+uint32(_consts[697]))) = v67
						*(*int32)(unsafe.Add(mBase, uint32(v47)+uint32(_consts[695]))) = int32(-1)
					} else {
					}
					v74 = *(*int32)(unsafe.Add(mBase, uint32(v47)+uint32(_consts[698])))
					if v74 == v57 {
						v80 = v52 + v57<<(uint(int32(4))%32)
						v81 = *(*int64)(unsafe.Add(mBase, uint32(v80)))
						*(*int64)(unsafe.Add(mBase, uint32(v47)+uint32(_consts[699]))) = v81
						v85 = *(*int64)(unsafe.Add(mBase, uint32(v80)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v47)+uint32(_consts[700]))) = v85
						*(*int32)(unsafe.Add(mBase, uint32(v47)+uint32(_consts[698]))) = int32(-1)
					} else {
					}
					v92 = *(*int32)(unsafe.Add(mBase, uint32(v47)+uint32(_consts[701])))
					if v92 == v57 {
						v98 = v52 + v57<<(uint(int32(4))%32)
						v99 = *(*int64)(unsafe.Add(mBase, uint32(v98)))
						*(*int64)(unsafe.Add(mBase, uint32(v47)+uint32(_consts[702]))) = v99
						v103 = *(*int64)(unsafe.Add(mBase, uint32(v98)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v47)+uint32(_consts[703]))) = v103
						*(*int32)(unsafe.Add(mBase, uint32(v47)+uint32(_consts[701]))) = int32(-1)
					} else {
					}
					v108 = int32(4)
					*(*int64)(unsafe.Add(mBase, uint32(v52+v53<<(uint(v108)%32)))) = l1
					v112 = *(*int32)(unsafe.Add(mBase, uint32(v47)+uint32(_consts[694])))
					*(*int64)(unsafe.Add(mBase, uint32(v47+v112<<(uint(v108)%32))+16)) = v28
					*(*int32)(unsafe.Add(mBase, uint32(v47)+uint32(_consts[694]))) = v57
				}
			}
			*(*int64)(unsafe.Add(mBase, _consts[691])) = v28
		}
	}
	if l3 == int32(0) {
		if v29 != 0 {
			return
		} else {
			v163 = *(*int64)(unsafe.Add(mBase, _consts[704]))
			v165 = *(*int32)(unsafe.Add(mBase, _consts[705]))
			v167 = base.I32_div_s(v165, int32(2))
			if v28 < v163+base.I64_extend_i32_s(v167)*int64(1000) {
				return
			} else {
				F_ProcessPendingWrites(m)
				mBase = m.M
				v174 = m.ExcPending
				if v174 != 0 {
					return
				} else {
					return
				}
			}
		}
	} else {
		v135 = *(*int32)(unsafe.Add(mBase, _consts[535]))
		if v135 <= int32(0) {
			if v29 != 0 {
				return
			} else {
				v163 = *(*int64)(unsafe.Add(mBase, _consts[704]))
				v165 = *(*int32)(unsafe.Add(mBase, _consts[705]))
				v167 = base.I32_div_s(v165, int32(2))
				if v28 < v163+base.I64_extend_i32_s(v167)*int64(1000) {
					return
				} else {
					F_ProcessPendingWrites(m)
					mBase = m.M
					v174 = m.ExcPending
					if v174 != 0 {
						return
					} else {
						return
					}
				}
			}
		} else {
			v139 = *(*int32)(unsafe.Add(mBase, _consts[174]))
			if v139 < int32(2) {
				if v29 != 0 {
					return
				} else {
					v163 = *(*int64)(unsafe.Add(mBase, _consts[704]))
					v165 = *(*int32)(unsafe.Add(mBase, _consts[705]))
					v167 = base.I32_div_s(v165, int32(2))
					if v28 < v163+base.I64_extend_i32_s(v167)*int64(1000) {
						return
					} else {
						F_ProcessPendingWrites(m)
						mBase = m.M
						v174 = m.ExcPending
						if v174 != 0 {
							return
						} else {
							return
						}
					}
				}
			} else {
				v143 = *(*int32)(unsafe.Add(mBase, _consts[706]))
				v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+48)))
				if v144&int32(2) == int32(0) {
					if v29 != 0 {
						return
					} else {
						v163 = *(*int64)(unsafe.Add(mBase, _consts[704]))
						v165 = *(*int32)(unsafe.Add(mBase, _consts[705]))
						v167 = base.I32_div_s(v165, int32(2))
						if v28 < v163+base.I64_extend_i32_s(v167)*int64(1000) {
							return
						} else {
							F_ProcessPendingWrites(m)
							mBase = m.M
							v174 = m.ExcPending
							if v174 != 0 {
								return
							} else {
								return
							}
						}
					}
				} else {
					F_WalSndKeepalive(m, int32(0), l1)
					mBase = m.M
					v151 = m.ExcPending
					if v151 != 0 {
						return
					} else {
						v153 = *(*int32)(unsafe.Add(mBase, _consts[707]))
						v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)+8))
						v155 = m.T0[v154].(func(*base.Module) int32)(m)
						mBase = m.M
						v156 = m.ExcPending
						if v156 != 0 {
							return
						} else {
							if v155 != 0 {
								F_WalSndShutdown(m)
								mBase = m.M
								v176 = m.ExcPending
								if v176 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v158 = *(*int32)(unsafe.Add(mBase, _consts[707]))
								v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)+12))
								v160 = m.T0[v159].(func(*base.Module) int32)(m)
								mBase = m.M
								v161 = m.ExcPending
								if v161 != 0 {
									return
								} else {
									if v160 != 0 {
										F_ProcessPendingWrites(m)
										mBase = m.M
										v174 = m.ExcPending
										if v174 != 0 {
											return
										} else {
											return
										}
									} else {
										if v29 != 0 {
											return
										} else {
											v163 = *(*int64)(unsafe.Add(mBase, _consts[704]))
											v165 = *(*int32)(unsafe.Add(mBase, _consts[705]))
											v167 = base.I32_div_s(v165, int32(2))
											if v28 < v163+base.I64_extend_i32_s(v167)*int64(1000) {
												return
											} else {
												F_ProcessPendingWrites(m)
												mBase = m.M
												v174 = m.ExcPending
												if v174 != 0 {
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
