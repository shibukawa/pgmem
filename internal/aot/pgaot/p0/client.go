package p0

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_ClientCheckTimeoutHandler(m *base.Module) {
	var v3 int32
	_ = v3
	Fn13831(m, int32(_a_F_ClientCheckTimeoutHandler_0))
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		return
	}
}
