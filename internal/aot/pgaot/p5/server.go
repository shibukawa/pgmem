package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CloseServerPorts(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	v3 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, _consts[466]))
	if v3 < v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = v3
	goto L4
L2:
	;
	goto L3
L3:
	;
	v42 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[466])) = v42
	v46 = *(*int32)(unsafe.Add(mBase, _consts[467]))
	if v46 == v42 {
		goto L14
	} else {
		goto L15
	}
L4:
	;
	v11 = *(*int32)(unsafe.Add(mBase, _consts[468]))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v11+v8<<(uint(int32(2))%32))))
	v16 = F_close(m, v15)
	mBase = m.M
	if v16 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v35 = v8 + int32(1)
	v37 = *(*int32)(unsafe.Add(mBase, _consts[466]))
	if v35 < v37 {
		v8 = v35
		goto L4
	} else {
		goto L13
	}
L7:
	;
	v21 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	if v21 == int32(0) {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	F_errmsg_internal(m, int32(307155), int32(0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(519676), int32(1428), int32(125745))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	goto L6
L13:
	;
	goto L5
L14:
	;
	*(*int32)(unsafe.Add(mBase, _consts[467])) = int32(0)
	return
L15:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v49 <= int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v52 = v42
	goto L17
L17:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54+v52<<(uint(int32(2))%32))))
	v59 = F_unlink(m, v58)
	mBase = m.M
	v61 = v52 + int32(1)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v61 < v62 {
		v52 = v61
		goto L17
	} else {
		goto L19
	}
L18:
	;
	goto L14
L19:
	;
	goto L18
}
