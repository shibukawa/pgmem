package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_vac_close_indexes(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l0 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v4 = l0
	goto L7
L5:
	;
	goto L6
L6:
	;
	F_pfree(m, l1)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L9
	} else {
		goto L12
	}
L7:
	;
	v8 = v4 - int32(1)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1+v8<<(uint(int32(2))%32))))
	F_relation_close(m, v12, l2)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	return
L10:
	;
	if v8 != 0 {
		v4 = v8
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	goto L3
}
