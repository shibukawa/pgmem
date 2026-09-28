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
	var v16 int64
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l2 != 0 {
		v16 = int64(111)
	} else {
		v16 = int64(115)
	}
	v18 = F_SearchSysCache3(m, int32(3), base.I64_extend_i32_u(l0), v16, base.I64_extend_i32_u(l1))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		if v18 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
				F_errmsg_internal(m, int32(_a_F_get_op_opfamily_properties_0), v10)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_get_op_opfamily_properties_1), int32(152), int32(_a_F_get_op_opfamily_properties_2))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
			v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+22)))
			v38 = v36 + v37
			v39 = int32(*(*int16)(unsafe.Add(mBase, uint32(v38)+16)))
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = v39
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l4))) = v41
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
			*(*int32)(unsafe.Add(mBase, uint32(l5))) = v43
			F_ReleaseCatCache(m, v18)
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return
			} else {
				m.G0 = v10 + int32(16)
				return
			}
		}
	}
}
