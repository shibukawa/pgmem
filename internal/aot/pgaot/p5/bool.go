package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_readBoolCols(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_pg_strtok(m, v8+int32(12))
	mBase = m.M
	if v12 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L9
	} else {
		goto L25
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L9
	} else {
		goto L22
	}
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	switch v13 {
	case 0:
		v64 = int32(0)
		goto L6
	case 1:
		goto L7
	default:
		goto L2
	}
L4:
	;
	goto L5
L5:
	;
	goto L1
L6:
	;
	m.G0 = v8 + int32(16)
	return v64
L7:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v14 != int32(40) {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v18 = F_palloc(m, l0)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	if int32(0) < l0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v26 = int32(0)
	goto L14
L12:
	;
	goto L13
L13:
	;
	v51 = F_pg_strtok(m, v8+int32(12))
	mBase = m.M
	if v51 == int32(0) {
		goto L1
	} else {
		goto L19
	}
L14:
	;
	v31 = F_pg_strtok(m, v8+int32(12))
	mBase = m.M
	if v31 == int32(0) {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	goto L13
L16:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if v34 == int32(41) {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v26+v18))) = uint8(base.B2i32(v34 == int32(116)))
	v42 = v26 + int32(1)
	if v42 != l0 {
		v26 = v42
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v54 != int32(1) {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v57 != int32(41) {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v64 = v18
	goto L6
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v12
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v74
	F_errmsg_internal(m, int32(_a_F_readBoolCols_0), v8)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L9
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_readBoolCols_1), int32(697), int32(_a_F_readBoolCols_2))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L9
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	F_errmsg_internal(m, int32(_a_F_readBoolCols_3), int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L9
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_readBoolCols_1), int32(697), int32(_a_F_readBoolCols_2))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L9
	} else {
		goto L27
	}
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
