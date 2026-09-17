package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_mbuf_create(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v5 = F_palloc(m, int32(20))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if l0 != 0 {
			v10 = l0
		} else {
			v10 = int32(_a_F_mbuf_create_0)
		}
		v11 = F_palloc(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = v11
			v14 = int32(256)
			*(*uint16)(unsafe.Add(mBase, uint32(v5)+16)) = uint16(v14)
			*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = v11
			*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v11
			*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = v11 + v10
			return v5
		}
	}
}
func F_mbuf_free(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)))
	if v3 == int32(1) {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v9 = v8 - v6
		if v9 != 0 {
			base.MemoryFill(m, v6, int32(0), v9)
		} else {
		}
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		F_pfree(m, v11)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			F_pfree(m, l0)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				return int32(0)
			}
		}
	} else {
		F_pfree(m, l0)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			return int32(0)
		}
	}
}
func F_mbuf_grab(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v5 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v5)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v8
	v10 = v7 - v8
	if l1 < v10 {
		v12 = l1
	} else {
		v12 = v10
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v8 + v12
	return v12
}
