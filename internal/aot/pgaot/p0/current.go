package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_GetCurrentDateTime(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	F_GetCurrentTimeUsec(m, l0, v5+int32(12), int32(0))
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		m.G0 = v5 + int32(16)
		return
	}
}
func F_GetCurrentReplayRecPtr(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentReplayRecPtr[0]))
	v9 = base.AtomicRmwXchg32(m, v6, int32(96), int32(1))
	if v9 != 0 {
		F_s_lock(m, v6+int32(96), int32(_a_F_GetCurrentReplayRecPtr_0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentReplayRecPtr[0]))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
			v20 = *(*int64)(unsafe.Add(mBase, uint32(v18)+48))
			v21 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v18)+96)), uint32(v21))
			if l0 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v19
			} else {
			}
			return v20
		}
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentReplayRecPtr[0]))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
		v20 = *(*int64)(unsafe.Add(mBase, uint32(v18)+48))
		v21 = int32(0)
		atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v18)+96)), uint32(v21))
		if l0 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v19
		} else {
		}
		return v20
	}
}
func F_GetCurrentRoleId(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentRoleId[0]))
	v5 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetCurrentRoleId[1])))
	if v5 != 0 {
		v6 = v2
	} else {
		v6 = int32(0)
	}
	return v6
}
func F_GetCurrentTimeUsec(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v11 int64
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int64
	_ = v35
	var v38 int64
	_ = v38
	var v41 int64
	_ = v41
	var v44 int64
	_ = v44
	var v47 int64
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentTimeUsec[0]))
	v9 = *(*int64)(unsafe.Add(mBase, _c_F_GetCurrentTimeUsec[1]))
	v11 = *(*int64)(unsafe.Add(mBase, _c_F_GetCurrentTimeUsec[2]))
	if v9 == v11 {
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentTimeUsec[3]))
		if v7 == v14 {
			v32 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentTimeUsec[4]))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v32
			v35 = *(*int64)(unsafe.Add(mBase, _c_F_GetCurrentTimeUsec[5]))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v35
			v38 = *(*int64)(unsafe.Add(mBase, _c_F_GetCurrentTimeUsec[6]))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v38
			v41 = *(*int64)(unsafe.Add(mBase, _c_F_GetCurrentTimeUsec[7]))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v41
			v44 = *(*int64)(unsafe.Add(mBase, _c_F_GetCurrentTimeUsec[8]))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v44
			v47 = *(*int64)(unsafe.Add(mBase, _c_F_GetCurrentTimeUsec[9]))
			*(*int64)(unsafe.Add(mBase, uint32(l0))) = v47
			v50 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentTimeUsec[10]))
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v50
			if l2 != 0 {
				v53 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentTimeUsec[11]))
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v53
			} else {
			}
			return
		} else {
			v17 = int32(0)
			*(*int32)(unsafe.Add(mBase, _c_F_GetCurrentTimeUsec[3])) = v17
			v23 = F_timestamp2tm(m, v9, int32(_a_F_GetCurrentTimeUsec_0), int32(_a_F_GetCurrentTimeUsec_1), int32(_a_F_GetCurrentTimeUsec_2), v17, v7)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				if v23 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return
						} else {
							F_errmsg(m, int32(_a_F_GetCurrentTimeUsec_3), int32(0))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_GetCurrentTimeUsec_4), int32(433), int32(_a_F_GetCurrentTimeUsec_5))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
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
					*(*int64)(unsafe.Add(mBase, _c_F_GetCurrentTimeUsec[2])) = v9
					v29 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentTimeUsec[0]))
					*(*int32)(unsafe.Add(mBase, _c_F_GetCurrentTimeUsec[3])) = v29
					v32 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentTimeUsec[4]))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v32
					v35 = *(*int64)(unsafe.Add(mBase, _c_F_GetCurrentTimeUsec[5]))
					*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v35
					v38 = *(*int64)(unsafe.Add(mBase, _c_F_GetCurrentTimeUsec[6]))
					*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v38
					v41 = *(*int64)(unsafe.Add(mBase, _c_F_GetCurrentTimeUsec[7]))
					*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v41
					v44 = *(*int64)(unsafe.Add(mBase, _c_F_GetCurrentTimeUsec[8]))
					*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v44
					v47 = *(*int64)(unsafe.Add(mBase, _c_F_GetCurrentTimeUsec[9]))
					*(*int64)(unsafe.Add(mBase, uint32(l0))) = v47
					v50 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentTimeUsec[10]))
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v50
					if l2 != 0 {
						v53 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentTimeUsec[11]))
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v53
					} else {
					}
					return
				}
			}
		}
	} else {
		v17 = int32(0)
		*(*int32)(unsafe.Add(mBase, _c_F_GetCurrentTimeUsec[3])) = v17
		v23 = F_timestamp2tm(m, v9, int32(_a_F_GetCurrentTimeUsec_0), int32(_a_F_GetCurrentTimeUsec_1), int32(_a_F_GetCurrentTimeUsec_2), v17, v7)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			if v23 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_GetCurrentTimeUsec_3), int32(0))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_GetCurrentTimeUsec_4), int32(433), int32(_a_F_GetCurrentTimeUsec_5))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
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
				*(*int64)(unsafe.Add(mBase, _c_F_GetCurrentTimeUsec[2])) = v9
				v29 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentTimeUsec[0]))
				*(*int32)(unsafe.Add(mBase, _c_F_GetCurrentTimeUsec[3])) = v29
				v32 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentTimeUsec[4]))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v32
				v35 = *(*int64)(unsafe.Add(mBase, _c_F_GetCurrentTimeUsec[5]))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v35
				v38 = *(*int64)(unsafe.Add(mBase, _c_F_GetCurrentTimeUsec[6]))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v38
				v41 = *(*int64)(unsafe.Add(mBase, _c_F_GetCurrentTimeUsec[7]))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v41
				v44 = *(*int64)(unsafe.Add(mBase, _c_F_GetCurrentTimeUsec[8]))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v44
				v47 = *(*int64)(unsafe.Add(mBase, _c_F_GetCurrentTimeUsec[9]))
				*(*int64)(unsafe.Add(mBase, uint32(l0))) = v47
				v50 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentTimeUsec[10]))
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v50
				if l2 != 0 {
					v53 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentTimeUsec[11]))
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v53
				} else {
				}
				return
			}
		}
	}
}
func F_GetCurrentVirtualXIDs(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	v3 = int32(0)
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentVirtualXIDs[0]))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v17 = F_palloc_mul(m, int32(8), v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentVirtualXIDs[1]))
		v26 = F_LWLockAcquire(m, v22+int32(512), int32(1))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
			if int32(0) < v28 {
				v34 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentVirtualXIDs[2]))
				v36 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentVirtualXIDs[3]))
				v41 = v28
				v43 = v3
				v44 = v3
				v46 = v34
				for {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v15+int32(36)+v43<<(uint(int32(2))%32))))
					v55 = v36 + v52*int32(768)
					if v55 == v46 {
						v100 = v41
						v101 = v44
						v102 = v46
					} else {
						v58 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentVirtualXIDs[4]))
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
						v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v43))))
						if v61&int32(7) != 0 {
							v100 = v41
							v101 = v44
							v102 = v46
						} else {
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v55)+20))
							v66 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentVirtualXIDs[5]))
							if v64 != v66 {
								v100 = v41
								v101 = v44
								v102 = v46
							} else {
								v68 = *(*int32)(unsafe.Add(mBase, uint32(v55)+52))
								if v68 == int32(0) {
									v100 = v41
									v101 = v44
									v102 = v46
								} else {
									if l0 == int32(0) {
										v84 = *(*int32)(unsafe.Add(mBase, uint32(v55)+44))
										if v84 == int32(0) {
											v100 = v41
											v101 = v44
											v102 = v46
										} else {
											v87 = *(*int32)(unsafe.Add(mBase, uint32(v55)+40))
											v90 = v17 + v44<<(uint(int32(3))%32)
											*(*int32)(unsafe.Add(mBase, uint32(v90)+4)) = v84
											*(*int32)(unsafe.Add(mBase, uint32(v90))) = v87
											v95 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
											v97 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentVirtualXIDs[2]))
											v100 = v95
											v101 = v44 + int32(1)
											v102 = v97
										}
									} else {
										v73 = int32(3)
										if base.B2i32(base.Ui32(l0) < base.Ui32(v73))|base.B2i32(base.Ui32(v68) < base.Ui32(v73)) == int32(0) {
											if v68-l0 <= int32(0) {
												v84 = *(*int32)(unsafe.Add(mBase, uint32(v55)+44))
												if v84 == int32(0) {
													v100 = v41
													v101 = v44
													v102 = v46
												} else {
													v87 = *(*int32)(unsafe.Add(mBase, uint32(v55)+40))
													v90 = v17 + v44<<(uint(int32(3))%32)
													*(*int32)(unsafe.Add(mBase, uint32(v90)+4)) = v84
													*(*int32)(unsafe.Add(mBase, uint32(v90))) = v87
													v95 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
													v97 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentVirtualXIDs[2]))
													v100 = v95
													v101 = v44 + int32(1)
													v102 = v97
												}
											} else {
												v100 = v41
												v101 = v44
												v102 = v46
											}
										} else {
											if base.Ui32(l0) < base.Ui32(v68) {
												v100 = v41
												v101 = v44
												v102 = v46
											} else {
												v84 = *(*int32)(unsafe.Add(mBase, uint32(v55)+44))
												if v84 == int32(0) {
													v100 = v41
													v101 = v44
													v102 = v46
												} else {
													v87 = *(*int32)(unsafe.Add(mBase, uint32(v55)+40))
													v90 = v17 + v44<<(uint(int32(3))%32)
													*(*int32)(unsafe.Add(mBase, uint32(v90)+4)) = v84
													*(*int32)(unsafe.Add(mBase, uint32(v90))) = v87
													v95 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
													v97 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentVirtualXIDs[2]))
													v100 = v95
													v101 = v44 + int32(1)
													v102 = v97
												}
											}
										}
									}
								}
							}
						}
					}
					v104 = v43 + int32(1)
					if v104 < v100 {
						v41 = v100
						v43 = v104
						v44 = v101
						v46 = v102
						continue
					} else {
						break
					}
					break
				}
				v113 = v101
			} else {
				v113 = v3
			}
			v119 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentVirtualXIDs[1]))
			F_LWLockRelease(m, v119+int32(512))
			mBase = m.M
			v123 = m.ExcPending
			if v123 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v113
				return v17
			}
		}
	}
}
func F_SetCurrentRoleId(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	if l0 == int32(0) {
		v6 = int32(0)
		*(*uint8)(unsafe.Add(mBase, _c_F_SetCurrentRoleId[0])) = uint8(v6)
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_SetCurrentRoleId[1]))
		if v9 == v6 {
			return
		} else {
			v13 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SetCurrentRoleId[2])))
			v17 = v9
			v18 = v13
			*(*int32)(unsafe.Add(mBase, _c_F_SetCurrentRoleId[3])) = v17
			*(*int32)(unsafe.Add(mBase, _c_F_SetCurrentRoleId[4])) = v17
			if v18&int32(1) != 0 {
				v28 = int32(_a_F_SetCurrentRoleId_0)
			} else {
				v28 = int32(_a_F_SetCurrentRoleId_1)
			}
			F_SetConfigOption(m, int32(_a_F_SetCurrentRoleId_2), v28, int32(0), int32(1))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		v15 = int32(1)
		*(*uint8)(unsafe.Add(mBase, _c_F_SetCurrentRoleId[0])) = uint8(v15)
		v17 = l0
		v18 = l1
		*(*int32)(unsafe.Add(mBase, _c_F_SetCurrentRoleId[3])) = v17
		*(*int32)(unsafe.Add(mBase, _c_F_SetCurrentRoleId[4])) = v17
		if v18&int32(1) != 0 {
			v28 = int32(_a_F_SetCurrentRoleId_0)
		} else {
			v28 = int32(_a_F_SetCurrentRoleId_1)
		}
		F_SetConfigOption(m, int32(_a_F_SetCurrentRoleId_2), v28, int32(0), int32(1))
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return
		} else {
			return
		}
	}
}
func F_SetCurrentStatementStartTimestamp(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_SetCurrentStatementStartTimestamp[0]))
	if v2 < int32(0) {
		v9 = m.G0
		v10 = int32(16)
		v11 = v9 - v10
		m.G0 = v11
		F_gettimeofday(m, v11)
		mBase = m.M
		v14 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
		v15 = int64(*(*int32)(unsafe.Add(mBase, uint32(v11)+8)))
		m.G0 = v11 + v10
		*(*int64)(unsafe.Add(mBase, _c_F_SetCurrentStatementStartTimestamp[1])) = v15 + v14*int64(1000000) - int64(946684800000000)
	} else {
	}
	return
}
func F_current_query(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_current_query[0]))
	if v4 != 0 {
		v5 = F_cstring_to_text(m, v4)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v5)
		}
	} else {
		v11 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v11)
		return int64(0)
	}
}
func F_isCurrentGroup(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
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
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int64
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int64
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int64
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+96))
	v20 = v18 - int32(1)
	if int32(0) <= v20 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L11
	} else {
		goto L26
	}
