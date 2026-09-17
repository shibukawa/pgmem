package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_btree_identify(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	v6 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(2))%32))&int32(60))+uint32(_c_F_btree_identify[0])))
	return v6
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
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
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
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
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
			v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
			if base.Ui32(v27) < base.Ui32(int32(25)) {
			} else {
				v33 = int32(base.Ui32(v27+int32(_a_F_btree_mask_1)) >> (uint(int32(2)) % 32))
				if v33&int32(_a_F_btree_mask_2) == int32(0) {
				} else {
					v38 = int32(1)
					v40 = l0 + int32(20)
					v44 = (v33 + v38) & int32(_a_F_btree_mask_2)
					if base.Ui32(int32(3)) <= base.Ui32(v44) {
						v47 = int32(2)
						if base.Ui32(v44) <= base.Ui32(v47) {
							v50 = v47
						} else {
							v50 = v44
						}
						v51 = int32(1)
						v52 = v50 - v51
						v60 = v51
						v61 = int32(0)
						for {
							v68 = v40 + v60<<(uint(int32(2))%32)
							v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
							if v69&int32(_a_F_btree_mask_3) != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v68))) = v69 & int32(-98305)
							} else {
							}
							v76 = v68 + int32(4)
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
							if v77&int32(_a_F_btree_mask_3) != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v76))) = v77 & int32(-98305)
							} else {
							}
							v83 = int32(2)
							v84 = v60 + v83
							v86 = v61 + v83
							if v86 != v52&int32(-2) {
								v60 = v84
								v61 = v86
								continue
							} else {
								break
							}
							break
						}
						if v52&v51 == int32(0) {
						} else {
							v91 = v84
							v99 = v40 + v91<<(uint(int32(2))%32)
							v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
							if v100&int32(_a_F_btree_mask_3) == int32(0) {
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v99))) = v100 & int32(-98305)
							}
						}
					} else {
						v91 = v38
						v99 = v40 + v91<<(uint(int32(2))%32)
						v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
						if v100&int32(_a_F_btree_mask_3) == int32(0) {
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v99))) = v100 & int32(-98305)
						}
					}
				}
			}
			v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+12)))
			v116 = v115
		} else {
			v116 = v18
		}
		v117 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(v17)+14)) = uint16(v117)
		v120 = v116 & int32(_a_F_btree_mask_4)
		*(*uint16)(unsafe.Add(mBase, uint32(v17)+12)) = uint16(v120)
		return
	}
}
