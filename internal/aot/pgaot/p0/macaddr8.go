package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_macaddr8_gt(m *base.Module, l0 int32) int32 {
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
	v10 = int32(65280)
	v12 = int32(8)
	v22 = v7<<(uint(v8)%32) | v7&v10<<(uint(v12)%32) | (int32(base.Ui32(v7)>>(uint(v12)%32))&v10 | int32(base.Ui32(v7)>>(uint(v8)%32)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v39 = v24<<(uint(v8)%32) | v24&v10<<(uint(v12)%32) | (int32(base.Ui32(v24)>>(uint(v12)%32))&v10 | int32(base.Ui32(v24)>>(uint(v8)%32)))
	if base.Ui32(v22) < base.Ui32(v39) {
		v80 = int32(0)
	} else {
		if base.Ui32(v39) < base.Ui32(v22) {
			v80 = int32(1)
		} else {
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
			v45 = int32(24)
			v47 = int32(65280)
			v49 = int32(8)
			v59 = v44<<(uint(v45)%32) | v44&v47<<(uint(v49)%32) | (int32(base.Ui32(v44)>>(uint(v49)%32))&v47 | int32(base.Ui32(v44)>>(uint(v45)%32)))
			v60 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
			v75 = v60<<(uint(v45)%32) | v60&v47<<(uint(v49)%32) | (int32(base.Ui32(v60)>>(uint(v49)%32))&v47 | int32(base.Ui32(v60)>>(uint(v45)%32)))
			if base.Ui32(v59) < base.Ui32(v75) {
				v80 = int32(0)
			} else {
				v80 = base.B2i32(base.Ui32(v75) < base.Ui32(v59))
			}
		}
	}
	return v80
}
func F_macaddr8_le(m *base.Module, l0 int32) int32 {
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
	v10 = int32(65280)
	v12 = int32(8)
	v22 = v7<<(uint(v8)%32) | v7&v10<<(uint(v12)%32) | (int32(base.Ui32(v7)>>(uint(v12)%32))&v10 | int32(base.Ui32(v7)>>(uint(v8)%32)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v39 = v24<<(uint(v8)%32) | v24&v10<<(uint(v12)%32) | (int32(base.Ui32(v24)>>(uint(v12)%32))&v10 | int32(base.Ui32(v24)>>(uint(v8)%32)))
	if base.Ui32(v22) < base.Ui32(v39) {
		v80 = int32(1)
	} else {
		if base.Ui32(v39) < base.Ui32(v22) {
			v80 = int32(0)
		} else {
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
			v45 = int32(24)
			v47 = int32(65280)
			v49 = int32(8)
			v59 = v44<<(uint(v45)%32) | v44&v47<<(uint(v49)%32) | (int32(base.Ui32(v44)>>(uint(v49)%32))&v47 | int32(base.Ui32(v44)>>(uint(v45)%32)))
			v60 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
			v75 = v60<<(uint(v45)%32) | v60&v47<<(uint(v49)%32) | (int32(base.Ui32(v60)>>(uint(v49)%32))&v47 | int32(base.Ui32(v60)>>(uint(v45)%32)))
			if base.Ui32(v59) < base.Ui32(v75) {
				v80 = int32(1)
			} else {
				v80 = base.B2i32(base.Ui32(v59) <= base.Ui32(v75))
			}
		}
	}
	return v80
}
func F_macaddr8_not(m *base.Module, l0 int32) int32 {
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
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_palloc0(m, int32(8))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
		v10 = int32(-1)
		v11 = v9 ^ v10
		*(*uint8)(unsafe.Add(mBase, uint32(v5))) = uint8(v11)
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+1)))
		v15 = v13 ^ v10
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+1)) = uint8(v15)
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+2)))
		v19 = v17 ^ v10
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+2)) = uint8(v19)
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+3)))
		v23 = v21 ^ v10
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+3)) = uint8(v23)
		v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+4)))
		v27 = v25 ^ v10
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+4)) = uint8(v27)
		v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+5)))
		v31 = v29 ^ v10
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+5)) = uint8(v31)
		v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+6)))
		v35 = v33 ^ v10
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+6)) = uint8(v35)
		v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+7)))
		v39 = v37 ^ v10
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+7)) = uint8(v39)
		return v5
	}
}
func F_macaddr8_or(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
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
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_palloc0(m, int32(8))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
		v13 = v11 | v12
		*(*uint8)(unsafe.Add(mBase, uint32(v7))) = uint8(v13)
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+1)))
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+1)))
		v17 = v15 | v16
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)) = uint8(v17)
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+2)))
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+2)))
		v21 = v19 | v20
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+2)) = uint8(v21)
		v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+3)))
		v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+3)))
		v25 = v23 | v24
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+3)) = uint8(v25)
		v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+4)))
		v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+4)))
		v29 = v27 | v28
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)) = uint8(v29)
		v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+5)))
		v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+5)))
		v33 = v31 | v32
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+5)) = uint8(v33)
		v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+6)))
		v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+6)))
		v37 = v35 | v36
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+6)) = uint8(v37)
		v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+7)))
		v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+7)))
		v41 = v39 | v40
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+7)) = uint8(v41)
		return v7
	}
}
