package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_GetAccessStrategyWithSize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	if base.Ui32(int32(15)) <= base.Ui32(l0+int32(7)) {
		v8 = *(*int32)(unsafe.Add(mBase, _c_F_GetAccessStrategyWithSize[0]))
		v9 = int32(8)
		v10 = base.I32_div_s(v8, v9)
		v12 = base.I32_div_s(l0, v9)
		if v10 < v12 {
			v14 = v10
		} else {
			v14 = v12
		}
		v19 = F_palloc0(m, v14<<(uint(int32(2))%32)+int32(12))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v14
			*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(3)
			v27 = v19
			return v27
		}
	} else {
		v27 = int32(0)
		return v27
	}
}
func F_GetComboCommandId(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
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
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_GetComboCommandId[0]))
	if v12 == int32(0) {
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_GetComboCommandId[1]))
		v18 = F_MemoryContextAlloc(m, v16, int32(800))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v23 = int32(100)
			*(*int32)(unsafe.Add(mBase, _c_F_GetComboCommandId[2])) = v23
			*(*int32)(unsafe.Add(mBase, _c_F_GetComboCommandId[3])) = v18
			*(*int32)(unsafe.Add(mBase, _c_F_GetComboCommandId[4])) = int32(0)
			*(*int64)(unsafe.Add(mBase, uint32(v9)+28)) = int64(51539607560)
			v33 = *(*int32)(unsafe.Add(mBase, _c_F_GetComboCommandId[1]))
			*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = v33
			v41 = F_hash_create(m, int32(_a_F_GetComboCommandId_0), v23, v7+int32(-52), int32(1064))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_GetComboCommandId[0])) = v41
				v44 = v41
				v46 = *(*int32)(unsafe.Add(mBase, _c_F_GetComboCommandId[2]))
				v48 = *(*int32)(unsafe.Add(mBase, _c_F_GetComboCommandId[4]))
				if v46 <= v48 {
					v51 = *(*int32)(unsafe.Add(mBase, _c_F_GetComboCommandId[3]))
					v54 = F_repalloc(m, v51, v46<<(uint(int32(4))%32))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_GetComboCommandId[2])) = v46 << (uint(int32(1)) % 32)
						*(*int32)(unsafe.Add(mBase, _c_F_GetComboCommandId[3])) = v54
						v63 = *(*int32)(unsafe.Add(mBase, _c_F_GetComboCommandId[0]))
						v64 = v63
						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l0
						v72 = F_hash_search(m, v64, v7+int32(-52), int32(1), v7+int32(-1))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int32(0)
						} else {
							v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+63)))
							if v74 == int32(1) {
								v77 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
								v92 = v77
							} else {
								v79 = *(*int32)(unsafe.Add(mBase, _c_F_GetComboCommandId[3]))
								v80 = int32(_a_F_GetComboCommandId_1)
								v81 = *(*int32)(unsafe.Add(mBase, _c_F_GetComboCommandId[4]))
								v84 = v79 + v81<<(uint(int32(3))%32)
								*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = l1
								*(*int32)(unsafe.Add(mBase, uint32(v84))) = l0
								*(*int32)(unsafe.Add(mBase, _c_F_GetComboCommandId[4])) = v81 + int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = v81
								v92 = v81
							}
							m.G0 = v9 - int32(-64)
							return v92
						}
					}
				} else {
					v64 = v44
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l0
					v72 = F_hash_search(m, v64, v7+int32(-52), int32(1), v7+int32(-1))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+63)))
						if v74 == int32(1) {
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
							v92 = v77
						} else {
							v79 = *(*int32)(unsafe.Add(mBase, _c_F_GetComboCommandId[3]))
							v80 = int32(_a_F_GetComboCommandId_1)
							v81 = *(*int32)(unsafe.Add(mBase, _c_F_GetComboCommandId[4]))
							v84 = v79 + v81<<(uint(int32(3))%32)
							*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v84))) = l0
							*(*int32)(unsafe.Add(mBase, _c_F_GetComboCommandId[4])) = v81 + int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = v81
							v92 = v81
						}
						m.G0 = v9 - int32(-64)
						return v92
					}
				}
			}
		}
	} else {
		v44 = v12
		v46 = *(*int32)(unsafe.Add(mBase, _c_F_GetComboCommandId[2]))
		v48 = *(*int32)(unsafe.Add(mBase, _c_F_GetComboCommandId[4]))
		if v46 <= v48 {
			v51 = *(*int32)(unsafe.Add(mBase, _c_F_GetComboCommandId[3]))
			v54 = F_repalloc(m, v51, v46<<(uint(int32(4))%32))
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_GetComboCommandId[2])) = v46 << (uint(int32(1)) % 32)
				*(*int32)(unsafe.Add(mBase, _c_F_GetComboCommandId[3])) = v54
				v63 = *(*int32)(unsafe.Add(mBase, _c_F_GetComboCommandId[0]))
				v64 = v63
				*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l0
				v72 = F_hash_search(m, v64, v7+int32(-52), int32(1), v7+int32(-1))
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return int32(0)
				} else {
					v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+63)))
					if v74 == int32(1) {
						v77 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
						v92 = v77
					} else {
						v79 = *(*int32)(unsafe.Add(mBase, _c_F_GetComboCommandId[3]))
						v80 = int32(_a_F_GetComboCommandId_1)
						v81 = *(*int32)(unsafe.Add(mBase, _c_F_GetComboCommandId[4]))
						v84 = v79 + v81<<(uint(int32(3))%32)
						*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v84))) = l0
						*(*int32)(unsafe.Add(mBase, _c_F_GetComboCommandId[4])) = v81 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = v81
						v92 = v81
					}
					m.G0 = v9 - int32(-64)
					return v92
				}
			}
		} else {
			v64 = v44
			*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l0
			v72 = F_hash_search(m, v64, v7+int32(-52), int32(1), v7+int32(-1))
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return int32(0)
			} else {
				v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+63)))
				if v74 == int32(1) {
					v77 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
					v92 = v77
				} else {
					v79 = *(*int32)(unsafe.Add(mBase, _c_F_GetComboCommandId[3]))
					v80 = int32(_a_F_GetComboCommandId_1)
					v81 = *(*int32)(unsafe.Add(mBase, _c_F_GetComboCommandId[4]))
					v84 = v79 + v81<<(uint(int32(3))%32)
					*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v84))) = l0
					*(*int32)(unsafe.Add(mBase, _c_F_GetComboCommandId[4])) = v81 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = v81
					v92 = v81
				}
				m.G0 = v9 - int32(-64)
				return v92
			}
		}
	}
}
func F_GetCompressionMethodName(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	switch l0 - int32(108) {
	case 0:
		v27 = int32(_a_F_GetCompressionMethodName_0)
		m.G0 = v6 + int32(16)
		return v27
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
			F_errmsg_internal(m, int32(_a_F_GetCompressionMethodName_1), v6)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_GetCompressionMethodName_2), int32(313), int32(_a_F_GetCompressionMethodName_3))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 4:
		v27 = int32(_a_F_GetCompressionMethodName_4)
		m.G0 = v6 + int32(16)
		return v27
	}
}
func F_GetFdwRoutineByServerId(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v10 = F_SearchSysCache1(m, int32(32), l0)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 != 0 {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+22)))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v14+v15)+72))
			F_ReleaseCatCache(m, v10)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v21 = F_SearchSysCache1(m, int32(30), v17)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					if v21 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v17
							F_errmsg_internal(m, int32(_a_F_GetFdwRoutineByServerId_0), v7+int32(16))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_GetFdwRoutineByServerId_1), int32(397), int32(_a_F_GetFdwRoutineByServerId_2))
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
						v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+22)))
						v27 = v25 + v26
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+72))
						if v28 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(325))
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v27 + int32(4)
									F_errmsg(m, int32(_a_F_GetFdwRoutineByServerId_3), v7+int32(32))
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_GetFdwRoutineByServerId_1), int32(406), int32(_a_F_GetFdwRoutineByServerId_2))
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							F_ReleaseCatCache(m, v21)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int32(0)
							} else {
								v33 = F_GetFdwRoutine(m, v28)
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
									return int32(0)
								} else {
									m.G0 = v7 + int32(48)
									return v33
								}
							}
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
				F_errmsg_internal(m, int32(_a_F_GetFdwRoutineByServerId_4), v7)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_GetFdwRoutineByServerId_1), int32(389), int32(_a_F_GetFdwRoutineByServerId_2))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
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
func F_GetLatestXTime(m *base.Module) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_GetLatestXTime[0]))
	v7 = base.AtomicRmwXchg32(m, v4, int32(96), int32(1))
	if v7 != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_GetLatestXTime[0]))
		F_s_lock(m, v9+int32(96), int32(_a_F_GetLatestXTime_0), int32(_a_F_GetLatestXTime_1), int32(_a_F_GetLatestXTime_2))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, _c_F_GetLatestXTime[0]))
			v21 = *(*int64)(unsafe.Add(mBase, uint32(v20)+64))
			v22 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v20)+96)), uint32(v22))
			return v21
		}
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_GetLatestXTime[0]))
		v21 = *(*int64)(unsafe.Add(mBase, uint32(v20)+64))
		v22 = int32(0)
		atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v20)+96)), uint32(v22))
		return v21
	}
}
func F_GetNamedDSMSegment(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	v7 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetNamedDSMSegment[0])))
	if v7 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L11
	} else {
		goto L62
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L11
	} else {
		goto L59
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L11
	} else {
		goto L56
	}
L4:
	;
	v9 = F_strlen(m, int32(_a_F_GetNamedDSMSegment_0))
	mBase = m.M
	if base.Ui32(int32(64)) <= base.Ui32(v9) {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L11
	} else {
		goto L53
	}
L7:
	;
	v12 = int32(_a_F_GetNamedDSMSegment_1)
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_GetNamedDSMSegment[1]))
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_GetNamedDSMSegment[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_GetNamedDSMSegment[1])) = v16
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_GetNamedDSMSegment[3]))
	if v19 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v97 = v19
	goto L10
L9:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_GetNamedDSMSegment[4]))
	v25 = F_LWLockAcquire(m, v21+int32(_a_F_GetNamedDSMSegment_2), int32(0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v99 = F_dshash_find_or_insert(m, v97, int32(_a_F_GetNamedDSMSegment_0), l0)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L11
	} else {
		goto L27
	}
L11:
	;
	return int32(0)
L12:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_GetNamedDSMSegment[5]))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v31 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_GetNamedDSMSegment[4]))
	F_LWLockRelease(m, v89+int32(_a_F_GetNamedDSMSegment_2))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L11
	} else {
		goto L26
	}
L14:
	;
	v38 = F_dsa_create_ext(m, int32(84), int32(_a_F_GetNamedDSMSegment_3), int32(134217728))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L11
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v71 = F_dsa_attach(m, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L11
	} else {
		goto L23
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetNamedDSMSegment[6])) = v38
	v44 = F_dshash_create(m, v38, int32(_a_F_GetNamedDSMSegment_4), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetNamedDSMSegment[3])) = v44
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_GetNamedDSMSegment[6]))
	F_dsa_pin(m, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L11
	} else {
		goto L19
	}
L19:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_GetNamedDSMSegment[6]))
	F_dsa_pin_mapping(m, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L11
	} else {
		goto L20
	}
L20:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_GetNamedDSMSegment[6]))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+28))
	goto L21
L21:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_GetNamedDSMSegment[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = v58
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_GetNamedDSMSegment[3]))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+32))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	goto L22
L22:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_GetNamedDSMSegment[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v67)+4)) = v65
	goto L13
L23:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetNamedDSMSegment[6])) = v71
	F_dsa_pin_mapping(m, v71)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L11
	} else {
		goto L24
	}
L24:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_GetNamedDSMSegment[6]))
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_GetNamedDSMSegment[5]))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v84 = F_dshash_attach(m, v78, int32(_a_F_GetNamedDSMSegment_4), v82, int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L11
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetNamedDSMSegment[3])) = v84
	goto L13
L26:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_GetNamedDSMSegment[3]))
	v97 = v95
	goto L10
L27:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v101 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v172)+24))
	v179 = *(*int32)(unsafe.Add(mBase, _c_F_GetNamedDSMSegment[3]))
	F_dshash_release_lock(m, v179, v99)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L11
	} else {
		goto L52
	}
L29:
	;
	v135 = *(*int32)(unsafe.Add(mBase, _c_F_GetNamedDSMSegment[7]))
	v136 = int32(0)
	if base.B2i32(v135 == v136)|base.B2i32(v135 == int32(_a_F_GetNamedDSMSegment_5)) == v136 {
		goto L41
	} else {
		goto L42
	}
L30:
	;
	v118 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v118)
	v122 = F_dsm_create(m, int32(44), v118)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L11
	} else {
		goto L36
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+68)) = int32(44)
	*(*int32)(unsafe.Add(mBase, uint32(v99)+64)) = int32(0)
	v116 = v99 - int32(-64)
	goto L30
L32:
	;
	goto L33
L33:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v99)+68))
	if v110 != int32(44) {
		goto L2
	} else {
		goto L34
	}
L34:
	;
	v114 = v99 - int32(-64)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v99)+64))
	if v115 != 0 {
		goto L29
	} else {
		goto L35
	}
L35:
	;
	v116 = v114
	goto L30
L36:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v122)+24))
	m.T0[int32(_a_F_GetNamedDSMSegment_6)].(func(*base.Module, int32))(m, v124)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L11
	} else {
		goto L37
	}
L37:
	;
	F_dsm_pin_segment(m, v122)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L11
	} else {
		goto L38
	}
L38:
	;
	F_dsm_pin_mapping(m, v122)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L11
	} else {
		goto L39
	}
L39:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = v132
	v172 = v122
	goto L28
L40:
	;
	if v164 != 0 {
		v172 = v164
		goto L28
	} else {
		goto L48
	}
L41:
	;
	v143 = v135
	goto L44
L42:
	;
	goto L43
L43:
	;
	v164 = int32(0)
	goto L40
L44:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v143)+12))
	if v115 == v148 {
		v164 = v143
		goto L40
	} else {
		goto L46
	}
L45:
	;
	goto L43
L46:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	if v150 != int32(_a_F_GetNamedDSMSegment_5) {
		v143 = v150
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	v166 = F_dsm_attach(m, v165)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L11
	} else {
		goto L49
	}
L49:
	;
	if v166 == int32(0) {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_dsm_pin_mapping(m, v166)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L11
	} else {
		goto L51
	}
L51:
	;
	v172 = v166
	goto L28
L52:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetNamedDSMSegment[1])) = v13
	return v177
L53:
	;
	F_errmsg(m, int32(_a_F_GetNamedDSMSegment_7), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L11
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_GetNamedDSMSegment_8), int32(144), int32(_a_F_GetNamedDSMSegment_9))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L11
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	F_errmsg(m, int32(_a_F_GetNamedDSMSegment_10), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L11
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_GetNamedDSMSegment_8), int32(148), int32(_a_F_GetNamedDSMSegment_9))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L11
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	F_errmsg(m, int32(_a_F_GetNamedDSMSegment_11), int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L11
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_GetNamedDSMSegment_8), int32(168), int32(_a_F_GetNamedDSMSegment_9))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L11
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	F_errmsg_internal(m, int32(_a_F_GetNamedDSMSegment_12), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L11
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_GetNamedDSMSegment_8), int32(192), int32(_a_F_GetNamedDSMSegment_9))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L11
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_GetOldestNonRemovableTransactionId(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	v3 = m.G0
	v5 = v3 - int32(48)
	m.G0 = v5
	F_ComputeXidHorizons(m, v5+int32(8))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = int32(0)
		if l0 == v13 {
			v44 = v13
			v48 = v44
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+117)))
			if v18 != 0 {
				v44 = v13
				v48 = v44
			} else {
				v19 = F_RecoveryInProgress(m)
				mBase = m.M
				if v19 != 0 {
					v44 = v13
					v48 = v44
				} else {
					v20 = int32(1)
					v21 = F_IsCatalogRelation(m, l0)
					mBase = m.M
					if v21 != 0 {
						v44 = v20
						v48 = v44
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestNonRemovableTransactionId[0]))
						if v23 < int32(2) {
							v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
							if v40 != 0 {
								v44 = int32(3)
								v48 = v44
							} else {
								v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								if v41 != 0 {
									v44 = int32(3)
									v48 = v44
								} else {
									v48 = int32(2)
								}
							}
						} else {
							v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+118)))
							if v27 != int32(112) {
								v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
								if v40 != 0 {
									v44 = int32(3)
									v48 = v44
								} else {
									v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
									if v41 != 0 {
										v44 = int32(3)
										v48 = v44
									} else {
										v48 = int32(2)
									}
								}
							} else {
								v30 = F_IsCatalogRelation(m, l0)
								mBase = m.M
								if v30 != 0 {
									v44 = v20
									v48 = v44
								} else {
									v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
									if v31 == int32(0) {
										v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
										if v40 != 0 {
											v44 = int32(3)
											v48 = v44
										} else {
											v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
											if v41 != 0 {
												v44 = int32(3)
												v48 = v44
											} else {
												v48 = int32(2)
											}
										}
									} else {
										v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
										v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+119)))
										switch v35 - int32(109) {
										case 0, 5:
											v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+104)))
											if v38 != 0 {
												v44 = v20
												v48 = v44
											} else {
												v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
												if v40 != 0 {
													v44 = int32(3)
													v48 = v44
												} else {
													v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
													if v41 != 0 {
														v44 = int32(3)
														v48 = v44
													} else {
														v48 = int32(2)
													}
												}
											}
										default:
											v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
											if v40 != 0 {
												v44 = int32(3)
												v48 = v44
											} else {
												v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
												if v41 != 0 {
													v44 = int32(3)
													v48 = v44
												} else {
													v48 = int32(2)
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
		switch v48 - int32(1) {
		case 0:
			v52 = *(*int32)(unsafe.Add(mBase, uint32(v5)+36))
			v55 = v52
		case 1:
			v53 = *(*int32)(unsafe.Add(mBase, uint32(v5)+40))
			v55 = v53
		case 2:
			v54 = *(*int32)(unsafe.Add(mBase, uint32(v5)+44))
			v55 = v54
		default:
			v51 = *(*int32)(unsafe.Add(mBase, uint32(v5)+28))
			v55 = v51
		}
		m.G0 = v5 + int32(48)
		return v55
	}
}
func F_GetOldestSafeDecodingTransactionId(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestSafeDecodingTransactionId[0]))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetOldestSafeDecodingTransactionId[1])))
	if v11 == int32(1) {
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestSafeDecodingTransactionId[2]))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+316))
		v19 = base.B2i32(v17 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _c_F_GetOldestSafeDecodingTransactionId[1])) = uint8(v19)
		v21 = v19
	} else {
		v21 = int32(0)
	}
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestSafeDecodingTransactionId[3]))
	v27 = F_LWLockAcquire(m, v23+int32(384), int32(1))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		return int32(0)
	} else {
		v32 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestSafeDecodingTransactionId[4]))
		v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
		v35 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestSafeDecodingTransactionId[0]))
		v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+28))
		if v36 == int32(0) {
			v56 = v33
		} else {
			if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v33))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v36)) == int32(0) {
				v50 = base.B2i32(base.Ui32(v36) < base.Ui32(v33))
			} else {
				v50 = int32(base.Ui32(v36-v33) >> (uint(int32(31)) % 32))
			}
			if v50 == int32(0) {
				v56 = v33
			} else {
				v54 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestSafeDecodingTransactionId[0]))
				v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+28))
				v56 = v55
			}
		}
		if l0 == int32(0) {
			v82 = v56
		} else {
			v60 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestSafeDecodingTransactionId[0]))
			v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+32))
			if v61 == int32(0) {
				v82 = v56
			} else {
				if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v56))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v61)) == int32(0) {
					v75 = base.B2i32(base.Ui32(v61) < base.Ui32(v56))
				} else {
					v75 = int32(base.Ui32(v61-v56) >> (uint(int32(31)) % 32))
				}
				if v75 == int32(0) {
					v82 = v56
				} else {
					v79 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestSafeDecodingTransactionId[0]))
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+32))
					v82 = v80
				}
			}
		}
		if v21 != 0 {
			v122 = v82
		} else {
			v83 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
			if v83 <= int32(0) {
				v122 = v82
			} else {
				v88 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestSafeDecodingTransactionId[5]))
				v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
				v90 = int32(0)
				v91 = v82
				v92 = v83
				for {
					v99 = *(*int32)(unsafe.Add(mBase, uint32(v89+v90<<(uint(int32(2))%32))))
					if base.Ui32(int32(3)) <= base.Ui32(v99) {
						if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v91))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v99)) == int32(0) {
							v113 = base.B2i32(base.Ui32(v99) < base.Ui32(v91))
						} else {
							v113 = int32(base.Ui32(v99-v91) >> (uint(int32(31)) % 32))
						}
						if v113 != 0 {
							v114 = v99
						} else {
							v114 = v91
						}
						v115 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
						v116 = v114
						v117 = v115
					} else {
						v116 = v91
						v117 = v92
					}
					v119 = v90 + int32(1)
					if v119 < v117 {
						v90 = v119
						v91 = v116
						v92 = v117
						continue
					} else {
						break
					}
					break
				}
				v122 = v116
			}
		}
		v128 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestSafeDecodingTransactionId[3]))
		F_LWLockRelease(m, v128+int32(384))
		mBase = m.M
		v132 = m.ExcPending
		if v132 != 0 {
			return int32(0)
		} else {
			return v122
		}
	}
}
func F_GetOldestUnsummarizedLSN(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int64
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int64
	_ = v64
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
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int64
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int64
	_ = v106
	var v109 int32
	_ = v109
	var v111 int64
	_ = v111
	var v112 int64
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v149 int64
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int64
	_ = v156
	var v157 int32
	_ = v157
	var v158 int64
	_ = v158
	var v159 int32
	_ = v159
	var v160 int64
	_ = v160
	var v161 int32
	_ = v161
	var v162 int64
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v186 int64
	_ = v186
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int64
	_ = v194
	var v195 int32
	_ = v195
	var v196 int64
	_ = v196
	var v204 int32
	_ = v204
	var v212 int64
	_ = v212
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v237 int64
	_ = v237
	var v238 int64
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v263 int64
	_ = v263
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	v3 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v23 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetOldestUnsummarizedLSN[0])))
	if v23 != int32(1) {
		v263 = int64(0)
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L7
	} else {
		goto L70
	}
L2:
	;
	m.G0 = v20 + int32(16)
	return v263
L3:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestUnsummarizedLSN[1]))
	if v27 != int32(15) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestUnsummarizedLSN[2]))
	v35 = F_LWLockAcquire(m, v31+int32(_a_F_GetOldestUnsummarizedLSN_0), int32(1))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v64 = F_GetLatestLSN(m, v20+int32(12))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L7
	} else {
		goto L20
	}
L7:
	;
	return int64(0)
L8:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestUnsummarizedLSN[3]))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	if v41 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v40)+8))
	if l0 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestUnsummarizedLSN[2]))
	F_LWLockRelease(m, v56+int32(_a_F_GetOldestUnsummarizedLSN_0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L7
	} else {
		goto L19
	}
L12:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v45
	goto L14
L13:
	;
	goto L14
L14:
	;
	if l1 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v47)
	goto L17
L16:
	;
	goto L17
L17:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestUnsummarizedLSN[2]))
	F_LWLockRelease(m, v50+int32(_a_F_GetOldestUnsummarizedLSN_0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L7
	} else {
		goto L18
	}
L18:
	;
	v263 = v44
	goto L2
L19:
	;
	goto L6
L20:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v67 = F_readTimeLineHistory(m, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	if v67 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	v71 = v69
	goto L24
L23:
	;
	v71 = int32(0)
	goto L24
L24:
	;
	v74 = v71
	goto L26
L25:
	;
	v112 = int64(0)
	v114 = F_GetWalSummaries(m, v109, v112, v112)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L7
	} else {
		goto L32
	}
L26:
	;
	v91 = v74 - int32(1)
	if v91 < int32(0) {
		v109 = v3
		v111 = int64(0)
		goto L25
	} else {
		goto L28
	}
L27:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v106 = int64(*(*int32)(unsafe.Add(mBase, _c_F_GetOldestUnsummarizedLSN[4])))
	v109 = v104
	v111 = v100 * v106
	goto L25
L28:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v94+v91<<(uint(int32(2))%32))))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v100 = F_XLogGetOldestSegno(m, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L7
	} else {
		goto L29
	}
L29:
	;
	if v100 == int64(0) {
		v74 = v91
		goto L26
	} else {
		goto L30
	}
L30:
	;
	goto L27
L31:
	;
	if v109 == int32(0) {
		goto L1
	} else {
		goto L55
	}
L32:
	;
	if v114 == int32(0) {
		v204 = v3
		v212 = v111
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	if v118 <= int32(0) {
		v204 = v3
		v212 = v111
		goto L31
	} else {
		goto L34
	}
L34:
	;
	if v118 == int32(1) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v189+v174<<(uint(int32(2))%32))))
	v194 = *(*int64)(unsafe.Add(mBase, uint32(v193)+8))
	v195 = base.B2i32(base.Ui64(v186) < base.Ui64(v194))
	if base.Ui64(v186) < base.Ui64(v194) {
		goto L52
	} else {
		goto L53
	}
L36:
	;
	v174 = int32(0)
	v178 = v3
	v186 = v111
	goto L35
L37:
	;
	goto L38
L38:
	;
	v124 = int32(0)
	if v124 < v118 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v127 = v118
	goto L41
L40:
	;
	v127 = v124
	goto L41
L41:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	v133 = int32(0)
	v137 = v133
	v140 = v133
	v141 = v3
	v149 = v111
	goto L42
L42:
	;
	v154 = v132 + v137<<(uint(int32(2))%32)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)+4))
	v156 = *(*int64)(unsafe.Add(mBase, uint32(v155)+8))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	v158 = *(*int64)(unsafe.Add(mBase, uint32(v157)+8))
	v159 = base.B2i32(base.Ui64(v149) < base.Ui64(v158))
	if base.Ui64(v149) < base.Ui64(v158) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	if v127&int32(1) == int32(0) {
		v204 = v164
		v212 = v162
		goto L31
	} else {
		goto L51
	}
L44:
	;
	v160 = v158
	goto L46
L45:
	;
	v160 = v149
	goto L46
L46:
	;
	v161 = base.B2i32(base.Ui64(v160) < base.Ui64(v156))
	if base.Ui64(v160) < base.Ui64(v156) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v162 = v156
	goto L49
L48:
	;
	v162 = v160
	goto L49
L49:
	;
	v164 = v159 | v161 | v141
	v165 = int32(2)
	v166 = v137 + v165
	v168 = v140 + v165
	if v168 != v127&int32(2147483646) {
		v137 = v166
		v140 = v168
		v141 = v164
		v149 = v162
		goto L42
	} else {
		goto L50
	}
L50:
	;
	goto L43
L51:
	;
	v174 = v166
	v178 = v164
	v186 = v162
	goto L35
L52:
	;
	v196 = v194
	goto L54
L53:
	;
	v196 = v186
	goto L54
L54:
	;
	v204 = v195 | v178
	v212 = v196
	goto L31
L55:
	;
	v218 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestUnsummarizedLSN[2]))
	v222 = F_LWLockAcquire(m, v218+int32(_a_F_GetOldestUnsummarizedLSN_0), int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L7
	} else {
		goto L56
	}
L56:
	;
	v225 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestUnsummarizedLSN[3]))
	if v27 != int32(15) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	if l0 != 0 {
		goto L63
	} else {
		goto L64
	}
L58:
	;
	v237 = *(*int64)(unsafe.Add(mBase, uint32(v225)+8))
	v238 = v237
	goto L57
L59:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225))))
	if v228 != 0 {
		goto L58
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v225)+8)) = v212
	v230 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v225))) = uint8(v230)
	*(*int64)(unsafe.Add(mBase, uint32(v225)+24)) = v212
	*(*int32)(unsafe.Add(mBase, uint32(v225)+4)) = v109
	v235 = v204 & v230
	*(*uint8)(unsafe.Add(mBase, uint32(v225)+16)) = uint8(v235)
	v238 = v212
	goto L57
L62:
	;
	goto L61
L63:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v225)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v239
	goto L65
L64:
	;
	goto L65
L65:
	;
	if l1 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v241)
	goto L68
L67:
	;
	goto L68
L68:
	;
	v244 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestUnsummarizedLSN[2]))
	F_LWLockRelease(m, v244+int32(_a_F_GetOldestUnsummarizedLSN_0))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L7
	} else {
		goto L69
	}
L69:
	;
	v263 = v238
	goto L2
L70:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L7
	} else {
		goto L71
	}
L71:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v277
	F_errmsg_internal(m, int32(_a_F_GetOldestUnsummarizedLSN_1), v20)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L7
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_GetOldestUnsummarizedLSN_2), int32(599), int32(_a_F_GetOldestUnsummarizedLSN_3))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L7
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_GetSerializableTransactionSnapshotInt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v111 int64
	_ = v111
	var v112 int64
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int64
	_ = v148
	var v160 int64
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int64
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int64
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int64
	_ = v231
	var v232 int64
	_ = v232
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v258 int64
	_ = v258
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v271 int64
	_ = v271
	var v273 int64
	_ = v273
	var v274 int64
	_ = v274
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v377 int32
	_ = v377
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v508 int64
	_ = v508
	var v510 int32
	_ = v510
	var v511 int64
	_ = v511
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v555 int32
	_ = v555
	var v560 int32
	_ = v560
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v592 int32
	_ = v592
	var v597 int32
	_ = v597
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v640 int32
	_ = v640
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v676 int32
	_ = v676
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v696 int32
	_ = v696
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v727 int32
	_ = v727
	var v732 int32
	_ = v732
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v772 int32
	_ = v772
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v788 int32
	_ = v788
	var v791 int32
	_ = v791
	var v795 int32
	_ = v795
	var v801 int32
	_ = v801
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v842 int32
	_ = v842
	var v845 int32
	_ = v845
	var v849 int32
	_ = v849
	var v853 int32
	_ = v853
	var v857 int32
	_ = v857
	var v860 int32
	_ = v860
	var v864 int32
	_ = v864
	var v868 int32
	_ = v868
	var v873 int32
	_ = v873
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v884 int32
	_ = v884
	var v888 int32
	_ = v888
	var v893 int32
	_ = v893
	v18 = m.G0
	v20 = v18 + int32(-64)
	m.G0 = v20
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[0]))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+72))
	if v25 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v431)))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v431)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v445)+4)) = v446
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v431)))
	*(*int32)(unsafe.Add(mBase, uint32(v446))) = v448
	v451 = v432 + int32(8)
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v432)+12))
	if v452 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L2:
	;
	if v28&int32(1) == int32(0) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	v28 = int32(1)
	goto L5
L4:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+76)))
	v28 = v27
	goto L5
L5:
	;
	goto L2
L6:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[1]))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+56))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+52))
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[2]))
	v42 = F_LWLockAcquire(m, v38+int32(3584), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L9
	} else {
		goto L85
	}
L9:
	;
	return int32(0)
L10:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[3]))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	if v48 != v47 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v51 = v48
	goto L13
L12:
	;
	v51 = int32(0)
	goto L13
L13:
	;
	if v51 != 0 {
		v431 = v48
		v432 = v47
		goto L1
	} else {
		goto L14
	}
L14:
	;
	goto L15
L15:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[2]))
	F_LWLockRelease(m, v70+int32(3584))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L9
	} else {
		goto L17
	}
L16:
	;
	v431 = v410
	v432 = v409
	goto L1
L17:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[2]))
	v80 = F_LWLockAcquire(m, v76+int32(3712), int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[4]))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	v85 = int32(0)
	if base.B2i32(v84 == v85)|base.B2i32(v84 == v83) == v85 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+4)) = v92
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	*(*int32)(unsafe.Add(mBase, uint32(v92))) = v94
	*(*int64)(unsafe.Add(mBase, uint32(v84))) = int64(0)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v84)+40))
	if v100 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v396 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[2]))
	F_LWLockRelease(m, v396+int32(3712))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L9
	} else {
		goto L82
	}
L22:
	;
	F_ReleaseOneSerializableXact(m, v84-int32(56), int32(0), int32(1))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L9
	} else {
		goto L81
	}
L23:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v84)+52))
	if v103&int32(32) != 0 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	if v103&int32(16) != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v84-int32(32))))
	v112 = v111
	goto L27
L26:
	;
	v112 = int64(-1)
	goto L27
L27:
	;
	v114 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[5]))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+28))
	v117 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[6])))
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[2]))
	v123 = F_LWLockAcquire(m, v119+int32(_a_F_GetSerializableTransactionSnapshotInt_0), int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	v126 = int32(base.Ui32(v100) >> (uint(int32(10)) % 32))
	v127 = base.I32_rem_u_s(v126, v117)
	v129 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[7]))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+12))
	if v130 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v352 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[2]))
	F_LWLockRelease(m, v352+int32(_a_F_GetSerializableTransactionSnapshotInt_0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L9
	} else {
		goto L80
	}
L30:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v130))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v100)) == int32(0) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	if v144 != 0 {
		goto L29
	} else {
		goto L35
	}
L32:
	;
	v144 = base.B2i32(base.Ui32(v100) < base.Ui32(v130))
	goto L31
L33:
	;
	goto L34
L34:
	;
	v144 = int32(base.Ui32(v100-v130) >> (uint(int32(31)) % 32))
	goto L31
L35:
	;
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[7]))
	v148 = *(*int64)(unsafe.Add(mBase, uint32(v147)))
	if v148 < int64(0) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v203 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[7]))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)+8))
	if v204 != 0 {
		goto L53
	} else {
		goto L54
	}
L37:
	;
	v198 = int32(1)
	v201 = base.I64_extend_i32_u(int32(base.Ui32(v130) >> (uint(int32(10)) % 32)))
	goto L36
L38:
	;
	goto L39
L39:
	;
	if base.Ui64(v148) <= base.Ui64(int64(4194302)) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v160 = v148 + int64(1)
	goto L42
L41:
	;
	v160 = int64(0)
	goto L42
L42:
	;
	v164 = int32(4)
	v165 = base.I32_wrap_i64(v148)<<(uint(int32(10))%32) | v164
	v167 = v100 & int32(-1024)
	v169 = v167 | v164
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v169))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v165)) == int32(0) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	if v181 == int32(0) {
		v198 = int32(0)
		v201 = v160
		goto L36
	} else {
		goto L47
	}
L44:
	;
	v181 = base.B2i32(base.Ui32(v165) < base.Ui32(v169))
	goto L43
L45:
	;
	goto L46
L46:
	;
	v181 = int32(base.Ui32(v165-v169) >> (uint(int32(31)) % 32))
	goto L43
L47:
	;
	v185 = v167 + int32(1027)
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v185))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v165)) == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v198 = v197
	v201 = v160
	goto L36
L49:
	;
	v197 = base.B2i32(base.Ui32(v165) < base.Ui32(v185))
	goto L48
L50:
	;
	goto L51
L51:
	;
	v197 = int32(base.Ui32(v165-v185) >> (uint(int32(31)) % 32))
	goto L48
L52:
	;
	v223 = base.I64_extend_i32_u(v126)
	if v198 != 0 {
		goto L62
	} else {
		goto L63
	}
L53:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v204))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v100)) == int32(0) {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	v221 = v203
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v221)+8)) = v100
	goto L52
L56:
	;
	if v216 == int32(0) {
		goto L52
	} else {
		goto L60
	}
L57:
	;
	v216 = base.B2i32(base.Ui32(v204) < base.Ui32(v100))
	goto L56
L58:
	;
	goto L59
L59:
	;
	v216 = base.B2i32(int32(0) < v100-v204)
	goto L56
L60:
	;
	v220 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[7]))
	v221 = v220
	goto L55
L61:
	;
	v313 = int32(_a_F_GetSerializableTransactionSnapshotInt_1)
	v314 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[5]))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v314)+4))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v315+v301<<(uint(int32(2))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v319+v100<<(uint(int32(3))%32)&int32(_a_F_GetSerializableTransactionSnapshotInt_2)))) = v112
	v327 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[5]))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v327)+12))
	v330 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v328+v301))) = uint8(v330)
	F_LWLockRelease(m, v299)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L9
	} else {
		goto L79
	}
L62:
	;
	v225 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[7]))
	*(*int64)(unsafe.Add(mBase, uint32(v225))) = v223
	v228 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[5]))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)+28))
	v231 = int64(*(*uint16)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[6])))
	v232 = base.I64_rem_u_s(v201, v231)
	v236 = v229 + base.I32_wrap_i64(v232)<<(uint(int32(7))%32)
	v238 = F_LWLockAcquire(m, v236, int32(0))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L9
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v288 = v115 + v127<<(uint(int32(7))%32)
	v290 = F_LWLockAcquire(m, v288, int32(0))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L9
	} else {
		goto L77
	}
L65:
	;
	v241 = F_SimpleLruZeroPage(m, int32(_a_F_GetSerializableTransactionSnapshotInt_1), v201)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L9
	} else {
		goto L66
	}
L66:
	;
	if v201 == v223 {
		v299 = v236
		v301 = v241
		goto L61
	} else {
		goto L67
	}
L67:
	;
	v247 = v236
	v258 = v201
	goto L68
L68:
	;
	F_LWLockRelease(m, v247)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L9
	} else {
		goto L70
	}
L69:
	;
	v299 = v278
	v301 = v283
	goto L61
L70:
	;
	v264 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[5]))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)+28))
	if v258 <= int64(4194302) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v271 = v258 + int64(1)
	goto L73
L72:
	;
	v271 = int64(0)
	goto L73
L73:
	;
	v273 = int64(*(*uint16)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[6])))
	v274 = base.I64_rem_s(v271, v273)
	v278 = v265 + base.I32_wrap_i64(v274)<<(uint(int32(7))%32)
	v280 = F_LWLockAcquire(m, v278, int32(0))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L9
	} else {
		goto L74
	}
L74:
	;
	v283 = F_SimpleLruZeroPage(m, int32(_a_F_GetSerializableTransactionSnapshotInt_1), v271)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L9
	} else {
		goto L75
	}
L75:
	;
	if v271 != v223 {
		v247 = v278
		v258 = v271
		goto L68
	} else {
		goto L76
	}
L76:
	;
	goto L69
L77:
	;
	v294 = F_SimpleLruReadPage(m, int32(_a_F_GetSerializableTransactionSnapshotInt_1), v223, int32(1), v100)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L9
	} else {
		goto L78
	}
L78:
	;
	v299 = v288
	v301 = v294
	goto L61
