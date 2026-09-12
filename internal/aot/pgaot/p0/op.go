package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_op_opfamily_properties(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l2 != 0 {
		v15 = int32(111)
	} else {
		v15 = int32(115)
	}
	v16 = F_SearchSysCache3(m, int32(3), l0, v15, l1)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		if v16 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
				F_errmsg_internal(m, int32(40791), v10)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					F_errfinish(m, int32(511204), int32(151), int32(171202))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
			v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+22)))
			v36 = v34 + v35
			v37 = int32(*(*int16)(unsafe.Add(mBase, uint32(v36)+16)))
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = v37
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l4))) = v39
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
			*(*int32)(unsafe.Add(mBase, uint32(l5))) = v41
			F_ReleaseCatCache(m, v16)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return
			} else {
				m.G0 = v10 + int32(16)
				return
			}
		}
	}
}
