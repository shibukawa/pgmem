package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SpGistInitPage(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	v2 = l1
	if l0&int32(3) != 0 {
	} else {
	}
	v29 = F___memset(m, l0, int32(0), int32(8192))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+10)) = int32(1572864)
	v35 = int32(8196)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)) = uint16(v35)
	v41 = int32(8184)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v41)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v41)
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	v45 = l0 + v44
	v46 = int32(65410)
	*(*uint16)(unsafe.Add(mBase, uint32(v45)+6)) = uint16(v46)
	*(*uint16)(unsafe.Add(mBase, uint32(v45))) = uint16(v2)
	return
}
func F_initSpGistState(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v14 int64
	_ = v14
	var v16 int32
	_ = v16
	var v18 int64
	_ = v18
	var v20 int32
	_ = v20
	var v22 int64
	_ = v22
	var v24 int32
	_ = v24
	var v26 int64
	_ = v26
	var v30 int32
	_ = v30
	var v32 int64
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	v10 = F_spgGetCache(m, l1)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		v12 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v12
		v14 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = v14
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v16
		v18 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+20)) = v18
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v20
		v22 = *(*int64)(unsafe.Add(mBase, uint32(v10)+28))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v22
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v24
		v26 = *(*int64)(unsafe.Add(mBase, uint32(v10)+40))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+44)) = v26
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v10)+60))
		*(*int32)(unsafe.Add(mBase, uint32(l0-int32(-64)))) = v30
		v32 = *(*int64)(unsafe.Add(mBase, uint32(v10)+52))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v32
		v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+52))
		v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
		v41 = *(*int32)(unsafe.Add(mBase, uint32(v36+v37<<(uint(int32(4))%32))+88))
		if v34 != v41 {
			v43 = F_CreateTupleDescCopy(m, v36)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return
			} else {
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
				v49 = v43 + v46<<(uint(int32(4))%32)
				*(*int32)(unsafe.Add(mBase, uint32(v49)+96)) = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(v49)+88)) = v45
				v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)))
				*(*uint16)(unsafe.Add(mBase, uint32(v49)+92)) = uint16(v53)
				v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+38)))
				*(*uint8)(unsafe.Add(mBase, uint32(v49)+102)) = uint8(v55)
				v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+39)))
				*(*uint8)(unsafe.Add(mBase, uint32(v49)+103)) = uint8(v57)
				v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
				v60 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v49)+116)) = v60
				*(*uint8)(unsafe.Add(mBase, uint32(v49)+105)) = uint8(v60)
				*(*uint8)(unsafe.Add(mBase, uint32(v49)+104)) = uint8(v59)
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
				if v65 < int32(2) {
				} else {
					v69 = v43 + int32(20)
					v70 = int32(1)
					v71 = v65 - v70
					v72 = int32(7)
					v73 = v71 & v72
					if base.Ui32(v72) <= base.Ui32(v65-int32(2)) {
						v84 = v70
						v86 = int32(0)
						for {
							v92 = v69 + v84<<(uint(int32(4))%32)
							v93 = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v92))) = v93
							*(*int32)(unsafe.Add(mBase, uint32(v92)+16)) = v93
							*(*int32)(unsafe.Add(mBase, uint32(v92)+32)) = v93
							*(*int32)(unsafe.Add(mBase, uint32(v92)+48)) = v93
							*(*int32)(unsafe.Add(mBase, uint32(v92-int32(-64)))) = v93
							*(*int32)(unsafe.Add(mBase, uint32(v92)+80)) = v93
							*(*int32)(unsafe.Add(mBase, uint32(v92)+96)) = v93
							*(*int32)(unsafe.Add(mBase, uint32(v92)+112)) = v93
							v111 = int32(8)
							v112 = v84 + v111
							v114 = v86 + v111
							if v114 != v71&int32(-8) {
								v84 = v112
								v86 = v114
								continue
							} else {
								break
							}
							break
						}
						v118 = v112
					} else {
						v118 = v70
					}
					if v73 == int32(0) {
					} else {
						v128 = int32(0)
						v129 = v118
						for {
							*(*int32)(unsafe.Add(mBase, uint32(v69+v129<<(uint(int32(4))%32)))) = int32(-1)
							v140 = int32(1)
							v143 = v128 + v140
							if v143 != v73 {
								v128 = v143
								v129 = v129 + v140
								continue
							} else {
								break
							}
							break
						}
					}
				}
				F_populate_compact_attribute(m, v43, int32(0))
				mBase = m.M
				v155 = m.ExcPending
				if v155 != 0 {
					return
				} else {
					v159 = v43
					*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v159
					v166 = F_palloc0(m, int32(16))
					mBase = m.M
					v167 = m.ExcPending
					if v167 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v166
						v170 = *(*int32)(unsafe.Add(mBase, _consts[83]))
						v171 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+80)) = uint8(v171)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v170
						return
					}
				}
			}
		} else {
			v159 = v36
			*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v159
			v166 = F_palloc0(m, int32(16))
			mBase = m.M
			v167 = m.ExcPending
			if v167 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v166
				v170 = *(*int32)(unsafe.Add(mBase, _consts[83]))
				v171 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+80)) = uint8(v171)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v170
				return
			}
		}
	}
}
