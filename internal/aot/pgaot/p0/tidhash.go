package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tidhash_destroy(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_pfree(m, v2)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		F_pfree(m, l0)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			return
		}
	}
}
func F_tidhash_iterate(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	v10 = v7
	goto L1
L1:
	;
	if v10&int32(1) != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	return v27
L3:
	;
	return int32(0)
L4:
	;
	goto L5
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v23 = v19 & (v20 - int32(1))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v23
	v27 = v18 + v20<<(uint(int32(3))%32)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v29 = v19 & v28
	if v29 == v23 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v31 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)) = uint8(v31)
	goto L8
L7:
	;
	goto L8
L8:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+6)))
	if v34 != int32(1) {
		v10 = base.B2i32(v23 == v29)
		goto L1
	} else {
		goto L9
	}
L9:
	;
	goto L2
}
