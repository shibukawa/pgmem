package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_replorigin_desc(m *base.Module, l0 int32, l1 int32) {
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
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v22 int64
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+64))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+48)))
	switch v12 & int32(240) {
	case 0:
		v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+8)))
		v16 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+10)))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v17
		*(*uint32)(unsafe.Add(mBase, uint32(v8)+8)) = uint32(v16)
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = v15
		v22 = int64(base.Ui64(v16) >> (uint(int64(32)) % 64))
		*(*uint32)(unsafe.Add(mBase, uint32(v8)+4)) = uint32(v22)
		F_appendStringInfo(m, l0, int32(487101), v8)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			m.G0 = v8 + int32(32)
			return
		}
	default:
		m.G0 = v8 + int32(32)
		return
	case 16:
		v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11))))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v27
		F_appendStringInfo(m, l0, int32(44174), v8+int32(16))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return
		} else {
			m.G0 = v8 + int32(32)
			return
		}
	}
}
func F_replorigin_drop_by_name(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v112 int64
	_ = v112
	var v113 int32
	_ = v113
	var v114 int64
	_ = v114
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v16 = F_table_open(m, int32(6000), int32(3))
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
	v19 = F_replorigin_by_name(m, l0, l1)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_LockSharedObject(m, int32(6000), v19, int32(8))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v25 = F_SearchSysCache1(m, int32(58), v19)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L47
	}
L6:
	;
	m.G0 = v12 + int32(32)
	return
L7:
	;
	if v25 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if l1 == int32(0) {
		goto L5
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v43 = F_LWLockAcquire(m, v39+int32(5120), int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L14
	}
L11:
	;
	F_UnlockSharedObject(m, int32(6000), v19, int32(8))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	F_sequence_close(m, v16, int32(3))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	goto L6
L14:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _consts[518]))
	if v46 <= int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v152 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v152+int32(5120))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L41
	}
L16:
	;
	v56 = v46
	goto L17
L17:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _consts[519]))
	v62 = int32(0)
	goto L19
L18:
	;
	goto L15
L19:
	;
	v72 = v59 + v62*int32(56)
	v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72))))
	if v19 != v73 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v72)+24))
	if v78 != 0 {
		goto L26
	} else {
		goto L27
	}
L21:
	;
	v76 = v62 + int32(1)
	if v56 != v76 {
		v62 = v76
		goto L19
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	goto L20
L24:
	;
	goto L15
L25:
	;
	v121 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v121+int32(5120))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L37
	}
L26:
	;
	if l2 == int32(0) {
		goto L25
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+30)) = uint16(v19)
	F_XLogBeginInsert(m)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L34
	}
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_errcode(m, int32(100663621))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72))))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v72)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v88
	F_errmsg(m, int32(479613), v12+int32(16))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(496692), int32(395), int32(230504))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L34:
	;
	F_XLogRegisterData(m, v12+int32(30), int32(2))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v112 = F_XLogInsert(m, int32(19), int32(16))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v114 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v72)+8)) = v114
	v116 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v72))) = uint16(v116)
	*(*int64)(unsafe.Add(mBase, uint32(v72)+16)) = v114
	goto L15
L37:
	;
	F_ConditionVariableSleep(m, v72+int32(28), int32(134217776))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v132 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v136 = F_LWLockAcquire(m, v132+int32(5120), int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _consts[518]))
	if int32(0) < v139 {
		v56 = v139
		goto L17
	} else {
		goto L40
	}
L40:
	;
	goto L18
L41:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_CatalogTupleDelete(m, v16, v25+int32(4))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_ReleaseCatCache(m, v25)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_sequence_close(m, v16, int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	goto L6
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v19
	F_errmsg_internal(m, int32(480086), v12)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(496692), int32(460), int32(378402))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_replorigin_reset(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v10 int32
	_ = v10
	v4 = int64(0)
	*(*int64)(unsafe.Add(mBase, _consts[110])) = v4
	*(*int64)(unsafe.Add(mBase, _consts[111])) = v4
	v10 = int32(0)
	*(*uint16)(unsafe.Add(mBase, _consts[109])) = uint16(v10)
	return
}
func F_replorigin_session_advance(m *base.Module, l0 int64, l1 int64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v16 int64
	_ = v16
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, _consts[520]))
	v9 = F_LWLockAcquire(m, v5+int32(40), int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, _consts[520]))
		v13 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
		if base.Ui64(v13) < base.Ui64(l1) {
			*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = l1
		} else {
		}
		v16 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
		if base.Ui64(v16) < base.Ui64(l0) {
			*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = l0
		} else {
		}
		F_LWLockRelease(m, v12+int32(40))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			return
		}
	}
}
func F_replorigin_session_get_progress(m *base.Module) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int64
	_ = v15
	var v19 int32
	_ = v19
	v4 = *(*int32)(unsafe.Add(mBase, _consts[520]))
	v8 = F_LWLockAcquire(m, v4+int32(40), int32(1))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, _consts[520]))
		v15 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
		F_LWLockRelease(m, v13+int32(40))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			return v15
		}
	}
}
