package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_socket(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v17 = m.Env.X__syscall_socket(m, l0, l1, l2, v4, v4, v4)
	mBase = m.M
	if base.B2i32(l1&int32(_a_F_socket_0) == v4)|base.B2i32(v17 != int32(-28))&base.B2i32(v17 != int32(-66)) != 0 {
		v48 = v17
	} else {
		v26 = int32(0)
		v29 = m.Env.X__syscall_socket(m, l0, l1&int32(-526337), l2, v26, v26, v26)
		mBase = m.M
		if v29 < v26 {
			v48 = v29
		} else {
			if l1&int32(_a_F_socket_1) != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = int32(1)
				v39 = m.Env.X__syscall_fcntl64(m, v29, int32(2), v8+int32(16))
				mBase = m.M
			} else {
			}
			if l1&int32(2048) == int32(0) {
				v48 = v29
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(2048)
				v47 = m.Env.X__syscall_fcntl64(m, v29, int32(4), v8)
				mBase = m.M
				v48 = v29
			}
		}
	}
	if base.Ui32(int32(-4095)) <= base.Ui32(v48) {
		*(*int32)(unsafe.Add(mBase, _c_F_socket[0])) = int32(0) - v48
		v56 = int32(-1)
	} else {
		v56 = v48
	}
	m.G0 = v8 + int32(32)
	return v56
}
func F_socket_flush_if_writable(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	v1 = int32(0)
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_socket_flush_if_writable[0]))
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_socket_flush_if_writable[1]))
	if v3 == v5 {
		v29 = v1
		return v29
	} else {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_socket_flush_if_writable[2])))
		if v8 != 0 {
			v29 = v1
			return v29
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, _c_F_socket_flush_if_writable[3]))
			if v10 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50332160))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_socket_flush_if_writable_0), int32(0))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_socket_flush_if_writable_1), int32(886), int32(_a_F_socket_flush_if_writable_2))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
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
				v13 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v10)+4)) = uint8(v13)
				*(*uint8)(unsafe.Add(mBase, _c_F_socket_flush_if_writable[2])) = uint8(v13)
				v19 = *(*int32)(unsafe.Add(mBase, _c_F_socket_flush_if_writable[4]))
				v22 = F_internal_flush_buffer(m, v19, int32(_a_F_socket_flush_if_writable_3), int32(_a_F_socket_flush_if_writable_4))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v27 = int32(0)
					*(*uint8)(unsafe.Add(mBase, _c_F_socket_flush_if_writable[2])) = uint8(v27)
					v29 = v22
					return v29
				}
			}
		}
	}
}
func F_socket_is_send_pending(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_socket_is_send_pending[0]))
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_socket_is_send_pending[1]))
	return base.B2i32(base.Ui32(v2) < base.Ui32(v4))
}
