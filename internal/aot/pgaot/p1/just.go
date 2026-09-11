package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecJustAssignOuterVarVirt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	v9 = int32(2)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v13+v14<<(uint(v9)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v6+v8<<(uint(v9)%32)))) = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v5)+20))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+v22))))
	*(*uint8)(unsafe.Add(mBase, uint32(v8+v20))) = uint8(v24)
	return int32(0)
}
func F_ExecJustAssignScanVarVirt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	v9 = int32(2)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v13+v14<<(uint(v9)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v6+v8<<(uint(v9)%32)))) = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v5)+20))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+v22))))
	*(*uint8)(unsafe.Add(mBase, uint32(v8+v20))) = uint8(v24)
	return int32(0)
}
func F_ExecJustOuterVar(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+56))
	v9 = v7 + int32(1)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v11 = int32(*(*int16)(unsafe.Add(mBase, uint32(v10)+6)))
	if v11 < v9 {
		F_slot_getsomeattrs_int(m, v10, v9)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+v7))))
			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v19)
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v21+v7<<(uint(int32(2))%32))))
			return v25
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+v7))))
		*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v19)
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v21+v7<<(uint(int32(2))%32))))
		return v25
	}
}
