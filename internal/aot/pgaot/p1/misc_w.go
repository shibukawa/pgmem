package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_WaitIO(m *base.Module, l0 int32) {
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int64
	_ = v86
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[745]))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = v11 + v12<<(uint(int32(4))%32)
	F_ConditionVariablePrepareToSleep(m, v15)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v19 = l0 + int32(36)
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = int32(234095)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = int32(6259)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = int32(506700)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = int64(0)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v36 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v35 | v36
	if v35&v36 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L33
	}
L5:
	;
	goto L8
L6:
	;
	v58 = v35
	goto L7
L7:
	;
	v64 = int32(4142620)
	v65 = *(*int32)(unsafe.Add(mBase, _consts[276]))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(24))+8))
	if v67 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L8:
	;
	F_perform_spin_delay(m, v8+int32(24))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	v58 = v50
	goto L7
L10:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v51 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v50 | v51
	if v50&v51 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v84
	v86 = *(*int64)(unsafe.Add(mBase, uint32(v19)))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v58 & int32(-4194305)
	if v58&int32(67108864) != 0 {
		goto L23
	} else {
		goto L24
	}
L13:
	;
	goto L12
L14:
	;
	*(*int32)(unsafe.Add(mBase, _consts[276])) = v82
	goto L13
L15:
	;
	if int32(999) < v65 {
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if v65 < int32(11) {
		goto L13
	} else {
		goto L22
	}
L18:
	;
	v72 = int32(900)
	if v72 <= v65 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v75 = v72
	goto L21
L20:
	;
	v75 = v65
	goto L21
L21:
	;
	v82 = v75 + int32(100)
	goto L14
L22:
	;
	v82 = v65 - int32(1)
	goto L14
L23:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(8))))
	goto L26
L24:
	;
	goto L25
L25:
	;
	goto L4
L26:
	;
	if v95 != int32(-1) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	F_pgaio_wref_wait(m, v8+int32(8))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	F_ConditionVariableSleep(m, v15, int32(134217736))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	F_ConditionVariablePrepareToSleep(m, v15)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	goto L3
L32:
	;
	goto L3
L33:
	;
	m.G0 = v8 + int32(48)
	return
}
func F_WaitLatchOrSocket(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
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
	var v47 int32
	_ = v47
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v14 = F_CreateWaitEventSet(m, v12, int32(3))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		F_AddWaitEventToSet(m, v14, int32(1), int32(-1), l0)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v23 = int32(*(*uint8)(unsafe.Add(mBase, _consts[131])))
			if v23 == int32(1) {
				F_AddWaitEventToSet(m, v14, int32(32), int32(-1), int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					F_AddWaitEventToSet(m, v14, int32(2), l1, int32(0))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						v36 = F_WaitEventSetWait(m, v14, l2, v9, int32(1), l3)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
							F_FreeWaitEventSet(m, v14)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								m.G0 = v9 + int32(16)
								if v36 != 0 {
									v47 = v38 & int32(151)
								} else {
									v47 = int32(8)
								}
								return v47
							}
						}
					}
				}
			} else {
				F_AddWaitEventToSet(m, v14, int32(2), l1, int32(0))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v36 = F_WaitEventSetWait(m, v14, l2, v9, int32(1), l3)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
						F_FreeWaitEventSet(m, v14)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(16)
							if v36 != 0 {
								v47 = v38 & int32(151)
							} else {
								v47 = int32(8)
							}
							return v47
						}
					}
				}
			}
		}
	}
}
func F___wasm_setjmp_test(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if l1 == v4 {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v7 = v6
	} else {
		v7 = int32(0)
	}
	return v7
}
func F_week_num(m *base.Module, l0 int32) int32 {
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
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	v5 = int32(53)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = int32(7)
	v11 = base.I32_rem_u_s(v7+int32(6), v10)
	v16 = base.I32_div_u_s(v6-v11+v10, v10)
	v17 = v7 - v6
	v21 = base.I32_rem_u_s(v17+int32(369), v10)
	v24 = v16 + base.B2i32(base.Ui32(v21) < base.Ui32(int32(3)))
	if v24 != v5 {
		if v24 != 0 {
			v89 = v24
			return v89
		} else {
			v27 = int32(52)
			v31 = base.I32_rem_u_s(v17+int32(6), int32(7))
			switch v31 - int32(4) {
			case 0:
				return int32(53)
			case 1:
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v36 = base.I32_rem_s(v34, int32(400))
				v38 = v36 - int32(1)
				if int32(2147481747) < v38 {
					v43 = v38 - int32(2000)
				} else {
					v43 = v38
				}
				if v43&int32(3) != 0 {
					v57 = int32(0)
				} else {
					v48 = v43 + int32(1900)
					v50 = base.I32_rem_s(v48, int32(100))
					if v50 != 0 {
						v57 = int32(1)
					} else {
						v53 = base.I32_rem_s(v48, int32(400))
						v57 = base.B2i32(v53 == int32(0))
					}
				}
				if v57 == int32(0) {
					v89 = v27
					return v89
				} else {
					return int32(53)
				}
			default:
				v89 = v27
				return v89
			}
		}
	} else {
		v65 = base.I32_rem_u_s(v17+int32(371), int32(7))
		switch v65 - int32(3) {
		case 0:
			v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if int32(2147481747) < v68 {
				v73 = v68 - int32(2000)
			} else {
				v73 = v68
			}
			if v73&int32(3) != 0 {
				v87 = int32(0)
			} else {
				v78 = v73 + int32(1900)
				v80 = base.I32_rem_s(v78, int32(100))
				if v80 != 0 {
					v87 = int32(1)
				} else {
					v83 = base.I32_rem_s(v78, int32(400))
					v87 = base.B2i32(v83 == int32(0))
				}
			}
			if v87 != 0 {
				v89 = v5
			} else {
				v89 = int32(1)
			}
		case 1:
			v89 = v5
		default:
			v89 = int32(1)
		}
		return v89
	}
}
func F_write(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l1
	v16 = m.Wasi_snapshot_preview1.Fd_write(m, l0, v7+int32(8), int32(1), v7+int32(4))
	mBase = m.M
	if v16 == int32(0) {
		v23 = int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[137])) = v16
		v23 = int32(-1)
	}
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	m.G0 = v7 + int32(16)
	if v23 != 0 {
		v29 = int32(-1)
	} else {
		v29 = v24
	}
	return v29
}
func F_writetup_index(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+6)))
	v13 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v10&int32(8191) + v13
	F_LogicalTapeWrite(m, l1, v7+int32(12), v13)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return
	} else {
		v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+6)))
		F_LogicalTapeWrite(m, l1, v9, v21&int32(8191))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
			if v26&int32(1) != 0 {
				F_LogicalTapeWrite(m, l1, v7+int32(12), int32(4))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					m.G0 = v7 + int32(16)
					return
				}
			} else {
				m.G0 = v7 + int32(16)
				return
			}
		}
	}
}
