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
	var v13 int32
	_ = v13
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v13 = m.Env.X__syscall_socket(m, l0, l1, l2, v4, v4, v4)
	mBase = m.M
	if base.B2i32(v13 != int32(-28))&base.B2i32(v13 != int32(-66)) != 0 {
		v47 = v13
	} else {
		if l1&int32(526336) == int32(0) {
			v47 = v13
		} else {
			v25 = int32(0)
			v28 = m.Env.X__syscall_socket(m, l0, l1&int32(-526337), l2, v25, v25, v25)
			mBase = m.M
			if v28 < v25 {
				v47 = v28
			} else {
				if l1&int32(524288) != 0 {
					*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = int64(1)
					v38 = m.Env.X__syscall_fcntl64(m, v28, int32(2), v8+int32(16))
					mBase = m.M
				} else {
				}
				if l1&int32(2048) == int32(0) {
					v47 = v28
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(2048)
					v46 = m.Env.X__syscall_fcntl64(m, v28, int32(4), v8)
					mBase = m.M
					v47 = v28
				}
			}
		}
	}
	if base.Ui32(int32(-4095)) <= base.Ui32(v47) {
		*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(0) - v47
		v55 = int32(-1)
	} else {
		v55 = v47
	}
	m.G0 = v8 + int32(32)
	return v55
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
	v3 = *(*int32)(unsafe.Add(mBase, _consts[456]))
	v5 = *(*int32)(unsafe.Add(mBase, _consts[457]))
	if v3 == v5 {
		v29 = v1
		return v29
	} else {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, _consts[458])))
		if v8 != 0 {
			v29 = v1
			return v29
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, _consts[454]))
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
						F_errmsg(m, int32(254881), int32(0))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(496855), int32(886), int32(335324))
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
				*(*uint8)(unsafe.Add(mBase, _consts[458])) = uint8(v13)
				v19 = *(*int32)(unsafe.Add(mBase, _consts[459]))
				v22 = F_internal_flush_buffer(m, v19, int32(4414040), int32(4414044))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v27 = int32(0)
					*(*uint8)(unsafe.Add(mBase, _consts[458])) = uint8(v27)
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
	v2 = *(*int32)(unsafe.Add(mBase, _consts[457]))
	v4 = *(*int32)(unsafe.Add(mBase, _consts[456]))
	return base.B2i32(base.Ui32(v2) < base.Ui32(v4))
}
