package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_byteaSetBit(m *base.Module, l0 int32) int64 {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum_copy(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		v24 = base.I64_extend_i32_u(int32(base.Ui32(v17)>>(uint(int32(2))%32)))<<(uint(int64(3))%64) - int64(32)
		v25 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		if base.B2i32(v25 < int64(0))|base.B2i32(v24 <= v25) == int32(0) {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			if base.Ui32(int32(2)) <= base.Ui32(v32) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_byteaSetBit_0), int32(0))
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_byteaSetBit_1), int32(767), int32(_a_F_byteaSetBit_2))
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
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
				v38 = v13 + base.I32_wrap_i64(int64(base.Ui64(v25)>>(uint(int64(3))%64)))
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+4)))
				v43 = base.I32_wrap_i64(v25) & int32(7)
				if v32 != 0 {
					v49 = v39 | int32(1)<<(uint(v43)%32)
				} else {
					v49 = v39 & base.I32_rotl(int32(-2), v43)
				}
				*(*uint8)(unsafe.Add(mBase, uint32(v38)+4)) = uint8(v49)
				m.G0 = v10 + int32(16)
				return base.I64_extend_i32_u(v13)
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(352845954))
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int64(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v24 - int64(1)
					*(*int64)(unsafe.Add(mBase, uint32(v10))) = v25
					F_errmsg(m, int32(_a_F_byteaSetBit_3), v10)
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_byteaSetBit_1), int32(755), int32(_a_F_byteaSetBit_2))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int64(0)
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
func F_bytea_string_agg_finalfn(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v5 != 0 {
		v30 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v30)
		return int64(0)
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if v6 == int32(0) {
			v30 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v30)
			return int64(0)
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
			v11 = v9 - v10
			v13 = v11 + int32(4)
			v14 = F_palloc(m, v13)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = v13 << (uint(int32(2)) % 32)
				if v11 != 0 {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
					base.MemoryCopy(m, v14+int32(4), v23+v24, v11)
				} else {
				}
				return base.I64_extend_i32_u(v14)
			}
		}
	}
}
