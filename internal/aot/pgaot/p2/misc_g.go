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
			*(*int32)(unsafe.Add(mBase, _c_F_GetComboCommandId[2])) = int32(100)
			*(*int32)(unsafe.Add(mBase, _c_F_GetComboCommandId[3])) = v18
			*(*int32)(unsafe.Add(mBase, _c_F_GetComboCommandId[4])) = int32(0)
			*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = int64(51539607560)
			v33 = *(*int32)(unsafe.Add(mBase, _c_F_GetComboCommandId[1]))
			*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = v33
			v41 = F_hash_create(m, int32(_a_F_GetComboCommandId_0), int64(100), v7+int32(-56), int32(1064))
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
						*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l0
						v72 = F_hash_search(m, v64, v7+int32(-56), int32(1), v7+int32(-1))
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
					*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l0
					v72 = F_hash_search(m, v64, v7+int32(-56), int32(1), v7+int32(-1))
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
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l0
				v72 = F_hash_search(m, v64, v7+int32(-56), int32(1), v7+int32(-1))
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
			*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l0
			v72 = F_hash_search(m, v64, v7+int32(-56), int32(1), v7+int32(-1))
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
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v11 = F_SearchSysCache1(m, int32(32), base.I64_extend_i32_u(l0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 != 0 {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+22)))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v15+v16)+72))
			F_ReleaseCatCache(m, v11)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v23 = F_SearchSysCache1(m, int32(30), base.I64_extend_i32_u(v18))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					if v23 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v18
							F_errmsg_internal(m, int32(_a_F_GetFdwRoutineByServerId_0), v7+int32(16))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_GetFdwRoutineByServerId_1), int32(440), int32(_a_F_GetFdwRoutineByServerId_2))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
						v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+22)))
						v29 = v27 + v28
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+72))
						if v30 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(325))
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v29 + int32(4)
									F_errmsg(m, int32(_a_F_GetFdwRoutineByServerId_3), v7+int32(32))
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_GetFdwRoutineByServerId_1), int32(449), int32(_a_F_GetFdwRoutineByServerId_2))
										mBase = m.M
										v88 = m.ExcPending
										if v88 != 0 {
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
							F_ReleaseCatCache(m, v23)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								v35 = F_GetFdwRoutine(m, v30)
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return int32(0)
								} else {
									m.G0 = v7 + int32(48)
									return v35
								}
							}
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
				F_errmsg_internal(m, int32(_a_F_GetFdwRoutineByServerId_4), v7)
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_GetFdwRoutineByServerId_1), int32(432), int32(_a_F_GetFdwRoutineByServerId_2))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_GetLatestXTime[0]))
	v7 = base.AtomicRmwXchg32(m, v4, int32(96), int32(1))
	if v7 != 0 {
		F_s_lock(m, v4+int32(96), int32(_a_F_GetLatestXTime_0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, _c_F_GetLatestXTime[0]))
			v17 = *(*int64)(unsafe.Add(mBase, uint32(v16)+64))
			v18 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v16)+96)), uint32(v18))
			return v17
		}
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_GetLatestXTime[0]))
		v17 = *(*int64)(unsafe.Add(mBase, uint32(v16)+64))
		v18 = int32(0)
		atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v16)+96)), uint32(v18))
		return v17
	}
}
func F_GetNamedDSMSegment(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	v6 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetNamedDSMSegment[0])))
	if v6 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L9
	} else {
		goto L52
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L9
	} else {
		goto L49
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L9
	} else {
		goto L46
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L9
	} else {
		goto L43
	}
L5:
	;
	v8 = F_strlen(m, int32(_a_F_GetNamedDSMSegment_0))
	mBase = m.M
	if base.Ui32(int32(64)) <= base.Ui32(v8) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L9
	} else {
		goto L40
	}
L8:
	;
	v11 = int32(_a_F_GetNamedDSMSegment_1)
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_GetNamedDSMSegment[1]))
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_GetNamedDSMSegment[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_GetNamedDSMSegment[1])) = v15
	F_init_dsm_registry(m)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_GetNamedDSMSegment[3]))
	v24 = F_dshash_find_or_insert_extended(m, v22, int32(_a_F_GetNamedDSMSegment_0), l0)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v26 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v86)+24))
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_GetNamedDSMSegment[3]))
	F_dshash_release_lock(m, v90, v24)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L9
	} else {
		goto L39
	}
L13:
	;
	v56 = int32(0)
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_GetNamedDSMSegment[4]))
	if base.B2i32(v58 == v56)|base.B2i32(v58 == int32(_a_F_GetNamedDSMSegment_2)) == v56 {
		goto L26
	} else {
		goto L27
	}
L14:
	;
	v39 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v39)
	v43 = F_dsm_create(m, int32(44), v39)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L9
	} else {
		goto L21
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+72)) = int32(44)
	*(*int64)(unsafe.Add(mBase, uint32(v24)+64)) = int64(0)
	goto L14
L16:
	;
	goto L17
L17:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v24)+64))
	if v33 != 0 {
		goto L3
	} else {
		goto L18
	}
L18:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v24)+72))
	if v34 != int32(44) {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v24)+68))
	if v37 != 0 {
		goto L13
	} else {
		goto L20
	}
L20:
	;
	goto L14
L21:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v43)+24))
	m.T0[int32(_a_F_GetNamedDSMSegment_3)].(func(*base.Module, int32, int32))(m, v45, int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L9
	} else {
		goto L22
	}
L22:
	;
	F_dsm_pin_segment(m, v43)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L9
	} else {
		goto L23
	}
L23:
	;
	F_dsm_pin_mapping(m, v43)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L9
	} else {
		goto L24
	}
L24:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+68)) = v54
	v86 = v43
	goto L12
L25:
	;
	if v78 != 0 {
		v86 = v78
		goto L12
	} else {
		goto L35
	}
L26:
	;
	v67 = v58
	goto L29
L27:
	;
	goto L28
L28:
	;
	v78 = int32(0)
	goto L25
L29:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	if v37 == v68 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L28
L31:
	;
	v78 = v67
	goto L25
L32:
	;
	goto L33
L33:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	if v70 != int32(_a_F_GetNamedDSMSegment_2) {
		v67 = v70
		goto L29
	} else {
		goto L34
	}
L34:
	;
	goto L30
L35:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v24)+68))
	v80 = F_dsm_attach(m, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L9
	} else {
		goto L36
	}
L36:
	;
	if v80 == int32(0) {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_dsm_pin_mapping(m, v80)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L9
	} else {
		goto L38
	}
L38:
	;
	v86 = v80
	goto L12
L39:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetNamedDSMSegment[1])) = v12
	return v88
L40:
	;
	F_errmsg(m, int32(_a_F_GetNamedDSMSegment_4), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L9
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_GetNamedDSMSegment_5), int32(204), int32(_a_F_GetNamedDSMSegment_6))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L9
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	F_errmsg(m, int32(_a_F_GetNamedDSMSegment_7), int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L9
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(_a_F_GetNamedDSMSegment_5), int32(208), int32(_a_F_GetNamedDSMSegment_6))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L9
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L46:
	;
	F_errmsg(m, int32(_a_F_GetNamedDSMSegment_8), int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L9
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_GetNamedDSMSegment_5), int32(230), int32(_a_F_GetNamedDSMSegment_6))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L9
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	F_errmsg(m, int32(_a_F_GetNamedDSMSegment_9), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L9
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_GetNamedDSMSegment_5), int32(233), int32(_a_F_GetNamedDSMSegment_6))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L9
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L52:
	;
	F_errmsg_internal(m, int32(_a_F_GetNamedDSMSegment_10), int32(0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L9
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_GetNamedDSMSegment_5), int32(257), int32(_a_F_GetNamedDSMSegment_6))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L9
	} else {
		goto L54
	}
L54:
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
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v145 int64
	_ = v145
	var v157 int32
	_ = v157
	var v168 int64
	_ = v168
	var v170 int32
	_ = v170
	var v171 int64
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v187 int64
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int64
	_ = v193
	var v194 int64
	_ = v194
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v220 int64
	_ = v220
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v233 int64
	_ = v233
	var v235 int64
	_ = v235
	var v236 int64
	_ = v236
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v342 int32
	_ = v342
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v476 int32
	_ = v476
	var v477 int64
	_ = v477
	var v479 int32
	_ = v479
	var v480 int64
	_ = v480
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v559 int32
	_ = v559
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v609 int32
	_ = v609
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v650 int32
	_ = v650
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v669 int32
	_ = v669
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v706 int32
	_ = v706
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v763 int32
	_ = v763
	var v769 int64
	_ = v769
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v817 int32
	_ = v817
	var v821 int32
	_ = v821
	var v825 int32
	_ = v825
	var v828 int32
	_ = v828
	var v832 int32
	_ = v832
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v841 int32
	_ = v841
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v852 int32
	_ = v852
	var v856 int32
	_ = v856
	var v861 int32
	_ = v861
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
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v396)))
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v396)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v410)+4)) = v411
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v396)))
	*(*int32)(unsafe.Add(mBase, uint32(v411))) = v413
	v416 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[1]))
	v418 = v416 + int32(8)
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v416)+12))
	if v419 == int32(0) {
		goto L79
	} else {
		goto L80
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
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[2]))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+44))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+40))
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[3]))
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
	v383 = m.ExcPending
	if v383 != 0 {
		goto L9
	} else {
		goto L76
	}
L9:
	;
	return int32(0)
L10:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[1]))
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
		v396 = v48
		goto L1
	} else {
		goto L14
	}
L14:
	;
	goto L15
L15:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[3]))
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
	v396 = v375
	goto L1
L17:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[3]))
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
	v361 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[3]))
	F_LWLockRelease(m, v361+int32(3712))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L9
	} else {
		goto L73
	}
L22:
	;
	F_ReleaseOneSerializableXact(m, v84-int32(56), int32(0), int32(1))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L9
	} else {
		goto L72
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
	v114 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v100
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[5]))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+28))
	v120 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[6])))
	v124 = F_LWLockAcquire(m, v114+int32(_a_F_GetSerializableTransactionSnapshotInt_0), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	v127 = int32(base.Ui32(v100) >> (uint(int32(10)) % 32))
	v128 = base.I32_rem_u_s(v127, v120)
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[7]))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+12))
	if v131 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v317 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[3]))
	F_LWLockRelease(m, v317+int32(_a_F_GetSerializableTransactionSnapshotInt_0))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L9
	} else {
		goto L71
	}
L30:
	;
	v134 = int32(3)
	if base.B2i32(base.Ui32(v100) < base.Ui32(v134))|base.B2i32(base.Ui32(v131) < base.Ui32(v134)) == int32(0) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v145 = *(*int64)(unsafe.Add(mBase, uint32(v130)))
	if v145 < int64(0) {
		goto L38
	} else {
		goto L39
	}
L32:
	;
	if int32(0) <= v100-v131 {
		goto L31
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	if base.Ui32(v100) < base.Ui32(v131) {
		goto L29
	} else {
		goto L36
	}
L35:
	;
	goto L29
L36:
	;
	goto L31
L37:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
	if v172 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L38:
	;
	v170 = int32(1)
	v171 = base.I64_extend_i32_u(int32(base.Ui32(v131) >> (uint(int32(10)) % 32)))
	goto L37
L39:
	;
	goto L40
L40:
	;
	v157 = base.I32_wrap_i64(v145)<<(uint(int32(10))%32) - v100&int32(-1024)
	if base.Ui64(v145) <= base.Ui64(int64(4194302)) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v168 = v145 + int64(1)
	goto L43
L42:
	;
	v168 = int64(0)
	goto L43
L43:
	;
	v170 = int32(base.Ui32((v157-int32(1023))&v157) >> (uint(int32(31)) % 32))
	v171 = v168
	goto L37
L44:
	;
	v187 = base.I64_extend_i32_u(v127)
	if v170 != 0 {
		goto L53
	} else {
		goto L54
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v130)+8)) = v100
	goto L44
L46:
	;
	v175 = int32(3)
	if base.B2i32(base.Ui32(v100) < base.Ui32(v175))|base.B2i32(base.Ui32(v172) < base.Ui32(v175)) == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	if int32(0) < v100-v172 {
		goto L45
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	if base.Ui32(v100) <= base.Ui32(v172) {
		goto L44
	} else {
		goto L51
	}
L50:
	;
	goto L44
L51:
	;
	goto L45
L52:
	;
	v278 = int32(_a_F_GetSerializableTransactionSnapshotInt_1)
	v279 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[5]))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v279)+4))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v280+v266<<(uint(int32(2))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v284+v265<<(uint(int32(3))%32)&int32(_a_F_GetSerializableTransactionSnapshotInt_2)))) = v112
	v292 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[5]))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v292)+12))
	v295 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v293+v266))) = uint8(v295)
	F_LWLockRelease(m, v264)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L9
	} else {
		goto L70
	}
L53:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v130))) = v187
	v190 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[5]))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v190)+28))
	v193 = int64(*(*uint16)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[6])))
	v194 = base.I64_rem_u_s(v171, v193)
	v198 = v191 + base.I32_wrap_i64(v194)<<(uint(int32(7))%32)
	v200 = F_LWLockAcquire(m, v198, int32(0))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L9
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v250 = v118 + v128<<(uint(int32(7))%32)
	v252 = F_LWLockAcquire(m, v250, int32(0))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L9
	} else {
		goto L68
	}
L56:
	;
	v203 = F_SimpleLruZeroPage(m, int32(_a_F_GetSerializableTransactionSnapshotInt_3), v171)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L9
	} else {
		goto L57
	}
L57:
	;
	if v171 == v187 {
		v264 = v198
		v265 = v100
		v266 = v203
		goto L52
	} else {
		goto L58
	}
L58:
	;
	v209 = v198
	v220 = v171
	goto L59
L59:
	;
	F_LWLockRelease(m, v209)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L9
	} else {
		goto L61
	}
L60:
	;
	v264 = v240
	v265 = v100
	v266 = v245
	goto L52
L61:
	;
	v226 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[5]))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+28))
	if v220 <= int64(4194302) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v233 = v220 + int64(1)
	goto L64
L63:
	;
	v233 = int64(0)
	goto L64
L64:
	;
	v235 = int64(*(*uint16)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[6])))
	v236 = base.I64_rem_s(v233, v235)
	v240 = v227 + base.I32_wrap_i64(v236)<<(uint(int32(7))%32)
	v242 = F_LWLockAcquire(m, v240, int32(0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L9
	} else {
		goto L65
	}
L65:
	;
	v245 = F_SimpleLruZeroPage(m, int32(_a_F_GetSerializableTransactionSnapshotInt_3), v233)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L9
	} else {
		goto L66
	}
L66:
	;
	if v233 != v187 {
		v209 = v240
		v220 = v233
		goto L59
	} else {
		goto L67
	}
L67:
	;
	goto L60
L68:
	;
	v258 = F_SimpleLruReadPage(m, int32(_a_F_GetSerializableTransactionSnapshotInt_3), v187, int32(1), v18+int32(-48))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L9
	} else {
		goto L69
	}
L69:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v264 = v250
	v265 = v260
	v266 = v258
	goto L52
L70:
	;
	goto L29
L71:
	;
	goto L22
L72:
	;
	goto L21
L73:
	;
	v367 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[3]))
	v371 = F_LWLockAcquire(m, v367+int32(3584), int32(0))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L9
	} else {
		goto L74
	}
L74:
	;
	v374 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[1]))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v374)+4))
	if base.B2i32(v375 == int32(0))|base.B2i32(v375 == v374) != 0 {
		goto L15
	} else {
		goto L75
	}
L75:
	;
	goto L16
L76:
	;
	F_errmsg_internal(m, int32(_a_F_GetSerializableTransactionSnapshotInt_4), int32(0))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L9
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F_GetSerializableTransactionSnapshotInt_5), int32(1713), int32(_a_F_GetSerializableTransactionSnapshotInt_6))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L9
	} else {
		goto L78
	}
L78:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v416)+12)) = v418
	*(*int32)(unsafe.Add(mBase, uint32(v416)+8)) = v418
	goto L81
L80:
	;
	goto L81
L81:
	;
	v425 = v396 + int32(-64)
	*(*int32)(unsafe.Add(mBase, uint32(v396)+4)) = v418
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v418)))
	*(*int32)(unsafe.Add(mBase, uint32(v396))) = v427
	*(*int32)(unsafe.Add(mBase, uint32(v427)+4)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v418))) = v396
	if l1 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L82:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L9
	} else {
		goto L160
	}
L83:
	;
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v425)+64))
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v425)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v797)+4)) = v798
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v425)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v798))) = v800
	v803 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[1]))
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v803)+4))
	if v804 == int32(0) {
		goto L151
	} else {
		goto L152
	}
L84:
	;
	v442 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[8])))
	if v442 != int32(1) {
		goto L92
	} else {
		goto L93
	}
L85:
	;
	v433 = F_GetSnapshotData(m, l0)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L9
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v436 = F_ProcArrayInstallImportedXmin(m, v435, l1)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L9
	} else {
		goto L89
	}
L88:
	;
	v440 = v433
	goto L84
L89:
	;
	if v436 == int32(0) {
		goto L83
	} else {
		goto L90
	}
L90:
	;
	v440 = l0
	goto L84
L91:
	;
	m.G0 = v20 - int32(-64)
	return v440
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v425))) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v396-int32(60)))) = v35
	v476 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[1]))
	v477 = *(*int64)(unsafe.Add(mBase, uint32(v476)+32))
	v479 = v396 - int32(56)
	v480 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v479))) = v480
	*(*int64)(unsafe.Add(mBase, uint32(v396-int32(40)))) = v477
	*(*int64)(unsafe.Add(mBase, uint32(v479)+8)) = v480
	v487 = int32(24)
	v488 = v396 + v487
	*(*int32)(unsafe.Add(mBase, uint32(v396)+28)) = v488
	*(*int32)(unsafe.Add(mBase, uint32(v396)+24)) = v488
	v494 = v396 - v487
	*(*int32)(unsafe.Add(mBase, uint32(v396-int32(20)))) = v494
	v499 = v396 - int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v396-int32(28)))) = v499
	*(*int32)(unsafe.Add(mBase, uint32(v494))) = v494
	*(*int32)(unsafe.Add(mBase, uint32(v499))) = v499
	v504 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[9]))
	v505 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v396)+36)) = v505
	*(*int32)(unsafe.Add(mBase, uint32(v396)+32)) = v504
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v440)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v396)+40)) = v508
	v511 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[10]))
	*(*int32)(unsafe.Add(mBase, uint32(v396)+48)) = v511
	v514 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[11]))
	*(*int32)(unsafe.Add(mBase, uint32(v396)+44)) = v505
	*(*int64)(unsafe.Add(mBase, uint32(v396-int32(8)))) = int64(0)
	v524 = v396 - int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v396-int32(12)))) = v524
	*(*int32)(unsafe.Add(mBase, uint32(v396)+52)) = v514
	*(*int32)(unsafe.Add(mBase, uint32(v524))) = v524
	v529 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[8])))
	if v529 == int32(1) {
		goto L100
	} else {
		goto L101
	}
L93:
	;
	v446 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[1]))
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v446)+24))
	if v447 != 0 {
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v396)))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v396)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v448)+4)) = v449
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v396)))
	*(*int32)(unsafe.Add(mBase, uint32(v449))) = v451
	v454 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[1]))
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v454)+4))
	if v455 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v454)+4)) = v454
	*(*int32)(unsafe.Add(mBase, uint32(v454))) = v454
	goto L97
L96:
	;
	goto L97
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v396)+4)) = v454
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v454)))
	*(*int32)(unsafe.Add(mBase, uint32(v396))) = v461
	*(*int32)(unsafe.Add(mBase, uint32(v461)+4)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v454))) = v396
	v466 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[3]))
	F_LWLockRelease(m, v466+int32(3584))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L9
	} else {
		goto L98
	}
L98:
	;
	goto L91
L99:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v440)+4))
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v669)+16))
	if v684 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v396)+44)) = int32(32)
	v535 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[1]))
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v535)+12))
	if v536 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	goto L102
L102:
	;
	v661 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[1]))
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v661)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v661)+24)) = v662 + int32(1)
	v669 = v661
	goto L99
L103:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v396)+28))
	v629 = int32(0)
	if base.B2i32(v628 == v629)|base.B2i32(v628 == v488) == v629 {
		goto L119
	} else {
		goto L120
	}
L104:
	;
	v540 = v535 + int32(8)
	if v536 == v540 {
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v544 = v536
	goto L106
L106:
	;
	v559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v544)+44)))
	if v559&int32(41) == int32(0) {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	goto L103
L108:
	;
	v565 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[12]))
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v565)+4))
	if base.B2i32(v566 == int32(0))|base.B2i32(v566 == v565) != 0 {
		goto L82
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v544)+4))
	if v609 != v540 {
		v544 = v609
		goto L106
	} else {
		goto L118
	}
L111:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v566)))
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v566)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v571)+4)) = v572
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v566)))
	*(*int32)(unsafe.Add(mBase, uint32(v572))) = v574
	*(*int32)(unsafe.Add(mBase, uint32(v566)+20)) = v425
	*(*int32)(unsafe.Add(mBase, uint32(v566)+16)) = v544 + int32(-64)
	v581 = v544 + int32(24)
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v544)+28))
	if v582 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v544)+28)) = v581
	*(*int32)(unsafe.Add(mBase, uint32(v544)+24)) = v581
	goto L114
L113:
	;
	goto L114
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v566)+4)) = v581
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v581)))
	*(*int32)(unsafe.Add(mBase, uint32(v566))) = v588
	*(*int32)(unsafe.Add(mBase, uint32(v588)+4)) = v566
	*(*int32)(unsafe.Add(mBase, uint32(v581))) = v566
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v396)+28))
	if v592 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v396)+28)) = v488
	*(*int32)(unsafe.Add(mBase, uint32(v396)+24)) = v396 + int32(24)
	goto L117
L116:
	;
	goto L117
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v566)+12)) = v488
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v488)))
	*(*int32)(unsafe.Add(mBase, uint32(v566)+8)) = v600
	v603 = v566 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v600)+4)) = v603
	*(*int32)(unsafe.Add(mBase, uint32(v488))) = v603
	goto L110
L118:
	;
	goto L107
L119:
	;
	v636 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[1]))
	v669 = v636
	goto L99
L120:
	;
	goto L121
L121:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v396)))
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v396)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v637)+4)) = v638
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v396)))
	*(*int32)(unsafe.Add(mBase, uint32(v638))) = v640
	v643 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[1]))
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v643)+4))
	if v644 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v643)+4)) = v643
	*(*int32)(unsafe.Add(mBase, uint32(v643))) = v643
	goto L124
L123:
	;
	goto L124
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v396)+4)) = v643
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v643)))
	*(*int32)(unsafe.Add(mBase, uint32(v396))) = v650
	*(*int32)(unsafe.Add(mBase, uint32(v650)+4)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v643))) = v396
	v655 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[3]))
	F_LWLockRelease(m, v655+int32(3584))
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L9
	} else {
		goto L125
	}
L125:
	;
	goto L91
L126:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[13])) = v425
	v756 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[14])) = uint8(v756)
	v759 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[3]))
	F_LWLockRelease(m, v759+int32(3584))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L9
	} else {
		goto L149
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v669)+20)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v669)+16)) = v683
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v440)+4))
	v692 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[3]))
	v696 = F_LWLockAcquire(m, v692+int32(_a_F_GetSerializableTransactionSnapshotInt_0), int32(0))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L9
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	if v683 != v684 {
		goto L126
	} else {
		goto L148
	}
L130:
	;
	if v690 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	v740 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[3]))
	F_LWLockRelease(m, v740+int32(_a_F_GetSerializableTransactionSnapshotInt_0))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L9
	} else {
		goto L147
	}
L132:
	;
	v701 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[7]))
	*(*int64)(unsafe.Add(mBase, uint32(v701)+8)) = int64(0)
	goto L131
L133:
	;
	goto L134
L134:
	;
	v706 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[15])))
	if v706 == int32(1) {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	v718 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[7]))
	if v716 == int32(0) {
		goto L139
	} else {
		goto L140
	}
L136:
	;
	v711 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[16]))
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v711)+308))
	v714 = base.B2i32(v712 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[15])) = uint8(v714)
	v716 = v714
	goto L138
L137:
	;
	v716 = int32(0)
	goto L138
L138:
	;
	goto L135
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v718)+12)) = v690
	goto L131
L140:
	;
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v718)+12))
	if v721 == int32(0) {
		goto L139
	} else {
		goto L141
	}
L141:
	;
	v724 = int32(3)
	if base.B2i32(base.Ui32(v690) < base.Ui32(v724))|base.B2i32(base.Ui32(v721) < base.Ui32(v724)) == int32(0) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	if v690-v721 < int32(0) {
		goto L139
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	if base.Ui32(v721) <= base.Ui32(v690) {
		goto L131
	} else {
		goto L146
	}
L145:
	;
	goto L131
L146:
	;
	goto L139
L147:
	;
	goto L126
L148:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v669)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v669)+20)) = v746 + int32(1)
	goto L126
L149:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+24)) = int64(103079215120)
	v769 = int64(*(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[17])))
	v773 = F_hash_create(m, int32(_a_F_GetSerializableTransactionSnapshotInt_7), v769, v18+int32(-48), int32(40))
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L9
	} else {
		goto L150
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[18])) = v773
	goto L91
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v803)+4)) = v803
	*(*int32)(unsafe.Add(mBase, uint32(v803))) = v803
	goto L153
L152:
	;
	goto L153
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v425)+68)) = v803
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v803)))
	*(*int32)(unsafe.Add(mBase, uint32(v425)+64)) = v810
	v813 = v425 - int32(-64)
	*(*int32)(unsafe.Add(mBase, uint32(v810)+4)) = v813
	*(*int32)(unsafe.Add(mBase, uint32(v803))) = v813
	v817 = *(*int32)(unsafe.Add(mBase, _c_F_GetSerializableTransactionSnapshotInt[3]))
	F_LWLockRelease(m, v817+int32(3584))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L9
	} else {
		goto L154
	}
L154:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L9
	} else {
		goto L155
	}
L155:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L9
	} else {
		goto L156
	}
L156:
	;
	F_errmsg(m, int32(_a_F_GetSerializableTransactionSnapshotInt_8), int32(0))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L9
	} else {
		goto L157
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = l2
	v835 = F_errdetail(m, int32(_a_F_GetSerializableTransactionSnapshotInt_9), v20)
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L9
	} else {
		goto L158
	}
L158:
	;
	F_errfinish(m, int32(_a_F_GetSerializableTransactionSnapshotInt_5), int32(1758), int32(_a_F_GetSerializableTransactionSnapshotInt_6))
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L9
	} else {
		goto L159
	}
