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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v32 int64
	_ = v32
	var v36 int32
	_ = v36
	var v46 int32
	_ = v46
	var v49 int64
	_ = v49
	var v53 int64
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	v3 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v7 == v3 {
		v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+96))
		v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+120)))
		v25 = v13
		v26 = v10 - v11
		if v25&int32(255) != base.B2i32(v7 != int32(0)) {
			if v25&int32(1) != 0 {
				v46 = v3
			} else {
				v46 = int32(1)
			}
		} else {
			v32 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
			if v26 <= v32 {
				if v25&int32(1) != 0 {
					v46 = v3
				} else {
					v46 = int32(1)
				}
			} else {
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+120)) = uint8(v25)
				*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v26
				v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v36
				if v25&int32(1) == int32(0) {
					v46 = int32(1)
				} else {
					v46 = v3
				}
			}
		}
	} else {
		v14 = *(*int64)(unsafe.Add(mBase, uint32(v7)+24))
		v15 = *(*int64)(unsafe.Add(mBase, uint32(v7)+32))
		v18 = (v14 - v15) << (uint(int64(13)) % 64)
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+120)))
		if v19 != 0 {
			v25 = v19
			v26 = v18
			if v25&int32(255) != base.B2i32(v7 != int32(0)) {
				if v25&int32(1) != 0 {
					v46 = v3
				} else {
					v46 = int32(1)
				}
			} else {
				v32 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
				if v26 <= v32 {
					if v25&int32(1) != 0 {
						v46 = v3
					} else {
						v46 = int32(1)
					}
				} else {
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+120)) = uint8(v25)
					*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v26
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v36
					if v25&int32(1) == int32(0) {
						v46 = int32(1)
					} else {
						v46 = v3
					}
				}
			}
		} else {
			v20 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+120)) = uint8(v20)
			*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v18
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v23
			v46 = v3
		}
	}
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v46
	v49 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	v53 = base.I64_div_s(v49+int64(1023), int64(1024))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v53
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	switch v55 - int32(3) {
	case 0:
		v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+69)))
		if v60 != 0 {
			v61 = int32(1)
		} else {
			v61 = int32(2)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v61
		return
	case 1:
		v66 = v55
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v66
		return
	case 2:
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(8)
		return
	default:
		v66 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v66
		return
	}
}
