package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ConditionVariableSignal(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
	if v7 != 0 {
		F_s_lock(m, l0, int32(_a_F_ConditionVariableSignal_0), int32(264), int32(_a_F_ConditionVariableSignal_1))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v15 == int32(-1) {
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
				return
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableSignal[0]))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				v25 = v22 + v15*int32(640)
				v27 = v25 + int32(84)
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+88))
				if v29 == int32(-1) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v28
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
					v38 = v33
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v22+v29*int32(640))+84)) = v28
					v38 = v29
				}
				if v28 == int32(-1) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v38
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableSignal[0]))
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
					*(*int32)(unsafe.Add(mBase, uint32(v44+v28*int32(640))+88)) = v38
				}
				*(*int64)(unsafe.Add(mBase, uint32(v27))) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
				v55 = v22 + v15*int32(640)
				if v55 != 0 {
					F_SetLatch(m, v55+int32(20))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return
					} else {
						return
					}
				} else {
					return
				}
			}
		}
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v15 == int32(-1) {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
			return
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableSignal[0]))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
			v25 = v22 + v15*int32(640)
			v27 = v25 + int32(84)
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+88))
			if v29 == int32(-1) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v28
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
				v38 = v33
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v22+v29*int32(640))+84)) = v28
				v38 = v29
			}
			if v28 == int32(-1) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v38
			} else {
				v43 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableSignal[0]))
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
				*(*int32)(unsafe.Add(mBase, uint32(v44+v28*int32(640))+88)) = v38
			}
			*(*int64)(unsafe.Add(mBase, uint32(v27))) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
			v55 = v22 + v15*int32(640)
			if v55 != 0 {
				F_SetLatch(m, v55+int32(20))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return
				} else {
					return
				}
			} else {
				return
			}
		}
	}
}
func F_ConditionalLockBuffer(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	if l0 < int32(0) {
		return int32(1)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBuffer[0]))
		v14 = F_LWLockConditionalAcquire(m, v7+l0<<(uint(int32(6))%32)-int32(16), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			return v14
		}
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
	var v35 int32
	_ = v35
	var v51 int32
	_ = v51
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v11 = int32(1)
	if l0 <= int32(3591) {
		if l0 <= int32(2670) {
			switch l0 - int32(1213) {
			case 0, 1, 19, 20, 47, 48, 49:
				v79 = v11
			case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
				v79 = int32(0)
			default:
				if base.Ui32(int32(2)) <= base.Ui32(l0-int32(2396)) {
					v79 = int32(0)
				} else {
					v79 = v11
				}
			}
		} else {
			v23 = l0 - int32(2671)
			if base.Ui32(int32(27)) < base.Ui32(v23) {
				if base.Ui32(l0-int32(2964)) < base.Ui32(int32(4)) {
					v79 = v11
				} else {
					if base.Ui32(l0-int32(2846)) < base.Ui32(int32(2)) {
						v79 = v11
					} else {
						v79 = int32(0)
					}
				}
			} else {
				if int32(1)<<(uint(v23)%32)&int32(226492515) == int32(0) {
					if base.Ui32(l0-int32(2964)) < base.Ui32(int32(4)) {
						v79 = v11
					} else {
						if base.Ui32(l0-int32(2846)) < base.Ui32(int32(2)) {
							v79 = v11
						} else {
							v79 = int32(0)
						}
					}
				} else {
					v79 = v11
				}
			}
		}
	} else {
		if l0 <= int32(_a_F_ConditionalLockRelationOid_0) {
			v35 = l0 - int32(_a_F_ConditionalLockRelationOid_1)
			if base.Ui32(int32(9)) < base.Ui32(v35) {
				if base.Ui32(l0-int32(3592)) < base.Ui32(int32(2)) {
					v79 = v11
				} else {
					if base.Ui32(int32(2)) <= base.Ui32(l0-int32(4060)) {
						v79 = int32(0)
					} else {
						v79 = v11
					}
				}
			} else {
				if int32(1)<<(uint(v35)%32)&int32(963) == int32(0) {
					if base.Ui32(l0-int32(3592)) < base.Ui32(int32(2)) {
						v79 = v11
					} else {
						if base.Ui32(int32(2)) <= base.Ui32(l0-int32(4060)) {
							v79 = int32(0)
						} else {
							v79 = v11
						}
					}
				} else {
					v79 = v11
				}
			}
		} else {
			switch l0 - int32(_a_F_ConditionalLockRelationOid_2) {
			case 0, 1, 2, 3, 4, 59, 60:
				v79 = v11
			case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
				v79 = int32(0)
			default:
				if base.Ui32(l0-int32(_a_F_ConditionalLockRelationOid_3)) < base.Ui32(int32(3)) {
					v79 = v11
				} else {
					v51 = l0 - int32(_a_F_ConditionalLockRelationOid_4)
					if base.Ui32(int32(15)) < base.Ui32(v51) {
						v79 = int32(0)
					} else {
						if int32(1)<<(uint(v51)%32)&int32(_a_F_ConditionalLockRelationOid_5) != 0 {
							v79 = v11
						} else {
							v79 = int32(0)
						}
					}
				}
			}
		}
	}
	*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = int64(72057594037927936)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l0
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockRelationOid[0]))
	if v79 != 0 {
		v86 = int32(0)
	} else {
		v86 = v85
	}
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v86
	v90 = int32(0)
	v95 = F_LockAcquireExtended(m, v7+int32(16), l1, v90, int32(1), v7+int32(12), v90)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		return int32(0)
	} else {
		switch v95 {
		case 0, 3:
			m.G0 = v7 + int32(32)
			return base.B2i32(v95 != int32(0))
		default:
			F_ReceiveSharedInvalidMessages(m)
			mBase = m.M
			v100 = m.ExcPending
			if v100 != 0 {
				return int32(0)
			} else {
				v101 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
				v102 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v101)+53)) = uint8(v102)
				m.G0 = v7 + int32(32)
				return base.B2i32(v95 != int32(0))
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
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
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
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
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	v3 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v15 = F_SearchSysCache1(m, int32(20), l0)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return v155
L2:
	;
	return int32(0)
L3:
	;
	if v15 == int32(0) {
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
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+22)))
	F_recomputeNamespacePath(m)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L2
	} else {
		goto L13
	}
L7:
	;
	v21 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v21)
	v155 = int32(0)
	goto L1
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
	F_errmsg_internal(m, int32(_a_F_ConversionIsVisibleExt_0), v12)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(_a_F_ConversionIsVisibleExt_1), int32(2536), int32(_a_F_ConversionIsVisibleExt_2))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
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
	v41 = v37 + v38
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+68))
	if v42 != int32(11) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	F_ReleaseCatCache(m, v15)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L2
	} else {
		goto L44
	}
L15:
	;
	v45 = int32(0)
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_ConversionIsVisibleExt[0]))
	if v47 == v45 {
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
	v91 = m.ExcPending
	if v91 != 0 {
		goto L2
	} else {
		goto L32
	}
L18:
	;
	if v86 == int32(0) {
		v144 = v45
		goto L14
	} else {
		goto L31
	}
L19:
	;
	v86 = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	if v54 <= int32(0) {
		v79 = v45
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v86 = v79
	goto L18
L23:
	;
	v57 = int32(0)
	if v57 < v54 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v60 = v54
	goto L26
L25:
	;
	v60 = v57
	goto L26
L26:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	v63 = int32(0)
	goto L27
L27:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v61+v63<<(uint(int32(2))%32))))
	v72 = base.B2i32(v71 == v42)
	if v71 == v42 {
		v79 = v72
		goto L22
	} else {
		goto L29
	}
L28:
	;
	v79 = v72
	goto L22
L29:
	;
	v74 = v63 + int32(1)
	if v74 != v60 {
		v63 = v74
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
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_ConversionIsVisibleExt[0]))
	if v93 == int32(0) {
		v136 = v3
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v144 = base.B2i32(l0 == v136)
	goto L14
L34:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	if v96 <= int32(0) {
		v136 = v3
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_ConversionIsVisibleExt[1]))
	v105 = int32(0)
	v107 = v102
	v111 = v96
	goto L36
L36:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v113+v105<<(uint(int32(2))%32))))
	if v107 != v117 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v136 = int32(0)
	goto L33
L38:
	;
	v120 = int32(0)
	v122 = F_GetSysCacheOid(m, int32(18), v41+int32(4), v117, v120, v120)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L2
	} else {
		goto L41
	}
L39:
	;
	v127 = v107
	v128 = v111
	goto L40
L40:
	;
	v130 = v105 + int32(1)
	if v130 < v128 {
		v105 = v130
		v107 = v127
		v111 = v128
		goto L36
	} else {
		goto L43
	}
L41:
	;
	if v122 != 0 {
		v136 = v122
		goto L33
	} else {
		goto L42
	}