L79:
	;
	goto L29
L80:
	;
	goto L22
L81:
	;
	goto L21
L82:
	;
	v402 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[2]))
	v406 = F_LWLockAcquire(m, v402+int32(3584), int32(0))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L9
	} else {
		goto L83
	}
L83:
	;
	v409 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[3]))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v409)+4))
	if base.B2i32(v410 == int32(0))|base.B2i32(v410 == v409) != 0 {
		goto L15
	} else {
		goto L84
	}
L84:
	;
	goto L16
L85:
	;
	F_errmsg_internal(m, int32(_a_F_GetSerializableTransactionSnapshotInt_3), int32(0))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L9
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(_a_F_GetSerializableTransactionSnapshotInt_4), int32(1784), int32(_a_F_GetSerializableTransactionSnapshotInt_5))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L9
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v432)+12)) = v451
	*(*int32)(unsafe.Add(mBase, uint32(v432)+8)) = v451
	goto L90
L89:
	;
	goto L90
L90:
	;
	v458 = v431 + int32(-64)
	*(*int32)(unsafe.Add(mBase, uint32(v431)+4)) = v451
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v451)))
	*(*int32)(unsafe.Add(mBase, uint32(v431))) = v460
	*(*int32)(unsafe.Add(mBase, uint32(v460)+4)) = v431
	*(*int32)(unsafe.Add(mBase, uint32(v451))) = v431
	if l1 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L91:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L9
	} else {
		goto L170
	}
L92:
	;
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v458)+64))
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v458)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v829)+4)) = v830
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v458)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v830))) = v832
	v835 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[3]))
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v835)+4))
	if v836 == int32(0) {
		goto L161
	} else {
		goto L162
	}
L93:
	;
	v475 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[3]))
	v477 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[8])))
	if v477 != int32(1) {
		goto L101
	} else {
		goto L102
	}
L94:
	;
	v466 = F_GetSnapshotData(m, l0)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L9
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v469 = F_ProcArrayInstallImportedXmin(m, v468, l1)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L9
	} else {
		goto L98
	}
L97:
	;
	v473 = v466
	goto L93
L98:
	;
	if v469 == int32(0) {
		goto L92
	} else {
		goto L99
	}
L99:
	;
	v473 = l0
	goto L93
L100:
	;
	m.G0 = v20 - int32(-64)
	return v473
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v458))) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v431-int32(60)))) = v35
	v508 = *(*int64)(unsafe.Add(mBase, uint32(v475)+32))
	v510 = v431 - int32(56)
	v511 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v510))) = v511
	*(*int64)(unsafe.Add(mBase, uint32(v431-int32(40)))) = v508
	*(*int64)(unsafe.Add(mBase, uint32(v510)+8)) = v511
	v518 = int32(24)
	v519 = v431 + v518
	*(*int32)(unsafe.Add(mBase, uint32(v431)+28)) = v519
	*(*int32)(unsafe.Add(mBase, uint32(v431)+24)) = v519
	v525 = v431 - v518
	*(*int32)(unsafe.Add(mBase, uint32(v431-int32(20)))) = v525
	v530 = v431 - int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v431-int32(28)))) = v530
	*(*int32)(unsafe.Add(mBase, uint32(v525))) = v525
	*(*int32)(unsafe.Add(mBase, uint32(v530))) = v530
	v535 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[9]))
	v536 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v431)+36)) = v536
	*(*int32)(unsafe.Add(mBase, uint32(v431)+32)) = v535
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v473)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v431)+40)) = v539
	v542 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[10]))
	*(*int32)(unsafe.Add(mBase, uint32(v431)+48)) = v542
	v545 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[11]))
	*(*int32)(unsafe.Add(mBase, uint32(v431)+44)) = v536
	*(*int64)(unsafe.Add(mBase, uint32(v431-int32(8)))) = int64(0)
	v555 = v431 - int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v431-int32(12)))) = v555
	*(*int32)(unsafe.Add(mBase, uint32(v431)+52)) = v545
	*(*int32)(unsafe.Add(mBase, uint32(v555))) = v555
	v560 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[8])))
	if v560 == int32(1) {
		goto L109
	} else {
		goto L110
	}
L102:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v475)+24))
	if v480 != 0 {
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v431)))
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v431)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v481)+4)) = v482
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v431)))
	*(*int32)(unsafe.Add(mBase, uint32(v482))) = v484
	v487 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[3]))
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v487)+4))
	if v488 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v487)+4)) = v487
	*(*int32)(unsafe.Add(mBase, uint32(v487))) = v487
	goto L106
L105:
	;
	goto L106
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v431)+4)) = v487
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v487)))
	*(*int32)(unsafe.Add(mBase, uint32(v431))) = v494
	*(*int32)(unsafe.Add(mBase, uint32(v494)+4)) = v431
	*(*int32)(unsafe.Add(mBase, uint32(v487))) = v431
	v499 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[2]))
	F_LWLockRelease(m, v499+int32(3584))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L9
	} else {
		goto L107
	}
L107:
	;
	goto L100
L108:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v473)+4))
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v696)+16))
	if v710 == int32(0) {
		goto L137
	} else {
		goto L138
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v431)+44)) = int32(32)
	v566 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[3]))
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v566)+12))
	if v567 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L110:
	;
	goto L111
L111:
	;
	v687 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[3]))
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v687)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v687)+24)) = v688 + int32(1)
	v696 = v687
	goto L108
L112:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v431)+28))
	if v659 != v519 {
		goto L128
	} else {
		goto L129
	}
L113:
	;
	v571 = v566 + int32(8)
	if v567 == v571 {
		goto L112
	} else {
		goto L114
	}
L114:
	;
	v574 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[12]))
	v577 = v567
	goto L115
L115:
	;
	v592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v577)+44)))
	if v592&int32(41) == int32(0) {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	goto L112
L117:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v574)+4))
	if base.B2i32(v597 == int32(0))|base.B2i32(v597 == v574) != 0 {
		goto L91
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v577)+4))
	if v640 != v571 {
		v577 = v640
		goto L115
	} else {
		goto L127
	}
L120:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v597)))
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v597)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v602)+4)) = v603
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v597)))
	*(*int32)(unsafe.Add(mBase, uint32(v603))) = v605
	*(*int32)(unsafe.Add(mBase, uint32(v597)+20)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v597)+16)) = v577 + int32(-64)
	v612 = v577 + int32(24)
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v577)+28))
	if v613 == int32(0) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v577)+28)) = v612
	*(*int32)(unsafe.Add(mBase, uint32(v577)+24)) = v612
	goto L123
L122:
	;
	goto L123
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v597)+4)) = v612
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v612)))
	*(*int32)(unsafe.Add(mBase, uint32(v597))) = v619
	*(*int32)(unsafe.Add(mBase, uint32(v619)+4)) = v597
	*(*int32)(unsafe.Add(mBase, uint32(v612))) = v597
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v431)+28))
	if v623 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v431)+28)) = v519
	*(*int32)(unsafe.Add(mBase, uint32(v431)+24)) = v431 + int32(24)
	goto L126
L125:
	;
	goto L126
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v597)+12)) = v519
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v519)))
	*(*int32)(unsafe.Add(mBase, uint32(v597)+8)) = v631
	v634 = v597 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v631)+4)) = v634
	*(*int32)(unsafe.Add(mBase, uint32(v519))) = v634
	goto L119
L127:
	;
	goto L116
L128:
	;
	v662 = v659
	goto L130
L129:
	;
	v662 = int32(0)
	goto L130
L130:
	;
	if v662 != 0 {
		v696 = v566
		goto L108
	} else {
		goto L131
	}
L131:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v431)))
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v431)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v663)+4)) = v664
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v431)))
	*(*int32)(unsafe.Add(mBase, uint32(v664))) = v666
	v669 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[3]))
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v669)+4))
	if v670 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v669)+4)) = v669
	*(*int32)(unsafe.Add(mBase, uint32(v669))) = v669
	goto L134
L133:
	;
	goto L134
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v431)+4)) = v669
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v669)))
	*(*int32)(unsafe.Add(mBase, uint32(v431))) = v676
	*(*int32)(unsafe.Add(mBase, uint32(v676)+4)) = v431
	*(*int32)(unsafe.Add(mBase, uint32(v669))) = v431
	v681 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[2]))
	F_LWLockRelease(m, v681+int32(3584))
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L9
	} else {
		goto L135
	}
L135:
	;
	goto L100
L136:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[13])) = v458
	v788 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[14])) = uint8(v788)
	v791 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[2]))
	F_LWLockRelease(m, v791+int32(3584))
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L9
	} else {
		goto L159
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v696)+20)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v696)+16)) = v709
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v473)+4))
	v718 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[2]))
	v722 = F_LWLockAcquire(m, v718+int32(_a_F_GetSerializableTransactionSnapshotInt_0), int32(0))
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L9
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	if v709 != v710 {
		goto L136
	} else {
		goto L158
	}
L140:
	;
	if v716 == int32(0) {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	v772 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[2]))
	F_LWLockRelease(m, v772+int32(_a_F_GetSerializableTransactionSnapshotInt_0))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L9
	} else {
		goto L157
	}
L142:
	;
	v727 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[7]))
	*(*int64)(unsafe.Add(mBase, uint32(v727)+8)) = int64(0)
	goto L141
L143:
	;
	goto L144
L144:
	;
	v732 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[15])))
	if v732 == int32(1) {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	v744 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[7]))
	if v742 == int32(0) {
		v767 = v744
		goto L149
	} else {
		goto L150
	}
L146:
	;
	v737 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[16]))
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v737)+316))
	v740 = base.B2i32(v738 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[15])) = uint8(v740)
	v742 = v740
	goto L148
L147:
	;
	v742 = int32(0)
	goto L148
L148:
	;
	goto L145
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v767)+12)) = v716
	goto L141
L150:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v744)+12))
	if v747 == int32(0) {
		v767 = v744
		goto L149
	} else {
		goto L151
	}
L151:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v747))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v716)) == int32(0) {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	if v761 == int32(0) {
		goto L141
	} else {
		goto L156
	}
L153:
	;
	v761 = base.B2i32(base.Ui32(v716) < base.Ui32(v747))
	goto L152
L154:
	;
	goto L155
L155:
	;
	v761 = int32(base.Ui32(v716-v747) >> (uint(int32(31)) % 32))
	goto L152
L156:
	;
	v765 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[7]))
	v767 = v765
	goto L149
L157:
	;
	goto L136
L158:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v696)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v696)+20)) = v778 + int32(1)
	goto L136
L159:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+32)) = int64(103079215120)
	v801 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[17]))
	v805 = F_hash_create(m, int32(_a_F_GetSerializableTransactionSnapshotInt_6), v801, v18+int32(-48), int32(40))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L9
	} else {
		goto L160
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[18])) = v805
	goto L100
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v835)+4)) = v835
	*(*int32)(unsafe.Add(mBase, uint32(v835))) = v835
	goto L163
L162:
	;
	goto L163
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v458)+68)) = v835
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v835)))
	*(*int32)(unsafe.Add(mBase, uint32(v458)+64)) = v842
	v845 = v458 - int32(-64)
	*(*int32)(unsafe.Add(mBase, uint32(v842)+4)) = v845
	*(*int32)(unsafe.Add(mBase, uint32(v835))) = v845
	v849 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[2]))
	F_LWLockRelease(m, v849+int32(3584))
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L9
	} else {
		goto L164
	}
L164:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L9
	} else {
		goto L165
	}
L165:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L9
	} else {
		goto L166
	}
L166:
	;
	F_errmsg(m, int32(_a_F_GetSerializableTransactionSnapshotInt_7), int32(0))
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L9
	} else {
		goto L167
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = l2
	F_errdetail(m, int32(_a_F_GetSerializableTransactionSnapshotInt_8), v20)
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L9
	} else {
		goto L168
	}
L168:
	;
	F_errfinish(m, int32(_a_F_GetSerializableTransactionSnapshotInt_4), int32(1829), int32(_a_F_GetSerializableTransactionSnapshotInt_5))
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L9
	} else {
		goto L169
	}
L169:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L170:
	;
	F_errcode(m, int32(_a_F_GetSerializableTransactionSnapshotInt_9))
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L9
	} else {
		goto L171
	}
L171:
	;
	F_errmsg(m, int32(_a_F_GetSerializableTransactionSnapshotInt_10), int32(0))
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L9
	} else {
		goto L172
	}
L172:
	;
	F_errhint(m, int32(_a_F_GetSerializableTransactionSnapshotInt_11), int32(0))
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L9
	} else {
		goto L173
	}
L173:
	;
	F_errfinish(m, int32(_a_F_GetSerializableTransactionSnapshotInt_4), int32(679), int32(_a_F_GetSerializableTransactionSnapshotInt_12))
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L9
	} else {
		goto L174
	}
L174:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_GetSnapshotData(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int64
	_ = v64
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int64
	_ = v108
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v192 int32
	_ = v192
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v390 int32
	_ = v390
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v461 int64
	_ = v461
	var v463 int64
	_ = v463
	var v468 int64
	_ = v468
	var v473 int64
	_ = v473
	var v475 int64
	_ = v475
	var v477 int64
	_ = v477
	var v479 int64
	_ = v479
	var v481 int64
	_ = v481
	var v483 int64
	_ = v483
	var v484 int64
	_ = v484
	var v486 int64
	_ = v486
	var v491 int32
	_ = v491
	var v493 int64
	_ = v493
	var v495 int64
	_ = v495
	var v497 int64
	_ = v497
	var v499 int64
	_ = v499
	var v507 int64
	_ = v507
	var v515 int64
	_ = v515
	var v519 int64
	_ = v519
	var v522 int64
	_ = v522
	var v524 int64
	_ = v524
	var v526 int64
	_ = v526
	var v527 int32
	_ = v527
	var v529 int64
	_ = v529
	var v532 int64
	_ = v532
	var v534 int64
	_ = v534
	var v535 int64
	_ = v535
	var v537 int32
	_ = v537
	var v539 int64
	_ = v539
	var v542 int64
	_ = v542
	var v544 int64
	_ = v544
	var v545 int64
	_ = v545
	var v548 int64
	_ = v548
	var v551 int64
	_ = v551
	var v553 int64
	_ = v553
	var v554 int64
	_ = v554
	var v563 int32
	_ = v563
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v609 int32
	_ = v609
	v2 = int32(0)
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[0]))
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[1]))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v33 == v2 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L8
	} else {
		goto L176
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L8
	} else {
		goto L172
	}
L3:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v39 = F_emscripten_builtin_malloc(m, v36<<(uint(int32(2))%32))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v39
	if v39 == int32(0) {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[2]))
	v60 = F_LWLockAcquire(m, v56+int32(512), int32(1))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[3]))
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[4]))
	v50 = F_emscripten_builtin_malloc(m, (v44+v46)*int32(260))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v50
	if v50 == int32(0) {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	return int32(0)
L9:
	;
	v64 = *(*int64)(unsafe.Add(mBase, uint32(l0)+64))
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[5]))
	v69 = *(*int64)(unsafe.Add(mBase, uint32(v68)+56))
	if base.B2i32(v64 == int64(0))|base.B2i32(v64 != v69) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[6]))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+40))
	if v77 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[6]))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+48))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v32+v102<<(uint(int32(2))%32))))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
	v108 = *(*int64)(unsafe.Add(mBase, uint32(v68)+48))
	v111 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetSnapshotData[7])))
	if v111 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[8])) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v76)+40)) = v74
	goto L15
L14:
	;
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[9])) = v74
	v86 = F_GetCurrentCommandId(m, int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+44)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v86
	v91 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)) = uint8(v91)
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[2]))
	F_LWLockRelease(m, v94+int32(512))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	return l0
L18:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)) = uint8(v121)
	v123 = int32(3)
	v124 = base.I32_wrap_i64(v108)
	v126 = v124 + int32(1)
	if base.Ui32(v126) <= base.Ui32(v123) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[10]))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+316))
	v119 = base.B2i32(v117 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_GetSnapshotData[7])) = uint8(v119)
	v121 = v119
	goto L21
L20:
	;
	v121 = int32(0)
	goto L21
L21:
	;
	goto L18
L22:
	;
	v129 = v123
	goto L24
L23:
	;
	v129 = v126
	goto L24
L24:
	;
	if v106-v129 < int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v133 = v106
	goto L27
L26:
	;
	v133 = v129
	goto L27
L27:
	;
	if base.Ui32(int32(2)) < base.Ui32(v106) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v136 = v133
	goto L30
L29:
	;
	v136 = v129
	goto L30
L30:
	;
	if v121 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v406 = *(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[0]))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v406)+32))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v406)+28))
	v410 = *(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[6]))
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v410)+40))
	if v411 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L32:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if v139 <= int32(0) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	v247 = *(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[0]))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v247)+20))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v247)+16))
	if v248 <= v250 {
		goto L58
	} else {
		goto L59
	}
L35:
	;
	v379 = v136
	v381 = int32(0)
	v383 = v2
	v390 = v2
	goto L31
L36:
	;
	goto L37
L37:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[1]))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+12))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v147)+8))
	v150 = int32(0)
	v153 = v136
	v155 = v150
	v156 = v150
	v157 = v2
	v164 = v2
	goto L38
L38:
	;
	v180 = v156 << (uint(int32(2)) % 32)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v32+v180)))
	v183 = int32(0)
	if base.B2i32(v182 == v183)|base.B2i32(v156 == v102)|base.B2i32(v183 <= v182-v129) != 0 {
		v235 = v153
		v237 = v155
		v238 = v157
		v240 = v164
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v379 = v235
	v381 = v237
	v383 = v238
	v390 = v240
	goto L31
L40:
	;
	v244 = v156 + int32(1)
	if v244 != v139 {
		v153 = v235
		v155 = v237
		v156 = v244
		v157 = v238
		v164 = v240
		goto L38
	} else {
		goto L56
	}
L41:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156+v148))))
	if v192&int32(18) != 0 {
		v235 = v153
		v237 = v155
		v238 = v157
		v240 = v164
		goto L40
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v145+v164<<(uint(int32(2))%32)))) = v182
	if v182-v153 < int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v202 = v182
	goto L45
L44:
	;
	v202 = v153
	goto L45
L45:
	;
	v203 = int32(1)
	v204 = v164 + v203
	if v157&v203 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v235 = v202
	v237 = v155
	v238 = int32(1)
	v240 = v204
	goto L40
L47:
	;
	goto L48
L48:
	;
	v208 = int32(1)
	v211 = v149 + v156<<(uint(v208)%32)
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+1)))
	if v212 != 0 {
		v235 = v202
		v237 = v155
		v238 = v208
		v240 = v204
		goto L40
	} else {
		goto L49
	}
L49:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211))))
	if v213 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v235 = v202
	v237 = v155
	v238 = int32(0)
	v240 = v204
	goto L40
L51:
	;
	goto L52
L52:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v180+(v29+int32(36)))))
	v221 = *(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[11]))
	v223 = v213 << (uint(int32(2)) % 32)
	if v223 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	base.MemoryCopy(m, v224+v155<<(uint(int32(2))%32), v221+v218*int32(640)+int32(280), v223)
	goto L55
L54:
	;
	goto L55
L55:
	;
	v235 = v202
	v237 = v213 + v155
	v238 = int32(0)
	v240 = v204
	goto L40
L56:
	;
	goto L39
L57:
	;
	v364 = *(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[0]))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v364)+24))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v365))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v337)) == int32(0) {
		goto L83
	} else {
		goto L84
	}
L58:
	;
	v337 = v136
	v339 = int32(0)
	goto L57
L59:
	;
	goto L60
L60:
	;
	v254 = *(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[12]))
	v257 = v136
	v259 = int32(0)
	v260 = v250
	v262 = v254
	goto L61
L61:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260+v262))))
	if v284 == int32(1) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v337 = v329
	v339 = v330
	goto L57
L63:
	;
	v288 = *(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[13]))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v288+v260<<(uint(int32(2))%32))))
	if v259 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	v329 = v257
	v330 = v259
	v332 = v262
	goto L65
L65:
	;
	v334 = v260 + int32(1)
	if v334 != v248 {
		v257 = v329
		v259 = v330
		v260 = v334
		v262 = v332
		goto L61
	} else {
		goto L81
	}
L66:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v257))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v292)) == int32(0) {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	v308 = v257
	goto L68
L68:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v129))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v292)) == int32(0) {
		goto L77
	} else {
		goto L78
	}
L69:
	;
	if v306 != 0 {
		goto L73
	} else {
		goto L74
	}
L70:
	;
	v306 = base.B2i32(base.Ui32(v292) < base.Ui32(v257))
	goto L69
L71:
	;
	goto L72
L72:
	;
	v306 = int32(base.Ui32(v292-v257) >> (uint(int32(31)) % 32))
	goto L69
L73:
	;
	v307 = v292
	goto L75
L74:
	;
	v307 = v257
	goto L75
L75:
	;
	v308 = v307
	goto L68
L76:
	;
	if v320 != 0 {
		v337 = v308
		v339 = v259
		goto L57
	} else {
		goto L80
	}
L77:
	;
	v320 = base.B2i32(base.Ui32(v129) <= base.Ui32(v292))
	goto L76
L78:
	;
	goto L79
L79:
	;
	v320 = base.B2i32(int32(0) <= v292-v129)
	goto L76
L80:
	;
	v322 = *(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[12]))
	*(*int32)(unsafe.Add(mBase, uint32(v249+v259<<(uint(int32(2))%32)))) = v292
	v329 = v308
	v330 = v259 + int32(1)
	v332 = v322
	goto L65
L81:
	;
	goto L62
L82:
	;
	v379 = v337
	v381 = v339
	v383 = v377
	v390 = v2
	goto L31
L83:
	;
	v377 = base.B2i32(base.Ui32(v337) <= base.Ui32(v365))
	goto L82
L84:
	;
	goto L85
L85:
	;
	v377 = base.B2i32(v337-v365 <= int32(0))
	goto L82
L86:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[8])) = v379
	*(*int32)(unsafe.Add(mBase, uint32(v410)+40)) = v379
	goto L88
L87:
	;
	goto L88
L88:
	;
	v418 = *(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[2]))
	F_LWLockRelease(m, v418+int32(512))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L8
	} else {
		goto L89
	}
L89:
	;
	if v379 == int32(0) {
		v440 = v408
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v461 = v108 + base.I64_extend_i32_s(v440-v124)
	v463 = *(*int64)(unsafe.Add(mBase, _c_F_GetSnapshotData[14]))
	v468 = v108 + base.I64_extend_i32_s(v458-v124)
	if base.I32_wrap_i64(v468) == int32(0) {
		goto L111
	} else {
		goto L112
	}
L91:
	;
	if v407 == int32(0) {
		v458 = v440
		goto L90
	} else {
		goto L101
	}
L92:
	;
	if v408 == int32(0) {
		v440 = v379
		goto L91
	} else {
		goto L93
	}
L93:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v408))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v379)) == int32(0) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	if v438 != 0 {
		goto L98
	} else {
		goto L99
	}
L95:
	;
	v438 = base.B2i32(base.Ui32(v379) < base.Ui32(v408))
	goto L94
L96:
	;
	goto L97
L97:
	;
	v438 = int32(base.Ui32(v379-v408) >> (uint(int32(31)) % 32))
	goto L94
L98:
	;
	v439 = v379
	goto L100
L99:
	;
	v439 = v408
	goto L100
L100:
	;
	v440 = v439
	goto L91
L101:
	;
	if v440 == int32(0) {
		v458 = v407
		goto L90
	} else {
		goto L102
	}
L102:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v440))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v407)) == int32(0) {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	if v456 != 0 {
		goto L107
	} else {
		goto L108
	}
L104:
	;
	v456 = base.B2i32(base.Ui32(v407) < base.Ui32(v440))
	goto L103
L105:
	;
	goto L106
L106:
	;
	v456 = int32(base.Ui32(v407-v440) >> (uint(int32(31)) % 32))
	goto L103
L107:
	;
	v457 = v407
	goto L109
L108:
	;
	v457 = v440
	goto L109
L109:
	;
	v458 = v457
	goto L90
L110:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_GetSnapshotData[15])) = v486
	*(*int64)(unsafe.Add(mBase, _c_F_GetSnapshotData[14])) = v484
	v491 = int32(_a_F_GetSnapshotData_0)
	v493 = *(*int64)(unsafe.Add(mBase, _c_F_GetSnapshotData[16]))
	if base.Ui64(v493) < base.Ui64(v461) {
		goto L126
	} else {
		goto L127
	}
L111:
	;
	v473 = *(*int64)(unsafe.Add(mBase, _c_F_GetSnapshotData[15]))
	v484 = v463
	v486 = v473
	goto L110
L112:
	;
	goto L113
L113:
	;
	if base.Ui64(v463) < base.Ui64(v468) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v475 = v468
	goto L116
L115:
	;
	v475 = v463
	goto L116
L116:
	;
	if base.I32_wrap_i64(v463) != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v477 = v475
	goto L119
L118:
	;
	v477 = v468
	goto L119
L119:
	;
	v479 = *(*int64)(unsafe.Add(mBase, _c_F_GetSnapshotData[15]))
	if base.Ui64(v479) < base.Ui64(v468) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v481 = v468
	goto L122
L121:
	;
	v481 = v479
	goto L122
L122:
	;
	if base.I32_wrap_i64(v479) != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v483 = v481
	goto L125
L124:
	;
	v483 = v468
	goto L125
L125:
	;
	v484 = v477
	v486 = v483
	goto L110
L126:
	;
	v495 = v461
	goto L128
L127:
	;
	v495 = v493
	goto L128
L128:
	;
	if base.I32_wrap_i64(v493) != 0 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v497 = v495
	goto L131
L130:
	;
	v497 = v461
	goto L131
L131:
	;
	if base.I32_wrap_i64(v461) != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v499 = v497
	goto L134
L133:
	;
	v499 = v493
	goto L134
L134:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_GetSnapshotData[16])) = v499
	if base.Ui32(int32(3)) <= base.Ui32(v106) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v524 = v108 + base.I64_extend_i32_s(v106-v124)
	goto L137
L136:
	;
	v507 = int64(1)
	v515 = v108 + v507
	if base.Ui32(base.I32_wrap_i64(v515)) < base.Ui32(int32(3)) {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_GetSnapshotData[17])) = v524
	v526 = v108 + base.I64_extend_i32_s(v107-v124)
	v527 = int32(_a_F_GetSnapshotData_1)
	v529 = *(*int64)(unsafe.Add(mBase, _c_F_GetSnapshotData[18]))
	if base.I32_wrap_i64(v529) != 0 {
		goto L144
	} else {
		goto L145
	}
L138:
	;
	v519 = v108 + (v507-v108)&int64(4294967295) + int64(2)
	goto L140
L139:
	;
	v519 = v515
	goto L140
L140:
	;
	if base.Ui64(int64(2)) < base.Ui64(v515) {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v522 = v519
	goto L143
L142:
	;
	v522 = v515
	goto L143
L143:
	;
	v524 = v522
	goto L137
L144:
	;
	if base.Ui64(v526) < base.Ui64(v529) {
		goto L147
	} else {
		goto L148
	}
L145:
	;
	v535 = v526
	goto L146
L146:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_GetSnapshotData[18])) = v535
	v537 = int32(_a_F_GetSnapshotData_2)
	v539 = *(*int64)(unsafe.Add(mBase, _c_F_GetSnapshotData[19]))
	if base.I32_wrap_i64(v539) != 0 {
		goto L153
	} else {
		goto L154
	}
L147:
	;
	v532 = v529
	goto L149
L148:
	;
	v532 = v526
	goto L149
L149:
	;
	if base.I32_wrap_i64(v526) != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v534 = v532
	goto L152
L151:
	;
	v534 = v529
	goto L152
L152:
	;
	v535 = v534
	goto L146
L153:
	;
	if base.Ui64(v526) < base.Ui64(v539) {
		goto L156
	} else {
		goto L157
	}
L154:
	;
	v545 = v526
	goto L155
L155:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_GetSnapshotData[19])) = v545
	v548 = *(*int64)(unsafe.Add(mBase, _c_F_GetSnapshotData[20]))
	if base.I32_wrap_i64(v548) != 0 {
		goto L162
	} else {
		goto L163
	}
L156:
	;
	v542 = v539
	goto L158
L157:
	;
	v542 = v526
	goto L158
L158:
	;
	if base.I32_wrap_i64(v526) != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v544 = v542
	goto L161
L160:
	;
	v544 = v539
	goto L161
L161:
	;
	v545 = v544
	goto L155
L162:
	;
	if base.Ui64(v526) < base.Ui64(v548) {
		goto L165
	} else {
		goto L166
	}
L163:
	;
	v554 = v526
	goto L164
L164:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_GetSnapshotData[21])) = v524
	*(*int64)(unsafe.Add(mBase, _c_F_GetSnapshotData[20])) = v554
	*(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[9])) = v379
	*(*int64)(unsafe.Add(mBase, uint32(l0)+64)) = v69
	v563 = v383 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v563)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v381
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v379
	v570 = F_GetCurrentCommandId(m, int32(0))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L8
	} else {
		goto L171
	}
L165:
	;
	v551 = v548
	goto L167
L166:
	;
	v551 = v526
	goto L167
L167:
	;
	if base.I32_wrap_i64(v526) != 0 {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v553 = v551
	goto L170
L169:
	;
	v553 = v548
	goto L170
L170:
	;
	v554 = v553
	goto L164
L171:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+44)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v570
	v575 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)) = uint8(v575)
	return l0
L172:
	;
	F_errcode(m, int32(_a_F_GetSnapshotData_3))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L8
	} else {
		goto L173
	}
L173:
	;
	F_errmsg(m, int32(_a_F_GetSnapshotData_4), int32(0))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L8
	} else {
		goto L174
	}
L174:
	;
	F_errfinish(m, int32(_a_F_GetSnapshotData_5), int32(2217), int32(_a_F_GetSnapshotData_6))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L8
	} else {
		goto L175
	}
L175:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L176:
	;
	F_errcode(m, int32(_a_F_GetSnapshotData_3))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L8
	} else {
		goto L177
	}
L177:
	;
	F_errmsg(m, int32(_a_F_GetSnapshotData_4), int32(0))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L8
	} else {
		goto L178
	}
L178:
	;
	F_errfinish(m, int32(_a_F_GetSnapshotData_5), int32(2224), int32(_a_F_GetSnapshotData_6))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L8
	} else {
		goto L179
	}
L179:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_GetVisibilityMapPins(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
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
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v263 int32
	_ = v263
	v8 = int32(0)
	if base.B2i32(l2 == v8)|base.B2i32(base.Ui32(l3) <= base.Ui32(l4)) != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v32 = v26 ^ int32(-1)
	v34 = v26 << (uint(int32(13)) % 32)
	if int32(0) <= v26 {
		goto L15
	} else {
		goto L16
	}
L2:
	;
	v22 = l1
	goto L4
L3:
	;
	v22 = v8
	goto L4
L4:
	;
	if v22 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v25 = l4
	v26 = l2
	v27 = l6
	v28 = l5
	v29 = l3
	v30 = l1
	goto L1
L6:
	;
	goto L7
L7:
	;
	v25 = l3
	v26 = l1
	v27 = l5
	v28 = l6
	v29 = l4
	v30 = l2
	goto L1
L8:
	;
	return v263
L9:
	;
	F_LockBuffer(m, v26, int32(2))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L39
	} else {
		goto L54
	}
L10:
	;
	F_visibilitymap_pin(m, l0, v29, v28)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L39
	} else {
		goto L53
	}
L11:
	;
	if v100 != 0 {
		goto L10
	} else {
		goto L52
	}
L12:
	;
	if v60 != 0 {
		v263 = v8
		goto L8
	} else {
		goto L48
	}
L13:
	;
	v69 = v30 ^ int32(-1)
	v71 = v30 << (uint(int32(13)) % 32)
	if int32(0) <= v30 {
		goto L28
	} else {
		goto L29
	}
L14:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+10)))
	if v49&int32(4) != 0 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_GetVisibilityMapPins[0]))
	v48 = v38 + v34 + int32(-8192)
	goto L14
L16:
	;
	goto L17
L17:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_GetVisibilityMapPins[1]))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43+v32<<(uint(int32(2))%32))))
	v48 = v47
	goto L14
L18:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if v52 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	goto L20
L20:
	;
	if v30 == int32(0) {
		v263 = v8
		goto L8
	} else {
		goto L26
	}
L21:
	;
	if v30 == int32(0) {
		goto L12
	} else {
		goto L25
	}
L22:
	;
	v60 = int32(0)
	goto L21
L23:
	;
	goto L24
L24:
	;
	v56 = F_BufferGetBlockNumber(m, v52)
	mBase = m.M
	v58 = base.I32_div_u_s(v25, int32(_a_F_GetVisibilityMapPins_0))
	v60 = base.B2i32(v56 == v58)
	goto L21
L25:
	;
	v67 = v60 ^ int32(1)
	goto L13
L26:
	;
	v67 = v8
	goto L13
L27:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+10)))
	if v86&int32(4) != 0 {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_GetVisibilityMapPins[0]))
	v85 = v75 + v71 + int32(-8192)
	goto L27
L29:
	;
	goto L30
L30:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_GetVisibilityMapPins[1]))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v80+v69<<(uint(int32(2))%32))))
	v85 = v84
	goto L27
L31:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v89 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	v100 = v8
	goto L33
L33:
	;
	if (v67|v100)&int32(1) == int32(0) {
		v263 = v8
		goto L8
	} else {
		goto L38
	}
L34:
	;
	v100 = v97 ^ int32(1)
	goto L33
L35:
	;
	v97 = int32(0)
	goto L34
L36:
	;
	goto L37
L37:
	;
	v93 = F_BufferGetBlockNumber(m, v89)
	mBase = m.M
	v95 = base.I32_div_u_s(v29, int32(_a_F_GetVisibilityMapPins_0))
	v97 = base.B2i32(v93 == v95)
	goto L34
L38:
	;
	F_LockBuffer(m, v26, int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	return int32(0)
L40:
	;
	v111 = int32(0)
	v114 = base.B2i32(v30 == v111) | base.B2i32(l1 == l2)
	if v114 == v111 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	F_LockBuffer(m, v30, int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L39
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	if v67 == int32(0) {
		goto L11
	} else {
		goto L45
	}
L44:
	;
	goto L43
L45:
	;
	F_visibilitymap_pin(m, l0, v25, v27)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L39
	} else {
		goto L46
	}
L46:
	;
	if v100 != 0 {
		goto L10
	} else {
		goto L47
	}
L47:
	;
	v138 = int32(0)
	goto L9
L48:
	;
	F_LockBuffer(m, v26, int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L39
	} else {
		goto L49
	}
L49:
	;
	F_visibilitymap_pin(m, l0, v25, v27)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L39
	} else {
		goto L50
	}
L50:
	;
	F_LockBuffer(m, v26, int32(2))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L39
	} else {
		goto L51
	}
L51:
	;
	return int32(1)
L52:
	;
	v138 = int32(0)
	goto L9
L53:
	;
	v138 = v67
	goto L9
L54:
	;
	v142 = int32(1)
	if v114 != 0 {
		v263 = v142
		goto L8
	} else {
		goto L55
	}
L55:
	;
	F_LockBuffer(m, v30, int32(2))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L39
	} else {
		goto L56
	}
L56:
	;
	if v138 != 0 {
		v263 = v142
		goto L8
	} else {
		goto L57
	}
L57:
	;
	goto L58
L58:
	;
	v166 = int32(0)
	if base.B2i32(int32(0) <= v26) == v166 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	v263 = v142
	goto L8
L60:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+10)))
	if v179&int32(4) != 0 {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	v170 = *(*int32)(unsafe.Add(mBase, _c_F_GetVisibilityMapPins[1]))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v170+v32<<(uint(int32(2))%32))))
	v178 = v172
	goto L60
L62:
	;
	goto L63
L63:
	;
	v174 = *(*int32)(unsafe.Add(mBase, _c_F_GetVisibilityMapPins[0]))
	v178 = v174 + v34 + int32(-8192)
	goto L60
L64:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if v182 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	v193 = v166
	goto L66
L66:
	;
	v194 = int32(0)
	if v30 < v194 {
		goto L72
	} else {
		goto L73
	}
L67:
	;
	v193 = v190 ^ int32(1)
	goto L66
L68:
	;
	v190 = int32(0)
	goto L67
L69:
	;
	goto L70
L70:
	;
	v186 = F_BufferGetBlockNumber(m, v182)
	mBase = m.M
	v188 = base.I32_div_u_s(v25, int32(_a_F_GetVisibilityMapPins_0))
	v190 = base.B2i32(v186 == v188)
	goto L67
L71:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+10)))
	if v209&int32(4) != 0 {
		goto L75
	} else {
		goto L76
	}
L72:
	;
	v198 = *(*int32)(unsafe.Add(mBase, _c_F_GetVisibilityMapPins[1]))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v198+v69<<(uint(int32(2))%32))))
	v208 = v202
	goto L71
L73:
	;
	goto L74
L74:
	;
	v204 = *(*int32)(unsafe.Add(mBase, _c_F_GetVisibilityMapPins[0]))
	v208 = v204 + v71 + int32(-8192)
	goto L71
L75:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v212 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L76:
	;
	v223 = v194
	goto L77
L77:
	;
	if (v193|v223)&int32(1) == int32(0) {
		v263 = v142
		goto L8
	} else {
		goto L82
	}
L78:
	;
	v223 = v220 ^ int32(1)
	goto L77
L79:
	;
	v220 = int32(0)
	goto L78
L80:
	;
	goto L81
L81:
	;
	v216 = F_BufferGetBlockNumber(m, v212)
	mBase = m.M
	v218 = base.I32_div_u_s(v29, int32(_a_F_GetVisibilityMapPins_0))
	v220 = base.B2i32(v216 == v218)
	goto L78
L82:
	;
	F_LockBuffer(m, v26, int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L39
	} else {
		goto L83
	}
L83:
	;
	F_LockBuffer(m, v30, int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L39
	} else {
		goto L84
	}
L84:
	;
	if v193 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	F_visibilitymap_pin(m, l0, v25, v27)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L39
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	if v223 != 0 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	goto L87
L89:
	;
	F_visibilitymap_pin(m, l0, v29, v28)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L39
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	F_LockBuffer(m, v26, int32(2))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L39
	} else {
		goto L93
	}
L92:
	;
	goto L91
L93:
	;
	F_LockBuffer(m, v30, int32(2))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L39
	} else {
		goto L94
	}
