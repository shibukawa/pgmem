package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_gtrgm_decompress(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		if v6 == v10 {
			return v4
		} else {
			v14 = F_palloc(m, int32(16))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = v6
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v17
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v19
				v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4)+12)))
				*(*uint16)(unsafe.Add(mBase, uint32(v14)+12)) = uint16(v21)
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+14)))
				*(*uint8)(unsafe.Add(mBase, uint32(v14)+14)) = uint8(v23)
				return v14
			}
		}
	}
}
