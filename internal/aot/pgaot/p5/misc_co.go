package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_ConditionVariableSignal(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
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
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	v8 = base.AtomicRmwXchg32(m, l0, int32(0), int32(1))
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_s_lock(m, l0, int32(_a_F_ConditionVariableSignal_0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v12 == int32(-1) {
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
	v15 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v15))
	return
L7:
	;
	goto L8
L8:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableSignal[0]))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v23 = v20 + v12*int32(768)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+356))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+360))
	if v25 == int32(-1) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v24 == int32(-1) {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v24
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v23)+360))
	v34 = v29
	goto L9
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20+v25*int32(768))+356)) = v24
	v34 = v25
	goto L9
L13:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+356)) = int64(0)
	v47 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v47))
	v51 = v23 + int32(316)
	v55 = base.AtomicRmwOr32(m, v47, int32(_a_F_ConditionVariableSignal_1), v47)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	if v56 != 0 {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v34
	goto L13
L15:
	;
	goto L16
L16:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableSignal[0]))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	*(*int32)(unsafe.Add(mBase, uint32(v40+v24*int32(768))+360)) = v34
	goto L13
L17:
	;
	return
L18:
	;
	goto L17
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = int32(1)
	v59 = int32(0)
	v62 = base.AtomicRmwOr32(m, v59, int32(_a_F_ConditionVariableSignal_1), v59)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if v63 == v59 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	if v66 == int32(0) {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableSignal[1]))
	if v70 == v66 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v72 = m.G0
	v74 = v72 - int32(16)
	m.G0 = v74
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableSignal[2]))
	if v77 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v100 = F_pgmem_kill(m, v66, int32(23))
	mBase = m.M
	goto L18
L25:
	;
	m.G0 = v74 + int32(16)
	goto L17
L26:
	;
	v80 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v74)+15)) = uint8(v80)
	goto L27
L27:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableSignal[3]))
	v88 = F_write(m, v84, v74+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v88 {
		goto L25
	} else {
		goto L29
	}
L28:
	;
	goto L25
L29:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableSignal[4]))
	if v92 == int32(27) {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
}
func F_ConditionalLockBuffer(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v43 int64
	_ = v43
	var v47 int64
	_ = v47
	var v49 int64
	_ = v49
	var v53 int32
	_ = v53
	var v54 int64
	_ = v54
	var v57 int64
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	if int32(0) <= l0 {
		v8 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBuffer[0]))
		v10 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBuffer[1]))
		if v10 != int32(-1) {
			v14 = v10 << (uint(int32(4)) % 32)
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)+uint32(_c_F_ConditionalLockBuffer[2])))
			if v17 == l0 {
				v25 = v14 + int32(_a_F_ConditionalLockBuffer_0)
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
				if v26 != 0 {
					return int32(0)
				} else {
					v29 = int32(_a_F_ConditionalLockBuffer_1)
					v31 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBuffer[3]))
					*(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBuffer[3])) = v31 + int32(1)
					v39 = v8 + l0*int32(56) - int32(32)
					v40 = int64(0)
					v43 = base.AtomicRmwCmpxchg64(m, v39, int32(0), v40, v40)
					v47 = v43
					for {
						v49 = int64(0)
						v53 = base.B2i32(v47&int64(18014381329612800) == v49)
						if v47&int64(18014381329612800) == v49 {
							v54 = int64(9007199254740992)
						} else {
							v54 = v49
						}
						v57 = base.AtomicRmwCmpxchg64(m, v39, int32(0), v47, v54|v47)
						if v47 != v57 {
							v47 = v57
							continue
						} else {
							break
						}
						break
					}
					if v53 == int32(0) {
						v61 = int32(_a_F_ConditionalLockBuffer_1)
						v63 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBuffer[3]))
						*(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBuffer[3])) = v63 - int32(1)
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = int32(3)
						return int32(1)
					}
				}
			} else {
				v21 = F_GetPrivateRefCountEntrySlow(m, l0, int32(1))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = v21
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
					if v26 != 0 {
						return int32(0)
					} else {
						v29 = int32(_a_F_ConditionalLockBuffer_1)
						v31 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBuffer[3]))
						*(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBuffer[3])) = v31 + int32(1)
						v39 = v8 + l0*int32(56) - int32(32)
						v40 = int64(0)
						v43 = base.AtomicRmwCmpxchg64(m, v39, int32(0), v40, v40)
						v47 = v43
						for {
							v49 = int64(0)
							v53 = base.B2i32(v47&int64(18014381329612800) == v49)
							if v47&int64(18014381329612800) == v49 {
								v54 = int64(9007199254740992)
							} else {
								v54 = v49
							}
							v57 = base.AtomicRmwCmpxchg64(m, v39, int32(0), v47, v54|v47)
							if v47 != v57 {
								v47 = v57
								continue
							} else {
								break
							}
							break
						}
						if v53 == int32(0) {
							v61 = int32(_a_F_ConditionalLockBuffer_1)
							v63 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBuffer[3]))
							*(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBuffer[3])) = v63 - int32(1)
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = int32(3)
							return int32(1)
						}
					}
				}
			}
		} else {
			v21 = F_GetPrivateRefCountEntrySlow(m, l0, int32(1))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = v21
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
				if v26 != 0 {
					return int32(0)
				} else {
					v29 = int32(_a_F_ConditionalLockBuffer_1)
					v31 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBuffer[3]))
					*(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBuffer[3])) = v31 + int32(1)
					v39 = v8 + l0*int32(56) - int32(32)
					v40 = int64(0)
					v43 = base.AtomicRmwCmpxchg64(m, v39, int32(0), v40, v40)
					v47 = v43
					for {
						v49 = int64(0)
						v53 = base.B2i32(v47&int64(18014381329612800) == v49)
						if v47&int64(18014381329612800) == v49 {
							v54 = int64(9007199254740992)
						} else {
							v54 = v49
						}
						v57 = base.AtomicRmwCmpxchg64(m, v39, int32(0), v47, v54|v47)
						if v47 != v57 {
							v47 = v57
							continue
						} else {
							break
						}
						break
					}
					if v53 == int32(0) {
						v61 = int32(_a_F_ConditionalLockBuffer_1)
						v63 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBuffer[3]))
						*(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBuffer[3])) = v63 - int32(1)
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = int32(3)
						return int32(1)
					}
				}
			}
		}
	} else {
		return int32(1)
	}
}
func F_ConditionalLockRelationOid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v23 int32
	_ = v23
	var v36 int32
	_ = v36
	var v53 int32
	_ = v53
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v11 = int32(1)
	if l0 <= int32(3591) {
		if l0 <= int32(2670) {
			switch l0 - int32(1213) {
			case 0, 1, 19, 20, 47, 48, 49:
				v82 = v11
			case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
				v82 = int32(0)
			default:
				if base.Ui32(int32(2)) <= base.Ui32(l0-int32(2396)) {
					v82 = int32(0)
				} else {
					v82 = v11
				}
			}
		} else {
			v23 = l0 - int32(2671)
			if base.B2i32(base.Ui32(int32(27)) < base.Ui32(v23))|base.B2i32(int32(1)<<(uint(v23)%32)&int32(226492515) == int32(0)) != 0 {
				if base.B2i32(base.Ui32(l0-int32(2964)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(l0-int32(2846)) < base.Ui32(int32(2))) != 0 {
					v82 = v11
				} else {
					v82 = int32(0)
				}
			} else {
				v82 = v11
			}
		}
	} else {
		if l0 <= int32(_a_F_ConditionalLockRelationOid_0) {
			v36 = l0 - int32(_a_F_ConditionalLockRelationOid_1)
			if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v36))|base.B2i32(int32(1)<<(uint(v36)%32)&int32(963) == int32(0)) != 0 {
				if base.Ui32(l0-int32(3592)) < base.Ui32(int32(2)) {
					v82 = v11
				} else {
					if base.Ui32(int32(2)) <= base.Ui32(l0-int32(4060)) {
						v82 = int32(0)
					} else {
						v82 = v11
					}
				}
			} else {
				v82 = v11
			}
		} else {
			switch l0 - int32(_a_F_ConditionalLockRelationOid_2) {
			case 0, 1, 2, 3, 4, 59, 60:
				v82 = v11
			case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
				v82 = int32(0)
			default:
				if base.Ui32(l0-int32(_a_F_ConditionalLockRelationOid_3)) < base.Ui32(int32(3)) {
					v82 = v11
				} else {
					v53 = l0 - int32(_a_F_ConditionalLockRelationOid_4)
					if base.Ui32(int32(15)) < base.Ui32(v53) {
						v82 = int32(0)
					} else {
						if int32(1)<<(uint(v53)%32)&int32(_a_F_ConditionalLockRelationOid_5) != 0 {
							v82 = v11
						} else {
							v82 = int32(0)
						}
					}
				}
			}
		}
	}
	*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = int64(72057594037927936)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l0
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockRelationOid[0]))
	if v82 != 0 {
		v89 = int32(0)
	} else {
		v89 = v88
	}
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v89
	v93 = int32(0)
	v98 = F_LockAcquireExtended(m, v7+int32(16), l1, v93, int32(1), v7+int32(12), v93)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		return int32(0)
	} else {
		switch v98 {
		case 0, 3:
			m.G0 = v7 + int32(32)
			return base.B2i32(v98 != int32(0))
		default:
			F_ReceiveSharedInvalidMessages(m)
			mBase = m.M
			v103 = m.ExcPending
			if v103 != 0 {
				return int32(0)
			} else {
				v104 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
				v105 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v104)+53)) = uint8(v105)
				m.G0 = v7 + int32(32)
				return base.B2i32(v98 != int32(0))
			}
		}
	}
}
func F_ConditionalLockTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	v5 = int32(0)
	v6 = m.G0
	v7 = int32(16)
	v8 = v6 - v7
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v12
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v14 | v15<<(uint(v7)%32)
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v21 = int32(260)
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+14)) = uint16(v21)
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+12)) = uint16(v20)
	v27 = F_LockAcquireExtended(m, v8, l2, v5, int32(1), v5, l3)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		return int32(0)
	} else {
		m.G0 = v8 + int32(16)
		return base.B2i32(v27 != int32(0))
	}
}
func F_ConversionIsVisibleExt(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v123 int64
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	v3 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v16 = F_SearchSysCache1(m, int32(20), base.I64_extend_i32_u(l0))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return v158
L2:
	;
	return int32(0)
L3:
	;
	if v16 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if l1 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+22)))
	F_recomputeNamespacePath(m)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L2
	} else {
		goto L13
	}
