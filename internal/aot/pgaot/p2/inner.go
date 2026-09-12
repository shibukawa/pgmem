package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_inner_int_contains(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = F_ArrayGetNItems(m, v10, l0+int32(16))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v20 = F_ArrayGetNItems(m, v17, l1+int32(16))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v22 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v32 = (v25<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L6
L5:
	;
	v32 = v22
	goto L6
L6:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v33 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v43 = (v36<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L9
L8:
	;
	v43 = v33
	goto L9
L9:
	;
	if v13 <= int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return base.B2i32(v20 == int32(0))
L11:
	;
	goto L12
L12:
	;
	if v20 <= int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return base.B2i32(v20 == int32(0))
L14:
	;
	goto L15
L15:
	;
	v56 = int32(0)
	v59 = v56
	v60 = v56
	v63 = int32(0)
	goto L16
L16:
	;
	v67 = int32(2)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0+v32+v63<<(uint(v67)%32))))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1+v43+v59<<(uint(v67)%32))))
	if v74 <= v70 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	return base.B2i32(v88 == v20)
L18:
	;
	goto L17
L19:
	;
	if v70 != v74 {
		v88 = v60
		goto L18
	} else {
		goto L22
	}
L20:
	;
	v81 = v59
	v82 = v60
	goto L21
L21:
	;
	v84 = v63 + int32(1)
	if v13 <= v84 {
		v88 = v82
		goto L18
	} else {
		goto L23
	}
L22:
	;
	v77 = int32(1)
	v81 = v59 + v77
	v82 = v60 + v77
	goto L21
L23:
	;
	if v81 < v20 {
		v59 = v81
		v60 = v82
		v63 = v84
		goto L16
	} else {
		goto L24
	}
L24:
	;
	v88 = v82
	goto L18
}
