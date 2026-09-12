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
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	if v4 != 0 {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v9 = v5<<(uint(int32(3))%32) + int32(8)
		v10 = F_palloc(m, v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v10
			if v9 != 0 {
				v13 = F__emscripten_memcpy_bulkmem(m, v10, v4, v9)
				mBase = m.M
			} else {
			}
			return
		}
	} else {
		return
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
		F_errmsg_internal(m, int32(242632), int32(0))
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			F_errfinish(m, int32(484474), int32(45), int32(225768))
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
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	v4 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v12 = base.B2i32(v8|v9 == v4)
	if v8 == v4 {
		v51 = v12
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return base.B2i32(v51 == int32(0))
L2:
	;
	if v9 == int32(0) {
		v51 = v12
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v18 != v19 {
		v51 = int32(0)
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v21 = int32(1)
	if v18 <= v21 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v24 = v21
	goto L7
L6:
	;
	v24 = v18
	goto L7
L7:
	;
	v25 = int32(8)
	v30 = int32(0)
	goto L8
L8:
	;
	v38 = v30 << (uint(int32(2)) % 32)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v8+v25+v38)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38+(v9+v25))))
	v43 = base.B2i32(v40 == v42)
	if v42 != v40 {
		v51 = v43
		goto L1
	} else {
		goto L10
	}
L9:
	;
	v51 = v43
	goto L1
L10:
	;
	v46 = v30 + int32(1)
	if v46 != v24 {
		v30 = v46
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
}
