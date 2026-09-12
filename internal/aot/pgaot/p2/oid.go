package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetNewOidWithIndex(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v25 int64
	_ = v25
	var v26 int64
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v81 int64
	_ = v81
	var v86 int64
	_ = v86
	var v87 int64
	_ = v87
	var v90 int64
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
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
	var v120 int32
	_ = v120
	v11 = m.G0
	v13 = v11 - int32(96)
	m.G0 = v13
	v16 = *(*int32)(unsafe.Add(mBase, _consts[75]))
	if v16 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 + int32(96)
	return v120
L2:
	;
	v25 = int64(0)
	v26 = int64(1000000)
	goto L5
L3:
	;
	goto L4
L4:
	;
	v114 = F_GetNewObjectId(m)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L10
	} else {
		goto L36
	}
L5:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v29 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if base.Ui64(v90) < base.Ui64(int64(1000001)) {
		v120 = v38
		goto L1
	} else {
		goto L31
	}
L7:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v38 = F_GetNewObjectId(m)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L10
	} else {
		goto L12
	}
L10:
	;
	return int32(0)
L11:
	;
	goto L9
L12:
	;
	F_ScanKeyInit(m, v13+int32(48), l2, int32(3), int32(184), v38)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v42 = int32(1)
	v47 = F_systable_beginscan(m, l0, l1, v42, int32(4216240), v42, v13+int32(48))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v49 = F_systable_getnext(m, v47)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	F_systable_endscan(m, v47)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	if base.Ui64(v26) <= base.Ui64(v25) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v56 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L10
	} else {
		goto L20
	}
L18:
	;
	v87 = v26
	goto L19
L19:
	;
	v90 = v25 + int64(1)
	if v49 != 0 {
		v25 = v90
		v26 = v87
		goto L5
	} else {
		goto L30
	}
L20:
	;
	if v56 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v58 + int32(4)
	F_errmsg(m, int32(740437), v13+int32(32))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L10
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v81 = v26 << (uint(int64(1)) % 64)
	if base.Ui64(v81) < base.Ui64(int64(128000001)) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v25
	F_errdetail_plural(m, int32(609976), int32(609893), base.I32_wrap_i64(v25), v13+int32(16))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L10
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(523112), int32(509), int32(30437))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L10
	} else {
		goto L26
	}
L26:
	;
	goto L23
L27:
	;
	v86 = v81
	goto L29
L28:
	;
	v86 = v26 + int64(128000000)
	goto L29
L29:
	;
	v87 = v86
	goto L19
L30:
	;
	goto L6
L31:
	;
	v95 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L10
	} else {
		goto L32
	}
L32:
	;
	if v95 == int32(0) {
		v120 = v38
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v99 + int32(4)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v90
	F_errmsg_plural(m, int32(12842), int32(178297), base.I32_wrap_i64(v90), v13)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L10
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(523112), int32(534), int32(30437))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L10
	} else {
		goto L35
	}
L35:
	;
	v120 = v38
	goto L1
L36:
	;
	v120 = v114
	goto L1
}
func F_OidFunctionCall0Coll(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	v3 = m.G0
	v5 = v3 + int32(-64)
	m.G0 = v5
	v10 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	F_fmgr_info_cxt_security(m, l0, v3+int32(-48), v10, int32(0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(v5)+62)) = uint16(v16)
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+60)) = uint8(v16)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+56)) = v16
		*(*int64)(unsafe.Add(mBase, uint32(v5)+48)) = int64(0)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+44)) = v3 + int32(-48)
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
		v30 = m.T0[v29].(func(*base.Module, int32) int32)(m, v3+int32(-20))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+60)))
			if v32 == int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v5)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v5))) = v39
					F_errmsg_internal(m, int32(558657), v5)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(519609), int32(1123), int32(318863))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				m.G0 = v5 - int32(-64)
				return v30
			}
		}
	}
}
func F_OidFunctionCall2Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	v13 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	F_fmgr_info_cxt_security(m, l0, v8+int32(16), v13, int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v8)+76)) = uint8(v19)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+72)) = l3
		*(*uint8)(unsafe.Add(mBase, uint32(v8)+68)) = uint8(v19)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = l2
		v25 = int32(2)
		*(*uint16)(unsafe.Add(mBase, uint32(v8)+62)) = uint16(v25)
		*(*uint8)(unsafe.Add(mBase, uint32(v8)+60)) = uint8(v19)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+56)) = l1
		*(*int64)(unsafe.Add(mBase, uint32(v8)+48)) = int64(0)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = v8 + int32(16)
		v37 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
		v38 = m.T0[v37].(func(*base.Module, int32) int32)(m, v8+int32(44))
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return int32(0)
		} else {
			v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+60)))
			if v40 == int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v47
					F_errmsg_internal(m, int32(558657), v8)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(519609), int32(1165), int32(318821))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				m.G0 = v8 + int32(80)
				return v38
			}
		}
	}
}
func F_oid_decrement(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v4 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(base.B2i32(l1 == v4))
	v8 = l1 - int32(1)
	if base.Ui32(v8) <= base.Ui32(l1) {
		v11 = v8
	} else {
		v11 = v4
	}
	return v11
}
