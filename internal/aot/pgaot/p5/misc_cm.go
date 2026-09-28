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
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v5 != v6 {
		v23 = v5
		v24 = v6
		if base.I32_extend8_s(v23) < base.I32_extend8_s(v24) {
			v30 = int32(-1)
		} else {
			v30 = int32(1)
		}
		return v30
	} else {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
		if v8 != v9 {
			v23 = v8
			v24 = v9
			if base.I32_extend8_s(v23) < base.I32_extend8_s(v24) {
				v30 = int32(-1)
			} else {
				v30 = int32(1)
			}
			return v30
		} else {
			v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
			v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
			if v11 != v12 {
				if base.I32_extend8_s(v11) < base.I32_extend8_s(v12) {
					v19 = int32(-1)
				} else {
					v19 = int32(1)
				}
				v21 = v19
			} else {
				v21 = int32(0)
			}
			return v21
		}
	}
}
