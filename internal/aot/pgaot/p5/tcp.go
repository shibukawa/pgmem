package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_assign_tcp_keepalives_interval(m *base.Module, l0 int32, l1 int32) {
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
	v4 = *(*int32)(unsafe.Add(mBase, _consts[381]))
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l0
	if v4 == int32(0) {
	} else {
		v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4)+12)))
		if v13 == int32(1) {
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v4)+404))
			if l0 == v16 {
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v4)+388))
				if int32(0) < v18 {
					if l0 == int32(0) {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v4)+388))
						*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v26
					} else {
					}
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v4)+404)) = v29
				} else {
					v21 = F_pq_getkeepalivesinterval(m, v4)
					mBase = m.M
					if int32(0) <= v21 {
						if l0 == int32(0) {
							v26 = *(*int32)(unsafe.Add(mBase, uint32(v4)+388))
							*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v26
						} else {
						}
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v4)+404)) = v29
					} else {
					}
				}
			}
		}
	}
	m.G0 = v8 + int32(16)
	return
}
func F_show_tcp_keepalives_interval(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	v1 = int32(0)
	v2 = m.G0
	v3 = int32(16)
	v4 = v2 - v3
	m.G0 = v4
	v7 = *(*int32)(unsafe.Add(mBase, _consts[381]))
	v10 = m.G0
	v12 = v10 - v3
	m.G0 = v12
	if v7 == v1 {
		v28 = v1
	} else {
		v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+12)))
		if v16 == int32(1) {
			v28 = v1
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v7)+404))
			if v19 != 0 {
				v28 = v19
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v7)+388))
				if v20 != 0 {
					v28 = v20
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(4)
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v7+int32(388))))
					v28 = v26
				}
			}
		}
	}
	m.G0 = v12 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v4))) = v28
	v36 = F_pg_snprintf(m, int32(4351008), int32(16), int32(467292), v4)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		return int32(0)
	} else {
		m.G0 = v4 + int32(16)
		return int32(4351008)
	}
}
func F_show_tcp_user_timeout(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	v1 = int32(0)
	v2 = m.G0
	v3 = int32(16)
	v4 = v2 - v3
	m.G0 = v4
	v7 = *(*int32)(unsafe.Add(mBase, _consts[381]))
	v10 = m.G0
	v12 = v10 - v3
	m.G0 = v12
	if v7 == v1 {
		v28 = v1
	} else {
		v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+12)))
		if v16 == int32(1) {
			v28 = v1
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v7)+412))
			if v19 != 0 {
				v28 = v19
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v7)+396))
				if v20 != 0 {
					v28 = v20
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(4)
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v7+int32(396))))
					v28 = v26
				}
			}
		}
	}
	m.G0 = v12 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v4))) = v28
	v36 = F_pg_snprintf(m, int32(4351040), int32(16), int32(467292), v4)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		return int32(0)
	} else {
		m.G0 = v4 + int32(16)
		return int32(4351040)
	}
}
