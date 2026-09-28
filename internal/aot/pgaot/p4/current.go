package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetCurrentLSNForWaitType(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int64
	_ = v13
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v29 int64
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v37 int64
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int64
	_ = v49
	var v56 int64
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int64
	_ = v71
	var v72 int32
	_ = v72
	var v75 int64
	_ = v75
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	switch l0 {
	case 0:
		v71 = F_GetXLogReplayRecPtr(m, int32(0))
		mBase = m.M
		v72 = m.ExcPending
		if v72 != 0 {
			return int64(0)
		} else {
			v75 = v71
			m.G0 = v7 + int32(16)
			return v75
		}
	case 1:
		v10 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentLSNForWaitType[0]))
		v13 = base.AtomicRmwOr64(m, v10, int32(1464), int64(0))
		v15 = F_GetXLogReplayRecPtr(m, int32(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			if base.Ui64(v15) < base.Ui64(v13) {
				v20 = v13
			} else {
				v20 = v15
			}
			v75 = v20
			m.G0 = v7 + int32(16)
			return v75
		}
	case 2:
		v21 = int32(0)
		v23 = F_GetWalRcvFlushRecPtr(m, v21, v21)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int64(0)
		} else {
			v26 = F_GetXLogReplayRecPtr(m, int32(0))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int64(0)
			} else {
				if base.Ui64(v26) < base.Ui64(v23) {
					v29 = v23
				} else {
					v29 = v26
				}
				v75 = v29
				m.G0 = v7 + int32(16)
				return v75
			}
		}
	case 3:
		v30 = int32(0)
		v32 = int32(_a_F_GetCurrentLSNForWaitType_0)
		v33 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentLSNForWaitType[1]))
		v34 = int64(0)
		v37 = base.AtomicRmwCmpxchg64(m, v33, int32(272), v34, v34)
		*(*int64)(unsafe.Add(mBase, _c_F_GetCurrentLSNForWaitType[2])) = v37
		v42 = base.AtomicRmwOr32(m, v30, int32(_a_F_GetCurrentLSNForWaitType_1), v30)
		v45 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentLSNForWaitType[1]))
		v49 = base.AtomicRmwCmpxchg64(m, v45, int32(264), v34, v34)
		*(*int64)(unsafe.Add(mBase, _c_F_GetCurrentLSNForWaitType[3])) = v49
		v56 = *(*int64)(unsafe.Add(mBase, _c_F_GetCurrentLSNForWaitType[2]))
		v75 = v56
		m.G0 = v7 + int32(16)
		return v75
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v60 = m.ExcPending
		if v60 != 0 {
			return int64(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
			F_errmsg_internal(m, int32(_a_F_GetCurrentLSNForWaitType_2), v7)
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, int32(_a_F_GetCurrentLSNForWaitType_3), int32(146), int32(_a_F_GetCurrentLSNForWaitType_4))
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return int64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_GetCurrentTransactionId(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentTransactionId[0]))
	v5 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v4))))
	if v5 == int64(0) {
		F_AssignTransactionId(m, v4)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v12 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
			v13 = v12
			return base.I32_wrap_i64(v13)
		}
	} else {
		v13 = v5
		return base.I32_wrap_i64(v13)
	}
}
func F_current_schemas(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v45 int64
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	v2 = int32(0)
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_fetch_search_path(m, base.B2i32(v6 != int64(0)))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_list_free(m, v9)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L2
	} else {
		goto L18
	}
L2:
	;
	return int64(0)
L3:
	;
	if v9 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v16 = F_palloc(m, int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v21 = F_palloc(m, v18<<(uint(int32(3))%32))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L2
	} else {
		goto L8
	}
L7:
	;
	v57 = v2
	v58 = v16
	goto L1
L8:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v23 <= int32(0) {
		v57 = v2
		v58 = v21
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v27 = int32(0)
	v29 = v2
	goto L10
L10:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32+v27<<(uint(int32(2))%32))))
	v37 = F_get_namespace_name(m, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L12
	}
L11:
	;
	v57 = v50
	v58 = v21
	goto L1
L12:
	;
	if v37 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v45 = F_DirectFunctionCall1Coll(m, int32(534), int32(0), base.I64_extend_i32_u(v37))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L2
	} else {
		goto L16
	}
L14:
	;
	v50 = v29
	goto L15
L15:
	;
	v52 = v27 + int32(1)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v52 < v53 {
		v27 = v52
		v29 = v50
		goto L10
	} else {
		goto L17
	}
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21+v29<<(uint(int32(3))%32)))) = v45
	v50 = v29 + int32(1)
	goto L15
L17:
	;
	goto L11
L18:
	;
	v63 = F_construct_array_builtin(m, v58, v57, int32(19))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	return base.I64_extend_i32_u(v63)
}
