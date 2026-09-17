package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tts_buffer_heap_copy_heap_tuple(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v4 != 0 {
		v32 = v4
		v34 = F_heap_copytuple(m, v32)
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return int32(0)
		} else {
			return v34
		}
	} else {
		v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
		if v6&int32(4) != 0 {
			v32 = int32(0)
			v34 = F_heap_copytuple(m, v32)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				return v34
			}
		} else {
			v9 = int32(_a_F_tts_buffer_heap_copy_heap_tuple_0)
			v10 = *(*int32)(unsafe.Add(mBase, _c_F_tts_buffer_heap_copy_heap_tuple[0]))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			*(*int32)(unsafe.Add(mBase, _c_F_tts_buffer_heap_copy_heap_tuple[0])) = v12
			v14 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v14)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v14
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v21 = F_heap_form_tuple(m, v18, v19, v20)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v21
				v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
				v28 = v26 | int32(4)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v28)
				*(*int32)(unsafe.Add(mBase, _c_F_tts_buffer_heap_copy_heap_tuple[0])) = v10
				v32 = v21
				v34 = F_heap_copytuple(m, v32)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					return v34
				}
			}
		}
	}
}
func F_tts_heap_is_current_xact_tuple(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14010(m, l0, int32(_a_F_tts_heap_is_current_xact_tuple_0), int32(391))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_tts_virtual_copy_heap_tuple(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_heap_form_tuple(m, v2, v3, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
