package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_prng_uint64_range(m *base.Module, l0 int32, l1 int64, l2 int64) int64 {
	mBase := m.M
	_ = mBase
	var v11 int64
	_ = v11
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v17 int64
	_ = v17
	var v19 int64
	_ = v19
	var v23 int64
	_ = v23
	var v25 int64
	_ = v25
	var v33 int64
	_ = v33
	var v38 int64
	_ = v38
	if base.Ui64(l2) <= base.Ui64(l1) {
		return l1
	} else {
		v11 = l2 - l1
		v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		v17 = v14
		v19 = v13
		for {
			v23 = v17 ^ v19
			v25 = base.I64_rotl(v23, int64(37))
			v33 = v23 ^ (v23<<(uint(int64(16))%64) ^ base.I64_rotl(v17, int64(24)))
			v38 = int64(base.Ui64(base.I64_rotl(v17*int64(5), int64(7))*int64(9)) >> (uint(base.I64_clz(v11)) % 64))
			if base.Ui64(v11) < base.Ui64(v38) {
				v17 = v33
				v19 = v25
				continue
			} else {
				break
			}
			break
		}
		*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v25
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v33
		return l1 + v38
	}
}