L42:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_ConversionIsVisibleExt[1]))
	v127 = v126
	v128 = v124
	goto L40
L43:
	;
	goto L37
L44:
	;
	v155 = v144
	goto L1
}
func F_CopyGetAttnums(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v205 int32
	_ = v205
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	v4 = int32(0)
	v12 = m.G0
	v14 = v12 + int32(-64)
	m.G0 = v14
	if l2 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L18
	} else {
		goto L73
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L18
	} else {
		goto L69
	}
L3:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v205 + int32(4)
	F_errmsg(m, int32(_a_F_CopyGetAttnums_0), v12+int32(-48))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L18
	} else {
		goto L67
	}
L4:
	;
	m.G0 = v14 - int32(-64)
	return v195
L5:
	;
	v57 = v4
	v62 = v4
	goto L21
L6:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if int32(0) < v16 {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v19 <= int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v195 = v4
	goto L4
L10:
	;
	v195 = v4
	goto L4
L11:
	;
	goto L12
L12:
	;
	v27 = v4
	v29 = v4
	goto L13
L13:
	;
	v37 = l0 + int32(20) + v27<<(uint(int32(4))%32)
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+9)))
	if v38 != 0 {
		v46 = v29
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v195 = v46
	goto L4
L15:
	;
	v48 = v27 + int32(1)
	if v48 != v19 {
		v27 = v48
		v29 = v46
		goto L13
	} else {
		goto L20
	}
L16:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+10)))
	if v39 != 0 {
		v46 = v29
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v42 = F_lappend_int(m, v29, v27+int32(1))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return int32(0)
L19:
	;
	v46 = v42
	goto L15
L20:
	;
	goto L14
L21:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63+v62<<(uint(int32(2))%32))))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	v69 = int32(0)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v70 <= v69 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v195 = v184
	goto L4
L23:
	;
	v145 = int32(0)
	if v57 == v145 {
		goto L52
	} else {
		goto L53
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L18
	} else {
		goto L46
	}
L25:
	;
	v76 = v69
	v80 = v70
	goto L26
L26:
	;
	v89 = l0 + int32(20) + v80<<(uint(int32(4))%32) + v76*int32(100)
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+91)))
	if v90 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+90)))
	if v116 != 0 {
		goto L1
	} else {
		goto L44
	}
L28:
	;
	goto L27
L29:
	;
	v94 = v89 + int32(4)
	if v94|v68 != 0 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v112 = v80
	goto L31
L31:
	;
	v114 = v76 + int32(1)
	if v114 < v112 {
		v76 = v114
		v80 = v112
		goto L26
	} else {
		goto L43
	}
L32:
	;
	if v108 == int32(0) {
		goto L28
	} else {
		goto L42
	}
L33:
	;
	v100 = int32(-1)
	goto L35
L34:
	;
	v100 = int32(0)
	goto L35
L35:
	;
	if v94 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v101 = int32(1)
	goto L38
L37:
	;
	v101 = v100
	goto L38
L38:
	;
	if v94 == int32(0) {
		v108 = v101
		goto L39
	} else {
		goto L40
	}
L39:
	;
	goto L32
L40:
	;
	if v68 == int32(0) {
		v108 = v101
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v107 = F_strncmp(m, v94, v68, int32(64))
	mBase = m.M
	v108 = v107
	goto L39
L42:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v112 = v111
	goto L31
L43:
	;
	goto L24
L44:
	;
	v117 = int32(*(*int16)(unsafe.Add(mBase, uint32(v89)+74)))
	if v117 != 0 {
		goto L23
	} else {
		goto L45
	}
L45:
	;
	goto L24
L46:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L18
	} else {
		goto L47
	}
L47:
	;
	if l1 != 0 {
		goto L3
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v68
	F_errmsg(m, int32(_a_F_CopyGetAttnums_1), v14)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L18
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_CopyGetAttnums_2), int32(1045), int32(_a_F_CopyGetAttnums_3))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L18
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	if v183 != 0 {
		goto L2
	} else {
		goto L64
	}
L52:
	;
	v183 = int32(0)
	goto L51
L53:
	;
	goto L54
L54:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v151 <= int32(0) {
		v176 = v145
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v183 = v176
	goto L51
L56:
	;
	v154 = int32(0)
	if v154 < v151 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v157 = v151
	goto L59
L58:
	;
	v157 = v154
	goto L59
L59:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v160 = int32(0)
	goto L60
L60:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v158+v160<<(uint(int32(2))%32))))
	v169 = base.B2i32(v168 == v117)
	if v168 == v117 {
		v176 = v169
		goto L55
	} else {
		goto L62
	}
L61:
	;
	v176 = v169
	goto L55
L62:
	;
	v171 = v160 + int32(1)
	if v171 != v157 {
		v160 = v171
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v184 = F_lappend_int(m, v57, v117)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L18
	} else {
		goto L65
	}
L65:
	;
	v187 = v62 + int32(1)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v187 < v188 {
		v57 = v184
		v62 = v187
		goto L21
	} else {
		goto L66
	}
L66:
	;
	goto L22
L67:
	;
	F_errfinish(m, int32(_a_F_CopyGetAttnums_2), int32(1040), int32(_a_F_CopyGetAttnums_3))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L18
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	F_errcode(m, int32(16806020))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L18
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v68
	F_errmsg(m, int32(_a_F_CopyGetAttnums_4), v12+int32(-32))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L18
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_CopyGetAttnums_2), int32(1052), int32(_a_F_CopyGetAttnums_3))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L18
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	F_errcode(m, int32(_a_F_CopyGetAttnums_5))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L18
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v68
	F_errmsg(m, int32(_a_F_CopyGetAttnums_6), v12+int32(-16))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L18
	} else {
		goto L75
	}
L75:
	;
	F_errdetail(m, int32(_a_F_CopyGetAttnums_7), int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L18
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_CopyGetAttnums_2), int32(1029), int32(_a_F_CopyGetAttnums_3))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L18
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_CopyLimitPrintoutLength(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
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
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	if l0&int32(3) == int32(0) {
		v26 = l0
		goto L3
	} else {
		goto L4
	}
L1:
	;
	if v59 <= int32(100) {
		goto L18
	} else {
		goto L19
	}
L2:
	;
	v59 = v51 - l0
	goto L1
L3:
	;
	v30 = v26
	goto L12
L4:
	;
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v10 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v59 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v15 = l0
	goto L8
L8:
	;
	v19 = v15 + int32(1)
	if v19&int32(3) == int32(0) {
		v26 = v19
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v51 = v19
	goto L2
L10:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v24 != 0 {
		v15 = v19
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v39 = int32(-2139062144)
	if (int32(16843008)-v36|v36)&v39 == v39 {
		v30 = v30 + int32(4)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v45 = v30
	goto L15
L14:
	;
	goto L13
L15:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v49 != 0 {
		v45 = v45 + int32(1)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v51 = v45
	goto L2
L17:
	;
	goto L16
L18:
	;
	v62 = F_pstrdup(m, l0)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v68 = F_pg_mbcliplen(m, l0, v59, int32(100))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L21
	} else {
		goto L23
	}
L21:
	;
	return int32(0)
L22:
	;
	return v62
L23:
	;
	v72 = F_palloc(m, v68+int32(4))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	if v68 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75+v68))) = int32(_a_F_CopyLimitPrintoutLength_0)
	return v75
L26:
	;
	v74 = F__emscripten_memcpy_bulkmem(m, v72, l0, v68)
	mBase = m.M
	v75 = v74
	goto L28
L27:
	;
	v75 = v72
	goto L28
L28:
	;
	goto L25
}
func F_CopyReadAttributesText(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
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
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v184 int32
	_ = v184
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v219 int32
	_ = v219
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
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
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
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
	var v414 int32
	_ = v414
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v461 int32
	_ = v461
	v2 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	if v20 <= v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v18 + int32(16)
	return v461
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	if v23 == int32(0) {
		v461 = v2
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	v47 = l0 + int32(264)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v49 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v48))) = uint8(v49)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v49
	goto L11
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
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
	v34 = m.ExcPending
	if v34 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	F_errmsg(m, int32(_a_F_CopyReadAttributesText_0), int32(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	F_errfinish(m, int32(_a_F_CopyReadAttributesText_1), int32(1581), int32(_a_F_CopyReadAttributesText_2))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
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
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v56 <= v55 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	F_enlargeStringInfo(m, v47, v55)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L6
	} else {
		goto L15
	}
L13:
	;
	v61 = v55
	goto L14
L14:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
	v63 = v61 + v62
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
	v69 = v64
	v70 = v62
	v73 = v2
	goto L16
L15:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	v61 = v60
	goto L14
L16:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	if v83 <= v73 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+268)) = v270 - v452
	v461 = v449
	goto L1
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+280)) = v83 << (uint(int32(1)) % 32)
	v90 = F_repalloc(m, v82, v83<<(uint(int32(3))%32))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L6
	} else {
		goto L21
	}
