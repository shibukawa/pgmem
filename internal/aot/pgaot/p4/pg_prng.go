package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_prng_bool(m *base.Module) int32 {
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
	v3 = int32(4508792)
	v4 = int32(4508784)
	v5 = *(*int64)(unsafe.Add(mBase, _consts[180]))
	v7 = *(*int64)(unsafe.Add(mBase, _consts[181]))
	v8 = v5 ^ v7
	*(*int64)(unsafe.Add(mBase, _consts[181])) = base.I64_rotl(v8, int64(37))
	*(*int64)(unsafe.Add(mBase, _consts[180])) = v8<<(uint(int64(16))%64) ^ base.I64_rotl(v5, int64(24)) ^ v8
	return base.I32_wrap_i64(int64(base.Ui64(base.I64_rotl(v5*int64(5), int64(7))*int64(9)) >> (uint(int64(63)) % 64)))
}
func F_pg_prng_seed_check(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v5 int64
	_ = v5
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if v2 != int64(0) {
	} else {
		v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		if v5 != int64(0) {
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(1442695040888963407)
			*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(6364136223846793005)
		}
	}
	return int32(1)
}