L94:
	;
	if v193&v223 != int32(1) {
		goto L58
	} else {
		goto L95
	}
L95:
	;
	goto L59
}
func F___getf2(m *base.Module, l0 int64, l1 int64, l2 int64) int32 {
	var v7 int32
	_ = v7
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v16 int32
	_ = v16
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	v7 = int32(-1)
	v11 = l1 & int64(9223372036854775807)
	v12 = int64(9223090561878065152)
	if v11 == v12 {
		v16 = base.B2i32(l0 != int64(0))
	} else {
		v16 = base.B2i32(base.Ui64(v12) < base.Ui64(v11))
	}
	if v16 != 0 {
		v50 = v7
		return v50
	} else {
		v18 = l2 & int64(9223372036854775807)
		v19 = int64(9223090561878065152)
		if base.B2i32(base.Ui64(v19) < base.Ui64(v18))&base.B2i32(v18 != v19) != 0 {
			v50 = v7
			return v50
		} else {
			if l0|(v11|v18) == int64(0) {
				return int32(0)
			} else {
				if int64(0) <= l1&l2 {
					if base.B2i32(l1 != l2)&base.B2i32(l1 < l2) != 0 {
						v50 = v7
						return v50
					} else {
						return base.B2i32(l0|(l1^l2) != int64(0))
					}
				} else {
					if l1 == l2 {
						v45 = base.B2i32(l0 != int64(0))
					} else {
						v45 = base.B2i32(l2 < l1)
					}
					if v45 != 0 {
						v50 = v7
					} else {
						v50 = base.B2i32(l0|(l1^l2) != int64(0))
					}
					return v50
				}
			}
		}
	}
}
func F_gbk_to_utf8(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v3 = int32(0)
	v7 = Fn13870(m, l0, int32(37), v3, v3, v3, int32(_a_F_gbk_to_utf8_0))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_generateClonedIndexStmt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
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
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int64
	_ = v79
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v324 int32
	_ = v324
	var v331 int32
	_ = v331
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
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
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v577 int32
	_ = v577
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v592 int32
	_ = v592
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v630 int32
	_ = v630
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v645 int32
	_ = v645
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v660 int32
	_ = v660
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v691 int32
	_ = v691
	var v698 int32
	_ = v698
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v842 int32
	_ = v842
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v862 int32
	_ = v862
	var v867 int32
	_ = v867
	v26 = m.G0
	v28 = v26 - int32(144)
	m.G0 = v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	if l3 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	goto L3
L2:
	;
	goto L3
L3:
	;
	v34 = F_SearchSysCache1(m, int32(57), v30)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L12
	} else {
		goto L13
	}
L4:
	;
	v691 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41)+8)))
	if v669 < v691 {
		goto L158
	} else {
		goto L159
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L12
	} else {
		goto L153
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L12
	} else {
		goto L150
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L12
	} else {
		goto L145
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L12
	} else {
		goto L142
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L12
	} else {
		goto L139
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L12
	} else {
		goto L136
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L12
	} else {
		goto L133
	}
L12:
	;
	return int32(0)
L13:
	;
	if v34 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+196))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+22)))
	v41 = v39 + v40
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+22)))
	v46 = v44 + v45
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+84))
	v48 = F_SearchSysCache1(m, int32(2), v47)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L12
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L12
	} else {
		goto L130
	}
L17:
	;
	if v48 == int32(0) {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+22)))
	v56 = F_SysCacheGetAttrNotNull(m, int32(34), v38, int32(17))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L12
	} else {
		goto L19
	}
L19:
	;
	v60 = F_SysCacheGetAttrNotNull(m, int32(34), v38, int32(18))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	v63 = F_palloc0(m, int32(72))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v63)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = int32(204)
	v71 = F_pstrdup(m, v52+v53+int32(4))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L12
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v63)+12)) = v71
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v46)+92))
	if v74 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v75 = F_get_tablespace_name(m, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L12
	} else {
		goto L26
	}
L24:
	;
	v78 = int32(0)
	goto L25
L25:
	;
	v79 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v63)+36)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v63)+16)) = v78
	*(*int64)(unsafe.Add(mBase, uint32(v63)+44)) = v79
	*(*int64)(unsafe.Add(mBase, uint32(v63)+52)) = v79
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v63)+60)) = uint8(v86)
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+13)))
	*(*uint8)(unsafe.Add(mBase, uint32(v63)+61)) = uint8(v88)
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+14)))
	*(*uint8)(unsafe.Add(mBase, uint32(v63)+62)) = uint8(v90)
	if v90 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v78 = v75
	goto L25
L27:
	;
	v100 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v63)+67)) = v100
	v102 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v63)+4)) = v102
	v105 = v99 & v100
	*(*uint8)(unsafe.Add(mBase, uint32(v63)+64)) = uint8(v105)
	if v90|v86 == v102 {
		goto L34
	} else {
		goto L35
	}
L28:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+12)))
	if v95 != int32(1) {
		v99 = int32(0)
		goto L27
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+15)))
	v99 = v98
	goto L27
L31:
	;
	goto L30
L32:
	;
	v283 = F_SysCacheGetAttr(m, int32(34), v38, int32(20), v28+int32(135))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L12
	} else {
		goto L66
	}
L33:
	;
	v252 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v63)+63)) = uint8(v252)
	goto L32
L34:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+15)))
	if v110 != int32(1) {
		goto L33
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v113 = F_get_index_constraint(m, v30)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L12
	} else {
		goto L38
	}
L37:
	;
	goto L36
L38:
	;
	if v113 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	if l3 != 0 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L41
L41:
	;
	v250 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v63)+63)) = uint8(v250)
	goto L32
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v113
	goto L44
L43:
	;
	goto L44
L44:
	;
	v117 = F_SearchSysCache1(m, int32(19), v113)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L12
	} else {
		goto L45
	}
L45:
	;
	if v117 == int32(0) {
		goto L10
	} else {
		goto L46
	}
L46:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v117)+16))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+22)))
	v123 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v63)+63)) = uint8(v123)
	v125 = v121 + v122
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125)+73)))
	*(*uint8)(unsafe.Add(mBase, uint32(v63)+65)) = uint8(v126)
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125)+74)))
	*(*uint8)(unsafe.Add(mBase, uint32(v63)+66)) = uint8(v128)
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+15)))
	if v130 != v123 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	F_ReleaseCatCache(m, v117)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L12
	} else {
		goto L65
	}
L48:
	;
	v136 = F_SysCacheGetAttrNotNull(m, int32(19), v117, int32(27))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L12
	} else {
		goto L49
	}
L49:
	;
	v138 = F_pg_detoast_datum(m, v136)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L12
	} else {
		goto L50
	}
L50:
	;
	F_deconstruct_array_builtin(m, v138, int32(26), v28+int32(140), int32(0), v28+int32(136))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L12
	} else {
		goto L51
	}
L51:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v28)+136))
	if v148 <= int32(0) {
		goto L47
	} else {
		goto L52
	}
L52:
	;
	v152 = v63 + int32(36)
	v156 = int32(0)
	goto L53
L53:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v28)+140))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v179+v156<<(uint(int32(2))%32))))
	v184 = F_SearchSysCache1(m, int32(40), v183)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L12
	} else {
		goto L55
	}
L54:
	;
	goto L47
L55:
	;
	if v184 == int32(0) {
		goto L9
	} else {
		goto L56
	}
L56:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v184)+16))
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+22)))
	v190 = v188 + v189
	v193 = F_pstrdup(m, v190+int32(4))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L12
	} else {
		goto L57
	}
L57:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v190)+68))
	v196 = F_get_namespace_name(m, v195)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L12
	} else {
		goto L58
	}
L58:
	;
	v198 = F_makeString(m, v196)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L12
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+128)) = v198
	v201 = F_makeString(m, v193)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L12
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+124)) = v201
	*(*int32)(unsafe.Add(mBase, uint32(v28)+116)) = v201
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v28)+128))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+120)) = v205
	v211 = F_list_make2_impl(m, v28+int32(120), v28+int32(116))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L12
	} else {
		goto L61
	}
L61:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	v214 = F_lappend(m, v213, v211)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L12
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v152))) = v214
	F_ReleaseCatCache(m, v184)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L12
	} else {
		goto L63
	}
L63:
	;
	v220 = v156 + int32(1)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v28)+136))
	if v220 < v221 {
		v156 = v220
		goto L53
	} else {
		goto L64
	}
L64:
	;
	goto L54
L65:
	;
	goto L32
L66:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+135)))
	if v285 == int32(1) {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	v314 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41)+10)))
	if v314 <= int32(0) {
		v669 = v314
		goto L4
	} else {
		goto L76
	}
L68:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v296)+12))
	v310 = v296
	v311 = v301
	v312 = v303
	v313 = v309
	goto L67
L69:
	;
	v307 = int32(0)
	v310 = v307
	v311 = v305
	v312 = v306
	v313 = v307
	goto L67
L70:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v63)+20)) = int64(0)
	v305 = v63 + int32(24)
	v306 = v63 + int32(20)
	goto L69
L71:
	;
	goto L72
L72:
	;
	v294 = F_text_to_cstring(m, v283)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L12
	} else {
		goto L73
	}
L73:
	;
	v296 = F_stringToNode(m, v294)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L12
	} else {
		goto L74
	}
L74:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v63)+20)) = int64(0)
	v301 = v63 + int32(24)
	v303 = v63 + int32(20)
	if v296 != 0 {
		goto L68
	} else {
		goto L75
	}
L75:
	;
	v305 = v301
	v306 = v303
	goto L69
L76:
	;
	v317 = int32(24)
	v324 = int32(0)
	v331 = v313
	goto L77
L77:
	;
	v350 = v324 << (uint(int32(1)) % 32)
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l1)+224))
	v353 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v350+v351))))
	v355 = int32(*(*int16)(unsafe.Add(mBase, uint32(v350+(v41+int32(48))))))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v356)))
	v359 = F_palloc0(m, int32(36))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L12
	} else {
		goto L79
	}
L78:
	;
	v669 = v537
	goto L4
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v359))) = int32(92)
	if v355 != 0 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v410 = F_pstrdup(m, v356+v357<<(uint(int32(4))%32)+v324*int32(100)+int32(24))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L12
	} else {
		goto L93
	}
L81:
	;
	v364 = F_get_attname(m, v42, v355, int32(0))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L12
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	if v331 == int32(0) {
		goto L8
	} else {
		goto L86
	}
L84:
	;
	v366 = F_get_atttype(m, v42, v355)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L12
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v359)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v359)+4)) = v364
	v399 = v331
	v400 = v366
	goto L80
L86:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v310)+4))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v310)+12))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v331)))
	v380 = F_map_variable_attnos(m, v375, int32(1), l2, int32(0), v28+int32(140))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L12
	} else {
		goto L87
	}
L87:
	;
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+140)))
	if v382 == int32(1) {
		goto L7
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v359)+8)) = v380
	v386 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v359)+4)) = v386
	v389 = v331 + int32(4)
	if base.Ui32(v389) < base.Ui32(v374+v373<<(uint(int32(2))%32)) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v395 = v389
	goto L91
L90:
	;
	v395 = v386
	goto L91
L91:
	;
	v396 = F_exprType(m, v380)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L12
	} else {
		goto L92
	}
L92:
	;
	v399 = v395
	v400 = v396
	goto L80
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v359)+12)) = v410
	v413 = int32(0)
	v415 = v324 << (uint(int32(2)) % 32)
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v56+v317+v415)))
	if v417 == v413 {
		v455 = v413
		goto L94
	} else {
		goto L95
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v359)+16)) = v455
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v415+(v60+v317))))
	v462 = F_SearchSysCache1(m, int32(14), v461)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L12
	} else {
		goto L106
	}
L95:
	;
	v420 = F_get_typcollation(m, v400)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L12
	} else {
		goto L96
	}
L96:
	;
	if v420 == v417 {
		v455 = v413
		goto L94
	} else {
		goto L97
	}
L97:
	;
	v424 = F_SearchSysCache1(m, int32(16), v417)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L12
	} else {
		goto L98
	}
L98:
	;
	if v424 == int32(0) {
		goto L6
	} else {
		goto L99
	}
L99:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v424)+16))
	v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v428)+22)))
	v430 = v428 + v429
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v430)+68))
	v432 = F_get_namespace_name(m, v431)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L12
	} else {
		goto L100
	}
L100:
	;
	v436 = F_pstrdup(m, v430+int32(4))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L12
	} else {
		goto L101
	}
L101:
	;
	v438 = F_makeString(m, v432)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L12
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+140)) = v438
	v441 = F_makeString(m, v436)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L12
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+136)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v28)+88)) = v441
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v28)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+92)) = v445
	v451 = F_list_make2_impl(m, v28+int32(92), v28+int32(88))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L12
	} else {
		goto L104
	}
L104:
	;
	F_ReleaseCatCache(m, v424)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L12
	} else {
		goto L105
	}
L105:
	;
	v455 = v451
	goto L94
L106:
	;
	if v462 == int32(0) {
		goto L5
	} else {
		goto L107
	}
L107:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v462)+16))
	v468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v467)+22)))
	v469 = v467 + v468
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v469)+4))
	v471 = F_GetDefaultOpClass(m, v400, v470)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L12
	} else {
		goto L108
	}
L108:
	;
	if v471 != v461 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v469)+72))
	v475 = F_get_namespace_name(m, v474)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L12
	} else {
		goto L112
	}
L110:
	;
	v498 = int32(0)
	goto L111
L111:
	;
	F_ReleaseCatCache(m, v462)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L12
	} else {
		goto L117
	}
L112:
	;
	v479 = F_pstrdup(m, v469+int32(8))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L12
	} else {
		goto L113
	}
L113:
	;
	v481 = F_makeString(m, v475)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L12
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+140)) = v481
	v484 = F_makeString(m, v479)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L12
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+136)) = v484
	*(*int32)(unsafe.Add(mBase, uint32(v28)+72)) = v484
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v28)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+76)) = v488
	v494 = F_list_make2_impl(m, v28+int32(76), v28+int32(72))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L12
	} else {
		goto L116
	}
L116:
	;
	v498 = v494
	goto L111
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v359)+20)) = v498
	v503 = v324 + int32(1)
	v505 = F_get_attoptions(m, v30, base.I32_extend16_s(v503))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L12
	} else {
		goto L118
	}
L118:
	;
	v507 = F_untransformRelOptions(m, v505)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L12
	} else {
		goto L119
	}
L119:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v359)+28)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v359)+24)) = v507
	v512 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	v513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512)+10)))
	if v513 != int32(1) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v312)))
	v534 = F_lappend(m, v533, v359)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L12
	} else {
		goto L128
	}
L121:
	;
	if v353&int32(1) != 0 {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v359)+32)) = v530
	goto L120
L123:
	;
	v518 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v359)+28)) = v518
	if v353&v518 == int32(0) {
		v530 = v518
		goto L122
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	if v353&int32(2) == int32(0) {
		goto L120
	} else {
		goto L127
	}
L126:
	;
	goto L120
L127:
	;
	v530 = int32(1)
	goto L122
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v312))) = v534
	v537 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41)+10)))
	if v503 < v537 {
		v324 = v503
		v331 = v399
		goto L77
	} else {
		goto L129
	}
L129:
	;
	goto L78
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v30
	F_errmsg_internal(m, int32(_a_F_generateClonedIndexStmt_0), v28)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L12
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(_a_F_generateClonedIndexStmt_1), int32(1723), int32(_a_F_generateClonedIndexStmt_2))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L12
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
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v46)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v556
	F_errmsg_internal(m, int32(_a_F_generateClonedIndexStmt_3), v28+int32(16))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L12
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(_a_F_generateClonedIndexStmt_1), int32(1735), int32(_a_F_generateClonedIndexStmt_2))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L12
	} else {
		goto L135
	}
L135:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+96)) = v113
	F_errmsg_internal(m, int32(_a_F_generateClonedIndexStmt_4), v28+int32(96))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L12
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(_a_F_generateClonedIndexStmt_1), int32(1800), int32(_a_F_generateClonedIndexStmt_2))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L12
	} else {
		goto L138
	}
L138:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+112)) = v183
	F_errmsg_internal(m, int32(_a_F_generateClonedIndexStmt_5), v28+int32(112))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L12
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_generateClonedIndexStmt_1), int32(1835), int32(_a_F_generateClonedIndexStmt_2))
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L12
	} else {
		goto L141
	}
L141:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L142:
	;
	F_errmsg_internal(m, int32(_a_F_generateClonedIndexStmt_6), int32(0))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L12
	} else {
		goto L143
	}
L143:
	;
	F_errfinish(m, int32(_a_F_generateClonedIndexStmt_1), int32(1902), int32(_a_F_generateClonedIndexStmt_2))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L12
	} else {
		goto L144
	}
L144:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L145:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L12
	} else {
		goto L146
	}
L146:
	;
	F_errmsg(m, int32(_a_F_generateClonedIndexStmt_7), int32(0))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L12
	} else {
		goto L147
	}
L147:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+32)) = v622 + int32(4)
	F_errdetail(m, int32(_a_F_generateClonedIndexStmt_8), v28+int32(32))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L12
	} else {
		goto L148
	}
L148:
	;
	F_errfinish(m, int32(_a_F_generateClonedIndexStmt_1), int32(1918), int32(_a_F_generateClonedIndexStmt_2))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L12
	} else {
		goto L149
	}
L149:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+80)) = v417
	F_errmsg_internal(m, int32(_a_F_generateClonedIndexStmt_9), v28+int32(80))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L12
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(_a_F_generateClonedIndexStmt_1), int32(2188), int32(_a_F_generateClonedIndexStmt_10))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L12
	} else {
		goto L152
	}
L152:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+48)) = v461
	F_errmsg_internal(m, int32(_a_F_generateClonedIndexStmt_11), v28+int32(48))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L12
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(_a_F_generateClonedIndexStmt_1), int32(2215), int32(_a_F_generateClonedIndexStmt_12))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L12
	} else {
		goto L155
	}
L155:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L156:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L12
	} else {
		goto L188
	}
L157:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L12
	} else {
		goto L184
	}
L158:
	;
	v698 = v669
	goto L161
L159:
	;
	goto L160
L160:
	;
	v787 = F_SysCacheGetAttr(m, int32(57), v34, int32(33), v28+int32(135))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L12
	} else {
		goto L169
	}
L161:
	;
	v723 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41+int32(48)+v698<<(uint(int32(1))%32)))))
	v724 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v724)))
	v727 = F_palloc0(m, int32(36))
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L12
	} else {
		goto L163
	}
L162:
	;
	goto L160
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v727))) = int32(92)
	if v723 == int32(0) {
		goto L157
	} else {
		goto L164
	}
L164:
	;
	v734 = F_get_attname(m, v42, v723, int32(0))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L12
	} else {
		goto L165
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v727)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v727)+4)) = v734
	v747 = F_pstrdup(m, v724+v725<<(uint(int32(4))%32)+v698*int32(100)+int32(24))
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L12
	} else {
		goto L166
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v727)+12)) = v747
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v311)))
	v751 = F_lappend(m, v750, v727)
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L12
	} else {
		goto L167
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v311))) = v751
	v755 = v698 + int32(1)
	v756 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41)+8)))
	if v755 < v756 {
		v698 = v755
		goto L161
	} else {
		goto L168
	}
L168:
	;
	goto L162
L169:
	;
	v789 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+135)))
	if v789 == int32(0) {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v792 = F_untransformRelOptions(m, v787)
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L12
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	v799 = F_SysCacheGetAttr(m, int32(34), v38, int32(21), v28+int32(135))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L12
	} else {
		goto L174
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v63)+28)) = v792
	goto L172
L174:
	;
	v801 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+135)))
	if v801 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v804 = F_text_to_cstring(m, v799)
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L12
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	F_ReleaseCatCache(m, v34)
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L12
	} else {
		goto L182
	}
L178:
	;
	v806 = F_stringToNode(m, v804)
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L12
	} else {
		goto L179
	}
L179:
	;
	v812 = F_map_variable_attnos(m, v806, int32(1), l2, int32(0), v28+int32(140))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L12
	} else {
		goto L180
	}
L180:
	;
	v814 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+140)))
	if v814 == int32(1) {
		goto L156
	} else {
		goto L181
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v63)+32)) = v812
	goto L177
L182:
	;
	F_ReleaseCatCache(m, v48)
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L12
	} else {
		goto L183
	}
L183:
	;
	m.G0 = v28 + int32(144)
	return v63
L184:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L12
	} else {
		goto L185
	}
L185:
	;
	F_errmsg(m, int32(_a_F_generateClonedIndexStmt_13), int32(0))
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L12
	} else {
		goto L186
	}
L186:
	;
	F_errfinish(m, int32(_a_F_generateClonedIndexStmt_1), int32(1988), int32(_a_F_generateClonedIndexStmt_2))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L12
	} else {
		goto L187
	}
L187:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L188:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L12
	} else {
		goto L189
	}
L189:
	;
	F_errmsg(m, int32(_a_F_generateClonedIndexStmt_7), int32(0))
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L12
	} else {
		goto L190
	}
L190:
	;
	v854 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+64)) = v854 + int32(4)
	F_errdetail(m, int32(_a_F_generateClonedIndexStmt_8), v28-int32(-64))
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L12
	} else {
		goto L191
	}
L191:
	;
	F_errfinish(m, int32(_a_F_generateClonedIndexStmt_1), int32(2026), int32(_a_F_generateClonedIndexStmt_2))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L12
	} else {
		goto L192
	}
L192:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_generate_combinations_recurse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if l1 < v6 {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v8 <= l2 {
		} else {
			v11 = l1 + int32(1)
			v17 = l2
			for {
				*(*int32)(unsafe.Add(mBase, uint32(l3+l1<<(uint(int32(2))%32)))) = v17
				v22 = v17 + int32(1)
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v11 < v24 {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					if v26 <= v22 {
					} else {
						v82 = int32(2)
						v35 = v22
						for {
							*(*int32)(unsafe.Add(mBase, uint32(l3+v11<<(uint(v82)%32)))) = v35
							v40 = v35 + int32(1)
							F_generate_combinations_recurse(m, l0, l1+v82, v40, l3)
							mBase = m.M
							v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							if v40 < v42 {
								v35 = v40
								continue
							} else {
								break
							}
							break
						}
					}
				} else {
					v45 = v24 << (uint(int32(2)) % 32)
					if v45 != 0 {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						base.MemoryCopy(m, v46+v47*v24<<(uint(int32(2))%32), l3, v45)
					} else {
					}
					v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v53 + int32(1)
				}
				v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v22 < v62 {
					v17 = v22
					continue
				} else {
					break
				}
				break
			}
		}
	} else {
		v65 = v6 << (uint(int32(2)) % 32)
		if v65 != 0 {
			v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			base.MemoryCopy(m, v66+v67*v6<<(uint(int32(2))%32), l3, v65)
		} else {
		}
		v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v73 + int32(1)
	}
	return
}
func F_generate_implied_equalities_for_column(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v85 int32
	_ = v85
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v145 int32
	_ = v145
	var v154 int32
	_ = v154
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v516 int32
	_ = v516
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v557 int32
	_ = v557
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v580 int32
	_ = v580
	var v613 int32
	_ = v613
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v669 int32
	_ = v669
	var v678 int32
	_ = v678
	v6 = int32(0)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v20 == int32(2) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v23 = F_find_childrel_parents(m, l0, l1)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v27 = v6
	goto L3
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+136))
	if v28 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	return int32(0)
L5:
	;
	v27 = v23
	goto L3
L6:
	;
	return v678
L7:
	;
	if v85 < int32(0) {
		v678 = v6
		goto L6
	} else {
		goto L18
	}
L8:
	;
	v85 = base.I32_ctz(v71) | v72<<(uint(int32(5))%32)
	goto L7
L9:
	;
	v85 = int32(-2)
	goto L7
L10:
	;
	v38 = base.I32_div_s(int32(0), int32(32))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v39 <= v38 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v42 = v28 + int32(8)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42+v38<<(uint(int32(2))%32))))
	v49 = v46 & int32(-1)
	if v49 != 0 {
		v71 = v49
		v72 = v38
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v51 = v38 + int32(1)
	if v51 == v39 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v54 = v51
	goto L14
L14:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v42+v54<<(uint(int32(2))%32))))
	if v61 != 0 {
		v71 = v61
		v72 = v54
		goto L8
	} else {
		goto L16
	}
L15:
	;
	goto L9
L16:
	;
	v63 = v54 + int32(1)
	if v63 != v39 {
		v54 = v63
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v100 = v85
	goto L19
L19:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+12))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v108+v100<<(uint(int32(2))%32))))
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+40)))
	if v113 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v678 = int32(0)
	goto L6
L21:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(l1)+136))
	if v613 == int32(0) {
		goto L141
	} else {
		goto L142
	}
L22:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v112)+16))
	if v114 == int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	if v117 < int32(2) {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v112)+20))
	if v122 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v123 = v120
	goto L27
L26:
	;
	v123 = int32(0)
	goto L27
L27:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	v131 = v114
	v132 = v124
	v134 = int32(-1)
	goto L28
L28:
	;
	if v132 != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v112)+16))
	if v324 == int32(0) {
		goto L21
	} else {
		goto L68
	}
L30:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v239)))
	if v253 == int32(0) {
		goto L21
	} else {
		goto L50
	}
L31:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v131)+12))
	v238 = v131
	v239 = v132
	v241 = v134
	v252 = v145
	goto L30
L32:
	;
	goto L33
L33:
	;
	v154 = v134
	goto L34
L34:
	;
	if v123 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v229)+12))
	v238 = v229
	v239 = v232
	v241 = v220
	v252 = v232
	goto L30
L36:
	;
	if v220 <= int32(0) {
		goto L21
	} else {
		goto L47
	}
L37:
	;
	v220 = base.I32_ctz(v206) | v207<<(uint(int32(5))%32)
	goto L36
L38:
	;
	v220 = int32(-2)
	goto L36
L39:
	;
	v171 = v154 + int32(1)
	v173 = base.I32_div_s(v171, int32(32))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	if v174 <= v173 {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v177 = v123 + int32(8)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v177+v173<<(uint(int32(2))%32))))
	v184 = v181 & (int32(-1) << (uint(v171) % 32))
	if v184 != 0 {
		v206 = v184
		v207 = v173
		goto L37
	} else {
		goto L41
	}
L41:
	;
	v186 = v173 + int32(1)
	if v186 == v174 {
		goto L38
	} else {
		goto L42
	}
L42:
	;
	v189 = v186
	goto L43
L43:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v177+v189<<(uint(int32(2))%32))))
	if v196 != 0 {
		v206 = v196
		v207 = v189
		goto L37
	} else {
		goto L45
	}
L44:
	;
	goto L38
L45:
	;
	v198 = v189 + int32(1)
	if v198 != v174 {
		v189 = v198
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
	if v223 <= v220 {
		goto L21
	} else {
		goto L48
	}
L48:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v112)+20))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v225+v220<<(uint(int32(2))%32))))
	if v229 == int32(0) {
		v154 = v220
		goto L34
	} else {
		goto L49
	}
L49:
	;
	goto L35
L50:
	;
	v257 = v239 + int32(4)
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v238)+4))
	if base.Ui32(v257) < base.Ui32(v252+v259<<(uint(int32(2))%32)) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v264 = v257
	goto L53
L52:
	;
	v264 = int32(0)
	goto L53
L53:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v253)+8))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v267 = int32(0)
	if base.B2i32(v265 == v267)|base.B2i32(v266 == v267) != 0 {
		v313 = base.B2i32(v265|v266 == v267)
		goto L55
	} else {
		goto L56
	}
L54:
	;
	if v313 == int32(0) {
		v131 = v238
		v132 = v264
		v134 = v241
		goto L28
	} else {
		goto L65
	}
L55:
	;
	goto L54
L56:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v265)+4))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v266)+4))
	if v281 != v282 {
		v313 = int32(0)
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v284 = int32(1)
	if v281 <= v284 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v287 = v284
	goto L60
L59:
	;
	v287 = v281
	goto L60
L60:
	;
	v288 = int32(8)
	v293 = int32(0)
	goto L61
L61:
	;
	v301 = v293 << (uint(int32(2)) % 32)
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v265+v288+v301)))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v266+v288+v301)))
	v306 = base.B2i32(v303 == v305)
	if v303 != v305 {
		v313 = v306
		goto L55
	} else {
		goto L63
	}
L62:
	;
	v313 = v306
	goto L55
L63:
	;
	v309 = v293 + int32(1)
	if v309 != v287 {
		v293 = v309
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	v320 = m.T0[l2].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, l0, l1, v112, v253, l3)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L4
	} else {
		goto L66
	}
L66:
	;
	if v320 == int32(0) {
		v131 = v238
		v132 = v264
		v134 = v241
		goto L28
	} else {
		goto L67
	}
L67:
	;
	goto L29
L68:
	;
	v327 = int32(0)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v324)+4))
	if v327 < v329 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v337 = v327
	v340 = v327
	goto L72
L70:
	;
	v580 = v327
	goto L71
L71:
	;
	if v580 != 0 {
		v678 = v580
		goto L6
	} else {
		goto L138
	}
L72:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v324)+12))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v351+v340<<(uint(int32(2))%32))))
	if v355 == v253 {
		v557 = v337
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v580 = v557
	goto L71
L74:
	;
	v572 = v340 + int32(1)
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v324)+4))
	if v572 < v573 {
		v337 = v557
		v340 = v572
		goto L72
	} else {
		goto L137
	}
L75:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v355)+8))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v359 = int32(0)
	if base.B2i32(v357 == v359)|base.B2i32(v358 == v359) != 0 {
		v404 = v359
		goto L77
	} else {
		goto L78
	}
L76:
	;
	if v404 != 0 {
		v557 = v337
		goto L74
	} else {
		goto L89
	}
L77:
	;
	goto L76
L78:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v357)+4))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v358)+4))
	if v369 < v370 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v372 = v369
	goto L81
L80:
	;
	v372 = v370
	goto L81
L81:
	;
	if v372 <= int32(1) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v375 = int32(1)
	goto L84
L83:
	;
	v375 = v372
	goto L84
L84:
	;
	v376 = int32(8)
	v381 = int32(0)
	goto L85
L85:
	;
	v388 = v381 << (uint(int32(2)) % 32)
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v358+v376+v388)))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v357+v376+v388)))
	v393 = v390 & v392
	v395 = base.B2i32(v393 != int32(0))
	if v393 != 0 {
		v404 = v395
		goto L77
	} else {
		goto L87
	}
L86:
	;
	v404 = v395
	goto L77
L87:
	;
	v397 = v381 + int32(1)
	if v397 != v375 {
		v381 = v397
		goto L85
	} else {
		goto L88
	}
L88:
	;
	goto L86
L89:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v355)+8))
	v406 = int32(0)
	if base.B2i32(v405 == v406)|base.B2i32(l4 == v406) != 0 {
		v451 = v406
		goto L91
	} else {
		goto L92
	}
L90:
	;
	if v451 != 0 {
		v557 = v337
		goto L74
	} else {
		goto L103
	}
L91:
	;
	goto L90
L92:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v405)+4))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v416 < v417 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v419 = v416
	goto L95
L94:
	;
	v419 = v417
	goto L95
L95:
	;
	if v419 <= int32(1) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v422 = int32(1)
	goto L98
L97:
	;
	v422 = v419
	goto L98
L98:
	;
	v423 = int32(8)
	v428 = int32(0)
	goto L99
L99:
	;
	v435 = v428 << (uint(int32(2)) % 32)
	v437 = *(*int32)(unsafe.Add(mBase, uint32(l4+v423+v435)))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v405+v423+v435)))
	v440 = v437 & v439
	v442 = base.B2i32(v440 != int32(0))
	if v440 != 0 {
		v451 = v442
		goto L91
	} else {
		goto L101
	}
L100:
	;
	v451 = v442
	goto L91
L101:
	;
	v444 = v428 + int32(1)
	if v444 != v422 {
		v428 = v444
		goto L99
	} else {
		goto L102
	}
L102:
	;
	goto L100
L103:
	;
	if v20 == int32(2) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v355)+8))
	v455 = int32(0)
	if base.B2i32(v27 == v455)|base.B2i32(v454 == v455) != 0 {
		v500 = v455
		goto L108
	} else {
		goto L109
	}
L105:
	;
	goto L106
L106:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	if v501 == int32(0) {
		v557 = v337
		goto L74
	} else {
		goto L121
	}
L107:
	;
	if v500 != 0 {
		v557 = v337
		goto L74
	} else {
		goto L120
	}
L108:
	;
	goto L107
L109:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v454)+4))
	if v465 < v466 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v468 = v465
	goto L112
L111:
	;
	v468 = v466
	goto L112
L112:
	;
	if v468 <= int32(1) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v471 = int32(1)
	goto L115
L114:
	;
	v471 = v468
	goto L115
L115:
	;
	v472 = int32(8)
	v477 = int32(0)
	goto L116
L116:
	;
	v484 = v477 << (uint(int32(2)) % 32)
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v454+v472+v484)))
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v27+v472+v484)))
	v489 = v486 & v488
	v491 = base.B2i32(v489 != int32(0))
	if v489 != 0 {
		v500 = v491
		goto L108
	} else {
		goto L118
	}
L117:
	;
	v500 = v491
	goto L108
L118:
	;
	v493 = v477 + int32(1)
	if v493 != v471 {
		v477 = v493
		goto L116
	} else {
		goto L119
	}
L119:
	;
	goto L117
L120:
	;
	goto L106
L121:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v501)+4))
	if v504 <= int32(0) {
		v557 = v337
		goto L74
	} else {
		goto L122
	}
L122:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v355)+16))
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v253)+16))
	v516 = int32(0)
	goto L123
L123:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v501)+12))
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v529+v516<<(uint(int32(2))%32))))
	v535 = F_get_opfamily_member_for_cmptype(m, v533, v508, v507, int32(3))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L4
	} else {
		goto L126
	}
L124:
	;
	v548 = F_create_join_clause(m, l0, v112, v535, v253, v355, v112)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L4
	} else {
		goto L135
	}
L125:
	;
	goto L124
L126:
	;
	if v535 != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v112)+52))
	if v537 == int32(0) {
		goto L125
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	v545 = v516 + int32(1)
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v501)+4))
	if v545 < v546 {
		v516 = v545
		goto L123
	} else {
		goto L134
	}
L130:
	;
	v540 = F_get_opcode(m, v535)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L4
	} else {
		goto L131
	}
L131:
	;
	v542 = F_get_func_leakproof(m, v540)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L4
	} else {
		goto L132
	}
L132:
	;
	if v542 != 0 {
		goto L125
	} else {
		goto L133
	}
L133:
	;
	goto L129
L134:
	;
	v557 = v337
	goto L74
L135:
	;
	v550 = F_lappend(m, v337, v548)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L4
	} else {
		goto L136
	}
L136:
	;
	v557 = v550
	goto L74
L137:
	;
	goto L73
L138:
	;
	goto L21
L139:
	;
	if int32(0) <= v669 {
		v100 = v669
		goto L19
	} else {
		goto L150
	}
L140:
	;
	v669 = base.I32_ctz(v655) | v656<<(uint(int32(5))%32)
	goto L139
L141:
	;
	v669 = int32(-2)
	goto L139
L142:
	;
	v620 = v100 + int32(1)
	v622 = base.I32_div_s(v620, int32(32))
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v613)+4))
	if v623 <= v622 {
		goto L141
	} else {
		goto L143
	}
L143:
	;
	v626 = v613 + int32(8)
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v626+v622<<(uint(int32(2))%32))))
	v633 = v630 & (int32(-1) << (uint(v620) % 32))
	if v633 != 0 {
		v655 = v633
		v656 = v622
		goto L140
	} else {
		goto L144
	}
L144:
	;
	v635 = v622 + int32(1)
	if v635 == v623 {
		goto L141
	} else {
		goto L145
	}
L145:
	;
	v638 = v635
	goto L146
L146:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v626+v638<<(uint(int32(2))%32))))
	if v645 != 0 {
		v655 = v645
		v656 = v638
		goto L140
	} else {
		goto L148
	}
L147:
	;
	goto L141
L148:
	;
	v647 = v638 + int32(1)
	if v647 != v623 {
		v638 = v647
		goto L146
	} else {
		goto L149
	}
L149:
	;
	goto L147
L150:
	;
	goto L20
}
func F_generate_partitionwise_join_paths(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
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
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	v3 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v8 - int32(1) {
	case 0, 2:
		goto L3
	default:
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+236)) = int32(0)
	return
L2:
	;
	return
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+232))
	if v11 == int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+240))
	if v14 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+236))
	if v17 <= int32(0) {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+252))
	if v20 == int32(0) {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v23 = int32(0)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v25 == v23 {
		v46 = v23
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v46 != 0 {
		goto L2
	} else {
		goto L18
	}
L9:
	;
	goto L8
L10:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v29 = v28
	goto L11
L11:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if base.Ui32(int32(2)) <= base.Ui32(v33-int32(301)) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v46 = int32(1)
	goto L9
L13:
	;
	if v33 != int32(290) {
		v46 = v23
		goto L9
	} else {
		goto L16
	}
L14:
	;
	v29 = v32 + int32(72)
	goto L11
L15:
	;
	goto L12
L16:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v32)+72))
	if v40 != 0 {
		v46 = v23
		goto L9
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return
L20:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+236))
	if int32(0) < v49 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	F_add_paths_to_append_rel(m, l0, l1, v99)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L19
	} else {
		goto L47
	}
L22:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+252))
	v56 = v3
	v57 = v3
	goto L25
L23:
	;
	goto L24
L24:
	;
	F_mark_dummy_rel(m, l1)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L19
	} else {
		goto L46
	}
L25:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v52+v57<<(uint(int32(2))%32))))
	if v63 == int32(0) {
		v99 = v56
		goto L27
	} else {
		goto L28
	}
L26:
	;
	if v99 != 0 {
		goto L21
	} else {
		goto L45
	}
L27:
	;
	v101 = v57 + int32(1)
	if v101 != v49 {
		v56 = v99
		v57 = v101
		goto L25
	} else {
		goto L44
	}
L28:
	;
	F_generate_partitionwise_join_paths(m, l0, v63)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L19
	} else {
		goto L29
	}
L29:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v63)+32))
	if v68 == int32(0) {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_set_cheapest(m, v63)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L19
	} else {
		goto L31
	}
