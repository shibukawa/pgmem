package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BumpGetChunkSpace(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(414910)
		F_errmsg_internal(m, int32(207886), v4)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			F_errfinish(m, int32(490425), int32(651), int32(414949))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_BumpIsEmpty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	v6 = int32(1)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v7 == int32(0) {
		v27 = v6
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v27
L2:
	;
	v11 = l0 + int32(60)
	if v7 == v11 {
		v27 = v6
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v14 = v7
	goto L4
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v20 = v14 + int32(16)
	v21 = base.B2i32(v18 == v20)
	if v18 != v20 {
		v27 = v21
		goto L1
	} else {
		goto L6
	}
L5:
	;
	v27 = v21
	goto L1
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v23 != v11 {
		v14 = v23
		goto L4
	} else {
		goto L7
	}
L7:
	;
	goto L5
}
