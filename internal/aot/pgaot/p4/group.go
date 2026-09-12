package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_create_group_result_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v34 float64
	_ = v34
	var v37 float64
	_ = v37
	var v39 float64
	_ = v39
	var v43 int32
	_ = v43
	var v44 float64
	_ = v44
	var v45 float64
	_ = v45
	var v46 float64
	_ = v46
	var v47 float64
	_ = v47
	var v50 float64
	_ = v50
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = F_palloc0(m, int32(80))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v13)+20)) = uint8(v17)
		*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v17
		*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = l1
		*(*int64)(unsafe.Add(mBase, uint32(v13))) = int64(1421634175268)
		v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
		*(*int32)(unsafe.Add(mBase, uint32(v13)+72)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v17
		*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v17
		*(*uint8)(unsafe.Add(mBase, uint32(v13)+21)) = uint8(v25)
		*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = int64(4607182418800017408)
		v34 = *(*float64)(unsafe.Add(mBase, uint32(l2)+16))
		*(*float64)(unsafe.Add(mBase, uint32(v13)+48)) = v34
		v37 = *(*float64)(unsafe.Add(mBase, _consts[611]))
		v39 = *(*float64)(unsafe.Add(mBase, uint32(l2)+24))
		*(*float64)(unsafe.Add(mBase, uint32(v13)+56)) = base.F64_add(base.F64_add(v34, v37), v39)
		if l3 != 0 {
			F_cost_qual_eval(m, v10, l3, l0)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				v44 = *(*float64)(unsafe.Add(mBase, uint32(v13)+48))
				v45 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
				v46 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
				v47 = base.F64_add(v45, v46)
				*(*float64)(unsafe.Add(mBase, uint32(v13)+48)) = base.F64_add(v44, v47)
				v50 = *(*float64)(unsafe.Add(mBase, uint32(v13)+56))
				*(*float64)(unsafe.Add(mBase, uint32(v13)+56)) = base.F64_add(v47, v50)
				m.G0 = v10 + int32(16)
				return v13
			}
		} else {
			m.G0 = v10 + int32(16)
			return v13
		}
	}
}
