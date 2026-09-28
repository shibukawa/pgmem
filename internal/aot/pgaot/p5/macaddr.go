package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_macaddr_abbrev_convert(m *base.Module, l0 int64, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v8 int64
	_ = v8
	var v13 int64
	_ = v13
	var v23 int64
	_ = v23
	v4 = base.I32_wrap_i64(l0)
	v5 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v4)+4)))
	v6 = int64(24)
	v8 = int64(8)
	v13 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v4))))
	v23 = v13 | v5<<(uint(int64(32))%64)
	return (v5<<(uint(v6)%64)|v5<<(uint(v8)%64))&int64(4294901760) | (v13<<(uint(int64(56))%64) | v13&int64(65280)<<(uint(int64(40))%64) | (v23&int64(16711680)<<(uint(v6)%64) | v23&int64(4278190080)<<(uint(v8)%64)))
}
func F_macaddr_send(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_pq_begintypsend(m, v7)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
		F_enlargeStringInfo(m, v7, int32(1))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int64(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			*(*uint8)(unsafe.Add(mBase, uint32(v18+v19))) = uint8(v14)
			v22 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v18 + v22
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
			F_enlargeStringInfo(m, v7, v22)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int64(0)
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				*(*uint8)(unsafe.Add(mBase, uint32(v29+v30))) = uint8(v25)
				v33 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v29 + v33
				v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+2)))
				F_enlargeStringInfo(m, v7, v33)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int64(0)
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
					*(*uint8)(unsafe.Add(mBase, uint32(v40+v41))) = uint8(v36)
					v44 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v40 + v44
					v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+3)))
					F_enlargeStringInfo(m, v7, v44)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int64(0)
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
						*(*uint8)(unsafe.Add(mBase, uint32(v51+v52))) = uint8(v47)
						v55 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v51 + v55
						v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+4)))
						F_enlargeStringInfo(m, v7, v55)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int64(0)
						} else {
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
							*(*uint8)(unsafe.Add(mBase, uint32(v62+v63))) = uint8(v58)
							v66 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v62 + v66
							v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+5)))
							F_enlargeStringInfo(m, v7, v66)
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return int64(0)
							} else {
								v73 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
								v74 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
								*(*uint8)(unsafe.Add(mBase, uint32(v73+v74))) = uint8(v69)
								v78 = v73 + int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v78
								v81 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
								*(*int32)(unsafe.Add(mBase, uint32(v81))) = v78 << (uint(int32(2)) % 32)
								m.G0 = v7 + int32(16)
								return base.I64_extend_i32_u(v81)
							}
						}
					}
				}
			}
		}
	}
}
func F_macaddr_trunc(m *base.Module, l0 int32) int64 {
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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_palloc(m, int32(6))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
		*(*uint8)(unsafe.Add(mBase, uint32(v5))) = uint8(v9)
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+1)))
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+1)) = uint8(v11)
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+2)))
		v14 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+5)) = uint8(v14)
		*(*uint16)(unsafe.Add(mBase, uint32(v5)+3)) = uint16(v14)
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+2)) = uint8(v13)
		return base.I64_extend_i32_u(v5)
	}
}
