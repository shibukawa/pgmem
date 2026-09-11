package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetCurrentSubTransactionId(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentSubTransactionId[0]))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+8))
	return v3
}
