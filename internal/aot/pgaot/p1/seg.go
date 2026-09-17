package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_seg_contains(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 float32
	_ = v4
	var v5 int32
	_ = v5
	var v6 float32
	_ = v6
	var v8 float32
	_ = v8
	var v9 float32
	_ = v9
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*float32)(unsafe.Add(mBase, uint32(v3)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*float32)(unsafe.Add(mBase, uint32(v5)))
	if base.F32_le(v4, v6) != 0 {
		v8 = *(*float32)(unsafe.Add(mBase, uint32(v3)+4))
		v9 = *(*float32)(unsafe.Add(mBase, uint32(v5)+4))
		v12 = base.F32_ge(v8, v9)
	} else {
		v12 = int32(0)
	}
	return v12
}
func F_seg_different(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_DirectFunctionCall2Coll(m, int32(_a_F_seg_different_0), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_seg_gt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_DirectFunctionCall2Coll(m, int32(_a_F_seg_gt_0), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(int32(0) < v6)
	}
}
func F_seg_scanner_init(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int64
	_ = v33
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	if l1 != 0 {
		v6 = F_palloc(m, int32(96))
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v6
			if v6 != 0 {
				v27 = int32(0)
				base.MemoryFill(m, v6, v27, int32(96))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				*(*int32)(unsafe.Add(mBase, uint32(v30)+60)) = v27
				v33 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v30)+52)) = v33
				*(*int32)(unsafe.Add(mBase, uint32(v30)+44)) = v27
				*(*int64)(unsafe.Add(mBase, uint32(v30)+36)) = v33
				*(*int64)(unsafe.Add(mBase, uint32(v30)+4)) = v33
				*(*int64)(unsafe.Add(mBase, uint32(v30)+12)) = v33
				*(*int32)(unsafe.Add(mBase, uint32(v30)+20)) = v27
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v46 = F_seg_yy_scan_string(m, l0, v45)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return
				} else {
					return
				}
			} else {
				v12 = int32(48)
				*(*int32)(unsafe.Add(mBase, _c_F_seg_scanner_init[0])) = v12
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(_a_F_seg_scanner_init_0), int32(0))
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_seg_scanner_init_1), int32(104), int32(_a_F_seg_scanner_init_2))
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	} else {
		v12 = int32(28)
		*(*int32)(unsafe.Add(mBase, _c_F_seg_scanner_init[0])) = v12
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_seg_scanner_init_0), int32(0))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_seg_scanner_init_1), int32(104), int32(_a_F_seg_scanner_init_2))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_seg_size(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 float32
	_ = v3
	var v4 float32
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*float32)(unsafe.Add(mBase, uint32(v2)+4))
	v4 = *(*float32)(unsafe.Add(mBase, uint32(v2)))
	return base.I32_reinterpret_f32(base.F32_sub(v3, v4)) & int32(2147483647)
}