L7:
	;
	v22 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v22)
	v158 = int32(0)
	goto L1
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
	F_errmsg_internal(m, int32(_a_F_ConversionIsVisibleExt_0), v12)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(_a_F_ConversionIsVisibleExt_1), int32(2605), int32(_a_F_ConversionIsVisibleExt_2))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L13:
	;
	v42 = v38 + v39
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+68))
	if v43 != int32(11) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	F_ReleaseCatCache(m, v16)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L2
	} else {
		goto L44
	}
L15:
	;
	v46 = int32(0)
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_ConversionIsVisibleExt[0]))
	if v48 == v46 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	F_recomputeNamespacePath(m)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L2
	} else {
		goto L32
	}
L18:
	;
	if v87 == int32(0) {
		v147 = v46
		goto L14
	} else {
		goto L31
	}
L19:
	;
	v87 = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	if v55 <= int32(0) {
		v81 = v46
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v87 = v81
	goto L18
L23:
	;
	v58 = int32(0)
	if v58 < v55 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v61 = v55
	goto L26
L25:
	;
	v61 = v58
	goto L26
L26:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	v64 = int32(0)
	goto L27
L27:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v62+v64<<(uint(int32(2))%32))))
	v73 = base.B2i32(v72 == v43)
	if v72 == v43 {
		v81 = v73
		goto L22
	} else {
		goto L29
	}
L28:
	;
	v81 = v73
	goto L22
L29:
	;
	v75 = v64 + int32(1)
	if v75 != v61 {
		v64 = v75
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	goto L17
L32:
	;
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_ConversionIsVisibleExt[0]))
	if v94 == int32(0) {
		v139 = v3
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v147 = base.B2i32(l0 == v139)
	goto L14
L34:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	if v97 <= int32(0) {
		v139 = v3
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_ConversionIsVisibleExt[1]))
	v107 = int32(0)
	v109 = v104
	v113 = v97
	goto L36
L36:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v94)+12))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v115+v107<<(uint(int32(2))%32))))
	if v109 != v119 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v139 = int32(0)
	goto L33
L38:
	;
	v123 = int64(0)
	v125 = F_GetSysCacheOid(m, int32(18), base.I64_extend_i32_u(v42+int32(4)), base.I64_extend_i32_u(v119), v123, v123)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L2
	} else {
		goto L41
	}
L39:
	;
	v130 = v109
	v131 = v113
	goto L40
L40:
	;
	v133 = v107 + int32(1)
	if v133 < v131 {
		v107 = v133
		v109 = v130
		v113 = v131
		goto L36
	} else {
		goto L43
	}
L41:
	;
	if v125 != 0 {
		v139 = v125
		goto L33
	} else {
		goto L42
	}
L42:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	v129 = *(*int32)(unsafe.Add(mBase, _c_F_ConversionIsVisibleExt[1]))
	v130 = v129
	v131 = v127
	goto L40
L43:
	;
	goto L37
L44:
	;
	v158 = v147
	goto L1
}
func F_CopyGetAttnums(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	v4 = int32(0)
	v11 = m.G0
	v13 = v11 + int32(-64)
	m.G0 = v13
	if l2 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 - int32(-64)
	return v245
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v15 <= int32(0) {
		v245 = v4
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v213 <= int32(0) {
		v245 = v4
		goto L1
	} else {
		goto L67
	}
L5:
	;
	v23 = v4
	v26 = v4
	goto L8
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L34
	} else {
		goto L62
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L34
	} else {
		goto L58
	}
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v26<<(uint(int32(2))%32))))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v34 = int32(0)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v35 <= v34 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v158 + int32(4)
	F_errmsg(m, int32(_a_F_CopyGetAttnums_0), v11+int32(-48))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L34
	} else {
		goto L56
	}
L10:
	;
	goto L9
L11:
	;
	v113 = int32(0)
	if v23 == v113 {
		goto L41
	} else {
		goto L42
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L34
	} else {
		goto L35
	}
L13:
	;
	v41 = v35
	v45 = v34
	goto L14
L14:
	;
	v53 = l0 + v41<<(uint(int32(3))%32) + v45*int32(100)
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+119)))
	if v54 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+90)))
	if v83 != 0 {
		goto L6
	} else {
		goto L32
	}
L16:
	;
	goto L15
L17:
	;
	v58 = v53 + int32(28)
	v60 = v53 + int32(32)
	if v60|v33 != 0 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v79 = v41
	goto L19
L19:
	;
	v81 = v45 + int32(1)
	if v81 < v79 {
		v41 = v79
		v45 = v81
		goto L14
	} else {
		goto L31
	}
L20:
	;
	if v75 == int32(0) {
		goto L16
	} else {
		goto L30
	}
L21:
	;
	v66 = int32(-1)
	goto L23
L22:
	;
	v66 = int32(0)
	goto L23
L23:
	;
	if v60 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v67 = int32(1)
	goto L26
L25:
	;
	v67 = v66
	goto L26
L26:
	;
	v68 = int32(0)
	if base.B2i32(v60 == v68)|base.B2i32(v33 == v68) != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v75 = v67
	goto L29
L28:
	;
	v74 = F_strncmp(m, v60, v33, int32(64))
	mBase = m.M
	v75 = v74
	goto L29
L29:
	;
	goto L20
L30:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v79 = v78
	goto L19
L31:
	;
	goto L12
L32:
	;
	v84 = int32(*(*int16)(unsafe.Add(mBase, uint32(v58)+74)))
	if v84 != 0 {
		goto L11
	} else {
		goto L33
	}
L33:
	;
	goto L12
L34:
	;
	return int32(0)
L35:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	if l1 != 0 {
		goto L10
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v33
	F_errmsg(m, int32(_a_F_CopyGetAttnums_1), v13)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L34
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_CopyGetAttnums_2), int32(1129), int32(_a_F_CopyGetAttnums_3))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L34
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	if v151 != 0 {
		goto L7
	} else {
		goto L53
	}
L41:
	;
	v151 = int32(0)
	goto L40
L42:
	;
	goto L43
L43:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v119 <= int32(0) {
		v145 = v113
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v151 = v145
	goto L40
L45:
	;
	v122 = int32(0)
	if v122 < v119 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v125 = v119
	goto L48
L47:
	;
	v125 = v122
	goto L48
L48:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v128 = int32(0)
	goto L49
L49:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v126+v128<<(uint(int32(2))%32))))
	v137 = base.B2i32(v136 == v84)
	if v136 == v84 {
		v145 = v137
		goto L44
	} else {
		goto L51
	}
L50:
	;
	v145 = v137
	goto L44
L51:
	;
	v139 = v128 + int32(1)
	if v139 != v125 {
		v128 = v139
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v152 = F_lappend_int(m, v23, v84)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L34
	} else {
		goto L54
	}
L54:
	;
	v155 = v26 + int32(1)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v155 < v156 {
		v23 = v152
		v26 = v155
		goto L8
	} else {
		goto L55
	}
L55:
	;
	v245 = v152
	goto L1
L56:
	;
	F_errfinish(m, int32(_a_F_CopyGetAttnums_2), int32(1124), int32(_a_F_CopyGetAttnums_3))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L34
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	F_errcode(m, int32(16806020))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L34
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v33
	F_errmsg(m, int32(_a_F_CopyGetAttnums_4), v11+int32(-32))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L34
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_CopyGetAttnums_2), int32(1136), int32(_a_F_CopyGetAttnums_3))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L34
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
	F_errcode(m, int32(_a_F_CopyGetAttnums_5))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L34
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v33
	F_errmsg(m, int32(_a_F_CopyGetAttnums_6), v11+int32(-16))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L34
	} else {
		goto L64
	}
L64:
	;
	v206 = F_errdetail(m, int32(_a_F_CopyGetAttnums_7), int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L34
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_CopyGetAttnums_2), int32(1113), int32(_a_F_CopyGetAttnums_3))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L34
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
	v219 = v4
	v221 = v4
	goto L68
L68:
	;
	v227 = v219 + int32(1)
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v219<<(uint(int32(3))%32))+34)))
	if v231&int32(12) == int32(0) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v245 = v238
	goto L1
L70:
	;
	v236 = F_lappend_int(m, v221, v227)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L34
	} else {
		goto L73
	}
L71:
	;
	v238 = v221
	goto L72
L72:
	;
	if v227 != v213 {
		v219 = v227
		v221 = v238
		goto L68
	} else {
		goto L74
	}
L73:
	;
	v238 = v236
	goto L72
L74:
	;
	goto L69
}
func F_CopyLimitPrintoutLength(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	v4 = F_strlen(m, l0)
	mBase = m.M
	if v4 <= int32(100) {
		v7 = F_pstrdup(m, l0)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			return v7
		}
	} else {
		v13 = F_pg_mbcliplen(m, l0, v4, int32(100))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v17 = F_palloc(m, v13+int32(4))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				if v13 != 0 {
					base.MemoryCopy(m, v17, l0, v13)
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(v13+v17))) = int32(_a_F_CopyLimitPrintoutLength_0)
				return v17
			}
		}
	}
}
func F_CopyReadAttributesText(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
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
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
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
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v458 int32
	_ = v458
	v2 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+296))
	if v19 <= v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v17 + int32(16)
	return v458
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	if v22 == int32(0) {
		v458 = v2
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	v46 = l0 + int32(280)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v48 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v47))) = uint8(v48)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+12)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v48
	goto L11
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	F_errmsg(m, int32(_a_F_CopyReadAttributesText_0), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	F_errfinish(m, int32(_a_F_CopyReadAttributesText_1), int32(1871), int32(_a_F_CopyReadAttributesText_2))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L11:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
	if v55 <= v54 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	F_enlargeStringInfo(m, v46, v54)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L6
	} else {
		goto L15
	}
L13:
	;
	v60 = v54
	goto L14
L14:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+304))
	v62 = v60 + v61
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	v65 = v63
	v67 = v61
	v71 = v2
	goto L16
L15:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	v60 = v59
	goto L14
L16:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+300))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+296))
	if v79 <= v71 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+284)) = v265 - v448
	v458 = v445
	goto L1
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+296)) = v79 << (uint(int32(1)) % 32)
	v86 = F_repalloc(m, v78, v79<<(uint(int32(3))%32))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L6
	} else {
		goto L21
	}
