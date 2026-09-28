package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DirectInputFunctionCallSafe(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int64
	_ = v16
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v39 int64
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	if l1 == int32(0) {
		*(*int64)(unsafe.Add(mBase, uint32(l4))) = int64(0)
		v69 = int32(1)
		m.G0 = v9 + int32(80)
		return v69
	} else {
		v16 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v16
		*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l3
		v19 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v19
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+24)) = uint8(v19)
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+72)) = uint8(v19)
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+56)) = uint8(v19)
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+40)) = uint8(v19)
		v29 = int32(3)
		*(*uint16)(unsafe.Add(mBase, uint32(v9)+26)) = uint16(v29)
		*(*int64)(unsafe.Add(mBase, uint32(v9)+64)) = base.I64_extend_i32_s(l2)
		*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = v16
		*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = base.I64_extend_i32_u(l1)
		v39 = m.T0[l0].(func(*base.Module, int32) int64)(m, v9+int32(8))
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return int32(0)
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(l4))) = v39
			if l3 == int32(0) {
				v52 = int32(1)
				v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+24)))
				if v53 != v52 {
					v69 = v52
					m.G0 = v9 + int32(80)
					return v69
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
						F_errmsg_internal(m, int32(_a_F_DirectInputFunctionCallSafe_0), v9)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_DirectInputFunctionCallSafe_1), int32(1671), int32(_a_F_DirectInputFunctionCallSafe_2))
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
				if v46 != int32(453) {
					v52 = int32(1)
					v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+24)))
					if v53 != v52 {
						v69 = v52
						m.G0 = v9 + int32(80)
						return v69
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
							F_errmsg_internal(m, int32(_a_F_DirectInputFunctionCallSafe_0), v9)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_DirectInputFunctionCallSafe_1), int32(1671), int32(_a_F_DirectInputFunctionCallSafe_2))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)))
					if v50 != 0 {
						v69 = int32(0)
						m.G0 = v9 + int32(80)
						return v69
					} else {
						v52 = int32(1)
						v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+24)))
						if v53 != v52 {
							v69 = v52
							m.G0 = v9 + int32(80)
							return v69
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
								F_errmsg_internal(m, int32(_a_F_DirectInputFunctionCallSafe_0), v9)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_DirectInputFunctionCallSafe_1), int32(1671), int32(_a_F_DirectInputFunctionCallSafe_2))
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
