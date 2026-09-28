package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DirectFunctionCall2Coll(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64) int64 {
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
	var v27 int64
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
	v8 = v6 + int32(-64)
	m.G0 = v8
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+56)) = uint8(v5)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+48)) = l3
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+40)) = uint8(v5)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = l2
	v16 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+26)) = uint16(v16)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+24)) = uint8(v5)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v5
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(0)
	v27 = m.T0[l0].(func(*base.Module, int32) int64)(m, v6+int32(-56))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		return int64(0)
	} else {
		v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+24)))
		if v31 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
				F_errmsg_internal(m, int32(_a_F_DirectFunctionCall2Coll_0), v8)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_DirectFunctionCall2Coll_1), int32(830), int32(_a_F_DirectFunctionCall2Coll_2))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			m.G0 = v8 - int32(-64)
			return v27
		}
	}
}
func F_DirectFunctionCall4Coll(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int64, l5 int64) int64 {
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
	var v35 int64
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
	v10 = v8 - int32(96)
	m.G0 = v10
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+88)) = uint8(v7)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+80)) = l5
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+72)) = uint8(v7)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+64)) = l4
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+56)) = uint8(v7)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = l3
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+40)) = uint8(v7)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = l2
	v24 = int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+26)) = uint16(v24)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)) = uint8(v7)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v7
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = int64(0)
	v35 = m.T0[l0].(func(*base.Module, int32) int64)(m, v10+int32(8))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		return int64(0)
	} else {
		v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
		if v39 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
				F_errmsg_internal(m, int32(_a_F_DirectFunctionCall4Coll_0), v10)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_DirectFunctionCall4Coll_1), int32(882), int32(_a_F_DirectFunctionCall4Coll_2))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			m.G0 = v10 + int32(96)
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
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v19 int32
	_ = v19
	v2 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+95)) = uint8(v2)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v6 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+28)))
	v7 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0)+76)))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v9 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
	v10 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+72)))
	v14 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+64)))
	v15 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+68)))
	v16 = F_FunctionCall8Coll(m, v4, v5, v6, v7, v8, v9, v10, base.I64_extend_i32_u(l0+int32(95)), v14, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v16 != int64(0))
	}
}
