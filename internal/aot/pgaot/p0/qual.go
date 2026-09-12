package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_EvalPlanQualInit(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	v16 = F_palloc0(m, v8<<(uint(int32(2))%32))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		v18 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(l0)+28)) = v18
		*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v16
		*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = v18
		*(*int64)(unsafe.Add(mBase, uint32(l0)+44)) = v18
		return
	}
}
func F_EvalPlanQualSlot(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v10 = v5 + l2<<(uint(int32(2))%32) - int32(4)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if v11 != 0 {
		v29 = v11
		return v29
	} else {
		v12 = int32(4562080)
		v13 = *(*int32)(unsafe.Add(mBase, _consts[0]))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+100))
		*(*int32)(unsafe.Add(mBase, _consts[0])) = v16
		v20 = F_table_slot_create(m, l1, l0+int32(12))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = v20
			*(*int32)(unsafe.Add(mBase, _consts[0])) = v13
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			v29 = v27
			return v29
		}
	}
}
