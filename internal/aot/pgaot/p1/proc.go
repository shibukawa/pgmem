package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CleanupProcSignalState(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(4405940)
	v10 = *(*int32)(unsafe.Add(mBase, _consts[785]))
	*(*int32)(unsafe.Add(mBase, _consts[785])) = int32(0)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+96)) = int32(1)
	v18 = v10 + int32(96)
	if v14 != 0 {
		F_s_lock(m, v18, int32(495788), int32(243), int32(352847))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			v26 = *(*int32)(unsafe.Add(mBase, _consts[156]))
			if v24 != v26 {
				v28 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v18))) = v28
				v32 = F_errstart(m, int32(15), v28)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					if v32 == int32(0) {
						m.G0 = v7 + int32(16)
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v24
						v38 = *(*int32)(unsafe.Add(mBase, _consts[156]))
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v38
						v41 = *(*int32)(unsafe.Add(mBase, _consts[786]))
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = (v10 - v41 - int32(8)) >> (uint(int32(7)) % 32)
						F_errmsg_internal(m, int32(468224), v7)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							F_errfinish(m, int32(495788), int32(253), int32(352847))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return
							} else {
								m.G0 = v7 + int32(16)
								return
							}
						}
					}
				}
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v10)+104)) = int64(-1)
				v58 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v58
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = v58
				*(*int32)(unsafe.Add(mBase, uint32(v10)+96)) = v58
				F_ConditionVariableBroadcast(m, v10+int32(116))
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return
				} else {
					m.G0 = v7 + int32(16)
					return
				}
			}
		}
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
		v26 = *(*int32)(unsafe.Add(mBase, _consts[156]))
		if v24 != v26 {
			v28 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v18))) = v28
			v32 = F_errstart(m, int32(15), v28)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return
			} else {
				if v32 == int32(0) {
					m.G0 = v7 + int32(16)
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v24
					v38 = *(*int32)(unsafe.Add(mBase, _consts[156]))
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v38
					v41 = *(*int32)(unsafe.Add(mBase, _consts[786]))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = (v10 - v41 - int32(8)) >> (uint(int32(7)) % 32)
					F_errmsg_internal(m, int32(468224), v7)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						F_errfinish(m, int32(495788), int32(253), int32(352847))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							m.G0 = v7 + int32(16)
							return
						}
					}
				}
			}
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v10)+104)) = int64(-1)
			v58 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v58
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = v58
			*(*int32)(unsafe.Add(mBase, uint32(v10)+96)) = v58
			F_ConditionVariableBroadcast(m, v10+int32(116))
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return
			} else {
				m.G0 = v7 + int32(16)
				return
			}
		}
	}
}
func F_ProcArrayEndTransactionInternal(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
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
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v62 int32
	_ = v62
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v81 int64
	_ = v81
	v3 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, _consts[125]))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v8+v9<<(uint(int32(2))%32)))) = v3
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v3
	*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v3
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+73)) = uint8(v3)
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+124)))
	if v23&int32(14) != 0 {
		v27 = v23 & int32(241)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+124)) = uint8(v27)
		v30 = *(*int32)(unsafe.Add(mBase, _consts[125]))
		v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
		v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		*(*uint8)(unsafe.Add(mBase, uint32(v31+v32))) = uint8(v27)
	} else {
	}
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+276)))
	if v36 == int32(0) {
		v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+277)))
		if v39 != int32(1) {
		} else {
			v43 = v9 << (uint(int32(1)) % 32)
			v44 = int32(4412592)
			v45 = *(*int32)(unsafe.Add(mBase, _consts[125]))
			v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
			v48 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v43+v46))) = uint8(v48)
			v51 = *(*int32)(unsafe.Add(mBase, _consts[125]))
			v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
			*(*uint8)(unsafe.Add(mBase, uint32(v52+v43)+1)) = uint8(v48)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+276)) = uint16(v48)
		}
	} else {
		v43 = v9 << (uint(int32(1)) % 32)
		v44 = int32(4412592)
		v45 = *(*int32)(unsafe.Add(mBase, _consts[125]))
		v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
		v48 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v43+v46))) = uint8(v48)
		v51 = *(*int32)(unsafe.Add(mBase, _consts[125]))
		v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
		*(*uint8)(unsafe.Add(mBase, uint32(v52+v43)+1)) = uint8(v48)
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+276)) = uint16(v48)
	}
	v60 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	v61 = *(*int64)(unsafe.Add(mBase, uint32(v60)+48))
	v62 = base.I32_wrap_i64(v61)
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l1))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v62)) == int32(0) {
		v74 = base.B2i32(base.Ui32(v62) < base.Ui32(l1))
	} else {
		v74 = int32(base.Ui32(v62-l1) >> (uint(int32(31)) % 32))
	}
	v76 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	if v74 != 0 {
		*(*int64)(unsafe.Add(mBase, uint32(v76)+48)) = v61 + base.I64_extend_i32_s(l1-v62)
	} else {
	}
	v81 = *(*int64)(unsafe.Add(mBase, uint32(v76)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v76)+56)) = v81 + int64(1)
	return
}
func F_RemoveProcFromArray(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	F_ProcArrayRemove(m, v4, int32(0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_proc_exit_prepare(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v18 int32
	_ = v18
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, _consts[48])) = v2
	*(*int32)(unsafe.Add(mBase, _consts[763])) = v2
	*(*int32)(unsafe.Add(mBase, _consts[764])) = v2
	v18 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[163])) = v18
	*(*uint8)(unsafe.Add(mBase, _consts[765])) = uint8(v18)
	*(*int32)(unsafe.Add(mBase, _consts[337])) = v2
	*(*int32)(unsafe.Add(mBase, _consts[5])) = v2
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v2
	F_shmem_exit(m, l0)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v36 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v36 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
	v40 = *(*int32)(unsafe.Add(mBase, _consts[766]))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v40
	F_errmsg_internal(m, int32(397566), v6)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v50 = int32(4405584)
	v52 = *(*int32)(unsafe.Add(mBase, _consts[766]))
	v54 = v52 - int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[766])) = v54
	if int32(0) <= v54 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	F_errfinish(m, int32(497985), int32(202), int32(364161))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	v59 = v54
	goto L12
L10:
	;
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, _consts[766])) = int32(0)
	m.G0 = v6 + int32(16)
	return
L12:
	;
	v62 = v59 << (uint(int32(3)) % 32)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v62)+uint32(_consts[767])))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v62)+uint32(_consts[768])))
	m.T0[v68].(func(*base.Module, int32, int32))(m, l0, v65)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	goto L11
L14:
	;
	v71 = int32(4405584)
	v73 = *(*int32)(unsafe.Add(mBase, _consts[766]))
	v75 = v73 - int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[766])) = v75
	if int32(0) <= v75 {
		v59 = v75
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
}
