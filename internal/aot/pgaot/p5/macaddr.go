package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_macaddr_abbrev_convert(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	*(*int64)(unsafe.Add(mBase, uint32(v4))) = v5 + int64(1)
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+8)))
	if v9 == int32(1) {
		v12 = int32(16)
		v13 = v4 + v12
		v18 = int32(711645284)
		v21 = v3 - int32(1636608428) ^ v18 - int32(1455628627)
		v26 = v21 ^ int32(-1636608428) - base.I32_rotl(v21, int32(25))
		v31 = v26 ^ v18 - base.I32_rotl(v26, v12)
		v35 = v31 ^ v21 - base.I32_rotl(v31, int32(4))
		v39 = v35 ^ v26 - base.I32_rotl(v35, int32(14))
		v43 = v39 ^ v31 - base.I32_rotl(v39, int32(24))
		v46 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
		v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
		v49 = int32(32) - v48
		v51 = v46 + int32(base.Ui32(v43)>>(uint(v49)%32))
		v52 = v43 << (uint(v48) % 32)
		if v52 != 0 {
			v59 = int32(32) - (base.I32_clz(v52) ^ int32(31))
			v60 = int32(255)
			if base.Ui32(v49&v60) < base.Ui32(v59&v60) {
				v65 = v49 + int32(1)
			} else {
				v65 = v59
			}
			v69 = v65
		} else {
			v69 = v49 + int32(1)
		}
		v71 = v69 & int32(255)
		v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
		if base.Ui32(v72) < base.Ui32(v71) {
			v74 = v71
		} else {
			v74 = v72
		}
		*(*uint8)(unsafe.Add(mBase, uint32(v51))) = uint8(v74)
	} else {
	}
	v78 = int32(16711935)
	return base.I32_rotr(v3, int32(24))&v78 | base.I32_rotr(v3&v78, int32(8))
}
func F_macaddr_send(m *base.Module, l0 int32) int32 {
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
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_pq_begintypsend(m, v7)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
		F_enlargeStringInfo(m, v7, int32(1))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
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
				return int32(0)
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
					return int32(0)
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
						return int32(0)
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
							return int32(0)
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
								return int32(0)
							} else {
								v73 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
								v74 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
								*(*uint8)(unsafe.Add(mBase, uint32(v73+v74))) = uint8(v69)
								v78 = v73 + int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v78
								v81 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
								*(*int32)(unsafe.Add(mBase, uint32(v81))) = v78 << (uint(int32(2)) % 32)
								m.G0 = v7 + int32(16)
								return v81
							}
						}
					}
				}
			}
		}
	}
}
func F_macaddr_trunc(m *base.Module, l0 int32) int32 {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_palloc(m, int32(6))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
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
		return v5
	}
}