L19:
	;
	v89 = v78
	goto L20
L20:
	;
	v91 = v71 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v89+v91))) = v65
	v94 = int32(0)
	if base.Ui32(v62) <= base.Ui32(v67) {
		v262 = v67
		v264 = v67
		v265 = v65
		v269 = v94
		v273 = v94
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+300)) = v86
	v89 = v86
	goto L20
L22:
	;
	v275 = v262 - v67
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v275 != v276 {
		goto L76
	} else {
		goto L77
	}
L23:
	;
	v98 = v67
	v101 = v65
	v105 = v94
	goto L24
L24:
	;
	v112 = v98 + int32(1)
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	v114 = base.B2i32(v113 == v44)
	if v113 == v44 {
		v262 = v98
		v264 = v112
		v265 = v101
		v269 = v105
		v273 = v114
		goto L22
	} else {
		goto L26
	}
L25:
	;
	v262 = v253
	v264 = v253
	v265 = v259
	v269 = v256
	v273 = v114
	goto L22
L26:
	;
	if v113 != int32(92) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v101))) = uint8(v254)
	v259 = v101 + int32(1)
	if base.Ui32(v253) < base.Ui32(v62) {
		v98 = v253
		v101 = v259
		v105 = v256
		goto L24
	} else {
		goto L74
	}
L28:
	;
	v253 = v112
	v254 = v113
	v256 = v105
	goto L27
L29:
	;
	goto L30
L30:
	;
	if base.Ui32(v62) <= base.Ui32(v112) {
		v262 = v98
		v264 = v112
		v265 = v101
		v269 = v105
		v273 = v114
		goto L22
	} else {
		goto L31
	}
L31:
	;
	v119 = v98 + int32(2)
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+1)))
	v122 = v120 - int32(48)
	switch v122 {
	case 0, 1, 2, 3, 4, 5, 6, 7:
		goto L39
	default:
		v253 = v119
		v254 = v120
		v256 = v105
		goto L27
	case 50:
		goto L37
	case 54:
		goto L36
	case 62:
		goto L35
	case 66:
		goto L34
	case 68:
		goto L33
	case 70:
		goto L32
	case 72:
		goto L38
	}
L32:
	;
	v253 = v119
	v254 = int32(11)
	v256 = v105
	goto L27
L33:
	;
	v253 = v119
	v254 = int32(9)
	v256 = v105
	goto L27
L34:
	;
	v253 = v119
	v254 = int32(13)
	v256 = v105
	goto L27
L35:
	;
	v253 = v119
	v254 = int32(10)
	v256 = v105
	goto L27
L36:
	;
	v253 = v119
	v254 = int32(12)
	v256 = v105
	goto L27
L37:
	;
	v253 = v119
	v254 = int32(8)
	v256 = v105
	goto L27
L38:
	;
	v156 = int32(120)
	if base.Ui32(v62) <= base.Ui32(v119) {
		v253 = v119
		v254 = v156
		v256 = v105
		goto L27
	} else {
		goto L49
	}
L39:
	;
	if base.Ui32(v62) <= base.Ui32(v119) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v253 = v149
	v254 = v150
	v256 = base.B2i32(base.I32_extend8_s(v150) <= int32(0)) | v105
	goto L27
L41:
	;
	v149 = v119
	v150 = v122
	goto L40
L42:
	;
	goto L43
L43:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	if v124&int32(248) != int32(48) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v149 = v119
	v150 = v122
	goto L40
L45:
	;
	goto L46
L46:
	;
	v129 = int32(3)
	v133 = v124 + v122<<(uint(v129)%32) - int32(48)
	v135 = v98 + v129
	if base.Ui32(v62) <= base.Ui32(v135) {
		v149 = v135
		v150 = v133
		goto L40
	} else {
		goto L47
	}
L47:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
	if v137&int32(248) != int32(48) {
		v149 = v135
		v150 = v133
		goto L40
	} else {
		goto L48
	}
L48:
	;
	v149 = v98 + int32(4)
	v150 = v137 + v133<<(uint(int32(3))%32) - int32(48)
	goto L40
L49:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	goto L50
L50:
	;
	if base.B2i32(base.Ui32(v158-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v158|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		v253 = v119
		v254 = v156
		v256 = v105
		goto L27
	} else {
		goto L51
	}
L51:
	;
	v179 = base.B2i32(base.Ui32((v158-int32(48))&int32(255)) < base.Ui32(int32(10)))
	if base.Ui32((v158-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v180 = int32(-48)
	goto L54
L53:
	;
	v180 = int32(-87)
	goto L54
L54:
	;
	if base.Ui32((v158-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v189 = v158 | int32(32)
	goto L57
L56:
	;
	v189 = v158
	goto L57
L57:
	;
	if base.Ui32((v158-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v190 = v158
	goto L60
L59:
	;
	v190 = v189
	goto L60
L60:
	;
	v191 = v180 + v190
	v193 = v98 + int32(3)
	if base.Ui32(v62) <= base.Ui32(v193) {
		v234 = v193
		v235 = v191
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v253 = v234
	v254 = v235
	v256 = base.B2i32(v235&int32(255) == int32(0)) | int32(base.Ui32(v235&int32(128))>>(uint(int32(7))%32)) | v105
	goto L27
L62:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193))))
	goto L63
L63:
	;
	if base.B2i32(base.Ui32(v195-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v195|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		v234 = v193
		v235 = v191
		goto L61
	} else {
		goto L64
	}
L64:
	;
	v218 = base.B2i32(base.Ui32((v195-int32(48))&int32(255)) < base.Ui32(int32(10)))
	if base.Ui32((v195-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v219 = int32(-48)
	goto L67
L66:
	;
	v219 = int32(-87)
	goto L67
L67:
	;
	if base.Ui32((v195-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v229 = v195 | int32(32)
	goto L70
L69:
	;
	v229 = v195
	goto L70
L70:
	;
	if base.Ui32((v195-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v230 = v195
	goto L73
L72:
	;
	v230 = v229
	goto L73
L73:
	;
	v234 = v98 + int32(4)
	v235 = v191<<(uint(int32(4))%32) + v219 + v230
	goto L61
L74:
	;
	goto L25
L75:
	;
	v442 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v265))) = uint8(v442)
	v444 = int32(1)
	v445 = v71 + v444
	if v273 != 0 {
		v65 = v265 + v444
		v67 = v264
		v71 = v445
		goto L16
	} else {
		goto L123
	}
L76:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v328 != 0 {
		goto L93
	} else {
		goto L94
	}
L77:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v275 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	if v323 != 0 {
		goto L76
	} else {
		goto L91
	}
L79:
	;
	v323 = int32(0)
	goto L78
L80:
	;
	goto L81
L81:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	if v284 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v285 = v67
	v286 = v278
	v287 = v275
	v288 = v284
	goto L86
L83:
	;
	v311 = v278
	v315 = int32(0)
	goto L84
L84:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311))))
	v323 = v315 - v316
	goto L78
L85:
	;
	v311 = v306
	v315 = v308
	goto L84
L86:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286))))
	if base.B2i32(v288 != v290)|base.B2i32(v290 == int32(0)) != 0 {
		v306 = v286
		v308 = v288
		goto L85
	} else {
		goto L88
	}
L87:
	;
	v306 = v300
	v308 = int32(0)
	goto L85
L88:
	;
	v296 = v287 - int32(1)
	if v296 == int32(0) {
		v306 = v286
		v308 = v288
		goto L85
	} else {
		goto L89
	}
L89:
	;
	v299 = int32(1)
	v300 = v286 + v299
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285)+1)))
	if v301 != 0 {
		v285 = v285 + v299
		v286 = v300
		v287 = v296
		v288 = v301
		goto L86
	} else {
		goto L90
	}
L90:
	;
	goto L87
L91:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l0)+300))
	*(*int32)(unsafe.Add(mBase, uint32(v324+v91))) = int32(0)
	goto L75
L92:
	;
	if v269&int32(1) == int32(0) {
		goto L75
	} else {
		goto L121
	}
L93:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v328)+4))
	v331 = v329
	goto L95
L94:
	;
	v331 = int32(0)
	goto L95
L95:
	;
	if v331 <= v71 {
		goto L92
	} else {
		goto L96
	}
L96:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v333 == int32(0) {
		goto L92
	} else {
		goto L97
	}
L97:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v275 != v336 {
		goto L92
	} else {
		goto L98
	}
L98:
	;
	if v275 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	if v382 != 0 {
		goto L92
	} else {
		goto L112
	}
L100:
	;
	v382 = int32(0)
	goto L99
L101:
	;
	goto L102
L102:
	;
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	if v343 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v344 = v67
	v345 = v333
	v346 = v275
	v347 = v343
	goto L107
L104:
	;
	v370 = v333
	v374 = int32(0)
	goto L105
L105:
	;
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v370))))
	v382 = v374 - v375
	goto L99
L106:
	;
	v370 = v365
	v374 = v367
	goto L105
L107:
	;
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345))))
	if base.B2i32(v347 != v349)|base.B2i32(v349 == int32(0)) != 0 {
		v365 = v345
		v367 = v347
		goto L106
	} else {
		goto L109
	}
L108:
	;
	v365 = v359
	v367 = int32(0)
	goto L106
L109:
	;
	v355 = v346 - int32(1)
	if v355 == int32(0) {
		v365 = v345
		v367 = v347
		goto L106
	} else {
		goto L110
	}
L110:
	;
	v358 = int32(1)
	v359 = v345 + v358
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344)+1)))
	if v360 != 0 {
		v344 = v344 + v358
		v345 = v359
		v346 = v355
		v347 = v360
		goto L107
	} else {
		goto L111
	}
L111:
	;
	goto L108
L112:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v328)+12))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v384+v91)))
	v388 = v386 - int32(1)
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v383+v388<<(uint(int32(2))%32))))
	if v392 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v395 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v393+v388))) = uint8(v395)
	goto L75
L114:
	;
	goto L115
L115:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v397)+52))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v398)))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L6
	} else {
		goto L116
	}
L116:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L6
	} else {
		goto L117
	}
L117:
	;
	F_errmsg(m, int32(_a_F_CopyReadAttributesText_3), int32(0))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L6
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v398 + v399<<(uint(int32(3))%32) + v388*int32(100) + int32(32)
	v421 = F_errdetail(m, int32(_a_F_CopyReadAttributesText_4), v17)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L6
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(_a_F_CopyReadAttributesText_1), int32(2065), int32(_a_F_CopyReadAttributesText_2))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L6
	} else {
		goto L120
	}
