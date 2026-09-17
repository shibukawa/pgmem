package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_entryIsMoveRight(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1+v9)))
	if v11 == int32(-1) {
		v54 = int32(0)
		m.G0 = v7 + int32(16)
		return v54
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+12)))
		if base.Ui32(int32(25)) <= base.Ui32(v16) {
			v26 = int32(base.Ui32(v16+int32(_a_F_entryIsMoveRight_0))>>(uint(int32(2))%32)) & int32(_a_F_entryIsMoveRight_1)
		} else {
			v26 = int32(0)
		}
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l1+v26<<(uint(int32(2))%32))+20))
		v33 = l1 + v30&int32(_a_F_entryIsMoveRight_2)
		v34 = F_gintuple_get_attrnum(m, v15, v33)
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return int32(0)
		} else {
			v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v41 = F_gintuple_get_key(m, v38, v33, v7+int32(15))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+54)))
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				v46 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+60)))
				v47 = int32(*(*int8)(unsafe.Add(mBase, uint32(v7)+15)))
				v48 = F_ginCompareAttEntries(m, v43, v44, v45, v46, v34, v41, v47)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					v54 = base.B2i32(int32(0) < v48)
					m.G0 = v7 + int32(16)
					return v54
				}
			}
		}
	}
}
