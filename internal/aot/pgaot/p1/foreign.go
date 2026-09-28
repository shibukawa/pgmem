package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ForeignNext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	v5 = int32(_a_F_ForeignNext_0)
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_ForeignNext[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ForeignNext[0])) = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v7)+80))
	if v15 == int32(1) {
		v18 = int32(20)
	} else {
		v18 = int32(96)
	}
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v12+v18)))
	v21 = m.T0[v20].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_ForeignNext[0])) = v6
		v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+120)))
		if base.B2i32(v21 == int32(0))|base.B2i32(v29 != int32(1)) != 0 {
		} else {
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+4)))
			if v33&int32(2) != 0 {
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+56))
				*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v37
			}
		}
		return v21
	}
}
func F_GetForeignServerByName(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int64
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v11 = int64(0)
	v14 = F_GetSysCacheOid(m, int32(31), base.I64_extend_i32_u(l0), v11, v11, v11)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if l1|v14 != 0 {
			if v14 != 0 {
				v20 = F_GetForeignServerExtended(m, v14, int32(0))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					v23 = v20
					m.G0 = v7 + int32(16)
					return v23
				}
			} else {
				v23 = int32(0)
				m.G0 = v7 + int32(16)
				return v23
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(67137668))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
					F_errmsg(m, int32(_a_F_GetForeignServerByName_0), v7)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_GetForeignServerByName_1), int32(757), int32(_a_F_GetForeignServerByName_2))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
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
func F_GetForeignServerIdByRelId(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14213(m, l0, int32(_a_F_GetForeignServerIdByRelId_0), int32(407), int32(_a_F_GetForeignServerIdByRelId_1), int32(_a_F_GetForeignServerIdByRelId_2), int32(33))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
