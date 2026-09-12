package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_errdetail_log_plural(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = int32(4508044)
	v14 = *(*int32)(unsafe.Add(mBase, _consts[1408]))
	*(*int32)(unsafe.Add(mBase, _consts[1408])) = v14 + int32(1)
	v19 = *(*int32)(unsafe.Add(mBase, _consts[1409]))
	if int32(0) <= v19 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = int32(4515120)
	v23 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v26 = v19 * int32(100)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_consts[1410])))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v29
	F_initStringInfo(m, v10+int32(16))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1409])) = int32(-1)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L4
	} else {
		goto L24
	}
L4:
	;
	return
L5:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_consts[1412])))
	*(*int32)(unsafe.Add(mBase, _consts[140])) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l3
	if l2 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v43 = l0
	goto L8
L7:
	;
	v43 = l1
	goto L8
L8:
	;
	v44 = F_appendStringInfoVA(m, v10+int32(16), v43, l3)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	if v44 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v48 = v44
	goto L13
L11:
	;
	goto L12
L12:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_consts[1421])))
	if v72 != 0 {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	F_enlargeStringInfo(m, v10+int32(16), v48)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L4
	} else {
		goto L15
	}
L14:
	;
	goto L12
L15:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_consts[1412])))
	*(*int32)(unsafe.Add(mBase, _consts[140])) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l3
	v63 = F_appendStringInfoVA(m, v10+int32(16), v43, l3)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	if v63 != 0 {
		v48 = v63
		goto L13
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	F_pfree(m, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L4
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v76 = F_pstrdup(m, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_consts[1421]))) = v76
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	F_pfree(m, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v23
	v84 = int32(4508044)
	v86 = *(*int32)(unsafe.Add(mBase, _consts[1408]))
	*(*int32)(unsafe.Add(mBase, _consts[1408])) = v86 - int32(1)
	m.G0 = v10 + int32(32)
	return
L24:
	;
	F_errmsg_internal(m, int32(453713), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(497857), int32(1280), int32(310053))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
