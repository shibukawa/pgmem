package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_macaddr8_and(m *base.Module, l0 int32) int32 {
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
		v13 = v11 & v12
		*(*uint8)(unsafe.Add(mBase, uint32(v7))) = uint8(v13)
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+1)))
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+1)))
		v17 = v15 & v16
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)) = uint8(v17)
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+2)))
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+2)))
		v21 = v19 & v20
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+2)) = uint8(v21)
		v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+3)))
		v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+3)))
		v25 = v23 & v24
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+3)) = uint8(v25)
		v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+4)))
		v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+4)))
		v29 = v27 & v28
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)) = uint8(v29)
		v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+5)))
		v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+5)))
		v33 = v31 & v32
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+5)) = uint8(v33)
		v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+6)))
		v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+6)))
		v37 = v35 & v36
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+6)) = uint8(v37)
		v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+7)))
		v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+7)))
		v41 = v39 & v40
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+7)) = uint8(v41)
		return v7
	}
}
func F_macaddr8_ne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	v5 = int32(1)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v8 = int32(24)
	v10 = int32(_a_F_macaddr8_ne_0)
	v12 = int32(8)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v7<<(uint(v8)%32)|v7&v10<<(uint(v12)%32)|(int32(base.Ui32(v7)>>(uint(v12)%32))&v10|int32(base.Ui32(v7)>>(uint(v8)%32))) != v24<<(uint(v8)%32)|v24&v10<<(uint(v12)%32)|(int32(base.Ui32(v24)>>(uint(v12)%32))&v10|int32(base.Ui32(v24)>>(uint(v8)%32))) {
		v77 = v5
	} else {
		v41 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
		v42 = int32(24)
		v44 = int32(_a_F_macaddr8_ne_0)
		v46 = int32(8)
		v56 = v41<<(uint(v42)%32) | v41&v44<<(uint(v46)%32) | (int32(base.Ui32(v41)>>(uint(v46)%32))&v44 | int32(base.Ui32(v41)>>(uint(v42)%32)))
		v57 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
		v72 = v57<<(uint(v42)%32) | v57&v44<<(uint(v46)%32) | (int32(base.Ui32(v57)>>(uint(v46)%32))&v44 | int32(base.Ui32(v57)>>(uint(v42)%32)))
		if base.Ui32(v56) < base.Ui32(v72) {
			v77 = v5
		} else {
			v77 = base.B2i32(base.Ui32(v72) < base.Ui32(v56))
		}
	}
	return v77
}
