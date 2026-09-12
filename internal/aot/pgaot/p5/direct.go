package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DirectFunctionCall2Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	v5 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+44)) = uint8(v5)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = l3
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+36)) = uint8(v5)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = l2
	v16 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+30)) = uint16(v16)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+28)) = uint8(v5)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v5
	*(*int64)(unsafe.Add(mBase, uint32(v8)+12)) = int64(0)
	v27 = m.T0[l0].(func(*base.Module, int32) int32)(m, v8+int32(12))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		return int32(0)
	} else {
		v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+28)))
		if v31 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
				F_errmsg_internal(m, int32(557280), v8)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(518227), int32(828), int32(318005))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			m.G0 = v8 + int32(48)
			return v27
		}
	}
}
func F_DirectFunctionCall4Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v24 int32
	_ = v24
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	v7 = int32(0)
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+60)) = uint8(v7)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = l5
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+52)) = uint8(v7)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = l4
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+44)) = uint8(v7)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = l3
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+36)) = uint8(v7)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = l2
	v24 = int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+30)) = uint16(v24)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+28)) = uint8(v7)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v7
	*(*int64)(unsafe.Add(mBase, uint32(v10)+12)) = int64(0)
	v35 = m.T0[l0].(func(*base.Module, int32) int32)(m, v8+int32(-52))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		return int32(0)
	} else {
		v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+28)))
		if v39 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
				F_errmsg_internal(m, int32(557280), v10)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(518227), int32(880), int32(317957))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			m.G0 = v10 - int32(-64)
			return v35
		}
	}
}
func F_directBoolConsistentFn(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	v2 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+87)) = uint8(v2)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+68)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v15 = F_FunctionCall8Coll(m, v4, v5, v6, v7, v8, v9, v10, l0+int32(87), v13, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v15 != int32(0))
	}
}
