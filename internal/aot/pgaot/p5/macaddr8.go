package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_macaddr8_cmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v8 = int32(24)
	v10 = int32(_a_F_macaddr8_cmp_0)
	v12 = int32(8)
	v22 = v7<<(uint(v8)%32) | v7&v10<<(uint(v12)%32) | (int32(base.Ui32(v7)>>(uint(v12)%32))&v10 | int32(base.Ui32(v7)>>(uint(v8)%32)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v39 = v24<<(uint(v8)%32) | v24&v10<<(uint(v12)%32) | (int32(base.Ui32(v24)>>(uint(v12)%32))&v10 | int32(base.Ui32(v24)>>(uint(v8)%32)))
	if base.Ui32(v22) < base.Ui32(v39) {
		v80 = int32(-1)
	} else {
		if base.Ui32(v39) < base.Ui32(v22) {
			v80 = int32(1)
		} else {
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
			v45 = int32(24)
			v47 = int32(_a_F_macaddr8_cmp_0)
			v49 = int32(8)
			v59 = v44<<(uint(v45)%32) | v44&v47<<(uint(v49)%32) | (int32(base.Ui32(v44)>>(uint(v49)%32))&v47 | int32(base.Ui32(v44)>>(uint(v45)%32)))
			v60 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
			v75 = v60<<(uint(v45)%32) | v60&v47<<(uint(v49)%32) | (int32(base.Ui32(v60)>>(uint(v49)%32))&v47 | int32(base.Ui32(v60)>>(uint(v45)%32)))
			if base.Ui32(v59) < base.Ui32(v75) {
				v80 = int32(-1)
			} else {
				v80 = base.B2i32(base.Ui32(v75) < base.Ui32(v59))
			}
		}
	}
	return v80
}
func F_macaddr8_set7bit(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_palloc0(m, int32(8))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
		v11 = v9 | int32(2)
		*(*uint8)(unsafe.Add(mBase, uint32(v5))) = uint8(v11)
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+1)))
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+1)) = uint8(v13)
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+2)))
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+2)) = uint8(v15)
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+3)))
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+3)) = uint8(v17)
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+4)))
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+4)) = uint8(v19)
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+5)))
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+5)) = uint8(v21)
		v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+6)))
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+6)) = uint8(v23)
		v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+7)))
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+7)) = uint8(v25)
		return v5
	}
}