L19:
	;
	v93 = v82
	goto L20
L20:
	;
	v95 = v73 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v93+v95))) = v69
	v98 = int32(0)
	if base.Ui32(v63) <= base.Ui32(v70) {
		v268 = v70
		v269 = v70
		v270 = v69
		v274 = v98
		v278 = v98
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+284)) = v90
	v93 = v90
	goto L20
L22:
	;
	v281 = v268 - v70
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v281 != v282 {
		goto L74
	} else {
		goto L75
	}
L23:
	;
	v103 = v70
	v105 = v69
	v109 = v98
	goto L24
L24:
	;
	v117 = v103 + int32(1)
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	v119 = base.B2i32(v118 == v45&int32(255))
	if v118 == v45&int32(255) {
		v268 = v103
		v269 = v117
		v270 = v105
		v274 = v109
		v278 = v119
		goto L22
	} else {
		goto L26
	}
L25:
	;
	v268 = v258
	v269 = v258
	v270 = v264
	v274 = v260
	v278 = v119
	goto L22
L26:
	;
	if v118 != int32(92) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v105))) = uint8(v257)
	v264 = v105 + int32(1)
	if base.Ui32(v258) < base.Ui32(v63) {
		v103 = v258
		v105 = v264
		v109 = v260
		goto L24
	} else {
		goto L72
	}
L28:
	;
	v257 = v118
	v258 = v117
	v260 = v109
	goto L27
L29:
	;
	goto L30
L30:
	;
	if base.Ui32(v63) <= base.Ui32(v117) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v268 = v103
	v269 = v117
	v270 = v105
	v274 = v109
	v278 = v119
	goto L22
L32:
	;
	goto L33
L33:
	;
	v124 = v103 + int32(2)
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+1)))
	v127 = v125 - int32(48)
	switch v127 {
	case 0, 1, 2, 3, 4, 5, 6, 7:
		goto L41
	default:
		v257 = v125
		v258 = v124
		v260 = v109
		goto L27
	case 50:
		goto L39
	case 54:
		goto L38
	case 62:
		goto L37
	case 66:
		goto L36
	case 68:
		goto L35
	case 70:
		goto L34
	case 72:
		goto L40
	}
L34:
	;
	v257 = int32(11)
	v258 = v124
	v260 = v109
	goto L27
L35:
	;
	v257 = int32(9)
	v258 = v124
	v260 = v109
	goto L27
L36:
	;
	v257 = int32(13)
	v258 = v124
	v260 = v109
	goto L27
L37:
	;
	v257 = int32(10)
	v258 = v124
	v260 = v109
	goto L27
L38:
	;
	v257 = int32(12)
	v258 = v124
	v260 = v109
	goto L27
L39:
	;
	v257 = int32(8)
	v258 = v124
	v260 = v109
	goto L27
L40:
	;
	v167 = int32(120)
	if base.Ui32(v63) <= base.Ui32(v124) {
		v257 = v167
		v258 = v124
		v260 = v109
		goto L27
	} else {
		goto L51
	}
L41:
	;
	if base.Ui32(v63) <= base.Ui32(v124) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v257 = v154
	v258 = v155
	v260 = base.B2i32(v154&int32(255) == int32(0)) | int32(base.Ui32(v154&int32(128))>>(uint(int32(7))%32)) | v109
	goto L27
L43:
	;
	v154 = v127
	v155 = v124
	goto L42
L44:
	;
	goto L45
L45:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	if v129&int32(248) != int32(48) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v154 = v127
	v155 = v124
	goto L42
L47:
	;
	goto L48
L48:
	;
	v134 = int32(3)
	v138 = v129 + v127<<(uint(v134)%32) - int32(48)
	v140 = v103 + v134
	if base.Ui32(v63) <= base.Ui32(v140) {
		v154 = v138
		v155 = v140
		goto L42
	} else {
		goto L49
	}
L49:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	if v142&int32(248) != int32(48) {
		v154 = v138
		v155 = v140
		goto L42
	} else {
		goto L50
	}
L50:
	;
	v154 = v142 + v138<<(uint(int32(3))%32) - int32(48)
	v155 = v103 + int32(4)
	goto L42
L51:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	goto L52
L52:
	;
	if base.B2i32(base.Ui32(v169-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v169|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		v257 = v167
		v258 = v124
		v260 = v109
		goto L27
	} else {
		goto L53
	}
L53:
	;
	v184 = v169 - int32(48)
	if base.Ui32(int32(9)) < base.Ui32(v184&int32(255)) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	if base.Ui32(v169-int32(65)) < base.Ui32(int32(26)) {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	v198 = v184
	goto L56
L56:
	;
	v200 = v103 + int32(3)
	if base.Ui32(v63) <= base.Ui32(v200) {
		v237 = v198
		v238 = v200
		goto L61
	} else {
		goto L62
	}
L57:
	;
	v198 = v195 - int32(87)
	goto L56
L58:
	;
	v195 = v169 | int32(32)
	goto L60
L59:
	;
	v195 = v169
	goto L60
L60:
	;
	goto L57
L61:
	;
	v257 = v237
	v258 = v238
	v260 = base.B2i32(v237&int32(255) == int32(0)) | int32(base.Ui32(v237&int32(128))>>(uint(int32(7))%32)) | v109
	goto L27
L62:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200))))
	goto L63
L63:
	;
	if base.B2i32(base.Ui32(v202-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v202|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		v237 = v198
		v238 = v200
		goto L61
	} else {
		goto L64
	}
L64:
	;
	v219 = v202 - int32(48)
	if base.Ui32(int32(9)) < base.Ui32(v219&int32(255)) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	if base.Ui32(v202-int32(65)) < base.Ui32(int32(26)) {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	v233 = v219
	goto L67
L67:
	;
	v237 = v233 + v198<<(uint(int32(4))%32)
	v238 = v103 + int32(4)
	goto L61
L68:
	;
	v233 = v230 - int32(87)
	goto L67
L69:
	;
	v230 = v202 | int32(32)
	goto L71
L70:
	;
	v230 = v202
	goto L71
L71:
	;
	goto L68
L72:
	;
	goto L25
L73:
	;
	v446 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v270))) = uint8(v446)
	v448 = int32(1)
	v449 = v73 + v448
	if v278 != 0 {
		v69 = v270 + v448
		v70 = v269
		v73 = v449
		goto L16
	} else {
		goto L123
	}
L74:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v333 != 0 {
		goto L92
	} else {
		goto L93
	}
L75:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v281 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	if v328 != 0 {
		goto L74
	} else {
		goto L90
	}
L77:
	;
	v328 = int32(0)
	goto L76
L78:
	;
	goto L79
L79:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v290 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v291 = v70
	v292 = v284
	v293 = v281
	v294 = v290
	goto L84
L81:
	;
	v316 = v284
	v320 = int32(0)
	goto L82
L82:
	;
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v316))))
	v328 = v320 - v321
	goto L76
L83:
	;
	v316 = v311
	v320 = v313
	goto L82
L84:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292))))
	if v294 != v296 {
		v311 = v292
		v313 = v294
		goto L83
	} else {
		goto L86
	}
L85:
	;
	v311 = v305
	v313 = int32(0)
	goto L83
L86:
	;
	if v296 == int32(0) {
		v311 = v292
		v313 = v294
		goto L83
	} else {
		goto L87
	}
L87:
	;
	v301 = v293 - int32(1)
	if v301 == int32(0) {
		v311 = v292
		v313 = v294
		goto L83
	} else {
		goto L88
	}
L88:
	;
	v304 = int32(1)
	v305 = v292 + v304
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291)+1)))
	if v306 != 0 {
		v291 = v291 + v304
		v292 = v305
		v293 = v301
		v294 = v306
		goto L84
	} else {
		goto L89
	}
L89:
	;
	goto L85
L90:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	*(*int32)(unsafe.Add(mBase, uint32(v329+v95))) = int32(0)
	goto L73
L91:
	;
	if v274&int32(1) == int32(0) {
		goto L73
	} else {
		goto L121
	}
L92:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v333)+4))
	v336 = v334
	goto L94
L93:
	;
	v336 = int32(0)
	goto L94
L94:
	;
	if v336 <= v73 {
		goto L91
	} else {
		goto L95
	}
L95:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v338 == int32(0) {
		goto L91
	} else {
		goto L96
	}
L96:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v281 != v341 {
		goto L91
	} else {
		goto L97
	}
L97:
	;
	if v281 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	if v386 != 0 {
		goto L91
	} else {
		goto L112
	}
L99:
	;
	v386 = int32(0)
	goto L98
L100:
	;
	goto L101
L101:
	;
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v348 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v349 = v70
	v350 = v338
	v351 = v281
	v352 = v348
	goto L106
