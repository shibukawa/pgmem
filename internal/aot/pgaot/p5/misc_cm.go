package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CMPTRGM_SIGNED(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v5 != v6 {
		if base.I32_extend8_s(v5) < base.I32_extend8_s(v6) {
			v13 = int32(-1)
		} else {
			v13 = int32(1)
		}
		return v13
	} else {
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
		if v15 != v16 {
			if base.I32_extend8_s(v15) < base.I32_extend8_s(v16) {
				v23 = int32(-1)
			} else {
				v23 = int32(1)
			}
			return v23
		} else {
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
			if v26 != v27 {
				if base.I32_extend8_s(v26) < base.I32_extend8_s(v27) {
					v34 = int32(-1)
				} else {
					v34 = int32(1)
				}
				v35 = v34
			} else {
				v35 = int32(0)
			}
			return v35
		}
	}
}
