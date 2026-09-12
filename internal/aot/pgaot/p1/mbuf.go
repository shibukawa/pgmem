package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_mbuf_append(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v7 == int32(1) {
		F_px_debug(m, int32(349149), int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			return int32(-12)
		}
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if base.Ui32(v18) < base.Ui32(v19+l2) {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v28 = v18 - v22 + (l2+int32(32767))&int32(-16384)
			v29 = F_repalloc(m, v22, v28)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v29 + v28
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v29
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v37 = v29 + (v35 - v33)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v37
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v29 + (v39 - v33)
				v43 = v37
				if l2 != 0 {
					v46 = F__emscripten_memcpy_bulkmem(m, v43, l1, l2)
					mBase = m.M
				} else {
				}
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v48 + l2
				return int32(0)
			}
		} else {
			v43 = v19
			if l2 != 0 {
				v46 = F__emscripten_memcpy_bulkmem(m, v43, l1, l2)
				mBase = m.M
			} else {
			}
			v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v48 + l2
			return int32(0)
		}
	}
}
func F_mbuf_steal_data(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	v4 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v4)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v7
	v9 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v9
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v9
	return v6 - v7
}