L120:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L121:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(l0)+300))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v433+v91)))
	F_pg_verifymbstr(m, v435, v265-v435)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L6
	} else {
		goto L122
	}
L122:
	;
	goto L75
L123:
	;
	goto L17
}
func F_colNameToVar(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	v5 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	if l0 == v5 {
		v120 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v15 + int32(16)
	return v120
L2:
	;
	v23 = l0
	v29 = v5
	goto L3
L3:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	if v31 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v120 = int32(0)
	goto L1
L5:
	;
	if l2|v103 != 0 {
		v120 = v103
		goto L1
	} else {
		goto L30
	}
L6:
	;
	v103 = int32(0)
	goto L5
L7:
	;
	goto L8
L8:
	;
	v35 = int32(0)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v37 <= v35 {
		v103 = v35
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v45 = v35
	v48 = v35
	goto L10
L10:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v52+v48<<(uint(int32(2))%32))))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+21)))
	if v57 == int32(0) {
		v74 = v45
		goto L13
	} else {
		goto L14
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L19
	} else {
		goto L25
	}
L12:
	;
	goto L11
L13:
	;
	v77 = v48 + int32(1)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v77 < v78 {
		v45 = v74
		v48 = v77
		goto L10
	} else {
		goto L24
	}
L14:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+22)))
	if v60 == int32(1) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+32)))
	if v63 != int32(1) {
		v74 = v45
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v66 = F_scanNSItemForColumn(m, l0, v56, v29, l1, l3)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L17
L19:
	;
	return int32(0)
L20:
	;
	if v66 == int32(0) {
		v74 = v45
		goto L13
	} else {
		goto L21
	}
L21:
	;
	if v45 != 0 {
		goto L12
	} else {
		goto L22
	}
L22:
	;
	F_check_lateral_ref_ok(m, v23, v56, l3)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v74 = v66
	goto L13
L24:
	;
	v103 = v74
	goto L5
L25:
	;
	F_errcode(m, int32(33583236))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L19
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l1
	F_errmsg(m, int32(_a_F_colNameToVar_0), v15)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L19
	} else {
		goto L27
	}
L27:
	;
	F_parser_errposition(m, v23, l3)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L19
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_colNameToVar_1), int32(960), int32(_a_F_colNameToVar_2))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L19
	} else {
		goto L29
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L30:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v113 != 0 {
		v23 = v113
		v29 = v29 + int32(1)
		goto L3
	} else {
		goto L31
	}
L31:
	;
	goto L4
}
func F_combo_init(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v11 = m.T0[v10].(func(*base.Module, int32) int32)(m, v9)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
		v16 = m.T0[v15].(func(*base.Module, int32) int32)(m, v9)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			if v16 == int32(0) {
				v30 = int32(0)
				v31 = F_palloc0(m, v11)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					if base.Ui32(l2) < base.Ui32(v11) {
						v34 = l2
					} else {
						v34 = v11
					}
					if v34 != 0 {
						base.MemoryCopy(m, v31, l1, v34)
					} else {
					}
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
					v37 = m.T0[v36].(func(*base.Module, int32, int32, int32, int32) int32)(m, v9, v31, v34, v30)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						if v30 != 0 {
							F_pfree(m, v30)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								F_pfree(m, v31)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									return v37
								}
							}
						} else {
							F_pfree(m, v31)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								return v37
							}
						}
					}
				}
			} else {
				v20 = F_palloc0(m, v16)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					if base.Ui32(l4) <= base.Ui32(v16) {
						if l4 == int32(0) {
							v30 = v20
						} else {
							v25 = l4
							if v25 == int32(0) {
								v30 = v20
							} else {
								base.MemoryCopy(m, v20, l3, v25)
								v30 = v20
							}
						}
					} else {
						v25 = v16
						if v25 == int32(0) {
							v30 = v20
						} else {
							base.MemoryCopy(m, v20, l3, v25)
							v30 = v20
						}
					}
					v31 = F_palloc0(m, v11)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						if base.Ui32(l2) < base.Ui32(v11) {
							v34 = l2
						} else {
							v34 = v11
						}
						if v34 != 0 {
							base.MemoryCopy(m, v31, l1, v34)
						} else {
						}
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
						v37 = m.T0[v36].(func(*base.Module, int32, int32, int32, int32) int32)(m, v9, v31, v34, v30)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							if v30 != 0 {
								F_pfree(m, v30)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int32(0)
								} else {
									F_pfree(m, v31)
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
										return int32(0)
									} else {
										return v37
									}
								}
							} else {
								F_pfree(m, v31)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									return v37
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_compact(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	v3 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = v3
	v13 = v9
	v14 = v3
	goto L4
L2:
	;
	v30 = v3
	v36 = int32(0)
	goto L3
L3:
	;
	v38 = F_palloc_extended(m, v30, int32(2))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v18 = int32(1)
	v19 = v12 + v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v23 = v14 + v20 + v18
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	if v24 != 0 {
		v12 = v19
		v13 = v24
		v14 = v23
		goto L4
	} else {
		goto L6
	}
L5:
	;
	v30 = v19
	v36 = v23 << (uint(int32(3)) % 32)
	goto L3
L6:
	;
	goto L5
L7:
	;
	return
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v38
	v41 = int32(2)
	v44 = F_palloc_extended(m, v30<<(uint(v41)%32), v41)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v44
	v48 = F_palloc_extended(m, v36, int32(2))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v48
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v51 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v30
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v73
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v76
	v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+56)))
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+20)) = uint16(v78)
	v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+58)))
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+22)) = uint16(v80)
	v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+60)))
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+24)) = uint16(v82)
	v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+62)))
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)) = uint16(v84)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	if v88 != 0 {
		goto L31
	} else {
		goto L32
	}
L12:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v48 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v57 != 0 {
		goto L20
	} else {
		goto L21
	}
L15:
	;
	v54 = v52
	goto L17
L16:
	;
	v54 = int32(0)
	goto L17
L17:
	;
	if v54 != 0 {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	F_pfree(m, v51)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	goto L14
L20:
	;
	F_pfree(m, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L7
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v60 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L22
L24:
	;
	F_pfree(m, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L7
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v63)+24)) = int32(101)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
	if v67 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L26
L28:
	;
	v69 = v67
	goto L30
L29:
	;
	v69 = int32(12)
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66)+12)) = v69
	return
L31:
	;
	v93 = int32(0)
	goto L33
L32:
	;
	v90 = int32(*(*int16)(unsafe.Add(mBase, uint32(v86)+12)))
	v93 = v90 + int32(1)
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v93
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v95
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v97
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v99
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v101 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v106 = v48
	v108 = v101
	goto L37
L35:
	;
	goto L36
L36:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+20))
	if v208 != 0 {
		goto L64
	} else {
		goto L65
	}
L37:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	v113 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v110+v111))) = uint8(v113)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	*(*int32)(unsafe.Add(mBase, uint32(v115+v116<<(uint(int32(2))%32)))) = v106
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v108)+20))
	if v121 != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L36
L39:
	;
	v124 = v106
	v125 = v121
	goto L42
L40:
	;
	v177 = v106
	goto L41
L41:
	;
	v185 = (v177 - v106) >> (uint(int32(3)) % 32)
	if base.Ui32(int32(2)) <= base.Ui32(v185) {
		goto L59
	} else {
		goto L60
	}
L42:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	if v130 != int32(76) {
		goto L47
	} else {
		goto L48
	}
L43:
	;
	v177 = v173
	goto L41
L44:
	;
	v173 = v124 + int32(8)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v125)+16))
	if v174 != 0 {
		v124 = v173
		v125 = v174
		goto L42
	} else {
		goto L58
	}
L45:
	;
	v161 = v142 + v140
	*(*uint16)(unsafe.Add(mBase, uint32(v124))) = uint16(v161)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	*(*int32)(unsafe.Add(mBase, uint32(v124)+4)) = v164
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v166 | int32(1)
	goto L44
L46:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v153)+24)) = int32(101)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+12))
	if v157 != 0 {
		goto L55
	} else {
		goto L56
	}
L47:
	;
	if v130 != int32(112) {
		goto L46
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v140 = int32(*(*int16)(unsafe.Add(mBase, uint32(v125)+4)))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v140 <= int32(_a_F_compact_0)-v142 {
		goto L45
	} else {
		goto L51
	}
L50:
	;
	v135 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v125)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v124))) = uint16(v135)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	*(*int32)(unsafe.Add(mBase, uint32(v124)+4)) = v138
	goto L44
L51:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v145)+24)) = int32(101)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+12))
	if v149 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v151 = v149
	goto L54
L53:
	;
	v151 = int32(20)
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v148)+12)) = v151
	return
L55:
	;
	v159 = v157
	goto L57
L56:
	;
	v159 = int32(15)
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v156)+12)) = v159
	return
L58:
	;
	goto L43
L59:
	;
	F_pg_qsort(m, v106, v185, int32(8), int32(1036))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L7
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v177)+4)) = int32(0)
	v194 = int32(_a_F_compact_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v177))) = uint16(v194)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v108)+28))
	if v198 != 0 {
		v106 = v177 + int32(8)
		v108 = v198
		goto L37
	} else {
		goto L63
	}
L62:
	;
	goto L61
L63:
	;
	goto L38
L64:
	;
	v212 = v208
	goto L67
L65:
	;
	v227 = v207
	goto L66
L66:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v227)))
	v236 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v233+v234))) = uint8(v236)
	return
L67:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v212)+12))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
	v221 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v217+v219))) = uint8(v221)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v212)+16))
	if v223 != 0 {
		v212 = v223
		goto L67
	} else {
		goto L69
	}
L68:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v227 = v224
	goto L66
