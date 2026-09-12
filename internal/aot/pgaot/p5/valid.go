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
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int64
	_ = v12
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+68)) = v3
	v12 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+60)) = v12
	*(*int64)(unsafe.Add(mBase, uint32(v8)+52)) = v12
	*(*int64)(unsafe.Add(mBase, uint32(v8)+44)) = v12
	*(*int64)(unsafe.Add(mBase, uint32(v8)+36)) = v12
	*(*int64)(unsafe.Add(mBase, uint32(v8)+28)) = v12
	*(*int64)(unsafe.Add(mBase, uint32(v8)+20)) = v12
	*(*int64)(unsafe.Add(mBase, uint32(v8)+12)) = v12
	if l1 != 0 {
		v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
		v28 = base.B2i32(v26 == int32(45))
		v29 = l0 + v28
		*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v29
		*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l1 - v28
		*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v29
		v40 = F_json_lex_number(m, v8+int32(4), v29, v8+int32(79), v8+int32(72))
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return int32(0)
		} else {
			v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+79)))
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v8)+72))
			v48 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
			v51 = (v44 ^ int32(-1)) & base.B2i32(v47 == v48)
			m.G0 = v8 + int32(80)
			return v51
		}
	} else {
		v51 = v3
		m.G0 = v8 + int32(80)
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
				F_errmsg(m, int32(218334), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					F_errfinish(m, int32(522856), int32(131), int32(218312))
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
					F_errmsg(m, int32(218334), int32(0))
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return
					} else {
						F_errfinish(m, int32(522856), int32(131), int32(218312))
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
						F_errmsg(m, int32(218334), int32(0))
						mBase = m.M
						v19 = m.ExcPending
						if v19 != 0 {
							return
						} else {
							F_errfinish(m, int32(522856), int32(131), int32(218312))
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
