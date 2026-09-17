package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecGetResultSlotOps(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+103)))
	if v4 == int32(1) {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
		if v7 != 0 {
			if l1 == int32(0) {
				v42 = v7
				return v42
			} else {
				v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+99)))
				*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v10)
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
				return v12
			}
		} else {
			if l1 == int32(0) {
			} else {
				v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+99)))
				v30 = v16
				*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v30)
			}
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
			if v34 == int32(0) {
				return int32(_a_F_ExecGetResultSlotOps_0)
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
				v42 = v39
				return v42
			}
		}
	} else {
		if l1 == int32(0) {
		} else {
			v19 = int32(0)
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
			if v20 == v19 {
				v30 = v19
			} else {
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+4)))
				v30 = int32(base.Ui32(v23)>>(uint(int32(4))%32)) & int32(1)
			}
			*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v30)
		}
		v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
		if v34 == int32(0) {
			return int32(_a_F_ExecGetResultSlotOps_0)
		} else {
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
			v42 = v39
			return v42
		}
	}
}
func F_get_call_result_type(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v5)+24))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v9 = F_internal_get_result_type(m, v6, v7, v8, l1, l2)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
func F_wait_result_is_any_signal(m *base.Module, l0 int32) int32 {
	var v10 int32
	_ = v10
	v10 = int32(255)
	return base.B2i32(l0&int32(127) == int32(0))&base.B2i32(base.Ui32(int32(125)) < base.Ui32(int32(base.Ui32(l0)>>(uint(int32(8))%32))&v10)) | base.B2i32(base.Ui32(l0&int32(_a_F_wait_result_is_any_signal_0)-int32(1)) < base.Ui32(v10))
}