L31:
	;
	v73 = int32(0)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v63)+32))
	if v75 == v73 {
		v96 = v73
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if v96 != 0 {
		v99 = v56
		goto L27
	} else {
		goto L42
	}
L33:
	;
	goto L32
L34:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
	v79 = v78
	goto L35
L35:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	if base.Ui32(int32(2)) <= base.Ui32(v83-int32(301)) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v96 = int32(1)
	goto L33
L37:
	;
	if v83 != int32(290) {
		v96 = v73
		goto L33
	} else {
		goto L40
	}
L38:
	;
	v79 = v82 + int32(72)
	goto L35
L39:
	;
	goto L36
L40:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v82)+72))
	if v90 != 0 {
		v96 = v73
		goto L33
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v97 = F_lappend(m, v56, v63)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L19
	} else {
		goto L43
	}
L43:
	;
	v99 = v97
	goto L27
L44:
	;
	goto L26
L45:
	;
	goto L24
L46:
	;
	return
L47:
	;
	F_list_free(m, v99)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L19
	} else {
		goto L48
	}
L48:
	;
	goto L2
}
func F_geqo_rand(m *base.Module, l0 int32) float64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v30 float64
	_ = v30
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+368))
	v4 = v2 + int32(8)
	v7 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(v4)+8))
	v9 = v7 ^ v8
	*(*int64)(unsafe.Add(mBase, uint32(v4)+8)) = base.I64_rotl(v9, int64(37))
	*(*int64)(unsafe.Add(mBase, uint32(v4))) = v9<<(uint(int64(16))%64) ^ base.I64_rotl(v7, int64(24)) ^ v9
	v30 = F_scalbn(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v7*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
	mBase = m.M
	return v30
}
func F_getTokenTypes(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	v3 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = F_lookup_ts_parser_cache(m, l0)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if l1 == int32(0) {
		v200 = v3
		goto L5
	} else {
		goto L6
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L50
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L47
	}
L5:
	;
	m.G0 = v14 + int32(32)
	return v200
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v22 == int32(0) {
		v200 = v3
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	if v25 == int32(0) {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v28 = int32(0)
	v30 = F_OidFunctionCall1Coll(m, v25, v28, v28)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v32 <= int32(0) {
		v200 = v3
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v39 = v3
	v40 = v32
	v41 = v3
	goto L11
L11:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46+v41<<(uint(int32(2))%32))))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v39 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v200 = v186
	goto L5
L13:
	;
	v194 = v41 + int32(1)
	if v194 < v187 {
		v39 = v186
		v40 = v187
		v41 = v194
		goto L11
	} else {
		goto L46
	}
L14:
	;
	if v30 == int32(0) {
		goto L3
	} else {
		goto L28
	}
L15:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v54 <= int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v59 = int32(0)
	goto L17
L17:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v57+v59<<(uint(int32(2))%32))))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	if base.B2i32(v77 == int32(0))|base.B2i32(v77 != v80) != 0 {
		v98 = v77
		v99 = v80
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L14
L19:
	;
	if v98-v99 == int32(0) {
		v186 = v39
		v187 = v40
		goto L13
	} else {
		goto L26
	}
L20:
	;
	goto L19
L21:
	;
	v83 = v51
	v84 = v74
	goto L22
L22:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+1)))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+1)))
	if v88 == int32(0) {
		v98 = v88
		v99 = v87
		goto L20
	} else {
		goto L24
	}
L23:
	;
	v98 = v88
	v99 = v87
	goto L20
L24:
	;
	v91 = int32(1)
	if v88 == v87 {
		v83 = v83 + v91
		v84 = v84 + v91
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v104 = v59 + int32(1)
	if v54 != v104 {
		v59 = v104
		goto L17
	} else {
		goto L27
	}
L27:
	;
	goto L18
L28:
	;
	v119 = int32(0)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if v120 == v119 {
		goto L3
	} else {
		goto L29
	}
L29:
	;
	v123 = v119
	goto L30
L30:
	;
	v136 = v30 + v123*int32(12)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
	if base.B2i32(v140 == int32(0))|base.B2i32(v140 != v143) != 0 {
		v161 = v140
		v162 = v143
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v171 = F_palloc0(m, int32(8))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L43
	}
L32:
	;
	if v161-v162 != 0 {
		goto L39
	} else {
		goto L40
	}
L33:
	;
	goto L32
L34:
	;
	v146 = v51
	v147 = v137
	goto L35
L35:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+1)))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+1)))
	if v151 == int32(0) {
		v161 = v151
		v162 = v150
		goto L33
	} else {
		goto L37
	}
L36:
	;
	v161 = v151
	v162 = v150
	goto L33
L37:
	;
	v154 = int32(1)
	if v151 == v150 {
		v146 = v146 + v154
		v147 = v147 + v154
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v165 = v123 + int32(1)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v30+v165*int32(12))))
	if v169 != 0 {
		v123 = v165
		goto L30
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	goto L31
L42:
	;
	goto L3
L43:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	*(*int32)(unsafe.Add(mBase, uint32(v171))) = v173
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	v176 = F_pstrdup(m, v175)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v171)+4)) = v176
	v179 = F_lappend(m, v39, v171)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v186 = v179
	v187 = v181
	goto L13
L46:
	;
	goto L12
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = l0
	F_errmsg_internal(m, int32(_a_F_getTokenTypes_0), v14)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_getTokenTypes_1), int32(1243), int32(_a_F_getTokenTypes_2))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v242
	F_errmsg(m, int32(_a_F_getTokenTypes_3), v14+int32(16))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_getTokenTypes_1), int32(1278), int32(_a_F_getTokenTypes_2))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_actual_clauses(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v2 = int32(0)
	if l0 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v8 <= int32(0) {
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
	v14 = v2
	v15 = v2
	goto L7
L7:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16+v15<<(uint(int32(2))%32))))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v22 = F_lappend(m, v14, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	return v22
L9:
	;
	return int32(0)
L10:
	;
	v27 = v15 + int32(1)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v27 < v28 {
		v14 = v22
		v15 = v27
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
}
func F_get_actual_variable_range(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
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
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
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
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v178 int32
	_ = v178
	v7 = int32(0)
	v12 = m.G0
	v14 = v12 + int32(-64)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v16 == v7 {
		v178 = v7
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v14 - int32(-64)
	return v178
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)+108))
	if v19 == int32(0) {
		v178 = v7
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)+68))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22+v23<<(uint(int32(2))%32))))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+21)))
	if v28 == int32(112) {
		v178 = v7
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v31 <= int32(0) {
		v178 = v7
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v42 = v7
	goto L6
L6:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45+v42<<(uint(int32(2))%32))))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+60))
	if v50 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v178 = int32(0)
	goto L1
L8:
	;
	v166 = v42 + int32(1)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v166 < v167 {
		v42 = v166
		goto L6
	} else {
		goto L46
	}
L9:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49)+88))
	if v53 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+104)))
	if v54 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v49)+76))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v56 != int32(1) {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v49)+48))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if l3 != v60 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v64 = F_match_index_to_operand(m, v62, int32(0), v49)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(0)
L15:
	;
	if v64 == int32(0) {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v49)+60))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v72 = F_get_op_opfamily_strategy(m, l2, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L14
	} else {
		goto L20
	}
L17:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_get_actual_variable_range[0]))
	v100 = F_AllocSetContextCreateInternal(m, v95, int32(_a_F_get_actual_variable_range_0), int32(0), int32(_a_F_get_actual_variable_range_1), int32(_a_F_get_actual_variable_range_2))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L14
	} else {
		goto L28
	}
L18:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v49)+64))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	if v91 != 0 {
		goto L25
	} else {
		goto L26
	}
L19:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v49)+64))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	if v86 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v49)+80))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v49)+60))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v79 = F_IndexAmTranslateStrategy(m, v72&int32(_a_F_get_actual_variable_range_3), v76, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L14
	} else {
		goto L21
	}
L21:
	;
	switch v79 - int32(1) {
	case 0:
		goto L19
	default:
		goto L8
	case 4:
		goto L18
	}
L22:
	;
	v87 = int32(-1)
	goto L24
L23:
	;
	v87 = int32(1)
	goto L24
L24:
	;
	v93 = v87
	goto L17
L25:
	;
	v92 = int32(1)
	goto L27
L26:
	;
	v92 = int32(-1)
	goto L27
L27:
	;
	v93 = v92
	goto L17
L28:
	;
	v102 = int32(_a_F_get_actual_variable_range_4)
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_get_actual_variable_range[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_get_actual_variable_range[0])) = v100
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	v108 = F_table_open(m, v106, int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L14
	} else {
		goto L29
	}
L29:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	v112 = F_index_open(m, v110, int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L14
	} else {
		goto L30
	}
L30:
	;
	v115 = F_table_slot_create(m, v108, int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L14
	} else {
		goto L31
	}
L31:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	F_get_typlenbyval(m, v117, v12+int32(-2), v12+int32(-3))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L14
	} else {
		goto L32
	}
L32:
	;
	v124 = int32(1)
	v127 = int32(0)
	F_ScanKeyEntryInitialize(m, v14, int32(129), v124, v127, v127, v127, v127, v127)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L14
	} else {
		goto L33
	}
L33:
	;
	if l4 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v134 = int32(*(*int16)(unsafe.Add(mBase, uint32(v14)+62)))
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+61)))
	v136 = F_get_actual_variable_endpoint(m, v108, v112, v93, v14, v134, v135, v115, v103, l4)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L14
	} else {
		goto L37
	}
L35:
	;
	v138 = v124
	goto L36
L36:
	;
	v139 = int32(0)
	if base.B2i32(l5 == v139)|base.B2i32(v138 == v139) == v139 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v138 = v136
	goto L36
L38:
	;
	v148 = int32(*(*int16)(unsafe.Add(mBase, uint32(v14)+62)))
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+61)))
	v150 = F_get_actual_variable_endpoint(m, v108, v112, int32(0)-v93, v14, v148, v149, v115, v103, l5)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L14
	} else {
		goto L41
	}
L39:
	;
	v152 = v138
	goto L40
L40:
	;
	F_ExecDropSingleTupleTableSlot(m, v115)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L14
	} else {
		goto L42
	}
L41:
	;
	v152 = v150
	goto L40
L42:
	;
	F_relation_close(m, v112, int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L14
	} else {
		goto L43
	}
L43:
	;
	F_relation_close(m, v108, int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L14
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_get_actual_variable_range[0])) = v103
	F_MemoryContextDelete(m, v100)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L14
	} else {
		goto L45
	}
L45:
	;
	v178 = v152
	goto L1
L46:
	;
	goto L7
}
func F_get_attavgwidth(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_get_attavgwidth[0]))
	if v5 != 0 {
		v6 = m.T0[v5].(func(*base.Module, int32, int32) int32)(m, l0, l1)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			if int32(0) < v6 {
				v30 = v6
				return v30
			} else {
				v15 = F_SearchSysCache3(m, int32(65), l0, l1, int32(0))
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int32(0)
				} else {
					if v15 != 0 {
						v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
						v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+22)))
						v20 = *(*int32)(unsafe.Add(mBase, uint32(v17+v18)+12))
						F_ReleaseCatCache(m, v15)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int32(0)
						} else {
							if int32(0) < v20 {
								v30 = v20
							} else {
								v30 = int32(0)
							}
							return v30
						}
					} else {
						v30 = int32(0)
						return v30
					}
				}
			}
		}
	} else {
		v15 = F_SearchSysCache3(m, int32(65), l0, l1, int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			if v15 != 0 {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
				v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+22)))
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v17+v18)+12))
				F_ReleaseCatCache(m, v15)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					if int32(0) < v20 {
						v30 = v20
					} else {
						v30 = int32(0)
					}
					return v30
				}
			} else {
				v30 = int32(0)
				return v30
			}
		}
	}
}
func F_get_baserel_parampathinfo(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
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
	var v153 int32
	_ = v153
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 float64
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 float64
	_ = v219
	var v220 int32
	_ = v220
	var v221 float64
	_ = v221
	var v230 float64
	_ = v230
	var v234 float64
	_ = v234
	var v235 float64
	_ = v235
	var v237 float64
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v256 int32
	_ = v256
	v4 = int32(0)
	if l2 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v16 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return v256
L5:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v106 = F_bms_union(m, v105, l2)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L23
	} else {
		goto L24
	}
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v19 <= int32(0) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v26 = v4
	goto L8
L8:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33+v26<<(uint(int32(2))%32))))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v39 = int32(0)
	if base.B2i32(v38 == v39)|base.B2i32(l2 == v39) != 0 {
		v85 = base.B2i32(v38|l2 == v39)
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L5
L10:
	;
	if v85 != 0 {
		v256 = v37
		goto L4
	} else {
		goto L21
	}
L11:
	;
	goto L10
L12:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v53 != v54 {
		v85 = int32(0)
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v56 = int32(1)
	if v53 <= v56 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v59 = v56
	goto L16
L15:
	;
	v59 = v53
	goto L16
L16:
	;
	v60 = int32(8)
	v65 = int32(0)
	goto L17
L17:
	;
	v73 = v65 << (uint(int32(2)) % 32)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v38+v60+v73)))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l2+v60+v73)))
	v78 = base.B2i32(v75 == v77)
	if v75 != v77 {
		v85 = v78
		goto L11
	} else {
		goto L19
	}
L18:
	;
	v85 = v78
	goto L11
L19:
	;
	v81 = v65 + int32(1)
	if v81 != v59 {
		v65 = v81
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v91 = v26 + int32(1)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v91 < v92 {
		v26 = v91
		goto L8
	} else {
		goto L22
	}
L22:
	;
	goto L9
L23:
	;
	return int32(0)
L24:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+212))
	if v110 == int32(0) {
		v161 = v4
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v167 = F_generate_join_implied_equalities(m, l0, v106, l2, l1, int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L23
	} else {
		goto L40
	}
L26:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	if v113 <= int32(0) {
		v161 = v4
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v121 = int32(0)
	v123 = v4
	goto L28
L28:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v110)+12))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v128+v121<<(uint(int32(2))%32))))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v134 = int32(0)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v132)+28))
	v136 = F_bms_is_subset(m, v135, v106)
	mBase = m.M
	if v136 == v134 {
		v147 = v134
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v161 = v150
	goto L25
L30:
	;
	if v147 != 0 {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	goto L30
L32:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v132)+28))
	v140 = F_bms_overlap(m, v133, v139)
	mBase = m.M
	if v140 == int32(0) {
		v147 = v134
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v132)+40))
	v144 = F_bms_overlap(m, v133, v143)
	mBase = m.M
	v147 = v144 ^ int32(1)
	goto L31
L34:
	;
	v148 = F_lappend(m, v123, v132)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L23
	} else {
		goto L37
	}
L35:
	;
	v150 = v123
	goto L36
L36:
	;
	v152 = v121 + int32(1)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	if v152 < v153 {
		v121 = v152
		v123 = v150
		goto L28
	} else {
		goto L38
	}
L37:
	;
	v150 = v148
	goto L36
L38:
	;
	goto L29
L39:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
	v212 = F_list_concat_copy(m, v169, v211)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L23
	} else {
		goto L48
	}
L40:
	;
	v169 = F_list_concat(m, v161, v167)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L23
	} else {
		goto L41
	}
L41:
	;
	if v169 == int32(0) {
		v207 = v4
		goto L39
	} else {
		goto L42
	}
L42:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v169)+4))
	if v173 <= int32(0) {
		v207 = v4
		goto L39
	} else {
		goto L43
	}
L43:
	;
	v181 = int32(0)
	v184 = v4
	goto L44
L44:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v169)+12))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v188+v181<<(uint(int32(2))%32))))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+56))
	v194 = F_bms_add_member(m, v184, v193)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L23
	} else {
		goto L46
	}
L45:
	;
	v207 = v194
	goto L39
L46:
	;
	v197 = v181 + int32(1)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v169)+4))
	if v197 < v198 {
		v181 = v197
		v184 = v194
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v215 = *(*float64)(unsafe.Add(mBase, uint32(l1)+120))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v217 = int32(0)
	v219 = F_clauselist_selectivity(m, l0, v212, v216, v217, v217)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L23
	} else {
		goto L50
	}
L49:
	;
	v235 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	if base.F64_gt(v234, v235) != 0 {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	v221 = base.F64_mul(v215, v219)
	if base.F64_gt(v221, float64(1e+100))|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v221)&int64(9223372036854775807))) != 0 {
		v234 = float64(1e+100)
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v230 = float64(1)
	if base.F64_le(v221, v230) != 0 {
		v234 = v230
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v234 = base.F64_nearest(v221)
	goto L49
L53:
	;
	v237 = v235
	goto L55
L54:
	;
	v237 = v234
	goto L55
L55:
	;
	v239 = F_palloc0(m, int32(24))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L23
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v239)+20)) = v207
	*(*int32)(unsafe.Add(mBase, uint32(v239)+16)) = v169
	*(*float64)(unsafe.Add(mBase, uint32(v239)+8)) = v237
	*(*int32)(unsafe.Add(mBase, uint32(v239)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v239))) = int32(278)
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v248 = F_lappend(m, v247, v239)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L23
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v248
	v256 = v239
	goto L4
}
func F_get_cheapest_fractional_path(m *base.Module, l0 int32, l1 float64) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v15 float64
	_ = v15
	var v21 float64
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 float64
	_ = v60
	var v61 float64
	_ = v61
	var v65 float64
	_ = v65
	var v66 float64
	_ = v66
	var v72 float64
	_ = v72
	var v73 float64
	_ = v73
	var v76 float64
	_ = v76
	var v77 float64
	_ = v77
	var v78 float64
	_ = v78
	var v81 float64
	_ = v81
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if base.F64_le(l1, float64(0)) != 0 {
		v101 = v8
	} else {
		if base.F64_ge(l1, float64(1)) == int32(0) {
			v21 = l1
		} else {
			v15 = *(*float64)(unsafe.Add(mBase, uint32(v8)+32))
			if base.F64_gt(v15, float64(0)) == int32(0) {
				v21 = l1
			} else {
				v21 = base.F64_div(l1, v15)
			}
		}
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		if v23 == int32(0) {
			v101 = v8
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
			if v26 <= int32(0) {
				v101 = v8
			} else {
				v31 = v8
				v34 = int32(0)
				for {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v36+v34<<(uint(int32(2))%32))))
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
					if v41 != 0 {
						v94 = v31
					} else {
						v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						if v40 == v42 {
							v94 = v31
						} else {
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v31)+40))
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v40)+40))
							if v47 != v48 {
								if v47 < v48 {
									v53 = int32(-1)
								} else {
									v53 = int32(1)
								}
								v90 = v53
							} else {
								if base.F64_le(v21, float64(0))|base.F64_ge(v21, float64(1)) != 0 {
									v59 = int32(-1)
									v60 = *(*float64)(unsafe.Add(mBase, uint32(v31)+56))
									v61 = *(*float64)(unsafe.Add(mBase, uint32(v40)+56))
									if base.F64_lt(v60, v61) != 0 {
										v86 = v59
										v90 = v86
									} else {
										if base.F64_gt(v60, v61) != 0 {
											v90 = int32(1)
										} else {
											v65 = *(*float64)(unsafe.Add(mBase, uint32(v31)+48))
											v66 = *(*float64)(unsafe.Add(mBase, uint32(v40)+48))
											if base.F64_lt(v65, v66) != 0 {
												v86 = v59
												v90 = v86
											} else {
												if base.F64_gt(v65, v66) != 0 {
													v86 = int32(1)
													v90 = v86
												} else {
													v90 = int32(0)
												}
											}
										}
									}
								} else {
									v72 = *(*float64)(unsafe.Add(mBase, uint32(v31)+56))
									v73 = *(*float64)(unsafe.Add(mBase, uint32(v31)+48))
									v76 = base.F64_add(base.F64_mul(v21, base.F64_sub(v72, v73)), v73)
									v77 = *(*float64)(unsafe.Add(mBase, uint32(v40)+56))
									v78 = *(*float64)(unsafe.Add(mBase, uint32(v40)+48))
									v81 = base.F64_add(base.F64_mul(v21, base.F64_sub(v77, v78)), v78)
									if base.F64_lt(v76, v81) != 0 {
										v86 = int32(-1)
									} else {
										v86 = base.F64_lt(v81, v76)
									}
									v90 = v86
								}
							}
							if v90 <= int32(0) {
								v93 = v31
							} else {
								v93 = v40
							}
							v94 = v93
						}
					}
					v96 = v34 + int32(1)
					v97 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
					if v96 < v97 {
						v31 = v94
						v34 = v96
						continue
					} else {
						break
					}
					break
				}
				v101 = v94
			}
		}
	}
	return v101
}
func F_get_cheapest_fractional_path_for_pathkeys(m *base.Module, l0 int32, l1 int32, l2 float64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 float64
	_ = v51
	var v52 float64
	_ = v52
	var v56 float64
	_ = v56
	var v57 float64
	_ = v57
	var v63 float64
	_ = v63
	var v64 float64
	_ = v64
	var v67 float64
	_ = v67
	var v68 float64
	_ = v68
	var v69 float64
	_ = v69
	var v72 float64
	_ = v72
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	v4 = int32(0)
	if l0 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v16 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v24 = v4
	v27 = v4
	goto L7
L5:
	;
	v218 = v4
	goto L6
L6:
	;
	return v218
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30+v27<<(uint(int32(2))%32))))
	if v24 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v218 = v203
	goto L6
L9:
	;
	v210 = v27 + int32(1)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v210 < v211 {
		v24 = v203
		v27 = v210
		goto L7
	} else {
		goto L70
	}
L10:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v34)+40))
	if v38 != v39 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	goto L12
L12:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v34)+64))
	if l1 == v84 {
		goto L32
	} else {
		goto L33
	}
L13:
	;
	if v81 <= int32(0) {
		v203 = v24
		goto L9
	} else {
		goto L31
	}
L14:
	;
	if v38 < v39 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	if base.F64_le(l2, float64(0))|base.F64_ge(l2, float64(1)) != 0 {
		goto L21
	} else {
		goto L22
	}
L17:
	;
	v44 = int32(-1)
	goto L19
L18:
	;
	v44 = int32(1)
	goto L19
L19:
	;
	v81 = v44
	goto L13
L20:
	;
	v81 = v77
	goto L13
L21:
	;
	v50 = int32(-1)
	v51 = *(*float64)(unsafe.Add(mBase, uint32(v24)+56))
	v52 = *(*float64)(unsafe.Add(mBase, uint32(v34)+56))
	if base.F64_lt(v51, v52) != 0 {
		v77 = v50
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v63 = *(*float64)(unsafe.Add(mBase, uint32(v24)+56))
	v64 = *(*float64)(unsafe.Add(mBase, uint32(v24)+48))
	v67 = base.F64_add(base.F64_mul(l2, base.F64_sub(v63, v64)), v64)
	v68 = *(*float64)(unsafe.Add(mBase, uint32(v34)+56))
	v69 = *(*float64)(unsafe.Add(mBase, uint32(v34)+48))
	v72 = base.F64_add(base.F64_mul(l2, base.F64_sub(v68, v69)), v69)
	if base.F64_lt(v67, v72) != 0 {
		v77 = int32(-1)
		goto L20
	} else {
		goto L30
	}
L24:
	;
	if base.F64_gt(v51, v52) != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v81 = int32(1)
	goto L13
L26:
	;
	goto L27
L27:
	;
	v56 = *(*float64)(unsafe.Add(mBase, uint32(v24)+48))
	v57 = *(*float64)(unsafe.Add(mBase, uint32(v34)+48))
	if base.F64_lt(v56, v57) != 0 {
		v77 = v50
		goto L20
	} else {
		goto L28
	}
L28:
	;
	if base.F64_gt(v56, v57) != 0 {
		v77 = int32(1)
		goto L20
	} else {
		goto L29
	}
L29:
	;
	v81 = int32(0)
	goto L13
L30:
	;
	v77 = base.F64_lt(v72, v67)
	goto L20
L31:
	;
	goto L12
L32:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	if v138 != 0 {
		goto L50
	} else {
		goto L51
	}
L33:
	;
	v90 = int32(0)
	goto L34
L34:
	;
	v98 = int32(0)
	if l1 == v98 {
		v108 = v98
		goto L36
	} else {
		goto L37
	}
L35:
	;
	if v108 != 0 {
		v203 = v24
		goto L9
	} else {
		goto L49
	}
L36:
	;
	if v84 != 0 {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v102 <= v90 {
		v108 = int32(0)
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v108 = v104 + v90<<(uint(int32(2))%32)
	goto L36
L39:
	;
	if v108 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L40:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v90 < v109 {
		goto L39
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	if v108 == int32(0) {
		goto L32
	} else {
		goto L44
	}
L43:
	;
	goto L42
L44:
	;
	v203 = v24
	goto L9
L45:
	;
	goto L35
L46:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
	if v115 == int32(0) {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v115+v90<<(uint(int32(2))%32))))
	if v122 == v124 {
		v90 = v90 + int32(1)
		goto L34
	} else {
		goto L48
	}
L48:
	;
	v203 = v24
	goto L9
L49:
	;
	goto L32
L50:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	v141 = v139
	goto L52
L51:
	;
	v141 = int32(0)
	goto L52
L52:
	;
	v142 = int32(0)
	if v141 == v142 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	if v196 != 0 {
		goto L67
	} else {
		goto L68
	}
L54:
	;
	v196 = int32(1)
	goto L53
L55:
	;
	goto L56
L56:
	;
	goto L57
L57:
	;
	v196 = v142
	goto L53
L67:
	;
	v197 = v34
	goto L69
L68:
	;
	v197 = v24
	goto L69
L69:
	;
	v203 = v197
	goto L9
L70:
	;
	goto L8
}
func F_get_cheapest_path_for_pathkeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v18 int32
	_ = v18
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 float64
	_ = v56
	var v57 float64
	_ = v57
	var v61 float64
	_ = v61
	var v62 float64
	_ = v62
	var v68 int32
	_ = v68
	var v69 float64
	_ = v69
	var v70 float64
	_ = v70
	var v74 float64
	_ = v74
	var v75 float64
	_ = v75
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v233 int32
	_ = v233
	v6 = int32(0)
	if l0 == v6 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v18 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v28 = v6
	v31 = v6
	goto L7
L5:
	;
	v233 = v6
	goto L6
L6:
	;
	return v233
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34+v31<<(uint(int32(2))%32))))
	if l4 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v233 = v216
	goto L6
L9:
	;
	v223 = v31 + int32(1)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v223 < v224 {
		v28 = v216
		v31 = v223
		goto L7
	} else {
		goto L80
	}
L10:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+21)))
	if v39 != int32(1) {
		v216 = v28
		goto L9
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	if v28 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L12
L14:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v28)+40))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v38)+40))
	if v46 != v47 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	goto L16
L16:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v38)+64))
	if l1 == v92 {
		goto L42
	} else {
		goto L43
	}
L17:
	;
	if v89 <= int32(0) {
		v216 = v28
		goto L9
	} else {
		goto L41
	}
L18:
	;
	if v46 < v47 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	if l3 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L21:
	;
	v52 = int32(-1)
	goto L23
L22:
	;
	v52 = int32(1)
	goto L23
L23:
	;
	v89 = v52
	goto L17
L24:
	;
	v89 = v83
	goto L17
L25:
	;
	v83 = int32(0)
	goto L24
L26:
	;
	v55 = int32(-1)
	v56 = *(*float64)(unsafe.Add(mBase, uint32(v28)+48))
	v57 = *(*float64)(unsafe.Add(mBase, uint32(v38)+48))
	if base.F64_lt(v56, v57) != 0 {
		v83 = v55
		goto L24
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v68 = int32(-1)
	v69 = *(*float64)(unsafe.Add(mBase, uint32(v28)+56))
	v70 = *(*float64)(unsafe.Add(mBase, uint32(v38)+56))
	if base.F64_lt(v69, v70) != 0 {
		v83 = v68
		goto L24
	} else {
		goto L35
	}
L29:
	;
	if base.F64_gt(v56, v57) != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v89 = int32(1)
	goto L17
L31:
	;
	goto L32
L32:
	;
	v61 = *(*float64)(unsafe.Add(mBase, uint32(v28)+56))
	v62 = *(*float64)(unsafe.Add(mBase, uint32(v38)+56))
	if base.F64_lt(v61, v62) != 0 {
		v83 = v55
		goto L24
	} else {
		goto L33
	}
L33:
	;
	if base.F64_gt(v61, v62) == int32(0) {
		goto L25
	} else {
		goto L34
	}
L34:
	;
	v83 = int32(1)
	goto L24
L35:
	;
	if base.F64_gt(v69, v70) != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v89 = int32(1)
	goto L17
L37:
	;
	goto L38
L38:
	;
	v74 = *(*float64)(unsafe.Add(mBase, uint32(v28)+48))
	v75 = *(*float64)(unsafe.Add(mBase, uint32(v38)+48))
	if base.F64_lt(v74, v75) != 0 {
		v83 = v68
		goto L24
	} else {
		goto L39
	}
L39:
	;
	if base.F64_gt(v74, v75) != 0 {
		v83 = int32(1)
		goto L24
	} else {
		goto L40
	}
L40:
	;
	goto L25
L41:
	;
	goto L16
L42:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	if v150 != 0 {
		goto L60
	} else {
		goto L61
	}
L43:
	;
	v100 = int32(0)
	goto L44
L44:
	;
	v108 = int32(0)
	if l1 == v108 {
		v118 = v108
		goto L46
	} else {
		goto L47
	}
L45:
	;
	if v118 != 0 {
		v216 = v28
		goto L9
	} else {
		goto L59
	}
L46:
	;
	if v92 != 0 {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v112 <= v100 {
		v118 = int32(0)
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v118 = v114 + v100<<(uint(int32(2))%32)
	goto L46
L49:
	;
	if v118 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L50:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	if v100 < v119 {
		goto L49
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	if v118 == int32(0) {
		goto L42
	} else {
		goto L54
	}
L53:
	;
	goto L52
L54:
	;
	v216 = v28
	goto L9
L55:
	;
	goto L45
L56:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v92)+12))
	if v125 == int32(0) {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v125+v100<<(uint(int32(2))%32))))
	if v132 == v134 {
		v100 = v100 + int32(1)
		goto L44
	} else {
		goto L58
	}
L58:
	;
	v216 = v28
	goto L9
L59:
	;
	goto L42
L60:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+4))
	v153 = v151
	goto L62
L61:
	;
	v153 = int32(0)
	goto L62
L62:
	;
	v154 = int32(0)
	if v153 == v154 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	if v207 != 0 {
		goto L77
	} else {
		goto L78
	}
L64:
	;
	v207 = int32(1)
	goto L63
L65:
	;
	goto L66
L66:
	;
	if l2 == int32(0) {
		v200 = v154
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v207 = v200
	goto L63
L68:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v164 < v163 {
		v200 = v154
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v166 = int32(1)
	if v163 <= v166 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v169 = v166
	goto L72
L71:
	;
	v169 = v163
	goto L72
L72:
	;
	v170 = int32(8)
	v175 = int32(0)
	goto L73
L73:
	;
	v182 = v175 << (uint(int32(2)) % 32)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v153+v170+v182)))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l2+v170+v182)))
	v189 = v184 & (v186 ^ int32(-1))
	v191 = base.B2i32(v189 == int32(0))
	if v189 != 0 {
		v200 = v191
		goto L67
	} else {
		goto L75
	}
L74:
	;
	v200 = v191
	goto L67
L75:
	;
	v193 = v175 + int32(1)
	if v193 != v169 {
		v175 = v193
		goto L73
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	v208 = v38
	goto L79
L78:
	;
	v208 = v28
	goto L79
L79:
	;
	v216 = v208
	goto L9
L80:
	;
	goto L8
}
func F_get_decomposed_size(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
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
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	v9 = l0 - int32(_a_F_get_decomposed_size_0)
	if base.Ui32(v9) <= base.Ui32(int32(_a_F_get_decomposed_size_1)) {
		v17 = base.I32_rem_u_s(v9&int32(_a_F_get_decomposed_size_2), int32(28))
		if v17 != 0 {
			v18 = int32(3)
		} else {
			v18 = int32(2)
		}
		return v18
	} else {
		v20 = int32(1)
		v21 = int32(16711935)
		v23 = int32(8)
		v24 = base.I32_rotr(l0&v21, v23)
		v25 = int32(24)
		v26 = base.I32_rotr(l0, v25)
		v28 = int32(255)
		v29 = (v24 | v26) & v28
		v30 = int32(_a_F_get_decomposed_size_3)
		v35 = int32(base.Ui32(v24)>>(uint(v23)%32)) & v28
		v42 = int32(base.Ui32(v26&v21) >> (uint(int32(16)) % 32))
		v47 = int32(base.Ui32(v24) >> (uint(v25) % 32))
		v51 = int32(_a_F_get_decomposed_size_4)
		v52 = base.I32_rem_u_s(((v29*v30+v35)*v30+v42)*v30+v47+int32(402620417), v51)
		v55 = int32(*(*int16)(unsafe.Add(mBase, uint32(v52<<(uint(v20)%32))+uint32(_c_F_get_decomposed_size[0]))))
		v56 = int32(257)
		v66 = base.I32_rem_u_s(((v29*v56+v35)*v56+v42)*v56+v47, v51)
		v69 = int32(*(*int16)(unsafe.Add(mBase, uint32(v66<<(uint(v20)%32))+uint32(_c_F_get_decomposed_size[0]))))
		v70 = v55 + v69
		if base.Ui32(int32(_a_F_get_decomposed_size_5)) < base.Ui32(v70) {
			v257 = v20
		} else {
			v74 = v70 << (uint(int32(3)) % 32)
			v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+uint32(_c_F_get_decomposed_size[1])))
			if l0 != v75 {
				v257 = v20
			} else {
				v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+uint32(_c_F_get_decomposed_size[2]))))
				v81 = v79 & int32(31)
				if v79&int32(32) != 0 {
					v87 = l1
				} else {
					v87 = int32(1)
				}
				if base.B2i32(v81 == int32(0))|base.B2i32(v87 == int32(0)) != 0 {
					v257 = v20
				} else {
					v91 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v74)+uint32(_c_F_get_decomposed_size[3]))))
					if v79&int32(64) != 0 {
						v94 = int32(_a_F_get_decomposed_size_6)
						*(*int32)(unsafe.Add(mBase, _c_F_get_decomposed_size[4])) = v91
						v102 = int32(1)
						v103 = v94
					} else {
						v102 = v81
						v103 = v91<<(uint(int32(2))%32) + int32(_a_F_get_decomposed_size_7)
					}
					v104 = int32(0)
					v106 = v104
					v109 = v104
					for {
						v116 = *(*int32)(unsafe.Add(mBase, uint32(v103+v106<<(uint(int32(2))%32))))
						v123 = v116 - int32(_a_F_get_decomposed_size_0)
						if base.Ui32(v123) <= base.Ui32(int32(_a_F_get_decomposed_size_1)) {
							v131 = base.I32_rem_u_s(v123&int32(_a_F_get_decomposed_size_2), int32(28))
							if v131 != 0 {
								v132 = int32(3)
							} else {
								v132 = int32(2)
							}
							v249 = v132
						} else {
							v133 = int32(1)
							v134 = int32(16711935)
							v136 = int32(8)
							v137 = base.I32_rotr(v116&v134, v136)
							v138 = int32(24)
							v139 = base.I32_rotr(v116, v138)
							v141 = int32(255)
							v142 = (v137 | v139) & v141
							v143 = int32(_a_F_get_decomposed_size_3)
							v148 = int32(base.Ui32(v137)>>(uint(v136)%32)) & v141
							v155 = int32(base.Ui32(v139&v134) >> (uint(int32(16)) % 32))
							v160 = int32(base.Ui32(v137) >> (uint(v138) % 32))
							v164 = int32(_a_F_get_decomposed_size_4)
							v165 = base.I32_rem_u_s(((v142*v143+v148)*v143+v155)*v143+v160+int32(402620417), v164)
							v168 = int32(*(*int16)(unsafe.Add(mBase, uint32(v165<<(uint(v133)%32))+uint32(_c_F_get_decomposed_size[0]))))
							v169 = int32(257)
							v179 = base.I32_rem_u_s(((v142*v169+v148)*v169+v155)*v169+v160, v164)
							v182 = int32(*(*int16)(unsafe.Add(mBase, uint32(v179<<(uint(v133)%32))+uint32(_c_F_get_decomposed_size[0]))))
							v183 = v168 + v182
							if base.Ui32(int32(_a_F_get_decomposed_size_5)) < base.Ui32(v183) {
								v238 = v133
							} else {
								v187 = v183 << (uint(int32(3)) % 32)
								v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_get_decomposed_size[1])))
								if v116 != v188 {
									v238 = v133
								} else {
									v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_get_decomposed_size[2]))))
									v194 = v192 & int32(31)
									if v192&int32(32) != 0 {
										v200 = l1
									} else {
										v200 = int32(1)
									}
									if base.B2i32(v194 == int32(0))|base.B2i32(v200 == int32(0)) != 0 {
										v238 = v133
									} else {
										v204 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_get_decomposed_size[3]))))
										if v192&int32(64) != 0 {
											v207 = int32(_a_F_get_decomposed_size_6)
											*(*int32)(unsafe.Add(mBase, _c_F_get_decomposed_size[4])) = v204
											v215 = int32(1)
											v216 = v207
										} else {
											v215 = v194
											v216 = v204<<(uint(int32(2))%32) + int32(_a_F_get_decomposed_size_7)
										}
										v217 = int32(0)
										v219 = v217
										v222 = v217
										for {
											v229 = *(*int32)(unsafe.Add(mBase, uint32(v216+v219<<(uint(int32(2))%32))))
											v230 = F_get_decomposed_size(m, v229, l1)
											mBase = m.M
											v231 = v230 + v222
											v233 = v219 + int32(1)
											if v233 != v215 {
												v219 = v233
												v222 = v231
												continue
											} else {
												break
											}
											break
										}
										v238 = v231
									}
								}
							}
							v249 = v238
						}
						v250 = v249 + v109
						v252 = v106 + int32(1)
						if v252 != v102 {
							v106 = v252
							v109 = v250
							continue
						} else {
							break
						}
						break
					}
					v257 = v250
				}
			}
		}
		return v257
	}
}
func F_get_element_type(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	v5 = F_SearchSysCache1(m, int32(82), l0)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 == int32(0) {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+22)))
			v15 = v13 + v14
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+92))
			if v16 != 0 {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
				if v17 == int32(_a_F_get_element_type_0) {
					v21 = v16
				} else {
					v21 = int32(0)
				}
			} else {
				v21 = int32(0)
			}
			F_ReleaseCatCache(m, v5)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				return v21
			}
		}
	}
}
func F_get_formatted_log_time(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	v2 = m.G0
	v4 = v2 - int32(32)
	m.G0 = v4
	v7 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_get_formatted_log_time[0])))
	if v7 == int32(0) {
		v11 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_get_formatted_log_time[1])))
		if v11 == int32(0) {
			F_gettimeofday(m, int32(_a_F_get_formatted_log_time_0))
			mBase = m.M
			v17 = int32(1)
			*(*uint8)(unsafe.Add(mBase, _c_F_get_formatted_log_time[1])) = uint8(v17)
		} else {
		}
		v20 = *(*int64)(unsafe.Add(mBase, _c_F_get_formatted_log_time[2]))
		*(*int64)(unsafe.Add(mBase, uint32(v4)+24)) = v20
		v28 = *(*int32)(unsafe.Add(mBase, _c_F_get_formatted_log_time[3]))
		v29 = F_pg_localtime(m, v4+int32(24), v28)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			v33 = F_pg_strftime(m, int32(_a_F_get_formatted_log_time_1), int32(128), int32(_a_F_get_formatted_log_time_2), v29)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, _c_F_get_formatted_log_time[4]))
				v38 = base.I32_div_s(v36, int32(1000))
				*(*int32)(unsafe.Add(mBase, uint32(v4))) = v38
				v43 = F_pg_sprintf(m, v4+int32(8), int32(_a_F_get_formatted_log_time_3), v4)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
					*(*int32)(unsafe.Add(mBase, _c_F_get_formatted_log_time[5])) = v46
					m.G0 = v4 + int32(32)
					return int32(_a_F_get_formatted_log_time_1)
				}
			}
		}
	} else {
		m.G0 = v4 + int32(32)
		return int32(_a_F_get_formatted_log_time_1)
	}
}
func F_get_formatted_start_time(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
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
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	v7 = *(*int64)(unsafe.Add(mBase, _c_F_get_formatted_start_time[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v4)+8)) = v7
	v10 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_get_formatted_start_time[1])))
	if v10 == int32(0) {
		v19 = *(*int32)(unsafe.Add(mBase, _c_F_get_formatted_start_time[2]))
		v20 = F_pg_localtime(m, v4+int32(8), v19)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v24 = F_pg_strftime(m, int32(_a_F_get_formatted_start_time_0), int32(128), int32(_a_F_get_formatted_start_time_1), v20)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				m.G0 = v4 + int32(16)
				return int32(_a_F_get_formatted_start_time_0)
			}
		}
	} else {
		m.G0 = v4 + int32(16)
		return int32(_a_F_get_formatted_start_time_0)
	}
}
func F_get_nullingrels_recurse(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
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
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
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
	var v102 int32
	_ = v102
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v8 + int32(32)
	return
L2:
	;
	v12 = l0
	v13 = l1
	goto L3
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v17 != int32(64) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	goto L1
L5:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	F_get_nullingrels_recurse(m, v99, v98, l2)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L16
	} else {
		goto L35
	}
