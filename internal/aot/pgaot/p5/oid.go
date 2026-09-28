package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_OidFunctionCall3Coll(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64) int64 {
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
	var v28 int32
	_ = v28
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	v6 = m.G0
	v8 = v6 - int32(112)
	m.G0 = v8
	v11 = v8 + int32(12)
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_OidFunctionCall3Coll[0]))
	F_fmgr_info_cxt_security(m, l0, v11, v13, int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int64(0)
	} else {
		v19 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v8)+104)) = uint8(v19)
		*(*int64)(unsafe.Add(mBase, uint32(v8)+96)) = l3
		*(*uint8)(unsafe.Add(mBase, uint32(v8)+88)) = uint8(v19)
		*(*int64)(unsafe.Add(mBase, uint32(v8)+80)) = l2
		*(*uint8)(unsafe.Add(mBase, uint32(v8)+72)) = uint8(v19)
		*(*int64)(unsafe.Add(mBase, uint32(v8)+64)) = l1
		v28 = int32(3)
		*(*uint16)(unsafe.Add(mBase, uint32(v8)+58)) = uint16(v28)
		*(*uint8)(unsafe.Add(mBase, uint32(v8)+56)) = uint8(v19)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = v19
		*(*int64)(unsafe.Add(mBase, uint32(v8)+44)) = int64(0)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v11
		v39 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
		v40 = m.T0[v39].(func(*base.Module, int32) int64)(m, v8+int32(40))
		mBase = m.M
		v41 = m.ExcPending
		if v41 != 0 {
			return int64(0)
		} else {
			v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+56)))
			if v42 == int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int64(0)
				} else {
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v49
					F_errmsg_internal(m, int32(_a_F_OidFunctionCall3Coll_0), v8)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_OidFunctionCall3Coll_1), int32(1192), int32(_a_F_OidFunctionCall3Coll_2))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				m.G0 = v8 + int32(112)
				return v40
			}
		}
	}
}
func F_OidOutputFunctionCall(m *base.Module, l0 int32, l1 int64) int32 {
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
	var v30 int64
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
	v6 = v4 - int32(80)
	m.G0 = v6
	v9 = v6 + int32(12)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_OidOutputFunctionCall[0]))
	F_fmgr_info_cxt_security(m, l0, v9, v11, int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v6)+44)) = v17
		*(*int64)(unsafe.Add(mBase, uint32(v6)+49)) = v17
		v21 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v6)+72)) = uint8(v21)
		*(*int64)(unsafe.Add(mBase, uint32(v6)+64)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = v9
		v25 = int32(1)
		*(*uint16)(unsafe.Add(mBase, uint32(v6)+58)) = uint16(v25)
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
		v30 = m.T0[v29].(func(*base.Module, int32) int64)(m, v6+int32(40))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+56)))
			if v32 == int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = v39
					F_errmsg_internal(m, int32(_a_F_OidOutputFunctionCall_0), v6)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_OidOutputFunctionCall_1), int32(1145), int32(_a_F_OidOutputFunctionCall_2))
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
				m.G0 = v6 + int32(80)
				return base.I32_wrap_i64(v30)
			}
		}
	}
}
func F_OidReceiveFunctionCall(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = v9 + int32(4)
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_OidReceiveFunctionCall[0]))
	F_fmgr_info_cxt_security(m, l0, v12, v14, int32(0))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int64(0)
	} else {
		v20 = F_ReceiveFunctionCall(m, v12, l1, l2, l3)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int64(0)
		} else {
			m.G0 = v9 + int32(32)
			return v20
		}
	}
}
