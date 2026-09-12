package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_prng_int64_range(m *base.Module, l0 int64, l1 int64) int64 {
	mBase := m.M
	_ = mBase
	var v9 int64
	_ = v9
	var v12 int64
	_ = v12
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v18 int64
	_ = v18
	var v22 int64
	_ = v22
	var v24 int64
	_ = v24
	var v32 int64
	_ = v32
	var v37 int64
	_ = v37
	var v51 int64
	_ = v51
	if l0 < l1 {
		v9 = l1 - l0
		v12 = *(*int64)(unsafe.Add(mBase, _consts[1127]))
		v14 = *(*int64)(unsafe.Add(mBase, _consts[1126]))
		v16 = v14
		v18 = v12
		for {
			v22 = v16 ^ v18
			v24 = base.I64_rotl(v22, int64(37))
			v32 = v22 ^ (v22<<(uint(int64(16))%64) ^ base.I64_rotl(v16, int64(24)))
			v37 = int64(base.Ui64(base.I64_rotl(v16*int64(5), int64(7))*int64(9)) >> (uint(base.I64_clz(v9)) % 64))
			if base.Ui64(v9) < base.Ui64(v37) {
				v16 = v32
				v18 = v24
				continue
			} else {
				break
			}
			break
		}
		*(*int64)(unsafe.Add(mBase, _consts[1127])) = v24
		*(*int64)(unsafe.Add(mBase, _consts[1126])) = v32
		v51 = l0 + v37
	} else {
		v51 = l0
	}
	return v51
}
