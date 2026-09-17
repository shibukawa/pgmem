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
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v6 = v4 ^ v5
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = base.I64_rotl(v6, int64(37))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v6<<(uint(int64(16))%64) ^ base.I64_rotl(v4, int64(24)) ^ v6
	return base.F64_mul(base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v4*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(971))<<(uint(int64(52))%64)))
}
func F_pg_prng_uint32(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	v4 = int32(_a_F_pg_prng_uint32_0)
	v5 = int32(_a_F_pg_prng_uint32_1)
	v6 = *(*int64)(unsafe.Add(mBase, _c_F_pg_prng_uint32[0]))
	v8 = *(*int64)(unsafe.Add(mBase, _c_F_pg_prng_uint32[1]))
	v9 = v6 ^ v8
	*(*int64)(unsafe.Add(mBase, _c_F_pg_prng_uint32[1])) = base.I64_rotl(v9, int64(37))
	*(*int64)(unsafe.Add(mBase, _c_F_pg_prng_uint32[0])) = v9<<(uint(int64(16))%64) ^ base.I64_rotl(v6, int64(24)) ^ v9
	return base.I32_wrap_i64(int64(base.Ui64(base.I64_rotl(v6*int64(5), int64(7))*int64(9)) >> (uint(int64(32)) % 64)))
}