L159:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L160:
	;
	F_errcode(m, int32(_a_F_GetSerializableTransactionSnapshotInt_10))
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L9
	} else {
		goto L161
	}
L161:
	;
	F_errmsg(m, int32(_a_F_GetSerializableTransactionSnapshotInt_11), int32(0))
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L9
	} else {
		goto L162
	}
L162:
	;
	F_errhint(m, int32(_a_F_GetSerializableTransactionSnapshotInt_12), int32(0))
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L9
	} else {
		goto L163
	}
L163:
	;
	F_errfinish(m, int32(_a_F_GetSerializableTransactionSnapshotInt_5), int32(693), int32(_a_F_GetSerializableTransactionSnapshotInt_13))
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L9
	} else {
		goto L164
	}
L164:
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
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int64
	_ = v65
	var v69 int32
	_ = v69
	var v70 int64
	_ = v70
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int64
	_ = v109
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v190 int32
	_ = v190
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
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
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v448 int64
	_ = v448
	var v450 int64
	_ = v450
	var v455 int64
	_ = v455
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
	var v470 int64
	_ = v470
	var v472 int64
	_ = v472
	var v473 int64
	_ = v473
	var v478 int32
	_ = v478
	var v480 int64
	_ = v480
	var v482 int64
	_ = v482
	var v484 int64
	_ = v484
	var v486 int64
	_ = v486
	var v494 int64
	_ = v494
	var v502 int64
	_ = v502
	var v506 int64
	_ = v506
	var v509 int64
	_ = v509
	var v511 int64
	_ = v511
	var v513 int64
	_ = v513
	var v514 int32
	_ = v514
	var v516 int64
	_ = v516
	var v519 int64
	_ = v519
	var v521 int64
	_ = v521
	var v522 int64
	_ = v522
	var v524 int32
	_ = v524
	var v526 int64
	_ = v526
	var v529 int64
	_ = v529
	var v531 int64
	_ = v531
	var v532 int64
	_ = v532
	var v535 int64
	_ = v535
	var v538 int64
	_ = v538
	var v540 int64
	_ = v540
	var v541 int64
	_ = v541
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v597 int32
	_ = v597
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
	F_emscripten_builtin_free(m, v39)
	mBase = m.M
	v580 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v580
	F_errstart_cold(m, int32(21), v580)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L8
	} else {
		goto L170
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L8
	} else {
		goto L166
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
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[2]))
	v61 = F_LWLockAcquire(m, v57+int32(512), int32(1))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
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
	v65 = *(*int64)(unsafe.Add(mBase, uint32(l0)+64))
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[5]))
	v70 = *(*int64)(unsafe.Add(mBase, uint32(v69)+56))
	if base.B2i32(v65 == int64(0))|base.B2i32(v65 != v70) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[6]))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+52))
	if v78 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[6]))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+32))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v32+v103<<(uint(int32(2))%32))))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
	v109 = *(*int64)(unsafe.Add(mBase, uint32(v69)+48))
	v112 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetSnapshotData[7])))
	if v112 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[8])) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v77)+52)) = v75
	goto L15
L14:
	;
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[9])) = v75
	v87 = F_GetCurrentCommandId(m, int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+44)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v87
	v92 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)) = uint8(v92)
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[2]))
	F_LWLockRelease(m, v95+int32(512))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	return l0
L18:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)) = uint8(v122)
	v124 = int32(3)
	v125 = base.I32_wrap_i64(v109)
	v127 = v125 + int32(1)
	if base.Ui32(v127) <= base.Ui32(v124) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[10]))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+308))
	v120 = base.B2i32(v118 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_GetSnapshotData[7])) = uint8(v120)
	v122 = v120
	goto L21
L20:
	;
	v122 = int32(0)
	goto L21
L21:
	;
	goto L18
L22:
	;
	v130 = v124
	goto L24
L23:
	;
	v130 = v127
	goto L24
L24:
	;
	if v107-v130 < int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v134 = v107
	goto L27
L26:
	;
	v134 = v130
	goto L27
L27:
	;
	if base.Ui32(int32(2)) < base.Ui32(v107) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v137 = v134
	goto L30
L29:
	;
	v137 = v130
	goto L30
L30:
	;
	if v122 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v397 = *(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[0]))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v397)+32))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v397)+28))
	v401 = *(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[6]))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v401)+52))
	if v402 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L32:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if v140 <= int32(0) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v248 = *(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[0]))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v248)+20))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v248)+16))
	v251 = int32(0)
	v254 = base.AtomicRmwOr32(m, v251, int32(_a_F_GetSnapshotData_0), v251)
	if v249 <= v250 {
		goto L58
	} else {
		goto L59
	}
L35:
	;
	v371 = v137
	v375 = v2
	v376 = v2
	v380 = v2
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
	v152 = v137
	v155 = v2
	v156 = v2
	v157 = v2
	v161 = v2
	goto L38
L38:
	;
	v178 = v155 << (uint(int32(2)) % 32)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v32+v178)))
	v181 = int32(0)
	if base.B2i32(v180 == v181)|base.B2i32(v155 == v103)|base.B2i32(v181 <= v180-v130) != 0 {
		v236 = v152
		v239 = v156
		v240 = v157
		v241 = v161
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v371 = v236
	v375 = v239
	v376 = v240
	v380 = v241
	goto L31
L40:
	;
	v244 = v155 + int32(1)
	if v244 != v140 {
		v152 = v236
		v155 = v244
		v156 = v239
		v157 = v240
		v161 = v241
		goto L38
	} else {
		goto L56
	}
L41:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155+v148))))
	if v190&int32(18) != 0 {
		v236 = v152
		v239 = v156
		v240 = v157
		v241 = v161
		goto L40
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v145+v161<<(uint(int32(2))%32)))) = v180
	if v180-v152 < int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v200 = v180
	goto L45
L44:
	;
	v200 = v152
	goto L45
L45:
	;
	v202 = v161 + int32(1)
	if v157 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v236 = v200
	v239 = v156
	v240 = int32(1)
	v241 = v202
	goto L40
L47:
	;
	goto L48
L48:
	;
	v204 = int32(1)
	v207 = v149 + v155<<(uint(v204)%32)
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+1)))
	if v208 != 0 {
		v236 = v200
		v239 = v156
		v240 = v204
		v241 = v202
		goto L40
	} else {
		goto L49
	}
L49:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
	if v209 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v236 = v200
	v239 = v156
	v240 = int32(0)
	v241 = v202
	goto L40
L51:
	;
	goto L52
L52:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v178+(v29+int32(36)))))
	v215 = int32(0)
	v217 = *(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[11]))
	v221 = base.AtomicRmwOr32(m, v215, int32(_a_F_GetSnapshotData_0), v215)
	v223 = v209 << (uint(int32(2)) % 32)
	if v223 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	base.MemoryCopy(m, v224+v156<<(uint(int32(2))%32), v217+v214*int32(768)+int32(60), v223)
	goto L55
L54:
	;
	goto L55
L55:
	;
	v236 = v200
	v239 = v156 + v209
	v240 = v215
	v241 = v202
	goto L40
L56:
	;
	goto L39
L57:
	;
	v356 = *(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[0]))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v356)+24))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v357))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v330)) == int32(0) {
		goto L79
	} else {
		goto L80
	}
L58:
	;
	v330 = v137
	v334 = v2
	goto L57
L59:
	;
	goto L60
L60:
	;
	v257 = *(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[12]))
	v259 = v257
	v260 = v137
	v263 = v250
	v264 = v2
	goto L61
L61:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259+v263))))
	if v286 == int32(1) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v330 = v323
	v334 = v324
	goto L57
L63:
	;
	v290 = *(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[13]))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v290+v263<<(uint(int32(2))%32))))
	if v264 != 0 {
		v306 = v260
		goto L68
	} else {
		goto L69
	}
L64:
	;
	v322 = v259
	v323 = v260
	v324 = v264
	goto L65
L65:
	;
	v326 = v263 + int32(1)
	if v326 != v249 {
		v259 = v322
		v260 = v323
		v263 = v326
		v264 = v324
		goto L61
	} else {
		goto L78
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v246+v264<<(uint(int32(2))%32)))) = v294
	v321 = *(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[12]))
	v322 = v321
	v323 = v313
	v324 = v264 + int32(1)
	goto L65
L67:
	;
	if int32(0) <= v294-v130 {
		v330 = v309
		v334 = v264
		goto L57
	} else {
		goto L77
	}
L68:
	;
	if base.Ui32(v294) < base.Ui32(int32(3)) {
		v313 = v306
		goto L66
	} else {
		goto L76
	}
L69:
	;
	v295 = int32(3)
	if base.B2i32(base.Ui32(v260) < base.Ui32(v295))|base.B2i32(base.Ui32(v294) < base.Ui32(v295)) == int32(0) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v306 = v294
	goto L68
L71:
	;
	if v294-v260 < int32(0) {
		goto L70
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	if base.Ui32(v260) <= base.Ui32(v294) {
		v306 = v260
		goto L68
	} else {
		goto L75
	}
L74:
	;
	v309 = v260
	goto L67
L75:
	;
	goto L70
L76:
	;
	v309 = v306
	goto L67
L77:
	;
	v313 = v309
	goto L66
L78:
	;
	goto L62
L79:
	;
	v371 = v330
	v375 = v334
	v376 = base.B2i32(base.Ui32(v330) <= base.Ui32(v357))
	v380 = v2
	goto L31
L80:
	;
	goto L81
L81:
	;
	v371 = v330
	v375 = v334
	v376 = base.B2i32(v330-v357 <= int32(0))
	v380 = v2
	goto L31
L82:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[8])) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v401)+52)) = v371
	goto L84
L83:
	;
	goto L84
L84:
	;
	v409 = *(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[2]))
	F_LWLockRelease(m, v409+int32(512))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L8
	} else {
		goto L85
	}
L85:
	;
	if v371 == int32(0) {
		v429 = v399
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v448 = v109 + base.I64_extend_i32_s(v429-v125)
	v450 = *(*int64)(unsafe.Add(mBase, _c_F_GetSnapshotData[14]))
	v455 = v109 + base.I64_extend_i32_s(v445-v125)
	if base.I32_wrap_i64(v455) == int32(0) {
		goto L105
	} else {
		goto L106
	}
L87:
	;
	if v398 == int32(0) {
		v445 = v429
		goto L86
	} else {
		goto L96
	}
L88:
	;
	if v399 == int32(0) {
		v429 = v371
		goto L87
	} else {
		goto L89
	}
L89:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v399))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v371)) != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v427 = int32(base.Ui32(v371-v399) >> (uint(int32(31)) % 32))
	goto L92
L91:
	;
	v427 = base.B2i32(base.Ui32(v371) < base.Ui32(v399))
	goto L92
L92:
	;
	if v427 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v428 = v371
	goto L95
L94:
	;
	v428 = v399
	goto L95
L95:
	;
	v429 = v428
	goto L87
L96:
	;
	if v429 == int32(0) {
		v445 = v398
		goto L86
	} else {
		goto L97
	}
L97:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v429))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v398)) != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v443 = int32(base.Ui32(v398-v429) >> (uint(int32(31)) % 32))
	goto L100
L99:
	;
	v443 = base.B2i32(base.Ui32(v398) < base.Ui32(v429))
	goto L100
L100:
	;
	if v443 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v444 = v398
	goto L103
L102:
	;
	v444 = v429
	goto L103
L103:
	;
	v445 = v444
	goto L86
L104:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_GetSnapshotData[15])) = v473
	*(*int64)(unsafe.Add(mBase, _c_F_GetSnapshotData[14])) = v472
	v478 = int32(_a_F_GetSnapshotData_1)
	v480 = *(*int64)(unsafe.Add(mBase, _c_F_GetSnapshotData[16]))
	if base.Ui64(v480) < base.Ui64(v448) {
		goto L120
	} else {
		goto L121
	}
L105:
	;
	v460 = *(*int64)(unsafe.Add(mBase, _c_F_GetSnapshotData[15]))
	v472 = v450
	v473 = v460
	goto L104
L106:
	;
	goto L107
L107:
	;
	if base.Ui64(v450) < base.Ui64(v455) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v462 = v455
	goto L110
L109:
	;
	v462 = v450
	goto L110
L110:
	;
	if base.I32_wrap_i64(v450) != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v464 = v462
	goto L113
L112:
	;
	v464 = v455
	goto L113
L113:
	;
	v466 = *(*int64)(unsafe.Add(mBase, _c_F_GetSnapshotData[15]))
	if base.Ui64(v466) < base.Ui64(v455) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v468 = v455
	goto L116
L115:
	;
	v468 = v466
	goto L116
L116:
	;
	if base.I32_wrap_i64(v466) != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v470 = v468
	goto L119
L118:
	;
	v470 = v455
	goto L119
L119:
	;
	v472 = v464
	v473 = v470
	goto L104
L120:
	;
	v482 = v448
	goto L122
L121:
	;
	v482 = v480
	goto L122
L122:
	;
	if base.I32_wrap_i64(v480) != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v484 = v482
	goto L125
L124:
	;
	v484 = v448
	goto L125
L125:
	;
	if base.I32_wrap_i64(v448) != 0 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v486 = v484
	goto L128
L127:
	;
	v486 = v480
	goto L128
L128:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_GetSnapshotData[16])) = v486
	if base.Ui32(int32(3)) <= base.Ui32(v107) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v511 = v109 + base.I64_extend_i32_s(v107-v125)
	goto L131
L130:
	;
	v494 = int64(1)
	v502 = v109 + v494
	if base.Ui32(base.I32_wrap_i64(v502)) < base.Ui32(int32(3)) {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_GetSnapshotData[17])) = v511
	v513 = base.I64_extend_i32_s(v108-v125) + v109
	v514 = int32(_a_F_GetSnapshotData_2)
	v516 = *(*int64)(unsafe.Add(mBase, _c_F_GetSnapshotData[18]))
	if base.I32_wrap_i64(v516) != 0 {
		goto L138
	} else {
		goto L139
	}
L132:
	;
	v506 = v109 + (v494-v109)&int64(4294967295) + int64(2)
	goto L134
L133:
	;
	v506 = v502
	goto L134
L134:
	;
	if base.Ui64(int64(2)) < base.Ui64(v502) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v509 = v506
	goto L137
L136:
	;
	v509 = v502
	goto L137
L137:
	;
	v511 = v509
	goto L131
L138:
	;
	if base.Ui64(v513) < base.Ui64(v516) {
		goto L141
	} else {
		goto L142
	}
L139:
	;
	v522 = v513
	goto L140
L140:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_GetSnapshotData[18])) = v522
	v524 = int32(_a_F_GetSnapshotData_3)
	v526 = *(*int64)(unsafe.Add(mBase, _c_F_GetSnapshotData[19]))
	if base.I32_wrap_i64(v526) != 0 {
		goto L147
	} else {
		goto L148
	}
L141:
	;
	v519 = v516
	goto L143
L142:
	;
	v519 = v513
	goto L143
L143:
	;
	if base.I32_wrap_i64(v513) != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v521 = v519
	goto L146
L145:
	;
	v521 = v516
	goto L146
L146:
	;
	v522 = v521
	goto L140
L147:
	;
	if base.Ui64(v513) < base.Ui64(v526) {
		goto L150
	} else {
		goto L151
	}
L148:
	;
	v532 = v513
	goto L149
L149:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_GetSnapshotData[19])) = v532
	v535 = *(*int64)(unsafe.Add(mBase, _c_F_GetSnapshotData[20]))
	if base.I32_wrap_i64(v535) != 0 {
		goto L156
	} else {
		goto L157
	}
L150:
	;
	v529 = v526
	goto L152
L151:
	;
	v529 = v513
	goto L152
L152:
	;
	if base.I32_wrap_i64(v513) != 0 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v531 = v529
	goto L155
L154:
	;
	v531 = v526
	goto L155
L155:
	;
	v532 = v531
	goto L149
L156:
	;
	if base.Ui64(v513) < base.Ui64(v535) {
		goto L159
	} else {
		goto L160
	}
L157:
	;
	v541 = v513
	goto L158
L158:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_GetSnapshotData[21])) = v511
	*(*int64)(unsafe.Add(mBase, _c_F_GetSnapshotData[20])) = v541
	*(*int32)(unsafe.Add(mBase, _c_F_GetSnapshotData[9])) = v371
	*(*int64)(unsafe.Add(mBase, uint32(l0)+64)) = v70
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v376)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v375
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v380
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v371
	v555 = F_GetCurrentCommandId(m, int32(0))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L8
	} else {
		goto L165
	}
L159:
	;
	v538 = v535
	goto L161
L160:
	;
	v538 = v513
	goto L161
L161:
	;
	if base.I32_wrap_i64(v513) != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v540 = v538
	goto L164
L163:
	;
	v540 = v535
	goto L164
L164:
	;
	v541 = v540
	goto L158
L165:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+44)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v555
	v560 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)) = uint8(v560)
	return l0
L166:
	;
	F_errcode(m, int32(_a_F_GetSnapshotData_4))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L8
	} else {
		goto L167
	}
L167:
	;
	F_errmsg(m, int32(_a_F_GetSnapshotData_5), int32(0))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L8
	} else {
		goto L168
	}
L168:
	;
	F_errfinish(m, int32(_a_F_GetSnapshotData_6), int32(2156), int32(_a_F_GetSnapshotData_7))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L8
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
	F_errcode(m, int32(_a_F_GetSnapshotData_4))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L8
	} else {
		goto L171
	}
L171:
	;
	F_errmsg(m, int32(_a_F_GetSnapshotData_5), int32(0))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L8
	} else {
		goto L172
	}
L172:
	;
	F_errfinish(m, int32(_a_F_GetSnapshotData_6), int32(2170), int32(_a_F_GetSnapshotData_7))
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L8
	} else {
		goto L173
	}
L173:
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
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v258 int32
	_ = v258
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
	return v258
L9:
	;
	F_LockBufferInternal(m, v26, int32(3))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L39
	} else {
		goto L54
	}
L10:
	;
	F_visibilitymap_pin(m, l0, v29, v28)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
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
		v258 = v8
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
		v258 = v8
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
		v258 = v8
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
	F_UnlockBuffer(m, v26)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	return int32(0)
L40:
	;
	v110 = int32(0)
	v113 = base.B2i32(v30 == v110) | base.B2i32(l1 == l2)
	if v113 == v110 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	F_UnlockBuffer(m, v30)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
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
	v121 = m.ExcPending
	if v121 != 0 {
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
	v135 = int32(0)
	goto L9
L48:
	;
	F_UnlockBuffer(m, v26)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L39
	} else {
		goto L49
	}
L49:
	;
	F_visibilitymap_pin(m, l0, v25, v27)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L39
	} else {
		goto L50
	}
L50:
	;
	F_LockBufferInternal(m, v26, int32(3))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L39
	} else {
		goto L51
	}
L51:
	;
	return int32(1)
L52:
	;
	v135 = int32(0)
	goto L9
L53:
	;
	v135 = v67
	goto L9
L54:
	;
	v139 = int32(1)
	if v113 != 0 {
		v258 = v139
		goto L8
	} else {
		goto L55
	}
L55:
	;
	F_LockBufferInternal(m, v30, int32(3))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L39
	} else {
		goto L56
	}
L56:
	;
	if v135 != 0 {
		v258 = v139
		goto L8
	} else {
		goto L57
	}
L57:
	;
	goto L58
L58:
	;
	v163 = int32(0)
	if base.B2i32(int32(0) <= v26) == v163 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	v258 = v139
	goto L8
L60:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+10)))
	if v176&int32(4) != 0 {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_GetVisibilityMapPins[1]))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v167+v32<<(uint(int32(2))%32))))
	v175 = v169
	goto L60
L62:
	;
	goto L63
L63:
	;
	v171 = *(*int32)(unsafe.Add(mBase, _c_F_GetVisibilityMapPins[0]))
	v175 = v171 + v34 + int32(-8192)
	goto L60
L64:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if v179 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	v190 = v163
	goto L66
L66:
	;
	v191 = int32(0)
	if v30 < v191 {
		goto L72
	} else {
		goto L73
	}
L67:
	;
	v190 = v187 ^ int32(1)
	goto L66
L68:
	;
	v187 = int32(0)
	goto L67
L69:
	;
	goto L70
L70:
	;
	v183 = F_BufferGetBlockNumber(m, v179)
	mBase = m.M
	v185 = base.I32_div_u_s(v25, int32(_a_F_GetVisibilityMapPins_0))
	v187 = base.B2i32(v183 == v185)
	goto L67
L71:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+10)))
	if v206&int32(4) != 0 {
		goto L75
	} else {
		goto L76
	}
L72:
	;
	v195 = *(*int32)(unsafe.Add(mBase, _c_F_GetVisibilityMapPins[1]))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v195+v69<<(uint(int32(2))%32))))
	v205 = v199
	goto L71
L73:
	;
	goto L74
L74:
	;
	v201 = *(*int32)(unsafe.Add(mBase, _c_F_GetVisibilityMapPins[0]))
	v205 = v201 + v71 + int32(-8192)
	goto L71
L75:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v209 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L76:
	;
	v220 = v191
	goto L77
L77:
	;
	if (v190|v220)&int32(1) == int32(0) {
		v258 = v139
		goto L8
	} else {
		goto L82
	}
L78:
	;
	v220 = v217 ^ int32(1)
	goto L77
L79:
	;
	v217 = int32(0)
	goto L78
L80:
	;
	goto L81
L81:
	;
	v213 = F_BufferGetBlockNumber(m, v209)
	mBase = m.M
	v215 = base.I32_div_u_s(v29, int32(_a_F_GetVisibilityMapPins_0))
	v217 = base.B2i32(v213 == v215)
	goto L78
L82:
	;
	F_UnlockBuffer(m, v26)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L39
	} else {
		goto L83
	}
L83:
	;
	F_UnlockBuffer(m, v30)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L39
	} else {
		goto L84
	}
L84:
	;
	if v190 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	F_visibilitymap_pin(m, l0, v25, v27)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L39
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	if v220 != 0 {
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
	v233 = m.ExcPending
	if v233 != 0 {
		goto L39
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	F_LockBufferInternal(m, v26, int32(3))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L39
	} else {
		goto L93
	}
L92:
	;
	goto L91
L93:
	;
	F_LockBufferInternal(m, v30, int32(3))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L39
	} else {
		goto L94
	}
L94:
	;
	if v190&v220 != int32(1) {
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
func F_gbk_to_utf8(m *base.Module, l0 int32) int64 {
	var v3 int32
	_ = v3
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v3 = int32(0)
	v7 = Fn14231(m, l0, int32(37), v3, v3, v3, int32(_a_F_gbk_to_utf8_0))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_generateClonedIndexStmt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
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
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int64
	_ = v60
	var v61 int32
	_ = v61
	var v64 int64
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int64
	_ = v83
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
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
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v141 int64
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v188 int32
	_ = v188
	var v192 int64
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
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
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v298 int64
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
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
	var v334 int32
	_ = v334
	var v342 int32
	_ = v342
	var v349 int32
	_ = v349
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v528 int64
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v541 int32
	_ = v541
	var v553 int32
	_ = v553
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v587 int32
	_ = v587
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v602 int32
	_ = v602
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v617 int32
	_ = v617
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v670 int32
	_ = v670
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v685 int32
	_ = v685
	var v690 int32
	_ = v690
	var v696 int32
	_ = v696
	var v719 int32
	_ = v719
	var v728 int32
	_ = v728
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v823 int64
	_ = v823
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v835 int64
	_ = v835
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v857 int32
	_ = v857
	var v859 int32
	_ = v859
	var v867 int32
	_ = v867
	var v870 int32
	_ = v870
	var v874 int32
	_ = v874
	var v879 int32
	_ = v879
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v904 int32
	_ = v904
	v29 = m.G0
	v31 = v29 - int32(144)
	m.G0 = v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
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
	v38 = F_SearchSysCache1(m, int32(57), base.I64_extend_i32_u(v33))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L12
	} else {
		goto L13
	}
L4:
	;
	v719 = int32(*(*int16)(unsafe.Add(mBase, uint32(v45)+8)))
	if v696 < v719 {
		goto L158
	} else {
		goto L159
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L12
	} else {
		goto L153
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L12
	} else {
		goto L150
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L12
	} else {
		goto L145
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L12
	} else {
		goto L142
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L12
	} else {
		goto L139
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L12
	} else {
		goto L136
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L12
	} else {
		goto L133
	}
L12:
	;
	return int32(0)
L13:
	;
	if v38 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+196))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+22)))
	v45 = v43 + v44
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+22)))
	v50 = v48 + v49
	v51 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v50)+84)))
	v52 = F_SearchSysCache1(m, int32(2), v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
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
	v567 = m.ExcPending
	if v567 != 0 {
		goto L12
	} else {
		goto L130
	}
L17:
	;
	if v52 == int32(0) {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+22)))
	v60 = F_SysCacheGetAttrNotNull(m, int32(34), v42, int32(17))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L12
	} else {
		goto L19
	}
L19:
	;
	v64 = F_SysCacheGetAttrNotNull(m, int32(34), v42, int32(18))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	v67 = F_palloc0(m, int32(72))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = int32(204)
	v75 = F_pstrdup(m, v57+v56+int32(4))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L12
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67)+12)) = v75
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v50)+92))
	if v78 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v79 = F_get_tablespace_name(m, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L12
	} else {
		goto L26
	}
L24:
	;
	v82 = int32(0)
	goto L25
L25:
	;
	v83 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v67)+36)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v67)+16)) = v82
	*(*int64)(unsafe.Add(mBase, uint32(v67)+44)) = v83
	*(*int64)(unsafe.Add(mBase, uint32(v67)+52)) = v83
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v67)+60)) = uint8(v90)
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+13)))
	*(*uint8)(unsafe.Add(mBase, uint32(v67)+61)) = uint8(v92)
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+14)))
	*(*uint8)(unsafe.Add(mBase, uint32(v67)+62)) = uint8(v94)
	if v94 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v82 = v79
	goto L25
L27:
	;
	v104 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v67)+67)) = v104
	v106 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v67)+4)) = v106
	v109 = v103 & v104
	*(*uint8)(unsafe.Add(mBase, uint32(v67)+64)) = uint8(v109)
	if v94|v90 == v106 {
		goto L34
	} else {
		goto L35
	}
L28:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+12)))
	if v99 != int32(1) {
		v103 = int32(0)
		goto L27
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+15)))
	v103 = v102
	goto L27
L31:
	;
	goto L30
L32:
	;
	v298 = F_SysCacheGetAttr(m, int32(34), v42, int32(20), v31+int32(135))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L12
	} else {
		goto L66
	}
