package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_byteaSetBit(m *base.Module, l0 int32) int32 {
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
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum_copy(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		v24 = base.I64_extend_i32_u(int32(base.Ui32(v17)>>(uint(int32(2))%32)))<<(uint(int64(3))%64) - int64(32)
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v26 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
		if v26 < int64(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(352845954))
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int32(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v24 - int64(1)
					*(*int64)(unsafe.Add(mBase, uint32(v10))) = v26
					F_errmsg(m, int32(449806), v10)
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(524383), int32(3428), int32(110408))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
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
			if v24 <= v26 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(352845954))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v24 - int64(1)
						*(*int64)(unsafe.Add(mBase, uint32(v10))) = v26
						F_errmsg(m, int32(449806), v10)
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(524383), int32(3428), int32(110408))
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
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
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				if base.Ui32(int32(2)) <= base.Ui32(v30) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(588492), int32(0))
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(524383), int32(3440), int32(110408))
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
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
					v38 = v13 + base.I32_wrap_i64(int64(base.Ui64(v26)>>(uint(int64(3))%64))) + int32(4)
					v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
					v43 = base.I32_wrap_i64(v26) & int32(7)
					if v30 != 0 {
						v49 = v39 | int32(1)<<(uint(v43)%32)
					} else {
						v49 = v39 & base.I32_rotl(int32(-2), v43)
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v38))) = uint8(v49)
					m.G0 = v10 + int32(16)
					return v13
				}
			}
		}
	}
}
func F_bytea_string_agg_finalfn(m *base.Module, l0 int32) int32 {
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
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v5 != 0 {
		v30 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v30)
		return int32(0)
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v6 == int32(0) {
			v30 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v30)
			return int32(0)
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
			v11 = v9 - v10
			v13 = v11 + int32(4)
			v14 = F_palloc(m, v13)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = v13 << (uint(int32(2)) % 32)
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
				if v11 != 0 {
					v26 = F__emscripten_memcpy_bulkmem(m, v14+int32(4), v23+v24, v11)
					mBase = m.M
				} else {
				}
				return v14
			}
		}
	}
}
