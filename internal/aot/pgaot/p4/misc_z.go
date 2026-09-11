package p4

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_zeroinfnan(m *base.Module, l0 int64) int32 {
	return base.B2i32(base.Ui64(l0<<(uint(int64(1))%64)+int64(9007199254740992)) < base.Ui64(int64(9007199254740993)))
}
