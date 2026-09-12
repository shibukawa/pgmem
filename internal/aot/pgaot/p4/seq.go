package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecSeqScanWithQual(m *base.Module, l0 int32) int32 {
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
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 float64
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	F_MemoryContextReset(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v12 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v10 + int32(16)
	return v68
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v22 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	goto L12
L7:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v25 = F_SeqNext(m, l0)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L9
L11:
	;
	v68 = v25
	goto L3
L12:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v35 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v38 = F_SeqNext(m, l0)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	if v38 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v68 = int32(0)
	goto L3
L20:
	;
	goto L21
L21:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+4)))
	if v43&int32(2) != 0 {
		v68 = v38
		goto L3
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v38
	v47 = int32(4536272)
	v48 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v50
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	v55 = m.T0[v54].(func(*base.Module, int32, int32, int32) int32)(m, v12, v13, v10+int32(15))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v48
	if v55 != 0 {
		v68 = v38
		goto L3
	} else {
		goto L24
	}
L24:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v59 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v60 = *(*float64)(unsafe.Add(mBase, uint32(v59)+240))
	*(*float64)(unsafe.Add(mBase, uint32(v59)+240)) = base.F64_add(v60, float64(1))
	goto L27
L26:
	;
	goto L27
L27:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	F_MemoryContextReset(m, v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	goto L12
}
func F_SeqNext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
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
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v9 == int32(0) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
		v14 = int32(0)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v12)+188))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
		v20 = m.T0[v19].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v12, v13, v14, v14, v14, int32(449))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v20
			v25 = v20
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+56))
			*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = v27
			v30 = *(*int32)(unsafe.Add(mBase, _consts[115]))
			if v30 == int32(0) {
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+188))
				v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+20))
				v54 = m.T0[v53].(func(*base.Module, int32, int32, int32) int32)(m, v25, v8, v6)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					if v54 != 0 {
						v56 = v6
					} else {
						v56 = int32(0)
					}
					return v56
				}
			} else {
				v34 = int32(*(*uint8)(unsafe.Add(mBase, _consts[116])))
				if v34&int32(1) != 0 {
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+188))
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+20))
					v54 = m.T0[v53].(func(*base.Module, int32, int32, int32) int32)(m, v25, v8, v6)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						if v54 != 0 {
							v56 = v6
						} else {
							v56 = int32(0)
						}
						return v56
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(343590), int32(0))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(333907), int32(1034), int32(86386))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
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
	} else {
		v25 = v9
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+56))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = v27
		v30 = *(*int32)(unsafe.Add(mBase, _consts[115]))
		if v30 == int32(0) {
			v51 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
			v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+188))
			v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+20))
			v54 = m.T0[v53].(func(*base.Module, int32, int32, int32) int32)(m, v25, v8, v6)
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return int32(0)
			} else {
				if v54 != 0 {
					v56 = v6
				} else {
					v56 = int32(0)
				}
				return v56
			}
		} else {
			v34 = int32(*(*uint8)(unsafe.Add(mBase, _consts[116])))
			if v34&int32(1) != 0 {
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+188))
				v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+20))
				v54 = m.T0[v53].(func(*base.Module, int32, int32, int32) int32)(m, v25, v8, v6)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					if v54 != 0 {
						v56 = v6
					} else {
						v56 = int32(0)
					}
					return v56
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(343590), int32(0))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(333907), int32(1034), int32(86386))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
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
}
func F_seq_identify(m *base.Module, l0 int32) int32 {
	var v6 int32
	_ = v6
	if base.Ui32(l0) < base.Ui32(int32(16)) {
		v6 = int32(548935)
	} else {
		v6 = int32(0)
	}
	return v6
}