L33:
	;
	v264 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v67)+63)) = uint8(v264)
	goto L32
L34:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+15)))
	if v114 != int32(1) {
		goto L33
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v117 = F_get_index_constraint(m, v33)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L12
	} else {
		goto L38
	}
L37:
	;
	goto L36
L38:
	;
	if v117 != 0 {
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
	v262 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v67)+63)) = uint8(v262)
	goto L32
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v117
	goto L44
L43:
	;
	goto L44
L44:
	;
	v122 = F_SearchSysCache1(m, int32(19), base.I64_extend_i32_u(v117))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L12
	} else {
		goto L45
	}
L45:
	;
	if v122 == int32(0) {
		goto L10
	} else {
		goto L46
	}
L46:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v122)+16))
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+22)))
	v128 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v67)+63)) = uint8(v128)
	v130 = v126 + v127
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+73)))
	*(*uint8)(unsafe.Add(mBase, uint32(v67)+65)) = uint8(v131)
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+74)))
	*(*uint8)(unsafe.Add(mBase, uint32(v67)+66)) = uint8(v133)
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+15)))
	if v135 != v128 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	F_ReleaseCatCache(m, v122)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L12
	} else {
		goto L65
	}
L48:
	;
	v141 = F_SysCacheGetAttrNotNull(m, int32(19), v122, int32(27))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L12
	} else {
		goto L49
	}
L49:
	;
	v144 = F_pg_detoast_datum(m, base.I32_wrap_i64(v141))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L12
	} else {
		goto L50
	}
L50:
	;
	F_deconstruct_array_builtin(m, v144, int32(26), v31+int32(140), int32(0), v31+int32(136))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L12
	} else {
		goto L51
	}
L51:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v31)+136))
	if v154 <= int32(0) {
		goto L47
	} else {
		goto L52
	}
L52:
	;
	v158 = v67 + int32(36)
	v162 = int32(0)
	goto L53
L53:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v31)+140))
	v192 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v188+v162<<(uint(int32(3))%32)))))
	v193 = F_SearchSysCache1(m, int32(40), v192)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L12
	} else {
		goto L55
	}
L54:
	;
	goto L47
L55:
	;
	if v193 == int32(0) {
		goto L9
	} else {
		goto L56
	}
L56:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v193)+16))
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+22)))
	v199 = v197 + v198
	v202 = F_pstrdup(m, v199+int32(4))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L12
	} else {
		goto L57
	}
L57:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v199)+68))
	v205 = F_get_namespace_name(m, v204)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L12
	} else {
		goto L58
	}
L58:
	;
	v207 = F_makeString(m, v205)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L12
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v207
	v210 = F_makeString(m, v202)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L12
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v210
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v210
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v31)+128))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v214
	v220 = F_list_make2_impl(m, v31+int32(120), v31+int32(116))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L12
	} else {
		goto L61
	}
L61:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	v223 = F_lappend(m, v222, v220)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L12
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v158))) = v223
	F_ReleaseCatCache(m, v193)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L12
	} else {
		goto L63
	}
L63:
	;
	v229 = v162 + int32(1)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v31)+136))
	if v229 < v230 {
		v162 = v229
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
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+135)))
	if v300 == int32(1) {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	v330 = int32(*(*int16)(unsafe.Add(mBase, uint32(v45)+10)))
	if v330 <= int32(0) {
		v696 = v330
		goto L4
	} else {
		goto L76
	}
L68:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v312)+12))
	v326 = v312
	v327 = v317
	v328 = v319
	v329 = v325
	goto L67
L69:
	;
	v323 = int32(0)
	v326 = v323
	v327 = v321
	v328 = v322
	v329 = v323
	goto L67
L70:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v67)+20)) = int64(0)
	v321 = v67 + int32(24)
	v322 = v67 + int32(20)
	goto L69
L71:
	;
	goto L72
L72:
	;
	v310 = F_text_to_cstring(m, base.I32_wrap_i64(v298))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L12
	} else {
		goto L73
	}
L73:
	;
	v312 = F_stringToNode(m, v310)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L12
	} else {
		goto L74
	}
L74:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v67)+20)) = int64(0)
	v317 = v67 + int32(24)
	v319 = v67 + int32(20)
	if v312 != 0 {
		goto L68
	} else {
		goto L75
	}
L75:
	;
	v321 = v317
	v322 = v319
	goto L69
L76:
	;
	v334 = int32(24)
	v342 = int32(0)
	v349 = v329
	goto L77
L77:
	;
	v371 = v342 << (uint(int32(1)) % 32)
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l1)+224))
	v374 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v371+v372))))
	v376 = int32(*(*int16)(unsafe.Add(mBase, uint32(v371+(v45+int32(48))))))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)))
	v380 = F_palloc0(m, int32(40))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L12
	} else {
		goto L79
	}
L78:
	;
	v696 = v562
	goto L4
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v380))) = int32(92)
	if v376 != 0 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v431 = F_pstrdup(m, v377+v378<<(uint(int32(3))%32)+v342*int32(100)+int32(32))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L12
	} else {
		goto L93
	}
L81:
	;
	v385 = F_get_attname(m, v46, v376, int32(0))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L12
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	if v349 == int32(0) {
		goto L8
	} else {
		goto L86
	}
L84:
	;
	v387 = F_get_atttype(m, v46, v376)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L12
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v380)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v380)+4)) = v385
	v420 = v349
	v421 = v387
	goto L80
L86:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v326)+4))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v326)+12))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v349)))
	v401 = F_map_variable_attnos(m, v396, int32(1), l2, int32(0), v31+int32(140))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L12
	} else {
		goto L87
	}
L87:
	;
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+140)))
	if v403 == int32(1) {
		goto L7
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v380)+8)) = v401
	v407 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v380)+4)) = v407
	v410 = v349 + int32(4)
	if base.Ui32(v410) < base.Ui32(v395+v394<<(uint(int32(2))%32)) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v416 = v410
	goto L91
L90:
	;
	v416 = v407
	goto L91
L91:
	;
	v417 = F_exprType(m, v401)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L12
	} else {
		goto L92
	}
L92:
	;
	v420 = v416
	v421 = v417
	goto L80
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v380)+12)) = v431
	v434 = int32(0)
	v436 = v342 << (uint(int32(2)) % 32)
	v438 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v60)+v334+v436)))
	if v438 == v434 {
		v477 = v434
		goto L94
	} else {
		goto L95
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v380)+16)) = v477
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v436+(base.I32_wrap_i64(v64)+v334))))
	v485 = F_SearchSysCache1(m, int32(14), base.I64_extend_i32_u(v483))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L12
	} else {
		goto L106
	}
L95:
	;
	v441 = F_get_typcollation(m, v421)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L12
	} else {
		goto L96
	}
L96:
	;
	if v441 == v438 {
		v477 = v434
		goto L94
	} else {
		goto L97
	}
L97:
	;
	v446 = F_SearchSysCache1(m, int32(16), base.I64_extend_i32_u(v438))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L12
	} else {
		goto L98
	}
L98:
	;
	if v446 == int32(0) {
		goto L6
	} else {
		goto L99
	}
L99:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v446)+16))
	v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v450)+22)))
	v452 = v450 + v451
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v452)+68))
	v454 = F_get_namespace_name(m, v453)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L12
	} else {
		goto L100
	}
L100:
	;
	v458 = F_pstrdup(m, v452+int32(4))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L12
	} else {
		goto L101
	}
L101:
	;
	v460 = F_makeString(m, v454)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L12
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v460
	v463 = F_makeString(m, v458)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L12
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v463
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v463
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v31)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+92)) = v467
	v473 = F_list_make2_impl(m, v31+int32(92), v31+int32(88))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L12
	} else {
		goto L104
	}
L104:
	;
	F_ReleaseCatCache(m, v446)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L12
	} else {
		goto L105
	}
L105:
	;
	v477 = v473
	goto L94
L106:
	;
	if v485 == int32(0) {
		goto L5
	} else {
		goto L107
	}
L107:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v485)+16))
	v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v490)+22)))
	v492 = v490 + v491
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v492)+4))
	v494 = F_GetDefaultOpClass(m, v421, v493)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L12
	} else {
		goto L108
	}
L108:
	;
	if v494 != v483 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v492)+72))
	v498 = F_get_namespace_name(m, v497)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L12
	} else {
		goto L112
	}
L110:
	;
	v521 = int32(0)
	goto L111
L111:
	;
	F_ReleaseCatCache(m, v485)
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L12
	} else {
		goto L117
	}
L112:
	;
	v502 = F_pstrdup(m, v492+int32(8))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L12
	} else {
		goto L113
	}
L113:
	;
	v504 = F_makeString(m, v498)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L12
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v504
	v507 = F_makeString(m, v502)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L12
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v31)+72)) = v507
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v31)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+76)) = v511
	v517 = F_list_make2_impl(m, v31+int32(76), v31+int32(72))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L12
	} else {
		goto L116
	}
L116:
	;
	v521 = v517
	goto L111
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v380)+20)) = v521
	v526 = v342 + int32(1)
	v528 = F_get_attoptions(m, v33, base.I32_extend16_s(v526))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L12
	} else {
		goto L118
	}
L118:
	;
	v530 = F_untransformRelOptions(m, v528)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L12
	} else {
		goto L119
	}
L119:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v380)+28)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v380)+24)) = v530
	v535 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v535)+10)))
	if v536 != int32(1) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v380)+36)) = int32(-1)
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v328)))
	v559 = F_lappend(m, v558, v380)
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L12
	} else {
		goto L128
	}
L121:
	;
	if v374&int32(1) != 0 {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v380)+32)) = v553
	goto L120
L123:
	;
	v541 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v380)+28)) = v541
	if v374&v541 == int32(0) {
		v553 = v541
		goto L122
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	if v374&int32(2) == int32(0) {
		goto L120
	} else {
		goto L127
	}
L126:
	;
	goto L120
L127:
	;
	v553 = int32(1)
	goto L122
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v328))) = v559
	v562 = int32(*(*int16)(unsafe.Add(mBase, uint32(v45)+10)))
	if v526 < v562 {
		v342 = v526
		v349 = v420
		goto L77
	} else {
		goto L129
	}
L129:
	;
	goto L78
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v33
	F_errmsg_internal(m, int32(_a_F_generateClonedIndexStmt_0), v31)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L12
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(_a_F_generateClonedIndexStmt_1), int32(1739), int32(_a_F_generateClonedIndexStmt_2))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
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
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v50)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v581
	F_errmsg_internal(m, int32(_a_F_generateClonedIndexStmt_3), v31+int32(16))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L12
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(_a_F_generateClonedIndexStmt_1), int32(1751), int32(_a_F_generateClonedIndexStmt_2))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v117
	F_errmsg_internal(m, int32(_a_F_generateClonedIndexStmt_4), v31+int32(96))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L12
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(_a_F_generateClonedIndexStmt_1), int32(1816), int32(_a_F_generateClonedIndexStmt_2))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
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
	*(*uint32)(unsafe.Add(mBase, uint32(v31)+112)) = uint32(v192)
	F_errmsg_internal(m, int32(_a_F_generateClonedIndexStmt_5), v31+int32(112))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L12
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_generateClonedIndexStmt_1), int32(1851), int32(_a_F_generateClonedIndexStmt_2))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
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
	v630 = m.ExcPending
	if v630 != 0 {
		goto L12
	} else {
		goto L143
	}
L143:
	;
	F_errfinish(m, int32(_a_F_generateClonedIndexStmt_1), int32(1918), int32(_a_F_generateClonedIndexStmt_2))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
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
	v642 = m.ExcPending
	if v642 != 0 {
		goto L12
	} else {
		goto L146
	}
L146:
	;
	F_errmsg(m, int32(_a_F_generateClonedIndexStmt_7), int32(0))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L12
	} else {
		goto L147
	}
L147:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+32)) = v647 + int32(4)
	v654 = F_errdetail(m, int32(_a_F_generateClonedIndexStmt_8), v31+int32(32))
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L12
	} else {
		goto L148
	}
L148:
	;
	F_errfinish(m, int32(_a_F_generateClonedIndexStmt_1), int32(1934), int32(_a_F_generateClonedIndexStmt_2))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v438
	F_errmsg_internal(m, int32(_a_F_generateClonedIndexStmt_9), v31+int32(80))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L12
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(_a_F_generateClonedIndexStmt_1), int32(2212), int32(_a_F_generateClonedIndexStmt_10))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v31)+48)) = v483
	F_errmsg_internal(m, int32(_a_F_generateClonedIndexStmt_11), v31+int32(48))
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L12
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(_a_F_generateClonedIndexStmt_1), int32(2239), int32(_a_F_generateClonedIndexStmt_12))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
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
	v883 = m.ExcPending
	if v883 != 0 {
		goto L12
	} else {
		goto L188
	}
L157:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L12
	} else {
		goto L184
	}
L158:
	;
	v728 = v696
	goto L161
L159:
	;
	goto L160
L160:
	;
	v823 = F_SysCacheGetAttr(m, int32(57), v38, int32(33), v31+int32(135))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L12
	} else {
		goto L169
	}
L161:
	;
	v754 = int32(*(*int16)(unsafe.Add(mBase, uint32(v45+int32(48)+v728<<(uint(int32(1))%32)))))
	v755 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v755)))
	v758 = F_palloc0(m, int32(40))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L12
	} else {
		goto L163
	}
L162:
	;
	goto L160
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v758))) = int32(92)
	if v754 == int32(0) {
		goto L157
	} else {
		goto L164
	}
L164:
	;
	v765 = F_get_attname(m, v46, v754, int32(0))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L12
	} else {
		goto L165
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v758)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v758)+4)) = v765
	v778 = F_pstrdup(m, v755+v756<<(uint(int32(3))%32)+v728*int32(100)+int32(32))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L12
	} else {
		goto L166
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v758)+36)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v758)+12)) = v778
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v327)))
	v784 = F_lappend(m, v783, v758)
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L12
	} else {
		goto L167
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v327))) = v784
	v788 = v728 + int32(1)
	v789 = int32(*(*int16)(unsafe.Add(mBase, uint32(v45)+8)))
	if v788 < v789 {
		v728 = v788
		goto L161
	} else {
		goto L168
	}
L168:
	;
	goto L162
L169:
	;
	v825 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+135)))
	if v825 == int32(0) {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v828 = F_untransformRelOptions(m, v823)
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L12
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	v835 = F_SysCacheGetAttr(m, int32(34), v42, int32(21), v31+int32(135))
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L12
	} else {
		goto L174
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67)+28)) = v828
	goto L172
L174:
	;
	v837 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+135)))
	if v837 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v841 = F_text_to_cstring(m, base.I32_wrap_i64(v835))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L12
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	F_ReleaseCatCache(m, v38)
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L12
	} else {
		goto L182
	}
L178:
	;
	v843 = F_stringToNode(m, v841)
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L12
	} else {
		goto L179
	}
L179:
	;
	v849 = F_map_variable_attnos(m, v843, int32(1), l2, int32(0), v31+int32(140))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L12
	} else {
		goto L180
	}
L180:
	;
	v851 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+140)))
	if v851 == int32(1) {
		goto L156
	} else {
		goto L181
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67)+32)) = v849
	goto L177
L182:
	;
	F_ReleaseCatCache(m, v52)
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L12
	} else {
		goto L183
	}
L183:
	;
	m.G0 = v31 + int32(144)
	return v67
L184:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L12
	} else {
		goto L185
	}
L185:
	;
	F_errmsg(m, int32(_a_F_generateClonedIndexStmt_13), int32(0))
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		goto L12
	} else {
		goto L186
	}
L186:
	;
	F_errfinish(m, int32(_a_F_generateClonedIndexStmt_1), int32(2006), int32(_a_F_generateClonedIndexStmt_2))
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
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
	v886 = m.ExcPending
	if v886 != 0 {
		goto L12
	} else {
		goto L189
	}
L189:
	;
	F_errmsg(m, int32(_a_F_generateClonedIndexStmt_7), int32(0))
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L12
	} else {
		goto L190
	}
L190:
	;
	v891 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+64)) = v891 + int32(4)
	v898 = F_errdetail(m, int32(_a_F_generateClonedIndexStmt_8), v31-int32(-64))
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L12
	} else {
		goto L191
	}
L191:
	;
	F_errfinish(m, int32(_a_F_generateClonedIndexStmt_1), int32(2046), int32(_a_F_generateClonedIndexStmt_2))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
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
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
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
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+144))
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
	v36 = int32(0)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v39 <= v36 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v42 = v28 + int32(8)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v49 = v46 & int32(-1)
	if v49 != 0 {
		v71 = v49
		v72 = v36
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v50 = int32(1)
	if v50 == v39 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v54 = v50
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
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
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
	v613 = *(*int32)(unsafe.Add(mBase, uint32(l1)+144))
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
	v173 = int32(base.Ui32(v171) >> (uint(int32(5)) % 32))
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
	v622 = int32(base.Ui32(v620) >> (uint(int32(5)) % 32))
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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	v3 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v11 - int32(1) {
	case 0, 2:
		goto L3
	default:
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+260)) = int32(0)
	return
L2:
	;
	return
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+256))
	if v14 == int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+264))
	if v17 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+260))
	if v20 <= int32(0) {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+276))
	if v23 == int32(0) {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v26 = int32(0)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	if v28 == v26 {
		v49 = v26
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v49 != 0 {
		goto L2
	} else {
		goto L18
	}
L9:
	;
	goto L8
L10:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v32 = v31
	goto L11
L11:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if base.Ui32(int32(2)) <= base.Ui32(v36-int32(303)) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v49 = int32(1)
	goto L9
L13:
	;
	if v36 != int32(293) {
		v49 = v26
		goto L9
	} else {
		goto L16
	}
L14:
	;
	v32 = v35 + int32(72)
	goto L11
L15:
	;
	goto L12
L16:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v35)+72))
	if v43 != 0 {
		v49 = v26
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
	v51 = m.ExcPending
	if v51 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return
L20:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+260))
	if int32(0) < v52 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	F_add_paths_to_append_rel(m, l0, l1, v179)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L19
	} else {
		goto L68
	}
L22:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+276))
	v63 = v3
	v65 = v3
	goto L25
L23:
	;
	goto L24
L24:
	;
	F_mark_dummy_rel(m, l1)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L19
	} else {
		goto L67
	}
L25:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v55+v65<<(uint(int32(2))%32))))
	if v73 == int32(0) {
		v179 = v63
		goto L27
	} else {
		goto L28
	}
L26:
	;
	if v179 != 0 {
		goto L21
	} else {
		goto L66
	}
L27:
	;
	v182 = v65 + int32(1)
	if v182 != v52 {
		v63 = v179
		v65 = v182
		goto L25
	} else {
		goto L65
	}
L28:
	;
	F_generate_partitionwise_join_paths(m, l0, v73)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L19
	} else {
		goto L29
	}
L29:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v73)+44))
	if v78 == int32(0) {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_set_cheapest(m, v73)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L19
	} else {
		goto L31
	}
L31:
	;
	v83 = int32(0)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v73)+44))
	if v85 == v83 {
		v106 = v83
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if v106 != 0 {
		v179 = v63
		goto L27
	} else {
		goto L42
	}
L33:
	;
	goto L32
L34:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
	v89 = v88
	goto L35
L35:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	if base.Ui32(int32(2)) <= base.Ui32(v93-int32(303)) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v106 = int32(1)
	goto L33
L37:
	;
	if v93 != int32(293) {
		v106 = v83
		goto L33
	} else {
		goto L40
	}
L38:
	;
	v89 = v92 + int32(72)
	goto L35
L39:
	;
	goto L36
L40:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v92)+72))
	if v100 != 0 {
		v106 = v83
		goto L33
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v73)+240))
	if v107 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v177 = F_lappend(m, v63, v73)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L19
	} else {
		goto L64
	}
L44:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v110) <= base.Ui32(int32(5)) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v120 = int32(0)
	if base.B2i32(v118 == v120)|base.B2i32(v119 == v120) != 0 {
		v166 = base.B2i32(v118|v119 == v120)
		goto L51
	} else {
		goto L52
	}
L46:
	;
	if int32(1)<<(uint(v110)%32)&int32(44) != 0 {
		v117 = l1 + int32(252)
		goto L45
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v117 = l1 + int32(8)
	goto L45
L49:
	;
	goto L48
L50:
	;
	if v166 != 0 {
		goto L43
	} else {
		goto L61
	}
L51:
	;
	goto L50
L52:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v118)+4))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
	if v134 != v135 {
		v166 = int32(0)
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v137 = int32(1)
	if v134 <= v137 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v140 = v137
	goto L56
L55:
	;
	v140 = v134
	goto L56
L56:
	;
	v141 = int32(8)
	v146 = int32(0)
	goto L57
L57:
	;
	v154 = v146 << (uint(int32(2)) % 32)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v118+v141+v154)))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v119+v141+v154)))
	v159 = base.B2i32(v156 == v158)
	if v156 != v158 {
		v166 = v159
		goto L51
	} else {
		goto L59
	}
L58:
	;
	v166 = v159
	goto L51
L59:
	;
	v162 = v146 + int32(1)
	if v162 != v140 {
		v146 = v162
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v73)+240))
	F_generate_grouped_paths(m, l0, v171, v73)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L19
	} else {
		goto L62
	}
L62:
	;
	F_set_cheapest(m, v171)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L19
	} else {
		goto L63
	}
L63:
	;
	goto L43
L64:
	;
	v179 = v177
	goto L27
L65:
	;
	goto L26
L66:
	;
	goto L24
L67:
	;
	return
L68:
	;
	F_list_free(m, v179)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L19
	} else {
		goto L69
	}
L69:
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
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v36 float64
	_ = v36
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+388))
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_geqo_rand[0]))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v2+v4<<(uint(int32(2))%32))))
	v10 = v8 + int32(8)
	v13 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	v14 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
	v15 = v13 ^ v14
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = base.I64_rotl(v15, int64(37))
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v15<<(uint(int64(16))%64) ^ base.I64_rotl(v13, int64(24)) ^ v15
	v36 = F_scalbn(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v13*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
	mBase = m.M
	return v36
}
func F_getTokenTypes(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v31 int64
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
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
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	v3 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v17 = F_lookup_ts_parser_cache(m, l0)
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
	if l1 == int32(0) {
		v207 = v3
		goto L5
	} else {
		goto L6
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L50
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L47
	}
L5:
	;
	m.G0 = v15 + int32(32)
	return v207
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v23 == int32(0) {
		v207 = v3
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	if v26 == int32(0) {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v31 = F_OidFunctionCall1Coll(m, v26, int32(0), int64(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v33 <= int32(0) {
		v207 = v3
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v36 = base.I32_wrap_i64(v31)
	v41 = v3
	v42 = v33
	v43 = v3
	goto L11
L11:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49+v43<<(uint(int32(2))%32))))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v41 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v207 = v192
	goto L5
L13:
	;
	v201 = v43 + int32(1)
	if v201 < v193 {
		v41 = v192
		v42 = v193
		v43 = v201
		goto L11
	} else {
		goto L46
	}
L14:
	;
	if v36 == int32(0) {
		goto L3
	} else {
		goto L28
	}
L15:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v57 <= int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v62 = int32(0)
	goto L17
L17:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v60+v62<<(uint(int32(2))%32))))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if base.B2i32(v81 == int32(0))|base.B2i32(v81 != v84) != 0 {
		v102 = v81
		v103 = v84
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L14
L19:
	;
	if v102-v103 == int32(0) {
		v192 = v41
		v193 = v42
		goto L13
	} else {
		goto L26
	}
L20:
	;
	goto L19
L21:
	;
	v87 = v54
	v88 = v78
	goto L22
L22:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+1)))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+1)))
	if v92 == int32(0) {
		v102 = v92
		v103 = v91
		goto L20
	} else {
		goto L24
	}
L23:
	;
	v102 = v92
	v103 = v91
	goto L20
L24:
	;
	v95 = int32(1)
	if v92 == v91 {
		v87 = v87 + v95
		v88 = v88 + v95
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v108 = v62 + int32(1)
	if v57 != v108 {
		v62 = v108
		goto L17
	} else {
		goto L27
	}
L27:
	;
	goto L18
L28:
	;
	v124 = int32(0)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	if v125 == v124 {
		goto L3
	} else {
		goto L29
	}
L29:
	;
	v128 = v124
	goto L30
L30:
	;
	v142 = v36 + v128*int32(12)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143))))
	if base.B2i32(v146 == int32(0))|base.B2i32(v146 != v149) != 0 {
		v167 = v146
		v168 = v149
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v177 = F_palloc0(m, int32(8))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L43
	}
L32:
	;
	if v167-v168 != 0 {
		goto L39
	} else {
		goto L40
	}
L33:
	;
	goto L32
L34:
	;
	v152 = v54
	v153 = v143
	goto L35
L35:
	;
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+1)))
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+1)))
	if v157 == int32(0) {
		v167 = v157
		v168 = v156
		goto L33
	} else {
		goto L37
	}
L36:
	;
	v167 = v157
	v168 = v156
	goto L33
L37:
	;
	v160 = int32(1)
	if v157 == v156 {
		v152 = v152 + v160
		v153 = v153 + v160
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v171 = v128 + int32(1)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v36+v171*int32(12))))
	if v175 != 0 {
		v128 = v171
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
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	*(*int32)(unsafe.Add(mBase, uint32(v177))) = v179
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v182 = F_pstrdup(m, v181)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v177)+4)) = v182
	v185 = F_lappend(m, v41, v177)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v192 = v185
	v193 = v187
	goto L13
