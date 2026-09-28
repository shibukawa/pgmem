package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CloseServerPorts(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_CloseServerPorts[0]))
	if v3 < v6 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v9 = v3
	goto L4
L2:
	;
	goto L3
L3:
	;
	v45 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_CloseServerPorts[0])) = v45
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_CloseServerPorts[1]))
	if v49 == v45 {
		goto L14
	} else {
		goto L15
	}
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_CloseServerPorts[2]))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13+v9<<(uint(int32(2))%32))))
	v18 = F_close(m, v17)
	mBase = m.M
	if v18 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v37 = v9 + int32(1)
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_CloseServerPorts[0]))
	if v37 < v39 {
		v9 = v37
		goto L4
	} else {
		goto L13
	}
L7:
	;
	v23 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	if v23 == int32(0) {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	F_errmsg_internal(m, int32(_a_F_CloseServerPorts_0), int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(_a_F_CloseServerPorts_1), int32(1442), int32(_a_F_CloseServerPorts_2))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
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
	return
L15:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v52 <= int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v55 = v45
	goto L17
L17:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58+v55<<(uint(int32(2))%32))))
	v63 = F_unlink(m, v62)
	mBase = m.M
	v65 = v55 + int32(1)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v65 < v66 {
		v55 = v65
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
