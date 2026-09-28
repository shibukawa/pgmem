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
	var v42 int32
	_ = v42
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
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int64
	_ = v80
	var v85 int64
	_ = v85
	var v86 int64
	_ = v86
	var v89 int64
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
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
	var v120 int32
	_ = v120
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
	v113 = F_GetNewObjectId(m)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
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
	if base.Ui64(v89) < base.Ui64(int64(1000001)) {
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
	v35 = v13 + int32(40)
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
	F_ScanKeyInit(m, v35, l2, int32(3), int32(184), base.I64_extend_i32_u(v38))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v43 = int32(1)
	v46 = F_systable_beginscan(m, l0, l1, v43, int32(_a_F_GetNewOidWithIndex_0), v43, v35)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v48 = F_systable_getnext(m, v46)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	F_systable_endscan(m, v46)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
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
	v55 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L10
	} else {
		goto L20
	}
L18:
	;
	v86 = v26
	goto L19
L19:
	;
	v89 = v25 + int64(1)
	if v48 != 0 {
		v25 = v89
		v26 = v86
		goto L5
	} else {
		goto L30
	}
L20:
	;
	if v55 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v57 + int32(4)
	F_errmsg(m, int32(_a_F_GetNewOidWithIndex_1), v13+int32(32))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L10
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v80 = v26 << (uint(int64(1)) % 64)
	if base.Ui64(v80) < base.Ui64(int64(128000001)) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v25
	F_errdetail_plural(m, int32(_a_F_GetNewOidWithIndex_2), int32(_a_F_GetNewOidWithIndex_3), base.I32_wrap_i64(v25), v13+int32(16))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L10
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(_a_F_GetNewOidWithIndex_4), int32(509), int32(_a_F_GetNewOidWithIndex_5))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L10
	} else {
		goto L26
	}
L26:
	;
	goto L23
L27:
	;
	v85 = v80
	goto L29
L28:
	;
	v85 = v26 + int64(128000000)
	goto L29
L29:
	;
	v86 = v85
	goto L19
L30:
	;
	goto L6
L31:
	;
	v94 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L10
	} else {
		goto L32
	}
L32:
	;
	if v94 == int32(0) {
		v120 = v38
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v98 + int32(4)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v89
	F_errmsg_plural(m, int32(_a_F_GetNewOidWithIndex_6), int32(_a_F_GetNewOidWithIndex_7), base.I32_wrap_i64(v89), v13)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L10
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_GetNewOidWithIndex_4), int32(534), int32(_a_F_GetNewOidWithIndex_5))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
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
	v120 = v113
	goto L1
}
func F_OidFunctionCall0Coll(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	v4 = m.G0
	v6 = v4 + int32(-64)
	m.G0 = v6
	v9 = v4 + int32(-52)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_OidFunctionCall0Coll[0]))
	F_fmgr_info_cxt_security(m, l0, v9, v11, int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(v6)+58)) = uint16(v17)
		*(*uint8)(unsafe.Add(mBase, uint32(v6)+56)) = uint8(v17)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+52)) = v17
		*(*int64)(unsafe.Add(mBase, uint32(v6)+44)) = int64(0)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = v9
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
		v29 = m.T0[v28].(func(*base.Module, int32) int64)(m, v4+int32(-24))
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int64(0)
		} else {
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+56)))
			if v31 == int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int64(0)
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = v38
					F_errmsg_internal(m, int32(_a_F_OidFunctionCall0Coll_0), v6)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_OidFunctionCall0Coll_1), int32(1125), int32(_a_F_OidFunctionCall0Coll_2))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				m.G0 = v6 - int32(-64)
				return v29
			}
		}
	}
}
func F_OidFunctionCall2Coll(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64) int64 {
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
	var v36 int64
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
	v8 = v6 - int32(96)
	m.G0 = v8
	v11 = v8 + int32(12)
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_OidFunctionCall2Coll[0]))
	F_fmgr_info_cxt_security(m, l0, v11, v13, int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int64(0)
	} else {
		v19 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v8)+88)) = uint8(v19)
		*(*int64)(unsafe.Add(mBase, uint32(v8)+80)) = l3
		*(*uint8)(unsafe.Add(mBase, uint32(v8)+72)) = uint8(v19)
		*(*int64)(unsafe.Add(mBase, uint32(v8)+64)) = l2
		v25 = int32(2)
		*(*uint16)(unsafe.Add(mBase, uint32(v8)+58)) = uint16(v25)
		*(*uint8)(unsafe.Add(mBase, uint32(v8)+56)) = uint8(v19)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = l1
		*(*int64)(unsafe.Add(mBase, uint32(v8)+44)) = int64(0)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v11
		v35 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
		v36 = m.T0[v35].(func(*base.Module, int32) int64)(m, v8+int32(40))
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return int64(0)
		} else {
			v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+56)))
			if v38 == int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int64(0)
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v45
					F_errmsg_internal(m, int32(_a_F_OidFunctionCall2Coll_0), v8)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_OidFunctionCall2Coll_1), int32(1167), int32(_a_F_OidFunctionCall2Coll_2))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				m.G0 = v8 + int32(96)
				return v36
			}
		}
	}
}
func F_oid_decrement(m *base.Module, l0 int32, l1 int64, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int64
	_ = v13
	v4 = base.I32_wrap_i64(l1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(base.B2i32(v4 == int32(0)))
	if v4 != 0 {
		v13 = (l1 - int64(1)) & int64(4294967295)
	} else {
		v13 = int64(0)
	}
	return v13
}
