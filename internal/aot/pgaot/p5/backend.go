package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BackendPidGetProc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	v2 = int32(0)
	if l0 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_BackendPidGetProc[0]))
	v16 = F_LWLockAcquire(m, v12+int32(512), int32(1))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_BackendPidGetProc[1]))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v22 <= int32(0) {
		v51 = v2
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_BackendPidGetProc[0]))
	F_LWLockRelease(m, v56+int32(512))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L4
	} else {
		goto L12
	}
L7:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_BackendPidGetProc[2]))
	v31 = int32(0)
	goto L8
L8:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v21+int32(36)+v31<<(uint(int32(2))%32))))
	v42 = v29 + v39*int32(768)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	if v43 == l0 {
		v51 = v42
		goto L6
	} else {
		goto L10
	}
L9:
	;
	v51 = int32(0)
	goto L6
L10:
	;
	v46 = v31 + int32(1)
	if v46 != v22 {
		v31 = v46
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	return v51
}