L69:
	;
	goto L68
}
func F_compareDocR(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
	v8 = int32(_a_F_compareDocR_0)
	v9 = v7 & v8
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
	v12 = v10 & v8
	if v9 == v12 {
		v14 = int32(14)
		v15 = int32(base.Ui32(v7) >> (uint(v14) % 32))
		v17 = int32(base.Ui32(v10) >> (uint(v14) % 32))
		if v15 == v17 {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v19 == v20 {
				return int32(0)
			} else {
				if base.Ui32(v20) < base.Ui32(v19) {
					v27 = int32(1)
				} else {
					v27 = int32(-1)
				}
				return v27
			}
		} else {
			if base.Ui32(v17) < base.Ui32(v15) {
				v32 = int32(1)
			} else {
				v32 = int32(-1)
			}
			return v32
		}
	} else {
		if base.Ui32(v12) < base.Ui32(v9) {
			v37 = int32(1)
		} else {
			v37 = int32(-1)
		}
		return v37
	}
}
func F_compare_scalars(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int64
	_ = v10
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v14 = m.T0[v13].(func(*base.Module, int64, int64, int32) int32)(m, v10, v11, v12)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v14 < int32(0) {
			v21 = int32(1)
		} else {
			v21 = int32(0) - v14
		}
		v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+8)))
		if v22 != 0 {
			v23 = v21
		} else {
			v23 = v14
		}
		if v23 != 0 {
			v42 = v23
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
			v27 = v24 + v7<<(uint(int32(2))%32)
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
			if v28 < v6 {
				*(*int32)(unsafe.Add(mBase, uint32(v27))) = v6
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
				v32 = v31
			} else {
				v32 = v24
			}
			v35 = v32 + v6<<(uint(int32(2))%32)
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
			if v36 < v7 {
				*(*int32)(unsafe.Add(mBase, uint32(v35))) = v7
			} else {
			}
			v42 = v7 - v6
		}
		return v42
	}
}
func F_computeRegionDelta(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	v13 = int32(-1)
	v15 = base.B2i32(base.Ui32(l3) < base.Ui32(l5))
	if base.Ui32(l3) < base.Ui32(l5) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v16 = l3
	goto L3
L2:
	;
	v16 = v13
	goto L3
L3:
	;
	if base.Ui32(l3) < base.Ui32(l5) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v138 < int32(0) {
		goto L45
	} else {
		goto L46
	}
L5:
	;
	v17 = l5
	goto L7
L6:
	;
	v17 = l3
	goto L7
L7:
	;
	if base.Ui32(l4) < base.Ui32(l6) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v19 = l4
	goto L10
L9:
	;
	v19 = l6
	goto L10
L10:
	;
	if base.Ui32(v19) <= base.Ui32(v17) {
		v138 = v16
		v139 = v13
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v26 = v17
	v30 = v16
	goto L12
L12:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v26))))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v26))))
	if v36 != v38 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v138 = v127
	v139 = v128
	goto L4
L14:
	;
	if v30 < int32(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v67 = v26
	v71 = v30
	goto L16
L16:
	;
	v77 = v67 + int32(1)
	if v77 < v19 {
		goto L26
	} else {
		goto L27
	}
L17:
	;
	v42 = v26
	goto L19
L18:
	;
	v42 = v30
	goto L19
L19:
	;
	v46 = v26
	goto L20
L20:
	;
	v56 = v46 + int32(1)
	if v19 <= v56 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v67 = v56
	v71 = v42
	goto L16
L22:
	;
	v138 = v42
	v139 = int32(-1)
	goto L4
L23:
	;
	goto L24
L24:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v56))))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v56))))
	if v60 != v62 {
		v46 = v56
		goto L20
	} else {
		goto L25
	}
L25:
	;
	goto L21
L26:
	;
	v79 = v19
	goto L28
L27:
	;
	v79 = v77
	goto L28
L28:
	;
	v85 = v67
	goto L29
L29:
	;
	if v85 == v79-int32(1) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	if v71 < int32(0) {
		goto L37
	} else {
		goto L38
	}
L31:
	;
	goto L30
L32:
	;
	v102 = v79
	goto L31
L33:
	;
	goto L34
L34:
	;
	v96 = v85 + int32(1)
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v96))))
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v96))))
	if v98 == v100 {
		v85 = v96
		goto L29
	} else {
		goto L35
	}
L35:
	;
	v102 = v96
	goto L31
L36:
	;
	if v102 < v19 {
		v26 = v102
		v30 = v127
		goto L12
	} else {
		goto L44
	}
L37:
	;
	v127 = int32(-1)
	v128 = v67
	goto L36
L38:
	;
	goto L39
L39:
	;
	if base.Ui32(v102-v67) < base.Ui32(int32(5)) {
		v127 = v71
		v128 = v67
		goto L36
	} else {
		goto L40
	}
L40:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v110 = l0 + int32(16) + v109
	v111 = v67 - v71
	*(*uint16)(unsafe.Add(mBase, uint32(v110)+2)) = uint16(v111)
	*(*uint16)(unsafe.Add(mBase, uint32(v110))) = uint16(v71)
	v115 = v111 & int32(_a_F_computeRegionDelta_0)
	if v115 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	base.MemoryCopy(m, v110+int32(4), l2+v71, v115)
	goto L43
L42:
	;
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v115 + v109 + int32(4)
	v124 = int32(-1)
	v127 = v124
	v128 = v124
	goto L36
L44:
	;
	goto L13
L45:
	;
	v145 = v19
	goto L47
L46:
	;
	v145 = v138
	goto L47
L47:
	;
	v146 = base.B2i32(base.Ui32(l6) < base.Ui32(l4))
	if base.Ui32(l6) < base.Ui32(l4) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v147 = v145
	goto L50
L49:
	;
	v147 = v138
	goto L50
L50:
	;
	if int32(0) <= v147 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v151 = l0 + v150
	*(*uint16)(unsafe.Add(mBase, uint32(v151)+16)) = uint16(v147)
	if base.Ui32(l6) < base.Ui32(l4) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L53
L53:
	;
	return
L54:
	;
	v153 = l4
	goto L56
L55:
	;
	v153 = v139
	goto L56
L56:
	;
	if v153 < int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v156 = l4
	goto L59
L58:
	;
	v156 = v153
	goto L59
L59:
	;
	v157 = v156 - v147
	*(*uint16)(unsafe.Add(mBase, uint32(v151)+18)) = uint16(v157)
	v160 = v157 & int32(_a_F_computeRegionDelta_0)
	if v160 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	base.MemoryCopy(m, v151+int32(20), v147+l2, v160)
	goto L62
L61:
	;
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v160 + v150 + int32(4)
	goto L53
}
func F_compute_remaining_iovec(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	v7 = l1
	v8 = l2
	v9 = l3
	goto L2
L1:
	;
	if l0 == v7 {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if base.Ui32(v9) < base.Ui32(v11) {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	v17 = v8 - int32(1)
	if v17 != 0 {
		v7 = v7 + int32(8)
		v8 = v17
		v9 = v9 - v11
		goto L2
	} else {
		goto L5
	}
L5:
	;
	goto L3
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v27 + v9
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v30 - v9
	return v8
L7:
	;
	v22 = v8 << (uint(int32(3)) % 32)
	if v22 == int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	base.MemoryCopy(m, l0, v7, v22)
	goto L6
}
func F_compute_trivial_stats(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v39 float64
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v54 int64
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v100 float64
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	v5 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+78)))
	if v19 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+76)))
	v25 = int32(_a_F_compute_trivial_stats_0)
	v30 = base.B2i32(v22 < int32(0))
	v31 = base.B2i32(v22&v25 == v25)
	goto L3
L2:
	;
	v30 = v5
	v31 = v5
	goto L3
L3:
	;
	if l2 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	m.G0 = v16 + int32(16)
	return
L5:
	;
	v39 = float64(0)
	v40 = int32(0)
	v44 = v5
	v46 = v5
	goto L6
L6:
	;
	F_vacuum_delay_point(m, int32(1))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if int32(0) < v101 {
		goto L32
	} else {
		goto L33
	}
L8:
	;
	return
L9:
	;
	v54 = m.T0[l1].(func(*base.Module, int32, int32, int32) int64)(m, l0, v46, v16+int32(15))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+15)))
	if v56 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v106 = v46 + int32(1)
	if v106 != l2 {
		v39 = v100
		v40 = v101
		v44 = v103
		v46 = v106
		goto L6
	} else {
		goto L30
	}
L12:
	;
	v100 = v39
	v101 = v40
	v103 = v44 + int32(1)
	goto L11
L13:
	;
	goto L14
L14:
	;
	v62 = v40 + int32(1)
	if v31 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v63 = base.I32_wrap_i64(v54)
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	if v64 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	if v30 == int32(0) {
		v100 = v39
		v101 = v62
		v103 = v44
		goto L11
	} else {
		goto L29
	}
L18:
	;
	v100 = base.F64_add(v39, base.F64_convert_i32_u(v89))
	v101 = v62
	v103 = v44
	goto L11
L19:
	;
	v68 = int32(18)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
	if v70 == v68 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v81 = int32(1)
	if v64&v81 != 0 {
		v89 = int32(base.Ui32(v64) >> (uint(v81) % 32))
		goto L18
	} else {
		goto L28
	}
L22:
	;
	v73 = v68
	goto L24
L23:
	;
	v73 = int32(2)
	goto L24
L24:
	;
	if base.Ui32((v70-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v80 = int32(6)
	goto L27
L26:
	;
	v80 = v73
	goto L27
L27:
	;
	v89 = v80
	goto L18
L28:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v89 = int32(base.Ui32(v85) >> (uint(int32(2)) % 32))
	goto L18
L29:
	;
	v95 = F_strlen(m, base.I32_wrap_i64(v54))
	mBase = m.M
	v100 = base.F64_add(v39, base.F64_convert_i32_u(v95+int32(1)))
	v101 = v62
	v103 = v44
	goto L11
L30:
	;
	goto L7
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v131
	goto L4
L32:
	;
	v110 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v110)
	*(*float32)(unsafe.Add(mBase, uint32(l0)+40)) = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v103), base.F64_convert_i32_s(l2)))
	if v30 != 0 {
		v131 = base.I32_trunc_sat_f64_s(base.F64_div(v100, base.F64_convert_i32_u(v101)))
		goto L31
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	if v103 <= int32(0) {
		goto L4
	} else {
		goto L36
	}
L35:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v121 = int32(*(*int16)(unsafe.Add(mBase, uint32(v120)+76)))
	v131 = v121
	goto L31
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(1065353216)
	v126 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v126)
	if v30 != 0 {
		v131 = int32(0)
		goto L31
	} else {
		goto L37
	}
