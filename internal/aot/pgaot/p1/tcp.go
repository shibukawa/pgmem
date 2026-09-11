package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_assign_tcp_keepalives_count(m *base.Module, l0 int32, l1 int32) {
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_assign_tcp_keepalives_count[0]))
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l0
	if v4 == int32(0) {
	} else {
		v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4)+12)))
		if v13 == int32(1) {
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v4)+408))
			if l0 == v16 {
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v4)+392))
				if int32(0) < v18 {
					if l0 == int32(0) {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v4)+392))
						*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v26
					} else {
					}
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v4)+408)) = v29
				} else {
					v21 = F_pq_getkeepalivescount(m, v4)
					mBase = m.M
					if int32(0) <= v21 {
						if l0 == int32(0) {
							v26 = *(*int32)(unsafe.Add(mBase, uint32(v4)+392))
							*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v26
						} else {
						}
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v4)+408)) = v29
					} else {
					}
				}
			}
		}
	}
	m.G0 = v8 + int32(16)
	return
}
func F_assign_tcp_user_timeout(m *base.Module, l0 int32, l1 int32) {
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_assign_tcp_user_timeout[0]))
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l0
	if v4 == int32(0) {
	} else {
		v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4)+12)))
		if v13 == int32(1) {
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v4)+412))
			if l0 == v16 {
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v4)+396))
				if int32(0) < v18 {
					if l0 == int32(0) {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v4)+396))
						*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v26
					} else {
					}
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v4)+412)) = v29
				} else {
					v21 = F_pq_gettcpusertimeout(m, v4)
					mBase = m.M
					if int32(0) <= v21 {
						if l0 == int32(0) {
							v26 = *(*int32)(unsafe.Add(mBase, uint32(v4)+396))
							*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v26
						} else {
						}
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v4)+412)) = v29
					} else {
					}
				}
			}
		}
	}
	m.G0 = v8 + int32(16)
	return
}
