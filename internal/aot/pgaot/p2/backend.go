package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BackendPidGetProcWithLock(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	v2 = int32(0)
	if l0 == v2 {
		v40 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v40
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_BackendPidGetProcWithLock[0]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if v11 <= int32(0) {
		v40 = v2
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_BackendPidGetProcWithLock[1]))
	v20 = int32(0)
	goto L4
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(36)+v20<<(uint(int32(2))%32))))
	v31 = v18 + v28*int32(768)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	if v32 == l0 {
		v40 = v31
		goto L1
	} else {
		goto L6
	}
L5:
	;
	v40 = int32(0)
	goto L1
L6:
	;
	v35 = v20 + int32(1)
	if v35 != v11 {
		v20 = v35
		goto L4
	} else {
		goto L7
	}
L7:
	;
	goto L5
}