L37:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v130 = int32(*(*int16)(unsafe.Add(mBase, uint32(v129)+76)))
	v131 = v130
	goto L31
}
func F_construct_empty_array(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_palloc0(m, int32(16))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = int32(0)
		*(*int64)(unsafe.Add(mBase, uint32(v4))) = int64(64)
		return v4
	}
}
func F_construct_md_array(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
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
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
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
	var v265 int32
	_ = v265
	v10 = int32(0)
	v19 = m.G0
	v21 = v19 + int32(-64)
	m.G0 = v21
	switch l8 - int32(99) {
	case 0:
		v44 = int32(1)
		goto L1
	case 1:
		goto L4
	default:
		goto L3
	case 6:
		goto L5
	case 16:
		goto L2
	}
L1:
	;
	if int32(0) <= l2 {
		goto L13
	} else {
		goto L14
	}
L2:
	;
	v44 = int32(2)
	goto L1
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v44 = int32(8)
	goto L1
L5:
	;
	v44 = int32(4)
	goto L1
L6:
	;
	return int32(0)
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = l8
	F_errmsg_internal(m, int32(_a_F_construct_md_array_0), v21)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	F_errfinish(m, int32(_a_F_construct_md_array_1), int32(322), int32(_a_F_construct_md_array_2))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L10:
	;
	m.G0 = v21 - int32(-64)
	return v265
L11:
	;
	v231 = v225 + v228
	v232 = F_palloc0(m, v231)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L6
	} else {
		goto L68
	}
L12:
	;
	v217 = base.I32_div_s(v49+int32(7), int32(8))
	v224 = (v217 + l2<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	v225 = v224
	v228 = v212
	v230 = v224
	goto L11
L13:
	;
	if base.Ui32(l2) < base.Ui32(int32(7)) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L6
	} else {
		goto L64
	}
L16:
	;
	v49 = F_ArrayGetNItemsSafe(m, l2, l3)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L6
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L6
	} else {
		goto L60
	}
L19:
	;
	F_ArrayCheckBounds(m, l2, l3, l4)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	if int32(0) < v49 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v55 = int32(0)
	v69 = v55
	v74 = v10
	v77 = v10
	goto L24
L22:
	;
	goto L23
L23:
	;
	v164 = F_palloc0(m, int32(16))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L6
	} else {
		goto L59
	}
L24:
	;
	if l1 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v212 = v74
	goto L12
L26:
	;
	v159 = int32(1)
	v161 = v69 + v159
	if v161 != v49 {
		v69 = v161
		v77 = v159
		goto L24
	} else {
		goto L58
	}
L27:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v69))))
	if v79 != 0 {
		goto L26
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if l6 != int32(-1) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L29
L31:
	;
	v127 = (v74 + (v44 - int32(1)) + v123) & (v55 - v44)
	if base.Ui32(v127) < base.Ui32(int32(1073741824)) {
		goto L49
	} else {
		goto L50
	}
L32:
	;
	if int32(0) < l6 {
		v123 = l6
		goto L31
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v93 = l0 + v69<<(uint(int32(3))%32)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	v95 = F_pg_detoast_datum(m, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L6
	} else {
		goto L36
	}
L35:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0+v69<<(uint(int32(3))%32))))
	v88 = F_strlen(m, v87)
	mBase = m.M
	v123 = v88 + int32(1)
	goto L31
L36:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v93))) = base.I64_extend_i32_u(v95)
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	if v99 == int32(1) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v103 = int32(18)
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+1)))
	if v105 == v103 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	if v99&int32(1) != 0 {
		goto L46
	} else {
		goto L47
	}
L40:
	;
	v108 = v103
	goto L42
L41:
	;
	v108 = int32(2)
	goto L42
L42:
	;
	if base.Ui32((v105-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v115 = int32(6)
	goto L45
L44:
	;
	v115 = v108
	goto L45
L45:
	;
	v123 = v115
	goto L31
L46:
	;
	v123 = int32(base.Ui32(v99) >> (uint(int32(1)) % 32))
	goto L31
L47:
	;
	goto L48
L48:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v123 = int32(base.Ui32(v120) >> (uint(int32(2)) % 32))
	goto L31
L49:
	;
	v131 = v69 + int32(1)
	if v131 != v49 {
		v69 = v131
		v74 = v127
		goto L24
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L6
	} else {
		goto L54
	}
L52:
	;
	if v77 != 0 {
		v212 = v127
		goto L12
	} else {
		goto L53
	}
L53:
	;
	v225 = (l2<<(uint(int32(3))%32) + int32(23)) & int32(120)
	v228 = v127
	v230 = int32(0)
	goto L11
L54:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = int32(1073741823)
	F_errmsg(m, int32(_a_F_construct_md_array_3), v19+int32(-16))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L6
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_construct_md_array_4), int32(3553), int32(_a_F_construct_md_array_5))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L6
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	goto L25
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v164)+12)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v164)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v164))) = int64(64)
	v265 = v164
	goto L10
L60:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = l2
	F_errmsg(m, int32(_a_F_construct_md_array_6), v19+int32(-32))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L6
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_construct_md_array_4), int32(3523), int32(_a_F_construct_md_array_5))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L6
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L6
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = l2
	F_errmsg(m, int32(_a_F_construct_md_array_7), v19+int32(-48))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L6
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_construct_md_array_4), int32(3518), int32(_a_F_construct_md_array_5))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L6
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+12)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v232)+8)) = v230
	*(*int32)(unsafe.Add(mBase, uint32(v232)+4)) = l2
	v237 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v232))) = v231 << (uint(v237) % 32)
	v241 = v232 + int32(16)
	v243 = l2 << (uint(v237) % 32)
	v244 = int32(0)
	v245 = base.B2i32(v243 == v244)
	if v245 == v244 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	base.MemoryCopy(m, v241, l3, v243)
	goto L71
L70:
	;
	goto L71
L71:
	;
	if v245 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	base.MemoryCopy(m, v243+v241, l4, v243)
	goto L74
L73:
	;
	goto L74
L74:
	;
	F_CopyArrayEls(m, v232, l0, l1, v49, l6, l7, l8, int32(0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L6
	} else {
		goto L75
	}
L75:
	;
	v265 = v232
	goto L10
}
func F_convert_tuples_by_position(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v304 int32
	_ = v304
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v375 int32
	_ = v375
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	v4 = int32(0)
	v18 = m.G0
	v20 = v18 + int32(-64)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v24 = F_palloc0(m, int32(8))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v375 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L2:
	;
	return int32(0)
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v22
	v30 = F_palloc0_mul(m, int32(2), v22)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v30
	v33 = int32(1)
	if v22 <= int32(0) {
		v197 = v4
		v201 = v33
		v203 = v4
		v207 = v30
		v209 = v4
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v211 <= v197 {
		v280 = v201
		v282 = v203
		goto L34
	} else {
		goto L35
	}
L6:
	;
	v38 = v4
	v42 = v33
	v44 = v4
	v46 = v4
	v48 = v30
	v50 = v4
	goto L7
L7:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v58 = l1 + v52<<(uint(int32(3))%32) + v46*int32(100)
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+119)))
	if v59 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L2
	} else {
		goto L27
	}
L9:
	;
	goto L8
L10:
	;
	v63 = v50 + int32(1)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v64 <= v38 {
		v117 = v38
		v123 = v44
		v127 = v48
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v140 = v38
	v144 = v42
	v146 = v44
	v150 = v48
	v152 = v50
	goto L12
L12:
	;
	v155 = v46 + int32(1)
	if v22 != v155 {
		v38 = v140
		v42 = v144
		v44 = v146
		v46 = v155
		v48 = v150
		v50 = v152
		goto L7
	} else {
		goto L26
	}
L13:
	;
	v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v127+v46<<(uint(int32(1))%32)))))
	v140 = v117
	v144 = base.B2i32(v134 != int32(0)) & v42
	v146 = v123
	v150 = v127
	v152 = v63
	goto L12
L14:
	;
	v66 = int32(28)
	v67 = v58 + v66
	v75 = v38
	goto L15
L15:
	;
	v91 = l0 + v64<<(uint(int32(3))%32) + v66 + v75*int32(100)
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+91)))
	if v92 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v117 = v64
	v123 = v44
	v127 = v48
	goto L13
L17:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v67)+68))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v91)+68))
	if v95 != v96 {
		goto L9
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v113 = v75 + int32(1)
	if v113 != v64 {
		v75 = v113
		goto L15
	} else {
		goto L25
	}
L20:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v67)+76))
	if int32(0) <= v98 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v91)+76))
	if v98 != v101 {
		goto L9
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v103 = int32(1)
	v109 = v75 + v103
	*(*uint16)(unsafe.Add(mBase, uint32(v48+v46<<(uint(v103)%32)))) = uint16(v109)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v117 = v109
	v123 = v44 + v103
	v127 = v111
	goto L13
L24:
	;
	goto L23
L25:
	;
	goto L16
L26:
	;
	v197 = v140
	v201 = v144
	v203 = v146
	v207 = v150
	v209 = v152
	goto L5
L27:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L2
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = l2
	F_errmsg_internal(m, int32(_a_F_convert_tuples_by_position_0), v18+int32(-16))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L2
	} else {
		goto L29
	}
L29:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v91)+68))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v91)+76))
	v173 = F_format_type_with_typemod(m, v171, v172)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L2
	} else {
		goto L30
	}
L30:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v67)+68))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v67)+76))
	v177 = F_format_type_with_typemod(m, v175, v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v58 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v173
	v188 = F_errdetail(m, int32(_a_F_convert_tuples_by_position_1), v18+int32(-32))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_convert_tuples_by_position_2), int32(124), int32(_a_F_convert_tuples_by_position_3))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L34:
	;
	if v280&int32(1) != 0 {
		goto L43
	} else {
		goto L44
	}
L35:
	;
	v213 = int32(1)
	v214 = v197 + v213
	if (v211-v197)&v213 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v197<<(uint(int32(3))%32))+34)))
	v223 = v221 & int32(4)
	v230 = v214
	v231 = int32(base.Ui32(v223)>>(uint(int32(2))%32)) & v201
	v232 = v203 + base.B2i32(v223 == int32(0))
	goto L38
L37:
	;
	v230 = v197
	v231 = v201
	v232 = v203
	goto L38
L38:
	;
	if v214 == v211 {
		v280 = v231
		v282 = v232
		goto L34
	} else {
		goto L39
	}
L39:
	;
	v236 = v230
	v240 = v231
	v242 = v232
	goto L40
