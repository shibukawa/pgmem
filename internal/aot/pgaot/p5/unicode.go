package p5

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_unicode_to_utf8word(m *base.Module, l0 int32) int32 {
	var v52 int32
	_ = v52
	if base.Ui32(int32(128)) <= base.Ui32(l0) {
		if base.Ui32(l0) <= base.Ui32(int32(2047)) {
			return l0&int32(63) | l0<<(uint(int32(2))%32)&int32(7936) | int32(49280)
		} else {
			if base.Ui32(l0) <= base.Ui32(int32(65535)) {
				return l0<<(uint(int32(4))%32)&int32(983040) | (l0&int32(63) | l0<<(uint(int32(2))%32)&int32(16128)) | int32(14712960)
			} else {
				v52 = l0<<(uint(int32(2))%32)&int32(16128) | (l0<<(uint(int32(6))%32)&int32(117440512) | (l0&int32(63) | l0<<(uint(int32(4))%32)&int32(4128768))) | int32(-260013952)
				return v52
			}
		}
	} else {
		v52 = l0
		return v52
	}
}
