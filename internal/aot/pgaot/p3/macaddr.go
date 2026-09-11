package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_macaddr_recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_palloc(m, int32(6))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = F_pq_getmsgbyte(m, v3)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			*(*uint8)(unsafe.Add(mBase, uint32(v5))) = uint8(v9)
			v12 = F_pq_getmsgbyte(m, v3)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				*(*uint8)(unsafe.Add(mBase, uint32(v5)+1)) = uint8(v12)
				v15 = F_pq_getmsgbyte(m, v3)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int32(0)
				} else {
					*(*uint8)(unsafe.Add(mBase, uint32(v5)+2)) = uint8(v15)
					v18 = F_pq_getmsgbyte(m, v3)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return int32(0)
					} else {
						*(*uint8)(unsafe.Add(mBase, uint32(v5)+3)) = uint8(v18)
						v21 = F_pq_getmsgbyte(m, v3)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int32(0)
						} else {
							*(*uint8)(unsafe.Add(mBase, uint32(v5)+4)) = uint8(v21)
							v24 = F_pq_getmsgbyte(m, v3)
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return int32(0)
							} else {
								*(*uint8)(unsafe.Add(mBase, uint32(v5)+5)) = uint8(v24)
								return v5
							}
						}
					}
				}
			}
		}
	}
}