L46:
	;
	goto L12
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l0
	F_errmsg_internal(m, int32(_a_F_getTokenTypes_0), v15)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_getTokenTypes_1), int32(1243), int32(_a_F_getTokenTypes_2))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
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
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v251
	F_errmsg(m, int32(_a_F_getTokenTypes_3), v15+int32(16))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_getTokenTypes_1), int32(1278), int32(_a_F_getTokenTypes_2))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
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
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)+116))
	if v19 == int32(0) {
		v178 = v7
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)+76))
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
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+105)))
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
	F_ScanKeyEntryInitialize(m, v14, int32(129), v124, v127, v127, v127, v127, int64(0))
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_get_attavgwidth[0]))
	if v5 != 0 {
		v6 = m.T0[v5].(func(*base.Module, int32, int32) int32)(m, l0, l1)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			if int32(0) < v6 {
				v32 = v6
				return v32
			} else {
				v17 = F_SearchSysCache3(m, int32(65), base.I64_extend_i32_u(l0), base.I64_extend_i32_s(l1), int64(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					if v17 != 0 {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
						v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+22)))
						v22 = *(*int32)(unsafe.Add(mBase, uint32(v19+v20)+12))
						F_ReleaseCatCache(m, v17)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return int32(0)
						} else {
							if int32(0) < v22 {
								v32 = v22
							} else {
								v32 = int32(0)
							}
							return v32
						}
					} else {
						v32 = int32(0)
						return v32
					}
				}
			}
		}
	} else {
		v17 = F_SearchSysCache3(m, int32(65), base.I64_extend_i32_u(l0), base.I64_extend_i32_s(l1), int64(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			if v17 != 0 {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+22)))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v19+v20)+12))
				F_ReleaseCatCache(m, v17)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					if int32(0) < v22 {
						v32 = v22
					} else {
						v32 = int32(0)
					}
					return v32
				}
			} else {
				v32 = int32(0)
				return v32
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
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
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
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
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
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 float64
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 float64
	_ = v285
	var v286 int32
	_ = v286
	var v287 float64
	_ = v287
	var v296 float64
	_ = v296
	var v300 float64
	_ = v300
	var v301 float64
	_ = v301
	var v303 float64
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
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
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	if v16 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v109 = F_bms_union(m, v108, l2)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L24
	} else {
		goto L25
	}
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v19 <= int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v26 = v4
	goto L7
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33+v26<<(uint(int32(2))%32))))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v39 = int32(0)
	if base.B2i32(v38 == v39)|base.B2i32(l2 == v39) != 0 {
		v85 = base.B2i32(v38|l2 == v39)
		goto L10
	} else {
		goto L11
	}
L8:
	;
	return v37
L9:
	;
	if v85 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L10:
	;
	goto L9
L11:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v53 != v54 {
		v85 = int32(0)
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v56 = int32(1)
	if v53 <= v56 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v59 = v56
	goto L15
L14:
	;
	v59 = v53
	goto L15
L15:
	;
	v60 = int32(8)
	v65 = int32(0)
	goto L16
L16:
	;
	v73 = v65 << (uint(int32(2)) % 32)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v38+v60+v73)))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l2+v60+v73)))
	v78 = base.B2i32(v75 == v77)
	if v75 != v77 {
		v85 = v78
		goto L10
	} else {
		goto L18
	}
L17:
	;
	v85 = v78
	goto L10
L18:
	;
	v81 = v65 + int32(1)
	if v81 != v59 {
		v65 = v81
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v93 = v26 + int32(1)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v93 < v94 {
		v26 = v93
		goto L7
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	goto L8
L23:
	;
	goto L4
L24:
	;
	return int32(0)
L25:
	;
	v113 = int32(0)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+228))
	if v114 == v113 {
		v226 = v113
		v229 = v4
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v233 = F_generate_join_implied_equalities(m, l0, v109, l2, l1, int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L24
	} else {
		goto L62
	}
L27:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	if v117 <= int32(0) {
		v226 = v113
		v229 = v4
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v125 = int32(0)
	v126 = v113
	v129 = v4
	goto L29
L29:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v132+v125<<(uint(int32(2))%32))))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v138 = int32(0)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v136)+28))
	v140 = F_bms_is_subset(m, v139, v109)
	mBase = m.M
	if v140 == v138 {
		v151 = v138
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v226 = v215
	v229 = v216
	goto L26
L31:
	;
	v218 = v125 + int32(1)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	if v218 < v219 {
		v125 = v218
		v126 = v215
		v129 = v216
		goto L29
	} else {
		goto L60
	}
L32:
	;
	if v151 == int32(0) {
		v215 = v126
		v216 = v129
		goto L31
	} else {
		goto L36
	}
L33:
	;
	goto L32
L34:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v136)+28))
	v144 = F_bms_overlap(m, v137, v143)
	mBase = m.M
	if v144 == int32(0) {
		v151 = v138
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v136)+40))
	v148 = F_bms_overlap(m, v137, v147)
	mBase = m.M
	v151 = v148 ^ int32(1)
	goto L33
L36:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+11)))
	if v154 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v210 = F_lappend(m, v129, v136)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L24
	} else {
		goto L58
	}
L38:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+12)))
	if v157 != int32(1) {
		goto L37
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v136)+36))
	v161 = int32(0)
	if base.B2i32(v160 == v161)|base.B2i32(v109 == v161) != 0 {
		v206 = v161
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L40
L42:
	;
	if v206 != 0 {
		v215 = v126
		v216 = v129
		goto L31
	} else {
		goto L55
	}
L43:
	;
	goto L42
L44:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	if v171 < v172 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v174 = v171
	goto L47
L46:
	;
	v174 = v172
	goto L47
L47:
	;
	if v174 <= int32(1) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v177 = int32(1)
	goto L50
L49:
	;
	v177 = v174
	goto L50
L50:
	;
	v178 = int32(8)
	v183 = int32(0)
	goto L51
L51:
	;
	v190 = v183 << (uint(int32(2)) % 32)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v109+v178+v190)))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v160+v178+v190)))
	v195 = v192 & v194
	v197 = base.B2i32(v195 != int32(0))
	if v195 != 0 {
		v206 = v197
		goto L43
	} else {
		goto L53
	}
L52:
	;
	v206 = v197
	goto L43
L53:
	;
	v199 = v183 + int32(1)
	if v199 != v177 {
		v183 = v199
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v136)+56))
	v208 = F_bms_is_member(m, v207, v126)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L24
	} else {
		goto L56
	}
L56:
	;
	if v208 != 0 {
		v215 = v126
		v216 = v129
		goto L31
	} else {
		goto L57
	}
L57:
	;
	goto L37
L58:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v136)+56))
	v213 = F_bms_add_member(m, v126, v212)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L24
	} else {
		goto L59
	}
L59:
	;
	v215 = v213
	v216 = v210
	goto L31
L60:
	;
	goto L30
L61:
	;
	v275 = F_list_concat(m, v229, v233)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L24
	} else {
		goto L69
	}
L62:
	;
	if v233 == int32(0) {
		v269 = v226
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v237 = int32(0)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v233)+4))
	if v238 <= v237 {
		v269 = v226
		goto L61
	} else {
		goto L64
	}
L64:
	;
	v245 = v237
	v246 = v226
	goto L65
L65:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v233)+12))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v252+v245<<(uint(int32(2))%32))))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)+56))
	v258 = F_bms_add_member(m, v246, v257)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L24
	} else {
		goto L67
	}
L66:
	;
	v269 = v258
	goto L61
L67:
	;
	v261 = v245 + int32(1)
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v233)+4))
	if v261 < v262 {
		v245 = v261
		v246 = v258
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	v278 = F_list_concat_copy(m, v275, v277)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L24
	} else {
		goto L70
	}
L70:
	;
	v281 = *(*float64)(unsafe.Add(mBase, uint32(l1)+128))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v283 = int32(0)
	v285 = F_clauselist_selectivity(m, l0, v278, v282, v283, v283)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L24
	} else {
		goto L72
	}
L71:
	;
	v301 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	if base.F64_gt(v300, v301) != 0 {
		goto L75
	} else {
		goto L76
	}
L72:
	;
	v287 = base.F64_mul(v281, v285)
	if base.F64_gt(v287, float64(1e+100))|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v287)&int64(9223372036854775807))) != 0 {
		v300 = float64(1e+100)
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v296 = float64(1)
	if base.F64_le(v287, v296) != 0 {
		v300 = v296
		goto L71
	} else {
		goto L74
	}
L74:
	;
	v300 = base.F64_nearest(v287)
	goto L71
L75:
	;
	v303 = v301
	goto L77
L76:
	;
	v303 = v300
	goto L77
L77:
	;
	v305 = F_palloc0(m, int32(24))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L24
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v305)+20)) = v269
	*(*int32)(unsafe.Add(mBase, uint32(v305)+16)) = v275
	*(*float64)(unsafe.Add(mBase, uint32(v305)+8)) = v303
	*(*int32)(unsafe.Add(mBase, uint32(v305)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v305))) = int32(281)
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v314 = F_lappend(m, v313, v305)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L24
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v314
	return v305
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
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
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
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
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
						v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
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
		v30 = int32(127)
		v35 = int32(base.Ui32(v24)>>(uint(v23)%32)) & v28
		v42 = int32(base.Ui32(v26&v21) >> (uint(int32(16)) % 32))
		v47 = int32(base.Ui32(v24) >> (uint(v25) % 32))
		v51 = int32(_a_F_get_decomposed_size_3)
		v52 = base.I32_rem_u_s(((v29*v30+v35)*v30+v42)*v30+v47+int32(260144641), v51)
		v55 = int32(*(*int16)(unsafe.Add(mBase, uint32(v52<<(uint(v20)%32))+uint32(_c_F_get_decomposed_size[0]))))
		v56 = int32(257)
		v66 = base.I32_rem_u_s(((v29*v56+v35)*v56+v42)*v56+v47, v51)
		v69 = int32(*(*int16)(unsafe.Add(mBase, uint32(v66<<(uint(v20)%32))+uint32(_c_F_get_decomposed_size[0]))))
		v70 = v55 + v69
		if base.Ui32(int32(_a_F_get_decomposed_size_4)) < base.Ui32(v70) {
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
						v94 = int32(_a_F_get_decomposed_size_5)
						*(*int32)(unsafe.Add(mBase, _c_F_get_decomposed_size[4])) = v91
						v102 = int32(1)
						v103 = v94
					} else {
						v102 = v81
						v103 = v91<<(uint(int32(2))%32) + int32(_a_F_get_decomposed_size_6)
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
							v143 = int32(127)
							v148 = int32(base.Ui32(v137)>>(uint(v136)%32)) & v141
							v155 = int32(base.Ui32(v139&v134) >> (uint(int32(16)) % 32))
							v160 = int32(base.Ui32(v137) >> (uint(v138) % 32))
							v164 = int32(_a_F_get_decomposed_size_3)
							v165 = base.I32_rem_u_s(((v142*v143+v148)*v143+v155)*v143+v160+int32(260144641), v164)
							v168 = int32(*(*int16)(unsafe.Add(mBase, uint32(v165<<(uint(v133)%32))+uint32(_c_F_get_decomposed_size[0]))))
							v169 = int32(257)
							v179 = base.I32_rem_u_s(((v142*v169+v148)*v169+v155)*v169+v160, v164)
							v182 = int32(*(*int16)(unsafe.Add(mBase, uint32(v179<<(uint(v133)%32))+uint32(_c_F_get_decomposed_size[0]))))
							v183 = v168 + v182
							if base.Ui32(int32(_a_F_get_decomposed_size_4)) < base.Ui32(v183) {
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
											v207 = int32(_a_F_get_decomposed_size_5)
											*(*int32)(unsafe.Add(mBase, _c_F_get_decomposed_size[4])) = v204
											v215 = int32(1)
											v216 = v207
										} else {
											v215 = v194
											v216 = v204<<(uint(int32(2))%32) + int32(_a_F_get_decomposed_size_6)
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
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
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
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	v6 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(l0))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 == int32(0) {
			return int32(0)
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+22)))
			v16 = v14 + v15
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+92))
			if v17 != 0 {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)+88))
				if v18 == int32(_a_F_get_element_type_0) {
					v22 = v17
				} else {
					v22 = int32(0)
				}
			} else {
				v22 = int32(0)
			}
			F_ReleaseCatCache(m, v6)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				return v22
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
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v5 = F_SearchSysCache1(m, int32(40), base.I64_extend_i32_u(l0))
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
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v13+v14)+104))
			F_ReleaseCatCache(m, v5)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				return v16
			}
		}
	}
}
func F_get_relname_relid(m *base.Module, l0 int32, l1 int32) int32 {
	var v6 int64
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v6 = int64(0)
	v8 = F_GetSysCacheOid(m, int32(56), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1), v6, v6)
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
func F_get_required_extension(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int64
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
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
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v16 = int64(0)
	v19 = F_GetSysCacheOid(m, int32(27), base.I64_extend_i32_u(l0), v16, v16, v16)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L2
	} else {
		goto L41
	}
L2:
	;
	return int32(0)
L3:
	;
	if v19 == int32(0) {
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
	v141 = v19
	goto L6
L6:
	;
	m.G0 = v12 + int32(48)
	return v141
L7:
	;
	F_check_valid_extension_name(m, l0)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
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
	v112 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L2
	} else {
		goto L32
	}
L10:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v31 <= int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v34 = int32(0)
	if v34 < v31 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v38 = v31
	goto L14
L13:
	;
	v38 = v34
	goto L14
L14:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v47 = v34
	goto L15
L15:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v39+v47<<(uint(int32(2))%32))))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v55 == int32(0))|base.B2i32(v55 != v58) != 0 {
		v76 = v55
		v77 = v58
		goto L18
	} else {
		goto L19
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L2
	} else {
		goto L28
	}
L17:
	;
	if v76-v77 != 0 {
		goto L24
	} else {
		goto L25
	}
L18:
	;
	goto L17
L19:
	;
	v61 = v52
	v62 = l0
	goto L20
L20:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+1)))
	if v66 == int32(0) {
		v76 = v66
		v77 = v65
		goto L18
	} else {
		goto L22
	}
L21:
	;
	v76 = v66
	v77 = v65
	goto L18
L22:
	;
	v69 = int32(1)
	if v66 == v65 {
		v61 = v61 + v69
		v62 = v62 + v69
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v80 = v47 + int32(1)
	if v38 != v80 {
		v47 = v80
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
	v88 = m.ExcPending
	if v88 != 0 {
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
	v95 = m.ExcPending
	if v95 != 0 {
		goto L2
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_get_required_extension_1), int32(2109), int32(_a_F_get_required_extension_2))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
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
	if v112 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
	F_errmsg(m, int32(_a_F_get_required_extension_3), v12)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L2
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v127 = F_list_copy(m, l4)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L2
	} else {
		goto L38
	}
L36:
	;
	F_errfinish(m, int32(_a_F_get_required_extension_1), int32(2114), int32(_a_F_get_required_extension_2))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v129 = F_lappend(m, v127, l1)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	F_CreateExtensionInternal(m, v12+int32(36), l0, l2, int32(0), int32(1), v129, l5)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L2
	} else {
		goto L40
	}
L40:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	v141 = v133
	goto L6
L41:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = l0
	F_errmsg(m, int32(_a_F_get_required_extension_4), v12+int32(32))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
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
	v163 = m.ExcPending
	if v163 != 0 {
		goto L2
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	F_errfinish(m, int32(_a_F_get_required_extension_1), int32(2139), int32(_a_F_get_required_extension_2))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v49 int32
	_ = v49
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
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
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
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
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v183 int32
	_ = v183
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
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_get_steps_using_prefix_recurse[0]))
	if v21 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v26 = int32(2)
	v27 = (l7 - v24) >> (uint(v26) % 32)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v24+v30<<(uint(v26)%32)-int32(4))))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	if v29 < v37 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L5
L7:
	;
	return v183
L8:
	;
	if v30 <= v27 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v114 = int32(0)
	if v30 <= v27 {
		v183 = v114
		goto L7
	} else {
		goto L31
	}
L11:
	;
	return int32(0)
L12:
	;
	goto L13
L13:
	;
	v49 = v27
	goto L15
L14:
	;
	v77 = int32(0)
	v81 = v27
	goto L19
L15:
	;
	v59 = v24 + v49<<(uint(int32(2))%32)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v29 < v61 {
		v68 = v59
		goto L14
	} else {
		goto L17
	}
L16:
	;
	v68 = int32(0)
	goto L14
L17:
	;
	v64 = v49 + int32(1)
	if v64 != v30 {
		v49 = v64
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v85+v81<<(uint(int32(2))%32))))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	if v90 != v29 {
		v183 = v77
		goto L7
	} else {
		goto L21
	}
L20:
	;
	v183 = v104
	goto L7
L21:
	;
	v92 = F_list_copy(m, l8)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v95 = F_lappend(m, v92, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v97 = F_list_copy(m, l9)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	v100 = F_lappend_oid(m, v97, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v102 = F_get_steps_using_prefix_recurse(m, l0, l1, l2, l3, l4, l5, l6, v68, v95, v100)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v104 = F_list_concat(m, v77, v102)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_list_free(m, v95)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_list_free(m, v100)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v111 = v81 + int32(1)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v111 < v112 {
		v77 = v104
		v81 = v111
		goto L19
	} else {
		goto L30
	}
L30:
	;
	goto L20
L31:
	;
	if l2 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v117 = int32(0)
	goto L34
L33:
	;
	v117 = l1
	goto L34
L34:
	;
	v125 = v114
	v129 = v27
	goto L35
L35:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v133+v129<<(uint(int32(2))%32))))
	v138 = F_list_copy(m, l8)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L37
	}
L36:
	;
	v183 = v170
	goto L7
L37:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
	v141 = F_lappend(m, v138, v140)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v143 = F_lappend(m, v141, l3)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v145 = F_list_copy(m, l9)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
	v148 = F_lappend_oid(m, v145, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v150 = F_lappend_oid(m, v148, l4)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v153 = F_palloc0(m, int32(24))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153))) = int32(381)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v157 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v153)+20)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v153)+16)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v153)+12)) = v143
	*(*uint16)(unsafe.Add(mBase, uint32(v153)+8)) = uint16(v117)
	*(*int32)(unsafe.Add(mBase, uint32(v153)+4)) = v157
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v167 = F_lappend(m, v166, v153)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v167
	v170 = F_lappend(m, v125, v153)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v173 = v129 + int32(1)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v173 < v174 {
		v125 = v170
		v129 = v173
		goto L35
	} else {
		goto L46
	}
L46:
	;
	goto L36
}
func F_get_typavgwidth(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	v7 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(l0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		if v7 != 0 {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
			v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)))
			v14 = int32(*(*int16)(unsafe.Add(mBase, uint32(v11+v12)+76)))
			F_ReleaseCatCache(m, v7)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				if int32(0) < v14 {
					v83 = v14
					return v83
				} else {
					v21 = int32(-1)
					if l1 < int32(0) {
						v61 = v21
						v63 = v61
					} else {
						if base.Ui32(int32(2)) <= base.Ui32(l0-int32(1042)) {
							switch l0 - int32(1560) {
							case 0, 2:
								v57 = int32(8)
								v58 = base.I32_div_s(l1+int32(7), v57)
								v61 = v58 + v57
								v63 = v61
							case 1:
								v61 = v21
								v63 = v61
							default:
								if l0 != int32(1700) {
									v61 = v21
									v63 = v61
								} else {
									v40 = int32(4)
									if l1 < v40 {
										v54 = int32(-1)
									} else {
										v54 = int32(base.Ui32(int32(base.Ui32(l1-v40)>>(uint(int32(16))%32))+int32(6))>>(uint(int32(1))%32))&int32(_a_F_get_typavgwidth_0) + int32(8)
									}
									v63 = v54
								}
							}
						} else {
							v30 = F_GetDatabaseEncoding(m)
							mBase = m.M
							v31 = F_pg_encoding_max_length(m, v30)
							mBase = m.M
							v32 = int32(4)
							v63 = v31*(l1-v32) + v32
						}
					}
					if v63 <= int32(0) {
						return int32(32)
					} else {
						if base.B2i32(l0 == int32(1042))|base.B2i32(base.Ui32(v63) < base.Ui32(int32(33))) != 0 {
							v83 = v63
							return v83
						} else {
							if base.Ui32(int32(999)) < base.Ui32(v63) {
								return int32(516)
							} else {
								v77 = int32(32)
								v83 = int32(base.Ui32(v63-v77)>>(uint(int32(1))%32)) + v77
								return v83
							}
						}
					}
				}
			}
		} else {
			v21 = int32(-1)
			if l1 < int32(0) {
				v61 = v21
				v63 = v61
			} else {
				if base.Ui32(int32(2)) <= base.Ui32(l0-int32(1042)) {
					switch l0 - int32(1560) {
					case 0, 2:
						v57 = int32(8)
						v58 = base.I32_div_s(l1+int32(7), v57)
						v61 = v58 + v57
						v63 = v61
					case 1:
						v61 = v21
						v63 = v61
					default:
						if l0 != int32(1700) {
							v61 = v21
							v63 = v61
						} else {
							v40 = int32(4)
							if l1 < v40 {
								v54 = int32(-1)
							} else {
								v54 = int32(base.Ui32(int32(base.Ui32(l1-v40)>>(uint(int32(16))%32))+int32(6))>>(uint(int32(1))%32))&int32(_a_F_get_typavgwidth_0) + int32(8)
							}
							v63 = v54
						}
					}
				} else {
					v30 = F_GetDatabaseEncoding(m)
					mBase = m.M
					v31 = F_pg_encoding_max_length(m, v30)
					mBase = m.M
					v32 = int32(4)
					v63 = v31*(l1-v32) + v32
				}
			}
			if v63 <= int32(0) {
				return int32(32)
			} else {
				if base.B2i32(l0 == int32(1042))|base.B2i32(base.Ui32(v63) < base.Ui32(int32(33))) != 0 {
					v83 = v63
					return v83
				} else {
					if base.Ui32(int32(999)) < base.Ui32(v63) {
						return int32(516)
					} else {
						v77 = int32(32)
						v83 = int32(base.Ui32(v63-v77)>>(uint(int32(1))%32)) + v77
						return v83
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
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
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
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(279)
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
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
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
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
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
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
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
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
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
	*(*int32)(unsafe.Add(mBase, uint32(v297))) = int32(279)
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
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
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
	var v131 int32
	_ = v131
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
	var v147 int32
	_ = v147
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
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
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v226)+72))
	v284 = F_quote_identifier(m, v283)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L13
	} else {
		goto L88
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L13
	} else {
		goto L84
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
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v126 == int32(1) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	F_appendStringInfoString(m, v17, int32(_a_F_get_windowfunc_expr_helper_4))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L13
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	F_appendStringInfoString(m, v17, int32(_a_F_get_windowfunc_expr_helper_5))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L13
	} else {
		goto L48
	}
L47:
	;
	goto L46
L48:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v135 != 0 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_get_rule_windowspec(m, v157, v264, l1)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L13
	} else {
		goto L83
	}
L50:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+4))
	if v136 <= int32(0) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v198 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L53:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L13
	} else {
		goto L64
	}
L54:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v135)+12))
	v147 = int32(0)
	goto L55
L55:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v140+v147<<(uint(int32(2))%32))))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+48))
	if v139 != v158 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	if v163 == int32(0) {
		goto L49
	} else {
		goto L61
	}
L57:
	;
	v161 = v147 + int32(1)
	if v136 != v161 {
		v147 = v161
		goto L55
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	goto L56
L60:
	;
	goto L53
L61:
	;
	v166 = F_quote_identifier(m, v163)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L13
	} else {
		goto L62
	}
L62:
	;
	F_appendStringInfoString(m, v17, v166)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L13
	} else {
		goto L63
	}
L63:
	;
	goto L1
L64:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v186
	F_errmsg_internal(m, int32(_a_F_get_windowfunc_expr_helper_6), v15+int32(16))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L13
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_get_windowfunc_expr_helper_7), int32(_a_F_get_windowfunc_expr_helper_8), int32(_a_F_get_windowfunc_expr_helper_9))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L13
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L13
	} else {
		goto L80
	}
L68:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v198)+4))
	if v201 <= int32(0) {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v204 = int32(0)
	if v204 < v201 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v208 = v201
	goto L72
L71:
	;
	v208 = v204
	goto L72
L72:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v198)+12))
	v215 = v204
	goto L73
L73:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v209+v215<<(uint(int32(2))%32))))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)+40))
	if v226 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	goto L67
L75:
	;
	v236 = v215 + int32(1)
	if v236 != v208 {
		v215 = v236
		goto L73
	} else {
		goto L79
	}
L76:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v226)))
	if v229 != int32(370) {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v226)+76))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v232 == v233 {
		goto L2
	} else {
		goto L78
	}
L78:
	;
	goto L75
L79:
	;
	goto L74
L80:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v254
	F_errmsg_internal(m, int32(_a_F_get_windowfunc_expr_helper_6), v15)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L13
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_get_windowfunc_expr_helper_7), int32(_a_F_get_windowfunc_expr_helper_10), int32(_a_F_get_windowfunc_expr_helper_9))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
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
	goto L1
L84:
	;
	F_errcode(m, int32(50856197))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L13
	} else {
		goto L85
	}
L85:
	;
	F_errmsg(m, int32(_a_F_get_windowfunc_expr_helper_11), int32(0))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L13
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(_a_F_get_windowfunc_expr_helper_7), int32(_a_F_get_windowfunc_expr_helper_12), int32(_a_F_get_windowfunc_expr_helper_9))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L13
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
	F_appendStringInfoString(m, v17, v284)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L13
	} else {
		goto L89
	}
L89:
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
func F_getdatabaseencoding(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_getdatabaseencoding[0]))
	v6 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v5))))
	v7 = F_DirectFunctionCall1Coll(m, int32(534), int32(0), v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
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
				F_errfinish(m, int32(_a_F_geterrcode_1), int32(1779), int32(_a_F_geterrcode_2))
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
				F_errfinish(m, int32(_a_F_geterrposition_1), int32(1796), int32(_a_F_geterrposition_2))
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
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
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v107 int32
	_ = v107
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
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
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
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v519 int32
	_ = v519
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v558 int32
	_ = v558
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v596 int32
	_ = v596
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v692 int32
	_ = v692
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v766 int32
	_ = v766
	var v770 int32
	_ = v770
	var v775 int32
	_ = v775
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v792 int32
	_ = v792
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	v7 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(160)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v30 = int32(1)
	v31 = v22
	v33 = v7
	v37 = v7
	v39 = v7
	v40 = v7
	goto L1
L1:
	;
	switch v30 - int32(1) {
	case 0:
		goto L29
	case 1:
		goto L27
	case 2:
		goto L28
	case 3:
		goto L26
	case 4:
		v77 = v31
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
	v828 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v829 = F_pg_mblen_cstr(m, v828)
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L40
	} else {
		goto L319
	}
L4:
	;
	v818 = int32(3)
	v819 = v31
	v821 = v33
	v824 = v37
	v826 = v816
	v827 = v40
	goto L3
L5:
	;
	v816 = int32(2)
	goto L4
L6:
	;
	m.G0 = v20 + int32(160)
	return v792
L7:
	;
	if base.B2i32(v87 == int32(32))|base.B2i32(base.Ui32(v87-int32(9)) < base.Ui32(int32(5))) != 0 {
		v818 = int32(1)
		v819 = v31
		v821 = v33
		v824 = v37
		v826 = v39
		v827 = v40
		goto L3
	} else {
		goto L314
	}
L8:
	;
	v752 = int32(0)
	v753 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v754 = F_errsave_start(m, v753)
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L40
	} else {
		goto L306
	}
L9:
	;
	if base.B2i32(v91&int32(1) == int32(0))|base.B2i32(v87 != int32(34)) != 0 {
		goto L7
	} else {
		goto L305
	}
