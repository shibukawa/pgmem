package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_makeParamList(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	v7 = F_palloc(m, l0*int32(12)+int32(32))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v7))) = v11
		*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(814)
		*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v11
		*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v7
		return v7
	}
}
