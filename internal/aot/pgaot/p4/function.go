package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_FunctionCall4Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
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
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l0
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v35 = m.T0[v34].(func(*base.Module, int32) int32)(m, v8+int32(-52))
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
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = v46
				F_errmsg_internal(m, int32(532930), v10)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(495561), int32(1217), int32(304296))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
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
func F_record_plan_function_dependency(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	if base.Ui32(int32(12000)) <= base.Ui32(l1) {
		v7 = F_palloc0(m, int32(12))
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v7))) = int64(201863463291)
			v13 = F_GetSysCacheHashValue(m, int32(47), l1, int32(0))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v13
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+60))
				v18 = F_lappend(m, v17, v7)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v20)+60)) = v18
					return
				}
			}
		}
	} else {
		return
	}
}
