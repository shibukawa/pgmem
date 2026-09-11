package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_zaptreesubs(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v4 <= int32(0) {
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if base.Ui32(v7) <= base.Ui32(v4) {
		} else {
			v10 = v4 << (uint(int32(3)) % 32)
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v13 = int32(-1)
			*(*int32)(unsafe.Add(mBase, uint32(v10+v11))) = v13
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			*(*int32)(unsafe.Add(mBase, uint32(v15+v10)+4)) = v13
		}
	}
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v20 != 0 {
		v22 = v20
		for {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
			if v25 <= int32(0) {
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if base.Ui32(v28) <= base.Ui32(v25) {
				} else {
					v31 = v25 << (uint(int32(3)) % 32)
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v34 = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(v31+v32))) = v34
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v36+v31)+4)) = v34
				}
			}
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
			if v41 != 0 {
				v43 = v41
				for {
					F_zaptreesubs(m, l0, v43)
					mBase = m.M
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+24))
					if v46 != 0 {
						v43 = v46
						continue
					} else {
						break
					}
					break
				}
			} else {
			}
			v50 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
			if v50 != 0 {
				v22 = v50
				continue
			} else {
				break
			}
			break
		}
	} else {
	}
	return
}
