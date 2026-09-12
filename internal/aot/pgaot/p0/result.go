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
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+103)))
	if v4 == int32(1) {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
		if v7 != 0 {
			if l1 == int32(0) {
				v39 = v7
				return v39
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
				v29 = v16
				*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v29)
			}
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
			if v32 == int32(0) {
				return int32(1617588)
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
				v39 = v37
				return v39
			}
		}
	} else {
		if l1 == int32(0) {
		} else {
			v19 = int32(0)
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
			if v20 == v19 {
				v29 = v19
			} else {
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+4)))
				v29 = int32(base.Ui32(v23)>>(uint(int32(4))%32)) & int32(1)
			}
			*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v29)
		}
		v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
		if v32 == int32(0) {
			return int32(1617588)
		} else {
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
			v39 = v37
			return v39
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
	var v3 int32
	_ = v3
	var v21 int32
	_ = v21
	v3 = int32(1)
	if base.Ui32(l0&int32(65535)-v3) < base.Ui32(int32(255)) {
		v21 = v3
	} else {
		if l0&int32(127) == int32(0) {
			if base.Ui32(int32(125)) < base.Ui32(int32(base.Ui32(l0)>>(uint(int32(8))%32))&int32(255)) {
				v21 = v3
			} else {
				v21 = int32(0)
			}
		} else {
			v21 = int32(0)
		}
	}
	return v21
}
