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
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
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
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v79 int64
	_ = v79
	var v84 int64
	_ = v84
	var v85 int64
	_ = v85
	var v88 int64
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	v11 = m.G0
	v13 = v11 - int32(96)
	m.G0 = v13
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_GetNewOidWithIndex[0]))
	if v16 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 + int32(96)
	return v119
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
	v112 = F_GetNewObjectId(m)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L10
	} else {
		goto L36
	}
L5:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_GetNewOidWithIndex[1]))
	if v29 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if base.Ui64(v88) < base.Ui64(int64(1000001)) {
		v119 = v38
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
	v35 = v13 + int32(48)
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
	F_ScanKeyInit(m, v35, l2, int32(3), int32(184), v38)
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
	v45 = F_systable_beginscan(m, l0, l1, v42, int32(_a_F_GetNewOidWithIndex_0), v42, v35)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v47 = F_systable_getnext(m, v45)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	F_systable_endscan(m, v45)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
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
	v54 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L10
	} else {
		goto L20
	}
L18:
	;
	v85 = v26
	goto L19
L19:
	;
	v88 = v25 + int64(1)
	if v47 != 0 {
		v25 = v88
		v26 = v85
		goto L5
	} else {
		goto L30
	}
L20:
	;
	if v54 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v56 + int32(4)
	F_errmsg(m, int32(_a_F_GetNewOidWithIndex_1), v13+int32(32))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L10
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v79 = v26 << (uint(int64(1)) % 64)
	if base.Ui64(v79) < base.Ui64(int64(128000001)) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v25
	F_errdetail_plural(m, int32(_a_F_GetNewOidWithIndex_2), int32(_a_F_GetNewOidWithIndex_3), base.I32_wrap_i64(v25), v13+int32(16))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L10
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(_a_F_GetNewOidWithIndex_4), int32(509), int32(_a_F_GetNewOidWithIndex_5))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L10
	} else {
		goto L26
	}
L26:
	;
	goto L23
L27:
	;
	v84 = v79
	goto L29
L28:
	;
	v84 = v26 + int64(128000000)
	goto L29
L29:
	;
	v85 = v84
	goto L19
L30:
	;
	goto L6
L31:
	;
	v93 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L10
	} else {
		goto L32
	}
L32:
	;
	if v93 == int32(0) {
		v119 = v38
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v97 + int32(4)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v88
	F_errmsg_plural(m, int32(_a_F_GetNewOidWithIndex_6), int32(_a_F_GetNewOidWithIndex_7), base.I32_wrap_i64(v88), v13)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L10
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_GetNewOidWithIndex_4), int32(534), int32(_a_F_GetNewOidWithIndex_5))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L10
	} else {
		goto L35
	}
L35:
	;
	v119 = v38
	goto L1
L36:
	;
	v119 = v112
	goto L1
}
func F_OidFunctionCall0Coll(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	v3 = m.G0
	v5 = v3 + int32(-64)
	m.G0 = v5
	v8 = v3 + int32(-48)
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_OidFunctionCall0Coll[0]))
	F_fmgr_info_cxt_security(m, l0, v8, v10, int32(0))
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
		*(*int32)(unsafe.Add(mBase, uint32(v5)+44)) = v8
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
		v28 = m.T0[v27].(func(*base.Module, int32) int32)(m, v3+int32(-20))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+60)))
			if v30 == int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v5)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v5))) = v37
					F_errmsg_internal(m, int32(_a_F_OidFunctionCall0Coll_0), v5)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_OidFunctionCall0Coll_1), int32(1123), int32(_a_F_OidFunctionCall0Coll_2))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
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
				return v28
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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	v11 = v8 + int32(16)
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_OidFunctionCall2Coll[0]))
	F_fmgr_info_cxt_security(m, l0, v11, v13, int32(0))
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
		*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = v11
		v35 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
		v36 = m.T0[v35].(func(*base.Module, int32) int32)(m, v8+int32(44))
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return int32(0)
		} else {
			v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+60)))
			if v38 == int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v45
					F_errmsg_internal(m, int32(_a_F_OidFunctionCall2Coll_0), v8)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_OidFunctionCall2Coll_1), int32(1165), int32(_a_F_OidFunctionCall2Coll_2))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
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
				return v36
			}
		}
	}
}
func F_oid_decrement(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(base.B2i32(l1 == v4))
	return l1 - base.B2i32(l1 != v4)
}