L40:
	;
	v252 = l0 + v236<<(uint(int32(3))%32)
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252)+42)))
	v254 = int32(2)
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252)+34)))
	v257 = int32(4)
	v260 = int32(base.Ui32(v256&v257) >> (uint(v254) % 32))
	v262 = int32(base.Ui32(v253)>>(uint(v254)%32)) & v260 & v240
	v265 = int32(0)
	v270 = base.B2i32(v253&v257 == v265) + (v242 + base.B2i32(v260 == v265))
	v272 = v236 + v254
	if v272 != v211 {
		v236 = v272
		v240 = v262
		v242 = v270
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v280 = v262
	v282 = v270
	goto L34
L42:
	;
	goto L41
L43:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v211 != v292 {
		v375 = v24
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L45
L45:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L2
	} else {
		goto L63
	}
L46:
	;
	m.G0 = v20 - int32(-64)
	goto L1
L47:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if int32(0) < v294 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v297 = int32(28)
	v304 = int32(0)
	goto L51
L49:
	;
	goto L50
L50:
	;
	F_pfree(m, v207)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L2
	} else {
		goto L61
	}
L51:
	;
	v319 = v304 << (uint(int32(3)) % 32)
	v320 = l0 + v297 + v319
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320)+6)))
	if v321&int32(2) != 0 {
		v375 = v24
		goto L46
	} else {
		goto L53
	}
L52:
	;
	goto L50
L53:
	;
	v324 = int32(1)
	v325 = v304 + v324
	v329 = int32(*(*int16)(unsafe.Add(mBase, uint32(v207+v304<<(uint(v324)%32)))))
	if v325 != v329 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	if base.B2i32(v321&int32(4) == int32(0))|v329 != 0 {
		v375 = v24
		goto L46
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	if v325 != v294 {
		v304 = v325
		goto L51
	} else {
		goto L60
	}
L57:
	;
	v336 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v320)+2)))
	v337 = l1 + v297 + v319
	v338 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v337)+2)))
	if v336 != v338 {
		v375 = v24
		goto L46
	} else {
		goto L58
	}
L58:
	;
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320)+5)))
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337)+5)))
	if v340 != v341 {
		v375 = v24
		goto L46
	} else {
		goto L59
	}
L59:
	;
	goto L56
L60:
	;
	goto L52
L61:
	;
	F_pfree(m, v24)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L2
	} else {
		goto L62
	}
L62:
	;
	v375 = int32(0)
	goto L46
L63:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L2
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = l2
	F_errmsg_internal(m, int32(_a_F_convert_tuples_by_position_0), v18+int32(-48))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L2
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v209
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v282
	v401 = F_errdetail(m, int32(_a_F_convert_tuples_by_position_4), v20)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L2
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_convert_tuples_by_position_2), int32(149), int32(_a_F_convert_tuples_by_position_3))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L2
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	return int32(0)
L69:
	;
	goto L70
L70:
	;
	v413 = F_palloc(m, int32(28))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L2
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v413)+8)) = v375
	*(*int32)(unsafe.Add(mBase, uint32(v413)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v413))) = l0
	v419 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v421 = v419 + int32(1)
	v422 = F_palloc_mul(m, int32(8), v421)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L2
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v413)+20)) = v422
	v426 = F_palloc_mul(m, int32(1), v421)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L2
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v413)+24)) = v426
	v430 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v432 = v430 + int32(1)
	v433 = F_palloc_mul(m, int32(8), v432)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L2
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v413)+12)) = v433
	v437 = F_palloc_mul(m, int32(1), v432)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L2
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v413)+16)) = v437
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v413)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v440))) = int64(0)
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v413)+16))
	v444 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v443))) = uint8(v444)
	return v413
}
func F_copy_addr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	v2 = l1
	if v2 != int32(10) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return
L2:
	;
	if base.Ui32(l4) < base.Ui32(v31) {
		goto L1
	} else {
		goto L12
	}
L3:
	;
	if v2 != int32(2) {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v17 = l2 + int32(8)
	v18 = int32(16)
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	switch v19 - int32(254) {
	case 0:
		goto L9
	case 1:
		goto L8
	default:
		v31 = v18
		v32 = v17
		goto L2
	}
L6:
	;
	v13 = int32(4)
	v31 = v13
	v32 = l2 + v13
	goto L2
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = l5
	v31 = v18
	v32 = v17
	goto L2
L8:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+1)))
	if v25&int32(15) != int32(2) {
		v31 = v18
		v32 = v17
		goto L2
	} else {
		goto L11
	}
L9:
	;
	v22 = int32(*(*int8)(unsafe.Add(mBase, uint32(l3)+1)))
	if v22 < int32(-64) {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v31 = v18
	v32 = v17
	goto L2
L11:
	;
	goto L7
L12:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l2))) = uint16(v2)
	if base.Ui32(int32(512)) <= base.Ui32(v31) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l2
	goto L1
L14:
	;
	if v31 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v41 = v32 + v31
	if (v32^l3)&int32(3) == int32(0) {
		goto L21
	} else {
		goto L22
	}
L17:
	;
	base.MemoryCopy(m, v32, l3, v31)
	goto L19
L18:
	;
	goto L19
L19:
	;
	goto L13
L20:
	;
	if base.Ui32(v173) < base.Ui32(v41) {
		goto L54
	} else {
		goto L55
	}
L21:
	;
	if v32&int32(3) == int32(0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	goto L23
L23:
	;
	if base.Ui32(v41) < base.Ui32(int32(4)) {
		goto L45
	} else {
		goto L46
	}
L24:
	;
	v77 = v41 & int32(-4)
	if base.Ui32(v41) < base.Ui32(int32(64)) {
		v127 = v71
		v128 = v72
		goto L35
	} else {
		goto L36
	}
L25:
	;
	v71 = l3
	v72 = v32
	goto L24
L26:
	;
	goto L27
L27:
	;
	if v31 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v71 = l3
	v72 = v32
	goto L24
L29:
	;
	goto L30
L30:
	;
	v54 = l3
	v55 = v32
	goto L31
L31:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	*(*uint8)(unsafe.Add(mBase, uint32(v55))) = uint8(v59)
	v61 = int32(1)
	v62 = v54 + v61
	v64 = v55 + v61
	if v64&int32(3) == int32(0) {
		v71 = v62
		v72 = v64
		goto L24
	} else {
		goto L33
	}
L32:
	;
	v71 = v62
	v72 = v64
	goto L24
L33:
	;
	if base.Ui32(v64) < base.Ui32(v41) {
		v54 = v62
		v55 = v64
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	if base.Ui32(v77) <= base.Ui32(v128) {
		v172 = v127
		v173 = v128
		goto L20
	} else {
		goto L41
	}
L36:
	;
	v81 = v77 + int32(-64)
	if base.Ui32(v81) < base.Ui32(v72) {
		v127 = v71
		v128 = v72
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v84 = v71
	v85 = v72
	goto L38
L38:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	*(*int32)(unsafe.Add(mBase, uint32(v85))) = v89
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+4)) = v91
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+8)) = v93
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+12)) = v95
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v84)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+16)) = v97
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v84)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+20)) = v99
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v84)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+24)) = v101
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v84)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+28)) = v103
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v84)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+32)) = v105
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v84)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+36)) = v107
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v84)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+40)) = v109
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v84)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+44)) = v111
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v84)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+48)) = v113
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v84)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+52)) = v115
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v84)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+56)) = v117
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v84)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+60)) = v119
	v121 = int32(-64)
	v122 = v84 - v121
	v124 = v85 - v121
	if base.Ui32(v124) <= base.Ui32(v81) {
		v84 = v122
		v85 = v124
		goto L38
	} else {
		goto L40
	}
L39:
	;
	v127 = v122
	v128 = v124
	goto L35
L40:
	;
	goto L39
L41:
	;
	v134 = v127
	v135 = v128
	goto L42
L42:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v139
	v141 = int32(4)
	v142 = v134 + v141
	v144 = v135 + v141
	if base.Ui32(v144) < base.Ui32(v77) {
		v134 = v142
		v135 = v144
		goto L42
	} else {
		goto L44
	}
L43:
	;
	v172 = v142
	v173 = v144
	goto L20
L44:
	;
	goto L43
L45:
	;
	v172 = l3
	v173 = v32
	goto L20
L46:
	;
	goto L47
L47:
	;
	if base.Ui32(v31) < base.Ui32(int32(4)) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v172 = l3
	v173 = v32
	goto L20
L49:
	;
	goto L50
L50:
	;
	v153 = l3
	v154 = v32
	goto L51
L51:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	*(*uint8)(unsafe.Add(mBase, uint32(v154))) = uint8(v158)
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v154)+1)) = uint8(v160)
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v154)+2)) = uint8(v162)
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v154)+3)) = uint8(v164)
	v166 = int32(4)
	v167 = v153 + v166
	v169 = v154 + v166
	if base.Ui32(v169) <= base.Ui32(v41-int32(4)) {
		v153 = v167
		v154 = v169
		goto L51
	} else {
		goto L53
	}
L52:
	;
	v172 = v167
	v173 = v169
	goto L20
L53:
	;
	goto L52
L54:
	;
	v179 = v172
	v180 = v173
	goto L57
L55:
	;
	goto L56
L56:
	;
	goto L13
L57:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
	*(*uint8)(unsafe.Add(mBase, uint32(v180))) = uint8(v184)
	v186 = int32(1)
	v189 = v180 + v186
	if v189 != v41 {
		v179 = v179 + v186
		v180 = v189
		goto L57
	} else {
		goto L59
	}
L58:
	;
	goto L56