L10:
	;
	if v87 == int32(124) {
		goto L8
	} else {
		goto L304
	}
L11:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v163 == int32(58) {
		goto L266
	} else {
		goto L267
	}
L12:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v612 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v613 = v31 - v612
	v614 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v611 <= v613+v614 {
		goto L238
	} else {
		goto L239
	}
L13:
	;
	if base.B2i32(v159 == int32(0))|base.B2i32(v163 != int32(34)) != 0 {
		goto L11
	} else {
		goto L237
	}
L14:
	;
	if v163 == int32(124) {
		goto L12
	} else {
		goto L236
	}
L15:
	;
	v574 = m.G0
	v576 = v574 - int32(16)
	m.G0 = v576
	v578 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v579 = F_errsave_start(m, v578)
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L40
	} else {
		goto L226
	}
L16:
	;
	if int32(1)<<(uint(v171)%32)&int32(134218145) == int32(0) {
		goto L14
	} else {
		goto L225
	}
L17:
	;
	if int32(1)<<(uint(v111)%32)&int32(134218145) == int32(0) {
		goto L10
	} else {
		goto L224
	}
L18:
	;
	v816 = int32(4)
	goto L4
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L40
	} else {
		goto L221
	}
L20:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v452))))
	switch v453 {
	case 0, 9, 10, 11, 12, 13, 32:
		goto L188
	default:
		goto L187
	case 42, 65, 97:
		goto L192
	case 44:
		v818 = int32(6)
		v819 = v31
		v821 = v33
		v824 = v37
		v826 = v39
		v827 = v40
		goto L3
	case 66, 98:
		goto L191
	case 67, 99:
		goto L190
	case 68, 100:
		goto L189
	}
L21:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
	if base.Ui32((v311-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
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
	v254 = int32(0)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v256 = F_errsave_start(m, v255)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L40
	} else {
		goto L108
	}
L25:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v233 = v31 - v47
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v232 <= v233+v234 {
		goto L100
	} else {
		goto L101
	}
L26:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+22)))
	if v176 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L27:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+22)))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
	if base.B2i32(v159 == int32(0))&base.B2i32(v163 == int32(92)) != 0 {
		goto L5
	} else {
		goto L70
	}
L28:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
	if v115 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L29:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	if v87 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L30:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if v82 != int32(58) {
		goto L22
	} else {
		goto L44
	}
L31:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+22)))
	if v43 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v55 = v31 - v52
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v54 <= v55+v56 {
		goto L37
	} else {
		goto L38
	}
L33:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v52 = v46
	goto L32
L34:
	;
	goto L35
L35:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	if v49 == int32(39) {
		goto L25
	} else {
		goto L36
	}
L36:
	;
	v52 = v47
	goto L32
L37:
	;
	v60 = v54 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v60
	v62 = F_repalloc(m, v52, v60)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v69 = v31
	goto L39
L39:
	;
	v71 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v69))) = uint8(v71)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v69 == v73 {
		goto L24
	} else {
		goto L42
	}
L40:
	;
	return int32(0)
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v62
	v69 = v62 + v55
	goto L39
L42:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v75 != 0 {
		goto L23
	} else {
		goto L43
	}
L43:
	;
	v77 = v69
	goto L30
L44:
	;
	v818 = int32(6)
	v819 = v77
	v821 = v33
	v824 = v37
	v826 = v39
	v827 = v40
	goto L3
L45:
	;
	v792 = int32(0)
	goto L6
L46:
	;
	goto L47
L47:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+22)))
	if v91&int32(1)|base.B2i32(v87 != int32(39)) == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v818 = int32(4)
	v819 = v31
	v821 = v33
	v824 = v37
	v826 = v39
	v827 = v40
	goto L3
L49:
	;
	goto L50
L50:
	;
	if base.B2i32(v91&int32(1) == int32(0))&base.B2i32(v87 == int32(92)) != 0 {
		goto L5
	} else {
		goto L51
	}
L51:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v107 != int32(1) {
		goto L9
	} else {
		goto L52
	}
L52:
	;
	v111 = v87 - int32(33)
	if base.Ui32(v111) <= base.Ui32(int32(27)) {
		goto L17
	} else {
		goto L53
	}
L53:
	;
	goto L10
L54:
	;
	v118 = int32(0)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v120 = F_errsave_start(m, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L40
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v141 = v31 - v140
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v139 <= v141+v142 {
		goto L62
	} else {
		goto L63
	}
L57:
	;
	if v120 == int32(0) {
		v792 = v118
		goto L6
	} else {
		goto L58
	}
L58:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L40
	} else {
		goto L59
	}
L59:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v127
	F_errmsg(m, int32(_a_F_gettoken_tsvector_0), v20+int32(32))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L40
	} else {
		goto L60
	}
L60:
	;
	F_errsave_finish(m, v119, int32(_a_F_gettoken_tsvector_1), int32(221), int32(_a_F_gettoken_tsvector_2))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L40
	} else {
		goto L61
	}
L61:
	;
	v792 = v118
	goto L6
L62:
	;
	v146 = v139 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v146
	v148 = F_repalloc(m, v140, v146)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L40
	} else {
		goto L65
	}
L63:
	;
	v153 = v114
	v154 = v31
	goto L64
L64:
	;
	v155 = F_pg_mblen_cstr(m, v153)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L40
	} else {
		goto L66
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v148
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v153 = v152
	v154 = v148 + v141
	goto L64
L66:
	;
	if v155 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	base.MemoryCopy(m, v154, v153, v155)
	goto L69
L68:
	;
	goto L69
L69:
	;
	v818 = v39
	v819 = v154 + v155
	v821 = v33
	v824 = v37
	v826 = v39
	v827 = v40
	goto L3
L70:
	;
	switch v163 {
	case 0, 9, 10, 11, 12, 13, 32:
		goto L12
	default:
		goto L71
	}
L71:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v167 != int32(1) {
		goto L13
	} else {
		goto L72
	}
L72:
	;
	v171 = v163 - int32(33)
	if base.Ui32(v171) <= base.Ui32(int32(27)) {
		goto L16
	} else {
		goto L73
	}
L73:
	;
	goto L14
L74:
	;
	if v175 == int32(39) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	if v175 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L77:
	;
	v818 = int32(8)
	v819 = v31
	v821 = v33
	v824 = v37
	v826 = v39
	v827 = v40
	goto L3
L78:
	;
	goto L79
L79:
	;
	if v175 == int32(92) {
		goto L18
	} else {
		goto L80
	}
L80:
	;
	goto L76
L81:
	;
	v186 = int32(0)
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v188 = F_errsave_start(m, v187)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L40
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v212 = v31 - v211
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v210 <= v212+v213 {
		goto L92
	} else {
		goto L93
	}
L84:
	;
	if v188 == int32(0) {
		v792 = v186
		goto L6
	} else {
		goto L85
	}
L85:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L40
	} else {
		goto L86
	}
L86:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = v196
	if v195 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v200 = int32(_a_F_gettoken_tsvector_3)
	goto L89
L88:
	;
	v200 = int32(_a_F_gettoken_tsvector_4)
	goto L89
L89:
	;
	F_errmsg(m, v200, v20+int32(80))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L40
	} else {
		goto L90
	}
L90:
	;
	F_errsave_finish(m, v187, int32(_a_F_gettoken_tsvector_1), int32(148), int32(_a_F_gettoken_tsvector_5))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L40
	} else {
		goto L91
	}
L91:
	;
	v792 = v186
	goto L6
L92:
	;
	v217 = v210 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v217
	v219 = F_repalloc(m, v211, v217)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L40
	} else {
		goto L95
	}
L93:
	;
	v225 = v31
	v226 = v174
	goto L94
L94:
	;
	v227 = F_pg_mblen_cstr(m, v226)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L40
	} else {
		goto L96
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v219
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v225 = v219 + v212
	v226 = v223
	goto L94
L96:
	;
	if v227 != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	base.MemoryCopy(m, v225, v226, v227)
	goto L99
L98:
	;
	goto L99
L99:
	;
	v818 = int32(4)
	v819 = v227 + v225
	v821 = v33
	v824 = v37
	v826 = v39
	v827 = v40
	goto L3
L100:
	;
	v238 = v232 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v238
	v240 = F_repalloc(m, v47, v238)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L40
	} else {
		goto L103
	}
L101:
	;
	v246 = v31
	v248 = v48
	goto L102
L102:
	;
	v249 = F_pg_mblen_cstr(m, v248)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L40
	} else {
		goto L104
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v240
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v246 = v240 + v233
	v248 = v244
	goto L102
L104:
	;
	if v249 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	base.MemoryCopy(m, v246, v248, v249)
	goto L107
L106:
	;
	goto L107
L107:
	;
	v818 = int32(4)
	v819 = v249 + v246
	v821 = v33
	v824 = v37
	v826 = v39
	v827 = v40
	goto L3
L108:
	;
	if v256 == int32(0) {
		v792 = v254
		goto L6
	} else {
		goto L109
	}
L109:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L40
	} else {
		goto L110
	}
L110:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+96)) = v264
	if v263 != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v268 = int32(_a_F_gettoken_tsvector_3)
	goto L113
L112:
	;
	v268 = int32(_a_F_gettoken_tsvector_4)
	goto L113
L113:
	;
	F_errmsg(m, v268, v20+int32(96))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L40
	} else {
		goto L114
	}
L114:
	;
	F_errsave_finish(m, v255, int32(_a_F_gettoken_tsvector_1), int32(148), int32(_a_F_gettoken_tsvector_5))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L40
	} else {
		goto L115
	}
L115:
	;
	v792 = v254
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
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v33
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v37
	goto L116
L118:
	;
	goto L119
L119:
	;
	if v33 == int32(0) {
		goto L116
	} else {
		goto L120
	}
L120:
	;
	F_pfree(m, v33)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L40
	} else {
		goto L121
	}
L121:
	;
	goto L116
L122:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v284
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
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v69 - v286
	goto L127
L126:
	;
	goto L127
L127:
	;
	v289 = int32(1)
	if l5 == int32(0) {
		v792 = v289
		goto L6
	} else {
		goto L128
	}
L128:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v292
	v792 = v289
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
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v33
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v37
	goto L129
L131:
	;
	goto L132
L132:
	;
	if v33 == int32(0) {
		goto L129
	} else {
		goto L133
	}
L133:
	;
	F_pfree(m, v33)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L40
	} else {
		goto L134
	}
L134:
	;
	goto L129
L135:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v300
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
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v77 - v302
	goto L140
L139:
	;
	goto L140
L140:
	;
	v305 = int32(1)
	if l5 == int32(0) {
		v792 = v305
		goto L6
	} else {
		goto L141
	}
L141:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v308
	v792 = v305
	goto L6
L142:
	;
	if v40 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L143:
	;
	goto L144
L144:
	;
	v427 = int32(0)
	v428 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v429 = F_errsave_start(m, v428)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L40
	} else {
		goto L179
	}
L145:
	;
	v339 = v334 + v335<<(uint(int32(1))%32)
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v345 = v341
	goto L153
L146:
	;
	v321 = int32(4)
	v324 = F_palloc_mul(m, int32(2), v321)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L40
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	if v37+int32(1) < v40 {
		v334 = v33
		v335 = v37
		v336 = v40
		goto L145
	} else {
		goto L150
	}
L149:
	;
	v334 = v324
	v335 = int32(0)
	v336 = v321
	goto L145
L150:
	;
	v331 = v40 << (uint(int32(1)) % 32)
	v332 = F_repalloc_mul(m, v33, int32(2), v331)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L40
	} else {
		goto L151
	}
L151:
	;
	v334 = v332
	v335 = v37
	v336 = v331
	goto L145
L152:
	;
	v390 = int32(_a_F_gettoken_tsvector_6)
	if v390 < v389 {
		goto L168
	} else {
		goto L169
	}
L153:
	;
	v350 = v345 + int32(1)
	v351 = int32(*(*int8)(unsafe.Add(mBase, uint32(v345))))
	v352 = F___isspace(m, v351)
	mBase = m.M
	if v352 != 0 {
		v345 = v350
		goto L153
	} else {
		goto L155
	}
L154:
	;
	v353 = int32(1)
	switch v351&int32(255) - int32(43) {
	case 0:
		v359 = v353
		goto L157
	default:
		v361 = v351
		v362 = v345
		v363 = v353
		goto L156
	case 2:
		goto L158
	}
L155:
	;
	goto L154
L156:
	;
	v364 = int32(0)
	v366 = v361 - int32(48)
	if base.Ui32(v366) <= base.Ui32(int32(9)) {
		goto L159
	} else {
		goto L160
	}
L157:
	;
	v360 = int32(*(*int8)(unsafe.Add(mBase, uint32(v350))))
	v361 = v360
	v362 = v350
	v363 = v359
	goto L156
L158:
	;
	v359 = int32(0)
	goto L157
L159:
	;
	v369 = v364
	v370 = v366
	v371 = v362
	goto L162
L160:
	;
	v383 = v364
	goto L161
L161:
	;
	if v363 != 0 {
		goto L165
	} else {
		goto L166
	}
L162:
	;
	v373 = int32(10)
	v375 = v369*v373 - v370
	v376 = int32(*(*int8)(unsafe.Add(mBase, uint32(v371)+1)))
	v380 = v376 - int32(48)
	if base.Ui32(v380) < base.Ui32(v373) {
		v369 = v375
		v370 = v380
		v371 = v371 + int32(1)
		goto L162
	} else {
		goto L164
	}
L163:
	;
	v383 = v375
	goto L161
L164:
	;
	goto L163
L165:
	;
	v389 = int32(0) - v383
	goto L167
L166:
	;
	v389 = v383
	goto L167
L167:
	;
	goto L152
L168:
	;
	v394 = int32(_a_F_gettoken_tsvector_6)
	goto L170
L169:
	;
	v394 = v389 & v390
	goto L170
L170:
	;
	v395 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v339))))
	v398 = v394 | v395&int32(_a_F_gettoken_tsvector_7)
	*(*uint16)(unsafe.Add(mBase, uint32(v339))) = uint16(v398)
	if v394 == int32(0) {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v402 = int32(0)
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v404 = F_errsave_start(m, v403)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L40
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v339))) = uint16(v394)
	v818 = int32(7)
	v819 = v31
	v821 = v334
	v824 = v335 + int32(1)
	v826 = v39
	v827 = v336
	goto L3
L174:
	;
	if v404 == int32(0) {
		v792 = v402
		goto L6
	} else {
		goto L175
	}
L175:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L40
	} else {
		goto L176
	}
L176:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+112)) = v411
	F_errmsg(m, int32(_a_F_gettoken_tsvector_8), v20+int32(112))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L40
	} else {
		goto L177
	}
L177:
	;
	F_errsave_finish(m, v403, int32(_a_F_gettoken_tsvector_1), int32(335), int32(_a_F_gettoken_tsvector_2))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L40
	} else {
		goto L178
	}
L178:
	;
	v792 = v402
	goto L6
L179:
	;
	if v429 == int32(0) {
		v792 = v427
		goto L6
	} else {
		goto L180
	}
L180:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L40
	} else {
		goto L181
	}
L181:
	;
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+128)) = v437
	if v436 != 0 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v441 = int32(_a_F_gettoken_tsvector_3)
	goto L184
L183:
	;
	v441 = int32(_a_F_gettoken_tsvector_4)
	goto L184
L184:
	;
	F_errmsg(m, v441, v20+int32(128))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L40
	} else {
		goto L185
	}
L185:
	;
	F_errsave_finish(m, v428, int32(_a_F_gettoken_tsvector_1), int32(148), int32(_a_F_gettoken_tsvector_5))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L40
	} else {
		goto L186
	}
L186:
	;
	v792 = v427
	goto L6
L187:
	;
	if base.Ui32(int32(10)) <= base.Ui32((v453-int32(48))&int32(255)) {
		goto L15
	} else {
		goto L220
	}
L188:
	;
	if l3 != 0 {
		goto L208
	} else {
		goto L209
	}
L189:
	;
	v519 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33+v37<<(uint(int32(1))%32)-int32(2)))))
	if base.Ui32(int32(_a_F_gettoken_tsvector_9)) <= base.Ui32(v519) {
		goto L15
	} else {
		goto L206
	}
L190:
	;
	v506 = v33 + v37<<(uint(int32(1))%32) - int32(2)
	v507 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v506))))
	if base.Ui32(int32(_a_F_gettoken_tsvector_9)) <= base.Ui32(v507) {
		goto L15
	} else {
		goto L205
	}
L191:
	;
	v494 = v33 + v37<<(uint(int32(1))%32) - int32(2)
	v495 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v494))))
	if base.Ui32(int32(_a_F_gettoken_tsvector_9)) <= base.Ui32(v495) {
		goto L15
	} else {
		goto L204
	}
L192:
	;
	v458 = v33 + v37<<(uint(int32(1))%32) - int32(2)
	v459 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v458))))
	if base.Ui32(int32(_a_F_gettoken_tsvector_9)) <= base.Ui32(v459) {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v462 = int32(0)
	v463 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v464 = F_errsave_start(m, v463)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L40
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	v487 = v459 | int32(_a_F_gettoken_tsvector_7)
	*(*uint16)(unsafe.Add(mBase, uint32(v458))) = uint16(v487)
	v818 = int32(7)
	v819 = v31
	v821 = v33
	v824 = v37
	v826 = v39
	v827 = v40
	goto L3
L196:
	;
	if v464 == int32(0) {
		v792 = v462
		goto L6
	} else {
		goto L197
	}
L197:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L40
	} else {
		goto L198
	}
L198:
	;
	v471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+144)) = v472
	if v471 != 0 {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v476 = int32(_a_F_gettoken_tsvector_3)
	goto L201
L200:
	;
	v476 = int32(_a_F_gettoken_tsvector_4)
	goto L201
L201:
	;
	F_errmsg(m, v476, v20+int32(144))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L40
	} else {
		goto L202
	}
L202:
	;
	F_errsave_finish(m, v463, int32(_a_F_gettoken_tsvector_1), int32(148), int32(_a_F_gettoken_tsvector_5))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L40
	} else {
		goto L203
	}
L203:
	;
	v792 = v462
	goto L6
L204:
	;
	v499 = v495 | int32(_a_F_gettoken_tsvector_10)
	*(*uint16)(unsafe.Add(mBase, uint32(v494))) = uint16(v499)
	v818 = int32(7)
	v819 = v31
	v821 = v33
	v824 = v37
	v826 = v39
	v827 = v40
	goto L3
L205:
	;
	v511 = v507 | int32(_a_F_gettoken_tsvector_9)
	*(*uint16)(unsafe.Add(mBase, uint32(v506))) = uint16(v511)
	v818 = int32(7)
	v819 = v31
	v821 = v33
	v824 = v37
	v826 = v39
	v827 = v40
	goto L3
L206:
	;
	v818 = int32(7)
	v819 = v31
	v821 = v33
	v824 = v37
	v826 = v39
	v827 = v40
	goto L3
L207:
	;
	if l1 != 0 {
		goto L213
	} else {
		goto L214
	}
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v33
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v37
	goto L207
L209:
	;
	goto L210
L210:
	;
	if v33 == int32(0) {
		goto L207
	} else {
		goto L211
	}
L211:
	;
	F_pfree(m, v33)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L40
	} else {
		goto L212
	}
L212:
	;
	goto L207
L213:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v529
	goto L215
L214:
	;
	goto L215
L215:
	;
	if l2 != 0 {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v31 - v531
	goto L218
L217:
	;
	goto L218
L218:
	;
	v534 = int32(1)
	if l5 == int32(0) {
		v792 = v534
		goto L6
	} else {
		goto L219
	}
L219:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v537
	v792 = v534
	goto L6
L220:
	;
	v818 = int32(7)
	v819 = v31
	v821 = v33
	v824 = v37
	v826 = v39
	v827 = v40
	goto L3
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v30
	F_errmsg_internal(m, int32(_a_F_gettoken_tsvector_11), v20)
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L40
	} else {
		goto L222
	}
L222:
	;
	F_errfinish(m, int32(_a_F_gettoken_tsvector_1), int32(378), int32(_a_F_gettoken_tsvector_2))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L40
	} else {
		goto L223
	}
L223:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L224:
	;
	goto L8
L225:
	;
	goto L12
L226:
	;
	if v579 != 0 {
		goto L227
	} else {
		goto L228
	}
L227:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L40
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	m.G0 = v576 + int32(16)
	v792 = int32(0)
	goto L6
L230:
	;
	v584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
	v585 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v576))) = v585
	if v584 != 0 {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	v589 = int32(_a_F_gettoken_tsvector_3)
	goto L233
L232:
	;
	v589 = int32(_a_F_gettoken_tsvector_4)
	goto L233
L233:
	;
	F_errmsg(m, v589, v576)
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L40
	} else {
		goto L234
	}
L234:
	;
	F_errsave_finish(m, v578, int32(_a_F_gettoken_tsvector_1), int32(148), int32(_a_F_gettoken_tsvector_5))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L40
	} else {
		goto L235
	}
L235:
	;
	goto L229
L236:
	;
	goto L13
L237:
	;
	goto L12
L238:
	;
	v618 = v611 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v618
	v620 = F_repalloc(m, v612, v618)
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L40
	} else {
		goto L241
	}
L239:
	;
	v624 = v612
	v625 = v31
	goto L240
L240:
	;
	if v624 == v625 {
		goto L242
	} else {
		goto L243
	}
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v620
	v624 = v620
	v625 = v620 + v613
	goto L240
L242:
	;
	v628 = int32(0)
	v629 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v630 = F_errsave_start(m, v629)
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L40
	} else {
		goto L245
	}
L243:
	;
	goto L244
L244:
	;
	v652 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v625))) = uint8(v652)
	if l3 != 0 {
		goto L254
	} else {
		goto L255
	}
L245:
	;
	if v630 == int32(0) {
		v792 = v628
		goto L6
	} else {
		goto L246
	}
L246:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L40
	} else {
		goto L247
	}
L247:
	;
	v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
	v638 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v638
	if v637 != 0 {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	v642 = int32(_a_F_gettoken_tsvector_3)
	goto L250
L249:
	;
	v642 = int32(_a_F_gettoken_tsvector_4)
	goto L250
L250:
	;
	F_errmsg(m, v642, v20+int32(48))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L40
	} else {
		goto L251
	}
L251:
	;
	F_errsave_finish(m, v629, int32(_a_F_gettoken_tsvector_1), int32(148), int32(_a_F_gettoken_tsvector_5))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L40
	} else {
		goto L252
	}
L252:
	;
	v792 = v628
	goto L6
L253:
	;
	if l1 != 0 {
		goto L259
	} else {
		goto L260
	}
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v33
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v37
	goto L253
L255:
	;
	goto L256
L256:
	;
	if v33 == int32(0) {
		goto L253
	} else {
		goto L257
	}
L257:
	;
	F_pfree(m, v33)
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L40
	} else {
		goto L258
	}
L258:
	;
	goto L253
L259:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v660
	goto L261
L260:
	;
	goto L261
L261:
	;
	if l2 != 0 {
		goto L262
	} else {
		goto L263
	}
L262:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v625 - v662
	goto L264
L263:
	;
	goto L264
L264:
	;
	v665 = int32(1)
	if l5 == int32(0) {
		v792 = v665
		goto L6
	} else {
		goto L265
	}
L265:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v668
	v792 = v665
	goto L6
L266:
	;
	if v31 == v670 {
		goto L269
	} else {
		goto L270
	}
L267:
	;
	goto L268
L268:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v721 = v31 - v670
	v722 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v720 <= v721+v722 {
		goto L296
	} else {
		goto L297
	}
L269:
	;
	v674 = int32(0)
	v675 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v676 = F_errsave_start(m, v675)
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L40
	} else {
		goto L272
	}
L270:
	;
	goto L271
L271:
	;
	v698 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v698)
	v700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v700 != int32(1) {
		goto L280
	} else {
		goto L281
	}
L272:
	;
	if v676 == int32(0) {
		v792 = v674
		goto L6
	} else {
		goto L273
	}
L273:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L40
	} else {
		goto L274
	}
L274:
	;
	v683 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
	v684 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v684
	if v683 != 0 {
		goto L275
	} else {
		goto L276
	}
L275:
	;
	v688 = int32(_a_F_gettoken_tsvector_3)
	goto L277
L276:
	;
	v688 = int32(_a_F_gettoken_tsvector_4)
	goto L277
L277:
	;
	F_errmsg(m, v688, v20-int32(-64))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L40
	} else {
		goto L278
	}
L278:
	;
	F_errsave_finish(m, v675, int32(_a_F_gettoken_tsvector_1), int32(148), int32(_a_F_gettoken_tsvector_5))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L40
	} else {
		goto L279
	}
L279:
	;
	v792 = v674
	goto L6
L280:
	;
	v818 = int32(6)
	v819 = v31
	v821 = v33
	v824 = v37
	v826 = v39
	v827 = v40
	goto L3
L281:
	;
	goto L282
L282:
	;
	if l3 != 0 {
		goto L284
	} else {
		goto L285
	}
L283:
	;
	if l1 != 0 {
		goto L289
	} else {
		goto L290
	}
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v33
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v37
	goto L283
L285:
	;
	goto L286
L286:
	;
	if v33 == int32(0) {
		goto L283
	} else {
		goto L287
	}
L287:
	;
	F_pfree(m, v33)
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L40
	} else {
		goto L288
	}
L288:
	;
	goto L283
L289:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v710
	goto L291
L290:
	;
	goto L291
L291:
	;
	if l2 != 0 {
		goto L292
	} else {
		goto L293
	}
L292:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v31 - v712
	goto L294
L293:
	;
	goto L294
L294:
	;
	v715 = int32(1)
	if l5 == int32(0) {
		v792 = v715
		goto L6
	} else {
		goto L295
	}
L295:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v718
	v792 = v715
	goto L6
L296:
	;
	v726 = v720 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v726
	v728 = F_repalloc(m, v670, v726)
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L40
	} else {
		goto L299
	}
L297:
	;
	v734 = v31
	v735 = v162
	goto L298
L298:
	;
	v736 = F_pg_mblen_cstr(m, v735)
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L40
	} else {
		goto L300
	}
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v728
	v732 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v734 = v728 + v721
	v735 = v732
	goto L298
L300:
	;
	if v736 != 0 {
		goto L301
	} else {
		goto L302
	}
L301:
	;
	base.MemoryCopy(m, v734, v735, v736)
	goto L303
L302:
	;
	goto L303
L303:
	;
	v818 = int32(2)
	v819 = v736 + v734
	v821 = v33
	v824 = v37
	v826 = v39
	v827 = v40
	goto L3
