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
	var v16 int32
	_ = v16
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	if l1 == int32(0) {
		*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
		v67 = int32(1)
		m.G0 = v9 + int32(48)
		return v67
	} else {
		v16 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)) = uint8(v16)
		*(*int64)(unsafe.Add(mBase, uint32(v9)+12)) = int64(0)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v16
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+44)) = uint8(v16)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = l2
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+36)) = uint8(v16)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v16
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)) = uint8(v16)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = l1
		v33 = int32(3)
		*(*uint16)(unsafe.Add(mBase, uint32(v9)+22)) = uint16(v33)
		v37 = m.T0[l0].(func(*base.Module, int32) int32)(m, v9+int32(4))
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l4))) = v37
			if l3 == int32(0) {
				v50 = int32(1)
				v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
				if v51 != v50 {
					v67 = v50
					m.G0 = v9 + int32(48)
					return v67
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
						F_errmsg_internal(m, int32(529218), v9)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(492035), int32(1670), int32(407262))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
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
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
				if v44 != int32(447) {
					v50 = int32(1)
					v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
					if v51 != v50 {
						v67 = v50
						m.G0 = v9 + int32(48)
						return v67
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
							F_errmsg_internal(m, int32(529218), v9)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(492035), int32(1670), int32(407262))
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
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
					v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)))
					if v48 != 0 {
						v67 = int32(0)
						m.G0 = v9 + int32(48)
						return v67
					} else {
						v50 = int32(1)
						v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
						if v51 != v50 {
							v67 = v50
							m.G0 = v9 + int32(48)
							return v67
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
								F_errmsg_internal(m, int32(529218), v9)
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(492035), int32(1670), int32(407262))
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
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
