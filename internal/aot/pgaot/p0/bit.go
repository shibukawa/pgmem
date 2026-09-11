package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bit_send(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		F_pq_begintypsend(m, v7)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
			F_enlargeStringInfo(m, v7, int32(4))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				v23 = int32(24)
				v25 = int32(65280)
				v27 = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v20+v21))) = v16<<(uint(v23)%32) | v16&v25<<(uint(v27)%32) | (int32(base.Ui32(v16)>>(uint(v27)%32))&v25 | int32(base.Ui32(v16)>>(uint(v23)%32)))
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v20 + int32(4)
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
				F_pq_sendbytes(m, v7, v10+v27, int32(base.Ui32(v44)>>(uint(int32(2))%32))-v27)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v52))) = v53 << (uint(int32(2)) % 32)
					m.G0 = v7 + int32(16)
					return v52
				}
			}
		}
	}
}
