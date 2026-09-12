package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bpchar_input(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l2 < int32(4) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v83
L2:
	;
	v65 = v61 + int32(4)
	v66 = F_palloc(m, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L6
	} else {
		goto L25
	}
L3:
	;
	v58 = l1
	v61 = l1
	goto L2
L4:
	;
	goto L5
L5:
	;
	v14 = F_pg_mbstrlen_with_len(m, l0, l1)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	v19 = l2 - int32(4)
	if base.Ui32(v19) < base.Ui32(v14) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v21 = F_pg_mbcharcliplen(m, l0, l1, v19)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v58 = l1
	v61 = l1 + v19 - v14
	goto L2
L11:
	;
	if base.Ui32(l1) <= base.Ui32(v21) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v58 = v21
	v61 = v21
	goto L2
L13:
	;
	goto L14
L14:
	;
	v26 = v21
	goto L16
L15:
	;
	v38 = int32(0)
	v39 = F_errsave_start(m, l3)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L6
	} else {
		goto L20
	}
L16:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v26))))
	if v32 != int32(32) {
		goto L15
	} else {
		goto L18
	}
L17:
	;
	v58 = v21
	v61 = v21
	goto L2
L18:
	;
	v36 = v26 + int32(1)
	if v36 != l1 {
		v26 = v36
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	if v39 == int32(0) {
		v83 = v38
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_errcode(m, int32(16777346))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v19
	F_errmsg(m, int32(678840), v10)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	F_errsave_finish(m, l3, int32(495859), int32(162), int32(64639))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L6
	} else {
		goto L24
	}
L24:
	;
	v83 = v38
	goto L1
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v65 << (uint(int32(2)) % 32)
	v72 = v66 + int32(4)
	if v58 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	if base.Ui32(v61) <= base.Ui32(v58) {
		v83 = v66
		goto L1
	} else {
		goto L30
	}
L27:
	;
	v73 = F__emscripten_memcpy_bulkmem(m, v72, l0, v58)
	mBase = m.M
	v74 = v73
	goto L29
L28:
	;
	v74 = v72
	goto L29
L29:
	;
	goto L26
L30:
	;
	v80 = F__emscripten_memset_bulkmem(m, v74+v58, base.I32_extend8_s(int32(32)), v61-v58)
	mBase = m.M
	goto L31
L31:
	;
	v83 = v66
	goto L1
}
