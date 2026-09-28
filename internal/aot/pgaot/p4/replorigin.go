package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_replorigin_advance(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v143 int64
	_ = v143
	var v144 int32
	_ = v144
	var v147 int64
	_ = v147
	var v154 int64
	_ = v154
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	v1 = l0
	v4 = l3
	v6 = int32(0)
	v13 = m.G0
	v15 = v13 + int32(-64)
	m.G0 = v15
	if v1 != int32(_a_F_replorigin_advance_0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_advance[0]))
	v24 = F_LWLockAcquire(m, v20+int32(_a_F_replorigin_advance_1), int32(0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v15 - int32(-64)
	return
L4:
	;
	return
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_advance[1]))
	if int32(0) < v27 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	if l4 != 0 {
		goto L39
	} else {
		goto L40
	}
L7:
	;
	v126 = F_LWLockAcquire(m, v87+int32(44), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L4
	} else {
		goto L38
	}
L8:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_advance[2]))
	v38 = v6
	v40 = v6
	goto L11
L9:
	;
	goto L10
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L33
	}
L11:
	;
	v46 = v31 + v40<<(uint(int32(6))%32)
	v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46))))
	if v47|v38 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	if v87 != 0 {
		goto L7
	} else {
		goto L32
	}
L13:
	;
	v89 = v40 + int32(1)
	if v89 != v27 {
		v38 = v87
		v40 = v89
		goto L11
	} else {
		goto L31
	}
L14:
	;
	v87 = v46
	goto L13
L15:
	;
	goto L16
L16:
	;
	if v1 != v47 {
		v87 = v38
		goto L13
	} else {
		goto L17
	}
L17:
	;
	v55 = F_LWLockAcquire(m, v46+int32(44), int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v46)+28))
	if v57 <= int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v129 = v46
	goto L6
L20:
	;
	goto L21
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	F_errcode(m, int32(100663621))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46))))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v46)+24))
	if v68 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	F_errfinish(m, int32(_a_F_replorigin_advance_2), int32(989), int32(_a_F_replorigin_advance_3))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L4
	} else {
		goto L30
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v67
	F_errmsg(m, int32(_a_F_replorigin_advance_4), v13+int32(-32))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L4
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v67
	F_errmsg(m, int32(_a_F_replorigin_advance_5), v13+int32(-48))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L4
	} else {
		goto L29
	}
L28:
	;
	goto L24
L29:
	;
	goto L24
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L31:
	;
	goto L12
L32:
	;
	goto L10
L33:
	;
	F_errcode(m, int32(_a_F_replorigin_advance_6))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v1
	F_errmsg(m, int32(_a_F_replorigin_advance_7), v15)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	F_errhint(m, int32(_a_F_replorigin_advance_8), int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_replorigin_advance_2), int32(1000), int32(_a_F_replorigin_advance_3))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L38:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v87))) = uint16(v1)
	v129 = v87
	goto L6
L39:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+58)) = uint8(v4)
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+56)) = uint16(v1)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+48)) = l1
	F_XLogBeginInsert(m)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L4
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	if v4 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L42:
	;
	F_XLogRegisterData(m, v13+int32(-16), int32(16))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	v143 = F_XLogInsert(m, int32(19), int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	goto L41
L45:
	;
	F_LWLockRelease(m, v129+int32(44))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L4
	} else {
		goto L58
	}
L46:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v129)+16)) = l2
	goto L45
L47:
	;
	v147 = *(*int64)(unsafe.Add(mBase, uint32(v129)+8))
	if base.Ui64(v147) < base.Ui64(l1) {
		goto L51
	} else {
		goto L52
	}
L48:
	;
	goto L49
L49:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v129)+8)) = l1
	if l2 == int64(0) {
		goto L45
	} else {
		goto L57
	}
L50:
	;
	v154 = *(*int64)(unsafe.Add(mBase, uint32(v129)+16))
	if base.Ui64(l2) <= base.Ui64(v154) {
		goto L45
	} else {
		goto L56
	}
L51:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v129)+8)) = l1
	if l2 != int64(0) {
		goto L50
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	if l2 == int64(0) {
		goto L45
	} else {
		goto L55
	}
L54:
	;
	goto L45
L55:
	;
	goto L50
L56:
	;
	goto L46
L57:
	;
	goto L46
L58:
	;
	v165 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_advance[0]))
	F_LWLockRelease(m, v165+int32(_a_F_replorigin_advance_1))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	goto L3
}
func F_replorigin_redo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int64
	_ = v40
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+48)))
	v14 = v12 & int32(240)
	if v14 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return
L2:
	;
	if v14 == int32(16) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59)+8)))
	v61 = *(*int64)(unsafe.Add(mBase, uint32(v59)))
	v62 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+10)))
	F_replorigin_advance(m, v60, v61, v62, v63, int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L15
	} else {
		goto L19
	}
L5:
	;
	v17 = int32(0)
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_redo[0]))
	if v19 <= v17 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L15
	} else {
		goto L16
	}
L8:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_redo[1]))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24))))
	v26 = v17
	goto L9
L9:
	;
	v34 = v23 + v26<<(uint(int32(6))%32)
	v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34))))
	if v25 != v35 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v40 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v34)+8)) = v40
	v42 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v34))) = uint16(v42)
	*(*int64)(unsafe.Add(mBase, uint32(v34)+16)) = v40
	goto L1
L11:
	;
	v38 = v26 + int32(1)
	if v19 != v38 {
		v26 = v38
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	goto L10
L14:
	;
	goto L1
L15:
	;
	return
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v14
	F_errmsg_internal(m, int32(_a_F_replorigin_redo_0), v9)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(_a_F_replorigin_redo_1), int32(908), int32(_a_F_replorigin_redo_2))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L19:
	;
	goto L1
}
func F_replorigin_xact_clear(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v8 int32
	_ = v8
	v2 = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_replorigin_xact_clear[0])) = v2
	*(*int64)(unsafe.Add(mBase, _c_F_replorigin_xact_clear[1])) = v2
	v8 = int32(0)
	*(*uint16)(unsafe.Add(mBase, _c_F_replorigin_xact_clear[2])) = uint16(v8)
	return
}