L103:
	;
	v374 = v338
	v378 = int32(0)
	goto L104
L104:
	;
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374))))
	v386 = v378 - v379
	goto L98
L105:
	;
	v374 = v369
	v378 = v371
	goto L104
L106:
	;
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350))))
	if v352 != v354 {
		v369 = v350
		v371 = v352
		goto L105
	} else {
		goto L108
	}
L107:
	;
	v369 = v363
	v371 = int32(0)
	goto L105
L108:
	;
	if v354 == int32(0) {
		v369 = v350
		v371 = v352
		goto L105
	} else {
		goto L109
	}
L109:
	;
	v359 = v351 - int32(1)
	if v359 == int32(0) {
		v369 = v350
		v371 = v352
		goto L105
	} else {
		goto L110
	}
L110:
	;
	v362 = int32(1)
	v363 = v350 + v362
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349)+1)))
	if v364 != 0 {
		v349 = v349 + v362
		v350 = v363
		v351 = v359
		v352 = v364
		goto L106
	} else {
		goto L111
	}
L111:
	;
	goto L107
L112:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v333)+12))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v388+v95)))
	v392 = v390 - int32(1)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v387+v392<<(uint(int32(2))%32))))
	if v396 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
	v399 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v397+v392))) = uint8(v399)
	goto L73
L114:
	;
	goto L115
L115:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v401)+52))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v402)))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L6
	} else {
		goto L116
	}
L116:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L6
	} else {
		goto L117
	}
L117:
	;
	F_errmsg(m, int32(_a_F_CopyReadAttributesText_3), int32(0))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L6
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v402 + v403<<(uint(int32(4))%32) + v392*int32(100) + int32(24)
	F_errdetail(m, int32(_a_F_CopyReadAttributesText_4), v18)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L6
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(_a_F_CopyReadAttributesText_1), int32(1775), int32(_a_F_CopyReadAttributesText_2))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
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
	v437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v437+v95)))
	F_pg_verifymbstr(m, v439, v270-v439)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L6
	} else {
		goto L122
	}
L122:
	;
	goto L73
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
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
	var v102 int32
	_ = v102
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	v5 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	if l0 == v5 {
		v118 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v15 + int32(16)
	return v118
L2:
	;
	v24 = l0
	v26 = v5
	goto L3
L3:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	if v31 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v118 = int32(0)
	goto L1
L5:
	;
	if l2 != 0 {
		v118 = v102
		goto L1
	} else {
		goto L30
	}
L6:
	;
	v102 = int32(0)
	goto L5
L7:
	;
	goto L8
L8:
	;
	v35 = int32(0)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v37 <= v35 {
		v102 = v35
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v44 = v35
	v49 = v35
	goto L10
L10:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v52+v49<<(uint(int32(2))%32))))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+21)))
	if v57 != int32(1) {
		v74 = v44
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
	v77 = v49 + int32(1)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v77 < v78 {
		v44 = v74
		v49 = v77
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
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+32)))
	if v63 != int32(1) {
		v74 = v44
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v66 = F_scanNSItemForColumn(m, l0, v56, v26, l1, l3)
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
		v74 = v44
		goto L13
	} else {
		goto L21
	}
L21:
	;
	if v44 != 0 {
		goto L12
	} else {
		goto L22
	}
L22:
	;
	F_check_lateral_ref_ok(m, v24, v56, l3)
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
	v102 = v74
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
	F_parser_errposition(m, v24, l3)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L19
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_colNameToVar_1), int32(932), int32(_a_F_colNameToVar_2))
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
	if v102 != 0 {
		v118 = v102
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v112 != 0 {
		v24 = v112
		v26 = v26 + int32(1)
		goto L3
	} else {
		goto L32
	}
L32:
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
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
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
				v29 = int32(0)
				v30 = F_palloc0(m, v11)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					if base.Ui32(l2) < base.Ui32(v11) {
						v33 = l2
					} else {
						v33 = v11
					}
					if v33 != 0 {
						v34 = F__emscripten_memcpy_bulkmem(m, v30, l1, v33)
						mBase = m.M
						v35 = v34
					} else {
						v35 = v30
					}
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
					v37 = m.T0[v36].(func(*base.Module, int32, int32, int32, int32) int32)(m, v9, v35, v33, v29)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						if v29 != 0 {
							F_pfree(m, v29)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								F_pfree(m, v35)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									return v37
								}
							}
						} else {
							F_pfree(m, v35)
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
							v29 = v20
						} else {
							v25 = l4
							if v25 != 0 {
								v26 = F__emscripten_memcpy_bulkmem(m, v20, l3, v25)
								mBase = m.M
							} else {
							}
							v29 = v20
						}
					} else {
						v25 = v16
						if v25 != 0 {
							v26 = F__emscripten_memcpy_bulkmem(m, v20, l3, v25)
							mBase = m.M
						} else {
						}
						v29 = v20
					}
					v30 = F_palloc0(m, v11)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						if base.Ui32(l2) < base.Ui32(v11) {
							v33 = l2
						} else {
							v33 = v11
						}
						if v33 != 0 {
							v34 = F__emscripten_memcpy_bulkmem(m, v30, l1, v33)
							mBase = m.M
							v35 = v34
						} else {
							v35 = v30
						}
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
						v37 = m.T0[v36].(func(*base.Module, int32, int32, int32, int32) int32)(m, v9, v35, v33, v29)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							if v29 != 0 {
								F_pfree(m, v29)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int32(0)
								} else {
									F_pfree(m, v35)
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
										return int32(0)
									} else {
										return v37
									}
								}
							} else {
								F_pfree(m, v35)
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
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
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
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
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
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
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
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	v3 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v8 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v35 = F_palloc_extended(m, v30, int32(2))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	v30 = v3
	v31 = v3
	goto L1
L3:
	;
	goto L4
L4:
	;
	v13 = v8
	v14 = v3
	v15 = v3
	goto L5
L5:
	;
	v18 = int32(1)
	v19 = v14 + v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v23 = v15 + v20 + v18
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	if v24 != 0 {
		v13 = v24
		v14 = v19
		v15 = v23
		goto L5
	} else {
		goto L7
	}
L6:
	;
	v30 = v19
	v31 = v23 << (uint(int32(3)) % 32)
	goto L1
L7:
	;
	goto L6
L8:
	;
	return
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v35
	v38 = int32(2)
	v41 = F_palloc_extended(m, v30<<(uint(v38)%32), v38)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v41
	v45 = F_palloc_extended(m, v31, int32(2))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v45
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v48 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v30
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v70
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v73
	v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+56)))
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+20)) = uint16(v75)
	v77 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+58)))
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+22)) = uint16(v77)
	v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+60)))
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+24)) = uint16(v79)
	v81 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+62)))
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)) = uint16(v81)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
	if v86 != 0 {
		goto L32
	} else {
		goto L33
	}
L13:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v45 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v54 != 0 {
		goto L21
	} else {
		goto L22
	}
L16:
	;
	v51 = v49
	goto L18
L17:
	;
	v51 = int32(0)
	goto L18
L18:
	;
	if v51 != 0 {
		goto L12
	} else {
		goto L19
	}
L19:
	;
	F_pfree(m, v48)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	goto L15
L21:
	;
	F_pfree(m, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L8
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v57 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L23
L25:
	;
	F_pfree(m, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L8
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+24)) = int32(101)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	if v64 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L27
L29:
	;
	v66 = v64
	goto L31
L30:
	;
	v66 = int32(12)
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v63)+12)) = v66
	return
L32:
	;
	v90 = int32(0)
	goto L34
L33:
	;
	v87 = int32(*(*int16)(unsafe.Add(mBase, uint32(v84)+12)))
	v90 = v87 + int32(1)
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v90
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v92
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v94
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v96
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v98 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v104 = v45
	v105 = v98
	goto L38
L36:
	;
	goto L37
L37:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+20))
	if v188 != 0 {
		goto L61
	} else {
		goto L62
	}
L38:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v109 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v106+v107))) = uint8(v109)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	*(*int32)(unsafe.Add(mBase, uint32(v111+v112<<(uint(int32(2))%32)))) = v104
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v105)+20))
	if v117 != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L37
L40:
	;
	v120 = v117
	v121 = v104
	goto L43
L41:
	;
	v160 = v104
	goto L42
L42:
	;
	v166 = (v160 - v104) >> (uint(int32(3)) % 32)
	if base.Ui32(int32(2)) <= base.Ui32(v166) {
		goto L56
	} else {
		goto L57
	}
L43:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	if v125 != int32(76) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	v160 = v155
	goto L42
L45:
	;
	v155 = v121 + int32(8)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v120)+16))
	if v156 != 0 {
		v120 = v156
		v121 = v155
		goto L43
	} else {
		goto L55
	}
