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
	var v63 int32
	_ = v63
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_pg_strtok(m, v8+int32(12))
	mBase = m.M
	if v12 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L10
	} else {
		goto L32
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L10
	} else {
		goto L29
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L10
	} else {
		goto L26
	}
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	switch v13 {
	case 0:
		v63 = int32(0)
		goto L7
	case 1:
		goto L8
	default:
		goto L3
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L10
	} else {
		goto L23
	}
L7:
	;
	m.G0 = v8 + int32(16)
	return v63
L8:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v14 != int32(40) {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v18 = F_palloc(m, l0)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	if int32(0) < l0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v26 = int32(0)
	goto L15
L13:
	;
	goto L14
L14:
	;
	v51 = F_pg_strtok(m, v8+int32(12))
	mBase = m.M
	if v51 == int32(0) {
		goto L2
	} else {
		goto L20
	}
L15:
	;
	v31 = F_pg_strtok(m, v8+int32(12))
	mBase = m.M
	if v31 == int32(0) {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	goto L14
L17:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if v34 == int32(41) {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v26+v18))) = uint8(base.B2i32(v34 == int32(116)))
	v42 = v26 + int32(1)
	if v42 != l0 {
		v26 = v42
		goto L15
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v54 != int32(1) {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v57 != int32(41) {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v63 = v18
	goto L7
L23:
	;
	F_errmsg_internal(m, int32(23721), int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L10
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(478148), int32(697), int32(144812))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L10
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
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v12
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v87
	F_errmsg_internal(m, int32(655739), v8)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L10
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(478148), int32(697), int32(144812))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L29:
	;
	F_errmsg_internal(m, int32(23721), int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L10
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(478148), int32(697), int32(144812))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L10
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	F_errmsg_internal(m, int32(23721), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L10
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(478148), int32(697), int32(144812))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L10
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
