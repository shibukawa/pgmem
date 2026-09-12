package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cstring_to_text(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	if l0&int32(3) == int32(0) {
		v28 = l0
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v63 = v61 + int32(4)
	v64 = F_palloc(m, v63)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L18
	} else {
		goto L19
	}
L2:
	;
	v61 = v53 - l0
	goto L1
L3:
	;
	v32 = v28
	goto L12
L4:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v12 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v61 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v17 = l0
	goto L8
L8:
	;
	v21 = v17 + int32(1)
	if v21&int32(3) == int32(0) {
		v28 = v21
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v53 = v21
	goto L2
L10:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v26 != 0 {
		v17 = v21
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v41 = int32(-2139062144)
	if (int32(16843008)-v38|v38)&v41 == v41 {
		v32 = v32 + int32(4)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v47 = v32
	goto L15
L14:
	;
	goto L13
L15:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	if v51 != 0 {
		v47 = v47 + int32(1)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v53 = v47
	goto L2
L17:
	;
	goto L16
L18:
	;
	return int32(0)
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v63 << (uint(int32(2)) % 32)
	if v61 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	return v64
L21:
	;
	v73 = F__emscripten_memcpy_bulkmem(m, v64+int32(4), l0, v61)
	mBase = m.M
	goto L23
L22:
	;
	goto L23
L23:
	;
	goto L20
}
