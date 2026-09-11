package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_btree_identify(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	v8 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(2))%32))&int32(60))+uint32(_c_F_btree_identify[0])))
	return v8
}
func F_btree_mask(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	v3 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)) = uint16(v3)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
	v12 = v10 & int32(_a_F_btree_mask_0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)) = uint16(v12)
	F_mask_unused_space(m, l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
		v17 = l0 + v16
		v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+12)))
		if v18&int32(1) != 0 {
			v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
			if base.Ui32(v28) < base.Ui32(int32(25)) {
			} else {
				v34 = int32(base.Ui32(v28+int32(_a_F_btree_mask_1)) >> (uint(int32(2)) % 32))
				if v34&int32(_a_F_btree_mask_2) == int32(0) {
				} else {
					v39 = int32(1)
					v40 = int32(2)
					v44 = (v34 + v39) & int32(_a_F_btree_mask_2)
					if base.Ui32(v44) <= base.Ui32(v40) {
						v47 = v40
					} else {
						v47 = v44
					}
					v48 = int32(1)
					v49 = v47 - v48
					v53 = l0 + int32(24)
					v54 = int32(0)
					if base.Ui32(int32(3)) <= base.Ui32(v44) {
						v59 = v54
						v60 = v39
						for {
							v69 = v60<<(uint(int32(2))%32) + v53
							v71 = v69 - int32(4)
							v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
							if v72&int32(_a_F_btree_mask_3) != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v71))) = v72 & int32(-98305)
							} else {
							}
							v78 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
							if v78&int32(_a_F_btree_mask_3) != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v69))) = v78 & int32(-98305)
							} else {
							}
							v84 = int32(2)
							v87 = v59 + v84
							if v87 != v49&int32(-2) {
								v59 = v87
								v60 = v60 + v84
								continue
							} else {
								break
							}
							break
						}
						v91 = v60 + int32(1)
					} else {
						v91 = v54
					}
					if v49&v48 == int32(0) {
					} else {
						v103 = v53 + v91<<(uint(int32(2))%32)
						v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
						if v104&int32(_a_F_btree_mask_3) == int32(0) {
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v103))) = v104 & int32(-98305)
						}
					}
				}
			}
			v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+12)))
			v121 = v120
		} else {
			v121 = v18
		}
		v122 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(v17)+14)) = uint16(v122)
		v125 = v121 & int32(_a_F_btree_mask_4)
		*(*uint16)(unsafe.Add(mBase, uint32(v17)+12)) = uint16(v125)
		return
	}
}
