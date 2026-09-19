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
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+92))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v9 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+84)) = v9
	*(*int64)(unsafe.Add(mBase, uint32(v6)+4)) = int64(0)
	if base.Ui32(v9) <= base.Ui32(l2) {
		v16 = v9
	} else {
		v16 = l2
	}
	if v16 != 0 {
		base.MemoryCopy(m, v6+int32(4), l1, v16)
	} else {
	}
	v21 = v6 + int32(68)
	if l3 != 0 {
		if v8 == int32(0) {
			return int32(0)
		} else {
			base.MemoryCopy(m, v21, l3, v8)
			return int32(0)
		}
	} else {
		if v8 == int32(0) {
		} else {
			base.MemoryFill(m, v21, int32(0), v8)
		}
		return int32(0)
	}
}
func F_pgmem_shmat(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v3 = int32(0)
	v4 = m.Env.Pgmem_shmat(m, l0, l1, v3)
	mBase = m.M
	if v4 <= v3 {
		if v4 != 0 {
			v11 = int32(0) - v4
		} else {
			v11 = int32(28)
		}
		*(*int32)(unsafe.Add(mBase, _c_F_pgmem_shmat[0])) = v11
		v14 = int32(-1)
	} else {
		v14 = v4
	}
	return v14
}
