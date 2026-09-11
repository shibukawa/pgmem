package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_socket_comm_reset(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[373])) = uint8(v2)
	return
}