L59:
	;
	goto L58
}
func F_copytup_heap(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	v5 = F_minimal_tuple_from_heap_tuple(m, l1, int32(0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = F_GetMemoryChunkSpace(m, v5)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v11 - base.I64_extend_i32_u(v9)
			return v5
		}
	}
}
func F_cost_material(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64, l4 float64, l5 float64, l6 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 float64
	_ = v14
	var v18 float64
	_ = v18
	var v26 float64
	_ = v26
	var v32 float64
	_ = v32
	var v38 float64
	_ = v38
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_cost_material[0]))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = l5
	v14 = *(*float64)(unsafe.Add(mBase, _c_F_cost_material[1]))
	v18 = base.F64_add(base.F64_mul(base.F64_add(v14, v14), l5), base.F64_sub(l4, l3))
	v26 = base.F64_mul(l5, base.F64_convert_i32_u((l6+int32(7))&int32(-8)+int32(24)))
	if base.F64_gt(v26, base.F64_convert_i32_u(v11<<(uint(int32(10))%32))) != 0 {
		v32 = *(*float64)(unsafe.Add(mBase, _c_F_cost_material[2]))
		v38 = base.F64_add(base.F64_mul(v32, base.F64_ceil(base.F64_mul(v26, float64(0.0001220703125)))), v18)
	} else {
		v38 = v18
	}
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = l3
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = base.F64_add(l3, v38)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l2 + (l1 ^ int32(1))
	return
}
func F_cost_samplescan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 float64
	_ = v9
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 float64
	_ = v46
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 float64
	_ = v57
	var v58 float64
	_ = v58
	var v59 int32
	_ = v59
	var v60 int64
	_ = v60
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 float64
	_ = v102
	var v103 float64
	_ = v103
	var v113 float64
	_ = v113
	var v120 float64
	_ = v120
	var v121 float64
	_ = v121
	var v123 float64
	_ = v123
	var v125 float64
	_ = v125
	var v126 float64
	_ = v126
	var v136 float64
	_ = v136
	var v143 float64
	_ = v143
	var v144 int32
	_ = v144
	var v145 float64
	_ = v145
	var v146 float64
	_ = v146
	var v148 float64
	_ = v148
	var v149 int64
	_ = v149
	var v152 float64
	_ = v152
	var v153 float64
	_ = v153
	var v157 int32
	_ = v157
	var v158 int64
	_ = v158
	var v163 float64
	_ = v163
	var v167 float64
	_ = v167
	v9 = float64(0)
	v17 = m.G0
	v19 = v17 - int32(48)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	if v21 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+32))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v39 = F_GetTsmRoutine(m, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
	v35 = v21 + v22<<(uint(int32(2))%32)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+52))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
	v35 = v28 + v29<<(uint(int32(2))%32) - int32(4)
	goto L1
L5:
	;
	return
L6:
	;
	if l3 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v45 = l3 + int32(8)
	goto L9
L8:
	;
	v45 = l2 + int32(16)
	goto L9
L9:
	;
	v46 = *(*float64)(unsafe.Add(mBase, uint32(v45)))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v46
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l2)+80))
	F_get_tablespace_page_costs(m, v48, v19+int32(8), v19+int32(16))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)+124))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v39)+24))
	v57 = *(*float64)(unsafe.Add(mBase, uint32(v19)+16))
	v58 = *(*float64)(unsafe.Add(mBase, uint32(v19)+8))
	if l3 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v145 = *(*float64)(unsafe.Add(mBase, uint32(v144)+24))
	v146 = *(*float64)(unsafe.Add(mBase, uint32(l2)+128))
	v148 = *(*float64)(unsafe.Add(mBase, _c_F_cost_samplescan[0]))
	v149 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
	v152 = *(*float64)(unsafe.Add(mBase, uint32(v144)+16))
	v153 = base.F64_add(base.F64_add(v143, float64(0)), v152)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = v153
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v157 != 0 {
		goto L22
	} else {
		goto L23
	}
L12:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v60 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+32)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v19)+40)) = v60
	if v59 == int32(0) {
		v113 = v9
		v120 = float64(0)
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v125 = *(*float64)(unsafe.Add(mBase, uint32(l2)+216))
	v126 = *(*float64)(unsafe.Add(mBase, uint32(l2)+208))
	v136 = v125
	v143 = v126
	goto L11
L15:
	;
	v121 = *(*float64)(unsafe.Add(mBase, uint32(l2)+216))
	v123 = *(*float64)(unsafe.Add(mBase, uint32(l2)+208))
	v136 = base.F64_add(v113, v121)
	v143 = base.F64_add(v120, v123)
	goto L11
L16:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v69 <= int32(0) {
		v113 = v9
		v120 = float64(0)
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v76 = int32(0)
	goto L18
L18:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89+v76<<(uint(int32(2))%32))))
	v96 = F_cost_qual_eval_walker(m, v93, v19+int32(24))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L5
	} else {
		goto L20
	}
L19:
	;
	v102 = *(*float64)(unsafe.Add(mBase, uint32(v19)+40))
	v103 = *(*float64)(unsafe.Add(mBase, uint32(v19)+32))
	v113 = v102
	v120 = v103
	goto L15
L20:
	;
	v99 = v76 + int32(1)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v99 < v100 {
		v76 = v99
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v158 = int64(-1)
	goto L24
L23:
	;
	v158 = int64(-262145)
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = base.B2i32(v149|v158 != int64(-1))
	v163 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
	if v56 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v167 = v58
	goto L27
L26:
	;
	v167 = v57
	goto L27
L27:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = base.F64_add(v153, base.F64_add(base.F64_mul(v145, v163), base.F64_add(base.F64_mul(v146, base.F64_add(v136, v148)), base.F64_add(base.F64_mul(v167, base.F64_convert_i32_u(v55)), float64(0)))))
	m.G0 = v19 + int32(48)
	return
}
func F_cost_subqueryscan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v9 float64
	_ = v9
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
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
	var v29 float64
	_ = v29
	var v30 int32
	_ = v30
	var v33 float64
	_ = v33
	var v34 int32
	_ = v34
	var v35 float64
	_ = v35
	var v44 float64
	_ = v44
	var v48 float64
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v55 int32
	_ = v55
	var v56 int64
	_ = v56
	var v62 float64
	_ = v62
	var v64 float64
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int64
	_ = v71
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 float64
	_ = v113
	var v114 float64
	_ = v114
	var v115 float64
	_ = v115
	var v116 int32
	_ = v116
	var v117 float64
	_ = v117
	var v118 float64
	_ = v118
	var v124 int32
	_ = v124
	var v127 float64
	_ = v127
	var v128 float64
	_ = v128
	var v130 float64
	_ = v130
	var v131 float64
	_ = v131
	var v135 float64
	_ = v135
	var v136 float64
	_ = v136
	var v138 float64
	_ = v138
	var v140 float64
	_ = v140
	var v141 float64
	_ = v141
	var v147 int32
	_ = v147
	var v150 float64
	_ = v150
	var v151 float64
	_ = v151
	var v153 float64
	_ = v153
	var v154 float64
	_ = v154
	var v158 float64
	_ = v158
	var v159 int32
	_ = v159
	var v160 float64
	_ = v160
	var v161 float64
	_ = v161
	var v163 float64
	_ = v163
	var v164 float64
	_ = v164
	var v165 float64
	_ = v165
	v9 = float64(0)
	v17 = m.G0
	v19 = v17 - int32(32)
	m.G0 = v19
	if l3 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v29 = *(*float64)(unsafe.Add(mBase, uint32(v28)+32))
	v30 = int32(0)
	v33 = F_clauselist_selectivity(m, l1, v26, v30, v30, v30)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L5
	} else {
		goto L8
	}
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)+204))
	v23 = F_list_concat_copy(m, v21, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)+204))
	v26 = v25
	goto L1
L5:
	;
	return
L6:
	;
	v26 = v23
	goto L1
L7:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v48
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+40))
	v52 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v55 != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v35 = base.F64_mul(v29, v33)
	if base.F64_gt(v35, float64(1e+100))|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v35)&int64(9223372036854775807))) != 0 {
		v48 = float64(1e+100)
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v44 = float64(1)
	if base.F64_le(v35, v44) != 0 {
		v48 = v44
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v48 = base.F64_nearest(v35)
	goto L7
L11:
	;
	v56 = int64(-1)
	goto L13
L12:
	;
	v56 = int64(-262145)
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v51 + base.B2i32(v52|v56 != int64(-1))
	v62 = *(*float64)(unsafe.Add(mBase, uint32(v50)+48))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = v62
	v64 = *(*float64)(unsafe.Add(mBase, uint32(v50)+56))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = v64
	if v26 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v67 = int32(0)
	goto L16
L15:
	;
	v67 = l4
	goto L16
L16:
	;
	if v67 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	if l3 != 0 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	goto L19
L19:
	;
	m.G0 = v19 + int32(32)
	return
L20:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v160 = *(*float64)(unsafe.Add(mBase, uint32(v159)+24))
	v161 = *(*float64)(unsafe.Add(mBase, uint32(v147)+32))
	v163 = *(*float64)(unsafe.Add(mBase, _c_F_cost_subqueryscan[0]))
	v164 = *(*float64)(unsafe.Add(mBase, uint32(v159)+16))
	v165 = base.F64_add(v158, v164)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = base.F64_add(v165, v150)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = base.F64_add(base.F64_add(v165, base.F64_add(base.F64_mul(v160, v151), base.F64_mul(v161, base.F64_add(v153, v163)))), v154)
	goto L19
L21:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v71 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v71
	if v70 == int32(0) {
		v124 = v50
		v127 = v62
		v128 = v48
		v130 = v9
		v131 = v64
		v135 = float64(0)
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v140 = *(*float64)(unsafe.Add(mBase, uint32(l2)+216))
	v141 = *(*float64)(unsafe.Add(mBase, uint32(l2)+208))
	v147 = v50
	v150 = v62
	v151 = v48
	v153 = v140
	v154 = v64
	v158 = v141
	goto L20
L24:
	;
	v136 = *(*float64)(unsafe.Add(mBase, uint32(l2)+216))
	v138 = *(*float64)(unsafe.Add(mBase, uint32(l2)+208))
	v147 = v124
	v150 = v127
	v151 = v128
	v153 = base.F64_add(v130, v136)
	v154 = v131
	v158 = base.F64_add(v135, v138)
	goto L20
L25:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	if v80 <= int32(0) {
		v124 = v50
		v127 = v62
		v128 = v48
		v130 = v9
		v131 = v64
		v135 = float64(0)
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v89 = int32(0)
	goto L27
L27:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v100+v89<<(uint(int32(2))%32))))
	v107 = F_cost_qual_eval_walker(m, v104, v19+int32(8))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L5
	} else {
		goto L29
	}
L28:
	;
	v113 = *(*float64)(unsafe.Add(mBase, uint32(l0)+56))
	v114 = *(*float64)(unsafe.Add(mBase, uint32(l0)+48))
	v115 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v117 = *(*float64)(unsafe.Add(mBase, uint32(v19)+24))
	v118 = *(*float64)(unsafe.Add(mBase, uint32(v19)+16))
	v124 = v116
	v127 = v114
	v128 = v115
	v130 = v117
	v131 = v113
	v135 = v118
	goto L24
L29:
	;
	v110 = v89 + int32(1)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	if v110 < v111 {
		v89 = v110
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
}