L46:
	;
	if v125 == int32(112) {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	goto L48
L48:
	;
	v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v120)+4)))
	v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v145 = v143 + v144
	*(*uint16)(unsafe.Add(mBase, uint32(v121))) = uint16(v145)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	*(*int32)(unsafe.Add(mBase, uint32(v121)+4)) = v148
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v150 | int32(1)
	goto L45
L49:
	;
	v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v120)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v121))) = uint16(v130)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v121)+4)) = v133
	goto L45
L50:
	;
	goto L51
L51:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v135)+24)) = int32(101)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+12))
	if v139 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v141 = v139
	goto L54
L53:
	;
	v141 = int32(15)
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v138)+12)) = v141
	return
L55:
	;
	goto L44
L56:
	;
	F_pg_qsort(m, v104, v166, int32(8), int32(971))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L8
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160)+4)) = int32(0)
	v175 = int32(_a_F_compact_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v160))) = uint16(v175)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v105)+28))
	if v179 != 0 {
		v104 = v160 + int32(8)
		v105 = v179
		goto L38
	} else {
		goto L60
	}
L59:
	;
	goto L58
L60:
	;
	goto L39
L61:
	;
	v191 = v188
	goto L64
L62:
	;
	v207 = v187
	goto L63
L63:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
	v214 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v211+v212))) = uint8(v214)
	return
L64:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v191)+12))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
	v200 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v196+v198))) = uint8(v200)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v191)+16))
	if v202 != 0 {
		v191 = v202
		goto L64
	} else {
		goto L66
	}
L65:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v207 = v203
	goto L63
L66:
	;
	goto L65
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
	var v10 int32
	_ = v10
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
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v14 = m.T0[v13].(func(*base.Module, int32, int32, int32) int32)(m, v10, v11, v12)
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
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
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
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
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
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	v7 = l1
	v8 = l2
	v9 = l3
	goto L2
L1:
	;
	if l0 != v7 {
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
	v22 = v8 << (uint(int32(3)) % 32)
	if l0 == v7 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	goto L8
L8:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v167 + v9
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v170 - v9
	return v8
L9:
	;
	goto L8
L10:
	;
	goto L9
L11:
	;
	v26 = l0 + v22
	if base.Ui32(v7-v26) <= base.Ui32(int32(0)-v22<<(uint(int32(1))%32)) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v33 = F___memcpy(m, l0, v7, v22)
	mBase = m.M
	goto L9
L13:
	;
	goto L14
L14:
	;
	v36 = (l0 ^ v7) & int32(3)
	if base.Ui32(l0) < base.Ui32(v7) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	if v138 == int32(0) {
		goto L10
	} else {
		goto L51
	}
L16:
	;
	if base.Ui32(v116) <= base.Ui32(int32(3)) {
		v137 = v115
		v138 = v116
		v139 = v117
		goto L15
	} else {
		goto L47
	}
L17:
	;
	if v36 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	if v36 != 0 {
		v98 = v22
		goto L30
	} else {
		goto L31
	}
L20:
	;
	v137 = v7
	v138 = v22
	v139 = l0
	goto L15
L21:
	;
	goto L22
L22:
	;
	if l0&int32(3) == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v115 = v7
	v116 = v22
	v117 = l0
	goto L16
L24:
	;
	goto L25
L25:
	;
	v43 = v7
	v44 = v22
	v45 = l0
	goto L26
L26:
	;
	if v44 == int32(0) {
		goto L10
	} else {
		goto L28
	}
L27:
	;
	v115 = v52
	v116 = v54
	v117 = v56
	goto L16
L28:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	*(*uint8)(unsafe.Add(mBase, uint32(v45))) = uint8(v49)
	v51 = int32(1)
	v52 = v43 + v51
	v54 = v44 - v51
	v56 = v45 + v51
	if v56&int32(3) != 0 {
		v43 = v52
		v44 = v54
		v45 = v56
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	if v98 == int32(0) {
		goto L10
	} else {
		goto L43
	}
L31:
	;
	if v26&int32(3) != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v63 = v22
	goto L35
L33:
	;
	v78 = v22
	goto L34
L34:
	;
	if base.Ui32(v78) <= base.Ui32(int32(3)) {
		v98 = v78
		goto L30
	} else {
		goto L39
	}
L35:
	;
	if v63 == int32(0) {
		goto L10
	} else {
		goto L37
	}
L36:
	;
	v78 = v69
	goto L34
L37:
	;
	v69 = v63 - int32(1)
	v70 = l0 + v69
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7+v69))))
	*(*uint8)(unsafe.Add(mBase, uint32(v70))) = uint8(v72)
	if v70&int32(3) != 0 {
		v63 = v69
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v85 = v78
	goto L40
L40:
	;
	v89 = v85 - int32(4)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v7+v89)))
	*(*int32)(unsafe.Add(mBase, uint32(l0+v89))) = v92
	if base.Ui32(int32(3)) < base.Ui32(v89) {
		v85 = v89
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v98 = v89
	goto L30
L42:
	;
	goto L41
L43:
	;
	v105 = v98
	goto L44
L44:
	;
	v109 = v105 - int32(1)
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7+v109))))
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v109))) = uint8(v112)
	if v109 != 0 {
		v105 = v109
		goto L44
	} else {
		goto L46
	}
L45:
	;
	goto L10
L46:
	;
	goto L45
L47:
	;
	v122 = v115
	v123 = v116
	v124 = v117
	goto L48
L48:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	*(*int32)(unsafe.Add(mBase, uint32(v124))) = v126
	v128 = int32(4)
	v129 = v122 + v128
	v131 = v124 + v128
	v133 = v123 - v128
	if base.Ui32(int32(3)) < base.Ui32(v133) {
		v122 = v129
		v123 = v133
		v124 = v131
		goto L48
	} else {
		goto L50
	}
L49:
	;
	v137 = v129
	v138 = v133
	v139 = v131
	goto L15
L50:
	;
	goto L49
L51:
	;
	v144 = v137
	v145 = v138
	v146 = v139
	goto L52
L52:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	*(*uint8)(unsafe.Add(mBase, uint32(v146))) = uint8(v148)
	v150 = int32(1)
	v155 = v145 - v150
	if v155 != 0 {
		v144 = v144 + v150
		v145 = v155
		v146 = v146 + v150
		goto L52
	} else {
		goto L54
	}
L53:
	;
	goto L10
L54:
	;
	goto L53
}
func F_compute_trivial_stats(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 float64
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v152 float64
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v170 float64
	_ = v170
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	v5 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+78)))
	if v18 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+76)))
	v24 = int32(_a_F_compute_trivial_stats_0)
	v29 = base.B2i32(v21 < int32(0))
	v30 = base.B2i32(v21&v24 == v24)
	goto L3
L2:
	;
	v29 = v5
	v30 = v5
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
	m.G0 = v15 + int32(16)
	return
L5:
	;
	v37 = float64(0)
	v39 = v5
	v41 = v5
	v42 = v5
	goto L6
L6:
	;
	F_vacuum_delay_point(m, int32(1))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if int32(0) < v155 {
		goto L47
	} else {
		goto L48
	}
L8:
	;
	return
L9:
	;
	v51 = m.T0[l1].(func(*base.Module, int32, int32, int32) int32)(m, l0, v39, v15+int32(15))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+15)))
	if v53 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v158 = v39 + int32(1)
	if v158 != l2 {
		v37 = v152
		v39 = v158
		v41 = v154
		v42 = v155
		goto L6
	} else {
		goto L45
	}
L12:
	;
	v152 = v37
	v154 = v41 + int32(1)
	v155 = v42
	goto L11
L13:
	;
	goto L14
L14:
	;
	v59 = v42 + int32(1)
	if v30 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v60 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	if v29 == int32(0) {
		v152 = v37
		v154 = v41
		v155 = v59
		goto L11
	} else {
		goto L27
	}
L18:
	;
	v152 = base.F64_add(v37, base.F64_convert_i32_u(v86))
	v154 = v41
	v155 = v59
	goto L11
