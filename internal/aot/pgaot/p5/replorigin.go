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
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v24 int64
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+64))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+48)))
	v14 = v12 & int32(240)
	if v14 != 0 {
		if v14 == int32(16) {
			v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11))))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v29
			F_appendStringInfo(m, l0, int32(_a_F_replorigin_desc_0), v8+int32(16))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return
			} else {
				m.G0 = v8 + int32(32)
				return
			}
		} else {
			m.G0 = v8 + int32(32)
			return
		}
	} else {
		v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+8)))
		v18 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+10)))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v19
		*(*uint32)(unsafe.Add(mBase, uint32(v8)+8)) = uint32(v18)
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = v17
		v24 = int64(base.Ui64(v18) >> (uint(int64(32)) % 64))
		*(*uint32)(unsafe.Add(mBase, uint32(v8)+4)) = uint32(v24)
		F_appendStringInfo(m, l0, int32(_a_F_replorigin_desc_1), v8)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v121 int64
	_ = v121
	var v122 int32
	_ = v122
	var v123 int64
	_ = v123
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v16 = F_table_open(m, int32(_a_F_replorigin_drop_by_name_0), int32(3))
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
	F_LockSharedObject(m, int32(_a_F_replorigin_drop_by_name_0), v19, int32(8))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v26 = F_SearchSysCache1(m, int32(58), base.I64_extend_i32_u(v19))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L52
	}
L6:
	;
	m.G0 = v12 + int32(48)
	return
L7:
	;
	if v26 == int32(0) {
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
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_drop_by_name[0]))
	v44 = F_LWLockAcquire(m, v40+int32(_a_F_replorigin_drop_by_name_1), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L14
	}
L11:
	;
	F_UnlockSharedObject(m, int32(_a_F_replorigin_drop_by_name_0), v19, int32(8))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	F_relation_close(m, v16, int32(3))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	goto L6
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_drop_by_name[1]))
	if v47 <= int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v161 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_drop_by_name[0]))
	F_LWLockRelease(m, v161+int32(_a_F_replorigin_drop_by_name_1))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L46
	}
L16:
	;
	v57 = v47
	goto L17
L17:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_drop_by_name[2]))
	v63 = int32(0)
	goto L19
L18:
	;
	goto L15
L19:
	;
	v73 = v60 + v63<<(uint(int32(6))%32)
	v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v73))))
	if v19 != v74 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v73)+28))
	if int32(0) < v79 {
		goto L26
	} else {
		goto L27
	}
L21:
	;
	v77 = v63 + int32(1)
	if v57 != v77 {
		v63 = v77
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
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_drop_by_name[0]))
	F_LWLockRelease(m, v130+int32(_a_F_replorigin_drop_by_name_1))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L42
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
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+46)) = uint16(v19)
	F_XLogBeginInsert(m)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L39
	}
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_errcode(m, int32(100663621))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v91 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v73))))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v73)+24))
	if v92 != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	F_errfinish(m, int32(_a_F_replorigin_drop_by_name_2), int32(415), int32(_a_F_replorigin_drop_by_name_3))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L38
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v91
	F_errmsg(m, int32(_a_F_replorigin_drop_by_name_4), v12+int32(32))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v91
	F_errmsg(m, int32(_a_F_replorigin_drop_by_name_5), v12+int32(16))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L37
	}
L36:
	;
	goto L32
L37:
	;
	goto L32
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L39:
	;
	F_XLogRegisterData(m, v12+int32(46), int32(2))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v121 = F_XLogInsert(m, int32(19), int32(16))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v123 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v73)+8)) = v123
	v125 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v73))) = uint16(v125)
	*(*int64)(unsafe.Add(mBase, uint32(v73)+16)) = v123
	goto L15
L42:
	;
	F_ConditionVariableSleep(m, v73+int32(32), int32(134217777))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v141 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_drop_by_name[0]))
	v145 = F_LWLockAcquire(m, v141+int32(_a_F_replorigin_drop_by_name_1), int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_drop_by_name[1]))
	if int32(0) < v148 {
		v57 = v148
		goto L17
	} else {
		goto L45
	}
L45:
	;
	goto L18
L46:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_simple_heap_delete(m, v16, v26+int32(4))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_ReleaseCatCache(m, v26)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_relation_close(m, v16, int32(0))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	goto L6
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v19
	F_errmsg_internal(m, int32(_a_F_replorigin_drop_by_name_6), v12)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_replorigin_drop_by_name_2), int32(480), int32(_a_F_replorigin_drop_by_name_7))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
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
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_session_advance[0]))
	v9 = F_LWLockAcquire(m, v5+int32(44), int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_session_advance[0]))
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
		F_LWLockRelease(m, v12+int32(44))
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
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_session_get_progress[0]))
	v8 = F_LWLockAcquire(m, v4+int32(44), int32(1))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_session_get_progress[0]))
		v15 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
		F_LWLockRelease(m, v13+int32(44))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			return v15
		}
	}
}
