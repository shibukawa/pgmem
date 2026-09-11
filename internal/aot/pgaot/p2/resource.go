package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ResourceOwnerDelete(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	goto L1
L1:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v7 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v10 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	F_ResourceOwnerDelete(m, v7)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	goto L2
L6:
	;
	return
L7:
	;
	goto L1
L8:
	;
	v29 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v29
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	if v33 != 0 {
		goto L17
	} else {
		goto L18
	}
L9:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v13 == l0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v15
	goto L8
L11:
	;
	goto L12
L12:
	;
	v18 = v13
	goto L13
L13:
	;
	if v18 == int32(0) {
		goto L8
	} else {
		goto L15
	}
L14:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v24
	goto L8
L15:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if l0 != v22 {
		v18 = v22
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	F_pfree(m, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L6
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	F_pfree(m, l0)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L6
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	return
}
func F_resource_priority_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v6 == v8 {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
		return base.B2i32(base.Ui32(v11) < base.Ui32(v10)) - base.B2i32(base.Ui32(v10) < base.Ui32(v11))
	} else {
		if base.Ui32(v8) < base.Ui32(v6) {
			v19 = int32(-1)
		} else {
			v19 = int32(1)
		}
		return v19
	}
}
