package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CleanQuerytext(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v8 < int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v85
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v84
	return v81
L2:
	;
	v26 = v22
	v29 = v23
	v30 = v24
	goto L9
L3:
	;
	v19 = F_strlen(m, v16)
	mBase = m.M
	if v19 <= int32(0) {
		v81 = v16
		v84 = v19
		v85 = v18
		goto L1
	} else {
		goto L8
	}
L4:
	;
	v16 = l0
	v18 = int32(0)
	goto L3
L5:
	;
	goto L6
L6:
	;
	v12 = l0 + v8
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if int32(0) < v13 {
		v22 = v12
		v23 = v13
		v24 = v8
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v16 = v12
	v18 = v8
	goto L3
L8:
	;
	v22 = v16
	v23 = v19
	v24 = v18
	goto L2
L9:
	;
	v33 = int32(*(*int8)(unsafe.Add(mBase, uint32(v26))))
	goto L11
L10:
	;
	v81 = v75
	v84 = int32(0)
	v85 = v23 + v24
	goto L1
L11:
	;
	if base.B2i32(v33 == int32(32))|base.B2i32(base.Ui32((v33-int32(9))&int32(255)) < base.Ui32(int32(5))) == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v48 = v29
	goto L15
L13:
	;
	goto L14
L14:
	;
	v72 = int32(1)
	v75 = v26 + v72
	if v72 < v29 {
		v26 = v75
		v29 = v29 - v72
		v30 = v30 + v72
		goto L9
	} else {
		goto L20
	}
L15:
	;
	v55 = int32(*(*int8)(unsafe.Add(mBase, uint32(v26+v48-int32(1)))))
	goto L17
L16:
	;
	v81 = v26
	v84 = int32(0)
	v85 = v30
	goto L1
L17:
	;
	if base.B2i32(v55 == int32(32))|base.B2i32(base.Ui32((v55-int32(9))&int32(255)) < base.Ui32(int32(5))) == int32(0) {
		v81 = v26
		v84 = v48
		v85 = v30
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v67 = int32(1)
	if v67 < v48 {
		v48 = v48 - v67
		goto L15
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	goto L10
}
func F_clearerr(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v2 & int32(-49)
	return
}
func F_clog_identify(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v5 = l0 & int32(240)
	if v5 == int32(16) {
		v8 = int32(_a_F_clog_identify_0)
	} else {
		v8 = int32(0)
	}
	if v5 != 0 {
		v10 = v8
	} else {
		v10 = int32(_a_F_clog_identify_1)
	}
	return v10
}
func F_clog_redo(m *base.Module, l0 int32) {
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
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v40 int32
	_ = v40
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+48)))
	v12 = v10 & int32(240)
	if v12 != 0 {
		if v12 == int32(16) {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v9)+64))
			v16 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
			F_AdvanceOldestClogXid(m, v17)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				F_SimpleLruTruncate(m, int32(_a_F_clog_redo_0), v16)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					m.G0 = v7 + int32(16)
					return
				}
			}
		} else {
			F_errstart_cold(m, int32(24), int32(0))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v12
				F_errmsg_internal(m, int32(_a_F_clog_redo_1), v7)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_clog_redo_2), int32(1115), int32(_a_F_clog_redo_3))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v37 = *(*int32)(unsafe.Add(mBase, uint32(v9)+64))
		v38 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
		F_SimpleLruZeroAndWritePage(m, int32(_a_F_clog_redo_0), v38)
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return
		} else {
			m.G0 = v7 + int32(16)
			return
		}
	}
}
func F_close(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	v2 = m.Wasi_snapshot_preview1.Fd_close(m, l0)
	mBase = m.M
	if v2 != int32(27) {
		v6 = v2
	} else {
		v6 = int32(0)
	}
	if v6 == int32(0) {
		v13 = int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_close[0])) = v6
		v13 = int32(-1)
	}
	return v13
}
func F_close_ps(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = F_palloc(m, int32(16))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v12 = F_lseg_closept_point(m, v8, v5, v6)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)) {
				v19 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v19)
				return int64(0)
			} else {
				return base.I64_extend_i32_u(v8)
			}
		}
	}
}