L304:
	;
	goto L9
L305:
	;
	goto L8
L306:
	;
	if v754 == int32(0) {
		v792 = v752
		goto L6
	} else {
		goto L307
	}
L307:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L40
	} else {
		goto L308
	}
L308:
	;
	v761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
	v762 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v762
	if v761 != 0 {
		goto L309
	} else {
		goto L310
	}
L309:
	;
	v766 = int32(_a_F_gettoken_tsvector_3)
	goto L311
L310:
	;
	v766 = int32(_a_F_gettoken_tsvector_4)
	goto L311
L311:
	;
	F_errmsg(m, v766, v20+int32(16))
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L40
	} else {
		goto L312
	}
L312:
	;
	F_errsave_finish(m, v753, int32(_a_F_gettoken_tsvector_1), int32(148), int32(_a_F_gettoken_tsvector_5))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L40
	} else {
		goto L313
	}
L313:
	;
	v792 = v752
	goto L6
L314:
	;
	v784 = F_pg_mblen_cstr(m, v86)
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L40
	} else {
		goto L315
	}
L315:
	;
	if v784 != 0 {
		goto L316
	} else {
		goto L317
	}
L316:
	;
	base.MemoryCopy(m, v31, v86, v784)
	goto L318
L317:
	;
	goto L318
L318:
	;
	v818 = int32(2)
	v819 = v784 + v31
	v821 = v33
	v824 = v37
	v826 = v39
	v827 = v40
	goto L3
L319:
	;
	v831 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v829 + v831
	v30 = v818
	v31 = v819
	v33 = v821
	v37 = v824
	v39 = v826
	v40 = v827
	goto L1
}
func F_ginadjustmembers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	var v13 int32
	_ = v13
	Fn14296(m, l0, l1, l2, l3, int32(_a_F_ginadjustmembers_0), int32(324), int32(_a_F_ginadjustmembers_1), int32(_a_F_ginadjustmembers_2), int32(12), int32(242), int32(7))
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
func F_ginint4_consistent(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int64
	_ = v45
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
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
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int64
	_ = v88
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v18 = base.I32_wrap_i64(v15) & int32(_a_F_ginint4_consistent_0)
	switch v18 - int32(3) {
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
	m.G0 = v10 + int32(16)
	return v88
L2:
	;
	v88 = int64(1)
	goto L1
L3:
	;
	v84 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v84)
	goto L2
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L23
	} else {
		goto L26
	}
L5:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v62 = F_pg_detoast_datum(m, v61)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L23
	} else {
		goto L24
	}
L6:
	;
	v42 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v42)
	v45 = int64(1)
	if v13 <= v42 {
		v88 = v45
		goto L1
	} else {
		goto L16
	}
L7:
	;
	v23 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v23)
	v25 = int32(0)
	v26 = int64(1)
	if v13 <= v25 {
		v88 = v26
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v21 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v21)
	goto L2
L9:
	;
	v29 = v25
	goto L10
L10:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+v14))))
	if v37 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v88 = int64(0)
	goto L1
L12:
	;
	v39 = v29 + int32(1)
	if v13 != v39 {
		v29 = v39
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
	v88 = v26
	goto L1
L16:
	;
	v48 = v42
	goto L17
L17:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48+v14))))
	if v56 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v88 = int64(0)
	goto L1
L19:
	;
	v58 = v48 + int32(1)
	if v13 != v58 {
		v48 = v58
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
	v88 = v45
	goto L1
L23:
	;
	return int64(0)
L24:
	;
	v66 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v66)
	v68 = F_gin_bool_consistent(m, v62, v14)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v88 = base.I64_extend_i32_u(v68)
	goto L1
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v18
	F_errmsg_internal(m, int32(_a_F_ginint4_consistent_1), v10)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_ginint4_consistent_2), int32(177), int32(_a_F_ginint4_consistent_3))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
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
				v15 = v11 * int32(56)
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
	var v24 int32
	_ = v24
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
	var v35 int32
	_ = v35
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int64
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
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
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int64
	_ = v134
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v298 int32
	_ = v298
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
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
	var v348 int32
	_ = v348
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
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
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v442 int32
	_ = v442
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
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
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v565 int32
	_ = v565
	var v571 int32
	_ = v571
	var v583 int64
	_ = v583
	var v584 int64
	_ = v584
	var v589 int32
	_ = v589
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v608 int32
	_ = v608
	var v617 int32
	_ = v617
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v625 int64
	_ = v625
	var v629 int64
	_ = v629
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v650 int32
	_ = v650
	var v660 int32
	_ = v660
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	v18 = m.G0
	v20 = v18 - int32(272)
	m.G0 = v20
	v24 = F_SearchSysCache1(m, int32(14), base.I64_extend_i32_u(l0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v41)+56))
	if int32(0) < v298 {
		goto L62
	} else {
		goto L63
	}
L2:
	;
	return int32(0)
L3:
	;
	if v24 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
	v30 = v28 + v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+84))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)+92))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)+80))
	v34 = F_get_opfamily_name(m, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
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
	v271 = m.ExcPending
	if v271 != 0 {
		goto L2
	} else {
		goto L59
	}
L7:
	;
	v38 = base.I64_extend_i32_u(v33)
	v39 = int64(0)
	v41 = F_SearchSysCacheList(m, int32(4), int32(1), v38, v39, v39)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v43 = int32(1)
	v46 = int64(0)
	v48 = F_SearchSysCacheList(m, int32(5), v43, v38, v46, v46)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v48)+56))
	if v50 <= int32(0) {
		v285 = v43
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v32 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v53 = v32
	goto L13
L12:
	;
	v53 = v31
	goto L13
L13:
	;
	v62 = int32(0)
	v64 = v43
	goto L14
L14:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v48-int32(-64)+v62<<(uint(int32(2))%32))))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+72))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+22)))
	v83 = v81 + v82
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
	if v84 == v85 {
		v114 = v64
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v285 = v262
	goto L1
L16:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	if v115 != v31 {
		v262 = v114
		goto L24
	} else {
		goto L25
	}
L17:
	;
	v87 = int32(0)
	v90 = F_errstart(m, int32(17), v87)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	if v90 == int32(0) {
		v114 = v87
		goto L16
	} else {
		goto L19
	}
L19:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v83)+20))
	v98 = F_format_procedure(m, v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+264)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v20)+260)) = int32(_a_F_ginvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+256)) = v34
	F_errmsg(m, int32(_a_F_ginvalidate_1), v20+int32(256))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_ginvalidate_2), int32(85), int32(_a_F_ginvalidate_3))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	v114 = v87
	goto L16
L24:
	;
	v265 = v62 + int32(1)
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v48)+56))
	if v265 < v266 {
		v62 = v265
		v64 = v262
		goto L14
	} else {
		goto L58
	}
L25:
	;
	v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83)+16)))
	switch v117 - int32(1) {
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
	v243 = m.ExcPending
	if v243 != 0 {
		goto L2
	} else {
		goto L54
	}
L27:
	;
	v230 = int32(0)
	v233 = F_errstart(m, int32(17), v230)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L2
	} else {
		goto L52
	}
L28:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v83)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+132)) = int64(9796820404457)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+128)) = v31
	v228 = F_check_amproc_signature(m, v218, int32(2281), int32(0), int32(2), int32(3), v20+int32(128))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L2
	} else {
		goto L50
	}
L29:
	;
	v209 = int32(0)
	v212 = F_errstart(m, int32(17), v209)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L2
	} else {
		goto L48
	}
L30:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v83)+20))
	v205 = F_check_amoptsproc_signature(m, v204)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L2
	} else {
		goto L46
	}
L31:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v83)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+244)) = int64(9796820404457)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+236)) = int64(9796820402199)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+232)) = v31
	*(*int64)(unsafe.Add(mBase, uint32(v20)+224)) = int64(90194315497)
	v196 = int32(7)
	v200 = F_check_amproc_signature(m, v186, int32(18), int32(0), v196, v196, v20+int32(224))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L2
	} else {
		goto L44
	}
L32:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v83)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+216)) = int64(9796820402197)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+212)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v20)+208)) = v53
	v178 = int32(4)
	v182 = F_check_amproc_signature(m, v171, int32(23), int32(0), v178, v178, v20+int32(208))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L2
	} else {
		goto L42
	}
L33:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v83)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v20+int32(204)))) = int32(2281)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+196)) = int64(9796820404457)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+188)) = int64(9796820402199)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+184)) = v31
	*(*int64)(unsafe.Add(mBase, uint32(v20)+176)) = int64(90194315497)
	v167 = F_check_amproc_signature(m, v151, int32(16), int32(0), int32(6), int32(8), v20+int32(176))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L2
	} else {
		goto L40
	}
L34:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v83)+20))
	v134 = int64(9796820404457)
	*(*int64)(unsafe.Add(mBase, uint32(v20+int32(164)))) = v134
	*(*int64)(unsafe.Add(mBase, uint32(v20)+156)) = v134
	*(*int64)(unsafe.Add(mBase, uint32(v20)+148)) = int64(90194315497)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+144)) = v31
	v147 = F_check_amproc_signature(m, v133, int32(2281), int32(0), int32(5), int32(7), v20+int32(144))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L2
	} else {
		goto L38
	}
L35:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v83)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+116)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v20)+112)) = v53
	v125 = int32(2)
	v129 = F_check_amproc_signature(m, v120, int32(23), int32(0), v125, v125, v20+int32(112))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	if v129 == int32(0) {
		goto L27
	} else {
		goto L37
	}
L37:
	;
	v262 = v114
	goto L24
L38:
	;
	if v147 == int32(0) {
		goto L27
	} else {
		goto L39
	}
L39:
	;
	v262 = v114
	goto L24
L40:
	;
	if v167 == int32(0) {
		goto L27
	} else {
		goto L41
	}
L41:
	;
	v262 = v114
	goto L24
L42:
	;
	if v182 == int32(0) {
		goto L27
	} else {
		goto L43
	}
L43:
	;
	v262 = v114
	goto L24
L44:
	;
	if v200 == int32(0) {
		goto L27
	} else {
		goto L45
	}
L45:
	;
	v262 = v114
	goto L24
L46:
	;
	if v205 == int32(0) {
		goto L27
	} else {
		goto L47
	}
L47:
	;
	v262 = v114
	goto L24
L48:
	;
	if v212 == int32(0) {
		v262 = v209
		goto L24
	} else {
		goto L49
	}
L49:
	;
	v239 = int32(145)
	v240 = int32(_a_F_ginvalidate_4)
	goto L26
L50:
	;
	if v228 != 0 {
		v262 = v114
		goto L24
	} else {
		goto L51
	}
L51:
	;
	goto L27
L52:
	;
	if v233 == int32(0) {
		v262 = v230
		goto L24
	} else {
		goto L53
	}
L53:
	;
	v239 = int32(157)
	v240 = int32(_a_F_ginvalidate_5)
	goto L26
L54:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v83)+20))
	v245 = F_format_procedure(m, v244)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L2
	} else {
		goto L55
	}
L55:
	;
	v247 = int32(*(*int16)(unsafe.Add(mBase, uint32(v83)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+108)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v20)+104)) = v245
	*(*int32)(unsafe.Add(mBase, uint32(v20)+100)) = int32(_a_F_ginvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+96)) = v34
	F_errmsg(m, v240, v20+int32(96))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L2
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_ginvalidate_2), v239, int32(_a_F_ginvalidate_3))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L2
	} else {
		goto L57
	}
L57:
	;
	v262 = int32(0)
	goto L24
L58:
	;
	goto L15
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = l0
	F_errmsg_internal(m, int32(_a_F_ginvalidate_6), v20)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L2
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_ginvalidate_2), int32(51), int32(_a_F_ginvalidate_3))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
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
	v306 = int32(0)
	v308 = v285
	goto L65
L63:
	;
	v442 = v285
	goto L64
L64:
	;
	v455 = F_identify_opfamily_groups(m, v41, v48)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L2
	} else {
		goto L97
	}
L65:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v41-int32(-64)+v306<<(uint(int32(2))%32))))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v324)+72))
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325)+22)))
	v327 = v325 + v326
	v328 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v327)+16)))
	if base.Ui32(int32(_a_F_ginvalidate_7)) < base.Ui32((v328+int32(-64))&int32(_a_F_ginvalidate_8)) {
		v364 = v308
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v442 = v433
	goto L64
L67:
	;
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v327)+18)))
	if v366 == int32(115) {
		goto L76
	} else {
		goto L77
	}
L68:
	;
	v335 = int32(0)
	v338 = F_errstart(m, int32(17), v335)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L2
	} else {
		goto L69
	}
L69:
	;
	if v338 == int32(0) {
		v364 = v335
		goto L67
	} else {
		goto L70
	}
L70:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L2
	} else {
		goto L71
	}
L71:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v327)+20))
	v346 = F_format_operator(m, v345)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L2
	} else {
		goto L72
	}
L72:
	;
	v348 = int32(*(*int16)(unsafe.Add(mBase, uint32(v327)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+92)) = v348
	*(*int32)(unsafe.Add(mBase, uint32(v20)+88)) = v346
	*(*int32)(unsafe.Add(mBase, uint32(v20)+84)) = int32(_a_F_ginvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = v34
	F_errmsg(m, int32(_a_F_ginvalidate_9), v20+int32(80))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L2
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(_a_F_ginvalidate_2), int32(176), int32(_a_F_ginvalidate_3))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L2
	} else {
		goto L74
	}
L74:
	;
	v364 = v335
	goto L67
L75:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v327)+20))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v327)+8))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v327)+12))
	v404 = F_check_amop_signature(m, v400, int32(16), v402, v403)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L2
	} else {
		goto L87
	}
L76:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v327)+28))
	if v369 == int32(0) {
		v399 = v364
		goto L75
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v372 = int32(0)
	v375 = F_errstart(m, int32(17), v372)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L2
	} else {
		goto L80
	}
L79:
	;
	goto L78
L80:
	;
	if v375 == int32(0) {
		v399 = v372
		goto L75
	} else {
		goto L81
	}
L81:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L2
	} else {
		goto L82
	}
L82:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v327)+20))
	v383 = F_format_operator(m, v382)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L2
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+72)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v20)+68)) = int32(_a_F_ginvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v34
	F_errmsg(m, int32(_a_F_ginvalidate_10), v20-int32(-64))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L2
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_ginvalidate_2), int32(188), int32(_a_F_ginvalidate_3))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L2
	} else {
		goto L85
	}
L85:
	;
	v399 = v372
	goto L75
L86:
	;
	v435 = v306 + int32(1)
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v41)+56))
	if v435 < v436 {
		v306 = v435
		v308 = v433
		goto L65
	} else {
		goto L95
	}
L87:
	;
	if v404 != 0 {
		v433 = v399
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v406 = int32(0)
	v409 = F_errstart(m, int32(17), v406)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L2
	} else {
		goto L89
	}
L89:
	;
	if v409 == int32(0) {
		v433 = v406
		goto L86
	} else {
		goto L90
	}
L90:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L2
	} else {
		goto L91
	}
L91:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v327)+20))
	v417 = F_format_operator(m, v416)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L2
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = v417
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = int32(_a_F_ginvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v34
	F_errmsg(m, int32(_a_F_ginvalidate_11), v20+int32(48))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L2
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(_a_F_ginvalidate_2), int32(201), int32(_a_F_ginvalidate_3))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L2
	} else {
		goto L94
	}
L94:
	;
	v433 = v406
	goto L86
L95:
	;
	goto L66
L96:
	;
	v565 = v30 + int32(8)
	v571 = v442
	v583 = int64(1)
	goto L132
L97:
	;
	if v455 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v547 = int32(0)
	goto L96
L99:
	;
	goto L100
L100:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v455)+4))
	if v460 <= int32(0) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v547 = int32(0)
	goto L96
L102:
	;
	goto L103
L103:
	;
	v464 = int32(0)
	if v460 != int32(1) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v468 = int32(0)
	if v468 < v460 {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	v520 = v464
	v522 = v464
	goto L106
L106:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v455)+12))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v537+v522<<(uint(int32(2))%32))))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v541)))
	if v542 != v31 {
		v547 = v520
		goto L96
	} else {
		goto L126
	}
L107:
	;
	v471 = v460
	goto L109
L108:
	;
	v471 = v468
	goto L109
L109:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v455)+12))
	v477 = int32(0)
	v479 = v477
	v481 = v464
	v485 = v477
	goto L110
L110:
	;
	v498 = v476 + v481<<(uint(int32(2))%32)
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v498)))
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v499)))
	if v31 == v500 {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	if v471&int32(1) == int32(0) {
		v547 = v512
		goto L96
	} else {
		goto L125
	}
L112:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v499)+4))
	if v502 == v31 {
		goto L115
	} else {
		goto L116
	}
L113:
	;
	v505 = v479
	goto L114
L114:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v498)+4))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v506)))
	if v31 == v507 {
		goto L118
	} else {
		goto L119
	}
L115:
	;
	v504 = v499
	goto L117
L116:
	;
	v504 = v479
	goto L117
L117:
	;
	v505 = v504
	goto L114
L118:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v506)+4))
	if v509 == v31 {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	v512 = v505
	goto L120
L120:
	;
	v513 = int32(2)
	v514 = v481 + v513
	v516 = v485 + v513
	if v516 != v471&int32(2147483646) {
		v479 = v512
		v481 = v514
		v485 = v516
		goto L110
	} else {
		goto L124
	}
L121:
	;
	v511 = v506
	goto L123
L122:
	;
	v511 = v505
	goto L123
L123:
	;
	v512 = v511
	goto L120
L124:
	;
	goto L111
L125:
	;
	v520 = v512
	v522 = v514
	goto L106
L126:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v541)+4))
	if v544 == v31 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v546 = v541
	goto L129
L128:
	;
	v546 = v520
	goto L129
L129:
	;
	v547 = v546
	goto L96
L130:
	;
	F_ReleaseCatCacheList(m, v48)
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L2
	} else {
		goto L156
	}
L131:
	;
	v641 = int32(0)
	v644 = F_errstart(m, int32(17), v641)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L2
	} else {
		goto L151
	}
L132:
	;
	if v547 != 0 {
		goto L136
	} else {
		goto L137
	}
L133:
	;
	v635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v547)+16)))
	if v635&int32(80) != 0 {
		v667 = v633
		goto L130
	} else {
		goto L150
	}
L134:
	;
	goto L133
L135:
	;
	v629 = v583 + int64(1)
	if v629 != int64(7) {
		v583 = v629
		goto L132
	} else {
		goto L149
	}
L136:
	;
	v584 = *(*int64)(unsafe.Add(mBase, uint32(v547)+16))
	if base.I32_wrap_i64(int64(base.Ui64(v584)>>(uint(v583)%64)))&int32(1) != 0 {
		goto L135
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	v589 = base.I32_wrap_i64(v583)
	if base.B2i32(v589&int32(5) == int32(4))|base.B2i32(v589&int32(3) == int32(1)) != 0 {
		v623 = v571
		goto L140
	} else {
		goto L141
	}
L139:
	;
	goto L138
L140:
	;
	v625 = v583 + int64(1)
	if v625 != int64(7) {
		v571 = v623
		v583 = v625
		goto L132
	} else {
		goto L147
	}
L141:
	;
	v599 = int32(0)
	v602 = F_errstart(m, int32(17), v599)
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L2
	} else {
		goto L142
	}
L142:
	;
	if v602 == int32(0) {
		v623 = v599
		goto L140
	} else {
		goto L143
	}
L143:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L2
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = int32(_a_F_ginvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v565
	F_errmsg(m, int32(_a_F_ginvalidate_12), v20+int32(32))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L2
	} else {
		goto L145
	}
L145:
	;
	F_errfinish(m, int32(_a_F_ginvalidate_2), int32(242), int32(_a_F_ginvalidate_3))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L2
	} else {
		goto L146
	}
L146:
	;
	v623 = v599
	goto L140
L147:
	;
	if v547 != 0 {
		v633 = v623
		goto L134
	} else {
		goto L148
	}
L148:
	;
	goto L131
L149:
	;
	v633 = v571
	goto L134
L150:
	;
	goto L131
L151:
	;
	if v644 == int32(0) {
		v667 = v641
		goto L130
	} else {
		goto L152
	}
L152:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L2
	} else {
		goto L153
	}
L153:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+24)) = int64(25769803780)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = int32(_a_F_ginvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v565
	F_errmsg(m, int32(_a_F_ginvalidate_13), v20+int32(16))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L2
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(_a_F_ginvalidate_2), int32(253), int32(_a_F_ginvalidate_3))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L2
	} else {
		goto L155
	}
L155:
	;
	v667 = v641
	goto L130
L156:
	;
	F_ReleaseCatCacheList(m, v41)
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L2
	} else {
		goto L157
	}
L157:
	;
	F_ReleaseCatCache(m, v24)
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L2
	} else {
		goto L158
	}
L158:
	;
	m.G0 = v20 + int32(272)
	return v667 & int32(1)
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
	Fn14297(m, l0, l1, int32(_a_F_gistcheckpage_0), int32(812), int32(_a_F_gistcheckpage_1), int32(801))
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
	F_LockBufferInternal(m, v16, int32(3))
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
	var v43 int32
	_ = v43
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
	var v225 int32
	_ = v225
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
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v454 int64
	_ = v454
	var v455 int32
	_ = v455
	var v461 int64
	_ = v461
	var v462 int32
	_ = v462
	var v468 int64
	_ = v468
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v502 int32
	_ = v502
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v569 int32
	_ = v569
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v629 int32
	_ = v629
	var v636 int32
	_ = v636
	var v641 int32
	_ = v641
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v697 int32
	_ = v697
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int64
	_ = v723
	var v728 int32
	_ = v728
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v750 int32
	_ = v750
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v774 int64
	_ = v774
	var v775 int32
	_ = v775
	var v791 int32
	_ = v791
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v810 int32
	_ = v810
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v824 int32
	_ = v824
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v835 int32
	_ = v835
	var v839 int32
	_ = v839
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v853 int32
	_ = v853
	var v858 int32
	_ = v858
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v878 int32
	_ = v878
	var v884 int32
	_ = v884
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v898 int32
	_ = v898
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v943 int32
	_ = v943
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v950 int32
	_ = v950
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v975 int32
	_ = v975
	var v980 int32
	_ = v980
	var v982 int32
	_ = v982
	var v985 int32
	_ = v985
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v997 int32
	_ = v997
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1011 int32
	_ = v1011
	var v1012 int64
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1014 int64
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1016 int64
	_ = v1016
	var v1020 int32
	_ = v1020
	var v1027 int32
	_ = v1027
	var v1031 int32
	_ = v1031
	var v1036 int32
	_ = v1036
	var v1040 int32
	_ = v1040
	var v1046 int32
	_ = v1046
	var v1051 int32
	_ = v1051
	var v1091 int32
	_ = v1091
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1112 int32
	_ = v1112
	var v1128 int32
	_ = v1128
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1152 int32
	_ = v1152
	var v1187 int32
	_ = v1187
	var v1193 int32
	_ = v1193
	var v1195 int32
	_ = v1195
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1211 int32
	_ = v1211
	var v1223 int32
	_ = v1223
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1272 int32
	_ = v1272
	var v1276 int32
	_ = v1276
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1282 int32
	_ = v1282
	var v1284 int32
	_ = v1284
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1326 int32
	_ = v1326
	var v1349 int32
	_ = v1349
	var v1374 int32
	_ = v1374
	var v1377 int32
	_ = v1377
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1410 int32
	_ = v1410
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1420 int32
	_ = v1420
	var v1421 int32
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1423 int32
	_ = v1423
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1437 int32
	_ = v1437
	var v1440 int32
	_ = v1440
	var v1442 int32
	_ = v1442
	var v1446 int32
	_ = v1446
	var v1474 int32
	_ = v1474
	var v1477 int32
	_ = v1477
	var v1480 int32
	_ = v1480
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1486 int32
	_ = v1486
	var v1487 int32
	_ = v1487
	var v1488 int32
	_ = v1488
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1495 int32
	_ = v1495
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1500 int32
	_ = v1500
	var v1501 int32
	_ = v1501
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1507 int32
	_ = v1507
	var v1514 int32
	_ = v1514
	var v1522 int32
	_ = v1522
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1539 int32
	_ = v1539
	var v1542 int32
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1544 int32
	_ = v1544
	var v1549 int32
	_ = v1549
	var v1550 int32
	_ = v1550
	var v1552 int32
	_ = v1552
	var v1571 int32
	_ = v1571
	var v1584 int32
	_ = v1584
	var v1586 int32
	_ = v1586
	var v1587 int32
	_ = v1587
	var v1615 int32
	_ = v1615
	var v1616 int32
	_ = v1616
	var v1617 int32
	_ = v1617
	var v1621 int32
	_ = v1621
	var v1627 int32
	_ = v1627
	var v1629 int32
	_ = v1629
	var v1635 int32
	_ = v1635
	var v1637 int32
	_ = v1637
	var v1638 int32
	_ = v1638
	var v1642 int32
	_ = v1642
	var v1648 int32
	_ = v1648
	var v1650 int32
	_ = v1650
	var v1656 int32
	_ = v1656
	var v1659 int32
	_ = v1659
	var v1660 int32
	_ = v1660
	var v1664 int32
	_ = v1664
	var v1667 int32
	_ = v1667
	var v1668 int32
	_ = v1668
	var v1669 int32
	_ = v1669
	var v1670 int32
	_ = v1670
	var v1672 int32
	_ = v1672
	var v1675 int32
	_ = v1675
	var v1676 int32
	_ = v1676
	var v1701 int32
	_ = v1701
	var v1702 int32
	_ = v1702
	var v1705 int32
	_ = v1705
	var v1735 int32
	_ = v1735
	var v1739 int32
	_ = v1739
	var v1744 int32
	_ = v1744
	var v1746 int32
	_ = v1746
	var v1748 int32
	_ = v1748
	var v1773 int32
	_ = v1773
	var v1774 int32
	_ = v1774
	var v1777 int32
	_ = v1777
	var v1778 int32
	_ = v1778
	var v1782 int32
	_ = v1782
	var v1783 int32
	_ = v1783
	var v1784 int32
	_ = v1784
	var v1786 int32
	_ = v1786
	var v1789 int32
	_ = v1789
	var v1818 int64
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1823 int64
	_ = v1823
	var v1824 int32
	_ = v1824
	var v1851 int64
	_ = v1851
	var v1869 int32
	_ = v1869
	var v1882 int32
	_ = v1882
	var v1884 int32
	_ = v1884
	var v1911 int32
	_ = v1911
	var v1927 int32
	_ = v1927
	var v1940 int32
	_ = v1940
	var v1942 int32
	_ = v1942
	var v1943 int32
	_ = v1943
	var v1985 int32
	_ = v1985
	var v1995 int64
	_ = v1995
	var v2000 int32
	_ = v2000
	var v2006 int32
	_ = v2006
	var v2008 int32
	_ = v2008
	var v2014 int32
	_ = v2014
	var v2015 int32
	_ = v2015
	var v2018 int64
	_ = v2018
	var v2020 int32
	_ = v2020
	var v2021 int32
	_ = v2021
	var v2022 int32
	_ = v2022
	var v2024 int32
	_ = v2024
	var v2030 int32
	_ = v2030
	var v2032 int32
	_ = v2032
	var v2043 int32
	_ = v2043
	var v2044 int32
	_ = v2044
	var v2052 int32
	_ = v2052
	var v2057 int32
	_ = v2057
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
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v34+(l3^int32(-1))*int32(56))+16))
	v49 = v40
	goto L1
