package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_prng_double(m *base.Module, l0 int32) float64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v27 float64
	_ = v27
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v6 = v4 ^ v5
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = base.I64_rotl(v6, int64(37))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v6<<(uint(int64(16))%64) ^ base.I64_rotl(v4, int64(24)) ^ v6
	v27 = F_scalbn(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v4*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
	mBase = m.M
	return v27
}
func F_pg_prng_uint32(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	v3 = int32(4599224)
	v4 = int32(4599216)
	v5 = *(*int64)(unsafe.Add(mBase, _consts[819]))
	v7 = *(*int64)(unsafe.Add(mBase, _consts[818]))
	v8 = v5 ^ v7
	*(*int64)(unsafe.Add(mBase, _consts[818])) = base.I64_rotl(v8, int64(37))
	*(*int64)(unsafe.Add(mBase, _consts[819])) = v8<<(uint(int64(16))%64) ^ base.I64_rotl(v5, int64(24)) ^ v8
	return base.I32_wrap_i64(int64(base.Ui64(base.I64_rotl(v5*int64(5), int64(7))*int64(9)) >> (uint(int64(32)) % 64)))
}
