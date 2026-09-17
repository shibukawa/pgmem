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
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	v2 = l1
	v3 = int32(_a_F_SpGistInitPage_0)
	v5 = int32(0)
	if v5|(l0&int32(3)|int32(1)) == v5 {
		v21 = l0 + v3
		v23 = l0 + int32(4)
		if base.Ui32(v23) < base.Ui32(v21) {
			v25 = v21
		} else {
			v25 = v23
		}
		v30 = (l0^int32(-1)+v25)&int32(-4) + int32(4)
		if v30 == int32(0) {
		} else {
			base.MemoryFill(m, l0, int32(0), v30)
		}
	} else {
		base.MemoryFill(m, l0, int32(0), v3)
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+10)) = int32(_a_F_SpGistInitPage_1)
	v44 = int32(_a_F_SpGistInitPage_2)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)) = uint16(v44)
	v50 = int32(_a_F_SpGistInitPage_3)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v50)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v50)
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	v54 = l0 + v53
	v55 = int32(_a_F_SpGistInitPage_4)
	*(*uint16)(unsafe.Add(mBase, uint32(v54)+6)) = uint16(v55)
	*(*uint16)(unsafe.Add(mBase, uint32(v54))) = uint16(v2)
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
	var v28 int32
	_ = v28
	var v30 int64
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	v10 = F_spgGetCache(m, l1)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		v12 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = v12
		v14 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v14
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
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v10)+60))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v28
		v30 = *(*int64)(unsafe.Add(mBase, uint32(v10)+52))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v30
		v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+52))
		v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
		v39 = *(*int32)(unsafe.Add(mBase, uint32(v34+v35<<(uint(int32(4))%32))+88))
		if v32 != v39 {
			v41 = F_CreateTupleDescCopy(m, v34)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return
			} else {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
				v47 = v41 + v44<<(uint(int32(4))%32)
				*(*int32)(unsafe.Add(mBase, uint32(v47)+96)) = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(v47)+88)) = v43
				v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)))
				*(*uint16)(unsafe.Add(mBase, uint32(v47)+92)) = uint16(v51)
				v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+38)))
				*(*uint8)(unsafe.Add(mBase, uint32(v47)+102)) = uint8(v53)
				v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+39)))
				*(*uint8)(unsafe.Add(mBase, uint32(v47)+103)) = uint8(v55)
				v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
				v58 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v47)+116)) = v58
				*(*uint8)(unsafe.Add(mBase, uint32(v47)+105)) = uint8(v58)
				*(*uint8)(unsafe.Add(mBase, uint32(v47)+104)) = uint8(v57)
				v63 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
				if v63 < int32(2) {
				} else {
					v67 = v41 + int32(20)
					v68 = int32(1)
					v69 = v63 - v68
					v70 = int32(7)
					v71 = v69 & v70
					if base.Ui32(v70) <= base.Ui32(v63-int32(2)) {
						v82 = v68
						v84 = int32(0)
						for {
							v90 = v67 + v82<<(uint(int32(4))%32)
							v91 = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v90)+112)) = v91
							*(*int32)(unsafe.Add(mBase, uint32(v90)+96)) = v91
							*(*int32)(unsafe.Add(mBase, uint32(v90)+80)) = v91
							*(*int32)(unsafe.Add(mBase, uint32(v90)+64)) = v91
							*(*int32)(unsafe.Add(mBase, uint32(v90)+48)) = v91
							*(*int32)(unsafe.Add(mBase, uint32(v90)+32)) = v91
							*(*int32)(unsafe.Add(mBase, uint32(v90)+16)) = v91
							*(*int32)(unsafe.Add(mBase, uint32(v90))) = v91
							v107 = int32(8)
							v108 = v82 + v107
							v110 = v84 + v107
							if v110 != v69&int32(-8) {
								v82 = v108
								v84 = v110
								continue
							} else {
								break
							}
							break
						}
						if v71 == int32(0) {
						} else {
							v116 = v108
							v124 = int32(0)
							v125 = v116
							for {
								*(*int32)(unsafe.Add(mBase, uint32(v67+v125<<(uint(int32(4))%32)))) = int32(-1)
								v136 = int32(1)
								v139 = v124 + v136
								if v139 != v71 {
									v124 = v139
									v125 = v125 + v136
									continue
								} else {
									break
								}
								break
							}
						}
					} else {
						v116 = v68
						v124 = int32(0)
						v125 = v116
						for {
							*(*int32)(unsafe.Add(mBase, uint32(v67+v125<<(uint(int32(4))%32)))) = int32(-1)
							v136 = int32(1)
							v139 = v124 + v136
							if v139 != v71 {
								v124 = v139
								v125 = v125 + v136
								continue
							} else {
								break
							}
							break
						}
					}
				}
				F_populate_compact_attribute(m, v41, int32(0))
				mBase = m.M
				v151 = m.ExcPending
				if v151 != 0 {
					return
				} else {
					v155 = v41
					*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v155
					v162 = F_palloc0(m, int32(16))
					mBase = m.M
					v163 = m.ExcPending
					if v163 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v162
						v166 = *(*int32)(unsafe.Add(mBase, _c_F_initSpGistState[0]))
						v167 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+80)) = uint8(v167)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v166
						return
					}
				}
			}
		} else {
			v155 = v34
			*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v155
			v162 = F_palloc0(m, int32(16))
			mBase = m.M
			v163 = m.ExcPending
			if v163 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v162
				v166 = *(*int32)(unsafe.Add(mBase, _c_F_initSpGistState[0]))
				v167 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+80)) = uint8(v167)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v166
				return
			}
		}
	}
}
