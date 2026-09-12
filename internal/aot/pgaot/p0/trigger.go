package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CreateTrigger(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	v10 = int32(0)
	F_CreateTriggerFiringOn(m, l0, l1, l2, l3, l4, l5, l6, v10, l7, v10, l8, v10, int32(79))
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		return
	}
}
func F_ExecGetTriggerOldSlot(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v5 == int32(0) {
		v8 = int32(4442992)
		v9 = *(*int32)(unsafe.Add(mBase, _consts[0]))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
		*(*int32)(unsafe.Add(mBase, _consts[0])) = v12
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+52))
		v15 = F_table_slot_callbacks(m, v10)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = F_ExecInitExtraTupleSlot(m, l0, v14, v15)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = v19
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v9
				v24 = v19
				return v24
			}
		}
	} else {
		v24 = v5
		return v24
	}
}
func F_FindTriggerIncompatibleWithInheritance(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	v2 = int32(0)
	if l0 == v2 {
		v31 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v31
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v7 <= int32(0) {
		v31 = v2
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = int32(0)
	goto L5
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v31 = v29
	goto L1
L5:
	;
	v18 = v10 + v12*int32(60)
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+12)))
	if v19&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	return int32(0)
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
	if v22 != 0 {
		goto L4
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v25 = v12 + int32(1)
	if v25 != v7 {
		v12 = v25
		goto L5
	} else {
		goto L12
	}
L10:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
	if v23 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	goto L6
}