L6:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v91+v92<<(uint(int32(2))%32)))) = v13
	goto L1
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L16
	} else {
		goto L32
	}
L8:
	;
	switch v17 - int32(63) {
	case 0:
		goto L6
	default:
		goto L7
	case 2:
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	switch v45 {
	case 0:
		v97 = v13
		v98 = v13
		goto L5
	case 1, 4, 5:
		goto L22
	case 2:
		goto L21
	case 3:
		goto L20
	default:
		goto L19
	}
L11:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v22 == int32(0) {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v25 <= int32(0) {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v29 = int32(0)
	goto L14
L14:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34+v29<<(uint(int32(2))%32))))
	F_get_nullingrels_recurse(m, v38, v13, l2)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L1
L16:
	;
	return
L17:
	;
	v42 = v29 + int32(1)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v42 < v43 {
		v29 = v42
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L16
	} else {
		goto L29
	}
L20:
	;
	v56 = F_bms_copy(m, v13)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L16
	} else {
		goto L27
	}
L21:
	;
	v51 = F_bms_copy(m, v13)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L16
	} else {
		goto L25
	}
L22:
	;
	v46 = F_bms_copy(m, v13)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L16
	} else {
		goto L23
	}
L23:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	v49 = F_bms_add_member(m, v46, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L16
	} else {
		goto L24
	}
L24:
	;
	v97 = v49
	v98 = v13
	goto L5
L25:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	v54 = F_bms_add_member(m, v51, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L16
	} else {
		goto L26
	}
L26:
	;
	v97 = v54
	v98 = v54
	goto L5
L27:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	v59 = F_bms_add_member(m, v56, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L16
	} else {
		goto L28
	}
L28:
	;
	v97 = v13
	v98 = v59
	goto L5
L29:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v65
	F_errmsg_internal(m, int32(_a_F_get_nullingrels_recurse_0), v8+int32(16))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L16
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_get_nullingrels_recurse_1), int32(_a_F_get_nullingrels_recurse_2), int32(_a_F_get_nullingrels_recurse_3))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L16
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v81
	F_errmsg_internal(m, int32(_a_F_get_nullingrels_recurse_4), v8)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L16
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_get_nullingrels_recurse_1), int32(_a_F_get_nullingrels_recurse_5), int32(_a_F_get_nullingrels_recurse_3))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L16
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	if v102 != 0 {
		v12 = v102
		v13 = v97
		goto L3
	} else {
		goto L36
	}
L36:
	;
	goto L4
}
func F_get_oprrest(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v4 = F_SearchSysCache1(m, int32(40), l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v12+v13)+104))
			F_ReleaseCatCache(m, v4)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	}
}
func F_get_ordering_op_for_equality_op(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	v2 = int32(0)
	v11 = F_SearchSysCacheList(m, int32(3), int32(1), l0, v2, v2)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_ReleaseCatCacheList(m, v11)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L2
	} else {
		goto L24
	}
L2:
	;
	return int32(0)
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	if v15 <= int32(0) {
		v77 = v2
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v21 = int32(0)
	goto L5
L5:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(48)+v21<<(uint(int32(2))%32))))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+56))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+22)))
	v33 = v31 + v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	if v34 <= int32(2741) {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	v77 = int32(0)
	goto L1
L7:
	;
	v72 = v21 + int32(1)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	if v72 < v73 {
		v21 = v72
		goto L5
	} else {
		goto L23
	}
L8:
	;
	v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33)+16)))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v61 = F_IndexAmTranslateStrategy(m, v59, v58, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L2
	} else {
		goto L19
	}
L9:
	;
	v50 = F_GetIndexAmRoutineByAmId(m, v34, int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L2
	} else {
		goto L16
	}
L10:
	;
	switch v34 - int32(403) {
	case 0:
		v58 = v34
		goto L8
	case 1:
		goto L9
	case 2:
		goto L7
	default:
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	if base.B2i32(v34 == int32(2742))|base.B2i32(v34 == int32(3580))|base.B2i32(v34 == int32(4000)) != 0 {
		goto L7
	} else {
		goto L15
	}
L13:
	;
	if v34 != int32(783) {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	goto L7
L15:
	;
	goto L9
L16:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+10)))
	F_pfree(m, v50)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	if v52 != int32(1) {
		goto L7
	} else {
		goto L18
	}
L18:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	v58 = v57
	goto L8
L19:
	;
	if v61 != int32(3) {
		goto L7
	} else {
		goto L20
	}
L20:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v68 = F_get_opfamily_member_for_cmptype(m, v65, v66, v66, int32(1))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	if v68 != 0 {
		v77 = v68
		goto L1
	} else {
		goto L22
	}
L22:
	;
	goto L7
L23:
	;
	goto L6
L24:
	;
	return v77
}
func F_get_relname_relid(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = int32(0)
	v6 = F_GetSysCacheOid(m, int32(56), l0, l1, v4, v4)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_get_required_extension(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	v7 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v18 = F_GetSysCacheOid(m, int32(27), l0, v7, v7, v7)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L2
	} else {
		goto L41
	}
L2:
	;
	return int32(0)
L3:
	;
	if v18 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if l3 == int32(0) {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	v140 = v18
	goto L6
L6:
	;
	m.G0 = v12 + int32(48)
	return v140
L7:
	;
	F_check_valid_extension_name(m, l0)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	if l4 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v111 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L2
	} else {
		goto L32
	}
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v30 <= int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v33 = int32(0)
	if v33 < v30 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v37 = v30
	goto L14
L13:
	;
	v37 = v33
	goto L14
L14:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v46 = v33
	goto L15
L15:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v38+v46<<(uint(int32(2))%32))))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v54 == int32(0))|base.B2i32(v54 != v57) != 0 {
		v75 = v54
		v76 = v57
		goto L18
	} else {
		goto L19
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L2
	} else {
		goto L28
	}
L17:
	;
	if v75-v76 != 0 {
		goto L24
	} else {
		goto L25
	}
L18:
	;
	goto L17
L19:
	;
	v60 = v51
	v61 = l0
	goto L20
L20:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+1)))
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+1)))
	if v65 == int32(0) {
		v75 = v65
		v76 = v64
		goto L18
	} else {
		goto L22
	}
L21:
	;
	v75 = v65
	v76 = v64
	goto L18
L22:
	;
	v68 = int32(1)
	if v65 == v64 {
		v60 = v60 + v68
		v61 = v61 + v68
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v79 = v46 + int32(1)
	if v37 != v79 {
		v46 = v79
		goto L15
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	goto L16
L27:
	;
	goto L9
L28:
	;
	F_errcode(m, int32(151388292))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L2
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l0
	F_errmsg(m, int32(_a_F_get_required_extension_0), v12+int32(16))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L2
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_get_required_extension_1), int32(2054), int32(_a_F_get_required_extension_2))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	if v111 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
	F_errmsg(m, int32(_a_F_get_required_extension_3), v12)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L2
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v126 = F_list_copy(m, l4)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L2
	} else {
		goto L38
	}
L36:
	;
	F_errfinish(m, int32(_a_F_get_required_extension_1), int32(2059), int32(_a_F_get_required_extension_2))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v128 = F_lappend(m, v126, l1)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	F_CreateExtensionInternal(m, v12+int32(36), l0, l2, int32(0), int32(1), v128, l5)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L2
	} else {
		goto L40
	}
L40:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	v140 = v132
	goto L6
L41:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = l0
	F_errmsg(m, int32(_a_F_get_required_extension_4), v12+int32(32))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	if l5 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	F_errhint(m, int32(_a_F_get_required_extension_5), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L2
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	F_errfinish(m, int32(_a_F_get_required_extension_1), int32(2084), int32(_a_F_get_required_extension_2))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L2
	} else {
		goto L48
	}
L47:
	;
	goto L46
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_sortgroupref_clause_noerr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v38 int32
	_ = v38
	if l1 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v38
L2:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v6 <= int32(0) {
		v38 = int32(0)
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v38 = int32(0)
	goto L1
L5:
	;
	v9 = int32(0)
	if v9 < v6 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v12 = v6
	goto L8
L7:
	;
	v12 = v9
	goto L8
L8:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v16 = int32(0)
	goto L9
L9:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v13+v16<<(uint(int32(2))%32))))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v24 == l0 {
		v38 = v23
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L4
L11:
	;
	v27 = v16 + int32(1)
	if v27 != v12 {
		v16 = v27
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
}
func F_get_steps_using_prefix_recurse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v45 int32
	_ = v45
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
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
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v179 int32
	_ = v179
	F_check_stack_depth(m)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v22 = int32(2)
	v23 = (l7 - v20) >> (uint(v22) % 32)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v20+v26<<(uint(v22)%32)-int32(4))))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if v25 < v33 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	return v179
L4:
	;
	if v26 <= v23 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v110 = int32(0)
	if v26 <= v23 {
		v179 = v110
		goto L3
	} else {
		goto L27
	}
L7:
	;
	return int32(0)
L8:
	;
	goto L9
L9:
	;
	v45 = v23
	goto L11
L10:
	;
	v73 = int32(0)
	v77 = v23
	goto L15
L11:
	;
	v55 = v20 + v45<<(uint(int32(2))%32)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if v25 < v57 {
		v64 = v55
		goto L10
	} else {
		goto L13
	}
L12:
	;
	v64 = int32(0)
	goto L10
L13:
	;
	v60 = v45 + int32(1)
	if v60 != v26 {
		v45 = v60
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v81+v77<<(uint(int32(2))%32))))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	if v86 != v25 {
		v179 = v73
		goto L3
	} else {
		goto L17
	}
L16:
	;
	v179 = v100
	goto L3
L17:
	;
	v88 = F_list_copy(m, l8)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
	v91 = F_lappend(m, v88, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v93 = F_list_copy(m, l9)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
	v96 = F_lappend_oid(m, v93, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v98 = F_get_steps_using_prefix_recurse(m, l0, l1, l2, l3, l4, l5, l6, v64, v91, v96)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v100 = F_list_concat(m, v73, v98)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_list_free(m, v91)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_list_free(m, v96)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v107 = v77 + int32(1)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v107 < v108 {
		v73 = v100
		v77 = v107
		goto L15
	} else {
		goto L26
	}
L26:
	;
	goto L16
L27:
	;
	if l2 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v113 = int32(0)
	goto L30
L29:
	;
	v113 = l1
	goto L30
L30:
	;
	v121 = v110
	v125 = v23
	goto L31
L31:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v129+v125<<(uint(int32(2))%32))))
	v134 = F_list_copy(m, l8)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L33
	}
L32:
	;
	v179 = v166
	goto L3
L33:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v133)+12))
	v137 = F_lappend(m, v134, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v139 = F_lappend(m, v137, l3)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v141 = F_list_copy(m, l9)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v133)+16))
	v144 = F_lappend_oid(m, v141, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v146 = F_lappend_oid(m, v144, l4)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v149 = F_palloc0(m, int32(24))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v149))) = int32(377)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v153 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v149)+20)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v149)+16)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v149)+12)) = v139
	*(*uint16)(unsafe.Add(mBase, uint32(v149)+8)) = uint16(v113)
	*(*int32)(unsafe.Add(mBase, uint32(v149)+4)) = v153
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v163 = F_lappend(m, v162, v149)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v163
	v166 = F_lappend(m, v121, v149)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v169 = v125 + int32(1)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v169 < v170 {
		v121 = v166
		v125 = v169
		goto L31
	} else {
		goto L42
	}
L42:
	;
	goto L32
}
func F_get_typavgwidth(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	v6 = F_SearchSysCache1(m, int32(82), l0)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 != 0 {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
			v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+22)))
			v13 = int32(*(*int16)(unsafe.Add(mBase, uint32(v10+v11)+76)))
			F_ReleaseCatCache(m, v6)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				if int32(0) < v13 {
					v82 = v13
					return v82
				} else {
					v20 = int32(-1)
					if l1 < int32(0) {
						v60 = v20
						v62 = v60
					} else {
						if base.Ui32(int32(2)) <= base.Ui32(l0-int32(1042)) {
							switch l0 - int32(1560) {
							case 0, 2:
								v56 = int32(8)
								v57 = base.I32_div_s(l1+int32(7), v56)
								v60 = v57 + v56
								v62 = v60
							case 1:
								v60 = v20
								v62 = v60
							default:
								if l0 != int32(1700) {
									v60 = v20
									v62 = v60
								} else {
									v39 = int32(4)
									if l1 < v39 {
										v53 = int32(-1)
									} else {
										v53 = int32(base.Ui32(int32(base.Ui32(l1-v39)>>(uint(int32(16))%32))+int32(6))>>(uint(int32(1))%32))&int32(_a_F_get_typavgwidth_0) + int32(8)
									}
									v62 = v53
								}
							}
						} else {
							v29 = F_GetDatabaseEncoding(m)
							mBase = m.M
							v30 = F_pg_encoding_max_length(m, v29)
							mBase = m.M
							v31 = int32(4)
							v62 = v30*(l1-v31) + v31
						}
					}
					if v62 <= int32(0) {
						return int32(32)
					} else {
						if base.B2i32(l0 == int32(1042))|base.B2i32(base.Ui32(v62) < base.Ui32(int32(33))) != 0 {
							v82 = v62
							return v82
						} else {
							if base.Ui32(int32(999)) < base.Ui32(v62) {
								return int32(516)
							} else {
								v76 = int32(32)
								v82 = int32(base.Ui32(v62-v76)>>(uint(int32(1))%32)) + v76
								return v82
							}
						}
					}
				}
			}
		} else {
			v20 = int32(-1)
			if l1 < int32(0) {
				v60 = v20
				v62 = v60
			} else {
				if base.Ui32(int32(2)) <= base.Ui32(l0-int32(1042)) {
					switch l0 - int32(1560) {
					case 0, 2:
						v56 = int32(8)
						v57 = base.I32_div_s(l1+int32(7), v56)
						v60 = v57 + v56
						v62 = v60
					case 1:
						v60 = v20
						v62 = v60
					default:
						if l0 != int32(1700) {
							v60 = v20
							v62 = v60
						} else {
							v39 = int32(4)
							if l1 < v39 {
								v53 = int32(-1)
							} else {
								v53 = int32(base.Ui32(int32(base.Ui32(l1-v39)>>(uint(int32(16))%32))+int32(6))>>(uint(int32(1))%32))&int32(_a_F_get_typavgwidth_0) + int32(8)
							}
							v62 = v53
						}
					}
				} else {
					v29 = F_GetDatabaseEncoding(m)
					mBase = m.M
					v30 = F_pg_encoding_max_length(m, v29)
					mBase = m.M
					v31 = int32(4)
					v62 = v30*(l1-v31) + v31
				}
			}
			if v62 <= int32(0) {
				return int32(32)
			} else {
				if base.B2i32(l0 == int32(1042))|base.B2i32(base.Ui32(v62) < base.Ui32(int32(33))) != 0 {
					v82 = v62
					return v82
				} else {
					if base.Ui32(int32(999)) < base.Ui32(v62) {
						return int32(516)
					} else {
						v76 = int32(32)
						v82 = int32(base.Ui32(v62-v76)>>(uint(int32(1))%32)) + v76
						return v82
					}
				}
			}
		}
	}
}
func F_get_useful_group_keys_orderings(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v314 int32
	_ = v314
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	v17 = F_palloc0(m, int32(12))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v14
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(276)
	v26 = F_lappend(m, int32(0), v17)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_get_useful_group_keys_orderings[0])))
	if v29 != int32(1) {
		v314 = v26
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return v314
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
	if v32 != 0 {
		v314 = v26
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	if v33 == int32(0) {
		v314 = v26
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	if v33 == v36 {
		v314 = v26
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v42 = int32(0)
	goto L10
L9:
	;
	if v14 == int32(0) {
		v314 = v26
		goto L4
	} else {
		goto L26
	}
L10:
	;
	if v42 < v38 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v58 == int32(0) {
		v314 = v26
		goto L4
	} else {
		goto L25
	}
L12:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v58 = v54 + v42<<(uint(int32(2))%32)
	goto L14
L13:
	;
	v58 = int32(0)
	goto L14
L14:
	;
	if v36 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if v58 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L16:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	if v42 < v59 {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	if v58 != 0 {
		goto L9
	} else {
		goto L20
	}
L19:
	;
	goto L18
L20:
	;
	v314 = v26
	goto L4
L21:
	;
	goto L11
L22:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	if v63 == int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v63+v42<<(uint(int32(2))%32))))
	if v70 == v72 {
		v42 = v42 + int32(1)
		goto L10
	} else {
		goto L24
	}
L24:
	;
	goto L9
L25:
	;
	goto L9
L26:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v83 = F_list_copy_head(m, v14, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v85 = int32(0)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v86 <= v85 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v230 = F_list_concat_unique_ptr(m, v218, v14)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L74
	}
L29:
	;
	v89 = int32(0)
	v218 = v85
	v221 = v89
	v229 = v89
	goto L28
L30:
	;
	goto L31
L31:
	;
	v91 = int32(0)
	if v91 < v82 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v94 = v82
	goto L34
L33:
	;
	v94 = v91
	goto L34
L34:
	;
	v95 = int32(0)
	v98 = v85
	v99 = v95
	v101 = v95
	goto L35
L35:
	;
	if v99 == v94 {
		v207 = v98
		v209 = v101
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v207 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L37:
	;
	goto L36
L38:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v110+v99<<(uint(int32(2))%32))))
	v115 = int32(0)
	if v83 == v115 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	if v153 == int32(0) {
		v207 = v98
		v209 = v101
		goto L37
	} else {
		goto L52
	}
L40:
	;
	v153 = int32(0)
	goto L39
L41:
	;
	goto L42
L42:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	if v121 <= int32(0) {
		v147 = v115
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v153 = v147
	goto L39
L44:
	;
	v124 = int32(0)
	if v124 < v121 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v127 = v121
	goto L47
L46:
	;
	v127 = v124
	goto L47
L47:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
	v130 = int32(0)
	goto L48
L48:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v128+v130<<(uint(int32(2))%32))))
	v139 = base.B2i32(v138 == v114)
	if v138 == v114 {
		v147 = v139
		goto L43
	} else {
		goto L50
	}
L49:
	;
	v147 = v139
	goto L43
L50:
	;
	v141 = v130 + int32(1)
	if v141 != v127 {
		v130 = v141
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+44))
	if v157 == int32(0) {
		v207 = v98
		v209 = v101
		goto L37
	} else {
		goto L53
	}
L53:
	;
	if v15 != 0 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	if v195 == int32(0) {
		v207 = v98
		v209 = v101
		goto L37
	} else {
		goto L67
	}
L55:
	;
	goto L54
L56:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v163 <= int32(0) {
		v195 = int32(0)
		goto L55
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v195 = int32(0)
	goto L55
L59:
	;
	v166 = int32(0)
	if v166 < v163 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v169 = v163
	goto L62
L61:
	;
	v169 = v166
	goto L62
L62:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v173 = int32(0)
	goto L63
L63:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v170+v173<<(uint(int32(2))%32))))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+4))
	if v181 == v157 {
		v195 = v180
		goto L55
	} else {
		goto L65
	}
L64:
	;
	goto L58
L65:
	;
	v184 = v173 + int32(1)
	if v184 != v169 {
		v173 = v184
		goto L63
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	v199 = F_lappend(m, v98, v114)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v201 = F_lappend(m, v101, v195)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v204 = v99 + int32(1)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v204 < v205 {
		v98 = v199
		v99 = v204
		v101 = v201
		goto L35
	} else {
		goto L70
	}
L70:
	;
	v207 = v199
	v209 = v201
	goto L37
L71:
	;
	v214 = int32(0)
	v218 = v214
	v221 = v209
	v229 = v214
	goto L28
L72:
	;
	goto L73
L73:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v207)+4))
	v218 = v207
	v221 = v209
	v229 = v216
	goto L28
L74:
	;
	v232 = F_list_concat_unique_ptr(m, v221, v15)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	F_list_free(m, v83)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	if v229 <= int32(0) {
		v314 = v26
		goto L4
	} else {
		goto L77
	}
L77:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_get_useful_group_keys_orderings[1])))
	if v239 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	if v229 != v242 {
		v314 = v26
		goto L4
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	if v230 == v244 {
		v314 = v26
		goto L4
	} else {
		goto L82
	}
L81:
	;
	goto L80
L82:
	;
	v247 = int32(0)
	goto L85
L83:
	;
	v297 = F_palloc0(m, int32(12))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L98
	}
L84:
	;
	if v269 == int32(0) {
		v314 = v26
		goto L4
	} else {
		goto L97
	}
L85:
	;
	v259 = int32(0)
	if v230 == v259 {
		v269 = v259
		goto L87
	} else {
		goto L88
	}
L86:
	;
	if v269|v276 != 0 {
		goto L83
	} else {
		goto L96
	}
L87:
	;
	if v244 == int32(0) {
		goto L84
	} else {
		goto L90
	}
L88:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v230)+4))
	if v263 <= v247 {
		v269 = int32(0)
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v230)+12))
	v269 = v265 + v247<<(uint(int32(2))%32)
	goto L87
L90:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
	if v272 <= v247 {
		goto L84
	} else {
		goto L91
	}
L91:
	;
	v274 = int32(0)
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v244)+12))
	if base.B2i32(v269 == v274)|base.B2i32(v276 == v274) == v274 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v269)))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v276+v247<<(uint(int32(2))%32))))
	if v286 == v288 {
		v247 = v247 + int32(1)
		goto L85
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	goto L86
L95:
	;
	goto L83
L96:
	;
	v314 = v26
	goto L4
L97:
	;
	goto L83
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v297)+8)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v297)+4)) = v230
	*(*int32)(unsafe.Add(mBase, uint32(v297))) = int32(276)
	v303 = F_lappend(m, v26, v297)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v314 = v303
	goto L4
}
func F_get_useful_pathkeys_for_distinct(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	v4 = int32(0)
	v9 = F_lappend(m, v4, l1)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_get_useful_pathkeys_for_distinct[0])))
	if base.B2i32(l2 == int32(0))|base.B2i32(v16 != int32(1)) != 0 {
		v218 = v9
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return v218
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v20 <= int32(0) {
		v131 = v4
		goto L5
	} else {
		goto L6
	}
L5:
	;
	if v131 == int32(0) {
		v218 = v9
		goto L3
	} else {
		goto L42
	}
L6:
	;
	v26 = v4
	v28 = v4
	goto L7
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30+v28<<(uint(int32(2))%32))))
	v35 = int32(0)
	if l1 == v35 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v131 = v122
	goto L5
L9:
	;
	if v73 == int32(0) {
		v131 = v26
		goto L5
	} else {
		goto L22
	}
L10:
	;
	v73 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v41 <= int32(0) {
		v67 = v35
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v73 = v67
	goto L9
L14:
	;
	v44 = int32(0)
	if v44 < v41 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v47 = v41
	goto L17
L16:
	;
	v47 = v44
	goto L17
L17:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v50 = int32(0)
	goto L18
L18:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v48+v50<<(uint(int32(2))%32))))
	v59 = base.B2i32(v58 == v34)
	if v58 == v34 {
		v67 = v59
		goto L13
	} else {
		goto L20
	}
L19:
	;
	v67 = v59
	goto L13
L20:
	;
	v61 = v50 + int32(1)
	if v61 != v47 {
		v50 = v61
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+40)))
	if v77 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	v81 = int32(0)
	if v80 == v81 {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	goto L25
L25:
	;
	v122 = F_lappend(m, v26, v34)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L40
	}
L26:
	;
	if v119 == int32(0) {
		v131 = v26
		goto L5
	} else {
		goto L39
	}
L27:
	;
	v119 = int32(0)
	goto L26
L28:
	;
	goto L29
L29:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	if v87 <= int32(0) {
		v113 = v81
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v119 = v113
	goto L26
L31:
	;
	v90 = int32(0)
	if v90 < v87 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v93 = v87
	goto L34
L33:
	;
	v93 = v90
	goto L34
L34:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	v96 = int32(0)
	goto L35
L35:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v94+v96<<(uint(int32(2))%32))))
	v105 = base.B2i32(v104 == v34)
	if v104 == v34 {
		v113 = v105
		goto L30
	} else {
		goto L37
	}
L36:
	;
	v113 = v105
	goto L30
L37:
	;
	v107 = v96 + int32(1)
	if v107 != v93 {
		v96 = v107
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	goto L25
L40:
	;
	v125 = v28 + int32(1)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v125 < v126 {
		v26 = v122
		v28 = v125
		goto L7
	} else {
		goto L41
	}
L41:
	;
	goto L8
L42:
	;
	if l1 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v139 = v137
	goto L45
L44:
	;
	v139 = int32(0)
	goto L45
L45:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	if v140 < v139 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_get_useful_pathkeys_for_distinct[1])))
	if v143&int32(1) == int32(0) {
		v218 = v9
		goto L3
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v148 = F_list_concat_unique_ptr(m, v131, l1)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L50
	}
L49:
	;
	goto L48
L50:
	;
	if l1 == v148 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	if v209 == int32(0) {
		v218 = v9
		goto L3
	} else {
		goto L75
	}
L52:
	;
	v209 = int32(0)
	goto L51
L53:
	;
	goto L54
L54:
	;
	v158 = int32(0)
	goto L57
L55:
	;
	if v198 != 0 {
		goto L72
	} else {
		goto L73
	}
L56:
	;
	v193 = int32(0)
	if v180 != 0 {
		goto L69
	} else {
		goto L70
	}
L57:
	;
	v162 = int32(0)
	if l1 == v162 {
		v172 = v162
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v209 = int32(3)
	goto L51
L59:
	;
	if v148 != 0 {
		goto L63
	} else {
		goto L64
	}
L60:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v166 <= v158 {
		v172 = int32(0)
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v172 = v168 + v158<<(uint(int32(2))%32)
	goto L59
L62:
	;
	v178 = int32(0)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v148)+12))
	if base.B2i32(v172 == v178)|base.B2i32(v180 == v178) != 0 {
		goto L56
	} else {
		goto L67
	}
L63:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	if v158 < v173 {
		goto L62
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v175 = int32(0)
	v198 = base.B2i32(v172 == v175)
	v200 = v175
	goto L55
L66:
	;
	goto L65
L67:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v180+v158<<(uint(int32(2))%32))))
	if v188 == v190 {
		v158 = v158 + int32(1)
		goto L57
	} else {
		goto L68
	}
L68:
	;
	goto L58
L69:
	;
	v197 = int32(2)
	goto L71
L70:
	;
	v197 = v193
	goto L71
L71:
	;
	v198 = base.B2i32(v172 == v193)
	v200 = v197
	goto L55
L72:
	;
	v202 = v200
	goto L74
L73:
	;
	v202 = int32(1)
	goto L74
L74:
	;
	v209 = v202
	goto L51
L75:
	;
	v212 = F_lappend(m, v9, v148)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v218 = v212
	goto L3
}
func F_get_windowfunc_expr_helper(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	v6 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(448)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v18 == v6 {
		v66 = v6
		v70 = v6
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v15 + int32(448)
	return
L2:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v217)+72))
	v275 = F_quote_identifier(m, v274)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L13
	} else {
		goto L83
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L13
	} else {
		goto L79
	}
L4:
	;
	if l2 != 0 {
		goto L17
	} else {
		goto L18
	}
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if int32(101) <= v21 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v24 <= int32(0) {
		v66 = v6
		v70 = v6
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v32 = v6
	v36 = v6
	goto L8
L8:
	;
	v40 = v32 << (uint(int32(2)) % 32)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40+v41)))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	if v44 == int32(16) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v66 = v58
	v70 = v50
	goto L4
L10:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v48 = F_lappend(m, v36, v47)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v50 = v36
	goto L12
L12:
	;
	v54 = F_exprType(m, v43)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L13
	} else {
		goto L15
	}
L13:
	;
	return
L14:
	;
	v50 = v48
	goto L12
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15+int32(48)+v40))) = v54
	v58 = v32 + int32(1)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v58 < v59 {
		v32 = v58
		v36 = v50
		goto L8
	} else {
		goto L16
	}
L16:
	;
	goto L9
L17:
	;
	v81 = l2
	goto L19
L18:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v76 = int32(0)
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+34)))
	v79 = F_generate_function_name(m, v73, v66, v70, v15+int32(48), v76, v76, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L13
	} else {
		goto L20
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v81
	F_appendStringInfo(m, v17, int32(_a_F_get_windowfunc_expr_helper_0), v15+int32(32))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L13
	} else {
		goto L21
	}
L20:
	;
	v81 = v79
	goto L19
L21:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v88 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if l3 != 0 {
		goto L34
	} else {
		goto L35
	}
L23:
	;
	F_appendStringInfoChar(m, v17, int32(42))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L13
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if l4 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L22
L27:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+12))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	F_get_rule_expr(m, v96, l1, int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L13
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	F_get_rule_expr(m, v94, l1, int32(1))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L13
	} else {
		goto L33
	}
L30:
	;
	F_appendStringInfoString(m, v17, int32(_a_F_get_windowfunc_expr_helper_1))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L13
	} else {
		goto L31
	}
L31:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+12))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	F_get_rule_expr(m, v105, l1, int32(0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L13
	} else {
		goto L32
	}
L32:
	;
	goto L22
L33:
	;
	goto L22
L34:
	;
	F_appendStringInfoString(m, v17, l3)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L13
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v115 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L36
L38:
	;
	F_appendStringInfoString(m, v17, int32(_a_F_get_windowfunc_expr_helper_2))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L13
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	F_appendStringInfoString(m, v17, int32(_a_F_get_windowfunc_expr_helper_3))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L13
	} else {
		goto L43
	}
L41:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_get_rule_expr(m, v119, l1, int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L13
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v126 != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_get_rule_windowspec(m, v148, v255, l1)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L13
	} else {
		goto L78
	}
L45:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	if v127 <= int32(0) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L47
L47:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v189 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L48:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L13
	} else {
		goto L59
	}
L49:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v126)+12))
	v138 = int32(0)
	goto L50
L50:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v131+v138<<(uint(int32(2))%32))))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+48))
	if v130 != v149 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	if v154 == int32(0) {
		goto L44
	} else {
		goto L56
	}
L52:
	;
	v152 = v138 + int32(1)
	if v127 != v152 {
		v138 = v152
		goto L50
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	goto L51
L55:
	;
	goto L48
L56:
	;
	v157 = F_quote_identifier(m, v154)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L13
	} else {
		goto L57
	}
L57:
	;
	F_appendStringInfoString(m, v17, v157)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L13
	} else {
		goto L58
	}
L58:
	;
	goto L1
L59:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v177
	F_errmsg_internal(m, int32(_a_F_get_windowfunc_expr_helper_4), v15+int32(16))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L13
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_get_windowfunc_expr_helper_5), int32(_a_F_get_windowfunc_expr_helper_6), int32(_a_F_get_windowfunc_expr_helper_7))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L13
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L13
	} else {
		goto L75
	}
L63:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v189)+4))
	if v192 <= int32(0) {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v195 = int32(0)
	if v195 < v192 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v199 = v192
	goto L67
L66:
	;
	v199 = v195
	goto L67
L67:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v189)+12))
	v206 = v195
	goto L68
L68:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v200+v206<<(uint(int32(2))%32))))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)+40))
	if v217 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	goto L62
L70:
	;
	v227 = v206 + int32(1)
	if v227 != v199 {
		v206 = v227
		goto L68
	} else {
		goto L74
	}
L71:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
	if v220 != int32(366) {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v217)+76))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v223 == v224 {
		goto L2
	} else {
		goto L73
	}
L73:
	;
	goto L70
L74:
	;
	goto L69
L75:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v245
	F_errmsg_internal(m, int32(_a_F_get_windowfunc_expr_helper_4), v15)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L13
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_get_windowfunc_expr_helper_5), int32(_a_F_get_windowfunc_expr_helper_8), int32(_a_F_get_windowfunc_expr_helper_7))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L13
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	goto L1
L79:
	;
	F_errcode(m, int32(50856197))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L13
	} else {
		goto L80
	}
L80:
	;
	F_errmsg(m, int32(_a_F_get_windowfunc_expr_helper_9), int32(0))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L13
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_get_windowfunc_expr_helper_5), int32(_a_F_get_windowfunc_expr_helper_10), int32(_a_F_get_windowfunc_expr_helper_7))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L13
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	F_appendStringInfoString(m, v17, v275)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L13
	} else {
		goto L84
	}
L84:
	;
	goto L1
}
func F_get_xmltable(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v157 int32
	_ = v157
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
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v345 int32
	_ = v345
	var v367 int32
	_ = v367
	v4 = int32(0)
	v20 = m.G0
	v22 = v20 + int32(-64)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_appendStringInfoString(m, v24, int32(_a_F_get_xmltable_0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v28 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_appendStringInfoString(m, v24, int32(_a_F_get_xmltable_1))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	F_appendStringInfoChar(m, v24, int32(40))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L40
	}
L6:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v33 == int32(0) {
		v40 = v4
		goto L7
	} else {
		goto L8
	}
L7:
	;
	if v32 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v36 <= int32(0) {
		v40 = v4
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v40 = v39
	goto L7
L10:
	;
	F_appendStringInfoString(m, v24, int32(_a_F_get_xmltable_2))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L39
	}
L11:
	;
	v43 = int32(0)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if base.B2i32(v40 == v43)|base.B2i32(v45 <= v43) != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	if v49 == int32(0) {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	if v53 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v74 = int32(1)
	goto L23
L15:
	;
	F_get_rule_expr(m, v52, l1, l2)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	F_appendStringInfoString(m, v24, int32(_a_F_get_xmltable_3))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L21
	}
L18:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v57 = F_quote_identifier(m, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v57
	F_appendStringInfo(m, v24, int32(_a_F_get_xmltable_4), v20+int32(-16))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	goto L14
L21:
	;
	F_get_rule_expr(m, v52, l1, l2)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	goto L14
L23:
	;
	v90 = int32(0)
	if v33 == v90 {
		v99 = v90
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if base.B2i32(v99 == int32(0))|base.B2i32(v102 <= v74) != 0 {
		goto L10
	} else {
		goto L28
	}
L26:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v93 <= v74 {
		v99 = v90
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v99 = v95 + v74<<(uint(int32(2))%32)
	goto L25
L28:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	if v105 == int32(0) {
		goto L10
	} else {
		goto L29
	}
L29:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v105+v74<<(uint(int32(2))%32))))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	F_appendStringInfoString(m, v24, int32(_a_F_get_xmltable_5))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	if v111 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	F_get_rule_expr(m, v112, l1, l2)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	F_appendStringInfoString(m, v24, int32(_a_F_get_xmltable_3))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L37
	}
L34:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v119 = F_quote_identifier(m, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v119
	F_appendStringInfo(m, v24, int32(_a_F_get_xmltable_4), v20+int32(-32))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v74 = v74 + int32(1)
	goto L23
L37:
	;
	F_get_rule_expr(m, v112, l1, l2)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v74 = v74 + int32(1)
	goto L23
L39:
	;
	goto L5
L40:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_get_rule_expr(m, v180, l1, l2)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_appendStringInfoString(m, v24, int32(_a_F_get_xmltable_6))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	F_get_rule_expr(m, v186, l1, l2)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_appendStringInfoChar(m, v24, int32(41))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v192 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	F_appendStringInfoChar(m, v24, int32(41))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L1
	} else {
		goto L92
	}
L46:
	;
	F_appendStringInfoString(m, v24, int32(_a_F_get_xmltable_7))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v207 = int32(0)
	goto L48
L48:
	;
	v223 = int32(0)
	if v202 == v223 {
		v233 = v223
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v234 = int32(0)
	if v201 == v234 {
		v245 = v234
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v202)+4))
	if v227 <= v207 {
		v233 = int32(0)
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v202)+12))
	v233 = v229 + v207<<(uint(int32(2))%32)
	goto L50
