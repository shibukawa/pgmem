package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_execute_attr_map_tuple(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_heap_deform_tuple(m, l0, v12, v13+int32(4), v16+int32(1))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		if int32(0) < v23 {
			v27 = int32(0)
			for {
				v35 = int32(2)
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				v39 = int32(1)
				v42 = int32(*(*int16)(unsafe.Add(mBase, uint32(v38+v27<<(uint(v39)%32)))))
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v13+v42<<(uint(v35)%32))))
				*(*int32)(unsafe.Add(mBase, uint32(v10+v27<<(uint(v35)%32)))) = v46
				v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+v42))))
				*(*uint8)(unsafe.Add(mBase, uint32(v27+v9))) = uint8(v50)
				v53 = v27 + v39
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
				if v53 < v54 {
					v27 = v53
					continue
				} else {
					break
				}
				break
			}
		} else {
		}
		v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v65 = F_heap_form_tuple(m, v64, v10, v9)
		mBase = m.M
		v66 = m.ExcPending
		if v66 != 0 {
			return int32(0)
		} else {
			return v65
		}
	}
}
