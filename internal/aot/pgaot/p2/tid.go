package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_TidStoreGetHandle(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	return v4
}
func F_TidStoreMemoryUsage(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v4 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v10 = F_LWLockAcquire(m, v6+int32(1476), int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v3)+24))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	if v26 != 0 {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	return int32(0)
L5:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+1448))
	F_LWLockRelease(m, v14+int32(1476))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	return v15
L7:
	;
	return v43
L8:
	;
	v29 = v25
	v30 = v26
	goto L11
L9:
	;
	v43 = v25
	goto L10
L10:
	;
	goto L7
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	v32 = v31 + v29
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	if v33 != 0 {
		v29 = v32
		v30 = v33
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v43 = v32
	goto L10
L13:
	;
	v35 = v30
	goto L14
L14:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+28))
	if v38 != 0 {
		v29 = v32
		v30 = v38
		goto L11
	} else {
		goto L16
	}
L15:
	;
	goto L12
L16:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
	if v39 != v21 {
		v35 = v39
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
}
