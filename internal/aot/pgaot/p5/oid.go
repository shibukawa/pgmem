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
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	v4 = m.G0
	v6 = v4 + int32(-64)
	m.G0 = v6
	v9 = v4 + int32(-56)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_OidOutputFunctionCall[0]))
	F_fmgr_info_cxt_security(m, l0, v9, v11, int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v6)+40)) = v17
		*(*int64)(unsafe.Add(mBase, uint32(v6)+45)) = v17
		v21 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v6)+60)) = uint8(v21)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+56)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = v9
		v25 = int32(1)
		*(*uint16)(unsafe.Add(mBase, uint32(v6)+54)) = uint16(v25)
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
		v30 = m.T0[v29].(func(*base.Module, int32) int32)(m, v4+int32(-28))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+52)))
			if v32 == int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = v39
					F_errmsg_internal(m, int32(_a_F_OidOutputFunctionCall_0), v6)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_OidOutputFunctionCall_1), int32(1143), int32(_a_F_OidOutputFunctionCall_2))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
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
				return v30
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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = v8 + int32(4)
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_OidReceiveFunctionCall[0]))
	F_fmgr_info_cxt_security(m, l0, v11, v13, int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = F_ReceiveFunctionCall(m, v11, l1, l2, l3)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			m.G0 = v8 + int32(32)
			return v19
		}
	}
}
