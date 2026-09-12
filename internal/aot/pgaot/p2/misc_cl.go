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
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v8 < int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v141
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v140
	return v137
L2:
	;
	v82 = v78
	v85 = v79
	v86 = v80
	goto L27
L3:
	;
	if v16&int32(3) == int32(0) {
		v42 = v16
		goto L10
	} else {
		goto L11
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
		v78 = v12
		v79 = v13
		v80 = v8
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
	if v75 <= int32(0) {
		v137 = v16
		v140 = v75
		v141 = v18
		goto L1
	} else {
		goto L25
	}
L9:
	;
	v75 = v67 - v16
	goto L8
L10:
	;
	v46 = v42
	goto L19
L11:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v26 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v75 = int32(0)
	goto L8
L13:
	;
	goto L14
L14:
	;
	v31 = v16
	goto L15
L15:
	;
	v35 = v31 + int32(1)
	if v35&int32(3) == int32(0) {
		v42 = v35
		goto L10
	} else {
		goto L17
	}
L16:
	;
	v67 = v35
	goto L9
L17:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v40 != 0 {
		v31 = v35
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v55 = int32(-2139062144)
	if (int32(16843008)-v52|v52)&v55 == v55 {
		v46 = v46 + int32(4)
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v61 = v46
	goto L22
L21:
	;
	goto L20
L22:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v65 != 0 {
		v61 = v61 + int32(1)
		goto L22
	} else {
		goto L24
	}
L23:
	;
	v67 = v61
	goto L9
L24:
	;
	goto L23
L25:
	;
	v78 = v16
	v79 = v75
	v80 = v18
	goto L2
L26:
	;
	v115 = v85
	goto L32
L27:
	;
	v89 = int32(*(*int8)(unsafe.Add(mBase, uint32(v82))))
	goto L29
L28:
	;
	v137 = v104
	v140 = int32(0)
	v141 = v79 + v80
	goto L1
L29:
	;
	if base.B2i32(v89 == int32(32))|base.B2i32(base.Ui32((v89-int32(9))&int32(255)) < base.Ui32(int32(5))) == int32(0) {
		goto L26
	} else {
		goto L30
	}
L30:
	;
	v101 = int32(1)
	v104 = v82 + v101
	if v101 < v85 {
		v82 = v104
		v85 = v85 - v101
		v86 = v86 + v101
		goto L27
	} else {
		goto L31
	}
L31:
	;
	goto L28
L32:
	;
	v120 = int32(*(*int8)(unsafe.Add(mBase, uint32(v115+(v82-int32(1))))))
	goto L34
L33:
	;
	v137 = v82
	v140 = int32(0)
	v141 = v86
	goto L1
L34:
	;
	if base.B2i32(v120 == int32(32))|base.B2i32(base.Ui32((v120-int32(9))&int32(255)) < base.Ui32(int32(5))) == int32(0) {
		v137 = v82
		v140 = v115
		v141 = v86
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v132 = int32(1)
	if v132 < v115 {
		v115 = v115 - v132
		goto L32
	} else {
		goto L36
	}
L36:
	;
	goto L33
}
func F_clear_setitimer(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int64
	_ = v6
	var v14 int32
	_ = v14
	v2 = m.G0
	v3 = int32(32)
	v4 = v2 - v3
	m.G0 = v4
	v6 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v4)+24)) = v6
	*(*int64)(unsafe.Add(mBase, uint32(v4)+16)) = v6
	*(*int64)(unsafe.Add(mBase, uint32(v4)+8)) = v6
	*(*int64)(unsafe.Add(mBase, uint32(v4))) = v6
	v14 = F_setitimer(m, v4)
	mBase = m.M
	m.G0 = v4 + v3
	return
}
func F_clearerr(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if int32(0) <= v2 {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v5 & int32(-49)
		return
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v9 & int32(-49)
		return
	}
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
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
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
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+48)))
	v12 = v10 & int32(240)
	switch v12 {
	case 0:
		v35 = *(*int32)(unsafe.Add(mBase, _c_F_clog_redo[0]))
		v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+28))
		v37 = *(*int32)(unsafe.Add(mBase, uint32(v9)+64))
		v38 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
		v40 = int64(*(*uint16)(unsafe.Add(mBase, _c_F_clog_redo[1])))
		v41 = base.I64_rem_s(v38, v40)
		v45 = v36 + base.I32_wrap_i64(v41)<<(uint(int32(7))%32)
		v47 = F_LWLockAcquire(m, v45, int32(0))
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return
		} else {
			v49 = int32(_a_F_clog_redo_0)
			v51 = F_SimpleLruZeroPage(m, v49, v38)
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return
			} else {
				F_SimpleLruWritePage(m, v49, v51)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return
				} else {
					F_LWLockRelease(m, v45)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						m.G0 = v7 + int32(16)
						return
					}
				}
			}
		}
	default:
		F_errstart_cold(m, int32(23), int32(0))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = v12
			F_errmsg_internal(m, int32(_a_F_clog_redo_1), v7)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_clog_redo_2), int32(1142), int32(_a_F_clog_redo_3))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 16:
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+64))
		v14 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
		F_AdvanceOldestClogXid(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			F_SimpleLruTruncate(m, int32(_a_F_clog_redo_0), v14)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				m.G0 = v7 + int32(16)
				return
			}
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
func F_close_ps(m *base.Module, l0 int32) int32 {
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
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_palloc(m, int32(16))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = F_lseg_closept_point(m, v8, v5, v6)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)) {
				v19 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v19)
				v22 = int32(0)
			} else {
				v22 = v8
			}
			return v22
		}
	}
}
