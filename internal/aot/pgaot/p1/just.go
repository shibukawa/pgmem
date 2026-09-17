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
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+56))
	v8 = v6 + int32(1)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v10 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9)+6)))
	if v10 < v8 {
		F_slot_getsomeattrs_int(m, v9, v8)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+v6))))
			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v18)
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v20+v6<<(uint(int32(2))%32))))
			return v24
		}
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+v6))))
		*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v18)
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v20+v6<<(uint(int32(2))%32))))
		return v24
	}
}
