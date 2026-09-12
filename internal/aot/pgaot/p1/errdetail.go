package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_errdetail_log(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = int32(4513196)
	v13 = *(*int32)(unsafe.Add(mBase, _consts[1172]))
	*(*int32)(unsafe.Add(mBase, _consts[1172])) = v13 + int32(1)
	v18 = *(*int32)(unsafe.Add(mBase, _consts[1173]))
	if int32(0) <= v18 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v21 = int32(4520272)
	v22 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v25 = v18 * int32(100)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_consts[1178])))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v28
	F_initStringInfo(m, v9+int32(16))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1173])) = int32(-1)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L21
	}
L4:
	;
	return
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_consts[1177])))
	*(*int32)(unsafe.Add(mBase, _consts[137])) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
	v40 = F_appendStringInfoVA(m, v9+int32(16), l0, l1)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if v40 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v45 = v40
	goto L10
L8:
	;
	goto L9
L9:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_consts[1181])))
	if v66 != 0 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	F_enlargeStringInfo(m, v9+int32(16), v45)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L4
	} else {
		goto L12
	}
L11:
	;
	goto L9
L12:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_consts[1177])))
	*(*int32)(unsafe.Add(mBase, _consts[137])) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
	v58 = F_appendStringInfoVA(m, v9+int32(16), l0, l1)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	if v58 != 0 {
		v45 = v58
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	F_pfree(m, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L4
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	v70 = F_pstrdup(m, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_consts[1181]))) = v70
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	F_pfree(m, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v22
	v78 = int32(4513196)
	v80 = *(*int32)(unsafe.Add(mBase, _consts[1172]))
	*(*int32)(unsafe.Add(mBase, _consts[1172])) = v80 - int32(1)
	m.G0 = v9 + int32(32)
	return
L21:
	;
	F_errmsg_internal(m, int32(455099), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(499310), int32(1258), int32(328135))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_errdetail_recovery_conflict(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v3 = l0 - int32(7)
	if base.Ui32(v3) <= base.Ui32(int32(6)) {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v3<<(uint(int32(2))%32))+uint32(_consts[828])))
		F_errdetail(m, v10, int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
func F_errdetail_relkind_not_supported(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	switch l0 - int32(73) {
	case 0:
		v33 = int32(597168)
		F_errdetail(m, v33, int32(0))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return
		} else {
			m.G0 = v6 + int32(16)
			return
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
			F_errmsg_internal(m, int32(691107), v6)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				F_errfinish(m, int32(494899), int32(49), int32(442682))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 10:
		v33 = int32(603367)
		F_errdetail(m, v33, int32(0))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return
		} else {
			m.G0 = v6 + int32(16)
			return
		}
	case 26:
		v33 = int32(599925)
		F_errdetail(m, v33, int32(0))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return
		} else {
			m.G0 = v6 + int32(16)
			return
		}
	case 29:
		v33 = int32(601866)
		F_errdetail(m, v33, int32(0))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return
		} else {
			m.G0 = v6 + int32(16)
			return
		}
	case 32:
		v33 = int32(597037)
		F_errdetail(m, v33, int32(0))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return
		} else {
			m.G0 = v6 + int32(16)
			return
		}
	case 36:
		v33 = int32(584642)
		F_errdetail(m, v33, int32(0))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return
		} else {
			m.G0 = v6 + int32(16)
			return
		}
	case 39:
		v33 = int32(602128)
		F_errdetail(m, v33, int32(0))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return
		} else {
			m.G0 = v6 + int32(16)
			return
		}
	case 41:
		v33 = int32(601658)
		F_errdetail(m, v33, int32(0))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return
		} else {
			m.G0 = v6 + int32(16)
			return
		}
	case 43:
		v33 = int32(602314)
		F_errdetail(m, v33, int32(0))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return
		} else {
			m.G0 = v6 + int32(16)
			return
		}
	case 45:
		v33 = int32(584599)
		F_errdetail(m, v33, int32(0))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return
		} else {
			m.G0 = v6 + int32(16)
			return
		}
	}
}