L19:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
	if base.Ui32((v64-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v86 = int32(6)
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v78 = int32(1)
	if v60&v78 != 0 {
		v86 = int32(base.Ui32(v60) >> (uint(v78) % 32))
		goto L18
	} else {
		goto L26
	}
L22:
	;
	v71 = int32(18)
	if v64&int32(255) == v71 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v77 = v71
	goto L25
L24:
	;
	v77 = int32(2)
	goto L25
L25:
	;
	v86 = v77
	goto L18
L26:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v86 = int32(base.Ui32(v82) >> (uint(int32(2)) % 32))
	goto L18
L27:
	;
	if v51&int32(3) == int32(0) {
		v114 = v51
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v152 = base.F64_add(v37, base.F64_convert_i32_u(v147+int32(1)))
	v154 = v41
	v155 = v59
	goto L11
L29:
	;
	v147 = v139 - v51
	goto L28
L30:
	;
	v118 = v114
	goto L39
L31:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v98 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v147 = int32(0)
	goto L28
L33:
	;
	goto L34
L34:
	;
	v103 = v51
	goto L35
L35:
	;
	v107 = v103 + int32(1)
	if v107&int32(3) == int32(0) {
		v114 = v107
		goto L30
	} else {
		goto L37
	}
L36:
	;
	v139 = v107
	goto L29
L37:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if v112 != 0 {
		v103 = v107
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	v127 = int32(-2139062144)
	if (int32(16843008)-v124|v124)&v127 == v127 {
		v118 = v118 + int32(4)
		goto L39
	} else {
		goto L41
	}
L40:
	;
	v133 = v118
	goto L42
L41:
	;
	goto L40
L42:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	if v137 != 0 {
		v133 = v133 + int32(1)
		goto L42
	} else {
		goto L44
	}
L43:
	;
	v139 = v133
	goto L29
L44:
	;
	goto L43
L45:
	;
	goto L7
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v188
	goto L4
L47:
	;
	v162 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v162)
	*(*float32)(unsafe.Add(mBase, uint32(l0)+40)) = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v154), base.F64_convert_i32_s(l2)))
	if v29 != 0 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	if v154 <= int32(0) {
		goto L4
	} else {
		goto L56
	}
L50:
	;
	v170 = base.F64_div(v152, base.F64_convert_i32_u(v155))
	if base.F64_lt(base.F64_abs(v170), float64(2.147483648e+09)) != 0 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v177 = int32(*(*int16)(unsafe.Add(mBase, uint32(v176)+76)))
	v188 = v177
	goto L46
L53:
	;
	v174 = base.I32_trunc_f64_s(v170)
	v188 = v174
	goto L46
L54:
	;
	goto L55
L55:
	;
	v188 = int32(-2147483648)
	goto L46
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(1065353216)
	v182 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v182)
	if v29 != 0 {
		v188 = int32(0)
		goto L46
	} else {
		goto L57
	}
L57:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v186 = int32(*(*int16)(unsafe.Add(mBase, uint32(v185)+76)))
	v188 = v186
	goto L46
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
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v83 int32
	_ = v83
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v248 int32
	_ = v248
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
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
	var v313 int32
	_ = v313
	var v323 int32
	_ = v323
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	v10 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(48)
	m.G0 = v21
	if v10 <= l2 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L7
	} else {
		goto L89
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L7
	} else {
		goto L85
	}
L3:
	;
	if base.Ui32(int32(7)) <= base.Ui32(l2) {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L7
	} else {
		goto L81
	}
L6:
	;
	v27 = F_ArrayGetNItems(m, l2, l3)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	F_ArrayCheckBounds(m, l2, l3, l4)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	if int32(0) < v27 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	m.G0 = v21 + int32(48)
	return v323
L11:
	;
	v56 = v10
	v57 = v10
	v62 = v10
	goto L19
L12:
	;
	goto L11
L13:
	;
	goto L14
L14:
	;
	v40 = F_palloc0(m, int32(16))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+12)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v40)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v40))) = int64(64)
	v323 = v40
	goto L10
L16:
	;
	v293 = v283 + v284
	v294 = F_palloc0(m, v293)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L7
	} else {
		goto L71
	}
L17:
	;
	v283 = (l2<<(uint(int32(3))%32) + int32(23)) & int32(120)
	v284 = v230
	v292 = int32(0)
	goto L16
L18:
	;
	v259 = base.I32_div_s(v27+int32(7), int32(8))
	v266 = (v259 + l2<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	v283 = v266
	v284 = v248
	v292 = v266
	goto L16
L19:
	;
	if l1 == int32(0) {
		v107 = v56
		v113 = v62
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if v113 == int32(0) {
		goto L17
	} else {
		goto L70
	}
L21:
	;
	if base.B2i32(l6 == int32(-1)) == int32(0) {
		goto L32
	} else {
		goto L33
	}
L22:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v56))))
	if v68 != int32(1) {
		v107 = v56
		v113 = v62
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v72 = v56 + int32(1)
	if v72 == v27 {
		v248 = v57
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v83 = v72
	goto L25
L25:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v83))))
	if v93 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v107 = v83
	v113 = int32(1)
	goto L21
L27:
	;
	v95 = v83 + int32(1)
	if v27 != v95 {
		v83 = v95
		goto L25
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	goto L26
L30:
	;
	v248 = v57
	goto L18
L31:
	;
	v217 = v57 + v215
	switch l8 - int32(99) {
	case 0:
		v230 = v217
		goto L64
	case 1:
		goto L66
	default:
		goto L65
	case 6:
		goto L67
	}
L32:
	;
	if int32(0) < l6 {
		v215 = l6
		goto L31
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v185 = l0 + v107<<(uint(int32(2))%32)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v185)))
	v187 = F_pg_detoast_datum(m, v186)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L7
	} else {
		goto L53
	}
L35:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0+v107<<(uint(int32(2))%32))))
	if v123&int32(3) == int32(0) {
		v147 = v123
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v215 = v180 + int32(1)
	goto L31
L37:
	;
	v180 = v172 - v123
	goto L36
L38:
	;
	v151 = v147
	goto L47
L39:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	if v131 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v180 = int32(0)
	goto L36
L41:
	;
	goto L42
L42:
	;
	v136 = v123
	goto L43
L43:
	;
	v140 = v136 + int32(1)
	if v140&int32(3) == int32(0) {
		v147 = v140
		goto L38
	} else {
		goto L45
	}
L44:
	;
	v172 = v140
	goto L37
L45:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	if v145 != 0 {
		v136 = v140
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
	v160 = int32(-2139062144)
	if (int32(16843008)-v157|v157)&v160 == v160 {
		v151 = v151 + int32(4)
		goto L47
	} else {
		goto L49
	}
L48:
	;
	v166 = v151
	goto L50
L49:
	;
	goto L48
L50:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166))))
	if v170 != 0 {
		v166 = v166 + int32(1)
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v172 = v166
	goto L37
L52:
	;
	goto L51
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v185))) = v187
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	if v190 == int32(1) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+1)))
	if base.Ui32((v194-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v215 = int32(6)
		goto L31
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	if v190&int32(1) != 0 {
		goto L61
	} else {
		goto L62
	}
L57:
	;
	v201 = int32(18)
	if v194&int32(255) == v201 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v207 = v201
	goto L60
L59:
	;
	v207 = int32(2)
	goto L60
L60:
	;
	v215 = v207
	goto L31
L61:
	;
	v215 = int32(base.Ui32(v190) >> (uint(int32(1)) % 32))
	goto L31
L62:
	;
	goto L63
L63:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	v215 = int32(base.Ui32(v212) >> (uint(int32(2)) % 32))
	goto L31
L64:
	;
	if base.Ui32(int32(1073741824)) <= base.Ui32(v230) {
		goto L1
	} else {
		goto L68
	}
L65:
	;
	v230 = (v217 + int32(1)) & int32(-2)
	goto L64
L66:
	;
	v230 = (v217 + int32(7)) & int32(-8)
	goto L64
L67:
	;
	v230 = (v217 + int32(3)) & int32(-4)
	goto L64
L68:
	;
	v234 = v107 + int32(1)
	if v234 != v27 {
		v56 = v234
		v57 = v230
		v62 = v113
		goto L19
	} else {
		goto L69
	}
L69:
	;
	goto L20
L70:
	;
	v248 = v230
	goto L18
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v294)+12)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v294)+8)) = v292
	*(*int32)(unsafe.Add(mBase, uint32(v294)+4)) = l2
	v299 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v294))) = v293 << (uint(v299) % 32)
	v303 = v294 + int32(16)
	v305 = l2 << (uint(v299) % 32)
	if v305 != 0 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	if v305 != 0 {
		goto L77
	} else {
		goto L78
	}
L73:
	;
	v306 = F__emscripten_memcpy_bulkmem(m, v303, l3, v305)
	mBase = m.M
	v307 = v306
	goto L75
L74:
	;
	v307 = v303
	goto L75
L75:
	;
	goto L72
L76:
	;
	F_CopyArrayEls(m, v294, l0, l1, v27, l6, l7, l8, int32(0))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L7
	} else {
		goto L80
	}
L77:
	;
	v309 = F__emscripten_memcpy_bulkmem(m, v307+v305, l4, v305)
	mBase = m.M
	goto L79
L78:
	;
	goto L79
L79:
	;
	goto L76
L80:
	;
	v323 = v294
	goto L10
