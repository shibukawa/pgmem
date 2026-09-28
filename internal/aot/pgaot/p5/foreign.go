package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ForeignRecheck(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = l1
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	F_MemoryContextReset(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v12)+112))
		if v20 != 0 {
			v22 = m.T0[v20].(func(*base.Module, int32, int32) int32)(m, l0, l1)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				if v22 == int32(0) {
					v47 = int32(0)
					m.G0 = v10 + int32(16)
					return v47
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
					if v27 == int32(0) {
						v47 = int32(1)
						m.G0 = v10 + int32(16)
						return v47
					} else {
						v31 = int32(_a_F_ForeignRecheck_0)
						v32 = *(*int32)(unsafe.Add(mBase, _c_F_ForeignRecheck[0]))
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
						*(*int32)(unsafe.Add(mBase, _c_F_ForeignRecheck[0])) = v34
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
						v39 = m.T0[v38].(func(*base.Module, int32, int32, int32) int64)(m, v27, v13, v10+int32(15))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_ForeignRecheck[0])) = v32
							v47 = base.B2i32(v39 != int64(0))
							m.G0 = v10 + int32(16)
							return v47
						}
					}
				}
			}
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
			if v27 == int32(0) {
				v47 = int32(1)
				m.G0 = v10 + int32(16)
				return v47
			} else {
				v31 = int32(_a_F_ForeignRecheck_0)
				v32 = *(*int32)(unsafe.Add(mBase, _c_F_ForeignRecheck[0]))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
				*(*int32)(unsafe.Add(mBase, _c_F_ForeignRecheck[0])) = v34
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
				v39 = m.T0[v38].(func(*base.Module, int32, int32, int32) int64)(m, v27, v13, v10+int32(15))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_ForeignRecheck[0])) = v32
					v47 = base.B2i32(v39 != int64(0))
					m.G0 = v10 + int32(16)
					return v47
				}
			}
		}
	}
}
func F_GetForeignDataWrapperByName(m *base.Module, l0 int32, l1 int32) int32 {
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
	v14 = F_GetSysCacheOid(m, int32(29), base.I64_extend_i32_u(l0), v11, v11, v11)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if l1|v14 != 0 {
			if v14 != 0 {
				v20 = F_GetForeignDataWrapperExtended(m, v14, int32(0))
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
					F_errmsg(m, int32(_a_F_GetForeignDataWrapperByName_0), v7)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_GetForeignDataWrapperByName_1), int32(736), int32(_a_F_GetForeignDataWrapperByName_2))
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
func F_get_foreign_data_wrapper_oid(m *base.Module, l0 int32, l1 int32) int32 {
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v9 = Fn14290(m, l0, l1, int32(_a_F_get_foreign_data_wrapper_oid_0), int32(736), int32(_a_F_get_foreign_data_wrapper_oid_1), int32(_a_F_get_foreign_data_wrapper_oid_2), int32(67137668), int32(29))
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
