package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_socket_close(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v4 = *(*int32)(unsafe.Add(mBase, _consts[475]))
	if v4 != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, _consts[475]))
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(-1)
	} else {
	}
	return
}
func F_socket_flush(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	v1 = int32(0)
	v3 = int32(*(*uint8)(unsafe.Add(mBase, _consts[481])))
	if v3 == v1 {
		v7 = int32(1)
		*(*uint8)(unsafe.Add(mBase, _consts[481])) = uint8(v7)
		v10 = *(*int32)(unsafe.Add(mBase, _consts[475]))
		if v10 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50332160))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(255169), int32(0))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(497281), int32(886), int32(335637))
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
		} else {
			v13 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v10)+4)) = uint8(v13)
			v16 = *(*int32)(unsafe.Add(mBase, _consts[479]))
			v19 = F_internal_flush_buffer(m, v16, int32(4414408), int32(4414412))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v24 = int32(0)
				*(*uint8)(unsafe.Add(mBase, _consts[481])) = uint8(v24)
				v26 = v19
				return v26
			}
		}
	} else {
		v26 = v1
		return v26
	}
}
func F_socket_putmessage(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	v1 = l0
	v4 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v1)
	v12 = int32(*(*uint8)(unsafe.Add(mBase, _consts[481])))
	if v12 == v4 {
		v16 = int32(1)
		*(*uint8)(unsafe.Add(mBase, _consts[481])) = uint8(v16)
		v21 = F_internal_putbytes(m, v7+int32(15), v16)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			if v21 != 0 {
				v54 = int32(-1)
				v56 = int32(0)
				*(*uint8)(unsafe.Add(mBase, _consts[481])) = uint8(v56)
				v58 = v54
				m.G0 = v7 + int32(16)
				return v58
			} else {
				v25 = int32(4)
				v26 = l2 + v25
				v27 = int32(24)
				v29 = int32(65280)
				v31 = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v26<<(uint(v27)%32) | v26&v29<<(uint(v31)%32) | (int32(base.Ui32(v26)>>(uint(v31)%32))&v29 | int32(base.Ui32(v26)>>(uint(v27)%32)))
				v46 = F_internal_putbytes(m, v7+v31, v25)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					if v46 != 0 {
						v54 = int32(-1)
						v56 = int32(0)
						*(*uint8)(unsafe.Add(mBase, _consts[481])) = uint8(v56)
						v58 = v54
						m.G0 = v7 + int32(16)
						return v58
					} else {
						v48 = F_internal_putbytes(m, l1, l2)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							if v48 != 0 {
								v54 = int32(-1)
							} else {
								v54 = int32(0)
							}
							v56 = int32(0)
							*(*uint8)(unsafe.Add(mBase, _consts[481])) = uint8(v56)
							v58 = v54
							m.G0 = v7 + int32(16)
							return v58
						}
					}
				}
			}
		}
	} else {
		v58 = v4
		m.G0 = v7 + int32(16)
		return v58
	}
}
