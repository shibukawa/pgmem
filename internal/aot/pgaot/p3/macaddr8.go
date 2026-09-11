package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_macaddr8_recv(m *base.Module, l0 int32) int32 {
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_palloc0(m, int32(8))
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
					v18 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
					if v18 == int32(6) {
						v21 = int32(255)
						*(*uint8)(unsafe.Add(mBase, uint32(v5)+3)) = uint8(v21)
						v29 = int32(254)
						*(*uint8)(unsafe.Add(mBase, uint32(v5)+4)) = uint8(v29)
						v31 = F_pq_getmsgbyte(m, v3)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							*(*uint8)(unsafe.Add(mBase, uint32(v5)+5)) = uint8(v31)
							v34 = F_pq_getmsgbyte(m, v3)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								*(*uint8)(unsafe.Add(mBase, uint32(v5)+6)) = uint8(v34)
								v37 = F_pq_getmsgbyte(m, v3)
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return int32(0)
								} else {
									*(*uint8)(unsafe.Add(mBase, uint32(v5)+7)) = uint8(v37)
									return v5
								}
							}
						}
					} else {
						v24 = F_pq_getmsgbyte(m, v3)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							*(*uint8)(unsafe.Add(mBase, uint32(v5)+3)) = uint8(v24)
							v27 = F_pq_getmsgbyte(m, v3)
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return int32(0)
							} else {
								v29 = v27
								*(*uint8)(unsafe.Add(mBase, uint32(v5)+4)) = uint8(v29)
								v31 = F_pq_getmsgbyte(m, v3)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return int32(0)
								} else {
									*(*uint8)(unsafe.Add(mBase, uint32(v5)+5)) = uint8(v31)
									v34 = F_pq_getmsgbyte(m, v3)
									mBase = m.M
									v35 = m.ExcPending
									if v35 != 0 {
										return int32(0)
									} else {
										*(*uint8)(unsafe.Add(mBase, uint32(v5)+6)) = uint8(v34)
										v37 = F_pq_getmsgbyte(m, v3)
										mBase = m.M
										v38 = m.ExcPending
										if v38 != 0 {
											return int32(0)
										} else {
											*(*uint8)(unsafe.Add(mBase, uint32(v5)+7)) = uint8(v37)
											return v5
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
