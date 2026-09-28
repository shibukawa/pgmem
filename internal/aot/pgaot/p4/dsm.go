package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_dsm_main_space_init(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_main_space_init[0]))
	if v3 != 0 {
		v5 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_main_space_init[1]))
		v6 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v5)+4)) = v6
		*(*int64)(unsafe.Add(mBase, uint32(v5)+12)) = v6
		*(*int64)(unsafe.Add(mBase, uint32(v5)+20)) = v6
		v12 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+28)) = v12
		v14 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+32)) = uint8(v14)
		if v5 != 0 {
			v20 = v5 - v5 + v14
		} else {
			v20 = v12
		}
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = v20
		base.MemoryFill(m, v5+int32(36), int32(0), int32(516))
		v27 = int32(1)
		v29 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_main_space_init[0]))
		F_FreePageManagerPut(m, v5, v27, int32(base.Ui32(v29)>>(uint(int32(12))%32))-v27)
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