L53:
	;
	if v200 == int32(0) {
		v254 = v234
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v201)+4))
	if v239 <= v207 {
		v245 = int32(0)
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v201)+12))
	v245 = v241 + v207<<(uint(int32(2))%32)
	goto L53
L56:
	;
	v255 = int32(0)
	if v199 == v255 {
		v266 = v255
		goto L59
	} else {
		goto L60
	}
L57:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v200)+4))
	if v248 <= v207 {
		v254 = v234
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v200)+12))
	v254 = v250 + v207<<(uint(int32(2))%32)
	goto L56
L59:
	;
	if v198 == int32(0) {
		v275 = v255
		goto L62
	} else {
		goto L63
	}
L60:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v199)+4))
	if v260 <= v207 {
		v266 = int32(0)
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v199)+12))
	v266 = v262 + v207<<(uint(int32(2))%32)
	goto L59
L62:
	;
	v276 = int32(0)
	if base.B2i32(v275 == v276)|(base.B2i32(v233 == v276)|base.B2i32(v245 == v276)|(base.B2i32(v254 == v276)|base.B2i32(v266 == v276))) != 0 {
		goto L45
	} else {
		goto L65
	}
L63:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v198)+4))
	if v269 <= v207 {
		v275 = v255
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v198)+12))
	v275 = v271 + v207<<(uint(int32(2))%32)
	goto L62
L65:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v275)))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v266)))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v254)))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v245)))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v295)+4))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v298 = F_bms_is_member(m, v207, v297)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	if int32(0) < v207 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	F_appendStringInfoString(m, v24, int32(_a_F_get_xmltable_5))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v306 = v207 + int32(1)
	v307 = F_quote_identifier(m, v296)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L71
	}
L70:
	;
	goto L69
L71:
	;
	if v207 != v290 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v310 = F_format_type_with_typemod(m, v294, v293)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = int32(_a_F_get_xmltable_8)
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v307
	F_appendStringInfo(m, v24, int32(_a_F_get_xmltable_9), v22)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L91
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v310
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v307
	F_appendStringInfo(m, v24, int32(_a_F_get_xmltable_9), v20+int32(-48))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	if v291 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	F_appendStringInfoString(m, v24, int32(_a_F_get_xmltable_10))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	if v292 != 0 {
		goto L83
	} else {
		goto L84
	}
L80:
	;
	F_get_rule_expr(m, v291, l1, l2)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_appendStringInfoChar(m, v24, int32(41))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	goto L79
L83:
	;
	F_appendStringInfoString(m, v24, int32(_a_F_get_xmltable_11))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	if v298 == int32(0) {
		v207 = v306
		goto L48
	} else {
		goto L89
	}
L86:
	;
	F_get_rule_expr(m, v292, l1, l2)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	F_appendStringInfoChar(m, v24, int32(41))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	goto L85
L89:
	;
	F_appendStringInfoString(m, v24, int32(_a_F_get_xmltable_12))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	v207 = v306
	goto L48
L91:
	;
	v207 = v306
	goto L48
L92:
	;
	m.G0 = v22 - int32(-64)
	return
}
func F_getcwd(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	v7 = m.G0
	v8 = int32(_a_F_getcwd_0)
	if l0 != 0 {
		v11 = int32(16)
	} else {
		v11 = v8
	}
	v12 = v7 - v11
	m.G0 = v12
	if l0 == int32(0) {
		v20 = v12
		v21 = v8
		v22 = int32(0)
		v23 = m.Env.X__syscall_getcwd(m, v20, v21)
		mBase = m.M
		if base.Ui32(int32(-4095)) <= base.Ui32(v23) {
			*(*int32)(unsafe.Add(mBase, _c_F_getcwd[0])) = int32(0) - v23
			v31 = int32(-1)
		} else {
			v31 = v23
		}
		if v31 < int32(0) {
			v52 = v22
		} else {
			if v31 != 0 {
				v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
				if v34 == int32(47) {
					if v20 != v12 {
						v52 = v20
					} else {
						v43 = F_strlen(m, v20)
						mBase = m.M
						v45 = v43 + int32(1)
						v46 = F_emscripten_builtin_malloc(m, v45)
						mBase = m.M
						if v46 == int32(0) {
							v51 = int32(0)
						} else {
							v50 = F___memcpy(m, v46, v20, v45)
							mBase = m.M
							v51 = v50
						}
						v52 = v51
					}
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_getcwd[0])) = int32(44)
					v52 = v22
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_getcwd[0])) = int32(44)
				v52 = v22
			}
		}
	} else {
		if l1 != 0 {
			v20 = l0
			v21 = l1
			v22 = int32(0)
			v23 = m.Env.X__syscall_getcwd(m, v20, v21)
			mBase = m.M
			if base.Ui32(int32(-4095)) <= base.Ui32(v23) {
				*(*int32)(unsafe.Add(mBase, _c_F_getcwd[0])) = int32(0) - v23
				v31 = int32(-1)
			} else {
				v31 = v23
			}
			if v31 < int32(0) {
				v52 = v22
			} else {
				if v31 != 0 {
					v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
					if v34 == int32(47) {
						if v20 != v12 {
							v52 = v20
						} else {
							v43 = F_strlen(m, v20)
							mBase = m.M
							v45 = v43 + int32(1)
							v46 = F_emscripten_builtin_malloc(m, v45)
							mBase = m.M
							if v46 == int32(0) {
								v51 = int32(0)
							} else {
								v50 = F___memcpy(m, v46, v20, v45)
								mBase = m.M
								v51 = v50
							}
							v52 = v51
						}
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_getcwd[0])) = int32(44)
						v52 = v22
					}
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_getcwd[0])) = int32(44)
					v52 = v22
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_getcwd[0])) = int32(28)
			v52 = int32(0)
		}
	}
	m.G0 = v7
	return v52
}
func F_getdatabaseencoding(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_getdatabaseencoding[0]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v7 = F_DirectFunctionCall1Coll(m, int32(500), int32(0), v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_geterrcode(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_geterrcode[0]))
	if v3 < int32(0) {
		*(*int32)(unsafe.Add(mBase, _c_F_geterrcode[0])) = int32(-1)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_geterrcode_0), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_geterrcode_1), int32(1588), int32(_a_F_geterrcode_2))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v3*int32(100))+uint32(_c_F_geterrcode[1])))
		return v26
	}
}
func F_geterrposition(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_geterrposition[0]))
	if v3 < int32(0) {
		*(*int32)(unsafe.Add(mBase, _c_F_geterrposition[0])) = int32(-1)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_geterrposition_0), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_geterrposition_1), int32(1605), int32(_a_F_geterrposition_2))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v3*int32(100))+uint32(_c_F_geterrposition[1])))
		return v26
	}
}
func F_gettimeofday(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 float64
	_ = v4
	var v5 float64
	_ = v5
	var v7 int64
	_ = v7
	v4 = m.Env.Emscripten_date_now(m)
	mBase = m.M
	v5 = float64(1000)
	v7 = base.I64_trunc_sat_f64_s(base.F64_div(v4, v5))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v7
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = base.I32_trunc_sat_f64_s(base.F64_mul(base.F64_sub(v4, base.F64_convert_i64_s(v7*int64(1000))), v5))
	return
}
func F_gettoken_tsvector(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
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
	var v41 int32
	_ = v41
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
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
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
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
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
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v513 int32
	_ = v513
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v552 int32
	_ = v552
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v590 int32
	_ = v590
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v686 int32
	_ = v686
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v759 int32
	_ = v759
	var v763 int32
	_ = v763
	var v768 int32
	_ = v768
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v785 int32
	_ = v785
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	v7 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(160)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v29 = int32(1)
	v31 = v21
	v32 = v7
	v35 = v7
	v37 = v7
	v38 = v7
	goto L1
L1:
	;
	switch v29 - int32(1) {
	case 0:
		goto L29
	case 1:
		goto L27
	case 2:
		goto L28
	case 3:
		goto L26
	case 4:
		v75 = v31
		goto L30
	case 5:
		goto L21
	case 6:
		goto L20
	case 7:
		goto L31
	default:
		goto L19
	}
L3:
	;
	v815 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v816 = F_pg_mblen_cstr(m, v815)
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L40
	} else {
		goto L321
	}
L4:
	;
	v806 = int32(3)
	v808 = v31
	v809 = v32
	v811 = v35
	v813 = v37
	v814 = v804
	goto L3
L5:
	;
	v804 = int32(2)
	goto L4
L6:
	;
	m.G0 = v19 + int32(160)
	return v785
L7:
	;
	if base.B2i32(v83 == int32(32))|base.B2i32(base.Ui32(v83-int32(9)) < base.Ui32(int32(5))) != 0 {
		v806 = int32(1)
		v808 = v31
		v809 = v32
		v811 = v35
		v813 = v37
		v814 = v38
		goto L3
	} else {
		goto L316
	}
L8:
	;
	v745 = int32(0)
	v746 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v747 = F_errsave_start(m, v746)
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L40
	} else {
		goto L308
	}
L9:
	;
	if base.B2i32(v87&int32(1) == int32(0))|base.B2i32(v83 != int32(34)) != 0 {
		goto L7
	} else {
		goto L307
	}
L10:
	;
	if v83 == int32(124) {
		goto L8
	} else {
		goto L306
	}
L11:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v160 == int32(58) {
		goto L268
	} else {
		goto L269
	}
L12:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v606 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v607 = v31 - v606
	v608 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v605 <= v607+v608 {
		goto L240
	} else {
		goto L241
	}
L13:
	;
	if base.B2i32(v156 == int32(0))|base.B2i32(v160 != int32(34)) != 0 {
		goto L11
	} else {
		goto L239
	}
L14:
	;
	if v160 == int32(124) {
		goto L12
	} else {
		goto L238
	}
L15:
	;
	v568 = m.G0
	v570 = v568 - int32(16)
	m.G0 = v570
	v572 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v573 = F_errsave_start(m, v572)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L40
	} else {
		goto L228
	}
L16:
	;
	if int32(1)<<(uint(v168)%32)&int32(134218145) == int32(0) {
		goto L14
	} else {
		goto L227
	}
L17:
	;
	if int32(1)<<(uint(v107)%32)&int32(134218145) == int32(0) {
		goto L10
	} else {
		goto L226
	}
L18:
	;
	v804 = int32(4)
	goto L4
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L40
	} else {
		goto L223
	}
L20:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v446))))
	switch v447 {
	case 0, 9, 10, 11, 12, 13, 32:
		goto L190
	default:
		goto L189
	case 42, 65, 97:
		goto L194
	case 44:
		v806 = int32(6)
		v808 = v31
		v809 = v32
		v811 = v35
		v813 = v37
		v814 = v38
		goto L3
	case 66, 98:
		goto L193
	case 67, 99:
		goto L192
	case 68, 100:
		goto L191
	}
L21:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304))))
	if base.Ui32((v305-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L142
	} else {
		goto L143
	}
L22:
	;
	if l3 != 0 {
		goto L130
	} else {
		goto L131
	}
L23:
	;
	if l3 != 0 {
		goto L117
	} else {
		goto L118
	}
L24:
	;
	v248 = int32(0)
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v250 = F_errsave_start(m, v249)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L40
	} else {
		goto L108
	}
L25:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v229 = v31 - v45
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v228 <= v229+v230 {
		goto L100
	} else {
		goto L101
	}
L26:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171))))
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+22)))
	if v173 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L27:
	;
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+22)))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	if base.B2i32(v156 == int32(0))&base.B2i32(v160 == int32(92)) != 0 {
		goto L5
	} else {
		goto L70
	}
L28:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	if v111 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L29:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if v83 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L30:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	if v78 != int32(58) {
		goto L22
	} else {
		goto L44
	}
L31:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+22)))
	if v41 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v53 = v31 - v50
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v52 <= v53+v54 {
		goto L37
	} else {
		goto L38
	}
L33:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v50 = v44
	goto L32
L34:
	;
	goto L35
L35:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v47 == int32(39) {
		goto L25
	} else {
		goto L36
	}
L36:
	;
	v50 = v45
	goto L32
L37:
	;
	v58 = v52 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v58
	v60 = F_repalloc(m, v50, v58)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v67 = v31
	goto L39
L39:
	;
	v68 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v67))) = uint8(v68)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v67 == v70 {
		goto L24
	} else {
		goto L42
	}
L40:
	;
	return int32(0)
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v60
	v67 = v60 + v53
	goto L39
L42:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v72 != 0 {
		goto L23
	} else {
		goto L43
	}
L43:
	;
	v75 = v67
	goto L30
L44:
	;
	v806 = int32(6)
	v808 = v75
	v809 = v32
	v811 = v35
	v813 = v37
	v814 = v38
	goto L3
L45:
	;
	v785 = int32(0)
	goto L6
L46:
	;
	goto L47
L47:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+22)))
	if v87&int32(1)|base.B2i32(v83 != int32(39)) == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v806 = int32(4)
	v808 = v31
	v809 = v32
	v811 = v35
	v813 = v37
	v814 = v38
	goto L3
L49:
	;
	goto L50
L50:
	;
	if base.B2i32(v87&int32(1) == int32(0))&base.B2i32(v83 == int32(92)) != 0 {
		goto L5
	} else {
		goto L51
	}
L51:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v103 != int32(1) {
		goto L9
	} else {
		goto L52
	}
L52:
	;
	v107 = v83 - int32(33)
	if base.Ui32(v107) <= base.Ui32(int32(27)) {
		goto L17
	} else {
		goto L53
	}
L53:
	;
	goto L10
L54:
	;
	v114 = int32(0)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v116 = F_errsave_start(m, v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L40
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v137 = v31 - v136
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v135 <= v137+v138 {
		goto L62
	} else {
		goto L63
	}
L57:
	;
	if v116 == int32(0) {
		v785 = v114
		goto L6
	} else {
		goto L58
	}
L58:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L40
	} else {
		goto L59
	}
L59:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v123
	F_errmsg(m, int32(_a_F_gettoken_tsvector_0), v19+int32(32))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L40
	} else {
		goto L60
	}
L60:
	;
	F_errsave_finish(m, v115, int32(_a_F_gettoken_tsvector_1), int32(221), int32(_a_F_gettoken_tsvector_2))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L40
	} else {
		goto L61
	}
L61:
	;
	v785 = v114
	goto L6
L62:
	;
	v142 = v135 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v142
	v144 = F_repalloc(m, v136, v142)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L40
	} else {
		goto L65
	}
L63:
	;
	v149 = v110
	v151 = v31
	goto L64
L64:
	;
	v152 = F_pg_mblen_cstr(m, v149)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L40
	} else {
		goto L66
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v144
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v149 = v148
	v151 = v144 + v137
	goto L64
L66:
	;
	if v152 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	base.MemoryCopy(m, v151, v149, v152)
	goto L69
L68:
	;
	goto L69
L69:
	;
	v806 = v38
	v808 = v152 + v151
	v809 = v32
	v811 = v35
	v813 = v37
	v814 = v38
	goto L3
L70:
	;
	switch v160 {
	case 0, 9, 10, 11, 12, 13, 32:
		goto L12
	default:
		goto L71
	}
L71:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v164 != int32(1) {
		goto L13
	} else {
		goto L72
	}
L72:
	;
	v168 = v160 - int32(33)
	if base.Ui32(v168) <= base.Ui32(int32(27)) {
		goto L16
	} else {
		goto L73
	}
L73:
	;
	goto L14
L74:
	;
	if v172 == int32(39) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	if v172 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L77:
	;
	v806 = int32(8)
	v808 = v31
	v809 = v32
	v811 = v35
	v813 = v37
	v814 = v38
	goto L3
L78:
	;
	goto L79
L79:
	;
	if v172 == int32(92) {
		goto L18
	} else {
		goto L80
	}
L80:
	;
	goto L76
L81:
	;
	v183 = int32(0)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v185 = F_errsave_start(m, v184)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L40
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v209 = v31 - v208
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v207 <= v209+v210 {
		goto L92
	} else {
		goto L93
	}
L84:
	;
	if v185 == int32(0) {
		v785 = v183
		goto L6
	} else {
		goto L85
	}
L85:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L40
	} else {
		goto L86
	}
L86:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+80)) = v193
	if v192 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v197 = int32(_a_F_gettoken_tsvector_3)
	goto L89
L88:
	;
	v197 = int32(_a_F_gettoken_tsvector_4)
	goto L89
L89:
	;
	F_errmsg(m, v197, v19+int32(80))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L40
	} else {
		goto L90
	}
L90:
	;
	F_errsave_finish(m, v184, int32(_a_F_gettoken_tsvector_1), int32(148), int32(_a_F_gettoken_tsvector_5))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L40
	} else {
		goto L91
	}
L91:
	;
	v785 = v183
	goto L6
L92:
	;
	v214 = v207 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v214
	v216 = F_repalloc(m, v208, v214)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L40
	} else {
		goto L95
	}
L93:
	;
	v221 = v171
	v222 = v31
	goto L94
L94:
	;
	v223 = F_pg_mblen_cstr(m, v221)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L40
	} else {
		goto L96
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v216
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v221 = v220
	v222 = v216 + v209
	goto L94
L96:
	;
	if v223 != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	base.MemoryCopy(m, v222, v221, v223)
	goto L99
L98:
	;
	goto L99
L99:
	;
	v806 = int32(4)
	v808 = v223 + v222
	v809 = v32
	v811 = v35
	v813 = v37
	v814 = v38
	goto L3
L100:
	;
	v234 = v228 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v234
	v236 = F_repalloc(m, v45, v234)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L40
	} else {
		goto L103
	}
L101:
	;
	v241 = v46
	v242 = v31
	goto L102
L102:
	;
	v243 = F_pg_mblen_cstr(m, v241)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L40
	} else {
		goto L104
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v236
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v241 = v240
	v242 = v236 + v229
	goto L102
L104:
	;
	if v243 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	base.MemoryCopy(m, v242, v241, v243)
	goto L107
L106:
	;
	goto L107
L107:
	;
	v806 = int32(4)
	v808 = v243 + v242
	v809 = v32
	v811 = v35
	v813 = v37
	v814 = v38
	goto L3
L108:
	;
	if v250 == int32(0) {
		v785 = v248
		goto L6
	} else {
		goto L109
	}
L109:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L40
	} else {
		goto L110
	}
L110:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+96)) = v258
	if v257 != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v262 = int32(_a_F_gettoken_tsvector_3)
	goto L113
L112:
	;
	v262 = int32(_a_F_gettoken_tsvector_4)
	goto L113
L113:
	;
	F_errmsg(m, v262, v19+int32(96))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L40
	} else {
		goto L114
	}
L114:
	;
	F_errsave_finish(m, v249, int32(_a_F_gettoken_tsvector_1), int32(148), int32(_a_F_gettoken_tsvector_5))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L40
	} else {
		goto L115
	}
L115:
	;
	v785 = v248
	goto L6
L116:
	;
	if l1 != 0 {
		goto L122
	} else {
		goto L123
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v32
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v35
	goto L116
L118:
	;
	goto L119
L119:
	;
	if v32 == int32(0) {
		goto L116
	} else {
		goto L120
	}
L120:
	;
	F_pfree(m, v32)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L40
	} else {
		goto L121
	}
L121:
	;
	goto L116
L122:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v278
	goto L124
L123:
	;
	goto L124
L124:
	;
	if l2 != 0 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v67 - v280
	goto L127
L126:
	;
	goto L127
L127:
	;
	v283 = int32(1)
	if l5 == int32(0) {
		v785 = v283
		goto L6
	} else {
		goto L128
	}
L128:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v286
	v785 = v283
	goto L6
L129:
	;
	if l1 != 0 {
		goto L135
	} else {
		goto L136
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v32
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v35
	goto L129
L131:
	;
	goto L132
L132:
	;
	if v32 == int32(0) {
		goto L129
	} else {
		goto L133
	}
L133:
	;
	F_pfree(m, v32)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L40
	} else {
		goto L134
	}
L134:
	;
	goto L129
L135:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v294
	goto L137
L136:
	;
	goto L137
L137:
	;
	if l2 != 0 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v75 - v296
	goto L140
L139:
	;
	goto L140
L140:
	;
	v299 = int32(1)
	if l5 == int32(0) {
		v785 = v299
		goto L6
	} else {
		goto L141
	}
L141:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v302
	v785 = v299
	goto L6
L142:
	;
	if v37 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L143:
	;
	goto L144
L144:
	;
	v421 = int32(0)
	v422 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v423 = F_errsave_start(m, v422)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L40
	} else {
		goto L181
	}
L145:
	;
	v333 = v329 + v330<<(uint(int32(1))%32)
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v339 = v335
	goto L155
L146:
	;
	v317 = F_palloc(m, int32(8))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L40
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	if v35+int32(1) < v37 {
		goto L150
	} else {
		goto L151
	}
L149:
	;
	v328 = int32(4)
	v329 = v317
	v330 = int32(0)
	goto L145
L150:
	;
	v328 = v37
	v329 = v32
	v330 = v35
	goto L145
L151:
	;
	goto L152
L152:
	;
	v326 = F_repalloc(m, v32, v37<<(uint(int32(2))%32))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L40
	} else {
		goto L153
	}
L153:
	;
	v328 = v37 << (uint(int32(1)) % 32)
	v329 = v326
	v330 = v35
	goto L145
L154:
	;
	v384 = int32(_a_F_gettoken_tsvector_6)
	if v384 < v383 {
		goto L170
	} else {
		goto L171
	}
L155:
	;
	v344 = v339 + int32(1)
	v345 = int32(*(*int8)(unsafe.Add(mBase, uint32(v339))))
	v346 = F___isspace(m, v345)
	mBase = m.M
	if v346 != 0 {
		v339 = v344
		goto L155
	} else {
		goto L157
	}
L156:
	;
	v347 = int32(1)
	switch v345&int32(255) - int32(43) {
	case 0:
		v353 = v347
		goto L159
	default:
		v355 = v345
		v356 = v339
		v357 = v347
		goto L158
	case 2:
		goto L160
	}
L157:
	;
	goto L156
L158:
	;
	v358 = int32(0)
	v360 = v355 - int32(48)
	if base.Ui32(v360) <= base.Ui32(int32(9)) {
		goto L161
	} else {
		goto L162
	}
L159:
	;
	v354 = int32(*(*int8)(unsafe.Add(mBase, uint32(v344))))
	v355 = v354
	v356 = v344
	v357 = v353
	goto L158
L160:
	;
	v353 = int32(0)
	goto L159
L161:
	;
	v363 = v358
	v364 = v360
	v365 = v356
	goto L164
L162:
	;
	v377 = v358
	goto L163
L163:
	;
	if v357 != 0 {
		goto L167
	} else {
		goto L168
	}
L164:
	;
	v367 = int32(10)
	v369 = v363*v367 - v364
	v370 = int32(*(*int8)(unsafe.Add(mBase, uint32(v365)+1)))
	v374 = v370 - int32(48)
	if base.Ui32(v374) < base.Ui32(v367) {
		v363 = v369
		v364 = v374
		v365 = v365 + int32(1)
		goto L164
	} else {
		goto L166
	}
L165:
	;
	v377 = v369
	goto L163
L166:
	;
	goto L165
L167:
	;
	v383 = int32(0) - v377
	goto L169
L168:
	;
	v383 = v377
	goto L169
L169:
	;
	goto L154
L170:
	;
	v388 = int32(_a_F_gettoken_tsvector_6)
	goto L172
L171:
	;
	v388 = v383 & v384
	goto L172
L172:
	;
	v389 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v333))))
	v392 = v388 | v389&int32(_a_F_gettoken_tsvector_7)
	*(*uint16)(unsafe.Add(mBase, uint32(v333))) = uint16(v392)
	if v388 == int32(0) {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v396 = int32(0)
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v398 = F_errsave_start(m, v397)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L40
	} else {
		goto L176
	}
L174:
	;
	goto L175
L175:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v333))) = uint16(v388)
	v806 = int32(7)
	v808 = v31
	v809 = v329
	v811 = v330 + int32(1)
	v813 = v328
	v814 = v38
	goto L3
L176:
	;
	if v398 == int32(0) {
		v785 = v396
		goto L6
	} else {
		goto L177
	}
L177:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L40
	} else {
		goto L178
	}
L178:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+112)) = v405
	F_errmsg(m, int32(_a_F_gettoken_tsvector_8), v19+int32(112))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L40
	} else {
		goto L179
	}
L179:
	;
	F_errsave_finish(m, v397, int32(_a_F_gettoken_tsvector_1), int32(335), int32(_a_F_gettoken_tsvector_2))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L40
	} else {
		goto L180
	}
L180:
	;
	v785 = v396
	goto L6
L181:
	;
	if v423 == int32(0) {
		v785 = v421
		goto L6
	} else {
		goto L182
	}
L182:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L40
	} else {
		goto L183
	}
L183:
	;
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+128)) = v431
	if v430 != 0 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v435 = int32(_a_F_gettoken_tsvector_3)
	goto L186
L185:
	;
	v435 = int32(_a_F_gettoken_tsvector_4)
	goto L186
L186:
	;
	F_errmsg(m, v435, v19+int32(128))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L40
	} else {
		goto L187
	}
L187:
	;
	F_errsave_finish(m, v422, int32(_a_F_gettoken_tsvector_1), int32(148), int32(_a_F_gettoken_tsvector_5))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L40
	} else {
		goto L188
	}
L188:
	;
	v785 = v421
	goto L6
L189:
	;
	if base.Ui32(int32(10)) <= base.Ui32((v447-int32(48))&int32(255)) {
		goto L15
	} else {
		goto L222
	}
L190:
	;
	if l3 != 0 {
		goto L210
	} else {
		goto L211
	}
L191:
	;
	v513 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32+v35<<(uint(int32(1))%32)-int32(2)))))
	if base.Ui32(int32(_a_F_gettoken_tsvector_9)) <= base.Ui32(v513) {
		goto L15
	} else {
		goto L208
	}
L192:
	;
	v500 = v32 + v35<<(uint(int32(1))%32) - int32(2)
	v501 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v500))))
	if base.Ui32(int32(_a_F_gettoken_tsvector_9)) <= base.Ui32(v501) {
		goto L15
	} else {
		goto L207
	}
L193:
	;
	v488 = v32 + v35<<(uint(int32(1))%32) - int32(2)
	v489 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v488))))
	if base.Ui32(int32(_a_F_gettoken_tsvector_9)) <= base.Ui32(v489) {
		goto L15
	} else {
		goto L206
	}
L194:
	;
	v452 = v32 + v35<<(uint(int32(1))%32) - int32(2)
	v453 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v452))))
	if base.Ui32(int32(_a_F_gettoken_tsvector_9)) <= base.Ui32(v453) {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v456 = int32(0)
	v457 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v458 = F_errsave_start(m, v457)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L40
	} else {
		goto L198
	}
L196:
	;
	goto L197
L197:
	;
	v481 = v453 | int32(_a_F_gettoken_tsvector_7)
	*(*uint16)(unsafe.Add(mBase, uint32(v452))) = uint16(v481)
	v806 = int32(7)
	v808 = v31
	v809 = v32
	v811 = v35
	v813 = v37
	v814 = v38
	goto L3
L198:
	;
	if v458 == int32(0) {
		v785 = v456
		goto L6
	} else {
		goto L199
	}
L199:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L40
	} else {
		goto L200
	}
L200:
	;
	v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+144)) = v466
	if v465 != 0 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v470 = int32(_a_F_gettoken_tsvector_3)
	goto L203
L202:
	;
	v470 = int32(_a_F_gettoken_tsvector_4)
	goto L203
L203:
	;
	F_errmsg(m, v470, v19+int32(144))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L40
	} else {
		goto L204
	}
L204:
	;
	F_errsave_finish(m, v457, int32(_a_F_gettoken_tsvector_1), int32(148), int32(_a_F_gettoken_tsvector_5))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L40
	} else {
		goto L205
	}
L205:
	;
	v785 = v456
	goto L6
L206:
	;
	v493 = v489 | int32(_a_F_gettoken_tsvector_10)
	*(*uint16)(unsafe.Add(mBase, uint32(v488))) = uint16(v493)
	v806 = int32(7)
	v808 = v31
	v809 = v32
	v811 = v35
	v813 = v37
	v814 = v38
	goto L3
L207:
	;
	v505 = v501 | int32(_a_F_gettoken_tsvector_9)
	*(*uint16)(unsafe.Add(mBase, uint32(v500))) = uint16(v505)
	v806 = int32(7)
	v808 = v31
	v809 = v32
	v811 = v35
	v813 = v37
	v814 = v38
	goto L3
L208:
	;
	v806 = int32(7)
	v808 = v31
	v809 = v32
	v811 = v35
	v813 = v37
	v814 = v38
	goto L3
L209:
	;
	if l1 != 0 {
		goto L215
	} else {
		goto L216
	}
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v32
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v35
	goto L209
L211:
	;
	goto L212
L212:
	;
	if v32 == int32(0) {
		goto L209
	} else {
		goto L213
	}
L213:
	;
	F_pfree(m, v32)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L40
	} else {
		goto L214
	}
L214:
	;
	goto L209
L215:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v523
	goto L217
L216:
	;
	goto L217
L217:
	;
	if l2 != 0 {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v31 - v525
	goto L220
L219:
	;
	goto L220
L220:
	;
	v528 = int32(1)
	if l5 == int32(0) {
		v785 = v528
		goto L6
	} else {
		goto L221
	}
L221:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v531
	v785 = v528
	goto L6
L222:
	;
	v806 = int32(7)
	v808 = v31
	v809 = v32
	v811 = v35
	v813 = v37
	v814 = v38
	goto L3
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v29
	F_errmsg_internal(m, int32(_a_F_gettoken_tsvector_11), v19)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L40
	} else {
		goto L224
	}
L224:
	;
	F_errfinish(m, int32(_a_F_gettoken_tsvector_1), int32(378), int32(_a_F_gettoken_tsvector_2))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L40
	} else {
		goto L225
	}
L225:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L226:
	;
	goto L8
L227:
	;
	goto L12
L228:
	;
	if v573 != 0 {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L40
	} else {
		goto L232
	}
L230:
	;
	goto L231
L231:
	;
	m.G0 = v570 + int32(16)
	v785 = int32(0)
	goto L6
L232:
	;
	v578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
	v579 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v570))) = v579
	if v578 != 0 {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	v583 = int32(_a_F_gettoken_tsvector_3)
	goto L235
L234:
	;
	v583 = int32(_a_F_gettoken_tsvector_4)
	goto L235
L235:
	;
	F_errmsg(m, v583, v570)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L40
	} else {
		goto L236
	}
L236:
	;
	F_errsave_finish(m, v572, int32(_a_F_gettoken_tsvector_1), int32(148), int32(_a_F_gettoken_tsvector_5))
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L40
	} else {
		goto L237
	}
L237:
	;
	goto L231
L238:
	;
	goto L13
L239:
	;
	goto L12
L240:
	;
	v612 = v605 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v612
	v614 = F_repalloc(m, v606, v612)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L40
	} else {
		goto L243
	}
L241:
	;
	v618 = v606
	v620 = v31
	goto L242
L242:
	;
	if v618 == v620 {
		goto L244
	} else {
		goto L245
	}
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v614
	v618 = v614
	v620 = v614 + v607
	goto L242
L244:
	;
	v622 = int32(0)
	v623 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v624 = F_errsave_start(m, v623)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L40
	} else {
		goto L247
	}
L245:
	;
	goto L246
L246:
	;
	v646 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v620))) = uint8(v646)
	if l3 != 0 {
		goto L256
	} else {
		goto L257
	}
L247:
	;
	if v624 == int32(0) {
		v785 = v622
		goto L6
	} else {
		goto L248
	}
L248:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L40
	} else {
		goto L249
	}
L249:
	;
	v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
	v632 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = v632
	if v631 != 0 {
		goto L250
	} else {
		goto L251
	}
L250:
	;
	v636 = int32(_a_F_gettoken_tsvector_3)
	goto L252
L251:
	;
	v636 = int32(_a_F_gettoken_tsvector_4)
	goto L252
L252:
	;
	F_errmsg(m, v636, v19+int32(48))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L40
	} else {
		goto L253
	}
L253:
	;
	F_errsave_finish(m, v623, int32(_a_F_gettoken_tsvector_1), int32(148), int32(_a_F_gettoken_tsvector_5))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L40
	} else {
		goto L254
	}
L254:
	;
	v785 = v622
	goto L6
L255:
	;
	if l1 != 0 {
		goto L261
	} else {
		goto L262
	}
L256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v32
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v35
	goto L255
L257:
	;
	goto L258
L258:
	;
	if v32 == int32(0) {
		goto L255
	} else {
		goto L259
	}
L259:
	;
	F_pfree(m, v32)
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L40
	} else {
		goto L260
	}
L260:
	;
	goto L255
L261:
	;
	v654 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v654
	goto L263
L262:
	;
	goto L263
L263:
	;
	if l2 != 0 {
		goto L264
	} else {
		goto L265
	}
L264:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v620 - v656
	goto L266
L265:
	;
	goto L266
L266:
	;
	v659 = int32(1)
	if l5 == int32(0) {
		v785 = v659
		goto L6
	} else {
		goto L267
	}
L267:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v662
	v785 = v659
	goto L6
L268:
	;
	if v31 == v664 {
		goto L271
	} else {
		goto L272
	}
L269:
	;
	goto L270
L270:
	;
	v714 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v715 = v31 - v664
	v716 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v714 <= v715+v716 {
		goto L298
	} else {
		goto L299
	}
L271:
	;
	v668 = int32(0)
	v669 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v670 = F_errsave_start(m, v669)
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L40
	} else {
		goto L274
	}
L272:
	;
	goto L273
L273:
	;
	v692 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v692)
	v694 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v694 != int32(1) {
		goto L282
	} else {
		goto L283
	}
L274:
	;
	if v670 == int32(0) {
		v785 = v668
		goto L6
	} else {
		goto L275
	}
L275:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L40
	} else {
		goto L276
	}
L276:
	;
	v677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
	v678 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = v678
	if v677 != 0 {
		goto L277
	} else {
		goto L278
	}
L277:
	;
	v682 = int32(_a_F_gettoken_tsvector_3)
	goto L279
L278:
	;
	v682 = int32(_a_F_gettoken_tsvector_4)
	goto L279
L279:
	;
	F_errmsg(m, v682, v19-int32(-64))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L40
	} else {
		goto L280
	}
L280:
	;
	F_errsave_finish(m, v669, int32(_a_F_gettoken_tsvector_1), int32(148), int32(_a_F_gettoken_tsvector_5))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L40
	} else {
		goto L281
	}
L281:
	;
	v785 = v668
	goto L6
L282:
	;
	v806 = int32(6)
	v808 = v31
	v809 = v32
	v811 = v35
	v813 = v37
	v814 = v38
	goto L3
L283:
	;
	goto L284
L284:
	;
	if l3 != 0 {
		goto L286
	} else {
		goto L287
	}
L285:
	;
	if l1 != 0 {
		goto L291
	} else {
		goto L292
	}
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v32
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v35
	goto L285
L287:
	;
	goto L288
L288:
	;
	if v32 == int32(0) {
		goto L285
	} else {
		goto L289
	}
L289:
	;
	F_pfree(m, v32)
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L40
	} else {
		goto L290
	}
L290:
	;
	goto L285
L291:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v704
	goto L293
L292:
	;
	goto L293
L293:
	;
	if l2 != 0 {
		goto L294
	} else {
		goto L295
	}
L294:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v31 - v706
	goto L296
L295:
	;
	goto L296
L296:
	;
	v709 = int32(1)
	if l5 == int32(0) {
		v785 = v709
		goto L6
	} else {
		goto L297
	}
L297:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v712
	v785 = v709
	goto L6
L298:
	;
	v720 = v714 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v720
	v722 = F_repalloc(m, v664, v720)
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L40
	} else {
		goto L301
	}
L299:
	;
	v727 = v159
	v728 = v31
	goto L300
L300:
	;
	v729 = F_pg_mblen_cstr(m, v727)
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L40
	} else {
		goto L302
	}
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v722
	v726 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v727 = v726
	v728 = v722 + v715
	goto L300
L302:
	;
	if v729 != 0 {
		goto L303
	} else {
		goto L304
	}
L303:
	;
	base.MemoryCopy(m, v728, v727, v729)
	goto L305
L304:
	;
	goto L305
L305:
	;
	v806 = int32(2)
	v808 = v729 + v728
	v809 = v32
	v811 = v35
	v813 = v37
	v814 = v38
	goto L3
L306:
	;
	goto L9
L307:
	;
	goto L8
L308:
	;
	if v747 == int32(0) {
		v785 = v745
		goto L6
	} else {
		goto L309
	}
L309:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L40
	} else {
		goto L310
	}
L310:
	;
	v754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
	v755 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v755
	if v754 != 0 {
		goto L311
	} else {
		goto L312
	}
L311:
	;
	v759 = int32(_a_F_gettoken_tsvector_3)
	goto L313
L312:
	;
	v759 = int32(_a_F_gettoken_tsvector_4)
	goto L313
L313:
	;
	F_errmsg(m, v759, v19+int32(16))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L40
	} else {
		goto L314
	}
L314:
	;
	F_errsave_finish(m, v746, int32(_a_F_gettoken_tsvector_1), int32(148), int32(_a_F_gettoken_tsvector_5))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L40
	} else {
		goto L315
	}
L315:
	;
	v785 = v745
	goto L6
L316:
	;
	v777 = F_pg_mblen_cstr(m, v82)
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L40
	} else {
		goto L317
	}
L317:
	;
	if v777 != 0 {
		goto L318
	} else {
		goto L319
	}
L318:
	;
	base.MemoryCopy(m, v31, v82, v777)
	goto L320
L319:
	;
	goto L320
L320:
	;
	v806 = int32(2)
	v808 = v777 + v31
	v809 = v32
	v811 = v35
	v813 = v37
	v814 = v38
	goto L3
