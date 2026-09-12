package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_emit_jsp_gin_entries(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	F_check_stack_depth(m)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v8 {
	case 0, 1:
		goto L4
	case 2:
		goto L5
	default:
		goto L3
	}
L3:
	;
	return
L4:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v39 <= int32(0) {
		goto L3
	} else {
		goto L16
	}
L5:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v10 < v11 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v30 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v29+v30<<(uint(int32(2))%32)))) = v9
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10
	return
L7:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v29 = v13
	goto L6
L8:
	;
	goto L9
L9:
	;
	if v11 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v27
	v29 = v27
	goto L6
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v11 << (uint(int32(1)) % 32)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v20 = F_repalloc(m, v17, v11<<(uint(int32(3))%32))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(8)
	v25 = F_palloc(m, int32(32))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	v27 = v20
	goto L10
L15:
	;
	v27 = v25
	goto L10
L16:
	;
	v46 = int32(0)
	goto L17
L17:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(8)+v46<<(uint(int32(2))%32))))
	F_emit_jsp_gin_entries(m, v52, l1)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L19
	}
L18:
	;
	goto L3
L19:
	;
	v56 = v46 + int32(1)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v56 < v57 {
		v46 = v56
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
}
func F_execute_jsp_gin_node(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v12 {
	case 0:
		goto L6
	case 1:
		goto L5
	case 2:
		goto L2
	default:
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return base.I32_extend8_s(v90)
L2:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v87 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1+v85))))
	v90 = v87
	goto L1
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L13
	} else {
		goto L22
	}
L4:
	;
	v51 = v3
	v52 = v3
	goto L16
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v17 <= int32(0) {
		v90 = int32(1)
		goto L1
	} else {
		goto L8
	}
L6:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v13 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v90 = v3
	goto L1
L8:
	;
	v26 = v3
	v27 = int32(1)
	goto L9
L9:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(8)+v26<<(uint(int32(2))%32))))
	v34 = F_execute_jsp_gin_node(m, v33, l1)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v90 = v41
	goto L1
L11:
	;
	v43 = v26 + int32(1)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v43 < v44 {
		v26 = v43
		v27 = v41
		goto L9
	} else {
		goto L15
	}
L12:
	;
	v41 = int32(2)
	goto L11
L13:
	;
	return int32(0)
L14:
	;
	switch v34 & int32(255) {
	case 0:
		v90 = v34
		goto L1
	default:
		v41 = v27
		goto L11
	case 2:
		goto L12
	}
L15:
	;
	goto L10
L16:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(8)+v51<<(uint(int32(2))%32))))
	v59 = F_execute_jsp_gin_node(m, v58, l1)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L13
	} else {
		goto L20
	}
L17:
	;
	v90 = v66
	goto L1
L18:
	;
	v68 = v51 + int32(1)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v68 < v69 {
		v51 = v68
		v52 = v66
		goto L16
	} else {
		goto L21
	}
L19:
	;
	v66 = int32(2)
	goto L18
L20:
	;
	switch v59&int32(255) - int32(1) {
	case 0:
		v90 = v59
		goto L1
	case 1:
		goto L19
	default:
		v66 = v52
		goto L18
	}
L21:
	;
	goto L17
L22:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v75
	F_errmsg_internal(m, int32(486080), v10)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L13
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(496890), int32(842), int32(411271))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L13
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
