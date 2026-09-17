package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_IsValidJsonNumber(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int64
	_ = v11
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(80)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+68)) = v3
	v11 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+60)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v7)+52)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v7)+44)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v7)+36)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v7)+28)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v7)+20)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v7)+12)) = v11
	if l1 != 0 {
		v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
		v27 = base.B2i32(v25 == int32(45))
		v28 = l0 + v27
		*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v28
		*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l1 - v27
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v28
		v39 = F_json_lex_number(m, v7+int32(4), v28, v7+int32(79), v7+int32(72))
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return int32(0)
		} else {
			v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+79)))
			v46 = *(*int32)(unsafe.Add(mBase, uint32(v7)+72))
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
			v51 = (v43 ^ int32(-1)) & base.B2i32(v46 == v47)
			m.G0 = v7 + int32(80)
			return v51
		}
	} else {
		v51 = v3
		m.G0 = v7 + int32(80)
		return v51
	}
}
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
				F_errmsg(m, int32(_a_F_check_valid_oidvector_0), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_check_valid_oidvector_1), int32(131), int32(_a_F_check_valid_oidvector_2))
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
					F_errmsg(m, int32(_a_F_check_valid_oidvector_0), int32(0))
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_check_valid_oidvector_1), int32(131), int32(_a_F_check_valid_oidvector_2))
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
						F_errmsg(m, int32(_a_F_check_valid_oidvector_0), int32(0))
						mBase = m.M
						v19 = m.ExcPending
						if v19 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_check_valid_oidvector_1), int32(131), int32(_a_F_check_valid_oidvector_2))
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