L3:
	;
	goto L4
L4:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[1]))
	v43 = int32(56)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v42+l3*v43-v43)+16))
	v49 = v48
	goto L1
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2043 = m.ExcPending
	if v2043 != 0 {
		goto L59
	} else {
		goto L348
	}
L6:
	;
	if l8 != 0 {
		goto L341
	} else {
		goto L342
	}
L7:
	;
	if v653 != 0 {
		goto L204
	} else {
		goto L205
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L59
	} else {
		goto L198
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
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L59
	} else {
		goto L195
	}
L16:
	;
	v943 = int32(_a_F_gistplacetopage_0)
	v945 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[4]))
	v946 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[4])) = v945 + v946
	v950 = v7 - v946
	if base.Ui32(v950&int32(_a_F_gistplacetopage_1)) <= base.Ui32(int32(2047)) {
		goto L162
	} else {
		goto L163
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
	v621 = F_gistextractpage(m, v67, v29+int32(44))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L59
	} else {
		goto L107
	}
L34:
	;
	v502 = int32(0)
	if l5 <= v502 {
		v569 = l1
		goto L94
	} else {
		goto L95
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
	v225 = int32(0)
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
	v272 = v225 + v269
	if v272 != v200&int32(-2) {
		v220 = v270
		v221 = v268
		v225 = v272
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
	*(*int64)(unsafe.Add(mBase, uint32(v67))) = v468
	v470 = int32(_a_F_gistplacetopage_0)
	v472 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[4])) = v472 - int32(1)
	goto L34
L64:
	;
	v461 = F_XLogGetFakeLSN(m, l0)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L59
	} else {
		goto L92
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
	if v399 <= int32(1) {
		goto L72
	} else {
		goto L73
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
	*(*uint8)(unsafe.Add(mBase, uint32(v396)+14)) = uint8(v432)
	*(*uint16)(unsafe.Add(mBase, uint32(v396)+12)) = uint16(v332)
	F_XLogBeginInsert(m)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L59
	} else {
		goto L87
	}
L72:
	;
	v403 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_gistplacetopage[6])))
	if v403&int32(1) == int32(0) {
		v432 = v393
		goto L71
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(l11)+48))
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408)+118)))
	if v409 != int32(112) {
		v432 = v393
		goto L71
	} else {
		goto L76
	}
L75:
	;
	goto L74
L76:
	;
	if int32(0) < v399 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(l11)+56))
	goto L81
L78:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(l11)+32))
	if v414 != 0 {
		v432 = v393
		goto L71
	} else {
		goto L79
	}
L79:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(l11)+40))
	if v415 == int32(0) {
		goto L77
	} else {
		goto L80
	}
L80:
	;
	v432 = v393
	goto L71
L81:
	;
	if base.Ui32(v419) < base.Ui32(int32(_a_F_gistplacetopage_7)) {
		v432 = int32(1)
		goto L71
	} else {
		goto L82
	}
L82:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(l11)+180))
	if v422 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v432 = int32(0)
	goto L71
L84:
	;
	goto L85
L85:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(l11)+48))
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v427)+119)))
	switch v428 - int32(109) {
	case 0, 5:
		goto L86
	default:
		v432 = int32(0)
		goto L71
	}
L86:
	;
	v431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v422)+112)))
	v432 = v431
	goto L71
L87:
	;
	v439 = int32(8)
	F_XLogRegisterData(m, v396+v439, v439)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L59
	} else {
		goto L88
	}
L88:
	;
	F_XLogRegisterData(m, v29+int32(48), v332<<(uint(int32(1))%32))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L59
	} else {
		goto L89
	}
L89:
	;
	F_XLogRegisterBuffer(m, int32(0), l3, int32(8))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L59
	} else {
		goto L90
	}
L90:
	;
	v454 = F_XLogInsert(m, int32(14), int32(16))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L59
	} else {
		goto L91
	}
L91:
	;
	m.G0 = v396 + int32(16)
	v468 = base.I64_rotl(v454, int64(32))
	goto L63
L92:
	;
	v468 = base.I64_rotl(v461, int64(32))
	goto L63
L93:
	;
	if base.B2i32(base.Ui32(v588+v587) < base.Ui32(v569)) == int32(0) {
		goto L16
	} else {
		goto L106
	}
L94:
	;
	if v7 != 0 {
		goto L103
	} else {
		goto L104
	}
L95:
	;
	if l5 != int32(1) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v517 = int32(0)
	v519 = l1
	v520 = v502
	goto L99
L97:
	;
	v550 = l1
	v551 = v502
	goto L98
L98:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(l4+v551<<(uint(int32(2))%32))))
	v559 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v558)+6)))
	v569 = v559&int32(_a_F_gistplacetopage_2) + v550 + int32(4)
	goto L94
L99:
	;
	v524 = int32(2)
	v526 = l4 + v520<<(uint(v524)%32)
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v526)))
	v528 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v527)+6)))
	v529 = int32(_a_F_gistplacetopage_2)
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v526)+4))
	v533 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v532)+6)))
	v538 = v519 + v528&v529 + v533&v529 + int32(8)
	v540 = v520 + v524
	v542 = v517 + v524
	if v542 != l5&int32(2147483646) {
		v517 = v542
		v519 = v538
		v520 = v540
		goto L99
	} else {
		goto L101
	}
L100:
	;
	if l5&int32(1) == int32(0) {
		v569 = v538
		goto L94
	} else {
		goto L102
	}
L101:
	;
	goto L100
L102:
	;
	v550 = v538
	v551 = v540
	goto L98
L103:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v67+v7<<(uint(int32(2))%32))+20))
	v581 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67+v577&int32(_a_F_gistplacetopage_3))+6)))
	v587 = v581&int32(_a_F_gistplacetopage_2) + int32(4)
	goto L105
L104:
	;
	v587 = int32(0)
	goto L105
L105:
	;
	v588 = F_PageGetFreeSpace(m, v67)
	mBase = m.M
	goto L93
L106:
	;
	goto L33
L107:
	;
	if base.Ui32(int32(2047)) < base.Ui32((v7-int32(1))&int32(_a_F_gistplacetopage_1)) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v647 = int32(0)
	v650 = F_gistjoinvector(m, v621, v29+int32(44), l4, l5)
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L59
	} else {
		goto L112
	}
L109:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v29)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v629 - int32(1)
	if v7 == v629 {
		goto L108
	} else {
		goto L110
	}
L110:
	;
	v636 = (v629 - v7) << (uint(int32(2)) % 32)
	if v636 == int32(0) {
		goto L108
	} else {
		goto L111
	}
L111:
	;
	v641 = v621 + v7<<(uint(int32(2))%32)
	base.MemoryCopy(m, v641-int32(4), v641, v636)
	goto L108
L112:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v29)+44))
	v653 = F_gistSplit(m, l0, v67, v650, v652, l2)
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L59
	} else {
		goto L113
	}
L113:
	;
	if v653 != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v668 = v647
	v669 = v653
	goto L117
L115:
	;
	v697 = v647
	goto L116
L116:
	;
	v712 = v697 + base.B2i32(v49 == int32(0))
	if int32(76) <= v712 {
		goto L8
	} else {
		goto L120
	}
L117:
	;
	v682 = v668 + int32(1)
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v669)+28))
	if v683 != 0 {
		v668 = v682
		v669 = v683
		goto L117
	} else {
		goto L119
	}
L118:
	;
	v697 = v682
	goto L116
L119:
	;
	goto L118
L120:
	;
	v716 = v70 & int32(1)
	if v49 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	if v775 == int32(0) {
		goto L7
	} else {
		goto L134
	}
L122:
	;
	v773 = int32(-1)
	v774 = int64(0)
	v775 = v653
	goto L121
L123:
	;
	goto L124
L124:
	;
	v720 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+16)))
	v721 = v67 + v720
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v721)+8))
	v723 = *(*int64)(unsafe.Add(mBase, uint32(v721)))
	*(*int32)(unsafe.Add(mBase, uint32(v653)+24)) = l3
	if l3 < int32(0) {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v653))) = v743
	if l3 < int32(0) {
		goto L130
	} else {
		goto L131
	}
L126:
	;
	v728 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[0]))
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v728+(l3^int32(-1))*int32(56))+16))
	v743 = v734
	goto L125
L127:
	;
	goto L128
L128:
	;
	v736 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[1]))
	v737 = int32(56)
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v736+l3*v737-v737)+16))
	v743 = v742
	goto L125
L129:
	;
	v765 = F_PageGetTempPageCopySpecial(m, v764)
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L59
	} else {
		goto L133
	}
L130:
	;
	v750 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[2]))
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v750+(l3^int32(-1))<<(uint(int32(2))%32))))
	v764 = v756
	goto L129
L131:
	;
	goto L132
L132:
	;
	v758 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[3]))
	v764 = v758 + l3<<(uint(int32(13))%32) + int32(-8192)
	goto L129
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v653)+20)) = v765
	v768 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v765)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v765+v768)+12)) = uint16(v716)
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v653)+28))
	v773 = v722
	v774 = base.I64_rotl(v723, int64(32))
	v775 = v771
	goto L121
L134:
	;
	v791 = v775
	goto L135
L135:
	;
	v804 = F_gistNewBuffer(m, l0, l11)
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L59
	} else {
		goto L137
	}
L136:
	;
	goto L7
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v791)+24)) = v804
	if v804 < int32(0) {
		goto L140
	} else {
		goto L141
	}
L138:
	;
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v791)+24))
	if v835 < int32(0) {
		goto L144
	} else {
		goto L145
	}
L139:
	;
	F_PageInit(m, v824, int32(_a_F_gistplacetopage_8), int32(16))
	mBase = m.M
	v828 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v824)+16)))
	v829 = v824 + v828
	v830 = int32(_a_F_gistplacetopage_9)
	*(*uint16)(unsafe.Add(mBase, uint32(v829)+14)) = uint16(v830)
	*(*uint16)(unsafe.Add(mBase, uint32(v829)+12)) = uint16(v716)
	*(*int32)(unsafe.Add(mBase, uint32(v829)+8)) = int32(-1)
	goto L138
L140:
	;
	v810 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[2]))
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v810+(v804^int32(-1))<<(uint(int32(2))%32))))
	v824 = v816
	goto L139
L141:
	;
	goto L142
L142:
	;
	v818 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[3]))
	v824 = v818 + v804<<(uint(int32(13))%32) + int32(-8192)
	goto L139
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v791)+20)) = v853
	if v835 < int32(0) {
		goto L148
	} else {
		goto L149
	}
L144:
	;
	v839 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[2]))
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v839+(v835^int32(-1))<<(uint(int32(2))%32))))
	v853 = v845
	goto L143
L145:
	;
	goto L146
L146:
	;
	v847 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[3]))
	v853 = v847 + v835<<(uint(int32(13))%32) + int32(-8192)
	goto L143
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v791))) = v873
	if l3 < int32(0) {
		goto L152
	} else {
		goto L153
	}
L148:
	;
	v858 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[0]))
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v858+(v835^int32(-1))*int32(56))+16))
	v873 = v864
	goto L147
L149:
	;
	goto L150
L150:
	;
	v866 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[1]))
	v867 = int32(56)
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v866+v835*v867-v867)+16))
	v873 = v872
	goto L147
L151:
	;
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v791)+24))
	if v894 < int32(0) {
		goto L156
	} else {
		goto L157
	}
L152:
	;
	v878 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[0]))
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v878+(l3^int32(-1))*int32(56))+16))
	v893 = v884
	goto L151
L153:
	;
	goto L154
L154:
	;
	v886 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[1]))
	v887 = int32(56)
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v886+l3*v887-v887)+16))
	v893 = v892
	goto L151
L155:
	;
	F_PredicateLockPageSplit(m, l0, v893, v913)
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L59
	} else {
		goto L159
	}
L156:
	;
	v898 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[0]))
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v898+(v894^int32(-1))*int32(56))+16))
	v913 = v904
	goto L155
L157:
	;
	goto L158
L158:
	;
	v906 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[1]))
	v907 = int32(56)
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v906+v894*v907-v907)+16))
	v913 = v912
	goto L155
L159:
	;
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v791)+28))
	if v916 != 0 {
		v791 = v916
		goto L135
	} else {
		goto L160
	}
L160:
	;
	goto L136
L161:
	;
	F_MarkBufferDirty(m, l3)
	mBase = m.M
	v988 = m.ExcPending
	if v988 != 0 {
		goto L59
	} else {
		goto L175
	}
L162:
	;
	if l5 == int32(1) {
		goto L165
	} else {
		goto L166
	}
L163:
	;
	goto L164
L164:
	;
	F_gistfillbuffer(m, v67, l4, l5, int32(0))
	mBase = m.M
	v985 = m.ExcPending
	if v985 != 0 {
		goto L59
	} else {
		goto L174
	}
L165:
	;
	v957 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v958 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v957)+6)))
	v961 = F_PageIndexTupleOverwrite(m, v67, v7, v957, v958&int32(_a_F_gistplacetopage_2))
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L59
	} else {
		goto L168
	}
L166:
	;
	goto L167
L167:
	;
	F_PageIndexTupleDelete(m, v67, v7)
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L59
	} else {
		goto L173
	}
L168:
	;
	if v961 != 0 {
		goto L161
	} else {
		goto L169
	}
L169:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L59
	} else {
		goto L170
	}
L170:
	;
	v967 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v967 + int32(4)
	F_errmsg_internal(m, int32(_a_F_gistplacetopage_10), v29+int32(32))
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L59
	} else {
		goto L171
	}
L171:
	;
	F_errfinish(m, int32(_a_F_gistplacetopage_11), int32(557), int32(_a_F_gistplacetopage_12))
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L59
	} else {
		goto L172
	}
L172:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L173:
	;
	goto L164
L174:
	;
	goto L161
L175:
	;
	if l8 != 0 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	F_MarkBufferDirty(m, l8)
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L59
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	if l12 != 0 {
		v1016 = int64(1)
		goto L180
	} else {
		goto L181
	}
L179:
	;
	goto L178
L180:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v67))) = base.I64_rotl(v1016, int64(32))
	v1020 = int32(0)
	if l7 == v1020 {
		v1985 = v1020
		v1995 = v1016
		goto L6
	} else {
		goto L194
	}
L181:
	;
	v992 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v993 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v992)+118)))
	if v993 != int32(112) {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v1014 = F_XLogGetFakeLSN(m, l0)
	mBase = m.M
	v1015 = m.ExcPending
	if v1015 != 0 {
		goto L59
	} else {
		goto L193
	}
L183:
	;
	v997 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[5]))
	if v997 <= int32(0) {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1000 != 0 {
		goto L182
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	if base.Ui32(v950&int32(_a_F_gistplacetopage_1)) <= base.Ui32(int32(2047)) {
		goto L189
	} else {
		goto L190
	}
L187:
	;
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v1001 != 0 {
		goto L182
	} else {
		goto L188
	}
L188:
	;
	goto L186
L189:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+48)) = uint16(v7)
	v1011 = int32(1)
	goto L191
L190:
	;
	v1011 = int32(0)
	goto L191
L191:
	;
	v1012 = F_gistXLogUpdate(m, l3, v29+int32(48), v1011, l4, l5, l8)
	mBase = m.M
	v1013 = m.ExcPending
	if v1013 != 0 {
		goto L59
	} else {
		goto L192
	}
L192:
	;
	v1016 = v1012
	goto L180
L193:
	;
	v1016 = v1014
	goto L180
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v49
	v1985 = v1020
	v1995 = v1016
	goto L6
L195:
	;
	F_errmsg_internal(m, int32(_a_F_gistplacetopage_13), int32(0))
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L59
	} else {
		goto L196
	}
L196:
	;
	F_errfinish(m, int32(_a_F_gistplacetopage_11), int32(257), int32(_a_F_gistplacetopage_12))
	mBase = m.M
	v1036 = m.ExcPending
	if v1036 != 0 {
		goto L59
	} else {
		goto L197
	}
L197:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = int32(75)
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v712
	F_errmsg_internal(m, int32(_a_F_gistplacetopage_14), v29)
	mBase = m.M
	v1046 = m.ExcPending
	if v1046 != 0 {
		goto L59
	} else {
		goto L199
	}
L199:
	;
	F_errfinish(m, int32(_a_F_gistplacetopage_11), int32(329), int32(_a_F_gistplacetopage_12))
	mBase = m.M
	v1051 = m.ExcPending
	if v1051 != 0 {
		goto L59
	} else {
		goto L200
	}
L200:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L201:
	;
	if l12 != 0 {
		goto L267
	} else {
		goto L268
	}
L202:
	;
	v1349 = v1326
	goto L235
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+72)) = l3
	if l3 < int32(0) {
		goto L218
	} else {
		goto L219
	}
L204:
	;
	v1091 = v653
	goto L207
L205:
	;
	goto L206
L206:
	;
	if v49 == int32(0) {
		goto L203
	} else {
		goto L216
	}
L207:
	;
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(v1091)+16))
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(v1091)))
	*(*int32)(unsafe.Add(mBase, uint32(v1104))) = base.I32_rotr(v1105, int32(16))
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(v1091)+16))
	v1110 = int32(_a_F_gistplacetopage_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1109)+4)) = uint16(v1110)
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v1091)+28))
	if v1112 != 0 {
		v1091 = v1112
		goto L207
	} else {
		goto L209
	}
L208:
	;
	if v49 == int32(0) {
		goto L203
	} else {
		goto L210
	}
L209:
	;
	goto L208
L210:
	;
	v1128 = v653
	goto L211
L211:
	;
	v1142 = F_palloc(m, int32(8))
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		goto L59
	} else {
		goto L213
	}
L212:
	;
	v1326 = v653
	goto L202
L213:
	;
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v1128)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1142))) = v1144
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v1128)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1142)+4)) = v1146
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(l9)))
	v1149 = F_lappend(m, v1148, v1142)
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L59
	} else {
		goto L214
	}
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l9))) = v1149
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v1128)+28))
	if v1152 != 0 {
		v1128 = v1152
		goto L211
	} else {
		goto L215
	}
L215:
	;
	goto L212
L216:
	;
	v1514 = int32(0)
	v1522 = int32(1)
	goto L201
L217:
	;
	v1202 = F_PageGetTempPageCopySpecial(m, v1201)
	mBase = m.M
	v1203 = m.ExcPending
	if v1203 != 0 {
		goto L59
	} else {
		goto L221
	}
L218:
	;
	v1187 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[2]))
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v1187+(l3^int32(-1))<<(uint(int32(2))%32))))
	v1201 = v1193
	goto L217
L219:
	;
	goto L220
L220:
	;
	v1195 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[3]))
	v1201 = v1195 + l3<<(uint(int32(13))%32) + int32(-8192)
	goto L217
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+68)) = v1202
	v1205 = int32(0)
	v1206 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1202)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1202+v1206)+12)) = uint16(v1205)
	if v653 != 0 {
		goto L223
	} else {
		goto L224
	}
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v1282
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = int32(0)
	v1312 = F_gistfillitupvec(m, v1284, v1282, v29+int32(60))
	mBase = m.M
	v1313 = m.ExcPending
	if v1313 != 0 {
		goto L59
	} else {
		goto L234
	}
L223:
	;
	v1211 = v1205
	v1223 = v653
	goto L226
L224:
	;
	goto L225
L225:
	;
	v1279 = F_palloc_mul(m, int32(4), int32(0))
	mBase = m.M
	v1280 = m.ExcPending
	if v1280 != 0 {
		goto L59
	} else {
		goto L233
	}
L226:
	;
	v1237 = v1211 + int32(1)
	v1238 = *(*int32)(unsafe.Add(mBase, uint32(v1223)+28))
	if v1238 != 0 {
		v1211 = v1237
		v1223 = v1238
		goto L226
	} else {
		goto L228
	}
L227:
	;
	v1241 = F_palloc_mul(m, int32(4), v1237)
	mBase = m.M
	v1242 = m.ExcPending
	if v1242 != 0 {
		goto L59
	} else {
		goto L229
	}
L228:
	;
	goto L227
L229:
	;
	v1256 = int32(0)
	v1257 = v653
	goto L230
L230:
	;
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(v1257)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1241+v1256<<(uint(int32(2))%32)))) = v1272
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(v1257)+28))
	if v1276 != 0 {
		v1256 = v1256 + int32(1)
		v1257 = v1276
		goto L230
	} else {
		goto L232
	}
L231:
	;
	v1282 = v1237
	v1284 = v1241
	goto L222
L232:
	;
	goto L231
L233:
	;
	v1282 = v1205
	v1284 = v1279
	goto L222
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+76)) = v653
	*(*int32)(unsafe.Add(mBase, uint32(v29)+64)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v1312
	v1326 = v29 + int32(48)
	goto L202
L235:
	;
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(v1349)+4))
	if int32(0) < v1374 {
		goto L237
	} else {
		goto L238
	}
L236:
	;
	v1514 = v1326
	v1522 = v1486
	goto L201
L237:
	;
	v1377 = *(*int32)(unsafe.Add(mBase, uint32(v1349)+8))
	v1392 = v1377
	v1393 = int32(0)
	goto L240
L238:
	;
	goto L239
L239:
	;
	v1474 = *(*int32)(unsafe.Add(mBase, uint32(v1349)+28))
	if v1474 == int32(0) {
		v1481 = v773
		goto L254
	} else {
		goto L255
	}
L240:
	;
	v1405 = *(*int32)(unsafe.Add(mBase, uint32(v1349)+20))
	v1406 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1392)+6)))
	v1410 = v1393 + int32(1)
	v1414 = F_PageAddItemExtended(m, v1405, v1392, v1406&int32(_a_F_gistplacetopage_2), v1410&int32(_a_F_gistplacetopage_1), int32(0))
	mBase = m.M
	v1415 = m.ExcPending
	if v1415 != 0 {
		goto L59
	} else {
		goto L242
	}
L241:
	;
	goto L239
L242:
	;
	if v1414 == int32(0) {
		goto L5
	} else {
		goto L243
	}
L243:
	;
	if l7 == int32(0) {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v1442 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1392)+6)))
	v1446 = *(*int32)(unsafe.Add(mBase, uint32(v1349)+4))
	if v1410 < v1446 {
		v1392 = v1392 + v1442&int32(_a_F_gistplacetopage_2)
		v1393 = v1410
		goto L240
	} else {
		goto L253
	}
L245:
	;
	v1420 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v1421 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1392)+2)))
	v1422 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1392))))
	v1423 = int32(16)
	v1426 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1420)+2)))
	v1427 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1420))))
	if v1421|v1422<<(uint(v1423)%32) == v1426|v1427<<(uint(v1423)%32) {
		goto L248
	} else {
		goto L249
	}
L246:
	;
	if v1437 == int32(0) {
		goto L244
	} else {
		goto L252
	}
L247:
	;
	goto L246
L248:
	;
	v1433 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1392)+4)))
	v1434 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1420)+4)))
	if v1433 == v1434 {
		v1437 = int32(1)
		goto L247
	} else {
		goto L251
	}
L249:
	;
	goto L250
L250:
	;
	v1437 = int32(0)
	goto L247
L251:
	;
	goto L250
L252:
	;
	v1440 = *(*int32)(unsafe.Add(mBase, uint32(v1349)))
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v1440
	goto L244
L253:
	;
	goto L241
L254:
	;
	v1482 = *(*int32)(unsafe.Add(mBase, uint32(v1349)+20))
	v1483 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1482)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v1482+v1483)+8)) = v1481
	v1486 = int32(0)
	v1487 = *(*int32)(unsafe.Add(mBase, uint32(v1349)+20))
	v1488 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1487)+16)))
	v1489 = v1487 + v1488
	v1490 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1489)+12)))
	if v49 != 0 {
		goto L257
	} else {
		goto L258
	}
L255:
	;
	v1477 = *(*int32)(unsafe.Add(mBase, uint32(v1349)))
	if v1477 == int32(0) {
		v1481 = v773
		goto L254
	} else {
		goto L256
	}
L256:
	;
	v1480 = *(*int32)(unsafe.Add(mBase, uint32(v1474)))
	v1481 = v1480
	goto L254
L257:
	;
	v1495 = int32(8)
	goto L259
L258:
	;
	v1495 = v1486
	goto L259
L259:
	;
	v1497 = *(*int32)(unsafe.Add(mBase, uint32(v1349)+28))
	if v1497 != 0 {
		goto L260
	} else {
		goto L261
	}
L260:
	;
	v1498 = v1495
	goto L262
L261:
	;
	v1498 = int32(0)
	goto L262
L262:
	;
	if v11 != 0 {
		goto L263
	} else {
		goto L264
	}
L263:
	;
	v1500 = v1498
	goto L265
L264:
	;
	v1500 = int32(0)
	goto L265
L265:
	;
	v1501 = v1490&int32(_a_F_gistplacetopage_15) | v1500
	*(*uint16)(unsafe.Add(mBase, uint32(v1489)+12)) = uint16(v1501)
	v1503 = *(*int32)(unsafe.Add(mBase, uint32(v1349)+20))
	v1504 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1503)+16)))
	*(*int64)(unsafe.Add(mBase, uint32(v1503+v1504))) = base.I64_rotl(v774, int64(32))
	v1507 = *(*int32)(unsafe.Add(mBase, uint32(v1349)+28))
	if v1507 != 0 {
		v1349 = v1507
		goto L235
	} else {
		goto L266
	}
L266:
	;
	goto L236
L267:
	;
	v1550 = int32(_a_F_gistplacetopage_0)
	v1552 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[4])) = v1552 + int32(1)
	if v1522 == int32(0) {
		goto L276
	} else {
		goto L277
	}
L268:
	;
	v1534 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1534)+118)))
	if v1535 != int32(112) {
		goto L267
	} else {
		goto L269
	}
L269:
	;
	v1539 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[5]))
	if v1539 <= int32(0) {
		goto L270
	} else {
		goto L271
	}
