package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tsquery_requires_match(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	F_check_stack_depth(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v13 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L19
	}
L4:
	;
	m.G0 = v7 + int32(16)
	return v51
L5:
	;
	v51 = int32(1)
	goto L4
L6:
	;
	goto L7
L7:
	;
	v17 = l0
	goto L8
L8:
	;
	v21 = int32(0)
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	switch v22 - int32(1) {
	case 0:
		v51 = v21
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
	v51 = v45
	goto L4
L10:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L17
	}
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v38 = F_tsquery_requires_match(m, v17+v34*int32(12))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L15
	}
L12:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v29 = F_tsquery_requires_match(m, v17+v25*int32(12))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	if v29 == int32(0) {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v51 = int32(1)
	goto L4
L15:
	;
	if v38 == int32(0) {
		v51 = v21
		goto L4
	} else {
		goto L16
	}
L16:
	;
	goto L10
L17:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+12)))
	v45 = int32(1)
	if v44 != v45 {
		v17 = v17 + int32(12)
		goto L8
	} else {
		goto L18
	}
L18:
	;
	goto L9
L19:
	;
	v62 = int32(*(*int8)(unsafe.Add(mBase, uint32(v17)+1)))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v62
	F_errmsg_internal(m, int32(_a_F_tsquery_requires_match_0), v7)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(_a_F_tsquery_requires_match_1), int32(2195), int32(_a_F_tsquery_requires_match_2))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
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
