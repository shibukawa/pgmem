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
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_socket_close[0]))
	if v4 != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, _c_F_socket_close[0]))
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
	v3 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_socket_flush[0])))
	if v3 == v1 {
		v7 = int32(1)
		*(*uint8)(unsafe.Add(mBase, _c_F_socket_flush[0])) = uint8(v7)
		v10 = *(*int32)(unsafe.Add(mBase, _c_F_socket_flush[1]))
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
					F_errmsg(m, int32(_a_F_socket_flush_0), int32(0))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_socket_flush_1), int32(886), int32(_a_F_socket_flush_2))
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
			v16 = *(*int32)(unsafe.Add(mBase, _c_F_socket_flush[2]))
			v19 = F_internal_flush_buffer(m, v16, int32(_a_F_socket_flush_3), int32(_a_F_socket_flush_4))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v24 = int32(0)
				*(*uint8)(unsafe.Add(mBase, _c_F_socket_flush[0])) = uint8(v24)
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	v1 = l0
	v4 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v1)
	v12 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_socket_putmessage[0])))
	if v12 == v4 {
		v16 = int32(1)
		*(*uint8)(unsafe.Add(mBase, _c_F_socket_putmessage[0])) = uint8(v16)
		v21 = F_internal_putbytes(m, v7+int32(15), v16)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			if v21 != 0 {
				v48 = int32(-1)
				v50 = int32(0)
				*(*uint8)(unsafe.Add(mBase, _c_F_socket_putmessage[0])) = uint8(v50)
				v52 = v48
				m.G0 = v7 + int32(16)
				return v52
			} else {
				v25 = int32(4)
				v26 = l2 + v25
				v27 = int32(16711935)
				v29 = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = base.I32_rotr(v26&v27, v29) | base.I32_rotr(v26, int32(24))&v27
				v40 = F_internal_putbytes(m, v7+v29, v25)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					if v40 != 0 {
						v48 = int32(-1)
						v50 = int32(0)
						*(*uint8)(unsafe.Add(mBase, _c_F_socket_putmessage[0])) = uint8(v50)
						v52 = v48
						m.G0 = v7 + int32(16)
						return v52
					} else {
						v42 = F_internal_putbytes(m, l1, l2)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							if v42 != 0 {
								v48 = int32(-1)
							} else {
								v48 = int32(0)
							}
							v50 = int32(0)
							*(*uint8)(unsafe.Add(mBase, _c_F_socket_putmessage[0])) = uint8(v50)
							v52 = v48
							m.G0 = v7 + int32(16)
							return v52
						}
					}
				}
			}
		}
	} else {
		v52 = v4
		m.G0 = v7 + int32(16)
		return v52
	}
}