L321:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v816 + v818
	v29 = v806
	v31 = v808
	v32 = v809
	v35 = v811
	v37 = v813
	v38 = v814
	goto L1
}
func F_ginadjustmembers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	var v13 int32
	_ = v13
	Fn13930(m, l0, l1, l2, l3, int32(_a_F_ginadjustmembers_0), int32(324), int32(_a_F_ginadjustmembers_1), int32(_a_F_ginadjustmembers_2), int32(12), int32(242), int32(7))
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		return
	}
}
func F_ginbuildphasename(m *base.Module, l0 int64) int32 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v3 = l0 - int64(1)
	if base.Ui64(v3) <= base.Ui64(int64(5)) {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v3)<<(uint(int32(2))%32))+uint32(_c_F_ginbuildphasename[0])))
		v11 = v9
	} else {
		v11 = int32(0)
	}
	return v11
}
func F_ginendscan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v12 int32
	_ = v12
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_ginFreeScanKeys(m, v2)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
		F_MemoryContextDelete(m, v5)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(v2)+uint32(_c_F_ginendscan[0])))
			F_MemoryContextDelete(m, v8)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				F_pfree(m, v2)
				mBase = m.M
				v12 = m.ExcPending
				if v12 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_ginint4_consistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
	switch v14 - int32(3) {
	case 0:
		goto L3
	default:
		goto L4
	case 3:
		goto L7
	case 4, 10:
		goto L6
	case 5, 11:
		goto L8
	case 17:
		goto L5
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return v83
L2:
	;
	v83 = int32(1)
	goto L1
L3:
	;
	v78 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11))) = uint8(v78)
	goto L2
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L23
	} else {
		goto L26
	}
L5:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v57 = F_pg_detoast_datum(m, v56)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L23
	} else {
		goto L24
	}
L6:
	;
	v38 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11))) = uint8(v38)
	if v12 <= v38 {
		goto L2
	} else {
		goto L16
	}
L7:
	;
	v20 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11))) = uint8(v20)
	v23 = int32(0)
	if v12 <= v23 {
		v83 = v20
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v17 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11))) = uint8(v17)
	v83 = v17
	goto L1
L9:
	;
	v26 = v23
	goto L10
L10:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+v13))))
	if v33 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v83 = int32(0)
	goto L1
L12:
	;
	v35 = v26 + int32(1)
	if v12 != v35 {
		v26 = v35
		goto L10
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	goto L11
L15:
	;
	v83 = v20
	goto L1
L16:
	;
	v43 = v38
	goto L17
L17:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+v13))))
	if v50 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v83 = int32(0)
	goto L1
L19:
	;
	v51 = int32(1)
	v53 = v43 + v51
	if v12 != v53 {
		v43 = v53
		goto L17
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	goto L18
L22:
	;
	v83 = v51
	goto L1
L23:
	;
	return int32(0)
L24:
	;
	v61 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11))) = uint8(v61)
	v63 = F_gin_bool_consistent(m, v57, v13)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v83 = v63
	goto L1
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v14
	F_errmsg_internal(m, int32(_a_F_ginint4_consistent_0), v9)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_ginint4_consistent_1), int32(176), int32(_a_F_ginint4_consistent_2))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L23
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ginrescan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_ginFreeScanKeys(m, v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		if l1 == int32(0) {
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v11 <= int32(0) {
			} else {
				v15 = v11 * int32(48)
				if v15 == int32(0) {
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					base.MemoryCopy(m, v18, l1, v15)
				}
			}
		}
		return
	}
}
func F_ginvalidate(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int64
	_ = v133
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v297 int32
	_ = v297
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
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
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
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
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v442 int32
	_ = v442
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v564 int32
	_ = v564
	var v571 int32
	_ = v571
	var v582 int64
	_ = v582
	var v583 int64
	_ = v583
	var v588 int32
	_ = v588
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v607 int32
	_ = v607
	var v616 int32
	_ = v616
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v624 int64
	_ = v624
	var v628 int64
	_ = v628
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v649 int32
	_ = v649
	var v659 int32
	_ = v659
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	v18 = m.G0
	v20 = v18 - int32(272)
	m.G0 = v20
	v23 = F_SearchSysCache1(m, int32(14), l0)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v39)+40))
	if int32(0) < v297 {
		goto L62
	} else {
		goto L63
	}
L2:
	;
	return int32(0)
L3:
	;
	if v23 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+22)))
	v29 = v27 + v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+84))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)+92))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+80))
	v33 = F_get_opfamily_name(m, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L2
	} else {
		goto L59
	}
L7:
	;
	v37 = int32(0)
	v39 = F_SearchSysCacheList(m, int32(4), int32(1), v32, v37, v37)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v41 = int32(1)
	v44 = int32(0)
	v46 = F_SearchSysCacheList(m, int32(5), v41, v32, v44, v44)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v46)+40))
	if v48 <= int32(0) {
		v285 = v41
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v31 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v51 = v31
	goto L13
L12:
	;
	v51 = v30
	goto L13
L13:
	;
	v61 = int32(0)
	v64 = v41
	goto L14
L14:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v46+int32(48)+v61<<(uint(int32(2))%32))))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+56))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+22)))
	v82 = v80 + v81
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+8))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	if v83 == v84 {
		v113 = v64
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v285 = v261
	goto L1
L16:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v82)+8))
	if v114 != v30 {
		v261 = v113
		goto L24
	} else {
		goto L25
	}
L17:
	;
	v86 = int32(0)
	v89 = F_errstart(m, int32(17), v86)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	if v89 == int32(0) {
		v113 = v86
		goto L16
	} else {
		goto L19
	}
L19:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	v97 = F_format_procedure(m, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+264)) = v97
	*(*int32)(unsafe.Add(mBase, uint32(v20)+260)) = int32(_a_F_ginvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+256)) = v33
	F_errmsg(m, int32(_a_F_ginvalidate_1), v20+int32(256))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_ginvalidate_2), int32(85), int32(_a_F_ginvalidate_3))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	v113 = v86
	goto L16
L24:
	;
	v264 = v61 + int32(1)
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v46)+40))
	if v264 < v265 {
		v61 = v264
		v64 = v261
		goto L14
	} else {
		goto L58
	}
L25:
	;
	v116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82)+16)))
	switch v116 - int32(1) {
	case 0:
		goto L35
	case 1:
		goto L28
	case 2:
		goto L34
	case 3:
		goto L33
	case 4:
		goto L32
	case 5:
		goto L31
	case 6:
		goto L30
	default:
		goto L29
	}
L26:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L2
	} else {
		goto L54
	}
L27:
	;
	v229 = int32(0)
	v232 = F_errstart(m, int32(17), v229)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L2
	} else {
		goto L52
	}
L28:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+132)) = int64(9796820404457)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+128)) = v30
	v227 = F_check_amproc_signature(m, v217, int32(2281), int32(0), int32(2), int32(3), v20+int32(128))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L2
	} else {
		goto L50
	}
L29:
	;
	v208 = int32(0)
	v211 = F_errstart(m, int32(17), v208)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L2
	} else {
		goto L48
	}
L30:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	v204 = F_check_amoptsproc_signature(m, v203)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L2
	} else {
		goto L46
	}
L31:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+244)) = int64(9796820404457)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+236)) = int64(9796820402199)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+232)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v20)+224)) = int64(90194315497)
	v195 = int32(7)
	v199 = F_check_amproc_signature(m, v185, int32(18), int32(0), v195, v195, v20+int32(224))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L2
	} else {
		goto L44
	}
L32:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+216)) = int64(9796820402197)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+212)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v20)+208)) = v51
	v177 = int32(4)
	v181 = F_check_amproc_signature(m, v170, int32(23), int32(0), v177, v177, v20+int32(208))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L2
	} else {
		goto L42
	}
L33:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v20+int32(204)))) = int32(2281)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+196)) = int64(9796820404457)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+188)) = int64(9796820402199)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+184)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v20)+176)) = int64(90194315497)
	v166 = F_check_amproc_signature(m, v150, int32(16), int32(0), int32(6), int32(8), v20+int32(176))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L2
	} else {
		goto L40
	}
L34:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	v133 = int64(9796820404457)
	*(*int64)(unsafe.Add(mBase, uint32(v20+int32(164)))) = v133
	*(*int64)(unsafe.Add(mBase, uint32(v20)+156)) = v133
	*(*int64)(unsafe.Add(mBase, uint32(v20)+148)) = int64(90194315497)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+144)) = v30
	v146 = F_check_amproc_signature(m, v132, int32(2281), int32(0), int32(5), int32(7), v20+int32(144))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L2
	} else {
		goto L38
	}
L35:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+116)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v20)+112)) = v51
	v124 = int32(2)
	v128 = F_check_amproc_signature(m, v119, int32(23), int32(0), v124, v124, v20+int32(112))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	if v128 == int32(0) {
		goto L27
	} else {
		goto L37
	}
L37:
	;
	v261 = v113
	goto L24
L38:
	;
	if v146 == int32(0) {
		goto L27
	} else {
		goto L39
	}
L39:
	;
	v261 = v113
	goto L24
L40:
	;
	if v166 == int32(0) {
		goto L27
	} else {
		goto L41
	}
L41:
	;
	v261 = v113
	goto L24
L42:
	;
	if v181 == int32(0) {
		goto L27
	} else {
		goto L43
	}
L43:
	;
	v261 = v113
	goto L24
L44:
	;
	if v199 == int32(0) {
		goto L27
	} else {
		goto L45
	}
L45:
	;
	v261 = v113
	goto L24
L46:
	;
	if v204 == int32(0) {
		goto L27
	} else {
		goto L47
	}
L47:
	;
	v261 = v113
	goto L24
L48:
	;
	if v211 == int32(0) {
		v261 = v208
		goto L24
	} else {
		goto L49
	}
L49:
	;
	v238 = int32(145)
	v239 = int32(_a_F_ginvalidate_4)
	goto L26
L50:
	;
	if v227 != 0 {
		v261 = v113
		goto L24
	} else {
		goto L51
	}
L51:
	;
	goto L27
L52:
	;
	if v232 == int32(0) {
		v261 = v229
		goto L24
	} else {
		goto L53
	}
L53:
	;
	v238 = int32(157)
	v239 = int32(_a_F_ginvalidate_5)
	goto L26
L54:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	v244 = F_format_procedure(m, v243)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L2
	} else {
		goto L55
	}
L55:
	;
	v246 = int32(*(*int16)(unsafe.Add(mBase, uint32(v82)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+108)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v20)+104)) = v244
	*(*int32)(unsafe.Add(mBase, uint32(v20)+100)) = int32(_a_F_ginvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+96)) = v33
	F_errmsg(m, v239, v20+int32(96))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L2
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_ginvalidate_2), v238, int32(_a_F_ginvalidate_3))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L2
	} else {
		goto L57
	}
L57:
	;
	v261 = int32(0)
	goto L24
L58:
	;
	goto L15
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = l0
	F_errmsg_internal(m, int32(_a_F_ginvalidate_6), v20)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L2
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_ginvalidate_2), int32(51), int32(_a_F_ginvalidate_3))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L2
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	v305 = int32(0)
	v308 = v285
	goto L65
L63:
	;
	v442 = v285
	goto L64
L64:
	;
	v454 = F_identify_opfamily_groups(m, v39, v46)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L2
	} else {
		goto L97
	}
L65:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v39+int32(48)+v305<<(uint(int32(2))%32))))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v323)+56))
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v324)+22)))
	v326 = v324 + v325
	v327 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v326)+16)))
	if base.Ui32(int32(_a_F_ginvalidate_7)) < base.Ui32((v327+int32(-64))&int32(_a_F_ginvalidate_8)) {
		v363 = v308
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v442 = v432
	goto L64
L67:
	;
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326)+18)))
	if v365 == int32(115) {
		goto L76
	} else {
		goto L77
	}
L68:
	;
	v334 = int32(0)
	v337 = F_errstart(m, int32(17), v334)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L2
	} else {
		goto L69
	}
L69:
	;
	if v337 == int32(0) {
		v363 = v334
		goto L67
	} else {
		goto L70
	}
L70:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L2
	} else {
		goto L71
	}
L71:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v326)+20))
	v345 = F_format_operator(m, v344)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L2
	} else {
		goto L72
	}
L72:
	;
	v347 = int32(*(*int16)(unsafe.Add(mBase, uint32(v326)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+92)) = v347
	*(*int32)(unsafe.Add(mBase, uint32(v20)+88)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v20)+84)) = int32(_a_F_ginvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = v33
	F_errmsg(m, int32(_a_F_ginvalidate_9), v20+int32(80))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L2
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(_a_F_ginvalidate_2), int32(176), int32(_a_F_ginvalidate_3))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L2
	} else {
		goto L74
	}
L74:
	;
	v363 = v334
	goto L67
L75:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v326)+20))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v326)+8))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v326)+12))
	v403 = F_check_amop_signature(m, v399, int32(16), v401, v402)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L2
	} else {
		goto L87
	}
L76:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v326)+28))
	if v368 == int32(0) {
		v398 = v363
		goto L75
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v371 = int32(0)
	v374 = F_errstart(m, int32(17), v371)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L2
	} else {
		goto L80
	}
L79:
	;
	goto L78
L80:
	;
	if v374 == int32(0) {
		v398 = v371
		goto L75
	} else {
		goto L81
	}
L81:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L2
	} else {
		goto L82
	}
L82:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v326)+20))
	v382 = F_format_operator(m, v381)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L2
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+72)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v20)+68)) = int32(_a_F_ginvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v33
	F_errmsg(m, int32(_a_F_ginvalidate_10), v20-int32(-64))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L2
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_ginvalidate_2), int32(188), int32(_a_F_ginvalidate_3))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L2
	} else {
		goto L85
	}
L85:
	;
	v398 = v371
	goto L75
L86:
	;
	v434 = v305 + int32(1)
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v39)+40))
	if v434 < v435 {
		v305 = v434
		v308 = v432
		goto L65
	} else {
		goto L95
	}
L87:
	;
	if v403 != 0 {
		v432 = v398
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v405 = int32(0)
	v408 = F_errstart(m, int32(17), v405)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L2
	} else {
		goto L89
	}
L89:
	;
	if v408 == int32(0) {
		v432 = v405
		goto L86
	} else {
		goto L90
	}
L90:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L2
	} else {
		goto L91
	}
L91:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v326)+20))
	v416 = F_format_operator(m, v415)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L2
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = v416
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = int32(_a_F_ginvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v33
	F_errmsg(m, int32(_a_F_ginvalidate_11), v20+int32(48))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L2
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(_a_F_ginvalidate_2), int32(201), int32(_a_F_ginvalidate_3))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L2
	} else {
		goto L94
	}
L94:
	;
	v432 = v405
	goto L86
L95:
	;
	goto L66
L96:
	;
	v564 = v29 + int32(8)
	v571 = v442
	v582 = int64(1)
	goto L132
L97:
	;
	if v454 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v546 = int32(0)
	goto L96
L99:
	;
	goto L100
L100:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v454)+4))
	if v459 <= int32(0) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v546 = int32(0)
	goto L96
L102:
	;
	goto L103
L103:
	;
	v463 = int32(0)
	if v459 != int32(1) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v467 = int32(0)
	if v467 < v459 {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	v519 = v463
	v521 = v463
	goto L106
L106:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v454)+12))
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v536+v521<<(uint(int32(2))%32))))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v540)))
	if v541 != v30 {
		v546 = v519
		goto L96
	} else {
		goto L126
	}
L107:
	;
	v470 = v459
	goto L109
L108:
	;
	v470 = v467
	goto L109
L109:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v454)+12))
	v476 = int32(0)
	v478 = v476
	v480 = v463
	v484 = v476
	goto L110
L110:
	;
	v497 = v475 + v480<<(uint(int32(2))%32)
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v497)))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v498)))
	if v30 == v499 {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	if v470&int32(1) == int32(0) {
		v546 = v511
		goto L96
	} else {
		goto L125
	}
L112:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v498)+4))
	if v501 == v30 {
		goto L115
	} else {
		goto L116
	}
L113:
	;
	v504 = v478
	goto L114
L114:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v497)+4))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v505)))
	if v30 == v506 {
		goto L118
	} else {
		goto L119
	}
L115:
	;
	v503 = v498
	goto L117
L116:
	;
	v503 = v478
	goto L117
L117:
	;
	v504 = v503
	goto L114
L118:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v505)+4))
	if v508 == v30 {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	v511 = v504
	goto L120
L120:
	;
	v512 = int32(2)
	v513 = v480 + v512
	v515 = v484 + v512
	if v515 != v470&int32(2147483646) {
		v478 = v511
		v480 = v513
		v484 = v515
		goto L110
	} else {
		goto L124
	}
L121:
	;
	v510 = v505
	goto L123
L122:
	;
	v510 = v504
	goto L123
L123:
	;
	v511 = v510
	goto L120
L124:
	;
	goto L111
L125:
	;
	v519 = v511
	v521 = v513
	goto L106
L126:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v540)+4))
	if v543 == v30 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v545 = v540
	goto L129
L128:
	;
	v545 = v519
	goto L129
L129:
	;
	v546 = v545
	goto L96
L130:
	;
	F_ReleaseCatCacheList(m, v46)
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L2
	} else {
		goto L156
	}
L131:
	;
	v640 = int32(0)
	v643 = F_errstart(m, int32(17), v640)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L2
	} else {
		goto L151
	}
L132:
	;
	if v546 != 0 {
		goto L136
	} else {
		goto L137
	}
L133:
	;
	v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v546)+16)))
	if v634&int32(80) != 0 {
		v666 = v632
		goto L130
	} else {
		goto L150
	}
L134:
	;
	goto L133
L135:
	;
	v628 = v582 + int64(1)
	if v628 != int64(7) {
		v582 = v628
		goto L132
	} else {
		goto L149
	}
L136:
	;
	v583 = *(*int64)(unsafe.Add(mBase, uint32(v546)+16))
	if base.I32_wrap_i64(int64(base.Ui64(v583)>>(uint(v582)%64)))&int32(1) != 0 {
		goto L135
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	v588 = base.I32_wrap_i64(v582)
	if base.B2i32(v588&int32(5) == int32(4))|base.B2i32(v588&int32(3) == int32(1)) != 0 {
		v622 = v571
		goto L140
	} else {
		goto L141
	}
L139:
	;
	goto L138
L140:
	;
	v624 = v582 + int64(1)
	if v624 != int64(7) {
		v571 = v622
		v582 = v624
		goto L132
	} else {
		goto L147
	}
L141:
	;
	v598 = int32(0)
	v601 = F_errstart(m, int32(17), v598)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L2
	} else {
		goto L142
	}
L142:
	;
	if v601 == int32(0) {
		v622 = v598
		goto L140
	} else {
		goto L143
	}
L143:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L2
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v588
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = int32(_a_F_ginvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v564
	F_errmsg(m, int32(_a_F_ginvalidate_12), v20+int32(32))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L2
	} else {
		goto L145
	}
L145:
	;
	F_errfinish(m, int32(_a_F_ginvalidate_2), int32(242), int32(_a_F_ginvalidate_3))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L2
	} else {
		goto L146
	}
L146:
	;
	v622 = v598
	goto L140
L147:
	;
	if v546 != 0 {
		v632 = v622
		goto L134
	} else {
		goto L148
	}
L148:
	;
	goto L131
L149:
	;
	v632 = v571
	goto L134
L150:
	;
	goto L131
L151:
	;
	if v643 == int32(0) {
		v666 = v640
		goto L130
	} else {
		goto L152
	}
L152:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L2
	} else {
		goto L153
	}
L153:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+24)) = int64(25769803780)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = int32(_a_F_ginvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v564
	F_errmsg(m, int32(_a_F_ginvalidate_13), v20+int32(16))
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L2
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(_a_F_ginvalidate_2), int32(253), int32(_a_F_ginvalidate_3))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L2
	} else {
		goto L155
	}
L155:
	;
	v666 = v640
	goto L130
L156:
	;
	F_ReleaseCatCacheList(m, v39)
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L2
	} else {
		goto L157
	}
L157:
	;
	F_ReleaseCatCache(m, v23)
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L2
	} else {
		goto L158
	}
L158:
	;
	m.G0 = v20 + int32(272)
	return v666 & int32(1)
}
func F_gistbulkdelete(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	if l1 == int32(0) {
		v8 = F_palloc0(m, int32(40))
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v12 = v8
			F_gistvacuumscan(m, l0, v12, l2, l3)
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				return v12
			}
		}
	} else {
		v12 = l1
		F_gistvacuumscan(m, l0, v12, l2, l3)
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			return v12
		}
	}
}
func F_gistcheckpage(m *base.Module, l0 int32, l1 int32) {
	var v8 int32
	_ = v8
	Fn13931(m, l0, l1, int32(_a_F_gistcheckpage_0), int32(812), int32(_a_F_gistcheckpage_1), int32(801))
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
	}
}
func F_gistfillitupvec(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	v4 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v4
	if v4 < l1 {
		if l1 != int32(1) {
			v23 = v4
			v24 = v4
			v28 = v4
			for {
				v29 = int32(2)
				v31 = l0 + v24<<(uint(v29)%32)
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
				v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+6)))
				v34 = int32(_a_F_gistfillitupvec_0)
				v36 = v23 + v33&v34
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v36
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
				v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+6)))
				v42 = v36 + v39&v34
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v42
				v45 = v24 + v29
				v47 = v28 + v29
				if v47 != l1&int32(2147483646) {
					v23 = v42
					v24 = v45
					v28 = v47
					continue
				} else {
					break
				}
				break
			}
			if l1&int32(1) == int32(0) {
				v72 = v42
			} else {
				v54 = v42
				v55 = v45
				v63 = *(*int32)(unsafe.Add(mBase, uint32(l0+v55<<(uint(int32(2))%32))))
				v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63)+6)))
				v67 = v54 + v64&int32(_a_F_gistfillitupvec_0)
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v67
				v72 = v67
			}
		} else {
			v54 = v4
			v55 = v4
			v63 = *(*int32)(unsafe.Add(mBase, uint32(l0+v55<<(uint(int32(2))%32))))
			v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63)+6)))
			v67 = v54 + v64&int32(_a_F_gistfillitupvec_0)
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v67
			v72 = v67
		}
		v79 = F_palloc(m, v72)
		mBase = m.M
		v82 = m.ExcPending
		if v82 != 0 {
			return int32(0)
		} else {
			v86 = v79
			v87 = int32(0)
			for {
				v94 = l0 + v87<<(uint(int32(2))%32)
				v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
				v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v95)+6)))
				v98 = v96 & int32(_a_F_gistfillitupvec_0)
				if v98 != 0 {
					base.MemoryCopy(m, v86, v95, v98)
				} else {
				}
				v100 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
				v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+6)))
				v106 = v87 + int32(1)
				if v106 != l1 {
					v86 = v86 + v101&int32(_a_F_gistfillitupvec_0)
					v87 = v106
					continue
				} else {
					break
				}
				break
			}
			return v79
		}
	} else {
		v110 = F_palloc(m, int32(0))
		mBase = m.M
		v111 = m.ExcPending
		if v111 != 0 {
			return int32(0)
		} else {
			return v110
		}
	}
}
func F_gistfinishsplit(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var __phi32 int32
	_ = __phi32
	var v33 int32
	_ = v33
	var __phi33 int32
	_ = __phi33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	F_LockBuffer(m, v16, int32(2))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if l3 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v81
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v83
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_gistFindCorrectParent(m, v85, l1)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L14
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v24 = v22 - int32(1)
	if v24 < int32(2) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	__phi32 = v22
	__phi33 = v24
	v32 = __phi32
	v33 = __phi33
	goto L6
L6:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v38 = int32(2)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v37+v32<<(uint(v38)%32)-int32(8))))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v37+v33<<(uint(v38)%32))))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_gistFindCorrectParent(m, v48, l1)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	goto L3
L8:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v55 = int32(0)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v60 = F_gistinserttuples(m, l0, v51, l2, v47+int32(4), int32(1), v55, v56, v57, v55, v55)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if v60 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v62 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)) = uint16(v62)
	goto L12
L11:
	;
	goto L12
L12:
	;
	if int32(2) < v33 {
		__phi32 = v33
		__phi33 = v33 - int32(1)
		v32 = __phi32
		v33 = __phi33
		goto L6
	} else {
		goto L13
	}
L13:
	;
	goto L7
L14:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v96 = F_gistinserttuples(m, l0, v88, l2, v13+int32(8), int32(2), v92, v93, v94, int32(1), l4)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v98 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v98)
	v100 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)) = uint16(v100)
	m.G0 = v13 + int32(16)
	return
}
func F_gistnospace(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	v6 = int32(0)
	if l2 <= v6 {
		v73 = l4
	} else {
		if l2 != int32(1) {
			v21 = int32(0)
			v23 = l4
			v24 = v6
			for {
				v28 = int32(2)
				v30 = l1 + v24<<(uint(v28)%32)
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
				v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31)+6)))
				v33 = int32(_a_F_gistnospace_0)
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
				v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+6)))
				v42 = v23 + v32&v33 + v37&v33 + int32(8)
				v44 = v24 + v28
				v46 = v21 + v28
				if v46 != l2&int32(2147483646) {
					v21 = v46
					v23 = v42
					v24 = v44
					continue
				} else {
					break
				}
				break
			}
			if l2&int32(1) == int32(0) {
				v73 = v42
			} else {
				v54 = v42
				v55 = v44
				v62 = *(*int32)(unsafe.Add(mBase, uint32(l1+v55<<(uint(int32(2))%32))))
				v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62)+6)))
				v73 = v63&int32(_a_F_gistnospace_0) + v54 + int32(4)
			}
		} else {
			v54 = l4
			v55 = v6
			v62 = *(*int32)(unsafe.Add(mBase, uint32(l1+v55<<(uint(int32(2))%32))))
			v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62)+6)))
			v73 = v63&int32(_a_F_gistnospace_0) + v54 + int32(4)
		}
	}
	if l3 != 0 {
		v81 = *(*int32)(unsafe.Add(mBase, uint32(l0+l3<<(uint(int32(2))%32))+20))
		v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+v81&int32(_a_F_gistnospace_1))+6)))
		v91 = v85&int32(_a_F_gistnospace_0) + int32(4)
	} else {
		v91 = int32(0)
	}
	v92 = int32(4)
	v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
	v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	v95 = v93 - v94
	if v95 <= v92 {
		v98 = v92
	} else {
		v98 = v95
	}
	return base.B2i32(base.Ui32(v98-int32(4)+v91) < base.Ui32(v73))
}
func F_gistplacetopage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v332 int32
	_ = v332
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v441 int64
	_ = v441
	var v442 int32
	_ = v442
	var v446 int64
	_ = v446
	var v447 int32
	_ = v447
	var v451 int64
	_ = v451
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v487 int32
	_ = v487
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v554 int32
	_ = v554
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v614 int32
	_ = v614
	var v621 int32
	_ = v621
	var v626 int32
	_ = v626
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v682 int32
	_ = v682
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int64
	_ = v708
	var v709 int64
	_ = v709
	var v714 int32
	_ = v714
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v737 int32
	_ = v737
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v761 int64
	_ = v761
	var v762 int32
	_ = v762
	var v778 int32
	_ = v778
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v797 int32
	_ = v797
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v811 int32
	_ = v811
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v822 int32
	_ = v822
	var v826 int32
	_ = v826
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v840 int32
	_ = v840
	var v845 int32
	_ = v845
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v865 int32
	_ = v865
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v885 int32
	_ = v885
	var v891 int32
	_ = v891
	var v893 int32
	_ = v893
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v930 int32
	_ = v930
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v937 int32
	_ = v937
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v962 int32
	_ = v962
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v984 int32
	_ = v984
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v998 int32
	_ = v998
	var v999 int64
	_ = v999
	var v1000 int32
	_ = v1000
	var v1001 int64
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1003 int64
	_ = v1003
	var v1007 int32
	_ = v1007
	var v1014 int32
	_ = v1014
	var v1018 int32
	_ = v1018
	var v1023 int32
	_ = v1023
	var v1027 int32
	_ = v1027
	var v1033 int32
	_ = v1033
	var v1038 int32
	_ = v1038
	var v1078 int32
	_ = v1078
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1099 int32
	_ = v1099
	var v1115 int32
	_ = v1115
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1133 int32
	_ = v1133
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1139 int32
	_ = v1139
	var v1174 int32
	_ = v1174
	var v1180 int32
	_ = v1180
	var v1182 int32
	_ = v1182
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1198 int32
	_ = v1198
	var v1210 int32
	_ = v1210
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1260 int32
	_ = v1260
	var v1264 int32
	_ = v1264
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1269 int32
	_ = v1269
	var v1271 int32
	_ = v1271
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1313 int32
	_ = v1313
	var v1338 int32
	_ = v1338
	var v1363 int32
	_ = v1363
	var v1366 int32
	_ = v1366
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1399 int32
	_ = v1399
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1411 int32
	_ = v1411
	var v1412 int32
	_ = v1412
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1422 int32
	_ = v1422
	var v1423 int32
	_ = v1423
	var v1426 int32
	_ = v1426
	var v1429 int32
	_ = v1429
	var v1431 int32
	_ = v1431
	var v1435 int32
	_ = v1435
	var v1463 int32
	_ = v1463
	var v1466 int32
	_ = v1466
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1472 int32
	_ = v1472
	var v1475 int32
	_ = v1475
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1484 int32
	_ = v1484
	var v1486 int32
	_ = v1486
	var v1487 int32
	_ = v1487
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1496 int32
	_ = v1496
	var v1497 int32
	_ = v1497
	var v1500 int32
	_ = v1500
	var v1507 int32
	_ = v1507
	var v1515 int32
	_ = v1515
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1532 int32
	_ = v1532
	var v1535 int32
	_ = v1535
	var v1536 int32
	_ = v1536
	var v1537 int32
	_ = v1537
	var v1542 int32
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1545 int32
	_ = v1545
	var v1564 int32
	_ = v1564
	var v1577 int32
	_ = v1577
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1614 int32
	_ = v1614
	var v1620 int32
	_ = v1620
	var v1622 int32
	_ = v1622
	var v1628 int32
	_ = v1628
	var v1630 int32
	_ = v1630
	var v1631 int32
	_ = v1631
	var v1635 int32
	_ = v1635
	var v1641 int32
	_ = v1641
	var v1643 int32
	_ = v1643
	var v1649 int32
	_ = v1649
	var v1652 int32
	_ = v1652
	var v1653 int32
	_ = v1653
	var v1657 int32
	_ = v1657
	var v1660 int32
	_ = v1660
	var v1661 int32
	_ = v1661
	var v1662 int32
	_ = v1662
	var v1663 int32
	_ = v1663
	var v1665 int32
	_ = v1665
	var v1668 int32
	_ = v1668
	var v1669 int32
	_ = v1669
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1698 int32
	_ = v1698
	var v1728 int32
	_ = v1728
	var v1732 int32
	_ = v1732
	var v1737 int32
	_ = v1737
	var v1739 int32
	_ = v1739
	var v1741 int32
	_ = v1741
	var v1766 int32
	_ = v1766
	var v1767 int32
	_ = v1767
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
	var v1779 int32
	_ = v1779
	var v1782 int32
	_ = v1782
	var v1811 int64
	_ = v1811
	var v1812 int32
	_ = v1812
	var v1816 int64
	_ = v1816
	var v1817 int32
	_ = v1817
	var v1844 int64
	_ = v1844
	var v1864 int32
	_ = v1864
	var v1877 int32
	_ = v1877
	var v1880 int32
	_ = v1880
	var v1907 int32
	_ = v1907
	var v1923 int32
	_ = v1923
	var v1936 int32
	_ = v1936
	var v1938 int32
	_ = v1938
	var v1939 int32
	_ = v1939
	var v1981 int32
	_ = v1981
	var v1991 int64
	_ = v1991
	var v1996 int32
	_ = v1996
	var v2002 int32
	_ = v2002
	var v2004 int32
	_ = v2004
	var v2010 int32
	_ = v2010
	var v2011 int32
	_ = v2011
	var v2015 int32
	_ = v2015
	var v2017 int32
	_ = v2017
	var v2019 int32
	_ = v2019
	var v2021 int32
	_ = v2021
	var v2022 int32
	_ = v2022
	var v2023 int32
	_ = v2023
	var v2025 int32
	_ = v2025
	var v2033 int32
	_ = v2033
	var v2035 int32
	_ = v2035
	var v2046 int32
	_ = v2046
	var v2047 int32
	_ = v2047
	var v2055 int32
	_ = v2055
	var v2060 int32
	_ = v2060
	v7 = l6
	v11 = l10
	v14 = int32(0)
	v27 = m.G0
	v29 = v27 - int32(864)
	m.G0 = v29
	if l3 < v14 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if l3 < int32(0) {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[0]))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v34+(l3^int32(-1))<<(uint(int32(6))%32))+16))
	v49 = v40
	goto L1
L3:
	;
	goto L4
L4:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[1]))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v42+l3<<(uint(int32(6))%32)+int32(-64))+16))
	v49 = v48
	goto L1
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2046 = m.ExcPending
	if v2046 != 0 {
		goto L59
	} else {
		goto L339
	}
L6:
	;
	if l8 != 0 {
		goto L332
	} else {
		goto L333
	}
L7:
	;
	if v638 != 0 {
		goto L195
	} else {
		goto L196
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L59
	} else {
		goto L189
	}
L9:
	;
	v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+16)))
	v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68+v67)+12)))
	if v70&int32(8) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[2]))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v53+(l3^int32(-1))<<(uint(int32(2))%32))))
	v67 = v59
	goto L9
L11:
	;
	goto L12
L12:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[3]))
	v67 = v61 + l3<<(uint(int32(13))%32) + int32(-8192)
	goto L9
L13:
	;
	v75 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l9))) = v75
	if l5 <= v75 {
		v144 = l1
		goto L18
	} else {
		goto L19
	}
L14:
	;
	goto L15
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1014 = m.ExcPending
	if v1014 != 0 {
		goto L59
	} else {
		goto L186
	}
L16:
	;
	v930 = int32(_a_F_gistplacetopage_0)
	v932 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[4]))
	v933 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[4])) = v932 + v933
	v937 = v7 - v933
	if base.Ui32(v937&int32(_a_F_gistplacetopage_1)) <= base.Ui32(int32(2047)) {
		goto L153
	} else {
		goto L154
	}
L17:
	;
	if base.B2i32(base.Ui32(v163+v162) < base.Ui32(v144)) == int32(0) {
		goto L16
	} else {
		goto L30
	}
L18:
	;
	if v7 != 0 {
		goto L27
	} else {
		goto L28
	}
L19:
	;
	if l5 != int32(1) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v92 = int32(0)
	v94 = l1
	v95 = v75
	goto L23
L21:
	;
	v125 = l1
	v126 = v75
	goto L22
L22:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l4+v126<<(uint(int32(2))%32))))
	v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v133)+6)))
	v144 = v134&int32(_a_F_gistplacetopage_2) + v125 + int32(4)
	goto L18
L23:
	;
	v99 = int32(2)
	v101 = l4 + v95<<(uint(v99)%32)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102)+6)))
	v104 = int32(_a_F_gistplacetopage_2)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v107)+6)))
	v113 = v94 + v103&v104 + v108&v104 + int32(8)
	v115 = v95 + v99
	v117 = v92 + v99
	if v117 != l5&int32(2147483646) {
		v92 = v117
		v94 = v113
		v95 = v115
		goto L23
	} else {
		goto L25
	}
L24:
	;
	if l5&int32(1) == int32(0) {
		v144 = v113
		goto L18
	} else {
		goto L26
	}
L25:
	;
	goto L24
L26:
	;
	v125 = v113
	v126 = v115
	goto L22
L27:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v67+v7<<(uint(int32(2))%32))+20))
	v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67+v152&int32(_a_F_gistplacetopage_3))+6)))
	v162 = v156&int32(_a_F_gistplacetopage_2) + int32(4)
	goto L29
L28:
	;
	v162 = int32(0)
	goto L29
L29:
	;
	v163 = F_PageGetFreeSpace(m, v67)
	mBase = m.M
	goto L17
L30:
	;
	v168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+16)))
	v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67+v168)+12)))
	v171 = int32(17)
	if v170&v171 == v171 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v175 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+12)))
	if base.Ui32(v175) < base.Ui32(int32(25)) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	v606 = F_gistextractpage(m, v67, v29+int32(44))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L59
	} else {
		goto L98
	}
L34:
	;
	v487 = int32(0)
	if l5 <= v487 {
		v554 = l1
		goto L85
	} else {
		goto L86
	}
L35:
	;
	v181 = int32(base.Ui32(v175+int32(_a_F_gistplacetopage_4)) >> (uint(int32(2)) % 32))
	if v181&int32(_a_F_gistplacetopage_1) == int32(0) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v186 = int32(1)
	v188 = v67 + int32(20)
	v192 = (v181 + v186) & int32(_a_F_gistplacetopage_1)
	if base.Ui32(int32(3)) <= base.Ui32(v192) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	if v332 <= int32(0) {
		goto L34
	} else {
		goto L55
	}
L38:
	;
	v195 = int32(2)
	if base.Ui32(v192) <= base.Ui32(v195) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	v289 = v186
	v290 = v14
	goto L40
L40:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v188+v289<<(uint(int32(2))%32))))
	v306 = int32(_a_F_gistplacetopage_5)
	if v305&v306 != v306 {
		v332 = v290
		goto L37
	} else {
		goto L54
	}
L41:
	;
	v198 = v195
	goto L43
L42:
	;
	v198 = v192
	goto L43
L43:
	;
	v199 = int32(1)
	v200 = v198 - v199
	v220 = v199
	v221 = v14
	v224 = int32(0)
	goto L44
L44:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v188+v220<<(uint(int32(2))%32))))
	v237 = int32(_a_F_gistplacetopage_5)
	if v236&v237 == v237 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	if v200&v199 == int32(0) {
		v332 = v268
		goto L37
	} else {
		goto L53
	}
L46:
	;
	v243 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v29+int32(48)+v221<<(uint(v243)%32)))) = uint16(v220)
	v249 = v221 + v243
	goto L48
L47:
	;
	v249 = v221
	goto L48
L48:
	;
	v251 = v220 + int32(1)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v188+v251<<(uint(int32(2))%32))))
	v256 = int32(_a_F_gistplacetopage_5)
	if v255&v256 == v256 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v262 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v29+int32(48)+v249<<(uint(v262)%32)))) = uint16(v251)
	v268 = v249 + v262
	goto L51
L50:
	;
	v268 = v249
	goto L51
L51:
	;
	v269 = int32(2)
	v270 = v220 + v269
	v272 = v224 + v269
	if v272 != v200&int32(-2) {
		v220 = v270
		v221 = v268
		v224 = v272
		goto L44
	} else {
		goto L52
	}
L52:
	;
	goto L45
L53:
	;
	v289 = v270
	v290 = v268
	goto L40
L54:
	;
	v312 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v29+int32(48)+v290<<(uint(v312)%32)))) = uint16(v289)
	v332 = v290 + v312
	goto L37
