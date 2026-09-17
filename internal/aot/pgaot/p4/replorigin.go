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
	var v37 int32
	_ = v37
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
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v137 int64
	_ = v137
	var v138 int32
	_ = v138
	var v141 int64
	_ = v141
	var v148 int64
	_ = v148
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	v1 = l0
	v4 = l3
	v6 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(48)
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
	m.G0 = v15 + int32(48)
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
		goto L34
	} else {
		goto L35
	}
L7:
	;
	v120 = F_LWLockAcquire(m, v81+int32(40), int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L4
	} else {
		goto L33
	}
L8:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_advance[2]))
	v37 = v6
	v40 = v6
	goto L11
L9:
	;
	goto L10
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L4
	} else {
		goto L28
	}
L11:
	;
	v46 = v31 + v40*int32(56)
	v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46))))
	if v47|v37 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	if v81 != 0 {
		goto L7
	} else {
		goto L27
	}
L13:
	;
	v83 = v40 + int32(1)
	if v83 != v27 {
		v37 = v81
		v40 = v83
		goto L11
	} else {
		goto L26
	}
L14:
	;
	v81 = v46
	goto L13
L15:
	;
	goto L16
L16:
	;
	if v1 != v47 {
		v81 = v37
		goto L13
	} else {
		goto L17
	}
L17:
	;
	v55 = F_LWLockAcquire(m, v46+int32(40), int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v46)+24))
	if v57 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v123 = v46
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v67
	F_errmsg(m, int32(_a_F_replorigin_advance_2), v15+int32(16))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_replorigin_advance_3), int32(969), int32(_a_F_replorigin_advance_4))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L26:
	;
	goto L12
L27:
	;
	goto L10
L28:
	;
	F_errcode(m, int32(_a_F_replorigin_advance_5))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v1
	F_errmsg(m, int32(_a_F_replorigin_advance_6), v15)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	F_errhint(m, int32(_a_F_replorigin_advance_7), int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_replorigin_advance_3), int32(980), int32(_a_F_replorigin_advance_4))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L33:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v81))) = uint16(v1)
	v123 = v81
	goto L6
L34:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+42)) = uint8(v4)
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+40)) = uint16(v1)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+32)) = l1
	F_XLogBeginInsert(m)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	if v4 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L37:
	;
	F_XLogRegisterData(m, v15+int32(32), int32(16))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	v137 = F_XLogInsert(m, int32(19), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	goto L36
L40:
	;
	F_LWLockRelease(m, v123+int32(40))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L4
	} else {
		goto L53
	}
L41:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v123)+16)) = l2
	goto L40
L42:
	;
	v141 = *(*int64)(unsafe.Add(mBase, uint32(v123)+8))
	if base.Ui64(v141) < base.Ui64(l1) {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	goto L44
L44:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v123)+8)) = l1
	if l2 == int64(0) {
		goto L40
	} else {
		goto L52
	}
L45:
	;
	v148 = *(*int64)(unsafe.Add(mBase, uint32(v123)+16))
	if base.Ui64(l2) <= base.Ui64(v148) {
		goto L40
	} else {
		goto L51
	}
L46:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v123)+8)) = l1
	if l2 != int64(0) {
		goto L45
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	if l2 == int64(0) {
		goto L40
	} else {
		goto L50
	}
L49:
	;
	goto L40
L50:
	;
	goto L45
L51:
	;
	goto L41
L52:
	;
	goto L41
L53:
	;
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_advance[0]))
	F_LWLockRelease(m, v159+int32(_a_F_replorigin_advance_1))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
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
	F_errstart_cold(m, int32(23), int32(0))
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
	v34 = v23 + v26*int32(56)
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
	F_errfinish(m, int32(_a_F_replorigin_redo_1), int32(891), int32(_a_F_replorigin_redo_2))
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
