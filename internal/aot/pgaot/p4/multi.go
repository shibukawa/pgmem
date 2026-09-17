package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_multi_sort_compare(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v85 int32
	_ = v85
	v4 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v4 < v9 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return int32(1)
L2:
	;
	return v85
L3:
	;
	v17 = v9
	v18 = v4
	goto L6
L4:
	;
	goto L5
L5:
	;
	v85 = int32(0)
	goto L2
L6:
	;
	v24 = l2 + int32(4) + v18*int32(36)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+v18))))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+v18))))
	if v30 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L5
L8:
	;
	v71 = v18 + int32(1)
	if v71 < v69 {
		v17 = v69
		v18 = v71
		goto L6
	} else {
		goto L29
	}
L9:
	;
	if v27&int32(1) != 0 {
		v69 = v17
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	if v27&int32(1) != 0 {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+9)))
	if v37 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v38 = int32(-1)
	goto L15
L14:
	;
	v38 = int32(1)
	goto L15
L15:
	;
	return v38
L16:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+9)))
	if v44 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v48 = v18 << (uint(int32(2)) % 32)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v48+v49)))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v52+v48)))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v56 = m.T0[v55].(func(*base.Module, int32, int32, int32) int32)(m, v51, v54, v24)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v45 = int32(1)
	goto L21
L20:
	;
	v45 = int32(-1)
	goto L21
L21:
	;
	return v45
L22:
	;
	return int32(0)
L23:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+8)))
	if v60 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	if v56 < int32(0) {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	v67 = v56
	goto L26
L26:
	;
	if v67 != 0 {
		v85 = v67
		goto L2
	} else {
		goto L28
	}
L27:
	;
	v67 = int32(0) - v56
	goto L26
L28:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v69 = v68
	goto L8
L29:
	;
	goto L7
}