L55:
	;
	v346 = int32(0)
	v348 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[5]))
	if v348 <= v346 {
		v362 = v346
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v363 = int32(_a_F_gistplacetopage_0)
	v365 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[4])) = v365 + int32(1)
	F_PageIndexMultiDelete(m, v67, v29+int32(48), v332)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L59
	} else {
		goto L61
	}
L57:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352)+118)))
	if v353 != int32(112) {
		v362 = int32(0)
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v358 = F_index_compute_xid_horizon_for_tuples(m, l0, l11, l3, v29+int32(48), v332)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	return int32(0)
L60:
	;
	v362 = v358
	goto L56
L61:
	;
	v373 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+16)))
	v374 = v67 + v373
	v375 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v374)+12)))
	v377 = v375 & int32(_a_F_gistplacetopage_6)
	*(*uint16)(unsafe.Add(mBase, uint32(v374)+12)) = uint16(v377)
	F_MarkBufferDirty(m, l3)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L59
	} else {
		goto L62
	}
L62:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381)+118)))
	if v382 != int32(112) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v67))) = base.I64_rotr(v451, int64(32))
	v455 = int32(_a_F_gistplacetopage_0)
	v457 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[4])) = v457 - int32(1)
	goto L34
L64:
	;
	v446 = F_gistGetFakeLSN(m, l0)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L59
	} else {
		goto L83
	}
L65:
	;
	v386 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[5]))
	if v386 <= int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v389 != 0 {
		goto L64
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v393 = int32(0)
	v394 = m.G0
	v396 = v394 - int32(16)
	m.G0 = v396
	v399 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[5]))
	if v399 < int32(2) {
		v419 = v393
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v390 != 0 {
		goto L64
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v396)+8)) = v362
	*(*uint8)(unsafe.Add(mBase, uint32(v396)+14)) = uint8(v419)
	*(*uint16)(unsafe.Add(mBase, uint32(v396)+12)) = uint16(v332)
	F_XLogBeginInsert(m)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L59
	} else {
		goto L78
	}
L72:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l11)+48))
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v402)+118)))
	if v403 != int32(112) {
		v419 = v393
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(l11)+56))
	goto L74
L74:
	;
	if base.Ui32(v407) < base.Ui32(int32(_a_F_gistplacetopage_7)) {
		v419 = int32(1)
		goto L71
	} else {
		goto L75
	}
L75:
	;
	v410 = int32(0)
	v411 = *(*int32)(unsafe.Add(mBase, uint32(l11)+180))
	if v411 == v410 {
		v419 = v410
		goto L71
	} else {
		goto L76
	}
L76:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(l11)+48))
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v414)+119)))
	switch v415 - int32(109) {
	case 0, 5:
		goto L77
	default:
		v419 = v410
		goto L71
	}
L77:
	;
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411)+104)))
	v419 = v418
	goto L71
L78:
	;
	v426 = int32(8)
	F_XLogRegisterData(m, v396+v426, v426)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L59
	} else {
		goto L79
	}
L79:
	;
	F_XLogRegisterData(m, v29+int32(48), v332<<(uint(int32(1))%32))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L59
	} else {
		goto L80
	}
L80:
	;
	F_XLogRegisterBuffer(m, int32(0), l3, int32(8))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L59
	} else {
		goto L81
	}
L81:
	;
	v441 = F_XLogInsert(m, int32(14), int32(16))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L59
	} else {
		goto L82
	}
L82:
	;
	m.G0 = v396 + int32(16)
	v451 = v441
	goto L63
L83:
	;
	v451 = v446
	goto L63
L84:
	;
	if base.B2i32(base.Ui32(v573+v572) < base.Ui32(v554)) == int32(0) {
		goto L16
	} else {
		goto L97
	}
L85:
	;
	if v7 != 0 {
		goto L94
	} else {
		goto L95
	}
L86:
	;
	if l5 != int32(1) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v502 = int32(0)
	v504 = l1
	v505 = v487
	goto L90
L88:
	;
	v535 = l1
	v536 = v487
	goto L89
L89:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(l4+v536<<(uint(int32(2))%32))))
	v544 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v543)+6)))
	v554 = v544&int32(_a_F_gistplacetopage_2) + v535 + int32(4)
	goto L85
L90:
	;
	v509 = int32(2)
	v511 = l4 + v505<<(uint(v509)%32)
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v511)))
	v513 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v512)+6)))
	v514 = int32(_a_F_gistplacetopage_2)
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v511)+4))
	v518 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v517)+6)))
	v523 = v504 + v513&v514 + v518&v514 + int32(8)
	v525 = v505 + v509
	v527 = v502 + v509
	if v527 != l5&int32(2147483646) {
		v502 = v527
		v504 = v523
		v505 = v525
		goto L90
	} else {
		goto L92
	}
L91:
	;
	if l5&int32(1) == int32(0) {
		v554 = v523
		goto L85
	} else {
		goto L93
	}
L92:
	;
	goto L91
L93:
	;
	v535 = v523
	v536 = v525
	goto L89
L94:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v67+v7<<(uint(int32(2))%32))+20))
	v566 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67+v562&int32(_a_F_gistplacetopage_3))+6)))
	v572 = v566&int32(_a_F_gistplacetopage_2) + int32(4)
	goto L96
L95:
	;
	v572 = int32(0)
	goto L96
L96:
	;
	v573 = F_PageGetFreeSpace(m, v67)
	mBase = m.M
	goto L84
L97:
	;
	goto L33
L98:
	;
	if base.Ui32(int32(2047)) < base.Ui32((v7-int32(1))&int32(_a_F_gistplacetopage_1)) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v632 = int32(0)
	v635 = F_gistjoinvector(m, v606, v29+int32(44), l4, l5)
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L59
	} else {
		goto L103
	}
L100:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v29)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v614 - int32(1)
	if v7 == v614 {
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v621 = (v614 - v7) << (uint(int32(2)) % 32)
	if v621 == int32(0) {
		goto L99
	} else {
		goto L102
	}
L102:
	;
	v626 = v606 + v7<<(uint(int32(2))%32)
	base.MemoryCopy(m, v626-int32(4), v626, v621)
	goto L99
L103:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v29)+44))
	v638 = F_gistSplit(m, l0, v67, v635, v637, l2)
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L59
	} else {
		goto L104
	}
L104:
	;
	if v638 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v653 = v632
	v654 = v638
	goto L108
L106:
	;
	v682 = v632
	goto L107
L107:
	;
	v697 = v682 + base.B2i32(v49 == int32(0))
	if int32(76) <= v697 {
		goto L8
	} else {
		goto L111
	}
L108:
	;
	v667 = v653 + int32(1)
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v654)+28))
	if v668 != 0 {
		v653 = v667
		v654 = v668
		goto L108
	} else {
		goto L110
	}
L109:
	;
	v682 = v667
	goto L107
L110:
	;
	goto L109
L111:
	;
	v701 = v70 & int32(1)
	if v49 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	if v762 == int32(0) {
		goto L7
	} else {
		goto L125
	}
L113:
	;
	v760 = int32(-1)
	v761 = int64(0)
	v762 = v638
	goto L112
L114:
	;
	goto L115
L115:
	;
	v705 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+16)))
	v706 = v67 + v705
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v706)+8))
	v708 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v706)+4)))
	v709 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v706))))
	*(*int32)(unsafe.Add(mBase, uint32(v638)+24)) = l3
	if l3 < int32(0) {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v638))) = v729
	if l3 < int32(0) {
		goto L121
	} else {
		goto L122
	}
L117:
	;
	v714 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[0]))
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v714+(l3^int32(-1))<<(uint(int32(6))%32))+16))
	v729 = v720
	goto L116
L118:
	;
	goto L119
L119:
	;
	v722 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[1]))
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v722+l3<<(uint(int32(6))%32)+int32(-64))+16))
	v729 = v728
	goto L116
L120:
	;
	v752 = F_PageGetTempPageCopySpecial(m, v751)
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L59
	} else {
		goto L124
	}
L121:
	;
	v737 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[2]))
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v737+(l3^int32(-1))<<(uint(int32(2))%32))))
	v751 = v743
	goto L120
L122:
	;
	goto L123
L123:
	;
	v745 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[3]))
	v751 = v745 + l3<<(uint(int32(13))%32) + int32(-8192)
	goto L120
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v638)+20)) = v752
	v755 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v752)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v752+v755)+12)) = uint16(v701)
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v638)+28))
	v760 = v707
	v761 = v709<<(uint(int64(32))%64) | v708
	v762 = v758
	goto L112
L125:
	;
	v778 = v762
	goto L126
L126:
	;
	v791 = F_gistNewBuffer(m, l0, l11)
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L59
	} else {
		goto L128
	}
L127:
	;
	goto L7
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v778)+24)) = v791
	if v791 < int32(0) {
		goto L131
	} else {
		goto L132
	}
L129:
	;
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v778)+24))
	if v822 < int32(0) {
		goto L135
	} else {
		goto L136
	}
L130:
	;
	F_PageInit(m, v811, int32(_a_F_gistplacetopage_8), int32(16))
	mBase = m.M
	v815 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v811)+16)))
	v816 = v811 + v815
	v817 = int32(_a_F_gistplacetopage_9)
	*(*uint16)(unsafe.Add(mBase, uint32(v816)+14)) = uint16(v817)
	*(*uint16)(unsafe.Add(mBase, uint32(v816)+12)) = uint16(v701)
	*(*int32)(unsafe.Add(mBase, uint32(v816)+8)) = int32(-1)
	goto L129
L131:
	;
	v797 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[2]))
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v797+(v791^int32(-1))<<(uint(int32(2))%32))))
	v811 = v803
	goto L130
L132:
	;
	goto L133
L133:
	;
	v805 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[3]))
	v811 = v805 + v791<<(uint(int32(13))%32) + int32(-8192)
	goto L130
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v778)+20)) = v840
	if v822 < int32(0) {
		goto L139
	} else {
		goto L140
	}
L135:
	;
	v826 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[2]))
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v826+(v822^int32(-1))<<(uint(int32(2))%32))))
	v840 = v832
	goto L134
L136:
	;
	goto L137
L137:
	;
	v834 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[3]))
	v840 = v834 + v822<<(uint(int32(13))%32) + int32(-8192)
	goto L134
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v778))) = v860
	if l3 < int32(0) {
		goto L143
	} else {
		goto L144
	}
L139:
	;
	v845 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[0]))
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v845+(v822^int32(-1))<<(uint(int32(6))%32))+16))
	v860 = v851
	goto L138
L140:
	;
	goto L141
L141:
	;
	v853 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[1]))
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v853+v822<<(uint(int32(6))%32)+int32(-64))+16))
	v860 = v859
	goto L138
L142:
	;
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v778)+24))
	if v881 < int32(0) {
		goto L147
	} else {
		goto L148
	}
L143:
	;
	v865 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[0]))
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v865+(l3^int32(-1))<<(uint(int32(6))%32))+16))
	v880 = v871
	goto L142
L144:
	;
	goto L145
L145:
	;
	v873 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[1]))
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v873+l3<<(uint(int32(6))%32)+int32(-64))+16))
	v880 = v879
	goto L142
L146:
	;
	F_PredicateLockPageSplit(m, l0, v880, v900)
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		goto L59
	} else {
		goto L150
	}
L147:
	;
	v885 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[0]))
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v885+(v881^int32(-1))<<(uint(int32(6))%32))+16))
	v900 = v891
	goto L146
L148:
	;
	goto L149
L149:
	;
	v893 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[1]))
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v893+v881<<(uint(int32(6))%32)+int32(-64))+16))
	v900 = v899
	goto L146
L150:
	;
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v778)+28))
	if v903 != 0 {
		v778 = v903
		goto L126
	} else {
		goto L151
	}
L151:
	;
	goto L127
L152:
	;
	F_MarkBufferDirty(m, l3)
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L59
	} else {
		goto L166
	}
L153:
	;
	if l5 == int32(1) {
		goto L156
	} else {
		goto L157
	}
L154:
	;
	goto L155
L155:
	;
	F_gistfillbuffer(m, v67, l4, l5, int32(0))
	mBase = m.M
	v972 = m.ExcPending
	if v972 != 0 {
		goto L59
	} else {
		goto L165
	}
L156:
	;
	v944 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v945 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v944)+6)))
	v948 = F_PageIndexTupleOverwrite(m, v67, v7, v944, v945&int32(_a_F_gistplacetopage_2))
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L59
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	F_PageIndexTupleDelete(m, v67, v7)
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L59
	} else {
		goto L164
	}
L159:
	;
	if v948 != 0 {
		goto L152
	} else {
		goto L160
	}
L160:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v953 = m.ExcPending
	if v953 != 0 {
		goto L59
	} else {
		goto L161
	}
L161:
	;
	v954 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v954 + int32(4)
	F_errmsg_internal(m, int32(_a_F_gistplacetopage_10), v29+int32(32))
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L59
	} else {
		goto L162
	}
L162:
	;
	F_errfinish(m, int32(_a_F_gistplacetopage_11), int32(557), int32(_a_F_gistplacetopage_12))
	mBase = m.M
	v967 = m.ExcPending
	if v967 != 0 {
		goto L59
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
	goto L155
L165:
	;
	goto L152
L166:
	;
	if l8 != 0 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	F_MarkBufferDirty(m, l8)
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L59
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	if l12 != 0 {
		v1003 = int64(1)
		goto L171
	} else {
		goto L172
	}
L170:
	;
	goto L169
L171:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v67))) = base.I64_rotr(v1003, int64(32))
	v1007 = int32(0)
	if l7 == v1007 {
		v1981 = v1007
		v1991 = v1003
		goto L6
	} else {
		goto L185
	}
L172:
	;
	v979 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v980 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v979)+118)))
	if v980 != int32(112) {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v1001 = F_gistGetFakeLSN(m, l0)
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L59
	} else {
		goto L184
	}
L174:
	;
	v984 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[5]))
	if v984 <= int32(0) {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v987 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v987 != 0 {
		goto L173
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	if base.Ui32(v937&int32(_a_F_gistplacetopage_1)) <= base.Ui32(int32(2047)) {
		goto L180
	} else {
		goto L181
	}
L178:
	;
	v988 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v988 != 0 {
		goto L173
	} else {
		goto L179
	}
L179:
	;
	goto L177
L180:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+48)) = uint16(v7)
	v998 = int32(1)
	goto L182
L181:
	;
	v998 = int32(0)
	goto L182
L182:
	;
	v999 = F_gistXLogUpdate(m, l3, v29+int32(48), v998, l4, l5, l8)
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L59
	} else {
		goto L183
	}
L183:
	;
	v1003 = v999
	goto L171
L184:
	;
	v1003 = v1001
	goto L171
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v49
	v1981 = v1007
	v1991 = v1003
	goto L6
L186:
	;
	F_errmsg_internal(m, int32(_a_F_gistplacetopage_13), int32(0))
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L59
	} else {
		goto L187
	}
L187:
	;
	F_errfinish(m, int32(_a_F_gistplacetopage_11), int32(256), int32(_a_F_gistplacetopage_12))
	mBase = m.M
	v1023 = m.ExcPending
	if v1023 != 0 {
		goto L59
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
	*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = int32(75)
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v697
	F_errmsg_internal(m, int32(_a_F_gistplacetopage_14), v29)
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L59
	} else {
		goto L190
	}
L190:
	;
	F_errfinish(m, int32(_a_F_gistplacetopage_11), int32(328), int32(_a_F_gistplacetopage_12))
	mBase = m.M
	v1038 = m.ExcPending
	if v1038 != 0 {
		goto L59
	} else {
		goto L191
	}
L191:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L192:
	;
	if l12 != 0 {
		goto L258
	} else {
		goto L259
	}
L193:
	;
	v1338 = v1313
	goto L226
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+72)) = l3
	if l3 < int32(0) {
		goto L209
	} else {
		goto L210
	}
L195:
	;
	v1078 = v638
	goto L198
L196:
	;
	goto L197
L197:
	;
	if v49 == int32(0) {
		goto L194
	} else {
		goto L207
	}
L198:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v1078)+16))
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v1078)))
	*(*int32)(unsafe.Add(mBase, uint32(v1091))) = base.I32_rotr(v1092, int32(16))
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v1078)+16))
	v1097 = int32(_a_F_gistplacetopage_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1096)+4)) = uint16(v1097)
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v1078)+28))
	if v1099 != 0 {
		v1078 = v1099
		goto L198
	} else {
		goto L200
	}
L199:
	;
	if v49 == int32(0) {
		goto L194
	} else {
		goto L201
	}
L200:
	;
	goto L199
L201:
	;
	v1115 = v638
	goto L202
L202:
	;
	v1129 = F_palloc(m, int32(8))
	mBase = m.M
	v1130 = m.ExcPending
	if v1130 != 0 {
		goto L59
	} else {
		goto L204
	}
L203:
	;
	v1313 = v638
	goto L193
L204:
	;
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(v1115)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1129))) = v1131
	v1133 = *(*int32)(unsafe.Add(mBase, uint32(v1115)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1129)+4)) = v1133
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(l9)))
	v1136 = F_lappend(m, v1135, v1129)
	mBase = m.M
	v1137 = m.ExcPending
	if v1137 != 0 {
		goto L59
	} else {
		goto L205
	}
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l9))) = v1136
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(v1115)+28))
	if v1139 != 0 {
		v1115 = v1139
		goto L202
	} else {
		goto L206
	}
L206:
	;
	goto L203
L207:
	;
	v1507 = int32(0)
	v1515 = int32(1)
	goto L192
L208:
	;
	v1189 = F_PageGetTempPageCopySpecial(m, v1188)
	mBase = m.M
	v1190 = m.ExcPending
	if v1190 != 0 {
		goto L59
	} else {
		goto L212
	}
L209:
	;
	v1174 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[2]))
	v1180 = *(*int32)(unsafe.Add(mBase, uint32(v1174+(l3^int32(-1))<<(uint(int32(2))%32))))
	v1188 = v1180
	goto L208
L210:
	;
	goto L211
L211:
	;
	v1182 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[3]))
	v1188 = v1182 + l3<<(uint(int32(13))%32) + int32(-8192)
	goto L208
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+68)) = v1189
	v1192 = int32(0)
	v1193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1189)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1189+v1193)+12)) = uint16(v1192)
	if v638 != 0 {
		goto L214
	} else {
		goto L215
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v1269
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = int32(0)
	v1299 = F_gistfillitupvec(m, v1271, v1269, v29+int32(60))
	mBase = m.M
	v1300 = m.ExcPending
	if v1300 != 0 {
		goto L59
	} else {
		goto L225
	}
L214:
	;
	v1198 = v1192
	v1210 = v638
	goto L217
L215:
	;
	goto L216
L216:
	;
	v1266 = F_palloc(m, int32(0))
	mBase = m.M
	v1267 = m.ExcPending
	if v1267 != 0 {
		goto L59
	} else {
		goto L224
	}
L217:
	;
	v1224 = v1198 + int32(1)
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(v1210)+28))
	if v1225 != 0 {
		v1198 = v1224
		v1210 = v1225
		goto L217
	} else {
		goto L219
	}
L218:
	;
	v1229 = F_palloc(m, v1224<<(uint(int32(2))%32))
	mBase = m.M
	v1230 = m.ExcPending
	if v1230 != 0 {
		goto L59
	} else {
		goto L220
	}
L219:
	;
	goto L218
L220:
	;
	v1244 = int32(0)
	v1245 = v638
	goto L221
L221:
	;
	v1260 = *(*int32)(unsafe.Add(mBase, uint32(v1245)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1229+v1244<<(uint(int32(2))%32)))) = v1260
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(v1245)+28))
	if v1264 != 0 {
		v1244 = v1244 + int32(1)
		v1245 = v1264
		goto L221
	} else {
		goto L223
	}
L222:
	;
	v1269 = v1224
	v1271 = v1229
	goto L213
L223:
	;
	goto L222
L224:
	;
	v1269 = v1192
	v1271 = v1266
	goto L213
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+76)) = v638
	*(*int32)(unsafe.Add(mBase, uint32(v29)+64)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v1299
	v1313 = v29 + int32(48)
	goto L193
L226:
	;
	v1363 = *(*int32)(unsafe.Add(mBase, uint32(v1338)+4))
	if int32(0) < v1363 {
		goto L228
	} else {
		goto L229
	}
L227:
	;
	v1507 = v1313
	v1515 = v1475
	goto L192
L228:
	;
	v1366 = *(*int32)(unsafe.Add(mBase, uint32(v1338)+8))
	v1381 = v1366
	v1382 = int32(0)
	goto L231
L229:
	;
	goto L230
L230:
	;
	v1463 = *(*int32)(unsafe.Add(mBase, uint32(v1338)+28))
	if v1463 == int32(0) {
		v1470 = v760
		goto L245
	} else {
		goto L246
	}
L231:
	;
	v1394 = *(*int32)(unsafe.Add(mBase, uint32(v1338)+20))
	v1395 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1381)+6)))
	v1399 = v1382 + int32(1)
	v1403 = F_PageAddItemExtended(m, v1394, v1381, v1395&int32(_a_F_gistplacetopage_2), v1399&int32(_a_F_gistplacetopage_1), int32(0))
	mBase = m.M
	v1404 = m.ExcPending
	if v1404 != 0 {
		goto L59
	} else {
		goto L233
	}
L232:
	;
	goto L230
L233:
	;
	if v1403 == int32(0) {
		goto L5
	} else {
		goto L234
	}
L234:
	;
	if l7 == int32(0) {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v1431 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1381)+6)))
	v1435 = *(*int32)(unsafe.Add(mBase, uint32(v1338)+4))
	if v1399 < v1435 {
		v1381 = v1381 + v1431&int32(_a_F_gistplacetopage_2)
		v1382 = v1399
		goto L231
	} else {
		goto L244
	}
L236:
	;
	v1409 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v1410 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1381)+2)))
	v1411 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1381))))
	v1412 = int32(16)
	v1415 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1409)+2)))
	v1416 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1409))))
	if v1410|v1411<<(uint(v1412)%32) == v1415|v1416<<(uint(v1412)%32) {
		goto L239
	} else {
		goto L240
	}
L237:
	;
	if v1426 == int32(0) {
		goto L235
	} else {
		goto L243
	}
L238:
	;
	goto L237
L239:
	;
	v1422 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1381)+4)))
	v1423 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1409)+4)))
	if v1422 == v1423 {
		v1426 = int32(1)
		goto L238
	} else {
		goto L242
	}
L240:
	;
	goto L241
L241:
	;
	v1426 = int32(0)
	goto L238
L242:
	;
	goto L241
L243:
	;
	v1429 = *(*int32)(unsafe.Add(mBase, uint32(v1338)))
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v1429
	goto L235
L244:
	;
	goto L232
L245:
	;
	v1471 = *(*int32)(unsafe.Add(mBase, uint32(v1338)+20))
	v1472 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1471)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v1471+v1472)+8)) = v1470
	v1475 = int32(0)
	v1476 = *(*int32)(unsafe.Add(mBase, uint32(v1338)+20))
	v1477 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1476)+16)))
	v1478 = v1476 + v1477
	v1479 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1478)+12)))
	if v49 != 0 {
		goto L248
	} else {
		goto L249
	}
L246:
	;
	v1466 = *(*int32)(unsafe.Add(mBase, uint32(v1338)))
	if v1466 == int32(0) {
		v1470 = v760
		goto L245
	} else {
		goto L247
	}
L247:
	;
	v1469 = *(*int32)(unsafe.Add(mBase, uint32(v1463)))
	v1470 = v1469
	goto L245
L248:
	;
	v1484 = int32(8)
	goto L250
L249:
	;
	v1484 = v1475
	goto L250
L250:
	;
	v1486 = *(*int32)(unsafe.Add(mBase, uint32(v1338)+28))
	if v1486 != 0 {
		goto L251
	} else {
		goto L252
	}
L251:
	;
	v1487 = v1484
	goto L253
L252:
	;
	v1487 = int32(0)
	goto L253
L253:
	;
	if v11 != 0 {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v1489 = v1487
	goto L256
L255:
	;
	v1489 = int32(0)
	goto L256
L256:
	;
	v1490 = v1479&int32(_a_F_gistplacetopage_15) | v1489
	*(*uint16)(unsafe.Add(mBase, uint32(v1478)+12)) = uint16(v1490)
	v1492 = *(*int32)(unsafe.Add(mBase, uint32(v1338)+20))
	v1493 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1492)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v1492+v1493))) = base.I32_wrap_i64(int64(base.Ui64(v761) >> (uint(int64(32)) % 64)))
	v1496 = *(*int32)(unsafe.Add(mBase, uint32(v1338)+20))
	v1497 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1496)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v1496+v1497)+4)) = base.I32_wrap_i64(v761)
	v1500 = *(*int32)(unsafe.Add(mBase, uint32(v1338)+28))
	if v1500 != 0 {
		v1338 = v1500
		goto L226
	} else {
		goto L257
	}
L257:
	;
	goto L227
L258:
	;
	v1543 = int32(_a_F_gistplacetopage_0)
	v1545 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[4])) = v1545 + int32(1)
	if v1515 == int32(0) {
		goto L267
	} else {
		goto L268
	}
L259:
	;
	v1527 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1527)+118)))
	if v1528 != int32(112) {
		goto L258
	} else {
		goto L260
	}
L260:
	;
	v1532 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[5]))
	if v1532 <= int32(0) {
		goto L261
	} else {
		goto L262
	}
L261:
	;
	v1535 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1535 != 0 {
		goto L258
	} else {
		goto L264
	}
L262:
	;
	goto L263
L263:
	;
	v1537 = int32(1)
	F_XLogEnsureRecordSpace(m, v697, v697<<(uint(v1537)%32)|v1537)
	mBase = m.M
	v1542 = m.ExcPending
	if v1542 != 0 {
		goto L59
	} else {
		goto L266
	}
L264:
	;
	v1536 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v1536 != 0 {
		goto L258
	} else {
		goto L265
	}
L265:
	;
	goto L263
L266:
	;
	goto L258
L267:
	;
	v1564 = v1507
	goto L270
L268:
	;
	goto L269
L269:
	;
	if l8 != 0 {
		goto L274
	} else {
		goto L275
	}
L270:
	;
	v1577 = *(*int32)(unsafe.Add(mBase, uint32(v1564)+24))
	F_MarkBufferDirty(m, v1577)
	mBase = m.M
	v1579 = m.ExcPending
	if v1579 != 0 {
		goto L59
	} else {
		goto L272
	}
L271:
	;
	goto L269
L272:
	;
	v1580 = *(*int32)(unsafe.Add(mBase, uint32(v1564)+28))
	if v1580 != 0 {
		v1564 = v1580
		goto L270
	} else {
		goto L273
	}
L273:
	;
	goto L271
L274:
	;
	F_MarkBufferDirty(m, l8)
	mBase = m.M
	v1608 = m.ExcPending
	if v1608 != 0 {
		goto L59
	} else {
		goto L277
	}
L275:
	;
	goto L276
L276:
	;
	v1609 = *(*int32)(unsafe.Add(mBase, uint32(v1507)+20))
	v1610 = *(*int32)(unsafe.Add(mBase, uint32(v1507)+24))
	if v1610 < int32(0) {
		goto L279
	} else {
		goto L280
	}
L277:
	;
	goto L276
L278:
	;
	F_PageRestoreTempPage(m, v1609, v1628)
	mBase = m.M
	v1630 = m.ExcPending
	if v1630 != 0 {
		goto L59
	} else {
		goto L282
	}
L279:
	;
	v1614 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[2]))
	v1620 = *(*int32)(unsafe.Add(mBase, uint32(v1614+(v1610^int32(-1))<<(uint(int32(2))%32))))
	v1628 = v1620
	goto L278
L280:
	;
	goto L281
L281:
	;
	v1622 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[3]))
	v1628 = v1622 + v1610<<(uint(int32(13))%32) + int32(-8192)
	goto L278
L282:
	;
	v1631 = *(*int32)(unsafe.Add(mBase, uint32(v1507)+24))
	if v1631 < int32(0) {
		goto L284
	} else {
		goto L285
	}
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1507)+20)) = v1649
	if l12 != 0 {
		v1844 = int64(1)
		goto L287
	} else {
		goto L288
	}
L284:
	;
	v1635 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[2]))
	v1641 = *(*int32)(unsafe.Add(mBase, uint32(v1635+(v1631^int32(-1))<<(uint(int32(2))%32))))
	v1649 = v1641
	goto L283
L285:
	;
	goto L286
L286:
	;
	v1643 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[3]))
	v1649 = v1643 + v1631<<(uint(int32(13))%32) + int32(-8192)
	goto L283
L287:
	;
	if v1515 == int32(0) {
		goto L319
	} else {
		goto L320
	}
L288:
	;
	v1652 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1652)+118)))
	if v1653 != int32(112) {
		goto L289
	} else {
		goto L290
	}
L289:
	;
	v1816 = F_gistGetFakeLSN(m, l0)
	mBase = m.M
	v1817 = m.ExcPending
	if v1817 != 0 {
		goto L59
	} else {
		goto L318
	}
L290:
	;
	v1657 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[5]))
	if v1657 <= int32(0) {
		goto L291
	} else {
		goto L292
	}
L291:
	;
	v1660 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1660 != 0 {
		goto L289
	} else {
		goto L294
	}
L292:
	;
	goto L293
L293:
	;
	v1662 = int32(0)
	v1663 = m.G0
	v1665 = v1663 - int32(32)
	m.G0 = v1665
	if v1507 != 0 {
		goto L296
	} else {
		goto L297
	}
L294:
	;
	v1661 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v1661 != 0 {
		goto L289
	} else {
		goto L295
	}
L295:
	;
	goto L293
L296:
	;
	v1668 = v1507
	v1669 = v1662
	goto L299
L297:
	;
	v1698 = v1662
	goto L298
L298:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1665)+28)) = uint8(v11)
	*(*uint16)(unsafe.Add(mBase, uint32(v1665)+26)) = uint16(v1698)
	*(*uint8)(unsafe.Add(mBase, uint32(v1665)+24)) = uint8(v701)
	*(*int64)(unsafe.Add(mBase, uint32(v1665)+16)) = v761
	*(*int32)(unsafe.Add(mBase, uint32(v1665)+8)) = v760
	F_XLogBeginInsert(m)
	mBase = m.M
	v1728 = m.ExcPending
	if v1728 != 0 {
		goto L59
	} else {
		goto L302
	}
L299:
	;
	v1694 = v1669 + int32(1)
	v1695 = *(*int32)(unsafe.Add(mBase, uint32(v1668)+28))
	if v1695 != 0 {
		v1668 = v1695
		v1669 = v1694
		goto L299
	} else {
		goto L301
	}
L300:
	;
	v1698 = v1694
	goto L298
L301:
	;
	goto L300
L302:
	;
	if l8 != 0 {
		goto L303
	} else {
		goto L304
	}
L303:
	;
	F_XLogRegisterBuffer(m, int32(0), l8, int32(8))
	mBase = m.M
	v1732 = m.ExcPending
	if v1732 != 0 {
		goto L59
	} else {
		goto L306
	}
L304:
	;
	goto L305
L305:
	;
	F_XLogRegisterData(m, v1665+int32(8), int32(24))
	mBase = m.M
	v1737 = m.ExcPending
	if v1737 != 0 {
		goto L59
	} else {
		goto L307
	}
L306:
	;
	goto L305
L307:
	;
	if v1507 != 0 {
		goto L308
	} else {
		goto L309
	}
L308:
	;
	v1739 = v1507
	v1741 = int32(1)
	goto L311
L309:
	;
	goto L310
L310:
	;
	v1811 = F_XLogInsert(m, int32(14), int32(48))
	mBase = m.M
	v1812 = m.ExcPending
	if v1812 != 0 {
		goto L59
	} else {
		goto L317
	}
L311:
	;
	v1766 = v1741 & int32(255)
	v1767 = *(*int32)(unsafe.Add(mBase, uint32(v1739)+24))
	F_XLogRegisterBuffer(m, v1766, v1767, int32(6))
	mBase = m.M
	v1770 = m.ExcPending
	if v1770 != 0 {
		goto L59
	} else {
		goto L313
	}
L312:
	;
	goto L310
L313:
	;
	v1771 = int32(4)
	F_XLogRegisterBufData(m, v1766, v1739+v1771, v1771)
	mBase = m.M
	v1775 = m.ExcPending
	if v1775 != 0 {
		goto L59
	} else {
		goto L314
	}
L314:
	;
	v1776 = *(*int32)(unsafe.Add(mBase, uint32(v1739)+8))
	v1777 = *(*int32)(unsafe.Add(mBase, uint32(v1739)+12))
	F_XLogRegisterBufData(m, v1766, v1776, v1777)
	mBase = m.M
	v1779 = m.ExcPending
	if v1779 != 0 {
		goto L59
	} else {
		goto L315
	}
L315:
	;
	v1782 = *(*int32)(unsafe.Add(mBase, uint32(v1739)+28))
	if v1782 != 0 {
		v1739 = v1782
		v1741 = v1741 + int32(1)
		goto L311
	} else {
		goto L316
	}
L316:
	;
	goto L312
L317:
	;
	m.G0 = v1665 + int32(32)
	v1844 = v1811
	goto L287
L318:
	;
	v1844 = v1816
	goto L287
L319:
	;
	v1864 = v1507
	goto L322
L320:
	;
	goto L321
L321:
	;
	if v49 != 0 {
		goto L325
	} else {
		goto L326
	}
L322:
	;
	v1877 = *(*int32)(unsafe.Add(mBase, uint32(v1864)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1877)+4)) = base.I32_wrap_i64(v1844)
	*(*int32)(unsafe.Add(mBase, uint32(v1877))) = base.I32_wrap_i64(int64(base.Ui64(v1844) >> (uint(int64(32)) % 64)))
	v1880 = *(*int32)(unsafe.Add(mBase, uint32(v1864)+28))
	if v1880 != 0 {
		v1864 = v1880
		goto L322
	} else {
		goto L324
	}
L323:
	;
	goto L321
L324:
	;
	goto L323
L325:
	;
	v1981 = int32(1)
	v1991 = v1844
	goto L6
L326:
	;
	v1907 = *(*int32)(unsafe.Add(mBase, uint32(v1507)+28))
	if v1907 == int32(0) {
		goto L325
	} else {
		goto L327
	}
L327:
	;
	v1923 = v1907
	goto L328
L328:
	;
	v1936 = *(*int32)(unsafe.Add(mBase, uint32(v1923)+24))
	F_UnlockReleaseBuffer(m, v1936)
	mBase = m.M
	v1938 = m.ExcPending
	if v1938 != 0 {
		goto L59
	} else {
		goto L330
	}
L329:
	;
	goto L325
L330:
	;
	v1939 = *(*int32)(unsafe.Add(mBase, uint32(v1923)+28))
	if v1939 != 0 {
		v1923 = v1939
		goto L328
	} else {
		goto L331
	}
L331:
	;
	goto L329
L332:
	;
	if l8 < int32(0) {
		goto L336
	} else {
		goto L337
	}
L333:
	;
	goto L334
L334:
	;
	v2033 = int32(_a_F_gistplacetopage_0)
	v2035 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[4])) = v2035 - int32(1)
	m.G0 = v29 + int32(864)
	return v1981
L335:
	;
	v2011 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2010)+16)))
	v2015 = base.I32_wrap_i64(int64(base.Ui64(v1991) >> (uint(int64(32)) % 64)))
	*(*int32)(unsafe.Add(mBase, uint32(v2011+v2010))) = v2015
	v2017 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2010)+16)))
	v2019 = base.I32_wrap_i64(v1991)
	*(*int32)(unsafe.Add(mBase, uint32(v2010+v2017)+4)) = v2019
	v2021 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2010)+16)))
	v2022 = v2010 + v2021
	v2023 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2022)+12)))
	v2025 = v2023 & int32(_a_F_gistplacetopage_15)
	*(*uint16)(unsafe.Add(mBase, uint32(v2022)+12)) = uint16(v2025)
	*(*int32)(unsafe.Add(mBase, uint32(v2010)+4)) = v2019
	*(*int32)(unsafe.Add(mBase, uint32(v2010))) = v2015
	goto L334
L336:
	;
	v1996 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[2]))
	v2002 = *(*int32)(unsafe.Add(mBase, uint32(v1996+(l8^int32(-1))<<(uint(int32(2))%32))))
	v2010 = v2002
	goto L335
L337:
	;
	goto L338
L338:
	;
	v2004 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[3]))
	v2010 = v2004 + l8<<(uint(int32(13))%32) + int32(-8192)
	goto L335
L339:
	;
	v2047 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v2047 + int32(4)
	F_errmsg_internal(m, int32(_a_F_gistplacetopage_10), v29+int32(16))
	mBase = m.M
	v2055 = m.ExcPending
	if v2055 != 0 {
		goto L59
	} else {
		goto L340
	}
L340:
	;
	F_errfinish(m, int32(_a_F_gistplacetopage_11), int32(434), int32(_a_F_gistplacetopage_12))
	mBase = m.M
	v2060 = m.ExcPending
	if v2060 != 0 {
		goto L59
	} else {
		goto L341
	}
L341:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gseg_penalty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 float32
	_ = v2
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 float32
	_ = v21
	var v22 float32
	_ = v22
	var v28 float32
	_ = v28
	var v29 int32
	_ = v29
	var v32 float32
	_ = v32
	var v33 float32
	_ = v33
	var v39 float32
	_ = v39
	v2 = float32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v15 = F_DirectFunctionCall2Coll(m, int32(_a_F_gseg_penalty_0), int32(0), v12, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		if v15 == int32(0) {
			v28 = v2
		} else {
			v21 = *(*float32)(unsafe.Add(mBase, uint32(v15)+4))
			v22 = *(*float32)(unsafe.Add(mBase, uint32(v15)))
			if base.F32_le(v21, v22) != 0 {
				v28 = v2
			} else {
				v28 = base.F32_abs(base.F32_sub(v21, v22))
			}
		}
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
		if v29 == int32(0) {
			v39 = v2
		} else {
			v32 = *(*float32)(unsafe.Add(mBase, uint32(v29)+4))
			v33 = *(*float32)(unsafe.Add(mBase, uint32(v29)))
			if base.F32_le(v32, v33) != 0 {
				v39 = v2
			} else {
				v39 = base.F32_abs(base.F32_sub(v32, v33))
			}
		}
		*(*float32)(unsafe.Add(mBase, uint32(v8))) = base.F32_sub(v28, v39)
		return v8
	}
}
func F_gtsvectorin(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13874(m, l0, int32(_a_F_gtsvectorin_0), int32(94), int32(_a_F_gtsvectorin_1), int32(_a_F_gtsvectorin_2), int32(_a_F_gtsvectorin_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
