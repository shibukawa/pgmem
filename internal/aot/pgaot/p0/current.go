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
	var v11 int32
	_ = v11
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentReplayRecPtr[0]))
	v9 = base.AtomicRmwXchg32(m, v6, int32(96), int32(1))
	if v9 != 0 {
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentReplayRecPtr[0]))
		F_s_lock(m, v11+int32(96), int32(_a_F_GetCurrentReplayRecPtr_0), int32(_a_F_GetCurrentReplayRecPtr_1), int32(_a_F_GetCurrentReplayRecPtr_2))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int64(0)
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentReplayRecPtr[0]))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
			v24 = *(*int64)(unsafe.Add(mBase, uint32(v22)+48))
			v25 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v22)+96)), uint32(v25))
			if l0 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v23
			} else {
			}
			return v24
		}
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentReplayRecPtr[0]))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
		v24 = *(*int64)(unsafe.Add(mBase, uint32(v22)+48))
		v25 = int32(0)
		atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v22)+96)), uint32(v25))
		if l0 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v23
		} else {
		}
		return v24
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
								F_errfinish(m, int32(_a_F_GetCurrentTimeUsec_4), int32(432), int32(_a_F_GetCurrentTimeUsec_5))
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
							F_errfinish(m, int32(_a_F_GetCurrentTimeUsec_4), int32(432), int32(_a_F_GetCurrentTimeUsec_5))
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	v3 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentVirtualXIDs[0]))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v17 = F_palloc(m, v14<<(uint(int32(3))%32))
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
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
			if int32(0) < v28 {
				v34 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentVirtualXIDs[2]))
				v40 = v3
				v41 = v3
				v43 = v34
				for {
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(36)+v40<<(uint(int32(2))%32))))
					v52 = v43 + v49*int32(640)
					v54 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentVirtualXIDs[3]))
					if v52 == v54 {
						v100 = v41
						v101 = v43
					} else {
						v57 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentVirtualXIDs[4]))
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
						v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58+v40))))
						if v60&int32(7) != 0 {
							v100 = v41
							v101 = v43
						} else {
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v52)+60))
							v65 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentVirtualXIDs[5]))
							if v63 != v65 {
								v100 = v41
								v101 = v43
							} else {
								v67 = *(*int32)(unsafe.Add(mBase, uint32(v52)+40))
								if v67 == int32(0) {
									v100 = v41
									v101 = v43
								} else {
									if l0 != 0 {
										if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l0))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v67)) == int32(0) {
											v81 = base.B2i32(base.Ui32(v67) <= base.Ui32(l0))
										} else {
											v81 = base.B2i32(v67-l0 <= int32(0))
										}
										v83 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentVirtualXIDs[2]))
										if v81 == int32(0) {
											v100 = v41
											v101 = v83
										} else {
											v86 = v83
											v87 = *(*int32)(unsafe.Add(mBase, uint32(v52)+56))
											if v87 == int32(0) {
												v100 = v41
												v101 = v86
											} else {
												v90 = *(*int32)(unsafe.Add(mBase, uint32(v52)+52))
												v93 = v17 + v41<<(uint(int32(3))%32)
												*(*int32)(unsafe.Add(mBase, uint32(v93)+4)) = v87
												*(*int32)(unsafe.Add(mBase, uint32(v93))) = v90
												v100 = v41 + int32(1)
												v101 = v86
											}
										}
									} else {
										v86 = v43
										v87 = *(*int32)(unsafe.Add(mBase, uint32(v52)+56))
										if v87 == int32(0) {
											v100 = v41
											v101 = v86
										} else {
											v90 = *(*int32)(unsafe.Add(mBase, uint32(v52)+52))
											v93 = v17 + v41<<(uint(int32(3))%32)
											*(*int32)(unsafe.Add(mBase, uint32(v93)+4)) = v87
											*(*int32)(unsafe.Add(mBase, uint32(v93))) = v90
											v100 = v41 + int32(1)
											v101 = v86
										}
									}
								}
							}
						}
					}
					v104 = v40 + int32(1)
					v105 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					if v104 < v105 {
						v40 = v104
						v41 = v100
						v43 = v101
						continue
					} else {
						break
					}
					break
				}
				v113 = v100
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
func F_current_query(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_current_query[0]))
	if v4 != 0 {
		v5 = F_cstring_to_text(m, v4)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			return v5
		}
	} else {
		v10 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v10)
		return int32(0)
	}
}
func F_isCurrentGroup(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
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
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+96))
	v19 = v17 - int32(1)
	if int32(0) <= v19 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L11
	} else {
		goto L26
	}
L2:
	;
	m.G0 = v14 + int32(16)
	return v115
L3:
	;
	v27 = v19
	goto L6
L4:
	;
	goto L5
L5:
	;
	v115 = int32(1)
	goto L2
L6:
	;
	v34 = v27 * int32(36)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v37 = int32(*(*int16)(unsafe.Add(mBase, uint32(v34+v35)+32)))
	v38 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
	if v38 < v37 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L5
L8:
	;
	F_slot_getsomeattrs_int(m, l1, v37)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v45 = v37 - int32(1)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+v46))))
	v50 = v45 << (uint(int32(2)) % 32)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v50+v51)))
	v54 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+6)))
	if v54 < v37 {
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
	F_slot_getsomeattrs_int(m, l2, v37)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L11
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58+v45))))
	if v48|v60 != 0 {
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
	if v60 == v48 {
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v64+v50)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v68 = v67 + v34
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v69)+20)) = v53
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v68)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v71)+28)) = v66
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v68)+28))
	v74 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v73)+16)) = uint8(v74)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v68)+28))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v79 = m.T0[v78].(func(*base.Module, int32) int32)(m, v76)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L11
	} else {
		goto L22
	}
L21:
	;
	v115 = int32(0)
	goto L2
L22:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v68)+28))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+16)))
	if v82 == int32(1) {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	if v79 != 0 {
		goto L17
	} else {
		goto L24
	}
L24:
	;
	v115 = int32(0)
	goto L2
L25:
	;
	goto L7
L26:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v124
	F_errmsg_internal(m, int32(_a_F_isCurrentGroup_0), v14)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L11
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_isCurrentGroup_1), int32(258), int32(_a_F_isCurrentGroup_2))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
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
