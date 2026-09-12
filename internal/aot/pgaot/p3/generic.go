package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GenericXLogStart(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	v6 = F_palloc_aligned(m, int32(69632), int32(4096), int32(0))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+118)))
		if v12 != int32(112) {
			v25 = int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, _consts[18]))
			if int32(0) < v17 {
				v25 = int32(1)
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				if v21 != 0 {
					v25 = int32(0)
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					v25 = base.B2i32(v22 == int32(0))
				}
			}
		}
		*(*uint8)(unsafe.Add(mBase, uint32(v6)+uint32(_consts[92]))) = uint8(v25)
		v27 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_consts[93]))) = v27
		*(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_consts[94]))) = v6 - int32(-8192)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_consts[95]))) = v6 + int32(16384)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_consts[96]))) = v27
		*(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_consts[97]))) = v6 + int32(24576)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_consts[98]))) = v27
		*(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_consts[99]))) = v27
		*(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_consts[100]))) = v6
		return v6
	}
}
