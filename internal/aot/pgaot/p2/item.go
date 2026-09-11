package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_compare_sort_item_count(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	return base.B2i32(v5 < v4) - base.B2i32(v4 < v5)
}
func F_sort_item_compare(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v8 == int32(1) {
		if v6&int32(1) != 0 {
			v44 = int32(0)
			return v44
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
			if v16 != 0 {
				v17 = int32(-1)
			} else {
				v17 = int32(1)
			}
			return v17
		}
	} else {
		if v6&int32(1) != 0 {
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
			if v23 != 0 {
				v24 = int32(1)
			} else {
				v24 = int32(-1)
			}
			return v24
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
			v31 = m.T0[v30].(func(*base.Module, int32, int32, int32) int32)(m, v27, v29, l2)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
				if v35 != int32(1) {
					v44 = v31
				} else {
					v39 = int32(0)
					if v31 < v39 {
						v43 = int32(1)
					} else {
						v43 = v39 - v31
					}
					v44 = v43
				}
				return v44
			}
		}
	}
}
