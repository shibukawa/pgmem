package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_OidOutputFunctionCall(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	v4 = m.G0
	v6 = v4 + int32(-64)
	m.G0 = v6
	v11 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	F_fmgr_info_cxt_security(m, l0, v4+int32(-56), v11, int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v6)+45)) = v17
		*(*int64)(unsafe.Add(mBase, uint32(v6)+40)) = v17
		v21 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v6)+60)) = uint8(v21)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+56)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = v4 + int32(-56)
		v27 = int32(1)
		*(*uint16)(unsafe.Add(mBase, uint32(v6)+54)) = uint16(v27)
		v31 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
		v32 = m.T0[v31].(func(*base.Module, int32) int32)(m, v4+int32(-28))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+52)))
			if v34 == int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = v41
					F_errmsg_internal(m, int32(510130), v6)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(473651), int32(1143), int32(290453))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				m.G0 = v6 - int32(-64)
				return v32
			}
		}
	}
}
func F_OidReceiveFunctionCall(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v13 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	F_fmgr_info_cxt_security(m, l0, v8+int32(4), v13, int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v21 = F_ReceiveFunctionCall(m, v8+int32(4), l1, l2, l3)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			m.G0 = v8 + int32(32)
			return v21
		}
	}
}
