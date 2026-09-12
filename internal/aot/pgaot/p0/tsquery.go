package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tsquery_requires_match(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	F_check_stack_depth(m)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v12 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L19
	}
L4:
	;
	m.G0 = v6 + int32(16)
	return v49
L5:
	;
	v49 = int32(1)
	goto L4
L6:
	;
	goto L7
L7:
	;
	v16 = l0
	goto L8
L8:
	;
	v19 = int32(0)
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	switch v20 - int32(1) {
	case 0:
		v49 = v19
		goto L4
	case 1, 3:
		goto L12
	case 2:
		goto L11
	default:
		goto L3
	}
L9:
	;
	v49 = v42
	goto L4
L10:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L17
	}
L11:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v36 = F_tsquery_requires_match(m, v16+v32*int32(12))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L15
	}
L12:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v27 = F_tsquery_requires_match(m, v16+v23*int32(12))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	if v27 == int32(0) {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v49 = int32(1)
	goto L4
L15:
	;
	if v36 == int32(0) {
		v49 = v19
		goto L4
	} else {
		goto L16
	}
L16:
	;
	goto L10
L17:
	;
	v42 = int32(1)
	v44 = v16 + int32(12)
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v45 != v42 {
		v16 = v44
		goto L8
	} else {
		goto L18
	}
L18:
	;
	goto L9
L19:
	;
	v59 = int32(*(*int8)(unsafe.Add(mBase, uint32(v16)+1)))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v59
	F_errmsg_internal(m, int32(459611), v6)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(472629), int32(2195), int32(308701))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
