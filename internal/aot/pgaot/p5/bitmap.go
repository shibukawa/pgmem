package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecBitmapIndexScanRetrieveInstrumentation(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	if v4 == int32(0) {
		return
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v11 = v7<<(uint(int32(3))%32) + int32(8)
		v12 = F_palloc(m, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v12
			if v11 == int32(0) {
			} else {
				base.MemoryCopy(m, v12, v4, v11)
			}
			return
		}
	}
}
func F_ExecBitmapOr(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	F_errstart_cold(m, int32(21), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		F_errmsg_internal(m, int32(_a_F_ExecBitmapOr_0), int32(0))
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			F_errfinish(m, int32(_a_F_ExecBitmapOr_1), int32(45), int32(_a_F_ExecBitmapOr_2))
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_bitmap_match(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	v4 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.B2i32(v8 == v4)|base.B2i32(v9 == v4) != 0 {
		v51 = base.B2i32(v8|v9 == v4)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return base.B2i32(v51 == int32(0))
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v19 != v20 {
		v51 = int32(0)
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = int32(1)
	if v19 <= v22 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v25 = v22
	goto L6
L5:
	;
	v25 = v19
	goto L6
L6:
	;
	v26 = int32(8)
	v31 = int32(0)
	goto L7
L7:
	;
	v39 = v31 << (uint(int32(2)) % 32)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v8+v26+v39)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v9+v26+v39)))
	v44 = base.B2i32(v41 == v43)
	if v41 != v43 {
		v51 = v44
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v51 = v44
	goto L1
L9:
	;
	v47 = v31 + int32(1)
	if v47 != v25 {
		v31 = v47
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
}
