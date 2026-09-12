package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_enum_ge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = F_enum_cmp_internal(m, v2, v3, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return int32(base.Ui32(v4^int32(-1)) >> (uint(int32(31)) % 32))
	}
}
func F_enum_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_SearchSysCache1(m, int32(23), v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50462850))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = v9
					F_errmsg(m, int32(58115), v6)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(493891), int32(167), int32(66816))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
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
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+22)))
			v37 = F_pstrdup(m, v32+v33+int32(12))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				F_ReleaseCatCache(m, v10)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					m.G0 = v6 + int32(16)
					return v37
				}
			}
		}
	}
}
func F_enum_send(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_SearchSysCache1(m, int32(23), v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50462850))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = v9
					F_errmsg(m, int32(58115), v6)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(493891), int32(233), int32(423748))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
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
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+22)))
			v36 = v32 + v33 + int32(12)
			F_pq_begintypsend(m, v6+int32(16))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				v43 = F_strlen(m, v36)
				mBase = m.M
				F_pq_sendtext(m, v6+int32(16), v36, v43)
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					F_ReleaseCatCache(m, v10)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						v49 = v6 + int32(16)
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v51))) = v52 << (uint(int32(2)) % 32)
						m.G0 = v6 + int32(32)
						return v51
					}
				}
			}
		}
	}
}
