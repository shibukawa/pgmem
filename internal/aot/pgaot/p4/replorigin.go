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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v117 int64
	_ = v117
	var v118 int32
	_ = v118
	var v119 int64
	_ = v119
	var v126 int64
	_ = v126
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
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
	if v27 <= int32(0) {
		v88 = v6
		v90 = v6
		goto L6
	} else {
		goto L7
	}
L6:
	;
	if v88|v90 != 0 {
		goto L25
	} else {
		goto L26
	}
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_advance[2]))
	v39 = v6
	v40 = v6
	goto L8
L8:
	;
	v46 = v31 + v40*int32(56)
	v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46))))
	if v47 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v88 = int32(0)
	v90 = v78
	goto L6
L10:
	;
	v80 = v40 + int32(1)
	if v80 != v27 {
		v39 = v78
		v40 = v80
		goto L8
	} else {
		goto L21
	}
L11:
	;
	if v1 != v47 {
		v78 = v39
		goto L10
	} else {
		goto L14
	}
L12:
	;
	if v39 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v78 = v46
	goto L10
L14:
	;
	v52 = F_LWLockAcquire(m, v46+int32(40), int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v46)+24))
	if v54 == int32(0) {
		v88 = v46
		v90 = v39
		goto L6
	} else {
		goto L16
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	F_errcode(m, int32(100663621))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46))))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v46)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v64
	F_errmsg(m, int32(_a_F_replorigin_advance_2), v15+int32(16))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_replorigin_advance_3), int32(969), int32(_a_F_replorigin_advance_4))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L21:
	;
	goto L9
L22:
	;
	F_LWLockRelease(m, v104+int32(40))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L4
	} else {
		goto L52
	}
L23:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v104)+16)) = l2
	goto L22
L24:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v104)+8)) = l1
	if l2 == int64(0) {
		goto L22
	} else {
		goto L51
	}
L25:
	;
	if v88 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L4
	} else {
		goto L46
	}
L28:
	;
	v101 = F_LWLockAcquire(m, v90+int32(40), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L31
	}
L29:
	;
	v104 = v88
	goto L30
L30:
	;
	if l4 != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v90))) = uint16(v1)
	v104 = v90
	goto L30
L32:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+42)) = uint8(v4)
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+40)) = uint16(v1)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+32)) = l1
	F_XLogBeginInsert(m)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L4
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	if v4 != 0 {
		goto L24
	} else {
		goto L38
	}
L35:
	;
	F_XLogRegisterData(m, v15+int32(32), int32(16))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	v117 = F_XLogInsert(m, int32(19), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	goto L34
L38:
	;
	v119 = *(*int64)(unsafe.Add(mBase, uint32(v104)+8))
	if base.Ui64(v119) < base.Ui64(l1) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v126 = *(*int64)(unsafe.Add(mBase, uint32(v104)+16))
	if base.Ui64(l2) <= base.Ui64(v126) {
		goto L22
	} else {
		goto L45
	}
L40:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v104)+8)) = l1
	if l2 != int64(0) {
		goto L39
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	if l2 == int64(0) {
		goto L22
	} else {
		goto L44
	}
L43:
	;
	goto L22
L44:
	;
	goto L39
L45:
	;
	goto L23
L46:
	;
	F_errcode(m, int32(_a_F_replorigin_advance_5))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v1
	F_errmsg(m, int32(_a_F_replorigin_advance_6), v15)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	F_errhint(m, int32(_a_F_replorigin_advance_7), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_replorigin_advance_3), int32(980), int32(_a_F_replorigin_advance_4))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	goto L23
L52:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_advance[0]))
	F_LWLockRelease(m, v157+int32(_a_F_replorigin_advance_1))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int64
	_ = v38
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v60 int64
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+48)))
	v14 = v12 & int32(240)
	switch v14 {
	case 0:
		goto L2
	default:
		goto L3
	case 16:
		goto L4
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return
L2:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+8)))
	v59 = *(*int64)(unsafe.Add(mBase, uint32(v57)))
	v60 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+10)))
	F_replorigin_advance(m, v58, v59, v60, v61, int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L12
	} else {
		goto L16
	}
L3:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L12
	} else {
		goto L13
	}
L4:
	;
	v15 = int32(0)
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_redo[0]))
	if v17 <= v15 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_replorigin_redo[1]))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22))))
	v24 = v15
	goto L6
L6:
	;
	v32 = v21 + v24*int32(56)
	v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32))))
	if v23 != v33 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v38 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+8)) = v38
	v40 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v32))) = uint16(v40)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+16)) = v38
	goto L1
L8:
	;
	v36 = v24 + int32(1)
	if v17 != v36 {
		v24 = v36
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	goto L7
L11:
	;
	goto L1
L12:
	;
	return
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v14
	F_errmsg_internal(m, int32(_a_F_replorigin_redo_0), v9)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(_a_F_replorigin_redo_1), int32(891), int32(_a_F_replorigin_redo_2))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L16:
	;
	goto L1
}
