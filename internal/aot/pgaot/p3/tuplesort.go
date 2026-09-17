package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tuplesort_get_stats(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int64
	_ = v10
	var v11 int64
	_ = v11
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v18 int64
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v33 int64
	_ = v33
	var v37 int32
	_ = v37
	var v47 int32
	_ = v47
	var v50 int64
	_ = v50
	var v54 int64
	_ = v54
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	v3 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v7 == v3 {
		v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+96))
		v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+120)))
		v26 = v13
		v27 = v10 - v11
		if v26&int32(255) != base.B2i32(v7 != int32(0)) {
			if v26&int32(1) != 0 {
				v47 = v3
			} else {
				v47 = int32(1)
			}
		} else {
			v33 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
			if v27 <= v33 {
				if v26&int32(1) != 0 {
					v47 = v3
				} else {
					v47 = int32(1)
				}
			} else {
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+120)) = uint8(v26)
				*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v27
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v37
				if v26&int32(1) == int32(0) {
					v47 = int32(1)
				} else {
					v47 = v3
				}
			}
		}
	} else {
		v14 = *(*int64)(unsafe.Add(mBase, uint32(v7)+24))
		v15 = *(*int64)(unsafe.Add(mBase, uint32(v7)+32))
		v18 = (v14 - v15) << (uint(int64(13)) % 64)
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+120)))
		if v20 != 0 {
			v26 = int32(1)
			v27 = v18
			if v26&int32(255) != base.B2i32(v7 != int32(0)) {
				if v26&int32(1) != 0 {
					v47 = v3
				} else {
					v47 = int32(1)
				}
			} else {
				v33 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
				if v27 <= v33 {
					if v26&int32(1) != 0 {
						v47 = v3
					} else {
						v47 = int32(1)
					}
				} else {
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+120)) = uint8(v26)
					*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v27
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v37
					if v26&int32(1) == int32(0) {
						v47 = int32(1)
					} else {
						v47 = v3
					}
				}
			}
		} else {
			v21 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+120)) = uint8(v21)
			*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v18
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v24
			v47 = v3
		}
	}
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v47
	v50 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	v54 = base.I64_div_s(v50+int64(1023), int64(1024))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v54
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	switch v56 - int32(3) {
	case 0:
		v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+69)))
		if v61 != 0 {
			v62 = int32(1)
		} else {
			v62 = int32(2)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v62
		return
	case 1:
		v67 = v56
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v67
		return
	case 2:
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(8)
		return
	default:
		v67 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v67
		return
	}
}
