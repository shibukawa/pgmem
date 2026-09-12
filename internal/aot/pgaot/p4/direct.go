package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DirectFunctionCall1Coll(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	v4 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+28)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = l2
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+20)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v4
	*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = int64(0)
	v19 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+22)) = uint16(v19)
	v23 = m.T0[l0].(func(*base.Module, int32) int32)(m, v7+int32(4))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		return int32(0)
	} else {
		v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+20)))
		if v27 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
				F_errmsg_internal(m, int32(524390), v7)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(487551), int32(806), int32(299654))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			m.G0 = v7 + int32(32)
			return v23
		}
	}
}
func F_DirectFunctionCall3Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v20 int32
	_ = v20
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	v6 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+44)) = uint8(v6)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = l4
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+36)) = uint8(v6)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = l3
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)) = uint8(v6)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = l2
	v20 = int32(3)
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+22)) = uint16(v20)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)) = uint8(v6)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v6
	*(*int64)(unsafe.Add(mBase, uint32(v9)+4)) = int64(0)
	v31 = m.T0[l0].(func(*base.Module, int32) int32)(m, v9+int32(4))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		return int32(0)
	} else {
		v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
		if v35 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
				F_errmsg_internal(m, int32(524390), v9)
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(487551), int32(853), int32(299606))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			m.G0 = v9 + int32(48)
			return v31
		}
	}
}
func F_directTriConsistentFn(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
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
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+68)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v11 = F_FunctionCall7Coll(m, v2, v3, v4, v5, v6, v7, v8, v9, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		return base.I32_extend8_s(v11)
	}
}
