package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SN_close_env(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	if l0 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v4 != 0 {
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
	if int32(0) < l1 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	F_pfree(m, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L12
	} else {
		goto L16
	}
L7:
	;
	v10 = int32(0)
	goto L10
L8:
	;
	v25 = v4
	goto L9
L9:
	;
	F_pfree(m, v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L12
	} else {
		goto L15
	}
L10:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v11+v10<<(uint(int32(2))%32))))
	F_lose_s(m, v15)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v25 = v21
	goto L9
L12:
	;
	return
L13:
	;
	v19 = v10 + int32(1)
	if v19 != l1 {
		v10 = v19
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	goto L6
L16:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v34 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	F_lose_s(m, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L12
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
	v38 = m.ExcPending
	if v38 != 0 {
		goto L12
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	goto L3
}