L2:
	;
	m.G0 = v15 + int32(16)
	return v126
L3:
	;
	v27 = v20
	goto L6
L4:
	;
	goto L5
L5:
	;
	v126 = int32(1)
	goto L2
L6:
	;
	v36 = v27 * int32(36)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v39 = int32(*(*int16)(unsafe.Add(mBase, uint32(v36+v37)+32)))
	v40 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
	if v40 < v39 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L5
L8:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	m.T0[v43].(func(*base.Module, int32, int32))(m, l1, v39)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v49 = v39 - int32(1)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49+v50))))
	v54 = v49 << (uint(int32(3)) % 32)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v57 = *(*int64)(unsafe.Add(mBase, uint32(v54+v55)))
	v58 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+6)))
	if v58 < v39 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	return int32(0)
L12:
	;
	goto L10
L13:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
	m.T0[v61].(func(*base.Module, int32, int32))(m, l2, v39)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L11
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64+v49))))
	if v52|v66 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L15
L17:
	;
	if int32(0) < v27 {
		v27 = v27 - int32(1)
		goto L6
	} else {
		goto L25
	}
L18:
	;
	if v66 == v52 {
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v72 = *(*int64)(unsafe.Add(mBase, uint32(v70+v54)))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v74 = v73 + v36
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+28))
	*(*int64)(unsafe.Add(mBase, uint32(v75)+24)) = v57
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v74)+28))
	*(*int64)(unsafe.Add(mBase, uint32(v77)+40)) = v72
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v74)+28))
	v80 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v79)+16)) = uint8(v80)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v74)+28))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v85 = m.T0[v84].(func(*base.Module, int32) int64)(m, v82)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L11
	} else {
		goto L22
	}
L21:
	;
	v126 = int32(0)
	goto L2
L22:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v74)+28))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+16)))
	if v88 == int32(1) {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	if v85 != int64(0) {
		goto L17
	} else {
		goto L24
	}
L24:
	;
	v126 = int32(0)
	goto L2
L25:
	;
	goto L7
L26:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v135
	F_errmsg_internal(m, int32(_a_F_isCurrentGroup_0), v15)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L11
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_isCurrentGroup_1), int32(258), int32(_a_F_isCurrentGroup_2))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L11
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
