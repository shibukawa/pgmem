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
	var v17 int32
	_ = v17
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
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
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1+v9)))
	if v11 == int32(-1) {
		v56 = int32(0)
		m.G0 = v7 + int32(16)
		return v56
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+12)))
		if base.Ui32(v17) < base.Ui32(int32(25)) {
			v28 = int32(-1)
		} else {
			v28 = int32(base.Ui32(v17+int32(262120))>>(uint(int32(2))%32))&int32(65535) - int32(1)
		}
		v32 = *(*int32)(unsafe.Add(mBase, uint32(l1+v28<<(uint(int32(2))%32))+24))
		v35 = l1 + v32&int32(32767)
		v36 = F_gintuple_get_attrnum(m, v15, v35)
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return int32(0)
		} else {
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v43 = F_gintuple_get_key(m, v40, v35, v7+int32(15))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int32(0)
			} else {
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+54)))
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				v48 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+60)))
				v49 = int32(*(*int8)(unsafe.Add(mBase, uint32(v7)+15)))
				v50 = F_ginCompareAttEntries(m, v45, v46, v47, v48, v36, v43, v49)
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					v56 = base.B2i32(int32(0) < v50)
					m.G0 = v7 + int32(16)
					return v56
				}
			}
		}
	}
}