L270:
	;
	v1542 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1542 != 0 {
		goto L267
	} else {
		goto L273
	}
L271:
	;
	goto L272
L272:
	;
	v1544 = int32(1)
	F_XLogEnsureRecordSpace(m, v712, v712<<(uint(v1544)%32)|v1544)
	mBase = m.M
	v1549 = m.ExcPending
	if v1549 != 0 {
		goto L59
	} else {
		goto L275
	}
L273:
	;
	v1543 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v1543 != 0 {
		goto L267
	} else {
		goto L274
	}
L274:
	;
	goto L272
L275:
	;
	goto L267
L276:
	;
	v1571 = v1514
	goto L279
L277:
	;
	goto L278
L278:
	;
	if l8 != 0 {
		goto L283
	} else {
		goto L284
	}
L279:
	;
	v1584 = *(*int32)(unsafe.Add(mBase, uint32(v1571)+24))
	F_MarkBufferDirty(m, v1584)
	mBase = m.M
	v1586 = m.ExcPending
	if v1586 != 0 {
		goto L59
	} else {
		goto L281
	}
L280:
	;
	goto L278
L281:
	;
	v1587 = *(*int32)(unsafe.Add(mBase, uint32(v1571)+28))
	if v1587 != 0 {
		v1571 = v1587
		goto L279
	} else {
		goto L282
	}
L282:
	;
	goto L280
L283:
	;
	F_MarkBufferDirty(m, l8)
	mBase = m.M
	v1615 = m.ExcPending
	if v1615 != 0 {
		goto L59
	} else {
		goto L286
	}
L284:
	;
	goto L285
L285:
	;
	v1616 = *(*int32)(unsafe.Add(mBase, uint32(v1514)+20))
	v1617 = *(*int32)(unsafe.Add(mBase, uint32(v1514)+24))
	if v1617 < int32(0) {
		goto L288
	} else {
		goto L289
	}
L286:
	;
	goto L285
L287:
	;
	F_PageRestoreTempPage(m, v1616, v1635)
	mBase = m.M
	v1637 = m.ExcPending
	if v1637 != 0 {
		goto L59
	} else {
		goto L291
	}
L288:
	;
	v1621 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[2]))
	v1627 = *(*int32)(unsafe.Add(mBase, uint32(v1621+(v1617^int32(-1))<<(uint(int32(2))%32))))
	v1635 = v1627
	goto L287
L289:
	;
	goto L290
L290:
	;
	v1629 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[3]))
	v1635 = v1629 + v1617<<(uint(int32(13))%32) + int32(-8192)
	goto L287
L291:
	;
	v1638 = *(*int32)(unsafe.Add(mBase, uint32(v1514)+24))
	if v1638 < int32(0) {
		goto L293
	} else {
		goto L294
	}
L292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1514)+20)) = v1656
	if l12 != 0 {
		v1851 = int64(1)
		goto L296
	} else {
		goto L297
	}
L293:
	;
	v1642 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[2]))
	v1648 = *(*int32)(unsafe.Add(mBase, uint32(v1642+(v1638^int32(-1))<<(uint(int32(2))%32))))
	v1656 = v1648
	goto L292
L294:
	;
	goto L295
L295:
	;
	v1650 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[3]))
	v1656 = v1650 + v1638<<(uint(int32(13))%32) + int32(-8192)
	goto L292
L296:
	;
	if v1522 == int32(0) {
		goto L328
	} else {
		goto L329
	}
L297:
	;
	v1659 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1659)+118)))
	if v1660 != int32(112) {
		goto L298
	} else {
		goto L299
	}
L298:
	;
	v1823 = F_XLogGetFakeLSN(m, l0)
	mBase = m.M
	v1824 = m.ExcPending
	if v1824 != 0 {
		goto L59
	} else {
		goto L327
	}
L299:
	;
	v1664 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[5]))
	if v1664 <= int32(0) {
		goto L300
	} else {
		goto L301
	}
L300:
	;
	v1667 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1667 != 0 {
		goto L298
	} else {
		goto L303
	}
L301:
	;
	goto L302
L302:
	;
	v1669 = int32(0)
	v1670 = m.G0
	v1672 = v1670 - int32(32)
	m.G0 = v1672
	if v1514 != 0 {
		goto L305
	} else {
		goto L306
	}
L303:
	;
	v1668 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v1668 != 0 {
		goto L298
	} else {
		goto L304
	}
L304:
	;
	goto L302
L305:
	;
	v1675 = v1514
	v1676 = v1669
	goto L308
L306:
	;
	v1705 = v1669
	goto L307
L307:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1672)+28)) = uint8(v11)
	*(*uint16)(unsafe.Add(mBase, uint32(v1672)+26)) = uint16(v1705)
	*(*uint8)(unsafe.Add(mBase, uint32(v1672)+24)) = uint8(v716)
	*(*int64)(unsafe.Add(mBase, uint32(v1672)+16)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v1672)+8)) = v773
	F_XLogBeginInsert(m)
	mBase = m.M
	v1735 = m.ExcPending
	if v1735 != 0 {
		goto L59
	} else {
		goto L311
	}
L308:
	;
	v1701 = v1676 + int32(1)
	v1702 = *(*int32)(unsafe.Add(mBase, uint32(v1675)+28))
	if v1702 != 0 {
		v1675 = v1702
		v1676 = v1701
		goto L308
	} else {
		goto L310
	}
L309:
	;
	v1705 = v1701
	goto L307
L310:
	;
	goto L309
L311:
	;
	if l8 != 0 {
		goto L312
	} else {
		goto L313
	}
L312:
	;
	F_XLogRegisterBuffer(m, int32(0), l8, int32(8))
	mBase = m.M
	v1739 = m.ExcPending
	if v1739 != 0 {
		goto L59
	} else {
		goto L315
	}
L313:
	;
	goto L314
L314:
	;
	F_XLogRegisterData(m, v1672+int32(8), int32(24))
	mBase = m.M
	v1744 = m.ExcPending
	if v1744 != 0 {
		goto L59
	} else {
		goto L316
	}
L315:
	;
	goto L314
L316:
	;
	if v1514 != 0 {
		goto L317
	} else {
		goto L318
	}
L317:
	;
	v1746 = v1514
	v1748 = int32(1)
	goto L320
L318:
	;
	goto L319
L319:
	;
	v1818 = F_XLogInsert(m, int32(14), int32(48))
	mBase = m.M
	v1819 = m.ExcPending
	if v1819 != 0 {
		goto L59
	} else {
		goto L326
	}
L320:
	;
	v1773 = v1748 & int32(255)
	v1774 = *(*int32)(unsafe.Add(mBase, uint32(v1746)+24))
	F_XLogRegisterBuffer(m, v1773, v1774, int32(6))
	mBase = m.M
	v1777 = m.ExcPending
	if v1777 != 0 {
		goto L59
	} else {
		goto L322
	}
L321:
	;
	goto L319
L322:
	;
	v1778 = int32(4)
	F_XLogRegisterBufData(m, v1773, v1746+v1778, v1778)
	mBase = m.M
	v1782 = m.ExcPending
	if v1782 != 0 {
		goto L59
	} else {
		goto L323
	}
L323:
	;
	v1783 = *(*int32)(unsafe.Add(mBase, uint32(v1746)+8))
	v1784 = *(*int32)(unsafe.Add(mBase, uint32(v1746)+12))
	F_XLogRegisterBufData(m, v1773, v1783, v1784)
	mBase = m.M
	v1786 = m.ExcPending
	if v1786 != 0 {
		goto L59
	} else {
		goto L324
	}
L324:
	;
	v1789 = *(*int32)(unsafe.Add(mBase, uint32(v1746)+28))
	if v1789 != 0 {
		v1746 = v1789
		v1748 = v1748 + int32(1)
		goto L320
	} else {
		goto L325
	}
L325:
	;
	goto L321
L326:
	;
	m.G0 = v1672 + int32(32)
	v1851 = v1818
	goto L296
L327:
	;
	v1851 = v1823
	goto L296
L328:
	;
	v1869 = v1514
	goto L331
L329:
	;
	goto L330
L330:
	;
	if v49 != 0 {
		goto L334
	} else {
		goto L335
	}
L331:
	;
	v1882 = *(*int32)(unsafe.Add(mBase, uint32(v1869)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v1882))) = base.I64_rotl(v1851, int64(32))
	v1884 = *(*int32)(unsafe.Add(mBase, uint32(v1869)+28))
	if v1884 != 0 {
		v1869 = v1884
		goto L331
	} else {
		goto L333
	}
L332:
	;
	goto L330
L333:
	;
	goto L332
L334:
	;
	v1985 = int32(1)
	v1995 = v1851
	goto L6
L335:
	;
	v1911 = *(*int32)(unsafe.Add(mBase, uint32(v1514)+28))
	if v1911 == int32(0) {
		goto L334
	} else {
		goto L336
	}
L336:
	;
	v1927 = v1911
	goto L337
L337:
	;
	v1940 = *(*int32)(unsafe.Add(mBase, uint32(v1927)+24))
	F_UnlockReleaseBuffer(m, v1940)
	mBase = m.M
	v1942 = m.ExcPending
	if v1942 != 0 {
		goto L59
	} else {
		goto L339
	}
L338:
	;
	goto L334
L339:
	;
	v1943 = *(*int32)(unsafe.Add(mBase, uint32(v1927)+28))
	if v1943 != 0 {
		v1927 = v1943
		goto L337
	} else {
		goto L340
	}
L340:
	;
	goto L338
L341:
	;
	if l8 < int32(0) {
		goto L345
	} else {
		goto L346
	}
L342:
	;
	goto L343
L343:
	;
	v2030 = int32(_a_F_gistplacetopage_0)
	v2032 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[4])) = v2032 - int32(1)
	m.G0 = v29 + int32(864)
	return v1985
L344:
	;
	v2015 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2014)+16)))
	v2018 = base.I64_rotl(v1995, int64(32))
	*(*int64)(unsafe.Add(mBase, uint32(v2015+v2014))) = v2018
	v2020 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2014)+16)))
	v2021 = v2014 + v2020
	v2022 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2021)+12)))
	v2024 = v2022 & int32(_a_F_gistplacetopage_15)
	*(*uint16)(unsafe.Add(mBase, uint32(v2021)+12)) = uint16(v2024)
	*(*int64)(unsafe.Add(mBase, uint32(v2014))) = v2018
	goto L343
L345:
	;
	v2000 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[2]))
	v2006 = *(*int32)(unsafe.Add(mBase, uint32(v2000+(l8^int32(-1))<<(uint(int32(2))%32))))
	v2014 = v2006
	goto L344
L346:
	;
	goto L347
L347:
	;
	v2008 = *(*int32)(unsafe.Add(mBase, _c_F_gistplacetopage[3]))
	v2014 = v2008 + l8<<(uint(int32(13))%32) + int32(-8192)
	goto L344
L348:
	;
	v2044 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v2044 + int32(4)
	F_errmsg_internal(m, int32(_a_F_gistplacetopage_10), v29+int32(16))
	mBase = m.M
	v2052 = m.ExcPending
	if v2052 != 0 {
		goto L59
	} else {
		goto L349
	}
L349:
	;
	F_errfinish(m, int32(_a_F_gistplacetopage_11), int32(435), int32(_a_F_gistplacetopage_12))
	mBase = m.M
	v2057 = m.ExcPending
	if v2057 != 0 {
		goto L59
	} else {
		goto L350
	}
L350:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_globalChannelTableHash(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
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
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
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
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	v4 = int32(4)
	v5 = l0 + v4
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = int32(711645284)
	v14 = v6 - int32(1636608428) ^ v11 - int32(1455628627)
	v19 = v14 ^ int32(-1636608428) - base.I32_rotl(v14, int32(25))
	v24 = v19 ^ v11 - base.I32_rotl(v19, int32(16))
	v28 = v24 ^ v14 - base.I32_rotl(v24, v4)
	v32 = v28 ^ v19 - base.I32_rotl(v28, int32(14))
	v37 = int32(64)
	v40 = F_memchr(m, v5, int32(0), v37)
	mBase = m.M
	if v40 != 0 {
		v42 = v40 - v5
	} else {
		v42 = v37
	}
	v48 = v42 - int32(1636608432)
	if v5&int32(3) != 0 {
		if base.Ui32(int32(11)) < base.Ui32(v42) {
			v157 = v5
			v158 = v42
			v159 = v48
			v160 = v48
			v161 = v48
			for {
				v163 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
				v164 = v163 + v160
				v165 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
				v167 = *(*int32)(unsafe.Add(mBase, uint32(v157)+8))
				v168 = v167 + v161
				v170 = int32(4)
				v172 = v165 + v159 - v168 ^ base.I32_rotl(v168, v170)
				v176 = v164 - v172 ^ base.I32_rotl(v172, int32(6))
				v177 = v168 + v164
				v178 = v172 + v177
				v179 = v176 + v178
				v183 = v177 - v176 ^ base.I32_rotl(v176, int32(8))
				v187 = v178 - v183 ^ base.I32_rotl(v183, int32(16))
				v191 = v179 - v187 ^ base.I32_rotl(v187, int32(19))
				v192 = v183 + v179
				v193 = v187 + v192
				v194 = v191 + v193
				v198 = v192 - v191 ^ base.I32_rotl(v191, v170)
				v199 = int32(12)
				v200 = v157 + v199
				v202 = v158 - v199
				if base.Ui32(int32(11)) < base.Ui32(v202) {
					v157 = v200
					v158 = v202
					v159 = v193
					v160 = v194
					v161 = v198
					continue
				} else {
					break
				}
				break
			}
			v205 = v200
			v206 = v202
			v207 = v193
			v208 = v194
			v209 = v198
		} else {
			v205 = v5
			v206 = v42
			v207 = v48
			v208 = v48
			v209 = v48
		}
		switch v206 - int32(1) {
		case 0:
			v268 = v207
			v269 = v208
			v270 = v209
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
			v275 = v268 + v271
			v276 = v269
			v277 = v270
		case 1:
			v261 = v207
			v262 = v208
			v263 = v209
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+1)))
			v268 = v264<<(uint(int32(8))%32) + v261
			v269 = v262
			v270 = v263
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
			v275 = v268 + v271
			v276 = v269
			v277 = v270
		case 2:
			v254 = v207
			v255 = v208
			v256 = v209
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+2)))
			v261 = v257<<(uint(int32(16))%32) + v254
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+1)))
			v268 = v264<<(uint(int32(8))%32) + v261
			v269 = v262
			v270 = v263
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
			v275 = v268 + v271
			v276 = v269
			v277 = v270
		case 3:
			v248 = v208
			v249 = v209
			v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+3)))
			v254 = v250<<(uint(int32(24))%32) + v207
			v255 = v248
			v256 = v249
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+2)))
			v261 = v257<<(uint(int32(16))%32) + v254
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+1)))
			v268 = v264<<(uint(int32(8))%32) + v261
			v269 = v262
			v270 = v263
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
			v275 = v268 + v271
			v276 = v269
			v277 = v270
		case 4:
			v244 = v208
			v245 = v209
			v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+4)))
			v248 = v244 + v246
			v249 = v245
			v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+3)))
			v254 = v250<<(uint(int32(24))%32) + v207
			v255 = v248
			v256 = v249
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+2)))
			v261 = v257<<(uint(int32(16))%32) + v254
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+1)))
			v268 = v264<<(uint(int32(8))%32) + v261
			v269 = v262
			v270 = v263
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
			v275 = v268 + v271
			v276 = v269
			v277 = v270
		case 5:
			v238 = v208
			v239 = v209
			v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+5)))
			v244 = v240<<(uint(int32(8))%32) + v238
			v245 = v239
			v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+4)))
			v248 = v244 + v246
			v249 = v245
			v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+3)))
			v254 = v250<<(uint(int32(24))%32) + v207
			v255 = v248
			v256 = v249
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+2)))
			v261 = v257<<(uint(int32(16))%32) + v254
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+1)))
			v268 = v264<<(uint(int32(8))%32) + v261
			v269 = v262
			v270 = v263
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
			v275 = v268 + v271
			v276 = v269
			v277 = v270
		case 6:
			v232 = v208
			v233 = v209
			v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+6)))
			v238 = v234<<(uint(int32(16))%32) + v232
			v239 = v233
			v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+5)))
			v244 = v240<<(uint(int32(8))%32) + v238
			v245 = v239
			v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+4)))
			v248 = v244 + v246
			v249 = v245
			v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+3)))
			v254 = v250<<(uint(int32(24))%32) + v207
			v255 = v248
			v256 = v249
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+2)))
			v261 = v257<<(uint(int32(16))%32) + v254
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+1)))
			v268 = v264<<(uint(int32(8))%32) + v261
			v269 = v262
			v270 = v263
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
			v275 = v268 + v271
			v276 = v269
			v277 = v270
		case 7:
			v227 = v209
			v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+7)))
			v232 = v228<<(uint(int32(24))%32) + v208
			v233 = v227
			v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+6)))
			v238 = v234<<(uint(int32(16))%32) + v232
			v239 = v233
			v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+5)))
			v244 = v240<<(uint(int32(8))%32) + v238
			v245 = v239
			v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+4)))
			v248 = v244 + v246
			v249 = v245
			v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+3)))
			v254 = v250<<(uint(int32(24))%32) + v207
			v255 = v248
			v256 = v249
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+2)))
			v261 = v257<<(uint(int32(16))%32) + v254
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+1)))
			v268 = v264<<(uint(int32(8))%32) + v261
			v269 = v262
			v270 = v263
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
			v275 = v268 + v271
			v276 = v269
			v277 = v270
		case 8:
			v222 = v209
			v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+8)))
			v227 = v223<<(uint(int32(8))%32) + v222
			v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+7)))
			v232 = v228<<(uint(int32(24))%32) + v208
			v233 = v227
			v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+6)))
			v238 = v234<<(uint(int32(16))%32) + v232
			v239 = v233
			v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+5)))
			v244 = v240<<(uint(int32(8))%32) + v238
			v245 = v239
			v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+4)))
			v248 = v244 + v246
			v249 = v245
			v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+3)))
			v254 = v250<<(uint(int32(24))%32) + v207
			v255 = v248
			v256 = v249
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+2)))
			v261 = v257<<(uint(int32(16))%32) + v254
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+1)))
			v268 = v264<<(uint(int32(8))%32) + v261
			v269 = v262
			v270 = v263
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
			v275 = v268 + v271
			v276 = v269
			v277 = v270
		case 9:
			v217 = v209
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+9)))
			v222 = v218<<(uint(int32(16))%32) + v217
			v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+8)))
			v227 = v223<<(uint(int32(8))%32) + v222
			v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+7)))
			v232 = v228<<(uint(int32(24))%32) + v208
			v233 = v227
			v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+6)))
			v238 = v234<<(uint(int32(16))%32) + v232
			v239 = v233
			v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+5)))
			v244 = v240<<(uint(int32(8))%32) + v238
			v245 = v239
			v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+4)))
			v248 = v244 + v246
			v249 = v245
			v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+3)))
			v254 = v250<<(uint(int32(24))%32) + v207
			v255 = v248
			v256 = v249
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+2)))
			v261 = v257<<(uint(int32(16))%32) + v254
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+1)))
			v268 = v264<<(uint(int32(8))%32) + v261
			v269 = v262
			v270 = v263
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
			v275 = v268 + v271
			v276 = v269
			v277 = v270
		case 10:
			v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+10)))
			v217 = v213<<(uint(int32(24))%32) + v209
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+9)))
			v222 = v218<<(uint(int32(16))%32) + v217
			v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+8)))
			v227 = v223<<(uint(int32(8))%32) + v222
			v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+7)))
			v232 = v228<<(uint(int32(24))%32) + v208
			v233 = v227
			v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+6)))
			v238 = v234<<(uint(int32(16))%32) + v232
			v239 = v233
			v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+5)))
			v244 = v240<<(uint(int32(8))%32) + v238
			v245 = v239
			v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+4)))
			v248 = v244 + v246
			v249 = v245
			v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+3)))
			v254 = v250<<(uint(int32(24))%32) + v207
			v255 = v248
			v256 = v249
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+2)))
			v261 = v257<<(uint(int32(16))%32) + v254
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+1)))
			v268 = v264<<(uint(int32(8))%32) + v261
			v269 = v262
			v270 = v263
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
			v275 = v268 + v271
			v276 = v269
			v277 = v270
		default:
			v275 = v207
			v276 = v208
			v277 = v209
		}
	} else {
		if base.Ui32(v42) < base.Ui32(int32(12)) {
			v103 = v5
			v104 = v42
			v105 = v48
			v106 = v48
			v107 = v48
		} else {
			v55 = v5
			v56 = v42
			v57 = v48
			v58 = v48
			v59 = v48
			for {
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
				v62 = v61 + v58
				v63 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
				v66 = v65 + v59
				v68 = int32(4)
				v70 = v63 + v57 - v66 ^ base.I32_rotl(v66, v68)
				v74 = v62 - v70 ^ base.I32_rotl(v70, int32(6))
				v75 = v66 + v62
				v76 = v70 + v75
				v77 = v74 + v76
				v81 = v75 - v74 ^ base.I32_rotl(v74, int32(8))
				v85 = v76 - v81 ^ base.I32_rotl(v81, int32(16))
				v89 = v77 - v85 ^ base.I32_rotl(v85, int32(19))
				v90 = v81 + v77
				v91 = v85 + v90
				v92 = v89 + v91
				v96 = v90 - v89 ^ base.I32_rotl(v89, v68)
				v97 = int32(12)
				v98 = v55 + v97
				v100 = v56 - v97
				if base.Ui32(int32(11)) < base.Ui32(v100) {
					v55 = v98
					v56 = v100
					v57 = v91
					v58 = v92
					v59 = v96
					continue
				} else {
					break
				}
				break
			}
			v103 = v98
			v104 = v100
			v105 = v91
			v106 = v92
			v107 = v96
		}
		switch v104 - int32(1) {
		case 0:
			v154 = v105
			v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
			v275 = v154 + v155
			v276 = v106
			v277 = v107
		case 1:
			v149 = v105
			v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+1)))
			v154 = v150<<(uint(int32(8))%32) + v149
			v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
			v275 = v154 + v155
			v276 = v106
			v277 = v107
		case 2:
			v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+2)))
			v149 = v145<<(uint(int32(16))%32) + v105
			v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+1)))
			v154 = v150<<(uint(int32(8))%32) + v149
			v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
			v275 = v154 + v155
			v276 = v106
			v277 = v107
		case 3:
			v142 = v106
			v143 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
			v275 = v143 + v105
			v276 = v142
			v277 = v107
		case 4:
			v139 = v106
			v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+4)))
			v142 = v139 + v140
			v143 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
			v275 = v143 + v105
			v276 = v142
			v277 = v107
		case 5:
			v134 = v106
			v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+5)))
			v139 = v135<<(uint(int32(8))%32) + v134
			v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+4)))
			v142 = v139 + v140
			v143 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
			v275 = v143 + v105
			v276 = v142
			v277 = v107
		case 6:
			v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+6)))
			v134 = v130<<(uint(int32(16))%32) + v106
			v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+5)))
			v139 = v135<<(uint(int32(8))%32) + v134
			v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+4)))
			v142 = v139 + v140
			v143 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
			v275 = v143 + v105
			v276 = v142
			v277 = v107
		case 7:
			v125 = v107
			v126 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
			v128 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
			v275 = v126 + v105
			v276 = v128 + v106
			v277 = v125
		case 8:
			v120 = v107
			v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+8)))
			v125 = v121<<(uint(int32(8))%32) + v120
			v126 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
			v128 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
			v275 = v126 + v105
			v276 = v128 + v106
			v277 = v125
		case 9:
			v115 = v107
			v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+9)))
			v120 = v116<<(uint(int32(16))%32) + v115
			v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+8)))
			v125 = v121<<(uint(int32(8))%32) + v120
			v126 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
			v128 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
			v275 = v126 + v105
			v276 = v128 + v106
			v277 = v125
		case 10:
			v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+10)))
			v115 = v111<<(uint(int32(24))%32) + v107
			v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+9)))
			v120 = v116<<(uint(int32(16))%32) + v115
			v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+8)))
			v125 = v121<<(uint(int32(8))%32) + v120
			v126 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
			v128 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
			v275 = v126 + v105
			v276 = v128 + v106
			v277 = v125
		default:
			v275 = v105
			v276 = v106
			v277 = v107
		}
	}
	v280 = int32(14)
	v282 = v276 ^ v277 - base.I32_rotl(v276, v280)
	v286 = v282 ^ v275 - base.I32_rotl(v282, int32(11))
	v290 = v286 ^ v276 - base.I32_rotl(v286, int32(25))
	v294 = v290 ^ v282 - base.I32_rotl(v290, int32(16))
	v298 = v294 ^ v286 - base.I32_rotl(v294, int32(4))
	v302 = v298 ^ v290 - base.I32_rotl(v298, v280)
	return v32 ^ v24 - base.I32_rotl(v32, int32(24)) ^ (v302 ^ v294 - base.I32_rotl(v302, int32(24)))
}
func F_gseg_penalty(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 float32
	_ = v2
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 float32
	_ = v22
	var v23 float32
	_ = v23
	var v29 float32
	_ = v29
	var v30 int32
	_ = v30
	var v33 float32
	_ = v33
	var v34 float32
	_ = v34
	var v40 float32
	_ = v40
	v2 = float32(0)
	v8 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	v15 = F_DirectFunctionCall2Coll(m, int32(_a_F_gseg_penalty_0), int32(0), v12, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int64(0)
	} else {
		v19 = base.I32_wrap_i64(v15)
		if v19 == int32(0) {
			v29 = v2
		} else {
			v22 = *(*float32)(unsafe.Add(mBase, uint32(v19)+4))
			v23 = *(*float32)(unsafe.Add(mBase, uint32(v19)))
			if base.F32_le(v22, v23) != 0 {
				v29 = v2
			} else {
				v29 = base.F32_abs(base.F32_sub(v22, v23))
			}
		}
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
		if v30 == int32(0) {
			v40 = v2
		} else {
			v33 = *(*float32)(unsafe.Add(mBase, uint32(v30)+4))
			v34 = *(*float32)(unsafe.Add(mBase, uint32(v30)))
			if base.F32_le(v33, v34) != 0 {
				v40 = v2
			} else {
				v40 = base.F32_abs(base.F32_sub(v33, v34))
			}
		}
		*(*float32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v8)))) = base.F32_sub(v29, v40)
		return v8
	}
}
func F_gtsvectorin(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_gtsvectorin_0), int32(94), int32(_a_F_gtsvectorin_1), int32(_a_F_gtsvectorin_2), int32(_a_F_gtsvectorin_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
