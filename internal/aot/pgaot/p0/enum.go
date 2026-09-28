package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_enum_ge(m *base.Module, l0 int32) int64 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v4 = F_enum_cmp_internal(m, v2, v3, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(base.B2i32(int32(0) <= v4))
	}
}
func F_enum_out(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+24)))
	v11 = F_SearchSysCache1(m, int32(23), v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		if v11 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(50462850))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int64(0)
				} else {
					*(*uint32)(unsafe.Add(mBase, uint32(v7))) = uint32(v10)
					F_errmsg(m, int32(_a_F_enum_out_0), v7)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_enum_out_1), int32(167), int32(_a_F_enum_out_2))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+22)))
			v38 = F_pstrdup(m, v33+v34+int32(12))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int64(0)
			} else {
				F_ReleaseCatCache(m, v11)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int64(0)
				} else {
					m.G0 = v7 + int32(16)
					return base.I64_extend_i32_u(v38)
				}
			}
		}
	}
}
func F_enum_send(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+24)))
	v12 = F_SearchSysCache1(m, int32(23), v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		if v12 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(50462850))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int64(0)
				} else {
					*(*uint32)(unsafe.Add(mBase, uint32(v8))) = uint32(v11)
					F_errmsg(m, int32(_a_F_enum_send_0), v8)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_enum_send_1), int32(233), int32(_a_F_enum_send_2))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+22)))
			v38 = v34 + v35 + int32(12)
			v40 = v8 + int32(16)
			F_pq_begintypsend(m, v40)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int64(0)
			} else {
				v43 = F_strlen(m, v38)
				mBase = m.M
				F_pq_sendtext(m, v40, v38, v43)
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int64(0)
				} else {
					F_ReleaseCatCache(m, v12)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int64(0)
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v49))) = v50 << (uint(int32(2)) % 32)
						m.G0 = v8 + int32(32)
						return base.I64_extend_i32_u(v49)
					}
				}
			}
		}
	}
}
