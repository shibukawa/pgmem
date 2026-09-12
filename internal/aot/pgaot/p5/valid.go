package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_check_valid_oidvector(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v2 != int32(1) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			F_errcode(m, int32(67141764))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				F_errmsg(m, int32(198511), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					F_errfinish(m, int32(477992), int32(131), int32(198489))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
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
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v5 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				F_errcode(m, int32(67141764))
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					F_errmsg(m, int32(198511), int32(0))
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return
					} else {
						F_errfinish(m, int32(477992), int32(131), int32(198489))
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
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
			v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v6 == int32(26) {
				return
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v12 = m.ExcPending
				if v12 != 0 {
					return
				} else {
					F_errcode(m, int32(67141764))
					mBase = m.M
					v15 = m.ExcPending
					if v15 != 0 {
						return
					} else {
						F_errmsg(m, int32(198511), int32(0))
						mBase = m.M
						v19 = m.ExcPending
						if v19 != 0 {
							return
						} else {
							F_errfinish(m, int32(477992), int32(131), int32(198489))
							mBase = m.M
							v24 = m.ExcPending
							if v24 != 0 {
								return
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