L81:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L7
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = l2
	F_errmsg(m, int32(_a_F_construct_md_array_0), v21)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L7
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_construct_md_array_1), int32(3511), int32(_a_F_construct_md_array_2))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L7
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L85:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L7
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = l2
	F_errmsg(m, int32(_a_F_construct_md_array_3), v21+int32(16))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L7
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_construct_md_array_1), int32(3516), int32(_a_F_construct_md_array_2))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L7
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L89:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L7
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = int32(1073741823)
	F_errmsg(m, int32(_a_F_construct_md_array_4), v21+int32(32))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L7
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(_a_F_construct_md_array_1), int32(3546), int32(_a_F_construct_md_array_2))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L7
	} else {
		goto L92
	}
L92:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_convert_tuples_by_position(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
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
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v318 int32
	_ = v318
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v391 int32
	_ = v391
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	v4 = int32(0)
	v20 = m.G0
	v22 = v20 + int32(-64)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v26 = F_palloc0(m, int32(8))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v391 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L2:
	;
	return int32(0)
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v24
	v31 = int32(1)
	v34 = F_palloc0(m, v24<<(uint(v31)%32))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v34
	v38 = l1 + int32(20)
	if v24 <= int32(0) {
		v208 = v4
		v213 = v31
		v214 = v4
		v218 = v34
		v219 = v4
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v224 <= v208 {
		v295 = v213
		v296 = v214
		goto L34
	} else {
		goto L35
	}
L6:
	;
	v45 = v4
	v48 = v4
	v50 = v31
	v51 = v4
	v55 = v34
	v56 = v4
	goto L7
L7:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v67 = v38 + v61<<(uint(int32(4))%32) + v48*int32(100)
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+91)))
	if v68 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L2
	} else {
		goto L27
	}
L9:
	;
	goto L8
L10:
	;
	v72 = v56 + int32(1)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v73 <= v45 {
		v124 = v45
		v130 = v51
		v134 = v55
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v149 = v45
	v154 = v50
	v155 = v51
	v159 = v55
	v160 = v56
	goto L12
L12:
	;
	v166 = v48 + int32(1)
	if v24 != v166 {
		v45 = v149
		v48 = v166
		v50 = v154
		v51 = v155
		v55 = v159
		v56 = v160
		goto L7
	} else {
		goto L26
	}
L13:
	;
	v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v134+v48<<(uint(int32(1))%32)))))
	v149 = v124
	v154 = base.B2i32(v143 != int32(0)) & v50
	v155 = v130
	v159 = v134
	v160 = v72
	goto L12
L14:
	;
	v80 = v45
	goto L15
L15:
	;
	v98 = l0 + int32(20) + v73<<(uint(int32(4))%32) + v80*int32(100)
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+91)))
	if v99 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v124 = v73
	v130 = v51
	v134 = v55
	goto L13
L17:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v67)+68))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v98)+68))
	if v102 != v103 {
		goto L9
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v120 = v80 + int32(1)
	if v120 != v73 {
		v80 = v120
		goto L15
	} else {
		goto L25
	}
L20:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v67)+76))
	if int32(0) <= v105 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v98)+76))
	if v105 != v108 {
		goto L9
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v110 = int32(1)
	v116 = v80 + v110
	*(*uint16)(unsafe.Add(mBase, uint32(v55+v48<<(uint(v110)%32)))) = uint16(v116)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v124 = v116
	v130 = v51 + v110
	v134 = v118
	goto L13
L24:
	;
	goto L23
L25:
	;
	goto L16
L26:
	;
	v208 = v149
	v213 = v154
	v214 = v155
	v218 = v159
	v219 = v160
	goto L5
L27:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L2
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = l2
	F_errmsg_internal(m, int32(_a_F_convert_tuples_by_position_0), v20+int32(-16))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L2
	} else {
		goto L29
	}
L29:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v98)+68))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v98)+76))
	v184 = F_format_type_with_typemod(m, v182, v183)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L2
	} else {
		goto L30
	}
L30:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v67)+68))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v67)+76))
	v188 = F_format_type_with_typemod(m, v186, v187)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+44)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v67 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = v188
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v184
	F_errdetail(m, int32(_a_F_convert_tuples_by_position_1), v20+int32(-32))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_convert_tuples_by_position_2), int32(124), int32(_a_F_convert_tuples_by_position_3))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
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
	if v295&int32(1) != 0 {
		goto L43
	} else {
		goto L44
	}
L35:
	;
	v226 = int32(1)
	v227 = v208 + v226
	v229 = l0 + int32(29)
	if (v224-v208)&v226 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229+v208<<(uint(int32(4))%32)))))
	v243 = v227
	v244 = v236 & v213
	v245 = v214 + (v236^int32(-1))&int32(1)
	goto L38
L37:
	;
	v243 = v208
	v244 = v213
	v245 = v214
	goto L38
L38:
	;
	if v227 == v224 {
		v295 = v244
		v296 = v245
		goto L34
	} else {
		goto L39
	}
L39:
	;
	v251 = v243
	v256 = v244
	v257 = v245
	goto L40
L40:
	;
	v268 = v251 << (uint(int32(4)) % 32)
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(45)+v268))))
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268+v229))))
	v274 = v270 & v272 & v256
	v275 = int32(-1)
	v277 = int32(1)
	v284 = v257 + (v272^v275)&v277 + (v270^v275)&v277
	v286 = v251 + int32(2)
	if v286 != v224 {
		v251 = v286
		v256 = v274
		v257 = v284
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v295 = v274
	v296 = v284
	goto L34
L42:
	;
	goto L41
L43:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v224 != v308 {
		v391 = v26
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
	v404 = m.ExcPending
	if v404 != 0 {
		goto L2
	} else {
		goto L64
	}
L46:
	;
	m.G0 = v22 - int32(-64)
	goto L1
L47:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if int32(0) < v310 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v318 = int32(0)
	goto L51
L49:
	;
	goto L50
L50:
	;
	F_pfree(m, v218)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L2
	} else {
		goto L62
	}
L51:
	;
	v335 = v318 << (uint(int32(4)) % 32)
	v336 = l0 + int32(20) + v335
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+8)))
	if v337 != 0 {
		v391 = v26
		goto L46
	} else {
		goto L53
	}
L52:
	;
	goto L50
L53:
	;
	v338 = int32(1)
	v339 = v318 + v338
	v343 = int32(*(*int16)(unsafe.Add(mBase, uint32(v218+v318<<(uint(v338)%32)))))
	if v339 != v343 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	if v343 != 0 {
		v391 = v26
		goto L46
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	if v339 != v310 {
		v318 = v339
		goto L51
	} else {
		goto L61
	}
L57:
	;
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+9)))
	if v345 != int32(1) {
		v391 = v26
		goto L46
	} else {
		goto L58
	}
L58:
	;
	v348 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v336)+4)))
	v349 = v335 + v38
	v350 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v349)+4)))
	if v348 != v350 {
		v391 = v26
		goto L46
	} else {
		goto L59
	}
L59:
	;
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+12)))
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349)+12)))
	if v352 != v353 {
		v391 = v26
		goto L46
	} else {
		goto L60
	}
L60:
	;
	goto L56
L61:
	;
	goto L52
L62:
	;
	F_pfree(m, v26)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L2
	} else {
		goto L63
	}
L63:
	;
	v391 = int32(0)
	goto L46
L64:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L2
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = l2
	F_errmsg_internal(m, int32(_a_F_convert_tuples_by_position_0), v20+int32(-48))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L2
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v296
	F_errdetail(m, int32(_a_F_convert_tuples_by_position_4), v22)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L2
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_convert_tuples_by_position_2), int32(149), int32(_a_F_convert_tuples_by_position_3))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	return int32(0)
L70:
	;
	goto L71
