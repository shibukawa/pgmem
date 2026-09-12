package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_mbuf_create_from_data(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v5 = F_palloc(m, int32(20))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = l0 + l1
		*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = v9
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
		v12 = int32(1)
		*(*uint16)(unsafe.Add(mBase, uint32(v5)+16)) = uint16(v12)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v9
		return v5
	}
}
