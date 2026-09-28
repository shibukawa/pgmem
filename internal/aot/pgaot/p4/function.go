package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_FunctionCall0Coll(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	*(*uint16)(unsafe.Add(mBase, uint32(v6)+26)) = uint16(v2)
	*(*uint8)(unsafe.Add(mBase, uint32(v6)+24)) = uint8(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v2
	*(*int64)(unsafe.Add(mBase, uint32(v6)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = l0
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v20 = m.T0[v19].(func(*base.Module, int32) int64)(m, v6+int32(8))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return int64(0)
	} else {
		v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+24)))
		if v24 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int64(0)
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v31
				F_errmsg_internal(m, int32(_a_F_FunctionCall0Coll_0), v6)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_FunctionCall0Coll_1), int32(1125), int32(_a_F_FunctionCall0Coll_2))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			m.G0 = v6 + int32(32)
			return v20
		}
	}
}
func F_FunctionCall4Coll(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int64, l5 int64) int64 {
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
	var v35 int64
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
	*(*int64)(unsafe.Add(mBase, uint32(v10)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = l0
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v35 = m.T0[v34].(func(*base.Module, int32) int64)(m, v10+int32(8))
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
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = v46
				F_errmsg_internal(m, int32(_a_F_FunctionCall4Coll_0), v10)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_FunctionCall4Coll_1), int32(1219), int32(_a_F_FunctionCall4Coll_2))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
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
func F_record_plan_function_dependency(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	if base.Ui32(int32(_a_F_record_plan_function_dependency_0)) <= base.Ui32(l1) {
		v7 = F_palloc0(m, int32(12))
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v7))) = int64(201863463295)
			v14 = F_GetSysCacheHashValue(m, int32(47), base.I64_extend_i32_u(l1), int64(0))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v14
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+68))
				v19 = F_lappend(m, v18, v7)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = v19
					return
				}
			}
		}
	} else {
		return
	}
}