L71:
	;
	v429 = F_palloc(m, int32(28))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L2
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v429)+8)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v429)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v429))) = l0
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v436 = v434 + int32(1)
	v439 = F_palloc(m, v436<<(uint(int32(2))%32))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L2
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v429)+20)) = v439
	v442 = F_palloc(m, v436)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L2
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v429)+24)) = v442
	v445 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v447 = v445 + int32(1)
	v450 = F_palloc(m, v447<<(uint(int32(2))%32))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L2
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v429)+12)) = v450
	v453 = F_palloc(m, v447)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L2
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v429)+16)) = v453
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v429)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v456))) = int32(0)
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v429)+16))
	v460 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v459))) = uint8(v460)
	return v429
}
func F_copy_addr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	v2 = l1
	if v2&int32(255) != int32(10) {
		if v2 != int32(2) {
		} else {
			v15 = int32(4)
			v33 = v15
			v34 = l2 + v15
			if base.Ui32(l4) < base.Ui32(v33) {
			} else {
				*(*uint16)(unsafe.Add(mBase, uint32(l2))) = uint16(v2)
				if v33 != 0 {
					v37 = F__emscripten_memcpy_bulkmem(m, v34, l3, v33)
					mBase = m.M
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = l2
			}
		}
	} else {
		v19 = l2 + int32(8)
		v20 = int32(16)
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
		switch v21 - int32(254) {
		case 0:
			v24 = int32(*(*int8)(unsafe.Add(mBase, uint32(l3)+1)))
			if v24 < int32(-64) {
				*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = l5
				v33 = v20
				v34 = v19
			} else {
				v33 = v20
				v34 = v19
			}
		case 1:
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+1)))
			if v27&int32(15) != int32(2) {
				v33 = v20
				v34 = v19
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = l5
				v33 = v20
				v34 = v19
			}
		default:
			v33 = v20
			v34 = v19
		}
		if base.Ui32(l4) < base.Ui32(v33) {
		} else {
			*(*uint16)(unsafe.Add(mBase, uint32(l2))) = uint16(v2)
			if v33 != 0 {
				v37 = F__emscripten_memcpy_bulkmem(m, v34, l3, v33)
				mBase = m.M
			} else {
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = l2
		}
	}
	return
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
func F_cost_material(m *base.Module, l0 int32, l1 int32, l2 float64, l3 float64, l4 float64, l5 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 float64
	_ = v13
	var v17 float64
	_ = v17
	var v25 float64
	_ = v25
	var v31 float64
	_ = v31
	var v37 float64
	_ = v37
	var v39 int32
	_ = v39
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_cost_material[0]))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = l4
	v13 = *(*float64)(unsafe.Add(mBase, _c_F_cost_material[1]))
	v17 = base.F64_add(base.F64_mul(base.F64_add(v13, v13), l4), base.F64_sub(l3, l2))
	v25 = base.F64_mul(l4, base.F64_convert_i32_u((l5+int32(7))&int32(-8)+int32(24)))
	if base.F64_gt(v25, base.F64_convert_i32_u(v10<<(uint(int32(10))%32))) != 0 {
		v31 = *(*float64)(unsafe.Add(mBase, _c_F_cost_material[2]))
		v37 = base.F64_add(base.F64_mul(v31, base.F64_ceil(base.F64_mul(v25, float64(0.0001220703125)))), v17)
	} else {
		v37 = v17
	}
	v39 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_cost_material[3])))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = base.F64_add(l2, v37)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l1 + (v39 ^ int32(1))
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
	var v112 float64
	_ = v112
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
	var v135 float64
	_ = v135
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
	var v149 float64
	_ = v149
	var v154 float64
	_ = v154
	var v156 float64
	_ = v156
	var v160 float64
	_ = v160
	v9 = float64(0)
	v17 = m.G0
	v19 = v17 - int32(48)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
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
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
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
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
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
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l2)+72))
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
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)+116))
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
	v146 = *(*float64)(unsafe.Add(mBase, uint32(l2)+120))
	v148 = *(*float64)(unsafe.Add(mBase, _c_F_cost_samplescan[0]))
	v149 = *(*float64)(unsafe.Add(mBase, uint32(v144)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
	v154 = base.F64_add(v149, base.F64_add(v143, float64(0)))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = v154
	v156 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
	if v56 != 0 {
		goto L22
	} else {
		goto L23
	}
L12:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v60 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+40)) = v60
	*(*int64)(unsafe.Add(mBase, uint32(v19)+32)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = l1
	if v59 == int32(0) {
		v112 = v9
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
	v125 = *(*float64)(unsafe.Add(mBase, uint32(l2)+200))
	v126 = *(*float64)(unsafe.Add(mBase, uint32(l2)+192))
	v135 = v125
	v143 = v126
	goto L11
L15:
	;
	v121 = *(*float64)(unsafe.Add(mBase, uint32(l2)+200))
	v123 = *(*float64)(unsafe.Add(mBase, uint32(l2)+192))
	v135 = base.F64_add(v112, v121)
	v143 = base.F64_add(v120, v123)
	goto L11
L16:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v69 <= int32(0) {
		v112 = v9
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
	v112 = v102
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
	v160 = v58
	goto L24
L23:
	;
	v160 = v57
	goto L24
L24:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = base.F64_add(v154, base.F64_add(base.F64_mul(v145, v156), base.F64_add(base.F64_mul(v146, base.F64_add(v135, v148)), base.F64_add(base.F64_mul(v160, base.F64_convert_i32_u(v55)), float64(0)))))
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
	var v27 float64
	_ = v27
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
	var v43 float64
	_ = v43
	var v47 float64
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 float64
	_ = v52
	var v54 float64
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 float64
	_ = v103
	var v104 float64
	_ = v104
	var v105 float64
	_ = v105
	var v106 int32
	_ = v106
	var v107 float64
	_ = v107
	var v108 float64
	_ = v108
	var v114 int32
	_ = v114
	var v117 float64
	_ = v117
	var v118 float64
	_ = v118
	var v119 float64
	_ = v119
	var v121 float64
	_ = v121
	var v125 float64
	_ = v125
	var v126 float64
	_ = v126
	var v128 float64
	_ = v128
	var v130 float64
	_ = v130
	var v131 float64
	_ = v131
	var v137 int32
	_ = v137
	var v140 float64
	_ = v140
	var v141 float64
	_ = v141
	var v142 float64
	_ = v142
	var v144 float64
	_ = v144
	var v148 float64
	_ = v148
	var v149 int32
	_ = v149
	var v150 float64
	_ = v150
	var v151 float64
	_ = v151
	var v153 float64
	_ = v153
	var v154 float64
	_ = v154
	var v155 float64
	_ = v155
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
	v27 = float64(1e+100)
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
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)+184))
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
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)+184))
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
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v47
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v50
	v52 = *(*float64)(unsafe.Add(mBase, uint32(v49)+48))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = v52
	v54 = *(*float64)(unsafe.Add(mBase, uint32(v49)+56))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = v54
	if v26 != 0 {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	v35 = base.F64_mul(v29, v33)
	if base.F64_gt(v35, float64(1e+100)) != 0 {
		v47 = v27
		goto L7
	} else {
		goto L9
	}
L9:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v35)&int64(9223372036854775807)) {
		v47 = v27
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v43 = float64(1)
	if base.F64_le(v35, v43) != 0 {
		v47 = v43
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v47 = base.F64_nearest(v35)
	goto L7
L12:
	;
	v57 = int32(0)
	goto L14
L13:
	;
	v57 = l4
	goto L14
L14:
	;
	if v57 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	if l3 != 0 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	m.G0 = v19 + int32(32)
	return
L18:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v150 = *(*float64)(unsafe.Add(mBase, uint32(v149)+24))
	v151 = *(*float64)(unsafe.Add(mBase, uint32(v137)+32))
	v153 = *(*float64)(unsafe.Add(mBase, _c_F_cost_subqueryscan[0]))
	v154 = *(*float64)(unsafe.Add(mBase, uint32(v149)+16))
	v155 = base.F64_add(v148, v154)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = base.F64_add(v155, v140)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = base.F64_add(base.F64_add(v155, base.F64_add(base.F64_mul(v150, v141), base.F64_mul(v151, base.F64_add(v142, v153)))), v144)
	goto L17
L19:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v61 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v61
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = l1
	if v60 == int32(0) {
		v114 = v49
		v117 = v52
		v118 = v47
		v119 = v9
		v121 = v54
		v125 = float64(0)
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v130 = *(*float64)(unsafe.Add(mBase, uint32(l2)+200))
	v131 = *(*float64)(unsafe.Add(mBase, uint32(l2)+192))
	v137 = v49
	v140 = v52
	v141 = v47
	v142 = v130
	v144 = v54
	v148 = v131
	goto L18
L22:
	;
	v126 = *(*float64)(unsafe.Add(mBase, uint32(l2)+200))
	v128 = *(*float64)(unsafe.Add(mBase, uint32(l2)+192))
	v137 = v114
	v140 = v117
	v141 = v118
	v142 = base.F64_add(v119, v126)
	v144 = v121
	v148 = base.F64_add(v125, v128)
	goto L18
L23:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v70 <= int32(0) {
		v114 = v49
		v117 = v52
		v118 = v47
		v119 = v9
		v121 = v54
		v125 = float64(0)
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v79 = int32(0)
	goto L25
L25:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v90+v79<<(uint(int32(2))%32))))
	v97 = F_cost_qual_eval_walker(m, v94, v19+int32(8))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L5
	} else {
		goto L27
	}
L26:
	;
	v103 = *(*float64)(unsafe.Add(mBase, uint32(l0)+56))
	v104 = *(*float64)(unsafe.Add(mBase, uint32(l0)+48))
	v105 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v107 = *(*float64)(unsafe.Add(mBase, uint32(v19)+24))
	v108 = *(*float64)(unsafe.Add(mBase, uint32(v19)+16))
	v114 = v106
	v117 = v104
	v118 = v105
	v119 = v107
	v121 = v103
	v125 = v108
	goto L22
L27:
	;
	v100 = v79 + int32(1)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v100 < v101 {
		v79 = v100
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
}
