package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ForeignRecheck(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = l1
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	F_MemoryContextReset(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v11)+112))
		if v19 != 0 {
			v21 = m.T0[v19].(func(*base.Module, int32, int32) int32)(m, l0, l1)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				if v21 == int32(0) {
					v46 = int32(0)
					m.G0 = v9 + int32(16)
					return v46
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
					if v26 == int32(0) {
						v46 = int32(1)
						m.G0 = v9 + int32(16)
						return v46
					} else {
						v30 = int32(4464496)
						v31 = *(*int32)(unsafe.Add(mBase, _consts[9]))
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
						*(*int32)(unsafe.Add(mBase, _consts[9])) = v33
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
						v38 = m.T0[v37].(func(*base.Module, int32, int32, int32) int32)(m, v26, v12, v9+int32(15))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[9])) = v31
							v46 = base.B2i32(v38 != int32(0))
							m.G0 = v9 + int32(16)
							return v46
						}
					}
				}
			}
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
			if v26 == int32(0) {
				v46 = int32(1)
				m.G0 = v9 + int32(16)
				return v46
			} else {
				v30 = int32(4464496)
				v31 = *(*int32)(unsafe.Add(mBase, _consts[9]))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
				*(*int32)(unsafe.Add(mBase, _consts[9])) = v33
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
				v38 = m.T0[v37].(func(*base.Module, int32, int32, int32) int32)(m, v26, v12, v9+int32(15))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[9])) = v31
					v46 = base.B2i32(v38 != int32(0))
					m.G0 = v9 + int32(16)
					return v46
				}
			}
		}
	}
}
func F_GetForeignDataWrapperByName(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v13 = F_GetSysCacheOid(m, int32(29), l0, v3, v3, v3)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if l1|v13 != 0 {
			if v13 != 0 {
				v19 = F_GetForeignDataWrapperExtended(m, v13, int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					v21 = v19
					m.G0 = v7 + int32(16)
					return v21
				}
			} else {
				v21 = v3
				m.G0 = v7 + int32(16)
				return v21
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(67137668))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
					F_errmsg(m, int32(69496), v7)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(485525), int32(693), int32(424437))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
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
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v13 = F_GetSysCacheOid(m, int32(29), l0, v3, v3, v3)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if l1 != 0 {
			m.G0 = v7 + int32(16)
			return v13
		} else {
			if v13 != 0 {
				m.G0 = v7 + int32(16)
				return v13
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
						F_errmsg(m, int32(69496), v7)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(485525), int32(693), int32(424437))
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
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
