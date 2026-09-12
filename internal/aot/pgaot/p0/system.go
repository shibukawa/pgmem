package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetSystemIdentifier(m *base.Module) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, _consts[178]))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	return v3
}
func F_SystemAttributeDefinition(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	if base.Ui32(l0) <= base.Ui32(int32(-7)) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
			F_errmsg_internal(m, int32(471902), v5)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(495987), int32(239), int32(250564))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v30 = *(*int32)(unsafe.Add(mBase, uint32((l0^int32(-1))<<(uint(int32(2))%32))+uint32(_consts[199])))
		m.G0 = v5 + int32(16)
		return v30
	}
}
func F_system_beginsamplescan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 float32
	_ = v7
	var v12 float64
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v28 float64
	_ = v28
	var v34 int64
	_ = v34
	var v36 int64
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	v7 = *(*float32)(unsafe.Add(mBase, uint32(l1)))
	if base.F32_lt(v7, float32(0)) != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v47 = m.ExcPending
		if v47 != 0 {
			return
		} else {
			F_errcode(m, int32(403177602))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return
			} else {
				F_errmsg(m, int32(566728), int32(0))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return
				} else {
					F_errfinish(m, int32(497098), int32(151), int32(284492))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		if base.F32_gt(v7, float32(100)) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return
			} else {
				F_errcode(m, int32(403177602))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return
				} else {
					F_errmsg(m, int32(566728), int32(0))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return
					} else {
						F_errfinish(m, int32(497098), int32(151), int32(284492))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v12 = base.F64_promote_f32(v7)
			if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return
				} else {
					F_errcode(m, int32(403177602))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						F_errmsg(m, int32(566728), int32(0))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							F_errfinish(m, int32(497098), int32(151), int32(284492))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v19 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v18)+16)) = uint16(v19)
				*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v19
				*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = l3
				v28 = base.F64_nearest(base.F64_div(base.F64_mul(v12, float64(4.294967296e+09)), float64(100)))
				if base.F64_lt(v28, float64(1.8446744073709552e+19))&base.F64_ge(v28, float64(0)) != 0 {
					v34 = base.I64_trunc_f64_u(v28)
					v36 = v34
				} else {
					v36 = int64(0)
				}
				*(*int64)(unsafe.Add(mBase, uint32(v18))) = v36
				v38 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+133)) = uint8(v38)
				v41 = base.F32_ge(v7, float32(1))
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+132)) = uint8(v41)
				return
			}
		}
	}
}
