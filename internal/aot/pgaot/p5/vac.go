package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_vac_cleanup_one_index(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 float64
	_ = v24
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v36 float64
	_ = v36
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	v10 = F_index_vacuum_cleanup(m, l0, l1)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 == int32(0) {
			m.G0 = v8 - int32(-64)
			return v10
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v18 = F_errstart(m, v16, int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				if v18 == int32(0) {
					m.G0 = v8 - int32(-64)
					return v10
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
					v24 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v25
					*(*float64)(unsafe.Add(mBase, uint32(v8)+40)) = v24
					*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v23 + int32(4)
					F_errmsg(m, int32(180713), v6+int32(-32))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						v36 = *(*float64)(unsafe.Add(mBase, uint32(v10)+16))
						v37 = *(*int64)(unsafe.Add(mBase, uint32(v10)+24))
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
						*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v38
						*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v37
						*(*float64)(unsafe.Add(mBase, uint32(v8))) = v36
						F_errdetail(m, int32(667868), v8)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(520592), int32(2687), int32(28877))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 - int32(-64)
								return v10
							}
						}
					}
				}
			}
		}
	}
}
