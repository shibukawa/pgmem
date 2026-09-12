package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgmem_des_init(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+92))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v9 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+84)) = v9
	*(*int64)(unsafe.Add(mBase, uint32(v6)+4)) = int64(0)
	if base.Ui32(v9) <= base.Ui32(l2) {
		v18 = v9
	} else {
		v18 = l2
	}
	if v18 != 0 {
		v19 = F__emscripten_memcpy_bulkmem(m, v6+int32(4), l1, v18)
		mBase = m.M
	} else {
	}
	v22 = v6 + int32(68)
	if l3 != 0 {
		if v8 != 0 {
			v23 = F__emscripten_memcpy_bulkmem(m, v22, l3, v8)
			mBase = m.M
		} else {
		}
		return int32(0)
	} else {
		v29 = F__emscripten_memset_bulkmem(m, v22, base.I32_extend8_s(int32(0)), v8)
		mBase = m.M
		return int32(0)
	}
}
