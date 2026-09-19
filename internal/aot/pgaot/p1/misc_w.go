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
	var v37 int32
	_ = v37
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v82 int64
	_ = v82
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_WaitIO[0]))
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
	*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = int32(_a_F_WaitIO_0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = int32(_a_F_WaitIO_1)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = int32(_a_F_WaitIO_2)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = int64(0)
	v35 = int32(_a_F_WaitIO_3)
	v37 = base.AtomicRmwOr32(m, l0, int32(24), v35)
	if v37&v35 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L33
	}
L5:
	;
	goto L8
L6:
	;
	v56 = v37
	goto L7
L7:
	;
	v62 = int32(_a_F_WaitIO_4)
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_WaitIO[1]))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(24))+8))
	if v65 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L8:
	;
	F_perform_spin_delay(m, v8+int32(24))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	v56 = v51
	goto L7
L10:
	;
	v49 = int32(_a_F_WaitIO_3)
	v51 = base.AtomicRmwOr32(m, l0, int32(24), v49)
	if v51&v49 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v82 = *(*int64)(unsafe.Add(mBase, uint32(v19)))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v82
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v56 & int32(-4194305)
	if v56&int32(67108864) != 0 {
		goto L23
	} else {
		goto L24
	}
L13:
	;
	goto L12
L14:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WaitIO[1])) = v80
	goto L13
L15:
	;
	if int32(999) < v63 {
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if v63 < int32(11) {
		goto L13
	} else {
		goto L22
	}
L18:
	;
	v70 = int32(900)
	if v70 <= v63 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v73 = v70
	goto L21
L20:
	;
	v73 = v63
	goto L21
L21:
	;
	v80 = v73 + int32(100)
	goto L14
L22:
	;
	v80 = v63 - int32(1)
	goto L14
L23:
	;
	v92 = v8 + int32(8)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	goto L26
L24:
	;
	goto L25
L25:
	;
	goto L4
L26:
	;
	if v93 != int32(-1) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	F_pgaio_wref_wait(m, v92)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
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
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	F_ConditionVariablePrepareToSleep(m, v15)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
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
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_WaitLatchOrSocket[0]))
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
			v23 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WaitLatchOrSocket[1])))
			if v23&int32(1) != 0 {
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
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if l1 == v3 {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v7 = v5
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
		*(*int32)(unsafe.Add(mBase, _c_F_write[0])) = v16
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
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+6)))
	v14 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v11&int32(_a_F_writetup_index_0) + v14
	v18 = v8 + int32(12)
	F_LogicalTapeWrite(m, l1, v18, v14)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return
	} else {
		v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+6)))
		F_LogicalTapeWrite(m, l1, v10, v22&int32(_a_F_writetup_index_0))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
			if v27&int32(1) != 0 {
				F_LogicalTapeWrite(m, l1, v18, int32(4))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					m.G0 = v8 + int32(16)
					return
				}
			} else {
				m.G0 = v8 + int32(16)
				return
			}
		}
	}
}
