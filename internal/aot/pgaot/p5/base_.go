package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_add_base_rels_to_query(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l1 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v7 + int32(16)
	return
L2:
	;
	v12 = l1
	goto L4
L3:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v62 = F_build_simple_rel(m, l0, v60, int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L15
	} else {
		goto L23
	}
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v15 != int32(64) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L15
	} else {
		goto L20
	}
L6:
	;
	goto L5
L7:
	;
	switch v15 - int32(63) {
	case 0:
		goto L3
	default:
		goto L6
	case 2:
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	F_add_base_rels_to_query(m, l0, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L15
	} else {
		goto L18
	}
L10:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v20 == int32(0) {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v23 <= int32(0) {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v28 = int32(0)
	goto L13
L13:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31+v28<<(uint(int32(2))%32))))
	F_add_base_rels_to_query(m, l0, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L1
L15:
	;
	return
L16:
	;
	v39 = v28 + int32(1)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v39 < v40 {
		v28 = v39
		goto L13
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	if v45 != 0 {
		v12 = v45
		goto L4
	} else {
		goto L19
	}
L19:
	;
	goto L1
L20:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v50
	F_errmsg_internal(m, int32(506946), v7)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L15
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(519107), int32(185), int32(15900))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L15
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	goto L1
}
func F_find_base_rel(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if base.Ui32(v8) <= base.Ui32(l1) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = l1
			F_errmsg_internal(m, int32(499575), v6)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(522301), int32(426), int32(321291))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10+l1<<(uint(int32(2))%32))))
		if v14 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l1
				F_errmsg_internal(m, int32(499575), v6)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(522301), int32(426), int32(321291))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			m.G0 = v6 + int32(16)
			return v14
		}
	}
}
