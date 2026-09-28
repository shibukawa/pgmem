package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_gtrgm_decompress(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		if v6 == v10 {
			return base.I64_extend_i32_u(v4)
		} else {
			v15 = F_palloc(m, int32(24))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int64(0)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v15))) = base.I64_extend_i32_u(v6)
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v19
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v21
				v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4)+16)))
				*(*uint16)(unsafe.Add(mBase, uint32(v15)+16)) = uint16(v23)
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+18)))
				*(*uint8)(unsafe.Add(mBase, uint32(v15)+18)) = uint8(v25)
				return base.I64_extend_i32_u(v15)
			}
		}
	}
}
